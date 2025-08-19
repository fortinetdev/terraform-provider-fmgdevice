---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_interface_ipv6_ip6routelist"
description: |-
  Advertised route list.
---

# fmgdevice_system_interface_ipv6_ip6routelist
Advertised route list.

~> This resource is a sub resource for variable `ip6_route_list` of resource `fmgdevice_system_interface_ipv6`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `interface` - Interface.

* `route` - IPv6 route.
* `route_life_time` - Route life time in seconds (0 - 65535, default = 1800).
* `route_pref` - Set route preference to the interface (default = medium). Valid values: `medium`, `high`, `low`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{route}}.

## Import

System InterfaceIpv6Ip6RouteList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "interface=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_interface_ipv6_ip6routelist.labelname {{route}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

