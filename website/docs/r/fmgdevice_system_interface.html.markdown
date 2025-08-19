---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_interface"
description: |-
  Configure interfaces.
---

# fmgdevice_system_interface
Configure interfaces.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `client_options`: `fmgdevice_system_interface_clientoptions`
>- `dhcp_snooping_server_list`: `fmgdevice_system_interface_dhcpsnoopingserverlist`
>- `egress_queues`: `fmgdevice_system_interface_egressqueues`
>- `ipv6`: `fmgdevice_system_interface_ipv6`
>- `l2tp_client_settings`: `fmgdevice_system_interface_l2tpclientsettings`
>- `mirroring_filter`: `fmgdevice_system_interface_mirroringfilter`
>- `phy_setting`: `fmgdevice_system_interface_physetting`
>- `secondaryip`: `fmgdevice_system_interface_secondaryip`
>- `tagging`: `fmgdevice_system_interface_tagging`
>- `vrrp`: `fmgdevice_system_interface_vrrp`
>- `wifi_networks`: `fmgdevice_system_interface_wifinetworks`



## Example Usage

```hcl
resource "fmgdevice_system_interface" "trname" {
  ac_name        = "your own value"
  aggregate      = "your own value"
  aggregate_type = "physical"
  algorithm      = "L3"
  alias          = "your own value"
  name           = "your own value"
  device_name    = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `ac_name` - PPPoE server name.
* `aggregate` - Aggregate.
* `aggregate_type` - Type of aggregation. Valid values: `physical`, `vxlan`.

* `algorithm` - Frame distribution algorithm. Valid values: `L2`, `L3`, `L4`, `LB`, `Source-MAC`.

* `alias` - Alias will be displayed with the interface name to make it easier to distinguish.
* `allowaccess` - Permitted types of management access to this interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `fgfm`, `auto-ipsec`, `radius-acct`, `probe-response`, `capwap`, `dnp`, `ftm`, `fabric`, `speed-test`, `icond`.

* `annex` - Configure xDSL annex type. Valid values: `a`, `b`, `j`, `bjm`, `i`, `al`, `m`, `aijlm`, `bj`.

* `ap_discover` - Enable/disable automatic registration of unknown FortiAP devices. Valid values: `disable`, `enable`.

* `arpforward` - Enable/disable ARP forwarding. Valid values: `disable`, `enable`.

* `atm_protocol` - ATM protocol. Valid values: `none`, `ipoa`.

* `auth_cert` - HTTPS server certificate.
* `auth_portal_addr` - Address of captive portal.
* `auth_type` - PPP authentication type to use. Valid values: `auto`, `pap`, `chap`, `mschapv1`, `mschapv2`.

* `auto_auth_extension_device` - Enable/disable automatic authorization of dedicated Fortinet extension device on this interface. Valid values: `disable`, `enable`.

* `bandwidth_measure_time` - Bandwidth measure time.
* `bfd` - Bidirectional Forwarding Detection (BFD) settings. Valid values: `global`, `enable`, `disable`.

* `bfd_desired_min_tx` - BFD desired minimal transmit interval.
* `bfd_detect_mult` - BFD detection multiplier.
* `bfd_required_min_rx` - BFD required minimal receive interval.
* `broadcast_forticlient_discovery` - Enable/disable broadcasting FortiClient discovery messages. Valid values: `disable`, `enable`.

* `broadcast_forward` - Enable/disable broadcast forwarding. Valid values: `disable`, `enable`.

* `captive_portal` - Enable/disable captive portal.
* `cli_conn_status` - Cli-Conn-Status.
* `client_options` - Client-Options. The structure of `client_options` block is documented below.
* `color` - Color of icon on the GUI.
* `dedicated_to` - Configure interface for single purpose. Valid values: `none`, `management`.

* `default_purdue_level` - default purdue level of device detected on this interface. Valid values: `1`, `2`, `3`, `4`, `5`, `1.5`, `2.5`, `3.5`, `5.5`.

* `defaultgw` - Enable to get the gateway IP from the DHCP or PPPoE server. Valid values: `disable`, `enable`.

* `description` - Description.
* `detected_peer_mtu` - Detected-Peer-Mtu.
* `device_identification` - Enable/disable passively gathering of device identity information about the devices on the network connected to this interface. Valid values: `disable`, `enable`.

* `device_user_identification` - Enable/disable passive gathering of user identity information about users on this interface. Valid values: `disable`, `enable`.

* `devindex` - Devindex.
* `dhcp_broadcast_flag` - Enable/disable setting of the broadcast flag in messages sent by the DHCP client (default = enable). Valid values: `disable`, `enable`.

* `dhcp_classless_route_addition` - Enable/disable addition of classless static routes retrieved from DHCP server. Valid values: `disable`, `enable`.

* `dhcp_client_identifier` - DHCP client identifier.
* `dhcp_relay_agent_option` - Enable/disable DHCP relay agent option. Valid values: `disable`, `enable`.

* `dhcp_relay_allow_no_end_option` - Enable/disable relaying DHCP messages with no end option. Valid values: `disable`, `enable`.

* `dhcp_relay_circuit_id` - DHCP relay circuit ID.
* `dhcp_relay_interface` - Specify outgoing interface to reach server.
* `dhcp_relay_interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `dhcp_relay_ip` - DHCP relay IP address.
* `dhcp_relay_link_selection` - DHCP relay link selection.
* `dhcp_relay_request_all_server` - Enable/disable sending of DHCP requests to all servers. Valid values: `disable`, `enable`.

* `dhcp_relay_service` - Enable/disable allowing this interface to act as a DHCP relay. Valid values: `disable`, `enable`.

* `dhcp_relay_source_ip` - IP address used by the DHCP relay as its source IP.
* `dhcp_relay_type` - DHCP relay type (regular or IPsec). Valid values: `regular`, `ipsec`.

* `dhcp_relay_vrf_select` - VRF ID used for connection to server.
* `dhcp_renew_time` - DHCP renew time in seconds (300-604800), 0 means use the renew time provided by the server.
* `dhcp_smart_relay` - Enable/disable DHCP smart relay. Valid values: `disable`, `enable`.

* `dhcp_snooping_server_list` - Dhcp-Snooping-Server-List. The structure of `dhcp_snooping_server_list` block is documented below.
* `disc_retry_timeout` - Time in seconds to wait before retrying to start a PPPoE discovery, 0 means no timeout.
* `disconnect_threshold` - Time in milliseconds to wait before sending a notification that this interface is down or disconnected.
* `distance` - Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.
* `dns_server_override` - Enable/disable use DNS acquired by DHCP or PPPoE. Valid values: `disable`, `enable`.

* `dns_server_protocol` - DNS transport protocols. Valid values: `cleartext`, `dot`, `doh`.

* `drop_fragment` - Enable/disable drop fragment packets. Valid values: `disable`, `enable`.

* `drop_overlapped_fragment` - Enable/disable drop overlapped fragment packets. Valid values: `disable`, `enable`.

* `eap_ca_cert` - EAP CA certificate name.
* `eap_identity` - EAP identity.
* `eap_method` - EAP method. Valid values: `tls`, `peap`.

* `eap_password` - EAP password.
* `eap_supplicant` - Enable/disable EAP-Supplicant. Valid values: `disable`, `enable`.

* `eap_user_cert` - EAP user certificate name.
* `egress_cos` - Override outgoing CoS in user VLAN tag. Valid values: `disable`, `cos0`, `cos1`, `cos2`, `cos3`, `cos4`, `cos5`, `cos6`, `cos7`.

* `egress_queues` - Egress-Queues. The structure of `egress_queues` block is documented below.
* `egress_shaping_profile` - Outgoing traffic shaping profile.
* `eip` - Eip.
* `estimated_downstream_bandwidth` - Estimated maximum downstream bandwidth (kbps). Used to estimate link utilization.
* `estimated_upstream_bandwidth` - Estimated maximum upstream bandwidth (kbps). Used to estimate link utilization.
* `exclude_signatures` - Exclude IOT or OT application signatures. Valid values: `iot`, `ot`.

* `explicit_ftp_proxy` - Enable/disable the explicit FTP proxy on this interface. Valid values: `disable`, `enable`.

* `explicit_web_proxy` - Enable/disable the explicit web proxy on this interface. Valid values: `disable`, `enable`.

* `external` - Enable/disable identifying the interface as an external interface (which usually means it's connected to the Internet). Valid values: `disable`, `enable`.

* `fail_action_on_extender` - Action on FortiExtender when interface fail. Valid values: `soft-restart`, `hard-restart`, `reboot`.

* `fail_alert_interfaces` - Names of the FortiGate interfaces to which the link failure alert is sent.
* `fail_alert_method` - Select link-failed-signal or link-down method to alert about a failed link. Valid values: `link-failed-signal`, `link-down`.

* `fail_detect` - Enable/disable fail detection features for this interface. Valid values: `disable`, `enable`.

* `fail_detect_option` - Options for detecting that this interface has failed. Valid values: `detectserver`, `link-down`.

* `fortilink` - Enable FortiLink to dedicate this interface to manage other Fortinet devices. Valid values: `disable`, `enable`.

* `fortilink_backup_link` - Fortilink-Backup-Link.
* `fortilink_neighbor_detect` - Protocol for FortiGate neighbor discovery. Valid values: `lldp`, `fortilink`.

* `fortilink_split_interface` - Enable/disable FortiLink split interface to connect member link to different FortiSwitch in stack for uplink redundancy. Valid values: `disable`, `enable`.

* `fortilink_stacking` - Enable/disable FortiLink switch-stacking on this interface. Valid values: `disable`, `enable`.

* `forward_domain` - Transparent mode forward domain.
* `forward_error_correction` - Configure forward error correction (FEC). Valid values: `disable`, `enable`, `rs-fec`, `base-r-fec`, `fec-cl91`, `fec-cl74`, `rs-544`, `none`, `cl91-rs-fec`, `cl74-fc-fec`, `auto`.

* `gateway_address` - Gateway address.
* `gi_gk` - Enable/disable Gi Gatekeeper. Valid values: `disable`, `enable`.

* `icmp_accept_redirect` - Enable/disable ICMP accept redirect. Valid values: `disable`, `enable`.

* `icmp_send_redirect` - Enable/disable sending of ICMP redirects. Valid values: `disable`, `enable`.

* `ident_accept` - Enable/disable authentication for this interface. Valid values: `disable`, `enable`.

* `idle_timeout` - PPPoE auto disconnect after idle timeout seconds, 0 means no timeout.
* `ike_saml_server` - Configure IKE authentication SAML server.
* `inbandwidth` - Bandwidth limit for incoming traffic (0 - 80000000 kbps), 0 means unlimited.
* `ingress_cos` - Override incoming CoS in user VLAN tag on VLAN interface or assign a priority VLAN tag on physical interface. Valid values: `disable`, `cos0`, `cos1`, `cos2`, `cos3`, `cos4`, `cos5`, `cos6`, `cos7`.

* `ingress_shaping_profile` - Incoming traffic shaping profile.
* `ingress_spillover_threshold` - Ingress Spillover threshold (0 - 16776000 kbps), 0 means unlimited.
* `interconnect_profile` - Set interconnect profile. Valid values: `default`, `profile1`, `profile2`.

* `interface` - Interface name.
* `internal` - Implicitly created.
* `ip` - Interface IPv4 address and subnet mask, syntax: X.X.X.X/24.
* `ip_managed_by_fortiipam` - Enable/disable automatic IP address assignment of this interface by FortiIPAM. Valid values: `disable`, `enable`, `inherit-global`.

* `ipmac` - Enable/disable IP/MAC binding. Valid values: `disable`, `enable`.

* `ips_sniffer_mode` - Enable/disable the use of this interface as a one-armed sniffer. Valid values: `disable`, `enable`.

* `ipunnumbered` - Unnumbered IP used for PPPoE interfaces for which no unique local address is provided.
* `ipv6` - Ipv6. The structure of `ipv6` block is documented below.
* `l2forward` - Enable/disable l2 forwarding. Valid values: `disable`, `enable`.

* `l2tp_client` - Enable/disable this interface as a Layer 2 Tunnelling Protocol (L2TP) client. Valid values: `disable`, `enable`.

* `l2tp_client_settings` - L2Tp-Client-Settings. The structure of `l2tp_client_settings` block is documented below.
* `lacp_ha_secondary` - LACP HA secondary member. Valid values: `disable`, `enable`.

* `lacp_ha_slave` - Lacp-Ha-Slave. Valid values: `disable`, `enable`.

* `lacp_mode` - LACP mode. Valid values: `static`, `passive`, `active`.

* `lacp_speed` - How often the interface sends LACP messages. Valid values: `slow`, `fast`.

* `lcp_echo_interval` - Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
* `lcp_max_echo_fails` - Maximum missed LCP echo messages before disconnect.
* `link_up_delay` - Number of milliseconds to wait before considering a link is up.
* `lldp_network_policy` - LLDP-MED network policy profile.
* `lldp_reception` - Enable/disable Link Layer Discovery Protocol (LLDP) reception. Valid values: `disable`, `enable`, `vdom`.

* `lldp_transmission` - Enable/disable Link Layer Discovery Protocol (LLDP) transmission. Valid values: `enable`, `disable`, `vdom`.

* `macaddr` - Change the interface's MAC address.
* `managed_subnetwork_size` - Number of IP addresses to be allocated by FortiIPAM and used by this FortiGate unit's DHCP server settings. Valid values: `32`, `64`, `128`, `256`, `512`, `1024`, `2048`, `4096`, `8192`, `16384`, `32768`, `65536`.

* `management_ip` - High Availability in-band management IP address of this interface.
* `max_egress_burst_rate` - Max egress burst rate (kbits per sec).
* `max_egress_rate` - Max egress rate (kbits per sec).
* `measured_downstream_bandwidth` - Measured downstream bandwidth (kbps).
* `measured_upstream_bandwidth` - Measured upstream bandwidth (kbps).
* `mediatype` - Select SFP media interface type Valid values: `serdes-sfp`, `sgmii-sfp`, `cfp2-sr10`, `cfp2-lr4`, `serdes-copper-sfp`, `sr`, `cr`, `lr`, `qsfp28-sr4`, `qsfp28-lr4`, `qsfp28-cr4`, `sr4`, `cr4`, `lr4`, `none`, `gmii`, `sgmii`, `sr2`, `lr2`, `cr2`, `sr8`, `lr8`, `cr8`.

* `member` - Physical interfaces that belong to the aggregate or redundant interface.
* `min_links` - Minimum number of aggregated ports that must be up.
* `min_links_down` - Action to take when less than the configured minimum number of links are active. Valid values: `operational`, `administrative`.

* `mirroring_direction` - Port mirroring direction. Valid values: `rx`, `tx`, `both`.

* `mirroring_filter` - Mirroring-Filter. The structure of `mirroring_filter` block is documented below.
* `mirroring_port` - Mirroring port.
* `mode` - Addressing mode (static, DHCP, PPPoE). Valid values: `static`, `dhcp`, `pppoe`, `pppoa`, `ipoa`, `eoa`.

* `monitor_bandwidth` - Enable monitoring bandwidth on this interface. Valid values: `disable`, `enable`.

* `mtu` - MTU value for this interface.
* `mtu_override` - Enable to set a custom MTU for this interface. Valid values: `disable`, `enable`.

* `mux_type` - Multiplexer type. Valid values: `llc-encaps`, `vc-encaps`.

* `name` - Name.
* `ndiscforward` - Enable/disable NDISC forwarding. Valid values: `disable`, `enable`.

* `netbios_forward` - Enable/disable NETBIOS forwarding. Valid values: `disable`, `enable`.

* `netflow_sample_rate` - NetFlow sample rate.  Sample one packet every configured number of packets (1 - 65535, default = 1, which means standard NetFlow where all packets are sampled).
* `netflow_sampler` - Enable/disable NetFlow on this interface and set the data that NetFlow collects (rx, tx, or both). Valid values: `disable`, `tx`, `rx`, `both`.

* `netflow_sampler_id` - Netflow sampler ID.
* `np_qos_profile` - NP QoS profile ID.
* `outbandwidth` - Bandwidth limit for outgoing traffic (0 - 80000000 kbps).
* `padt_retry_timeout` - PPPoE Active Discovery Terminate (PADT) used to terminate sessions after an idle time.
* `password` - PPPoE account's password.
* `phy_mode` - DSL physical mode. Valid values: `auto`, `adsl`, `vdsl`, `adsl-auto`, `vdsl2`, `adsl2+`, `adsl2`, `g.dmt`, `g-dmt`, `t1.413`, `t1-413`, `g.lite`, `g-lite`.

* `phy_setting` - Phy-Setting. The structure of `phy_setting` block is documented below.
* `poe` - Enable/disable PoE status. Valid values: `disable`, `enable`.

* `polling_interval` - sFlow polling interval in seconds (1 - 255).
* `port_mirroring` - Enable/disable NP port mirroring. Valid values: `disable`, `enable`.

* `pppoe_egress_cos` - CoS in VLAN tag for outgoing PPPoE/PPP packets. Valid values: `cos0`, `cos1`, `cos2`, `cos3`, `cos4`, `cos5`, `cos6`, `cos7`.

* `pppoe_unnumbered_negotiate` - Enable/disable PPPoE unnumbered negotiation. Valid values: `disable`, `enable`.

* `pptp_auth_type` - PPTP authentication type. Valid values: `auto`, `pap`, `chap`, `mschapv1`, `mschapv2`.

* `pptp_client` - Enable/disable PPTP client. Valid values: `disable`, `enable`.

* `pptp_password` - PPTP password.
* `pptp_server_ip` - PPTP server IP address.
* `pptp_timeout` - Idle timer in minutes (0 for disabled).
* `pptp_user` - PPTP user name.
* `preserve_session_route` - Enable/disable preservation of session route when dirty. Valid values: `disable`, `enable`.

* `priority` - Priority of learned routes.
* `priority_override` - Enable/disable fail back to higher priority port once recovered. Valid values: `disable`, `enable`.

* `proxy_captive_portal` - Enable/disable proxy captive portal on this interface. Valid values: `disable`, `enable`.

* `pvc_atm_qos` - SFP-DSL ADSL Fallback PVC ATM QoS. Valid values: `cbr`, `rt-vbr`, `nrt-vbr`, `ubr`.

* `pvc_chan` - SFP-DSL ADSL Fallback PVC Channel.
* `pvc_crc` - SFP-DSL ADSL Fallback PVC CRC Option:  bit0: sar LLC preserve, bit1: ream LLC preserve, bit2: ream VC-MUX has crc.
* `pvc_pcr` - SFP-DSL ADSL Fallback PVC Packet Cell Rate in cells (0 - 5500).
* `pvc_scr` - SFP-DSL ADSL Fallback PVC Sustainable Cell Rate in cells (0 - 5500).
* `pvc_vlan_id` - SFP-DSL ADSL Fallback PVC VLAN ID.
* `pvc_vlan_rx_id` - SFP-DSL ADSL Fallback PVC VLANID RX.
* `pvc_vlan_rx_op` - SFP-DSL ADSL Fallback PVC VLAN RX op. Valid values: `pass-through`, `replace`, `remove`.

* `pvc_vlan_tx_id` - SFP-DSL ADSL Fallback PVC VLAN ID TX.
* `pvc_vlan_tx_op` - SFP-DSL ADSL Fallback PVC VLAN TX op. Valid values: `pass-through`, `replace`, `remove`.

* `reachable_time` - IPv4 reachable time in milliseconds (30000 - 3600000, default = 30000).
* `redundant_interface` - Redundant-Interface.
* `remote_ip` - Remote IP address of tunnel.
* `replacemsg_override_group` - Replacement message override group.
* `retransmission` - Enable/disable DSL retransmission. Valid values: `disable`, `enable`.

* `ring_rx` - RX ring size.
* `ring_tx` - TX ring size.
* `role` - Interface role. Valid values: `lan`, `wan`, `dmz`, `undefined`.

* `sample_direction` - Data that NetFlow collects (rx, tx, or both). Valid values: `rx`, `tx`, `both`.

* `sample_rate` - sFlow sample rate (10 - 99999).
* `secondary_ip` - Enable/disable adding a secondary IP to this interface. Valid values: `disable`, `enable`.

* `secondaryip` - Secondaryip. The structure of `secondaryip` block is documented below.
* `security_8021x_dynamic_vlan_id` - VLAN ID for virtual switch.
* `security_8021x_master` - 802.1X master virtual-switch.
* `security_8021x_member_mode` - 802.1X member mode. Valid values: `disable`, `switch`.

* `security_8021x_mode` - 802.1X mode. Valid values: `default`, `dynamic-vlan`, `fallback`, `slave`.

* `security_exempt_list` - Name of security-exempt-list.
* `security_external_logout` - URL of external authentication logout server.
* `security_external_web` - URL of external authentication web server.
* `security_groups` - User groups that can authenticate with the captive portal.
* `security_ip_auth_bypass` - Enable/disable IP authentication bypass. Valid values: `disable`, `enable`.

* `security_mac_auth_bypass` - Enable/disable MAC authentication bypass. Valid values: `disable`, `enable`, `mac-auth-only`.

* `security_mode` - Turn on captive portal authentication for this interface. Valid values: `none`, `captive-portal`, `802.1X`.

* `security_redirect_url` - URL redirection after disclaimer/authentication.
* `select_profile_30a_35b` - Select VDSL Profile 30a or 35b. Valid values: `30A`, `35B`.

* `service_name` - PPPoE service name.
* `sflow_sampler` - Enable/disable sFlow on this interface. Valid values: `disable`, `enable`.

* `sfp_dsl` - Enable/disable SFP DSL. Valid values: `disable`, `enable`.

* `sfp_dsl_adsl_fallback` - Enable/disable SFP DSL ADSL fallback. Valid values: `disable`, `enable`.

* `sfp_dsl_autodetect` - Enable/disable SFP DSL MAC address autodetect. Valid values: `disable`, `enable`.

* `sfp_dsl_mac` - SFP DSL MAC address.
* `snmp_index` - Permanent SNMP Index of the interface.
* `speed` - Interface speed. The default setting and the options available depend on the interface hardware. Valid values: `auto`, `10full`, `10half`, `100full`, `100half`, `1000full`, `1000half`, `10000full`, `1000auto`, `10000auto`, `40000full`, `100Gfull`, `25000full`, `50000full`, `40000auto`, `25000auto`, `100Gauto`, `2500auto`, `400Gfull`, `400Gauto`, `5000auto`, `50000auto`, `200Gfull`, `200Gauto`, `100auto`.

* `spillover_threshold` - Egress Spillover threshold (0 - 16776000 kbps), 0 means unlimited.
* `src_check` - Enable/disable source IP check. Valid values: `disable`, `enable`.

* `status` - Bring the interface up or shut the interface down. Valid values: `down`, `up`.

* `stp` - Enable/disable STP. Valid values: `disable`, `enable`.

* `stp_ha_slave` - Control STP behaviour on HA slave. Valid values: `disable`, `enable`, `priority-adjust`.

* `stp_edge` - Enable/disable as STP edge port. Valid values: `disable`, `enable`.

* `stp_ha_secondary` - Control STP behavior on HA secondary. Valid values: `disable`, `enable`, `priority-adjust`.

* `stpforward` - Enable/disable STP forwarding. Valid values: `disable`, `enable`.

* `stpforward_mode` - Configure STP forwarding mode. Valid values: `rpl-all-ext-id`, `rpl-bridge-ext-id`, `rpl-nothing`.

* `subst` - Enable to always send packets from this interface to a destination MAC address. Valid values: `disable`, `enable`.

* `substitute_dst_mac` - Destination MAC address that all packets are sent to from this interface.
* `sw_algorithm` - Frame distribution algorithm for switch. Valid values: `l2`, `l3`, `eh`, `default`.

* `swc_first_create` - Initial create for switch-controller VLANs.
* `swc_vlan` - Swc-Vlan.
* `switch` - Switch.
* `switch_controller_access_vlan` - Block FortiSwitch port-to-port traffic. Valid values: `disable`, `enable`.

* `switch_controller_arp_inspection` - Enable/disable/Monitor FortiSwitch ARP inspection. Valid values: `disable`, `enable`, `monitor`.

* `switch_controller_dhcp_snooping` - Switch controller DHCP snooping. Valid values: `disable`, `enable`.

* `switch_controller_dhcp_snooping_option82` - Switch controller DHCP snooping option82. Valid values: `disable`, `enable`.

* `switch_controller_dhcp_snooping_verify_mac` - Switch controller DHCP snooping verify MAC. Valid values: `disable`, `enable`.

* `switch_controller_dynamic` - Integrated FortiLink settings for managed FortiSwitch.
* `switch_controller_feature` - Interface's purpose when assigning traffic (read only). Valid values: `none`, `default-vlan`, `quarantine`, `sniffer`, `voice`, `camera`, `rspan`, `video`, `nac`, `nac-segment`.

* `switch_controller_igmp_snooping` - Switch controller IGMP snooping. Valid values: `disable`, `enable`.

* `switch_controller_igmp_snooping_fast_leave` - Switch controller IGMP snooping fast-leave. Valid values: `disable`, `enable`.

* `switch_controller_igmp_snooping_proxy` - Switch controller IGMP snooping proxy. Valid values: `disable`, `enable`.

* `switch_controller_iot_scanning` - Enable/disable managed FortiSwitch IoT scanning. Valid values: `disable`, `enable`.

* `switch_controller_learning_limit` - Limit the number of dynamic MAC addresses on this VLAN (1 - 128, 0 = no limit, default).
* `switch_controller_mgmt_vlan` - VLAN to use for FortiLink management purposes.
* `switch_controller_nac` - Integrated FortiLink settings for managed FortiSwitch.
* `switch_controller_netflow_collect` - NetFlow collection and processing. Valid values: `disable`, `enable`.

* `switch_controller_offload` - Enable/disable managed FortiSwitch routing offload. Valid values: `disable`, `enable`.

* `switch_controller_offload_gw` - Enable/disable managed FortiSwitch routing offload gateway. Valid values: `disable`, `enable`.

* `switch_controller_offload_ip` - IP for routing offload on FortiSwitch.
* `switch_controller_rspan_mode` - Stop Layer2 MAC learning and interception of BPDUs and other packets on this interface. Valid values: `disable`, `enable`.

* `switch_controller_source_ip` - Source IP address used in FortiLink over L3 connections. Valid values: `outbound`, `fixed`.

* `switch_controller_traffic_policy` - Switch controller traffic policy for the VLAN.
* `system_id` - Define a system ID for the aggregate interface.
* `system_id_type` - Method in which system ID is generated. Valid values: `auto`, `user`.

* `tagging` - Tagging. The structure of `tagging` block is documented below.
* `tc_mode` - DSL transfer mode. Valid values: `ptm`, `atm`.

* `tcp_mss` - TCP maximum segment size. 0 means do not change segment size.
* `trunk` - Enable/disable VLAN trunk. Valid values: `disable`, `enable`.

* `trust_ip_1` - Trusted host for dedicated management traffic (0.0.0.0/24 for all hosts).
* `trust_ip_2` - Trusted host for dedicated management traffic (0.0.0.0/24 for all hosts).
* `trust_ip_3` - Trusted host for dedicated management traffic (0.0.0.0/24 for all hosts).
* `trust_ip6_1` - Trusted IPv6 host for dedicated management traffic (::/0 for all hosts).
* `trust_ip6_2` - Trusted IPv6 host for dedicated management traffic (::/0 for all hosts).
* `trust_ip6_3` - Trusted IPv6 host for dedicated management traffic (::/0 for all hosts).
* `type` - Interface type. Valid values: `physical`, `vlan`, `aggregate`, `redundant`, `tunnel`, `wireless`, `vdom-link`, `loopback`, `switch`, `hard-switch`, `hdlc`, `vap-switch`, `wl-mesh`, `fortilink`, `switch-vlan`, `fctrl-trunk`, `tdm`, `fext-wan`, `vxlan`, `emac-vlan`, `geneve`, `ssl`, `lan-extension`.

* `username` - Username of the PPPoE account, provided by your ISP.
* `vci` - Virtual Channel ID.
* `vdom` - Interface is in this virtual domain (VDOM).
* `vectoring` - Enable/disable DSL vectoring. Valid values: `disable`, `enable`.

* `vindex` - Vindex.
* `virtual_mac` - Change the interface's virtual MAC address.
* `vlan_id` - Vlan ID.
* `vlan_op_mode` - Configure DSL 802.1q mode. Valid values: `tag`, `untag`, `passthrough`.

* `vlan_protocol` - Ethernet protocol of VLAN. Valid values: `8021q`, `8021ad`.

* `vlanforward` - Enable/disable traffic forwarding between VLANs on this interface. Valid values: `disable`, `enable`.

* `vlanid` - VLAN ID (1 - 4094).
* `vpi` - Virtual Path ID.
* `vrf` - Virtual Routing Forwarding ID.
* `vrrp` - Vrrp. The structure of `vrrp` block is documented below.
* `vrrp_virtual_mac` - Enable/disable use of virtual MAC for VRRP. Valid values: `disable`, `enable`.

* `wccp` - Enable/disable WCCP on this interface. Used for encapsulated WCCP communication between WCCP clients and servers. Valid values: `disable`, `enable`.

* `weight` - Default weight for static routes (if route has no weight configured).
* `wifi_5g_threshold` - Minimal signal strength to be considered as a good 5G AP.
* `wifi_acl` - Access control for MAC addresses in the MAC list. Valid values: `deny`, `allow`.

* `wifi_ap_band` - How to select the AP to connect. Valid values: `any`, `5g-preferred`, `5g-only`.

* `wifi_auth` - WiFi authentication. Valid values: `PSK`, `radius`, `usergroup`.

* `wifi_auto_connect` - Enable/disable WiFi network auto connect. Valid values: `disable`, `enable`.

* `wifi_auto_save` - Enable/disable WiFi network automatic save. Valid values: `disable`, `enable`.

* `wifi_broadcast_ssid` - Enable/disable SSID broadcast in the beacon. Valid values: `disable`, `enable`.

* `wifi_dns_server1` - DNS server 1.
* `wifi_dns_server2` - DNS server 2.
* `wifi_encrypt` - Data encryption. Valid values: `TKIP`, `AES`.

* `wifi_fragment_threshold` - WiFi fragment threshold (800 - 2346).
* `wifi_gateway` - IPv4 default gateway IP address.
* `wifi_key` - WiFi WEP Key.
* `wifi_keyindex` - WEP key index (1 - 4).
* `wifi_mac_filter` - Enable/disable MAC filter status. Valid values: `disable`, `enable`.

* `wifi_networks` - Wifi-Networks. The structure of `wifi_networks` block is documented below.
* `wifi_passphrase` - WiFi pre-shared key for WPA.
* `wifi_radius_server` - WiFi RADIUS server for WPA.
* `wifi_rts_threshold` - WiFi RTS threshold (256 - 2346).
* `wifi_security` - Wireless access security of SSID. Valid values: `None`, `wep64`, `wep128`, `WPA_PSK`, `WPA_RADIUS`, `WPA`, `WPA2`, `WPA2_AUTO`, `open`, `wpa-personal`, `wpa-enterprise`, `wpa-only-personal`, `wpa-only-enterprise`, `wpa2-only-personal`, `wpa2-only-enterprise`.

* `wifi_ssid` - IEEE 802.11 Service Set Identifier.
* `wifi_usergroup` - WiFi user group for WPA.
* `wins_ip` - WINS server IP.
* `autogenerated` - Indicates whether the interface is automatically created by FortiManager, for example, created during the VPN creation process. If it is, set it to "auto", else keep it empty.
* `type` - Indicates whether the type of the interface is physical. If it is, set it to "physical", else keep it empty.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `client_options` block supports:

* `code` - DHCP client option code.
* `id` - ID.
* `ip` - DHCP option IPs.
* `type` - DHCP client option type. Valid values: `hex`, `string`, `ip`, `fqdn`.

* `value` - DHCP client option value.

The `dhcp_snooping_server_list` block supports:

* `name` - DHCP server name.
* `server_ip` - IP address for DHCP server.

The `egress_queues` block supports:

* `cos0` - CoS profile name for CoS 0.
* `cos1` - CoS profile name for CoS 1.
* `cos2` - CoS profile name for CoS 2.
* `cos3` - CoS profile name for CoS 3.
* `cos4` - CoS profile name for CoS 4.
* `cos5` - CoS profile name for CoS 5.
* `cos6` - CoS profile name for CoS 6.
* `cos7` - CoS profile name for CoS 7.

The `ipv6` block supports:

* `autoconf` - Enable/disable address auto config. Valid values: `disable`, `enable`.

* `cli_conn6_status` - Cli-Conn6-Status.
* `client_options` - Client-Options. The structure of `client_options` block is documented below.
* `dhcp6_client_options` - Dhcp6-Client-Options. Valid values: `rapid`, `iapd`, `iana`, `dns`, `dnsname`.

* `dhcp6_iapd_list` - Dhcp6-Iapd-List. The structure of `dhcp6_iapd_list` block is documented below.
* `dhcp6_information_request` - Enable/disable DHCPv6 information request. Valid values: `disable`, `enable`.

* `dhcp6_prefix_delegation` - Enable/disable DHCPv6 prefix delegation. Valid values: `disable`, `enable`.

* `dhcp6_prefix_hint` - DHCPv6 prefix that will be used as a hint to the upstream DHCPv6 server.
* `dhcp6_prefix_hint_plt` - DHCPv6 prefix hint preferred life time (sec), 0 means unlimited lease time.
* `dhcp6_prefix_hint_vlt` - DHCPv6 prefix hint valid life time (sec).
* `dhcp6_relay_interface_id` - DHCP6 relay interface ID.
* `dhcp6_relay_ip` - DHCPv6 relay IP address.
* `dhcp6_relay_service` - Enable/disable DHCPv6 relay. Valid values: `disable`, `enable`.

* `dhcp6_relay_source_interface` - Enable/disable use of address on this interface as the source address of the relay message. Valid values: `disable`, `enable`.

* `dhcp6_relay_source_ip` - IPv6 address used by the DHCP6 relay as its source IP.
* `dhcp6_relay_type` - DHCPv6 relay type. Valid values: `regular`.

* `icmp6_send_redirect` - Enable/disable sending of ICMPv6 redirects. Valid values: `disable`, `enable`.

* `interface_identifier` - IPv6 interface identifier.
* `ip6_address` - Primary IPv6 address prefix. Syntax: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx.
* `ip6_adv_rio` - Enable/disable sending advertisements with route information option. Valid values: `disable`, `enable`.

* `ip6_allowaccess` - Allow management access to the interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `fgfm`, `capwap`, `fabric`.

* `ip6_default_life` - Default life (sec).
* `ip6_delegated_prefix_iaid` - IAID of obtained delegated-prefix from the upstream interface.
* `ip6_delegated_prefix_list` - Ip6-Delegated-Prefix-List. The structure of `ip6_delegated_prefix_list` block is documented below.
* `ip6_dns_server_override` - Enable/disable using the DNS server acquired by DHCP. Valid values: `disable`, `enable`.

* `ip6_dnssl_list` - Ip6-Dnssl-List. The structure of `ip6_dnssl_list` block is documented below.
* `ip6_extra_addr` - Ip6-Extra-Addr. The structure of `ip6_extra_addr` block is documented below.
* `ip6_hop_limit` - Hop limit (0 means unspecified).
* `ip6_link_mtu` - IPv6 link MTU.
* `ip6_manage_flag` - Enable/disable the managed flag. Valid values: `disable`, `enable`.

* `ip6_max_interval` - IPv6 maximum interval (4 to 1800 sec).
* `ip6_min_interval` - IPv6 minimum interval (3 to 1350 sec).
* `ip6_mode` - Addressing mode (static, DHCP, delegated). Valid values: `static`, `dhcp`, `pppoe`, `delegated`.

* `ip6_other_flag` - Enable/disable the other IPv6 flag. Valid values: `disable`, `enable`.

* `ip6_prefix_list` - Ip6-Prefix-List. The structure of `ip6_prefix_list` block is documented below.
* `ip6_prefix_mode` - Assigning a prefix from DHCP or RA. Valid values: `dhcp6`, `ra`.

* `ip6_rdnss_list` - Ip6-Rdnss-List. The structure of `ip6_rdnss_list` block is documented below.
* `ip6_reachable_time` - IPv6 reachable time (milliseconds; 0 means unspecified).
* `ip6_retrans_time` - IPv6 retransmit time (milliseconds; 0 means unspecified).
* `ip6_route_list` - Ip6-Route-List. The structure of `ip6_route_list` block is documented below.
* `ip6_route_pref` - Set route preference to the interface (default = medium). Valid values: `medium`, `high`, `low`.

* `ip6_send_adv` - Enable/disable sending advertisements about the interface. Valid values: `disable`, `enable`.

* `ip6_subnet` - Subnet to routing prefix. Syntax: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx.
* `ip6_upstream_interface` - Interface name providing delegated information.
* `nd_cert` - Neighbor discovery certificate.
* `nd_cga_modifier` - Neighbor discovery CGA modifier.
* `nd_mode` - Neighbor discovery mode. Valid values: `basic`, `SEND-compatible`.

* `nd_security_level` - Neighbor discovery security level (0 - 7; 0 = least secure, default = 0).
* `nd_timestamp_delta` - Neighbor discovery timestamp delta value (1 - 3600 sec; default = 300).
* `nd_timestamp_fuzz` - Neighbor discovery timestamp fuzz factor (1 - 60 sec; default = 1).
* `ra_send_mtu` - Enable/disable sending link MTU in RA packet. Valid values: `disable`, `enable`.

* `unique_autoconf_addr` - Enable/disable unique auto config address. Valid values: `disable`, `enable`.

* `vrip6_link_local` - Link-local IPv6 address of virtual router.
* `vrrp_virtual_mac6` - Enable/disable virtual MAC for VRRP. Valid values: `disable`, `enable`.

* `vrrp6` - Vrrp6. The structure of `vrrp6` block is documented below.

The `client_options` block supports:

* `code` - DHCPv6 option code.
* `id` - ID.
* `ip6` - DHCP option IP6s.
* `type` - DHCPv6 option type. Valid values: `hex`, `string`, `ip6`, `fqdn`.

* `value` - DHCPv6 option value (hexadecimal value must be even).

The `dhcp6_iapd_list` block supports:

* `iaid` - Identity association identifier.
* `prefix_hint` - DHCPv6 prefix that will be used as a hint to the upstream DHCPv6 server.
* `prefix_hint_plt` - DHCPv6 prefix hint preferred life time (sec), 0 means unlimited lease time.
* `prefix_hint_vlt` - DHCPv6 prefix hint valid life time (sec).

The `ip6_delegated_prefix_list` block supports:

* `autonomous_flag` - Enable/disable the autonomous flag. Valid values: `disable`, `enable`.

* `delegated_prefix_iaid` - IAID of obtained delegated-prefix from the upstream interface.
* `onlink_flag` - Enable/disable the onlink flag. Valid values: `disable`, `enable`.

* `prefix_id` - Prefix ID.
* `rdnss` - Recursive DNS server option.
* `rdnss_service` - Recursive DNS service option. Valid values: `delegated`, `default`, `specify`.

* `subnet` - Add subnet ID to routing prefix.
* `upstream_interface` - Name of the interface that provides delegated information.

The `ip6_dnssl_list` block supports:

* `dnssl_life_time` - DNS search list time in seconds (0 - 4294967295, default = 1800).
* `domain` - Domain name.

The `ip6_extra_addr` block supports:

* `prefix` - IPv6 address prefix.

The `ip6_prefix_list` block supports:

* `autonomous_flag` - Enable/disable the autonomous flag. Valid values: `disable`, `enable`.

* `dnssl` - DNS search list option.
* `onlink_flag` - Enable/disable the onlink flag. Valid values: `disable`, `enable`.

* `preferred_life_time` - Preferred life time (sec).
* `prefix` - IPv6 prefix.
* `rdnss` - Recursive DNS server option.
* `valid_life_time` - Valid life time (sec).

The `ip6_rdnss_list` block supports:

* `rdnss` - Recursive DNS server option.
* `rdnss_life_time` - Recursive DNS server life time in seconds (0 - 4294967295, default = 1800).

The `ip6_route_list` block supports:

* `route` - IPv6 route.
* `route_life_time` - Route life time in seconds (0 - 65535, default = 1800).
* `route_pref` - Set route preference to the interface (default = medium). Valid values: `medium`, `high`, `low`.


The `vrrp6` block supports:

* `accept_mode` - Enable/disable accept mode. Valid values: `disable`, `enable`.

* `adv_interval` - Advertisement interval (1 - 255 seconds).
* `ignore_default_route` - Enable/disable ignoring of default route when checking destination. Valid values: `disable`, `enable`.

* `preempt` - Enable/disable preempt mode. Valid values: `disable`, `enable`.

* `priority` - Priority of the virtual router (1 - 255).
* `start_time` - Startup time (1 - 255 seconds).
* `status` - Enable/disable VRRP. Valid values: `disable`, `enable`.

* `vrdst_priority` - Priority of the virtual router when the virtual router destination becomes unreachable (0 - 254).
* `vrdst6` - Monitor the route to this destination.
* `vrgrp` - VRRP group ID (1 - 65535).
* `vrid` - Virtual router identifier (1 - 255).
* `vrip6` - IPv6 address of the virtual router.

The `l2tp_client_settings` block supports:

* `auth_type` - L2TP authentication type. Valid values: `auto`, `pap`, `chap`, `mschapv1`, `mschapv2`.

* `defaultgw` - Enable/disable default gateway. Valid values: `disable`, `enable`.

* `distance` - Distance of learned routes.
* `hello_interval` - L2TP hello message interval in seconds (0 - 3600 sec, default = 60).
* `ip` - Ip.
* `mtu` - L2TP MTU.
* `password` - L2TP password.
* `peer_host` - L2TP peer host address.
* `peer_mask` - L2TP peer mask.
* `peer_port` - L2TP peer port number.
* `priority` - Priority of learned routes.
* `user` - L2TP user name.

The `mirroring_filter` block supports:

* `dbg_dump` - Debug mirroring filter, driver dump data/mask pdq.
* `filter_dport` - Destinatin port of mirroring filter.
* `filter_dstip` - Destinatin IP and mask of mirroring filter.
* `filter_protocol` - Protocol of mirroring filter.
* `filter_sport` - Source port of mirroring filter.
* `filter_srcip` - Source IP and mask of mirroring filter.

The `phy_setting` block supports:

* `signal_ok_threshold_value` - Signal-ok-threshold value(0 - 12).

The `secondaryip` block supports:

* `allowaccess` - Management access settings for the secondary IP address. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `fgfm`, `auto-ipsec`, `radius-acct`, `probe-response`, `capwap`, `dnp`, `ftm`, `fabric`, `speed-test`, `icond`.

* `id` - ID.
* `ip` - Secondary IP address of the interface.
* `secip_relay_ip` - DHCP relay IP address.

The `tagging` block supports:

* `category` - Tag category.
* `name` - Tagging entry name.
* `tags` - Tags.

The `vrrp` block supports:

* `accept_mode` - Enable/disable accept mode. Valid values: `disable`, `enable`.

* `adv_interval` - Advertisement interval (1 - 255 seconds).
* `ignore_default_route` - Enable/disable ignoring of default route when checking destination. Valid values: `disable`, `enable`.

* `preempt` - Enable/disable preempt mode. Valid values: `disable`, `enable`.

* `priority` - Priority of the virtual router (1 - 255).
* `proxy_arp` - Proxy-Arp. The structure of `proxy_arp` block is documented below.
* `start_time` - Startup time (1 - 255 seconds).
* `status` - Enable/disable this VRRP configuration. Valid values: `disable`, `enable`.

* `version` - VRRP version. Valid values: `2`, `3`.

* `vrdst` - Monitor the route to this destination.
* `vrdst_priority` - Priority of the virtual router when the virtual router destination becomes unreachable (0 - 254).
* `vrgrp` - VRRP group ID (1 - 65535).
* `vrid` - Virtual router identifier (1 - 255).
* `vrip` - IP address of the virtual router.

The `proxy_arp` block supports:

* `id` - ID.
* `ip` - Set IP addresses of proxy ARP.

The `wifi_networks` block supports:

* `id` - ID.
* `obsolete_security_options` - Enable/disable obsolete security options. Valid values: `disable`, `enable`.

* `wifi_ca_certificate` - CA certificate for WPA2/WPA3-ENTERPRISE.
* `wifi_client_certificate` - Client certificate for WPA2/WPA3-ENTERPRISE.
* `wifi_eap_type` - WPA2/WPA3-ENTERPRISE EAP Method. Valid values: `tls`, `peap`, `both`.

* `wifi_encrypt` - Data encryption. Valid values: `TKIP`, `AES`.

* `wifi_key` - WiFi WEP Key.
* `wifi_keyindex` - WEP key index (1 - 4).
* `wifi_passphrase` - WiFi pre-shared key for WPA-PSK or password for WPA3-SAE and WPA2/WPA3-ENTERPRISE.
* `wifi_private_key` - Private key for WPA2/WPA3-ENTERPRISE.
* `wifi_private_key_password` - Password for private key file for WPA2/WPA3-ENTERPRISE.
* `wifi_security` - Wireless access security of SSID. Valid values: `wep64`, `wep128`, `open`, `wpa-personal`, `wpa-enterprise`, `wpa-only-personal`, `wpa2-only-personal`, `owe`, `wpa3-sae`.

* `wifi_ssid` - IEEE 802.11 Service Set Identifier.
* `wifi_username` - Username for WPA2/WPA3-ENTERPRISE.

-> If you want to modify the configuration of the interface which is automatically created by FortiManager (for example, the interface created during the VPN creation process), please set the `autogenerated` argument to "auto".


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Interface can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_interface.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

