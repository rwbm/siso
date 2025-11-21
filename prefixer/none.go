package prefixer

// No prefixer.
var None *noPrefierxer = &noPrefierxer{}

type noPrefierxer struct{}

func (n *noPrefierxer) Encode(length int, data []byte) error {
	return nil
}

func (n *noPrefierxer) Decode(data []byte, offset int) (int, error) {
	return -1, nil
}

func (n *noPrefierxer) PackedLen() int {
	return 0
}
