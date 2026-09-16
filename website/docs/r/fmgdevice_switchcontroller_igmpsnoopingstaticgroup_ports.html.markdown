---
subcategory: "Switch Controller"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_igmpsnoopingstaticgroup_ports"
description: |-
  Configure static group in switches.
---

# fmgdevice_switchcontroller_igmpsnoopingstaticgroup_ports
Configure static group in switches.

~> This resource is a sub resource for variable `ports` of resource `fmgdevice_switchcontroller_igmpsnoopingstaticgroup`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `igmp_snooping_static_group` - Igmp Snooping Static Group.

* `fosid` - ID.
* `ports` - Port members.
* `switch_id` - Switch ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

SwitchController IgmpSnoopingStaticGroupPorts can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "igmp_snooping_static_group=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_igmpsnoopingstaticgroup_ports.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

