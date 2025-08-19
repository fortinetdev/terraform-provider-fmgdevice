// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Virtual Access Points (VAPs).

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerVap() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerVapCreate,
		Read:   resourceWirelessControllerVapRead,
		Update: resourceWirelessControllerVapUpdate,
		Delete: resourceWirelessControllerVapDelete,

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
			"n80211k": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"n80211v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"_centmgmt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"_dhcp_svr_id": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"_intf_allowaccess": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"_intf_device_access_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"_intf_device_identification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"_intf_device_netscan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"_intf_dhcp_relay_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"_intf_dhcp_relay_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"_intf_dhcp_relay_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"_intf_dhcp6_relay_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"_intf_dhcp6_relay_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"_intf_dhcp6_relay_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"_intf_ip": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"_intf_ip_managed_by_fortiipam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"_intf_ip6_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"_intf_ip6_allowaccess": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"_intf_listen_forticlient_connection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"_intf_managed_subnetwork_size": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"_intf_role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"_is_factory_setting": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"acct_interim_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"access_control_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"additional_akms": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"address_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"address_group_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"akm24_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"alias": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"antivirus_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"application_detection_engine": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"application_dscp_marking": &schema.Schema{
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
			"application_report_intv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"atf_weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"auth_portal_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"beacon_advertising": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"beacon_protection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"broadcast_ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"broadcast_suppression": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"bss_color_partial": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bstm_disassociation_imminent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bstm_load_balancing_disassoc_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"bstm_rssi_disassoc_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"called_station_id_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"captive_portal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"captive_portal_ac_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"captive_portal_macauth_radius_secret": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"captive_portal_macauth_radius_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"captive_portal_radius_secret": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"captive_portal_radius_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"captive_portal_session_timeout_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"captive_portal_auth_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"captive_portal_fw_accounting": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_address_enforcement": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_lease_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dhcp_option43_insertion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_option82_circuit_id_insertion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_option82_insertion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_option82_remote_id_insertion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain_name_stripping": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_mapping": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"n80211k": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"n80211v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"_centmgmt": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"_dhcp_svr_id": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"_intf_allowaccess": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"_intf_device_access_list": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"_intf_device_identification": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"_intf_device_netscan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"_intf_dhcp_relay_ip": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"_intf_dhcp_relay_service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"_intf_dhcp_relay_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"_intf_dhcp6_relay_ip": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"_intf_dhcp6_relay_service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"_intf_dhcp6_relay_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"_intf_ip": &schema.Schema{
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"_intf_ip_managed_by_fortiipam": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"_intf_ip6_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"_intf_ip6_allowaccess": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"_intf_listen_forticlient_connection": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"_intf_managed_subnetwork_size": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"_intf_role": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"_is_factory_setting": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"_scope": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"vdom": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"access_control_list": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"acct_interim_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"additional_akms": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"address_group": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"address_group_policy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"akm24_only": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"alias": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"antivirus_profile": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"application_detection_engine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"application_dscp_marking": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"application_list": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"application_report_intv": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"atf_weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"auth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"auth_cert": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"auth_portal_addr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"beacon_advertising": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"beacon_protection": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"broadcast_ssid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"broadcast_suppression": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"bss_color_partial": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"bstm_disassociation_imminent": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"bstm_load_balancing_disassoc_timer": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"bstm_rssi_disassoc_timer": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"called_station_id_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"captive_portal": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"captive_portal_ac_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"captive_portal_auth_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"captive_portal_fw_accounting": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"captive_portal_macauth_radius_secret": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"captive_portal_macauth_radius_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"captive_portal_radius_secret": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"captive_portal_radius_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"captive_portal_session_timeout_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"client_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"dhcp_address_enforcement": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dhcp_lease_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"dhcp_option43_insertion": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dhcp_option82_circuit_id_insertion": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dhcp_option82_insertion": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dhcp_option82_remote_id_insertion": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"domain_name_stripping": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dynamic_vlan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"eap_reauth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"eap_reauth_intv": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"eapol_key_retries": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"encrypt": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"external_fast_roaming": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"external_logout": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"external_pre_auth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"external_web": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"external_web_format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fast_bss_transition": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fast_roaming": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ft_mobility_domain": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ft_over_ds": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ft_r0_key_lifetime": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"gas_comeback_delay": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"gas_fragmentation_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"gtk_rekey": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"gtk_rekey_intv": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"high_efficiency": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"hotspot20_profile": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"igmp_snooping": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"intra_vap_privacy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"ips_sensor": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"ipv6_rules": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"key": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"keyindex": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"l3_roaming": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"l3_roaming_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ldpc": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"local_authentication": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"local_bridging": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"local_lan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"local_lan_partition": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"local_standalone": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"local_standalone_dns": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"local_standalone_dns_ip": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"local_standalone_nat": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"local_switching": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mac_auth_bypass": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mac_called_station_delimiter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mac_calling_station_delimiter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mac_case": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mac_filter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mac_filter_policy_other": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mac_password_delimiter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mac_username_delimiter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"max_clients": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"max_clients_ap": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mbo": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mbo_cell_data_conn_pref": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"me_disable_thresh": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mesh_backhaul": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mpsk": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mpsk_concurrent_clients": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mpsk_profile": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"mu_mimo": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"multicast_enhance": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"multicast_rate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"nac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"nac_profile": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"nas_filter_rule": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"neighbor_report_dual_band": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"okc": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"osen": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"owe_groups": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"owe_transition": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"owe_transition_ssid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"passphrase": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"pmf": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pmf_assoc_comeback_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"pmf_sa_query_retry_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"port_macauth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_macauth_reauth_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"port_macauth_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"portal_message_override_group": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"portal_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pre_auth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"primary_wag_profile": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"probe_resp_suppression": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"probe_resp_threshold": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ptk_rekey": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ptk_rekey_intv": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"qos_profile": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"radio_2g_threshold": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"radio_5g_threshold": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"radio_sensitivity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"radius_mac_auth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"radius_mac_auth_block_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"radius_mac_auth_server": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"radius_mac_auth_usergroups": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"radius_mac_mpsk_auth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"radius_mac_mpsk_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"radius_server": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"rates_11a": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"rates_11ac_mcs_map": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rates_11ac_ss12": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"rates_11ac_ss34": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"rates_11ax_mcs_map": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rates_11ax_ss12": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"rates_11ax_ss34": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"rates_11be_mcs_map": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rates_11be_mcs_map_160": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rates_11be_mcs_map_320": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rates_11bg": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"rates_11n_ss12": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"rates_11n_ss34": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"roaming_acct_interim_update": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sae_groups": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"sae_h2e_only": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sae_hnp_only": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sae_password": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"sae_pk": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sae_private_key": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"scan_botnet_connections": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"schedule": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"secondary_wag_profile": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"security": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"security_exempt_list": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"security_obsolete_option": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"security_redirect_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"selected_usergroups": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"split_tunneling": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sticky_client_remove": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sticky_client_threshold_2g": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sticky_client_threshold_5g": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sticky_client_threshold_6g": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"target_wake_time": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tkip_counter_measure": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tunnel_echo_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"tunnel_fallback_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"usergroup": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"utm_log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"utm_profile": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"utm_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vdom": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"vlan_auto": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vlan_pooling": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vlanid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"voice_enterprise": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"webfilter_profile": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"eap_reauth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eap_reauth_intv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"eapol_key_retries": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encrypt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_fast_roaming": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_logout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_pre_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_web": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_web_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fast_bss_transition": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fast_roaming": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ft_mobility_domain": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ft_over_ds": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ft_r0_key_lifetime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"gas_comeback_delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gas_fragmentation_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gtk_rekey": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gtk_rekey_intv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"high_efficiency": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hotspot20_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"igmp_snooping": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"intra_vap_privacy": &schema.Schema{
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
			"ips_sensor": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ipv6_rules": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"keyindex": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"l3_roaming": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"l3_roaming_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldpc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"local_bridging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_lan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_lan_partition": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"local_standalone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_standalone_dns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_standalone_dns_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"local_standalone_nat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mac_auth_bypass": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_called_station_delimiter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_calling_station_delimiter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_case": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_filter_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mac_filter_policy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"mac_filter_policy_other": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_password_delimiter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_username_delimiter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_clients": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_clients_ap": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mbo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mbo_cell_data_conn_pref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"me_disable_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mesh_backhaul": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mpsk": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mpsk_concurrent_clients": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mpsk_key": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"concurrent_clients": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"key_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mpsk_schedules": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"passphrase": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
					},
				},
			},
			"mpsk_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"mu_mimo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multicast_enhance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multicast_rate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nac_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"nas_filter_rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"neighbor_report_dual_band": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"okc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"osen": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"owe_groups": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"owe_transition": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"owe_transition_ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"passphrase": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"pmf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pmf_assoc_comeback_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pmf_sa_query_retry_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"port_macauth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port_macauth_reauth_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"port_macauth_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"portal_message_override_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"portal_message_overrides": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_disclaimer_page": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"auth_login_failed_page": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"auth_login_page": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"auth_reject_page": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"portal_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pre_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"primary_wag_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"probe_resp_suppression": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"probe_resp_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ptk_rekey": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ptk_rekey_intv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"qos_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"quarantine": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radio_2g_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radio_5g_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radio_sensitivity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_mac_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_mac_auth_block_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"radius_mac_auth_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"radius_mac_auth_usergroups": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"radius_mac_mpsk_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_mac_mpsk_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"radius_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"rates_11a": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"rates_11ac_mcs_map": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rates_11ac_ss12": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"rates_11ac_ss34": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"rates_11ax_mcs_map": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rates_11ax_ss12": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"rates_11ax_ss34": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"rates_11be_mcs_map": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rates_11be_mcs_map_160": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rates_11be_mcs_map_320": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rates_11bg": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"rates_11n_ss12": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"rates_11n_ss34": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"roaming_acct_interim_update": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sae_groups": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sae_h2e_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sae_hnp_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sae_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"sae_pk": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sae_private_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"scan_botnet_connections": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"schedule": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"secondary_wag_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"security": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_exempt_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"security_obsolete_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_redirect_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"selected_usergroups": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"split_tunneling": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sticky_client_remove": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sticky_client_threshold_2g": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sticky_client_threshold_5g": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sticky_client_threshold_6g": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"target_wake_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tkip_counter_measure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tunnel_echo_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tunnel_fallback_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"usergroup": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"utm_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"utm_profile": &schema.Schema{
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
			"vlan_auto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan_name": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vlan_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"vlan_pool": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"wtp_group": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"vlan_pooling": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlanid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"voice_enterprise": &schema.Schema{
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceWirelessControllerVapCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWirelessControllerVap(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerVap resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerVap(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerVap resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerVapRead(d, m)
}

func resourceWirelessControllerVapUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWirelessControllerVap(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerVap resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerVap(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerVap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerVapRead(d, m)
}

func resourceWirelessControllerVapDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWirelessControllerVap(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerVap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerVapRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerVap(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerVap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerVap(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerVap resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerVap80211K(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVap80211V(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapCentmgmt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDhcpSvrId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapIntfAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapIntfDeviceAccessList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapIntfDeviceIdentification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapIntfDeviceNetscan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapIntfDhcpRelayIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapIntfDhcpRelayService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapIntfDhcpRelayType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapIntfDhcp6RelayIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapIntfDhcp6RelayService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapIntfDhcp6RelayType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapIntfIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapIntfIpManagedByFortiipam(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapIntfIp6Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapIntfIp6Allowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapIntfListenForticlientConnection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapIntfManagedSubnetworkSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapIntfRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapIsFactorySetting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapAcctInterimInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapAccessControlList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapAdditionalAkms(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapAddressGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapAddressGroupPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapAkm24Only(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapAlias(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapAntivirusProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapApplicationDetectionEngine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapApplicationDscpMarking(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapApplicationList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapApplicationReportIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapAtfWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapAuthCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapAuthPortalAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapBeaconAdvertising(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapBeaconProtection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapBroadcastSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapBroadcastSuppression(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapBssColorPartial(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapBstmDisassociationImminent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapBstmLoadBalancingDisassocTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapBstmRssiDisassocTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapCalledStationIdType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapCaptivePortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapCaptivePortalAcName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapCaptivePortalMacauthRadiusServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapCaptivePortalRadiusServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapCaptivePortalSessionTimeoutInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapCaptivePortalAuthTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapCaptivePortalFwAccounting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDhcpAddressEnforcement(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDhcpLeaseTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDhcpOption43Insertion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDhcpOption82CircuitIdInsertion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDhcpOption82Insertion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDhcpOption82RemoteIdInsertion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDomainNameStripping(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "n80211k"
		if _, ok := i["80211k"]; ok {
			v := flattenWirelessControllerVapDynamicMapping80211K(i["80211k"], d, pre_append)
			tmp["n80211k"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-80211K")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "n80211v"
		if _, ok := i["80211v"]; ok {
			v := flattenWirelessControllerVapDynamicMapping80211V(i["80211v"], d, pre_append)
			tmp["n80211v"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-80211V")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_centmgmt"
		if _, ok := i["_centmgmt"]; ok {
			v := flattenWirelessControllerVapDynamicMappingCentmgmt(i["_centmgmt"], d, pre_append)
			tmp["_centmgmt"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Centmgmt")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_dhcp_svr_id"
		if _, ok := i["_dhcp_svr_id"]; ok {
			v := flattenWirelessControllerVapDynamicMappingDhcpSvrId(i["_dhcp_svr_id"], d, pre_append)
			tmp["_dhcp_svr_id"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-DhcpSvrId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_allowaccess"
		if _, ok := i["_intf_allowaccess"]; ok {
			v := flattenWirelessControllerVapDynamicMappingIntfAllowaccess(i["_intf_allowaccess"], d, pre_append)
			tmp["_intf_allowaccess"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-IntfAllowaccess")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_device_access_list"
		if _, ok := i["_intf_device-access-list"]; ok {
			v := flattenWirelessControllerVapDynamicMappingIntfDeviceAccessList(i["_intf_device-access-list"], d, pre_append)
			tmp["_intf_device_access_list"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-IntfDeviceAccessList")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_device_identification"
		if _, ok := i["_intf_device-identification"]; ok {
			v := flattenWirelessControllerVapDynamicMappingIntfDeviceIdentification(i["_intf_device-identification"], d, pre_append)
			tmp["_intf_device_identification"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-IntfDeviceIdentification")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_device_netscan"
		if _, ok := i["_intf_device-netscan"]; ok {
			v := flattenWirelessControllerVapDynamicMappingIntfDeviceNetscan(i["_intf_device-netscan"], d, pre_append)
			tmp["_intf_device_netscan"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-IntfDeviceNetscan")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_dhcp_relay_ip"
		if _, ok := i["_intf_dhcp-relay-ip"]; ok {
			v := flattenWirelessControllerVapDynamicMappingIntfDhcpRelayIp(i["_intf_dhcp-relay-ip"], d, pre_append)
			tmp["_intf_dhcp_relay_ip"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-IntfDhcpRelayIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_dhcp_relay_service"
		if _, ok := i["_intf_dhcp-relay-service"]; ok {
			v := flattenWirelessControllerVapDynamicMappingIntfDhcpRelayService(i["_intf_dhcp-relay-service"], d, pre_append)
			tmp["_intf_dhcp_relay_service"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-IntfDhcpRelayService")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_dhcp_relay_type"
		if _, ok := i["_intf_dhcp-relay-type"]; ok {
			v := flattenWirelessControllerVapDynamicMappingIntfDhcpRelayType(i["_intf_dhcp-relay-type"], d, pre_append)
			tmp["_intf_dhcp_relay_type"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-IntfDhcpRelayType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_dhcp6_relay_ip"
		if _, ok := i["_intf_dhcp6-relay-ip"]; ok {
			v := flattenWirelessControllerVapDynamicMappingIntfDhcp6RelayIp(i["_intf_dhcp6-relay-ip"], d, pre_append)
			tmp["_intf_dhcp6_relay_ip"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-IntfDhcp6RelayIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_dhcp6_relay_service"
		if _, ok := i["_intf_dhcp6-relay-service"]; ok {
			v := flattenWirelessControllerVapDynamicMappingIntfDhcp6RelayService(i["_intf_dhcp6-relay-service"], d, pre_append)
			tmp["_intf_dhcp6_relay_service"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-IntfDhcp6RelayService")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_dhcp6_relay_type"
		if _, ok := i["_intf_dhcp6-relay-type"]; ok {
			v := flattenWirelessControllerVapDynamicMappingIntfDhcp6RelayType(i["_intf_dhcp6-relay-type"], d, pre_append)
			tmp["_intf_dhcp6_relay_type"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-IntfDhcp6RelayType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_ip"
		if _, ok := i["_intf_ip"]; ok {
			v := flattenWirelessControllerVapDynamicMappingIntfIp(i["_intf_ip"], d, pre_append)
			tmp["_intf_ip"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-IntfIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_ip_managed_by_fortiipam"
		if _, ok := i["_intf_ip-managed-by-fortiipam"]; ok {
			v := flattenWirelessControllerVapDynamicMappingIntfIpManagedByFortiipam(i["_intf_ip-managed-by-fortiipam"], d, pre_append)
			tmp["_intf_ip_managed_by_fortiipam"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-IntfIpManagedByFortiipam")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_ip6_address"
		if _, ok := i["_intf_ip6-address"]; ok {
			v := flattenWirelessControllerVapDynamicMappingIntfIp6Address(i["_intf_ip6-address"], d, pre_append)
			tmp["_intf_ip6_address"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-IntfIp6Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_ip6_allowaccess"
		if _, ok := i["_intf_ip6-allowaccess"]; ok {
			v := flattenWirelessControllerVapDynamicMappingIntfIp6Allowaccess(i["_intf_ip6-allowaccess"], d, pre_append)
			tmp["_intf_ip6_allowaccess"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-IntfIp6Allowaccess")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_listen_forticlient_connection"
		if _, ok := i["_intf_listen-forticlient-connection"]; ok {
			v := flattenWirelessControllerVapDynamicMappingIntfListenForticlientConnection(i["_intf_listen-forticlient-connection"], d, pre_append)
			tmp["_intf_listen_forticlient_connection"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-IntfListenForticlientConnection")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_managed_subnetwork_size"
		if _, ok := i["_intf_managed-subnetwork-size"]; ok {
			v := flattenWirelessControllerVapDynamicMappingIntfManagedSubnetworkSize(i["_intf_managed-subnetwork-size"], d, pre_append)
			tmp["_intf_managed_subnetwork_size"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-IntfManagedSubnetworkSize")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_role"
		if _, ok := i["_intf_role"]; ok {
			v := flattenWirelessControllerVapDynamicMappingIntfRole(i["_intf_role"], d, pre_append)
			tmp["_intf_role"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-IntfRole")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_is_factory_setting"
		if _, ok := i["_is_factory_setting"]; ok {
			v := flattenWirelessControllerVapDynamicMappingIsFactorySetting(i["_is_factory_setting"], d, pre_append)
			tmp["_is_factory_setting"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-IsFactorySetting")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := i["_scope"]; ok {
			v := flattenWirelessControllerVapDynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_control_list"
		if _, ok := i["access-control-list"]; ok {
			v := flattenWirelessControllerVapDynamicMappingAccessControlList(i["access-control-list"], d, pre_append)
			tmp["access_control_list"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-AccessControlList")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "acct_interim_interval"
		if _, ok := i["acct-interim-interval"]; ok {
			v := flattenWirelessControllerVapDynamicMappingAcctInterimInterval(i["acct-interim-interval"], d, pre_append)
			tmp["acct_interim_interval"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-AcctInterimInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_akms"
		if _, ok := i["additional-akms"]; ok {
			v := flattenWirelessControllerVapDynamicMappingAdditionalAkms(i["additional-akms"], d, pre_append)
			tmp["additional_akms"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-AdditionalAkms")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address_group"
		if _, ok := i["address-group"]; ok {
			v := flattenWirelessControllerVapDynamicMappingAddressGroup(i["address-group"], d, pre_append)
			tmp["address_group"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-AddressGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address_group_policy"
		if _, ok := i["address-group-policy"]; ok {
			v := flattenWirelessControllerVapDynamicMappingAddressGroupPolicy(i["address-group-policy"], d, pre_append)
			tmp["address_group_policy"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-AddressGroupPolicy")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "akm24_only"
		if _, ok := i["akm24-only"]; ok {
			v := flattenWirelessControllerVapDynamicMappingAkm24Only(i["akm24-only"], d, pre_append)
			tmp["akm24_only"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Akm24Only")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "alias"
		if _, ok := i["alias"]; ok {
			v := flattenWirelessControllerVapDynamicMappingAlias(i["alias"], d, pre_append)
			tmp["alias"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Alias")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "antivirus_profile"
		if _, ok := i["antivirus-profile"]; ok {
			v := flattenWirelessControllerVapDynamicMappingAntivirusProfile(i["antivirus-profile"], d, pre_append)
			tmp["antivirus_profile"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-AntivirusProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application_detection_engine"
		if _, ok := i["application-detection-engine"]; ok {
			v := flattenWirelessControllerVapDynamicMappingApplicationDetectionEngine(i["application-detection-engine"], d, pre_append)
			tmp["application_detection_engine"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-ApplicationDetectionEngine")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application_dscp_marking"
		if _, ok := i["application-dscp-marking"]; ok {
			v := flattenWirelessControllerVapDynamicMappingApplicationDscpMarking(i["application-dscp-marking"], d, pre_append)
			tmp["application_dscp_marking"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-ApplicationDscpMarking")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application_list"
		if _, ok := i["application-list"]; ok {
			v := flattenWirelessControllerVapDynamicMappingApplicationList(i["application-list"], d, pre_append)
			tmp["application_list"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-ApplicationList")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application_report_intv"
		if _, ok := i["application-report-intv"]; ok {
			v := flattenWirelessControllerVapDynamicMappingApplicationReportIntv(i["application-report-intv"], d, pre_append)
			tmp["application_report_intv"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-ApplicationReportIntv")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "atf_weight"
		if _, ok := i["atf-weight"]; ok {
			v := flattenWirelessControllerVapDynamicMappingAtfWeight(i["atf-weight"], d, pre_append)
			tmp["atf_weight"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-AtfWeight")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth"
		if _, ok := i["auth"]; ok {
			v := flattenWirelessControllerVapDynamicMappingAuth(i["auth"], d, pre_append)
			tmp["auth"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Auth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_cert"
		if _, ok := i["auth-cert"]; ok {
			v := flattenWirelessControllerVapDynamicMappingAuthCert(i["auth-cert"], d, pre_append)
			tmp["auth_cert"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-AuthCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_portal_addr"
		if _, ok := i["auth-portal-addr"]; ok {
			v := flattenWirelessControllerVapDynamicMappingAuthPortalAddr(i["auth-portal-addr"], d, pre_append)
			tmp["auth_portal_addr"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-AuthPortalAddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "beacon_advertising"
		if _, ok := i["beacon-advertising"]; ok {
			v := flattenWirelessControllerVapDynamicMappingBeaconAdvertising(i["beacon-advertising"], d, pre_append)
			tmp["beacon_advertising"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-BeaconAdvertising")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "beacon_protection"
		if _, ok := i["beacon-protection"]; ok {
			v := flattenWirelessControllerVapDynamicMappingBeaconProtection(i["beacon-protection"], d, pre_append)
			tmp["beacon_protection"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-BeaconProtection")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "broadcast_ssid"
		if _, ok := i["broadcast-ssid"]; ok {
			v := flattenWirelessControllerVapDynamicMappingBroadcastSsid(i["broadcast-ssid"], d, pre_append)
			tmp["broadcast_ssid"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-BroadcastSsid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "broadcast_suppression"
		if _, ok := i["broadcast-suppression"]; ok {
			v := flattenWirelessControllerVapDynamicMappingBroadcastSuppression(i["broadcast-suppression"], d, pre_append)
			tmp["broadcast_suppression"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-BroadcastSuppression")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bss_color_partial"
		if _, ok := i["bss-color-partial"]; ok {
			v := flattenWirelessControllerVapDynamicMappingBssColorPartial(i["bss-color-partial"], d, pre_append)
			tmp["bss_color_partial"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-BssColorPartial")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bstm_disassociation_imminent"
		if _, ok := i["bstm-disassociation-imminent"]; ok {
			v := flattenWirelessControllerVapDynamicMappingBstmDisassociationImminent(i["bstm-disassociation-imminent"], d, pre_append)
			tmp["bstm_disassociation_imminent"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-BstmDisassociationImminent")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bstm_load_balancing_disassoc_timer"
		if _, ok := i["bstm-load-balancing-disassoc-timer"]; ok {
			v := flattenWirelessControllerVapDynamicMappingBstmLoadBalancingDisassocTimer(i["bstm-load-balancing-disassoc-timer"], d, pre_append)
			tmp["bstm_load_balancing_disassoc_timer"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-BstmLoadBalancingDisassocTimer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bstm_rssi_disassoc_timer"
		if _, ok := i["bstm-rssi-disassoc-timer"]; ok {
			v := flattenWirelessControllerVapDynamicMappingBstmRssiDisassocTimer(i["bstm-rssi-disassoc-timer"], d, pre_append)
			tmp["bstm_rssi_disassoc_timer"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-BstmRssiDisassocTimer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "called_station_id_type"
		if _, ok := i["called-station-id-type"]; ok {
			v := flattenWirelessControllerVapDynamicMappingCalledStationIdType(i["called-station-id-type"], d, pre_append)
			tmp["called_station_id_type"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-CalledStationIdType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal"
		if _, ok := i["captive-portal"]; ok {
			v := flattenWirelessControllerVapDynamicMappingCaptivePortal(i["captive-portal"], d, pre_append)
			tmp["captive_portal"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-CaptivePortal")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal_ac_name"
		if _, ok := i["captive-portal-ac-name"]; ok {
			v := flattenWirelessControllerVapDynamicMappingCaptivePortalAcName(i["captive-portal-ac-name"], d, pre_append)
			tmp["captive_portal_ac_name"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-CaptivePortalAcName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal_auth_timeout"
		if _, ok := i["captive-portal-auth-timeout"]; ok {
			v := flattenWirelessControllerVapDynamicMappingCaptivePortalAuthTimeout(i["captive-portal-auth-timeout"], d, pre_append)
			tmp["captive_portal_auth_timeout"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-CaptivePortalAuthTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal_fw_accounting"
		if _, ok := i["captive-portal-fw-accounting"]; ok {
			v := flattenWirelessControllerVapDynamicMappingCaptivePortalFwAccounting(i["captive-portal-fw-accounting"], d, pre_append)
			tmp["captive_portal_fw_accounting"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-CaptivePortalFwAccounting")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal_macauth_radius_server"
		if _, ok := i["captive-portal-macauth-radius-server"]; ok {
			v := flattenWirelessControllerVapDynamicMappingCaptivePortalMacauthRadiusServer(i["captive-portal-macauth-radius-server"], d, pre_append)
			tmp["captive_portal_macauth_radius_server"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-CaptivePortalMacauthRadiusServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal_radius_server"
		if _, ok := i["captive-portal-radius-server"]; ok {
			v := flattenWirelessControllerVapDynamicMappingCaptivePortalRadiusServer(i["captive-portal-radius-server"], d, pre_append)
			tmp["captive_portal_radius_server"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-CaptivePortalRadiusServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal_session_timeout_interval"
		if _, ok := i["captive-portal-session-timeout-interval"]; ok {
			v := flattenWirelessControllerVapDynamicMappingCaptivePortalSessionTimeoutInterval(i["captive-portal-session-timeout-interval"], d, pre_append)
			tmp["captive_portal_session_timeout_interval"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-CaptivePortalSessionTimeoutInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_count"
		if _, ok := i["client-count"]; ok {
			v := flattenWirelessControllerVapDynamicMappingClientCount(i["client-count"], d, pre_append)
			tmp["client_count"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-ClientCount")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_address_enforcement"
		if _, ok := i["dhcp-address-enforcement"]; ok {
			v := flattenWirelessControllerVapDynamicMappingDhcpAddressEnforcement(i["dhcp-address-enforcement"], d, pre_append)
			tmp["dhcp_address_enforcement"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-DhcpAddressEnforcement")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_lease_time"
		if _, ok := i["dhcp-lease-time"]; ok {
			v := flattenWirelessControllerVapDynamicMappingDhcpLeaseTime(i["dhcp-lease-time"], d, pre_append)
			tmp["dhcp_lease_time"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-DhcpLeaseTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_option43_insertion"
		if _, ok := i["dhcp-option43-insertion"]; ok {
			v := flattenWirelessControllerVapDynamicMappingDhcpOption43Insertion(i["dhcp-option43-insertion"], d, pre_append)
			tmp["dhcp_option43_insertion"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-DhcpOption43Insertion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_option82_circuit_id_insertion"
		if _, ok := i["dhcp-option82-circuit-id-insertion"]; ok {
			v := flattenWirelessControllerVapDynamicMappingDhcpOption82CircuitIdInsertion(i["dhcp-option82-circuit-id-insertion"], d, pre_append)
			tmp["dhcp_option82_circuit_id_insertion"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-DhcpOption82CircuitIdInsertion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_option82_insertion"
		if _, ok := i["dhcp-option82-insertion"]; ok {
			v := flattenWirelessControllerVapDynamicMappingDhcpOption82Insertion(i["dhcp-option82-insertion"], d, pre_append)
			tmp["dhcp_option82_insertion"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-DhcpOption82Insertion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_option82_remote_id_insertion"
		if _, ok := i["dhcp-option82-remote-id-insertion"]; ok {
			v := flattenWirelessControllerVapDynamicMappingDhcpOption82RemoteIdInsertion(i["dhcp-option82-remote-id-insertion"], d, pre_append)
			tmp["dhcp_option82_remote_id_insertion"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-DhcpOption82RemoteIdInsertion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain_name_stripping"
		if _, ok := i["domain-name-stripping"]; ok {
			v := flattenWirelessControllerVapDynamicMappingDomainNameStripping(i["domain-name-stripping"], d, pre_append)
			tmp["domain_name_stripping"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-DomainNameStripping")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dynamic_vlan"
		if _, ok := i["dynamic-vlan"]; ok {
			v := flattenWirelessControllerVapDynamicMappingDynamicVlan(i["dynamic-vlan"], d, pre_append)
			tmp["dynamic_vlan"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-DynamicVlan")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "eap_reauth"
		if _, ok := i["eap-reauth"]; ok {
			v := flattenWirelessControllerVapDynamicMappingEapReauth(i["eap-reauth"], d, pre_append)
			tmp["eap_reauth"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-EapReauth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "eap_reauth_intv"
		if _, ok := i["eap-reauth-intv"]; ok {
			v := flattenWirelessControllerVapDynamicMappingEapReauthIntv(i["eap-reauth-intv"], d, pre_append)
			tmp["eap_reauth_intv"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-EapReauthIntv")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "eapol_key_retries"
		if _, ok := i["eapol-key-retries"]; ok {
			v := flattenWirelessControllerVapDynamicMappingEapolKeyRetries(i["eapol-key-retries"], d, pre_append)
			tmp["eapol_key_retries"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-EapolKeyRetries")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "encrypt"
		if _, ok := i["encrypt"]; ok {
			v := flattenWirelessControllerVapDynamicMappingEncrypt(i["encrypt"], d, pre_append)
			tmp["encrypt"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Encrypt")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_fast_roaming"
		if _, ok := i["external-fast-roaming"]; ok {
			v := flattenWirelessControllerVapDynamicMappingExternalFastRoaming(i["external-fast-roaming"], d, pre_append)
			tmp["external_fast_roaming"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-ExternalFastRoaming")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_logout"
		if _, ok := i["external-logout"]; ok {
			v := flattenWirelessControllerVapDynamicMappingExternalLogout(i["external-logout"], d, pre_append)
			tmp["external_logout"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-ExternalLogout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_pre_auth"
		if _, ok := i["external-pre-auth"]; ok {
			v := flattenWirelessControllerVapDynamicMappingExternalPreAuth(i["external-pre-auth"], d, pre_append)
			tmp["external_pre_auth"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-ExternalPreAuth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_web"
		if _, ok := i["external-web"]; ok {
			v := flattenWirelessControllerVapDynamicMappingExternalWeb(i["external-web"], d, pre_append)
			tmp["external_web"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-ExternalWeb")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_web_format"
		if _, ok := i["external-web-format"]; ok {
			v := flattenWirelessControllerVapDynamicMappingExternalWebFormat(i["external-web-format"], d, pre_append)
			tmp["external_web_format"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-ExternalWebFormat")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fast_bss_transition"
		if _, ok := i["fast-bss-transition"]; ok {
			v := flattenWirelessControllerVapDynamicMappingFastBssTransition(i["fast-bss-transition"], d, pre_append)
			tmp["fast_bss_transition"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-FastBssTransition")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fast_roaming"
		if _, ok := i["fast-roaming"]; ok {
			v := flattenWirelessControllerVapDynamicMappingFastRoaming(i["fast-roaming"], d, pre_append)
			tmp["fast_roaming"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-FastRoaming")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ft_mobility_domain"
		if _, ok := i["ft-mobility-domain"]; ok {
			v := flattenWirelessControllerVapDynamicMappingFtMobilityDomain(i["ft-mobility-domain"], d, pre_append)
			tmp["ft_mobility_domain"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-FtMobilityDomain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ft_over_ds"
		if _, ok := i["ft-over-ds"]; ok {
			v := flattenWirelessControllerVapDynamicMappingFtOverDs(i["ft-over-ds"], d, pre_append)
			tmp["ft_over_ds"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-FtOverDs")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ft_r0_key_lifetime"
		if _, ok := i["ft-r0-key-lifetime"]; ok {
			v := flattenWirelessControllerVapDynamicMappingFtR0KeyLifetime(i["ft-r0-key-lifetime"], d, pre_append)
			tmp["ft_r0_key_lifetime"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-FtR0KeyLifetime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gas_comeback_delay"
		if _, ok := i["gas-comeback-delay"]; ok {
			v := flattenWirelessControllerVapDynamicMappingGasComebackDelay(i["gas-comeback-delay"], d, pre_append)
			tmp["gas_comeback_delay"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-GasComebackDelay")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gas_fragmentation_limit"
		if _, ok := i["gas-fragmentation-limit"]; ok {
			v := flattenWirelessControllerVapDynamicMappingGasFragmentationLimit(i["gas-fragmentation-limit"], d, pre_append)
			tmp["gas_fragmentation_limit"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-GasFragmentationLimit")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gtk_rekey"
		if _, ok := i["gtk-rekey"]; ok {
			v := flattenWirelessControllerVapDynamicMappingGtkRekey(i["gtk-rekey"], d, pre_append)
			tmp["gtk_rekey"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-GtkRekey")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gtk_rekey_intv"
		if _, ok := i["gtk-rekey-intv"]; ok {
			v := flattenWirelessControllerVapDynamicMappingGtkRekeyIntv(i["gtk-rekey-intv"], d, pre_append)
			tmp["gtk_rekey_intv"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-GtkRekeyIntv")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "high_efficiency"
		if _, ok := i["high-efficiency"]; ok {
			v := flattenWirelessControllerVapDynamicMappingHighEfficiency(i["high-efficiency"], d, pre_append)
			tmp["high_efficiency"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-HighEfficiency")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hotspot20_profile"
		if _, ok := i["hotspot20-profile"]; ok {
			v := flattenWirelessControllerVapDynamicMappingHotspot20Profile(i["hotspot20-profile"], d, pre_append)
			tmp["hotspot20_profile"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Hotspot20Profile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmp_snooping"
		if _, ok := i["igmp-snooping"]; ok {
			v := flattenWirelessControllerVapDynamicMappingIgmpSnooping(i["igmp-snooping"], d, pre_append)
			tmp["igmp_snooping"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-IgmpSnooping")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "intra_vap_privacy"
		if _, ok := i["intra-vap-privacy"]; ok {
			v := flattenWirelessControllerVapDynamicMappingIntraVapPrivacy(i["intra-vap-privacy"], d, pre_append)
			tmp["intra_vap_privacy"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-IntraVapPrivacy")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenWirelessControllerVapDynamicMappingIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ips_sensor"
		if _, ok := i["ips-sensor"]; ok {
			v := flattenWirelessControllerVapDynamicMappingIpsSensor(i["ips-sensor"], d, pre_append)
			tmp["ips_sensor"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-IpsSensor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_rules"
		if _, ok := i["ipv6-rules"]; ok {
			v := flattenWirelessControllerVapDynamicMappingIpv6Rules(i["ipv6-rules"], d, pre_append)
			tmp["ipv6_rules"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Ipv6Rules")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keyindex"
		if _, ok := i["keyindex"]; ok {
			v := flattenWirelessControllerVapDynamicMappingKeyindex(i["keyindex"], d, pre_append)
			tmp["keyindex"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Keyindex")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "l3_roaming"
		if _, ok := i["l3-roaming"]; ok {
			v := flattenWirelessControllerVapDynamicMappingL3Roaming(i["l3-roaming"], d, pre_append)
			tmp["l3_roaming"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-L3Roaming")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "l3_roaming_mode"
		if _, ok := i["l3-roaming-mode"]; ok {
			v := flattenWirelessControllerVapDynamicMappingL3RoamingMode(i["l3-roaming-mode"], d, pre_append)
			tmp["l3_roaming_mode"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-L3RoamingMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldpc"
		if _, ok := i["ldpc"]; ok {
			v := flattenWirelessControllerVapDynamicMappingLdpc(i["ldpc"], d, pre_append)
			tmp["ldpc"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Ldpc")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_authentication"
		if _, ok := i["local-authentication"]; ok {
			v := flattenWirelessControllerVapDynamicMappingLocalAuthentication(i["local-authentication"], d, pre_append)
			tmp["local_authentication"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-LocalAuthentication")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_bridging"
		if _, ok := i["local-bridging"]; ok {
			v := flattenWirelessControllerVapDynamicMappingLocalBridging(i["local-bridging"], d, pre_append)
			tmp["local_bridging"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-LocalBridging")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_lan"
		if _, ok := i["local-lan"]; ok {
			v := flattenWirelessControllerVapDynamicMappingLocalLan(i["local-lan"], d, pre_append)
			tmp["local_lan"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-LocalLan")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_lan_partition"
		if _, ok := i["local-lan-partition"]; ok {
			v := flattenWirelessControllerVapDynamicMappingLocalLanPartition(i["local-lan-partition"], d, pre_append)
			tmp["local_lan_partition"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-LocalLanPartition")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_standalone"
		if _, ok := i["local-standalone"]; ok {
			v := flattenWirelessControllerVapDynamicMappingLocalStandalone(i["local-standalone"], d, pre_append)
			tmp["local_standalone"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-LocalStandalone")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_standalone_dns"
		if _, ok := i["local-standalone-dns"]; ok {
			v := flattenWirelessControllerVapDynamicMappingLocalStandaloneDns(i["local-standalone-dns"], d, pre_append)
			tmp["local_standalone_dns"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-LocalStandaloneDns")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_standalone_dns_ip"
		if _, ok := i["local-standalone-dns-ip"]; ok {
			v := flattenWirelessControllerVapDynamicMappingLocalStandaloneDnsIp(i["local-standalone-dns-ip"], d, pre_append)
			tmp["local_standalone_dns_ip"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-LocalStandaloneDnsIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_standalone_nat"
		if _, ok := i["local-standalone-nat"]; ok {
			v := flattenWirelessControllerVapDynamicMappingLocalStandaloneNat(i["local-standalone-nat"], d, pre_append)
			tmp["local_standalone_nat"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-LocalStandaloneNat")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_switching"
		if _, ok := i["local-switching"]; ok {
			v := flattenWirelessControllerVapDynamicMappingLocalSwitching(i["local-switching"], d, pre_append)
			tmp["local_switching"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-LocalSwitching")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_auth_bypass"
		if _, ok := i["mac-auth-bypass"]; ok {
			v := flattenWirelessControllerVapDynamicMappingMacAuthBypass(i["mac-auth-bypass"], d, pre_append)
			tmp["mac_auth_bypass"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-MacAuthBypass")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_called_station_delimiter"
		if _, ok := i["mac-called-station-delimiter"]; ok {
			v := flattenWirelessControllerVapDynamicMappingMacCalledStationDelimiter(i["mac-called-station-delimiter"], d, pre_append)
			tmp["mac_called_station_delimiter"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-MacCalledStationDelimiter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_calling_station_delimiter"
		if _, ok := i["mac-calling-station-delimiter"]; ok {
			v := flattenWirelessControllerVapDynamicMappingMacCallingStationDelimiter(i["mac-calling-station-delimiter"], d, pre_append)
			tmp["mac_calling_station_delimiter"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-MacCallingStationDelimiter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_case"
		if _, ok := i["mac-case"]; ok {
			v := flattenWirelessControllerVapDynamicMappingMacCase(i["mac-case"], d, pre_append)
			tmp["mac_case"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-MacCase")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_filter"
		if _, ok := i["mac-filter"]; ok {
			v := flattenWirelessControllerVapDynamicMappingMacFilter(i["mac-filter"], d, pre_append)
			tmp["mac_filter"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-MacFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_filter_policy_other"
		if _, ok := i["mac-filter-policy-other"]; ok {
			v := flattenWirelessControllerVapDynamicMappingMacFilterPolicyOther(i["mac-filter-policy-other"], d, pre_append)
			tmp["mac_filter_policy_other"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-MacFilterPolicyOther")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_password_delimiter"
		if _, ok := i["mac-password-delimiter"]; ok {
			v := flattenWirelessControllerVapDynamicMappingMacPasswordDelimiter(i["mac-password-delimiter"], d, pre_append)
			tmp["mac_password_delimiter"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-MacPasswordDelimiter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_username_delimiter"
		if _, ok := i["mac-username-delimiter"]; ok {
			v := flattenWirelessControllerVapDynamicMappingMacUsernameDelimiter(i["mac-username-delimiter"], d, pre_append)
			tmp["mac_username_delimiter"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-MacUsernameDelimiter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_clients"
		if _, ok := i["max-clients"]; ok {
			v := flattenWirelessControllerVapDynamicMappingMaxClients(i["max-clients"], d, pre_append)
			tmp["max_clients"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-MaxClients")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_clients_ap"
		if _, ok := i["max-clients-ap"]; ok {
			v := flattenWirelessControllerVapDynamicMappingMaxClientsAp(i["max-clients-ap"], d, pre_append)
			tmp["max_clients_ap"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-MaxClientsAp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mbo"
		if _, ok := i["mbo"]; ok {
			v := flattenWirelessControllerVapDynamicMappingMbo(i["mbo"], d, pre_append)
			tmp["mbo"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Mbo")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mbo_cell_data_conn_pref"
		if _, ok := i["mbo-cell-data-conn-pref"]; ok {
			v := flattenWirelessControllerVapDynamicMappingMboCellDataConnPref(i["mbo-cell-data-conn-pref"], d, pre_append)
			tmp["mbo_cell_data_conn_pref"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-MboCellDataConnPref")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "me_disable_thresh"
		if _, ok := i["me-disable-thresh"]; ok {
			v := flattenWirelessControllerVapDynamicMappingMeDisableThresh(i["me-disable-thresh"], d, pre_append)
			tmp["me_disable_thresh"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-MeDisableThresh")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mesh_backhaul"
		if _, ok := i["mesh-backhaul"]; ok {
			v := flattenWirelessControllerVapDynamicMappingMeshBackhaul(i["mesh-backhaul"], d, pre_append)
			tmp["mesh_backhaul"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-MeshBackhaul")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk"
		if _, ok := i["mpsk"]; ok {
			v := flattenWirelessControllerVapDynamicMappingMpsk(i["mpsk"], d, pre_append)
			tmp["mpsk"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Mpsk")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk_concurrent_clients"
		if _, ok := i["mpsk-concurrent-clients"]; ok {
			v := flattenWirelessControllerVapDynamicMappingMpskConcurrentClients(i["mpsk-concurrent-clients"], d, pre_append)
			tmp["mpsk_concurrent_clients"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-MpskConcurrentClients")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk_profile"
		if _, ok := i["mpsk-profile"]; ok {
			v := flattenWirelessControllerVapDynamicMappingMpskProfile(i["mpsk-profile"], d, pre_append)
			tmp["mpsk_profile"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-MpskProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mu_mimo"
		if _, ok := i["mu-mimo"]; ok {
			v := flattenWirelessControllerVapDynamicMappingMuMimo(i["mu-mimo"], d, pre_append)
			tmp["mu_mimo"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-MuMimo")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "multicast_enhance"
		if _, ok := i["multicast-enhance"]; ok {
			v := flattenWirelessControllerVapDynamicMappingMulticastEnhance(i["multicast-enhance"], d, pre_append)
			tmp["multicast_enhance"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-MulticastEnhance")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "multicast_rate"
		if _, ok := i["multicast-rate"]; ok {
			v := flattenWirelessControllerVapDynamicMappingMulticastRate(i["multicast-rate"], d, pre_append)
			tmp["multicast_rate"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-MulticastRate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nac"
		if _, ok := i["nac"]; ok {
			v := flattenWirelessControllerVapDynamicMappingNac(i["nac"], d, pre_append)
			tmp["nac"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Nac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nac_profile"
		if _, ok := i["nac-profile"]; ok {
			v := flattenWirelessControllerVapDynamicMappingNacProfile(i["nac-profile"], d, pre_append)
			tmp["nac_profile"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-NacProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nas_filter_rule"
		if _, ok := i["nas-filter-rule"]; ok {
			v := flattenWirelessControllerVapDynamicMappingNasFilterRule(i["nas-filter-rule"], d, pre_append)
			tmp["nas_filter_rule"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-NasFilterRule")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbor_report_dual_band"
		if _, ok := i["neighbor-report-dual-band"]; ok {
			v := flattenWirelessControllerVapDynamicMappingNeighborReportDualBand(i["neighbor-report-dual-band"], d, pre_append)
			tmp["neighbor_report_dual_band"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-NeighborReportDualBand")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "okc"
		if _, ok := i["okc"]; ok {
			v := flattenWirelessControllerVapDynamicMappingOkc(i["okc"], d, pre_append)
			tmp["okc"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Okc")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "osen"
		if _, ok := i["osen"]; ok {
			v := flattenWirelessControllerVapDynamicMappingOsen(i["osen"], d, pre_append)
			tmp["osen"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Osen")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "owe_groups"
		if _, ok := i["owe-groups"]; ok {
			v := flattenWirelessControllerVapDynamicMappingOweGroups(i["owe-groups"], d, pre_append)
			tmp["owe_groups"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-OweGroups")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "owe_transition"
		if _, ok := i["owe-transition"]; ok {
			v := flattenWirelessControllerVapDynamicMappingOweTransition(i["owe-transition"], d, pre_append)
			tmp["owe_transition"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-OweTransition")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "owe_transition_ssid"
		if _, ok := i["owe-transition-ssid"]; ok {
			v := flattenWirelessControllerVapDynamicMappingOweTransitionSsid(i["owe-transition-ssid"], d, pre_append)
			tmp["owe_transition_ssid"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-OweTransitionSsid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pmf"
		if _, ok := i["pmf"]; ok {
			v := flattenWirelessControllerVapDynamicMappingPmf(i["pmf"], d, pre_append)
			tmp["pmf"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Pmf")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pmf_assoc_comeback_timeout"
		if _, ok := i["pmf-assoc-comeback-timeout"]; ok {
			v := flattenWirelessControllerVapDynamicMappingPmfAssocComebackTimeout(i["pmf-assoc-comeback-timeout"], d, pre_append)
			tmp["pmf_assoc_comeback_timeout"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-PmfAssocComebackTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pmf_sa_query_retry_timeout"
		if _, ok := i["pmf-sa-query-retry-timeout"]; ok {
			v := flattenWirelessControllerVapDynamicMappingPmfSaQueryRetryTimeout(i["pmf-sa-query-retry-timeout"], d, pre_append)
			tmp["pmf_sa_query_retry_timeout"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-PmfSaQueryRetryTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_macauth"
		if _, ok := i["port-macauth"]; ok {
			v := flattenWirelessControllerVapDynamicMappingPortMacauth(i["port-macauth"], d, pre_append)
			tmp["port_macauth"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-PortMacauth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_macauth_reauth_timeout"
		if _, ok := i["port-macauth-reauth-timeout"]; ok {
			v := flattenWirelessControllerVapDynamicMappingPortMacauthReauthTimeout(i["port-macauth-reauth-timeout"], d, pre_append)
			tmp["port_macauth_reauth_timeout"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-PortMacauthReauthTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_macauth_timeout"
		if _, ok := i["port-macauth-timeout"]; ok {
			v := flattenWirelessControllerVapDynamicMappingPortMacauthTimeout(i["port-macauth-timeout"], d, pre_append)
			tmp["port_macauth_timeout"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-PortMacauthTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "portal_message_override_group"
		if _, ok := i["portal-message-override-group"]; ok {
			v := flattenWirelessControllerVapDynamicMappingPortalMessageOverrideGroup(i["portal-message-override-group"], d, pre_append)
			tmp["portal_message_override_group"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-PortalMessageOverrideGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "portal_type"
		if _, ok := i["portal-type"]; ok {
			v := flattenWirelessControllerVapDynamicMappingPortalType(i["portal-type"], d, pre_append)
			tmp["portal_type"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-PortalType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pre_auth"
		if _, ok := i["pre-auth"]; ok {
			v := flattenWirelessControllerVapDynamicMappingPreAuth(i["pre-auth"], d, pre_append)
			tmp["pre_auth"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-PreAuth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "primary_wag_profile"
		if _, ok := i["primary-wag-profile"]; ok {
			v := flattenWirelessControllerVapDynamicMappingPrimaryWagProfile(i["primary-wag-profile"], d, pre_append)
			tmp["primary_wag_profile"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-PrimaryWagProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_resp_suppression"
		if _, ok := i["probe-resp-suppression"]; ok {
			v := flattenWirelessControllerVapDynamicMappingProbeRespSuppression(i["probe-resp-suppression"], d, pre_append)
			tmp["probe_resp_suppression"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-ProbeRespSuppression")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_resp_threshold"
		if _, ok := i["probe-resp-threshold"]; ok {
			v := flattenWirelessControllerVapDynamicMappingProbeRespThreshold(i["probe-resp-threshold"], d, pre_append)
			tmp["probe_resp_threshold"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-ProbeRespThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ptk_rekey"
		if _, ok := i["ptk-rekey"]; ok {
			v := flattenWirelessControllerVapDynamicMappingPtkRekey(i["ptk-rekey"], d, pre_append)
			tmp["ptk_rekey"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-PtkRekey")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ptk_rekey_intv"
		if _, ok := i["ptk-rekey-intv"]; ok {
			v := flattenWirelessControllerVapDynamicMappingPtkRekeyIntv(i["ptk-rekey-intv"], d, pre_append)
			tmp["ptk_rekey_intv"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-PtkRekeyIntv")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "qos_profile"
		if _, ok := i["qos-profile"]; ok {
			v := flattenWirelessControllerVapDynamicMappingQosProfile(i["qos-profile"], d, pre_append)
			tmp["qos_profile"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-QosProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := i["quarantine"]; ok {
			v := flattenWirelessControllerVapDynamicMappingQuarantine(i["quarantine"], d, pre_append)
			tmp["quarantine"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Quarantine")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radio_2g_threshold"
		if _, ok := i["radio-2g-threshold"]; ok {
			v := flattenWirelessControllerVapDynamicMappingRadio2GThreshold(i["radio-2g-threshold"], d, pre_append)
			tmp["radio_2g_threshold"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Radio2GThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radio_5g_threshold"
		if _, ok := i["radio-5g-threshold"]; ok {
			v := flattenWirelessControllerVapDynamicMappingRadio5GThreshold(i["radio-5g-threshold"], d, pre_append)
			tmp["radio_5g_threshold"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Radio5GThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radio_sensitivity"
		if _, ok := i["radio-sensitivity"]; ok {
			v := flattenWirelessControllerVapDynamicMappingRadioSensitivity(i["radio-sensitivity"], d, pre_append)
			tmp["radio_sensitivity"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-RadioSensitivity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_mac_auth"
		if _, ok := i["radius-mac-auth"]; ok {
			v := flattenWirelessControllerVapDynamicMappingRadiusMacAuth(i["radius-mac-auth"], d, pre_append)
			tmp["radius_mac_auth"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-RadiusMacAuth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_mac_auth_block_interval"
		if _, ok := i["radius-mac-auth-block-interval"]; ok {
			v := flattenWirelessControllerVapDynamicMappingRadiusMacAuthBlockInterval(i["radius-mac-auth-block-interval"], d, pre_append)
			tmp["radius_mac_auth_block_interval"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-RadiusMacAuthBlockInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_mac_auth_server"
		if _, ok := i["radius-mac-auth-server"]; ok {
			v := flattenWirelessControllerVapDynamicMappingRadiusMacAuthServer(i["radius-mac-auth-server"], d, pre_append)
			tmp["radius_mac_auth_server"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-RadiusMacAuthServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_mac_auth_usergroups"
		if _, ok := i["radius-mac-auth-usergroups"]; ok {
			v := flattenWirelessControllerVapDynamicMappingRadiusMacAuthUsergroups(i["radius-mac-auth-usergroups"], d, pre_append)
			tmp["radius_mac_auth_usergroups"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-RadiusMacAuthUsergroups")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_mac_mpsk_auth"
		if _, ok := i["radius-mac-mpsk-auth"]; ok {
			v := flattenWirelessControllerVapDynamicMappingRadiusMacMpskAuth(i["radius-mac-mpsk-auth"], d, pre_append)
			tmp["radius_mac_mpsk_auth"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-RadiusMacMpskAuth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_mac_mpsk_timeout"
		if _, ok := i["radius-mac-mpsk-timeout"]; ok {
			v := flattenWirelessControllerVapDynamicMappingRadiusMacMpskTimeout(i["radius-mac-mpsk-timeout"], d, pre_append)
			tmp["radius_mac_mpsk_timeout"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-RadiusMacMpskTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_server"
		if _, ok := i["radius-server"]; ok {
			v := flattenWirelessControllerVapDynamicMappingRadiusServer(i["radius-server"], d, pre_append)
			tmp["radius_server"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-RadiusServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11a"
		if _, ok := i["rates-11a"]; ok {
			v := flattenWirelessControllerVapDynamicMappingRates11A(i["rates-11a"], d, pre_append)
			tmp["rates_11a"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Rates11A")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11ac_mcs_map"
		if _, ok := i["rates-11ac-mcs-map"]; ok {
			v := flattenWirelessControllerVapDynamicMappingRates11AcMcsMap(i["rates-11ac-mcs-map"], d, pre_append)
			tmp["rates_11ac_mcs_map"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Rates11AcMcsMap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11ac_ss12"
		if _, ok := i["rates-11ac-ss12"]; ok {
			v := flattenWirelessControllerVapDynamicMappingRates11AcSs12(i["rates-11ac-ss12"], d, pre_append)
			tmp["rates_11ac_ss12"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Rates11AcSs12")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11ac_ss34"
		if _, ok := i["rates-11ac-ss34"]; ok {
			v := flattenWirelessControllerVapDynamicMappingRates11AcSs34(i["rates-11ac-ss34"], d, pre_append)
			tmp["rates_11ac_ss34"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Rates11AcSs34")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11ax_mcs_map"
		if _, ok := i["rates-11ax-mcs-map"]; ok {
			v := flattenWirelessControllerVapDynamicMappingRates11AxMcsMap(i["rates-11ax-mcs-map"], d, pre_append)
			tmp["rates_11ax_mcs_map"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Rates11AxMcsMap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11ax_ss12"
		if _, ok := i["rates-11ax-ss12"]; ok {
			v := flattenWirelessControllerVapDynamicMappingRates11AxSs12(i["rates-11ax-ss12"], d, pre_append)
			tmp["rates_11ax_ss12"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Rates11AxSs12")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11ax_ss34"
		if _, ok := i["rates-11ax-ss34"]; ok {
			v := flattenWirelessControllerVapDynamicMappingRates11AxSs34(i["rates-11ax-ss34"], d, pre_append)
			tmp["rates_11ax_ss34"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Rates11AxSs34")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11be_mcs_map"
		if _, ok := i["rates-11be-mcs-map"]; ok {
			v := flattenWirelessControllerVapDynamicMappingRates11BeMcsMap(i["rates-11be-mcs-map"], d, pre_append)
			tmp["rates_11be_mcs_map"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Rates11BeMcsMap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11be_mcs_map_160"
		if _, ok := i["rates-11be-mcs-map-160"]; ok {
			v := flattenWirelessControllerVapDynamicMappingRates11BeMcsMap160(i["rates-11be-mcs-map-160"], d, pre_append)
			tmp["rates_11be_mcs_map_160"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Rates11BeMcsMap160")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11be_mcs_map_320"
		if _, ok := i["rates-11be-mcs-map-320"]; ok {
			v := flattenWirelessControllerVapDynamicMappingRates11BeMcsMap320(i["rates-11be-mcs-map-320"], d, pre_append)
			tmp["rates_11be_mcs_map_320"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Rates11BeMcsMap320")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11bg"
		if _, ok := i["rates-11bg"]; ok {
			v := flattenWirelessControllerVapDynamicMappingRates11Bg(i["rates-11bg"], d, pre_append)
			tmp["rates_11bg"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Rates11Bg")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11n_ss12"
		if _, ok := i["rates-11n-ss12"]; ok {
			v := flattenWirelessControllerVapDynamicMappingRates11NSs12(i["rates-11n-ss12"], d, pre_append)
			tmp["rates_11n_ss12"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Rates11NSs12")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11n_ss34"
		if _, ok := i["rates-11n-ss34"]; ok {
			v := flattenWirelessControllerVapDynamicMappingRates11NSs34(i["rates-11n-ss34"], d, pre_append)
			tmp["rates_11n_ss34"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Rates11NSs34")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "roaming_acct_interim_update"
		if _, ok := i["roaming-acct-interim-update"]; ok {
			v := flattenWirelessControllerVapDynamicMappingRoamingAcctInterimUpdate(i["roaming-acct-interim-update"], d, pre_append)
			tmp["roaming_acct_interim_update"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-RoamingAcctInterimUpdate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_groups"
		if _, ok := i["sae-groups"]; ok {
			v := flattenWirelessControllerVapDynamicMappingSaeGroups(i["sae-groups"], d, pre_append)
			tmp["sae_groups"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-SaeGroups")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_h2e_only"
		if _, ok := i["sae-h2e-only"]; ok {
			v := flattenWirelessControllerVapDynamicMappingSaeH2EOnly(i["sae-h2e-only"], d, pre_append)
			tmp["sae_h2e_only"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-SaeH2EOnly")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_hnp_only"
		if _, ok := i["sae-hnp-only"]; ok {
			v := flattenWirelessControllerVapDynamicMappingSaeHnpOnly(i["sae-hnp-only"], d, pre_append)
			tmp["sae_hnp_only"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-SaeHnpOnly")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_pk"
		if _, ok := i["sae-pk"]; ok {
			v := flattenWirelessControllerVapDynamicMappingSaePk(i["sae-pk"], d, pre_append)
			tmp["sae_pk"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-SaePk")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_private_key"
		if _, ok := i["sae-private-key"]; ok {
			v := flattenWirelessControllerVapDynamicMappingSaePrivateKey(i["sae-private-key"], d, pre_append)
			tmp["sae_private_key"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-SaePrivateKey")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "scan_botnet_connections"
		if _, ok := i["scan-botnet-connections"]; ok {
			v := flattenWirelessControllerVapDynamicMappingScanBotnetConnections(i["scan-botnet-connections"], d, pre_append)
			tmp["scan_botnet_connections"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-ScanBotnetConnections")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "schedule"
		if _, ok := i["schedule"]; ok {
			v := flattenWirelessControllerVapDynamicMappingSchedule(i["schedule"], d, pre_append)
			tmp["schedule"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Schedule")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secondary_wag_profile"
		if _, ok := i["secondary-wag-profile"]; ok {
			v := flattenWirelessControllerVapDynamicMappingSecondaryWagProfile(i["secondary-wag-profile"], d, pre_append)
			tmp["secondary_wag_profile"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-SecondaryWagProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security"
		if _, ok := i["security"]; ok {
			v := flattenWirelessControllerVapDynamicMappingSecurity(i["security"], d, pre_append)
			tmp["security"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Security")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_exempt_list"
		if _, ok := i["security-exempt-list"]; ok {
			v := flattenWirelessControllerVapDynamicMappingSecurityExemptList(i["security-exempt-list"], d, pre_append)
			tmp["security_exempt_list"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-SecurityExemptList")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_obsolete_option"
		if _, ok := i["security-obsolete-option"]; ok {
			v := flattenWirelessControllerVapDynamicMappingSecurityObsoleteOption(i["security-obsolete-option"], d, pre_append)
			tmp["security_obsolete_option"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-SecurityObsoleteOption")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_redirect_url"
		if _, ok := i["security-redirect-url"]; ok {
			v := flattenWirelessControllerVapDynamicMappingSecurityRedirectUrl(i["security-redirect-url"], d, pre_append)
			tmp["security_redirect_url"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-SecurityRedirectUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "selected_usergroups"
		if _, ok := i["selected-usergroups"]; ok {
			v := flattenWirelessControllerVapDynamicMappingSelectedUsergroups(i["selected-usergroups"], d, pre_append)
			tmp["selected_usergroups"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-SelectedUsergroups")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_tunneling"
		if _, ok := i["split-tunneling"]; ok {
			v := flattenWirelessControllerVapDynamicMappingSplitTunneling(i["split-tunneling"], d, pre_append)
			tmp["split_tunneling"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-SplitTunneling")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssid"
		if _, ok := i["ssid"]; ok {
			v := flattenWirelessControllerVapDynamicMappingSsid(i["ssid"], d, pre_append)
			tmp["ssid"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Ssid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sticky_client_remove"
		if _, ok := i["sticky-client-remove"]; ok {
			v := flattenWirelessControllerVapDynamicMappingStickyClientRemove(i["sticky-client-remove"], d, pre_append)
			tmp["sticky_client_remove"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-StickyClientRemove")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sticky_client_threshold_2g"
		if _, ok := i["sticky-client-threshold-2g"]; ok {
			v := flattenWirelessControllerVapDynamicMappingStickyClientThreshold2G(i["sticky-client-threshold-2g"], d, pre_append)
			tmp["sticky_client_threshold_2g"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-StickyClientThreshold2G")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sticky_client_threshold_5g"
		if _, ok := i["sticky-client-threshold-5g"]; ok {
			v := flattenWirelessControllerVapDynamicMappingStickyClientThreshold5G(i["sticky-client-threshold-5g"], d, pre_append)
			tmp["sticky_client_threshold_5g"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-StickyClientThreshold5G")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sticky_client_threshold_6g"
		if _, ok := i["sticky-client-threshold-6g"]; ok {
			v := flattenWirelessControllerVapDynamicMappingStickyClientThreshold6G(i["sticky-client-threshold-6g"], d, pre_append)
			tmp["sticky_client_threshold_6g"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-StickyClientThreshold6G")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target_wake_time"
		if _, ok := i["target-wake-time"]; ok {
			v := flattenWirelessControllerVapDynamicMappingTargetWakeTime(i["target-wake-time"], d, pre_append)
			tmp["target_wake_time"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-TargetWakeTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tkip_counter_measure"
		if _, ok := i["tkip-counter-measure"]; ok {
			v := flattenWirelessControllerVapDynamicMappingTkipCounterMeasure(i["tkip-counter-measure"], d, pre_append)
			tmp["tkip_counter_measure"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-TkipCounterMeasure")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tunnel_echo_interval"
		if _, ok := i["tunnel-echo-interval"]; ok {
			v := flattenWirelessControllerVapDynamicMappingTunnelEchoInterval(i["tunnel-echo-interval"], d, pre_append)
			tmp["tunnel_echo_interval"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-TunnelEchoInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tunnel_fallback_interval"
		if _, ok := i["tunnel-fallback-interval"]; ok {
			v := flattenWirelessControllerVapDynamicMappingTunnelFallbackInterval(i["tunnel-fallback-interval"], d, pre_append)
			tmp["tunnel_fallback_interval"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-TunnelFallbackInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "usergroup"
		if _, ok := i["usergroup"]; ok {
			v := flattenWirelessControllerVapDynamicMappingUsergroup(i["usergroup"], d, pre_append)
			tmp["usergroup"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Usergroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "utm_log"
		if _, ok := i["utm-log"]; ok {
			v := flattenWirelessControllerVapDynamicMappingUtmLog(i["utm-log"], d, pre_append)
			tmp["utm_log"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-UtmLog")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "utm_profile"
		if _, ok := i["utm-profile"]; ok {
			v := flattenWirelessControllerVapDynamicMappingUtmProfile(i["utm-profile"], d, pre_append)
			tmp["utm_profile"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-UtmProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "utm_status"
		if _, ok := i["utm-status"]; ok {
			v := flattenWirelessControllerVapDynamicMappingUtmStatus(i["utm-status"], d, pre_append)
			tmp["utm_status"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-UtmStatus")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenWirelessControllerVapDynamicMappingVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Vdom")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_auto"
		if _, ok := i["vlan-auto"]; ok {
			v := flattenWirelessControllerVapDynamicMappingVlanAuto(i["vlan-auto"], d, pre_append)
			tmp["vlan_auto"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-VlanAuto")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_pooling"
		if _, ok := i["vlan-pooling"]; ok {
			v := flattenWirelessControllerVapDynamicMappingVlanPooling(i["vlan-pooling"], d, pre_append)
			tmp["vlan_pooling"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-VlanPooling")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlanid"
		if _, ok := i["vlanid"]; ok {
			v := flattenWirelessControllerVapDynamicMappingVlanid(i["vlanid"], d, pre_append)
			tmp["vlanid"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-Vlanid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "voice_enterprise"
		if _, ok := i["voice-enterprise"]; ok {
			v := flattenWirelessControllerVapDynamicMappingVoiceEnterprise(i["voice-enterprise"], d, pre_append)
			tmp["voice_enterprise"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-VoiceEnterprise")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "webfilter_profile"
		if _, ok := i["webfilter-profile"]; ok {
			v := flattenWirelessControllerVapDynamicMappingWebfilterProfile(i["webfilter-profile"], d, pre_append)
			tmp["webfilter_profile"] = fortiAPISubPartPatch(v, "WirelessControllerVap-DynamicMapping-WebfilterProfile")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerVapDynamicMapping80211K(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMapping80211V(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingCentmgmt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingDhcpSvrId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingIntfAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingIntfDeviceAccessList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingIntfDeviceIdentification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingIntfDeviceNetscan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingIntfDhcpRelayIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingIntfDhcpRelayService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingIntfDhcpRelayType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingIntfDhcp6RelayIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingIntfDhcp6RelayService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingIntfDhcp6RelayType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingIntfIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingIntfIpManagedByFortiipam(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingIntfIp6Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingIntfIp6Allowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingIntfListenForticlientConnection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingIntfManagedSubnetworkSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingIntfRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingIsFactorySetting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenWirelessControllerVapDynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "WirelessControllerVapDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenWirelessControllerVapDynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "WirelessControllerVapDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerVapDynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingAccessControlList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingAcctInterimInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingAdditionalAkms(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingAddressGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingAddressGroupPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingAkm24Only(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingAlias(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingAntivirusProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingApplicationDetectionEngine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingApplicationDscpMarking(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingApplicationList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingApplicationReportIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingAtfWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingAuthCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingAuthPortalAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingBeaconAdvertising(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingBeaconProtection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingBroadcastSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingBroadcastSuppression(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingBssColorPartial(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingBstmDisassociationImminent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingBstmLoadBalancingDisassocTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingBstmRssiDisassocTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingCalledStationIdType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingCaptivePortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingCaptivePortalAcName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingCaptivePortalAuthTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingCaptivePortalFwAccounting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingCaptivePortalMacauthRadiusServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingCaptivePortalRadiusServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingCaptivePortalSessionTimeoutInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingClientCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingDhcpAddressEnforcement(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingDhcpLeaseTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingDhcpOption43Insertion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingDhcpOption82CircuitIdInsertion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingDhcpOption82Insertion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingDhcpOption82RemoteIdInsertion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingDomainNameStripping(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingDynamicVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingEapReauth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingEapReauthIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingEapolKeyRetries(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingEncrypt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingExternalFastRoaming(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingExternalLogout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingExternalPreAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingExternalWeb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingExternalWebFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingFastBssTransition(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingFastRoaming(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingFtMobilityDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingFtOverDs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingFtR0KeyLifetime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingGasComebackDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingGasFragmentationLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingGtkRekey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingGtkRekeyIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingHighEfficiency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingHotspot20Profile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingIgmpSnooping(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingIntraVapPrivacy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingIpsSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingIpv6Rules(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingKeyindex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingL3Roaming(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingL3RoamingMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingLdpc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingLocalAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingLocalBridging(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingLocalLan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingLocalLanPartition(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingLocalStandalone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingLocalStandaloneDns(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingLocalStandaloneDnsIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingLocalStandaloneNat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingLocalSwitching(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingMacAuthBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingMacCalledStationDelimiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingMacCallingStationDelimiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingMacCase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingMacFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingMacFilterPolicyOther(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingMacPasswordDelimiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingMacUsernameDelimiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingMaxClients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingMaxClientsAp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingMbo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingMboCellDataConnPref(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingMeDisableThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingMeshBackhaul(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingMpsk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingMpskConcurrentClients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingMpskProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingMuMimo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingMulticastEnhance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingMulticastRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingNac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingNacProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingNasFilterRule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingNeighborReportDualBand(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingOkc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingOsen(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingOweGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingOweTransition(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingOweTransitionSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingPmf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingPmfAssocComebackTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingPmfSaQueryRetryTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingPortMacauth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingPortMacauthReauthTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingPortMacauthTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingPortalMessageOverrideGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingPortalType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingPreAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingPrimaryWagProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingProbeRespSuppression(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingProbeRespThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingPtkRekey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingPtkRekeyIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingQosProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingRadio2GThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingRadio5GThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingRadioSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingRadiusMacAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingRadiusMacAuthBlockInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingRadiusMacAuthServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingRadiusMacAuthUsergroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingRadiusMacMpskAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingRadiusMacMpskTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingRadiusServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingRates11A(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingRates11AcMcsMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingRates11AcSs12(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingRates11AcSs34(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingRates11AxMcsMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingRates11AxSs12(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingRates11AxSs34(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingRates11BeMcsMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingRates11BeMcsMap160(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingRates11BeMcsMap320(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingRates11Bg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingRates11NSs12(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingRates11NSs34(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingRoamingAcctInterimUpdate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingSaeGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingSaeH2EOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingSaeHnpOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingSaePk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingSaePrivateKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingScanBotnetConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingSecondaryWagProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingSecurity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingSecurityExemptList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingSecurityObsoleteOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingSecurityRedirectUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingSelectedUsergroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingSplitTunneling(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingStickyClientRemove(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingStickyClientThreshold2G(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingStickyClientThreshold5G(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingStickyClientThreshold6G(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingTargetWakeTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingTkipCounterMeasure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingTunnelEchoInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingTunnelFallbackInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingUsergroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingUtmLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingUtmProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingUtmStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingVlanAuto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingVlanPooling(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingVlanid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingVoiceEnterprise(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingWebfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapEapReauth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapEapReauthIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapEapolKeyRetries(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapEncrypt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapExternalFastRoaming(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapExternalLogout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapExternalPreAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapExternalWeb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapExternalWebFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapFastBssTransition(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapFastRoaming(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapFtMobilityDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapFtOverDs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapFtR0KeyLifetime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapGasComebackDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapGasFragmentationLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapGtkRekey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapGtkRekeyIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapHighEfficiency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapHotspot20Profile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapIgmpSnooping(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapIntraVapPrivacy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapIpsSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapIpv6Rules(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapKeyindex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapL3Roaming(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapL3RoamingMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapLdpc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapLocalAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapLocalBridging(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapLocalLan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapLocalLanPartition(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapLocalStandalone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapLocalStandaloneDns(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapLocalStandaloneDnsIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapLocalStandaloneNat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMacAuthBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMacCalledStationDelimiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMacCallingStationDelimiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMacCase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMacFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMacFilterList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWirelessControllerVapMacFilterListId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WirelessControllerVap-MacFilterList-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			v := flattenWirelessControllerVapMacFilterListMac(i["mac"], d, pre_append)
			tmp["mac"] = fortiAPISubPartPatch(v, "WirelessControllerVap-MacFilterList-Mac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_filter_policy"
		if _, ok := i["mac-filter-policy"]; ok {
			v := flattenWirelessControllerVapMacFilterListMacFilterPolicy(i["mac-filter-policy"], d, pre_append)
			tmp["mac_filter_policy"] = fortiAPISubPartPatch(v, "WirelessControllerVap-MacFilterList-MacFilterPolicy")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerVapMacFilterListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMacFilterListMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMacFilterListMacFilterPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMacFilterPolicyOther(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMacPasswordDelimiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMacUsernameDelimiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMaxClients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMaxClientsAp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMbo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMboCellDataConnPref(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMeDisableThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMeshBackhaul(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMpsk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMpskConcurrentClients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMpskKey(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenWirelessControllerVapMpskKeyComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "WirelessControllerVap-MpskKey-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "concurrent_clients"
		if _, ok := i["concurrent-clients"]; ok {
			v := flattenWirelessControllerVapMpskKeyConcurrentClients(i["concurrent-clients"], d, pre_append)
			tmp["concurrent_clients"] = fortiAPISubPartPatch(v, "WirelessControllerVap-MpskKey-ConcurrentClients")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_name"
		if _, ok := i["key-name"]; ok {
			v := flattenWirelessControllerVapMpskKeyKeyName(i["key-name"], d, pre_append)
			tmp["key_name"] = fortiAPISubPartPatch(v, "WirelessControllerVap-MpskKey-KeyName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk_schedules"
		if _, ok := i["mpsk-schedules"]; ok {
			v := flattenWirelessControllerVapMpskKeyMpskSchedules(i["mpsk-schedules"], d, pre_append)
			tmp["mpsk_schedules"] = fortiAPISubPartPatch(v, "WirelessControllerVap-MpskKey-MpskSchedules")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerVapMpskKeyComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMpskKeyConcurrentClients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMpskKeyKeyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMpskKeyMpskSchedules(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapMpskProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapMuMimo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMulticastEnhance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMulticastRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapNac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapNacProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapNasFilterRule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapNeighborReportDualBand(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapOkc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapOsen(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapOweGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapOweTransition(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapOweTransitionSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapPmf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapPmfAssocComebackTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapPmfSaQueryRetryTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapPortMacauth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapPortMacauthReauthTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapPortMacauthTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapPortalMessageOverrideGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapPortalMessageOverrides(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auth_disclaimer_page"
	if _, ok := i["auth-disclaimer-page"]; ok {
		result["auth_disclaimer_page"] = flattenWirelessControllerVapPortalMessageOverridesAuthDisclaimerPage(i["auth-disclaimer-page"], d, pre_append)
	}

	pre_append = pre + ".0." + "auth_login_failed_page"
	if _, ok := i["auth-login-failed-page"]; ok {
		result["auth_login_failed_page"] = flattenWirelessControllerVapPortalMessageOverridesAuthLoginFailedPage(i["auth-login-failed-page"], d, pre_append)
	}

	pre_append = pre + ".0." + "auth_login_page"
	if _, ok := i["auth-login-page"]; ok {
		result["auth_login_page"] = flattenWirelessControllerVapPortalMessageOverridesAuthLoginPage(i["auth-login-page"], d, pre_append)
	}

	pre_append = pre + ".0." + "auth_reject_page"
	if _, ok := i["auth-reject-page"]; ok {
		result["auth_reject_page"] = flattenWirelessControllerVapPortalMessageOverridesAuthRejectPage(i["auth-reject-page"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWirelessControllerVapPortalMessageOverridesAuthDisclaimerPage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapPortalMessageOverridesAuthLoginFailedPage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapPortalMessageOverridesAuthLoginPage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapPortalMessageOverridesAuthRejectPage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapPortalType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapPreAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapPrimaryWagProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapProbeRespSuppression(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapProbeRespThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapPtkRekey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapPtkRekeyIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapQosProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapRadio2GThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapRadio5GThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapRadioSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapRadiusMacAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapRadiusMacAuthBlockInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapRadiusMacAuthServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapRadiusMacAuthUsergroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapRadiusMacMpskAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapRadiusMacMpskTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapRadiusServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapRates11A(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapRates11AcMcsMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapRates11AcSs12(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapRates11AcSs34(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapRates11AxMcsMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapRates11AxSs12(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapRates11AxSs34(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapRates11BeMcsMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapRates11BeMcsMap160(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapRates11BeMcsMap320(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapRates11Bg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapRates11NSs12(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapRates11NSs34(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapRoamingAcctInterimUpdate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapSaeGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapSaeH2EOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapSaeHnpOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapSaePk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapSaePrivateKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapScanBotnetConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapSecondaryWagProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapSecurity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapSecurityExemptList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapSecurityObsoleteOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapSecurityRedirectUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapSelectedUsergroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapSplitTunneling(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapStickyClientRemove(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapStickyClientThreshold2G(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapStickyClientThreshold5G(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapStickyClientThreshold6G(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapTargetWakeTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapTkipCounterMeasure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapTunnelEchoInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapTunnelFallbackInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapUsergroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapUtmLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapUtmProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapUtmStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapVlanAuto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapVlanName(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenWirelessControllerVapVlanNameName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "WirelessControllerVap-VlanName-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_id"
		if _, ok := i["vlan-id"]; ok {
			v := flattenWirelessControllerVapVlanNameVlanId(i["vlan-id"], d, pre_append)
			tmp["vlan_id"] = fortiAPISubPartPatch(v, "WirelessControllerVap-VlanName-VlanId")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerVapVlanNameName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapVlanNameVlanId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapVlanPool(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWirelessControllerVapVlanPoolId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WirelessControllerVap-VlanPool-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wtp_group"
		if _, ok := i["wtp-group"]; ok {
			v := flattenWirelessControllerVapVlanPoolWtpGroup(i["wtp-group"], d, pre_append)
			tmp["wtp_group"] = fortiAPISubPartPatch(v, "WirelessControllerVap-VlanPool-WtpGroup")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerVapVlanPoolId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapVlanPoolWtpGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapVlanPooling(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapVlanid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapVoiceEnterprise(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapWebfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectWirelessControllerVap(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("n80211k", flattenWirelessControllerVap80211K(o["80211k"], d, "n80211k")); err != nil {
		if vv, ok := fortiAPIPatch(o["80211k"], "WirelessControllerVap-80211K"); ok {
			if err = d.Set("n80211k", vv); err != nil {
				return fmt.Errorf("Error reading n80211k: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading n80211k: %v", err)
		}
	}

	if err = d.Set("n80211v", flattenWirelessControllerVap80211V(o["80211v"], d, "n80211v")); err != nil {
		if vv, ok := fortiAPIPatch(o["80211v"], "WirelessControllerVap-80211V"); ok {
			if err = d.Set("n80211v", vv); err != nil {
				return fmt.Errorf("Error reading n80211v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading n80211v: %v", err)
		}
	}

	if err = d.Set("_centmgmt", flattenWirelessControllerVapCentmgmt(o["_centmgmt"], d, "_centmgmt")); err != nil {
		if vv, ok := fortiAPIPatch(o["_centmgmt"], "WirelessControllerVap-Centmgmt"); ok {
			if err = d.Set("_centmgmt", vv); err != nil {
				return fmt.Errorf("Error reading _centmgmt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _centmgmt: %v", err)
		}
	}

	if err = d.Set("_dhcp_svr_id", flattenWirelessControllerVapDhcpSvrId(o["_dhcp_svr_id"], d, "_dhcp_svr_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["_dhcp_svr_id"], "WirelessControllerVap-DhcpSvrId"); ok {
			if err = d.Set("_dhcp_svr_id", vv); err != nil {
				return fmt.Errorf("Error reading _dhcp_svr_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _dhcp_svr_id: %v", err)
		}
	}

	if err = d.Set("_intf_allowaccess", flattenWirelessControllerVapIntfAllowaccess(o["_intf_allowaccess"], d, "_intf_allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_allowaccess"], "WirelessControllerVap-IntfAllowaccess"); ok {
			if err = d.Set("_intf_allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading _intf_allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_allowaccess: %v", err)
		}
	}

	if err = d.Set("_intf_device_access_list", flattenWirelessControllerVapIntfDeviceAccessList(o["_intf_device-access-list"], d, "_intf_device_access_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_device-access-list"], "WirelessControllerVap-IntfDeviceAccessList"); ok {
			if err = d.Set("_intf_device_access_list", vv); err != nil {
				return fmt.Errorf("Error reading _intf_device_access_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_device_access_list: %v", err)
		}
	}

	if err = d.Set("_intf_device_identification", flattenWirelessControllerVapIntfDeviceIdentification(o["_intf_device-identification"], d, "_intf_device_identification")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_device-identification"], "WirelessControllerVap-IntfDeviceIdentification"); ok {
			if err = d.Set("_intf_device_identification", vv); err != nil {
				return fmt.Errorf("Error reading _intf_device_identification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_device_identification: %v", err)
		}
	}

	if err = d.Set("_intf_device_netscan", flattenWirelessControllerVapIntfDeviceNetscan(o["_intf_device-netscan"], d, "_intf_device_netscan")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_device-netscan"], "WirelessControllerVap-IntfDeviceNetscan"); ok {
			if err = d.Set("_intf_device_netscan", vv); err != nil {
				return fmt.Errorf("Error reading _intf_device_netscan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_device_netscan: %v", err)
		}
	}

	if err = d.Set("_intf_dhcp_relay_ip", flattenWirelessControllerVapIntfDhcpRelayIp(o["_intf_dhcp-relay-ip"], d, "_intf_dhcp_relay_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_dhcp-relay-ip"], "WirelessControllerVap-IntfDhcpRelayIp"); ok {
			if err = d.Set("_intf_dhcp_relay_ip", vv); err != nil {
				return fmt.Errorf("Error reading _intf_dhcp_relay_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_dhcp_relay_ip: %v", err)
		}
	}

	if err = d.Set("_intf_dhcp_relay_service", flattenWirelessControllerVapIntfDhcpRelayService(o["_intf_dhcp-relay-service"], d, "_intf_dhcp_relay_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_dhcp-relay-service"], "WirelessControllerVap-IntfDhcpRelayService"); ok {
			if err = d.Set("_intf_dhcp_relay_service", vv); err != nil {
				return fmt.Errorf("Error reading _intf_dhcp_relay_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_dhcp_relay_service: %v", err)
		}
	}

	if err = d.Set("_intf_dhcp_relay_type", flattenWirelessControllerVapIntfDhcpRelayType(o["_intf_dhcp-relay-type"], d, "_intf_dhcp_relay_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_dhcp-relay-type"], "WirelessControllerVap-IntfDhcpRelayType"); ok {
			if err = d.Set("_intf_dhcp_relay_type", vv); err != nil {
				return fmt.Errorf("Error reading _intf_dhcp_relay_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_dhcp_relay_type: %v", err)
		}
	}

	if err = d.Set("_intf_dhcp6_relay_ip", flattenWirelessControllerVapIntfDhcp6RelayIp(o["_intf_dhcp6-relay-ip"], d, "_intf_dhcp6_relay_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_dhcp6-relay-ip"], "WirelessControllerVap-IntfDhcp6RelayIp"); ok {
			if err = d.Set("_intf_dhcp6_relay_ip", vv); err != nil {
				return fmt.Errorf("Error reading _intf_dhcp6_relay_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_dhcp6_relay_ip: %v", err)
		}
	}

	if err = d.Set("_intf_dhcp6_relay_service", flattenWirelessControllerVapIntfDhcp6RelayService(o["_intf_dhcp6-relay-service"], d, "_intf_dhcp6_relay_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_dhcp6-relay-service"], "WirelessControllerVap-IntfDhcp6RelayService"); ok {
			if err = d.Set("_intf_dhcp6_relay_service", vv); err != nil {
				return fmt.Errorf("Error reading _intf_dhcp6_relay_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_dhcp6_relay_service: %v", err)
		}
	}

	if err = d.Set("_intf_dhcp6_relay_type", flattenWirelessControllerVapIntfDhcp6RelayType(o["_intf_dhcp6-relay-type"], d, "_intf_dhcp6_relay_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_dhcp6-relay-type"], "WirelessControllerVap-IntfDhcp6RelayType"); ok {
			if err = d.Set("_intf_dhcp6_relay_type", vv); err != nil {
				return fmt.Errorf("Error reading _intf_dhcp6_relay_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_dhcp6_relay_type: %v", err)
		}
	}

	if err = d.Set("_intf_ip", flattenWirelessControllerVapIntfIp(o["_intf_ip"], d, "_intf_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_ip"], "WirelessControllerVap-IntfIp"); ok {
			if err = d.Set("_intf_ip", vv); err != nil {
				return fmt.Errorf("Error reading _intf_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_ip: %v", err)
		}
	}

	if err = d.Set("_intf_ip_managed_by_fortiipam", flattenWirelessControllerVapIntfIpManagedByFortiipam(o["_intf_ip-managed-by-fortiipam"], d, "_intf_ip_managed_by_fortiipam")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_ip-managed-by-fortiipam"], "WirelessControllerVap-IntfIpManagedByFortiipam"); ok {
			if err = d.Set("_intf_ip_managed_by_fortiipam", vv); err != nil {
				return fmt.Errorf("Error reading _intf_ip_managed_by_fortiipam: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_ip_managed_by_fortiipam: %v", err)
		}
	}

	if err = d.Set("_intf_ip6_address", flattenWirelessControllerVapIntfIp6Address(o["_intf_ip6-address"], d, "_intf_ip6_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_ip6-address"], "WirelessControllerVap-IntfIp6Address"); ok {
			if err = d.Set("_intf_ip6_address", vv); err != nil {
				return fmt.Errorf("Error reading _intf_ip6_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_ip6_address: %v", err)
		}
	}

	if err = d.Set("_intf_ip6_allowaccess", flattenWirelessControllerVapIntfIp6Allowaccess(o["_intf_ip6-allowaccess"], d, "_intf_ip6_allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_ip6-allowaccess"], "WirelessControllerVap-IntfIp6Allowaccess"); ok {
			if err = d.Set("_intf_ip6_allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading _intf_ip6_allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_ip6_allowaccess: %v", err)
		}
	}

	if err = d.Set("_intf_listen_forticlient_connection", flattenWirelessControllerVapIntfListenForticlientConnection(o["_intf_listen-forticlient-connection"], d, "_intf_listen_forticlient_connection")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_listen-forticlient-connection"], "WirelessControllerVap-IntfListenForticlientConnection"); ok {
			if err = d.Set("_intf_listen_forticlient_connection", vv); err != nil {
				return fmt.Errorf("Error reading _intf_listen_forticlient_connection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_listen_forticlient_connection: %v", err)
		}
	}

	if err = d.Set("_intf_managed_subnetwork_size", flattenWirelessControllerVapIntfManagedSubnetworkSize(o["_intf_managed-subnetwork-size"], d, "_intf_managed_subnetwork_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_managed-subnetwork-size"], "WirelessControllerVap-IntfManagedSubnetworkSize"); ok {
			if err = d.Set("_intf_managed_subnetwork_size", vv); err != nil {
				return fmt.Errorf("Error reading _intf_managed_subnetwork_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_managed_subnetwork_size: %v", err)
		}
	}

	if err = d.Set("_intf_role", flattenWirelessControllerVapIntfRole(o["_intf_role"], d, "_intf_role")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_role"], "WirelessControllerVap-IntfRole"); ok {
			if err = d.Set("_intf_role", vv); err != nil {
				return fmt.Errorf("Error reading _intf_role: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_role: %v", err)
		}
	}

	if err = d.Set("_is_factory_setting", flattenWirelessControllerVapIsFactorySetting(o["_is_factory_setting"], d, "_is_factory_setting")); err != nil {
		if vv, ok := fortiAPIPatch(o["_is_factory_setting"], "WirelessControllerVap-IsFactorySetting"); ok {
			if err = d.Set("_is_factory_setting", vv); err != nil {
				return fmt.Errorf("Error reading _is_factory_setting: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _is_factory_setting: %v", err)
		}
	}

	if err = d.Set("acct_interim_interval", flattenWirelessControllerVapAcctInterimInterval(o["acct-interim-interval"], d, "acct_interim_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["acct-interim-interval"], "WirelessControllerVap-AcctInterimInterval"); ok {
			if err = d.Set("acct_interim_interval", vv); err != nil {
				return fmt.Errorf("Error reading acct_interim_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading acct_interim_interval: %v", err)
		}
	}

	if err = d.Set("access_control_list", flattenWirelessControllerVapAccessControlList(o["access-control-list"], d, "access_control_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-control-list"], "WirelessControllerVap-AccessControlList"); ok {
			if err = d.Set("access_control_list", vv); err != nil {
				return fmt.Errorf("Error reading access_control_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_control_list: %v", err)
		}
	}

	if err = d.Set("additional_akms", flattenWirelessControllerVapAdditionalAkms(o["additional-akms"], d, "additional_akms")); err != nil {
		if vv, ok := fortiAPIPatch(o["additional-akms"], "WirelessControllerVap-AdditionalAkms"); ok {
			if err = d.Set("additional_akms", vv); err != nil {
				return fmt.Errorf("Error reading additional_akms: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading additional_akms: %v", err)
		}
	}

	if err = d.Set("address_group", flattenWirelessControllerVapAddressGroup(o["address-group"], d, "address_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["address-group"], "WirelessControllerVap-AddressGroup"); ok {
			if err = d.Set("address_group", vv); err != nil {
				return fmt.Errorf("Error reading address_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address_group: %v", err)
		}
	}

	if err = d.Set("address_group_policy", flattenWirelessControllerVapAddressGroupPolicy(o["address-group-policy"], d, "address_group_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["address-group-policy"], "WirelessControllerVap-AddressGroupPolicy"); ok {
			if err = d.Set("address_group_policy", vv); err != nil {
				return fmt.Errorf("Error reading address_group_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address_group_policy: %v", err)
		}
	}

	if err = d.Set("akm24_only", flattenWirelessControllerVapAkm24Only(o["akm24-only"], d, "akm24_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["akm24-only"], "WirelessControllerVap-Akm24Only"); ok {
			if err = d.Set("akm24_only", vv); err != nil {
				return fmt.Errorf("Error reading akm24_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading akm24_only: %v", err)
		}
	}

	if err = d.Set("alias", flattenWirelessControllerVapAlias(o["alias"], d, "alias")); err != nil {
		if vv, ok := fortiAPIPatch(o["alias"], "WirelessControllerVap-Alias"); ok {
			if err = d.Set("alias", vv); err != nil {
				return fmt.Errorf("Error reading alias: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alias: %v", err)
		}
	}

	if err = d.Set("antivirus_profile", flattenWirelessControllerVapAntivirusProfile(o["antivirus-profile"], d, "antivirus_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["antivirus-profile"], "WirelessControllerVap-AntivirusProfile"); ok {
			if err = d.Set("antivirus_profile", vv); err != nil {
				return fmt.Errorf("Error reading antivirus_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading antivirus_profile: %v", err)
		}
	}

	if err = d.Set("application_detection_engine", flattenWirelessControllerVapApplicationDetectionEngine(o["application-detection-engine"], d, "application_detection_engine")); err != nil {
		if vv, ok := fortiAPIPatch(o["application-detection-engine"], "WirelessControllerVap-ApplicationDetectionEngine"); ok {
			if err = d.Set("application_detection_engine", vv); err != nil {
				return fmt.Errorf("Error reading application_detection_engine: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application_detection_engine: %v", err)
		}
	}

	if err = d.Set("application_dscp_marking", flattenWirelessControllerVapApplicationDscpMarking(o["application-dscp-marking"], d, "application_dscp_marking")); err != nil {
		if vv, ok := fortiAPIPatch(o["application-dscp-marking"], "WirelessControllerVap-ApplicationDscpMarking"); ok {
			if err = d.Set("application_dscp_marking", vv); err != nil {
				return fmt.Errorf("Error reading application_dscp_marking: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application_dscp_marking: %v", err)
		}
	}

	if err = d.Set("application_list", flattenWirelessControllerVapApplicationList(o["application-list"], d, "application_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["application-list"], "WirelessControllerVap-ApplicationList"); ok {
			if err = d.Set("application_list", vv); err != nil {
				return fmt.Errorf("Error reading application_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application_list: %v", err)
		}
	}

	if err = d.Set("application_report_intv", flattenWirelessControllerVapApplicationReportIntv(o["application-report-intv"], d, "application_report_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["application-report-intv"], "WirelessControllerVap-ApplicationReportIntv"); ok {
			if err = d.Set("application_report_intv", vv); err != nil {
				return fmt.Errorf("Error reading application_report_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application_report_intv: %v", err)
		}
	}

	if err = d.Set("atf_weight", flattenWirelessControllerVapAtfWeight(o["atf-weight"], d, "atf_weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["atf-weight"], "WirelessControllerVap-AtfWeight"); ok {
			if err = d.Set("atf_weight", vv); err != nil {
				return fmt.Errorf("Error reading atf_weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading atf_weight: %v", err)
		}
	}

	if err = d.Set("auth", flattenWirelessControllerVapAuth(o["auth"], d, "auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth"], "WirelessControllerVap-Auth"); ok {
			if err = d.Set("auth", vv); err != nil {
				return fmt.Errorf("Error reading auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth: %v", err)
		}
	}

	if err = d.Set("auth_cert", flattenWirelessControllerVapAuthCert(o["auth-cert"], d, "auth_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-cert"], "WirelessControllerVap-AuthCert"); ok {
			if err = d.Set("auth_cert", vv); err != nil {
				return fmt.Errorf("Error reading auth_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_cert: %v", err)
		}
	}

	if err = d.Set("auth_portal_addr", flattenWirelessControllerVapAuthPortalAddr(o["auth-portal-addr"], d, "auth_portal_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-portal-addr"], "WirelessControllerVap-AuthPortalAddr"); ok {
			if err = d.Set("auth_portal_addr", vv); err != nil {
				return fmt.Errorf("Error reading auth_portal_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_portal_addr: %v", err)
		}
	}

	if err = d.Set("beacon_advertising", flattenWirelessControllerVapBeaconAdvertising(o["beacon-advertising"], d, "beacon_advertising")); err != nil {
		if vv, ok := fortiAPIPatch(o["beacon-advertising"], "WirelessControllerVap-BeaconAdvertising"); ok {
			if err = d.Set("beacon_advertising", vv); err != nil {
				return fmt.Errorf("Error reading beacon_advertising: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading beacon_advertising: %v", err)
		}
	}

	if err = d.Set("beacon_protection", flattenWirelessControllerVapBeaconProtection(o["beacon-protection"], d, "beacon_protection")); err != nil {
		if vv, ok := fortiAPIPatch(o["beacon-protection"], "WirelessControllerVap-BeaconProtection"); ok {
			if err = d.Set("beacon_protection", vv); err != nil {
				return fmt.Errorf("Error reading beacon_protection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading beacon_protection: %v", err)
		}
	}

	if err = d.Set("broadcast_ssid", flattenWirelessControllerVapBroadcastSsid(o["broadcast-ssid"], d, "broadcast_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["broadcast-ssid"], "WirelessControllerVap-BroadcastSsid"); ok {
			if err = d.Set("broadcast_ssid", vv); err != nil {
				return fmt.Errorf("Error reading broadcast_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading broadcast_ssid: %v", err)
		}
	}

	if err = d.Set("broadcast_suppression", flattenWirelessControllerVapBroadcastSuppression(o["broadcast-suppression"], d, "broadcast_suppression")); err != nil {
		if vv, ok := fortiAPIPatch(o["broadcast-suppression"], "WirelessControllerVap-BroadcastSuppression"); ok {
			if err = d.Set("broadcast_suppression", vv); err != nil {
				return fmt.Errorf("Error reading broadcast_suppression: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading broadcast_suppression: %v", err)
		}
	}

	if err = d.Set("bss_color_partial", flattenWirelessControllerVapBssColorPartial(o["bss-color-partial"], d, "bss_color_partial")); err != nil {
		if vv, ok := fortiAPIPatch(o["bss-color-partial"], "WirelessControllerVap-BssColorPartial"); ok {
			if err = d.Set("bss_color_partial", vv); err != nil {
				return fmt.Errorf("Error reading bss_color_partial: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bss_color_partial: %v", err)
		}
	}

	if err = d.Set("bstm_disassociation_imminent", flattenWirelessControllerVapBstmDisassociationImminent(o["bstm-disassociation-imminent"], d, "bstm_disassociation_imminent")); err != nil {
		if vv, ok := fortiAPIPatch(o["bstm-disassociation-imminent"], "WirelessControllerVap-BstmDisassociationImminent"); ok {
			if err = d.Set("bstm_disassociation_imminent", vv); err != nil {
				return fmt.Errorf("Error reading bstm_disassociation_imminent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bstm_disassociation_imminent: %v", err)
		}
	}

	if err = d.Set("bstm_load_balancing_disassoc_timer", flattenWirelessControllerVapBstmLoadBalancingDisassocTimer(o["bstm-load-balancing-disassoc-timer"], d, "bstm_load_balancing_disassoc_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["bstm-load-balancing-disassoc-timer"], "WirelessControllerVap-BstmLoadBalancingDisassocTimer"); ok {
			if err = d.Set("bstm_load_balancing_disassoc_timer", vv); err != nil {
				return fmt.Errorf("Error reading bstm_load_balancing_disassoc_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bstm_load_balancing_disassoc_timer: %v", err)
		}
	}

	if err = d.Set("bstm_rssi_disassoc_timer", flattenWirelessControllerVapBstmRssiDisassocTimer(o["bstm-rssi-disassoc-timer"], d, "bstm_rssi_disassoc_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["bstm-rssi-disassoc-timer"], "WirelessControllerVap-BstmRssiDisassocTimer"); ok {
			if err = d.Set("bstm_rssi_disassoc_timer", vv); err != nil {
				return fmt.Errorf("Error reading bstm_rssi_disassoc_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bstm_rssi_disassoc_timer: %v", err)
		}
	}

	if err = d.Set("called_station_id_type", flattenWirelessControllerVapCalledStationIdType(o["called-station-id-type"], d, "called_station_id_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["called-station-id-type"], "WirelessControllerVap-CalledStationIdType"); ok {
			if err = d.Set("called_station_id_type", vv); err != nil {
				return fmt.Errorf("Error reading called_station_id_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading called_station_id_type: %v", err)
		}
	}

	if err = d.Set("captive_portal", flattenWirelessControllerVapCaptivePortal(o["captive-portal"], d, "captive_portal")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal"], "WirelessControllerVap-CaptivePortal"); ok {
			if err = d.Set("captive_portal", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal: %v", err)
		}
	}

	if err = d.Set("captive_portal_ac_name", flattenWirelessControllerVapCaptivePortalAcName(o["captive-portal-ac-name"], d, "captive_portal_ac_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-ac-name"], "WirelessControllerVap-CaptivePortalAcName"); ok {
			if err = d.Set("captive_portal_ac_name", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_ac_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_ac_name: %v", err)
		}
	}

	if err = d.Set("captive_portal_macauth_radius_server", flattenWirelessControllerVapCaptivePortalMacauthRadiusServer(o["captive-portal-macauth-radius-server"], d, "captive_portal_macauth_radius_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-macauth-radius-server"], "WirelessControllerVap-CaptivePortalMacauthRadiusServer"); ok {
			if err = d.Set("captive_portal_macauth_radius_server", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_macauth_radius_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_macauth_radius_server: %v", err)
		}
	}

	if err = d.Set("captive_portal_radius_server", flattenWirelessControllerVapCaptivePortalRadiusServer(o["captive-portal-radius-server"], d, "captive_portal_radius_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-radius-server"], "WirelessControllerVap-CaptivePortalRadiusServer"); ok {
			if err = d.Set("captive_portal_radius_server", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_radius_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_radius_server: %v", err)
		}
	}

	if err = d.Set("captive_portal_session_timeout_interval", flattenWirelessControllerVapCaptivePortalSessionTimeoutInterval(o["captive-portal-session-timeout-interval"], d, "captive_portal_session_timeout_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-session-timeout-interval"], "WirelessControllerVap-CaptivePortalSessionTimeoutInterval"); ok {
			if err = d.Set("captive_portal_session_timeout_interval", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_session_timeout_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_session_timeout_interval: %v", err)
		}
	}

	if err = d.Set("captive_portal_auth_timeout", flattenWirelessControllerVapCaptivePortalAuthTimeout(o["captive-portal-auth-timeout"], d, "captive_portal_auth_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-auth-timeout"], "WirelessControllerVap-CaptivePortalAuthTimeout"); ok {
			if err = d.Set("captive_portal_auth_timeout", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_auth_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_auth_timeout: %v", err)
		}
	}

	if err = d.Set("captive_portal_fw_accounting", flattenWirelessControllerVapCaptivePortalFwAccounting(o["captive-portal-fw-accounting"], d, "captive_portal_fw_accounting")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-fw-accounting"], "WirelessControllerVap-CaptivePortalFwAccounting"); ok {
			if err = d.Set("captive_portal_fw_accounting", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_fw_accounting: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_fw_accounting: %v", err)
		}
	}

	if err = d.Set("dhcp_address_enforcement", flattenWirelessControllerVapDhcpAddressEnforcement(o["dhcp-address-enforcement"], d, "dhcp_address_enforcement")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-address-enforcement"], "WirelessControllerVap-DhcpAddressEnforcement"); ok {
			if err = d.Set("dhcp_address_enforcement", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_address_enforcement: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_address_enforcement: %v", err)
		}
	}

	if err = d.Set("dhcp_lease_time", flattenWirelessControllerVapDhcpLeaseTime(o["dhcp-lease-time"], d, "dhcp_lease_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-lease-time"], "WirelessControllerVap-DhcpLeaseTime"); ok {
			if err = d.Set("dhcp_lease_time", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_lease_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_lease_time: %v", err)
		}
	}

	if err = d.Set("dhcp_option43_insertion", flattenWirelessControllerVapDhcpOption43Insertion(o["dhcp-option43-insertion"], d, "dhcp_option43_insertion")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-option43-insertion"], "WirelessControllerVap-DhcpOption43Insertion"); ok {
			if err = d.Set("dhcp_option43_insertion", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_option43_insertion: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_option43_insertion: %v", err)
		}
	}

	if err = d.Set("dhcp_option82_circuit_id_insertion", flattenWirelessControllerVapDhcpOption82CircuitIdInsertion(o["dhcp-option82-circuit-id-insertion"], d, "dhcp_option82_circuit_id_insertion")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-option82-circuit-id-insertion"], "WirelessControllerVap-DhcpOption82CircuitIdInsertion"); ok {
			if err = d.Set("dhcp_option82_circuit_id_insertion", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_option82_circuit_id_insertion: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_option82_circuit_id_insertion: %v", err)
		}
	}

	if err = d.Set("dhcp_option82_insertion", flattenWirelessControllerVapDhcpOption82Insertion(o["dhcp-option82-insertion"], d, "dhcp_option82_insertion")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-option82-insertion"], "WirelessControllerVap-DhcpOption82Insertion"); ok {
			if err = d.Set("dhcp_option82_insertion", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_option82_insertion: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_option82_insertion: %v", err)
		}
	}

	if err = d.Set("dhcp_option82_remote_id_insertion", flattenWirelessControllerVapDhcpOption82RemoteIdInsertion(o["dhcp-option82-remote-id-insertion"], d, "dhcp_option82_remote_id_insertion")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-option82-remote-id-insertion"], "WirelessControllerVap-DhcpOption82RemoteIdInsertion"); ok {
			if err = d.Set("dhcp_option82_remote_id_insertion", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_option82_remote_id_insertion: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_option82_remote_id_insertion: %v", err)
		}
	}

	if err = d.Set("domain_name_stripping", flattenWirelessControllerVapDomainNameStripping(o["domain-name-stripping"], d, "domain_name_stripping")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain-name-stripping"], "WirelessControllerVap-DomainNameStripping"); ok {
			if err = d.Set("domain_name_stripping", vv); err != nil {
				return fmt.Errorf("Error reading domain_name_stripping: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain_name_stripping: %v", err)
		}
	}

	if err = d.Set("dynamic_vlan", flattenWirelessControllerVapDynamicVlan(o["dynamic-vlan"], d, "dynamic_vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["dynamic-vlan"], "WirelessControllerVap-DynamicVlan"); ok {
			if err = d.Set("dynamic_vlan", vv); err != nil {
				return fmt.Errorf("Error reading dynamic_vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dynamic_vlan: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenWirelessControllerVapDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "WirelessControllerVap-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenWirelessControllerVapDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "WirelessControllerVap-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("eap_reauth", flattenWirelessControllerVapEapReauth(o["eap-reauth"], d, "eap_reauth")); err != nil {
		if vv, ok := fortiAPIPatch(o["eap-reauth"], "WirelessControllerVap-EapReauth"); ok {
			if err = d.Set("eap_reauth", vv); err != nil {
				return fmt.Errorf("Error reading eap_reauth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eap_reauth: %v", err)
		}
	}

	if err = d.Set("eap_reauth_intv", flattenWirelessControllerVapEapReauthIntv(o["eap-reauth-intv"], d, "eap_reauth_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["eap-reauth-intv"], "WirelessControllerVap-EapReauthIntv"); ok {
			if err = d.Set("eap_reauth_intv", vv); err != nil {
				return fmt.Errorf("Error reading eap_reauth_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eap_reauth_intv: %v", err)
		}
	}

	if err = d.Set("eapol_key_retries", flattenWirelessControllerVapEapolKeyRetries(o["eapol-key-retries"], d, "eapol_key_retries")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-key-retries"], "WirelessControllerVap-EapolKeyRetries"); ok {
			if err = d.Set("eapol_key_retries", vv); err != nil {
				return fmt.Errorf("Error reading eapol_key_retries: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_key_retries: %v", err)
		}
	}

	if err = d.Set("encrypt", flattenWirelessControllerVapEncrypt(o["encrypt"], d, "encrypt")); err != nil {
		if vv, ok := fortiAPIPatch(o["encrypt"], "WirelessControllerVap-Encrypt"); ok {
			if err = d.Set("encrypt", vv); err != nil {
				return fmt.Errorf("Error reading encrypt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading encrypt: %v", err)
		}
	}

	if err = d.Set("external_fast_roaming", flattenWirelessControllerVapExternalFastRoaming(o["external-fast-roaming"], d, "external_fast_roaming")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-fast-roaming"], "WirelessControllerVap-ExternalFastRoaming"); ok {
			if err = d.Set("external_fast_roaming", vv); err != nil {
				return fmt.Errorf("Error reading external_fast_roaming: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_fast_roaming: %v", err)
		}
	}

	if err = d.Set("external_logout", flattenWirelessControllerVapExternalLogout(o["external-logout"], d, "external_logout")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-logout"], "WirelessControllerVap-ExternalLogout"); ok {
			if err = d.Set("external_logout", vv); err != nil {
				return fmt.Errorf("Error reading external_logout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_logout: %v", err)
		}
	}

	if err = d.Set("external_pre_auth", flattenWirelessControllerVapExternalPreAuth(o["external-pre-auth"], d, "external_pre_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-pre-auth"], "WirelessControllerVap-ExternalPreAuth"); ok {
			if err = d.Set("external_pre_auth", vv); err != nil {
				return fmt.Errorf("Error reading external_pre_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_pre_auth: %v", err)
		}
	}

	if err = d.Set("external_web", flattenWirelessControllerVapExternalWeb(o["external-web"], d, "external_web")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-web"], "WirelessControllerVap-ExternalWeb"); ok {
			if err = d.Set("external_web", vv); err != nil {
				return fmt.Errorf("Error reading external_web: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_web: %v", err)
		}
	}

	if err = d.Set("external_web_format", flattenWirelessControllerVapExternalWebFormat(o["external-web-format"], d, "external_web_format")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-web-format"], "WirelessControllerVap-ExternalWebFormat"); ok {
			if err = d.Set("external_web_format", vv); err != nil {
				return fmt.Errorf("Error reading external_web_format: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_web_format: %v", err)
		}
	}

	if err = d.Set("fast_bss_transition", flattenWirelessControllerVapFastBssTransition(o["fast-bss-transition"], d, "fast_bss_transition")); err != nil {
		if vv, ok := fortiAPIPatch(o["fast-bss-transition"], "WirelessControllerVap-FastBssTransition"); ok {
			if err = d.Set("fast_bss_transition", vv); err != nil {
				return fmt.Errorf("Error reading fast_bss_transition: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fast_bss_transition: %v", err)
		}
	}

	if err = d.Set("fast_roaming", flattenWirelessControllerVapFastRoaming(o["fast-roaming"], d, "fast_roaming")); err != nil {
		if vv, ok := fortiAPIPatch(o["fast-roaming"], "WirelessControllerVap-FastRoaming"); ok {
			if err = d.Set("fast_roaming", vv); err != nil {
				return fmt.Errorf("Error reading fast_roaming: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fast_roaming: %v", err)
		}
	}

	if err = d.Set("ft_mobility_domain", flattenWirelessControllerVapFtMobilityDomain(o["ft-mobility-domain"], d, "ft_mobility_domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["ft-mobility-domain"], "WirelessControllerVap-FtMobilityDomain"); ok {
			if err = d.Set("ft_mobility_domain", vv); err != nil {
				return fmt.Errorf("Error reading ft_mobility_domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ft_mobility_domain: %v", err)
		}
	}

	if err = d.Set("ft_over_ds", flattenWirelessControllerVapFtOverDs(o["ft-over-ds"], d, "ft_over_ds")); err != nil {
		if vv, ok := fortiAPIPatch(o["ft-over-ds"], "WirelessControllerVap-FtOverDs"); ok {
			if err = d.Set("ft_over_ds", vv); err != nil {
				return fmt.Errorf("Error reading ft_over_ds: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ft_over_ds: %v", err)
		}
	}

	if err = d.Set("ft_r0_key_lifetime", flattenWirelessControllerVapFtR0KeyLifetime(o["ft-r0-key-lifetime"], d, "ft_r0_key_lifetime")); err != nil {
		if vv, ok := fortiAPIPatch(o["ft-r0-key-lifetime"], "WirelessControllerVap-FtR0KeyLifetime"); ok {
			if err = d.Set("ft_r0_key_lifetime", vv); err != nil {
				return fmt.Errorf("Error reading ft_r0_key_lifetime: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ft_r0_key_lifetime: %v", err)
		}
	}

	if err = d.Set("gas_comeback_delay", flattenWirelessControllerVapGasComebackDelay(o["gas-comeback-delay"], d, "gas_comeback_delay")); err != nil {
		if vv, ok := fortiAPIPatch(o["gas-comeback-delay"], "WirelessControllerVap-GasComebackDelay"); ok {
			if err = d.Set("gas_comeback_delay", vv); err != nil {
				return fmt.Errorf("Error reading gas_comeback_delay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gas_comeback_delay: %v", err)
		}
	}

	if err = d.Set("gas_fragmentation_limit", flattenWirelessControllerVapGasFragmentationLimit(o["gas-fragmentation-limit"], d, "gas_fragmentation_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["gas-fragmentation-limit"], "WirelessControllerVap-GasFragmentationLimit"); ok {
			if err = d.Set("gas_fragmentation_limit", vv); err != nil {
				return fmt.Errorf("Error reading gas_fragmentation_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gas_fragmentation_limit: %v", err)
		}
	}

	if err = d.Set("gtk_rekey", flattenWirelessControllerVapGtkRekey(o["gtk-rekey"], d, "gtk_rekey")); err != nil {
		if vv, ok := fortiAPIPatch(o["gtk-rekey"], "WirelessControllerVap-GtkRekey"); ok {
			if err = d.Set("gtk_rekey", vv); err != nil {
				return fmt.Errorf("Error reading gtk_rekey: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gtk_rekey: %v", err)
		}
	}

	if err = d.Set("gtk_rekey_intv", flattenWirelessControllerVapGtkRekeyIntv(o["gtk-rekey-intv"], d, "gtk_rekey_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["gtk-rekey-intv"], "WirelessControllerVap-GtkRekeyIntv"); ok {
			if err = d.Set("gtk_rekey_intv", vv); err != nil {
				return fmt.Errorf("Error reading gtk_rekey_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gtk_rekey_intv: %v", err)
		}
	}

	if err = d.Set("high_efficiency", flattenWirelessControllerVapHighEfficiency(o["high-efficiency"], d, "high_efficiency")); err != nil {
		if vv, ok := fortiAPIPatch(o["high-efficiency"], "WirelessControllerVap-HighEfficiency"); ok {
			if err = d.Set("high_efficiency", vv); err != nil {
				return fmt.Errorf("Error reading high_efficiency: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading high_efficiency: %v", err)
		}
	}

	if err = d.Set("hotspot20_profile", flattenWirelessControllerVapHotspot20Profile(o["hotspot20-profile"], d, "hotspot20_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["hotspot20-profile"], "WirelessControllerVap-Hotspot20Profile"); ok {
			if err = d.Set("hotspot20_profile", vv); err != nil {
				return fmt.Errorf("Error reading hotspot20_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hotspot20_profile: %v", err)
		}
	}

	if err = d.Set("igmp_snooping", flattenWirelessControllerVapIgmpSnooping(o["igmp-snooping"], d, "igmp_snooping")); err != nil {
		if vv, ok := fortiAPIPatch(o["igmp-snooping"], "WirelessControllerVap-IgmpSnooping"); ok {
			if err = d.Set("igmp_snooping", vv); err != nil {
				return fmt.Errorf("Error reading igmp_snooping: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading igmp_snooping: %v", err)
		}
	}

	if err = d.Set("intra_vap_privacy", flattenWirelessControllerVapIntraVapPrivacy(o["intra-vap-privacy"], d, "intra_vap_privacy")); err != nil {
		if vv, ok := fortiAPIPatch(o["intra-vap-privacy"], "WirelessControllerVap-IntraVapPrivacy"); ok {
			if err = d.Set("intra_vap_privacy", vv); err != nil {
				return fmt.Errorf("Error reading intra_vap_privacy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading intra_vap_privacy: %v", err)
		}
	}

	if err = d.Set("ip", flattenWirelessControllerVapIp(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "WirelessControllerVap-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("ips_sensor", flattenWirelessControllerVapIpsSensor(o["ips-sensor"], d, "ips_sensor")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-sensor"], "WirelessControllerVap-IpsSensor"); ok {
			if err = d.Set("ips_sensor", vv); err != nil {
				return fmt.Errorf("Error reading ips_sensor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_sensor: %v", err)
		}
	}

	if err = d.Set("ipv6_rules", flattenWirelessControllerVapIpv6Rules(o["ipv6-rules"], d, "ipv6_rules")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-rules"], "WirelessControllerVap-Ipv6Rules"); ok {
			if err = d.Set("ipv6_rules", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_rules: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_rules: %v", err)
		}
	}

	if err = d.Set("keyindex", flattenWirelessControllerVapKeyindex(o["keyindex"], d, "keyindex")); err != nil {
		if vv, ok := fortiAPIPatch(o["keyindex"], "WirelessControllerVap-Keyindex"); ok {
			if err = d.Set("keyindex", vv); err != nil {
				return fmt.Errorf("Error reading keyindex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keyindex: %v", err)
		}
	}

	if err = d.Set("l3_roaming", flattenWirelessControllerVapL3Roaming(o["l3-roaming"], d, "l3_roaming")); err != nil {
		if vv, ok := fortiAPIPatch(o["l3-roaming"], "WirelessControllerVap-L3Roaming"); ok {
			if err = d.Set("l3_roaming", vv); err != nil {
				return fmt.Errorf("Error reading l3_roaming: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading l3_roaming: %v", err)
		}
	}

	if err = d.Set("l3_roaming_mode", flattenWirelessControllerVapL3RoamingMode(o["l3-roaming-mode"], d, "l3_roaming_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["l3-roaming-mode"], "WirelessControllerVap-L3RoamingMode"); ok {
			if err = d.Set("l3_roaming_mode", vv); err != nil {
				return fmt.Errorf("Error reading l3_roaming_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading l3_roaming_mode: %v", err)
		}
	}

	if err = d.Set("ldpc", flattenWirelessControllerVapLdpc(o["ldpc"], d, "ldpc")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldpc"], "WirelessControllerVap-Ldpc"); ok {
			if err = d.Set("ldpc", vv); err != nil {
				return fmt.Errorf("Error reading ldpc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldpc: %v", err)
		}
	}

	if err = d.Set("local_authentication", flattenWirelessControllerVapLocalAuthentication(o["local-authentication"], d, "local_authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-authentication"], "WirelessControllerVap-LocalAuthentication"); ok {
			if err = d.Set("local_authentication", vv); err != nil {
				return fmt.Errorf("Error reading local_authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_authentication: %v", err)
		}
	}

	if err = d.Set("local_bridging", flattenWirelessControllerVapLocalBridging(o["local-bridging"], d, "local_bridging")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-bridging"], "WirelessControllerVap-LocalBridging"); ok {
			if err = d.Set("local_bridging", vv); err != nil {
				return fmt.Errorf("Error reading local_bridging: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_bridging: %v", err)
		}
	}

	if err = d.Set("local_lan", flattenWirelessControllerVapLocalLan(o["local-lan"], d, "local_lan")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-lan"], "WirelessControllerVap-LocalLan"); ok {
			if err = d.Set("local_lan", vv); err != nil {
				return fmt.Errorf("Error reading local_lan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_lan: %v", err)
		}
	}

	if err = d.Set("local_lan_partition", flattenWirelessControllerVapLocalLanPartition(o["local-lan-partition"], d, "local_lan_partition")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-lan-partition"], "WirelessControllerVap-LocalLanPartition"); ok {
			if err = d.Set("local_lan_partition", vv); err != nil {
				return fmt.Errorf("Error reading local_lan_partition: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_lan_partition: %v", err)
		}
	}

	if err = d.Set("local_standalone", flattenWirelessControllerVapLocalStandalone(o["local-standalone"], d, "local_standalone")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-standalone"], "WirelessControllerVap-LocalStandalone"); ok {
			if err = d.Set("local_standalone", vv); err != nil {
				return fmt.Errorf("Error reading local_standalone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_standalone: %v", err)
		}
	}

	if err = d.Set("local_standalone_dns", flattenWirelessControllerVapLocalStandaloneDns(o["local-standalone-dns"], d, "local_standalone_dns")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-standalone-dns"], "WirelessControllerVap-LocalStandaloneDns"); ok {
			if err = d.Set("local_standalone_dns", vv); err != nil {
				return fmt.Errorf("Error reading local_standalone_dns: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_standalone_dns: %v", err)
		}
	}

	if err = d.Set("local_standalone_dns_ip", flattenWirelessControllerVapLocalStandaloneDnsIp(o["local-standalone-dns-ip"], d, "local_standalone_dns_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-standalone-dns-ip"], "WirelessControllerVap-LocalStandaloneDnsIp"); ok {
			if err = d.Set("local_standalone_dns_ip", vv); err != nil {
				return fmt.Errorf("Error reading local_standalone_dns_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_standalone_dns_ip: %v", err)
		}
	}

	if err = d.Set("local_standalone_nat", flattenWirelessControllerVapLocalStandaloneNat(o["local-standalone-nat"], d, "local_standalone_nat")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-standalone-nat"], "WirelessControllerVap-LocalStandaloneNat"); ok {
			if err = d.Set("local_standalone_nat", vv); err != nil {
				return fmt.Errorf("Error reading local_standalone_nat: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_standalone_nat: %v", err)
		}
	}

	if err = d.Set("mac_auth_bypass", flattenWirelessControllerVapMacAuthBypass(o["mac-auth-bypass"], d, "mac_auth_bypass")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-auth-bypass"], "WirelessControllerVap-MacAuthBypass"); ok {
			if err = d.Set("mac_auth_bypass", vv); err != nil {
				return fmt.Errorf("Error reading mac_auth_bypass: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_auth_bypass: %v", err)
		}
	}

	if err = d.Set("mac_called_station_delimiter", flattenWirelessControllerVapMacCalledStationDelimiter(o["mac-called-station-delimiter"], d, "mac_called_station_delimiter")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-called-station-delimiter"], "WirelessControllerVap-MacCalledStationDelimiter"); ok {
			if err = d.Set("mac_called_station_delimiter", vv); err != nil {
				return fmt.Errorf("Error reading mac_called_station_delimiter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_called_station_delimiter: %v", err)
		}
	}

	if err = d.Set("mac_calling_station_delimiter", flattenWirelessControllerVapMacCallingStationDelimiter(o["mac-calling-station-delimiter"], d, "mac_calling_station_delimiter")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-calling-station-delimiter"], "WirelessControllerVap-MacCallingStationDelimiter"); ok {
			if err = d.Set("mac_calling_station_delimiter", vv); err != nil {
				return fmt.Errorf("Error reading mac_calling_station_delimiter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_calling_station_delimiter: %v", err)
		}
	}

	if err = d.Set("mac_case", flattenWirelessControllerVapMacCase(o["mac-case"], d, "mac_case")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-case"], "WirelessControllerVap-MacCase"); ok {
			if err = d.Set("mac_case", vv); err != nil {
				return fmt.Errorf("Error reading mac_case: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_case: %v", err)
		}
	}

	if err = d.Set("mac_filter", flattenWirelessControllerVapMacFilter(o["mac-filter"], d, "mac_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-filter"], "WirelessControllerVap-MacFilter"); ok {
			if err = d.Set("mac_filter", vv); err != nil {
				return fmt.Errorf("Error reading mac_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_filter: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("mac_filter_list", flattenWirelessControllerVapMacFilterList(o["mac-filter-list"], d, "mac_filter_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["mac-filter-list"], "WirelessControllerVap-MacFilterList"); ok {
				if err = d.Set("mac_filter_list", vv); err != nil {
					return fmt.Errorf("Error reading mac_filter_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mac_filter_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mac_filter_list"); ok {
			if err = d.Set("mac_filter_list", flattenWirelessControllerVapMacFilterList(o["mac-filter-list"], d, "mac_filter_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["mac-filter-list"], "WirelessControllerVap-MacFilterList"); ok {
					if err = d.Set("mac_filter_list", vv); err != nil {
						return fmt.Errorf("Error reading mac_filter_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mac_filter_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("mac_filter_policy_other", flattenWirelessControllerVapMacFilterPolicyOther(o["mac-filter-policy-other"], d, "mac_filter_policy_other")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-filter-policy-other"], "WirelessControllerVap-MacFilterPolicyOther"); ok {
			if err = d.Set("mac_filter_policy_other", vv); err != nil {
				return fmt.Errorf("Error reading mac_filter_policy_other: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_filter_policy_other: %v", err)
		}
	}

	if err = d.Set("mac_password_delimiter", flattenWirelessControllerVapMacPasswordDelimiter(o["mac-password-delimiter"], d, "mac_password_delimiter")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-password-delimiter"], "WirelessControllerVap-MacPasswordDelimiter"); ok {
			if err = d.Set("mac_password_delimiter", vv); err != nil {
				return fmt.Errorf("Error reading mac_password_delimiter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_password_delimiter: %v", err)
		}
	}

	if err = d.Set("mac_username_delimiter", flattenWirelessControllerVapMacUsernameDelimiter(o["mac-username-delimiter"], d, "mac_username_delimiter")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-username-delimiter"], "WirelessControllerVap-MacUsernameDelimiter"); ok {
			if err = d.Set("mac_username_delimiter", vv); err != nil {
				return fmt.Errorf("Error reading mac_username_delimiter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_username_delimiter: %v", err)
		}
	}

	if err = d.Set("max_clients", flattenWirelessControllerVapMaxClients(o["max-clients"], d, "max_clients")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-clients"], "WirelessControllerVap-MaxClients"); ok {
			if err = d.Set("max_clients", vv); err != nil {
				return fmt.Errorf("Error reading max_clients: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_clients: %v", err)
		}
	}

	if err = d.Set("max_clients_ap", flattenWirelessControllerVapMaxClientsAp(o["max-clients-ap"], d, "max_clients_ap")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-clients-ap"], "WirelessControllerVap-MaxClientsAp"); ok {
			if err = d.Set("max_clients_ap", vv); err != nil {
				return fmt.Errorf("Error reading max_clients_ap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_clients_ap: %v", err)
		}
	}

	if err = d.Set("mbo", flattenWirelessControllerVapMbo(o["mbo"], d, "mbo")); err != nil {
		if vv, ok := fortiAPIPatch(o["mbo"], "WirelessControllerVap-Mbo"); ok {
			if err = d.Set("mbo", vv); err != nil {
				return fmt.Errorf("Error reading mbo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mbo: %v", err)
		}
	}

	if err = d.Set("mbo_cell_data_conn_pref", flattenWirelessControllerVapMboCellDataConnPref(o["mbo-cell-data-conn-pref"], d, "mbo_cell_data_conn_pref")); err != nil {
		if vv, ok := fortiAPIPatch(o["mbo-cell-data-conn-pref"], "WirelessControllerVap-MboCellDataConnPref"); ok {
			if err = d.Set("mbo_cell_data_conn_pref", vv); err != nil {
				return fmt.Errorf("Error reading mbo_cell_data_conn_pref: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mbo_cell_data_conn_pref: %v", err)
		}
	}

	if err = d.Set("me_disable_thresh", flattenWirelessControllerVapMeDisableThresh(o["me-disable-thresh"], d, "me_disable_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["me-disable-thresh"], "WirelessControllerVap-MeDisableThresh"); ok {
			if err = d.Set("me_disable_thresh", vv); err != nil {
				return fmt.Errorf("Error reading me_disable_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading me_disable_thresh: %v", err)
		}
	}

	if err = d.Set("mesh_backhaul", flattenWirelessControllerVapMeshBackhaul(o["mesh-backhaul"], d, "mesh_backhaul")); err != nil {
		if vv, ok := fortiAPIPatch(o["mesh-backhaul"], "WirelessControllerVap-MeshBackhaul"); ok {
			if err = d.Set("mesh_backhaul", vv); err != nil {
				return fmt.Errorf("Error reading mesh_backhaul: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mesh_backhaul: %v", err)
		}
	}

	if err = d.Set("mpsk", flattenWirelessControllerVapMpsk(o["mpsk"], d, "mpsk")); err != nil {
		if vv, ok := fortiAPIPatch(o["mpsk"], "WirelessControllerVap-Mpsk"); ok {
			if err = d.Set("mpsk", vv); err != nil {
				return fmt.Errorf("Error reading mpsk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mpsk: %v", err)
		}
	}

	if err = d.Set("mpsk_concurrent_clients", flattenWirelessControllerVapMpskConcurrentClients(o["mpsk-concurrent-clients"], d, "mpsk_concurrent_clients")); err != nil {
		if vv, ok := fortiAPIPatch(o["mpsk-concurrent-clients"], "WirelessControllerVap-MpskConcurrentClients"); ok {
			if err = d.Set("mpsk_concurrent_clients", vv); err != nil {
				return fmt.Errorf("Error reading mpsk_concurrent_clients: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mpsk_concurrent_clients: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("mpsk_key", flattenWirelessControllerVapMpskKey(o["mpsk-key"], d, "mpsk_key")); err != nil {
			if vv, ok := fortiAPIPatch(o["mpsk-key"], "WirelessControllerVap-MpskKey"); ok {
				if err = d.Set("mpsk_key", vv); err != nil {
					return fmt.Errorf("Error reading mpsk_key: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mpsk_key: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mpsk_key"); ok {
			if err = d.Set("mpsk_key", flattenWirelessControllerVapMpskKey(o["mpsk-key"], d, "mpsk_key")); err != nil {
				if vv, ok := fortiAPIPatch(o["mpsk-key"], "WirelessControllerVap-MpskKey"); ok {
					if err = d.Set("mpsk_key", vv); err != nil {
						return fmt.Errorf("Error reading mpsk_key: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mpsk_key: %v", err)
				}
			}
		}
	}

	if err = d.Set("mpsk_profile", flattenWirelessControllerVapMpskProfile(o["mpsk-profile"], d, "mpsk_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["mpsk-profile"], "WirelessControllerVap-MpskProfile"); ok {
			if err = d.Set("mpsk_profile", vv); err != nil {
				return fmt.Errorf("Error reading mpsk_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mpsk_profile: %v", err)
		}
	}

	if err = d.Set("mu_mimo", flattenWirelessControllerVapMuMimo(o["mu-mimo"], d, "mu_mimo")); err != nil {
		if vv, ok := fortiAPIPatch(o["mu-mimo"], "WirelessControllerVap-MuMimo"); ok {
			if err = d.Set("mu_mimo", vv); err != nil {
				return fmt.Errorf("Error reading mu_mimo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mu_mimo: %v", err)
		}
	}

	if err = d.Set("multicast_enhance", flattenWirelessControllerVapMulticastEnhance(o["multicast-enhance"], d, "multicast_enhance")); err != nil {
		if vv, ok := fortiAPIPatch(o["multicast-enhance"], "WirelessControllerVap-MulticastEnhance"); ok {
			if err = d.Set("multicast_enhance", vv); err != nil {
				return fmt.Errorf("Error reading multicast_enhance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multicast_enhance: %v", err)
		}
	}

	if err = d.Set("multicast_rate", flattenWirelessControllerVapMulticastRate(o["multicast-rate"], d, "multicast_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["multicast-rate"], "WirelessControllerVap-MulticastRate"); ok {
			if err = d.Set("multicast_rate", vv); err != nil {
				return fmt.Errorf("Error reading multicast_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multicast_rate: %v", err)
		}
	}

	if err = d.Set("nac", flattenWirelessControllerVapNac(o["nac"], d, "nac")); err != nil {
		if vv, ok := fortiAPIPatch(o["nac"], "WirelessControllerVap-Nac"); ok {
			if err = d.Set("nac", vv); err != nil {
				return fmt.Errorf("Error reading nac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nac: %v", err)
		}
	}

	if err = d.Set("nac_profile", flattenWirelessControllerVapNacProfile(o["nac-profile"], d, "nac_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["nac-profile"], "WirelessControllerVap-NacProfile"); ok {
			if err = d.Set("nac_profile", vv); err != nil {
				return fmt.Errorf("Error reading nac_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nac_profile: %v", err)
		}
	}

	if err = d.Set("name", flattenWirelessControllerVapName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerVap-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("nas_filter_rule", flattenWirelessControllerVapNasFilterRule(o["nas-filter-rule"], d, "nas_filter_rule")); err != nil {
		if vv, ok := fortiAPIPatch(o["nas-filter-rule"], "WirelessControllerVap-NasFilterRule"); ok {
			if err = d.Set("nas_filter_rule", vv); err != nil {
				return fmt.Errorf("Error reading nas_filter_rule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nas_filter_rule: %v", err)
		}
	}

	if err = d.Set("neighbor_report_dual_band", flattenWirelessControllerVapNeighborReportDualBand(o["neighbor-report-dual-band"], d, "neighbor_report_dual_band")); err != nil {
		if vv, ok := fortiAPIPatch(o["neighbor-report-dual-band"], "WirelessControllerVap-NeighborReportDualBand"); ok {
			if err = d.Set("neighbor_report_dual_band", vv); err != nil {
				return fmt.Errorf("Error reading neighbor_report_dual_band: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading neighbor_report_dual_band: %v", err)
		}
	}

	if err = d.Set("okc", flattenWirelessControllerVapOkc(o["okc"], d, "okc")); err != nil {
		if vv, ok := fortiAPIPatch(o["okc"], "WirelessControllerVap-Okc"); ok {
			if err = d.Set("okc", vv); err != nil {
				return fmt.Errorf("Error reading okc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading okc: %v", err)
		}
	}

	if err = d.Set("osen", flattenWirelessControllerVapOsen(o["osen"], d, "osen")); err != nil {
		if vv, ok := fortiAPIPatch(o["osen"], "WirelessControllerVap-Osen"); ok {
			if err = d.Set("osen", vv); err != nil {
				return fmt.Errorf("Error reading osen: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading osen: %v", err)
		}
	}

	if err = d.Set("owe_groups", flattenWirelessControllerVapOweGroups(o["owe-groups"], d, "owe_groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["owe-groups"], "WirelessControllerVap-OweGroups"); ok {
			if err = d.Set("owe_groups", vv); err != nil {
				return fmt.Errorf("Error reading owe_groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading owe_groups: %v", err)
		}
	}

	if err = d.Set("owe_transition", flattenWirelessControllerVapOweTransition(o["owe-transition"], d, "owe_transition")); err != nil {
		if vv, ok := fortiAPIPatch(o["owe-transition"], "WirelessControllerVap-OweTransition"); ok {
			if err = d.Set("owe_transition", vv); err != nil {
				return fmt.Errorf("Error reading owe_transition: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading owe_transition: %v", err)
		}
	}

	if err = d.Set("owe_transition_ssid", flattenWirelessControllerVapOweTransitionSsid(o["owe-transition-ssid"], d, "owe_transition_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["owe-transition-ssid"], "WirelessControllerVap-OweTransitionSsid"); ok {
			if err = d.Set("owe_transition_ssid", vv); err != nil {
				return fmt.Errorf("Error reading owe_transition_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading owe_transition_ssid: %v", err)
		}
	}

	if err = d.Set("pmf", flattenWirelessControllerVapPmf(o["pmf"], d, "pmf")); err != nil {
		if vv, ok := fortiAPIPatch(o["pmf"], "WirelessControllerVap-Pmf"); ok {
			if err = d.Set("pmf", vv); err != nil {
				return fmt.Errorf("Error reading pmf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pmf: %v", err)
		}
	}

	if err = d.Set("pmf_assoc_comeback_timeout", flattenWirelessControllerVapPmfAssocComebackTimeout(o["pmf-assoc-comeback-timeout"], d, "pmf_assoc_comeback_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["pmf-assoc-comeback-timeout"], "WirelessControllerVap-PmfAssocComebackTimeout"); ok {
			if err = d.Set("pmf_assoc_comeback_timeout", vv); err != nil {
				return fmt.Errorf("Error reading pmf_assoc_comeback_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pmf_assoc_comeback_timeout: %v", err)
		}
	}

	if err = d.Set("pmf_sa_query_retry_timeout", flattenWirelessControllerVapPmfSaQueryRetryTimeout(o["pmf-sa-query-retry-timeout"], d, "pmf_sa_query_retry_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["pmf-sa-query-retry-timeout"], "WirelessControllerVap-PmfSaQueryRetryTimeout"); ok {
			if err = d.Set("pmf_sa_query_retry_timeout", vv); err != nil {
				return fmt.Errorf("Error reading pmf_sa_query_retry_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pmf_sa_query_retry_timeout: %v", err)
		}
	}

	if err = d.Set("port_macauth", flattenWirelessControllerVapPortMacauth(o["port-macauth"], d, "port_macauth")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-macauth"], "WirelessControllerVap-PortMacauth"); ok {
			if err = d.Set("port_macauth", vv); err != nil {
				return fmt.Errorf("Error reading port_macauth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_macauth: %v", err)
		}
	}

	if err = d.Set("port_macauth_reauth_timeout", flattenWirelessControllerVapPortMacauthReauthTimeout(o["port-macauth-reauth-timeout"], d, "port_macauth_reauth_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-macauth-reauth-timeout"], "WirelessControllerVap-PortMacauthReauthTimeout"); ok {
			if err = d.Set("port_macauth_reauth_timeout", vv); err != nil {
				return fmt.Errorf("Error reading port_macauth_reauth_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_macauth_reauth_timeout: %v", err)
		}
	}

	if err = d.Set("port_macauth_timeout", flattenWirelessControllerVapPortMacauthTimeout(o["port-macauth-timeout"], d, "port_macauth_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-macauth-timeout"], "WirelessControllerVap-PortMacauthTimeout"); ok {
			if err = d.Set("port_macauth_timeout", vv); err != nil {
				return fmt.Errorf("Error reading port_macauth_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_macauth_timeout: %v", err)
		}
	}

	if err = d.Set("portal_message_override_group", flattenWirelessControllerVapPortalMessageOverrideGroup(o["portal-message-override-group"], d, "portal_message_override_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["portal-message-override-group"], "WirelessControllerVap-PortalMessageOverrideGroup"); ok {
			if err = d.Set("portal_message_override_group", vv); err != nil {
				return fmt.Errorf("Error reading portal_message_override_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading portal_message_override_group: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("portal_message_overrides", flattenWirelessControllerVapPortalMessageOverrides(o["portal-message-overrides"], d, "portal_message_overrides")); err != nil {
			if vv, ok := fortiAPIPatch(o["portal-message-overrides"], "WirelessControllerVap-PortalMessageOverrides"); ok {
				if err = d.Set("portal_message_overrides", vv); err != nil {
					return fmt.Errorf("Error reading portal_message_overrides: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading portal_message_overrides: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("portal_message_overrides"); ok {
			if err = d.Set("portal_message_overrides", flattenWirelessControllerVapPortalMessageOverrides(o["portal-message-overrides"], d, "portal_message_overrides")); err != nil {
				if vv, ok := fortiAPIPatch(o["portal-message-overrides"], "WirelessControllerVap-PortalMessageOverrides"); ok {
					if err = d.Set("portal_message_overrides", vv); err != nil {
						return fmt.Errorf("Error reading portal_message_overrides: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading portal_message_overrides: %v", err)
				}
			}
		}
	}

	if err = d.Set("portal_type", flattenWirelessControllerVapPortalType(o["portal-type"], d, "portal_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["portal-type"], "WirelessControllerVap-PortalType"); ok {
			if err = d.Set("portal_type", vv); err != nil {
				return fmt.Errorf("Error reading portal_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading portal_type: %v", err)
		}
	}

	if err = d.Set("pre_auth", flattenWirelessControllerVapPreAuth(o["pre-auth"], d, "pre_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["pre-auth"], "WirelessControllerVap-PreAuth"); ok {
			if err = d.Set("pre_auth", vv); err != nil {
				return fmt.Errorf("Error reading pre_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pre_auth: %v", err)
		}
	}

	if err = d.Set("primary_wag_profile", flattenWirelessControllerVapPrimaryWagProfile(o["primary-wag-profile"], d, "primary_wag_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["primary-wag-profile"], "WirelessControllerVap-PrimaryWagProfile"); ok {
			if err = d.Set("primary_wag_profile", vv); err != nil {
				return fmt.Errorf("Error reading primary_wag_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading primary_wag_profile: %v", err)
		}
	}

	if err = d.Set("probe_resp_suppression", flattenWirelessControllerVapProbeRespSuppression(o["probe-resp-suppression"], d, "probe_resp_suppression")); err != nil {
		if vv, ok := fortiAPIPatch(o["probe-resp-suppression"], "WirelessControllerVap-ProbeRespSuppression"); ok {
			if err = d.Set("probe_resp_suppression", vv); err != nil {
				return fmt.Errorf("Error reading probe_resp_suppression: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading probe_resp_suppression: %v", err)
		}
	}

	if err = d.Set("probe_resp_threshold", flattenWirelessControllerVapProbeRespThreshold(o["probe-resp-threshold"], d, "probe_resp_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["probe-resp-threshold"], "WirelessControllerVap-ProbeRespThreshold"); ok {
			if err = d.Set("probe_resp_threshold", vv); err != nil {
				return fmt.Errorf("Error reading probe_resp_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading probe_resp_threshold: %v", err)
		}
	}

	if err = d.Set("ptk_rekey", flattenWirelessControllerVapPtkRekey(o["ptk-rekey"], d, "ptk_rekey")); err != nil {
		if vv, ok := fortiAPIPatch(o["ptk-rekey"], "WirelessControllerVap-PtkRekey"); ok {
			if err = d.Set("ptk_rekey", vv); err != nil {
				return fmt.Errorf("Error reading ptk_rekey: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ptk_rekey: %v", err)
		}
	}

	if err = d.Set("ptk_rekey_intv", flattenWirelessControllerVapPtkRekeyIntv(o["ptk-rekey-intv"], d, "ptk_rekey_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["ptk-rekey-intv"], "WirelessControllerVap-PtkRekeyIntv"); ok {
			if err = d.Set("ptk_rekey_intv", vv); err != nil {
				return fmt.Errorf("Error reading ptk_rekey_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ptk_rekey_intv: %v", err)
		}
	}

	if err = d.Set("qos_profile", flattenWirelessControllerVapQosProfile(o["qos-profile"], d, "qos_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["qos-profile"], "WirelessControllerVap-QosProfile"); ok {
			if err = d.Set("qos_profile", vv); err != nil {
				return fmt.Errorf("Error reading qos_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading qos_profile: %v", err)
		}
	}

	if err = d.Set("quarantine", flattenWirelessControllerVapQuarantine(o["quarantine"], d, "quarantine")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine"], "WirelessControllerVap-Quarantine"); ok {
			if err = d.Set("quarantine", vv); err != nil {
				return fmt.Errorf("Error reading quarantine: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine: %v", err)
		}
	}

	if err = d.Set("radio_2g_threshold", flattenWirelessControllerVapRadio2GThreshold(o["radio-2g-threshold"], d, "radio_2g_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["radio-2g-threshold"], "WirelessControllerVap-Radio2GThreshold"); ok {
			if err = d.Set("radio_2g_threshold", vv); err != nil {
				return fmt.Errorf("Error reading radio_2g_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radio_2g_threshold: %v", err)
		}
	}

	if err = d.Set("radio_5g_threshold", flattenWirelessControllerVapRadio5GThreshold(o["radio-5g-threshold"], d, "radio_5g_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["radio-5g-threshold"], "WirelessControllerVap-Radio5GThreshold"); ok {
			if err = d.Set("radio_5g_threshold", vv); err != nil {
				return fmt.Errorf("Error reading radio_5g_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radio_5g_threshold: %v", err)
		}
	}

	if err = d.Set("radio_sensitivity", flattenWirelessControllerVapRadioSensitivity(o["radio-sensitivity"], d, "radio_sensitivity")); err != nil {
		if vv, ok := fortiAPIPatch(o["radio-sensitivity"], "WirelessControllerVap-RadioSensitivity"); ok {
			if err = d.Set("radio_sensitivity", vv); err != nil {
				return fmt.Errorf("Error reading radio_sensitivity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radio_sensitivity: %v", err)
		}
	}

	if err = d.Set("radius_mac_auth", flattenWirelessControllerVapRadiusMacAuth(o["radius-mac-auth"], d, "radius_mac_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-mac-auth"], "WirelessControllerVap-RadiusMacAuth"); ok {
			if err = d.Set("radius_mac_auth", vv); err != nil {
				return fmt.Errorf("Error reading radius_mac_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_mac_auth: %v", err)
		}
	}

	if err = d.Set("radius_mac_auth_block_interval", flattenWirelessControllerVapRadiusMacAuthBlockInterval(o["radius-mac-auth-block-interval"], d, "radius_mac_auth_block_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-mac-auth-block-interval"], "WirelessControllerVap-RadiusMacAuthBlockInterval"); ok {
			if err = d.Set("radius_mac_auth_block_interval", vv); err != nil {
				return fmt.Errorf("Error reading radius_mac_auth_block_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_mac_auth_block_interval: %v", err)
		}
	}

	if err = d.Set("radius_mac_auth_server", flattenWirelessControllerVapRadiusMacAuthServer(o["radius-mac-auth-server"], d, "radius_mac_auth_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-mac-auth-server"], "WirelessControllerVap-RadiusMacAuthServer"); ok {
			if err = d.Set("radius_mac_auth_server", vv); err != nil {
				return fmt.Errorf("Error reading radius_mac_auth_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_mac_auth_server: %v", err)
		}
	}

	if err = d.Set("radius_mac_auth_usergroups", flattenWirelessControllerVapRadiusMacAuthUsergroups(o["radius-mac-auth-usergroups"], d, "radius_mac_auth_usergroups")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-mac-auth-usergroups"], "WirelessControllerVap-RadiusMacAuthUsergroups"); ok {
			if err = d.Set("radius_mac_auth_usergroups", vv); err != nil {
				return fmt.Errorf("Error reading radius_mac_auth_usergroups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_mac_auth_usergroups: %v", err)
		}
	}

	if err = d.Set("radius_mac_mpsk_auth", flattenWirelessControllerVapRadiusMacMpskAuth(o["radius-mac-mpsk-auth"], d, "radius_mac_mpsk_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-mac-mpsk-auth"], "WirelessControllerVap-RadiusMacMpskAuth"); ok {
			if err = d.Set("radius_mac_mpsk_auth", vv); err != nil {
				return fmt.Errorf("Error reading radius_mac_mpsk_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_mac_mpsk_auth: %v", err)
		}
	}

	if err = d.Set("radius_mac_mpsk_timeout", flattenWirelessControllerVapRadiusMacMpskTimeout(o["radius-mac-mpsk-timeout"], d, "radius_mac_mpsk_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-mac-mpsk-timeout"], "WirelessControllerVap-RadiusMacMpskTimeout"); ok {
			if err = d.Set("radius_mac_mpsk_timeout", vv); err != nil {
				return fmt.Errorf("Error reading radius_mac_mpsk_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_mac_mpsk_timeout: %v", err)
		}
	}

	if err = d.Set("radius_server", flattenWirelessControllerVapRadiusServer(o["radius-server"], d, "radius_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-server"], "WirelessControllerVap-RadiusServer"); ok {
			if err = d.Set("radius_server", vv); err != nil {
				return fmt.Errorf("Error reading radius_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_server: %v", err)
		}
	}

	if err = d.Set("rates_11a", flattenWirelessControllerVapRates11A(o["rates-11a"], d, "rates_11a")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11a"], "WirelessControllerVap-Rates11A"); ok {
			if err = d.Set("rates_11a", vv); err != nil {
				return fmt.Errorf("Error reading rates_11a: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11a: %v", err)
		}
	}

	if err = d.Set("rates_11ac_mcs_map", flattenWirelessControllerVapRates11AcMcsMap(o["rates-11ac-mcs-map"], d, "rates_11ac_mcs_map")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11ac-mcs-map"], "WirelessControllerVap-Rates11AcMcsMap"); ok {
			if err = d.Set("rates_11ac_mcs_map", vv); err != nil {
				return fmt.Errorf("Error reading rates_11ac_mcs_map: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11ac_mcs_map: %v", err)
		}
	}

	if err = d.Set("rates_11ac_ss12", flattenWirelessControllerVapRates11AcSs12(o["rates-11ac-ss12"], d, "rates_11ac_ss12")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11ac-ss12"], "WirelessControllerVap-Rates11AcSs12"); ok {
			if err = d.Set("rates_11ac_ss12", vv); err != nil {
				return fmt.Errorf("Error reading rates_11ac_ss12: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11ac_ss12: %v", err)
		}
	}

	if err = d.Set("rates_11ac_ss34", flattenWirelessControllerVapRates11AcSs34(o["rates-11ac-ss34"], d, "rates_11ac_ss34")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11ac-ss34"], "WirelessControllerVap-Rates11AcSs34"); ok {
			if err = d.Set("rates_11ac_ss34", vv); err != nil {
				return fmt.Errorf("Error reading rates_11ac_ss34: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11ac_ss34: %v", err)
		}
	}

	if err = d.Set("rates_11ax_mcs_map", flattenWirelessControllerVapRates11AxMcsMap(o["rates-11ax-mcs-map"], d, "rates_11ax_mcs_map")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11ax-mcs-map"], "WirelessControllerVap-Rates11AxMcsMap"); ok {
			if err = d.Set("rates_11ax_mcs_map", vv); err != nil {
				return fmt.Errorf("Error reading rates_11ax_mcs_map: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11ax_mcs_map: %v", err)
		}
	}

	if err = d.Set("rates_11ax_ss12", flattenWirelessControllerVapRates11AxSs12(o["rates-11ax-ss12"], d, "rates_11ax_ss12")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11ax-ss12"], "WirelessControllerVap-Rates11AxSs12"); ok {
			if err = d.Set("rates_11ax_ss12", vv); err != nil {
				return fmt.Errorf("Error reading rates_11ax_ss12: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11ax_ss12: %v", err)
		}
	}

	if err = d.Set("rates_11ax_ss34", flattenWirelessControllerVapRates11AxSs34(o["rates-11ax-ss34"], d, "rates_11ax_ss34")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11ax-ss34"], "WirelessControllerVap-Rates11AxSs34"); ok {
			if err = d.Set("rates_11ax_ss34", vv); err != nil {
				return fmt.Errorf("Error reading rates_11ax_ss34: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11ax_ss34: %v", err)
		}
	}

	if err = d.Set("rates_11be_mcs_map", flattenWirelessControllerVapRates11BeMcsMap(o["rates-11be-mcs-map"], d, "rates_11be_mcs_map")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11be-mcs-map"], "WirelessControllerVap-Rates11BeMcsMap"); ok {
			if err = d.Set("rates_11be_mcs_map", vv); err != nil {
				return fmt.Errorf("Error reading rates_11be_mcs_map: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11be_mcs_map: %v", err)
		}
	}

	if err = d.Set("rates_11be_mcs_map_160", flattenWirelessControllerVapRates11BeMcsMap160(o["rates-11be-mcs-map-160"], d, "rates_11be_mcs_map_160")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11be-mcs-map-160"], "WirelessControllerVap-Rates11BeMcsMap160"); ok {
			if err = d.Set("rates_11be_mcs_map_160", vv); err != nil {
				return fmt.Errorf("Error reading rates_11be_mcs_map_160: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11be_mcs_map_160: %v", err)
		}
	}

	if err = d.Set("rates_11be_mcs_map_320", flattenWirelessControllerVapRates11BeMcsMap320(o["rates-11be-mcs-map-320"], d, "rates_11be_mcs_map_320")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11be-mcs-map-320"], "WirelessControllerVap-Rates11BeMcsMap320"); ok {
			if err = d.Set("rates_11be_mcs_map_320", vv); err != nil {
				return fmt.Errorf("Error reading rates_11be_mcs_map_320: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11be_mcs_map_320: %v", err)
		}
	}

	if err = d.Set("rates_11bg", flattenWirelessControllerVapRates11Bg(o["rates-11bg"], d, "rates_11bg")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11bg"], "WirelessControllerVap-Rates11Bg"); ok {
			if err = d.Set("rates_11bg", vv); err != nil {
				return fmt.Errorf("Error reading rates_11bg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11bg: %v", err)
		}
	}

	if err = d.Set("rates_11n_ss12", flattenWirelessControllerVapRates11NSs12(o["rates-11n-ss12"], d, "rates_11n_ss12")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11n-ss12"], "WirelessControllerVap-Rates11NSs12"); ok {
			if err = d.Set("rates_11n_ss12", vv); err != nil {
				return fmt.Errorf("Error reading rates_11n_ss12: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11n_ss12: %v", err)
		}
	}

	if err = d.Set("rates_11n_ss34", flattenWirelessControllerVapRates11NSs34(o["rates-11n-ss34"], d, "rates_11n_ss34")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11n-ss34"], "WirelessControllerVap-Rates11NSs34"); ok {
			if err = d.Set("rates_11n_ss34", vv); err != nil {
				return fmt.Errorf("Error reading rates_11n_ss34: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11n_ss34: %v", err)
		}
	}

	if err = d.Set("roaming_acct_interim_update", flattenWirelessControllerVapRoamingAcctInterimUpdate(o["roaming-acct-interim-update"], d, "roaming_acct_interim_update")); err != nil {
		if vv, ok := fortiAPIPatch(o["roaming-acct-interim-update"], "WirelessControllerVap-RoamingAcctInterimUpdate"); ok {
			if err = d.Set("roaming_acct_interim_update", vv); err != nil {
				return fmt.Errorf("Error reading roaming_acct_interim_update: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading roaming_acct_interim_update: %v", err)
		}
	}

	if err = d.Set("sae_groups", flattenWirelessControllerVapSaeGroups(o["sae-groups"], d, "sae_groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["sae-groups"], "WirelessControllerVap-SaeGroups"); ok {
			if err = d.Set("sae_groups", vv); err != nil {
				return fmt.Errorf("Error reading sae_groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sae_groups: %v", err)
		}
	}

	if err = d.Set("sae_h2e_only", flattenWirelessControllerVapSaeH2EOnly(o["sae-h2e-only"], d, "sae_h2e_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["sae-h2e-only"], "WirelessControllerVap-SaeH2EOnly"); ok {
			if err = d.Set("sae_h2e_only", vv); err != nil {
				return fmt.Errorf("Error reading sae_h2e_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sae_h2e_only: %v", err)
		}
	}

	if err = d.Set("sae_hnp_only", flattenWirelessControllerVapSaeHnpOnly(o["sae-hnp-only"], d, "sae_hnp_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["sae-hnp-only"], "WirelessControllerVap-SaeHnpOnly"); ok {
			if err = d.Set("sae_hnp_only", vv); err != nil {
				return fmt.Errorf("Error reading sae_hnp_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sae_hnp_only: %v", err)
		}
	}

	if err = d.Set("sae_pk", flattenWirelessControllerVapSaePk(o["sae-pk"], d, "sae_pk")); err != nil {
		if vv, ok := fortiAPIPatch(o["sae-pk"], "WirelessControllerVap-SaePk"); ok {
			if err = d.Set("sae_pk", vv); err != nil {
				return fmt.Errorf("Error reading sae_pk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sae_pk: %v", err)
		}
	}

	if err = d.Set("sae_private_key", flattenWirelessControllerVapSaePrivateKey(o["sae-private-key"], d, "sae_private_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["sae-private-key"], "WirelessControllerVap-SaePrivateKey"); ok {
			if err = d.Set("sae_private_key", vv); err != nil {
				return fmt.Errorf("Error reading sae_private_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sae_private_key: %v", err)
		}
	}

	if err = d.Set("scan_botnet_connections", flattenWirelessControllerVapScanBotnetConnections(o["scan-botnet-connections"], d, "scan_botnet_connections")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-botnet-connections"], "WirelessControllerVap-ScanBotnetConnections"); ok {
			if err = d.Set("scan_botnet_connections", vv); err != nil {
				return fmt.Errorf("Error reading scan_botnet_connections: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_botnet_connections: %v", err)
		}
	}

	if err = d.Set("schedule", flattenWirelessControllerVapSchedule(o["schedule"], d, "schedule")); err != nil {
		if vv, ok := fortiAPIPatch(o["schedule"], "WirelessControllerVap-Schedule"); ok {
			if err = d.Set("schedule", vv); err != nil {
				return fmt.Errorf("Error reading schedule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("secondary_wag_profile", flattenWirelessControllerVapSecondaryWagProfile(o["secondary-wag-profile"], d, "secondary_wag_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary-wag-profile"], "WirelessControllerVap-SecondaryWagProfile"); ok {
			if err = d.Set("secondary_wag_profile", vv); err != nil {
				return fmt.Errorf("Error reading secondary_wag_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_wag_profile: %v", err)
		}
	}

	if err = d.Set("security", flattenWirelessControllerVapSecurity(o["security"], d, "security")); err != nil {
		if vv, ok := fortiAPIPatch(o["security"], "WirelessControllerVap-Security"); ok {
			if err = d.Set("security", vv); err != nil {
				return fmt.Errorf("Error reading security: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security: %v", err)
		}
	}

	if err = d.Set("security_exempt_list", flattenWirelessControllerVapSecurityExemptList(o["security-exempt-list"], d, "security_exempt_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-exempt-list"], "WirelessControllerVap-SecurityExemptList"); ok {
			if err = d.Set("security_exempt_list", vv); err != nil {
				return fmt.Errorf("Error reading security_exempt_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_exempt_list: %v", err)
		}
	}

	if err = d.Set("security_obsolete_option", flattenWirelessControllerVapSecurityObsoleteOption(o["security-obsolete-option"], d, "security_obsolete_option")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-obsolete-option"], "WirelessControllerVap-SecurityObsoleteOption"); ok {
			if err = d.Set("security_obsolete_option", vv); err != nil {
				return fmt.Errorf("Error reading security_obsolete_option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_obsolete_option: %v", err)
		}
	}

	if err = d.Set("security_redirect_url", flattenWirelessControllerVapSecurityRedirectUrl(o["security-redirect-url"], d, "security_redirect_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-redirect-url"], "WirelessControllerVap-SecurityRedirectUrl"); ok {
			if err = d.Set("security_redirect_url", vv); err != nil {
				return fmt.Errorf("Error reading security_redirect_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_redirect_url: %v", err)
		}
	}

	if err = d.Set("selected_usergroups", flattenWirelessControllerVapSelectedUsergroups(o["selected-usergroups"], d, "selected_usergroups")); err != nil {
		if vv, ok := fortiAPIPatch(o["selected-usergroups"], "WirelessControllerVap-SelectedUsergroups"); ok {
			if err = d.Set("selected_usergroups", vv); err != nil {
				return fmt.Errorf("Error reading selected_usergroups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading selected_usergroups: %v", err)
		}
	}

	if err = d.Set("split_tunneling", flattenWirelessControllerVapSplitTunneling(o["split-tunneling"], d, "split_tunneling")); err != nil {
		if vv, ok := fortiAPIPatch(o["split-tunneling"], "WirelessControllerVap-SplitTunneling"); ok {
			if err = d.Set("split_tunneling", vv); err != nil {
				return fmt.Errorf("Error reading split_tunneling: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading split_tunneling: %v", err)
		}
	}

	if err = d.Set("ssid", flattenWirelessControllerVapSsid(o["ssid"], d, "ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssid"], "WirelessControllerVap-Ssid"); ok {
			if err = d.Set("ssid", vv); err != nil {
				return fmt.Errorf("Error reading ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssid: %v", err)
		}
	}

	if err = d.Set("sticky_client_remove", flattenWirelessControllerVapStickyClientRemove(o["sticky-client-remove"], d, "sticky_client_remove")); err != nil {
		if vv, ok := fortiAPIPatch(o["sticky-client-remove"], "WirelessControllerVap-StickyClientRemove"); ok {
			if err = d.Set("sticky_client_remove", vv); err != nil {
				return fmt.Errorf("Error reading sticky_client_remove: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sticky_client_remove: %v", err)
		}
	}

	if err = d.Set("sticky_client_threshold_2g", flattenWirelessControllerVapStickyClientThreshold2G(o["sticky-client-threshold-2g"], d, "sticky_client_threshold_2g")); err != nil {
		if vv, ok := fortiAPIPatch(o["sticky-client-threshold-2g"], "WirelessControllerVap-StickyClientThreshold2G"); ok {
			if err = d.Set("sticky_client_threshold_2g", vv); err != nil {
				return fmt.Errorf("Error reading sticky_client_threshold_2g: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sticky_client_threshold_2g: %v", err)
		}
	}

	if err = d.Set("sticky_client_threshold_5g", flattenWirelessControllerVapStickyClientThreshold5G(o["sticky-client-threshold-5g"], d, "sticky_client_threshold_5g")); err != nil {
		if vv, ok := fortiAPIPatch(o["sticky-client-threshold-5g"], "WirelessControllerVap-StickyClientThreshold5G"); ok {
			if err = d.Set("sticky_client_threshold_5g", vv); err != nil {
				return fmt.Errorf("Error reading sticky_client_threshold_5g: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sticky_client_threshold_5g: %v", err)
		}
	}

	if err = d.Set("sticky_client_threshold_6g", flattenWirelessControllerVapStickyClientThreshold6G(o["sticky-client-threshold-6g"], d, "sticky_client_threshold_6g")); err != nil {
		if vv, ok := fortiAPIPatch(o["sticky-client-threshold-6g"], "WirelessControllerVap-StickyClientThreshold6G"); ok {
			if err = d.Set("sticky_client_threshold_6g", vv); err != nil {
				return fmt.Errorf("Error reading sticky_client_threshold_6g: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sticky_client_threshold_6g: %v", err)
		}
	}

	if err = d.Set("target_wake_time", flattenWirelessControllerVapTargetWakeTime(o["target-wake-time"], d, "target_wake_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["target-wake-time"], "WirelessControllerVap-TargetWakeTime"); ok {
			if err = d.Set("target_wake_time", vv); err != nil {
				return fmt.Errorf("Error reading target_wake_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading target_wake_time: %v", err)
		}
	}

	if err = d.Set("tkip_counter_measure", flattenWirelessControllerVapTkipCounterMeasure(o["tkip-counter-measure"], d, "tkip_counter_measure")); err != nil {
		if vv, ok := fortiAPIPatch(o["tkip-counter-measure"], "WirelessControllerVap-TkipCounterMeasure"); ok {
			if err = d.Set("tkip_counter_measure", vv); err != nil {
				return fmt.Errorf("Error reading tkip_counter_measure: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tkip_counter_measure: %v", err)
		}
	}

	if err = d.Set("tunnel_echo_interval", flattenWirelessControllerVapTunnelEchoInterval(o["tunnel-echo-interval"], d, "tunnel_echo_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-echo-interval"], "WirelessControllerVap-TunnelEchoInterval"); ok {
			if err = d.Set("tunnel_echo_interval", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_echo_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_echo_interval: %v", err)
		}
	}

	if err = d.Set("tunnel_fallback_interval", flattenWirelessControllerVapTunnelFallbackInterval(o["tunnel-fallback-interval"], d, "tunnel_fallback_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-fallback-interval"], "WirelessControllerVap-TunnelFallbackInterval"); ok {
			if err = d.Set("tunnel_fallback_interval", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_fallback_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_fallback_interval: %v", err)
		}
	}

	if err = d.Set("usergroup", flattenWirelessControllerVapUsergroup(o["usergroup"], d, "usergroup")); err != nil {
		if vv, ok := fortiAPIPatch(o["usergroup"], "WirelessControllerVap-Usergroup"); ok {
			if err = d.Set("usergroup", vv); err != nil {
				return fmt.Errorf("Error reading usergroup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading usergroup: %v", err)
		}
	}

	if err = d.Set("utm_log", flattenWirelessControllerVapUtmLog(o["utm-log"], d, "utm_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["utm-log"], "WirelessControllerVap-UtmLog"); ok {
			if err = d.Set("utm_log", vv); err != nil {
				return fmt.Errorf("Error reading utm_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading utm_log: %v", err)
		}
	}

	if err = d.Set("utm_profile", flattenWirelessControllerVapUtmProfile(o["utm-profile"], d, "utm_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["utm-profile"], "WirelessControllerVap-UtmProfile"); ok {
			if err = d.Set("utm_profile", vv); err != nil {
				return fmt.Errorf("Error reading utm_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading utm_profile: %v", err)
		}
	}

	if err = d.Set("utm_status", flattenWirelessControllerVapUtmStatus(o["utm-status"], d, "utm_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["utm-status"], "WirelessControllerVap-UtmStatus"); ok {
			if err = d.Set("utm_status", vv); err != nil {
				return fmt.Errorf("Error reading utm_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading utm_status: %v", err)
		}
	}

	if err = d.Set("vlan_auto", flattenWirelessControllerVapVlanAuto(o["vlan-auto"], d, "vlan_auto")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-auto"], "WirelessControllerVap-VlanAuto"); ok {
			if err = d.Set("vlan_auto", vv); err != nil {
				return fmt.Errorf("Error reading vlan_auto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_auto: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("vlan_name", flattenWirelessControllerVapVlanName(o["vlan-name"], d, "vlan_name")); err != nil {
			if vv, ok := fortiAPIPatch(o["vlan-name"], "WirelessControllerVap-VlanName"); ok {
				if err = d.Set("vlan_name", vv); err != nil {
					return fmt.Errorf("Error reading vlan_name: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading vlan_name: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vlan_name"); ok {
			if err = d.Set("vlan_name", flattenWirelessControllerVapVlanName(o["vlan-name"], d, "vlan_name")); err != nil {
				if vv, ok := fortiAPIPatch(o["vlan-name"], "WirelessControllerVap-VlanName"); ok {
					if err = d.Set("vlan_name", vv); err != nil {
						return fmt.Errorf("Error reading vlan_name: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading vlan_name: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("vlan_pool", flattenWirelessControllerVapVlanPool(o["vlan-pool"], d, "vlan_pool")); err != nil {
			if vv, ok := fortiAPIPatch(o["vlan-pool"], "WirelessControllerVap-VlanPool"); ok {
				if err = d.Set("vlan_pool", vv); err != nil {
					return fmt.Errorf("Error reading vlan_pool: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading vlan_pool: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vlan_pool"); ok {
			if err = d.Set("vlan_pool", flattenWirelessControllerVapVlanPool(o["vlan-pool"], d, "vlan_pool")); err != nil {
				if vv, ok := fortiAPIPatch(o["vlan-pool"], "WirelessControllerVap-VlanPool"); ok {
					if err = d.Set("vlan_pool", vv); err != nil {
						return fmt.Errorf("Error reading vlan_pool: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading vlan_pool: %v", err)
				}
			}
		}
	}

	if err = d.Set("vlan_pooling", flattenWirelessControllerVapVlanPooling(o["vlan-pooling"], d, "vlan_pooling")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-pooling"], "WirelessControllerVap-VlanPooling"); ok {
			if err = d.Set("vlan_pooling", vv); err != nil {
				return fmt.Errorf("Error reading vlan_pooling: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_pooling: %v", err)
		}
	}

	if err = d.Set("vlanid", flattenWirelessControllerVapVlanid(o["vlanid"], d, "vlanid")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlanid"], "WirelessControllerVap-Vlanid"); ok {
			if err = d.Set("vlanid", vv); err != nil {
				return fmt.Errorf("Error reading vlanid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlanid: %v", err)
		}
	}

	if err = d.Set("voice_enterprise", flattenWirelessControllerVapVoiceEnterprise(o["voice-enterprise"], d, "voice_enterprise")); err != nil {
		if vv, ok := fortiAPIPatch(o["voice-enterprise"], "WirelessControllerVap-VoiceEnterprise"); ok {
			if err = d.Set("voice_enterprise", vv); err != nil {
				return fmt.Errorf("Error reading voice_enterprise: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading voice_enterprise: %v", err)
		}
	}

	if err = d.Set("webfilter_profile", flattenWirelessControllerVapWebfilterProfile(o["webfilter-profile"], d, "webfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-profile"], "WirelessControllerVap-WebfilterProfile"); ok {
			if err = d.Set("webfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_profile: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerVapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerVap80211K(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVap80211V(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapCentmgmt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDhcpSvrId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapIntfAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapIntfDeviceAccessList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapIntfDeviceIdentification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapIntfDeviceNetscan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapIntfDhcpRelayIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapIntfDhcpRelayService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapIntfDhcpRelayType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapIntfDhcp6RelayIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapIntfDhcp6RelayService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapIntfDhcp6RelayType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapIntfIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandWirelessControllerVapIntfIpManagedByFortiipam(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapIntfIp6Address(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapIntfIp6Allowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapIntfListenForticlientConnection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapIntfManagedSubnetworkSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapIntfRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapIsFactorySetting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapAcctInterimInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapAccessControlList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapAdditionalAkms(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapAddressGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapAddressGroupPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapAkm24Only(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapAlias(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapAntivirusProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapApplicationDetectionEngine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapApplicationDscpMarking(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapApplicationList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapApplicationReportIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapAtfWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapAuthCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapAuthPortalAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapBeaconAdvertising(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapBeaconProtection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapBroadcastSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapBroadcastSuppression(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapBssColorPartial(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapBstmDisassociationImminent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapBstmLoadBalancingDisassocTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapBstmRssiDisassocTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapCalledStationIdType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapCaptivePortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapCaptivePortalAcName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapCaptivePortalMacauthRadiusSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapCaptivePortalMacauthRadiusServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapCaptivePortalRadiusSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapCaptivePortalRadiusServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapCaptivePortalSessionTimeoutInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapCaptivePortalAuthTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapCaptivePortalFwAccounting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDhcpAddressEnforcement(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDhcpLeaseTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDhcpOption43Insertion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDhcpOption82CircuitIdInsertion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDhcpOption82Insertion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDhcpOption82RemoteIdInsertion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDomainNameStripping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "n80211k"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["80211k"], _ = expandWirelessControllerVapDynamicMapping80211K(d, i["n80211k"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "n80211v"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["80211v"], _ = expandWirelessControllerVapDynamicMapping80211V(d, i["n80211v"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_centmgmt"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_centmgmt"], _ = expandWirelessControllerVapDynamicMappingCentmgmt(d, i["_centmgmt"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_dhcp_svr_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_dhcp_svr_id"], _ = expandWirelessControllerVapDynamicMappingDhcpSvrId(d, i["_dhcp_svr_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_allowaccess"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_allowaccess"], _ = expandWirelessControllerVapDynamicMappingIntfAllowaccess(d, i["_intf_allowaccess"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_device_access_list"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_device-access-list"], _ = expandWirelessControllerVapDynamicMappingIntfDeviceAccessList(d, i["_intf_device_access_list"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_device_identification"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_device-identification"], _ = expandWirelessControllerVapDynamicMappingIntfDeviceIdentification(d, i["_intf_device_identification"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_device_netscan"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_device-netscan"], _ = expandWirelessControllerVapDynamicMappingIntfDeviceNetscan(d, i["_intf_device_netscan"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_dhcp_relay_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_dhcp-relay-ip"], _ = expandWirelessControllerVapDynamicMappingIntfDhcpRelayIp(d, i["_intf_dhcp_relay_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_dhcp_relay_service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_dhcp-relay-service"], _ = expandWirelessControllerVapDynamicMappingIntfDhcpRelayService(d, i["_intf_dhcp_relay_service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_dhcp_relay_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_dhcp-relay-type"], _ = expandWirelessControllerVapDynamicMappingIntfDhcpRelayType(d, i["_intf_dhcp_relay_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_dhcp6_relay_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_dhcp6-relay-ip"], _ = expandWirelessControllerVapDynamicMappingIntfDhcp6RelayIp(d, i["_intf_dhcp6_relay_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_dhcp6_relay_service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_dhcp6-relay-service"], _ = expandWirelessControllerVapDynamicMappingIntfDhcp6RelayService(d, i["_intf_dhcp6_relay_service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_dhcp6_relay_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_dhcp6-relay-type"], _ = expandWirelessControllerVapDynamicMappingIntfDhcp6RelayType(d, i["_intf_dhcp6_relay_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_ip"], _ = expandWirelessControllerVapDynamicMappingIntfIp(d, i["_intf_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_ip_managed_by_fortiipam"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_ip-managed-by-fortiipam"], _ = expandWirelessControllerVapDynamicMappingIntfIpManagedByFortiipam(d, i["_intf_ip_managed_by_fortiipam"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_ip6_address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_ip6-address"], _ = expandWirelessControllerVapDynamicMappingIntfIp6Address(d, i["_intf_ip6_address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_ip6_allowaccess"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_ip6-allowaccess"], _ = expandWirelessControllerVapDynamicMappingIntfIp6Allowaccess(d, i["_intf_ip6_allowaccess"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_listen_forticlient_connection"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_listen-forticlient-connection"], _ = expandWirelessControllerVapDynamicMappingIntfListenForticlientConnection(d, i["_intf_listen_forticlient_connection"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_managed_subnetwork_size"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_managed-subnetwork-size"], _ = expandWirelessControllerVapDynamicMappingIntfManagedSubnetworkSize(d, i["_intf_managed_subnetwork_size"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_role"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_role"], _ = expandWirelessControllerVapDynamicMappingIntfRole(d, i["_intf_role"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_is_factory_setting"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_is_factory_setting"], _ = expandWirelessControllerVapDynamicMappingIsFactorySetting(d, i["_is_factory_setting"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandWirelessControllerVapDynamicMappingScope(d, i["_scope"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["_scope"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_control_list"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["access-control-list"], _ = expandWirelessControllerVapDynamicMappingAccessControlList(d, i["access_control_list"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "acct_interim_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["acct-interim-interval"], _ = expandWirelessControllerVapDynamicMappingAcctInterimInterval(d, i["acct_interim_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_akms"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["additional-akms"], _ = expandWirelessControllerVapDynamicMappingAdditionalAkms(d, i["additional_akms"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["address-group"], _ = expandWirelessControllerVapDynamicMappingAddressGroup(d, i["address_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address_group_policy"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["address-group-policy"], _ = expandWirelessControllerVapDynamicMappingAddressGroupPolicy(d, i["address_group_policy"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "akm24_only"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["akm24-only"], _ = expandWirelessControllerVapDynamicMappingAkm24Only(d, i["akm24_only"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "alias"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["alias"], _ = expandWirelessControllerVapDynamicMappingAlias(d, i["alias"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "antivirus_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["antivirus-profile"], _ = expandWirelessControllerVapDynamicMappingAntivirusProfile(d, i["antivirus_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application_detection_engine"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["application-detection-engine"], _ = expandWirelessControllerVapDynamicMappingApplicationDetectionEngine(d, i["application_detection_engine"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application_dscp_marking"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["application-dscp-marking"], _ = expandWirelessControllerVapDynamicMappingApplicationDscpMarking(d, i["application_dscp_marking"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application_list"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["application-list"], _ = expandWirelessControllerVapDynamicMappingApplicationList(d, i["application_list"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application_report_intv"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["application-report-intv"], _ = expandWirelessControllerVapDynamicMappingApplicationReportIntv(d, i["application_report_intv"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "atf_weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["atf-weight"], _ = expandWirelessControllerVapDynamicMappingAtfWeight(d, i["atf_weight"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth"], _ = expandWirelessControllerVapDynamicMappingAuth(d, i["auth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-cert"], _ = expandWirelessControllerVapDynamicMappingAuthCert(d, i["auth_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_portal_addr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-portal-addr"], _ = expandWirelessControllerVapDynamicMappingAuthPortalAddr(d, i["auth_portal_addr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "beacon_advertising"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["beacon-advertising"], _ = expandWirelessControllerVapDynamicMappingBeaconAdvertising(d, i["beacon_advertising"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "beacon_protection"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["beacon-protection"], _ = expandWirelessControllerVapDynamicMappingBeaconProtection(d, i["beacon_protection"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "broadcast_ssid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["broadcast-ssid"], _ = expandWirelessControllerVapDynamicMappingBroadcastSsid(d, i["broadcast_ssid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "broadcast_suppression"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["broadcast-suppression"], _ = expandWirelessControllerVapDynamicMappingBroadcastSuppression(d, i["broadcast_suppression"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bss_color_partial"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bss-color-partial"], _ = expandWirelessControllerVapDynamicMappingBssColorPartial(d, i["bss_color_partial"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bstm_disassociation_imminent"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bstm-disassociation-imminent"], _ = expandWirelessControllerVapDynamicMappingBstmDisassociationImminent(d, i["bstm_disassociation_imminent"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bstm_load_balancing_disassoc_timer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bstm-load-balancing-disassoc-timer"], _ = expandWirelessControllerVapDynamicMappingBstmLoadBalancingDisassocTimer(d, i["bstm_load_balancing_disassoc_timer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bstm_rssi_disassoc_timer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bstm-rssi-disassoc-timer"], _ = expandWirelessControllerVapDynamicMappingBstmRssiDisassocTimer(d, i["bstm_rssi_disassoc_timer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "called_station_id_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["called-station-id-type"], _ = expandWirelessControllerVapDynamicMappingCalledStationIdType(d, i["called_station_id_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["captive-portal"], _ = expandWirelessControllerVapDynamicMappingCaptivePortal(d, i["captive_portal"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal_ac_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["captive-portal-ac-name"], _ = expandWirelessControllerVapDynamicMappingCaptivePortalAcName(d, i["captive_portal_ac_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal_auth_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["captive-portal-auth-timeout"], _ = expandWirelessControllerVapDynamicMappingCaptivePortalAuthTimeout(d, i["captive_portal_auth_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal_fw_accounting"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["captive-portal-fw-accounting"], _ = expandWirelessControllerVapDynamicMappingCaptivePortalFwAccounting(d, i["captive_portal_fw_accounting"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal_macauth_radius_secret"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["captive-portal-macauth-radius-secret"], _ = expandWirelessControllerVapDynamicMappingCaptivePortalMacauthRadiusSecret(d, i["captive_portal_macauth_radius_secret"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal_macauth_radius_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["captive-portal-macauth-radius-server"], _ = expandWirelessControllerVapDynamicMappingCaptivePortalMacauthRadiusServer(d, i["captive_portal_macauth_radius_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal_radius_secret"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["captive-portal-radius-secret"], _ = expandWirelessControllerVapDynamicMappingCaptivePortalRadiusSecret(d, i["captive_portal_radius_secret"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal_radius_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["captive-portal-radius-server"], _ = expandWirelessControllerVapDynamicMappingCaptivePortalRadiusServer(d, i["captive_portal_radius_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal_session_timeout_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["captive-portal-session-timeout-interval"], _ = expandWirelessControllerVapDynamicMappingCaptivePortalSessionTimeoutInterval(d, i["captive_portal_session_timeout_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_count"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["client-count"], _ = expandWirelessControllerVapDynamicMappingClientCount(d, i["client_count"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_address_enforcement"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dhcp-address-enforcement"], _ = expandWirelessControllerVapDynamicMappingDhcpAddressEnforcement(d, i["dhcp_address_enforcement"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_lease_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dhcp-lease-time"], _ = expandWirelessControllerVapDynamicMappingDhcpLeaseTime(d, i["dhcp_lease_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_option43_insertion"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dhcp-option43-insertion"], _ = expandWirelessControllerVapDynamicMappingDhcpOption43Insertion(d, i["dhcp_option43_insertion"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_option82_circuit_id_insertion"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dhcp-option82-circuit-id-insertion"], _ = expandWirelessControllerVapDynamicMappingDhcpOption82CircuitIdInsertion(d, i["dhcp_option82_circuit_id_insertion"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_option82_insertion"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dhcp-option82-insertion"], _ = expandWirelessControllerVapDynamicMappingDhcpOption82Insertion(d, i["dhcp_option82_insertion"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_option82_remote_id_insertion"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dhcp-option82-remote-id-insertion"], _ = expandWirelessControllerVapDynamicMappingDhcpOption82RemoteIdInsertion(d, i["dhcp_option82_remote_id_insertion"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain_name_stripping"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["domain-name-stripping"], _ = expandWirelessControllerVapDynamicMappingDomainNameStripping(d, i["domain_name_stripping"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dynamic_vlan"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dynamic-vlan"], _ = expandWirelessControllerVapDynamicMappingDynamicVlan(d, i["dynamic_vlan"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "eap_reauth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["eap-reauth"], _ = expandWirelessControllerVapDynamicMappingEapReauth(d, i["eap_reauth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "eap_reauth_intv"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["eap-reauth-intv"], _ = expandWirelessControllerVapDynamicMappingEapReauthIntv(d, i["eap_reauth_intv"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "eapol_key_retries"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["eapol-key-retries"], _ = expandWirelessControllerVapDynamicMappingEapolKeyRetries(d, i["eapol_key_retries"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "encrypt"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["encrypt"], _ = expandWirelessControllerVapDynamicMappingEncrypt(d, i["encrypt"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_fast_roaming"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["external-fast-roaming"], _ = expandWirelessControllerVapDynamicMappingExternalFastRoaming(d, i["external_fast_roaming"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_logout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["external-logout"], _ = expandWirelessControllerVapDynamicMappingExternalLogout(d, i["external_logout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_pre_auth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["external-pre-auth"], _ = expandWirelessControllerVapDynamicMappingExternalPreAuth(d, i["external_pre_auth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_web"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["external-web"], _ = expandWirelessControllerVapDynamicMappingExternalWeb(d, i["external_web"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_web_format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["external-web-format"], _ = expandWirelessControllerVapDynamicMappingExternalWebFormat(d, i["external_web_format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fast_bss_transition"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fast-bss-transition"], _ = expandWirelessControllerVapDynamicMappingFastBssTransition(d, i["fast_bss_transition"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fast_roaming"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fast-roaming"], _ = expandWirelessControllerVapDynamicMappingFastRoaming(d, i["fast_roaming"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ft_mobility_domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ft-mobility-domain"], _ = expandWirelessControllerVapDynamicMappingFtMobilityDomain(d, i["ft_mobility_domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ft_over_ds"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ft-over-ds"], _ = expandWirelessControllerVapDynamicMappingFtOverDs(d, i["ft_over_ds"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ft_r0_key_lifetime"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ft-r0-key-lifetime"], _ = expandWirelessControllerVapDynamicMappingFtR0KeyLifetime(d, i["ft_r0_key_lifetime"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gas_comeback_delay"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gas-comeback-delay"], _ = expandWirelessControllerVapDynamicMappingGasComebackDelay(d, i["gas_comeback_delay"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gas_fragmentation_limit"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gas-fragmentation-limit"], _ = expandWirelessControllerVapDynamicMappingGasFragmentationLimit(d, i["gas_fragmentation_limit"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gtk_rekey"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gtk-rekey"], _ = expandWirelessControllerVapDynamicMappingGtkRekey(d, i["gtk_rekey"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gtk_rekey_intv"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gtk-rekey-intv"], _ = expandWirelessControllerVapDynamicMappingGtkRekeyIntv(d, i["gtk_rekey_intv"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "high_efficiency"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["high-efficiency"], _ = expandWirelessControllerVapDynamicMappingHighEfficiency(d, i["high_efficiency"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hotspot20_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hotspot20-profile"], _ = expandWirelessControllerVapDynamicMappingHotspot20Profile(d, i["hotspot20_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmp_snooping"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["igmp-snooping"], _ = expandWirelessControllerVapDynamicMappingIgmpSnooping(d, i["igmp_snooping"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "intra_vap_privacy"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["intra-vap-privacy"], _ = expandWirelessControllerVapDynamicMappingIntraVapPrivacy(d, i["intra_vap_privacy"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandWirelessControllerVapDynamicMappingIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ips_sensor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ips-sensor"], _ = expandWirelessControllerVapDynamicMappingIpsSensor(d, i["ips_sensor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_rules"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipv6-rules"], _ = expandWirelessControllerVapDynamicMappingIpv6Rules(d, i["ipv6_rules"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["key"], _ = expandWirelessControllerVapDynamicMappingKey(d, i["key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keyindex"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["keyindex"], _ = expandWirelessControllerVapDynamicMappingKeyindex(d, i["keyindex"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "l3_roaming"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["l3-roaming"], _ = expandWirelessControllerVapDynamicMappingL3Roaming(d, i["l3_roaming"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "l3_roaming_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["l3-roaming-mode"], _ = expandWirelessControllerVapDynamicMappingL3RoamingMode(d, i["l3_roaming_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldpc"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ldpc"], _ = expandWirelessControllerVapDynamicMappingLdpc(d, i["ldpc"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_authentication"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["local-authentication"], _ = expandWirelessControllerVapDynamicMappingLocalAuthentication(d, i["local_authentication"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_bridging"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["local-bridging"], _ = expandWirelessControllerVapDynamicMappingLocalBridging(d, i["local_bridging"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_lan"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["local-lan"], _ = expandWirelessControllerVapDynamicMappingLocalLan(d, i["local_lan"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_lan_partition"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["local-lan-partition"], _ = expandWirelessControllerVapDynamicMappingLocalLanPartition(d, i["local_lan_partition"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_standalone"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["local-standalone"], _ = expandWirelessControllerVapDynamicMappingLocalStandalone(d, i["local_standalone"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_standalone_dns"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["local-standalone-dns"], _ = expandWirelessControllerVapDynamicMappingLocalStandaloneDns(d, i["local_standalone_dns"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_standalone_dns_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["local-standalone-dns-ip"], _ = expandWirelessControllerVapDynamicMappingLocalStandaloneDnsIp(d, i["local_standalone_dns_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_standalone_nat"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["local-standalone-nat"], _ = expandWirelessControllerVapDynamicMappingLocalStandaloneNat(d, i["local_standalone_nat"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_switching"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["local-switching"], _ = expandWirelessControllerVapDynamicMappingLocalSwitching(d, i["local_switching"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_auth_bypass"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac-auth-bypass"], _ = expandWirelessControllerVapDynamicMappingMacAuthBypass(d, i["mac_auth_bypass"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_called_station_delimiter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac-called-station-delimiter"], _ = expandWirelessControllerVapDynamicMappingMacCalledStationDelimiter(d, i["mac_called_station_delimiter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_calling_station_delimiter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac-calling-station-delimiter"], _ = expandWirelessControllerVapDynamicMappingMacCallingStationDelimiter(d, i["mac_calling_station_delimiter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_case"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac-case"], _ = expandWirelessControllerVapDynamicMappingMacCase(d, i["mac_case"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac-filter"], _ = expandWirelessControllerVapDynamicMappingMacFilter(d, i["mac_filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_filter_policy_other"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac-filter-policy-other"], _ = expandWirelessControllerVapDynamicMappingMacFilterPolicyOther(d, i["mac_filter_policy_other"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_password_delimiter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac-password-delimiter"], _ = expandWirelessControllerVapDynamicMappingMacPasswordDelimiter(d, i["mac_password_delimiter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_username_delimiter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac-username-delimiter"], _ = expandWirelessControllerVapDynamicMappingMacUsernameDelimiter(d, i["mac_username_delimiter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_clients"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-clients"], _ = expandWirelessControllerVapDynamicMappingMaxClients(d, i["max_clients"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_clients_ap"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-clients-ap"], _ = expandWirelessControllerVapDynamicMappingMaxClientsAp(d, i["max_clients_ap"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mbo"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mbo"], _ = expandWirelessControllerVapDynamicMappingMbo(d, i["mbo"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mbo_cell_data_conn_pref"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mbo-cell-data-conn-pref"], _ = expandWirelessControllerVapDynamicMappingMboCellDataConnPref(d, i["mbo_cell_data_conn_pref"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "me_disable_thresh"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["me-disable-thresh"], _ = expandWirelessControllerVapDynamicMappingMeDisableThresh(d, i["me_disable_thresh"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mesh_backhaul"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mesh-backhaul"], _ = expandWirelessControllerVapDynamicMappingMeshBackhaul(d, i["mesh_backhaul"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mpsk"], _ = expandWirelessControllerVapDynamicMappingMpsk(d, i["mpsk"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk_concurrent_clients"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mpsk-concurrent-clients"], _ = expandWirelessControllerVapDynamicMappingMpskConcurrentClients(d, i["mpsk_concurrent_clients"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mpsk-profile"], _ = expandWirelessControllerVapDynamicMappingMpskProfile(d, i["mpsk_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mu_mimo"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mu-mimo"], _ = expandWirelessControllerVapDynamicMappingMuMimo(d, i["mu_mimo"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "multicast_enhance"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["multicast-enhance"], _ = expandWirelessControllerVapDynamicMappingMulticastEnhance(d, i["multicast_enhance"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "multicast_rate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["multicast-rate"], _ = expandWirelessControllerVapDynamicMappingMulticastRate(d, i["multicast_rate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nac"], _ = expandWirelessControllerVapDynamicMappingNac(d, i["nac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nac_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nac-profile"], _ = expandWirelessControllerVapDynamicMappingNacProfile(d, i["nac_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nas_filter_rule"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nas-filter-rule"], _ = expandWirelessControllerVapDynamicMappingNasFilterRule(d, i["nas_filter_rule"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbor_report_dual_band"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["neighbor-report-dual-band"], _ = expandWirelessControllerVapDynamicMappingNeighborReportDualBand(d, i["neighbor_report_dual_band"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "okc"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["okc"], _ = expandWirelessControllerVapDynamicMappingOkc(d, i["okc"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "osen"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["osen"], _ = expandWirelessControllerVapDynamicMappingOsen(d, i["osen"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "owe_groups"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["owe-groups"], _ = expandWirelessControllerVapDynamicMappingOweGroups(d, i["owe_groups"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "owe_transition"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["owe-transition"], _ = expandWirelessControllerVapDynamicMappingOweTransition(d, i["owe_transition"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "owe_transition_ssid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["owe-transition-ssid"], _ = expandWirelessControllerVapDynamicMappingOweTransitionSsid(d, i["owe_transition_ssid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passphrase"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["passphrase"], _ = expandWirelessControllerVapDynamicMappingPassphrase(d, i["passphrase"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pmf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pmf"], _ = expandWirelessControllerVapDynamicMappingPmf(d, i["pmf"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pmf_assoc_comeback_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pmf-assoc-comeback-timeout"], _ = expandWirelessControllerVapDynamicMappingPmfAssocComebackTimeout(d, i["pmf_assoc_comeback_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pmf_sa_query_retry_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pmf-sa-query-retry-timeout"], _ = expandWirelessControllerVapDynamicMappingPmfSaQueryRetryTimeout(d, i["pmf_sa_query_retry_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_macauth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port-macauth"], _ = expandWirelessControllerVapDynamicMappingPortMacauth(d, i["port_macauth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_macauth_reauth_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port-macauth-reauth-timeout"], _ = expandWirelessControllerVapDynamicMappingPortMacauthReauthTimeout(d, i["port_macauth_reauth_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_macauth_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port-macauth-timeout"], _ = expandWirelessControllerVapDynamicMappingPortMacauthTimeout(d, i["port_macauth_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "portal_message_override_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["portal-message-override-group"], _ = expandWirelessControllerVapDynamicMappingPortalMessageOverrideGroup(d, i["portal_message_override_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "portal_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["portal-type"], _ = expandWirelessControllerVapDynamicMappingPortalType(d, i["portal_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pre_auth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pre-auth"], _ = expandWirelessControllerVapDynamicMappingPreAuth(d, i["pre_auth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "primary_wag_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["primary-wag-profile"], _ = expandWirelessControllerVapDynamicMappingPrimaryWagProfile(d, i["primary_wag_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_resp_suppression"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["probe-resp-suppression"], _ = expandWirelessControllerVapDynamicMappingProbeRespSuppression(d, i["probe_resp_suppression"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_resp_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["probe-resp-threshold"], _ = expandWirelessControllerVapDynamicMappingProbeRespThreshold(d, i["probe_resp_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ptk_rekey"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ptk-rekey"], _ = expandWirelessControllerVapDynamicMappingPtkRekey(d, i["ptk_rekey"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ptk_rekey_intv"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ptk-rekey-intv"], _ = expandWirelessControllerVapDynamicMappingPtkRekeyIntv(d, i["ptk_rekey_intv"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "qos_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["qos-profile"], _ = expandWirelessControllerVapDynamicMappingQosProfile(d, i["qos_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine"], _ = expandWirelessControllerVapDynamicMappingQuarantine(d, i["quarantine"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radio_2g_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["radio-2g-threshold"], _ = expandWirelessControllerVapDynamicMappingRadio2GThreshold(d, i["radio_2g_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radio_5g_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["radio-5g-threshold"], _ = expandWirelessControllerVapDynamicMappingRadio5GThreshold(d, i["radio_5g_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radio_sensitivity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["radio-sensitivity"], _ = expandWirelessControllerVapDynamicMappingRadioSensitivity(d, i["radio_sensitivity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_mac_auth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["radius-mac-auth"], _ = expandWirelessControllerVapDynamicMappingRadiusMacAuth(d, i["radius_mac_auth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_mac_auth_block_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["radius-mac-auth-block-interval"], _ = expandWirelessControllerVapDynamicMappingRadiusMacAuthBlockInterval(d, i["radius_mac_auth_block_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_mac_auth_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["radius-mac-auth-server"], _ = expandWirelessControllerVapDynamicMappingRadiusMacAuthServer(d, i["radius_mac_auth_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_mac_auth_usergroups"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["radius-mac-auth-usergroups"], _ = expandWirelessControllerVapDynamicMappingRadiusMacAuthUsergroups(d, i["radius_mac_auth_usergroups"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_mac_mpsk_auth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["radius-mac-mpsk-auth"], _ = expandWirelessControllerVapDynamicMappingRadiusMacMpskAuth(d, i["radius_mac_mpsk_auth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_mac_mpsk_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["radius-mac-mpsk-timeout"], _ = expandWirelessControllerVapDynamicMappingRadiusMacMpskTimeout(d, i["radius_mac_mpsk_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["radius-server"], _ = expandWirelessControllerVapDynamicMappingRadiusServer(d, i["radius_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11a"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rates-11a"], _ = expandWirelessControllerVapDynamicMappingRates11A(d, i["rates_11a"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11ac_mcs_map"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rates-11ac-mcs-map"], _ = expandWirelessControllerVapDynamicMappingRates11AcMcsMap(d, i["rates_11ac_mcs_map"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11ac_ss12"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rates-11ac-ss12"], _ = expandWirelessControllerVapDynamicMappingRates11AcSs12(d, i["rates_11ac_ss12"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11ac_ss34"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rates-11ac-ss34"], _ = expandWirelessControllerVapDynamicMappingRates11AcSs34(d, i["rates_11ac_ss34"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11ax_mcs_map"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rates-11ax-mcs-map"], _ = expandWirelessControllerVapDynamicMappingRates11AxMcsMap(d, i["rates_11ax_mcs_map"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11ax_ss12"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rates-11ax-ss12"], _ = expandWirelessControllerVapDynamicMappingRates11AxSs12(d, i["rates_11ax_ss12"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11ax_ss34"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rates-11ax-ss34"], _ = expandWirelessControllerVapDynamicMappingRates11AxSs34(d, i["rates_11ax_ss34"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11be_mcs_map"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rates-11be-mcs-map"], _ = expandWirelessControllerVapDynamicMappingRates11BeMcsMap(d, i["rates_11be_mcs_map"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11be_mcs_map_160"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rates-11be-mcs-map-160"], _ = expandWirelessControllerVapDynamicMappingRates11BeMcsMap160(d, i["rates_11be_mcs_map_160"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11be_mcs_map_320"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rates-11be-mcs-map-320"], _ = expandWirelessControllerVapDynamicMappingRates11BeMcsMap320(d, i["rates_11be_mcs_map_320"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11bg"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rates-11bg"], _ = expandWirelessControllerVapDynamicMappingRates11Bg(d, i["rates_11bg"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11n_ss12"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rates-11n-ss12"], _ = expandWirelessControllerVapDynamicMappingRates11NSs12(d, i["rates_11n_ss12"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11n_ss34"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rates-11n-ss34"], _ = expandWirelessControllerVapDynamicMappingRates11NSs34(d, i["rates_11n_ss34"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "roaming_acct_interim_update"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["roaming-acct-interim-update"], _ = expandWirelessControllerVapDynamicMappingRoamingAcctInterimUpdate(d, i["roaming_acct_interim_update"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_groups"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sae-groups"], _ = expandWirelessControllerVapDynamicMappingSaeGroups(d, i["sae_groups"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_h2e_only"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sae-h2e-only"], _ = expandWirelessControllerVapDynamicMappingSaeH2EOnly(d, i["sae_h2e_only"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_hnp_only"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sae-hnp-only"], _ = expandWirelessControllerVapDynamicMappingSaeHnpOnly(d, i["sae_hnp_only"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sae-password"], _ = expandWirelessControllerVapDynamicMappingSaePassword(d, i["sae_password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_pk"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sae-pk"], _ = expandWirelessControllerVapDynamicMappingSaePk(d, i["sae_pk"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_private_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sae-private-key"], _ = expandWirelessControllerVapDynamicMappingSaePrivateKey(d, i["sae_private_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "scan_botnet_connections"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["scan-botnet-connections"], _ = expandWirelessControllerVapDynamicMappingScanBotnetConnections(d, i["scan_botnet_connections"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "schedule"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["schedule"], _ = expandWirelessControllerVapDynamicMappingSchedule(d, i["schedule"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secondary_wag_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["secondary-wag-profile"], _ = expandWirelessControllerVapDynamicMappingSecondaryWagProfile(d, i["secondary_wag_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["security"], _ = expandWirelessControllerVapDynamicMappingSecurity(d, i["security"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_exempt_list"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["security-exempt-list"], _ = expandWirelessControllerVapDynamicMappingSecurityExemptList(d, i["security_exempt_list"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_obsolete_option"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["security-obsolete-option"], _ = expandWirelessControllerVapDynamicMappingSecurityObsoleteOption(d, i["security_obsolete_option"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_redirect_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["security-redirect-url"], _ = expandWirelessControllerVapDynamicMappingSecurityRedirectUrl(d, i["security_redirect_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "selected_usergroups"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["selected-usergroups"], _ = expandWirelessControllerVapDynamicMappingSelectedUsergroups(d, i["selected_usergroups"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_tunneling"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["split-tunneling"], _ = expandWirelessControllerVapDynamicMappingSplitTunneling(d, i["split_tunneling"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssid"], _ = expandWirelessControllerVapDynamicMappingSsid(d, i["ssid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sticky_client_remove"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sticky-client-remove"], _ = expandWirelessControllerVapDynamicMappingStickyClientRemove(d, i["sticky_client_remove"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sticky_client_threshold_2g"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sticky-client-threshold-2g"], _ = expandWirelessControllerVapDynamicMappingStickyClientThreshold2G(d, i["sticky_client_threshold_2g"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sticky_client_threshold_5g"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sticky-client-threshold-5g"], _ = expandWirelessControllerVapDynamicMappingStickyClientThreshold5G(d, i["sticky_client_threshold_5g"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sticky_client_threshold_6g"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sticky-client-threshold-6g"], _ = expandWirelessControllerVapDynamicMappingStickyClientThreshold6G(d, i["sticky_client_threshold_6g"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target_wake_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["target-wake-time"], _ = expandWirelessControllerVapDynamicMappingTargetWakeTime(d, i["target_wake_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tkip_counter_measure"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tkip-counter-measure"], _ = expandWirelessControllerVapDynamicMappingTkipCounterMeasure(d, i["tkip_counter_measure"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tunnel_echo_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tunnel-echo-interval"], _ = expandWirelessControllerVapDynamicMappingTunnelEchoInterval(d, i["tunnel_echo_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tunnel_fallback_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tunnel-fallback-interval"], _ = expandWirelessControllerVapDynamicMappingTunnelFallbackInterval(d, i["tunnel_fallback_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "usergroup"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["usergroup"], _ = expandWirelessControllerVapDynamicMappingUsergroup(d, i["usergroup"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "utm_log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["utm-log"], _ = expandWirelessControllerVapDynamicMappingUtmLog(d, i["utm_log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "utm_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["utm-profile"], _ = expandWirelessControllerVapDynamicMappingUtmProfile(d, i["utm_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "utm_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["utm-status"], _ = expandWirelessControllerVapDynamicMappingUtmStatus(d, i["utm_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandWirelessControllerVapDynamicMappingVdom(d, i["vdom"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_auto"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlan-auto"], _ = expandWirelessControllerVapDynamicMappingVlanAuto(d, i["vlan_auto"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_pooling"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlan-pooling"], _ = expandWirelessControllerVapDynamicMappingVlanPooling(d, i["vlan_pooling"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlanid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlanid"], _ = expandWirelessControllerVapDynamicMappingVlanid(d, i["vlanid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "voice_enterprise"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["voice-enterprise"], _ = expandWirelessControllerVapDynamicMappingVoiceEnterprise(d, i["voice_enterprise"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "webfilter_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["webfilter-profile"], _ = expandWirelessControllerVapDynamicMappingWebfilterProfile(d, i["webfilter_profile"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerVapDynamicMapping80211K(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMapping80211V(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingCentmgmt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingDhcpSvrId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingIntfAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingIntfDeviceAccessList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingIntfDeviceIdentification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingIntfDeviceNetscan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingIntfDhcpRelayIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingIntfDhcpRelayService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingIntfDhcpRelayType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingIntfDhcp6RelayIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingIntfDhcp6RelayService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingIntfDhcp6RelayType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingIntfIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandWirelessControllerVapDynamicMappingIntfIpManagedByFortiipam(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingIntfIp6Address(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingIntfIp6Allowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingIntfListenForticlientConnection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingIntfManagedSubnetworkSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingIntfRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingIsFactorySetting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandWirelessControllerVapDynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandWirelessControllerVapDynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerVapDynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingAccessControlList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingAcctInterimInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingAdditionalAkms(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingAddressGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingAddressGroupPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingAkm24Only(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingAlias(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingAntivirusProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingApplicationDetectionEngine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingApplicationDscpMarking(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingApplicationList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingApplicationReportIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingAtfWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingAuthCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingAuthPortalAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingBeaconAdvertising(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingBeaconProtection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingBroadcastSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingBroadcastSuppression(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingBssColorPartial(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingBstmDisassociationImminent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingBstmLoadBalancingDisassocTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingBstmRssiDisassocTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingCalledStationIdType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingCaptivePortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingCaptivePortalAcName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingCaptivePortalAuthTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingCaptivePortalFwAccounting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingCaptivePortalMacauthRadiusSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingCaptivePortalMacauthRadiusServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingCaptivePortalRadiusSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingCaptivePortalRadiusServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingCaptivePortalSessionTimeoutInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingClientCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingDhcpAddressEnforcement(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingDhcpLeaseTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingDhcpOption43Insertion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingDhcpOption82CircuitIdInsertion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingDhcpOption82Insertion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingDhcpOption82RemoteIdInsertion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingDomainNameStripping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingDynamicVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingEapReauth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingEapReauthIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingEapolKeyRetries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingEncrypt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingExternalFastRoaming(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingExternalLogout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingExternalPreAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingExternalWeb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingExternalWebFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingFastBssTransition(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingFastRoaming(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingFtMobilityDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingFtOverDs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingFtR0KeyLifetime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingGasComebackDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingGasFragmentationLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingGtkRekey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingGtkRekeyIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingHighEfficiency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingHotspot20Profile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingIgmpSnooping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingIntraVapPrivacy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandWirelessControllerVapDynamicMappingIpsSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingIpv6Rules(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingKeyindex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingL3Roaming(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingL3RoamingMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingLdpc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingLocalAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingLocalBridging(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingLocalLan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingLocalLanPartition(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingLocalStandalone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingLocalStandaloneDns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingLocalStandaloneDnsIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingLocalStandaloneNat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingLocalSwitching(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingMacAuthBypass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingMacCalledStationDelimiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingMacCallingStationDelimiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingMacCase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingMacFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingMacFilterPolicyOther(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingMacPasswordDelimiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingMacUsernameDelimiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingMaxClients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingMaxClientsAp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingMbo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingMboCellDataConnPref(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingMeDisableThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingMeshBackhaul(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingMpsk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingMpskConcurrentClients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingMpskProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingMuMimo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingMulticastEnhance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingMulticastRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingNac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingNacProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingNasFilterRule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingNeighborReportDualBand(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingOkc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingOsen(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingOweGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingOweTransition(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingOweTransitionSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingPassphrase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingPmf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingPmfAssocComebackTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingPmfSaQueryRetryTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingPortMacauth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingPortMacauthReauthTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingPortMacauthTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingPortalMessageOverrideGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingPortalType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingPreAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingPrimaryWagProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingProbeRespSuppression(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingProbeRespThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingPtkRekey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingPtkRekeyIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingQosProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingRadio2GThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingRadio5GThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingRadioSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingRadiusMacAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingRadiusMacAuthBlockInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingRadiusMacAuthServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingRadiusMacAuthUsergroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingRadiusMacMpskAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingRadiusMacMpskTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingRadiusServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingRates11A(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingRates11AcMcsMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingRates11AcSs12(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingRates11AcSs34(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingRates11AxMcsMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingRates11AxSs12(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingRates11AxSs34(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingRates11BeMcsMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingRates11BeMcsMap160(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingRates11BeMcsMap320(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingRates11Bg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingRates11NSs12(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingRates11NSs34(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingRoamingAcctInterimUpdate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingSaeGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingSaeH2EOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingSaeHnpOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingSaePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingSaePk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingSaePrivateKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingScanBotnetConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingSchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingSecondaryWagProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingSecurity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingSecurityExemptList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingSecurityObsoleteOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingSecurityRedirectUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingSelectedUsergroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingSplitTunneling(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingStickyClientRemove(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingStickyClientThreshold2G(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingStickyClientThreshold5G(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingStickyClientThreshold6G(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingTargetWakeTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingTkipCounterMeasure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingTunnelEchoInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingTunnelFallbackInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingUsergroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingUtmLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingUtmProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingUtmStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingVlanAuto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingVlanPooling(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingVlanid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingVoiceEnterprise(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingWebfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapEapReauth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapEapReauthIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapEapolKeyRetries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapEncrypt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapExternalFastRoaming(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapExternalLogout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapExternalPreAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapExternalWeb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapExternalWebFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapFastBssTransition(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapFastRoaming(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapFtMobilityDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapFtOverDs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapFtR0KeyLifetime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapGasComebackDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapGasFragmentationLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapGtkRekey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapGtkRekeyIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapHighEfficiency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapHotspot20Profile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapIgmpSnooping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapIntraVapPrivacy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandWirelessControllerVapIpsSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapIpv6Rules(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapKeyindex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapL3Roaming(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapL3RoamingMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapLdpc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapLocalAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapLocalBridging(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapLocalLan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapLocalLanPartition(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapLocalStandalone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapLocalStandaloneDns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapLocalStandaloneDnsIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapLocalStandaloneNat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMacAuthBypass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMacCalledStationDelimiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMacCallingStationDelimiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMacCase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMacFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMacFilterList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWirelessControllerVapMacFilterListId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac"], _ = expandWirelessControllerVapMacFilterListMac(d, i["mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_filter_policy"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac-filter-policy"], _ = expandWirelessControllerVapMacFilterListMacFilterPolicy(d, i["mac_filter_policy"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerVapMacFilterListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMacFilterListMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMacFilterListMacFilterPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMacFilterPolicyOther(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMacPasswordDelimiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMacUsernameDelimiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMaxClients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMaxClientsAp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMbo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMboCellDataConnPref(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMeDisableThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMeshBackhaul(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMpsk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMpskConcurrentClients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMpskKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandWirelessControllerVapMpskKeyComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "concurrent_clients"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["concurrent-clients"], _ = expandWirelessControllerVapMpskKeyConcurrentClients(d, i["concurrent_clients"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["key-name"], _ = expandWirelessControllerVapMpskKeyKeyName(d, i["key_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk_schedules"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mpsk-schedules"], _ = expandWirelessControllerVapMpskKeyMpskSchedules(d, i["mpsk_schedules"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passphrase"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["passphrase"], _ = expandWirelessControllerVapMpskKeyPassphrase(d, i["passphrase"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerVapMpskKeyComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMpskKeyConcurrentClients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMpskKeyKeyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMpskKeyMpskSchedules(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapMpskKeyPassphrase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapMpskProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapMuMimo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMulticastEnhance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMulticastRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapNac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapNacProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapNasFilterRule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapNeighborReportDualBand(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapOkc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapOsen(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapOweGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapOweTransition(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapOweTransitionSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPassphrase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapPmf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPmfAssocComebackTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPmfSaQueryRetryTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPortMacauth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPortMacauthReauthTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPortMacauthTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPortalMessageOverrideGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapPortalMessageOverrides(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auth_disclaimer_page"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auth-disclaimer-page"], _ = expandWirelessControllerVapPortalMessageOverridesAuthDisclaimerPage(d, i["auth_disclaimer_page"], pre_append)
	}
	pre_append = pre + ".0." + "auth_login_failed_page"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auth-login-failed-page"], _ = expandWirelessControllerVapPortalMessageOverridesAuthLoginFailedPage(d, i["auth_login_failed_page"], pre_append)
	}
	pre_append = pre + ".0." + "auth_login_page"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auth-login-page"], _ = expandWirelessControllerVapPortalMessageOverridesAuthLoginPage(d, i["auth_login_page"], pre_append)
	}
	pre_append = pre + ".0." + "auth_reject_page"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auth-reject-page"], _ = expandWirelessControllerVapPortalMessageOverridesAuthRejectPage(d, i["auth_reject_page"], pre_append)
	}

	return result, nil
}

func expandWirelessControllerVapPortalMessageOverridesAuthDisclaimerPage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPortalMessageOverridesAuthLoginFailedPage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPortalMessageOverridesAuthLoginPage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPortalMessageOverridesAuthRejectPage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPortalType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPreAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPrimaryWagProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapProbeRespSuppression(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapProbeRespThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPtkRekey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPtkRekeyIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapQosProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRadio2GThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRadio5GThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRadioSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRadiusMacAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRadiusMacAuthBlockInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRadiusMacAuthServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapRadiusMacAuthUsergroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapRadiusMacMpskAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRadiusMacMpskTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRadiusServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapRates11A(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapRates11AcMcsMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRates11AcSs12(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapRates11AcSs34(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapRates11AxMcsMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRates11AxSs12(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapRates11AxSs34(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapRates11BeMcsMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRates11BeMcsMap160(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRates11BeMcsMap320(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRates11Bg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapRates11NSs12(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapRates11NSs34(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapRoamingAcctInterimUpdate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSaeGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapSaeH2EOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSaeHnpOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSaePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapSaePk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSaePrivateKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapScanBotnetConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapSecondaryWagProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapSecurity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSecurityExemptList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapSecurityObsoleteOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSecurityRedirectUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSelectedUsergroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapSplitTunneling(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapStickyClientRemove(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapStickyClientThreshold2G(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapStickyClientThreshold5G(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapStickyClientThreshold6G(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapTargetWakeTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapTkipCounterMeasure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapTunnelEchoInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapTunnelFallbackInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapUsergroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapUtmLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapUtmProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapUtmStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapVlanAuto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapVlanName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandWirelessControllerVapVlanNameName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlan-id"], _ = expandWirelessControllerVapVlanNameVlanId(d, i["vlan_id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerVapVlanNameName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapVlanNameVlanId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapVlanPool(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWirelessControllerVapVlanPoolId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wtp_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["wtp-group"], _ = expandWirelessControllerVapVlanPoolWtpGroup(d, i["wtp_group"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerVapVlanPoolId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapVlanPoolWtpGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapVlanPooling(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapVlanid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapVoiceEnterprise(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapWebfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectWirelessControllerVap(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("n80211k"); ok || d.HasChange("n80211k") {
		t, err := expandWirelessControllerVap80211K(d, v, "n80211k")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["80211k"] = t
		}
	}

	if v, ok := d.GetOk("n80211v"); ok || d.HasChange("n80211v") {
		t, err := expandWirelessControllerVap80211V(d, v, "n80211v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["80211v"] = t
		}
	}

	if v, ok := d.GetOk("_centmgmt"); ok || d.HasChange("_centmgmt") {
		t, err := expandWirelessControllerVapCentmgmt(d, v, "_centmgmt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_centmgmt"] = t
		}
	}

	if v, ok := d.GetOk("_dhcp_svr_id"); ok || d.HasChange("_dhcp_svr_id") {
		t, err := expandWirelessControllerVapDhcpSvrId(d, v, "_dhcp_svr_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_dhcp_svr_id"] = t
		}
	}

	if v, ok := d.GetOk("_intf_allowaccess"); ok || d.HasChange("_intf_allowaccess") {
		t, err := expandWirelessControllerVapIntfAllowaccess(d, v, "_intf_allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("_intf_device_access_list"); ok || d.HasChange("_intf_device_access_list") {
		t, err := expandWirelessControllerVapIntfDeviceAccessList(d, v, "_intf_device_access_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_device-access-list"] = t
		}
	}

	if v, ok := d.GetOk("_intf_device_identification"); ok || d.HasChange("_intf_device_identification") {
		t, err := expandWirelessControllerVapIntfDeviceIdentification(d, v, "_intf_device_identification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_device-identification"] = t
		}
	}

	if v, ok := d.GetOk("_intf_device_netscan"); ok || d.HasChange("_intf_device_netscan") {
		t, err := expandWirelessControllerVapIntfDeviceNetscan(d, v, "_intf_device_netscan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_device-netscan"] = t
		}
	}

	if v, ok := d.GetOk("_intf_dhcp_relay_ip"); ok || d.HasChange("_intf_dhcp_relay_ip") {
		t, err := expandWirelessControllerVapIntfDhcpRelayIp(d, v, "_intf_dhcp_relay_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_dhcp-relay-ip"] = t
		}
	}

	if v, ok := d.GetOk("_intf_dhcp_relay_service"); ok || d.HasChange("_intf_dhcp_relay_service") {
		t, err := expandWirelessControllerVapIntfDhcpRelayService(d, v, "_intf_dhcp_relay_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_dhcp-relay-service"] = t
		}
	}

	if v, ok := d.GetOk("_intf_dhcp_relay_type"); ok || d.HasChange("_intf_dhcp_relay_type") {
		t, err := expandWirelessControllerVapIntfDhcpRelayType(d, v, "_intf_dhcp_relay_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_dhcp-relay-type"] = t
		}
	}

	if v, ok := d.GetOk("_intf_dhcp6_relay_ip"); ok || d.HasChange("_intf_dhcp6_relay_ip") {
		t, err := expandWirelessControllerVapIntfDhcp6RelayIp(d, v, "_intf_dhcp6_relay_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_dhcp6-relay-ip"] = t
		}
	}

	if v, ok := d.GetOk("_intf_dhcp6_relay_service"); ok || d.HasChange("_intf_dhcp6_relay_service") {
		t, err := expandWirelessControllerVapIntfDhcp6RelayService(d, v, "_intf_dhcp6_relay_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_dhcp6-relay-service"] = t
		}
	}

	if v, ok := d.GetOk("_intf_dhcp6_relay_type"); ok || d.HasChange("_intf_dhcp6_relay_type") {
		t, err := expandWirelessControllerVapIntfDhcp6RelayType(d, v, "_intf_dhcp6_relay_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_dhcp6-relay-type"] = t
		}
	}

	if v, ok := d.GetOk("_intf_ip"); ok || d.HasChange("_intf_ip") {
		t, err := expandWirelessControllerVapIntfIp(d, v, "_intf_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_ip"] = t
		}
	}

	if v, ok := d.GetOk("_intf_ip_managed_by_fortiipam"); ok || d.HasChange("_intf_ip_managed_by_fortiipam") {
		t, err := expandWirelessControllerVapIntfIpManagedByFortiipam(d, v, "_intf_ip_managed_by_fortiipam")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_ip-managed-by-fortiipam"] = t
		}
	}

	if v, ok := d.GetOk("_intf_ip6_address"); ok || d.HasChange("_intf_ip6_address") {
		t, err := expandWirelessControllerVapIntfIp6Address(d, v, "_intf_ip6_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_ip6-address"] = t
		}
	}

	if v, ok := d.GetOk("_intf_ip6_allowaccess"); ok || d.HasChange("_intf_ip6_allowaccess") {
		t, err := expandWirelessControllerVapIntfIp6Allowaccess(d, v, "_intf_ip6_allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_ip6-allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("_intf_listen_forticlient_connection"); ok || d.HasChange("_intf_listen_forticlient_connection") {
		t, err := expandWirelessControllerVapIntfListenForticlientConnection(d, v, "_intf_listen_forticlient_connection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_listen-forticlient-connection"] = t
		}
	}

	if v, ok := d.GetOk("_intf_managed_subnetwork_size"); ok || d.HasChange("_intf_managed_subnetwork_size") {
		t, err := expandWirelessControllerVapIntfManagedSubnetworkSize(d, v, "_intf_managed_subnetwork_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_managed-subnetwork-size"] = t
		}
	}

	if v, ok := d.GetOk("_intf_role"); ok || d.HasChange("_intf_role") {
		t, err := expandWirelessControllerVapIntfRole(d, v, "_intf_role")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_role"] = t
		}
	}

	if v, ok := d.GetOk("_is_factory_setting"); ok || d.HasChange("_is_factory_setting") {
		t, err := expandWirelessControllerVapIsFactorySetting(d, v, "_is_factory_setting")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_is_factory_setting"] = t
		}
	}

	if v, ok := d.GetOk("acct_interim_interval"); ok || d.HasChange("acct_interim_interval") {
		t, err := expandWirelessControllerVapAcctInterimInterval(d, v, "acct_interim_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acct-interim-interval"] = t
		}
	}

	if v, ok := d.GetOk("access_control_list"); ok || d.HasChange("access_control_list") {
		t, err := expandWirelessControllerVapAccessControlList(d, v, "access_control_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-control-list"] = t
		}
	}

	if v, ok := d.GetOk("additional_akms"); ok || d.HasChange("additional_akms") {
		t, err := expandWirelessControllerVapAdditionalAkms(d, v, "additional_akms")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-akms"] = t
		}
	}

	if v, ok := d.GetOk("address_group"); ok || d.HasChange("address_group") {
		t, err := expandWirelessControllerVapAddressGroup(d, v, "address_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address-group"] = t
		}
	}

	if v, ok := d.GetOk("address_group_policy"); ok || d.HasChange("address_group_policy") {
		t, err := expandWirelessControllerVapAddressGroupPolicy(d, v, "address_group_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address-group-policy"] = t
		}
	}

	if v, ok := d.GetOk("akm24_only"); ok || d.HasChange("akm24_only") {
		t, err := expandWirelessControllerVapAkm24Only(d, v, "akm24_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["akm24-only"] = t
		}
	}

	if v, ok := d.GetOk("alias"); ok || d.HasChange("alias") {
		t, err := expandWirelessControllerVapAlias(d, v, "alias")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alias"] = t
		}
	}

	if v, ok := d.GetOk("antivirus_profile"); ok || d.HasChange("antivirus_profile") {
		t, err := expandWirelessControllerVapAntivirusProfile(d, v, "antivirus_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antivirus-profile"] = t
		}
	}

	if v, ok := d.GetOk("application_detection_engine"); ok || d.HasChange("application_detection_engine") {
		t, err := expandWirelessControllerVapApplicationDetectionEngine(d, v, "application_detection_engine")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-detection-engine"] = t
		}
	}

	if v, ok := d.GetOk("application_dscp_marking"); ok || d.HasChange("application_dscp_marking") {
		t, err := expandWirelessControllerVapApplicationDscpMarking(d, v, "application_dscp_marking")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-dscp-marking"] = t
		}
	}

	if v, ok := d.GetOk("application_list"); ok || d.HasChange("application_list") {
		t, err := expandWirelessControllerVapApplicationList(d, v, "application_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-list"] = t
		}
	}

	if v, ok := d.GetOk("application_report_intv"); ok || d.HasChange("application_report_intv") {
		t, err := expandWirelessControllerVapApplicationReportIntv(d, v, "application_report_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-report-intv"] = t
		}
	}

	if v, ok := d.GetOk("atf_weight"); ok || d.HasChange("atf_weight") {
		t, err := expandWirelessControllerVapAtfWeight(d, v, "atf_weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["atf-weight"] = t
		}
	}

	if v, ok := d.GetOk("auth"); ok || d.HasChange("auth") {
		t, err := expandWirelessControllerVapAuth(d, v, "auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth"] = t
		}
	}

	if v, ok := d.GetOk("auth_cert"); ok || d.HasChange("auth_cert") {
		t, err := expandWirelessControllerVapAuthCert(d, v, "auth_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-cert"] = t
		}
	}

	if v, ok := d.GetOk("auth_portal_addr"); ok || d.HasChange("auth_portal_addr") {
		t, err := expandWirelessControllerVapAuthPortalAddr(d, v, "auth_portal_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-portal-addr"] = t
		}
	}

	if v, ok := d.GetOk("beacon_advertising"); ok || d.HasChange("beacon_advertising") {
		t, err := expandWirelessControllerVapBeaconAdvertising(d, v, "beacon_advertising")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["beacon-advertising"] = t
		}
	}

	if v, ok := d.GetOk("beacon_protection"); ok || d.HasChange("beacon_protection") {
		t, err := expandWirelessControllerVapBeaconProtection(d, v, "beacon_protection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["beacon-protection"] = t
		}
	}

	if v, ok := d.GetOk("broadcast_ssid"); ok || d.HasChange("broadcast_ssid") {
		t, err := expandWirelessControllerVapBroadcastSsid(d, v, "broadcast_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["broadcast-ssid"] = t
		}
	}

	if v, ok := d.GetOk("broadcast_suppression"); ok || d.HasChange("broadcast_suppression") {
		t, err := expandWirelessControllerVapBroadcastSuppression(d, v, "broadcast_suppression")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["broadcast-suppression"] = t
		}
	}

	if v, ok := d.GetOk("bss_color_partial"); ok || d.HasChange("bss_color_partial") {
		t, err := expandWirelessControllerVapBssColorPartial(d, v, "bss_color_partial")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bss-color-partial"] = t
		}
	}

	if v, ok := d.GetOk("bstm_disassociation_imminent"); ok || d.HasChange("bstm_disassociation_imminent") {
		t, err := expandWirelessControllerVapBstmDisassociationImminent(d, v, "bstm_disassociation_imminent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bstm-disassociation-imminent"] = t
		}
	}

	if v, ok := d.GetOk("bstm_load_balancing_disassoc_timer"); ok || d.HasChange("bstm_load_balancing_disassoc_timer") {
		t, err := expandWirelessControllerVapBstmLoadBalancingDisassocTimer(d, v, "bstm_load_balancing_disassoc_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bstm-load-balancing-disassoc-timer"] = t
		}
	}

	if v, ok := d.GetOk("bstm_rssi_disassoc_timer"); ok || d.HasChange("bstm_rssi_disassoc_timer") {
		t, err := expandWirelessControllerVapBstmRssiDisassocTimer(d, v, "bstm_rssi_disassoc_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bstm-rssi-disassoc-timer"] = t
		}
	}

	if v, ok := d.GetOk("called_station_id_type"); ok || d.HasChange("called_station_id_type") {
		t, err := expandWirelessControllerVapCalledStationIdType(d, v, "called_station_id_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["called-station-id-type"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal"); ok || d.HasChange("captive_portal") {
		t, err := expandWirelessControllerVapCaptivePortal(d, v, "captive_portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_ac_name"); ok || d.HasChange("captive_portal_ac_name") {
		t, err := expandWirelessControllerVapCaptivePortalAcName(d, v, "captive_portal_ac_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-ac-name"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_macauth_radius_secret"); ok || d.HasChange("captive_portal_macauth_radius_secret") {
		t, err := expandWirelessControllerVapCaptivePortalMacauthRadiusSecret(d, v, "captive_portal_macauth_radius_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-macauth-radius-secret"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_macauth_radius_server"); ok || d.HasChange("captive_portal_macauth_radius_server") {
		t, err := expandWirelessControllerVapCaptivePortalMacauthRadiusServer(d, v, "captive_portal_macauth_radius_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-macauth-radius-server"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_radius_secret"); ok || d.HasChange("captive_portal_radius_secret") {
		t, err := expandWirelessControllerVapCaptivePortalRadiusSecret(d, v, "captive_portal_radius_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-radius-secret"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_radius_server"); ok || d.HasChange("captive_portal_radius_server") {
		t, err := expandWirelessControllerVapCaptivePortalRadiusServer(d, v, "captive_portal_radius_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-radius-server"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_session_timeout_interval"); ok || d.HasChange("captive_portal_session_timeout_interval") {
		t, err := expandWirelessControllerVapCaptivePortalSessionTimeoutInterval(d, v, "captive_portal_session_timeout_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-session-timeout-interval"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_auth_timeout"); ok || d.HasChange("captive_portal_auth_timeout") {
		t, err := expandWirelessControllerVapCaptivePortalAuthTimeout(d, v, "captive_portal_auth_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-auth-timeout"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_fw_accounting"); ok || d.HasChange("captive_portal_fw_accounting") {
		t, err := expandWirelessControllerVapCaptivePortalFwAccounting(d, v, "captive_portal_fw_accounting")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-fw-accounting"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_address_enforcement"); ok || d.HasChange("dhcp_address_enforcement") {
		t, err := expandWirelessControllerVapDhcpAddressEnforcement(d, v, "dhcp_address_enforcement")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-address-enforcement"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_lease_time"); ok || d.HasChange("dhcp_lease_time") {
		t, err := expandWirelessControllerVapDhcpLeaseTime(d, v, "dhcp_lease_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-lease-time"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_option43_insertion"); ok || d.HasChange("dhcp_option43_insertion") {
		t, err := expandWirelessControllerVapDhcpOption43Insertion(d, v, "dhcp_option43_insertion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-option43-insertion"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_option82_circuit_id_insertion"); ok || d.HasChange("dhcp_option82_circuit_id_insertion") {
		t, err := expandWirelessControllerVapDhcpOption82CircuitIdInsertion(d, v, "dhcp_option82_circuit_id_insertion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-option82-circuit-id-insertion"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_option82_insertion"); ok || d.HasChange("dhcp_option82_insertion") {
		t, err := expandWirelessControllerVapDhcpOption82Insertion(d, v, "dhcp_option82_insertion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-option82-insertion"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_option82_remote_id_insertion"); ok || d.HasChange("dhcp_option82_remote_id_insertion") {
		t, err := expandWirelessControllerVapDhcpOption82RemoteIdInsertion(d, v, "dhcp_option82_remote_id_insertion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-option82-remote-id-insertion"] = t
		}
	}

	if v, ok := d.GetOk("domain_name_stripping"); ok || d.HasChange("domain_name_stripping") {
		t, err := expandWirelessControllerVapDomainNameStripping(d, v, "domain_name_stripping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-name-stripping"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_vlan"); ok || d.HasChange("dynamic_vlan") {
		t, err := expandWirelessControllerVapDynamicVlan(d, v, "dynamic_vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic-vlan"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandWirelessControllerVapDynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("eap_reauth"); ok || d.HasChange("eap_reauth") {
		t, err := expandWirelessControllerVapEapReauth(d, v, "eap_reauth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-reauth"] = t
		}
	}

	if v, ok := d.GetOk("eap_reauth_intv"); ok || d.HasChange("eap_reauth_intv") {
		t, err := expandWirelessControllerVapEapReauthIntv(d, v, "eap_reauth_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-reauth-intv"] = t
		}
	}

	if v, ok := d.GetOk("eapol_key_retries"); ok || d.HasChange("eapol_key_retries") {
		t, err := expandWirelessControllerVapEapolKeyRetries(d, v, "eapol_key_retries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-key-retries"] = t
		}
	}

	if v, ok := d.GetOk("encrypt"); ok || d.HasChange("encrypt") {
		t, err := expandWirelessControllerVapEncrypt(d, v, "encrypt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encrypt"] = t
		}
	}

	if v, ok := d.GetOk("external_fast_roaming"); ok || d.HasChange("external_fast_roaming") {
		t, err := expandWirelessControllerVapExternalFastRoaming(d, v, "external_fast_roaming")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-fast-roaming"] = t
		}
	}

	if v, ok := d.GetOk("external_logout"); ok || d.HasChange("external_logout") {
		t, err := expandWirelessControllerVapExternalLogout(d, v, "external_logout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-logout"] = t
		}
	}

	if v, ok := d.GetOk("external_pre_auth"); ok || d.HasChange("external_pre_auth") {
		t, err := expandWirelessControllerVapExternalPreAuth(d, v, "external_pre_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-pre-auth"] = t
		}
	}

	if v, ok := d.GetOk("external_web"); ok || d.HasChange("external_web") {
		t, err := expandWirelessControllerVapExternalWeb(d, v, "external_web")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-web"] = t
		}
	}

	if v, ok := d.GetOk("external_web_format"); ok || d.HasChange("external_web_format") {
		t, err := expandWirelessControllerVapExternalWebFormat(d, v, "external_web_format")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-web-format"] = t
		}
	}

	if v, ok := d.GetOk("fast_bss_transition"); ok || d.HasChange("fast_bss_transition") {
		t, err := expandWirelessControllerVapFastBssTransition(d, v, "fast_bss_transition")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fast-bss-transition"] = t
		}
	}

	if v, ok := d.GetOk("fast_roaming"); ok || d.HasChange("fast_roaming") {
		t, err := expandWirelessControllerVapFastRoaming(d, v, "fast_roaming")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fast-roaming"] = t
		}
	}

	if v, ok := d.GetOk("ft_mobility_domain"); ok || d.HasChange("ft_mobility_domain") {
		t, err := expandWirelessControllerVapFtMobilityDomain(d, v, "ft_mobility_domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ft-mobility-domain"] = t
		}
	}

	if v, ok := d.GetOk("ft_over_ds"); ok || d.HasChange("ft_over_ds") {
		t, err := expandWirelessControllerVapFtOverDs(d, v, "ft_over_ds")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ft-over-ds"] = t
		}
	}

	if v, ok := d.GetOk("ft_r0_key_lifetime"); ok || d.HasChange("ft_r0_key_lifetime") {
		t, err := expandWirelessControllerVapFtR0KeyLifetime(d, v, "ft_r0_key_lifetime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ft-r0-key-lifetime"] = t
		}
	}

	if v, ok := d.GetOk("gas_comeback_delay"); ok || d.HasChange("gas_comeback_delay") {
		t, err := expandWirelessControllerVapGasComebackDelay(d, v, "gas_comeback_delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gas-comeback-delay"] = t
		}
	}

	if v, ok := d.GetOk("gas_fragmentation_limit"); ok || d.HasChange("gas_fragmentation_limit") {
		t, err := expandWirelessControllerVapGasFragmentationLimit(d, v, "gas_fragmentation_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gas-fragmentation-limit"] = t
		}
	}

	if v, ok := d.GetOk("gtk_rekey"); ok || d.HasChange("gtk_rekey") {
		t, err := expandWirelessControllerVapGtkRekey(d, v, "gtk_rekey")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtk-rekey"] = t
		}
	}

	if v, ok := d.GetOk("gtk_rekey_intv"); ok || d.HasChange("gtk_rekey_intv") {
		t, err := expandWirelessControllerVapGtkRekeyIntv(d, v, "gtk_rekey_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtk-rekey-intv"] = t
		}
	}

	if v, ok := d.GetOk("high_efficiency"); ok || d.HasChange("high_efficiency") {
		t, err := expandWirelessControllerVapHighEfficiency(d, v, "high_efficiency")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["high-efficiency"] = t
		}
	}

	if v, ok := d.GetOk("hotspot20_profile"); ok || d.HasChange("hotspot20_profile") {
		t, err := expandWirelessControllerVapHotspot20Profile(d, v, "hotspot20_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hotspot20-profile"] = t
		}
	}

	if v, ok := d.GetOk("igmp_snooping"); ok || d.HasChange("igmp_snooping") {
		t, err := expandWirelessControllerVapIgmpSnooping(d, v, "igmp_snooping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["igmp-snooping"] = t
		}
	}

	if v, ok := d.GetOk("intra_vap_privacy"); ok || d.HasChange("intra_vap_privacy") {
		t, err := expandWirelessControllerVapIntraVapPrivacy(d, v, "intra_vap_privacy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intra-vap-privacy"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandWirelessControllerVapIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("ips_sensor"); ok || d.HasChange("ips_sensor") {
		t, err := expandWirelessControllerVapIpsSensor(d, v, "ips_sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_rules"); ok || d.HasChange("ipv6_rules") {
		t, err := expandWirelessControllerVapIpv6Rules(d, v, "ipv6_rules")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-rules"] = t
		}
	}

	if v, ok := d.GetOk("key"); ok || d.HasChange("key") {
		t, err := expandWirelessControllerVapKey(d, v, "key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key"] = t
		}
	}

	if v, ok := d.GetOk("keyindex"); ok || d.HasChange("keyindex") {
		t, err := expandWirelessControllerVapKeyindex(d, v, "keyindex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keyindex"] = t
		}
	}

	if v, ok := d.GetOk("l3_roaming"); ok || d.HasChange("l3_roaming") {
		t, err := expandWirelessControllerVapL3Roaming(d, v, "l3_roaming")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l3-roaming"] = t
		}
	}

	if v, ok := d.GetOk("l3_roaming_mode"); ok || d.HasChange("l3_roaming_mode") {
		t, err := expandWirelessControllerVapL3RoamingMode(d, v, "l3_roaming_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l3-roaming-mode"] = t
		}
	}

	if v, ok := d.GetOk("ldpc"); ok || d.HasChange("ldpc") {
		t, err := expandWirelessControllerVapLdpc(d, v, "ldpc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldpc"] = t
		}
	}

	if v, ok := d.GetOk("local_authentication"); ok || d.HasChange("local_authentication") {
		t, err := expandWirelessControllerVapLocalAuthentication(d, v, "local_authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-authentication"] = t
		}
	}

	if v, ok := d.GetOk("local_bridging"); ok || d.HasChange("local_bridging") {
		t, err := expandWirelessControllerVapLocalBridging(d, v, "local_bridging")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-bridging"] = t
		}
	}

	if v, ok := d.GetOk("local_lan"); ok || d.HasChange("local_lan") {
		t, err := expandWirelessControllerVapLocalLan(d, v, "local_lan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-lan"] = t
		}
	}

	if v, ok := d.GetOk("local_lan_partition"); ok || d.HasChange("local_lan_partition") {
		t, err := expandWirelessControllerVapLocalLanPartition(d, v, "local_lan_partition")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-lan-partition"] = t
		}
	}

	if v, ok := d.GetOk("local_standalone"); ok || d.HasChange("local_standalone") {
		t, err := expandWirelessControllerVapLocalStandalone(d, v, "local_standalone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-standalone"] = t
		}
	}

	if v, ok := d.GetOk("local_standalone_dns"); ok || d.HasChange("local_standalone_dns") {
		t, err := expandWirelessControllerVapLocalStandaloneDns(d, v, "local_standalone_dns")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-standalone-dns"] = t
		}
	}

	if v, ok := d.GetOk("local_standalone_dns_ip"); ok || d.HasChange("local_standalone_dns_ip") {
		t, err := expandWirelessControllerVapLocalStandaloneDnsIp(d, v, "local_standalone_dns_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-standalone-dns-ip"] = t
		}
	}

	if v, ok := d.GetOk("local_standalone_nat"); ok || d.HasChange("local_standalone_nat") {
		t, err := expandWirelessControllerVapLocalStandaloneNat(d, v, "local_standalone_nat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-standalone-nat"] = t
		}
	}

	if v, ok := d.GetOk("mac_auth_bypass"); ok || d.HasChange("mac_auth_bypass") {
		t, err := expandWirelessControllerVapMacAuthBypass(d, v, "mac_auth_bypass")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-auth-bypass"] = t
		}
	}

	if v, ok := d.GetOk("mac_called_station_delimiter"); ok || d.HasChange("mac_called_station_delimiter") {
		t, err := expandWirelessControllerVapMacCalledStationDelimiter(d, v, "mac_called_station_delimiter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-called-station-delimiter"] = t
		}
	}

	if v, ok := d.GetOk("mac_calling_station_delimiter"); ok || d.HasChange("mac_calling_station_delimiter") {
		t, err := expandWirelessControllerVapMacCallingStationDelimiter(d, v, "mac_calling_station_delimiter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-calling-station-delimiter"] = t
		}
	}

	if v, ok := d.GetOk("mac_case"); ok || d.HasChange("mac_case") {
		t, err := expandWirelessControllerVapMacCase(d, v, "mac_case")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-case"] = t
		}
	}

	if v, ok := d.GetOk("mac_filter"); ok || d.HasChange("mac_filter") {
		t, err := expandWirelessControllerVapMacFilter(d, v, "mac_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-filter"] = t
		}
	}

	if v, ok := d.GetOk("mac_filter_list"); ok || d.HasChange("mac_filter_list") {
		t, err := expandWirelessControllerVapMacFilterList(d, v, "mac_filter_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-filter-list"] = t
		}
	}

	if v, ok := d.GetOk("mac_filter_policy_other"); ok || d.HasChange("mac_filter_policy_other") {
		t, err := expandWirelessControllerVapMacFilterPolicyOther(d, v, "mac_filter_policy_other")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-filter-policy-other"] = t
		}
	}

	if v, ok := d.GetOk("mac_password_delimiter"); ok || d.HasChange("mac_password_delimiter") {
		t, err := expandWirelessControllerVapMacPasswordDelimiter(d, v, "mac_password_delimiter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-password-delimiter"] = t
		}
	}

	if v, ok := d.GetOk("mac_username_delimiter"); ok || d.HasChange("mac_username_delimiter") {
		t, err := expandWirelessControllerVapMacUsernameDelimiter(d, v, "mac_username_delimiter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-username-delimiter"] = t
		}
	}

	if v, ok := d.GetOk("max_clients"); ok || d.HasChange("max_clients") {
		t, err := expandWirelessControllerVapMaxClients(d, v, "max_clients")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-clients"] = t
		}
	}

	if v, ok := d.GetOk("max_clients_ap"); ok || d.HasChange("max_clients_ap") {
		t, err := expandWirelessControllerVapMaxClientsAp(d, v, "max_clients_ap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-clients-ap"] = t
		}
	}

	if v, ok := d.GetOk("mbo"); ok || d.HasChange("mbo") {
		t, err := expandWirelessControllerVapMbo(d, v, "mbo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mbo"] = t
		}
	}

	if v, ok := d.GetOk("mbo_cell_data_conn_pref"); ok || d.HasChange("mbo_cell_data_conn_pref") {
		t, err := expandWirelessControllerVapMboCellDataConnPref(d, v, "mbo_cell_data_conn_pref")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mbo-cell-data-conn-pref"] = t
		}
	}

	if v, ok := d.GetOk("me_disable_thresh"); ok || d.HasChange("me_disable_thresh") {
		t, err := expandWirelessControllerVapMeDisableThresh(d, v, "me_disable_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["me-disable-thresh"] = t
		}
	}

	if v, ok := d.GetOk("mesh_backhaul"); ok || d.HasChange("mesh_backhaul") {
		t, err := expandWirelessControllerVapMeshBackhaul(d, v, "mesh_backhaul")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mesh-backhaul"] = t
		}
	}

	if v, ok := d.GetOk("mpsk"); ok || d.HasChange("mpsk") {
		t, err := expandWirelessControllerVapMpsk(d, v, "mpsk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk"] = t
		}
	}

	if v, ok := d.GetOk("mpsk_concurrent_clients"); ok || d.HasChange("mpsk_concurrent_clients") {
		t, err := expandWirelessControllerVapMpskConcurrentClients(d, v, "mpsk_concurrent_clients")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk-concurrent-clients"] = t
		}
	}

	if v, ok := d.GetOk("mpsk_key"); ok || d.HasChange("mpsk_key") {
		t, err := expandWirelessControllerVapMpskKey(d, v, "mpsk_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk-key"] = t
		}
	}

	if v, ok := d.GetOk("mpsk_profile"); ok || d.HasChange("mpsk_profile") {
		t, err := expandWirelessControllerVapMpskProfile(d, v, "mpsk_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk-profile"] = t
		}
	}

	if v, ok := d.GetOk("mu_mimo"); ok || d.HasChange("mu_mimo") {
		t, err := expandWirelessControllerVapMuMimo(d, v, "mu_mimo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mu-mimo"] = t
		}
	}

	if v, ok := d.GetOk("multicast_enhance"); ok || d.HasChange("multicast_enhance") {
		t, err := expandWirelessControllerVapMulticastEnhance(d, v, "multicast_enhance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-enhance"] = t
		}
	}

	if v, ok := d.GetOk("multicast_rate"); ok || d.HasChange("multicast_rate") {
		t, err := expandWirelessControllerVapMulticastRate(d, v, "multicast_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-rate"] = t
		}
	}

	if v, ok := d.GetOk("nac"); ok || d.HasChange("nac") {
		t, err := expandWirelessControllerVapNac(d, v, "nac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac"] = t
		}
	}

	if v, ok := d.GetOk("nac_profile"); ok || d.HasChange("nac_profile") {
		t, err := expandWirelessControllerVapNacProfile(d, v, "nac_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac-profile"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerVapName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nas_filter_rule"); ok || d.HasChange("nas_filter_rule") {
		t, err := expandWirelessControllerVapNasFilterRule(d, v, "nas_filter_rule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nas-filter-rule"] = t
		}
	}

	if v, ok := d.GetOk("neighbor_report_dual_band"); ok || d.HasChange("neighbor_report_dual_band") {
		t, err := expandWirelessControllerVapNeighborReportDualBand(d, v, "neighbor_report_dual_band")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-report-dual-band"] = t
		}
	}

	if v, ok := d.GetOk("okc"); ok || d.HasChange("okc") {
		t, err := expandWirelessControllerVapOkc(d, v, "okc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["okc"] = t
		}
	}

	if v, ok := d.GetOk("osen"); ok || d.HasChange("osen") {
		t, err := expandWirelessControllerVapOsen(d, v, "osen")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["osen"] = t
		}
	}

	if v, ok := d.GetOk("owe_groups"); ok || d.HasChange("owe_groups") {
		t, err := expandWirelessControllerVapOweGroups(d, v, "owe_groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["owe-groups"] = t
		}
	}

	if v, ok := d.GetOk("owe_transition"); ok || d.HasChange("owe_transition") {
		t, err := expandWirelessControllerVapOweTransition(d, v, "owe_transition")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["owe-transition"] = t
		}
	}

	if v, ok := d.GetOk("owe_transition_ssid"); ok || d.HasChange("owe_transition_ssid") {
		t, err := expandWirelessControllerVapOweTransitionSsid(d, v, "owe_transition_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["owe-transition-ssid"] = t
		}
	}

	if v, ok := d.GetOk("passphrase"); ok || d.HasChange("passphrase") {
		t, err := expandWirelessControllerVapPassphrase(d, v, "passphrase")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passphrase"] = t
		}
	}

	if v, ok := d.GetOk("pmf"); ok || d.HasChange("pmf") {
		t, err := expandWirelessControllerVapPmf(d, v, "pmf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pmf"] = t
		}
	}

	if v, ok := d.GetOk("pmf_assoc_comeback_timeout"); ok || d.HasChange("pmf_assoc_comeback_timeout") {
		t, err := expandWirelessControllerVapPmfAssocComebackTimeout(d, v, "pmf_assoc_comeback_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pmf-assoc-comeback-timeout"] = t
		}
	}

	if v, ok := d.GetOk("pmf_sa_query_retry_timeout"); ok || d.HasChange("pmf_sa_query_retry_timeout") {
		t, err := expandWirelessControllerVapPmfSaQueryRetryTimeout(d, v, "pmf_sa_query_retry_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pmf-sa-query-retry-timeout"] = t
		}
	}

	if v, ok := d.GetOk("port_macauth"); ok || d.HasChange("port_macauth") {
		t, err := expandWirelessControllerVapPortMacauth(d, v, "port_macauth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-macauth"] = t
		}
	}

	if v, ok := d.GetOk("port_macauth_reauth_timeout"); ok || d.HasChange("port_macauth_reauth_timeout") {
		t, err := expandWirelessControllerVapPortMacauthReauthTimeout(d, v, "port_macauth_reauth_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-macauth-reauth-timeout"] = t
		}
	}

	if v, ok := d.GetOk("port_macauth_timeout"); ok || d.HasChange("port_macauth_timeout") {
		t, err := expandWirelessControllerVapPortMacauthTimeout(d, v, "port_macauth_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-macauth-timeout"] = t
		}
	}

	if v, ok := d.GetOk("portal_message_override_group"); ok || d.HasChange("portal_message_override_group") {
		t, err := expandWirelessControllerVapPortalMessageOverrideGroup(d, v, "portal_message_override_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portal-message-override-group"] = t
		}
	}

	if v, ok := d.GetOk("portal_message_overrides"); ok || d.HasChange("portal_message_overrides") {
		t, err := expandWirelessControllerVapPortalMessageOverrides(d, v, "portal_message_overrides")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portal-message-overrides"] = t
		}
	}

	if v, ok := d.GetOk("portal_type"); ok || d.HasChange("portal_type") {
		t, err := expandWirelessControllerVapPortalType(d, v, "portal_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portal-type"] = t
		}
	}

	if v, ok := d.GetOk("pre_auth"); ok || d.HasChange("pre_auth") {
		t, err := expandWirelessControllerVapPreAuth(d, v, "pre_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pre-auth"] = t
		}
	}

	if v, ok := d.GetOk("primary_wag_profile"); ok || d.HasChange("primary_wag_profile") {
		t, err := expandWirelessControllerVapPrimaryWagProfile(d, v, "primary_wag_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary-wag-profile"] = t
		}
	}

	if v, ok := d.GetOk("probe_resp_suppression"); ok || d.HasChange("probe_resp_suppression") {
		t, err := expandWirelessControllerVapProbeRespSuppression(d, v, "probe_resp_suppression")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-resp-suppression"] = t
		}
	}

	if v, ok := d.GetOk("probe_resp_threshold"); ok || d.HasChange("probe_resp_threshold") {
		t, err := expandWirelessControllerVapProbeRespThreshold(d, v, "probe_resp_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-resp-threshold"] = t
		}
	}

	if v, ok := d.GetOk("ptk_rekey"); ok || d.HasChange("ptk_rekey") {
		t, err := expandWirelessControllerVapPtkRekey(d, v, "ptk_rekey")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ptk-rekey"] = t
		}
	}

	if v, ok := d.GetOk("ptk_rekey_intv"); ok || d.HasChange("ptk_rekey_intv") {
		t, err := expandWirelessControllerVapPtkRekeyIntv(d, v, "ptk_rekey_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ptk-rekey-intv"] = t
		}
	}

	if v, ok := d.GetOk("qos_profile"); ok || d.HasChange("qos_profile") {
		t, err := expandWirelessControllerVapQosProfile(d, v, "qos_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qos-profile"] = t
		}
	}

	if v, ok := d.GetOk("quarantine"); ok || d.HasChange("quarantine") {
		t, err := expandWirelessControllerVapQuarantine(d, v, "quarantine")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine"] = t
		}
	}

	if v, ok := d.GetOk("radio_2g_threshold"); ok || d.HasChange("radio_2g_threshold") {
		t, err := expandWirelessControllerVapRadio2GThreshold(d, v, "radio_2g_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-2g-threshold"] = t
		}
	}

	if v, ok := d.GetOk("radio_5g_threshold"); ok || d.HasChange("radio_5g_threshold") {
		t, err := expandWirelessControllerVapRadio5GThreshold(d, v, "radio_5g_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-5g-threshold"] = t
		}
	}

	if v, ok := d.GetOk("radio_sensitivity"); ok || d.HasChange("radio_sensitivity") {
		t, err := expandWirelessControllerVapRadioSensitivity(d, v, "radio_sensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-sensitivity"] = t
		}
	}

	if v, ok := d.GetOk("radius_mac_auth"); ok || d.HasChange("radius_mac_auth") {
		t, err := expandWirelessControllerVapRadiusMacAuth(d, v, "radius_mac_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-auth"] = t
		}
	}

	if v, ok := d.GetOk("radius_mac_auth_block_interval"); ok || d.HasChange("radius_mac_auth_block_interval") {
		t, err := expandWirelessControllerVapRadiusMacAuthBlockInterval(d, v, "radius_mac_auth_block_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-auth-block-interval"] = t
		}
	}

	if v, ok := d.GetOk("radius_mac_auth_server"); ok || d.HasChange("radius_mac_auth_server") {
		t, err := expandWirelessControllerVapRadiusMacAuthServer(d, v, "radius_mac_auth_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-auth-server"] = t
		}
	}

	if v, ok := d.GetOk("radius_mac_auth_usergroups"); ok || d.HasChange("radius_mac_auth_usergroups") {
		t, err := expandWirelessControllerVapRadiusMacAuthUsergroups(d, v, "radius_mac_auth_usergroups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-auth-usergroups"] = t
		}
	}

	if v, ok := d.GetOk("radius_mac_mpsk_auth"); ok || d.HasChange("radius_mac_mpsk_auth") {
		t, err := expandWirelessControllerVapRadiusMacMpskAuth(d, v, "radius_mac_mpsk_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-mpsk-auth"] = t
		}
	}

	if v, ok := d.GetOk("radius_mac_mpsk_timeout"); ok || d.HasChange("radius_mac_mpsk_timeout") {
		t, err := expandWirelessControllerVapRadiusMacMpskTimeout(d, v, "radius_mac_mpsk_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-mpsk-timeout"] = t
		}
	}

	if v, ok := d.GetOk("radius_server"); ok || d.HasChange("radius_server") {
		t, err := expandWirelessControllerVapRadiusServer(d, v, "radius_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-server"] = t
		}
	}

	if v, ok := d.GetOk("rates_11a"); ok || d.HasChange("rates_11a") {
		t, err := expandWirelessControllerVapRates11A(d, v, "rates_11a")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11a"] = t
		}
	}

	if v, ok := d.GetOk("rates_11ac_mcs_map"); ok || d.HasChange("rates_11ac_mcs_map") {
		t, err := expandWirelessControllerVapRates11AcMcsMap(d, v, "rates_11ac_mcs_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11ac-mcs-map"] = t
		}
	}

	if v, ok := d.GetOk("rates_11ac_ss12"); ok || d.HasChange("rates_11ac_ss12") {
		t, err := expandWirelessControllerVapRates11AcSs12(d, v, "rates_11ac_ss12")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11ac-ss12"] = t
		}
	}

	if v, ok := d.GetOk("rates_11ac_ss34"); ok || d.HasChange("rates_11ac_ss34") {
		t, err := expandWirelessControllerVapRates11AcSs34(d, v, "rates_11ac_ss34")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11ac-ss34"] = t
		}
	}

	if v, ok := d.GetOk("rates_11ax_mcs_map"); ok || d.HasChange("rates_11ax_mcs_map") {
		t, err := expandWirelessControllerVapRates11AxMcsMap(d, v, "rates_11ax_mcs_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11ax-mcs-map"] = t
		}
	}

	if v, ok := d.GetOk("rates_11ax_ss12"); ok || d.HasChange("rates_11ax_ss12") {
		t, err := expandWirelessControllerVapRates11AxSs12(d, v, "rates_11ax_ss12")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11ax-ss12"] = t
		}
	}

	if v, ok := d.GetOk("rates_11ax_ss34"); ok || d.HasChange("rates_11ax_ss34") {
		t, err := expandWirelessControllerVapRates11AxSs34(d, v, "rates_11ax_ss34")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11ax-ss34"] = t
		}
	}

	if v, ok := d.GetOk("rates_11be_mcs_map"); ok || d.HasChange("rates_11be_mcs_map") {
		t, err := expandWirelessControllerVapRates11BeMcsMap(d, v, "rates_11be_mcs_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11be-mcs-map"] = t
		}
	}

	if v, ok := d.GetOk("rates_11be_mcs_map_160"); ok || d.HasChange("rates_11be_mcs_map_160") {
		t, err := expandWirelessControllerVapRates11BeMcsMap160(d, v, "rates_11be_mcs_map_160")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11be-mcs-map-160"] = t
		}
	}

	if v, ok := d.GetOk("rates_11be_mcs_map_320"); ok || d.HasChange("rates_11be_mcs_map_320") {
		t, err := expandWirelessControllerVapRates11BeMcsMap320(d, v, "rates_11be_mcs_map_320")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11be-mcs-map-320"] = t
		}
	}

	if v, ok := d.GetOk("rates_11bg"); ok || d.HasChange("rates_11bg") {
		t, err := expandWirelessControllerVapRates11Bg(d, v, "rates_11bg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11bg"] = t
		}
	}

	if v, ok := d.GetOk("rates_11n_ss12"); ok || d.HasChange("rates_11n_ss12") {
		t, err := expandWirelessControllerVapRates11NSs12(d, v, "rates_11n_ss12")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11n-ss12"] = t
		}
	}

	if v, ok := d.GetOk("rates_11n_ss34"); ok || d.HasChange("rates_11n_ss34") {
		t, err := expandWirelessControllerVapRates11NSs34(d, v, "rates_11n_ss34")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11n-ss34"] = t
		}
	}

	if v, ok := d.GetOk("roaming_acct_interim_update"); ok || d.HasChange("roaming_acct_interim_update") {
		t, err := expandWirelessControllerVapRoamingAcctInterimUpdate(d, v, "roaming_acct_interim_update")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["roaming-acct-interim-update"] = t
		}
	}

	if v, ok := d.GetOk("sae_groups"); ok || d.HasChange("sae_groups") {
		t, err := expandWirelessControllerVapSaeGroups(d, v, "sae_groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sae-groups"] = t
		}
	}

	if v, ok := d.GetOk("sae_h2e_only"); ok || d.HasChange("sae_h2e_only") {
		t, err := expandWirelessControllerVapSaeH2EOnly(d, v, "sae_h2e_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sae-h2e-only"] = t
		}
	}

	if v, ok := d.GetOk("sae_hnp_only"); ok || d.HasChange("sae_hnp_only") {
		t, err := expandWirelessControllerVapSaeHnpOnly(d, v, "sae_hnp_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sae-hnp-only"] = t
		}
	}

	if v, ok := d.GetOk("sae_password"); ok || d.HasChange("sae_password") {
		t, err := expandWirelessControllerVapSaePassword(d, v, "sae_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sae-password"] = t
		}
	}

	if v, ok := d.GetOk("sae_pk"); ok || d.HasChange("sae_pk") {
		t, err := expandWirelessControllerVapSaePk(d, v, "sae_pk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sae-pk"] = t
		}
	}

	if v, ok := d.GetOk("sae_private_key"); ok || d.HasChange("sae_private_key") {
		t, err := expandWirelessControllerVapSaePrivateKey(d, v, "sae_private_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sae-private-key"] = t
		}
	}

	if v, ok := d.GetOk("scan_botnet_connections"); ok || d.HasChange("scan_botnet_connections") {
		t, err := expandWirelessControllerVapScanBotnetConnections(d, v, "scan_botnet_connections")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-botnet-connections"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok || d.HasChange("schedule") {
		t, err := expandWirelessControllerVapSchedule(d, v, "schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("secondary_wag_profile"); ok || d.HasChange("secondary_wag_profile") {
		t, err := expandWirelessControllerVapSecondaryWagProfile(d, v, "secondary_wag_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-wag-profile"] = t
		}
	}

	if v, ok := d.GetOk("security"); ok || d.HasChange("security") {
		t, err := expandWirelessControllerVapSecurity(d, v, "security")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security"] = t
		}
	}

	if v, ok := d.GetOk("security_exempt_list"); ok || d.HasChange("security_exempt_list") {
		t, err := expandWirelessControllerVapSecurityExemptList(d, v, "security_exempt_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-exempt-list"] = t
		}
	}

	if v, ok := d.GetOk("security_obsolete_option"); ok || d.HasChange("security_obsolete_option") {
		t, err := expandWirelessControllerVapSecurityObsoleteOption(d, v, "security_obsolete_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-obsolete-option"] = t
		}
	}

	if v, ok := d.GetOk("security_redirect_url"); ok || d.HasChange("security_redirect_url") {
		t, err := expandWirelessControllerVapSecurityRedirectUrl(d, v, "security_redirect_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-redirect-url"] = t
		}
	}

	if v, ok := d.GetOk("selected_usergroups"); ok || d.HasChange("selected_usergroups") {
		t, err := expandWirelessControllerVapSelectedUsergroups(d, v, "selected_usergroups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["selected-usergroups"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling"); ok || d.HasChange("split_tunneling") {
		t, err := expandWirelessControllerVapSplitTunneling(d, v, "split_tunneling")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling"] = t
		}
	}

	if v, ok := d.GetOk("ssid"); ok || d.HasChange("ssid") {
		t, err := expandWirelessControllerVapSsid(d, v, "ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssid"] = t
		}
	}

	if v, ok := d.GetOk("sticky_client_remove"); ok || d.HasChange("sticky_client_remove") {
		t, err := expandWirelessControllerVapStickyClientRemove(d, v, "sticky_client_remove")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sticky-client-remove"] = t
		}
	}

	if v, ok := d.GetOk("sticky_client_threshold_2g"); ok || d.HasChange("sticky_client_threshold_2g") {
		t, err := expandWirelessControllerVapStickyClientThreshold2G(d, v, "sticky_client_threshold_2g")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sticky-client-threshold-2g"] = t
		}
	}

	if v, ok := d.GetOk("sticky_client_threshold_5g"); ok || d.HasChange("sticky_client_threshold_5g") {
		t, err := expandWirelessControllerVapStickyClientThreshold5G(d, v, "sticky_client_threshold_5g")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sticky-client-threshold-5g"] = t
		}
	}

	if v, ok := d.GetOk("sticky_client_threshold_6g"); ok || d.HasChange("sticky_client_threshold_6g") {
		t, err := expandWirelessControllerVapStickyClientThreshold6G(d, v, "sticky_client_threshold_6g")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sticky-client-threshold-6g"] = t
		}
	}

	if v, ok := d.GetOk("target_wake_time"); ok || d.HasChange("target_wake_time") {
		t, err := expandWirelessControllerVapTargetWakeTime(d, v, "target_wake_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target-wake-time"] = t
		}
	}

	if v, ok := d.GetOk("tkip_counter_measure"); ok || d.HasChange("tkip_counter_measure") {
		t, err := expandWirelessControllerVapTkipCounterMeasure(d, v, "tkip_counter_measure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tkip-counter-measure"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_echo_interval"); ok || d.HasChange("tunnel_echo_interval") {
		t, err := expandWirelessControllerVapTunnelEchoInterval(d, v, "tunnel_echo_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-echo-interval"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_fallback_interval"); ok || d.HasChange("tunnel_fallback_interval") {
		t, err := expandWirelessControllerVapTunnelFallbackInterval(d, v, "tunnel_fallback_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-fallback-interval"] = t
		}
	}

	if v, ok := d.GetOk("usergroup"); ok || d.HasChange("usergroup") {
		t, err := expandWirelessControllerVapUsergroup(d, v, "usergroup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["usergroup"] = t
		}
	}

	if v, ok := d.GetOk("utm_log"); ok || d.HasChange("utm_log") {
		t, err := expandWirelessControllerVapUtmLog(d, v, "utm_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm-log"] = t
		}
	}

	if v, ok := d.GetOk("utm_profile"); ok || d.HasChange("utm_profile") {
		t, err := expandWirelessControllerVapUtmProfile(d, v, "utm_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm-profile"] = t
		}
	}

	if v, ok := d.GetOk("utm_status"); ok || d.HasChange("utm_status") {
		t, err := expandWirelessControllerVapUtmStatus(d, v, "utm_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm-status"] = t
		}
	}

	if v, ok := d.GetOk("vlan_auto"); ok || d.HasChange("vlan_auto") {
		t, err := expandWirelessControllerVapVlanAuto(d, v, "vlan_auto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-auto"] = t
		}
	}

	if v, ok := d.GetOk("vlan_name"); ok || d.HasChange("vlan_name") {
		t, err := expandWirelessControllerVapVlanName(d, v, "vlan_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-name"] = t
		}
	}

	if v, ok := d.GetOk("vlan_pool"); ok || d.HasChange("vlan_pool") {
		t, err := expandWirelessControllerVapVlanPool(d, v, "vlan_pool")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-pool"] = t
		}
	}

	if v, ok := d.GetOk("vlan_pooling"); ok || d.HasChange("vlan_pooling") {
		t, err := expandWirelessControllerVapVlanPooling(d, v, "vlan_pooling")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-pooling"] = t
		}
	}

	if v, ok := d.GetOk("vlanid"); ok || d.HasChange("vlanid") {
		t, err := expandWirelessControllerVapVlanid(d, v, "vlanid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlanid"] = t
		}
	}

	if v, ok := d.GetOk("voice_enterprise"); ok || d.HasChange("voice_enterprise") {
		t, err := expandWirelessControllerVapVoiceEnterprise(d, v, "voice_enterprise")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voice-enterprise"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_profile"); ok || d.HasChange("webfilter_profile") {
		t, err := expandWirelessControllerVapWebfilterProfile(d, v, "webfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile"] = t
		}
	}

	return &obj, nil
}
