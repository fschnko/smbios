package smbios

var _CacheErrorCorrection = _EnumDictionary{
	"Not set", // dictionary starts from 0x01
	"Other",
	"Unknown",
	"None",
	"Parity",
	"Single-bit ECC",
	"Multi-bit ECC",
}

// CacheErrorCorrection represents an Error Correction Type field.
type CacheErrorCorrection byte

func (c CacheErrorCorrection) String() string {
	return _CacheErrorCorrection.str(int(c))
}
