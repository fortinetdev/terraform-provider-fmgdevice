---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_casb_profile_saasapplication_advancedtenantcontrol"
description: |-
  <i>This object will be purged after policy copy and install.</i> CASB profile advanced tenant control.
---

# fmgdevice_casb_profile_saasapplication_advancedtenantcontrol
<i>This object will be purged after policy copy and install.</i> CASB profile advanced tenant control.

~> This resource is a sub resource for variable `advanced_tenant_control` of resource `fmgdevice_casb_profile_saasapplication`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `attribute`: `fmgdevice_casb_profile_saasapplication_advancedtenantcontrol_attribute`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.
* `saas_application` - Saas Application.

* `attribute` - Attribute. The structure of `attribute` block is documented below.
* `name` - CASB advanced tenant control name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `attribute` block supports:

* `input` - CASB extend user input value.
* `name` - CASB extend user input name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Casb ProfileSaasApplicationAdvancedTenantControl can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE", "saas_application=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_casb_profile_saasapplication_advancedtenantcontrol.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

