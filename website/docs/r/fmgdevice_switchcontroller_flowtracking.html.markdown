---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_flowtracking"
description: |-
  Configure FortiSwitch flow tracking and export via ipfix/netflow.
---

# fmgdevice_switchcontroller_flowtracking
Configure FortiSwitch flow tracking and export via ipfix/netflow.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `aggregates`: `fmgdevice_switchcontroller_flowtracking_aggregates`
>- `collectors`: `fmgdevice_switchcontroller_flowtracking_collectors`



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_flowtracking" "trname" {
  aggregates {
    fosid = 10
    ip    = ["your own value"]
  }

  collector_ip   = "your own value"
  collector_port = 10
  collectors {
    ip        = "your own value"
    name      = "your own value"
    port      = 10
    transport = "sctp"
  }

  format      = "netflow5"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `aggregates` - Aggregates. The structure of `aggregates` block is documented below.
* `collector_ip` - Configure collector ip address.
* `collector_port` - Configure collector port number(0-65535, default=0).
* `collectors` - Collectors. The structure of `collectors` block is documented below.
* `format` - Configure flow tracking protocol. Valid values: `netflow1`, `netflow5`, `netflow9`, `ipfix`.

* `level` - Configure flow tracking level. Valid values: `vlan`, `ip`, `port`, `proto`, `mac`.

* `max_export_pkt_size` - Configure flow max export packet size (512-9216, default=512 bytes).
* `sample_mode` - Configure sample mode for the flow tracking. Valid values: `local`, `perimeter`, `device-ingress`.

* `sample_rate` - Configure sample rate for the perimeter and device-ingress sampling(0 - 99999).
* `template_export_period` - Configure template export period (1-60, default=5 minutes).
* `timeout_general` - Configure flow session general timeout (60-604800, default=3600 seconds).
* `timeout_icmp` - Configure flow session ICMP timeout (60-604800, default=300 seconds).
* `timeout_max` - Configure flow session max timeout (60-604800, default=604800 seconds).
* `timeout_tcp` - Configure flow session TCP timeout (60-604800, default=3600 seconds).
* `timeout_tcp_fin` - Configure flow session TCP FIN timeout (60-604800, default=300 seconds).
* `timeout_tcp_rst` - Configure flow session TCP RST timeout (60-604800, default=120 seconds).
* `timeout_udp` - Configure flow session UDP timeout (60-604800, default=300 seconds).
* `transport` - Configure L4 transport protocol for exporting packets. Valid values: `udp`, `tcp`, `sctp`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `aggregates` block supports:

* `id` - Aggregate id.
* `ip` - IP address to group all matching traffic sessions to a flow.

The `collectors` block supports:

* `ip` - Collector IP address.
* `name` - Collector name.
* `port` - Collector port number(0-65535, default:0, netflow:2055, ipfix:4739).
* `transport` - Collector L4 transport protocol for exporting packets. Valid values: `udp`, `tcp`, `sctp`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController FlowTracking can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_flowtracking.labelname SwitchControllerFlowTracking
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

