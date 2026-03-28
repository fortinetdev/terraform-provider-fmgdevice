---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_emailfilter_bword"
description: |-
  Configure AntiSpam banned word list.
---

# fmgdevice_emailfilter_bword
Configure AntiSpam banned word list.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `entries`: `fmgdevice_emailfilter_bword_entries`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comment` - Optional comments.
* `entries` - Entries. The structure of `entries` block is documented below.
* `fosid` - ID.
* `name` - Name of table.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `entries` block supports:

* `action` - Mark spam or good. Valid values: `spam`, `clear`.

* `id` - Banned word entry ID.
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

Emailfilter Bword can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_emailfilter_bword.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

