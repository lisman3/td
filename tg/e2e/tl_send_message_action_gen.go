// Code generated by gotdgen, DO NOT EDIT.

package e2e

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

// SendMessageTypingAction represents TL type `sendMessageTypingAction#16bf744e`.
// User is typing.
//
// See https://core.telegram.org/constructor/sendMessageTypingAction for reference.
type SendMessageTypingAction struct {
}

// SendMessageTypingActionTypeID is TL type id of SendMessageTypingAction.
const SendMessageTypingActionTypeID = 0x16bf744e

func (s *SendMessageTypingAction) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SendMessageTypingAction) String() string {
	if s == nil {
		return "SendMessageTypingAction(nil)"
	}
	var sb strings.Builder
	sb.WriteString("SendMessageTypingAction")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SendMessageTypingAction) TypeID() uint32 {
	return SendMessageTypingActionTypeID
}

// Encode implements bin.Encoder.
func (s *SendMessageTypingAction) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessageTypingAction#16bf744e as nil")
	}
	b.PutID(SendMessageTypingActionTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendMessageTypingAction) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessageTypingAction#16bf744e to nil")
	}
	if err := b.ConsumeID(SendMessageTypingActionTypeID); err != nil {
		return fmt.Errorf("unable to decode sendMessageTypingAction#16bf744e: %w", err)
	}
	return nil
}

// construct implements constructor of SendMessageActionClass.
func (s SendMessageTypingAction) construct() SendMessageActionClass { return &s }

// Ensuring interfaces in compile-time for SendMessageTypingAction.
var (
	_ bin.Encoder = &SendMessageTypingAction{}
	_ bin.Decoder = &SendMessageTypingAction{}

	_ SendMessageActionClass = &SendMessageTypingAction{}
)

// SendMessageCancelAction represents TL type `sendMessageCancelAction#fd5ec8f5`.
// Invalidate all previous action updates. E.g. when user deletes entered text or aborts a video upload.
//
// See https://core.telegram.org/constructor/sendMessageCancelAction for reference.
type SendMessageCancelAction struct {
}

// SendMessageCancelActionTypeID is TL type id of SendMessageCancelAction.
const SendMessageCancelActionTypeID = 0xfd5ec8f5

func (s *SendMessageCancelAction) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SendMessageCancelAction) String() string {
	if s == nil {
		return "SendMessageCancelAction(nil)"
	}
	var sb strings.Builder
	sb.WriteString("SendMessageCancelAction")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SendMessageCancelAction) TypeID() uint32 {
	return SendMessageCancelActionTypeID
}

// Encode implements bin.Encoder.
func (s *SendMessageCancelAction) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessageCancelAction#fd5ec8f5 as nil")
	}
	b.PutID(SendMessageCancelActionTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendMessageCancelAction) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessageCancelAction#fd5ec8f5 to nil")
	}
	if err := b.ConsumeID(SendMessageCancelActionTypeID); err != nil {
		return fmt.Errorf("unable to decode sendMessageCancelAction#fd5ec8f5: %w", err)
	}
	return nil
}

// construct implements constructor of SendMessageActionClass.
func (s SendMessageCancelAction) construct() SendMessageActionClass { return &s }

// Ensuring interfaces in compile-time for SendMessageCancelAction.
var (
	_ bin.Encoder = &SendMessageCancelAction{}
	_ bin.Decoder = &SendMessageCancelAction{}

	_ SendMessageActionClass = &SendMessageCancelAction{}
)

// SendMessageRecordVideoAction represents TL type `sendMessageRecordVideoAction#a187d66f`.
// User is recording a video.
//
// See https://core.telegram.org/constructor/sendMessageRecordVideoAction for reference.
type SendMessageRecordVideoAction struct {
}

// SendMessageRecordVideoActionTypeID is TL type id of SendMessageRecordVideoAction.
const SendMessageRecordVideoActionTypeID = 0xa187d66f

func (s *SendMessageRecordVideoAction) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SendMessageRecordVideoAction) String() string {
	if s == nil {
		return "SendMessageRecordVideoAction(nil)"
	}
	var sb strings.Builder
	sb.WriteString("SendMessageRecordVideoAction")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SendMessageRecordVideoAction) TypeID() uint32 {
	return SendMessageRecordVideoActionTypeID
}

// Encode implements bin.Encoder.
func (s *SendMessageRecordVideoAction) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessageRecordVideoAction#a187d66f as nil")
	}
	b.PutID(SendMessageRecordVideoActionTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendMessageRecordVideoAction) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessageRecordVideoAction#a187d66f to nil")
	}
	if err := b.ConsumeID(SendMessageRecordVideoActionTypeID); err != nil {
		return fmt.Errorf("unable to decode sendMessageRecordVideoAction#a187d66f: %w", err)
	}
	return nil
}

// construct implements constructor of SendMessageActionClass.
func (s SendMessageRecordVideoAction) construct() SendMessageActionClass { return &s }

// Ensuring interfaces in compile-time for SendMessageRecordVideoAction.
var (
	_ bin.Encoder = &SendMessageRecordVideoAction{}
	_ bin.Decoder = &SendMessageRecordVideoAction{}

	_ SendMessageActionClass = &SendMessageRecordVideoAction{}
)

// SendMessageUploadVideoAction represents TL type `sendMessageUploadVideoAction#92042ff7`.
// User is uploading a video.
//
// See https://core.telegram.org/constructor/sendMessageUploadVideoAction for reference.
type SendMessageUploadVideoAction struct {
}

// SendMessageUploadVideoActionTypeID is TL type id of SendMessageUploadVideoAction.
const SendMessageUploadVideoActionTypeID = 0x92042ff7

func (s *SendMessageUploadVideoAction) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SendMessageUploadVideoAction) String() string {
	if s == nil {
		return "SendMessageUploadVideoAction(nil)"
	}
	var sb strings.Builder
	sb.WriteString("SendMessageUploadVideoAction")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SendMessageUploadVideoAction) TypeID() uint32 {
	return SendMessageUploadVideoActionTypeID
}

// Encode implements bin.Encoder.
func (s *SendMessageUploadVideoAction) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessageUploadVideoAction#92042ff7 as nil")
	}
	b.PutID(SendMessageUploadVideoActionTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendMessageUploadVideoAction) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessageUploadVideoAction#92042ff7 to nil")
	}
	if err := b.ConsumeID(SendMessageUploadVideoActionTypeID); err != nil {
		return fmt.Errorf("unable to decode sendMessageUploadVideoAction#92042ff7: %w", err)
	}
	return nil
}

// construct implements constructor of SendMessageActionClass.
func (s SendMessageUploadVideoAction) construct() SendMessageActionClass { return &s }

// Ensuring interfaces in compile-time for SendMessageUploadVideoAction.
var (
	_ bin.Encoder = &SendMessageUploadVideoAction{}
	_ bin.Decoder = &SendMessageUploadVideoAction{}

	_ SendMessageActionClass = &SendMessageUploadVideoAction{}
)

// SendMessageRecordAudioAction represents TL type `sendMessageRecordAudioAction#d52f73f7`.
// User is recording a voice message.
//
// See https://core.telegram.org/constructor/sendMessageRecordAudioAction for reference.
type SendMessageRecordAudioAction struct {
}

// SendMessageRecordAudioActionTypeID is TL type id of SendMessageRecordAudioAction.
const SendMessageRecordAudioActionTypeID = 0xd52f73f7

func (s *SendMessageRecordAudioAction) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SendMessageRecordAudioAction) String() string {
	if s == nil {
		return "SendMessageRecordAudioAction(nil)"
	}
	var sb strings.Builder
	sb.WriteString("SendMessageRecordAudioAction")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SendMessageRecordAudioAction) TypeID() uint32 {
	return SendMessageRecordAudioActionTypeID
}

// Encode implements bin.Encoder.
func (s *SendMessageRecordAudioAction) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessageRecordAudioAction#d52f73f7 as nil")
	}
	b.PutID(SendMessageRecordAudioActionTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendMessageRecordAudioAction) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessageRecordAudioAction#d52f73f7 to nil")
	}
	if err := b.ConsumeID(SendMessageRecordAudioActionTypeID); err != nil {
		return fmt.Errorf("unable to decode sendMessageRecordAudioAction#d52f73f7: %w", err)
	}
	return nil
}

// construct implements constructor of SendMessageActionClass.
func (s SendMessageRecordAudioAction) construct() SendMessageActionClass { return &s }

// Ensuring interfaces in compile-time for SendMessageRecordAudioAction.
var (
	_ bin.Encoder = &SendMessageRecordAudioAction{}
	_ bin.Decoder = &SendMessageRecordAudioAction{}

	_ SendMessageActionClass = &SendMessageRecordAudioAction{}
)

// SendMessageUploadAudioAction represents TL type `sendMessageUploadAudioAction#e6ac8a6f`.
// User is uploading a voice message.
//
// See https://core.telegram.org/constructor/sendMessageUploadAudioAction for reference.
type SendMessageUploadAudioAction struct {
}

// SendMessageUploadAudioActionTypeID is TL type id of SendMessageUploadAudioAction.
const SendMessageUploadAudioActionTypeID = 0xe6ac8a6f

func (s *SendMessageUploadAudioAction) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SendMessageUploadAudioAction) String() string {
	if s == nil {
		return "SendMessageUploadAudioAction(nil)"
	}
	var sb strings.Builder
	sb.WriteString("SendMessageUploadAudioAction")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SendMessageUploadAudioAction) TypeID() uint32 {
	return SendMessageUploadAudioActionTypeID
}

// Encode implements bin.Encoder.
func (s *SendMessageUploadAudioAction) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessageUploadAudioAction#e6ac8a6f as nil")
	}
	b.PutID(SendMessageUploadAudioActionTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendMessageUploadAudioAction) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessageUploadAudioAction#e6ac8a6f to nil")
	}
	if err := b.ConsumeID(SendMessageUploadAudioActionTypeID); err != nil {
		return fmt.Errorf("unable to decode sendMessageUploadAudioAction#e6ac8a6f: %w", err)
	}
	return nil
}

// construct implements constructor of SendMessageActionClass.
func (s SendMessageUploadAudioAction) construct() SendMessageActionClass { return &s }

// Ensuring interfaces in compile-time for SendMessageUploadAudioAction.
var (
	_ bin.Encoder = &SendMessageUploadAudioAction{}
	_ bin.Decoder = &SendMessageUploadAudioAction{}

	_ SendMessageActionClass = &SendMessageUploadAudioAction{}
)

// SendMessageUploadPhotoAction represents TL type `sendMessageUploadPhotoAction#990a3c1a`.
// User is uploading a photo.
//
// See https://core.telegram.org/constructor/sendMessageUploadPhotoAction for reference.
type SendMessageUploadPhotoAction struct {
}

// SendMessageUploadPhotoActionTypeID is TL type id of SendMessageUploadPhotoAction.
const SendMessageUploadPhotoActionTypeID = 0x990a3c1a

func (s *SendMessageUploadPhotoAction) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SendMessageUploadPhotoAction) String() string {
	if s == nil {
		return "SendMessageUploadPhotoAction(nil)"
	}
	var sb strings.Builder
	sb.WriteString("SendMessageUploadPhotoAction")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SendMessageUploadPhotoAction) TypeID() uint32 {
	return SendMessageUploadPhotoActionTypeID
}

// Encode implements bin.Encoder.
func (s *SendMessageUploadPhotoAction) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessageUploadPhotoAction#990a3c1a as nil")
	}
	b.PutID(SendMessageUploadPhotoActionTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendMessageUploadPhotoAction) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessageUploadPhotoAction#990a3c1a to nil")
	}
	if err := b.ConsumeID(SendMessageUploadPhotoActionTypeID); err != nil {
		return fmt.Errorf("unable to decode sendMessageUploadPhotoAction#990a3c1a: %w", err)
	}
	return nil
}

// construct implements constructor of SendMessageActionClass.
func (s SendMessageUploadPhotoAction) construct() SendMessageActionClass { return &s }

// Ensuring interfaces in compile-time for SendMessageUploadPhotoAction.
var (
	_ bin.Encoder = &SendMessageUploadPhotoAction{}
	_ bin.Decoder = &SendMessageUploadPhotoAction{}

	_ SendMessageActionClass = &SendMessageUploadPhotoAction{}
)

// SendMessageUploadDocumentAction represents TL type `sendMessageUploadDocumentAction#8faee98e`.
// User is uploading a file.
//
// See https://core.telegram.org/constructor/sendMessageUploadDocumentAction for reference.
type SendMessageUploadDocumentAction struct {
}

// SendMessageUploadDocumentActionTypeID is TL type id of SendMessageUploadDocumentAction.
const SendMessageUploadDocumentActionTypeID = 0x8faee98e

func (s *SendMessageUploadDocumentAction) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SendMessageUploadDocumentAction) String() string {
	if s == nil {
		return "SendMessageUploadDocumentAction(nil)"
	}
	var sb strings.Builder
	sb.WriteString("SendMessageUploadDocumentAction")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SendMessageUploadDocumentAction) TypeID() uint32 {
	return SendMessageUploadDocumentActionTypeID
}

// Encode implements bin.Encoder.
func (s *SendMessageUploadDocumentAction) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessageUploadDocumentAction#8faee98e as nil")
	}
	b.PutID(SendMessageUploadDocumentActionTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendMessageUploadDocumentAction) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessageUploadDocumentAction#8faee98e to nil")
	}
	if err := b.ConsumeID(SendMessageUploadDocumentActionTypeID); err != nil {
		return fmt.Errorf("unable to decode sendMessageUploadDocumentAction#8faee98e: %w", err)
	}
	return nil
}

// construct implements constructor of SendMessageActionClass.
func (s SendMessageUploadDocumentAction) construct() SendMessageActionClass { return &s }

// Ensuring interfaces in compile-time for SendMessageUploadDocumentAction.
var (
	_ bin.Encoder = &SendMessageUploadDocumentAction{}
	_ bin.Decoder = &SendMessageUploadDocumentAction{}

	_ SendMessageActionClass = &SendMessageUploadDocumentAction{}
)

// SendMessageGeoLocationAction represents TL type `sendMessageGeoLocationAction#176f8ba1`.
// User is selecting a location to share.
//
// See https://core.telegram.org/constructor/sendMessageGeoLocationAction for reference.
type SendMessageGeoLocationAction struct {
}

// SendMessageGeoLocationActionTypeID is TL type id of SendMessageGeoLocationAction.
const SendMessageGeoLocationActionTypeID = 0x176f8ba1

func (s *SendMessageGeoLocationAction) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SendMessageGeoLocationAction) String() string {
	if s == nil {
		return "SendMessageGeoLocationAction(nil)"
	}
	var sb strings.Builder
	sb.WriteString("SendMessageGeoLocationAction")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SendMessageGeoLocationAction) TypeID() uint32 {
	return SendMessageGeoLocationActionTypeID
}

// Encode implements bin.Encoder.
func (s *SendMessageGeoLocationAction) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessageGeoLocationAction#176f8ba1 as nil")
	}
	b.PutID(SendMessageGeoLocationActionTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendMessageGeoLocationAction) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessageGeoLocationAction#176f8ba1 to nil")
	}
	if err := b.ConsumeID(SendMessageGeoLocationActionTypeID); err != nil {
		return fmt.Errorf("unable to decode sendMessageGeoLocationAction#176f8ba1: %w", err)
	}
	return nil
}

// construct implements constructor of SendMessageActionClass.
func (s SendMessageGeoLocationAction) construct() SendMessageActionClass { return &s }

// Ensuring interfaces in compile-time for SendMessageGeoLocationAction.
var (
	_ bin.Encoder = &SendMessageGeoLocationAction{}
	_ bin.Decoder = &SendMessageGeoLocationAction{}

	_ SendMessageActionClass = &SendMessageGeoLocationAction{}
)

// SendMessageChooseContactAction represents TL type `sendMessageChooseContactAction#628cbc6f`.
// User is selecting a contact to share.
//
// See https://core.telegram.org/constructor/sendMessageChooseContactAction for reference.
type SendMessageChooseContactAction struct {
}

// SendMessageChooseContactActionTypeID is TL type id of SendMessageChooseContactAction.
const SendMessageChooseContactActionTypeID = 0x628cbc6f

func (s *SendMessageChooseContactAction) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SendMessageChooseContactAction) String() string {
	if s == nil {
		return "SendMessageChooseContactAction(nil)"
	}
	var sb strings.Builder
	sb.WriteString("SendMessageChooseContactAction")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SendMessageChooseContactAction) TypeID() uint32 {
	return SendMessageChooseContactActionTypeID
}

// Encode implements bin.Encoder.
func (s *SendMessageChooseContactAction) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessageChooseContactAction#628cbc6f as nil")
	}
	b.PutID(SendMessageChooseContactActionTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendMessageChooseContactAction) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessageChooseContactAction#628cbc6f to nil")
	}
	if err := b.ConsumeID(SendMessageChooseContactActionTypeID); err != nil {
		return fmt.Errorf("unable to decode sendMessageChooseContactAction#628cbc6f: %w", err)
	}
	return nil
}

// construct implements constructor of SendMessageActionClass.
func (s SendMessageChooseContactAction) construct() SendMessageActionClass { return &s }

// Ensuring interfaces in compile-time for SendMessageChooseContactAction.
var (
	_ bin.Encoder = &SendMessageChooseContactAction{}
	_ bin.Decoder = &SendMessageChooseContactAction{}

	_ SendMessageActionClass = &SendMessageChooseContactAction{}
)

// SendMessageRecordRoundAction represents TL type `sendMessageRecordRoundAction#88f27fbc`.
// User is recording a round video to share
//
// See https://core.telegram.org/constructor/sendMessageRecordRoundAction for reference.
type SendMessageRecordRoundAction struct {
}

// SendMessageRecordRoundActionTypeID is TL type id of SendMessageRecordRoundAction.
const SendMessageRecordRoundActionTypeID = 0x88f27fbc

func (s *SendMessageRecordRoundAction) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SendMessageRecordRoundAction) String() string {
	if s == nil {
		return "SendMessageRecordRoundAction(nil)"
	}
	var sb strings.Builder
	sb.WriteString("SendMessageRecordRoundAction")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SendMessageRecordRoundAction) TypeID() uint32 {
	return SendMessageRecordRoundActionTypeID
}

// Encode implements bin.Encoder.
func (s *SendMessageRecordRoundAction) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessageRecordRoundAction#88f27fbc as nil")
	}
	b.PutID(SendMessageRecordRoundActionTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendMessageRecordRoundAction) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessageRecordRoundAction#88f27fbc to nil")
	}
	if err := b.ConsumeID(SendMessageRecordRoundActionTypeID); err != nil {
		return fmt.Errorf("unable to decode sendMessageRecordRoundAction#88f27fbc: %w", err)
	}
	return nil
}

// construct implements constructor of SendMessageActionClass.
func (s SendMessageRecordRoundAction) construct() SendMessageActionClass { return &s }

// Ensuring interfaces in compile-time for SendMessageRecordRoundAction.
var (
	_ bin.Encoder = &SendMessageRecordRoundAction{}
	_ bin.Decoder = &SendMessageRecordRoundAction{}

	_ SendMessageActionClass = &SendMessageRecordRoundAction{}
)

// SendMessageUploadRoundAction represents TL type `sendMessageUploadRoundAction#bb718624`.
// User is uploading a round video
//
// See https://core.telegram.org/constructor/sendMessageUploadRoundAction for reference.
type SendMessageUploadRoundAction struct {
}

// SendMessageUploadRoundActionTypeID is TL type id of SendMessageUploadRoundAction.
const SendMessageUploadRoundActionTypeID = 0xbb718624

func (s *SendMessageUploadRoundAction) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SendMessageUploadRoundAction) String() string {
	if s == nil {
		return "SendMessageUploadRoundAction(nil)"
	}
	var sb strings.Builder
	sb.WriteString("SendMessageUploadRoundAction")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SendMessageUploadRoundAction) TypeID() uint32 {
	return SendMessageUploadRoundActionTypeID
}

// Encode implements bin.Encoder.
func (s *SendMessageUploadRoundAction) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessageUploadRoundAction#bb718624 as nil")
	}
	b.PutID(SendMessageUploadRoundActionTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendMessageUploadRoundAction) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessageUploadRoundAction#bb718624 to nil")
	}
	if err := b.ConsumeID(SendMessageUploadRoundActionTypeID); err != nil {
		return fmt.Errorf("unable to decode sendMessageUploadRoundAction#bb718624: %w", err)
	}
	return nil
}

// construct implements constructor of SendMessageActionClass.
func (s SendMessageUploadRoundAction) construct() SendMessageActionClass { return &s }

// Ensuring interfaces in compile-time for SendMessageUploadRoundAction.
var (
	_ bin.Encoder = &SendMessageUploadRoundAction{}
	_ bin.Decoder = &SendMessageUploadRoundAction{}

	_ SendMessageActionClass = &SendMessageUploadRoundAction{}
)

// SendMessageActionClass represents SendMessageAction generic type.
//
// See https://core.telegram.org/type/SendMessageAction for reference.
//
// Example:
//  g, err := DecodeSendMessageAction(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *SendMessageTypingAction: // sendMessageTypingAction#16bf744e
//  case *SendMessageCancelAction: // sendMessageCancelAction#fd5ec8f5
//  case *SendMessageRecordVideoAction: // sendMessageRecordVideoAction#a187d66f
//  case *SendMessageUploadVideoAction: // sendMessageUploadVideoAction#92042ff7
//  case *SendMessageRecordAudioAction: // sendMessageRecordAudioAction#d52f73f7
//  case *SendMessageUploadAudioAction: // sendMessageUploadAudioAction#e6ac8a6f
//  case *SendMessageUploadPhotoAction: // sendMessageUploadPhotoAction#990a3c1a
//  case *SendMessageUploadDocumentAction: // sendMessageUploadDocumentAction#8faee98e
//  case *SendMessageGeoLocationAction: // sendMessageGeoLocationAction#176f8ba1
//  case *SendMessageChooseContactAction: // sendMessageChooseContactAction#628cbc6f
//  case *SendMessageRecordRoundAction: // sendMessageRecordRoundAction#88f27fbc
//  case *SendMessageUploadRoundAction: // sendMessageUploadRoundAction#bb718624
//  default: panic(v)
//  }
type SendMessageActionClass interface {
	bin.Encoder
	bin.Decoder
	construct() SendMessageActionClass

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeSendMessageAction implements binary de-serialization for SendMessageActionClass.
func DecodeSendMessageAction(buf *bin.Buffer) (SendMessageActionClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case SendMessageTypingActionTypeID:
		// Decoding sendMessageTypingAction#16bf744e.
		v := SendMessageTypingAction{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SendMessageActionClass: %w", err)
		}
		return &v, nil
	case SendMessageCancelActionTypeID:
		// Decoding sendMessageCancelAction#fd5ec8f5.
		v := SendMessageCancelAction{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SendMessageActionClass: %w", err)
		}
		return &v, nil
	case SendMessageRecordVideoActionTypeID:
		// Decoding sendMessageRecordVideoAction#a187d66f.
		v := SendMessageRecordVideoAction{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SendMessageActionClass: %w", err)
		}
		return &v, nil
	case SendMessageUploadVideoActionTypeID:
		// Decoding sendMessageUploadVideoAction#92042ff7.
		v := SendMessageUploadVideoAction{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SendMessageActionClass: %w", err)
		}
		return &v, nil
	case SendMessageRecordAudioActionTypeID:
		// Decoding sendMessageRecordAudioAction#d52f73f7.
		v := SendMessageRecordAudioAction{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SendMessageActionClass: %w", err)
		}
		return &v, nil
	case SendMessageUploadAudioActionTypeID:
		// Decoding sendMessageUploadAudioAction#e6ac8a6f.
		v := SendMessageUploadAudioAction{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SendMessageActionClass: %w", err)
		}
		return &v, nil
	case SendMessageUploadPhotoActionTypeID:
		// Decoding sendMessageUploadPhotoAction#990a3c1a.
		v := SendMessageUploadPhotoAction{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SendMessageActionClass: %w", err)
		}
		return &v, nil
	case SendMessageUploadDocumentActionTypeID:
		// Decoding sendMessageUploadDocumentAction#8faee98e.
		v := SendMessageUploadDocumentAction{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SendMessageActionClass: %w", err)
		}
		return &v, nil
	case SendMessageGeoLocationActionTypeID:
		// Decoding sendMessageGeoLocationAction#176f8ba1.
		v := SendMessageGeoLocationAction{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SendMessageActionClass: %w", err)
		}
		return &v, nil
	case SendMessageChooseContactActionTypeID:
		// Decoding sendMessageChooseContactAction#628cbc6f.
		v := SendMessageChooseContactAction{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SendMessageActionClass: %w", err)
		}
		return &v, nil
	case SendMessageRecordRoundActionTypeID:
		// Decoding sendMessageRecordRoundAction#88f27fbc.
		v := SendMessageRecordRoundAction{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SendMessageActionClass: %w", err)
		}
		return &v, nil
	case SendMessageUploadRoundActionTypeID:
		// Decoding sendMessageUploadRoundAction#bb718624.
		v := SendMessageUploadRoundAction{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SendMessageActionClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode SendMessageActionClass: %w", bin.NewUnexpectedID(id))
	}
}

// SendMessageAction boxes the SendMessageActionClass providing a helper.
type SendMessageActionBox struct {
	SendMessageAction SendMessageActionClass
}

// Decode implements bin.Decoder for SendMessageActionBox.
func (b *SendMessageActionBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode SendMessageActionBox to nil")
	}
	v, err := DecodeSendMessageAction(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.SendMessageAction = v
	return nil
}

// Encode implements bin.Encode for SendMessageActionBox.
func (b *SendMessageActionBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.SendMessageAction == nil {
		return fmt.Errorf("unable to encode SendMessageActionClass as nil")
	}
	return b.SendMessageAction.Encode(buf)
}
