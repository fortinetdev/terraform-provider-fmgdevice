---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_replacemsggroup_devicedetectionportal"
description: |-
  <i>This object will be purged after policy copy and install.</i>
---

# fmgdevice_system_replacemsggroup_devicedetectionportal
<i>This object will be purged after policy copy and install.</i>

~> This resource is a sub resource for variable `device_detection_portal` of resource `fmgdevice_system_replacemsggroup`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `replacemsg_group` - Replacemsg Group.

* `buffer` - Buffer.
* `format` - Format. Valid values: `none`, `text`, `html`.

* `header` - Header. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Msg-Type.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{msg_type}}.

## Import

System ReplacemsgGroupDeviceDetectionPortal can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "replacemsg_group=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_replacemsggroup_devicedetectionportal.labelname {{msg_type}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

