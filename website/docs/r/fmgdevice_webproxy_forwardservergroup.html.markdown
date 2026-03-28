---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webproxy_forwardservergroup"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure a forward server group consisting or multiple forward servers. Supports failover and load balancing.
---

# fmgdevice_webproxy_forwardservergroup
<i>This object will be purged after policy copy and install.</i> Configure a forward server group consisting or multiple forward servers. Supports failover and load balancing.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `server_list`: `fmgdevice_webproxy_forwardservergroup_serverlist`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `affinity` - Enable/disable affinity, attaching a source-ip's traffic to the assigned forwarding server until the forward-server-affinity-timeout is reached (under web-proxy global). Valid values: `disable`, `enable`.

* `group_down_option` - Action to take when all of the servers in the forward server group are down: block sessions until at least one server is back up or pass sessions to their destination. Valid values: `block`, `pass`.

* `ldb_method` - Load balance method: weighted or least-session. Valid values: `weighted`, `least-session`, `active-passive`.

* `name` - Configure a forward server group consisting one or multiple forward servers. Supports failover and load balancing.
* `server_list` - Server-List. The structure of `server_list` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `server_list` block supports:

* `name` - Forward server name.
* `weight` - Optionally assign a weight of the forwarding server for weighted load balancing (1 - 100, default = 10).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WebProxy ForwardServerGroup can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webproxy_forwardservergroup.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

