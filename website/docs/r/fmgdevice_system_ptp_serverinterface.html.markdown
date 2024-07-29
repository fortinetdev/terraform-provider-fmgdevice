---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_ptp_serverinterface"
description: |-
  FortiGate interface(s) with PTP server mode enabled. Devices on your network can contact these interfaces for PTP services.
---

# fmgdevice_system_ptp_serverinterface
FortiGate interface(s) with PTP server mode enabled. Devices on your network can contact these interfaces for PTP services.

~> This resource is a sub resource for variable `server_interface` of resource `fmgdevice_system_ptp`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_ptp_serverinterface" "trname" {
  delay_mechanism       = "E2E"
  fosid                 = 10
  server_interface_name = ["port2"]
  device_name           = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `delay_mechanism` - End to end delay detection or peer to peer delay detection. Valid values: `E2E`, `P2P`.

* `fosid` - ID.
* `server_interface_name` - Interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System PtpServerInterface can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_ptp_serverinterface.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

