---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_videofilter_profile_filters"
description: |-
  <i>This object will be purged after policy copy and install.</i> YouTube filter entries.
---

# fmgdevice_videofilter_profile_filters
<i>This object will be purged after policy copy and install.</i> YouTube filter entries.

~> This resource is a sub resource for variable `filters` of resource `fmgdevice_videofilter_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `action` - Video filter action. Valid values: `block`, `monitor`, `allow`.

* `category` - FortiGuard category ID.
* `channel` - Channel ID.
* `comment` - Comment.
* `fosid` - ID.
* `keyword` - Video filter keyword ID.
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `type` - Filter type. Valid values: `category`, `channel`, `title`, `description`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Videofilter ProfileFilters can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_videofilter_profile_filters.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

