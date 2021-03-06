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

// MsgsStateReq represents TL type `msgs_state_req#da69fb52`.
type MsgsStateReq struct {
	// MsgIds field of MsgsStateReq.
	MsgIds []int64
}

// MsgsStateReqTypeID is TL type id of MsgsStateReq.
const MsgsStateReqTypeID = 0xda69fb52

func (m *MsgsStateReq) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.MsgIds == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MsgsStateReq) String() string {
	if m == nil {
		return "MsgsStateReq(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MsgsStateReq")
	sb.WriteString("{\n")
	sb.WriteByte('[')
	for _, v := range m.MsgIds {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (m *MsgsStateReq) TypeID() uint32 {
	return MsgsStateReqTypeID
}

// Encode implements bin.Encoder.
func (m *MsgsStateReq) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode msgs_state_req#da69fb52 as nil")
	}
	b.PutID(MsgsStateReqTypeID)
	b.PutVectorHeader(len(m.MsgIds))
	for _, v := range m.MsgIds {
		b.PutLong(v)
	}
	return nil
}

// GetMsgIds returns value of MsgIds field.
func (m *MsgsStateReq) GetMsgIds() (value []int64) {
	return m.MsgIds
}

// Decode implements bin.Decoder.
func (m *MsgsStateReq) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode msgs_state_req#da69fb52 to nil")
	}
	if err := b.ConsumeID(MsgsStateReqTypeID); err != nil {
		return fmt.Errorf("unable to decode msgs_state_req#da69fb52: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode msgs_state_req#da69fb52: field msg_ids: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode msgs_state_req#da69fb52: field msg_ids: %w", err)
			}
			m.MsgIds = append(m.MsgIds, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MsgsStateReq.
var (
	_ bin.Encoder = &MsgsStateReq{}
	_ bin.Decoder = &MsgsStateReq{}
)
