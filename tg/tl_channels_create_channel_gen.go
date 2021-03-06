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

// ChannelsCreateChannelRequest represents TL type `channels.createChannel#3d5fb10f`.
// Create a supergroup/channel¹.
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.createChannel for reference.
type ChannelsCreateChannelRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to create a channel¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Broadcast bool
	// Whether to create a supergroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Megagroup bool
	// ForImport field of ChannelsCreateChannelRequest.
	ForImport bool
	// Channel title
	Title string
	// Channel description
	About string
	// Geogroup location
	//
	// Use SetGeoPoint and GetGeoPoint helpers.
	GeoPoint InputGeoPointClass
	// Geogroup address
	//
	// Use SetAddress and GetAddress helpers.
	Address string
}

// ChannelsCreateChannelRequestTypeID is TL type id of ChannelsCreateChannelRequest.
const ChannelsCreateChannelRequestTypeID = 0x3d5fb10f

func (c *ChannelsCreateChannelRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Broadcast == false) {
		return false
	}
	if !(c.Megagroup == false) {
		return false
	}
	if !(c.ForImport == false) {
		return false
	}
	if !(c.Title == "") {
		return false
	}
	if !(c.About == "") {
		return false
	}
	if !(c.GeoPoint == nil) {
		return false
	}
	if !(c.Address == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelsCreateChannelRequest) String() string {
	if c == nil {
		return "ChannelsCreateChannelRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChannelsCreateChannelRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(c.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tTitle: ")
	sb.WriteString(fmt.Sprint(c.Title))
	sb.WriteString(",\n")
	sb.WriteString("\tAbout: ")
	sb.WriteString(fmt.Sprint(c.About))
	sb.WriteString(",\n")
	if c.Flags.Has(2) {
		sb.WriteString("\tGeoPoint: ")
		sb.WriteString(fmt.Sprint(c.GeoPoint))
		sb.WriteString(",\n")
	}
	if c.Flags.Has(2) {
		sb.WriteString("\tAddress: ")
		sb.WriteString(fmt.Sprint(c.Address))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *ChannelsCreateChannelRequest) TypeID() uint32 {
	return ChannelsCreateChannelRequestTypeID
}

// Encode implements bin.Encoder.
func (c *ChannelsCreateChannelRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channels.createChannel#3d5fb10f as nil")
	}
	b.PutID(ChannelsCreateChannelRequestTypeID)
	if !(c.Broadcast == false) {
		c.Flags.Set(0)
	}
	if !(c.Megagroup == false) {
		c.Flags.Set(1)
	}
	if !(c.ForImport == false) {
		c.Flags.Set(3)
	}
	if !(c.GeoPoint == nil) {
		c.Flags.Set(2)
	}
	if !(c.Address == "") {
		c.Flags.Set(2)
	}
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.createChannel#3d5fb10f: field flags: %w", err)
	}
	b.PutString(c.Title)
	b.PutString(c.About)
	if c.Flags.Has(2) {
		if c.GeoPoint == nil {
			return fmt.Errorf("unable to encode channels.createChannel#3d5fb10f: field geo_point is nil")
		}
		if err := c.GeoPoint.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channels.createChannel#3d5fb10f: field geo_point: %w", err)
		}
	}
	if c.Flags.Has(2) {
		b.PutString(c.Address)
	}
	return nil
}

// SetBroadcast sets value of Broadcast conditional field.
func (c *ChannelsCreateChannelRequest) SetBroadcast(value bool) {
	if value {
		c.Flags.Set(0)
		c.Broadcast = true
	} else {
		c.Flags.Unset(0)
		c.Broadcast = false
	}
}

// GetBroadcast returns value of Broadcast conditional field.
func (c *ChannelsCreateChannelRequest) GetBroadcast() (value bool) {
	return c.Flags.Has(0)
}

// SetMegagroup sets value of Megagroup conditional field.
func (c *ChannelsCreateChannelRequest) SetMegagroup(value bool) {
	if value {
		c.Flags.Set(1)
		c.Megagroup = true
	} else {
		c.Flags.Unset(1)
		c.Megagroup = false
	}
}

// GetMegagroup returns value of Megagroup conditional field.
func (c *ChannelsCreateChannelRequest) GetMegagroup() (value bool) {
	return c.Flags.Has(1)
}

// SetForImport sets value of ForImport conditional field.
func (c *ChannelsCreateChannelRequest) SetForImport(value bool) {
	if value {
		c.Flags.Set(3)
		c.ForImport = true
	} else {
		c.Flags.Unset(3)
		c.ForImport = false
	}
}

// GetForImport returns value of ForImport conditional field.
func (c *ChannelsCreateChannelRequest) GetForImport() (value bool) {
	return c.Flags.Has(3)
}

// GetTitle returns value of Title field.
func (c *ChannelsCreateChannelRequest) GetTitle() (value string) {
	return c.Title
}

// GetAbout returns value of About field.
func (c *ChannelsCreateChannelRequest) GetAbout() (value string) {
	return c.About
}

// SetGeoPoint sets value of GeoPoint conditional field.
func (c *ChannelsCreateChannelRequest) SetGeoPoint(value InputGeoPointClass) {
	c.Flags.Set(2)
	c.GeoPoint = value
}

// GetGeoPoint returns value of GeoPoint conditional field and
// boolean which is true if field was set.
func (c *ChannelsCreateChannelRequest) GetGeoPoint() (value InputGeoPointClass, ok bool) {
	if !c.Flags.Has(2) {
		return value, false
	}
	return c.GeoPoint, true
}

// SetAddress sets value of Address conditional field.
func (c *ChannelsCreateChannelRequest) SetAddress(value string) {
	c.Flags.Set(2)
	c.Address = value
}

// GetAddress returns value of Address conditional field and
// boolean which is true if field was set.
func (c *ChannelsCreateChannelRequest) GetAddress() (value string, ok bool) {
	if !c.Flags.Has(2) {
		return value, false
	}
	return c.Address, true
}

// Decode implements bin.Decoder.
func (c *ChannelsCreateChannelRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channels.createChannel#3d5fb10f to nil")
	}
	if err := b.ConsumeID(ChannelsCreateChannelRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.createChannel#3d5fb10f: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channels.createChannel#3d5fb10f: field flags: %w", err)
		}
	}
	c.Broadcast = c.Flags.Has(0)
	c.Megagroup = c.Flags.Has(1)
	c.ForImport = c.Flags.Has(3)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channels.createChannel#3d5fb10f: field title: %w", err)
		}
		c.Title = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channels.createChannel#3d5fb10f: field about: %w", err)
		}
		c.About = value
	}
	if c.Flags.Has(2) {
		value, err := DecodeInputGeoPoint(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.createChannel#3d5fb10f: field geo_point: %w", err)
		}
		c.GeoPoint = value
	}
	if c.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channels.createChannel#3d5fb10f: field address: %w", err)
		}
		c.Address = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsCreateChannelRequest.
var (
	_ bin.Encoder = &ChannelsCreateChannelRequest{}
	_ bin.Decoder = &ChannelsCreateChannelRequest{}
)

// ChannelsCreateChannel invokes method channels.createChannel#3d5fb10f returning error if any.
// Create a supergroup/channel¹.
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// Possible errors:
//  400 CHANNELS_TOO_MUCH: You have joined too many channels/supergroups
//  400 CHAT_ABOUT_TOO_LONG: Chat about too long
//  400 CHAT_TITLE_EMPTY: No chat title provided
//  403 USER_RESTRICTED: You're spamreported, you can't create channels or chats.
//
// See https://core.telegram.org/method/channels.createChannel for reference.
func (c *Client) ChannelsCreateChannel(ctx context.Context, request *ChannelsCreateChannelRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
