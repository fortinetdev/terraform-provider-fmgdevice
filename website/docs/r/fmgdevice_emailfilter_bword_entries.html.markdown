---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_emailfilter_bword_entries"
description: |-
  Spam filter banned word.
---

# fmgdevice_emailfilter_bword_entries
Spam filter banned word.

~> This resource is a sub resource for variable `entries` of resource `fmgdevice_emailfilter_bword`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `bword` - Bword.

* `action` - Mark spam or good. Valid values: `spam`, `clear`.

* `fosid` - Banned word entry ID.
* `language` - Language for the banned word. Valid values: `western`, `simch`, `trach`, `japanese`, `korean`, `french`, `thai`, `spanish`.

* `pattern` - Pattern for the banned word.
* `pattern_type` - Wildcard pattern or regular expression. Valid values: `wildcard`, `regexp`.

* `score` - Score value.
* `status` - Enable/disable status. Valid values: `disable`, `enable`.

* `where` - Component of the email to be scanned. Valid values: `all`, `subject`, `body`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Emailfilter BwordEntries can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "bword=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_emailfilter_bword_entries.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

