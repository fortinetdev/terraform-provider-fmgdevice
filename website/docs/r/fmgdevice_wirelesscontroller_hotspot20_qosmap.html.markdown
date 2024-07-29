---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_hotspot20_qosmap"
description: |-
  Configure QoS map set.
---

# fmgdevice_wirelesscontroller_hotspot20_qosmap
Configure QoS map set.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `dscp_except`: `fmgdevice_wirelesscontroller_hotspot20_qosmap_dscpexcept`
>- `dscp_range`: `fmgdevice_wirelesscontroller_hotspot20_qosmap_dscprange`



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_hotspot20_qosmap" "trname" {
  dscp_except {
    dscp  = 10
    index = 10
    up    = 10
  }

  dscp_range {
    high  = 10
    index = 10
    low   = 10
    up    = 10
  }

  name        = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `dscp_except` - Dscp-Except. The structure of `dscp_except` block is documented below.
* `dscp_range` - Dscp-Range. The structure of `dscp_range` block is documented below.
* `name` - QOS-MAP name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dscp_except` block supports:

* `dscp` - DSCP value.
* `index` - DSCP exception index.
* `up` - User priority.

The `dscp_range` block supports:

* `high` - DSCP high value.
* `index` - DSCP range index.
* `low` - DSCP low value.
* `up` - User priority.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController Hotspot20QosMap can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_hotspot20_qosmap.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

