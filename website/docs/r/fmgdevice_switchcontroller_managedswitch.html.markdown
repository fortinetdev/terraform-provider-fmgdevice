---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_managedswitch"
description: |-
  Configure FortiSwitch devices that are managed by this FortiGate.
---

# fmgdevice_switchcontroller_managedswitch
Configure FortiSwitch devices that are managed by this FortiGate.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `n802_1x_settings`: `fmgdevice_switchcontroller_managedswitch_8021xsettings`
>- `custom_command`: `fmgdevice_switchcontroller_managedswitch_customcommand`
>- `dhcp_snooping_static_client`: `fmgdevice_switchcontroller_managedswitch_dhcpsnoopingstaticclient`
>- `igmp_snooping`: `fmgdevice_switchcontroller_managedswitch_igmpsnooping`
>- `ip_source_guard`: `fmgdevice_switchcontroller_managedswitch_ipsourceguard`
>- `mirror`: `fmgdevice_switchcontroller_managedswitch_mirror`
>- `ports`: `fmgdevice_switchcontroller_managedswitch_ports`
>- `remote_log`: `fmgdevice_switchcontroller_managedswitch_remotelog`
>- `route_offload_router`: `fmgdevice_switchcontroller_managedswitch_routeoffloadrouter`
>- `snmp_community`: `fmgdevice_switchcontroller_managedswitch_snmpcommunity`
>- `snmp_sysinfo`: `fmgdevice_switchcontroller_managedswitch_snmpsysinfo`
>- `snmp_trap_threshold`: `fmgdevice_switchcontroller_managedswitch_snmptrapthreshold`
>- `snmp_user`: `fmgdevice_switchcontroller_managedswitch_snmpuser`
>- `static_mac`: `fmgdevice_switchcontroller_managedswitch_staticmac`
>- `storm_control`: `fmgdevice_switchcontroller_managedswitch_stormcontrol`
>- `stp_instance`: `fmgdevice_switchcontroller_managedswitch_stpinstance`
>- `stp_settings`: `fmgdevice_switchcontroller_managedswitch_stpsettings`
>- `switch_log`: `fmgdevice_switchcontroller_managedswitch_switchlog`
>- `vlan`: `fmgdevice_switchcontroller_managedswitch_vlan`



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_managedswitch" "trname" {
  n802_1x_settings {
    link_down_auth                = "no-action"
    local_override                = "enable"
    mab_reauth                    = "disable"
    mac_called_station_delimiter  = "none"
    mac_calling_station_delimiter = "single-hyphen"
    mac_case                      = "uppercase"
    mac_password_delimiter        = "hyphen"
    mac_username_delimiter        = "single-hyphen"
    max_reauth_attempt            = 10
    reauth_period                 = 10
    tx_period                     = 10
  }

  _platform      = "your own value"
  access_profile = ["your own value"]
  custom_command {
    command_entry = "your own value"
    command_name  = ["your own value"]
  }

  delayed_restart_trigger = 10
  switch_id               = "your own value"
  device_name             = var.device_name # not required if setting is at provider
  device_vdom             = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `n802_1x_settings` - 802-1X-Settings. The structure of `n802_1x_settings` block is documented below.
* `_platform` - _Platform.
* `access_profile` - FortiSwitch access profile.
* `custom_command` - Custom-Command. The structure of `custom_command` block is documented below.
* `delayed_restart_trigger` - Delayed restart triggered for this FortiSwitch.
* `description` - Description.
* `dhcp_server_access_list` - DHCP snooping server access list. Valid values: `disable`, `enable`, `global`.

* `dhcp_snooping_static_client` - Dhcp-Snooping-Static-Client. The structure of `dhcp_snooping_static_client` block is documented below.
* `directly_connected` - Directly-Connected.
* `dynamic_capability` - List of features this FortiSwitch supports (not configurable) that is sent to the FortiGate device for subsequent configuration initiated by the FortiGate device.
* `dynamically_discovered` - Dynamically-Discovered.
* `firmware_provision` - Enable/disable provisioning of firmware to FortiSwitches on join connection. Valid values: `disable`, `enable`.

* `firmware_provision_latest` - Enable/disable one-time automatic provisioning of the latest firmware version. Valid values: `disable`, `once`.

* `firmware_provision_version` - Firmware version to provision to this FortiSwitch on bootup (major.minor.build, i.e. 6.2.1234).
* `flow_identity` - Flow-tracking netflow ipfix switch identity in hex format(00000000-FFFFFFFF default=0).
* `fsw_wan1_admin` - FortiSwitch WAN1 admin status; enable to authorize the FortiSwitch as a managed switch. Valid values: `disable`, `enable`, `discovered`.

* `fsw_wan1_peer` - FortiSwitch WAN1 peer port.
* `fsw_wan2_admin` - FortiSwitch WAN2 admin status; enable to authorize the FortiSwitch as a managed switch. Valid values: `disable`, `enable`, `discovered`.

* `fsw_wan2_peer` - FortiSwitch WAN2 peer port.
* `igmp_snooping` - Igmp-Snooping. The structure of `igmp_snooping` block is documented below.
* `ip_source_guard` - Ip-Source-Guard. The structure of `ip_source_guard` block is documented below.
* `l3_discovered` - L3-Discovered.
* `max_allowed_trunk_members` - FortiSwitch maximum allowed trunk members.
* `mclag_igmp_snooping_aware` - Enable/disable MCLAG IGMP-snooping awareness. Valid values: `disable`, `enable`.

* `mgmt_mode` - FortiLink management mode.
* `mirror` - Mirror. The structure of `mirror` block is documented below.
* `name` - Managed-switch name.
* `override_snmp_community` - Enable/disable overriding the global SNMP communities. Valid values: `disable`, `enable`.

* `override_snmp_sysinfo` - Enable/disable overriding the global SNMP system information. Valid values: `disable`, `enable`.

* `override_snmp_trap_threshold` - Enable/disable overriding the global SNMP trap threshold values. Valid values: `disable`, `enable`.

* `override_snmp_user` - Enable/disable overriding the global SNMP users. Valid values: `disable`, `enable`.

* `owner_vdom` - VDOM which owner of port belongs to.
* `poe_detection_type` - PoE detection type for FortiSwitch.
* `poe_lldp_detection` - Enable/disable PoE LLDP detection. Valid values: `disable`, `enable`.

* `poe_pre_standard_detection` - Enable/disable PoE pre-standard detection. Valid values: `disable`, `enable`.

* `ports` - Ports. The structure of `ports` block is documented below.
* `pre_provisioned` - Pre-provisioned managed switch.
* `ptp_profile` - PTP profile configuration.
* `ptp_status` - Enable/disable PTP profile on this FortiSwitch. Valid values: `disable`, `enable`.

* `purdue_level` - Purdue Level of this FortiSwitch. Valid values: `1`, `2`, `3`, `4`, `5`, `1.5`, `2.5`, `3.5`, `5.5`.

* `qos_drop_policy` - Set QoS drop-policy. Valid values: `taildrop`, `random-early-detection`.

* `qos_red_probability` - Set QoS RED/WRED drop probability.
* `radius_nas_ip` - NAS-IP address.
* `radius_nas_ip_override` - Use locally defined NAS-IP. Valid values: `disable`, `enable`.

* `remote_log` - Remote-Log. The structure of `remote_log` block is documented below.
* `route_offload` - Enable/disable route offload on this FortiSwitch. Valid values: `disable`, `enable`.

* `route_offload_mclag` - Enable/disable route offload MCLAG on this FortiSwitch. Valid values: `disable`, `enable`.

* `route_offload_router` - Route-Offload-Router. The structure of `route_offload_router` block is documented below.
* `sn` - Managed-switch serial number.
* `snmp_community` - Snmp-Community. The structure of `snmp_community` block is documented below.
* `snmp_sysinfo` - Snmp-Sysinfo. The structure of `snmp_sysinfo` block is documented below.
* `snmp_trap_threshold` - Snmp-Trap-Threshold. The structure of `snmp_trap_threshold` block is documented below.
* `snmp_user` - Snmp-User. The structure of `snmp_user` block is documented below.
* `staged_image_version` - Staged image version for FortiSwitch.
* `static_mac` - Static-Mac. The structure of `static_mac` block is documented below.
* `storm_control` - Storm-Control. The structure of `storm_control` block is documented below.
* `stp_instance` - Stp-Instance. The structure of `stp_instance` block is documented below.
* `stp_settings` - Stp-Settings. The structure of `stp_settings` block is documented below.
* `switch_device_tag` - User definable label/tag.
* `switch_dhcp_opt43_key` - DHCP option43 key.
* `switch_id` - Managed-switch name.
* `switch_log` - Switch-Log. The structure of `switch_log` block is documented below.
* `switch_profile` - FortiSwitch profile.
* `tdr_supported` - Tdr-Supported.
* `tunnel_discovered` - Tunnel-Discovered.
* `type` - Indication of switch type, physical or virtual. Valid values: `physical`, `virtual`.

* `version` - FortiSwitch version.
* `vlan` - Vlan. The structure of `vlan` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `n802_1x_settings` block supports:

* `link_down_auth` - Authentication state to set if a link is down. Valid values: `set-unauth`, `no-action`.

* `local_override` - Enable to override global 802.1X settings on individual FortiSwitches. Valid values: `disable`, `enable`.

* `mab_reauth` - Enable or disable MAB reauthentication settings. Valid values: `disable`, `enable`.

* `mac_called_station_delimiter` - MAC called station delimiter (default = hyphen). Valid values: `hyphen`, `single-hyphen`, `colon`, `none`.

* `mac_calling_station_delimiter` - MAC calling station delimiter (default = hyphen). Valid values: `hyphen`, `single-hyphen`, `colon`, `none`.

* `mac_case` - MAC case (default = lowercase). Valid values: `uppercase`, `lowercase`.

* `mac_password_delimiter` - MAC authentication password delimiter (default = hyphen). Valid values: `hyphen`, `single-hyphen`, `colon`, `none`.

* `mac_username_delimiter` - MAC authentication username delimiter (default = hyphen). Valid values: `hyphen`, `single-hyphen`, `colon`, `none`.

* `max_reauth_attempt` - Maximum number of authentication attempts (0 - 15, default = 3).
* `reauth_period` - Reauthentication time interval (1 - 1440 min, default = 60, 0 = disable).
* `tx_period` - 802.1X Tx period (seconds, default=30).

The `custom_command` block supports:

* `command_entry` - List of FortiSwitch commands.
* `command_name` - Names of commands to be pushed to this FortiSwitch device, as configured under config switch-controller custom-command.

The `dhcp_snooping_static_client` block supports:

* `ip` - Client static IP address.
* `mac` - Client MAC address.
* `name` - Client name.
* `port` - Interface name.
* `vlan` - VLAN name.

The `igmp_snooping` block supports:

* `aging_time` - Maximum time to retain a multicast snooping entry for which no packets have been seen (15 - 3600 sec, default = 300).
* `flood_unknown_multicast` - Enable/disable unknown multicast flooding. Valid values: `disable`, `enable`.

* `local_override` - Enable/disable overriding the global IGMP snooping configuration. Valid values: `disable`, `enable`.

* `vlans` - Vlans. The structure of `vlans` block is documented below.

The `vlans` block supports:

* `proxy` - IGMP snooping proxy for the VLAN interface. Valid values: `disable`, `enable`, `global`.

* `querier` - Enable/disable IGMP snooping querier for the VLAN interface. Valid values: `disable`, `enable`.

* `querier_addr` - IGMP snooping querier address.
* `version` - IGMP snooping querying version.
* `vlan_name` - List of FortiSwitch VLANs.

The `ip_source_guard` block supports:

* `binding_entry` - Binding-Entry. The structure of `binding_entry` block is documented below.
* `description` - Description.
* `port` - Ingress interface to which source guard is bound.

The `binding_entry` block supports:

* `entry_name` - Configure binding pair.
* `ip` - Source IP for this rule.
* `mac` - MAC address for this rule.

The `mirror` block supports:

* `dst` - Destination port.
* `name` - Mirror name.
* `src_egress` - Source egress interfaces.
* `src_ingress` - Source ingress interfaces.
* `status` - Active/inactive mirror configuration. Valid values: `inactive`, `active`.

* `switching_packet` - Enable/disable switching functionality when mirroring. Valid values: `disable`, `enable`.


The `ports` block supports:

* `access_mode` - Access mode of the port. Valid values: `normal`, `nac`, `dynamic`, `static`.

* `acl_group` - ACL groups on this port.
* `aggregator_mode` - LACP member select mode. Valid values: `bandwidth`, `count`.

* `allow_arp_monitor` - Enable/Disable allow ARP monitor. Valid values: `disable`, `enable`.

* `allowed_vlans` - Configure switch port tagged VLANs.
* `allowed_vlans_all` - Enable/disable all defined vlans on this port. Valid values: `disable`, `enable`.

* `arp_inspection_trust` - Trusted or untrusted dynamic ARP inspection. Valid values: `untrusted`, `trusted`.

* `authenticated_port` - Authenticated-Port.
* `bundle` - Enable/disable Link Aggregation Group (LAG) bundling for non-FortiLink interfaces. Valid values: `disable`, `enable`.

* `description` - Description for port.
* `dhcp_snoop_option82_override` - Dhcp-Snoop-Option82-Override. The structure of `dhcp_snoop_option82_override` block is documented below.
* `dhcp_snoop_option82_trust` - Enable/disable allowance of DHCP with option-82 on untrusted interface. Valid values: `disable`, `enable`.

* `dhcp_snooping` - Trusted or untrusted DHCP-snooping interface. Valid values: `trusted`, `untrusted`.

* `discard_mode` - Configure discard mode for port. Valid values: `none`, `all-untagged`, `all-tagged`.

* `dsl_profile` - DSL policy configuration.
* `edge_port` - Enable/disable this interface as an edge port, bridging connections between workstations and/or computers. Valid values: `disable`, `enable`.

* `export_tags` - Configure export tag(s) for FortiSwitch port when exported to a virtual port pool.
* `encrypted_port` - Encrypted-Port.
* `export_to` - Export managed-switch port to a tenant VDOM.
* `export_to_pool` - Switch controller export port to pool-list.
* `export_to_pool_flag` - Switch controller export port to pool-list.
* `fallback_port` - LACP fallback port.
* `fec_capable` - FEC capable.
* `fec_state` - State of forward error correction. Valid values: `disabled`, `cl74`, `cl91`, `detect-by-module`.

* `fgt_peer_device_name` - Fgt-Peer-Device-Name.
* `fgt_peer_port_name` - Fgt-Peer-Port-Name.
* `fiber_port` - Fiber-Port.
* `flags` - Flags.
* `flap_duration` - Period over which flap events are calculated (seconds).
* `flap_rate` - Number of stage change events needed within flap-duration.
* `flap_timeout` - Flap guard disabling protection (min).
* `flapguard` - Enable/disable flap guard. Valid values: `disable`, `enable`.

* `flow_control` - Flow control direction. Valid values: `disable`, `tx`, `rx`, `both`.

* `fortilink_port` - Fortilink-Port.
* `igmp_snooping` - Set IGMP snooping mode for the physical port interface. Valid values: `disable`, `enable`.

* `fortiswitch_acls` - ACLs on this port.
* `igmp_snooping_flood_reports` - Enable/disable flooding of IGMP reports to this interface when igmp-snooping enabled. Valid values: `disable`, `enable`.

* `igmps_flood_reports` - Igmps-Flood-Reports. Valid values: `disable`, `enable`.

* `igmps_flood_traffic` - Igmps-Flood-Traffic. Valid values: `disable`, `enable`.

* `interface_tags` - Tag(s) associated with the interface for various features including virtual port pool, dynamic port policy.
* `ip_source_guard` - Enable/disable IP source guard. Valid values: `disable`, `enable`.

* `isl_local_trunk_name` - Isl-Local-Trunk-Name.
* `isl_peer_device_name` - Isl-Peer-Device-Name.
* `isl_peer_device_sn` - Isl-Peer-Device-Sn.
* `isl_peer_port_name` - Isl-Peer-Port-Name.
* `lacp_speed` - End Link Aggregation Control Protocol (LACP) messages every 30 seconds (slow) or every second (fast). Valid values: `slow`, `fast`.

* `learning_limit` - Limit the number of dynamic MAC addresses on this Port (1 - 128, 0 = no limit, default).
* `link_status` - Link-Status. Valid values: `down`, `up`.

* `lldp_profile` - LLDP port TLV profile.
* `lldp_status` - LLDP transmit and receive status. Valid values: `disable`, `rx-only`, `tx-only`, `tx-rx`.

* `log_mac_event` - Enable/disable logging for dynamic MAC address events. Valid values: `disable`, `enable`.

* `loop_guard` - Enable/disable loop-guard on this interface, an STP optimization used to prevent network loops. Valid values: `disabled`, `enabled`.

* `loop_guard_timeout` - Loop-guard timeout (0 - 120 min, default = 45).
* `mac_addr` - Port/Trunk MAC.
* `matched_dpp_intf_tags` - Matched interface tags in the dynamic port policy.
* `matched_dpp_policy` - Matched child policy in the dynamic port policy.
* `max_bundle` - Maximum size of LAG bundle (1 - 24, default = 24).
* `mcast_snooping_flood_traffic` - Enable/disable flooding of IGMP snooping traffic to this interface. Valid values: `disable`, `enable`.

* `mclag` - Enable/disable multi-chassis link aggregation (MCLAG). Valid values: `disable`, `enable`.

* `mclag_icl_port` - Mclag-Icl-Port.
* `media_type` - Media-Type.
* `member_withdrawal_behavior` - Port behavior after it withdraws because of loss of control packets. Valid values: `forward`, `block`.

* `members` - Aggregated LAG bundle interfaces.
* `min_bundle` - Minimum size of LAG bundle (1 - 24, default = 1).
* `mode` - LACP mode: ignore and do not send control messages, or negotiate 802.3ad aggregation passively or actively. Valid values: `static`, `lacp-passive`, `lacp-active`.

* `p2p_port` - P2P-Port.
* `packet_sample_rate` - Packet sampling rate (0 - 99999 p/sec).
* `packet_sampler` - Enable/disable packet sampling on this interface. Valid values: `disabled`, `enabled`.

* `pause_meter` - Configure ingress pause metering rate, in kbps (default = 0, disabled).
* `pause_meter_resume` - Resume threshold for resuming traffic on ingress port. Valid values: `25%`, `50%`, `75%`.

* `pd_capable` - Powered device capable.
* `poe_capable` - PoE capable.
* `poe_max_power` - Poe-Max-Power.
* `poe_mode_bt_cabable` - PoE mode IEEE 802.3BT capable.
* `poe_port_mode` - Configure PoE port mode. Valid values: `ieee802-3af`, `ieee802-3at`, `ieee802-3bt`.

* `poe_port_power` - Configure PoE port power. Valid values: `normal`, `perpetual`, `perpetual-fast`.

* `poe_port_priority` - Configure PoE port priority. Valid values: `critical-priority`, `high-priority`, `low-priority`, `medium-priority`.

* `poe_pre_standard_detection` - Enable/disable PoE pre-standard detection. Valid values: `disable`, `enable`.

* `poe_standard` - Poe-Standard.
* `poe_status` - Enable/disable PoE status. Valid values: `disable`, `enable`.

* `port_name` - Switch port name.
* `port_number` - Port-Number.
* `port_owner` - Switch port name.
* `port_policy` - Switch controller dynamic port policy from available options.
* `port_prefix_type` - Port-Prefix-Type.
* `port_security_policy` - Switch controller authentication policy to apply to this managed switch from available options.
* `port_selection_criteria` - Algorithm for aggregate port selection. Valid values: `src-mac`, `dst-mac`, `src-dst-mac`, `src-ip`, `dst-ip`, `src-dst-ip`.

* `ptp_policy` - PTP policy configuration.
* `ptp_status` - Enable/disable PTP policy on this FortiSwitch port. Valid values: `disable`, `enable`.

* `qnq` - 802.1AD VLANs in the VDom.
* `qos_policy` - Switch controller QoS policy from available options.
* `restricted_auth_port` - Restricted-Auth-Port.
* `rpvst_port` - Enable/disable inter-operability with rapid PVST on this interface. Valid values: `disabled`, `enabled`.

* `sample_direction` - Packet sampling direction. Valid values: `rx`, `tx`, `both`.

* `sflow_counter_interval` - sFlow sampling counter polling interval in seconds (0 - 255).
* `speed` - Switch port speed; default and available settings depend on hardware. Valid values: `auto`, `10full`, `10half`, `100full`, `100half`, `1000full`, `10000full`, `1000auto`, `40000full`, `1000fiber`, `10000`, `40000`, `auto-module`, `100FX-half`, `100FX-full`, `100000full`, `2500full`, `25000full`, `50000full`, `40000auto`, `10000cr`, `10000sr`, `100000sr4`, `100000cr4`, `25000cr4`, `25000sr4`, `5000full`, `2500auto`, `5000auto`, `1000full-fiber`, `40000sr4`, `40000cr4`, `25000cr`, `25000sr`, `50000cr`, `50000sr`.

* `speed_mask` - Switch port speed mask.
* `stacking_port` - Stacking-Port.
* `status` - Switch port admin status: up or down. Valid values: `down`, `up`.

* `sticky_mac` - Enable or disable sticky-mac on the interface. Valid values: `disable`, `enable`.

* `storm_control_policy` - Switch controller storm control policy from available options.
* `stp_bpdu_guard` - Enable/disable STP BPDU guard on this interface. Valid values: `disabled`, `enabled`.

* `stp_bpdu_guard_timeout` - BPDU Guard disabling protection (0 - 120 min).
* `stp_root_guard` - Enable/disable STP root guard on this interface. Valid values: `disabled`, `enabled`.

* `stp_state` - Enable/disable Spanning Tree Protocol (STP) on this interface. Valid values: `disabled`, `enabled`.

* `switch_id` - Switch-Id.
* `trunk_member` - Trunk member.
* `type` - Interface type: physical or trunk port. Valid values: `physical`, `trunk`.

* `untagged_vlans` - Configure switch port untagged VLANs.
* `virtual_port` - Virtualized switch port.
* `vlan` - Assign switch ports to a VLAN.

The `dhcp_snoop_option82_override` block supports:

* `circuit_id` - Circuit ID string.
* `remote_id` - Remote ID string.
* `vlan_name` - DHCP snooping option 82 VLAN.

The `remote_log` block supports:

* `csv` - Enable/disable comma-separated value (CSV) strings. Valid values: `disable`, `enable`.

* `facility` - Facility to log to remote syslog server. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.

* `name` - Remote log name.
* `port` - Remote syslog server listening port.
* `server` - IPv4 address of the remote syslog server.
* `severity` - Severity of logs to be transferred to remote log server. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.

* `status` - Enable/disable logging by FortiSwitch device to a remote syslog server. Valid values: `disable`, `enable`.


The `route_offload_router` block supports:

* `router_ip` - Router IP address.
* `vlan_name` - VLAN name.

The `snmp_community` block supports:

* `events` - SNMP notifications (traps) to send. Valid values: `cpu-high`, `mem-low`, `log-full`, `intf-ip`, `ent-conf-change`.

* `hosts` - Hosts. The structure of `hosts` block is documented below.
* `id` - SNMP community ID.
* `name` - SNMP community name.
* `query_v1_port` - SNMP v1 query port (default = 161).
* `query_v1_status` - Enable/disable SNMP v1 queries. Valid values: `disable`, `enable`.

* `query_v2c_port` - SNMP v2c query port (default = 161).
* `query_v2c_status` - Enable/disable SNMP v2c queries. Valid values: `disable`, `enable`.

* `status` - Enable/disable this SNMP community. Valid values: `disable`, `enable`.

* `trap_v1_lport` - SNMP v2c trap local port (default = 162).
* `trap_v1_rport` - SNMP v2c trap remote port (default = 162).
* `trap_v1_status` - Enable/disable SNMP v1 traps. Valid values: `disable`, `enable`.

* `trap_v2c_lport` - SNMP v2c trap local port (default = 162).
* `trap_v2c_rport` - SNMP v2c trap remote port (default = 162).
* `trap_v2c_status` - Enable/disable SNMP v2c traps. Valid values: `disable`, `enable`.


The `hosts` block supports:

* `id` - Host entry ID.
* `ip` - IPv4 address of the SNMP manager (host).

The `snmp_sysinfo` block supports:

* `contact_info` - Contact information.
* `description` - System description.
* `engine_id` - Local SNMP engine ID string (max 24 char).
* `location` - System location.
* `status` - Enable/disable SNMP. Valid values: `disable`, `enable`.


The `snmp_trap_threshold` block supports:

* `trap_high_cpu_threshold` - CPU usage when trap is sent.
* `trap_log_full_threshold` - Log disk usage when trap is sent.
* `trap_low_memory_threshold` - Memory usage when trap is sent.

The `snmp_user` block supports:

* `auth_proto` - Authentication protocol. Valid values: `md5`, `sha`, `sha1`, `sha256`, `sha384`, `sha512`, `sha224`.

* `auth_pwd` - Password for authentication protocol.
* `name` - SNMP user name.
* `priv_proto` - Privacy (encryption) protocol. Valid values: `des`, `aes`, `aes128`, `aes192`, `aes256`, `aes192c`, `aes256c`.

* `priv_pwd` - Password for privacy (encryption) protocol.
* `queries` - Enable/disable SNMP queries for this user. Valid values: `disable`, `enable`.

* `query_port` - SNMPv3 query port (default = 161).
* `security_level` - Security level for message authentication and encryption. Valid values: `no-auth-no-priv`, `auth-no-priv`, `auth-priv`.


The `static_mac` block supports:

* `description` - Description.
* `id` - ID.
* `interface` - Interface name.
* `mac` - MAC address.
* `type` - Type. Valid values: `static`, `sticky`.

* `vlan` - Vlan.

The `storm_control` block supports:

* `broadcast` - Enable/disable storm control to drop broadcast traffic. Valid values: `disable`, `enable`.

* `local_override` - Enable to override global FortiSwitch storm control settings for this FortiSwitch. Valid values: `disable`, `enable`.

* `rate` - Rate in packets per second at which storm control drops excess traffic(0-10000000, default=500, drop-all=0).
* `unknown_multicast` - Enable/disable storm control to drop unknown multicast traffic. Valid values: `disable`, `enable`.

* `unknown_unicast` - Enable/disable storm control to drop unknown unicast traffic. Valid values: `disable`, `enable`.


The `stp_instance` block supports:

* `id` - Instance ID.
* `priority` - Priority. Valid values: `0`, `4096`, `8192`, `12288`, `12328`, `16384`, `20480`, `24576`, `28672`, `32768`, `36864`, `40960`, `45056`, `49152`, `53248`, `57344`, `61440`.


The `stp_settings` block supports:

* `forward_time` - Period of time a port is in listening and learning state (4 - 30 sec, default = 15).
* `hello_time` - Period of time between successive STP frame Bridge Protocol Data Units (BPDUs) sent on a port (1 - 10 sec, default = 2).
* `local_override` - Enable to configure local STP settings that override global STP settings. Valid values: `disable`, `enable`.

* `max_age` - Maximum time before a bridge port saves its configuration BPDU information (6 - 40 sec, default = 20).
* `max_hops` - Maximum number of hops between the root bridge and the furthest bridge (1- 40, default = 20).
* `name` - Name of local STP settings configuration.
* `pending_timer` - Pending time (1 - 15 sec, default = 4).
* `revision` - STP revision number (0 - 65535).
* `status` - Enable/disable STP. Valid values: `disable`, `enable`.


The `switch_log` block supports:

* `local_override` - Enable to configure local logging settings that override global logging settings. Valid values: `disable`, `enable`.

* `severity` - Severity of FortiSwitch logs that are added to the FortiGate event log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.

* `status` - Enable/disable adding FortiSwitch logs to the FortiGate event log. Valid values: `disable`, `enable`.


The `vlan` block supports:

* `assignment_priority` - 802.1x Radius (Tunnel-Private-Group-Id) VLANID assign-by-name priority. A smaller value has a higher priority.
* `vlan_name` - VLAN name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{switch_id}}.

## Import

SwitchController ManagedSwitch can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_managedswitch.labelname {{switch_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

