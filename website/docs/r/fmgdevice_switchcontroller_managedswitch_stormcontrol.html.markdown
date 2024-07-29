---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_managedswitch_stormcontrol"
description: |-
  Configuration method to edit FortiSwitch storm control for measuring traffic activity using data rates to prevent traffic disruption.
---

# fmgdevice_switchcontroller_managedswitch_stormcontrol
Configuration method to edit FortiSwitch storm control for measuring traffic activity using data rates to prevent traffic disruption.

~> This resource is a sub resource for variable `storm_control` of resource `fmgdevice_switchcontroller_managedswitch`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_managedswitch_stormcontrol" "trname" {
  broadcast         = "disable"
  local_override    = "enable"
  rate              = 10
  unknown_multicast = "disable"
  unknown_unicast   = "enable"
  device_name       = var.device_name # not required if setting is at provider
  device_vdom       = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `managed_switch` - Managed Switch.

* `broadcast` - Enable/disable storm control to drop broadcast traffic. Valid values: `disable`, `enable`.

* `local_override` - Enable to override global FortiSwitch storm control settings for this FortiSwitch. Valid values: `disable`, `enable`.

* `rate` - Rate in packets per second at which storm control drops excess traffic(0-10000000, default=500, drop-all=0).
* `unknown_multicast` - Enable/disable storm control to drop unknown multicast traffic. Valid values: `disable`, `enable`.

* `unknown_unicast` - Enable/disable storm control to drop unknown unicast traffic. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController ManagedSwitchStormControl can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "managed_switch=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_managedswitch_stormcontrol.labelname SwitchControllerManagedSwitchStormControl
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

