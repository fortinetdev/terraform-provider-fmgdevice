---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_sittunnel"
description: |-
  Configure IPv6 tunnel over IPv4.
---

# fmgdevice_system_sittunnel
Configure IPv6 tunnel over IPv4.

## Example Usage

```hcl
resource "fmgdevice_system_sittunnel" "trname" {
  auto_asic_offload = "enable"
  destination       = "your own value"
  interface         = ["port2"]
  ip6               = "your own value"
  name              = "your own value"
  device_name       = var.device_name # not required if setting is at provider
  device_vdom       = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `auto_asic_offload` - Enable/disable tunnel ASIC offloading. Valid values: `disable`, `enable`.

* `destination` - Destination IP address of the tunnel.
* `interface` - Interface name.
* `ip6` - IPv6 address of the tunnel.
* `name` - Tunnel name.
* `source` - Source IP address of the tunnel.
* `use_sdwan` - Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System SitTunnel can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_sittunnel.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

