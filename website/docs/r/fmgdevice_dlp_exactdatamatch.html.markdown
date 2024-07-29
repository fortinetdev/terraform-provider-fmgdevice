---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_dlp_exactdatamatch"
description: |-
  Configure exact-data-match template used by DLP scan.
---

# fmgdevice_dlp_exactdatamatch
Configure exact-data-match template used by DLP scan.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `columns`: `fmgdevice_dlp_exactdatamatch_columns`



## Example Usage

```hcl
resource "fmgdevice_dlp_exactdatamatch" "trname" {
  columns {
    index    = 10
    optional = "disable"
    type     = ["your own value"]
  }

  data        = ["your own value"]
  name        = "your own value"
  optional    = 10
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `columns` - Columns. The structure of `columns` block is documented below.
* `data` - External resource for exact data match.
* `name` - Name of table containing the exact-data-match template.
* `optional` - Number of optional columns need to match.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `columns` block supports:

* `index` - Column index.
* `optional` - Enable/disable optional match. Valid values: `disable`, `enable`.

* `type` - Data-type for this column.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Dlp ExactDataMatch can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_dlp_exactdatamatch.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

