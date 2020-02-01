package smbios

var _BoardFeature = _FlagDictionary64{
	"Board is a hosting board",
	"Board requires daughter board or auxiliary card",
	"Board is removable",
	"Board is replaceable",
	"Board is hot swappable",
}

// BoardFeature collection of flags that identify features of this baseboard.
type BoardFeature byte

func (b BoardFeature) String() string {
	return _BoardFeature.str(uint64(b))
}

// Slice returns list of flags.
func (b BoardFeature) Slice() []string {
	return _BoardFeature.slice(uint64(b))
}
