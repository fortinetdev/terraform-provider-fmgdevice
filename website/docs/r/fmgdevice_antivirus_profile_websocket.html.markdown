---
subcategory: "Antivirus"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_antivirus_profile_websocket"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure WEBSOCKET AntiVirus options.
---

# fmgdevice_antivirus_profile_websocket
<i>This object will be purged after policy copy and install.</i> Configure WEBSOCKET AntiVirus options.

~> This resource is a sub resource for variable `websocket` of resource `fmgdevice_antivirus_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `archive_block` - Select the archive types to block. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `timeout`.

* `archive_log` - Select the archive types to log. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `timeout`.

* `av_scan` - Enable/disable AntiVirus scan service. Valid values: `disable`, `monitor`, `block`.

* `emulator` - Enable/disable the virus emulator. Valid values: `disable`, `enable`.

* `external_blocklist` - Enable/disable external-blocklist. Analyzes files including the content of archives. Valid values: `disable`, `monitor`, `block`.

* `fortindr` - Enable/disable scanning of files by FortiNDR. Valid values: `disable`, `monitor`, `block`.

* `fortisandbox` - Enable/disable scanning of files by FortiSandbox. Valid values: `disable`, `monitor`, `block`.

* `malware_stream` - Enable/disable 0-day malware-stream scanning. Analyzes files including the content of archives. Valid values: `disable`, `monitor`, `block`.

* `outbreak_prevention` - Enable/disable virus outbreak prevention service. Valid values: `disable`, `monitor`, `block`.

* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Antivirus ProfileWebsocket can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_antivirus_profile_websocket.labelname AntivirusProfileWebsocket
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

