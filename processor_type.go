package smbios

var _ProcessorType = _EnumDictionary{
	"Not set", // dictionary starts from 0x01
	"Other",
	"Unknown",
	"Central Processor",
	"Math Processor",
	"DSP Processor",
	"Video Processor",
}

// ProcessorType represents a processor type field.
type ProcessorType byte

func (pt ProcessorType) String() string {
	return _ProcessorType.str(int(pt))
}
