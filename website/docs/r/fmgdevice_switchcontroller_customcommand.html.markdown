---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_customcommand"
description: |-
  Configure the FortiGate switch controller to send custom commands to managed FortiSwitch devices.
---

# fmgdevice_switchcontroller_customcommand
Configure the FortiGate switch controller to send custom commands to managed FortiSwitch devices.

## Example Usage

```hcl
resource "fmgdevice_switchcontroller_customcommand" "trname" {
  command      = "your own value"
  command_name = "your own value"
  description  = "your own value"
  device_name  = var.device_name # not required if setting is at provider
  device_vdom  = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `command` - String of commands to send to FortiSwitch devices (For example (%0a = return key): config switch trunk %0a edit myTrunk %0a set members port1 port2 %0a end %0a).
* `command_name` - Command name called by the FortiGate switch controller in the execute command.
* `description` - Description.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{command_name}}.

## Import

SwitchController CustomCommand can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_customcommand.labelname {{command_name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

