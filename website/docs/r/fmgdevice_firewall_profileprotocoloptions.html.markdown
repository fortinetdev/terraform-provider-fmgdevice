---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_profileprotocoloptions"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure protocol options.
---

# fmgdevice_firewall_profileprotocoloptions
<i>This object will be purged after policy copy and install.</i> Configure protocol options.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `cifs`: `fmgdevice_firewall_profileprotocoloptions_cifs`
>- `dns`: `fmgdevice_firewall_profileprotocoloptions_dns`
>- `ftp`: `fmgdevice_firewall_profileprotocoloptions_ftp`
>- `http`: `fmgdevice_firewall_profileprotocoloptions_http`
>- `imap`: `fmgdevice_firewall_profileprotocoloptions_imap`
>- `mail_signature`: `fmgdevice_firewall_profileprotocoloptions_mailsignature`
>- `mapi`: `fmgdevice_firewall_profileprotocoloptions_mapi`
>- `nntp`: `fmgdevice_firewall_profileprotocoloptions_nntp`
>- `pop3`: `fmgdevice_firewall_profileprotocoloptions_pop3`
>- `proxy_redirect`: `fmgdevice_firewall_profileprotocoloptions_proxyredirect`
>- `smtp`: `fmgdevice_firewall_profileprotocoloptions_smtp`
>- `ssh`: `fmgdevice_firewall_profileprotocoloptions_ssh`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `cifs` - Cifs. The structure of `cifs` block is documented below.
* `comment` - Optional comments.
* `dns` - Dns. The structure of `dns` block is documented below.
* `feature_set` - Feature-Set. Valid values: `proxy`, `flow`.

* `ftp` - Ftp. The structure of `ftp` block is documented below.
* `http` - Http. The structure of `http` block is documented below.
* `imap` - Imap. The structure of `imap` block is documented below.
* `mail_signature` - Mail-Signature. The structure of `mail_signature` block is documented below.
* `mapi` - Mapi. The structure of `mapi` block is documented below.
* `name` - Name.
* `nntp` - Nntp. The structure of `nntp` block is documented below.
* `oversize_log` - Enable/disable logging for antivirus oversize file blocking. Valid values: `disable`, `enable`.

* `pop3` - Pop3. The structure of `pop3` block is documented below.
* `replacemsg_group` - Name of the replacement message group to be used.
* `rpc_over_http` - Enable/disable inspection of RPC over HTTP. Valid values: `disable`, `enable`.

* `smtp` - Smtp. The structure of `smtp` block is documented below.
* `ssh` - Ssh. The structure of `ssh` block is documented below.
* `switching_protocols_log` - Enable/disable logging for HTTP/HTTPS switching protocols. Valid values: `disable`, `enable`.

* `proxy_redirect` - Proxy-Redirect. The structure of `proxy_redirect` block is documented below.

The `cifs` block supports:

* `domain_controller` - Domain for which to decrypt CIFS traffic.
* `file_filter` - File-Filter. The structure of `file_filter` block is documented below.
* `options` - One or more options that can be applied to the session. Valid values: `oversize`.

* `oversize_limit` - Maximum in-memory file size that can be scanned (MB).
* `ports` - Ports to scan for content (1 - 65535, default = 445).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `disable`, `enable`.

* `server_credential_type` - CIFS server credential type. Valid values: `none`, `credential-replication`, `credential-keytab`.

* `server_keytab` - Server-Keytab. The structure of `server_keytab` block is documented below.
* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `disable`, `enable`.

* `tcp_window_maximum` - Maximum dynamic TCP window size.
* `tcp_window_minimum` - Minimum dynamic TCP window size.
* `tcp_window_size` - Set TCP static window size.
* `tcp_window_type` - TCP window type to use for this protocol. Valid values: `system`, `static`, `dynamic`, `auto-tuning`.

* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (MB).

The `file_filter` block supports:

* `entries` - Entries. The structure of `entries` block is documented below.
* `log` - Log. Valid values: `disable`, `enable`.

* `status` - Status. Valid values: `disable`, `enable`.


The `entries` block supports:

* `action` - Action. Valid values: `log`, `block`.

* `comment` - Comment.
* `direction` - Direction. Valid values: `any`, `incoming`, `outgoing`.

* `file_type` - File-Type.
* `filter` - Filter.
* `protocol` - Protocol. Valid values: `cifs`.


The `server_keytab` block supports:

* `keytab` - Base64 encoded keytab file containing credential of the server.
* `password` - Password for keytab.
* `principal` - Service principal. For example, host/cifsserver.example.com@example.com.

The `dns` block supports:

* `ports` - Ports to scan for content (1 - 65535, default = 53).
* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `disable`, `enable`.


The `ftp` block supports:

* `comfort_amount` - Number of bytes to send in each transmission for client comforting (bytes).
* `comfort_interval` - Interval between successive transmissions of data for client comforting (seconds).
* `explicit_ftp_tls` - Enable/disable FTP redirection for explicit FTPS. Valid values: `disable`, `enable`.

* `inspect_all` - Enable/disable the inspection of all ports for the protocol. Valid values: `disable`, `enable`.

* `options` - One or more options that can be applied to the session. Valid values: `clientcomfort`, `oversize`, `splice`, `bypass-rest-command`, `bypass-mode-command`.

* `oversize_limit` - Maximum in-memory file size that can be scanned (MB).
* `ports` - Ports to scan for content (1 - 65535, default = 21).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `disable`, `enable`.

* `ssl_offloaded` - SSL decryption and encryption performed by an external device. Valid values: `no`, `yes`.

* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `disable`, `enable`.

* `stream_based_uncompressed_limit` - Maximum stream-based uncompressed data size that will be scanned in megabytes. Stream-based uncompression used only under certain conditions (unlimited = 0, default = 0).
* `tcp_window_maximum` - Maximum dynamic TCP window size.
* `tcp_window_minimum` - Minimum dynamic TCP window size.
* `tcp_window_size` - Set TCP static window size.
* `tcp_window_type` - TCP window type to use for this protocol. Valid values: `system`, `static`, `dynamic`, `auto-tuning`.

* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (MB).

The `http` block supports:

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


The `imap` block supports:

* `inspect_all` - Enable/disable the inspection of all ports for the protocol. Valid values: `disable`, `enable`.

* `options` - One or more options that can be applied to the session. Valid values: `oversize`, `fragmail`.

* `oversize_limit` - Maximum in-memory file size that can be scanned (MB).
* `ports` - Ports to scan for content (1 - 65535, default = 143).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `disable`, `enable`.

* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `disable`, `enable`.

* `ssl_offloaded` - SSL decryption and encryption performed by an external device. Valid values: `no`, `yes`.

* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `disable`, `enable`.

* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (MB).
* `address_ip_rating` - Address-Ip-Rating. Valid values: `disable`, `enable`.


The `mail_signature` block supports:

* `signature` - Email signature to be added to outgoing email (if the signature contains spaces, enclose with quotation marks).
* `status` - Enable/disable adding an email signature to SMTP email messages as they pass through the FortiGate. Valid values: `disable`, `enable`.


The `mapi` block supports:

* `options` - One or more options that can be applied to the session. Valid values: `fragmail`, `oversize`.

* `oversize_limit` - Maximum in-memory file size that can be scanned (MB).
* `ports` - Ports to scan for content (1 - 65535, default = 135).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `disable`, `enable`.

* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `disable`, `enable`.

* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (MB).

The `nntp` block supports:

* `inspect_all` - Enable/disable the inspection of all ports for the protocol. Valid values: `disable`, `enable`.

* `options` - One or more options that can be applied to the session. Valid values: `oversize`, `splice`.

* `oversize_limit` - Maximum in-memory file size that can be scanned (MB).
* `ports` - Ports to scan for content (1 - 65535, default = 119).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `disable`, `enable`.

* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `disable`, `enable`.

* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `disable`, `enable`.

* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (MB).

The `pop3` block supports:

* `inspect_all` - Enable/disable the inspection of all ports for the protocol. Valid values: `disable`, `enable`.

* `options` - One or more options that can be applied to the session. Valid values: `oversize`, `fragmail`.

* `oversize_limit` - Maximum in-memory file size that can be scanned (MB).
* `ports` - Ports to scan for content (1 - 65535, default = 110).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `disable`, `enable`.

* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `disable`, `enable`.

* `ssl_offloaded` - SSL decryption and encryption performed by an external device. Valid values: `no`, `yes`.

* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `disable`, `enable`.

* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (MB).

The `smtp` block supports:

* `inspect_all` - Enable/disable the inspection of all ports for the protocol. Valid values: `disable`, `enable`.

* `options` - One or more options that can be applied to the session. Valid values: `oversize`, `fragmail`, `splice`.

* `oversize_limit` - Maximum in-memory file size that can be scanned (MB).
* `ports` - Ports to scan for content (1 - 65535, default = 25).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `disable`, `enable`.

* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `disable`, `enable`.

* `server_busy` - Enable/disable SMTP server busy when server not available. Valid values: `disable`, `enable`.

* `ssl_offloaded` - SSL decryption and encryption performed by an external device. Valid values: `no`, `yes`.

* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `disable`, `enable`.

* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (MB).

The `ssh` block supports:

* `comfort_amount` - Number of bytes to send in each transmission for client comforting (bytes).
* `comfort_interval` - Interval between successive transmissions of data for client comforting (seconds).
* `options` - One or more options that can be applied to the session. Valid values: `oversize`, `clientcomfort`, `servercomfort`.

* `oversize_limit` - Maximum in-memory file size that can be scanned (MB).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `disable`, `enable`.

* `ssl_offloaded` - SSL decryption and encryption performed by an external device. Valid values: `no`, `yes`.

* `stream_based_uncompressed_limit` - Maximum stream-based uncompressed data size that will be scanned in megabytes. Stream-based uncompression used only under certain conditions (unlimited = 0, default = 0).
* `tcp_window_maximum` - Maximum dynamic TCP window size.
* `tcp_window_minimum` - Minimum dynamic TCP window size.
* `tcp_window_size` - Set TCP static window size.
* `tcp_window_type` - TCP window type to use for this protocol. Valid values: `system`, `static`, `dynamic`, `auto-tuning`.

* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (MB).
* `explicit_ftp_tls` - Explicit-Ftp-Tls. Valid values: `disable`, `enable`.


The `proxy_redirect` block supports:

* `ports` - Ports.
* `status` - Status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall ProfileProtocolOptions can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_profileprotocoloptions.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

