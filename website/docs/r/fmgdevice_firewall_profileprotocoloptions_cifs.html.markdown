---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_profileprotocoloptions_cifs"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure CIFS protocol options.
---

# fmgdevice_firewall_profileprotocoloptions_cifs
<i>This object will be purged after policy copy and install.</i> Configure CIFS protocol options.

~> This resource is a sub resource for variable `cifs` of resource `fmgdevice_firewall_profileprotocoloptions`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `file_filter`: `fmgdevice_firewall_profileprotocoloptions_cifs_filefilter`
>- `server_keytab`: `fmgdevice_firewall_profileprotocoloptions_cifs_serverkeytab`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile_protocol_options` - Profile Protocol Options.

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
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

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


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Firewall ProfileProtocolOptionsCifs can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile_protocol_options=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_profileprotocoloptions_cifs.labelname FirewallProfileProtocolOptionsCifs
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

