---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_ospf6_ospf6interface_neighbor"
description: |-
  OSPFv3 neighbors are used when OSPFv3 runs on non-broadcast media.
---

# fmgdevice_router_ospf6_ospf6interface_neighbor
OSPFv3 neighbors are used when OSPFv3 runs on non-broadcast media.

~> This resource is a sub resource for variable `neighbor` of resource `fmgdevice_router_ospf6_ospf6interface`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_router_ospf6_ospf6interface_neighbor" "trname" {
  cost          = 10
  ip6           = "your own value"
  poll_interval = 10
  priority      = 10
  device_name   = var.device_name # not required if setting is at provider
  device_vdom   = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `ospf6_interface` - Ospf6 Interface.

* `cost` - Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
* `ip6` - IPv6 link local address of the neighbor.
* `poll_interval` - Poll interval time in seconds.
* `priority` - Priority.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{ip6}}.

## Import

Router Ospf6Ospf6InterfaceNeighbor can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "ospf6_interface=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_ospf6_ospf6interface_neighbor.labelname {{ip6}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

