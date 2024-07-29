---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_loadbalance_setting_workers"
description: |-
  Worker blade used by this group.
---

# fmgdevice_loadbalance_setting_workers
Worker blade used by this group.

~> This resource is a sub resource for variable `workers` of resource `fmgdevice_loadbalance_setting`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_loadbalance_setting_workers" "trname" {
  slot        = 10
  status      = "enable"
  weight      = 10
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `slot` - slot number
* `status` - Enable/disable this worker. Valid values: `disable`, `enable`.

* `weight` - load balancing weight (1-10)


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{slot}}.

## Import

LoadBalance SettingWorkers can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_loadbalance_setting_workers.labelname {{slot}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

