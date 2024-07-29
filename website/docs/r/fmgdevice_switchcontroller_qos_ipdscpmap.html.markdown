---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_qos_ipdscpmap"
description: |-
  Configure FortiSwitch QoS IP precedence/DSCP.
---

# fmgdevice_switchcontroller_qos_ipdscpmap
Configure FortiSwitch QoS IP precedence/DSCP.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `map`: `fmgdevice_switchcontroller_qos_ipdscpmap_map`



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_qos_ipdscpmap" "trname" {
  description = "your own value"
  map {
    cos_queue     = 10
    diffserv      = ["AF43"]
    ip_precedence = ["priority"]
    name          = "your own value"
    value         = "your own value"
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

* `description` - Description of the ip-dscp map name.
* `map` - Map. The structure of `map` block is documented below.
* `name` - Dscp map name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `map` block supports:

* `cos_queue` - COS queue number.
* `diffserv` - Differentiated service. Valid values: `CS0`, `CS1`, `AF11`, `AF12`, `AF13`, `CS2`, `AF21`, `AF22`, `AF23`, `CS3`, `AF31`, `AF32`, `AF33`, `CS4`, `AF41`, `AF42`, `AF43`, `CS5`, `EF`, `CS6`, `CS7`.

* `ip_precedence` - IP Precedence. Valid values: `network-control`, `internetwork-control`, `critic-ecp`, `flashoverride`, `flash`, `immediate`, `priority`, `routine`.

* `name` - Dscp mapping entry name.
* `value` - Raw values of DSCP (0 - 63).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController QosIpDscpMap can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_qos_ipdscpmap.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

