---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_npu_portnpumap"
description: |-
  Configure port to NPU group mapping.
---

# fmgdevice_system_npu_portnpumap
Configure port to NPU group mapping.

~> This resource is a sub resource for variable `port_npu_map` of resource `fmgdevice_system_npu`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `interface` - Set NPU interface port for NPU group mapping.
* `npu_group_index` - Mapping NPU group index.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{interface}}.

## Import

System NpuPortNpuMap can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_npu_portnpumap.labelname {{interface}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

