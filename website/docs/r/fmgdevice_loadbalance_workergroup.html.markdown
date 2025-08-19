---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_loadbalance_workergroup"
description: |-
  Worker group configuration.
---

# fmgdevice_loadbalance_workergroup
Worker group configuration.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `member` - FPM in the group.
* `worker_group_name` - Name of worker group.  Default worker group can not be changed.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{worker_group_name}}.

## Import

LoadBalance WorkerGroup can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_loadbalance_workergroup.labelname {{worker_group_name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

