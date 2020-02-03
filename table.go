package smbios

import "encoding/binary"

const (
	byteLen = 1 << (1 * iota)
	wordLen
	dwordLen
	qwordLen
)

const headOffset = 4

type table struct {
	Type    uint8
	Length  uint8
	Handle  Handle
	Data    []byte
	Strings []string
}

func (t *table) byte(offset int) byte {
	offset -= headOffset
	if len(t.Data) > offset {
		return t.Data[offset]
	}
	return 0
}

func (t *table) word(offset int) uint16 {
	return binary.LittleEndian.Uint16(t.bytes(offset, 2))
}

func (t *table) dword(offset int) uint32 {
	return binary.LittleEndian.Uint32(t.bytes(offset, 4))
}

func (t *table) qword(offset int) uint64 {
	return binary.LittleEndian.Uint64(t.bytes(offset, 8))
}

func (t *table) str(offset int) string {
	snum := t.byte(offset) - 1
	if len(t.Strings) > int(snum) {
		return t.Strings[snum]
	}
	return ""
}

// bytes returns bytes in the direct order.
func (t *table) bytes(offset, length int) []byte {
	offset -= headOffset
	if offset >= 0 && len(t.Data) >= offset+length {
		return t.Data[offset : offset+length]
	}
	return nil
}

// rbytes returns bytes in the reverse order.
func (t *table) rbytes(offset, length int) []byte {
	data := t.bytes(offset, length)

	result := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		result[i] = data[len(data)-1-i]
	}
	return result
}

// bytes returns bytes in the direct order.
func (t *table) tail(offset int) []byte {
	offset -= headOffset
	if offset >= 0 && len(t.Data) > offset {
		return t.Data[offset:]
	}
	return nil
}
