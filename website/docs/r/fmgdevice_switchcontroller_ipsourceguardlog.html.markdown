---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_ipsourceguardlog"
description: |-
  Configure FortiSwitch ip source guard log.
---

# fmgdevice_switchcontroller_ipsourceguardlog
Configure FortiSwitch ip source guard log.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `log_violations` - Enable/Disable log violations for IP source guard logging. Valid values: `disable`, `enable`.

* `violation_timer` - IP source gurad log violation timer in seconds (0 - 1500, default = 0).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController IpSourceGuardLog can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_ipsourceguardlog.labelname SwitchControllerIpSourceGuardLog
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

