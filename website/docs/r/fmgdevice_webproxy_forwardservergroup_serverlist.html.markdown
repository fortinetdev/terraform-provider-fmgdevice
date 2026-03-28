---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webproxy_forwardservergroup_serverlist"
description: |-
  <i>This object will be purged after policy copy and install.</i> Add web forward servers to a list to form a server group. Optionally assign weights to each server.
---

# fmgdevice_webproxy_forwardservergroup_serverlist
<i>This object will be purged after policy copy and install.</i> Add web forward servers to a list to form a server group. Optionally assign weights to each server.

~> This resource is a sub resource for variable `server_list` of resource `fmgdevice_webproxy_forwardservergroup`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `forward_server_group` - Forward Server Group.

* `name` - Forward server name.
* `weight` - Optionally assign a weight of the forwarding server for weighted load balancing (1 - 100, default = 10).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WebProxy ForwardServerGroupServerList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "forward_server_group=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webproxy_forwardservergroup_serverlist.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

