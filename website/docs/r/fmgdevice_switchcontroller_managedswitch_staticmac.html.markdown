---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_managedswitch_staticmac"
description: |-
  Configuration method to edit FortiSwitch Static and Sticky MAC.
---

# fmgdevice_switchcontroller_managedswitch_staticmac
Configuration method to edit FortiSwitch Static and Sticky MAC.

~> This resource is a sub resource for variable `static_mac` of resource `fmgdevice_switchcontroller_managedswitch`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_managedswitch_staticmac" "trname" {
  description = "your own value"
  fosid       = 10
  interface   = "your own value"
  mac         = "your own value"
  type        = "static"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `managed_switch` - Managed Switch.

* `description` - Description.
* `fosid` - ID.
* `interface` - Interface name.
* `mac` - MAC address.
* `type` - Type. Valid values: `static`, `sticky`.

* `vlan` - Vlan.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

SwitchController ManagedSwitchStaticMac can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "managed_switch=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_managedswitch_staticmac.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

