---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_clustersync"
description: |-
  Device SystemClusterSync
---

# fmgdevice_system_clustersync
Device SystemClusterSync

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `session_sync_filter`: `fmgdevice_system_clustersync_sessionsyncfilter`



## Example Usage

```hcl
resource "fmgdevice_system_clustersync" "trname" {
  down_intfs_before_sess_sync = ["your own value"]
  hb_interval                 = 10
  hb_lost_threshold           = 10
  ike_heartbeat_interval      = 10
  ike_monitor                 = "enable"
  sync_id                     = 10
  device_name                 = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `down_intfs_before_sess_sync` - Down-Intfs-Before-Sess-Sync.
* `hb_interval` - Hb-Interval.
* `hb_lost_threshold` - Hb-Lost-Threshold.
* `ike_heartbeat_interval` - Ike-Heartbeat-Interval.
* `ike_monitor` - Ike-Monitor. Valid values: `disable`, `enable`.

* `ike_monitor_interval` - Ike-Monitor-Interval.
* `ike_seqjump_speed` - Ike-Seqjump-Speed.
* `ike_use_rfc6311` - Ike-Use-Rfc6311. Valid values: `disable`, `enable`.

* `ipsec_tunnel_sync` - Ipsec-Tunnel-Sync. Valid values: `disable`, `enable`.

* `peerip` - Peerip.
* `peervd` - Peervd.
* `secondary_add_ipsec_routes` - Secondary-Add-Ipsec-Routes. Valid values: `disable`, `enable`.

* `session_sync_filter` - Session-Sync-Filter. The structure of `session_sync_filter` block is documented below.
* `slave_add_ike_routes` - Slave-Add-Ike-Routes. Valid values: `disable`, `enable`.

* `sync_id` - Sync-Id.
* `syncvd` - Syncvd.

The `session_sync_filter` block supports:

* `custom_service` - Custom-Service. The structure of `custom_service` block is documented below.
* `dstaddr` - Dstaddr.
* `dstaddr6` - Dstaddr6.
* `dstintf` - Dstintf.
* `srcaddr` - Srcaddr.
* `srcaddr6` - Srcaddr6.
* `srcintf` - Srcintf.

The `custom_service` block supports:

* `dst_port_range` - Dst-Port-Range.
* `id` - Id.
* `src_port_range` - Src-Port-Range.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{sync_id}}.

## Import

System ClusterSync can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_clustersync.labelname {{sync_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

