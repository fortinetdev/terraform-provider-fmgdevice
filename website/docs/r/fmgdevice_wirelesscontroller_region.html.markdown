---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_region"
description: |-
  Configure FortiAP regions (for floor plans and maps).
---

# fmgdevice_wirelesscontroller_region
Configure FortiAP regions (for floor plans and maps).

## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_region" "trname" {
  comments    = "your own value"
  grayscale   = "disable"
  image_type  = "png"
  name        = "your own value"
  opacity     = 10
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comments` - Comments.
* `grayscale` - Region image grayscale. Valid values: `disable`, `enable`.

* `image_type` - FortiAP region image type (png|jpeg|gif). Valid values: `gif`, `jpeg`, `png`.

* `name` - FortiAP region name.
* `opacity` - Region image opacity (0 - 100).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController Region can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_region.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

