package smbios

type Baseboard struct {
	table
}

// Manufacturer returns baseboard manufacturer name.
func (b *Baseboard) Manufacturer() string {
	return b.str(0x04)
}

// Product returns baseboard product name.
func (b *Baseboard) Product() string {
	return b.str(0x05)
}

// Version returns baseboard product version.
func (b *Baseboard) Version() string {
	return b.str(0x06)
}

// SerialNumber returns baseboard serial number.
func (b *Baseboard) SerialNumber() string {
	return b.str(0x07)
}

// AssetTag returns baseboard asset tag.
func (b *Baseboard) AssetTag() string {
	return b.str(0x08)
}

// Features returns a collection of flags that identify features of the board.
// See 7.3.1
func (b *Baseboard) Features() BoardFeature {
	return BoardFeature(b.byte(0x09))
}

// Location returns location within the chassis referenced by the
// Chassis Handle.
func (b *Baseboard) Location() string {
	return b.str(0x0a)
}

// ChassisHandle returns handle associated with the chassis in
// which the board resides (see 7.4)
func (b *Baseboard) ChassisHandle() Handle {
	return Handle(b.word(0x0b))
}

// Type returns type of the board (see 7.3.2).
func (b *Baseboard) Type() BaseboardType {
	return BaseboardType(b.byte(0x0d))
}

// ObjectsNumber returns number
func (b *Baseboard) ObjectsNumber() byte {
	return b.byte(0x0e)
}

// Objects returns list of objects that are contained by the board.
func (b *Baseboard) Objects() []Handle {
	objs := make([]Handle, 0)
	for i := 0x0f; len(b.Data) > i+wordLen; i += wordLen {
		objs = append(objs, Handle(b.word(i)))
	}
	return objs
}
