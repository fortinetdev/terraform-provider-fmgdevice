---
subcategory: "DLP"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_dlp_settings"
description: |-
  Designate logical storage for DLP fingerprint database.
---

# fmgdevice_dlp_settings
Designate logical storage for DLP fingerprint database.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `ocr`: `fmgdevice_dlp_settings_ocr`



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
* `config_builder_timeout` - Maximum time allowed for building a single DLP profile (default 60 seconds).
* `db_mode` - Behavior when the maximum size is reached. Valid values: `stop-adding`, `remove-modified-then-oldest`, `remove-oldest`.

* `ocr` - Ocr. The structure of `ocr` block is documented below.
* `size` - Maximum total size of files within the storage (MB).
* `storage_device` - Storage device name.

The `ocr` block supports:

* `confidence` - Minimum confidence threshold for the OCR converted content to be scanned (0 - 100, default = 80).
* `filetype_ignore_list` - List of file types to be exempt from OCR scanning.
* `max_file_size` - Maximum file size for an image to be a candidate for OCR conversion in kilobytes (0 - 4193280, 0 = unlimited).
* `scan` - Enable/disable OCR conversion of images for DLP content scanning. Valid values: `disable`, `enable`.



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

