---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_dlp_filepattern"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure file patterns used by DLP blocking.
---

# fmgdevice_dlp_filepattern
<i>This object will be purged after policy copy and install.</i> Configure file patterns used by DLP blocking.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `entries`: `fmgdevice_dlp_filepattern_entries`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comment` - Optional comments.
* `entries` - Entries. The structure of `entries` block is documented below.
* `fosid` - ID.
* `name` - Name of table containing the file pattern list.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `entries` block supports:

* `file_type` - Select a file type. Valid values: `unknown`, `exe`, `elf`, `bat`, `javascript`, `html`, `hta`, `msoffice`, `gzip`, `rar`, `tar`, `lzh`, `upx`, `zip`, `cab`, `bzip2`, `bzip`, `activemime`, `mime`, `hlp`, `arj`, `base64`, `binhex`, `uue`, `fsg`, `aspack`, `genscript`, `shellscript`, `petite`, `jpeg`, `gif`, `tiff`, `png`, `bmp`, `msi`, `mpeg`, `mov`, `mp3`, `wma`, `wav`, `pdf`, `avi`, `rm`, `torrent`, `hibun`, `7z`, `xz`, `msofficex`, `mach-o`, `dmg`, `.net`, `xar`, `chm`, `iso`, `crx`, `flac`, `jad`, `class`, `cod`, `sis`, `registry`, `hwp`, `rpm`, `c/cpp`, `pfile`, `lzip`, `wasm`, `sylk`, `dll`, `jnlp`, `python`.

* `filter_type` - Filter by file name pattern or by file type. Valid values: `pattern`, `type`.

* `pattern` - Add a file name pattern.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Dlp Filepattern can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_dlp_filepattern.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

