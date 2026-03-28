---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_log_syslogd_overridesetting_logtemplates"
description: |-
  Device LogSyslogdOverrideSettingLogTemplates
---

# fmgdevice_log_syslogd_overridesetting_logtemplates
Device LogSyslogdOverrideSettingLogTemplates

~> This resource is a sub resource for variable `log_templates` of resource `fmgdevice_log_syslogd_overridesetting`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `category` - Category. Valid values: `app-ctrl`, `attack`, `dlp`, `event`, `traffic`, `virus`, `voip`, `webfilter`, `spam`, `anomaly`, `waf`, `dns`, `ssh`, `ssl`, `file-filter`, `icap`, `virtual-patch`.

* `empty_value_indicator` - Empty-Value-Indicator.
* `fosid` - Id.
* `template` - Template.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Log SyslogdOverrideSettingLogTemplates can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_log_syslogd_overridesetting_logtemplates.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

