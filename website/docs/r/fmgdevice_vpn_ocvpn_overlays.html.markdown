---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_ocvpn_overlays"
description: |-
  Network overlays to register with Overlay Controller VPN service.
---

# fmgdevice_vpn_ocvpn_overlays
Network overlays to register with Overlay Controller VPN service.

~> This resource is a sub resource for variable `overlays` of resource `fmgdevice_vpn_ocvpn`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `subnets`: `fmgdevice_vpn_ocvpn_overlays_subnets`



## Example Usage

```hcl
resource "fmgdevice_vpn_ocvpn_overlays" "trname" {
  assign_ip     = "enable"
  fosid         = 10
  inter_overlay = "deny"
  ipv4_end_ip   = "your own value"
  ipv4_start_ip = "your own value"
  overlay_name  = "your own value"
  device_name   = var.device_name # not required if setting is at provider
  device_vdom   = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `assign_ip` - Assign-Ip. Valid values: `disable`, `enable`.

* `fosid` - ID.
* `inter_overlay` - Allow or deny traffic from other overlays. Valid values: `deny`, `allow`.

* `ipv4_end_ip` - Ipv4-End-Ip.
* `ipv4_start_ip` - Ipv4-Start-Ip.
* `name` - Overlay name.
* `overlay_name` - Overlay name.
* `subnets` - Subnets. The structure of `subnets` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `subnets` block supports:

* `id` - ID.
* `interface` - LAN interface.
* `subnet` - IPv4 address and subnet mask.
* `type` - Subnet type. Valid values: `subnet`, `interface`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{overlay_name}}.

## Import

Vpn OcvpnOverlays can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_ocvpn_overlays.labelname {{overlay_name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

