---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_virtualwirepair"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure virtual wire pairs.
---

# fmgdevice_system_virtualwirepair
<i>This object will be purged after policy copy and install.</i> Configure virtual wire pairs.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `member` - Interfaces belong to the virtual-wire-pair.
* `name` - Virtual-wire-pair name. Must be a unique interface name.
* `outer_vlan_id` - Outer VLAN ID.
* `poweroff_bypass` - Enable/disable interface bypass state when power off. Valid values: `disable`, `enable`.

* `poweron_bypass` - Enable/disable interface bypass state when power on. Valid values: `disable`, `enable`.

* `vlan_filter` - VLAN ranges to allow
* `wildcard_vlan` - Enable/disable wildcard VLAN. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System VirtualWirePair can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_virtualwirepair.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

