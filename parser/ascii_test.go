package parser

import (
	"strings"
	"testing"

	"github.com/rwbm/siso/padder"
)

func TestAsciiParserSuccess(t *testing.T) {
	value, rest, err := Ascii.Parse([]byte("000123XYZ"), 6, padder.LeftZero)
	if err != nil {
		t.Fatalf("Parse returned error: %v", err)
	}
	if value != "123" {
		t.Fatalf("Parse value = %q, want %q", value, "123")
	}
	if string(rest) != "XYZ" {
		t.Fatalf("Parse rest = %q, want %q", string(rest), "XYZ")
	}
}

func TestAsciiParserWithoutPadder(t *testing.T) {
	value, rest, err := Ascii.Parse([]byte("9876"), 4, nil)
	if err != nil {
		t.Fatalf("Parse returned error: %v", err)
	}
	if value != "9876" {
		t.Fatalf("Parse value = %q, want %q", value, "9876")
	}
	if len(rest) != 0 {
		t.Fatalf("Parse rest length = %d, want 0", len(rest))
	}
}

func TestAsciiParserErrors(t *testing.T) {
	tests := []struct {
		name   string
		data   []byte
		length int
		pad    padder.Padder
		want   string
	}{
		{
			name:   "insufficient data",
			data:   []byte("12"),
			length: 3,
			pad:    padder.LeftZero,
			want:   "insufficient data",
		},
		{
			name:   "non ASCII byte",
			data:   []byte{0x31, 0xFF},
			length: 2,
			pad:    padder.LeftZero,
			want:   "non-ASCII",
		},
		{
			name:   "non numeric after unpadding",
			data:   []byte("00A1"),
			length: 4,
			pad:    padder.LeftZero,
			want:   "non-numeric",
		},
		{
			name:   "negative length",
			data:   []byte("123"),
			length: -1,
			pad:    padder.LeftZero,
			want:   "length must be non-negative",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if _, _, err := Ascii.Parse(tc.data, tc.length, tc.pad); err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("Parse error %v, want substring %q", err, tc.want)
			}
		})
	}
}
