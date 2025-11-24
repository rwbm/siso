package codec

import (
	"fmt"
	"unicode/utf8"

	"github.com/rwbm/siso/padder"
	"github.com/rwbm/siso/parser"
	"github.com/rwbm/siso/prefixer"
)

func NewAsciiNumeric() *AsciiNumeric {
	an := &AsciiNumeric{
		prefixer: prefixer.None,
		padder:   padder.LeftZero,
		parser:   parser.Ascii,
	}
	return an
}

type AsciiNumeric struct {
	prefixer prefixer.Prefixer
	padder   padder.Padder
	parser   parser.Parser
}

func (a *AsciiNumeric) Encode(value string) ([]byte, error) {
	if err := a.ensureNumeric([]byte(value)); err != nil {
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

func (a *AsciiNumeric) Decode(data []byte) (string, error) {
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
	value, rest, err := a.parser.Parse(segment, fieldLen, a.padder)
	if err != nil {
		return "", fmt.Errorf("decode parse: %w", err)
	}
	if len(rest) > 0 {
		return "", fmt.Errorf("decode: unexpected %d trailing bytes after parse", len(rest))
	}

	return value, nil
}

func (a *AsciiNumeric) ensureNumeric(b []byte) error {
	if !utf8.Valid(b) {
		return fmt.Errorf("encode: value is not valid UTF-8")
	}
	for idx, r := range string(b) {
		if r < '0' || r > '9' {
			return fmt.Errorf("encode: non-numeric character %q at position %d", r, idx)
		}
	}
	return nil
}
