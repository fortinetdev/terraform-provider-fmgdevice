---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_log_syslogd3_overridesetting_customfieldname"
description: |-
  Custom field name for CEF format logging.
---

# fmgdevice_log_syslogd3_overridesetting_customfieldname
Custom field name for CEF format logging.

~> This resource is a sub resource for variable `custom_field_name` of resource `fmgdevice_log_syslogd3_overridesetting`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_log_syslogd3_overridesetting_customfieldname" "trname" {
  custom      = "your own value"
  fosid       = 10
  name        = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `custom` - Field custom name [A-Za-z0-9_].
* `fosid` - Entry ID.
* `name` - Field name [A-Za-z0-9_].


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Log Syslogd3OverrideSettingCustomFieldName can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_log_syslogd3_overridesetting_customfieldname.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

