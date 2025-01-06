---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_ipsec_phase2interface"
description: |-
  Configure VPN autokey tunnel.
---

# fmgdevice_vpn_ipsec_phase2interface
Configure VPN autokey tunnel.

## Example Usage

```hcl
resource "fmgdevice_vpn_ipsec_phase2interface" "trname" {
  add_route                = "phase1"
  auto_discovery_forwarder = "enable"
  auto_discovery_sender    = "enable"
  auto_negotiate           = "disable"
  comments                 = "your own value"
  name                     = "your own value"
  device_name              = var.device_name # not required if setting is at provider
  device_vdom              = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `add_route` - Enable/disable automatic route addition. Valid values: `disable`, `enable`, `phase1`.

* `addke1` - phase2 ADDKE1 group. Valid values: `0`, `1080`, `1081`, `1082`.

* `addke2` - phase2 ADDKE2 group. Valid values: `0`, `1080`, `1081`, `1082`.

* `addke3` - phase2 ADDKE3 group. Valid values: `0`, `1080`, `1081`, `1082`.

* `addke4` - phase2 ADDKE4 group. Valid values: `0`, `1080`, `1081`, `1082`.

* `addke5` - phase2 ADDKE5 group. Valid values: `0`, `1080`, `1081`, `1082`.

* `addke6` - phase2 ADDKE6 group. Valid values: `0`, `1080`, `1081`, `1082`.

* `addke7` - phase2 ADDKE7 group. Valid values: `0`, `1080`, `1081`, `1082`.

* `auto_discovery_forwarder` - Enable/disable forwarding short-cut messages. Valid values: `disable`, `enable`, `phase1`.

* `auto_discovery_sender` - Enable/disable sending short-cut messages. Valid values: `disable`, `enable`, `phase1`.

* `auto_negotiate` - Enable/disable IPsec SA auto-negotiation. Valid values: `disable`, `enable`.

* `comments` - Comment.
* `dhcp_ipsec` - Enable/disable DHCP-IPsec. Valid values: `disable`, `enable`.

* `dhgrp` - Phase2 DH group. Valid values: `1`, `2`, `5`, `14`, `15`, `16`, `17`, `18`, `19`, `20`, `21`, `27`, `28`, `29`, `30`, `31`, `32`.

* `diffserv` - Enable/disable applying DSCP value to the IPsec tunnel outer IP header. Valid values: `disable`, `enable`.

* `diffservcode` - DSCP value to be applied to the IPsec tunnel outer IP header.
* `dst_addr_type` - Remote proxy ID type. Valid values: `subnet`, `range`, `ip`, `name`, `subnet6`, `range6`, `ip6`, `name6`.

* `dst_end_ip` - Remote proxy ID IPv4 end.
* `dst_end_ip6` - Remote proxy ID IPv6 end.
* `dst_name` - Remote proxy ID name.
* `dst_name6` - Remote proxy ID name.
* `dst_port` - Quick mode destination port (1 - 65535 or 0 for all).
* `dst_start_ip` - Remote proxy ID IPv4 start.
* `dst_start_ip6` - Remote proxy ID IPv6 start.
* `dst_subnet` - Remote proxy ID IPv4 subnet.
* `dst_subnet6` - Remote proxy ID IPv6 subnet.
* `encapsulation` - ESP encapsulation mode. Valid values: `tunnel-mode`, `transport-mode`.

* `inbound_dscp_copy` - Enable/disable copying of the DSCP in the ESP header to the inner IP header. Valid values: `disable`, `enable`, `phase1`.

* `initiator_ts_narrow` - Enable/disable traffic selector narrowing for IKEv2 initiator. Valid values: `disable`, `enable`.

* `ipv4_df` - Enable/disable setting and resetting of IPv4 'Don't Fragment' bit. Valid values: `disable`, `enable`.

* `keepalive` - Enable/disable keep alive. Valid values: `disable`, `enable`.

* `keylife_type` - Keylife type. Valid values: `seconds`, `kbs`, `both`.

* `keylifekbs` - Phase2 key life in number of kilobytes of traffic (5120 - 4294967295).
* `keylifeseconds` - Phase2 key life in time in seconds (120 - 172800).
* `l2tp` - Enable/disable L2TP over IPsec. Valid values: `disable`, `enable`.

* `name` - IPsec tunnel name.
* `pfs` - Enable/disable PFS feature. Valid values: `disable`, `enable`.

* `phase1name` - Phase 1 determines the options required for phase 2.
* `proposal` - Phase2 proposal. Valid values: `null-md5`, `null-sha1`, `des-null`, `3des-null`, `des-md5`, `des-sha1`, `3des-md5`, `3des-sha1`, `aes128-md5`, `aes128-sha1`, `aes192-md5`, `aes192-sha1`, `aes256-md5`, `aes256-sha1`, `aes128-null`, `aes192-null`, `aes256-null`, `null-sha256`, `des-sha256`, `3des-sha256`, `aes128-sha256`, `aes192-sha256`, `aes256-sha256`, `des-sha384`, `des-sha512`, `3des-sha384`, `3des-sha512`, `aes128-sha384`, `aes128-sha512`, `aes192-sha384`, `aes192-sha512`, `aes256-sha384`, `aes256-sha512`, `null-sha384`, `null-sha512`, `aria128-null`, `aria128-md5`, `aria128-sha1`, `aria128-sha256`, `aria128-sha384`, `aria128-sha512`, `aria192-null`, `aria192-md5`, `aria192-sha1`, `aria192-sha256`, `aria192-sha384`, `aria192-sha512`, `aria256-null`, `aria256-md5`, `aria256-sha1`, `aria256-sha256`, `aria256-sha384`, `aria256-sha512`, `seed-null`, `seed-md5`, `seed-sha1`, `seed-sha256`, `seed-sha384`, `seed-sha512`, `aes128gcm`, `aes256gcm`, `chacha20poly1305`.

* `protocol` - Quick mode protocol selector (1 - 255 or 0 for all).
* `replay` - Enable/disable replay detection. Valid values: `disable`, `enable`.

* `route_overlap` - Action for overlapping routes. Valid values: `use-old`, `use-new`, `allow`.

* `single_source` - Enable/disable single source IP restriction. Valid values: `disable`, `enable`.

* `src_addr_type` - Local proxy ID type. Valid values: `subnet`, `range`, `ip`, `name`, `subnet6`, `range6`, `ip6`, `name6`.

* `src_end_ip` - Local proxy ID end.
* `src_end_ip6` - Local proxy ID IPv6 end.
* `src_name` - Local proxy ID name.
* `src_name6` - Local proxy ID name.
* `src_port` - Quick mode source port (1 - 65535 or 0 for all).
* `src_start_ip` - Local proxy ID start.
* `src_start_ip6` - Local proxy ID IPv6 start.
* `src_subnet` - Local proxy ID subnet.
* `src_subnet6` - Local proxy ID IPv6 subnet.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Vpn IpsecPhase2Interface can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_ipsec_phase2interface.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

