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

// AccountUpdateUsernameRequest represents TL type `account.updateUsername#3e0bdd7c`.
// Changes username for the current user.
//
// See https://core.telegram.org/method/account.updateUsername for reference.
type AccountUpdateUsernameRequest struct {
	// username or empty string if username is to be removedAccepted characters: a-z (case-insensitive), 0-9 and underscores.Length: 5-32 characters.
	Username string
}

// AccountUpdateUsernameRequestTypeID is TL type id of AccountUpdateUsernameRequest.
const AccountUpdateUsernameRequestTypeID = 0x3e0bdd7c

func (u *AccountUpdateUsernameRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Username == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *AccountUpdateUsernameRequest) String() string {
	if u == nil {
		return "AccountUpdateUsernameRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AccountUpdateUsernameRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tUsername: ")
	sb.WriteString(fmt.Sprint(u.Username))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (u *AccountUpdateUsernameRequest) TypeID() uint32 {
	return AccountUpdateUsernameRequestTypeID
}

// Encode implements bin.Encoder.
func (u *AccountUpdateUsernameRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateUsername#3e0bdd7c as nil")
	}
	b.PutID(AccountUpdateUsernameRequestTypeID)
	b.PutString(u.Username)
	return nil
}

// GetUsername returns value of Username field.
func (u *AccountUpdateUsernameRequest) GetUsername() (value string) {
	return u.Username
}

// Decode implements bin.Decoder.
func (u *AccountUpdateUsernameRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateUsername#3e0bdd7c to nil")
	}
	if err := b.ConsumeID(AccountUpdateUsernameRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.updateUsername#3e0bdd7c: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.updateUsername#3e0bdd7c: field username: %w", err)
		}
		u.Username = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountUpdateUsernameRequest.
var (
	_ bin.Encoder = &AccountUpdateUsernameRequest{}
	_ bin.Decoder = &AccountUpdateUsernameRequest{}
)

// AccountUpdateUsername invokes method account.updateUsername#3e0bdd7c returning error if any.
// Changes username for the current user.
//
// Possible errors:
//  401 AUTH_KEY_PERM_EMPTY: The temporary auth key must be binded to the permanent auth key to use these methods.
//  400 USERNAME_INVALID: Unacceptable username
//  400 USERNAME_NOT_MODIFIED: Username is not different from the current username
//  400 USERNAME_OCCUPIED: Username is taken
//
// See https://core.telegram.org/method/account.updateUsername for reference.
func (c *Client) AccountUpdateUsername(ctx context.Context, username string) (UserClass, error) {
	var result UserBox

	request := &AccountUpdateUsernameRequest{
		Username: username,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.User, nil
}
