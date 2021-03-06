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

// AuthPasswordRecovery represents TL type `auth.passwordRecovery#137948a5`.
// Recovery info of a 2FA password¹, only for accounts with a recovery email configured².
//
// Links:
//  1) https://core.telegram.org/api/srp
//  2) https://core.telegram.org/api/srp#email-verification
//
// See https://core.telegram.org/constructor/auth.passwordRecovery for reference.
type AuthPasswordRecovery struct {
	// The email to which the recovery code was sent must match this pattern¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/pattern
	EmailPattern string
}

// AuthPasswordRecoveryTypeID is TL type id of AuthPasswordRecovery.
const AuthPasswordRecoveryTypeID = 0x137948a5

func (p *AuthPasswordRecovery) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.EmailPattern == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *AuthPasswordRecovery) String() string {
	if p == nil {
		return "AuthPasswordRecovery(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AuthPasswordRecovery")
	sb.WriteString("{\n")
	sb.WriteString("\tEmailPattern: ")
	sb.WriteString(fmt.Sprint(p.EmailPattern))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *AuthPasswordRecovery) TypeID() uint32 {
	return AuthPasswordRecoveryTypeID
}

// Encode implements bin.Encoder.
func (p *AuthPasswordRecovery) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode auth.passwordRecovery#137948a5 as nil")
	}
	b.PutID(AuthPasswordRecoveryTypeID)
	b.PutString(p.EmailPattern)
	return nil
}

// GetEmailPattern returns value of EmailPattern field.
func (p *AuthPasswordRecovery) GetEmailPattern() (value string) {
	return p.EmailPattern
}

// Decode implements bin.Decoder.
func (p *AuthPasswordRecovery) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode auth.passwordRecovery#137948a5 to nil")
	}
	if err := b.ConsumeID(AuthPasswordRecoveryTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.passwordRecovery#137948a5: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.passwordRecovery#137948a5: field email_pattern: %w", err)
		}
		p.EmailPattern = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AuthPasswordRecovery.
var (
	_ bin.Encoder = &AuthPasswordRecovery{}
	_ bin.Decoder = &AuthPasswordRecovery{}
)
