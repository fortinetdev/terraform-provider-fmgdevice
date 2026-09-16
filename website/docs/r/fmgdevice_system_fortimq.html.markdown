---
subcategory: "System"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_fortimq"
description: |-
  Configure FortiMQ settings.
---

# fmgdevice_system_fortimq
Configure FortiMQ settings.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `ocsp_check` - Enable/disable OCSP check when connecting to FortiMQ cloud service (default = enable). Valid values: `disable`, `enable`.

* `publish_metadata` - Enable/disable the publishing of FortiGate metadata to the FortiMQ cloud sevice (default = enable). Valid values: `disable`, `enable`.

* `status` - Enable/disable FortiMQ features (default = enable). Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Fortimq can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_fortimq.labelname SystemFortimq
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

