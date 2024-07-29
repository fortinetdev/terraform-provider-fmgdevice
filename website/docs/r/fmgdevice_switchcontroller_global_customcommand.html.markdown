---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_global_customcommand"
description: |-
  List of custom commands to be pushed to all FortiSwitches in the VDOM.
---

# fmgdevice_switchcontroller_global_customcommand
List of custom commands to be pushed to all FortiSwitches in the VDOM.

~> This resource is a sub resource for variable `custom_command` of resource `fmgdevice_switchcontroller_global`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_global_customcommand" "trname" {
  command_entry = "your own value"
  command_name  = ["your own value"]
  device_name   = var.device_name # not required if setting is at provider
  device_vdom   = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `command_entry` - List of FortiSwitch commands.
* `command_name` - Name of custom command to push to all FortiSwitches in VDOM.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{command_entry}}.

## Import

SwitchController GlobalCustomCommand can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_global_customcommand.labelname {{command_entry}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

