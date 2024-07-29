---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_multicast_interface_igmp"
description: |-
  IGMP configuration options.
---

# fmgdevice_router_multicast_interface_igmp
IGMP configuration options.

~> This resource is a sub resource for variable `igmp` of resource `fmgdevice_router_multicast_interface`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_router_multicast_interface_igmp" "trname" {
  access_group               = ["your own value"]
  immediate_leave_group      = ["your own value"]
  last_member_query_count    = 10
  last_member_query_interval = 10
  query_interval             = 10
  device_name                = var.device_name # not required if setting is at provider
  device_vdom                = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `interface` - Interface.

* `access_group` - Groups IGMP hosts are allowed to join.
* `immediate_leave_group` - Groups to drop membership for immediately after receiving IGMPv2 leave.
* `last_member_query_count` - Number of group specific queries before removing group (2 - 7, default = 2).
* `last_member_query_interval` - Timeout between IGMPv2 leave and removing group (1 - 65535 msec, default = 1000).
* `query_interval` - Interval between queries to IGMP hosts (1 - 65535 sec, default = 125).
* `query_max_response_time` - Maximum time to wait for a IGMP query response (1 - 25 sec, default = 10).
* `query_timeout` - Timeout between queries before becoming querying unit for network (60 - 900, default = 255).
* `router_alert_check` - Enable/disable require IGMP packets contain router alert option. Valid values: `disable`, `enable`.

* `version` - Maximum version of IGMP to support. Valid values: `1`, `2`, `3`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router MulticastInterfaceIgmp can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "interface=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_multicast_interface_igmp.labelname RouterMulticastInterfaceIgmp
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

