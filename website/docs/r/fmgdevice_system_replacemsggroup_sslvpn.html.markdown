---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_replacemsggroup_sslvpn"
description: |-
  <i>This object will be purged after policy copy and install.</i> Replacement message table entries.
---

# fmgdevice_system_replacemsggroup_sslvpn
<i>This object will be purged after policy copy and install.</i> Replacement message table entries.

~> This resource is a sub resource for variable `sslvpn` of resource `fmgdevice_system_replacemsggroup`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `replacemsg_group` - Replacemsg Group.

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{msg_type}}.

## Import

System ReplacemsgGroupSslvpn can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "replacemsg_group=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_replacemsggroup_sslvpn.labelname {{msg_type}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

