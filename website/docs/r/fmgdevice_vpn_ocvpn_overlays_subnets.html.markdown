---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_ocvpn_overlays_subnets"
description: |-
  Internal subnets to register with OCVPN service.
---

# fmgdevice_vpn_ocvpn_overlays_subnets
Internal subnets to register with OCVPN service.

~> This resource is a sub resource for variable `subnets` of resource `fmgdevice_vpn_ocvpn_overlays`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_vpn_ocvpn_overlays_subnets" "trname" {
  fosid       = 10
  interface   = ["port2"]
  subnet      = ["your own value"]
  type        = "subnet"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `overlays` - Overlays.

* `fosid` - ID.
* `interface` - LAN interface.
* `subnet` - IPv4 address and subnet mask.
* `type` - Subnet type. Valid values: `subnet`, `interface`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Vpn OcvpnOverlaysSubnets can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "overlays=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_ocvpn_overlays_subnets.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

