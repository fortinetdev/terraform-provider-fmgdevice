---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_dhcp6_server_options"
description: |-
  DHCPv6 options.
---

# fmgdevice_system_dhcp6_server_options
DHCPv6 options.

~> This resource is a sub resource for variable `options` of resource `fmgdevice_system_dhcp6_server`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `server` - Server.

* `code` - DHCPv6 option code.
* `fosid` - ID.
* `ip6` - DHCP option IP6s.
* `type` - DHCPv6 option type. Valid values: `hex`, `string`, `ip6`, `fqdn`.

* `value` - DHCPv6 option value (hexadecimal value must be even).
* `vci_match` - Enable/disable vendor class option matching. When enabled only DHCP requests with a matching VCI are served with this option. Valid values: `disable`, `enable`.

* `vci_string` - One or more VCI strings in quotes separated by spaces.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System Dhcp6ServerOptions can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "server=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_dhcp6_server_options.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

