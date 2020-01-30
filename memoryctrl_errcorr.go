package smbios

var _MemoryControllerErrCorrecting = _FlagDictionary64{
	"Other",
	"Unknown",
	"None",
	"Single-Bit Error Correcting",
	"Double-Bit Error Correcting",
	"Error Scrubbing",
}

// MemoryControllerErrCorrecting represents a memory controller's
// error correcting capability.
type MemoryControllerErrCorrecting byte

func (c MemoryControllerErrCorrecting) String() string {
	return _MemoryControllerErrDecting.str(int(c))
}
