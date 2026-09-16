---
subcategory: "Log"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_log_customformat"
description: |-
  Configure custom log format.
---

# fmgdevice_log_customformat
Configure custom log format.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `log_templates`: `fmgdevice_log_customformat_logtemplates`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `empty_value_indicator` - A character to indicate log field is empty.
* `field_exclusion_list` - Log fields to exclude in the log.
* `log_templates` - Log-Templates. The structure of `log_templates` block is documented below.
* `name` - Format name string.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `log_templates` block supports:

* `category` - Log category. Valid values: `traffic`, `event`, `virus`, `webfilter`, `attack`, `spam`, `voip`, `dlp`, `app-ctrl`, `anomaly`, `waf`, `gtp`, `dns`, `ssh`, `ssl`, `file-filter`, `icap`, `virtual-patch`, `debug`.

* `name` - Template name string.
* `subtypes` - Log subtypes to apply template to.
* `template` - Log template string.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Log CustomFormat can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_log_customformat.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

