package padder

import (
	"fmt"
	"strings"
)

// Pads to the left with zeros.
var LeftZero *leftPadding = &leftPadding{"0"}

// Pads to the left with spaces.
var LeftSpaces *leftPadding = &leftPadding{" "}

type leftPadding struct {
	filler string
}

func (p *leftPadding) Pad(s string, maxLen int) string {
	if len(s) > maxLen {
		return s[:maxLen]
	}
	if len(s) == maxLen {
		return s
	}
	return fmt.Sprintf("%s%s", strings.Repeat(p.filler, maxLen-len(s)), s)
}

func (p *leftPadding) Unpad(s string) string {
	if len(p.filler) == 0 {
		return s
	}

	for strings.HasPrefix(s, p.filler) {
		s = s[len(p.filler):]
	}
	return s
}
