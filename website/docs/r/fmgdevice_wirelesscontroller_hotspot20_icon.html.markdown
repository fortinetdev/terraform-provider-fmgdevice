---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_hotspot20_icon"
description: |-
  Configure OSU provider icon.
---

# fmgdevice_wirelesscontroller_hotspot20_icon
Configure OSU provider icon.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `icon_list`: `fmgdevice_wirelesscontroller_hotspot20_icon_iconlist`



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_hotspot20_icon" "trname" {
  icon_list {
    file   = "your own value"
    height = 10
    lang   = "your own value"
    name   = "your own value"
    type   = "bmp"
    width  = 10
  }

  name        = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `icon_list` - Icon-List. The structure of `icon_list` block is documented below.
* `name` - Icon list ID.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `icon_list` block supports:

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

WirelessController Hotspot20Icon can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_hotspot20_icon.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

