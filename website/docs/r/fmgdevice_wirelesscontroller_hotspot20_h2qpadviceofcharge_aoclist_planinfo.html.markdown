---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_hotspot20_h2qpadviceofcharge_aoclist_planinfo"
description: |-
  Plan info.
---

# fmgdevice_wirelesscontroller_hotspot20_h2qpadviceofcharge_aoclist_planinfo
Plan info.

~> This resource is a sub resource for variable `plan_info` of resource `fmgdevice_wirelesscontroller_hotspot20_h2qpadviceofcharge_aoclist`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_hotspot20_h2qpadviceofcharge_aoclist_planinfo" "trname" {
  currency    = "your own value"
  info_file   = "your own value"
  lang        = "your own value"
  name        = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `h2qp_advice_of_charge` - H2Qp Advice Of Charge.
* `aoc_list` - Aoc List.

* `currency` - Currency code.
* `info_file` - Info file.
* `lang` - Language code.
* `name` - Plan name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController Hotspot20H2QpAdviceOfChargeAocListPlanInfo can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "h2qp_advice_of_charge=YOUR_VALUE", "aoc_list=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_hotspot20_h2qpadviceofcharge_aoclist_planinfo.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

