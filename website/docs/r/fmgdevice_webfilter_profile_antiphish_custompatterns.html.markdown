---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webfilter_profile_antiphish_custompatterns"
description: |-
  <i>This object will be purged after policy copy and install.</i> Custom username and password regex patterns.
---

# fmgdevice_webfilter_profile_antiphish_custompatterns
<i>This object will be purged after policy copy and install.</i> Custom username and password regex patterns.

~> This resource is a sub resource for variable `custom_patterns` of resource `fmgdevice_webfilter_profile_antiphish`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `category` - Category that the pattern matches. Valid values: `username`, `password`.

* `pattern` - Target pattern.
* `type` - Pattern will be treated either as a regex pattern or literal string. Valid values: `regex`, `literal`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{pattern}}.

## Import

Webfilter ProfileAntiphishCustomPatterns can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webfilter_profile_antiphish_custompatterns.labelname {{pattern}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

