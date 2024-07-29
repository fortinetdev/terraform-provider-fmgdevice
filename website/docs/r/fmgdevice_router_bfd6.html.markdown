---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_bfd6"
description: |-
  Configure IPv6 BFD.
---

# fmgdevice_router_bfd6
Configure IPv6 BFD.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `multihop_template`: `fmgdevice_router_bfd6_multihoptemplate`
>- `neighbor`: `fmgdevice_router_bfd6_neighbor`



## Example Usage

```hcl
resource "fmgdevice_router_bfd6" "trname" {
  multihop_template {
    auth_mode           = "none"
    bfd_desired_min_tx  = 10
    bfd_detect_mult     = 10
    bfd_required_min_rx = 10
    dst                 = "your own value"
    fosid               = 10
    md5_key             = ["your own value"]
    src                 = "your own value"
  }

  neighbor {
    interface   = ["port2"]
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

* `multihop_template` - Multihop-Template. The structure of `multihop_template` block is documented below.
* `neighbor` - Neighbor. The structure of `neighbor` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `multihop_template` block supports:

* `auth_mode` - Authentication mode. Valid values: `none`, `md5`.

* `bfd_desired_min_tx` - BFD desired minimal transmit interval (milliseconds).
* `bfd_detect_mult` - BFD detection multiplier.
* `bfd_required_min_rx` - BFD required minimal receive interval (milliseconds).
* `dst` - Destination prefix.
* `id` - ID.
* `md5_key` - MD5 key of key ID 1.
* `src` - Source prefix.

The `neighbor` block supports:

* `interface` - Interface to the BFD neighbor.
* `ip6_address` - IPv6 address of the BFD neighbor.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router Bfd6 can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_bfd6.labelname RouterBfd6
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

