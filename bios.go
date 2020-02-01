package smbios

import (
	"time"
)

// BIOS represents the BIOS Information structure (Type 0).
type BIOS struct {
	table
}

// Vendor returns the BIOS Vendor’s Name.
func (b *BIOS) Vendor() string {
	return b.str(0x04)
}

// Version returns the BIOS Version. This value is a free-form string
// that may contain Core and OEM version information.
func (b *BIOS) Version() string {
	return b.str(0x05)
}

// Address returns location of BIOS starting address
// (for example, 0E800h).
// NOTE: The size of the runtime BIOS image can be computed
// by subtracting the Starting Address Segment
// from 10000h and multiplying the result by 16.
func (b *BIOS) Address() MemoryAddress {
	return MemoryAddress(b.word(0x06))
}

// ReleaseDate returns BIOS release date.
func (b *BIOS) ReleaseDate() time.Time {
	return biosTime(b.str(0x08))
}

func biosTime(str string) time.Time {
	// The date string, if supplied, is in either
	// mm/dd/yy or mm/dd/yyyy format. If the year
	// portion of the string is two digits, the year is
	// assumed to be 19yy.
	// NOTE: The mm/dd/yyyy format is required for
	// SMBIOS version 2.3 and later.
	var t time.Time
	switch len(str) {
	case 10:
		t, _ = time.Parse("01/02/2006", str)
	case 8:
		t, _ = time.Parse("01/02/06", str)
	}
	return t

}

// ROMSize returns size of the physical device(s) containing the BIOS, rounded up if needed.
func (b *BIOS) ROMSize() MemorySize {
	return b.size(0x09, 0x18)
}

// Characteristics returns which functions the BIOS supports:
// PCI, PCMCIA, Flash, etc. (see 7.1.1).
func (b *BIOS) Characteristics() BIOSCharacteristics {
	return BIOSCharacteristics{
		Low:   uint64(b.qword(0x0A)),
		Hight: uint64(b.dword(0x12)),
	}
}

// Release returns the release of the System BIOS.
func (b *BIOS) Release() *Version {
	if b.byte(0x14) != 0xff {
		return NewVersion(
			int(b.byte(0x14)),
			int(b.byte(0x15)),
			-1,
		)
	}
	return NewVersion(-1, -1, -1)
}

// FirmwareRelease returns the release of the
// embedded controller firmware.
func (b *BIOS) FirmwareRelease() *Version {
	// If the system does not have field
	// upgradeable embedded controller firmware,
	// the value is 0FFh.
	if b.byte(0x16) != 0xff {
		return NewVersion(
			int(b.byte(0x16)),
			int(b.byte(0x17)),
			-1,
		)
	}
	return NewVersion(-1, -1, -1)
}
