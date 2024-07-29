---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_networkmonitorsettings"
description: |-
  Configure network monitor settings.
---

# fmgdevice_switchcontroller_networkmonitorsettings
Configure network monitor settings.

## Example Usage

```hcl
resource "fmgdevice_switchcontroller_networkmonitorsettings" "trname" {
  network_monitoring = "disable"
  device_name        = var.device_name # not required if setting is at provider
  device_vdom        = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `network_monitoring` - Enable/disable passive gathering of information by FortiSwitch units concerning other network devices. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController NetworkMonitorSettings can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_networkmonitorsettings.labelname SwitchControllerNetworkMonitorSettings
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

