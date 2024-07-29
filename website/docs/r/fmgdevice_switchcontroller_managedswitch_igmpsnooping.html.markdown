---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_managedswitch_igmpsnooping"
description: |-
  Configure FortiSwitch IGMP snooping global settings.
---

# fmgdevice_switchcontroller_managedswitch_igmpsnooping
Configure FortiSwitch IGMP snooping global settings.

~> This resource is a sub resource for variable `igmp_snooping` of resource `fmgdevice_switchcontroller_managedswitch`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `vlans`: `fmgdevice_switchcontroller_managedswitch_igmpsnooping_vlans`



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_managedswitch_igmpsnooping" "trname" {
  aging_time              = 10
  flood_unknown_multicast = "enable"
  local_override          = "disable"
  vlans {
    proxy        = "disable"
    querier      = "enable"
    querier_addr = "your own value"
    version      = 10
    vlan_name    = ["your own value"]
  }

  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `managed_switch` - Managed Switch.

* `aging_time` - Maximum time to retain a multicast snooping entry for which no packets have been seen (15 - 3600 sec, default = 300).
* `flood_unknown_multicast` - Enable/disable unknown multicast flooding. Valid values: `disable`, `enable`.

* `local_override` - Enable/disable overriding the global IGMP snooping configuration. Valid values: `disable`, `enable`.

* `vlans` - Vlans. The structure of `vlans` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `vlans` block supports:

* `proxy` - IGMP snooping proxy for the VLAN interface. Valid values: `disable`, `enable`, `global`.

* `querier` - Enable/disable IGMP snooping querier for the VLAN interface. Valid values: `disable`, `enable`.

* `querier_addr` - IGMP snooping querier address.
* `version` - IGMP snooping querying version.
* `vlan_name` - List of FortiSwitch VLANs.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController ManagedSwitchIgmpSnooping can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "managed_switch=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_managedswitch_igmpsnooping.labelname SwitchControllerManagedSwitchIgmpSnooping
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

