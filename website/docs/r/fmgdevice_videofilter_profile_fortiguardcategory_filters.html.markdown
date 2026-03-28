---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_videofilter_profile_fortiguardcategory_filters"
description: |-
  <i>This object will be purged after policy copy and install.</i>
---

# fmgdevice_videofilter_profile_fortiguardcategory_filters
<i>This object will be purged after policy copy and install.</i>

~> This resource is a sub resource for variable `filters` of resource `fmgdevice_videofilter_profile_fortiguardcategory`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `action` - Action. Valid values: `block`, `monitor`, `allow`.

* `category_id` - Category-Id.
* `fosid` - Id.
* `log` - Log. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Videofilter ProfileFortiguardCategoryFilters can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_videofilter_profile_fortiguardcategory_filters.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

