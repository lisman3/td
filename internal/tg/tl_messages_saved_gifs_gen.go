// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/ernado/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// MessagesSavedGifsNotModified represents TL type `messages.savedGifsNotModified#e8025ca2`.
type MessagesSavedGifsNotModified struct {
}

// MessagesSavedGifsNotModifiedTypeID is TL type id of MessagesSavedGifsNotModified.
const MessagesSavedGifsNotModifiedTypeID = 0xe8025ca2

// Encode implements bin.Encoder.
func (s *MessagesSavedGifsNotModified) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.savedGifsNotModified#e8025ca2 as nil")
	}
	b.PutID(MessagesSavedGifsNotModifiedTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSavedGifsNotModified) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.savedGifsNotModified#e8025ca2 to nil")
	}
	if err := b.ConsumeID(MessagesSavedGifsNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.savedGifsNotModified#e8025ca2: %w", err)
	}
	return nil
}

// construct implements constructor of MessagesSavedGifsClass.
func (s MessagesSavedGifsNotModified) construct() MessagesSavedGifsClass { return &s }

// Ensuring interfaces in compile-time for MessagesSavedGifsNotModified.
var (
	_ bin.Encoder = &MessagesSavedGifsNotModified{}
	_ bin.Decoder = &MessagesSavedGifsNotModified{}

	_ MessagesSavedGifsClass = &MessagesSavedGifsNotModified{}
)

// MessagesSavedGifs represents TL type `messages.savedGifs#2e0709a5`.
type MessagesSavedGifs struct {
	// Hash field of MessagesSavedGifs.
	Hash int
	// Gifs field of MessagesSavedGifs.
	Gifs []DocumentClass
}

// MessagesSavedGifsTypeID is TL type id of MessagesSavedGifs.
const MessagesSavedGifsTypeID = 0x2e0709a5

// Encode implements bin.Encoder.
func (s *MessagesSavedGifs) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.savedGifs#2e0709a5 as nil")
	}
	b.PutID(MessagesSavedGifsTypeID)
	b.PutInt(s.Hash)
	b.PutVectorHeader(len(s.Gifs))
	for idx, v := range s.Gifs {
		if v == nil {
			return fmt.Errorf("unable to encode messages.savedGifs#2e0709a5: field gifs element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.savedGifs#2e0709a5: field gifs element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSavedGifs) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.savedGifs#2e0709a5 to nil")
	}
	if err := b.ConsumeID(MessagesSavedGifsTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.savedGifs#2e0709a5: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.savedGifs#2e0709a5: field hash: %w", err)
		}
		s.Hash = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.savedGifs#2e0709a5: field gifs: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDocument(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.savedGifs#2e0709a5: field gifs: %w", err)
			}
			s.Gifs = append(s.Gifs, value)
		}
	}
	return nil
}

// construct implements constructor of MessagesSavedGifsClass.
func (s MessagesSavedGifs) construct() MessagesSavedGifsClass { return &s }

// Ensuring interfaces in compile-time for MessagesSavedGifs.
var (
	_ bin.Encoder = &MessagesSavedGifs{}
	_ bin.Decoder = &MessagesSavedGifs{}

	_ MessagesSavedGifsClass = &MessagesSavedGifs{}
)

// MessagesSavedGifsClass represents messages.SavedGifs generic type.
//
// Example:
//  g, err := DecodeMessagesSavedGifs(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *MessagesSavedGifsNotModified: // messages.savedGifsNotModified#e8025ca2
//  case *MessagesSavedGifs: // messages.savedGifs#2e0709a5
//  default: panic(v)
//  }
type MessagesSavedGifsClass interface {
	bin.Encoder
	bin.Decoder
	construct() MessagesSavedGifsClass
}

// DecodeMessagesSavedGifs implements binary de-serialization for MessagesSavedGifsClass.
func DecodeMessagesSavedGifs(buf *bin.Buffer) (MessagesSavedGifsClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesSavedGifsNotModifiedTypeID:
		// Decoding messages.savedGifsNotModified#e8025ca2.
		v := MessagesSavedGifsNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesSavedGifsClass: %w", err)
		}
		return &v, nil
	case MessagesSavedGifsTypeID:
		// Decoding messages.savedGifs#2e0709a5.
		v := MessagesSavedGifs{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesSavedGifsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesSavedGifsClass: %w", bin.NewUnexpectedID(id))
	}
}