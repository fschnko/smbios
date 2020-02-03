package smbios

var _SystemBootStatus = _EnumDictionary{
	"No errors detected",
	"No bootable media",
	"Operating system failed to load",
	"Firmware-detected hardware failure",
	"Operating system-detected hardware failure",
	"User-requested boot",
	"System security violation",
	"Previously-requested image",
	"System watchdog timer expired",
}

// systemBootStatusList represents a list of the boot statuses.
type systemBootStatusList []SystemBootStatus

// Add adds unique status to the list.
func (s *systemBootStatusList) Add(status SystemBootStatus) {
	if !s.exists(status) {
		*s = append(*s, status)
	}
}

func (s *systemBootStatusList) exists(status SystemBootStatus) bool {
	for _, v := range *s {
		if status == v {
			return true
		}
	}
	return false
}

// SystemBootStatus represents status and additional data fields
// that identify the boot status.
type SystemBootStatus byte

func (s SystemBootStatus) String() string {
	switch {
	case s <= 8:
		return _SystemBootStatus.str(int(s))
	case s >= 128 && s <= 191:
		// 128-191 Vendor/OEM-specific implementations
		return "OEM-specific"
	case s >= 192:
		// 192-255  Product-specific implementations
		return "Product-specific"
	default:
		return "Unknown"
	}
}
