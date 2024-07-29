---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_hotspot20_anqproamingconsortium"
description: |-
  Configure roaming consortium.
---

# fmgdevice_wirelesscontroller_hotspot20_anqproamingconsortium
Configure roaming consortium.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `oi_list`: `fmgdevice_wirelesscontroller_hotspot20_anqproamingconsortium_oilist`



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_hotspot20_anqproamingconsortium" "trname" {
  name = "your own value"
  oi_list {
    comment = "your own value"
    index   = 10
    oi      = "your own value"
  }

  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `name` - Roaming consortium name.
* `oi_list` - Oi-List. The structure of `oi_list` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `oi_list` block supports:

* `comment` - Comment.
* `index` - OI index.
* `oi` - Organization identifier.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController Hotspot20AnqpRoamingConsortium can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_hotspot20_anqproamingconsortium.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

