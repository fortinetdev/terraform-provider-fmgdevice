---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webproxy_wisp"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure Websense Integrated Services Protocol (WISP) servers.
---

# fmgdevice_webproxy_wisp
<i>This object will be purged after policy copy and install.</i> Configure Websense Integrated Services Protocol (WISP) servers.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comment` - Comment.
* `max_connections` - Maximum number of web proxy WISP connections (4 - 4096, default = 64).
* `name` - Server name.
* `outgoing_ip` - WISP outgoing IP address.
* `server_ip` - WISP server IP address.
* `server_port` - WISP server port (1 - 65535, default = 15868).
* `timeout` - Period of time before WISP requests time out (1 - 15 sec, default = 5).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WebProxy Wisp can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webproxy_wisp.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

