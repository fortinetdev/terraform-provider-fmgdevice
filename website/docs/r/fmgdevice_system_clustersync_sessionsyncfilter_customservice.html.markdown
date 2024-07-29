---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_clustersync_sessionsyncfilter_customservice"
description: |-
  Device SystemClusterSyncSessionSyncFilterCustomService
---

# fmgdevice_system_clustersync_sessionsyncfilter_customservice
Device SystemClusterSyncSessionSyncFilterCustomService

~> This resource is a sub resource for variable `custom_service` of resource `fmgdevice_system_clustersync_sessionsyncfilter`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_clustersync_sessionsyncfilter_customservice" "trname" {
  dst_port_range = "your own value"
  fosid          = 10
  src_port_range = "your own value"
  device_name    = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `cluster_sync` - Cluster Sync.

* `dst_port_range` - Dst-Port-Range.
* `fosid` - Id.
* `src_port_range` - Src-Port-Range.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System ClusterSyncSessionSyncFilterCustomService can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "cluster_sync=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_clustersync_sessionsyncfilter_customservice.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

