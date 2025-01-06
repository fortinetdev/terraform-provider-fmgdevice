---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_snmp_community_hosts6"
description: |-
  Configure IPv6 SNMP managers.
---

# fmgdevice_system_snmp_community_hosts6
Configure IPv6 SNMP managers.

~> This resource is a sub resource for variable `hosts6` of resource `fmgdevice_system_snmp_community`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_snmp_community_hosts6" "trname" {
  ha_direct   = "disable"
  host_type   = "query"
  fosid       = 10
  ipv6        = "your own value"
  source_ipv6 = "your own value"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `community` - Community.

* `ha_direct` - Enable/disable direct management of HA cluster members. Valid values: `disable`, `enable`.

* `host_type` - Control whether the SNMP manager sends SNMP queries, receives SNMP traps, or both. Valid values: `any`, `query`, `trap`.

* `fosid` - Host6 entry ID.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `ipv6` - SNMP manager IPv6 address prefix.
* `source_ipv6` - Source IPv6 address for SNMP traps.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System SnmpCommunityHosts6 can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "community=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_snmp_community_hosts6.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

