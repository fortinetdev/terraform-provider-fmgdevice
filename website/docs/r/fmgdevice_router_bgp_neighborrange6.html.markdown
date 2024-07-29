---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_bgp_neighborrange6"
description: |-
  BGP IPv6 neighbor range table.
---

# fmgdevice_router_bgp_neighborrange6
BGP IPv6 neighbor range table.

~> This resource is a sub resource for variable `neighbor_range6` of resource `fmgdevice_router_bgp`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_router_bgp_neighborrange6" "trname" {
  fosid            = 10
  max_neighbor_num = 10
  neighbor_group   = ["your own value"]
  prefix6          = "your own value"
  device_name      = var.device_name # not required if setting is at provider
  device_vdom      = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `fosid` - IPv6 neighbor range ID.
* `max_neighbor_num` - Maximum number of neighbors.
* `neighbor_group` - Neighbor group name.
* `prefix6` - IPv6 prefix.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Router BgpNeighborRange6 can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_bgp_neighborrange6.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

