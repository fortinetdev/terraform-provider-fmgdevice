---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_sslserver"
description: |-
  Configure SSL servers.
---

# fmgdevice_firewall_sslserver
Configure SSL servers.

## Example Usage

```hcl
resource "fmgdevice_firewall_sslserver" "trname" {
  add_header_x_forwarded_proto = "disable"
  ip                           = "your own value"
  mapped_port                  = 10
  name                         = "your own value"
  port                         = 10
  device_name                  = var.device_name # not required if setting is at provider
  device_vdom                  = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `add_header_x_forwarded_proto` - Enable/disable adding an X-Forwarded-Proto header to forwarded requests. Valid values: `disable`, `enable`.

* `ip` - IPv4 address of the SSL server.
* `mapped_port` - Mapped server service port (1 - 65535, default = 80).
* `name` - Server name.
* `port` - Server service port (1 - 65535, default = 443).
* `ssl_algorithm` - Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.

* `ssl_cert` - List of certificate names to use for SSL connections to this server. (default = "Fortinet_SSL").
* `ssl_client_renegotiation` - Allow or block client renegotiation by server. Valid values: `deny`, `allow`, `secure`.

* `ssl_dh_bits` - Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.

* `ssl_max_version` - Highest SSL/TLS version to negotiate. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.

* `ssl_min_version` - Lowest SSL/TLS version to negotiate. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.

* `ssl_mode` - SSL/TLS mode for encryption and decryption of traffic. Valid values: `half`, `full`.

* `ssl_send_empty_frags` - Enable/disable sending empty fragments to avoid attack on CBC IV. Valid values: `disable`, `enable`.

* `url_rewrite` - Enable/disable rewriting the URL. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall SslServer can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_sslserver.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

