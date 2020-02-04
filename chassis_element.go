package smbios

const (
	_ChassisElementTypeIdentifierMask = 0b_1000_0000
	_ChassisElementTypeDataMask       = _ChassisElementTypeIdentifierMask ^ 0xff
)

// ChassisElement represents a Chassis Contained Element.
type ChassisElement []byte

// Type returns the type of element associated with a record.
func (el *ChassisElement) Type() ChassisElementType {
	return ChassisElementType(el.byte(0x00))
}

// MinCount returns the minimum number of the element type
// that can be installed in the chassis for the chassis to
// properly operate.
func (el *ChassisElement) MinCount() byte {
	return el.byte(0x01)
}

// MaxCount returns the maximum number of the element type
// that can be installed in the chassis.
func (el *ChassisElement) MaxCount() byte {
	return el.byte(0x02)
}

func (el *ChassisElement) byte(offset int) byte {
	if el != nil && len(*el) > offset {
		return (*el)[offset]
	}
	return 0
}

// ChassisElementType represents a chassis contained element type.
type ChassisElementType byte

// IsBaseboard returns true if contains an SMBIOS baseboard type enumeration.
func (et ChassisElementType) IsBaseboard() bool {
	return et&_ChassisElementTypeIdentifierMask == 0
}

// IsTable returns true if contains an SMBIOS table type enumeration.
func (et ChassisElementType) IsTable() bool {
	return et&_ChassisElementTypeIdentifierMask > 0
}

// Value returns either an SMBIOS Board Type enumeration
// or an SMBIOS table type, dependent on the setting of the
// Type Select.
func (et ChassisElementType) Value() byte {
	return byte(et & _ChassisElementTypeDataMask)
}

func (et ChassisElementType) String() string {
	if et.IsBaseboard() {
		return "board: " + BaseboardType(et.Value()).String()
	}

	if et.IsTable() {
		return "table: " + TableType(et.Value()).String()
	}

	return "Unknown"
}
