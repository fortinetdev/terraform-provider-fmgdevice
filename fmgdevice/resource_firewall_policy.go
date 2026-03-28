// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure IPv4/IPv6 policies.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallPolicyCreate,
		Read:   resourceFirewallPolicyRead,
		Update: resourceFirewallPolicyUpdate,
		Delete: resourceFirewallPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
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
			"_policy_block": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"anti_replay": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"app_monitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"application_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"auth_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"auth_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_redirect_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_asic_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"av_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"block_notification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"captive_portal_exempt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"capture_packet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cifs_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"casb_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"cgn_eif": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cgn_eim": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cgn_log_server_grp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cgn_resource_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cgn_session_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cgn_sw_eif_ctrl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"custom_log_fields": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"decrypted_traffic_mirror": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"delay_tcp_npu_session": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diameter_filter_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"diffserv_copy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffserv_forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffserv_reverse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffservcode_forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"diffservcode_rev": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"disclaimer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dlp_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dnsfilter_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dsri": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstaddr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dstaddr_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstaddr6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dstaddr6_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstintf": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dynamic_shaping": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eif_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"eif_learn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"email_collect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"emailfilter_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fec": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"file_filter_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"firewall_session_dirty": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fixedport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fsso_agent_for_ntlm": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fsso_groups": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"geoip_anycast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"geoip_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"global_label": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"groups": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"gtp_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"http_policy_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"icap_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"identity_based_route": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"inbound": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"inspection_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_custom": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_custom_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_fortiguard": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_name": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service_src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_src_custom": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_src_custom_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_src_fortiguard": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_src_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_src_name": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_src_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service6_custom": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service6_custom_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service6_fortiguard": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service6_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service6_name": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service6_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service6_src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service6_src_custom": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service6_src_custom_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service6_src_fortiguard": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service6_src_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service6_src_name": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service6_src_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_version_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ippool": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ips_sensor": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ips_voip_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"label": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_http_transaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logtraffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logtraffic_start": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"match_vip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"match_vip_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat46": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat64": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"natinbound": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"natip": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"natoutbound": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_service_dynamic": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"network_service_src_dynamic": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"np_acceleration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ntlm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ntlm_enabled_browsers": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ntlm_guest": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"outbound": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"passive_wan_health_measurement": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pcp_inbound": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pcp_outbound": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pcp_poolname": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"per_ip_shaper": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"permit_any_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"permit_stun_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pfcp_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"policy_behaviour_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policy_expiry": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policy_expiry_date": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policy_expiry_date_utc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"policy_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"policyid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"poolname": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"poolname6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"port_preserve": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port_random": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"profile_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"profile_protocol_options": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"profile_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"radius_ip_auth_bypass": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_mac_auth_bypass": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirect_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"replacemsg_override_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"reputation_direction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reputation_direction6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reputation_minimum": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reputation_minimum6": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rtp_addr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"rtp_nat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"saml_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"schedule": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"schedule_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scim": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"scim_groups": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"scim_users": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sctp_filter_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"send_deny_packet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"service_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session_ttl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sgt": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"sgt_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_vendor_mac": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"srcaddr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"srcaddr_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcaddr6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"srcaddr6_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcintf": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssh_filter_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssh_policy_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_ssh_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_mss_receiver": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tcp_mss_sender": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tcp_session_without_syn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_timeout_pid": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"telemetry_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"timeout_send_rst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tos_mask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tos_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_shaper": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"traffic_shaper_reverse": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"udp_timeout_pid": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"users": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"utm_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"videofilter_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"virtual_patch_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vlan_cos_fwd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vlan_cos_rev": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vlan_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"voip_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vpntunnel": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"waf_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"wanopt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wanopt_detection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wanopt_passive_opt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wanopt_peer": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"wanopt_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"wccp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webcache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"webcache_https": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webfilter_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"webproxy_forward_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"webproxy_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ztna_device_ownership": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ztna_ems_tag": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ztna_ems_tag_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ztna_ems_tag_secondary": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ztna_geo_tag": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ztna_policy_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ztna_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ztna_tags_match_logic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"access_proxy": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"app_category": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"app_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"application": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"detect_https_in_http_request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_ownership": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dlp_sensor": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dynamic_bypass": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"explicit_web_proxy": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"extended_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"force_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_tunnel_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"https_sub_category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ia_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"implicit_proxy_detection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"isolator_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"isolator_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"max_session_per_user": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pass_through": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"redirect_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"reverse_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_connector": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssh_policy_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"transparent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"url_category": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"url_risk": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ztna_proxy": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

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

	obj, err := getObjectFirewallPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallPolicy resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("policyid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallPolicy(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallPolicy(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallPolicy resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateFirewallPolicy(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallPolicy resource: %v", err)
		}

		if v != nil && v["policyid"] != nil {
			if vidn, ok := v["policyid"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceFirewallPolicyRead(d, m)
			} else {
				return fmt.Errorf("Error creating FirewallPolicy resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourceFirewallPolicyRead(d, m)
}

func resourceFirewallPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

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

	obj, err := getObjectFirewallPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallPolicy resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallPolicy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourceFirewallPolicyRead(d, m)
}

func resourceFirewallPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

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

	wsParams["adom"] = adomv

	err = c.DeleteFirewallPolicy(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallPolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallPolicy(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallPolicy resource from API: %v", err)
	}
	return nil
}

func flattenFirewallPolicyPolicyBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyAntiReplay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyAppMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyApplicationList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyAuthCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyAuthPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyAuthRedirectAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyAutoAsicOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyAvProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyBlockNotification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyCaptivePortalExempt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyCapturePacket(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyCifsProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyCasbProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyCgnEif(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyCgnEim(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyCgnLogServerGrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyCgnResourceQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyCgnSessionQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyCgnSwEifCtrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyCustomLogFields(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyDecryptedTrafficMirror(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyDelayTcpNpuSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyDiameterFilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyDiffservCopy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyDiffservForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyDiffservReverse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyDiffservcodeForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyDiffservcodeRev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyDisclaimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyDlpProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyDnsfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyDsri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyDstaddrNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyDstaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyDstaddr6Negate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyDstintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyDynamicShaping(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyEifCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyEifLearn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyEmailCollect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyEmailfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyFec(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyFileFilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyFirewallSessionDirty(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyFixedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyFssoAgentForNtlm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyFssoGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyGeoipAnycast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyGeoipMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyGlobalLabel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyGtpProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyHttpPolicyRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyIcapProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyIdentityBasedRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyInbound(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyInspectionMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyInternetService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyInternetServiceCustom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyInternetServiceCustomGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyInternetServiceFortiguard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyInternetServiceGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyInternetServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyInternetServiceNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyInternetServiceSrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyInternetServiceSrcCustom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyInternetServiceSrcCustomGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyInternetServiceSrcFortiguard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyInternetServiceSrcGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyInternetServiceSrcName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyInternetServiceSrcNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyInternetService6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyInternetService6Custom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyInternetService6CustomGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyInternetService6Fortiguard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyInternetService6Group(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyInternetService6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyInternetService6Negate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyInternetService6Src(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyInternetService6SrcCustom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyInternetService6SrcCustomGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyInternetService6SrcFortiguard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyInternetService6SrcGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyInternetService6SrcName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyInternetService6SrcNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyIpVersionType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyIppool(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyIpsSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyIpsVoipFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyLabel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyLogHttpTransaction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyLogtraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyLogtrafficStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyMatchVip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyMatchVipOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyNat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyNat46(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyNat64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyNatinbound(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyNatip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyNatoutbound(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyNetworkServiceDynamic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyNetworkServiceSrcDynamic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyNpAcceleration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyNtlm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyNtlmEnabledBrowsers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyNtlmGuest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyOutbound(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyPassiveWanHealthMeasurement(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyPcpInbound(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyPcpOutbound(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyPcpPoolname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyPerIpShaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyPermitAnyHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyPermitStunHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyPfcpProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyPolicyBehaviourType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyPolicyExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyPolicyExpiryDate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyPolicyExpiryDateUtc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyPolicyOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyPolicyid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyPoolname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyPoolname6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyPortPreserve(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyPortRandom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyProfileGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyProfileProtocolOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyProfileType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyRadiusIpAuthBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyRadiusMacAuthBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyRedirectUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyReplacemsgOverrideGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyReputationDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyReputationDirection6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyReputationMinimum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyReputationMinimum6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyRtpAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyRtpNat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicySamlServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicySchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyScheduleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyScim(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyScimGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyScimUsers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicySctpFilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicySendDenyPacket(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyServiceNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicySessionTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicySgt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenFirewallPolicySgtCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicySrcVendorMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicySrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicySrcaddrNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicySrcaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicySrcaddr6Negate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicySrcintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicySshFilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicySshPolicyRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicySslSshProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyTcpMssReceiver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyTcpMssSender(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyTcpSessionWithoutSyn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyTcpTimeoutPid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyTelemetryProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyTimeoutSendRst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyTos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyTosMask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyTosNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyTrafficShaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyTrafficShaperReverse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyUdpTimeoutPid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyUsers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyUtmStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyVideofilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyVirtualPatchProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyVlanCosFwd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyVlanCosRev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyVlanFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyVoipProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyVpntunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyWafProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyWanopt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyWanoptDetection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyWanoptPassiveOpt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyWanoptPeer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyWanoptProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyWccp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyWebcache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyWebcacheHttps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyWebfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyWebproxyForwardServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyWebproxyProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyZtnaDeviceOwnership(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyZtnaEmsTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyZtnaEmsTagNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyZtnaEmsTagSecondary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyZtnaGeoTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyZtnaPolicyRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyZtnaStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyZtnaTagsMatchLogic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyAccessProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyAppCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyAppGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyApplication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenFirewallPolicyDetectHttpsInHttpRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyDeviceOwnership(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyDlpSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyDynamicBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyExplicitWebProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyExtendedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyForceProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyHttpTunnelAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyHttpsSubCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyIaProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyImplicitProxyDetection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyIsolatorProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyIsolatorServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyMaxSessionPerUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyPassThrough(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyRedirectProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyReverseCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyServiceConnector(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicySshPolicyCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyTransparent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyUrlCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyUrlRisk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyZtnaProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectFirewallPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("_policy_block", flattenFirewallPolicyPolicyBlock(o["_policy_block"], d, "_policy_block")); err != nil {
		if vv, ok := fortiAPIPatch(o["_policy_block"], "FirewallPolicy-PolicyBlock"); ok {
			if err = d.Set("_policy_block", vv); err != nil {
				return fmt.Errorf("Error reading _policy_block: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _policy_block: %v", err)
		}
	}

	if err = d.Set("action", flattenFirewallPolicyAction(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "FirewallPolicy-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("anti_replay", flattenFirewallPolicyAntiReplay(o["anti-replay"], d, "anti_replay")); err != nil {
		if vv, ok := fortiAPIPatch(o["anti-replay"], "FirewallPolicy-AntiReplay"); ok {
			if err = d.Set("anti_replay", vv); err != nil {
				return fmt.Errorf("Error reading anti_replay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading anti_replay: %v", err)
		}
	}

	if err = d.Set("app_monitor", flattenFirewallPolicyAppMonitor(o["app-monitor"], d, "app_monitor")); err != nil {
		if vv, ok := fortiAPIPatch(o["app-monitor"], "FirewallPolicy-AppMonitor"); ok {
			if err = d.Set("app_monitor", vv); err != nil {
				return fmt.Errorf("Error reading app_monitor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading app_monitor: %v", err)
		}
	}

	if err = d.Set("application_list", flattenFirewallPolicyApplicationList(o["application-list"], d, "application_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["application-list"], "FirewallPolicy-ApplicationList"); ok {
			if err = d.Set("application_list", vv); err != nil {
				return fmt.Errorf("Error reading application_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application_list: %v", err)
		}
	}

	if err = d.Set("auth_cert", flattenFirewallPolicyAuthCert(o["auth-cert"], d, "auth_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-cert"], "FirewallPolicy-AuthCert"); ok {
			if err = d.Set("auth_cert", vv); err != nil {
				return fmt.Errorf("Error reading auth_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_cert: %v", err)
		}
	}

	if err = d.Set("auth_path", flattenFirewallPolicyAuthPath(o["auth-path"], d, "auth_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-path"], "FirewallPolicy-AuthPath"); ok {
			if err = d.Set("auth_path", vv); err != nil {
				return fmt.Errorf("Error reading auth_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_path: %v", err)
		}
	}

	if err = d.Set("auth_redirect_addr", flattenFirewallPolicyAuthRedirectAddr(o["auth-redirect-addr"], d, "auth_redirect_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-redirect-addr"], "FirewallPolicy-AuthRedirectAddr"); ok {
			if err = d.Set("auth_redirect_addr", vv); err != nil {
				return fmt.Errorf("Error reading auth_redirect_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_redirect_addr: %v", err)
		}
	}

	if err = d.Set("auto_asic_offload", flattenFirewallPolicyAutoAsicOffload(o["auto-asic-offload"], d, "auto_asic_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-asic-offload"], "FirewallPolicy-AutoAsicOffload"); ok {
			if err = d.Set("auto_asic_offload", vv); err != nil {
				return fmt.Errorf("Error reading auto_asic_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_asic_offload: %v", err)
		}
	}

	if err = d.Set("av_profile", flattenFirewallPolicyAvProfile(o["av-profile"], d, "av_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-profile"], "FirewallPolicy-AvProfile"); ok {
			if err = d.Set("av_profile", vv); err != nil {
				return fmt.Errorf("Error reading av_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_profile: %v", err)
		}
	}

	if err = d.Set("block_notification", flattenFirewallPolicyBlockNotification(o["block-notification"], d, "block_notification")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-notification"], "FirewallPolicy-BlockNotification"); ok {
			if err = d.Set("block_notification", vv); err != nil {
				return fmt.Errorf("Error reading block_notification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_notification: %v", err)
		}
	}

	if err = d.Set("captive_portal_exempt", flattenFirewallPolicyCaptivePortalExempt(o["captive-portal-exempt"], d, "captive_portal_exempt")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-exempt"], "FirewallPolicy-CaptivePortalExempt"); ok {
			if err = d.Set("captive_portal_exempt", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_exempt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_exempt: %v", err)
		}
	}

	if err = d.Set("capture_packet", flattenFirewallPolicyCapturePacket(o["capture-packet"], d, "capture_packet")); err != nil {
		if vv, ok := fortiAPIPatch(o["capture-packet"], "FirewallPolicy-CapturePacket"); ok {
			if err = d.Set("capture_packet", vv); err != nil {
				return fmt.Errorf("Error reading capture_packet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capture_packet: %v", err)
		}
	}

	if err = d.Set("cifs_profile", flattenFirewallPolicyCifsProfile(o["cifs-profile"], d, "cifs_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["cifs-profile"], "FirewallPolicy-CifsProfile"); ok {
			if err = d.Set("cifs_profile", vv); err != nil {
				return fmt.Errorf("Error reading cifs_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cifs_profile: %v", err)
		}
	}

	if err = d.Set("casb_profile", flattenFirewallPolicyCasbProfile(o["casb-profile"], d, "casb_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["casb-profile"], "FirewallPolicy-CasbProfile"); ok {
			if err = d.Set("casb_profile", vv); err != nil {
				return fmt.Errorf("Error reading casb_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading casb_profile: %v", err)
		}
	}

	if err = d.Set("cgn_eif", flattenFirewallPolicyCgnEif(o["cgn-eif"], d, "cgn_eif")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-eif"], "FirewallPolicy-CgnEif"); ok {
			if err = d.Set("cgn_eif", vv); err != nil {
				return fmt.Errorf("Error reading cgn_eif: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_eif: %v", err)
		}
	}

	if err = d.Set("cgn_eim", flattenFirewallPolicyCgnEim(o["cgn-eim"], d, "cgn_eim")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-eim"], "FirewallPolicy-CgnEim"); ok {
			if err = d.Set("cgn_eim", vv); err != nil {
				return fmt.Errorf("Error reading cgn_eim: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_eim: %v", err)
		}
	}

	if err = d.Set("cgn_log_server_grp", flattenFirewallPolicyCgnLogServerGrp(o["cgn-log-server-grp"], d, "cgn_log_server_grp")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-log-server-grp"], "FirewallPolicy-CgnLogServerGrp"); ok {
			if err = d.Set("cgn_log_server_grp", vv); err != nil {
				return fmt.Errorf("Error reading cgn_log_server_grp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_log_server_grp: %v", err)
		}
	}

	if err = d.Set("cgn_resource_quota", flattenFirewallPolicyCgnResourceQuota(o["cgn-resource-quota"], d, "cgn_resource_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-resource-quota"], "FirewallPolicy-CgnResourceQuota"); ok {
			if err = d.Set("cgn_resource_quota", vv); err != nil {
				return fmt.Errorf("Error reading cgn_resource_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_resource_quota: %v", err)
		}
	}

	if err = d.Set("cgn_session_quota", flattenFirewallPolicyCgnSessionQuota(o["cgn-session-quota"], d, "cgn_session_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-session-quota"], "FirewallPolicy-CgnSessionQuota"); ok {
			if err = d.Set("cgn_session_quota", vv); err != nil {
				return fmt.Errorf("Error reading cgn_session_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_session_quota: %v", err)
		}
	}

	if err = d.Set("cgn_sw_eif_ctrl", flattenFirewallPolicyCgnSwEifCtrl(o["cgn-sw-eif-ctrl"], d, "cgn_sw_eif_ctrl")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-sw-eif-ctrl"], "FirewallPolicy-CgnSwEifCtrl"); ok {
			if err = d.Set("cgn_sw_eif_ctrl", vv); err != nil {
				return fmt.Errorf("Error reading cgn_sw_eif_ctrl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_sw_eif_ctrl: %v", err)
		}
	}

	if err = d.Set("comments", flattenFirewallPolicyComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "FirewallPolicy-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("custom_log_fields", flattenFirewallPolicyCustomLogFields(o["custom-log-fields"], d, "custom_log_fields")); err != nil {
		if vv, ok := fortiAPIPatch(o["custom-log-fields"], "FirewallPolicy-CustomLogFields"); ok {
			if err = d.Set("custom_log_fields", vv); err != nil {
				return fmt.Errorf("Error reading custom_log_fields: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading custom_log_fields: %v", err)
		}
	}

	if err = d.Set("decrypted_traffic_mirror", flattenFirewallPolicyDecryptedTrafficMirror(o["decrypted-traffic-mirror"], d, "decrypted_traffic_mirror")); err != nil {
		if vv, ok := fortiAPIPatch(o["decrypted-traffic-mirror"], "FirewallPolicy-DecryptedTrafficMirror"); ok {
			if err = d.Set("decrypted_traffic_mirror", vv); err != nil {
				return fmt.Errorf("Error reading decrypted_traffic_mirror: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading decrypted_traffic_mirror: %v", err)
		}
	}

	if err = d.Set("delay_tcp_npu_session", flattenFirewallPolicyDelayTcpNpuSession(o["delay-tcp-npu-session"], d, "delay_tcp_npu_session")); err != nil {
		if vv, ok := fortiAPIPatch(o["delay-tcp-npu-session"], "FirewallPolicy-DelayTcpNpuSession"); ok {
			if err = d.Set("delay_tcp_npu_session", vv); err != nil {
				return fmt.Errorf("Error reading delay_tcp_npu_session: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading delay_tcp_npu_session: %v", err)
		}
	}

	if err = d.Set("diameter_filter_profile", flattenFirewallPolicyDiameterFilterProfile(o["diameter-filter-profile"], d, "diameter_filter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["diameter-filter-profile"], "FirewallPolicy-DiameterFilterProfile"); ok {
			if err = d.Set("diameter_filter_profile", vv); err != nil {
				return fmt.Errorf("Error reading diameter_filter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diameter_filter_profile: %v", err)
		}
	}

	if err = d.Set("diffserv_copy", flattenFirewallPolicyDiffservCopy(o["diffserv-copy"], d, "diffserv_copy")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffserv-copy"], "FirewallPolicy-DiffservCopy"); ok {
			if err = d.Set("diffserv_copy", vv); err != nil {
				return fmt.Errorf("Error reading diffserv_copy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffserv_copy: %v", err)
		}
	}

	if err = d.Set("diffserv_forward", flattenFirewallPolicyDiffservForward(o["diffserv-forward"], d, "diffserv_forward")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffserv-forward"], "FirewallPolicy-DiffservForward"); ok {
			if err = d.Set("diffserv_forward", vv); err != nil {
				return fmt.Errorf("Error reading diffserv_forward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffserv_forward: %v", err)
		}
	}

	if err = d.Set("diffserv_reverse", flattenFirewallPolicyDiffservReverse(o["diffserv-reverse"], d, "diffserv_reverse")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffserv-reverse"], "FirewallPolicy-DiffservReverse"); ok {
			if err = d.Set("diffserv_reverse", vv); err != nil {
				return fmt.Errorf("Error reading diffserv_reverse: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffserv_reverse: %v", err)
		}
	}

	if err = d.Set("diffservcode_forward", flattenFirewallPolicyDiffservcodeForward(o["diffservcode-forward"], d, "diffservcode_forward")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffservcode-forward"], "FirewallPolicy-DiffservcodeForward"); ok {
			if err = d.Set("diffservcode_forward", vv); err != nil {
				return fmt.Errorf("Error reading diffservcode_forward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffservcode_forward: %v", err)
		}
	}

	if err = d.Set("diffservcode_rev", flattenFirewallPolicyDiffservcodeRev(o["diffservcode-rev"], d, "diffservcode_rev")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffservcode-rev"], "FirewallPolicy-DiffservcodeRev"); ok {
			if err = d.Set("diffservcode_rev", vv); err != nil {
				return fmt.Errorf("Error reading diffservcode_rev: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffservcode_rev: %v", err)
		}
	}

	if err = d.Set("disclaimer", flattenFirewallPolicyDisclaimer(o["disclaimer"], d, "disclaimer")); err != nil {
		if vv, ok := fortiAPIPatch(o["disclaimer"], "FirewallPolicy-Disclaimer"); ok {
			if err = d.Set("disclaimer", vv); err != nil {
				return fmt.Errorf("Error reading disclaimer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading disclaimer: %v", err)
		}
	}

	if err = d.Set("dlp_profile", flattenFirewallPolicyDlpProfile(o["dlp-profile"], d, "dlp_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-profile"], "FirewallPolicy-DlpProfile"); ok {
			if err = d.Set("dlp_profile", vv); err != nil {
				return fmt.Errorf("Error reading dlp_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_profile: %v", err)
		}
	}

	if err = d.Set("dnsfilter_profile", flattenFirewallPolicyDnsfilterProfile(o["dnsfilter-profile"], d, "dnsfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["dnsfilter-profile"], "FirewallPolicy-DnsfilterProfile"); ok {
			if err = d.Set("dnsfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading dnsfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dnsfilter_profile: %v", err)
		}
	}

	if err = d.Set("dsri", flattenFirewallPolicyDsri(o["dsri"], d, "dsri")); err != nil {
		if vv, ok := fortiAPIPatch(o["dsri"], "FirewallPolicy-Dsri"); ok {
			if err = d.Set("dsri", vv); err != nil {
				return fmt.Errorf("Error reading dsri: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dsri: %v", err)
		}
	}

	if err = d.Set("dstaddr", flattenFirewallPolicyDstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "FirewallPolicy-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("dstaddr_negate", flattenFirewallPolicyDstaddrNegate(o["dstaddr-negate"], d, "dstaddr_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr-negate"], "FirewallPolicy-DstaddrNegate"); ok {
			if err = d.Set("dstaddr_negate", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr_negate: %v", err)
		}
	}

	if err = d.Set("dstaddr6", flattenFirewallPolicyDstaddr6(o["dstaddr6"], d, "dstaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr6"], "FirewallPolicy-Dstaddr6"); ok {
			if err = d.Set("dstaddr6", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr6: %v", err)
		}
	}

	if err = d.Set("dstaddr6_negate", flattenFirewallPolicyDstaddr6Negate(o["dstaddr6-negate"], d, "dstaddr6_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr6-negate"], "FirewallPolicy-Dstaddr6Negate"); ok {
			if err = d.Set("dstaddr6_negate", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr6_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr6_negate: %v", err)
		}
	}

	if err = d.Set("dstintf", flattenFirewallPolicyDstintf(o["dstintf"], d, "dstintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstintf"], "FirewallPolicy-Dstintf"); ok {
			if err = d.Set("dstintf", vv); err != nil {
				return fmt.Errorf("Error reading dstintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstintf: %v", err)
		}
	}

	if err = d.Set("dynamic_shaping", flattenFirewallPolicyDynamicShaping(o["dynamic-shaping"], d, "dynamic_shaping")); err != nil {
		if vv, ok := fortiAPIPatch(o["dynamic-shaping"], "FirewallPolicy-DynamicShaping"); ok {
			if err = d.Set("dynamic_shaping", vv); err != nil {
				return fmt.Errorf("Error reading dynamic_shaping: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dynamic_shaping: %v", err)
		}
	}

	if err = d.Set("eif_check", flattenFirewallPolicyEifCheck(o["eif-check"], d, "eif_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["eif-check"], "FirewallPolicy-EifCheck"); ok {
			if err = d.Set("eif_check", vv); err != nil {
				return fmt.Errorf("Error reading eif_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eif_check: %v", err)
		}
	}

	if err = d.Set("eif_learn", flattenFirewallPolicyEifLearn(o["eif-learn"], d, "eif_learn")); err != nil {
		if vv, ok := fortiAPIPatch(o["eif-learn"], "FirewallPolicy-EifLearn"); ok {
			if err = d.Set("eif_learn", vv); err != nil {
				return fmt.Errorf("Error reading eif_learn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eif_learn: %v", err)
		}
	}

	if err = d.Set("email_collect", flattenFirewallPolicyEmailCollect(o["email-collect"], d, "email_collect")); err != nil {
		if vv, ok := fortiAPIPatch(o["email-collect"], "FirewallPolicy-EmailCollect"); ok {
			if err = d.Set("email_collect", vv); err != nil {
				return fmt.Errorf("Error reading email_collect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading email_collect: %v", err)
		}
	}

	if err = d.Set("emailfilter_profile", flattenFirewallPolicyEmailfilterProfile(o["emailfilter-profile"], d, "emailfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["emailfilter-profile"], "FirewallPolicy-EmailfilterProfile"); ok {
			if err = d.Set("emailfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading emailfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading emailfilter_profile: %v", err)
		}
	}

	if err = d.Set("fec", flattenFirewallPolicyFec(o["fec"], d, "fec")); err != nil {
		if vv, ok := fortiAPIPatch(o["fec"], "FirewallPolicy-Fec"); ok {
			if err = d.Set("fec", vv); err != nil {
				return fmt.Errorf("Error reading fec: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fec: %v", err)
		}
	}

	if err = d.Set("file_filter_profile", flattenFirewallPolicyFileFilterProfile(o["file-filter-profile"], d, "file_filter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-filter-profile"], "FirewallPolicy-FileFilterProfile"); ok {
			if err = d.Set("file_filter_profile", vv); err != nil {
				return fmt.Errorf("Error reading file_filter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_filter_profile: %v", err)
		}
	}

	if err = d.Set("firewall_session_dirty", flattenFirewallPolicyFirewallSessionDirty(o["firewall-session-dirty"], d, "firewall_session_dirty")); err != nil {
		if vv, ok := fortiAPIPatch(o["firewall-session-dirty"], "FirewallPolicy-FirewallSessionDirty"); ok {
			if err = d.Set("firewall_session_dirty", vv); err != nil {
				return fmt.Errorf("Error reading firewall_session_dirty: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading firewall_session_dirty: %v", err)
		}
	}

	if err = d.Set("fixedport", flattenFirewallPolicyFixedport(o["fixedport"], d, "fixedport")); err != nil {
		if vv, ok := fortiAPIPatch(o["fixedport"], "FirewallPolicy-Fixedport"); ok {
			if err = d.Set("fixedport", vv); err != nil {
				return fmt.Errorf("Error reading fixedport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fixedport: %v", err)
		}
	}

	if err = d.Set("fsso_agent_for_ntlm", flattenFirewallPolicyFssoAgentForNtlm(o["fsso-agent-for-ntlm"], d, "fsso_agent_for_ntlm")); err != nil {
		if vv, ok := fortiAPIPatch(o["fsso-agent-for-ntlm"], "FirewallPolicy-FssoAgentForNtlm"); ok {
			if err = d.Set("fsso_agent_for_ntlm", vv); err != nil {
				return fmt.Errorf("Error reading fsso_agent_for_ntlm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fsso_agent_for_ntlm: %v", err)
		}
	}

	if err = d.Set("fsso_groups", flattenFirewallPolicyFssoGroups(o["fsso-groups"], d, "fsso_groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["fsso-groups"], "FirewallPolicy-FssoGroups"); ok {
			if err = d.Set("fsso_groups", vv); err != nil {
				return fmt.Errorf("Error reading fsso_groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fsso_groups: %v", err)
		}
	}

	if err = d.Set("geoip_anycast", flattenFirewallPolicyGeoipAnycast(o["geoip-anycast"], d, "geoip_anycast")); err != nil {
		if vv, ok := fortiAPIPatch(o["geoip-anycast"], "FirewallPolicy-GeoipAnycast"); ok {
			if err = d.Set("geoip_anycast", vv); err != nil {
				return fmt.Errorf("Error reading geoip_anycast: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading geoip_anycast: %v", err)
		}
	}

	if err = d.Set("geoip_match", flattenFirewallPolicyGeoipMatch(o["geoip-match"], d, "geoip_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["geoip-match"], "FirewallPolicy-GeoipMatch"); ok {
			if err = d.Set("geoip_match", vv); err != nil {
				return fmt.Errorf("Error reading geoip_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading geoip_match: %v", err)
		}
	}

	if err = d.Set("global_label", flattenFirewallPolicyGlobalLabel(o["global-label"], d, "global_label")); err != nil {
		if vv, ok := fortiAPIPatch(o["global-label"], "FirewallPolicy-GlobalLabel"); ok {
			if err = d.Set("global_label", vv); err != nil {
				return fmt.Errorf("Error reading global_label: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading global_label: %v", err)
		}
	}

	if err = d.Set("groups", flattenFirewallPolicyGroups(o["groups"], d, "groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["groups"], "FirewallPolicy-Groups"); ok {
			if err = d.Set("groups", vv); err != nil {
				return fmt.Errorf("Error reading groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading groups: %v", err)
		}
	}

	if err = d.Set("gtp_profile", flattenFirewallPolicyGtpProfile(o["gtp-profile"], d, "gtp_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["gtp-profile"], "FirewallPolicy-GtpProfile"); ok {
			if err = d.Set("gtp_profile", vv); err != nil {
				return fmt.Errorf("Error reading gtp_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gtp_profile: %v", err)
		}
	}

	if err = d.Set("http_policy_redirect", flattenFirewallPolicyHttpPolicyRedirect(o["http-policy-redirect"], d, "http_policy_redirect")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-policy-redirect"], "FirewallPolicy-HttpPolicyRedirect"); ok {
			if err = d.Set("http_policy_redirect", vv); err != nil {
				return fmt.Errorf("Error reading http_policy_redirect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_policy_redirect: %v", err)
		}
	}

	if err = d.Set("icap_profile", flattenFirewallPolicyIcapProfile(o["icap-profile"], d, "icap_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["icap-profile"], "FirewallPolicy-IcapProfile"); ok {
			if err = d.Set("icap_profile", vv); err != nil {
				return fmt.Errorf("Error reading icap_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icap_profile: %v", err)
		}
	}

	if err = d.Set("identity_based_route", flattenFirewallPolicyIdentityBasedRoute(o["identity-based-route"], d, "identity_based_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["identity-based-route"], "FirewallPolicy-IdentityBasedRoute"); ok {
			if err = d.Set("identity_based_route", vv); err != nil {
				return fmt.Errorf("Error reading identity_based_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading identity_based_route: %v", err)
		}
	}

	if err = d.Set("inbound", flattenFirewallPolicyInbound(o["inbound"], d, "inbound")); err != nil {
		if vv, ok := fortiAPIPatch(o["inbound"], "FirewallPolicy-Inbound"); ok {
			if err = d.Set("inbound", vv); err != nil {
				return fmt.Errorf("Error reading inbound: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inbound: %v", err)
		}
	}

	if err = d.Set("inspection_mode", flattenFirewallPolicyInspectionMode(o["inspection-mode"], d, "inspection_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["inspection-mode"], "FirewallPolicy-InspectionMode"); ok {
			if err = d.Set("inspection_mode", vv); err != nil {
				return fmt.Errorf("Error reading inspection_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inspection_mode: %v", err)
		}
	}

	if err = d.Set("internet_service", flattenFirewallPolicyInternetService(o["internet-service"], d, "internet_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service"], "FirewallPolicy-InternetService"); ok {
			if err = d.Set("internet_service", vv); err != nil {
				return fmt.Errorf("Error reading internet_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service: %v", err)
		}
	}

	if err = d.Set("internet_service_custom", flattenFirewallPolicyInternetServiceCustom(o["internet-service-custom"], d, "internet_service_custom")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-custom"], "FirewallPolicy-InternetServiceCustom"); ok {
			if err = d.Set("internet_service_custom", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_custom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_custom: %v", err)
		}
	}

	if err = d.Set("internet_service_custom_group", flattenFirewallPolicyInternetServiceCustomGroup(o["internet-service-custom-group"], d, "internet_service_custom_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-custom-group"], "FirewallPolicy-InternetServiceCustomGroup"); ok {
			if err = d.Set("internet_service_custom_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_custom_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_custom_group: %v", err)
		}
	}

	if err = d.Set("internet_service_fortiguard", flattenFirewallPolicyInternetServiceFortiguard(o["internet-service-fortiguard"], d, "internet_service_fortiguard")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-fortiguard"], "FirewallPolicy-InternetServiceFortiguard"); ok {
			if err = d.Set("internet_service_fortiguard", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_fortiguard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_fortiguard: %v", err)
		}
	}

	if err = d.Set("internet_service_group", flattenFirewallPolicyInternetServiceGroup(o["internet-service-group"], d, "internet_service_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-group"], "FirewallPolicy-InternetServiceGroup"); ok {
			if err = d.Set("internet_service_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_group: %v", err)
		}
	}

	if err = d.Set("internet_service_name", flattenFirewallPolicyInternetServiceName(o["internet-service-name"], d, "internet_service_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-name"], "FirewallPolicy-InternetServiceName"); ok {
			if err = d.Set("internet_service_name", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_name: %v", err)
		}
	}

	if err = d.Set("internet_service_negate", flattenFirewallPolicyInternetServiceNegate(o["internet-service-negate"], d, "internet_service_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-negate"], "FirewallPolicy-InternetServiceNegate"); ok {
			if err = d.Set("internet_service_negate", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_negate: %v", err)
		}
	}

	if err = d.Set("internet_service_src", flattenFirewallPolicyInternetServiceSrc(o["internet-service-src"], d, "internet_service_src")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src"], "FirewallPolicy-InternetServiceSrc"); ok {
			if err = d.Set("internet_service_src", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src: %v", err)
		}
	}

	if err = d.Set("internet_service_src_custom", flattenFirewallPolicyInternetServiceSrcCustom(o["internet-service-src-custom"], d, "internet_service_src_custom")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-custom"], "FirewallPolicy-InternetServiceSrcCustom"); ok {
			if err = d.Set("internet_service_src_custom", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_custom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_custom: %v", err)
		}
	}

	if err = d.Set("internet_service_src_custom_group", flattenFirewallPolicyInternetServiceSrcCustomGroup(o["internet-service-src-custom-group"], d, "internet_service_src_custom_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-custom-group"], "FirewallPolicy-InternetServiceSrcCustomGroup"); ok {
			if err = d.Set("internet_service_src_custom_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_custom_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_custom_group: %v", err)
		}
	}

	if err = d.Set("internet_service_src_fortiguard", flattenFirewallPolicyInternetServiceSrcFortiguard(o["internet-service-src-fortiguard"], d, "internet_service_src_fortiguard")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-fortiguard"], "FirewallPolicy-InternetServiceSrcFortiguard"); ok {
			if err = d.Set("internet_service_src_fortiguard", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_fortiguard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_fortiguard: %v", err)
		}
	}

	if err = d.Set("internet_service_src_group", flattenFirewallPolicyInternetServiceSrcGroup(o["internet-service-src-group"], d, "internet_service_src_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-group"], "FirewallPolicy-InternetServiceSrcGroup"); ok {
			if err = d.Set("internet_service_src_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_group: %v", err)
		}
	}

	if err = d.Set("internet_service_src_name", flattenFirewallPolicyInternetServiceSrcName(o["internet-service-src-name"], d, "internet_service_src_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-name"], "FirewallPolicy-InternetServiceSrcName"); ok {
			if err = d.Set("internet_service_src_name", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_name: %v", err)
		}
	}

	if err = d.Set("internet_service_src_negate", flattenFirewallPolicyInternetServiceSrcNegate(o["internet-service-src-negate"], d, "internet_service_src_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-negate"], "FirewallPolicy-InternetServiceSrcNegate"); ok {
			if err = d.Set("internet_service_src_negate", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_negate: %v", err)
		}
	}

	if err = d.Set("internet_service6", flattenFirewallPolicyInternetService6(o["internet-service6"], d, "internet_service6")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6"], "FirewallPolicy-InternetService6"); ok {
			if err = d.Set("internet_service6", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6: %v", err)
		}
	}

	if err = d.Set("internet_service6_custom", flattenFirewallPolicyInternetService6Custom(o["internet-service6-custom"], d, "internet_service6_custom")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-custom"], "FirewallPolicy-InternetService6Custom"); ok {
			if err = d.Set("internet_service6_custom", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_custom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_custom: %v", err)
		}
	}

	if err = d.Set("internet_service6_custom_group", flattenFirewallPolicyInternetService6CustomGroup(o["internet-service6-custom-group"], d, "internet_service6_custom_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-custom-group"], "FirewallPolicy-InternetService6CustomGroup"); ok {
			if err = d.Set("internet_service6_custom_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_custom_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_custom_group: %v", err)
		}
	}

	if err = d.Set("internet_service6_fortiguard", flattenFirewallPolicyInternetService6Fortiguard(o["internet-service6-fortiguard"], d, "internet_service6_fortiguard")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-fortiguard"], "FirewallPolicy-InternetService6Fortiguard"); ok {
			if err = d.Set("internet_service6_fortiguard", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_fortiguard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_fortiguard: %v", err)
		}
	}

	if err = d.Set("internet_service6_group", flattenFirewallPolicyInternetService6Group(o["internet-service6-group"], d, "internet_service6_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-group"], "FirewallPolicy-InternetService6Group"); ok {
			if err = d.Set("internet_service6_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_group: %v", err)
		}
	}

	if err = d.Set("internet_service6_name", flattenFirewallPolicyInternetService6Name(o["internet-service6-name"], d, "internet_service6_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-name"], "FirewallPolicy-InternetService6Name"); ok {
			if err = d.Set("internet_service6_name", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_name: %v", err)
		}
	}

	if err = d.Set("internet_service6_negate", flattenFirewallPolicyInternetService6Negate(o["internet-service6-negate"], d, "internet_service6_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-negate"], "FirewallPolicy-InternetService6Negate"); ok {
			if err = d.Set("internet_service6_negate", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_negate: %v", err)
		}
	}

	if err = d.Set("internet_service6_src", flattenFirewallPolicyInternetService6Src(o["internet-service6-src"], d, "internet_service6_src")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-src"], "FirewallPolicy-InternetService6Src"); ok {
			if err = d.Set("internet_service6_src", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_src: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_src: %v", err)
		}
	}

	if err = d.Set("internet_service6_src_custom", flattenFirewallPolicyInternetService6SrcCustom(o["internet-service6-src-custom"], d, "internet_service6_src_custom")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-src-custom"], "FirewallPolicy-InternetService6SrcCustom"); ok {
			if err = d.Set("internet_service6_src_custom", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_src_custom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_src_custom: %v", err)
		}
	}

	if err = d.Set("internet_service6_src_custom_group", flattenFirewallPolicyInternetService6SrcCustomGroup(o["internet-service6-src-custom-group"], d, "internet_service6_src_custom_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-src-custom-group"], "FirewallPolicy-InternetService6SrcCustomGroup"); ok {
			if err = d.Set("internet_service6_src_custom_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_src_custom_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_src_custom_group: %v", err)
		}
	}

	if err = d.Set("internet_service6_src_fortiguard", flattenFirewallPolicyInternetService6SrcFortiguard(o["internet-service6-src-fortiguard"], d, "internet_service6_src_fortiguard")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-src-fortiguard"], "FirewallPolicy-InternetService6SrcFortiguard"); ok {
			if err = d.Set("internet_service6_src_fortiguard", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_src_fortiguard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_src_fortiguard: %v", err)
		}
	}

	if err = d.Set("internet_service6_src_group", flattenFirewallPolicyInternetService6SrcGroup(o["internet-service6-src-group"], d, "internet_service6_src_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-src-group"], "FirewallPolicy-InternetService6SrcGroup"); ok {
			if err = d.Set("internet_service6_src_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_src_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_src_group: %v", err)
		}
	}

	if err = d.Set("internet_service6_src_name", flattenFirewallPolicyInternetService6SrcName(o["internet-service6-src-name"], d, "internet_service6_src_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-src-name"], "FirewallPolicy-InternetService6SrcName"); ok {
			if err = d.Set("internet_service6_src_name", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_src_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_src_name: %v", err)
		}
	}

	if err = d.Set("internet_service6_src_negate", flattenFirewallPolicyInternetService6SrcNegate(o["internet-service6-src-negate"], d, "internet_service6_src_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-src-negate"], "FirewallPolicy-InternetService6SrcNegate"); ok {
			if err = d.Set("internet_service6_src_negate", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_src_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_src_negate: %v", err)
		}
	}

	if err = d.Set("ip_version_type", flattenFirewallPolicyIpVersionType(o["ip-version-type"], d, "ip_version_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-version-type"], "FirewallPolicy-IpVersionType"); ok {
			if err = d.Set("ip_version_type", vv); err != nil {
				return fmt.Errorf("Error reading ip_version_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_version_type: %v", err)
		}
	}

	if err = d.Set("ippool", flattenFirewallPolicyIppool(o["ippool"], d, "ippool")); err != nil {
		if vv, ok := fortiAPIPatch(o["ippool"], "FirewallPolicy-Ippool"); ok {
			if err = d.Set("ippool", vv); err != nil {
				return fmt.Errorf("Error reading ippool: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ippool: %v", err)
		}
	}

	if err = d.Set("ips_sensor", flattenFirewallPolicyIpsSensor(o["ips-sensor"], d, "ips_sensor")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-sensor"], "FirewallPolicy-IpsSensor"); ok {
			if err = d.Set("ips_sensor", vv); err != nil {
				return fmt.Errorf("Error reading ips_sensor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_sensor: %v", err)
		}
	}

	if err = d.Set("ips_voip_filter", flattenFirewallPolicyIpsVoipFilter(o["ips-voip-filter"], d, "ips_voip_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-voip-filter"], "FirewallPolicy-IpsVoipFilter"); ok {
			if err = d.Set("ips_voip_filter", vv); err != nil {
				return fmt.Errorf("Error reading ips_voip_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_voip_filter: %v", err)
		}
	}

	if err = d.Set("label", flattenFirewallPolicyLabel(o["label"], d, "label")); err != nil {
		if vv, ok := fortiAPIPatch(o["label"], "FirewallPolicy-Label"); ok {
			if err = d.Set("label", vv); err != nil {
				return fmt.Errorf("Error reading label: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading label: %v", err)
		}
	}

	if err = d.Set("log_http_transaction", flattenFirewallPolicyLogHttpTransaction(o["log-http-transaction"], d, "log_http_transaction")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-http-transaction"], "FirewallPolicy-LogHttpTransaction"); ok {
			if err = d.Set("log_http_transaction", vv); err != nil {
				return fmt.Errorf("Error reading log_http_transaction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_http_transaction: %v", err)
		}
	}

	if err = d.Set("logtraffic", flattenFirewallPolicyLogtraffic(o["logtraffic"], d, "logtraffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["logtraffic"], "FirewallPolicy-Logtraffic"); ok {
			if err = d.Set("logtraffic", vv); err != nil {
				return fmt.Errorf("Error reading logtraffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logtraffic: %v", err)
		}
	}

	if err = d.Set("logtraffic_start", flattenFirewallPolicyLogtrafficStart(o["logtraffic-start"], d, "logtraffic_start")); err != nil {
		if vv, ok := fortiAPIPatch(o["logtraffic-start"], "FirewallPolicy-LogtrafficStart"); ok {
			if err = d.Set("logtraffic_start", vv); err != nil {
				return fmt.Errorf("Error reading logtraffic_start: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logtraffic_start: %v", err)
		}
	}

	if err = d.Set("match_vip", flattenFirewallPolicyMatchVip(o["match-vip"], d, "match_vip")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-vip"], "FirewallPolicy-MatchVip"); ok {
			if err = d.Set("match_vip", vv); err != nil {
				return fmt.Errorf("Error reading match_vip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_vip: %v", err)
		}
	}

	if err = d.Set("match_vip_only", flattenFirewallPolicyMatchVipOnly(o["match-vip-only"], d, "match_vip_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-vip-only"], "FirewallPolicy-MatchVipOnly"); ok {
			if err = d.Set("match_vip_only", vv); err != nil {
				return fmt.Errorf("Error reading match_vip_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_vip_only: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallPolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("nat", flattenFirewallPolicyNat(o["nat"], d, "nat")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat"], "FirewallPolicy-Nat"); ok {
			if err = d.Set("nat", vv); err != nil {
				return fmt.Errorf("Error reading nat: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat: %v", err)
		}
	}

	if err = d.Set("nat46", flattenFirewallPolicyNat46(o["nat46"], d, "nat46")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat46"], "FirewallPolicy-Nat46"); ok {
			if err = d.Set("nat46", vv); err != nil {
				return fmt.Errorf("Error reading nat46: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat46: %v", err)
		}
	}

	if err = d.Set("nat64", flattenFirewallPolicyNat64(o["nat64"], d, "nat64")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat64"], "FirewallPolicy-Nat64"); ok {
			if err = d.Set("nat64", vv); err != nil {
				return fmt.Errorf("Error reading nat64: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat64: %v", err)
		}
	}

	if err = d.Set("natinbound", flattenFirewallPolicyNatinbound(o["natinbound"], d, "natinbound")); err != nil {
		if vv, ok := fortiAPIPatch(o["natinbound"], "FirewallPolicy-Natinbound"); ok {
			if err = d.Set("natinbound", vv); err != nil {
				return fmt.Errorf("Error reading natinbound: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading natinbound: %v", err)
		}
	}

	if err = d.Set("natip", flattenFirewallPolicyNatip(o["natip"], d, "natip")); err != nil {
		if vv, ok := fortiAPIPatch(o["natip"], "FirewallPolicy-Natip"); ok {
			if err = d.Set("natip", vv); err != nil {
				return fmt.Errorf("Error reading natip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading natip: %v", err)
		}
	}

	if err = d.Set("natoutbound", flattenFirewallPolicyNatoutbound(o["natoutbound"], d, "natoutbound")); err != nil {
		if vv, ok := fortiAPIPatch(o["natoutbound"], "FirewallPolicy-Natoutbound"); ok {
			if err = d.Set("natoutbound", vv); err != nil {
				return fmt.Errorf("Error reading natoutbound: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading natoutbound: %v", err)
		}
	}

	if err = d.Set("network_service_dynamic", flattenFirewallPolicyNetworkServiceDynamic(o["network-service-dynamic"], d, "network_service_dynamic")); err != nil {
		if vv, ok := fortiAPIPatch(o["network-service-dynamic"], "FirewallPolicy-NetworkServiceDynamic"); ok {
			if err = d.Set("network_service_dynamic", vv); err != nil {
				return fmt.Errorf("Error reading network_service_dynamic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading network_service_dynamic: %v", err)
		}
	}

	if err = d.Set("network_service_src_dynamic", flattenFirewallPolicyNetworkServiceSrcDynamic(o["network-service-src-dynamic"], d, "network_service_src_dynamic")); err != nil {
		if vv, ok := fortiAPIPatch(o["network-service-src-dynamic"], "FirewallPolicy-NetworkServiceSrcDynamic"); ok {
			if err = d.Set("network_service_src_dynamic", vv); err != nil {
				return fmt.Errorf("Error reading network_service_src_dynamic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading network_service_src_dynamic: %v", err)
		}
	}

	if err = d.Set("np_acceleration", flattenFirewallPolicyNpAcceleration(o["np-acceleration"], d, "np_acceleration")); err != nil {
		if vv, ok := fortiAPIPatch(o["np-acceleration"], "FirewallPolicy-NpAcceleration"); ok {
			if err = d.Set("np_acceleration", vv); err != nil {
				return fmt.Errorf("Error reading np_acceleration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading np_acceleration: %v", err)
		}
	}

	if err = d.Set("ntlm", flattenFirewallPolicyNtlm(o["ntlm"], d, "ntlm")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntlm"], "FirewallPolicy-Ntlm"); ok {
			if err = d.Set("ntlm", vv); err != nil {
				return fmt.Errorf("Error reading ntlm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntlm: %v", err)
		}
	}

	if err = d.Set("ntlm_enabled_browsers", flattenFirewallPolicyNtlmEnabledBrowsers(o["ntlm-enabled-browsers"], d, "ntlm_enabled_browsers")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntlm-enabled-browsers"], "FirewallPolicy-NtlmEnabledBrowsers"); ok {
			if err = d.Set("ntlm_enabled_browsers", vv); err != nil {
				return fmt.Errorf("Error reading ntlm_enabled_browsers: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntlm_enabled_browsers: %v", err)
		}
	}

	if err = d.Set("ntlm_guest", flattenFirewallPolicyNtlmGuest(o["ntlm-guest"], d, "ntlm_guest")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntlm-guest"], "FirewallPolicy-NtlmGuest"); ok {
			if err = d.Set("ntlm_guest", vv); err != nil {
				return fmt.Errorf("Error reading ntlm_guest: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntlm_guest: %v", err)
		}
	}

	if err = d.Set("outbound", flattenFirewallPolicyOutbound(o["outbound"], d, "outbound")); err != nil {
		if vv, ok := fortiAPIPatch(o["outbound"], "FirewallPolicy-Outbound"); ok {
			if err = d.Set("outbound", vv); err != nil {
				return fmt.Errorf("Error reading outbound: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outbound: %v", err)
		}
	}

	if err = d.Set("passive_wan_health_measurement", flattenFirewallPolicyPassiveWanHealthMeasurement(o["passive-wan-health-measurement"], d, "passive_wan_health_measurement")); err != nil {
		if vv, ok := fortiAPIPatch(o["passive-wan-health-measurement"], "FirewallPolicy-PassiveWanHealthMeasurement"); ok {
			if err = d.Set("passive_wan_health_measurement", vv); err != nil {
				return fmt.Errorf("Error reading passive_wan_health_measurement: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading passive_wan_health_measurement: %v", err)
		}
	}

	if err = d.Set("pcp_inbound", flattenFirewallPolicyPcpInbound(o["pcp-inbound"], d, "pcp_inbound")); err != nil {
		if vv, ok := fortiAPIPatch(o["pcp-inbound"], "FirewallPolicy-PcpInbound"); ok {
			if err = d.Set("pcp_inbound", vv); err != nil {
				return fmt.Errorf("Error reading pcp_inbound: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pcp_inbound: %v", err)
		}
	}

	if err = d.Set("pcp_outbound", flattenFirewallPolicyPcpOutbound(o["pcp-outbound"], d, "pcp_outbound")); err != nil {
		if vv, ok := fortiAPIPatch(o["pcp-outbound"], "FirewallPolicy-PcpOutbound"); ok {
			if err = d.Set("pcp_outbound", vv); err != nil {
				return fmt.Errorf("Error reading pcp_outbound: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pcp_outbound: %v", err)
		}
	}

	if err = d.Set("pcp_poolname", flattenFirewallPolicyPcpPoolname(o["pcp-poolname"], d, "pcp_poolname")); err != nil {
		if vv, ok := fortiAPIPatch(o["pcp-poolname"], "FirewallPolicy-PcpPoolname"); ok {
			if err = d.Set("pcp_poolname", vv); err != nil {
				return fmt.Errorf("Error reading pcp_poolname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pcp_poolname: %v", err)
		}
	}

	if err = d.Set("per_ip_shaper", flattenFirewallPolicyPerIpShaper(o["per-ip-shaper"], d, "per_ip_shaper")); err != nil {
		if vv, ok := fortiAPIPatch(o["per-ip-shaper"], "FirewallPolicy-PerIpShaper"); ok {
			if err = d.Set("per_ip_shaper", vv); err != nil {
				return fmt.Errorf("Error reading per_ip_shaper: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading per_ip_shaper: %v", err)
		}
	}

	if err = d.Set("permit_any_host", flattenFirewallPolicyPermitAnyHost(o["permit-any-host"], d, "permit_any_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["permit-any-host"], "FirewallPolicy-PermitAnyHost"); ok {
			if err = d.Set("permit_any_host", vv); err != nil {
				return fmt.Errorf("Error reading permit_any_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading permit_any_host: %v", err)
		}
	}

	if err = d.Set("permit_stun_host", flattenFirewallPolicyPermitStunHost(o["permit-stun-host"], d, "permit_stun_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["permit-stun-host"], "FirewallPolicy-PermitStunHost"); ok {
			if err = d.Set("permit_stun_host", vv); err != nil {
				return fmt.Errorf("Error reading permit_stun_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading permit_stun_host: %v", err)
		}
	}

	if err = d.Set("pfcp_profile", flattenFirewallPolicyPfcpProfile(o["pfcp-profile"], d, "pfcp_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["pfcp-profile"], "FirewallPolicy-PfcpProfile"); ok {
			if err = d.Set("pfcp_profile", vv); err != nil {
				return fmt.Errorf("Error reading pfcp_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pfcp_profile: %v", err)
		}
	}

	if err = d.Set("policy_behaviour_type", flattenFirewallPolicyPolicyBehaviourType(o["policy-behaviour-type"], d, "policy_behaviour_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-behaviour-type"], "FirewallPolicy-PolicyBehaviourType"); ok {
			if err = d.Set("policy_behaviour_type", vv); err != nil {
				return fmt.Errorf("Error reading policy_behaviour_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_behaviour_type: %v", err)
		}
	}

	if err = d.Set("policy_expiry", flattenFirewallPolicyPolicyExpiry(o["policy-expiry"], d, "policy_expiry")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-expiry"], "FirewallPolicy-PolicyExpiry"); ok {
			if err = d.Set("policy_expiry", vv); err != nil {
				return fmt.Errorf("Error reading policy_expiry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_expiry: %v", err)
		}
	}

	if err = d.Set("policy_expiry_date", flattenFirewallPolicyPolicyExpiryDate(o["policy-expiry-date"], d, "policy_expiry_date")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-expiry-date"], "FirewallPolicy-PolicyExpiryDate"); ok {
			if err = d.Set("policy_expiry_date", vv); err != nil {
				return fmt.Errorf("Error reading policy_expiry_date: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_expiry_date: %v", err)
		}
	}

	if err = d.Set("policy_expiry_date_utc", flattenFirewallPolicyPolicyExpiryDateUtc(o["policy-expiry-date-utc"], d, "policy_expiry_date_utc")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-expiry-date-utc"], "FirewallPolicy-PolicyExpiryDateUtc"); ok {
			if err = d.Set("policy_expiry_date_utc", vv); err != nil {
				return fmt.Errorf("Error reading policy_expiry_date_utc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_expiry_date_utc: %v", err)
		}
	}

	if err = d.Set("policy_offload", flattenFirewallPolicyPolicyOffload(o["policy-offload"], d, "policy_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-offload"], "FirewallPolicy-PolicyOffload"); ok {
			if err = d.Set("policy_offload", vv); err != nil {
				return fmt.Errorf("Error reading policy_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_offload: %v", err)
		}
	}

	if err = d.Set("policyid", flattenFirewallPolicyPolicyid(o["policyid"], d, "policyid")); err != nil {
		if vv, ok := fortiAPIPatch(o["policyid"], "FirewallPolicy-Policyid"); ok {
			if err = d.Set("policyid", vv); err != nil {
				return fmt.Errorf("Error reading policyid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policyid: %v", err)
		}
	}

	if err = d.Set("poolname", flattenFirewallPolicyPoolname(o["poolname"], d, "poolname")); err != nil {
		if vv, ok := fortiAPIPatch(o["poolname"], "FirewallPolicy-Poolname"); ok {
			if err = d.Set("poolname", vv); err != nil {
				return fmt.Errorf("Error reading poolname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poolname: %v", err)
		}
	}

	if err = d.Set("poolname6", flattenFirewallPolicyPoolname6(o["poolname6"], d, "poolname6")); err != nil {
		if vv, ok := fortiAPIPatch(o["poolname6"], "FirewallPolicy-Poolname6"); ok {
			if err = d.Set("poolname6", vv); err != nil {
				return fmt.Errorf("Error reading poolname6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poolname6: %v", err)
		}
	}

	if err = d.Set("port_preserve", flattenFirewallPolicyPortPreserve(o["port-preserve"], d, "port_preserve")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-preserve"], "FirewallPolicy-PortPreserve"); ok {
			if err = d.Set("port_preserve", vv); err != nil {
				return fmt.Errorf("Error reading port_preserve: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_preserve: %v", err)
		}
	}

	if err = d.Set("port_random", flattenFirewallPolicyPortRandom(o["port-random"], d, "port_random")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-random"], "FirewallPolicy-PortRandom"); ok {
			if err = d.Set("port_random", vv); err != nil {
				return fmt.Errorf("Error reading port_random: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_random: %v", err)
		}
	}

	if err = d.Set("profile_group", flattenFirewallPolicyProfileGroup(o["profile-group"], d, "profile_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-group"], "FirewallPolicy-ProfileGroup"); ok {
			if err = d.Set("profile_group", vv); err != nil {
				return fmt.Errorf("Error reading profile_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_group: %v", err)
		}
	}

	if err = d.Set("profile_protocol_options", flattenFirewallPolicyProfileProtocolOptions(o["profile-protocol-options"], d, "profile_protocol_options")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-protocol-options"], "FirewallPolicy-ProfileProtocolOptions"); ok {
			if err = d.Set("profile_protocol_options", vv); err != nil {
				return fmt.Errorf("Error reading profile_protocol_options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_protocol_options: %v", err)
		}
	}

	if err = d.Set("profile_type", flattenFirewallPolicyProfileType(o["profile-type"], d, "profile_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-type"], "FirewallPolicy-ProfileType"); ok {
			if err = d.Set("profile_type", vv); err != nil {
				return fmt.Errorf("Error reading profile_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_type: %v", err)
		}
	}

	if err = d.Set("radius_ip_auth_bypass", flattenFirewallPolicyRadiusIpAuthBypass(o["radius-ip-auth-bypass"], d, "radius_ip_auth_bypass")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-ip-auth-bypass"], "FirewallPolicy-RadiusIpAuthBypass"); ok {
			if err = d.Set("radius_ip_auth_bypass", vv); err != nil {
				return fmt.Errorf("Error reading radius_ip_auth_bypass: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_ip_auth_bypass: %v", err)
		}
	}

	if err = d.Set("radius_mac_auth_bypass", flattenFirewallPolicyRadiusMacAuthBypass(o["radius-mac-auth-bypass"], d, "radius_mac_auth_bypass")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-mac-auth-bypass"], "FirewallPolicy-RadiusMacAuthBypass"); ok {
			if err = d.Set("radius_mac_auth_bypass", vv); err != nil {
				return fmt.Errorf("Error reading radius_mac_auth_bypass: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_mac_auth_bypass: %v", err)
		}
	}

	if err = d.Set("redirect_url", flattenFirewallPolicyRedirectUrl(o["redirect-url"], d, "redirect_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["redirect-url"], "FirewallPolicy-RedirectUrl"); ok {
			if err = d.Set("redirect_url", vv); err != nil {
				return fmt.Errorf("Error reading redirect_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redirect_url: %v", err)
		}
	}

	if err = d.Set("replacemsg_override_group", flattenFirewallPolicyReplacemsgOverrideGroup(o["replacemsg-override-group"], d, "replacemsg_override_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["replacemsg-override-group"], "FirewallPolicy-ReplacemsgOverrideGroup"); ok {
			if err = d.Set("replacemsg_override_group", vv); err != nil {
				return fmt.Errorf("Error reading replacemsg_override_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replacemsg_override_group: %v", err)
		}
	}

	if err = d.Set("reputation_direction", flattenFirewallPolicyReputationDirection(o["reputation-direction"], d, "reputation_direction")); err != nil {
		if vv, ok := fortiAPIPatch(o["reputation-direction"], "FirewallPolicy-ReputationDirection"); ok {
			if err = d.Set("reputation_direction", vv); err != nil {
				return fmt.Errorf("Error reading reputation_direction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reputation_direction: %v", err)
		}
	}

	if err = d.Set("reputation_direction6", flattenFirewallPolicyReputationDirection6(o["reputation-direction6"], d, "reputation_direction6")); err != nil {
		if vv, ok := fortiAPIPatch(o["reputation-direction6"], "FirewallPolicy-ReputationDirection6"); ok {
			if err = d.Set("reputation_direction6", vv); err != nil {
				return fmt.Errorf("Error reading reputation_direction6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reputation_direction6: %v", err)
		}
	}

	if err = d.Set("reputation_minimum", flattenFirewallPolicyReputationMinimum(o["reputation-minimum"], d, "reputation_minimum")); err != nil {
		if vv, ok := fortiAPIPatch(o["reputation-minimum"], "FirewallPolicy-ReputationMinimum"); ok {
			if err = d.Set("reputation_minimum", vv); err != nil {
				return fmt.Errorf("Error reading reputation_minimum: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reputation_minimum: %v", err)
		}
	}

	if err = d.Set("reputation_minimum6", flattenFirewallPolicyReputationMinimum6(o["reputation-minimum6"], d, "reputation_minimum6")); err != nil {
		if vv, ok := fortiAPIPatch(o["reputation-minimum6"], "FirewallPolicy-ReputationMinimum6"); ok {
			if err = d.Set("reputation_minimum6", vv); err != nil {
				return fmt.Errorf("Error reading reputation_minimum6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reputation_minimum6: %v", err)
		}
	}

	if err = d.Set("rtp_addr", flattenFirewallPolicyRtpAddr(o["rtp-addr"], d, "rtp_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["rtp-addr"], "FirewallPolicy-RtpAddr"); ok {
			if err = d.Set("rtp_addr", vv); err != nil {
				return fmt.Errorf("Error reading rtp_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rtp_addr: %v", err)
		}
	}

	if err = d.Set("rtp_nat", flattenFirewallPolicyRtpNat(o["rtp-nat"], d, "rtp_nat")); err != nil {
		if vv, ok := fortiAPIPatch(o["rtp-nat"], "FirewallPolicy-RtpNat"); ok {
			if err = d.Set("rtp_nat", vv); err != nil {
				return fmt.Errorf("Error reading rtp_nat: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rtp_nat: %v", err)
		}
	}

	if err = d.Set("saml_server", flattenFirewallPolicySamlServer(o["saml-server"], d, "saml_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["saml-server"], "FirewallPolicy-SamlServer"); ok {
			if err = d.Set("saml_server", vv); err != nil {
				return fmt.Errorf("Error reading saml_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading saml_server: %v", err)
		}
	}

	if err = d.Set("schedule", flattenFirewallPolicySchedule(o["schedule"], d, "schedule")); err != nil {
		if vv, ok := fortiAPIPatch(o["schedule"], "FirewallPolicy-Schedule"); ok {
			if err = d.Set("schedule", vv); err != nil {
				return fmt.Errorf("Error reading schedule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("schedule_timeout", flattenFirewallPolicyScheduleTimeout(o["schedule-timeout"], d, "schedule_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["schedule-timeout"], "FirewallPolicy-ScheduleTimeout"); ok {
			if err = d.Set("schedule_timeout", vv); err != nil {
				return fmt.Errorf("Error reading schedule_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading schedule_timeout: %v", err)
		}
	}

	if err = d.Set("scim", flattenFirewallPolicyScim(o["scim"], d, "scim")); err != nil {
		if vv, ok := fortiAPIPatch(o["scim"], "FirewallPolicy-Scim"); ok {
			if err = d.Set("scim", vv); err != nil {
				return fmt.Errorf("Error reading scim: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scim: %v", err)
		}
	}

	if err = d.Set("scim_groups", flattenFirewallPolicyScimGroups(o["scim-groups"], d, "scim_groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["scim-groups"], "FirewallPolicy-ScimGroups"); ok {
			if err = d.Set("scim_groups", vv); err != nil {
				return fmt.Errorf("Error reading scim_groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scim_groups: %v", err)
		}
	}

	if err = d.Set("scim_users", flattenFirewallPolicyScimUsers(o["scim-users"], d, "scim_users")); err != nil {
		if vv, ok := fortiAPIPatch(o["scim-users"], "FirewallPolicy-ScimUsers"); ok {
			if err = d.Set("scim_users", vv); err != nil {
				return fmt.Errorf("Error reading scim_users: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scim_users: %v", err)
		}
	}

	if err = d.Set("sctp_filter_profile", flattenFirewallPolicySctpFilterProfile(o["sctp-filter-profile"], d, "sctp_filter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["sctp-filter-profile"], "FirewallPolicy-SctpFilterProfile"); ok {
			if err = d.Set("sctp_filter_profile", vv); err != nil {
				return fmt.Errorf("Error reading sctp_filter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sctp_filter_profile: %v", err)
		}
	}

	if err = d.Set("send_deny_packet", flattenFirewallPolicySendDenyPacket(o["send-deny-packet"], d, "send_deny_packet")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-deny-packet"], "FirewallPolicy-SendDenyPacket"); ok {
			if err = d.Set("send_deny_packet", vv); err != nil {
				return fmt.Errorf("Error reading send_deny_packet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_deny_packet: %v", err)
		}
	}

	if err = d.Set("service", flattenFirewallPolicyService(o["service"], d, "service")); err != nil {
		if vv, ok := fortiAPIPatch(o["service"], "FirewallPolicy-Service"); ok {
			if err = d.Set("service", vv); err != nil {
				return fmt.Errorf("Error reading service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("service_negate", flattenFirewallPolicyServiceNegate(o["service-negate"], d, "service_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-negate"], "FirewallPolicy-ServiceNegate"); ok {
			if err = d.Set("service_negate", vv); err != nil {
				return fmt.Errorf("Error reading service_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_negate: %v", err)
		}
	}

	if err = d.Set("session_ttl", flattenFirewallPolicySessionTtl(o["session-ttl"], d, "session_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["session-ttl"], "FirewallPolicy-SessionTtl"); ok {
			if err = d.Set("session_ttl", vv); err != nil {
				return fmt.Errorf("Error reading session_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session_ttl: %v", err)
		}
	}

	if err = d.Set("sgt", flattenFirewallPolicySgt(o["sgt"], d, "sgt")); err != nil {
		if vv, ok := fortiAPIPatch(o["sgt"], "FirewallPolicy-Sgt"); ok {
			if err = d.Set("sgt", vv); err != nil {
				return fmt.Errorf("Error reading sgt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sgt: %v", err)
		}
	}

	if err = d.Set("sgt_check", flattenFirewallPolicySgtCheck(o["sgt-check"], d, "sgt_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["sgt-check"], "FirewallPolicy-SgtCheck"); ok {
			if err = d.Set("sgt_check", vv); err != nil {
				return fmt.Errorf("Error reading sgt_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sgt_check: %v", err)
		}
	}

	if err = d.Set("src_vendor_mac", flattenFirewallPolicySrcVendorMac(o["src-vendor-mac"], d, "src_vendor_mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-vendor-mac"], "FirewallPolicy-SrcVendorMac"); ok {
			if err = d.Set("src_vendor_mac", vv); err != nil {
				return fmt.Errorf("Error reading src_vendor_mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_vendor_mac: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenFirewallPolicySrcaddr(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "FirewallPolicy-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("srcaddr_negate", flattenFirewallPolicySrcaddrNegate(o["srcaddr-negate"], d, "srcaddr_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr-negate"], "FirewallPolicy-SrcaddrNegate"); ok {
			if err = d.Set("srcaddr_negate", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr_negate: %v", err)
		}
	}

	if err = d.Set("srcaddr6", flattenFirewallPolicySrcaddr6(o["srcaddr6"], d, "srcaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr6"], "FirewallPolicy-Srcaddr6"); ok {
			if err = d.Set("srcaddr6", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr6: %v", err)
		}
	}

	if err = d.Set("srcaddr6_negate", flattenFirewallPolicySrcaddr6Negate(o["srcaddr6-negate"], d, "srcaddr6_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr6-negate"], "FirewallPolicy-Srcaddr6Negate"); ok {
			if err = d.Set("srcaddr6_negate", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr6_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr6_negate: %v", err)
		}
	}

	if err = d.Set("srcintf", flattenFirewallPolicySrcintf(o["srcintf"], d, "srcintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcintf"], "FirewallPolicy-Srcintf"); ok {
			if err = d.Set("srcintf", vv); err != nil {
				return fmt.Errorf("Error reading srcintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcintf: %v", err)
		}
	}

	if err = d.Set("ssh_filter_profile", flattenFirewallPolicySshFilterProfile(o["ssh-filter-profile"], d, "ssh_filter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-filter-profile"], "FirewallPolicy-SshFilterProfile"); ok {
			if err = d.Set("ssh_filter_profile", vv); err != nil {
				return fmt.Errorf("Error reading ssh_filter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_filter_profile: %v", err)
		}
	}

	if err = d.Set("ssh_policy_redirect", flattenFirewallPolicySshPolicyRedirect(o["ssh-policy-redirect"], d, "ssh_policy_redirect")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-policy-redirect"], "FirewallPolicy-SshPolicyRedirect"); ok {
			if err = d.Set("ssh_policy_redirect", vv); err != nil {
				return fmt.Errorf("Error reading ssh_policy_redirect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_policy_redirect: %v", err)
		}
	}

	if err = d.Set("ssl_ssh_profile", flattenFirewallPolicySslSshProfile(o["ssl-ssh-profile"], d, "ssl_ssh_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-ssh-profile"], "FirewallPolicy-SslSshProfile"); ok {
			if err = d.Set("ssl_ssh_profile", vv); err != nil {
				return fmt.Errorf("Error reading ssl_ssh_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_ssh_profile: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallPolicyStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FirewallPolicy-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tcp_mss_receiver", flattenFirewallPolicyTcpMssReceiver(o["tcp-mss-receiver"], d, "tcp_mss_receiver")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-mss-receiver"], "FirewallPolicy-TcpMssReceiver"); ok {
			if err = d.Set("tcp_mss_receiver", vv); err != nil {
				return fmt.Errorf("Error reading tcp_mss_receiver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_mss_receiver: %v", err)
		}
	}

	if err = d.Set("tcp_mss_sender", flattenFirewallPolicyTcpMssSender(o["tcp-mss-sender"], d, "tcp_mss_sender")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-mss-sender"], "FirewallPolicy-TcpMssSender"); ok {
			if err = d.Set("tcp_mss_sender", vv); err != nil {
				return fmt.Errorf("Error reading tcp_mss_sender: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_mss_sender: %v", err)
		}
	}

	if err = d.Set("tcp_session_without_syn", flattenFirewallPolicyTcpSessionWithoutSyn(o["tcp-session-without-syn"], d, "tcp_session_without_syn")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-session-without-syn"], "FirewallPolicy-TcpSessionWithoutSyn"); ok {
			if err = d.Set("tcp_session_without_syn", vv); err != nil {
				return fmt.Errorf("Error reading tcp_session_without_syn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_session_without_syn: %v", err)
		}
	}

	if err = d.Set("tcp_timeout_pid", flattenFirewallPolicyTcpTimeoutPid(o["tcp-timeout-pid"], d, "tcp_timeout_pid")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-timeout-pid"], "FirewallPolicy-TcpTimeoutPid"); ok {
			if err = d.Set("tcp_timeout_pid", vv); err != nil {
				return fmt.Errorf("Error reading tcp_timeout_pid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_timeout_pid: %v", err)
		}
	}

	if err = d.Set("telemetry_profile", flattenFirewallPolicyTelemetryProfile(o["telemetry-profile"], d, "telemetry_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["telemetry-profile"], "FirewallPolicy-TelemetryProfile"); ok {
			if err = d.Set("telemetry_profile", vv); err != nil {
				return fmt.Errorf("Error reading telemetry_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading telemetry_profile: %v", err)
		}
	}

	if err = d.Set("timeout_send_rst", flattenFirewallPolicyTimeoutSendRst(o["timeout-send-rst"], d, "timeout_send_rst")); err != nil {
		if vv, ok := fortiAPIPatch(o["timeout-send-rst"], "FirewallPolicy-TimeoutSendRst"); ok {
			if err = d.Set("timeout_send_rst", vv); err != nil {
				return fmt.Errorf("Error reading timeout_send_rst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timeout_send_rst: %v", err)
		}
	}

	if err = d.Set("tos", flattenFirewallPolicyTos(o["tos"], d, "tos")); err != nil {
		if vv, ok := fortiAPIPatch(o["tos"], "FirewallPolicy-Tos"); ok {
			if err = d.Set("tos", vv); err != nil {
				return fmt.Errorf("Error reading tos: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tos: %v", err)
		}
	}

	if err = d.Set("tos_mask", flattenFirewallPolicyTosMask(o["tos-mask"], d, "tos_mask")); err != nil {
		if vv, ok := fortiAPIPatch(o["tos-mask"], "FirewallPolicy-TosMask"); ok {
			if err = d.Set("tos_mask", vv); err != nil {
				return fmt.Errorf("Error reading tos_mask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tos_mask: %v", err)
		}
	}

	if err = d.Set("tos_negate", flattenFirewallPolicyTosNegate(o["tos-negate"], d, "tos_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["tos-negate"], "FirewallPolicy-TosNegate"); ok {
			if err = d.Set("tos_negate", vv); err != nil {
				return fmt.Errorf("Error reading tos_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tos_negate: %v", err)
		}
	}

	if err = d.Set("traffic_shaper", flattenFirewallPolicyTrafficShaper(o["traffic-shaper"], d, "traffic_shaper")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-shaper"], "FirewallPolicy-TrafficShaper"); ok {
			if err = d.Set("traffic_shaper", vv); err != nil {
				return fmt.Errorf("Error reading traffic_shaper: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_shaper: %v", err)
		}
	}

	if err = d.Set("traffic_shaper_reverse", flattenFirewallPolicyTrafficShaperReverse(o["traffic-shaper-reverse"], d, "traffic_shaper_reverse")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-shaper-reverse"], "FirewallPolicy-TrafficShaperReverse"); ok {
			if err = d.Set("traffic_shaper_reverse", vv); err != nil {
				return fmt.Errorf("Error reading traffic_shaper_reverse: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_shaper_reverse: %v", err)
		}
	}

	if err = d.Set("udp_timeout_pid", flattenFirewallPolicyUdpTimeoutPid(o["udp-timeout-pid"], d, "udp_timeout_pid")); err != nil {
		if vv, ok := fortiAPIPatch(o["udp-timeout-pid"], "FirewallPolicy-UdpTimeoutPid"); ok {
			if err = d.Set("udp_timeout_pid", vv); err != nil {
				return fmt.Errorf("Error reading udp_timeout_pid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udp_timeout_pid: %v", err)
		}
	}

	if err = d.Set("users", flattenFirewallPolicyUsers(o["users"], d, "users")); err != nil {
		if vv, ok := fortiAPIPatch(o["users"], "FirewallPolicy-Users"); ok {
			if err = d.Set("users", vv); err != nil {
				return fmt.Errorf("Error reading users: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading users: %v", err)
		}
	}

	if err = d.Set("utm_status", flattenFirewallPolicyUtmStatus(o["utm-status"], d, "utm_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["utm-status"], "FirewallPolicy-UtmStatus"); ok {
			if err = d.Set("utm_status", vv); err != nil {
				return fmt.Errorf("Error reading utm_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading utm_status: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallPolicyUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "FirewallPolicy-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("videofilter_profile", flattenFirewallPolicyVideofilterProfile(o["videofilter-profile"], d, "videofilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["videofilter-profile"], "FirewallPolicy-VideofilterProfile"); ok {
			if err = d.Set("videofilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading videofilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading videofilter_profile: %v", err)
		}
	}

	if err = d.Set("virtual_patch_profile", flattenFirewallPolicyVirtualPatchProfile(o["virtual-patch-profile"], d, "virtual_patch_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["virtual-patch-profile"], "FirewallPolicy-VirtualPatchProfile"); ok {
			if err = d.Set("virtual_patch_profile", vv); err != nil {
				return fmt.Errorf("Error reading virtual_patch_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading virtual_patch_profile: %v", err)
		}
	}

	if err = d.Set("vlan_cos_fwd", flattenFirewallPolicyVlanCosFwd(o["vlan-cos-fwd"], d, "vlan_cos_fwd")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-cos-fwd"], "FirewallPolicy-VlanCosFwd"); ok {
			if err = d.Set("vlan_cos_fwd", vv); err != nil {
				return fmt.Errorf("Error reading vlan_cos_fwd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_cos_fwd: %v", err)
		}
	}

	if err = d.Set("vlan_cos_rev", flattenFirewallPolicyVlanCosRev(o["vlan-cos-rev"], d, "vlan_cos_rev")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-cos-rev"], "FirewallPolicy-VlanCosRev"); ok {
			if err = d.Set("vlan_cos_rev", vv); err != nil {
				return fmt.Errorf("Error reading vlan_cos_rev: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_cos_rev: %v", err)
		}
	}

	if err = d.Set("vlan_filter", flattenFirewallPolicyVlanFilter(o["vlan-filter"], d, "vlan_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-filter"], "FirewallPolicy-VlanFilter"); ok {
			if err = d.Set("vlan_filter", vv); err != nil {
				return fmt.Errorf("Error reading vlan_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_filter: %v", err)
		}
	}

	if err = d.Set("voip_profile", flattenFirewallPolicyVoipProfile(o["voip-profile"], d, "voip_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["voip-profile"], "FirewallPolicy-VoipProfile"); ok {
			if err = d.Set("voip_profile", vv); err != nil {
				return fmt.Errorf("Error reading voip_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading voip_profile: %v", err)
		}
	}

	if err = d.Set("vpntunnel", flattenFirewallPolicyVpntunnel(o["vpntunnel"], d, "vpntunnel")); err != nil {
		if vv, ok := fortiAPIPatch(o["vpntunnel"], "FirewallPolicy-Vpntunnel"); ok {
			if err = d.Set("vpntunnel", vv); err != nil {
				return fmt.Errorf("Error reading vpntunnel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vpntunnel: %v", err)
		}
	}

	if err = d.Set("waf_profile", flattenFirewallPolicyWafProfile(o["waf-profile"], d, "waf_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["waf-profile"], "FirewallPolicy-WafProfile"); ok {
			if err = d.Set("waf_profile", vv); err != nil {
				return fmt.Errorf("Error reading waf_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading waf_profile: %v", err)
		}
	}

	if err = d.Set("wanopt", flattenFirewallPolicyWanopt(o["wanopt"], d, "wanopt")); err != nil {
		if vv, ok := fortiAPIPatch(o["wanopt"], "FirewallPolicy-Wanopt"); ok {
			if err = d.Set("wanopt", vv); err != nil {
				return fmt.Errorf("Error reading wanopt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wanopt: %v", err)
		}
	}

	if err = d.Set("wanopt_detection", flattenFirewallPolicyWanoptDetection(o["wanopt-detection"], d, "wanopt_detection")); err != nil {
		if vv, ok := fortiAPIPatch(o["wanopt-detection"], "FirewallPolicy-WanoptDetection"); ok {
			if err = d.Set("wanopt_detection", vv); err != nil {
				return fmt.Errorf("Error reading wanopt_detection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wanopt_detection: %v", err)
		}
	}

	if err = d.Set("wanopt_passive_opt", flattenFirewallPolicyWanoptPassiveOpt(o["wanopt-passive-opt"], d, "wanopt_passive_opt")); err != nil {
		if vv, ok := fortiAPIPatch(o["wanopt-passive-opt"], "FirewallPolicy-WanoptPassiveOpt"); ok {
			if err = d.Set("wanopt_passive_opt", vv); err != nil {
				return fmt.Errorf("Error reading wanopt_passive_opt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wanopt_passive_opt: %v", err)
		}
	}

	if err = d.Set("wanopt_peer", flattenFirewallPolicyWanoptPeer(o["wanopt-peer"], d, "wanopt_peer")); err != nil {
		if vv, ok := fortiAPIPatch(o["wanopt-peer"], "FirewallPolicy-WanoptPeer"); ok {
			if err = d.Set("wanopt_peer", vv); err != nil {
				return fmt.Errorf("Error reading wanopt_peer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wanopt_peer: %v", err)
		}
	}

	if err = d.Set("wanopt_profile", flattenFirewallPolicyWanoptProfile(o["wanopt-profile"], d, "wanopt_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["wanopt-profile"], "FirewallPolicy-WanoptProfile"); ok {
			if err = d.Set("wanopt_profile", vv); err != nil {
				return fmt.Errorf("Error reading wanopt_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wanopt_profile: %v", err)
		}
	}

	if err = d.Set("wccp", flattenFirewallPolicyWccp(o["wccp"], d, "wccp")); err != nil {
		if vv, ok := fortiAPIPatch(o["wccp"], "FirewallPolicy-Wccp"); ok {
			if err = d.Set("wccp", vv); err != nil {
				return fmt.Errorf("Error reading wccp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wccp: %v", err)
		}
	}

	if err = d.Set("webcache", flattenFirewallPolicyWebcache(o["webcache"], d, "webcache")); err != nil {
		if vv, ok := fortiAPIPatch(o["webcache"], "FirewallPolicy-Webcache"); ok {
			if err = d.Set("webcache", vv); err != nil {
				return fmt.Errorf("Error reading webcache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webcache: %v", err)
		}
	}

	if err = d.Set("webcache_https", flattenFirewallPolicyWebcacheHttps(o["webcache-https"], d, "webcache_https")); err != nil {
		if vv, ok := fortiAPIPatch(o["webcache-https"], "FirewallPolicy-WebcacheHttps"); ok {
			if err = d.Set("webcache_https", vv); err != nil {
				return fmt.Errorf("Error reading webcache_https: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webcache_https: %v", err)
		}
	}

	if err = d.Set("webfilter_profile", flattenFirewallPolicyWebfilterProfile(o["webfilter-profile"], d, "webfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-profile"], "FirewallPolicy-WebfilterProfile"); ok {
			if err = d.Set("webfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_profile: %v", err)
		}
	}

	if err = d.Set("webproxy_forward_server", flattenFirewallPolicyWebproxyForwardServer(o["webproxy-forward-server"], d, "webproxy_forward_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["webproxy-forward-server"], "FirewallPolicy-WebproxyForwardServer"); ok {
			if err = d.Set("webproxy_forward_server", vv); err != nil {
				return fmt.Errorf("Error reading webproxy_forward_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webproxy_forward_server: %v", err)
		}
	}

	if err = d.Set("webproxy_profile", flattenFirewallPolicyWebproxyProfile(o["webproxy-profile"], d, "webproxy_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["webproxy-profile"], "FirewallPolicy-WebproxyProfile"); ok {
			if err = d.Set("webproxy_profile", vv); err != nil {
				return fmt.Errorf("Error reading webproxy_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webproxy_profile: %v", err)
		}
	}

	if err = d.Set("ztna_device_ownership", flattenFirewallPolicyZtnaDeviceOwnership(o["ztna-device-ownership"], d, "ztna_device_ownership")); err != nil {
		if vv, ok := fortiAPIPatch(o["ztna-device-ownership"], "FirewallPolicy-ZtnaDeviceOwnership"); ok {
			if err = d.Set("ztna_device_ownership", vv); err != nil {
				return fmt.Errorf("Error reading ztna_device_ownership: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ztna_device_ownership: %v", err)
		}
	}

	if err = d.Set("ztna_ems_tag", flattenFirewallPolicyZtnaEmsTag(o["ztna-ems-tag"], d, "ztna_ems_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["ztna-ems-tag"], "FirewallPolicy-ZtnaEmsTag"); ok {
			if err = d.Set("ztna_ems_tag", vv); err != nil {
				return fmt.Errorf("Error reading ztna_ems_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ztna_ems_tag: %v", err)
		}
	}

	if err = d.Set("ztna_ems_tag_negate", flattenFirewallPolicyZtnaEmsTagNegate(o["ztna-ems-tag-negate"], d, "ztna_ems_tag_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["ztna-ems-tag-negate"], "FirewallPolicy-ZtnaEmsTagNegate"); ok {
			if err = d.Set("ztna_ems_tag_negate", vv); err != nil {
				return fmt.Errorf("Error reading ztna_ems_tag_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ztna_ems_tag_negate: %v", err)
		}
	}

	if err = d.Set("ztna_ems_tag_secondary", flattenFirewallPolicyZtnaEmsTagSecondary(o["ztna-ems-tag-secondary"], d, "ztna_ems_tag_secondary")); err != nil {
		if vv, ok := fortiAPIPatch(o["ztna-ems-tag-secondary"], "FirewallPolicy-ZtnaEmsTagSecondary"); ok {
			if err = d.Set("ztna_ems_tag_secondary", vv); err != nil {
				return fmt.Errorf("Error reading ztna_ems_tag_secondary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ztna_ems_tag_secondary: %v", err)
		}
	}

	if err = d.Set("ztna_geo_tag", flattenFirewallPolicyZtnaGeoTag(o["ztna-geo-tag"], d, "ztna_geo_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["ztna-geo-tag"], "FirewallPolicy-ZtnaGeoTag"); ok {
			if err = d.Set("ztna_geo_tag", vv); err != nil {
				return fmt.Errorf("Error reading ztna_geo_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ztna_geo_tag: %v", err)
		}
	}

	if err = d.Set("ztna_policy_redirect", flattenFirewallPolicyZtnaPolicyRedirect(o["ztna-policy-redirect"], d, "ztna_policy_redirect")); err != nil {
		if vv, ok := fortiAPIPatch(o["ztna-policy-redirect"], "FirewallPolicy-ZtnaPolicyRedirect"); ok {
			if err = d.Set("ztna_policy_redirect", vv); err != nil {
				return fmt.Errorf("Error reading ztna_policy_redirect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ztna_policy_redirect: %v", err)
		}
	}

	if err = d.Set("ztna_status", flattenFirewallPolicyZtnaStatus(o["ztna-status"], d, "ztna_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["ztna-status"], "FirewallPolicy-ZtnaStatus"); ok {
			if err = d.Set("ztna_status", vv); err != nil {
				return fmt.Errorf("Error reading ztna_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ztna_status: %v", err)
		}
	}

	if err = d.Set("ztna_tags_match_logic", flattenFirewallPolicyZtnaTagsMatchLogic(o["ztna-tags-match-logic"], d, "ztna_tags_match_logic")); err != nil {
		if vv, ok := fortiAPIPatch(o["ztna-tags-match-logic"], "FirewallPolicy-ZtnaTagsMatchLogic"); ok {
			if err = d.Set("ztna_tags_match_logic", vv); err != nil {
				return fmt.Errorf("Error reading ztna_tags_match_logic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ztna_tags_match_logic: %v", err)
		}
	}

	if err = d.Set("access_proxy", flattenFirewallPolicyAccessProxy(o["access-proxy"], d, "access_proxy")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-proxy"], "FirewallPolicy-AccessProxy"); ok {
			if err = d.Set("access_proxy", vv); err != nil {
				return fmt.Errorf("Error reading access_proxy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_proxy: %v", err)
		}
	}

	if err = d.Set("app_category", flattenFirewallPolicyAppCategory(o["app-category"], d, "app_category")); err != nil {
		if vv, ok := fortiAPIPatch(o["app-category"], "FirewallPolicy-AppCategory"); ok {
			if err = d.Set("app_category", vv); err != nil {
				return fmt.Errorf("Error reading app_category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading app_category: %v", err)
		}
	}

	if err = d.Set("app_group", flattenFirewallPolicyAppGroup(o["app-group"], d, "app_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["app-group"], "FirewallPolicy-AppGroup"); ok {
			if err = d.Set("app_group", vv); err != nil {
				return fmt.Errorf("Error reading app_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading app_group: %v", err)
		}
	}

	if err = d.Set("application", flattenFirewallPolicyApplication(o["application"], d, "application")); err != nil {
		if vv, ok := fortiAPIPatch(o["application"], "FirewallPolicy-Application"); ok {
			if err = d.Set("application", vv); err != nil {
				return fmt.Errorf("Error reading application: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if err = d.Set("detect_https_in_http_request", flattenFirewallPolicyDetectHttpsInHttpRequest(o["detect-https-in-http-request"], d, "detect_https_in_http_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["detect-https-in-http-request"], "FirewallPolicy-DetectHttpsInHttpRequest"); ok {
			if err = d.Set("detect_https_in_http_request", vv); err != nil {
				return fmt.Errorf("Error reading detect_https_in_http_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading detect_https_in_http_request: %v", err)
		}
	}

	if err = d.Set("device_ownership", flattenFirewallPolicyDeviceOwnership(o["device-ownership"], d, "device_ownership")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-ownership"], "FirewallPolicy-DeviceOwnership"); ok {
			if err = d.Set("device_ownership", vv); err != nil {
				return fmt.Errorf("Error reading device_ownership: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_ownership: %v", err)
		}
	}

	if err = d.Set("dlp_sensor", flattenFirewallPolicyDlpSensor(o["dlp-sensor"], d, "dlp_sensor")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-sensor"], "FirewallPolicy-DlpSensor"); ok {
			if err = d.Set("dlp_sensor", vv); err != nil {
				return fmt.Errorf("Error reading dlp_sensor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_sensor: %v", err)
		}
	}

	if err = d.Set("dynamic_bypass", flattenFirewallPolicyDynamicBypass(o["dynamic-bypass"], d, "dynamic_bypass")); err != nil {
		if vv, ok := fortiAPIPatch(o["dynamic-bypass"], "FirewallPolicy-DynamicBypass"); ok {
			if err = d.Set("dynamic_bypass", vv); err != nil {
				return fmt.Errorf("Error reading dynamic_bypass: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dynamic_bypass: %v", err)
		}
	}

	if err = d.Set("explicit_web_proxy", flattenFirewallPolicyExplicitWebProxy(o["explicit-web-proxy"], d, "explicit_web_proxy")); err != nil {
		if vv, ok := fortiAPIPatch(o["explicit-web-proxy"], "FirewallPolicy-ExplicitWebProxy"); ok {
			if err = d.Set("explicit_web_proxy", vv); err != nil {
				return fmt.Errorf("Error reading explicit_web_proxy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading explicit_web_proxy: %v", err)
		}
	}

	if err = d.Set("extended_log", flattenFirewallPolicyExtendedLog(o["extended-log"], d, "extended_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["extended-log"], "FirewallPolicy-ExtendedLog"); ok {
			if err = d.Set("extended_log", vv); err != nil {
				return fmt.Errorf("Error reading extended_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extended_log: %v", err)
		}
	}

	if err = d.Set("force_proxy", flattenFirewallPolicyForceProxy(o["force-proxy"], d, "force_proxy")); err != nil {
		if vv, ok := fortiAPIPatch(o["force-proxy"], "FirewallPolicy-ForceProxy"); ok {
			if err = d.Set("force_proxy", vv); err != nil {
				return fmt.Errorf("Error reading force_proxy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading force_proxy: %v", err)
		}
	}

	if err = d.Set("http_tunnel_auth", flattenFirewallPolicyHttpTunnelAuth(o["http-tunnel-auth"], d, "http_tunnel_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-tunnel-auth"], "FirewallPolicy-HttpTunnelAuth"); ok {
			if err = d.Set("http_tunnel_auth", vv); err != nil {
				return fmt.Errorf("Error reading http_tunnel_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_tunnel_auth: %v", err)
		}
	}

	if err = d.Set("https_sub_category", flattenFirewallPolicyHttpsSubCategory(o["https-sub-category"], d, "https_sub_category")); err != nil {
		if vv, ok := fortiAPIPatch(o["https-sub-category"], "FirewallPolicy-HttpsSubCategory"); ok {
			if err = d.Set("https_sub_category", vv); err != nil {
				return fmt.Errorf("Error reading https_sub_category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading https_sub_category: %v", err)
		}
	}

	if err = d.Set("ia_profile", flattenFirewallPolicyIaProfile(o["ia-profile"], d, "ia_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["ia-profile"], "FirewallPolicy-IaProfile"); ok {
			if err = d.Set("ia_profile", vv); err != nil {
				return fmt.Errorf("Error reading ia_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ia_profile: %v", err)
		}
	}

	if err = d.Set("implicit_proxy_detection", flattenFirewallPolicyImplicitProxyDetection(o["implicit-proxy-detection"], d, "implicit_proxy_detection")); err != nil {
		if vv, ok := fortiAPIPatch(o["implicit-proxy-detection"], "FirewallPolicy-ImplicitProxyDetection"); ok {
			if err = d.Set("implicit_proxy_detection", vv); err != nil {
				return fmt.Errorf("Error reading implicit_proxy_detection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading implicit_proxy_detection: %v", err)
		}
	}

	if err = d.Set("isolator_profile", flattenFirewallPolicyIsolatorProfile(o["isolator-profile"], d, "isolator_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["isolator-profile"], "FirewallPolicy-IsolatorProfile"); ok {
			if err = d.Set("isolator_profile", vv); err != nil {
				return fmt.Errorf("Error reading isolator_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading isolator_profile: %v", err)
		}
	}

	if err = d.Set("isolator_server", flattenFirewallPolicyIsolatorServer(o["isolator-server"], d, "isolator_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["isolator-server"], "FirewallPolicy-IsolatorServer"); ok {
			if err = d.Set("isolator_server", vv); err != nil {
				return fmt.Errorf("Error reading isolator_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading isolator_server: %v", err)
		}
	}

	if err = d.Set("max_session_per_user", flattenFirewallPolicyMaxSessionPerUser(o["max-session-per-user"], d, "max_session_per_user")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-session-per-user"], "FirewallPolicy-MaxSessionPerUser"); ok {
			if err = d.Set("max_session_per_user", vv); err != nil {
				return fmt.Errorf("Error reading max_session_per_user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_session_per_user: %v", err)
		}
	}

	if err = d.Set("pass_through", flattenFirewallPolicyPassThrough(o["pass-through"], d, "pass_through")); err != nil {
		if vv, ok := fortiAPIPatch(o["pass-through"], "FirewallPolicy-PassThrough"); ok {
			if err = d.Set("pass_through", vv); err != nil {
				return fmt.Errorf("Error reading pass_through: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pass_through: %v", err)
		}
	}

	if err = d.Set("redirect_profile", flattenFirewallPolicyRedirectProfile(o["redirect-profile"], d, "redirect_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["redirect-profile"], "FirewallPolicy-RedirectProfile"); ok {
			if err = d.Set("redirect_profile", vv); err != nil {
				return fmt.Errorf("Error reading redirect_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redirect_profile: %v", err)
		}
	}

	if err = d.Set("reverse_cache", flattenFirewallPolicyReverseCache(o["reverse-cache"], d, "reverse_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["reverse-cache"], "FirewallPolicy-ReverseCache"); ok {
			if err = d.Set("reverse_cache", vv); err != nil {
				return fmt.Errorf("Error reading reverse_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reverse_cache: %v", err)
		}
	}

	if err = d.Set("service_connector", flattenFirewallPolicyServiceConnector(o["service-connector"], d, "service_connector")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-connector"], "FirewallPolicy-ServiceConnector"); ok {
			if err = d.Set("service_connector", vv); err != nil {
				return fmt.Errorf("Error reading service_connector: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_connector: %v", err)
		}
	}

	if err = d.Set("ssh_policy_check", flattenFirewallPolicySshPolicyCheck(o["ssh-policy-check"], d, "ssh_policy_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-policy-check"], "FirewallPolicy-SshPolicyCheck"); ok {
			if err = d.Set("ssh_policy_check", vv); err != nil {
				return fmt.Errorf("Error reading ssh_policy_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_policy_check: %v", err)
		}
	}

	if err = d.Set("transparent", flattenFirewallPolicyTransparent(o["transparent"], d, "transparent")); err != nil {
		if vv, ok := fortiAPIPatch(o["transparent"], "FirewallPolicy-Transparent"); ok {
			if err = d.Set("transparent", vv); err != nil {
				return fmt.Errorf("Error reading transparent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading transparent: %v", err)
		}
	}

	if err = d.Set("type", flattenFirewallPolicyType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "FirewallPolicy-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("url_category", flattenFirewallPolicyUrlCategory(o["url-category"], d, "url_category")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-category"], "FirewallPolicy-UrlCategory"); ok {
			if err = d.Set("url_category", vv); err != nil {
				return fmt.Errorf("Error reading url_category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_category: %v", err)
		}
	}

	if err = d.Set("url_risk", flattenFirewallPolicyUrlRisk(o["url-risk"], d, "url_risk")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-risk"], "FirewallPolicy-UrlRisk"); ok {
			if err = d.Set("url_risk", vv); err != nil {
				return fmt.Errorf("Error reading url_risk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_risk: %v", err)
		}
	}

	if err = d.Set("ztna_proxy", flattenFirewallPolicyZtnaProxy(o["ztna-proxy"], d, "ztna_proxy")); err != nil {
		if vv, ok := fortiAPIPatch(o["ztna-proxy"], "FirewallPolicy-ZtnaProxy"); ok {
			if err = d.Set("ztna_proxy", vv); err != nil {
				return fmt.Errorf("Error reading ztna_proxy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ztna_proxy: %v", err)
		}
	}

	return nil
}

func flattenFirewallPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallPolicyPolicyBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyAntiReplay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyAppMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyApplicationList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyAuthCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyAuthPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyAuthRedirectAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyAutoAsicOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyAvProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyBlockNotification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyCaptivePortalExempt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyCapturePacket(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyCifsProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyCasbProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyCgnEif(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyCgnEim(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyCgnLogServerGrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyCgnResourceQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyCgnSessionQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyCgnSwEifCtrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyCustomLogFields(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyDecryptedTrafficMirror(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyDelayTcpNpuSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDiameterFilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyDiffservCopy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDiffservForward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDiffservReverse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDiffservcodeForward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDiffservcodeRev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDisclaimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDlpProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyDnsfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyDsri(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyDstaddrNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDstaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyDstaddr6Negate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDstintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyDynamicShaping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyEifCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyEifLearn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyEmailCollect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyEmailfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyFec(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyFileFilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyFirewallSessionDirty(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyFixedport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyFssoAgentForNtlm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyFssoGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyGeoipAnycast(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyGeoipMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyGlobalLabel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyGtpProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyHttpPolicyRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyIcapProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyIdentityBasedRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyInbound(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInspectionMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetServiceCustom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyInternetServiceCustomGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyInternetServiceFortiguard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyInternetServiceGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyInternetServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyInternetServiceNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetServiceSrc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetServiceSrcCustom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyInternetServiceSrcCustomGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyInternetServiceSrcFortiguard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyInternetServiceSrcGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyInternetServiceSrcName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyInternetServiceSrcNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetService6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetService6Custom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyInternetService6CustomGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyInternetService6Fortiguard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyInternetService6Group(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyInternetService6Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyInternetService6Negate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetService6Src(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetService6SrcCustom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyInternetService6SrcCustomGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyInternetService6SrcFortiguard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyInternetService6SrcGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyInternetService6SrcName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyInternetService6SrcNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyIpVersionType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyIppool(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyIpsSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyIpsVoipFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyLabel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyLogHttpTransaction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyLogtraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyLogtrafficStart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyMatchVip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyMatchVipOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyNat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyNat46(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyNat64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyNatinbound(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyNatip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandFirewallPolicyNatoutbound(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyNetworkServiceDynamic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyNetworkServiceSrcDynamic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyNpAcceleration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyNtlm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyNtlmEnabledBrowsers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyNtlmGuest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyOutbound(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPassiveWanHealthMeasurement(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPcpInbound(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPcpOutbound(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPcpPoolname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyPerIpShaper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyPermitAnyHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPermitStunHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPfcpProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyPolicyBehaviourType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPolicyExpiry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPolicyExpiryDate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPolicyExpiryDateUtc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPolicyOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPolicyid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPoolname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyPoolname6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyPortPreserve(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPortRandom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyProfileGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyProfileProtocolOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyProfileType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyRadiusIpAuthBypass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyRadiusMacAuthBypass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyRedirectUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyReplacemsgOverrideGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyReputationDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyReputationDirection6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyReputationMinimum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyReputationMinimum6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyRtpAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyRtpNat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySamlServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicySchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyScheduleTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyScim(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyScimGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyScimUsers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicySctpFilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicySendDenyPacket(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyServiceNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySessionTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySgt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicySgtCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySrcVendorMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicySrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicySrcaddrNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySrcaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicySrcaddr6Negate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySrcintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicySshFilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicySshPolicyRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySslSshProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyTcpMssReceiver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyTcpMssSender(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyTcpSessionWithoutSyn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyTcpTimeoutPid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyTelemetryProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyTimeoutSendRst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyTos(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyTosMask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyTosNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyTrafficShaper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyTrafficShaperReverse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyUdpTimeoutPid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyUsers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyUtmStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyVideofilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyVirtualPatchProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyVlanCosFwd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyVlanCosRev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyVlanFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyVoipProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyVpntunnel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyWafProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyWanopt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWanoptDetection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWanoptPassiveOpt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWanoptPeer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyWanoptProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyWccp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWebcache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWebcacheHttps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWebfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyWebproxyForwardServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyWebproxyProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyZtnaDeviceOwnership(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyZtnaEmsTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyZtnaEmsTagNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyZtnaEmsTagSecondary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyZtnaGeoTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyZtnaPolicyRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyZtnaStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyZtnaTagsMatchLogic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyAccessProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyAppCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyAppGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyDetectHttpsInHttpRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDeviceOwnership(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDlpSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyDynamicBypass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyExplicitWebProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyExtendedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyForceProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyHttpTunnelAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyHttpsSubCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyIaProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyImplicitProxyDetection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyIsolatorProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyIsolatorServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyMaxSessionPerUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPassThrough(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyRedirectProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyReverseCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyServiceConnector(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicySshPolicyCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyTransparent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyUrlCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyUrlRisk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyZtnaProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectFirewallPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_policy_block"); ok || d.HasChange("_policy_block") {
		t, err := expandFirewallPolicyPolicyBlock(d, v, "_policy_block")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_policy_block"] = t
		}
	}

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandFirewallPolicyAction(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("anti_replay"); ok || d.HasChange("anti_replay") {
		t, err := expandFirewallPolicyAntiReplay(d, v, "anti_replay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anti-replay"] = t
		}
	}

	if v, ok := d.GetOk("app_monitor"); ok || d.HasChange("app_monitor") {
		t, err := expandFirewallPolicyAppMonitor(d, v, "app_monitor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-monitor"] = t
		}
	}

	if v, ok := d.GetOk("application_list"); ok || d.HasChange("application_list") {
		t, err := expandFirewallPolicyApplicationList(d, v, "application_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-list"] = t
		}
	}

	if v, ok := d.GetOk("auth_cert"); ok || d.HasChange("auth_cert") {
		t, err := expandFirewallPolicyAuthCert(d, v, "auth_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-cert"] = t
		}
	}

	if v, ok := d.GetOk("auth_path"); ok || d.HasChange("auth_path") {
		t, err := expandFirewallPolicyAuthPath(d, v, "auth_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-path"] = t
		}
	}

	if v, ok := d.GetOk("auth_redirect_addr"); ok || d.HasChange("auth_redirect_addr") {
		t, err := expandFirewallPolicyAuthRedirectAddr(d, v, "auth_redirect_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-redirect-addr"] = t
		}
	}

	if v, ok := d.GetOk("auto_asic_offload"); ok || d.HasChange("auto_asic_offload") {
		t, err := expandFirewallPolicyAutoAsicOffload(d, v, "auto_asic_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-asic-offload"] = t
		}
	}

	if v, ok := d.GetOk("av_profile"); ok || d.HasChange("av_profile") {
		t, err := expandFirewallPolicyAvProfile(d, v, "av_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-profile"] = t
		}
	}

	if v, ok := d.GetOk("block_notification"); ok || d.HasChange("block_notification") {
		t, err := expandFirewallPolicyBlockNotification(d, v, "block_notification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-notification"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_exempt"); ok || d.HasChange("captive_portal_exempt") {
		t, err := expandFirewallPolicyCaptivePortalExempt(d, v, "captive_portal_exempt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-exempt"] = t
		}
	}

	if v, ok := d.GetOk("capture_packet"); ok || d.HasChange("capture_packet") {
		t, err := expandFirewallPolicyCapturePacket(d, v, "capture_packet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capture-packet"] = t
		}
	}

	if v, ok := d.GetOk("cifs_profile"); ok || d.HasChange("cifs_profile") {
		t, err := expandFirewallPolicyCifsProfile(d, v, "cifs_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cifs-profile"] = t
		}
	}

	if v, ok := d.GetOk("casb_profile"); ok || d.HasChange("casb_profile") {
		t, err := expandFirewallPolicyCasbProfile(d, v, "casb_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["casb-profile"] = t
		}
	}

	if v, ok := d.GetOk("cgn_eif"); ok || d.HasChange("cgn_eif") {
		t, err := expandFirewallPolicyCgnEif(d, v, "cgn_eif")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-eif"] = t
		}
	}

	if v, ok := d.GetOk("cgn_eim"); ok || d.HasChange("cgn_eim") {
		t, err := expandFirewallPolicyCgnEim(d, v, "cgn_eim")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-eim"] = t
		}
	}

	if v, ok := d.GetOk("cgn_log_server_grp"); ok || d.HasChange("cgn_log_server_grp") {
		t, err := expandFirewallPolicyCgnLogServerGrp(d, v, "cgn_log_server_grp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-log-server-grp"] = t
		}
	}

	if v, ok := d.GetOk("cgn_resource_quota"); ok || d.HasChange("cgn_resource_quota") {
		t, err := expandFirewallPolicyCgnResourceQuota(d, v, "cgn_resource_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-resource-quota"] = t
		}
	}

	if v, ok := d.GetOk("cgn_session_quota"); ok || d.HasChange("cgn_session_quota") {
		t, err := expandFirewallPolicyCgnSessionQuota(d, v, "cgn_session_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-session-quota"] = t
		}
	}

	if v, ok := d.GetOk("cgn_sw_eif_ctrl"); ok || d.HasChange("cgn_sw_eif_ctrl") {
		t, err := expandFirewallPolicyCgnSwEifCtrl(d, v, "cgn_sw_eif_ctrl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-sw-eif-ctrl"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandFirewallPolicyComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("custom_log_fields"); ok || d.HasChange("custom_log_fields") {
		t, err := expandFirewallPolicyCustomLogFields(d, v, "custom_log_fields")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-log-fields"] = t
		}
	}

	if v, ok := d.GetOk("decrypted_traffic_mirror"); ok || d.HasChange("decrypted_traffic_mirror") {
		t, err := expandFirewallPolicyDecryptedTrafficMirror(d, v, "decrypted_traffic_mirror")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["decrypted-traffic-mirror"] = t
		}
	}

	if v, ok := d.GetOk("delay_tcp_npu_session"); ok || d.HasChange("delay_tcp_npu_session") {
		t, err := expandFirewallPolicyDelayTcpNpuSession(d, v, "delay_tcp_npu_session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delay-tcp-npu-session"] = t
		}
	}

	if v, ok := d.GetOk("diameter_filter_profile"); ok || d.HasChange("diameter_filter_profile") {
		t, err := expandFirewallPolicyDiameterFilterProfile(d, v, "diameter_filter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diameter-filter-profile"] = t
		}
	}

	if v, ok := d.GetOk("diffserv_copy"); ok || d.HasChange("diffserv_copy") {
		t, err := expandFirewallPolicyDiffservCopy(d, v, "diffserv_copy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv-copy"] = t
		}
	}

	if v, ok := d.GetOk("diffserv_forward"); ok || d.HasChange("diffserv_forward") {
		t, err := expandFirewallPolicyDiffservForward(d, v, "diffserv_forward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv-forward"] = t
		}
	}

	if v, ok := d.GetOk("diffserv_reverse"); ok || d.HasChange("diffserv_reverse") {
		t, err := expandFirewallPolicyDiffservReverse(d, v, "diffserv_reverse")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv-reverse"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode_forward"); ok || d.HasChange("diffservcode_forward") {
		t, err := expandFirewallPolicyDiffservcodeForward(d, v, "diffservcode_forward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode-forward"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode_rev"); ok || d.HasChange("diffservcode_rev") {
		t, err := expandFirewallPolicyDiffservcodeRev(d, v, "diffservcode_rev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode-rev"] = t
		}
	}

	if v, ok := d.GetOk("disclaimer"); ok || d.HasChange("disclaimer") {
		t, err := expandFirewallPolicyDisclaimer(d, v, "disclaimer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disclaimer"] = t
		}
	}

	if v, ok := d.GetOk("dlp_profile"); ok || d.HasChange("dlp_profile") {
		t, err := expandFirewallPolicyDlpProfile(d, v, "dlp_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-profile"] = t
		}
	}

	if v, ok := d.GetOk("dnsfilter_profile"); ok || d.HasChange("dnsfilter_profile") {
		t, err := expandFirewallPolicyDnsfilterProfile(d, v, "dnsfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dnsfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("dsri"); ok || d.HasChange("dsri") {
		t, err := expandFirewallPolicyDsri(d, v, "dsri")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dsri"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandFirewallPolicyDstaddr(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr_negate"); ok || d.HasChange("dstaddr_negate") {
		t, err := expandFirewallPolicyDstaddrNegate(d, v, "dstaddr_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr-negate"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr6"); ok || d.HasChange("dstaddr6") {
		t, err := expandFirewallPolicyDstaddr6(d, v, "dstaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr6"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr6_negate"); ok || d.HasChange("dstaddr6_negate") {
		t, err := expandFirewallPolicyDstaddr6Negate(d, v, "dstaddr6_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr6-negate"] = t
		}
	}

	if v, ok := d.GetOk("dstintf"); ok || d.HasChange("dstintf") {
		t, err := expandFirewallPolicyDstintf(d, v, "dstintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstintf"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_shaping"); ok || d.HasChange("dynamic_shaping") {
		t, err := expandFirewallPolicyDynamicShaping(d, v, "dynamic_shaping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic-shaping"] = t
		}
	}

	if v, ok := d.GetOk("eif_check"); ok || d.HasChange("eif_check") {
		t, err := expandFirewallPolicyEifCheck(d, v, "eif_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eif-check"] = t
		}
	}

	if v, ok := d.GetOk("eif_learn"); ok || d.HasChange("eif_learn") {
		t, err := expandFirewallPolicyEifLearn(d, v, "eif_learn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eif-learn"] = t
		}
	}

	if v, ok := d.GetOk("email_collect"); ok || d.HasChange("email_collect") {
		t, err := expandFirewallPolicyEmailCollect(d, v, "email_collect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-collect"] = t
		}
	}

	if v, ok := d.GetOk("emailfilter_profile"); ok || d.HasChange("emailfilter_profile") {
		t, err := expandFirewallPolicyEmailfilterProfile(d, v, "emailfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["emailfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("fec"); ok || d.HasChange("fec") {
		t, err := expandFirewallPolicyFec(d, v, "fec")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec"] = t
		}
	}

	if v, ok := d.GetOk("file_filter_profile"); ok || d.HasChange("file_filter_profile") {
		t, err := expandFirewallPolicyFileFilterProfile(d, v, "file_filter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-filter-profile"] = t
		}
	}

	if v, ok := d.GetOk("firewall_session_dirty"); ok || d.HasChange("firewall_session_dirty") {
		t, err := expandFirewallPolicyFirewallSessionDirty(d, v, "firewall_session_dirty")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firewall-session-dirty"] = t
		}
	}

	if v, ok := d.GetOk("fixedport"); ok || d.HasChange("fixedport") {
		t, err := expandFirewallPolicyFixedport(d, v, "fixedport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fixedport"] = t
		}
	}

	if v, ok := d.GetOk("fsso_agent_for_ntlm"); ok || d.HasChange("fsso_agent_for_ntlm") {
		t, err := expandFirewallPolicyFssoAgentForNtlm(d, v, "fsso_agent_for_ntlm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsso-agent-for-ntlm"] = t
		}
	}

	if v, ok := d.GetOk("fsso_groups"); ok || d.HasChange("fsso_groups") {
		t, err := expandFirewallPolicyFssoGroups(d, v, "fsso_groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsso-groups"] = t
		}
	}

	if v, ok := d.GetOk("geoip_anycast"); ok || d.HasChange("geoip_anycast") {
		t, err := expandFirewallPolicyGeoipAnycast(d, v, "geoip_anycast")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["geoip-anycast"] = t
		}
	}

	if v, ok := d.GetOk("geoip_match"); ok || d.HasChange("geoip_match") {
		t, err := expandFirewallPolicyGeoipMatch(d, v, "geoip_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["geoip-match"] = t
		}
	}

	if v, ok := d.GetOk("global_label"); ok || d.HasChange("global_label") {
		t, err := expandFirewallPolicyGlobalLabel(d, v, "global_label")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["global-label"] = t
		}
	}

	if v, ok := d.GetOk("groups"); ok || d.HasChange("groups") {
		t, err := expandFirewallPolicyGroups(d, v, "groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["groups"] = t
		}
	}

	if v, ok := d.GetOk("gtp_profile"); ok || d.HasChange("gtp_profile") {
		t, err := expandFirewallPolicyGtpProfile(d, v, "gtp_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtp-profile"] = t
		}
	}

	if v, ok := d.GetOk("http_policy_redirect"); ok || d.HasChange("http_policy_redirect") {
		t, err := expandFirewallPolicyHttpPolicyRedirect(d, v, "http_policy_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-policy-redirect"] = t
		}
	}

	if v, ok := d.GetOk("icap_profile"); ok || d.HasChange("icap_profile") {
		t, err := expandFirewallPolicyIcapProfile(d, v, "icap_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icap-profile"] = t
		}
	}

	if v, ok := d.GetOk("identity_based_route"); ok || d.HasChange("identity_based_route") {
		t, err := expandFirewallPolicyIdentityBasedRoute(d, v, "identity_based_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["identity-based-route"] = t
		}
	}

	if v, ok := d.GetOk("inbound"); ok || d.HasChange("inbound") {
		t, err := expandFirewallPolicyInbound(d, v, "inbound")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inbound"] = t
		}
	}

	if v, ok := d.GetOk("inspection_mode"); ok || d.HasChange("inspection_mode") {
		t, err := expandFirewallPolicyInspectionMode(d, v, "inspection_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inspection-mode"] = t
		}
	}

	if v, ok := d.GetOk("internet_service"); ok || d.HasChange("internet_service") {
		t, err := expandFirewallPolicyInternetService(d, v, "internet_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom"); ok || d.HasChange("internet_service_custom") {
		t, err := expandFirewallPolicyInternetServiceCustom(d, v, "internet_service_custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom_group"); ok || d.HasChange("internet_service_custom_group") {
		t, err := expandFirewallPolicyInternetServiceCustomGroup(d, v, "internet_service_custom_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_fortiguard"); ok || d.HasChange("internet_service_fortiguard") {
		t, err := expandFirewallPolicyInternetServiceFortiguard(d, v, "internet_service_fortiguard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-fortiguard"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_group"); ok || d.HasChange("internet_service_group") {
		t, err := expandFirewallPolicyInternetServiceGroup(d, v, "internet_service_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_name"); ok || d.HasChange("internet_service_name") {
		t, err := expandFirewallPolicyInternetServiceName(d, v, "internet_service_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-name"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_negate"); ok || d.HasChange("internet_service_negate") {
		t, err := expandFirewallPolicyInternetServiceNegate(d, v, "internet_service_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-negate"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src"); ok || d.HasChange("internet_service_src") {
		t, err := expandFirewallPolicyInternetServiceSrc(d, v, "internet_service_src")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_custom"); ok || d.HasChange("internet_service_src_custom") {
		t, err := expandFirewallPolicyInternetServiceSrcCustom(d, v, "internet_service_src_custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_custom_group"); ok || d.HasChange("internet_service_src_custom_group") {
		t, err := expandFirewallPolicyInternetServiceSrcCustomGroup(d, v, "internet_service_src_custom_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_fortiguard"); ok || d.HasChange("internet_service_src_fortiguard") {
		t, err := expandFirewallPolicyInternetServiceSrcFortiguard(d, v, "internet_service_src_fortiguard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-fortiguard"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_group"); ok || d.HasChange("internet_service_src_group") {
		t, err := expandFirewallPolicyInternetServiceSrcGroup(d, v, "internet_service_src_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_name"); ok || d.HasChange("internet_service_src_name") {
		t, err := expandFirewallPolicyInternetServiceSrcName(d, v, "internet_service_src_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-name"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_negate"); ok || d.HasChange("internet_service_src_negate") {
		t, err := expandFirewallPolicyInternetServiceSrcNegate(d, v, "internet_service_src_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-negate"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6"); ok || d.HasChange("internet_service6") {
		t, err := expandFirewallPolicyInternetService6(d, v, "internet_service6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_custom"); ok || d.HasChange("internet_service6_custom") {
		t, err := expandFirewallPolicyInternetService6Custom(d, v, "internet_service6_custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_custom_group"); ok || d.HasChange("internet_service6_custom_group") {
		t, err := expandFirewallPolicyInternetService6CustomGroup(d, v, "internet_service6_custom_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_fortiguard"); ok || d.HasChange("internet_service6_fortiguard") {
		t, err := expandFirewallPolicyInternetService6Fortiguard(d, v, "internet_service6_fortiguard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-fortiguard"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_group"); ok || d.HasChange("internet_service6_group") {
		t, err := expandFirewallPolicyInternetService6Group(d, v, "internet_service6_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_name"); ok || d.HasChange("internet_service6_name") {
		t, err := expandFirewallPolicyInternetService6Name(d, v, "internet_service6_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-name"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_negate"); ok || d.HasChange("internet_service6_negate") {
		t, err := expandFirewallPolicyInternetService6Negate(d, v, "internet_service6_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-negate"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src"); ok || d.HasChange("internet_service6_src") {
		t, err := expandFirewallPolicyInternetService6Src(d, v, "internet_service6_src")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_custom"); ok || d.HasChange("internet_service6_src_custom") {
		t, err := expandFirewallPolicyInternetService6SrcCustom(d, v, "internet_service6_src_custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_custom_group"); ok || d.HasChange("internet_service6_src_custom_group") {
		t, err := expandFirewallPolicyInternetService6SrcCustomGroup(d, v, "internet_service6_src_custom_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_fortiguard"); ok || d.HasChange("internet_service6_src_fortiguard") {
		t, err := expandFirewallPolicyInternetService6SrcFortiguard(d, v, "internet_service6_src_fortiguard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-fortiguard"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_group"); ok || d.HasChange("internet_service6_src_group") {
		t, err := expandFirewallPolicyInternetService6SrcGroup(d, v, "internet_service6_src_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_name"); ok || d.HasChange("internet_service6_src_name") {
		t, err := expandFirewallPolicyInternetService6SrcName(d, v, "internet_service6_src_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-name"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_negate"); ok || d.HasChange("internet_service6_src_negate") {
		t, err := expandFirewallPolicyInternetService6SrcNegate(d, v, "internet_service6_src_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-negate"] = t
		}
	}

	if v, ok := d.GetOk("ip_version_type"); ok || d.HasChange("ip_version_type") {
		t, err := expandFirewallPolicyIpVersionType(d, v, "ip_version_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-version-type"] = t
		}
	}

	if v, ok := d.GetOk("ippool"); ok || d.HasChange("ippool") {
		t, err := expandFirewallPolicyIppool(d, v, "ippool")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ippool"] = t
		}
	}

	if v, ok := d.GetOk("ips_sensor"); ok || d.HasChange("ips_sensor") {
		t, err := expandFirewallPolicyIpsSensor(d, v, "ips_sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor"] = t
		}
	}

	if v, ok := d.GetOk("ips_voip_filter"); ok || d.HasChange("ips_voip_filter") {
		t, err := expandFirewallPolicyIpsVoipFilter(d, v, "ips_voip_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-voip-filter"] = t
		}
	}

	if v, ok := d.GetOk("label"); ok || d.HasChange("label") {
		t, err := expandFirewallPolicyLabel(d, v, "label")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["label"] = t
		}
	}

	if v, ok := d.GetOk("log_http_transaction"); ok || d.HasChange("log_http_transaction") {
		t, err := expandFirewallPolicyLogHttpTransaction(d, v, "log_http_transaction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-http-transaction"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic"); ok || d.HasChange("logtraffic") {
		t, err := expandFirewallPolicyLogtraffic(d, v, "logtraffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic_start"); ok || d.HasChange("logtraffic_start") {
		t, err := expandFirewallPolicyLogtrafficStart(d, v, "logtraffic_start")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic-start"] = t
		}
	}

	if v, ok := d.GetOk("match_vip"); ok || d.HasChange("match_vip") {
		t, err := expandFirewallPolicyMatchVip(d, v, "match_vip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-vip"] = t
		}
	}

	if v, ok := d.GetOk("match_vip_only"); ok || d.HasChange("match_vip_only") {
		t, err := expandFirewallPolicyMatchVipOnly(d, v, "match_vip_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-vip-only"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nat"); ok || d.HasChange("nat") {
		t, err := expandFirewallPolicyNat(d, v, "nat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat"] = t
		}
	}

	if v, ok := d.GetOk("nat46"); ok || d.HasChange("nat46") {
		t, err := expandFirewallPolicyNat46(d, v, "nat46")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat46"] = t
		}
	}

	if v, ok := d.GetOk("nat64"); ok || d.HasChange("nat64") {
		t, err := expandFirewallPolicyNat64(d, v, "nat64")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat64"] = t
		}
	}

	if v, ok := d.GetOk("natinbound"); ok || d.HasChange("natinbound") {
		t, err := expandFirewallPolicyNatinbound(d, v, "natinbound")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["natinbound"] = t
		}
	}

	if v, ok := d.GetOk("natip"); ok || d.HasChange("natip") {
		t, err := expandFirewallPolicyNatip(d, v, "natip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["natip"] = t
		}
	}

	if v, ok := d.GetOk("natoutbound"); ok || d.HasChange("natoutbound") {
		t, err := expandFirewallPolicyNatoutbound(d, v, "natoutbound")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["natoutbound"] = t
		}
	}

	if v, ok := d.GetOk("network_service_dynamic"); ok || d.HasChange("network_service_dynamic") {
		t, err := expandFirewallPolicyNetworkServiceDynamic(d, v, "network_service_dynamic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-service-dynamic"] = t
		}
	}

	if v, ok := d.GetOk("network_service_src_dynamic"); ok || d.HasChange("network_service_src_dynamic") {
		t, err := expandFirewallPolicyNetworkServiceSrcDynamic(d, v, "network_service_src_dynamic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-service-src-dynamic"] = t
		}
	}

	if v, ok := d.GetOk("np_acceleration"); ok || d.HasChange("np_acceleration") {
		t, err := expandFirewallPolicyNpAcceleration(d, v, "np_acceleration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["np-acceleration"] = t
		}
	}

	if v, ok := d.GetOk("ntlm"); ok || d.HasChange("ntlm") {
		t, err := expandFirewallPolicyNtlm(d, v, "ntlm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntlm"] = t
		}
	}

	if v, ok := d.GetOk("ntlm_enabled_browsers"); ok || d.HasChange("ntlm_enabled_browsers") {
		t, err := expandFirewallPolicyNtlmEnabledBrowsers(d, v, "ntlm_enabled_browsers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntlm-enabled-browsers"] = t
		}
	}

	if v, ok := d.GetOk("ntlm_guest"); ok || d.HasChange("ntlm_guest") {
		t, err := expandFirewallPolicyNtlmGuest(d, v, "ntlm_guest")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntlm-guest"] = t
		}
	}

	if v, ok := d.GetOk("outbound"); ok || d.HasChange("outbound") {
		t, err := expandFirewallPolicyOutbound(d, v, "outbound")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbound"] = t
		}
	}

	if v, ok := d.GetOk("passive_wan_health_measurement"); ok || d.HasChange("passive_wan_health_measurement") {
		t, err := expandFirewallPolicyPassiveWanHealthMeasurement(d, v, "passive_wan_health_measurement")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passive-wan-health-measurement"] = t
		}
	}

	if v, ok := d.GetOk("pcp_inbound"); ok || d.HasChange("pcp_inbound") {
		t, err := expandFirewallPolicyPcpInbound(d, v, "pcp_inbound")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pcp-inbound"] = t
		}
	}

	if v, ok := d.GetOk("pcp_outbound"); ok || d.HasChange("pcp_outbound") {
		t, err := expandFirewallPolicyPcpOutbound(d, v, "pcp_outbound")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pcp-outbound"] = t
		}
	}

	if v, ok := d.GetOk("pcp_poolname"); ok || d.HasChange("pcp_poolname") {
		t, err := expandFirewallPolicyPcpPoolname(d, v, "pcp_poolname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pcp-poolname"] = t
		}
	}

	if v, ok := d.GetOk("per_ip_shaper"); ok || d.HasChange("per_ip_shaper") {
		t, err := expandFirewallPolicyPerIpShaper(d, v, "per_ip_shaper")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-ip-shaper"] = t
		}
	}

	if v, ok := d.GetOk("permit_any_host"); ok || d.HasChange("permit_any_host") {
		t, err := expandFirewallPolicyPermitAnyHost(d, v, "permit_any_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-any-host"] = t
		}
	}

	if v, ok := d.GetOk("permit_stun_host"); ok || d.HasChange("permit_stun_host") {
		t, err := expandFirewallPolicyPermitStunHost(d, v, "permit_stun_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-stun-host"] = t
		}
	}

	if v, ok := d.GetOk("pfcp_profile"); ok || d.HasChange("pfcp_profile") {
		t, err := expandFirewallPolicyPfcpProfile(d, v, "pfcp_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pfcp-profile"] = t
		}
	}

	if v, ok := d.GetOk("policy_behaviour_type"); ok || d.HasChange("policy_behaviour_type") {
		t, err := expandFirewallPolicyPolicyBehaviourType(d, v, "policy_behaviour_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-behaviour-type"] = t
		}
	}

	if v, ok := d.GetOk("policy_expiry"); ok || d.HasChange("policy_expiry") {
		t, err := expandFirewallPolicyPolicyExpiry(d, v, "policy_expiry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-expiry"] = t
		}
	}

	if v, ok := d.GetOk("policy_expiry_date"); ok || d.HasChange("policy_expiry_date") {
		t, err := expandFirewallPolicyPolicyExpiryDate(d, v, "policy_expiry_date")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-expiry-date"] = t
		}
	}

	if v, ok := d.GetOk("policy_expiry_date_utc"); ok || d.HasChange("policy_expiry_date_utc") {
		t, err := expandFirewallPolicyPolicyExpiryDateUtc(d, v, "policy_expiry_date_utc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-expiry-date-utc"] = t
		}
	}

	if v, ok := d.GetOk("policy_offload"); ok || d.HasChange("policy_offload") {
		t, err := expandFirewallPolicyPolicyOffload(d, v, "policy_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-offload"] = t
		}
	}

	if v, ok := d.GetOk("policyid"); ok || d.HasChange("policyid") {
		t, err := expandFirewallPolicyPolicyid(d, v, "policyid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyid"] = t
		}
	}

	if v, ok := d.GetOk("poolname"); ok || d.HasChange("poolname") {
		t, err := expandFirewallPolicyPoolname(d, v, "poolname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poolname"] = t
		}
	}

	if v, ok := d.GetOk("poolname6"); ok || d.HasChange("poolname6") {
		t, err := expandFirewallPolicyPoolname6(d, v, "poolname6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poolname6"] = t
		}
	}

	if v, ok := d.GetOk("port_preserve"); ok || d.HasChange("port_preserve") {
		t, err := expandFirewallPolicyPortPreserve(d, v, "port_preserve")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-preserve"] = t
		}
	}

	if v, ok := d.GetOk("port_random"); ok || d.HasChange("port_random") {
		t, err := expandFirewallPolicyPortRandom(d, v, "port_random")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-random"] = t
		}
	}

	if v, ok := d.GetOk("profile_group"); ok || d.HasChange("profile_group") {
		t, err := expandFirewallPolicyProfileGroup(d, v, "profile_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-group"] = t
		}
	}

	if v, ok := d.GetOk("profile_protocol_options"); ok || d.HasChange("profile_protocol_options") {
		t, err := expandFirewallPolicyProfileProtocolOptions(d, v, "profile_protocol_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-protocol-options"] = t
		}
	}

	if v, ok := d.GetOk("profile_type"); ok || d.HasChange("profile_type") {
		t, err := expandFirewallPolicyProfileType(d, v, "profile_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-type"] = t
		}
	}

	if v, ok := d.GetOk("radius_ip_auth_bypass"); ok || d.HasChange("radius_ip_auth_bypass") {
		t, err := expandFirewallPolicyRadiusIpAuthBypass(d, v, "radius_ip_auth_bypass")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-ip-auth-bypass"] = t
		}
	}

	if v, ok := d.GetOk("radius_mac_auth_bypass"); ok || d.HasChange("radius_mac_auth_bypass") {
		t, err := expandFirewallPolicyRadiusMacAuthBypass(d, v, "radius_mac_auth_bypass")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-auth-bypass"] = t
		}
	}

	if v, ok := d.GetOk("redirect_url"); ok || d.HasChange("redirect_url") {
		t, err := expandFirewallPolicyRedirectUrl(d, v, "redirect_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redirect-url"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_override_group"); ok || d.HasChange("replacemsg_override_group") {
		t, err := expandFirewallPolicyReplacemsgOverrideGroup(d, v, "replacemsg_override_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-override-group"] = t
		}
	}

	if v, ok := d.GetOk("reputation_direction"); ok || d.HasChange("reputation_direction") {
		t, err := expandFirewallPolicyReputationDirection(d, v, "reputation_direction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reputation-direction"] = t
		}
	}

	if v, ok := d.GetOk("reputation_direction6"); ok || d.HasChange("reputation_direction6") {
		t, err := expandFirewallPolicyReputationDirection6(d, v, "reputation_direction6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reputation-direction6"] = t
		}
	}

	if v, ok := d.GetOk("reputation_minimum"); ok || d.HasChange("reputation_minimum") {
		t, err := expandFirewallPolicyReputationMinimum(d, v, "reputation_minimum")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reputation-minimum"] = t
		}
	}

	if v, ok := d.GetOk("reputation_minimum6"); ok || d.HasChange("reputation_minimum6") {
		t, err := expandFirewallPolicyReputationMinimum6(d, v, "reputation_minimum6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reputation-minimum6"] = t
		}
	}

	if v, ok := d.GetOk("rtp_addr"); ok || d.HasChange("rtp_addr") {
		t, err := expandFirewallPolicyRtpAddr(d, v, "rtp_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rtp-addr"] = t
		}
	}

	if v, ok := d.GetOk("rtp_nat"); ok || d.HasChange("rtp_nat") {
		t, err := expandFirewallPolicyRtpNat(d, v, "rtp_nat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rtp-nat"] = t
		}
	}

	if v, ok := d.GetOk("saml_server"); ok || d.HasChange("saml_server") {
		t, err := expandFirewallPolicySamlServer(d, v, "saml_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["saml-server"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok || d.HasChange("schedule") {
		t, err := expandFirewallPolicySchedule(d, v, "schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("schedule_timeout"); ok || d.HasChange("schedule_timeout") {
		t, err := expandFirewallPolicyScheduleTimeout(d, v, "schedule_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule-timeout"] = t
		}
	}

	if v, ok := d.GetOk("scim"); ok || d.HasChange("scim") {
		t, err := expandFirewallPolicyScim(d, v, "scim")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scim"] = t
		}
	}

	if v, ok := d.GetOk("scim_groups"); ok || d.HasChange("scim_groups") {
		t, err := expandFirewallPolicyScimGroups(d, v, "scim_groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scim-groups"] = t
		}
	}

	if v, ok := d.GetOk("scim_users"); ok || d.HasChange("scim_users") {
		t, err := expandFirewallPolicyScimUsers(d, v, "scim_users")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scim-users"] = t
		}
	}

	if v, ok := d.GetOk("sctp_filter_profile"); ok || d.HasChange("sctp_filter_profile") {
		t, err := expandFirewallPolicySctpFilterProfile(d, v, "sctp_filter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sctp-filter-profile"] = t
		}
	}

	if v, ok := d.GetOk("send_deny_packet"); ok || d.HasChange("send_deny_packet") {
		t, err := expandFirewallPolicySendDenyPacket(d, v, "send_deny_packet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-deny-packet"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandFirewallPolicyService(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("service_negate"); ok || d.HasChange("service_negate") {
		t, err := expandFirewallPolicyServiceNegate(d, v, "service_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-negate"] = t
		}
	}

	if v, ok := d.GetOk("session_ttl"); ok || d.HasChange("session_ttl") {
		t, err := expandFirewallPolicySessionTtl(d, v, "session_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-ttl"] = t
		}
	}

	if v, ok := d.GetOk("sgt"); ok || d.HasChange("sgt") {
		t, err := expandFirewallPolicySgt(d, v, "sgt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sgt"] = t
		}
	}

	if v, ok := d.GetOk("sgt_check"); ok || d.HasChange("sgt_check") {
		t, err := expandFirewallPolicySgtCheck(d, v, "sgt_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sgt-check"] = t
		}
	}

	if v, ok := d.GetOk("src_vendor_mac"); ok || d.HasChange("src_vendor_mac") {
		t, err := expandFirewallPolicySrcVendorMac(d, v, "src_vendor_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-vendor-mac"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandFirewallPolicySrcaddr(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr_negate"); ok || d.HasChange("srcaddr_negate") {
		t, err := expandFirewallPolicySrcaddrNegate(d, v, "srcaddr_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr-negate"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr6"); ok || d.HasChange("srcaddr6") {
		t, err := expandFirewallPolicySrcaddr6(d, v, "srcaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr6"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr6_negate"); ok || d.HasChange("srcaddr6_negate") {
		t, err := expandFirewallPolicySrcaddr6Negate(d, v, "srcaddr6_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr6-negate"] = t
		}
	}

	if v, ok := d.GetOk("srcintf"); ok || d.HasChange("srcintf") {
		t, err := expandFirewallPolicySrcintf(d, v, "srcintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf"] = t
		}
	}

	if v, ok := d.GetOk("ssh_filter_profile"); ok || d.HasChange("ssh_filter_profile") {
		t, err := expandFirewallPolicySshFilterProfile(d, v, "ssh_filter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-filter-profile"] = t
		}
	}

	if v, ok := d.GetOk("ssh_policy_redirect"); ok || d.HasChange("ssh_policy_redirect") {
		t, err := expandFirewallPolicySshPolicyRedirect(d, v, "ssh_policy_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-policy-redirect"] = t
		}
	}

	if v, ok := d.GetOk("ssl_ssh_profile"); ok || d.HasChange("ssl_ssh_profile") {
		t, err := expandFirewallPolicySslSshProfile(d, v, "ssl_ssh_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-ssh-profile"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFirewallPolicyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tcp_mss_receiver"); ok || d.HasChange("tcp_mss_receiver") {
		t, err := expandFirewallPolicyTcpMssReceiver(d, v, "tcp_mss_receiver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-mss-receiver"] = t
		}
	}

	if v, ok := d.GetOk("tcp_mss_sender"); ok || d.HasChange("tcp_mss_sender") {
		t, err := expandFirewallPolicyTcpMssSender(d, v, "tcp_mss_sender")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-mss-sender"] = t
		}
	}

	if v, ok := d.GetOk("tcp_session_without_syn"); ok || d.HasChange("tcp_session_without_syn") {
		t, err := expandFirewallPolicyTcpSessionWithoutSyn(d, v, "tcp_session_without_syn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-session-without-syn"] = t
		}
	}

	if v, ok := d.GetOk("tcp_timeout_pid"); ok || d.HasChange("tcp_timeout_pid") {
		t, err := expandFirewallPolicyTcpTimeoutPid(d, v, "tcp_timeout_pid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-timeout-pid"] = t
		}
	}

	if v, ok := d.GetOk("telemetry_profile"); ok || d.HasChange("telemetry_profile") {
		t, err := expandFirewallPolicyTelemetryProfile(d, v, "telemetry_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["telemetry-profile"] = t
		}
	}

	if v, ok := d.GetOk("timeout_send_rst"); ok || d.HasChange("timeout_send_rst") {
		t, err := expandFirewallPolicyTimeoutSendRst(d, v, "timeout_send_rst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout-send-rst"] = t
		}
	}

	if v, ok := d.GetOk("tos"); ok || d.HasChange("tos") {
		t, err := expandFirewallPolicyTos(d, v, "tos")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos"] = t
		}
	}

	if v, ok := d.GetOk("tos_mask"); ok || d.HasChange("tos_mask") {
		t, err := expandFirewallPolicyTosMask(d, v, "tos_mask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos-mask"] = t
		}
	}

	if v, ok := d.GetOk("tos_negate"); ok || d.HasChange("tos_negate") {
		t, err := expandFirewallPolicyTosNegate(d, v, "tos_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos-negate"] = t
		}
	}

	if v, ok := d.GetOk("traffic_shaper"); ok || d.HasChange("traffic_shaper") {
		t, err := expandFirewallPolicyTrafficShaper(d, v, "traffic_shaper")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-shaper"] = t
		}
	}

	if v, ok := d.GetOk("traffic_shaper_reverse"); ok || d.HasChange("traffic_shaper_reverse") {
		t, err := expandFirewallPolicyTrafficShaperReverse(d, v, "traffic_shaper_reverse")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-shaper-reverse"] = t
		}
	}

	if v, ok := d.GetOk("udp_timeout_pid"); ok || d.HasChange("udp_timeout_pid") {
		t, err := expandFirewallPolicyUdpTimeoutPid(d, v, "udp_timeout_pid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-timeout-pid"] = t
		}
	}

	if v, ok := d.GetOk("users"); ok || d.HasChange("users") {
		t, err := expandFirewallPolicyUsers(d, v, "users")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["users"] = t
		}
	}

	if v, ok := d.GetOk("utm_status"); ok || d.HasChange("utm_status") {
		t, err := expandFirewallPolicyUtmStatus(d, v, "utm_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm-status"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandFirewallPolicyUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("videofilter_profile"); ok || d.HasChange("videofilter_profile") {
		t, err := expandFirewallPolicyVideofilterProfile(d, v, "videofilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["videofilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("virtual_patch_profile"); ok || d.HasChange("virtual_patch_profile") {
		t, err := expandFirewallPolicyVirtualPatchProfile(d, v, "virtual_patch_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-patch-profile"] = t
		}
	}

	if v, ok := d.GetOk("vlan_cos_fwd"); ok || d.HasChange("vlan_cos_fwd") {
		t, err := expandFirewallPolicyVlanCosFwd(d, v, "vlan_cos_fwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-cos-fwd"] = t
		}
	}

	if v, ok := d.GetOk("vlan_cos_rev"); ok || d.HasChange("vlan_cos_rev") {
		t, err := expandFirewallPolicyVlanCosRev(d, v, "vlan_cos_rev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-cos-rev"] = t
		}
	}

	if v, ok := d.GetOk("vlan_filter"); ok || d.HasChange("vlan_filter") {
		t, err := expandFirewallPolicyVlanFilter(d, v, "vlan_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-filter"] = t
		}
	}

	if v, ok := d.GetOk("voip_profile"); ok || d.HasChange("voip_profile") {
		t, err := expandFirewallPolicyVoipProfile(d, v, "voip_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voip-profile"] = t
		}
	}

	if v, ok := d.GetOk("vpntunnel"); ok || d.HasChange("vpntunnel") {
		t, err := expandFirewallPolicyVpntunnel(d, v, "vpntunnel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpntunnel"] = t
		}
	}

	if v, ok := d.GetOk("waf_profile"); ok || d.HasChange("waf_profile") {
		t, err := expandFirewallPolicyWafProfile(d, v, "waf_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["waf-profile"] = t
		}
	}

	if v, ok := d.GetOk("wanopt"); ok || d.HasChange("wanopt") {
		t, err := expandFirewallPolicyWanopt(d, v, "wanopt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanopt"] = t
		}
	}

	if v, ok := d.GetOk("wanopt_detection"); ok || d.HasChange("wanopt_detection") {
		t, err := expandFirewallPolicyWanoptDetection(d, v, "wanopt_detection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanopt-detection"] = t
		}
	}

	if v, ok := d.GetOk("wanopt_passive_opt"); ok || d.HasChange("wanopt_passive_opt") {
		t, err := expandFirewallPolicyWanoptPassiveOpt(d, v, "wanopt_passive_opt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanopt-passive-opt"] = t
		}
	}

	if v, ok := d.GetOk("wanopt_peer"); ok || d.HasChange("wanopt_peer") {
		t, err := expandFirewallPolicyWanoptPeer(d, v, "wanopt_peer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanopt-peer"] = t
		}
	}

	if v, ok := d.GetOk("wanopt_profile"); ok || d.HasChange("wanopt_profile") {
		t, err := expandFirewallPolicyWanoptProfile(d, v, "wanopt_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanopt-profile"] = t
		}
	}

	if v, ok := d.GetOk("wccp"); ok || d.HasChange("wccp") {
		t, err := expandFirewallPolicyWccp(d, v, "wccp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wccp"] = t
		}
	}

	if v, ok := d.GetOk("webcache"); ok || d.HasChange("webcache") {
		t, err := expandFirewallPolicyWebcache(d, v, "webcache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webcache"] = t
		}
	}

	if v, ok := d.GetOk("webcache_https"); ok || d.HasChange("webcache_https") {
		t, err := expandFirewallPolicyWebcacheHttps(d, v, "webcache_https")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webcache-https"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_profile"); ok || d.HasChange("webfilter_profile") {
		t, err := expandFirewallPolicyWebfilterProfile(d, v, "webfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("webproxy_forward_server"); ok || d.HasChange("webproxy_forward_server") {
		t, err := expandFirewallPolicyWebproxyForwardServer(d, v, "webproxy_forward_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webproxy-forward-server"] = t
		}
	}

	if v, ok := d.GetOk("webproxy_profile"); ok || d.HasChange("webproxy_profile") {
		t, err := expandFirewallPolicyWebproxyProfile(d, v, "webproxy_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webproxy-profile"] = t
		}
	}

	if v, ok := d.GetOk("ztna_device_ownership"); ok || d.HasChange("ztna_device_ownership") {
		t, err := expandFirewallPolicyZtnaDeviceOwnership(d, v, "ztna_device_ownership")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna-device-ownership"] = t
		}
	}

	if v, ok := d.GetOk("ztna_ems_tag"); ok || d.HasChange("ztna_ems_tag") {
		t, err := expandFirewallPolicyZtnaEmsTag(d, v, "ztna_ems_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna-ems-tag"] = t
		}
	}

	if v, ok := d.GetOk("ztna_ems_tag_negate"); ok || d.HasChange("ztna_ems_tag_negate") {
		t, err := expandFirewallPolicyZtnaEmsTagNegate(d, v, "ztna_ems_tag_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna-ems-tag-negate"] = t
		}
	}

	if v, ok := d.GetOk("ztna_ems_tag_secondary"); ok || d.HasChange("ztna_ems_tag_secondary") {
		t, err := expandFirewallPolicyZtnaEmsTagSecondary(d, v, "ztna_ems_tag_secondary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna-ems-tag-secondary"] = t
		}
	}

	if v, ok := d.GetOk("ztna_geo_tag"); ok || d.HasChange("ztna_geo_tag") {
		t, err := expandFirewallPolicyZtnaGeoTag(d, v, "ztna_geo_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna-geo-tag"] = t
		}
	}

	if v, ok := d.GetOk("ztna_policy_redirect"); ok || d.HasChange("ztna_policy_redirect") {
		t, err := expandFirewallPolicyZtnaPolicyRedirect(d, v, "ztna_policy_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna-policy-redirect"] = t
		}
	}

	if v, ok := d.GetOk("ztna_status"); ok || d.HasChange("ztna_status") {
		t, err := expandFirewallPolicyZtnaStatus(d, v, "ztna_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna-status"] = t
		}
	}

	if v, ok := d.GetOk("ztna_tags_match_logic"); ok || d.HasChange("ztna_tags_match_logic") {
		t, err := expandFirewallPolicyZtnaTagsMatchLogic(d, v, "ztna_tags_match_logic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna-tags-match-logic"] = t
		}
	}

	if v, ok := d.GetOk("access_proxy"); ok || d.HasChange("access_proxy") {
		t, err := expandFirewallPolicyAccessProxy(d, v, "access_proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-proxy"] = t
		}
	}

	if v, ok := d.GetOk("app_category"); ok || d.HasChange("app_category") {
		t, err := expandFirewallPolicyAppCategory(d, v, "app_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-category"] = t
		}
	}

	if v, ok := d.GetOk("app_group"); ok || d.HasChange("app_group") {
		t, err := expandFirewallPolicyAppGroup(d, v, "app_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-group"] = t
		}
	}

	if v, ok := d.GetOk("application"); ok || d.HasChange("application") {
		t, err := expandFirewallPolicyApplication(d, v, "application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("detect_https_in_http_request"); ok || d.HasChange("detect_https_in_http_request") {
		t, err := expandFirewallPolicyDetectHttpsInHttpRequest(d, v, "detect_https_in_http_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["detect-https-in-http-request"] = t
		}
	}

	if v, ok := d.GetOk("device_ownership"); ok || d.HasChange("device_ownership") {
		t, err := expandFirewallPolicyDeviceOwnership(d, v, "device_ownership")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-ownership"] = t
		}
	}

	if v, ok := d.GetOk("dlp_sensor"); ok || d.HasChange("dlp_sensor") {
		t, err := expandFirewallPolicyDlpSensor(d, v, "dlp_sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-sensor"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_bypass"); ok || d.HasChange("dynamic_bypass") {
		t, err := expandFirewallPolicyDynamicBypass(d, v, "dynamic_bypass")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic-bypass"] = t
		}
	}

	if v, ok := d.GetOk("explicit_web_proxy"); ok || d.HasChange("explicit_web_proxy") {
		t, err := expandFirewallPolicyExplicitWebProxy(d, v, "explicit_web_proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["explicit-web-proxy"] = t
		}
	}

	if v, ok := d.GetOk("extended_log"); ok || d.HasChange("extended_log") {
		t, err := expandFirewallPolicyExtendedLog(d, v, "extended_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-log"] = t
		}
	}

	if v, ok := d.GetOk("force_proxy"); ok || d.HasChange("force_proxy") {
		t, err := expandFirewallPolicyForceProxy(d, v, "force_proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["force-proxy"] = t
		}
	}

	if v, ok := d.GetOk("http_tunnel_auth"); ok || d.HasChange("http_tunnel_auth") {
		t, err := expandFirewallPolicyHttpTunnelAuth(d, v, "http_tunnel_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-tunnel-auth"] = t
		}
	}

	if v, ok := d.GetOk("https_sub_category"); ok || d.HasChange("https_sub_category") {
		t, err := expandFirewallPolicyHttpsSubCategory(d, v, "https_sub_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-sub-category"] = t
		}
	}

	if v, ok := d.GetOk("ia_profile"); ok || d.HasChange("ia_profile") {
		t, err := expandFirewallPolicyIaProfile(d, v, "ia_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ia-profile"] = t
		}
	}

	if v, ok := d.GetOk("implicit_proxy_detection"); ok || d.HasChange("implicit_proxy_detection") {
		t, err := expandFirewallPolicyImplicitProxyDetection(d, v, "implicit_proxy_detection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["implicit-proxy-detection"] = t
		}
	}

	if v, ok := d.GetOk("isolator_profile"); ok || d.HasChange("isolator_profile") {
		t, err := expandFirewallPolicyIsolatorProfile(d, v, "isolator_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["isolator-profile"] = t
		}
	}

	if v, ok := d.GetOk("isolator_server"); ok || d.HasChange("isolator_server") {
		t, err := expandFirewallPolicyIsolatorServer(d, v, "isolator_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["isolator-server"] = t
		}
	}

	if v, ok := d.GetOk("max_session_per_user"); ok || d.HasChange("max_session_per_user") {
		t, err := expandFirewallPolicyMaxSessionPerUser(d, v, "max_session_per_user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-session-per-user"] = t
		}
	}

	if v, ok := d.GetOk("pass_through"); ok || d.HasChange("pass_through") {
		t, err := expandFirewallPolicyPassThrough(d, v, "pass_through")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pass-through"] = t
		}
	}

	if v, ok := d.GetOk("redirect_profile"); ok || d.HasChange("redirect_profile") {
		t, err := expandFirewallPolicyRedirectProfile(d, v, "redirect_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redirect-profile"] = t
		}
	}

	if v, ok := d.GetOk("reverse_cache"); ok || d.HasChange("reverse_cache") {
		t, err := expandFirewallPolicyReverseCache(d, v, "reverse_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reverse-cache"] = t
		}
	}

	if v, ok := d.GetOk("service_connector"); ok || d.HasChange("service_connector") {
		t, err := expandFirewallPolicyServiceConnector(d, v, "service_connector")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-connector"] = t
		}
	}

	if v, ok := d.GetOk("ssh_policy_check"); ok || d.HasChange("ssh_policy_check") {
		t, err := expandFirewallPolicySshPolicyCheck(d, v, "ssh_policy_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-policy-check"] = t
		}
	}

	if v, ok := d.GetOk("transparent"); ok || d.HasChange("transparent") {
		t, err := expandFirewallPolicyTransparent(d, v, "transparent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transparent"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandFirewallPolicyType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("url_category"); ok || d.HasChange("url_category") {
		t, err := expandFirewallPolicyUrlCategory(d, v, "url_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-category"] = t
		}
	}

	if v, ok := d.GetOk("url_risk"); ok || d.HasChange("url_risk") {
		t, err := expandFirewallPolicyUrlRisk(d, v, "url_risk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-risk"] = t
		}
	}

	if v, ok := d.GetOk("ztna_proxy"); ok || d.HasChange("ztna_proxy") {
		t, err := expandFirewallPolicyZtnaProxy(d, v, "ztna_proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna-proxy"] = t
		}
	}

	return &obj, nil
}
