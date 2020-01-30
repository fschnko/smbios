package smbios

import "fmt"

// Version represents a semversion style version.
type Version struct {
	Major    int
	Minor    int
	Revision int
}

// NewVersion returns a new version instance.
func NewVersion(maj, min, rev int) *Version {
	return &Version{
		Major:    maj,
		Minor:    min,
		Revision: rev,
	}
}

func (v *Version) String() string {
	var ver string
	switch {
	case v.Major == -1:
		ver = "Unknown"
	case v.Revision == -1:
		ver = fmt.Sprintf("%d.%d", v.Major, v.Minor)
	default:
		ver = fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Revision)
	}
	return ver
}
