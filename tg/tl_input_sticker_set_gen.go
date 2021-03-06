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

// InputStickerSetEmpty represents TL type `inputStickerSetEmpty#ffb62b95`.
// Empty constructor
//
// See https://core.telegram.org/constructor/inputStickerSetEmpty for reference.
type InputStickerSetEmpty struct {
}

// InputStickerSetEmptyTypeID is TL type id of InputStickerSetEmpty.
const InputStickerSetEmptyTypeID = 0xffb62b95

func (i *InputStickerSetEmpty) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputStickerSetEmpty) String() string {
	if i == nil {
		return "InputStickerSetEmpty(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputStickerSetEmpty")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputStickerSetEmpty) TypeID() uint32 {
	return InputStickerSetEmptyTypeID
}

// Encode implements bin.Encoder.
func (i *InputStickerSetEmpty) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStickerSetEmpty#ffb62b95 as nil")
	}
	b.PutID(InputStickerSetEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputStickerSetEmpty) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStickerSetEmpty#ffb62b95 to nil")
	}
	if err := b.ConsumeID(InputStickerSetEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputStickerSetEmpty#ffb62b95: %w", err)
	}
	return nil
}

// construct implements constructor of InputStickerSetClass.
func (i InputStickerSetEmpty) construct() InputStickerSetClass { return &i }

// Ensuring interfaces in compile-time for InputStickerSetEmpty.
var (
	_ bin.Encoder = &InputStickerSetEmpty{}
	_ bin.Decoder = &InputStickerSetEmpty{}

	_ InputStickerSetClass = &InputStickerSetEmpty{}
)

// InputStickerSetID represents TL type `inputStickerSetID#9de7a269`.
// Stickerset by ID
//
// See https://core.telegram.org/constructor/inputStickerSetID for reference.
type InputStickerSetID struct {
	// ID
	ID int64
	// Access hash
	AccessHash int64
}

// InputStickerSetIDTypeID is TL type id of InputStickerSetID.
const InputStickerSetIDTypeID = 0x9de7a269

func (i *InputStickerSetID) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == 0) {
		return false
	}
	if !(i.AccessHash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputStickerSetID) String() string {
	if i == nil {
		return "InputStickerSetID(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputStickerSetID")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(i.ID))
	sb.WriteString(",\n")
	sb.WriteString("\tAccessHash: ")
	sb.WriteString(fmt.Sprint(i.AccessHash))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputStickerSetID) TypeID() uint32 {
	return InputStickerSetIDTypeID
}

// Encode implements bin.Encoder.
func (i *InputStickerSetID) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStickerSetID#9de7a269 as nil")
	}
	b.PutID(InputStickerSetIDTypeID)
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	return nil
}

// GetID returns value of ID field.
func (i *InputStickerSetID) GetID() (value int64) {
	return i.ID
}

// GetAccessHash returns value of AccessHash field.
func (i *InputStickerSetID) GetAccessHash() (value int64) {
	return i.AccessHash
}

// Decode implements bin.Decoder.
func (i *InputStickerSetID) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStickerSetID#9de7a269 to nil")
	}
	if err := b.ConsumeID(InputStickerSetIDTypeID); err != nil {
		return fmt.Errorf("unable to decode inputStickerSetID#9de7a269: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputStickerSetID#9de7a269: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputStickerSetID#9de7a269: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// construct implements constructor of InputStickerSetClass.
func (i InputStickerSetID) construct() InputStickerSetClass { return &i }

// Ensuring interfaces in compile-time for InputStickerSetID.
var (
	_ bin.Encoder = &InputStickerSetID{}
	_ bin.Decoder = &InputStickerSetID{}

	_ InputStickerSetClass = &InputStickerSetID{}
)

// InputStickerSetShortName represents TL type `inputStickerSetShortName#861cc8a0`.
// Stickerset by short name, from tg://addstickers?set=short_name
//
// See https://core.telegram.org/constructor/inputStickerSetShortName for reference.
type InputStickerSetShortName struct {
	// From tg://addstickers?set=short_name
	ShortName string
}

// InputStickerSetShortNameTypeID is TL type id of InputStickerSetShortName.
const InputStickerSetShortNameTypeID = 0x861cc8a0

func (i *InputStickerSetShortName) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ShortName == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputStickerSetShortName) String() string {
	if i == nil {
		return "InputStickerSetShortName(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputStickerSetShortName")
	sb.WriteString("{\n")
	sb.WriteString("\tShortName: ")
	sb.WriteString(fmt.Sprint(i.ShortName))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputStickerSetShortName) TypeID() uint32 {
	return InputStickerSetShortNameTypeID
}

// Encode implements bin.Encoder.
func (i *InputStickerSetShortName) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStickerSetShortName#861cc8a0 as nil")
	}
	b.PutID(InputStickerSetShortNameTypeID)
	b.PutString(i.ShortName)
	return nil
}

// GetShortName returns value of ShortName field.
func (i *InputStickerSetShortName) GetShortName() (value string) {
	return i.ShortName
}

// Decode implements bin.Decoder.
func (i *InputStickerSetShortName) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStickerSetShortName#861cc8a0 to nil")
	}
	if err := b.ConsumeID(InputStickerSetShortNameTypeID); err != nil {
		return fmt.Errorf("unable to decode inputStickerSetShortName#861cc8a0: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputStickerSetShortName#861cc8a0: field short_name: %w", err)
		}
		i.ShortName = value
	}
	return nil
}

// construct implements constructor of InputStickerSetClass.
func (i InputStickerSetShortName) construct() InputStickerSetClass { return &i }

// Ensuring interfaces in compile-time for InputStickerSetShortName.
var (
	_ bin.Encoder = &InputStickerSetShortName{}
	_ bin.Decoder = &InputStickerSetShortName{}

	_ InputStickerSetClass = &InputStickerSetShortName{}
)

// InputStickerSetAnimatedEmoji represents TL type `inputStickerSetAnimatedEmoji#28703c8`.
// Animated emojis stickerset
//
// See https://core.telegram.org/constructor/inputStickerSetAnimatedEmoji for reference.
type InputStickerSetAnimatedEmoji struct {
}

// InputStickerSetAnimatedEmojiTypeID is TL type id of InputStickerSetAnimatedEmoji.
const InputStickerSetAnimatedEmojiTypeID = 0x28703c8

func (i *InputStickerSetAnimatedEmoji) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputStickerSetAnimatedEmoji) String() string {
	if i == nil {
		return "InputStickerSetAnimatedEmoji(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputStickerSetAnimatedEmoji")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputStickerSetAnimatedEmoji) TypeID() uint32 {
	return InputStickerSetAnimatedEmojiTypeID
}

// Encode implements bin.Encoder.
func (i *InputStickerSetAnimatedEmoji) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStickerSetAnimatedEmoji#28703c8 as nil")
	}
	b.PutID(InputStickerSetAnimatedEmojiTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputStickerSetAnimatedEmoji) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStickerSetAnimatedEmoji#28703c8 to nil")
	}
	if err := b.ConsumeID(InputStickerSetAnimatedEmojiTypeID); err != nil {
		return fmt.Errorf("unable to decode inputStickerSetAnimatedEmoji#28703c8: %w", err)
	}
	return nil
}

// construct implements constructor of InputStickerSetClass.
func (i InputStickerSetAnimatedEmoji) construct() InputStickerSetClass { return &i }

// Ensuring interfaces in compile-time for InputStickerSetAnimatedEmoji.
var (
	_ bin.Encoder = &InputStickerSetAnimatedEmoji{}
	_ bin.Decoder = &InputStickerSetAnimatedEmoji{}

	_ InputStickerSetClass = &InputStickerSetAnimatedEmoji{}
)

// InputStickerSetDice represents TL type `inputStickerSetDice#e67f520e`.
// Used for fetching animated dice stickers¹
//
// Links:
//  1) https://core.telegram.org/api/dice
//
// See https://core.telegram.org/constructor/inputStickerSetDice for reference.
type InputStickerSetDice struct {
	// The emoji, for now ,  and  are supported
	Emoticon string
}

// InputStickerSetDiceTypeID is TL type id of InputStickerSetDice.
const InputStickerSetDiceTypeID = 0xe67f520e

func (i *InputStickerSetDice) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Emoticon == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputStickerSetDice) String() string {
	if i == nil {
		return "InputStickerSetDice(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputStickerSetDice")
	sb.WriteString("{\n")
	sb.WriteString("\tEmoticon: ")
	sb.WriteString(fmt.Sprint(i.Emoticon))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputStickerSetDice) TypeID() uint32 {
	return InputStickerSetDiceTypeID
}

// Encode implements bin.Encoder.
func (i *InputStickerSetDice) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStickerSetDice#e67f520e as nil")
	}
	b.PutID(InputStickerSetDiceTypeID)
	b.PutString(i.Emoticon)
	return nil
}

// GetEmoticon returns value of Emoticon field.
func (i *InputStickerSetDice) GetEmoticon() (value string) {
	return i.Emoticon
}

// Decode implements bin.Decoder.
func (i *InputStickerSetDice) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStickerSetDice#e67f520e to nil")
	}
	if err := b.ConsumeID(InputStickerSetDiceTypeID); err != nil {
		return fmt.Errorf("unable to decode inputStickerSetDice#e67f520e: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputStickerSetDice#e67f520e: field emoticon: %w", err)
		}
		i.Emoticon = value
	}
	return nil
}

// construct implements constructor of InputStickerSetClass.
func (i InputStickerSetDice) construct() InputStickerSetClass { return &i }

// Ensuring interfaces in compile-time for InputStickerSetDice.
var (
	_ bin.Encoder = &InputStickerSetDice{}
	_ bin.Decoder = &InputStickerSetDice{}

	_ InputStickerSetClass = &InputStickerSetDice{}
)

// InputStickerSetClass represents InputStickerSet generic type.
//
// See https://core.telegram.org/type/InputStickerSet for reference.
//
// Example:
//  g, err := DecodeInputStickerSet(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InputStickerSetEmpty: // inputStickerSetEmpty#ffb62b95
//  case *InputStickerSetID: // inputStickerSetID#9de7a269
//  case *InputStickerSetShortName: // inputStickerSetShortName#861cc8a0
//  case *InputStickerSetAnimatedEmoji: // inputStickerSetAnimatedEmoji#28703c8
//  case *InputStickerSetDice: // inputStickerSetDice#e67f520e
//  default: panic(v)
//  }
type InputStickerSetClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputStickerSetClass

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeInputStickerSet implements binary de-serialization for InputStickerSetClass.
func DecodeInputStickerSet(buf *bin.Buffer) (InputStickerSetClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputStickerSetEmptyTypeID:
		// Decoding inputStickerSetEmpty#ffb62b95.
		v := InputStickerSetEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStickerSetClass: %w", err)
		}
		return &v, nil
	case InputStickerSetIDTypeID:
		// Decoding inputStickerSetID#9de7a269.
		v := InputStickerSetID{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStickerSetClass: %w", err)
		}
		return &v, nil
	case InputStickerSetShortNameTypeID:
		// Decoding inputStickerSetShortName#861cc8a0.
		v := InputStickerSetShortName{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStickerSetClass: %w", err)
		}
		return &v, nil
	case InputStickerSetAnimatedEmojiTypeID:
		// Decoding inputStickerSetAnimatedEmoji#28703c8.
		v := InputStickerSetAnimatedEmoji{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStickerSetClass: %w", err)
		}
		return &v, nil
	case InputStickerSetDiceTypeID:
		// Decoding inputStickerSetDice#e67f520e.
		v := InputStickerSetDice{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStickerSetClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputStickerSetClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputStickerSet boxes the InputStickerSetClass providing a helper.
type InputStickerSetBox struct {
	InputStickerSet InputStickerSetClass
}

// Decode implements bin.Decoder for InputStickerSetBox.
func (b *InputStickerSetBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputStickerSetBox to nil")
	}
	v, err := DecodeInputStickerSet(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputStickerSet = v
	return nil
}

// Encode implements bin.Encode for InputStickerSetBox.
func (b *InputStickerSetBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputStickerSet == nil {
		return fmt.Errorf("unable to encode InputStickerSetClass as nil")
	}
	return b.InputStickerSet.Encode(buf)
}
