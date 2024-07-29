---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_nsxt_servicechain_serviceindex"
description: |-
  Configure service index.
---

# fmgdevice_nsxt_servicechain_serviceindex
Configure service index.

~> This resource is a sub resource for variable `service_index` of resource `fmgdevice_nsxt_servicechain`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_nsxt_servicechain_serviceindex" "trname" {
  fosid         = 10
  name          = "your own value"
  reverse_index = 10
  vd            = ["your own value"]
  device_name   = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `service_chain` - Service Chain.

* `fosid` - Service index.
* `name` - Index name.
* `reverse_index` - Reverse service index.
* `vd` - VDOM name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Nsxt ServiceChainServiceIndex can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "service_chain=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_nsxt_servicechain_serviceindex.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

