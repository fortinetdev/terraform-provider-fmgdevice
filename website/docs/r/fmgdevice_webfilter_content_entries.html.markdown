---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webfilter_content_entries"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure banned word entries.
---

# fmgdevice_webfilter_content_entries
<i>This object will be purged after policy copy and install.</i> Configure banned word entries.

~> This resource is a sub resource for variable `entries` of resource `fmgdevice_webfilter_content`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `content` - Content.

* `action` - Block or exempt word when a match is found. Valid values: `exempt`, `block`.

* `lang` - Lang. Valid values: `western`, `simch`, `trach`, `japanese`, `korean`, `french`, `thai`, `spanish`, `cyrillic`.

* `name` - Banned word.
* `pattern_type` - Banned word pattern type: wildcard pattern or Perl regular expression. Valid values: `wildcard`, `regexp`.

* `score` - Score, to be applied every time the word appears on a web page (0 - 4294967295, default = 10).
* `status` - Enable/disable banned word. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Webfilter ContentEntries can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "content=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webfilter_content_entries.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

