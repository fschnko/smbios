package smbios

import "fmt"

// Memory size units according to the International Electrotechnical Commission.
// Numeric prefixes (kibi-, mebi-, gibi-, and so on, abbreviated Ki, Mi, and Gi)
// based explicitly on powers of 2.
const (
	Byte MemorySize = 1 << (10 * iota)
	KibiByte
	MebiByte
	GibiByte
	TebiByte
	PebiByte
	ExbiByte

	// Granularity units
	KibiByte64 MemorySize = 64 * KibiByte
)

// MemorySize represents a memory size in bytes.
type MemorySize uint64

// As returns value in given size unit.
func (s MemorySize) As(gran MemorySize) uint64 {
	return uint64(s / gran)
}

// ConvertableTo returns true if value can be converted to given size unit.
func (s MemorySize) ConvertableTo(gran MemorySize) bool {
	return gran != 0 && s%gran == 0
}

func (s MemorySize) String() string {
	var (
		gran    = Byte
		symbols = "B"
	)

	switch {
	case s == 0:
		gran, symbols = Byte, "B"
	case s.ConvertableTo(ExbiByte):
		gran, symbols = ExbiByte, "EiB"
	case s.ConvertableTo(PebiByte):
		gran, symbols = PebiByte, "PiB"
	case s.ConvertableTo(TebiByte):
		gran, symbols = TebiByte, "TiB"
	case s.ConvertableTo(GibiByte):
		gran, symbols = GibiByte, "GiB"
	case s.ConvertableTo(MebiByte):
		gran, symbols = MebiByte, "MiB"
	case s.ConvertableTo(KibiByte):
		gran, symbols = KibiByte, "KiB"
	}

	return fmt.Sprintf("%d%s", s.As(gran), symbols)
}
