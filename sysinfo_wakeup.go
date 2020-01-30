package smbios

var _SystemWakeUp = _EnumDictionary{
	"Reserved",
	"Other",
	"Unknown",
	"APM Timer",
	"Modem Ring",
	"LAN Remote",
	"Power Switch",
	"PCI PME#",
	"AC Power Restored",
}

// SystemWakeUp represents a system wake up type field.
type SystemWakeUp byte

func (s SystemWakeUp) String() string {
	return _SystemWakeUp.str(int(s))
}
