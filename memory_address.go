package smbios

import "fmt"

type MemoryAddress uint64

func (ma MemoryAddress) String() string {
	// %#X prints 0XFFFF insted of 0xFFFF
	return fmt.Sprintf("0x%X", uint64(ma))
}
