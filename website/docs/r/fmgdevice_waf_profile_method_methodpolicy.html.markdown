---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_waf_profile_method_methodpolicy"
description: |-
  <i>This object will be purged after policy copy and install.</i> HTTP method policy.
---

# fmgdevice_waf_profile_method_methodpolicy
<i>This object will be purged after policy copy and install.</i> HTTP method policy.

~> This resource is a sub resource for variable `method_policy` of resource `fmgdevice_waf_profile_method`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `address` - Host address.
* `allowed_methods` - Allowed Methods. Valid values: `delete`, `get`, `head`, `options`, `post`, `put`, `trace`, `others`, `connect`.

* `fosid` - HTTP method policy ID.
* `pattern` - URL pattern.
* `regex` - Enable/disable regular expression based pattern match. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Waf ProfileMethodMethodPolicy can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_waf_profile_method_methodpolicy.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

