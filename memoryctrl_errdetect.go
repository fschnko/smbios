package smbios

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

// MemoryControllerErrDecting represents a memory controller's
// error detecting method field.
type MemoryControllerErrDecting byte

func (d MemoryControllerErrDecting) String() string {
	return _MemoryControllerErrDecting.str(int(d))
}
