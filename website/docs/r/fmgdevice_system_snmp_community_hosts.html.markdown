---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_snmp_community_hosts"
description: |-
  Configure IPv4 SNMP managers (hosts).
---

# fmgdevice_system_snmp_community_hosts
Configure IPv4 SNMP managers (hosts).

~> This resource is a sub resource for variable `hosts` of resource `fmgdevice_system_snmp_community`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_snmp_community_hosts" "trname" {
  ha_direct   = "disable"
  host_type   = "any"
  fosid       = 10
  ip          = ["your own value"]
  source_ip   = "your own value"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `community` - Community.

* `ha_direct` - Enable/disable direct management of HA cluster members. Valid values: `disable`, `enable`.

* `host_type` - Control whether the SNMP manager sends SNMP queries, receives SNMP traps, or both. No traps will be sent when IP type is subnet. Valid values: `any`, `query`, `trap`.

* `fosid` - Host entry ID.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `ip` - IPv4 address of the SNMP manager (host).
* `source_ip` - Source IPv4 address for SNMP traps.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System SnmpCommunityHosts can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "community=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_snmp_community_hosts.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

