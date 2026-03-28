---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_managedswitch_routerstatic"
description: |-
  Configure static routes.
---

# fmgdevice_switchcontroller_managedswitch_routerstatic
Configure static routes.

~> This resource is a sub resource for variable `router_static` of resource `fmgdevice_switchcontroller_managedswitch`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `managed_switch` - Managed Switch.

* `blackhole` - Enable/disable blackhole on this route. Valid values: `disable`, `enable`.

* `comment` - Comment.
* `device` - Gateway out interface.
* `distance` - Administrative distance for the route (1 - 255, default = 10).
* `dst` - Destination ip and mask for this route.
* `dynamic_gateway` - Enable/disable dynamic gateway. Valid values: `disable`, `enable`.

* `gateway` - Gateway ip for this route.
* `fosid` - Entry sequence number.
* `status` - Enable/disable route status. Valid values: `disable`, `enable`.

* `switch_id` - Switch ID.
* `vrf` - VRF for this route.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

SwitchController ManagedSwitchRouterStatic can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "managed_switch=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_managedswitch_routerstatic.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

