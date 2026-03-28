---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_extensioncontroller_extenderprofile_cellular_smsnotification_receiver"
description: |-
  SMS notification receiver list.
---

# fmgdevice_extensioncontroller_extenderprofile_cellular_smsnotification_receiver
SMS notification receiver list.

~> This resource is a sub resource for variable `receiver` of resource `fmgdevice_extensioncontroller_extenderprofile_cellular_smsnotification`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `extender_profile` - Extender Profile.

* `alert` - Alert multi-options. Valid values: `system-reboot`, `data-exhausted`, `session-disconnect`, `low-signal-strength`, `mode-switch`, `os-image-fallback`, `fgt-backup-mode-switch`.

* `name` - FortiExtender SMS notification receiver name.
* `phone_number` - Receiver phone number. Format: [+][country code][area code][local phone number]. For example, +16501234567.
* `status` - SMS notification receiver status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ExtensionController ExtenderProfileCellularSmsNotificationReceiver can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "extender_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_extensioncontroller_extenderprofile_cellular_smsnotification_receiver.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

