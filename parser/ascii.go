package parser

import (
	"fmt"

	"github.com/rwbm/siso/padder"
)

// Ascii is the default parser for ASCII encoded fields.
var Ascii Parser = &AsciiParser{}

type AsciiParser struct{}

func (a *AsciiParser) Parse(data []byte, length int, pad padder.Padder) (string, []byte, error) {
	if length < 0 {
		return "", data, fmt.Errorf("length must be non-negative, got %d", length)
	}
	if len(data) < length {
		return "", data, fmt.Errorf("insufficient data: need %d bytes, have %d", length, len(data))
	}

	segment := data[:length]

	for idx, b := range segment {
		if b > 0x7F {
			return "", data, fmt.Errorf("non-ASCII byte 0x%X at position %d", b, idx)
		}
	}

	value := string(segment)
	if pad != nil {
		value = pad.Unpad(value)
	}

	for idx, r := range value {
		if r < '0' || r > '9' {
			return "", data, fmt.Errorf("non-numeric character %q after unpadding at position %d", r, idx)
		}
	}

	return value, data[length:], nil
}
