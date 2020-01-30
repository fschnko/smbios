package smbios

// Processor represents a processor information table.
type Processor struct {
	table
}

// SocketDesignation returns number for Reference Designation.
func (p *Processor) SocketDesignation() string {
	return p.str(0x04)
}

// ProcessorType returns processor type information.
func (p *Processor) ProcessorType() ProcessorType {
	return ProcessorType(p.byte(0x05))
}

// ProcessorFamily return processor family information.
func (p *Processor) ProcessorFamily() ProcessorFamily {
	// 0xFE indicator to obtain the processor family
	// from the Processor Family 2 field
	if p.byte(0x06) == 0xfe {
		return ProcessorFamily(p.word(0x28))
	}
	return ProcessorFamily(p.byte(0x06))
}

// Manufacturer returns processor manufacturer name.
func (p *Processor) Manufacturer() string {
	return p.str(0x07)
}

// ID returns raw processor identification data.
// See 7.5.3 for details.
func (p *Processor) ID() []byte {
	return p.bytes(0x08, qwordLen)
}

// Version returns processor version.
func (p *Processor) Version() string {
	return p.str(0x10)
}

// Voltage returns information of processor voltage.
func (p *Processor) Voltage() ProcessorVoltage {
	return ProcessorVoltage(p.byte(0x11))
}

// ExternalClock returns External Clock Frequency, in MHz.
// If the value is unknown, the field is set to 0.
func (p *Processor) ExternalClock() ProcessorSpeed {
	return ProcessorSpeed(p.word(0x12))
}

// MaxSpeed returns maximum processor speed (in MHz) supported
// by the system for this processor socket.
// If the value is unknown, the field is set to 0.
// NOTE: This field identifies a capability for the system,
// not the processor itself.
func (p *Processor) MaxSpeed() ProcessorSpeed {
	return ProcessorSpeed(p.word(0x14))
}

// CurrentSpeed returns the processor's speed at
// system boot; the processor may support
// more than one speed.
func (p *Processor) CurrentSpeed() ProcessorSpeed {
	return ProcessorSpeed(p.word(0x16))
}

// Status returns the processor status information.
func (p *Processor) Status() ProcessorStatus {
	return ProcessorStatus(p.byte(0x18))
}

// ProcessorUpgrade returns the processor upgrade information.
func (p *Processor) ProcessorUpgrade() ProcessorUpgrade {
	return ProcessorUpgrade(p.byte(0x19))
}

// L1Cache returns Handle of a Cache Information structure that
// defines the attributes of the primary (Level 1)
// cache for this processor
// For version 2.1 and version 2.2
// implementations, the value is 0FFFFh if the
// processor has no L1 cache. For version 2.3 and
// later implementations, the value is 0FFFFh if
// the Cache Information structure is not provided.
func (p *Processor) L1Cache() Handle {
	return Handle(p.word(0x1a))
}

// L2Cache returns Handle of a Cache Information structure that
// defines the attributes of the secondary (Level 2)
// cache for this processor
// For version 2.1 and version 2.2
// implementations, the value is 0FFFFh if the
// processor has no L2 cache. For version 2.3 and
// later implementations, the value is 0FFFFh if
// the Cache Information structure is not provided.
func (p *Processor) L2Cache() Handle {
	return Handle(p.word(0x1c))
}

// L3Cache returns Handle of a Cache Information structure that
// defines the attributes of the tertiary (Level 3)
// cache for this processor
// For version 2.1 and version 2.2
// implementations, the value is 0FFFFh if the
// processor has no L3 cache. For version 2.3 and
// later implementations, the value is 0FFFFh if
// the Cache Information structure is not provided.
func (p *Processor) L3Cache() Handle {
	return Handle(p.word(0x1e))
}

// SerialNumber returns the serial number of the processor.
// This value is set by the manufacturer and normally not changeable.
func (p *Processor) SerialNumber() string {
	return p.str(0x20)
}

// AssetTag returns the asset tag of the processor.
func (p *Processor) AssetTag() string {
	return p.str(0x21)
}

// PartNumber returns the part number of the processor.
// This value is set by the manufacturer and normally not changeable.
func (p *Processor) PartNumber() string {
	return p.str(0x22)
}

// CoreCount returns number of cores per processor socket.
// See 7.5.6. If the value is unknown, the field is set to 0.
func (p *Processor) CoreCount() uint16 {
	// For core counts of 256 or greater, the Core Count field is set to FFh
	// and the Core Count 2 field is set to the number of cores.
	if p.byte(0x23) == 0xff {
		return p.word(0x2a)
	}
	return uint16(p.byte(0x23))
}

// CoreEnabled returns number of enabled cores per processor socket
// See 7.5.7. If the value is unknown, the field is set 0.
func (p *Processor) CoreEnabled() uint16 {
	// For core counts of 256 or greater, the Core Enabled field is set to FFh
	// and the Core Enabled 2 field is set to the number of enabled cores.
	if p.byte(0x24) == 0xff {
		return p.word(0x2c)
	}
	return uint16(p.byte(0x24))
}

// ThreadCount returns number of threads per processor socket
// See 7.5.8. If the value is unknown, the field is set to 0.
func (p *Processor) ThreadCount() uint16 {
	// For thread counts of 256 or greater, the Thread Count field is set to FFh
	// and the Thread Count 2 field is set to the number of threads.
	if p.byte(0x25) == 0xff {
		return p.word(0x2e)
	}
	return uint16(p.byte(0x25))
}

// Characteristics returns which functions the processor supports.
// See 7.5.9.
func (p *Processor) Characteristics() ProcessorCharacteristics {
	return ProcessorCharacteristics(p.word(0x26))
}
