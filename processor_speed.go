package smbios

import "fmt"

// ProcessorSpeed represents a processor speed information fields.
type ProcessorSpeed uint16

func (s ProcessorSpeed) String() string {
	if s == 0 {
		return "Unknown"
	}
	return fmt.Sprintf("%d MHz", byte(s))
}
