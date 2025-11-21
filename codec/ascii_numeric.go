package codec

import (
	"fmt"
	"unicode/utf8"

	"github.com/rwbm/siso/padder"
	"github.com/rwbm/siso/parser"
	"github.com/rwbm/siso/prefixer"
)

func NewAsciiNumeric(value string) *AsciiNumeric {
	an := &AsciiNumeric{
		value:    []byte(value),
		prefixer: prefixer.None,
		padder:   padder.LeftZero,
		parser:   parser.Ascii,
	}
	return an
}

type AsciiNumeric struct {
	value    []byte
	prefixer prefixer.Prefixer
	padder   padder.Padder
	parser   parser.Parser
}

func (a *AsciiNumeric) String() string {
	return string(a.value)
}

func (a *AsciiNumeric) Length() int {
	return len(a.value)
}

func (a *AsciiNumeric) Encode(value string) ([]byte, error) {
	a.value = []byte(value)

	if err := a.ensureNumeric(); err != nil {
		return nil, err
	}

	content := a.value
	prefixLen := 0
	if a.prefixer != nil {
		prefixLen = a.prefixer.PackedLen()
	}

	out := make([]byte, prefixLen+len(content))
	copy(out[prefixLen:], content)

	if a.prefixer != nil {
		if err := a.prefixer.Encode(len(content), out); err != nil {
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

	a.value = []byte(value)
	return value, nil
}

func (a *AsciiNumeric) ensureNumeric() error {
	if !utf8.Valid(a.value) {
		return fmt.Errorf("encode: value is not valid UTF-8")
	}
	for idx, r := range string(a.value) {
		if r < '0' || r > '9' {
			return fmt.Errorf("encode: non-numeric character %q at position %d", r, idx)
		}
	}
	return nil
}
