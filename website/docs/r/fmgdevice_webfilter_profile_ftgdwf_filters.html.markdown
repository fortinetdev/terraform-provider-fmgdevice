---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webfilter_profile_ftgdwf_filters"
description: |-
  <i>This object will be purged after policy copy and install.</i> FortiGuard filters.
---

# fmgdevice_webfilter_profile_ftgdwf_filters
<i>This object will be purged after policy copy and install.</i> FortiGuard filters.

~> This resource is a sub resource for variable `filters` of resource `fmgdevice_webfilter_profile_ftgdwf`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `action` - Action to take for matches. Valid values: `block`, `monitor`, `warning`, `authenticate`.

* `auth_usr_grp` - Groups with permission to authenticate.
* `category` - Categories and groups the filter examines.
* `fosid` - ID number.
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `override_replacemsg` - Override replacement message.
* `warn_duration` - Duration of warnings.
* `warning_duration_type` - Re-display warning after closing browser or after a timeout. Valid values: `session`, `timeout`.

* `warning_prompt` - Warning prompts in each category or each domain. Valid values: `per-domain`, `per-category`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Webfilter ProfileFtgdWfFilters can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webfilter_profile_ftgdwf_filters.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

