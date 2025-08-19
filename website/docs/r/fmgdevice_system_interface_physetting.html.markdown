---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_interface_physetting"
description: |-
  PHY settings
---

# fmgdevice_system_interface_physetting
PHY settings

~> This resource is a sub resource for variable `phy_setting` of resource `fmgdevice_system_interface`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `interface` - Interface.

* `signal_ok_threshold_value` - Signal-ok-threshold value(0 - 12).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System InterfacePhySetting can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "interface=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_interface_physetting.labelname SystemInterfacePhySetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

