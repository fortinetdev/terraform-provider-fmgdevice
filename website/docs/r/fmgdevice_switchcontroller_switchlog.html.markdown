---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_switchlog"
description: |-
  Configure FortiSwitch logging (logs are transferred to and inserted into FortiGate event log).
---

# fmgdevice_switchcontroller_switchlog
Configure FortiSwitch logging (logs are transferred to and inserted into FortiGate event log).

## Example Usage

```hcl
resource "fmgdevice_switchcontroller_switchlog" "trname" {
  severity    = "critical"
  status      = "enable"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `severity` - Severity of FortiSwitch logs that are added to the FortiGate event log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.

* `status` - Enable/disable adding FortiSwitch logs to FortiGate event log. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController SwitchLog can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_switchlog.labelname SwitchControllerSwitchLog
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

