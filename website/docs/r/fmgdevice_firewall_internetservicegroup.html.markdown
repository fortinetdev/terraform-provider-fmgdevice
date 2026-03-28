---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_internetservicegroup"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure group of Internet Service.
---

# fmgdevice_firewall_internetservicegroup
<i>This object will be purged after policy copy and install.</i> Configure group of Internet Service.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comment` - Comment.
* `direction` - How this service may be used (source, destination or both). Valid values: `both`, `source`, `destination`.

* `member` - Internet Service group member.
* `name` - Internet Service group name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall InternetServiceGroup can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_internetservicegroup.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

