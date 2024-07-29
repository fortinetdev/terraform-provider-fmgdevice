---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_ipam"
description: |-
  Configure IP address management services.
---

# fmgdevice_system_ipam
Configure IP address management services.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `pools`: `fmgdevice_system_ipam_pools`
>- `rules`: `fmgdevice_system_ipam_rules`



## Example Usage

```hcl
resource "fmgdevice_system_ipam" "trname" {
  pool_subnet                    = ["your own value"]
  automatic_conflict_resolution  = "enable"
  manage_lan_addresses           = "disable"
  manage_lan_extension_addresses = "disable"
  manage_ssid_addresses          = "disable"
  device_name                    = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `pool_subnet` - Pool-Subnet.
* `automatic_conflict_resolution` - Enable/disable automatic conflict resolution. Valid values: `disable`, `enable`.

* `manage_lan_addresses` - Enable/disable default management of LAN interface addresses. Valid values: `disable`, `enable`.

* `manage_lan_extension_addresses` - Enable/disable default management of FortiExtender LAN extension interface addresses. Valid values: `disable`, `enable`.

* `manage_ssid_addresses` - Enable/disable default management of FortiAP SSID addresses. Valid values: `disable`, `enable`.

* `pools` - Pools. The structure of `pools` block is documented below.
* `require_subnet_size_match` - Enable/disable reassignment of subnets to make requested and actual sizes match. Valid values: `disable`, `enable`.

* `rules` - Rules. The structure of `rules` block is documented below.
* `server_type` - Configure the type of IPAM server to use. Valid values: `cloud`, `fabric-root`.

* `status` - Enable/disable IP address management services. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `pools` block supports:

* `description` - Description.
* `exclude` - Exclude. The structure of `exclude` block is documented below.
* `name` - IPAM pool name.
* `subnet` - Configure IPAM pool subnet, Class A - Class B subnet.

The `exclude` block supports:

* `id` - Exclude ID.
* `exclude_subnet` - Configure subnet to exclude from the IPAM pool.

The `rules` block supports:

* `description` - Description.
* `device` - Configure serial number or wildcard of FortiGate to match.
* `dhcp` - Enable/disable DHCP server for matching IPAM interfaces. Valid values: `disable`, `enable`.

* `interface` - Configure name or wildcard of interface to match.
* `name` - IPAM rule name.
* `pool` - Configure name of IPAM pool to use.
* `role` - Configure role of interface to match. Valid values: `any`, `lan`, `wan`, `dmz`, `undefined`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Ipam can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_ipam.labelname SystemIpam
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

