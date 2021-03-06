package codec

import (
	"encoding/binary"
	"fmt"
	"io"

	"golang.org/x/xerrors"

	"github.com/gotd/td/bin"
)

// AbridgedClientStart is starting bytes sent by client in Abridged mode.
//
// Note that server does not respond with it.
var AbridgedClientStart = [1]byte{0xef}

// Abridged is intermediate MTProto transport.
//
// See https://core.telegram.org/mtproto/mtproto-transports#abridged
type Abridged struct{}

// WriteHeader sends protocol tag.
func (i Abridged) WriteHeader(w io.Writer) error {
	if _, err := w.Write(AbridgedClientStart[:]); err != nil {
		return xerrors.Errorf("write abridged header: %w", err)
	}

	return nil
}

// ReadHeader reads protocol tag.
func (i Abridged) ReadHeader(r io.Reader) error {
	var b [1]byte
	if _, err := r.Read(b[:]); err != nil {
		return xerrors.Errorf("read abridged header: %w", err)
	}

	if b != AbridgedClientStart {
		return ErrProtocolHeaderMismatch
	}

	return nil
}

// ObfuscatedTag returns protocol tag for obfuscation.
func (i Abridged) ObfuscatedTag() (r [4]byte) {
	return [4]byte{0xef, 0xef, 0xef, 0xef}
}

// Write encode to writer message from given buffer.
func (i Abridged) Write(w io.Writer, b *bin.Buffer) error {
	if err := checkOutgoingMessage(b); err != nil {
		return err
	}

	if err := checkAlign(b, 4); err != nil {
		return err
	}

	if err := writeAbridged(w, b); err != nil {
		return xerrors.Errorf("write abridged: %w", err)
	}

	return nil
}

// Read fills buffer with received message.
func (i Abridged) Read(r io.Reader, b *bin.Buffer) error {
	if err := readAbridged(r, b); err != nil {
		return xerrors.Errorf("read abridged: %w", err)
	}

	if err := checkProtocolError(b); err != nil {
		return err
	}

	return nil
}

func writeAbridged(w io.Writer, b *bin.Buffer) error {
	length := b.Len() >> 2

	// Re-using b.Buf if possible to reduce allocations.
	buf := append(b.Buf[len(b.Buf):], make([]byte, 0, 4)...)
	inner := bin.Buffer{Buf: buf}

	if length < 127 {
		inner.Put([]byte{byte(length)})
	} else {
		var buf [5]byte
		buf[0] = 0x7f
		binary.LittleEndian.PutUint32(buf[1:], uint32(length))
		inner.Put(buf[:4])
	}

	if _, err := w.Write(inner.Buf); err != nil {
		return err
	}
	if _, err := w.Write(b.Raw()); err != nil {
		return err
	}
	return nil
}

func readAbridged(r io.Reader, b *bin.Buffer) error {
	b.ResetN(bin.Word)

	_, err := io.ReadFull(r, b.Buf[:1])
	if err != nil {
		return err
	}

	if b.Buf[0] >= 127 {
		_, err := io.ReadFull(r, b.Buf[0:3])
		if err != nil {
			return err
		}
	}

	n, err := b.Int()
	if err != nil {
		return err
	}

	b.ResetN(n << 2)
	if _, err := io.ReadFull(r, b.Buf); err != nil {
		return fmt.Errorf("failed to read payload: %w", err)
	}

	return nil
}
