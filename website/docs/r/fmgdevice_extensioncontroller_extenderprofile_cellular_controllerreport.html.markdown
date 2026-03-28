---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_extensioncontroller_extenderprofile_cellular_controllerreport"
description: |-
  FortiExtender controller report configuration.
---

# fmgdevice_extensioncontroller_extenderprofile_cellular_controllerreport
FortiExtender controller report configuration.

~> This resource is a sub resource for variable `controller_report` of resource `fmgdevice_extensioncontroller_extenderprofile_cellular`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `extender_profile` - Extender Profile.

* `interval` - Controller report interval.
* `signal_threshold` - Controller report signal threshold.
* `status` - FortiExtender controller report status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ExtensionController ExtenderProfileCellularControllerReport can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "extender_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_extensioncontroller_extenderprofile_cellular_controllerreport.labelname ExtensionControllerExtenderProfileCellularControllerReport
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

