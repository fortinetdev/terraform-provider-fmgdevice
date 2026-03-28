---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_ippool6"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure IPv6 IP pools.
---

# fmgdevice_firewall_ippool6
<i>This object will be purged after policy copy and install.</i> Configure IPv6 IP pools.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `dynamic_mapping`: `fmgdevice_firewall_ippool6_dynamic_mapping`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `add_nat46_route` - Enable/disable adding NAT46 route. Valid values: `disable`, `enable`.

* `comments` - Comment.
* `dynamic_mapping` - Dynamic_Mapping. The structure of `dynamic_mapping` block is documented below.
* `endip` - Final IPv6 address (inclusive) in the range for the address pool (format = xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, default = ::).
* `external_prefix` - External NPTv6 prefix length (32 - 64).
* `internal_prefix` - Internal NPTv6 prefix length (32 - 64).
* `name` - IPv6 IP pool name.
* `nat46` - Enable/disable NAT46. Valid values: `disable`, `enable`.

* `startip` - First IPv6 address (inclusive) in the range for the address pool (format = xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, default = ::).
* `type` - Configure IPv6 pool type (overload or NPTv6). Valid values: `overload`, `nptv6`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dynamic_mapping` block supports:

* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `add_nat46_route` - Enable/disable adding NAT46 route. Valid values: `disable`, `enable`.

* `comments` - Comment.
* `endip` - Final IPv6 address (inclusive) in the range for the address pool (format = xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, default = ::).
* `external_prefix` - External NPTv6 prefix length (32 - 64).
* `internal_prefix` - Internal NPTv6 prefix length (32 - 64).
* `nat46` - Enable/disable NAT46. Valid values: `disable`, `enable`.

* `startip` - First IPv6 address (inclusive) in the range for the address pool (format = xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, default = ::).
* `type` - Configure IPv6 pool type (overload or NPTv6). Valid values: `overload`, `nptv6`.


The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall Ippool6 can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_ippool6.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

