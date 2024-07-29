---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_intercontroller_intercontrollerpeer"
description: |-
  Fast failover peer wireless controller list.
---

# fmgdevice_wirelesscontroller_intercontroller_intercontrollerpeer
Fast failover peer wireless controller list.

~> This resource is a sub resource for variable `inter_controller_peer` of resource `fmgdevice_wirelesscontroller_intercontroller`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_intercontroller_intercontrollerpeer" "trname" {
  fosid         = 10
  peer_ip       = "your own value"
  peer_port     = 10
  peer_priority = "primary"
  device_name   = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `fosid` - ID.
* `peer_ip` - Peer wireless controller's IP address.
* `peer_port` - Port used by the wireless controller's for inter-controller communications (1024 - 49150, default = 5246).
* `peer_priority` - Peer wireless controller's priority (primary or secondary, default = primary). Valid values: `primary`, `secondary`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

WirelessController InterControllerInterControllerPeer can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_intercontroller_intercontrollerpeer.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

