package smbios

// Cache defines the attributes of CPU cache device in the system.
type Cache struct {
	table
}

// SocketDesignation returns number for reference designation.
func (c *Cache) SocketDesignation() string {
	return c.str(0x04)
}

// Configuration returns cache configuration information.
func (c *Cache) Configuration() CacheConfiguration {
	return CacheConfiguration(c.word(0x05))
}

// MaximumSize returns maximum size that can be installed
func (c *Cache) MaximumSize() MemorySize {
	return c.size(0x07, 0x13)
}

// InstalledSize returns installed cache size.
// Absent or set to 0 if no cache is installed.
func (c *Cache) InstalledSize() MemorySize {
	return c.size(0x09, 0x17)
}

// SupportedSRAM returns supported SRAM type.
func (c *Cache) SupportedSRAM() CacheSRAM {
	return CacheSRAM(c.word(0x0b))
}

// CurrentSRAM returns current SRAM type.
func (c *Cache) CurrentSRAM() CacheSRAM {
	return CacheSRAM(c.word(0x0d))
}

// Speed returns a cache module speed, in nanoseconds.
// The value is 0 if the speed is unknown.
func (c *Cache) Speed() byte {
	return c.byte(0x0f)
}

// ErrorCorrection returns Error-correction scheme supported by the cache component.
func (c *Cache) ErrorCorrection() CacheErrorCorrection {
	return CacheErrorCorrection(c.byte(0x10))
}

// Type returns a logical type of the cache.
func (c *Cache) Type() CacheType {
	return CacheType(c.byte(0x11))
}

// Associativity returns an Associativity of the cache.
func (c *Cache) Associativity() CacheAssociativity {
	return CacheAssociativity(c.byte(0x12))
}
