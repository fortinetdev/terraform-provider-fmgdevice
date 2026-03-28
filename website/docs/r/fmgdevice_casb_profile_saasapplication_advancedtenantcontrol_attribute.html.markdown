---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_casb_profile_saasapplication_advancedtenantcontrol_attribute"
description: |-
  <i>This object will be purged after policy copy and install.</i> CASB advanced tenant control attribute.
---

# fmgdevice_casb_profile_saasapplication_advancedtenantcontrol_attribute
<i>This object will be purged after policy copy and install.</i> CASB advanced tenant control attribute.

~> This resource is a sub resource for variable `attribute` of resource `fmgdevice_casb_profile_saasapplication_advancedtenantcontrol`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.
* `saas_application` - Saas Application.
* `advanced_tenant_control` - Advanced Tenant Control.

* `input` - CASB extend user input value.
* `name` - CASB extend user input name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Casb ProfileSaasApplicationAdvancedTenantControlAttribute can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE", "saas_application=YOUR_VALUE", "advanced_tenant_control=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_casb_profile_saasapplication_advancedtenantcontrol_attribute.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

