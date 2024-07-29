---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_bfd6_multihoptemplate"
description: |-
  BFD IPv6 multi-hop template table.
---

# fmgdevice_router_bfd6_multihoptemplate
BFD IPv6 multi-hop template table.

~> This resource is a sub resource for variable `multihop_template` of resource `fmgdevice_router_bfd6`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_router_bfd6_multihoptemplate" "trname" {
  auth_mode           = "md5"
  bfd_desired_min_tx  = 10
  bfd_detect_mult     = 10
  bfd_required_min_rx = 10
  dst                 = "your own value"
  fosid               = 10
  device_name         = var.device_name # not required if setting is at provider
  device_vdom         = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `auth_mode` - Authentication mode. Valid values: `none`, `md5`.

* `bfd_desired_min_tx` - BFD desired minimal transmit interval (milliseconds).
* `bfd_detect_mult` - BFD detection multiplier.
* `bfd_required_min_rx` - BFD required minimal receive interval (milliseconds).
* `dst` - Destination prefix.
* `fosid` - ID.
* `md5_key` - MD5 key of key ID 1.
* `src` - Source prefix.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Router Bfd6MultihopTemplate can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_bfd6_multihoptemplate.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

