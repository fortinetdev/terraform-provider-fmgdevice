---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_wagprofile"
description: |-
  Configure wireless access gateway (WAG) profiles used for tunnels on AP.
---

# fmgdevice_wirelesscontroller_wagprofile
Configure wireless access gateway (WAG) profiles used for tunnels on AP.

## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_wagprofile" "trname" {
  comment       = "your own value"
  dhcp_ip_addr  = "your own value"
  name          = "your own value"
  ping_interval = 10
  ping_number   = 10
  device_name   = var.device_name # not required if setting is at provider
  device_vdom   = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comment` - Comment.
* `dhcp_ip_addr` - IP address of the monitoring DHCP request packet sent through the tunnel.
* `name` - Tunnel profile name.
* `ping_interval` - Interval between two tunnel monitoring echo packets (1 - 65535 sec, default = 1).
* `ping_number` - Number of the tunnel monitoring echo packets (1 - 65535, default = 5).
* `return_packet_timeout` - Window of time for the return packets from the tunnel's remote end (1 - 65535 sec, default = 160).
* `tunnel_type` - Tunnel type. Valid values: `gre`, `l2tpv3`.

* `wag_ip` - IP Address of the wireless access gateway.
* `wag_port` - UDP port of the wireless access gateway.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController WagProfile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_wagprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

