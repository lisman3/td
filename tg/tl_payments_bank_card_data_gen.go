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

// PaymentsBankCardData represents TL type `payments.bankCardData#3e24e573`.
// Credit card info, provided by the card's bank(s)
//
// See https://core.telegram.org/constructor/payments.bankCardData for reference.
type PaymentsBankCardData struct {
	// Credit card title
	Title string
	// Info URL(s) provided by the card's bank(s)
	OpenUrls []BankCardOpenUrl
}

// PaymentsBankCardDataTypeID is TL type id of PaymentsBankCardData.
const PaymentsBankCardDataTypeID = 0x3e24e573

func (b *PaymentsBankCardData) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Title == "") {
		return false
	}
	if !(b.OpenUrls == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *PaymentsBankCardData) String() string {
	if b == nil {
		return "PaymentsBankCardData(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PaymentsBankCardData")
	sb.WriteString("{\n")
	sb.WriteString("\tTitle: ")
	sb.WriteString(fmt.Sprint(b.Title))
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range b.OpenUrls {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (b *PaymentsBankCardData) TypeID() uint32 {
	return PaymentsBankCardDataTypeID
}

// Encode implements bin.Encoder.
func (b *PaymentsBankCardData) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode payments.bankCardData#3e24e573 as nil")
	}
	buf.PutID(PaymentsBankCardDataTypeID)
	buf.PutString(b.Title)
	buf.PutVectorHeader(len(b.OpenUrls))
	for idx, v := range b.OpenUrls {
		if err := v.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode payments.bankCardData#3e24e573: field open_urls element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetTitle returns value of Title field.
func (b *PaymentsBankCardData) GetTitle() (value string) {
	return b.Title
}

// GetOpenUrls returns value of OpenUrls field.
func (b *PaymentsBankCardData) GetOpenUrls() (value []BankCardOpenUrl) {
	return b.OpenUrls
}

// Decode implements bin.Decoder.
func (b *PaymentsBankCardData) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode payments.bankCardData#3e24e573 to nil")
	}
	if err := buf.ConsumeID(PaymentsBankCardDataTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.bankCardData#3e24e573: %w", err)
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode payments.bankCardData#3e24e573: field title: %w", err)
		}
		b.Title = value
	}
	{
		headerLen, err := buf.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode payments.bankCardData#3e24e573: field open_urls: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value BankCardOpenUrl
			if err := value.Decode(buf); err != nil {
				return fmt.Errorf("unable to decode payments.bankCardData#3e24e573: field open_urls: %w", err)
			}
			b.OpenUrls = append(b.OpenUrls, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PaymentsBankCardData.
var (
	_ bin.Encoder = &PaymentsBankCardData{}
	_ bin.Decoder = &PaymentsBankCardData{}
)
