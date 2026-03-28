---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_dlp_filepattern_entries"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure file patterns used by DLP blocking.
---

# fmgdevice_dlp_filepattern_entries
<i>This object will be purged after policy copy and install.</i> Configure file patterns used by DLP blocking.

~> This resource is a sub resource for variable `entries` of resource `fmgdevice_dlp_filepattern`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `filepattern` - Filepattern.

* `file_type` - Select a file type. Valid values: `unknown`, `exe`, `elf`, `bat`, `javascript`, `html`, `hta`, `msoffice`, `gzip`, `rar`, `tar`, `lzh`, `upx`, `zip`, `cab`, `bzip2`, `bzip`, `activemime`, `mime`, `hlp`, `arj`, `base64`, `binhex`, `uue`, `fsg`, `aspack`, `genscript`, `shellscript`, `petite`, `jpeg`, `gif`, `tiff`, `png`, `bmp`, `msi`, `mpeg`, `mov`, `mp3`, `wma`, `wav`, `pdf`, `avi`, `rm`, `torrent`, `hibun`, `7z`, `xz`, `msofficex`, `mach-o`, `dmg`, `.net`, `xar`, `chm`, `iso`, `crx`, `flac`, `jad`, `class`, `cod`, `sis`, `registry`, `hwp`, `rpm`, `c/cpp`, `pfile`, `lzip`, `wasm`, `sylk`, `dll`, `jnlp`, `python`.

* `filter_type` - Filter by file name pattern or by file type. Valid values: `pattern`, `type`.

* `pattern` - Add a file name pattern.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{pattern}}.

## Import

Dlp FilepatternEntries can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "filepattern=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_dlp_filepattern_entries.labelname {{pattern}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

