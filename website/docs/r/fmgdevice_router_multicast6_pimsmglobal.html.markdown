---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_multicast6_pimsmglobal"
description: |-
  PIM sparse-mode global settings.
---

# fmgdevice_router_multicast6_pimsmglobal
PIM sparse-mode global settings.

~> This resource is a sub resource for variable `pim_sm_global` of resource `fmgdevice_router_multicast6`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `rp_address`: `fmgdevice_router_multicast6_pimsmglobal_rpaddress`



## Example Usage

```hcl
resource "fmgdevice_router_multicast6_pimsmglobal" "trname" {
  register_rate_limit = 10
  rp_address {
    fosid       = 10
    ip6_address = "your own value"
  }

  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `pim_use_sdwan` - Enable/disable use of SDWAN when checking RPF neighbor and sending of REG packet. Valid values: `disable`, `enable`.

* `register_rate_limit` - Limit of packets/sec per source registered through this RP (0 means unlimited).
* `rp_address` - Rp-Address. The structure of `rp_address` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `rp_address` block supports:

* `id` - ID of the entry.
* `ip6_address` - RP router IPv6 address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router Multicast6PimSmGlobal can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_multicast6_pimsmglobal.labelname RouterMulticast6PimSmGlobal
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

