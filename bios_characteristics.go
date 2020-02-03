package smbios

const (
	_BIOSCharacteristicsLowMask   = 0b_00000000_00000000_00000000_00000000_11111111_11111111_11111111_11111100
	_BIOSCharacteristicsHightMask = 0b_00011111_11111111
)

var _BIOSCharacteristics = _FlagDictionary128{
	"Reserved",
	"Reserved",
	"Unknown",
	"BIOS Characteristics are not supported",
	"ISA is supported",
	"MCA is supported",
	"EISA is supported",
	"PCI is supported",
	"PC card (PCMCIA) is supported",
	"Plug and Play is supported",
	"APM is supported",
	"BIOS is upgradeable (Flash)",
	"BIOS shadowing is allowed",
	"VL-VESA is supported",
	"ESCD support is available",
	"Boot from CD is supported",
	"Selectable boot is supported",
	"BIOS ROM is socketed",
	"Boot from PC card (PCMCIA) is supported",
	"EDD specification is supported",
	"Japanese floppy for NEC 9800 1.2 MB is supported (int 13h)",
	"Japanese floppy for Toshiba 1.2 MB is supported (int 13h)",
	"5.25” / 360 KB floppy services are supported (int 13h)",
	"5.25” /1.2 MB floppy services are supported (int 13h)",
	"3.5” / 720 KB floppy services are supported (int 13h)",
	"3.5” / 2.88 MB floppy services are supported (int 13h)",
	"Print screen Service is supported (int 5h)",
	"8042 keyboard services are supported (int 9h)",
	"Serial services are supported (int 14h)",
	"Printer services are supported (int 17h)",
	"CGA/Mono Video Services are supported (int 10h)",
	"NEC PC-98",
	// Bits32:47 Reserved for BIOS vendor.
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	// Bits 48:63 Reserved for system vendor.
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	// BIOS Characteristics Extension Byte 1
	"ACPI is supported",
	"USB Legacy is supported",
	"AGP is supported",
	"I2O boot is supported",
	"LS-120 SuperDisk boot is supported",
	"ATAPI ZIP drive boot is supported",
	"1394 boot is supported",
	"Smart battery is supported",
	// BIOS Characteristics Extension Byte 2
	"BIOS Boot Specification is supported",
	"Function key-initiated network service boot is supported",
	"Targeted content distribution is supported",
	"UEFI is supported",
	"SMBIOS table describes a virtual machine",
	// Bits 5:7 Reserved for future assignment.
	"Reserved",
	"Reserved",
	"Reserved",
}

// BIOSCharacteristics represents functions which the BIOS supports.
type BIOSCharacteristics struct {
	Low, Hight uint64
}

func (c BIOSCharacteristics) String() string {
	return _BIOSCharacteristics.str(
		_BIOSCharacteristicsHightMask&c.Hight,
		_BIOSCharacteristicsLowMask&c.Low,
	)
}

// Slice returns a list of characteristics.
func (c BIOSCharacteristics) Slice() []string {
	return _BIOSCharacteristics.slice(
		_BIOSCharacteristicsHightMask&c.Hight,
		_BIOSCharacteristicsLowMask&c.Low,
	)
}
