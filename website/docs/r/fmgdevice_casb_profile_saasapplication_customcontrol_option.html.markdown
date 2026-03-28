---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_casb_profile_saasapplication_customcontrol_option"
description: |-
  <i>This object will be purged after policy copy and install.</i> CASB custom control option.
---

# fmgdevice_casb_profile_saasapplication_customcontrol_option
<i>This object will be purged after policy copy and install.</i> CASB custom control option.

~> This resource is a sub resource for variable `option` of resource `fmgdevice_casb_profile_saasapplication_customcontrol`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.
* `saas_application` - Saas Application.
* `custom_control` - Custom Control.

* `name` - CASB custom control option name.
* `user_input` - CASB custom control user input.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Casb ProfileSaasApplicationCustomControlOption can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE", "saas_application=YOUR_VALUE", "custom_control=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_casb_profile_saasapplication_customcontrol_option.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

