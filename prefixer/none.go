package prefixer

// No prefixer.
var None *noPrefixer = &noPrefixer{}

type noPrefixer struct{}

func (n *noPrefixer) Encode(length int, data []byte) error {
	return nil
}

func (n *noPrefixer) Decode(data []byte, offset int) (int, error) {
	return -1, nil
}

func (n *noPrefixer) PackedLen() int {
	return 0
}
