---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_centralsnatmap"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure IPv4 and IPv6 central SNAT policies.
---

# fmgdevice_firewall_centralsnatmap
<i>This object will be purged after policy copy and install.</i> Configure IPv4 and IPv6 central SNAT policies.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comments` - Comment.
* `dst_addr` - IPv4 Destination address.
* `dst_addr6` - IPv6 Destination address.
* `dst_port` - Destination port or port range (1 to 65535, 0 means any port).
* `dstintf` - Destination interface name from available interfaces.
* `nat` - Enable/disable source NAT. Valid values: `disable`, `enable`.

* `nat_ippool` - Name of the IP pools to be used to translate addresses from available IP Pools.
* `nat_ippool6` - IPv6 pools to be used for source NAT.
* `nat_port` - Translated port or port range (1 to 65535, 0 means any port).
* `nat46` - Enable/disable NAT46. Valid values: `disable`, `enable`.

* `nat64` - Enable/disable NAT64. Valid values: `disable`, `enable`.

* `orig_addr` - IPv4 Original address.
* `orig_addr6` - IPv6 Original address.
* `orig_port` - Original TCP port (1 to 65535, 0 means any port).
* `policyid` - Policy ID.
* `port_preserve` - Enable/disable preservation of the original source port from source NAT if it has not been used. Valid values: `disable`, `enable`.

* `port_random` - Enable/disable random source port selection for source NAT. Valid values: `disable`, `enable`.

* `protocol` - Integer value for the protocol type (0 - 255).
* `srcintf` - Source interface name from available interfaces.
* `status` - Enable/disable the active status of this policy. Valid values: `disable`, `enable`.

* `type` - IPv4/IPv6 source NAT. Valid values: `ipv4`, `ipv6`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `action` - Action. Valid values: `bypass`, `masquerade`, `ippool`.

* `ipv6` - Ipv6. Valid values: `disable`, `enable`.

* `src_addr` - Src-Addr.
* `src_addr6` - Src-Addr6.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Firewall CentralSnatMap can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_centralsnatmap.labelname {{policyid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

