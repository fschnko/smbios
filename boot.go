package smbios

// SystemBoot represents the System Boot Information table (Type 32).
type SystemBoot struct {
	table
}

// Status provides information about system boot status.
func (s *SystemBoot) Status() []SystemBootStatus {
	result := &systemBootStatusList{}
	for _, b := range s.tail(0x0a) {
		result.Add(SystemBootStatus(b))
	}

	return *result
}
