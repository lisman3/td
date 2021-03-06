// Code generated by gotdgen, DO NOT EDIT.

package td

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

// TestVectorBytes represents TL type `testVectorBytes#a590fb25`.
//
// See https://localhost:80/doc/constructor/testVectorBytes for reference.
type TestVectorBytes struct {
	// Value field of TestVectorBytes.
	Value [][]byte
}

// TestVectorBytesTypeID is TL type id of TestVectorBytes.
const TestVectorBytesTypeID = 0xa590fb25

func (t *TestVectorBytes) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Value == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TestVectorBytes) String() string {
	if t == nil {
		return "TestVectorBytes(nil)"
	}
	var sb strings.Builder
	sb.WriteString("TestVectorBytes")
	sb.WriteString("{\n")
	sb.WriteByte('[')
	for _, v := range t.Value {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (t *TestVectorBytes) TypeID() uint32 {
	return TestVectorBytesTypeID
}

// Encode implements bin.Encoder.
func (t *TestVectorBytes) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testVectorBytes#a590fb25 as nil")
	}
	b.PutID(TestVectorBytesTypeID)
	b.PutVectorHeader(len(t.Value))
	for _, v := range t.Value {
		b.PutBytes(v)
	}
	return nil
}

// GetValue returns value of Value field.
func (t *TestVectorBytes) GetValue() (value [][]byte) {
	return t.Value
}

// Decode implements bin.Decoder.
func (t *TestVectorBytes) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testVectorBytes#a590fb25 to nil")
	}
	if err := b.ConsumeID(TestVectorBytesTypeID); err != nil {
		return fmt.Errorf("unable to decode testVectorBytes#a590fb25: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode testVectorBytes#a590fb25: field value: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Bytes()
			if err != nil {
				return fmt.Errorf("unable to decode testVectorBytes#a590fb25: field value: %w", err)
			}
			t.Value = append(t.Value, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for TestVectorBytes.
var (
	_ bin.Encoder = &TestVectorBytes{}
	_ bin.Decoder = &TestVectorBytes{}
)
