---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_hotspot20_h2qpwanmetric"
description: |-
  Configure WAN metrics.
---

# fmgdevice_wirelesscontroller_hotspot20_h2qpwanmetric
Configure WAN metrics.

## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_hotspot20_h2qpwanmetric" "trname" {
  downlink_load             = 10
  downlink_speed            = 10
  link_at_capacity          = "enable"
  link_status               = "up"
  load_measurement_duration = 10
  name                      = "your own value"
  device_name               = var.device_name # not required if setting is at provider
  device_vdom               = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `downlink_load` - Downlink load.
* `downlink_speed` - Downlink speed (in kilobits/s).
* `link_at_capacity` - Link at capacity. Valid values: `disable`, `enable`.

* `link_status` - Link status. Valid values: `down`, `up`, `in-test`.

* `load_measurement_duration` - Load measurement duration (in tenths of a second).
* `name` - WAN metric name.
* `symmetric_wan_link` - WAN link symmetry. Valid values: `asymmetric`, `symmetric`.

* `uplink_load` - Uplink load.
* `uplink_speed` - Uplink speed (in kilobits/s).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController Hotspot20H2QpWanMetric can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_hotspot20_h2qpwanmetric.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

