---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_dedicatedmgmt"
description: |-
  Configure dedicated management.
---

# fmgdevice_system_dedicatedmgmt
Configure dedicated management.

## Example Usage

```hcl
resource "fmgdevice_system_dedicatedmgmt" "trname" {
  default_gateway = "your own value"
  dhcp_end_ip     = "your own value"
  dhcp_netmask    = "your own value"
  dhcp_server     = "disable"
  dhcp_start_ip   = "your own value"
  device_name     = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `default_gateway` - Default gateway for dedicated management interface.
* `dhcp_end_ip` - DHCP end IP for dedicated management.
* `dhcp_netmask` - DHCP netmask.
* `dhcp_server` - Enable/disable DHCP server on management interface. Valid values: `disable`, `enable`.

* `dhcp_start_ip` - DHCP start IP for dedicated management.
* `interface` - Dedicated management interface.
* `status` - Enable/disable dedicated management. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System DedicatedMgmt can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_dedicatedmgmt.labelname SystemDedicatedMgmt
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

