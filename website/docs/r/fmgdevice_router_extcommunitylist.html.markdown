---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_extcommunitylist"
description: |-
  Configure extended community lists.
---

# fmgdevice_router_extcommunitylist
Configure extended community lists.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `rule`: `fmgdevice_router_extcommunitylist_rule`



## Example Usage

```hcl
resource "fmgdevice_router_extcommunitylist" "trname" {
  name = "your own value"
  rule {
    action = "deny"
    fosid  = 10
    match  = "your own value"
    regexp = "your own value"
    type   = "soo"
  }

  type        = "standard"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `name` - Extended community list name.
* `rule` - Rule. The structure of `rule` block is documented below.
* `type` - Extended community list type (standard or expanded). Valid values: `standard`, `expanded`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `rule` block supports:

* `action` - Permit or deny route-based operations, based on the route's EXTENDED COMMUNITY attribute. Valid values: `permit`, `deny`.

* `id` - ID.
* `match` - Extended community specifications for matching a reserved extended community.
* `regexp` - Ordered list of EXTENDED COMMUNITY attributes as a regular expression.
* `type` - Type of extended community. Valid values: `rt`, `soo`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router ExtcommunityList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_extcommunitylist.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

