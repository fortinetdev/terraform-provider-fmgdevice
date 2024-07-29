---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_bfd6_neighbor"
description: |-
  Configure neighbor of IPv6 BFD.
---

# fmgdevice_router_bfd6_neighbor
Configure neighbor of IPv6 BFD.

~> This resource is a sub resource for variable `neighbor` of resource `fmgdevice_router_bfd6`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_router_bfd6_neighbor" "trname" {
  interface   = ["port2"]
  ip6_address = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `interface` - Interface to the BFD neighbor.
* `ip6_address` - IPv6 address of the BFD neighbor.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{ip6_address}}.

## Import

Router Bfd6Neighbor can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_bfd6_neighbor.labelname {{ip6_address}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

