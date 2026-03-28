---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_icap_remoteservergroup_serverlist"
description: |-
  Device IcapRemoteServerGroupServerList
---

# fmgdevice_icap_remoteservergroup_serverlist
Device IcapRemoteServerGroupServerList

~> This resource is a sub resource for variable `server_list` of resource `fmgdevice_icap_remoteservergroup`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `remote_server_group` - Remote Server Group.

* `name` - Name.
* `weight` - Weight.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Icap RemoteServerGroupServerList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "remote_server_group=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_icap_remoteservergroup_serverlist.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

