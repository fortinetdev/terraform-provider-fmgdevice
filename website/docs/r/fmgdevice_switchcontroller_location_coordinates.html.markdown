---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_location_coordinates"
description: |-
  Configure location GPS coordinates.
---

# fmgdevice_switchcontroller_location_coordinates
Configure location GPS coordinates.

~> This resource is a sub resource for variable `coordinates` of resource `fmgdevice_switchcontroller_location`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_location_coordinates" "trname" {
  altitude      = "your own value"
  altitude_unit = "m"
  datum         = "NAD83/MLLW"
  latitude      = "your own value"
  longitude     = "your own value"
  device_name   = var.device_name # not required if setting is at provider
  device_vdom   = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `location` - Location.

* `altitude` - Plus or minus floating point number. For example, 117.47.
* `altitude_unit` - Configure the unit for which the altitude is to (m = meters, f = floors of a building). Valid values: `m`, `f`.

* `datum` - WGS84, NAD83, NAD83/MLLW. Valid values: `WGS84`, `NAD83`, `NAD83/MLLW`.

* `latitude` - Floating point starting with +/- or ending with (N or S). For example, +/-16.67 or 16.67N.
* `longitude` - Floating point starting with +/- or ending with (N or S). For example, +/-26.789 or 26.789E.
* `parent_key` - Parent-Key.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController LocationCoordinates can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "location=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_location_coordinates.labelname SwitchControllerLocationCoordinates
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

