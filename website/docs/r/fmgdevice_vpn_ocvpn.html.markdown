---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_ocvpn"
description: |-
  Configure Overlay Controller VPN settings.
---

# fmgdevice_vpn_ocvpn
Configure Overlay Controller VPN settings.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `forticlient_access`: `fmgdevice_vpn_ocvpn_forticlientaccess`
>- `overlays`: `fmgdevice_vpn_ocvpn_overlays`



## Example Usage

```hcl
resource "fmgdevice_vpn_ocvpn" "trname" {
  auto_discovery               = "enable"
  auto_discovery_shortcut_mode = "independent"
  eap                          = "disable"
  eap_users                    = ["your own value"]
  forticlient_access {
    auth_groups {
      auth_group = ["your own value"]
      name       = "your own value"
      overlays   = ["your own value"]
    }

    psksecret = ["your own value"]
    status    = "enable"
  }

  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `auto_discovery` - Enable/disable auto-discovery shortcuts. Valid values: `disable`, `enable`.

* `auto_discovery_shortcut_mode` - Control deletion of child short-cut tunnels when the parent tunnel goes down. Valid values: `independent`, `dependent`.

* `eap` - Enable/disable EAP client authentication. Valid values: `disable`, `enable`.

* `eap_users` - EAP authentication user group.
* `forticlient_access` - Forticlient-Access. The structure of `forticlient_access` block is documented below.
* `ha_alias` - Hidden HA alias.
* `ip_allocation_block` - Class B subnet reserved for private IP address assignment.
* `multipath` - Enable/disable multipath redundancy. Valid values: `disable`, `enable`.

* `nat` - Enable/disable NAT support. Valid values: `disable`, `enable`.

* `overlays` - Overlays. The structure of `overlays` block is documented below.
* `poll_interval` - Overlay Controller VPN polling interval.
* `role` - Set device role. Valid values: `spoke`, `primary-hub`, `secondary-hub`, `client`.

* `sdwan` - Enable/disable adding OCVPN tunnels to SD-WAN. Valid values: `disable`, `enable`.

* `sdwan_zone` - Set SD-WAN zone.
* `status` - Enable/disable Overlay Controller cloud assisted VPN. Valid values: `disable`, `enable`.

* `wan_interface` - FortiGate WAN interfaces to use with OCVPN.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `forticlient_access` block supports:

* `auth_groups` - Auth-Groups. The structure of `auth_groups` block is documented below.
* `psksecret` - Pre-shared secret for FortiClient PSK authentication (ASCII string or hexadecimal encoded with a leading 0x).
* `status` - Enable/disable FortiClient to access OCVPN networks. Valid values: `disable`, `enable`.


The `auth_groups` block supports:

* `auth_group` - Authentication user group for FortiClient access.
* `name` - Group name.
* `overlays` - OCVPN overlays to allow access to.

The `overlays` block supports:

* `assign_ip` - Assign-Ip. Valid values: `disable`, `enable`.

* `id` - ID.
* `inter_overlay` - Allow or deny traffic from other overlays. Valid values: `deny`, `allow`.

* `ipv4_end_ip` - Ipv4-End-Ip.
* `ipv4_start_ip` - Ipv4-Start-Ip.
* `name` - Overlay name.
* `overlay_name` - Overlay name.
* `subnets` - Subnets. The structure of `subnets` block is documented below.

The `subnets` block supports:

* `id` - ID.
* `interface` - LAN interface.
* `subnet` - IPv4 address and subnet mask.
* `type` - Subnet type. Valid values: `subnet`, `interface`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Vpn Ocvpn can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_ocvpn.labelname VpnOcvpn
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

