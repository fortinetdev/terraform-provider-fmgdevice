---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_extensioncontroller_extenderprofile_lanextension_trafficsplitservices"
description: |-
  Config FortiExtender traffic split interface for LAN extension.
---

# fmgdevice_extensioncontroller_extenderprofile_lanextension_trafficsplitservices
Config FortiExtender traffic split interface for LAN extension.

~> This resource is a sub resource for variable `traffic_split_services` of resource `fmgdevice_extensioncontroller_extenderprofile_lanextension`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `extender_profile` - Extender Profile.

* `address` - Address selection.
* `name` - FortiExtender LAN extension tunnel split entry name.
* `service` - Service selection.
* `vsdb` - Set video streaming traffic goes through local WAN [enable/disable]. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ExtensionController ExtenderProfileLanExtensionTrafficSplitServices can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "extender_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_extensioncontroller_extenderprofile_lanextension_trafficsplitservices.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

