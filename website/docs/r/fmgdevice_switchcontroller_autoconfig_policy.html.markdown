---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_autoconfig_policy"
description: |-
  Policy definitions which can define the behavior on auto configured interfaces.
---

# fmgdevice_switchcontroller_autoconfig_policy
Policy definitions which can define the behavior on auto configured interfaces.

## Example Usage

```hcl
resource "fmgdevice_switchcontroller_autoconfig_policy" "trname" {
  igmp_flood_report  = "disable"
  igmp_flood_traffic = "enable"
  name               = "your own value"
  poe_status         = "enable"
  qos_policy         = ["your own value"]
  device_name        = var.device_name # not required if setting is at provider
  device_vdom        = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `igmp_flood_report` - Enable/disable IGMP flood report. Valid values: `disable`, `enable`.

* `igmp_flood_traffic` - Enable/disable IGMP flood traffic. Valid values: `disable`, `enable`.

* `name` - Auto-config policy name.
* `poe_status` - Enable/disable PoE status. Valid values: `disable`, `enable`.

* `qos_policy` - Auto-Config QoS policy.
* `storm_control_policy` - Auto-Config storm control policy.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController AutoConfigPolicy can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_autoconfig_policy.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

