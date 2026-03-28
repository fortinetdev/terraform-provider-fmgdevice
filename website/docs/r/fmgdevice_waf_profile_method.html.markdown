---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_waf_profile_method"
description: |-
  <i>This object will be purged after policy copy and install.</i> Method restriction.
---

# fmgdevice_waf_profile_method
<i>This object will be purged after policy copy and install.</i> Method restriction.

~> This resource is a sub resource for variable `method` of resource `fmgdevice_waf_profile`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `method_policy`: `fmgdevice_waf_profile_method_methodpolicy`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `default_allowed_methods` - Methods. Valid values: `delete`, `get`, `head`, `options`, `post`, `put`, `trace`, `others`, `connect`.

* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `method_policy` - Method-Policy. The structure of `method_policy` block is documented below.
* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Status. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `method_policy` block supports:

* `address` - Host address.
* `allowed_methods` - Allowed Methods. Valid values: `delete`, `get`, `head`, `options`, `post`, `put`, `trace`, `others`, `connect`.

* `id` - HTTP method policy ID.
* `pattern` - URL pattern.
* `regex` - Enable/disable regular expression based pattern match. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Waf ProfileMethod can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_waf_profile_method.labelname WafProfileMethod
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

