---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_log_memory_setting"
description: |-
  Settings for memory buffer.
---

# fmgdevice_log_memory_setting
Settings for memory buffer.

## Example Usage

```hcl
resource "fmgdevice_log_memory_setting" "trname" {
  diskfull    = "overwrite"
  status      = "enable"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `diskfull` - Diskfull. Valid values: `overwrite`, `blocktraffic`.

* `status` - Enable/disable logging to the FortiGate's memory. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Log MemorySetting can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_log_memory_setting.labelname LogMemorySetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

