---
subcategory: "Router"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_multicast6flow_flows"
description: |-
  IPv6 multicast-flow entries.
---

# fmgdevice_router_multicast6flow_flows
IPv6 multicast-flow entries.

~> This resource is a sub resource for variable `flows` of resource `fmgdevice_router_multicast6flow`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `multicast6_flow` - Multicast6 Flow.

* `group_addr` - Multicast group IP address.
* `fosid` - Flow ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Router Multicast6FlowFlows can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "multicast6_flow=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_multicast6flow_flows.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

