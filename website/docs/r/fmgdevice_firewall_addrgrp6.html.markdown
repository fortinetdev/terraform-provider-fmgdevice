---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_addrgrp6"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure IPv6 address groups.
---

# fmgdevice_firewall_addrgrp6
<i>This object will be purged after policy copy and install.</i> Configure IPv6 address groups.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `dynamic_mapping`: `fmgdevice_firewall_addrgrp6_dynamic_mapping`
>- `tagging`: `fmgdevice_firewall_addrgrp6_tagging`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `_image_base64` - _Image-Base64.
* `color` - Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets the value to 1).
* `comment` - Comment.
* `dynamic_mapping` - Dynamic_Mapping. The structure of `dynamic_mapping` block is documented below.
* `exclude` - Enable/disable address6 exclusion. Valid values: `disable`, `enable`.

* `exclude_member` - Address6 exclusion member.
* `fabric_object` - Security Fabric global object setting. Valid values: `disable`, `enable`.

* `global_object` - Global-Object.
* `member` - Address objects contained within the group.
* `name` - IPv6 address group name.
* `tagging` - Tagging. The structure of `tagging` block is documented below.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dynamic_mapping` block supports:

* `_image_base64` - _Image-Base64.
* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `color` - Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets the value to 1).
* `comment` - Comment.
* `exclude` - Enable/disable address6 exclusion. Valid values: `disable`, `enable`.

* `exclude_member` - Address6 exclusion member.
* `fabric_object` - Security Fabric global object setting. Valid values: `disable`, `enable`.

* `global_object` - Global-Object.
* `member` - Address objects contained within the group.
* `tags` - Tags.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `visibility` - Visibility. Valid values: `disable`, `enable`.


The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.

The `tagging` block supports:

* `category` - Tag category.
* `name` - Tagging entry name.
* `tags` - Tags.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall Addrgrp6 can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_addrgrp6.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

