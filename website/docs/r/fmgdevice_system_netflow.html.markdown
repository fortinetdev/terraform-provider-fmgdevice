---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_netflow"
description: |-
  Configure NetFlow.
---

# fmgdevice_system_netflow
Configure NetFlow.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `collectors`: `fmgdevice_system_netflow_collectors`



## Example Usage

```hcl
resource "fmgdevice_system_netflow" "trname" {
  active_flow_timeout = 10
  collector_ip        = "your own value"
  collector_port      = 10
  collectors {
    collector_ip            = "your own value"
    collector_port          = 10
    fosid                   = 10
    interface               = ["port2"]
    interface_select_method = "auto"
    source_ip               = "your own value"
  }

  inactive_flow_timeout = 10
  device_name           = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `active_flow_timeout` - Timeout to report active flows (60 - 3600 sec, default = 1800).
* `collector_ip` - Collector-Ip.
* `collector_port` - Collector-Port.
* `collectors` - Collectors. The structure of `collectors` block is documented below.
* `inactive_flow_timeout` - Timeout for periodic report of finished flows (10 - 600 sec, default = 15).
* `interface` - Interface.
* `interface_select_method` - Interface-Select-Method. Valid values: `auto`, `sdwan`, `specify`.

* `source_ip` - Source-Ip.
* `template_tx_counter` - Counter of flowset records before resending a template flowset record.
* `template_tx_timeout` - Timeout for periodic template flowset transmission (60 - 86400 sec, default = 1800).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `collectors` block supports:

* `collector_ip` - Collector IP.
* `collector_port` - NetFlow collector port number.
* `id` - ID.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `source_ip` - Source IP address for communication with the NetFlow agent.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Netflow can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_netflow.labelname SystemNetflow
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

