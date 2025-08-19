// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure VDOM settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSettingsUpdate,
		Read:   resourceSystemSettingsRead,
		Update: resourceSystemSettingsUpdate,
		Delete: resourceSystemSettingsDelete,

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
			"allow_linkdown_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_subnet_overlap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"application_bandwidth_tracking": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"asymroute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"asymroute_icmp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"asymroute6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"asymroute6_icmp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auxiliary_session": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bfd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bfd_desired_min_tx": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"bfd_detect_mult": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"bfd_dont_enforce_src_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"bfd_required_min_rx": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"block_land_attack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"central_nat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"consolidated_firewall_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"default_app_port_as_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_policy_expiry_days": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"default_voip_alg_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"deny_tcp_with_icmp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"detect_unknown_esp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dhcp_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_proxy_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dhcp_proxy_interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_proxy_vrf_select": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dhcp_server_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dhcp6_server_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"discovered_device_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dp_load_distribution_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dp_load_distribution_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dyn_addr_session_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ecmp_max_paths": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"email_portal_check_dns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ext_resource_session_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"firewall_session_dirty": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fqdn_session_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fw_session_hairpin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gateway6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gtp_asym_fgsp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gtp_monitor_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gui_advanced_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_advanced_wireless_features": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_allow_unnamed_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_antivirus": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_ap_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_application_control": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_casb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_default_policy_columns": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"gui_dhcp_advanced": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_dlp_advanced": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_dlp_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_dns_database": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_dnsfilter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_domain_ip_reputation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gui_dos_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_dynamic_profile_display": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gui_dynamic_device_os_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_dynamic_routing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_email_collection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_endpoint_control": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_endpoint_control_advanced": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_enforce_change_summary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_explicit_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_file_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_fortiap_split_tunneling": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_fortiextender_controller": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_gtp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_icap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_implicit_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_ips": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_load_balance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_local_in_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_local_reports": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_multicast_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_multiple_interface_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_multiple_utm_profiles": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gui_nat46_64": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gui_object_colors": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_per_policy_disclaimer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gui_ot": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_policy_based_ipsec": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_policy_disclaimer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_replacement_message_groups": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gui_proxy_inspection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_route_tag_address_creation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_security_profile_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_spamfilter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_sslvpn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_sslvpn_clients": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_sslvpn_personal_bookmarks": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_sslvpn_realms": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_switch_controller": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_traffic_shaping": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_videofilter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_virtual_patch_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_voip_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_vpn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_waf_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_wan_load_balancing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_wanopt_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_webfilter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_webfilter_advanced": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_wireless_controller": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_ztna": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"h323_direct_model": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_external_dest": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hyperscale_default_policy_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ike_dn_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ike_policy_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ike_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ike_quick_crash_detect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ike_session_resume": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"implicit_allow_dns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ike_tcp_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"internet_service_app_ctrl_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"internet_service_database_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"intree_ses_best_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lan_extension_controller_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"link_down_access": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lldp_reception": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lldp_transmission": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"location_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"manageip": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"manageip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"motherboard_traffic_forwarding": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"multicast_forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multicast_skip_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"multicast_ttl_notchange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat46_force_ipv4_packet_forwarding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat46_generate_ipv6_fragment_header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat64_force_ipv6_packet_forwarding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ngfw_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nonat_eif_key_sel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"npu_group_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"opmode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pfcp_monitor_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"policy_offload_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"prp_trailer_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sccp_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sctp_session_without_init": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ses_denied_multicast_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ses_denied_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session_insert_trial": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sip_expectation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sip_nat_trace": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sip_ssl_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sip_tcp_port": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"sip_udp_port": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"snat_hairpin_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"strict_src_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_session_without_syn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trap_local_session": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"trap_session_flag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"utf8_spam_tagging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"v4_ecmp_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vdom_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vpn_stats_log": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vpn_stats_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"wccp_cache_engine": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemSettingsUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSettings(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemSettings")

	return resourceSystemSettingsRead(d, m)
}

func resourceSystemSettingsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemSettings(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSettingsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSettings resource from API: %v", err)
	}
	return nil
}

func flattenSystemSettingsAllowLinkdownPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsAllowSubnetOverlap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsApplicationBandwidthTracking(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsAsymroute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsAsymrouteIcmp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsAsymroute6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsAsymroute6Icmp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsAuxiliarySession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsBfdDesiredMinTx(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsBfdDetectMult(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsBfdDontEnforceSrcPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsBfdRequiredMinRx(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsBlockLandAttack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsCentralNat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsConsolidatedFirewallMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsDefaultAppPortAsService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsDefaultPolicyExpiryDays(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsDefaultVoipAlgMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsDenyTcpWithIcmp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsDetectUnknownEsp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSettingsDhcpProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsDhcpProxyInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSettingsDhcpProxyInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsDhcpProxyVrfSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsDhcpServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSettingsDhcp6ServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSettingsDiscoveredDeviceTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsDpLoadDistributionGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSettingsDpLoadDistributionMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsDynAddrSessionCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsEcmpMaxPaths(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsEmailPortalCheckDns(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsExtResourceSessionCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsFirewallSessionDirty(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsFqdnSessionCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsFwSessionHairpin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGateway6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGtpAsymFgsp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGtpMonitorMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiAdvancedPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiAdvancedWirelessFeatures(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiAllowUnnamedPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiAntivirus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiApProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiApplicationControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiCasb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiDefaultPolicyColumns(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSettingsGuiDhcpAdvanced(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiDlpAdvanced(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiDlpProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiDnsDatabase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiDnsfilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiDomainIpReputation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiDosPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiDynamicProfileDisplay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiDynamicDeviceOsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiDynamicRouting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiEmailCollection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiEndpointControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiEndpointControlAdvanced(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiEnforceChangeSummary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiExplicitProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiFileFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiFortiapSplitTunneling(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiFortiextenderController(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiGtp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiIcap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiImplicitPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiIps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiLoadBalance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiLocalInPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiLocalReports(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiMulticastPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiMultipleInterfacePolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiMultipleUtmProfiles(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiNat4664(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiObjectColors(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiPerPolicyDisclaimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiOt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiPolicyBasedIpsec(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiPolicyDisclaimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiReplacementMessageGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiProxyInspection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiRouteTagAddressCreation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiSecurityProfileGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiSpamfilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiSslvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiSslvpnClients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiSslvpnPersonalBookmarks(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiSslvpnRealms(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiSwitchController(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiThreatWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiTrafficShaping(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiVideofilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiVirtualPatchProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiVoipProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiVpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiWafProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiWanLoadBalancing(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiWanoptCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiWebfilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiWebfilterAdvanced(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiWirelessController(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsGuiZtna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsH323DirectModel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsHttpExternalDest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsHyperscaleDefaultPolicyAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsIkeDnFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsIkePolicyRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsIkePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsIkeQuickCrashDetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsIkeSessionResume(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsImplicitAllowDns(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsIkeTcpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsInternetServiceAppCtrlSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsInternetServiceDatabaseCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsIntreeSesBestRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSettingsIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsLanExtensionControllerAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsLinkDownAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsLldpReception(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsLldpTransmission(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsLocationId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsMacTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsManageip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSettingsManageip6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsMotherboardTrafficForwarding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSettingsMulticastForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsMulticastSkipPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsMulticastTtlNotchange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsNat46ForceIpv4PacketForwarding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsNat46GenerateIpv6FragmentHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsNat64ForceIpv6PacketForwarding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsNgfwMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsNonatEifKeySel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsNpuGroupId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsOpmode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsPfcpMonitorMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsPolicyOffloadLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsPrpTrailerAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsSccpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsSctpSessionWithoutInit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsSesDeniedMulticastTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsSesDeniedTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsSessionInsertTrial(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsSipExpectation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsSipNatTrace(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsSipSslPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsSipTcpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenSystemSettingsSipUdpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenSystemSettingsSnatHairpinTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsStrictSrcCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsTcpSessionWithoutSyn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsTrapLocalSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsTrapSessionFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsUtf8SpamTagging(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsV4EcmpMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsVdomType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsVpnStatsLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSettingsVpnStatsPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSettingsWccpCacheEngine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("allow_linkdown_path", flattenSystemSettingsAllowLinkdownPath(o["allow-linkdown-path"], d, "allow_linkdown_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["allow-linkdown-path"], "SystemSettings-AllowLinkdownPath"); ok {
			if err = d.Set("allow_linkdown_path", vv); err != nil {
				return fmt.Errorf("Error reading allow_linkdown_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allow_linkdown_path: %v", err)
		}
	}

	if err = d.Set("allow_subnet_overlap", flattenSystemSettingsAllowSubnetOverlap(o["allow-subnet-overlap"], d, "allow_subnet_overlap")); err != nil {
		if vv, ok := fortiAPIPatch(o["allow-subnet-overlap"], "SystemSettings-AllowSubnetOverlap"); ok {
			if err = d.Set("allow_subnet_overlap", vv); err != nil {
				return fmt.Errorf("Error reading allow_subnet_overlap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allow_subnet_overlap: %v", err)
		}
	}

	if err = d.Set("application_bandwidth_tracking", flattenSystemSettingsApplicationBandwidthTracking(o["application-bandwidth-tracking"], d, "application_bandwidth_tracking")); err != nil {
		if vv, ok := fortiAPIPatch(o["application-bandwidth-tracking"], "SystemSettings-ApplicationBandwidthTracking"); ok {
			if err = d.Set("application_bandwidth_tracking", vv); err != nil {
				return fmt.Errorf("Error reading application_bandwidth_tracking: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application_bandwidth_tracking: %v", err)
		}
	}

	if err = d.Set("asymroute", flattenSystemSettingsAsymroute(o["asymroute"], d, "asymroute")); err != nil {
		if vv, ok := fortiAPIPatch(o["asymroute"], "SystemSettings-Asymroute"); ok {
			if err = d.Set("asymroute", vv); err != nil {
				return fmt.Errorf("Error reading asymroute: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading asymroute: %v", err)
		}
	}

	if err = d.Set("asymroute_icmp", flattenSystemSettingsAsymrouteIcmp(o["asymroute-icmp"], d, "asymroute_icmp")); err != nil {
		if vv, ok := fortiAPIPatch(o["asymroute-icmp"], "SystemSettings-AsymrouteIcmp"); ok {
			if err = d.Set("asymroute_icmp", vv); err != nil {
				return fmt.Errorf("Error reading asymroute_icmp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading asymroute_icmp: %v", err)
		}
	}

	if err = d.Set("asymroute6", flattenSystemSettingsAsymroute6(o["asymroute6"], d, "asymroute6")); err != nil {
		if vv, ok := fortiAPIPatch(o["asymroute6"], "SystemSettings-Asymroute6"); ok {
			if err = d.Set("asymroute6", vv); err != nil {
				return fmt.Errorf("Error reading asymroute6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading asymroute6: %v", err)
		}
	}

	if err = d.Set("asymroute6_icmp", flattenSystemSettingsAsymroute6Icmp(o["asymroute6-icmp"], d, "asymroute6_icmp")); err != nil {
		if vv, ok := fortiAPIPatch(o["asymroute6-icmp"], "SystemSettings-Asymroute6Icmp"); ok {
			if err = d.Set("asymroute6_icmp", vv); err != nil {
				return fmt.Errorf("Error reading asymroute6_icmp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading asymroute6_icmp: %v", err)
		}
	}

	if err = d.Set("auxiliary_session", flattenSystemSettingsAuxiliarySession(o["auxiliary-session"], d, "auxiliary_session")); err != nil {
		if vv, ok := fortiAPIPatch(o["auxiliary-session"], "SystemSettings-AuxiliarySession"); ok {
			if err = d.Set("auxiliary_session", vv); err != nil {
				return fmt.Errorf("Error reading auxiliary_session: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auxiliary_session: %v", err)
		}
	}

	if err = d.Set("bfd", flattenSystemSettingsBfd(o["bfd"], d, "bfd")); err != nil {
		if vv, ok := fortiAPIPatch(o["bfd"], "SystemSettings-Bfd"); ok {
			if err = d.Set("bfd", vv); err != nil {
				return fmt.Errorf("Error reading bfd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	if err = d.Set("bfd_desired_min_tx", flattenSystemSettingsBfdDesiredMinTx(o["bfd-desired-min-tx"], d, "bfd_desired_min_tx")); err != nil {
		if vv, ok := fortiAPIPatch(o["bfd-desired-min-tx"], "SystemSettings-BfdDesiredMinTx"); ok {
			if err = d.Set("bfd_desired_min_tx", vv); err != nil {
				return fmt.Errorf("Error reading bfd_desired_min_tx: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bfd_desired_min_tx: %v", err)
		}
	}

	if err = d.Set("bfd_detect_mult", flattenSystemSettingsBfdDetectMult(o["bfd-detect-mult"], d, "bfd_detect_mult")); err != nil {
		if vv, ok := fortiAPIPatch(o["bfd-detect-mult"], "SystemSettings-BfdDetectMult"); ok {
			if err = d.Set("bfd_detect_mult", vv); err != nil {
				return fmt.Errorf("Error reading bfd_detect_mult: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bfd_detect_mult: %v", err)
		}
	}

	if err = d.Set("bfd_dont_enforce_src_port", flattenSystemSettingsBfdDontEnforceSrcPort(o["bfd-dont-enforce-src-port"], d, "bfd_dont_enforce_src_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["bfd-dont-enforce-src-port"], "SystemSettings-BfdDontEnforceSrcPort"); ok {
			if err = d.Set("bfd_dont_enforce_src_port", vv); err != nil {
				return fmt.Errorf("Error reading bfd_dont_enforce_src_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bfd_dont_enforce_src_port: %v", err)
		}
	}

	if err = d.Set("bfd_required_min_rx", flattenSystemSettingsBfdRequiredMinRx(o["bfd-required-min-rx"], d, "bfd_required_min_rx")); err != nil {
		if vv, ok := fortiAPIPatch(o["bfd-required-min-rx"], "SystemSettings-BfdRequiredMinRx"); ok {
			if err = d.Set("bfd_required_min_rx", vv); err != nil {
				return fmt.Errorf("Error reading bfd_required_min_rx: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bfd_required_min_rx: %v", err)
		}
	}

	if err = d.Set("block_land_attack", flattenSystemSettingsBlockLandAttack(o["block-land-attack"], d, "block_land_attack")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-land-attack"], "SystemSettings-BlockLandAttack"); ok {
			if err = d.Set("block_land_attack", vv); err != nil {
				return fmt.Errorf("Error reading block_land_attack: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_land_attack: %v", err)
		}
	}

	if err = d.Set("central_nat", flattenSystemSettingsCentralNat(o["central-nat"], d, "central_nat")); err != nil {
		if vv, ok := fortiAPIPatch(o["central-nat"], "SystemSettings-CentralNat"); ok {
			if err = d.Set("central_nat", vv); err != nil {
				return fmt.Errorf("Error reading central_nat: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading central_nat: %v", err)
		}
	}

	if err = d.Set("comments", flattenSystemSettingsComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "SystemSettings-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("consolidated_firewall_mode", flattenSystemSettingsConsolidatedFirewallMode(o["consolidated-firewall-mode"], d, "consolidated_firewall_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["consolidated-firewall-mode"], "SystemSettings-ConsolidatedFirewallMode"); ok {
			if err = d.Set("consolidated_firewall_mode", vv); err != nil {
				return fmt.Errorf("Error reading consolidated_firewall_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading consolidated_firewall_mode: %v", err)
		}
	}

	if err = d.Set("default_app_port_as_service", flattenSystemSettingsDefaultAppPortAsService(o["default-app-port-as-service"], d, "default_app_port_as_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-app-port-as-service"], "SystemSettings-DefaultAppPortAsService"); ok {
			if err = d.Set("default_app_port_as_service", vv); err != nil {
				return fmt.Errorf("Error reading default_app_port_as_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_app_port_as_service: %v", err)
		}
	}

	if err = d.Set("default_policy_expiry_days", flattenSystemSettingsDefaultPolicyExpiryDays(o["default-policy-expiry-days"], d, "default_policy_expiry_days")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-policy-expiry-days"], "SystemSettings-DefaultPolicyExpiryDays"); ok {
			if err = d.Set("default_policy_expiry_days", vv); err != nil {
				return fmt.Errorf("Error reading default_policy_expiry_days: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_policy_expiry_days: %v", err)
		}
	}

	if err = d.Set("default_voip_alg_mode", flattenSystemSettingsDefaultVoipAlgMode(o["default-voip-alg-mode"], d, "default_voip_alg_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-voip-alg-mode"], "SystemSettings-DefaultVoipAlgMode"); ok {
			if err = d.Set("default_voip_alg_mode", vv); err != nil {
				return fmt.Errorf("Error reading default_voip_alg_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_voip_alg_mode: %v", err)
		}
	}

	if err = d.Set("deny_tcp_with_icmp", flattenSystemSettingsDenyTcpWithIcmp(o["deny-tcp-with-icmp"], d, "deny_tcp_with_icmp")); err != nil {
		if vv, ok := fortiAPIPatch(o["deny-tcp-with-icmp"], "SystemSettings-DenyTcpWithIcmp"); ok {
			if err = d.Set("deny_tcp_with_icmp", vv); err != nil {
				return fmt.Errorf("Error reading deny_tcp_with_icmp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading deny_tcp_with_icmp: %v", err)
		}
	}

	if err = d.Set("detect_unknown_esp", flattenSystemSettingsDetectUnknownEsp(o["detect-unknown-esp"], d, "detect_unknown_esp")); err != nil {
		if vv, ok := fortiAPIPatch(o["detect-unknown-esp"], "SystemSettings-DetectUnknownEsp"); ok {
			if err = d.Set("detect_unknown_esp", vv); err != nil {
				return fmt.Errorf("Error reading detect_unknown_esp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading detect_unknown_esp: %v", err)
		}
	}

	if err = d.Set("device", flattenSystemSettingsDevice(o["device"], d, "device")); err != nil {
		if vv, ok := fortiAPIPatch(o["device"], "SystemSettings-Device"); ok {
			if err = d.Set("device", vv); err != nil {
				return fmt.Errorf("Error reading device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("dhcp_proxy", flattenSystemSettingsDhcpProxy(o["dhcp-proxy"], d, "dhcp_proxy")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-proxy"], "SystemSettings-DhcpProxy"); ok {
			if err = d.Set("dhcp_proxy", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_proxy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_proxy: %v", err)
		}
	}

	if err = d.Set("dhcp_proxy_interface", flattenSystemSettingsDhcpProxyInterface(o["dhcp-proxy-interface"], d, "dhcp_proxy_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-proxy-interface"], "SystemSettings-DhcpProxyInterface"); ok {
			if err = d.Set("dhcp_proxy_interface", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_proxy_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_proxy_interface: %v", err)
		}
	}

	if err = d.Set("dhcp_proxy_interface_select_method", flattenSystemSettingsDhcpProxyInterfaceSelectMethod(o["dhcp-proxy-interface-select-method"], d, "dhcp_proxy_interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-proxy-interface-select-method"], "SystemSettings-DhcpProxyInterfaceSelectMethod"); ok {
			if err = d.Set("dhcp_proxy_interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_proxy_interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_proxy_interface_select_method: %v", err)
		}
	}

	if err = d.Set("dhcp_proxy_vrf_select", flattenSystemSettingsDhcpProxyVrfSelect(o["dhcp-proxy-vrf-select"], d, "dhcp_proxy_vrf_select")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-proxy-vrf-select"], "SystemSettings-DhcpProxyVrfSelect"); ok {
			if err = d.Set("dhcp_proxy_vrf_select", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_proxy_vrf_select: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_proxy_vrf_select: %v", err)
		}
	}

	if err = d.Set("dhcp_server_ip", flattenSystemSettingsDhcpServerIp(o["dhcp-server-ip"], d, "dhcp_server_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-server-ip"], "SystemSettings-DhcpServerIp"); ok {
			if err = d.Set("dhcp_server_ip", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_server_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_server_ip: %v", err)
		}
	}

	if err = d.Set("dhcp6_server_ip", flattenSystemSettingsDhcp6ServerIp(o["dhcp6-server-ip"], d, "dhcp6_server_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp6-server-ip"], "SystemSettings-Dhcp6ServerIp"); ok {
			if err = d.Set("dhcp6_server_ip", vv); err != nil {
				return fmt.Errorf("Error reading dhcp6_server_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp6_server_ip: %v", err)
		}
	}

	if err = d.Set("discovered_device_timeout", flattenSystemSettingsDiscoveredDeviceTimeout(o["discovered-device-timeout"], d, "discovered_device_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["discovered-device-timeout"], "SystemSettings-DiscoveredDeviceTimeout"); ok {
			if err = d.Set("discovered_device_timeout", vv); err != nil {
				return fmt.Errorf("Error reading discovered_device_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading discovered_device_timeout: %v", err)
		}
	}

	if err = d.Set("dp_load_distribution_group", flattenSystemSettingsDpLoadDistributionGroup(o["dp-load-distribution-group"], d, "dp_load_distribution_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["dp-load-distribution-group"], "SystemSettings-DpLoadDistributionGroup"); ok {
			if err = d.Set("dp_load_distribution_group", vv); err != nil {
				return fmt.Errorf("Error reading dp_load_distribution_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dp_load_distribution_group: %v", err)
		}
	}

	if err = d.Set("dp_load_distribution_method", flattenSystemSettingsDpLoadDistributionMethod(o["dp-load-distribution-method"], d, "dp_load_distribution_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["dp-load-distribution-method"], "SystemSettings-DpLoadDistributionMethod"); ok {
			if err = d.Set("dp_load_distribution_method", vv); err != nil {
				return fmt.Errorf("Error reading dp_load_distribution_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dp_load_distribution_method: %v", err)
		}
	}

	if err = d.Set("dyn_addr_session_check", flattenSystemSettingsDynAddrSessionCheck(o["dyn-addr-session-check"], d, "dyn_addr_session_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["dyn-addr-session-check"], "SystemSettings-DynAddrSessionCheck"); ok {
			if err = d.Set("dyn_addr_session_check", vv); err != nil {
				return fmt.Errorf("Error reading dyn_addr_session_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dyn_addr_session_check: %v", err)
		}
	}

	if err = d.Set("ecmp_max_paths", flattenSystemSettingsEcmpMaxPaths(o["ecmp-max-paths"], d, "ecmp_max_paths")); err != nil {
		if vv, ok := fortiAPIPatch(o["ecmp-max-paths"], "SystemSettings-EcmpMaxPaths"); ok {
			if err = d.Set("ecmp_max_paths", vv); err != nil {
				return fmt.Errorf("Error reading ecmp_max_paths: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ecmp_max_paths: %v", err)
		}
	}

	if err = d.Set("email_portal_check_dns", flattenSystemSettingsEmailPortalCheckDns(o["email-portal-check-dns"], d, "email_portal_check_dns")); err != nil {
		if vv, ok := fortiAPIPatch(o["email-portal-check-dns"], "SystemSettings-EmailPortalCheckDns"); ok {
			if err = d.Set("email_portal_check_dns", vv); err != nil {
				return fmt.Errorf("Error reading email_portal_check_dns: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading email_portal_check_dns: %v", err)
		}
	}

	if err = d.Set("ext_resource_session_check", flattenSystemSettingsExtResourceSessionCheck(o["ext-resource-session-check"], d, "ext_resource_session_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["ext-resource-session-check"], "SystemSettings-ExtResourceSessionCheck"); ok {
			if err = d.Set("ext_resource_session_check", vv); err != nil {
				return fmt.Errorf("Error reading ext_resource_session_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ext_resource_session_check: %v", err)
		}
	}

	if err = d.Set("firewall_session_dirty", flattenSystemSettingsFirewallSessionDirty(o["firewall-session-dirty"], d, "firewall_session_dirty")); err != nil {
		if vv, ok := fortiAPIPatch(o["firewall-session-dirty"], "SystemSettings-FirewallSessionDirty"); ok {
			if err = d.Set("firewall_session_dirty", vv); err != nil {
				return fmt.Errorf("Error reading firewall_session_dirty: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading firewall_session_dirty: %v", err)
		}
	}

	if err = d.Set("fqdn_session_check", flattenSystemSettingsFqdnSessionCheck(o["fqdn-session-check"], d, "fqdn_session_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["fqdn-session-check"], "SystemSettings-FqdnSessionCheck"); ok {
			if err = d.Set("fqdn_session_check", vv); err != nil {
				return fmt.Errorf("Error reading fqdn_session_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fqdn_session_check: %v", err)
		}
	}

	if err = d.Set("fw_session_hairpin", flattenSystemSettingsFwSessionHairpin(o["fw-session-hairpin"], d, "fw_session_hairpin")); err != nil {
		if vv, ok := fortiAPIPatch(o["fw-session-hairpin"], "SystemSettings-FwSessionHairpin"); ok {
			if err = d.Set("fw_session_hairpin", vv); err != nil {
				return fmt.Errorf("Error reading fw_session_hairpin: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fw_session_hairpin: %v", err)
		}
	}

	if err = d.Set("gateway", flattenSystemSettingsGateway(o["gateway"], d, "gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["gateway"], "SystemSettings-Gateway"); ok {
			if err = d.Set("gateway", vv); err != nil {
				return fmt.Errorf("Error reading gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("gateway6", flattenSystemSettingsGateway6(o["gateway6"], d, "gateway6")); err != nil {
		if vv, ok := fortiAPIPatch(o["gateway6"], "SystemSettings-Gateway6"); ok {
			if err = d.Set("gateway6", vv); err != nil {
				return fmt.Errorf("Error reading gateway6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gateway6: %v", err)
		}
	}

	if err = d.Set("gtp_asym_fgsp", flattenSystemSettingsGtpAsymFgsp(o["gtp-asym-fgsp"], d, "gtp_asym_fgsp")); err != nil {
		if vv, ok := fortiAPIPatch(o["gtp-asym-fgsp"], "SystemSettings-GtpAsymFgsp"); ok {
			if err = d.Set("gtp_asym_fgsp", vv); err != nil {
				return fmt.Errorf("Error reading gtp_asym_fgsp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gtp_asym_fgsp: %v", err)
		}
	}

	if err = d.Set("gtp_monitor_mode", flattenSystemSettingsGtpMonitorMode(o["gtp-monitor-mode"], d, "gtp_monitor_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["gtp-monitor-mode"], "SystemSettings-GtpMonitorMode"); ok {
			if err = d.Set("gtp_monitor_mode", vv); err != nil {
				return fmt.Errorf("Error reading gtp_monitor_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gtp_monitor_mode: %v", err)
		}
	}

	if err = d.Set("gui_advanced_policy", flattenSystemSettingsGuiAdvancedPolicy(o["gui-advanced-policy"], d, "gui_advanced_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-advanced-policy"], "SystemSettings-GuiAdvancedPolicy"); ok {
			if err = d.Set("gui_advanced_policy", vv); err != nil {
				return fmt.Errorf("Error reading gui_advanced_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_advanced_policy: %v", err)
		}
	}

	if err = d.Set("gui_advanced_wireless_features", flattenSystemSettingsGuiAdvancedWirelessFeatures(o["gui-advanced-wireless-features"], d, "gui_advanced_wireless_features")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-advanced-wireless-features"], "SystemSettings-GuiAdvancedWirelessFeatures"); ok {
			if err = d.Set("gui_advanced_wireless_features", vv); err != nil {
				return fmt.Errorf("Error reading gui_advanced_wireless_features: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_advanced_wireless_features: %v", err)
		}
	}

	if err = d.Set("gui_allow_unnamed_policy", flattenSystemSettingsGuiAllowUnnamedPolicy(o["gui-allow-unnamed-policy"], d, "gui_allow_unnamed_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-allow-unnamed-policy"], "SystemSettings-GuiAllowUnnamedPolicy"); ok {
			if err = d.Set("gui_allow_unnamed_policy", vv); err != nil {
				return fmt.Errorf("Error reading gui_allow_unnamed_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_allow_unnamed_policy: %v", err)
		}
	}

	if err = d.Set("gui_antivirus", flattenSystemSettingsGuiAntivirus(o["gui-antivirus"], d, "gui_antivirus")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-antivirus"], "SystemSettings-GuiAntivirus"); ok {
			if err = d.Set("gui_antivirus", vv); err != nil {
				return fmt.Errorf("Error reading gui_antivirus: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_antivirus: %v", err)
		}
	}

	if err = d.Set("gui_ap_profile", flattenSystemSettingsGuiApProfile(o["gui-ap-profile"], d, "gui_ap_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-ap-profile"], "SystemSettings-GuiApProfile"); ok {
			if err = d.Set("gui_ap_profile", vv); err != nil {
				return fmt.Errorf("Error reading gui_ap_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_ap_profile: %v", err)
		}
	}

	if err = d.Set("gui_application_control", flattenSystemSettingsGuiApplicationControl(o["gui-application-control"], d, "gui_application_control")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-application-control"], "SystemSettings-GuiApplicationControl"); ok {
			if err = d.Set("gui_application_control", vv); err != nil {
				return fmt.Errorf("Error reading gui_application_control: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_application_control: %v", err)
		}
	}

	if err = d.Set("gui_casb", flattenSystemSettingsGuiCasb(o["gui-casb"], d, "gui_casb")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-casb"], "SystemSettings-GuiCasb"); ok {
			if err = d.Set("gui_casb", vv); err != nil {
				return fmt.Errorf("Error reading gui_casb: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_casb: %v", err)
		}
	}

	if err = d.Set("gui_default_policy_columns", flattenSystemSettingsGuiDefaultPolicyColumns(o["gui-default-policy-columns"], d, "gui_default_policy_columns")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-default-policy-columns"], "SystemSettings-GuiDefaultPolicyColumns"); ok {
			if err = d.Set("gui_default_policy_columns", vv); err != nil {
				return fmt.Errorf("Error reading gui_default_policy_columns: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_default_policy_columns: %v", err)
		}
	}

	if err = d.Set("gui_dhcp_advanced", flattenSystemSettingsGuiDhcpAdvanced(o["gui-dhcp-advanced"], d, "gui_dhcp_advanced")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-dhcp-advanced"], "SystemSettings-GuiDhcpAdvanced"); ok {
			if err = d.Set("gui_dhcp_advanced", vv); err != nil {
				return fmt.Errorf("Error reading gui_dhcp_advanced: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_dhcp_advanced: %v", err)
		}
	}

	if err = d.Set("gui_dlp_advanced", flattenSystemSettingsGuiDlpAdvanced(o["gui-dlp-advanced"], d, "gui_dlp_advanced")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-dlp-advanced"], "SystemSettings-GuiDlpAdvanced"); ok {
			if err = d.Set("gui_dlp_advanced", vv); err != nil {
				return fmt.Errorf("Error reading gui_dlp_advanced: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_dlp_advanced: %v", err)
		}
	}

	if err = d.Set("gui_dlp_profile", flattenSystemSettingsGuiDlpProfile(o["gui-dlp-profile"], d, "gui_dlp_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-dlp-profile"], "SystemSettings-GuiDlpProfile"); ok {
			if err = d.Set("gui_dlp_profile", vv); err != nil {
				return fmt.Errorf("Error reading gui_dlp_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_dlp_profile: %v", err)
		}
	}

	if err = d.Set("gui_dns_database", flattenSystemSettingsGuiDnsDatabase(o["gui-dns-database"], d, "gui_dns_database")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-dns-database"], "SystemSettings-GuiDnsDatabase"); ok {
			if err = d.Set("gui_dns_database", vv); err != nil {
				return fmt.Errorf("Error reading gui_dns_database: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_dns_database: %v", err)
		}
	}

	if err = d.Set("gui_dnsfilter", flattenSystemSettingsGuiDnsfilter(o["gui-dnsfilter"], d, "gui_dnsfilter")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-dnsfilter"], "SystemSettings-GuiDnsfilter"); ok {
			if err = d.Set("gui_dnsfilter", vv); err != nil {
				return fmt.Errorf("Error reading gui_dnsfilter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_dnsfilter: %v", err)
		}
	}

	if err = d.Set("gui_domain_ip_reputation", flattenSystemSettingsGuiDomainIpReputation(o["gui-domain-ip-reputation"], d, "gui_domain_ip_reputation")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-domain-ip-reputation"], "SystemSettings-GuiDomainIpReputation"); ok {
			if err = d.Set("gui_domain_ip_reputation", vv); err != nil {
				return fmt.Errorf("Error reading gui_domain_ip_reputation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_domain_ip_reputation: %v", err)
		}
	}

	if err = d.Set("gui_dos_policy", flattenSystemSettingsGuiDosPolicy(o["gui-dos-policy"], d, "gui_dos_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-dos-policy"], "SystemSettings-GuiDosPolicy"); ok {
			if err = d.Set("gui_dos_policy", vv); err != nil {
				return fmt.Errorf("Error reading gui_dos_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_dos_policy: %v", err)
		}
	}

	if err = d.Set("gui_dynamic_profile_display", flattenSystemSettingsGuiDynamicProfileDisplay(o["gui-dynamic-profile-display"], d, "gui_dynamic_profile_display")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-dynamic-profile-display"], "SystemSettings-GuiDynamicProfileDisplay"); ok {
			if err = d.Set("gui_dynamic_profile_display", vv); err != nil {
				return fmt.Errorf("Error reading gui_dynamic_profile_display: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_dynamic_profile_display: %v", err)
		}
	}

	if err = d.Set("gui_dynamic_device_os_id", flattenSystemSettingsGuiDynamicDeviceOsId(o["gui-dynamic-device-os-id"], d, "gui_dynamic_device_os_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-dynamic-device-os-id"], "SystemSettings-GuiDynamicDeviceOsId"); ok {
			if err = d.Set("gui_dynamic_device_os_id", vv); err != nil {
				return fmt.Errorf("Error reading gui_dynamic_device_os_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_dynamic_device_os_id: %v", err)
		}
	}

	if err = d.Set("gui_dynamic_routing", flattenSystemSettingsGuiDynamicRouting(o["gui-dynamic-routing"], d, "gui_dynamic_routing")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-dynamic-routing"], "SystemSettings-GuiDynamicRouting"); ok {
			if err = d.Set("gui_dynamic_routing", vv); err != nil {
				return fmt.Errorf("Error reading gui_dynamic_routing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_dynamic_routing: %v", err)
		}
	}

	if err = d.Set("gui_email_collection", flattenSystemSettingsGuiEmailCollection(o["gui-email-collection"], d, "gui_email_collection")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-email-collection"], "SystemSettings-GuiEmailCollection"); ok {
			if err = d.Set("gui_email_collection", vv); err != nil {
				return fmt.Errorf("Error reading gui_email_collection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_email_collection: %v", err)
		}
	}

	if err = d.Set("gui_endpoint_control", flattenSystemSettingsGuiEndpointControl(o["gui-endpoint-control"], d, "gui_endpoint_control")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-endpoint-control"], "SystemSettings-GuiEndpointControl"); ok {
			if err = d.Set("gui_endpoint_control", vv); err != nil {
				return fmt.Errorf("Error reading gui_endpoint_control: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_endpoint_control: %v", err)
		}
	}

	if err = d.Set("gui_endpoint_control_advanced", flattenSystemSettingsGuiEndpointControlAdvanced(o["gui-endpoint-control-advanced"], d, "gui_endpoint_control_advanced")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-endpoint-control-advanced"], "SystemSettings-GuiEndpointControlAdvanced"); ok {
			if err = d.Set("gui_endpoint_control_advanced", vv); err != nil {
				return fmt.Errorf("Error reading gui_endpoint_control_advanced: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_endpoint_control_advanced: %v", err)
		}
	}

	if err = d.Set("gui_enforce_change_summary", flattenSystemSettingsGuiEnforceChangeSummary(o["gui-enforce-change-summary"], d, "gui_enforce_change_summary")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-enforce-change-summary"], "SystemSettings-GuiEnforceChangeSummary"); ok {
			if err = d.Set("gui_enforce_change_summary", vv); err != nil {
				return fmt.Errorf("Error reading gui_enforce_change_summary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_enforce_change_summary: %v", err)
		}
	}

	if err = d.Set("gui_explicit_proxy", flattenSystemSettingsGuiExplicitProxy(o["gui-explicit-proxy"], d, "gui_explicit_proxy")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-explicit-proxy"], "SystemSettings-GuiExplicitProxy"); ok {
			if err = d.Set("gui_explicit_proxy", vv); err != nil {
				return fmt.Errorf("Error reading gui_explicit_proxy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_explicit_proxy: %v", err)
		}
	}

	if err = d.Set("gui_file_filter", flattenSystemSettingsGuiFileFilter(o["gui-file-filter"], d, "gui_file_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-file-filter"], "SystemSettings-GuiFileFilter"); ok {
			if err = d.Set("gui_file_filter", vv); err != nil {
				return fmt.Errorf("Error reading gui_file_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_file_filter: %v", err)
		}
	}

	if err = d.Set("gui_fortiap_split_tunneling", flattenSystemSettingsGuiFortiapSplitTunneling(o["gui-fortiap-split-tunneling"], d, "gui_fortiap_split_tunneling")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-fortiap-split-tunneling"], "SystemSettings-GuiFortiapSplitTunneling"); ok {
			if err = d.Set("gui_fortiap_split_tunneling", vv); err != nil {
				return fmt.Errorf("Error reading gui_fortiap_split_tunneling: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_fortiap_split_tunneling: %v", err)
		}
	}

	if err = d.Set("gui_fortiextender_controller", flattenSystemSettingsGuiFortiextenderController(o["gui-fortiextender-controller"], d, "gui_fortiextender_controller")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-fortiextender-controller"], "SystemSettings-GuiFortiextenderController"); ok {
			if err = d.Set("gui_fortiextender_controller", vv); err != nil {
				return fmt.Errorf("Error reading gui_fortiextender_controller: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_fortiextender_controller: %v", err)
		}
	}

	if err = d.Set("gui_gtp", flattenSystemSettingsGuiGtp(o["gui-gtp"], d, "gui_gtp")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-gtp"], "SystemSettings-GuiGtp"); ok {
			if err = d.Set("gui_gtp", vv); err != nil {
				return fmt.Errorf("Error reading gui_gtp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_gtp: %v", err)
		}
	}

	if err = d.Set("gui_icap", flattenSystemSettingsGuiIcap(o["gui-icap"], d, "gui_icap")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-icap"], "SystemSettings-GuiIcap"); ok {
			if err = d.Set("gui_icap", vv); err != nil {
				return fmt.Errorf("Error reading gui_icap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_icap: %v", err)
		}
	}

	if err = d.Set("gui_implicit_policy", flattenSystemSettingsGuiImplicitPolicy(o["gui-implicit-policy"], d, "gui_implicit_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-implicit-policy"], "SystemSettings-GuiImplicitPolicy"); ok {
			if err = d.Set("gui_implicit_policy", vv); err != nil {
				return fmt.Errorf("Error reading gui_implicit_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_implicit_policy: %v", err)
		}
	}

	if err = d.Set("gui_ips", flattenSystemSettingsGuiIps(o["gui-ips"], d, "gui_ips")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-ips"], "SystemSettings-GuiIps"); ok {
			if err = d.Set("gui_ips", vv); err != nil {
				return fmt.Errorf("Error reading gui_ips: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_ips: %v", err)
		}
	}

	if err = d.Set("gui_load_balance", flattenSystemSettingsGuiLoadBalance(o["gui-load-balance"], d, "gui_load_balance")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-load-balance"], "SystemSettings-GuiLoadBalance"); ok {
			if err = d.Set("gui_load_balance", vv); err != nil {
				return fmt.Errorf("Error reading gui_load_balance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_load_balance: %v", err)
		}
	}

	if err = d.Set("gui_local_in_policy", flattenSystemSettingsGuiLocalInPolicy(o["gui-local-in-policy"], d, "gui_local_in_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-local-in-policy"], "SystemSettings-GuiLocalInPolicy"); ok {
			if err = d.Set("gui_local_in_policy", vv); err != nil {
				return fmt.Errorf("Error reading gui_local_in_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_local_in_policy: %v", err)
		}
	}

	if err = d.Set("gui_local_reports", flattenSystemSettingsGuiLocalReports(o["gui-local-reports"], d, "gui_local_reports")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-local-reports"], "SystemSettings-GuiLocalReports"); ok {
			if err = d.Set("gui_local_reports", vv); err != nil {
				return fmt.Errorf("Error reading gui_local_reports: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_local_reports: %v", err)
		}
	}

	if err = d.Set("gui_multicast_policy", flattenSystemSettingsGuiMulticastPolicy(o["gui-multicast-policy"], d, "gui_multicast_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-multicast-policy"], "SystemSettings-GuiMulticastPolicy"); ok {
			if err = d.Set("gui_multicast_policy", vv); err != nil {
				return fmt.Errorf("Error reading gui_multicast_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_multicast_policy: %v", err)
		}
	}

	if err = d.Set("gui_multiple_interface_policy", flattenSystemSettingsGuiMultipleInterfacePolicy(o["gui-multiple-interface-policy"], d, "gui_multiple_interface_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-multiple-interface-policy"], "SystemSettings-GuiMultipleInterfacePolicy"); ok {
			if err = d.Set("gui_multiple_interface_policy", vv); err != nil {
				return fmt.Errorf("Error reading gui_multiple_interface_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_multiple_interface_policy: %v", err)
		}
	}

	if err = d.Set("gui_multiple_utm_profiles", flattenSystemSettingsGuiMultipleUtmProfiles(o["gui-multiple-utm-profiles"], d, "gui_multiple_utm_profiles")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-multiple-utm-profiles"], "SystemSettings-GuiMultipleUtmProfiles"); ok {
			if err = d.Set("gui_multiple_utm_profiles", vv); err != nil {
				return fmt.Errorf("Error reading gui_multiple_utm_profiles: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_multiple_utm_profiles: %v", err)
		}
	}

	if err = d.Set("gui_nat46_64", flattenSystemSettingsGuiNat4664(o["gui-nat46-64"], d, "gui_nat46_64")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-nat46-64"], "SystemSettings-GuiNat4664"); ok {
			if err = d.Set("gui_nat46_64", vv); err != nil {
				return fmt.Errorf("Error reading gui_nat46_64: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_nat46_64: %v", err)
		}
	}

	if err = d.Set("gui_object_colors", flattenSystemSettingsGuiObjectColors(o["gui-object-colors"], d, "gui_object_colors")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-object-colors"], "SystemSettings-GuiObjectColors"); ok {
			if err = d.Set("gui_object_colors", vv); err != nil {
				return fmt.Errorf("Error reading gui_object_colors: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_object_colors: %v", err)
		}
	}

	if err = d.Set("gui_per_policy_disclaimer", flattenSystemSettingsGuiPerPolicyDisclaimer(o["gui-per-policy-disclaimer"], d, "gui_per_policy_disclaimer")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-per-policy-disclaimer"], "SystemSettings-GuiPerPolicyDisclaimer"); ok {
			if err = d.Set("gui_per_policy_disclaimer", vv); err != nil {
				return fmt.Errorf("Error reading gui_per_policy_disclaimer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_per_policy_disclaimer: %v", err)
		}
	}

	if err = d.Set("gui_ot", flattenSystemSettingsGuiOt(o["gui-ot"], d, "gui_ot")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-ot"], "SystemSettings-GuiOt"); ok {
			if err = d.Set("gui_ot", vv); err != nil {
				return fmt.Errorf("Error reading gui_ot: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_ot: %v", err)
		}
	}

	if err = d.Set("gui_policy_based_ipsec", flattenSystemSettingsGuiPolicyBasedIpsec(o["gui-policy-based-ipsec"], d, "gui_policy_based_ipsec")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-policy-based-ipsec"], "SystemSettings-GuiPolicyBasedIpsec"); ok {
			if err = d.Set("gui_policy_based_ipsec", vv); err != nil {
				return fmt.Errorf("Error reading gui_policy_based_ipsec: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_policy_based_ipsec: %v", err)
		}
	}

	if err = d.Set("gui_policy_disclaimer", flattenSystemSettingsGuiPolicyDisclaimer(o["gui-policy-disclaimer"], d, "gui_policy_disclaimer")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-policy-disclaimer"], "SystemSettings-GuiPolicyDisclaimer"); ok {
			if err = d.Set("gui_policy_disclaimer", vv); err != nil {
				return fmt.Errorf("Error reading gui_policy_disclaimer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_policy_disclaimer: %v", err)
		}
	}

	if err = d.Set("gui_replacement_message_groups", flattenSystemSettingsGuiReplacementMessageGroups(o["gui-replacement-message-groups"], d, "gui_replacement_message_groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-replacement-message-groups"], "SystemSettings-GuiReplacementMessageGroups"); ok {
			if err = d.Set("gui_replacement_message_groups", vv); err != nil {
				return fmt.Errorf("Error reading gui_replacement_message_groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_replacement_message_groups: %v", err)
		}
	}

	if err = d.Set("gui_proxy_inspection", flattenSystemSettingsGuiProxyInspection(o["gui-proxy-inspection"], d, "gui_proxy_inspection")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-proxy-inspection"], "SystemSettings-GuiProxyInspection"); ok {
			if err = d.Set("gui_proxy_inspection", vv); err != nil {
				return fmt.Errorf("Error reading gui_proxy_inspection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_proxy_inspection: %v", err)
		}
	}

	if err = d.Set("gui_route_tag_address_creation", flattenSystemSettingsGuiRouteTagAddressCreation(o["gui-route-tag-address-creation"], d, "gui_route_tag_address_creation")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-route-tag-address-creation"], "SystemSettings-GuiRouteTagAddressCreation"); ok {
			if err = d.Set("gui_route_tag_address_creation", vv); err != nil {
				return fmt.Errorf("Error reading gui_route_tag_address_creation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_route_tag_address_creation: %v", err)
		}
	}

	if err = d.Set("gui_security_profile_group", flattenSystemSettingsGuiSecurityProfileGroup(o["gui-security-profile-group"], d, "gui_security_profile_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-security-profile-group"], "SystemSettings-GuiSecurityProfileGroup"); ok {
			if err = d.Set("gui_security_profile_group", vv); err != nil {
				return fmt.Errorf("Error reading gui_security_profile_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_security_profile_group: %v", err)
		}
	}

	if err = d.Set("gui_spamfilter", flattenSystemSettingsGuiSpamfilter(o["gui-spamfilter"], d, "gui_spamfilter")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-spamfilter"], "SystemSettings-GuiSpamfilter"); ok {
			if err = d.Set("gui_spamfilter", vv); err != nil {
				return fmt.Errorf("Error reading gui_spamfilter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_spamfilter: %v", err)
		}
	}

	if err = d.Set("gui_sslvpn", flattenSystemSettingsGuiSslvpn(o["gui-sslvpn"], d, "gui_sslvpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-sslvpn"], "SystemSettings-GuiSslvpn"); ok {
			if err = d.Set("gui_sslvpn", vv); err != nil {
				return fmt.Errorf("Error reading gui_sslvpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_sslvpn: %v", err)
		}
	}

	if err = d.Set("gui_sslvpn_clients", flattenSystemSettingsGuiSslvpnClients(o["gui-sslvpn-clients"], d, "gui_sslvpn_clients")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-sslvpn-clients"], "SystemSettings-GuiSslvpnClients"); ok {
			if err = d.Set("gui_sslvpn_clients", vv); err != nil {
				return fmt.Errorf("Error reading gui_sslvpn_clients: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_sslvpn_clients: %v", err)
		}
	}

	if err = d.Set("gui_sslvpn_personal_bookmarks", flattenSystemSettingsGuiSslvpnPersonalBookmarks(o["gui-sslvpn-personal-bookmarks"], d, "gui_sslvpn_personal_bookmarks")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-sslvpn-personal-bookmarks"], "SystemSettings-GuiSslvpnPersonalBookmarks"); ok {
			if err = d.Set("gui_sslvpn_personal_bookmarks", vv); err != nil {
				return fmt.Errorf("Error reading gui_sslvpn_personal_bookmarks: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_sslvpn_personal_bookmarks: %v", err)
		}
	}

	if err = d.Set("gui_sslvpn_realms", flattenSystemSettingsGuiSslvpnRealms(o["gui-sslvpn-realms"], d, "gui_sslvpn_realms")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-sslvpn-realms"], "SystemSettings-GuiSslvpnRealms"); ok {
			if err = d.Set("gui_sslvpn_realms", vv); err != nil {
				return fmt.Errorf("Error reading gui_sslvpn_realms: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_sslvpn_realms: %v", err)
		}
	}

	if err = d.Set("gui_switch_controller", flattenSystemSettingsGuiSwitchController(o["gui-switch-controller"], d, "gui_switch_controller")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-switch-controller"], "SystemSettings-GuiSwitchController"); ok {
			if err = d.Set("gui_switch_controller", vv); err != nil {
				return fmt.Errorf("Error reading gui_switch_controller: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_switch_controller: %v", err)
		}
	}

	if err = d.Set("gui_threat_weight", flattenSystemSettingsGuiThreatWeight(o["gui-threat-weight"], d, "gui_threat_weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-threat-weight"], "SystemSettings-GuiThreatWeight"); ok {
			if err = d.Set("gui_threat_weight", vv); err != nil {
				return fmt.Errorf("Error reading gui_threat_weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_threat_weight: %v", err)
		}
	}

	if err = d.Set("gui_traffic_shaping", flattenSystemSettingsGuiTrafficShaping(o["gui-traffic-shaping"], d, "gui_traffic_shaping")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-traffic-shaping"], "SystemSettings-GuiTrafficShaping"); ok {
			if err = d.Set("gui_traffic_shaping", vv); err != nil {
				return fmt.Errorf("Error reading gui_traffic_shaping: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_traffic_shaping: %v", err)
		}
	}

	if err = d.Set("gui_videofilter", flattenSystemSettingsGuiVideofilter(o["gui-videofilter"], d, "gui_videofilter")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-videofilter"], "SystemSettings-GuiVideofilter"); ok {
			if err = d.Set("gui_videofilter", vv); err != nil {
				return fmt.Errorf("Error reading gui_videofilter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_videofilter: %v", err)
		}
	}

	if err = d.Set("gui_virtual_patch_profile", flattenSystemSettingsGuiVirtualPatchProfile(o["gui-virtual-patch-profile"], d, "gui_virtual_patch_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-virtual-patch-profile"], "SystemSettings-GuiVirtualPatchProfile"); ok {
			if err = d.Set("gui_virtual_patch_profile", vv); err != nil {
				return fmt.Errorf("Error reading gui_virtual_patch_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_virtual_patch_profile: %v", err)
		}
	}

	if err = d.Set("gui_voip_profile", flattenSystemSettingsGuiVoipProfile(o["gui-voip-profile"], d, "gui_voip_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-voip-profile"], "SystemSettings-GuiVoipProfile"); ok {
			if err = d.Set("gui_voip_profile", vv); err != nil {
				return fmt.Errorf("Error reading gui_voip_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_voip_profile: %v", err)
		}
	}

	if err = d.Set("gui_vpn", flattenSystemSettingsGuiVpn(o["gui-vpn"], d, "gui_vpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-vpn"], "SystemSettings-GuiVpn"); ok {
			if err = d.Set("gui_vpn", vv); err != nil {
				return fmt.Errorf("Error reading gui_vpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_vpn: %v", err)
		}
	}

	if err = d.Set("gui_waf_profile", flattenSystemSettingsGuiWafProfile(o["gui-waf-profile"], d, "gui_waf_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-waf-profile"], "SystemSettings-GuiWafProfile"); ok {
			if err = d.Set("gui_waf_profile", vv); err != nil {
				return fmt.Errorf("Error reading gui_waf_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_waf_profile: %v", err)
		}
	}

	if err = d.Set("gui_wan_load_balancing", flattenSystemSettingsGuiWanLoadBalancing(o["gui-wan-load-balancing"], d, "gui_wan_load_balancing")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-wan-load-balancing"], "SystemSettings-GuiWanLoadBalancing"); ok {
			if err = d.Set("gui_wan_load_balancing", vv); err != nil {
				return fmt.Errorf("Error reading gui_wan_load_balancing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_wan_load_balancing: %v", err)
		}
	}

	if err = d.Set("gui_wanopt_cache", flattenSystemSettingsGuiWanoptCache(o["gui-wanopt-cache"], d, "gui_wanopt_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-wanopt-cache"], "SystemSettings-GuiWanoptCache"); ok {
			if err = d.Set("gui_wanopt_cache", vv); err != nil {
				return fmt.Errorf("Error reading gui_wanopt_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_wanopt_cache: %v", err)
		}
	}

	if err = d.Set("gui_webfilter", flattenSystemSettingsGuiWebfilter(o["gui-webfilter"], d, "gui_webfilter")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-webfilter"], "SystemSettings-GuiWebfilter"); ok {
			if err = d.Set("gui_webfilter", vv); err != nil {
				return fmt.Errorf("Error reading gui_webfilter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_webfilter: %v", err)
		}
	}

	if err = d.Set("gui_webfilter_advanced", flattenSystemSettingsGuiWebfilterAdvanced(o["gui-webfilter-advanced"], d, "gui_webfilter_advanced")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-webfilter-advanced"], "SystemSettings-GuiWebfilterAdvanced"); ok {
			if err = d.Set("gui_webfilter_advanced", vv); err != nil {
				return fmt.Errorf("Error reading gui_webfilter_advanced: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_webfilter_advanced: %v", err)
		}
	}

	if err = d.Set("gui_wireless_controller", flattenSystemSettingsGuiWirelessController(o["gui-wireless-controller"], d, "gui_wireless_controller")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-wireless-controller"], "SystemSettings-GuiWirelessController"); ok {
			if err = d.Set("gui_wireless_controller", vv); err != nil {
				return fmt.Errorf("Error reading gui_wireless_controller: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_wireless_controller: %v", err)
		}
	}

	if err = d.Set("gui_ztna", flattenSystemSettingsGuiZtna(o["gui-ztna"], d, "gui_ztna")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-ztna"], "SystemSettings-GuiZtna"); ok {
			if err = d.Set("gui_ztna", vv); err != nil {
				return fmt.Errorf("Error reading gui_ztna: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_ztna: %v", err)
		}
	}

	if err = d.Set("h323_direct_model", flattenSystemSettingsH323DirectModel(o["h323-direct-model"], d, "h323_direct_model")); err != nil {
		if vv, ok := fortiAPIPatch(o["h323-direct-model"], "SystemSettings-H323DirectModel"); ok {
			if err = d.Set("h323_direct_model", vv); err != nil {
				return fmt.Errorf("Error reading h323_direct_model: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading h323_direct_model: %v", err)
		}
	}

	if err = d.Set("http_external_dest", flattenSystemSettingsHttpExternalDest(o["http-external-dest"], d, "http_external_dest")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-external-dest"], "SystemSettings-HttpExternalDest"); ok {
			if err = d.Set("http_external_dest", vv); err != nil {
				return fmt.Errorf("Error reading http_external_dest: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_external_dest: %v", err)
		}
	}

	if err = d.Set("hyperscale_default_policy_action", flattenSystemSettingsHyperscaleDefaultPolicyAction(o["hyperscale-default-policy-action"], d, "hyperscale_default_policy_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["hyperscale-default-policy-action"], "SystemSettings-HyperscaleDefaultPolicyAction"); ok {
			if err = d.Set("hyperscale_default_policy_action", vv); err != nil {
				return fmt.Errorf("Error reading hyperscale_default_policy_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hyperscale_default_policy_action: %v", err)
		}
	}

	if err = d.Set("ike_dn_format", flattenSystemSettingsIkeDnFormat(o["ike-dn-format"], d, "ike_dn_format")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike-dn-format"], "SystemSettings-IkeDnFormat"); ok {
			if err = d.Set("ike_dn_format", vv); err != nil {
				return fmt.Errorf("Error reading ike_dn_format: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike_dn_format: %v", err)
		}
	}

	if err = d.Set("ike_policy_route", flattenSystemSettingsIkePolicyRoute(o["ike-policy-route"], d, "ike_policy_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike-policy-route"], "SystemSettings-IkePolicyRoute"); ok {
			if err = d.Set("ike_policy_route", vv); err != nil {
				return fmt.Errorf("Error reading ike_policy_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike_policy_route: %v", err)
		}
	}

	if err = d.Set("ike_port", flattenSystemSettingsIkePort(o["ike-port"], d, "ike_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike-port"], "SystemSettings-IkePort"); ok {
			if err = d.Set("ike_port", vv); err != nil {
				return fmt.Errorf("Error reading ike_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike_port: %v", err)
		}
	}

	if err = d.Set("ike_quick_crash_detect", flattenSystemSettingsIkeQuickCrashDetect(o["ike-quick-crash-detect"], d, "ike_quick_crash_detect")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike-quick-crash-detect"], "SystemSettings-IkeQuickCrashDetect"); ok {
			if err = d.Set("ike_quick_crash_detect", vv); err != nil {
				return fmt.Errorf("Error reading ike_quick_crash_detect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike_quick_crash_detect: %v", err)
		}
	}

	if err = d.Set("ike_session_resume", flattenSystemSettingsIkeSessionResume(o["ike-session-resume"], d, "ike_session_resume")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike-session-resume"], "SystemSettings-IkeSessionResume"); ok {
			if err = d.Set("ike_session_resume", vv); err != nil {
				return fmt.Errorf("Error reading ike_session_resume: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike_session_resume: %v", err)
		}
	}

	if err = d.Set("implicit_allow_dns", flattenSystemSettingsImplicitAllowDns(o["implicit-allow-dns"], d, "implicit_allow_dns")); err != nil {
		if vv, ok := fortiAPIPatch(o["implicit-allow-dns"], "SystemSettings-ImplicitAllowDns"); ok {
			if err = d.Set("implicit_allow_dns", vv); err != nil {
				return fmt.Errorf("Error reading implicit_allow_dns: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading implicit_allow_dns: %v", err)
		}
	}

	if err = d.Set("ike_tcp_port", flattenSystemSettingsIkeTcpPort(o["ike-tcp-port"], d, "ike_tcp_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike-tcp-port"], "SystemSettings-IkeTcpPort"); ok {
			if err = d.Set("ike_tcp_port", vv); err != nil {
				return fmt.Errorf("Error reading ike_tcp_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike_tcp_port: %v", err)
		}
	}

	if err = d.Set("internet_service_app_ctrl_size", flattenSystemSettingsInternetServiceAppCtrlSize(o["internet-service-app-ctrl-size"], d, "internet_service_app_ctrl_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-app-ctrl-size"], "SystemSettings-InternetServiceAppCtrlSize"); ok {
			if err = d.Set("internet_service_app_ctrl_size", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_app_ctrl_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_app_ctrl_size: %v", err)
		}
	}

	if err = d.Set("internet_service_database_cache", flattenSystemSettingsInternetServiceDatabaseCache(o["internet-service-database-cache"], d, "internet_service_database_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-database-cache"], "SystemSettings-InternetServiceDatabaseCache"); ok {
			if err = d.Set("internet_service_database_cache", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_database_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_database_cache: %v", err)
		}
	}

	if err = d.Set("intree_ses_best_route", flattenSystemSettingsIntreeSesBestRoute(o["intree-ses-best-route"], d, "intree_ses_best_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["intree-ses-best-route"], "SystemSettings-IntreeSesBestRoute"); ok {
			if err = d.Set("intree_ses_best_route", vv); err != nil {
				return fmt.Errorf("Error reading intree_ses_best_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading intree_ses_best_route: %v", err)
		}
	}

	if err = d.Set("ip", flattenSystemSettingsIp(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "SystemSettings-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("ip6", flattenSystemSettingsIp6(o["ip6"], d, "ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6"], "SystemSettings-Ip6"); ok {
			if err = d.Set("ip6", vv); err != nil {
				return fmt.Errorf("Error reading ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("lan_extension_controller_addr", flattenSystemSettingsLanExtensionControllerAddr(o["lan-extension-controller-addr"], d, "lan_extension_controller_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["lan-extension-controller-addr"], "SystemSettings-LanExtensionControllerAddr"); ok {
			if err = d.Set("lan_extension_controller_addr", vv); err != nil {
				return fmt.Errorf("Error reading lan_extension_controller_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lan_extension_controller_addr: %v", err)
		}
	}

	if err = d.Set("link_down_access", flattenSystemSettingsLinkDownAccess(o["link-down-access"], d, "link_down_access")); err != nil {
		if vv, ok := fortiAPIPatch(o["link-down-access"], "SystemSettings-LinkDownAccess"); ok {
			if err = d.Set("link_down_access", vv); err != nil {
				return fmt.Errorf("Error reading link_down_access: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading link_down_access: %v", err)
		}
	}

	if err = d.Set("lldp_reception", flattenSystemSettingsLldpReception(o["lldp-reception"], d, "lldp_reception")); err != nil {
		if vv, ok := fortiAPIPatch(o["lldp-reception"], "SystemSettings-LldpReception"); ok {
			if err = d.Set("lldp_reception", vv); err != nil {
				return fmt.Errorf("Error reading lldp_reception: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lldp_reception: %v", err)
		}
	}

	if err = d.Set("lldp_transmission", flattenSystemSettingsLldpTransmission(o["lldp-transmission"], d, "lldp_transmission")); err != nil {
		if vv, ok := fortiAPIPatch(o["lldp-transmission"], "SystemSettings-LldpTransmission"); ok {
			if err = d.Set("lldp_transmission", vv); err != nil {
				return fmt.Errorf("Error reading lldp_transmission: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lldp_transmission: %v", err)
		}
	}

	if err = d.Set("location_id", flattenSystemSettingsLocationId(o["location-id"], d, "location_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["location-id"], "SystemSettings-LocationId"); ok {
			if err = d.Set("location_id", vv); err != nil {
				return fmt.Errorf("Error reading location_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading location_id: %v", err)
		}
	}

	if err = d.Set("mac_ttl", flattenSystemSettingsMacTtl(o["mac-ttl"], d, "mac_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-ttl"], "SystemSettings-MacTtl"); ok {
			if err = d.Set("mac_ttl", vv); err != nil {
				return fmt.Errorf("Error reading mac_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_ttl: %v", err)
		}
	}

	if err = d.Set("manageip", flattenSystemSettingsManageip(o["manageip"], d, "manageip")); err != nil {
		if vv, ok := fortiAPIPatch(o["manageip"], "SystemSettings-Manageip"); ok {
			if err = d.Set("manageip", vv); err != nil {
				return fmt.Errorf("Error reading manageip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading manageip: %v", err)
		}
	}

	if err = d.Set("manageip6", flattenSystemSettingsManageip6(o["manageip6"], d, "manageip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["manageip6"], "SystemSettings-Manageip6"); ok {
			if err = d.Set("manageip6", vv); err != nil {
				return fmt.Errorf("Error reading manageip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading manageip6: %v", err)
		}
	}

	if err = d.Set("motherboard_traffic_forwarding", flattenSystemSettingsMotherboardTrafficForwarding(o["motherboard-traffic-forwarding"], d, "motherboard_traffic_forwarding")); err != nil {
		if vv, ok := fortiAPIPatch(o["motherboard-traffic-forwarding"], "SystemSettings-MotherboardTrafficForwarding"); ok {
			if err = d.Set("motherboard_traffic_forwarding", vv); err != nil {
				return fmt.Errorf("Error reading motherboard_traffic_forwarding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading motherboard_traffic_forwarding: %v", err)
		}
	}

	if err = d.Set("multicast_forward", flattenSystemSettingsMulticastForward(o["multicast-forward"], d, "multicast_forward")); err != nil {
		if vv, ok := fortiAPIPatch(o["multicast-forward"], "SystemSettings-MulticastForward"); ok {
			if err = d.Set("multicast_forward", vv); err != nil {
				return fmt.Errorf("Error reading multicast_forward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multicast_forward: %v", err)
		}
	}

	if err = d.Set("multicast_skip_policy", flattenSystemSettingsMulticastSkipPolicy(o["multicast-skip-policy"], d, "multicast_skip_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["multicast-skip-policy"], "SystemSettings-MulticastSkipPolicy"); ok {
			if err = d.Set("multicast_skip_policy", vv); err != nil {
				return fmt.Errorf("Error reading multicast_skip_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multicast_skip_policy: %v", err)
		}
	}

	if err = d.Set("multicast_ttl_notchange", flattenSystemSettingsMulticastTtlNotchange(o["multicast-ttl-notchange"], d, "multicast_ttl_notchange")); err != nil {
		if vv, ok := fortiAPIPatch(o["multicast-ttl-notchange"], "SystemSettings-MulticastTtlNotchange"); ok {
			if err = d.Set("multicast_ttl_notchange", vv); err != nil {
				return fmt.Errorf("Error reading multicast_ttl_notchange: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multicast_ttl_notchange: %v", err)
		}
	}

	if err = d.Set("nat46_force_ipv4_packet_forwarding", flattenSystemSettingsNat46ForceIpv4PacketForwarding(o["nat46-force-ipv4-packet-forwarding"], d, "nat46_force_ipv4_packet_forwarding")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat46-force-ipv4-packet-forwarding"], "SystemSettings-Nat46ForceIpv4PacketForwarding"); ok {
			if err = d.Set("nat46_force_ipv4_packet_forwarding", vv); err != nil {
				return fmt.Errorf("Error reading nat46_force_ipv4_packet_forwarding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat46_force_ipv4_packet_forwarding: %v", err)
		}
	}

	if err = d.Set("nat46_generate_ipv6_fragment_header", flattenSystemSettingsNat46GenerateIpv6FragmentHeader(o["nat46-generate-ipv6-fragment-header"], d, "nat46_generate_ipv6_fragment_header")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat46-generate-ipv6-fragment-header"], "SystemSettings-Nat46GenerateIpv6FragmentHeader"); ok {
			if err = d.Set("nat46_generate_ipv6_fragment_header", vv); err != nil {
				return fmt.Errorf("Error reading nat46_generate_ipv6_fragment_header: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat46_generate_ipv6_fragment_header: %v", err)
		}
	}

	if err = d.Set("nat64_force_ipv6_packet_forwarding", flattenSystemSettingsNat64ForceIpv6PacketForwarding(o["nat64-force-ipv6-packet-forwarding"], d, "nat64_force_ipv6_packet_forwarding")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat64-force-ipv6-packet-forwarding"], "SystemSettings-Nat64ForceIpv6PacketForwarding"); ok {
			if err = d.Set("nat64_force_ipv6_packet_forwarding", vv); err != nil {
				return fmt.Errorf("Error reading nat64_force_ipv6_packet_forwarding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat64_force_ipv6_packet_forwarding: %v", err)
		}
	}

	if err = d.Set("ngfw_mode", flattenSystemSettingsNgfwMode(o["ngfw-mode"], d, "ngfw_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ngfw-mode"], "SystemSettings-NgfwMode"); ok {
			if err = d.Set("ngfw_mode", vv); err != nil {
				return fmt.Errorf("Error reading ngfw_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ngfw_mode: %v", err)
		}
	}

	if err = d.Set("nonat_eif_key_sel", flattenSystemSettingsNonatEifKeySel(o["nonat-eif-key-sel"], d, "nonat_eif_key_sel")); err != nil {
		if vv, ok := fortiAPIPatch(o["nonat-eif-key-sel"], "SystemSettings-NonatEifKeySel"); ok {
			if err = d.Set("nonat_eif_key_sel", vv); err != nil {
				return fmt.Errorf("Error reading nonat_eif_key_sel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nonat_eif_key_sel: %v", err)
		}
	}

	if err = d.Set("npu_group_id", flattenSystemSettingsNpuGroupId(o["npu-group-id"], d, "npu_group_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["npu-group-id"], "SystemSettings-NpuGroupId"); ok {
			if err = d.Set("npu_group_id", vv); err != nil {
				return fmt.Errorf("Error reading npu_group_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading npu_group_id: %v", err)
		}
	}

	if err = d.Set("opmode", flattenSystemSettingsOpmode(o["opmode"], d, "opmode")); err != nil {
		if vv, ok := fortiAPIPatch(o["opmode"], "SystemSettings-Opmode"); ok {
			if err = d.Set("opmode", vv); err != nil {
				return fmt.Errorf("Error reading opmode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading opmode: %v", err)
		}
	}

	if err = d.Set("pfcp_monitor_mode", flattenSystemSettingsPfcpMonitorMode(o["pfcp-monitor-mode"], d, "pfcp_monitor_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["pfcp-monitor-mode"], "SystemSettings-PfcpMonitorMode"); ok {
			if err = d.Set("pfcp_monitor_mode", vv); err != nil {
				return fmt.Errorf("Error reading pfcp_monitor_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pfcp_monitor_mode: %v", err)
		}
	}

	if err = d.Set("policy_offload_level", flattenSystemSettingsPolicyOffloadLevel(o["policy-offload-level"], d, "policy_offload_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-offload-level"], "SystemSettings-PolicyOffloadLevel"); ok {
			if err = d.Set("policy_offload_level", vv); err != nil {
				return fmt.Errorf("Error reading policy_offload_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_offload_level: %v", err)
		}
	}

	if err = d.Set("prp_trailer_action", flattenSystemSettingsPrpTrailerAction(o["prp-trailer-action"], d, "prp_trailer_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["prp-trailer-action"], "SystemSettings-PrpTrailerAction"); ok {
			if err = d.Set("prp_trailer_action", vv); err != nil {
				return fmt.Errorf("Error reading prp_trailer_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prp_trailer_action: %v", err)
		}
	}

	if err = d.Set("sccp_port", flattenSystemSettingsSccpPort(o["sccp-port"], d, "sccp_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["sccp-port"], "SystemSettings-SccpPort"); ok {
			if err = d.Set("sccp_port", vv); err != nil {
				return fmt.Errorf("Error reading sccp_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sccp_port: %v", err)
		}
	}

	if err = d.Set("sctp_session_without_init", flattenSystemSettingsSctpSessionWithoutInit(o["sctp-session-without-init"], d, "sctp_session_without_init")); err != nil {
		if vv, ok := fortiAPIPatch(o["sctp-session-without-init"], "SystemSettings-SctpSessionWithoutInit"); ok {
			if err = d.Set("sctp_session_without_init", vv); err != nil {
				return fmt.Errorf("Error reading sctp_session_without_init: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sctp_session_without_init: %v", err)
		}
	}

	if err = d.Set("ses_denied_multicast_traffic", flattenSystemSettingsSesDeniedMulticastTraffic(o["ses-denied-multicast-traffic"], d, "ses_denied_multicast_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["ses-denied-multicast-traffic"], "SystemSettings-SesDeniedMulticastTraffic"); ok {
			if err = d.Set("ses_denied_multicast_traffic", vv); err != nil {
				return fmt.Errorf("Error reading ses_denied_multicast_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ses_denied_multicast_traffic: %v", err)
		}
	}

	if err = d.Set("ses_denied_traffic", flattenSystemSettingsSesDeniedTraffic(o["ses-denied-traffic"], d, "ses_denied_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["ses-denied-traffic"], "SystemSettings-SesDeniedTraffic"); ok {
			if err = d.Set("ses_denied_traffic", vv); err != nil {
				return fmt.Errorf("Error reading ses_denied_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ses_denied_traffic: %v", err)
		}
	}

	if err = d.Set("session_insert_trial", flattenSystemSettingsSessionInsertTrial(o["session-insert-trial"], d, "session_insert_trial")); err != nil {
		if vv, ok := fortiAPIPatch(o["session-insert-trial"], "SystemSettings-SessionInsertTrial"); ok {
			if err = d.Set("session_insert_trial", vv); err != nil {
				return fmt.Errorf("Error reading session_insert_trial: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session_insert_trial: %v", err)
		}
	}

	if err = d.Set("sip_expectation", flattenSystemSettingsSipExpectation(o["sip-expectation"], d, "sip_expectation")); err != nil {
		if vv, ok := fortiAPIPatch(o["sip-expectation"], "SystemSettings-SipExpectation"); ok {
			if err = d.Set("sip_expectation", vv); err != nil {
				return fmt.Errorf("Error reading sip_expectation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sip_expectation: %v", err)
		}
	}

	if err = d.Set("sip_nat_trace", flattenSystemSettingsSipNatTrace(o["sip-nat-trace"], d, "sip_nat_trace")); err != nil {
		if vv, ok := fortiAPIPatch(o["sip-nat-trace"], "SystemSettings-SipNatTrace"); ok {
			if err = d.Set("sip_nat_trace", vv); err != nil {
				return fmt.Errorf("Error reading sip_nat_trace: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sip_nat_trace: %v", err)
		}
	}

	if err = d.Set("sip_ssl_port", flattenSystemSettingsSipSslPort(o["sip-ssl-port"], d, "sip_ssl_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["sip-ssl-port"], "SystemSettings-SipSslPort"); ok {
			if err = d.Set("sip_ssl_port", vv); err != nil {
				return fmt.Errorf("Error reading sip_ssl_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sip_ssl_port: %v", err)
		}
	}

	if err = d.Set("sip_tcp_port", flattenSystemSettingsSipTcpPort(o["sip-tcp-port"], d, "sip_tcp_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["sip-tcp-port"], "SystemSettings-SipTcpPort"); ok {
			if err = d.Set("sip_tcp_port", vv); err != nil {
				return fmt.Errorf("Error reading sip_tcp_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sip_tcp_port: %v", err)
		}
	}

	if err = d.Set("sip_udp_port", flattenSystemSettingsSipUdpPort(o["sip-udp-port"], d, "sip_udp_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["sip-udp-port"], "SystemSettings-SipUdpPort"); ok {
			if err = d.Set("sip_udp_port", vv); err != nil {
				return fmt.Errorf("Error reading sip_udp_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sip_udp_port: %v", err)
		}
	}

	if err = d.Set("snat_hairpin_traffic", flattenSystemSettingsSnatHairpinTraffic(o["snat-hairpin-traffic"], d, "snat_hairpin_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["snat-hairpin-traffic"], "SystemSettings-SnatHairpinTraffic"); ok {
			if err = d.Set("snat_hairpin_traffic", vv); err != nil {
				return fmt.Errorf("Error reading snat_hairpin_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading snat_hairpin_traffic: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemSettingsStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemSettings-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("strict_src_check", flattenSystemSettingsStrictSrcCheck(o["strict-src-check"], d, "strict_src_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["strict-src-check"], "SystemSettings-StrictSrcCheck"); ok {
			if err = d.Set("strict_src_check", vv); err != nil {
				return fmt.Errorf("Error reading strict_src_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading strict_src_check: %v", err)
		}
	}

	if err = d.Set("tcp_session_without_syn", flattenSystemSettingsTcpSessionWithoutSyn(o["tcp-session-without-syn"], d, "tcp_session_without_syn")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-session-without-syn"], "SystemSettings-TcpSessionWithoutSyn"); ok {
			if err = d.Set("tcp_session_without_syn", vv); err != nil {
				return fmt.Errorf("Error reading tcp_session_without_syn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_session_without_syn: %v", err)
		}
	}

	if err = d.Set("trap_local_session", flattenSystemSettingsTrapLocalSession(o["trap-local-session"], d, "trap_local_session")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-local-session"], "SystemSettings-TrapLocalSession"); ok {
			if err = d.Set("trap_local_session", vv); err != nil {
				return fmt.Errorf("Error reading trap_local_session: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_local_session: %v", err)
		}
	}

	if err = d.Set("trap_session_flag", flattenSystemSettingsTrapSessionFlag(o["trap-session-flag"], d, "trap_session_flag")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-session-flag"], "SystemSettings-TrapSessionFlag"); ok {
			if err = d.Set("trap_session_flag", vv); err != nil {
				return fmt.Errorf("Error reading trap_session_flag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_session_flag: %v", err)
		}
	}

	if err = d.Set("utf8_spam_tagging", flattenSystemSettingsUtf8SpamTagging(o["utf8-spam-tagging"], d, "utf8_spam_tagging")); err != nil {
		if vv, ok := fortiAPIPatch(o["utf8-spam-tagging"], "SystemSettings-Utf8SpamTagging"); ok {
			if err = d.Set("utf8_spam_tagging", vv); err != nil {
				return fmt.Errorf("Error reading utf8_spam_tagging: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading utf8_spam_tagging: %v", err)
		}
	}

	if err = d.Set("v4_ecmp_mode", flattenSystemSettingsV4EcmpMode(o["v4-ecmp-mode"], d, "v4_ecmp_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["v4-ecmp-mode"], "SystemSettings-V4EcmpMode"); ok {
			if err = d.Set("v4_ecmp_mode", vv); err != nil {
				return fmt.Errorf("Error reading v4_ecmp_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading v4_ecmp_mode: %v", err)
		}
	}

	if err = d.Set("vdom_type", flattenSystemSettingsVdomType(o["vdom-type"], d, "vdom_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom-type"], "SystemSettings-VdomType"); ok {
			if err = d.Set("vdom_type", vv); err != nil {
				return fmt.Errorf("Error reading vdom_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom_type: %v", err)
		}
	}

	if err = d.Set("vpn_stats_log", flattenSystemSettingsVpnStatsLog(o["vpn-stats-log"], d, "vpn_stats_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["vpn-stats-log"], "SystemSettings-VpnStatsLog"); ok {
			if err = d.Set("vpn_stats_log", vv); err != nil {
				return fmt.Errorf("Error reading vpn_stats_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vpn_stats_log: %v", err)
		}
	}

	if err = d.Set("vpn_stats_period", flattenSystemSettingsVpnStatsPeriod(o["vpn-stats-period"], d, "vpn_stats_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["vpn-stats-period"], "SystemSettings-VpnStatsPeriod"); ok {
			if err = d.Set("vpn_stats_period", vv); err != nil {
				return fmt.Errorf("Error reading vpn_stats_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vpn_stats_period: %v", err)
		}
	}

	if err = d.Set("wccp_cache_engine", flattenSystemSettingsWccpCacheEngine(o["wccp-cache-engine"], d, "wccp_cache_engine")); err != nil {
		if vv, ok := fortiAPIPatch(o["wccp-cache-engine"], "SystemSettings-WccpCacheEngine"); ok {
			if err = d.Set("wccp_cache_engine", vv); err != nil {
				return fmt.Errorf("Error reading wccp_cache_engine: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wccp_cache_engine: %v", err)
		}
	}

	return nil
}

func flattenSystemSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSettingsAllowLinkdownPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsAllowSubnetOverlap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsApplicationBandwidthTracking(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsAsymroute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsAsymrouteIcmp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsAsymroute6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsAsymroute6Icmp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsAuxiliarySession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsBfd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsBfdDesiredMinTx(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsBfdDetectMult(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsBfdDontEnforceSrcPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsBfdRequiredMinRx(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsBlockLandAttack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsCentralNat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsConsolidatedFirewallMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsDefaultAppPortAsService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsDefaultPolicyExpiryDays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsDefaultVoipAlgMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsDenyTcpWithIcmp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsDetectUnknownEsp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSettingsDhcpProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsDhcpProxyInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSettingsDhcpProxyInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsDhcpProxyVrfSelect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsDhcpServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSettingsDhcp6ServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSettingsDiscoveredDeviceTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsDpLoadDistributionGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSettingsDpLoadDistributionMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsDynAddrSessionCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsEcmpMaxPaths(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsEmailPortalCheckDns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsExtResourceSessionCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsFirewallSessionDirty(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsFqdnSessionCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsFwSessionHairpin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGateway6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGtpAsymFgsp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGtpMonitorMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiAdvancedPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiAdvancedWirelessFeatures(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiAllowUnnamedPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiAntivirus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiApProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiApplicationControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiCasb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiDefaultPolicyColumns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSettingsGuiDhcpAdvanced(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiDlpAdvanced(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiDlpProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiDnsDatabase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiDnsfilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiDomainIpReputation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiDosPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiDynamicProfileDisplay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiDynamicDeviceOsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiDynamicRouting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiEmailCollection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiEndpointControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiEndpointControlAdvanced(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiEnforceChangeSummary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiExplicitProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiFileFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiFortiapSplitTunneling(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiFortiextenderController(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiGtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiIcap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiImplicitPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiIps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiLoadBalance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiLocalInPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiLocalReports(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiMulticastPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiMultipleInterfacePolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiMultipleUtmProfiles(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiNat4664(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiObjectColors(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiPerPolicyDisclaimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiOt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiPolicyBasedIpsec(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiPolicyDisclaimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiReplacementMessageGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiProxyInspection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiRouteTagAddressCreation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiSecurityProfileGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiSpamfilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiSslvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiSslvpnClients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiSslvpnPersonalBookmarks(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiSslvpnRealms(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiSwitchController(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiThreatWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiTrafficShaping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiVideofilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiVirtualPatchProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiVoipProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiVpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiWafProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiWanLoadBalancing(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiWanoptCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiWebfilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiWebfilterAdvanced(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiWirelessController(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiZtna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsH323DirectModel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsHttpExternalDest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsHyperscaleDefaultPolicyAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsIkeDnFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsIkePolicyRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsIkePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsIkeQuickCrashDetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsIkeSessionResume(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsImplicitAllowDns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsIkeTcpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsInternetServiceAppCtrlSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsInternetServiceDatabaseCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsIntreeSesBestRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemSettingsIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsLanExtensionControllerAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsLinkDownAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsLldpReception(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsLldpTransmission(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsLocationId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsMacTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsManageip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSettingsManageip6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsMotherboardTrafficForwarding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSettingsMulticastForward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsMulticastSkipPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsMulticastTtlNotchange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsNat46ForceIpv4PacketForwarding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsNat46GenerateIpv6FragmentHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsNat64ForceIpv6PacketForwarding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsNgfwMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsNonatEifKeySel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsNpuGroupId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsOpmode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsPfcpMonitorMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsPolicyOffloadLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsPrpTrailerAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsSccpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsSctpSessionWithoutInit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsSesDeniedMulticastTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsSesDeniedTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsSessionInsertTrial(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsSipExpectation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsSipNatTrace(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsSipSslPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsSipTcpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandSystemSettingsSipUdpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandSystemSettingsSnatHairpinTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsStrictSrcCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsTcpSessionWithoutSyn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsTrapLocalSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsTrapSessionFlag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsUtf8SpamTagging(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsV4EcmpMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsVdomType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsVpnStatsLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSettingsVpnStatsPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsWccpCacheEngine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("allow_linkdown_path"); ok || d.HasChange("allow_linkdown_path") {
		t, err := expandSystemSettingsAllowLinkdownPath(d, v, "allow_linkdown_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-linkdown-path"] = t
		}
	}

	if v, ok := d.GetOk("allow_subnet_overlap"); ok || d.HasChange("allow_subnet_overlap") {
		t, err := expandSystemSettingsAllowSubnetOverlap(d, v, "allow_subnet_overlap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-subnet-overlap"] = t
		}
	}

	if v, ok := d.GetOk("application_bandwidth_tracking"); ok || d.HasChange("application_bandwidth_tracking") {
		t, err := expandSystemSettingsApplicationBandwidthTracking(d, v, "application_bandwidth_tracking")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-bandwidth-tracking"] = t
		}
	}

	if v, ok := d.GetOk("asymroute"); ok || d.HasChange("asymroute") {
		t, err := expandSystemSettingsAsymroute(d, v, "asymroute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["asymroute"] = t
		}
	}

	if v, ok := d.GetOk("asymroute_icmp"); ok || d.HasChange("asymroute_icmp") {
		t, err := expandSystemSettingsAsymrouteIcmp(d, v, "asymroute_icmp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["asymroute-icmp"] = t
		}
	}

	if v, ok := d.GetOk("asymroute6"); ok || d.HasChange("asymroute6") {
		t, err := expandSystemSettingsAsymroute6(d, v, "asymroute6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["asymroute6"] = t
		}
	}

	if v, ok := d.GetOk("asymroute6_icmp"); ok || d.HasChange("asymroute6_icmp") {
		t, err := expandSystemSettingsAsymroute6Icmp(d, v, "asymroute6_icmp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["asymroute6-icmp"] = t
		}
	}

	if v, ok := d.GetOk("auxiliary_session"); ok || d.HasChange("auxiliary_session") {
		t, err := expandSystemSettingsAuxiliarySession(d, v, "auxiliary_session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auxiliary-session"] = t
		}
	}

	if v, ok := d.GetOk("bfd"); ok || d.HasChange("bfd") {
		t, err := expandSystemSettingsBfd(d, v, "bfd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd"] = t
		}
	}

	if v, ok := d.GetOk("bfd_desired_min_tx"); ok || d.HasChange("bfd_desired_min_tx") {
		t, err := expandSystemSettingsBfdDesiredMinTx(d, v, "bfd_desired_min_tx")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd-desired-min-tx"] = t
		}
	}

	if v, ok := d.GetOk("bfd_detect_mult"); ok || d.HasChange("bfd_detect_mult") {
		t, err := expandSystemSettingsBfdDetectMult(d, v, "bfd_detect_mult")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd-detect-mult"] = t
		}
	}

	if v, ok := d.GetOk("bfd_dont_enforce_src_port"); ok || d.HasChange("bfd_dont_enforce_src_port") {
		t, err := expandSystemSettingsBfdDontEnforceSrcPort(d, v, "bfd_dont_enforce_src_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd-dont-enforce-src-port"] = t
		}
	}

	if v, ok := d.GetOk("bfd_required_min_rx"); ok || d.HasChange("bfd_required_min_rx") {
		t, err := expandSystemSettingsBfdRequiredMinRx(d, v, "bfd_required_min_rx")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd-required-min-rx"] = t
		}
	}

	if v, ok := d.GetOk("block_land_attack"); ok || d.HasChange("block_land_attack") {
		t, err := expandSystemSettingsBlockLandAttack(d, v, "block_land_attack")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-land-attack"] = t
		}
	}

	if v, ok := d.GetOk("central_nat"); ok || d.HasChange("central_nat") {
		t, err := expandSystemSettingsCentralNat(d, v, "central_nat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["central-nat"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandSystemSettingsComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("consolidated_firewall_mode"); ok || d.HasChange("consolidated_firewall_mode") {
		t, err := expandSystemSettingsConsolidatedFirewallMode(d, v, "consolidated_firewall_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["consolidated-firewall-mode"] = t
		}
	}

	if v, ok := d.GetOk("default_app_port_as_service"); ok || d.HasChange("default_app_port_as_service") {
		t, err := expandSystemSettingsDefaultAppPortAsService(d, v, "default_app_port_as_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-app-port-as-service"] = t
		}
	}

	if v, ok := d.GetOk("default_policy_expiry_days"); ok || d.HasChange("default_policy_expiry_days") {
		t, err := expandSystemSettingsDefaultPolicyExpiryDays(d, v, "default_policy_expiry_days")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-policy-expiry-days"] = t
		}
	}

	if v, ok := d.GetOk("default_voip_alg_mode"); ok || d.HasChange("default_voip_alg_mode") {
		t, err := expandSystemSettingsDefaultVoipAlgMode(d, v, "default_voip_alg_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-voip-alg-mode"] = t
		}
	}

	if v, ok := d.GetOk("deny_tcp_with_icmp"); ok || d.HasChange("deny_tcp_with_icmp") {
		t, err := expandSystemSettingsDenyTcpWithIcmp(d, v, "deny_tcp_with_icmp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deny-tcp-with-icmp"] = t
		}
	}

	if v, ok := d.GetOk("detect_unknown_esp"); ok || d.HasChange("detect_unknown_esp") {
		t, err := expandSystemSettingsDetectUnknownEsp(d, v, "detect_unknown_esp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["detect-unknown-esp"] = t
		}
	}

	if v, ok := d.GetOk("device"); ok || d.HasChange("device") {
		t, err := expandSystemSettingsDevice(d, v, "device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_proxy"); ok || d.HasChange("dhcp_proxy") {
		t, err := expandSystemSettingsDhcpProxy(d, v, "dhcp_proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-proxy"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_proxy_interface"); ok || d.HasChange("dhcp_proxy_interface") {
		t, err := expandSystemSettingsDhcpProxyInterface(d, v, "dhcp_proxy_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-proxy-interface"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_proxy_interface_select_method"); ok || d.HasChange("dhcp_proxy_interface_select_method") {
		t, err := expandSystemSettingsDhcpProxyInterfaceSelectMethod(d, v, "dhcp_proxy_interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-proxy-interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_proxy_vrf_select"); ok || d.HasChange("dhcp_proxy_vrf_select") {
		t, err := expandSystemSettingsDhcpProxyVrfSelect(d, v, "dhcp_proxy_vrf_select")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-proxy-vrf-select"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_server_ip"); ok || d.HasChange("dhcp_server_ip") {
		t, err := expandSystemSettingsDhcpServerIp(d, v, "dhcp_server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_server_ip"); ok || d.HasChange("dhcp6_server_ip") {
		t, err := expandSystemSettingsDhcp6ServerIp(d, v, "dhcp6_server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("discovered_device_timeout"); ok || d.HasChange("discovered_device_timeout") {
		t, err := expandSystemSettingsDiscoveredDeviceTimeout(d, v, "discovered_device_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["discovered-device-timeout"] = t
		}
	}

	if v, ok := d.GetOk("dp_load_distribution_group"); ok || d.HasChange("dp_load_distribution_group") {
		t, err := expandSystemSettingsDpLoadDistributionGroup(d, v, "dp_load_distribution_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dp-load-distribution-group"] = t
		}
	}

	if v, ok := d.GetOk("dp_load_distribution_method"); ok || d.HasChange("dp_load_distribution_method") {
		t, err := expandSystemSettingsDpLoadDistributionMethod(d, v, "dp_load_distribution_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dp-load-distribution-method"] = t
		}
	}

	if v, ok := d.GetOk("dyn_addr_session_check"); ok || d.HasChange("dyn_addr_session_check") {
		t, err := expandSystemSettingsDynAddrSessionCheck(d, v, "dyn_addr_session_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dyn-addr-session-check"] = t
		}
	}

	if v, ok := d.GetOk("ecmp_max_paths"); ok || d.HasChange("ecmp_max_paths") {
		t, err := expandSystemSettingsEcmpMaxPaths(d, v, "ecmp_max_paths")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ecmp-max-paths"] = t
		}
	}

	if v, ok := d.GetOk("email_portal_check_dns"); ok || d.HasChange("email_portal_check_dns") {
		t, err := expandSystemSettingsEmailPortalCheckDns(d, v, "email_portal_check_dns")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-portal-check-dns"] = t
		}
	}

	if v, ok := d.GetOk("ext_resource_session_check"); ok || d.HasChange("ext_resource_session_check") {
		t, err := expandSystemSettingsExtResourceSessionCheck(d, v, "ext_resource_session_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ext-resource-session-check"] = t
		}
	}

	if v, ok := d.GetOk("firewall_session_dirty"); ok || d.HasChange("firewall_session_dirty") {
		t, err := expandSystemSettingsFirewallSessionDirty(d, v, "firewall_session_dirty")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firewall-session-dirty"] = t
		}
	}

	if v, ok := d.GetOk("fqdn_session_check"); ok || d.HasChange("fqdn_session_check") {
		t, err := expandSystemSettingsFqdnSessionCheck(d, v, "fqdn_session_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fqdn-session-check"] = t
		}
	}

	if v, ok := d.GetOk("fw_session_hairpin"); ok || d.HasChange("fw_session_hairpin") {
		t, err := expandSystemSettingsFwSessionHairpin(d, v, "fw_session_hairpin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fw-session-hairpin"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok || d.HasChange("gateway") {
		t, err := expandSystemSettingsGateway(d, v, "gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}

	if v, ok := d.GetOk("gateway6"); ok || d.HasChange("gateway6") {
		t, err := expandSystemSettingsGateway6(d, v, "gateway6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway6"] = t
		}
	}

	if v, ok := d.GetOk("gtp_asym_fgsp"); ok || d.HasChange("gtp_asym_fgsp") {
		t, err := expandSystemSettingsGtpAsymFgsp(d, v, "gtp_asym_fgsp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtp-asym-fgsp"] = t
		}
	}

	if v, ok := d.GetOk("gtp_monitor_mode"); ok || d.HasChange("gtp_monitor_mode") {
		t, err := expandSystemSettingsGtpMonitorMode(d, v, "gtp_monitor_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtp-monitor-mode"] = t
		}
	}

	if v, ok := d.GetOk("gui_advanced_policy"); ok || d.HasChange("gui_advanced_policy") {
		t, err := expandSystemSettingsGuiAdvancedPolicy(d, v, "gui_advanced_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-advanced-policy"] = t
		}
	}

	if v, ok := d.GetOk("gui_advanced_wireless_features"); ok || d.HasChange("gui_advanced_wireless_features") {
		t, err := expandSystemSettingsGuiAdvancedWirelessFeatures(d, v, "gui_advanced_wireless_features")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-advanced-wireless-features"] = t
		}
	}

	if v, ok := d.GetOk("gui_allow_unnamed_policy"); ok || d.HasChange("gui_allow_unnamed_policy") {
		t, err := expandSystemSettingsGuiAllowUnnamedPolicy(d, v, "gui_allow_unnamed_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-allow-unnamed-policy"] = t
		}
	}

	if v, ok := d.GetOk("gui_antivirus"); ok || d.HasChange("gui_antivirus") {
		t, err := expandSystemSettingsGuiAntivirus(d, v, "gui_antivirus")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-antivirus"] = t
		}
	}

	if v, ok := d.GetOk("gui_ap_profile"); ok || d.HasChange("gui_ap_profile") {
		t, err := expandSystemSettingsGuiApProfile(d, v, "gui_ap_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-ap-profile"] = t
		}
	}

	if v, ok := d.GetOk("gui_application_control"); ok || d.HasChange("gui_application_control") {
		t, err := expandSystemSettingsGuiApplicationControl(d, v, "gui_application_control")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-application-control"] = t
		}
	}

	if v, ok := d.GetOk("gui_casb"); ok || d.HasChange("gui_casb") {
		t, err := expandSystemSettingsGuiCasb(d, v, "gui_casb")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-casb"] = t
		}
	}

	if v, ok := d.GetOk("gui_default_policy_columns"); ok || d.HasChange("gui_default_policy_columns") {
		t, err := expandSystemSettingsGuiDefaultPolicyColumns(d, v, "gui_default_policy_columns")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-default-policy-columns"] = t
		}
	}

	if v, ok := d.GetOk("gui_dhcp_advanced"); ok || d.HasChange("gui_dhcp_advanced") {
		t, err := expandSystemSettingsGuiDhcpAdvanced(d, v, "gui_dhcp_advanced")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-dhcp-advanced"] = t
		}
	}

	if v, ok := d.GetOk("gui_dlp_advanced"); ok || d.HasChange("gui_dlp_advanced") {
		t, err := expandSystemSettingsGuiDlpAdvanced(d, v, "gui_dlp_advanced")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-dlp-advanced"] = t
		}
	}

	if v, ok := d.GetOk("gui_dlp_profile"); ok || d.HasChange("gui_dlp_profile") {
		t, err := expandSystemSettingsGuiDlpProfile(d, v, "gui_dlp_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-dlp-profile"] = t
		}
	}

	if v, ok := d.GetOk("gui_dns_database"); ok || d.HasChange("gui_dns_database") {
		t, err := expandSystemSettingsGuiDnsDatabase(d, v, "gui_dns_database")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-dns-database"] = t
		}
	}

	if v, ok := d.GetOk("gui_dnsfilter"); ok || d.HasChange("gui_dnsfilter") {
		t, err := expandSystemSettingsGuiDnsfilter(d, v, "gui_dnsfilter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-dnsfilter"] = t
		}
	}

	if v, ok := d.GetOk("gui_domain_ip_reputation"); ok || d.HasChange("gui_domain_ip_reputation") {
		t, err := expandSystemSettingsGuiDomainIpReputation(d, v, "gui_domain_ip_reputation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-domain-ip-reputation"] = t
		}
	}

	if v, ok := d.GetOk("gui_dos_policy"); ok || d.HasChange("gui_dos_policy") {
		t, err := expandSystemSettingsGuiDosPolicy(d, v, "gui_dos_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-dos-policy"] = t
		}
	}

	if v, ok := d.GetOk("gui_dynamic_profile_display"); ok || d.HasChange("gui_dynamic_profile_display") {
		t, err := expandSystemSettingsGuiDynamicProfileDisplay(d, v, "gui_dynamic_profile_display")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-dynamic-profile-display"] = t
		}
	}

	if v, ok := d.GetOk("gui_dynamic_device_os_id"); ok || d.HasChange("gui_dynamic_device_os_id") {
		t, err := expandSystemSettingsGuiDynamicDeviceOsId(d, v, "gui_dynamic_device_os_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-dynamic-device-os-id"] = t
		}
	}

	if v, ok := d.GetOk("gui_dynamic_routing"); ok || d.HasChange("gui_dynamic_routing") {
		t, err := expandSystemSettingsGuiDynamicRouting(d, v, "gui_dynamic_routing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-dynamic-routing"] = t
		}
	}

	if v, ok := d.GetOk("gui_email_collection"); ok || d.HasChange("gui_email_collection") {
		t, err := expandSystemSettingsGuiEmailCollection(d, v, "gui_email_collection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-email-collection"] = t
		}
	}

	if v, ok := d.GetOk("gui_endpoint_control"); ok || d.HasChange("gui_endpoint_control") {
		t, err := expandSystemSettingsGuiEndpointControl(d, v, "gui_endpoint_control")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-endpoint-control"] = t
		}
	}

	if v, ok := d.GetOk("gui_endpoint_control_advanced"); ok || d.HasChange("gui_endpoint_control_advanced") {
		t, err := expandSystemSettingsGuiEndpointControlAdvanced(d, v, "gui_endpoint_control_advanced")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-endpoint-control-advanced"] = t
		}
	}

	if v, ok := d.GetOk("gui_enforce_change_summary"); ok || d.HasChange("gui_enforce_change_summary") {
		t, err := expandSystemSettingsGuiEnforceChangeSummary(d, v, "gui_enforce_change_summary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-enforce-change-summary"] = t
		}
	}

	if v, ok := d.GetOk("gui_explicit_proxy"); ok || d.HasChange("gui_explicit_proxy") {
		t, err := expandSystemSettingsGuiExplicitProxy(d, v, "gui_explicit_proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-explicit-proxy"] = t
		}
	}

	if v, ok := d.GetOk("gui_file_filter"); ok || d.HasChange("gui_file_filter") {
		t, err := expandSystemSettingsGuiFileFilter(d, v, "gui_file_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-file-filter"] = t
		}
	}

	if v, ok := d.GetOk("gui_fortiap_split_tunneling"); ok || d.HasChange("gui_fortiap_split_tunneling") {
		t, err := expandSystemSettingsGuiFortiapSplitTunneling(d, v, "gui_fortiap_split_tunneling")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-fortiap-split-tunneling"] = t
		}
	}

	if v, ok := d.GetOk("gui_fortiextender_controller"); ok || d.HasChange("gui_fortiextender_controller") {
		t, err := expandSystemSettingsGuiFortiextenderController(d, v, "gui_fortiextender_controller")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-fortiextender-controller"] = t
		}
	}

	if v, ok := d.GetOk("gui_gtp"); ok || d.HasChange("gui_gtp") {
		t, err := expandSystemSettingsGuiGtp(d, v, "gui_gtp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-gtp"] = t
		}
	}

	if v, ok := d.GetOk("gui_icap"); ok || d.HasChange("gui_icap") {
		t, err := expandSystemSettingsGuiIcap(d, v, "gui_icap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-icap"] = t
		}
	}

	if v, ok := d.GetOk("gui_implicit_policy"); ok || d.HasChange("gui_implicit_policy") {
		t, err := expandSystemSettingsGuiImplicitPolicy(d, v, "gui_implicit_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-implicit-policy"] = t
		}
	}

	if v, ok := d.GetOk("gui_ips"); ok || d.HasChange("gui_ips") {
		t, err := expandSystemSettingsGuiIps(d, v, "gui_ips")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-ips"] = t
		}
	}

	if v, ok := d.GetOk("gui_load_balance"); ok || d.HasChange("gui_load_balance") {
		t, err := expandSystemSettingsGuiLoadBalance(d, v, "gui_load_balance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-load-balance"] = t
		}
	}

	if v, ok := d.GetOk("gui_local_in_policy"); ok || d.HasChange("gui_local_in_policy") {
		t, err := expandSystemSettingsGuiLocalInPolicy(d, v, "gui_local_in_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-local-in-policy"] = t
		}
	}

	if v, ok := d.GetOk("gui_local_reports"); ok || d.HasChange("gui_local_reports") {
		t, err := expandSystemSettingsGuiLocalReports(d, v, "gui_local_reports")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-local-reports"] = t
		}
	}

	if v, ok := d.GetOk("gui_multicast_policy"); ok || d.HasChange("gui_multicast_policy") {
		t, err := expandSystemSettingsGuiMulticastPolicy(d, v, "gui_multicast_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-multicast-policy"] = t
		}
	}

	if v, ok := d.GetOk("gui_multiple_interface_policy"); ok || d.HasChange("gui_multiple_interface_policy") {
		t, err := expandSystemSettingsGuiMultipleInterfacePolicy(d, v, "gui_multiple_interface_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-multiple-interface-policy"] = t
		}
	}

	if v, ok := d.GetOk("gui_multiple_utm_profiles"); ok || d.HasChange("gui_multiple_utm_profiles") {
		t, err := expandSystemSettingsGuiMultipleUtmProfiles(d, v, "gui_multiple_utm_profiles")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-multiple-utm-profiles"] = t
		}
	}

	if v, ok := d.GetOk("gui_nat46_64"); ok || d.HasChange("gui_nat46_64") {
		t, err := expandSystemSettingsGuiNat4664(d, v, "gui_nat46_64")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-nat46-64"] = t
		}
	}

	if v, ok := d.GetOk("gui_object_colors"); ok || d.HasChange("gui_object_colors") {
		t, err := expandSystemSettingsGuiObjectColors(d, v, "gui_object_colors")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-object-colors"] = t
		}
	}

	if v, ok := d.GetOk("gui_per_policy_disclaimer"); ok || d.HasChange("gui_per_policy_disclaimer") {
		t, err := expandSystemSettingsGuiPerPolicyDisclaimer(d, v, "gui_per_policy_disclaimer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-per-policy-disclaimer"] = t
		}
	}

	if v, ok := d.GetOk("gui_ot"); ok || d.HasChange("gui_ot") {
		t, err := expandSystemSettingsGuiOt(d, v, "gui_ot")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-ot"] = t
		}
	}

	if v, ok := d.GetOk("gui_policy_based_ipsec"); ok || d.HasChange("gui_policy_based_ipsec") {
		t, err := expandSystemSettingsGuiPolicyBasedIpsec(d, v, "gui_policy_based_ipsec")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-policy-based-ipsec"] = t
		}
	}

	if v, ok := d.GetOk("gui_policy_disclaimer"); ok || d.HasChange("gui_policy_disclaimer") {
		t, err := expandSystemSettingsGuiPolicyDisclaimer(d, v, "gui_policy_disclaimer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-policy-disclaimer"] = t
		}
	}

	if v, ok := d.GetOk("gui_replacement_message_groups"); ok || d.HasChange("gui_replacement_message_groups") {
		t, err := expandSystemSettingsGuiReplacementMessageGroups(d, v, "gui_replacement_message_groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-replacement-message-groups"] = t
		}
	}

	if v, ok := d.GetOk("gui_proxy_inspection"); ok || d.HasChange("gui_proxy_inspection") {
		t, err := expandSystemSettingsGuiProxyInspection(d, v, "gui_proxy_inspection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-proxy-inspection"] = t
		}
	}

	if v, ok := d.GetOk("gui_route_tag_address_creation"); ok || d.HasChange("gui_route_tag_address_creation") {
		t, err := expandSystemSettingsGuiRouteTagAddressCreation(d, v, "gui_route_tag_address_creation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-route-tag-address-creation"] = t
		}
	}

	if v, ok := d.GetOk("gui_security_profile_group"); ok || d.HasChange("gui_security_profile_group") {
		t, err := expandSystemSettingsGuiSecurityProfileGroup(d, v, "gui_security_profile_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-security-profile-group"] = t
		}
	}

	if v, ok := d.GetOk("gui_spamfilter"); ok || d.HasChange("gui_spamfilter") {
		t, err := expandSystemSettingsGuiSpamfilter(d, v, "gui_spamfilter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-spamfilter"] = t
		}
	}

	if v, ok := d.GetOk("gui_sslvpn"); ok || d.HasChange("gui_sslvpn") {
		t, err := expandSystemSettingsGuiSslvpn(d, v, "gui_sslvpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-sslvpn"] = t
		}
	}

	if v, ok := d.GetOk("gui_sslvpn_clients"); ok || d.HasChange("gui_sslvpn_clients") {
		t, err := expandSystemSettingsGuiSslvpnClients(d, v, "gui_sslvpn_clients")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-sslvpn-clients"] = t
		}
	}

	if v, ok := d.GetOk("gui_sslvpn_personal_bookmarks"); ok || d.HasChange("gui_sslvpn_personal_bookmarks") {
		t, err := expandSystemSettingsGuiSslvpnPersonalBookmarks(d, v, "gui_sslvpn_personal_bookmarks")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-sslvpn-personal-bookmarks"] = t
		}
	}

	if v, ok := d.GetOk("gui_sslvpn_realms"); ok || d.HasChange("gui_sslvpn_realms") {
		t, err := expandSystemSettingsGuiSslvpnRealms(d, v, "gui_sslvpn_realms")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-sslvpn-realms"] = t
		}
	}

	if v, ok := d.GetOk("gui_switch_controller"); ok || d.HasChange("gui_switch_controller") {
		t, err := expandSystemSettingsGuiSwitchController(d, v, "gui_switch_controller")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-switch-controller"] = t
		}
	}

	if v, ok := d.GetOk("gui_threat_weight"); ok || d.HasChange("gui_threat_weight") {
		t, err := expandSystemSettingsGuiThreatWeight(d, v, "gui_threat_weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-threat-weight"] = t
		}
	}

	if v, ok := d.GetOk("gui_traffic_shaping"); ok || d.HasChange("gui_traffic_shaping") {
		t, err := expandSystemSettingsGuiTrafficShaping(d, v, "gui_traffic_shaping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-traffic-shaping"] = t
		}
	}

	if v, ok := d.GetOk("gui_videofilter"); ok || d.HasChange("gui_videofilter") {
		t, err := expandSystemSettingsGuiVideofilter(d, v, "gui_videofilter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-videofilter"] = t
		}
	}

	if v, ok := d.GetOk("gui_virtual_patch_profile"); ok || d.HasChange("gui_virtual_patch_profile") {
		t, err := expandSystemSettingsGuiVirtualPatchProfile(d, v, "gui_virtual_patch_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-virtual-patch-profile"] = t
		}
	}

	if v, ok := d.GetOk("gui_voip_profile"); ok || d.HasChange("gui_voip_profile") {
		t, err := expandSystemSettingsGuiVoipProfile(d, v, "gui_voip_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-voip-profile"] = t
		}
	}

	if v, ok := d.GetOk("gui_vpn"); ok || d.HasChange("gui_vpn") {
		t, err := expandSystemSettingsGuiVpn(d, v, "gui_vpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-vpn"] = t
		}
	}

	if v, ok := d.GetOk("gui_waf_profile"); ok || d.HasChange("gui_waf_profile") {
		t, err := expandSystemSettingsGuiWafProfile(d, v, "gui_waf_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-waf-profile"] = t
		}
	}

	if v, ok := d.GetOk("gui_wan_load_balancing"); ok || d.HasChange("gui_wan_load_balancing") {
		t, err := expandSystemSettingsGuiWanLoadBalancing(d, v, "gui_wan_load_balancing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-wan-load-balancing"] = t
		}
	}

	if v, ok := d.GetOk("gui_wanopt_cache"); ok || d.HasChange("gui_wanopt_cache") {
		t, err := expandSystemSettingsGuiWanoptCache(d, v, "gui_wanopt_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-wanopt-cache"] = t
		}
	}

	if v, ok := d.GetOk("gui_webfilter"); ok || d.HasChange("gui_webfilter") {
		t, err := expandSystemSettingsGuiWebfilter(d, v, "gui_webfilter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-webfilter"] = t
		}
	}

	if v, ok := d.GetOk("gui_webfilter_advanced"); ok || d.HasChange("gui_webfilter_advanced") {
		t, err := expandSystemSettingsGuiWebfilterAdvanced(d, v, "gui_webfilter_advanced")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-webfilter-advanced"] = t
		}
	}

	if v, ok := d.GetOk("gui_wireless_controller"); ok || d.HasChange("gui_wireless_controller") {
		t, err := expandSystemSettingsGuiWirelessController(d, v, "gui_wireless_controller")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-wireless-controller"] = t
		}
	}

	if v, ok := d.GetOk("gui_ztna"); ok || d.HasChange("gui_ztna") {
		t, err := expandSystemSettingsGuiZtna(d, v, "gui_ztna")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-ztna"] = t
		}
	}

	if v, ok := d.GetOk("h323_direct_model"); ok || d.HasChange("h323_direct_model") {
		t, err := expandSystemSettingsH323DirectModel(d, v, "h323_direct_model")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h323-direct-model"] = t
		}
	}

	if v, ok := d.GetOk("http_external_dest"); ok || d.HasChange("http_external_dest") {
		t, err := expandSystemSettingsHttpExternalDest(d, v, "http_external_dest")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-external-dest"] = t
		}
	}

	if v, ok := d.GetOk("hyperscale_default_policy_action"); ok || d.HasChange("hyperscale_default_policy_action") {
		t, err := expandSystemSettingsHyperscaleDefaultPolicyAction(d, v, "hyperscale_default_policy_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hyperscale-default-policy-action"] = t
		}
	}

	if v, ok := d.GetOk("ike_dn_format"); ok || d.HasChange("ike_dn_format") {
		t, err := expandSystemSettingsIkeDnFormat(d, v, "ike_dn_format")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-dn-format"] = t
		}
	}

	if v, ok := d.GetOk("ike_policy_route"); ok || d.HasChange("ike_policy_route") {
		t, err := expandSystemSettingsIkePolicyRoute(d, v, "ike_policy_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-policy-route"] = t
		}
	}

	if v, ok := d.GetOk("ike_port"); ok || d.HasChange("ike_port") {
		t, err := expandSystemSettingsIkePort(d, v, "ike_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-port"] = t
		}
	}

	if v, ok := d.GetOk("ike_quick_crash_detect"); ok || d.HasChange("ike_quick_crash_detect") {
		t, err := expandSystemSettingsIkeQuickCrashDetect(d, v, "ike_quick_crash_detect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-quick-crash-detect"] = t
		}
	}

	if v, ok := d.GetOk("ike_session_resume"); ok || d.HasChange("ike_session_resume") {
		t, err := expandSystemSettingsIkeSessionResume(d, v, "ike_session_resume")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-session-resume"] = t
		}
	}

	if v, ok := d.GetOk("implicit_allow_dns"); ok || d.HasChange("implicit_allow_dns") {
		t, err := expandSystemSettingsImplicitAllowDns(d, v, "implicit_allow_dns")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["implicit-allow-dns"] = t
		}
	}

	if v, ok := d.GetOk("ike_tcp_port"); ok || d.HasChange("ike_tcp_port") {
		t, err := expandSystemSettingsIkeTcpPort(d, v, "ike_tcp_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-tcp-port"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_app_ctrl_size"); ok || d.HasChange("internet_service_app_ctrl_size") {
		t, err := expandSystemSettingsInternetServiceAppCtrlSize(d, v, "internet_service_app_ctrl_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-app-ctrl-size"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_database_cache"); ok || d.HasChange("internet_service_database_cache") {
		t, err := expandSystemSettingsInternetServiceDatabaseCache(d, v, "internet_service_database_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-database-cache"] = t
		}
	}

	if v, ok := d.GetOk("intree_ses_best_route"); ok || d.HasChange("intree_ses_best_route") {
		t, err := expandSystemSettingsIntreeSesBestRoute(d, v, "intree_ses_best_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intree-ses-best-route"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandSystemSettingsIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("ip6"); ok || d.HasChange("ip6") {
		t, err := expandSystemSettingsIp6(d, v, "ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6"] = t
		}
	}

	if v, ok := d.GetOk("lan_extension_controller_addr"); ok || d.HasChange("lan_extension_controller_addr") {
		t, err := expandSystemSettingsLanExtensionControllerAddr(d, v, "lan_extension_controller_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lan-extension-controller-addr"] = t
		}
	}

	if v, ok := d.GetOk("link_down_access"); ok || d.HasChange("link_down_access") {
		t, err := expandSystemSettingsLinkDownAccess(d, v, "link_down_access")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-down-access"] = t
		}
	}

	if v, ok := d.GetOk("lldp_reception"); ok || d.HasChange("lldp_reception") {
		t, err := expandSystemSettingsLldpReception(d, v, "lldp_reception")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lldp-reception"] = t
		}
	}

	if v, ok := d.GetOk("lldp_transmission"); ok || d.HasChange("lldp_transmission") {
		t, err := expandSystemSettingsLldpTransmission(d, v, "lldp_transmission")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lldp-transmission"] = t
		}
	}

	if v, ok := d.GetOk("location_id"); ok || d.HasChange("location_id") {
		t, err := expandSystemSettingsLocationId(d, v, "location_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["location-id"] = t
		}
	}

	if v, ok := d.GetOk("mac_ttl"); ok || d.HasChange("mac_ttl") {
		t, err := expandSystemSettingsMacTtl(d, v, "mac_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-ttl"] = t
		}
	}

	if v, ok := d.GetOk("manageip"); ok || d.HasChange("manageip") {
		t, err := expandSystemSettingsManageip(d, v, "manageip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["manageip"] = t
		}
	}

	if v, ok := d.GetOk("manageip6"); ok || d.HasChange("manageip6") {
		t, err := expandSystemSettingsManageip6(d, v, "manageip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["manageip6"] = t
		}
	}

	if v, ok := d.GetOk("motherboard_traffic_forwarding"); ok || d.HasChange("motherboard_traffic_forwarding") {
		t, err := expandSystemSettingsMotherboardTrafficForwarding(d, v, "motherboard_traffic_forwarding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["motherboard-traffic-forwarding"] = t
		}
	}

	if v, ok := d.GetOk("multicast_forward"); ok || d.HasChange("multicast_forward") {
		t, err := expandSystemSettingsMulticastForward(d, v, "multicast_forward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-forward"] = t
		}
	}

	if v, ok := d.GetOk("multicast_skip_policy"); ok || d.HasChange("multicast_skip_policy") {
		t, err := expandSystemSettingsMulticastSkipPolicy(d, v, "multicast_skip_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-skip-policy"] = t
		}
	}

	if v, ok := d.GetOk("multicast_ttl_notchange"); ok || d.HasChange("multicast_ttl_notchange") {
		t, err := expandSystemSettingsMulticastTtlNotchange(d, v, "multicast_ttl_notchange")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-ttl-notchange"] = t
		}
	}

	if v, ok := d.GetOk("nat46_force_ipv4_packet_forwarding"); ok || d.HasChange("nat46_force_ipv4_packet_forwarding") {
		t, err := expandSystemSettingsNat46ForceIpv4PacketForwarding(d, v, "nat46_force_ipv4_packet_forwarding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat46-force-ipv4-packet-forwarding"] = t
		}
	}

	if v, ok := d.GetOk("nat46_generate_ipv6_fragment_header"); ok || d.HasChange("nat46_generate_ipv6_fragment_header") {
		t, err := expandSystemSettingsNat46GenerateIpv6FragmentHeader(d, v, "nat46_generate_ipv6_fragment_header")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat46-generate-ipv6-fragment-header"] = t
		}
	}

	if v, ok := d.GetOk("nat64_force_ipv6_packet_forwarding"); ok || d.HasChange("nat64_force_ipv6_packet_forwarding") {
		t, err := expandSystemSettingsNat64ForceIpv6PacketForwarding(d, v, "nat64_force_ipv6_packet_forwarding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat64-force-ipv6-packet-forwarding"] = t
		}
	}

	if v, ok := d.GetOk("ngfw_mode"); ok || d.HasChange("ngfw_mode") {
		t, err := expandSystemSettingsNgfwMode(d, v, "ngfw_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ngfw-mode"] = t
		}
	}

	if v, ok := d.GetOk("nonat_eif_key_sel"); ok || d.HasChange("nonat_eif_key_sel") {
		t, err := expandSystemSettingsNonatEifKeySel(d, v, "nonat_eif_key_sel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nonat-eif-key-sel"] = t
		}
	}

	if v, ok := d.GetOk("npu_group_id"); ok || d.HasChange("npu_group_id") {
		t, err := expandSystemSettingsNpuGroupId(d, v, "npu_group_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["npu-group-id"] = t
		}
	}

	if v, ok := d.GetOk("opmode"); ok || d.HasChange("opmode") {
		t, err := expandSystemSettingsOpmode(d, v, "opmode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["opmode"] = t
		}
	}

	if v, ok := d.GetOk("pfcp_monitor_mode"); ok || d.HasChange("pfcp_monitor_mode") {
		t, err := expandSystemSettingsPfcpMonitorMode(d, v, "pfcp_monitor_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pfcp-monitor-mode"] = t
		}
	}

	if v, ok := d.GetOk("policy_offload_level"); ok || d.HasChange("policy_offload_level") {
		t, err := expandSystemSettingsPolicyOffloadLevel(d, v, "policy_offload_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-offload-level"] = t
		}
	}

	if v, ok := d.GetOk("prp_trailer_action"); ok || d.HasChange("prp_trailer_action") {
		t, err := expandSystemSettingsPrpTrailerAction(d, v, "prp_trailer_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prp-trailer-action"] = t
		}
	}

	if v, ok := d.GetOk("sccp_port"); ok || d.HasChange("sccp_port") {
		t, err := expandSystemSettingsSccpPort(d, v, "sccp_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sccp-port"] = t
		}
	}

	if v, ok := d.GetOk("sctp_session_without_init"); ok || d.HasChange("sctp_session_without_init") {
		t, err := expandSystemSettingsSctpSessionWithoutInit(d, v, "sctp_session_without_init")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sctp-session-without-init"] = t
		}
	}

	if v, ok := d.GetOk("ses_denied_multicast_traffic"); ok || d.HasChange("ses_denied_multicast_traffic") {
		t, err := expandSystemSettingsSesDeniedMulticastTraffic(d, v, "ses_denied_multicast_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ses-denied-multicast-traffic"] = t
		}
	}

	if v, ok := d.GetOk("ses_denied_traffic"); ok || d.HasChange("ses_denied_traffic") {
		t, err := expandSystemSettingsSesDeniedTraffic(d, v, "ses_denied_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ses-denied-traffic"] = t
		}
	}

	if v, ok := d.GetOk("session_insert_trial"); ok || d.HasChange("session_insert_trial") {
		t, err := expandSystemSettingsSessionInsertTrial(d, v, "session_insert_trial")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-insert-trial"] = t
		}
	}

	if v, ok := d.GetOk("sip_expectation"); ok || d.HasChange("sip_expectation") {
		t, err := expandSystemSettingsSipExpectation(d, v, "sip_expectation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sip-expectation"] = t
		}
	}

	if v, ok := d.GetOk("sip_nat_trace"); ok || d.HasChange("sip_nat_trace") {
		t, err := expandSystemSettingsSipNatTrace(d, v, "sip_nat_trace")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sip-nat-trace"] = t
		}
	}

	if v, ok := d.GetOk("sip_ssl_port"); ok || d.HasChange("sip_ssl_port") {
		t, err := expandSystemSettingsSipSslPort(d, v, "sip_ssl_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sip-ssl-port"] = t
		}
	}

	if v, ok := d.GetOk("sip_tcp_port"); ok || d.HasChange("sip_tcp_port") {
		t, err := expandSystemSettingsSipTcpPort(d, v, "sip_tcp_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sip-tcp-port"] = t
		}
	}

	if v, ok := d.GetOk("sip_udp_port"); ok || d.HasChange("sip_udp_port") {
		t, err := expandSystemSettingsSipUdpPort(d, v, "sip_udp_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sip-udp-port"] = t
		}
	}

	if v, ok := d.GetOk("snat_hairpin_traffic"); ok || d.HasChange("snat_hairpin_traffic") {
		t, err := expandSystemSettingsSnatHairpinTraffic(d, v, "snat_hairpin_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["snat-hairpin-traffic"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemSettingsStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("strict_src_check"); ok || d.HasChange("strict_src_check") {
		t, err := expandSystemSettingsStrictSrcCheck(d, v, "strict_src_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strict-src-check"] = t
		}
	}

	if v, ok := d.GetOk("tcp_session_without_syn"); ok || d.HasChange("tcp_session_without_syn") {
		t, err := expandSystemSettingsTcpSessionWithoutSyn(d, v, "tcp_session_without_syn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-session-without-syn"] = t
		}
	}

	if v, ok := d.GetOk("trap_local_session"); ok || d.HasChange("trap_local_session") {
		t, err := expandSystemSettingsTrapLocalSession(d, v, "trap_local_session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-local-session"] = t
		}
	}

	if v, ok := d.GetOk("trap_session_flag"); ok || d.HasChange("trap_session_flag") {
		t, err := expandSystemSettingsTrapSessionFlag(d, v, "trap_session_flag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-session-flag"] = t
		}
	}

	if v, ok := d.GetOk("utf8_spam_tagging"); ok || d.HasChange("utf8_spam_tagging") {
		t, err := expandSystemSettingsUtf8SpamTagging(d, v, "utf8_spam_tagging")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utf8-spam-tagging"] = t
		}
	}

	if v, ok := d.GetOk("v4_ecmp_mode"); ok || d.HasChange("v4_ecmp_mode") {
		t, err := expandSystemSettingsV4EcmpMode(d, v, "v4_ecmp_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["v4-ecmp-mode"] = t
		}
	}

	if v, ok := d.GetOk("vdom_type"); ok || d.HasChange("vdom_type") {
		t, err := expandSystemSettingsVdomType(d, v, "vdom_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom-type"] = t
		}
	}

	if v, ok := d.GetOk("vpn_stats_log"); ok || d.HasChange("vpn_stats_log") {
		t, err := expandSystemSettingsVpnStatsLog(d, v, "vpn_stats_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpn-stats-log"] = t
		}
	}

	if v, ok := d.GetOk("vpn_stats_period"); ok || d.HasChange("vpn_stats_period") {
		t, err := expandSystemSettingsVpnStatsPeriod(d, v, "vpn_stats_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpn-stats-period"] = t
		}
	}

	if v, ok := d.GetOk("wccp_cache_engine"); ok || d.HasChange("wccp_cache_engine") {
		t, err := expandSystemSettingsWccpCacheEngine(d, v, "wccp_cache_engine")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wccp-cache-engine"] = t
		}
	}

	return &obj, nil
}
