package codec

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func bitsFromHex(t *testing.T, hex string) string {
	t.Helper()

	var bits strings.Builder
	bits.Grow(len(hex) * 4)

	for pos, c := range strings.ToUpper(hex) {
		n, err := strconv.ParseUint(string(c), 16, 4)
		if err != nil {
			t.Fatalf("invalid hex %q at position %d: %v", c, pos, err)
		}
		bits.WriteString(fmt.Sprintf("%04b", n))
	}

	return bits.String()
}

func TestAsciiBitmapEncodePrimary(t *testing.T) {
	bitmap := bitsFromHex(t, "0123456789ABCDEF")
	encoded, err := NewAsciiBitmap().Encode(bitmap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got, want := string(encoded), "0123456789ABCDEF"; got != want {
		t.Fatalf("unexpected encoded bitmap: got %s want %s", got, want)
	}
}

func TestAsciiBitmapEncodeSecondary(t *testing.T) {
	bitmap := bitsFromHex(t, "F123456789ABCDEF0123456789ABCDEF")
	encoded, err := NewAsciiBitmap().Encode(bitmap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got, want := string(encoded), "F123456789ABCDEF0123456789ABCDEF"; got != want {
		t.Fatalf("unexpected encoded bitmap: got %s want %s", got, want)
	}
}

func TestAsciiBitmapEncodeRejectsSecondaryWithoutBits(t *testing.T) {
	// First bit set indicates a secondary bitmap should follow.
	bitmap := bitsFromHex(t, "8000000000000000")

	if _, err := NewAsciiBitmap().Encode(bitmap); err == nil {
		t.Fatal("expected error for secondary bitmap flag without 128 bits, got nil")
	}
}

func TestAsciiBitmapDecodePrimary(t *testing.T) {
	hexBitmap := "0123456789ABCDEF"
	bitmap := bitsFromHex(t, hexBitmap)

	field := NewAsciiBitmap()
	if _, err := field.Decode([]byte(hexBitmap)); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if field.String() != bitmap {
		t.Fatalf("decoded bitmap mismatch: got %q want %q", field.String(), bitmap)
	}
}

func TestAsciiBitmapDecodeSecondary(t *testing.T) {
	hexBitmap := "F123456789ABCDEF0123456789ABCDEF"
	bitmap := bitsFromHex(t, hexBitmap)

	field := NewAsciiBitmap()
	if _, err := field.Decode([]byte(hexBitmap)); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(field.String()) != 128 {
		t.Fatalf("expected 128 bits, got %d", len(field.String()))
	}

	if field.String() != bitmap {
		t.Fatalf("decoded bitmap mismatch: got %q want %q", field.String(), bitmap)
	}
}

func TestAsciiBitmapDecodeMissingSecondary(t *testing.T) {
	// Primary bitmap indicates a secondary bitmap should follow.
	if _, err := NewAsciiBitmap().Decode([]byte("8000000000000000")); err == nil {
		t.Fatal("expected error when secondary bitmap flag set but no secondary bitmap provided")
	}
}

func TestAsciiBitmapDecodeRejectsInvalidHex(t *testing.T) {
	if _, err := NewAsciiBitmap().Decode([]byte("ZZ00000000000000")); err == nil {
		t.Fatal("expected error for invalid hex characters, got nil")
	}
}

func TestAsciiBitmapSetPrimaryBit(t *testing.T) {
	bm := NewAsciiBitmap()

	if err := bm.Set(3); err != nil {
		t.Fatalf("set returned error: %v", err)
	}

	if bm.Length() != 64 {
		t.Fatalf("expected 64-bit bitmap, got %d", bm.Length())
	}

	if !bm.IsSet(3) {
		t.Fatal("expected bit 3 to be set")
	}
}

func TestAsciiBitmapSetSecondaryBitExtends(t *testing.T) {
	bm := NewAsciiBitmap()

	if err := bm.Set(70); err != nil {
		t.Fatalf("set returned error: %v", err)
	}

	if bm.Length() != 128 {
		t.Fatalf("expected bitmap to expand to 128 bits, got %d", bm.Length())
	}

	if !bm.IsSet(1) || !bm.IsSet(70) {
		t.Fatalf("expected bits 1 and 70 to be set, got bit1 %v bit70 %v", bm.IsSet(1), bm.IsSet(70))
	}
}

func TestAsciiBitmapClearSecondaryBitShrinks(t *testing.T) {
	bm := NewAsciiBitmap()

	if err := bm.Set(70); err != nil {
		t.Fatalf("set returned error: %v", err)
	}

	if err := bm.Clear(70); err != nil {
		t.Fatalf("clear returned error: %v", err)
	}

	if bm.Length() != 64 {
		t.Fatalf("expected bitmap to shrink back to 64 bits, got %d", bm.Length())
	}

	if bm.IsSet(1) {
		t.Fatal("expected bit 1 to be cleared when secondary bitmap is empty")
	}

	if bm.IsSet(70) {
		t.Fatal("expected bit 70 to be cleared")
	}
}

func TestAsciiBitmapIsSetOutOfRange(t *testing.T) {
	bm := NewAsciiBitmap()

	if bm.IsSet(0) {
		t.Fatal("expected IsSet to return false for out-of-range position")
	}

	if bm.IsSet(129) {
		t.Fatal("expected IsSet to return false for out-of-range position")
	}
}

func TestNewEmptyBitmap(t *testing.T) {
	bm := NewEmptyAsciiBitmap()

	if bm.Length() != 64 {
		t.Fatalf("expected 64-bit bitmap, got %d", bm.Length())
	}

	if bm.IsSet(1) {
		t.Fatal("expected all bits to be cleared in empty bitmap")
	}
}

func TestAsciiBitmapBitmapList(t *testing.T) {
	bm := NewEmptyAsciiBitmap()

	must := func(err error) {
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}

	must(bm.Set(1))
	must(bm.Set(3))
	must(bm.Set(64))
	must(bm.Set(70))

	want := []int{1, 3, 64, 70}
	if got := bm.Bitmap(); !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected bitmap list: got %v want %v", got, want)
	}
}
