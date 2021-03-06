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

// StickersRemoveStickerFromSetRequest represents TL type `stickers.removeStickerFromSet#f7760f51`.
// Remove a sticker from the set where it belongs, bots only. The sticker set must have been created by the bot.
//
// See https://core.telegram.org/method/stickers.removeStickerFromSet for reference.
type StickersRemoveStickerFromSetRequest struct {
	// The sticker to remove
	Sticker InputDocumentClass
}

// StickersRemoveStickerFromSetRequestTypeID is TL type id of StickersRemoveStickerFromSetRequest.
const StickersRemoveStickerFromSetRequestTypeID = 0xf7760f51

func (r *StickersRemoveStickerFromSetRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Sticker == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *StickersRemoveStickerFromSetRequest) String() string {
	if r == nil {
		return "StickersRemoveStickerFromSetRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("StickersRemoveStickerFromSetRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tSticker: ")
	sb.WriteString(fmt.Sprint(r.Sticker))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (r *StickersRemoveStickerFromSetRequest) TypeID() uint32 {
	return StickersRemoveStickerFromSetRequestTypeID
}

// Encode implements bin.Encoder.
func (r *StickersRemoveStickerFromSetRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode stickers.removeStickerFromSet#f7760f51 as nil")
	}
	b.PutID(StickersRemoveStickerFromSetRequestTypeID)
	if r.Sticker == nil {
		return fmt.Errorf("unable to encode stickers.removeStickerFromSet#f7760f51: field sticker is nil")
	}
	if err := r.Sticker.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stickers.removeStickerFromSet#f7760f51: field sticker: %w", err)
	}
	return nil
}

// GetSticker returns value of Sticker field.
func (r *StickersRemoveStickerFromSetRequest) GetSticker() (value InputDocumentClass) {
	return r.Sticker
}

// Decode implements bin.Decoder.
func (r *StickersRemoveStickerFromSetRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode stickers.removeStickerFromSet#f7760f51 to nil")
	}
	if err := b.ConsumeID(StickersRemoveStickerFromSetRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stickers.removeStickerFromSet#f7760f51: %w", err)
	}
	{
		value, err := DecodeInputDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode stickers.removeStickerFromSet#f7760f51: field sticker: %w", err)
		}
		r.Sticker = value
	}
	return nil
}

// Ensuring interfaces in compile-time for StickersRemoveStickerFromSetRequest.
var (
	_ bin.Encoder = &StickersRemoveStickerFromSetRequest{}
	_ bin.Decoder = &StickersRemoveStickerFromSetRequest{}
)

// StickersRemoveStickerFromSet invokes method stickers.removeStickerFromSet#f7760f51 returning error if any.
// Remove a sticker from the set where it belongs, bots only. The sticker set must have been created by the bot.
//
// Possible errors:
//  400 BOT_MISSING: This method can only be run by a bot
//  400 STICKER_INVALID: The provided sticker is invalid
//
// See https://core.telegram.org/method/stickers.removeStickerFromSet for reference.
// Can be used by bots.
func (c *Client) StickersRemoveStickerFromSet(ctx context.Context, sticker InputDocumentClass) (*MessagesStickerSet, error) {
	var result MessagesStickerSet

	request := &StickersRemoveStickerFromSetRequest{
		Sticker: sticker,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
