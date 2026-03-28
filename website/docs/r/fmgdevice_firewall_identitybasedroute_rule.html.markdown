---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_identitybasedroute_rule"
description: |-
  <i>This object will be purged after policy copy and install.</i> Rule.
---

# fmgdevice_firewall_identitybasedroute_rule
<i>This object will be purged after policy copy and install.</i> Rule.

~> This resource is a sub resource for variable `rule` of resource `fmgdevice_firewall_identitybasedroute`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `identity_based_route` - Identity Based Route.

* `device` - Outgoing interface for the rule.
* `gateway` - IPv4 address of the gateway (Format: xxx.xxx.xxx.xxx , Default: 0.0.0.0).
* `groups` - Select one or more group(s) from available groups that are allowed to use this route. Separate group names with a space.
* `fosid` - Rule ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall IdentityBasedRouteRule can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "identity_based_route=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_identitybasedroute_rule.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

