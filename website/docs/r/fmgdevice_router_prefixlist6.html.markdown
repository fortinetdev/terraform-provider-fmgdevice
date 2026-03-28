---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_prefixlist6"
description: |-
  Configure IPv6 prefix lists.
---

# fmgdevice_router_prefixlist6
Configure IPv6 prefix lists.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `rule`: `fmgdevice_router_prefixlist6_rule`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comments` - Comment.
* `name` - Name.
* `rule` - Rule. The structure of `rule` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `rule` block supports:

* `action` - Permit or deny packets that match this rule. Valid values: `permit`, `deny`.

* `flags` - Flags.
* `ge` - Minimum prefix length to be matched (0 - 128).
* `id` - Rule ID.
* `le` - Maximum prefix length to be matched (0 - 128).
* `prefix6` - IPv6 prefix to define regular filter criteria, such as "any" or subnets.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router PrefixList6 can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_prefixlist6.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

