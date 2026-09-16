---
subcategory: "Router"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_multicast6_pimsmglobalvrf_rpaddress"
description: |-
  Statically configured RP addresses.
---

# fmgdevice_router_multicast6_pimsmglobalvrf_rpaddress
Statically configured RP addresses.

~> This resource is a sub resource for variable `rp_address` of resource `fmgdevice_router_multicast6_pimsmglobalvrf`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `pim_sm_global_vrf` - Pim Sm Global Vrf.

* `group` - Groups to use this RP.
* `fosid` - ID of the entry.
* `ip6_address` - RP router IPv6 address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Router Multicast6PimSmGlobalVrfRpAddress can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "pim_sm_global_vrf=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_multicast6_pimsmglobalvrf_rpaddress.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

