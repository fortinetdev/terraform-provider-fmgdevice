---
subcategory: "Router"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_multicast6"
description: |-
  Configure IPv6 multicast.
---

# fmgdevice_router_multicast6
Configure IPv6 multicast.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `interface`: `fmgdevice_router_multicast6_interface`
>- `pim_sm_global`: `fmgdevice_router_multicast6_pimsmglobal`
>- `pim_sm_global_vrf`: `fmgdevice_router_multicast6_pimsmglobalvrf`



## Example Usage

```hcl
resource "fmgdevice_router_multicast6" "trname" {
  interface {
    hello_holdtime = 10
    hello_interval = 10
    name           = "your own value"
  }

  multicast_pmtu    = "enable"
  multicast_routing = "disable"
  pim_sm_global {
    register_rate_limit = 10
    rp_address {
      fosid       = 10
      ip6_address = "your own value"
    }

  }

  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `interface` - Interface. The structure of `interface` block is documented below.
* `multicast_pmtu` - Enable/disable PMTU for IPv6 multicast. Valid values: `disable`, `enable`.

* `multicast_routing` - Enable/disable IPv6 multicast routing. Valid values: `disable`, `enable`.

* `pim_sm_global` - Pim-Sm-Global. The structure of `pim_sm_global` block is documented below.
* `pim_sm_global_vrf` - Pim-Sm-Global-Vrf. The structure of `pim_sm_global_vrf` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `interface` block supports:

* `hello_holdtime` - Time before old neighbor information expires in seconds (1 - 65535, default = 105).
* `hello_interval` - Interval between sending PIM hello messages in seconds (1 - 65535, default = 30).
* `name` - Interface name.
* `rp_candidate` - Enable/disable compete to become RP in elections. Valid values: `disable`, `enable`.

* `rp_candidate_group` - Multicast groups managed by this RP.
* `rp_candidate_interval` - RP candidate advertisement interval (1 - 16383 sec, default = 60).
* `rp_candidate_priority` - Router's priority as RP.
* `static_group` - Statically set IPv6 multicast groups to forward out.

The `pim_sm_global` block supports:

* `bsr_allow_quick_refresh` - Enable/disable accept BSR quick refresh packets from neighbors. Valid values: `disable`, `enable`.

* `bsr_candidate` - Enable/disable allowing this router to become a bootstrap router (BSR). Valid values: `disable`, `enable`.

* `bsr_hash` - BSR hash length (0 - 128, default = 126).
* `bsr_interface` - Interface to advertise as candidate BSR.
* `bsr_priority` - BSR priority (0 - 255, default = 0).
* `cisco_crp_prefix` - Enable/disable making candidate RP compatible with old Cisco IOS. Valid values: `disable`, `enable`.

* `cisco_ignore_rp_set_priority` - Use only hash for RP selection (compatibility with old Cisco IOS). Valid values: `disable`, `enable`.

* `pim_use_sdwan` - Enable/disable use of SDWAN when checking RPF neighbor and sending of REG packet. Valid values: `disable`, `enable`.

* `register_rate_limit` - Limit of packets/sec per source registered through this RP (0 means unlimited).
* `rp_address` - Rp-Address. The structure of `rp_address` block is documented below.
* `spt_threshold` - Enable/disable switching to source specific trees. Valid values: `disable`, `enable`.

* `spt_threshold_group` - Groups allowed to switch to source tree.

The `rp_address` block supports:

* `group` - Groups to use this RP.
* `id` - ID of the entry.
* `ip6_address` - RP router IPv6 address.

The `pim_sm_global_vrf` block supports:

* `bsr_allow_quick_refresh` - Enable/disable accept BSR quick refresh packets from neighbors. Valid values: `disable`, `enable`.

* `bsr_candidate` - Enable/disable allowing this router to become a bootstrap router (BSR). Valid values: `disable`, `enable`.

* `bsr_hash` - BSR hash length (0 - 128, default = 126).
* `bsr_interface` - Interface to advertise as candidate BSR.
* `bsr_priority` - BSR priority (0 - 255, default = 0).
* `cisco_crp_prefix` - Enable/disable making candidate RP compatible with old Cisco IOS. Valid values: `disable`, `enable`.

* `rp_address` - Rp-Address. The structure of `rp_address` block is documented below.
* `vrf` - VRF ID.

The `rp_address` block supports:

* `group` - Groups to use this RP.
* `id` - ID of the entry.
* `ip6_address` - RP router IPv6 address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router Multicast6 can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_multicast6.labelname RouterMulticast6
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

