---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_static6"
description: |-
  Configure IPv6 static routing tables.
---

# fmgdevice_router_static6
Configure IPv6 static routing tables.

## Example Usage

```hcl
resource "fmgdevice_router_static6" "trname" {
  bfd         = "disable"
  blackhole   = "disable"
  comment     = "your own value"
  device      = ["your own value"]
  devindex    = 10
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
* `devindex` - Device index (0 - 4294967295).
* `distance` - Administrative distance (1 - 255).
* `dst` - Destination IPv6 prefix.
* `dstaddr` - Name of firewall address or address group.
* `dynamic_gateway` - Enable use of dynamic gateway retrieved from Router Advertisement (RA). Valid values: `disable`, `enable`.

* `gateway` - IPv6 address of the gateway.
* `link_monitor_exempt` - Enable/disable withdrawal of this static route when link monitor or health check is down. Valid values: `disable`, `enable`.

* `priority` - Administrative priority (1 - 65535).
* `sdwan` - Sdwan. Valid values: `disable`, `enable`.

* `sdwan_zone` - Choose SD-WAN Zone.
* `seq_num` - Sequence number.
* `status` - Enable/disable this static route. Valid values: `disable`, `enable`.

* `virtual_wan_link` - Virtual-Wan-Link. Valid values: `disable`, `enable`.

* `vrf` - Virtual Routing Forwarding ID.
* `weight` - Administrative weight (0 - 255).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{seq_num}}.

## Import

Router Static6 can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_static6.labelname {{seq_num}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

