---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_dlp_settings"
description: |-
  Designate logical storage for DLP fingerprint database.
---

# fmgdevice_dlp_settings
Designate logical storage for DLP fingerprint database.

## Example Usage

```hcl
resource "fmgdevice_dlp_settings" "trname" {
  cache_mem_percent = 10
  chunk_size        = 10
  db_mode           = "remove-oldest"
  size              = 10
  storage_device    = "your own value"
  device_name       = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `cache_mem_percent` - Maximum percentage of available memory allocated to caching (1 - 15).
* `chunk_size` - Maximum fingerprint chunk size. Caution, changing this setting will flush the entire database.
* `db_mode` - Behavior when the maximum size is reached. Valid values: `stop-adding`, `remove-modified-then-oldest`, `remove-oldest`.

* `size` - Maximum total size of files within the storage (MB).
* `storage_device` - Storage device name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Dlp Settings can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_dlp_settings.labelname DlpSettings
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

