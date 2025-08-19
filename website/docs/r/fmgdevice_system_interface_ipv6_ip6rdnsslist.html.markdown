---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_interface_ipv6_ip6rdnsslist"
description: |-
  Advertised IPv6 RDNSS list.
---

# fmgdevice_system_interface_ipv6_ip6rdnsslist
Advertised IPv6 RDNSS list.

~> This resource is a sub resource for variable `ip6_rdnss_list` of resource `fmgdevice_system_interface_ipv6`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `interface` - Interface.

* `rdnss` - Recursive DNS server option.
* `rdnss_life_time` - Recursive DNS server life time in seconds (0 - 4294967295, default = 1800).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{rdnss}}.

## Import

System InterfaceIpv6Ip6RdnssList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "interface=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_interface_ipv6_ip6rdnsslist.labelname {{rdnss}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

