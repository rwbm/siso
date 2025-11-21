package siso

import (
	"fmt"
	"strconv"

	"github.com/rwbm/siso/codec"
	"github.com/shopspring/decimal"
)

// Fields with especial meanings.
const (
	FieldMessageTypeIndicator = 0
	FieldPrimaryBitmap        = 1
	FieldSecondaryBitmap      = 128
)

func NewField(id int, value string) Field {
	return Field{ID: id, Value: value}
}

// Field is used to represent a field value in a message.
type Field struct {
	ID          int
	Description string
	Value       string
	Type        codec.DataType
}

func (f *Field) ValueAsNumber() (int64, error) {
	if f.Value == "" {
		return 0, nil
	}
	return strconv.ParseInt(f.Value, 10, 64)
}

func (f *Field) ValueAsDecimal(decimalPlaces int) (decimal.Decimal, error) {
	if f.Value == "" {
		return decimal.Zero, nil
	}

	if decimalPlaces < 0 {
		return decimal.Zero, fmt.Errorf("decimalPlaces must be non-negative")
	}

	numVal, err := f.ValueAsNumber()
	if err != nil {
		return decimal.Zero, err
	}

	return decimal.NewFromInt(numVal).Shift(int32(-decimalPlaces)), nil
}

func (f *Field) String() string {
	return f.Value
}
