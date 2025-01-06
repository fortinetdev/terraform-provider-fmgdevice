---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_dhcp6_server"
description: |-
  Configure DHCPv6 servers.
---

# fmgdevice_system_dhcp6_server
Configure DHCPv6 servers.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `ip_range`: `fmgdevice_system_dhcp6_server_iprange`
>- `options`: `fmgdevice_system_dhcp6_server_options`
>- `prefix_range`: `fmgdevice_system_dhcp6_server_prefixrange`



## Example Usage

```hcl
resource "fmgdevice_system_dhcp6_server" "trname" {
  delegated_prefix_iaid = 10
  dns_search_list       = "delegated"
  dns_server1           = "your own value"
  dns_server2           = "your own value"
  dns_server3           = "your own value"
  fosid                 = 10
  device_name           = var.device_name # not required if setting is at provider
  device_vdom           = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `delegated_prefix_iaid` - IAID of obtained delegated-prefix from the upstream interface.
* `dns_search_list` - DNS search list options. Valid values: `specify`, `delegated`.

* `dns_server1` - DNS server 1.
* `dns_server2` - DNS server 2.
* `dns_server3` - DNS server 3.
* `dns_server4` - DNS server 4.
* `dns_service` - Options for assigning DNS servers to DHCPv6 clients. Valid values: `default`, `specify`, `delegated`.

* `domain` - Domain name suffix for the IP addresses that the DHCP server assigns to clients.
* `fosid` - ID.
* `interface` - DHCP server can assign IP configurations to clients connected to this interface.
* `ip_mode` - Method used to assign client IP. Valid values: `range`, `delegated`.

* `ip_range` - Ip-Range. The structure of `ip_range` block is documented below.
* `lease_time` - Lease time in seconds, 0 means unlimited.
* `options` - Options. The structure of `options` block is documented below.
* `option1` - Option 1.
* `option2` - Option 2.
* `option3` - Option 3.
* `prefix_mode` - Assigning a prefix from a DHCPv6 client or RA. Valid values: `dhcp6`, `ra`.

* `prefix_range` - Prefix-Range. The structure of `prefix_range` block is documented below.
* `rapid_commit` - Enable/disable allow/disallow rapid commit. Valid values: `disable`, `enable`.

* `status` - Enable/disable this DHCPv6 configuration. Valid values: `disable`, `enable`.

* `subnet` - Subnet or subnet-id if the IP mode is delegated.
* `upstream_interface` - Interface name from where delegated information is provided.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `ip_range` block supports:

* `end_ip` - End of IP range.
* `id` - ID.
* `start_ip` - Start of IP range.
* `vci_match` - Enable/disable vendor class option matching. When enabled only DHCP requests with a matching VC are served with this range. Valid values: `disable`, `enable`.

* `vci_string` - One or more VCI strings in quotes separated by spaces.

The `options` block supports:

* `code` - DHCPv6 option code.
* `id` - ID.
* `ip6` - DHCP option IP6s.
* `type` - DHCPv6 option type. Valid values: `hex`, `string`, `ip6`, `fqdn`.

* `value` - DHCPv6 option value (hexadecimal value must be even).
* `vci_match` - Enable/disable vendor class option matching. When enabled only DHCP requests with a matching VCI are served with this option. Valid values: `disable`, `enable`.

* `vci_string` - One or more VCI strings in quotes separated by spaces.

The `prefix_range` block supports:

* `end_prefix` - End of prefix range.
* `id` - ID.
* `prefix_length` - Prefix length.
* `start_prefix` - Start of prefix range.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System Dhcp6Server can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_dhcp6_server.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

