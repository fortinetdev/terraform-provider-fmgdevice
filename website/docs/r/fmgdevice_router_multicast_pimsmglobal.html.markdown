---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_multicast_pimsmglobal"
description: |-
  PIM sparse-mode global settings.
---

# fmgdevice_router_multicast_pimsmglobal
PIM sparse-mode global settings.

~> This resource is a sub resource for variable `pim_sm_global` of resource `fmgdevice_router_multicast`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `rp_address`: `fmgdevice_router_multicast_pimsmglobal_rpaddress`



## Example Usage

```hcl
resource "fmgdevice_router_multicast_pimsmglobal" "trname" {
  accept_register_list    = ["your own value"]
  accept_source_list      = ["your own value"]
  bsr_allow_quick_refresh = "disable"
  bsr_candidate           = "disable"
  bsr_hash                = 10
  device_name             = var.device_name # not required if setting is at provider
  device_vdom             = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

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
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `rp_address` block supports:

* `group` - Groups to use this RP.
* `id` - ID.
* `ip_address` - RP router address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router MulticastPimSmGlobal can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_multicast_pimsmglobal.labelname RouterMulticastPimSmGlobal
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

