---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_hotspot20_h2qpoperatorname"
description: |-
  Configure operator friendly name.
---

# fmgdevice_wirelesscontroller_hotspot20_h2qpoperatorname
Configure operator friendly name.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `value_list`: `fmgdevice_wirelesscontroller_hotspot20_h2qpoperatorname_valuelist`



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_hotspot20_h2qpoperatorname" "trname" {
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

* `name` - Friendly name ID.
* `value_list` - Value-List. The structure of `value_list` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `value_list` block supports:

* `index` - Value index.
* `lang` - Language code.
* `value` - Friendly name value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController Hotspot20H2QpOperatorName can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_hotspot20_h2qpoperatorname.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

