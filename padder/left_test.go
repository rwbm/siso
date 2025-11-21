package padder

import "testing"

func TestLeftPaddingPad(t *testing.T) {
	t.Run("zero filler pads on the left", func(t *testing.T) {
		got := LeftZero.Pad("123", 5)
		if got != "00123" {
			t.Fatalf("Pad returned %q, want %q", got, "00123")
		}
	})

	t.Run("space filler pads on the left", func(t *testing.T) {
		got := LeftSpaces.Pad("A", 3)
		if got != "  A" {
			t.Fatalf("Pad returned %q, want %q", got, "  A")
		}
	})

	t.Run("truncates when input exceeds max length", func(t *testing.T) {
		got := LeftZero.Pad("abcdef", 4)
		if got != "abcd" {
			t.Fatalf("Pad returned %q, want %q", got, "abcd")
		}
	})
}

func TestLeftPaddingUnpad(t *testing.T) {
	t.Run("removes leading zero filler", func(t *testing.T) {
		got := LeftZero.Unpad("000123")
		if got != "123" {
			t.Fatalf("Unpad returned %q, want %q", got, "123")
		}
	})

	t.Run("removes leading space filler", func(t *testing.T) {
		got := LeftSpaces.Unpad("   ABC")
		if got != "ABC" {
			t.Fatalf("Unpad returned %q, want %q", got, "ABC")
		}
	})

	t.Run("does not strip non-filler prefix", func(t *testing.T) {
		got := LeftZero.Unpad("A001")
		if got != "A001" {
			t.Fatalf("Unpad returned %q, want %q", got, "A001")
		}
	})

	t.Run("returns empty when only filler present", func(t *testing.T) {
		got := LeftZero.Unpad("0000")
		if got != "" {
			t.Fatalf("Unpad returned %q, want empty string", got)
		}
	})
}
