---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_hotspot20_anqp3gppcellular"
description: |-
  Configure 3GPP public land mobile network (PLMN).
---

# fmgdevice_wirelesscontroller_hotspot20_anqp3gppcellular
Configure 3GPP public land mobile network (PLMN).

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `mcc_mnc_list`: `fmgdevice_wirelesscontroller_hotspot20_anqp3gppcellular_mccmnclist`



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_hotspot20_anqp3gppcellular" "trname" {
  mcc_mnc_list {
    fosid = 10
    mcc   = "your own value"
    mnc   = "your own value"
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

* `mcc_mnc_list` - Mcc-Mnc-List. The structure of `mcc_mnc_list` block is documented below.
* `name` - 3GPP PLMN name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `mcc_mnc_list` block supports:

* `id` - ID.
* `mcc` - Mobile country code.
* `mnc` - Mobile network code.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController Hotspot20Anqp3GppCellular can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_hotspot20_anqp3gppcellular.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

