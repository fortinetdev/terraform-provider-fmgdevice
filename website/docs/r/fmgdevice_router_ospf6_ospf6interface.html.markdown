---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_ospf6_ospf6interface"
description: |-
  OSPF6 interface configuration.
---

# fmgdevice_router_ospf6_ospf6interface
OSPF6 interface configuration.

~> This resource is a sub resource for variable `ospf6_interface` of resource `fmgdevice_router_ospf6`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `ipsec_keys`: `fmgdevice_router_ospf6_ospf6interface_ipseckeys`
>- `neighbor`: `fmgdevice_router_ospf6_ospf6interface_neighbor`



## Example Usage

```hcl
resource "fmgdevice_router_ospf6_ospf6interface" "trname" {
  area_id        = "your own value"
  authentication = "none"
  bfd            = "disable"
  cost           = 10
  dead_interval  = 10
  name           = "your own value"
  device_name    = var.device_name # not required if setting is at provider
  device_vdom    = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `area_id` - A.B.C.D, in IPv4 address format.
* `authentication` - Authentication mode. Valid values: `none`, `ah`, `esp`, `area`.

* `bfd` - Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `global`, `enable`, `disable`.

* `cost` - Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `interface` - Configuration interface name.
* `ipsec_auth_alg` - Authentication algorithm. Valid values: `md5`, `sha1`, `sha256`, `sha384`, `sha512`.

* `ipsec_enc_alg` - Encryption algorithm. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`.

* `ipsec_keys` - Ipsec-Keys. The structure of `ipsec_keys` block is documented below.
* `key_rollover_interval` - Key roll-over interval.
* `mtu` - MTU for OSPFv3 packets.
* `mtu_ignore` - Enable/disable ignoring MTU field in DBD packets. Valid values: `disable`, `enable`.

* `name` - Interface entry name.
* `neighbor` - Neighbor. The structure of `neighbor` block is documented below.
* `network_type` - Network type. Valid values: `broadcast`, `non-broadcast`, `point-to-point`, `point-to-multipoint`, `point-to-multipoint-non-broadcast`.

* `priority` - Priority.
* `retransmit_interval` - Retransmit interval.
* `status` - Enable/disable OSPF6 routing on this interface. Valid values: `disable`, `enable`.

* `transmit_delay` - Transmit delay.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `ipsec_keys` block supports:

* `auth_key` - Authentication key.
* `enc_key` - Encryption key.
* `spi` - Security Parameters Index.

The `neighbor` block supports:

* `cost` - Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
* `ip6` - IPv6 link local address of the neighbor.
* `poll_interval` - Poll interval time in seconds.
* `priority` - Priority.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router Ospf6Ospf6Interface can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_ospf6_ospf6interface.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

