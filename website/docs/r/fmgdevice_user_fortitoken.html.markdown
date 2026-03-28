---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_user_fortitoken"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure FortiToken.
---

# fmgdevice_user_fortitoken
<i>This object will be purged after policy copy and install.</i> Configure FortiToken.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comments` - Comment.
* `license` - Mobile token license.
* `serial_number` - Serial number.
* `status` - Status. Valid values: `lock`, `active`.

* `os_ver` - Device Mobile Version.
* `reg_id` - Device Reg ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{serial_number}}.

## Import

User Fortitoken can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_user_fortitoken.labelname {{serial_number}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

