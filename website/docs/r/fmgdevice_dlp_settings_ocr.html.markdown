---
subcategory: "DLP"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_dlp_settings_ocr"
description: |-
  Configure settings for optical character recognition (OCR) conversion.
---

# fmgdevice_dlp_settings_ocr
Configure settings for optical character recognition (OCR) conversion.

~> This resource is a sub resource for variable `ocr` of resource `fmgdevice_dlp_settings`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `confidence` - Minimum confidence threshold for the OCR converted content to be scanned (0 - 100, default = 80).
* `filetype_ignore_list` - List of file types to be exempt from OCR scanning.
* `max_file_size` - Maximum file size for an image to be a candidate for OCR conversion in kilobytes (0 - 4193280, 0 = unlimited).
* `scan` - Enable/disable OCR conversion of images for DLP content scanning. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Dlp SettingsOcr can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_dlp_settings_ocr.labelname DlpSettingsOcr
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

