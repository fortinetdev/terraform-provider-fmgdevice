// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure BGP.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterBgp() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterBgpUpdate,
		Read:   resourceRouterBgpRead,
		Update: resourceRouterBgpUpdate,
		Delete: resourceRouterBgpDelete,

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
			"additional_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"additional_path_select": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"additional_path_select_vpnv4": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"additional_path_select_vpnv6": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"additional_path_select6": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"additional_path_vpnv4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"additional_path_vpnv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"additional_path6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_distance": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"distance": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"neighbour_prefix": &schema.Schema{
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"route_list": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"aggregate_address": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"as_set": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"summary_only": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"aggregate_address6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"as_set": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"prefix6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"summary_only": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"always_compare_med": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"as": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"bestpath_as_path_ignore": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bestpath_cmp_confed_aspath": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bestpath_cmp_routerid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bestpath_med_confed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bestpath_med_missing_as_worst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_to_client_reflection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cluster_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"confederation_identifier": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"confederation_peers": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"cross_family_conditional_adv": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dampening": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dampening_max_suppress_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dampening_reachability_half_life": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dampening_reuse": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dampening_route_map": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dampening_suppress": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dampening_unreachability_half_life": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"default_local_preference": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"deterministic_med": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"distance_external": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"distance_internal": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"distance_local": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ebgp_multipath": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enforce_first_as": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fast_external_failover": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"graceful_end_on_timer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"graceful_restart": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"graceful_restart_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"graceful_stalepath_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"graceful_update_delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"holdtime_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ibgp_multipath": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ignore_optional_capability": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keepalive_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"log_neighbour_changes": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multipath_recursive_distance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"neighbor": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"activate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"activate_evpn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"activate_vpnv4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"activate_vpnv6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"activate6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"additional_path": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"additional_path_vpnv4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"additional_path_vpnv6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"additional_path6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"adv_additional_path": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"adv_additional_path_vpnv4": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"adv_additional_path_vpnv6": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"adv_additional_path6": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"advertisement_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"allowas_in": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"allowas_in_enable": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"allowas_in_enable_evpn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"allowas_in_enable_vpnv4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"allowas_in_enable_vpnv6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"allowas_in_enable6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"allowas_in_evpn": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"allowas_in_vpnv4": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"allowas_in_vpnv6": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"allowas_in6": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"as_override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"as_override6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"attribute_unchanged": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"attribute_unchanged_vpnv4": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"attribute_unchanged_vpnv6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"attribute_unchanged6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"auth_options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"bfd": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"capability_default_originate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"capability_default_originate6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"capability_dynamic": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"capability_graceful_restart": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"capability_graceful_restart_evpn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"capability_graceful_restart_vpnv4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"capability_graceful_restart_vpnv6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"capability_graceful_restart6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"capability_orf": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"capability_orf6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"capability_route_refresh": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"conditional_advertise": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"advertise_routemap": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"condition_routemap": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"condition_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"conditional_advertise6": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"advertise_routemap": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"condition_routemap": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"condition_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"connect_timer": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"default_originate_routemap": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"default_originate_routemap6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"distribute_list_in": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"distribute_list_in_vpnv4": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"distribute_list_in_vpnv6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"distribute_list_in6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"distribute_list_out": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"distribute_list_out_vpnv4": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"distribute_list_out_vpnv6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"distribute_list_out6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"dont_capability_negotiate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ebgp_enforce_multihop": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ebgp_multihop_ttl": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"filter_list_in": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"filter_list_in_vpnv4": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"filter_list_in_vpnv6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"filter_list_in6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"filter_list_out": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"filter_list_out_vpnv4": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"filter_list_out_vpnv6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"filter_list_out6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"holdtime_timer": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"keep_alive_timer": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"link_down_failover": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"local_as": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"local_as_no_prepend": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"local_as_replace_as": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"maximum_prefix": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"maximum_prefix_evpn": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"maximum_prefix_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"maximum_prefix_threshold_evpn": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"maximum_prefix_threshold_vpnv4": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"maximum_prefix_threshold_vpnv6": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"maximum_prefix_threshold6": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"maximum_prefix_vpnv4": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"maximum_prefix_vpnv6": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"maximum_prefix_warning_only": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"maximum_prefix_warning_only_evpn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"maximum_prefix_warning_only_vpnv4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"maximum_prefix_warning_only_vpnv6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"maximum_prefix_warning_only6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"maximum_prefix6": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"next_hop_self": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"next_hop_self_rr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"next_hop_self_rr6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"next_hop_self_vpnv4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"next_hop_self_vpnv6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"next_hop_self6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"override_capability": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"passive": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"password": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"prefix_list_in": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"prefix_list_in_vpnv4": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"prefix_list_in_vpnv6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"prefix_list_in6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"prefix_list_out": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"prefix_list_out_vpnv4": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"prefix_list_out_vpnv6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"prefix_list_out6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"remote_as": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"remove_private_as": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"remove_private_as_evpn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"remove_private_as_vpnv4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"remove_private_as_vpnv6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"remove_private_as6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"restart_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"retain_stale_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"route_map_in": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"route_map_in_evpn": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"route_map_in_vpnv4": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"route_map_in_vpnv6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"route_map_in6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"route_map_out": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"route_map_out_evpn": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"route_map_out_preferable": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"route_map_out_vpnv4": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"route_map_out_vpnv4_preferable": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"route_map_out_vpnv6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"route_map_out_vpnv6_preferable": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"route_map_out6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"route_map_out6_preferable": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"route_reflector_client": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"route_reflector_client_evpn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"route_reflector_client_vpnv4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"route_reflector_client_vpnv6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"route_reflector_client6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"route_server_client": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"route_server_client_evpn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"route_server_client_vpnv4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"route_server_client_vpnv6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"route_server_client6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"send_community": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"send_community_evpn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"send_community_vpnv4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"send_community_vpnv6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"send_community6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"shutdown": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"soft_reconfiguration": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"soft_reconfiguration_evpn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"soft_reconfiguration_vpnv4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"soft_reconfiguration_vpnv6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"soft_reconfiguration6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"stale_route": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"strict_capability_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsuppress_map": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"unsuppress_map6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"update_source": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"neighbor_group": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"activate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"activate_evpn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"activate_vpnv4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"activate_vpnv6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"activate6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"additional_path": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"additional_path_vpnv4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"additional_path_vpnv6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"additional_path6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"adv_additional_path": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"adv_additional_path_vpnv4": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"adv_additional_path_vpnv6": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"adv_additional_path6": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"advertisement_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"allowas_in": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"allowas_in_enable": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"allowas_in_enable_evpn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"allowas_in_enable_vpnv4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"allowas_in_enable_vpnv6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"allowas_in_enable6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"allowas_in_evpn": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"allowas_in_vpnv4": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"allowas_in_vpnv6": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"allowas_in6": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"as_override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"as_override6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"attribute_unchanged": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"attribute_unchanged_vpnv4": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"attribute_unchanged_vpnv6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"attribute_unchanged6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"auth_options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"bfd": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"capability_default_originate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"capability_default_originate6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"capability_dynamic": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"capability_graceful_restart": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"capability_graceful_restart_evpn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"capability_graceful_restart_vpnv4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"capability_graceful_restart_vpnv6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"capability_graceful_restart6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"capability_orf": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"capability_orf6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"capability_route_refresh": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"connect_timer": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"default_originate_routemap": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"default_originate_routemap6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"distribute_list_in": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"distribute_list_in_vpnv4": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"distribute_list_in_vpnv6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"distribute_list_in6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"distribute_list_out": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"distribute_list_out_vpnv4": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"distribute_list_out_vpnv6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"distribute_list_out6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"dont_capability_negotiate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ebgp_enforce_multihop": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ebgp_multihop_ttl": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"filter_list_in": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"filter_list_in_vpnv4": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"filter_list_in_vpnv6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"filter_list_in6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"filter_list_out": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"filter_list_out_vpnv4": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"filter_list_out_vpnv6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"filter_list_out6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"holdtime_timer": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"keep_alive_timer": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"link_down_failover": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"local_as": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"local_as_no_prepend": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"local_as_replace_as": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"maximum_prefix": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"maximum_prefix_evpn": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"maximum_prefix_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"maximum_prefix_threshold_evpn": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"maximum_prefix_threshold_vpnv4": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"maximum_prefix_threshold_vpnv6": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"maximum_prefix_threshold6": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"maximum_prefix_vpnv4": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"maximum_prefix_vpnv6": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"maximum_prefix_warning_only": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"maximum_prefix_warning_only_evpn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"maximum_prefix_warning_only_vpnv4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"maximum_prefix_warning_only_vpnv6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"maximum_prefix_warning_only6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"maximum_prefix6": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"next_hop_self": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"next_hop_self_rr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"next_hop_self_rr6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"next_hop_self_vpnv4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"next_hop_self_vpnv6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"next_hop_self6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"override_capability": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"passive": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"password": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"prefix_list_in": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"prefix_list_in_vpnv4": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"prefix_list_in_vpnv6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"prefix_list_in6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"prefix_list_out": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"prefix_list_out_vpnv4": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"prefix_list_out_vpnv6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"prefix_list_out6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"remote_as": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"remote_as_filter": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"remove_private_as": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"remove_private_as_evpn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"remove_private_as_vpnv4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"remove_private_as_vpnv6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"remove_private_as6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"restart_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"retain_stale_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"route_map_in": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"route_map_in_evpn": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"route_map_in_vpnv4": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"route_map_in_vpnv6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"route_map_in6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"route_map_out": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"route_map_out_evpn": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"route_map_out_preferable": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"route_map_out_vpnv4": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"route_map_out_vpnv4_preferable": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"route_map_out_vpnv6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"route_map_out_vpnv6_preferable": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"route_map_out6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"route_map_out6_preferable": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"route_reflector_client": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"route_reflector_client_evpn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"route_reflector_client_vpnv4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"route_reflector_client_vpnv6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"route_reflector_client6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"route_server_client": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"route_server_client_evpn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"route_server_client_vpnv4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"route_server_client_vpnv6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"route_server_client6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"send_community": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"send_community_evpn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"send_community_vpnv4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"send_community_vpnv6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"send_community6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"shutdown": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"soft_reconfiguration": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"soft_reconfiguration_evpn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"soft_reconfiguration_vpnv4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"soft_reconfiguration_vpnv6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"soft_reconfiguration6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"stale_route": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"strict_capability_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsuppress_map": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"unsuppress_map6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"update_source": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"neighbor_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"max_neighbor_num": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"neighbor_group": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"neighbor_range6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"max_neighbor_num": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"neighbor_group": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"prefix6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"network": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backdoor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"network_import_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"route_map": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"network_import_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backdoor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"network_import_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"prefix6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"route_map": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"recursive_inherit_priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"recursive_next_hop": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redistribute": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"route_map": &schema.Schema{
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
					},
				},
			},
			"redistribute6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"route_map": &schema.Schema{
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
					},
				},
			},
			"router_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"scan_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"synchronization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tag_resolve_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vrf_leak": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"target": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"route_map": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"vrf": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"vrf": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"vrf_leak6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"target": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"route_map": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"vrf": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"vrf": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"vrf": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"export_rt": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"import_route_map": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"import_rt": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"leak_target": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"route_map": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"vrf": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"rd": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"role": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vrf": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"vrf6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"export_rt": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"import_route_map": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"import_rt": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"leak_target": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"route_map": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"vrf": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"rd": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"role": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vrf": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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

func resourceRouterBgpUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterBgp(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgp resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterBgp(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("RouterBgp")

	return resourceRouterBgpRead(d, m)
}

func resourceRouterBgpDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterBgp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting RouterBgp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterBgpRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterBgp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterBgp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgp resource from API: %v", err)
	}
	return nil
}

func flattenRouterBgpAdditionalPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpAdditionalPathSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpAdditionalPathSelectVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpAdditionalPathSelectVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpAdditionalPathSelect6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpAdditionalPathVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpAdditionalPathVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpAdditionalPath6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpAdminDistance(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distance"
		if _, ok := i["distance"]; ok {
			v := flattenRouterBgpAdminDistanceDistance(i["distance"], d, pre_append)
			tmp["distance"] = fortiAPISubPartPatch(v, "RouterBgp-AdminDistance-Distance")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterBgpAdminDistanceId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterBgp-AdminDistance-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbour_prefix"
		if _, ok := i["neighbour-prefix"]; ok {
			v := flattenRouterBgpAdminDistanceNeighbourPrefix(i["neighbour-prefix"], d, pre_append)
			tmp["neighbour_prefix"] = fortiAPISubPartPatch(v, "RouterBgp-AdminDistance-NeighbourPrefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_list"
		if _, ok := i["route-list"]; ok {
			v := flattenRouterBgpAdminDistanceRouteList(i["route-list"], d, pre_append)
			tmp["route_list"] = fortiAPISubPartPatch(v, "RouterBgp-AdminDistance-RouteList")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterBgpAdminDistanceDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpAdminDistanceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpAdminDistanceNeighbourPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpAdminDistanceRouteList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpAggregateAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_set"
		if _, ok := i["as-set"]; ok {
			v := flattenRouterBgpAggregateAddressAsSet(i["as-set"], d, pre_append)
			tmp["as_set"] = fortiAPISubPartPatch(v, "RouterBgp-AggregateAddress-AsSet")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterBgpAggregateAddressId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterBgp-AggregateAddress-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			v := flattenRouterBgpAggregateAddressPrefix(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "RouterBgp-AggregateAddress-Prefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "summary_only"
		if _, ok := i["summary-only"]; ok {
			v := flattenRouterBgpAggregateAddressSummaryOnly(i["summary-only"], d, pre_append)
			tmp["summary_only"] = fortiAPISubPartPatch(v, "RouterBgp-AggregateAddress-SummaryOnly")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterBgpAggregateAddressAsSet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpAggregateAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpAggregateAddressPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpAggregateAddressSummaryOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpAggregateAddress6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_set"
		if _, ok := i["as-set"]; ok {
			v := flattenRouterBgpAggregateAddress6AsSet(i["as-set"], d, pre_append)
			tmp["as_set"] = fortiAPISubPartPatch(v, "RouterBgp-AggregateAddress6-AsSet")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterBgpAggregateAddress6Id(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterBgp-AggregateAddress6-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {
			v := flattenRouterBgpAggregateAddress6Prefix6(i["prefix6"], d, pre_append)
			tmp["prefix6"] = fortiAPISubPartPatch(v, "RouterBgp-AggregateAddress6-Prefix6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "summary_only"
		if _, ok := i["summary-only"]; ok {
			v := flattenRouterBgpAggregateAddress6SummaryOnly(i["summary-only"], d, pre_append)
			tmp["summary_only"] = fortiAPISubPartPatch(v, "RouterBgp-AggregateAddress6-SummaryOnly")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterBgpAggregateAddress6AsSet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpAggregateAddress6Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpAggregateAddress6Prefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpAggregateAddress6SummaryOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpAlwaysCompareMed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpBestpathAsPathIgnore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpBestpathCmpConfedAspath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpBestpathCmpRouterid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpBestpathMedConfed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpBestpathMedMissingAsWorst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpClientToClientReflection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpClusterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpConfederationIdentifier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpConfederationPeers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpCrossFamilyConditionalAdv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpDampening(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpDampeningMaxSuppressTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpDampeningReachabilityHalfLife(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpDampeningReuse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpDampeningRouteMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpDampeningSuppress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpDampeningUnreachabilityHalfLife(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpDefaultLocalPreference(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpDeterministicMed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpDistanceExternal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpDistanceInternal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpDistanceLocal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpEbgpMultipath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpEnforceFirstAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpFastExternalFailover(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpGracefulEndOnTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpGracefulRestart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpGracefulRestartTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpGracefulStalepathTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpGracefulUpdateDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpHoldtimeTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpIbgpMultipath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpIgnoreOptionalCapability(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpKeepaliveTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpLogNeighbourChanges(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpMultipathRecursiveDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighbor(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate"
		if _, ok := i["activate"]; ok {
			v := flattenRouterBgpNeighborActivate(i["activate"], d, pre_append)
			tmp["activate"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-Activate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate_evpn"
		if _, ok := i["activate-evpn"]; ok {
			v := flattenRouterBgpNeighborActivateEvpn(i["activate-evpn"], d, pre_append)
			tmp["activate_evpn"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-ActivateEvpn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate_vpnv4"
		if _, ok := i["activate-vpnv4"]; ok {
			v := flattenRouterBgpNeighborActivateVpnv4(i["activate-vpnv4"], d, pre_append)
			tmp["activate_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-ActivateVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate_vpnv6"
		if _, ok := i["activate-vpnv6"]; ok {
			v := flattenRouterBgpNeighborActivateVpnv6(i["activate-vpnv6"], d, pre_append)
			tmp["activate_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-ActivateVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate6"
		if _, ok := i["activate6"]; ok {
			v := flattenRouterBgpNeighborActivate6(i["activate6"], d, pre_append)
			tmp["activate6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-Activate6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path"
		if _, ok := i["additional-path"]; ok {
			v := flattenRouterBgpNeighborAdditionalPath(i["additional-path"], d, pre_append)
			tmp["additional_path"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-AdditionalPath")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path_vpnv4"
		if _, ok := i["additional-path-vpnv4"]; ok {
			v := flattenRouterBgpNeighborAdditionalPathVpnv4(i["additional-path-vpnv4"], d, pre_append)
			tmp["additional_path_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-AdditionalPathVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path_vpnv6"
		if _, ok := i["additional-path-vpnv6"]; ok {
			v := flattenRouterBgpNeighborAdditionalPathVpnv6(i["additional-path-vpnv6"], d, pre_append)
			tmp["additional_path_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-AdditionalPathVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path6"
		if _, ok := i["additional-path6"]; ok {
			v := flattenRouterBgpNeighborAdditionalPath6(i["additional-path6"], d, pre_append)
			tmp["additional_path6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-AdditionalPath6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path"
		if _, ok := i["adv-additional-path"]; ok {
			v := flattenRouterBgpNeighborAdvAdditionalPath(i["adv-additional-path"], d, pre_append)
			tmp["adv_additional_path"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-AdvAdditionalPath")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path_vpnv4"
		if _, ok := i["adv-additional-path-vpnv4"]; ok {
			v := flattenRouterBgpNeighborAdvAdditionalPathVpnv4(i["adv-additional-path-vpnv4"], d, pre_append)
			tmp["adv_additional_path_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-AdvAdditionalPathVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path_vpnv6"
		if _, ok := i["adv-additional-path-vpnv6"]; ok {
			v := flattenRouterBgpNeighborAdvAdditionalPathVpnv6(i["adv-additional-path-vpnv6"], d, pre_append)
			tmp["adv_additional_path_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-AdvAdditionalPathVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path6"
		if _, ok := i["adv-additional-path6"]; ok {
			v := flattenRouterBgpNeighborAdvAdditionalPath6(i["adv-additional-path6"], d, pre_append)
			tmp["adv_additional_path6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-AdvAdditionalPath6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertisement_interval"
		if _, ok := i["advertisement-interval"]; ok {
			v := flattenRouterBgpNeighborAdvertisementInterval(i["advertisement-interval"], d, pre_append)
			tmp["advertisement_interval"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-AdvertisementInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in"
		if _, ok := i["allowas-in"]; ok {
			v := flattenRouterBgpNeighborAllowasIn(i["allowas-in"], d, pre_append)
			tmp["allowas_in"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-AllowasIn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable"
		if _, ok := i["allowas-in-enable"]; ok {
			v := flattenRouterBgpNeighborAllowasInEnable(i["allowas-in-enable"], d, pre_append)
			tmp["allowas_in_enable"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-AllowasInEnable")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable_evpn"
		if _, ok := i["allowas-in-enable-evpn"]; ok {
			v := flattenRouterBgpNeighborAllowasInEnableEvpn(i["allowas-in-enable-evpn"], d, pre_append)
			tmp["allowas_in_enable_evpn"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-AllowasInEnableEvpn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable_vpnv4"
		if _, ok := i["allowas-in-enable-vpnv4"]; ok {
			v := flattenRouterBgpNeighborAllowasInEnableVpnv4(i["allowas-in-enable-vpnv4"], d, pre_append)
			tmp["allowas_in_enable_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-AllowasInEnableVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable_vpnv6"
		if _, ok := i["allowas-in-enable-vpnv6"]; ok {
			v := flattenRouterBgpNeighborAllowasInEnableVpnv6(i["allowas-in-enable-vpnv6"], d, pre_append)
			tmp["allowas_in_enable_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-AllowasInEnableVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable6"
		if _, ok := i["allowas-in-enable6"]; ok {
			v := flattenRouterBgpNeighborAllowasInEnable6(i["allowas-in-enable6"], d, pre_append)
			tmp["allowas_in_enable6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-AllowasInEnable6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_evpn"
		if _, ok := i["allowas-in-evpn"]; ok {
			v := flattenRouterBgpNeighborAllowasInEvpn(i["allowas-in-evpn"], d, pre_append)
			tmp["allowas_in_evpn"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-AllowasInEvpn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_vpnv4"
		if _, ok := i["allowas-in-vpnv4"]; ok {
			v := flattenRouterBgpNeighborAllowasInVpnv4(i["allowas-in-vpnv4"], d, pre_append)
			tmp["allowas_in_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-AllowasInVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_vpnv6"
		if _, ok := i["allowas-in-vpnv6"]; ok {
			v := flattenRouterBgpNeighborAllowasInVpnv6(i["allowas-in-vpnv6"], d, pre_append)
			tmp["allowas_in_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-AllowasInVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in6"
		if _, ok := i["allowas-in6"]; ok {
			v := flattenRouterBgpNeighborAllowasIn6(i["allowas-in6"], d, pre_append)
			tmp["allowas_in6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-AllowasIn6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_override"
		if _, ok := i["as-override"]; ok {
			v := flattenRouterBgpNeighborAsOverride(i["as-override"], d, pre_append)
			tmp["as_override"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-AsOverride")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_override6"
		if _, ok := i["as-override6"]; ok {
			v := flattenRouterBgpNeighborAsOverride6(i["as-override6"], d, pre_append)
			tmp["as_override6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-AsOverride6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged"
		if _, ok := i["attribute-unchanged"]; ok {
			v := flattenRouterBgpNeighborAttributeUnchanged(i["attribute-unchanged"], d, pre_append)
			tmp["attribute_unchanged"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-AttributeUnchanged")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged_vpnv4"
		if _, ok := i["attribute-unchanged-vpnv4"]; ok {
			v := flattenRouterBgpNeighborAttributeUnchangedVpnv4(i["attribute-unchanged-vpnv4"], d, pre_append)
			tmp["attribute_unchanged_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-AttributeUnchangedVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged_vpnv6"
		if _, ok := i["attribute-unchanged-vpnv6"]; ok {
			v := flattenRouterBgpNeighborAttributeUnchangedVpnv6(i["attribute-unchanged-vpnv6"], d, pre_append)
			tmp["attribute_unchanged_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-AttributeUnchangedVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged6"
		if _, ok := i["attribute-unchanged6"]; ok {
			v := flattenRouterBgpNeighborAttributeUnchanged6(i["attribute-unchanged6"], d, pre_append)
			tmp["attribute_unchanged6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-AttributeUnchanged6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_options"
		if _, ok := i["auth-options"]; ok {
			v := flattenRouterBgpNeighborAuthOptions(i["auth-options"], d, pre_append)
			tmp["auth_options"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-AuthOptions")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := i["bfd"]; ok {
			v := flattenRouterBgpNeighborBfd(i["bfd"], d, pre_append)
			tmp["bfd"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-Bfd")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_default_originate"
		if _, ok := i["capability-default-originate"]; ok {
			v := flattenRouterBgpNeighborCapabilityDefaultOriginate(i["capability-default-originate"], d, pre_append)
			tmp["capability_default_originate"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-CapabilityDefaultOriginate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_default_originate6"
		if _, ok := i["capability-default-originate6"]; ok {
			v := flattenRouterBgpNeighborCapabilityDefaultOriginate6(i["capability-default-originate6"], d, pre_append)
			tmp["capability_default_originate6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-CapabilityDefaultOriginate6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_dynamic"
		if _, ok := i["capability-dynamic"]; ok {
			v := flattenRouterBgpNeighborCapabilityDynamic(i["capability-dynamic"], d, pre_append)
			tmp["capability_dynamic"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-CapabilityDynamic")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart"
		if _, ok := i["capability-graceful-restart"]; ok {
			v := flattenRouterBgpNeighborCapabilityGracefulRestart(i["capability-graceful-restart"], d, pre_append)
			tmp["capability_graceful_restart"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-CapabilityGracefulRestart")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart_evpn"
		if _, ok := i["capability-graceful-restart-evpn"]; ok {
			v := flattenRouterBgpNeighborCapabilityGracefulRestartEvpn(i["capability-graceful-restart-evpn"], d, pre_append)
			tmp["capability_graceful_restart_evpn"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-CapabilityGracefulRestartEvpn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart_vpnv4"
		if _, ok := i["capability-graceful-restart-vpnv4"]; ok {
			v := flattenRouterBgpNeighborCapabilityGracefulRestartVpnv4(i["capability-graceful-restart-vpnv4"], d, pre_append)
			tmp["capability_graceful_restart_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-CapabilityGracefulRestartVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart_vpnv6"
		if _, ok := i["capability-graceful-restart-vpnv6"]; ok {
			v := flattenRouterBgpNeighborCapabilityGracefulRestartVpnv6(i["capability-graceful-restart-vpnv6"], d, pre_append)
			tmp["capability_graceful_restart_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-CapabilityGracefulRestartVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart6"
		if _, ok := i["capability-graceful-restart6"]; ok {
			v := flattenRouterBgpNeighborCapabilityGracefulRestart6(i["capability-graceful-restart6"], d, pre_append)
			tmp["capability_graceful_restart6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-CapabilityGracefulRestart6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_orf"
		if _, ok := i["capability-orf"]; ok {
			v := flattenRouterBgpNeighborCapabilityOrf(i["capability-orf"], d, pre_append)
			tmp["capability_orf"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-CapabilityOrf")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_orf6"
		if _, ok := i["capability-orf6"]; ok {
			v := flattenRouterBgpNeighborCapabilityOrf6(i["capability-orf6"], d, pre_append)
			tmp["capability_orf6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-CapabilityOrf6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_route_refresh"
		if _, ok := i["capability-route-refresh"]; ok {
			v := flattenRouterBgpNeighborCapabilityRouteRefresh(i["capability-route-refresh"], d, pre_append)
			tmp["capability_route_refresh"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-CapabilityRouteRefresh")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "conditional_advertise"
		if _, ok := i["conditional-advertise"]; ok {
			v := flattenRouterBgpNeighborConditionalAdvertise(i["conditional-advertise"], d, pre_append)
			tmp["conditional_advertise"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-ConditionalAdvertise")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "conditional_advertise6"
		if _, ok := i["conditional-advertise6"]; ok {
			v := flattenRouterBgpNeighborConditionalAdvertise6(i["conditional-advertise6"], d, pre_append)
			tmp["conditional_advertise6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-ConditionalAdvertise6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "connect_timer"
		if _, ok := i["connect-timer"]; ok {
			v := flattenRouterBgpNeighborConnectTimer(i["connect-timer"], d, pre_append)
			tmp["connect_timer"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-ConnectTimer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_originate_routemap"
		if _, ok := i["default-originate-routemap"]; ok {
			v := flattenRouterBgpNeighborDefaultOriginateRoutemap(i["default-originate-routemap"], d, pre_append)
			tmp["default_originate_routemap"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-DefaultOriginateRoutemap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_originate_routemap6"
		if _, ok := i["default-originate-routemap6"]; ok {
			v := flattenRouterBgpNeighborDefaultOriginateRoutemap6(i["default-originate-routemap6"], d, pre_append)
			tmp["default_originate_routemap6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-DefaultOriginateRoutemap6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenRouterBgpNeighborDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in"
		if _, ok := i["distribute-list-in"]; ok {
			v := flattenRouterBgpNeighborDistributeListIn(i["distribute-list-in"], d, pre_append)
			tmp["distribute_list_in"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-DistributeListIn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in_vpnv4"
		if _, ok := i["distribute-list-in-vpnv4"]; ok {
			v := flattenRouterBgpNeighborDistributeListInVpnv4(i["distribute-list-in-vpnv4"], d, pre_append)
			tmp["distribute_list_in_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-DistributeListInVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in_vpnv6"
		if _, ok := i["distribute-list-in-vpnv6"]; ok {
			v := flattenRouterBgpNeighborDistributeListInVpnv6(i["distribute-list-in-vpnv6"], d, pre_append)
			tmp["distribute_list_in_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-DistributeListInVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in6"
		if _, ok := i["distribute-list-in6"]; ok {
			v := flattenRouterBgpNeighborDistributeListIn6(i["distribute-list-in6"], d, pre_append)
			tmp["distribute_list_in6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-DistributeListIn6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out"
		if _, ok := i["distribute-list-out"]; ok {
			v := flattenRouterBgpNeighborDistributeListOut(i["distribute-list-out"], d, pre_append)
			tmp["distribute_list_out"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-DistributeListOut")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out_vpnv4"
		if _, ok := i["distribute-list-out-vpnv4"]; ok {
			v := flattenRouterBgpNeighborDistributeListOutVpnv4(i["distribute-list-out-vpnv4"], d, pre_append)
			tmp["distribute_list_out_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-DistributeListOutVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out_vpnv6"
		if _, ok := i["distribute-list-out-vpnv6"]; ok {
			v := flattenRouterBgpNeighborDistributeListOutVpnv6(i["distribute-list-out-vpnv6"], d, pre_append)
			tmp["distribute_list_out_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-DistributeListOutVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out6"
		if _, ok := i["distribute-list-out6"]; ok {
			v := flattenRouterBgpNeighborDistributeListOut6(i["distribute-list-out6"], d, pre_append)
			tmp["distribute_list_out6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-DistributeListOut6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dont_capability_negotiate"
		if _, ok := i["dont-capability-negotiate"]; ok {
			v := flattenRouterBgpNeighborDontCapabilityNegotiate(i["dont-capability-negotiate"], d, pre_append)
			tmp["dont_capability_negotiate"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-DontCapabilityNegotiate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_enforce_multihop"
		if _, ok := i["ebgp-enforce-multihop"]; ok {
			v := flattenRouterBgpNeighborEbgpEnforceMultihop(i["ebgp-enforce-multihop"], d, pre_append)
			tmp["ebgp_enforce_multihop"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-EbgpEnforceMultihop")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_multihop_ttl"
		if _, ok := i["ebgp-multihop-ttl"]; ok {
			v := flattenRouterBgpNeighborEbgpMultihopTtl(i["ebgp-multihop-ttl"], d, pre_append)
			tmp["ebgp_multihop_ttl"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-EbgpMultihopTtl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in"
		if _, ok := i["filter-list-in"]; ok {
			v := flattenRouterBgpNeighborFilterListIn(i["filter-list-in"], d, pre_append)
			tmp["filter_list_in"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-FilterListIn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in_vpnv4"
		if _, ok := i["filter-list-in-vpnv4"]; ok {
			v := flattenRouterBgpNeighborFilterListInVpnv4(i["filter-list-in-vpnv4"], d, pre_append)
			tmp["filter_list_in_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-FilterListInVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in_vpnv6"
		if _, ok := i["filter-list-in-vpnv6"]; ok {
			v := flattenRouterBgpNeighborFilterListInVpnv6(i["filter-list-in-vpnv6"], d, pre_append)
			tmp["filter_list_in_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-FilterListInVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in6"
		if _, ok := i["filter-list-in6"]; ok {
			v := flattenRouterBgpNeighborFilterListIn6(i["filter-list-in6"], d, pre_append)
			tmp["filter_list_in6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-FilterListIn6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out"
		if _, ok := i["filter-list-out"]; ok {
			v := flattenRouterBgpNeighborFilterListOut(i["filter-list-out"], d, pre_append)
			tmp["filter_list_out"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-FilterListOut")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out_vpnv4"
		if _, ok := i["filter-list-out-vpnv4"]; ok {
			v := flattenRouterBgpNeighborFilterListOutVpnv4(i["filter-list-out-vpnv4"], d, pre_append)
			tmp["filter_list_out_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-FilterListOutVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out_vpnv6"
		if _, ok := i["filter-list-out-vpnv6"]; ok {
			v := flattenRouterBgpNeighborFilterListOutVpnv6(i["filter-list-out-vpnv6"], d, pre_append)
			tmp["filter_list_out_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-FilterListOutVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out6"
		if _, ok := i["filter-list-out6"]; ok {
			v := flattenRouterBgpNeighborFilterListOut6(i["filter-list-out6"], d, pre_append)
			tmp["filter_list_out6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-FilterListOut6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holdtime_timer"
		if _, ok := i["holdtime-timer"]; ok {
			v := flattenRouterBgpNeighborHoldtimeTimer(i["holdtime-timer"], d, pre_append)
			tmp["holdtime_timer"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-HoldtimeTimer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenRouterBgpNeighborInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenRouterBgpNeighborIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keep_alive_timer"
		if _, ok := i["keep-alive-timer"]; ok {
			v := flattenRouterBgpNeighborKeepAliveTimer(i["keep-alive-timer"], d, pre_append)
			tmp["keep_alive_timer"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-KeepAliveTimer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_down_failover"
		if _, ok := i["link-down-failover"]; ok {
			v := flattenRouterBgpNeighborLinkDownFailover(i["link-down-failover"], d, pre_append)
			tmp["link_down_failover"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-LinkDownFailover")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as"
		if _, ok := i["local-as"]; ok {
			v := flattenRouterBgpNeighborLocalAs(i["local-as"], d, pre_append)
			tmp["local_as"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-LocalAs")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as_no_prepend"
		if _, ok := i["local-as-no-prepend"]; ok {
			v := flattenRouterBgpNeighborLocalAsNoPrepend(i["local-as-no-prepend"], d, pre_append)
			tmp["local_as_no_prepend"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-LocalAsNoPrepend")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as_replace_as"
		if _, ok := i["local-as-replace-as"]; ok {
			v := flattenRouterBgpNeighborLocalAsReplaceAs(i["local-as-replace-as"], d, pre_append)
			tmp["local_as_replace_as"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-LocalAsReplaceAs")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix"
		if _, ok := i["maximum-prefix"]; ok {
			v := flattenRouterBgpNeighborMaximumPrefix(i["maximum-prefix"], d, pre_append)
			tmp["maximum_prefix"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-MaximumPrefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_evpn"
		if _, ok := i["maximum-prefix-evpn"]; ok {
			v := flattenRouterBgpNeighborMaximumPrefixEvpn(i["maximum-prefix-evpn"], d, pre_append)
			tmp["maximum_prefix_evpn"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-MaximumPrefixEvpn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold"
		if _, ok := i["maximum-prefix-threshold"]; ok {
			v := flattenRouterBgpNeighborMaximumPrefixThreshold(i["maximum-prefix-threshold"], d, pre_append)
			tmp["maximum_prefix_threshold"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-MaximumPrefixThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold_evpn"
		if _, ok := i["maximum-prefix-threshold-evpn"]; ok {
			v := flattenRouterBgpNeighborMaximumPrefixThresholdEvpn(i["maximum-prefix-threshold-evpn"], d, pre_append)
			tmp["maximum_prefix_threshold_evpn"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-MaximumPrefixThresholdEvpn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold_vpnv4"
		if _, ok := i["maximum-prefix-threshold-vpnv4"]; ok {
			v := flattenRouterBgpNeighborMaximumPrefixThresholdVpnv4(i["maximum-prefix-threshold-vpnv4"], d, pre_append)
			tmp["maximum_prefix_threshold_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-MaximumPrefixThresholdVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold_vpnv6"
		if _, ok := i["maximum-prefix-threshold-vpnv6"]; ok {
			v := flattenRouterBgpNeighborMaximumPrefixThresholdVpnv6(i["maximum-prefix-threshold-vpnv6"], d, pre_append)
			tmp["maximum_prefix_threshold_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-MaximumPrefixThresholdVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold6"
		if _, ok := i["maximum-prefix-threshold6"]; ok {
			v := flattenRouterBgpNeighborMaximumPrefixThreshold6(i["maximum-prefix-threshold6"], d, pre_append)
			tmp["maximum_prefix_threshold6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-MaximumPrefixThreshold6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_vpnv4"
		if _, ok := i["maximum-prefix-vpnv4"]; ok {
			v := flattenRouterBgpNeighborMaximumPrefixVpnv4(i["maximum-prefix-vpnv4"], d, pre_append)
			tmp["maximum_prefix_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-MaximumPrefixVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_vpnv6"
		if _, ok := i["maximum-prefix-vpnv6"]; ok {
			v := flattenRouterBgpNeighborMaximumPrefixVpnv6(i["maximum-prefix-vpnv6"], d, pre_append)
			tmp["maximum_prefix_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-MaximumPrefixVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only"
		if _, ok := i["maximum-prefix-warning-only"]; ok {
			v := flattenRouterBgpNeighborMaximumPrefixWarningOnly(i["maximum-prefix-warning-only"], d, pre_append)
			tmp["maximum_prefix_warning_only"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-MaximumPrefixWarningOnly")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only_evpn"
		if _, ok := i["maximum-prefix-warning-only-evpn"]; ok {
			v := flattenRouterBgpNeighborMaximumPrefixWarningOnlyEvpn(i["maximum-prefix-warning-only-evpn"], d, pre_append)
			tmp["maximum_prefix_warning_only_evpn"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-MaximumPrefixWarningOnlyEvpn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only_vpnv4"
		if _, ok := i["maximum-prefix-warning-only-vpnv4"]; ok {
			v := flattenRouterBgpNeighborMaximumPrefixWarningOnlyVpnv4(i["maximum-prefix-warning-only-vpnv4"], d, pre_append)
			tmp["maximum_prefix_warning_only_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-MaximumPrefixWarningOnlyVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only_vpnv6"
		if _, ok := i["maximum-prefix-warning-only-vpnv6"]; ok {
			v := flattenRouterBgpNeighborMaximumPrefixWarningOnlyVpnv6(i["maximum-prefix-warning-only-vpnv6"], d, pre_append)
			tmp["maximum_prefix_warning_only_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-MaximumPrefixWarningOnlyVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only6"
		if _, ok := i["maximum-prefix-warning-only6"]; ok {
			v := flattenRouterBgpNeighborMaximumPrefixWarningOnly6(i["maximum-prefix-warning-only6"], d, pre_append)
			tmp["maximum_prefix_warning_only6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-MaximumPrefixWarningOnly6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix6"
		if _, ok := i["maximum-prefix6"]; ok {
			v := flattenRouterBgpNeighborMaximumPrefix6(i["maximum-prefix6"], d, pre_append)
			tmp["maximum_prefix6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-MaximumPrefix6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self"
		if _, ok := i["next-hop-self"]; ok {
			v := flattenRouterBgpNeighborNextHopSelf(i["next-hop-self"], d, pre_append)
			tmp["next_hop_self"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-NextHopSelf")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_rr"
		if _, ok := i["next-hop-self-rr"]; ok {
			v := flattenRouterBgpNeighborNextHopSelfRr(i["next-hop-self-rr"], d, pre_append)
			tmp["next_hop_self_rr"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-NextHopSelfRr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_rr6"
		if _, ok := i["next-hop-self-rr6"]; ok {
			v := flattenRouterBgpNeighborNextHopSelfRr6(i["next-hop-self-rr6"], d, pre_append)
			tmp["next_hop_self_rr6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-NextHopSelfRr6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_vpnv4"
		if _, ok := i["next-hop-self-vpnv4"]; ok {
			v := flattenRouterBgpNeighborNextHopSelfVpnv4(i["next-hop-self-vpnv4"], d, pre_append)
			tmp["next_hop_self_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-NextHopSelfVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_vpnv6"
		if _, ok := i["next-hop-self-vpnv6"]; ok {
			v := flattenRouterBgpNeighborNextHopSelfVpnv6(i["next-hop-self-vpnv6"], d, pre_append)
			tmp["next_hop_self_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-NextHopSelfVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self6"
		if _, ok := i["next-hop-self6"]; ok {
			v := flattenRouterBgpNeighborNextHopSelf6(i["next-hop-self6"], d, pre_append)
			tmp["next_hop_self6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-NextHopSelf6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_capability"
		if _, ok := i["override-capability"]; ok {
			v := flattenRouterBgpNeighborOverrideCapability(i["override-capability"], d, pre_append)
			tmp["override_capability"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-OverrideCapability")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive"
		if _, ok := i["passive"]; ok {
			v := flattenRouterBgpNeighborPassive(i["passive"], d, pre_append)
			tmp["passive"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-Passive")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in"
		if _, ok := i["prefix-list-in"]; ok {
			v := flattenRouterBgpNeighborPrefixListIn(i["prefix-list-in"], d, pre_append)
			tmp["prefix_list_in"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-PrefixListIn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in_vpnv4"
		if _, ok := i["prefix-list-in-vpnv4"]; ok {
			v := flattenRouterBgpNeighborPrefixListInVpnv4(i["prefix-list-in-vpnv4"], d, pre_append)
			tmp["prefix_list_in_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-PrefixListInVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in_vpnv6"
		if _, ok := i["prefix-list-in-vpnv6"]; ok {
			v := flattenRouterBgpNeighborPrefixListInVpnv6(i["prefix-list-in-vpnv6"], d, pre_append)
			tmp["prefix_list_in_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-PrefixListInVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in6"
		if _, ok := i["prefix-list-in6"]; ok {
			v := flattenRouterBgpNeighborPrefixListIn6(i["prefix-list-in6"], d, pre_append)
			tmp["prefix_list_in6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-PrefixListIn6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out"
		if _, ok := i["prefix-list-out"]; ok {
			v := flattenRouterBgpNeighborPrefixListOut(i["prefix-list-out"], d, pre_append)
			tmp["prefix_list_out"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-PrefixListOut")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out_vpnv4"
		if _, ok := i["prefix-list-out-vpnv4"]; ok {
			v := flattenRouterBgpNeighborPrefixListOutVpnv4(i["prefix-list-out-vpnv4"], d, pre_append)
			tmp["prefix_list_out_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-PrefixListOutVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out_vpnv6"
		if _, ok := i["prefix-list-out-vpnv6"]; ok {
			v := flattenRouterBgpNeighborPrefixListOutVpnv6(i["prefix-list-out-vpnv6"], d, pre_append)
			tmp["prefix_list_out_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-PrefixListOutVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out6"
		if _, ok := i["prefix-list-out6"]; ok {
			v := flattenRouterBgpNeighborPrefixListOut6(i["prefix-list-out6"], d, pre_append)
			tmp["prefix_list_out6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-PrefixListOut6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_as"
		if _, ok := i["remote-as"]; ok {
			v := flattenRouterBgpNeighborRemoteAs(i["remote-as"], d, pre_append)
			tmp["remote_as"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RemoteAs")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as"
		if _, ok := i["remove-private-as"]; ok {
			v := flattenRouterBgpNeighborRemovePrivateAs(i["remove-private-as"], d, pre_append)
			tmp["remove_private_as"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RemovePrivateAs")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as_evpn"
		if _, ok := i["remove-private-as-evpn"]; ok {
			v := flattenRouterBgpNeighborRemovePrivateAsEvpn(i["remove-private-as-evpn"], d, pre_append)
			tmp["remove_private_as_evpn"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RemovePrivateAsEvpn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as_vpnv4"
		if _, ok := i["remove-private-as-vpnv4"]; ok {
			v := flattenRouterBgpNeighborRemovePrivateAsVpnv4(i["remove-private-as-vpnv4"], d, pre_append)
			tmp["remove_private_as_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RemovePrivateAsVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as_vpnv6"
		if _, ok := i["remove-private-as-vpnv6"]; ok {
			v := flattenRouterBgpNeighborRemovePrivateAsVpnv6(i["remove-private-as-vpnv6"], d, pre_append)
			tmp["remove_private_as_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RemovePrivateAsVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as6"
		if _, ok := i["remove-private-as6"]; ok {
			v := flattenRouterBgpNeighborRemovePrivateAs6(i["remove-private-as6"], d, pre_append)
			tmp["remove_private_as6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RemovePrivateAs6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restart_time"
		if _, ok := i["restart-time"]; ok {
			v := flattenRouterBgpNeighborRestartTime(i["restart-time"], d, pre_append)
			tmp["restart_time"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RestartTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retain_stale_time"
		if _, ok := i["retain-stale-time"]; ok {
			v := flattenRouterBgpNeighborRetainStaleTime(i["retain-stale-time"], d, pre_append)
			tmp["retain_stale_time"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RetainStaleTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in"
		if _, ok := i["route-map-in"]; ok {
			v := flattenRouterBgpNeighborRouteMapIn(i["route-map-in"], d, pre_append)
			tmp["route_map_in"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RouteMapIn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in_evpn"
		if _, ok := i["route-map-in-evpn"]; ok {
			v := flattenRouterBgpNeighborRouteMapInEvpn(i["route-map-in-evpn"], d, pre_append)
			tmp["route_map_in_evpn"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RouteMapInEvpn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in_vpnv4"
		if _, ok := i["route-map-in-vpnv4"]; ok {
			v := flattenRouterBgpNeighborRouteMapInVpnv4(i["route-map-in-vpnv4"], d, pre_append)
			tmp["route_map_in_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RouteMapInVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in_vpnv6"
		if _, ok := i["route-map-in-vpnv6"]; ok {
			v := flattenRouterBgpNeighborRouteMapInVpnv6(i["route-map-in-vpnv6"], d, pre_append)
			tmp["route_map_in_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RouteMapInVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in6"
		if _, ok := i["route-map-in6"]; ok {
			v := flattenRouterBgpNeighborRouteMapIn6(i["route-map-in6"], d, pre_append)
			tmp["route_map_in6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RouteMapIn6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out"
		if _, ok := i["route-map-out"]; ok {
			v := flattenRouterBgpNeighborRouteMapOut(i["route-map-out"], d, pre_append)
			tmp["route_map_out"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RouteMapOut")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_evpn"
		if _, ok := i["route-map-out-evpn"]; ok {
			v := flattenRouterBgpNeighborRouteMapOutEvpn(i["route-map-out-evpn"], d, pre_append)
			tmp["route_map_out_evpn"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RouteMapOutEvpn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_preferable"
		if _, ok := i["route-map-out-preferable"]; ok {
			v := flattenRouterBgpNeighborRouteMapOutPreferable(i["route-map-out-preferable"], d, pre_append)
			tmp["route_map_out_preferable"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RouteMapOutPreferable")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv4"
		if _, ok := i["route-map-out-vpnv4"]; ok {
			v := flattenRouterBgpNeighborRouteMapOutVpnv4(i["route-map-out-vpnv4"], d, pre_append)
			tmp["route_map_out_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RouteMapOutVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv4_preferable"
		if _, ok := i["route-map-out-vpnv4-preferable"]; ok {
			v := flattenRouterBgpNeighborRouteMapOutVpnv4Preferable(i["route-map-out-vpnv4-preferable"], d, pre_append)
			tmp["route_map_out_vpnv4_preferable"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RouteMapOutVpnv4Preferable")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv6"
		if _, ok := i["route-map-out-vpnv6"]; ok {
			v := flattenRouterBgpNeighborRouteMapOutVpnv6(i["route-map-out-vpnv6"], d, pre_append)
			tmp["route_map_out_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RouteMapOutVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv6_preferable"
		if _, ok := i["route-map-out-vpnv6-preferable"]; ok {
			v := flattenRouterBgpNeighborRouteMapOutVpnv6Preferable(i["route-map-out-vpnv6-preferable"], d, pre_append)
			tmp["route_map_out_vpnv6_preferable"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RouteMapOutVpnv6Preferable")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out6"
		if _, ok := i["route-map-out6"]; ok {
			v := flattenRouterBgpNeighborRouteMapOut6(i["route-map-out6"], d, pre_append)
			tmp["route_map_out6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RouteMapOut6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out6_preferable"
		if _, ok := i["route-map-out6-preferable"]; ok {
			v := flattenRouterBgpNeighborRouteMapOut6Preferable(i["route-map-out6-preferable"], d, pre_append)
			tmp["route_map_out6_preferable"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RouteMapOut6Preferable")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client"
		if _, ok := i["route-reflector-client"]; ok {
			v := flattenRouterBgpNeighborRouteReflectorClient(i["route-reflector-client"], d, pre_append)
			tmp["route_reflector_client"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RouteReflectorClient")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client_evpn"
		if _, ok := i["route-reflector-client-evpn"]; ok {
			v := flattenRouterBgpNeighborRouteReflectorClientEvpn(i["route-reflector-client-evpn"], d, pre_append)
			tmp["route_reflector_client_evpn"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RouteReflectorClientEvpn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client_vpnv4"
		if _, ok := i["route-reflector-client-vpnv4"]; ok {
			v := flattenRouterBgpNeighborRouteReflectorClientVpnv4(i["route-reflector-client-vpnv4"], d, pre_append)
			tmp["route_reflector_client_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RouteReflectorClientVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client_vpnv6"
		if _, ok := i["route-reflector-client-vpnv6"]; ok {
			v := flattenRouterBgpNeighborRouteReflectorClientVpnv6(i["route-reflector-client-vpnv6"], d, pre_append)
			tmp["route_reflector_client_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RouteReflectorClientVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client6"
		if _, ok := i["route-reflector-client6"]; ok {
			v := flattenRouterBgpNeighborRouteReflectorClient6(i["route-reflector-client6"], d, pre_append)
			tmp["route_reflector_client6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RouteReflectorClient6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client"
		if _, ok := i["route-server-client"]; ok {
			v := flattenRouterBgpNeighborRouteServerClient(i["route-server-client"], d, pre_append)
			tmp["route_server_client"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RouteServerClient")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client_evpn"
		if _, ok := i["route-server-client-evpn"]; ok {
			v := flattenRouterBgpNeighborRouteServerClientEvpn(i["route-server-client-evpn"], d, pre_append)
			tmp["route_server_client_evpn"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RouteServerClientEvpn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client_vpnv4"
		if _, ok := i["route-server-client-vpnv4"]; ok {
			v := flattenRouterBgpNeighborRouteServerClientVpnv4(i["route-server-client-vpnv4"], d, pre_append)
			tmp["route_server_client_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RouteServerClientVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client_vpnv6"
		if _, ok := i["route-server-client-vpnv6"]; ok {
			v := flattenRouterBgpNeighborRouteServerClientVpnv6(i["route-server-client-vpnv6"], d, pre_append)
			tmp["route_server_client_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RouteServerClientVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client6"
		if _, ok := i["route-server-client6"]; ok {
			v := flattenRouterBgpNeighborRouteServerClient6(i["route-server-client6"], d, pre_append)
			tmp["route_server_client6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-RouteServerClient6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community"
		if _, ok := i["send-community"]; ok {
			v := flattenRouterBgpNeighborSendCommunity(i["send-community"], d, pre_append)
			tmp["send_community"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-SendCommunity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community_evpn"
		if _, ok := i["send-community-evpn"]; ok {
			v := flattenRouterBgpNeighborSendCommunityEvpn(i["send-community-evpn"], d, pre_append)
			tmp["send_community_evpn"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-SendCommunityEvpn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community_vpnv4"
		if _, ok := i["send-community-vpnv4"]; ok {
			v := flattenRouterBgpNeighborSendCommunityVpnv4(i["send-community-vpnv4"], d, pre_append)
			tmp["send_community_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-SendCommunityVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community_vpnv6"
		if _, ok := i["send-community-vpnv6"]; ok {
			v := flattenRouterBgpNeighborSendCommunityVpnv6(i["send-community-vpnv6"], d, pre_append)
			tmp["send_community_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-SendCommunityVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community6"
		if _, ok := i["send-community6"]; ok {
			v := flattenRouterBgpNeighborSendCommunity6(i["send-community6"], d, pre_append)
			tmp["send_community6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-SendCommunity6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shutdown"
		if _, ok := i["shutdown"]; ok {
			v := flattenRouterBgpNeighborShutdown(i["shutdown"], d, pre_append)
			tmp["shutdown"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-Shutdown")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration"
		if _, ok := i["soft-reconfiguration"]; ok {
			v := flattenRouterBgpNeighborSoftReconfiguration(i["soft-reconfiguration"], d, pre_append)
			tmp["soft_reconfiguration"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-SoftReconfiguration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration_evpn"
		if _, ok := i["soft-reconfiguration-evpn"]; ok {
			v := flattenRouterBgpNeighborSoftReconfigurationEvpn(i["soft-reconfiguration-evpn"], d, pre_append)
			tmp["soft_reconfiguration_evpn"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-SoftReconfigurationEvpn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration_vpnv4"
		if _, ok := i["soft-reconfiguration-vpnv4"]; ok {
			v := flattenRouterBgpNeighborSoftReconfigurationVpnv4(i["soft-reconfiguration-vpnv4"], d, pre_append)
			tmp["soft_reconfiguration_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-SoftReconfigurationVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration_vpnv6"
		if _, ok := i["soft-reconfiguration-vpnv6"]; ok {
			v := flattenRouterBgpNeighborSoftReconfigurationVpnv6(i["soft-reconfiguration-vpnv6"], d, pre_append)
			tmp["soft_reconfiguration_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-SoftReconfigurationVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration6"
		if _, ok := i["soft-reconfiguration6"]; ok {
			v := flattenRouterBgpNeighborSoftReconfiguration6(i["soft-reconfiguration6"], d, pre_append)
			tmp["soft_reconfiguration6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-SoftReconfiguration6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stale_route"
		if _, ok := i["stale-route"]; ok {
			v := flattenRouterBgpNeighborStaleRoute(i["stale-route"], d, pre_append)
			tmp["stale_route"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-StaleRoute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "strict_capability_match"
		if _, ok := i["strict-capability-match"]; ok {
			v := flattenRouterBgpNeighborStrictCapabilityMatch(i["strict-capability-match"], d, pre_append)
			tmp["strict_capability_match"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-StrictCapabilityMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unsuppress_map"
		if _, ok := i["unsuppress-map"]; ok {
			v := flattenRouterBgpNeighborUnsuppressMap(i["unsuppress-map"], d, pre_append)
			tmp["unsuppress_map"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-UnsuppressMap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unsuppress_map6"
		if _, ok := i["unsuppress-map6"]; ok {
			v := flattenRouterBgpNeighborUnsuppressMap6(i["unsuppress-map6"], d, pre_append)
			tmp["unsuppress_map6"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-UnsuppressMap6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_source"
		if _, ok := i["update-source"]; ok {
			v := flattenRouterBgpNeighborUpdateSource(i["update-source"], d, pre_append)
			tmp["update_source"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-UpdateSource")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenRouterBgpNeighborWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "RouterBgp-Neighbor-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterBgpNeighborActivate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborActivateEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborActivateVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborActivateVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborActivate6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAdditionalPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAdditionalPathVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAdditionalPathVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAdditionalPath6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAdvAdditionalPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAdvAdditionalPathVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAdvAdditionalPathVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAdvAdditionalPath6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAdvertisementInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAllowasIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAllowasInEnable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAllowasInEnableEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAllowasInEnableVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAllowasInEnableVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAllowasInEnable6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAllowasInEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAllowasInVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAllowasInVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAllowasIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAsOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAsOverride6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAttributeUnchanged(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborAttributeUnchangedVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborAttributeUnchangedVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborAttributeUnchanged6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborAuthOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityDefaultOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityDefaultOriginate6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityDynamic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityGracefulRestart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityGracefulRestartEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityGracefulRestartVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityGracefulRestartVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityGracefulRestart6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityOrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityOrf6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityRouteRefresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborConditionalAdvertise(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise_routemap"
		if _, ok := i["advertise-routemap"]; ok {
			v := flattenRouterBgpNeighborConditionalAdvertiseAdvertiseRoutemap(i["advertise-routemap"], d, pre_append)
			tmp["advertise_routemap"] = fortiAPISubPartPatch(v, "RouterBgpNeighbor-ConditionalAdvertise-AdvertiseRoutemap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_routemap"
		if _, ok := i["condition-routemap"]; ok {
			v := flattenRouterBgpNeighborConditionalAdvertiseConditionRoutemap(i["condition-routemap"], d, pre_append)
			tmp["condition_routemap"] = fortiAPISubPartPatch(v, "RouterBgpNeighbor-ConditionalAdvertise-ConditionRoutemap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_type"
		if _, ok := i["condition-type"]; ok {
			v := flattenRouterBgpNeighborConditionalAdvertiseConditionType(i["condition-type"], d, pre_append)
			tmp["condition_type"] = fortiAPISubPartPatch(v, "RouterBgpNeighbor-ConditionalAdvertise-ConditionType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterBgpNeighborConditionalAdvertiseAdvertiseRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborConditionalAdvertiseConditionRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborConditionalAdvertiseConditionType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborConditionalAdvertise6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise_routemap"
		if _, ok := i["advertise-routemap"]; ok {
			v := flattenRouterBgpNeighborConditionalAdvertise6AdvertiseRoutemap(i["advertise-routemap"], d, pre_append)
			tmp["advertise_routemap"] = fortiAPISubPartPatch(v, "RouterBgpNeighbor-ConditionalAdvertise6-AdvertiseRoutemap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_routemap"
		if _, ok := i["condition-routemap"]; ok {
			v := flattenRouterBgpNeighborConditionalAdvertise6ConditionRoutemap(i["condition-routemap"], d, pre_append)
			tmp["condition_routemap"] = fortiAPISubPartPatch(v, "RouterBgpNeighbor-ConditionalAdvertise6-ConditionRoutemap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_type"
		if _, ok := i["condition-type"]; ok {
			v := flattenRouterBgpNeighborConditionalAdvertise6ConditionType(i["condition-type"], d, pre_append)
			tmp["condition_type"] = fortiAPISubPartPatch(v, "RouterBgpNeighbor-ConditionalAdvertise6-ConditionType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterBgpNeighborConditionalAdvertise6AdvertiseRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborConditionalAdvertise6ConditionRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborConditionalAdvertise6ConditionType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborConnectTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborDefaultOriginateRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborDefaultOriginateRoutemap6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborDistributeListIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborDistributeListInVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborDistributeListInVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborDistributeListIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborDistributeListOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborDistributeListOutVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborDistributeListOutVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborDistributeListOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborDontCapabilityNegotiate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborEbgpEnforceMultihop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborEbgpMultihopTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborFilterListIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborFilterListInVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborFilterListInVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborFilterListIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborFilterListOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborFilterListOutVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborFilterListOutVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborFilterListOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborHoldtimeTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborKeepAliveTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborLinkDownFailover(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborLocalAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborLocalAsNoPrepend(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborLocalAsReplaceAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixThresholdEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixThresholdVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixThresholdVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixThreshold6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixWarningOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixWarningOnlyEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixWarningOnlyVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixWarningOnlyVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixWarningOnly6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborNextHopSelf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborNextHopSelfRr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborNextHopSelfRr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborNextHopSelfVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborNextHopSelfVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborNextHopSelf6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborOverrideCapability(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborPassive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborPrefixListIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborPrefixListInVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborPrefixListInVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborPrefixListIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborPrefixListOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborPrefixListOutVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborPrefixListOutVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborPrefixListOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRemoteAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRemovePrivateAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRemovePrivateAsEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRemovePrivateAsVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRemovePrivateAsVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRemovePrivateAs6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRestartTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRetainStaleTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteMapIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRouteMapInEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRouteMapInVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRouteMapInVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRouteMapIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRouteMapOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRouteMapOutEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRouteMapOutPreferable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRouteMapOutVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRouteMapOutVpnv4Preferable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRouteMapOutVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRouteMapOutVpnv6Preferable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRouteMapOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRouteMapOut6Preferable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRouteReflectorClient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteReflectorClientEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteReflectorClientVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteReflectorClientVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteReflectorClient6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteServerClient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteServerClientEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteServerClientVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteServerClientVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteServerClient6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborSendCommunity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborSendCommunityEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborSendCommunityVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborSendCommunityVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborSendCommunity6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborShutdown(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborSoftReconfiguration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborSoftReconfigurationEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborSoftReconfigurationVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborSoftReconfigurationVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborSoftReconfiguration6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborStaleRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborStrictCapabilityMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborUnsuppressMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborUnsuppressMap6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborUpdateSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate"
		if _, ok := i["activate"]; ok {
			v := flattenRouterBgpNeighborGroupActivate(i["activate"], d, pre_append)
			tmp["activate"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-Activate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate_evpn"
		if _, ok := i["activate-evpn"]; ok {
			v := flattenRouterBgpNeighborGroupActivateEvpn(i["activate-evpn"], d, pre_append)
			tmp["activate_evpn"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-ActivateEvpn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate_vpnv4"
		if _, ok := i["activate-vpnv4"]; ok {
			v := flattenRouterBgpNeighborGroupActivateVpnv4(i["activate-vpnv4"], d, pre_append)
			tmp["activate_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-ActivateVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate_vpnv6"
		if _, ok := i["activate-vpnv6"]; ok {
			v := flattenRouterBgpNeighborGroupActivateVpnv6(i["activate-vpnv6"], d, pre_append)
			tmp["activate_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-ActivateVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate6"
		if _, ok := i["activate6"]; ok {
			v := flattenRouterBgpNeighborGroupActivate6(i["activate6"], d, pre_append)
			tmp["activate6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-Activate6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path"
		if _, ok := i["additional-path"]; ok {
			v := flattenRouterBgpNeighborGroupAdditionalPath(i["additional-path"], d, pre_append)
			tmp["additional_path"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-AdditionalPath")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path_vpnv4"
		if _, ok := i["additional-path-vpnv4"]; ok {
			v := flattenRouterBgpNeighborGroupAdditionalPathVpnv4(i["additional-path-vpnv4"], d, pre_append)
			tmp["additional_path_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-AdditionalPathVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path_vpnv6"
		if _, ok := i["additional-path-vpnv6"]; ok {
			v := flattenRouterBgpNeighborGroupAdditionalPathVpnv6(i["additional-path-vpnv6"], d, pre_append)
			tmp["additional_path_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-AdditionalPathVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path6"
		if _, ok := i["additional-path6"]; ok {
			v := flattenRouterBgpNeighborGroupAdditionalPath6(i["additional-path6"], d, pre_append)
			tmp["additional_path6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-AdditionalPath6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path"
		if _, ok := i["adv-additional-path"]; ok {
			v := flattenRouterBgpNeighborGroupAdvAdditionalPath(i["adv-additional-path"], d, pre_append)
			tmp["adv_additional_path"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-AdvAdditionalPath")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path_vpnv4"
		if _, ok := i["adv-additional-path-vpnv4"]; ok {
			v := flattenRouterBgpNeighborGroupAdvAdditionalPathVpnv4(i["adv-additional-path-vpnv4"], d, pre_append)
			tmp["adv_additional_path_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-AdvAdditionalPathVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path_vpnv6"
		if _, ok := i["adv-additional-path-vpnv6"]; ok {
			v := flattenRouterBgpNeighborGroupAdvAdditionalPathVpnv6(i["adv-additional-path-vpnv6"], d, pre_append)
			tmp["adv_additional_path_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-AdvAdditionalPathVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path6"
		if _, ok := i["adv-additional-path6"]; ok {
			v := flattenRouterBgpNeighborGroupAdvAdditionalPath6(i["adv-additional-path6"], d, pre_append)
			tmp["adv_additional_path6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-AdvAdditionalPath6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertisement_interval"
		if _, ok := i["advertisement-interval"]; ok {
			v := flattenRouterBgpNeighborGroupAdvertisementInterval(i["advertisement-interval"], d, pre_append)
			tmp["advertisement_interval"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-AdvertisementInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in"
		if _, ok := i["allowas-in"]; ok {
			v := flattenRouterBgpNeighborGroupAllowasIn(i["allowas-in"], d, pre_append)
			tmp["allowas_in"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-AllowasIn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable"
		if _, ok := i["allowas-in-enable"]; ok {
			v := flattenRouterBgpNeighborGroupAllowasInEnable(i["allowas-in-enable"], d, pre_append)
			tmp["allowas_in_enable"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-AllowasInEnable")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable_evpn"
		if _, ok := i["allowas-in-enable-evpn"]; ok {
			v := flattenRouterBgpNeighborGroupAllowasInEnableEvpn(i["allowas-in-enable-evpn"], d, pre_append)
			tmp["allowas_in_enable_evpn"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-AllowasInEnableEvpn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable_vpnv4"
		if _, ok := i["allowas-in-enable-vpnv4"]; ok {
			v := flattenRouterBgpNeighborGroupAllowasInEnableVpnv4(i["allowas-in-enable-vpnv4"], d, pre_append)
			tmp["allowas_in_enable_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-AllowasInEnableVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable_vpnv6"
		if _, ok := i["allowas-in-enable-vpnv6"]; ok {
			v := flattenRouterBgpNeighborGroupAllowasInEnableVpnv6(i["allowas-in-enable-vpnv6"], d, pre_append)
			tmp["allowas_in_enable_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-AllowasInEnableVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable6"
		if _, ok := i["allowas-in-enable6"]; ok {
			v := flattenRouterBgpNeighborGroupAllowasInEnable6(i["allowas-in-enable6"], d, pre_append)
			tmp["allowas_in_enable6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-AllowasInEnable6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_evpn"
		if _, ok := i["allowas-in-evpn"]; ok {
			v := flattenRouterBgpNeighborGroupAllowasInEvpn(i["allowas-in-evpn"], d, pre_append)
			tmp["allowas_in_evpn"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-AllowasInEvpn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_vpnv4"
		if _, ok := i["allowas-in-vpnv4"]; ok {
			v := flattenRouterBgpNeighborGroupAllowasInVpnv4(i["allowas-in-vpnv4"], d, pre_append)
			tmp["allowas_in_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-AllowasInVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_vpnv6"
		if _, ok := i["allowas-in-vpnv6"]; ok {
			v := flattenRouterBgpNeighborGroupAllowasInVpnv6(i["allowas-in-vpnv6"], d, pre_append)
			tmp["allowas_in_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-AllowasInVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in6"
		if _, ok := i["allowas-in6"]; ok {
			v := flattenRouterBgpNeighborGroupAllowasIn6(i["allowas-in6"], d, pre_append)
			tmp["allowas_in6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-AllowasIn6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_override"
		if _, ok := i["as-override"]; ok {
			v := flattenRouterBgpNeighborGroupAsOverride(i["as-override"], d, pre_append)
			tmp["as_override"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-AsOverride")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_override6"
		if _, ok := i["as-override6"]; ok {
			v := flattenRouterBgpNeighborGroupAsOverride6(i["as-override6"], d, pre_append)
			tmp["as_override6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-AsOverride6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged"
		if _, ok := i["attribute-unchanged"]; ok {
			v := flattenRouterBgpNeighborGroupAttributeUnchanged(i["attribute-unchanged"], d, pre_append)
			tmp["attribute_unchanged"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-AttributeUnchanged")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged_vpnv4"
		if _, ok := i["attribute-unchanged-vpnv4"]; ok {
			v := flattenRouterBgpNeighborGroupAttributeUnchangedVpnv4(i["attribute-unchanged-vpnv4"], d, pre_append)
			tmp["attribute_unchanged_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-AttributeUnchangedVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged_vpnv6"
		if _, ok := i["attribute-unchanged-vpnv6"]; ok {
			v := flattenRouterBgpNeighborGroupAttributeUnchangedVpnv6(i["attribute-unchanged-vpnv6"], d, pre_append)
			tmp["attribute_unchanged_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-AttributeUnchangedVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged6"
		if _, ok := i["attribute-unchanged6"]; ok {
			v := flattenRouterBgpNeighborGroupAttributeUnchanged6(i["attribute-unchanged6"], d, pre_append)
			tmp["attribute_unchanged6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-AttributeUnchanged6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_options"
		if _, ok := i["auth-options"]; ok {
			v := flattenRouterBgpNeighborGroupAuthOptions(i["auth-options"], d, pre_append)
			tmp["auth_options"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-AuthOptions")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := i["bfd"]; ok {
			v := flattenRouterBgpNeighborGroupBfd(i["bfd"], d, pre_append)
			tmp["bfd"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-Bfd")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_default_originate"
		if _, ok := i["capability-default-originate"]; ok {
			v := flattenRouterBgpNeighborGroupCapabilityDefaultOriginate(i["capability-default-originate"], d, pre_append)
			tmp["capability_default_originate"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-CapabilityDefaultOriginate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_default_originate6"
		if _, ok := i["capability-default-originate6"]; ok {
			v := flattenRouterBgpNeighborGroupCapabilityDefaultOriginate6(i["capability-default-originate6"], d, pre_append)
			tmp["capability_default_originate6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-CapabilityDefaultOriginate6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_dynamic"
		if _, ok := i["capability-dynamic"]; ok {
			v := flattenRouterBgpNeighborGroupCapabilityDynamic(i["capability-dynamic"], d, pre_append)
			tmp["capability_dynamic"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-CapabilityDynamic")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart"
		if _, ok := i["capability-graceful-restart"]; ok {
			v := flattenRouterBgpNeighborGroupCapabilityGracefulRestart(i["capability-graceful-restart"], d, pre_append)
			tmp["capability_graceful_restart"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-CapabilityGracefulRestart")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart_evpn"
		if _, ok := i["capability-graceful-restart-evpn"]; ok {
			v := flattenRouterBgpNeighborGroupCapabilityGracefulRestartEvpn(i["capability-graceful-restart-evpn"], d, pre_append)
			tmp["capability_graceful_restart_evpn"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-CapabilityGracefulRestartEvpn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart_vpnv4"
		if _, ok := i["capability-graceful-restart-vpnv4"]; ok {
			v := flattenRouterBgpNeighborGroupCapabilityGracefulRestartVpnv4(i["capability-graceful-restart-vpnv4"], d, pre_append)
			tmp["capability_graceful_restart_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-CapabilityGracefulRestartVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart_vpnv6"
		if _, ok := i["capability-graceful-restart-vpnv6"]; ok {
			v := flattenRouterBgpNeighborGroupCapabilityGracefulRestartVpnv6(i["capability-graceful-restart-vpnv6"], d, pre_append)
			tmp["capability_graceful_restart_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-CapabilityGracefulRestartVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart6"
		if _, ok := i["capability-graceful-restart6"]; ok {
			v := flattenRouterBgpNeighborGroupCapabilityGracefulRestart6(i["capability-graceful-restart6"], d, pre_append)
			tmp["capability_graceful_restart6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-CapabilityGracefulRestart6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_orf"
		if _, ok := i["capability-orf"]; ok {
			v := flattenRouterBgpNeighborGroupCapabilityOrf(i["capability-orf"], d, pre_append)
			tmp["capability_orf"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-CapabilityOrf")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_orf6"
		if _, ok := i["capability-orf6"]; ok {
			v := flattenRouterBgpNeighborGroupCapabilityOrf6(i["capability-orf6"], d, pre_append)
			tmp["capability_orf6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-CapabilityOrf6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_route_refresh"
		if _, ok := i["capability-route-refresh"]; ok {
			v := flattenRouterBgpNeighborGroupCapabilityRouteRefresh(i["capability-route-refresh"], d, pre_append)
			tmp["capability_route_refresh"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-CapabilityRouteRefresh")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "connect_timer"
		if _, ok := i["connect-timer"]; ok {
			v := flattenRouterBgpNeighborGroupConnectTimer(i["connect-timer"], d, pre_append)
			tmp["connect_timer"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-ConnectTimer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_originate_routemap"
		if _, ok := i["default-originate-routemap"]; ok {
			v := flattenRouterBgpNeighborGroupDefaultOriginateRoutemap(i["default-originate-routemap"], d, pre_append)
			tmp["default_originate_routemap"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-DefaultOriginateRoutemap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_originate_routemap6"
		if _, ok := i["default-originate-routemap6"]; ok {
			v := flattenRouterBgpNeighborGroupDefaultOriginateRoutemap6(i["default-originate-routemap6"], d, pre_append)
			tmp["default_originate_routemap6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-DefaultOriginateRoutemap6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenRouterBgpNeighborGroupDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in"
		if _, ok := i["distribute-list-in"]; ok {
			v := flattenRouterBgpNeighborGroupDistributeListIn(i["distribute-list-in"], d, pre_append)
			tmp["distribute_list_in"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-DistributeListIn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in_vpnv4"
		if _, ok := i["distribute-list-in-vpnv4"]; ok {
			v := flattenRouterBgpNeighborGroupDistributeListInVpnv4(i["distribute-list-in-vpnv4"], d, pre_append)
			tmp["distribute_list_in_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-DistributeListInVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in_vpnv6"
		if _, ok := i["distribute-list-in-vpnv6"]; ok {
			v := flattenRouterBgpNeighborGroupDistributeListInVpnv6(i["distribute-list-in-vpnv6"], d, pre_append)
			tmp["distribute_list_in_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-DistributeListInVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in6"
		if _, ok := i["distribute-list-in6"]; ok {
			v := flattenRouterBgpNeighborGroupDistributeListIn6(i["distribute-list-in6"], d, pre_append)
			tmp["distribute_list_in6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-DistributeListIn6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out"
		if _, ok := i["distribute-list-out"]; ok {
			v := flattenRouterBgpNeighborGroupDistributeListOut(i["distribute-list-out"], d, pre_append)
			tmp["distribute_list_out"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-DistributeListOut")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out_vpnv4"
		if _, ok := i["distribute-list-out-vpnv4"]; ok {
			v := flattenRouterBgpNeighborGroupDistributeListOutVpnv4(i["distribute-list-out-vpnv4"], d, pre_append)
			tmp["distribute_list_out_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-DistributeListOutVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out_vpnv6"
		if _, ok := i["distribute-list-out-vpnv6"]; ok {
			v := flattenRouterBgpNeighborGroupDistributeListOutVpnv6(i["distribute-list-out-vpnv6"], d, pre_append)
			tmp["distribute_list_out_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-DistributeListOutVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out6"
		if _, ok := i["distribute-list-out6"]; ok {
			v := flattenRouterBgpNeighborGroupDistributeListOut6(i["distribute-list-out6"], d, pre_append)
			tmp["distribute_list_out6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-DistributeListOut6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dont_capability_negotiate"
		if _, ok := i["dont-capability-negotiate"]; ok {
			v := flattenRouterBgpNeighborGroupDontCapabilityNegotiate(i["dont-capability-negotiate"], d, pre_append)
			tmp["dont_capability_negotiate"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-DontCapabilityNegotiate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_enforce_multihop"
		if _, ok := i["ebgp-enforce-multihop"]; ok {
			v := flattenRouterBgpNeighborGroupEbgpEnforceMultihop(i["ebgp-enforce-multihop"], d, pre_append)
			tmp["ebgp_enforce_multihop"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-EbgpEnforceMultihop")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_multihop_ttl"
		if _, ok := i["ebgp-multihop-ttl"]; ok {
			v := flattenRouterBgpNeighborGroupEbgpMultihopTtl(i["ebgp-multihop-ttl"], d, pre_append)
			tmp["ebgp_multihop_ttl"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-EbgpMultihopTtl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in"
		if _, ok := i["filter-list-in"]; ok {
			v := flattenRouterBgpNeighborGroupFilterListIn(i["filter-list-in"], d, pre_append)
			tmp["filter_list_in"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-FilterListIn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in_vpnv4"
		if _, ok := i["filter-list-in-vpnv4"]; ok {
			v := flattenRouterBgpNeighborGroupFilterListInVpnv4(i["filter-list-in-vpnv4"], d, pre_append)
			tmp["filter_list_in_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-FilterListInVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in_vpnv6"
		if _, ok := i["filter-list-in-vpnv6"]; ok {
			v := flattenRouterBgpNeighborGroupFilterListInVpnv6(i["filter-list-in-vpnv6"], d, pre_append)
			tmp["filter_list_in_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-FilterListInVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in6"
		if _, ok := i["filter-list-in6"]; ok {
			v := flattenRouterBgpNeighborGroupFilterListIn6(i["filter-list-in6"], d, pre_append)
			tmp["filter_list_in6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-FilterListIn6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out"
		if _, ok := i["filter-list-out"]; ok {
			v := flattenRouterBgpNeighborGroupFilterListOut(i["filter-list-out"], d, pre_append)
			tmp["filter_list_out"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-FilterListOut")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out_vpnv4"
		if _, ok := i["filter-list-out-vpnv4"]; ok {
			v := flattenRouterBgpNeighborGroupFilterListOutVpnv4(i["filter-list-out-vpnv4"], d, pre_append)
			tmp["filter_list_out_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-FilterListOutVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out_vpnv6"
		if _, ok := i["filter-list-out-vpnv6"]; ok {
			v := flattenRouterBgpNeighborGroupFilterListOutVpnv6(i["filter-list-out-vpnv6"], d, pre_append)
			tmp["filter_list_out_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-FilterListOutVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out6"
		if _, ok := i["filter-list-out6"]; ok {
			v := flattenRouterBgpNeighborGroupFilterListOut6(i["filter-list-out6"], d, pre_append)
			tmp["filter_list_out6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-FilterListOut6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holdtime_timer"
		if _, ok := i["holdtime-timer"]; ok {
			v := flattenRouterBgpNeighborGroupHoldtimeTimer(i["holdtime-timer"], d, pre_append)
			tmp["holdtime_timer"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-HoldtimeTimer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenRouterBgpNeighborGroupInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keep_alive_timer"
		if _, ok := i["keep-alive-timer"]; ok {
			v := flattenRouterBgpNeighborGroupKeepAliveTimer(i["keep-alive-timer"], d, pre_append)
			tmp["keep_alive_timer"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-KeepAliveTimer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_down_failover"
		if _, ok := i["link-down-failover"]; ok {
			v := flattenRouterBgpNeighborGroupLinkDownFailover(i["link-down-failover"], d, pre_append)
			tmp["link_down_failover"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-LinkDownFailover")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as"
		if _, ok := i["local-as"]; ok {
			v := flattenRouterBgpNeighborGroupLocalAs(i["local-as"], d, pre_append)
			tmp["local_as"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-LocalAs")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as_no_prepend"
		if _, ok := i["local-as-no-prepend"]; ok {
			v := flattenRouterBgpNeighborGroupLocalAsNoPrepend(i["local-as-no-prepend"], d, pre_append)
			tmp["local_as_no_prepend"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-LocalAsNoPrepend")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as_replace_as"
		if _, ok := i["local-as-replace-as"]; ok {
			v := flattenRouterBgpNeighborGroupLocalAsReplaceAs(i["local-as-replace-as"], d, pre_append)
			tmp["local_as_replace_as"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-LocalAsReplaceAs")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix"
		if _, ok := i["maximum-prefix"]; ok {
			v := flattenRouterBgpNeighborGroupMaximumPrefix(i["maximum-prefix"], d, pre_append)
			tmp["maximum_prefix"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-MaximumPrefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_evpn"
		if _, ok := i["maximum-prefix-evpn"]; ok {
			v := flattenRouterBgpNeighborGroupMaximumPrefixEvpn(i["maximum-prefix-evpn"], d, pre_append)
			tmp["maximum_prefix_evpn"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-MaximumPrefixEvpn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold"
		if _, ok := i["maximum-prefix-threshold"]; ok {
			v := flattenRouterBgpNeighborGroupMaximumPrefixThreshold(i["maximum-prefix-threshold"], d, pre_append)
			tmp["maximum_prefix_threshold"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-MaximumPrefixThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold_evpn"
		if _, ok := i["maximum-prefix-threshold-evpn"]; ok {
			v := flattenRouterBgpNeighborGroupMaximumPrefixThresholdEvpn(i["maximum-prefix-threshold-evpn"], d, pre_append)
			tmp["maximum_prefix_threshold_evpn"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-MaximumPrefixThresholdEvpn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold_vpnv4"
		if _, ok := i["maximum-prefix-threshold-vpnv4"]; ok {
			v := flattenRouterBgpNeighborGroupMaximumPrefixThresholdVpnv4(i["maximum-prefix-threshold-vpnv4"], d, pre_append)
			tmp["maximum_prefix_threshold_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-MaximumPrefixThresholdVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold_vpnv6"
		if _, ok := i["maximum-prefix-threshold-vpnv6"]; ok {
			v := flattenRouterBgpNeighborGroupMaximumPrefixThresholdVpnv6(i["maximum-prefix-threshold-vpnv6"], d, pre_append)
			tmp["maximum_prefix_threshold_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-MaximumPrefixThresholdVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold6"
		if _, ok := i["maximum-prefix-threshold6"]; ok {
			v := flattenRouterBgpNeighborGroupMaximumPrefixThreshold6(i["maximum-prefix-threshold6"], d, pre_append)
			tmp["maximum_prefix_threshold6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-MaximumPrefixThreshold6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_vpnv4"
		if _, ok := i["maximum-prefix-vpnv4"]; ok {
			v := flattenRouterBgpNeighborGroupMaximumPrefixVpnv4(i["maximum-prefix-vpnv4"], d, pre_append)
			tmp["maximum_prefix_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-MaximumPrefixVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_vpnv6"
		if _, ok := i["maximum-prefix-vpnv6"]; ok {
			v := flattenRouterBgpNeighborGroupMaximumPrefixVpnv6(i["maximum-prefix-vpnv6"], d, pre_append)
			tmp["maximum_prefix_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-MaximumPrefixVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only"
		if _, ok := i["maximum-prefix-warning-only"]; ok {
			v := flattenRouterBgpNeighborGroupMaximumPrefixWarningOnly(i["maximum-prefix-warning-only"], d, pre_append)
			tmp["maximum_prefix_warning_only"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-MaximumPrefixWarningOnly")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only_evpn"
		if _, ok := i["maximum-prefix-warning-only-evpn"]; ok {
			v := flattenRouterBgpNeighborGroupMaximumPrefixWarningOnlyEvpn(i["maximum-prefix-warning-only-evpn"], d, pre_append)
			tmp["maximum_prefix_warning_only_evpn"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-MaximumPrefixWarningOnlyEvpn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only_vpnv4"
		if _, ok := i["maximum-prefix-warning-only-vpnv4"]; ok {
			v := flattenRouterBgpNeighborGroupMaximumPrefixWarningOnlyVpnv4(i["maximum-prefix-warning-only-vpnv4"], d, pre_append)
			tmp["maximum_prefix_warning_only_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-MaximumPrefixWarningOnlyVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only_vpnv6"
		if _, ok := i["maximum-prefix-warning-only-vpnv6"]; ok {
			v := flattenRouterBgpNeighborGroupMaximumPrefixWarningOnlyVpnv6(i["maximum-prefix-warning-only-vpnv6"], d, pre_append)
			tmp["maximum_prefix_warning_only_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-MaximumPrefixWarningOnlyVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only6"
		if _, ok := i["maximum-prefix-warning-only6"]; ok {
			v := flattenRouterBgpNeighborGroupMaximumPrefixWarningOnly6(i["maximum-prefix-warning-only6"], d, pre_append)
			tmp["maximum_prefix_warning_only6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-MaximumPrefixWarningOnly6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix6"
		if _, ok := i["maximum-prefix6"]; ok {
			v := flattenRouterBgpNeighborGroupMaximumPrefix6(i["maximum-prefix6"], d, pre_append)
			tmp["maximum_prefix6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-MaximumPrefix6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenRouterBgpNeighborGroupName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self"
		if _, ok := i["next-hop-self"]; ok {
			v := flattenRouterBgpNeighborGroupNextHopSelf(i["next-hop-self"], d, pre_append)
			tmp["next_hop_self"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-NextHopSelf")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_rr"
		if _, ok := i["next-hop-self-rr"]; ok {
			v := flattenRouterBgpNeighborGroupNextHopSelfRr(i["next-hop-self-rr"], d, pre_append)
			tmp["next_hop_self_rr"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-NextHopSelfRr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_rr6"
		if _, ok := i["next-hop-self-rr6"]; ok {
			v := flattenRouterBgpNeighborGroupNextHopSelfRr6(i["next-hop-self-rr6"], d, pre_append)
			tmp["next_hop_self_rr6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-NextHopSelfRr6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_vpnv4"
		if _, ok := i["next-hop-self-vpnv4"]; ok {
			v := flattenRouterBgpNeighborGroupNextHopSelfVpnv4(i["next-hop-self-vpnv4"], d, pre_append)
			tmp["next_hop_self_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-NextHopSelfVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_vpnv6"
		if _, ok := i["next-hop-self-vpnv6"]; ok {
			v := flattenRouterBgpNeighborGroupNextHopSelfVpnv6(i["next-hop-self-vpnv6"], d, pre_append)
			tmp["next_hop_self_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-NextHopSelfVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self6"
		if _, ok := i["next-hop-self6"]; ok {
			v := flattenRouterBgpNeighborGroupNextHopSelf6(i["next-hop-self6"], d, pre_append)
			tmp["next_hop_self6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-NextHopSelf6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_capability"
		if _, ok := i["override-capability"]; ok {
			v := flattenRouterBgpNeighborGroupOverrideCapability(i["override-capability"], d, pre_append)
			tmp["override_capability"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-OverrideCapability")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive"
		if _, ok := i["passive"]; ok {
			v := flattenRouterBgpNeighborGroupPassive(i["passive"], d, pre_append)
			tmp["passive"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-Passive")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in"
		if _, ok := i["prefix-list-in"]; ok {
			v := flattenRouterBgpNeighborGroupPrefixListIn(i["prefix-list-in"], d, pre_append)
			tmp["prefix_list_in"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-PrefixListIn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in_vpnv4"
		if _, ok := i["prefix-list-in-vpnv4"]; ok {
			v := flattenRouterBgpNeighborGroupPrefixListInVpnv4(i["prefix-list-in-vpnv4"], d, pre_append)
			tmp["prefix_list_in_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-PrefixListInVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in_vpnv6"
		if _, ok := i["prefix-list-in-vpnv6"]; ok {
			v := flattenRouterBgpNeighborGroupPrefixListInVpnv6(i["prefix-list-in-vpnv6"], d, pre_append)
			tmp["prefix_list_in_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-PrefixListInVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in6"
		if _, ok := i["prefix-list-in6"]; ok {
			v := flattenRouterBgpNeighborGroupPrefixListIn6(i["prefix-list-in6"], d, pre_append)
			tmp["prefix_list_in6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-PrefixListIn6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out"
		if _, ok := i["prefix-list-out"]; ok {
			v := flattenRouterBgpNeighborGroupPrefixListOut(i["prefix-list-out"], d, pre_append)
			tmp["prefix_list_out"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-PrefixListOut")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out_vpnv4"
		if _, ok := i["prefix-list-out-vpnv4"]; ok {
			v := flattenRouterBgpNeighborGroupPrefixListOutVpnv4(i["prefix-list-out-vpnv4"], d, pre_append)
			tmp["prefix_list_out_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-PrefixListOutVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out_vpnv6"
		if _, ok := i["prefix-list-out-vpnv6"]; ok {
			v := flattenRouterBgpNeighborGroupPrefixListOutVpnv6(i["prefix-list-out-vpnv6"], d, pre_append)
			tmp["prefix_list_out_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-PrefixListOutVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out6"
		if _, ok := i["prefix-list-out6"]; ok {
			v := flattenRouterBgpNeighborGroupPrefixListOut6(i["prefix-list-out6"], d, pre_append)
			tmp["prefix_list_out6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-PrefixListOut6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_as"
		if _, ok := i["remote-as"]; ok {
			v := flattenRouterBgpNeighborGroupRemoteAs(i["remote-as"], d, pre_append)
			tmp["remote_as"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RemoteAs")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_as_filter"
		if _, ok := i["remote-as-filter"]; ok {
			v := flattenRouterBgpNeighborGroupRemoteAsFilter(i["remote-as-filter"], d, pre_append)
			tmp["remote_as_filter"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RemoteAsFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as"
		if _, ok := i["remove-private-as"]; ok {
			v := flattenRouterBgpNeighborGroupRemovePrivateAs(i["remove-private-as"], d, pre_append)
			tmp["remove_private_as"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RemovePrivateAs")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as_evpn"
		if _, ok := i["remove-private-as-evpn"]; ok {
			v := flattenRouterBgpNeighborGroupRemovePrivateAsEvpn(i["remove-private-as-evpn"], d, pre_append)
			tmp["remove_private_as_evpn"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RemovePrivateAsEvpn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as_vpnv4"
		if _, ok := i["remove-private-as-vpnv4"]; ok {
			v := flattenRouterBgpNeighborGroupRemovePrivateAsVpnv4(i["remove-private-as-vpnv4"], d, pre_append)
			tmp["remove_private_as_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RemovePrivateAsVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as_vpnv6"
		if _, ok := i["remove-private-as-vpnv6"]; ok {
			v := flattenRouterBgpNeighborGroupRemovePrivateAsVpnv6(i["remove-private-as-vpnv6"], d, pre_append)
			tmp["remove_private_as_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RemovePrivateAsVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as6"
		if _, ok := i["remove-private-as6"]; ok {
			v := flattenRouterBgpNeighborGroupRemovePrivateAs6(i["remove-private-as6"], d, pre_append)
			tmp["remove_private_as6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RemovePrivateAs6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restart_time"
		if _, ok := i["restart-time"]; ok {
			v := flattenRouterBgpNeighborGroupRestartTime(i["restart-time"], d, pre_append)
			tmp["restart_time"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RestartTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retain_stale_time"
		if _, ok := i["retain-stale-time"]; ok {
			v := flattenRouterBgpNeighborGroupRetainStaleTime(i["retain-stale-time"], d, pre_append)
			tmp["retain_stale_time"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RetainStaleTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in"
		if _, ok := i["route-map-in"]; ok {
			v := flattenRouterBgpNeighborGroupRouteMapIn(i["route-map-in"], d, pre_append)
			tmp["route_map_in"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RouteMapIn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in_evpn"
		if _, ok := i["route-map-in-evpn"]; ok {
			v := flattenRouterBgpNeighborGroupRouteMapInEvpn(i["route-map-in-evpn"], d, pre_append)
			tmp["route_map_in_evpn"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RouteMapInEvpn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in_vpnv4"
		if _, ok := i["route-map-in-vpnv4"]; ok {
			v := flattenRouterBgpNeighborGroupRouteMapInVpnv4(i["route-map-in-vpnv4"], d, pre_append)
			tmp["route_map_in_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RouteMapInVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in_vpnv6"
		if _, ok := i["route-map-in-vpnv6"]; ok {
			v := flattenRouterBgpNeighborGroupRouteMapInVpnv6(i["route-map-in-vpnv6"], d, pre_append)
			tmp["route_map_in_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RouteMapInVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in6"
		if _, ok := i["route-map-in6"]; ok {
			v := flattenRouterBgpNeighborGroupRouteMapIn6(i["route-map-in6"], d, pre_append)
			tmp["route_map_in6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RouteMapIn6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out"
		if _, ok := i["route-map-out"]; ok {
			v := flattenRouterBgpNeighborGroupRouteMapOut(i["route-map-out"], d, pre_append)
			tmp["route_map_out"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RouteMapOut")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_evpn"
		if _, ok := i["route-map-out-evpn"]; ok {
			v := flattenRouterBgpNeighborGroupRouteMapOutEvpn(i["route-map-out-evpn"], d, pre_append)
			tmp["route_map_out_evpn"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RouteMapOutEvpn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_preferable"
		if _, ok := i["route-map-out-preferable"]; ok {
			v := flattenRouterBgpNeighborGroupRouteMapOutPreferable(i["route-map-out-preferable"], d, pre_append)
			tmp["route_map_out_preferable"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RouteMapOutPreferable")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv4"
		if _, ok := i["route-map-out-vpnv4"]; ok {
			v := flattenRouterBgpNeighborGroupRouteMapOutVpnv4(i["route-map-out-vpnv4"], d, pre_append)
			tmp["route_map_out_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RouteMapOutVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv4_preferable"
		if _, ok := i["route-map-out-vpnv4-preferable"]; ok {
			v := flattenRouterBgpNeighborGroupRouteMapOutVpnv4Preferable(i["route-map-out-vpnv4-preferable"], d, pre_append)
			tmp["route_map_out_vpnv4_preferable"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RouteMapOutVpnv4Preferable")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv6"
		if _, ok := i["route-map-out-vpnv6"]; ok {
			v := flattenRouterBgpNeighborGroupRouteMapOutVpnv6(i["route-map-out-vpnv6"], d, pre_append)
			tmp["route_map_out_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RouteMapOutVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv6_preferable"
		if _, ok := i["route-map-out-vpnv6-preferable"]; ok {
			v := flattenRouterBgpNeighborGroupRouteMapOutVpnv6Preferable(i["route-map-out-vpnv6-preferable"], d, pre_append)
			tmp["route_map_out_vpnv6_preferable"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RouteMapOutVpnv6Preferable")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out6"
		if _, ok := i["route-map-out6"]; ok {
			v := flattenRouterBgpNeighborGroupRouteMapOut6(i["route-map-out6"], d, pre_append)
			tmp["route_map_out6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RouteMapOut6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out6_preferable"
		if _, ok := i["route-map-out6-preferable"]; ok {
			v := flattenRouterBgpNeighborGroupRouteMapOut6Preferable(i["route-map-out6-preferable"], d, pre_append)
			tmp["route_map_out6_preferable"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RouteMapOut6Preferable")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client"
		if _, ok := i["route-reflector-client"]; ok {
			v := flattenRouterBgpNeighborGroupRouteReflectorClient(i["route-reflector-client"], d, pre_append)
			tmp["route_reflector_client"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RouteReflectorClient")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client_evpn"
		if _, ok := i["route-reflector-client-evpn"]; ok {
			v := flattenRouterBgpNeighborGroupRouteReflectorClientEvpn(i["route-reflector-client-evpn"], d, pre_append)
			tmp["route_reflector_client_evpn"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RouteReflectorClientEvpn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client_vpnv4"
		if _, ok := i["route-reflector-client-vpnv4"]; ok {
			v := flattenRouterBgpNeighborGroupRouteReflectorClientVpnv4(i["route-reflector-client-vpnv4"], d, pre_append)
			tmp["route_reflector_client_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RouteReflectorClientVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client_vpnv6"
		if _, ok := i["route-reflector-client-vpnv6"]; ok {
			v := flattenRouterBgpNeighborGroupRouteReflectorClientVpnv6(i["route-reflector-client-vpnv6"], d, pre_append)
			tmp["route_reflector_client_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RouteReflectorClientVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client6"
		if _, ok := i["route-reflector-client6"]; ok {
			v := flattenRouterBgpNeighborGroupRouteReflectorClient6(i["route-reflector-client6"], d, pre_append)
			tmp["route_reflector_client6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RouteReflectorClient6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client"
		if _, ok := i["route-server-client"]; ok {
			v := flattenRouterBgpNeighborGroupRouteServerClient(i["route-server-client"], d, pre_append)
			tmp["route_server_client"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RouteServerClient")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client_evpn"
		if _, ok := i["route-server-client-evpn"]; ok {
			v := flattenRouterBgpNeighborGroupRouteServerClientEvpn(i["route-server-client-evpn"], d, pre_append)
			tmp["route_server_client_evpn"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RouteServerClientEvpn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client_vpnv4"
		if _, ok := i["route-server-client-vpnv4"]; ok {
			v := flattenRouterBgpNeighborGroupRouteServerClientVpnv4(i["route-server-client-vpnv4"], d, pre_append)
			tmp["route_server_client_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RouteServerClientVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client_vpnv6"
		if _, ok := i["route-server-client-vpnv6"]; ok {
			v := flattenRouterBgpNeighborGroupRouteServerClientVpnv6(i["route-server-client-vpnv6"], d, pre_append)
			tmp["route_server_client_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RouteServerClientVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client6"
		if _, ok := i["route-server-client6"]; ok {
			v := flattenRouterBgpNeighborGroupRouteServerClient6(i["route-server-client6"], d, pre_append)
			tmp["route_server_client6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-RouteServerClient6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community"
		if _, ok := i["send-community"]; ok {
			v := flattenRouterBgpNeighborGroupSendCommunity(i["send-community"], d, pre_append)
			tmp["send_community"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-SendCommunity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community_evpn"
		if _, ok := i["send-community-evpn"]; ok {
			v := flattenRouterBgpNeighborGroupSendCommunityEvpn(i["send-community-evpn"], d, pre_append)
			tmp["send_community_evpn"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-SendCommunityEvpn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community_vpnv4"
		if _, ok := i["send-community-vpnv4"]; ok {
			v := flattenRouterBgpNeighborGroupSendCommunityVpnv4(i["send-community-vpnv4"], d, pre_append)
			tmp["send_community_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-SendCommunityVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community_vpnv6"
		if _, ok := i["send-community-vpnv6"]; ok {
			v := flattenRouterBgpNeighborGroupSendCommunityVpnv6(i["send-community-vpnv6"], d, pre_append)
			tmp["send_community_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-SendCommunityVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community6"
		if _, ok := i["send-community6"]; ok {
			v := flattenRouterBgpNeighborGroupSendCommunity6(i["send-community6"], d, pre_append)
			tmp["send_community6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-SendCommunity6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shutdown"
		if _, ok := i["shutdown"]; ok {
			v := flattenRouterBgpNeighborGroupShutdown(i["shutdown"], d, pre_append)
			tmp["shutdown"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-Shutdown")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration"
		if _, ok := i["soft-reconfiguration"]; ok {
			v := flattenRouterBgpNeighborGroupSoftReconfiguration(i["soft-reconfiguration"], d, pre_append)
			tmp["soft_reconfiguration"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-SoftReconfiguration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration_evpn"
		if _, ok := i["soft-reconfiguration-evpn"]; ok {
			v := flattenRouterBgpNeighborGroupSoftReconfigurationEvpn(i["soft-reconfiguration-evpn"], d, pre_append)
			tmp["soft_reconfiguration_evpn"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-SoftReconfigurationEvpn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration_vpnv4"
		if _, ok := i["soft-reconfiguration-vpnv4"]; ok {
			v := flattenRouterBgpNeighborGroupSoftReconfigurationVpnv4(i["soft-reconfiguration-vpnv4"], d, pre_append)
			tmp["soft_reconfiguration_vpnv4"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-SoftReconfigurationVpnv4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration_vpnv6"
		if _, ok := i["soft-reconfiguration-vpnv6"]; ok {
			v := flattenRouterBgpNeighborGroupSoftReconfigurationVpnv6(i["soft-reconfiguration-vpnv6"], d, pre_append)
			tmp["soft_reconfiguration_vpnv6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-SoftReconfigurationVpnv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration6"
		if _, ok := i["soft-reconfiguration6"]; ok {
			v := flattenRouterBgpNeighborGroupSoftReconfiguration6(i["soft-reconfiguration6"], d, pre_append)
			tmp["soft_reconfiguration6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-SoftReconfiguration6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stale_route"
		if _, ok := i["stale-route"]; ok {
			v := flattenRouterBgpNeighborGroupStaleRoute(i["stale-route"], d, pre_append)
			tmp["stale_route"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-StaleRoute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "strict_capability_match"
		if _, ok := i["strict-capability-match"]; ok {
			v := flattenRouterBgpNeighborGroupStrictCapabilityMatch(i["strict-capability-match"], d, pre_append)
			tmp["strict_capability_match"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-StrictCapabilityMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unsuppress_map"
		if _, ok := i["unsuppress-map"]; ok {
			v := flattenRouterBgpNeighborGroupUnsuppressMap(i["unsuppress-map"], d, pre_append)
			tmp["unsuppress_map"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-UnsuppressMap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unsuppress_map6"
		if _, ok := i["unsuppress-map6"]; ok {
			v := flattenRouterBgpNeighborGroupUnsuppressMap6(i["unsuppress-map6"], d, pre_append)
			tmp["unsuppress_map6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-UnsuppressMap6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_source"
		if _, ok := i["update-source"]; ok {
			v := flattenRouterBgpNeighborGroupUpdateSource(i["update-source"], d, pre_append)
			tmp["update_source"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-UpdateSource")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenRouterBgpNeighborGroupWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborGroup-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterBgpNeighborGroupActivate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupActivateEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupActivateVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupActivateVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupActivate6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAdditionalPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAdditionalPathVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAdditionalPathVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAdditionalPath6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAdvAdditionalPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAdvAdditionalPathVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAdvAdditionalPathVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAdvAdditionalPath6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAdvertisementInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAllowasIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAllowasInEnable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAllowasInEnableEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAllowasInEnableVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAllowasInEnableVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAllowasInEnable6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAllowasInEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAllowasInVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAllowasInVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAllowasIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAsOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAsOverride6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAttributeUnchanged(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupAttributeUnchangedVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupAttributeUnchangedVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupAttributeUnchanged6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupAuthOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityDefaultOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityDefaultOriginate6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityDynamic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityGracefulRestart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityGracefulRestartEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityGracefulRestartVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityGracefulRestartVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityGracefulRestart6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityOrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityOrf6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityRouteRefresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupConnectTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupDefaultOriginateRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupDefaultOriginateRoutemap6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupDistributeListIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupDistributeListInVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupDistributeListInVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupDistributeListIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupDistributeListOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupDistributeListOutVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupDistributeListOutVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupDistributeListOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupDontCapabilityNegotiate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupEbgpEnforceMultihop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupEbgpMultihopTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupFilterListIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupFilterListInVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupFilterListInVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupFilterListIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupFilterListOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupFilterListOutVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupFilterListOutVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupFilterListOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupHoldtimeTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupKeepAliveTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupLinkDownFailover(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupLocalAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupLocalAsNoPrepend(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupLocalAsReplaceAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixThresholdEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixThresholdVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixThresholdVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixThreshold6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixWarningOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixWarningOnlyEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixWarningOnlyVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixWarningOnlyVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixWarningOnly6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupNextHopSelf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupNextHopSelfRr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupNextHopSelfRr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupNextHopSelfVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupNextHopSelfVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupNextHopSelf6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupOverrideCapability(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupPassive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupPrefixListIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupPrefixListInVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupPrefixListInVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupPrefixListIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupPrefixListOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupPrefixListOutVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupPrefixListOutVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupPrefixListOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRemoteAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRemoteAsFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRemovePrivateAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRemovePrivateAsEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRemovePrivateAsVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRemovePrivateAsVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRemovePrivateAs6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRestartTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRetainStaleTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteMapIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRouteMapInEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRouteMapInVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRouteMapInVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRouteMapIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRouteMapOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRouteMapOutEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRouteMapOutPreferable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRouteMapOutVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRouteMapOutVpnv4Preferable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRouteMapOutVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRouteMapOutVpnv6Preferable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRouteMapOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRouteMapOut6Preferable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRouteReflectorClient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteReflectorClientEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteReflectorClientVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteReflectorClientVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteReflectorClient6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteServerClient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteServerClientEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteServerClientVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteServerClientVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteServerClient6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSendCommunity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSendCommunityEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSendCommunityVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSendCommunityVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSendCommunity6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupShutdown(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSoftReconfiguration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSoftReconfigurationEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSoftReconfigurationVpnv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSoftReconfigurationVpnv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSoftReconfiguration6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupStaleRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupStrictCapabilityMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupUnsuppressMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupUnsuppressMap6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupUpdateSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterBgpNeighborRangeId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_neighbor_num"
		if _, ok := i["max-neighbor-num"]; ok {
			v := flattenRouterBgpNeighborRangeMaxNeighborNum(i["max-neighbor-num"], d, pre_append)
			tmp["max_neighbor_num"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborRange-MaxNeighborNum")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbor_group"
		if _, ok := i["neighbor-group"]; ok {
			v := flattenRouterBgpNeighborRangeNeighborGroup(i["neighbor-group"], d, pre_append)
			tmp["neighbor_group"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborRange-NeighborGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			v := flattenRouterBgpNeighborRangePrefix(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborRange-Prefix")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterBgpNeighborRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRangeMaxNeighborNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRangeNeighborGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRangePrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRange6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterBgpNeighborRange6Id(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborRange6-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_neighbor_num"
		if _, ok := i["max-neighbor-num"]; ok {
			v := flattenRouterBgpNeighborRange6MaxNeighborNum(i["max-neighbor-num"], d, pre_append)
			tmp["max_neighbor_num"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborRange6-MaxNeighborNum")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbor_group"
		if _, ok := i["neighbor-group"]; ok {
			v := flattenRouterBgpNeighborRange6NeighborGroup(i["neighbor-group"], d, pre_append)
			tmp["neighbor_group"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborRange6-NeighborGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {
			v := flattenRouterBgpNeighborRange6Prefix6(i["prefix6"], d, pre_append)
			tmp["prefix6"] = fortiAPISubPartPatch(v, "RouterBgp-NeighborRange6-Prefix6")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterBgpNeighborRange6Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRange6MaxNeighborNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRange6NeighborGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRange6Prefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNetwork(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "backdoor"
		if _, ok := i["backdoor"]; ok {
			v := flattenRouterBgpNetworkBackdoor(i["backdoor"], d, pre_append)
			tmp["backdoor"] = fortiAPISubPartPatch(v, "RouterBgp-Network-Backdoor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterBgpNetworkId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterBgp-Network-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "network_import_check"
		if _, ok := i["network-import-check"]; ok {
			v := flattenRouterBgpNetworkNetworkImportCheck(i["network-import-check"], d, pre_append)
			tmp["network_import_check"] = fortiAPISubPartPatch(v, "RouterBgp-Network-NetworkImportCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			v := flattenRouterBgpNetworkPrefix(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "RouterBgp-Network-Prefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := i["route-map"]; ok {
			v := flattenRouterBgpNetworkRouteMap(i["route-map"], d, pre_append)
			tmp["route_map"] = fortiAPISubPartPatch(v, "RouterBgp-Network-RouteMap")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterBgpNetworkBackdoor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNetworkId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNetworkNetworkImportCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNetworkPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNetworkRouteMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNetworkImportCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNetwork6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "backdoor"
		if _, ok := i["backdoor"]; ok {
			v := flattenRouterBgpNetwork6Backdoor(i["backdoor"], d, pre_append)
			tmp["backdoor"] = fortiAPISubPartPatch(v, "RouterBgp-Network6-Backdoor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterBgpNetwork6Id(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterBgp-Network6-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "network_import_check"
		if _, ok := i["network-import-check"]; ok {
			v := flattenRouterBgpNetwork6NetworkImportCheck(i["network-import-check"], d, pre_append)
			tmp["network_import_check"] = fortiAPISubPartPatch(v, "RouterBgp-Network6-NetworkImportCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {
			v := flattenRouterBgpNetwork6Prefix6(i["prefix6"], d, pre_append)
			tmp["prefix6"] = fortiAPISubPartPatch(v, "RouterBgp-Network6-Prefix6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := i["route-map"]; ok {
			v := flattenRouterBgpNetwork6RouteMap(i["route-map"], d, pre_append)
			tmp["route_map"] = fortiAPISubPartPatch(v, "RouterBgp-Network6-RouteMap")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterBgpNetwork6Backdoor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNetwork6Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNetwork6NetworkImportCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNetwork6Prefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNetwork6RouteMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpRecursiveInheritPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpRecursiveNextHop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpRedistribute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterBgpRedistributeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "RouterBgp-Redistribute-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := i["route-map"]; ok {
			v := flattenRouterBgpRedistributeRouteMap(i["route-map"], d, pre_append)
			tmp["route_map"] = fortiAPISubPartPatch(v, "RouterBgp-Redistribute-RouteMap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenRouterBgpRedistributeStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "RouterBgp-Redistribute-Status")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterBgpRedistributeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpRedistributeRouteMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpRedistributeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpRedistribute6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterBgpRedistribute6Name(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "RouterBgp-Redistribute6-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := i["route-map"]; ok {
			v := flattenRouterBgpRedistribute6RouteMap(i["route-map"], d, pre_append)
			tmp["route_map"] = fortiAPISubPartPatch(v, "RouterBgp-Redistribute6-RouteMap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenRouterBgpRedistribute6Status(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "RouterBgp-Redistribute6-Status")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterBgpRedistribute6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpRedistribute6RouteMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpRedistribute6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpRouterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpScanTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpSynchronization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpTagResolveMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpVrfLeak(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := i["target"]; ok {
			v := flattenRouterBgpVrfLeakTarget(i["target"], d, pre_append)
			tmp["target"] = fortiAPISubPartPatch(v, "RouterBgp-VrfLeak-Target")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := i["vrf"]; ok {
			v := flattenRouterBgpVrfLeakVrf(i["vrf"], d, pre_append)
			tmp["vrf"] = fortiAPISubPartPatch(v, "RouterBgp-VrfLeak-Vrf")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterBgpVrfLeakTarget(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterBgpVrfLeakTargetInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "RouterBgpVrfLeak-Target-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := i["route-map"]; ok {
			v := flattenRouterBgpVrfLeakTargetRouteMap(i["route-map"], d, pre_append)
			tmp["route_map"] = fortiAPISubPartPatch(v, "RouterBgpVrfLeak-Target-RouteMap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := i["vrf"]; ok {
			v := flattenRouterBgpVrfLeakTargetVrf(i["vrf"], d, pre_append)
			tmp["vrf"] = fortiAPISubPartPatch(v, "RouterBgpVrfLeak-Target-Vrf")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterBgpVrfLeakTargetInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpVrfLeakTargetRouteMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpVrfLeakTargetVrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpVrfLeakVrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpVrfLeak6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := i["target"]; ok {
			v := flattenRouterBgpVrfLeak6Target(i["target"], d, pre_append)
			tmp["target"] = fortiAPISubPartPatch(v, "RouterBgp-VrfLeak6-Target")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := i["vrf"]; ok {
			v := flattenRouterBgpVrfLeak6Vrf(i["vrf"], d, pre_append)
			tmp["vrf"] = fortiAPISubPartPatch(v, "RouterBgp-VrfLeak6-Vrf")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterBgpVrfLeak6Target(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterBgpVrfLeak6TargetInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "RouterBgpVrfLeak6-Target-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := i["route-map"]; ok {
			v := flattenRouterBgpVrfLeak6TargetRouteMap(i["route-map"], d, pre_append)
			tmp["route_map"] = fortiAPISubPartPatch(v, "RouterBgpVrfLeak6-Target-RouteMap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := i["vrf"]; ok {
			v := flattenRouterBgpVrfLeak6TargetVrf(i["vrf"], d, pre_append)
			tmp["vrf"] = fortiAPISubPartPatch(v, "RouterBgpVrfLeak6-Target-Vrf")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterBgpVrfLeak6TargetInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpVrfLeak6TargetRouteMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpVrfLeak6TargetVrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpVrfLeak6Vrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpVrf(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "export_rt"
		if _, ok := i["export-rt"]; ok {
			v := flattenRouterBgpVrfExportRt(i["export-rt"], d, pre_append)
			tmp["export_rt"] = fortiAPISubPartPatch(v, "RouterBgp-Vrf-ExportRt")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "import_route_map"
		if _, ok := i["import-route-map"]; ok {
			v := flattenRouterBgpVrfImportRouteMap(i["import-route-map"], d, pre_append)
			tmp["import_route_map"] = fortiAPISubPartPatch(v, "RouterBgp-Vrf-ImportRouteMap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "import_rt"
		if _, ok := i["import-rt"]; ok {
			v := flattenRouterBgpVrfImportRt(i["import-rt"], d, pre_append)
			tmp["import_rt"] = fortiAPISubPartPatch(v, "RouterBgp-Vrf-ImportRt")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "leak_target"
		if _, ok := i["leak-target"]; ok {
			v := flattenRouterBgpVrfLeakTargetU(i["leak-target"], d, pre_append)
			tmp["leak_target"] = fortiAPISubPartPatch(v, "RouterBgp-Vrf-LeakTarget")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rd"
		if _, ok := i["rd"]; ok {
			v := flattenRouterBgpVrfRd(i["rd"], d, pre_append)
			tmp["rd"] = fortiAPISubPartPatch(v, "RouterBgp-Vrf-Rd")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := i["role"]; ok {
			v := flattenRouterBgpVrfRole(i["role"], d, pre_append)
			tmp["role"] = fortiAPISubPartPatch(v, "RouterBgp-Vrf-Role")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := i["vrf"]; ok {
			v := flattenRouterBgpVrfVrf(i["vrf"], d, pre_append)
			tmp["vrf"] = fortiAPISubPartPatch(v, "RouterBgp-Vrf-Vrf")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterBgpVrfExportRt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpVrfImportRouteMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpVrfImportRt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpVrfLeakTargetU(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterBgpVrfLeakTargetUInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "RouterBgpVrf-LeakTarget-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := i["route-map"]; ok {
			v := flattenRouterBgpVrfLeakTargetURouteMap(i["route-map"], d, pre_append)
			tmp["route_map"] = fortiAPISubPartPatch(v, "RouterBgpVrf-LeakTarget-RouteMap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := i["vrf"]; ok {
			v := flattenRouterBgpVrfLeakTargetUVrf(i["vrf"], d, pre_append)
			tmp["vrf"] = fortiAPISubPartPatch(v, "RouterBgpVrf-LeakTarget-Vrf")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterBgpVrfLeakTargetUInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpVrfLeakTargetURouteMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpVrfLeakTargetUVrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpVrfRd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpVrfRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpVrfVrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpVrf6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "export_rt"
		if _, ok := i["export-rt"]; ok {
			v := flattenRouterBgpVrf6ExportRt(i["export-rt"], d, pre_append)
			tmp["export_rt"] = fortiAPISubPartPatch(v, "RouterBgp-Vrf6-ExportRt")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "import_route_map"
		if _, ok := i["import-route-map"]; ok {
			v := flattenRouterBgpVrf6ImportRouteMap(i["import-route-map"], d, pre_append)
			tmp["import_route_map"] = fortiAPISubPartPatch(v, "RouterBgp-Vrf6-ImportRouteMap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "import_rt"
		if _, ok := i["import-rt"]; ok {
			v := flattenRouterBgpVrf6ImportRt(i["import-rt"], d, pre_append)
			tmp["import_rt"] = fortiAPISubPartPatch(v, "RouterBgp-Vrf6-ImportRt")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "leak_target"
		if _, ok := i["leak-target"]; ok {
			v := flattenRouterBgpVrf6LeakTarget(i["leak-target"], d, pre_append)
			tmp["leak_target"] = fortiAPISubPartPatch(v, "RouterBgp-Vrf6-LeakTarget")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rd"
		if _, ok := i["rd"]; ok {
			v := flattenRouterBgpVrf6Rd(i["rd"], d, pre_append)
			tmp["rd"] = fortiAPISubPartPatch(v, "RouterBgp-Vrf6-Rd")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := i["role"]; ok {
			v := flattenRouterBgpVrf6Role(i["role"], d, pre_append)
			tmp["role"] = fortiAPISubPartPatch(v, "RouterBgp-Vrf6-Role")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := i["vrf"]; ok {
			v := flattenRouterBgpVrf6Vrf(i["vrf"], d, pre_append)
			tmp["vrf"] = fortiAPISubPartPatch(v, "RouterBgp-Vrf6-Vrf")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterBgpVrf6ExportRt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpVrf6ImportRouteMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpVrf6ImportRt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpVrf6LeakTarget(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterBgpVrf6LeakTargetInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "RouterBgpVrf6-LeakTarget-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := i["route-map"]; ok {
			v := flattenRouterBgpVrf6LeakTargetRouteMap(i["route-map"], d, pre_append)
			tmp["route_map"] = fortiAPISubPartPatch(v, "RouterBgpVrf6-LeakTarget-RouteMap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := i["vrf"]; ok {
			v := flattenRouterBgpVrf6LeakTargetVrf(i["vrf"], d, pre_append)
			tmp["vrf"] = fortiAPISubPartPatch(v, "RouterBgpVrf6-LeakTarget-Vrf")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterBgpVrf6LeakTargetInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpVrf6LeakTargetRouteMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpVrf6LeakTargetVrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpVrf6Rd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpVrf6Role(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpVrf6Vrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterBgp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("additional_path", flattenRouterBgpAdditionalPath(o["additional-path"], d, "additional_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["additional-path"], "RouterBgp-AdditionalPath"); ok {
			if err = d.Set("additional_path", vv); err != nil {
				return fmt.Errorf("Error reading additional_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading additional_path: %v", err)
		}
	}

	if err = d.Set("additional_path_select", flattenRouterBgpAdditionalPathSelect(o["additional-path-select"], d, "additional_path_select")); err != nil {
		if vv, ok := fortiAPIPatch(o["additional-path-select"], "RouterBgp-AdditionalPathSelect"); ok {
			if err = d.Set("additional_path_select", vv); err != nil {
				return fmt.Errorf("Error reading additional_path_select: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading additional_path_select: %v", err)
		}
	}

	if err = d.Set("additional_path_select_vpnv4", flattenRouterBgpAdditionalPathSelectVpnv4(o["additional-path-select-vpnv4"], d, "additional_path_select_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["additional-path-select-vpnv4"], "RouterBgp-AdditionalPathSelectVpnv4"); ok {
			if err = d.Set("additional_path_select_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading additional_path_select_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading additional_path_select_vpnv4: %v", err)
		}
	}

	if err = d.Set("additional_path_select_vpnv6", flattenRouterBgpAdditionalPathSelectVpnv6(o["additional-path-select-vpnv6"], d, "additional_path_select_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["additional-path-select-vpnv6"], "RouterBgp-AdditionalPathSelectVpnv6"); ok {
			if err = d.Set("additional_path_select_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading additional_path_select_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading additional_path_select_vpnv6: %v", err)
		}
	}

	if err = d.Set("additional_path_select6", flattenRouterBgpAdditionalPathSelect6(o["additional-path-select6"], d, "additional_path_select6")); err != nil {
		if vv, ok := fortiAPIPatch(o["additional-path-select6"], "RouterBgp-AdditionalPathSelect6"); ok {
			if err = d.Set("additional_path_select6", vv); err != nil {
				return fmt.Errorf("Error reading additional_path_select6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading additional_path_select6: %v", err)
		}
	}

	if err = d.Set("additional_path_vpnv4", flattenRouterBgpAdditionalPathVpnv4(o["additional-path-vpnv4"], d, "additional_path_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["additional-path-vpnv4"], "RouterBgp-AdditionalPathVpnv4"); ok {
			if err = d.Set("additional_path_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading additional_path_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading additional_path_vpnv4: %v", err)
		}
	}

	if err = d.Set("additional_path_vpnv6", flattenRouterBgpAdditionalPathVpnv6(o["additional-path-vpnv6"], d, "additional_path_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["additional-path-vpnv6"], "RouterBgp-AdditionalPathVpnv6"); ok {
			if err = d.Set("additional_path_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading additional_path_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading additional_path_vpnv6: %v", err)
		}
	}

	if err = d.Set("additional_path6", flattenRouterBgpAdditionalPath6(o["additional-path6"], d, "additional_path6")); err != nil {
		if vv, ok := fortiAPIPatch(o["additional-path6"], "RouterBgp-AdditionalPath6"); ok {
			if err = d.Set("additional_path6", vv); err != nil {
				return fmt.Errorf("Error reading additional_path6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading additional_path6: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("admin_distance", flattenRouterBgpAdminDistance(o["admin-distance"], d, "admin_distance")); err != nil {
			if vv, ok := fortiAPIPatch(o["admin-distance"], "RouterBgp-AdminDistance"); ok {
				if err = d.Set("admin_distance", vv); err != nil {
					return fmt.Errorf("Error reading admin_distance: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading admin_distance: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("admin_distance"); ok {
			if err = d.Set("admin_distance", flattenRouterBgpAdminDistance(o["admin-distance"], d, "admin_distance")); err != nil {
				if vv, ok := fortiAPIPatch(o["admin-distance"], "RouterBgp-AdminDistance"); ok {
					if err = d.Set("admin_distance", vv); err != nil {
						return fmt.Errorf("Error reading admin_distance: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading admin_distance: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("aggregate_address", flattenRouterBgpAggregateAddress(o["aggregate-address"], d, "aggregate_address")); err != nil {
			if vv, ok := fortiAPIPatch(o["aggregate-address"], "RouterBgp-AggregateAddress"); ok {
				if err = d.Set("aggregate_address", vv); err != nil {
					return fmt.Errorf("Error reading aggregate_address: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading aggregate_address: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("aggregate_address"); ok {
			if err = d.Set("aggregate_address", flattenRouterBgpAggregateAddress(o["aggregate-address"], d, "aggregate_address")); err != nil {
				if vv, ok := fortiAPIPatch(o["aggregate-address"], "RouterBgp-AggregateAddress"); ok {
					if err = d.Set("aggregate_address", vv); err != nil {
						return fmt.Errorf("Error reading aggregate_address: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading aggregate_address: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("aggregate_address6", flattenRouterBgpAggregateAddress6(o["aggregate-address6"], d, "aggregate_address6")); err != nil {
			if vv, ok := fortiAPIPatch(o["aggregate-address6"], "RouterBgp-AggregateAddress6"); ok {
				if err = d.Set("aggregate_address6", vv); err != nil {
					return fmt.Errorf("Error reading aggregate_address6: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading aggregate_address6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("aggregate_address6"); ok {
			if err = d.Set("aggregate_address6", flattenRouterBgpAggregateAddress6(o["aggregate-address6"], d, "aggregate_address6")); err != nil {
				if vv, ok := fortiAPIPatch(o["aggregate-address6"], "RouterBgp-AggregateAddress6"); ok {
					if err = d.Set("aggregate_address6", vv); err != nil {
						return fmt.Errorf("Error reading aggregate_address6: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading aggregate_address6: %v", err)
				}
			}
		}
	}

	if err = d.Set("always_compare_med", flattenRouterBgpAlwaysCompareMed(o["always-compare-med"], d, "always_compare_med")); err != nil {
		if vv, ok := fortiAPIPatch(o["always-compare-med"], "RouterBgp-AlwaysCompareMed"); ok {
			if err = d.Set("always_compare_med", vv); err != nil {
				return fmt.Errorf("Error reading always_compare_med: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading always_compare_med: %v", err)
		}
	}

	if err = d.Set("as", flattenRouterBgpAs(o["as"], d, "as")); err != nil {
		if vv, ok := fortiAPIPatch(o["as"], "RouterBgp-As"); ok {
			if err = d.Set("as", vv); err != nil {
				return fmt.Errorf("Error reading as: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading as: %v", err)
		}
	}

	if err = d.Set("bestpath_as_path_ignore", flattenRouterBgpBestpathAsPathIgnore(o["bestpath-as-path-ignore"], d, "bestpath_as_path_ignore")); err != nil {
		if vv, ok := fortiAPIPatch(o["bestpath-as-path-ignore"], "RouterBgp-BestpathAsPathIgnore"); ok {
			if err = d.Set("bestpath_as_path_ignore", vv); err != nil {
				return fmt.Errorf("Error reading bestpath_as_path_ignore: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bestpath_as_path_ignore: %v", err)
		}
	}

	if err = d.Set("bestpath_cmp_confed_aspath", flattenRouterBgpBestpathCmpConfedAspath(o["bestpath-cmp-confed-aspath"], d, "bestpath_cmp_confed_aspath")); err != nil {
		if vv, ok := fortiAPIPatch(o["bestpath-cmp-confed-aspath"], "RouterBgp-BestpathCmpConfedAspath"); ok {
			if err = d.Set("bestpath_cmp_confed_aspath", vv); err != nil {
				return fmt.Errorf("Error reading bestpath_cmp_confed_aspath: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bestpath_cmp_confed_aspath: %v", err)
		}
	}

	if err = d.Set("bestpath_cmp_routerid", flattenRouterBgpBestpathCmpRouterid(o["bestpath-cmp-routerid"], d, "bestpath_cmp_routerid")); err != nil {
		if vv, ok := fortiAPIPatch(o["bestpath-cmp-routerid"], "RouterBgp-BestpathCmpRouterid"); ok {
			if err = d.Set("bestpath_cmp_routerid", vv); err != nil {
				return fmt.Errorf("Error reading bestpath_cmp_routerid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bestpath_cmp_routerid: %v", err)
		}
	}

	if err = d.Set("bestpath_med_confed", flattenRouterBgpBestpathMedConfed(o["bestpath-med-confed"], d, "bestpath_med_confed")); err != nil {
		if vv, ok := fortiAPIPatch(o["bestpath-med-confed"], "RouterBgp-BestpathMedConfed"); ok {
			if err = d.Set("bestpath_med_confed", vv); err != nil {
				return fmt.Errorf("Error reading bestpath_med_confed: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bestpath_med_confed: %v", err)
		}
	}

	if err = d.Set("bestpath_med_missing_as_worst", flattenRouterBgpBestpathMedMissingAsWorst(o["bestpath-med-missing-as-worst"], d, "bestpath_med_missing_as_worst")); err != nil {
		if vv, ok := fortiAPIPatch(o["bestpath-med-missing-as-worst"], "RouterBgp-BestpathMedMissingAsWorst"); ok {
			if err = d.Set("bestpath_med_missing_as_worst", vv); err != nil {
				return fmt.Errorf("Error reading bestpath_med_missing_as_worst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bestpath_med_missing_as_worst: %v", err)
		}
	}

	if err = d.Set("client_to_client_reflection", flattenRouterBgpClientToClientReflection(o["client-to-client-reflection"], d, "client_to_client_reflection")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-to-client-reflection"], "RouterBgp-ClientToClientReflection"); ok {
			if err = d.Set("client_to_client_reflection", vv); err != nil {
				return fmt.Errorf("Error reading client_to_client_reflection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_to_client_reflection: %v", err)
		}
	}

	if err = d.Set("cluster_id", flattenRouterBgpClusterId(o["cluster-id"], d, "cluster_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["cluster-id"], "RouterBgp-ClusterId"); ok {
			if err = d.Set("cluster_id", vv); err != nil {
				return fmt.Errorf("Error reading cluster_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cluster_id: %v", err)
		}
	}

	if err = d.Set("confederation_identifier", flattenRouterBgpConfederationIdentifier(o["confederation-identifier"], d, "confederation_identifier")); err != nil {
		if vv, ok := fortiAPIPatch(o["confederation-identifier"], "RouterBgp-ConfederationIdentifier"); ok {
			if err = d.Set("confederation_identifier", vv); err != nil {
				return fmt.Errorf("Error reading confederation_identifier: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading confederation_identifier: %v", err)
		}
	}

	if err = d.Set("confederation_peers", flattenRouterBgpConfederationPeers(o["confederation-peers"], d, "confederation_peers")); err != nil {
		if vv, ok := fortiAPIPatch(o["confederation-peers"], "RouterBgp-ConfederationPeers"); ok {
			if err = d.Set("confederation_peers", vv); err != nil {
				return fmt.Errorf("Error reading confederation_peers: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading confederation_peers: %v", err)
		}
	}

	if err = d.Set("cross_family_conditional_adv", flattenRouterBgpCrossFamilyConditionalAdv(o["cross-family-conditional-adv"], d, "cross_family_conditional_adv")); err != nil {
		if vv, ok := fortiAPIPatch(o["cross-family-conditional-adv"], "RouterBgp-CrossFamilyConditionalAdv"); ok {
			if err = d.Set("cross_family_conditional_adv", vv); err != nil {
				return fmt.Errorf("Error reading cross_family_conditional_adv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cross_family_conditional_adv: %v", err)
		}
	}

	if err = d.Set("dampening", flattenRouterBgpDampening(o["dampening"], d, "dampening")); err != nil {
		if vv, ok := fortiAPIPatch(o["dampening"], "RouterBgp-Dampening"); ok {
			if err = d.Set("dampening", vv); err != nil {
				return fmt.Errorf("Error reading dampening: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dampening: %v", err)
		}
	}

	if err = d.Set("dampening_max_suppress_time", flattenRouterBgpDampeningMaxSuppressTime(o["dampening-max-suppress-time"], d, "dampening_max_suppress_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["dampening-max-suppress-time"], "RouterBgp-DampeningMaxSuppressTime"); ok {
			if err = d.Set("dampening_max_suppress_time", vv); err != nil {
				return fmt.Errorf("Error reading dampening_max_suppress_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dampening_max_suppress_time: %v", err)
		}
	}

	if err = d.Set("dampening_reachability_half_life", flattenRouterBgpDampeningReachabilityHalfLife(o["dampening-reachability-half-life"], d, "dampening_reachability_half_life")); err != nil {
		if vv, ok := fortiAPIPatch(o["dampening-reachability-half-life"], "RouterBgp-DampeningReachabilityHalfLife"); ok {
			if err = d.Set("dampening_reachability_half_life", vv); err != nil {
				return fmt.Errorf("Error reading dampening_reachability_half_life: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dampening_reachability_half_life: %v", err)
		}
	}

	if err = d.Set("dampening_reuse", flattenRouterBgpDampeningReuse(o["dampening-reuse"], d, "dampening_reuse")); err != nil {
		if vv, ok := fortiAPIPatch(o["dampening-reuse"], "RouterBgp-DampeningReuse"); ok {
			if err = d.Set("dampening_reuse", vv); err != nil {
				return fmt.Errorf("Error reading dampening_reuse: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dampening_reuse: %v", err)
		}
	}

	if err = d.Set("dampening_route_map", flattenRouterBgpDampeningRouteMap(o["dampening-route-map"], d, "dampening_route_map")); err != nil {
		if vv, ok := fortiAPIPatch(o["dampening-route-map"], "RouterBgp-DampeningRouteMap"); ok {
			if err = d.Set("dampening_route_map", vv); err != nil {
				return fmt.Errorf("Error reading dampening_route_map: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dampening_route_map: %v", err)
		}
	}

	if err = d.Set("dampening_suppress", flattenRouterBgpDampeningSuppress(o["dampening-suppress"], d, "dampening_suppress")); err != nil {
		if vv, ok := fortiAPIPatch(o["dampening-suppress"], "RouterBgp-DampeningSuppress"); ok {
			if err = d.Set("dampening_suppress", vv); err != nil {
				return fmt.Errorf("Error reading dampening_suppress: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dampening_suppress: %v", err)
		}
	}

	if err = d.Set("dampening_unreachability_half_life", flattenRouterBgpDampeningUnreachabilityHalfLife(o["dampening-unreachability-half-life"], d, "dampening_unreachability_half_life")); err != nil {
		if vv, ok := fortiAPIPatch(o["dampening-unreachability-half-life"], "RouterBgp-DampeningUnreachabilityHalfLife"); ok {
			if err = d.Set("dampening_unreachability_half_life", vv); err != nil {
				return fmt.Errorf("Error reading dampening_unreachability_half_life: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dampening_unreachability_half_life: %v", err)
		}
	}

	if err = d.Set("default_local_preference", flattenRouterBgpDefaultLocalPreference(o["default-local-preference"], d, "default_local_preference")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-local-preference"], "RouterBgp-DefaultLocalPreference"); ok {
			if err = d.Set("default_local_preference", vv); err != nil {
				return fmt.Errorf("Error reading default_local_preference: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_local_preference: %v", err)
		}
	}

	if err = d.Set("deterministic_med", flattenRouterBgpDeterministicMed(o["deterministic-med"], d, "deterministic_med")); err != nil {
		if vv, ok := fortiAPIPatch(o["deterministic-med"], "RouterBgp-DeterministicMed"); ok {
			if err = d.Set("deterministic_med", vv); err != nil {
				return fmt.Errorf("Error reading deterministic_med: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading deterministic_med: %v", err)
		}
	}

	if err = d.Set("distance_external", flattenRouterBgpDistanceExternal(o["distance-external"], d, "distance_external")); err != nil {
		if vv, ok := fortiAPIPatch(o["distance-external"], "RouterBgp-DistanceExternal"); ok {
			if err = d.Set("distance_external", vv); err != nil {
				return fmt.Errorf("Error reading distance_external: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distance_external: %v", err)
		}
	}

	if err = d.Set("distance_internal", flattenRouterBgpDistanceInternal(o["distance-internal"], d, "distance_internal")); err != nil {
		if vv, ok := fortiAPIPatch(o["distance-internal"], "RouterBgp-DistanceInternal"); ok {
			if err = d.Set("distance_internal", vv); err != nil {
				return fmt.Errorf("Error reading distance_internal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distance_internal: %v", err)
		}
	}

	if err = d.Set("distance_local", flattenRouterBgpDistanceLocal(o["distance-local"], d, "distance_local")); err != nil {
		if vv, ok := fortiAPIPatch(o["distance-local"], "RouterBgp-DistanceLocal"); ok {
			if err = d.Set("distance_local", vv); err != nil {
				return fmt.Errorf("Error reading distance_local: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distance_local: %v", err)
		}
	}

	if err = d.Set("ebgp_multipath", flattenRouterBgpEbgpMultipath(o["ebgp-multipath"], d, "ebgp_multipath")); err != nil {
		if vv, ok := fortiAPIPatch(o["ebgp-multipath"], "RouterBgp-EbgpMultipath"); ok {
			if err = d.Set("ebgp_multipath", vv); err != nil {
				return fmt.Errorf("Error reading ebgp_multipath: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ebgp_multipath: %v", err)
		}
	}

	if err = d.Set("enforce_first_as", flattenRouterBgpEnforceFirstAs(o["enforce-first-as"], d, "enforce_first_as")); err != nil {
		if vv, ok := fortiAPIPatch(o["enforce-first-as"], "RouterBgp-EnforceFirstAs"); ok {
			if err = d.Set("enforce_first_as", vv); err != nil {
				return fmt.Errorf("Error reading enforce_first_as: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enforce_first_as: %v", err)
		}
	}

	if err = d.Set("fast_external_failover", flattenRouterBgpFastExternalFailover(o["fast-external-failover"], d, "fast_external_failover")); err != nil {
		if vv, ok := fortiAPIPatch(o["fast-external-failover"], "RouterBgp-FastExternalFailover"); ok {
			if err = d.Set("fast_external_failover", vv); err != nil {
				return fmt.Errorf("Error reading fast_external_failover: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fast_external_failover: %v", err)
		}
	}

	if err = d.Set("graceful_end_on_timer", flattenRouterBgpGracefulEndOnTimer(o["graceful-end-on-timer"], d, "graceful_end_on_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["graceful-end-on-timer"], "RouterBgp-GracefulEndOnTimer"); ok {
			if err = d.Set("graceful_end_on_timer", vv); err != nil {
				return fmt.Errorf("Error reading graceful_end_on_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading graceful_end_on_timer: %v", err)
		}
	}

	if err = d.Set("graceful_restart", flattenRouterBgpGracefulRestart(o["graceful-restart"], d, "graceful_restart")); err != nil {
		if vv, ok := fortiAPIPatch(o["graceful-restart"], "RouterBgp-GracefulRestart"); ok {
			if err = d.Set("graceful_restart", vv); err != nil {
				return fmt.Errorf("Error reading graceful_restart: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading graceful_restart: %v", err)
		}
	}

	if err = d.Set("graceful_restart_time", flattenRouterBgpGracefulRestartTime(o["graceful-restart-time"], d, "graceful_restart_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["graceful-restart-time"], "RouterBgp-GracefulRestartTime"); ok {
			if err = d.Set("graceful_restart_time", vv); err != nil {
				return fmt.Errorf("Error reading graceful_restart_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading graceful_restart_time: %v", err)
		}
	}

	if err = d.Set("graceful_stalepath_time", flattenRouterBgpGracefulStalepathTime(o["graceful-stalepath-time"], d, "graceful_stalepath_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["graceful-stalepath-time"], "RouterBgp-GracefulStalepathTime"); ok {
			if err = d.Set("graceful_stalepath_time", vv); err != nil {
				return fmt.Errorf("Error reading graceful_stalepath_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading graceful_stalepath_time: %v", err)
		}
	}

	if err = d.Set("graceful_update_delay", flattenRouterBgpGracefulUpdateDelay(o["graceful-update-delay"], d, "graceful_update_delay")); err != nil {
		if vv, ok := fortiAPIPatch(o["graceful-update-delay"], "RouterBgp-GracefulUpdateDelay"); ok {
			if err = d.Set("graceful_update_delay", vv); err != nil {
				return fmt.Errorf("Error reading graceful_update_delay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading graceful_update_delay: %v", err)
		}
	}

	if err = d.Set("holdtime_timer", flattenRouterBgpHoldtimeTimer(o["holdtime-timer"], d, "holdtime_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["holdtime-timer"], "RouterBgp-HoldtimeTimer"); ok {
			if err = d.Set("holdtime_timer", vv); err != nil {
				return fmt.Errorf("Error reading holdtime_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading holdtime_timer: %v", err)
		}
	}

	if err = d.Set("ibgp_multipath", flattenRouterBgpIbgpMultipath(o["ibgp-multipath"], d, "ibgp_multipath")); err != nil {
		if vv, ok := fortiAPIPatch(o["ibgp-multipath"], "RouterBgp-IbgpMultipath"); ok {
			if err = d.Set("ibgp_multipath", vv); err != nil {
				return fmt.Errorf("Error reading ibgp_multipath: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ibgp_multipath: %v", err)
		}
	}

	if err = d.Set("ignore_optional_capability", flattenRouterBgpIgnoreOptionalCapability(o["ignore-optional-capability"], d, "ignore_optional_capability")); err != nil {
		if vv, ok := fortiAPIPatch(o["ignore-optional-capability"], "RouterBgp-IgnoreOptionalCapability"); ok {
			if err = d.Set("ignore_optional_capability", vv); err != nil {
				return fmt.Errorf("Error reading ignore_optional_capability: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ignore_optional_capability: %v", err)
		}
	}

	if err = d.Set("keepalive_timer", flattenRouterBgpKeepaliveTimer(o["keepalive-timer"], d, "keepalive_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["keepalive-timer"], "RouterBgp-KeepaliveTimer"); ok {
			if err = d.Set("keepalive_timer", vv); err != nil {
				return fmt.Errorf("Error reading keepalive_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keepalive_timer: %v", err)
		}
	}

	if err = d.Set("log_neighbour_changes", flattenRouterBgpLogNeighbourChanges(o["log-neighbour-changes"], d, "log_neighbour_changes")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-neighbour-changes"], "RouterBgp-LogNeighbourChanges"); ok {
			if err = d.Set("log_neighbour_changes", vv); err != nil {
				return fmt.Errorf("Error reading log_neighbour_changes: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_neighbour_changes: %v", err)
		}
	}

	if err = d.Set("multipath_recursive_distance", flattenRouterBgpMultipathRecursiveDistance(o["multipath-recursive-distance"], d, "multipath_recursive_distance")); err != nil {
		if vv, ok := fortiAPIPatch(o["multipath-recursive-distance"], "RouterBgp-MultipathRecursiveDistance"); ok {
			if err = d.Set("multipath_recursive_distance", vv); err != nil {
				return fmt.Errorf("Error reading multipath_recursive_distance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multipath_recursive_distance: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("neighbor", flattenRouterBgpNeighbor(o["neighbor"], d, "neighbor")); err != nil {
			if vv, ok := fortiAPIPatch(o["neighbor"], "RouterBgp-Neighbor"); ok {
				if err = d.Set("neighbor", vv); err != nil {
					return fmt.Errorf("Error reading neighbor: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading neighbor: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("neighbor"); ok {
			if err = d.Set("neighbor", flattenRouterBgpNeighbor(o["neighbor"], d, "neighbor")); err != nil {
				if vv, ok := fortiAPIPatch(o["neighbor"], "RouterBgp-Neighbor"); ok {
					if err = d.Set("neighbor", vv); err != nil {
						return fmt.Errorf("Error reading neighbor: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading neighbor: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("neighbor_group", flattenRouterBgpNeighborGroup(o["neighbor-group"], d, "neighbor_group")); err != nil {
			if vv, ok := fortiAPIPatch(o["neighbor-group"], "RouterBgp-NeighborGroup"); ok {
				if err = d.Set("neighbor_group", vv); err != nil {
					return fmt.Errorf("Error reading neighbor_group: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading neighbor_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("neighbor_group"); ok {
			if err = d.Set("neighbor_group", flattenRouterBgpNeighborGroup(o["neighbor-group"], d, "neighbor_group")); err != nil {
				if vv, ok := fortiAPIPatch(o["neighbor-group"], "RouterBgp-NeighborGroup"); ok {
					if err = d.Set("neighbor_group", vv); err != nil {
						return fmt.Errorf("Error reading neighbor_group: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading neighbor_group: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("neighbor_range", flattenRouterBgpNeighborRange(o["neighbor-range"], d, "neighbor_range")); err != nil {
			if vv, ok := fortiAPIPatch(o["neighbor-range"], "RouterBgp-NeighborRange"); ok {
				if err = d.Set("neighbor_range", vv); err != nil {
					return fmt.Errorf("Error reading neighbor_range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading neighbor_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("neighbor_range"); ok {
			if err = d.Set("neighbor_range", flattenRouterBgpNeighborRange(o["neighbor-range"], d, "neighbor_range")); err != nil {
				if vv, ok := fortiAPIPatch(o["neighbor-range"], "RouterBgp-NeighborRange"); ok {
					if err = d.Set("neighbor_range", vv); err != nil {
						return fmt.Errorf("Error reading neighbor_range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading neighbor_range: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("neighbor_range6", flattenRouterBgpNeighborRange6(o["neighbor-range6"], d, "neighbor_range6")); err != nil {
			if vv, ok := fortiAPIPatch(o["neighbor-range6"], "RouterBgp-NeighborRange6"); ok {
				if err = d.Set("neighbor_range6", vv); err != nil {
					return fmt.Errorf("Error reading neighbor_range6: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading neighbor_range6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("neighbor_range6"); ok {
			if err = d.Set("neighbor_range6", flattenRouterBgpNeighborRange6(o["neighbor-range6"], d, "neighbor_range6")); err != nil {
				if vv, ok := fortiAPIPatch(o["neighbor-range6"], "RouterBgp-NeighborRange6"); ok {
					if err = d.Set("neighbor_range6", vv); err != nil {
						return fmt.Errorf("Error reading neighbor_range6: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading neighbor_range6: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("network", flattenRouterBgpNetwork(o["network"], d, "network")); err != nil {
			if vv, ok := fortiAPIPatch(o["network"], "RouterBgp-Network"); ok {
				if err = d.Set("network", vv); err != nil {
					return fmt.Errorf("Error reading network: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading network: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("network"); ok {
			if err = d.Set("network", flattenRouterBgpNetwork(o["network"], d, "network")); err != nil {
				if vv, ok := fortiAPIPatch(o["network"], "RouterBgp-Network"); ok {
					if err = d.Set("network", vv); err != nil {
						return fmt.Errorf("Error reading network: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading network: %v", err)
				}
			}
		}
	}

	if err = d.Set("network_import_check", flattenRouterBgpNetworkImportCheck(o["network-import-check"], d, "network_import_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["network-import-check"], "RouterBgp-NetworkImportCheck"); ok {
			if err = d.Set("network_import_check", vv); err != nil {
				return fmt.Errorf("Error reading network_import_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading network_import_check: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("network6", flattenRouterBgpNetwork6(o["network6"], d, "network6")); err != nil {
			if vv, ok := fortiAPIPatch(o["network6"], "RouterBgp-Network6"); ok {
				if err = d.Set("network6", vv); err != nil {
					return fmt.Errorf("Error reading network6: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading network6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("network6"); ok {
			if err = d.Set("network6", flattenRouterBgpNetwork6(o["network6"], d, "network6")); err != nil {
				if vv, ok := fortiAPIPatch(o["network6"], "RouterBgp-Network6"); ok {
					if err = d.Set("network6", vv); err != nil {
						return fmt.Errorf("Error reading network6: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading network6: %v", err)
				}
			}
		}
	}

	if err = d.Set("recursive_inherit_priority", flattenRouterBgpRecursiveInheritPriority(o["recursive-inherit-priority"], d, "recursive_inherit_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["recursive-inherit-priority"], "RouterBgp-RecursiveInheritPriority"); ok {
			if err = d.Set("recursive_inherit_priority", vv); err != nil {
				return fmt.Errorf("Error reading recursive_inherit_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading recursive_inherit_priority: %v", err)
		}
	}

	if err = d.Set("recursive_next_hop", flattenRouterBgpRecursiveNextHop(o["recursive-next-hop"], d, "recursive_next_hop")); err != nil {
		if vv, ok := fortiAPIPatch(o["recursive-next-hop"], "RouterBgp-RecursiveNextHop"); ok {
			if err = d.Set("recursive_next_hop", vv); err != nil {
				return fmt.Errorf("Error reading recursive_next_hop: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading recursive_next_hop: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("redistribute", flattenRouterBgpRedistribute(o["redistribute"], d, "redistribute")); err != nil {
			if vv, ok := fortiAPIPatch(o["redistribute"], "RouterBgp-Redistribute"); ok {
				if err = d.Set("redistribute", vv); err != nil {
					return fmt.Errorf("Error reading redistribute: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading redistribute: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("redistribute"); ok {
			if err = d.Set("redistribute", flattenRouterBgpRedistribute(o["redistribute"], d, "redistribute")); err != nil {
				if vv, ok := fortiAPIPatch(o["redistribute"], "RouterBgp-Redistribute"); ok {
					if err = d.Set("redistribute", vv); err != nil {
						return fmt.Errorf("Error reading redistribute: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading redistribute: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("redistribute6", flattenRouterBgpRedistribute6(o["redistribute6"], d, "redistribute6")); err != nil {
			if vv, ok := fortiAPIPatch(o["redistribute6"], "RouterBgp-Redistribute6"); ok {
				if err = d.Set("redistribute6", vv); err != nil {
					return fmt.Errorf("Error reading redistribute6: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading redistribute6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("redistribute6"); ok {
			if err = d.Set("redistribute6", flattenRouterBgpRedistribute6(o["redistribute6"], d, "redistribute6")); err != nil {
				if vv, ok := fortiAPIPatch(o["redistribute6"], "RouterBgp-Redistribute6"); ok {
					if err = d.Set("redistribute6", vv); err != nil {
						return fmt.Errorf("Error reading redistribute6: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading redistribute6: %v", err)
				}
			}
		}
	}

	if err = d.Set("router_id", flattenRouterBgpRouterId(o["router-id"], d, "router_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["router-id"], "RouterBgp-RouterId"); ok {
			if err = d.Set("router_id", vv); err != nil {
				return fmt.Errorf("Error reading router_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading router_id: %v", err)
		}
	}

	if err = d.Set("scan_time", flattenRouterBgpScanTime(o["scan-time"], d, "scan_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-time"], "RouterBgp-ScanTime"); ok {
			if err = d.Set("scan_time", vv); err != nil {
				return fmt.Errorf("Error reading scan_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_time: %v", err)
		}
	}

	if err = d.Set("synchronization", flattenRouterBgpSynchronization(o["synchronization"], d, "synchronization")); err != nil {
		if vv, ok := fortiAPIPatch(o["synchronization"], "RouterBgp-Synchronization"); ok {
			if err = d.Set("synchronization", vv); err != nil {
				return fmt.Errorf("Error reading synchronization: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading synchronization: %v", err)
		}
	}

	if err = d.Set("tag_resolve_mode", flattenRouterBgpTagResolveMode(o["tag-resolve-mode"], d, "tag_resolve_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["tag-resolve-mode"], "RouterBgp-TagResolveMode"); ok {
			if err = d.Set("tag_resolve_mode", vv); err != nil {
				return fmt.Errorf("Error reading tag_resolve_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tag_resolve_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("vrf_leak", flattenRouterBgpVrfLeak(o["vrf-leak"], d, "vrf_leak")); err != nil {
			if vv, ok := fortiAPIPatch(o["vrf-leak"], "RouterBgp-VrfLeak"); ok {
				if err = d.Set("vrf_leak", vv); err != nil {
					return fmt.Errorf("Error reading vrf_leak: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading vrf_leak: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vrf_leak"); ok {
			if err = d.Set("vrf_leak", flattenRouterBgpVrfLeak(o["vrf-leak"], d, "vrf_leak")); err != nil {
				if vv, ok := fortiAPIPatch(o["vrf-leak"], "RouterBgp-VrfLeak"); ok {
					if err = d.Set("vrf_leak", vv); err != nil {
						return fmt.Errorf("Error reading vrf_leak: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading vrf_leak: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("vrf_leak6", flattenRouterBgpVrfLeak6(o["vrf-leak6"], d, "vrf_leak6")); err != nil {
			if vv, ok := fortiAPIPatch(o["vrf-leak6"], "RouterBgp-VrfLeak6"); ok {
				if err = d.Set("vrf_leak6", vv); err != nil {
					return fmt.Errorf("Error reading vrf_leak6: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading vrf_leak6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vrf_leak6"); ok {
			if err = d.Set("vrf_leak6", flattenRouterBgpVrfLeak6(o["vrf-leak6"], d, "vrf_leak6")); err != nil {
				if vv, ok := fortiAPIPatch(o["vrf-leak6"], "RouterBgp-VrfLeak6"); ok {
					if err = d.Set("vrf_leak6", vv); err != nil {
						return fmt.Errorf("Error reading vrf_leak6: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading vrf_leak6: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("vrf", flattenRouterBgpVrf(o["vrf"], d, "vrf")); err != nil {
			if vv, ok := fortiAPIPatch(o["vrf"], "RouterBgp-Vrf"); ok {
				if err = d.Set("vrf", vv); err != nil {
					return fmt.Errorf("Error reading vrf: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading vrf: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vrf"); ok {
			if err = d.Set("vrf", flattenRouterBgpVrf(o["vrf"], d, "vrf")); err != nil {
				if vv, ok := fortiAPIPatch(o["vrf"], "RouterBgp-Vrf"); ok {
					if err = d.Set("vrf", vv); err != nil {
						return fmt.Errorf("Error reading vrf: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading vrf: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("vrf6", flattenRouterBgpVrf6(o["vrf6"], d, "vrf6")); err != nil {
			if vv, ok := fortiAPIPatch(o["vrf6"], "RouterBgp-Vrf6"); ok {
				if err = d.Set("vrf6", vv); err != nil {
					return fmt.Errorf("Error reading vrf6: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading vrf6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vrf6"); ok {
			if err = d.Set("vrf6", flattenRouterBgpVrf6(o["vrf6"], d, "vrf6")); err != nil {
				if vv, ok := fortiAPIPatch(o["vrf6"], "RouterBgp-Vrf6"); ok {
					if err = d.Set("vrf6", vv); err != nil {
						return fmt.Errorf("Error reading vrf6: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading vrf6: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenRouterBgpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterBgpAdditionalPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdditionalPathSelect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdditionalPathSelectVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdditionalPathSelectVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdditionalPathSelect6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdditionalPathVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdditionalPathVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdditionalPath6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdminDistance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distance"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["distance"], _ = expandRouterBgpAdminDistanceDistance(d, i["distance"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterBgpAdminDistanceId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbour_prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["neighbour-prefix"], _ = expandRouterBgpAdminDistanceNeighbourPrefix(d, i["neighbour_prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_list"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-list"], _ = expandRouterBgpAdminDistanceRouteList(d, i["route_list"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterBgpAdminDistanceDistance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdminDistanceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdminDistanceNeighbourPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandRouterBgpAdminDistanceRouteList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpAggregateAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_set"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["as-set"], _ = expandRouterBgpAggregateAddressAsSet(d, i["as_set"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterBgpAggregateAddressId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandRouterBgpAggregateAddressPrefix(d, i["prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "summary_only"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["summary-only"], _ = expandRouterBgpAggregateAddressSummaryOnly(d, i["summary_only"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterBgpAggregateAddressAsSet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAggregateAddressId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAggregateAddressPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandRouterBgpAggregateAddressSummaryOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAggregateAddress6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_set"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["as-set"], _ = expandRouterBgpAggregateAddress6AsSet(d, i["as_set"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterBgpAggregateAddress6Id(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix6"], _ = expandRouterBgpAggregateAddress6Prefix6(d, i["prefix6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "summary_only"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["summary-only"], _ = expandRouterBgpAggregateAddress6SummaryOnly(d, i["summary_only"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterBgpAggregateAddress6AsSet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAggregateAddress6Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAggregateAddress6Prefix6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAggregateAddress6SummaryOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAlwaysCompareMed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpBestpathAsPathIgnore(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpBestpathCmpConfedAspath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpBestpathCmpRouterid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpBestpathMedConfed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpBestpathMedMissingAsWorst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpClientToClientReflection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpClusterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpConfederationIdentifier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpConfederationPeers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpCrossFamilyConditionalAdv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDampening(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDampeningMaxSuppressTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDampeningReachabilityHalfLife(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDampeningReuse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDampeningRouteMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpDampeningSuppress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDampeningUnreachabilityHalfLife(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDefaultLocalPreference(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDeterministicMed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDistanceExternal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDistanceInternal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDistanceLocal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpEbgpMultipath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpEnforceFirstAs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpFastExternalFailover(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpGracefulEndOnTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpGracefulRestart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpGracefulRestartTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpGracefulStalepathTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpGracefulUpdateDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpHoldtimeTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpIbgpMultipath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpIgnoreOptionalCapability(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpKeepaliveTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpLogNeighbourChanges(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpMultipathRecursiveDistance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighbor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["activate"], _ = expandRouterBgpNeighborActivate(d, i["activate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate_evpn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["activate-evpn"], _ = expandRouterBgpNeighborActivateEvpn(d, i["activate_evpn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["activate-vpnv4"], _ = expandRouterBgpNeighborActivateVpnv4(d, i["activate_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["activate-vpnv6"], _ = expandRouterBgpNeighborActivateVpnv6(d, i["activate_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["activate6"], _ = expandRouterBgpNeighborActivate6(d, i["activate6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["additional-path"], _ = expandRouterBgpNeighborAdditionalPath(d, i["additional_path"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["additional-path-vpnv4"], _ = expandRouterBgpNeighborAdditionalPathVpnv4(d, i["additional_path_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["additional-path-vpnv6"], _ = expandRouterBgpNeighborAdditionalPathVpnv6(d, i["additional_path_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["additional-path6"], _ = expandRouterBgpNeighborAdditionalPath6(d, i["additional_path6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["adv-additional-path"], _ = expandRouterBgpNeighborAdvAdditionalPath(d, i["adv_additional_path"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["adv-additional-path-vpnv4"], _ = expandRouterBgpNeighborAdvAdditionalPathVpnv4(d, i["adv_additional_path_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["adv-additional-path-vpnv6"], _ = expandRouterBgpNeighborAdvAdditionalPathVpnv6(d, i["adv_additional_path_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["adv-additional-path6"], _ = expandRouterBgpNeighborAdvAdditionalPath6(d, i["adv_additional_path6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertisement_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["advertisement-interval"], _ = expandRouterBgpNeighborAdvertisementInterval(d, i["advertisement_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowas-in"], _ = expandRouterBgpNeighborAllowasIn(d, i["allowas_in"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowas-in-enable"], _ = expandRouterBgpNeighborAllowasInEnable(d, i["allowas_in_enable"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable_evpn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowas-in-enable-evpn"], _ = expandRouterBgpNeighborAllowasInEnableEvpn(d, i["allowas_in_enable_evpn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowas-in-enable-vpnv4"], _ = expandRouterBgpNeighborAllowasInEnableVpnv4(d, i["allowas_in_enable_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowas-in-enable-vpnv6"], _ = expandRouterBgpNeighborAllowasInEnableVpnv6(d, i["allowas_in_enable_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowas-in-enable6"], _ = expandRouterBgpNeighborAllowasInEnable6(d, i["allowas_in_enable6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_evpn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowas-in-evpn"], _ = expandRouterBgpNeighborAllowasInEvpn(d, i["allowas_in_evpn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowas-in-vpnv4"], _ = expandRouterBgpNeighborAllowasInVpnv4(d, i["allowas_in_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowas-in-vpnv6"], _ = expandRouterBgpNeighborAllowasInVpnv6(d, i["allowas_in_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowas-in6"], _ = expandRouterBgpNeighborAllowasIn6(d, i["allowas_in6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_override"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["as-override"], _ = expandRouterBgpNeighborAsOverride(d, i["as_override"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_override6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["as-override6"], _ = expandRouterBgpNeighborAsOverride6(d, i["as_override6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["attribute-unchanged"], _ = expandRouterBgpNeighborAttributeUnchanged(d, i["attribute_unchanged"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["attribute-unchanged-vpnv4"], _ = expandRouterBgpNeighborAttributeUnchangedVpnv4(d, i["attribute_unchanged_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["attribute-unchanged-vpnv6"], _ = expandRouterBgpNeighborAttributeUnchangedVpnv6(d, i["attribute_unchanged_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["attribute-unchanged6"], _ = expandRouterBgpNeighborAttributeUnchanged6(d, i["attribute_unchanged6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_options"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-options"], _ = expandRouterBgpNeighborAuthOptions(d, i["auth_options"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bfd"], _ = expandRouterBgpNeighborBfd(d, i["bfd"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_default_originate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["capability-default-originate"], _ = expandRouterBgpNeighborCapabilityDefaultOriginate(d, i["capability_default_originate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_default_originate6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["capability-default-originate6"], _ = expandRouterBgpNeighborCapabilityDefaultOriginate6(d, i["capability_default_originate6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_dynamic"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["capability-dynamic"], _ = expandRouterBgpNeighborCapabilityDynamic(d, i["capability_dynamic"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["capability-graceful-restart"], _ = expandRouterBgpNeighborCapabilityGracefulRestart(d, i["capability_graceful_restart"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart_evpn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["capability-graceful-restart-evpn"], _ = expandRouterBgpNeighborCapabilityGracefulRestartEvpn(d, i["capability_graceful_restart_evpn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["capability-graceful-restart-vpnv4"], _ = expandRouterBgpNeighborCapabilityGracefulRestartVpnv4(d, i["capability_graceful_restart_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["capability-graceful-restart-vpnv6"], _ = expandRouterBgpNeighborCapabilityGracefulRestartVpnv6(d, i["capability_graceful_restart_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["capability-graceful-restart6"], _ = expandRouterBgpNeighborCapabilityGracefulRestart6(d, i["capability_graceful_restart6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_orf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["capability-orf"], _ = expandRouterBgpNeighborCapabilityOrf(d, i["capability_orf"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_orf6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["capability-orf6"], _ = expandRouterBgpNeighborCapabilityOrf6(d, i["capability_orf6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_route_refresh"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["capability-route-refresh"], _ = expandRouterBgpNeighborCapabilityRouteRefresh(d, i["capability_route_refresh"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "conditional_advertise"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandRouterBgpNeighborConditionalAdvertise(d, i["conditional_advertise"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["conditional-advertise"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "conditional_advertise6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandRouterBgpNeighborConditionalAdvertise6(d, i["conditional_advertise6"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["conditional-advertise6"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "connect_timer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["connect-timer"], _ = expandRouterBgpNeighborConnectTimer(d, i["connect_timer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_originate_routemap"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["default-originate-routemap"], _ = expandRouterBgpNeighborDefaultOriginateRoutemap(d, i["default_originate_routemap"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_originate_routemap6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["default-originate-routemap6"], _ = expandRouterBgpNeighborDefaultOriginateRoutemap6(d, i["default_originate_routemap6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandRouterBgpNeighborDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["distribute-list-in"], _ = expandRouterBgpNeighborDistributeListIn(d, i["distribute_list_in"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["distribute-list-in-vpnv4"], _ = expandRouterBgpNeighborDistributeListInVpnv4(d, i["distribute_list_in_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["distribute-list-in-vpnv6"], _ = expandRouterBgpNeighborDistributeListInVpnv6(d, i["distribute_list_in_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["distribute-list-in6"], _ = expandRouterBgpNeighborDistributeListIn6(d, i["distribute_list_in6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["distribute-list-out"], _ = expandRouterBgpNeighborDistributeListOut(d, i["distribute_list_out"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["distribute-list-out-vpnv4"], _ = expandRouterBgpNeighborDistributeListOutVpnv4(d, i["distribute_list_out_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["distribute-list-out-vpnv6"], _ = expandRouterBgpNeighborDistributeListOutVpnv6(d, i["distribute_list_out_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["distribute-list-out6"], _ = expandRouterBgpNeighborDistributeListOut6(d, i["distribute_list_out6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dont_capability_negotiate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dont-capability-negotiate"], _ = expandRouterBgpNeighborDontCapabilityNegotiate(d, i["dont_capability_negotiate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_enforce_multihop"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ebgp-enforce-multihop"], _ = expandRouterBgpNeighborEbgpEnforceMultihop(d, i["ebgp_enforce_multihop"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_multihop_ttl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ebgp-multihop-ttl"], _ = expandRouterBgpNeighborEbgpMultihopTtl(d, i["ebgp_multihop_ttl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter-list-in"], _ = expandRouterBgpNeighborFilterListIn(d, i["filter_list_in"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter-list-in-vpnv4"], _ = expandRouterBgpNeighborFilterListInVpnv4(d, i["filter_list_in_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter-list-in-vpnv6"], _ = expandRouterBgpNeighborFilterListInVpnv6(d, i["filter_list_in_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter-list-in6"], _ = expandRouterBgpNeighborFilterListIn6(d, i["filter_list_in6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter-list-out"], _ = expandRouterBgpNeighborFilterListOut(d, i["filter_list_out"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter-list-out-vpnv4"], _ = expandRouterBgpNeighborFilterListOutVpnv4(d, i["filter_list_out_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter-list-out-vpnv6"], _ = expandRouterBgpNeighborFilterListOutVpnv6(d, i["filter_list_out_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter-list-out6"], _ = expandRouterBgpNeighborFilterListOut6(d, i["filter_list_out6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holdtime_timer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["holdtime-timer"], _ = expandRouterBgpNeighborHoldtimeTimer(d, i["holdtime_timer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandRouterBgpNeighborInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandRouterBgpNeighborIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keep_alive_timer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["keep-alive-timer"], _ = expandRouterBgpNeighborKeepAliveTimer(d, i["keep_alive_timer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_down_failover"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["link-down-failover"], _ = expandRouterBgpNeighborLinkDownFailover(d, i["link_down_failover"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["local-as"], _ = expandRouterBgpNeighborLocalAs(d, i["local_as"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as_no_prepend"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["local-as-no-prepend"], _ = expandRouterBgpNeighborLocalAsNoPrepend(d, i["local_as_no_prepend"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as_replace_as"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["local-as-replace-as"], _ = expandRouterBgpNeighborLocalAsReplaceAs(d, i["local_as_replace_as"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-prefix"], _ = expandRouterBgpNeighborMaximumPrefix(d, i["maximum_prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_evpn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-prefix-evpn"], _ = expandRouterBgpNeighborMaximumPrefixEvpn(d, i["maximum_prefix_evpn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-prefix-threshold"], _ = expandRouterBgpNeighborMaximumPrefixThreshold(d, i["maximum_prefix_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold_evpn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-prefix-threshold-evpn"], _ = expandRouterBgpNeighborMaximumPrefixThresholdEvpn(d, i["maximum_prefix_threshold_evpn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-prefix-threshold-vpnv4"], _ = expandRouterBgpNeighborMaximumPrefixThresholdVpnv4(d, i["maximum_prefix_threshold_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-prefix-threshold-vpnv6"], _ = expandRouterBgpNeighborMaximumPrefixThresholdVpnv6(d, i["maximum_prefix_threshold_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-prefix-threshold6"], _ = expandRouterBgpNeighborMaximumPrefixThreshold6(d, i["maximum_prefix_threshold6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-prefix-vpnv4"], _ = expandRouterBgpNeighborMaximumPrefixVpnv4(d, i["maximum_prefix_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-prefix-vpnv6"], _ = expandRouterBgpNeighborMaximumPrefixVpnv6(d, i["maximum_prefix_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-prefix-warning-only"], _ = expandRouterBgpNeighborMaximumPrefixWarningOnly(d, i["maximum_prefix_warning_only"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only_evpn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-prefix-warning-only-evpn"], _ = expandRouterBgpNeighborMaximumPrefixWarningOnlyEvpn(d, i["maximum_prefix_warning_only_evpn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-prefix-warning-only-vpnv4"], _ = expandRouterBgpNeighborMaximumPrefixWarningOnlyVpnv4(d, i["maximum_prefix_warning_only_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-prefix-warning-only-vpnv6"], _ = expandRouterBgpNeighborMaximumPrefixWarningOnlyVpnv6(d, i["maximum_prefix_warning_only_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-prefix-warning-only6"], _ = expandRouterBgpNeighborMaximumPrefixWarningOnly6(d, i["maximum_prefix_warning_only6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-prefix6"], _ = expandRouterBgpNeighborMaximumPrefix6(d, i["maximum_prefix6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["next-hop-self"], _ = expandRouterBgpNeighborNextHopSelf(d, i["next_hop_self"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_rr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["next-hop-self-rr"], _ = expandRouterBgpNeighborNextHopSelfRr(d, i["next_hop_self_rr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_rr6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["next-hop-self-rr6"], _ = expandRouterBgpNeighborNextHopSelfRr6(d, i["next_hop_self_rr6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["next-hop-self-vpnv4"], _ = expandRouterBgpNeighborNextHopSelfVpnv4(d, i["next_hop_self_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["next-hop-self-vpnv6"], _ = expandRouterBgpNeighborNextHopSelfVpnv6(d, i["next_hop_self_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["next-hop-self6"], _ = expandRouterBgpNeighborNextHopSelf6(d, i["next_hop_self6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_capability"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["override-capability"], _ = expandRouterBgpNeighborOverrideCapability(d, i["override_capability"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["passive"], _ = expandRouterBgpNeighborPassive(d, i["passive"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password"], _ = expandRouterBgpNeighborPassword(d, i["password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix-list-in"], _ = expandRouterBgpNeighborPrefixListIn(d, i["prefix_list_in"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix-list-in-vpnv4"], _ = expandRouterBgpNeighborPrefixListInVpnv4(d, i["prefix_list_in_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix-list-in-vpnv6"], _ = expandRouterBgpNeighborPrefixListInVpnv6(d, i["prefix_list_in_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix-list-in6"], _ = expandRouterBgpNeighborPrefixListIn6(d, i["prefix_list_in6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix-list-out"], _ = expandRouterBgpNeighborPrefixListOut(d, i["prefix_list_out"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix-list-out-vpnv4"], _ = expandRouterBgpNeighborPrefixListOutVpnv4(d, i["prefix_list_out_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix-list-out-vpnv6"], _ = expandRouterBgpNeighborPrefixListOutVpnv6(d, i["prefix_list_out_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix-list-out6"], _ = expandRouterBgpNeighborPrefixListOut6(d, i["prefix_list_out6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_as"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remote-as"], _ = expandRouterBgpNeighborRemoteAs(d, i["remote_as"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remove-private-as"], _ = expandRouterBgpNeighborRemovePrivateAs(d, i["remove_private_as"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as_evpn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remove-private-as-evpn"], _ = expandRouterBgpNeighborRemovePrivateAsEvpn(d, i["remove_private_as_evpn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remove-private-as-vpnv4"], _ = expandRouterBgpNeighborRemovePrivateAsVpnv4(d, i["remove_private_as_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remove-private-as-vpnv6"], _ = expandRouterBgpNeighborRemovePrivateAsVpnv6(d, i["remove_private_as_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remove-private-as6"], _ = expandRouterBgpNeighborRemovePrivateAs6(d, i["remove_private_as6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restart_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["restart-time"], _ = expandRouterBgpNeighborRestartTime(d, i["restart_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retain_stale_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["retain-stale-time"], _ = expandRouterBgpNeighborRetainStaleTime(d, i["retain_stale_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map-in"], _ = expandRouterBgpNeighborRouteMapIn(d, i["route_map_in"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in_evpn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map-in-evpn"], _ = expandRouterBgpNeighborRouteMapInEvpn(d, i["route_map_in_evpn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map-in-vpnv4"], _ = expandRouterBgpNeighborRouteMapInVpnv4(d, i["route_map_in_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map-in-vpnv6"], _ = expandRouterBgpNeighborRouteMapInVpnv6(d, i["route_map_in_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map-in6"], _ = expandRouterBgpNeighborRouteMapIn6(d, i["route_map_in6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map-out"], _ = expandRouterBgpNeighborRouteMapOut(d, i["route_map_out"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_evpn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map-out-evpn"], _ = expandRouterBgpNeighborRouteMapOutEvpn(d, i["route_map_out_evpn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_preferable"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map-out-preferable"], _ = expandRouterBgpNeighborRouteMapOutPreferable(d, i["route_map_out_preferable"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map-out-vpnv4"], _ = expandRouterBgpNeighborRouteMapOutVpnv4(d, i["route_map_out_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv4_preferable"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map-out-vpnv4-preferable"], _ = expandRouterBgpNeighborRouteMapOutVpnv4Preferable(d, i["route_map_out_vpnv4_preferable"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map-out-vpnv6"], _ = expandRouterBgpNeighborRouteMapOutVpnv6(d, i["route_map_out_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv6_preferable"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map-out-vpnv6-preferable"], _ = expandRouterBgpNeighborRouteMapOutVpnv6Preferable(d, i["route_map_out_vpnv6_preferable"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map-out6"], _ = expandRouterBgpNeighborRouteMapOut6(d, i["route_map_out6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out6_preferable"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map-out6-preferable"], _ = expandRouterBgpNeighborRouteMapOut6Preferable(d, i["route_map_out6_preferable"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-reflector-client"], _ = expandRouterBgpNeighborRouteReflectorClient(d, i["route_reflector_client"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client_evpn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-reflector-client-evpn"], _ = expandRouterBgpNeighborRouteReflectorClientEvpn(d, i["route_reflector_client_evpn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-reflector-client-vpnv4"], _ = expandRouterBgpNeighborRouteReflectorClientVpnv4(d, i["route_reflector_client_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-reflector-client-vpnv6"], _ = expandRouterBgpNeighborRouteReflectorClientVpnv6(d, i["route_reflector_client_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-reflector-client6"], _ = expandRouterBgpNeighborRouteReflectorClient6(d, i["route_reflector_client6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-server-client"], _ = expandRouterBgpNeighborRouteServerClient(d, i["route_server_client"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client_evpn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-server-client-evpn"], _ = expandRouterBgpNeighborRouteServerClientEvpn(d, i["route_server_client_evpn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-server-client-vpnv4"], _ = expandRouterBgpNeighborRouteServerClientVpnv4(d, i["route_server_client_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-server-client-vpnv6"], _ = expandRouterBgpNeighborRouteServerClientVpnv6(d, i["route_server_client_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-server-client6"], _ = expandRouterBgpNeighborRouteServerClient6(d, i["route_server_client6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["send-community"], _ = expandRouterBgpNeighborSendCommunity(d, i["send_community"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community_evpn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["send-community-evpn"], _ = expandRouterBgpNeighborSendCommunityEvpn(d, i["send_community_evpn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["send-community-vpnv4"], _ = expandRouterBgpNeighborSendCommunityVpnv4(d, i["send_community_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["send-community-vpnv6"], _ = expandRouterBgpNeighborSendCommunityVpnv6(d, i["send_community_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["send-community6"], _ = expandRouterBgpNeighborSendCommunity6(d, i["send_community6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shutdown"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["shutdown"], _ = expandRouterBgpNeighborShutdown(d, i["shutdown"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["soft-reconfiguration"], _ = expandRouterBgpNeighborSoftReconfiguration(d, i["soft_reconfiguration"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration_evpn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["soft-reconfiguration-evpn"], _ = expandRouterBgpNeighborSoftReconfigurationEvpn(d, i["soft_reconfiguration_evpn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["soft-reconfiguration-vpnv4"], _ = expandRouterBgpNeighborSoftReconfigurationVpnv4(d, i["soft_reconfiguration_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["soft-reconfiguration-vpnv6"], _ = expandRouterBgpNeighborSoftReconfigurationVpnv6(d, i["soft_reconfiguration_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["soft-reconfiguration6"], _ = expandRouterBgpNeighborSoftReconfiguration6(d, i["soft_reconfiguration6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stale_route"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["stale-route"], _ = expandRouterBgpNeighborStaleRoute(d, i["stale_route"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "strict_capability_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["strict-capability-match"], _ = expandRouterBgpNeighborStrictCapabilityMatch(d, i["strict_capability_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unsuppress_map"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["unsuppress-map"], _ = expandRouterBgpNeighborUnsuppressMap(d, i["unsuppress_map"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unsuppress_map6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["unsuppress-map6"], _ = expandRouterBgpNeighborUnsuppressMap6(d, i["unsuppress_map6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_source"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["update-source"], _ = expandRouterBgpNeighborUpdateSource(d, i["update_source"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandRouterBgpNeighborWeight(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterBgpNeighborActivate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborActivateEvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborActivateVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborActivateVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborActivate6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdditionalPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdditionalPathVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdditionalPathVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdditionalPath6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdvAdditionalPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdvAdditionalPathVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdvAdditionalPathVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdvAdditionalPath6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdvertisementInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasIn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasInEnable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasInEnableEvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasInEnableVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasInEnableVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasInEnable6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasInEvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasInVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasInVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasIn6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAsOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAsOverride6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAttributeUnchanged(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborAttributeUnchangedVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborAttributeUnchangedVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborAttributeUnchanged6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborAuthOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborBfd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityDefaultOriginate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityDefaultOriginate6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityDynamic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityGracefulRestart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityGracefulRestartEvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityGracefulRestartVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityGracefulRestartVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityGracefulRestart6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityOrf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityOrf6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityRouteRefresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborConditionalAdvertise(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise_routemap"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["advertise-routemap"], _ = expandRouterBgpNeighborConditionalAdvertiseAdvertiseRoutemap(d, i["advertise_routemap"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_routemap"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["condition-routemap"], _ = expandRouterBgpNeighborConditionalAdvertiseConditionRoutemap(d, i["condition_routemap"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["condition-type"], _ = expandRouterBgpNeighborConditionalAdvertiseConditionType(d, i["condition_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterBgpNeighborConditionalAdvertiseAdvertiseRoutemap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborConditionalAdvertiseConditionRoutemap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborConditionalAdvertiseConditionType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborConditionalAdvertise6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise_routemap"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["advertise-routemap"], _ = expandRouterBgpNeighborConditionalAdvertise6AdvertiseRoutemap(d, i["advertise_routemap"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_routemap"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["condition-routemap"], _ = expandRouterBgpNeighborConditionalAdvertise6ConditionRoutemap(d, i["condition_routemap"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["condition-type"], _ = expandRouterBgpNeighborConditionalAdvertise6ConditionType(d, i["condition_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterBgpNeighborConditionalAdvertise6AdvertiseRoutemap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborConditionalAdvertise6ConditionRoutemap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborConditionalAdvertise6ConditionType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborConnectTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborDefaultOriginateRoutemap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborDefaultOriginateRoutemap6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborDistributeListIn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborDistributeListInVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborDistributeListInVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborDistributeListIn6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborDistributeListOut(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborDistributeListOutVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborDistributeListOutVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborDistributeListOut6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborDontCapabilityNegotiate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborEbgpEnforceMultihop(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborEbgpMultihopTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborFilterListIn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborFilterListInVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborFilterListInVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborFilterListIn6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborFilterListOut(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborFilterListOutVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborFilterListOutVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborFilterListOut6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborHoldtimeTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborKeepAliveTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborLinkDownFailover(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborLocalAs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborLocalAsNoPrepend(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborLocalAsReplaceAs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixEvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixThresholdEvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixThresholdVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixThresholdVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixThreshold6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixWarningOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixWarningOnlyEvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixWarningOnlyVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixWarningOnlyVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixWarningOnly6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefix6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborNextHopSelf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborNextHopSelfRr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborNextHopSelfRr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborNextHopSelfVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborNextHopSelfVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborNextHopSelf6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborOverrideCapability(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborPassive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborPrefixListIn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborPrefixListInVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborPrefixListInVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborPrefixListIn6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborPrefixListOut(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborPrefixListOutVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborPrefixListOutVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborPrefixListOut6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRemoteAs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRemovePrivateAs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRemovePrivateAsEvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRemovePrivateAsVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRemovePrivateAsVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRemovePrivateAs6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRestartTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRetainStaleTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteMapIn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRouteMapInEvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRouteMapInVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRouteMapInVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRouteMapIn6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRouteMapOut(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRouteMapOutEvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRouteMapOutPreferable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRouteMapOutVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRouteMapOutVpnv4Preferable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRouteMapOutVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRouteMapOutVpnv6Preferable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRouteMapOut6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRouteMapOut6Preferable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRouteReflectorClient(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteReflectorClientEvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteReflectorClientVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteReflectorClientVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteReflectorClient6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteServerClient(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteServerClientEvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteServerClientVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteServerClientVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteServerClient6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSendCommunity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSendCommunityEvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSendCommunityVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSendCommunityVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSendCommunity6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborShutdown(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSoftReconfiguration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSoftReconfigurationEvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSoftReconfigurationVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSoftReconfigurationVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSoftReconfiguration6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborStaleRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborStrictCapabilityMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborUnsuppressMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborUnsuppressMap6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborUpdateSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["activate"], _ = expandRouterBgpNeighborGroupActivate(d, i["activate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate_evpn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["activate-evpn"], _ = expandRouterBgpNeighborGroupActivateEvpn(d, i["activate_evpn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["activate-vpnv4"], _ = expandRouterBgpNeighborGroupActivateVpnv4(d, i["activate_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["activate-vpnv6"], _ = expandRouterBgpNeighborGroupActivateVpnv6(d, i["activate_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["activate6"], _ = expandRouterBgpNeighborGroupActivate6(d, i["activate6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["additional-path"], _ = expandRouterBgpNeighborGroupAdditionalPath(d, i["additional_path"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["additional-path-vpnv4"], _ = expandRouterBgpNeighborGroupAdditionalPathVpnv4(d, i["additional_path_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["additional-path-vpnv6"], _ = expandRouterBgpNeighborGroupAdditionalPathVpnv6(d, i["additional_path_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["additional-path6"], _ = expandRouterBgpNeighborGroupAdditionalPath6(d, i["additional_path6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["adv-additional-path"], _ = expandRouterBgpNeighborGroupAdvAdditionalPath(d, i["adv_additional_path"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["adv-additional-path-vpnv4"], _ = expandRouterBgpNeighborGroupAdvAdditionalPathVpnv4(d, i["adv_additional_path_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["adv-additional-path-vpnv6"], _ = expandRouterBgpNeighborGroupAdvAdditionalPathVpnv6(d, i["adv_additional_path_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["adv-additional-path6"], _ = expandRouterBgpNeighborGroupAdvAdditionalPath6(d, i["adv_additional_path6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertisement_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["advertisement-interval"], _ = expandRouterBgpNeighborGroupAdvertisementInterval(d, i["advertisement_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowas-in"], _ = expandRouterBgpNeighborGroupAllowasIn(d, i["allowas_in"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowas-in-enable"], _ = expandRouterBgpNeighborGroupAllowasInEnable(d, i["allowas_in_enable"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable_evpn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowas-in-enable-evpn"], _ = expandRouterBgpNeighborGroupAllowasInEnableEvpn(d, i["allowas_in_enable_evpn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowas-in-enable-vpnv4"], _ = expandRouterBgpNeighborGroupAllowasInEnableVpnv4(d, i["allowas_in_enable_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowas-in-enable-vpnv6"], _ = expandRouterBgpNeighborGroupAllowasInEnableVpnv6(d, i["allowas_in_enable_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowas-in-enable6"], _ = expandRouterBgpNeighborGroupAllowasInEnable6(d, i["allowas_in_enable6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_evpn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowas-in-evpn"], _ = expandRouterBgpNeighborGroupAllowasInEvpn(d, i["allowas_in_evpn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowas-in-vpnv4"], _ = expandRouterBgpNeighborGroupAllowasInVpnv4(d, i["allowas_in_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowas-in-vpnv6"], _ = expandRouterBgpNeighborGroupAllowasInVpnv6(d, i["allowas_in_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowas-in6"], _ = expandRouterBgpNeighborGroupAllowasIn6(d, i["allowas_in6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_override"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["as-override"], _ = expandRouterBgpNeighborGroupAsOverride(d, i["as_override"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_override6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["as-override6"], _ = expandRouterBgpNeighborGroupAsOverride6(d, i["as_override6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["attribute-unchanged"], _ = expandRouterBgpNeighborGroupAttributeUnchanged(d, i["attribute_unchanged"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["attribute-unchanged-vpnv4"], _ = expandRouterBgpNeighborGroupAttributeUnchangedVpnv4(d, i["attribute_unchanged_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["attribute-unchanged-vpnv6"], _ = expandRouterBgpNeighborGroupAttributeUnchangedVpnv6(d, i["attribute_unchanged_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["attribute-unchanged6"], _ = expandRouterBgpNeighborGroupAttributeUnchanged6(d, i["attribute_unchanged6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_options"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-options"], _ = expandRouterBgpNeighborGroupAuthOptions(d, i["auth_options"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bfd"], _ = expandRouterBgpNeighborGroupBfd(d, i["bfd"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_default_originate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["capability-default-originate"], _ = expandRouterBgpNeighborGroupCapabilityDefaultOriginate(d, i["capability_default_originate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_default_originate6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["capability-default-originate6"], _ = expandRouterBgpNeighborGroupCapabilityDefaultOriginate6(d, i["capability_default_originate6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_dynamic"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["capability-dynamic"], _ = expandRouterBgpNeighborGroupCapabilityDynamic(d, i["capability_dynamic"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["capability-graceful-restart"], _ = expandRouterBgpNeighborGroupCapabilityGracefulRestart(d, i["capability_graceful_restart"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart_evpn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["capability-graceful-restart-evpn"], _ = expandRouterBgpNeighborGroupCapabilityGracefulRestartEvpn(d, i["capability_graceful_restart_evpn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["capability-graceful-restart-vpnv4"], _ = expandRouterBgpNeighborGroupCapabilityGracefulRestartVpnv4(d, i["capability_graceful_restart_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["capability-graceful-restart-vpnv6"], _ = expandRouterBgpNeighborGroupCapabilityGracefulRestartVpnv6(d, i["capability_graceful_restart_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["capability-graceful-restart6"], _ = expandRouterBgpNeighborGroupCapabilityGracefulRestart6(d, i["capability_graceful_restart6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_orf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["capability-orf"], _ = expandRouterBgpNeighborGroupCapabilityOrf(d, i["capability_orf"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_orf6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["capability-orf6"], _ = expandRouterBgpNeighborGroupCapabilityOrf6(d, i["capability_orf6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_route_refresh"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["capability-route-refresh"], _ = expandRouterBgpNeighborGroupCapabilityRouteRefresh(d, i["capability_route_refresh"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "connect_timer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["connect-timer"], _ = expandRouterBgpNeighborGroupConnectTimer(d, i["connect_timer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_originate_routemap"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["default-originate-routemap"], _ = expandRouterBgpNeighborGroupDefaultOriginateRoutemap(d, i["default_originate_routemap"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_originate_routemap6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["default-originate-routemap6"], _ = expandRouterBgpNeighborGroupDefaultOriginateRoutemap6(d, i["default_originate_routemap6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandRouterBgpNeighborGroupDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["distribute-list-in"], _ = expandRouterBgpNeighborGroupDistributeListIn(d, i["distribute_list_in"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["distribute-list-in-vpnv4"], _ = expandRouterBgpNeighborGroupDistributeListInVpnv4(d, i["distribute_list_in_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["distribute-list-in-vpnv6"], _ = expandRouterBgpNeighborGroupDistributeListInVpnv6(d, i["distribute_list_in_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["distribute-list-in6"], _ = expandRouterBgpNeighborGroupDistributeListIn6(d, i["distribute_list_in6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["distribute-list-out"], _ = expandRouterBgpNeighborGroupDistributeListOut(d, i["distribute_list_out"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["distribute-list-out-vpnv4"], _ = expandRouterBgpNeighborGroupDistributeListOutVpnv4(d, i["distribute_list_out_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["distribute-list-out-vpnv6"], _ = expandRouterBgpNeighborGroupDistributeListOutVpnv6(d, i["distribute_list_out_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["distribute-list-out6"], _ = expandRouterBgpNeighborGroupDistributeListOut6(d, i["distribute_list_out6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dont_capability_negotiate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dont-capability-negotiate"], _ = expandRouterBgpNeighborGroupDontCapabilityNegotiate(d, i["dont_capability_negotiate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_enforce_multihop"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ebgp-enforce-multihop"], _ = expandRouterBgpNeighborGroupEbgpEnforceMultihop(d, i["ebgp_enforce_multihop"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_multihop_ttl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ebgp-multihop-ttl"], _ = expandRouterBgpNeighborGroupEbgpMultihopTtl(d, i["ebgp_multihop_ttl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter-list-in"], _ = expandRouterBgpNeighborGroupFilterListIn(d, i["filter_list_in"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter-list-in-vpnv4"], _ = expandRouterBgpNeighborGroupFilterListInVpnv4(d, i["filter_list_in_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter-list-in-vpnv6"], _ = expandRouterBgpNeighborGroupFilterListInVpnv6(d, i["filter_list_in_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter-list-in6"], _ = expandRouterBgpNeighborGroupFilterListIn6(d, i["filter_list_in6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter-list-out"], _ = expandRouterBgpNeighborGroupFilterListOut(d, i["filter_list_out"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter-list-out-vpnv4"], _ = expandRouterBgpNeighborGroupFilterListOutVpnv4(d, i["filter_list_out_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter-list-out-vpnv6"], _ = expandRouterBgpNeighborGroupFilterListOutVpnv6(d, i["filter_list_out_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter-list-out6"], _ = expandRouterBgpNeighborGroupFilterListOut6(d, i["filter_list_out6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holdtime_timer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["holdtime-timer"], _ = expandRouterBgpNeighborGroupHoldtimeTimer(d, i["holdtime_timer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandRouterBgpNeighborGroupInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keep_alive_timer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["keep-alive-timer"], _ = expandRouterBgpNeighborGroupKeepAliveTimer(d, i["keep_alive_timer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_down_failover"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["link-down-failover"], _ = expandRouterBgpNeighborGroupLinkDownFailover(d, i["link_down_failover"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["local-as"], _ = expandRouterBgpNeighborGroupLocalAs(d, i["local_as"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as_no_prepend"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["local-as-no-prepend"], _ = expandRouterBgpNeighborGroupLocalAsNoPrepend(d, i["local_as_no_prepend"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as_replace_as"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["local-as-replace-as"], _ = expandRouterBgpNeighborGroupLocalAsReplaceAs(d, i["local_as_replace_as"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-prefix"], _ = expandRouterBgpNeighborGroupMaximumPrefix(d, i["maximum_prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_evpn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-prefix-evpn"], _ = expandRouterBgpNeighborGroupMaximumPrefixEvpn(d, i["maximum_prefix_evpn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-prefix-threshold"], _ = expandRouterBgpNeighborGroupMaximumPrefixThreshold(d, i["maximum_prefix_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold_evpn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-prefix-threshold-evpn"], _ = expandRouterBgpNeighborGroupMaximumPrefixThresholdEvpn(d, i["maximum_prefix_threshold_evpn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-prefix-threshold-vpnv4"], _ = expandRouterBgpNeighborGroupMaximumPrefixThresholdVpnv4(d, i["maximum_prefix_threshold_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-prefix-threshold-vpnv6"], _ = expandRouterBgpNeighborGroupMaximumPrefixThresholdVpnv6(d, i["maximum_prefix_threshold_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-prefix-threshold6"], _ = expandRouterBgpNeighborGroupMaximumPrefixThreshold6(d, i["maximum_prefix_threshold6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-prefix-vpnv4"], _ = expandRouterBgpNeighborGroupMaximumPrefixVpnv4(d, i["maximum_prefix_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-prefix-vpnv6"], _ = expandRouterBgpNeighborGroupMaximumPrefixVpnv6(d, i["maximum_prefix_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-prefix-warning-only"], _ = expandRouterBgpNeighborGroupMaximumPrefixWarningOnly(d, i["maximum_prefix_warning_only"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only_evpn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-prefix-warning-only-evpn"], _ = expandRouterBgpNeighborGroupMaximumPrefixWarningOnlyEvpn(d, i["maximum_prefix_warning_only_evpn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-prefix-warning-only-vpnv4"], _ = expandRouterBgpNeighborGroupMaximumPrefixWarningOnlyVpnv4(d, i["maximum_prefix_warning_only_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-prefix-warning-only-vpnv6"], _ = expandRouterBgpNeighborGroupMaximumPrefixWarningOnlyVpnv6(d, i["maximum_prefix_warning_only_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-prefix-warning-only6"], _ = expandRouterBgpNeighborGroupMaximumPrefixWarningOnly6(d, i["maximum_prefix_warning_only6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-prefix6"], _ = expandRouterBgpNeighborGroupMaximumPrefix6(d, i["maximum_prefix6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandRouterBgpNeighborGroupName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["next-hop-self"], _ = expandRouterBgpNeighborGroupNextHopSelf(d, i["next_hop_self"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_rr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["next-hop-self-rr"], _ = expandRouterBgpNeighborGroupNextHopSelfRr(d, i["next_hop_self_rr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_rr6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["next-hop-self-rr6"], _ = expandRouterBgpNeighborGroupNextHopSelfRr6(d, i["next_hop_self_rr6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["next-hop-self-vpnv4"], _ = expandRouterBgpNeighborGroupNextHopSelfVpnv4(d, i["next_hop_self_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["next-hop-self-vpnv6"], _ = expandRouterBgpNeighborGroupNextHopSelfVpnv6(d, i["next_hop_self_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["next-hop-self6"], _ = expandRouterBgpNeighborGroupNextHopSelf6(d, i["next_hop_self6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_capability"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["override-capability"], _ = expandRouterBgpNeighborGroupOverrideCapability(d, i["override_capability"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["passive"], _ = expandRouterBgpNeighborGroupPassive(d, i["passive"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password"], _ = expandRouterBgpNeighborGroupPassword(d, i["password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix-list-in"], _ = expandRouterBgpNeighborGroupPrefixListIn(d, i["prefix_list_in"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix-list-in-vpnv4"], _ = expandRouterBgpNeighborGroupPrefixListInVpnv4(d, i["prefix_list_in_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix-list-in-vpnv6"], _ = expandRouterBgpNeighborGroupPrefixListInVpnv6(d, i["prefix_list_in_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix-list-in6"], _ = expandRouterBgpNeighborGroupPrefixListIn6(d, i["prefix_list_in6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix-list-out"], _ = expandRouterBgpNeighborGroupPrefixListOut(d, i["prefix_list_out"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix-list-out-vpnv4"], _ = expandRouterBgpNeighborGroupPrefixListOutVpnv4(d, i["prefix_list_out_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix-list-out-vpnv6"], _ = expandRouterBgpNeighborGroupPrefixListOutVpnv6(d, i["prefix_list_out_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix-list-out6"], _ = expandRouterBgpNeighborGroupPrefixListOut6(d, i["prefix_list_out6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_as"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remote-as"], _ = expandRouterBgpNeighborGroupRemoteAs(d, i["remote_as"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_as_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remote-as-filter"], _ = expandRouterBgpNeighborGroupRemoteAsFilter(d, i["remote_as_filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remove-private-as"], _ = expandRouterBgpNeighborGroupRemovePrivateAs(d, i["remove_private_as"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as_evpn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remove-private-as-evpn"], _ = expandRouterBgpNeighborGroupRemovePrivateAsEvpn(d, i["remove_private_as_evpn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remove-private-as-vpnv4"], _ = expandRouterBgpNeighborGroupRemovePrivateAsVpnv4(d, i["remove_private_as_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remove-private-as-vpnv6"], _ = expandRouterBgpNeighborGroupRemovePrivateAsVpnv6(d, i["remove_private_as_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remove-private-as6"], _ = expandRouterBgpNeighborGroupRemovePrivateAs6(d, i["remove_private_as6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restart_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["restart-time"], _ = expandRouterBgpNeighborGroupRestartTime(d, i["restart_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retain_stale_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["retain-stale-time"], _ = expandRouterBgpNeighborGroupRetainStaleTime(d, i["retain_stale_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map-in"], _ = expandRouterBgpNeighborGroupRouteMapIn(d, i["route_map_in"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in_evpn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map-in-evpn"], _ = expandRouterBgpNeighborGroupRouteMapInEvpn(d, i["route_map_in_evpn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map-in-vpnv4"], _ = expandRouterBgpNeighborGroupRouteMapInVpnv4(d, i["route_map_in_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map-in-vpnv6"], _ = expandRouterBgpNeighborGroupRouteMapInVpnv6(d, i["route_map_in_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map-in6"], _ = expandRouterBgpNeighborGroupRouteMapIn6(d, i["route_map_in6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map-out"], _ = expandRouterBgpNeighborGroupRouteMapOut(d, i["route_map_out"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_evpn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map-out-evpn"], _ = expandRouterBgpNeighborGroupRouteMapOutEvpn(d, i["route_map_out_evpn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_preferable"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map-out-preferable"], _ = expandRouterBgpNeighborGroupRouteMapOutPreferable(d, i["route_map_out_preferable"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map-out-vpnv4"], _ = expandRouterBgpNeighborGroupRouteMapOutVpnv4(d, i["route_map_out_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv4_preferable"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map-out-vpnv4-preferable"], _ = expandRouterBgpNeighborGroupRouteMapOutVpnv4Preferable(d, i["route_map_out_vpnv4_preferable"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map-out-vpnv6"], _ = expandRouterBgpNeighborGroupRouteMapOutVpnv6(d, i["route_map_out_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv6_preferable"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map-out-vpnv6-preferable"], _ = expandRouterBgpNeighborGroupRouteMapOutVpnv6Preferable(d, i["route_map_out_vpnv6_preferable"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map-out6"], _ = expandRouterBgpNeighborGroupRouteMapOut6(d, i["route_map_out6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out6_preferable"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map-out6-preferable"], _ = expandRouterBgpNeighborGroupRouteMapOut6Preferable(d, i["route_map_out6_preferable"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-reflector-client"], _ = expandRouterBgpNeighborGroupRouteReflectorClient(d, i["route_reflector_client"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client_evpn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-reflector-client-evpn"], _ = expandRouterBgpNeighborGroupRouteReflectorClientEvpn(d, i["route_reflector_client_evpn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-reflector-client-vpnv4"], _ = expandRouterBgpNeighborGroupRouteReflectorClientVpnv4(d, i["route_reflector_client_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-reflector-client-vpnv6"], _ = expandRouterBgpNeighborGroupRouteReflectorClientVpnv6(d, i["route_reflector_client_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-reflector-client6"], _ = expandRouterBgpNeighborGroupRouteReflectorClient6(d, i["route_reflector_client6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-server-client"], _ = expandRouterBgpNeighborGroupRouteServerClient(d, i["route_server_client"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client_evpn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-server-client-evpn"], _ = expandRouterBgpNeighborGroupRouteServerClientEvpn(d, i["route_server_client_evpn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-server-client-vpnv4"], _ = expandRouterBgpNeighborGroupRouteServerClientVpnv4(d, i["route_server_client_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-server-client-vpnv6"], _ = expandRouterBgpNeighborGroupRouteServerClientVpnv6(d, i["route_server_client_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-server-client6"], _ = expandRouterBgpNeighborGroupRouteServerClient6(d, i["route_server_client6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["send-community"], _ = expandRouterBgpNeighborGroupSendCommunity(d, i["send_community"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community_evpn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["send-community-evpn"], _ = expandRouterBgpNeighborGroupSendCommunityEvpn(d, i["send_community_evpn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["send-community-vpnv4"], _ = expandRouterBgpNeighborGroupSendCommunityVpnv4(d, i["send_community_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["send-community-vpnv6"], _ = expandRouterBgpNeighborGroupSendCommunityVpnv6(d, i["send_community_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["send-community6"], _ = expandRouterBgpNeighborGroupSendCommunity6(d, i["send_community6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shutdown"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["shutdown"], _ = expandRouterBgpNeighborGroupShutdown(d, i["shutdown"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["soft-reconfiguration"], _ = expandRouterBgpNeighborGroupSoftReconfiguration(d, i["soft_reconfiguration"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration_evpn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["soft-reconfiguration-evpn"], _ = expandRouterBgpNeighborGroupSoftReconfigurationEvpn(d, i["soft_reconfiguration_evpn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration_vpnv4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["soft-reconfiguration-vpnv4"], _ = expandRouterBgpNeighborGroupSoftReconfigurationVpnv4(d, i["soft_reconfiguration_vpnv4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration_vpnv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["soft-reconfiguration-vpnv6"], _ = expandRouterBgpNeighborGroupSoftReconfigurationVpnv6(d, i["soft_reconfiguration_vpnv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["soft-reconfiguration6"], _ = expandRouterBgpNeighborGroupSoftReconfiguration6(d, i["soft_reconfiguration6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stale_route"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["stale-route"], _ = expandRouterBgpNeighborGroupStaleRoute(d, i["stale_route"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "strict_capability_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["strict-capability-match"], _ = expandRouterBgpNeighborGroupStrictCapabilityMatch(d, i["strict_capability_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unsuppress_map"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["unsuppress-map"], _ = expandRouterBgpNeighborGroupUnsuppressMap(d, i["unsuppress_map"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unsuppress_map6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["unsuppress-map6"], _ = expandRouterBgpNeighborGroupUnsuppressMap6(d, i["unsuppress_map6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_source"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["update-source"], _ = expandRouterBgpNeighborGroupUpdateSource(d, i["update_source"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandRouterBgpNeighborGroupWeight(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterBgpNeighborGroupActivate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupActivateEvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupActivateVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupActivateVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupActivate6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdditionalPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdditionalPathVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdditionalPathVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdditionalPath6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdvAdditionalPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdvAdditionalPathVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdvAdditionalPathVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdvAdditionalPath6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdvertisementInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasIn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasInEnable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasInEnableEvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasInEnableVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasInEnableVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasInEnable6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasInEvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasInVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasInVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasIn6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAsOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAsOverride6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAttributeUnchanged(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupAttributeUnchangedVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupAttributeUnchangedVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupAttributeUnchanged6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupAuthOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupBfd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityDefaultOriginate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityDefaultOriginate6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityDynamic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityGracefulRestart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityGracefulRestartEvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityGracefulRestartVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityGracefulRestartVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityGracefulRestart6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityOrf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityOrf6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityRouteRefresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupConnectTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupDefaultOriginateRoutemap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupDefaultOriginateRoutemap6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupDistributeListIn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupDistributeListInVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupDistributeListInVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupDistributeListIn6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupDistributeListOut(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupDistributeListOutVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupDistributeListOutVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupDistributeListOut6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupDontCapabilityNegotiate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupEbgpEnforceMultihop(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupEbgpMultihopTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupFilterListIn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupFilterListInVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupFilterListInVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupFilterListIn6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupFilterListOut(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupFilterListOutVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupFilterListOutVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupFilterListOut6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupHoldtimeTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupKeepAliveTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupLinkDownFailover(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupLocalAs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupLocalAsNoPrepend(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupLocalAsReplaceAs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixEvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixThresholdEvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixThresholdVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixThresholdVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixThreshold6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixWarningOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixWarningOnlyEvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixWarningOnlyVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixWarningOnlyVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixWarningOnly6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefix6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupNextHopSelf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupNextHopSelfRr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupNextHopSelfRr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupNextHopSelfVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupNextHopSelfVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupNextHopSelf6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupOverrideCapability(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupPassive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupPrefixListIn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupPrefixListInVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupPrefixListInVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupPrefixListIn6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupPrefixListOut(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupPrefixListOutVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupPrefixListOutVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupPrefixListOut6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRemoteAs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRemoteAsFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRemovePrivateAs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRemovePrivateAsEvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRemovePrivateAsVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRemovePrivateAsVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRemovePrivateAs6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRestartTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRetainStaleTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteMapIn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRouteMapInEvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRouteMapInVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRouteMapInVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRouteMapIn6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRouteMapOut(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRouteMapOutEvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRouteMapOutPreferable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRouteMapOutVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRouteMapOutVpnv4Preferable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRouteMapOutVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRouteMapOutVpnv6Preferable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRouteMapOut6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRouteMapOut6Preferable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRouteReflectorClient(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteReflectorClientEvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteReflectorClientVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteReflectorClientVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteReflectorClient6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteServerClient(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteServerClientEvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteServerClientVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteServerClientVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteServerClient6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSendCommunity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSendCommunityEvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSendCommunityVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSendCommunityVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSendCommunity6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupShutdown(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSoftReconfiguration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSoftReconfigurationEvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSoftReconfigurationVpnv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSoftReconfigurationVpnv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSoftReconfiguration6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupStaleRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupStrictCapabilityMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupUnsuppressMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupUnsuppressMap6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupUpdateSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterBgpNeighborRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_neighbor_num"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-neighbor-num"], _ = expandRouterBgpNeighborRangeMaxNeighborNum(d, i["max_neighbor_num"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbor_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["neighbor-group"], _ = expandRouterBgpNeighborRangeNeighborGroup(d, i["neighbor_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandRouterBgpNeighborRangePrefix(d, i["prefix"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterBgpNeighborRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRangeMaxNeighborNum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRangeNeighborGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRangePrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandRouterBgpNeighborRange6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterBgpNeighborRange6Id(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_neighbor_num"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-neighbor-num"], _ = expandRouterBgpNeighborRange6MaxNeighborNum(d, i["max_neighbor_num"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbor_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["neighbor-group"], _ = expandRouterBgpNeighborRange6NeighborGroup(d, i["neighbor_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix6"], _ = expandRouterBgpNeighborRange6Prefix6(d, i["prefix6"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterBgpNeighborRange6Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRange6MaxNeighborNum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRange6NeighborGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRange6Prefix6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetwork(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "backdoor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["backdoor"], _ = expandRouterBgpNetworkBackdoor(d, i["backdoor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterBgpNetworkId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "network_import_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["network-import-check"], _ = expandRouterBgpNetworkNetworkImportCheck(d, i["network_import_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandRouterBgpNetworkPrefix(d, i["prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map"], _ = expandRouterBgpNetworkRouteMap(d, i["route_map"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterBgpNetworkBackdoor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetworkId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetworkNetworkImportCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetworkPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandRouterBgpNetworkRouteMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNetworkImportCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetwork6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "backdoor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["backdoor"], _ = expandRouterBgpNetwork6Backdoor(d, i["backdoor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterBgpNetwork6Id(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "network_import_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["network-import-check"], _ = expandRouterBgpNetwork6NetworkImportCheck(d, i["network_import_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix6"], _ = expandRouterBgpNetwork6Prefix6(d, i["prefix6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map"], _ = expandRouterBgpNetwork6RouteMap(d, i["route_map"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterBgpNetwork6Backdoor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetwork6Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetwork6NetworkImportCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetwork6Prefix6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetwork6RouteMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpRecursiveInheritPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpRecursiveNextHop(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpRedistribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandRouterBgpRedistributeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map"], _ = expandRouterBgpRedistributeRouteMap(d, i["route_map"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandRouterBgpRedistributeStatus(d, i["status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterBgpRedistributeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpRedistributeRouteMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpRedistributeStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpRedistribute6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandRouterBgpRedistribute6Name(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map"], _ = expandRouterBgpRedistribute6RouteMap(d, i["route_map"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandRouterBgpRedistribute6Status(d, i["status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterBgpRedistribute6Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpRedistribute6RouteMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpRedistribute6Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpRouterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpScanTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpSynchronization(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpTagResolveMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfLeak(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandRouterBgpVrfLeakTarget(d, i["target"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["target"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrf"], _ = expandRouterBgpVrfLeakVrf(d, i["vrf"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterBgpVrfLeakTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["interface"], _ = expandRouterBgpVrfLeakTargetInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map"], _ = expandRouterBgpVrfLeakTargetRouteMap(d, i["route_map"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrf"], _ = expandRouterBgpVrfLeakTargetVrf(d, i["vrf"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterBgpVrfLeakTargetInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpVrfLeakTargetRouteMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpVrfLeakTargetVrf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfLeakVrf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfLeak6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandRouterBgpVrfLeak6Target(d, i["target"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["target"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrf"], _ = expandRouterBgpVrfLeak6Vrf(d, i["vrf"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterBgpVrfLeak6Target(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["interface"], _ = expandRouterBgpVrfLeak6TargetInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map"], _ = expandRouterBgpVrfLeak6TargetRouteMap(d, i["route_map"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrf"], _ = expandRouterBgpVrfLeak6TargetVrf(d, i["vrf"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterBgpVrfLeak6TargetInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpVrfLeak6TargetRouteMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpVrfLeak6TargetVrf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfLeak6Vrf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "export_rt"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["export-rt"], _ = expandRouterBgpVrfExportRt(d, i["export_rt"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "import_route_map"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["import-route-map"], _ = expandRouterBgpVrfImportRouteMap(d, i["import_route_map"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "import_rt"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["import-rt"], _ = expandRouterBgpVrfImportRt(d, i["import_rt"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "leak_target"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandRouterBgpVrfLeakTargetU(d, i["leak_target"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["leak-target"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rd"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rd"], _ = expandRouterBgpVrfRd(d, i["rd"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["role"], _ = expandRouterBgpVrfRole(d, i["role"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrf"], _ = expandRouterBgpVrfVrf(d, i["vrf"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterBgpVrfExportRt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpVrfImportRouteMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpVrfImportRt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpVrfLeakTargetU(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["interface"], _ = expandRouterBgpVrfLeakTargetUInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map"], _ = expandRouterBgpVrfLeakTargetURouteMap(d, i["route_map"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrf"], _ = expandRouterBgpVrfLeakTargetUVrf(d, i["vrf"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterBgpVrfLeakTargetUInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpVrfLeakTargetURouteMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpVrfLeakTargetUVrf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfRd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfVrf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrf6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "export_rt"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["export-rt"], _ = expandRouterBgpVrf6ExportRt(d, i["export_rt"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "import_route_map"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["import-route-map"], _ = expandRouterBgpVrf6ImportRouteMap(d, i["import_route_map"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "import_rt"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["import-rt"], _ = expandRouterBgpVrf6ImportRt(d, i["import_rt"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "leak_target"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandRouterBgpVrf6LeakTarget(d, i["leak_target"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["leak-target"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rd"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rd"], _ = expandRouterBgpVrf6Rd(d, i["rd"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["role"], _ = expandRouterBgpVrf6Role(d, i["role"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrf"], _ = expandRouterBgpVrf6Vrf(d, i["vrf"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterBgpVrf6ExportRt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpVrf6ImportRouteMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpVrf6ImportRt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpVrf6LeakTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["interface"], _ = expandRouterBgpVrf6LeakTargetInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map"], _ = expandRouterBgpVrf6LeakTargetRouteMap(d, i["route_map"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrf"], _ = expandRouterBgpVrf6LeakTargetVrf(d, i["vrf"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterBgpVrf6LeakTargetInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpVrf6LeakTargetRouteMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpVrf6LeakTargetVrf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrf6Rd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrf6Role(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrf6Vrf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterBgp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("additional_path"); ok || d.HasChange("additional_path") {
		t, err := expandRouterBgpAdditionalPath(d, v, "additional_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-path"] = t
		}
	}

	if v, ok := d.GetOk("additional_path_select"); ok || d.HasChange("additional_path_select") {
		t, err := expandRouterBgpAdditionalPathSelect(d, v, "additional_path_select")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-path-select"] = t
		}
	}

	if v, ok := d.GetOk("additional_path_select_vpnv4"); ok || d.HasChange("additional_path_select_vpnv4") {
		t, err := expandRouterBgpAdditionalPathSelectVpnv4(d, v, "additional_path_select_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-path-select-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("additional_path_select_vpnv6"); ok || d.HasChange("additional_path_select_vpnv6") {
		t, err := expandRouterBgpAdditionalPathSelectVpnv6(d, v, "additional_path_select_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-path-select-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("additional_path_select6"); ok || d.HasChange("additional_path_select6") {
		t, err := expandRouterBgpAdditionalPathSelect6(d, v, "additional_path_select6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-path-select6"] = t
		}
	}

	if v, ok := d.GetOk("additional_path_vpnv4"); ok || d.HasChange("additional_path_vpnv4") {
		t, err := expandRouterBgpAdditionalPathVpnv4(d, v, "additional_path_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-path-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("additional_path_vpnv6"); ok || d.HasChange("additional_path_vpnv6") {
		t, err := expandRouterBgpAdditionalPathVpnv6(d, v, "additional_path_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-path-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("additional_path6"); ok || d.HasChange("additional_path6") {
		t, err := expandRouterBgpAdditionalPath6(d, v, "additional_path6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-path6"] = t
		}
	}

	if v, ok := d.GetOk("admin_distance"); ok || d.HasChange("admin_distance") {
		t, err := expandRouterBgpAdminDistance(d, v, "admin_distance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-distance"] = t
		}
	}

	if v, ok := d.GetOk("aggregate_address"); ok || d.HasChange("aggregate_address") {
		t, err := expandRouterBgpAggregateAddress(d, v, "aggregate_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aggregate-address"] = t
		}
	}

	if v, ok := d.GetOk("aggregate_address6"); ok || d.HasChange("aggregate_address6") {
		t, err := expandRouterBgpAggregateAddress6(d, v, "aggregate_address6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aggregate-address6"] = t
		}
	}

	if v, ok := d.GetOk("always_compare_med"); ok || d.HasChange("always_compare_med") {
		t, err := expandRouterBgpAlwaysCompareMed(d, v, "always_compare_med")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["always-compare-med"] = t
		}
	}

	if v, ok := d.GetOk("as"); ok || d.HasChange("as") {
		t, err := expandRouterBgpAs(d, v, "as")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["as"] = t
		}
	}

	if v, ok := d.GetOk("bestpath_as_path_ignore"); ok || d.HasChange("bestpath_as_path_ignore") {
		t, err := expandRouterBgpBestpathAsPathIgnore(d, v, "bestpath_as_path_ignore")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bestpath-as-path-ignore"] = t
		}
	}

	if v, ok := d.GetOk("bestpath_cmp_confed_aspath"); ok || d.HasChange("bestpath_cmp_confed_aspath") {
		t, err := expandRouterBgpBestpathCmpConfedAspath(d, v, "bestpath_cmp_confed_aspath")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bestpath-cmp-confed-aspath"] = t
		}
	}

	if v, ok := d.GetOk("bestpath_cmp_routerid"); ok || d.HasChange("bestpath_cmp_routerid") {
		t, err := expandRouterBgpBestpathCmpRouterid(d, v, "bestpath_cmp_routerid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bestpath-cmp-routerid"] = t
		}
	}

	if v, ok := d.GetOk("bestpath_med_confed"); ok || d.HasChange("bestpath_med_confed") {
		t, err := expandRouterBgpBestpathMedConfed(d, v, "bestpath_med_confed")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bestpath-med-confed"] = t
		}
	}

	if v, ok := d.GetOk("bestpath_med_missing_as_worst"); ok || d.HasChange("bestpath_med_missing_as_worst") {
		t, err := expandRouterBgpBestpathMedMissingAsWorst(d, v, "bestpath_med_missing_as_worst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bestpath-med-missing-as-worst"] = t
		}
	}

	if v, ok := d.GetOk("client_to_client_reflection"); ok || d.HasChange("client_to_client_reflection") {
		t, err := expandRouterBgpClientToClientReflection(d, v, "client_to_client_reflection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-to-client-reflection"] = t
		}
	}

	if v, ok := d.GetOk("cluster_id"); ok || d.HasChange("cluster_id") {
		t, err := expandRouterBgpClusterId(d, v, "cluster_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cluster-id"] = t
		}
	}

	if v, ok := d.GetOk("confederation_identifier"); ok || d.HasChange("confederation_identifier") {
		t, err := expandRouterBgpConfederationIdentifier(d, v, "confederation_identifier")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["confederation-identifier"] = t
		}
	}

	if v, ok := d.GetOk("confederation_peers"); ok || d.HasChange("confederation_peers") {
		t, err := expandRouterBgpConfederationPeers(d, v, "confederation_peers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["confederation-peers"] = t
		}
	}

	if v, ok := d.GetOk("cross_family_conditional_adv"); ok || d.HasChange("cross_family_conditional_adv") {
		t, err := expandRouterBgpCrossFamilyConditionalAdv(d, v, "cross_family_conditional_adv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cross-family-conditional-adv"] = t
		}
	}

	if v, ok := d.GetOk("dampening"); ok || d.HasChange("dampening") {
		t, err := expandRouterBgpDampening(d, v, "dampening")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dampening"] = t
		}
	}

	if v, ok := d.GetOk("dampening_max_suppress_time"); ok || d.HasChange("dampening_max_suppress_time") {
		t, err := expandRouterBgpDampeningMaxSuppressTime(d, v, "dampening_max_suppress_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dampening-max-suppress-time"] = t
		}
	}

	if v, ok := d.GetOk("dampening_reachability_half_life"); ok || d.HasChange("dampening_reachability_half_life") {
		t, err := expandRouterBgpDampeningReachabilityHalfLife(d, v, "dampening_reachability_half_life")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dampening-reachability-half-life"] = t
		}
	}

	if v, ok := d.GetOk("dampening_reuse"); ok || d.HasChange("dampening_reuse") {
		t, err := expandRouterBgpDampeningReuse(d, v, "dampening_reuse")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dampening-reuse"] = t
		}
	}

	if v, ok := d.GetOk("dampening_route_map"); ok || d.HasChange("dampening_route_map") {
		t, err := expandRouterBgpDampeningRouteMap(d, v, "dampening_route_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dampening-route-map"] = t
		}
	}

	if v, ok := d.GetOk("dampening_suppress"); ok || d.HasChange("dampening_suppress") {
		t, err := expandRouterBgpDampeningSuppress(d, v, "dampening_suppress")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dampening-suppress"] = t
		}
	}

	if v, ok := d.GetOk("dampening_unreachability_half_life"); ok || d.HasChange("dampening_unreachability_half_life") {
		t, err := expandRouterBgpDampeningUnreachabilityHalfLife(d, v, "dampening_unreachability_half_life")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dampening-unreachability-half-life"] = t
		}
	}

	if v, ok := d.GetOk("default_local_preference"); ok || d.HasChange("default_local_preference") {
		t, err := expandRouterBgpDefaultLocalPreference(d, v, "default_local_preference")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-local-preference"] = t
		}
	}

	if v, ok := d.GetOk("deterministic_med"); ok || d.HasChange("deterministic_med") {
		t, err := expandRouterBgpDeterministicMed(d, v, "deterministic_med")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deterministic-med"] = t
		}
	}

	if v, ok := d.GetOk("distance_external"); ok || d.HasChange("distance_external") {
		t, err := expandRouterBgpDistanceExternal(d, v, "distance_external")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance-external"] = t
		}
	}

	if v, ok := d.GetOk("distance_internal"); ok || d.HasChange("distance_internal") {
		t, err := expandRouterBgpDistanceInternal(d, v, "distance_internal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance-internal"] = t
		}
	}

	if v, ok := d.GetOk("distance_local"); ok || d.HasChange("distance_local") {
		t, err := expandRouterBgpDistanceLocal(d, v, "distance_local")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance-local"] = t
		}
	}

	if v, ok := d.GetOk("ebgp_multipath"); ok || d.HasChange("ebgp_multipath") {
		t, err := expandRouterBgpEbgpMultipath(d, v, "ebgp_multipath")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ebgp-multipath"] = t
		}
	}

	if v, ok := d.GetOk("enforce_first_as"); ok || d.HasChange("enforce_first_as") {
		t, err := expandRouterBgpEnforceFirstAs(d, v, "enforce_first_as")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enforce-first-as"] = t
		}
	}

	if v, ok := d.GetOk("fast_external_failover"); ok || d.HasChange("fast_external_failover") {
		t, err := expandRouterBgpFastExternalFailover(d, v, "fast_external_failover")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fast-external-failover"] = t
		}
	}

	if v, ok := d.GetOk("graceful_end_on_timer"); ok || d.HasChange("graceful_end_on_timer") {
		t, err := expandRouterBgpGracefulEndOnTimer(d, v, "graceful_end_on_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["graceful-end-on-timer"] = t
		}
	}

	if v, ok := d.GetOk("graceful_restart"); ok || d.HasChange("graceful_restart") {
		t, err := expandRouterBgpGracefulRestart(d, v, "graceful_restart")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["graceful-restart"] = t
		}
	}

	if v, ok := d.GetOk("graceful_restart_time"); ok || d.HasChange("graceful_restart_time") {
		t, err := expandRouterBgpGracefulRestartTime(d, v, "graceful_restart_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["graceful-restart-time"] = t
		}
	}

	if v, ok := d.GetOk("graceful_stalepath_time"); ok || d.HasChange("graceful_stalepath_time") {
		t, err := expandRouterBgpGracefulStalepathTime(d, v, "graceful_stalepath_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["graceful-stalepath-time"] = t
		}
	}

	if v, ok := d.GetOk("graceful_update_delay"); ok || d.HasChange("graceful_update_delay") {
		t, err := expandRouterBgpGracefulUpdateDelay(d, v, "graceful_update_delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["graceful-update-delay"] = t
		}
	}

	if v, ok := d.GetOk("holdtime_timer"); ok || d.HasChange("holdtime_timer") {
		t, err := expandRouterBgpHoldtimeTimer(d, v, "holdtime_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["holdtime-timer"] = t
		}
	}

	if v, ok := d.GetOk("ibgp_multipath"); ok || d.HasChange("ibgp_multipath") {
		t, err := expandRouterBgpIbgpMultipath(d, v, "ibgp_multipath")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ibgp-multipath"] = t
		}
	}

	if v, ok := d.GetOk("ignore_optional_capability"); ok || d.HasChange("ignore_optional_capability") {
		t, err := expandRouterBgpIgnoreOptionalCapability(d, v, "ignore_optional_capability")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ignore-optional-capability"] = t
		}
	}

	if v, ok := d.GetOk("keepalive_timer"); ok || d.HasChange("keepalive_timer") {
		t, err := expandRouterBgpKeepaliveTimer(d, v, "keepalive_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keepalive-timer"] = t
		}
	}

	if v, ok := d.GetOk("log_neighbour_changes"); ok || d.HasChange("log_neighbour_changes") {
		t, err := expandRouterBgpLogNeighbourChanges(d, v, "log_neighbour_changes")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-neighbour-changes"] = t
		}
	}

	if v, ok := d.GetOk("multipath_recursive_distance"); ok || d.HasChange("multipath_recursive_distance") {
		t, err := expandRouterBgpMultipathRecursiveDistance(d, v, "multipath_recursive_distance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multipath-recursive-distance"] = t
		}
	}

	if v, ok := d.GetOk("neighbor"); ok || d.HasChange("neighbor") {
		t, err := expandRouterBgpNeighbor(d, v, "neighbor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor"] = t
		}
	}

	if v, ok := d.GetOk("neighbor_group"); ok || d.HasChange("neighbor_group") {
		t, err := expandRouterBgpNeighborGroup(d, v, "neighbor_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-group"] = t
		}
	}

	if v, ok := d.GetOk("neighbor_range"); ok || d.HasChange("neighbor_range") {
		t, err := expandRouterBgpNeighborRange(d, v, "neighbor_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-range"] = t
		}
	}

	if v, ok := d.GetOk("neighbor_range6"); ok || d.HasChange("neighbor_range6") {
		t, err := expandRouterBgpNeighborRange6(d, v, "neighbor_range6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-range6"] = t
		}
	}

	if v, ok := d.GetOk("network"); ok || d.HasChange("network") {
		t, err := expandRouterBgpNetwork(d, v, "network")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network"] = t
		}
	}

	if v, ok := d.GetOk("network_import_check"); ok || d.HasChange("network_import_check") {
		t, err := expandRouterBgpNetworkImportCheck(d, v, "network_import_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-import-check"] = t
		}
	}

	if v, ok := d.GetOk("network6"); ok || d.HasChange("network6") {
		t, err := expandRouterBgpNetwork6(d, v, "network6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network6"] = t
		}
	}

	if v, ok := d.GetOk("recursive_inherit_priority"); ok || d.HasChange("recursive_inherit_priority") {
		t, err := expandRouterBgpRecursiveInheritPriority(d, v, "recursive_inherit_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["recursive-inherit-priority"] = t
		}
	}

	if v, ok := d.GetOk("recursive_next_hop"); ok || d.HasChange("recursive_next_hop") {
		t, err := expandRouterBgpRecursiveNextHop(d, v, "recursive_next_hop")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["recursive-next-hop"] = t
		}
	}

	if v, ok := d.GetOk("redistribute"); ok || d.HasChange("redistribute") {
		t, err := expandRouterBgpRedistribute(d, v, "redistribute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redistribute"] = t
		}
	}

	if v, ok := d.GetOk("redistribute6"); ok || d.HasChange("redistribute6") {
		t, err := expandRouterBgpRedistribute6(d, v, "redistribute6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redistribute6"] = t
		}
	}

	if v, ok := d.GetOk("router_id"); ok || d.HasChange("router_id") {
		t, err := expandRouterBgpRouterId(d, v, "router_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["router-id"] = t
		}
	}

	if v, ok := d.GetOk("scan_time"); ok || d.HasChange("scan_time") {
		t, err := expandRouterBgpScanTime(d, v, "scan_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-time"] = t
		}
	}

	if v, ok := d.GetOk("synchronization"); ok || d.HasChange("synchronization") {
		t, err := expandRouterBgpSynchronization(d, v, "synchronization")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["synchronization"] = t
		}
	}

	if v, ok := d.GetOk("tag_resolve_mode"); ok || d.HasChange("tag_resolve_mode") {
		t, err := expandRouterBgpTagResolveMode(d, v, "tag_resolve_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tag-resolve-mode"] = t
		}
	}

	if v, ok := d.GetOk("vrf_leak"); ok || d.HasChange("vrf_leak") {
		t, err := expandRouterBgpVrfLeak(d, v, "vrf_leak")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf-leak"] = t
		}
	}

	if v, ok := d.GetOk("vrf_leak6"); ok || d.HasChange("vrf_leak6") {
		t, err := expandRouterBgpVrfLeak6(d, v, "vrf_leak6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf-leak6"] = t
		}
	}

	if v, ok := d.GetOk("vrf"); ok || d.HasChange("vrf") {
		t, err := expandRouterBgpVrf(d, v, "vrf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf"] = t
		}
	}

	if v, ok := d.GetOk("vrf6"); ok || d.HasChange("vrf6") {
		t, err := expandRouterBgpVrf6(d, v, "vrf6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf6"] = t
		}
	}

	return &obj, nil
}
