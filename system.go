package smbios

// SystemInformation represents a system information table (Type 1).
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

// UUID returns an Universal unique ID number.
func (s *SystemInformation) UUID() UUID {
	var uuid UUID
	// time_low 4 bytes Low field of the timestamp
	copy(uuid[:4], s.rbytes(0x08, 4))
	// time_mid 2 bytes Middle field of the timestamp
	copy(uuid[4:6], s.rbytes(0x0c, 2))
	// time_hi_and_version 2 bytes High field of the timestamp multiplexed with the version number
	copy(uuid[6:8], s.rbytes(0x0e, 2))
	// clock_seq_hi_and_reserved byte High field of the clock sequence multiplexed with the variant
	copy(uuid[8:9], s.bytes(0x10, 1))
	// clock_seq_low byte Low field of the clock sequence
	copy(uuid[9:10], s.bytes(0x11, 1))
	// Node 6 bytes Spatially unique node identifier
	copy(uuid[10:], s.bytes(0x12, 6))

	return uuid
}

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
