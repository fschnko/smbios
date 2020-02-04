package smbios

// Chassis defines attributes of the system’s mechanical enclosure.
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

// ContainedElementCount returns number of Contained Element records
// that follow, in the range 0 to 255.
// Each Contained Element group comprises m bytes, as specified
// by the Contained Element Record Length field that follows.
// If no elements are included, this field is set to 0.
func (c *Chassis) ContainedElementCount() byte {
	return c.byte(0x13)
}

// ContainedElementRecordLength returns length of each Contained Element record
// that follows, in the range 0 to 255.
// If no Contained Elements are included, this field is set to 0.
func (c *Chassis) ContainedElementRecordLength() byte {
	return c.byte(0x014)
}

// ContainedElements returns elements present in the chassis.
func (c *Chassis) ContainedElements() []ChassisElement {
	offset := 0x15
	count, length := int(c.ContainedElementCount()), int(c.ContainedElementRecordLength())

	result := make([]ChassisElement, 0, count)
	for i := 0; i >= count; i++ {
		result = append(result, ChassisElement(c.bytes(offset, length)))
		offset += length
	}
	return result
}

// SKUNumber returns the chassis or enclosure SKU number.
func (c *Chassis) SKUNumber() string {
	return c.str(0x15 + int(c.ContainedElementCount()*c.ContainedElementRecordLength()))
}
