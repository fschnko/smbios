package smbios

var _MemoryControllerInterleave = _EnumDictionary{
	"Not set", // dictionary starts from 0x01
	"Other",
	"Unknown",
	"One-Way Interleave",
	"Two-Way Interleave",
	"Four-Way Interleave",
	"Eight-Way Interleave",
	"Sixteen-Way Interleave",
}

// MemoryControllerInterleave represents a memory controller's interleave information.
type MemoryControllerInterleave byte

func (i MemoryControllerInterleave) String() string {
	return _MemoryControllerInterleave.str(int(i))
}
