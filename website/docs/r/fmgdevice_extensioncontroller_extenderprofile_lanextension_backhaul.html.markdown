---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_extensioncontroller_extenderprofile_lanextension_backhaul"
description: |-
  LAN extension backhaul tunnel configuration.
---

# fmgdevice_extensioncontroller_extenderprofile_lanextension_backhaul
LAN extension backhaul tunnel configuration.

~> This resource is a sub resource for variable `backhaul` of resource `fmgdevice_extensioncontroller_extenderprofile_lanextension`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `extender_profile` - Extender Profile.

* `name` - FortiExtender LAN extension backhaul name.
* `port` - FortiExtender uplink port. Valid values: `wan`, `lte1`, `lte2`, `port1`, `port2`, `port3`, `port4`, `port5`, `sfp`.

* `role` - FortiExtender uplink port. Valid values: `primary`, `secondary`.

* `weight` - WRR weight parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ExtensionController ExtenderProfileLanExtensionBackhaul can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "extender_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_extensioncontroller_extenderprofile_lanextension_backhaul.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

