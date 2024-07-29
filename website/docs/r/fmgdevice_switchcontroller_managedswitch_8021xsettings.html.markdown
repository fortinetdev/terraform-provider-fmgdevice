---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_managedswitch_8021xsettings"
description: |-
  Configuration method to edit FortiSwitch 802.1X global settings.
---

# fmgdevice_switchcontroller_managedswitch_8021xsettings
Configuration method to edit FortiSwitch 802.1X global settings.

~> This resource is a sub resource for variable `n802_1x_settings` of resource `fmgdevice_switchcontroller_managedswitch`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `managed_switch` - Managed Switch.

* `link_down_auth` - Authentication state to set if a link is down. Valid values: `set-unauth`, `no-action`.

* `local_override` - Enable to override global 802.1X settings on individual FortiSwitches. Valid values: `disable`, `enable`.

* `mab_reauth` - Enable or disable MAB reauthentication settings. Valid values: `disable`, `enable`.

* `mac_called_station_delimiter` - MAC called station delimiter (default = hyphen). Valid values: `hyphen`, `single-hyphen`, `colon`, `none`.

* `mac_calling_station_delimiter` - MAC calling station delimiter (default = hyphen). Valid values: `hyphen`, `single-hyphen`, `colon`, `none`.

* `mac_case` - MAC case (default = lowercase). Valid values: `uppercase`, `lowercase`.

* `mac_password_delimiter` - MAC authentication password delimiter (default = hyphen). Valid values: `hyphen`, `single-hyphen`, `colon`, `none`.

* `mac_username_delimiter` - MAC authentication username delimiter (default = hyphen). Valid values: `hyphen`, `single-hyphen`, `colon`, `none`.

* `max_reauth_attempt` - Maximum number of authentication attempts (0 - 15, default = 3).
* `reauth_period` - Reauthentication time interval (1 - 1440 min, default = 60, 0 = disable).
* `tx_period` - 802.1X Tx period (seconds, default=30).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController ManagedSwitch8021XSettings can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "managed_switch=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_managedswitch_8021xsettings.labelname SwitchControllerManagedSwitch8021XSettings
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

