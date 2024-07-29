---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_ha_hamgmtinterfaces"
description: |-
  Reserve interfaces to manage individual cluster units.
---

# fmgdevice_system_ha_hamgmtinterfaces
Reserve interfaces to manage individual cluster units.

~> This resource is a sub resource for variable `ha_mgmt_interfaces` of resource `fmgdevice_system_ha`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_ha_hamgmtinterfaces" "trname" {
  dst         = ["your own value"]
  gateway     = "your own value"
  gateway6    = "your own value"
  fosid       = 10
  interface   = ["port2"]
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `dst` - Default route destination for reserved HA management interface.
* `gateway` - Default route gateway for reserved HA management interface.
* `gateway6` - Default IPv6 gateway for reserved HA management interface.
* `fosid` - Table ID.
* `interface` - Interface to reserve for HA management.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System HaHaMgmtInterfaces can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_ha_hamgmtinterfaces.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

