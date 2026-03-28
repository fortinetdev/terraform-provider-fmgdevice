---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_npu_portcpumap"
description: |-
  Configure NPU interface to CPU core mapping.
---

# fmgdevice_system_npu_portcpumap
Configure NPU interface to CPU core mapping.

~> This resource is a sub resource for variable `port_cpu_map` of resource `fmgdevice_system_npu`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `cpu_core` - The CPU core to map to an interface.
* `interface` - The interface to map to a CPU core.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{interface}}.

## Import

System NpuPortCpuMap can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_npu_portcpumap.labelname {{interface}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

