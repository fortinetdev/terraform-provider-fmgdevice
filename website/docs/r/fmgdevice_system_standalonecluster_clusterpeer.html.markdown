---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_standalonecluster_clusterpeer"
description: |-
  Configure FortiGate Session Life Support Protocol (FGSP) session synchronization.
---

# fmgdevice_system_standalonecluster_clusterpeer
Configure FortiGate Session Life Support Protocol (FGSP) session synchronization.

~> This resource is a sub resource for variable `cluster_peer` of resource `fmgdevice_system_standalonecluster`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `session_sync_filter`: `fmgdevice_system_standalonecluster_clusterpeer_sessionsyncfilter`



## Example Usage

```hcl
resource "fmgdevice_system_standalonecluster_clusterpeer" "trname" {
  down_intfs_before_sess_sync = ["your own value"]
  hb_interval                 = 10
  hb_lost_threshold           = 10
  ike_heartbeat_interval      = 10
  ike_monitor                 = "disable"
  sync_id                     = 10
  device_name                 = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `down_intfs_before_sess_sync` - List of interfaces to be turned down before session synchronization is complete.
* `hb_interval` - Heartbeat interval (1 - 20 (100*ms). Increase to reduce false positives.
* `hb_lost_threshold` - Lost heartbeat threshold (1 - 60). Increase to reduce false positives.
* `ike_heartbeat_interval` - Ike-Heartbeat-Interval.
* `ike_monitor` - Ike-Monitor. Valid values: `disable`, `enable`.

* `ike_monitor_interval` - Ike-Monitor-Interval.
* `ike_use_rfc6311` - Ike-Use-Rfc6311. Valid values: `disable`, `enable`.

* `ipsec_tunnel_sync` - Enable/disable IPsec tunnel synchronization. Valid values: `disable`, `enable`.

* `peerip` - IP address of the interface on the peer unit that is used for the session synchronization link.
* `peervd` - VDOM that contains the session synchronization link interface on the peer unit. Usually both peers would have the same peervd.
* `secondary_add_ipsec_routes` - Enable/disable IKE route announcement on the backup unit. Valid values: `disable`, `enable`.

* `session_sync_filter` - Session-Sync-Filter. The structure of `session_sync_filter` block is documented below.
* `sync_id` - Sync ID.
* `syncvd` - Sessions from these VDOMs are synchronized using this session synchronization configuration.

The `session_sync_filter` block supports:

* `custom_service` - Custom-Service. The structure of `custom_service` block is documented below.
* `dstaddr` - Only sessions to this IPv4 address are synchronized.
* `dstaddr6` - Only sessions to this IPv6 address are synchronized.
* `dstintf` - Only sessions to this interface are synchronized.
* `srcaddr` - Only sessions from this IPv4 address are synchronized.
* `srcaddr6` - Only sessions from this IPv6 address are synchronized.
* `srcintf` - Only sessions from this interface are synchronized.

The `custom_service` block supports:

* `dst_port_range` - Custom service destination port range.
* `id` - Custom service ID.
* `src_port_range` - Custom service source port range.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{sync_id}}.

## Import

System StandaloneClusterClusterPeer can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_standalonecluster_clusterpeer.labelname {{sync_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

