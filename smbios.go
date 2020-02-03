package smbios

import (
	"fmt"

	gosmbios "github.com/digitalocean/go-smbios/smbios"
)

func New() (*Explorer, error) {
	rc, ep, err := gosmbios.Stream()
	if err != nil {
		return nil, fmt.Errorf("failed to open stream: %w", err)
	}
	defer rc.Close()

	exp := &Explorer{
		Version: NewVersion(ep.Version()),
	}

	ss, err := gosmbios.NewDecoder(rc).Decode()
	if err != nil {
		return nil, fmt.Errorf("failed to decode structures: %w", err)
	}

	for _, s := range ss {
		exp.tables = append(exp.tables, table{
			Handle:  Handle(s.Header.Handle),
			Type:    s.Header.Type,
			Length:  s.Header.Length,
			Data:    s.Formatted,
			Strings: s.Strings,
		})
	}

	return exp, nil
}

type Explorer struct {
	Version *Version
	tables  []table
}

func (exp *Explorer) BIOS() *BIOS {
	for _, t := range exp.tables {
		if t.Type == 0 {
			return &BIOS{t}
		}
	}
	return &BIOS{}
}

func (exp *Explorer) SystemInformation() *SystemInformation {
	for _, t := range exp.tables {
		if t.Type == 1 {
			return &SystemInformation{t}
		}
	}
	return &SystemInformation{}
}

func (exp *Explorer) Baseboard() *Baseboard {
	for _, t := range exp.tables {
		if t.Type == 2 {
			return &Baseboard{t}
		}
	}
	return &Baseboard{}
}

func (exp *Explorer) Chassis() *Chassis {
	for _, t := range exp.tables {
		if t.Type == 3 {
			return &Chassis{t}
		}
	}
	return &Chassis{}
}

func (exp *Explorer) Processor() *Processor {
	for _, t := range exp.tables {
		if t.Type == 4 {
			return &Processor{t}
		}
	}
	return &Processor{}
}

func (exp *Explorer) SystemBootInformation() *SystemBoot {
	for _, t := range exp.tables {
		if t.Type == 32 {
			return &SystemBoot{t}
		}
	}
	return &SystemBoot{}
}
