---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_extensioncontroller_extenderprofile_cellular_smsnotification"
description: |-
  FortiExtender cellular SMS notification configuration.
---

# fmgdevice_extensioncontroller_extenderprofile_cellular_smsnotification
FortiExtender cellular SMS notification configuration.

~> This resource is a sub resource for variable `sms_notification` of resource `fmgdevice_extensioncontroller_extenderprofile_cellular`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `alert`: `fmgdevice_extensioncontroller_extenderprofile_cellular_smsnotification_alert`
>- `receiver`: `fmgdevice_extensioncontroller_extenderprofile_cellular_smsnotification_receiver`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `extender_profile` - Extender Profile.

* `alert` - Alert. The structure of `alert` block is documented below.
* `receiver` - Receiver. The structure of `receiver` block is documented below.
* `status` - FortiExtender SMS notification status. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `alert` block supports:

* `data_exhausted` - Display string when data exhausted.
* `fgt_backup_mode_switch` - Display string when FortiGate backup mode switched.
* `low_signal_strength` - Display string when signal strength is low.
* `mode_switch` - Display string when mode is switched.
* `os_image_fallback` - Display string when falling back to a previous OS image.
* `session_disconnect` - Display string when session disconnected.
* `system_reboot` - Display string when system rebooted.

The `receiver` block supports:

* `alert` - Alert multi-options. Valid values: `system-reboot`, `data-exhausted`, `session-disconnect`, `low-signal-strength`, `mode-switch`, `os-image-fallback`, `fgt-backup-mode-switch`.

* `name` - FortiExtender SMS notification receiver name.
* `phone_number` - Receiver phone number. Format: [+][country code][area code][local phone number]. For example, +16501234567.
* `status` - SMS notification receiver status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ExtensionController ExtenderProfileCellularSmsNotification can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "extender_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_extensioncontroller_extenderprofile_cellular_smsnotification.labelname ExtensionControllerExtenderProfileCellularSmsNotification
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

