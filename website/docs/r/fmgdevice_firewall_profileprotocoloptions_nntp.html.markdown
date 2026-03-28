---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_profileprotocoloptions_nntp"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure NNTP protocol options.
---

# fmgdevice_firewall_profileprotocoloptions_nntp
<i>This object will be purged after policy copy and install.</i> Configure NNTP protocol options.

~> This resource is a sub resource for variable `nntp` of resource `fmgdevice_firewall_profileprotocoloptions`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile_protocol_options` - Profile Protocol Options.

* `inspect_all` - Enable/disable the inspection of all ports for the protocol. Valid values: `disable`, `enable`.

* `options` - One or more options that can be applied to the session. Valid values: `oversize`, `splice`.

* `oversize_limit` - Maximum in-memory file size that can be scanned (MB).
* `ports` - Ports to scan for content (1 - 65535, default = 119).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `disable`, `enable`.

* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `disable`, `enable`.

* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `disable`, `enable`.

* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (MB).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Firewall ProfileProtocolOptionsNntp can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile_protocol_options=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_profileprotocoloptions_nntp.labelname FirewallProfileProtocolOptionsNntp
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

