---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_autoconfig_custom_switchbinding"
description: |-
  Switch binding list.
---

# fmgdevice_switchcontroller_autoconfig_custom_switchbinding
Switch binding list.

~> This resource is a sub resource for variable `switch_binding` of resource `fmgdevice_switchcontroller_autoconfig_custom`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_autoconfig_custom_switchbinding" "trname" {
  policy      = ["your own value"]
  switch_id   = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `custom` - Custom.

* `policy` - Custom auto-config policy.
* `switch_id` - Switch name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{switch_id}}.

## Import

SwitchController AutoConfigCustomSwitchBinding can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "custom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_autoconfig_custom_switchbinding.labelname {{switch_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

