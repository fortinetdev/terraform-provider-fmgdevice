// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch devices that are managed by this FortiGate.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerManagedSwitch() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerManagedSwitchCreate,
		Read:   resourceSwitchControllerManagedSwitchRead,
		Update: resourceSwitchControllerManagedSwitchUpdate,
		Delete: resourceSwitchControllerManagedSwitchDelete,

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
			"n802_1x_settings": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"link_down_auth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"local_override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mab_reauth": &schema.Schema{
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
						"max_reauth_attempt": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"reauth_period": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"tx_period": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"_platform": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"access_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"custom_command": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"command_entry": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"command_name": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"delayed_restart_trigger": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp_server_access_list": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_snooping_static_client": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vlan": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"directly_connected": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dynamic_capability": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamically_discovered": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"firmware_provision": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"firmware_provision_latest": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"firmware_provision_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"flow_identity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fsw_wan1_admin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fsw_wan1_peer": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fsw_wan2_admin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fsw_wan2_peer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"igmp_snooping": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aging_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"flood_unknown_multicast": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"local_override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlans": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"proxy": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"querier": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"querier_addr": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"version": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"vlan_name": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"ip_source_guard": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"binding_entry": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"entry_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mac": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"l3_discovered": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_allowed_trunk_members": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mclag_igmp_snooping_aware": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mgmt_mode": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mirror": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dst": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"src_egress": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"src_ingress": &schema.Schema{
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
						"switching_packet": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"override_snmp_community": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_snmp_sysinfo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_snmp_trap_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_snmp_user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"owner_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"poe_detection_type": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"poe_lldp_detection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"poe_pre_standard_detection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ports": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"acl_group": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"aggregator_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"allow_arp_monitor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"allowed_vlans": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"allowed_vlans_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"arp_inspection_trust": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"authenticated_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"bundle": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dhcp_snoop_option82_override": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"circuit_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"remote_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"vlan_name": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"dhcp_snoop_option82_trust": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dhcp_snooping": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"discard_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dsl_profile": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"edge_port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"export_tags": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"encrypted_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"export_to": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"export_to_pool": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"export_to_pool_flag": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"fallback_port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fec_capable": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"fec_state": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fgt_peer_device_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fgt_peer_port_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fiber_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"flags": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"flap_duration": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"flap_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"flap_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"flapguard": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"flow_control": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortilink_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"igmp_snooping": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortiswitch_acls": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"igmp_snooping_flood_reports": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"igmps_flood_reports": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"igmps_flood_traffic": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"interface_tags": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"ip_source_guard": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"isl_local_trunk_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"isl_peer_device_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"isl_peer_device_sn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"isl_peer_port_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"lacp_speed": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"learning_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"link_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"lldp_profile": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"lldp_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"loop_guard": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"loop_guard_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mac_addr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"matched_dpp_intf_tags": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"matched_dpp_policy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"max_bundle": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mcast_snooping_flood_traffic": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mclag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mclag_icl_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"media_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"member_withdrawal_behavior": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"members": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"min_bundle": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"p2p_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"packet_sample_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"packet_sampler": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pause_meter": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"pause_meter_resume": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"poe_capable": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"poe_max_power": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"poe_mode_bt_cabable": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"poe_port_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"poe_port_power": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"poe_port_priority": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"poe_pre_standard_detection": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"poe_standard": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"poe_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_number": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"port_owner": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_policy": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"port_prefix_type": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"port_security_policy": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"port_selection_criteria": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ptp_policy": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"ptp_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"qos_policy": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"restricted_auth_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"rpvst_port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sample_direction": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sflow_counter_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"speed": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"speed_mask": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"stacking_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sticky_mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"storm_control_policy": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"stp_bpdu_guard": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"stp_bpdu_guard_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"stp_root_guard": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"stp_state": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"switch_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"trunk_member": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"untagged_vlans": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"virtual_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"vlan": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"pre_provisioned": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ptp_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ptp_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"purdue_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"qos_drop_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"qos_red_probability": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"radius_nas_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"radius_nas_ip_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_log": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"csv": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"facility": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"route_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_offload_mclag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_offload_router": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"router_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan_name": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"sn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"snmp_community": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"events": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"hosts": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ip": &schema.Schema{
										Type:     schema.TypeList,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"query_v1_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"query_v1_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"query_v2c_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"query_v2c_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"trap_v1_lport": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"trap_v1_rport": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"trap_v1_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"trap_v2c_lport": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"trap_v2c_rport": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"trap_v2c_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"snmp_sysinfo": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"contact_info": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"engine_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"location": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"snmp_trap_threshold": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"trap_high_cpu_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"trap_log_full_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"trap_low_memory_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"snmp_user": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_proto": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auth_pwd": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"priv_proto": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"priv_pwd": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"queries": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"query_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"security_level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"staged_image_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"static_mac": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"storm_control": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"broadcast": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"local_override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"unknown_multicast": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unknown_unicast": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"stp_instance": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"stp_settings": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"forward_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"hello_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"local_override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_age": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"max_hops": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pending_timer": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"revision": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"switch_device_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"switch_dhcp_opt43_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"switch_id": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"switch_log": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"local_override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"switch_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"tdr_supported": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tunnel_discovered": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vlan": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"assignment_priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"vlan_name": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSwitchControllerManagedSwitchCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerManagedSwitch(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerManagedSwitch resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerManagedSwitch(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerManagedSwitch resource: %v", err)
	}

	d.SetId(getStringKey(d, "switch_id"))

	return resourceSwitchControllerManagedSwitchRead(d, m)
}

func resourceSwitchControllerManagedSwitchUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerManagedSwitch(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitch resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerManagedSwitch(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitch resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "switch_id"))

	return resourceSwitchControllerManagedSwitchRead(d, m)
}

func resourceSwitchControllerManagedSwitchDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerManagedSwitch(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerManagedSwitch resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerManagedSwitchRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerManagedSwitch(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitch resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerManagedSwitch(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitch resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerManagedSwitch8021XSettings(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "link_down_auth"
	if _, ok := i["link-down-auth"]; ok {
		result["link_down_auth"] = flattenSwitchControllerManagedSwitch8021XSettingsLinkDownAuth(i["link-down-auth"], d, pre_append)
	}

	pre_append = pre + ".0." + "local_override"
	if _, ok := i["local-override"]; ok {
		result["local_override"] = flattenSwitchControllerManagedSwitch8021XSettingsLocalOverride(i["local-override"], d, pre_append)
	}

	pre_append = pre + ".0." + "mab_reauth"
	if _, ok := i["mab-reauth"]; ok {
		result["mab_reauth"] = flattenSwitchControllerManagedSwitch8021XSettingsMabReauth(i["mab-reauth"], d, pre_append)
	}

	pre_append = pre + ".0." + "mac_called_station_delimiter"
	if _, ok := i["mac-called-station-delimiter"]; ok {
		result["mac_called_station_delimiter"] = flattenSwitchControllerManagedSwitch8021XSettingsMacCalledStationDelimiter(i["mac-called-station-delimiter"], d, pre_append)
	}

	pre_append = pre + ".0." + "mac_calling_station_delimiter"
	if _, ok := i["mac-calling-station-delimiter"]; ok {
		result["mac_calling_station_delimiter"] = flattenSwitchControllerManagedSwitch8021XSettingsMacCallingStationDelimiter(i["mac-calling-station-delimiter"], d, pre_append)
	}

	pre_append = pre + ".0." + "mac_case"
	if _, ok := i["mac-case"]; ok {
		result["mac_case"] = flattenSwitchControllerManagedSwitch8021XSettingsMacCase(i["mac-case"], d, pre_append)
	}

	pre_append = pre + ".0." + "mac_password_delimiter"
	if _, ok := i["mac-password-delimiter"]; ok {
		result["mac_password_delimiter"] = flattenSwitchControllerManagedSwitch8021XSettingsMacPasswordDelimiter(i["mac-password-delimiter"], d, pre_append)
	}

	pre_append = pre + ".0." + "mac_username_delimiter"
	if _, ok := i["mac-username-delimiter"]; ok {
		result["mac_username_delimiter"] = flattenSwitchControllerManagedSwitch8021XSettingsMacUsernameDelimiter(i["mac-username-delimiter"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_reauth_attempt"
	if _, ok := i["max-reauth-attempt"]; ok {
		result["max_reauth_attempt"] = flattenSwitchControllerManagedSwitch8021XSettingsMaxReauthAttempt(i["max-reauth-attempt"], d, pre_append)
	}

	pre_append = pre + ".0." + "reauth_period"
	if _, ok := i["reauth-period"]; ok {
		result["reauth_period"] = flattenSwitchControllerManagedSwitch8021XSettingsReauthPeriod(i["reauth-period"], d, pre_append)
	}

	pre_append = pre + ".0." + "tx_period"
	if _, ok := i["tx-period"]; ok {
		result["tx_period"] = flattenSwitchControllerManagedSwitch8021XSettingsTxPeriod(i["tx-period"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchControllerManagedSwitch8021XSettingsLinkDownAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettingsLocalOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettingsMabReauth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettingsMacCalledStationDelimiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettingsMacCallingStationDelimiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettingsMacCase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettingsMacPasswordDelimiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettingsMacUsernameDelimiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettingsMaxReauthAttempt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettingsReauthPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettingsTxPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPlatform(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchAccessProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchCustomCommand(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "command_entry"
		if _, ok := i["command-entry"]; ok {
			v := flattenSwitchControllerManagedSwitchCustomCommandCommandEntry(i["command-entry"], d, pre_append)
			tmp["command_entry"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-CustomCommand-CommandEntry")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "command_name"
		if _, ok := i["command-name"]; ok {
			v := flattenSwitchControllerManagedSwitchCustomCommandCommandName(i["command-name"], d, pre_append)
			tmp["command_name"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-CustomCommand-CommandName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerManagedSwitchCustomCommandCommandEntry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchCustomCommandCommandName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchDelayedRestartTrigger(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchDhcpServerAccessList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchDhcpSnoopingStaticClient(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenSwitchControllerManagedSwitchDhcpSnoopingStaticClientIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-DhcpSnoopingStaticClient-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			v := flattenSwitchControllerManagedSwitchDhcpSnoopingStaticClientMac(i["mac"], d, pre_append)
			tmp["mac"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-DhcpSnoopingStaticClient-Mac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSwitchControllerManagedSwitchDhcpSnoopingStaticClientName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-DhcpSnoopingStaticClient-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenSwitchControllerManagedSwitchDhcpSnoopingStaticClientPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-DhcpSnoopingStaticClient-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan"
		if _, ok := i["vlan"]; ok {
			v := flattenSwitchControllerManagedSwitchDhcpSnoopingStaticClientVlan(i["vlan"], d, pre_append)
			tmp["vlan"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-DhcpSnoopingStaticClient-Vlan")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerManagedSwitchDhcpSnoopingStaticClientIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchDhcpSnoopingStaticClientMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchDhcpSnoopingStaticClientName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchDhcpSnoopingStaticClientPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchDhcpSnoopingStaticClientVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchDirectlyConnected(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchDynamicCapability(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchDynamicallyDiscovered(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchFirmwareProvision(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchFirmwareProvisionLatest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchFirmwareProvisionVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchFlowIdentity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchFswWan1Admin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchFswWan1Peer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convstr2list(v, d.Get(pre))
}

func flattenSwitchControllerManagedSwitchFswWan2Admin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchFswWan2Peer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIgmpSnooping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "aging_time"
	if _, ok := i["aging-time"]; ok {
		result["aging_time"] = flattenSwitchControllerManagedSwitchIgmpSnoopingAgingTime(i["aging-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "flood_unknown_multicast"
	if _, ok := i["flood-unknown-multicast"]; ok {
		result["flood_unknown_multicast"] = flattenSwitchControllerManagedSwitchIgmpSnoopingFloodUnknownMulticast(i["flood-unknown-multicast"], d, pre_append)
	}

	pre_append = pre + ".0." + "local_override"
	if _, ok := i["local-override"]; ok {
		result["local_override"] = flattenSwitchControllerManagedSwitchIgmpSnoopingLocalOverride(i["local-override"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlans"
	if _, ok := i["vlans"]; ok {
		result["vlans"] = flattenSwitchControllerManagedSwitchIgmpSnoopingVlans(i["vlans"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingAgingTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingFloodUnknownMulticast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingLocalOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingVlans(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "proxy"
		if _, ok := i["proxy"]; ok {
			v := flattenSwitchControllerManagedSwitchIgmpSnoopingVlansProxy(i["proxy"], d, pre_append)
			tmp["proxy"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchIgmpSnooping-Vlans-Proxy")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "querier"
		if _, ok := i["querier"]; ok {
			v := flattenSwitchControllerManagedSwitchIgmpSnoopingVlansQuerier(i["querier"], d, pre_append)
			tmp["querier"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchIgmpSnooping-Vlans-Querier")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "querier_addr"
		if _, ok := i["querier-addr"]; ok {
			v := flattenSwitchControllerManagedSwitchIgmpSnoopingVlansQuerierAddr(i["querier-addr"], d, pre_append)
			tmp["querier_addr"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchIgmpSnooping-Vlans-QuerierAddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := i["version"]; ok {
			v := flattenSwitchControllerManagedSwitchIgmpSnoopingVlansVersion(i["version"], d, pre_append)
			tmp["version"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchIgmpSnooping-Vlans-Version")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if _, ok := i["vlan-name"]; ok {
			v := flattenSwitchControllerManagedSwitchIgmpSnoopingVlansVlanName(i["vlan-name"], d, pre_append)
			tmp["vlan_name"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchIgmpSnooping-Vlans-VlanName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingVlansProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingVlansQuerier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingVlansQuerierAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingVlansVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingVlansVlanName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchIpSourceGuard(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "binding_entry"
		if _, ok := i["binding-entry"]; ok {
			v := flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntry(i["binding-entry"], d, pre_append)
			tmp["binding_entry"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-IpSourceGuard-BindingEntry")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenSwitchControllerManagedSwitchIpSourceGuardDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-IpSourceGuard-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenSwitchControllerManagedSwitchIpSourceGuardPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-IpSourceGuard-Port")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntry(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entry_name"
		if _, ok := i["entry-name"]; ok {
			v := flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntryEntryName(i["entry-name"], d, pre_append)
			tmp["entry_name"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchIpSourceGuard-BindingEntry-EntryName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntryIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchIpSourceGuard-BindingEntry-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			v := flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntryMac(i["mac"], d, pre_append)
			tmp["mac"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchIpSourceGuard-BindingEntry-Mac")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntryEntryName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntryIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntryMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIpSourceGuardDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIpSourceGuardPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchL3Discovered(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchMaxAllowedTrunkMembers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchMclagIgmpSnoopingAware(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchMgmtMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchMirror(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := i["dst"]; ok {
			v := flattenSwitchControllerManagedSwitchMirrorDst(i["dst"], d, pre_append)
			tmp["dst"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Mirror-Dst")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSwitchControllerManagedSwitchMirrorName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Mirror-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_egress"
		if _, ok := i["src-egress"]; ok {
			v := flattenSwitchControllerManagedSwitchMirrorSrcEgress(i["src-egress"], d, pre_append)
			tmp["src_egress"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Mirror-SrcEgress")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_ingress"
		if _, ok := i["src-ingress"]; ok {
			v := flattenSwitchControllerManagedSwitchMirrorSrcIngress(i["src-ingress"], d, pre_append)
			tmp["src_ingress"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Mirror-SrcIngress")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenSwitchControllerManagedSwitchMirrorStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Mirror-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switching_packet"
		if _, ok := i["switching-packet"]; ok {
			v := flattenSwitchControllerManagedSwitchMirrorSwitchingPacket(i["switching-packet"], d, pre_append)
			tmp["switching_packet"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Mirror-SwitchingPacket")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerManagedSwitchMirrorDst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchMirrorName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchMirrorSrcEgress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchMirrorSrcIngress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchMirrorStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchMirrorSwitchingPacket(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchOverrideSnmpCommunity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchOverrideSnmpSysinfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchOverrideSnmpTrapThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchOverrideSnmpUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchOwnerVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPoeDetectionType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPoeLldpDetection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPoePreStandardDetection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPorts(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_mode"
		if _, ok := i["access-mode"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsAccessMode(i["access-mode"], d, pre_append)
			tmp["access_mode"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-AccessMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "acl_group"
		if _, ok := i["acl-group"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsAclGroup(i["acl-group"], d, pre_append)
			tmp["acl_group"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-AclGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "aggregator_mode"
		if _, ok := i["aggregator-mode"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsAggregatorMode(i["aggregator-mode"], d, pre_append)
			tmp["aggregator_mode"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-AggregatorMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allow_arp_monitor"
		if _, ok := i["allow-arp-monitor"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsAllowArpMonitor(i["allow-arp-monitor"], d, pre_append)
			tmp["allow_arp_monitor"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-AllowArpMonitor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowed_vlans"
		if _, ok := i["allowed-vlans"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsAllowedVlans(i["allowed-vlans"], d, pre_append)
			tmp["allowed_vlans"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-AllowedVlans")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowed_vlans_all"
		if _, ok := i["allowed-vlans-all"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsAllowedVlansAll(i["allowed-vlans-all"], d, pre_append)
			tmp["allowed_vlans_all"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-AllowedVlansAll")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "arp_inspection_trust"
		if _, ok := i["arp-inspection-trust"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsArpInspectionTrust(i["arp-inspection-trust"], d, pre_append)
			tmp["arp_inspection_trust"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-ArpInspectionTrust")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authenticated_port"
		if _, ok := i["authenticated-port"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsAuthenticatedPort(i["authenticated-port"], d, pre_append)
			tmp["authenticated_port"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-AuthenticatedPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bundle"
		if _, ok := i["bundle"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsBundle(i["bundle"], d, pre_append)
			tmp["bundle"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-Bundle")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_snoop_option82_override"
		if _, ok := i["dhcp-snoop-option82-override"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsDhcpSnoopOption82Override(i["dhcp-snoop-option82-override"], d, pre_append)
			tmp["dhcp_snoop_option82_override"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-DhcpSnoopOption82Override")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_snoop_option82_trust"
		if _, ok := i["dhcp-snoop-option82-trust"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsDhcpSnoopOption82Trust(i["dhcp-snoop-option82-trust"], d, pre_append)
			tmp["dhcp_snoop_option82_trust"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-DhcpSnoopOption82Trust")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_snooping"
		if _, ok := i["dhcp-snooping"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsDhcpSnooping(i["dhcp-snooping"], d, pre_append)
			tmp["dhcp_snooping"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-DhcpSnooping")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "discard_mode"
		if _, ok := i["discard-mode"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsDiscardMode(i["discard-mode"], d, pre_append)
			tmp["discard_mode"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-DiscardMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dsl_profile"
		if _, ok := i["dsl-profile"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsDslProfile(i["dsl-profile"], d, pre_append)
			tmp["dsl_profile"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-DslProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "edge_port"
		if _, ok := i["edge-port"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsEdgePort(i["edge-port"], d, pre_append)
			tmp["edge_port"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-EdgePort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "export_tags"
		if _, ok := i["export-tags"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsExportTags(i["export-tags"], d, pre_append)
			tmp["export_tags"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-ExportTags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "encrypted_port"
		if _, ok := i["encrypted-port"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsEncryptedPort(i["encrypted-port"], d, pre_append)
			tmp["encrypted_port"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-EncryptedPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "export_to"
		if _, ok := i["export-to"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsExportTo(i["export-to"], d, pre_append)
			tmp["export_to"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-ExportTo")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "export_to_pool"
		if _, ok := i["export-to-pool"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsExportToPool(i["export-to-pool"], d, pre_append)
			tmp["export_to_pool"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-ExportToPool")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "export_to_pool_flag"
		if _, ok := i["export-to-pool-flag"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsExportToPoolFlag(i["export-to-pool-flag"], d, pre_append)
			tmp["export_to_pool_flag"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-ExportToPoolFlag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fallback_port"
		if _, ok := i["fallback-port"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsFallbackPort(i["fallback-port"], d, pre_append)
			tmp["fallback_port"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-FallbackPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fec_capable"
		if _, ok := i["fec-capable"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsFecCapable(i["fec-capable"], d, pre_append)
			tmp["fec_capable"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-FecCapable")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fec_state"
		if _, ok := i["fec-state"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsFecState(i["fec-state"], d, pre_append)
			tmp["fec_state"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-FecState")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fgt_peer_device_name"
		if _, ok := i["fgt-peer-device-name"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsFgtPeerDeviceName(i["fgt-peer-device-name"], d, pre_append)
			tmp["fgt_peer_device_name"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-FgtPeerDeviceName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fgt_peer_port_name"
		if _, ok := i["fgt-peer-port-name"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsFgtPeerPortName(i["fgt-peer-port-name"], d, pre_append)
			tmp["fgt_peer_port_name"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-FgtPeerPortName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fiber_port"
		if _, ok := i["fiber-port"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsFiberPort(i["fiber-port"], d, pre_append)
			tmp["fiber_port"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-FiberPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := i["flags"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsFlags(i["flags"], d, pre_append)
			tmp["flags"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-Flags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flap_duration"
		if _, ok := i["flap-duration"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsFlapDuration(i["flap-duration"], d, pre_append)
			tmp["flap_duration"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-FlapDuration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flap_rate"
		if _, ok := i["flap-rate"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsFlapRate(i["flap-rate"], d, pre_append)
			tmp["flap_rate"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-FlapRate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flap_timeout"
		if _, ok := i["flap-timeout"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsFlapTimeout(i["flap-timeout"], d, pre_append)
			tmp["flap_timeout"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-FlapTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flapguard"
		if _, ok := i["flapguard"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsFlapguard(i["flapguard"], d, pre_append)
			tmp["flapguard"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-Flapguard")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flow_control"
		if _, ok := i["flow-control"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsFlowControl(i["flow-control"], d, pre_append)
			tmp["flow_control"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-FlowControl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortilink_port"
		if _, ok := i["fortilink-port"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsFortilinkPort(i["fortilink-port"], d, pre_append)
			tmp["fortilink_port"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-FortilinkPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmp_snooping"
		if _, ok := i["igmp-snooping"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsIgmpSnooping(i["igmp-snooping"], d, pre_append)
			tmp["igmp_snooping"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-IgmpSnooping")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortiswitch_acls"
		if _, ok := i["fortiswitch-acls"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsFortiswitchAcls(i["fortiswitch-acls"], d, pre_append)
			tmp["fortiswitch_acls"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-FortiswitchAcls")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmp_snooping_flood_reports"
		if _, ok := i["igmp-snooping-flood-reports"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsIgmpSnoopingFloodReports(i["igmp-snooping-flood-reports"], d, pre_append)
			tmp["igmp_snooping_flood_reports"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-IgmpSnoopingFloodReports")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmps_flood_reports"
		if _, ok := i["igmps-flood-reports"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsIgmpsFloodReports(i["igmps-flood-reports"], d, pre_append)
			tmp["igmps_flood_reports"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-IgmpsFloodReports")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmps_flood_traffic"
		if _, ok := i["igmps-flood-traffic"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsIgmpsFloodTraffic(i["igmps-flood-traffic"], d, pre_append)
			tmp["igmps_flood_traffic"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-IgmpsFloodTraffic")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_tags"
		if _, ok := i["interface-tags"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsInterfaceTags(i["interface-tags"], d, pre_append)
			tmp["interface_tags"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-InterfaceTags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_source_guard"
		if _, ok := i["ip-source-guard"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsIpSourceGuard(i["ip-source-guard"], d, pre_append)
			tmp["ip_source_guard"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-IpSourceGuard")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "isl_local_trunk_name"
		if _, ok := i["isl-local-trunk-name"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsIslLocalTrunkName(i["isl-local-trunk-name"], d, pre_append)
			tmp["isl_local_trunk_name"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-IslLocalTrunkName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "isl_peer_device_name"
		if _, ok := i["isl-peer-device-name"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsIslPeerDeviceName(i["isl-peer-device-name"], d, pre_append)
			tmp["isl_peer_device_name"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-IslPeerDeviceName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "isl_peer_device_sn"
		if _, ok := i["isl-peer-device-sn"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsIslPeerDeviceSn(i["isl-peer-device-sn"], d, pre_append)
			tmp["isl_peer_device_sn"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-IslPeerDeviceSn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "isl_peer_port_name"
		if _, ok := i["isl-peer-port-name"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsIslPeerPortName(i["isl-peer-port-name"], d, pre_append)
			tmp["isl_peer_port_name"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-IslPeerPortName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lacp_speed"
		if _, ok := i["lacp-speed"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsLacpSpeed(i["lacp-speed"], d, pre_append)
			tmp["lacp_speed"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-LacpSpeed")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "learning_limit"
		if _, ok := i["learning-limit"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsLearningLimit(i["learning-limit"], d, pre_append)
			tmp["learning_limit"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-LearningLimit")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_status"
		if _, ok := i["link-status"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsLinkStatus(i["link-status"], d, pre_append)
			tmp["link_status"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-LinkStatus")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lldp_profile"
		if _, ok := i["lldp-profile"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsLldpProfile(i["lldp-profile"], d, pre_append)
			tmp["lldp_profile"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-LldpProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lldp_status"
		if _, ok := i["lldp-status"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsLldpStatus(i["lldp-status"], d, pre_append)
			tmp["lldp_status"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-LldpStatus")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "loop_guard"
		if _, ok := i["loop-guard"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsLoopGuard(i["loop-guard"], d, pre_append)
			tmp["loop_guard"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-LoopGuard")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "loop_guard_timeout"
		if _, ok := i["loop-guard-timeout"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsLoopGuardTimeout(i["loop-guard-timeout"], d, pre_append)
			tmp["loop_guard_timeout"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-LoopGuardTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_addr"
		if _, ok := i["mac-addr"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsMacAddr(i["mac-addr"], d, pre_append)
			tmp["mac_addr"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-MacAddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "matched_dpp_intf_tags"
		if _, ok := i["matched-dpp-intf-tags"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsMatchedDppIntfTags(i["matched-dpp-intf-tags"], d, pre_append)
			tmp["matched_dpp_intf_tags"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-MatchedDppIntfTags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "matched_dpp_policy"
		if _, ok := i["matched-dpp-policy"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsMatchedDppPolicy(i["matched-dpp-policy"], d, pre_append)
			tmp["matched_dpp_policy"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-MatchedDppPolicy")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_bundle"
		if _, ok := i["max-bundle"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsMaxBundle(i["max-bundle"], d, pre_append)
			tmp["max_bundle"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-MaxBundle")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mcast_snooping_flood_traffic"
		if _, ok := i["mcast-snooping-flood-traffic"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsMcastSnoopingFloodTraffic(i["mcast-snooping-flood-traffic"], d, pre_append)
			tmp["mcast_snooping_flood_traffic"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-McastSnoopingFloodTraffic")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mclag"
		if _, ok := i["mclag"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsMclag(i["mclag"], d, pre_append)
			tmp["mclag"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-Mclag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mclag_icl_port"
		if _, ok := i["mclag-icl-port"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsMclagIclPort(i["mclag-icl-port"], d, pre_append)
			tmp["mclag_icl_port"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-MclagIclPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "media_type"
		if _, ok := i["media-type"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsMediaType(i["media-type"], d, pre_append)
			tmp["media_type"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-MediaType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member_withdrawal_behavior"
		if _, ok := i["member-withdrawal-behavior"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsMemberWithdrawalBehavior(i["member-withdrawal-behavior"], d, pre_append)
			tmp["member_withdrawal_behavior"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-MemberWithdrawalBehavior")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := i["members"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsMembers(i["members"], d, pre_append)
			tmp["members"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-Members")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min_bundle"
		if _, ok := i["min-bundle"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsMinBundle(i["min-bundle"], d, pre_append)
			tmp["min_bundle"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-MinBundle")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := i["mode"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsMode(i["mode"], d, pre_append)
			tmp["mode"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-Mode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "p2p_port"
		if _, ok := i["p2p-port"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsP2PPort(i["p2p-port"], d, pre_append)
			tmp["p2p_port"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-P2PPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_sample_rate"
		if _, ok := i["packet-sample-rate"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsPacketSampleRate(i["packet-sample-rate"], d, pre_append)
			tmp["packet_sample_rate"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-PacketSampleRate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_sampler"
		if _, ok := i["packet-sampler"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsPacketSampler(i["packet-sampler"], d, pre_append)
			tmp["packet_sampler"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-PacketSampler")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pause_meter"
		if _, ok := i["pause-meter"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsPauseMeter(i["pause-meter"], d, pre_append)
			tmp["pause_meter"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-PauseMeter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pause_meter_resume"
		if _, ok := i["pause-meter-resume"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsPauseMeterResume(i["pause-meter-resume"], d, pre_append)
			tmp["pause_meter_resume"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-PauseMeterResume")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_capable"
		if _, ok := i["poe-capable"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsPoeCapable(i["poe-capable"], d, pre_append)
			tmp["poe_capable"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-PoeCapable")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_max_power"
		if _, ok := i["poe-max-power"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsPoeMaxPower(i["poe-max-power"], d, pre_append)
			tmp["poe_max_power"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-PoeMaxPower")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_mode_bt_cabable"
		if _, ok := i["poe-mode-bt-cabable"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsPoeModeBtCabable(i["poe-mode-bt-cabable"], d, pre_append)
			tmp["poe_mode_bt_cabable"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-PoeModeBtCabable")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_port_mode"
		if _, ok := i["poe-port-mode"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsPoePortMode(i["poe-port-mode"], d, pre_append)
			tmp["poe_port_mode"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-PoePortMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_port_power"
		if _, ok := i["poe-port-power"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsPoePortPower(i["poe-port-power"], d, pre_append)
			tmp["poe_port_power"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-PoePortPower")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_port_priority"
		if _, ok := i["poe-port-priority"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsPoePortPriority(i["poe-port-priority"], d, pre_append)
			tmp["poe_port_priority"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-PoePortPriority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_pre_standard_detection"
		if _, ok := i["poe-pre-standard-detection"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsPoePreStandardDetection(i["poe-pre-standard-detection"], d, pre_append)
			tmp["poe_pre_standard_detection"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-PoePreStandardDetection")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_standard"
		if _, ok := i["poe-standard"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsPoeStandard(i["poe-standard"], d, pre_append)
			tmp["poe_standard"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-PoeStandard")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_status"
		if _, ok := i["poe-status"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsPoeStatus(i["poe-status"], d, pre_append)
			tmp["poe_status"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-PoeStatus")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_name"
		if _, ok := i["port-name"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsPortName(i["port-name"], d, pre_append)
			tmp["port_name"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-PortName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_number"
		if _, ok := i["port-number"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsPortNumber(i["port-number"], d, pre_append)
			tmp["port_number"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-PortNumber")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_owner"
		if _, ok := i["port-owner"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsPortOwner(i["port-owner"], d, pre_append)
			tmp["port_owner"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-PortOwner")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_policy"
		if _, ok := i["port-policy"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsPortPolicy(i["port-policy"], d, pre_append)
			tmp["port_policy"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-PortPolicy")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_prefix_type"
		if _, ok := i["port-prefix-type"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsPortPrefixType(i["port-prefix-type"], d, pre_append)
			tmp["port_prefix_type"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-PortPrefixType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_security_policy"
		if _, ok := i["port-security-policy"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsPortSecurityPolicy(i["port-security-policy"], d, pre_append)
			tmp["port_security_policy"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-PortSecurityPolicy")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_selection_criteria"
		if _, ok := i["port-selection-criteria"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsPortSelectionCriteria(i["port-selection-criteria"], d, pre_append)
			tmp["port_selection_criteria"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-PortSelectionCriteria")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ptp_policy"
		if _, ok := i["ptp-policy"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsPtpPolicy(i["ptp-policy"], d, pre_append)
			tmp["ptp_policy"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-PtpPolicy")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ptp_status"
		if _, ok := i["ptp-status"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsPtpStatus(i["ptp-status"], d, pre_append)
			tmp["ptp_status"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-PtpStatus")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "qos_policy"
		if _, ok := i["qos-policy"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsQosPolicy(i["qos-policy"], d, pre_append)
			tmp["qos_policy"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-QosPolicy")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restricted_auth_port"
		if _, ok := i["restricted-auth-port"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsRestrictedAuthPort(i["restricted-auth-port"], d, pre_append)
			tmp["restricted_auth_port"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-RestrictedAuthPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rpvst_port"
		if _, ok := i["rpvst-port"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsRpvstPort(i["rpvst-port"], d, pre_append)
			tmp["rpvst_port"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-RpvstPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sample_direction"
		if _, ok := i["sample-direction"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsSampleDirection(i["sample-direction"], d, pre_append)
			tmp["sample_direction"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-SampleDirection")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sflow_counter_interval"
		if _, ok := i["sflow-counter-interval"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsSflowCounterInterval(i["sflow-counter-interval"], d, pre_append)
			tmp["sflow_counter_interval"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-SflowCounterInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "speed"
		if _, ok := i["speed"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsSpeed(i["speed"], d, pre_append)
			tmp["speed"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-Speed")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "speed_mask"
		if _, ok := i["speed-mask"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsSpeedMask(i["speed-mask"], d, pre_append)
			tmp["speed_mask"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-SpeedMask")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stacking_port"
		if _, ok := i["stacking-port"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsStackingPort(i["stacking-port"], d, pre_append)
			tmp["stacking_port"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-StackingPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sticky_mac"
		if _, ok := i["sticky-mac"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsStickyMac(i["sticky-mac"], d, pre_append)
			tmp["sticky_mac"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-StickyMac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "storm_control_policy"
		if _, ok := i["storm-control-policy"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsStormControlPolicy(i["storm-control-policy"], d, pre_append)
			tmp["storm_control_policy"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-StormControlPolicy")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stp_bpdu_guard"
		if _, ok := i["stp-bpdu-guard"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsStpBpduGuard(i["stp-bpdu-guard"], d, pre_append)
			tmp["stp_bpdu_guard"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-StpBpduGuard")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stp_bpdu_guard_timeout"
		if _, ok := i["stp-bpdu-guard-timeout"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsStpBpduGuardTimeout(i["stp-bpdu-guard-timeout"], d, pre_append)
			tmp["stp_bpdu_guard_timeout"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-StpBpduGuardTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stp_root_guard"
		if _, ok := i["stp-root-guard"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsStpRootGuard(i["stp-root-guard"], d, pre_append)
			tmp["stp_root_guard"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-StpRootGuard")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stp_state"
		if _, ok := i["stp-state"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsStpState(i["stp-state"], d, pre_append)
			tmp["stp_state"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-StpState")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_id"
		if _, ok := i["switch-id"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsSwitchId(i["switch-id"], d, pre_append)
			tmp["switch_id"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-SwitchId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trunk_member"
		if _, ok := i["trunk-member"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsTrunkMember(i["trunk-member"], d, pre_append)
			tmp["trunk_member"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-TrunkMember")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "untagged_vlans"
		if _, ok := i["untagged-vlans"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsUntaggedVlans(i["untagged-vlans"], d, pre_append)
			tmp["untagged_vlans"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-UntaggedVlans")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_port"
		if _, ok := i["virtual-port"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsVirtualPort(i["virtual-port"], d, pre_append)
			tmp["virtual_port"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-VirtualPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan"
		if _, ok := i["vlan"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsVlan(i["vlan"], d, pre_append)
			tmp["vlan"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Ports-Vlan")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerManagedSwitchPortsAccessMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsAclGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchPortsAggregatorMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsAllowArpMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsAllowedVlans(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchPortsAllowedVlansAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsArpInspectionTrust(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsAuthenticatedPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsBundle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsDhcpSnoopOption82Override(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id"
		if _, ok := i["circuit-id"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideCircuitId(i["circuit-id"], d, pre_append)
			tmp["circuit_id"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchPorts-DhcpSnoopOption82Override-CircuitId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
		if _, ok := i["remote-id"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideRemoteId(i["remote-id"], d, pre_append)
			tmp["remote_id"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchPorts-DhcpSnoopOption82Override-RemoteId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if _, ok := i["vlan-name"]; ok {
			v := flattenSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideVlanName(i["vlan-name"], d, pre_append)
			tmp["vlan_name"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchPorts-DhcpSnoopOption82Override-VlanName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideCircuitId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideRemoteId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideVlanName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchPortsDhcpSnoopOption82Trust(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsDhcpSnooping(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsDiscardMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsDslProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchPortsEdgePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsExportTags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchPortsEncryptedPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsExportTo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchPortsExportToPool(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchPortsExportToPoolFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsFallbackPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsFecCapable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsFecState(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsFgtPeerDeviceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsFgtPeerPortName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsFiberPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsFlapDuration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsFlapRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsFlapTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsFlapguard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsFlowControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsFortilinkPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsIgmpSnooping(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsFortiswitchAcls(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenSwitchControllerManagedSwitchPortsIgmpSnoopingFloodReports(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsIgmpsFloodReports(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsIgmpsFloodTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsInterfaceTags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchPortsIpSourceGuard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsIslLocalTrunkName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsIslPeerDeviceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsIslPeerDeviceSn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsIslPeerPortName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsLacpSpeed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsLearningLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsLinkStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsLldpProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchPortsLldpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsLoopGuard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsLoopGuardTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsMacAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsMatchedDppIntfTags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsMatchedDppPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsMaxBundle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsMcastSnoopingFloodTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsMclag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsMclagIclPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsMediaType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsMemberWithdrawalBehavior(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsMembers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchPortsMinBundle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsP2PPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPacketSampleRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPacketSampler(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPauseMeter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPauseMeterResume(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPoeCapable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPoeMaxPower(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPoeModeBtCabable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPoePortMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPoePortPower(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPoePortPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPoePreStandardDetection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPoeStandard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPoeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPortName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPortNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPortOwner(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPortPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchPortsPortPrefixType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPortSecurityPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchPortsPortSelectionCriteria(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPtpPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchPortsPtpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsQosPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchPortsRestrictedAuthPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsRpvstPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsSampleDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsSflowCounterInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsSpeed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsSpeedMask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsStackingPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsStickyMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsStormControlPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchPortsStpBpduGuard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsStpBpduGuardTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsStpRootGuard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsStpState(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsSwitchId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsTrunkMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsUntaggedVlans(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchPortsVirtualPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchPreProvisioned(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPtpProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchPtpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPurdueLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchQosDropPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchQosRedProbability(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRadiusNasIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRadiusNasIpOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRemoteLog(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "csv"
		if _, ok := i["csv"]; ok {
			v := flattenSwitchControllerManagedSwitchRemoteLogCsv(i["csv"], d, pre_append)
			tmp["csv"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-RemoteLog-Csv")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "facility"
		if _, ok := i["facility"]; ok {
			v := flattenSwitchControllerManagedSwitchRemoteLogFacility(i["facility"], d, pre_append)
			tmp["facility"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-RemoteLog-Facility")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSwitchControllerManagedSwitchRemoteLogName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-RemoteLog-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenSwitchControllerManagedSwitchRemoteLogPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-RemoteLog-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {
			v := flattenSwitchControllerManagedSwitchRemoteLogServer(i["server"], d, pre_append)
			tmp["server"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-RemoteLog-Server")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := i["severity"]; ok {
			v := flattenSwitchControllerManagedSwitchRemoteLogSeverity(i["severity"], d, pre_append)
			tmp["severity"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-RemoteLog-Severity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenSwitchControllerManagedSwitchRemoteLogStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-RemoteLog-Status")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerManagedSwitchRemoteLogCsv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRemoteLogFacility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRemoteLogName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRemoteLogPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRemoteLogServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRemoteLogSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRemoteLogStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRouteOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRouteOffloadMclag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRouteOffloadRouter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "router_ip"
		if _, ok := i["router-ip"]; ok {
			v := flattenSwitchControllerManagedSwitchRouteOffloadRouterRouterIp(i["router-ip"], d, pre_append)
			tmp["router_ip"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-RouteOffloadRouter-RouterIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if _, ok := i["vlan-name"]; ok {
			v := flattenSwitchControllerManagedSwitchRouteOffloadRouterVlanName(i["vlan-name"], d, pre_append)
			tmp["vlan_name"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-RouteOffloadRouter-VlanName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerManagedSwitchRouteOffloadRouterRouterIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRouteOffloadRouterVlanName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchSn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunity(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "events"
		if _, ok := i["events"]; ok {
			v := flattenSwitchControllerManagedSwitchSnmpCommunityEvents(i["events"], d, pre_append)
			tmp["events"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-SnmpCommunity-Events")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hosts"
		if _, ok := i["hosts"]; ok {
			v := flattenSwitchControllerManagedSwitchSnmpCommunityHosts(i["hosts"], d, pre_append)
			tmp["hosts"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-SnmpCommunity-Hosts")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSwitchControllerManagedSwitchSnmpCommunityId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-SnmpCommunity-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSwitchControllerManagedSwitchSnmpCommunityName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-SnmpCommunity-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_v1_port"
		if _, ok := i["query-v1-port"]; ok {
			v := flattenSwitchControllerManagedSwitchSnmpCommunityQueryV1Port(i["query-v1-port"], d, pre_append)
			tmp["query_v1_port"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-SnmpCommunity-QueryV1Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_v1_status"
		if _, ok := i["query-v1-status"]; ok {
			v := flattenSwitchControllerManagedSwitchSnmpCommunityQueryV1Status(i["query-v1-status"], d, pre_append)
			tmp["query_v1_status"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-SnmpCommunity-QueryV1Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_v2c_port"
		if _, ok := i["query-v2c-port"]; ok {
			v := flattenSwitchControllerManagedSwitchSnmpCommunityQueryV2CPort(i["query-v2c-port"], d, pre_append)
			tmp["query_v2c_port"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-SnmpCommunity-QueryV2CPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_v2c_status"
		if _, ok := i["query-v2c-status"]; ok {
			v := flattenSwitchControllerManagedSwitchSnmpCommunityQueryV2CStatus(i["query-v2c-status"], d, pre_append)
			tmp["query_v2c_status"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-SnmpCommunity-QueryV2CStatus")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenSwitchControllerManagedSwitchSnmpCommunityStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-SnmpCommunity-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v1_lport"
		if _, ok := i["trap-v1-lport"]; ok {
			v := flattenSwitchControllerManagedSwitchSnmpCommunityTrapV1Lport(i["trap-v1-lport"], d, pre_append)
			tmp["trap_v1_lport"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-SnmpCommunity-TrapV1Lport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v1_rport"
		if _, ok := i["trap-v1-rport"]; ok {
			v := flattenSwitchControllerManagedSwitchSnmpCommunityTrapV1Rport(i["trap-v1-rport"], d, pre_append)
			tmp["trap_v1_rport"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-SnmpCommunity-TrapV1Rport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v1_status"
		if _, ok := i["trap-v1-status"]; ok {
			v := flattenSwitchControllerManagedSwitchSnmpCommunityTrapV1Status(i["trap-v1-status"], d, pre_append)
			tmp["trap_v1_status"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-SnmpCommunity-TrapV1Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v2c_lport"
		if _, ok := i["trap-v2c-lport"]; ok {
			v := flattenSwitchControllerManagedSwitchSnmpCommunityTrapV2CLport(i["trap-v2c-lport"], d, pre_append)
			tmp["trap_v2c_lport"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-SnmpCommunity-TrapV2CLport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v2c_rport"
		if _, ok := i["trap-v2c-rport"]; ok {
			v := flattenSwitchControllerManagedSwitchSnmpCommunityTrapV2CRport(i["trap-v2c-rport"], d, pre_append)
			tmp["trap_v2c_rport"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-SnmpCommunity-TrapV2CRport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v2c_status"
		if _, ok := i["trap-v2c-status"]; ok {
			v := flattenSwitchControllerManagedSwitchSnmpCommunityTrapV2CStatus(i["trap-v2c-status"], d, pre_append)
			tmp["trap_v2c_status"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-SnmpCommunity-TrapV2CStatus")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerManagedSwitchSnmpCommunityEvents(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchSnmpCommunityHosts(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSwitchControllerManagedSwitchSnmpCommunityHostsId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchSnmpCommunity-Hosts-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenSwitchControllerManagedSwitchSnmpCommunityHostsIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchSnmpCommunity-Hosts-Ip")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerManagedSwitchSnmpCommunityHostsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityHostsIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchSnmpCommunityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityQueryV1Port(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityQueryV1Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityQueryV2CPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityQueryV2CStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityTrapV1Lport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityTrapV1Rport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityTrapV1Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityTrapV2CLport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityTrapV2CRport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityTrapV2CStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpSysinfo(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "contact_info"
	if _, ok := i["contact-info"]; ok {
		result["contact_info"] = flattenSwitchControllerManagedSwitchSnmpSysinfoContactInfo(i["contact-info"], d, pre_append)
	}

	pre_append = pre + ".0." + "description"
	if _, ok := i["description"]; ok {
		result["description"] = flattenSwitchControllerManagedSwitchSnmpSysinfoDescription(i["description"], d, pre_append)
	}

	pre_append = pre + ".0." + "engine_id"
	if _, ok := i["engine-id"]; ok {
		result["engine_id"] = flattenSwitchControllerManagedSwitchSnmpSysinfoEngineId(i["engine-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "location"
	if _, ok := i["location"]; ok {
		result["location"] = flattenSwitchControllerManagedSwitchSnmpSysinfoLocation(i["location"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSwitchControllerManagedSwitchSnmpSysinfoStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchControllerManagedSwitchSnmpSysinfoContactInfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpSysinfoDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpSysinfoEngineId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpSysinfoLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpSysinfoStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpTrapThreshold(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "trap_high_cpu_threshold"
	if _, ok := i["trap-high-cpu-threshold"]; ok {
		result["trap_high_cpu_threshold"] = flattenSwitchControllerManagedSwitchSnmpTrapThresholdTrapHighCpuThreshold(i["trap-high-cpu-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "trap_log_full_threshold"
	if _, ok := i["trap-log-full-threshold"]; ok {
		result["trap_log_full_threshold"] = flattenSwitchControllerManagedSwitchSnmpTrapThresholdTrapLogFullThreshold(i["trap-log-full-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "trap_low_memory_threshold"
	if _, ok := i["trap-low-memory-threshold"]; ok {
		result["trap_low_memory_threshold"] = flattenSwitchControllerManagedSwitchSnmpTrapThresholdTrapLowMemoryThreshold(i["trap-low-memory-threshold"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchControllerManagedSwitchSnmpTrapThresholdTrapHighCpuThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpTrapThresholdTrapLogFullThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpTrapThresholdTrapLowMemoryThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpUser(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_proto"
		if _, ok := i["auth-proto"]; ok {
			v := flattenSwitchControllerManagedSwitchSnmpUserAuthProto(i["auth-proto"], d, pre_append)
			tmp["auth_proto"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-SnmpUser-AuthProto")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSwitchControllerManagedSwitchSnmpUserName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-SnmpUser-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priv_proto"
		if _, ok := i["priv-proto"]; ok {
			v := flattenSwitchControllerManagedSwitchSnmpUserPrivProto(i["priv-proto"], d, pre_append)
			tmp["priv_proto"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-SnmpUser-PrivProto")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queries"
		if _, ok := i["queries"]; ok {
			v := flattenSwitchControllerManagedSwitchSnmpUserQueries(i["queries"], d, pre_append)
			tmp["queries"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-SnmpUser-Queries")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_port"
		if _, ok := i["query-port"]; ok {
			v := flattenSwitchControllerManagedSwitchSnmpUserQueryPort(i["query-port"], d, pre_append)
			tmp["query_port"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-SnmpUser-QueryPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_level"
		if _, ok := i["security-level"]; ok {
			v := flattenSwitchControllerManagedSwitchSnmpUserSecurityLevel(i["security-level"], d, pre_append)
			tmp["security_level"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-SnmpUser-SecurityLevel")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerManagedSwitchSnmpUserAuthProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpUserName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpUserPrivProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpUserQueries(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpUserQueryPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpUserSecurityLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStagedImageVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStaticMac(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenSwitchControllerManagedSwitchStaticMacDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-StaticMac-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSwitchControllerManagedSwitchStaticMacId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-StaticMac-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenSwitchControllerManagedSwitchStaticMacInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-StaticMac-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			v := flattenSwitchControllerManagedSwitchStaticMacMac(i["mac"], d, pre_append)
			tmp["mac"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-StaticMac-Mac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenSwitchControllerManagedSwitchStaticMacType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-StaticMac-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan"
		if _, ok := i["vlan"]; ok {
			v := flattenSwitchControllerManagedSwitchStaticMacVlan(i["vlan"], d, pre_append)
			tmp["vlan"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-StaticMac-Vlan")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerManagedSwitchStaticMacDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStaticMacId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStaticMacInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStaticMacMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStaticMacType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStaticMacVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchStormControl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "broadcast"
	if _, ok := i["broadcast"]; ok {
		result["broadcast"] = flattenSwitchControllerManagedSwitchStormControlBroadcast(i["broadcast"], d, pre_append)
	}

	pre_append = pre + ".0." + "local_override"
	if _, ok := i["local-override"]; ok {
		result["local_override"] = flattenSwitchControllerManagedSwitchStormControlLocalOverride(i["local-override"], d, pre_append)
	}

	pre_append = pre + ".0." + "rate"
	if _, ok := i["rate"]; ok {
		result["rate"] = flattenSwitchControllerManagedSwitchStormControlRate(i["rate"], d, pre_append)
	}

	pre_append = pre + ".0." + "unknown_multicast"
	if _, ok := i["unknown-multicast"]; ok {
		result["unknown_multicast"] = flattenSwitchControllerManagedSwitchStormControlUnknownMulticast(i["unknown-multicast"], d, pre_append)
	}

	pre_append = pre + ".0." + "unknown_unicast"
	if _, ok := i["unknown-unicast"]; ok {
		result["unknown_unicast"] = flattenSwitchControllerManagedSwitchStormControlUnknownUnicast(i["unknown-unicast"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchControllerManagedSwitchStormControlBroadcast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStormControlLocalOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStormControlRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStormControlUnknownMulticast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStormControlUnknownUnicast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStpInstance(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSwitchControllerManagedSwitchStpInstanceId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-StpInstance-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenSwitchControllerManagedSwitchStpInstancePriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-StpInstance-Priority")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerManagedSwitchStpInstanceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStpInstancePriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStpSettings(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "forward_time"
	if _, ok := i["forward-time"]; ok {
		result["forward_time"] = flattenSwitchControllerManagedSwitchStpSettingsForwardTime(i["forward-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "hello_time"
	if _, ok := i["hello-time"]; ok {
		result["hello_time"] = flattenSwitchControllerManagedSwitchStpSettingsHelloTime(i["hello-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "local_override"
	if _, ok := i["local-override"]; ok {
		result["local_override"] = flattenSwitchControllerManagedSwitchStpSettingsLocalOverride(i["local-override"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_age"
	if _, ok := i["max-age"]; ok {
		result["max_age"] = flattenSwitchControllerManagedSwitchStpSettingsMaxAge(i["max-age"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_hops"
	if _, ok := i["max-hops"]; ok {
		result["max_hops"] = flattenSwitchControllerManagedSwitchStpSettingsMaxHops(i["max-hops"], d, pre_append)
	}

	pre_append = pre + ".0." + "name"
	if _, ok := i["name"]; ok {
		result["name"] = flattenSwitchControllerManagedSwitchStpSettingsName(i["name"], d, pre_append)
	}

	pre_append = pre + ".0." + "pending_timer"
	if _, ok := i["pending-timer"]; ok {
		result["pending_timer"] = flattenSwitchControllerManagedSwitchStpSettingsPendingTimer(i["pending-timer"], d, pre_append)
	}

	pre_append = pre + ".0." + "revision"
	if _, ok := i["revision"]; ok {
		result["revision"] = flattenSwitchControllerManagedSwitchStpSettingsRevision(i["revision"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSwitchControllerManagedSwitchStpSettingsStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchControllerManagedSwitchStpSettingsForwardTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStpSettingsHelloTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStpSettingsLocalOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStpSettingsMaxAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStpSettingsMaxHops(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStpSettingsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStpSettingsPendingTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStpSettingsRevision(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStpSettingsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSwitchDeviceTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSwitchDhcpOpt43Key(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSwitchId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSwitchLog(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "local_override"
	if _, ok := i["local-override"]; ok {
		result["local_override"] = flattenSwitchControllerManagedSwitchSwitchLogLocalOverride(i["local-override"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenSwitchControllerManagedSwitchSwitchLogSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSwitchControllerManagedSwitchSwitchLogStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchControllerManagedSwitchSwitchLogLocalOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSwitchLogSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSwitchLogStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSwitchProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchTdrSupported(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchTunnelDiscovered(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchVlan(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "assignment_priority"
		if _, ok := i["assignment-priority"]; ok {
			v := flattenSwitchControllerManagedSwitchVlanAssignmentPriority(i["assignment-priority"], d, pre_append)
			tmp["assignment_priority"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Vlan-AssignmentPriority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if _, ok := i["vlan-name"]; ok {
			v := flattenSwitchControllerManagedSwitchVlanVlanName(i["vlan-name"], d, pre_append)
			tmp["vlan_name"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitch-Vlan-VlanName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerManagedSwitchVlanAssignmentPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchVlanVlanName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSwitchControllerManagedSwitch(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("n802_1x_settings", flattenSwitchControllerManagedSwitch8021XSettings(o["802-1X-settings"], d, "n802_1x_settings")); err != nil {
			if vv, ok := fortiAPIPatch(o["802-1X-settings"], "SwitchControllerManagedSwitch-8021XSettings"); ok {
				if err = d.Set("n802_1x_settings", vv); err != nil {
					return fmt.Errorf("Error reading n802_1x_settings: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading n802_1x_settings: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("n802_1x_settings"); ok {
			if err = d.Set("n802_1x_settings", flattenSwitchControllerManagedSwitch8021XSettings(o["802-1X-settings"], d, "n802_1x_settings")); err != nil {
				if vv, ok := fortiAPIPatch(o["802-1X-settings"], "SwitchControllerManagedSwitch-8021XSettings"); ok {
					if err = d.Set("n802_1x_settings", vv); err != nil {
						return fmt.Errorf("Error reading n802_1x_settings: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading n802_1x_settings: %v", err)
				}
			}
		}
	}

	if err = d.Set("_platform", flattenSwitchControllerManagedSwitchPlatform(o["_platform"], d, "_platform")); err != nil {
		if vv, ok := fortiAPIPatch(o["_platform"], "SwitchControllerManagedSwitch-Platform"); ok {
			if err = d.Set("_platform", vv); err != nil {
				return fmt.Errorf("Error reading _platform: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _platform: %v", err)
		}
	}

	if err = d.Set("access_profile", flattenSwitchControllerManagedSwitchAccessProfile(o["access-profile"], d, "access_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-profile"], "SwitchControllerManagedSwitch-AccessProfile"); ok {
			if err = d.Set("access_profile", vv); err != nil {
				return fmt.Errorf("Error reading access_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_profile: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("custom_command", flattenSwitchControllerManagedSwitchCustomCommand(o["custom-command"], d, "custom_command")); err != nil {
			if vv, ok := fortiAPIPatch(o["custom-command"], "SwitchControllerManagedSwitch-CustomCommand"); ok {
				if err = d.Set("custom_command", vv); err != nil {
					return fmt.Errorf("Error reading custom_command: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading custom_command: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_command"); ok {
			if err = d.Set("custom_command", flattenSwitchControllerManagedSwitchCustomCommand(o["custom-command"], d, "custom_command")); err != nil {
				if vv, ok := fortiAPIPatch(o["custom-command"], "SwitchControllerManagedSwitch-CustomCommand"); ok {
					if err = d.Set("custom_command", vv); err != nil {
						return fmt.Errorf("Error reading custom_command: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading custom_command: %v", err)
				}
			}
		}
	}

	if err = d.Set("delayed_restart_trigger", flattenSwitchControllerManagedSwitchDelayedRestartTrigger(o["delayed-restart-trigger"], d, "delayed_restart_trigger")); err != nil {
		if vv, ok := fortiAPIPatch(o["delayed-restart-trigger"], "SwitchControllerManagedSwitch-DelayedRestartTrigger"); ok {
			if err = d.Set("delayed_restart_trigger", vv); err != nil {
				return fmt.Errorf("Error reading delayed_restart_trigger: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading delayed_restart_trigger: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchControllerManagedSwitchDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SwitchControllerManagedSwitch-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("dhcp_server_access_list", flattenSwitchControllerManagedSwitchDhcpServerAccessList(o["dhcp-server-access-list"], d, "dhcp_server_access_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-server-access-list"], "SwitchControllerManagedSwitch-DhcpServerAccessList"); ok {
			if err = d.Set("dhcp_server_access_list", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_server_access_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_server_access_list: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dhcp_snooping_static_client", flattenSwitchControllerManagedSwitchDhcpSnoopingStaticClient(o["dhcp-snooping-static-client"], d, "dhcp_snooping_static_client")); err != nil {
			if vv, ok := fortiAPIPatch(o["dhcp-snooping-static-client"], "SwitchControllerManagedSwitch-DhcpSnoopingStaticClient"); ok {
				if err = d.Set("dhcp_snooping_static_client", vv); err != nil {
					return fmt.Errorf("Error reading dhcp_snooping_static_client: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dhcp_snooping_static_client: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dhcp_snooping_static_client"); ok {
			if err = d.Set("dhcp_snooping_static_client", flattenSwitchControllerManagedSwitchDhcpSnoopingStaticClient(o["dhcp-snooping-static-client"], d, "dhcp_snooping_static_client")); err != nil {
				if vv, ok := fortiAPIPatch(o["dhcp-snooping-static-client"], "SwitchControllerManagedSwitch-DhcpSnoopingStaticClient"); ok {
					if err = d.Set("dhcp_snooping_static_client", vv); err != nil {
						return fmt.Errorf("Error reading dhcp_snooping_static_client: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dhcp_snooping_static_client: %v", err)
				}
			}
		}
	}

	if err = d.Set("directly_connected", flattenSwitchControllerManagedSwitchDirectlyConnected(o["directly-connected"], d, "directly_connected")); err != nil {
		if vv, ok := fortiAPIPatch(o["directly-connected"], "SwitchControllerManagedSwitch-DirectlyConnected"); ok {
			if err = d.Set("directly_connected", vv); err != nil {
				return fmt.Errorf("Error reading directly_connected: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading directly_connected: %v", err)
		}
	}

	if err = d.Set("dynamic_capability", flattenSwitchControllerManagedSwitchDynamicCapability(o["dynamic-capability"], d, "dynamic_capability")); err != nil {
		if vv, ok := fortiAPIPatch(o["dynamic-capability"], "SwitchControllerManagedSwitch-DynamicCapability"); ok {
			if err = d.Set("dynamic_capability", vv); err != nil {
				return fmt.Errorf("Error reading dynamic_capability: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dynamic_capability: %v", err)
		}
	}

	if err = d.Set("dynamically_discovered", flattenSwitchControllerManagedSwitchDynamicallyDiscovered(o["dynamically-discovered"], d, "dynamically_discovered")); err != nil {
		if vv, ok := fortiAPIPatch(o["dynamically-discovered"], "SwitchControllerManagedSwitch-DynamicallyDiscovered"); ok {
			if err = d.Set("dynamically_discovered", vv); err != nil {
				return fmt.Errorf("Error reading dynamically_discovered: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dynamically_discovered: %v", err)
		}
	}

	if err = d.Set("firmware_provision", flattenSwitchControllerManagedSwitchFirmwareProvision(o["firmware-provision"], d, "firmware_provision")); err != nil {
		if vv, ok := fortiAPIPatch(o["firmware-provision"], "SwitchControllerManagedSwitch-FirmwareProvision"); ok {
			if err = d.Set("firmware_provision", vv); err != nil {
				return fmt.Errorf("Error reading firmware_provision: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading firmware_provision: %v", err)
		}
	}

	if err = d.Set("firmware_provision_latest", flattenSwitchControllerManagedSwitchFirmwareProvisionLatest(o["firmware-provision-latest"], d, "firmware_provision_latest")); err != nil {
		if vv, ok := fortiAPIPatch(o["firmware-provision-latest"], "SwitchControllerManagedSwitch-FirmwareProvisionLatest"); ok {
			if err = d.Set("firmware_provision_latest", vv); err != nil {
				return fmt.Errorf("Error reading firmware_provision_latest: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading firmware_provision_latest: %v", err)
		}
	}

	if err = d.Set("firmware_provision_version", flattenSwitchControllerManagedSwitchFirmwareProvisionVersion(o["firmware-provision-version"], d, "firmware_provision_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["firmware-provision-version"], "SwitchControllerManagedSwitch-FirmwareProvisionVersion"); ok {
			if err = d.Set("firmware_provision_version", vv); err != nil {
				return fmt.Errorf("Error reading firmware_provision_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading firmware_provision_version: %v", err)
		}
	}

	if err = d.Set("flow_identity", flattenSwitchControllerManagedSwitchFlowIdentity(o["flow-identity"], d, "flow_identity")); err != nil {
		if vv, ok := fortiAPIPatch(o["flow-identity"], "SwitchControllerManagedSwitch-FlowIdentity"); ok {
			if err = d.Set("flow_identity", vv); err != nil {
				return fmt.Errorf("Error reading flow_identity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flow_identity: %v", err)
		}
	}

	if err = d.Set("fsw_wan1_admin", flattenSwitchControllerManagedSwitchFswWan1Admin(o["fsw-wan1-admin"], d, "fsw_wan1_admin")); err != nil {
		if vv, ok := fortiAPIPatch(o["fsw-wan1-admin"], "SwitchControllerManagedSwitch-FswWan1Admin"); ok {
			if err = d.Set("fsw_wan1_admin", vv); err != nil {
				return fmt.Errorf("Error reading fsw_wan1_admin: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fsw_wan1_admin: %v", err)
		}
	}

	if err = d.Set("fsw_wan1_peer", flattenSwitchControllerManagedSwitchFswWan1Peer(o["fsw-wan1-peer"], d, "fsw_wan1_peer")); err != nil {
		if vv, ok := fortiAPIPatch(o["fsw-wan1-peer"], "SwitchControllerManagedSwitch-FswWan1Peer"); ok {
			if err = d.Set("fsw_wan1_peer", vv); err != nil {
				return fmt.Errorf("Error reading fsw_wan1_peer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fsw_wan1_peer: %v", err)
		}
	}

	if err = d.Set("fsw_wan2_admin", flattenSwitchControllerManagedSwitchFswWan2Admin(o["fsw-wan2-admin"], d, "fsw_wan2_admin")); err != nil {
		if vv, ok := fortiAPIPatch(o["fsw-wan2-admin"], "SwitchControllerManagedSwitch-FswWan2Admin"); ok {
			if err = d.Set("fsw_wan2_admin", vv); err != nil {
				return fmt.Errorf("Error reading fsw_wan2_admin: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fsw_wan2_admin: %v", err)
		}
	}

	if err = d.Set("fsw_wan2_peer", flattenSwitchControllerManagedSwitchFswWan2Peer(o["fsw-wan2-peer"], d, "fsw_wan2_peer")); err != nil {
		if vv, ok := fortiAPIPatch(o["fsw-wan2-peer"], "SwitchControllerManagedSwitch-FswWan2Peer"); ok {
			if err = d.Set("fsw_wan2_peer", vv); err != nil {
				return fmt.Errorf("Error reading fsw_wan2_peer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fsw_wan2_peer: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("igmp_snooping", flattenSwitchControllerManagedSwitchIgmpSnooping(o["igmp-snooping"], d, "igmp_snooping")); err != nil {
			if vv, ok := fortiAPIPatch(o["igmp-snooping"], "SwitchControllerManagedSwitch-IgmpSnooping"); ok {
				if err = d.Set("igmp_snooping", vv); err != nil {
					return fmt.Errorf("Error reading igmp_snooping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading igmp_snooping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("igmp_snooping"); ok {
			if err = d.Set("igmp_snooping", flattenSwitchControllerManagedSwitchIgmpSnooping(o["igmp-snooping"], d, "igmp_snooping")); err != nil {
				if vv, ok := fortiAPIPatch(o["igmp-snooping"], "SwitchControllerManagedSwitch-IgmpSnooping"); ok {
					if err = d.Set("igmp_snooping", vv); err != nil {
						return fmt.Errorf("Error reading igmp_snooping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading igmp_snooping: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ip_source_guard", flattenSwitchControllerManagedSwitchIpSourceGuard(o["ip-source-guard"], d, "ip_source_guard")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip-source-guard"], "SwitchControllerManagedSwitch-IpSourceGuard"); ok {
				if err = d.Set("ip_source_guard", vv); err != nil {
					return fmt.Errorf("Error reading ip_source_guard: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip_source_guard: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip_source_guard"); ok {
			if err = d.Set("ip_source_guard", flattenSwitchControllerManagedSwitchIpSourceGuard(o["ip-source-guard"], d, "ip_source_guard")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip-source-guard"], "SwitchControllerManagedSwitch-IpSourceGuard"); ok {
					if err = d.Set("ip_source_guard", vv); err != nil {
						return fmt.Errorf("Error reading ip_source_guard: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip_source_guard: %v", err)
				}
			}
		}
	}

	if err = d.Set("l3_discovered", flattenSwitchControllerManagedSwitchL3Discovered(o["l3-discovered"], d, "l3_discovered")); err != nil {
		if vv, ok := fortiAPIPatch(o["l3-discovered"], "SwitchControllerManagedSwitch-L3Discovered"); ok {
			if err = d.Set("l3_discovered", vv); err != nil {
				return fmt.Errorf("Error reading l3_discovered: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading l3_discovered: %v", err)
		}
	}

	if err = d.Set("max_allowed_trunk_members", flattenSwitchControllerManagedSwitchMaxAllowedTrunkMembers(o["max-allowed-trunk-members"], d, "max_allowed_trunk_members")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-allowed-trunk-members"], "SwitchControllerManagedSwitch-MaxAllowedTrunkMembers"); ok {
			if err = d.Set("max_allowed_trunk_members", vv); err != nil {
				return fmt.Errorf("Error reading max_allowed_trunk_members: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_allowed_trunk_members: %v", err)
		}
	}

	if err = d.Set("mclag_igmp_snooping_aware", flattenSwitchControllerManagedSwitchMclagIgmpSnoopingAware(o["mclag-igmp-snooping-aware"], d, "mclag_igmp_snooping_aware")); err != nil {
		if vv, ok := fortiAPIPatch(o["mclag-igmp-snooping-aware"], "SwitchControllerManagedSwitch-MclagIgmpSnoopingAware"); ok {
			if err = d.Set("mclag_igmp_snooping_aware", vv); err != nil {
				return fmt.Errorf("Error reading mclag_igmp_snooping_aware: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mclag_igmp_snooping_aware: %v", err)
		}
	}

	if err = d.Set("mgmt_mode", flattenSwitchControllerManagedSwitchMgmtMode(o["mgmt-mode"], d, "mgmt_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mgmt-mode"], "SwitchControllerManagedSwitch-MgmtMode"); ok {
			if err = d.Set("mgmt_mode", vv); err != nil {
				return fmt.Errorf("Error reading mgmt_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mgmt_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("mirror", flattenSwitchControllerManagedSwitchMirror(o["mirror"], d, "mirror")); err != nil {
			if vv, ok := fortiAPIPatch(o["mirror"], "SwitchControllerManagedSwitch-Mirror"); ok {
				if err = d.Set("mirror", vv); err != nil {
					return fmt.Errorf("Error reading mirror: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mirror: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mirror"); ok {
			if err = d.Set("mirror", flattenSwitchControllerManagedSwitchMirror(o["mirror"], d, "mirror")); err != nil {
				if vv, ok := fortiAPIPatch(o["mirror"], "SwitchControllerManagedSwitch-Mirror"); ok {
					if err = d.Set("mirror", vv); err != nil {
						return fmt.Errorf("Error reading mirror: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mirror: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenSwitchControllerManagedSwitchName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerManagedSwitch-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("override_snmp_community", flattenSwitchControllerManagedSwitchOverrideSnmpCommunity(o["override-snmp-community"], d, "override_snmp_community")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-snmp-community"], "SwitchControllerManagedSwitch-OverrideSnmpCommunity"); ok {
			if err = d.Set("override_snmp_community", vv); err != nil {
				return fmt.Errorf("Error reading override_snmp_community: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_snmp_community: %v", err)
		}
	}

	if err = d.Set("override_snmp_sysinfo", flattenSwitchControllerManagedSwitchOverrideSnmpSysinfo(o["override-snmp-sysinfo"], d, "override_snmp_sysinfo")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-snmp-sysinfo"], "SwitchControllerManagedSwitch-OverrideSnmpSysinfo"); ok {
			if err = d.Set("override_snmp_sysinfo", vv); err != nil {
				return fmt.Errorf("Error reading override_snmp_sysinfo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_snmp_sysinfo: %v", err)
		}
	}

	if err = d.Set("override_snmp_trap_threshold", flattenSwitchControllerManagedSwitchOverrideSnmpTrapThreshold(o["override-snmp-trap-threshold"], d, "override_snmp_trap_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-snmp-trap-threshold"], "SwitchControllerManagedSwitch-OverrideSnmpTrapThreshold"); ok {
			if err = d.Set("override_snmp_trap_threshold", vv); err != nil {
				return fmt.Errorf("Error reading override_snmp_trap_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_snmp_trap_threshold: %v", err)
		}
	}

	if err = d.Set("override_snmp_user", flattenSwitchControllerManagedSwitchOverrideSnmpUser(o["override-snmp-user"], d, "override_snmp_user")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-snmp-user"], "SwitchControllerManagedSwitch-OverrideSnmpUser"); ok {
			if err = d.Set("override_snmp_user", vv); err != nil {
				return fmt.Errorf("Error reading override_snmp_user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_snmp_user: %v", err)
		}
	}

	if err = d.Set("owner_vdom", flattenSwitchControllerManagedSwitchOwnerVdom(o["owner-vdom"], d, "owner_vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["owner-vdom"], "SwitchControllerManagedSwitch-OwnerVdom"); ok {
			if err = d.Set("owner_vdom", vv); err != nil {
				return fmt.Errorf("Error reading owner_vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading owner_vdom: %v", err)
		}
	}

	if err = d.Set("poe_detection_type", flattenSwitchControllerManagedSwitchPoeDetectionType(o["poe-detection-type"], d, "poe_detection_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["poe-detection-type"], "SwitchControllerManagedSwitch-PoeDetectionType"); ok {
			if err = d.Set("poe_detection_type", vv); err != nil {
				return fmt.Errorf("Error reading poe_detection_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poe_detection_type: %v", err)
		}
	}

	if err = d.Set("poe_lldp_detection", flattenSwitchControllerManagedSwitchPoeLldpDetection(o["poe-lldp-detection"], d, "poe_lldp_detection")); err != nil {
		if vv, ok := fortiAPIPatch(o["poe-lldp-detection"], "SwitchControllerManagedSwitch-PoeLldpDetection"); ok {
			if err = d.Set("poe_lldp_detection", vv); err != nil {
				return fmt.Errorf("Error reading poe_lldp_detection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poe_lldp_detection: %v", err)
		}
	}

	if err = d.Set("poe_pre_standard_detection", flattenSwitchControllerManagedSwitchPoePreStandardDetection(o["poe-pre-standard-detection"], d, "poe_pre_standard_detection")); err != nil {
		if vv, ok := fortiAPIPatch(o["poe-pre-standard-detection"], "SwitchControllerManagedSwitch-PoePreStandardDetection"); ok {
			if err = d.Set("poe_pre_standard_detection", vv); err != nil {
				return fmt.Errorf("Error reading poe_pre_standard_detection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poe_pre_standard_detection: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ports", flattenSwitchControllerManagedSwitchPorts(o["ports"], d, "ports")); err != nil {
			if vv, ok := fortiAPIPatch(o["ports"], "SwitchControllerManagedSwitch-Ports"); ok {
				if err = d.Set("ports", vv); err != nil {
					return fmt.Errorf("Error reading ports: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ports: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ports"); ok {
			if err = d.Set("ports", flattenSwitchControllerManagedSwitchPorts(o["ports"], d, "ports")); err != nil {
				if vv, ok := fortiAPIPatch(o["ports"], "SwitchControllerManagedSwitch-Ports"); ok {
					if err = d.Set("ports", vv); err != nil {
						return fmt.Errorf("Error reading ports: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ports: %v", err)
				}
			}
		}
	}

	if err = d.Set("pre_provisioned", flattenSwitchControllerManagedSwitchPreProvisioned(o["pre-provisioned"], d, "pre_provisioned")); err != nil {
		if vv, ok := fortiAPIPatch(o["pre-provisioned"], "SwitchControllerManagedSwitch-PreProvisioned"); ok {
			if err = d.Set("pre_provisioned", vv); err != nil {
				return fmt.Errorf("Error reading pre_provisioned: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pre_provisioned: %v", err)
		}
	}

	if err = d.Set("ptp_profile", flattenSwitchControllerManagedSwitchPtpProfile(o["ptp-profile"], d, "ptp_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["ptp-profile"], "SwitchControllerManagedSwitch-PtpProfile"); ok {
			if err = d.Set("ptp_profile", vv); err != nil {
				return fmt.Errorf("Error reading ptp_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ptp_profile: %v", err)
		}
	}

	if err = d.Set("ptp_status", flattenSwitchControllerManagedSwitchPtpStatus(o["ptp-status"], d, "ptp_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["ptp-status"], "SwitchControllerManagedSwitch-PtpStatus"); ok {
			if err = d.Set("ptp_status", vv); err != nil {
				return fmt.Errorf("Error reading ptp_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ptp_status: %v", err)
		}
	}

	if err = d.Set("purdue_level", flattenSwitchControllerManagedSwitchPurdueLevel(o["purdue-level"], d, "purdue_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["purdue-level"], "SwitchControllerManagedSwitch-PurdueLevel"); ok {
			if err = d.Set("purdue_level", vv); err != nil {
				return fmt.Errorf("Error reading purdue_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading purdue_level: %v", err)
		}
	}

	if err = d.Set("qos_drop_policy", flattenSwitchControllerManagedSwitchQosDropPolicy(o["qos-drop-policy"], d, "qos_drop_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["qos-drop-policy"], "SwitchControllerManagedSwitch-QosDropPolicy"); ok {
			if err = d.Set("qos_drop_policy", vv); err != nil {
				return fmt.Errorf("Error reading qos_drop_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading qos_drop_policy: %v", err)
		}
	}

	if err = d.Set("qos_red_probability", flattenSwitchControllerManagedSwitchQosRedProbability(o["qos-red-probability"], d, "qos_red_probability")); err != nil {
		if vv, ok := fortiAPIPatch(o["qos-red-probability"], "SwitchControllerManagedSwitch-QosRedProbability"); ok {
			if err = d.Set("qos_red_probability", vv); err != nil {
				return fmt.Errorf("Error reading qos_red_probability: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading qos_red_probability: %v", err)
		}
	}

	if err = d.Set("radius_nas_ip", flattenSwitchControllerManagedSwitchRadiusNasIp(o["radius-nas-ip"], d, "radius_nas_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-nas-ip"], "SwitchControllerManagedSwitch-RadiusNasIp"); ok {
			if err = d.Set("radius_nas_ip", vv); err != nil {
				return fmt.Errorf("Error reading radius_nas_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_nas_ip: %v", err)
		}
	}

	if err = d.Set("radius_nas_ip_override", flattenSwitchControllerManagedSwitchRadiusNasIpOverride(o["radius-nas-ip-override"], d, "radius_nas_ip_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-nas-ip-override"], "SwitchControllerManagedSwitch-RadiusNasIpOverride"); ok {
			if err = d.Set("radius_nas_ip_override", vv); err != nil {
				return fmt.Errorf("Error reading radius_nas_ip_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_nas_ip_override: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("remote_log", flattenSwitchControllerManagedSwitchRemoteLog(o["remote-log"], d, "remote_log")); err != nil {
			if vv, ok := fortiAPIPatch(o["remote-log"], "SwitchControllerManagedSwitch-RemoteLog"); ok {
				if err = d.Set("remote_log", vv); err != nil {
					return fmt.Errorf("Error reading remote_log: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading remote_log: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("remote_log"); ok {
			if err = d.Set("remote_log", flattenSwitchControllerManagedSwitchRemoteLog(o["remote-log"], d, "remote_log")); err != nil {
				if vv, ok := fortiAPIPatch(o["remote-log"], "SwitchControllerManagedSwitch-RemoteLog"); ok {
					if err = d.Set("remote_log", vv); err != nil {
						return fmt.Errorf("Error reading remote_log: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading remote_log: %v", err)
				}
			}
		}
	}

	if err = d.Set("route_offload", flattenSwitchControllerManagedSwitchRouteOffload(o["route-offload"], d, "route_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-offload"], "SwitchControllerManagedSwitch-RouteOffload"); ok {
			if err = d.Set("route_offload", vv); err != nil {
				return fmt.Errorf("Error reading route_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_offload: %v", err)
		}
	}

	if err = d.Set("route_offload_mclag", flattenSwitchControllerManagedSwitchRouteOffloadMclag(o["route-offload-mclag"], d, "route_offload_mclag")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-offload-mclag"], "SwitchControllerManagedSwitch-RouteOffloadMclag"); ok {
			if err = d.Set("route_offload_mclag", vv); err != nil {
				return fmt.Errorf("Error reading route_offload_mclag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_offload_mclag: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("route_offload_router", flattenSwitchControllerManagedSwitchRouteOffloadRouter(o["route-offload-router"], d, "route_offload_router")); err != nil {
			if vv, ok := fortiAPIPatch(o["route-offload-router"], "SwitchControllerManagedSwitch-RouteOffloadRouter"); ok {
				if err = d.Set("route_offload_router", vv); err != nil {
					return fmt.Errorf("Error reading route_offload_router: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading route_offload_router: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("route_offload_router"); ok {
			if err = d.Set("route_offload_router", flattenSwitchControllerManagedSwitchRouteOffloadRouter(o["route-offload-router"], d, "route_offload_router")); err != nil {
				if vv, ok := fortiAPIPatch(o["route-offload-router"], "SwitchControllerManagedSwitch-RouteOffloadRouter"); ok {
					if err = d.Set("route_offload_router", vv); err != nil {
						return fmt.Errorf("Error reading route_offload_router: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading route_offload_router: %v", err)
				}
			}
		}
	}

	if err = d.Set("sn", flattenSwitchControllerManagedSwitchSn(o["sn"], d, "sn")); err != nil {
		if vv, ok := fortiAPIPatch(o["sn"], "SwitchControllerManagedSwitch-Sn"); ok {
			if err = d.Set("sn", vv); err != nil {
				return fmt.Errorf("Error reading sn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sn: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("snmp_community", flattenSwitchControllerManagedSwitchSnmpCommunity(o["snmp-community"], d, "snmp_community")); err != nil {
			if vv, ok := fortiAPIPatch(o["snmp-community"], "SwitchControllerManagedSwitch-SnmpCommunity"); ok {
				if err = d.Set("snmp_community", vv); err != nil {
					return fmt.Errorf("Error reading snmp_community: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading snmp_community: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("snmp_community"); ok {
			if err = d.Set("snmp_community", flattenSwitchControllerManagedSwitchSnmpCommunity(o["snmp-community"], d, "snmp_community")); err != nil {
				if vv, ok := fortiAPIPatch(o["snmp-community"], "SwitchControllerManagedSwitch-SnmpCommunity"); ok {
					if err = d.Set("snmp_community", vv); err != nil {
						return fmt.Errorf("Error reading snmp_community: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading snmp_community: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("snmp_sysinfo", flattenSwitchControllerManagedSwitchSnmpSysinfo(o["snmp-sysinfo"], d, "snmp_sysinfo")); err != nil {
			if vv, ok := fortiAPIPatch(o["snmp-sysinfo"], "SwitchControllerManagedSwitch-SnmpSysinfo"); ok {
				if err = d.Set("snmp_sysinfo", vv); err != nil {
					return fmt.Errorf("Error reading snmp_sysinfo: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading snmp_sysinfo: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("snmp_sysinfo"); ok {
			if err = d.Set("snmp_sysinfo", flattenSwitchControllerManagedSwitchSnmpSysinfo(o["snmp-sysinfo"], d, "snmp_sysinfo")); err != nil {
				if vv, ok := fortiAPIPatch(o["snmp-sysinfo"], "SwitchControllerManagedSwitch-SnmpSysinfo"); ok {
					if err = d.Set("snmp_sysinfo", vv); err != nil {
						return fmt.Errorf("Error reading snmp_sysinfo: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading snmp_sysinfo: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("snmp_trap_threshold", flattenSwitchControllerManagedSwitchSnmpTrapThreshold(o["snmp-trap-threshold"], d, "snmp_trap_threshold")); err != nil {
			if vv, ok := fortiAPIPatch(o["snmp-trap-threshold"], "SwitchControllerManagedSwitch-SnmpTrapThreshold"); ok {
				if err = d.Set("snmp_trap_threshold", vv); err != nil {
					return fmt.Errorf("Error reading snmp_trap_threshold: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading snmp_trap_threshold: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("snmp_trap_threshold"); ok {
			if err = d.Set("snmp_trap_threshold", flattenSwitchControllerManagedSwitchSnmpTrapThreshold(o["snmp-trap-threshold"], d, "snmp_trap_threshold")); err != nil {
				if vv, ok := fortiAPIPatch(o["snmp-trap-threshold"], "SwitchControllerManagedSwitch-SnmpTrapThreshold"); ok {
					if err = d.Set("snmp_trap_threshold", vv); err != nil {
						return fmt.Errorf("Error reading snmp_trap_threshold: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading snmp_trap_threshold: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("snmp_user", flattenSwitchControllerManagedSwitchSnmpUser(o["snmp-user"], d, "snmp_user")); err != nil {
			if vv, ok := fortiAPIPatch(o["snmp-user"], "SwitchControllerManagedSwitch-SnmpUser"); ok {
				if err = d.Set("snmp_user", vv); err != nil {
					return fmt.Errorf("Error reading snmp_user: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading snmp_user: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("snmp_user"); ok {
			if err = d.Set("snmp_user", flattenSwitchControllerManagedSwitchSnmpUser(o["snmp-user"], d, "snmp_user")); err != nil {
				if vv, ok := fortiAPIPatch(o["snmp-user"], "SwitchControllerManagedSwitch-SnmpUser"); ok {
					if err = d.Set("snmp_user", vv); err != nil {
						return fmt.Errorf("Error reading snmp_user: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading snmp_user: %v", err)
				}
			}
		}
	}

	if err = d.Set("staged_image_version", flattenSwitchControllerManagedSwitchStagedImageVersion(o["staged-image-version"], d, "staged_image_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["staged-image-version"], "SwitchControllerManagedSwitch-StagedImageVersion"); ok {
			if err = d.Set("staged_image_version", vv); err != nil {
				return fmt.Errorf("Error reading staged_image_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading staged_image_version: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("static_mac", flattenSwitchControllerManagedSwitchStaticMac(o["static-mac"], d, "static_mac")); err != nil {
			if vv, ok := fortiAPIPatch(o["static-mac"], "SwitchControllerManagedSwitch-StaticMac"); ok {
				if err = d.Set("static_mac", vv); err != nil {
					return fmt.Errorf("Error reading static_mac: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading static_mac: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("static_mac"); ok {
			if err = d.Set("static_mac", flattenSwitchControllerManagedSwitchStaticMac(o["static-mac"], d, "static_mac")); err != nil {
				if vv, ok := fortiAPIPatch(o["static-mac"], "SwitchControllerManagedSwitch-StaticMac"); ok {
					if err = d.Set("static_mac", vv); err != nil {
						return fmt.Errorf("Error reading static_mac: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading static_mac: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("storm_control", flattenSwitchControllerManagedSwitchStormControl(o["storm-control"], d, "storm_control")); err != nil {
			if vv, ok := fortiAPIPatch(o["storm-control"], "SwitchControllerManagedSwitch-StormControl"); ok {
				if err = d.Set("storm_control", vv); err != nil {
					return fmt.Errorf("Error reading storm_control: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading storm_control: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("storm_control"); ok {
			if err = d.Set("storm_control", flattenSwitchControllerManagedSwitchStormControl(o["storm-control"], d, "storm_control")); err != nil {
				if vv, ok := fortiAPIPatch(o["storm-control"], "SwitchControllerManagedSwitch-StormControl"); ok {
					if err = d.Set("storm_control", vv); err != nil {
						return fmt.Errorf("Error reading storm_control: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading storm_control: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("stp_instance", flattenSwitchControllerManagedSwitchStpInstance(o["stp-instance"], d, "stp_instance")); err != nil {
			if vv, ok := fortiAPIPatch(o["stp-instance"], "SwitchControllerManagedSwitch-StpInstance"); ok {
				if err = d.Set("stp_instance", vv); err != nil {
					return fmt.Errorf("Error reading stp_instance: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading stp_instance: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("stp_instance"); ok {
			if err = d.Set("stp_instance", flattenSwitchControllerManagedSwitchStpInstance(o["stp-instance"], d, "stp_instance")); err != nil {
				if vv, ok := fortiAPIPatch(o["stp-instance"], "SwitchControllerManagedSwitch-StpInstance"); ok {
					if err = d.Set("stp_instance", vv); err != nil {
						return fmt.Errorf("Error reading stp_instance: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading stp_instance: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("stp_settings", flattenSwitchControllerManagedSwitchStpSettings(o["stp-settings"], d, "stp_settings")); err != nil {
			if vv, ok := fortiAPIPatch(o["stp-settings"], "SwitchControllerManagedSwitch-StpSettings"); ok {
				if err = d.Set("stp_settings", vv); err != nil {
					return fmt.Errorf("Error reading stp_settings: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading stp_settings: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("stp_settings"); ok {
			if err = d.Set("stp_settings", flattenSwitchControllerManagedSwitchStpSettings(o["stp-settings"], d, "stp_settings")); err != nil {
				if vv, ok := fortiAPIPatch(o["stp-settings"], "SwitchControllerManagedSwitch-StpSettings"); ok {
					if err = d.Set("stp_settings", vv); err != nil {
						return fmt.Errorf("Error reading stp_settings: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading stp_settings: %v", err)
				}
			}
		}
	}

	if err = d.Set("switch_device_tag", flattenSwitchControllerManagedSwitchSwitchDeviceTag(o["switch-device-tag"], d, "switch_device_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-device-tag"], "SwitchControllerManagedSwitch-SwitchDeviceTag"); ok {
			if err = d.Set("switch_device_tag", vv); err != nil {
				return fmt.Errorf("Error reading switch_device_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_device_tag: %v", err)
		}
	}

	if err = d.Set("switch_dhcp_opt43_key", flattenSwitchControllerManagedSwitchSwitchDhcpOpt43Key(o["switch-dhcp_opt43_key"], d, "switch_dhcp_opt43_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-dhcp_opt43_key"], "SwitchControllerManagedSwitch-SwitchDhcpOpt43Key"); ok {
			if err = d.Set("switch_dhcp_opt43_key", vv); err != nil {
				return fmt.Errorf("Error reading switch_dhcp_opt43_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_dhcp_opt43_key: %v", err)
		}
	}

	if err = d.Set("switch_id", flattenSwitchControllerManagedSwitchSwitchId(o["switch-id"], d, "switch_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-id"], "SwitchControllerManagedSwitch-SwitchId"); ok {
			if err = d.Set("switch_id", vv); err != nil {
				return fmt.Errorf("Error reading switch_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_id: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("switch_log", flattenSwitchControllerManagedSwitchSwitchLog(o["switch-log"], d, "switch_log")); err != nil {
			if vv, ok := fortiAPIPatch(o["switch-log"], "SwitchControllerManagedSwitch-SwitchLog"); ok {
				if err = d.Set("switch_log", vv); err != nil {
					return fmt.Errorf("Error reading switch_log: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading switch_log: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("switch_log"); ok {
			if err = d.Set("switch_log", flattenSwitchControllerManagedSwitchSwitchLog(o["switch-log"], d, "switch_log")); err != nil {
				if vv, ok := fortiAPIPatch(o["switch-log"], "SwitchControllerManagedSwitch-SwitchLog"); ok {
					if err = d.Set("switch_log", vv); err != nil {
						return fmt.Errorf("Error reading switch_log: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading switch_log: %v", err)
				}
			}
		}
	}

	if err = d.Set("switch_profile", flattenSwitchControllerManagedSwitchSwitchProfile(o["switch-profile"], d, "switch_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-profile"], "SwitchControllerManagedSwitch-SwitchProfile"); ok {
			if err = d.Set("switch_profile", vv); err != nil {
				return fmt.Errorf("Error reading switch_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_profile: %v", err)
		}
	}

	if err = d.Set("tdr_supported", flattenSwitchControllerManagedSwitchTdrSupported(o["tdr-supported"], d, "tdr_supported")); err != nil {
		if vv, ok := fortiAPIPatch(o["tdr-supported"], "SwitchControllerManagedSwitch-TdrSupported"); ok {
			if err = d.Set("tdr_supported", vv); err != nil {
				return fmt.Errorf("Error reading tdr_supported: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tdr_supported: %v", err)
		}
	}

	if err = d.Set("tunnel_discovered", flattenSwitchControllerManagedSwitchTunnelDiscovered(o["tunnel-discovered"], d, "tunnel_discovered")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-discovered"], "SwitchControllerManagedSwitch-TunnelDiscovered"); ok {
			if err = d.Set("tunnel_discovered", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_discovered: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_discovered: %v", err)
		}
	}

	if err = d.Set("type", flattenSwitchControllerManagedSwitchType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "SwitchControllerManagedSwitch-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("version", flattenSwitchControllerManagedSwitchVersion(o["version"], d, "version")); err != nil {
		if vv, ok := fortiAPIPatch(o["version"], "SwitchControllerManagedSwitch-Version"); ok {
			if err = d.Set("version", vv); err != nil {
				return fmt.Errorf("Error reading version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading version: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("vlan", flattenSwitchControllerManagedSwitchVlan(o["vlan"], d, "vlan")); err != nil {
			if vv, ok := fortiAPIPatch(o["vlan"], "SwitchControllerManagedSwitch-Vlan"); ok {
				if err = d.Set("vlan", vv); err != nil {
					return fmt.Errorf("Error reading vlan: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading vlan: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vlan"); ok {
			if err = d.Set("vlan", flattenSwitchControllerManagedSwitchVlan(o["vlan"], d, "vlan")); err != nil {
				if vv, ok := fortiAPIPatch(o["vlan"], "SwitchControllerManagedSwitch-Vlan"); ok {
					if err = d.Set("vlan", vv); err != nil {
						return fmt.Errorf("Error reading vlan: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading vlan: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSwitchControllerManagedSwitchFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerManagedSwitch8021XSettings(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "link_down_auth"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["link-down-auth"], _ = expandSwitchControllerManagedSwitch8021XSettingsLinkDownAuth(d, i["link_down_auth"], pre_append)
	}
	pre_append = pre + ".0." + "local_override"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["local-override"], _ = expandSwitchControllerManagedSwitch8021XSettingsLocalOverride(d, i["local_override"], pre_append)
	}
	pre_append = pre + ".0." + "mab_reauth"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mab-reauth"], _ = expandSwitchControllerManagedSwitch8021XSettingsMabReauth(d, i["mab_reauth"], pre_append)
	}
	pre_append = pre + ".0." + "mac_called_station_delimiter"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mac-called-station-delimiter"], _ = expandSwitchControllerManagedSwitch8021XSettingsMacCalledStationDelimiter(d, i["mac_called_station_delimiter"], pre_append)
	}
	pre_append = pre + ".0." + "mac_calling_station_delimiter"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mac-calling-station-delimiter"], _ = expandSwitchControllerManagedSwitch8021XSettingsMacCallingStationDelimiter(d, i["mac_calling_station_delimiter"], pre_append)
	}
	pre_append = pre + ".0." + "mac_case"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mac-case"], _ = expandSwitchControllerManagedSwitch8021XSettingsMacCase(d, i["mac_case"], pre_append)
	}
	pre_append = pre + ".0." + "mac_password_delimiter"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mac-password-delimiter"], _ = expandSwitchControllerManagedSwitch8021XSettingsMacPasswordDelimiter(d, i["mac_password_delimiter"], pre_append)
	}
	pre_append = pre + ".0." + "mac_username_delimiter"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mac-username-delimiter"], _ = expandSwitchControllerManagedSwitch8021XSettingsMacUsernameDelimiter(d, i["mac_username_delimiter"], pre_append)
	}
	pre_append = pre + ".0." + "max_reauth_attempt"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-reauth-attempt"], _ = expandSwitchControllerManagedSwitch8021XSettingsMaxReauthAttempt(d, i["max_reauth_attempt"], pre_append)
	}
	pre_append = pre + ".0." + "reauth_period"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["reauth-period"], _ = expandSwitchControllerManagedSwitch8021XSettingsReauthPeriod(d, i["reauth_period"], pre_append)
	}
	pre_append = pre + ".0." + "tx_period"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tx-period"], _ = expandSwitchControllerManagedSwitch8021XSettingsTxPeriod(d, i["tx_period"], pre_append)
	}

	return result, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsLinkDownAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsLocalOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsMabReauth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsMacCalledStationDelimiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsMacCallingStationDelimiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsMacCase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsMacPasswordDelimiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsMacUsernameDelimiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsMaxReauthAttempt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsReauthPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsTxPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPlatform(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchAccessProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchCustomCommand(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "command_entry"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["command-entry"], _ = expandSwitchControllerManagedSwitchCustomCommandCommandEntry(d, i["command_entry"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "command_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["command-name"], _ = expandSwitchControllerManagedSwitchCustomCommandCommandName(d, i["command_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchCustomCommandCommandEntry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchCustomCommandCommandName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchDelayedRestartTrigger(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchDhcpServerAccessList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchDhcpSnoopingStaticClient(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandSwitchControllerManagedSwitchDhcpSnoopingStaticClientIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac"], _ = expandSwitchControllerManagedSwitchDhcpSnoopingStaticClientMac(d, i["mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSwitchControllerManagedSwitchDhcpSnoopingStaticClientName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandSwitchControllerManagedSwitchDhcpSnoopingStaticClientPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlan"], _ = expandSwitchControllerManagedSwitchDhcpSnoopingStaticClientVlan(d, i["vlan"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchDhcpSnoopingStaticClientIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchDhcpSnoopingStaticClientMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchDhcpSnoopingStaticClientName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchDhcpSnoopingStaticClientPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchDhcpSnoopingStaticClientVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchDirectlyConnected(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchDynamicCapability(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchDynamicallyDiscovered(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchFirmwareProvision(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchFirmwareProvisionLatest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchFirmwareProvisionVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchFlowIdentity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchFswWan1Admin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchFswWan1Peer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchFswWan2Admin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchFswWan2Peer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIgmpSnooping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "aging_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["aging-time"], _ = expandSwitchControllerManagedSwitchIgmpSnoopingAgingTime(d, i["aging_time"], pre_append)
	}
	pre_append = pre + ".0." + "flood_unknown_multicast"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["flood-unknown-multicast"], _ = expandSwitchControllerManagedSwitchIgmpSnoopingFloodUnknownMulticast(d, i["flood_unknown_multicast"], pre_append)
	}
	pre_append = pre + ".0." + "local_override"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["local-override"], _ = expandSwitchControllerManagedSwitchIgmpSnoopingLocalOverride(d, i["local_override"], pre_append)
	}
	pre_append = pre + ".0." + "vlans"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandSwitchControllerManagedSwitchIgmpSnoopingVlans(d, i["vlans"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["vlans"] = t
		}
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingAgingTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingFloodUnknownMulticast(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingLocalOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingVlans(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "proxy"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["proxy"], _ = expandSwitchControllerManagedSwitchIgmpSnoopingVlansProxy(d, i["proxy"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "querier"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["querier"], _ = expandSwitchControllerManagedSwitchIgmpSnoopingVlansQuerier(d, i["querier"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "querier_addr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["querier-addr"], _ = expandSwitchControllerManagedSwitchIgmpSnoopingVlansQuerierAddr(d, i["querier_addr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["version"], _ = expandSwitchControllerManagedSwitchIgmpSnoopingVlansVersion(d, i["version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlan-name"], _ = expandSwitchControllerManagedSwitchIgmpSnoopingVlansVlanName(d, i["vlan_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingVlansProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingVlansQuerier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingVlansQuerierAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingVlansVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingVlansVlanName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchIpSourceGuard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "binding_entry"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandSwitchControllerManagedSwitchIpSourceGuardBindingEntry(d, i["binding_entry"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["binding-entry"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandSwitchControllerManagedSwitchIpSourceGuardDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandSwitchControllerManagedSwitchIpSourceGuardPort(d, i["port"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchIpSourceGuardBindingEntry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entry_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["entry-name"], _ = expandSwitchControllerManagedSwitchIpSourceGuardBindingEntryEntryName(d, i["entry_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandSwitchControllerManagedSwitchIpSourceGuardBindingEntryIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac"], _ = expandSwitchControllerManagedSwitchIpSourceGuardBindingEntryMac(d, i["mac"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchIpSourceGuardBindingEntryEntryName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIpSourceGuardBindingEntryIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIpSourceGuardBindingEntryMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIpSourceGuardDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIpSourceGuardPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchL3Discovered(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchMaxAllowedTrunkMembers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchMclagIgmpSnoopingAware(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchMgmtMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchMirror(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dst"], _ = expandSwitchControllerManagedSwitchMirrorDst(d, i["dst"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSwitchControllerManagedSwitchMirrorName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_egress"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-egress"], _ = expandSwitchControllerManagedSwitchMirrorSrcEgress(d, i["src_egress"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_ingress"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-ingress"], _ = expandSwitchControllerManagedSwitchMirrorSrcIngress(d, i["src_ingress"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandSwitchControllerManagedSwitchMirrorStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switching_packet"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["switching-packet"], _ = expandSwitchControllerManagedSwitchMirrorSwitchingPacket(d, i["switching_packet"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchMirrorDst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchMirrorName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchMirrorSrcEgress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchMirrorSrcIngress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchMirrorStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchMirrorSwitchingPacket(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchOverrideSnmpCommunity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchOverrideSnmpSysinfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchOverrideSnmpTrapThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchOverrideSnmpUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchOwnerVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPoeDetectionType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPoeLldpDetection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPoePreStandardDetection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["access-mode"], _ = expandSwitchControllerManagedSwitchPortsAccessMode(d, i["access_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "acl_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["acl-group"], _ = expandSwitchControllerManagedSwitchPortsAclGroup(d, i["acl_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "aggregator_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["aggregator-mode"], _ = expandSwitchControllerManagedSwitchPortsAggregatorMode(d, i["aggregator_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allow_arp_monitor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allow-arp-monitor"], _ = expandSwitchControllerManagedSwitchPortsAllowArpMonitor(d, i["allow_arp_monitor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowed_vlans"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowed-vlans"], _ = expandSwitchControllerManagedSwitchPortsAllowedVlans(d, i["allowed_vlans"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowed_vlans_all"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowed-vlans-all"], _ = expandSwitchControllerManagedSwitchPortsAllowedVlansAll(d, i["allowed_vlans_all"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "arp_inspection_trust"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["arp-inspection-trust"], _ = expandSwitchControllerManagedSwitchPortsArpInspectionTrust(d, i["arp_inspection_trust"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authenticated_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["authenticated-port"], _ = expandSwitchControllerManagedSwitchPortsAuthenticatedPort(d, i["authenticated_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bundle"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bundle"], _ = expandSwitchControllerManagedSwitchPortsBundle(d, i["bundle"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandSwitchControllerManagedSwitchPortsDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_snoop_option82_override"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandSwitchControllerManagedSwitchPortsDhcpSnoopOption82Override(d, i["dhcp_snoop_option82_override"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["dhcp-snoop-option82-override"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_snoop_option82_trust"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dhcp-snoop-option82-trust"], _ = expandSwitchControllerManagedSwitchPortsDhcpSnoopOption82Trust(d, i["dhcp_snoop_option82_trust"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_snooping"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dhcp-snooping"], _ = expandSwitchControllerManagedSwitchPortsDhcpSnooping(d, i["dhcp_snooping"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "discard_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["discard-mode"], _ = expandSwitchControllerManagedSwitchPortsDiscardMode(d, i["discard_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dsl_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dsl-profile"], _ = expandSwitchControllerManagedSwitchPortsDslProfile(d, i["dsl_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "edge_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["edge-port"], _ = expandSwitchControllerManagedSwitchPortsEdgePort(d, i["edge_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "export_tags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["export-tags"], _ = expandSwitchControllerManagedSwitchPortsExportTags(d, i["export_tags"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "encrypted_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["encrypted-port"], _ = expandSwitchControllerManagedSwitchPortsEncryptedPort(d, i["encrypted_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "export_to"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["export-to"], _ = expandSwitchControllerManagedSwitchPortsExportTo(d, i["export_to"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "export_to_pool"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["export-to-pool"], _ = expandSwitchControllerManagedSwitchPortsExportToPool(d, i["export_to_pool"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "export_to_pool_flag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["export-to-pool-flag"], _ = expandSwitchControllerManagedSwitchPortsExportToPoolFlag(d, i["export_to_pool_flag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fallback_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fallback-port"], _ = expandSwitchControllerManagedSwitchPortsFallbackPort(d, i["fallback_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fec_capable"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fec-capable"], _ = expandSwitchControllerManagedSwitchPortsFecCapable(d, i["fec_capable"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fec_state"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fec-state"], _ = expandSwitchControllerManagedSwitchPortsFecState(d, i["fec_state"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fgt_peer_device_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fgt-peer-device-name"], _ = expandSwitchControllerManagedSwitchPortsFgtPeerDeviceName(d, i["fgt_peer_device_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fgt_peer_port_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fgt-peer-port-name"], _ = expandSwitchControllerManagedSwitchPortsFgtPeerPortName(d, i["fgt_peer_port_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fiber_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fiber-port"], _ = expandSwitchControllerManagedSwitchPortsFiberPort(d, i["fiber_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["flags"], _ = expandSwitchControllerManagedSwitchPortsFlags(d, i["flags"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flap_duration"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["flap-duration"], _ = expandSwitchControllerManagedSwitchPortsFlapDuration(d, i["flap_duration"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flap_rate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["flap-rate"], _ = expandSwitchControllerManagedSwitchPortsFlapRate(d, i["flap_rate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flap_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["flap-timeout"], _ = expandSwitchControllerManagedSwitchPortsFlapTimeout(d, i["flap_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flapguard"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["flapguard"], _ = expandSwitchControllerManagedSwitchPortsFlapguard(d, i["flapguard"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flow_control"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["flow-control"], _ = expandSwitchControllerManagedSwitchPortsFlowControl(d, i["flow_control"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortilink_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fortilink-port"], _ = expandSwitchControllerManagedSwitchPortsFortilinkPort(d, i["fortilink_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmp_snooping"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["igmp-snooping"], _ = expandSwitchControllerManagedSwitchPortsIgmpSnooping(d, i["igmp_snooping"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortiswitch_acls"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fortiswitch-acls"], _ = expandSwitchControllerManagedSwitchPortsFortiswitchAcls(d, i["fortiswitch_acls"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmp_snooping_flood_reports"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["igmp-snooping-flood-reports"], _ = expandSwitchControllerManagedSwitchPortsIgmpSnoopingFloodReports(d, i["igmp_snooping_flood_reports"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmps_flood_reports"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["igmps-flood-reports"], _ = expandSwitchControllerManagedSwitchPortsIgmpsFloodReports(d, i["igmps_flood_reports"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmps_flood_traffic"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["igmps-flood-traffic"], _ = expandSwitchControllerManagedSwitchPortsIgmpsFloodTraffic(d, i["igmps_flood_traffic"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_tags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface-tags"], _ = expandSwitchControllerManagedSwitchPortsInterfaceTags(d, i["interface_tags"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_source_guard"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip-source-guard"], _ = expandSwitchControllerManagedSwitchPortsIpSourceGuard(d, i["ip_source_guard"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "isl_local_trunk_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["isl-local-trunk-name"], _ = expandSwitchControllerManagedSwitchPortsIslLocalTrunkName(d, i["isl_local_trunk_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "isl_peer_device_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["isl-peer-device-name"], _ = expandSwitchControllerManagedSwitchPortsIslPeerDeviceName(d, i["isl_peer_device_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "isl_peer_device_sn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["isl-peer-device-sn"], _ = expandSwitchControllerManagedSwitchPortsIslPeerDeviceSn(d, i["isl_peer_device_sn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "isl_peer_port_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["isl-peer-port-name"], _ = expandSwitchControllerManagedSwitchPortsIslPeerPortName(d, i["isl_peer_port_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lacp_speed"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["lacp-speed"], _ = expandSwitchControllerManagedSwitchPortsLacpSpeed(d, i["lacp_speed"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "learning_limit"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["learning-limit"], _ = expandSwitchControllerManagedSwitchPortsLearningLimit(d, i["learning_limit"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["link-status"], _ = expandSwitchControllerManagedSwitchPortsLinkStatus(d, i["link_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lldp_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["lldp-profile"], _ = expandSwitchControllerManagedSwitchPortsLldpProfile(d, i["lldp_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lldp_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["lldp-status"], _ = expandSwitchControllerManagedSwitchPortsLldpStatus(d, i["lldp_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "loop_guard"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["loop-guard"], _ = expandSwitchControllerManagedSwitchPortsLoopGuard(d, i["loop_guard"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "loop_guard_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["loop-guard-timeout"], _ = expandSwitchControllerManagedSwitchPortsLoopGuardTimeout(d, i["loop_guard_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_addr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac-addr"], _ = expandSwitchControllerManagedSwitchPortsMacAddr(d, i["mac_addr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "matched_dpp_intf_tags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["matched-dpp-intf-tags"], _ = expandSwitchControllerManagedSwitchPortsMatchedDppIntfTags(d, i["matched_dpp_intf_tags"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "matched_dpp_policy"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["matched-dpp-policy"], _ = expandSwitchControllerManagedSwitchPortsMatchedDppPolicy(d, i["matched_dpp_policy"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_bundle"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-bundle"], _ = expandSwitchControllerManagedSwitchPortsMaxBundle(d, i["max_bundle"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mcast_snooping_flood_traffic"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mcast-snooping-flood-traffic"], _ = expandSwitchControllerManagedSwitchPortsMcastSnoopingFloodTraffic(d, i["mcast_snooping_flood_traffic"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mclag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mclag"], _ = expandSwitchControllerManagedSwitchPortsMclag(d, i["mclag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mclag_icl_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mclag-icl-port"], _ = expandSwitchControllerManagedSwitchPortsMclagIclPort(d, i["mclag_icl_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "media_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["media-type"], _ = expandSwitchControllerManagedSwitchPortsMediaType(d, i["media_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member_withdrawal_behavior"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["member-withdrawal-behavior"], _ = expandSwitchControllerManagedSwitchPortsMemberWithdrawalBehavior(d, i["member_withdrawal_behavior"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["members"], _ = expandSwitchControllerManagedSwitchPortsMembers(d, i["members"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min_bundle"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["min-bundle"], _ = expandSwitchControllerManagedSwitchPortsMinBundle(d, i["min_bundle"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mode"], _ = expandSwitchControllerManagedSwitchPortsMode(d, i["mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "p2p_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["p2p-port"], _ = expandSwitchControllerManagedSwitchPortsP2PPort(d, i["p2p_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_sample_rate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packet-sample-rate"], _ = expandSwitchControllerManagedSwitchPortsPacketSampleRate(d, i["packet_sample_rate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_sampler"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packet-sampler"], _ = expandSwitchControllerManagedSwitchPortsPacketSampler(d, i["packet_sampler"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pause_meter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pause-meter"], _ = expandSwitchControllerManagedSwitchPortsPauseMeter(d, i["pause_meter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pause_meter_resume"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pause-meter-resume"], _ = expandSwitchControllerManagedSwitchPortsPauseMeterResume(d, i["pause_meter_resume"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_capable"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["poe-capable"], _ = expandSwitchControllerManagedSwitchPortsPoeCapable(d, i["poe_capable"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_max_power"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["poe-max-power"], _ = expandSwitchControllerManagedSwitchPortsPoeMaxPower(d, i["poe_max_power"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_mode_bt_cabable"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["poe-mode-bt-cabable"], _ = expandSwitchControllerManagedSwitchPortsPoeModeBtCabable(d, i["poe_mode_bt_cabable"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_port_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["poe-port-mode"], _ = expandSwitchControllerManagedSwitchPortsPoePortMode(d, i["poe_port_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_port_power"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["poe-port-power"], _ = expandSwitchControllerManagedSwitchPortsPoePortPower(d, i["poe_port_power"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_port_priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["poe-port-priority"], _ = expandSwitchControllerManagedSwitchPortsPoePortPriority(d, i["poe_port_priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_pre_standard_detection"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["poe-pre-standard-detection"], _ = expandSwitchControllerManagedSwitchPortsPoePreStandardDetection(d, i["poe_pre_standard_detection"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_standard"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["poe-standard"], _ = expandSwitchControllerManagedSwitchPortsPoeStandard(d, i["poe_standard"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["poe-status"], _ = expandSwitchControllerManagedSwitchPortsPoeStatus(d, i["poe_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port-name"], _ = expandSwitchControllerManagedSwitchPortsPortName(d, i["port_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_number"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port-number"], _ = expandSwitchControllerManagedSwitchPortsPortNumber(d, i["port_number"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_owner"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port-owner"], _ = expandSwitchControllerManagedSwitchPortsPortOwner(d, i["port_owner"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_policy"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port-policy"], _ = expandSwitchControllerManagedSwitchPortsPortPolicy(d, i["port_policy"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_prefix_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port-prefix-type"], _ = expandSwitchControllerManagedSwitchPortsPortPrefixType(d, i["port_prefix_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_security_policy"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port-security-policy"], _ = expandSwitchControllerManagedSwitchPortsPortSecurityPolicy(d, i["port_security_policy"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_selection_criteria"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port-selection-criteria"], _ = expandSwitchControllerManagedSwitchPortsPortSelectionCriteria(d, i["port_selection_criteria"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ptp_policy"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ptp-policy"], _ = expandSwitchControllerManagedSwitchPortsPtpPolicy(d, i["ptp_policy"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ptp_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ptp-status"], _ = expandSwitchControllerManagedSwitchPortsPtpStatus(d, i["ptp_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "qos_policy"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["qos-policy"], _ = expandSwitchControllerManagedSwitchPortsQosPolicy(d, i["qos_policy"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restricted_auth_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["restricted-auth-port"], _ = expandSwitchControllerManagedSwitchPortsRestrictedAuthPort(d, i["restricted_auth_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rpvst_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rpvst-port"], _ = expandSwitchControllerManagedSwitchPortsRpvstPort(d, i["rpvst_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sample_direction"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sample-direction"], _ = expandSwitchControllerManagedSwitchPortsSampleDirection(d, i["sample_direction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sflow_counter_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sflow-counter-interval"], _ = expandSwitchControllerManagedSwitchPortsSflowCounterInterval(d, i["sflow_counter_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "speed"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["speed"], _ = expandSwitchControllerManagedSwitchPortsSpeed(d, i["speed"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "speed_mask"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["speed-mask"], _ = expandSwitchControllerManagedSwitchPortsSpeedMask(d, i["speed_mask"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stacking_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["stacking-port"], _ = expandSwitchControllerManagedSwitchPortsStackingPort(d, i["stacking_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandSwitchControllerManagedSwitchPortsStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sticky_mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sticky-mac"], _ = expandSwitchControllerManagedSwitchPortsStickyMac(d, i["sticky_mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "storm_control_policy"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["storm-control-policy"], _ = expandSwitchControllerManagedSwitchPortsStormControlPolicy(d, i["storm_control_policy"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stp_bpdu_guard"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["stp-bpdu-guard"], _ = expandSwitchControllerManagedSwitchPortsStpBpduGuard(d, i["stp_bpdu_guard"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stp_bpdu_guard_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["stp-bpdu-guard-timeout"], _ = expandSwitchControllerManagedSwitchPortsStpBpduGuardTimeout(d, i["stp_bpdu_guard_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stp_root_guard"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["stp-root-guard"], _ = expandSwitchControllerManagedSwitchPortsStpRootGuard(d, i["stp_root_guard"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stp_state"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["stp-state"], _ = expandSwitchControllerManagedSwitchPortsStpState(d, i["stp_state"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["switch-id"], _ = expandSwitchControllerManagedSwitchPortsSwitchId(d, i["switch_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trunk_member"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["trunk-member"], _ = expandSwitchControllerManagedSwitchPortsTrunkMember(d, i["trunk_member"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandSwitchControllerManagedSwitchPortsType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "untagged_vlans"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["untagged-vlans"], _ = expandSwitchControllerManagedSwitchPortsUntaggedVlans(d, i["untagged_vlans"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["virtual-port"], _ = expandSwitchControllerManagedSwitchPortsVirtualPort(d, i["virtual_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlan"], _ = expandSwitchControllerManagedSwitchPortsVlan(d, i["vlan"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchPortsAccessMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsAclGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchPortsAggregatorMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsAllowArpMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsAllowedVlans(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchPortsAllowedVlansAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsArpInspectionTrust(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsAuthenticatedPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsBundle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsDhcpSnoopOption82Override(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["circuit-id"], _ = expandSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideCircuitId(d, i["circuit_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remote-id"], _ = expandSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideRemoteId(d, i["remote_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlan-name"], _ = expandSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideVlanName(d, i["vlan_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideCircuitId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideRemoteId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideVlanName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchPortsDhcpSnoopOption82Trust(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsDhcpSnooping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsDiscardMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsDslProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchPortsEdgePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsExportTags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchPortsEncryptedPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsExportTo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchPortsExportToPool(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchPortsExportToPoolFlag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsFallbackPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsFecCapable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsFecState(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsFgtPeerDeviceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsFgtPeerPortName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsFiberPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsFlapDuration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsFlapRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsFlapTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsFlapguard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsFlowControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsFortilinkPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsIgmpSnooping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsFortiswitchAcls(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchPortsIgmpSnoopingFloodReports(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsIgmpsFloodReports(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsIgmpsFloodTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsInterfaceTags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchPortsIpSourceGuard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsIslLocalTrunkName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsIslPeerDeviceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsIslPeerDeviceSn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsIslPeerPortName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsLacpSpeed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsLearningLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsLinkStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsLldpProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchPortsLldpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsLoopGuard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsLoopGuardTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsMacAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsMatchedDppIntfTags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsMatchedDppPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsMaxBundle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsMcastSnoopingFloodTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsMclag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsMclagIclPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsMediaType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsMemberWithdrawalBehavior(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsMembers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchPortsMinBundle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsP2PPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPacketSampleRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPacketSampler(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPauseMeter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPauseMeterResume(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPoeCapable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPoeMaxPower(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPoeModeBtCabable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPoePortMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPoePortPower(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPoePortPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPoePreStandardDetection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPoeStandard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPoeStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPortName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPortNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPortOwner(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPortPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchPortsPortPrefixType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPortSecurityPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchPortsPortSelectionCriteria(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPtpPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchPortsPtpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsQosPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchPortsRestrictedAuthPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsRpvstPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsSampleDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsSflowCounterInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsSpeed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsSpeedMask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsStackingPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsStickyMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsStormControlPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchPortsStpBpduGuard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsStpBpduGuardTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsStpRootGuard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsStpState(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsSwitchId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsTrunkMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsUntaggedVlans(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchPortsVirtualPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchPreProvisioned(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPtpProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchPtpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPurdueLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchQosDropPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchQosRedProbability(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRadiusNasIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRadiusNasIpOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRemoteLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "csv"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["csv"], _ = expandSwitchControllerManagedSwitchRemoteLogCsv(d, i["csv"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "facility"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["facility"], _ = expandSwitchControllerManagedSwitchRemoteLogFacility(d, i["facility"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSwitchControllerManagedSwitchRemoteLogName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandSwitchControllerManagedSwitchRemoteLogPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server"], _ = expandSwitchControllerManagedSwitchRemoteLogServer(d, i["server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["severity"], _ = expandSwitchControllerManagedSwitchRemoteLogSeverity(d, i["severity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandSwitchControllerManagedSwitchRemoteLogStatus(d, i["status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchRemoteLogCsv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRemoteLogFacility(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRemoteLogName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRemoteLogPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRemoteLogServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRemoteLogSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRemoteLogStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRouteOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRouteOffloadMclag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRouteOffloadRouter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "router_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["router-ip"], _ = expandSwitchControllerManagedSwitchRouteOffloadRouterRouterIp(d, i["router_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlan-name"], _ = expandSwitchControllerManagedSwitchRouteOffloadRouterVlanName(d, i["vlan_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchRouteOffloadRouterRouterIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRouteOffloadRouterVlanName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchSn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "events"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["events"], _ = expandSwitchControllerManagedSwitchSnmpCommunityEvents(d, i["events"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hosts"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandSwitchControllerManagedSwitchSnmpCommunityHosts(d, i["hosts"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["hosts"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSwitchControllerManagedSwitchSnmpCommunityId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSwitchControllerManagedSwitchSnmpCommunityName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_v1_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["query-v1-port"], _ = expandSwitchControllerManagedSwitchSnmpCommunityQueryV1Port(d, i["query_v1_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_v1_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["query-v1-status"], _ = expandSwitchControllerManagedSwitchSnmpCommunityQueryV1Status(d, i["query_v1_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_v2c_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["query-v2c-port"], _ = expandSwitchControllerManagedSwitchSnmpCommunityQueryV2CPort(d, i["query_v2c_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_v2c_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["query-v2c-status"], _ = expandSwitchControllerManagedSwitchSnmpCommunityQueryV2CStatus(d, i["query_v2c_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandSwitchControllerManagedSwitchSnmpCommunityStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v1_lport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["trap-v1-lport"], _ = expandSwitchControllerManagedSwitchSnmpCommunityTrapV1Lport(d, i["trap_v1_lport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v1_rport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["trap-v1-rport"], _ = expandSwitchControllerManagedSwitchSnmpCommunityTrapV1Rport(d, i["trap_v1_rport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v1_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["trap-v1-status"], _ = expandSwitchControllerManagedSwitchSnmpCommunityTrapV1Status(d, i["trap_v1_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v2c_lport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["trap-v2c-lport"], _ = expandSwitchControllerManagedSwitchSnmpCommunityTrapV2CLport(d, i["trap_v2c_lport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v2c_rport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["trap-v2c-rport"], _ = expandSwitchControllerManagedSwitchSnmpCommunityTrapV2CRport(d, i["trap_v2c_rport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v2c_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["trap-v2c-status"], _ = expandSwitchControllerManagedSwitchSnmpCommunityTrapV2CStatus(d, i["trap_v2c_status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityEvents(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityHosts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandSwitchControllerManagedSwitchSnmpCommunityHostsId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandSwitchControllerManagedSwitchSnmpCommunityHostsIp(d, i["ip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityHostsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityHostsIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityQueryV1Port(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityQueryV1Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityQueryV2CPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityQueryV2CStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityTrapV1Lport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityTrapV1Rport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityTrapV1Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityTrapV2CLport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityTrapV2CRport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityTrapV2CStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpSysinfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "contact_info"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["contact-info"], _ = expandSwitchControllerManagedSwitchSnmpSysinfoContactInfo(d, i["contact_info"], pre_append)
	}
	pre_append = pre + ".0." + "description"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["description"], _ = expandSwitchControllerManagedSwitchSnmpSysinfoDescription(d, i["description"], pre_append)
	}
	pre_append = pre + ".0." + "engine_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["engine-id"], _ = expandSwitchControllerManagedSwitchSnmpSysinfoEngineId(d, i["engine_id"], pre_append)
	}
	pre_append = pre + ".0." + "location"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["location"], _ = expandSwitchControllerManagedSwitchSnmpSysinfoLocation(d, i["location"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandSwitchControllerManagedSwitchSnmpSysinfoStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchSnmpSysinfoContactInfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpSysinfoDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpSysinfoEngineId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpSysinfoLocation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpSysinfoStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpTrapThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "trap_high_cpu_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["trap-high-cpu-threshold"], _ = expandSwitchControllerManagedSwitchSnmpTrapThresholdTrapHighCpuThreshold(d, i["trap_high_cpu_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "trap_log_full_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["trap-log-full-threshold"], _ = expandSwitchControllerManagedSwitchSnmpTrapThresholdTrapLogFullThreshold(d, i["trap_log_full_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "trap_low_memory_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["trap-low-memory-threshold"], _ = expandSwitchControllerManagedSwitchSnmpTrapThresholdTrapLowMemoryThreshold(d, i["trap_low_memory_threshold"], pre_append)
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchSnmpTrapThresholdTrapHighCpuThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpTrapThresholdTrapLogFullThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpTrapThresholdTrapLowMemoryThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_proto"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-proto"], _ = expandSwitchControllerManagedSwitchSnmpUserAuthProto(d, i["auth_proto"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_pwd"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-pwd"], _ = expandSwitchControllerManagedSwitchSnmpUserAuthPwd(d, i["auth_pwd"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSwitchControllerManagedSwitchSnmpUserName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priv_proto"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priv-proto"], _ = expandSwitchControllerManagedSwitchSnmpUserPrivProto(d, i["priv_proto"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priv_pwd"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priv-pwd"], _ = expandSwitchControllerManagedSwitchSnmpUserPrivPwd(d, i["priv_pwd"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queries"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["queries"], _ = expandSwitchControllerManagedSwitchSnmpUserQueries(d, i["queries"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["query-port"], _ = expandSwitchControllerManagedSwitchSnmpUserQueryPort(d, i["query_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_level"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["security-level"], _ = expandSwitchControllerManagedSwitchSnmpUserSecurityLevel(d, i["security_level"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchSnmpUserAuthProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpUserAuthPwd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchSnmpUserName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpUserPrivProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpUserPrivPwd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchSnmpUserQueries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpUserQueryPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpUserSecurityLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStagedImageVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStaticMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandSwitchControllerManagedSwitchStaticMacDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSwitchControllerManagedSwitchStaticMacId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandSwitchControllerManagedSwitchStaticMacInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac"], _ = expandSwitchControllerManagedSwitchStaticMacMac(d, i["mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandSwitchControllerManagedSwitchStaticMacType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlan"], _ = expandSwitchControllerManagedSwitchStaticMacVlan(d, i["vlan"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchStaticMacDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStaticMacId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStaticMacInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStaticMacMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStaticMacType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStaticMacVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchStormControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "broadcast"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["broadcast"], _ = expandSwitchControllerManagedSwitchStormControlBroadcast(d, i["broadcast"], pre_append)
	}
	pre_append = pre + ".0." + "local_override"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["local-override"], _ = expandSwitchControllerManagedSwitchStormControlLocalOverride(d, i["local_override"], pre_append)
	}
	pre_append = pre + ".0." + "rate"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["rate"], _ = expandSwitchControllerManagedSwitchStormControlRate(d, i["rate"], pre_append)
	}
	pre_append = pre + ".0." + "unknown_multicast"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unknown-multicast"], _ = expandSwitchControllerManagedSwitchStormControlUnknownMulticast(d, i["unknown_multicast"], pre_append)
	}
	pre_append = pre + ".0." + "unknown_unicast"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unknown-unicast"], _ = expandSwitchControllerManagedSwitchStormControlUnknownUnicast(d, i["unknown_unicast"], pre_append)
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchStormControlBroadcast(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStormControlLocalOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStormControlRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStormControlUnknownMulticast(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStormControlUnknownUnicast(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStpInstance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandSwitchControllerManagedSwitchStpInstanceId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandSwitchControllerManagedSwitchStpInstancePriority(d, i["priority"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchStpInstanceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStpInstancePriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStpSettings(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "forward_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["forward-time"], _ = expandSwitchControllerManagedSwitchStpSettingsForwardTime(d, i["forward_time"], pre_append)
	}
	pre_append = pre + ".0." + "hello_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["hello-time"], _ = expandSwitchControllerManagedSwitchStpSettingsHelloTime(d, i["hello_time"], pre_append)
	}
	pre_append = pre + ".0." + "local_override"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["local-override"], _ = expandSwitchControllerManagedSwitchStpSettingsLocalOverride(d, i["local_override"], pre_append)
	}
	pre_append = pre + ".0." + "max_age"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-age"], _ = expandSwitchControllerManagedSwitchStpSettingsMaxAge(d, i["max_age"], pre_append)
	}
	pre_append = pre + ".0." + "max_hops"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-hops"], _ = expandSwitchControllerManagedSwitchStpSettingsMaxHops(d, i["max_hops"], pre_append)
	}
	pre_append = pre + ".0." + "name"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["name"], _ = expandSwitchControllerManagedSwitchStpSettingsName(d, i["name"], pre_append)
	}
	pre_append = pre + ".0." + "pending_timer"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pending-timer"], _ = expandSwitchControllerManagedSwitchStpSettingsPendingTimer(d, i["pending_timer"], pre_append)
	}
	pre_append = pre + ".0." + "revision"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["revision"], _ = expandSwitchControllerManagedSwitchStpSettingsRevision(d, i["revision"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandSwitchControllerManagedSwitchStpSettingsStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchStpSettingsForwardTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStpSettingsHelloTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStpSettingsLocalOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStpSettingsMaxAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStpSettingsMaxHops(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStpSettingsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStpSettingsPendingTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStpSettingsRevision(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStpSettingsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSwitchDeviceTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSwitchDhcpOpt43Key(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSwitchId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSwitchLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "local_override"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["local-override"], _ = expandSwitchControllerManagedSwitchSwitchLogLocalOverride(d, i["local_override"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandSwitchControllerManagedSwitchSwitchLogSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandSwitchControllerManagedSwitchSwitchLogStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchSwitchLogLocalOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSwitchLogSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSwitchLogStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSwitchProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchTdrSupported(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchTunnelDiscovered(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "assignment_priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["assignment-priority"], _ = expandSwitchControllerManagedSwitchVlanAssignmentPriority(d, i["assignment_priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlan-name"], _ = expandSwitchControllerManagedSwitchVlanVlanName(d, i["vlan_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchVlanAssignmentPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchVlanVlanName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSwitchControllerManagedSwitch(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("n802_1x_settings"); ok || d.HasChange("n802_1x_settings") {
		t, err := expandSwitchControllerManagedSwitch8021XSettings(d, v, "n802_1x_settings")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["802-1X-settings"] = t
		}
	}

	if v, ok := d.GetOk("_platform"); ok || d.HasChange("_platform") {
		t, err := expandSwitchControllerManagedSwitchPlatform(d, v, "_platform")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_platform"] = t
		}
	}

	if v, ok := d.GetOk("access_profile"); ok || d.HasChange("access_profile") {
		t, err := expandSwitchControllerManagedSwitchAccessProfile(d, v, "access_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-profile"] = t
		}
	}

	if v, ok := d.GetOk("custom_command"); ok || d.HasChange("custom_command") {
		t, err := expandSwitchControllerManagedSwitchCustomCommand(d, v, "custom_command")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-command"] = t
		}
	}

	if v, ok := d.GetOk("delayed_restart_trigger"); ok || d.HasChange("delayed_restart_trigger") {
		t, err := expandSwitchControllerManagedSwitchDelayedRestartTrigger(d, v, "delayed_restart_trigger")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delayed-restart-trigger"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSwitchControllerManagedSwitchDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_server_access_list"); ok || d.HasChange("dhcp_server_access_list") {
		t, err := expandSwitchControllerManagedSwitchDhcpServerAccessList(d, v, "dhcp_server_access_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-server-access-list"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_snooping_static_client"); ok || d.HasChange("dhcp_snooping_static_client") {
		t, err := expandSwitchControllerManagedSwitchDhcpSnoopingStaticClient(d, v, "dhcp_snooping_static_client")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-snooping-static-client"] = t
		}
	}

	if v, ok := d.GetOk("directly_connected"); ok || d.HasChange("directly_connected") {
		t, err := expandSwitchControllerManagedSwitchDirectlyConnected(d, v, "directly_connected")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["directly-connected"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_capability"); ok || d.HasChange("dynamic_capability") {
		t, err := expandSwitchControllerManagedSwitchDynamicCapability(d, v, "dynamic_capability")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic-capability"] = t
		}
	}

	if v, ok := d.GetOk("dynamically_discovered"); ok || d.HasChange("dynamically_discovered") {
		t, err := expandSwitchControllerManagedSwitchDynamicallyDiscovered(d, v, "dynamically_discovered")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamically-discovered"] = t
		}
	}

	if v, ok := d.GetOk("firmware_provision"); ok || d.HasChange("firmware_provision") {
		t, err := expandSwitchControllerManagedSwitchFirmwareProvision(d, v, "firmware_provision")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firmware-provision"] = t
		}
	}

	if v, ok := d.GetOk("firmware_provision_latest"); ok || d.HasChange("firmware_provision_latest") {
		t, err := expandSwitchControllerManagedSwitchFirmwareProvisionLatest(d, v, "firmware_provision_latest")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firmware-provision-latest"] = t
		}
	}

	if v, ok := d.GetOk("firmware_provision_version"); ok || d.HasChange("firmware_provision_version") {
		t, err := expandSwitchControllerManagedSwitchFirmwareProvisionVersion(d, v, "firmware_provision_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firmware-provision-version"] = t
		}
	}

	if v, ok := d.GetOk("flow_identity"); ok || d.HasChange("flow_identity") {
		t, err := expandSwitchControllerManagedSwitchFlowIdentity(d, v, "flow_identity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flow-identity"] = t
		}
	}

	if v, ok := d.GetOk("fsw_wan1_admin"); ok || d.HasChange("fsw_wan1_admin") {
		t, err := expandSwitchControllerManagedSwitchFswWan1Admin(d, v, "fsw_wan1_admin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsw-wan1-admin"] = t
		}
	}

	if v, ok := d.GetOk("fsw_wan1_peer"); ok || d.HasChange("fsw_wan1_peer") {
		t, err := expandSwitchControllerManagedSwitchFswWan1Peer(d, v, "fsw_wan1_peer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsw-wan1-peer"] = t
		}
	}

	if v, ok := d.GetOk("fsw_wan2_admin"); ok || d.HasChange("fsw_wan2_admin") {
		t, err := expandSwitchControllerManagedSwitchFswWan2Admin(d, v, "fsw_wan2_admin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsw-wan2-admin"] = t
		}
	}

	if v, ok := d.GetOk("fsw_wan2_peer"); ok || d.HasChange("fsw_wan2_peer") {
		t, err := expandSwitchControllerManagedSwitchFswWan2Peer(d, v, "fsw_wan2_peer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsw-wan2-peer"] = t
		}
	}

	if v, ok := d.GetOk("igmp_snooping"); ok || d.HasChange("igmp_snooping") {
		t, err := expandSwitchControllerManagedSwitchIgmpSnooping(d, v, "igmp_snooping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["igmp-snooping"] = t
		}
	}

	if v, ok := d.GetOk("ip_source_guard"); ok || d.HasChange("ip_source_guard") {
		t, err := expandSwitchControllerManagedSwitchIpSourceGuard(d, v, "ip_source_guard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-source-guard"] = t
		}
	}

	if v, ok := d.GetOk("l3_discovered"); ok || d.HasChange("l3_discovered") {
		t, err := expandSwitchControllerManagedSwitchL3Discovered(d, v, "l3_discovered")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l3-discovered"] = t
		}
	}

	if v, ok := d.GetOk("max_allowed_trunk_members"); ok || d.HasChange("max_allowed_trunk_members") {
		t, err := expandSwitchControllerManagedSwitchMaxAllowedTrunkMembers(d, v, "max_allowed_trunk_members")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-allowed-trunk-members"] = t
		}
	}

	if v, ok := d.GetOk("mclag_igmp_snooping_aware"); ok || d.HasChange("mclag_igmp_snooping_aware") {
		t, err := expandSwitchControllerManagedSwitchMclagIgmpSnoopingAware(d, v, "mclag_igmp_snooping_aware")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mclag-igmp-snooping-aware"] = t
		}
	}

	if v, ok := d.GetOk("mgmt_mode"); ok || d.HasChange("mgmt_mode") {
		t, err := expandSwitchControllerManagedSwitchMgmtMode(d, v, "mgmt_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mgmt-mode"] = t
		}
	}

	if v, ok := d.GetOk("mirror"); ok || d.HasChange("mirror") {
		t, err := expandSwitchControllerManagedSwitchMirror(d, v, "mirror")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mirror"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerManagedSwitchName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("override_snmp_community"); ok || d.HasChange("override_snmp_community") {
		t, err := expandSwitchControllerManagedSwitchOverrideSnmpCommunity(d, v, "override_snmp_community")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-snmp-community"] = t
		}
	}

	if v, ok := d.GetOk("override_snmp_sysinfo"); ok || d.HasChange("override_snmp_sysinfo") {
		t, err := expandSwitchControllerManagedSwitchOverrideSnmpSysinfo(d, v, "override_snmp_sysinfo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-snmp-sysinfo"] = t
		}
	}

	if v, ok := d.GetOk("override_snmp_trap_threshold"); ok || d.HasChange("override_snmp_trap_threshold") {
		t, err := expandSwitchControllerManagedSwitchOverrideSnmpTrapThreshold(d, v, "override_snmp_trap_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-snmp-trap-threshold"] = t
		}
	}

	if v, ok := d.GetOk("override_snmp_user"); ok || d.HasChange("override_snmp_user") {
		t, err := expandSwitchControllerManagedSwitchOverrideSnmpUser(d, v, "override_snmp_user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-snmp-user"] = t
		}
	}

	if v, ok := d.GetOk("owner_vdom"); ok || d.HasChange("owner_vdom") {
		t, err := expandSwitchControllerManagedSwitchOwnerVdom(d, v, "owner_vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["owner-vdom"] = t
		}
	}

	if v, ok := d.GetOk("poe_detection_type"); ok || d.HasChange("poe_detection_type") {
		t, err := expandSwitchControllerManagedSwitchPoeDetectionType(d, v, "poe_detection_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poe-detection-type"] = t
		}
	}

	if v, ok := d.GetOk("poe_lldp_detection"); ok || d.HasChange("poe_lldp_detection") {
		t, err := expandSwitchControllerManagedSwitchPoeLldpDetection(d, v, "poe_lldp_detection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poe-lldp-detection"] = t
		}
	}

	if v, ok := d.GetOk("poe_pre_standard_detection"); ok || d.HasChange("poe_pre_standard_detection") {
		t, err := expandSwitchControllerManagedSwitchPoePreStandardDetection(d, v, "poe_pre_standard_detection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poe-pre-standard-detection"] = t
		}
	}

	if v, ok := d.GetOk("ports"); ok || d.HasChange("ports") {
		t, err := expandSwitchControllerManagedSwitchPorts(d, v, "ports")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ports"] = t
		}
	}

	if v, ok := d.GetOk("pre_provisioned"); ok || d.HasChange("pre_provisioned") {
		t, err := expandSwitchControllerManagedSwitchPreProvisioned(d, v, "pre_provisioned")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pre-provisioned"] = t
		}
	}

	if v, ok := d.GetOk("ptp_profile"); ok || d.HasChange("ptp_profile") {
		t, err := expandSwitchControllerManagedSwitchPtpProfile(d, v, "ptp_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ptp-profile"] = t
		}
	}

	if v, ok := d.GetOk("ptp_status"); ok || d.HasChange("ptp_status") {
		t, err := expandSwitchControllerManagedSwitchPtpStatus(d, v, "ptp_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ptp-status"] = t
		}
	}

	if v, ok := d.GetOk("purdue_level"); ok || d.HasChange("purdue_level") {
		t, err := expandSwitchControllerManagedSwitchPurdueLevel(d, v, "purdue_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["purdue-level"] = t
		}
	}

	if v, ok := d.GetOk("qos_drop_policy"); ok || d.HasChange("qos_drop_policy") {
		t, err := expandSwitchControllerManagedSwitchQosDropPolicy(d, v, "qos_drop_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qos-drop-policy"] = t
		}
	}

	if v, ok := d.GetOk("qos_red_probability"); ok || d.HasChange("qos_red_probability") {
		t, err := expandSwitchControllerManagedSwitchQosRedProbability(d, v, "qos_red_probability")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qos-red-probability"] = t
		}
	}

	if v, ok := d.GetOk("radius_nas_ip"); ok || d.HasChange("radius_nas_ip") {
		t, err := expandSwitchControllerManagedSwitchRadiusNasIp(d, v, "radius_nas_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-nas-ip"] = t
		}
	}

	if v, ok := d.GetOk("radius_nas_ip_override"); ok || d.HasChange("radius_nas_ip_override") {
		t, err := expandSwitchControllerManagedSwitchRadiusNasIpOverride(d, v, "radius_nas_ip_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-nas-ip-override"] = t
		}
	}

	if v, ok := d.GetOk("remote_log"); ok || d.HasChange("remote_log") {
		t, err := expandSwitchControllerManagedSwitchRemoteLog(d, v, "remote_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-log"] = t
		}
	}

	if v, ok := d.GetOk("route_offload"); ok || d.HasChange("route_offload") {
		t, err := expandSwitchControllerManagedSwitchRouteOffload(d, v, "route_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-offload"] = t
		}
	}

	if v, ok := d.GetOk("route_offload_mclag"); ok || d.HasChange("route_offload_mclag") {
		t, err := expandSwitchControllerManagedSwitchRouteOffloadMclag(d, v, "route_offload_mclag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-offload-mclag"] = t
		}
	}

	if v, ok := d.GetOk("route_offload_router"); ok || d.HasChange("route_offload_router") {
		t, err := expandSwitchControllerManagedSwitchRouteOffloadRouter(d, v, "route_offload_router")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-offload-router"] = t
		}
	}

	if v, ok := d.GetOk("sn"); ok || d.HasChange("sn") {
		t, err := expandSwitchControllerManagedSwitchSn(d, v, "sn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sn"] = t
		}
	}

	if v, ok := d.GetOk("snmp_community"); ok || d.HasChange("snmp_community") {
		t, err := expandSwitchControllerManagedSwitchSnmpCommunity(d, v, "snmp_community")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["snmp-community"] = t
		}
	}

	if v, ok := d.GetOk("snmp_sysinfo"); ok || d.HasChange("snmp_sysinfo") {
		t, err := expandSwitchControllerManagedSwitchSnmpSysinfo(d, v, "snmp_sysinfo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["snmp-sysinfo"] = t
		}
	}

	if v, ok := d.GetOk("snmp_trap_threshold"); ok || d.HasChange("snmp_trap_threshold") {
		t, err := expandSwitchControllerManagedSwitchSnmpTrapThreshold(d, v, "snmp_trap_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["snmp-trap-threshold"] = t
		}
	}

	if v, ok := d.GetOk("snmp_user"); ok || d.HasChange("snmp_user") {
		t, err := expandSwitchControllerManagedSwitchSnmpUser(d, v, "snmp_user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["snmp-user"] = t
		}
	}

	if v, ok := d.GetOk("staged_image_version"); ok || d.HasChange("staged_image_version") {
		t, err := expandSwitchControllerManagedSwitchStagedImageVersion(d, v, "staged_image_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["staged-image-version"] = t
		}
	}

	if v, ok := d.GetOk("static_mac"); ok || d.HasChange("static_mac") {
		t, err := expandSwitchControllerManagedSwitchStaticMac(d, v, "static_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["static-mac"] = t
		}
	}

	if v, ok := d.GetOk("storm_control"); ok || d.HasChange("storm_control") {
		t, err := expandSwitchControllerManagedSwitchStormControl(d, v, "storm_control")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["storm-control"] = t
		}
	}

	if v, ok := d.GetOk("stp_instance"); ok || d.HasChange("stp_instance") {
		t, err := expandSwitchControllerManagedSwitchStpInstance(d, v, "stp_instance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stp-instance"] = t
		}
	}

	if v, ok := d.GetOk("stp_settings"); ok || d.HasChange("stp_settings") {
		t, err := expandSwitchControllerManagedSwitchStpSettings(d, v, "stp_settings")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stp-settings"] = t
		}
	}

	if v, ok := d.GetOk("switch_device_tag"); ok || d.HasChange("switch_device_tag") {
		t, err := expandSwitchControllerManagedSwitchSwitchDeviceTag(d, v, "switch_device_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-device-tag"] = t
		}
	}

	if v, ok := d.GetOk("switch_dhcp_opt43_key"); ok || d.HasChange("switch_dhcp_opt43_key") {
		t, err := expandSwitchControllerManagedSwitchSwitchDhcpOpt43Key(d, v, "switch_dhcp_opt43_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-dhcp_opt43_key"] = t
		}
	}

	if v, ok := d.GetOk("switch_id"); ok || d.HasChange("switch_id") {
		t, err := expandSwitchControllerManagedSwitchSwitchId(d, v, "switch_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-id"] = t
		}
	}

	if v, ok := d.GetOk("switch_log"); ok || d.HasChange("switch_log") {
		t, err := expandSwitchControllerManagedSwitchSwitchLog(d, v, "switch_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-log"] = t
		}
	}

	if v, ok := d.GetOk("switch_profile"); ok || d.HasChange("switch_profile") {
		t, err := expandSwitchControllerManagedSwitchSwitchProfile(d, v, "switch_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-profile"] = t
		}
	}

	if v, ok := d.GetOk("tdr_supported"); ok || d.HasChange("tdr_supported") {
		t, err := expandSwitchControllerManagedSwitchTdrSupported(d, v, "tdr_supported")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tdr-supported"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_discovered"); ok || d.HasChange("tunnel_discovered") {
		t, err := expandSwitchControllerManagedSwitchTunnelDiscovered(d, v, "tunnel_discovered")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-discovered"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandSwitchControllerManagedSwitchType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("version"); ok || d.HasChange("version") {
		t, err := expandSwitchControllerManagedSwitchVersion(d, v, "version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["version"] = t
		}
	}

	if v, ok := d.GetOk("vlan"); ok || d.HasChange("vlan") {
		t, err := expandSwitchControllerManagedSwitchVlan(d, v, "vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan"] = t
		}
	}

	return &obj, nil
}
