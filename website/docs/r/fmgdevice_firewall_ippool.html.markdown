---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_ippool"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure IPv4 IP pools.
---

# fmgdevice_firewall_ippool
<i>This object will be purged after policy copy and install.</i> Configure IPv4 IP pools.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `dynamic_mapping`: `fmgdevice_firewall_ippool_dynamic_mapping`
>- `grp`: `fmgdevice_firewall_ippool_grp`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `add_nat64_route` - Enable/disable adding NAT64 route. Valid values: `disable`, `enable`.

* `arp_intf` - Select an interface from available options that will reply to ARP requests. (If blank, any is selected).
* `arp_reply` - Enable/disable replying to ARP requests when an IP Pool is added to a policy (default = enable). Valid values: `disable`, `enable`.

* `associated_interface` - Associated interface name.
* `block_size` - Number of addresses in a block (64 - 4096, default = 128).
* `cgn_block_size` - Number of ports in a block(64 to 4096 in unit of 64, default = 128).
* `cgn_client_endip` - Final client IPv4 address (inclusive) (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `cgn_client_ipv6shift` - IPv6 shift for fixed-allocation.(default 0)
* `cgn_client_startip` - First client IPv4 address (inclusive) (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `cgn_fixedalloc` - Enable/disable fixed-allocation mode. Valid values: `disable`, `enable`.

* `cgn_overload` - Enable/disable overload mode. Valid values: `disable`, `enable`.

* `cgn_port_end` - Ending public port can be allocated.
* `cgn_port_start` - Starting public port can be allocated.
* `cgn_spa` - Enable/disable single port allocation mode. Valid values: `disable`, `enable`.

* `client_prefix_length` - Subnet length of a single deterministic NAT64 client (1 - 128, default = 64).
* `comments` - Comment.
* `dynamic_mapping` - Dynamic_Mapping. The structure of `dynamic_mapping` block is documented below.
* `endip` - Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `endport` - Final port number (inclusive) in the range for the address pool (1024 - 65535, Default: 65533).
* `exclude_ip` - Exclude IPs x.x.x.x.
* `icmp_session_quota` - Maximum number of concurrent ICMP sessions allowed per client (0 - 2097000, default = 0 which means no limit).
* `name` - IP pool name.
* `nat64` - Enable/disable NAT64. Valid values: `disable`, `enable`.

* `num_blocks_per_user` - Number of addresses blocks that can be used by a user (1 to 128, default = 8).
* `pba_interim_log` - Port block allocation interim logging interval (600 - 86400 seconds, default = 0 which disables interim logging).
* `pba_timeout` - Port block allocation timeout (seconds).
* `permit_any_host` - Enable/disable full cone NAT. Valid values: `disable`, `enable`.

* `port_per_user` - Number of port for each user (32 - 60416, default = 0, which is auto).
* `privileged_port_use_pba` - Enable/disable selection of the external port from the port block allocation for NAT'ing privileged ports (deafult = disable). Valid values: `disable`, `enable`.

* `source_endip` - Final IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `source_prefix6` - Source IPv6 network to be translated (format = xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx, default = ::/0).
* `source_startip` - First IPv4 address (inclusive) in the range of the source addresses to be translated (format = xxx.xxx.xxx.xxx, default = 0.0.0.0).
* `startip` - First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `startport` - First port number (inclusive) in the range for the address pool (1024 - 65535, Default: 5117).
* `subnet_broadcast_in_ippool` - Enable/disable inclusion of the subnetwork address and broadcast IP address in the NAT64 IP pool. Valid values: `disable`, `enable`.

* `tcp_session_quota` - Maximum number of concurrent TCP sessions allowed per client (0 - 2097000, default = 0 which means no limit).
* `type` - IP pool type: overload, one-to-one, fixed-port-range, port-block-allocation, cgn-resource-allocation (hyperscale vdom only) Valid values: `overload`, `one-to-one`, `fixed-port-range`, `port-block-allocation`, `cgn-resource-allocation`.

* `udp_session_quota` - Maximum number of concurrent UDP sessions allowed per client (0 - 2097000, default = 0 which means no limit).
* `utilization_alarm_clear` - Pool utilization alarm clear threshold (40-100).
* `utilization_alarm_raise` - Pool utilization alarm raise threshold (50-100).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dynamic_mapping` block supports:

* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `add_nat64_route` - Enable/disable adding NAT64 route. Valid values: `disable`, `enable`.

* `arp_intf` - Select an interface from available options that will reply to ARP requests. (If blank, any is selected).
* `arp_reply` - Enable/disable replying to ARP requests when an IP Pool is added to a policy (default = enable). Valid values: `disable`, `enable`.

* `associated_interface` - Associated interface name.
* `block_size` - Number of addresses in a block (64 - 4096, default = 128).
* `cgn_block_size` - Number of ports in a block(64 to 4096 in unit of 64, default = 128).
* `cgn_client_endip` - Final client IPv4 address (inclusive) (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `cgn_client_ipv6shift` - IPv6 shift for fixed-allocation.(default 0)
* `cgn_client_startip` - First client IPv4 address (inclusive) (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `cgn_fixedalloc` - Enable/disable fixed-allocation mode. Valid values: `disable`, `enable`.

* `cgn_overload` - Enable/disable overload mode. Valid values: `disable`, `enable`.

* `cgn_port_end` - Ending public port can be allocated.
* `cgn_port_start` - Starting public port can be allocated.
* `cgn_spa` - Enable/disable single port allocation mode. Valid values: `disable`, `enable`.

* `client_prefix_length` - Subnet length of a single deterministic NAT64 client (1 - 128, default = 64).
* `comments` - Comment.
* `endip` - Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `endport` - Final port number (inclusive) in the range for the address pool (1024 - 65535, Default: 65533).
* `exclude_ip` - Exclude IPs x.x.x.x.
* `icmp_session_quota` - Maximum number of concurrent ICMP sessions allowed per client (0 - 2097000, default = 0 which means no limit).
* `nat64` - Enable/disable NAT64. Valid values: `disable`, `enable`.

* `num_blocks_per_user` - Number of addresses blocks that can be used by a user (1 to 128, default = 8).
* `pba_interim_log` - Port block allocation interim logging interval (600 - 86400 seconds, default = 0 which disables interim logging).
* `pba_timeout` - Port block allocation timeout (seconds).
* `permit_any_host` - Enable/disable full cone NAT. Valid values: `disable`, `enable`.

* `port_per_user` - Number of port for each user (32 - 60416, default = 0, which is auto).
* `privileged_port_use_pba` - Enable/disable selection of the external port from the port block allocation for NAT'ing privileged ports (deafult = disable). Valid values: `disable`, `enable`.

* `source_endip` - Final IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `source_prefix6` - Source IPv6 network to be translated (format = xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx, default = ::/0).
* `source_startip` - First IPv4 address (inclusive) in the range of the source addresses to be translated (format = xxx.xxx.xxx.xxx, default = 0.0.0.0).
* `startip` - First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `startport` - First port number (inclusive) in the range for the address pool (1024 - 65535, Default: 5117).
* `subnet_broadcast_in_ippool` - Enable/disable inclusion of the subnetwork address and broadcast IP address in the NAT64 IP pool. Valid values: `disable`, `enable`.

* `tcp_session_quota` - Maximum number of concurrent TCP sessions allowed per client (0 - 2097000, default = 0 which means no limit).
* `type` - IP pool type: overload, one-to-one, fixed-port-range, port-block-allocation, cgn-resource-allocation (hyperscale vdom only) Valid values: `overload`, `one-to-one`, `fixed-port-range`, `port-block-allocation`, `cgn-resource-allocation`.

* `udp_session_quota` - Maximum number of concurrent UDP sessions allowed per client (0 - 2097000, default = 0 which means no limit).
* `utilization_alarm_clear` - Pool utilization alarm clear threshold (40-100).
* `utilization_alarm_raise` - Pool utilization alarm raise threshold (50-100).

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall Ippool can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_ippool.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

