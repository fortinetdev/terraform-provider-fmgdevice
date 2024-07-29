---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_hotspot20_icon_iconlist"
description: |-
  Icon list.
---

# fmgdevice_wirelesscontroller_hotspot20_icon_iconlist
Icon list.

~> This resource is a sub resource for variable `icon_list` of resource `fmgdevice_wirelesscontroller_hotspot20_icon`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_hotspot20_icon_iconlist" "trname" {
  file        = "your own value"
  height      = 10
  lang        = "your own value"
  name        = "your own value"
  type        = "png"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `icon` - Icon.

* `file` - Icon file.
* `height` - Icon height.
* `lang` - Language code.
* `name` - Icon name.
* `type` - Icon type. Valid values: `bmp`, `gif`, `jpeg`, `png`, `tiff`.

* `width` - Icon width.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController Hotspot20IconIconList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "icon=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_hotspot20_icon_iconlist.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

