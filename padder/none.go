package padder

// No padding.
var None *noPadding = &noPadding{}

type noPadding struct{}

func (n *noPadding) Pad(s string, maxLen int) string {
	return s
}

func (n *noPadding) Unpad(s string) string {
	return s
}
