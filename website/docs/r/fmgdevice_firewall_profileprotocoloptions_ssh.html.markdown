---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_profileprotocoloptions_ssh"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure SFTP and SCP protocol options.
---

# fmgdevice_firewall_profileprotocoloptions_ssh
<i>This object will be purged after policy copy and install.</i> Configure SFTP and SCP protocol options.

~> This resource is a sub resource for variable `ssh` of resource `fmgdevice_firewall_profileprotocoloptions`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile_protocol_options` - Profile Protocol Options.

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



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Firewall ProfileProtocolOptionsSsh can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile_protocol_options=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_profileprotocoloptions_ssh.labelname FirewallProfileProtocolOptionsSsh
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

