---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_nsxt_setting"
description: |-
  Configure NSX-T setting.
---

# fmgdevice_nsxt_setting
Configure NSX-T setting.

## Example Usage

```hcl
resource "fmgdevice_nsxt_setting" "trname" {
  liveness    = "disable"
  service     = "your own value"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `liveness` - Enable/disable liveness detection packet forwarding. Valid values: `disable`, `enable`.

* `service` - Service name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Nsxt Setting can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_nsxt_setting.labelname NsxtSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

