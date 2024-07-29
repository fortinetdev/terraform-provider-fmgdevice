---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_isis_summaryaddress"
description: |-
  IS-IS summary addresses.
---

# fmgdevice_router_isis_summaryaddress
IS-IS summary addresses.

~> This resource is a sub resource for variable `summary_address` of resource `fmgdevice_router_isis`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_router_isis_summaryaddress" "trname" {
  fosid       = 10
  level       = "level-1"
  prefix      = ["your own value"]
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `fosid` - Summary address entry ID.
* `level` - Level. Valid values: `level-1-2`, `level-1`, `level-2`.

* `prefix` - Prefix.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Router IsisSummaryAddress can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_isis_summaryaddress.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

