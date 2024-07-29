---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_autoupdate_schedule"
description: |-
  Configure update schedule.
---

# fmgdevice_system_autoupdate_schedule
Configure update schedule.

## Example Usage

```hcl
resource "fmgdevice_system_autoupdate_schedule" "trname" {
  day         = "Thursday"
  frequency   = "daily"
  status      = "enable"
  time        = "your own value"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `day` - Update day. Valid values: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.

* `frequency` - Update frequency. Valid values: `every`, `daily`, `weekly`, `automatic`.

* `status` - Enable/disable scheduled updates. Valid values: `disable`, `enable`.

* `time` - Update time.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System AutoupdateSchedule can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_autoupdate_schedule.labelname SystemAutoupdateSchedule
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

