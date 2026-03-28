---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_emailfilter_mheader_entries"
description: |-
  Spam filter mime header content.
---

# fmgdevice_emailfilter_mheader_entries
Spam filter mime header content.

~> This resource is a sub resource for variable `entries` of resource `fmgdevice_emailfilter_mheader`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `mheader` - Mheader.

* `action` - Mark spam or good. Valid values: `spam`, `clear`.

* `fieldbody` - Pattern for the header field body.
* `fieldname` - Pattern for header field name.
* `fosid` - Mime header entry ID.
* `pattern_type` - Wildcard pattern or regular expression. Valid values: `wildcard`, `regexp`.

* `status` - Enable/disable status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Emailfilter MheaderEntries can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "mheader=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_emailfilter_mheader_entries.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

