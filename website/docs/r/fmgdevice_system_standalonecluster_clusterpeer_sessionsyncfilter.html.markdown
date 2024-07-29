---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_standalonecluster_clusterpeer_sessionsyncfilter"
description: |-
  Add one or more filters if you only want to synchronize some sessions. Use the filter to configure the types of sessions to synchronize.
---

# fmgdevice_system_standalonecluster_clusterpeer_sessionsyncfilter
Add one or more filters if you only want to synchronize some sessions. Use the filter to configure the types of sessions to synchronize.

~> This resource is a sub resource for variable `session_sync_filter` of resource `fmgdevice_system_standalonecluster_clusterpeer`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `custom_service`: `fmgdevice_system_standalonecluster_clusterpeer_sessionsyncfilter_customservice`



## Example Usage

```hcl
resource "fmgdevice_system_standalonecluster_clusterpeer_sessionsyncfilter" "trname" {
  custom_service {
    dst_port_range = "your own value"
    fosid          = 10
    src_port_range = "your own value"
  }

  dstaddr     = ["your own value"]
  dstaddr6    = "your own value"
  dstintf     = ["your own value"]
  srcaddr     = ["your own value"]
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `cluster_peer` - Cluster Peer.

* `custom_service` - Custom-Service. The structure of `custom_service` block is documented below.
* `dstaddr` - Only sessions to this IPv4 address are synchronized.
* `dstaddr6` - Only sessions to this IPv6 address are synchronized.
* `dstintf` - Only sessions to this interface are synchronized.
* `srcaddr` - Only sessions from this IPv4 address are synchronized.
* `srcaddr6` - Only sessions from this IPv6 address are synchronized.
* `srcintf` - Only sessions from this interface are synchronized.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `custom_service` block supports:

* `dst_port_range` - Custom service destination port range.
* `id` - Custom service ID.
* `src_port_range` - Custom service source port range.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System StandaloneClusterClusterPeerSessionSyncFilter can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "cluster_peer=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_standalonecluster_clusterpeer_sessionsyncfilter.labelname SystemStandaloneClusterClusterPeerSessionSyncFilter
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

