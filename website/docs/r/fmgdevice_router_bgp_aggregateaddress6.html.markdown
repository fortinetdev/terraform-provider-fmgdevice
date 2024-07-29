---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_bgp_aggregateaddress6"
description: |-
  BGP IPv6 aggregate address table.
---

# fmgdevice_router_bgp_aggregateaddress6
BGP IPv6 aggregate address table.

~> This resource is a sub resource for variable `aggregate_address6` of resource `fmgdevice_router_bgp`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_router_bgp_aggregateaddress6" "trname" {
  as_set       = "disable"
  fosid        = 10
  prefix6      = "your own value"
  summary_only = "disable"
  device_name  = var.device_name # not required if setting is at provider
  device_vdom  = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `as_set` - Enable/disable generate AS set path information. Valid values: `disable`, `enable`.

* `fosid` - ID.
* `prefix6` - Aggregate IPv6 prefix.
* `summary_only` - Enable/disable filter more specific routes from updates. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Router BgpAggregateAddress6 can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_bgp_aggregateaddress6.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

