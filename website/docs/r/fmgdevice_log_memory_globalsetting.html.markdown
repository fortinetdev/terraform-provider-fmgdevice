---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_log_memory_globalsetting"
description: |-
  Global settings for memory logging.
---

# fmgdevice_log_memory_globalsetting
Global settings for memory logging.

## Example Usage

```hcl
resource "fmgdevice_log_memory_globalsetting" "trname" {
  full_final_warning_threshold  = 10
  full_first_warning_threshold  = 10
  full_second_warning_threshold = 10
  max_size                      = 10
  device_name                   = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `full_final_warning_threshold` - Log full final warning threshold as a percent (3 - 100, default = 95).
* `full_first_warning_threshold` - Log full first warning threshold as a percent (1 - 98, default = 75).
* `full_second_warning_threshold` - Log full second warning threshold as a percent (2 - 99, default = 90).
* `max_size` - Maximum amount of memory that can be used for memory logging in bytes.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Log MemoryGlobalSetting can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_log_memory_globalsetting.labelname LogMemoryGlobalSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

