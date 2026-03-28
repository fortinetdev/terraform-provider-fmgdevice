---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webfilter_ftgdlocalcat"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure FortiGuard Web Filter local categories.
---

# fmgdevice_webfilter_ftgdlocalcat
<i>This object will be purged after policy copy and install.</i> Configure FortiGuard Web Filter local categories.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `desc` - Local category description.
* `fosid` - Local category ID.
* `status` - Enable/disable the local category. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{desc}}.

## Import

Webfilter FtgdLocalCat can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webfilter_ftgdlocalcat.labelname {{desc}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

