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

// PostAddress represents TL type `postAddress#1e8caaeb`.
// Shipping address
//
// See https://core.telegram.org/constructor/postAddress for reference.
type PostAddress struct {
	// First line for the address
	StreetLine1 string
	// Second line for the address
	StreetLine2 string
	// City
	City string
	// State, if applicable (empty otherwise)
	State string
	// ISO 3166-1 alpha-2 country code
	CountryIso2 string
	// Address post code
	PostCode string
}

// PostAddressTypeID is TL type id of PostAddress.
const PostAddressTypeID = 0x1e8caaeb

func (p *PostAddress) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.StreetLine1 == "") {
		return false
	}
	if !(p.StreetLine2 == "") {
		return false
	}
	if !(p.City == "") {
		return false
	}
	if !(p.State == "") {
		return false
	}
	if !(p.CountryIso2 == "") {
		return false
	}
	if !(p.PostCode == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PostAddress) String() string {
	if p == nil {
		return "PostAddress(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PostAddress")
	sb.WriteString("{\n")
	sb.WriteString("\tStreetLine1: ")
	sb.WriteString(fmt.Sprint(p.StreetLine1))
	sb.WriteString(",\n")
	sb.WriteString("\tStreetLine2: ")
	sb.WriteString(fmt.Sprint(p.StreetLine2))
	sb.WriteString(",\n")
	sb.WriteString("\tCity: ")
	sb.WriteString(fmt.Sprint(p.City))
	sb.WriteString(",\n")
	sb.WriteString("\tState: ")
	sb.WriteString(fmt.Sprint(p.State))
	sb.WriteString(",\n")
	sb.WriteString("\tCountryIso2: ")
	sb.WriteString(fmt.Sprint(p.CountryIso2))
	sb.WriteString(",\n")
	sb.WriteString("\tPostCode: ")
	sb.WriteString(fmt.Sprint(p.PostCode))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PostAddress) TypeID() uint32 {
	return PostAddressTypeID
}

// Encode implements bin.Encoder.
func (p *PostAddress) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode postAddress#1e8caaeb as nil")
	}
	b.PutID(PostAddressTypeID)
	b.PutString(p.StreetLine1)
	b.PutString(p.StreetLine2)
	b.PutString(p.City)
	b.PutString(p.State)
	b.PutString(p.CountryIso2)
	b.PutString(p.PostCode)
	return nil
}

// GetStreetLine1 returns value of StreetLine1 field.
func (p *PostAddress) GetStreetLine1() (value string) {
	return p.StreetLine1
}

// GetStreetLine2 returns value of StreetLine2 field.
func (p *PostAddress) GetStreetLine2() (value string) {
	return p.StreetLine2
}

// GetCity returns value of City field.
func (p *PostAddress) GetCity() (value string) {
	return p.City
}

// GetState returns value of State field.
func (p *PostAddress) GetState() (value string) {
	return p.State
}

// GetCountryIso2 returns value of CountryIso2 field.
func (p *PostAddress) GetCountryIso2() (value string) {
	return p.CountryIso2
}

// GetPostCode returns value of PostCode field.
func (p *PostAddress) GetPostCode() (value string) {
	return p.PostCode
}

// Decode implements bin.Decoder.
func (p *PostAddress) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode postAddress#1e8caaeb to nil")
	}
	if err := b.ConsumeID(PostAddressTypeID); err != nil {
		return fmt.Errorf("unable to decode postAddress#1e8caaeb: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode postAddress#1e8caaeb: field street_line1: %w", err)
		}
		p.StreetLine1 = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode postAddress#1e8caaeb: field street_line2: %w", err)
		}
		p.StreetLine2 = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode postAddress#1e8caaeb: field city: %w", err)
		}
		p.City = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode postAddress#1e8caaeb: field state: %w", err)
		}
		p.State = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode postAddress#1e8caaeb: field country_iso2: %w", err)
		}
		p.CountryIso2 = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode postAddress#1e8caaeb: field post_code: %w", err)
		}
		p.PostCode = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PostAddress.
var (
	_ bin.Encoder = &PostAddress{}
	_ bin.Decoder = &PostAddress{}
)
