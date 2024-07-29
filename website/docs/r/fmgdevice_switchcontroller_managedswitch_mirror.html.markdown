---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_managedswitch_mirror"
description: |-
  Configuration method to edit FortiSwitch packet mirror.
---

# fmgdevice_switchcontroller_managedswitch_mirror
Configuration method to edit FortiSwitch packet mirror.

~> This resource is a sub resource for variable `mirror` of resource `fmgdevice_switchcontroller_managedswitch`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_managedswitch_mirror" "trname" {
  dst         = "your own value"
  name        = "your own value"
  src_egress  = ["your own value"]
  src_ingress = ["your own value"]
  status      = "inactive"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `managed_switch` - Managed Switch.

* `dst` - Destination port.
* `name` - Mirror name.
* `src_egress` - Source egress interfaces.
* `src_ingress` - Source ingress interfaces.
* `status` - Active/inactive mirror configuration. Valid values: `inactive`, `active`.

* `switching_packet` - Enable/disable switching functionality when mirroring. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController ManagedSwitchMirror can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "managed_switch=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_managedswitch_mirror.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

