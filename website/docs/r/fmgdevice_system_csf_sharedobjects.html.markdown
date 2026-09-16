---
subcategory: "System"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_csf_sharedobjects"
description: |-
  Fabric-wide objects shared by non-root nodes.
---

# fmgdevice_system_csf_sharedobjects
Fabric-wide objects shared by non-root nodes.

~> This resource is a sub resource for variable `shared_objects` of resource `fmgdevice_system_csf`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `objects`: `fmgdevice_system_csf_sharedobjects_objects`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `name` - UID of the source device.
* `objects` - Objects. The structure of `objects` block is documented below.
* `trusted_list_entry` - Trusted list entry name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `objects` block supports:

* `keys` - Keys. The structure of `keys` block is documented below.
* `pathname` - CMDB path and object name.

The `keys` block supports:

* `name` - key.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System CsfSharedObjects can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_csf_sharedobjects.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

