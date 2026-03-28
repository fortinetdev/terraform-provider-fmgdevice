---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_managedswitch_routervrf"
description: |-
  Configure VRF.
---

# fmgdevice_switchcontroller_managedswitch_routervrf
Configure VRF.

~> This resource is a sub resource for variable `router_vrf` of resource `fmgdevice_switchcontroller_managedswitch`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `managed_switch` - Managed Switch.

* `name` - VRF entry name.
* `switch_id` - Switch ID.
* `vrfid` - VRF ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController ManagedSwitchRouterVrf can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "managed_switch=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_managedswitch_routervrf.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

