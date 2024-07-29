---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_managedswitch_dhcpsnoopingstaticclient"
description: |-
  Configure FortiSwitch DHCP snooping static clients.
---

# fmgdevice_switchcontroller_managedswitch_dhcpsnoopingstaticclient
Configure FortiSwitch DHCP snooping static clients.

~> This resource is a sub resource for variable `dhcp_snooping_static_client` of resource `fmgdevice_switchcontroller_managedswitch`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_managedswitch_dhcpsnoopingstaticclient" "trname" {
  ip          = "your own value"
  mac         = "your own value"
  name        = "your own value"
  port        = "your own value"
  vlan        = ["your own value"]
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `managed_switch` - Managed Switch.

* `ip` - Client static IP address.
* `mac` - Client MAC address.
* `name` - Client name.
* `port` - Interface name.
* `vlan` - VLAN name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController ManagedSwitchDhcpSnoopingStaticClient can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "managed_switch=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_managedswitch_dhcpsnoopingstaticclient.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

