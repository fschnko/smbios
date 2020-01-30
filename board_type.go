package smbios

var _BaseboardType = _EnumDictionary{
	"Not set", // dictionary starts from 0x01
	"Unknown",
	"Other",
	"Server Blade",
	"Connectivity Switch",
	"System Management Module",
	"Processor Module",
	"I/O Module",
	"Memory Module",
	"Daughter board",
	"Motherboard",
	"Processor/Memory Module",
	"Processor/IO Module",
	"Interconnect board",
}

// BaseboardType represents a baseboard type field.
type BaseboardType byte

func (bt BaseboardType) String() string {
	return _BaseboardType.str(int(bt))
}
