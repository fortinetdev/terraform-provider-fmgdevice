---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_accesslist6_rule"
description: |-
  Rule.
---

# fmgdevice_router_accesslist6_rule
Rule.

~> This resource is a sub resource for variable `rule` of resource `fmgdevice_router_accesslist6`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `access_list6` - Access List6.

* `action` - Permit or deny this IP address and netmask prefix. Valid values: `permit`, `deny`.

* `exact_match` - Enable/disable exact prefix match. Valid values: `disable`, `enable`.

* `flags` - Flags.
* `fosid` - Rule ID.
* `prefix6` - IPv6 prefix to define regular filter criteria, such as "any" or subnets.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Router AccessList6Rule can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "access_list6=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_accesslist6_rule.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

