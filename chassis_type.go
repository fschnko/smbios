package smbios

const (
	// Bit 7 Chassis lock is present if 1.
	_ChassisTypeLockMask = 0b_1000_0000
	// Bits 6:0 Enumeration value; see below.
	_ChassisTypeValueMask = _ChassisTypeLockMask ^ 0xff // invert chassis lock mask
)

var _ChassisType = _EnumDictionary{
	"Not set", // dictionary starts from 0x01
	"Other",
	"Unknown",
	"Desktop",
	"Low Profile Desktop",
	"Pizza Box",
	"Mini Tower",
	"Tower",
	"Portable",
	"Laptop",
	"Notebook",
	"Hand Held",
	"Docking Station",
	"All in One",
	"Sub Notebook",
	"Space-saving",
	"Lunch Box",
	"Main Server Chassis",
	"Expansion Chassis",
	"SubChassis",
	"Bus Expansion Chassis",
	"Peripheral Chassis",
	"RAID Chassis",
	"Rack Mount Chassis",
	"Sealed-case PC",
	"Multi-system chassis",
	"Compact PCI",
	"Advanced TCA",
	"Blade",
	"Tablet",
	"Convertible",
	"Detachable",
	"IoT Gateway",
	"Embedded PC",
	"Mini PC",
	"Stick PC",
}

// ChassisType represents the System Enclosure or Chassis Types field.
type ChassisType byte

func (ct ChassisType) String() string {
	return _ChassisType.str(ct.Number())
}

// Number returns number of the chassis type.
func (ct ChassisType) Number() int {
	return int(ct & _ChassisTypeValueMask)
}

// ChassisLock shows the presence of the chassis lock.
type ChassisLock byte

func (cl ChassisLock) String() string {
	if cl.Present() {
		return "Present"
	}
	return "Not Present"
}

// Present returns true if chassis lock is present.
func (cl ChassisLock) Present() bool {
	return cl&_ChassisTypeLockMask != 0

}
