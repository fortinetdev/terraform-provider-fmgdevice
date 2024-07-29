---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_location_elinnumber"
description: |-
  Configure location ELIN number.
---

# fmgdevice_switchcontroller_location_elinnumber
Configure location ELIN number.

~> This resource is a sub resource for variable `elin_number` of resource `fmgdevice_switchcontroller_location`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_location_elinnumber" "trname" {
  elin_num    = "your own value"
  parent_key  = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `location` - Location.

* `elin_num` - Configure ELIN callback number.
* `parent_key` - Parent-Key.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController LocationElinNumber can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "location=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_location_elinnumber.labelname SwitchControllerLocationElinNumber
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

