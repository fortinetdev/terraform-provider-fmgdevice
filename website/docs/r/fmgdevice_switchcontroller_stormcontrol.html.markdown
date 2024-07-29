---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_stormcontrol"
description: |-
  Configure FortiSwitch storm control.
---

# fmgdevice_switchcontroller_stormcontrol
Configure FortiSwitch storm control.

## Example Usage

```hcl
resource "fmgdevice_switchcontroller_stormcontrol" "trname" {
  broadcast         = "enable"
  rate              = 10
  unknown_multicast = "enable"
  unknown_unicast   = "disable"
  device_name       = var.device_name # not required if setting is at provider
  device_vdom       = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `broadcast` - Enable/disable storm control to drop broadcast traffic. Valid values: `disable`, `enable`.

* `rate` - Rate in packets per second at which storm control drops excess traffic(0-10000000, default=500, drop-all=0).
* `unknown_multicast` - Enable/disable storm control to drop unknown multicast traffic. Valid values: `disable`, `enable`.

* `unknown_unicast` - Enable/disable storm control to drop unknown unicast traffic. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController StormControl can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_stormcontrol.labelname SwitchControllerStormControl
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

