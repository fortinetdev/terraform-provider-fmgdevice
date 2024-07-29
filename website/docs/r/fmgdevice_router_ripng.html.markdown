---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_ripng"
description: |-
  Configure RIPng.
---

# fmgdevice_router_ripng
Configure RIPng.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `aggregate_address`: `fmgdevice_router_ripng_aggregateaddress`
>- `distance`: `fmgdevice_router_ripng_distance`
>- `distribute_list`: `fmgdevice_router_ripng_distributelist`
>- `interface`: `fmgdevice_router_ripng_interface`
>- `neighbor`: `fmgdevice_router_ripng_neighbor`
>- `network`: `fmgdevice_router_ripng_network`
>- `offset_list`: `fmgdevice_router_ripng_offsetlist`
>- `redistribute`: `fmgdevice_router_ripng_redistribute`



## Example Usage

```hcl
resource "fmgdevice_router_ripng" "trname" {
  aggregate_address {
    fosid   = 10
    prefix6 = "your own value"
  }

  default_information_originate = "enable"
  default_metric                = 10
  distance {
    access_list6 = ["your own value"]
    distance     = 10
    fosid        = 10
    prefix6      = "your own value"
  }

  distribute_list {
    direction = "in"
    fosid     = 10
    interface = ["port2"]
    listname  = ["your own value"]
    status    = "enable"
  }

  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `aggregate_address` - Aggregate-Address. The structure of `aggregate_address` block is documented below.
* `default_information_originate` - Enable/disable generation of default route. Valid values: `disable`, `enable`.

* `default_metric` - Default metric.
* `distance` - Distance. The structure of `distance` block is documented below.
* `distribute_list` - Distribute-List. The structure of `distribute_list` block is documented below.
* `garbage_timer` - Garbage timer.
* `interface` - Interface. The structure of `interface` block is documented below.
* `max_out_metric` - Maximum metric allowed to output(0 means 'not set').
* `neighbor` - Neighbor. The structure of `neighbor` block is documented below.
* `network` - Network. The structure of `network` block is documented below.
* `offset_list` - Offset-List. The structure of `offset_list` block is documented below.
* `passive_interface` - Passive interface configuration.
* `redistribute` - Redistribute. The structure of `redistribute` block is documented below.
* `timeout_timer` - Timeout timer.
* `update_timer` - Update timer.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `aggregate_address` block supports:

* `id` - Aggregate address entry ID.
* `prefix6` - Aggregate address prefix.

The `distance` block supports:

* `access_list6` - Access list for route destination.
* `distance` - Distance (1 - 255).
* `id` - Distance ID.
* `prefix6` - Distance prefix6.

The `distribute_list` block supports:

* `direction` - Distribute list direction. Valid values: `out`, `in`.

* `id` - Distribute list ID.
* `interface` - Distribute list interface name.
* `listname` - Distribute access/prefix list name.
* `status` - Status. Valid values: `disable`, `enable`.


The `interface` block supports:

* `flags` - Flags.
* `name` - Interface name.
* `split_horizon` - Enable/disable split horizon. Valid values: `poisoned`, `regular`.

* `split_horizon_status` - Enable/disable split horizon. Valid values: `disable`, `enable`.


The `neighbor` block supports:

* `id` - Neighbor entry ID.
* `interface` - Interface name.
* `ip6` - IPv6 link-local address.

The `network` block supports:

* `id` - Network entry ID.
* `prefix` - Network IPv6 link-local prefix.

The `offset_list` block supports:

* `access_list6` - IPv6 access list name.
* `direction` - Offset list direction. Valid values: `out`, `in`.

* `id` - Offset-list ID.
* `interface` - Interface name.
* `offset` - Offset.
* `status` - Status. Valid values: `disable`, `enable`.


The `redistribute` block supports:

* `metric` - Redistribute metric setting.
* `name` - Redistribute name.
* `routemap` - Route map name.
* `status` - Status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router Ripng can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_ripng.labelname RouterRipng
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

