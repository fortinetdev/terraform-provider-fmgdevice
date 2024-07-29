---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_snmptrapthreshold"
description: |-
  Configure FortiSwitch SNMP trap threshold values globally.
---

# fmgdevice_switchcontroller_snmptrapthreshold
Configure FortiSwitch SNMP trap threshold values globally.

## Example Usage

```hcl
resource "fmgdevice_switchcontroller_snmptrapthreshold" "trname" {
  trap_high_cpu_threshold   = 10
  trap_log_full_threshold   = 10
  trap_low_memory_threshold = 10
  device_name               = var.device_name # not required if setting is at provider
  device_vdom               = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `trap_high_cpu_threshold` - CPU usage when trap is sent.
* `trap_log_full_threshold` - Log disk usage when trap is sent.
* `trap_low_memory_threshold` - Memory usage when trap is sent.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController SnmpTrapThreshold can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_snmptrapthreshold.labelname SwitchControllerSnmpTrapThreshold
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

