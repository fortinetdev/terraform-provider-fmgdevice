---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_internetserviceaddition_entry"
description: |-
  <i>This object will be purged after policy copy and install.</i> Entries added to the Internet Service addition database.
---

# fmgdevice_firewall_internetserviceaddition_entry
<i>This object will be purged after policy copy and install.</i> Entries added to the Internet Service addition database.

~> This resource is a sub resource for variable `entry` of resource `fmgdevice_firewall_internetserviceaddition`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `port_range`: `fmgdevice_firewall_internetserviceaddition_entry_portrange`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `internet_service_addition` - Internet Service Addition.

* `addr_mode` - Address mode (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.

* `fosid` - Entry ID(1-255).
* `port_range` - Port-Range. The structure of `port_range` block is documented below.
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `port_range` block supports:

* `end_port` - Integer value for ending TCP/UDP/SCTP destination port in range (0 to 65535).
* `id` - Custom entry port range ID.
* `start_port` - Integer value for starting TCP/UDP/SCTP destination port in range (0 to 65535).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall InternetServiceAdditionEntry can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "internet_service_addition=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_internetserviceaddition_entry.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

