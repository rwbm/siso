package common

// DataType represents the content encoding and length semantics of a field.
type DataType string

const (
	// ASCII
	AsciiNumeric     DataType = "ASCII_NUMERIC"
	AsciiLLNumeric   DataType = "ASCII_LLNUM"
	AsciiLLLNumeric  DataType = "ASCII_LLLNUM"
	AsciiLLLLNumeric DataType = "ASCII_LLLLNUM"
	AsciiChar        DataType = "ASCII_CHAR"
	AsciiLLChar      DataType = "ASCII_LLCHAR"
	AsciiLLLChar     DataType = "ASCII_LLLCHAR"
	AsciiLLLLChar    DataType = "ASCII_LLLLCHAR"
	AsciiBinary      DataType = "ASCII_BINARY"
	AsciiLLBinary    DataType = "ASCII_LLBINARY"
	AsciiLLLBinary   DataType = "ASCII_LLLBINARY"
	AsciiLLLLBinary  DataType = "ASCII_LLLLBINARY"

	// EBCDIC
	EbcdicNumeric     DataType = "EBCDIC_NUMERIC"
	EbcdicLLNumeric   DataType = "EBCDIC_LLNUM"
	EbcdicLLLNumeric  DataType = "EBCDIC_LLLNUM"
	EbcdicLLLLNumeric DataType = "EBCDIC_LLLLNUM"
	EbcdicChar        DataType = "EBCDIC_CHAR"
	EbcdicLLChar      DataType = "EBCDIC_LLCHAR"
	EbcdicLLLChar     DataType = "EBCDIC_LLLCHAR"
	EbcdicLLLLChar    DataType = "EBCDIC_LLLLCHAR"
	EbcdicBinary      DataType = "EBCDIC_BINARY"
	EbcdicLLBinary    DataType = "EBCDIC_LLBINARY"
	EbcdicLLLBinary   DataType = "EBCDIC_LLLBINARY"
	EbcdicLLLLBinary  DataType = "EBCDIC_LLLLBINARY"

	// Binary (length digits/contents are binary encoded)
	BinaryNumeric     DataType = "BINARY_NUMERIC"
	BinaryLLNumeric   DataType = "BINARY_LLNUM"
	BinaryLLLNumeric  DataType = "BINARY_LLLNUM"
	BinaryLLLLNumeric DataType = "BINARY_LLLLNUM"
	BinaryChar        DataType = "BINARY_CHAR"
	BinaryLLChar      DataType = "BINARY_LLCHAR"
	BinaryLLLChar     DataType = "BINARY_LLLCHAR"
	BinaryLLLLChar    DataType = "BINARY_LLLLCHAR"
	BinaryBinary      DataType = "BINARY_BINARY"
	BinaryLLBinary    DataType = "BINARY_LLBINARY"
	BinaryLLLBinary   DataType = "BINARY_LLLBINARY"
	BinaryLLLLBinary  DataType = "BINARY_LLLLBINARY"
)
