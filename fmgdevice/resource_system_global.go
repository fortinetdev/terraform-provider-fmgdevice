// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure global attributes.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemGlobal() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemGlobalUpdate,
		Read:   resourceSystemGlobalRead,
		Update: resourceSystemGlobalUpdate,
		Delete: resourceSystemGlobalDelete,

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
			"admin_ble_button": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"admin_concurrent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_console_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"admin_forticloud_sso_default_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"admin_forticloud_sso_login": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"admin_hsts_max_age": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"admin_https_pki_required": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_https_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_https_ssl_banned_ciphers": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"admin_https_ssl_ciphersuites": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"admin_https_ssl_versions": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"admin_lockout_duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"admin_lockout_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"admin_login_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"admin_maintainer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"admin_reset_button": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"admin_restrict_local": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_scp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_server_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"admin_sport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"admin_ssh_grace_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"admin_ssh_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_ssh_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"admin_ssh_v1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_telnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_telnet_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"admintimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"airplane_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"alias": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_traffic_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"anti_replay": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"arp_max_entry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auth_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"auth_http_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auth_https_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auth_ike_saml_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auth_keepalive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_session_limit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_auth_extension_device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"autorun_log_fsck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"av_affinity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"av_failopen": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"av_failopen_session": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"batch_cmdb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bfd_affinity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_session_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"br_fdb_max_entry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cert_chain_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cfg_revert_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cfg_save": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"check_protocol_header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"check_reset_range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cli_audit_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cloud_communication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clt_cert_req": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cmdbsvr_affinity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cpu_use_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"csr_ca_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"daily_restart": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_service_source_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_idle_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dh_params": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_lease_backup_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dnsproxy_worker_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dp_fragment_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dp_pinhole_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dp_rsync_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dp_tcp_normal_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dp_udp_idle_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"early_tcp_npu_session": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"endpoint_control_fds_access": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"edit_vdom_prompt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extender_controller_reserved_network": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"faz_disk_buffer_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fds_statistics": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fds_statistics_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fec_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fgd_alert_subscription": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"forticarrier_bypass": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"forticontroller_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"forticontroller_proxy_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"forticonverter_config_upload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forticonverter_integration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiextender": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiextender_data_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fortiextender_discovery_lockdown": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiextender_provision_on_authorization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiextender_vlan_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiipam_integration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fortigslb_integration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiservice_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fortitoken_cloud": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortitoken_cloud_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gui_allow_default_hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortitoken_cloud_push_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortitoken_cloud_sync_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gui_allow_incompatible_fabric_fgt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_app_detection_sdwan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_auto_upgrade_setup_warning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_cdn_domain_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gui_cdn_usage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_certificates": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_custom_language": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_date_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_date_time_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_device_latitude": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gui_device_longitude": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gui_display_hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_firmware_upgrade_setup_warning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gui_firmware_upgrade_warning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_forticare_registration_setup_warning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_fortisandbox_cloud": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gui_fortigate_cloud_sandbox": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_fortiguard_resource_fetch": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_lines_per_page": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"gui_local_out": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_replacement_message_groups": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_rest_api_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_theme": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_wireless_opensecurity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_workflow_management": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ha_affinity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"honor_df": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hw_switch_ether_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_request_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http_unauthenticated_request_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hyper_scale_vdom_num": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"igmp_state_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"interface_subnet_usage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internal_switch_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internal_switch_speed": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_database": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_download_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ip_fragment_mem_thresholds": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ip_src_port_range": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ips_affinity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec_asic_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipsec_ha_seqjump_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ipsec_hmac_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipsec_qat_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec_round_robin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec_soft_dec_async": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_accept_dad": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ipv6_allow_anycast_probe": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_allow_local_in_silent_drop": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_allow_local_in_slient_drop": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_allow_multicast_probe": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_allow_traffic_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"irq_time_accounting": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"language": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldapconntimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"legacy_poe_device_support": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"log_single_cpu_high": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_ssl_connection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_uuid_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_uuid_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"login_timestamp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"long_vdom_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"management_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"management_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"management_port_use_admin_sport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"management_vdom": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"max_route_cache_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"memory_use_threshold_extreme": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"memory_use_threshold_green": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"memory_use_threshold_red": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"miglog_affinity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"miglogd_children": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"multi_factor_authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ndp_max_entry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"npu_neighbor_update": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"optimize_flow_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"per_user_bwl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"per_user_bal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pmtu_discovery": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policy_auth_concurrent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"post_login_banner": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pre_login_banner": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"private_data_encryption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_and_explicit_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"proxy_auth_lifetime": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_auth_lifetime_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"proxy_auth_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"proxy_cipher_hardware_acceleration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"proxy_kxp_hardware_acceleration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"proxy_cert_use_mgmt_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_hardware_acceleration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"proxy_keep_alive_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_re_authentication_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_re_authentication_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"proxy_resource_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_worker_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"purdue_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"qsfp28_40g_port": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"qsfpdd_100g_port": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"qsfpdd_split8_port": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"quic_ack_thresold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"quic_congestion_control_algo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"quic_max_datagram_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"quic_pmtud": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"quic_tls_handshake_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"quic_udp_payload_size_shaping_per_cid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"reboot_upon_config_restore": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"refresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"remoteauthtimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"reset_sessionless_tcp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"restart_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"revision_backup_on_logout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"revision_image_auto_backup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scanunit_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"security_rating_result_submission": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_rating_run_on_schedule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"send_pmtu_icmp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sflowd_max_children_num": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"show_backplane_intf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"snat_route_change": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"special_file_23_support": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"speedtest_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"speedtestd_ctrl_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"speedtestd_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"split_port": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"split_port_mode": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"split_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ssd_trim_date": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ssd_trim_freq": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssd_trim_hour": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ssd_trim_min": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ssd_trim_weekday": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_cbc_cipher": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_enc_algo": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssh_hmac_md5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_hostkey": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_hostkey_algo": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssh_hostkey_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_hostkey_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"ssh_kex_algo": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssh_kex_sha1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_mac_algo": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssh_mac_weak": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_min_proto_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_static_key_ciphers": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslvpn_cipher_hardware_acceleration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sslvpn_ems_sn_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslvpn_kxp_hardware_acceleration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sslvpn_max_worker_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sslvpn_plugin_version_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslvpn_web_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"strict_dirty_session_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"strong_crypto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller_reserved_network": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sys_file_check_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sys_perf_log_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"syslog_affinity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_halfclose_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcp_halfopen_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcp_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_rst_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcp_timewait_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tftp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timezone": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"traffic_priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_priority_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"two_factor_email_expiry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"two_factor_fac_expiry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"two_factor_ftk_expiry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"two_factor_ftm_expiry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"two_factor_sms_expiry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"udp_idle_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"url_filter_affinity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url_filter_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"user_device_store_max_devices": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"user_device_store_max_unified_mem": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"user_device_store_max_users": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"user_server_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vdom_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip_arp_range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"virtual_server_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"virtual_server_hardware_acceleration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtual_switch_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vpn_ems_sn_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wad_affinity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wad_csvc_cs_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"wad_csvc_db_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"wad_memory_change_granularity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"wad_restart_end_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wad_restart_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wad_restart_start_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wad_source_affinity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wad_worker_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"wifi_ca_certificate": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"wifi_certificate": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"wimax_4g_usb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wireless_controller": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wireless_controller_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"wireless_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"xstools_update_frequency": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemGlobalUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemGlobal(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemGlobal resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemGlobal(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemGlobal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemGlobal")

	return resourceSystemGlobalRead(d, m)
}

func resourceSystemGlobalDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	err = c.DeleteSystemGlobal(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemGlobal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemGlobalRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	paradict["device"] = device_name

	o, err := c.ReadSystemGlobal(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemGlobal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemGlobal(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemGlobal resource from API: %v", err)
	}
	return nil
}

func flattenSystemGlobalAdminBleButton(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdminConcurrent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdminConsoleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdminForticloudSsoDefaultProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalAdminForticloudSsoLogin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdminHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdminHstsMaxAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdminHttpsPkiRequired(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdminHttpsRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdminHttpsSslBannedCiphers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalAdminHttpsSslCiphersuites(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalAdminHttpsSslVersions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalAdminLockoutDuration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdminLockoutThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdminLoginMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdminMaintainer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdminPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdminResetButton(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdminRestrictLocal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdminScp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdminServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalAdminSport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdminSshGraceTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdminSshPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdminSshPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdminSshV1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdminTelnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdminTelnetPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdmintimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAirplaneMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAlias(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAllowTrafficRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAntiReplay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalArpMaxEntry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAuthCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalAuthHttpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAuthHttpsPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAuthIkeSamlPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAuthKeepalive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAuthSessionLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAutoAuthExtensionDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAutorunLogFsck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAvAffinity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAvFailopen(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAvFailopenSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalBatchCmdb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalBfdAffinity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalBlockSessionTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalBrFdbMaxEntry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalCertChainMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalCfgRevertTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalCfgSave(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalCheckProtocolHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalCheckResetRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalCliAuditLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalCloudCommunication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalCltCertReq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalCmdbsvrAffinity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalCpuUseThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalCsrCaAttribute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalDailyRestart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalDefaultServiceSourcePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalDeviceIdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalDhParams(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalDhcpLeaseBackupInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalDnsproxyWorkerCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalDpFragmentTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalDpPinholeTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalDpRsyncTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalDpTcpNormalTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalDpUdpIdleTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalDst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalEarlyTcpNpuSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalEndpointControlFdsAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalEditVdomPrompt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalExtenderControllerReservedNetwork(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalFazDiskBufferSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFdsStatistics(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFdsStatisticsPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFecPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFgdAlertSubscription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalForticarrierBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalForticontrollerProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalForticontrollerProxyPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalForticonverterConfigUpload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalForticonverterIntegration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFortiextender(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFortiextenderDataPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFortiextenderDiscoveryLockdown(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFortiextenderProvisionOnAuthorization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFortiextenderVlanMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFortiipamIntegration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFortigslbIntegration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFortiservicePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFortitokenCloud(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFortitokenCloudService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGuiAllowDefaultHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFortitokenCloudPushStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFortitokenCloudSyncInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGuiAllowIncompatibleFabricFgt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGuiAppDetectionSdwan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGuiAutoUpgradeSetupWarning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGuiCdnDomainOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGuiCdnUsage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGuiCertificates(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGuiCustomLanguage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGuiDateFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGuiDateTimeSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGuiDeviceLatitude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGuiDeviceLongitude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGuiDisplayHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGuiFirmwareUpgradeSetupWarning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGuiFirmwareUpgradeWarning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGuiForticareRegistrationSetupWarning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGuiFortisandboxCloud(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGuiFortigateCloudSandbox(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGuiFortiguardResourceFetch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGuiIpv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGuiLinesPerPage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGuiLocalOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGuiReplacementMessageGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGuiRestApiCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGuiTheme(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGuiWirelessOpensecurity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGuiWorkflowManagement(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalHaAffinity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalHonorDf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalHwSwitchEtherFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalHttpRequestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalHttpUnauthenticatedRequestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalHyperScaleVdomNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalIgmpStateLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalInterfaceSubnetUsage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalInternalSwitchMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalInternalSwitchSpeed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalInternetServiceDatabase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalInternetServiceDownloadList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalIpFragmentMemThresholds(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalIpSrcPortRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalIpsAffinity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalIpsecAsicOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalIpsecHaSeqjumpRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalIpsecHmacOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalIpsecQatOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalIpsecRoundRobin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalIpsecSoftDecAsync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalIpv6AcceptDad(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalIpv6AllowAnycastProbe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalIpv6AllowLocalInSilentDrop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalIpv6AllowLocalInSlientDrop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalIpv6AllowMulticastProbe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalIpv6AllowTrafficRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalIrqTimeAccounting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalLanguage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalLdapconntimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalLegacyPoeDeviceSupport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalLldpReception(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalLldpTransmission(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalLogSingleCpuHigh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalLogSslConnection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalLogUuidAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalLogUuidPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalLoginTimestamp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalLongVdomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalManagementIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalManagementPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalManagementPortUseAdminSport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalManagementVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalMaxRouteCacheSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalMemoryUseThresholdExtreme(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalMemoryUseThresholdGreen(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalMemoryUseThresholdRed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalMiglogAffinity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalMiglogdChildren(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalMultiFactorAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalNdpMaxEntry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalNpuNeighborUpdate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalOptimizeFlowMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalPerUserBwl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalPerUserBal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalPmtuDiscovery(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalPolicyAuthConcurrent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalPostLoginBanner(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalPreLoginBanner(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalPrivateDataEncryption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalProxyAndExplicitProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalProxyAuthLifetime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalProxyAuthLifetimeTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalProxyAuthTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalProxyCipherHardwareAcceleration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalProxyKxpHardwareAcceleration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalProxyCertUseMgmtVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalProxyHardwareAcceleration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalProxyKeepAliveMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalProxyReAuthenticationMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalProxyReAuthenticationTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalProxyResourceMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalProxyWorkerCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalPurdueLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalQsfp2840GPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalQsfpdd100GPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalQsfpddSplit8Port(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalQuicAckThresold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalQuicCongestionControlAlgo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalQuicMaxDatagramSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalQuicPmtud(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalQuicTlsHandshakeTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalQuicUdpPayloadSizeShapingPerCid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalRadiusPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalRebootUponConfigRestore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalRefresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalRemoteauthtimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalResetSessionlessTcp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalRestartTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalRevisionBackupOnLogout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalRevisionImageAutoBackup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalScanunitCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSecurityRatingResultSubmission(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSecurityRatingRunOnSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSendPmtuIcmp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSflowdMaxChildrenNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalShowBackplaneIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSnatRouteChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSpecialFile23Support(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSpeedtestServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSpeedtestdCtrlPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSpeedtestdServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSplitPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalSplitPortMode(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenSystemGlobalSplitPortModeInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "SystemGlobal-SplitPortMode-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_mode"
		if _, ok := i["split-mode"]; ok {
			v := flattenSystemGlobalSplitPortModeSplitMode(i["split-mode"], d, pre_append)
			tmp["split_mode"] = fortiAPISubPartPatch(v, "SystemGlobal-SplitPortMode-SplitMode")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemGlobalSplitPortModeInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSplitPortModeSplitMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSsdTrimDate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSsdTrimFreq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSsdTrimHour(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSsdTrimMin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSsdTrimWeekday(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSshCbcCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSshEncAlgo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalSshHmacMd5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSshHostkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSshHostkeyAlgo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalSshHostkeyOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSshKexAlgo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalSshKexSha1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSshMacAlgo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalSshMacWeak(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSslStaticKeyCiphers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSslvpnCipherHardwareAcceleration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSslvpnEmsSnCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSslvpnKxpHardwareAcceleration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSslvpnMaxWorkerCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSslvpnPluginVersionCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSslvpnWebMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalStrictDirtySessionCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalStrongCrypto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSwitchController(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSwitchControllerReservedNetwork(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalSysFileCheckInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSysPerfLogInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSyslogAffinity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalTcpHalfcloseTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalTcpHalfopenTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalTcpOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalTcpRstTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalTcpTimewaitTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalTftp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalTimezone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convstr2list(v, d.Get(pre))
}

func flattenSystemGlobalTrafficPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalTrafficPriorityLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalTwoFactorEmailExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalTwoFactorFacExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalTwoFactorFtkExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalTwoFactorFtmExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalTwoFactorSmsExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalUdpIdleTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalUrlFilterAffinity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalUrlFilterCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalUserDeviceStoreMaxDevices(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalUserDeviceStoreMaxUnifiedMem(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalUserDeviceStoreMaxUsers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalUserServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalVdomMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalVipArpRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalVirtualServerCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalVirtualServerHardwareAcceleration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalVirtualSwitchVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalVpnEmsSnCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalWadAffinity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalWadCsvcCsCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalWadCsvcDbCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalWadMemoryChangeGranularity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalWadRestartEndTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalWadRestartMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalWadRestartStartTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalWadSourceAffinity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalWadWorkerCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalWifiCaCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalWifiCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalWimax4GUsb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalWirelessController(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalWirelessControllerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalWirelessMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalXstoolsUpdateFrequency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemGlobal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("admin_ble_button", flattenSystemGlobalAdminBleButton(o["admin-ble-button"], d, "admin_ble_button")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-ble-button"], "SystemGlobal-AdminBleButton"); ok {
			if err = d.Set("admin_ble_button", vv); err != nil {
				return fmt.Errorf("Error reading admin_ble_button: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_ble_button: %v", err)
		}
	}

	if err = d.Set("admin_concurrent", flattenSystemGlobalAdminConcurrent(o["admin-concurrent"], d, "admin_concurrent")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-concurrent"], "SystemGlobal-AdminConcurrent"); ok {
			if err = d.Set("admin_concurrent", vv); err != nil {
				return fmt.Errorf("Error reading admin_concurrent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_concurrent: %v", err)
		}
	}

	if err = d.Set("admin_console_timeout", flattenSystemGlobalAdminConsoleTimeout(o["admin-console-timeout"], d, "admin_console_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-console-timeout"], "SystemGlobal-AdminConsoleTimeout"); ok {
			if err = d.Set("admin_console_timeout", vv); err != nil {
				return fmt.Errorf("Error reading admin_console_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_console_timeout: %v", err)
		}
	}

	if err = d.Set("admin_forticloud_sso_default_profile", flattenSystemGlobalAdminForticloudSsoDefaultProfile(o["admin-forticloud-sso-default-profile"], d, "admin_forticloud_sso_default_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-forticloud-sso-default-profile"], "SystemGlobal-AdminForticloudSsoDefaultProfile"); ok {
			if err = d.Set("admin_forticloud_sso_default_profile", vv); err != nil {
				return fmt.Errorf("Error reading admin_forticloud_sso_default_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_forticloud_sso_default_profile: %v", err)
		}
	}

	if err = d.Set("admin_forticloud_sso_login", flattenSystemGlobalAdminForticloudSsoLogin(o["admin-forticloud-sso-login"], d, "admin_forticloud_sso_login")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-forticloud-sso-login"], "SystemGlobal-AdminForticloudSsoLogin"); ok {
			if err = d.Set("admin_forticloud_sso_login", vv); err != nil {
				return fmt.Errorf("Error reading admin_forticloud_sso_login: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_forticloud_sso_login: %v", err)
		}
	}

	if err = d.Set("admin_host", flattenSystemGlobalAdminHost(o["admin-host"], d, "admin_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-host"], "SystemGlobal-AdminHost"); ok {
			if err = d.Set("admin_host", vv); err != nil {
				return fmt.Errorf("Error reading admin_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_host: %v", err)
		}
	}

	if err = d.Set("admin_hsts_max_age", flattenSystemGlobalAdminHstsMaxAge(o["admin-hsts-max-age"], d, "admin_hsts_max_age")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-hsts-max-age"], "SystemGlobal-AdminHstsMaxAge"); ok {
			if err = d.Set("admin_hsts_max_age", vv); err != nil {
				return fmt.Errorf("Error reading admin_hsts_max_age: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_hsts_max_age: %v", err)
		}
	}

	if err = d.Set("admin_https_pki_required", flattenSystemGlobalAdminHttpsPkiRequired(o["admin-https-pki-required"], d, "admin_https_pki_required")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-https-pki-required"], "SystemGlobal-AdminHttpsPkiRequired"); ok {
			if err = d.Set("admin_https_pki_required", vv); err != nil {
				return fmt.Errorf("Error reading admin_https_pki_required: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_https_pki_required: %v", err)
		}
	}

	if err = d.Set("admin_https_redirect", flattenSystemGlobalAdminHttpsRedirect(o["admin-https-redirect"], d, "admin_https_redirect")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-https-redirect"], "SystemGlobal-AdminHttpsRedirect"); ok {
			if err = d.Set("admin_https_redirect", vv); err != nil {
				return fmt.Errorf("Error reading admin_https_redirect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_https_redirect: %v", err)
		}
	}

	if err = d.Set("admin_https_ssl_banned_ciphers", flattenSystemGlobalAdminHttpsSslBannedCiphers(o["admin-https-ssl-banned-ciphers"], d, "admin_https_ssl_banned_ciphers")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-https-ssl-banned-ciphers"], "SystemGlobal-AdminHttpsSslBannedCiphers"); ok {
			if err = d.Set("admin_https_ssl_banned_ciphers", vv); err != nil {
				return fmt.Errorf("Error reading admin_https_ssl_banned_ciphers: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_https_ssl_banned_ciphers: %v", err)
		}
	}

	if err = d.Set("admin_https_ssl_ciphersuites", flattenSystemGlobalAdminHttpsSslCiphersuites(o["admin-https-ssl-ciphersuites"], d, "admin_https_ssl_ciphersuites")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-https-ssl-ciphersuites"], "SystemGlobal-AdminHttpsSslCiphersuites"); ok {
			if err = d.Set("admin_https_ssl_ciphersuites", vv); err != nil {
				return fmt.Errorf("Error reading admin_https_ssl_ciphersuites: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_https_ssl_ciphersuites: %v", err)
		}
	}

	if err = d.Set("admin_https_ssl_versions", flattenSystemGlobalAdminHttpsSslVersions(o["admin-https-ssl-versions"], d, "admin_https_ssl_versions")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-https-ssl-versions"], "SystemGlobal-AdminHttpsSslVersions"); ok {
			if err = d.Set("admin_https_ssl_versions", vv); err != nil {
				return fmt.Errorf("Error reading admin_https_ssl_versions: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_https_ssl_versions: %v", err)
		}
	}

	if err = d.Set("admin_lockout_duration", flattenSystemGlobalAdminLockoutDuration(o["admin-lockout-duration"], d, "admin_lockout_duration")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-lockout-duration"], "SystemGlobal-AdminLockoutDuration"); ok {
			if err = d.Set("admin_lockout_duration", vv); err != nil {
				return fmt.Errorf("Error reading admin_lockout_duration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_lockout_duration: %v", err)
		}
	}

	if err = d.Set("admin_lockout_threshold", flattenSystemGlobalAdminLockoutThreshold(o["admin-lockout-threshold"], d, "admin_lockout_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-lockout-threshold"], "SystemGlobal-AdminLockoutThreshold"); ok {
			if err = d.Set("admin_lockout_threshold", vv); err != nil {
				return fmt.Errorf("Error reading admin_lockout_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_lockout_threshold: %v", err)
		}
	}

	if err = d.Set("admin_login_max", flattenSystemGlobalAdminLoginMax(o["admin-login-max"], d, "admin_login_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-login-max"], "SystemGlobal-AdminLoginMax"); ok {
			if err = d.Set("admin_login_max", vv); err != nil {
				return fmt.Errorf("Error reading admin_login_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_login_max: %v", err)
		}
	}

	if err = d.Set("admin_maintainer", flattenSystemGlobalAdminMaintainer(o["admin-maintainer"], d, "admin_maintainer")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-maintainer"], "SystemGlobal-AdminMaintainer"); ok {
			if err = d.Set("admin_maintainer", vv); err != nil {
				return fmt.Errorf("Error reading admin_maintainer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_maintainer: %v", err)
		}
	}

	if err = d.Set("admin_port", flattenSystemGlobalAdminPort(o["admin-port"], d, "admin_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-port"], "SystemGlobal-AdminPort"); ok {
			if err = d.Set("admin_port", vv); err != nil {
				return fmt.Errorf("Error reading admin_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_port: %v", err)
		}
	}

	if err = d.Set("admin_reset_button", flattenSystemGlobalAdminResetButton(o["admin-reset-button"], d, "admin_reset_button")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-reset-button"], "SystemGlobal-AdminResetButton"); ok {
			if err = d.Set("admin_reset_button", vv); err != nil {
				return fmt.Errorf("Error reading admin_reset_button: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_reset_button: %v", err)
		}
	}

	if err = d.Set("admin_restrict_local", flattenSystemGlobalAdminRestrictLocal(o["admin-restrict-local"], d, "admin_restrict_local")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-restrict-local"], "SystemGlobal-AdminRestrictLocal"); ok {
			if err = d.Set("admin_restrict_local", vv); err != nil {
				return fmt.Errorf("Error reading admin_restrict_local: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_restrict_local: %v", err)
		}
	}

	if err = d.Set("admin_scp", flattenSystemGlobalAdminScp(o["admin-scp"], d, "admin_scp")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-scp"], "SystemGlobal-AdminScp"); ok {
			if err = d.Set("admin_scp", vv); err != nil {
				return fmt.Errorf("Error reading admin_scp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_scp: %v", err)
		}
	}

	if err = d.Set("admin_server_cert", flattenSystemGlobalAdminServerCert(o["admin-server-cert"], d, "admin_server_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-server-cert"], "SystemGlobal-AdminServerCert"); ok {
			if err = d.Set("admin_server_cert", vv); err != nil {
				return fmt.Errorf("Error reading admin_server_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_server_cert: %v", err)
		}
	}

	if err = d.Set("admin_sport", flattenSystemGlobalAdminSport(o["admin-sport"], d, "admin_sport")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-sport"], "SystemGlobal-AdminSport"); ok {
			if err = d.Set("admin_sport", vv); err != nil {
				return fmt.Errorf("Error reading admin_sport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_sport: %v", err)
		}
	}

	if err = d.Set("admin_ssh_grace_time", flattenSystemGlobalAdminSshGraceTime(o["admin-ssh-grace-time"], d, "admin_ssh_grace_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-ssh-grace-time"], "SystemGlobal-AdminSshGraceTime"); ok {
			if err = d.Set("admin_ssh_grace_time", vv); err != nil {
				return fmt.Errorf("Error reading admin_ssh_grace_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_ssh_grace_time: %v", err)
		}
	}

	if err = d.Set("admin_ssh_password", flattenSystemGlobalAdminSshPassword(o["admin-ssh-password"], d, "admin_ssh_password")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-ssh-password"], "SystemGlobal-AdminSshPassword"); ok {
			if err = d.Set("admin_ssh_password", vv); err != nil {
				return fmt.Errorf("Error reading admin_ssh_password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_ssh_password: %v", err)
		}
	}

	if err = d.Set("admin_ssh_port", flattenSystemGlobalAdminSshPort(o["admin-ssh-port"], d, "admin_ssh_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-ssh-port"], "SystemGlobal-AdminSshPort"); ok {
			if err = d.Set("admin_ssh_port", vv); err != nil {
				return fmt.Errorf("Error reading admin_ssh_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_ssh_port: %v", err)
		}
	}

	if err = d.Set("admin_ssh_v1", flattenSystemGlobalAdminSshV1(o["admin-ssh-v1"], d, "admin_ssh_v1")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-ssh-v1"], "SystemGlobal-AdminSshV1"); ok {
			if err = d.Set("admin_ssh_v1", vv); err != nil {
				return fmt.Errorf("Error reading admin_ssh_v1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_ssh_v1: %v", err)
		}
	}

	if err = d.Set("admin_telnet", flattenSystemGlobalAdminTelnet(o["admin-telnet"], d, "admin_telnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-telnet"], "SystemGlobal-AdminTelnet"); ok {
			if err = d.Set("admin_telnet", vv); err != nil {
				return fmt.Errorf("Error reading admin_telnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_telnet: %v", err)
		}
	}

	if err = d.Set("admin_telnet_port", flattenSystemGlobalAdminTelnetPort(o["admin-telnet-port"], d, "admin_telnet_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-telnet-port"], "SystemGlobal-AdminTelnetPort"); ok {
			if err = d.Set("admin_telnet_port", vv); err != nil {
				return fmt.Errorf("Error reading admin_telnet_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_telnet_port: %v", err)
		}
	}

	if err = d.Set("admintimeout", flattenSystemGlobalAdmintimeout(o["admintimeout"], d, "admintimeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["admintimeout"], "SystemGlobal-Admintimeout"); ok {
			if err = d.Set("admintimeout", vv); err != nil {
				return fmt.Errorf("Error reading admintimeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admintimeout: %v", err)
		}
	}

	if err = d.Set("airplane_mode", flattenSystemGlobalAirplaneMode(o["airplane-mode"], d, "airplane_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["airplane-mode"], "SystemGlobal-AirplaneMode"); ok {
			if err = d.Set("airplane_mode", vv); err != nil {
				return fmt.Errorf("Error reading airplane_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading airplane_mode: %v", err)
		}
	}

	if err = d.Set("alias", flattenSystemGlobalAlias(o["alias"], d, "alias")); err != nil {
		if vv, ok := fortiAPIPatch(o["alias"], "SystemGlobal-Alias"); ok {
			if err = d.Set("alias", vv); err != nil {
				return fmt.Errorf("Error reading alias: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alias: %v", err)
		}
	}

	if err = d.Set("allow_traffic_redirect", flattenSystemGlobalAllowTrafficRedirect(o["allow-traffic-redirect"], d, "allow_traffic_redirect")); err != nil {
		if vv, ok := fortiAPIPatch(o["allow-traffic-redirect"], "SystemGlobal-AllowTrafficRedirect"); ok {
			if err = d.Set("allow_traffic_redirect", vv); err != nil {
				return fmt.Errorf("Error reading allow_traffic_redirect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allow_traffic_redirect: %v", err)
		}
	}

	if err = d.Set("anti_replay", flattenSystemGlobalAntiReplay(o["anti-replay"], d, "anti_replay")); err != nil {
		if vv, ok := fortiAPIPatch(o["anti-replay"], "SystemGlobal-AntiReplay"); ok {
			if err = d.Set("anti_replay", vv); err != nil {
				return fmt.Errorf("Error reading anti_replay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading anti_replay: %v", err)
		}
	}

	if err = d.Set("arp_max_entry", flattenSystemGlobalArpMaxEntry(o["arp-max-entry"], d, "arp_max_entry")); err != nil {
		if vv, ok := fortiAPIPatch(o["arp-max-entry"], "SystemGlobal-ArpMaxEntry"); ok {
			if err = d.Set("arp_max_entry", vv); err != nil {
				return fmt.Errorf("Error reading arp_max_entry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading arp_max_entry: %v", err)
		}
	}

	if err = d.Set("auth_cert", flattenSystemGlobalAuthCert(o["auth-cert"], d, "auth_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-cert"], "SystemGlobal-AuthCert"); ok {
			if err = d.Set("auth_cert", vv); err != nil {
				return fmt.Errorf("Error reading auth_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_cert: %v", err)
		}
	}

	if err = d.Set("auth_http_port", flattenSystemGlobalAuthHttpPort(o["auth-http-port"], d, "auth_http_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-http-port"], "SystemGlobal-AuthHttpPort"); ok {
			if err = d.Set("auth_http_port", vv); err != nil {
				return fmt.Errorf("Error reading auth_http_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_http_port: %v", err)
		}
	}

	if err = d.Set("auth_https_port", flattenSystemGlobalAuthHttpsPort(o["auth-https-port"], d, "auth_https_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-https-port"], "SystemGlobal-AuthHttpsPort"); ok {
			if err = d.Set("auth_https_port", vv); err != nil {
				return fmt.Errorf("Error reading auth_https_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_https_port: %v", err)
		}
	}

	if err = d.Set("auth_ike_saml_port", flattenSystemGlobalAuthIkeSamlPort(o["auth-ike-saml-port"], d, "auth_ike_saml_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-ike-saml-port"], "SystemGlobal-AuthIkeSamlPort"); ok {
			if err = d.Set("auth_ike_saml_port", vv); err != nil {
				return fmt.Errorf("Error reading auth_ike_saml_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_ike_saml_port: %v", err)
		}
	}

	if err = d.Set("auth_keepalive", flattenSystemGlobalAuthKeepalive(o["auth-keepalive"], d, "auth_keepalive")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-keepalive"], "SystemGlobal-AuthKeepalive"); ok {
			if err = d.Set("auth_keepalive", vv); err != nil {
				return fmt.Errorf("Error reading auth_keepalive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_keepalive: %v", err)
		}
	}

	if err = d.Set("auth_session_limit", flattenSystemGlobalAuthSessionLimit(o["auth-session-limit"], d, "auth_session_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-session-limit"], "SystemGlobal-AuthSessionLimit"); ok {
			if err = d.Set("auth_session_limit", vv); err != nil {
				return fmt.Errorf("Error reading auth_session_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_session_limit: %v", err)
		}
	}

	if err = d.Set("auto_auth_extension_device", flattenSystemGlobalAutoAuthExtensionDevice(o["auto-auth-extension-device"], d, "auto_auth_extension_device")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-auth-extension-device"], "SystemGlobal-AutoAuthExtensionDevice"); ok {
			if err = d.Set("auto_auth_extension_device", vv); err != nil {
				return fmt.Errorf("Error reading auto_auth_extension_device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_auth_extension_device: %v", err)
		}
	}

	if err = d.Set("autorun_log_fsck", flattenSystemGlobalAutorunLogFsck(o["autorun-log-fsck"], d, "autorun_log_fsck")); err != nil {
		if vv, ok := fortiAPIPatch(o["autorun-log-fsck"], "SystemGlobal-AutorunLogFsck"); ok {
			if err = d.Set("autorun_log_fsck", vv); err != nil {
				return fmt.Errorf("Error reading autorun_log_fsck: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading autorun_log_fsck: %v", err)
		}
	}

	if err = d.Set("av_affinity", flattenSystemGlobalAvAffinity(o["av-affinity"], d, "av_affinity")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-affinity"], "SystemGlobal-AvAffinity"); ok {
			if err = d.Set("av_affinity", vv); err != nil {
				return fmt.Errorf("Error reading av_affinity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_affinity: %v", err)
		}
	}

	if err = d.Set("av_failopen", flattenSystemGlobalAvFailopen(o["av-failopen"], d, "av_failopen")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-failopen"], "SystemGlobal-AvFailopen"); ok {
			if err = d.Set("av_failopen", vv); err != nil {
				return fmt.Errorf("Error reading av_failopen: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_failopen: %v", err)
		}
	}

	if err = d.Set("av_failopen_session", flattenSystemGlobalAvFailopenSession(o["av-failopen-session"], d, "av_failopen_session")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-failopen-session"], "SystemGlobal-AvFailopenSession"); ok {
			if err = d.Set("av_failopen_session", vv); err != nil {
				return fmt.Errorf("Error reading av_failopen_session: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_failopen_session: %v", err)
		}
	}

	if err = d.Set("batch_cmdb", flattenSystemGlobalBatchCmdb(o["batch-cmdb"], d, "batch_cmdb")); err != nil {
		if vv, ok := fortiAPIPatch(o["batch-cmdb"], "SystemGlobal-BatchCmdb"); ok {
			if err = d.Set("batch_cmdb", vv); err != nil {
				return fmt.Errorf("Error reading batch_cmdb: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading batch_cmdb: %v", err)
		}
	}

	if err = d.Set("bfd_affinity", flattenSystemGlobalBfdAffinity(o["bfd-affinity"], d, "bfd_affinity")); err != nil {
		if vv, ok := fortiAPIPatch(o["bfd-affinity"], "SystemGlobal-BfdAffinity"); ok {
			if err = d.Set("bfd_affinity", vv); err != nil {
				return fmt.Errorf("Error reading bfd_affinity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bfd_affinity: %v", err)
		}
	}

	if err = d.Set("block_session_timer", flattenSystemGlobalBlockSessionTimer(o["block-session-timer"], d, "block_session_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-session-timer"], "SystemGlobal-BlockSessionTimer"); ok {
			if err = d.Set("block_session_timer", vv); err != nil {
				return fmt.Errorf("Error reading block_session_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_session_timer: %v", err)
		}
	}

	if err = d.Set("br_fdb_max_entry", flattenSystemGlobalBrFdbMaxEntry(o["br-fdb-max-entry"], d, "br_fdb_max_entry")); err != nil {
		if vv, ok := fortiAPIPatch(o["br-fdb-max-entry"], "SystemGlobal-BrFdbMaxEntry"); ok {
			if err = d.Set("br_fdb_max_entry", vv); err != nil {
				return fmt.Errorf("Error reading br_fdb_max_entry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading br_fdb_max_entry: %v", err)
		}
	}

	if err = d.Set("cert_chain_max", flattenSystemGlobalCertChainMax(o["cert-chain-max"], d, "cert_chain_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert-chain-max"], "SystemGlobal-CertChainMax"); ok {
			if err = d.Set("cert_chain_max", vv); err != nil {
				return fmt.Errorf("Error reading cert_chain_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert_chain_max: %v", err)
		}
	}

	if err = d.Set("cfg_revert_timeout", flattenSystemGlobalCfgRevertTimeout(o["cfg-revert-timeout"], d, "cfg_revert_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["cfg-revert-timeout"], "SystemGlobal-CfgRevertTimeout"); ok {
			if err = d.Set("cfg_revert_timeout", vv); err != nil {
				return fmt.Errorf("Error reading cfg_revert_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cfg_revert_timeout: %v", err)
		}
	}

	if err = d.Set("cfg_save", flattenSystemGlobalCfgSave(o["cfg-save"], d, "cfg_save")); err != nil {
		if vv, ok := fortiAPIPatch(o["cfg-save"], "SystemGlobal-CfgSave"); ok {
			if err = d.Set("cfg_save", vv); err != nil {
				return fmt.Errorf("Error reading cfg_save: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cfg_save: %v", err)
		}
	}

	if err = d.Set("check_protocol_header", flattenSystemGlobalCheckProtocolHeader(o["check-protocol-header"], d, "check_protocol_header")); err != nil {
		if vv, ok := fortiAPIPatch(o["check-protocol-header"], "SystemGlobal-CheckProtocolHeader"); ok {
			if err = d.Set("check_protocol_header", vv); err != nil {
				return fmt.Errorf("Error reading check_protocol_header: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading check_protocol_header: %v", err)
		}
	}

	if err = d.Set("check_reset_range", flattenSystemGlobalCheckResetRange(o["check-reset-range"], d, "check_reset_range")); err != nil {
		if vv, ok := fortiAPIPatch(o["check-reset-range"], "SystemGlobal-CheckResetRange"); ok {
			if err = d.Set("check_reset_range", vv); err != nil {
				return fmt.Errorf("Error reading check_reset_range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading check_reset_range: %v", err)
		}
	}

	if err = d.Set("cli_audit_log", flattenSystemGlobalCliAuditLog(o["cli-audit-log"], d, "cli_audit_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["cli-audit-log"], "SystemGlobal-CliAuditLog"); ok {
			if err = d.Set("cli_audit_log", vv); err != nil {
				return fmt.Errorf("Error reading cli_audit_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cli_audit_log: %v", err)
		}
	}

	if err = d.Set("cloud_communication", flattenSystemGlobalCloudCommunication(o["cloud-communication"], d, "cloud_communication")); err != nil {
		if vv, ok := fortiAPIPatch(o["cloud-communication"], "SystemGlobal-CloudCommunication"); ok {
			if err = d.Set("cloud_communication", vv); err != nil {
				return fmt.Errorf("Error reading cloud_communication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cloud_communication: %v", err)
		}
	}

	if err = d.Set("clt_cert_req", flattenSystemGlobalCltCertReq(o["clt-cert-req"], d, "clt_cert_req")); err != nil {
		if vv, ok := fortiAPIPatch(o["clt-cert-req"], "SystemGlobal-CltCertReq"); ok {
			if err = d.Set("clt_cert_req", vv); err != nil {
				return fmt.Errorf("Error reading clt_cert_req: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading clt_cert_req: %v", err)
		}
	}

	if err = d.Set("cmdbsvr_affinity", flattenSystemGlobalCmdbsvrAffinity(o["cmdbsvr-affinity"], d, "cmdbsvr_affinity")); err != nil {
		if vv, ok := fortiAPIPatch(o["cmdbsvr-affinity"], "SystemGlobal-CmdbsvrAffinity"); ok {
			if err = d.Set("cmdbsvr_affinity", vv); err != nil {
				return fmt.Errorf("Error reading cmdbsvr_affinity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cmdbsvr_affinity: %v", err)
		}
	}

	if err = d.Set("cpu_use_threshold", flattenSystemGlobalCpuUseThreshold(o["cpu-use-threshold"], d, "cpu_use_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["cpu-use-threshold"], "SystemGlobal-CpuUseThreshold"); ok {
			if err = d.Set("cpu_use_threshold", vv); err != nil {
				return fmt.Errorf("Error reading cpu_use_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cpu_use_threshold: %v", err)
		}
	}

	if err = d.Set("csr_ca_attribute", flattenSystemGlobalCsrCaAttribute(o["csr-ca-attribute"], d, "csr_ca_attribute")); err != nil {
		if vv, ok := fortiAPIPatch(o["csr-ca-attribute"], "SystemGlobal-CsrCaAttribute"); ok {
			if err = d.Set("csr_ca_attribute", vv); err != nil {
				return fmt.Errorf("Error reading csr_ca_attribute: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading csr_ca_attribute: %v", err)
		}
	}

	if err = d.Set("daily_restart", flattenSystemGlobalDailyRestart(o["daily-restart"], d, "daily_restart")); err != nil {
		if vv, ok := fortiAPIPatch(o["daily-restart"], "SystemGlobal-DailyRestart"); ok {
			if err = d.Set("daily_restart", vv); err != nil {
				return fmt.Errorf("Error reading daily_restart: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading daily_restart: %v", err)
		}
	}

	if err = d.Set("default_service_source_port", flattenSystemGlobalDefaultServiceSourcePort(o["default-service-source-port"], d, "default_service_source_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-service-source-port"], "SystemGlobal-DefaultServiceSourcePort"); ok {
			if err = d.Set("default_service_source_port", vv); err != nil {
				return fmt.Errorf("Error reading default_service_source_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_service_source_port: %v", err)
		}
	}

	if err = d.Set("device_idle_timeout", flattenSystemGlobalDeviceIdleTimeout(o["device-idle-timeout"], d, "device_idle_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-idle-timeout"], "SystemGlobal-DeviceIdleTimeout"); ok {
			if err = d.Set("device_idle_timeout", vv); err != nil {
				return fmt.Errorf("Error reading device_idle_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_idle_timeout: %v", err)
		}
	}

	if err = d.Set("dh_params", flattenSystemGlobalDhParams(o["dh-params"], d, "dh_params")); err != nil {
		if vv, ok := fortiAPIPatch(o["dh-params"], "SystemGlobal-DhParams"); ok {
			if err = d.Set("dh_params", vv); err != nil {
				return fmt.Errorf("Error reading dh_params: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dh_params: %v", err)
		}
	}

	if err = d.Set("dhcp_lease_backup_interval", flattenSystemGlobalDhcpLeaseBackupInterval(o["dhcp-lease-backup-interval"], d, "dhcp_lease_backup_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-lease-backup-interval"], "SystemGlobal-DhcpLeaseBackupInterval"); ok {
			if err = d.Set("dhcp_lease_backup_interval", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_lease_backup_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_lease_backup_interval: %v", err)
		}
	}

	if err = d.Set("dnsproxy_worker_count", flattenSystemGlobalDnsproxyWorkerCount(o["dnsproxy-worker-count"], d, "dnsproxy_worker_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["dnsproxy-worker-count"], "SystemGlobal-DnsproxyWorkerCount"); ok {
			if err = d.Set("dnsproxy_worker_count", vv); err != nil {
				return fmt.Errorf("Error reading dnsproxy_worker_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dnsproxy_worker_count: %v", err)
		}
	}

	if err = d.Set("dp_fragment_timer", flattenSystemGlobalDpFragmentTimer(o["dp-fragment-timer"], d, "dp_fragment_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["dp-fragment-timer"], "SystemGlobal-DpFragmentTimer"); ok {
			if err = d.Set("dp_fragment_timer", vv); err != nil {
				return fmt.Errorf("Error reading dp_fragment_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dp_fragment_timer: %v", err)
		}
	}

	if err = d.Set("dp_pinhole_timer", flattenSystemGlobalDpPinholeTimer(o["dp-pinhole-timer"], d, "dp_pinhole_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["dp-pinhole-timer"], "SystemGlobal-DpPinholeTimer"); ok {
			if err = d.Set("dp_pinhole_timer", vv); err != nil {
				return fmt.Errorf("Error reading dp_pinhole_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dp_pinhole_timer: %v", err)
		}
	}

	if err = d.Set("dp_rsync_timer", flattenSystemGlobalDpRsyncTimer(o["dp-rsync-timer"], d, "dp_rsync_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["dp-rsync-timer"], "SystemGlobal-DpRsyncTimer"); ok {
			if err = d.Set("dp_rsync_timer", vv); err != nil {
				return fmt.Errorf("Error reading dp_rsync_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dp_rsync_timer: %v", err)
		}
	}

	if err = d.Set("dp_tcp_normal_timer", flattenSystemGlobalDpTcpNormalTimer(o["dp-tcp-normal-timer"], d, "dp_tcp_normal_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["dp-tcp-normal-timer"], "SystemGlobal-DpTcpNormalTimer"); ok {
			if err = d.Set("dp_tcp_normal_timer", vv); err != nil {
				return fmt.Errorf("Error reading dp_tcp_normal_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dp_tcp_normal_timer: %v", err)
		}
	}

	if err = d.Set("dp_udp_idle_timer", flattenSystemGlobalDpUdpIdleTimer(o["dp-udp-idle-timer"], d, "dp_udp_idle_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["dp-udp-idle-timer"], "SystemGlobal-DpUdpIdleTimer"); ok {
			if err = d.Set("dp_udp_idle_timer", vv); err != nil {
				return fmt.Errorf("Error reading dp_udp_idle_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dp_udp_idle_timer: %v", err)
		}
	}

	if err = d.Set("dst", flattenSystemGlobalDst(o["dst"], d, "dst")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst"], "SystemGlobal-Dst"); ok {
			if err = d.Set("dst", vv); err != nil {
				return fmt.Errorf("Error reading dst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	if err = d.Set("early_tcp_npu_session", flattenSystemGlobalEarlyTcpNpuSession(o["early-tcp-npu-session"], d, "early_tcp_npu_session")); err != nil {
		if vv, ok := fortiAPIPatch(o["early-tcp-npu-session"], "SystemGlobal-EarlyTcpNpuSession"); ok {
			if err = d.Set("early_tcp_npu_session", vv); err != nil {
				return fmt.Errorf("Error reading early_tcp_npu_session: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading early_tcp_npu_session: %v", err)
		}
	}

	if err = d.Set("endpoint_control_fds_access", flattenSystemGlobalEndpointControlFdsAccess(o["endpoint-control-fds-access"], d, "endpoint_control_fds_access")); err != nil {
		if vv, ok := fortiAPIPatch(o["endpoint-control-fds-access"], "SystemGlobal-EndpointControlFdsAccess"); ok {
			if err = d.Set("endpoint_control_fds_access", vv); err != nil {
				return fmt.Errorf("Error reading endpoint_control_fds_access: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading endpoint_control_fds_access: %v", err)
		}
	}

	if err = d.Set("edit_vdom_prompt", flattenSystemGlobalEditVdomPrompt(o["edit-vdom-prompt"], d, "edit_vdom_prompt")); err != nil {
		if vv, ok := fortiAPIPatch(o["edit-vdom-prompt"], "SystemGlobal-EditVdomPrompt"); ok {
			if err = d.Set("edit_vdom_prompt", vv); err != nil {
				return fmt.Errorf("Error reading edit_vdom_prompt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading edit_vdom_prompt: %v", err)
		}
	}

	if err = d.Set("extender_controller_reserved_network", flattenSystemGlobalExtenderControllerReservedNetwork(o["extender-controller-reserved-network"], d, "extender_controller_reserved_network")); err != nil {
		if vv, ok := fortiAPIPatch(o["extender-controller-reserved-network"], "SystemGlobal-ExtenderControllerReservedNetwork"); ok {
			if err = d.Set("extender_controller_reserved_network", vv); err != nil {
				return fmt.Errorf("Error reading extender_controller_reserved_network: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extender_controller_reserved_network: %v", err)
		}
	}

	if err = d.Set("faz_disk_buffer_size", flattenSystemGlobalFazDiskBufferSize(o["faz-disk-buffer-size"], d, "faz_disk_buffer_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz-disk-buffer-size"], "SystemGlobal-FazDiskBufferSize"); ok {
			if err = d.Set("faz_disk_buffer_size", vv); err != nil {
				return fmt.Errorf("Error reading faz_disk_buffer_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading faz_disk_buffer_size: %v", err)
		}
	}

	if err = d.Set("fds_statistics", flattenSystemGlobalFdsStatistics(o["fds-statistics"], d, "fds_statistics")); err != nil {
		if vv, ok := fortiAPIPatch(o["fds-statistics"], "SystemGlobal-FdsStatistics"); ok {
			if err = d.Set("fds_statistics", vv); err != nil {
				return fmt.Errorf("Error reading fds_statistics: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fds_statistics: %v", err)
		}
	}

	if err = d.Set("fds_statistics_period", flattenSystemGlobalFdsStatisticsPeriod(o["fds-statistics-period"], d, "fds_statistics_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["fds-statistics-period"], "SystemGlobal-FdsStatisticsPeriod"); ok {
			if err = d.Set("fds_statistics_period", vv); err != nil {
				return fmt.Errorf("Error reading fds_statistics_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fds_statistics_period: %v", err)
		}
	}

	if err = d.Set("fec_port", flattenSystemGlobalFecPort(o["fec-port"], d, "fec_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["fec-port"], "SystemGlobal-FecPort"); ok {
			if err = d.Set("fec_port", vv); err != nil {
				return fmt.Errorf("Error reading fec_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fec_port: %v", err)
		}
	}

	if err = d.Set("fgd_alert_subscription", flattenSystemGlobalFgdAlertSubscription(o["fgd-alert-subscription"], d, "fgd_alert_subscription")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgd-alert-subscription"], "SystemGlobal-FgdAlertSubscription"); ok {
			if err = d.Set("fgd_alert_subscription", vv); err != nil {
				return fmt.Errorf("Error reading fgd_alert_subscription: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgd_alert_subscription: %v", err)
		}
	}

	if err = d.Set("forticarrier_bypass", flattenSystemGlobalForticarrierBypass(o["forticarrier-bypass"], d, "forticarrier_bypass")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticarrier-bypass"], "SystemGlobal-ForticarrierBypass"); ok {
			if err = d.Set("forticarrier_bypass", vv); err != nil {
				return fmt.Errorf("Error reading forticarrier_bypass: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticarrier_bypass: %v", err)
		}
	}

	if err = d.Set("forticontroller_proxy", flattenSystemGlobalForticontrollerProxy(o["forticontroller-proxy"], d, "forticontroller_proxy")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticontroller-proxy"], "SystemGlobal-ForticontrollerProxy"); ok {
			if err = d.Set("forticontroller_proxy", vv); err != nil {
				return fmt.Errorf("Error reading forticontroller_proxy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticontroller_proxy: %v", err)
		}
	}

	if err = d.Set("forticontroller_proxy_port", flattenSystemGlobalForticontrollerProxyPort(o["forticontroller-proxy-port"], d, "forticontroller_proxy_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticontroller-proxy-port"], "SystemGlobal-ForticontrollerProxyPort"); ok {
			if err = d.Set("forticontroller_proxy_port", vv); err != nil {
				return fmt.Errorf("Error reading forticontroller_proxy_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticontroller_proxy_port: %v", err)
		}
	}

	if err = d.Set("forticonverter_config_upload", flattenSystemGlobalForticonverterConfigUpload(o["forticonverter-config-upload"], d, "forticonverter_config_upload")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticonverter-config-upload"], "SystemGlobal-ForticonverterConfigUpload"); ok {
			if err = d.Set("forticonverter_config_upload", vv); err != nil {
				return fmt.Errorf("Error reading forticonverter_config_upload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticonverter_config_upload: %v", err)
		}
	}

	if err = d.Set("forticonverter_integration", flattenSystemGlobalForticonverterIntegration(o["forticonverter-integration"], d, "forticonverter_integration")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticonverter-integration"], "SystemGlobal-ForticonverterIntegration"); ok {
			if err = d.Set("forticonverter_integration", vv); err != nil {
				return fmt.Errorf("Error reading forticonverter_integration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticonverter_integration: %v", err)
		}
	}

	if err = d.Set("fortiextender", flattenSystemGlobalFortiextender(o["fortiextender"], d, "fortiextender")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiextender"], "SystemGlobal-Fortiextender"); ok {
			if err = d.Set("fortiextender", vv); err != nil {
				return fmt.Errorf("Error reading fortiextender: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiextender: %v", err)
		}
	}

	if err = d.Set("fortiextender_data_port", flattenSystemGlobalFortiextenderDataPort(o["fortiextender-data-port"], d, "fortiextender_data_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiextender-data-port"], "SystemGlobal-FortiextenderDataPort"); ok {
			if err = d.Set("fortiextender_data_port", vv); err != nil {
				return fmt.Errorf("Error reading fortiextender_data_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiextender_data_port: %v", err)
		}
	}

	if err = d.Set("fortiextender_discovery_lockdown", flattenSystemGlobalFortiextenderDiscoveryLockdown(o["fortiextender-discovery-lockdown"], d, "fortiextender_discovery_lockdown")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiextender-discovery-lockdown"], "SystemGlobal-FortiextenderDiscoveryLockdown"); ok {
			if err = d.Set("fortiextender_discovery_lockdown", vv); err != nil {
				return fmt.Errorf("Error reading fortiextender_discovery_lockdown: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiextender_discovery_lockdown: %v", err)
		}
	}

	if err = d.Set("fortiextender_provision_on_authorization", flattenSystemGlobalFortiextenderProvisionOnAuthorization(o["fortiextender-provision-on-authorization"], d, "fortiextender_provision_on_authorization")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiextender-provision-on-authorization"], "SystemGlobal-FortiextenderProvisionOnAuthorization"); ok {
			if err = d.Set("fortiextender_provision_on_authorization", vv); err != nil {
				return fmt.Errorf("Error reading fortiextender_provision_on_authorization: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiextender_provision_on_authorization: %v", err)
		}
	}

	if err = d.Set("fortiextender_vlan_mode", flattenSystemGlobalFortiextenderVlanMode(o["fortiextender-vlan-mode"], d, "fortiextender_vlan_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiextender-vlan-mode"], "SystemGlobal-FortiextenderVlanMode"); ok {
			if err = d.Set("fortiextender_vlan_mode", vv); err != nil {
				return fmt.Errorf("Error reading fortiextender_vlan_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiextender_vlan_mode: %v", err)
		}
	}

	if err = d.Set("fortiipam_integration", flattenSystemGlobalFortiipamIntegration(o["fortiipam-integration"], d, "fortiipam_integration")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiipam-integration"], "SystemGlobal-FortiipamIntegration"); ok {
			if err = d.Set("fortiipam_integration", vv); err != nil {
				return fmt.Errorf("Error reading fortiipam_integration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiipam_integration: %v", err)
		}
	}

	if err = d.Set("fortigslb_integration", flattenSystemGlobalFortigslbIntegration(o["fortigslb-integration"], d, "fortigslb_integration")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortigslb-integration"], "SystemGlobal-FortigslbIntegration"); ok {
			if err = d.Set("fortigslb_integration", vv); err != nil {
				return fmt.Errorf("Error reading fortigslb_integration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortigslb_integration: %v", err)
		}
	}

	if err = d.Set("fortiservice_port", flattenSystemGlobalFortiservicePort(o["fortiservice-port"], d, "fortiservice_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiservice-port"], "SystemGlobal-FortiservicePort"); ok {
			if err = d.Set("fortiservice_port", vv); err != nil {
				return fmt.Errorf("Error reading fortiservice_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiservice_port: %v", err)
		}
	}

	if err = d.Set("fortitoken_cloud", flattenSystemGlobalFortitokenCloud(o["fortitoken-cloud"], d, "fortitoken_cloud")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortitoken-cloud"], "SystemGlobal-FortitokenCloud"); ok {
			if err = d.Set("fortitoken_cloud", vv); err != nil {
				return fmt.Errorf("Error reading fortitoken_cloud: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortitoken_cloud: %v", err)
		}
	}

	if err = d.Set("fortitoken_cloud_service", flattenSystemGlobalFortitokenCloudService(o["fortitoken-cloud-service"], d, "fortitoken_cloud_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortitoken-cloud-service"], "SystemGlobal-FortitokenCloudService"); ok {
			if err = d.Set("fortitoken_cloud_service", vv); err != nil {
				return fmt.Errorf("Error reading fortitoken_cloud_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortitoken_cloud_service: %v", err)
		}
	}

	if err = d.Set("gui_allow_default_hostname", flattenSystemGlobalGuiAllowDefaultHostname(o["gui-allow-default-hostname"], d, "gui_allow_default_hostname")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-allow-default-hostname"], "SystemGlobal-GuiAllowDefaultHostname"); ok {
			if err = d.Set("gui_allow_default_hostname", vv); err != nil {
				return fmt.Errorf("Error reading gui_allow_default_hostname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_allow_default_hostname: %v", err)
		}
	}

	if err = d.Set("fortitoken_cloud_push_status", flattenSystemGlobalFortitokenCloudPushStatus(o["fortitoken-cloud-push-status"], d, "fortitoken_cloud_push_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortitoken-cloud-push-status"], "SystemGlobal-FortitokenCloudPushStatus"); ok {
			if err = d.Set("fortitoken_cloud_push_status", vv); err != nil {
				return fmt.Errorf("Error reading fortitoken_cloud_push_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortitoken_cloud_push_status: %v", err)
		}
	}

	if err = d.Set("fortitoken_cloud_sync_interval", flattenSystemGlobalFortitokenCloudSyncInterval(o["fortitoken-cloud-sync-interval"], d, "fortitoken_cloud_sync_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortitoken-cloud-sync-interval"], "SystemGlobal-FortitokenCloudSyncInterval"); ok {
			if err = d.Set("fortitoken_cloud_sync_interval", vv); err != nil {
				return fmt.Errorf("Error reading fortitoken_cloud_sync_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortitoken_cloud_sync_interval: %v", err)
		}
	}

	if err = d.Set("gui_allow_incompatible_fabric_fgt", flattenSystemGlobalGuiAllowIncompatibleFabricFgt(o["gui-allow-incompatible-fabric-fgt"], d, "gui_allow_incompatible_fabric_fgt")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-allow-incompatible-fabric-fgt"], "SystemGlobal-GuiAllowIncompatibleFabricFgt"); ok {
			if err = d.Set("gui_allow_incompatible_fabric_fgt", vv); err != nil {
				return fmt.Errorf("Error reading gui_allow_incompatible_fabric_fgt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_allow_incompatible_fabric_fgt: %v", err)
		}
	}

	if err = d.Set("gui_app_detection_sdwan", flattenSystemGlobalGuiAppDetectionSdwan(o["gui-app-detection-sdwan"], d, "gui_app_detection_sdwan")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-app-detection-sdwan"], "SystemGlobal-GuiAppDetectionSdwan"); ok {
			if err = d.Set("gui_app_detection_sdwan", vv); err != nil {
				return fmt.Errorf("Error reading gui_app_detection_sdwan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_app_detection_sdwan: %v", err)
		}
	}

	if err = d.Set("gui_auto_upgrade_setup_warning", flattenSystemGlobalGuiAutoUpgradeSetupWarning(o["gui-auto-upgrade-setup-warning"], d, "gui_auto_upgrade_setup_warning")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-auto-upgrade-setup-warning"], "SystemGlobal-GuiAutoUpgradeSetupWarning"); ok {
			if err = d.Set("gui_auto_upgrade_setup_warning", vv); err != nil {
				return fmt.Errorf("Error reading gui_auto_upgrade_setup_warning: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_auto_upgrade_setup_warning: %v", err)
		}
	}

	if err = d.Set("gui_cdn_domain_override", flattenSystemGlobalGuiCdnDomainOverride(o["gui-cdn-domain-override"], d, "gui_cdn_domain_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-cdn-domain-override"], "SystemGlobal-GuiCdnDomainOverride"); ok {
			if err = d.Set("gui_cdn_domain_override", vv); err != nil {
				return fmt.Errorf("Error reading gui_cdn_domain_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_cdn_domain_override: %v", err)
		}
	}

	if err = d.Set("gui_cdn_usage", flattenSystemGlobalGuiCdnUsage(o["gui-cdn-usage"], d, "gui_cdn_usage")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-cdn-usage"], "SystemGlobal-GuiCdnUsage"); ok {
			if err = d.Set("gui_cdn_usage", vv); err != nil {
				return fmt.Errorf("Error reading gui_cdn_usage: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_cdn_usage: %v", err)
		}
	}

	if err = d.Set("gui_certificates", flattenSystemGlobalGuiCertificates(o["gui-certificates"], d, "gui_certificates")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-certificates"], "SystemGlobal-GuiCertificates"); ok {
			if err = d.Set("gui_certificates", vv); err != nil {
				return fmt.Errorf("Error reading gui_certificates: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_certificates: %v", err)
		}
	}

	if err = d.Set("gui_custom_language", flattenSystemGlobalGuiCustomLanguage(o["gui-custom-language"], d, "gui_custom_language")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-custom-language"], "SystemGlobal-GuiCustomLanguage"); ok {
			if err = d.Set("gui_custom_language", vv); err != nil {
				return fmt.Errorf("Error reading gui_custom_language: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_custom_language: %v", err)
		}
	}

	if err = d.Set("gui_date_format", flattenSystemGlobalGuiDateFormat(o["gui-date-format"], d, "gui_date_format")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-date-format"], "SystemGlobal-GuiDateFormat"); ok {
			if err = d.Set("gui_date_format", vv); err != nil {
				return fmt.Errorf("Error reading gui_date_format: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_date_format: %v", err)
		}
	}

	if err = d.Set("gui_date_time_source", flattenSystemGlobalGuiDateTimeSource(o["gui-date-time-source"], d, "gui_date_time_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-date-time-source"], "SystemGlobal-GuiDateTimeSource"); ok {
			if err = d.Set("gui_date_time_source", vv); err != nil {
				return fmt.Errorf("Error reading gui_date_time_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_date_time_source: %v", err)
		}
	}

	if err = d.Set("gui_device_latitude", flattenSystemGlobalGuiDeviceLatitude(o["gui-device-latitude"], d, "gui_device_latitude")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-device-latitude"], "SystemGlobal-GuiDeviceLatitude"); ok {
			if err = d.Set("gui_device_latitude", vv); err != nil {
				return fmt.Errorf("Error reading gui_device_latitude: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_device_latitude: %v", err)
		}
	}

	if err = d.Set("gui_device_longitude", flattenSystemGlobalGuiDeviceLongitude(o["gui-device-longitude"], d, "gui_device_longitude")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-device-longitude"], "SystemGlobal-GuiDeviceLongitude"); ok {
			if err = d.Set("gui_device_longitude", vv); err != nil {
				return fmt.Errorf("Error reading gui_device_longitude: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_device_longitude: %v", err)
		}
	}

	if err = d.Set("gui_display_hostname", flattenSystemGlobalGuiDisplayHostname(o["gui-display-hostname"], d, "gui_display_hostname")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-display-hostname"], "SystemGlobal-GuiDisplayHostname"); ok {
			if err = d.Set("gui_display_hostname", vv); err != nil {
				return fmt.Errorf("Error reading gui_display_hostname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_display_hostname: %v", err)
		}
	}

	if err = d.Set("gui_firmware_upgrade_setup_warning", flattenSystemGlobalGuiFirmwareUpgradeSetupWarning(o["gui-firmware-upgrade-setup-warning"], d, "gui_firmware_upgrade_setup_warning")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-firmware-upgrade-setup-warning"], "SystemGlobal-GuiFirmwareUpgradeSetupWarning"); ok {
			if err = d.Set("gui_firmware_upgrade_setup_warning", vv); err != nil {
				return fmt.Errorf("Error reading gui_firmware_upgrade_setup_warning: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_firmware_upgrade_setup_warning: %v", err)
		}
	}

	if err = d.Set("gui_firmware_upgrade_warning", flattenSystemGlobalGuiFirmwareUpgradeWarning(o["gui-firmware-upgrade-warning"], d, "gui_firmware_upgrade_warning")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-firmware-upgrade-warning"], "SystemGlobal-GuiFirmwareUpgradeWarning"); ok {
			if err = d.Set("gui_firmware_upgrade_warning", vv); err != nil {
				return fmt.Errorf("Error reading gui_firmware_upgrade_warning: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_firmware_upgrade_warning: %v", err)
		}
	}

	if err = d.Set("gui_forticare_registration_setup_warning", flattenSystemGlobalGuiForticareRegistrationSetupWarning(o["gui-forticare-registration-setup-warning"], d, "gui_forticare_registration_setup_warning")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-forticare-registration-setup-warning"], "SystemGlobal-GuiForticareRegistrationSetupWarning"); ok {
			if err = d.Set("gui_forticare_registration_setup_warning", vv); err != nil {
				return fmt.Errorf("Error reading gui_forticare_registration_setup_warning: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_forticare_registration_setup_warning: %v", err)
		}
	}

	if err = d.Set("gui_fortisandbox_cloud", flattenSystemGlobalGuiFortisandboxCloud(o["gui-fortisandbox-cloud"], d, "gui_fortisandbox_cloud")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-fortisandbox-cloud"], "SystemGlobal-GuiFortisandboxCloud"); ok {
			if err = d.Set("gui_fortisandbox_cloud", vv); err != nil {
				return fmt.Errorf("Error reading gui_fortisandbox_cloud: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_fortisandbox_cloud: %v", err)
		}
	}

	if err = d.Set("gui_fortigate_cloud_sandbox", flattenSystemGlobalGuiFortigateCloudSandbox(o["gui-fortigate-cloud-sandbox"], d, "gui_fortigate_cloud_sandbox")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-fortigate-cloud-sandbox"], "SystemGlobal-GuiFortigateCloudSandbox"); ok {
			if err = d.Set("gui_fortigate_cloud_sandbox", vv); err != nil {
				return fmt.Errorf("Error reading gui_fortigate_cloud_sandbox: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_fortigate_cloud_sandbox: %v", err)
		}
	}

	if err = d.Set("gui_fortiguard_resource_fetch", flattenSystemGlobalGuiFortiguardResourceFetch(o["gui-fortiguard-resource-fetch"], d, "gui_fortiguard_resource_fetch")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-fortiguard-resource-fetch"], "SystemGlobal-GuiFortiguardResourceFetch"); ok {
			if err = d.Set("gui_fortiguard_resource_fetch", vv); err != nil {
				return fmt.Errorf("Error reading gui_fortiguard_resource_fetch: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_fortiguard_resource_fetch: %v", err)
		}
	}

	if err = d.Set("gui_ipv6", flattenSystemGlobalGuiIpv6(o["gui-ipv6"], d, "gui_ipv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-ipv6"], "SystemGlobal-GuiIpv6"); ok {
			if err = d.Set("gui_ipv6", vv); err != nil {
				return fmt.Errorf("Error reading gui_ipv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_ipv6: %v", err)
		}
	}

	if err = d.Set("gui_lines_per_page", flattenSystemGlobalGuiLinesPerPage(o["gui-lines-per-page"], d, "gui_lines_per_page")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-lines-per-page"], "SystemGlobal-GuiLinesPerPage"); ok {
			if err = d.Set("gui_lines_per_page", vv); err != nil {
				return fmt.Errorf("Error reading gui_lines_per_page: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_lines_per_page: %v", err)
		}
	}

	if err = d.Set("gui_local_out", flattenSystemGlobalGuiLocalOut(o["gui-local-out"], d, "gui_local_out")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-local-out"], "SystemGlobal-GuiLocalOut"); ok {
			if err = d.Set("gui_local_out", vv); err != nil {
				return fmt.Errorf("Error reading gui_local_out: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_local_out: %v", err)
		}
	}

	if err = d.Set("gui_replacement_message_groups", flattenSystemGlobalGuiReplacementMessageGroups(o["gui-replacement-message-groups"], d, "gui_replacement_message_groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-replacement-message-groups"], "SystemGlobal-GuiReplacementMessageGroups"); ok {
			if err = d.Set("gui_replacement_message_groups", vv); err != nil {
				return fmt.Errorf("Error reading gui_replacement_message_groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_replacement_message_groups: %v", err)
		}
	}

	if err = d.Set("gui_rest_api_cache", flattenSystemGlobalGuiRestApiCache(o["gui-rest-api-cache"], d, "gui_rest_api_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-rest-api-cache"], "SystemGlobal-GuiRestApiCache"); ok {
			if err = d.Set("gui_rest_api_cache", vv); err != nil {
				return fmt.Errorf("Error reading gui_rest_api_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_rest_api_cache: %v", err)
		}
	}

	if err = d.Set("gui_theme", flattenSystemGlobalGuiTheme(o["gui-theme"], d, "gui_theme")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-theme"], "SystemGlobal-GuiTheme"); ok {
			if err = d.Set("gui_theme", vv); err != nil {
				return fmt.Errorf("Error reading gui_theme: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_theme: %v", err)
		}
	}

	if err = d.Set("gui_wireless_opensecurity", flattenSystemGlobalGuiWirelessOpensecurity(o["gui-wireless-opensecurity"], d, "gui_wireless_opensecurity")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-wireless-opensecurity"], "SystemGlobal-GuiWirelessOpensecurity"); ok {
			if err = d.Set("gui_wireless_opensecurity", vv); err != nil {
				return fmt.Errorf("Error reading gui_wireless_opensecurity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_wireless_opensecurity: %v", err)
		}
	}

	if err = d.Set("gui_workflow_management", flattenSystemGlobalGuiWorkflowManagement(o["gui-workflow-management"], d, "gui_workflow_management")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-workflow-management"], "SystemGlobal-GuiWorkflowManagement"); ok {
			if err = d.Set("gui_workflow_management", vv); err != nil {
				return fmt.Errorf("Error reading gui_workflow_management: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_workflow_management: %v", err)
		}
	}

	if err = d.Set("ha_affinity", flattenSystemGlobalHaAffinity(o["ha-affinity"], d, "ha_affinity")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-affinity"], "SystemGlobal-HaAffinity"); ok {
			if err = d.Set("ha_affinity", vv); err != nil {
				return fmt.Errorf("Error reading ha_affinity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_affinity: %v", err)
		}
	}

	if err = d.Set("honor_df", flattenSystemGlobalHonorDf(o["honor-df"], d, "honor_df")); err != nil {
		if vv, ok := fortiAPIPatch(o["honor-df"], "SystemGlobal-HonorDf"); ok {
			if err = d.Set("honor_df", vv); err != nil {
				return fmt.Errorf("Error reading honor_df: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading honor_df: %v", err)
		}
	}

	if err = d.Set("hostname", flattenSystemGlobalHostname(o["hostname"], d, "hostname")); err != nil {
		if vv, ok := fortiAPIPatch(o["hostname"], "SystemGlobal-Hostname"); ok {
			if err = d.Set("hostname", vv); err != nil {
				return fmt.Errorf("Error reading hostname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("hw_switch_ether_filter", flattenSystemGlobalHwSwitchEtherFilter(o["hw-switch-ether-filter"], d, "hw_switch_ether_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["hw-switch-ether-filter"], "SystemGlobal-HwSwitchEtherFilter"); ok {
			if err = d.Set("hw_switch_ether_filter", vv); err != nil {
				return fmt.Errorf("Error reading hw_switch_ether_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hw_switch_ether_filter: %v", err)
		}
	}

	if err = d.Set("http_request_limit", flattenSystemGlobalHttpRequestLimit(o["http-request-limit"], d, "http_request_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-request-limit"], "SystemGlobal-HttpRequestLimit"); ok {
			if err = d.Set("http_request_limit", vv); err != nil {
				return fmt.Errorf("Error reading http_request_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_request_limit: %v", err)
		}
	}

	if err = d.Set("http_unauthenticated_request_limit", flattenSystemGlobalHttpUnauthenticatedRequestLimit(o["http-unauthenticated-request-limit"], d, "http_unauthenticated_request_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-unauthenticated-request-limit"], "SystemGlobal-HttpUnauthenticatedRequestLimit"); ok {
			if err = d.Set("http_unauthenticated_request_limit", vv); err != nil {
				return fmt.Errorf("Error reading http_unauthenticated_request_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_unauthenticated_request_limit: %v", err)
		}
	}

	if err = d.Set("hyper_scale_vdom_num", flattenSystemGlobalHyperScaleVdomNum(o["hyper-scale-vdom-num"], d, "hyper_scale_vdom_num")); err != nil {
		if vv, ok := fortiAPIPatch(o["hyper-scale-vdom-num"], "SystemGlobal-HyperScaleVdomNum"); ok {
			if err = d.Set("hyper_scale_vdom_num", vv); err != nil {
				return fmt.Errorf("Error reading hyper_scale_vdom_num: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hyper_scale_vdom_num: %v", err)
		}
	}

	if err = d.Set("igmp_state_limit", flattenSystemGlobalIgmpStateLimit(o["igmp-state-limit"], d, "igmp_state_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["igmp-state-limit"], "SystemGlobal-IgmpStateLimit"); ok {
			if err = d.Set("igmp_state_limit", vv); err != nil {
				return fmt.Errorf("Error reading igmp_state_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading igmp_state_limit: %v", err)
		}
	}

	if err = d.Set("interface_subnet_usage", flattenSystemGlobalInterfaceSubnetUsage(o["interface-subnet-usage"], d, "interface_subnet_usage")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-subnet-usage"], "SystemGlobal-InterfaceSubnetUsage"); ok {
			if err = d.Set("interface_subnet_usage", vv); err != nil {
				return fmt.Errorf("Error reading interface_subnet_usage: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_subnet_usage: %v", err)
		}
	}

	if err = d.Set("internal_switch_mode", flattenSystemGlobalInternalSwitchMode(o["internal-switch-mode"], d, "internal_switch_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["internal-switch-mode"], "SystemGlobal-InternalSwitchMode"); ok {
			if err = d.Set("internal_switch_mode", vv); err != nil {
				return fmt.Errorf("Error reading internal_switch_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internal_switch_mode: %v", err)
		}
	}

	if err = d.Set("internal_switch_speed", flattenSystemGlobalInternalSwitchSpeed(o["internal-switch-speed"], d, "internal_switch_speed")); err != nil {
		if vv, ok := fortiAPIPatch(o["internal-switch-speed"], "SystemGlobal-InternalSwitchSpeed"); ok {
			if err = d.Set("internal_switch_speed", vv); err != nil {
				return fmt.Errorf("Error reading internal_switch_speed: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internal_switch_speed: %v", err)
		}
	}

	if err = d.Set("internet_service_database", flattenSystemGlobalInternetServiceDatabase(o["internet-service-database"], d, "internet_service_database")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-database"], "SystemGlobal-InternetServiceDatabase"); ok {
			if err = d.Set("internet_service_database", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_database: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_database: %v", err)
		}
	}

	if err = d.Set("internet_service_download_list", flattenSystemGlobalInternetServiceDownloadList(o["internet-service-download-list"], d, "internet_service_download_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-download-list"], "SystemGlobal-InternetServiceDownloadList"); ok {
			if err = d.Set("internet_service_download_list", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_download_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_download_list: %v", err)
		}
	}

	if err = d.Set("ip_fragment_mem_thresholds", flattenSystemGlobalIpFragmentMemThresholds(o["ip-fragment-mem-thresholds"], d, "ip_fragment_mem_thresholds")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-fragment-mem-thresholds"], "SystemGlobal-IpFragmentMemThresholds"); ok {
			if err = d.Set("ip_fragment_mem_thresholds", vv); err != nil {
				return fmt.Errorf("Error reading ip_fragment_mem_thresholds: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_fragment_mem_thresholds: %v", err)
		}
	}

	if err = d.Set("ip_src_port_range", flattenSystemGlobalIpSrcPortRange(o["ip-src-port-range"], d, "ip_src_port_range")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-src-port-range"], "SystemGlobal-IpSrcPortRange"); ok {
			if err = d.Set("ip_src_port_range", vv); err != nil {
				return fmt.Errorf("Error reading ip_src_port_range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_src_port_range: %v", err)
		}
	}

	if err = d.Set("ips_affinity", flattenSystemGlobalIpsAffinity(o["ips-affinity"], d, "ips_affinity")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-affinity"], "SystemGlobal-IpsAffinity"); ok {
			if err = d.Set("ips_affinity", vv); err != nil {
				return fmt.Errorf("Error reading ips_affinity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_affinity: %v", err)
		}
	}

	if err = d.Set("ipsec_asic_offload", flattenSystemGlobalIpsecAsicOffload(o["ipsec-asic-offload"], d, "ipsec_asic_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-asic-offload"], "SystemGlobal-IpsecAsicOffload"); ok {
			if err = d.Set("ipsec_asic_offload", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_asic_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_asic_offload: %v", err)
		}
	}

	if err = d.Set("ipsec_ha_seqjump_rate", flattenSystemGlobalIpsecHaSeqjumpRate(o["ipsec-ha-seqjump-rate"], d, "ipsec_ha_seqjump_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-ha-seqjump-rate"], "SystemGlobal-IpsecHaSeqjumpRate"); ok {
			if err = d.Set("ipsec_ha_seqjump_rate", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_ha_seqjump_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_ha_seqjump_rate: %v", err)
		}
	}

	if err = d.Set("ipsec_hmac_offload", flattenSystemGlobalIpsecHmacOffload(o["ipsec-hmac-offload"], d, "ipsec_hmac_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-hmac-offload"], "SystemGlobal-IpsecHmacOffload"); ok {
			if err = d.Set("ipsec_hmac_offload", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_hmac_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_hmac_offload: %v", err)
		}
	}

	if err = d.Set("ipsec_qat_offload", flattenSystemGlobalIpsecQatOffload(o["ipsec-qat-offload"], d, "ipsec_qat_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-qat-offload"], "SystemGlobal-IpsecQatOffload"); ok {
			if err = d.Set("ipsec_qat_offload", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_qat_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_qat_offload: %v", err)
		}
	}

	if err = d.Set("ipsec_round_robin", flattenSystemGlobalIpsecRoundRobin(o["ipsec-round-robin"], d, "ipsec_round_robin")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-round-robin"], "SystemGlobal-IpsecRoundRobin"); ok {
			if err = d.Set("ipsec_round_robin", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_round_robin: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_round_robin: %v", err)
		}
	}

	if err = d.Set("ipsec_soft_dec_async", flattenSystemGlobalIpsecSoftDecAsync(o["ipsec-soft-dec-async"], d, "ipsec_soft_dec_async")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-soft-dec-async"], "SystemGlobal-IpsecSoftDecAsync"); ok {
			if err = d.Set("ipsec_soft_dec_async", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_soft_dec_async: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_soft_dec_async: %v", err)
		}
	}

	if err = d.Set("ipv6_accept_dad", flattenSystemGlobalIpv6AcceptDad(o["ipv6-accept-dad"], d, "ipv6_accept_dad")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-accept-dad"], "SystemGlobal-Ipv6AcceptDad"); ok {
			if err = d.Set("ipv6_accept_dad", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_accept_dad: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_accept_dad: %v", err)
		}
	}

	if err = d.Set("ipv6_allow_anycast_probe", flattenSystemGlobalIpv6AllowAnycastProbe(o["ipv6-allow-anycast-probe"], d, "ipv6_allow_anycast_probe")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-allow-anycast-probe"], "SystemGlobal-Ipv6AllowAnycastProbe"); ok {
			if err = d.Set("ipv6_allow_anycast_probe", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_allow_anycast_probe: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_allow_anycast_probe: %v", err)
		}
	}

	if err = d.Set("ipv6_allow_local_in_silent_drop", flattenSystemGlobalIpv6AllowLocalInSilentDrop(o["ipv6-allow-local-in-silent-drop"], d, "ipv6_allow_local_in_silent_drop")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-allow-local-in-silent-drop"], "SystemGlobal-Ipv6AllowLocalInSilentDrop"); ok {
			if err = d.Set("ipv6_allow_local_in_silent_drop", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_allow_local_in_silent_drop: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_allow_local_in_silent_drop: %v", err)
		}
	}

	if err = d.Set("ipv6_allow_local_in_slient_drop", flattenSystemGlobalIpv6AllowLocalInSlientDrop(o["ipv6-allow-local-in-slient-drop"], d, "ipv6_allow_local_in_slient_drop")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-allow-local-in-slient-drop"], "SystemGlobal-Ipv6AllowLocalInSlientDrop"); ok {
			if err = d.Set("ipv6_allow_local_in_slient_drop", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_allow_local_in_slient_drop: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_allow_local_in_slient_drop: %v", err)
		}
	}

	if err = d.Set("ipv6_allow_multicast_probe", flattenSystemGlobalIpv6AllowMulticastProbe(o["ipv6-allow-multicast-probe"], d, "ipv6_allow_multicast_probe")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-allow-multicast-probe"], "SystemGlobal-Ipv6AllowMulticastProbe"); ok {
			if err = d.Set("ipv6_allow_multicast_probe", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_allow_multicast_probe: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_allow_multicast_probe: %v", err)
		}
	}

	if err = d.Set("ipv6_allow_traffic_redirect", flattenSystemGlobalIpv6AllowTrafficRedirect(o["ipv6-allow-traffic-redirect"], d, "ipv6_allow_traffic_redirect")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-allow-traffic-redirect"], "SystemGlobal-Ipv6AllowTrafficRedirect"); ok {
			if err = d.Set("ipv6_allow_traffic_redirect", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_allow_traffic_redirect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_allow_traffic_redirect: %v", err)
		}
	}

	if err = d.Set("irq_time_accounting", flattenSystemGlobalIrqTimeAccounting(o["irq-time-accounting"], d, "irq_time_accounting")); err != nil {
		if vv, ok := fortiAPIPatch(o["irq-time-accounting"], "SystemGlobal-IrqTimeAccounting"); ok {
			if err = d.Set("irq_time_accounting", vv); err != nil {
				return fmt.Errorf("Error reading irq_time_accounting: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading irq_time_accounting: %v", err)
		}
	}

	if err = d.Set("language", flattenSystemGlobalLanguage(o["language"], d, "language")); err != nil {
		if vv, ok := fortiAPIPatch(o["language"], "SystemGlobal-Language"); ok {
			if err = d.Set("language", vv); err != nil {
				return fmt.Errorf("Error reading language: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading language: %v", err)
		}
	}

	if err = d.Set("ldapconntimeout", flattenSystemGlobalLdapconntimeout(o["ldapconntimeout"], d, "ldapconntimeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldapconntimeout"], "SystemGlobal-Ldapconntimeout"); ok {
			if err = d.Set("ldapconntimeout", vv); err != nil {
				return fmt.Errorf("Error reading ldapconntimeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldapconntimeout: %v", err)
		}
	}

	if err = d.Set("legacy_poe_device_support", flattenSystemGlobalLegacyPoeDeviceSupport(o["legacy-poe-device-support"], d, "legacy_poe_device_support")); err != nil {
		if vv, ok := fortiAPIPatch(o["legacy-poe-device-support"], "SystemGlobal-LegacyPoeDeviceSupport"); ok {
			if err = d.Set("legacy_poe_device_support", vv); err != nil {
				return fmt.Errorf("Error reading legacy_poe_device_support: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading legacy_poe_device_support: %v", err)
		}
	}

	if err = d.Set("lldp_reception", flattenSystemGlobalLldpReception(o["lldp-reception"], d, "lldp_reception")); err != nil {
		if vv, ok := fortiAPIPatch(o["lldp-reception"], "SystemGlobal-LldpReception"); ok {
			if err = d.Set("lldp_reception", vv); err != nil {
				return fmt.Errorf("Error reading lldp_reception: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lldp_reception: %v", err)
		}
	}

	if err = d.Set("lldp_transmission", flattenSystemGlobalLldpTransmission(o["lldp-transmission"], d, "lldp_transmission")); err != nil {
		if vv, ok := fortiAPIPatch(o["lldp-transmission"], "SystemGlobal-LldpTransmission"); ok {
			if err = d.Set("lldp_transmission", vv); err != nil {
				return fmt.Errorf("Error reading lldp_transmission: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lldp_transmission: %v", err)
		}
	}

	if err = d.Set("log_single_cpu_high", flattenSystemGlobalLogSingleCpuHigh(o["log-single-cpu-high"], d, "log_single_cpu_high")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-single-cpu-high"], "SystemGlobal-LogSingleCpuHigh"); ok {
			if err = d.Set("log_single_cpu_high", vv); err != nil {
				return fmt.Errorf("Error reading log_single_cpu_high: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_single_cpu_high: %v", err)
		}
	}

	if err = d.Set("log_ssl_connection", flattenSystemGlobalLogSslConnection(o["log-ssl-connection"], d, "log_ssl_connection")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-ssl-connection"], "SystemGlobal-LogSslConnection"); ok {
			if err = d.Set("log_ssl_connection", vv); err != nil {
				return fmt.Errorf("Error reading log_ssl_connection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_ssl_connection: %v", err)
		}
	}

	if err = d.Set("log_uuid_address", flattenSystemGlobalLogUuidAddress(o["log-uuid-address"], d, "log_uuid_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-uuid-address"], "SystemGlobal-LogUuidAddress"); ok {
			if err = d.Set("log_uuid_address", vv); err != nil {
				return fmt.Errorf("Error reading log_uuid_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_uuid_address: %v", err)
		}
	}

	if err = d.Set("log_uuid_policy", flattenSystemGlobalLogUuidPolicy(o["log-uuid-policy"], d, "log_uuid_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-uuid-policy"], "SystemGlobal-LogUuidPolicy"); ok {
			if err = d.Set("log_uuid_policy", vv); err != nil {
				return fmt.Errorf("Error reading log_uuid_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_uuid_policy: %v", err)
		}
	}

	if err = d.Set("login_timestamp", flattenSystemGlobalLoginTimestamp(o["login-timestamp"], d, "login_timestamp")); err != nil {
		if vv, ok := fortiAPIPatch(o["login-timestamp"], "SystemGlobal-LoginTimestamp"); ok {
			if err = d.Set("login_timestamp", vv); err != nil {
				return fmt.Errorf("Error reading login_timestamp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading login_timestamp: %v", err)
		}
	}

	if err = d.Set("long_vdom_name", flattenSystemGlobalLongVdomName(o["long-vdom-name"], d, "long_vdom_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["long-vdom-name"], "SystemGlobal-LongVdomName"); ok {
			if err = d.Set("long_vdom_name", vv); err != nil {
				return fmt.Errorf("Error reading long_vdom_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading long_vdom_name: %v", err)
		}
	}

	if err = d.Set("management_ip", flattenSystemGlobalManagementIp(o["management-ip"], d, "management_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["management-ip"], "SystemGlobal-ManagementIp"); ok {
			if err = d.Set("management_ip", vv); err != nil {
				return fmt.Errorf("Error reading management_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading management_ip: %v", err)
		}
	}

	if err = d.Set("management_port", flattenSystemGlobalManagementPort(o["management-port"], d, "management_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["management-port"], "SystemGlobal-ManagementPort"); ok {
			if err = d.Set("management_port", vv); err != nil {
				return fmt.Errorf("Error reading management_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading management_port: %v", err)
		}
	}

	if err = d.Set("management_port_use_admin_sport", flattenSystemGlobalManagementPortUseAdminSport(o["management-port-use-admin-sport"], d, "management_port_use_admin_sport")); err != nil {
		if vv, ok := fortiAPIPatch(o["management-port-use-admin-sport"], "SystemGlobal-ManagementPortUseAdminSport"); ok {
			if err = d.Set("management_port_use_admin_sport", vv); err != nil {
				return fmt.Errorf("Error reading management_port_use_admin_sport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading management_port_use_admin_sport: %v", err)
		}
	}

	if err = d.Set("management_vdom", flattenSystemGlobalManagementVdom(o["management-vdom"], d, "management_vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["management-vdom"], "SystemGlobal-ManagementVdom"); ok {
			if err = d.Set("management_vdom", vv); err != nil {
				return fmt.Errorf("Error reading management_vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading management_vdom: %v", err)
		}
	}

	if err = d.Set("max_route_cache_size", flattenSystemGlobalMaxRouteCacheSize(o["max-route-cache-size"], d, "max_route_cache_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-route-cache-size"], "SystemGlobal-MaxRouteCacheSize"); ok {
			if err = d.Set("max_route_cache_size", vv); err != nil {
				return fmt.Errorf("Error reading max_route_cache_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_route_cache_size: %v", err)
		}
	}

	if err = d.Set("memory_use_threshold_extreme", flattenSystemGlobalMemoryUseThresholdExtreme(o["memory-use-threshold-extreme"], d, "memory_use_threshold_extreme")); err != nil {
		if vv, ok := fortiAPIPatch(o["memory-use-threshold-extreme"], "SystemGlobal-MemoryUseThresholdExtreme"); ok {
			if err = d.Set("memory_use_threshold_extreme", vv); err != nil {
				return fmt.Errorf("Error reading memory_use_threshold_extreme: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading memory_use_threshold_extreme: %v", err)
		}
	}

	if err = d.Set("memory_use_threshold_green", flattenSystemGlobalMemoryUseThresholdGreen(o["memory-use-threshold-green"], d, "memory_use_threshold_green")); err != nil {
		if vv, ok := fortiAPIPatch(o["memory-use-threshold-green"], "SystemGlobal-MemoryUseThresholdGreen"); ok {
			if err = d.Set("memory_use_threshold_green", vv); err != nil {
				return fmt.Errorf("Error reading memory_use_threshold_green: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading memory_use_threshold_green: %v", err)
		}
	}

	if err = d.Set("memory_use_threshold_red", flattenSystemGlobalMemoryUseThresholdRed(o["memory-use-threshold-red"], d, "memory_use_threshold_red")); err != nil {
		if vv, ok := fortiAPIPatch(o["memory-use-threshold-red"], "SystemGlobal-MemoryUseThresholdRed"); ok {
			if err = d.Set("memory_use_threshold_red", vv); err != nil {
				return fmt.Errorf("Error reading memory_use_threshold_red: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading memory_use_threshold_red: %v", err)
		}
	}

	if err = d.Set("miglog_affinity", flattenSystemGlobalMiglogAffinity(o["miglog-affinity"], d, "miglog_affinity")); err != nil {
		if vv, ok := fortiAPIPatch(o["miglog-affinity"], "SystemGlobal-MiglogAffinity"); ok {
			if err = d.Set("miglog_affinity", vv); err != nil {
				return fmt.Errorf("Error reading miglog_affinity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading miglog_affinity: %v", err)
		}
	}

	if err = d.Set("miglogd_children", flattenSystemGlobalMiglogdChildren(o["miglogd-children"], d, "miglogd_children")); err != nil {
		if vv, ok := fortiAPIPatch(o["miglogd-children"], "SystemGlobal-MiglogdChildren"); ok {
			if err = d.Set("miglogd_children", vv); err != nil {
				return fmt.Errorf("Error reading miglogd_children: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading miglogd_children: %v", err)
		}
	}

	if err = d.Set("multi_factor_authentication", flattenSystemGlobalMultiFactorAuthentication(o["multi-factor-authentication"], d, "multi_factor_authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["multi-factor-authentication"], "SystemGlobal-MultiFactorAuthentication"); ok {
			if err = d.Set("multi_factor_authentication", vv); err != nil {
				return fmt.Errorf("Error reading multi_factor_authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multi_factor_authentication: %v", err)
		}
	}

	if err = d.Set("ndp_max_entry", flattenSystemGlobalNdpMaxEntry(o["ndp-max-entry"], d, "ndp_max_entry")); err != nil {
		if vv, ok := fortiAPIPatch(o["ndp-max-entry"], "SystemGlobal-NdpMaxEntry"); ok {
			if err = d.Set("ndp_max_entry", vv); err != nil {
				return fmt.Errorf("Error reading ndp_max_entry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ndp_max_entry: %v", err)
		}
	}

	if err = d.Set("npu_neighbor_update", flattenSystemGlobalNpuNeighborUpdate(o["npu-neighbor-update"], d, "npu_neighbor_update")); err != nil {
		if vv, ok := fortiAPIPatch(o["npu-neighbor-update"], "SystemGlobal-NpuNeighborUpdate"); ok {
			if err = d.Set("npu_neighbor_update", vv); err != nil {
				return fmt.Errorf("Error reading npu_neighbor_update: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading npu_neighbor_update: %v", err)
		}
	}

	if err = d.Set("optimize_flow_mode", flattenSystemGlobalOptimizeFlowMode(o["optimize-flow-mode"], d, "optimize_flow_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["optimize-flow-mode"], "SystemGlobal-OptimizeFlowMode"); ok {
			if err = d.Set("optimize_flow_mode", vv); err != nil {
				return fmt.Errorf("Error reading optimize_flow_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading optimize_flow_mode: %v", err)
		}
	}

	if err = d.Set("per_user_bwl", flattenSystemGlobalPerUserBwl(o["per-user-bwl"], d, "per_user_bwl")); err != nil {
		if vv, ok := fortiAPIPatch(o["per-user-bwl"], "SystemGlobal-PerUserBwl"); ok {
			if err = d.Set("per_user_bwl", vv); err != nil {
				return fmt.Errorf("Error reading per_user_bwl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading per_user_bwl: %v", err)
		}
	}

	if err = d.Set("per_user_bal", flattenSystemGlobalPerUserBal(o["per-user-bal"], d, "per_user_bal")); err != nil {
		if vv, ok := fortiAPIPatch(o["per-user-bal"], "SystemGlobal-PerUserBal"); ok {
			if err = d.Set("per_user_bal", vv); err != nil {
				return fmt.Errorf("Error reading per_user_bal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading per_user_bal: %v", err)
		}
	}

	if err = d.Set("pmtu_discovery", flattenSystemGlobalPmtuDiscovery(o["pmtu-discovery"], d, "pmtu_discovery")); err != nil {
		if vv, ok := fortiAPIPatch(o["pmtu-discovery"], "SystemGlobal-PmtuDiscovery"); ok {
			if err = d.Set("pmtu_discovery", vv); err != nil {
				return fmt.Errorf("Error reading pmtu_discovery: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pmtu_discovery: %v", err)
		}
	}

	if err = d.Set("policy_auth_concurrent", flattenSystemGlobalPolicyAuthConcurrent(o["policy-auth-concurrent"], d, "policy_auth_concurrent")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-auth-concurrent"], "SystemGlobal-PolicyAuthConcurrent"); ok {
			if err = d.Set("policy_auth_concurrent", vv); err != nil {
				return fmt.Errorf("Error reading policy_auth_concurrent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_auth_concurrent: %v", err)
		}
	}

	if err = d.Set("post_login_banner", flattenSystemGlobalPostLoginBanner(o["post-login-banner"], d, "post_login_banner")); err != nil {
		if vv, ok := fortiAPIPatch(o["post-login-banner"], "SystemGlobal-PostLoginBanner"); ok {
			if err = d.Set("post_login_banner", vv); err != nil {
				return fmt.Errorf("Error reading post_login_banner: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading post_login_banner: %v", err)
		}
	}

	if err = d.Set("pre_login_banner", flattenSystemGlobalPreLoginBanner(o["pre-login-banner"], d, "pre_login_banner")); err != nil {
		if vv, ok := fortiAPIPatch(o["pre-login-banner"], "SystemGlobal-PreLoginBanner"); ok {
			if err = d.Set("pre_login_banner", vv); err != nil {
				return fmt.Errorf("Error reading pre_login_banner: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pre_login_banner: %v", err)
		}
	}

	if err = d.Set("private_data_encryption", flattenSystemGlobalPrivateDataEncryption(o["private-data-encryption"], d, "private_data_encryption")); err != nil {
		if vv, ok := fortiAPIPatch(o["private-data-encryption"], "SystemGlobal-PrivateDataEncryption"); ok {
			if err = d.Set("private_data_encryption", vv); err != nil {
				return fmt.Errorf("Error reading private_data_encryption: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading private_data_encryption: %v", err)
		}
	}

	if err = d.Set("proxy_and_explicit_proxy", flattenSystemGlobalProxyAndExplicitProxy(o["proxy-and-explicit-proxy"], d, "proxy_and_explicit_proxy")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-and-explicit-proxy"], "SystemGlobal-ProxyAndExplicitProxy"); ok {
			if err = d.Set("proxy_and_explicit_proxy", vv); err != nil {
				return fmt.Errorf("Error reading proxy_and_explicit_proxy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_and_explicit_proxy: %v", err)
		}
	}

	if err = d.Set("proxy_auth_lifetime", flattenSystemGlobalProxyAuthLifetime(o["proxy-auth-lifetime"], d, "proxy_auth_lifetime")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-auth-lifetime"], "SystemGlobal-ProxyAuthLifetime"); ok {
			if err = d.Set("proxy_auth_lifetime", vv); err != nil {
				return fmt.Errorf("Error reading proxy_auth_lifetime: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_auth_lifetime: %v", err)
		}
	}

	if err = d.Set("proxy_auth_lifetime_timeout", flattenSystemGlobalProxyAuthLifetimeTimeout(o["proxy-auth-lifetime-timeout"], d, "proxy_auth_lifetime_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-auth-lifetime-timeout"], "SystemGlobal-ProxyAuthLifetimeTimeout"); ok {
			if err = d.Set("proxy_auth_lifetime_timeout", vv); err != nil {
				return fmt.Errorf("Error reading proxy_auth_lifetime_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_auth_lifetime_timeout: %v", err)
		}
	}

	if err = d.Set("proxy_auth_timeout", flattenSystemGlobalProxyAuthTimeout(o["proxy-auth-timeout"], d, "proxy_auth_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-auth-timeout"], "SystemGlobal-ProxyAuthTimeout"); ok {
			if err = d.Set("proxy_auth_timeout", vv); err != nil {
				return fmt.Errorf("Error reading proxy_auth_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_auth_timeout: %v", err)
		}
	}

	if err = d.Set("proxy_cipher_hardware_acceleration", flattenSystemGlobalProxyCipherHardwareAcceleration(o["proxy-cipher-hardware-acceleration"], d, "proxy_cipher_hardware_acceleration")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-cipher-hardware-acceleration"], "SystemGlobal-ProxyCipherHardwareAcceleration"); ok {
			if err = d.Set("proxy_cipher_hardware_acceleration", vv); err != nil {
				return fmt.Errorf("Error reading proxy_cipher_hardware_acceleration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_cipher_hardware_acceleration: %v", err)
		}
	}

	if err = d.Set("proxy_kxp_hardware_acceleration", flattenSystemGlobalProxyKxpHardwareAcceleration(o["proxy-kxp-hardware-acceleration"], d, "proxy_kxp_hardware_acceleration")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-kxp-hardware-acceleration"], "SystemGlobal-ProxyKxpHardwareAcceleration"); ok {
			if err = d.Set("proxy_kxp_hardware_acceleration", vv); err != nil {
				return fmt.Errorf("Error reading proxy_kxp_hardware_acceleration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_kxp_hardware_acceleration: %v", err)
		}
	}

	if err = d.Set("proxy_cert_use_mgmt_vdom", flattenSystemGlobalProxyCertUseMgmtVdom(o["proxy-cert-use-mgmt-vdom"], d, "proxy_cert_use_mgmt_vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-cert-use-mgmt-vdom"], "SystemGlobal-ProxyCertUseMgmtVdom"); ok {
			if err = d.Set("proxy_cert_use_mgmt_vdom", vv); err != nil {
				return fmt.Errorf("Error reading proxy_cert_use_mgmt_vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_cert_use_mgmt_vdom: %v", err)
		}
	}

	if err = d.Set("proxy_hardware_acceleration", flattenSystemGlobalProxyHardwareAcceleration(o["proxy-hardware-acceleration"], d, "proxy_hardware_acceleration")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-hardware-acceleration"], "SystemGlobal-ProxyHardwareAcceleration"); ok {
			if err = d.Set("proxy_hardware_acceleration", vv); err != nil {
				return fmt.Errorf("Error reading proxy_hardware_acceleration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_hardware_acceleration: %v", err)
		}
	}

	if err = d.Set("proxy_keep_alive_mode", flattenSystemGlobalProxyKeepAliveMode(o["proxy-keep-alive-mode"], d, "proxy_keep_alive_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-keep-alive-mode"], "SystemGlobal-ProxyKeepAliveMode"); ok {
			if err = d.Set("proxy_keep_alive_mode", vv); err != nil {
				return fmt.Errorf("Error reading proxy_keep_alive_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_keep_alive_mode: %v", err)
		}
	}

	if err = d.Set("proxy_re_authentication_mode", flattenSystemGlobalProxyReAuthenticationMode(o["proxy-re-authentication-mode"], d, "proxy_re_authentication_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-re-authentication-mode"], "SystemGlobal-ProxyReAuthenticationMode"); ok {
			if err = d.Set("proxy_re_authentication_mode", vv); err != nil {
				return fmt.Errorf("Error reading proxy_re_authentication_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_re_authentication_mode: %v", err)
		}
	}

	if err = d.Set("proxy_re_authentication_time", flattenSystemGlobalProxyReAuthenticationTime(o["proxy-re-authentication-time"], d, "proxy_re_authentication_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-re-authentication-time"], "SystemGlobal-ProxyReAuthenticationTime"); ok {
			if err = d.Set("proxy_re_authentication_time", vv); err != nil {
				return fmt.Errorf("Error reading proxy_re_authentication_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_re_authentication_time: %v", err)
		}
	}

	if err = d.Set("proxy_resource_mode", flattenSystemGlobalProxyResourceMode(o["proxy-resource-mode"], d, "proxy_resource_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-resource-mode"], "SystemGlobal-ProxyResourceMode"); ok {
			if err = d.Set("proxy_resource_mode", vv); err != nil {
				return fmt.Errorf("Error reading proxy_resource_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_resource_mode: %v", err)
		}
	}

	if err = d.Set("proxy_worker_count", flattenSystemGlobalProxyWorkerCount(o["proxy-worker-count"], d, "proxy_worker_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-worker-count"], "SystemGlobal-ProxyWorkerCount"); ok {
			if err = d.Set("proxy_worker_count", vv); err != nil {
				return fmt.Errorf("Error reading proxy_worker_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_worker_count: %v", err)
		}
	}

	if err = d.Set("purdue_level", flattenSystemGlobalPurdueLevel(o["purdue-level"], d, "purdue_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["purdue-level"], "SystemGlobal-PurdueLevel"); ok {
			if err = d.Set("purdue_level", vv); err != nil {
				return fmt.Errorf("Error reading purdue_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading purdue_level: %v", err)
		}
	}

	if err = d.Set("qsfp28_40g_port", flattenSystemGlobalQsfp2840GPort(o["qsfp28-40g-port"], d, "qsfp28_40g_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["qsfp28-40g-port"], "SystemGlobal-Qsfp2840GPort"); ok {
			if err = d.Set("qsfp28_40g_port", vv); err != nil {
				return fmt.Errorf("Error reading qsfp28_40g_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading qsfp28_40g_port: %v", err)
		}
	}

	if err = d.Set("qsfpdd_100g_port", flattenSystemGlobalQsfpdd100GPort(o["qsfpdd-100g-port"], d, "qsfpdd_100g_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["qsfpdd-100g-port"], "SystemGlobal-Qsfpdd100GPort"); ok {
			if err = d.Set("qsfpdd_100g_port", vv); err != nil {
				return fmt.Errorf("Error reading qsfpdd_100g_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading qsfpdd_100g_port: %v", err)
		}
	}

	if err = d.Set("qsfpdd_split8_port", flattenSystemGlobalQsfpddSplit8Port(o["qsfpdd-split8-port"], d, "qsfpdd_split8_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["qsfpdd-split8-port"], "SystemGlobal-QsfpddSplit8Port"); ok {
			if err = d.Set("qsfpdd_split8_port", vv); err != nil {
				return fmt.Errorf("Error reading qsfpdd_split8_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading qsfpdd_split8_port: %v", err)
		}
	}

	if err = d.Set("quic_ack_thresold", flattenSystemGlobalQuicAckThresold(o["quic-ack-thresold"], d, "quic_ack_thresold")); err != nil {
		if vv, ok := fortiAPIPatch(o["quic-ack-thresold"], "SystemGlobal-QuicAckThresold"); ok {
			if err = d.Set("quic_ack_thresold", vv); err != nil {
				return fmt.Errorf("Error reading quic_ack_thresold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quic_ack_thresold: %v", err)
		}
	}

	if err = d.Set("quic_congestion_control_algo", flattenSystemGlobalQuicCongestionControlAlgo(o["quic-congestion-control-algo"], d, "quic_congestion_control_algo")); err != nil {
		if vv, ok := fortiAPIPatch(o["quic-congestion-control-algo"], "SystemGlobal-QuicCongestionControlAlgo"); ok {
			if err = d.Set("quic_congestion_control_algo", vv); err != nil {
				return fmt.Errorf("Error reading quic_congestion_control_algo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quic_congestion_control_algo: %v", err)
		}
	}

	if err = d.Set("quic_max_datagram_size", flattenSystemGlobalQuicMaxDatagramSize(o["quic-max-datagram-size"], d, "quic_max_datagram_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["quic-max-datagram-size"], "SystemGlobal-QuicMaxDatagramSize"); ok {
			if err = d.Set("quic_max_datagram_size", vv); err != nil {
				return fmt.Errorf("Error reading quic_max_datagram_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quic_max_datagram_size: %v", err)
		}
	}

	if err = d.Set("quic_pmtud", flattenSystemGlobalQuicPmtud(o["quic-pmtud"], d, "quic_pmtud")); err != nil {
		if vv, ok := fortiAPIPatch(o["quic-pmtud"], "SystemGlobal-QuicPmtud"); ok {
			if err = d.Set("quic_pmtud", vv); err != nil {
				return fmt.Errorf("Error reading quic_pmtud: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quic_pmtud: %v", err)
		}
	}

	if err = d.Set("quic_tls_handshake_timeout", flattenSystemGlobalQuicTlsHandshakeTimeout(o["quic-tls-handshake-timeout"], d, "quic_tls_handshake_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["quic-tls-handshake-timeout"], "SystemGlobal-QuicTlsHandshakeTimeout"); ok {
			if err = d.Set("quic_tls_handshake_timeout", vv); err != nil {
				return fmt.Errorf("Error reading quic_tls_handshake_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quic_tls_handshake_timeout: %v", err)
		}
	}

	if err = d.Set("quic_udp_payload_size_shaping_per_cid", flattenSystemGlobalQuicUdpPayloadSizeShapingPerCid(o["quic-udp-payload-size-shaping-per-cid"], d, "quic_udp_payload_size_shaping_per_cid")); err != nil {
		if vv, ok := fortiAPIPatch(o["quic-udp-payload-size-shaping-per-cid"], "SystemGlobal-QuicUdpPayloadSizeShapingPerCid"); ok {
			if err = d.Set("quic_udp_payload_size_shaping_per_cid", vv); err != nil {
				return fmt.Errorf("Error reading quic_udp_payload_size_shaping_per_cid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quic_udp_payload_size_shaping_per_cid: %v", err)
		}
	}

	if err = d.Set("radius_port", flattenSystemGlobalRadiusPort(o["radius-port"], d, "radius_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-port"], "SystemGlobal-RadiusPort"); ok {
			if err = d.Set("radius_port", vv); err != nil {
				return fmt.Errorf("Error reading radius_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_port: %v", err)
		}
	}

	if err = d.Set("reboot_upon_config_restore", flattenSystemGlobalRebootUponConfigRestore(o["reboot-upon-config-restore"], d, "reboot_upon_config_restore")); err != nil {
		if vv, ok := fortiAPIPatch(o["reboot-upon-config-restore"], "SystemGlobal-RebootUponConfigRestore"); ok {
			if err = d.Set("reboot_upon_config_restore", vv); err != nil {
				return fmt.Errorf("Error reading reboot_upon_config_restore: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reboot_upon_config_restore: %v", err)
		}
	}

	if err = d.Set("refresh", flattenSystemGlobalRefresh(o["refresh"], d, "refresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["refresh"], "SystemGlobal-Refresh"); ok {
			if err = d.Set("refresh", vv); err != nil {
				return fmt.Errorf("Error reading refresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading refresh: %v", err)
		}
	}

	if err = d.Set("remoteauthtimeout", flattenSystemGlobalRemoteauthtimeout(o["remoteauthtimeout"], d, "remoteauthtimeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["remoteauthtimeout"], "SystemGlobal-Remoteauthtimeout"); ok {
			if err = d.Set("remoteauthtimeout", vv); err != nil {
				return fmt.Errorf("Error reading remoteauthtimeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remoteauthtimeout: %v", err)
		}
	}

	if err = d.Set("reset_sessionless_tcp", flattenSystemGlobalResetSessionlessTcp(o["reset-sessionless-tcp"], d, "reset_sessionless_tcp")); err != nil {
		if vv, ok := fortiAPIPatch(o["reset-sessionless-tcp"], "SystemGlobal-ResetSessionlessTcp"); ok {
			if err = d.Set("reset_sessionless_tcp", vv); err != nil {
				return fmt.Errorf("Error reading reset_sessionless_tcp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reset_sessionless_tcp: %v", err)
		}
	}

	if err = d.Set("restart_time", flattenSystemGlobalRestartTime(o["restart-time"], d, "restart_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["restart-time"], "SystemGlobal-RestartTime"); ok {
			if err = d.Set("restart_time", vv); err != nil {
				return fmt.Errorf("Error reading restart_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restart_time: %v", err)
		}
	}

	if err = d.Set("revision_backup_on_logout", flattenSystemGlobalRevisionBackupOnLogout(o["revision-backup-on-logout"], d, "revision_backup_on_logout")); err != nil {
		if vv, ok := fortiAPIPatch(o["revision-backup-on-logout"], "SystemGlobal-RevisionBackupOnLogout"); ok {
			if err = d.Set("revision_backup_on_logout", vv); err != nil {
				return fmt.Errorf("Error reading revision_backup_on_logout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading revision_backup_on_logout: %v", err)
		}
	}

	if err = d.Set("revision_image_auto_backup", flattenSystemGlobalRevisionImageAutoBackup(o["revision-image-auto-backup"], d, "revision_image_auto_backup")); err != nil {
		if vv, ok := fortiAPIPatch(o["revision-image-auto-backup"], "SystemGlobal-RevisionImageAutoBackup"); ok {
			if err = d.Set("revision_image_auto_backup", vv); err != nil {
				return fmt.Errorf("Error reading revision_image_auto_backup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading revision_image_auto_backup: %v", err)
		}
	}

	if err = d.Set("scanunit_count", flattenSystemGlobalScanunitCount(o["scanunit-count"], d, "scanunit_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["scanunit-count"], "SystemGlobal-ScanunitCount"); ok {
			if err = d.Set("scanunit_count", vv); err != nil {
				return fmt.Errorf("Error reading scanunit_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scanunit_count: %v", err)
		}
	}

	if err = d.Set("security_rating_result_submission", flattenSystemGlobalSecurityRatingResultSubmission(o["security-rating-result-submission"], d, "security_rating_result_submission")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-rating-result-submission"], "SystemGlobal-SecurityRatingResultSubmission"); ok {
			if err = d.Set("security_rating_result_submission", vv); err != nil {
				return fmt.Errorf("Error reading security_rating_result_submission: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_rating_result_submission: %v", err)
		}
	}

	if err = d.Set("security_rating_run_on_schedule", flattenSystemGlobalSecurityRatingRunOnSchedule(o["security-rating-run-on-schedule"], d, "security_rating_run_on_schedule")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-rating-run-on-schedule"], "SystemGlobal-SecurityRatingRunOnSchedule"); ok {
			if err = d.Set("security_rating_run_on_schedule", vv); err != nil {
				return fmt.Errorf("Error reading security_rating_run_on_schedule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_rating_run_on_schedule: %v", err)
		}
	}

	if err = d.Set("send_pmtu_icmp", flattenSystemGlobalSendPmtuIcmp(o["send-pmtu-icmp"], d, "send_pmtu_icmp")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-pmtu-icmp"], "SystemGlobal-SendPmtuIcmp"); ok {
			if err = d.Set("send_pmtu_icmp", vv); err != nil {
				return fmt.Errorf("Error reading send_pmtu_icmp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_pmtu_icmp: %v", err)
		}
	}

	if err = d.Set("sflowd_max_children_num", flattenSystemGlobalSflowdMaxChildrenNum(o["sflowd-max-children-num"], d, "sflowd_max_children_num")); err != nil {
		if vv, ok := fortiAPIPatch(o["sflowd-max-children-num"], "SystemGlobal-SflowdMaxChildrenNum"); ok {
			if err = d.Set("sflowd_max_children_num", vv); err != nil {
				return fmt.Errorf("Error reading sflowd_max_children_num: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sflowd_max_children_num: %v", err)
		}
	}

	if err = d.Set("show_backplane_intf", flattenSystemGlobalShowBackplaneIntf(o["show-backplane-intf"], d, "show_backplane_intf")); err != nil {
		if vv, ok := fortiAPIPatch(o["show-backplane-intf"], "SystemGlobal-ShowBackplaneIntf"); ok {
			if err = d.Set("show_backplane_intf", vv); err != nil {
				return fmt.Errorf("Error reading show_backplane_intf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading show_backplane_intf: %v", err)
		}
	}

	if err = d.Set("snat_route_change", flattenSystemGlobalSnatRouteChange(o["snat-route-change"], d, "snat_route_change")); err != nil {
		if vv, ok := fortiAPIPatch(o["snat-route-change"], "SystemGlobal-SnatRouteChange"); ok {
			if err = d.Set("snat_route_change", vv); err != nil {
				return fmt.Errorf("Error reading snat_route_change: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading snat_route_change: %v", err)
		}
	}

	if err = d.Set("special_file_23_support", flattenSystemGlobalSpecialFile23Support(o["special-file-23-support"], d, "special_file_23_support")); err != nil {
		if vv, ok := fortiAPIPatch(o["special-file-23-support"], "SystemGlobal-SpecialFile23Support"); ok {
			if err = d.Set("special_file_23_support", vv); err != nil {
				return fmt.Errorf("Error reading special_file_23_support: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading special_file_23_support: %v", err)
		}
	}

	if err = d.Set("speedtest_server", flattenSystemGlobalSpeedtestServer(o["speedtest-server"], d, "speedtest_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["speedtest-server"], "SystemGlobal-SpeedtestServer"); ok {
			if err = d.Set("speedtest_server", vv); err != nil {
				return fmt.Errorf("Error reading speedtest_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading speedtest_server: %v", err)
		}
	}

	if err = d.Set("speedtestd_ctrl_port", flattenSystemGlobalSpeedtestdCtrlPort(o["speedtestd-ctrl-port"], d, "speedtestd_ctrl_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["speedtestd-ctrl-port"], "SystemGlobal-SpeedtestdCtrlPort"); ok {
			if err = d.Set("speedtestd_ctrl_port", vv); err != nil {
				return fmt.Errorf("Error reading speedtestd_ctrl_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading speedtestd_ctrl_port: %v", err)
		}
	}

	if err = d.Set("speedtestd_server_port", flattenSystemGlobalSpeedtestdServerPort(o["speedtestd-server-port"], d, "speedtestd_server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["speedtestd-server-port"], "SystemGlobal-SpeedtestdServerPort"); ok {
			if err = d.Set("speedtestd_server_port", vv); err != nil {
				return fmt.Errorf("Error reading speedtestd_server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading speedtestd_server_port: %v", err)
		}
	}

	if err = d.Set("split_port", flattenSystemGlobalSplitPort(o["split-port"], d, "split_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["split-port"], "SystemGlobal-SplitPort"); ok {
			if err = d.Set("split_port", vv); err != nil {
				return fmt.Errorf("Error reading split_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading split_port: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("split_port_mode", flattenSystemGlobalSplitPortMode(o["split-port-mode"], d, "split_port_mode")); err != nil {
			if vv, ok := fortiAPIPatch(o["split-port-mode"], "SystemGlobal-SplitPortMode"); ok {
				if err = d.Set("split_port_mode", vv); err != nil {
					return fmt.Errorf("Error reading split_port_mode: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading split_port_mode: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("split_port_mode"); ok {
			if err = d.Set("split_port_mode", flattenSystemGlobalSplitPortMode(o["split-port-mode"], d, "split_port_mode")); err != nil {
				if vv, ok := fortiAPIPatch(o["split-port-mode"], "SystemGlobal-SplitPortMode"); ok {
					if err = d.Set("split_port_mode", vv); err != nil {
						return fmt.Errorf("Error reading split_port_mode: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading split_port_mode: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssd_trim_date", flattenSystemGlobalSsdTrimDate(o["ssd-trim-date"], d, "ssd_trim_date")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssd-trim-date"], "SystemGlobal-SsdTrimDate"); ok {
			if err = d.Set("ssd_trim_date", vv); err != nil {
				return fmt.Errorf("Error reading ssd_trim_date: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssd_trim_date: %v", err)
		}
	}

	if err = d.Set("ssd_trim_freq", flattenSystemGlobalSsdTrimFreq(o["ssd-trim-freq"], d, "ssd_trim_freq")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssd-trim-freq"], "SystemGlobal-SsdTrimFreq"); ok {
			if err = d.Set("ssd_trim_freq", vv); err != nil {
				return fmt.Errorf("Error reading ssd_trim_freq: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssd_trim_freq: %v", err)
		}
	}

	if err = d.Set("ssd_trim_hour", flattenSystemGlobalSsdTrimHour(o["ssd-trim-hour"], d, "ssd_trim_hour")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssd-trim-hour"], "SystemGlobal-SsdTrimHour"); ok {
			if err = d.Set("ssd_trim_hour", vv); err != nil {
				return fmt.Errorf("Error reading ssd_trim_hour: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssd_trim_hour: %v", err)
		}
	}

	if err = d.Set("ssd_trim_min", flattenSystemGlobalSsdTrimMin(o["ssd-trim-min"], d, "ssd_trim_min")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssd-trim-min"], "SystemGlobal-SsdTrimMin"); ok {
			if err = d.Set("ssd_trim_min", vv); err != nil {
				return fmt.Errorf("Error reading ssd_trim_min: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssd_trim_min: %v", err)
		}
	}

	if err = d.Set("ssd_trim_weekday", flattenSystemGlobalSsdTrimWeekday(o["ssd-trim-weekday"], d, "ssd_trim_weekday")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssd-trim-weekday"], "SystemGlobal-SsdTrimWeekday"); ok {
			if err = d.Set("ssd_trim_weekday", vv); err != nil {
				return fmt.Errorf("Error reading ssd_trim_weekday: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssd_trim_weekday: %v", err)
		}
	}

	if err = d.Set("ssh_cbc_cipher", flattenSystemGlobalSshCbcCipher(o["ssh-cbc-cipher"], d, "ssh_cbc_cipher")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-cbc-cipher"], "SystemGlobal-SshCbcCipher"); ok {
			if err = d.Set("ssh_cbc_cipher", vv); err != nil {
				return fmt.Errorf("Error reading ssh_cbc_cipher: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_cbc_cipher: %v", err)
		}
	}

	if err = d.Set("ssh_enc_algo", flattenSystemGlobalSshEncAlgo(o["ssh-enc-algo"], d, "ssh_enc_algo")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-enc-algo"], "SystemGlobal-SshEncAlgo"); ok {
			if err = d.Set("ssh_enc_algo", vv); err != nil {
				return fmt.Errorf("Error reading ssh_enc_algo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_enc_algo: %v", err)
		}
	}

	if err = d.Set("ssh_hmac_md5", flattenSystemGlobalSshHmacMd5(o["ssh-hmac-md5"], d, "ssh_hmac_md5")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-hmac-md5"], "SystemGlobal-SshHmacMd5"); ok {
			if err = d.Set("ssh_hmac_md5", vv); err != nil {
				return fmt.Errorf("Error reading ssh_hmac_md5: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_hmac_md5: %v", err)
		}
	}

	if err = d.Set("ssh_hostkey", flattenSystemGlobalSshHostkey(o["ssh-hostkey"], d, "ssh_hostkey")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-hostkey"], "SystemGlobal-SshHostkey"); ok {
			if err = d.Set("ssh_hostkey", vv); err != nil {
				return fmt.Errorf("Error reading ssh_hostkey: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_hostkey: %v", err)
		}
	}

	if err = d.Set("ssh_hostkey_algo", flattenSystemGlobalSshHostkeyAlgo(o["ssh-hostkey-algo"], d, "ssh_hostkey_algo")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-hostkey-algo"], "SystemGlobal-SshHostkeyAlgo"); ok {
			if err = d.Set("ssh_hostkey_algo", vv); err != nil {
				return fmt.Errorf("Error reading ssh_hostkey_algo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_hostkey_algo: %v", err)
		}
	}

	if err = d.Set("ssh_hostkey_override", flattenSystemGlobalSshHostkeyOverride(o["ssh-hostkey-override"], d, "ssh_hostkey_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-hostkey-override"], "SystemGlobal-SshHostkeyOverride"); ok {
			if err = d.Set("ssh_hostkey_override", vv); err != nil {
				return fmt.Errorf("Error reading ssh_hostkey_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_hostkey_override: %v", err)
		}
	}

	if err = d.Set("ssh_kex_algo", flattenSystemGlobalSshKexAlgo(o["ssh-kex-algo"], d, "ssh_kex_algo")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-kex-algo"], "SystemGlobal-SshKexAlgo"); ok {
			if err = d.Set("ssh_kex_algo", vv); err != nil {
				return fmt.Errorf("Error reading ssh_kex_algo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_kex_algo: %v", err)
		}
	}

	if err = d.Set("ssh_kex_sha1", flattenSystemGlobalSshKexSha1(o["ssh-kex-sha1"], d, "ssh_kex_sha1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-kex-sha1"], "SystemGlobal-SshKexSha1"); ok {
			if err = d.Set("ssh_kex_sha1", vv); err != nil {
				return fmt.Errorf("Error reading ssh_kex_sha1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_kex_sha1: %v", err)
		}
	}

	if err = d.Set("ssh_mac_algo", flattenSystemGlobalSshMacAlgo(o["ssh-mac-algo"], d, "ssh_mac_algo")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-mac-algo"], "SystemGlobal-SshMacAlgo"); ok {
			if err = d.Set("ssh_mac_algo", vv); err != nil {
				return fmt.Errorf("Error reading ssh_mac_algo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_mac_algo: %v", err)
		}
	}

	if err = d.Set("ssh_mac_weak", flattenSystemGlobalSshMacWeak(o["ssh-mac-weak"], d, "ssh_mac_weak")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-mac-weak"], "SystemGlobal-SshMacWeak"); ok {
			if err = d.Set("ssh_mac_weak", vv); err != nil {
				return fmt.Errorf("Error reading ssh_mac_weak: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_mac_weak: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", flattenSystemGlobalSslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-min-proto-version"], "SystemGlobal-SslMinProtoVersion"); ok {
			if err = d.Set("ssl_min_proto_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	if err = d.Set("ssl_static_key_ciphers", flattenSystemGlobalSslStaticKeyCiphers(o["ssl-static-key-ciphers"], d, "ssl_static_key_ciphers")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-static-key-ciphers"], "SystemGlobal-SslStaticKeyCiphers"); ok {
			if err = d.Set("ssl_static_key_ciphers", vv); err != nil {
				return fmt.Errorf("Error reading ssl_static_key_ciphers: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_static_key_ciphers: %v", err)
		}
	}

	if err = d.Set("sslvpn_cipher_hardware_acceleration", flattenSystemGlobalSslvpnCipherHardwareAcceleration(o["sslvpn-cipher-hardware-acceleration"], d, "sslvpn_cipher_hardware_acceleration")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn-cipher-hardware-acceleration"], "SystemGlobal-SslvpnCipherHardwareAcceleration"); ok {
			if err = d.Set("sslvpn_cipher_hardware_acceleration", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn_cipher_hardware_acceleration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn_cipher_hardware_acceleration: %v", err)
		}
	}

	if err = d.Set("sslvpn_ems_sn_check", flattenSystemGlobalSslvpnEmsSnCheck(o["sslvpn-ems-sn-check"], d, "sslvpn_ems_sn_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn-ems-sn-check"], "SystemGlobal-SslvpnEmsSnCheck"); ok {
			if err = d.Set("sslvpn_ems_sn_check", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn_ems_sn_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn_ems_sn_check: %v", err)
		}
	}

	if err = d.Set("sslvpn_kxp_hardware_acceleration", flattenSystemGlobalSslvpnKxpHardwareAcceleration(o["sslvpn-kxp-hardware-acceleration"], d, "sslvpn_kxp_hardware_acceleration")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn-kxp-hardware-acceleration"], "SystemGlobal-SslvpnKxpHardwareAcceleration"); ok {
			if err = d.Set("sslvpn_kxp_hardware_acceleration", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn_kxp_hardware_acceleration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn_kxp_hardware_acceleration: %v", err)
		}
	}

	if err = d.Set("sslvpn_max_worker_count", flattenSystemGlobalSslvpnMaxWorkerCount(o["sslvpn-max-worker-count"], d, "sslvpn_max_worker_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn-max-worker-count"], "SystemGlobal-SslvpnMaxWorkerCount"); ok {
			if err = d.Set("sslvpn_max_worker_count", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn_max_worker_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn_max_worker_count: %v", err)
		}
	}

	if err = d.Set("sslvpn_plugin_version_check", flattenSystemGlobalSslvpnPluginVersionCheck(o["sslvpn-plugin-version-check"], d, "sslvpn_plugin_version_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn-plugin-version-check"], "SystemGlobal-SslvpnPluginVersionCheck"); ok {
			if err = d.Set("sslvpn_plugin_version_check", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn_plugin_version_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn_plugin_version_check: %v", err)
		}
	}

	if err = d.Set("sslvpn_web_mode", flattenSystemGlobalSslvpnWebMode(o["sslvpn-web-mode"], d, "sslvpn_web_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn-web-mode"], "SystemGlobal-SslvpnWebMode"); ok {
			if err = d.Set("sslvpn_web_mode", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn_web_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn_web_mode: %v", err)
		}
	}

	if err = d.Set("strict_dirty_session_check", flattenSystemGlobalStrictDirtySessionCheck(o["strict-dirty-session-check"], d, "strict_dirty_session_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["strict-dirty-session-check"], "SystemGlobal-StrictDirtySessionCheck"); ok {
			if err = d.Set("strict_dirty_session_check", vv); err != nil {
				return fmt.Errorf("Error reading strict_dirty_session_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading strict_dirty_session_check: %v", err)
		}
	}

	if err = d.Set("strong_crypto", flattenSystemGlobalStrongCrypto(o["strong-crypto"], d, "strong_crypto")); err != nil {
		if vv, ok := fortiAPIPatch(o["strong-crypto"], "SystemGlobal-StrongCrypto"); ok {
			if err = d.Set("strong_crypto", vv); err != nil {
				return fmt.Errorf("Error reading strong_crypto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading strong_crypto: %v", err)
		}
	}

	if err = d.Set("switch_controller", flattenSystemGlobalSwitchController(o["switch-controller"], d, "switch_controller")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller"], "SystemGlobal-SwitchController"); ok {
			if err = d.Set("switch_controller", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller: %v", err)
		}
	}

	if err = d.Set("switch_controller_reserved_network", flattenSystemGlobalSwitchControllerReservedNetwork(o["switch-controller-reserved-network"], d, "switch_controller_reserved_network")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-reserved-network"], "SystemGlobal-SwitchControllerReservedNetwork"); ok {
			if err = d.Set("switch_controller_reserved_network", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_reserved_network: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_reserved_network: %v", err)
		}
	}

	if err = d.Set("sys_file_check_interval", flattenSystemGlobalSysFileCheckInterval(o["sys-file-check-interval"], d, "sys_file_check_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["sys-file-check-interval"], "SystemGlobal-SysFileCheckInterval"); ok {
			if err = d.Set("sys_file_check_interval", vv); err != nil {
				return fmt.Errorf("Error reading sys_file_check_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sys_file_check_interval: %v", err)
		}
	}

	if err = d.Set("sys_perf_log_interval", flattenSystemGlobalSysPerfLogInterval(o["sys-perf-log-interval"], d, "sys_perf_log_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["sys-perf-log-interval"], "SystemGlobal-SysPerfLogInterval"); ok {
			if err = d.Set("sys_perf_log_interval", vv); err != nil {
				return fmt.Errorf("Error reading sys_perf_log_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sys_perf_log_interval: %v", err)
		}
	}

	if err = d.Set("syslog_affinity", flattenSystemGlobalSyslogAffinity(o["syslog-affinity"], d, "syslog_affinity")); err != nil {
		if vv, ok := fortiAPIPatch(o["syslog-affinity"], "SystemGlobal-SyslogAffinity"); ok {
			if err = d.Set("syslog_affinity", vv); err != nil {
				return fmt.Errorf("Error reading syslog_affinity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading syslog_affinity: %v", err)
		}
	}

	if err = d.Set("tcp_halfclose_timer", flattenSystemGlobalTcpHalfcloseTimer(o["tcp-halfclose-timer"], d, "tcp_halfclose_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-halfclose-timer"], "SystemGlobal-TcpHalfcloseTimer"); ok {
			if err = d.Set("tcp_halfclose_timer", vv); err != nil {
				return fmt.Errorf("Error reading tcp_halfclose_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_halfclose_timer: %v", err)
		}
	}

	if err = d.Set("tcp_halfopen_timer", flattenSystemGlobalTcpHalfopenTimer(o["tcp-halfopen-timer"], d, "tcp_halfopen_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-halfopen-timer"], "SystemGlobal-TcpHalfopenTimer"); ok {
			if err = d.Set("tcp_halfopen_timer", vv); err != nil {
				return fmt.Errorf("Error reading tcp_halfopen_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_halfopen_timer: %v", err)
		}
	}

	if err = d.Set("tcp_option", flattenSystemGlobalTcpOption(o["tcp-option"], d, "tcp_option")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-option"], "SystemGlobal-TcpOption"); ok {
			if err = d.Set("tcp_option", vv); err != nil {
				return fmt.Errorf("Error reading tcp_option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_option: %v", err)
		}
	}

	if err = d.Set("tcp_rst_timer", flattenSystemGlobalTcpRstTimer(o["tcp-rst-timer"], d, "tcp_rst_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-rst-timer"], "SystemGlobal-TcpRstTimer"); ok {
			if err = d.Set("tcp_rst_timer", vv); err != nil {
				return fmt.Errorf("Error reading tcp_rst_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_rst_timer: %v", err)
		}
	}

	if err = d.Set("tcp_timewait_timer", flattenSystemGlobalTcpTimewaitTimer(o["tcp-timewait-timer"], d, "tcp_timewait_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-timewait-timer"], "SystemGlobal-TcpTimewaitTimer"); ok {
			if err = d.Set("tcp_timewait_timer", vv); err != nil {
				return fmt.Errorf("Error reading tcp_timewait_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_timewait_timer: %v", err)
		}
	}

	if err = d.Set("tftp", flattenSystemGlobalTftp(o["tftp"], d, "tftp")); err != nil {
		if vv, ok := fortiAPIPatch(o["tftp"], "SystemGlobal-Tftp"); ok {
			if err = d.Set("tftp", vv); err != nil {
				return fmt.Errorf("Error reading tftp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tftp: %v", err)
		}
	}

	if err = d.Set("timezone", flattenSystemGlobalTimezone(o["timezone"], d, "timezone")); err != nil {
		if vv, ok := fortiAPIPatch(o["timezone"], "SystemGlobal-Timezone"); ok {
			if err = d.Set("timezone", vv); err != nil {
				return fmt.Errorf("Error reading timezone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timezone: %v", err)
		}
	}

	if err = d.Set("traffic_priority", flattenSystemGlobalTrafficPriority(o["traffic-priority"], d, "traffic_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-priority"], "SystemGlobal-TrafficPriority"); ok {
			if err = d.Set("traffic_priority", vv); err != nil {
				return fmt.Errorf("Error reading traffic_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_priority: %v", err)
		}
	}

	if err = d.Set("traffic_priority_level", flattenSystemGlobalTrafficPriorityLevel(o["traffic-priority-level"], d, "traffic_priority_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-priority-level"], "SystemGlobal-TrafficPriorityLevel"); ok {
			if err = d.Set("traffic_priority_level", vv); err != nil {
				return fmt.Errorf("Error reading traffic_priority_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_priority_level: %v", err)
		}
	}

	if err = d.Set("two_factor_email_expiry", flattenSystemGlobalTwoFactorEmailExpiry(o["two-factor-email-expiry"], d, "two_factor_email_expiry")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor-email-expiry"], "SystemGlobal-TwoFactorEmailExpiry"); ok {
			if err = d.Set("two_factor_email_expiry", vv); err != nil {
				return fmt.Errorf("Error reading two_factor_email_expiry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor_email_expiry: %v", err)
		}
	}

	if err = d.Set("two_factor_fac_expiry", flattenSystemGlobalTwoFactorFacExpiry(o["two-factor-fac-expiry"], d, "two_factor_fac_expiry")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor-fac-expiry"], "SystemGlobal-TwoFactorFacExpiry"); ok {
			if err = d.Set("two_factor_fac_expiry", vv); err != nil {
				return fmt.Errorf("Error reading two_factor_fac_expiry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor_fac_expiry: %v", err)
		}
	}

	if err = d.Set("two_factor_ftk_expiry", flattenSystemGlobalTwoFactorFtkExpiry(o["two-factor-ftk-expiry"], d, "two_factor_ftk_expiry")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor-ftk-expiry"], "SystemGlobal-TwoFactorFtkExpiry"); ok {
			if err = d.Set("two_factor_ftk_expiry", vv); err != nil {
				return fmt.Errorf("Error reading two_factor_ftk_expiry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor_ftk_expiry: %v", err)
		}
	}

	if err = d.Set("two_factor_ftm_expiry", flattenSystemGlobalTwoFactorFtmExpiry(o["two-factor-ftm-expiry"], d, "two_factor_ftm_expiry")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor-ftm-expiry"], "SystemGlobal-TwoFactorFtmExpiry"); ok {
			if err = d.Set("two_factor_ftm_expiry", vv); err != nil {
				return fmt.Errorf("Error reading two_factor_ftm_expiry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor_ftm_expiry: %v", err)
		}
	}

	if err = d.Set("two_factor_sms_expiry", flattenSystemGlobalTwoFactorSmsExpiry(o["two-factor-sms-expiry"], d, "two_factor_sms_expiry")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor-sms-expiry"], "SystemGlobal-TwoFactorSmsExpiry"); ok {
			if err = d.Set("two_factor_sms_expiry", vv); err != nil {
				return fmt.Errorf("Error reading two_factor_sms_expiry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor_sms_expiry: %v", err)
		}
	}

	if err = d.Set("udp_idle_timer", flattenSystemGlobalUdpIdleTimer(o["udp-idle-timer"], d, "udp_idle_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["udp-idle-timer"], "SystemGlobal-UdpIdleTimer"); ok {
			if err = d.Set("udp_idle_timer", vv); err != nil {
				return fmt.Errorf("Error reading udp_idle_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udp_idle_timer: %v", err)
		}
	}

	if err = d.Set("url_filter_affinity", flattenSystemGlobalUrlFilterAffinity(o["url-filter-affinity"], d, "url_filter_affinity")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-filter-affinity"], "SystemGlobal-UrlFilterAffinity"); ok {
			if err = d.Set("url_filter_affinity", vv); err != nil {
				return fmt.Errorf("Error reading url_filter_affinity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_filter_affinity: %v", err)
		}
	}

	if err = d.Set("url_filter_count", flattenSystemGlobalUrlFilterCount(o["url-filter-count"], d, "url_filter_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-filter-count"], "SystemGlobal-UrlFilterCount"); ok {
			if err = d.Set("url_filter_count", vv); err != nil {
				return fmt.Errorf("Error reading url_filter_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_filter_count: %v", err)
		}
	}

	if err = d.Set("user_device_store_max_devices", flattenSystemGlobalUserDeviceStoreMaxDevices(o["user-device-store-max-devices"], d, "user_device_store_max_devices")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-device-store-max-devices"], "SystemGlobal-UserDeviceStoreMaxDevices"); ok {
			if err = d.Set("user_device_store_max_devices", vv); err != nil {
				return fmt.Errorf("Error reading user_device_store_max_devices: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_device_store_max_devices: %v", err)
		}
	}

	if err = d.Set("user_device_store_max_unified_mem", flattenSystemGlobalUserDeviceStoreMaxUnifiedMem(o["user-device-store-max-unified-mem"], d, "user_device_store_max_unified_mem")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-device-store-max-unified-mem"], "SystemGlobal-UserDeviceStoreMaxUnifiedMem"); ok {
			if err = d.Set("user_device_store_max_unified_mem", vv); err != nil {
				return fmt.Errorf("Error reading user_device_store_max_unified_mem: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_device_store_max_unified_mem: %v", err)
		}
	}

	if err = d.Set("user_device_store_max_users", flattenSystemGlobalUserDeviceStoreMaxUsers(o["user-device-store-max-users"], d, "user_device_store_max_users")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-device-store-max-users"], "SystemGlobal-UserDeviceStoreMaxUsers"); ok {
			if err = d.Set("user_device_store_max_users", vv); err != nil {
				return fmt.Errorf("Error reading user_device_store_max_users: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_device_store_max_users: %v", err)
		}
	}

	if err = d.Set("user_server_cert", flattenSystemGlobalUserServerCert(o["user-server-cert"], d, "user_server_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-server-cert"], "SystemGlobal-UserServerCert"); ok {
			if err = d.Set("user_server_cert", vv); err != nil {
				return fmt.Errorf("Error reading user_server_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_server_cert: %v", err)
		}
	}

	if err = d.Set("vdom_mode", flattenSystemGlobalVdomMode(o["vdom-mode"], d, "vdom_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom-mode"], "SystemGlobal-VdomMode"); ok {
			if err = d.Set("vdom_mode", vv); err != nil {
				return fmt.Errorf("Error reading vdom_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom_mode: %v", err)
		}
	}

	if err = d.Set("vip_arp_range", flattenSystemGlobalVipArpRange(o["vip-arp-range"], d, "vip_arp_range")); err != nil {
		if vv, ok := fortiAPIPatch(o["vip-arp-range"], "SystemGlobal-VipArpRange"); ok {
			if err = d.Set("vip_arp_range", vv); err != nil {
				return fmt.Errorf("Error reading vip_arp_range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vip_arp_range: %v", err)
		}
	}

	if err = d.Set("virtual_server_count", flattenSystemGlobalVirtualServerCount(o["virtual-server-count"], d, "virtual_server_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["virtual-server-count"], "SystemGlobal-VirtualServerCount"); ok {
			if err = d.Set("virtual_server_count", vv); err != nil {
				return fmt.Errorf("Error reading virtual_server_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading virtual_server_count: %v", err)
		}
	}

	if err = d.Set("virtual_server_hardware_acceleration", flattenSystemGlobalVirtualServerHardwareAcceleration(o["virtual-server-hardware-acceleration"], d, "virtual_server_hardware_acceleration")); err != nil {
		if vv, ok := fortiAPIPatch(o["virtual-server-hardware-acceleration"], "SystemGlobal-VirtualServerHardwareAcceleration"); ok {
			if err = d.Set("virtual_server_hardware_acceleration", vv); err != nil {
				return fmt.Errorf("Error reading virtual_server_hardware_acceleration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading virtual_server_hardware_acceleration: %v", err)
		}
	}

	if err = d.Set("virtual_switch_vlan", flattenSystemGlobalVirtualSwitchVlan(o["virtual-switch-vlan"], d, "virtual_switch_vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["virtual-switch-vlan"], "SystemGlobal-VirtualSwitchVlan"); ok {
			if err = d.Set("virtual_switch_vlan", vv); err != nil {
				return fmt.Errorf("Error reading virtual_switch_vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading virtual_switch_vlan: %v", err)
		}
	}

	if err = d.Set("vpn_ems_sn_check", flattenSystemGlobalVpnEmsSnCheck(o["vpn-ems-sn-check"], d, "vpn_ems_sn_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["vpn-ems-sn-check"], "SystemGlobal-VpnEmsSnCheck"); ok {
			if err = d.Set("vpn_ems_sn_check", vv); err != nil {
				return fmt.Errorf("Error reading vpn_ems_sn_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vpn_ems_sn_check: %v", err)
		}
	}

	if err = d.Set("wad_affinity", flattenSystemGlobalWadAffinity(o["wad-affinity"], d, "wad_affinity")); err != nil {
		if vv, ok := fortiAPIPatch(o["wad-affinity"], "SystemGlobal-WadAffinity"); ok {
			if err = d.Set("wad_affinity", vv); err != nil {
				return fmt.Errorf("Error reading wad_affinity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wad_affinity: %v", err)
		}
	}

	if err = d.Set("wad_csvc_cs_count", flattenSystemGlobalWadCsvcCsCount(o["wad-csvc-cs-count"], d, "wad_csvc_cs_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["wad-csvc-cs-count"], "SystemGlobal-WadCsvcCsCount"); ok {
			if err = d.Set("wad_csvc_cs_count", vv); err != nil {
				return fmt.Errorf("Error reading wad_csvc_cs_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wad_csvc_cs_count: %v", err)
		}
	}

	if err = d.Set("wad_csvc_db_count", flattenSystemGlobalWadCsvcDbCount(o["wad-csvc-db-count"], d, "wad_csvc_db_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["wad-csvc-db-count"], "SystemGlobal-WadCsvcDbCount"); ok {
			if err = d.Set("wad_csvc_db_count", vv); err != nil {
				return fmt.Errorf("Error reading wad_csvc_db_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wad_csvc_db_count: %v", err)
		}
	}

	if err = d.Set("wad_memory_change_granularity", flattenSystemGlobalWadMemoryChangeGranularity(o["wad-memory-change-granularity"], d, "wad_memory_change_granularity")); err != nil {
		if vv, ok := fortiAPIPatch(o["wad-memory-change-granularity"], "SystemGlobal-WadMemoryChangeGranularity"); ok {
			if err = d.Set("wad_memory_change_granularity", vv); err != nil {
				return fmt.Errorf("Error reading wad_memory_change_granularity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wad_memory_change_granularity: %v", err)
		}
	}

	if err = d.Set("wad_restart_end_time", flattenSystemGlobalWadRestartEndTime(o["wad-restart-end-time"], d, "wad_restart_end_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["wad-restart-end-time"], "SystemGlobal-WadRestartEndTime"); ok {
			if err = d.Set("wad_restart_end_time", vv); err != nil {
				return fmt.Errorf("Error reading wad_restart_end_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wad_restart_end_time: %v", err)
		}
	}

	if err = d.Set("wad_restart_mode", flattenSystemGlobalWadRestartMode(o["wad-restart-mode"], d, "wad_restart_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["wad-restart-mode"], "SystemGlobal-WadRestartMode"); ok {
			if err = d.Set("wad_restart_mode", vv); err != nil {
				return fmt.Errorf("Error reading wad_restart_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wad_restart_mode: %v", err)
		}
	}

	if err = d.Set("wad_restart_start_time", flattenSystemGlobalWadRestartStartTime(o["wad-restart-start-time"], d, "wad_restart_start_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["wad-restart-start-time"], "SystemGlobal-WadRestartStartTime"); ok {
			if err = d.Set("wad_restart_start_time", vv); err != nil {
				return fmt.Errorf("Error reading wad_restart_start_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wad_restart_start_time: %v", err)
		}
	}

	if err = d.Set("wad_source_affinity", flattenSystemGlobalWadSourceAffinity(o["wad-source-affinity"], d, "wad_source_affinity")); err != nil {
		if vv, ok := fortiAPIPatch(o["wad-source-affinity"], "SystemGlobal-WadSourceAffinity"); ok {
			if err = d.Set("wad_source_affinity", vv); err != nil {
				return fmt.Errorf("Error reading wad_source_affinity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wad_source_affinity: %v", err)
		}
	}

	if err = d.Set("wad_worker_count", flattenSystemGlobalWadWorkerCount(o["wad-worker-count"], d, "wad_worker_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["wad-worker-count"], "SystemGlobal-WadWorkerCount"); ok {
			if err = d.Set("wad_worker_count", vv); err != nil {
				return fmt.Errorf("Error reading wad_worker_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wad_worker_count: %v", err)
		}
	}

	if err = d.Set("wifi_ca_certificate", flattenSystemGlobalWifiCaCertificate(o["wifi-ca-certificate"], d, "wifi_ca_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-ca-certificate"], "SystemGlobal-WifiCaCertificate"); ok {
			if err = d.Set("wifi_ca_certificate", vv); err != nil {
				return fmt.Errorf("Error reading wifi_ca_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_ca_certificate: %v", err)
		}
	}

	if err = d.Set("wifi_certificate", flattenSystemGlobalWifiCertificate(o["wifi-certificate"], d, "wifi_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-certificate"], "SystemGlobal-WifiCertificate"); ok {
			if err = d.Set("wifi_certificate", vv); err != nil {
				return fmt.Errorf("Error reading wifi_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_certificate: %v", err)
		}
	}

	if err = d.Set("wimax_4g_usb", flattenSystemGlobalWimax4GUsb(o["wimax-4g-usb"], d, "wimax_4g_usb")); err != nil {
		if vv, ok := fortiAPIPatch(o["wimax-4g-usb"], "SystemGlobal-Wimax4GUsb"); ok {
			if err = d.Set("wimax_4g_usb", vv); err != nil {
				return fmt.Errorf("Error reading wimax_4g_usb: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wimax_4g_usb: %v", err)
		}
	}

	if err = d.Set("wireless_controller", flattenSystemGlobalWirelessController(o["wireless-controller"], d, "wireless_controller")); err != nil {
		if vv, ok := fortiAPIPatch(o["wireless-controller"], "SystemGlobal-WirelessController"); ok {
			if err = d.Set("wireless_controller", vv); err != nil {
				return fmt.Errorf("Error reading wireless_controller: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wireless_controller: %v", err)
		}
	}

	if err = d.Set("wireless_controller_port", flattenSystemGlobalWirelessControllerPort(o["wireless-controller-port"], d, "wireless_controller_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["wireless-controller-port"], "SystemGlobal-WirelessControllerPort"); ok {
			if err = d.Set("wireless_controller_port", vv); err != nil {
				return fmt.Errorf("Error reading wireless_controller_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wireless_controller_port: %v", err)
		}
	}

	if err = d.Set("wireless_mode", flattenSystemGlobalWirelessMode(o["wireless-mode"], d, "wireless_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["wireless-mode"], "SystemGlobal-WirelessMode"); ok {
			if err = d.Set("wireless_mode", vv); err != nil {
				return fmt.Errorf("Error reading wireless_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wireless_mode: %v", err)
		}
	}

	if err = d.Set("xstools_update_frequency", flattenSystemGlobalXstoolsUpdateFrequency(o["xstools-update-frequency"], d, "xstools_update_frequency")); err != nil {
		if vv, ok := fortiAPIPatch(o["xstools-update-frequency"], "SystemGlobal-XstoolsUpdateFrequency"); ok {
			if err = d.Set("xstools_update_frequency", vv); err != nil {
				return fmt.Errorf("Error reading xstools_update_frequency: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading xstools_update_frequency: %v", err)
		}
	}

	return nil
}

func flattenSystemGlobalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemGlobalAdminBleButton(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminConcurrent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminConsoleTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminForticloudSsoDefaultProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalAdminForticloudSsoLogin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminHstsMaxAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminHttpsPkiRequired(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminHttpsRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminHttpsSslBannedCiphers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalAdminHttpsSslCiphersuites(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalAdminHttpsSslVersions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalAdminLockoutDuration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminLockoutThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminLoginMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminMaintainer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminResetButton(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminRestrictLocal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminScp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalAdminSport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminSshGraceTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminSshPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminSshPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminSshV1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminTelnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminTelnetPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdmintimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAirplaneMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAlias(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAllowTrafficRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAntiReplay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalArpMaxEntry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAuthCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalAuthHttpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAuthHttpsPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAuthIkeSamlPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAuthKeepalive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAuthSessionLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAutoAuthExtensionDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAutorunLogFsck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAvAffinity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAvFailopen(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAvFailopenSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalBatchCmdb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalBfdAffinity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalBlockSessionTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalBrFdbMaxEntry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCertChainMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCfgRevertTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCfgSave(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCheckProtocolHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCheckResetRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCliAuditLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCloudCommunication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCltCertReq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCmdbsvrAffinity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCpuUseThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCsrCaAttribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDailyRestart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDefaultServiceSourcePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDeviceIdleTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDhParams(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDhcpLeaseBackupInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDnsproxyWorkerCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDpFragmentTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDpPinholeTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDpRsyncTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDpTcpNormalTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDpUdpIdleTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalEarlyTcpNpuSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalEndpointControlFdsAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalEditVdomPrompt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalExtenderControllerReservedNetwork(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemGlobalFazDiskBufferSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFdsStatistics(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFdsStatisticsPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFecPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFgdAlertSubscription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalForticarrierBypass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalForticontrollerProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalForticontrollerProxyPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalForticonverterConfigUpload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalForticonverterIntegration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFortiextender(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFortiextenderDataPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFortiextenderDiscoveryLockdown(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFortiextenderProvisionOnAuthorization(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFortiextenderVlanMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFortiipamIntegration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFortigslbIntegration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFortiservicePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFortitokenCloud(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFortitokenCloudService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiAllowDefaultHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFortitokenCloudPushStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFortitokenCloudSyncInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiAllowIncompatibleFabricFgt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiAppDetectionSdwan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiAutoUpgradeSetupWarning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiCdnDomainOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiCdnUsage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiCertificates(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiCustomLanguage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiDateFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiDateTimeSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiDeviceLatitude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiDeviceLongitude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiDisplayHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiFirmwareUpgradeSetupWarning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiFirmwareUpgradeWarning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiForticareRegistrationSetupWarning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiFortisandboxCloud(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiFortigateCloudSandbox(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiFortiguardResourceFetch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiIpv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiLinesPerPage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiLocalOut(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiReplacementMessageGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiRestApiCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiTheme(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiWirelessOpensecurity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiWorkflowManagement(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalHaAffinity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalHonorDf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalHwSwitchEtherFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalHttpRequestLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalHttpUnauthenticatedRequestLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalHyperScaleVdomNum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalIgmpStateLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalInterfaceSubnetUsage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalInternalSwitchMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalInternalSwitchSpeed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalInternetServiceDatabase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalInternetServiceDownloadList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalIpFragmentMemThresholds(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalIpSrcPortRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalIpsAffinity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalIpsecAsicOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalIpsecHaSeqjumpRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalIpsecHmacOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalIpsecQatOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalIpsecRoundRobin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalIpsecSoftDecAsync(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalIpv6AcceptDad(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalIpv6AllowAnycastProbe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalIpv6AllowLocalInSilentDrop(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalIpv6AllowLocalInSlientDrop(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalIpv6AllowMulticastProbe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalIpv6AllowTrafficRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalIrqTimeAccounting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLanguage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLdapconntimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLegacyPoeDeviceSupport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLldpReception(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLldpTransmission(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLogSingleCpuHigh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLogSslConnection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLogUuidAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLogUuidPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLoginTimestamp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLongVdomName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalManagementIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalManagementPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalManagementPortUseAdminSport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalManagementVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalMaxRouteCacheSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalMemoryUseThresholdExtreme(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalMemoryUseThresholdGreen(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalMemoryUseThresholdRed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalMiglogAffinity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalMiglogdChildren(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalMultiFactorAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalNdpMaxEntry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalNpuNeighborUpdate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalOptimizeFlowMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPerUserBwl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPerUserBal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPmtuDiscovery(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPolicyAuthConcurrent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPostLoginBanner(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPreLoginBanner(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPrivateDataEncryption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalProxyAndExplicitProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalProxyAuthLifetime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalProxyAuthLifetimeTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalProxyAuthTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalProxyCipherHardwareAcceleration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalProxyKxpHardwareAcceleration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalProxyCertUseMgmtVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalProxyHardwareAcceleration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalProxyKeepAliveMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalProxyReAuthenticationMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalProxyReAuthenticationTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalProxyResourceMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalProxyWorkerCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPurdueLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalQsfp2840GPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalQsfpdd100GPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalQsfpddSplit8Port(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalQuicAckThresold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalQuicCongestionControlAlgo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalQuicMaxDatagramSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalQuicPmtud(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalQuicTlsHandshakeTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalQuicUdpPayloadSizeShapingPerCid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalRadiusPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalRebootUponConfigRestore(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalRefresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalRemoteauthtimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalResetSessionlessTcp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalRestartTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalRevisionBackupOnLogout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalRevisionImageAutoBackup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalScanunitCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSecurityRatingResultSubmission(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSecurityRatingRunOnSchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSendPmtuIcmp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSflowdMaxChildrenNum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalShowBackplaneIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSnatRouteChange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSpecialFile23Support(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSpeedtestServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSpeedtestdCtrlPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSpeedtestdServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSplitPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalSplitPortMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandSystemGlobalSplitPortModeInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["split-mode"], _ = expandSystemGlobalSplitPortModeSplitMode(d, i["split_mode"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemGlobalSplitPortModeInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSplitPortModeSplitMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSsdTrimDate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSsdTrimFreq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSsdTrimHour(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSsdTrimMin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSsdTrimWeekday(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSshCbcCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSshEncAlgo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalSshHmacMd5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSshHostkey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSshHostkeyAlgo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalSshHostkeyOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSshHostkeyPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalSshKexAlgo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalSshKexSha1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSshMacAlgo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalSshMacWeak(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSslStaticKeyCiphers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSslvpnCipherHardwareAcceleration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSslvpnEmsSnCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSslvpnKxpHardwareAcceleration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSslvpnMaxWorkerCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSslvpnPluginVersionCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSslvpnWebMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalStrictDirtySessionCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalStrongCrypto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSwitchController(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSwitchControllerReservedNetwork(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemGlobalSysFileCheckInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSysPerfLogInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSyslogAffinity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTcpHalfcloseTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTcpHalfopenTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTcpOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTcpRstTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTcpTimewaitTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTftp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTimezone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalTrafficPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTrafficPriorityLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTwoFactorEmailExpiry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTwoFactorFacExpiry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTwoFactorFtkExpiry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTwoFactorFtmExpiry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTwoFactorSmsExpiry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalUdpIdleTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalUrlFilterAffinity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalUrlFilterCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalUserDeviceStoreMaxDevices(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalUserDeviceStoreMaxUnifiedMem(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalUserDeviceStoreMaxUsers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalUserServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalVdomMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalVipArpRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalVirtualServerCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalVirtualServerHardwareAcceleration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalVirtualSwitchVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalVpnEmsSnCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalWadAffinity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalWadCsvcCsCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalWadCsvcDbCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalWadMemoryChangeGranularity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalWadRestartEndTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalWadRestartMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalWadRestartStartTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalWadSourceAffinity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalWadWorkerCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalWifiCaCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalWifiCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalWimax4GUsb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalWirelessController(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalWirelessControllerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalWirelessMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalXstoolsUpdateFrequency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemGlobal(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("admin_ble_button"); ok || d.HasChange("admin_ble_button") {
		t, err := expandSystemGlobalAdminBleButton(d, v, "admin_ble_button")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-ble-button"] = t
		}
	}

	if v, ok := d.GetOk("admin_concurrent"); ok || d.HasChange("admin_concurrent") {
		t, err := expandSystemGlobalAdminConcurrent(d, v, "admin_concurrent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-concurrent"] = t
		}
	}

	if v, ok := d.GetOk("admin_console_timeout"); ok || d.HasChange("admin_console_timeout") {
		t, err := expandSystemGlobalAdminConsoleTimeout(d, v, "admin_console_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-console-timeout"] = t
		}
	}

	if v, ok := d.GetOk("admin_forticloud_sso_default_profile"); ok || d.HasChange("admin_forticloud_sso_default_profile") {
		t, err := expandSystemGlobalAdminForticloudSsoDefaultProfile(d, v, "admin_forticloud_sso_default_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-forticloud-sso-default-profile"] = t
		}
	}

	if v, ok := d.GetOk("admin_forticloud_sso_login"); ok || d.HasChange("admin_forticloud_sso_login") {
		t, err := expandSystemGlobalAdminForticloudSsoLogin(d, v, "admin_forticloud_sso_login")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-forticloud-sso-login"] = t
		}
	}

	if v, ok := d.GetOk("admin_host"); ok || d.HasChange("admin_host") {
		t, err := expandSystemGlobalAdminHost(d, v, "admin_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-host"] = t
		}
	}

	if v, ok := d.GetOk("admin_hsts_max_age"); ok || d.HasChange("admin_hsts_max_age") {
		t, err := expandSystemGlobalAdminHstsMaxAge(d, v, "admin_hsts_max_age")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-hsts-max-age"] = t
		}
	}

	if v, ok := d.GetOk("admin_https_pki_required"); ok || d.HasChange("admin_https_pki_required") {
		t, err := expandSystemGlobalAdminHttpsPkiRequired(d, v, "admin_https_pki_required")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-https-pki-required"] = t
		}
	}

	if v, ok := d.GetOk("admin_https_redirect"); ok || d.HasChange("admin_https_redirect") {
		t, err := expandSystemGlobalAdminHttpsRedirect(d, v, "admin_https_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-https-redirect"] = t
		}
	}

	if v, ok := d.GetOk("admin_https_ssl_banned_ciphers"); ok || d.HasChange("admin_https_ssl_banned_ciphers") {
		t, err := expandSystemGlobalAdminHttpsSslBannedCiphers(d, v, "admin_https_ssl_banned_ciphers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-https-ssl-banned-ciphers"] = t
		}
	}

	if v, ok := d.GetOk("admin_https_ssl_ciphersuites"); ok || d.HasChange("admin_https_ssl_ciphersuites") {
		t, err := expandSystemGlobalAdminHttpsSslCiphersuites(d, v, "admin_https_ssl_ciphersuites")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-https-ssl-ciphersuites"] = t
		}
	}

	if v, ok := d.GetOk("admin_https_ssl_versions"); ok || d.HasChange("admin_https_ssl_versions") {
		t, err := expandSystemGlobalAdminHttpsSslVersions(d, v, "admin_https_ssl_versions")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-https-ssl-versions"] = t
		}
	}

	if v, ok := d.GetOk("admin_lockout_duration"); ok || d.HasChange("admin_lockout_duration") {
		t, err := expandSystemGlobalAdminLockoutDuration(d, v, "admin_lockout_duration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-lockout-duration"] = t
		}
	}

	if v, ok := d.GetOk("admin_lockout_threshold"); ok || d.HasChange("admin_lockout_threshold") {
		t, err := expandSystemGlobalAdminLockoutThreshold(d, v, "admin_lockout_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-lockout-threshold"] = t
		}
	}

	if v, ok := d.GetOk("admin_login_max"); ok || d.HasChange("admin_login_max") {
		t, err := expandSystemGlobalAdminLoginMax(d, v, "admin_login_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-login-max"] = t
		}
	}

	if v, ok := d.GetOk("admin_maintainer"); ok || d.HasChange("admin_maintainer") {
		t, err := expandSystemGlobalAdminMaintainer(d, v, "admin_maintainer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-maintainer"] = t
		}
	}

	if v, ok := d.GetOk("admin_port"); ok || d.HasChange("admin_port") {
		t, err := expandSystemGlobalAdminPort(d, v, "admin_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-port"] = t
		}
	}

	if v, ok := d.GetOk("admin_reset_button"); ok || d.HasChange("admin_reset_button") {
		t, err := expandSystemGlobalAdminResetButton(d, v, "admin_reset_button")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-reset-button"] = t
		}
	}

	if v, ok := d.GetOk("admin_restrict_local"); ok || d.HasChange("admin_restrict_local") {
		t, err := expandSystemGlobalAdminRestrictLocal(d, v, "admin_restrict_local")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-restrict-local"] = t
		}
	}

	if v, ok := d.GetOk("admin_scp"); ok || d.HasChange("admin_scp") {
		t, err := expandSystemGlobalAdminScp(d, v, "admin_scp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-scp"] = t
		}
	}

	if v, ok := d.GetOk("admin_server_cert"); ok || d.HasChange("admin_server_cert") {
		t, err := expandSystemGlobalAdminServerCert(d, v, "admin_server_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-server-cert"] = t
		}
	}

	if v, ok := d.GetOk("admin_sport"); ok || d.HasChange("admin_sport") {
		t, err := expandSystemGlobalAdminSport(d, v, "admin_sport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-sport"] = t
		}
	}

	if v, ok := d.GetOk("admin_ssh_grace_time"); ok || d.HasChange("admin_ssh_grace_time") {
		t, err := expandSystemGlobalAdminSshGraceTime(d, v, "admin_ssh_grace_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-ssh-grace-time"] = t
		}
	}

	if v, ok := d.GetOk("admin_ssh_password"); ok || d.HasChange("admin_ssh_password") {
		t, err := expandSystemGlobalAdminSshPassword(d, v, "admin_ssh_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-ssh-password"] = t
		}
	}

	if v, ok := d.GetOk("admin_ssh_port"); ok || d.HasChange("admin_ssh_port") {
		t, err := expandSystemGlobalAdminSshPort(d, v, "admin_ssh_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-ssh-port"] = t
		}
	}

	if v, ok := d.GetOk("admin_ssh_v1"); ok || d.HasChange("admin_ssh_v1") {
		t, err := expandSystemGlobalAdminSshV1(d, v, "admin_ssh_v1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-ssh-v1"] = t
		}
	}

	if v, ok := d.GetOk("admin_telnet"); ok || d.HasChange("admin_telnet") {
		t, err := expandSystemGlobalAdminTelnet(d, v, "admin_telnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-telnet"] = t
		}
	}

	if v, ok := d.GetOk("admin_telnet_port"); ok || d.HasChange("admin_telnet_port") {
		t, err := expandSystemGlobalAdminTelnetPort(d, v, "admin_telnet_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-telnet-port"] = t
		}
	}

	if v, ok := d.GetOk("admintimeout"); ok || d.HasChange("admintimeout") {
		t, err := expandSystemGlobalAdmintimeout(d, v, "admintimeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admintimeout"] = t
		}
	}

	if v, ok := d.GetOk("airplane_mode"); ok || d.HasChange("airplane_mode") {
		t, err := expandSystemGlobalAirplaneMode(d, v, "airplane_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["airplane-mode"] = t
		}
	}

	if v, ok := d.GetOk("alias"); ok || d.HasChange("alias") {
		t, err := expandSystemGlobalAlias(d, v, "alias")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alias"] = t
		}
	}

	if v, ok := d.GetOk("allow_traffic_redirect"); ok || d.HasChange("allow_traffic_redirect") {
		t, err := expandSystemGlobalAllowTrafficRedirect(d, v, "allow_traffic_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-traffic-redirect"] = t
		}
	}

	if v, ok := d.GetOk("anti_replay"); ok || d.HasChange("anti_replay") {
		t, err := expandSystemGlobalAntiReplay(d, v, "anti_replay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anti-replay"] = t
		}
	}

	if v, ok := d.GetOk("arp_max_entry"); ok || d.HasChange("arp_max_entry") {
		t, err := expandSystemGlobalArpMaxEntry(d, v, "arp_max_entry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arp-max-entry"] = t
		}
	}

	if v, ok := d.GetOk("auth_cert"); ok || d.HasChange("auth_cert") {
		t, err := expandSystemGlobalAuthCert(d, v, "auth_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-cert"] = t
		}
	}

	if v, ok := d.GetOk("auth_http_port"); ok || d.HasChange("auth_http_port") {
		t, err := expandSystemGlobalAuthHttpPort(d, v, "auth_http_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-http-port"] = t
		}
	}

	if v, ok := d.GetOk("auth_https_port"); ok || d.HasChange("auth_https_port") {
		t, err := expandSystemGlobalAuthHttpsPort(d, v, "auth_https_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-https-port"] = t
		}
	}

	if v, ok := d.GetOk("auth_ike_saml_port"); ok || d.HasChange("auth_ike_saml_port") {
		t, err := expandSystemGlobalAuthIkeSamlPort(d, v, "auth_ike_saml_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-ike-saml-port"] = t
		}
	}

	if v, ok := d.GetOk("auth_keepalive"); ok || d.HasChange("auth_keepalive") {
		t, err := expandSystemGlobalAuthKeepalive(d, v, "auth_keepalive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-keepalive"] = t
		}
	}

	if v, ok := d.GetOk("auth_session_limit"); ok || d.HasChange("auth_session_limit") {
		t, err := expandSystemGlobalAuthSessionLimit(d, v, "auth_session_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-session-limit"] = t
		}
	}

	if v, ok := d.GetOk("auto_auth_extension_device"); ok || d.HasChange("auto_auth_extension_device") {
		t, err := expandSystemGlobalAutoAuthExtensionDevice(d, v, "auto_auth_extension_device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-auth-extension-device"] = t
		}
	}

	if v, ok := d.GetOk("autorun_log_fsck"); ok || d.HasChange("autorun_log_fsck") {
		t, err := expandSystemGlobalAutorunLogFsck(d, v, "autorun_log_fsck")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["autorun-log-fsck"] = t
		}
	}

	if v, ok := d.GetOk("av_affinity"); ok || d.HasChange("av_affinity") {
		t, err := expandSystemGlobalAvAffinity(d, v, "av_affinity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-affinity"] = t
		}
	}

	if v, ok := d.GetOk("av_failopen"); ok || d.HasChange("av_failopen") {
		t, err := expandSystemGlobalAvFailopen(d, v, "av_failopen")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-failopen"] = t
		}
	}

	if v, ok := d.GetOk("av_failopen_session"); ok || d.HasChange("av_failopen_session") {
		t, err := expandSystemGlobalAvFailopenSession(d, v, "av_failopen_session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-failopen-session"] = t
		}
	}

	if v, ok := d.GetOk("batch_cmdb"); ok || d.HasChange("batch_cmdb") {
		t, err := expandSystemGlobalBatchCmdb(d, v, "batch_cmdb")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["batch-cmdb"] = t
		}
	}

	if v, ok := d.GetOk("bfd_affinity"); ok || d.HasChange("bfd_affinity") {
		t, err := expandSystemGlobalBfdAffinity(d, v, "bfd_affinity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd-affinity"] = t
		}
	}

	if v, ok := d.GetOk("block_session_timer"); ok || d.HasChange("block_session_timer") {
		t, err := expandSystemGlobalBlockSessionTimer(d, v, "block_session_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-session-timer"] = t
		}
	}

	if v, ok := d.GetOk("br_fdb_max_entry"); ok || d.HasChange("br_fdb_max_entry") {
		t, err := expandSystemGlobalBrFdbMaxEntry(d, v, "br_fdb_max_entry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["br-fdb-max-entry"] = t
		}
	}

	if v, ok := d.GetOk("cert_chain_max"); ok || d.HasChange("cert_chain_max") {
		t, err := expandSystemGlobalCertChainMax(d, v, "cert_chain_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-chain-max"] = t
		}
	}

	if v, ok := d.GetOk("cfg_revert_timeout"); ok || d.HasChange("cfg_revert_timeout") {
		t, err := expandSystemGlobalCfgRevertTimeout(d, v, "cfg_revert_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cfg-revert-timeout"] = t
		}
	}

	if v, ok := d.GetOk("cfg_save"); ok || d.HasChange("cfg_save") {
		t, err := expandSystemGlobalCfgSave(d, v, "cfg_save")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cfg-save"] = t
		}
	}

	if v, ok := d.GetOk("check_protocol_header"); ok || d.HasChange("check_protocol_header") {
		t, err := expandSystemGlobalCheckProtocolHeader(d, v, "check_protocol_header")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["check-protocol-header"] = t
		}
	}

	if v, ok := d.GetOk("check_reset_range"); ok || d.HasChange("check_reset_range") {
		t, err := expandSystemGlobalCheckResetRange(d, v, "check_reset_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["check-reset-range"] = t
		}
	}

	if v, ok := d.GetOk("cli_audit_log"); ok || d.HasChange("cli_audit_log") {
		t, err := expandSystemGlobalCliAuditLog(d, v, "cli_audit_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cli-audit-log"] = t
		}
	}

	if v, ok := d.GetOk("cloud_communication"); ok || d.HasChange("cloud_communication") {
		t, err := expandSystemGlobalCloudCommunication(d, v, "cloud_communication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cloud-communication"] = t
		}
	}

	if v, ok := d.GetOk("clt_cert_req"); ok || d.HasChange("clt_cert_req") {
		t, err := expandSystemGlobalCltCertReq(d, v, "clt_cert_req")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clt-cert-req"] = t
		}
	}

	if v, ok := d.GetOk("cmdbsvr_affinity"); ok || d.HasChange("cmdbsvr_affinity") {
		t, err := expandSystemGlobalCmdbsvrAffinity(d, v, "cmdbsvr_affinity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cmdbsvr-affinity"] = t
		}
	}

	if v, ok := d.GetOk("cpu_use_threshold"); ok || d.HasChange("cpu_use_threshold") {
		t, err := expandSystemGlobalCpuUseThreshold(d, v, "cpu_use_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cpu-use-threshold"] = t
		}
	}

	if v, ok := d.GetOk("csr_ca_attribute"); ok || d.HasChange("csr_ca_attribute") {
		t, err := expandSystemGlobalCsrCaAttribute(d, v, "csr_ca_attribute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["csr-ca-attribute"] = t
		}
	}

	if v, ok := d.GetOk("daily_restart"); ok || d.HasChange("daily_restart") {
		t, err := expandSystemGlobalDailyRestart(d, v, "daily_restart")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["daily-restart"] = t
		}
	}

	if v, ok := d.GetOk("default_service_source_port"); ok || d.HasChange("default_service_source_port") {
		t, err := expandSystemGlobalDefaultServiceSourcePort(d, v, "default_service_source_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-service-source-port"] = t
		}
	}

	if v, ok := d.GetOk("device_idle_timeout"); ok || d.HasChange("device_idle_timeout") {
		t, err := expandSystemGlobalDeviceIdleTimeout(d, v, "device_idle_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-idle-timeout"] = t
		}
	}

	if v, ok := d.GetOk("dh_params"); ok || d.HasChange("dh_params") {
		t, err := expandSystemGlobalDhParams(d, v, "dh_params")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-params"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_lease_backup_interval"); ok || d.HasChange("dhcp_lease_backup_interval") {
		t, err := expandSystemGlobalDhcpLeaseBackupInterval(d, v, "dhcp_lease_backup_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-lease-backup-interval"] = t
		}
	}

	if v, ok := d.GetOk("dnsproxy_worker_count"); ok || d.HasChange("dnsproxy_worker_count") {
		t, err := expandSystemGlobalDnsproxyWorkerCount(d, v, "dnsproxy_worker_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dnsproxy-worker-count"] = t
		}
	}

	if v, ok := d.GetOk("dp_fragment_timer"); ok || d.HasChange("dp_fragment_timer") {
		t, err := expandSystemGlobalDpFragmentTimer(d, v, "dp_fragment_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dp-fragment-timer"] = t
		}
	}

	if v, ok := d.GetOk("dp_pinhole_timer"); ok || d.HasChange("dp_pinhole_timer") {
		t, err := expandSystemGlobalDpPinholeTimer(d, v, "dp_pinhole_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dp-pinhole-timer"] = t
		}
	}

	if v, ok := d.GetOk("dp_rsync_timer"); ok || d.HasChange("dp_rsync_timer") {
		t, err := expandSystemGlobalDpRsyncTimer(d, v, "dp_rsync_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dp-rsync-timer"] = t
		}
	}

	if v, ok := d.GetOk("dp_tcp_normal_timer"); ok || d.HasChange("dp_tcp_normal_timer") {
		t, err := expandSystemGlobalDpTcpNormalTimer(d, v, "dp_tcp_normal_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dp-tcp-normal-timer"] = t
		}
	}

	if v, ok := d.GetOk("dp_udp_idle_timer"); ok || d.HasChange("dp_udp_idle_timer") {
		t, err := expandSystemGlobalDpUdpIdleTimer(d, v, "dp_udp_idle_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dp-udp-idle-timer"] = t
		}
	}

	if v, ok := d.GetOk("dst"); ok || d.HasChange("dst") {
		t, err := expandSystemGlobalDst(d, v, "dst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("early_tcp_npu_session"); ok || d.HasChange("early_tcp_npu_session") {
		t, err := expandSystemGlobalEarlyTcpNpuSession(d, v, "early_tcp_npu_session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["early-tcp-npu-session"] = t
		}
	}

	if v, ok := d.GetOk("endpoint_control_fds_access"); ok || d.HasChange("endpoint_control_fds_access") {
		t, err := expandSystemGlobalEndpointControlFdsAccess(d, v, "endpoint_control_fds_access")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["endpoint-control-fds-access"] = t
		}
	}

	if v, ok := d.GetOk("edit_vdom_prompt"); ok || d.HasChange("edit_vdom_prompt") {
		t, err := expandSystemGlobalEditVdomPrompt(d, v, "edit_vdom_prompt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["edit-vdom-prompt"] = t
		}
	}

	if v, ok := d.GetOk("extender_controller_reserved_network"); ok || d.HasChange("extender_controller_reserved_network") {
		t, err := expandSystemGlobalExtenderControllerReservedNetwork(d, v, "extender_controller_reserved_network")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extender-controller-reserved-network"] = t
		}
	}

	if v, ok := d.GetOk("faz_disk_buffer_size"); ok || d.HasChange("faz_disk_buffer_size") {
		t, err := expandSystemGlobalFazDiskBufferSize(d, v, "faz_disk_buffer_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-disk-buffer-size"] = t
		}
	}

	if v, ok := d.GetOk("fds_statistics"); ok || d.HasChange("fds_statistics") {
		t, err := expandSystemGlobalFdsStatistics(d, v, "fds_statistics")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fds-statistics"] = t
		}
	}

	if v, ok := d.GetOk("fds_statistics_period"); ok || d.HasChange("fds_statistics_period") {
		t, err := expandSystemGlobalFdsStatisticsPeriod(d, v, "fds_statistics_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fds-statistics-period"] = t
		}
	}

	if v, ok := d.GetOk("fec_port"); ok || d.HasChange("fec_port") {
		t, err := expandSystemGlobalFecPort(d, v, "fec_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-port"] = t
		}
	}

	if v, ok := d.GetOk("fgd_alert_subscription"); ok || d.HasChange("fgd_alert_subscription") {
		t, err := expandSystemGlobalFgdAlertSubscription(d, v, "fgd_alert_subscription")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgd-alert-subscription"] = t
		}
	}

	if v, ok := d.GetOk("forticarrier_bypass"); ok || d.HasChange("forticarrier_bypass") {
		t, err := expandSystemGlobalForticarrierBypass(d, v, "forticarrier_bypass")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticarrier-bypass"] = t
		}
	}

	if v, ok := d.GetOk("forticontroller_proxy"); ok || d.HasChange("forticontroller_proxy") {
		t, err := expandSystemGlobalForticontrollerProxy(d, v, "forticontroller_proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticontroller-proxy"] = t
		}
	}

	if v, ok := d.GetOk("forticontroller_proxy_port"); ok || d.HasChange("forticontroller_proxy_port") {
		t, err := expandSystemGlobalForticontrollerProxyPort(d, v, "forticontroller_proxy_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticontroller-proxy-port"] = t
		}
	}

	if v, ok := d.GetOk("forticonverter_config_upload"); ok || d.HasChange("forticonverter_config_upload") {
		t, err := expandSystemGlobalForticonverterConfigUpload(d, v, "forticonverter_config_upload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticonverter-config-upload"] = t
		}
	}

	if v, ok := d.GetOk("forticonverter_integration"); ok || d.HasChange("forticonverter_integration") {
		t, err := expandSystemGlobalForticonverterIntegration(d, v, "forticonverter_integration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticonverter-integration"] = t
		}
	}

	if v, ok := d.GetOk("fortiextender"); ok || d.HasChange("fortiextender") {
		t, err := expandSystemGlobalFortiextender(d, v, "fortiextender")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiextender"] = t
		}
	}

	if v, ok := d.GetOk("fortiextender_data_port"); ok || d.HasChange("fortiextender_data_port") {
		t, err := expandSystemGlobalFortiextenderDataPort(d, v, "fortiextender_data_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiextender-data-port"] = t
		}
	}

	if v, ok := d.GetOk("fortiextender_discovery_lockdown"); ok || d.HasChange("fortiextender_discovery_lockdown") {
		t, err := expandSystemGlobalFortiextenderDiscoveryLockdown(d, v, "fortiextender_discovery_lockdown")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiextender-discovery-lockdown"] = t
		}
	}

	if v, ok := d.GetOk("fortiextender_provision_on_authorization"); ok || d.HasChange("fortiextender_provision_on_authorization") {
		t, err := expandSystemGlobalFortiextenderProvisionOnAuthorization(d, v, "fortiextender_provision_on_authorization")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiextender-provision-on-authorization"] = t
		}
	}

	if v, ok := d.GetOk("fortiextender_vlan_mode"); ok || d.HasChange("fortiextender_vlan_mode") {
		t, err := expandSystemGlobalFortiextenderVlanMode(d, v, "fortiextender_vlan_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiextender-vlan-mode"] = t
		}
	}

	if v, ok := d.GetOk("fortiipam_integration"); ok || d.HasChange("fortiipam_integration") {
		t, err := expandSystemGlobalFortiipamIntegration(d, v, "fortiipam_integration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiipam-integration"] = t
		}
	}

	if v, ok := d.GetOk("fortigslb_integration"); ok || d.HasChange("fortigslb_integration") {
		t, err := expandSystemGlobalFortigslbIntegration(d, v, "fortigslb_integration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortigslb-integration"] = t
		}
	}

	if v, ok := d.GetOk("fortiservice_port"); ok || d.HasChange("fortiservice_port") {
		t, err := expandSystemGlobalFortiservicePort(d, v, "fortiservice_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiservice-port"] = t
		}
	}

	if v, ok := d.GetOk("fortitoken_cloud"); ok || d.HasChange("fortitoken_cloud") {
		t, err := expandSystemGlobalFortitokenCloud(d, v, "fortitoken_cloud")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortitoken-cloud"] = t
		}
	}

	if v, ok := d.GetOk("fortitoken_cloud_service"); ok || d.HasChange("fortitoken_cloud_service") {
		t, err := expandSystemGlobalFortitokenCloudService(d, v, "fortitoken_cloud_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortitoken-cloud-service"] = t
		}
	}

	if v, ok := d.GetOk("gui_allow_default_hostname"); ok || d.HasChange("gui_allow_default_hostname") {
		t, err := expandSystemGlobalGuiAllowDefaultHostname(d, v, "gui_allow_default_hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-allow-default-hostname"] = t
		}
	}

	if v, ok := d.GetOk("fortitoken_cloud_push_status"); ok || d.HasChange("fortitoken_cloud_push_status") {
		t, err := expandSystemGlobalFortitokenCloudPushStatus(d, v, "fortitoken_cloud_push_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortitoken-cloud-push-status"] = t
		}
	}

	if v, ok := d.GetOk("fortitoken_cloud_sync_interval"); ok || d.HasChange("fortitoken_cloud_sync_interval") {
		t, err := expandSystemGlobalFortitokenCloudSyncInterval(d, v, "fortitoken_cloud_sync_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortitoken-cloud-sync-interval"] = t
		}
	}

	if v, ok := d.GetOk("gui_allow_incompatible_fabric_fgt"); ok || d.HasChange("gui_allow_incompatible_fabric_fgt") {
		t, err := expandSystemGlobalGuiAllowIncompatibleFabricFgt(d, v, "gui_allow_incompatible_fabric_fgt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-allow-incompatible-fabric-fgt"] = t
		}
	}

	if v, ok := d.GetOk("gui_app_detection_sdwan"); ok || d.HasChange("gui_app_detection_sdwan") {
		t, err := expandSystemGlobalGuiAppDetectionSdwan(d, v, "gui_app_detection_sdwan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-app-detection-sdwan"] = t
		}
	}

	if v, ok := d.GetOk("gui_auto_upgrade_setup_warning"); ok || d.HasChange("gui_auto_upgrade_setup_warning") {
		t, err := expandSystemGlobalGuiAutoUpgradeSetupWarning(d, v, "gui_auto_upgrade_setup_warning")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-auto-upgrade-setup-warning"] = t
		}
	}

	if v, ok := d.GetOk("gui_cdn_domain_override"); ok || d.HasChange("gui_cdn_domain_override") {
		t, err := expandSystemGlobalGuiCdnDomainOverride(d, v, "gui_cdn_domain_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-cdn-domain-override"] = t
		}
	}

	if v, ok := d.GetOk("gui_cdn_usage"); ok || d.HasChange("gui_cdn_usage") {
		t, err := expandSystemGlobalGuiCdnUsage(d, v, "gui_cdn_usage")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-cdn-usage"] = t
		}
	}

	if v, ok := d.GetOk("gui_certificates"); ok || d.HasChange("gui_certificates") {
		t, err := expandSystemGlobalGuiCertificates(d, v, "gui_certificates")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-certificates"] = t
		}
	}

	if v, ok := d.GetOk("gui_custom_language"); ok || d.HasChange("gui_custom_language") {
		t, err := expandSystemGlobalGuiCustomLanguage(d, v, "gui_custom_language")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-custom-language"] = t
		}
	}

	if v, ok := d.GetOk("gui_date_format"); ok || d.HasChange("gui_date_format") {
		t, err := expandSystemGlobalGuiDateFormat(d, v, "gui_date_format")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-date-format"] = t
		}
	}

	if v, ok := d.GetOk("gui_date_time_source"); ok || d.HasChange("gui_date_time_source") {
		t, err := expandSystemGlobalGuiDateTimeSource(d, v, "gui_date_time_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-date-time-source"] = t
		}
	}

	if v, ok := d.GetOk("gui_device_latitude"); ok || d.HasChange("gui_device_latitude") {
		t, err := expandSystemGlobalGuiDeviceLatitude(d, v, "gui_device_latitude")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-device-latitude"] = t
		}
	}

	if v, ok := d.GetOk("gui_device_longitude"); ok || d.HasChange("gui_device_longitude") {
		t, err := expandSystemGlobalGuiDeviceLongitude(d, v, "gui_device_longitude")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-device-longitude"] = t
		}
	}

	if v, ok := d.GetOk("gui_display_hostname"); ok || d.HasChange("gui_display_hostname") {
		t, err := expandSystemGlobalGuiDisplayHostname(d, v, "gui_display_hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-display-hostname"] = t
		}
	}

	if v, ok := d.GetOk("gui_firmware_upgrade_setup_warning"); ok || d.HasChange("gui_firmware_upgrade_setup_warning") {
		t, err := expandSystemGlobalGuiFirmwareUpgradeSetupWarning(d, v, "gui_firmware_upgrade_setup_warning")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-firmware-upgrade-setup-warning"] = t
		}
	}

	if v, ok := d.GetOk("gui_firmware_upgrade_warning"); ok || d.HasChange("gui_firmware_upgrade_warning") {
		t, err := expandSystemGlobalGuiFirmwareUpgradeWarning(d, v, "gui_firmware_upgrade_warning")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-firmware-upgrade-warning"] = t
		}
	}

	if v, ok := d.GetOk("gui_forticare_registration_setup_warning"); ok || d.HasChange("gui_forticare_registration_setup_warning") {
		t, err := expandSystemGlobalGuiForticareRegistrationSetupWarning(d, v, "gui_forticare_registration_setup_warning")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-forticare-registration-setup-warning"] = t
		}
	}

	if v, ok := d.GetOk("gui_fortisandbox_cloud"); ok || d.HasChange("gui_fortisandbox_cloud") {
		t, err := expandSystemGlobalGuiFortisandboxCloud(d, v, "gui_fortisandbox_cloud")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-fortisandbox-cloud"] = t
		}
	}

	if v, ok := d.GetOk("gui_fortigate_cloud_sandbox"); ok || d.HasChange("gui_fortigate_cloud_sandbox") {
		t, err := expandSystemGlobalGuiFortigateCloudSandbox(d, v, "gui_fortigate_cloud_sandbox")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-fortigate-cloud-sandbox"] = t
		}
	}

	if v, ok := d.GetOk("gui_fortiguard_resource_fetch"); ok || d.HasChange("gui_fortiguard_resource_fetch") {
		t, err := expandSystemGlobalGuiFortiguardResourceFetch(d, v, "gui_fortiguard_resource_fetch")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-fortiguard-resource-fetch"] = t
		}
	}

	if v, ok := d.GetOk("gui_ipv6"); ok || d.HasChange("gui_ipv6") {
		t, err := expandSystemGlobalGuiIpv6(d, v, "gui_ipv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-ipv6"] = t
		}
	}

	if v, ok := d.GetOk("gui_lines_per_page"); ok || d.HasChange("gui_lines_per_page") {
		t, err := expandSystemGlobalGuiLinesPerPage(d, v, "gui_lines_per_page")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-lines-per-page"] = t
		}
	}

	if v, ok := d.GetOk("gui_local_out"); ok || d.HasChange("gui_local_out") {
		t, err := expandSystemGlobalGuiLocalOut(d, v, "gui_local_out")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-local-out"] = t
		}
	}

	if v, ok := d.GetOk("gui_replacement_message_groups"); ok || d.HasChange("gui_replacement_message_groups") {
		t, err := expandSystemGlobalGuiReplacementMessageGroups(d, v, "gui_replacement_message_groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-replacement-message-groups"] = t
		}
	}

	if v, ok := d.GetOk("gui_rest_api_cache"); ok || d.HasChange("gui_rest_api_cache") {
		t, err := expandSystemGlobalGuiRestApiCache(d, v, "gui_rest_api_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-rest-api-cache"] = t
		}
	}

	if v, ok := d.GetOk("gui_theme"); ok || d.HasChange("gui_theme") {
		t, err := expandSystemGlobalGuiTheme(d, v, "gui_theme")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-theme"] = t
		}
	}

	if v, ok := d.GetOk("gui_wireless_opensecurity"); ok || d.HasChange("gui_wireless_opensecurity") {
		t, err := expandSystemGlobalGuiWirelessOpensecurity(d, v, "gui_wireless_opensecurity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-wireless-opensecurity"] = t
		}
	}

	if v, ok := d.GetOk("gui_workflow_management"); ok || d.HasChange("gui_workflow_management") {
		t, err := expandSystemGlobalGuiWorkflowManagement(d, v, "gui_workflow_management")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-workflow-management"] = t
		}
	}

	if v, ok := d.GetOk("ha_affinity"); ok || d.HasChange("ha_affinity") {
		t, err := expandSystemGlobalHaAffinity(d, v, "ha_affinity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-affinity"] = t
		}
	}

	if v, ok := d.GetOk("honor_df"); ok || d.HasChange("honor_df") {
		t, err := expandSystemGlobalHonorDf(d, v, "honor_df")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["honor-df"] = t
		}
	}

	if v, ok := d.GetOk("hostname"); ok || d.HasChange("hostname") {
		t, err := expandSystemGlobalHostname(d, v, "hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostname"] = t
		}
	}

	if v, ok := d.GetOk("hw_switch_ether_filter"); ok || d.HasChange("hw_switch_ether_filter") {
		t, err := expandSystemGlobalHwSwitchEtherFilter(d, v, "hw_switch_ether_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hw-switch-ether-filter"] = t
		}
	}

	if v, ok := d.GetOk("http_request_limit"); ok || d.HasChange("http_request_limit") {
		t, err := expandSystemGlobalHttpRequestLimit(d, v, "http_request_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-request-limit"] = t
		}
	}

	if v, ok := d.GetOk("http_unauthenticated_request_limit"); ok || d.HasChange("http_unauthenticated_request_limit") {
		t, err := expandSystemGlobalHttpUnauthenticatedRequestLimit(d, v, "http_unauthenticated_request_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-unauthenticated-request-limit"] = t
		}
	}

	if v, ok := d.GetOk("hyper_scale_vdom_num"); ok || d.HasChange("hyper_scale_vdom_num") {
		t, err := expandSystemGlobalHyperScaleVdomNum(d, v, "hyper_scale_vdom_num")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hyper-scale-vdom-num"] = t
		}
	}

	if v, ok := d.GetOk("igmp_state_limit"); ok || d.HasChange("igmp_state_limit") {
		t, err := expandSystemGlobalIgmpStateLimit(d, v, "igmp_state_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["igmp-state-limit"] = t
		}
	}

	if v, ok := d.GetOk("interface_subnet_usage"); ok || d.HasChange("interface_subnet_usage") {
		t, err := expandSystemGlobalInterfaceSubnetUsage(d, v, "interface_subnet_usage")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-subnet-usage"] = t
		}
	}

	if v, ok := d.GetOk("internal_switch_mode"); ok || d.HasChange("internal_switch_mode") {
		t, err := expandSystemGlobalInternalSwitchMode(d, v, "internal_switch_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internal-switch-mode"] = t
		}
	}

	if v, ok := d.GetOk("internal_switch_speed"); ok || d.HasChange("internal_switch_speed") {
		t, err := expandSystemGlobalInternalSwitchSpeed(d, v, "internal_switch_speed")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internal-switch-speed"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_database"); ok || d.HasChange("internet_service_database") {
		t, err := expandSystemGlobalInternetServiceDatabase(d, v, "internet_service_database")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-database"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_download_list"); ok || d.HasChange("internet_service_download_list") {
		t, err := expandSystemGlobalInternetServiceDownloadList(d, v, "internet_service_download_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-download-list"] = t
		}
	}

	if v, ok := d.GetOk("ip_fragment_mem_thresholds"); ok || d.HasChange("ip_fragment_mem_thresholds") {
		t, err := expandSystemGlobalIpFragmentMemThresholds(d, v, "ip_fragment_mem_thresholds")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-fragment-mem-thresholds"] = t
		}
	}

	if v, ok := d.GetOk("ip_src_port_range"); ok || d.HasChange("ip_src_port_range") {
		t, err := expandSystemGlobalIpSrcPortRange(d, v, "ip_src_port_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-src-port-range"] = t
		}
	}

	if v, ok := d.GetOk("ips_affinity"); ok || d.HasChange("ips_affinity") {
		t, err := expandSystemGlobalIpsAffinity(d, v, "ips_affinity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-affinity"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_asic_offload"); ok || d.HasChange("ipsec_asic_offload") {
		t, err := expandSystemGlobalIpsecAsicOffload(d, v, "ipsec_asic_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-asic-offload"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_ha_seqjump_rate"); ok || d.HasChange("ipsec_ha_seqjump_rate") {
		t, err := expandSystemGlobalIpsecHaSeqjumpRate(d, v, "ipsec_ha_seqjump_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-ha-seqjump-rate"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_hmac_offload"); ok || d.HasChange("ipsec_hmac_offload") {
		t, err := expandSystemGlobalIpsecHmacOffload(d, v, "ipsec_hmac_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-hmac-offload"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_qat_offload"); ok || d.HasChange("ipsec_qat_offload") {
		t, err := expandSystemGlobalIpsecQatOffload(d, v, "ipsec_qat_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-qat-offload"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_round_robin"); ok || d.HasChange("ipsec_round_robin") {
		t, err := expandSystemGlobalIpsecRoundRobin(d, v, "ipsec_round_robin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-round-robin"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_soft_dec_async"); ok || d.HasChange("ipsec_soft_dec_async") {
		t, err := expandSystemGlobalIpsecSoftDecAsync(d, v, "ipsec_soft_dec_async")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-soft-dec-async"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_accept_dad"); ok || d.HasChange("ipv6_accept_dad") {
		t, err := expandSystemGlobalIpv6AcceptDad(d, v, "ipv6_accept_dad")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-accept-dad"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_allow_anycast_probe"); ok || d.HasChange("ipv6_allow_anycast_probe") {
		t, err := expandSystemGlobalIpv6AllowAnycastProbe(d, v, "ipv6_allow_anycast_probe")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-allow-anycast-probe"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_allow_local_in_silent_drop"); ok || d.HasChange("ipv6_allow_local_in_silent_drop") {
		t, err := expandSystemGlobalIpv6AllowLocalInSilentDrop(d, v, "ipv6_allow_local_in_silent_drop")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-allow-local-in-silent-drop"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_allow_local_in_slient_drop"); ok || d.HasChange("ipv6_allow_local_in_slient_drop") {
		t, err := expandSystemGlobalIpv6AllowLocalInSlientDrop(d, v, "ipv6_allow_local_in_slient_drop")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-allow-local-in-slient-drop"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_allow_multicast_probe"); ok || d.HasChange("ipv6_allow_multicast_probe") {
		t, err := expandSystemGlobalIpv6AllowMulticastProbe(d, v, "ipv6_allow_multicast_probe")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-allow-multicast-probe"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_allow_traffic_redirect"); ok || d.HasChange("ipv6_allow_traffic_redirect") {
		t, err := expandSystemGlobalIpv6AllowTrafficRedirect(d, v, "ipv6_allow_traffic_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-allow-traffic-redirect"] = t
		}
	}

	if v, ok := d.GetOk("irq_time_accounting"); ok || d.HasChange("irq_time_accounting") {
		t, err := expandSystemGlobalIrqTimeAccounting(d, v, "irq_time_accounting")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["irq-time-accounting"] = t
		}
	}

	if v, ok := d.GetOk("language"); ok || d.HasChange("language") {
		t, err := expandSystemGlobalLanguage(d, v, "language")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["language"] = t
		}
	}

	if v, ok := d.GetOk("ldapconntimeout"); ok || d.HasChange("ldapconntimeout") {
		t, err := expandSystemGlobalLdapconntimeout(d, v, "ldapconntimeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldapconntimeout"] = t
		}
	}

	if v, ok := d.GetOk("legacy_poe_device_support"); ok || d.HasChange("legacy_poe_device_support") {
		t, err := expandSystemGlobalLegacyPoeDeviceSupport(d, v, "legacy_poe_device_support")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["legacy-poe-device-support"] = t
		}
	}

	if v, ok := d.GetOk("lldp_reception"); ok || d.HasChange("lldp_reception") {
		t, err := expandSystemGlobalLldpReception(d, v, "lldp_reception")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lldp-reception"] = t
		}
	}

	if v, ok := d.GetOk("lldp_transmission"); ok || d.HasChange("lldp_transmission") {
		t, err := expandSystemGlobalLldpTransmission(d, v, "lldp_transmission")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lldp-transmission"] = t
		}
	}

	if v, ok := d.GetOk("log_single_cpu_high"); ok || d.HasChange("log_single_cpu_high") {
		t, err := expandSystemGlobalLogSingleCpuHigh(d, v, "log_single_cpu_high")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-single-cpu-high"] = t
		}
	}

	if v, ok := d.GetOk("log_ssl_connection"); ok || d.HasChange("log_ssl_connection") {
		t, err := expandSystemGlobalLogSslConnection(d, v, "log_ssl_connection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-ssl-connection"] = t
		}
	}

	if v, ok := d.GetOk("log_uuid_address"); ok || d.HasChange("log_uuid_address") {
		t, err := expandSystemGlobalLogUuidAddress(d, v, "log_uuid_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-uuid-address"] = t
		}
	}

	if v, ok := d.GetOk("log_uuid_policy"); ok || d.HasChange("log_uuid_policy") {
		t, err := expandSystemGlobalLogUuidPolicy(d, v, "log_uuid_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-uuid-policy"] = t
		}
	}

	if v, ok := d.GetOk("login_timestamp"); ok || d.HasChange("login_timestamp") {
		t, err := expandSystemGlobalLoginTimestamp(d, v, "login_timestamp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-timestamp"] = t
		}
	}

	if v, ok := d.GetOk("long_vdom_name"); ok || d.HasChange("long_vdom_name") {
		t, err := expandSystemGlobalLongVdomName(d, v, "long_vdom_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["long-vdom-name"] = t
		}
	}

	if v, ok := d.GetOk("management_ip"); ok || d.HasChange("management_ip") {
		t, err := expandSystemGlobalManagementIp(d, v, "management_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["management-ip"] = t
		}
	}

	if v, ok := d.GetOk("management_port"); ok || d.HasChange("management_port") {
		t, err := expandSystemGlobalManagementPort(d, v, "management_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["management-port"] = t
		}
	}

	if v, ok := d.GetOk("management_port_use_admin_sport"); ok || d.HasChange("management_port_use_admin_sport") {
		t, err := expandSystemGlobalManagementPortUseAdminSport(d, v, "management_port_use_admin_sport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["management-port-use-admin-sport"] = t
		}
	}

	if v, ok := d.GetOk("management_vdom"); ok || d.HasChange("management_vdom") {
		t, err := expandSystemGlobalManagementVdom(d, v, "management_vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["management-vdom"] = t
		}
	}

	if v, ok := d.GetOk("max_route_cache_size"); ok || d.HasChange("max_route_cache_size") {
		t, err := expandSystemGlobalMaxRouteCacheSize(d, v, "max_route_cache_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-route-cache-size"] = t
		}
	}

	if v, ok := d.GetOk("memory_use_threshold_extreme"); ok || d.HasChange("memory_use_threshold_extreme") {
		t, err := expandSystemGlobalMemoryUseThresholdExtreme(d, v, "memory_use_threshold_extreme")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["memory-use-threshold-extreme"] = t
		}
	}

	if v, ok := d.GetOk("memory_use_threshold_green"); ok || d.HasChange("memory_use_threshold_green") {
		t, err := expandSystemGlobalMemoryUseThresholdGreen(d, v, "memory_use_threshold_green")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["memory-use-threshold-green"] = t
		}
	}

	if v, ok := d.GetOk("memory_use_threshold_red"); ok || d.HasChange("memory_use_threshold_red") {
		t, err := expandSystemGlobalMemoryUseThresholdRed(d, v, "memory_use_threshold_red")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["memory-use-threshold-red"] = t
		}
	}

	if v, ok := d.GetOk("miglog_affinity"); ok || d.HasChange("miglog_affinity") {
		t, err := expandSystemGlobalMiglogAffinity(d, v, "miglog_affinity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["miglog-affinity"] = t
		}
	}

	if v, ok := d.GetOk("miglogd_children"); ok || d.HasChange("miglogd_children") {
		t, err := expandSystemGlobalMiglogdChildren(d, v, "miglogd_children")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["miglogd-children"] = t
		}
	}

	if v, ok := d.GetOk("multi_factor_authentication"); ok || d.HasChange("multi_factor_authentication") {
		t, err := expandSystemGlobalMultiFactorAuthentication(d, v, "multi_factor_authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multi-factor-authentication"] = t
		}
	}

	if v, ok := d.GetOk("ndp_max_entry"); ok || d.HasChange("ndp_max_entry") {
		t, err := expandSystemGlobalNdpMaxEntry(d, v, "ndp_max_entry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ndp-max-entry"] = t
		}
	}

	if v, ok := d.GetOk("npu_neighbor_update"); ok || d.HasChange("npu_neighbor_update") {
		t, err := expandSystemGlobalNpuNeighborUpdate(d, v, "npu_neighbor_update")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["npu-neighbor-update"] = t
		}
	}

	if v, ok := d.GetOk("optimize_flow_mode"); ok || d.HasChange("optimize_flow_mode") {
		t, err := expandSystemGlobalOptimizeFlowMode(d, v, "optimize_flow_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["optimize-flow-mode"] = t
		}
	}

	if v, ok := d.GetOk("per_user_bwl"); ok || d.HasChange("per_user_bwl") {
		t, err := expandSystemGlobalPerUserBwl(d, v, "per_user_bwl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-user-bwl"] = t
		}
	}

	if v, ok := d.GetOk("per_user_bal"); ok || d.HasChange("per_user_bal") {
		t, err := expandSystemGlobalPerUserBal(d, v, "per_user_bal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-user-bal"] = t
		}
	}

	if v, ok := d.GetOk("pmtu_discovery"); ok || d.HasChange("pmtu_discovery") {
		t, err := expandSystemGlobalPmtuDiscovery(d, v, "pmtu_discovery")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pmtu-discovery"] = t
		}
	}

	if v, ok := d.GetOk("policy_auth_concurrent"); ok || d.HasChange("policy_auth_concurrent") {
		t, err := expandSystemGlobalPolicyAuthConcurrent(d, v, "policy_auth_concurrent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-auth-concurrent"] = t
		}
	}

	if v, ok := d.GetOk("post_login_banner"); ok || d.HasChange("post_login_banner") {
		t, err := expandSystemGlobalPostLoginBanner(d, v, "post_login_banner")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["post-login-banner"] = t
		}
	}

	if v, ok := d.GetOk("pre_login_banner"); ok || d.HasChange("pre_login_banner") {
		t, err := expandSystemGlobalPreLoginBanner(d, v, "pre_login_banner")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pre-login-banner"] = t
		}
	}

	if v, ok := d.GetOk("private_data_encryption"); ok || d.HasChange("private_data_encryption") {
		t, err := expandSystemGlobalPrivateDataEncryption(d, v, "private_data_encryption")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-data-encryption"] = t
		}
	}

	if v, ok := d.GetOk("proxy_and_explicit_proxy"); ok || d.HasChange("proxy_and_explicit_proxy") {
		t, err := expandSystemGlobalProxyAndExplicitProxy(d, v, "proxy_and_explicit_proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-and-explicit-proxy"] = t
		}
	}

	if v, ok := d.GetOk("proxy_auth_lifetime"); ok || d.HasChange("proxy_auth_lifetime") {
		t, err := expandSystemGlobalProxyAuthLifetime(d, v, "proxy_auth_lifetime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-auth-lifetime"] = t
		}
	}

	if v, ok := d.GetOk("proxy_auth_lifetime_timeout"); ok || d.HasChange("proxy_auth_lifetime_timeout") {
		t, err := expandSystemGlobalProxyAuthLifetimeTimeout(d, v, "proxy_auth_lifetime_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-auth-lifetime-timeout"] = t
		}
	}

	if v, ok := d.GetOk("proxy_auth_timeout"); ok || d.HasChange("proxy_auth_timeout") {
		t, err := expandSystemGlobalProxyAuthTimeout(d, v, "proxy_auth_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-auth-timeout"] = t
		}
	}

	if v, ok := d.GetOk("proxy_cipher_hardware_acceleration"); ok || d.HasChange("proxy_cipher_hardware_acceleration") {
		t, err := expandSystemGlobalProxyCipherHardwareAcceleration(d, v, "proxy_cipher_hardware_acceleration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-cipher-hardware-acceleration"] = t
		}
	}

	if v, ok := d.GetOk("proxy_kxp_hardware_acceleration"); ok || d.HasChange("proxy_kxp_hardware_acceleration") {
		t, err := expandSystemGlobalProxyKxpHardwareAcceleration(d, v, "proxy_kxp_hardware_acceleration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-kxp-hardware-acceleration"] = t
		}
	}

	if v, ok := d.GetOk("proxy_cert_use_mgmt_vdom"); ok || d.HasChange("proxy_cert_use_mgmt_vdom") {
		t, err := expandSystemGlobalProxyCertUseMgmtVdom(d, v, "proxy_cert_use_mgmt_vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-cert-use-mgmt-vdom"] = t
		}
	}

	if v, ok := d.GetOk("proxy_hardware_acceleration"); ok || d.HasChange("proxy_hardware_acceleration") {
		t, err := expandSystemGlobalProxyHardwareAcceleration(d, v, "proxy_hardware_acceleration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-hardware-acceleration"] = t
		}
	}

	if v, ok := d.GetOk("proxy_keep_alive_mode"); ok || d.HasChange("proxy_keep_alive_mode") {
		t, err := expandSystemGlobalProxyKeepAliveMode(d, v, "proxy_keep_alive_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-keep-alive-mode"] = t
		}
	}

	if v, ok := d.GetOk("proxy_re_authentication_mode"); ok || d.HasChange("proxy_re_authentication_mode") {
		t, err := expandSystemGlobalProxyReAuthenticationMode(d, v, "proxy_re_authentication_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-re-authentication-mode"] = t
		}
	}

	if v, ok := d.GetOk("proxy_re_authentication_time"); ok || d.HasChange("proxy_re_authentication_time") {
		t, err := expandSystemGlobalProxyReAuthenticationTime(d, v, "proxy_re_authentication_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-re-authentication-time"] = t
		}
	}

	if v, ok := d.GetOk("proxy_resource_mode"); ok || d.HasChange("proxy_resource_mode") {
		t, err := expandSystemGlobalProxyResourceMode(d, v, "proxy_resource_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-resource-mode"] = t
		}
	}

	if v, ok := d.GetOk("proxy_worker_count"); ok || d.HasChange("proxy_worker_count") {
		t, err := expandSystemGlobalProxyWorkerCount(d, v, "proxy_worker_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-worker-count"] = t
		}
	}

	if v, ok := d.GetOk("purdue_level"); ok || d.HasChange("purdue_level") {
		t, err := expandSystemGlobalPurdueLevel(d, v, "purdue_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["purdue-level"] = t
		}
	}

	if v, ok := d.GetOk("qsfp28_40g_port"); ok || d.HasChange("qsfp28_40g_port") {
		t, err := expandSystemGlobalQsfp2840GPort(d, v, "qsfp28_40g_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qsfp28-40g-port"] = t
		}
	}

	if v, ok := d.GetOk("qsfpdd_100g_port"); ok || d.HasChange("qsfpdd_100g_port") {
		t, err := expandSystemGlobalQsfpdd100GPort(d, v, "qsfpdd_100g_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qsfpdd-100g-port"] = t
		}
	}

	if v, ok := d.GetOk("qsfpdd_split8_port"); ok || d.HasChange("qsfpdd_split8_port") {
		t, err := expandSystemGlobalQsfpddSplit8Port(d, v, "qsfpdd_split8_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qsfpdd-split8-port"] = t
		}
	}

	if v, ok := d.GetOk("quic_ack_thresold"); ok || d.HasChange("quic_ack_thresold") {
		t, err := expandSystemGlobalQuicAckThresold(d, v, "quic_ack_thresold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quic-ack-thresold"] = t
		}
	}

	if v, ok := d.GetOk("quic_congestion_control_algo"); ok || d.HasChange("quic_congestion_control_algo") {
		t, err := expandSystemGlobalQuicCongestionControlAlgo(d, v, "quic_congestion_control_algo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quic-congestion-control-algo"] = t
		}
	}

	if v, ok := d.GetOk("quic_max_datagram_size"); ok || d.HasChange("quic_max_datagram_size") {
		t, err := expandSystemGlobalQuicMaxDatagramSize(d, v, "quic_max_datagram_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quic-max-datagram-size"] = t
		}
	}

	if v, ok := d.GetOk("quic_pmtud"); ok || d.HasChange("quic_pmtud") {
		t, err := expandSystemGlobalQuicPmtud(d, v, "quic_pmtud")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quic-pmtud"] = t
		}
	}

	if v, ok := d.GetOk("quic_tls_handshake_timeout"); ok || d.HasChange("quic_tls_handshake_timeout") {
		t, err := expandSystemGlobalQuicTlsHandshakeTimeout(d, v, "quic_tls_handshake_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quic-tls-handshake-timeout"] = t
		}
	}

	if v, ok := d.GetOk("quic_udp_payload_size_shaping_per_cid"); ok || d.HasChange("quic_udp_payload_size_shaping_per_cid") {
		t, err := expandSystemGlobalQuicUdpPayloadSizeShapingPerCid(d, v, "quic_udp_payload_size_shaping_per_cid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quic-udp-payload-size-shaping-per-cid"] = t
		}
	}

	if v, ok := d.GetOk("radius_port"); ok || d.HasChange("radius_port") {
		t, err := expandSystemGlobalRadiusPort(d, v, "radius_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-port"] = t
		}
	}

	if v, ok := d.GetOk("reboot_upon_config_restore"); ok || d.HasChange("reboot_upon_config_restore") {
		t, err := expandSystemGlobalRebootUponConfigRestore(d, v, "reboot_upon_config_restore")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reboot-upon-config-restore"] = t
		}
	}

	if v, ok := d.GetOk("refresh"); ok || d.HasChange("refresh") {
		t, err := expandSystemGlobalRefresh(d, v, "refresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["refresh"] = t
		}
	}

	if v, ok := d.GetOk("remoteauthtimeout"); ok || d.HasChange("remoteauthtimeout") {
		t, err := expandSystemGlobalRemoteauthtimeout(d, v, "remoteauthtimeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remoteauthtimeout"] = t
		}
	}

	if v, ok := d.GetOk("reset_sessionless_tcp"); ok || d.HasChange("reset_sessionless_tcp") {
		t, err := expandSystemGlobalResetSessionlessTcp(d, v, "reset_sessionless_tcp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reset-sessionless-tcp"] = t
		}
	}

	if v, ok := d.GetOk("restart_time"); ok || d.HasChange("restart_time") {
		t, err := expandSystemGlobalRestartTime(d, v, "restart_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restart-time"] = t
		}
	}

	if v, ok := d.GetOk("revision_backup_on_logout"); ok || d.HasChange("revision_backup_on_logout") {
		t, err := expandSystemGlobalRevisionBackupOnLogout(d, v, "revision_backup_on_logout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["revision-backup-on-logout"] = t
		}
	}

	if v, ok := d.GetOk("revision_image_auto_backup"); ok || d.HasChange("revision_image_auto_backup") {
		t, err := expandSystemGlobalRevisionImageAutoBackup(d, v, "revision_image_auto_backup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["revision-image-auto-backup"] = t
		}
	}

	if v, ok := d.GetOk("scanunit_count"); ok || d.HasChange("scanunit_count") {
		t, err := expandSystemGlobalScanunitCount(d, v, "scanunit_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scanunit-count"] = t
		}
	}

	if v, ok := d.GetOk("security_rating_result_submission"); ok || d.HasChange("security_rating_result_submission") {
		t, err := expandSystemGlobalSecurityRatingResultSubmission(d, v, "security_rating_result_submission")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-rating-result-submission"] = t
		}
	}

	if v, ok := d.GetOk("security_rating_run_on_schedule"); ok || d.HasChange("security_rating_run_on_schedule") {
		t, err := expandSystemGlobalSecurityRatingRunOnSchedule(d, v, "security_rating_run_on_schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-rating-run-on-schedule"] = t
		}
	}

	if v, ok := d.GetOk("send_pmtu_icmp"); ok || d.HasChange("send_pmtu_icmp") {
		t, err := expandSystemGlobalSendPmtuIcmp(d, v, "send_pmtu_icmp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-pmtu-icmp"] = t
		}
	}

	if v, ok := d.GetOk("sflowd_max_children_num"); ok || d.HasChange("sflowd_max_children_num") {
		t, err := expandSystemGlobalSflowdMaxChildrenNum(d, v, "sflowd_max_children_num")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sflowd-max-children-num"] = t
		}
	}

	if v, ok := d.GetOk("show_backplane_intf"); ok || d.HasChange("show_backplane_intf") {
		t, err := expandSystemGlobalShowBackplaneIntf(d, v, "show_backplane_intf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["show-backplane-intf"] = t
		}
	}

	if v, ok := d.GetOk("snat_route_change"); ok || d.HasChange("snat_route_change") {
		t, err := expandSystemGlobalSnatRouteChange(d, v, "snat_route_change")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["snat-route-change"] = t
		}
	}

	if v, ok := d.GetOk("special_file_23_support"); ok || d.HasChange("special_file_23_support") {
		t, err := expandSystemGlobalSpecialFile23Support(d, v, "special_file_23_support")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["special-file-23-support"] = t
		}
	}

	if v, ok := d.GetOk("speedtest_server"); ok || d.HasChange("speedtest_server") {
		t, err := expandSystemGlobalSpeedtestServer(d, v, "speedtest_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["speedtest-server"] = t
		}
	}

	if v, ok := d.GetOk("speedtestd_ctrl_port"); ok || d.HasChange("speedtestd_ctrl_port") {
		t, err := expandSystemGlobalSpeedtestdCtrlPort(d, v, "speedtestd_ctrl_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["speedtestd-ctrl-port"] = t
		}
	}

	if v, ok := d.GetOk("speedtestd_server_port"); ok || d.HasChange("speedtestd_server_port") {
		t, err := expandSystemGlobalSpeedtestdServerPort(d, v, "speedtestd_server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["speedtestd-server-port"] = t
		}
	}

	if v, ok := d.GetOk("split_port"); ok || d.HasChange("split_port") {
		t, err := expandSystemGlobalSplitPort(d, v, "split_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-port"] = t
		}
	}

	if v, ok := d.GetOk("split_port_mode"); ok || d.HasChange("split_port_mode") {
		t, err := expandSystemGlobalSplitPortMode(d, v, "split_port_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-port-mode"] = t
		}
	}

	if v, ok := d.GetOk("ssd_trim_date"); ok || d.HasChange("ssd_trim_date") {
		t, err := expandSystemGlobalSsdTrimDate(d, v, "ssd_trim_date")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssd-trim-date"] = t
		}
	}

	if v, ok := d.GetOk("ssd_trim_freq"); ok || d.HasChange("ssd_trim_freq") {
		t, err := expandSystemGlobalSsdTrimFreq(d, v, "ssd_trim_freq")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssd-trim-freq"] = t
		}
	}

	if v, ok := d.GetOk("ssd_trim_hour"); ok || d.HasChange("ssd_trim_hour") {
		t, err := expandSystemGlobalSsdTrimHour(d, v, "ssd_trim_hour")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssd-trim-hour"] = t
		}
	}

	if v, ok := d.GetOk("ssd_trim_min"); ok || d.HasChange("ssd_trim_min") {
		t, err := expandSystemGlobalSsdTrimMin(d, v, "ssd_trim_min")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssd-trim-min"] = t
		}
	}

	if v, ok := d.GetOk("ssd_trim_weekday"); ok || d.HasChange("ssd_trim_weekday") {
		t, err := expandSystemGlobalSsdTrimWeekday(d, v, "ssd_trim_weekday")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssd-trim-weekday"] = t
		}
	}

	if v, ok := d.GetOk("ssh_cbc_cipher"); ok || d.HasChange("ssh_cbc_cipher") {
		t, err := expandSystemGlobalSshCbcCipher(d, v, "ssh_cbc_cipher")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-cbc-cipher"] = t
		}
	}

	if v, ok := d.GetOk("ssh_enc_algo"); ok || d.HasChange("ssh_enc_algo") {
		t, err := expandSystemGlobalSshEncAlgo(d, v, "ssh_enc_algo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-enc-algo"] = t
		}
	}

	if v, ok := d.GetOk("ssh_hmac_md5"); ok || d.HasChange("ssh_hmac_md5") {
		t, err := expandSystemGlobalSshHmacMd5(d, v, "ssh_hmac_md5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-hmac-md5"] = t
		}
	}

	if v, ok := d.GetOk("ssh_hostkey"); ok || d.HasChange("ssh_hostkey") {
		t, err := expandSystemGlobalSshHostkey(d, v, "ssh_hostkey")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-hostkey"] = t
		}
	}

	if v, ok := d.GetOk("ssh_hostkey_algo"); ok || d.HasChange("ssh_hostkey_algo") {
		t, err := expandSystemGlobalSshHostkeyAlgo(d, v, "ssh_hostkey_algo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-hostkey-algo"] = t
		}
	}

	if v, ok := d.GetOk("ssh_hostkey_override"); ok || d.HasChange("ssh_hostkey_override") {
		t, err := expandSystemGlobalSshHostkeyOverride(d, v, "ssh_hostkey_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-hostkey-override"] = t
		}
	}

	if v, ok := d.GetOk("ssh_hostkey_password"); ok || d.HasChange("ssh_hostkey_password") {
		t, err := expandSystemGlobalSshHostkeyPassword(d, v, "ssh_hostkey_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-hostkey-password"] = t
		}
	}

	if v, ok := d.GetOk("ssh_kex_algo"); ok || d.HasChange("ssh_kex_algo") {
		t, err := expandSystemGlobalSshKexAlgo(d, v, "ssh_kex_algo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-kex-algo"] = t
		}
	}

	if v, ok := d.GetOk("ssh_kex_sha1"); ok || d.HasChange("ssh_kex_sha1") {
		t, err := expandSystemGlobalSshKexSha1(d, v, "ssh_kex_sha1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-kex-sha1"] = t
		}
	}

	if v, ok := d.GetOk("ssh_mac_algo"); ok || d.HasChange("ssh_mac_algo") {
		t, err := expandSystemGlobalSshMacAlgo(d, v, "ssh_mac_algo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-mac-algo"] = t
		}
	}

	if v, ok := d.GetOk("ssh_mac_weak"); ok || d.HasChange("ssh_mac_weak") {
		t, err := expandSystemGlobalSshMacWeak(d, v, "ssh_mac_weak")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-mac-weak"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_version"); ok || d.HasChange("ssl_min_proto_version") {
		t, err := expandSystemGlobalSslMinProtoVersion(d, v, "ssl_min_proto_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-proto-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_static_key_ciphers"); ok || d.HasChange("ssl_static_key_ciphers") {
		t, err := expandSystemGlobalSslStaticKeyCiphers(d, v, "ssl_static_key_ciphers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-static-key-ciphers"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_cipher_hardware_acceleration"); ok || d.HasChange("sslvpn_cipher_hardware_acceleration") {
		t, err := expandSystemGlobalSslvpnCipherHardwareAcceleration(d, v, "sslvpn_cipher_hardware_acceleration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-cipher-hardware-acceleration"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_ems_sn_check"); ok || d.HasChange("sslvpn_ems_sn_check") {
		t, err := expandSystemGlobalSslvpnEmsSnCheck(d, v, "sslvpn_ems_sn_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-ems-sn-check"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_kxp_hardware_acceleration"); ok || d.HasChange("sslvpn_kxp_hardware_acceleration") {
		t, err := expandSystemGlobalSslvpnKxpHardwareAcceleration(d, v, "sslvpn_kxp_hardware_acceleration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-kxp-hardware-acceleration"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_max_worker_count"); ok || d.HasChange("sslvpn_max_worker_count") {
		t, err := expandSystemGlobalSslvpnMaxWorkerCount(d, v, "sslvpn_max_worker_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-max-worker-count"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_plugin_version_check"); ok || d.HasChange("sslvpn_plugin_version_check") {
		t, err := expandSystemGlobalSslvpnPluginVersionCheck(d, v, "sslvpn_plugin_version_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-plugin-version-check"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_web_mode"); ok || d.HasChange("sslvpn_web_mode") {
		t, err := expandSystemGlobalSslvpnWebMode(d, v, "sslvpn_web_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-web-mode"] = t
		}
	}

	if v, ok := d.GetOk("strict_dirty_session_check"); ok || d.HasChange("strict_dirty_session_check") {
		t, err := expandSystemGlobalStrictDirtySessionCheck(d, v, "strict_dirty_session_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strict-dirty-session-check"] = t
		}
	}

	if v, ok := d.GetOk("strong_crypto"); ok || d.HasChange("strong_crypto") {
		t, err := expandSystemGlobalStrongCrypto(d, v, "strong_crypto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strong-crypto"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller"); ok || d.HasChange("switch_controller") {
		t, err := expandSystemGlobalSwitchController(d, v, "switch_controller")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_reserved_network"); ok || d.HasChange("switch_controller_reserved_network") {
		t, err := expandSystemGlobalSwitchControllerReservedNetwork(d, v, "switch_controller_reserved_network")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-reserved-network"] = t
		}
	}

	if v, ok := d.GetOk("sys_file_check_interval"); ok || d.HasChange("sys_file_check_interval") {
		t, err := expandSystemGlobalSysFileCheckInterval(d, v, "sys_file_check_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sys-file-check-interval"] = t
		}
	}

	if v, ok := d.GetOk("sys_perf_log_interval"); ok || d.HasChange("sys_perf_log_interval") {
		t, err := expandSystemGlobalSysPerfLogInterval(d, v, "sys_perf_log_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sys-perf-log-interval"] = t
		}
	}

	if v, ok := d.GetOk("syslog_affinity"); ok || d.HasChange("syslog_affinity") {
		t, err := expandSystemGlobalSyslogAffinity(d, v, "syslog_affinity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["syslog-affinity"] = t
		}
	}

	if v, ok := d.GetOk("tcp_halfclose_timer"); ok || d.HasChange("tcp_halfclose_timer") {
		t, err := expandSystemGlobalTcpHalfcloseTimer(d, v, "tcp_halfclose_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-halfclose-timer"] = t
		}
	}

	if v, ok := d.GetOk("tcp_halfopen_timer"); ok || d.HasChange("tcp_halfopen_timer") {
		t, err := expandSystemGlobalTcpHalfopenTimer(d, v, "tcp_halfopen_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-halfopen-timer"] = t
		}
	}

	if v, ok := d.GetOk("tcp_option"); ok || d.HasChange("tcp_option") {
		t, err := expandSystemGlobalTcpOption(d, v, "tcp_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-option"] = t
		}
	}

	if v, ok := d.GetOk("tcp_rst_timer"); ok || d.HasChange("tcp_rst_timer") {
		t, err := expandSystemGlobalTcpRstTimer(d, v, "tcp_rst_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-rst-timer"] = t
		}
	}

	if v, ok := d.GetOk("tcp_timewait_timer"); ok || d.HasChange("tcp_timewait_timer") {
		t, err := expandSystemGlobalTcpTimewaitTimer(d, v, "tcp_timewait_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-timewait-timer"] = t
		}
	}

	if v, ok := d.GetOk("tftp"); ok || d.HasChange("tftp") {
		t, err := expandSystemGlobalTftp(d, v, "tftp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tftp"] = t
		}
	}

	if v, ok := d.GetOk("timezone"); ok || d.HasChange("timezone") {
		t, err := expandSystemGlobalTimezone(d, v, "timezone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timezone"] = t
		}
	}

	if v, ok := d.GetOk("traffic_priority"); ok || d.HasChange("traffic_priority") {
		t, err := expandSystemGlobalTrafficPriority(d, v, "traffic_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-priority"] = t
		}
	}

	if v, ok := d.GetOk("traffic_priority_level"); ok || d.HasChange("traffic_priority_level") {
		t, err := expandSystemGlobalTrafficPriorityLevel(d, v, "traffic_priority_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-priority-level"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_email_expiry"); ok || d.HasChange("two_factor_email_expiry") {
		t, err := expandSystemGlobalTwoFactorEmailExpiry(d, v, "two_factor_email_expiry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-email-expiry"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_fac_expiry"); ok || d.HasChange("two_factor_fac_expiry") {
		t, err := expandSystemGlobalTwoFactorFacExpiry(d, v, "two_factor_fac_expiry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-fac-expiry"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_ftk_expiry"); ok || d.HasChange("two_factor_ftk_expiry") {
		t, err := expandSystemGlobalTwoFactorFtkExpiry(d, v, "two_factor_ftk_expiry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-ftk-expiry"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_ftm_expiry"); ok || d.HasChange("two_factor_ftm_expiry") {
		t, err := expandSystemGlobalTwoFactorFtmExpiry(d, v, "two_factor_ftm_expiry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-ftm-expiry"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_sms_expiry"); ok || d.HasChange("two_factor_sms_expiry") {
		t, err := expandSystemGlobalTwoFactorSmsExpiry(d, v, "two_factor_sms_expiry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-sms-expiry"] = t
		}
	}

	if v, ok := d.GetOk("udp_idle_timer"); ok || d.HasChange("udp_idle_timer") {
		t, err := expandSystemGlobalUdpIdleTimer(d, v, "udp_idle_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-idle-timer"] = t
		}
	}

	if v, ok := d.GetOk("url_filter_affinity"); ok || d.HasChange("url_filter_affinity") {
		t, err := expandSystemGlobalUrlFilterAffinity(d, v, "url_filter_affinity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-filter-affinity"] = t
		}
	}

	if v, ok := d.GetOk("url_filter_count"); ok || d.HasChange("url_filter_count") {
		t, err := expandSystemGlobalUrlFilterCount(d, v, "url_filter_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-filter-count"] = t
		}
	}

	if v, ok := d.GetOk("user_device_store_max_devices"); ok || d.HasChange("user_device_store_max_devices") {
		t, err := expandSystemGlobalUserDeviceStoreMaxDevices(d, v, "user_device_store_max_devices")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-device-store-max-devices"] = t
		}
	}

	if v, ok := d.GetOk("user_device_store_max_unified_mem"); ok || d.HasChange("user_device_store_max_unified_mem") {
		t, err := expandSystemGlobalUserDeviceStoreMaxUnifiedMem(d, v, "user_device_store_max_unified_mem")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-device-store-max-unified-mem"] = t
		}
	}

	if v, ok := d.GetOk("user_device_store_max_users"); ok || d.HasChange("user_device_store_max_users") {
		t, err := expandSystemGlobalUserDeviceStoreMaxUsers(d, v, "user_device_store_max_users")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-device-store-max-users"] = t
		}
	}

	if v, ok := d.GetOk("user_server_cert"); ok || d.HasChange("user_server_cert") {
		t, err := expandSystemGlobalUserServerCert(d, v, "user_server_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-server-cert"] = t
		}
	}

	if v, ok := d.GetOk("vdom_mode"); ok || d.HasChange("vdom_mode") {
		t, err := expandSystemGlobalVdomMode(d, v, "vdom_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom-mode"] = t
		}
	}

	if v, ok := d.GetOk("vip_arp_range"); ok || d.HasChange("vip_arp_range") {
		t, err := expandSystemGlobalVipArpRange(d, v, "vip_arp_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vip-arp-range"] = t
		}
	}

	if v, ok := d.GetOk("virtual_server_count"); ok || d.HasChange("virtual_server_count") {
		t, err := expandSystemGlobalVirtualServerCount(d, v, "virtual_server_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-server-count"] = t
		}
	}

	if v, ok := d.GetOk("virtual_server_hardware_acceleration"); ok || d.HasChange("virtual_server_hardware_acceleration") {
		t, err := expandSystemGlobalVirtualServerHardwareAcceleration(d, v, "virtual_server_hardware_acceleration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-server-hardware-acceleration"] = t
		}
	}

	if v, ok := d.GetOk("virtual_switch_vlan"); ok || d.HasChange("virtual_switch_vlan") {
		t, err := expandSystemGlobalVirtualSwitchVlan(d, v, "virtual_switch_vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-switch-vlan"] = t
		}
	}

	if v, ok := d.GetOk("vpn_ems_sn_check"); ok || d.HasChange("vpn_ems_sn_check") {
		t, err := expandSystemGlobalVpnEmsSnCheck(d, v, "vpn_ems_sn_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpn-ems-sn-check"] = t
		}
	}

	if v, ok := d.GetOk("wad_affinity"); ok || d.HasChange("wad_affinity") {
		t, err := expandSystemGlobalWadAffinity(d, v, "wad_affinity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wad-affinity"] = t
		}
	}

	if v, ok := d.GetOk("wad_csvc_cs_count"); ok || d.HasChange("wad_csvc_cs_count") {
		t, err := expandSystemGlobalWadCsvcCsCount(d, v, "wad_csvc_cs_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wad-csvc-cs-count"] = t
		}
	}

	if v, ok := d.GetOk("wad_csvc_db_count"); ok || d.HasChange("wad_csvc_db_count") {
		t, err := expandSystemGlobalWadCsvcDbCount(d, v, "wad_csvc_db_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wad-csvc-db-count"] = t
		}
	}

	if v, ok := d.GetOk("wad_memory_change_granularity"); ok || d.HasChange("wad_memory_change_granularity") {
		t, err := expandSystemGlobalWadMemoryChangeGranularity(d, v, "wad_memory_change_granularity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wad-memory-change-granularity"] = t
		}
	}

	if v, ok := d.GetOk("wad_restart_end_time"); ok || d.HasChange("wad_restart_end_time") {
		t, err := expandSystemGlobalWadRestartEndTime(d, v, "wad_restart_end_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wad-restart-end-time"] = t
		}
	}

	if v, ok := d.GetOk("wad_restart_mode"); ok || d.HasChange("wad_restart_mode") {
		t, err := expandSystemGlobalWadRestartMode(d, v, "wad_restart_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wad-restart-mode"] = t
		}
	}

	if v, ok := d.GetOk("wad_restart_start_time"); ok || d.HasChange("wad_restart_start_time") {
		t, err := expandSystemGlobalWadRestartStartTime(d, v, "wad_restart_start_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wad-restart-start-time"] = t
		}
	}

	if v, ok := d.GetOk("wad_source_affinity"); ok || d.HasChange("wad_source_affinity") {
		t, err := expandSystemGlobalWadSourceAffinity(d, v, "wad_source_affinity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wad-source-affinity"] = t
		}
	}

	if v, ok := d.GetOk("wad_worker_count"); ok || d.HasChange("wad_worker_count") {
		t, err := expandSystemGlobalWadWorkerCount(d, v, "wad_worker_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wad-worker-count"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ca_certificate"); ok || d.HasChange("wifi_ca_certificate") {
		t, err := expandSystemGlobalWifiCaCertificate(d, v, "wifi_ca_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ca-certificate"] = t
		}
	}

	if v, ok := d.GetOk("wifi_certificate"); ok || d.HasChange("wifi_certificate") {
		t, err := expandSystemGlobalWifiCertificate(d, v, "wifi_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-certificate"] = t
		}
	}

	if v, ok := d.GetOk("wimax_4g_usb"); ok || d.HasChange("wimax_4g_usb") {
		t, err := expandSystemGlobalWimax4GUsb(d, v, "wimax_4g_usb")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wimax-4g-usb"] = t
		}
	}

	if v, ok := d.GetOk("wireless_controller"); ok || d.HasChange("wireless_controller") {
		t, err := expandSystemGlobalWirelessController(d, v, "wireless_controller")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wireless-controller"] = t
		}
	}

	if v, ok := d.GetOk("wireless_controller_port"); ok || d.HasChange("wireless_controller_port") {
		t, err := expandSystemGlobalWirelessControllerPort(d, v, "wireless_controller_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wireless-controller-port"] = t
		}
	}

	if v, ok := d.GetOk("wireless_mode"); ok || d.HasChange("wireless_mode") {
		t, err := expandSystemGlobalWirelessMode(d, v, "wireless_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wireless-mode"] = t
		}
	}

	if v, ok := d.GetOk("xstools_update_frequency"); ok || d.HasChange("xstools_update_frequency") {
		t, err := expandSystemGlobalXstoolsUpdateFrequency(d, v, "xstools_update_frequency")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["xstools-update-frequency"] = t
		}
	}

	return &obj, nil
}
