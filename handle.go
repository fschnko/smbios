package smbios

import "fmt"

type Handle uint16

func (h Handle) String() string {
	// %#X prints 0XFFFF insted of 0xFFFF
	return fmt.Sprintf("0x%X", uint16(h))
}
