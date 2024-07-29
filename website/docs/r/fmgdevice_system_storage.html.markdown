---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_storage"
description: |-
  Configure logical storage.
---

# fmgdevice_system_storage
Configure logical storage.

## Example Usage

```hcl
resource "fmgdevice_system_storage" "trname" {
  device       = "your own value"
  media_status = "disable"
  name         = "your own value"
  order        = 10
  partition    = "your own value"
  device_name  = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `device` - Partition device.
* `media_status` - The physical status of current media. Valid values: `disable`, `enable`, `fail`.

* `name` - Storage name.
* `order` - Set storage order.
* `partition` - Label of underlying partition.
* `size` - Partition size.
* `status` - Enable/disable storage. Valid values: `disable`, `enable`.

* `usage` - Use hard disk for logging and WAN Optimization. Valid values: `log`, `wanopt`, `mix`, `unused`.

* `wanopt_mode` - WAN Optimization mode (default = mix). Valid values: `mix`, `wanopt`, `webcache`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Storage can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_storage.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

