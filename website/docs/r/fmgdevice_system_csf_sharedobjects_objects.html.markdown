---
subcategory: "System"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_csf_sharedobjects_objects"
description: |-
  CMDB table entries.
---

# fmgdevice_system_csf_sharedobjects_objects
CMDB table entries.

~> This resource is a sub resource for variable `objects` of resource `fmgdevice_system_csf_sharedobjects`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `keys`: `fmgdevice_system_csf_sharedobjects_objects_keys`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `shared_objects` - Shared Objects.

* `keys` - Keys. The structure of `keys` block is documented below.
* `pathname` - CMDB path and object name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `keys` block supports:

* `name` - key.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{pathname}}.

## Import

System CsfSharedObjectsObjects can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "shared_objects=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_csf_sharedobjects_objects.labelname {{pathname}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

