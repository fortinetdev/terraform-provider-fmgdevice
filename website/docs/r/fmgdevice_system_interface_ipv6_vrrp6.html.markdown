---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_interface_ipv6_vrrp6"
description: |-
  IPv6 VRRP configuration.
---

# fmgdevice_system_interface_ipv6_vrrp6
IPv6 VRRP configuration.

~> This resource is a sub resource for variable `vrrp6` of resource `fmgdevice_system_interface_ipv6`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_interface_ipv6_vrrp6" "trname" {
  accept_mode          = "disable"
  adv_interval         = 10
  ignore_default_route = "enable"
  preempt              = "disable"
  priority             = 10
  vrid                 = 10
  device_name          = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `interface` - Interface.

* `accept_mode` - Enable/disable accept mode. Valid values: `disable`, `enable`.

* `adv_interval` - Advertisement interval (1 - 255 seconds).
* `ignore_default_route` - Enable/disable ignoring of default route when checking destination. Valid values: `disable`, `enable`.

* `preempt` - Enable/disable preempt mode. Valid values: `disable`, `enable`.

* `priority` - Priority of the virtual router (1 - 255).
* `start_time` - Startup time (1 - 255 seconds).
* `status` - Enable/disable VRRP. Valid values: `disable`, `enable`.

* `vrdst_priority` - Priority of the virtual router when the virtual router destination becomes unreachable (0 - 254).
* `vrdst6` - Monitor the route to this destination.
* `vrgrp` - VRRP group ID (1 - 65535).
* `vrid` - Virtual router identifier (1 - 255).
* `vrip6` - IPv6 address of the virtual router.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{vrid}}.

## Import

System InterfaceIpv6Vrrp6 can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "interface=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_interface_ipv6_vrrp6.labelname {{vrid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

