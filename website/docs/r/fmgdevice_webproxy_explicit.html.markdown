---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webproxy_explicit"
description: |-
  Configure explicit Web proxy settings.
---

# fmgdevice_webproxy_explicit
Configure explicit Web proxy settings.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `pac_policy`: `fmgdevice_webproxy_explicit_pacpolicy`



## Example Usage

```hcl
resource "fmgdevice_webproxy_explicit" "trname" {
  client_cert          = "enable"
  empty_cert_action    = "accept"
  ftp_incoming_port    = ["your own value"]
  ftp_over_http        = "enable"
  http_connection_mode = "multiplex"
  device_name          = var.device_name # not required if setting is at provider
  device_vdom          = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `client_cert` - Enable/disable to request client certificate. Valid values: `disable`, `enable`.

* `empty_cert_action` - Action of an empty client certificate. Valid values: `block`, `accept`, `accept-unmanageable`.

* `ftp_incoming_port` - Accept incoming FTP-over-HTTP requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).
* `ftp_over_http` - Enable to proxy FTP-over-HTTP sessions sent from a web browser. Valid values: `disable`, `enable`.

* `http_connection_mode` - HTTP connection mode (default = static). Valid values: `static`, `multiplex`, `serverpool`.

* `http_incoming_port` - Accept incoming HTTP requests on one or more ports (0 - 65535, default = 8080).
* `https_incoming_port` - Accept incoming HTTPS requests on one or more ports (0 - 65535, default = 0, use the same as HTTP).
* `https_replacement_message` - Enable/disable sending the client a replacement message for HTTPS requests. Valid values: `disable`, `enable`.

* `incoming_ip` - Restrict the explicit HTTP proxy to only accept sessions from this IP address. An interface must have this IP address.
* `incoming_ip6` - Restrict the explicit web proxy to only accept sessions from this IPv6 address. An interface must have this IPv6 address.
* `ipv6_status` - Enable/disable allowing an IPv6 web proxy destination in policies and all IPv6 related entries in this command. Valid values: `disable`, `enable`.

* `message_upon_server_error` - Enable/disable displaying a replacement message when a server error is detected. Valid values: `disable`, `enable`.

* `outgoing_ip` - Outgoing HTTP requests will have this IP address as their source address. An interface must have this IP address.
* `outgoing_ip6` - Outgoing HTTP requests will leave this IPv6. Multiple interfaces can be specified. Interfaces must have these IPv6 addresses.
* `pac_file_data` - PAC file contents enclosed in quotes (maximum of 256K bytes).
* `pac_file_name` - Pac file name.
* `pac_file_server_port` - Port number that PAC traffic from client web browsers uses to connect to the explicit web proxy (0 - 65535, default = 0; use the same as HTTP).
* `pac_file_server_status` - Enable/disable Proxy Auto-Configuration (PAC) for users of this explicit proxy profile. Valid values: `disable`, `enable`.

* `pac_file_through_https` - Enable/disable to get Proxy Auto-Configuration (PAC) through HTTPS. Valid values: `disable`, `enable`.

* `pac_file_url` - Pac-File-Url.
* `pac_policy` - Pac-Policy. The structure of `pac_policy` block is documented below.
* `pref_dns_result` - Prefer resolving addresses using the configured IPv4 or IPv6 DNS server (default = ipv4). Valid values: `ipv4`, `ipv6`, `ipv4-strict`, `ipv6-strict`.

* `realm` - Authentication realm used to identify the explicit web proxy (maximum of 63 characters).
* `sec_default_action` - Accept or deny explicit web proxy sessions when no web proxy firewall policy exists. Valid values: `deny`, `accept`.

* `secure_web_proxy` - Enable/disable/require the secure web proxy for HTTP and HTTPS session. Valid values: `disable`, `enable`, `secure`.

* `secure_web_proxy_cert` - Name of certificates for secure web proxy.
* `socks` - Enable/disable the SOCKS proxy. Valid values: `disable`, `enable`.

* `socks_incoming_port` - Accept incoming SOCKS proxy requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).
* `ssl_algorithm` - Relative strength of encryption algorithms accepted in HTTPS deep scan: high, medium, or low. Valid values: `high`, `medium`, `low`.

* `ssl_dh_bits` - Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.

* `status` - Enable/disable the explicit Web proxy for HTTP and HTTPS session. Valid values: `disable`, `enable`.

* `strict_guest` - Enable/disable strict guest user checking by the explicit web proxy. Valid values: `disable`, `enable`.

* `trace_auth_no_rsp` - Enable/disable logging timed-out authentication requests. Valid values: `disable`, `enable`.

* `unknown_http_version` - How to handle HTTP sessions that do not comply with HTTP 0.9, 1.0, or 1.1. Valid values: `best-effort`, `reject`, `tunnel`.

* `user_agent_detect` - Enable/disable to detect device type by HTTP user-agent if no client certificate provided. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `pac_policy` block supports:

* `comments` - Optional comments.
* `dstaddr` - Destination address objects.
* `pac_file_data` - PAC file contents enclosed in quotes (maximum of 256K bytes).
* `pac_file_name` - Pac file name.
* `policyid` - Policy ID.
* `srcaddr` - Source address objects.
* `srcaddr6` - Source address6 objects.
* `status` - Enable/disable policy. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

WebProxy Explicit can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webproxy_explicit.labelname WebProxyExplicit
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

