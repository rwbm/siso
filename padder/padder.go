package padder

type Padder interface {
	Pad(s string, maxLen int) string
	Unpad(s string) string
}
