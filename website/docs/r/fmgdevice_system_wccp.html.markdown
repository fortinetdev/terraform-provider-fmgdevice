---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_wccp"
description: |-
  Configure WCCP.
---

# fmgdevice_system_wccp
Configure WCCP.

## Example Usage

```hcl
resource "fmgdevice_system_wccp" "trname" {
  assignment_bucket_format = "wccp-v2"
  assignment_dstaddr_mask  = "your own value"
  assignment_method        = "any"
  assignment_srcaddr_mask  = "your own value"
  assignment_weight        = 10
  service_id               = "your own value"
  device_name              = var.device_name # not required if setting is at provider
  device_vdom              = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `assignment_bucket_format` - Assignment bucket format for the WCCP cache engine. Valid values: `cisco-implementation`, `wccp-v2`.

* `assignment_dstaddr_mask` - Assignment destination address mask.
* `assignment_method` - Hash key assignment preference. Valid values: `HASH`, `MASK`, `any`.

* `assignment_srcaddr_mask` - Assignment source address mask.
* `assignment_weight` - Assignment of hash weight/ratio for the WCCP cache engine.
* `authentication` - Enable/disable MD5 authentication. Valid values: `disable`, `enable`.

* `cache_engine_method` - Method used to forward traffic to the routers or to return to the cache engine. Valid values: `GRE`, `L2`.

* `cache_id` - IP address known to all routers. If the addresses are the same, use the default 0.0.0.0.
* `forward_method` - Method used to forward traffic to the cache servers. Valid values: `GRE`, `L2`, `any`.

* `group_address` - IP multicast address used by the cache routers. For the FortiGate to ignore multicast WCCP traffic, use the default 0.0.0.0.
* `password` - Password for MD5 authentication.
* `ports` - Service ports.
* `ports_defined` - Match method. Valid values: `source`, `destination`.

* `primary_hash` - Hash method. Valid values: `src-ip`, `dst-ip`, `src-port`, `dst-port`, `ports-defined`, `ports-source`.

* `priority` - Service priority.
* `protocol` - Service protocol.
* `return_method` - Method used to decline a redirected packet and return it to the FortiGate unit. Valid values: `GRE`, `L2`, `any`.

* `router_id` - IP address known to all cache engines. If all cache engines connect to the same FortiGate interface, use the default 0.0.0.0.
* `router_list` - IP addresses of one or more WCCP routers.
* `server_list` - IP addresses and netmasks for up to four cache servers.
* `server_type` - Cache server type. Valid values: `forward`, `proxy`.

* `service_id` - Service ID.
* `service_type` - WCCP service type used by the cache server for logical interception and redirection of traffic. Valid values: `standard`, `dynamic`, `auto`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{service_id}}.

## Import

System Wccp can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_wccp.labelname {{service_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

