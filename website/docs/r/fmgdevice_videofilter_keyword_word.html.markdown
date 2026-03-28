---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_videofilter_keyword_word"
description: |-
  <i>This object will be purged after policy copy and install.</i> List of keywords.
---

# fmgdevice_videofilter_keyword_word
<i>This object will be purged after policy copy and install.</i> List of keywords.

~> This resource is a sub resource for variable `word` of resource `fmgdevice_videofilter_keyword`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `keyword` - Keyword.

* `comment` - Comment.
* `name` - Name.
* `pattern_type` - Pattern type. Valid values: `wildcard`, `regex`.

* `status` - Enable(consider)/disable(ignore) this keyword. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Videofilter KeywordWord can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "keyword=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_videofilter_keyword_word.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

