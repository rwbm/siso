package siso

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestValueAsNumber(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  int64
	}{
		{name: "empty", value: "", want: 0},
		{name: "positive", value: "12345", want: 12345},
		{name: "negative", value: "-987", want: -987},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			field := Field{value: tc.value}
			got, err := field.ValueAsNumber()
			if err != nil {
				t.Fatalf("ValueAsNumber returned error: %v", err)
			}
			if got != tc.want {
				t.Fatalf("ValueAsNumber = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestValueAsDecimal(t *testing.T) {
	tests := []struct {
		name          string
		value         string
		decimalPlaces int
		want          string
	}{
		{name: "empty", value: "", decimalPlaces: 2, want: "0"},
		{name: "integer no decimals", value: "123", decimalPlaces: 0, want: "123"},
		{name: "with decimals", value: "12345", decimalPlaces: 2, want: "123.45"},
		{name: "small fractional", value: "1", decimalPlaces: 3, want: "0.001"},
		{name: "negative", value: "-9876", decimalPlaces: 2, want: "-98.76"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			field := Field{value: tc.value}
			got, err := field.ValueAsDecimal(tc.decimalPlaces)
			if err != nil {
				t.Fatalf("ValueAsDecimal returned error: %v", err)
			}
			if got.Cmp(decimal.RequireFromString(tc.want)) != 0 {
				t.Fatalf("ValueAsDecimal = %s, want %s", got.String(), tc.want)
			}
		})
	}
}
