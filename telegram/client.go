package telegram

import (
	"context"
	"errors"
	"fmt"
	"io"
	"runtime/debug"
	"strings"
	"sync"

	"github.com/cenkalti/backoff/v4"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"golang.org/x/xerrors"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/clock"
	"github.com/gotd/td/internal/crypto"
	"github.com/gotd/td/internal/mt"
	"github.com/gotd/td/internal/proto"
	"github.com/gotd/td/internal/tdsync"
	"github.com/gotd/td/internal/tmap"
	"github.com/gotd/td/mtproto"
	"github.com/gotd/td/session"
	"github.com/gotd/td/tg"
)

// UpdateHandler will be called on received updates from Telegram.
type UpdateHandler interface {
	Handle(ctx context.Context, u *tg.Updates) error
	HandleShort(ctx context.Context, u *tg.UpdateShort) error
}

// Available MTProto default server addresses.
//
// See https://my.telegram.org/apps.
const (
	AddrProduction = "149.154.167.50:443"
	AddrTest       = "149.154.167.40:443"
)

// Test-only credentials. Can be used with AddrTest and TestAuth to
// test authentication.
//
// Reference:
//	* https://github.com/telegramdesktop/tdesktop/blob/5f665b8ecb48802cd13cfb48ec834b946459274a/docs/api_credentials.md
const (
	TestAppID   = 17349
	TestAppHash = "344583e45741c457fe1862106095a5eb"
)

type clientStorage interface {
	Load(ctx context.Context) (*session.Data, error)
	Save(ctx context.Context, data *session.Data) error
}

type clientConn interface {
	Run(ctx context.Context) error
	InvokeRaw(ctx context.Context, input bin.Encoder, output bin.Decoder) error
}

// Client represents a MTProto client to Telegram.
type Client struct {
	// tg provides RPC calls via Client.
	tg *tg.Client // immutable

	// Telegram device information.
	device DeviceConfig // immutable

	// MTProto options.
	opts mtproto.Options // immutable

	// Connection state. Guarded by connMux.
	session   *syncSessionData
	cfg       *atomicConfig
	conn      clientConn
	connMux   sync.Mutex
	primaryDC atomic.Int64

	// Restart signal channel.
	restart  chan struct{}                      // immutable
	exported chan *tg.AuthExportedAuthorization // immutable

	// Wrappers for external world, like logs or PRNG.
	rand  io.Reader   // immutable
	log   *zap.Logger // immutable
	clock clock.Clock // immutable

	// Client context. Will be canceled by Run on exit.
	ctx    context.Context    // immutable
	cancel context.CancelFunc // immutable

	// Client config.
	appID   int    // immutable
	appHash string // immutable
	// Session storage.
	storage clientStorage // immutable,nilable

	// Ready signal channel, sends signal when client connection is ready.
	// Resets on reconnect.
	ready *tdsync.ResetReady // immutable

	connsCounter atomic.Int64

	// Telegram updates handler.
	updateHandler UpdateHandler // immutable
}

func (c *Client) onMessage(b *bin.Buffer) error {
	return c.handleUpdates(b)
}

// getVersion optimistically gets current client version.
//
// Does not handle replace directives.
func getVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return ""
	}
	// Hard-coded package name. Probably we can generate this via parsing
	// the go.mod file.
	const pkg = "github.com/gotd/td"
	for _, d := range info.Deps {
		if strings.HasPrefix(d.Path, pkg) {
			return d.Version
		}
	}
	return ""
}

// Port is default port used by telegram.
const Port = 443

// NewClient creates new unstarted client.
func NewClient(appID int, appHash string, opt Options) *Client {
	// Set default values, if user does not set.
	opt.setDefaults()

	clientCtx, clientCancel := context.WithCancel(context.Background())
	client := &Client{
		rand:          opt.Random,
		log:           opt.Logger,
		ctx:           clientCtx,
		cancel:        clientCancel,
		appID:         appID,
		appHash:       appHash,
		updateHandler: opt.UpdateHandler,
		session: newSyncSessionData(sessionData{
			Addr: opt.Addr,
			DC:   opt.DC,
		}),
		cfg:       &atomicConfig{},
		primaryDC: *atomic.NewInt64(int64(opt.DC)),
		clock:     opt.Clock,
		device:    opt.Device,
		ready:     tdsync.NewResetReady(),
		restart:   make(chan struct{}),
		exported:  make(chan *tg.AuthExportedAuthorization, 1),
	}

	// Including version into client logger to help with debugging.
	if v := getVersion(); v != "" {
		client.log = client.log.With(zap.String("v", v))
	}

	if opt.SessionStorage != nil {
		client.storage = &session.Loader{
			Storage: opt.SessionStorage,
		}
	}

	client.opts = mtproto.Options{
		PublicKeys:    opt.PublicKeys,
		Transport:     opt.Transport,
		Network:       opt.Network,
		Random:        opt.Random,
		Logger:        opt.Logger,
		AckBatchSize:  opt.AckBatchSize,
		AckInterval:   opt.AckInterval,
		RetryInterval: opt.RetryInterval,
		MaxRetries:    opt.MaxRetries,
		MessageID:     opt.MessageID,
		Clock:         opt.Clock,

		Types: tmap.New(
			tg.TypesMap(),
			mt.TypesMap(),
			proto.TypesMap(),
		),
	}
	client.conn = client.createConn(connModeUpdates)

	// Initializing internal RPC caller.
	client.tg = tg.NewClient(client)

	return client
}

func (c *Client) restoreConnection(ctx context.Context) error {
	if c.storage == nil {
		return nil
	}
	data, err := c.storage.Load(ctx)
	if errors.Is(err, session.ErrNotFound) {
		return nil
	}
	if err != nil {
		return xerrors.Errorf("load: %w", err)
	}

	// Restoring persisted auth key.
	var key crypto.AuthKey
	copy(key.Value[:], data.AuthKey)
	copy(key.ID[:], data.AuthKeyID)

	if key.Value.ID() != key.ID {
		return xerrors.New("corrupted key")
	}

	// Re-initializing connection from persisted state.
	c.log.Info("Connection restored from state",
		zap.String("addr", data.Addr),
		zap.String("key_id", fmt.Sprintf("%x", data.AuthKeyID)),
	)

	c.connMux.Lock()
	c.session.Store(sessionData{
		DC:      data.DC,
		Addr:    data.Addr,
		AuthKey: key,
		Salt:    data.Salt,
	})
	c.primaryDC.Store(int64(data.DC))
	c.conn = c.createConn(connModeUpdates)
	c.connMux.Unlock()

	return nil
}

func (c *Client) runUntilRestart(ctx context.Context) error {
	g := tdsync.NewCancellableGroup(ctx)
	g.Go(c.conn.Run)
	g.Go(func(groupCtx context.Context) error {
		_, err := c.tg.HelpGetConfig(ctx)
		if err != nil {
			c.log.Warn("Get config failed", zap.Error(err))
		}

		return nil
	})
	g.Go(func(gCtx context.Context) error {
		select {
		case <-gCtx.Done():
			return gCtx.Err()
		case <-c.restart:
			c.log.Debug("Restart triggered")
			// Should call cancel() to cancel group.
			g.Cancel()

			return nil
		}
	})

	return g.Wait()
}

func (c *Client) reconnectUntilClosed(ctx context.Context) error {
	// TODO: Make this configurable.
	// Note that we currently have no timeout on connection, so this is
	// potentially eternal.
	b := backoff.NewExponentialBackOff()
	b.Clock = c.clock
	b.MaxElapsedTime = 0

	for {
		err := c.runUntilRestart(ctx)
		if err == nil {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-c.clock.After(b.NextBackOff()):
			c.log.Info("Restarting connection", zap.Error(err))

			c.connMux.Lock()
			c.conn = c.buildConn(connModeUpdates).WithSetup(func(ctx context.Context, invoker tg.Invoker) error {
				select {
				case export := <-c.exported:
					_, err := tg.NewClient(invoker).AuthImportAuthorization(ctx, &tg.AuthImportAuthorizationRequest{
						ID:    export.ID,
						Bytes: export.Bytes,
					})
					return err
				default:
				}
				return nil
			}).Build()
			c.connMux.Unlock()
		}
	}
}

func (c *Client) onReady() {
	c.log.Debug("Ready")
	c.ready.Signal()
}

func (c *Client) resetReady() {
	c.ready.Reset()
}

// Run starts client session and block until connection close.
// The f callback is called on successful session initialization and Run
// will return on f() result.
//
// Context of callback will be canceled if fatal error is detected.
func (c *Client) Run(ctx context.Context, f func(ctx context.Context) error) error {
	select {
	case <-c.ctx.Done():
		return xerrors.Errorf("client already closed: %w", c.ctx.Err())
	default:
	}

	c.log.Info("Starting")
	defer c.log.Info("Closed")
	// Cancel client on exit.
	defer c.cancel()

	c.resetReady()
	if err := c.restoreConnection(ctx); err != nil {
		return err
	}

	g := tdsync.NewCancellableGroup(ctx)
	g.Go(c.reconnectUntilClosed)
	g.Go(func(gCtx context.Context) error {
		select {
		case <-gCtx.Done():
			return gCtx.Err()
		case <-c.ready.Ready():
			if err := f(gCtx); err != nil {
				return xerrors.Errorf("callback: %w", err)
			}
			// Should call cancel() to cancel gCtx.
			// This will terminate c.conn.Run().
			c.log.Debug("Callback returned, stopping")
			g.Cancel()
			return nil
		}
	})
	if err := g.Wait(); !xerrors.Is(err, context.Canceled) {
		return err
	}

	return nil
}

func (c *Client) saveSession(addr string, cfg tg.Config, s mtproto.Session) error {
	if c.storage == nil {
		return nil
	}

	// Do not save session for non-primary DC.
	if cfg.ThisDC != 0 && c.primaryDC.Load() != int64(cfg.ThisDC) {
		return nil
	}

	data, err := c.storage.Load(c.ctx)
	if errors.Is(err, session.ErrNotFound) {
		// Initializing new state.
		err = nil
		data = &session.Data{}
	}
	if err != nil {
		return xerrors.Errorf("load: %w", err)
	}

	// Updating previous data.
	data.Config = cfg
	data.AuthKey = s.Key.Value[:]
	data.AuthKeyID = s.Key.ID[:]
	data.DC = cfg.ThisDC
	data.Addr = addr
	data.Salt = s.Salt

	if err := c.storage.Save(c.ctx, data); err != nil {
		return xerrors.Errorf("save: %w", err)
	}

	c.log.Debug("Data saved",
		zap.String("key_id", fmt.Sprintf("%x", data.AuthKeyID)),
	)
	return nil
}

func (c *Client) onSession(addr string, cfg tg.Config, s mtproto.Session) error {
	if err := c.saveSession(addr, cfg, s); err != nil {
		return xerrors.Errorf("save: %w", err)
	}

	c.connMux.Lock()
	c.session.Store(sessionData{
		DC:      cfg.ThisDC,
		Addr:    addr,
		Salt:    s.Salt,
		AuthKey: s.Key,
	})
	c.cfg.Store(cfg)
	c.onReady()
	c.connMux.Unlock()

	return nil
}

func (c *Client) createConn(mode connMode) *conn {
	return c.buildConn(mode).Build()
}
