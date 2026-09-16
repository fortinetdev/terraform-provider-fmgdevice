---
subcategory: "System"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_csf_fabricdatasourceexemption"
description: |-
  Disable the fabric datasource check on the tables when synchronizing them.
---

# fmgdevice_system_csf_fabricdatasourceexemption
Disable the fabric datasource check on the tables when synchronizing them.

~> This resource is a sub resource for variable `fabric_datasource_exemption` of resource `fmgdevice_system_csf`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `name` - Name.
* `status` - Enable/disable the fabric datasource check on the target table. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System CsfFabricDatasourceExemption can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_csf_fabricdatasourceexemption.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

