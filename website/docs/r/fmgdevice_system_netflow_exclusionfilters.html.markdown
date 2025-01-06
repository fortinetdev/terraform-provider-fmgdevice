---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_netflow_exclusionfilters"
description: |-
  Exclusion filters
---

# fmgdevice_system_netflow_exclusionfilters
Exclusion filters

~> This resource is a sub resource for variable `exclusion_filters` of resource `fmgdevice_system_netflow`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `destination_ip` - Session destination address.
* `destination_port` - Session destination port number or range.
* `fosid` - Filter ID.
* `protocol` - Session IP protocol (0 - 255, default = 255, meaning any).
* `source_ip` - Session source address.
* `source_port` - Session source port number or range.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System NetflowExclusionFilters can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_netflow_exclusionfilters.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

