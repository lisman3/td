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

// MessagesGetEmojiKeywordsRequest represents TL type `messages.getEmojiKeywords#35a0e062`.
// Get localized emoji keywords
//
// See https://core.telegram.org/method/messages.getEmojiKeywords for reference.
type MessagesGetEmojiKeywordsRequest struct {
	// Language code
	LangCode string
}

// MessagesGetEmojiKeywordsRequestTypeID is TL type id of MessagesGetEmojiKeywordsRequest.
const MessagesGetEmojiKeywordsRequestTypeID = 0x35a0e062

func (g *MessagesGetEmojiKeywordsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.LangCode == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetEmojiKeywordsRequest) String() string {
	if g == nil {
		return "MessagesGetEmojiKeywordsRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesGetEmojiKeywordsRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tLangCode: ")
	sb.WriteString(fmt.Sprint(g.LangCode))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *MessagesGetEmojiKeywordsRequest) TypeID() uint32 {
	return MessagesGetEmojiKeywordsRequestTypeID
}

// Encode implements bin.Encoder.
func (g *MessagesGetEmojiKeywordsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getEmojiKeywords#35a0e062 as nil")
	}
	b.PutID(MessagesGetEmojiKeywordsRequestTypeID)
	b.PutString(g.LangCode)
	return nil
}

// GetLangCode returns value of LangCode field.
func (g *MessagesGetEmojiKeywordsRequest) GetLangCode() (value string) {
	return g.LangCode
}

// Decode implements bin.Decoder.
func (g *MessagesGetEmojiKeywordsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getEmojiKeywords#35a0e062 to nil")
	}
	if err := b.ConsumeID(MessagesGetEmojiKeywordsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getEmojiKeywords#35a0e062: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getEmojiKeywords#35a0e062: field lang_code: %w", err)
		}
		g.LangCode = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetEmojiKeywordsRequest.
var (
	_ bin.Encoder = &MessagesGetEmojiKeywordsRequest{}
	_ bin.Decoder = &MessagesGetEmojiKeywordsRequest{}
)

// MessagesGetEmojiKeywords invokes method messages.getEmojiKeywords#35a0e062 returning error if any.
// Get localized emoji keywords
//
// See https://core.telegram.org/method/messages.getEmojiKeywords for reference.
func (c *Client) MessagesGetEmojiKeywords(ctx context.Context, langcode string) (*EmojiKeywordsDifference, error) {
	var result EmojiKeywordsDifference

	request := &MessagesGetEmojiKeywordsRequest{
		LangCode: langcode,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
