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
	return Field{ID: id, value: value}
}

// Field is used to represent a field value in a message.
type Field struct {
	ID          int
	Description string

	value    string
	dataType codec.DataType
}

func (f *Field) WithID(id int) *Field {
	f.ID = id
	return f
}

func (f *Field) WithValue(val string) *Field {
	f.value = val
	return f
}

func (f *Field) WithDescription(description string) *Field {
	f.Description = description
	return f
}

func (f *Field) WithType(dataType codec.DataType) *Field {
	f.dataType = dataType
	return f
}

// ValueAsNumber tries to convert the value to an integer value.
func (f *Field) ValueAsNumber() (int64, error) {
	if f.value == "" {
		return 0, nil
	}
	return strconv.ParseInt(f.value, 10, 64)
}

// Tries to create a decimal value from the value of the field. This is
// useful when dealing with amount fields.
//
// This uses library github.com/shopspring/decimal which is an unofåficial
// standard for dealing with amount data types.
func (f *Field) ValueAsDecimal(decimalPlaces int) (decimal.Decimal, error) {
	if f.value == "" {
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
	return f.value
}
