package padder

// No padding.
var None *noPadding = &noPadding{}

type noPadding struct{}

func (p *noPadding) Pad(s string, maxLen int) string {
	return s
}

func (p *noPadding) Unpad(s string) string {
	return s
}
