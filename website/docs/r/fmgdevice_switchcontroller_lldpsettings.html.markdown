---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_lldpsettings"
description: |-
  Configure FortiSwitch LLDP settings.
---

# fmgdevice_switchcontroller_lldpsettings
Configure FortiSwitch LLDP settings.

## Example Usage

```hcl
resource "fmgdevice_switchcontroller_lldpsettings" "trname" {
  device_detection     = "disable"
  fast_start_interval  = 10
  management_interface = "mgmt"
  tx_hold              = 10
  tx_interval          = 10
  device_name          = var.device_name # not required if setting is at provider
  device_vdom          = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `device_detection` - Enable/disable dynamic detection of LLDP neighbor devices for VLAN assignment. Valid values: `disable`, `enable`.

* `fast_start_interval` - Frequency of LLDP PDU transmission from FortiSwitch for the first 4 packets when the link is up (2 - 5 sec, default = 2, 0 = disable fast start).
* `management_interface` - Primary management interface to be advertised in LLDP and CDP PDUs. Valid values: `internal`, `mgmt`.

* `tx_hold` - Number of tx-intervals before local LLDP data expires (1 - 16, default = 4). Packet TTL is tx-hold * tx-interval.
* `tx_interval` - Frequency of LLDP PDU transmission from FortiSwitch (5 - 4095 sec, default = 30). Packet TTL is tx-hold * tx-interval.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController LldpSettings can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_lldpsettings.labelname SwitchControllerLldpSettings
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

