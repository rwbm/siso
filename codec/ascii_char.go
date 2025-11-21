package codec

import (
	"fmt"
	"unicode/utf8"

	"github.com/rwbm/siso/padder"
	"github.com/rwbm/siso/parser"
	"github.com/rwbm/siso/prefixer"
)

func NewAsciiChar(value string) *AsciiChar {
	ac := &AsciiChar{
		value:    []byte(value),
		prefixer: prefixer.None,
		padder:   padder.LeftZero,
		parser:   parser.Ascii,
	}
	return ac
}

type AsciiChar struct {
	value    []byte
	prefixer prefixer.Prefixer
	padder   padder.Padder
	parser   parser.Parser
}

func (a *AsciiChar) String() string {
	return string(a.value)
}

func (a *AsciiChar) Length() int {
	return len(a.value)
}

func (a *AsciiChar) Encode(value string) ([]byte, error) {
	a.value = []byte(value)

	if err := a.ensureASCII(); err != nil {
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

	a.value = []byte(value)
	return value, nil
}

func (a *AsciiChar) ensureASCII() error {
	if !utf8.Valid(a.value) {
		return fmt.Errorf("encode: value is not valid UTF-8")
	}
	for idx, r := range string(a.value) {
		if r > 0x7F {
			return fmt.Errorf("encode: non-ASCII rune %q at position %d", r, idx)
		}
	}
	return nil
}
