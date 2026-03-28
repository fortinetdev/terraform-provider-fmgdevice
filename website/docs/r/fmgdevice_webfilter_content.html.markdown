---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webfilter_content"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure Web filter banned word table.
---

# fmgdevice_webfilter_content
<i>This object will be purged after policy copy and install.</i> Configure Web filter banned word table.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `entries`: `fmgdevice_webfilter_content_entries`



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

* `action` - Block or exempt word when a match is found. Valid values: `exempt`, `block`.

* `lang` - Lang. Valid values: `western`, `simch`, `trach`, `japanese`, `korean`, `french`, `thai`, `spanish`, `cyrillic`.

* `name` - Banned word.
* `pattern_type` - Banned word pattern type: wildcard pattern or Perl regular expression. Valid values: `wildcard`, `regexp`.

* `score` - Score, to be applied every time the word appears on a web page (0 - 4294967295, default = 10).
* `status` - Enable/disable banned word. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Webfilter Content can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webfilter_content.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

