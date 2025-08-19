---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_interface_ipv6"
description: |-
  IPv6 of interface.
---

# fmgdevice_system_interface_ipv6
IPv6 of interface.

~> This resource is a sub resource for variable `ipv6` of resource `fmgdevice_system_interface`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `client_options`: `fmgdevice_system_interface_ipv6_clientoptions`
>- `dhcp6_iapd_list`: `fmgdevice_system_interface_ipv6_dhcp6iapdlist`
>- `ip6_delegated_prefix_list`: `fmgdevice_system_interface_ipv6_ip6delegatedprefixlist`
>- `ip6_dnssl_list`: `fmgdevice_system_interface_ipv6_ip6dnssllist`
>- `ip6_extra_addr`: `fmgdevice_system_interface_ipv6_ip6extraaddr`
>- `ip6_prefix_list`: `fmgdevice_system_interface_ipv6_ip6prefixlist`
>- `ip6_rdnss_list`: `fmgdevice_system_interface_ipv6_ip6rdnsslist`
>- `ip6_route_list`: `fmgdevice_system_interface_ipv6_ip6routelist`
>- `vrrp6`: `fmgdevice_system_interface_ipv6_vrrp6`



## Example Usage

```hcl
resource "fmgdevice_system_interface_ipv6" "trname" {
  autoconf             = "enable"
  cli_conn6_status     = 10
  dhcp6_client_options = ["iana"]
  dhcp6_iapd_list {
    iaid            = 10
    prefix_hint     = "your own value"
    prefix_hint_plt = 10
    prefix_hint_vlt = 10
  }

  dhcp6_information_request = "enable"
  device_name               = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `interface` - Interface.

* `autoconf` - Enable/disable address auto config. Valid values: `disable`, `enable`.

* `cli_conn6_status` - Cli-Conn6-Status.
* `client_options` - Client-Options. The structure of `client_options` block is documented below.
* `dhcp6_client_options` - Dhcp6-Client-Options. Valid values: `rapid`, `iapd`, `iana`, `dns`, `dnsname`.

* `dhcp6_iapd_list` - Dhcp6-Iapd-List. The structure of `dhcp6_iapd_list` block is documented below.
* `dhcp6_information_request` - Enable/disable DHCPv6 information request. Valid values: `disable`, `enable`.

* `dhcp6_prefix_delegation` - Enable/disable DHCPv6 prefix delegation. Valid values: `disable`, `enable`.

* `dhcp6_prefix_hint` - DHCPv6 prefix that will be used as a hint to the upstream DHCPv6 server.
* `dhcp6_prefix_hint_plt` - DHCPv6 prefix hint preferred life time (sec), 0 means unlimited lease time.
* `dhcp6_prefix_hint_vlt` - DHCPv6 prefix hint valid life time (sec).
* `dhcp6_relay_interface_id` - DHCP6 relay interface ID.
* `dhcp6_relay_ip` - DHCPv6 relay IP address.
* `dhcp6_relay_service` - Enable/disable DHCPv6 relay. Valid values: `disable`, `enable`.

* `dhcp6_relay_source_interface` - Enable/disable use of address on this interface as the source address of the relay message. Valid values: `disable`, `enable`.

* `dhcp6_relay_source_ip` - IPv6 address used by the DHCP6 relay as its source IP.
* `dhcp6_relay_type` - DHCPv6 relay type. Valid values: `regular`.

* `icmp6_send_redirect` - Enable/disable sending of ICMPv6 redirects. Valid values: `disable`, `enable`.

* `interface_identifier` - IPv6 interface identifier.
* `ip6_address` - Primary IPv6 address prefix. Syntax: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx.
* `ip6_adv_rio` - Enable/disable sending advertisements with route information option. Valid values: `disable`, `enable`.

* `ip6_allowaccess` - Allow management access to the interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `fgfm`, `capwap`, `fabric`.

* `ip6_default_life` - Default life (sec).
* `ip6_delegated_prefix_iaid` - IAID of obtained delegated-prefix from the upstream interface.
* `ip6_delegated_prefix_list` - Ip6-Delegated-Prefix-List. The structure of `ip6_delegated_prefix_list` block is documented below.
* `ip6_dns_server_override` - Enable/disable using the DNS server acquired by DHCP. Valid values: `disable`, `enable`.

* `ip6_dnssl_list` - Ip6-Dnssl-List. The structure of `ip6_dnssl_list` block is documented below.
* `ip6_extra_addr` - Ip6-Extra-Addr. The structure of `ip6_extra_addr` block is documented below.
* `ip6_hop_limit` - Hop limit (0 means unspecified).
* `ip6_link_mtu` - IPv6 link MTU.
* `ip6_manage_flag` - Enable/disable the managed flag. Valid values: `disable`, `enable`.

* `ip6_max_interval` - IPv6 maximum interval (4 to 1800 sec).
* `ip6_min_interval` - IPv6 minimum interval (3 to 1350 sec).
* `ip6_mode` - Addressing mode (static, DHCP, delegated). Valid values: `static`, `dhcp`, `pppoe`, `delegated`.

* `ip6_other_flag` - Enable/disable the other IPv6 flag. Valid values: `disable`, `enable`.

* `ip6_prefix_list` - Ip6-Prefix-List. The structure of `ip6_prefix_list` block is documented below.
* `ip6_prefix_mode` - Assigning a prefix from DHCP or RA. Valid values: `dhcp6`, `ra`.

* `ip6_rdnss_list` - Ip6-Rdnss-List. The structure of `ip6_rdnss_list` block is documented below.
* `ip6_reachable_time` - IPv6 reachable time (milliseconds; 0 means unspecified).
* `ip6_retrans_time` - IPv6 retransmit time (milliseconds; 0 means unspecified).
* `ip6_route_list` - Ip6-Route-List. The structure of `ip6_route_list` block is documented below.
* `ip6_route_pref` - Set route preference to the interface (default = medium). Valid values: `medium`, `high`, `low`.

* `ip6_send_adv` - Enable/disable sending advertisements about the interface. Valid values: `disable`, `enable`.

* `ip6_subnet` - Subnet to routing prefix. Syntax: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx.
* `ip6_upstream_interface` - Interface name providing delegated information.
* `nd_cert` - Neighbor discovery certificate.
* `nd_cga_modifier` - Neighbor discovery CGA modifier.
* `nd_mode` - Neighbor discovery mode. Valid values: `basic`, `SEND-compatible`.

* `nd_security_level` - Neighbor discovery security level (0 - 7; 0 = least secure, default = 0).
* `nd_timestamp_delta` - Neighbor discovery timestamp delta value (1 - 3600 sec; default = 300).
* `nd_timestamp_fuzz` - Neighbor discovery timestamp fuzz factor (1 - 60 sec; default = 1).
* `ra_send_mtu` - Enable/disable sending link MTU in RA packet. Valid values: `disable`, `enable`.

* `unique_autoconf_addr` - Enable/disable unique auto config address. Valid values: `disable`, `enable`.

* `vrip6_link_local` - Link-local IPv6 address of virtual router.
* `vrrp_virtual_mac6` - Enable/disable virtual MAC for VRRP. Valid values: `disable`, `enable`.

* `vrrp6` - Vrrp6. The structure of `vrrp6` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `client_options` block supports:

* `code` - DHCPv6 option code.
* `id` - ID.
* `ip6` - DHCP option IP6s.
* `type` - DHCPv6 option type. Valid values: `hex`, `string`, `ip6`, `fqdn`.

* `value` - DHCPv6 option value (hexadecimal value must be even).

The `dhcp6_iapd_list` block supports:

* `iaid` - Identity association identifier.
* `prefix_hint` - DHCPv6 prefix that will be used as a hint to the upstream DHCPv6 server.
* `prefix_hint_plt` - DHCPv6 prefix hint preferred life time (sec), 0 means unlimited lease time.
* `prefix_hint_vlt` - DHCPv6 prefix hint valid life time (sec).

The `ip6_delegated_prefix_list` block supports:

* `autonomous_flag` - Enable/disable the autonomous flag. Valid values: `disable`, `enable`.

* `delegated_prefix_iaid` - IAID of obtained delegated-prefix from the upstream interface.
* `onlink_flag` - Enable/disable the onlink flag. Valid values: `disable`, `enable`.

* `prefix_id` - Prefix ID.
* `rdnss` - Recursive DNS server option.
* `rdnss_service` - Recursive DNS service option. Valid values: `delegated`, `default`, `specify`.

* `subnet` - Add subnet ID to routing prefix.
* `upstream_interface` - Name of the interface that provides delegated information.

The `ip6_dnssl_list` block supports:

* `dnssl_life_time` - DNS search list time in seconds (0 - 4294967295, default = 1800).
* `domain` - Domain name.

The `ip6_extra_addr` block supports:

* `prefix` - IPv6 address prefix.

The `ip6_prefix_list` block supports:

* `autonomous_flag` - Enable/disable the autonomous flag. Valid values: `disable`, `enable`.

* `dnssl` - DNS search list option.
* `onlink_flag` - Enable/disable the onlink flag. Valid values: `disable`, `enable`.

* `preferred_life_time` - Preferred life time (sec).
* `prefix` - IPv6 prefix.
* `rdnss` - Recursive DNS server option.
* `valid_life_time` - Valid life time (sec).

The `ip6_rdnss_list` block supports:

* `rdnss` - Recursive DNS server option.
* `rdnss_life_time` - Recursive DNS server life time in seconds (0 - 4294967295, default = 1800).

The `ip6_route_list` block supports:

* `route` - IPv6 route.
* `route_life_time` - Route life time in seconds (0 - 65535, default = 1800).
* `route_pref` - Set route preference to the interface (default = medium). Valid values: `medium`, `high`, `low`.


The `vrrp6` block supports:

* `accept_mode` - Enable/disable accept mode. Valid values: `disable`, `enable`.

* `adv_interval` - Advertisement interval (1 - 255 seconds).
* `ignore_default_route` - Enable/disable ignoring of default route when checking destination. Valid values: `disable`, `enable`.

* `preempt` - Enable/disable preempt mode. Valid values: `disable`, `enable`.

* `priority` - Priority of the virtual router (1 - 255).
* `start_time` - Startup time (1 - 255 seconds).
* `status` - Enable/disable VRRP. Valid values: `disable`, `enable`.

* `vrdst_priority` - Priority of the virtual router when the virtual router destination becomes unreachable (0 - 254).
* `vrdst6` - Monitor the route to this destination.
* `vrgrp` - VRRP group ID (1 - 65535).
* `vrid` - Virtual router identifier (1 - 255).
* `vrip6` - IPv6 address of the virtual router.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System InterfaceIpv6 can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "interface=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_interface_ipv6.labelname SystemInterfaceIpv6
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

