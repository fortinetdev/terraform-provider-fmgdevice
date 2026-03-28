---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_address6"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure IPv6 firewall addresses.
---

# fmgdevice_firewall_address6
<i>This object will be purged after policy copy and install.</i> Configure IPv6 firewall addresses.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `dynamic_mapping`: `fmgdevice_firewall_address6_dynamic_mapping`
>- `list`: `fmgdevice_firewall_address6_list`
>- `subnet_segment`: `fmgdevice_firewall_address6_subnetsegment`
>- `tagging`: `fmgdevice_firewall_address6_tagging`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `_image_base64` - _Image-Base64.
* `cache_ttl` - Minimal TTL of individual IPv6 addresses in FQDN cache.
* `color` - Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
* `comment` - Comment.
* `country` - IPv6 addresses associated to a specific country.
* `dynamic_mapping` - Dynamic_Mapping. The structure of `dynamic_mapping` block is documented below.
* `end_ip` - Final IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
* `end_mac` - Last MAC address in the range.
* `epg_name` - Endpoint group name.
* `fabric_object` - Security Fabric global object setting. Valid values: `disable`, `enable`.

* `filter` - Match criteria filter.
* `fqdn` - Fully qualified domain name.
* `global_object` - Global-Object.
* `host` - Host Address.
* `host_type` - Host type. Valid values: `any`, `specific`.

* `ip6` - IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
* `list` - List. The structure of `list` block is documented below.
* `macaddr` - Multiple MAC address ranges.
* `name` - Address name.
* `obj_id` - Object ID for NSX.
* `route_tag` - route-tag address.
* `sdn` - SDN.
* `sdn_addr_type` - Type of addresses to collect. Valid values: `all`, `private`, `public`.

* `sdn_tag` - SDN Tag.
* `start_ip` - First IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
* `start_mac` - First MAC address in the range.
* `subnet_segment` - Subnet-Segment. The structure of `subnet_segment` block is documented below.
* `tagging` - Tagging. The structure of `tagging` block is documented below.
* `template` - IPv6 address template.
* `tenant` - Tenant.
* `type` - Type of IPv6 address object (default = ipprefix). Valid values: `ipprefix`, `iprange`, `dynamic`, `fqdn`, `template`, `mac`, `geography`, `route-tag`, `wildcard`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `tags` - Tags.
* `wildcard` - IPv6 address and wildcard netmask.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dynamic_mapping` block supports:

* `_image_base64` - _Image-Base64.
* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `cache_ttl` - Minimal TTL of individual IPv6 addresses in FQDN cache.
* `color` - Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
* `comment` - Comment.
* `country` - IPv6 addresses associated to a specific country.
* `end_ip` - Final IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
* `end_mac` - End-Mac.
* `epg_name` - Endpoint group name.
* `fabric_object` - Security Fabric global object setting. Valid values: `disable`, `enable`.

* `filter` - Match criteria filter.
* `fqdn` - Fully qualified domain name.
* `global_object` - Global-Object.
* `host` - Host Address.
* `host_type` - Host type. Valid values: `any`, `specific`.

* `ip6` - IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
* `macaddr` - Multiple MAC address ranges.
* `obj_id` - Object ID for NSX.
* `route_tag` - route-tag address.
* `sdn` - SDN.
* `sdn_addr_type` - Type of addresses to collect. Valid values: `all`, `private`, `public`.

* `sdn_tag` - SDN Tag.
* `start_ip` - First IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
* `start_mac` - Start-Mac.
* `subnet_segment` - Subnet-Segment. The structure of `subnet_segment` block is documented below.
* `tags` - Tags.
* `template` - IPv6 address template.
* `tenant` - Tenant.
* `type` - Type of IPv6 address object (default = ipprefix). Valid values: `ipprefix`, `iprange`, `dynamic`, `fqdn`, `template`, `mac`, `geography`, `route-tag`, `wildcard`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `visibility` - Visibility. Valid values: `disable`, `enable`.

* `wildcard` - IPv6 address and wildcard netmask.

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.

The `subnet_segment` block supports:

* `name` - Name.
* `type` - Subnet segment type. Valid values: `any`, `specific`.

* `value` - Subnet segment value.

The `list` block supports:

* `ip` - IP.
* `net_id` - Network ID.
* `obj_id` - Object ID.

The `subnet_segment` block supports:

* `name` - Name.
* `type` - Subnet segment type. Valid values: `any`, `specific`.

* `value` - Subnet segment value.

The `tagging` block supports:

* `category` - Tag category.
* `name` - Tagging entry name.
* `tags` - Tags.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall Address6 can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_address6.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

