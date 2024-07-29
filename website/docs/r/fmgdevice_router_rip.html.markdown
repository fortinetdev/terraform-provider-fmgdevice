---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_rip"
description: |-
  Configure RIP.
---

# fmgdevice_router_rip
Configure RIP.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `distance`: `fmgdevice_router_rip_distance`
>- `distribute_list`: `fmgdevice_router_rip_distributelist`
>- `interface`: `fmgdevice_router_rip_interface`
>- `neighbor`: `fmgdevice_router_rip_neighbor`
>- `network`: `fmgdevice_router_rip_network`
>- `offset_list`: `fmgdevice_router_rip_offsetlist`
>- `redistribute`: `fmgdevice_router_rip_redistribute`



## Example Usage

```hcl
resource "fmgdevice_router_rip" "trname" {
  default_information_originate = "disable"
  default_metric                = 10
  distance {
    access_list = ["your own value"]
    distance    = 10
    fosid       = 10
    prefix      = ["your own value"]
  }

  distribute_list {
    direction = "out"
    fosid     = 10
    interface = ["port2"]
    listname  = ["your own value"]
    status    = "enable"
  }

  garbage_timer = 10
  device_name   = var.device_name # not required if setting is at provider
  device_vdom   = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `default_information_originate` - Enable/disable generation of default route. Valid values: `disable`, `enable`.

* `default_metric` - Default metric.
* `distance` - Distance. The structure of `distance` block is documented below.
* `distribute_list` - Distribute-List. The structure of `distribute_list` block is documented below.
* `garbage_timer` - Garbage timer in seconds.
* `interface` - Interface. The structure of `interface` block is documented below.
* `max_out_metric` - Maximum metric allowed to output(0 means 'not set').
* `neighbor` - Neighbor. The structure of `neighbor` block is documented below.
* `network` - Network. The structure of `network` block is documented below.
* `offset_list` - Offset-List. The structure of `offset_list` block is documented below.
* `passive_interface` - Passive interface configuration.
* `recv_buffer_size` - Receiving buffer size.
* `redistribute` - Redistribute. The structure of `redistribute` block is documented below.
* `timeout_timer` - Timeout timer in seconds.
* `update_timer` - Update timer in seconds.
* `version` - RIP version. Valid values: `1`, `2`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `distance` block supports:

* `access_list` - Access list for route destination.
* `distance` - Distance (1 - 255).
* `id` - Distance ID.
* `prefix` - Distance prefix.

The `distribute_list` block supports:

* `direction` - Distribute list direction. Valid values: `out`, `in`.

* `id` - Distribute list ID.
* `interface` - Distribute list interface name.
* `listname` - Distribute access/prefix list name.
* `status` - Status. Valid values: `disable`, `enable`.


The `interface` block supports:

* `auth_keychain` - Authentication key-chain name.
* `auth_mode` - Authentication mode. Valid values: `none`, `md5`, `text`.

* `auth_string` - Authentication string/password.
* `flags` - Flags.
* `name` - Interface name.
* `receive_version` - Receive version. Valid values: `1`, `2`.

* `send_version` - Send version. Valid values: `1`, `2`.

* `send_version2_broadcast` - Enable/disable broadcast version 1 compatible packets. Valid values: `disable`, `enable`.

* `split_horizon` - Enable/disable split horizon. Valid values: `poisoned`, `regular`.

* `split_horizon_status` - Enable/disable split horizon. Valid values: `disable`, `enable`.


The `neighbor` block supports:

* `id` - Neighbor entry ID.
* `ip` - IP address.

The `network` block supports:

* `id` - Network entry ID.
* `prefix` - Network prefix.

The `offset_list` block supports:

* `access_list` - Access list name.
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

Router Rip can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_rip.labelname RouterRip
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

