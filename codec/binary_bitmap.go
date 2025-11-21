package codec

import (
	"errors"
	"fmt"

	"github.com/rwbm/siso/padder"
	"github.com/rwbm/siso/prefixer"
)

// NewBinaryBitmap creates a binary bitmap from a bit-string (64 or 128 chars of 0/1).
func NewBinaryBitmap(bits string) *BinaryBitmap {
	return &BinaryBitmap{
		value:    []byte(bits),
		prefixer: prefixer.None,
		padder:   padder.None,
	}
}

// NewEmptyBinaryBitmap returns a 64-bit bitmap initialized to all zeros.
func NewEmptyBinaryBitmap() *BinaryBitmap {
	return NewBinaryBitmap("")
}

// BinaryBitmap represents a packed bitmap (binary representation).
type BinaryBitmap struct {
	value    []byte // holds the bits ('0'/'1'), not the packed form
	prefixer prefixer.Prefixer
	padder   padder.Padder
}

func (b *BinaryBitmap) ensureInitialized() {
	if len(b.value) == 0 {
		b.value = make([]byte, 64)
		for i := range b.value {
			b.value[i] = '0'
		}
	}
}

func (b *BinaryBitmap) String() string {
	return string(b.value)
}

func (b *BinaryBitmap) Length() int {
	return len(b.value)
}

func (b *BinaryBitmap) IsSet(pos int) bool {
	if pos < 1 || pos > 128 {
		return false
	}
	if len(b.value) == 0 || pos > len(b.value) {
		return false
	}
	return b.value[pos-1] == '1'
}

func (b *BinaryBitmap) Set(pos int) error {
	if err := validatePosition(pos); err != nil {
		return err
	}

	b.ensureInitialized()

	if pos > 64 && len(b.value) == 64 {
		expanded := make([]byte, 128)
		copy(expanded, b.value)
		for idx := 64; idx < 128; idx++ {
			expanded[idx] = '0'
		}
		b.value = expanded
	}

	bits := make([]byte, len(b.value))
	copy(bits, b.value)

	if pos > 64 {
		bits[0] = '1'
	}

	bits[pos-1] = '1'
	b.value = bits
	return nil
}

func (b *BinaryBitmap) Clear(pos int) error {
	if err := validatePosition(pos); err != nil {
		return err
	}

	b.ensureInitialized()
	if pos > len(b.value) {
		return nil
	}

	bits := make([]byte, len(b.value))
	copy(bits, b.value)
	bits[pos-1] = '0'

	if len(bits) == 128 && !secondaryBitsSet(bits) {
		bits[0] = '0'
		primary := make([]byte, 64)
		copy(primary, bits[:64])
		bits = primary
	}

	b.value = bits
	return nil
}

// Bitmap returns the positions (1-based) of all bits that are set.
func (b *BinaryBitmap) Bitmap() []int {
	b.ensureInitialized()

	positions := make([]int, 0, len(b.value)/2)
	for idx, bit := range b.value {
		if bit == '1' {
			positions = append(positions, idx+1)
		}
	}
	return positions
}

// Encode produces the packed binary representation (8 or 16 bytes).
func (b *BinaryBitmap) Encode(value string) ([]byte, error) {
	b.value = []byte(value)

	if len(b.value) == 0 {
		return nil, errors.New("bitmap value is empty")
	}

	if len(b.value) != 64 && len(b.value) != 128 {
		return nil, fmt.Errorf("bitmap length must be 64 or 128 bits, got %d", len(b.value))
	}

	if len(b.value) == 64 && b.value[0] == '1' {
		return nil, errors.New("secondary bitmap flag set but only 64 bits provided")
	}

	for idx, c := range b.value {
		if c != '0' && c != '1' {
			return nil, fmt.Errorf("invalid bit %q at position %d", c, idx)
		}
	}

	out := make([]byte, 0, len(b.value)/8)
	for i := 0; i < len(b.value); i += 8 {
		bVal, err := parseByte(b.value[i : i+8])
		if err != nil {
			return nil, fmt.Errorf("failed to parse byte at offset %d: %w", i, err)
		}
		out = append(out, bVal)
	}

	return out, nil
}

// Decode reads the packed binary representation and populates the bit string.
func (b *BinaryBitmap) Decode(data []byte) (string, error) {
	if len(data) == 0 {
		return "", errors.New("bitmap data is empty")
	}

	if len(data) != 8 && len(data) != 16 {
		return "", fmt.Errorf("bitmap data must be 8 or 16 bytes, got %d", len(data))
	}

	bits := make([]byte, 0, len(data)*8)
	for pos, v := range data {
		for bit := 7; bit >= 0; bit-- {
			mask := byte(1 << bit)
			if v&mask != 0 {
				bits = append(bits, '1')
			} else {
				bits = append(bits, '0')
			}
		}
		// pos unused, but left for clarity of loop variable
		_ = pos
	}

	if len(data) == 8 && bits[0] == '1' {
		return "", errors.New("secondary bitmap indicated but only primary bitmap provided")
	}

	b.value = bits
	return string(bits), nil
}

func parseByte(bits []byte) (byte, error) {
	if len(bits) != 8 {
		return 0, fmt.Errorf("byte must have 8 bits, got %d", len(bits))
	}
	var v byte
	for idx, c := range bits {
		switch c {
		case '0':
			// no-op
		case '1':
			v |= 1 << (7 - idx)
		default:
			return 0, fmt.Errorf("invalid bit %q", c)
		}
	}
	return v, nil
}
