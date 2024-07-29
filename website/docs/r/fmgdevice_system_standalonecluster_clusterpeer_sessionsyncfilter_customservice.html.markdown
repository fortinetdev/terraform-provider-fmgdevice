---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_standalonecluster_clusterpeer_sessionsyncfilter_customservice"
description: |-
  Only sessions using these custom services are synchronized. Use source and destination port ranges to define these custom services.
---

# fmgdevice_system_standalonecluster_clusterpeer_sessionsyncfilter_customservice
Only sessions using these custom services are synchronized. Use source and destination port ranges to define these custom services.

~> This resource is a sub resource for variable `custom_service` of resource `fmgdevice_system_standalonecluster_clusterpeer_sessionsyncfilter`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_standalonecluster_clusterpeer_sessionsyncfilter_customservice" "trname" {
  dst_port_range = "your own value"
  fosid          = 10
  src_port_range = "your own value"
  device_name    = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `cluster_peer` - Cluster Peer.

* `dst_port_range` - Custom service destination port range.
* `fosid` - Custom service ID.
* `src_port_range` - Custom service source port range.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System StandaloneClusterClusterPeerSessionSyncFilterCustomService can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "cluster_peer=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_standalonecluster_clusterpeer_sessionsyncfilter_customservice.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

