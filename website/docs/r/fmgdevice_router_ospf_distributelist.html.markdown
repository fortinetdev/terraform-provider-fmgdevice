---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_ospf_distributelist"
description: |-
  Distribute list configuration.
---

# fmgdevice_router_ospf_distributelist
Distribute list configuration.

~> This resource is a sub resource for variable `distribute_list` of resource `fmgdevice_router_ospf`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_router_ospf_distributelist" "trname" {
  access_list = ["your own value"]
  fosid       = 10
  protocol    = "rip"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `access_list` - Access list name.
* `fosid` - Distribute list entry ID.
* `protocol` - Protocol type. Valid values: `connected`, `static`, `rip`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Router OspfDistributeList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_ospf_distributelist.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

