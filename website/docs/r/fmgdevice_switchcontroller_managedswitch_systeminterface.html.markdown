---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_managedswitch_systeminterface"
description: |-
  Configure system interface on FortiSwitch.
---

# fmgdevice_switchcontroller_managedswitch_systeminterface
Configure system interface on FortiSwitch.

~> This resource is a sub resource for variable `system_interface` of resource `fmgdevice_switchcontroller_managedswitch`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `managed_switch` - Managed Switch.

* `allowaccess` - Permitted types of management access to this interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `radius-acct`.

* `interface` - Interface name.
* `ip` - IP and mask for this interface.
* `mode` - Interface addressing mode. Valid values: `static`, `dhcp`.

* `name` - Interface name.
* `status` - Enable/disable interface status. Valid values: `disable`, `enable`.

* `switch_id` - Switch ID.
* `type` - Interface type. Valid values: `physical`, `vlan`.

* `vlan` - VLAN name.
* `vrf` - VRF for this route.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController ManagedSwitchSystemInterface can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "managed_switch=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_managedswitch_systeminterface.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

