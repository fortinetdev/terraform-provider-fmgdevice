---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webproxy_isolatorserver"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure forward-server addresses.
---

# fmgdevice_webproxy_isolatorserver
<i>This object will be purged after policy copy and install.</i> Configure forward-server addresses.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `addr_type` - Address type of the forwarding proxy server: IP or FQDN. Valid values: `fqdn`, `ipv6`, `ip`.

* `comment` - Comment.
* `fqdn` - Forward server Fully Qualified Domain Name (FQDN).
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `ip` - Forward proxy server IP address.
* `ipv6` - Forward proxy server IPv6 address.
* `masquerade` - Enable/disable use of the of the IP address of the outgoing interface as the client IP address (default = enable) Valid values: `disable`, `enable`.

* `name` - Server name.
* `port` - Port number that the forwarding server expects to receive HTTP sessions on (1 - 65535, default = 3128).
* `vrf_select` - VRF ID used for connection to server.
* `ippool` - Ippool.
* `protocol` - Protocol. Valid values: `http`, `socks`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WebProxy IsolatorServer can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webproxy_isolatorserver.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

