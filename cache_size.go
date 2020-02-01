package smbios

const (
	// Bit 15 Granularity, 0 – 1K granularity, 1 – 64K granularity
	_CacheSizeFirstFieldUnitMask = 0b_1000_0000_0000_0000
	// Bits 14:0 Max size in given granularity
	_CacheSizeFirstFieldDataMask = _CacheSizeFirstFieldUnitMask ^ 0xffff
	// Bit 31 Granularity 0 – 1K granularity, 1 – 64K granularity (always 1b for cache sizes >2047 MB)
	_CacheSizeSecondFieldUnitMask = 0b_1000_0000_0000_0000_0000_0000_0000_0000
	// Bits 30:0 Max size in given granularity
	_CacheSizeSecondFieldDataMask = _CacheSizeSecondFieldUnitMask ^ 0xffffffff
)

func (c *Cache) size(first, second int) MemorySize {
	var value uint64

	gran := KibiByte
	if c.word(first) != 0xffff {
		if c.word(first)&_CacheSizeFirstFieldUnitMask > 0 {
			gran = KibiByte64
		}

		value = uint64(c.word(first) & _CacheSizeFirstFieldDataMask)
	} else {
		if c.dword(second)&_CacheSizeSecondFieldUnitMask > 0 {
			gran = KibiByte64
		}

		value = uint64(c.dword(second) & _CacheSizeSecondFieldDataMask)
	}

	return gran * MemorySize(value)
}
