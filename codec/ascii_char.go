package codec

import (
	"fmt"
	"unicode/utf8"

	"github.com/rwbm/siso/padder"
	"github.com/rwbm/siso/parser"
	"github.com/rwbm/siso/prefixer"
)

func NewAsciiChar() *AsciiChar {
	ac := &AsciiChar{
		prefixer: prefixer.None,
		padder:   padder.LeftZero,
		parser:   parser.Ascii,
	}
	return ac
}

type AsciiChar struct {
	prefixer prefixer.Prefixer
	padder   padder.Padder
	parser   parser.Parser
}

func (a *AsciiChar) Encode(value string) ([]byte, error) {
	if err := a.ensureASCII([]byte(value)); err != nil {
		return nil, err
	}

	prefixLen := 0
	if a.prefixer != nil {
		prefixLen = a.prefixer.PackedLen()
	}

	out := make([]byte, prefixLen+len(value))
	copy(out[prefixLen:], value)

	if a.prefixer != nil {
		if err := a.prefixer.Encode(len(value), out); err != nil {
			return nil, fmt.Errorf("encode prefix: %w", err)
		}
	}

	return out, nil
}

func (a *AsciiChar) Decode(data []byte) (string, error) {
	if len(data) == 0 {
		return "", fmt.Errorf("decode: no data provided")
	}

	prefixLen := 0
	if a.prefixer != nil {
		prefixLen = a.prefixer.PackedLen()
		if len(data) < prefixLen {
			return "", fmt.Errorf("decode: data shorter than prefix length")
		}
	}

	fieldLen := len(data) - prefixLen
	if a.prefixer != nil {
		l, err := a.prefixer.Decode(data, 0)
		if err != nil {
			return "", fmt.Errorf("decode prefix: %w", err)
		}
		if l >= 0 {
			fieldLen = l
		}
	}

	if fieldLen < 0 {
		return "", fmt.Errorf("decode: negative field length %d", fieldLen)
	}
	if len(data) < prefixLen+fieldLen {
		return "", fmt.Errorf("decode: insufficient data for field length %d (have %d)", fieldLen, len(data)-prefixLen)
	}

	segment := data[prefixLen : prefixLen+fieldLen]
	for idx, b := range segment {
		if b > 0x7F {
			return "", fmt.Errorf("decode: non-ASCII byte 0x%X at position %d", b, idx)
		}
	}

	value := string(segment)
	if a.padder != nil {
		value = a.padder.Unpad(value)
	}

	return value, nil
}

func (a *AsciiChar) ensureASCII(b []byte) error {
	if !utf8.Valid(b) {
		return fmt.Errorf("encode: value is not valid UTF-8")
	}
	for idx, r := range string(b) {
		if r > 0x7F {
			return fmt.Errorf("encode: non-ASCII rune %q at position %d", r, idx)
		}
	}
	return nil
}
