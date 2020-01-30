package smbios

const (
	// Bits 2:0 CPU Status
	_ProcessorStatusDataMask = 0x07

	// Bit 6 CPU Socket Populated
	_ProcessorStatusPopulatedMask = 0x80
)

var _ProcessorStatus = _EnumDictionary{
	"Unknown",
	"CPU Enabled",
	"CPU Disabled by User through BIOS Setup",
	"CPU Disabled By BIOS (POST Error)",
	"CPU is Idle, waiting to be	enabled",
}

// ProcessorStatus retpresents processor status information field.
type ProcessorStatus byte

// Status returns processor status information.
func (s ProcessorStatus) Status() string {
	return _ProcessorStatus.str(int(s & _ProcessorStatusDataMask))
}

// Populated returns true if the processor socket is populated.
func (s ProcessorStatus) Populated() bool {
	return s&_ProcessorStatusPopulatedMask != 0
}

func (s ProcessorStatus) String() string {
	status := s.Status() + "/CPU Socket "
	if s.Populated() {
		status += "Populated"
	} else {
		status += "Unpopulated"
	}
	return status
}
