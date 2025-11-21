package siso

import (
	"github.com/rwbm/siso/field"
)

type IsoMessage struct {
	Header string
	Length int
	Fields map[int]field.IsoField
}
