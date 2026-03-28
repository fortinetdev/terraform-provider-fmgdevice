---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_extensioncontroller_extenderprofile_cellular_modem1_autoswitch"
description: |-
  FortiExtender auto switch configuration.
---

# fmgdevice_extensioncontroller_extenderprofile_cellular_modem1_autoswitch
FortiExtender auto switch configuration.

~> This resource is a sub resource for variable `auto_switch` of resource `fmgdevice_extensioncontroller_extenderprofile_cellular_modem1`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `extender_profile` - Extender Profile.

* `dataplan` - Automatically switch based on data usage. Valid values: `disable`, `enable`.

* `disconnect` - Auto switch by disconnect. Valid values: `disable`, `enable`.

* `disconnect_period` - Automatically switch based on disconnect period.
* `disconnect_threshold` - Automatically switch based on disconnect threshold.
* `signal` - Automatically switch based on signal strength. Valid values: `disable`, `enable`.

* `switch_back` - Auto switch with switch back multi-options. Valid values: `time`, `timer`.

* `switch_back_time` - Automatically switch over to preferred SIM/carrier at a specified time in UTC (HH:MM).
* `switch_back_timer` - Automatically switch over to preferred SIM/carrier after the given time (3600 - 2147483647 sec).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ExtensionController ExtenderProfileCellularModem1AutoSwitch can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "extender_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_extensioncontroller_extenderprofile_cellular_modem1_autoswitch.labelname ExtensionControllerExtenderProfileCellularModem1AutoSwitch
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

