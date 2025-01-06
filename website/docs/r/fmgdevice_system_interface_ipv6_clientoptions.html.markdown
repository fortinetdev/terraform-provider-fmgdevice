---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_interface_ipv6_clientoptions"
description: |-
  DHCP6 client options.
---

# fmgdevice_system_interface_ipv6_clientoptions
DHCP6 client options.

~> This resource is a sub resource for variable `client_options` of resource `fmgdevice_system_interface_ipv6`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `interface` - Interface.

* `code` - DHCPv6 option code.
* `fosid` - ID.
* `ip6` - DHCP option IP6s.
* `type` - DHCPv6 option type. Valid values: `hex`, `string`, `ip6`, `fqdn`.

* `value` - DHCPv6 option value (hexadecimal value must be even).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System InterfaceIpv6ClientOptions can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "interface=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_interface_ipv6_clientoptions.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

