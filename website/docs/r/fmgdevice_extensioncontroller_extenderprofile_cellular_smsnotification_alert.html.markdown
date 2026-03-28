---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_extensioncontroller_extenderprofile_cellular_smsnotification_alert"
description: |-
  SMS alert list.
---

# fmgdevice_extensioncontroller_extenderprofile_cellular_smsnotification_alert
SMS alert list.

~> This resource is a sub resource for variable `alert` of resource `fmgdevice_extensioncontroller_extenderprofile_cellular_smsnotification`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `extender_profile` - Extender Profile.

* `data_exhausted` - Display string when data exhausted.
* `fgt_backup_mode_switch` - Display string when FortiGate backup mode switched.
* `low_signal_strength` - Display string when signal strength is low.
* `mode_switch` - Display string when mode is switched.
* `os_image_fallback` - Display string when falling back to a previous OS image.
* `session_disconnect` - Display string when session disconnected.
* `system_reboot` - Display string when system rebooted.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ExtensionController ExtenderProfileCellularSmsNotificationAlert can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "extender_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_extensioncontroller_extenderprofile_cellular_smsnotification_alert.labelname ExtensionControllerExtenderProfileCellularSmsNotificationAlert
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

