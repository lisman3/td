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

// AuthDropTempAuthKeysRequest represents TL type `auth.dropTempAuthKeys#8e48a188`.
type AuthDropTempAuthKeysRequest struct {
	// ExceptAuthKeys field of AuthDropTempAuthKeysRequest.
	ExceptAuthKeys []int64
}

// AuthDropTempAuthKeysRequestTypeID is TL type id of AuthDropTempAuthKeysRequest.
const AuthDropTempAuthKeysRequestTypeID = 0x8e48a188

// Encode implements bin.Encoder.
func (d *AuthDropTempAuthKeysRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode auth.dropTempAuthKeys#8e48a188 as nil")
	}
	b.PutID(AuthDropTempAuthKeysRequestTypeID)
	b.PutVectorHeader(len(d.ExceptAuthKeys))
	for _, v := range d.ExceptAuthKeys {
		b.PutLong(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *AuthDropTempAuthKeysRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode auth.dropTempAuthKeys#8e48a188 to nil")
	}
	if err := b.ConsumeID(AuthDropTempAuthKeysRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.dropTempAuthKeys#8e48a188: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode auth.dropTempAuthKeys#8e48a188: field except_auth_keys: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode auth.dropTempAuthKeys#8e48a188: field except_auth_keys: %w", err)
			}
			d.ExceptAuthKeys = append(d.ExceptAuthKeys, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AuthDropTempAuthKeysRequest.
var (
	_ bin.Encoder = &AuthDropTempAuthKeysRequest{}
	_ bin.Decoder = &AuthDropTempAuthKeysRequest{}
)

// AuthDropTempAuthKeys invokes method auth.dropTempAuthKeys#8e48a188 returning error if any.
func (c *Client) AuthDropTempAuthKeys(ctx context.Context, request *AuthDropTempAuthKeysRequest) (BoolClass, error) {
	var result BoolBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Bool, nil
}