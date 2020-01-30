package smbios

type Cache struct {
	table
}

// SocketDesignation returns number for reference designation.
func (c *Cache) SocketDesignation() string {
	return c.str(0x04)
}

// 05h 2.0+ Cache Configuration WORD Varies Bits 15:10 Reserved, must be zero
// Bits 9:8 Operational Mode
// 00b – Write Through
// 01b – Write Back
// 10b – Varies with Memory
// Address
// 11b – Unknown
// Bit 7 Enabled/Disabled (at boot time)
// 1b – Enabled
// 0b – Disabled
// Bits 6:5 Location, relative to the CPU
// module:
// 00b – Internal
// 01b – External
// 10b – Reserved
// 11b – Unknown
// Bit 4 Reserved, must be zero
// Bit 3 Cache Socketed
// 1b – Socketed
// 0b – Not Socketed
// Bits 2:0 Cache Level – 1 through 8 (For
// example, an L1 cache would
// use value 000b and an L3
// cache would use 010b.)

// 07h 2.0+ Maximum Cache Size WORD Varies Maximum size that can be installed
// Bit 15 Granularity
// 0 – 1K granularity
// 1 – 64K granularity
// Bits 14:0 Max size in given granularity
// See 7.8.1.

// 09h 2.0+ Installed Size WORD Varies Same format as Max Cache Size field; set to
// 0 if no cache is installed
// See 7.8.1.

// 0Bh 2.0+ Supported SRAM
// Type
// WORD Bit Field See 7.8.2.

// 0Dh 2.0+ Current SRAM Type WORD Bit Field See 7.8.2.

// Speed returns a cache module speed, in nanoseconds.
// The value is 0 if the speed is unknown.
func (c *Cache) Speed() byte {
	return c.byte(0x0f)
}

// 10h 2.1+ Error Correction Type BYTE ENUM Error-correction scheme supported by this
// cache component; see 7.8.3

// 11h 2.1+ System Cache Type BYTE ENUM Logical type of cache; see 7.8.4

// 12h 2.1+ Associativity BYTE ENUM Associativity of the cache; see 7.8.5

// 13h 3.1+ Maximum Cache Size
// 2
// DWORD Bit Field If this field is present, for cache sizes of
// 2047 MB or smaller the value in the Max
// size in given granularity portion of the field
// equals the size given in the corresponding
// portion of the Maximum Cache Size field,
// and the Granularity bit matches the value of
// the Granularity bit in the Maximum Cache
// Size field.
// For Cache sizes greater than 2047 MB, the
// Maximum Cache Size field is set to 0xFFFF
// and the Maximum Cache Size 2 field is
// present, the Granularity bit is set to 1b, and
// the size set as required; see 7.8.1.
// Bit 31 Granularity
// 0 – 1K granularity
// 1 – 64K granularity (always 1b
// for cache sizes >2047 MB)
// Bits 30:0 Max size in given granularity

// 17h 3.1+ Installed Cache Size 2 DWORD Bit Field Same format as Maximum Cache Size 2
// field; Absent or set to 0 if no cache is
// installed.
// See 7.8.1.
