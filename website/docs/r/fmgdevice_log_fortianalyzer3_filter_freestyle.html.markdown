---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_log_fortianalyzer3_filter_freestyle"
description: |-
  Free style filters.
---

# fmgdevice_log_fortianalyzer3_filter_freestyle
Free style filters.

~> This resource is a sub resource for variable `free_style` of resource `fmgdevice_log_fortianalyzer3_filter`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_log_fortianalyzer3_filter_freestyle" "trname" {
  category    = "anomaly"
  filter      = "your own value"
  filter_type = "include"
  fosid       = 10
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `category` - Log category. Valid values: `traffic`, `event`, `virus`, `webfilter`, `attack`, `spam`, `voip`, `dlp`, `app-ctrl`, `anomaly`, `waf`, `gtp`, `dns`, `ssh`, `ssl`, `file-filter`, `icap`, `ztna`, `virtual-patch`.

* `filter` - Free style filter string.
* `filter_type` - Include/exclude logs that match the filter. Valid values: `include`, `exclude`.

* `fosid` - Entry ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Log Fortianalyzer3FilterFreeStyle can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_log_fortianalyzer3_filter_freestyle.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

