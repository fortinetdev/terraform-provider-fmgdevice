---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_bgp_neighbor_conditionaladvertise"
description: |-
  Conditional advertisement.
---

# fmgdevice_router_bgp_neighbor_conditionaladvertise
Conditional advertisement.

~> This resource is a sub resource for variable `conditional_advertise` of resource `fmgdevice_router_bgp_neighbor`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_router_bgp_neighbor_conditionaladvertise" "trname" {
  advertise_routemap = "your own value"
  condition_routemap = ["your own value"]
  condition_type     = "non-exist"
  device_name        = var.device_name # not required if setting is at provider
  device_vdom        = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `neighbor` - Neighbor.

* `advertise_routemap` - Name of advertising route map.
* `condition_routemap` - List of conditional route maps.
* `condition_type` - Type of condition. Valid values: `exist`, `non-exist`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{advertise_routemap}}.

## Import

Router BgpNeighborConditionalAdvertise can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "neighbor=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_bgp_neighbor_conditionaladvertise.labelname {{advertise_routemap}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

