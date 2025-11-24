package codec

// Codec represents the interface definition for field encoding/decoding.
type Codec interface {
	Encode(value string) ([]byte, error)
	Decode([]byte) (string, error)
}
