package smbios

var _ChassisSecurityStatus = _EnumDictionary{
	"Not set", // dictionary starts from 0x01
	"Other",
	"Unknown",
	"None",
	"External interface locked out",
	"External interface enabled",
}

// ChassisSecurityStatus represents a chassis Security Status field.
type ChassisSecurityStatus byte

func (s ChassisSecurityStatus) String() string {
	return _ChassisSecurityStatus.str(int(s))
}
