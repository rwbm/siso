package encoder

type asciiEncoder struct{}

func (e asciiEncoder) Encode(data []byte) ([]byte, error) {
	return nil, nil
}

func (e asciiEncoder) Decode(data []byte, length int) ([]byte, int, error) {
	return nil, 0, nil
}
