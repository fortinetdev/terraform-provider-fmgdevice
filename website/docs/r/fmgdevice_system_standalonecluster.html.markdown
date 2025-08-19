---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_standalonecluster"
description: |-
  Configure FortiGate Session Life Support Protocol (FGSP) cluster attributes.
---

# fmgdevice_system_standalonecluster
Configure FortiGate Session Life Support Protocol (FGSP) cluster attributes.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `cluster_peer`: `fmgdevice_system_standalonecluster_clusterpeer`
>- `monitor_prefix`: `fmgdevice_system_standalonecluster_monitorprefix`



## Example Usage

```hcl
resource "fmgdevice_system_standalonecluster" "trname" {
  asymmetric_traffic_control = "strict-anti-replay"
  cluster_peer {
    down_intfs_before_sess_sync = ["your own value"]
    hb_interval                 = 10
    hb_lost_threshold           = 10
    ike_heartbeat_interval      = 10
    ike_monitor                 = "disable"
    ike_monitor_interval        = 10
    ike_use_rfc6311             = "disable"
    ipsec_tunnel_sync           = "disable"
    peerip                      = "your own value"
    peervd                      = ["your own value"]
    secondary_add_ipsec_routes  = "enable"
    session_sync_filter {
      custom_service {
        dst_port_range = "your own value"
        fosid          = 10
        src_port_range = "your own value"
      }

      dstaddr  = ["your own value"]
      dstaddr6 = "your own value"
      dstintf  = ["your own value"]
      srcaddr  = ["your own value"]
      srcaddr6 = "your own value"
      srcintf  = ["your own value"]
    }

    sync_id = 10
    syncvd  = ["your own value"]
  }

  data_intf_session_sync_dev = "your own value"
  encryption                 = "enable"
  group_member_id            = 10
  device_name                = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `asymmetric_traffic_control` - Asymmetric traffic control mode. Valid values: `cps-preferred`, `strict-anti-replay`.

* `cluster_peer` - Cluster-Peer. The structure of `cluster_peer` block is documented below.
* `data_intf_session_sync_dev` - Reserve data interfaces for session sync only.
* `encryption` - Enable/disable encryption when synchronizing sessions. Valid values: `disable`, `enable`.

* `group_member_id` - Cluster member ID (0 - 15).
* `layer2_connection` - Indicate whether layer 2 connections are present among FGSP members. Valid values: `unavailable`, `available`.

* `monitor_interface` - Configure a list of interfaces on which to monitor itself. Monitoring is performed on the status of the interface.
* `monitor_prefix` - Monitor-Prefix. The structure of `monitor_prefix` block is documented below.
* `pingsvr_monitor_interface` - List of pingsvr monitor interface to check for remote IP monitoring.
* `psksecret` - Pre-shared secret for session synchronization (ASCII string or hexadecimal encoded with a leading 0x).
* `session_sync_dev` - Offload session-sync process to kernel and sync sessions using connected interface(s) directly.
* `standalone_group_id` - Cluster group ID (0 - 255). Must be the same for all members.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `cluster_peer` block supports:

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

The `monitor_prefix` block supports:

* `id` - ID.
* `prefix` - Prefix.
* `vdom` - VDOM name.
* `vrf` - VRF ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System StandaloneCluster can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_standalonecluster.labelname SystemStandaloneCluster
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

