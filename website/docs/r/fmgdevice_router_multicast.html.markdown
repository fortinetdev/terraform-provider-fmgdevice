---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_multicast"
description: |-
  Configure router multicast.
---

# fmgdevice_router_multicast
Configure router multicast.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `interface`: `fmgdevice_router_multicast_interface`
>- `pim_sm_global`: `fmgdevice_router_multicast_pimsmglobal`



## Example Usage

```hcl
resource "fmgdevice_router_multicast" "trname" {
  interface {
    bfd                 = "enable"
    cisco_exclude_genid = "disable"
    dr_priority         = 10
    hello_holdtime      = 10
    hello_interval      = 10
    igmp {
      access_group               = ["your own value"]
      immediate_leave_group      = ["your own value"]
      last_member_query_count    = 10
      last_member_query_interval = 10
      query_interval             = 10
      query_max_response_time    = 10
      query_timeout              = 10
      router_alert_check         = "disable"
      version                    = "3"
    }

    join_group {
      address = "your own value"
    }

    multicast_flow           = ["your own value"]
    name                     = "your own value"
    neighbour_filter         = ["your own value"]
    passive                  = "disable"
    pim_mode                 = "sparse-mode"
    propagation_delay        = 10
    rp_candidate             = "enable"
    rp_candidate_group       = ["your own value"]
    rp_candidate_interval    = 10
    rp_candidate_priority    = 10
    rpf_nbr_fail_back        = "enable"
    rpf_nbr_fail_back_filter = ["your own value"]
    state_refresh_interval   = 10
    static_group             = ["your own value"]
    ttl_threshold            = 10
  }

  multicast_routing = "disable"
  pim_sm_global {
    accept_register_list          = ["your own value"]
    accept_source_list            = ["your own value"]
    bsr_allow_quick_refresh       = "enable"
    bsr_candidate                 = "enable"
    bsr_hash                      = 10
    bsr_interface                 = ["port2"]
    bsr_priority                  = 10
    cisco_crp_prefix              = "enable"
    cisco_ignore_rp_set_priority  = "enable"
    cisco_register_checksum       = "disable"
    cisco_register_checksum_group = ["your own value"]
    join_prune_holdtime           = 10
    message_interval              = 10
    null_register_retries         = 10
    pim_use_sdwan                 = "disable"
    register_rate_limit           = 10
    register_rp_reachability      = "enable"
    register_source               = "interface"
    register_source_interface     = ["port2"]
    register_source_ip            = "your own value"
    register_supression           = 10
    rp_address {
      group      = ["your own value"]
      fosid      = 10
      ip_address = "your own value"
    }

    rp_register_keepalive = 10
    spt_threshold         = "enable"
    spt_threshold_group   = ["your own value"]
    ssm                   = "disable"
    ssm_range             = ["your own value"]
  }

  route_limit     = 10
  route_threshold = 10
  device_name     = var.device_name # not required if setting is at provider
  device_vdom     = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `interface` - Interface. The structure of `interface` block is documented below.
* `multicast_routing` - Enable/disable IP multicast routing. Valid values: `disable`, `enable`.

* `pim_sm_global` - Pim-Sm-Global. The structure of `pim_sm_global` block is documented below.
* `route_limit` - Maximum number of multicast routes.
* `route_threshold` - Generate warnings when the number of multicast routes exceeds this number, must not be greater than route-limit.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `interface` block supports:

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

The `pim_sm_global` block supports:

* `accept_register_list` - Sources allowed to register packets with this Rendezvous Point (RP).
* `accept_source_list` - Sources allowed to send multicast traffic.
* `bsr_allow_quick_refresh` - Enable/disable accept BSR quick refresh packets from neighbors. Valid values: `disable`, `enable`.

* `bsr_candidate` - Enable/disable allowing this router to become a bootstrap router (BSR). Valid values: `disable`, `enable`.

* `bsr_hash` - BSR hash length (0 - 32, default = 10).
* `bsr_interface` - Interface to advertise as candidate BSR.
* `bsr_priority` - BSR priority (0 - 255, default = 0).
* `cisco_crp_prefix` - Enable/disable making candidate RP compatible with old Cisco IOS. Valid values: `disable`, `enable`.

* `cisco_ignore_rp_set_priority` - Use only hash for RP selection (compatibility with old Cisco IOS). Valid values: `disable`, `enable`.

* `cisco_register_checksum` - Checksum entire register packet(for old Cisco IOS compatibility). Valid values: `disable`, `enable`.

* `cisco_register_checksum_group` - Cisco register checksum only these groups.
* `join_prune_holdtime` - Join/prune holdtime (1 - 65535, default = 210).
* `message_interval` - Period of time between sending periodic PIM join/prune messages in seconds (1 - 65535, default = 60).
* `null_register_retries` - Maximum retries of null register (1 - 20, default = 1).
* `pim_use_sdwan` - Enable/disable use of SDWAN when checking RPF neighbor and sending of REG packet. Valid values: `disable`, `enable`.

* `register_rate_limit` - Limit of packets/sec per source registered through this RP (0 - 65535, default = 0 which means unlimited).
* `register_rp_reachability` - Enable/disable check RP is reachable before registering packets. Valid values: `disable`, `enable`.

* `register_source` - Override source address in register packets. Valid values: `disable`, `ip-address`, `interface`.

* `register_source_interface` - Override with primary interface address.
* `register_source_ip` - Override with local IP address.
* `register_supression` - Period of time to honor register-stop message (1 - 65535 sec, default = 60).
* `rp_address` - Rp-Address. The structure of `rp_address` block is documented below.
* `rp_register_keepalive` - Timeout for RP receiving data on (S,G) tree (1 - 65535 sec, default = 185).
* `spt_threshold` - Enable/disable switching to source specific trees. Valid values: `disable`, `enable`.

* `spt_threshold_group` - Groups allowed to switch to source tree.
* `ssm` - Enable/disable source specific multicast. Valid values: `disable`, `enable`.

* `ssm_range` - Groups allowed to source specific multicast.

The `rp_address` block supports:

* `group` - Groups to use this RP.
* `id` - ID.
* `ip_address` - RP router address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router Multicast can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_multicast.labelname RouterMulticast
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

