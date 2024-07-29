---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_ssl_client"
description: |-
  Client.
---

# fmgdevice_vpn_ssl_client
Client.

## Example Usage

```hcl
resource "fmgdevice_vpn_ssl_client" "trname" {
  certificate = ["your own value"]
  class_id    = ["your own value"]
  comment     = "your own value"
  distance    = 10
  interface   = ["port2"]
  name        = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `certificate` - Certificate to offer to SSL-VPN server if it requests one.
* `class_id` - Traffic class ID.
* `comment` - Comment.
* `distance` - Distance for routes added by SSL-VPN (1 - 255).
* `interface` - SSL interface to send/receive traffic over.
* `ipv4_subnets` - IPv4 subnets that the client is protecting.
* `ipv6_subnets` - IPv6 subnets that the client is protecting.
* `name` - SSL-VPN tunnel name.
* `peer` - Authenticate peer's certificate with the peer/peergrp.
* `port` - SSL-VPN server port.
* `priority` - Priority for routes added by SSL-VPN (1 - 65535).
* `psk` - Pre-shared secret to authenticate with the server (ASCII string or hexadecimal encoded with a leading 0x).
* `realm` - Realm name configured on SSL-VPN server.
* `server` - IPv4, IPv6 or DNS address of the SSL-VPN server.
* `source_ip` - IPv4 or IPv6 address to use as a source for the SSL-VPN connection to the server.
* `status` - Enable/disable this SSL-VPN client configuration. Valid values: `disable`, `enable`.

* `user` - Username to offer to the peer to authenticate the client.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Vpn SslClient can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_ssl_client.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

