package codec

// IsoBitmap extends IsoField with helpers to toggle and inspect bitmap bits.
type IsoBitmap interface {
	IsoField
	IsSet(pos int) bool
	Set(pos int) error
	Clear(pos int) error
	String() string
	Bitmap() []int
}
