---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_ha"
description: |-
  Configure HA.
---

# fmgdevice_system_ha
Configure HA.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `frup_settings`: `fmgdevice_system_ha_frupsettings`
>- `ha_mgmt_interfaces`: `fmgdevice_system_ha_hamgmtinterfaces`
>- `secondary_vcluster`: `fmgdevice_system_ha_secondaryvcluster`
>- `unicast_peers`: `fmgdevice_system_ha_unicastpeers`
>- `vcluster`: `fmgdevice_system_ha_vcluster`



## Example Usage

```hcl
resource "fmgdevice_system_ha" "trname" {
  arps                     = 10
  arps_interval            = 10
  authentication           = "disable"
  board_failover_tolerance = 10
  chassis_id               = 10
  device_name              = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `arps` - Number of gratuitous ARPs (1 - 60). Lower to reduce traffic. Higher to reduce failover time.
* `arps_interval` - Time between gratuitous ARPs  (1 - 20 sec). Lower to reduce failover time. Higher to reduce traffic.
* `authentication` - Enable/disable heartbeat message authentication. Valid values: `disable`, `enable`.

* `board_failover_tolerance` - Worker board failure failover threshold.
* `chassis_id` - chassis id
* `cpu_threshold` - Dynamic weighted load balancing CPU usage weight and high and low thresholds.
* `encryption` - Enable/disable heartbeat message encryption. Valid values: `disable`, `enable`.

* `frup` - Frup. Valid values: `disable`, `enable`.

* `frup_settings` - Frup-Settings. The structure of `frup_settings` block is documented below.
* `evpn_ttl` - HA EVPN FDB TTL on primary box (5 - 3600 sec).
* `failover_hold_time` - Time to wait before failover (0 - 300 sec, default = 0), to avoid flip.
* `ftp_proxy_threshold` - Dynamic weighted load balancing weight and high and low number of FTP proxy sessions.
* `gratuitous_arps` - Enable/disable gratuitous ARPs. Disable if link-failed-signal enabled. Valid values: `disable`, `enable`.

* `group_id` - HA group ID  (0 - 1023;  or 0 - 7 when there are more than 2 vclusters). Must be the same for all members.
* `group_name` - Cluster group name. Must be the same for all members.
* `ha_direct` - Enable/disable using ha-mgmt interface for syslog, remote authentication (RADIUS), FortiAnalyzer, FortiSandbox, sFlow, and Netflow. Valid values: `disable`, `enable`.

* `ha_eth_type` - HA heartbeat packet Ethertype (4-digit hex).
* `ha_mgmt_interfaces` - Ha-Mgmt-Interfaces. The structure of `ha_mgmt_interfaces` block is documented below.
* `ha_mgmt_status` - Enable to reserve interfaces to manage individual cluster units. Valid values: `disable`, `enable`.

* `ha_port_dtag_mode` - HA port double-tagging mode. Valid values: `proprietary`, `double-tagging`.

* `ha_port_outer_tpid` - Set HA port outer tpid. Valid values: `0x8100`, `0x88a8`, `0x9100`.

* `ha_uptime_diff_margin` - Normally you would only reduce this value for failover testing.
* `hb_interval` - Time between sending heartbeat packets (1 - 20). Increase to reduce false positives.
* `hb_interval_in_milliseconds` - Units of heartbeat interval time between sending heartbeat packets. Default is 100ms. Valid values: `100ms`, `10ms`.

* `hb_lost_threshold` - Number of lost heartbeats to signal a failure (1 - 60). Increase to reduce false positives.
* `hbdev` - Heartbeat interfaces. Must be the same for all members.
* `hbdev_second_vlan_id` - second VLAN id to use for HA heartbeat (1-4094).
* `hbdev_vlan_id` - VLAN id to use for HA heartbeat (1-4094).
* `hc_eth_type` - Transparent mode HA heartbeat packet Ethertype (4-digit hex).
* `hello_holddown` - Time to wait before changing from hello to work state (5 - 300 sec).
* `http_proxy_threshold` - Dynamic weighted load balancing weight and high and low number of HTTP proxy sessions.
* `hw_session_hold_time` - Time to hold sessions before purging on secondary node (0 - 180 sec, default = 10).
* `hw_session_sync_delay` - Time to wait before session sync starts on primary node (0 - 3600 sec, default = 150).
* `hw_session_sync_dev` - Hardware session sync interface.
* `imap_proxy_threshold` - Dynamic weighted load balancing weight and high and low number of IMAP proxy sessions.
* `inter_cluster_session_sync` - Enable/disable synchronization of sessions among HA clusters. Valid values: `disable`, `enable`.

* `ipsec_phase2_proposal` - IPsec phase2 proposal. Valid values: `aes128-sha1`, `aes128-sha256`, `aes128-sha384`, `aes128-sha512`, `aes192-sha1`, `aes192-sha256`, `aes192-sha384`, `aes192-sha512`, `aes256-sha1`, `aes256-sha256`, `aes256-sha384`, `aes256-sha512`, `aes128gcm`, `aes256gcm`, `chacha20poly1305`.

* `key` - Key.
* `l2ep_eth_type` - Telnet session HA heartbeat packet Ethertype (4-digit hex).
* `link_failed_signal` - Enable to shut down all interfaces for 1 sec after a failover. Use if gratuitous ARPs do not update network. Valid values: `disable`, `enable`.

* `load_balance_all` - Enable to load balance TCP sessions. Disable to load balance proxy sessions only. Valid values: `disable`, `enable`.

* `logical_sn` - Enable/disable usage of the logical serial number. Valid values: `disable`, `enable`.

* `memory_based_failover` - Enable/disable memory based failover. Valid values: `disable`, `enable`.

* `memory_compatible_mode` - Enable/disable memory compatible mode. Valid values: `disable`, `enable`.

* `memory_failover_flip_timeout` - Time to wait between subsequent memory based failovers in minutes (6 - 2147483647, default = 6).
* `memory_failover_monitor_period` - Duration of high memory usage before memory based failover is triggered in seconds (1 - 300, default = 60).
* `memory_failover_sample_rate` - Rate at which memory usage is sampled in order to measure memory usage in seconds (1 - 60, default = 1).
* `memory_failover_threshold` - Memory usage threshold to trigger memory based failover (0 means using conserve mode threshold in system.global).
* `memory_threshold` - Dynamic weighted load balancing memory usage weight and high and low thresholds.
* `minimum_worker_threshold` - The minimum number of operating workers to cause a content clustering chassis failover.
* `mode` - HA mode. Must be the same for all members. FGSP requires standalone. Valid values: `standalone`, `a-a`, `a-p`, `config-sync-only`, `active-passive`.

* `monitor` - Interfaces to check for port monitoring (or link failure).
* `multicast_ttl` - HA multicast TTL on primary (5 - 3600 sec).
* `nntp_proxy_threshold` - Dynamic weighted load balancing weight and high and low number of NNTP proxy sessions.
* `override` - Enable and increase the priority of the unit that should always be primary (master). Valid values: `disable`, `enable`.

* `override_wait_time` - Delay negotiating if override is enabled (0 - 3600 sec). Reduces how often the cluster negotiates.
* `password` - Cluster password. Must be the same for all members.
* `pingserver_failover_threshold` - Remote IP monitoring failover threshold (0 - 50).
* `pingserver_flip_timeout` - Time to wait in minutes before renegotiating after a remote IP monitoring failover.
* `pingserver_monitor_interface` - Interfaces to check for remote IP monitoring.
* `pingserver_secondary_force_reset` - Enable to force the cluster to negotiate after a remote IP monitoring failover. Valid values: `disable`, `enable`.

* `pingserver_slave_force_reset` - Pingserver-Slave-Force-Reset. Valid values: `disable`, `enable`.

* `pop3_proxy_threshold` - Dynamic weighted load balancing weight and high and low number of POP3 proxy sessions.
* `priority` - Increase the priority to select the primary unit (0 - 255).
* `route_hold` - Time to wait between routing table updates to the cluster (0 - 3600 sec).
* `route_ttl` - TTL for primary unit routes (5 - 3600 sec). Increase to maintain active routes during failover.
* `route_wait` - Time to wait before sending new routes to the cluster (0 - 3600 sec).
* `schedule` - Type of A-A load balancing. Use none if you have external load balancers. Valid values: `none`, `hub`, `leastconnection`, `round-robin`, `weight-round-robin`, `random`, `ip`, `ipport`.

* `secondary_switch_standby` - Enable to force content clustering subordinate unit standby mode. Valid values: `disable`, `enable`.

* `secondary_vcluster` - Secondary-Vcluster. The structure of `secondary_vcluster` block is documented below.
* `session_pickup` - Enable/disable session pickup. Enabling it can reduce session down time when fail over happens. Valid values: `disable`, `enable`.

* `session_pickup_connectionless` - Enable/disable UDP and ICMP session sync. Valid values: `disable`, `enable`.

* `session_pickup_delay` - Enable to sync sessions longer than 30 sec. Only longer lived sessions need to be synced. Valid values: `disable`, `enable`.

* `session_pickup_expectation` - Enable/disable session helper expectation session sync for FGSP. Valid values: `disable`, `enable`.

* `session_pickup_nat` - Enable/disable NAT session sync for FGSP. Valid values: `disable`, `enable`.

* `session_sync_dev` - Offload session-sync process to kernel and sync sessions using connected interface(s) directly.
* `slave_switch_standby` - Slave-Switch-Standby. Valid values: `disable`, `enable`.

* `smtp_proxy_threshold` - Dynamic weighted load balancing weight and high and low number of SMTP proxy sessions.
* `ssd_failover` - Enable/disable automatic HA failover on SSD disk failure. Valid values: `disable`, `enable`.

* `standalone_config_sync` - Enable/disable FGSP configuration synchronization. Valid values: `disable`, `enable`.

* `standalone_mgmt_vdom` - Enable/disable standalone management VDOM. Valid values: `disable`, `enable`.

* `sync_config` - Enable/disable configuration synchronization. Valid values: `disable`, `enable`.

* `sync_packet_balance` - Enable/disable HA packet distribution to multiple CPUs. Valid values: `disable`, `enable`.

* `unicast_gateway` - Default route gateway for unicast interface.
* `unicast_hb` - Enable/disable unicast heartbeat. Valid values: `disable`, `enable`.

* `unicast_hb_netmask` - Unicast heartbeat netmask.
* `unicast_hb_peerip` - Unicast heartbeat peer IP.
* `unicast_peers` - Unicast-Peers. The structure of `unicast_peers` block is documented below.
* `unicast_status` - Enable/disable unicast connection. Valid values: `disable`, `enable`.

* `uninterruptible_primary_wait` - Number of minutes the primary HA unit waits before the secondary HA unit is considered upgraded and the system is started before starting its own upgrade (15 - 300, default = 30).
* `uninterruptible_upgrade` - Enable to upgrade a cluster without blocking network traffic. Valid values: `disable`, `enable`.

* `vcluster_id` - Vcluster-Id.
* `vcluster2` - Enable/disable virtual cluster 2 for virtual clustering. Valid values: `disable`, `enable`.

* `vdom` - VDOMs in virtual cluster 1.
* `upgrade_mode` - The mode to upgrade a cluster. Valid values: `simultaneous`, `uninterruptible`, `local-only`, `secondary-only`.

* `vcluster` - Vcluster. The structure of `vcluster` block is documented below.
* `vcluster_status` - Enable/disable virtual cluster for virtual clustering. Valid values: `disable`, `enable`.

* `weight` - Weight-round-robin weight for each cluster unit. Syntax &lt;priority&gt; &lt;weight&gt;.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `frup_settings` block supports:

* `active_interface` - Active-Interface.
* `active_switch_port` - Active-Switch-Port. Valid values: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `13`, `14`, `15`, `16`, `17`, `18`, `19`, `20`, `21`, `22`, `23`, `24`, `25`, `26`, `27`, `28`, `29`, `30`, `31`, `32`, `33`, `34`, `35`, `36`, `37`, `38`, `39`, `40`, `41`, `42`, `43`, `44`, `45`, `46`, `47`, `48`, `49`, `50`, `51`, `52`, `53`, `54`, `55`, `56`, `57`, `58`, `59`, `60`, `61`, `62`, `63`, `64`, `65`, `66`, `67`, `68`, `69`, `70`, `71`, `72`, `73`, `74`, `75`, `76`, `77`, `78`, `79`, `80`, `81`, `82`, `83`, `84`.

* `backup_interface` - Backup-Interface.

The `ha_mgmt_interfaces` block supports:

* `dst` - Default route destination for reserved HA management interface.
* `gateway` - Default route gateway for reserved HA management interface.
* `gateway6` - Default IPv6 gateway for reserved HA management interface.
* `id` - Table ID.
* `interface` - Interface to reserve for HA management.

The `secondary_vcluster` block supports:

* `monitor` - Interfaces to check for port monitoring (or link failure).
* `override` - Enable and increase the priority of the unit that should always be primary. Valid values: `disable`, `enable`.

* `override_wait_time` - Delay negotiating if override is enabled (0 - 3600 sec). Reduces how often the cluster negotiates.
* `pingserver_failover_threshold` - Remote IP monitoring failover threshold (0 - 50).
* `pingserver_monitor_interface` - Interfaces to check for remote IP monitoring.
* `pingserver_secondary_force_reset` - Enable to force the cluster to negotiate after a remote IP monitoring failover. Valid values: `disable`, `enable`.

* `pingserver_slave_force_reset` - Pingserver-Slave-Force-Reset. Valid values: `disable`, `enable`.

* `priority` - Increase the priority to select the primary unit (0 - 255).
* `vcluster_id` - Vcluster-Id.
* `vdom` - VDOMs in virtual cluster 2.

The `unicast_peers` block supports:

* `id` - Table ID.
* `peer_ip` - Unicast peer IP.

The `vcluster` block supports:

* `monitor` - Interfaces to check for port monitoring (or link failure).
* `override` - Enable and increase the priority of the unit that should always be primary (master). Valid values: `disable`, `enable`.

* `override_wait_time` - Delay negotiating if override is enabled (0 - 3600 sec). Reduces how often the cluster negotiates.
* `pingserver_failover_threshold` - Remote IP monitoring failover threshold (0 - 50).
* `pingserver_flip_timeout` - Time to wait in minutes before renegotiating after a remote IP monitoring failover.
* `pingserver_monitor_interface` - Interfaces to check for remote IP monitoring.
* `pingserver_secondary_force_reset` - Enable to force the cluster to negotiate after a remote IP monitoring failover. Valid values: `disable`, `enable`.

* `pingserver_slave_force_reset` - Pingserver-Slave-Force-Reset. Valid values: `disable`, `enable`.

* `priority` - Increase the priority to select the primary unit (0 - 255).
* `vcluster_id` - ID.
* `vdom` - Virtual domain(s) in the virtual cluster.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Ha can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_ha.labelname SystemHa
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

