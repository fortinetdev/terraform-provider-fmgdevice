---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_extcommunitylist_rule"
description: |-
  Extended community list rule.
---

# fmgdevice_router_extcommunitylist_rule
Extended community list rule.

~> This resource is a sub resource for variable `rule` of resource `fmgdevice_router_extcommunitylist`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_router_extcommunitylist_rule" "trname" {
  action      = "permit"
  fosid       = 10
  match       = "your own value"
  regexp      = "your own value"
  type        = "soo"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `extcommunity_list` - Extcommunity List.

* `action` - Permit or deny route-based operations, based on the route's EXTENDED COMMUNITY attribute. Valid values: `permit`, `deny`.

* `fosid` - ID.
* `match` - Extended community specifications for matching a reserved extended community.
* `regexp` - Ordered list of EXTENDED COMMUNITY attributes as a regular expression.
* `type` - Type of extended community. Valid values: `rt`, `soo`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Router ExtcommunityListRule can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "extcommunity_list=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_extcommunitylist_rule.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

