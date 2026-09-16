---
subcategory: "Log"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_log_customformat_logtemplates"
description: |-
  Custom log templates.
---

# fmgdevice_log_customformat_logtemplates
Custom log templates.

~> This resource is a sub resource for variable `log_templates` of resource `fmgdevice_log_customformat`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `custom_format` - Custom Format.

* `category` - Log category. Valid values: `traffic`, `event`, `virus`, `webfilter`, `attack`, `spam`, `voip`, `dlp`, `app-ctrl`, `anomaly`, `waf`, `gtp`, `dns`, `ssh`, `ssl`, `file-filter`, `icap`, `virtual-patch`, `debug`.

* `name` - Template name string.
* `subtypes` - Log subtypes to apply template to.
* `template` - Log template string.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Log CustomFormatLogTemplates can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "custom_format=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_log_customformat_logtemplates.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

