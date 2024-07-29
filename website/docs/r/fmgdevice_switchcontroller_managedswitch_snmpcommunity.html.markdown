---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_managedswitch_snmpcommunity"
description: |-
  Configuration method to edit Simple Network Management Protocol (SNMP) communities.
---

# fmgdevice_switchcontroller_managedswitch_snmpcommunity
Configuration method to edit Simple Network Management Protocol (SNMP) communities.

~> This resource is a sub resource for variable `snmp_community` of resource `fmgdevice_switchcontroller_managedswitch`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `hosts`: `fmgdevice_switchcontroller_managedswitch_snmpcommunity_hosts`



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_managedswitch_snmpcommunity" "trname" {
  events = ["cpu-high"]
  hosts {
    fosid = 10
    ip    = ["your own value"]
  }

  fosid         = 10
  name          = "your own value"
  query_v1_port = 10
  device_name   = var.device_name # not required if setting is at provider
  device_vdom   = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `managed_switch` - Managed Switch.

* `events` - SNMP notifications (traps) to send. Valid values: `cpu-high`, `mem-low`, `log-full`, `intf-ip`, `ent-conf-change`.

* `hosts` - Hosts. The structure of `hosts` block is documented below.
* `fosid` - SNMP community ID.
* `name` - SNMP community name.
* `query_v1_port` - SNMP v1 query port (default = 161).
* `query_v1_status` - Enable/disable SNMP v1 queries. Valid values: `disable`, `enable`.

* `query_v2c_port` - SNMP v2c query port (default = 161).
* `query_v2c_status` - Enable/disable SNMP v2c queries. Valid values: `disable`, `enable`.

* `status` - Enable/disable this SNMP community. Valid values: `disable`, `enable`.

* `trap_v1_lport` - SNMP v2c trap local port (default = 162).
* `trap_v1_rport` - SNMP v2c trap remote port (default = 162).
* `trap_v1_status` - Enable/disable SNMP v1 traps. Valid values: `disable`, `enable`.

* `trap_v2c_lport` - SNMP v2c trap local port (default = 162).
* `trap_v2c_rport` - SNMP v2c trap remote port (default = 162).
* `trap_v2c_status` - Enable/disable SNMP v2c traps. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `hosts` block supports:

* `id` - Host entry ID.
* `ip` - IPv4 address of the SNMP manager (host).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

SwitchController ManagedSwitchSnmpCommunity can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "managed_switch=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_managedswitch_snmpcommunity.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

