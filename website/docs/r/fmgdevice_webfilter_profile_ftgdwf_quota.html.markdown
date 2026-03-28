---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webfilter_profile_ftgdwf_quota"
description: |-
  <i>This object will be purged after policy copy and install.</i> FortiGuard traffic quota settings.
---

# fmgdevice_webfilter_profile_ftgdwf_quota
<i>This object will be purged after policy copy and install.</i> FortiGuard traffic quota settings.

~> This resource is a sub resource for variable `quota` of resource `fmgdevice_webfilter_profile_ftgdwf`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `category` - FortiGuard categories to apply quota to (category action must be set to monitor).
* `duration` - Duration of quota.
* `fosid` - ID number.
* `override_replacemsg` - Override replacement message.
* `type` - Quota type. Valid values: `time`, `traffic`.

* `unit` - Traffic quota unit of measurement. Valid values: `B`, `KB`, `MB`, `GB`.

* `value` - Traffic quota value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Webfilter ProfileFtgdWfQuota can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webfilter_profile_ftgdwf_quota.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

