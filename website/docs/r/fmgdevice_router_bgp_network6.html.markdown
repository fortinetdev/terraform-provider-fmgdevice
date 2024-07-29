---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_bgp_network6"
description: |-
  BGP IPv6 network table.
---

# fmgdevice_router_bgp_network6
BGP IPv6 network table.

~> This resource is a sub resource for variable `network6` of resource `fmgdevice_router_bgp`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_router_bgp_network6" "trname" {
  backdoor             = "enable"
  fosid                = 10
  network_import_check = "global"
  prefix6              = "your own value"
  route_map            = ["your own value"]
  device_name          = var.device_name # not required if setting is at provider
  device_vdom          = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `backdoor` - Enable/disable route as backdoor. Valid values: `disable`, `enable`.

* `fosid` - ID.
* `network_import_check` - Configure insurance of BGP network route existence in IGP. Valid values: `disable`, `enable`, `global`.

* `prefix6` - Network IPv6 prefix.
* `route_map` - Route map to modify generated route.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Router BgpNetwork6 can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_bgp_network6.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

