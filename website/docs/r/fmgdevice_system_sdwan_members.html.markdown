---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_sdwan_members"
description: |-
  FortiGate interfaces added to the SD-WAN.
---

# fmgdevice_system_sdwan_members
FortiGate interfaces added to the SD-WAN.

~> This resource is a sub resource for variable `members` of resource `fmgdevice_system_sdwan`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_sdwan_members" "trname" {
  comment                     = "your own value"
  cost                        = 10
  gateway                     = "your own value"
  gateway6                    = "your own value"
  ingress_spillover_threshold = 10
  seq_num                     = 10
  device_name                 = var.device_name # not required if setting is at provider
  device_vdom                 = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comment` - Comments.
* `cost` - Cost of this interface for services in SLA mode (0 - 4294967295, default = 0).
* `gateway` - The default gateway for this interface. Usually the default gateway of the Internet service provider that this interface is connected to.
* `gateway6` - IPv6 gateway.
* `ingress_spillover_threshold` - Ingress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.
* `interface` - Interface name.
* `preferred_source` - Preferred source of route for this member.
* `priority` - Priority of the interface for IPv4 (1 - 65535, default = 1). Used for SD-WAN rules or priority rules.
* `priority6` - Priority of the interface for IPv6 (1 - 65535, default = 1024). Used for SD-WAN rules or priority rules.
* `seq_num` - Sequence number(1-512).
* `source` - Source IP address used in the health-check packet to the server.
* `source6` - Source IPv6 address used in the health-check packet to the server.
* `spillover_threshold` - Egress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.
* `status` - Enable/disable this interface in the SD-WAN. Valid values: `disable`, `enable`.

* `transport_group` - Measured transport group (0 - 255).
* `volume_ratio` - Measured volume ratio (this value / sum of all values = percentage of link volume, 1 - 255).
* `weight` - Weight of this interface for weighted load balancing. (1 - 255) More traffic is directed to interfaces with higher weights.
* `zone` - Zone name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{seq_num}}.

## Import

System SdwanMembers can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_sdwan_members.labelname {{seq_num}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

