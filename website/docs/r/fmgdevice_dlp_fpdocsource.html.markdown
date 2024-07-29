---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_dlp_fpdocsource"
description: |-
  Create a DLP fingerprint database by allowing the FortiGate to access a file server containing files from which to create fingerprints.
---

# fmgdevice_dlp_fpdocsource
Create a DLP fingerprint database by allowing the FortiGate to access a file server containing files from which to create fingerprints.

## Example Usage

```hcl
resource "fmgdevice_dlp_fpdocsource" "trname" {
  date          = 10
  file_path     = "your own value"
  file_pattern  = "your own value"
  keep_modified = "disable"
  name          = "your own value"
  device_name   = var.device_name # not required if setting is at provider
  device_vdom   = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `date` - Day of the month on which to scan the server (1 - 31).
* `file_path` - Path on the server to the fingerprint files (max 119 characters).
* `file_pattern` - Files matching this pattern on the server are fingerprinted. Optionally use the * and ? wildcards.
* `keep_modified` - Enable so that when a file is changed on the server the FortiGate keeps the old fingerprint and adds a new fingerprint to the database. Valid values: `disable`, `enable`.

* `name` - Name of the DLP fingerprint database.
* `password` - Password required to log into the file server.
* `period` - Frequency for which the FortiGate checks the server for new or changed files. Valid values: `none`, `daily`, `weekly`, `monthly`.

* `remove_deleted` - Enable to keep the fingerprint database up to date when a file is deleted from the server. Valid values: `disable`, `enable`.

* `scan_on_creation` - Enable to keep the fingerprint database up to date when a file is added or changed on the server. Valid values: `disable`, `enable`.

* `scan_subdirectories` - Enable/disable scanning subdirectories to find files to create fingerprints from. Valid values: `disable`, `enable`.

* `sensitivity` - Select a sensitivity or threat level for matches with this fingerprint database. Add sensitivities using sensitivity.
* `server` - IPv4 or IPv6 address of the server.
* `server_type` - Protocol used to communicate with the file server. Currently only Samba (SMB) servers are supported. Valid values: `samba`.

* `tod_hour` - Hour of the day on which to scan the server (0 - 23, default = 1).
* `tod_min` - Minute of the hour on which to scan the server (0 - 59).
* `username` - User name required to log into the file server.
* `vdom` - Select the VDOM that can communicate with the file server. Valid values: `mgmt`, `current`.

* `weekday` - Day of the week on which to scan the server. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Dlp FpDocSource can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_dlp_fpdocsource.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

