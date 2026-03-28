---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_internetservicecustomgroup"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure custom Internet Service group.
---

# fmgdevice_firewall_internetservicecustomgroup
<i>This object will be purged after policy copy and install.</i> Configure custom Internet Service group.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comment` - Comment.
* `member` - Custom Internet Service group members.
* `name` - Custom Internet Service group name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall InternetServiceCustomGroup can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_internetservicecustomgroup.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

