---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_managedswitch_igmpsnooping_vlans"
description: |-
  Configure IGMP snooping VLAN.
---

# fmgdevice_switchcontroller_managedswitch_igmpsnooping_vlans
Configure IGMP snooping VLAN.

~> This resource is a sub resource for variable `vlans` of resource `fmgdevice_switchcontroller_managedswitch_igmpsnooping`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_managedswitch_igmpsnooping_vlans" "trname" {
  proxy        = "disable"
  querier      = "disable"
  querier_addr = "your own value"
  version      = 10
  vlan_name    = ["your own value"]
  device_name  = var.device_name # not required if setting is at provider
  device_vdom  = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `managed_switch` - Managed Switch.

* `proxy` - IGMP snooping proxy for the VLAN interface. Valid values: `disable`, `enable`, `global`.

* `querier` - Enable/disable IGMP snooping querier for the VLAN interface. Valid values: `disable`, `enable`.

* `querier_addr` - IGMP snooping querier address.
* `version` - IGMP snooping querying version.
* `vlan_name` - List of FortiSwitch VLANs.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{vlan_name}}.

## Import

SwitchController ManagedSwitchIgmpSnoopingVlans can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "managed_switch=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_managedswitch_igmpsnooping_vlans.labelname {{vlan_name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

