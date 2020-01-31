package smbios

var _CacheSRAMType = _FlagDictionary64{
	"Other",
	"Unknown",
	"Non-Burst",
	"Burst",
	"Pipeline Burst",
	"Synchronous",
	"Asynchronous",
	// Bits 7:15 Reserved, must be zero
}

// CacheSRAM represents cache SRAM type.
type CacheSRAM uint16

func (c CacheSRAM) String() string {
	return _CacheSRAMType.str(uint64(c))
}
