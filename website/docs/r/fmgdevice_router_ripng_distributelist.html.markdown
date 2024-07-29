---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_ripng_distributelist"
description: |-
  Distribute list.
---

# fmgdevice_router_ripng_distributelist
Distribute list.

~> This resource is a sub resource for variable `distribute_list` of resource `fmgdevice_router_ripng`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_router_ripng_distributelist" "trname" {
  direction   = "out"
  fosid       = 10
  interface   = ["port2"]
  listname    = ["your own value"]
  status      = "enable"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `direction` - Distribute list direction. Valid values: `out`, `in`.

* `fosid` - Distribute list ID.
* `interface` - Distribute list interface name.
* `listname` - Distribute access/prefix list name.
* `status` - Status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Router RipngDistributeList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_ripng_distributelist.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

