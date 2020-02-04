package smbios

var _ChassisState = _EnumDictionary{
	"Not set", // dictionary starts from 0x01
	"Other",
	"Unknown",
	"Safe",
	"Warning",
	"Critical",
	"Non-recoverable",
}

// ChassisState represents a chassis State field.
type ChassisState byte

func (s ChassisState) String() string {
	return _ChassisState.str(int(s))
}
