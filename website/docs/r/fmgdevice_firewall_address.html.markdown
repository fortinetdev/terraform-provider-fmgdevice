---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_address"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure IPv4 addresses.
---

# fmgdevice_firewall_address
<i>This object will be purged after policy copy and install.</i> Configure IPv4 addresses.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `dynamic_mapping`: `fmgdevice_firewall_address_dynamic_mapping`
>- `list`: `fmgdevice_firewall_address_list`
>- `tagging`: `fmgdevice_firewall_address_tagging`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `_image_base64` - _Image-Base64.
* `agent_id` - Telemetry agent id.
* `allow_routing` - Enable/disable use of this address in routing configurations. Valid values: `disable`, `enable`.

* `associated_interface` - Network interface associated with address.
* `cache_ttl` - Defines the minimal TTL of individual IP addresses in FQDN cache measured in seconds.
* `clearpass_spt` - SPT (System Posture Token) value. Valid values: `unknown`, `healthy`, `quarantine`, `checkup`, `infected`, `transient`.

* `color` - Color of icon on the GUI.
* `comment` - Comment.
* `country` - IP addresses associated to a specific country.
* `dirty` - To be deleted address. Valid values: `dirty`, `clean`.

* `dynamic_mapping` - Dynamic_Mapping. The structure of `dynamic_mapping` block is documented below.
* `end_ip` - Final IP address (inclusive) in the range for the address.
* `end_mac` - Last MAC address in the range.
* `epg_name` - Endpoint group name.
* `fabric_object` - Security Fabric global object setting. Valid values: `disable`, `enable`.

* `filter` - Match criteria filter.
* `fqdn` - Fully Qualified Domain Name address.
* `fsso_group` - FSSO group(s).
* `global_object` - Global-Object.
* `hw_model` - Dynamic address matching hardware model.
* `hw_vendor` - Dynamic address matching hardware vendor.
* `interface` - Name of interface whose IP address is to be used.
* `list` - List. The structure of `list` block is documented below.
* `macaddr` - Multiple MAC address ranges.
* `name` - Address name.
* `node_ip_only` - Enable/disable collection of node addresses only in Kubernetes. Valid values: `disable`, `enable`.

* `obj_id` - Object ID for NSX.
* `obj_tag` - Tag of dynamic address object.
* `obj_type` - Object type. Valid values: `ip`, `mac`.

* `organization` - Organization domain name (Syntax: organization/domain).
* `os` - Dynamic address matching operating system.
* `policy_group` - Policy group name.
* `route_tag` - route-tag address.
* `sdn` - SDN.
* `sdn_addr_type` - Type of addresses to collect. Valid values: `private`, `public`, `all`.

* `sdn_tag` - SDN Tag.
* `sso_attribute_value` - RADIUS attributes value.
* `start_ip` - First IP address (inclusive) in the range for the address.
* `start_mac` - First MAC address in the range.
* `sub_type` - Sub-type of address. Valid values: `sdn`, `clearpass-spt`, `fsso`, `ems-tag`, `swc-tag`, `fortivoice-tag`, `fortinac-tag`, `fortipolicy-tag`, `device-identification`, `rsso`, `external-resource`, `obsolete`, `telemetry`.

* `subnet` - IP address and subnet mask of address.
* `subnet_name` - Subnet name.
* `sw_version` - Dynamic address matching software version.
* `tag_detection_level` - Tag detection level of dynamic address object.
* `tag_type` - Tag type of dynamic address object.
* `tag_uuid` - Foreign UUID of dynamic address object.
* `tagging` - Tagging. The structure of `tagging` block is documented below.
* `tenant` - Tenant.
* `type` - Type of address. Valid values: `ipmask`, `iprange`, `fqdn`, `wildcard`, `geography`, `dynamic`, `interface-subnet`, `mac`, `route-tag`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `wildcard` - IP address and wildcard netmask.
* `wildcard_fqdn` - Fully Qualified Domain Name with wildcard characters.
* `pattern_end` - Pattern-End.
* `pattern_start` - Pattern-Start.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dynamic_mapping` block supports:

* `_image_base64` - _Image-Base64.
* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `agent_id` - Telemetry agent id.
* `allow_routing` - Enable/disable use of this address in routing configurations. Valid values: `disable`, `enable`.

* `associated_interface` - Network interface associated with address.
* `cache_ttl` - Defines the minimal TTL of individual IP addresses in FQDN cache measured in seconds.
* `clearpass_spt` - SPT (System Posture Token) value. Valid values: `unknown`, `healthy`, `quarantine`, `checkup`, `infected`, `transient`.

* `color` - Color of icon on the GUI.
* `comment` - Comment.
* `country` - IP addresses associated to a specific country.
* `dirty` - To be deleted address. Valid values: `dirty`, `clean`.

* `end_ip` - Final IP address (inclusive) in the range for the address.
* `end_mac` - End-Mac.
* `epg_name` - Endpoint group name.
* `fabric_object` - Security Fabric global object setting. Valid values: `disable`, `enable`.

* `filter` - Match criteria filter.
* `fqdn` - Fully Qualified Domain Name address.
* `fsso_group` - FSSO group(s).
* `global_object` - Global-Object.
* `hw_model` - Dynamic address matching hardware model.
* `hw_vendor` - Dynamic address matching hardware vendor.
* `interface` - Name of interface whose IP address is to be used.
* `macaddr` - Multiple MAC address ranges.
* `node_ip_only` - Enable/disable collection of node addresses only in Kubernetes. Valid values: `disable`, `enable`.

* `obj_id` - Object ID for NSX.
* `obj_tag` - Tag of dynamic address object.
* `obj_type` - Object type. Valid values: `ip`, `mac`.

* `organization` - Organization domain name (Syntax: organization/domain).
* `os` - Dynamic address matching operating system.
* `pattern_end` - Pattern-End.
* `pattern_start` - Pattern-Start.
* `policy_group` - Policy group name.
* `route_tag` - route-tag address.
* `sdn` - SDN.
* `sdn_addr_type` - Type of addresses to collect. Valid values: `private`, `public`, `all`.

* `sdn_tag` - SDN Tag.
* `sso_attribute_value` - RADIUS attributes value.
* `start_ip` - First IP address (inclusive) in the range for the address.
* `start_mac` - Start-Mac.
* `sub_type` - Sub-type of address. Valid values: `sdn`, `clearpass-spt`, `fsso`, `ems-tag`, `swc-tag`, `fortivoice-tag`, `fortinac-tag`, `fortipolicy-tag`, `device-identification`, `rsso`, `external-resource`, `obsolete`, `telemetry`.

* `subnet` - IP address and subnet mask of address.
* `subnet_name` - Subnet name.
* `sw_version` - Dynamic address matching software version.
* `tag_detection_level` - Tag detection level of dynamic address object.
* `tag_type` - Tag type of dynamic address object.
* `tag_uuid` - Foreign UUID of dynamic address object.
* `tags` - Tags.
* `tenant` - Tenant.
* `type` - Type of address. Valid values: `ipmask`, `iprange`, `fqdn`, `wildcard`, `geography`, `dynamic`, `interface-subnet`, `mac`, `route-tag`.

* `url` - Url.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `visibility` - Visibility. Valid values: `disable`, `enable`.

* `wildcard` - IP address and wildcard netmask.
* `wildcard_fqdn` - Fully Qualified Domain Name with wildcard characters.

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.

The `list` block supports:

* `ip` - IP.
* `net_id` - Network ID.
* `obj_id` - Object ID.

The `tagging` block supports:

* `category` - Tag category.
* `name` - Tagging entry name.
* `tags` - Tags.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall Address can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_address.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

