---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_log_threatweight_level"
description: |-
  Score mapping for threat weight levels.
---

# fmgdevice_log_threatweight_level
Score mapping for threat weight levels.

~> This resource is a sub resource for variable `level` of resource `fmgdevice_log_threatweight`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `critical` - Critical level score value (1 - 100).
* `high` - High level score value (1 - 100).
* `low` - Low level score value (1 - 100).
* `medium` - Medium level score value (1 - 100).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Log ThreatWeightLevel can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_log_threatweight_level.labelname LogThreatWeightLevel
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

