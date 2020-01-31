package smbios

import "fmt"

const (
	// Bits 9:8 Operational Mode
	_CacheConfigurationModeMask = 0b_0000_0000_0011_0000_0000
	// Bit 7 Enabled/Disabled (at boot time)
	_CacheConfigurationEnabledMask = 0b_0000_0000_0000_1000_0000
	// Bits 6:5 Location, relative to the CPU module
	_CacheConfigurationLocationMask = 0b_0000_0000_0000_0110_0000
	// Bit 3 Cache Socketed
	_CacheConfigurationSocketedMask = 0b_0000_0000_0000_0000_1000
	// Bits 2:0 Cache Level
	_CacheConfigurationLevelMask = 0b_0000_0000_0000_0000_0111
)

// CacheConfiguration represents cache configuration information.
type CacheConfiguration uint16

// Mode returns the cache operational mode.
func (c CacheConfiguration) Mode() CacheMode {
	return CacheMode(c & _CacheConfigurationModeMask >> 8)
}

// Enabled returns true if cache enabled at boot time.
func (c CacheConfiguration) Enabled() bool {
	// 1 – Enabled, 0 – Disabled
	return c&_CacheConfigurationEnabledMask > 0
}

// Location returns location, relative to the CPU module.
func (c CacheConfiguration) Location() CacheLocation {
	return CacheLocation(c & _CacheConfigurationLocationMask >> 5)
}

// Socketed returns true if the cache is socketed.
func (c CacheConfiguration) Socketed() bool {
	// 1 – Socketed, 0 – Not Socketed
	return c&_CacheConfigurationSocketedMask > 0
}

// Level returns cache level information.
func (c CacheConfiguration) Level() CacheLevel {
	return CacheLevel(c & _CacheConfigurationLevelMask)
}

// CacheMode represents cache operational mode.
type CacheMode byte

// Cache operational modes.
const (
	CacheModeWriteThrough CacheMode = iota
	CacheModeWriteBack
	CacheModeVariesMemoryAddress
	CacheModeUnknown
)

func (c CacheMode) String() string {
	switch c {
	case CacheModeWriteThrough:
		return "Write Through"
	case CacheModeWriteBack:
		return "Write Back"
	case CacheModeVariesMemoryAddress:
		return "Varies with Memory Address"
	default:
		return "Unknown"
	}
}

// CacheLocation represents location, relative to the CPU module.
type CacheLocation byte

// Cache locations, relative to the CPU module.
const (
	CacheLocationInternal CacheLocation = iota
	CacheLocationExternal
)

func (c CacheLocation) String() string {
	switch c {
	case CacheLocationInternal:
		// 00b – Internal
		return "Internal"
	case CacheLocationExternal:
		// 01b – External
		return "External"
	default:
		// 10b – Reserved, 11b – Unknown
		return "Unknown"

	}
}

// CacheLevel represents a cache level – 1 through 8
// For example, an L1 cache would use value 000b
// and an L3 cache would use 010b.
type CacheLevel byte

func (c CacheLevel) String() string {
	return fmt.Sprintf("L%d", int(c)+1)
}
