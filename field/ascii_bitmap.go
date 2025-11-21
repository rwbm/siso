package field

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/rwbm/siso/padder"
	"github.com/rwbm/siso/prefixer"
)

const hexDigits = "0123456789ABCDEF"

func NewBitmap(value string) *AsciiBitmap {
	ab := &AsciiBitmap{
		value:    []byte(value),
		prefixer: prefixer.None,
		padder:   padder.None,
	}
	return ab
}

// NewEmptyBitmap returns a 64-bit bitmap initialized to all zeros.
func NewEmptyBitmap() *AsciiBitmap {
	return NewBitmap(strings.Repeat("0", 64))
}

// Represents a bitmap field in ASCII format.
type AsciiBitmap struct {
	value    []byte
	prefixer prefixer.Prefixer
	padder   padder.Padder
}

func (i *AsciiBitmap) ensureInitialized() {
	if len(i.value) == 0 {
		i.value = []byte(strings.Repeat("0", 64))
	}
}

func (i *AsciiBitmap) Value() string {
	return string(i.value)
}

func (i *AsciiBitmap) String() string {
	return i.Value()
}

func (i *AsciiBitmap) Length() int {
	return len(i.value)
}

func (i *AsciiBitmap) Prefixer() prefixer.Prefixer {
	return i.prefixer
}

func (i *AsciiBitmap) Padder() padder.Padder {
	return i.padder
}

func (i *AsciiBitmap) IsSet(pos int) bool {
	if pos < 1 || pos > 128 {
		return false
	}

	if len(i.value) == 0 || pos > len(i.value) {
		return false
	}

	return i.value[pos-1] == '1'
}

func (i *AsciiBitmap) Set(pos int) error {
	if err := validatePosition(pos); err != nil {
		return err
	}

	i.ensureInitialized()

	if pos > 64 && len(i.value) == 64 {
		expanded := make([]byte, 128)
		copy(expanded, i.value)
		for idx := 64; idx < 128; idx++ {
			expanded[idx] = '0'
		}
		i.value = expanded
	}

	bits := make([]byte, len(i.value))
	copy(bits, i.value)

	if pos > 64 {
		bits[0] = '1'
	}

	bits[pos-1] = '1'

	i.value = bits
	return nil
}

func (i *AsciiBitmap) Clear(pos int) error {
	if err := validatePosition(pos); err != nil {
		return err
	}

	i.ensureInitialized()

	if pos > len(i.value) {
		return nil
	}

	bits := make([]byte, len(i.value))
	copy(bits, i.value)
	bits[pos-1] = '0'

	if len(bits) == 128 && !secondaryBitsSet(bits) {
		bits[0] = '0'
		primary := make([]byte, 64)
		copy(primary, bits[:64])
		bits = primary
	}

	i.value = bits
	return nil
}

// Bitmap returns the positions (1-based) of all bits that are set.
func (i *AsciiBitmap) Bitmap() []int {
	i.ensureInitialized()

	bitsOn := make([]int, 0, len(i.value)/2)
	for idx, b := range i.value {
		if b == '1' {
			bitsOn = append(bitsOn, idx+1)
		}
	}
	return bitsOn
}

func (i *AsciiBitmap) Encode() ([]byte, error) {
	if len(i.value) == 0 {
		return nil, errors.New("bitmap value is empty")
	}

	if len(i.value)%4 != 0 {
		return nil, fmt.Errorf("bitmap length must be divisible by 4, got %d", len(i.value))
	}

	if len(i.value) != 64 && len(i.value) != 128 {
		return nil, fmt.Errorf("bitmap length must be 64 or 128 bits, got %d", len(i.value))
	}

	if len(i.value) == 64 && i.value[0] == '1' {
		return nil, errors.New("secondary bitmap flag set but only 64 bits provided")
	}

	for idx, c := range i.value {
		if c != '0' && c != '1' {
			return nil, fmt.Errorf("invalid bit %q at position %d", c, idx)
		}
	}

	var b strings.Builder
	b.Grow(len(i.value) / 4)

	for idx := 0; idx < len(i.value); idx += 4 {
		nibble := i.value[idx : idx+4]
		n, err := parseNibble(nibble)
		if err != nil {
			return nil, fmt.Errorf("failed to parse nibble at offset %d: %w", idx, err)
		}
		b.WriteByte(hexDigits[n])
	}

	return []byte(b.String()), nil
}

func (i *AsciiBitmap) Decode(data []byte) error {
	if len(data) == 0 {
		return errors.New("bitmap data is empty")
	}

	if len(data) != 16 && len(data) != 32 {
		return fmt.Errorf("bitmap data must be 16 or 32 hex characters, got %d", len(data))
	}

	hexString := strings.ToUpper(string(data))

	var bits strings.Builder
	bits.Grow(len(hexString) * 4)

	for pos, c := range hexString {
		n, err := strconv.ParseUint(string(c), 16, 4)
		if err != nil {
			return fmt.Errorf("invalid hex character %q at position %d", c, pos)
		}
		bits.WriteString(fmt.Sprintf("%04b", n))
	}

	bitmap := []byte(bits.String())

	if len(data) == 16 && bitmap[0] == '1' {
		return errors.New("secondary bitmap indicated but only primary bitmap provided")
	}

	i.value = bitmap
	return nil
}

func validatePosition(pos int) error {
	if pos < 1 || pos > 128 {
		return fmt.Errorf("bitmap position out of range: %d (valid 1-128)", pos)
	}
	return nil
}

func secondaryBitsSet(bits []byte) bool {
	for idx := 64; idx < len(bits); idx++ {
		if bits[idx] == '1' {
			return true
		}
	}
	return false
}

func parseNibble(bits []byte) (uint8, error) {
	if len(bits) != 4 {
		return 0, fmt.Errorf("nibble must have 4 bits, got %d", len(bits))
	}

	var n uint8
	for idx, c := range bits {
		switch c {
		case '0':
			// no-op
		case '1':
			n |= 1 << (3 - idx)
		default:
			return 0, fmt.Errorf("invalid bit %q", c)
		}
	}
	return n, nil
}
