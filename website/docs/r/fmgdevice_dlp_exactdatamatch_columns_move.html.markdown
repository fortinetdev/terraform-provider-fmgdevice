---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_dlp_exactdatamatch_columns_move"
description: |-
  DLP exact-data-match column types.
---

# fmgdevice_dlp_exactdatamatch_columns_move
DLP exact-data-match column types.

## Example Usage

```hcl
resource "fmgdevice_dlp_exactdatamatch_columns_move" "trname" {
  exact_data_match = fmgdevice_dlp_exactdatamatch.trname1.name
  columns          = 11
  target           = 10
  option           = "before"
  device_name      = var.device_name # not required if setting is at provider
  device_vdom      = var.device_vdom # not required if setting is at provider
  depends_on       = [fmgdevice_dlp_exactdatamatch_columns.trname1, fmgdevice_dlp_exactdatamatch_columns.trname2]
}

resource "fmgdevice_dlp_exactdatamatch_columns" "trname2" {
  index            = 11
  exact_data_match = fmgdevice_dlp_exactdatamatch.trname1.name
  device_name      = var.device_name # not required if setting is at provider
  device_vdom      = var.device_vdom # not required if setting is at provider
  depends_on       = [fmgdevice_dlp_exactdatamatch.trname1]
}


resource "fmgdevice_dlp_exactdatamatch_columns" "trname1" {
  index            = 10
  exact_data_match = fmgdevice_dlp_exactdatamatch.trname1.name
  device_name      = var.device_name # not required if setting is at provider
  device_vdom      = var.device_vdom # not required if setting is at provider
  depends_on       = [fmgdevice_dlp_exactdatamatch.trname1]
}

resource "fmgdevice_dlp_exactdatamatch" "trname1" {
  name        = "test2"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `exact_data_match` - Exact Data Match.
* `columns` - Columns.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{index}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the columns changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of column.
