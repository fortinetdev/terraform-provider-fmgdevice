---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_settings"
description: |-
  Configure VDOM settings.
---

# fmgdevice_system_settings
Configure VDOM settings.

## Example Usage

```hcl
resource "fmgdevice_system_settings" "trname" {
  allow_linkdown_path            = "enable"
  allow_subnet_overlap           = "disable"
  application_bandwidth_tracking = "enable"
  asymroute                      = "enable"
  asymroute_icmp                 = "disable"
  device_name                    = var.device_name # not required if setting is at provider
  device_vdom                    = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `allow_linkdown_path` - Enable/disable link down path. Valid values: `disable`, `enable`.

* `allow_subnet_overlap` - Enable/disable allowing interface subnets to use overlapping IP addresses. Valid values: `disable`, `enable`.

* `application_bandwidth_tracking` - Enable/disable application bandwidth tracking. Valid values: `disable`, `enable`.

* `asymroute` - Enable/disable IPv4 asymmetric routing. Valid values: `disable`, `enable`.

* `asymroute_icmp` - Enable/disable ICMP asymmetric routing. Valid values: `disable`, `enable`.

* `asymroute6` - Enable/disable asymmetric IPv6 routing. Valid values: `disable`, `enable`.

* `asymroute6_icmp` - Enable/disable asymmetric ICMPv6 routing. Valid values: `disable`, `enable`.

* `auxiliary_session` - Enable/disable auxiliary session. Valid values: `disable`, `enable`.

* `bfd` - Enable/disable Bi-directional Forwarding Detection (BFD) on all interfaces. Valid values: `disable`, `enable`.

* `bfd_desired_min_tx` - BFD desired minimal transmit interval (1 - 100000 ms, default = 250).
* `bfd_detect_mult` - BFD detection multiplier (1 - 50, default = 3).
* `bfd_dont_enforce_src_port` - Enable to not enforce verifying the source port of BFD Packets. Valid values: `disable`, `enable`.

* `bfd_required_min_rx` - BFD required minimal receive interval (1 - 100000 ms, default = 250).
* `block_land_attack` - Enable/disable blocking of land attacks. Valid values: `disable`, `enable`.

* `central_nat` - Enable/disable central NAT. Valid values: `disable`, `enable`.

* `comments` - VDOM comments.
* `consolidated_firewall_mode` - Consolidated firewall mode. Valid values: `disable`, `enable`.

* `default_app_port_as_service` - Enable/disable policy service enforcement based on application default ports. Valid values: `disable`, `enable`.

* `default_policy_expiry_days` - Default policy expiry in days (0 - 365 days, default = 30).
* `default_voip_alg_mode` - Configure how the FortiGate handles VoIP traffic when a policy that accepts the traffic doesn't include a VoIP profile. Valid values: `proxy-based`, `kernel-helper-based`.

* `deny_tcp_with_icmp` - Enable/disable denying TCP by sending an ICMP communication prohibited packet. Valid values: `disable`, `enable`.

* `detect_unknown_esp` - Enable/disable detection of unknown ESP packets (default = enable). Valid values: `disable`, `enable`.

* `device` - Interface to use for management access for NAT mode.
* `dhcp_proxy` - Enable/disable the DHCP Proxy. Valid values: `disable`, `enable`.

* `dhcp_proxy_interface` - Specify outgoing interface to reach server.
* `dhcp_proxy_interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `dhcp_server_ip` - DHCP Server IPv4 address.
* `dhcp6_server_ip` - DHCPv6 server IPv6 address.
* `discovered_device_timeout` - Timeout for discovered devices (1 - 365 days, default = 28).
* `dp_load_distribution_method` - Per VDOM DP load distribution method. Valid values: `src-ip`, `dst-ip`, `src-dst-ip`, `src-ip-sport`, `dst-ip-dport`, `src-dst-ip-sport-dport`, `to-master`, `derived`, `to-primary`.

* `dyn_addr_session_check` - Enable/disable dirty session check caused by dynamic address updates. Valid values: `disable`, `enable`.

* `ecmp_max_paths` - Maximum number of Equal Cost Multi-Path (ECMP) next-hops. Set to 1 to disable ECMP routing (1 - 255, default = 255).
* `email_portal_check_dns` - Enable/disable using DNS to validate email addresses collected by a captive portal. Valid values: `disable`, `enable`.

* `ext_resource_session_check` - Enable/disable dirty session check caused by external resource updates. Valid values: `disable`, `enable`.

* `firewall_session_dirty` - Select how to manage sessions affected by firewall policy configuration changes. Valid values: `check-all`, `check-new`, `check-policy-option`.

* `fqdn_session_check` - Enable/disable dirty session check caused by FQDN updates. Valid values: `disable`, `enable`.

* `fw_session_hairpin` - Enable/disable checking for a matching policy each time hairpin traffic goes through the FortiGate. Valid values: `disable`, `enable`.

* `gateway` - Transparent mode IPv4 default gateway IP address.
* `gateway6` - Transparent mode IPv6 default gateway IP address.
* `gtp_asym_fgsp` - Enable/disable GTP asymmetric traffic handling on FGSP. Valid values: `disable`, `enable`.

* `gtp_monitor_mode` - Enable/disable GTP monitor mode (VDOM level). Valid values: `disable`, `enable`.

* `gui_advanced_policy` - Enable/disable advanced policy configuration on the GUI. Valid values: `disable`, `enable`.

* `gui_advanced_wireless_features` - Enable/disable advanced wireless features in GUI. Valid values: `disable`, `enable`.

* `gui_allow_unnamed_policy` - Enable/disable the requirement for policy naming on the GUI. Valid values: `disable`, `enable`.

* `gui_antivirus` - Enable/disable AntiVirus on the GUI. Valid values: `disable`, `enable`.

* `gui_ap_profile` - Enable/disable FortiAP profiles on the GUI. Valid values: `disable`, `enable`.

* `gui_application_control` - Enable/disable application control on the GUI. Valid values: `disable`, `enable`.

* `gui_casb` - Enable/disable Inline-CASB on the GUI. Valid values: `disable`, `enable`.

* `gui_default_policy_columns` - Default columns to display for policy lists on GUI.
* `gui_dhcp_advanced` - Enable/disable advanced DHCP options on the GUI. Valid values: `disable`, `enable`.

* `gui_dlp_profile` - Enable/disable Data Loss Prevention on the GUI. Valid values: `disable`, `enable`.

* `gui_dns_database` - Enable/disable DNS database settings on the GUI. Valid values: `disable`, `enable`.

* `gui_dnsfilter` - Enable/disable DNS Filtering on the GUI. Valid values: `disable`, `enable`.

* `gui_domain_ip_reputation` - Gui-Domain-Ip-Reputation. Valid values: `disable`, `enable`.

* `gui_dos_policy` - Enable/disable DoS policies on the GUI. Valid values: `disable`, `enable`.

* `gui_dynamic_profile_display` - Gui-Dynamic-Profile-Display. Valid values: `disable`, `enable`.

* `gui_dynamic_device_os_id` - Enable/disable Create dynamic addresses to manage known devices. Valid values: `disable`, `enable`.

* `gui_dynamic_routing` - Enable/disable dynamic routing on the GUI. Valid values: `disable`, `enable`.

* `gui_email_collection` - Enable/disable email collection on the GUI. Valid values: `disable`, `enable`.

* `gui_endpoint_control` - Enable/disable endpoint control on the GUI. Valid values: `disable`, `enable`.

* `gui_endpoint_control_advanced` - Enable/disable advanced endpoint control options on the GUI. Valid values: `disable`, `enable`.

* `gui_enforce_change_summary` - Enforce change summaries for select tables in the GUI. Valid values: `disable`, `require`, `optional`.

* `gui_explicit_proxy` - Enable/disable the explicit proxy on the GUI. Valid values: `disable`, `enable`.

* `gui_file_filter` - Enable/disable File-filter on the GUI. Valid values: `disable`, `enable`.

* `gui_fortiap_split_tunneling` - Enable/disable FortiAP split tunneling on the GUI. Valid values: `disable`, `enable`.

* `gui_fortiextender_controller` - Enable/disable FortiExtender on the GUI. Valid values: `disable`, `enable`.

* `gui_icap` - Enable/disable ICAP on the GUI. Valid values: `disable`, `enable`.

* `gui_implicit_policy` - Enable/disable implicit firewall policies on the GUI. Valid values: `disable`, `enable`.

* `gui_ips` - Enable/disable IPS on the GUI. Valid values: `disable`, `enable`.

* `gui_load_balance` - Enable/disable server load balancing on the GUI. Valid values: `disable`, `enable`.

* `gui_local_in_policy` - Enable/disable Local-In policies on the GUI. Valid values: `disable`, `enable`.

* `gui_local_reports` - Enable/disable local reports on the GUI. Valid values: `disable`, `enable`.

* `gui_multicast_policy` - Enable/disable multicast firewall policies on the GUI. Valid values: `disable`, `enable`.

* `gui_multiple_interface_policy` - Enable/disable adding multiple interfaces to a policy on the GUI. Valid values: `disable`, `enable`.

* `gui_multiple_utm_profiles` - Gui-Multiple-Utm-Profiles. Valid values: `disable`, `enable`.

* `gui_nat46_64` - Gui-Nat46-64. Valid values: `disable`, `enable`.

* `gui_object_colors` - Enable/disable object colors on the GUI. Valid values: `disable`, `enable`.

* `gui_per_policy_disclaimer` - Enable/disable policy disclaimer on the GUI. Valid values: `disable`, `enable`.

* `gui_ot` - Enable/disable Operational technology features on the GUI. Valid values: `disable`, `enable`.

* `gui_policy_based_ipsec` - Enable/disable policy-based IPsec VPN on the GUI. Valid values: `disable`, `enable`.

* `gui_policy_disclaimer` - Enable/disable policy disclaimer on the GUI. Valid values: `disable`, `enable`.

* `gui_replacement_message_groups` - Enable/disable replacement message groups on the GUI. Valid values: `disable`, `enable`.

* `gui_proxy_inspection` - Enable/disable the proxy features on the GUI. Valid values: `disable`, `enable`.

* `gui_route_tag_address_creation` - Enable/disable route-tag addresses on the GUI. Valid values: `disable`, `enable`.

* `gui_security_profile_group` - Enable/disable Security Profile Groups on the GUI. Valid values: `disable`, `enable`.

* `gui_spamfilter` - Enable/disable Antispam on the GUI. Valid values: `disable`, `enable`.

* `gui_sslvpn` - Enable/disable SSL-VPN settings pages on the GUI. Valid values: `disable`, `enable`.

* `gui_sslvpn_personal_bookmarks` - Enable/disable SSL-VPN personal bookmark management on the GUI. Valid values: `disable`, `enable`.

* `gui_sslvpn_realms` - Enable/disable SSL-VPN realms on the GUI. Valid values: `disable`, `enable`.

* `gui_switch_controller` - Enable/disable the switch controller on the GUI. Valid values: `disable`, `enable`.

* `gui_threat_weight` - Enable/disable threat weight on the GUI. Valid values: `disable`, `enable`.

* `gui_traffic_shaping` - Enable/disable traffic shaping on the GUI. Valid values: `disable`, `enable`.

* `gui_videofilter` - Enable/disable Video filtering on the GUI. Valid values: `disable`, `enable`.

* `gui_virtual_patch_profile` - Enable/disable Virtual Patching on the GUI. Valid values: `disable`, `enable`.

* `gui_voip_profile` - Enable/disable VoIP profiles on the GUI. Valid values: `disable`, `enable`.

* `gui_vpn` - Enable/disable IPsec VPN settings pages on the GUI. Valid values: `disable`, `enable`.

* `gui_waf_profile` - Enable/disable Web Application Firewall on the GUI. Valid values: `disable`, `enable`.

* `gui_wan_load_balancing` - Enable/disable SD-WAN on the GUI. Valid values: `disable`, `enable`.

* `gui_wanopt_cache` - Enable/disable WAN Optimization and Web Caching on the GUI. Valid values: `disable`, `enable`.

* `gui_webfilter` - Enable/disable Web filtering on the GUI. Valid values: `disable`, `enable`.

* `gui_webfilter_advanced` - Enable/disable advanced web filtering on the GUI. Valid values: `disable`, `enable`.

* `gui_wireless_controller` - Enable/disable the wireless controller on the GUI. Valid values: `disable`, `enable`.

* `gui_ztna` - Enable/disable Zero Trust Network Access features on the GUI. Valid values: `disable`, `enable`.

* `h323_direct_model` - Enable/disable H323 direct model. Valid values: `disable`, `enable`.

* `http_external_dest` - Offload HTTP traffic to FortiWeb or FortiCache. Valid values: `fortiweb`, `forticache`.

* `hyperscale_default_policy_action` - Hyperscale default policy action. Valid values: `drop-on-hardware`, `forward-to-host`.

* `ike_dn_format` - Configure IKE ASN.1 Distinguished Name format conventions. Valid values: `with-space`, `no-space`.

* `ike_policy_route` - Enable/disable IKE Policy Based Routing (PBR). Valid values: `disable`, `enable`.

* `ike_port` - UDP port for IKE/IPsec traffic (default 500).
* `ike_quick_crash_detect` - Enable/disable IKE quick crash detection (RFC 6290). Valid values: `disable`, `enable`.

* `ike_session_resume` - Enable/disable IKEv2 session resumption (RFC 5723). Valid values: `disable`, `enable`.

* `implicit_allow_dns` - Enable/disable implicitly allowing DNS traffic. Valid values: `disable`, `enable`.

* `ike_tcp_port` - TCP port for IKE/IPsec traffic (default 4500).
* `internet_service_app_ctrl_size` - Maximum number of tuple entries (protocol, port, IP address, application ID) stored by the FortiGate unit (0 - 4294967295, default = 32768). A smaller value limits the FortiGate unit from learning about internet applications.
* `internet_service_database_cache` - Enable/disable Internet Service database caching. Valid values: `disable`, `enable`.

* `ip` - IP address and netmask.
* `ip6` - IPv6 address prefix for NAT mode.
* `lan_extension_controller_addr` - Controller IP address or FQDN to connect.
* `link_down_access` - Enable/disable link down access traffic. Valid values: `disable`, `enable`.

* `lldp_reception` - Enable/disable Link Layer Discovery Protocol (LLDP) reception for this VDOM or apply global settings to this VDOM. Valid values: `disable`, `enable`, `global`.

* `lldp_transmission` - Enable/disable Link Layer Discovery Protocol (LLDP) transmission for this VDOM or apply global settings to this VDOM. Valid values: `enable`, `disable`, `global`.

* `location_id` - Local location ID in the form of an IPv4 address.
* `mac_ttl` - Duration of MAC addresses in Transparent mode (300 - 8640000 sec, default = 300).
* `manageip` - Transparent mode IPv4 management IP address and netmask.
* `manageip6` - Transparent mode IPv6 management IP address and netmask.
* `motherboard_traffic_forwarding` - Motherboard-Traffic-Forwarding. Valid values: `icmp`, `admin`, `auth`.

* `multicast_forward` - Enable/disable multicast forwarding. Valid values: `disable`, `enable`.

* `multicast_skip_policy` - Enable/disable allowing multicast traffic through the FortiGate without a policy check. Valid values: `disable`, `enable`.

* `multicast_ttl_notchange` - Enable/disable preventing the FortiGate from changing the TTL for forwarded multicast packets. Valid values: `disable`, `enable`.

* `nat46_force_ipv4_packet_forwarding` - Enable/disable mandatory IPv4 packet forwarding in NAT46. Valid values: `disable`, `enable`.

* `nat46_generate_ipv6_fragment_header` - Enable/disable NAT46 IPv6 fragment header generation. Valid values: `disable`, `enable`.

* `nat64_force_ipv6_packet_forwarding` - Enable/disable mandatory IPv6 packet forwarding in NAT64. Valid values: `disable`, `enable`.

* `ngfw_mode` - Next Generation Firewall (NGFW) mode. Valid values: `profile-based`, `policy-based`.

* `npu_group_id` - npu-group-index.
* `opmode` - Firewall operation mode (NAT or Transparent). Valid values: `nat`, `transparent`.

* `pfcp_monitor_mode` - Enable/disable PFCP monitor mode (VDOM level). Valid values: `disable`, `enable`.

* `policy_offload_level` - Configure firewall policy offload level. Valid values: `disable`, `default`, `dos-offload`, `full-offload`.

* `prp_trailer_action` - Enable/disable action to take on PRP trailer. Valid values: `disable`, `enable`.

* `sccp_port` - TCP port the SCCP proxy monitors for SCCP traffic (0 - 65535, default = 2000).
* `sctp_session_without_init` - Enable/disable SCTP session creation without SCTP INIT. Valid values: `disable`, `enable`.

* `ses_denied_traffic` - Enable/disable including denied session in the session table. Valid values: `disable`, `enable`.

* `session_insert_trial` - Trial session insert. Valid values: `disable`, `enable`.

* `sip_expectation` - Enable/disable the SIP kernel session helper to create an expectation for port 5060. Valid values: `disable`, `enable`.

* `sip_nat_trace` - Enable/disable recording the original SIP source IP address when NAT is used. Valid values: `disable`, `enable`.

* `sip_ssl_port` - TCP port the SIP proxy monitors for SIP SSL/TLS traffic (0 - 65535, default = 5061).
* `sip_tcp_port` - TCP port the SIP proxy monitors for SIP traffic (0 - 65535, default = 5060).
* `sip_udp_port` - UDP port the SIP proxy monitors for SIP traffic (0 - 65535, default = 5060).
* `snat_hairpin_traffic` - Enable/disable source NAT (SNAT) for hairpin traffic. Valid values: `disable`, `enable`.

* `status` - Enable/disable this VDOM. Valid values: `disable`, `enable`.

* `strict_src_check` - Enable/disable strict source verification. Valid values: `disable`, `enable`.

* `tcp_session_without_syn` - Enable/disable allowing TCP session without SYN flags. Valid values: `disable`, `enable`.

* `trap_local_session` - Enable/disable local-in traffic session traps. Valid values: `disable`, `enable`.

* `trap_session_flag` - Trap session operation flags. Valid values: `udp-both`, `udp-reply`, `tcpudp-both`, `tcpudp-reply`, `trap-none`.

* `utf8_spam_tagging` - Enable/disable converting antispam tags to UTF-8 for better non-ASCII character support. Valid values: `disable`, `enable`.

* `v4_ecmp_mode` - IPv4 Equal-cost multi-path (ECMP) routing and load balancing mode. Valid values: `source-ip-based`, `weight-based`, `usage-based`, `source-dest-ip-based`.

* `vdom_type` - Vdom type (traffic, lan-extension or admin). Valid values: `traffic`, `admin`, `lan-extension`.

* `vpn_stats_log` - Enable/disable periodic VPN log statistics for one or more types of VPN. Separate names with a space. Valid values: `ipsec`, `pptp`, `l2tp`, `ssl`.

* `vpn_stats_period` - Period to send VPN log statistics (0 or 60 - 86400 sec).
* `wccp_cache_engine` - Enable/disable WCCP cache engine. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Settings can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_settings.labelname SystemSettings
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

