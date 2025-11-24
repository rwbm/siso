package codec

// DataType represents the content encoding and length semantics of a field.
type DataType string

const (
	TypeUndefined DataType = ""

	// ASCII
	TypeAsciiNumeric     DataType = "ASCII_NUMERIC"
	TypeAsciiLLNumeric   DataType = "ASCII_LLNUM"
	TypeAsciiLLLNumeric  DataType = "ASCII_LLLNUM"
	TypeAsciiLLLLNumeric DataType = "ASCII_LLLLNUM"
	TypeAsciiChar        DataType = "ASCII_CHAR"
	TypeAsciiLLChar      DataType = "ASCII_LLCHAR"
	TypeAsciiLLLChar     DataType = "ASCII_LLLCHAR"
	TypeAsciiLLLLChar    DataType = "ASCII_LLLLCHAR"
	TypeAsciiBinary      DataType = "ASCII_BINARY"
	TypeAsciiLLBinary    DataType = "ASCII_LLBINARY"
	TypeAsciiLLLBinary   DataType = "ASCII_LLLBINARY"
	TypeAsciiLLLLBinary  DataType = "ASCII_LLLLBINARY"

	// EBCDIC
	TypeEbcdicNumeric     DataType = "EBCDIC_NUMERIC"
	TypeEbcdicLLNumeric   DataType = "EBCDIC_LLNUM"
	TypeEbcdicLLLNumeric  DataType = "EBCDIC_LLLNUM"
	TypeEbcdicLLLLNumeric DataType = "EBCDIC_LLLLNUM"
	TypeEbcdicChar        DataType = "EBCDIC_CHAR"
	TypeEbcdicLLChar      DataType = "EBCDIC_LLCHAR"
	TypeEbcdicLLLChar     DataType = "EBCDIC_LLLCHAR"
	TypeEbcdicLLLLChar    DataType = "EBCDIC_LLLLCHAR"
	TypeEbcdicBinary      DataType = "EBCDIC_BINARY"
	TypeEbcdicLLBinary    DataType = "EBCDIC_LLBINARY"
	TypeEbcdicLLLBinary   DataType = "EBCDIC_LLLBINARY"
	TypeEbcdicLLLLBinary  DataType = "EBCDIC_LLLLBINARY"

	// Binary (length digits/contents are binary encoded)
	TypeBinaryNumeric     DataType = "BINARY_NUMERIC"
	TypeBinaryLLNumeric   DataType = "BINARY_LLNUM"
	TypeBinaryLLLNumeric  DataType = "BINARY_LLLNUM"
	TypeBinaryLLLLNumeric DataType = "BINARY_LLLLNUM"
	TypeBinaryChar        DataType = "BINARY_CHAR"
	TypeBinaryLLChar      DataType = "BINARY_LLCHAR"
	TypeBinaryLLLChar     DataType = "BINARY_LLLCHAR"
	TypeBinaryLLLLChar    DataType = "BINARY_LLLLCHAR"
	TypeBinaryBinary      DataType = "BINARY_BINARY"
	TypeBinaryLLBinary    DataType = "BINARY_LLBINARY"
	TypeBinaryLLLBinary   DataType = "BINARY_LLLBINARY"
	TypeBinaryLLLLBinary  DataType = "BINARY_LLLLBINARY"
)
