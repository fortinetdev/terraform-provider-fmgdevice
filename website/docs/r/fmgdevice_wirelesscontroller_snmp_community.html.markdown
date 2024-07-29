---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_snmp_community"
description: |-
  SNMP Community Configuration.
---

# fmgdevice_wirelesscontroller_snmp_community
SNMP Community Configuration.

~> This resource is a sub resource for variable `community` of resource `fmgdevice_wirelesscontroller_snmp`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `hosts`: `fmgdevice_wirelesscontroller_snmp_community_hosts`



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_snmp_community" "trname" {
  hosts {
    fosid = 10
    ip    = "your own value"
  }

  fosid            = 10
  name             = "your own value"
  query_v1_status  = "disable"
  query_v2c_status = "disable"
  device_name      = var.device_name # not required if setting is at provider
  device_vdom      = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `hosts` - Hosts. The structure of `hosts` block is documented below.
* `fosid` - Community ID.
* `name` - Community name.
* `query_v1_status` - Enable/disable SNMP v1 queries. Valid values: `disable`, `enable`.

* `query_v2c_status` - Enable/disable SNMP v2c queries. Valid values: `disable`, `enable`.

* `status` - Enable/disable this SNMP community. Valid values: `disable`, `enable`.

* `trap_v1_status` - Enable/disable SNMP v1 traps. Valid values: `disable`, `enable`.

* `trap_v2c_status` - Enable/disable SNMP v2c traps. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `hosts` block supports:

* `id` - Host entry ID.
* `ip` - IPv4 address of the SNMP manager (host).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

WirelessController SnmpCommunity can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_snmp_community.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

