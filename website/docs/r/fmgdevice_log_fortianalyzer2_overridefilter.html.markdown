---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_log_fortianalyzer2_overridefilter"
description: |-
  Override filters for FortiAnalyzer.
---

# fmgdevice_log_fortianalyzer2_overridefilter
Override filters for FortiAnalyzer.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `free_style`: `fmgdevice_log_fortianalyzer2_overridefilter_freestyle`



## Example Usage

```hcl
resource "fmgdevice_log_fortianalyzer2_overridefilter" "trname" {
  anomaly      = "disable"
  dlp_archive  = "disable"
  filter       = "your own value"
  filter_type  = "exclude"
  forti_switch = "disable"
  device_name  = var.device_name # not required if setting is at provider
  device_vdom  = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `anomaly` - Enable/disable anomaly logging. Valid values: `disable`, `enable`.

* `dlp_archive` - Enable/disable DLP archive logging. Valid values: `disable`, `enable`.

* `filter` - FortiAnalyzer 2 log filter.
* `filter_type` - Include/exclude logs that match the filter. Valid values: `include`, `exclude`.

* `forti_switch` - Enable/disable Forti-Switch logging. Valid values: `disable`, `enable`.

* `forward_traffic` - Enable/disable forward traffic logging. Valid values: `disable`, `enable`.

* `free_style` - Free-Style. The structure of `free_style` block is documented below.
* `gtp` - Enable/disable GTP messages logging. Valid values: `disable`, `enable`.

* `http_transaction` - Enable/disable log HTTP transaction messages. Valid values: `disable`, `enable`.

* `local_traffic` - Enable/disable local in or out traffic logging. Valid values: `disable`, `enable`.

* `multicast_traffic` - Enable/disable multicast traffic logging. Valid values: `disable`, `enable`.

* `severity` - Log every message above and including this severity level. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.

* `sniffer_traffic` - Enable/disable sniffer traffic logging. Valid values: `disable`, `enable`.

* `voip` - Enable/disable VoIP logging. Valid values: `disable`, `enable`.

* `ztna_traffic` - Enable/disable ztna traffic logging. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `free_style` block supports:

* `category` - Log category. Valid values: `traffic`, `event`, `virus`, `webfilter`, `attack`, `spam`, `voip`, `dlp`, `app-ctrl`, `anomaly`, `waf`, `gtp`, `dns`, `ssh`, `ssl`, `file-filter`, `icap`, `ztna`, `virtual-patch`.

* `filter` - Free style filter string.
* `filter_type` - Include/exclude logs that match the filter. Valid values: `include`, `exclude`.

* `id` - Entry ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Log Fortianalyzer2OverrideFilter can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_log_fortianalyzer2_overridefilter.labelname LogFortianalyzer2OverrideFilter
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

