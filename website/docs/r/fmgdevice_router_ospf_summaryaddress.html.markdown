---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_ospf_summaryaddress"
description: |-
  IP address summary configuration.
---

# fmgdevice_router_ospf_summaryaddress
IP address summary configuration.

~> This resource is a sub resource for variable `summary_address` of resource `fmgdevice_router_ospf`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_router_ospf_summaryaddress" "trname" {
  advertise   = "disable"
  fosid       = 10
  prefix      = ["your own value"]
  tag         = 10
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `advertise` - Enable/disable advertise status. Valid values: `disable`, `enable`.

* `fosid` - Summary address entry ID.
* `prefix` - Prefix.
* `tag` - Tag value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Router OspfSummaryAddress can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_ospf_summaryaddress.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

