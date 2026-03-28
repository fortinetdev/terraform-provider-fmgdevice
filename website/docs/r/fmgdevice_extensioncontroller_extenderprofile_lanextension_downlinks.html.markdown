---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_extensioncontroller_extenderprofile_lanextension_downlinks"
description: |-
  Config FortiExtender downlink interface for LAN extension.
---

# fmgdevice_extensioncontroller_extenderprofile_lanextension_downlinks
Config FortiExtender downlink interface for LAN extension.

~> This resource is a sub resource for variable `downlinks` of resource `fmgdevice_extensioncontroller_extenderprofile_lanextension`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `extender_profile` - Extender Profile.

* `name` - FortiExtender LAN extension downlink config entry name.
* `port` - FortiExtender LAN extension downlink port. Valid values: `port1`, `port2`, `port3`, `port4`, `port5`, `lan1`, `lan2`, `lan`.

* `pvid` - FortiExtender LAN extension downlink PVID (1 - 4089).
* `type` - FortiExtender LAN extension downlink type [port/vap]. Valid values: `port`, `vap`.

* `vap` - FortiExtender LAN extension downlink vap.
* `vids` - FortiExtender LAN extension downlink VIDs.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ExtensionController ExtenderProfileLanExtensionDownlinks can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "extender_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_extensioncontroller_extenderprofile_lanextension_downlinks.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

