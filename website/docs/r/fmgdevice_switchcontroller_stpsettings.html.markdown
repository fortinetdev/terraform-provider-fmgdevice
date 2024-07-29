---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_stpsettings"
description: |-
  Configure FortiSwitch spanning tree protocol (STP).
---

# fmgdevice_switchcontroller_stpsettings
Configure FortiSwitch spanning tree protocol (STP).

## Example Usage

```hcl
resource "fmgdevice_switchcontroller_stpsettings" "trname" {
  forward_time = 10
  hello_time   = 10
  max_age      = 10
  max_hops     = 10
  name         = "your own value"
  device_name  = var.device_name # not required if setting is at provider
  device_vdom  = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `forward_time` - Period of time a port is in listening and learning state (4 - 30 sec, default = 15).
* `hello_time` - Period of time between successive STP frame Bridge Protocol Data Units (BPDUs) sent on a port (1 - 10 sec, default = 2).
* `max_age` - Maximum time before a bridge port expires its configuration BPDU information (6 - 40 sec, default = 20).
* `max_hops` - Maximum number of hops between the root bridge and the furthest bridge (1- 40, default = 20).
* `name` - Name of global STP settings configuration.
* `pending_timer` - Pending time (1 - 15 sec, default = 4).
* `revision` - STP revision number (0 - 65535).
* `status` - Enable/disable STP. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController StpSettings can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_stpsettings.labelname SwitchControllerStpSettings
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

