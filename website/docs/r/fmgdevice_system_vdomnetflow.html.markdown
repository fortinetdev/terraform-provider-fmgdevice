---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_vdomnetflow"
description: |-
  Configure NetFlow per VDOM.
---

# fmgdevice_system_vdomnetflow
Configure NetFlow per VDOM.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `collectors`: `fmgdevice_system_vdomnetflow_collectors`



## Example Usage

```hcl
resource "fmgdevice_system_vdomnetflow" "trname" {
  collector_ip   = "your own value"
  collector_port = 10
  collectors {
    collector_ip            = "your own value"
    collector_port          = 10
    fosid                   = 10
    interface               = ["port2"]
    interface_select_method = "auto"
    source_ip               = "your own value"
  }

  interface               = ["port2"]
  interface_select_method = "auto"
  device_name             = var.device_name # not required if setting is at provider
  device_vdom             = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `collector_ip` - Collector-Ip.
* `collector_port` - Collector-Port.
* `collectors` - Collectors. The structure of `collectors` block is documented below.
* `interface` - Interface.
* `interface_select_method` - Interface-Select-Method. Valid values: `auto`, `sdwan`, `specify`.

* `source_ip` - Source-Ip.
* `vdom_netflow` - Enable/disable NetFlow per VDOM. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `collectors` block supports:

* `collector_ip` - Collector IP.
* `collector_port` - NetFlow collector port number.
* `id` - ID.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `source_ip` - Source IP address for communication with the NetFlow agent.
* `source_ip_interface` - Name of the interface used to determine the source IP for exporting packets.
* `vrf_select` - VRF ID used for connection to server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System VdomNetflow can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_vdomnetflow.labelname SystemVdomNetflow
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

