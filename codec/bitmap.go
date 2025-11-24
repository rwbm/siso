package codec

// Bitmap extends IsoField with helpers to toggle and inspect bitmap bits.
type Bitmap interface {
	Field
	IsSet(pos int) bool
	Set(pos int) error
	Clear(pos int) error
	String() string
	Bitmap() []int
}
