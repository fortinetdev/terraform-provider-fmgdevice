---
subcategory: "Log"
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

* `fabric_force_sync` - Enable/disable forced synchronization of configuration objects from the root FortiGate unit to the downstream devices.  Configuration conflict check is skipped. Valid values: `disable`, `enable`.

* `fabric_object` - Security Fabric global object setting. Valid values: `disable`, `enable`.

* `fabric_object_source` - Source of truth for fabric object. Valid values: `member`, `local`, `root`.

* `fosid` - Field ID string.
* `name` - Field name (max: 15 characters).
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
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

