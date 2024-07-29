---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_ripng_aggregateaddress"
description: |-
  Aggregate address.
---

# fmgdevice_router_ripng_aggregateaddress
Aggregate address.

~> This resource is a sub resource for variable `aggregate_address` of resource `fmgdevice_router_ripng`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_router_ripng_aggregateaddress" "trname" {
  fosid       = 10
  prefix6     = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `fosid` - Aggregate address entry ID.
* `prefix6` - Aggregate address prefix.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Router RipngAggregateAddress can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_ripng_aggregateaddress.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

