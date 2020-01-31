package smbios

const (
	// Bit 15 Granularity, 0 – 1K granularity, 1 – 64K granularity
	// Bits 14:0 Max size in given granularity
	_CacheSizeFirstFieldMask = 0b_1000_0000_0000_0000
	// Bit 31 Granularity 0 – 1K granularity, 1 – 64K granularity (always 1b for cache sizes >2047 MB)
	// Bits 30:0 Max size in given granularity
	_CacheSizeSecondFieldMask = 0b_1000_0000_0000_0000_0000_0000_0000_0000
)

func (c *Cache) size(first, second int) MemorySize {
	var value uint64

	gran := KibiByte
	if c.word(first) != 0xffff {
		if c.word(first)&_CacheSizeFirstFieldMask > 0 {
			gran = KibiByte64
		}

		value = uint64(c.word(first) ^ _CacheSizeFirstFieldMask)
	} else {
		if c.dword(second)&_CacheSizeSecondFieldMask > 0 {
			gran = KibiByte64
		}

		value = uint64(c.dword(second) ^ _CacheSizeSecondFieldMask)
	}

	return gran * MemorySize(value)
}
