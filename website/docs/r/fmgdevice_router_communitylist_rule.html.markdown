---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_communitylist_rule"
description: |-
  Community list rule.
---

# fmgdevice_router_communitylist_rule
Community list rule.

~> This resource is a sub resource for variable `rule` of resource `fmgdevice_router_communitylist`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `community_list` - Community List.

* `action` - Permit or deny route-based operations, based on the route's COMMUNITY attribute. Valid values: `deny`, `permit`.

* `fosid` - ID.
* `match` - Community specifications for matching a reserved community.
* `regexp` - Ordered list of COMMUNITY attributes as a regular expression.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Router CommunityListRule can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "community_list=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_communitylist_rule.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

