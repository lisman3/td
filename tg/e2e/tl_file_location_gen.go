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

// FileLocationUnavailable represents TL type `fileLocationUnavailable#7c596b46`.
//
// See https://core.telegram.org/constructor/fileLocationUnavailable for reference.
type FileLocationUnavailable struct {
	// VolumeID field of FileLocationUnavailable.
	VolumeID int64
	// LocalID field of FileLocationUnavailable.
	LocalID int
	// Secret field of FileLocationUnavailable.
	Secret int64
}

// FileLocationUnavailableTypeID is TL type id of FileLocationUnavailable.
const FileLocationUnavailableTypeID = 0x7c596b46

func (f *FileLocationUnavailable) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.VolumeID == 0) {
		return false
	}
	if !(f.LocalID == 0) {
		return false
	}
	if !(f.Secret == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *FileLocationUnavailable) String() string {
	if f == nil {
		return "FileLocationUnavailable(nil)"
	}
	var sb strings.Builder
	sb.WriteString("FileLocationUnavailable")
	sb.WriteString("{\n")
	sb.WriteString("\tVolumeID: ")
	sb.WriteString(fmt.Sprint(f.VolumeID))
	sb.WriteString(",\n")
	sb.WriteString("\tLocalID: ")
	sb.WriteString(fmt.Sprint(f.LocalID))
	sb.WriteString(",\n")
	sb.WriteString("\tSecret: ")
	sb.WriteString(fmt.Sprint(f.Secret))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (f *FileLocationUnavailable) TypeID() uint32 {
	return FileLocationUnavailableTypeID
}

// Encode implements bin.Encoder.
func (f *FileLocationUnavailable) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode fileLocationUnavailable#7c596b46 as nil")
	}
	b.PutID(FileLocationUnavailableTypeID)
	b.PutLong(f.VolumeID)
	b.PutInt(f.LocalID)
	b.PutLong(f.Secret)
	return nil
}

// GetVolumeID returns value of VolumeID field.
func (f *FileLocationUnavailable) GetVolumeID() (value int64) {
	return f.VolumeID
}

// GetLocalID returns value of LocalID field.
func (f *FileLocationUnavailable) GetLocalID() (value int) {
	return f.LocalID
}

// GetSecret returns value of Secret field.
func (f *FileLocationUnavailable) GetSecret() (value int64) {
	return f.Secret
}

// Decode implements bin.Decoder.
func (f *FileLocationUnavailable) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode fileLocationUnavailable#7c596b46 to nil")
	}
	if err := b.ConsumeID(FileLocationUnavailableTypeID); err != nil {
		return fmt.Errorf("unable to decode fileLocationUnavailable#7c596b46: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode fileLocationUnavailable#7c596b46: field volume_id: %w", err)
		}
		f.VolumeID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode fileLocationUnavailable#7c596b46: field local_id: %w", err)
		}
		f.LocalID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode fileLocationUnavailable#7c596b46: field secret: %w", err)
		}
		f.Secret = value
	}
	return nil
}

// construct implements constructor of FileLocationClass.
func (f FileLocationUnavailable) construct() FileLocationClass { return &f }

// Ensuring interfaces in compile-time for FileLocationUnavailable.
var (
	_ bin.Encoder = &FileLocationUnavailable{}
	_ bin.Decoder = &FileLocationUnavailable{}

	_ FileLocationClass = &FileLocationUnavailable{}
)

// FileLocation represents TL type `fileLocation#53d69076`.
//
// See https://core.telegram.org/constructor/fileLocation for reference.
type FileLocation struct {
	// DCID field of FileLocation.
	DCID int
	// VolumeID field of FileLocation.
	VolumeID int64
	// LocalID field of FileLocation.
	LocalID int
	// Secret field of FileLocation.
	Secret int64
}

// FileLocationTypeID is TL type id of FileLocation.
const FileLocationTypeID = 0x53d69076

func (f *FileLocation) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.DCID == 0) {
		return false
	}
	if !(f.VolumeID == 0) {
		return false
	}
	if !(f.LocalID == 0) {
		return false
	}
	if !(f.Secret == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *FileLocation) String() string {
	if f == nil {
		return "FileLocation(nil)"
	}
	var sb strings.Builder
	sb.WriteString("FileLocation")
	sb.WriteString("{\n")
	sb.WriteString("\tDCID: ")
	sb.WriteString(fmt.Sprint(f.DCID))
	sb.WriteString(",\n")
	sb.WriteString("\tVolumeID: ")
	sb.WriteString(fmt.Sprint(f.VolumeID))
	sb.WriteString(",\n")
	sb.WriteString("\tLocalID: ")
	sb.WriteString(fmt.Sprint(f.LocalID))
	sb.WriteString(",\n")
	sb.WriteString("\tSecret: ")
	sb.WriteString(fmt.Sprint(f.Secret))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (f *FileLocation) TypeID() uint32 {
	return FileLocationTypeID
}

// Encode implements bin.Encoder.
func (f *FileLocation) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode fileLocation#53d69076 as nil")
	}
	b.PutID(FileLocationTypeID)
	b.PutInt(f.DCID)
	b.PutLong(f.VolumeID)
	b.PutInt(f.LocalID)
	b.PutLong(f.Secret)
	return nil
}

// GetDCID returns value of DCID field.
func (f *FileLocation) GetDCID() (value int) {
	return f.DCID
}

// GetVolumeID returns value of VolumeID field.
func (f *FileLocation) GetVolumeID() (value int64) {
	return f.VolumeID
}

// GetLocalID returns value of LocalID field.
func (f *FileLocation) GetLocalID() (value int) {
	return f.LocalID
}

// GetSecret returns value of Secret field.
func (f *FileLocation) GetSecret() (value int64) {
	return f.Secret
}

// Decode implements bin.Decoder.
func (f *FileLocation) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode fileLocation#53d69076 to nil")
	}
	if err := b.ConsumeID(FileLocationTypeID); err != nil {
		return fmt.Errorf("unable to decode fileLocation#53d69076: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode fileLocation#53d69076: field dc_id: %w", err)
		}
		f.DCID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode fileLocation#53d69076: field volume_id: %w", err)
		}
		f.VolumeID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode fileLocation#53d69076: field local_id: %w", err)
		}
		f.LocalID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode fileLocation#53d69076: field secret: %w", err)
		}
		f.Secret = value
	}
	return nil
}

// construct implements constructor of FileLocationClass.
func (f FileLocation) construct() FileLocationClass { return &f }

// Ensuring interfaces in compile-time for FileLocation.
var (
	_ bin.Encoder = &FileLocation{}
	_ bin.Decoder = &FileLocation{}

	_ FileLocationClass = &FileLocation{}
)

// FileLocationClass represents FileLocation generic type.
//
// See https://core.telegram.org/type/FileLocation for reference.
//
// Example:
//  g, err := DecodeFileLocation(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *FileLocationUnavailable: // fileLocationUnavailable#7c596b46
//  case *FileLocation: // fileLocation#53d69076
//  default: panic(v)
//  }
type FileLocationClass interface {
	bin.Encoder
	bin.Decoder
	construct() FileLocationClass

	// VolumeID field of FileLocationUnavailable.
	GetVolumeID() (value int64)
	// LocalID field of FileLocationUnavailable.
	GetLocalID() (value int)
	// Secret field of FileLocationUnavailable.
	GetSecret() (value int64)

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeFileLocation implements binary de-serialization for FileLocationClass.
func DecodeFileLocation(buf *bin.Buffer) (FileLocationClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case FileLocationUnavailableTypeID:
		// Decoding fileLocationUnavailable#7c596b46.
		v := FileLocationUnavailable{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode FileLocationClass: %w", err)
		}
		return &v, nil
	case FileLocationTypeID:
		// Decoding fileLocation#53d69076.
		v := FileLocation{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode FileLocationClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode FileLocationClass: %w", bin.NewUnexpectedID(id))
	}
}

// FileLocation boxes the FileLocationClass providing a helper.
type FileLocationBox struct {
	FileLocation FileLocationClass
}

// Decode implements bin.Decoder for FileLocationBox.
func (b *FileLocationBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode FileLocationBox to nil")
	}
	v, err := DecodeFileLocation(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.FileLocation = v
	return nil
}

// Encode implements bin.Encode for FileLocationBox.
func (b *FileLocationBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.FileLocation == nil {
		return fmt.Errorf("unable to encode FileLocationClass as nil")
	}
	return b.FileLocation.Encode(buf)
}
