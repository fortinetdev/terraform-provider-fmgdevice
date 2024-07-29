---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_rip_offsetlist"
description: |-
  Offset list.
---

# fmgdevice_router_rip_offsetlist
Offset list.

~> This resource is a sub resource for variable `offset_list` of resource `fmgdevice_router_rip`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_router_rip_offsetlist" "trname" {
  access_list = ["your own value"]
  direction   = "in"
  fosid       = 10
  interface   = ["port2"]
  offset      = 10
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `access_list` - Access list name.
* `direction` - Offset list direction. Valid values: `out`, `in`.

* `fosid` - Offset-list ID.
* `interface` - Interface name.
* `offset` - Offset.
* `status` - Status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Router RipOffsetList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_rip_offsetlist.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

