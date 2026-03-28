---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_replacemsgimage"
description: |-
  Configure replacement message images.
---

# fmgdevice_system_replacemsgimage
Configure replacement message images.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `image_base64` - Image data.
* `image_type` - Image type. Valid values: `gif`, `jpg`, `tiff`, `png`.

* `name` - Image name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System ReplacemsgImage can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_replacemsgimage.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

