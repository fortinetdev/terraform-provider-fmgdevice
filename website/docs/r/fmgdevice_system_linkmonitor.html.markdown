---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_linkmonitor"
description: |-
  Configure Link Health Monitor.
---

# fmgdevice_system_linkmonitor
Configure Link Health Monitor.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `server_list`: `fmgdevice_system_linkmonitor_serverlist`



## Example Usage

```hcl
resource "fmgdevice_system_linkmonitor" "trname" {
  addr_mode    = "ipv6"
  class_id     = ["your own value"]
  diffservcode = "your own value"
  fail_weight  = 10
  failtime     = 10
  name         = "your own value"
  device_name  = var.device_name # not required if setting is at provider
  device_vdom  = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `addr_mode` - Address mode (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.

* `class_id` - Traffic class ID.
* `diffservcode` - Differentiated services code point (DSCP) in the IP header of the probe packet.
* `fail_weight` - Threshold weight to trigger link failure alert.
* `failtime` - Number of retry attempts before the server is considered down (1 - 3600, default = 5).
* `gateway_ip` - Gateway IP address used to probe the server.
* `gateway_ip6` - Gateway IPv6 address used to probe the server.
* `http_agent` - String in the http-agent field in the HTTP header.
* `http_get` - If you are monitoring an HTML server you can send an HTTP-GET request with a custom string. Use this option to define the string.
* `http_match` - String that you expect to see in the HTTP-GET requests of the traffic to be monitored.
* `interval` - Detection interval in milliseconds (20 - 3600 * 1000 msec, default = 500).
* `name` - Link monitor name.
* `packet_size` - Packet size of a TWAMP test session (124/158 - 1024).
* `password` - TWAMP controller password in authentication mode.
* `port` - Port number of the traffic to be used to monitor the server.
* `probe_count` - Number of most recent probes that should be used to calculate latency and jitter (5 - 30, default = 30).
* `probe_timeout` - Time to wait before a probe packet is considered lost (20 - 5000 msec, default = 500).
* `protocol` - Protocols used to monitor the server. Valid values: `ping`, `tcp-echo`, `udp-echo`, `http`, `twamp`, `ping6`, `https`.

* `recoverytime` - Number of successful responses received before server is considered recovered (1 - 3600, default = 5).
* `route` - Subnet to monitor.
* `security_mode` - Twamp controller security mode. Valid values: `none`, `authentication`.

* `server` - IP address of the server(s) to be monitored.
* `server_config` - Mode of server configuration. Valid values: `default`, `individual`.

* `server_list` - Server-List. The structure of `server_list` block is documented below.
* `server_type` - Server type (static or dynamic). Valid values: `static`, `dynamic`.

* `service_detection` - Only use monitor to read quality values. If enabled, static routes and cascade interfaces will not be updated. Valid values: `disable`, `enable`.

* `source_ip` - Source IP address used in packet to the server.
* `source_ip6` - Source IPv6 address used in packet to the server.
* `srcintf` - Interface that receives the traffic to be monitored.
* `status` - Enable/disable this link monitor. Valid values: `disable`, `enable`.

* `update_cascade_interface` - Enable/disable update cascade interface. Valid values: `disable`, `enable`.

* `update_policy_route` - Enable/disable updating the policy route. Valid values: `disable`, `enable`.

* `update_static_route` - Enable/disable updating the static route. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `server_list` block supports:

* `dst` - IP address of the server to be monitored.
* `id` - Server ID.
* `port` - Port number of the traffic to be used to monitor the server.
* `protocol` - Protocols used to monitor the server. Valid values: `ping`, `tcp-echo`, `udp-echo`, `http`, `twamp`, `https`.

* `weight` - Weight of the monitor to this dst (0 - 255).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System LinkMonitor can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_linkmonitor.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

