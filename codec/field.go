package codec

// Field represents the interface definition for field implementations.
type Field interface {
	Length() int
	Encode(value string) ([]byte, error)
	Decode([]byte) (string, error)
}
