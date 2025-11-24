package codec

import "testing"

func TestAsciiNumericEncode(t *testing.T) {
	an := NewAsciiNumeric()
	encoded, err := an.Encode("12345")
	if err != nil {
		t.Fatalf("Encode returned error: %v", err)
	}
	if string(encoded) != "12345" {
		t.Fatalf("Encode = %q, want %q", string(encoded), "12345")
	}
}

func TestAsciiNumericEncodeRejectsNonNumeric(t *testing.T) {
	an := NewAsciiNumeric()
	if _, err := an.Encode("12A"); err == nil {
		t.Fatalf("Encode succeeded, expected error for non-numeric input")
	}
}

func TestAsciiNumericDecode(t *testing.T) {
	an := NewAsciiNumeric()
	val, err := an.Decode([]byte("000789"))
	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	if val != "789" {
		t.Fatalf("Decode value = %q, want %q", val, "789")
	}
}

func TestAsciiNumericDecodeRejectsInvalid(t *testing.T) {
	an := NewAsciiNumeric()
	if _, err := an.Decode([]byte("12A")); err == nil {
		t.Fatalf("Decode succeeded, expected error for non-numeric data")
	}
}
