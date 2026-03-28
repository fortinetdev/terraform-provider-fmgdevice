---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_service_group"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure service groups.
---

# fmgdevice_firewall_service_group
<i>This object will be purged after policy copy and install.</i> Configure service groups.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `color` - Color of icon on the GUI.
* `comment` - Comment.
* `fabric_object` - Security Fabric global object setting. Valid values: `disable`, `enable`.

* `global_object` - Global-Object.
* `member` - Service objects contained within the group.
* `name` - Service group name.
* `proxy` - Enable/disable web proxy service group. Valid values: `disable`, `enable`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall ServiceGroup can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_service_group.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

