---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_user_radius_accountingserver"
description: |-
  <i>This object will be purged after policy copy and install.</i> Additional accounting servers.
---

# fmgdevice_user_radius_accountingserver
<i>This object will be purged after policy copy and install.</i> Additional accounting servers.

~> This resource is a sub resource for variable `accounting_server` of resource `fmgdevice_user_radius`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `radius` - Radius.

* `fosid` - ID (0 - 4294967295).
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `port` - RADIUS accounting port number.
* `secret` - Secret key.
* `server` - Server CN domain name or IP address.
* `source_ip` - Source IP address for communications to the RADIUS server.
* `status` - Status. Valid values: `disable`, `enable`.

* `vrf_select` - VRF ID used for connection to server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

User RadiusAccountingServer can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "radius=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_user_radius_accountingserver.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

