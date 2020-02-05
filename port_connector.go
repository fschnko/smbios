package smbios

// PortConnector defines the attributes of a system port connector (Type 8).
// For example, parallel, serial, keyboard, or mouse ports.
type PortConnector struct {
	table
}

// InternalReferenceDesignator returns number for internal to the system enclosure.
func (p *PortConnector) InternalReferenceDesignator() string {
	return p.str(0x04)
}

// InternalConnectorType returns an internal connector type.
func (p *PortConnector) InternalConnectorType() ConnectorType {
	return ConnectorType(p.byte(0x05))
}

// ExternalReferenceDesignator returns number for the to the system enclosure.
func (p *PortConnector) ExternalReferenceDesignator() string {
	return p.str(0x06)
}

// ExternalConnectorType returns an external connector type.
// See 7.9.2.
func (p *PortConnector) ExternalConnectorType() ConnectorType {
	return ConnectorType(p.byte(0x07))
}

// PortType returns the function of the port.
func (p *PortConnector) PortType() PortType {
	return PortType(p.byte(0x08))
}
