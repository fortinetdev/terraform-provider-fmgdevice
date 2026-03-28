---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_nethsm_hagroups"
description: |-
  Device SystemNethsmHagroups
---

# fmgdevice_system_nethsm_hagroups
Device SystemNethsmHagroups

~> This resource is a sub resource for variable `hagroups` of resource `fmgdevice_system_nethsm`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `member` - Member.
* `name` - Name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System NethsmHagroups can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_nethsm_hagroups.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

