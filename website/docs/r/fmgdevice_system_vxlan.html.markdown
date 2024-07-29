---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_vxlan"
description: |-
  Configure VXLAN devices.
---

# fmgdevice_system_vxlan
Configure VXLAN devices.

## Example Usage

```hcl
resource "fmgdevice_system_vxlan" "trname" {
  dstport            = 10
  evpn_id            = ["your own value"]
  interface          = ["port2"]
  ip_version         = "ipv4-multicast"
  learn_from_traffic = "enable"
  name               = "your own value"
  device_name        = var.device_name # not required if setting is at provider
  device_vdom        = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `dstport` - VXLAN destination port (1 - 65535, default = 4789).
* `evpn_id` - EVPN instance.
* `interface` - Outgoing interface for VXLAN encapsulated traffic.
* `ip_version` - IP version to use for the VXLAN interface and so for communication over the VXLAN. IPv4 or IPv6 unicast or multicast. Valid values: `ipv4-unicast`, `ipv6-unicast`, `ipv4-multicast`, `ipv6-multicast`.

* `learn_from_traffic` - Enable/disable VXLAN MAC learning from traffic. Valid values: `disable`, `enable`.

* `multicast_ttl` - VXLAN multicast TTL (1-255, default = 0).
* `name` - VXLAN device or interface name. Must be a unique interface name.
* `remote_ip` - IPv4 address of the VXLAN interface on the device at the remote end of the VXLAN.
* `remote_ip6` - IPv6 IP address of the VXLAN interface on the device at the remote end of the VXLAN.
* `vni` - VXLAN network ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Vxlan can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_vxlan.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

