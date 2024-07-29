---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_dlp_exactdatamatch_columns"
description: |-
  DLP exact-data-match column types.
---

# fmgdevice_dlp_exactdatamatch_columns
DLP exact-data-match column types.

~> This resource is a sub resource for variable `columns` of resource `fmgdevice_dlp_exactdatamatch`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_dlp_exactdatamatch_columns" "trname" {
  index       = 10
  optional    = "enable"
  type        = ["your own value"]
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `exact_data_match` - Exact Data Match.

* `index` - Column index.
* `optional` - Enable/disable optional match. Valid values: `disable`, `enable`.

* `type` - Data-type for this column.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{index}}.

## Import

Dlp ExactDataMatchColumns can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "exact_data_match=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_dlp_exactdatamatch_columns.labelname {{index}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

