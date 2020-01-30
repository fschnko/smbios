package smbios

type Chassis struct {
	table
}

// Manufacturer returns a chassis manufacturer name.
func (c *Chassis) Manufacturer() string {
	return c.str(0x04)
}

// Type returns a chassis type.
func (c *Chassis) Type() ChassisType {
	return ChassisType(c.byte(0x05))
}

// Lock returns a chassis lock information.
func (c *Chassis) Lock() ChassisLock {
	return ChassisLock(c.byte(0x05))
}

// Version returns a chassis version.
func (c *Chassis) Version() string {
	return c.str(0x06)
}

// SerialNumber returns a chassis serial number.
func (c *Chassis) SerialNumber() string {
	return c.str(0x07)
}

// AssetTag returns a chassis asset tag.
func (c *Chassis) AssetTag() string {
	return c.str(0x08)
}

// BootState returns state of the chassis when it was last booted.
func (c *Chassis) BootState() ChassisState {
	return ChassisState(c.byte(0x09))
}

// PowerSupplyState returns power supply state of the chassis when last booted.
func (c *Chassis) PowerSupplyState() ChassisState {
	return ChassisState(c.byte(0x0a))
}

// ThermalState returns thermal state of the chassis when last booted.
func (c *Chassis) ThermalState() ChassisState {
	return ChassisState(c.byte(0x0b))
}

// SecurityStatus returns physical security status of the chassis when
// last booted.
func (c *Chassis) SecurityStatus() ChassisSecurityStatus {
	return ChassisSecurityStatus(c.byte(0x0c))
}

// OEMInformation returns OEM or BIOS vendor-specific information.
func (c *Chassis) OEMInformation() uint32 {
	return c.dword(0x0d)
}

// Height returns height of the chassis, in 'U's.
// A U is a standard unit of measure for the
// height of a rack or rack-mountable component
// and is equal to 1.75 inches or 4.445 cm.
// A value of 00h indicates that the chassis
// height is unspecified.
func (c *Chassis) Height() byte {
	return c.byte(0x11)
}

// NumberPowerCords returns number of power cords
// associated with the chassis. A value of 00h indicates
// that the number is unspecified.
func (c *Chassis) NumberPowerCords() byte {
	return c.byte(0x12)
}

// 13h 2.3+ ContainedElementCount (n)
// BYTE Varies Number of Contained Element records that
// follow, in the range 0 to 255
// Each Contained Element group comprises m
// bytes, as specified by the Contained Element
// Record Length field that follows. If no
// Contained Elements are included, this field is
// set to 0.

// 14h 2.3+ Contained
// Element
// Record
// Length (m)
// BYTE Varies Byte length of each Contained Element record
// that follows, in the range 0 to 255
// If no Contained Elements are included, this
// field is set to 0. For version 2.3.2 and later of
// this specification, this field is set to at least 03h
// when Contained Elements are specified.

// 15h 2.3+ Contained
// Elements
// n * m
// BYTEs
// Varies Elements, possibly defined by other SMBIOS
// structures, present in this chassis; see 7.4.4
// for definitions

// 15h +
// n*m
// 2.7+ SKU Number BYTE STRING Number of null-terminated string describing the
// chassis or enclosure SKU number
