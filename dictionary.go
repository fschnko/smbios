package smbios

import "strings"

type _FlagDictionary64 []string

func (d _FlagDictionary64) str(flags uint64) string {
	return strings.Join(d.slice(flags), "\n")
}

func (d _FlagDictionary64) slice(flags uint64) []string {
	var sl []string
	for i, str := range d {
		if flags&(1<<i) != 0 {
			sl = append(sl, str)
		}
	}
	return sl
}

type _FlagDictionary128 []string

func (d _FlagDictionary128) str(hi, lo uint64) string {
	return strings.Join(d.slice(hi, lo), "\n")
}

func (d _FlagDictionary128) slice(hi, lo uint64) []string {
	var sl []string
	for i, str := range d {
		var mask, offset uint64
		switch {
		case i <= 63:
			mask = lo
			offset = uint64(i)
		case 63 < i && i < 127:
			mask = hi
			offset = uint64(i - 64)
		}

		if mask&(1<<offset) != 0 {
			sl = append(sl, str)
		}
	}
	return sl
}

type _EnumDictionary []string

func (d _EnumDictionary) str(id int) string {
	if len(d) > id && id >= 0 {
		return d[id]
	}
	return "Unknown"
}

type _MapDictionary map[uint16]string

func (d _MapDictionary) str(key uint16) string {
	if val, ok := d[key]; ok {
		return val
	}
	return "Unknown"
}
