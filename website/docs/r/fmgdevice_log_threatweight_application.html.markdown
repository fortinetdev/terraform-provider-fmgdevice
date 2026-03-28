---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_log_threatweight_application"
description: |-
  Application-control threat weight settings.
---

# fmgdevice_log_threatweight_application
Application-control threat weight settings.

~> This resource is a sub resource for variable `application` of resource `fmgdevice_log_threatweight`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `category` - Application category.
* `fosid` - Entry ID.
* `level` - Threat weight score for Application events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Log ThreatWeightApplication can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_log_threatweight_application.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

