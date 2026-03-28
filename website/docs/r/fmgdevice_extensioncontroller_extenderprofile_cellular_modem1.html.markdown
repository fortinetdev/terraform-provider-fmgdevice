---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_extensioncontroller_extenderprofile_cellular_modem1"
description: |-
  Configuration options for modem 1.
---

# fmgdevice_extensioncontroller_extenderprofile_cellular_modem1
Configuration options for modem 1.

~> This resource is a sub resource for variable `modem1` of resource `fmgdevice_extensioncontroller_extenderprofile_cellular`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `auto_switch`: `fmgdevice_extensioncontroller_extenderprofile_cellular_modem1_autoswitch`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `extender_profile` - Extender Profile.

* `auto_switch` - Auto-Switch. The structure of `auto_switch` block is documented below.
* `conn_status` - Conn-Status.
* `default_sim` - Default SIM selection. Valid values: `sim1`, `sim2`, `carrier`, `cost`.

* `gps` - FortiExtender GPS enable/disable. Valid values: `disable`, `enable`.

* `modem_id` - Modem ID.
* `multiple_pdn` - Multiple-PDN enable/disable. Valid values: `disable`, `enable`.

* `pdn1_dataplan` - PDN1-dataplan.
* `pdn2_dataplan` - PDN2-dataplan.
* `pdn3_dataplan` - PDN3-dataplan.
* `pdn4_dataplan` - PDN4-dataplan.
* `preferred_carrier` - Preferred carrier.
* `redundant_intf` - Redundant interface.
* `redundant_mode` - FortiExtender mode. Valid values: `disable`, `enable`.

* `sim1_pin` - SIM #1 PIN status. Valid values: `disable`, `enable`.

* `sim1_pin_code` - SIM #1 PIN password.
* `sim2_pin` - SIM #2 PIN status. Valid values: `disable`, `enable`.

* `sim2_pin_code` - SIM #2 PIN password.

The `auto_switch` block supports:

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

ExtensionController ExtenderProfileCellularModem1 can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "extender_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_extensioncontroller_extenderprofile_cellular_modem1.labelname ExtensionControllerExtenderProfileCellularModem1
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

