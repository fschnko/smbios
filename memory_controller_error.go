package smbios

var _MemoryControllerErrCorrecting = _FlagDictionary64{
	"Other",
	"Unknown",
	"None",
	"Single-Bit Error Correcting",
	"Double-Bit Error Correcting",
	"Error Scrubbing",
}

var _MemoryControllerErrDecting = _EnumDictionary{
	"Not set", // dictionary starts from 0x01
	"Other",
	"Unknown",
	"None",
	"8-bit Parity",
	"32-bit ECC",
	"64-bit ECC",
	"128-bit ECC",
	"CRC",
}

// MemoryControllerErrCorrecting represents a memory controller's
// error correcting capability.
type MemoryControllerErrCorrecting byte

func (c MemoryControllerErrCorrecting) String() string {
	return _MemoryControllerErrDecting.str(int(c))
}

// MemoryControllerErrDecting represents a memory controller's
// error detecting method field.
type MemoryControllerErrDecting byte

func (d MemoryControllerErrDecting) String() string {
	return _MemoryControllerErrDecting.str(int(d))
}
