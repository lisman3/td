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

// FileLocationToBeDeprecated represents TL type `fileLocationToBeDeprecated#bc7fc6cd`.
// Indicates the location of a photo, will be deprecated soon
//
// See https://core.telegram.org/constructor/fileLocationToBeDeprecated for reference.
type FileLocationToBeDeprecated struct {
	// Volume ID
	VolumeID int64
	// Local ID
	LocalID int
}

// FileLocationToBeDeprecatedTypeID is TL type id of FileLocationToBeDeprecated.
const FileLocationToBeDeprecatedTypeID = 0xbc7fc6cd

func (f *FileLocationToBeDeprecated) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.VolumeID == 0) {
		return false
	}
	if !(f.LocalID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *FileLocationToBeDeprecated) String() string {
	if f == nil {
		return "FileLocationToBeDeprecated(nil)"
	}
	var sb strings.Builder
	sb.WriteString("FileLocationToBeDeprecated")
	sb.WriteString("{\n")
	sb.WriteString("\tVolumeID: ")
	sb.WriteString(fmt.Sprint(f.VolumeID))
	sb.WriteString(",\n")
	sb.WriteString("\tLocalID: ")
	sb.WriteString(fmt.Sprint(f.LocalID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (f *FileLocationToBeDeprecated) TypeID() uint32 {
	return FileLocationToBeDeprecatedTypeID
}

// Encode implements bin.Encoder.
func (f *FileLocationToBeDeprecated) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode fileLocationToBeDeprecated#bc7fc6cd as nil")
	}
	b.PutID(FileLocationToBeDeprecatedTypeID)
	b.PutLong(f.VolumeID)
	b.PutInt(f.LocalID)
	return nil
}

// GetVolumeID returns value of VolumeID field.
func (f *FileLocationToBeDeprecated) GetVolumeID() (value int64) {
	return f.VolumeID
}

// GetLocalID returns value of LocalID field.
func (f *FileLocationToBeDeprecated) GetLocalID() (value int) {
	return f.LocalID
}

// Decode implements bin.Decoder.
func (f *FileLocationToBeDeprecated) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode fileLocationToBeDeprecated#bc7fc6cd to nil")
	}
	if err := b.ConsumeID(FileLocationToBeDeprecatedTypeID); err != nil {
		return fmt.Errorf("unable to decode fileLocationToBeDeprecated#bc7fc6cd: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode fileLocationToBeDeprecated#bc7fc6cd: field volume_id: %w", err)
		}
		f.VolumeID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode fileLocationToBeDeprecated#bc7fc6cd: field local_id: %w", err)
		}
		f.LocalID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for FileLocationToBeDeprecated.
var (
	_ bin.Encoder = &FileLocationToBeDeprecated{}
	_ bin.Decoder = &FileLocationToBeDeprecated{}
)
