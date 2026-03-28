---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_vipgrp"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure IPv4 virtual IP groups.
---

# fmgdevice_firewall_vipgrp
<i>This object will be purged after policy copy and install.</i> Configure IPv4 virtual IP groups.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `dynamic_mapping`: `fmgdevice_firewall_vipgrp_dynamic_mapping`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `color` - Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
* `comments` - Comment.
* `dynamic_mapping` - Dynamic_Mapping. The structure of `dynamic_mapping` block is documented below.
* `interface` - Interface.
* `member` - Member VIP objects of the group (Separate multiple objects with a space).
* `name` - VIP group name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dynamic_mapping` block supports:

* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `color` - Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
* `comments` - Comment.
* `interface` - Interface.
* `member` - Member VIP objects of the group (Separate multiple objects with a space).
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall Vipgrp can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_vipgrp.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

