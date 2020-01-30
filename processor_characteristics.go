package smbios

const _ProcessorCharacteristicsMask = 0x0e // Skip Reserved fields

var _ProcessorCharacteristics = _FlagDictionary64{
	"Reserved",
	"Unknown",
	"64-bit Capable",
	"Multi-Core",
	"Hardware Thread",
	"Execute Protection",
	"Enhanced Virtualization",
	"Power/Performance Control",
	"128-bit Capable",
	//Bits 9:15 Reserved
}

// ProcessorCharacteristics represents a processor characteristics field.
type ProcessorCharacteristics uint16

func (c ProcessorCharacteristics) String() string {
	return _ProcessorCharacteristics.str(uint64(c & _ProcessorCharacteristicsMask))
}
