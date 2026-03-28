---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_profileprotocoloptions_http"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure HTTP protocol options.
---

# fmgdevice_firewall_profileprotocoloptions_http
<i>This object will be purged after policy copy and install.</i> Configure HTTP protocol options.

~> This resource is a sub resource for variable `http` of resource `fmgdevice_firewall_profileprotocoloptions`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile_protocol_options` - Profile Protocol Options.

* `address_ip_rating` - Enable/disable IP based URL rating. Valid values: `disable`, `enable`.

* `block_page_status_code` - Code number returned for blocked HTTP pages (non-FortiGuard only) (100 - 599, default = 403).
* `comfort_amount` - Number of bytes to send in each transmission for client comforting (bytes).
* `comfort_interval` - Interval between successive transmissions of data for client comforting (seconds).
* `fortinet_bar` - Fortinet-Bar. Valid values: `disable`, `enable`.

* `fortinet_bar_port` - Fortinet-Bar-Port.
* `domain_fronting` - Configure HTTP domain fronting (default = block). Valid values: `strict`, `monitor`, `block`, `allow`.

* `h2c` - Enable/disable h2c HTTP connection upgrade. Valid values: `disable`, `enable`.

* `http_09` - Configure action to take upon receipt of HTTP 0.9 request. Valid values: `block`, `allow`.

* `inspect_all` - Enable/disable the inspection of all ports for the protocol. Valid values: `disable`, `enable`.

* `options` - One or more options that can be applied to the session. Valid values: `oversize`, `chunkedbypass`, `clientcomfort`, `servercomfort`.

* `oversize_limit` - Maximum in-memory file size that can be scanned (MB).
* `ports` - Ports to scan for content (1 - 65535, default = 80).
* `post_lang` - ID codes for character sets to be used to convert to UTF-8 for banned words and DLP on HTTP posts (maximum of 5 character sets). Valid values: `jisx0201`, `jisx0208`, `jisx0212`, `gb2312`, `ksc5601-ex`, `euc-jp`, `sjis`, `iso2022-jp`, `iso2022-jp-1`, `iso2022-jp-2`, `euc-cn`, `ces-gbk`, `hz`, `ces-big5`, `euc-kr`, `iso2022-jp-3`, `iso8859-1`, `tis620`, `cp874`, `cp1252`, `cp1251`.

* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `disable`, `enable`.

* `range_block` - Enable/disable blocking of partial downloads. Valid values: `disable`, `enable`.

* `retry_count` - Number of attempts to retry HTTP connection (0 - 100, default = 0).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `disable`, `enable`.

* `ssl_offloaded` - SSL decryption and encryption performed by an external device. Valid values: `no`, `yes`.

* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `disable`, `enable`.

* `stream_based_uncompressed_limit` - Maximum stream-based uncompressed data size that will be scanned in megabytes. Stream-based uncompression used only under certain conditions (unlimited = 0, default = 0).
* `streaming_content_bypass` - Enable/disable bypassing of streaming content from buffering. Valid values: `disable`, `enable`.

* `strip_x_forwarded_for` - Enable/disable stripping of HTTP X-Forwarded-For header. Valid values: `disable`, `enable`.

* `switching_protocols` - Bypass from scanning, or block a connection that attempts to switch protocol. Valid values: `bypass`, `block`.

* `tcp_window_maximum` - Maximum dynamic TCP window size.
* `tcp_window_minimum` - Minimum dynamic TCP window size.
* `tcp_window_size` - Set TCP static window size.
* `tcp_window_type` - TCP window type to use for this protocol. Valid values: `system`, `static`, `dynamic`, `auto-tuning`.

* `tunnel_non_http` - Configure how to process non-HTTP traffic when a profile configured for HTTP traffic accepts a non-HTTP session. Can occur if an application sends non-HTTP traffic using an HTTP destination port. Valid values: `disable`, `enable`.

* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (MB).
* `unknown_content_encoding` - Configure the action the FortiGate unit will take on unknown content-encoding. Valid values: `block`, `inspect`, `bypass`.

* `unknown_http_version` - How to handle HTTP sessions that do not comply with HTTP 0.9, 1.0, or 1.1. Valid values: `best-effort`, `reject`, `tunnel`.

* `verify_dns_for_policy_matching` - Enable/disable verification of DNS for policy matching. Valid values: `disable`, `enable`.

* `dns_protection` - Dns-Protection. Valid values: `disable`, `enable`.

* `encrypted_file` - Encrypted-File. Valid values: `block`, `pass`, `inspect`.

* `encrypted_file_log` - Encrypted-File-Log. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Firewall ProfileProtocolOptionsHttp can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile_protocol_options=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_profileprotocoloptions_http.labelname FirewallProfileProtocolOptionsHttp
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

