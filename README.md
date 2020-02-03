# SMBIOS

This is a draft of the library for System Management BIOS data structures.
It based on [System Management BIOS (SMBIOS) Reference Specification v.3.3.0](https://www.dmtf.org/sites/default/files/standards/documents/DSP0134_3.3.0.pdf)

> NOTICE:
>
> Currently, work in progress and have no deadlines. You can see the status of the work in the table below.

## Status

| Table name                                | Type              | \*Status | \*\*Required |
|-------------------------------------------|-------------------|:--------:|:------------:|
| BIOS Information                          | Type 0            |    ✔     |      ✔       |
| System Information                        | Type 1            |    ✔     |      ✔       |
| Baseboard (or Module) Information         | Type 2            |    ✔     |              |
| System Enclosure or Chassis               | Type 3            |    ⏳     |      ✔       |
| Processor Information                     | Type 4            |    ✔     |      ✔       |
| Memory Controller Information             | Type 5, Obsolete  |    ⏳     |              |
| Memory Module Information                 | Type 6, Obsolete  |    ⏳     |              |
| Cache Information                         | Type 7            |    ✔     |      ✔       |
| Port Connector Information                | Type 8            |          |              |
| System Slots                              | Type 9            |          |      ✔       |
| On Board Devices Information              | Type 10, Obsolete |          |              |
| OEM Strings                               | Type 11           |          |              |
| System Configuration Options              | Type 12           |          |              |
| BIOS Language Information                 | Type 13           |          |              |
| Group Associations                        | Type 14           |          |              |
| System Event Log                          | Type 15           |          |              |
| Physical Memory Array                     | Type 16           |          |      ✔       |
| Memory Device                             | Type 17           |          |      ✔       |
| 32-Bit Memory Error Information           | Type 18           |          |              |
| Memory Array Mapped Address               | Type 19           |          |      ✔       |
| Memory Device Mapped Address              | Type 20           |          |              |
| Built-in Pointing Device                  | Type 21           |          |              |
| Portable Battery                          | Type 22           |          |              |
| System Reset                              | Type 23           |          |              |
| Hardware Security                         | Type 24           |          |              |
| System Power Controls                     | Type 25           |          |              |
| Voltage Probe                             | Type 26           |          |              |
| Cooling Device                            | Type 27           |          |              |
| Temperature Probe                         | Type 28           |          |              |
| Electrical Current Probe                  | Type 29           |          |              |
| Out-of-Band Remote Access                 | Type 30           |          |              |
| Boot Integrity Services (BIS) Entry Point | Type 31           |          |              |
| System Boot Information                   | Type 32           |     ✔    |      ✔       |
| 64-Bit Memory Error Information           | Type 33           |          |              |
| Management Device                         | Type 34           |          |              |
| Management Device Component               | Type 35           |          |              |
| Management Device Threshold Data          | Type 36           |          |              |
| Memory Channel                            | Type 37           |          |              |
| IPMI Device Information                   | Type 38           |          |              |
| System Power Supply                       | Type 39           |          |              |
| Additional Information                    | Type 40           |          |              |
| Onboard Devices Extended Information      | Type 41           |          |              |
| Management Controller Host Interface      | Type 42           |          |              |
| TPM Device                                | Type 43           |          |              |
| Processor Additional Information          | Type 44           |          |              |
| Inactive                                  | Type 126          |          |              |
| End-of-Table                              | Type 127          |          |              |

> \* ⏳ - work in progress, ✔ - done.
>
> \*\* Beginning with SMBIOS version 2.3, compliant SMBIOS implementations include the base set of required structures and data within those structures.
