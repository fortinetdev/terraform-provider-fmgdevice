---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_static"
description: |-
  Configure IPv4 static routing tables.
---

# fmgdevice_router_static
Configure IPv4 static routing tables.

## Example Usage

```hcl
resource "fmgdevice_router_static" "trname" {
  bfd         = "disable"
  blackhole   = "disable"
  comment     = "your own value"
  device      = ["your own value"]
  distance    = 10
  seq_num     = 10
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `bfd` - Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `disable`, `enable`.

* `blackhole` - Enable/disable black hole. Valid values: `disable`, `enable`.

* `comment` - Optional comments.
* `device` - Gateway out interface or tunnel.
* `distance` - Administrative distance (1 - 255).
* `dst` - Destination IP and mask for this route.
* `dst_type` - Dst-Type. Valid values: `ipmask`, `addrname`, `service-id`, `service-custom`.

* `dstaddr` - Name of firewall address or address group.
* `dynamic_gateway` - Enable use of dynamic gateway retrieved from a DHCP or PPP server. Valid values: `disable`, `enable`.

* `gateway` - Gateway IP for this route.
* `internet_service` - Application ID in the Internet service database.
* `internet_service_custom` - Application name in the Internet service custom database.
* `link_monitor_exempt` - Enable/disable withdrawal of this static route when link monitor or health check is down. Valid values: `disable`, `enable`.

* `preferred_source` - Preferred source IP for this route.
* `priority` - Administrative priority (1 - 65535).
* `sdwan` - Sdwan. Valid values: `disable`, `enable`.

* `sdwan_zone` - Choose SD-WAN Zone.
* `seq_num` - Sequence number.
* `src` - Source prefix for this route.
* `status` - Enable/disable this static route. Valid values: `disable`, `enable`.

* `virtual_wan_link` - Virtual-Wan-Link. Valid values: `disable`, `enable`.

* `tag` - Route tag.
* `vrf` - Virtual Routing Forwarding ID.
* `weight` - Administrative weight (0 - 255).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{seq_num}}.

## Import

Router Static can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_static.labelname {{seq_num}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

