---
subcategory: "Switch Controller"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_igmpsnoopingstaticgroup"
description: |-
  Configure FortiSwitch IGMP snooping static group settings.
---

# fmgdevice_switchcontroller_igmpsnoopingstaticgroup
Configure FortiSwitch IGMP snooping static group settings.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `ports`: `fmgdevice_switchcontroller_igmpsnoopingstaticgroup_ports`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `description` - IGMP snooping static group description.
* `ignore_reports` - Enable/disable this ignore-reports. Valid values: `disable`, `enable`.

* `mcast_addr` - IGMP snooping static group multicast IP.
* `name` - IGMP snooping static group name.
* `ports` - Ports. The structure of `ports` block is documented below.
* `vlan` - VLAN name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `ports` block supports:

* `id` - ID.
* `ports` - Port members.
* `switch_id` - Switch ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController IgmpSnoopingStaticGroup can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_igmpsnoopingstaticgroup.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

