---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_proxyaddrgrp6"
description: |-
  <i>This object will be purged after policy copy and install.</i>
---

# fmgdevice_firewall_proxyaddrgrp6
<i>This object will be purged after policy copy and install.</i>

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `tagging`: `fmgdevice_firewall_proxyaddrgrp6_tagging`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `color` - Color.
* `comment` - Comment.
* `member` - Member.
* `name` - Name.
* `tagging` - Tagging. The structure of `tagging` block is documented below.
* `type` - Type. Valid values: `src`, `dst`.

* `uuid` - Uuid.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `tagging` block supports:

* `category` - Category.
* `name` - Name.
* `tags` - Tags.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall ProxyAddrgrp6 can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_proxyaddrgrp6.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

