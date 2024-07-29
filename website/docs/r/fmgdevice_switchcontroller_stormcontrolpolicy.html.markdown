---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_stormcontrolpolicy"
description: |-
  Configure FortiSwitch storm control policy to be applied on managed-switch ports.
---

# fmgdevice_switchcontroller_stormcontrolpolicy
Configure FortiSwitch storm control policy to be applied on managed-switch ports.

## Example Usage

```hcl
resource "fmgdevice_switchcontroller_stormcontrolpolicy" "trname" {
  broadcast          = "enable"
  description        = "your own value"
  name               = "your own value"
  rate               = 10
  storm_control_mode = "global"
  device_name        = var.device_name # not required if setting is at provider
  device_vdom        = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `broadcast` - Enable/disable storm control to drop/allow broadcast traffic in override mode. Valid values: `disable`, `enable`.

* `description` - Description of the storm control policy.
* `name` - Storm control policy name.
* `rate` - Threshold rate in packets per second at which storm traffic is controlled in override mode (default=500, 0 to drop all).
* `storm_control_mode` - Set Storm control mode. Valid values: `disabled`, `global`, `override`.

* `unknown_multicast` - Enable/disable storm control to drop/allow unknown multicast traffic in override mode. Valid values: `disable`, `enable`.

* `unknown_unicast` - Enable/disable storm control to drop/allow unknown unicast traffic in override mode. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController StormControlPolicy can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_stormcontrolpolicy.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

