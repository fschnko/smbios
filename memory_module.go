package smbios

type MemoryModule struct {
	table
}

// SocketDesignation returns string number for reference designation.
func (m *MemoryModule) SocketDesignation() string {
	return m.str(0x04)
}

// 05h Bank Connections BYTE Varies Each nibble indicates a bank (RAS#) connection; 0xF
// means no connection.
// EXAMPLE: If banks 1 & 3 (RAS# 1 & 3) were connected to a
// SIMM socket the byte for that socket would be 13h. If only bank 2
// (RAS 2) were connected, the byte for that socket would be 2Fh.

// CurrentSpeed returns speed of the memory module,
// in ns (for example, 70d for a 70ns module).
// If the speed is unknown, the field is set to 0.
func (m *MemoryModule) CurrentSpeed() byte {
	return m.byte(0x06)
}

// CurrentType returns the physical characteristics of the memory modules
// that are currently installed in the system.
func (m *MemoryModule) CurrentType() MemoryType {
	return MemoryType(m.word(0x07))
}

// 09h Installed Size BYTE Varies See 7.7.2.

// 0Ah Enabled Size BYTE Varies See 7.7.2.

// 0Bh Error Status BYTE Varies Bits 7:3 Reserved, set to 0s
// Bit 2 If set, the Error Status information should be
// obtained from the event log; bits 1and 0 are
// reserved.
// Bit 1 Correctable errors received for the module, if
// set. This bit is reset only during a system reset.
// Bit 0 Uncorrectable errors received for the module, if
// set. All or a portion of the module has been
// disabled. This bit is only reset on power-on.
