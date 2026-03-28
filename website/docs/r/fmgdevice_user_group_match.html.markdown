---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_user_group_match"
description: |-
  <i>This object will be purged after policy copy and install.</i> Group matches.
---

# fmgdevice_user_group_match
<i>This object will be purged after policy copy and install.</i> Group matches.

~> This resource is a sub resource for variable `match` of resource `fmgdevice_user_group`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `group` - Group.

* `_gui_meta` - _Gui_Meta.
* `group_name` - Name of matching user or group on remote authentication server or SCIM.
* `fosid` - ID.
* `server_name` - Name of remote auth server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

User GroupMatch can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "group=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_user_group_match.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

