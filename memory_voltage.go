package smbios

import "fmt"

// Bits 7:3 Reserved, must be zero
const _MemoryVoltageDataMask = 0x07

// MemoryVoltage describes the required voltages
// for the memory module sockets.
type MemoryVoltage byte

// Value returns memory socket voltage value.
// Value of -1.0 indicates the sockets are configurable.
func (v MemoryVoltage) Value() float32 {
	switch v & _MemoryVoltageDataMask {
	case 0:
		return 0.0
	case 1:
		// Bit 0 – 5V
		return 5.0
	case 2:
		// Bit 1 – 3.3V
		return 3.3
	case 4:
		// Bit 2 – 2.9V
		return 2.9
	default:
		// Setting of multiple bits indicates that the
		// sockets are configurable.
		return -1.0
	}
}

func (v MemoryVoltage) String() string {
	val := v.Value()
	switch val {
	case 0:
		return "Unknown"
	case -1:
		return "Configurable"
	default:
		return fmt.Sprintf("%.1fV", val)
	}
}
