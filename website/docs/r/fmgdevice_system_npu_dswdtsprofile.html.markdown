---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_npu_dswdtsprofile"
description: |-
  Configure NPU DSW DTS profile.
---

# fmgdevice_system_npu_dswdtsprofile
Configure NPU DSW DTS profile.

~> This resource is a sub resource for variable `dsw_dts_profile` of resource `fmgdevice_system_npu`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `action` - Set NPU DSW DTS profile action. Valid values: `wait`, `drop`, `drop_tmr_0`, `drop_tmr_1`, `enque`, `enque_0`, `enque_1`.

* `min_limit` - Set NPU DSW DTS profile min-limt.
* `profile_id` - Set NPU DSW DTS profile profile id.
* `step` - Set NPU DSW DTS profile step.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{profile_id}}.

## Import

System NpuDswDtsProfile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_npu_dswdtsprofile.labelname {{profile_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

