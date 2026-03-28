---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_filefilter_profile"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure file-filter profiles.
---

# fmgdevice_filefilter_profile
<i>This object will be purged after policy copy and install.</i> Configure file-filter profiles.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `rules`: `fmgdevice_filefilter_profile_rules`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comment` - Comment.
* `extended_log` - Enable/disable file-filter extended logging. Valid values: `disable`, `enable`.

* `feature_set` - Flow/proxy feature set. Valid values: `proxy`, `flow`.

* `log` - Enable/disable file-filter logging. Valid values: `disable`, `enable`.

* `name` - Profile name.
* `replacemsg_group` - Replacement message group.
* `rules` - Rules. The structure of `rules` block is documented below.
* `scan_archive_contents` - Enable/disable archive contents scan. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `rules` block supports:

* `action` - Action taken for matched file. Valid values: `log-only`, `block`.

* `comment` - Comment.
* `direction` - Traffic direction (HTTP, FTP, SSH, CIFS, and MAPI only). Valid values: `any`, `incoming`, `outgoing`.

* `file_type` - Select file type.
* `name` - File-filter rule name.
* `password_protected` - Match password-protected files. Valid values: `any`, `yes`.

* `protocol` - Protocols to apply rule to. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `mapi`, `cifs`, `ssh`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

FileFilter Profile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_filefilter_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

