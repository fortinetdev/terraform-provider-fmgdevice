---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_bgp_vrf"
description: |-
  BGP VRF leaking table.
---

# fmgdevice_router_bgp_vrf
BGP VRF leaking table.

~> This resource is a sub resource for variable `vrf` of resource `fmgdevice_router_bgp`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `leak_target`: `fmgdevice_router_bgp_vrf_leaktarget`



## Example Usage

```hcl
resource "fmgdevice_router_bgp_vrf" "trname" {
  export_rt        = ["your own value"]
  import_route_map = ["your own value"]
  import_rt        = ["your own value"]
  leak_target {
    interface = ["port2"]
    route_map = ["your own value"]
    vrf       = "your own value"
  }

  rd          = "your own value"
  vrf         = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `export_rt` - List of export route target.
* `import_route_map` - Import route map.
* `import_rt` - List of import route target.
* `leak_target` - Leak-Target. The structure of `leak_target` block is documented below.
* `rd` - Route Distinguisher: AA:NN|A.B.C.D:NN.
* `role` - VRF role. Valid values: `standalone`, `ce`, `pe`.

* `vrf` - Origin VRF ID (0 - 251).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `leak_target` block supports:

* `interface` - Interface which is used to leak routes to target VRF.
* `route_map` - Route map of VRF leaking.
* `vrf` - Target VRF ID (0 - 251).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{vrf}}.

## Import

Router BgpVrf can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_bgp_vrf.labelname {{vrf}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

