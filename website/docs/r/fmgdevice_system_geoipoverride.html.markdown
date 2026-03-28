---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_geoipoverride"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure geographical location mapping for IP address(es) to override mappings from FortiGuard.
---

# fmgdevice_system_geoipoverride
<i>This object will be purged after policy copy and install.</i> Configure geographical location mapping for IP address(es) to override mappings from FortiGuard.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `ip6_range`: `fmgdevice_system_geoipoverride_ip6range`
>- `ip_range`: `fmgdevice_system_geoipoverride_iprange`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `country_id` - Two character Country ID code.
* `description` - Description.
* `ip_range` - Ip-Range. The structure of `ip_range` block is documented below.
* `ip6_range` - Ip6-Range. The structure of `ip6_range` block is documented below.
* `name` - Location name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `ip_range` block supports:

* `end_ip` - Ending IP address, inclusive, of the address range (format: xxx.xxx.xxx.xxx).
* `id` - ID of individual entry in the IP range table.
* `start_ip` - Starting IP address, inclusive, of the address range (format: xxx.xxx.xxx.xxx).

The `ip6_range` block supports:

* `end_ip` - Ending IP address, inclusive, of the address range (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
* `id` - ID of individual entry in the IPv6 range table.
* `start_ip` - Starting IP address, inclusive, of the address range (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System GeoipOverride can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_geoipoverride.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

