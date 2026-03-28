---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_prefixlist_rule"
description: |-
  IPv4 prefix list rule.
---

# fmgdevice_router_prefixlist_rule
IPv4 prefix list rule.

~> This resource is a sub resource for variable `rule` of resource `fmgdevice_router_prefixlist`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `prefix_list` - Prefix List.

* `action` - Permit or deny this IP address and netmask prefix. Valid values: `permit`, `deny`.

* `flags` - Flags.
* `ge` - Minimum prefix length to be matched (0 - 32).
* `fosid` - Rule ID.
* `le` - Maximum prefix length to be matched (0 - 32).
* `prefix` - IPv4 prefix to define regular filter criteria, such as "any" or subnets.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Router PrefixListRule can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "prefix_list=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_prefixlist_rule.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

