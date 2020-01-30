package smbios

import (
	"time"
)

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

// 09h 2.0+ BIOS ROM Size BYTE Varies (n) Size (n) where 64K * (n+1) is the size of the
// physical device containing the BIOS, in
// bytes.
// FFh - size is 16MB or greater, see Extended
// BIOS ROM Size for actual size
// 18h 3.1+ Extended BIOS
// ROM Size
// WORD Bit Field Extended size of the physical device(s)
// containing the BIOS, rounded up if needed.
// Bits 15:14 Unit
// 00b - megabytes
// 01b - gigabytes
// 10b - reserved
// 11b - reserved
// Bits 13:0 Size
// Examples: a 16 MB device would be
// represented as 0010h. A 48 GB device set
// would be represented as
// 0100_0000_0011_0000b or 4030h.

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
