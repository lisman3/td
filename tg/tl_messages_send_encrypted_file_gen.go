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

// MessagesSendEncryptedFileRequest represents TL type `messages.sendEncryptedFile#5559481d`.
// Sends a message with a file attachment to a secret chat
//
// See https://core.telegram.org/method/messages.sendEncryptedFile for reference.
type MessagesSendEncryptedFileRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to send the file without triggering a notification
	Silent bool
	// Secret chat ID
	Peer InputEncryptedChat
	// Unique client message ID necessary to prevent message resending
	RandomID int64
	// TL-serialization of DecryptedMessage¹ type, encrypted with a key generated during chat initialization
	//
	// Links:
	//  1) https://core.telegram.org/type/DecryptedMessage
	Data []byte
	// File attachment for the secret chat
	File InputEncryptedFileClass
}

// MessagesSendEncryptedFileRequestTypeID is TL type id of MessagesSendEncryptedFileRequest.
const MessagesSendEncryptedFileRequestTypeID = 0x5559481d

func (s *MessagesSendEncryptedFileRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Silent == false) {
		return false
	}
	if !(s.Peer.Zero()) {
		return false
	}
	if !(s.RandomID == 0) {
		return false
	}
	if !(s.Data == nil) {
		return false
	}
	if !(s.File == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSendEncryptedFileRequest) String() string {
	if s == nil {
		return "MessagesSendEncryptedFileRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesSendEncryptedFileRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(s.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(fmt.Sprint(s.Peer))
	sb.WriteString(",\n")
	sb.WriteString("\tRandomID: ")
	sb.WriteString(fmt.Sprint(s.RandomID))
	sb.WriteString(",\n")
	sb.WriteString("\tData: ")
	sb.WriteString(fmt.Sprint(s.Data))
	sb.WriteString(",\n")
	sb.WriteString("\tFile: ")
	sb.WriteString(fmt.Sprint(s.File))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *MessagesSendEncryptedFileRequest) TypeID() uint32 {
	return MessagesSendEncryptedFileRequestTypeID
}

// Encode implements bin.Encoder.
func (s *MessagesSendEncryptedFileRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendEncryptedFile#5559481d as nil")
	}
	b.PutID(MessagesSendEncryptedFileRequestTypeID)
	if !(s.Silent == false) {
		s.Flags.Set(0)
	}
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendEncryptedFile#5559481d: field flags: %w", err)
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendEncryptedFile#5559481d: field peer: %w", err)
	}
	b.PutLong(s.RandomID)
	b.PutBytes(s.Data)
	if s.File == nil {
		return fmt.Errorf("unable to encode messages.sendEncryptedFile#5559481d: field file is nil")
	}
	if err := s.File.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendEncryptedFile#5559481d: field file: %w", err)
	}
	return nil
}

// SetSilent sets value of Silent conditional field.
func (s *MessagesSendEncryptedFileRequest) SetSilent(value bool) {
	if value {
		s.Flags.Set(0)
		s.Silent = true
	} else {
		s.Flags.Unset(0)
		s.Silent = false
	}
}

// GetSilent returns value of Silent conditional field.
func (s *MessagesSendEncryptedFileRequest) GetSilent() (value bool) {
	return s.Flags.Has(0)
}

// GetPeer returns value of Peer field.
func (s *MessagesSendEncryptedFileRequest) GetPeer() (value InputEncryptedChat) {
	return s.Peer
}

// GetRandomID returns value of RandomID field.
func (s *MessagesSendEncryptedFileRequest) GetRandomID() (value int64) {
	return s.RandomID
}

// GetData returns value of Data field.
func (s *MessagesSendEncryptedFileRequest) GetData() (value []byte) {
	return s.Data
}

// GetFile returns value of File field.
func (s *MessagesSendEncryptedFileRequest) GetFile() (value InputEncryptedFileClass) {
	return s.File
}

// Decode implements bin.Decoder.
func (s *MessagesSendEncryptedFileRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendEncryptedFile#5559481d to nil")
	}
	if err := b.ConsumeID(MessagesSendEncryptedFileRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.sendEncryptedFile#5559481d: %w", err)
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.sendEncryptedFile#5559481d: field flags: %w", err)
		}
	}
	s.Silent = s.Flags.Has(0)
	{
		if err := s.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.sendEncryptedFile#5559481d: field peer: %w", err)
		}
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendEncryptedFile#5559481d: field random_id: %w", err)
		}
		s.RandomID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendEncryptedFile#5559481d: field data: %w", err)
		}
		s.Data = value
	}
	{
		value, err := DecodeInputEncryptedFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendEncryptedFile#5559481d: field file: %w", err)
		}
		s.File = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSendEncryptedFileRequest.
var (
	_ bin.Encoder = &MessagesSendEncryptedFileRequest{}
	_ bin.Decoder = &MessagesSendEncryptedFileRequest{}
)

// MessagesSendEncryptedFile invokes method messages.sendEncryptedFile#5559481d returning error if any.
// Sends a message with a file attachment to a secret chat
//
// Possible errors:
//  400 DATA_TOO_LONG: Data too long
//  400 ENCRYPTION_DECLINED: The secret chat was declined
//  400 MD5_CHECKSUM_INVALID: The MD5 checksums do not match
//  400 MSG_WAIT_FAILED: A waiting call returned an error
//
// See https://core.telegram.org/method/messages.sendEncryptedFile for reference.
func (c *Client) MessagesSendEncryptedFile(ctx context.Context, request *MessagesSendEncryptedFileRequest) (MessagesSentEncryptedMessageClass, error) {
	var result MessagesSentEncryptedMessageBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.SentEncryptedMessage, nil
}
