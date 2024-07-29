---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_managedswitch_ports_dhcpsnoopoption82override"
description: |-
  Configure DHCP snooping option 82 override.
---

# fmgdevice_switchcontroller_managedswitch_ports_dhcpsnoopoption82override
Configure DHCP snooping option 82 override.

~> This resource is a sub resource for variable `dhcp_snoop_option82_override` of resource `fmgdevice_switchcontroller_managedswitch_ports`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_managedswitch_ports_dhcpsnoopoption82override" "trname" {
  circuit_id  = "your own value"
  remote_id   = "your own value"
  vlan_name   = ["your own value"]
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `managed_switch` - Managed Switch.
* `ports` - Ports.

* `circuit_id` - Circuit ID string.
* `remote_id` - Remote ID string.
* `vlan_name` - DHCP snooping option 82 VLAN.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{vlan_name}}.

## Import

SwitchController ManagedSwitchPortsDhcpSnoopOption82Override can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "managed_switch=YOUR_VALUE", "ports=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_managedswitch_ports_dhcpsnoopoption82override.labelname {{vlan_name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

