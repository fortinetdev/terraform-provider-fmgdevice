---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_fabricvpn"
description: |-
  Setup for self orchestrated fabric auto discovery VPN.
---

# fmgdevice_system_fabricvpn
Setup for self orchestrated fabric auto discovery VPN.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `advertised_subnets`: `fmgdevice_system_fabricvpn_advertisedsubnets`
>- `overlays`: `fmgdevice_system_fabricvpn_overlays`



## Example Usage

```hcl
resource "fmgdevice_system_fabricvpn" "trname" {
  advertised_subnets {
    access           = "inbound"
    bgp_network      = ["your own value"]
    firewall_address = ["your own value"]
    fosid            = 10
    policies         = ["your own value"]
    prefix           = ["your own value"]
  }

  bgp_as                 = 10
  branch_name            = "your own value"
  health_checks          = ["your own value"]
  loopback_address_block = ["your own value"]
  device_name            = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `advertised_subnets` - Advertised-Subnets. The structure of `advertised_subnets` block is documented below.
* `bgp_as` - BGP Router AS number, valid from 1 to 4294967295.
* `branch_name` - Branch name.
* `health_checks` - Underlying health checks.
* `loopback_address_block` - IPv4 address and subnet mask for hub's loopback address, syntax: X.X.X.X/24.
* `loopback_advertised_subnet` - Loopback advertised subnet reference.
* `loopback_interface` - Loopback interface.
* `overlays` - Overlays. The structure of `overlays` block is documented below.
* `policy_rule` - Policy creation rule. Valid values: `health-check`, `manual`, `auto`.

* `populated` - Populated the setting in tables.
* `psksecret` - Pre-shared secret for ADVPN.
* `sdwan_zone` - Reference to created SD-WAN zone.
* `status` - Enable/disable Fabric VPN. Valid values: `disable`, `enable`.

* `sync_mode` - Setting synchronised by fabric or manual. Valid values: `disable`, `enable`.

* `vpn_role` - Fabric VPN role. Valid values: `hub`, `spoke`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `advertised_subnets` block supports:

* `access` - Access policy direction. Valid values: `inbound`, `bidirectional`.

* `bgp_network` - Underlying BGP network.
* `firewall_address` - Underlying firewall address.
* `id` - ID.
* `policies` - Underlying policies.
* `prefix` - Network prefix.

The `overlays` block supports:

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
* `id` - an identifier for the resource.

## Import

System FabricVpn can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_fabricvpn.labelname SystemFabricVpn
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

