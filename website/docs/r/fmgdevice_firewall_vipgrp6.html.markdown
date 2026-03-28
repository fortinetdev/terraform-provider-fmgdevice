---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_vipgrp6"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure IPv6 virtual IP groups.
---

# fmgdevice_firewall_vipgrp6
<i>This object will be purged after policy copy and install.</i> Configure IPv6 virtual IP groups.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `color` - Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
* `comments` - Comment.
* `member` - Member VIP objects of the group (Separate multiple objects with a space).
* `name` - IPv6 VIP group name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall Vipgrp6 can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_vipgrp6.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

