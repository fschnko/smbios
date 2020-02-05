package smbios

var (
	_ConnectorType = _EnumDictionary{
		"None",
		"Centronics",
		"Mini Centronics",
		"Proprietary",
		"DB-25 pin male",
		"DB-25 pin female",
		"DB-15 pin male",
		"DB-15 pin female",
		"DB-9 pin male",
		"DB-9 pin female",
		"RJ-11",
		"RJ-45",
		"50-pin MiniSCSI",
		"Mini-DIN",
		"Micro-DIN",
		"PS/2",
		"Infrared",
		"HP-HIL",
		"Access Bus (USB)",
		"SSA SCSI",
		"Circular DIN-8 male",
		"Circular DIN-8 female",
		"On Board IDE",
		"On Board Floppy",
		"9-pin Dual Inline (pin 10 cut)",
		"25-pin Dual Inline (pin 26 cut)",
		"50-pin Dual Inline",
		"68-pin Dual Inline",
		"On Board Sound Input from CD-ROM",
		"Mini-Centronics Type-14",
		"Mini-Centronics Type-26",
		"Mini-jack (headphones)",
		"BNC",
		"1394",
		"SAS/SATA Plug Receptacle",
		"USB Type-C Receptacle",
	}

	_ConnectorType0xA0 = _EnumDictionary{
		"PC-98",
		"PC-98Hireso",
		"PC-H98",
		"PC-98Note",
		"PC-98Full",
	}
)

// ConnectorType represents a connector type.
type ConnectorType byte

func (c ConnectorType) String() string {
	switch {
	case c >= 0 && c <= 0x23:
		return _ConnectorType.str(int(c))
	case c >= 0xa0 && c <= 0xa4:
		return _ConnectorType0xA0.str(int(c - 0xa0))
	case c == 0xff:
		return "Other"
	default:
		return "Unknown"
	}
}

var (
	_PortType = _EnumDictionary{
		"None",
		"Parallel Port XT/AT Compatible",
		"Parallel Port PS/2",
		"Parallel Port ECP",
		"Parallel Port EPP",
		"Parallel Port ECP/EPP",
		"Serial Port XT/AT Compatible",
		"Serial Port 16450 Compatible",
		"Serial Port 16550 Compatible",
		"Serial Port 16550A Compatible",
		"SCSI Port",
		"MIDI Port",
		"Joy Stick Port",
		"Keyboard Port",
		"Mouse Port",
		"SSA SCSI",
		"USB",
		"FireWire (IEEE P1394)",
		"PCMCIA Type I2",
		"PCMCIA Type II",
		"PCMCIA Type III",
		"Cardbus",
		"Access Bus Port",
		"SCSI II",
		"SCSI Wide",
		"PC-98",
		"PC-98-Hireso",
		"PC-H98",
		"Video Port",
		"Audio Port",
		"Modem Port",
		"Network Port",
		"SATA",
		"SAS",
		"MFDP (Multi-Function Display Port)",
		"Thunderbolt",
	}

	_PortType0xA0 = _EnumDictionary{
		"8251 Compatible",
		"8251 FIFO Compatible",
	}
)

// PortType represents a port type.
type PortType byte

func (p PortType) String() string {
	switch {
	case p >= 0 && p <= 0x23:
		return _PortType.str(int(p))
	case p >= 0xa0 && p <= 0xa1:
		return _PortType0xA0.str(int(p - 0xa0))
	case p == 0xff:
		return "Other"
	default:
		return "Unknown"
	}
}
