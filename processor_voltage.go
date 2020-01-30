package smbios

import "fmt"

const (
	// Bit 7 Set to 0, indicating ‘legacy’ mode for processor voltage
	_ProcessorVoltageModeFlagMask = 0x80
	// Bits 3:0 Voltage Capability
	// Bit 3 – Reserved, must be zero.
	_ProcessorVoltageLegacyDataMask = 0x07

	_ProcessorVoltageDataMask = 0x7f
)

// The ProcessorVoltage describes information of processor voltage.
type ProcessorVoltage byte

// Value returns current processor voltage value.
// If Legacy Mode enabled returns specific voltage
// that the processor socket can accept.
// Value of -1.0 indicates the socket is configurable.
func (v ProcessorVoltage) Value() float32 {
	if v.LegacyMode() {
		switch v & _ProcessorVoltageLegacyDataMask {
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
			// Setting of multiple bits indicates the socket is
			// configurable.
			return -1.0
		}
	}
	// If bit 7 is set to 1, the remaining seven bits of the field
	// are set to contain the processor’s current voltage times 10.
	return float32(v&_ProcessorVoltageDataMask) / 10
}

// LegacyMode indicates that the voltage is supported.
func (v ProcessorVoltage) LegacyMode() bool {
	return v&_ProcessorVoltageModeFlagMask == 0
}

func (v ProcessorVoltage) String() string {
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
