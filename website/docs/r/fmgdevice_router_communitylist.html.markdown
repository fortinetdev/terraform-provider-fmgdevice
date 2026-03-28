---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_communitylist"
description: |-
  Configure community lists.
---

# fmgdevice_router_communitylist
Configure community lists.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `rule`: `fmgdevice_router_communitylist_rule`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `name` - Community list name.
* `rule` - Rule. The structure of `rule` block is documented below.
* `type` - Community list type (standard or expanded). Valid values: `standard`, `expanded`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `rule` block supports:

* `action` - Permit or deny route-based operations, based on the route's COMMUNITY attribute. Valid values: `deny`, `permit`.

* `id` - ID.
* `match` - Community specifications for matching a reserved community.
* `regexp` - Ordered list of COMMUNITY attributes as a regular expression.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router CommunityList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_communitylist.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

