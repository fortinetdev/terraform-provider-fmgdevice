---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_switchgroup"
description: |-
  Configure FortiSwitch switch groups.
---

# fmgdevice_switchcontroller_switchgroup
Configure FortiSwitch switch groups.

## Example Usage

```hcl
resource "fmgdevice_switchcontroller_switchgroup" "trname" {
  description = "your own value"
  fortilink   = ["your own value"]
  members     = ["your own value"]
  name        = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `description` - Optional switch group description.
* `fortilink` - FortiLink interface to which switch group members belong.
* `members` - FortiSwitch members belonging to this switch group.
* `name` - Switch group name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController SwitchGroup can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_switchgroup.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

