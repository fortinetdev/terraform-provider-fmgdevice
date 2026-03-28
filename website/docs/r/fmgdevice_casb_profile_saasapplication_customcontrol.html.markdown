---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_casb_profile_saasapplication_customcontrol"
description: |-
  <i>This object will be purged after policy copy and install.</i> CASB profile custom control.
---

# fmgdevice_casb_profile_saasapplication_customcontrol
<i>This object will be purged after policy copy and install.</i> CASB profile custom control.

~> This resource is a sub resource for variable `custom_control` of resource `fmgdevice_casb_profile_saasapplication`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `attribute_filter`: `fmgdevice_casb_profile_saasapplication_customcontrol_attributefilter`
>- `option`: `fmgdevice_casb_profile_saasapplication_customcontrol_option`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.
* `saas_application` - Saas Application.

* `attribute_filter` - Attribute-Filter. The structure of `attribute_filter` block is documented below.
* `name` - CASB custom control user activity name.
* `option` - Option. The structure of `option` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `attribute_filter` block supports:

* `action` - CASB access rule tenant control action. Valid values: `block`, `monitor`, `bypass`.

* `attribute_match` - CASB access rule tenant match.
* `id` - CASB tenant control ID.

The `option` block supports:

* `name` - CASB custom control option name.
* `user_input` - CASB custom control user input.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Casb ProfileSaasApplicationCustomControl can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE", "saas_application=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_casb_profile_saasapplication_customcontrol.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

