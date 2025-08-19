---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_interface_ipv6_ip6dnssllist"
description: |-
  Advertised IPv6 DNSS list.
---

# fmgdevice_system_interface_ipv6_ip6dnssllist
Advertised IPv6 DNSS list.

~> This resource is a sub resource for variable `ip6_dnssl_list` of resource `fmgdevice_system_interface_ipv6`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `interface` - Interface.

* `dnssl_life_time` - DNS search list time in seconds (0 - 4294967295, default = 1800).
* `domain` - Domain name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{domain}}.

## Import

System InterfaceIpv6Ip6DnsslList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "interface=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_interface_ipv6_ip6dnssllist.labelname {{domain}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

