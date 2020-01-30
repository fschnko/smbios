package smbios

var _MemoryType = _FlagDictionary64{
	"Other",
	"Unknown",
	"Standard",
	"Fast Page Mode",
	"EDO",
	"Parity",
	"ECC",
	"SIMM",
	"DIMM",
	"Burst EDO",
	"Bit 10 SDRAM",
	// Bits 11:15 Reserved, must be zero
}

// MemoryType describes the physical characteristics of the memory modules
// that are supported by (and currently installed in) the system.
type MemoryType uint16

func (t MemoryType) String() string {
	return _MemoryType.str(uint64(t))
}
