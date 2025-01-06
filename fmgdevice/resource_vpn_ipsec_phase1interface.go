// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure VPN remote gateway.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnIpsecPhase1Interface() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnIpsecPhase1InterfaceCreate,
		Read:   resourceVpnIpsecPhase1InterfaceRead,
		Update: resourceVpnIpsecPhase1InterfaceUpdate,
		Delete: resourceVpnIpsecPhase1InterfaceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"device_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"device_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"acct_verify": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"add_gw_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"add_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"addke1": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"addke2": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"addke3": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"addke4": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"addke5": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"addke6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"addke7": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"aggregate_member": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"aggregate_weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"assign_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"assign_ip_from": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"authmethod": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authmethod_remote": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"authpasswd": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"authusr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"authusrgrp": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"auto_discovery_crossover": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_discovery_forwarder": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_discovery_offer_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auto_discovery_psk": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_discovery_receiver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_discovery_sender": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_discovery_shortcuts": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_negotiate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_transport_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"azure_ad_autoconnect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"backup_gateway": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"banner": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cert_id_validation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cert_peer_username_strip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cert_peer_username_validation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cert_trust_store": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"certificate": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"childless_ike": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_auto_negotiate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_keep_alive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_resume": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_resume_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"default_gw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"default_gw_priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dev_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dev_id_notification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_ra_giaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp6_ra_linkaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhgrp": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"digital_signature_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"distance": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dns_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dpd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dpd_retrycount": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dpd_retryinterval": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"eap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"eap_cert_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eap_exclude_peergrp": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"eap_identity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ems_sn_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encap_local_gw4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encap_local_gw6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encap_remote_gw4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encap_remote_gw6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encapsulation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encapsulation_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enforce_unique_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"esn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"exchange_fgt_device_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"exchange_interface_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"exchange_ip_addr4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"exchange_ip_addr6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fallback_tcp_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fec_base": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fec_codec": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fec_egress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fec_health_check": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fec_ingress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fec_mapping_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fec_receive_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fec_redundant": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fec_send_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fgsp_sync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forticlient_enforcement": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortinet_esp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fragmentation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fragmentation_mtu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"group_authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"group_authentication_secret": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"ha_sync_esp_seqno": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"idle_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"idle_timeoutinterval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ike_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"inbound_dscp_copy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"include_local_lan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internal_domain_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ip_delay_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip_fragmentation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec_tunnel_slot": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv4_dns_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv4_dns_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv4_dns_server3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv4_end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv4_exclude_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"start_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ipv4_name": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ipv4_netmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv4_split_exclude": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ipv4_split_include": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ipv4_start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv4_wins_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv4_wins_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_auto_linklocal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_dns_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_dns_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_dns_server3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_exclude_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"start_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ipv6_name": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ipv6_prefix": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ipv6_split_exclude": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ipv6_split_include": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ipv6_start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"keepalive": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"keylife": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"kms": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"link_cost": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"local_gw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_gw6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"localid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"localid_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"loopback_asymroute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mesh_selector_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode_cfg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode_cfg_allow_client_selector": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"monitor": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"monitor_hold_down_delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"monitor_hold_down_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monitor_hold_down_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monitor_hold_down_weekday": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"monitor_min": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"nattraversal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"negotiate_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"net_device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"network_overlay": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"npu_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"packet_redistribution": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"passive_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"peer": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"peergrp": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"peerid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"peertype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ppk": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ppk_identity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ppk_secret": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"proposal": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"psksecret": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"psksecret_remote": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"qkd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"qkd_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"reauth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rekey": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_gw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_gw_country": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_gw_end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_gw_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_gw_start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_gw_subnet": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"remote_gw_ztna_tags": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"remote_gw6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_gw6_country": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_gw6_end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_gw6_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_gw6_start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_gw6_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remotegw_ddns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rsa_signature_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rsa_signature_hash_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"save_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"send_cert_chain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"shared_idle_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"signature_hash_alg": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"split_include_service": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"suite_b": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"transit_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tunnel_search": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"transport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unity_support": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"usrgrp": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vni": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"wizard_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"xauthtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceVpnIpsecPhase1InterfaceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectVpnIpsecPhase1Interface(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecPhase1Interface resource while getting object: %v", err)
	}

	_, err = c.CreateVpnIpsecPhase1Interface(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecPhase1Interface resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceVpnIpsecPhase1InterfaceRead(d, m)
}

func resourceVpnIpsecPhase1InterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectVpnIpsecPhase1Interface(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecPhase1Interface resource while getting object: %v", err)
	}

	_, err = c.UpdateVpnIpsecPhase1Interface(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecPhase1Interface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceVpnIpsecPhase1InterfaceRead(d, m)
}

func resourceVpnIpsecPhase1InterfaceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	err = c.DeleteVpnIpsecPhase1Interface(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting VpnIpsecPhase1Interface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnIpsecPhase1InterfaceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if device_vdom == "" {
		device_vdom = importOptionChecking(m.(*FortiClient).Cfg, "device_vdom")
		if device_vdom == "" {
			return fmt.Errorf("Parameter device_vdom is missing")
		}
		if err = d.Set("device_vdom", device_vdom); err != nil {
			return fmt.Errorf("Error set params device_vdom: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	o, err := c.ReadVpnIpsecPhase1Interface(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecPhase1Interface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnIpsecPhase1Interface(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecPhase1Interface resource from API: %v", err)
	}
	return nil
}

func flattenVpnIpsecPhase1InterfaceAcctVerify(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAddGwRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAddRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAddke1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceAddke2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceAddke3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceAddke4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceAddke5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceAddke6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceAddke7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceAggregateMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAggregateWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAssignIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAssignIpFrom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAuthmethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAuthmethodRemote(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAuthusr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAuthusrgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceAutoDiscoveryCrossover(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAutoDiscoveryForwarder(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAutoDiscoveryOfferInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAutoDiscoveryPsk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAutoDiscoveryReceiver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAutoDiscoverySender(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAutoDiscoveryShortcuts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAutoNegotiate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAutoTransportThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAzureAdAutoconnect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceBackupGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceBanner(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceCertIdValidation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceCertPeerUsernameStrip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceCertPeerUsernameValidation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceCertTrustStore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceChildlessIke(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceClientAutoNegotiate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceClientKeepAlive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceClientResume(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceClientResumeInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceDefaultGw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceDefaultGwPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceDevId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceDevIdNotification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceDhcpRaGiaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceDhcp6RaLinkaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceDhgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceDigitalSignatureAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceDnsMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceDpd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceDpdRetrycount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceDpdRetryinterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenVpnIpsecPhase1InterfaceEap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceEapCertAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceEapExcludePeergrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceEapIdentity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceEmsSnCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceEncapLocalGw4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceEncapLocalGw6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceEncapRemoteGw4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceEncapRemoteGw6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceEncapsulation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceEncapsulationAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceEnforceUniqueId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceEsn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceExchangeFgtDeviceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceExchangeInterfaceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceExchangeIpAddr4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceExchangeIpAddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceFallbackTcpThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceFecBase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceFecCodec(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceFecEgress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceFecHealthCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceFecIngress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceFecMappingProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceFecReceiveTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceFecRedundant(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceFecSendTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceFgspSync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceForticlientEnforcement(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceFortinetEsp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceFragmentation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceFragmentationMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceGroupAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceHaSyncEspSeqno(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIdleTimeoutinterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIkeVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceInboundDscpCopy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIncludeLocalLan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceInternalDomainList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceIpDelayInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpFragmentation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpsecTunnelSlot(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv4DnsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv4DnsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv4DnsServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv4EndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv4ExcludeRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := i["end-ip"]; ok {
			v := flattenVpnIpsecPhase1InterfaceIpv4ExcludeRangeEndIp(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "VpnIpsecPhase1Interface-Ipv4ExcludeRange-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenVpnIpsecPhase1InterfaceIpv4ExcludeRangeId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "VpnIpsecPhase1Interface-Ipv4ExcludeRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenVpnIpsecPhase1InterfaceIpv4ExcludeRangeStartIp(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "VpnIpsecPhase1Interface-Ipv4ExcludeRange-StartIp")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenVpnIpsecPhase1InterfaceIpv4ExcludeRangeEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv4ExcludeRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv4ExcludeRangeStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv4Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceIpv4Netmask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv4SplitExclude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceIpv4SplitInclude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceIpv4StartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv4WinsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv4WinsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv6AutoLinklocal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv6DnsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv6DnsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv6DnsServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv6EndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv6ExcludeRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := i["end-ip"]; ok {
			v := flattenVpnIpsecPhase1InterfaceIpv6ExcludeRangeEndIp(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "VpnIpsecPhase1Interface-Ipv6ExcludeRange-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenVpnIpsecPhase1InterfaceIpv6ExcludeRangeId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "VpnIpsecPhase1Interface-Ipv6ExcludeRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenVpnIpsecPhase1InterfaceIpv6ExcludeRangeStartIp(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "VpnIpsecPhase1Interface-Ipv6ExcludeRange-StartIp")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenVpnIpsecPhase1InterfaceIpv6ExcludeRangeEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv6ExcludeRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv6ExcludeRangeStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceIpv6Prefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv6SplitExclude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceIpv6SplitInclude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceIpv6StartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceKeepalive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceKeylife(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceKms(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceLinkCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceLocalGw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceLocalGw6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceLocalid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceLocalidType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceLoopbackAsymroute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceMeshSelectorType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceModeCfg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceModeCfgAllowClientSelector(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceMonitorHoldDownDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceMonitorHoldDownTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceMonitorHoldDownType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceMonitorHoldDownWeekday(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceMonitorMin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceNattraversal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceNegotiateTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceNetDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceNetworkId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceNetworkOverlay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceNpuOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfacePacketRedistribution(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfacePassiveMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfacePeer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfacePeergrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfacePeerid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfacePeertype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfacePpk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfacePpkIdentity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfacePriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceProposal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceQkd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceQkdProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceReauth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceRekey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceRemoteGw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceRemoteGwCountry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceRemoteGwEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceRemoteGwMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceRemoteGwStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceRemoteGwSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceRemoteGwZtnaTags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceRemoteGw6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceRemoteGw6Country(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceRemoteGw6EndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceRemoteGw6Match(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceRemoteGw6StartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceRemoteGw6Subnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceRemotegwDdns(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceRsaSignatureFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceRsaSignatureHashOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceSavePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceSendCertChain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceSharedIdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceSignatureHashAlg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceSplitIncludeService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceSuiteB(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceTransitGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceTunnelSearch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceTransport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceUnitySupport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceUsrgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InterfaceVni(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceWizardType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceXauthtype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnIpsecPhase1Interface(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("acct_verify", flattenVpnIpsecPhase1InterfaceAcctVerify(o["acct-verify"], d, "acct_verify")); err != nil {
		if vv, ok := fortiAPIPatch(o["acct-verify"], "VpnIpsecPhase1Interface-AcctVerify"); ok {
			if err = d.Set("acct_verify", vv); err != nil {
				return fmt.Errorf("Error reading acct_verify: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading acct_verify: %v", err)
		}
	}

	if err = d.Set("add_gw_route", flattenVpnIpsecPhase1InterfaceAddGwRoute(o["add-gw-route"], d, "add_gw_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["add-gw-route"], "VpnIpsecPhase1Interface-AddGwRoute"); ok {
			if err = d.Set("add_gw_route", vv); err != nil {
				return fmt.Errorf("Error reading add_gw_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading add_gw_route: %v", err)
		}
	}

	if err = d.Set("add_route", flattenVpnIpsecPhase1InterfaceAddRoute(o["add-route"], d, "add_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["add-route"], "VpnIpsecPhase1Interface-AddRoute"); ok {
			if err = d.Set("add_route", vv); err != nil {
				return fmt.Errorf("Error reading add_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading add_route: %v", err)
		}
	}

	if err = d.Set("addke1", flattenVpnIpsecPhase1InterfaceAddke1(o["addke1"], d, "addke1")); err != nil {
		if vv, ok := fortiAPIPatch(o["addke1"], "VpnIpsecPhase1Interface-Addke1"); ok {
			if err = d.Set("addke1", vv); err != nil {
				return fmt.Errorf("Error reading addke1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addke1: %v", err)
		}
	}

	if err = d.Set("addke2", flattenVpnIpsecPhase1InterfaceAddke2(o["addke2"], d, "addke2")); err != nil {
		if vv, ok := fortiAPIPatch(o["addke2"], "VpnIpsecPhase1Interface-Addke2"); ok {
			if err = d.Set("addke2", vv); err != nil {
				return fmt.Errorf("Error reading addke2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addke2: %v", err)
		}
	}

	if err = d.Set("addke3", flattenVpnIpsecPhase1InterfaceAddke3(o["addke3"], d, "addke3")); err != nil {
		if vv, ok := fortiAPIPatch(o["addke3"], "VpnIpsecPhase1Interface-Addke3"); ok {
			if err = d.Set("addke3", vv); err != nil {
				return fmt.Errorf("Error reading addke3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addke3: %v", err)
		}
	}

	if err = d.Set("addke4", flattenVpnIpsecPhase1InterfaceAddke4(o["addke4"], d, "addke4")); err != nil {
		if vv, ok := fortiAPIPatch(o["addke4"], "VpnIpsecPhase1Interface-Addke4"); ok {
			if err = d.Set("addke4", vv); err != nil {
				return fmt.Errorf("Error reading addke4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addke4: %v", err)
		}
	}

	if err = d.Set("addke5", flattenVpnIpsecPhase1InterfaceAddke5(o["addke5"], d, "addke5")); err != nil {
		if vv, ok := fortiAPIPatch(o["addke5"], "VpnIpsecPhase1Interface-Addke5"); ok {
			if err = d.Set("addke5", vv); err != nil {
				return fmt.Errorf("Error reading addke5: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addke5: %v", err)
		}
	}

	if err = d.Set("addke6", flattenVpnIpsecPhase1InterfaceAddke6(o["addke6"], d, "addke6")); err != nil {
		if vv, ok := fortiAPIPatch(o["addke6"], "VpnIpsecPhase1Interface-Addke6"); ok {
			if err = d.Set("addke6", vv); err != nil {
				return fmt.Errorf("Error reading addke6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addke6: %v", err)
		}
	}

	if err = d.Set("addke7", flattenVpnIpsecPhase1InterfaceAddke7(o["addke7"], d, "addke7")); err != nil {
		if vv, ok := fortiAPIPatch(o["addke7"], "VpnIpsecPhase1Interface-Addke7"); ok {
			if err = d.Set("addke7", vv); err != nil {
				return fmt.Errorf("Error reading addke7: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addke7: %v", err)
		}
	}

	if err = d.Set("aggregate_member", flattenVpnIpsecPhase1InterfaceAggregateMember(o["aggregate-member"], d, "aggregate_member")); err != nil {
		if vv, ok := fortiAPIPatch(o["aggregate-member"], "VpnIpsecPhase1Interface-AggregateMember"); ok {
			if err = d.Set("aggregate_member", vv); err != nil {
				return fmt.Errorf("Error reading aggregate_member: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aggregate_member: %v", err)
		}
	}

	if err = d.Set("aggregate_weight", flattenVpnIpsecPhase1InterfaceAggregateWeight(o["aggregate-weight"], d, "aggregate_weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["aggregate-weight"], "VpnIpsecPhase1Interface-AggregateWeight"); ok {
			if err = d.Set("aggregate_weight", vv); err != nil {
				return fmt.Errorf("Error reading aggregate_weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aggregate_weight: %v", err)
		}
	}

	if err = d.Set("assign_ip", flattenVpnIpsecPhase1InterfaceAssignIp(o["assign-ip"], d, "assign_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["assign-ip"], "VpnIpsecPhase1Interface-AssignIp"); ok {
			if err = d.Set("assign_ip", vv); err != nil {
				return fmt.Errorf("Error reading assign_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading assign_ip: %v", err)
		}
	}

	if err = d.Set("assign_ip_from", flattenVpnIpsecPhase1InterfaceAssignIpFrom(o["assign-ip-from"], d, "assign_ip_from")); err != nil {
		if vv, ok := fortiAPIPatch(o["assign-ip-from"], "VpnIpsecPhase1Interface-AssignIpFrom"); ok {
			if err = d.Set("assign_ip_from", vv); err != nil {
				return fmt.Errorf("Error reading assign_ip_from: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading assign_ip_from: %v", err)
		}
	}

	if err = d.Set("authmethod", flattenVpnIpsecPhase1InterfaceAuthmethod(o["authmethod"], d, "authmethod")); err != nil {
		if vv, ok := fortiAPIPatch(o["authmethod"], "VpnIpsecPhase1Interface-Authmethod"); ok {
			if err = d.Set("authmethod", vv); err != nil {
				return fmt.Errorf("Error reading authmethod: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authmethod: %v", err)
		}
	}

	if err = d.Set("authmethod_remote", flattenVpnIpsecPhase1InterfaceAuthmethodRemote(o["authmethod-remote"], d, "authmethod_remote")); err != nil {
		if vv, ok := fortiAPIPatch(o["authmethod-remote"], "VpnIpsecPhase1Interface-AuthmethodRemote"); ok {
			if err = d.Set("authmethod_remote", vv); err != nil {
				return fmt.Errorf("Error reading authmethod_remote: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authmethod_remote: %v", err)
		}
	}

	if err = d.Set("authusr", flattenVpnIpsecPhase1InterfaceAuthusr(o["authusr"], d, "authusr")); err != nil {
		if vv, ok := fortiAPIPatch(o["authusr"], "VpnIpsecPhase1Interface-Authusr"); ok {
			if err = d.Set("authusr", vv); err != nil {
				return fmt.Errorf("Error reading authusr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authusr: %v", err)
		}
	}

	if err = d.Set("authusrgrp", flattenVpnIpsecPhase1InterfaceAuthusrgrp(o["authusrgrp"], d, "authusrgrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["authusrgrp"], "VpnIpsecPhase1Interface-Authusrgrp"); ok {
			if err = d.Set("authusrgrp", vv); err != nil {
				return fmt.Errorf("Error reading authusrgrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authusrgrp: %v", err)
		}
	}

	if err = d.Set("auto_discovery_crossover", flattenVpnIpsecPhase1InterfaceAutoDiscoveryCrossover(o["auto-discovery-crossover"], d, "auto_discovery_crossover")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-discovery-crossover"], "VpnIpsecPhase1Interface-AutoDiscoveryCrossover"); ok {
			if err = d.Set("auto_discovery_crossover", vv); err != nil {
				return fmt.Errorf("Error reading auto_discovery_crossover: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_discovery_crossover: %v", err)
		}
	}

	if err = d.Set("auto_discovery_forwarder", flattenVpnIpsecPhase1InterfaceAutoDiscoveryForwarder(o["auto-discovery-forwarder"], d, "auto_discovery_forwarder")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-discovery-forwarder"], "VpnIpsecPhase1Interface-AutoDiscoveryForwarder"); ok {
			if err = d.Set("auto_discovery_forwarder", vv); err != nil {
				return fmt.Errorf("Error reading auto_discovery_forwarder: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_discovery_forwarder: %v", err)
		}
	}

	if err = d.Set("auto_discovery_offer_interval", flattenVpnIpsecPhase1InterfaceAutoDiscoveryOfferInterval(o["auto-discovery-offer-interval"], d, "auto_discovery_offer_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-discovery-offer-interval"], "VpnIpsecPhase1Interface-AutoDiscoveryOfferInterval"); ok {
			if err = d.Set("auto_discovery_offer_interval", vv); err != nil {
				return fmt.Errorf("Error reading auto_discovery_offer_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_discovery_offer_interval: %v", err)
		}
	}

	if err = d.Set("auto_discovery_psk", flattenVpnIpsecPhase1InterfaceAutoDiscoveryPsk(o["auto-discovery-psk"], d, "auto_discovery_psk")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-discovery-psk"], "VpnIpsecPhase1Interface-AutoDiscoveryPsk"); ok {
			if err = d.Set("auto_discovery_psk", vv); err != nil {
				return fmt.Errorf("Error reading auto_discovery_psk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_discovery_psk: %v", err)
		}
	}

	if err = d.Set("auto_discovery_receiver", flattenVpnIpsecPhase1InterfaceAutoDiscoveryReceiver(o["auto-discovery-receiver"], d, "auto_discovery_receiver")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-discovery-receiver"], "VpnIpsecPhase1Interface-AutoDiscoveryReceiver"); ok {
			if err = d.Set("auto_discovery_receiver", vv); err != nil {
				return fmt.Errorf("Error reading auto_discovery_receiver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_discovery_receiver: %v", err)
		}
	}

	if err = d.Set("auto_discovery_sender", flattenVpnIpsecPhase1InterfaceAutoDiscoverySender(o["auto-discovery-sender"], d, "auto_discovery_sender")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-discovery-sender"], "VpnIpsecPhase1Interface-AutoDiscoverySender"); ok {
			if err = d.Set("auto_discovery_sender", vv); err != nil {
				return fmt.Errorf("Error reading auto_discovery_sender: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_discovery_sender: %v", err)
		}
	}

	if err = d.Set("auto_discovery_shortcuts", flattenVpnIpsecPhase1InterfaceAutoDiscoveryShortcuts(o["auto-discovery-shortcuts"], d, "auto_discovery_shortcuts")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-discovery-shortcuts"], "VpnIpsecPhase1Interface-AutoDiscoveryShortcuts"); ok {
			if err = d.Set("auto_discovery_shortcuts", vv); err != nil {
				return fmt.Errorf("Error reading auto_discovery_shortcuts: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_discovery_shortcuts: %v", err)
		}
	}

	if err = d.Set("auto_negotiate", flattenVpnIpsecPhase1InterfaceAutoNegotiate(o["auto-negotiate"], d, "auto_negotiate")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-negotiate"], "VpnIpsecPhase1Interface-AutoNegotiate"); ok {
			if err = d.Set("auto_negotiate", vv); err != nil {
				return fmt.Errorf("Error reading auto_negotiate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_negotiate: %v", err)
		}
	}

	if err = d.Set("auto_transport_threshold", flattenVpnIpsecPhase1InterfaceAutoTransportThreshold(o["auto-transport-threshold"], d, "auto_transport_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-transport-threshold"], "VpnIpsecPhase1Interface-AutoTransportThreshold"); ok {
			if err = d.Set("auto_transport_threshold", vv); err != nil {
				return fmt.Errorf("Error reading auto_transport_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_transport_threshold: %v", err)
		}
	}

	if err = d.Set("azure_ad_autoconnect", flattenVpnIpsecPhase1InterfaceAzureAdAutoconnect(o["azure-ad-autoconnect"], d, "azure_ad_autoconnect")); err != nil {
		if vv, ok := fortiAPIPatch(o["azure-ad-autoconnect"], "VpnIpsecPhase1Interface-AzureAdAutoconnect"); ok {
			if err = d.Set("azure_ad_autoconnect", vv); err != nil {
				return fmt.Errorf("Error reading azure_ad_autoconnect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading azure_ad_autoconnect: %v", err)
		}
	}

	if err = d.Set("backup_gateway", flattenVpnIpsecPhase1InterfaceBackupGateway(o["backup-gateway"], d, "backup_gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["backup-gateway"], "VpnIpsecPhase1Interface-BackupGateway"); ok {
			if err = d.Set("backup_gateway", vv); err != nil {
				return fmt.Errorf("Error reading backup_gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading backup_gateway: %v", err)
		}
	}

	if err = d.Set("banner", flattenVpnIpsecPhase1InterfaceBanner(o["banner"], d, "banner")); err != nil {
		if vv, ok := fortiAPIPatch(o["banner"], "VpnIpsecPhase1Interface-Banner"); ok {
			if err = d.Set("banner", vv); err != nil {
				return fmt.Errorf("Error reading banner: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading banner: %v", err)
		}
	}

	if err = d.Set("cert_id_validation", flattenVpnIpsecPhase1InterfaceCertIdValidation(o["cert-id-validation"], d, "cert_id_validation")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert-id-validation"], "VpnIpsecPhase1Interface-CertIdValidation"); ok {
			if err = d.Set("cert_id_validation", vv); err != nil {
				return fmt.Errorf("Error reading cert_id_validation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert_id_validation: %v", err)
		}
	}

	if err = d.Set("cert_peer_username_strip", flattenVpnIpsecPhase1InterfaceCertPeerUsernameStrip(o["cert-peer-username-strip"], d, "cert_peer_username_strip")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert-peer-username-strip"], "VpnIpsecPhase1Interface-CertPeerUsernameStrip"); ok {
			if err = d.Set("cert_peer_username_strip", vv); err != nil {
				return fmt.Errorf("Error reading cert_peer_username_strip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert_peer_username_strip: %v", err)
		}
	}

	if err = d.Set("cert_peer_username_validation", flattenVpnIpsecPhase1InterfaceCertPeerUsernameValidation(o["cert-peer-username-validation"], d, "cert_peer_username_validation")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert-peer-username-validation"], "VpnIpsecPhase1Interface-CertPeerUsernameValidation"); ok {
			if err = d.Set("cert_peer_username_validation", vv); err != nil {
				return fmt.Errorf("Error reading cert_peer_username_validation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert_peer_username_validation: %v", err)
		}
	}

	if err = d.Set("cert_trust_store", flattenVpnIpsecPhase1InterfaceCertTrustStore(o["cert-trust-store"], d, "cert_trust_store")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert-trust-store"], "VpnIpsecPhase1Interface-CertTrustStore"); ok {
			if err = d.Set("cert_trust_store", vv); err != nil {
				return fmt.Errorf("Error reading cert_trust_store: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert_trust_store: %v", err)
		}
	}

	if err = d.Set("certificate", flattenVpnIpsecPhase1InterfaceCertificate(o["certificate"], d, "certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["certificate"], "VpnIpsecPhase1Interface-Certificate"); ok {
			if err = d.Set("certificate", vv); err != nil {
				return fmt.Errorf("Error reading certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("childless_ike", flattenVpnIpsecPhase1InterfaceChildlessIke(o["childless-ike"], d, "childless_ike")); err != nil {
		if vv, ok := fortiAPIPatch(o["childless-ike"], "VpnIpsecPhase1Interface-ChildlessIke"); ok {
			if err = d.Set("childless_ike", vv); err != nil {
				return fmt.Errorf("Error reading childless_ike: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading childless_ike: %v", err)
		}
	}

	if err = d.Set("client_auto_negotiate", flattenVpnIpsecPhase1InterfaceClientAutoNegotiate(o["client-auto-negotiate"], d, "client_auto_negotiate")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-auto-negotiate"], "VpnIpsecPhase1Interface-ClientAutoNegotiate"); ok {
			if err = d.Set("client_auto_negotiate", vv); err != nil {
				return fmt.Errorf("Error reading client_auto_negotiate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_auto_negotiate: %v", err)
		}
	}

	if err = d.Set("client_keep_alive", flattenVpnIpsecPhase1InterfaceClientKeepAlive(o["client-keep-alive"], d, "client_keep_alive")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-keep-alive"], "VpnIpsecPhase1Interface-ClientKeepAlive"); ok {
			if err = d.Set("client_keep_alive", vv); err != nil {
				return fmt.Errorf("Error reading client_keep_alive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_keep_alive: %v", err)
		}
	}

	if err = d.Set("client_resume", flattenVpnIpsecPhase1InterfaceClientResume(o["client-resume"], d, "client_resume")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-resume"], "VpnIpsecPhase1Interface-ClientResume"); ok {
			if err = d.Set("client_resume", vv); err != nil {
				return fmt.Errorf("Error reading client_resume: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_resume: %v", err)
		}
	}

	if err = d.Set("client_resume_interval", flattenVpnIpsecPhase1InterfaceClientResumeInterval(o["client-resume-interval"], d, "client_resume_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-resume-interval"], "VpnIpsecPhase1Interface-ClientResumeInterval"); ok {
			if err = d.Set("client_resume_interval", vv); err != nil {
				return fmt.Errorf("Error reading client_resume_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_resume_interval: %v", err)
		}
	}

	if err = d.Set("comments", flattenVpnIpsecPhase1InterfaceComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "VpnIpsecPhase1Interface-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("default_gw", flattenVpnIpsecPhase1InterfaceDefaultGw(o["default-gw"], d, "default_gw")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-gw"], "VpnIpsecPhase1Interface-DefaultGw"); ok {
			if err = d.Set("default_gw", vv); err != nil {
				return fmt.Errorf("Error reading default_gw: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_gw: %v", err)
		}
	}

	if err = d.Set("default_gw_priority", flattenVpnIpsecPhase1InterfaceDefaultGwPriority(o["default-gw-priority"], d, "default_gw_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-gw-priority"], "VpnIpsecPhase1Interface-DefaultGwPriority"); ok {
			if err = d.Set("default_gw_priority", vv); err != nil {
				return fmt.Errorf("Error reading default_gw_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_gw_priority: %v", err)
		}
	}

	if err = d.Set("dev_id", flattenVpnIpsecPhase1InterfaceDevId(o["dev-id"], d, "dev_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["dev-id"], "VpnIpsecPhase1Interface-DevId"); ok {
			if err = d.Set("dev_id", vv); err != nil {
				return fmt.Errorf("Error reading dev_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dev_id: %v", err)
		}
	}

	if err = d.Set("dev_id_notification", flattenVpnIpsecPhase1InterfaceDevIdNotification(o["dev-id-notification"], d, "dev_id_notification")); err != nil {
		if vv, ok := fortiAPIPatch(o["dev-id-notification"], "VpnIpsecPhase1Interface-DevIdNotification"); ok {
			if err = d.Set("dev_id_notification", vv); err != nil {
				return fmt.Errorf("Error reading dev_id_notification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dev_id_notification: %v", err)
		}
	}

	if err = d.Set("dhcp_ra_giaddr", flattenVpnIpsecPhase1InterfaceDhcpRaGiaddr(o["dhcp-ra-giaddr"], d, "dhcp_ra_giaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-ra-giaddr"], "VpnIpsecPhase1Interface-DhcpRaGiaddr"); ok {
			if err = d.Set("dhcp_ra_giaddr", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_ra_giaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_ra_giaddr: %v", err)
		}
	}

	if err = d.Set("dhcp6_ra_linkaddr", flattenVpnIpsecPhase1InterfaceDhcp6RaLinkaddr(o["dhcp6-ra-linkaddr"], d, "dhcp6_ra_linkaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp6-ra-linkaddr"], "VpnIpsecPhase1Interface-Dhcp6RaLinkaddr"); ok {
			if err = d.Set("dhcp6_ra_linkaddr", vv); err != nil {
				return fmt.Errorf("Error reading dhcp6_ra_linkaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp6_ra_linkaddr: %v", err)
		}
	}

	if err = d.Set("dhgrp", flattenVpnIpsecPhase1InterfaceDhgrp(o["dhgrp"], d, "dhgrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhgrp"], "VpnIpsecPhase1Interface-Dhgrp"); ok {
			if err = d.Set("dhgrp", vv); err != nil {
				return fmt.Errorf("Error reading dhgrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhgrp: %v", err)
		}
	}

	if err = d.Set("digital_signature_auth", flattenVpnIpsecPhase1InterfaceDigitalSignatureAuth(o["digital-signature-auth"], d, "digital_signature_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["digital-signature-auth"], "VpnIpsecPhase1Interface-DigitalSignatureAuth"); ok {
			if err = d.Set("digital_signature_auth", vv); err != nil {
				return fmt.Errorf("Error reading digital_signature_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading digital_signature_auth: %v", err)
		}
	}

	if err = d.Set("distance", flattenVpnIpsecPhase1InterfaceDistance(o["distance"], d, "distance")); err != nil {
		if vv, ok := fortiAPIPatch(o["distance"], "VpnIpsecPhase1Interface-Distance"); ok {
			if err = d.Set("distance", vv); err != nil {
				return fmt.Errorf("Error reading distance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	if err = d.Set("dns_mode", flattenVpnIpsecPhase1InterfaceDnsMode(o["dns-mode"], d, "dns_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-mode"], "VpnIpsecPhase1Interface-DnsMode"); ok {
			if err = d.Set("dns_mode", vv); err != nil {
				return fmt.Errorf("Error reading dns_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_mode: %v", err)
		}
	}

	if err = d.Set("domain", flattenVpnIpsecPhase1InterfaceDomain(o["domain"], d, "domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain"], "VpnIpsecPhase1Interface-Domain"); ok {
			if err = d.Set("domain", vv); err != nil {
				return fmt.Errorf("Error reading domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("dpd", flattenVpnIpsecPhase1InterfaceDpd(o["dpd"], d, "dpd")); err != nil {
		if vv, ok := fortiAPIPatch(o["dpd"], "VpnIpsecPhase1Interface-Dpd"); ok {
			if err = d.Set("dpd", vv); err != nil {
				return fmt.Errorf("Error reading dpd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dpd: %v", err)
		}
	}

	if err = d.Set("dpd_retrycount", flattenVpnIpsecPhase1InterfaceDpdRetrycount(o["dpd-retrycount"], d, "dpd_retrycount")); err != nil {
		if vv, ok := fortiAPIPatch(o["dpd-retrycount"], "VpnIpsecPhase1Interface-DpdRetrycount"); ok {
			if err = d.Set("dpd_retrycount", vv); err != nil {
				return fmt.Errorf("Error reading dpd_retrycount: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dpd_retrycount: %v", err)
		}
	}

	if err = d.Set("dpd_retryinterval", flattenVpnIpsecPhase1InterfaceDpdRetryinterval(o["dpd-retryinterval"], d, "dpd_retryinterval")); err != nil {
		if vv, ok := fortiAPIPatch(o["dpd-retryinterval"], "VpnIpsecPhase1Interface-DpdRetryinterval"); ok {
			if err = d.Set("dpd_retryinterval", vv); err != nil {
				return fmt.Errorf("Error reading dpd_retryinterval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dpd_retryinterval: %v", err)
		}
	}

	if err = d.Set("eap", flattenVpnIpsecPhase1InterfaceEap(o["eap"], d, "eap")); err != nil {
		if vv, ok := fortiAPIPatch(o["eap"], "VpnIpsecPhase1Interface-Eap"); ok {
			if err = d.Set("eap", vv); err != nil {
				return fmt.Errorf("Error reading eap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eap: %v", err)
		}
	}

	if err = d.Set("eap_cert_auth", flattenVpnIpsecPhase1InterfaceEapCertAuth(o["eap-cert-auth"], d, "eap_cert_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["eap-cert-auth"], "VpnIpsecPhase1Interface-EapCertAuth"); ok {
			if err = d.Set("eap_cert_auth", vv); err != nil {
				return fmt.Errorf("Error reading eap_cert_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eap_cert_auth: %v", err)
		}
	}

	if err = d.Set("eap_exclude_peergrp", flattenVpnIpsecPhase1InterfaceEapExcludePeergrp(o["eap-exclude-peergrp"], d, "eap_exclude_peergrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["eap-exclude-peergrp"], "VpnIpsecPhase1Interface-EapExcludePeergrp"); ok {
			if err = d.Set("eap_exclude_peergrp", vv); err != nil {
				return fmt.Errorf("Error reading eap_exclude_peergrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eap_exclude_peergrp: %v", err)
		}
	}

	if err = d.Set("eap_identity", flattenVpnIpsecPhase1InterfaceEapIdentity(o["eap-identity"], d, "eap_identity")); err != nil {
		if vv, ok := fortiAPIPatch(o["eap-identity"], "VpnIpsecPhase1Interface-EapIdentity"); ok {
			if err = d.Set("eap_identity", vv); err != nil {
				return fmt.Errorf("Error reading eap_identity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eap_identity: %v", err)
		}
	}

	if err = d.Set("ems_sn_check", flattenVpnIpsecPhase1InterfaceEmsSnCheck(o["ems-sn-check"], d, "ems_sn_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["ems-sn-check"], "VpnIpsecPhase1Interface-EmsSnCheck"); ok {
			if err = d.Set("ems_sn_check", vv); err != nil {
				return fmt.Errorf("Error reading ems_sn_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ems_sn_check: %v", err)
		}
	}

	if err = d.Set("encap_local_gw4", flattenVpnIpsecPhase1InterfaceEncapLocalGw4(o["encap-local-gw4"], d, "encap_local_gw4")); err != nil {
		if vv, ok := fortiAPIPatch(o["encap-local-gw4"], "VpnIpsecPhase1Interface-EncapLocalGw4"); ok {
			if err = d.Set("encap_local_gw4", vv); err != nil {
				return fmt.Errorf("Error reading encap_local_gw4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading encap_local_gw4: %v", err)
		}
	}

	if err = d.Set("encap_local_gw6", flattenVpnIpsecPhase1InterfaceEncapLocalGw6(o["encap-local-gw6"], d, "encap_local_gw6")); err != nil {
		if vv, ok := fortiAPIPatch(o["encap-local-gw6"], "VpnIpsecPhase1Interface-EncapLocalGw6"); ok {
			if err = d.Set("encap_local_gw6", vv); err != nil {
				return fmt.Errorf("Error reading encap_local_gw6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading encap_local_gw6: %v", err)
		}
	}

	if err = d.Set("encap_remote_gw4", flattenVpnIpsecPhase1InterfaceEncapRemoteGw4(o["encap-remote-gw4"], d, "encap_remote_gw4")); err != nil {
		if vv, ok := fortiAPIPatch(o["encap-remote-gw4"], "VpnIpsecPhase1Interface-EncapRemoteGw4"); ok {
			if err = d.Set("encap_remote_gw4", vv); err != nil {
				return fmt.Errorf("Error reading encap_remote_gw4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading encap_remote_gw4: %v", err)
		}
	}

	if err = d.Set("encap_remote_gw6", flattenVpnIpsecPhase1InterfaceEncapRemoteGw6(o["encap-remote-gw6"], d, "encap_remote_gw6")); err != nil {
		if vv, ok := fortiAPIPatch(o["encap-remote-gw6"], "VpnIpsecPhase1Interface-EncapRemoteGw6"); ok {
			if err = d.Set("encap_remote_gw6", vv); err != nil {
				return fmt.Errorf("Error reading encap_remote_gw6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading encap_remote_gw6: %v", err)
		}
	}

	if err = d.Set("encapsulation", flattenVpnIpsecPhase1InterfaceEncapsulation(o["encapsulation"], d, "encapsulation")); err != nil {
		if vv, ok := fortiAPIPatch(o["encapsulation"], "VpnIpsecPhase1Interface-Encapsulation"); ok {
			if err = d.Set("encapsulation", vv); err != nil {
				return fmt.Errorf("Error reading encapsulation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading encapsulation: %v", err)
		}
	}

	if err = d.Set("encapsulation_address", flattenVpnIpsecPhase1InterfaceEncapsulationAddress(o["encapsulation-address"], d, "encapsulation_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["encapsulation-address"], "VpnIpsecPhase1Interface-EncapsulationAddress"); ok {
			if err = d.Set("encapsulation_address", vv); err != nil {
				return fmt.Errorf("Error reading encapsulation_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading encapsulation_address: %v", err)
		}
	}

	if err = d.Set("enforce_unique_id", flattenVpnIpsecPhase1InterfaceEnforceUniqueId(o["enforce-unique-id"], d, "enforce_unique_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["enforce-unique-id"], "VpnIpsecPhase1Interface-EnforceUniqueId"); ok {
			if err = d.Set("enforce_unique_id", vv); err != nil {
				return fmt.Errorf("Error reading enforce_unique_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enforce_unique_id: %v", err)
		}
	}

	if err = d.Set("esn", flattenVpnIpsecPhase1InterfaceEsn(o["esn"], d, "esn")); err != nil {
		if vv, ok := fortiAPIPatch(o["esn"], "VpnIpsecPhase1Interface-Esn"); ok {
			if err = d.Set("esn", vv); err != nil {
				return fmt.Errorf("Error reading esn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading esn: %v", err)
		}
	}

	if err = d.Set("exchange_fgt_device_id", flattenVpnIpsecPhase1InterfaceExchangeFgtDeviceId(o["exchange-fgt-device-id"], d, "exchange_fgt_device_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["exchange-fgt-device-id"], "VpnIpsecPhase1Interface-ExchangeFgtDeviceId"); ok {
			if err = d.Set("exchange_fgt_device_id", vv); err != nil {
				return fmt.Errorf("Error reading exchange_fgt_device_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exchange_fgt_device_id: %v", err)
		}
	}

	if err = d.Set("exchange_interface_ip", flattenVpnIpsecPhase1InterfaceExchangeInterfaceIp(o["exchange-interface-ip"], d, "exchange_interface_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["exchange-interface-ip"], "VpnIpsecPhase1Interface-ExchangeInterfaceIp"); ok {
			if err = d.Set("exchange_interface_ip", vv); err != nil {
				return fmt.Errorf("Error reading exchange_interface_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exchange_interface_ip: %v", err)
		}
	}

	if err = d.Set("exchange_ip_addr4", flattenVpnIpsecPhase1InterfaceExchangeIpAddr4(o["exchange-ip-addr4"], d, "exchange_ip_addr4")); err != nil {
		if vv, ok := fortiAPIPatch(o["exchange-ip-addr4"], "VpnIpsecPhase1Interface-ExchangeIpAddr4"); ok {
			if err = d.Set("exchange_ip_addr4", vv); err != nil {
				return fmt.Errorf("Error reading exchange_ip_addr4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exchange_ip_addr4: %v", err)
		}
	}

	if err = d.Set("exchange_ip_addr6", flattenVpnIpsecPhase1InterfaceExchangeIpAddr6(o["exchange-ip-addr6"], d, "exchange_ip_addr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["exchange-ip-addr6"], "VpnIpsecPhase1Interface-ExchangeIpAddr6"); ok {
			if err = d.Set("exchange_ip_addr6", vv); err != nil {
				return fmt.Errorf("Error reading exchange_ip_addr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exchange_ip_addr6: %v", err)
		}
	}

	if err = d.Set("fallback_tcp_threshold", flattenVpnIpsecPhase1InterfaceFallbackTcpThreshold(o["fallback-tcp-threshold"], d, "fallback_tcp_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["fallback-tcp-threshold"], "VpnIpsecPhase1Interface-FallbackTcpThreshold"); ok {
			if err = d.Set("fallback_tcp_threshold", vv); err != nil {
				return fmt.Errorf("Error reading fallback_tcp_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fallback_tcp_threshold: %v", err)
		}
	}

	if err = d.Set("fec_base", flattenVpnIpsecPhase1InterfaceFecBase(o["fec-base"], d, "fec_base")); err != nil {
		if vv, ok := fortiAPIPatch(o["fec-base"], "VpnIpsecPhase1Interface-FecBase"); ok {
			if err = d.Set("fec_base", vv); err != nil {
				return fmt.Errorf("Error reading fec_base: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fec_base: %v", err)
		}
	}

	if err = d.Set("fec_codec", flattenVpnIpsecPhase1InterfaceFecCodec(o["fec-codec"], d, "fec_codec")); err != nil {
		if vv, ok := fortiAPIPatch(o["fec-codec"], "VpnIpsecPhase1Interface-FecCodec"); ok {
			if err = d.Set("fec_codec", vv); err != nil {
				return fmt.Errorf("Error reading fec_codec: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fec_codec: %v", err)
		}
	}

	if err = d.Set("fec_egress", flattenVpnIpsecPhase1InterfaceFecEgress(o["fec-egress"], d, "fec_egress")); err != nil {
		if vv, ok := fortiAPIPatch(o["fec-egress"], "VpnIpsecPhase1Interface-FecEgress"); ok {
			if err = d.Set("fec_egress", vv); err != nil {
				return fmt.Errorf("Error reading fec_egress: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fec_egress: %v", err)
		}
	}

	if err = d.Set("fec_health_check", flattenVpnIpsecPhase1InterfaceFecHealthCheck(o["fec-health-check"], d, "fec_health_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["fec-health-check"], "VpnIpsecPhase1Interface-FecHealthCheck"); ok {
			if err = d.Set("fec_health_check", vv); err != nil {
				return fmt.Errorf("Error reading fec_health_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fec_health_check: %v", err)
		}
	}

	if err = d.Set("fec_ingress", flattenVpnIpsecPhase1InterfaceFecIngress(o["fec-ingress"], d, "fec_ingress")); err != nil {
		if vv, ok := fortiAPIPatch(o["fec-ingress"], "VpnIpsecPhase1Interface-FecIngress"); ok {
			if err = d.Set("fec_ingress", vv); err != nil {
				return fmt.Errorf("Error reading fec_ingress: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fec_ingress: %v", err)
		}
	}

	if err = d.Set("fec_mapping_profile", flattenVpnIpsecPhase1InterfaceFecMappingProfile(o["fec-mapping-profile"], d, "fec_mapping_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["fec-mapping-profile"], "VpnIpsecPhase1Interface-FecMappingProfile"); ok {
			if err = d.Set("fec_mapping_profile", vv); err != nil {
				return fmt.Errorf("Error reading fec_mapping_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fec_mapping_profile: %v", err)
		}
	}

	if err = d.Set("fec_receive_timeout", flattenVpnIpsecPhase1InterfaceFecReceiveTimeout(o["fec-receive-timeout"], d, "fec_receive_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["fec-receive-timeout"], "VpnIpsecPhase1Interface-FecReceiveTimeout"); ok {
			if err = d.Set("fec_receive_timeout", vv); err != nil {
				return fmt.Errorf("Error reading fec_receive_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fec_receive_timeout: %v", err)
		}
	}

	if err = d.Set("fec_redundant", flattenVpnIpsecPhase1InterfaceFecRedundant(o["fec-redundant"], d, "fec_redundant")); err != nil {
		if vv, ok := fortiAPIPatch(o["fec-redundant"], "VpnIpsecPhase1Interface-FecRedundant"); ok {
			if err = d.Set("fec_redundant", vv); err != nil {
				return fmt.Errorf("Error reading fec_redundant: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fec_redundant: %v", err)
		}
	}

	if err = d.Set("fec_send_timeout", flattenVpnIpsecPhase1InterfaceFecSendTimeout(o["fec-send-timeout"], d, "fec_send_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["fec-send-timeout"], "VpnIpsecPhase1Interface-FecSendTimeout"); ok {
			if err = d.Set("fec_send_timeout", vv); err != nil {
				return fmt.Errorf("Error reading fec_send_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fec_send_timeout: %v", err)
		}
	}

	if err = d.Set("fgsp_sync", flattenVpnIpsecPhase1InterfaceFgspSync(o["fgsp-sync"], d, "fgsp_sync")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgsp-sync"], "VpnIpsecPhase1Interface-FgspSync"); ok {
			if err = d.Set("fgsp_sync", vv); err != nil {
				return fmt.Errorf("Error reading fgsp_sync: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgsp_sync: %v", err)
		}
	}

	if err = d.Set("forticlient_enforcement", flattenVpnIpsecPhase1InterfaceForticlientEnforcement(o["forticlient-enforcement"], d, "forticlient_enforcement")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticlient-enforcement"], "VpnIpsecPhase1Interface-ForticlientEnforcement"); ok {
			if err = d.Set("forticlient_enforcement", vv); err != nil {
				return fmt.Errorf("Error reading forticlient_enforcement: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticlient_enforcement: %v", err)
		}
	}

	if err = d.Set("fortinet_esp", flattenVpnIpsecPhase1InterfaceFortinetEsp(o["fortinet-esp"], d, "fortinet_esp")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortinet-esp"], "VpnIpsecPhase1Interface-FortinetEsp"); ok {
			if err = d.Set("fortinet_esp", vv); err != nil {
				return fmt.Errorf("Error reading fortinet_esp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortinet_esp: %v", err)
		}
	}

	if err = d.Set("fragmentation", flattenVpnIpsecPhase1InterfaceFragmentation(o["fragmentation"], d, "fragmentation")); err != nil {
		if vv, ok := fortiAPIPatch(o["fragmentation"], "VpnIpsecPhase1Interface-Fragmentation"); ok {
			if err = d.Set("fragmentation", vv); err != nil {
				return fmt.Errorf("Error reading fragmentation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fragmentation: %v", err)
		}
	}

	if err = d.Set("fragmentation_mtu", flattenVpnIpsecPhase1InterfaceFragmentationMtu(o["fragmentation-mtu"], d, "fragmentation_mtu")); err != nil {
		if vv, ok := fortiAPIPatch(o["fragmentation-mtu"], "VpnIpsecPhase1Interface-FragmentationMtu"); ok {
			if err = d.Set("fragmentation_mtu", vv); err != nil {
				return fmt.Errorf("Error reading fragmentation_mtu: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fragmentation_mtu: %v", err)
		}
	}

	if err = d.Set("group_authentication", flattenVpnIpsecPhase1InterfaceGroupAuthentication(o["group-authentication"], d, "group_authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-authentication"], "VpnIpsecPhase1Interface-GroupAuthentication"); ok {
			if err = d.Set("group_authentication", vv); err != nil {
				return fmt.Errorf("Error reading group_authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_authentication: %v", err)
		}
	}

	if err = d.Set("ha_sync_esp_seqno", flattenVpnIpsecPhase1InterfaceHaSyncEspSeqno(o["ha-sync-esp-seqno"], d, "ha_sync_esp_seqno")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-sync-esp-seqno"], "VpnIpsecPhase1Interface-HaSyncEspSeqno"); ok {
			if err = d.Set("ha_sync_esp_seqno", vv); err != nil {
				return fmt.Errorf("Error reading ha_sync_esp_seqno: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_sync_esp_seqno: %v", err)
		}
	}

	if err = d.Set("idle_timeout", flattenVpnIpsecPhase1InterfaceIdleTimeout(o["idle-timeout"], d, "idle_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["idle-timeout"], "VpnIpsecPhase1Interface-IdleTimeout"); ok {
			if err = d.Set("idle_timeout", vv); err != nil {
				return fmt.Errorf("Error reading idle_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idle_timeout: %v", err)
		}
	}

	if err = d.Set("idle_timeoutinterval", flattenVpnIpsecPhase1InterfaceIdleTimeoutinterval(o["idle-timeoutinterval"], d, "idle_timeoutinterval")); err != nil {
		if vv, ok := fortiAPIPatch(o["idle-timeoutinterval"], "VpnIpsecPhase1Interface-IdleTimeoutinterval"); ok {
			if err = d.Set("idle_timeoutinterval", vv); err != nil {
				return fmt.Errorf("Error reading idle_timeoutinterval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idle_timeoutinterval: %v", err)
		}
	}

	if err = d.Set("ike_version", flattenVpnIpsecPhase1InterfaceIkeVersion(o["ike-version"], d, "ike_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike-version"], "VpnIpsecPhase1Interface-IkeVersion"); ok {
			if err = d.Set("ike_version", vv); err != nil {
				return fmt.Errorf("Error reading ike_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike_version: %v", err)
		}
	}

	if err = d.Set("inbound_dscp_copy", flattenVpnIpsecPhase1InterfaceInboundDscpCopy(o["inbound-dscp-copy"], d, "inbound_dscp_copy")); err != nil {
		if vv, ok := fortiAPIPatch(o["inbound-dscp-copy"], "VpnIpsecPhase1Interface-InboundDscpCopy"); ok {
			if err = d.Set("inbound_dscp_copy", vv); err != nil {
				return fmt.Errorf("Error reading inbound_dscp_copy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inbound_dscp_copy: %v", err)
		}
	}

	if err = d.Set("include_local_lan", flattenVpnIpsecPhase1InterfaceIncludeLocalLan(o["include-local-lan"], d, "include_local_lan")); err != nil {
		if vv, ok := fortiAPIPatch(o["include-local-lan"], "VpnIpsecPhase1Interface-IncludeLocalLan"); ok {
			if err = d.Set("include_local_lan", vv); err != nil {
				return fmt.Errorf("Error reading include_local_lan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading include_local_lan: %v", err)
		}
	}

	if err = d.Set("interface", flattenVpnIpsecPhase1InterfaceInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "VpnIpsecPhase1Interface-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("internal_domain_list", flattenVpnIpsecPhase1InterfaceInternalDomainList(o["internal-domain-list"], d, "internal_domain_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["internal-domain-list"], "VpnIpsecPhase1Interface-InternalDomainList"); ok {
			if err = d.Set("internal_domain_list", vv); err != nil {
				return fmt.Errorf("Error reading internal_domain_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internal_domain_list: %v", err)
		}
	}

	if err = d.Set("ip_delay_interval", flattenVpnIpsecPhase1InterfaceIpDelayInterval(o["ip-delay-interval"], d, "ip_delay_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-delay-interval"], "VpnIpsecPhase1Interface-IpDelayInterval"); ok {
			if err = d.Set("ip_delay_interval", vv); err != nil {
				return fmt.Errorf("Error reading ip_delay_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_delay_interval: %v", err)
		}
	}

	if err = d.Set("ip_fragmentation", flattenVpnIpsecPhase1InterfaceIpFragmentation(o["ip-fragmentation"], d, "ip_fragmentation")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-fragmentation"], "VpnIpsecPhase1Interface-IpFragmentation"); ok {
			if err = d.Set("ip_fragmentation", vv); err != nil {
				return fmt.Errorf("Error reading ip_fragmentation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_fragmentation: %v", err)
		}
	}

	if err = d.Set("ip_version", flattenVpnIpsecPhase1InterfaceIpVersion(o["ip-version"], d, "ip_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-version"], "VpnIpsecPhase1Interface-IpVersion"); ok {
			if err = d.Set("ip_version", vv); err != nil {
				return fmt.Errorf("Error reading ip_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_version: %v", err)
		}
	}

	if err = d.Set("ipsec_tunnel_slot", flattenVpnIpsecPhase1InterfaceIpsecTunnelSlot(o["ipsec-tunnel-slot"], d, "ipsec_tunnel_slot")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-tunnel-slot"], "VpnIpsecPhase1Interface-IpsecTunnelSlot"); ok {
			if err = d.Set("ipsec_tunnel_slot", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_tunnel_slot: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_tunnel_slot: %v", err)
		}
	}

	if err = d.Set("ipv4_dns_server1", flattenVpnIpsecPhase1InterfaceIpv4DnsServer1(o["ipv4-dns-server1"], d, "ipv4_dns_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-dns-server1"], "VpnIpsecPhase1Interface-Ipv4DnsServer1"); ok {
			if err = d.Set("ipv4_dns_server1", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_dns_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_dns_server1: %v", err)
		}
	}

	if err = d.Set("ipv4_dns_server2", flattenVpnIpsecPhase1InterfaceIpv4DnsServer2(o["ipv4-dns-server2"], d, "ipv4_dns_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-dns-server2"], "VpnIpsecPhase1Interface-Ipv4DnsServer2"); ok {
			if err = d.Set("ipv4_dns_server2", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_dns_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_dns_server2: %v", err)
		}
	}

	if err = d.Set("ipv4_dns_server3", flattenVpnIpsecPhase1InterfaceIpv4DnsServer3(o["ipv4-dns-server3"], d, "ipv4_dns_server3")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-dns-server3"], "VpnIpsecPhase1Interface-Ipv4DnsServer3"); ok {
			if err = d.Set("ipv4_dns_server3", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_dns_server3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_dns_server3: %v", err)
		}
	}

	if err = d.Set("ipv4_end_ip", flattenVpnIpsecPhase1InterfaceIpv4EndIp(o["ipv4-end-ip"], d, "ipv4_end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-end-ip"], "VpnIpsecPhase1Interface-Ipv4EndIp"); ok {
			if err = d.Set("ipv4_end_ip", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_end_ip: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ipv4_exclude_range", flattenVpnIpsecPhase1InterfaceIpv4ExcludeRange(o["ipv4-exclude-range"], d, "ipv4_exclude_range")); err != nil {
			if vv, ok := fortiAPIPatch(o["ipv4-exclude-range"], "VpnIpsecPhase1Interface-Ipv4ExcludeRange"); ok {
				if err = d.Set("ipv4_exclude_range", vv); err != nil {
					return fmt.Errorf("Error reading ipv4_exclude_range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ipv4_exclude_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ipv4_exclude_range"); ok {
			if err = d.Set("ipv4_exclude_range", flattenVpnIpsecPhase1InterfaceIpv4ExcludeRange(o["ipv4-exclude-range"], d, "ipv4_exclude_range")); err != nil {
				if vv, ok := fortiAPIPatch(o["ipv4-exclude-range"], "VpnIpsecPhase1Interface-Ipv4ExcludeRange"); ok {
					if err = d.Set("ipv4_exclude_range", vv); err != nil {
						return fmt.Errorf("Error reading ipv4_exclude_range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ipv4_exclude_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("ipv4_name", flattenVpnIpsecPhase1InterfaceIpv4Name(o["ipv4-name"], d, "ipv4_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-name"], "VpnIpsecPhase1Interface-Ipv4Name"); ok {
			if err = d.Set("ipv4_name", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_name: %v", err)
		}
	}

	if err = d.Set("ipv4_netmask", flattenVpnIpsecPhase1InterfaceIpv4Netmask(o["ipv4-netmask"], d, "ipv4_netmask")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-netmask"], "VpnIpsecPhase1Interface-Ipv4Netmask"); ok {
			if err = d.Set("ipv4_netmask", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_netmask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_netmask: %v", err)
		}
	}

	if err = d.Set("ipv4_split_exclude", flattenVpnIpsecPhase1InterfaceIpv4SplitExclude(o["ipv4-split-exclude"], d, "ipv4_split_exclude")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-split-exclude"], "VpnIpsecPhase1Interface-Ipv4SplitExclude"); ok {
			if err = d.Set("ipv4_split_exclude", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_split_exclude: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_split_exclude: %v", err)
		}
	}

	if err = d.Set("ipv4_split_include", flattenVpnIpsecPhase1InterfaceIpv4SplitInclude(o["ipv4-split-include"], d, "ipv4_split_include")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-split-include"], "VpnIpsecPhase1Interface-Ipv4SplitInclude"); ok {
			if err = d.Set("ipv4_split_include", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_split_include: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_split_include: %v", err)
		}
	}

	if err = d.Set("ipv4_start_ip", flattenVpnIpsecPhase1InterfaceIpv4StartIp(o["ipv4-start-ip"], d, "ipv4_start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-start-ip"], "VpnIpsecPhase1Interface-Ipv4StartIp"); ok {
			if err = d.Set("ipv4_start_ip", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_start_ip: %v", err)
		}
	}

	if err = d.Set("ipv4_wins_server1", flattenVpnIpsecPhase1InterfaceIpv4WinsServer1(o["ipv4-wins-server1"], d, "ipv4_wins_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-wins-server1"], "VpnIpsecPhase1Interface-Ipv4WinsServer1"); ok {
			if err = d.Set("ipv4_wins_server1", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_wins_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_wins_server1: %v", err)
		}
	}

	if err = d.Set("ipv4_wins_server2", flattenVpnIpsecPhase1InterfaceIpv4WinsServer2(o["ipv4-wins-server2"], d, "ipv4_wins_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-wins-server2"], "VpnIpsecPhase1Interface-Ipv4WinsServer2"); ok {
			if err = d.Set("ipv4_wins_server2", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_wins_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_wins_server2: %v", err)
		}
	}

	if err = d.Set("ipv6_auto_linklocal", flattenVpnIpsecPhase1InterfaceIpv6AutoLinklocal(o["ipv6-auto-linklocal"], d, "ipv6_auto_linklocal")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-auto-linklocal"], "VpnIpsecPhase1Interface-Ipv6AutoLinklocal"); ok {
			if err = d.Set("ipv6_auto_linklocal", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_auto_linklocal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_auto_linklocal: %v", err)
		}
	}

	if err = d.Set("ipv6_dns_server1", flattenVpnIpsecPhase1InterfaceIpv6DnsServer1(o["ipv6-dns-server1"], d, "ipv6_dns_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-dns-server1"], "VpnIpsecPhase1Interface-Ipv6DnsServer1"); ok {
			if err = d.Set("ipv6_dns_server1", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_dns_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_dns_server1: %v", err)
		}
	}

	if err = d.Set("ipv6_dns_server2", flattenVpnIpsecPhase1InterfaceIpv6DnsServer2(o["ipv6-dns-server2"], d, "ipv6_dns_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-dns-server2"], "VpnIpsecPhase1Interface-Ipv6DnsServer2"); ok {
			if err = d.Set("ipv6_dns_server2", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_dns_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_dns_server2: %v", err)
		}
	}

	if err = d.Set("ipv6_dns_server3", flattenVpnIpsecPhase1InterfaceIpv6DnsServer3(o["ipv6-dns-server3"], d, "ipv6_dns_server3")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-dns-server3"], "VpnIpsecPhase1Interface-Ipv6DnsServer3"); ok {
			if err = d.Set("ipv6_dns_server3", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_dns_server3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_dns_server3: %v", err)
		}
	}

	if err = d.Set("ipv6_end_ip", flattenVpnIpsecPhase1InterfaceIpv6EndIp(o["ipv6-end-ip"], d, "ipv6_end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-end-ip"], "VpnIpsecPhase1Interface-Ipv6EndIp"); ok {
			if err = d.Set("ipv6_end_ip", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_end_ip: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ipv6_exclude_range", flattenVpnIpsecPhase1InterfaceIpv6ExcludeRange(o["ipv6-exclude-range"], d, "ipv6_exclude_range")); err != nil {
			if vv, ok := fortiAPIPatch(o["ipv6-exclude-range"], "VpnIpsecPhase1Interface-Ipv6ExcludeRange"); ok {
				if err = d.Set("ipv6_exclude_range", vv); err != nil {
					return fmt.Errorf("Error reading ipv6_exclude_range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ipv6_exclude_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ipv6_exclude_range"); ok {
			if err = d.Set("ipv6_exclude_range", flattenVpnIpsecPhase1InterfaceIpv6ExcludeRange(o["ipv6-exclude-range"], d, "ipv6_exclude_range")); err != nil {
				if vv, ok := fortiAPIPatch(o["ipv6-exclude-range"], "VpnIpsecPhase1Interface-Ipv6ExcludeRange"); ok {
					if err = d.Set("ipv6_exclude_range", vv); err != nil {
						return fmt.Errorf("Error reading ipv6_exclude_range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ipv6_exclude_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("ipv6_name", flattenVpnIpsecPhase1InterfaceIpv6Name(o["ipv6-name"], d, "ipv6_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-name"], "VpnIpsecPhase1Interface-Ipv6Name"); ok {
			if err = d.Set("ipv6_name", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_name: %v", err)
		}
	}

	if err = d.Set("ipv6_prefix", flattenVpnIpsecPhase1InterfaceIpv6Prefix(o["ipv6-prefix"], d, "ipv6_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-prefix"], "VpnIpsecPhase1Interface-Ipv6Prefix"); ok {
			if err = d.Set("ipv6_prefix", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_prefix: %v", err)
		}
	}

	if err = d.Set("ipv6_split_exclude", flattenVpnIpsecPhase1InterfaceIpv6SplitExclude(o["ipv6-split-exclude"], d, "ipv6_split_exclude")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-split-exclude"], "VpnIpsecPhase1Interface-Ipv6SplitExclude"); ok {
			if err = d.Set("ipv6_split_exclude", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_split_exclude: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_split_exclude: %v", err)
		}
	}

	if err = d.Set("ipv6_split_include", flattenVpnIpsecPhase1InterfaceIpv6SplitInclude(o["ipv6-split-include"], d, "ipv6_split_include")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-split-include"], "VpnIpsecPhase1Interface-Ipv6SplitInclude"); ok {
			if err = d.Set("ipv6_split_include", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_split_include: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_split_include: %v", err)
		}
	}

	if err = d.Set("ipv6_start_ip", flattenVpnIpsecPhase1InterfaceIpv6StartIp(o["ipv6-start-ip"], d, "ipv6_start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-start-ip"], "VpnIpsecPhase1Interface-Ipv6StartIp"); ok {
			if err = d.Set("ipv6_start_ip", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_start_ip: %v", err)
		}
	}

	if err = d.Set("keepalive", flattenVpnIpsecPhase1InterfaceKeepalive(o["keepalive"], d, "keepalive")); err != nil {
		if vv, ok := fortiAPIPatch(o["keepalive"], "VpnIpsecPhase1Interface-Keepalive"); ok {
			if err = d.Set("keepalive", vv); err != nil {
				return fmt.Errorf("Error reading keepalive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keepalive: %v", err)
		}
	}

	if err = d.Set("keylife", flattenVpnIpsecPhase1InterfaceKeylife(o["keylife"], d, "keylife")); err != nil {
		if vv, ok := fortiAPIPatch(o["keylife"], "VpnIpsecPhase1Interface-Keylife"); ok {
			if err = d.Set("keylife", vv); err != nil {
				return fmt.Errorf("Error reading keylife: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keylife: %v", err)
		}
	}

	if err = d.Set("kms", flattenVpnIpsecPhase1InterfaceKms(o["kms"], d, "kms")); err != nil {
		if vv, ok := fortiAPIPatch(o["kms"], "VpnIpsecPhase1Interface-Kms"); ok {
			if err = d.Set("kms", vv); err != nil {
				return fmt.Errorf("Error reading kms: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading kms: %v", err)
		}
	}

	if err = d.Set("link_cost", flattenVpnIpsecPhase1InterfaceLinkCost(o["link-cost"], d, "link_cost")); err != nil {
		if vv, ok := fortiAPIPatch(o["link-cost"], "VpnIpsecPhase1Interface-LinkCost"); ok {
			if err = d.Set("link_cost", vv); err != nil {
				return fmt.Errorf("Error reading link_cost: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading link_cost: %v", err)
		}
	}

	if err = d.Set("local_gw", flattenVpnIpsecPhase1InterfaceLocalGw(o["local-gw"], d, "local_gw")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-gw"], "VpnIpsecPhase1Interface-LocalGw"); ok {
			if err = d.Set("local_gw", vv); err != nil {
				return fmt.Errorf("Error reading local_gw: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_gw: %v", err)
		}
	}

	if err = d.Set("local_gw6", flattenVpnIpsecPhase1InterfaceLocalGw6(o["local-gw6"], d, "local_gw6")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-gw6"], "VpnIpsecPhase1Interface-LocalGw6"); ok {
			if err = d.Set("local_gw6", vv); err != nil {
				return fmt.Errorf("Error reading local_gw6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_gw6: %v", err)
		}
	}

	if err = d.Set("localid", flattenVpnIpsecPhase1InterfaceLocalid(o["localid"], d, "localid")); err != nil {
		if vv, ok := fortiAPIPatch(o["localid"], "VpnIpsecPhase1Interface-Localid"); ok {
			if err = d.Set("localid", vv); err != nil {
				return fmt.Errorf("Error reading localid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading localid: %v", err)
		}
	}

	if err = d.Set("localid_type", flattenVpnIpsecPhase1InterfaceLocalidType(o["localid-type"], d, "localid_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["localid-type"], "VpnIpsecPhase1Interface-LocalidType"); ok {
			if err = d.Set("localid_type", vv); err != nil {
				return fmt.Errorf("Error reading localid_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading localid_type: %v", err)
		}
	}

	if err = d.Set("loopback_asymroute", flattenVpnIpsecPhase1InterfaceLoopbackAsymroute(o["loopback-asymroute"], d, "loopback_asymroute")); err != nil {
		if vv, ok := fortiAPIPatch(o["loopback-asymroute"], "VpnIpsecPhase1Interface-LoopbackAsymroute"); ok {
			if err = d.Set("loopback_asymroute", vv); err != nil {
				return fmt.Errorf("Error reading loopback_asymroute: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading loopback_asymroute: %v", err)
		}
	}

	if err = d.Set("mesh_selector_type", flattenVpnIpsecPhase1InterfaceMeshSelectorType(o["mesh-selector-type"], d, "mesh_selector_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["mesh-selector-type"], "VpnIpsecPhase1Interface-MeshSelectorType"); ok {
			if err = d.Set("mesh_selector_type", vv); err != nil {
				return fmt.Errorf("Error reading mesh_selector_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mesh_selector_type: %v", err)
		}
	}

	if err = d.Set("mode", flattenVpnIpsecPhase1InterfaceMode(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "VpnIpsecPhase1Interface-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("mode_cfg", flattenVpnIpsecPhase1InterfaceModeCfg(o["mode-cfg"], d, "mode_cfg")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode-cfg"], "VpnIpsecPhase1Interface-ModeCfg"); ok {
			if err = d.Set("mode_cfg", vv); err != nil {
				return fmt.Errorf("Error reading mode_cfg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode_cfg: %v", err)
		}
	}

	if err = d.Set("mode_cfg_allow_client_selector", flattenVpnIpsecPhase1InterfaceModeCfgAllowClientSelector(o["mode-cfg-allow-client-selector"], d, "mode_cfg_allow_client_selector")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode-cfg-allow-client-selector"], "VpnIpsecPhase1Interface-ModeCfgAllowClientSelector"); ok {
			if err = d.Set("mode_cfg_allow_client_selector", vv); err != nil {
				return fmt.Errorf("Error reading mode_cfg_allow_client_selector: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode_cfg_allow_client_selector: %v", err)
		}
	}

	if err = d.Set("monitor", flattenVpnIpsecPhase1InterfaceMonitor(o["monitor"], d, "monitor")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor"], "VpnIpsecPhase1Interface-Monitor"); ok {
			if err = d.Set("monitor", vv); err != nil {
				return fmt.Errorf("Error reading monitor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor: %v", err)
		}
	}

	if err = d.Set("monitor_hold_down_delay", flattenVpnIpsecPhase1InterfaceMonitorHoldDownDelay(o["monitor-hold-down-delay"], d, "monitor_hold_down_delay")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor-hold-down-delay"], "VpnIpsecPhase1Interface-MonitorHoldDownDelay"); ok {
			if err = d.Set("monitor_hold_down_delay", vv); err != nil {
				return fmt.Errorf("Error reading monitor_hold_down_delay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor_hold_down_delay: %v", err)
		}
	}

	if err = d.Set("monitor_hold_down_time", flattenVpnIpsecPhase1InterfaceMonitorHoldDownTime(o["monitor-hold-down-time"], d, "monitor_hold_down_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor-hold-down-time"], "VpnIpsecPhase1Interface-MonitorHoldDownTime"); ok {
			if err = d.Set("monitor_hold_down_time", vv); err != nil {
				return fmt.Errorf("Error reading monitor_hold_down_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor_hold_down_time: %v", err)
		}
	}

	if err = d.Set("monitor_hold_down_type", flattenVpnIpsecPhase1InterfaceMonitorHoldDownType(o["monitor-hold-down-type"], d, "monitor_hold_down_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor-hold-down-type"], "VpnIpsecPhase1Interface-MonitorHoldDownType"); ok {
			if err = d.Set("monitor_hold_down_type", vv); err != nil {
				return fmt.Errorf("Error reading monitor_hold_down_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor_hold_down_type: %v", err)
		}
	}

	if err = d.Set("monitor_hold_down_weekday", flattenVpnIpsecPhase1InterfaceMonitorHoldDownWeekday(o["monitor-hold-down-weekday"], d, "monitor_hold_down_weekday")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor-hold-down-weekday"], "VpnIpsecPhase1Interface-MonitorHoldDownWeekday"); ok {
			if err = d.Set("monitor_hold_down_weekday", vv); err != nil {
				return fmt.Errorf("Error reading monitor_hold_down_weekday: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor_hold_down_weekday: %v", err)
		}
	}

	if err = d.Set("monitor_min", flattenVpnIpsecPhase1InterfaceMonitorMin(o["monitor-min"], d, "monitor_min")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor-min"], "VpnIpsecPhase1Interface-MonitorMin"); ok {
			if err = d.Set("monitor_min", vv); err != nil {
				return fmt.Errorf("Error reading monitor_min: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor_min: %v", err)
		}
	}

	if err = d.Set("name", flattenVpnIpsecPhase1InterfaceName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "VpnIpsecPhase1Interface-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("nattraversal", flattenVpnIpsecPhase1InterfaceNattraversal(o["nattraversal"], d, "nattraversal")); err != nil {
		if vv, ok := fortiAPIPatch(o["nattraversal"], "VpnIpsecPhase1Interface-Nattraversal"); ok {
			if err = d.Set("nattraversal", vv); err != nil {
				return fmt.Errorf("Error reading nattraversal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nattraversal: %v", err)
		}
	}

	if err = d.Set("negotiate_timeout", flattenVpnIpsecPhase1InterfaceNegotiateTimeout(o["negotiate-timeout"], d, "negotiate_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["negotiate-timeout"], "VpnIpsecPhase1Interface-NegotiateTimeout"); ok {
			if err = d.Set("negotiate_timeout", vv); err != nil {
				return fmt.Errorf("Error reading negotiate_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading negotiate_timeout: %v", err)
		}
	}

	if err = d.Set("net_device", flattenVpnIpsecPhase1InterfaceNetDevice(o["net-device"], d, "net_device")); err != nil {
		if vv, ok := fortiAPIPatch(o["net-device"], "VpnIpsecPhase1Interface-NetDevice"); ok {
			if err = d.Set("net_device", vv); err != nil {
				return fmt.Errorf("Error reading net_device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading net_device: %v", err)
		}
	}

	if err = d.Set("network_id", flattenVpnIpsecPhase1InterfaceNetworkId(o["network-id"], d, "network_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["network-id"], "VpnIpsecPhase1Interface-NetworkId"); ok {
			if err = d.Set("network_id", vv); err != nil {
				return fmt.Errorf("Error reading network_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading network_id: %v", err)
		}
	}

	if err = d.Set("network_overlay", flattenVpnIpsecPhase1InterfaceNetworkOverlay(o["network-overlay"], d, "network_overlay")); err != nil {
		if vv, ok := fortiAPIPatch(o["network-overlay"], "VpnIpsecPhase1Interface-NetworkOverlay"); ok {
			if err = d.Set("network_overlay", vv); err != nil {
				return fmt.Errorf("Error reading network_overlay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading network_overlay: %v", err)
		}
	}

	if err = d.Set("npu_offload", flattenVpnIpsecPhase1InterfaceNpuOffload(o["npu-offload"], d, "npu_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["npu-offload"], "VpnIpsecPhase1Interface-NpuOffload"); ok {
			if err = d.Set("npu_offload", vv); err != nil {
				return fmt.Errorf("Error reading npu_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading npu_offload: %v", err)
		}
	}

	if err = d.Set("packet_redistribution", flattenVpnIpsecPhase1InterfacePacketRedistribution(o["packet-redistribution"], d, "packet_redistribution")); err != nil {
		if vv, ok := fortiAPIPatch(o["packet-redistribution"], "VpnIpsecPhase1Interface-PacketRedistribution"); ok {
			if err = d.Set("packet_redistribution", vv); err != nil {
				return fmt.Errorf("Error reading packet_redistribution: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packet_redistribution: %v", err)
		}
	}

	if err = d.Set("passive_mode", flattenVpnIpsecPhase1InterfacePassiveMode(o["passive-mode"], d, "passive_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["passive-mode"], "VpnIpsecPhase1Interface-PassiveMode"); ok {
			if err = d.Set("passive_mode", vv); err != nil {
				return fmt.Errorf("Error reading passive_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading passive_mode: %v", err)
		}
	}

	if err = d.Set("peer", flattenVpnIpsecPhase1InterfacePeer(o["peer"], d, "peer")); err != nil {
		if vv, ok := fortiAPIPatch(o["peer"], "VpnIpsecPhase1Interface-Peer"); ok {
			if err = d.Set("peer", vv); err != nil {
				return fmt.Errorf("Error reading peer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peer: %v", err)
		}
	}

	if err = d.Set("peergrp", flattenVpnIpsecPhase1InterfacePeergrp(o["peergrp"], d, "peergrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["peergrp"], "VpnIpsecPhase1Interface-Peergrp"); ok {
			if err = d.Set("peergrp", vv); err != nil {
				return fmt.Errorf("Error reading peergrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peergrp: %v", err)
		}
	}

	if err = d.Set("peerid", flattenVpnIpsecPhase1InterfacePeerid(o["peerid"], d, "peerid")); err != nil {
		if vv, ok := fortiAPIPatch(o["peerid"], "VpnIpsecPhase1Interface-Peerid"); ok {
			if err = d.Set("peerid", vv); err != nil {
				return fmt.Errorf("Error reading peerid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peerid: %v", err)
		}
	}

	if err = d.Set("peertype", flattenVpnIpsecPhase1InterfacePeertype(o["peertype"], d, "peertype")); err != nil {
		if vv, ok := fortiAPIPatch(o["peertype"], "VpnIpsecPhase1Interface-Peertype"); ok {
			if err = d.Set("peertype", vv); err != nil {
				return fmt.Errorf("Error reading peertype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peertype: %v", err)
		}
	}

	if err = d.Set("ppk", flattenVpnIpsecPhase1InterfacePpk(o["ppk"], d, "ppk")); err != nil {
		if vv, ok := fortiAPIPatch(o["ppk"], "VpnIpsecPhase1Interface-Ppk"); ok {
			if err = d.Set("ppk", vv); err != nil {
				return fmt.Errorf("Error reading ppk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ppk: %v", err)
		}
	}

	if err = d.Set("ppk_identity", flattenVpnIpsecPhase1InterfacePpkIdentity(o["ppk-identity"], d, "ppk_identity")); err != nil {
		if vv, ok := fortiAPIPatch(o["ppk-identity"], "VpnIpsecPhase1Interface-PpkIdentity"); ok {
			if err = d.Set("ppk_identity", vv); err != nil {
				return fmt.Errorf("Error reading ppk_identity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ppk_identity: %v", err)
		}
	}

	if err = d.Set("priority", flattenVpnIpsecPhase1InterfacePriority(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "VpnIpsecPhase1Interface-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("proposal", flattenVpnIpsecPhase1InterfaceProposal(o["proposal"], d, "proposal")); err != nil {
		if vv, ok := fortiAPIPatch(o["proposal"], "VpnIpsecPhase1Interface-Proposal"); ok {
			if err = d.Set("proposal", vv); err != nil {
				return fmt.Errorf("Error reading proposal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proposal: %v", err)
		}
	}

	if err = d.Set("qkd", flattenVpnIpsecPhase1InterfaceQkd(o["qkd"], d, "qkd")); err != nil {
		if vv, ok := fortiAPIPatch(o["qkd"], "VpnIpsecPhase1Interface-Qkd"); ok {
			if err = d.Set("qkd", vv); err != nil {
				return fmt.Errorf("Error reading qkd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading qkd: %v", err)
		}
	}

	if err = d.Set("qkd_profile", flattenVpnIpsecPhase1InterfaceQkdProfile(o["qkd-profile"], d, "qkd_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["qkd-profile"], "VpnIpsecPhase1Interface-QkdProfile"); ok {
			if err = d.Set("qkd_profile", vv); err != nil {
				return fmt.Errorf("Error reading qkd_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading qkd_profile: %v", err)
		}
	}

	if err = d.Set("reauth", flattenVpnIpsecPhase1InterfaceReauth(o["reauth"], d, "reauth")); err != nil {
		if vv, ok := fortiAPIPatch(o["reauth"], "VpnIpsecPhase1Interface-Reauth"); ok {
			if err = d.Set("reauth", vv); err != nil {
				return fmt.Errorf("Error reading reauth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reauth: %v", err)
		}
	}

	if err = d.Set("rekey", flattenVpnIpsecPhase1InterfaceRekey(o["rekey"], d, "rekey")); err != nil {
		if vv, ok := fortiAPIPatch(o["rekey"], "VpnIpsecPhase1Interface-Rekey"); ok {
			if err = d.Set("rekey", vv); err != nil {
				return fmt.Errorf("Error reading rekey: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rekey: %v", err)
		}
	}

	if err = d.Set("remote_gw", flattenVpnIpsecPhase1InterfaceRemoteGw(o["remote-gw"], d, "remote_gw")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-gw"], "VpnIpsecPhase1Interface-RemoteGw"); ok {
			if err = d.Set("remote_gw", vv); err != nil {
				return fmt.Errorf("Error reading remote_gw: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_gw: %v", err)
		}
	}

	if err = d.Set("remote_gw_country", flattenVpnIpsecPhase1InterfaceRemoteGwCountry(o["remote-gw-country"], d, "remote_gw_country")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-gw-country"], "VpnIpsecPhase1Interface-RemoteGwCountry"); ok {
			if err = d.Set("remote_gw_country", vv); err != nil {
				return fmt.Errorf("Error reading remote_gw_country: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_gw_country: %v", err)
		}
	}

	if err = d.Set("remote_gw_end_ip", flattenVpnIpsecPhase1InterfaceRemoteGwEndIp(o["remote-gw-end-ip"], d, "remote_gw_end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-gw-end-ip"], "VpnIpsecPhase1Interface-RemoteGwEndIp"); ok {
			if err = d.Set("remote_gw_end_ip", vv); err != nil {
				return fmt.Errorf("Error reading remote_gw_end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_gw_end_ip: %v", err)
		}
	}

	if err = d.Set("remote_gw_match", flattenVpnIpsecPhase1InterfaceRemoteGwMatch(o["remote-gw-match"], d, "remote_gw_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-gw-match"], "VpnIpsecPhase1Interface-RemoteGwMatch"); ok {
			if err = d.Set("remote_gw_match", vv); err != nil {
				return fmt.Errorf("Error reading remote_gw_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_gw_match: %v", err)
		}
	}

	if err = d.Set("remote_gw_start_ip", flattenVpnIpsecPhase1InterfaceRemoteGwStartIp(o["remote-gw-start-ip"], d, "remote_gw_start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-gw-start-ip"], "VpnIpsecPhase1Interface-RemoteGwStartIp"); ok {
			if err = d.Set("remote_gw_start_ip", vv); err != nil {
				return fmt.Errorf("Error reading remote_gw_start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_gw_start_ip: %v", err)
		}
	}

	if err = d.Set("remote_gw_subnet", flattenVpnIpsecPhase1InterfaceRemoteGwSubnet(o["remote-gw-subnet"], d, "remote_gw_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-gw-subnet"], "VpnIpsecPhase1Interface-RemoteGwSubnet"); ok {
			if err = d.Set("remote_gw_subnet", vv); err != nil {
				return fmt.Errorf("Error reading remote_gw_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_gw_subnet: %v", err)
		}
	}

	if err = d.Set("remote_gw_ztna_tags", flattenVpnIpsecPhase1InterfaceRemoteGwZtnaTags(o["remote-gw-ztna-tags"], d, "remote_gw_ztna_tags")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-gw-ztna-tags"], "VpnIpsecPhase1Interface-RemoteGwZtnaTags"); ok {
			if err = d.Set("remote_gw_ztna_tags", vv); err != nil {
				return fmt.Errorf("Error reading remote_gw_ztna_tags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_gw_ztna_tags: %v", err)
		}
	}

	if err = d.Set("remote_gw6", flattenVpnIpsecPhase1InterfaceRemoteGw6(o["remote-gw6"], d, "remote_gw6")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-gw6"], "VpnIpsecPhase1Interface-RemoteGw6"); ok {
			if err = d.Set("remote_gw6", vv); err != nil {
				return fmt.Errorf("Error reading remote_gw6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_gw6: %v", err)
		}
	}

	if err = d.Set("remote_gw6_country", flattenVpnIpsecPhase1InterfaceRemoteGw6Country(o["remote-gw6-country"], d, "remote_gw6_country")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-gw6-country"], "VpnIpsecPhase1Interface-RemoteGw6Country"); ok {
			if err = d.Set("remote_gw6_country", vv); err != nil {
				return fmt.Errorf("Error reading remote_gw6_country: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_gw6_country: %v", err)
		}
	}

	if err = d.Set("remote_gw6_end_ip", flattenVpnIpsecPhase1InterfaceRemoteGw6EndIp(o["remote-gw6-end-ip"], d, "remote_gw6_end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-gw6-end-ip"], "VpnIpsecPhase1Interface-RemoteGw6EndIp"); ok {
			if err = d.Set("remote_gw6_end_ip", vv); err != nil {
				return fmt.Errorf("Error reading remote_gw6_end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_gw6_end_ip: %v", err)
		}
	}

	if err = d.Set("remote_gw6_match", flattenVpnIpsecPhase1InterfaceRemoteGw6Match(o["remote-gw6-match"], d, "remote_gw6_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-gw6-match"], "VpnIpsecPhase1Interface-RemoteGw6Match"); ok {
			if err = d.Set("remote_gw6_match", vv); err != nil {
				return fmt.Errorf("Error reading remote_gw6_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_gw6_match: %v", err)
		}
	}

	if err = d.Set("remote_gw6_start_ip", flattenVpnIpsecPhase1InterfaceRemoteGw6StartIp(o["remote-gw6-start-ip"], d, "remote_gw6_start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-gw6-start-ip"], "VpnIpsecPhase1Interface-RemoteGw6StartIp"); ok {
			if err = d.Set("remote_gw6_start_ip", vv); err != nil {
				return fmt.Errorf("Error reading remote_gw6_start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_gw6_start_ip: %v", err)
		}
	}

	if err = d.Set("remote_gw6_subnet", flattenVpnIpsecPhase1InterfaceRemoteGw6Subnet(o["remote-gw6-subnet"], d, "remote_gw6_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-gw6-subnet"], "VpnIpsecPhase1Interface-RemoteGw6Subnet"); ok {
			if err = d.Set("remote_gw6_subnet", vv); err != nil {
				return fmt.Errorf("Error reading remote_gw6_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_gw6_subnet: %v", err)
		}
	}

	if err = d.Set("remotegw_ddns", flattenVpnIpsecPhase1InterfaceRemotegwDdns(o["remotegw-ddns"], d, "remotegw_ddns")); err != nil {
		if vv, ok := fortiAPIPatch(o["remotegw-ddns"], "VpnIpsecPhase1Interface-RemotegwDdns"); ok {
			if err = d.Set("remotegw_ddns", vv); err != nil {
				return fmt.Errorf("Error reading remotegw_ddns: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remotegw_ddns: %v", err)
		}
	}

	if err = d.Set("rsa_signature_format", flattenVpnIpsecPhase1InterfaceRsaSignatureFormat(o["rsa-signature-format"], d, "rsa_signature_format")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsa-signature-format"], "VpnIpsecPhase1Interface-RsaSignatureFormat"); ok {
			if err = d.Set("rsa_signature_format", vv); err != nil {
				return fmt.Errorf("Error reading rsa_signature_format: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsa_signature_format: %v", err)
		}
	}

	if err = d.Set("rsa_signature_hash_override", flattenVpnIpsecPhase1InterfaceRsaSignatureHashOverride(o["rsa-signature-hash-override"], d, "rsa_signature_hash_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsa-signature-hash-override"], "VpnIpsecPhase1Interface-RsaSignatureHashOverride"); ok {
			if err = d.Set("rsa_signature_hash_override", vv); err != nil {
				return fmt.Errorf("Error reading rsa_signature_hash_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsa_signature_hash_override: %v", err)
		}
	}

	if err = d.Set("save_password", flattenVpnIpsecPhase1InterfaceSavePassword(o["save-password"], d, "save_password")); err != nil {
		if vv, ok := fortiAPIPatch(o["save-password"], "VpnIpsecPhase1Interface-SavePassword"); ok {
			if err = d.Set("save_password", vv); err != nil {
				return fmt.Errorf("Error reading save_password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading save_password: %v", err)
		}
	}

	if err = d.Set("send_cert_chain", flattenVpnIpsecPhase1InterfaceSendCertChain(o["send-cert-chain"], d, "send_cert_chain")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-cert-chain"], "VpnIpsecPhase1Interface-SendCertChain"); ok {
			if err = d.Set("send_cert_chain", vv); err != nil {
				return fmt.Errorf("Error reading send_cert_chain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_cert_chain: %v", err)
		}
	}

	if err = d.Set("shared_idle_timeout", flattenVpnIpsecPhase1InterfaceSharedIdleTimeout(o["shared-idle-timeout"], d, "shared_idle_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["shared-idle-timeout"], "VpnIpsecPhase1Interface-SharedIdleTimeout"); ok {
			if err = d.Set("shared_idle_timeout", vv); err != nil {
				return fmt.Errorf("Error reading shared_idle_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading shared_idle_timeout: %v", err)
		}
	}

	if err = d.Set("signature_hash_alg", flattenVpnIpsecPhase1InterfaceSignatureHashAlg(o["signature-hash-alg"], d, "signature_hash_alg")); err != nil {
		if vv, ok := fortiAPIPatch(o["signature-hash-alg"], "VpnIpsecPhase1Interface-SignatureHashAlg"); ok {
			if err = d.Set("signature_hash_alg", vv); err != nil {
				return fmt.Errorf("Error reading signature_hash_alg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading signature_hash_alg: %v", err)
		}
	}

	if err = d.Set("split_include_service", flattenVpnIpsecPhase1InterfaceSplitIncludeService(o["split-include-service"], d, "split_include_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["split-include-service"], "VpnIpsecPhase1Interface-SplitIncludeService"); ok {
			if err = d.Set("split_include_service", vv); err != nil {
				return fmt.Errorf("Error reading split_include_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading split_include_service: %v", err)
		}
	}

	if err = d.Set("suite_b", flattenVpnIpsecPhase1InterfaceSuiteB(o["suite-b"], d, "suite_b")); err != nil {
		if vv, ok := fortiAPIPatch(o["suite-b"], "VpnIpsecPhase1Interface-SuiteB"); ok {
			if err = d.Set("suite_b", vv); err != nil {
				return fmt.Errorf("Error reading suite_b: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading suite_b: %v", err)
		}
	}

	if err = d.Set("transit_gateway", flattenVpnIpsecPhase1InterfaceTransitGateway(o["transit-gateway"], d, "transit_gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["transit-gateway"], "VpnIpsecPhase1Interface-TransitGateway"); ok {
			if err = d.Set("transit_gateway", vv); err != nil {
				return fmt.Errorf("Error reading transit_gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading transit_gateway: %v", err)
		}
	}

	if err = d.Set("tunnel_search", flattenVpnIpsecPhase1InterfaceTunnelSearch(o["tunnel-search"], d, "tunnel_search")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-search"], "VpnIpsecPhase1Interface-TunnelSearch"); ok {
			if err = d.Set("tunnel_search", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_search: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_search: %v", err)
		}
	}

	if err = d.Set("transport", flattenVpnIpsecPhase1InterfaceTransport(o["transport"], d, "transport")); err != nil {
		if vv, ok := fortiAPIPatch(o["transport"], "VpnIpsecPhase1Interface-Transport"); ok {
			if err = d.Set("transport", vv); err != nil {
				return fmt.Errorf("Error reading transport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading transport: %v", err)
		}
	}

	if err = d.Set("type", flattenVpnIpsecPhase1InterfaceType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "VpnIpsecPhase1Interface-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("unity_support", flattenVpnIpsecPhase1InterfaceUnitySupport(o["unity-support"], d, "unity_support")); err != nil {
		if vv, ok := fortiAPIPatch(o["unity-support"], "VpnIpsecPhase1Interface-UnitySupport"); ok {
			if err = d.Set("unity_support", vv); err != nil {
				return fmt.Errorf("Error reading unity_support: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unity_support: %v", err)
		}
	}

	if err = d.Set("usrgrp", flattenVpnIpsecPhase1InterfaceUsrgrp(o["usrgrp"], d, "usrgrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["usrgrp"], "VpnIpsecPhase1Interface-Usrgrp"); ok {
			if err = d.Set("usrgrp", vv); err != nil {
				return fmt.Errorf("Error reading usrgrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading usrgrp: %v", err)
		}
	}

	if err = d.Set("vni", flattenVpnIpsecPhase1InterfaceVni(o["vni"], d, "vni")); err != nil {
		if vv, ok := fortiAPIPatch(o["vni"], "VpnIpsecPhase1Interface-Vni"); ok {
			if err = d.Set("vni", vv); err != nil {
				return fmt.Errorf("Error reading vni: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vni: %v", err)
		}
	}

	if err = d.Set("wizard_type", flattenVpnIpsecPhase1InterfaceWizardType(o["wizard-type"], d, "wizard_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["wizard-type"], "VpnIpsecPhase1Interface-WizardType"); ok {
			if err = d.Set("wizard_type", vv); err != nil {
				return fmt.Errorf("Error reading wizard_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wizard_type: %v", err)
		}
	}

	if err = d.Set("xauthtype", flattenVpnIpsecPhase1InterfaceXauthtype(o["xauthtype"], d, "xauthtype")); err != nil {
		if vv, ok := fortiAPIPatch(o["xauthtype"], "VpnIpsecPhase1Interface-Xauthtype"); ok {
			if err = d.Set("xauthtype", vv); err != nil {
				return fmt.Errorf("Error reading xauthtype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading xauthtype: %v", err)
		}
	}

	return nil
}

func flattenVpnIpsecPhase1InterfaceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnIpsecPhase1InterfaceAcctVerify(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAddGwRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAddRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAddke1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceAddke2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceAddke3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceAddke4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceAddke5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceAddke6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceAddke7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceAggregateMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAggregateWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAssignIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAssignIpFrom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAuthmethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAuthmethodRemote(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAuthpasswd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceAuthusr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAuthusrgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceAutoDiscoveryCrossover(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAutoDiscoveryForwarder(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAutoDiscoveryOfferInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAutoDiscoveryPsk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAutoDiscoveryReceiver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAutoDiscoverySender(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAutoDiscoveryShortcuts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAutoNegotiate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAutoTransportThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAzureAdAutoconnect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceBackupGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceBanner(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceCertIdValidation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceCertPeerUsernameStrip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceCertPeerUsernameValidation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceCertTrustStore(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceChildlessIke(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceClientAutoNegotiate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceClientKeepAlive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceClientResume(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceClientResumeInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceDefaultGw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceDefaultGwPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceDevId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceDevIdNotification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceDhcpRaGiaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceDhcp6RaLinkaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceDhgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceDigitalSignatureAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceDistance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceDnsMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceDpd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceDpdRetrycount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceDpdRetryinterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceEap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceEapCertAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceEapExcludePeergrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceEapIdentity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceEmsSnCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceEncapLocalGw4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceEncapLocalGw6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceEncapRemoteGw4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceEncapRemoteGw6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceEncapsulation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceEncapsulationAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceEnforceUniqueId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceEsn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceExchangeFgtDeviceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceExchangeInterfaceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceExchangeIpAddr4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceExchangeIpAddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceFallbackTcpThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceFecBase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceFecCodec(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceFecEgress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceFecHealthCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceFecIngress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceFecMappingProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceFecReceiveTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceFecRedundant(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceFecSendTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceFgspSync(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceForticlientEnforcement(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceFortinetEsp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceFragmentation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceFragmentationMtu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceGroupAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceGroupAuthenticationSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceHaSyncEspSeqno(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIdleTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIdleTimeoutinterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIkeVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceInboundDscpCopy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIncludeLocalLan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceInternalDomainList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceIpDelayInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpFragmentation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpsecTunnelSlot(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv4DnsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv4DnsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv4DnsServer3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv4EndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv4ExcludeRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-ip"], _ = expandVpnIpsecPhase1InterfaceIpv4ExcludeRangeEndIp(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandVpnIpsecPhase1InterfaceIpv4ExcludeRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip"], _ = expandVpnIpsecPhase1InterfaceIpv4ExcludeRangeStartIp(d, i["start_ip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandVpnIpsecPhase1InterfaceIpv4ExcludeRangeEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv4ExcludeRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv4ExcludeRangeStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv4Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceIpv4Netmask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv4SplitExclude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceIpv4SplitInclude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceIpv4StartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv4WinsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv4WinsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv6AutoLinklocal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv6DnsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv6DnsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv6DnsServer3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv6EndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv6ExcludeRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-ip"], _ = expandVpnIpsecPhase1InterfaceIpv6ExcludeRangeEndIp(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandVpnIpsecPhase1InterfaceIpv6ExcludeRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip"], _ = expandVpnIpsecPhase1InterfaceIpv6ExcludeRangeStartIp(d, i["start_ip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandVpnIpsecPhase1InterfaceIpv6ExcludeRangeEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv6ExcludeRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv6ExcludeRangeStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv6Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceIpv6Prefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv6SplitExclude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceIpv6SplitInclude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceIpv6StartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceKeepalive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceKeylife(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceKms(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceLinkCost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceLocalGw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceLocalGw6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceLocalid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceLocalidType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceLoopbackAsymroute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceMeshSelectorType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceModeCfg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceModeCfgAllowClientSelector(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceMonitorHoldDownDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceMonitorHoldDownTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceMonitorHoldDownType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceMonitorHoldDownWeekday(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceMonitorMin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceNattraversal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceNegotiateTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceNetDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceNetworkId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceNetworkOverlay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceNpuOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfacePacketRedistribution(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfacePassiveMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfacePeer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfacePeergrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfacePeerid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfacePeertype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfacePpk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfacePpkIdentity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfacePpkSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfacePriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceProposal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfacePsksecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfacePsksecretRemote(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceQkd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceQkdProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceReauth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRekey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRemoteGw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRemoteGwCountry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRemoteGwEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRemoteGwMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRemoteGwStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRemoteGwSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandVpnIpsecPhase1InterfaceRemoteGwZtnaTags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceRemoteGw6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRemoteGw6Country(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRemoteGw6EndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRemoteGw6Match(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRemoteGw6StartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRemoteGw6Subnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRemotegwDdns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRsaSignatureFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRsaSignatureHashOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceSavePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceSendCertChain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceSharedIdleTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceSignatureHashAlg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceSplitIncludeService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceSuiteB(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceTransitGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceTunnelSearch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceTransport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceUnitySupport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceUsrgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InterfaceVni(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceWizardType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceXauthtype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnIpsecPhase1Interface(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("acct_verify"); ok || d.HasChange("acct_verify") {
		t, err := expandVpnIpsecPhase1InterfaceAcctVerify(d, v, "acct_verify")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acct-verify"] = t
		}
	}

	if v, ok := d.GetOk("add_gw_route"); ok || d.HasChange("add_gw_route") {
		t, err := expandVpnIpsecPhase1InterfaceAddGwRoute(d, v, "add_gw_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-gw-route"] = t
		}
	}

	if v, ok := d.GetOk("add_route"); ok || d.HasChange("add_route") {
		t, err := expandVpnIpsecPhase1InterfaceAddRoute(d, v, "add_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-route"] = t
		}
	}

	if v, ok := d.GetOk("addke1"); ok || d.HasChange("addke1") {
		t, err := expandVpnIpsecPhase1InterfaceAddke1(d, v, "addke1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke1"] = t
		}
	}

	if v, ok := d.GetOk("addke2"); ok || d.HasChange("addke2") {
		t, err := expandVpnIpsecPhase1InterfaceAddke2(d, v, "addke2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke2"] = t
		}
	}

	if v, ok := d.GetOk("addke3"); ok || d.HasChange("addke3") {
		t, err := expandVpnIpsecPhase1InterfaceAddke3(d, v, "addke3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke3"] = t
		}
	}

	if v, ok := d.GetOk("addke4"); ok || d.HasChange("addke4") {
		t, err := expandVpnIpsecPhase1InterfaceAddke4(d, v, "addke4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke4"] = t
		}
	}

	if v, ok := d.GetOk("addke5"); ok || d.HasChange("addke5") {
		t, err := expandVpnIpsecPhase1InterfaceAddke5(d, v, "addke5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke5"] = t
		}
	}

	if v, ok := d.GetOk("addke6"); ok || d.HasChange("addke6") {
		t, err := expandVpnIpsecPhase1InterfaceAddke6(d, v, "addke6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke6"] = t
		}
	}

	if v, ok := d.GetOk("addke7"); ok || d.HasChange("addke7") {
		t, err := expandVpnIpsecPhase1InterfaceAddke7(d, v, "addke7")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke7"] = t
		}
	}

	if v, ok := d.GetOk("aggregate_member"); ok || d.HasChange("aggregate_member") {
		t, err := expandVpnIpsecPhase1InterfaceAggregateMember(d, v, "aggregate_member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aggregate-member"] = t
		}
	}

	if v, ok := d.GetOk("aggregate_weight"); ok || d.HasChange("aggregate_weight") {
		t, err := expandVpnIpsecPhase1InterfaceAggregateWeight(d, v, "aggregate_weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aggregate-weight"] = t
		}
	}

	if v, ok := d.GetOk("assign_ip"); ok || d.HasChange("assign_ip") {
		t, err := expandVpnIpsecPhase1InterfaceAssignIp(d, v, "assign_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["assign-ip"] = t
		}
	}

	if v, ok := d.GetOk("assign_ip_from"); ok || d.HasChange("assign_ip_from") {
		t, err := expandVpnIpsecPhase1InterfaceAssignIpFrom(d, v, "assign_ip_from")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["assign-ip-from"] = t
		}
	}

	if v, ok := d.GetOk("authmethod"); ok || d.HasChange("authmethod") {
		t, err := expandVpnIpsecPhase1InterfaceAuthmethod(d, v, "authmethod")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authmethod"] = t
		}
	}

	if v, ok := d.GetOk("authmethod_remote"); ok || d.HasChange("authmethod_remote") {
		t, err := expandVpnIpsecPhase1InterfaceAuthmethodRemote(d, v, "authmethod_remote")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authmethod-remote"] = t
		}
	}

	if v, ok := d.GetOk("authpasswd"); ok || d.HasChange("authpasswd") {
		t, err := expandVpnIpsecPhase1InterfaceAuthpasswd(d, v, "authpasswd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authpasswd"] = t
		}
	}

	if v, ok := d.GetOk("authusr"); ok || d.HasChange("authusr") {
		t, err := expandVpnIpsecPhase1InterfaceAuthusr(d, v, "authusr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authusr"] = t
		}
	}

	if v, ok := d.GetOk("authusrgrp"); ok || d.HasChange("authusrgrp") {
		t, err := expandVpnIpsecPhase1InterfaceAuthusrgrp(d, v, "authusrgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authusrgrp"] = t
		}
	}

	if v, ok := d.GetOk("auto_discovery_crossover"); ok || d.HasChange("auto_discovery_crossover") {
		t, err := expandVpnIpsecPhase1InterfaceAutoDiscoveryCrossover(d, v, "auto_discovery_crossover")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-discovery-crossover"] = t
		}
	}

	if v, ok := d.GetOk("auto_discovery_forwarder"); ok || d.HasChange("auto_discovery_forwarder") {
		t, err := expandVpnIpsecPhase1InterfaceAutoDiscoveryForwarder(d, v, "auto_discovery_forwarder")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-discovery-forwarder"] = t
		}
	}

	if v, ok := d.GetOk("auto_discovery_offer_interval"); ok || d.HasChange("auto_discovery_offer_interval") {
		t, err := expandVpnIpsecPhase1InterfaceAutoDiscoveryOfferInterval(d, v, "auto_discovery_offer_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-discovery-offer-interval"] = t
		}
	}

	if v, ok := d.GetOk("auto_discovery_psk"); ok || d.HasChange("auto_discovery_psk") {
		t, err := expandVpnIpsecPhase1InterfaceAutoDiscoveryPsk(d, v, "auto_discovery_psk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-discovery-psk"] = t
		}
	}

	if v, ok := d.GetOk("auto_discovery_receiver"); ok || d.HasChange("auto_discovery_receiver") {
		t, err := expandVpnIpsecPhase1InterfaceAutoDiscoveryReceiver(d, v, "auto_discovery_receiver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-discovery-receiver"] = t
		}
	}

	if v, ok := d.GetOk("auto_discovery_sender"); ok || d.HasChange("auto_discovery_sender") {
		t, err := expandVpnIpsecPhase1InterfaceAutoDiscoverySender(d, v, "auto_discovery_sender")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-discovery-sender"] = t
		}
	}

	if v, ok := d.GetOk("auto_discovery_shortcuts"); ok || d.HasChange("auto_discovery_shortcuts") {
		t, err := expandVpnIpsecPhase1InterfaceAutoDiscoveryShortcuts(d, v, "auto_discovery_shortcuts")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-discovery-shortcuts"] = t
		}
	}

	if v, ok := d.GetOk("auto_negotiate"); ok || d.HasChange("auto_negotiate") {
		t, err := expandVpnIpsecPhase1InterfaceAutoNegotiate(d, v, "auto_negotiate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-negotiate"] = t
		}
	}

	if v, ok := d.GetOk("auto_transport_threshold"); ok || d.HasChange("auto_transport_threshold") {
		t, err := expandVpnIpsecPhase1InterfaceAutoTransportThreshold(d, v, "auto_transport_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-transport-threshold"] = t
		}
	}

	if v, ok := d.GetOk("azure_ad_autoconnect"); ok || d.HasChange("azure_ad_autoconnect") {
		t, err := expandVpnIpsecPhase1InterfaceAzureAdAutoconnect(d, v, "azure_ad_autoconnect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["azure-ad-autoconnect"] = t
		}
	}

	if v, ok := d.GetOk("backup_gateway"); ok || d.HasChange("backup_gateway") {
		t, err := expandVpnIpsecPhase1InterfaceBackupGateway(d, v, "backup_gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backup-gateway"] = t
		}
	}

	if v, ok := d.GetOk("banner"); ok || d.HasChange("banner") {
		t, err := expandVpnIpsecPhase1InterfaceBanner(d, v, "banner")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["banner"] = t
		}
	}

	if v, ok := d.GetOk("cert_id_validation"); ok || d.HasChange("cert_id_validation") {
		t, err := expandVpnIpsecPhase1InterfaceCertIdValidation(d, v, "cert_id_validation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-id-validation"] = t
		}
	}

	if v, ok := d.GetOk("cert_peer_username_strip"); ok || d.HasChange("cert_peer_username_strip") {
		t, err := expandVpnIpsecPhase1InterfaceCertPeerUsernameStrip(d, v, "cert_peer_username_strip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-peer-username-strip"] = t
		}
	}

	if v, ok := d.GetOk("cert_peer_username_validation"); ok || d.HasChange("cert_peer_username_validation") {
		t, err := expandVpnIpsecPhase1InterfaceCertPeerUsernameValidation(d, v, "cert_peer_username_validation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-peer-username-validation"] = t
		}
	}

	if v, ok := d.GetOk("cert_trust_store"); ok || d.HasChange("cert_trust_store") {
		t, err := expandVpnIpsecPhase1InterfaceCertTrustStore(d, v, "cert_trust_store")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-trust-store"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok || d.HasChange("certificate") {
		t, err := expandVpnIpsecPhase1InterfaceCertificate(d, v, "certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("childless_ike"); ok || d.HasChange("childless_ike") {
		t, err := expandVpnIpsecPhase1InterfaceChildlessIke(d, v, "childless_ike")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["childless-ike"] = t
		}
	}

	if v, ok := d.GetOk("client_auto_negotiate"); ok || d.HasChange("client_auto_negotiate") {
		t, err := expandVpnIpsecPhase1InterfaceClientAutoNegotiate(d, v, "client_auto_negotiate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-auto-negotiate"] = t
		}
	}

	if v, ok := d.GetOk("client_keep_alive"); ok || d.HasChange("client_keep_alive") {
		t, err := expandVpnIpsecPhase1InterfaceClientKeepAlive(d, v, "client_keep_alive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-keep-alive"] = t
		}
	}

	if v, ok := d.GetOk("client_resume"); ok || d.HasChange("client_resume") {
		t, err := expandVpnIpsecPhase1InterfaceClientResume(d, v, "client_resume")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-resume"] = t
		}
	}

	if v, ok := d.GetOk("client_resume_interval"); ok || d.HasChange("client_resume_interval") {
		t, err := expandVpnIpsecPhase1InterfaceClientResumeInterval(d, v, "client_resume_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-resume-interval"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandVpnIpsecPhase1InterfaceComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("default_gw"); ok || d.HasChange("default_gw") {
		t, err := expandVpnIpsecPhase1InterfaceDefaultGw(d, v, "default_gw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-gw"] = t
		}
	}

	if v, ok := d.GetOk("default_gw_priority"); ok || d.HasChange("default_gw_priority") {
		t, err := expandVpnIpsecPhase1InterfaceDefaultGwPriority(d, v, "default_gw_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-gw-priority"] = t
		}
	}

	if v, ok := d.GetOk("dev_id"); ok || d.HasChange("dev_id") {
		t, err := expandVpnIpsecPhase1InterfaceDevId(d, v, "dev_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dev-id"] = t
		}
	}

	if v, ok := d.GetOk("dev_id_notification"); ok || d.HasChange("dev_id_notification") {
		t, err := expandVpnIpsecPhase1InterfaceDevIdNotification(d, v, "dev_id_notification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dev-id-notification"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_ra_giaddr"); ok || d.HasChange("dhcp_ra_giaddr") {
		t, err := expandVpnIpsecPhase1InterfaceDhcpRaGiaddr(d, v, "dhcp_ra_giaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-ra-giaddr"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_ra_linkaddr"); ok || d.HasChange("dhcp6_ra_linkaddr") {
		t, err := expandVpnIpsecPhase1InterfaceDhcp6RaLinkaddr(d, v, "dhcp6_ra_linkaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-ra-linkaddr"] = t
		}
	}

	if v, ok := d.GetOk("dhgrp"); ok || d.HasChange("dhgrp") {
		t, err := expandVpnIpsecPhase1InterfaceDhgrp(d, v, "dhgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhgrp"] = t
		}
	}

	if v, ok := d.GetOk("digital_signature_auth"); ok || d.HasChange("digital_signature_auth") {
		t, err := expandVpnIpsecPhase1InterfaceDigitalSignatureAuth(d, v, "digital_signature_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["digital-signature-auth"] = t
		}
	}

	if v, ok := d.GetOk("distance"); ok || d.HasChange("distance") {
		t, err := expandVpnIpsecPhase1InterfaceDistance(d, v, "distance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance"] = t
		}
	}

	if v, ok := d.GetOk("dns_mode"); ok || d.HasChange("dns_mode") {
		t, err := expandVpnIpsecPhase1InterfaceDnsMode(d, v, "dns_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-mode"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok || d.HasChange("domain") {
		t, err := expandVpnIpsecPhase1InterfaceDomain(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("dpd"); ok || d.HasChange("dpd") {
		t, err := expandVpnIpsecPhase1InterfaceDpd(d, v, "dpd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dpd"] = t
		}
	}

	if v, ok := d.GetOk("dpd_retrycount"); ok || d.HasChange("dpd_retrycount") {
		t, err := expandVpnIpsecPhase1InterfaceDpdRetrycount(d, v, "dpd_retrycount")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dpd-retrycount"] = t
		}
	}

	if v, ok := d.GetOk("dpd_retryinterval"); ok || d.HasChange("dpd_retryinterval") {
		t, err := expandVpnIpsecPhase1InterfaceDpdRetryinterval(d, v, "dpd_retryinterval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dpd-retryinterval"] = t
		}
	}

	if v, ok := d.GetOk("eap"); ok || d.HasChange("eap") {
		t, err := expandVpnIpsecPhase1InterfaceEap(d, v, "eap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap"] = t
		}
	}

	if v, ok := d.GetOk("eap_cert_auth"); ok || d.HasChange("eap_cert_auth") {
		t, err := expandVpnIpsecPhase1InterfaceEapCertAuth(d, v, "eap_cert_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-cert-auth"] = t
		}
	}

	if v, ok := d.GetOk("eap_exclude_peergrp"); ok || d.HasChange("eap_exclude_peergrp") {
		t, err := expandVpnIpsecPhase1InterfaceEapExcludePeergrp(d, v, "eap_exclude_peergrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-exclude-peergrp"] = t
		}
	}

	if v, ok := d.GetOk("eap_identity"); ok || d.HasChange("eap_identity") {
		t, err := expandVpnIpsecPhase1InterfaceEapIdentity(d, v, "eap_identity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-identity"] = t
		}
	}

	if v, ok := d.GetOk("ems_sn_check"); ok || d.HasChange("ems_sn_check") {
		t, err := expandVpnIpsecPhase1InterfaceEmsSnCheck(d, v, "ems_sn_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ems-sn-check"] = t
		}
	}

	if v, ok := d.GetOk("encap_local_gw4"); ok || d.HasChange("encap_local_gw4") {
		t, err := expandVpnIpsecPhase1InterfaceEncapLocalGw4(d, v, "encap_local_gw4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encap-local-gw4"] = t
		}
	}

	if v, ok := d.GetOk("encap_local_gw6"); ok || d.HasChange("encap_local_gw6") {
		t, err := expandVpnIpsecPhase1InterfaceEncapLocalGw6(d, v, "encap_local_gw6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encap-local-gw6"] = t
		}
	}

	if v, ok := d.GetOk("encap_remote_gw4"); ok || d.HasChange("encap_remote_gw4") {
		t, err := expandVpnIpsecPhase1InterfaceEncapRemoteGw4(d, v, "encap_remote_gw4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encap-remote-gw4"] = t
		}
	}

	if v, ok := d.GetOk("encap_remote_gw6"); ok || d.HasChange("encap_remote_gw6") {
		t, err := expandVpnIpsecPhase1InterfaceEncapRemoteGw6(d, v, "encap_remote_gw6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encap-remote-gw6"] = t
		}
	}

	if v, ok := d.GetOk("encapsulation"); ok || d.HasChange("encapsulation") {
		t, err := expandVpnIpsecPhase1InterfaceEncapsulation(d, v, "encapsulation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encapsulation"] = t
		}
	}

	if v, ok := d.GetOk("encapsulation_address"); ok || d.HasChange("encapsulation_address") {
		t, err := expandVpnIpsecPhase1InterfaceEncapsulationAddress(d, v, "encapsulation_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encapsulation-address"] = t
		}
	}

	if v, ok := d.GetOk("enforce_unique_id"); ok || d.HasChange("enforce_unique_id") {
		t, err := expandVpnIpsecPhase1InterfaceEnforceUniqueId(d, v, "enforce_unique_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enforce-unique-id"] = t
		}
	}

	if v, ok := d.GetOk("esn"); ok || d.HasChange("esn") {
		t, err := expandVpnIpsecPhase1InterfaceEsn(d, v, "esn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["esn"] = t
		}
	}

	if v, ok := d.GetOk("exchange_fgt_device_id"); ok || d.HasChange("exchange_fgt_device_id") {
		t, err := expandVpnIpsecPhase1InterfaceExchangeFgtDeviceId(d, v, "exchange_fgt_device_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exchange-fgt-device-id"] = t
		}
	}

	if v, ok := d.GetOk("exchange_interface_ip"); ok || d.HasChange("exchange_interface_ip") {
		t, err := expandVpnIpsecPhase1InterfaceExchangeInterfaceIp(d, v, "exchange_interface_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exchange-interface-ip"] = t
		}
	}

	if v, ok := d.GetOk("exchange_ip_addr4"); ok || d.HasChange("exchange_ip_addr4") {
		t, err := expandVpnIpsecPhase1InterfaceExchangeIpAddr4(d, v, "exchange_ip_addr4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exchange-ip-addr4"] = t
		}
	}

	if v, ok := d.GetOk("exchange_ip_addr6"); ok || d.HasChange("exchange_ip_addr6") {
		t, err := expandVpnIpsecPhase1InterfaceExchangeIpAddr6(d, v, "exchange_ip_addr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exchange-ip-addr6"] = t
		}
	}

	if v, ok := d.GetOk("fallback_tcp_threshold"); ok || d.HasChange("fallback_tcp_threshold") {
		t, err := expandVpnIpsecPhase1InterfaceFallbackTcpThreshold(d, v, "fallback_tcp_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fallback-tcp-threshold"] = t
		}
	}

	if v, ok := d.GetOk("fec_base"); ok || d.HasChange("fec_base") {
		t, err := expandVpnIpsecPhase1InterfaceFecBase(d, v, "fec_base")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-base"] = t
		}
	}

	if v, ok := d.GetOk("fec_codec"); ok || d.HasChange("fec_codec") {
		t, err := expandVpnIpsecPhase1InterfaceFecCodec(d, v, "fec_codec")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-codec"] = t
		}
	}

	if v, ok := d.GetOk("fec_egress"); ok || d.HasChange("fec_egress") {
		t, err := expandVpnIpsecPhase1InterfaceFecEgress(d, v, "fec_egress")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-egress"] = t
		}
	}

	if v, ok := d.GetOk("fec_health_check"); ok || d.HasChange("fec_health_check") {
		t, err := expandVpnIpsecPhase1InterfaceFecHealthCheck(d, v, "fec_health_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-health-check"] = t
		}
	}

	if v, ok := d.GetOk("fec_ingress"); ok || d.HasChange("fec_ingress") {
		t, err := expandVpnIpsecPhase1InterfaceFecIngress(d, v, "fec_ingress")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-ingress"] = t
		}
	}

	if v, ok := d.GetOk("fec_mapping_profile"); ok || d.HasChange("fec_mapping_profile") {
		t, err := expandVpnIpsecPhase1InterfaceFecMappingProfile(d, v, "fec_mapping_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-mapping-profile"] = t
		}
	}

	if v, ok := d.GetOk("fec_receive_timeout"); ok || d.HasChange("fec_receive_timeout") {
		t, err := expandVpnIpsecPhase1InterfaceFecReceiveTimeout(d, v, "fec_receive_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-receive-timeout"] = t
		}
	}

	if v, ok := d.GetOk("fec_redundant"); ok || d.HasChange("fec_redundant") {
		t, err := expandVpnIpsecPhase1InterfaceFecRedundant(d, v, "fec_redundant")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-redundant"] = t
		}
	}

	if v, ok := d.GetOk("fec_send_timeout"); ok || d.HasChange("fec_send_timeout") {
		t, err := expandVpnIpsecPhase1InterfaceFecSendTimeout(d, v, "fec_send_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-send-timeout"] = t
		}
	}

	if v, ok := d.GetOk("fgsp_sync"); ok || d.HasChange("fgsp_sync") {
		t, err := expandVpnIpsecPhase1InterfaceFgspSync(d, v, "fgsp_sync")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgsp-sync"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_enforcement"); ok || d.HasChange("forticlient_enforcement") {
		t, err := expandVpnIpsecPhase1InterfaceForticlientEnforcement(d, v, "forticlient_enforcement")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-enforcement"] = t
		}
	}

	if v, ok := d.GetOk("fortinet_esp"); ok || d.HasChange("fortinet_esp") {
		t, err := expandVpnIpsecPhase1InterfaceFortinetEsp(d, v, "fortinet_esp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortinet-esp"] = t
		}
	}

	if v, ok := d.GetOk("fragmentation"); ok || d.HasChange("fragmentation") {
		t, err := expandVpnIpsecPhase1InterfaceFragmentation(d, v, "fragmentation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fragmentation"] = t
		}
	}

	if v, ok := d.GetOk("fragmentation_mtu"); ok || d.HasChange("fragmentation_mtu") {
		t, err := expandVpnIpsecPhase1InterfaceFragmentationMtu(d, v, "fragmentation_mtu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fragmentation-mtu"] = t
		}
	}

	if v, ok := d.GetOk("group_authentication"); ok || d.HasChange("group_authentication") {
		t, err := expandVpnIpsecPhase1InterfaceGroupAuthentication(d, v, "group_authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-authentication"] = t
		}
	}

	if v, ok := d.GetOk("group_authentication_secret"); ok || d.HasChange("group_authentication_secret") {
		t, err := expandVpnIpsecPhase1InterfaceGroupAuthenticationSecret(d, v, "group_authentication_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-authentication-secret"] = t
		}
	}

	if v, ok := d.GetOk("ha_sync_esp_seqno"); ok || d.HasChange("ha_sync_esp_seqno") {
		t, err := expandVpnIpsecPhase1InterfaceHaSyncEspSeqno(d, v, "ha_sync_esp_seqno")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-sync-esp-seqno"] = t
		}
	}

	if v, ok := d.GetOk("idle_timeout"); ok || d.HasChange("idle_timeout") {
		t, err := expandVpnIpsecPhase1InterfaceIdleTimeout(d, v, "idle_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idle-timeout"] = t
		}
	}

	if v, ok := d.GetOk("idle_timeoutinterval"); ok || d.HasChange("idle_timeoutinterval") {
		t, err := expandVpnIpsecPhase1InterfaceIdleTimeoutinterval(d, v, "idle_timeoutinterval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idle-timeoutinterval"] = t
		}
	}

	if v, ok := d.GetOk("ike_version"); ok || d.HasChange("ike_version") {
		t, err := expandVpnIpsecPhase1InterfaceIkeVersion(d, v, "ike_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-version"] = t
		}
	}

	if v, ok := d.GetOk("inbound_dscp_copy"); ok || d.HasChange("inbound_dscp_copy") {
		t, err := expandVpnIpsecPhase1InterfaceInboundDscpCopy(d, v, "inbound_dscp_copy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inbound-dscp-copy"] = t
		}
	}

	if v, ok := d.GetOk("include_local_lan"); ok || d.HasChange("include_local_lan") {
		t, err := expandVpnIpsecPhase1InterfaceIncludeLocalLan(d, v, "include_local_lan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["include-local-lan"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandVpnIpsecPhase1InterfaceInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("internal_domain_list"); ok || d.HasChange("internal_domain_list") {
		t, err := expandVpnIpsecPhase1InterfaceInternalDomainList(d, v, "internal_domain_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internal-domain-list"] = t
		}
	}

	if v, ok := d.GetOk("ip_delay_interval"); ok || d.HasChange("ip_delay_interval") {
		t, err := expandVpnIpsecPhase1InterfaceIpDelayInterval(d, v, "ip_delay_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-delay-interval"] = t
		}
	}

	if v, ok := d.GetOk("ip_fragmentation"); ok || d.HasChange("ip_fragmentation") {
		t, err := expandVpnIpsecPhase1InterfaceIpFragmentation(d, v, "ip_fragmentation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-fragmentation"] = t
		}
	}

	if v, ok := d.GetOk("ip_version"); ok || d.HasChange("ip_version") {
		t, err := expandVpnIpsecPhase1InterfaceIpVersion(d, v, "ip_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-version"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_tunnel_slot"); ok || d.HasChange("ipsec_tunnel_slot") {
		t, err := expandVpnIpsecPhase1InterfaceIpsecTunnelSlot(d, v, "ipsec_tunnel_slot")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-tunnel-slot"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_dns_server1"); ok || d.HasChange("ipv4_dns_server1") {
		t, err := expandVpnIpsecPhase1InterfaceIpv4DnsServer1(d, v, "ipv4_dns_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_dns_server2"); ok || d.HasChange("ipv4_dns_server2") {
		t, err := expandVpnIpsecPhase1InterfaceIpv4DnsServer2(d, v, "ipv4_dns_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_dns_server3"); ok || d.HasChange("ipv4_dns_server3") {
		t, err := expandVpnIpsecPhase1InterfaceIpv4DnsServer3(d, v, "ipv4_dns_server3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-dns-server3"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_end_ip"); ok || d.HasChange("ipv4_end_ip") {
		t, err := expandVpnIpsecPhase1InterfaceIpv4EndIp(d, v, "ipv4_end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-end-ip"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_exclude_range"); ok || d.HasChange("ipv4_exclude_range") {
		t, err := expandVpnIpsecPhase1InterfaceIpv4ExcludeRange(d, v, "ipv4_exclude_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-exclude-range"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_name"); ok || d.HasChange("ipv4_name") {
		t, err := expandVpnIpsecPhase1InterfaceIpv4Name(d, v, "ipv4_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-name"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_netmask"); ok || d.HasChange("ipv4_netmask") {
		t, err := expandVpnIpsecPhase1InterfaceIpv4Netmask(d, v, "ipv4_netmask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-netmask"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_split_exclude"); ok || d.HasChange("ipv4_split_exclude") {
		t, err := expandVpnIpsecPhase1InterfaceIpv4SplitExclude(d, v, "ipv4_split_exclude")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-split-exclude"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_split_include"); ok || d.HasChange("ipv4_split_include") {
		t, err := expandVpnIpsecPhase1InterfaceIpv4SplitInclude(d, v, "ipv4_split_include")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-split-include"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_start_ip"); ok || d.HasChange("ipv4_start_ip") {
		t, err := expandVpnIpsecPhase1InterfaceIpv4StartIp(d, v, "ipv4_start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-start-ip"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_wins_server1"); ok || d.HasChange("ipv4_wins_server1") {
		t, err := expandVpnIpsecPhase1InterfaceIpv4WinsServer1(d, v, "ipv4_wins_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-wins-server1"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_wins_server2"); ok || d.HasChange("ipv4_wins_server2") {
		t, err := expandVpnIpsecPhase1InterfaceIpv4WinsServer2(d, v, "ipv4_wins_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-wins-server2"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_auto_linklocal"); ok || d.HasChange("ipv6_auto_linklocal") {
		t, err := expandVpnIpsecPhase1InterfaceIpv6AutoLinklocal(d, v, "ipv6_auto_linklocal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-auto-linklocal"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_dns_server1"); ok || d.HasChange("ipv6_dns_server1") {
		t, err := expandVpnIpsecPhase1InterfaceIpv6DnsServer1(d, v, "ipv6_dns_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_dns_server2"); ok || d.HasChange("ipv6_dns_server2") {
		t, err := expandVpnIpsecPhase1InterfaceIpv6DnsServer2(d, v, "ipv6_dns_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_dns_server3"); ok || d.HasChange("ipv6_dns_server3") {
		t, err := expandVpnIpsecPhase1InterfaceIpv6DnsServer3(d, v, "ipv6_dns_server3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-dns-server3"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_end_ip"); ok || d.HasChange("ipv6_end_ip") {
		t, err := expandVpnIpsecPhase1InterfaceIpv6EndIp(d, v, "ipv6_end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-end-ip"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_exclude_range"); ok || d.HasChange("ipv6_exclude_range") {
		t, err := expandVpnIpsecPhase1InterfaceIpv6ExcludeRange(d, v, "ipv6_exclude_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-exclude-range"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_name"); ok || d.HasChange("ipv6_name") {
		t, err := expandVpnIpsecPhase1InterfaceIpv6Name(d, v, "ipv6_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-name"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_prefix"); ok || d.HasChange("ipv6_prefix") {
		t, err := expandVpnIpsecPhase1InterfaceIpv6Prefix(d, v, "ipv6_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-prefix"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_split_exclude"); ok || d.HasChange("ipv6_split_exclude") {
		t, err := expandVpnIpsecPhase1InterfaceIpv6SplitExclude(d, v, "ipv6_split_exclude")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-split-exclude"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_split_include"); ok || d.HasChange("ipv6_split_include") {
		t, err := expandVpnIpsecPhase1InterfaceIpv6SplitInclude(d, v, "ipv6_split_include")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-split-include"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_start_ip"); ok || d.HasChange("ipv6_start_ip") {
		t, err := expandVpnIpsecPhase1InterfaceIpv6StartIp(d, v, "ipv6_start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-start-ip"] = t
		}
	}

	if v, ok := d.GetOk("keepalive"); ok || d.HasChange("keepalive") {
		t, err := expandVpnIpsecPhase1InterfaceKeepalive(d, v, "keepalive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keepalive"] = t
		}
	}

	if v, ok := d.GetOk("keylife"); ok || d.HasChange("keylife") {
		t, err := expandVpnIpsecPhase1InterfaceKeylife(d, v, "keylife")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keylife"] = t
		}
	}

	if v, ok := d.GetOk("kms"); ok || d.HasChange("kms") {
		t, err := expandVpnIpsecPhase1InterfaceKms(d, v, "kms")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["kms"] = t
		}
	}

	if v, ok := d.GetOk("link_cost"); ok || d.HasChange("link_cost") {
		t, err := expandVpnIpsecPhase1InterfaceLinkCost(d, v, "link_cost")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-cost"] = t
		}
	}

	if v, ok := d.GetOk("local_gw"); ok || d.HasChange("local_gw") {
		t, err := expandVpnIpsecPhase1InterfaceLocalGw(d, v, "local_gw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-gw"] = t
		}
	}

	if v, ok := d.GetOk("local_gw6"); ok || d.HasChange("local_gw6") {
		t, err := expandVpnIpsecPhase1InterfaceLocalGw6(d, v, "local_gw6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-gw6"] = t
		}
	}

	if v, ok := d.GetOk("localid"); ok || d.HasChange("localid") {
		t, err := expandVpnIpsecPhase1InterfaceLocalid(d, v, "localid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["localid"] = t
		}
	}

	if v, ok := d.GetOk("localid_type"); ok || d.HasChange("localid_type") {
		t, err := expandVpnIpsecPhase1InterfaceLocalidType(d, v, "localid_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["localid-type"] = t
		}
	}

	if v, ok := d.GetOk("loopback_asymroute"); ok || d.HasChange("loopback_asymroute") {
		t, err := expandVpnIpsecPhase1InterfaceLoopbackAsymroute(d, v, "loopback_asymroute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["loopback-asymroute"] = t
		}
	}

	if v, ok := d.GetOk("mesh_selector_type"); ok || d.HasChange("mesh_selector_type") {
		t, err := expandVpnIpsecPhase1InterfaceMeshSelectorType(d, v, "mesh_selector_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mesh-selector-type"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandVpnIpsecPhase1InterfaceMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("mode_cfg"); ok || d.HasChange("mode_cfg") {
		t, err := expandVpnIpsecPhase1InterfaceModeCfg(d, v, "mode_cfg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode-cfg"] = t
		}
	}

	if v, ok := d.GetOk("mode_cfg_allow_client_selector"); ok || d.HasChange("mode_cfg_allow_client_selector") {
		t, err := expandVpnIpsecPhase1InterfaceModeCfgAllowClientSelector(d, v, "mode_cfg_allow_client_selector")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode-cfg-allow-client-selector"] = t
		}
	}

	if v, ok := d.GetOk("monitor"); ok || d.HasChange("monitor") {
		t, err := expandVpnIpsecPhase1InterfaceMonitor(d, v, "monitor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor"] = t
		}
	}

	if v, ok := d.GetOk("monitor_hold_down_delay"); ok || d.HasChange("monitor_hold_down_delay") {
		t, err := expandVpnIpsecPhase1InterfaceMonitorHoldDownDelay(d, v, "monitor_hold_down_delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-hold-down-delay"] = t
		}
	}

	if v, ok := d.GetOk("monitor_hold_down_time"); ok || d.HasChange("monitor_hold_down_time") {
		t, err := expandVpnIpsecPhase1InterfaceMonitorHoldDownTime(d, v, "monitor_hold_down_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-hold-down-time"] = t
		}
	}

	if v, ok := d.GetOk("monitor_hold_down_type"); ok || d.HasChange("monitor_hold_down_type") {
		t, err := expandVpnIpsecPhase1InterfaceMonitorHoldDownType(d, v, "monitor_hold_down_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-hold-down-type"] = t
		}
	}

	if v, ok := d.GetOk("monitor_hold_down_weekday"); ok || d.HasChange("monitor_hold_down_weekday") {
		t, err := expandVpnIpsecPhase1InterfaceMonitorHoldDownWeekday(d, v, "monitor_hold_down_weekday")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-hold-down-weekday"] = t
		}
	}

	if v, ok := d.GetOk("monitor_min"); ok || d.HasChange("monitor_min") {
		t, err := expandVpnIpsecPhase1InterfaceMonitorMin(d, v, "monitor_min")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-min"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandVpnIpsecPhase1InterfaceName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nattraversal"); ok || d.HasChange("nattraversal") {
		t, err := expandVpnIpsecPhase1InterfaceNattraversal(d, v, "nattraversal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nattraversal"] = t
		}
	}

	if v, ok := d.GetOk("negotiate_timeout"); ok || d.HasChange("negotiate_timeout") {
		t, err := expandVpnIpsecPhase1InterfaceNegotiateTimeout(d, v, "negotiate_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["negotiate-timeout"] = t
		}
	}

	if v, ok := d.GetOk("net_device"); ok || d.HasChange("net_device") {
		t, err := expandVpnIpsecPhase1InterfaceNetDevice(d, v, "net_device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["net-device"] = t
		}
	}

	if v, ok := d.GetOk("network_id"); ok || d.HasChange("network_id") {
		t, err := expandVpnIpsecPhase1InterfaceNetworkId(d, v, "network_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-id"] = t
		}
	}

	if v, ok := d.GetOk("network_overlay"); ok || d.HasChange("network_overlay") {
		t, err := expandVpnIpsecPhase1InterfaceNetworkOverlay(d, v, "network_overlay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-overlay"] = t
		}
	}

	if v, ok := d.GetOk("npu_offload"); ok || d.HasChange("npu_offload") {
		t, err := expandVpnIpsecPhase1InterfaceNpuOffload(d, v, "npu_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["npu-offload"] = t
		}
	}

	if v, ok := d.GetOk("packet_redistribution"); ok || d.HasChange("packet_redistribution") {
		t, err := expandVpnIpsecPhase1InterfacePacketRedistribution(d, v, "packet_redistribution")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-redistribution"] = t
		}
	}

	if v, ok := d.GetOk("passive_mode"); ok || d.HasChange("passive_mode") {
		t, err := expandVpnIpsecPhase1InterfacePassiveMode(d, v, "passive_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passive-mode"] = t
		}
	}

	if v, ok := d.GetOk("peer"); ok || d.HasChange("peer") {
		t, err := expandVpnIpsecPhase1InterfacePeer(d, v, "peer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer"] = t
		}
	}

	if v, ok := d.GetOk("peergrp"); ok || d.HasChange("peergrp") {
		t, err := expandVpnIpsecPhase1InterfacePeergrp(d, v, "peergrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peergrp"] = t
		}
	}

	if v, ok := d.GetOk("peerid"); ok || d.HasChange("peerid") {
		t, err := expandVpnIpsecPhase1InterfacePeerid(d, v, "peerid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peerid"] = t
		}
	}

	if v, ok := d.GetOk("peertype"); ok || d.HasChange("peertype") {
		t, err := expandVpnIpsecPhase1InterfacePeertype(d, v, "peertype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peertype"] = t
		}
	}

	if v, ok := d.GetOk("ppk"); ok || d.HasChange("ppk") {
		t, err := expandVpnIpsecPhase1InterfacePpk(d, v, "ppk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppk"] = t
		}
	}

	if v, ok := d.GetOk("ppk_identity"); ok || d.HasChange("ppk_identity") {
		t, err := expandVpnIpsecPhase1InterfacePpkIdentity(d, v, "ppk_identity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppk-identity"] = t
		}
	}

	if v, ok := d.GetOk("ppk_secret"); ok || d.HasChange("ppk_secret") {
		t, err := expandVpnIpsecPhase1InterfacePpkSecret(d, v, "ppk_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppk-secret"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandVpnIpsecPhase1InterfacePriority(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("proposal"); ok || d.HasChange("proposal") {
		t, err := expandVpnIpsecPhase1InterfaceProposal(d, v, "proposal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proposal"] = t
		}
	}

	if v, ok := d.GetOk("psksecret"); ok || d.HasChange("psksecret") {
		t, err := expandVpnIpsecPhase1InterfacePsksecret(d, v, "psksecret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["psksecret"] = t
		}
	}

	if v, ok := d.GetOk("psksecret_remote"); ok || d.HasChange("psksecret_remote") {
		t, err := expandVpnIpsecPhase1InterfacePsksecretRemote(d, v, "psksecret_remote")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["psksecret-remote"] = t
		}
	}

	if v, ok := d.GetOk("qkd"); ok || d.HasChange("qkd") {
		t, err := expandVpnIpsecPhase1InterfaceQkd(d, v, "qkd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qkd"] = t
		}
	}

	if v, ok := d.GetOk("qkd_profile"); ok || d.HasChange("qkd_profile") {
		t, err := expandVpnIpsecPhase1InterfaceQkdProfile(d, v, "qkd_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qkd-profile"] = t
		}
	}

	if v, ok := d.GetOk("reauth"); ok || d.HasChange("reauth") {
		t, err := expandVpnIpsecPhase1InterfaceReauth(d, v, "reauth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reauth"] = t
		}
	}

	if v, ok := d.GetOk("rekey"); ok || d.HasChange("rekey") {
		t, err := expandVpnIpsecPhase1InterfaceRekey(d, v, "rekey")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rekey"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw"); ok || d.HasChange("remote_gw") {
		t, err := expandVpnIpsecPhase1InterfaceRemoteGw(d, v, "remote_gw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw_country"); ok || d.HasChange("remote_gw_country") {
		t, err := expandVpnIpsecPhase1InterfaceRemoteGwCountry(d, v, "remote_gw_country")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw-country"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw_end_ip"); ok || d.HasChange("remote_gw_end_ip") {
		t, err := expandVpnIpsecPhase1InterfaceRemoteGwEndIp(d, v, "remote_gw_end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw-end-ip"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw_match"); ok || d.HasChange("remote_gw_match") {
		t, err := expandVpnIpsecPhase1InterfaceRemoteGwMatch(d, v, "remote_gw_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw-match"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw_start_ip"); ok || d.HasChange("remote_gw_start_ip") {
		t, err := expandVpnIpsecPhase1InterfaceRemoteGwStartIp(d, v, "remote_gw_start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw-start-ip"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw_subnet"); ok || d.HasChange("remote_gw_subnet") {
		t, err := expandVpnIpsecPhase1InterfaceRemoteGwSubnet(d, v, "remote_gw_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw-subnet"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw_ztna_tags"); ok || d.HasChange("remote_gw_ztna_tags") {
		t, err := expandVpnIpsecPhase1InterfaceRemoteGwZtnaTags(d, v, "remote_gw_ztna_tags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw-ztna-tags"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw6"); ok || d.HasChange("remote_gw6") {
		t, err := expandVpnIpsecPhase1InterfaceRemoteGw6(d, v, "remote_gw6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw6"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw6_country"); ok || d.HasChange("remote_gw6_country") {
		t, err := expandVpnIpsecPhase1InterfaceRemoteGw6Country(d, v, "remote_gw6_country")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw6-country"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw6_end_ip"); ok || d.HasChange("remote_gw6_end_ip") {
		t, err := expandVpnIpsecPhase1InterfaceRemoteGw6EndIp(d, v, "remote_gw6_end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw6-end-ip"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw6_match"); ok || d.HasChange("remote_gw6_match") {
		t, err := expandVpnIpsecPhase1InterfaceRemoteGw6Match(d, v, "remote_gw6_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw6-match"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw6_start_ip"); ok || d.HasChange("remote_gw6_start_ip") {
		t, err := expandVpnIpsecPhase1InterfaceRemoteGw6StartIp(d, v, "remote_gw6_start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw6-start-ip"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw6_subnet"); ok || d.HasChange("remote_gw6_subnet") {
		t, err := expandVpnIpsecPhase1InterfaceRemoteGw6Subnet(d, v, "remote_gw6_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw6-subnet"] = t
		}
	}

	if v, ok := d.GetOk("remotegw_ddns"); ok || d.HasChange("remotegw_ddns") {
		t, err := expandVpnIpsecPhase1InterfaceRemotegwDdns(d, v, "remotegw_ddns")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remotegw-ddns"] = t
		}
	}

	if v, ok := d.GetOk("rsa_signature_format"); ok || d.HasChange("rsa_signature_format") {
		t, err := expandVpnIpsecPhase1InterfaceRsaSignatureFormat(d, v, "rsa_signature_format")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsa-signature-format"] = t
		}
	}

	if v, ok := d.GetOk("rsa_signature_hash_override"); ok || d.HasChange("rsa_signature_hash_override") {
		t, err := expandVpnIpsecPhase1InterfaceRsaSignatureHashOverride(d, v, "rsa_signature_hash_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsa-signature-hash-override"] = t
		}
	}

	if v, ok := d.GetOk("save_password"); ok || d.HasChange("save_password") {
		t, err := expandVpnIpsecPhase1InterfaceSavePassword(d, v, "save_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["save-password"] = t
		}
	}

	if v, ok := d.GetOk("send_cert_chain"); ok || d.HasChange("send_cert_chain") {
		t, err := expandVpnIpsecPhase1InterfaceSendCertChain(d, v, "send_cert_chain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-cert-chain"] = t
		}
	}

	if v, ok := d.GetOk("shared_idle_timeout"); ok || d.HasChange("shared_idle_timeout") {
		t, err := expandVpnIpsecPhase1InterfaceSharedIdleTimeout(d, v, "shared_idle_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shared-idle-timeout"] = t
		}
	}

	if v, ok := d.GetOk("signature_hash_alg"); ok || d.HasChange("signature_hash_alg") {
		t, err := expandVpnIpsecPhase1InterfaceSignatureHashAlg(d, v, "signature_hash_alg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signature-hash-alg"] = t
		}
	}

	if v, ok := d.GetOk("split_include_service"); ok || d.HasChange("split_include_service") {
		t, err := expandVpnIpsecPhase1InterfaceSplitIncludeService(d, v, "split_include_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-include-service"] = t
		}
	}

	if v, ok := d.GetOk("suite_b"); ok || d.HasChange("suite_b") {
		t, err := expandVpnIpsecPhase1InterfaceSuiteB(d, v, "suite_b")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["suite-b"] = t
		}
	}

	if v, ok := d.GetOk("transit_gateway"); ok || d.HasChange("transit_gateway") {
		t, err := expandVpnIpsecPhase1InterfaceTransitGateway(d, v, "transit_gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transit-gateway"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_search"); ok || d.HasChange("tunnel_search") {
		t, err := expandVpnIpsecPhase1InterfaceTunnelSearch(d, v, "tunnel_search")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-search"] = t
		}
	}

	if v, ok := d.GetOk("transport"); ok || d.HasChange("transport") {
		t, err := expandVpnIpsecPhase1InterfaceTransport(d, v, "transport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transport"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandVpnIpsecPhase1InterfaceType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("unity_support"); ok || d.HasChange("unity_support") {
		t, err := expandVpnIpsecPhase1InterfaceUnitySupport(d, v, "unity_support")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unity-support"] = t
		}
	}

	if v, ok := d.GetOk("usrgrp"); ok || d.HasChange("usrgrp") {
		t, err := expandVpnIpsecPhase1InterfaceUsrgrp(d, v, "usrgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["usrgrp"] = t
		}
	}

	if v, ok := d.GetOk("vni"); ok || d.HasChange("vni") {
		t, err := expandVpnIpsecPhase1InterfaceVni(d, v, "vni")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vni"] = t
		}
	}

	if v, ok := d.GetOk("wizard_type"); ok || d.HasChange("wizard_type") {
		t, err := expandVpnIpsecPhase1InterfaceWizardType(d, v, "wizard_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wizard-type"] = t
		}
	}

	if v, ok := d.GetOk("xauthtype"); ok || d.HasChange("xauthtype") {
		t, err := expandVpnIpsecPhase1InterfaceXauthtype(d, v, "xauthtype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["xauthtype"] = t
		}
	}

	return &obj, nil
}
