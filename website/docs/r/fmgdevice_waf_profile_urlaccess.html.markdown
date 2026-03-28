---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_waf_profile_urlaccess"
description: |-
  <i>This object will be purged after policy copy and install.</i> URL access list.
---

# fmgdevice_waf_profile_urlaccess
<i>This object will be purged after policy copy and install.</i> URL access list.

~> This resource is a sub resource for variable `url_access` of resource `fmgdevice_waf_profile`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `access_pattern`: `fmgdevice_waf_profile_urlaccess_accesspattern`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `access_pattern` - Access-Pattern. The structure of `access_pattern` block is documented below.
* `action` - Action. Valid values: `bypass`, `permit`, `block`.

* `address` - Host address.
* `fosid` - URL access ID.
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `access_pattern` block supports:

* `id` - URL access pattern ID.
* `negate` - Enable/disable match negation. Valid values: `disable`, `enable`.

* `pattern` - URL pattern.
* `regex` - Enable/disable regular expression based pattern match. Valid values: `disable`, `enable`.

* `srcaddr` - Source address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Waf ProfileUrlAccess can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_waf_profile_urlaccess.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

