---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_sflow"
description: |-
  Configure sFlow.
---

# fmgdevice_system_sflow
Configure sFlow.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `collectors`: `fmgdevice_system_sflow_collectors`



## Example Usage

```hcl
resource "fmgdevice_system_sflow" "trname" {
  collector_ip            = "your own value"
  collector_port          = 10
  interface               = ["port2"]
  interface_select_method = "specify"
  source_ip               = "your own value"
  device_name             = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `collector_ip` - IP address of the sFlow collector that sFlow agents added to interfaces in this VDOM send sFlow datagrams to (default = 0.0.0.0).
* `collector_port` - UDP port number used for sending sFlow datagrams (configure only if required by your sFlow collector or your network configuration) (0 - 65535, default = 6343).
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `source_ip` - Source IP address for sFlow agent.
* `collectors` - Collectors. The structure of `collectors` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `collectors` block supports:

* `collector_ip` - IP addresses of the sFlow collectors that sFlow agents added to interfaces in this VDOM send sFlow datagrams to.
* `collector_port` - UDP port number used for sending sFlow datagrams (configure only if required by your sFlow collector or your network configuration) (0 - 65535, default = 6343).
* `id` - ID.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `source_ip` - Source IP address for sFlow agent.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Sflow can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_sflow.labelname SystemSflow
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

