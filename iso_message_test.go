package siso

import "testing"

func TestIsoMessageXml(t *testing.T) {
	msg := IsoMessage{Header: NewField(-1, "ISO")}
	msg.
		WithField(Field{ID: 0, Description: "Message Type Indicator", Value: "0200"}).
		WithField(Field{ID: 2, Description: "Primary Account Number", Value: "5185660000006004"}).
		WithField(Field{ID: 3, Description: "Processing Code", Value: "12345"}).
		WithField(Field{ID: 4, Description: "Transaction Amount", Value: "10000"})

	t.Log("\n", msg.StringXml())
}
