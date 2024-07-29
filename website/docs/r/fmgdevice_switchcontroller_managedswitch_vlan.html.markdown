---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_managedswitch_vlan"
description: |-
  Configure VLAN assignment priority.
---

# fmgdevice_switchcontroller_managedswitch_vlan
Configure VLAN assignment priority.

~> This resource is a sub resource for variable `vlan` of resource `fmgdevice_switchcontroller_managedswitch`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_managedswitch_vlan" "trname" {
  assignment_priority = 10
  vlan_name           = ["your own value"]
  device_name         = var.device_name # not required if setting is at provider
  device_vdom         = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `managed_switch` - Managed Switch.

* `assignment_priority` - 802.1x Radius (Tunnel-Private-Group-Id) VLANID assign-by-name priority. A smaller value has a higher priority.
* `vlan_name` - VLAN name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{vlan_name}}.

## Import

SwitchController ManagedSwitchVlan can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "managed_switch=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_managedswitch_vlan.labelname {{vlan_name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

