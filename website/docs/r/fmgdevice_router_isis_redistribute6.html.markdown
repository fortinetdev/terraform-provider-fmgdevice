---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_isis_redistribute6"
description: |-
  IS-IS IPv6 redistribution for routing protocols.
---

# fmgdevice_router_isis_redistribute6
IS-IS IPv6 redistribution for routing protocols.

~> This resource is a sub resource for variable `redistribute6` of resource `fmgdevice_router_isis`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_router_isis_redistribute6" "trname" {
  level       = "level-1-2"
  metric      = 10
  metric_type = "external"
  protocol    = "your own value"
  routemap    = ["your own value"]
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `level` - Level. Valid values: `level-1-2`, `level-1`, `level-2`.

* `metric` - Metric.
* `metric_type` - Metric type. Valid values: `external`, `internal`.

* `protocol` - Protocol name.
* `routemap` - Route map name.
* `status` - Enable/disable redistribution. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{protocol}}.

## Import

Router IsisRedistribute6 can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_isis_redistribute6.labelname {{protocol}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

