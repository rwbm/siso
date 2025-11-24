package codec

import "testing"

func TestAsciiCharEncode(t *testing.T) {
	ac := NewAsciiChar()
	encoded, err := ac.Encode("ABC")
	if err != nil {
		t.Fatalf("Encode returned error: %v", err)
	}
	if string(encoded) != "ABC" {
		t.Fatalf("Encode = %q, want %q", string(encoded), "ABC")
	}
}

func TestAsciiCharEncodeRejectsNonASCII(t *testing.T) {
	ac := NewAsciiChar()
	if _, err := ac.Encode("olé"); err == nil {
		t.Fatalf("expected error for non-ASCII input")
	}
}

func TestAsciiCharDecode(t *testing.T) {
	ac := NewAsciiChar()
	val, err := ac.Decode([]byte("XYZ"))
	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	if val != "XYZ" {
		t.Fatalf("Decode value = %q, want %q", val, "XYZ")
	}
}

func TestAsciiCharDecodeRejectsNonASCII(t *testing.T) {
	ac := NewAsciiChar()
	if _, err := ac.Decode([]byte{0xFF, 0x41}); err == nil {
		t.Fatalf("expected error for non-ASCII byte")
	}
}
