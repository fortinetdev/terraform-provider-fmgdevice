---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webproxy_explicitproxy"
description: |-
  <i>This object will be purged after policy copy and install.</i>
---

# fmgdevice_webproxy_explicitproxy
<i>This object will be purged after policy copy and install.</i>

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `detect_https_in_http_request` - Detect-Https-In-Http-Request. Valid values: `disable`, `enable`.

* `client_cert` - Client-Cert. Valid values: `disable`, `enable`.

* `dns_mode` - Dns-Mode. Valid values: `recursive`, `non-recursive`, `forward-only`.

* `dstport_from_incoming` - Dstport-From-Incoming. Valid values: `disable`, `enable`.

* `empty_cert_action` - Empty-Cert-Action. Valid values: `block`, `accept`, `accept-unmanageable`.

* `ftp_incoming_port` - Ftp-Incoming-Port.
* `ftp_over_http` - Ftp-Over-Http. Valid values: `disable`, `enable`.

* `header_proxy_agent` - Header-Proxy-Agent. Valid values: `disable`, `enable`.

* `http` - Http. Valid values: `disable`, `enable`.

* `http_connection_mode` - Http-Connection-Mode. Valid values: `static`, `multiplex`, `serverpool`.

* `http_incoming_port` - Http-Incoming-Port.
* `https_incoming_port` - Https-Incoming-Port.
* `incoming_ip` - Incoming-Ip.
* `incoming_ip6` - Incoming-Ip6.
* `interface` - Interface.
* `ipv6_status` - Ipv6-Status. Valid values: `disable`, `enable`.

* `learn_dst_from_sni` - Learn-Dst-From-Sni. Valid values: `disable`, `enable`.

* `name` - Name.
* `pac_file_data` - Pac-File-Data.
* `pac_file_name` - Pac-File-Name.
* `pac_file_server_port` - Pac-File-Server-Port.
* `pac_file_server_status` - Pac-File-Server-Status. Valid values: `disable`, `enable`.

* `pac_file_through_https` - Pac-File-Through-Https. Valid values: `disable`, `enable`.

* `pac_file_url` - Pac-File-Url.
* `pref_dns_result` - Pref-Dns-Result. Valid values: `ipv4`, `ipv6`, `ipv4-strict`, `ipv6-strict`.

* `realm` - Realm.
* `return_to_sender` - Return-To-Sender. Valid values: `disable`, `enable`.

* `sec_default_action` - Sec-Default-Action. Valid values: `deny`, `accept`.

* `secure_web_proxy` - Secure-Web-Proxy. Valid values: `disable`, `enable`, `secure`.

* `secure_web_proxy_cert` - Secure-Web-Proxy-Cert.
* `socks` - Socks. Valid values: `disable`, `enable`.

* `socks_incoming_port` - Socks-Incoming-Port.
* `ssl_algorithm` - Ssl-Algorithm. Valid values: `high`, `low`, `medium`.

* `ssl_dh_bits` - Ssl-Dh-Bits. Valid values: `768`, `1024`, `1536`, `2048`.

* `status` - Status. Valid values: `disable`, `enable`.

* `unknown_http_version` - Unknown-Http-Version. Valid values: `best-effort`, `reject`.

* `user_agent_detect` - User-Agent-Detect. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WebProxy ExplicitProxy can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webproxy_explicitproxy.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

