---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_filefilter_profile_rules"
description: |-
  <i>This object will be purged after policy copy and install.</i> File filter rules.
---

# fmgdevice_filefilter_profile_rules
<i>This object will be purged after policy copy and install.</i> File filter rules.

~> This resource is a sub resource for variable `rules` of resource `fmgdevice_filefilter_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

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

FileFilter ProfileRules can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_filefilter_profile_rules.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

