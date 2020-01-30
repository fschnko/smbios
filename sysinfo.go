package smbios

// SystemInformation represents a system information table.
type SystemInformation struct {
	table
}

// Manufacturer returns  product manufacturer name.
func (s *SystemInformation) Manufacturer() string {
	return s.str(0x04)
}

// ProductName returns product name.
func (s *SystemInformation) ProductName() string {
	return s.str(0x05)
}

// Version returns product version.
func (s *SystemInformation) Version() string {
	return s.str(0x06)
}

// SerialNumber returns product serial number.
func (s *SystemInformation) SerialNumber() string {
	return s.str(0x07)
}

// 08h 2.1+ UUID 16 BYTEs Varies Universal unique ID number; see 7.2.1.

// WakeUp returns the event that caused the system to power up.
// See 7.2.2.
func (s *SystemInformation) WakeUp() SystemWakeUp {
	return SystemWakeUp(s.byte(0x18))
}

// SKUNumber returns string that identifies a particular computer
// configuration for sale. It is sometimes also
// called a product ID or purchase order number.
// This number is frequently found in existing
// fields, but there is no standard format.
// Typically for a given system board from a
// given OEM, there are tens of unique
// processor, memory, hard drive, and optical
// drive configurations.
func (s *SystemInformation) SKUNumber() string {
	return s.str(0x19)
}

// Family returns string that identifies the family to which a
// particular computer belongs. A family refers to
// a set of computers that are similar but not
// identical from a hardware or software point of
// view. Typically, a family is composed of
// different computer models, which have
// different configurations and pricing points.
// Computers in the same family often have
// similar branding and cosmetic features.
func (s *SystemInformation) Family() string {
	return s.str(0x1a)
}
