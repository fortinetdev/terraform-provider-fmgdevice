---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webfilter_profile_ftgdwf_risk"
description: |-
  <i>This object will be purged after policy copy and install.</i> FortiGuard risk level settings.
---

# fmgdevice_webfilter_profile_ftgdwf_risk
<i>This object will be purged after policy copy and install.</i> FortiGuard risk level settings.

~> This resource is a sub resource for variable `risk` of resource `fmgdevice_webfilter_profile_ftgdwf`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `action` - Action to take for matches. Valid values: `block`, `monitor`.

* `fosid` - ID number.
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `risk_level` - Risk level to be examined.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Webfilter ProfileFtgdWfRisk can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webfilter_profile_ftgdwf_risk.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

