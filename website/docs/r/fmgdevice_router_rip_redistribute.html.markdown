---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_rip_redistribute"
description: |-
  Redistribute configuration.
---

# fmgdevice_router_rip_redistribute
Redistribute configuration.

~> This resource is a sub resource for variable `redistribute` of resource `fmgdevice_router_rip`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_router_rip_redistribute" "trname" {
  metric      = 10
  name        = "your own value"
  routemap    = ["your own value"]
  status      = "disable"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `metric` - Redistribute metric setting.
* `name` - Redistribute name.
* `routemap` - Route map name.
* `status` - Status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router RipRedistribute can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_rip_redistribute.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

