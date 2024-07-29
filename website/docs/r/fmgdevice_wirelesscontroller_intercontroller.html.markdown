---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_intercontroller"
description: |-
  Configure inter wireless controller operation.
---

# fmgdevice_wirelesscontroller_intercontroller
Configure inter wireless controller operation.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `inter_controller_peer`: `fmgdevice_wirelesscontroller_intercontroller_intercontrollerpeer`



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_intercontroller" "trname" {
  fast_failover_max     = 10
  fast_failover_wait    = 10
  inter_controller_key  = ["your own value"]
  inter_controller_mode = "l2-roaming"
  inter_controller_peer {
    fosid         = 10
    peer_ip       = "your own value"
    peer_port     = 10
    peer_priority = "primary"
  }

  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `fast_failover_max` - Maximum number of retransmissions for fast failover HA messages between peer wireless controllers (3 - 64, default = 10).
* `fast_failover_wait` - Minimum wait time before an AP transitions from secondary controller to primary controller (10 - 86400 sec, default = 10).
* `inter_controller_key` - Secret key for inter-controller communications.
* `inter_controller_mode` - Configure inter-controller mode (disable, l2-roaming, 1+1, default = disable). Valid values: `disable`, `l2-roaming`, `1+1`.

* `inter_controller_peer` - Inter-Controller-Peer. The structure of `inter_controller_peer` block is documented below.
* `inter_controller_pri` - Configure inter-controller's priority (primary or secondary, default = primary). Valid values: `primary`, `secondary`.

* `l3_roaming` - Enable/disable layer 3 roaming (default = disable). Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `inter_controller_peer` block supports:

* `id` - ID.
* `peer_ip` - Peer wireless controller's IP address.
* `peer_port` - Port used by the wireless controller's for inter-controller communications (1024 - 49150, default = 5246).
* `peer_priority` - Peer wireless controller's priority (primary or secondary, default = primary). Valid values: `primary`, `secondary`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

WirelessController InterController can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_intercontroller.labelname WirelessControllerInterController
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

