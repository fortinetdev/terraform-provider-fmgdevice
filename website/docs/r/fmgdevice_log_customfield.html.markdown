---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_log_customfield"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure custom log fields.
---

# fmgdevice_log_customfield
<i>This object will be purged after policy copy and install.</i> Configure custom log fields.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `fosid` - Field ID string.
* `name` - Field name (max: 15 characters).
* `value` - Field value (max: 15 characters).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Log CustomField can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_log_customfield.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

