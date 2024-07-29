---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_ospf_redistribute"
description: |-
  Redistribute configuration.
---

# fmgdevice_router_ospf_redistribute
Redistribute configuration.

~> This resource is a sub resource for variable `redistribute` of resource `fmgdevice_router_ospf`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_router_ospf_redistribute" "trname" {
  metric      = 10
  metric_type = "1"
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
* `metric_type` - Metric type. Valid values: `2`, `1`.

* `name` - Redistribute name.
* `routemap` - Route map name.
* `status` - Status. Valid values: `disable`, `enable`.

* `tag` - Tag value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router OspfRedistribute can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_ospf_redistribute.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

