package parser

import "github.com/rwbm/siso/padder"

// Parser turns raw field bytes into a string value and returns any remaining bytes
// that have not been consumed. Implementations handle their own encoding rules
// (ASCII, EBCDIC, binary, ...).
type Parser interface {
	Parse(data []byte, length int, pad padder.Padder) (string, []byte, error)
}
