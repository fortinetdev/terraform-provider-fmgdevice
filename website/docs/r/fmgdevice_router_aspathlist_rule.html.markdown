---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_aspathlist_rule"
description: |-
  AS path list rule.
---

# fmgdevice_router_aspathlist_rule
AS path list rule.

~> This resource is a sub resource for variable `rule` of resource `fmgdevice_router_aspathlist`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `aspath_list` - Aspath List.

* `action` - Permit or deny route-based operations, based on the route's AS_PATH attribute. Valid values: `deny`, `permit`.

* `fosid` - ID.
* `regexp` - Regular-expression to match the Border Gateway Protocol (BGP) AS paths.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Router AspathListRule can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "aspath_list=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_aspathlist_rule.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

