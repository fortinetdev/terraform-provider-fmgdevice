---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_isolator_setting"
description: |-
  Device IsolatorSetting
---

# fmgdevice_isolator_setting
Device IsolatorSetting

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `default_isolator_profile` - Default-Isolator-Profile.
* `defective_session` - Defective-Session. Valid values: `block`, `pass`, `use-default-profile`.

* `unmatched_session` - Unmatched-Session. Valid values: `block`, `pass`, `use-default-profile`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Isolator Setting can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_isolator_setting.labelname IsolatorSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

