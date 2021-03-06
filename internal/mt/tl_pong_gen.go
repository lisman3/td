// Code generated by gotdgen, DO NOT EDIT.

package mt

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is

// Pong represents TL type `pong#347773c5`.
type Pong struct {
	// MsgID field of Pong.
	MsgID int64
	// PingID field of Pong.
	PingID int64
}

// PongTypeID is TL type id of Pong.
const PongTypeID = 0x347773c5

func (p *Pong) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.MsgID == 0) {
		return false
	}
	if !(p.PingID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *Pong) String() string {
	if p == nil {
		return "Pong(nil)"
	}
	var sb strings.Builder
	sb.WriteString("Pong")
	sb.WriteString("{\n")
	sb.WriteString("\tMsgID: ")
	sb.WriteString(fmt.Sprint(p.MsgID))
	sb.WriteString(",\n")
	sb.WriteString("\tPingID: ")
	sb.WriteString(fmt.Sprint(p.PingID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *Pong) TypeID() uint32 {
	return PongTypeID
}

// Encode implements bin.Encoder.
func (p *Pong) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pong#347773c5 as nil")
	}
	b.PutID(PongTypeID)
	b.PutLong(p.MsgID)
	b.PutLong(p.PingID)
	return nil
}

// GetMsgID returns value of MsgID field.
func (p *Pong) GetMsgID() (value int64) {
	return p.MsgID
}

// GetPingID returns value of PingID field.
func (p *Pong) GetPingID() (value int64) {
	return p.PingID
}

// Decode implements bin.Decoder.
func (p *Pong) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pong#347773c5 to nil")
	}
	if err := b.ConsumeID(PongTypeID); err != nil {
		return fmt.Errorf("unable to decode pong#347773c5: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode pong#347773c5: field msg_id: %w", err)
		}
		p.MsgID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode pong#347773c5: field ping_id: %w", err)
		}
		p.PingID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for Pong.
var (
	_ bin.Encoder = &Pong{}
	_ bin.Decoder = &Pong{}
)
