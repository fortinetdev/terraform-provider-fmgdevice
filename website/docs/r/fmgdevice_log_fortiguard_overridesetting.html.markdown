---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_log_fortiguard_overridesetting"
description: |-
  Override global FortiCloud logging settings for this VDOM.
---

# fmgdevice_log_fortiguard_overridesetting
Override global FortiCloud logging settings for this VDOM.

## Example Usage

```hcl
resource "fmgdevice_log_fortiguard_overridesetting" "trname" {
  access_config = "enable"
  max_log_rate  = 10
  override      = "disable"
  priority      = "low"
  status        = "disable"
  device_name   = var.device_name # not required if setting is at provider
  device_vdom   = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `access_config` - Enable/disable FortiCloud access to configuration and data. Valid values: `disable`, `enable`.

* `max_log_rate` - FortiCloud maximum log rate in MBps (0 = unlimited).
* `override` - Overriding FortiCloud settings for this VDOM or use global settings. Valid values: `disable`, `enable`.

* `priority` - Set log transmission priority. Valid values: `low`, `default`.

* `status` - Enable/disable logging to FortiCloud. Valid values: `disable`, `enable`.

* `upload_day` - Day of week to roll logs.
* `upload_interval` - Frequency of uploading log files to FortiCloud. Valid values: `daily`, `weekly`, `monthly`.

* `upload_option` - Configure how log messages are sent to FortiCloud. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.

* `upload_time` - Time of day to roll logs (hh:mm).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Log FortiguardOverrideSetting can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_log_fortiguard_overridesetting.labelname LogFortiguardOverrideSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

