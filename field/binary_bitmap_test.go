package field

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func bitsFromHexBinary(t *testing.T, hexStr string) string {
	t.Helper()

	var bits strings.Builder
	bits.Grow(len(hexStr) * 4)

	for pos, c := range strings.ToUpper(hexStr) {
		n, err := strconv.ParseUint(string(c), 16, 4)
		if err != nil {
			t.Fatalf("invalid hex %q at position %d: %v", c, pos, err)
		}
		bits.WriteString(fmt.Sprintf("%04b", n))
	}

	return bits.String()
}

func hexToBytes(t *testing.T, hexStr string) []byte {
	t.Helper()
	data, err := hex.DecodeString(hexStr)
	if err != nil {
		t.Fatalf("failed to decode hex %q: %v", hexStr, err)
	}
	return data
}

func TestBinaryBitmapEncodePrimary(t *testing.T) {
	hexStr := "0123456789ABCDEF"
	bits := bitsFromHexBinary(t, hexStr)

	encoded, err := NewBinaryBitmap(bits).Encode()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := hexToBytes(t, hexStr)
	if !reflect.DeepEqual(encoded, want) {
		t.Fatalf("unexpected encoded bytes: got % X want % X", encoded, want)
	}
}

func TestBinaryBitmapEncodeSecondary(t *testing.T) {
	hexStr := "F123456789ABCDEF0123456789ABCDEF"
	bits := bitsFromHexBinary(t, hexStr)

	encoded, err := NewBinaryBitmap(bits).Encode()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := hexToBytes(t, hexStr)
	if !reflect.DeepEqual(encoded, want) {
		t.Fatalf("unexpected encoded bytes: got % X want % X", encoded, want)
	}
}

func TestBinaryBitmapEncodeRejectsSecondaryWithoutBits(t *testing.T) {
	bitmap := bitsFromHexBinary(t, "8000000000000000")

	if _, err := NewBinaryBitmap(bitmap).Encode(); err == nil {
		t.Fatal("expected error for secondary bitmap flag without 128 bits, got nil")
	}
}

func TestBinaryBitmapDecodePrimary(t *testing.T) {
	hexStr := "0123456789ABCDEF"
	data := hexToBytes(t, hexStr)
	bits := bitsFromHexBinary(t, hexStr)

	bm := NewBinaryBitmap("")
	if err := bm.Decode(data); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if bm.Value() != bits {
		t.Fatalf("decoded bits mismatch: got %q want %q", bm.Value(), bits)
	}
}

func TestBinaryBitmapDecodeSecondary(t *testing.T) {
	hexStr := "F123456789ABCDEF0123456789ABCDEF"
	data := hexToBytes(t, hexStr)
	bits := bitsFromHexBinary(t, hexStr)

	bm := NewBinaryBitmap("")
	if err := bm.Decode(data); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if bm.Value() != bits {
		t.Fatalf("decoded bits mismatch: got %q want %q", bm.Value(), bits)
	}
}

func TestBinaryBitmapDecodeMissingSecondary(t *testing.T) {
	data := hexToBytes(t, "8000000000000000")
	if err := NewBinaryBitmap("").Decode(data); err == nil {
		t.Fatal("expected error when secondary bitmap flag set but only primary provided")
	}
}

func TestBinaryBitmapSetClearAndBitmapList(t *testing.T) {
	bm := NewEmptyBinaryBitmap()

	must := func(err error) {
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}

	must(bm.Set(2))
	must(bm.Set(65))

	if bm.Length() != 128 {
		t.Fatalf("expected expansion to 128 bits, got %d", bm.Length())
	}

	if !bm.IsSet(1) || !bm.IsSet(2) || !bm.IsSet(65) {
		t.Fatalf("expected bits 1, 2, 65 to be set, got bitmap %v", bm.Bitmap())
	}

	must(bm.Clear(65))

	if bm.Length() != 64 {
		t.Fatalf("expected shrink back to 64 bits, got %d", bm.Length())
	}

	want := []int{2}
	if got := bm.Bitmap(); !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected bitmap list: got %v want %v", got, want)
	}
}
