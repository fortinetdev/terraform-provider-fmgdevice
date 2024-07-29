---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_automation_setting"
description: |-
  Automation setting configuration.
---

# fmgdevice_automation_setting
Automation setting configuration.

## Example Usage

```hcl
resource "fmgdevice_automation_setting" "trname" {
  fabric_sync             = "disable"
  max_concurrent_stitches = 10
  device_name             = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `fabric_sync` - Enable/disable synchronization of automation settings with security fabric. Valid values: `disable`, `enable`.

* `max_concurrent_stitches` - Maximum number of automation stitches that are allowed to run concurrently.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Automation Setting can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_automation_setting.labelname AutomationSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

