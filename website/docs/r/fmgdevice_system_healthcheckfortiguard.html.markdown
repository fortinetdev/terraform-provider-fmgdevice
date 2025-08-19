---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_healthcheckfortiguard"
description: |-
  SD-WAN status checking or health checking. Identify a server predefine by FortiGuard and determine how SD-WAN verifies that FGT can communicate with it.
---

# fmgdevice_system_healthcheckfortiguard
SD-WAN status checking or health checking. Identify a server predefine by FortiGuard and determine how SD-WAN verifies that FGT can communicate with it.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `name` - Status check or predefined health-check targets name.
* `obsolete` - Obsolete.
* `protocol` - Protocol name. Valid values: `ping`, `tcp-echo`, `udp-echo`, `http`, `twamp`, `dns`, `tcp-connect`, `ftp`, `https`.

* `server` - Status check or predefined health-check domain name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System HealthCheckFortiguard can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_healthcheckfortiguard.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

