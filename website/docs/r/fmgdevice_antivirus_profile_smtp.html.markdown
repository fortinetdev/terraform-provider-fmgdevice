---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_antivirus_profile_smtp"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure SMTP AntiVirus options.
---

# fmgdevice_antivirus_profile_smtp
<i>This object will be purged after policy copy and install.</i> Configure SMTP AntiVirus options.

~> This resource is a sub resource for variable `smtp` of resource `fmgdevice_antivirus_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `archive_block` - Select the archive types to block. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `timeout`.

* `archive_log` - Select the archive types to log. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `timeout`.

* `av_scan` - Enable AntiVirus scan service. Valid values: `disable`, `monitor`, `block`.

* `content_disarm` - Enable/disable Content Disarm and Reconstruction when performing AntiVirus scan. Valid values: `disable`, `enable`.

* `emulator` - Enable/disable the virus emulator. Valid values: `disable`, `enable`.

* `executables` - Treat Windows executable files as viruses for the purpose of blocking or monitoring. Valid values: `default`, `virus`.

* `options` - Enable/disable SMTP AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `quarantine`, `avmonitor`.

* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives. Valid values: `disable`, `monitor`, `block`.

* `fortiai` - Enable/disable scanning of files by FortiAI server. Valid values: `disable`, `monitor`, `block`.

* `fortindr` - Enable scanning of files by FortiNDR. Valid values: `disable`, `block`, `monitor`.

* `fortisandbox` - Enable scanning of files by FortiSandbox. Valid values: `disable`, `block`, `monitor`.

* `malware_stream` - Enable 0-day malware-stream scanning. Analyzes files including the content of archives. Valid values: `disable`, `monitor`, `block`.

* `outbreak_prevention` - Enable virus outbreak prevention service. Valid values: `disable`, `block`, `monitor`.

* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Antivirus ProfileSmtp can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_antivirus_profile_smtp.labelname AntivirusProfileSmtp
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

