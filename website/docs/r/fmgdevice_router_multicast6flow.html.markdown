---
subcategory: "Router"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_multicast6flow"
description: |-
  Configure IPv6 multicast-flow.
---

# fmgdevice_router_multicast6flow
Configure IPv6 multicast-flow.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `flows`: `fmgdevice_router_multicast6flow_flows`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comments` - Comment.
* `flows` - Flows. The structure of `flows` block is documented below.
* `name` - Name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `flows` block supports:

* `group_addr` - Multicast group IP address.
* `id` - Flow ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router Multicast6Flow can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_multicast6flow.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

