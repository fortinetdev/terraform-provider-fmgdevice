---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_hotspot20_anqpvenuename"
description: |-
  Configure venue name duple.
---

# fmgdevice_wirelesscontroller_hotspot20_anqpvenuename
Configure venue name duple.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `value_list`: `fmgdevice_wirelesscontroller_hotspot20_anqpvenuename_valuelist`



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_hotspot20_anqpvenuename" "trname" {
  name = "your own value"
  value_list {
    index = 10
    lang  = "your own value"
    value = "your own value"
  }

  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `name` - Name of venue name duple.
* `value_list` - Value-List. The structure of `value_list` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `value_list` block supports:

* `index` - Value index.
* `lang` - Language code.
* `value` - Venue name value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController Hotspot20AnqpVenueName can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_hotspot20_anqpvenuename.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

