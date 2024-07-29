---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_hotspot20_h2qpadviceofcharge"
description: |-
  Configure advice of charge.
---

# fmgdevice_wirelesscontroller_hotspot20_h2qpadviceofcharge
Configure advice of charge.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `aoc_list`: `fmgdevice_wirelesscontroller_hotspot20_h2qpadviceofcharge_aoclist`



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_hotspot20_h2qpadviceofcharge" "trname" {
  aoc_list {
    nai_realm          = "your own value"
    nai_realm_encoding = "your own value"
    name               = "your own value"
    plan_info {
      currency  = "your own value"
      info_file = "your own value"
      lang      = "your own value"
      name      = "your own value"
    }

    type = "unlimited"
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

* `aoc_list` - Aoc-List. The structure of `aoc_list` block is documented below.
* `name` - Plan name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `aoc_list` block supports:

* `nai_realm` - NAI realm list name.
* `nai_realm_encoding` - NAI realm encoding.
* `name` - Advice of charge ID.
* `plan_info` - Plan-Info. The structure of `plan_info` block is documented below.
* `type` - Usage charge type. Valid values: `time-based`, `volume-based`, `time-and-volume-based`, `unlimited`.


The `plan_info` block supports:

* `currency` - Currency code.
* `info_file` - Info file.
* `lang` - Language code.
* `name` - Plan name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController Hotspot20H2QpAdviceOfCharge can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_hotspot20_h2qpadviceofcharge.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

