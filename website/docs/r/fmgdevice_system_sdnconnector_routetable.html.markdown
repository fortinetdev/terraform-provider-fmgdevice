---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_sdnconnector_routetable"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure Azure route table.
---

# fmgdevice_system_sdnconnector_routetable
<i>This object will be purged after policy copy and install.</i> Configure Azure route table.

~> This resource is a sub resource for variable `route_table` of resource `fmgdevice_system_sdnconnector`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `route`: `fmgdevice_system_sdnconnector_routetable_route`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `sdn_connector` - Sdn Connector.

* `name` - Route table name.
* `resource_group` - Resource group of Azure route table.
* `route` - Route. The structure of `route` block is documented below.
* `subscription_id` - Subscription ID of Azure route table.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `route` block supports:

* `name` - Route name.
* `next_hop` - Next hop address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System SdnConnectorRouteTable can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "sdn_connector=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_sdnconnector_routetable.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

