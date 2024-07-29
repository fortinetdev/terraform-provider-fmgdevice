---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_hotspot20_qosmap_dscprange"
description: |-
  Differentiated Services Code Point (DSCP) ranges.
---

# fmgdevice_wirelesscontroller_hotspot20_qosmap_dscprange
Differentiated Services Code Point (DSCP) ranges.

~> This resource is a sub resource for variable `dscp_range` of resource `fmgdevice_wirelesscontroller_hotspot20_qosmap`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_hotspot20_qosmap_dscprange" "trname" {
  high        = 10
  index       = 10
  low         = 10
  up          = 10
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `qos_map` - Qos Map.

* `high` - DSCP high value.
* `index` - DSCP range index.
* `low` - DSCP low value.
* `up` - User priority.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{index}}.

## Import

WirelessController Hotspot20QosMapDscpRange can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "qos_map=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_hotspot20_qosmap_dscprange.labelname {{index}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

