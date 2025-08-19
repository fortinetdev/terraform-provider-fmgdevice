---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_multicast_pimsmglobalvrf"
description: |-
  per-VRF PIM sparse-mode global settings.
---

# fmgdevice_router_multicast_pimsmglobalvrf
per-VRF PIM sparse-mode global settings.

~> This resource is a sub resource for variable `pim_sm_global_vrf` of resource `fmgdevice_router_multicast`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `rp_address`: `fmgdevice_router_multicast_pimsmglobalvrf_rpaddress`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `bsr_allow_quick_refresh` - Enable/disable accept BSR quick refresh packets from neighbors. Valid values: `disable`, `enable`.

* `bsr_candidate` - Enable/disable allowing this router to become a bootstrap router (BSR). Valid values: `disable`, `enable`.

* `bsr_hash` - BSR hash length (0 - 32, default = 10).
* `bsr_interface` - Interface to advertise as candidate BSR.
* `bsr_priority` - BSR priority (0 - 255, default = 0).
* `cisco_crp_prefix` - Enable/disable making candidate RP compatible with old Cisco IOS. Valid values: `disable`, `enable`.

* `rp_address` - Rp-Address. The structure of `rp_address` block is documented below.
* `vrf` - VRF ID.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `rp_address` block supports:

* `group` - Groups to use this RP.
* `id` - ID.
* `ip_address` - RP router address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{vrf}}.

## Import

Router MulticastPimSmGlobalVrf can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_multicast_pimsmglobalvrf.labelname {{vrf}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

