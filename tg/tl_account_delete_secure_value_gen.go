// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// AccountDeleteSecureValueRequest represents TL type `account.deleteSecureValue#b880bc4b`.
// Delete stored Telegram Passport¹ documents, for more info see the passport docs »²
//
// Links:
//  1) https://core.telegram.org/passport
//  2) https://core.telegram.org/passport/encryption#encryption
//
// See https://core.telegram.org/method/account.deleteSecureValue for reference.
type AccountDeleteSecureValueRequest struct {
	// Document types to delete
	Types []SecureValueTypeClass
}

// AccountDeleteSecureValueRequestTypeID is TL type id of AccountDeleteSecureValueRequest.
const AccountDeleteSecureValueRequestTypeID = 0xb880bc4b

func (d *AccountDeleteSecureValueRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Types == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *AccountDeleteSecureValueRequest) String() string {
	if d == nil {
		return "AccountDeleteSecureValueRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AccountDeleteSecureValueRequest")
	sb.WriteString("{\n")
	sb.WriteByte('[')
	for _, v := range d.Types {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (d *AccountDeleteSecureValueRequest) TypeID() uint32 {
	return AccountDeleteSecureValueRequestTypeID
}

// Encode implements bin.Encoder.
func (d *AccountDeleteSecureValueRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode account.deleteSecureValue#b880bc4b as nil")
	}
	b.PutID(AccountDeleteSecureValueRequestTypeID)
	b.PutVectorHeader(len(d.Types))
	for idx, v := range d.Types {
		if v == nil {
			return fmt.Errorf("unable to encode account.deleteSecureValue#b880bc4b: field types element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.deleteSecureValue#b880bc4b: field types element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetTypes returns value of Types field.
func (d *AccountDeleteSecureValueRequest) GetTypes() (value []SecureValueTypeClass) {
	return d.Types
}

// Decode implements bin.Decoder.
func (d *AccountDeleteSecureValueRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode account.deleteSecureValue#b880bc4b to nil")
	}
	if err := b.ConsumeID(AccountDeleteSecureValueRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.deleteSecureValue#b880bc4b: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.deleteSecureValue#b880bc4b: field types: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeSecureValueType(b)
			if err != nil {
				return fmt.Errorf("unable to decode account.deleteSecureValue#b880bc4b: field types: %w", err)
			}
			d.Types = append(d.Types, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountDeleteSecureValueRequest.
var (
	_ bin.Encoder = &AccountDeleteSecureValueRequest{}
	_ bin.Decoder = &AccountDeleteSecureValueRequest{}
)

// AccountDeleteSecureValue invokes method account.deleteSecureValue#b880bc4b returning error if any.
// Delete stored Telegram Passport¹ documents, for more info see the passport docs »²
//
// Links:
//  1) https://core.telegram.org/passport
//  2) https://core.telegram.org/passport/encryption#encryption
//
// See https://core.telegram.org/method/account.deleteSecureValue for reference.
func (c *Client) AccountDeleteSecureValue(ctx context.Context, types []SecureValueTypeClass) (bool, error) {
	var result BoolBox

	request := &AccountDeleteSecureValueRequest{
		Types: types,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
