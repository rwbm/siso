package field

import (
	"github.com/rwbm/siso/padder"
	"github.com/rwbm/siso/prefixer"
)

// IsoField represents the interface definition for field implementations.
type IsoField interface {
	Value() any
	Length() int
	Prefixer() prefixer.Prefixer
	Padder() padder.Padder
	Encode() ([]byte, error)
	Decode([]byte) error
}
