---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_ipsec_phase1interface"
description: |-
  Configure VPN remote gateway.
---

# fmgdevice_vpn_ipsec_phase1interface
Configure VPN remote gateway.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `ipv4_exclude_range`: `fmgdevice_vpn_ipsec_phase1interface_ipv4excluderange`
>- `ipv6_exclude_range`: `fmgdevice_vpn_ipsec_phase1interface_ipv6excluderange`



## Example Usage

```hcl
resource "fmgdevice_vpn_ipsec_phase1interface" "trname" {
  acct_verify      = "disable"
  add_gw_route     = "enable"
  add_route        = "disable"
  aggregate_member = "disable"
  aggregate_weight = 10
  name             = "your own value"
  device_name      = var.device_name # not required if setting is at provider
  device_vdom      = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `acct_verify` - Enable/disable verification of RADIUS accounting record. Valid values: `disable`, `enable`.

* `add_gw_route` - Enable/disable automatically add a route to the remote gateway. Valid values: `disable`, `enable`.

* `add_route` - Enable/disable control addition of a route to peer destination selector. Valid values: `disable`, `enable`.

* `addke1` - ADDKE1 group. Valid values: `0`, `1080`, `1081`, `1082`.

* `addke2` - ADDKE2 group. Valid values: `0`, `1080`, `1081`, `1082`.

* `addke3` - ADDKE3 group. Valid values: `0`, `1080`, `1081`, `1082`.

* `addke4` - ADDKE4 group. Valid values: `0`, `1080`, `1081`, `1082`.

* `addke5` - ADDKE5 group. Valid values: `0`, `1080`, `1081`, `1082`.

* `addke6` - ADDKE6 group. Valid values: `0`, `1080`, `1081`, `1082`.

* `addke7` - ADDKE7 group. Valid values: `0`, `1080`, `1081`, `1082`.

* `aggregate_member` - Enable/disable use as an aggregate member. Valid values: `disable`, `enable`.

* `aggregate_weight` - Link weight for aggregate.
* `assign_ip` - Enable/disable assignment of IP to IPsec interface via configuration method. Valid values: `disable`, `enable`.

* `assign_ip_from` - Method by which the IP address will be assigned. Valid values: `range`, `usrgrp`, `dhcp`, `name`.

* `authmethod` - Authentication method. Valid values: `psk`, `rsa-signature`, `signature`.

* `authmethod_remote` - Authentication method (remote side). Valid values: `psk`, `signature`.

* `authpasswd` - XAuth password (max 35 characters).
* `authusr` - XAuth user name.
* `authusrgrp` - Authentication user group.
* `auto_discovery_crossover` - Allow/block set-up of short-cut tunnels between different network IDs. Valid values: `block`, `allow`.

* `auto_discovery_forwarder` - Enable/disable forwarding auto-discovery short-cut messages. Valid values: `disable`, `enable`.

* `auto_discovery_offer_interval` - Interval between shortcut offer messages in seconds (1 - 300, default = 5).
* `auto_discovery_psk` - Enable/disable use of pre-shared secrets for authentication of auto-discovery tunnels. Valid values: `disable`, `enable`.

* `auto_discovery_receiver` - Enable/disable accepting auto-discovery short-cut messages. Valid values: `disable`, `enable`.

* `auto_discovery_sender` - Enable/disable sending auto-discovery short-cut messages. Valid values: `disable`, `enable`.

* `auto_discovery_shortcuts` - Control deletion of child short-cut tunnels when the parent tunnel goes down. Valid values: `independent`, `dependent`.

* `auto_negotiate` - Enable/disable automatic initiation of IKE SA negotiation. Valid values: `disable`, `enable`.

* `auto_transport_threshold` - Timeout in seconds before falling back to next transport protocol.
* `azure_ad_autoconnect` - Enable/disable Azure AD Auto-Connect for FortiClient. Valid values: `disable`, `enable`.

* `backup_gateway` - Instruct unity clients about the backup gateway address(es).
* `banner` - Message that unity client should display after connecting.
* `cert_id_validation` - Enable/disable cross validation of peer ID and the identity in the peer's certificate as specified in RFC 4945. Valid values: `disable`, `enable`.

* `cert_peer_username_strip` - Enable/disable domain stripping on certificate identity. Valid values: `disable`, `enable`.

* `cert_peer_username_validation` - Enable/disable cross validation of peer username and the identity in the peer's certificate. Valid values: `othername`, `rfc822name`, `cn`, `none`.

* `cert_trust_store` - CA certificate trust store. Valid values: `local`, `ems`.

* `certificate` - The names of up to 4 signed personal certificates.
* `childless_ike` - Enable/disable childless IKEv2 initiation (RFC 6023). Valid values: `disable`, `enable`.

* `client_auto_negotiate` - Enable/disable allowing the VPN client to bring up the tunnel when there is no traffic. Valid values: `disable`, `enable`.

* `client_keep_alive` - Enable/disable allowing the VPN client to keep the tunnel up when there is no traffic. Valid values: `disable`, `enable`.

* `client_resume` - Enable/disable resumption of offline FortiClient sessions.  When a FortiClient enabled laptop is closed or enters sleep/hibernate mode, enabling this feature allows FortiClient to keep the tunnel during this period, and allows users to immediately resume using the IPsec tunnel when the device wakes up. Valid values: `disable`, `enable`.

* `client_resume_interval` - Maximum time in seconds during which a VPN client may resume using a tunnel after a client PC has entered sleep mode or temporarily lost its network connection (120 - 172800, default = 1800).
* `comments` - Comment.
* `default_gw` - IPv4 address of default route gateway to use for traffic exiting the interface.
* `default_gw_priority` - Priority for default gateway route. A higher priority number signifies a less preferred route.
* `dev_id` - Device ID carried by the device ID notification.
* `dev_id_notification` - Enable/disable device ID notification. Valid values: `disable`, `enable`.

* `dhcp_ra_giaddr` - Relay agent gateway IP address to use in the giaddr field of DHCP requests.
* `dhcp6_ra_linkaddr` - Relay agent IPv6 link address to use in DHCP6 requests.
* `dhgrp` - DH group. Valid values: `1`, `2`, `5`, `14`, `15`, `16`, `17`, `18`, `19`, `20`, `21`, `27`, `28`, `29`, `30`, `31`, `32`.

* `digital_signature_auth` - Enable/disable IKEv2 Digital Signature Authentication (RFC 7427). Valid values: `disable`, `enable`.

* `distance` - Distance for routes added by IKE (1 - 255).
* `dns_mode` - DNS server mode. Valid values: `auto`, `manual`.

* `domain` - Instruct unity clients about the single default DNS domain.
* `dpd` - Dead Peer Detection mode. Valid values: `disable`, `enable`, `on-idle`, `on-demand`.

* `dpd_retrycount` - Number of DPD retry attempts.
* `dpd_retryinterval` - DPD retry interval.
* `eap` - Enable/disable IKEv2 EAP authentication. Valid values: `disable`, `enable`.

* `eap_cert_auth` - Enable/disable peer certificate authentication in addition to EAP if peer is a FortiClient endpoint. Valid values: `disable`, `enable`.

* `eap_exclude_peergrp` - Peer group excluded from EAP authentication.
* `eap_identity` - IKEv2 EAP peer identity type. Valid values: `use-id-payload`, `send-request`.

* `ems_sn_check` - Enable/disable verification of EMS serial number. Valid values: `use-global`, `enable`, `disable`.

* `encap_local_gw4` - Local IPv4 address of GRE/VXLAN tunnel.
* `encap_local_gw6` - Local IPv6 address of GRE/VXLAN tunnel.
* `encap_remote_gw4` - Remote IPv4 address of GRE/VXLAN tunnel.
* `encap_remote_gw6` - Remote IPv6 address of GRE/VXLAN tunnel.
* `encapsulation` - Enable/disable GRE/VXLAN/VPNID encapsulation. Valid values: `none`, `gre`, `vxlan`, `vpn-id-ipip`.

* `encapsulation_address` - Source for GRE/VXLAN tunnel address. Valid values: `ike`, `ipv4`, `ipv6`.

* `enforce_unique_id` - Enable/disable peer ID uniqueness check. Valid values: `disable`, `keep-new`, `keep-old`.

* `esn` - Extended sequence number (ESN) negotiation. Valid values: `disable`, `require`, `allow`.

* `exchange_fgt_device_id` - Enable/disable device identifier exchange with peer FortiGate units for use of VPN monitor data by FortiManager. Valid values: `disable`, `enable`.

* `exchange_interface_ip` - Enable/disable exchange of IPsec interface IP address. Valid values: `disable`, `enable`.

* `exchange_ip_addr4` - IPv4 address to exchange with peers.
* `exchange_ip_addr6` - IPv6 address to exchange with peers.
* `fallback_tcp_threshold` - Timeout in seconds before falling back IKE/IPsec traffic to tcp.
* `fec_base` - Number of base Forward Error Correction packets (1 - 20).
* `fec_codec` - Forward Error Correction encoding/decoding algorithm. Valid values: `rs`, `xor`.

* `fec_egress` - Enable/disable Forward Error Correction for egress IPsec traffic. Valid values: `disable`, `enable`.

* `fec_health_check` - SD-WAN health check.
* `fec_ingress` - Enable/disable Forward Error Correction for ingress IPsec traffic. Valid values: `disable`, `enable`.

* `fec_mapping_profile` - Forward Error Correction (FEC) mapping profile.
* `fec_receive_timeout` - Timeout in milliseconds before dropping Forward Error Correction packets (1 - 1000).
* `fec_redundant` - Number of redundant Forward Error Correction packets (1 - 5 for reed-solomon, 1 for xor).
* `fec_send_timeout` - Timeout in milliseconds before sending Forward Error Correction packets (1 - 1000).
* `fgsp_sync` - Enable/disable IPsec syncing of tunnels for FGSP IPsec. Valid values: `disable`, `enable`.

* `forticlient_enforcement` - Enable/disable FortiClient enforcement. Valid values: `disable`, `enable`.

* `fortinet_esp` - Enable/disable Fortinet ESP encapsulaton. Valid values: `disable`, `enable`.

* `fragmentation` - Enable/disable fragment IKE message on re-transmission. Valid values: `disable`, `enable`.

* `fragmentation_mtu` - IKE fragmentation MTU (500 - 16000).
* `group_authentication` - Enable/disable IKEv2 IDi group authentication. Valid values: `disable`, `enable`.

* `group_authentication_secret` - Password for IKEv2 ID group authentication. ASCII string or hexadecimal indicated by a leading 0x.
* `ha_sync_esp_seqno` - Enable/disable sequence number jump ahead for IPsec HA. Valid values: `disable`, `enable`.

* `idle_timeout` - Enable/disable IPsec tunnel idle timeout. Valid values: `disable`, `enable`.

* `idle_timeoutinterval` - IPsec tunnel idle timeout in minutes (5 - 43200).
* `ike_version` - IKE protocol version. Valid values: `1`, `2`.

* `inbound_dscp_copy` - Enable/disable copy the dscp in the ESP header to the inner IP Header. Valid values: `disable`, `enable`.

* `include_local_lan` - Enable/disable allow local LAN access on unity clients. Valid values: `disable`, `enable`.

* `interface` - Local physical, aggregate, or VLAN outgoing interface.
* `internal_domain_list` - One or more internal domain names in quotes separated by spaces.
* `ip_delay_interval` - IP address reuse delay interval in seconds (0 - 28800).
* `ip_fragmentation` - Determine whether IP packets are fragmented before or after IPsec encapsulation. Valid values: `pre-encapsulation`, `post-encapsulation`.

* `ip_version` - IP version to use for VPN interface. Valid values: `4`, `6`.

* `ipsec_tunnel_slot` - Slot at which IPsec tunnel will be establishd. Valid values: `master`, `auto`, `FPC1`, `FPC2`, `FPC3`, `FPC4`, `FPC5`, `FPC6`, `FPC7`, `FPC8`, `FPC9`, `FPC10`, `FPM3`, `FPM4`, `FPM5`, `FPM6`, `FPM7`, `FPM8`, `FPM9`, `FPM10`, `FPM11`, `FPM12`.

* `ipv4_dns_server1` - IPv4 DNS server 1.
* `ipv4_dns_server2` - IPv4 DNS server 2.
* `ipv4_dns_server3` - IPv4 DNS server 3.
* `ipv4_end_ip` - End of IPv4 range.
* `ipv4_exclude_range` - Ipv4-Exclude-Range. The structure of `ipv4_exclude_range` block is documented below.
* `ipv4_name` - IPv4 address name.
* `ipv4_netmask` - IPv4 Netmask.
* `ipv4_split_exclude` - IPv4 subnets that should not be sent over the IPsec tunnel.
* `ipv4_split_include` - IPv4 split-include subnets.
* `ipv4_start_ip` - Start of IPv4 range.
* `ipv4_wins_server1` - WINS server 1.
* `ipv4_wins_server2` - WINS server 2.
* `ipv6_auto_linklocal` - Enable/disable auto generation of IPv6 link-local address using last 8 bytes of mode-cfg assigned IPv6 address. Valid values: `disable`, `enable`.

* `ipv6_dns_server1` - IPv6 DNS server 1.
* `ipv6_dns_server2` - IPv6 DNS server 2.
* `ipv6_dns_server3` - IPv6 DNS server 3.
* `ipv6_end_ip` - End of IPv6 range.
* `ipv6_exclude_range` - Ipv6-Exclude-Range. The structure of `ipv6_exclude_range` block is documented below.
* `ipv6_name` - IPv6 address name.
* `ipv6_prefix` - IPv6 prefix.
* `ipv6_split_exclude` - IPv6 subnets that should not be sent over the IPsec tunnel.
* `ipv6_split_include` - IPv6 split-include subnets.
* `ipv6_start_ip` - Start of IPv6 range.
* `keepalive` - NAT-T keep alive interval.
* `keylife` - Time to wait in seconds before phase 1 encryption key expires.
* `kms` - Key Management Services server.
* `link_cost` - VPN tunnel underlay link cost.
* `local_gw` - IPv4 address of the local gateway's external interface.
* `local_gw6` - IPv6 address of the local gateway's external interface.
* `localid` - Local ID.
* `localid_type` - Local ID type. Valid values: `auto`, `fqdn`, `user-fqdn`, `keyid`, `address`, `asn1dn`.

* `loopback_asymroute` - Enable/disable asymmetric routing for IKE traffic on loopback interface. Valid values: `disable`, `enable`.

* `mesh_selector_type` - Add selectors containing subsets of the configuration depending on traffic. Valid values: `disable`, `subnet`, `host`.

* `mode` - The ID protection mode used to establish a secure channel. Valid values: `main`, `aggressive`.

* `mode_cfg` - Enable/disable configuration method. Valid values: `disable`, `enable`.

* `mode_cfg_allow_client_selector` - Enable/disable mode-cfg client to use custom phase2 selectors. Valid values: `disable`, `enable`.

* `monitor` - IPsec interface as backup for primary interface.
* `monitor_hold_down_delay` - Time to wait in seconds before recovery once primary re-establishes.
* `monitor_hold_down_time` - Time of day at which to fail back to primary after it re-establishes.
* `monitor_hold_down_type` - Recovery time method when primary interface re-establishes. Valid values: `immediate`, `delay`, `time`.

* `monitor_hold_down_weekday` - Day of the week to recover once primary re-establishes. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `everyday`.

* `monitor_min` - Minimum number of links to become degraded before activating this interface.  Zero (0) means all links must be down before activating this interface.
* `name` - IPsec remote gateway name.
* `nattraversal` - Enable/disable NAT traversal. Valid values: `disable`, `enable`, `forced`.

* `negotiate_timeout` - IKE SA negotiation timeout in seconds (1 - 300).
* `net_device` - Enable/disable kernel device creation. Valid values: `disable`, `enable`.

* `network_id` - VPN gateway network ID.
* `network_overlay` - Enable/disable network overlays. Valid values: `disable`, `enable`.

* `npu_offload` - Enable/disable offloading NPU. Valid values: `disable`, `enable`.

* `packet_redistribution` - Enable/disable packet distribution (RPS) on the IPsec interface. Valid values: `disable`, `enable`.

* `passive_mode` - Enable/disable IPsec passive mode for static tunnels. Valid values: `disable`, `enable`.

* `peer` - Accept this peer certificate.
* `peergrp` - Accept this peer certificate group.
* `peerid` - Accept this peer identity.
* `peertype` - Accept this peer type. Valid values: `any`, `one`, `dialup`, `peer`, `peergrp`.

* `ppk` - Enable/disable IKEv2 Postquantum Preshared Key (PPK). Valid values: `disable`, `allow`, `require`.

* `ppk_identity` - IKEv2 Postquantum Preshared Key Identity.
* `ppk_secret` - IKEv2 Postquantum Preshared Key (ASCII string or hexadecimal encoded with a leading 0x).
* `priority` - Priority for routes added by IKE (1 - 65535).
* `proposal` - Phase1 proposal. Valid values: `des-md5`, `des-sha1`, `3des-md5`, `3des-sha1`, `aes128-md5`, `aes128-sha1`, `aes192-md5`, `aes192-sha1`, `aes256-md5`, `aes256-sha1`, `des-sha256`, `3des-sha256`, `aes128-sha256`, `aes192-sha256`, `aes256-sha256`, `des-sha384`, `des-sha512`, `3des-sha384`, `3des-sha512`, `aes128-sha384`, `aes128-sha512`, `aes192-sha384`, `aes192-sha512`, `aes256-sha384`, `aes256-sha512`, `aria128-md5`, `aria128-sha1`, `aria128-sha256`, `aria128-sha384`, `aria128-sha512`, `aria192-md5`, `aria192-sha1`, `aria192-sha256`, `aria192-sha384`, `aria192-sha512`, `aria256-md5`, `aria256-sha1`, `aria256-sha256`, `aria256-sha384`, `aria256-sha512`, `seed-md5`, `seed-sha1`, `seed-sha256`, `seed-sha384`, `seed-sha512`, `aes128gcm-prfsha1`, `aes128gcm-prfsha256`, `aes128gcm-prfsha384`, `aes128gcm-prfsha512`, `aes256gcm-prfsha1`, `aes256gcm-prfsha256`, `aes256gcm-prfsha384`, `aes256gcm-prfsha512`, `chacha20poly1305-prfsha1`, `chacha20poly1305-prfsha256`, `chacha20poly1305-prfsha384`, `chacha20poly1305-prfsha512`.

* `psksecret` - Pre-shared secret for PSK authentication (ASCII string or hexadecimal encoded with a leading 0x).
* `psksecret_remote` - Pre-shared secret for remote side PSK authentication (ASCII string or hexadecimal encoded with a leading 0x).
* `qkd` - Enable/disable use of Quantum Key Distribution (QKD) server. Valid values: `disable`, `allow`, `require`.

* `qkd_profile` - Quantum Key Distribution (QKD) server profile.
* `reauth` - Enable/disable re-authentication upon IKE SA lifetime expiration. Valid values: `disable`, `enable`.

* `rekey` - Enable/disable phase1 rekey. Valid values: `disable`, `enable`.

* `remote_gw` - IPv4 address of the remote gateway's external interface.
* `remote_gw_country` - IPv4 addresses associated to a specific country.
* `remote_gw_end_ip` - Last IPv4 address in the range.
* `remote_gw_match` - Set type of IPv4 remote gateway address matching. Valid values: `any`, `ipmask`, `iprange`, `geography`.

* `remote_gw_start_ip` - First IPv4 address in the range.
* `remote_gw_subnet` - IPv4 address and subnet mask.
* `remote_gw_ztna_tags` - IPv4 ZTNA posture tags.
* `remote_gw6` - IPv6 address of the remote gateway's external interface.
* `remote_gw6_country` - IPv6 addresses associated to a specific country.
* `remote_gw6_end_ip` - Last IPv6 address in the range.
* `remote_gw6_match` - Set type of IPv6 remote gateway address matching. Valid values: `any`, `iprange`, `geography`, `ipprefix`.

* `remote_gw6_start_ip` - First IPv6 address in the range.
* `remote_gw6_subnet` - IPv6 address and prefix.
* `remotegw_ddns` - Domain name of remote gateway. For example, name.ddns.com.
* `rsa_signature_format` - Digital Signature Authentication RSA signature format. Valid values: `pkcs1`, `pss`.

* `rsa_signature_hash_override` - Enable/disable IKEv2 RSA signature hash algorithm override. Valid values: `disable`, `enable`.

* `save_password` - Enable/disable saving XAuth username and password on VPN clients. Valid values: `disable`, `enable`.

* `send_cert_chain` - Enable/disable sending certificate chain. Valid values: `disable`, `enable`.

* `shared_idle_timeout` - Enable/disable IPsec tunnel shared idle timeout. Valid values: `disable`, `enable`.

* `signature_hash_alg` - Digital Signature Authentication hash algorithms. Valid values: `sha1`, `sha2-256`, `sha2-384`, `sha2-512`.

* `split_include_service` - Split-include services.
* `suite_b` - Use Suite-B. Valid values: `disable`, `suite-b-gcm-128`, `suite-b-gcm-256`.

* `transit_gateway` - IPsec tunnel created by autoscaling to be used as a transit gateway. Valid values: `disable`, `enable`.

* `tunnel_search` - Tunnel search method for when the interface is shared. Valid values: `selectors`, `nexthop`.

* `transport` - Set IKE transport protocol. Valid values: `udp`, `tcp`, `udp-fallback-tcp`.

* `type` - Remote gateway type. Valid values: `static`, `dynamic`, `ddns`.

* `unity_support` - Enable/disable support for Cisco UNITY Configuration Method extensions. Valid values: `disable`, `enable`.

* `usrgrp` - User group name for dialup peers.
* `vni` - VNI of VXLAN tunnel.
* `wizard_type` - GUI VPN Wizard Type. Valid values: `custom`, `dialup-forticlient`, `dialup-ios`, `dialup-android`, `dialup-cisco`, `static-fortigate`, `static-cisco`, `dialup-windows`, `dialup-fortigate`, `dialup-cisco-fw`, `simplified-static-fortigate`, `hub-fortigate-auto-discovery`, `spoke-fortigate-auto-discovery`, `static-fortiproxy`, `dialup-fortiproxy`, `simplified-static-fortiproxy`, `hub-fortiproxy-auto-discovery`, `spoke-fortiproxy-auto-discovery`.

* `xauthtype` - XAuth type. Valid values: `disable`, `client`, `pap`, `chap`, `auto`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `ipv4_exclude_range` block supports:

* `end_ip` - End of IPv4 exclusive range.
* `id` - ID.
* `start_ip` - Start of IPv4 exclusive range.

The `ipv6_exclude_range` block supports:

* `end_ip` - End of IPv6 exclusive range.
* `id` - ID.
* `start_ip` - Start of IPv6 exclusive range.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Vpn IpsecPhase1Interface can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_ipsec_phase1interface.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

