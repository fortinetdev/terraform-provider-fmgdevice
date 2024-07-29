---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_multicast_interface"
description: |-
  PIM interfaces.
---

# fmgdevice_router_multicast_interface
PIM interfaces.

~> This resource is a sub resource for variable `interface` of resource `fmgdevice_router_multicast`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `igmp`: `fmgdevice_router_multicast_interface_igmp`
>- `join_group`: `fmgdevice_router_multicast_interface_joingroup`



## Example Usage

```hcl
resource "fmgdevice_router_multicast_interface" "trname" {
  bfd                 = "enable"
  cisco_exclude_genid = "disable"
  dr_priority         = 10
  hello_holdtime      = 10
  hello_interval      = 10
  name                = "your own value"
  device_name         = var.device_name # not required if setting is at provider
  device_vdom         = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `bfd` - Enable/disable Protocol Independent Multicast (PIM) Bidirectional Forwarding Detection (BFD). Valid values: `disable`, `enable`.

* `cisco_exclude_genid` - Exclude GenID from hello packets (compatibility with old Cisco IOS). Valid values: `disable`, `enable`.

* `dr_priority` - DR election priority.
* `hello_holdtime` - Time before old neighbor information expires (0 - 65535 sec, default = 105).
* `hello_interval` - Interval between sending PIM hello messages (0 - 65535 sec, default = 30).
* `igmp` - Igmp. The structure of `igmp` block is documented below.
* `join_group` - Join-Group. The structure of `join_group` block is documented below.
* `multicast_flow` - Acceptable source for multicast group.
* `name` - Interface name.
* `neighbour_filter` - Routers acknowledged as neighbor routers.
* `passive` - Enable/disable listening to IGMP but not participating in PIM. Valid values: `disable`, `enable`.

* `pim_mode` - PIM operation mode. Valid values: `sparse-mode`, `dense-mode`.

* `propagation_delay` - Delay flooding packets on this interface (100 - 5000 msec, default = 500).
* `rp_candidate` - Enable/disable compete to become RP in elections. Valid values: `disable`, `enable`.

* `rp_candidate_group` - Multicast groups managed by this RP.
* `rp_candidate_interval` - RP candidate advertisement interval (1 - 16383 sec, default = 60).
* `rp_candidate_priority` - Router's priority as RP.
* `rpf_nbr_fail_back` - Enable/disable fail back for RPF neighbor query. Valid values: `disable`, `enable`.

* `rpf_nbr_fail_back_filter` - Filter for fail back RPF neighbors.
* `state_refresh_interval` - Interval between sending state-refresh packets (1 - 100 sec, default = 60).
* `static_group` - Statically set multicast groups to forward out.
* `ttl_threshold` - Minimum TTL of multicast packets that will be forwarded (applied only to new multicast routes) (1 - 255, default = 1).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `igmp` block supports:

* `access_group` - Groups IGMP hosts are allowed to join.
* `immediate_leave_group` - Groups to drop membership for immediately after receiving IGMPv2 leave.
* `last_member_query_count` - Number of group specific queries before removing group (2 - 7, default = 2).
* `last_member_query_interval` - Timeout between IGMPv2 leave and removing group (1 - 65535 msec, default = 1000).
* `query_interval` - Interval between queries to IGMP hosts (1 - 65535 sec, default = 125).
* `query_max_response_time` - Maximum time to wait for a IGMP query response (1 - 25 sec, default = 10).
* `query_timeout` - Timeout between queries before becoming querying unit for network (60 - 900, default = 255).
* `router_alert_check` - Enable/disable require IGMP packets contain router alert option. Valid values: `disable`, `enable`.

* `version` - Maximum version of IGMP to support. Valid values: `1`, `2`, `3`.


The `join_group` block supports:

* `address` - Multicast group IP address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router MulticastInterface can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_multicast_interface.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

