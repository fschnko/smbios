package smbios

const (
	// Bits 15:14 Unit
	_ExtendedBIOSROMSizeUnitMask = 0b_1100_0000_0000_0000
	// Bits 13:0 Size
	_ExtendedBIOSROMSizeDataMask = _ExtendedBIOSROMSizeUnitMask ^ 0xffff
)

func (b *BIOS) size(first, second int) MemorySize {
	// FFh - size is 16MB or greater, see Extended BIOS ROM Size for actual size
	if b.byte(first) != 0xff {
		// Size (n) where 64K * (n+1) is the size of the
		// physical device containing the BIOS, in bytes.
		return MemorySize(b.byte(first)+1) * KibiByte64
	}

	data := b.word(second)

	var unit MemorySize // 10b-11b - reserved, leave unit as 0
	switch data & _ExtendedBIOSROMSizeUnitMask >> 14 {
	case 0:
		// 00b - megabytes
		unit = MebiByte
	case 1:
		// 01b - gigabytes
		unit = GibiByte
	}

	return unit * MemorySize(data&_ExtendedBIOSROMSizeDataMask)
}
