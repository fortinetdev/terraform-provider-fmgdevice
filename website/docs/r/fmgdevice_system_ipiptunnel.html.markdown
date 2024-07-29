---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_ipiptunnel"
description: |-
  Configure IP in IP Tunneling.
---

# fmgdevice_system_ipiptunnel
Configure IP in IP Tunneling.

## Example Usage

```hcl
resource "fmgdevice_system_ipiptunnel" "trname" {
  auto_asic_offload = "enable"
  interface         = ["port2"]
  local_gw          = "your own value"
  name              = "your own value"
  remote_gw         = "your own value"
  device_name       = var.device_name # not required if setting is at provider
  device_vdom       = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `auto_asic_offload` - Enable/disable tunnel ASIC offloading. Valid values: `disable`, `enable`.

* `interface` - Interface name that is associated with the incoming traffic from available options.
* `local_gw` - IPv4 address for the local gateway.
* `name` - IPIP Tunnel name.
* `remote_gw` - IPv4 address for the remote gateway.
* `use_sdwan` - Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System IpipTunnel can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_ipiptunnel.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

