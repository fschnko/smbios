package smbios

var _CacheType = _EnumDictionary{
	"Not set", // dictionary starts from 0x01
	"Other",
	"Unknown",
	"Instruction",
	"Data",
	"Unified",
}

// CacheType represents System Cache Type field.
type CacheType byte

func (c CacheType) String() string {
	return _CacheType.str(int(c))
}
