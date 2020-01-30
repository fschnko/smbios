package smbios

var _BoardFeature = _FlagDictionary64{
	"Board is a hosting board",
	"Board requires daughter board or auxiliary card",
	"Board is removable",
	"Board is replaceable",
	"Board is hot swappable",
}

type BoardFeature byte

func (b BoardFeature) String() string {
	return _BoardFeature.str(uint64(b))
}

func (b BoardFeature) Slice() []string {
	return _BoardFeature.slice(uint64(b))
}
