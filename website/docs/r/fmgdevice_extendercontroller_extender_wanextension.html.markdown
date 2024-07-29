---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_extendercontroller_extender_wanextension"
description: |-
  Device ExtenderControllerExtenderWanExtension
---

# fmgdevice_extendercontroller_extender_wanextension
Device ExtenderControllerExtenderWanExtension

~> This resource is a sub resource for variable `wan_extension` of resource `fmgdevice_extendercontroller_extender`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_extendercontroller_extender_wanextension" "trname" {
  modem1_extension = ["your own value"]
  modem2_extension = ["your own value"]
  device_name      = var.device_name # not required if setting is at provider
  device_vdom      = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `extender` - Extender.

* `modem1_extension` - Modem1-Extension.
* `modem2_extension` - Modem2-Extension.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ExtenderController ExtenderWanExtension can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "extender=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_extendercontroller_extender_wanextension.labelname ExtenderControllerExtenderWanExtension
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

