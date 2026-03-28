---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_videofilter_keyword"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure video filter keywords.
---

# fmgdevice_videofilter_keyword
<i>This object will be purged after policy copy and install.</i> Configure video filter keywords.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `word`: `fmgdevice_videofilter_keyword_word`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comment` - Comment.
* `fosid` - ID.
* `match` - Keyword matching logic. Valid values: `or`, `and`.

* `name` - Name.
* `word` - Word. The structure of `word` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `word` block supports:

* `comment` - Comment.
* `name` - Name.
* `pattern_type` - Pattern type. Valid values: `wildcard`, `regex`.

* `status` - Enable(consider)/disable(ignore) this keyword. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Videofilter Keyword can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_videofilter_keyword.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

