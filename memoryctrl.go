package smbios

// MemoryController represents a memory controller table (Obsolete).
type MemoryController struct {
	table
}

// ErrorDetecting returns a memory controller's error detecting method.
func (c *MemoryController) ErrorDetecting() MemoryControllerErrDecting {
	return MemoryControllerErrDecting(c.byte(0x04))
}

// ErrorCorrecting returns a memory controller's error correcting capability.
func (c *MemoryController) ErrorCorrecting() MemoryControllerErrCorrecting {
	return MemoryControllerErrCorrecting(c.byte(0x05))
}

// SupportedInterleave returns a memory controller's supported interleave information.
func (c *MemoryController) SupportedInterleave() MemoryControllerInterleave {
	return MemoryControllerInterleave(c.byte(0x06))
}

// CurrentInterleave returns a memory controller's current interleave information.
func (c *MemoryController) CurrentInterleave() MemoryControllerInterleave {
	return MemoryControllerInterleave(c.byte(0x07))
}

// 08h 2.0+ MaxModuleSize
// BYTE Varies (n) Size of the largest memory module supported
// (per slot), specified as n, where 2**n is the
// maximum size in MB
// The maximum amount of memory supported by
// this controller is that value times the number of
// slots, as specified in offset 0Eh of this structure.

// SupportedSpeeds returns a memory controller supported speed information.
func (c *MemoryController) SupportedSpeeds() MemoryControllerSpeed {
	return MemoryControllerSpeed(c.word(0x09))
}

// SupportedTypes returns the physical characteristics of the memory modules
// that are supported by the system.
func (c *MemoryController) SupportedTypes() MemoryType {
	return MemoryType(c.word(0x0b))
}

// ModuleVoltage returns the required voltages for each of the
// memory module sockets controlled by the controller.
func (c *MemoryController) ModuleVoltage() MemoryVoltage {
	return MemoryVoltage(c.byte(0x0d))
}

// AssociatedSlots returns how many of the Memory Module
// Information blocks are controlled by this controller
func (c *MemoryController) AssociatedSlots() byte {
	return c.byte(0x0e)
}

// 0Fh to
// 0Fh +
// (2*x) - 1
// 2.0+ Memory
// Module
// Configuration
// Handles
// x
// WORDs
// Varies Lists memory information structure handles
// controlled by this controller
// Value in offset 0Eh (x) defines the count.

// 0Fh +
// (2*x)
// 2.1+ Enabled Error
// Correcting
// Capabilities
// BYTE Bit Field Identifies the error-correcting capabilities that
// were enabled when the structure was built
// See 7.6.2 for bit-wise definitions.
