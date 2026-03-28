---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_waf_profile_constraint_maxcookie"
description: |-
  <i>This object will be purged after policy copy and install.</i> Maximum number of cookies in HTTP request.
---

# fmgdevice_waf_profile_constraint_maxcookie
<i>This object will be purged after policy copy and install.</i> Maximum number of cookies in HTTP request.

~> This resource is a sub resource for variable `max_cookie` of resource `fmgdevice_waf_profile_constraint`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `action` - Action. Valid values: `allow`, `block`.

* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `max_cookie` - Maximum number of cookies in HTTP request (0 to 2147483647).
* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Enable/disable the constraint. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Waf ProfileConstraintMaxCookie can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_waf_profile_constraint_maxcookie.labelname WafProfileConstraintMaxCookie
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

