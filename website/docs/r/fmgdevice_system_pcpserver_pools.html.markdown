---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_pcpserver_pools"
description: |-
  Configure PCP pools.
---

# fmgdevice_system_pcpserver_pools
Configure PCP pools.

~> This resource is a sub resource for variable `pools` of resource `fmgdevice_system_pcpserver`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_pcpserver_pools" "trname" {
  allow_opcode         = ["map"]
  announcement_count   = 10
  arp_reply            = "disable"
  client_mapping_limit = 10
  client_subnet        = ["your own value"]
  name                 = "your own value"
  device_name          = var.device_name # not required if setting is at provider
  device_vdom          = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `allow_opcode` - Allowed PCP opcode. Valid values: `map`, `peer`, `announce`.

* `announcement_count` - Number of multicast announcements.
* `arp_reply` - Enable to respond to ARP requests for external IP (default = enable). Valid values: `disable`, `enable`.

* `client_mapping_limit` - Mapping limit per client (0 - 65535, default = 0, 0 = unlimited).
* `client_subnet` - Subnets from which PCP requests are accepted.
* `description` - Description.
* `ext_intf` - External interface name.
* `extip` - IP address or address range on the external interface that you want to map to an address on the internal network.
* `extport` - Incoming port number range that you want to map to a port number on the internal network.
* `fosid` - ID.
* `intl_intf` - Internal interface name.
* `mapping_filter_limit` - Filter limit per mapping (0 - 5, default = 1).
* `maximal_lifetime` - Maximal lifetime of a PCP mapping in seconds (3600 - 604800, default = 86400).
* `minimal_lifetime` - Minimal lifetime of a PCP mapping in seconds (60 - 300, default = 120).
* `multicast_announcement` - Enable/disable multicast announcements. Valid values: `disable`, `enable`.

* `name` - PCP pool name.
* `recycle_delay` - Minimum delay (in seconds) the PCP Server will wait before recycling mappings that have expired (0 - 3600, default = 0).
* `third_party` - Allow/disallow third party option. Valid values: `disallow`, `allow`.

* `third_party_subnet` - Subnets from which third party requests are accepted.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System PcpServerPools can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_pcpserver_pools.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

