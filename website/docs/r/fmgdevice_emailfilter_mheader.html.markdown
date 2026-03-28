---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_emailfilter_mheader"
description: |-
  Configure AntiSpam MIME header.
---

# fmgdevice_emailfilter_mheader
Configure AntiSpam MIME header.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `entries`: `fmgdevice_emailfilter_mheader_entries`



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

* `fieldbody` - Pattern for the header field body.
* `fieldname` - Pattern for header field name.
* `id` - Mime header entry ID.
* `pattern_type` - Wildcard pattern or regular expression. Valid values: `wildcard`, `regexp`.

* `status` - Enable/disable status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Emailfilter Mheader can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_emailfilter_mheader.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

