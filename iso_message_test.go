package siso

import (
	"testing"
)

func TestIsoMessageXml(t *testing.T) {
	msg := IsoMessage{Header: NewField(-1, "ISO")}
	msg.
		WithField(Field{ID: 0, Description: "Message Type Indicator", value: "0200"}).
		WithField(Field{ID: 2, Description: "Primary Account Number", value: "5185660000006004"}).
		WithField(Field{ID: 3, Description: "Processing Code", value: "12345"}).
		WithField(Field{ID: 4, Description: "Transaction Amount", value: "10000"})

	t.Log("\n", msg.StringXml())
}
