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

func resourceVpnIpsecPhase1() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnIpsecPhase1Create,
		Read:   resourceVpnIpsecPhase1Read,
		Update: resourceVpnIpsecPhase1Update,
		Delete: resourceVpnIpsecPhase1Delete,

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
				Computed: true,
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
			"assign_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			},
			"client_keep_alive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"dev_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dev_id_notification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"distance": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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
				Computed: true,
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
			},
			"ipv6_dns_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_dns_server3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
				Computed: true,
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
				Computed: true,
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
				Computed: true,
			},
			"remote_gw_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_gw_start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"remote_gw6_country": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_gw6_end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_gw6_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_gw6_start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_gw6_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"transport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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

func resourceVpnIpsecPhase1Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

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

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectVpnIpsecPhase1(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecPhase1 resource while getting object: %v", err)
	}

	_, err = c.CreateVpnIpsecPhase1(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecPhase1 resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceVpnIpsecPhase1Read(d, m)
}

func resourceVpnIpsecPhase1Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

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

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectVpnIpsecPhase1(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecPhase1 resource while getting object: %v", err)
	}

	_, err = c.UpdateVpnIpsecPhase1(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecPhase1 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceVpnIpsecPhase1Read(d, m)
}

func resourceVpnIpsecPhase1Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

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

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteVpnIpsecPhase1(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VpnIpsecPhase1 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnIpsecPhase1Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnIpsecPhase1(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecPhase1 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnIpsecPhase1(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecPhase1 resource from API: %v", err)
	}
	return nil
}

func flattenVpnIpsecPhase1AcctVerify(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1AddGwRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1AddRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Addke1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1Addke2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1Addke3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1Addke4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1Addke5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1Addke6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1Addke7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1AssignIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1AssignIpFrom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Authmethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1AuthmethodRemote(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Authusr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Authusrgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1AutoNegotiate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1AutoTransportThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1AzureAdAutoconnect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1BackupGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1Banner(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1CertIdValidation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1CertPeerUsernameStrip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1CertPeerUsernameValidation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1CertTrustStore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Certificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1ChildlessIke(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1ClientAutoNegotiate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1ClientKeepAlive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1ClientResume(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1ClientResumeInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Comments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1DevId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1DevIdNotification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1DhcpRaGiaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Dhcp6RaLinkaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Dhgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1DigitalSignatureAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Distance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1DnsMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Domain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Dpd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1DpdRetrycount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1DpdRetryinterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenVpnIpsecPhase1Eap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1EapCertAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1EapExcludePeergrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1EapIdentity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1EmsSnCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1EnforceUniqueId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Esn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1ExchangeFgtDeviceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1FallbackTcpThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1FecBase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1FecCodec(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1FecEgress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1FecHealthCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1FecIngress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1FecMappingProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1FecReceiveTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1FecRedundant(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1FecSendTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1FgspSync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1ForticlientEnforcement(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1FortinetEsp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Fragmentation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1FragmentationMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1GroupAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1HaSyncEspSeqno(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1IdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1IdleTimeoutinterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1IkeVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InboundDscpCopy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1IncludeLocalLan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Interface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1InternalDomainList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1IpDelayInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4DnsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4DnsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4DnsServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4EndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4ExcludeRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenVpnIpsecPhase1Ipv4ExcludeRangeEndIp(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "VpnIpsecPhase1-Ipv4ExcludeRange-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenVpnIpsecPhase1Ipv4ExcludeRangeId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "VpnIpsecPhase1-Ipv4ExcludeRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenVpnIpsecPhase1Ipv4ExcludeRangeStartIp(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "VpnIpsecPhase1-Ipv4ExcludeRange-StartIp")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenVpnIpsecPhase1Ipv4ExcludeRangeEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4ExcludeRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4ExcludeRangeStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1Ipv4Netmask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4SplitExclude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1Ipv4SplitInclude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1Ipv4StartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4WinsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4WinsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6AutoLinklocal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6DnsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6DnsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6DnsServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6EndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6ExcludeRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenVpnIpsecPhase1Ipv6ExcludeRangeEndIp(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "VpnIpsecPhase1-Ipv6ExcludeRange-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenVpnIpsecPhase1Ipv6ExcludeRangeId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "VpnIpsecPhase1-Ipv6ExcludeRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenVpnIpsecPhase1Ipv6ExcludeRangeStartIp(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "VpnIpsecPhase1-Ipv6ExcludeRange-StartIp")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenVpnIpsecPhase1Ipv6ExcludeRangeEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6ExcludeRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6ExcludeRangeStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1Ipv6Prefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6SplitExclude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1Ipv6SplitInclude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1Ipv6StartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Keepalive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Keylife(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Kms(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1LinkCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1LocalGw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Localid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1LocalidType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1LoopbackAsymroute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1MeshSelectorType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1ModeCfg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1ModeCfgAllowClientSelector(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Nattraversal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1NegotiateTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1NetworkId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1NetworkOverlay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1NpuOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Peer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1Peergrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1Peerid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Peertype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ppk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1PpkIdentity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Priority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Proposal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1Qkd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1QkdProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1Reauth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Rekey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1RemoteGw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1RemoteGwCountry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1RemoteGwEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1RemoteGwMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1RemoteGwStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1RemoteGwSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1RemoteGwZtnaTags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1RemoteGw6Country(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1RemoteGw6EndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1RemoteGw6Match(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1RemoteGw6StartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1RemoteGw6Subnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1RemotegwDdns(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1RsaSignatureFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1RsaSignatureHashOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1SavePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1SendCertChain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1SharedIdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1SignatureHashAlg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1SplitIncludeService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1SuiteB(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1TransitGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Transport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Type(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1UnitySupport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Usrgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase1WizardType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Xauthtype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnIpsecPhase1(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("acct_verify", flattenVpnIpsecPhase1AcctVerify(o["acct-verify"], d, "acct_verify")); err != nil {
		if vv, ok := fortiAPIPatch(o["acct-verify"], "VpnIpsecPhase1-AcctVerify"); ok {
			if err = d.Set("acct_verify", vv); err != nil {
				return fmt.Errorf("Error reading acct_verify: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading acct_verify: %v", err)
		}
	}

	if err = d.Set("add_gw_route", flattenVpnIpsecPhase1AddGwRoute(o["add-gw-route"], d, "add_gw_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["add-gw-route"], "VpnIpsecPhase1-AddGwRoute"); ok {
			if err = d.Set("add_gw_route", vv); err != nil {
				return fmt.Errorf("Error reading add_gw_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading add_gw_route: %v", err)
		}
	}

	if err = d.Set("add_route", flattenVpnIpsecPhase1AddRoute(o["add-route"], d, "add_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["add-route"], "VpnIpsecPhase1-AddRoute"); ok {
			if err = d.Set("add_route", vv); err != nil {
				return fmt.Errorf("Error reading add_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading add_route: %v", err)
		}
	}

	if err = d.Set("addke1", flattenVpnIpsecPhase1Addke1(o["addke1"], d, "addke1")); err != nil {
		if vv, ok := fortiAPIPatch(o["addke1"], "VpnIpsecPhase1-Addke1"); ok {
			if err = d.Set("addke1", vv); err != nil {
				return fmt.Errorf("Error reading addke1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addke1: %v", err)
		}
	}

	if err = d.Set("addke2", flattenVpnIpsecPhase1Addke2(o["addke2"], d, "addke2")); err != nil {
		if vv, ok := fortiAPIPatch(o["addke2"], "VpnIpsecPhase1-Addke2"); ok {
			if err = d.Set("addke2", vv); err != nil {
				return fmt.Errorf("Error reading addke2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addke2: %v", err)
		}
	}

	if err = d.Set("addke3", flattenVpnIpsecPhase1Addke3(o["addke3"], d, "addke3")); err != nil {
		if vv, ok := fortiAPIPatch(o["addke3"], "VpnIpsecPhase1-Addke3"); ok {
			if err = d.Set("addke3", vv); err != nil {
				return fmt.Errorf("Error reading addke3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addke3: %v", err)
		}
	}

	if err = d.Set("addke4", flattenVpnIpsecPhase1Addke4(o["addke4"], d, "addke4")); err != nil {
		if vv, ok := fortiAPIPatch(o["addke4"], "VpnIpsecPhase1-Addke4"); ok {
			if err = d.Set("addke4", vv); err != nil {
				return fmt.Errorf("Error reading addke4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addke4: %v", err)
		}
	}

	if err = d.Set("addke5", flattenVpnIpsecPhase1Addke5(o["addke5"], d, "addke5")); err != nil {
		if vv, ok := fortiAPIPatch(o["addke5"], "VpnIpsecPhase1-Addke5"); ok {
			if err = d.Set("addke5", vv); err != nil {
				return fmt.Errorf("Error reading addke5: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addke5: %v", err)
		}
	}

	if err = d.Set("addke6", flattenVpnIpsecPhase1Addke6(o["addke6"], d, "addke6")); err != nil {
		if vv, ok := fortiAPIPatch(o["addke6"], "VpnIpsecPhase1-Addke6"); ok {
			if err = d.Set("addke6", vv); err != nil {
				return fmt.Errorf("Error reading addke6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addke6: %v", err)
		}
	}

	if err = d.Set("addke7", flattenVpnIpsecPhase1Addke7(o["addke7"], d, "addke7")); err != nil {
		if vv, ok := fortiAPIPatch(o["addke7"], "VpnIpsecPhase1-Addke7"); ok {
			if err = d.Set("addke7", vv); err != nil {
				return fmt.Errorf("Error reading addke7: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addke7: %v", err)
		}
	}

	if err = d.Set("assign_ip", flattenVpnIpsecPhase1AssignIp(o["assign-ip"], d, "assign_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["assign-ip"], "VpnIpsecPhase1-AssignIp"); ok {
			if err = d.Set("assign_ip", vv); err != nil {
				return fmt.Errorf("Error reading assign_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading assign_ip: %v", err)
		}
	}

	if err = d.Set("assign_ip_from", flattenVpnIpsecPhase1AssignIpFrom(o["assign-ip-from"], d, "assign_ip_from")); err != nil {
		if vv, ok := fortiAPIPatch(o["assign-ip-from"], "VpnIpsecPhase1-AssignIpFrom"); ok {
			if err = d.Set("assign_ip_from", vv); err != nil {
				return fmt.Errorf("Error reading assign_ip_from: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading assign_ip_from: %v", err)
		}
	}

	if err = d.Set("authmethod", flattenVpnIpsecPhase1Authmethod(o["authmethod"], d, "authmethod")); err != nil {
		if vv, ok := fortiAPIPatch(o["authmethod"], "VpnIpsecPhase1-Authmethod"); ok {
			if err = d.Set("authmethod", vv); err != nil {
				return fmt.Errorf("Error reading authmethod: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authmethod: %v", err)
		}
	}

	if err = d.Set("authmethod_remote", flattenVpnIpsecPhase1AuthmethodRemote(o["authmethod-remote"], d, "authmethod_remote")); err != nil {
		if vv, ok := fortiAPIPatch(o["authmethod-remote"], "VpnIpsecPhase1-AuthmethodRemote"); ok {
			if err = d.Set("authmethod_remote", vv); err != nil {
				return fmt.Errorf("Error reading authmethod_remote: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authmethod_remote: %v", err)
		}
	}

	if err = d.Set("authusr", flattenVpnIpsecPhase1Authusr(o["authusr"], d, "authusr")); err != nil {
		if vv, ok := fortiAPIPatch(o["authusr"], "VpnIpsecPhase1-Authusr"); ok {
			if err = d.Set("authusr", vv); err != nil {
				return fmt.Errorf("Error reading authusr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authusr: %v", err)
		}
	}

	if err = d.Set("authusrgrp", flattenVpnIpsecPhase1Authusrgrp(o["authusrgrp"], d, "authusrgrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["authusrgrp"], "VpnIpsecPhase1-Authusrgrp"); ok {
			if err = d.Set("authusrgrp", vv); err != nil {
				return fmt.Errorf("Error reading authusrgrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authusrgrp: %v", err)
		}
	}

	if err = d.Set("auto_negotiate", flattenVpnIpsecPhase1AutoNegotiate(o["auto-negotiate"], d, "auto_negotiate")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-negotiate"], "VpnIpsecPhase1-AutoNegotiate"); ok {
			if err = d.Set("auto_negotiate", vv); err != nil {
				return fmt.Errorf("Error reading auto_negotiate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_negotiate: %v", err)
		}
	}

	if err = d.Set("auto_transport_threshold", flattenVpnIpsecPhase1AutoTransportThreshold(o["auto-transport-threshold"], d, "auto_transport_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-transport-threshold"], "VpnIpsecPhase1-AutoTransportThreshold"); ok {
			if err = d.Set("auto_transport_threshold", vv); err != nil {
				return fmt.Errorf("Error reading auto_transport_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_transport_threshold: %v", err)
		}
	}

	if err = d.Set("azure_ad_autoconnect", flattenVpnIpsecPhase1AzureAdAutoconnect(o["azure-ad-autoconnect"], d, "azure_ad_autoconnect")); err != nil {
		if vv, ok := fortiAPIPatch(o["azure-ad-autoconnect"], "VpnIpsecPhase1-AzureAdAutoconnect"); ok {
			if err = d.Set("azure_ad_autoconnect", vv); err != nil {
				return fmt.Errorf("Error reading azure_ad_autoconnect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading azure_ad_autoconnect: %v", err)
		}
	}

	if err = d.Set("backup_gateway", flattenVpnIpsecPhase1BackupGateway(o["backup-gateway"], d, "backup_gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["backup-gateway"], "VpnIpsecPhase1-BackupGateway"); ok {
			if err = d.Set("backup_gateway", vv); err != nil {
				return fmt.Errorf("Error reading backup_gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading backup_gateway: %v", err)
		}
	}

	if err = d.Set("banner", flattenVpnIpsecPhase1Banner(o["banner"], d, "banner")); err != nil {
		if vv, ok := fortiAPIPatch(o["banner"], "VpnIpsecPhase1-Banner"); ok {
			if err = d.Set("banner", vv); err != nil {
				return fmt.Errorf("Error reading banner: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading banner: %v", err)
		}
	}

	if err = d.Set("cert_id_validation", flattenVpnIpsecPhase1CertIdValidation(o["cert-id-validation"], d, "cert_id_validation")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert-id-validation"], "VpnIpsecPhase1-CertIdValidation"); ok {
			if err = d.Set("cert_id_validation", vv); err != nil {
				return fmt.Errorf("Error reading cert_id_validation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert_id_validation: %v", err)
		}
	}

	if err = d.Set("cert_peer_username_strip", flattenVpnIpsecPhase1CertPeerUsernameStrip(o["cert-peer-username-strip"], d, "cert_peer_username_strip")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert-peer-username-strip"], "VpnIpsecPhase1-CertPeerUsernameStrip"); ok {
			if err = d.Set("cert_peer_username_strip", vv); err != nil {
				return fmt.Errorf("Error reading cert_peer_username_strip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert_peer_username_strip: %v", err)
		}
	}

	if err = d.Set("cert_peer_username_validation", flattenVpnIpsecPhase1CertPeerUsernameValidation(o["cert-peer-username-validation"], d, "cert_peer_username_validation")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert-peer-username-validation"], "VpnIpsecPhase1-CertPeerUsernameValidation"); ok {
			if err = d.Set("cert_peer_username_validation", vv); err != nil {
				return fmt.Errorf("Error reading cert_peer_username_validation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert_peer_username_validation: %v", err)
		}
	}

	if err = d.Set("cert_trust_store", flattenVpnIpsecPhase1CertTrustStore(o["cert-trust-store"], d, "cert_trust_store")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert-trust-store"], "VpnIpsecPhase1-CertTrustStore"); ok {
			if err = d.Set("cert_trust_store", vv); err != nil {
				return fmt.Errorf("Error reading cert_trust_store: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert_trust_store: %v", err)
		}
	}

	if err = d.Set("certificate", flattenVpnIpsecPhase1Certificate(o["certificate"], d, "certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["certificate"], "VpnIpsecPhase1-Certificate"); ok {
			if err = d.Set("certificate", vv); err != nil {
				return fmt.Errorf("Error reading certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("childless_ike", flattenVpnIpsecPhase1ChildlessIke(o["childless-ike"], d, "childless_ike")); err != nil {
		if vv, ok := fortiAPIPatch(o["childless-ike"], "VpnIpsecPhase1-ChildlessIke"); ok {
			if err = d.Set("childless_ike", vv); err != nil {
				return fmt.Errorf("Error reading childless_ike: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading childless_ike: %v", err)
		}
	}

	if err = d.Set("client_auto_negotiate", flattenVpnIpsecPhase1ClientAutoNegotiate(o["client-auto-negotiate"], d, "client_auto_negotiate")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-auto-negotiate"], "VpnIpsecPhase1-ClientAutoNegotiate"); ok {
			if err = d.Set("client_auto_negotiate", vv); err != nil {
				return fmt.Errorf("Error reading client_auto_negotiate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_auto_negotiate: %v", err)
		}
	}

	if err = d.Set("client_keep_alive", flattenVpnIpsecPhase1ClientKeepAlive(o["client-keep-alive"], d, "client_keep_alive")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-keep-alive"], "VpnIpsecPhase1-ClientKeepAlive"); ok {
			if err = d.Set("client_keep_alive", vv); err != nil {
				return fmt.Errorf("Error reading client_keep_alive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_keep_alive: %v", err)
		}
	}

	if err = d.Set("client_resume", flattenVpnIpsecPhase1ClientResume(o["client-resume"], d, "client_resume")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-resume"], "VpnIpsecPhase1-ClientResume"); ok {
			if err = d.Set("client_resume", vv); err != nil {
				return fmt.Errorf("Error reading client_resume: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_resume: %v", err)
		}
	}

	if err = d.Set("client_resume_interval", flattenVpnIpsecPhase1ClientResumeInterval(o["client-resume-interval"], d, "client_resume_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-resume-interval"], "VpnIpsecPhase1-ClientResumeInterval"); ok {
			if err = d.Set("client_resume_interval", vv); err != nil {
				return fmt.Errorf("Error reading client_resume_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_resume_interval: %v", err)
		}
	}

	if err = d.Set("comments", flattenVpnIpsecPhase1Comments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "VpnIpsecPhase1-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("dev_id", flattenVpnIpsecPhase1DevId(o["dev-id"], d, "dev_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["dev-id"], "VpnIpsecPhase1-DevId"); ok {
			if err = d.Set("dev_id", vv); err != nil {
				return fmt.Errorf("Error reading dev_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dev_id: %v", err)
		}
	}

	if err = d.Set("dev_id_notification", flattenVpnIpsecPhase1DevIdNotification(o["dev-id-notification"], d, "dev_id_notification")); err != nil {
		if vv, ok := fortiAPIPatch(o["dev-id-notification"], "VpnIpsecPhase1-DevIdNotification"); ok {
			if err = d.Set("dev_id_notification", vv); err != nil {
				return fmt.Errorf("Error reading dev_id_notification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dev_id_notification: %v", err)
		}
	}

	if err = d.Set("dhcp_ra_giaddr", flattenVpnIpsecPhase1DhcpRaGiaddr(o["dhcp-ra-giaddr"], d, "dhcp_ra_giaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-ra-giaddr"], "VpnIpsecPhase1-DhcpRaGiaddr"); ok {
			if err = d.Set("dhcp_ra_giaddr", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_ra_giaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_ra_giaddr: %v", err)
		}
	}

	if err = d.Set("dhcp6_ra_linkaddr", flattenVpnIpsecPhase1Dhcp6RaLinkaddr(o["dhcp6-ra-linkaddr"], d, "dhcp6_ra_linkaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp6-ra-linkaddr"], "VpnIpsecPhase1-Dhcp6RaLinkaddr"); ok {
			if err = d.Set("dhcp6_ra_linkaddr", vv); err != nil {
				return fmt.Errorf("Error reading dhcp6_ra_linkaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp6_ra_linkaddr: %v", err)
		}
	}

	if err = d.Set("dhgrp", flattenVpnIpsecPhase1Dhgrp(o["dhgrp"], d, "dhgrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhgrp"], "VpnIpsecPhase1-Dhgrp"); ok {
			if err = d.Set("dhgrp", vv); err != nil {
				return fmt.Errorf("Error reading dhgrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhgrp: %v", err)
		}
	}

	if err = d.Set("digital_signature_auth", flattenVpnIpsecPhase1DigitalSignatureAuth(o["digital-signature-auth"], d, "digital_signature_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["digital-signature-auth"], "VpnIpsecPhase1-DigitalSignatureAuth"); ok {
			if err = d.Set("digital_signature_auth", vv); err != nil {
				return fmt.Errorf("Error reading digital_signature_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading digital_signature_auth: %v", err)
		}
	}

	if err = d.Set("distance", flattenVpnIpsecPhase1Distance(o["distance"], d, "distance")); err != nil {
		if vv, ok := fortiAPIPatch(o["distance"], "VpnIpsecPhase1-Distance"); ok {
			if err = d.Set("distance", vv); err != nil {
				return fmt.Errorf("Error reading distance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	if err = d.Set("dns_mode", flattenVpnIpsecPhase1DnsMode(o["dns-mode"], d, "dns_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-mode"], "VpnIpsecPhase1-DnsMode"); ok {
			if err = d.Set("dns_mode", vv); err != nil {
				return fmt.Errorf("Error reading dns_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_mode: %v", err)
		}
	}

	if err = d.Set("domain", flattenVpnIpsecPhase1Domain(o["domain"], d, "domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain"], "VpnIpsecPhase1-Domain"); ok {
			if err = d.Set("domain", vv); err != nil {
				return fmt.Errorf("Error reading domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("dpd", flattenVpnIpsecPhase1Dpd(o["dpd"], d, "dpd")); err != nil {
		if vv, ok := fortiAPIPatch(o["dpd"], "VpnIpsecPhase1-Dpd"); ok {
			if err = d.Set("dpd", vv); err != nil {
				return fmt.Errorf("Error reading dpd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dpd: %v", err)
		}
	}

	if err = d.Set("dpd_retrycount", flattenVpnIpsecPhase1DpdRetrycount(o["dpd-retrycount"], d, "dpd_retrycount")); err != nil {
		if vv, ok := fortiAPIPatch(o["dpd-retrycount"], "VpnIpsecPhase1-DpdRetrycount"); ok {
			if err = d.Set("dpd_retrycount", vv); err != nil {
				return fmt.Errorf("Error reading dpd_retrycount: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dpd_retrycount: %v", err)
		}
	}

	if err = d.Set("dpd_retryinterval", flattenVpnIpsecPhase1DpdRetryinterval(o["dpd-retryinterval"], d, "dpd_retryinterval")); err != nil {
		if vv, ok := fortiAPIPatch(o["dpd-retryinterval"], "VpnIpsecPhase1-DpdRetryinterval"); ok {
			if err = d.Set("dpd_retryinterval", vv); err != nil {
				return fmt.Errorf("Error reading dpd_retryinterval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dpd_retryinterval: %v", err)
		}
	}

	if err = d.Set("eap", flattenVpnIpsecPhase1Eap(o["eap"], d, "eap")); err != nil {
		if vv, ok := fortiAPIPatch(o["eap"], "VpnIpsecPhase1-Eap"); ok {
			if err = d.Set("eap", vv); err != nil {
				return fmt.Errorf("Error reading eap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eap: %v", err)
		}
	}

	if err = d.Set("eap_cert_auth", flattenVpnIpsecPhase1EapCertAuth(o["eap-cert-auth"], d, "eap_cert_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["eap-cert-auth"], "VpnIpsecPhase1-EapCertAuth"); ok {
			if err = d.Set("eap_cert_auth", vv); err != nil {
				return fmt.Errorf("Error reading eap_cert_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eap_cert_auth: %v", err)
		}
	}

	if err = d.Set("eap_exclude_peergrp", flattenVpnIpsecPhase1EapExcludePeergrp(o["eap-exclude-peergrp"], d, "eap_exclude_peergrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["eap-exclude-peergrp"], "VpnIpsecPhase1-EapExcludePeergrp"); ok {
			if err = d.Set("eap_exclude_peergrp", vv); err != nil {
				return fmt.Errorf("Error reading eap_exclude_peergrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eap_exclude_peergrp: %v", err)
		}
	}

	if err = d.Set("eap_identity", flattenVpnIpsecPhase1EapIdentity(o["eap-identity"], d, "eap_identity")); err != nil {
		if vv, ok := fortiAPIPatch(o["eap-identity"], "VpnIpsecPhase1-EapIdentity"); ok {
			if err = d.Set("eap_identity", vv); err != nil {
				return fmt.Errorf("Error reading eap_identity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eap_identity: %v", err)
		}
	}

	if err = d.Set("ems_sn_check", flattenVpnIpsecPhase1EmsSnCheck(o["ems-sn-check"], d, "ems_sn_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["ems-sn-check"], "VpnIpsecPhase1-EmsSnCheck"); ok {
			if err = d.Set("ems_sn_check", vv); err != nil {
				return fmt.Errorf("Error reading ems_sn_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ems_sn_check: %v", err)
		}
	}

	if err = d.Set("enforce_unique_id", flattenVpnIpsecPhase1EnforceUniqueId(o["enforce-unique-id"], d, "enforce_unique_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["enforce-unique-id"], "VpnIpsecPhase1-EnforceUniqueId"); ok {
			if err = d.Set("enforce_unique_id", vv); err != nil {
				return fmt.Errorf("Error reading enforce_unique_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enforce_unique_id: %v", err)
		}
	}

	if err = d.Set("esn", flattenVpnIpsecPhase1Esn(o["esn"], d, "esn")); err != nil {
		if vv, ok := fortiAPIPatch(o["esn"], "VpnIpsecPhase1-Esn"); ok {
			if err = d.Set("esn", vv); err != nil {
				return fmt.Errorf("Error reading esn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading esn: %v", err)
		}
	}

	if err = d.Set("exchange_fgt_device_id", flattenVpnIpsecPhase1ExchangeFgtDeviceId(o["exchange-fgt-device-id"], d, "exchange_fgt_device_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["exchange-fgt-device-id"], "VpnIpsecPhase1-ExchangeFgtDeviceId"); ok {
			if err = d.Set("exchange_fgt_device_id", vv); err != nil {
				return fmt.Errorf("Error reading exchange_fgt_device_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exchange_fgt_device_id: %v", err)
		}
	}

	if err = d.Set("fallback_tcp_threshold", flattenVpnIpsecPhase1FallbackTcpThreshold(o["fallback-tcp-threshold"], d, "fallback_tcp_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["fallback-tcp-threshold"], "VpnIpsecPhase1-FallbackTcpThreshold"); ok {
			if err = d.Set("fallback_tcp_threshold", vv); err != nil {
				return fmt.Errorf("Error reading fallback_tcp_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fallback_tcp_threshold: %v", err)
		}
	}

	if err = d.Set("fec_base", flattenVpnIpsecPhase1FecBase(o["fec-base"], d, "fec_base")); err != nil {
		if vv, ok := fortiAPIPatch(o["fec-base"], "VpnIpsecPhase1-FecBase"); ok {
			if err = d.Set("fec_base", vv); err != nil {
				return fmt.Errorf("Error reading fec_base: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fec_base: %v", err)
		}
	}

	if err = d.Set("fec_codec", flattenVpnIpsecPhase1FecCodec(o["fec-codec"], d, "fec_codec")); err != nil {
		if vv, ok := fortiAPIPatch(o["fec-codec"], "VpnIpsecPhase1-FecCodec"); ok {
			if err = d.Set("fec_codec", vv); err != nil {
				return fmt.Errorf("Error reading fec_codec: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fec_codec: %v", err)
		}
	}

	if err = d.Set("fec_egress", flattenVpnIpsecPhase1FecEgress(o["fec-egress"], d, "fec_egress")); err != nil {
		if vv, ok := fortiAPIPatch(o["fec-egress"], "VpnIpsecPhase1-FecEgress"); ok {
			if err = d.Set("fec_egress", vv); err != nil {
				return fmt.Errorf("Error reading fec_egress: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fec_egress: %v", err)
		}
	}

	if err = d.Set("fec_health_check", flattenVpnIpsecPhase1FecHealthCheck(o["fec-health-check"], d, "fec_health_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["fec-health-check"], "VpnIpsecPhase1-FecHealthCheck"); ok {
			if err = d.Set("fec_health_check", vv); err != nil {
				return fmt.Errorf("Error reading fec_health_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fec_health_check: %v", err)
		}
	}

	if err = d.Set("fec_ingress", flattenVpnIpsecPhase1FecIngress(o["fec-ingress"], d, "fec_ingress")); err != nil {
		if vv, ok := fortiAPIPatch(o["fec-ingress"], "VpnIpsecPhase1-FecIngress"); ok {
			if err = d.Set("fec_ingress", vv); err != nil {
				return fmt.Errorf("Error reading fec_ingress: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fec_ingress: %v", err)
		}
	}

	if err = d.Set("fec_mapping_profile", flattenVpnIpsecPhase1FecMappingProfile(o["fec-mapping-profile"], d, "fec_mapping_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["fec-mapping-profile"], "VpnIpsecPhase1-FecMappingProfile"); ok {
			if err = d.Set("fec_mapping_profile", vv); err != nil {
				return fmt.Errorf("Error reading fec_mapping_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fec_mapping_profile: %v", err)
		}
	}

	if err = d.Set("fec_receive_timeout", flattenVpnIpsecPhase1FecReceiveTimeout(o["fec-receive-timeout"], d, "fec_receive_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["fec-receive-timeout"], "VpnIpsecPhase1-FecReceiveTimeout"); ok {
			if err = d.Set("fec_receive_timeout", vv); err != nil {
				return fmt.Errorf("Error reading fec_receive_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fec_receive_timeout: %v", err)
		}
	}

	if err = d.Set("fec_redundant", flattenVpnIpsecPhase1FecRedundant(o["fec-redundant"], d, "fec_redundant")); err != nil {
		if vv, ok := fortiAPIPatch(o["fec-redundant"], "VpnIpsecPhase1-FecRedundant"); ok {
			if err = d.Set("fec_redundant", vv); err != nil {
				return fmt.Errorf("Error reading fec_redundant: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fec_redundant: %v", err)
		}
	}

	if err = d.Set("fec_send_timeout", flattenVpnIpsecPhase1FecSendTimeout(o["fec-send-timeout"], d, "fec_send_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["fec-send-timeout"], "VpnIpsecPhase1-FecSendTimeout"); ok {
			if err = d.Set("fec_send_timeout", vv); err != nil {
				return fmt.Errorf("Error reading fec_send_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fec_send_timeout: %v", err)
		}
	}

	if err = d.Set("fgsp_sync", flattenVpnIpsecPhase1FgspSync(o["fgsp-sync"], d, "fgsp_sync")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgsp-sync"], "VpnIpsecPhase1-FgspSync"); ok {
			if err = d.Set("fgsp_sync", vv); err != nil {
				return fmt.Errorf("Error reading fgsp_sync: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgsp_sync: %v", err)
		}
	}

	if err = d.Set("forticlient_enforcement", flattenVpnIpsecPhase1ForticlientEnforcement(o["forticlient-enforcement"], d, "forticlient_enforcement")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticlient-enforcement"], "VpnIpsecPhase1-ForticlientEnforcement"); ok {
			if err = d.Set("forticlient_enforcement", vv); err != nil {
				return fmt.Errorf("Error reading forticlient_enforcement: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticlient_enforcement: %v", err)
		}
	}

	if err = d.Set("fortinet_esp", flattenVpnIpsecPhase1FortinetEsp(o["fortinet-esp"], d, "fortinet_esp")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortinet-esp"], "VpnIpsecPhase1-FortinetEsp"); ok {
			if err = d.Set("fortinet_esp", vv); err != nil {
				return fmt.Errorf("Error reading fortinet_esp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortinet_esp: %v", err)
		}
	}

	if err = d.Set("fragmentation", flattenVpnIpsecPhase1Fragmentation(o["fragmentation"], d, "fragmentation")); err != nil {
		if vv, ok := fortiAPIPatch(o["fragmentation"], "VpnIpsecPhase1-Fragmentation"); ok {
			if err = d.Set("fragmentation", vv); err != nil {
				return fmt.Errorf("Error reading fragmentation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fragmentation: %v", err)
		}
	}

	if err = d.Set("fragmentation_mtu", flattenVpnIpsecPhase1FragmentationMtu(o["fragmentation-mtu"], d, "fragmentation_mtu")); err != nil {
		if vv, ok := fortiAPIPatch(o["fragmentation-mtu"], "VpnIpsecPhase1-FragmentationMtu"); ok {
			if err = d.Set("fragmentation_mtu", vv); err != nil {
				return fmt.Errorf("Error reading fragmentation_mtu: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fragmentation_mtu: %v", err)
		}
	}

	if err = d.Set("group_authentication", flattenVpnIpsecPhase1GroupAuthentication(o["group-authentication"], d, "group_authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-authentication"], "VpnIpsecPhase1-GroupAuthentication"); ok {
			if err = d.Set("group_authentication", vv); err != nil {
				return fmt.Errorf("Error reading group_authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_authentication: %v", err)
		}
	}

	if err = d.Set("ha_sync_esp_seqno", flattenVpnIpsecPhase1HaSyncEspSeqno(o["ha-sync-esp-seqno"], d, "ha_sync_esp_seqno")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-sync-esp-seqno"], "VpnIpsecPhase1-HaSyncEspSeqno"); ok {
			if err = d.Set("ha_sync_esp_seqno", vv); err != nil {
				return fmt.Errorf("Error reading ha_sync_esp_seqno: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_sync_esp_seqno: %v", err)
		}
	}

	if err = d.Set("idle_timeout", flattenVpnIpsecPhase1IdleTimeout(o["idle-timeout"], d, "idle_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["idle-timeout"], "VpnIpsecPhase1-IdleTimeout"); ok {
			if err = d.Set("idle_timeout", vv); err != nil {
				return fmt.Errorf("Error reading idle_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idle_timeout: %v", err)
		}
	}

	if err = d.Set("idle_timeoutinterval", flattenVpnIpsecPhase1IdleTimeoutinterval(o["idle-timeoutinterval"], d, "idle_timeoutinterval")); err != nil {
		if vv, ok := fortiAPIPatch(o["idle-timeoutinterval"], "VpnIpsecPhase1-IdleTimeoutinterval"); ok {
			if err = d.Set("idle_timeoutinterval", vv); err != nil {
				return fmt.Errorf("Error reading idle_timeoutinterval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idle_timeoutinterval: %v", err)
		}
	}

	if err = d.Set("ike_version", flattenVpnIpsecPhase1IkeVersion(o["ike-version"], d, "ike_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike-version"], "VpnIpsecPhase1-IkeVersion"); ok {
			if err = d.Set("ike_version", vv); err != nil {
				return fmt.Errorf("Error reading ike_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike_version: %v", err)
		}
	}

	if err = d.Set("inbound_dscp_copy", flattenVpnIpsecPhase1InboundDscpCopy(o["inbound-dscp-copy"], d, "inbound_dscp_copy")); err != nil {
		if vv, ok := fortiAPIPatch(o["inbound-dscp-copy"], "VpnIpsecPhase1-InboundDscpCopy"); ok {
			if err = d.Set("inbound_dscp_copy", vv); err != nil {
				return fmt.Errorf("Error reading inbound_dscp_copy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inbound_dscp_copy: %v", err)
		}
	}

	if err = d.Set("include_local_lan", flattenVpnIpsecPhase1IncludeLocalLan(o["include-local-lan"], d, "include_local_lan")); err != nil {
		if vv, ok := fortiAPIPatch(o["include-local-lan"], "VpnIpsecPhase1-IncludeLocalLan"); ok {
			if err = d.Set("include_local_lan", vv); err != nil {
				return fmt.Errorf("Error reading include_local_lan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading include_local_lan: %v", err)
		}
	}

	if err = d.Set("interface", flattenVpnIpsecPhase1Interface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "VpnIpsecPhase1-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("internal_domain_list", flattenVpnIpsecPhase1InternalDomainList(o["internal-domain-list"], d, "internal_domain_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["internal-domain-list"], "VpnIpsecPhase1-InternalDomainList"); ok {
			if err = d.Set("internal_domain_list", vv); err != nil {
				return fmt.Errorf("Error reading internal_domain_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internal_domain_list: %v", err)
		}
	}

	if err = d.Set("ip_delay_interval", flattenVpnIpsecPhase1IpDelayInterval(o["ip-delay-interval"], d, "ip_delay_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-delay-interval"], "VpnIpsecPhase1-IpDelayInterval"); ok {
			if err = d.Set("ip_delay_interval", vv); err != nil {
				return fmt.Errorf("Error reading ip_delay_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_delay_interval: %v", err)
		}
	}

	if err = d.Set("ipv4_dns_server1", flattenVpnIpsecPhase1Ipv4DnsServer1(o["ipv4-dns-server1"], d, "ipv4_dns_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-dns-server1"], "VpnIpsecPhase1-Ipv4DnsServer1"); ok {
			if err = d.Set("ipv4_dns_server1", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_dns_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_dns_server1: %v", err)
		}
	}

	if err = d.Set("ipv4_dns_server2", flattenVpnIpsecPhase1Ipv4DnsServer2(o["ipv4-dns-server2"], d, "ipv4_dns_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-dns-server2"], "VpnIpsecPhase1-Ipv4DnsServer2"); ok {
			if err = d.Set("ipv4_dns_server2", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_dns_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_dns_server2: %v", err)
		}
	}

	if err = d.Set("ipv4_dns_server3", flattenVpnIpsecPhase1Ipv4DnsServer3(o["ipv4-dns-server3"], d, "ipv4_dns_server3")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-dns-server3"], "VpnIpsecPhase1-Ipv4DnsServer3"); ok {
			if err = d.Set("ipv4_dns_server3", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_dns_server3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_dns_server3: %v", err)
		}
	}

	if err = d.Set("ipv4_end_ip", flattenVpnIpsecPhase1Ipv4EndIp(o["ipv4-end-ip"], d, "ipv4_end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-end-ip"], "VpnIpsecPhase1-Ipv4EndIp"); ok {
			if err = d.Set("ipv4_end_ip", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_end_ip: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ipv4_exclude_range", flattenVpnIpsecPhase1Ipv4ExcludeRange(o["ipv4-exclude-range"], d, "ipv4_exclude_range")); err != nil {
			if vv, ok := fortiAPIPatch(o["ipv4-exclude-range"], "VpnIpsecPhase1-Ipv4ExcludeRange"); ok {
				if err = d.Set("ipv4_exclude_range", vv); err != nil {
					return fmt.Errorf("Error reading ipv4_exclude_range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ipv4_exclude_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ipv4_exclude_range"); ok {
			if err = d.Set("ipv4_exclude_range", flattenVpnIpsecPhase1Ipv4ExcludeRange(o["ipv4-exclude-range"], d, "ipv4_exclude_range")); err != nil {
				if vv, ok := fortiAPIPatch(o["ipv4-exclude-range"], "VpnIpsecPhase1-Ipv4ExcludeRange"); ok {
					if err = d.Set("ipv4_exclude_range", vv); err != nil {
						return fmt.Errorf("Error reading ipv4_exclude_range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ipv4_exclude_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("ipv4_name", flattenVpnIpsecPhase1Ipv4Name(o["ipv4-name"], d, "ipv4_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-name"], "VpnIpsecPhase1-Ipv4Name"); ok {
			if err = d.Set("ipv4_name", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_name: %v", err)
		}
	}

	if err = d.Set("ipv4_netmask", flattenVpnIpsecPhase1Ipv4Netmask(o["ipv4-netmask"], d, "ipv4_netmask")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-netmask"], "VpnIpsecPhase1-Ipv4Netmask"); ok {
			if err = d.Set("ipv4_netmask", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_netmask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_netmask: %v", err)
		}
	}

	if err = d.Set("ipv4_split_exclude", flattenVpnIpsecPhase1Ipv4SplitExclude(o["ipv4-split-exclude"], d, "ipv4_split_exclude")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-split-exclude"], "VpnIpsecPhase1-Ipv4SplitExclude"); ok {
			if err = d.Set("ipv4_split_exclude", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_split_exclude: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_split_exclude: %v", err)
		}
	}

	if err = d.Set("ipv4_split_include", flattenVpnIpsecPhase1Ipv4SplitInclude(o["ipv4-split-include"], d, "ipv4_split_include")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-split-include"], "VpnIpsecPhase1-Ipv4SplitInclude"); ok {
			if err = d.Set("ipv4_split_include", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_split_include: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_split_include: %v", err)
		}
	}

	if err = d.Set("ipv4_start_ip", flattenVpnIpsecPhase1Ipv4StartIp(o["ipv4-start-ip"], d, "ipv4_start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-start-ip"], "VpnIpsecPhase1-Ipv4StartIp"); ok {
			if err = d.Set("ipv4_start_ip", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_start_ip: %v", err)
		}
	}

	if err = d.Set("ipv4_wins_server1", flattenVpnIpsecPhase1Ipv4WinsServer1(o["ipv4-wins-server1"], d, "ipv4_wins_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-wins-server1"], "VpnIpsecPhase1-Ipv4WinsServer1"); ok {
			if err = d.Set("ipv4_wins_server1", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_wins_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_wins_server1: %v", err)
		}
	}

	if err = d.Set("ipv4_wins_server2", flattenVpnIpsecPhase1Ipv4WinsServer2(o["ipv4-wins-server2"], d, "ipv4_wins_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-wins-server2"], "VpnIpsecPhase1-Ipv4WinsServer2"); ok {
			if err = d.Set("ipv4_wins_server2", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_wins_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_wins_server2: %v", err)
		}
	}

	if err = d.Set("ipv6_auto_linklocal", flattenVpnIpsecPhase1Ipv6AutoLinklocal(o["ipv6-auto-linklocal"], d, "ipv6_auto_linklocal")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-auto-linklocal"], "VpnIpsecPhase1-Ipv6AutoLinklocal"); ok {
			if err = d.Set("ipv6_auto_linklocal", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_auto_linklocal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_auto_linklocal: %v", err)
		}
	}

	if err = d.Set("ipv6_dns_server1", flattenVpnIpsecPhase1Ipv6DnsServer1(o["ipv6-dns-server1"], d, "ipv6_dns_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-dns-server1"], "VpnIpsecPhase1-Ipv6DnsServer1"); ok {
			if err = d.Set("ipv6_dns_server1", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_dns_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_dns_server1: %v", err)
		}
	}

	if err = d.Set("ipv6_dns_server2", flattenVpnIpsecPhase1Ipv6DnsServer2(o["ipv6-dns-server2"], d, "ipv6_dns_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-dns-server2"], "VpnIpsecPhase1-Ipv6DnsServer2"); ok {
			if err = d.Set("ipv6_dns_server2", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_dns_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_dns_server2: %v", err)
		}
	}

	if err = d.Set("ipv6_dns_server3", flattenVpnIpsecPhase1Ipv6DnsServer3(o["ipv6-dns-server3"], d, "ipv6_dns_server3")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-dns-server3"], "VpnIpsecPhase1-Ipv6DnsServer3"); ok {
			if err = d.Set("ipv6_dns_server3", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_dns_server3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_dns_server3: %v", err)
		}
	}

	if err = d.Set("ipv6_end_ip", flattenVpnIpsecPhase1Ipv6EndIp(o["ipv6-end-ip"], d, "ipv6_end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-end-ip"], "VpnIpsecPhase1-Ipv6EndIp"); ok {
			if err = d.Set("ipv6_end_ip", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_end_ip: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ipv6_exclude_range", flattenVpnIpsecPhase1Ipv6ExcludeRange(o["ipv6-exclude-range"], d, "ipv6_exclude_range")); err != nil {
			if vv, ok := fortiAPIPatch(o["ipv6-exclude-range"], "VpnIpsecPhase1-Ipv6ExcludeRange"); ok {
				if err = d.Set("ipv6_exclude_range", vv); err != nil {
					return fmt.Errorf("Error reading ipv6_exclude_range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ipv6_exclude_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ipv6_exclude_range"); ok {
			if err = d.Set("ipv6_exclude_range", flattenVpnIpsecPhase1Ipv6ExcludeRange(o["ipv6-exclude-range"], d, "ipv6_exclude_range")); err != nil {
				if vv, ok := fortiAPIPatch(o["ipv6-exclude-range"], "VpnIpsecPhase1-Ipv6ExcludeRange"); ok {
					if err = d.Set("ipv6_exclude_range", vv); err != nil {
						return fmt.Errorf("Error reading ipv6_exclude_range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ipv6_exclude_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("ipv6_name", flattenVpnIpsecPhase1Ipv6Name(o["ipv6-name"], d, "ipv6_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-name"], "VpnIpsecPhase1-Ipv6Name"); ok {
			if err = d.Set("ipv6_name", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_name: %v", err)
		}
	}

	if err = d.Set("ipv6_prefix", flattenVpnIpsecPhase1Ipv6Prefix(o["ipv6-prefix"], d, "ipv6_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-prefix"], "VpnIpsecPhase1-Ipv6Prefix"); ok {
			if err = d.Set("ipv6_prefix", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_prefix: %v", err)
		}
	}

	if err = d.Set("ipv6_split_exclude", flattenVpnIpsecPhase1Ipv6SplitExclude(o["ipv6-split-exclude"], d, "ipv6_split_exclude")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-split-exclude"], "VpnIpsecPhase1-Ipv6SplitExclude"); ok {
			if err = d.Set("ipv6_split_exclude", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_split_exclude: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_split_exclude: %v", err)
		}
	}

	if err = d.Set("ipv6_split_include", flattenVpnIpsecPhase1Ipv6SplitInclude(o["ipv6-split-include"], d, "ipv6_split_include")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-split-include"], "VpnIpsecPhase1-Ipv6SplitInclude"); ok {
			if err = d.Set("ipv6_split_include", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_split_include: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_split_include: %v", err)
		}
	}

	if err = d.Set("ipv6_start_ip", flattenVpnIpsecPhase1Ipv6StartIp(o["ipv6-start-ip"], d, "ipv6_start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-start-ip"], "VpnIpsecPhase1-Ipv6StartIp"); ok {
			if err = d.Set("ipv6_start_ip", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_start_ip: %v", err)
		}
	}

	if err = d.Set("keepalive", flattenVpnIpsecPhase1Keepalive(o["keepalive"], d, "keepalive")); err != nil {
		if vv, ok := fortiAPIPatch(o["keepalive"], "VpnIpsecPhase1-Keepalive"); ok {
			if err = d.Set("keepalive", vv); err != nil {
				return fmt.Errorf("Error reading keepalive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keepalive: %v", err)
		}
	}

	if err = d.Set("keylife", flattenVpnIpsecPhase1Keylife(o["keylife"], d, "keylife")); err != nil {
		if vv, ok := fortiAPIPatch(o["keylife"], "VpnIpsecPhase1-Keylife"); ok {
			if err = d.Set("keylife", vv); err != nil {
				return fmt.Errorf("Error reading keylife: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keylife: %v", err)
		}
	}

	if err = d.Set("kms", flattenVpnIpsecPhase1Kms(o["kms"], d, "kms")); err != nil {
		if vv, ok := fortiAPIPatch(o["kms"], "VpnIpsecPhase1-Kms"); ok {
			if err = d.Set("kms", vv); err != nil {
				return fmt.Errorf("Error reading kms: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading kms: %v", err)
		}
	}

	if err = d.Set("link_cost", flattenVpnIpsecPhase1LinkCost(o["link-cost"], d, "link_cost")); err != nil {
		if vv, ok := fortiAPIPatch(o["link-cost"], "VpnIpsecPhase1-LinkCost"); ok {
			if err = d.Set("link_cost", vv); err != nil {
				return fmt.Errorf("Error reading link_cost: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading link_cost: %v", err)
		}
	}

	if err = d.Set("local_gw", flattenVpnIpsecPhase1LocalGw(o["local-gw"], d, "local_gw")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-gw"], "VpnIpsecPhase1-LocalGw"); ok {
			if err = d.Set("local_gw", vv); err != nil {
				return fmt.Errorf("Error reading local_gw: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_gw: %v", err)
		}
	}

	if err = d.Set("localid", flattenVpnIpsecPhase1Localid(o["localid"], d, "localid")); err != nil {
		if vv, ok := fortiAPIPatch(o["localid"], "VpnIpsecPhase1-Localid"); ok {
			if err = d.Set("localid", vv); err != nil {
				return fmt.Errorf("Error reading localid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading localid: %v", err)
		}
	}

	if err = d.Set("localid_type", flattenVpnIpsecPhase1LocalidType(o["localid-type"], d, "localid_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["localid-type"], "VpnIpsecPhase1-LocalidType"); ok {
			if err = d.Set("localid_type", vv); err != nil {
				return fmt.Errorf("Error reading localid_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading localid_type: %v", err)
		}
	}

	if err = d.Set("loopback_asymroute", flattenVpnIpsecPhase1LoopbackAsymroute(o["loopback-asymroute"], d, "loopback_asymroute")); err != nil {
		if vv, ok := fortiAPIPatch(o["loopback-asymroute"], "VpnIpsecPhase1-LoopbackAsymroute"); ok {
			if err = d.Set("loopback_asymroute", vv); err != nil {
				return fmt.Errorf("Error reading loopback_asymroute: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading loopback_asymroute: %v", err)
		}
	}

	if err = d.Set("mesh_selector_type", flattenVpnIpsecPhase1MeshSelectorType(o["mesh-selector-type"], d, "mesh_selector_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["mesh-selector-type"], "VpnIpsecPhase1-MeshSelectorType"); ok {
			if err = d.Set("mesh_selector_type", vv); err != nil {
				return fmt.Errorf("Error reading mesh_selector_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mesh_selector_type: %v", err)
		}
	}

	if err = d.Set("mode", flattenVpnIpsecPhase1Mode(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "VpnIpsecPhase1-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("mode_cfg", flattenVpnIpsecPhase1ModeCfg(o["mode-cfg"], d, "mode_cfg")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode-cfg"], "VpnIpsecPhase1-ModeCfg"); ok {
			if err = d.Set("mode_cfg", vv); err != nil {
				return fmt.Errorf("Error reading mode_cfg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode_cfg: %v", err)
		}
	}

	if err = d.Set("mode_cfg_allow_client_selector", flattenVpnIpsecPhase1ModeCfgAllowClientSelector(o["mode-cfg-allow-client-selector"], d, "mode_cfg_allow_client_selector")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode-cfg-allow-client-selector"], "VpnIpsecPhase1-ModeCfgAllowClientSelector"); ok {
			if err = d.Set("mode_cfg_allow_client_selector", vv); err != nil {
				return fmt.Errorf("Error reading mode_cfg_allow_client_selector: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode_cfg_allow_client_selector: %v", err)
		}
	}

	if err = d.Set("name", flattenVpnIpsecPhase1Name(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "VpnIpsecPhase1-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("nattraversal", flattenVpnIpsecPhase1Nattraversal(o["nattraversal"], d, "nattraversal")); err != nil {
		if vv, ok := fortiAPIPatch(o["nattraversal"], "VpnIpsecPhase1-Nattraversal"); ok {
			if err = d.Set("nattraversal", vv); err != nil {
				return fmt.Errorf("Error reading nattraversal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nattraversal: %v", err)
		}
	}

	if err = d.Set("negotiate_timeout", flattenVpnIpsecPhase1NegotiateTimeout(o["negotiate-timeout"], d, "negotiate_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["negotiate-timeout"], "VpnIpsecPhase1-NegotiateTimeout"); ok {
			if err = d.Set("negotiate_timeout", vv); err != nil {
				return fmt.Errorf("Error reading negotiate_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading negotiate_timeout: %v", err)
		}
	}

	if err = d.Set("network_id", flattenVpnIpsecPhase1NetworkId(o["network-id"], d, "network_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["network-id"], "VpnIpsecPhase1-NetworkId"); ok {
			if err = d.Set("network_id", vv); err != nil {
				return fmt.Errorf("Error reading network_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading network_id: %v", err)
		}
	}

	if err = d.Set("network_overlay", flattenVpnIpsecPhase1NetworkOverlay(o["network-overlay"], d, "network_overlay")); err != nil {
		if vv, ok := fortiAPIPatch(o["network-overlay"], "VpnIpsecPhase1-NetworkOverlay"); ok {
			if err = d.Set("network_overlay", vv); err != nil {
				return fmt.Errorf("Error reading network_overlay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading network_overlay: %v", err)
		}
	}

	if err = d.Set("npu_offload", flattenVpnIpsecPhase1NpuOffload(o["npu-offload"], d, "npu_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["npu-offload"], "VpnIpsecPhase1-NpuOffload"); ok {
			if err = d.Set("npu_offload", vv); err != nil {
				return fmt.Errorf("Error reading npu_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading npu_offload: %v", err)
		}
	}

	if err = d.Set("peer", flattenVpnIpsecPhase1Peer(o["peer"], d, "peer")); err != nil {
		if vv, ok := fortiAPIPatch(o["peer"], "VpnIpsecPhase1-Peer"); ok {
			if err = d.Set("peer", vv); err != nil {
				return fmt.Errorf("Error reading peer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peer: %v", err)
		}
	}

	if err = d.Set("peergrp", flattenVpnIpsecPhase1Peergrp(o["peergrp"], d, "peergrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["peergrp"], "VpnIpsecPhase1-Peergrp"); ok {
			if err = d.Set("peergrp", vv); err != nil {
				return fmt.Errorf("Error reading peergrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peergrp: %v", err)
		}
	}

	if err = d.Set("peerid", flattenVpnIpsecPhase1Peerid(o["peerid"], d, "peerid")); err != nil {
		if vv, ok := fortiAPIPatch(o["peerid"], "VpnIpsecPhase1-Peerid"); ok {
			if err = d.Set("peerid", vv); err != nil {
				return fmt.Errorf("Error reading peerid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peerid: %v", err)
		}
	}

	if err = d.Set("peertype", flattenVpnIpsecPhase1Peertype(o["peertype"], d, "peertype")); err != nil {
		if vv, ok := fortiAPIPatch(o["peertype"], "VpnIpsecPhase1-Peertype"); ok {
			if err = d.Set("peertype", vv); err != nil {
				return fmt.Errorf("Error reading peertype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peertype: %v", err)
		}
	}

	if err = d.Set("ppk", flattenVpnIpsecPhase1Ppk(o["ppk"], d, "ppk")); err != nil {
		if vv, ok := fortiAPIPatch(o["ppk"], "VpnIpsecPhase1-Ppk"); ok {
			if err = d.Set("ppk", vv); err != nil {
				return fmt.Errorf("Error reading ppk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ppk: %v", err)
		}
	}

	if err = d.Set("ppk_identity", flattenVpnIpsecPhase1PpkIdentity(o["ppk-identity"], d, "ppk_identity")); err != nil {
		if vv, ok := fortiAPIPatch(o["ppk-identity"], "VpnIpsecPhase1-PpkIdentity"); ok {
			if err = d.Set("ppk_identity", vv); err != nil {
				return fmt.Errorf("Error reading ppk_identity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ppk_identity: %v", err)
		}
	}

	if err = d.Set("priority", flattenVpnIpsecPhase1Priority(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "VpnIpsecPhase1-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("proposal", flattenVpnIpsecPhase1Proposal(o["proposal"], d, "proposal")); err != nil {
		if vv, ok := fortiAPIPatch(o["proposal"], "VpnIpsecPhase1-Proposal"); ok {
			if err = d.Set("proposal", vv); err != nil {
				return fmt.Errorf("Error reading proposal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proposal: %v", err)
		}
	}

	if err = d.Set("qkd", flattenVpnIpsecPhase1Qkd(o["qkd"], d, "qkd")); err != nil {
		if vv, ok := fortiAPIPatch(o["qkd"], "VpnIpsecPhase1-Qkd"); ok {
			if err = d.Set("qkd", vv); err != nil {
				return fmt.Errorf("Error reading qkd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading qkd: %v", err)
		}
	}

	if err = d.Set("qkd_profile", flattenVpnIpsecPhase1QkdProfile(o["qkd-profile"], d, "qkd_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["qkd-profile"], "VpnIpsecPhase1-QkdProfile"); ok {
			if err = d.Set("qkd_profile", vv); err != nil {
				return fmt.Errorf("Error reading qkd_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading qkd_profile: %v", err)
		}
	}

	if err = d.Set("reauth", flattenVpnIpsecPhase1Reauth(o["reauth"], d, "reauth")); err != nil {
		if vv, ok := fortiAPIPatch(o["reauth"], "VpnIpsecPhase1-Reauth"); ok {
			if err = d.Set("reauth", vv); err != nil {
				return fmt.Errorf("Error reading reauth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reauth: %v", err)
		}
	}

	if err = d.Set("rekey", flattenVpnIpsecPhase1Rekey(o["rekey"], d, "rekey")); err != nil {
		if vv, ok := fortiAPIPatch(o["rekey"], "VpnIpsecPhase1-Rekey"); ok {
			if err = d.Set("rekey", vv); err != nil {
				return fmt.Errorf("Error reading rekey: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rekey: %v", err)
		}
	}

	if err = d.Set("remote_gw", flattenVpnIpsecPhase1RemoteGw(o["remote-gw"], d, "remote_gw")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-gw"], "VpnIpsecPhase1-RemoteGw"); ok {
			if err = d.Set("remote_gw", vv); err != nil {
				return fmt.Errorf("Error reading remote_gw: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_gw: %v", err)
		}
	}

	if err = d.Set("remote_gw_country", flattenVpnIpsecPhase1RemoteGwCountry(o["remote-gw-country"], d, "remote_gw_country")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-gw-country"], "VpnIpsecPhase1-RemoteGwCountry"); ok {
			if err = d.Set("remote_gw_country", vv); err != nil {
				return fmt.Errorf("Error reading remote_gw_country: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_gw_country: %v", err)
		}
	}

	if err = d.Set("remote_gw_end_ip", flattenVpnIpsecPhase1RemoteGwEndIp(o["remote-gw-end-ip"], d, "remote_gw_end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-gw-end-ip"], "VpnIpsecPhase1-RemoteGwEndIp"); ok {
			if err = d.Set("remote_gw_end_ip", vv); err != nil {
				return fmt.Errorf("Error reading remote_gw_end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_gw_end_ip: %v", err)
		}
	}

	if err = d.Set("remote_gw_match", flattenVpnIpsecPhase1RemoteGwMatch(o["remote-gw-match"], d, "remote_gw_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-gw-match"], "VpnIpsecPhase1-RemoteGwMatch"); ok {
			if err = d.Set("remote_gw_match", vv); err != nil {
				return fmt.Errorf("Error reading remote_gw_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_gw_match: %v", err)
		}
	}

	if err = d.Set("remote_gw_start_ip", flattenVpnIpsecPhase1RemoteGwStartIp(o["remote-gw-start-ip"], d, "remote_gw_start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-gw-start-ip"], "VpnIpsecPhase1-RemoteGwStartIp"); ok {
			if err = d.Set("remote_gw_start_ip", vv); err != nil {
				return fmt.Errorf("Error reading remote_gw_start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_gw_start_ip: %v", err)
		}
	}

	if err = d.Set("remote_gw_subnet", flattenVpnIpsecPhase1RemoteGwSubnet(o["remote-gw-subnet"], d, "remote_gw_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-gw-subnet"], "VpnIpsecPhase1-RemoteGwSubnet"); ok {
			if err = d.Set("remote_gw_subnet", vv); err != nil {
				return fmt.Errorf("Error reading remote_gw_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_gw_subnet: %v", err)
		}
	}

	if err = d.Set("remote_gw_ztna_tags", flattenVpnIpsecPhase1RemoteGwZtnaTags(o["remote-gw-ztna-tags"], d, "remote_gw_ztna_tags")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-gw-ztna-tags"], "VpnIpsecPhase1-RemoteGwZtnaTags"); ok {
			if err = d.Set("remote_gw_ztna_tags", vv); err != nil {
				return fmt.Errorf("Error reading remote_gw_ztna_tags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_gw_ztna_tags: %v", err)
		}
	}

	if err = d.Set("remote_gw6_country", flattenVpnIpsecPhase1RemoteGw6Country(o["remote-gw6-country"], d, "remote_gw6_country")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-gw6-country"], "VpnIpsecPhase1-RemoteGw6Country"); ok {
			if err = d.Set("remote_gw6_country", vv); err != nil {
				return fmt.Errorf("Error reading remote_gw6_country: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_gw6_country: %v", err)
		}
	}

	if err = d.Set("remote_gw6_end_ip", flattenVpnIpsecPhase1RemoteGw6EndIp(o["remote-gw6-end-ip"], d, "remote_gw6_end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-gw6-end-ip"], "VpnIpsecPhase1-RemoteGw6EndIp"); ok {
			if err = d.Set("remote_gw6_end_ip", vv); err != nil {
				return fmt.Errorf("Error reading remote_gw6_end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_gw6_end_ip: %v", err)
		}
	}

	if err = d.Set("remote_gw6_match", flattenVpnIpsecPhase1RemoteGw6Match(o["remote-gw6-match"], d, "remote_gw6_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-gw6-match"], "VpnIpsecPhase1-RemoteGw6Match"); ok {
			if err = d.Set("remote_gw6_match", vv); err != nil {
				return fmt.Errorf("Error reading remote_gw6_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_gw6_match: %v", err)
		}
	}

	if err = d.Set("remote_gw6_start_ip", flattenVpnIpsecPhase1RemoteGw6StartIp(o["remote-gw6-start-ip"], d, "remote_gw6_start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-gw6-start-ip"], "VpnIpsecPhase1-RemoteGw6StartIp"); ok {
			if err = d.Set("remote_gw6_start_ip", vv); err != nil {
				return fmt.Errorf("Error reading remote_gw6_start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_gw6_start_ip: %v", err)
		}
	}

	if err = d.Set("remote_gw6_subnet", flattenVpnIpsecPhase1RemoteGw6Subnet(o["remote-gw6-subnet"], d, "remote_gw6_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-gw6-subnet"], "VpnIpsecPhase1-RemoteGw6Subnet"); ok {
			if err = d.Set("remote_gw6_subnet", vv); err != nil {
				return fmt.Errorf("Error reading remote_gw6_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_gw6_subnet: %v", err)
		}
	}

	if err = d.Set("remotegw_ddns", flattenVpnIpsecPhase1RemotegwDdns(o["remotegw-ddns"], d, "remotegw_ddns")); err != nil {
		if vv, ok := fortiAPIPatch(o["remotegw-ddns"], "VpnIpsecPhase1-RemotegwDdns"); ok {
			if err = d.Set("remotegw_ddns", vv); err != nil {
				return fmt.Errorf("Error reading remotegw_ddns: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remotegw_ddns: %v", err)
		}
	}

	if err = d.Set("rsa_signature_format", flattenVpnIpsecPhase1RsaSignatureFormat(o["rsa-signature-format"], d, "rsa_signature_format")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsa-signature-format"], "VpnIpsecPhase1-RsaSignatureFormat"); ok {
			if err = d.Set("rsa_signature_format", vv); err != nil {
				return fmt.Errorf("Error reading rsa_signature_format: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsa_signature_format: %v", err)
		}
	}

	if err = d.Set("rsa_signature_hash_override", flattenVpnIpsecPhase1RsaSignatureHashOverride(o["rsa-signature-hash-override"], d, "rsa_signature_hash_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsa-signature-hash-override"], "VpnIpsecPhase1-RsaSignatureHashOverride"); ok {
			if err = d.Set("rsa_signature_hash_override", vv); err != nil {
				return fmt.Errorf("Error reading rsa_signature_hash_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsa_signature_hash_override: %v", err)
		}
	}

	if err = d.Set("save_password", flattenVpnIpsecPhase1SavePassword(o["save-password"], d, "save_password")); err != nil {
		if vv, ok := fortiAPIPatch(o["save-password"], "VpnIpsecPhase1-SavePassword"); ok {
			if err = d.Set("save_password", vv); err != nil {
				return fmt.Errorf("Error reading save_password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading save_password: %v", err)
		}
	}

	if err = d.Set("send_cert_chain", flattenVpnIpsecPhase1SendCertChain(o["send-cert-chain"], d, "send_cert_chain")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-cert-chain"], "VpnIpsecPhase1-SendCertChain"); ok {
			if err = d.Set("send_cert_chain", vv); err != nil {
				return fmt.Errorf("Error reading send_cert_chain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_cert_chain: %v", err)
		}
	}

	if err = d.Set("shared_idle_timeout", flattenVpnIpsecPhase1SharedIdleTimeout(o["shared-idle-timeout"], d, "shared_idle_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["shared-idle-timeout"], "VpnIpsecPhase1-SharedIdleTimeout"); ok {
			if err = d.Set("shared_idle_timeout", vv); err != nil {
				return fmt.Errorf("Error reading shared_idle_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading shared_idle_timeout: %v", err)
		}
	}

	if err = d.Set("signature_hash_alg", flattenVpnIpsecPhase1SignatureHashAlg(o["signature-hash-alg"], d, "signature_hash_alg")); err != nil {
		if vv, ok := fortiAPIPatch(o["signature-hash-alg"], "VpnIpsecPhase1-SignatureHashAlg"); ok {
			if err = d.Set("signature_hash_alg", vv); err != nil {
				return fmt.Errorf("Error reading signature_hash_alg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading signature_hash_alg: %v", err)
		}
	}

	if err = d.Set("split_include_service", flattenVpnIpsecPhase1SplitIncludeService(o["split-include-service"], d, "split_include_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["split-include-service"], "VpnIpsecPhase1-SplitIncludeService"); ok {
			if err = d.Set("split_include_service", vv); err != nil {
				return fmt.Errorf("Error reading split_include_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading split_include_service: %v", err)
		}
	}

	if err = d.Set("suite_b", flattenVpnIpsecPhase1SuiteB(o["suite-b"], d, "suite_b")); err != nil {
		if vv, ok := fortiAPIPatch(o["suite-b"], "VpnIpsecPhase1-SuiteB"); ok {
			if err = d.Set("suite_b", vv); err != nil {
				return fmt.Errorf("Error reading suite_b: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading suite_b: %v", err)
		}
	}

	if err = d.Set("transit_gateway", flattenVpnIpsecPhase1TransitGateway(o["transit-gateway"], d, "transit_gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["transit-gateway"], "VpnIpsecPhase1-TransitGateway"); ok {
			if err = d.Set("transit_gateway", vv); err != nil {
				return fmt.Errorf("Error reading transit_gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading transit_gateway: %v", err)
		}
	}

	if err = d.Set("transport", flattenVpnIpsecPhase1Transport(o["transport"], d, "transport")); err != nil {
		if vv, ok := fortiAPIPatch(o["transport"], "VpnIpsecPhase1-Transport"); ok {
			if err = d.Set("transport", vv); err != nil {
				return fmt.Errorf("Error reading transport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading transport: %v", err)
		}
	}

	if err = d.Set("type", flattenVpnIpsecPhase1Type(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "VpnIpsecPhase1-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("unity_support", flattenVpnIpsecPhase1UnitySupport(o["unity-support"], d, "unity_support")); err != nil {
		if vv, ok := fortiAPIPatch(o["unity-support"], "VpnIpsecPhase1-UnitySupport"); ok {
			if err = d.Set("unity_support", vv); err != nil {
				return fmt.Errorf("Error reading unity_support: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unity_support: %v", err)
		}
	}

	if err = d.Set("usrgrp", flattenVpnIpsecPhase1Usrgrp(o["usrgrp"], d, "usrgrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["usrgrp"], "VpnIpsecPhase1-Usrgrp"); ok {
			if err = d.Set("usrgrp", vv); err != nil {
				return fmt.Errorf("Error reading usrgrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading usrgrp: %v", err)
		}
	}

	if err = d.Set("wizard_type", flattenVpnIpsecPhase1WizardType(o["wizard-type"], d, "wizard_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["wizard-type"], "VpnIpsecPhase1-WizardType"); ok {
			if err = d.Set("wizard_type", vv); err != nil {
				return fmt.Errorf("Error reading wizard_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wizard_type: %v", err)
		}
	}

	if err = d.Set("xauthtype", flattenVpnIpsecPhase1Xauthtype(o["xauthtype"], d, "xauthtype")); err != nil {
		if vv, ok := fortiAPIPatch(o["xauthtype"], "VpnIpsecPhase1-Xauthtype"); ok {
			if err = d.Set("xauthtype", vv); err != nil {
				return fmt.Errorf("Error reading xauthtype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading xauthtype: %v", err)
		}
	}

	return nil
}

func flattenVpnIpsecPhase1FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnIpsecPhase1AcctVerify(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1AddGwRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1AddRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Addke1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1Addke2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1Addke3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1Addke4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1Addke5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1Addke6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1Addke7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1AssignIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1AssignIpFrom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Authmethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1AuthmethodRemote(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Authpasswd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1Authusr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Authusrgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1AutoNegotiate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1AutoTransportThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1AzureAdAutoconnect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1BackupGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1Banner(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1CertIdValidation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1CertPeerUsernameStrip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1CertPeerUsernameValidation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1CertTrustStore(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Certificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1ChildlessIke(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1ClientAutoNegotiate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1ClientKeepAlive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1ClientResume(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1ClientResumeInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Comments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1DevId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1DevIdNotification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1DhcpRaGiaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Dhcp6RaLinkaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Dhgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1DigitalSignatureAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Distance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1DnsMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Domain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Dpd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1DpdRetrycount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1DpdRetryinterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1Eap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1EapCertAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1EapExcludePeergrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1EapIdentity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1EmsSnCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1EnforceUniqueId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Esn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1ExchangeFgtDeviceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1FallbackTcpThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1FecBase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1FecCodec(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1FecEgress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1FecHealthCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1FecIngress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1FecMappingProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1FecReceiveTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1FecRedundant(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1FecSendTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1FgspSync(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1ForticlientEnforcement(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1FortinetEsp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Fragmentation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1FragmentationMtu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1GroupAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1GroupAuthenticationSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1HaSyncEspSeqno(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1IdleTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1IdleTimeoutinterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1IkeVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InboundDscpCopy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1IncludeLocalLan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Interface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1InternalDomainList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1IpDelayInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4DnsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4DnsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4DnsServer3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4EndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4ExcludeRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["end-ip"], _ = expandVpnIpsecPhase1Ipv4ExcludeRangeEndIp(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandVpnIpsecPhase1Ipv4ExcludeRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip"], _ = expandVpnIpsecPhase1Ipv4ExcludeRangeStartIp(d, i["start_ip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandVpnIpsecPhase1Ipv4ExcludeRangeEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4ExcludeRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4ExcludeRangeStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1Ipv4Netmask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4SplitExclude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1Ipv4SplitInclude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1Ipv4StartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4WinsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4WinsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6AutoLinklocal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6DnsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6DnsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6DnsServer3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6EndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6ExcludeRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["end-ip"], _ = expandVpnIpsecPhase1Ipv6ExcludeRangeEndIp(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandVpnIpsecPhase1Ipv6ExcludeRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip"], _ = expandVpnIpsecPhase1Ipv6ExcludeRangeStartIp(d, i["start_ip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandVpnIpsecPhase1Ipv6ExcludeRangeEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6ExcludeRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6ExcludeRangeStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1Ipv6Prefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6SplitExclude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1Ipv6SplitInclude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1Ipv6StartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Keepalive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Keylife(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Kms(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1LinkCost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1LocalGw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Localid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1LocalidType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1LoopbackAsymroute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1MeshSelectorType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1ModeCfg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1ModeCfgAllowClientSelector(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Nattraversal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1NegotiateTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1NetworkId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1NetworkOverlay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1NpuOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Peer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1Peergrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1Peerid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Peertype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ppk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1PpkIdentity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1PpkSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1Priority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Proposal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1Psksecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1PsksecretRemote(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1Qkd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1QkdProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1Reauth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Rekey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RemoteGw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RemoteGwCountry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RemoteGwEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RemoteGwMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RemoteGwStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RemoteGwSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandVpnIpsecPhase1RemoteGwZtnaTags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1RemoteGw6Country(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RemoteGw6EndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RemoteGw6Match(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RemoteGw6StartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RemoteGw6Subnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RemotegwDdns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RsaSignatureFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RsaSignatureHashOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1SavePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1SendCertChain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1SharedIdleTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1SignatureHashAlg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1SplitIncludeService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1SuiteB(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1TransitGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Transport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Type(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1UnitySupport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Usrgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase1WizardType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Xauthtype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnIpsecPhase1(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("acct_verify"); ok || d.HasChange("acct_verify") {
		t, err := expandVpnIpsecPhase1AcctVerify(d, v, "acct_verify")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acct-verify"] = t
		}
	}

	if v, ok := d.GetOk("add_gw_route"); ok || d.HasChange("add_gw_route") {
		t, err := expandVpnIpsecPhase1AddGwRoute(d, v, "add_gw_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-gw-route"] = t
		}
	}

	if v, ok := d.GetOk("add_route"); ok || d.HasChange("add_route") {
		t, err := expandVpnIpsecPhase1AddRoute(d, v, "add_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-route"] = t
		}
	}

	if v, ok := d.GetOk("addke1"); ok || d.HasChange("addke1") {
		t, err := expandVpnIpsecPhase1Addke1(d, v, "addke1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke1"] = t
		}
	}

	if v, ok := d.GetOk("addke2"); ok || d.HasChange("addke2") {
		t, err := expandVpnIpsecPhase1Addke2(d, v, "addke2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke2"] = t
		}
	}

	if v, ok := d.GetOk("addke3"); ok || d.HasChange("addke3") {
		t, err := expandVpnIpsecPhase1Addke3(d, v, "addke3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke3"] = t
		}
	}

	if v, ok := d.GetOk("addke4"); ok || d.HasChange("addke4") {
		t, err := expandVpnIpsecPhase1Addke4(d, v, "addke4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke4"] = t
		}
	}

	if v, ok := d.GetOk("addke5"); ok || d.HasChange("addke5") {
		t, err := expandVpnIpsecPhase1Addke5(d, v, "addke5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke5"] = t
		}
	}

	if v, ok := d.GetOk("addke6"); ok || d.HasChange("addke6") {
		t, err := expandVpnIpsecPhase1Addke6(d, v, "addke6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke6"] = t
		}
	}

	if v, ok := d.GetOk("addke7"); ok || d.HasChange("addke7") {
		t, err := expandVpnIpsecPhase1Addke7(d, v, "addke7")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke7"] = t
		}
	}

	if v, ok := d.GetOk("assign_ip"); ok || d.HasChange("assign_ip") {
		t, err := expandVpnIpsecPhase1AssignIp(d, v, "assign_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["assign-ip"] = t
		}
	}

	if v, ok := d.GetOk("assign_ip_from"); ok || d.HasChange("assign_ip_from") {
		t, err := expandVpnIpsecPhase1AssignIpFrom(d, v, "assign_ip_from")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["assign-ip-from"] = t
		}
	}

	if v, ok := d.GetOk("authmethod"); ok || d.HasChange("authmethod") {
		t, err := expandVpnIpsecPhase1Authmethod(d, v, "authmethod")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authmethod"] = t
		}
	}

	if v, ok := d.GetOk("authmethod_remote"); ok || d.HasChange("authmethod_remote") {
		t, err := expandVpnIpsecPhase1AuthmethodRemote(d, v, "authmethod_remote")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authmethod-remote"] = t
		}
	}

	if v, ok := d.GetOk("authpasswd"); ok || d.HasChange("authpasswd") {
		t, err := expandVpnIpsecPhase1Authpasswd(d, v, "authpasswd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authpasswd"] = t
		}
	}

	if v, ok := d.GetOk("authusr"); ok || d.HasChange("authusr") {
		t, err := expandVpnIpsecPhase1Authusr(d, v, "authusr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authusr"] = t
		}
	}

	if v, ok := d.GetOk("authusrgrp"); ok || d.HasChange("authusrgrp") {
		t, err := expandVpnIpsecPhase1Authusrgrp(d, v, "authusrgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authusrgrp"] = t
		}
	}

	if v, ok := d.GetOk("auto_negotiate"); ok || d.HasChange("auto_negotiate") {
		t, err := expandVpnIpsecPhase1AutoNegotiate(d, v, "auto_negotiate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-negotiate"] = t
		}
	}

	if v, ok := d.GetOk("auto_transport_threshold"); ok || d.HasChange("auto_transport_threshold") {
		t, err := expandVpnIpsecPhase1AutoTransportThreshold(d, v, "auto_transport_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-transport-threshold"] = t
		}
	}

	if v, ok := d.GetOk("azure_ad_autoconnect"); ok || d.HasChange("azure_ad_autoconnect") {
		t, err := expandVpnIpsecPhase1AzureAdAutoconnect(d, v, "azure_ad_autoconnect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["azure-ad-autoconnect"] = t
		}
	}

	if v, ok := d.GetOk("backup_gateway"); ok || d.HasChange("backup_gateway") {
		t, err := expandVpnIpsecPhase1BackupGateway(d, v, "backup_gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backup-gateway"] = t
		}
	}

	if v, ok := d.GetOk("banner"); ok || d.HasChange("banner") {
		t, err := expandVpnIpsecPhase1Banner(d, v, "banner")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["banner"] = t
		}
	}

	if v, ok := d.GetOk("cert_id_validation"); ok || d.HasChange("cert_id_validation") {
		t, err := expandVpnIpsecPhase1CertIdValidation(d, v, "cert_id_validation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-id-validation"] = t
		}
	}

	if v, ok := d.GetOk("cert_peer_username_strip"); ok || d.HasChange("cert_peer_username_strip") {
		t, err := expandVpnIpsecPhase1CertPeerUsernameStrip(d, v, "cert_peer_username_strip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-peer-username-strip"] = t
		}
	}

	if v, ok := d.GetOk("cert_peer_username_validation"); ok || d.HasChange("cert_peer_username_validation") {
		t, err := expandVpnIpsecPhase1CertPeerUsernameValidation(d, v, "cert_peer_username_validation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-peer-username-validation"] = t
		}
	}

	if v, ok := d.GetOk("cert_trust_store"); ok || d.HasChange("cert_trust_store") {
		t, err := expandVpnIpsecPhase1CertTrustStore(d, v, "cert_trust_store")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-trust-store"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok || d.HasChange("certificate") {
		t, err := expandVpnIpsecPhase1Certificate(d, v, "certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("childless_ike"); ok || d.HasChange("childless_ike") {
		t, err := expandVpnIpsecPhase1ChildlessIke(d, v, "childless_ike")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["childless-ike"] = t
		}
	}

	if v, ok := d.GetOk("client_auto_negotiate"); ok || d.HasChange("client_auto_negotiate") {
		t, err := expandVpnIpsecPhase1ClientAutoNegotiate(d, v, "client_auto_negotiate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-auto-negotiate"] = t
		}
	}

	if v, ok := d.GetOk("client_keep_alive"); ok || d.HasChange("client_keep_alive") {
		t, err := expandVpnIpsecPhase1ClientKeepAlive(d, v, "client_keep_alive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-keep-alive"] = t
		}
	}

	if v, ok := d.GetOk("client_resume"); ok || d.HasChange("client_resume") {
		t, err := expandVpnIpsecPhase1ClientResume(d, v, "client_resume")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-resume"] = t
		}
	}

	if v, ok := d.GetOk("client_resume_interval"); ok || d.HasChange("client_resume_interval") {
		t, err := expandVpnIpsecPhase1ClientResumeInterval(d, v, "client_resume_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-resume-interval"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandVpnIpsecPhase1Comments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("dev_id"); ok || d.HasChange("dev_id") {
		t, err := expandVpnIpsecPhase1DevId(d, v, "dev_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dev-id"] = t
		}
	}

	if v, ok := d.GetOk("dev_id_notification"); ok || d.HasChange("dev_id_notification") {
		t, err := expandVpnIpsecPhase1DevIdNotification(d, v, "dev_id_notification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dev-id-notification"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_ra_giaddr"); ok || d.HasChange("dhcp_ra_giaddr") {
		t, err := expandVpnIpsecPhase1DhcpRaGiaddr(d, v, "dhcp_ra_giaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-ra-giaddr"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_ra_linkaddr"); ok || d.HasChange("dhcp6_ra_linkaddr") {
		t, err := expandVpnIpsecPhase1Dhcp6RaLinkaddr(d, v, "dhcp6_ra_linkaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-ra-linkaddr"] = t
		}
	}

	if v, ok := d.GetOk("dhgrp"); ok || d.HasChange("dhgrp") {
		t, err := expandVpnIpsecPhase1Dhgrp(d, v, "dhgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhgrp"] = t
		}
	}

	if v, ok := d.GetOk("digital_signature_auth"); ok || d.HasChange("digital_signature_auth") {
		t, err := expandVpnIpsecPhase1DigitalSignatureAuth(d, v, "digital_signature_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["digital-signature-auth"] = t
		}
	}

	if v, ok := d.GetOk("distance"); ok || d.HasChange("distance") {
		t, err := expandVpnIpsecPhase1Distance(d, v, "distance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance"] = t
		}
	}

	if v, ok := d.GetOk("dns_mode"); ok || d.HasChange("dns_mode") {
		t, err := expandVpnIpsecPhase1DnsMode(d, v, "dns_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-mode"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok || d.HasChange("domain") {
		t, err := expandVpnIpsecPhase1Domain(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("dpd"); ok || d.HasChange("dpd") {
		t, err := expandVpnIpsecPhase1Dpd(d, v, "dpd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dpd"] = t
		}
	}

	if v, ok := d.GetOk("dpd_retrycount"); ok || d.HasChange("dpd_retrycount") {
		t, err := expandVpnIpsecPhase1DpdRetrycount(d, v, "dpd_retrycount")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dpd-retrycount"] = t
		}
	}

	if v, ok := d.GetOk("dpd_retryinterval"); ok || d.HasChange("dpd_retryinterval") {
		t, err := expandVpnIpsecPhase1DpdRetryinterval(d, v, "dpd_retryinterval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dpd-retryinterval"] = t
		}
	}

	if v, ok := d.GetOk("eap"); ok || d.HasChange("eap") {
		t, err := expandVpnIpsecPhase1Eap(d, v, "eap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap"] = t
		}
	}

	if v, ok := d.GetOk("eap_cert_auth"); ok || d.HasChange("eap_cert_auth") {
		t, err := expandVpnIpsecPhase1EapCertAuth(d, v, "eap_cert_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-cert-auth"] = t
		}
	}

	if v, ok := d.GetOk("eap_exclude_peergrp"); ok || d.HasChange("eap_exclude_peergrp") {
		t, err := expandVpnIpsecPhase1EapExcludePeergrp(d, v, "eap_exclude_peergrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-exclude-peergrp"] = t
		}
	}

	if v, ok := d.GetOk("eap_identity"); ok || d.HasChange("eap_identity") {
		t, err := expandVpnIpsecPhase1EapIdentity(d, v, "eap_identity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-identity"] = t
		}
	}

	if v, ok := d.GetOk("ems_sn_check"); ok || d.HasChange("ems_sn_check") {
		t, err := expandVpnIpsecPhase1EmsSnCheck(d, v, "ems_sn_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ems-sn-check"] = t
		}
	}

	if v, ok := d.GetOk("enforce_unique_id"); ok || d.HasChange("enforce_unique_id") {
		t, err := expandVpnIpsecPhase1EnforceUniqueId(d, v, "enforce_unique_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enforce-unique-id"] = t
		}
	}

	if v, ok := d.GetOk("esn"); ok || d.HasChange("esn") {
		t, err := expandVpnIpsecPhase1Esn(d, v, "esn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["esn"] = t
		}
	}

	if v, ok := d.GetOk("exchange_fgt_device_id"); ok || d.HasChange("exchange_fgt_device_id") {
		t, err := expandVpnIpsecPhase1ExchangeFgtDeviceId(d, v, "exchange_fgt_device_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exchange-fgt-device-id"] = t
		}
	}

	if v, ok := d.GetOk("fallback_tcp_threshold"); ok || d.HasChange("fallback_tcp_threshold") {
		t, err := expandVpnIpsecPhase1FallbackTcpThreshold(d, v, "fallback_tcp_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fallback-tcp-threshold"] = t
		}
	}

	if v, ok := d.GetOk("fec_base"); ok || d.HasChange("fec_base") {
		t, err := expandVpnIpsecPhase1FecBase(d, v, "fec_base")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-base"] = t
		}
	}

	if v, ok := d.GetOk("fec_codec"); ok || d.HasChange("fec_codec") {
		t, err := expandVpnIpsecPhase1FecCodec(d, v, "fec_codec")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-codec"] = t
		}
	}

	if v, ok := d.GetOk("fec_egress"); ok || d.HasChange("fec_egress") {
		t, err := expandVpnIpsecPhase1FecEgress(d, v, "fec_egress")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-egress"] = t
		}
	}

	if v, ok := d.GetOk("fec_health_check"); ok || d.HasChange("fec_health_check") {
		t, err := expandVpnIpsecPhase1FecHealthCheck(d, v, "fec_health_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-health-check"] = t
		}
	}

	if v, ok := d.GetOk("fec_ingress"); ok || d.HasChange("fec_ingress") {
		t, err := expandVpnIpsecPhase1FecIngress(d, v, "fec_ingress")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-ingress"] = t
		}
	}

	if v, ok := d.GetOk("fec_mapping_profile"); ok || d.HasChange("fec_mapping_profile") {
		t, err := expandVpnIpsecPhase1FecMappingProfile(d, v, "fec_mapping_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-mapping-profile"] = t
		}
	}

	if v, ok := d.GetOk("fec_receive_timeout"); ok || d.HasChange("fec_receive_timeout") {
		t, err := expandVpnIpsecPhase1FecReceiveTimeout(d, v, "fec_receive_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-receive-timeout"] = t
		}
	}

	if v, ok := d.GetOk("fec_redundant"); ok || d.HasChange("fec_redundant") {
		t, err := expandVpnIpsecPhase1FecRedundant(d, v, "fec_redundant")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-redundant"] = t
		}
	}

	if v, ok := d.GetOk("fec_send_timeout"); ok || d.HasChange("fec_send_timeout") {
		t, err := expandVpnIpsecPhase1FecSendTimeout(d, v, "fec_send_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-send-timeout"] = t
		}
	}

	if v, ok := d.GetOk("fgsp_sync"); ok || d.HasChange("fgsp_sync") {
		t, err := expandVpnIpsecPhase1FgspSync(d, v, "fgsp_sync")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgsp-sync"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_enforcement"); ok || d.HasChange("forticlient_enforcement") {
		t, err := expandVpnIpsecPhase1ForticlientEnforcement(d, v, "forticlient_enforcement")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-enforcement"] = t
		}
	}

	if v, ok := d.GetOk("fortinet_esp"); ok || d.HasChange("fortinet_esp") {
		t, err := expandVpnIpsecPhase1FortinetEsp(d, v, "fortinet_esp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortinet-esp"] = t
		}
	}

	if v, ok := d.GetOk("fragmentation"); ok || d.HasChange("fragmentation") {
		t, err := expandVpnIpsecPhase1Fragmentation(d, v, "fragmentation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fragmentation"] = t
		}
	}

	if v, ok := d.GetOk("fragmentation_mtu"); ok || d.HasChange("fragmentation_mtu") {
		t, err := expandVpnIpsecPhase1FragmentationMtu(d, v, "fragmentation_mtu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fragmentation-mtu"] = t
		}
	}

	if v, ok := d.GetOk("group_authentication"); ok || d.HasChange("group_authentication") {
		t, err := expandVpnIpsecPhase1GroupAuthentication(d, v, "group_authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-authentication"] = t
		}
	}

	if v, ok := d.GetOk("group_authentication_secret"); ok || d.HasChange("group_authentication_secret") {
		t, err := expandVpnIpsecPhase1GroupAuthenticationSecret(d, v, "group_authentication_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-authentication-secret"] = t
		}
	}

	if v, ok := d.GetOk("ha_sync_esp_seqno"); ok || d.HasChange("ha_sync_esp_seqno") {
		t, err := expandVpnIpsecPhase1HaSyncEspSeqno(d, v, "ha_sync_esp_seqno")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-sync-esp-seqno"] = t
		}
	}

	if v, ok := d.GetOk("idle_timeout"); ok || d.HasChange("idle_timeout") {
		t, err := expandVpnIpsecPhase1IdleTimeout(d, v, "idle_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idle-timeout"] = t
		}
	}

	if v, ok := d.GetOk("idle_timeoutinterval"); ok || d.HasChange("idle_timeoutinterval") {
		t, err := expandVpnIpsecPhase1IdleTimeoutinterval(d, v, "idle_timeoutinterval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idle-timeoutinterval"] = t
		}
	}

	if v, ok := d.GetOk("ike_version"); ok || d.HasChange("ike_version") {
		t, err := expandVpnIpsecPhase1IkeVersion(d, v, "ike_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-version"] = t
		}
	}

	if v, ok := d.GetOk("inbound_dscp_copy"); ok || d.HasChange("inbound_dscp_copy") {
		t, err := expandVpnIpsecPhase1InboundDscpCopy(d, v, "inbound_dscp_copy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inbound-dscp-copy"] = t
		}
	}

	if v, ok := d.GetOk("include_local_lan"); ok || d.HasChange("include_local_lan") {
		t, err := expandVpnIpsecPhase1IncludeLocalLan(d, v, "include_local_lan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["include-local-lan"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandVpnIpsecPhase1Interface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("internal_domain_list"); ok || d.HasChange("internal_domain_list") {
		t, err := expandVpnIpsecPhase1InternalDomainList(d, v, "internal_domain_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internal-domain-list"] = t
		}
	}

	if v, ok := d.GetOk("ip_delay_interval"); ok || d.HasChange("ip_delay_interval") {
		t, err := expandVpnIpsecPhase1IpDelayInterval(d, v, "ip_delay_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-delay-interval"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_dns_server1"); ok || d.HasChange("ipv4_dns_server1") {
		t, err := expandVpnIpsecPhase1Ipv4DnsServer1(d, v, "ipv4_dns_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_dns_server2"); ok || d.HasChange("ipv4_dns_server2") {
		t, err := expandVpnIpsecPhase1Ipv4DnsServer2(d, v, "ipv4_dns_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_dns_server3"); ok || d.HasChange("ipv4_dns_server3") {
		t, err := expandVpnIpsecPhase1Ipv4DnsServer3(d, v, "ipv4_dns_server3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-dns-server3"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_end_ip"); ok || d.HasChange("ipv4_end_ip") {
		t, err := expandVpnIpsecPhase1Ipv4EndIp(d, v, "ipv4_end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-end-ip"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_exclude_range"); ok || d.HasChange("ipv4_exclude_range") {
		t, err := expandVpnIpsecPhase1Ipv4ExcludeRange(d, v, "ipv4_exclude_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-exclude-range"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_name"); ok || d.HasChange("ipv4_name") {
		t, err := expandVpnIpsecPhase1Ipv4Name(d, v, "ipv4_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-name"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_netmask"); ok || d.HasChange("ipv4_netmask") {
		t, err := expandVpnIpsecPhase1Ipv4Netmask(d, v, "ipv4_netmask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-netmask"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_split_exclude"); ok || d.HasChange("ipv4_split_exclude") {
		t, err := expandVpnIpsecPhase1Ipv4SplitExclude(d, v, "ipv4_split_exclude")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-split-exclude"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_split_include"); ok || d.HasChange("ipv4_split_include") {
		t, err := expandVpnIpsecPhase1Ipv4SplitInclude(d, v, "ipv4_split_include")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-split-include"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_start_ip"); ok || d.HasChange("ipv4_start_ip") {
		t, err := expandVpnIpsecPhase1Ipv4StartIp(d, v, "ipv4_start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-start-ip"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_wins_server1"); ok || d.HasChange("ipv4_wins_server1") {
		t, err := expandVpnIpsecPhase1Ipv4WinsServer1(d, v, "ipv4_wins_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-wins-server1"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_wins_server2"); ok || d.HasChange("ipv4_wins_server2") {
		t, err := expandVpnIpsecPhase1Ipv4WinsServer2(d, v, "ipv4_wins_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-wins-server2"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_auto_linklocal"); ok || d.HasChange("ipv6_auto_linklocal") {
		t, err := expandVpnIpsecPhase1Ipv6AutoLinklocal(d, v, "ipv6_auto_linklocal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-auto-linklocal"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_dns_server1"); ok || d.HasChange("ipv6_dns_server1") {
		t, err := expandVpnIpsecPhase1Ipv6DnsServer1(d, v, "ipv6_dns_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_dns_server2"); ok || d.HasChange("ipv6_dns_server2") {
		t, err := expandVpnIpsecPhase1Ipv6DnsServer2(d, v, "ipv6_dns_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_dns_server3"); ok || d.HasChange("ipv6_dns_server3") {
		t, err := expandVpnIpsecPhase1Ipv6DnsServer3(d, v, "ipv6_dns_server3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-dns-server3"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_end_ip"); ok || d.HasChange("ipv6_end_ip") {
		t, err := expandVpnIpsecPhase1Ipv6EndIp(d, v, "ipv6_end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-end-ip"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_exclude_range"); ok || d.HasChange("ipv6_exclude_range") {
		t, err := expandVpnIpsecPhase1Ipv6ExcludeRange(d, v, "ipv6_exclude_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-exclude-range"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_name"); ok || d.HasChange("ipv6_name") {
		t, err := expandVpnIpsecPhase1Ipv6Name(d, v, "ipv6_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-name"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_prefix"); ok || d.HasChange("ipv6_prefix") {
		t, err := expandVpnIpsecPhase1Ipv6Prefix(d, v, "ipv6_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-prefix"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_split_exclude"); ok || d.HasChange("ipv6_split_exclude") {
		t, err := expandVpnIpsecPhase1Ipv6SplitExclude(d, v, "ipv6_split_exclude")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-split-exclude"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_split_include"); ok || d.HasChange("ipv6_split_include") {
		t, err := expandVpnIpsecPhase1Ipv6SplitInclude(d, v, "ipv6_split_include")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-split-include"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_start_ip"); ok || d.HasChange("ipv6_start_ip") {
		t, err := expandVpnIpsecPhase1Ipv6StartIp(d, v, "ipv6_start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-start-ip"] = t
		}
	}

	if v, ok := d.GetOk("keepalive"); ok || d.HasChange("keepalive") {
		t, err := expandVpnIpsecPhase1Keepalive(d, v, "keepalive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keepalive"] = t
		}
	}

	if v, ok := d.GetOk("keylife"); ok || d.HasChange("keylife") {
		t, err := expandVpnIpsecPhase1Keylife(d, v, "keylife")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keylife"] = t
		}
	}

	if v, ok := d.GetOk("kms"); ok || d.HasChange("kms") {
		t, err := expandVpnIpsecPhase1Kms(d, v, "kms")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["kms"] = t
		}
	}

	if v, ok := d.GetOk("link_cost"); ok || d.HasChange("link_cost") {
		t, err := expandVpnIpsecPhase1LinkCost(d, v, "link_cost")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-cost"] = t
		}
	}

	if v, ok := d.GetOk("local_gw"); ok || d.HasChange("local_gw") {
		t, err := expandVpnIpsecPhase1LocalGw(d, v, "local_gw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-gw"] = t
		}
	}

	if v, ok := d.GetOk("localid"); ok || d.HasChange("localid") {
		t, err := expandVpnIpsecPhase1Localid(d, v, "localid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["localid"] = t
		}
	}

	if v, ok := d.GetOk("localid_type"); ok || d.HasChange("localid_type") {
		t, err := expandVpnIpsecPhase1LocalidType(d, v, "localid_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["localid-type"] = t
		}
	}

	if v, ok := d.GetOk("loopback_asymroute"); ok || d.HasChange("loopback_asymroute") {
		t, err := expandVpnIpsecPhase1LoopbackAsymroute(d, v, "loopback_asymroute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["loopback-asymroute"] = t
		}
	}

	if v, ok := d.GetOk("mesh_selector_type"); ok || d.HasChange("mesh_selector_type") {
		t, err := expandVpnIpsecPhase1MeshSelectorType(d, v, "mesh_selector_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mesh-selector-type"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandVpnIpsecPhase1Mode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("mode_cfg"); ok || d.HasChange("mode_cfg") {
		t, err := expandVpnIpsecPhase1ModeCfg(d, v, "mode_cfg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode-cfg"] = t
		}
	}

	if v, ok := d.GetOk("mode_cfg_allow_client_selector"); ok || d.HasChange("mode_cfg_allow_client_selector") {
		t, err := expandVpnIpsecPhase1ModeCfgAllowClientSelector(d, v, "mode_cfg_allow_client_selector")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode-cfg-allow-client-selector"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandVpnIpsecPhase1Name(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nattraversal"); ok || d.HasChange("nattraversal") {
		t, err := expandVpnIpsecPhase1Nattraversal(d, v, "nattraversal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nattraversal"] = t
		}
	}

	if v, ok := d.GetOk("negotiate_timeout"); ok || d.HasChange("negotiate_timeout") {
		t, err := expandVpnIpsecPhase1NegotiateTimeout(d, v, "negotiate_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["negotiate-timeout"] = t
		}
	}

	if v, ok := d.GetOk("network_id"); ok || d.HasChange("network_id") {
		t, err := expandVpnIpsecPhase1NetworkId(d, v, "network_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-id"] = t
		}
	}

	if v, ok := d.GetOk("network_overlay"); ok || d.HasChange("network_overlay") {
		t, err := expandVpnIpsecPhase1NetworkOverlay(d, v, "network_overlay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-overlay"] = t
		}
	}

	if v, ok := d.GetOk("npu_offload"); ok || d.HasChange("npu_offload") {
		t, err := expandVpnIpsecPhase1NpuOffload(d, v, "npu_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["npu-offload"] = t
		}
	}

	if v, ok := d.GetOk("peer"); ok || d.HasChange("peer") {
		t, err := expandVpnIpsecPhase1Peer(d, v, "peer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer"] = t
		}
	}

	if v, ok := d.GetOk("peergrp"); ok || d.HasChange("peergrp") {
		t, err := expandVpnIpsecPhase1Peergrp(d, v, "peergrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peergrp"] = t
		}
	}

	if v, ok := d.GetOk("peerid"); ok || d.HasChange("peerid") {
		t, err := expandVpnIpsecPhase1Peerid(d, v, "peerid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peerid"] = t
		}
	}

	if v, ok := d.GetOk("peertype"); ok || d.HasChange("peertype") {
		t, err := expandVpnIpsecPhase1Peertype(d, v, "peertype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peertype"] = t
		}
	}

	if v, ok := d.GetOk("ppk"); ok || d.HasChange("ppk") {
		t, err := expandVpnIpsecPhase1Ppk(d, v, "ppk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppk"] = t
		}
	}

	if v, ok := d.GetOk("ppk_identity"); ok || d.HasChange("ppk_identity") {
		t, err := expandVpnIpsecPhase1PpkIdentity(d, v, "ppk_identity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppk-identity"] = t
		}
	}

	if v, ok := d.GetOk("ppk_secret"); ok || d.HasChange("ppk_secret") {
		t, err := expandVpnIpsecPhase1PpkSecret(d, v, "ppk_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppk-secret"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandVpnIpsecPhase1Priority(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("proposal"); ok || d.HasChange("proposal") {
		t, err := expandVpnIpsecPhase1Proposal(d, v, "proposal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proposal"] = t
		}
	}

	if v, ok := d.GetOk("psksecret"); ok || d.HasChange("psksecret") {
		t, err := expandVpnIpsecPhase1Psksecret(d, v, "psksecret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["psksecret"] = t
		}
	}

	if v, ok := d.GetOk("psksecret_remote"); ok || d.HasChange("psksecret_remote") {
		t, err := expandVpnIpsecPhase1PsksecretRemote(d, v, "psksecret_remote")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["psksecret-remote"] = t
		}
	}

	if v, ok := d.GetOk("qkd"); ok || d.HasChange("qkd") {
		t, err := expandVpnIpsecPhase1Qkd(d, v, "qkd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qkd"] = t
		}
	}

	if v, ok := d.GetOk("qkd_profile"); ok || d.HasChange("qkd_profile") {
		t, err := expandVpnIpsecPhase1QkdProfile(d, v, "qkd_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qkd-profile"] = t
		}
	}

	if v, ok := d.GetOk("reauth"); ok || d.HasChange("reauth") {
		t, err := expandVpnIpsecPhase1Reauth(d, v, "reauth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reauth"] = t
		}
	}

	if v, ok := d.GetOk("rekey"); ok || d.HasChange("rekey") {
		t, err := expandVpnIpsecPhase1Rekey(d, v, "rekey")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rekey"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw"); ok || d.HasChange("remote_gw") {
		t, err := expandVpnIpsecPhase1RemoteGw(d, v, "remote_gw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw_country"); ok || d.HasChange("remote_gw_country") {
		t, err := expandVpnIpsecPhase1RemoteGwCountry(d, v, "remote_gw_country")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw-country"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw_end_ip"); ok || d.HasChange("remote_gw_end_ip") {
		t, err := expandVpnIpsecPhase1RemoteGwEndIp(d, v, "remote_gw_end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw-end-ip"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw_match"); ok || d.HasChange("remote_gw_match") {
		t, err := expandVpnIpsecPhase1RemoteGwMatch(d, v, "remote_gw_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw-match"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw_start_ip"); ok || d.HasChange("remote_gw_start_ip") {
		t, err := expandVpnIpsecPhase1RemoteGwStartIp(d, v, "remote_gw_start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw-start-ip"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw_subnet"); ok || d.HasChange("remote_gw_subnet") {
		t, err := expandVpnIpsecPhase1RemoteGwSubnet(d, v, "remote_gw_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw-subnet"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw_ztna_tags"); ok || d.HasChange("remote_gw_ztna_tags") {
		t, err := expandVpnIpsecPhase1RemoteGwZtnaTags(d, v, "remote_gw_ztna_tags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw-ztna-tags"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw6_country"); ok || d.HasChange("remote_gw6_country") {
		t, err := expandVpnIpsecPhase1RemoteGw6Country(d, v, "remote_gw6_country")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw6-country"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw6_end_ip"); ok || d.HasChange("remote_gw6_end_ip") {
		t, err := expandVpnIpsecPhase1RemoteGw6EndIp(d, v, "remote_gw6_end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw6-end-ip"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw6_match"); ok || d.HasChange("remote_gw6_match") {
		t, err := expandVpnIpsecPhase1RemoteGw6Match(d, v, "remote_gw6_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw6-match"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw6_start_ip"); ok || d.HasChange("remote_gw6_start_ip") {
		t, err := expandVpnIpsecPhase1RemoteGw6StartIp(d, v, "remote_gw6_start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw6-start-ip"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw6_subnet"); ok || d.HasChange("remote_gw6_subnet") {
		t, err := expandVpnIpsecPhase1RemoteGw6Subnet(d, v, "remote_gw6_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw6-subnet"] = t
		}
	}

	if v, ok := d.GetOk("remotegw_ddns"); ok || d.HasChange("remotegw_ddns") {
		t, err := expandVpnIpsecPhase1RemotegwDdns(d, v, "remotegw_ddns")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remotegw-ddns"] = t
		}
	}

	if v, ok := d.GetOk("rsa_signature_format"); ok || d.HasChange("rsa_signature_format") {
		t, err := expandVpnIpsecPhase1RsaSignatureFormat(d, v, "rsa_signature_format")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsa-signature-format"] = t
		}
	}

	if v, ok := d.GetOk("rsa_signature_hash_override"); ok || d.HasChange("rsa_signature_hash_override") {
		t, err := expandVpnIpsecPhase1RsaSignatureHashOverride(d, v, "rsa_signature_hash_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsa-signature-hash-override"] = t
		}
	}

	if v, ok := d.GetOk("save_password"); ok || d.HasChange("save_password") {
		t, err := expandVpnIpsecPhase1SavePassword(d, v, "save_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["save-password"] = t
		}
	}

	if v, ok := d.GetOk("send_cert_chain"); ok || d.HasChange("send_cert_chain") {
		t, err := expandVpnIpsecPhase1SendCertChain(d, v, "send_cert_chain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-cert-chain"] = t
		}
	}

	if v, ok := d.GetOk("shared_idle_timeout"); ok || d.HasChange("shared_idle_timeout") {
		t, err := expandVpnIpsecPhase1SharedIdleTimeout(d, v, "shared_idle_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shared-idle-timeout"] = t
		}
	}

	if v, ok := d.GetOk("signature_hash_alg"); ok || d.HasChange("signature_hash_alg") {
		t, err := expandVpnIpsecPhase1SignatureHashAlg(d, v, "signature_hash_alg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signature-hash-alg"] = t
		}
	}

	if v, ok := d.GetOk("split_include_service"); ok || d.HasChange("split_include_service") {
		t, err := expandVpnIpsecPhase1SplitIncludeService(d, v, "split_include_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-include-service"] = t
		}
	}

	if v, ok := d.GetOk("suite_b"); ok || d.HasChange("suite_b") {
		t, err := expandVpnIpsecPhase1SuiteB(d, v, "suite_b")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["suite-b"] = t
		}
	}

	if v, ok := d.GetOk("transit_gateway"); ok || d.HasChange("transit_gateway") {
		t, err := expandVpnIpsecPhase1TransitGateway(d, v, "transit_gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transit-gateway"] = t
		}
	}

	if v, ok := d.GetOk("transport"); ok || d.HasChange("transport") {
		t, err := expandVpnIpsecPhase1Transport(d, v, "transport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transport"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandVpnIpsecPhase1Type(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("unity_support"); ok || d.HasChange("unity_support") {
		t, err := expandVpnIpsecPhase1UnitySupport(d, v, "unity_support")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unity-support"] = t
		}
	}

	if v, ok := d.GetOk("usrgrp"); ok || d.HasChange("usrgrp") {
		t, err := expandVpnIpsecPhase1Usrgrp(d, v, "usrgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["usrgrp"] = t
		}
	}

	if v, ok := d.GetOk("wizard_type"); ok || d.HasChange("wizard_type") {
		t, err := expandVpnIpsecPhase1WizardType(d, v, "wizard_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wizard-type"] = t
		}
	}

	if v, ok := d.GetOk("xauthtype"); ok || d.HasChange("xauthtype") {
		t, err := expandVpnIpsecPhase1Xauthtype(d, v, "xauthtype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["xauthtype"] = t
		}
	}

	return &obj, nil
}
