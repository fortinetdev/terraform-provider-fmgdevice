---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_ptp"
description: |-
  Configure system PTP information.
---

# fmgdevice_system_ptp
Configure system PTP information.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `server_interface`: `fmgdevice_system_ptp_serverinterface`



## Example Usage

```hcl
resource "fmgdevice_system_ptp" "trname" {
  delay_mechanism  = "E2E"
  interface        = ["port2"]
  mode             = "multicast"
  request_interval = 10
  server_interface {
    delay_mechanism       = "E2E"
    fosid                 = 10
    server_interface_name = ["port2"]
  }

  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `delay_mechanism` - End to end delay detection or peer to peer delay detection. Valid values: `E2E`, `P2P`.

* `interface` - PTP client will reply through this interface.
* `mode` - Multicast transmission or hybrid transmission. Valid values: `multicast`, `hybrid`.

* `request_interval` - The delay request value is the logarithmic mean interval in seconds between the delay request messages sent by the slave to the master.
* `server_interface` - Server-Interface. The structure of `server_interface` block is documented below.
* `server_mode` - Enable/disable FortiGate PTP server mode. Your FortiGate becomes an PTP server for other devices on your network. Valid values: `disable`, `enable`.

* `status` - Enable/disable setting the FortiGate system time by synchronizing with an PTP Server. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `server_interface` block supports:

* `delay_mechanism` - End to end delay detection or peer to peer delay detection. Valid values: `E2E`, `P2P`.

* `id` - ID.
* `server_interface_name` - Interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Ptp can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_ptp.labelname SystemPtp
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

