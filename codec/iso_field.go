package codec

// IsoField represents the interface definition for field implementations.
type IsoField interface {
	Length() int
	Encode(value string) ([]byte, error)
	Decode([]byte) (string, error)
}
