package smbios

var _CacheAssociativity = _EnumDictionary{
	"Not set", // dictionary starts from 0x01
	"Other",
	"Unknown",
	"Direct Mapped",
	"2-way Set-Associative",
	"4-way Set-Associative",
	"Fully Associative",
	"8-way Set-Associative",
	"16-way Set-Associative",
	"12-way Set-Associative",
	"24-way Set-Associative",
	"32-way Set-Associative",
	"48-way Set-Associative",
	"64-way Set-Associative",
	"20-way Set-Associative",
}

// CacheAssociativity represents a Cache Associativity field.
type CacheAssociativity byte

func (c CacheAssociativity) String() string {
	return _CacheAssociativity.str(int(c))
}
