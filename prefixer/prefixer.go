package prefixer

type Prefixer interface {
	Encode(length int, data []byte) error
	Decode(data []byte, offset int) (int, error)
	PackedLen() int
}
