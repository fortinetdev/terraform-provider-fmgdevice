---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_standalonecluster_monitorprefix"
description: |-
  Configure a list of routing prefixes to monitor.
---

# fmgdevice_system_standalonecluster_monitorprefix
Configure a list of routing prefixes to monitor.

~> This resource is a sub resource for variable `monitor_prefix` of resource `fmgdevice_system_standalonecluster`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `fosid` - ID.
* `prefix` - Prefix.
* `vdom` - VDOM name.
* `vrf` - VRF ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System StandaloneClusterMonitorPrefix can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_standalonecluster_monitorprefix.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

