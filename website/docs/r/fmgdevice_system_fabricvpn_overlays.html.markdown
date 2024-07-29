---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_fabricvpn_overlays"
description: |-
  Local overlay interfaces table.
---

# fmgdevice_system_fabricvpn_overlays
Local overlay interfaces table.

~> This resource is a sub resource for variable `overlays` of resource `fmgdevice_system_fabricvpn`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_fabricvpn_overlays" "trname" {
  bgp_neighbor       = ["your own value"]
  bgp_neighbor_group = ["your own value"]
  bgp_neighbor_range = ["your own value"]
  bgp_network        = ["your own value"]
  interface          = ["port2"]
  name               = "your own value"
  device_name        = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `bgp_neighbor` - Underlying BGP neighbor entry.
* `bgp_neighbor_group` - Underlying BGP neighbor group entry.
* `bgp_neighbor_range` - Underlying BGP neighbor range entry.
* `bgp_network` - Underlying BGP network.
* `interface` - Underlying interface name.
* `ipsec_phase1` - IPsec interface.
* `name` - Overlay name.
* `overlay_policy` - The overlay policy to allow ADVPN thru traffic.
* `overlay_tunnel_block` - IPv4 address and subnet mask for the overlay tunnel , syntax: X.X.X.X/24.
* `remote_gw` - IP address of the hub gateway (Set by hub).
* `route_policy` - Underlying router policy.
* `sdwan_member` - Reference to SD-WAN member entry.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System FabricVpnOverlays can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_fabricvpn_overlays.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

