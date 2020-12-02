// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// AccountGetNotifyExceptionsRequest represents TL type `account.getNotifyExceptions#53577479`.
type AccountGetNotifyExceptionsRequest struct {
	// Flags field of AccountGetNotifyExceptionsRequest.
	Flags bin.Fields
	// CompareSound field of AccountGetNotifyExceptionsRequest.
	CompareSound bool
	// Peer field of AccountGetNotifyExceptionsRequest.
	//
	// Use SetPeer and GetPeer helpers.
	Peer InputNotifyPeerClass
}

// AccountGetNotifyExceptionsRequestTypeID is TL type id of AccountGetNotifyExceptionsRequest.
const AccountGetNotifyExceptionsRequestTypeID = 0x53577479

// Encode implements bin.Encoder.
func (g *AccountGetNotifyExceptionsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getNotifyExceptions#53577479 as nil")
	}
	b.PutID(AccountGetNotifyExceptionsRequestTypeID)
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.getNotifyExceptions#53577479: field flags: %w", err)
	}
	if g.Flags.Has(0) {
		if g.Peer == nil {
			return fmt.Errorf("unable to encode account.getNotifyExceptions#53577479: field peer is nil")
		}
		if err := g.Peer.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.getNotifyExceptions#53577479: field peer: %w", err)
		}
	}
	return nil
}

// SetCompareSound sets value of CompareSound conditional field.
func (g *AccountGetNotifyExceptionsRequest) SetCompareSound(value bool) {
	if value {
		g.Flags.Set(1)
	} else {
		g.Flags.Unset(1)
	}
}

// SetPeer sets value of Peer conditional field.
func (g *AccountGetNotifyExceptionsRequest) SetPeer(value InputNotifyPeerClass) {
	g.Flags.Set(0)
	g.Peer = value
}

// GetPeer returns value of Peer conditional field and
// boolean which is true if field was set.
func (g *AccountGetNotifyExceptionsRequest) GetPeer() (value InputNotifyPeerClass, ok bool) {
	if !g.Flags.Has(0) {
		return value, false
	}
	return g.Peer, true
}

// Decode implements bin.Decoder.
func (g *AccountGetNotifyExceptionsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getNotifyExceptions#53577479 to nil")
	}
	if err := b.ConsumeID(AccountGetNotifyExceptionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getNotifyExceptions#53577479: %w", err)
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.getNotifyExceptions#53577479: field flags: %w", err)
		}
	}
	g.CompareSound = g.Flags.Has(1)
	if g.Flags.Has(0) {
		value, err := DecodeInputNotifyPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.getNotifyExceptions#53577479: field peer: %w", err)
		}
		g.Peer = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountGetNotifyExceptionsRequest.
var (
	_ bin.Encoder = &AccountGetNotifyExceptionsRequest{}
	_ bin.Decoder = &AccountGetNotifyExceptionsRequest{}
)

// AccountGetNotifyExceptions invokes method account.getNotifyExceptions#53577479 returning error if any.
func (c *Client) AccountGetNotifyExceptions(ctx context.Context, request *AccountGetNotifyExceptionsRequest) (UpdatesClass, error) {
	var result UpdatesBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}