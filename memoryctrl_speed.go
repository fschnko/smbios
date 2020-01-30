package smbios

var _MemoryControllerSpeed = _FlagDictionary64{
	"Other",
	"Unknown",
	"70ns",
	"60ns",
	"50ns",
	// Bits 5:15 Reserved, must be zero
}

// MemoryControllerSpeed represents a memory controller supported speed information.
type MemoryControllerSpeed uint16

func (s MemoryControllerSpeed) String() string {
	return _MemoryControllerSpeed.str(uint64(s))
}
