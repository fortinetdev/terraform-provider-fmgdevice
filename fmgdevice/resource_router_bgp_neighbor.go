// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: BGP neighbor table.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterBgpNeighbor() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterBgpNeighborCreate,
		Read:   resourceRouterBgpNeighborRead,
		Update: resourceRouterBgpNeighborUpdate,
		Delete: resourceRouterBgpNeighborDelete,

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
				ForceNew: true,
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
			"rr_attr_allow_change": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rr_attr_allow_change_evpn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rr_attr_allow_change_vpnv4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rr_attr_allow_change_vpnv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rr_attr_allow_change6": &schema.Schema{
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceRouterBgpNeighborCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterBgpNeighbor(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterBgpNeighbor resource while getting object: %v", err)
	}

	_, err = c.CreateRouterBgpNeighbor(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating RouterBgpNeighbor resource: %v", err)
	}

	d.SetId(getStringKey(d, "ip"))

	return resourceRouterBgpNeighborRead(d, m)
}

func resourceRouterBgpNeighborUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterBgpNeighbor(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgpNeighbor resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterBgpNeighbor(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgpNeighbor resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "ip"))

	return resourceRouterBgpNeighborRead(d, m)
}

func resourceRouterBgpNeighborDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterBgpNeighbor(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterBgpNeighbor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterBgpNeighborRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterBgpNeighbor(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgpNeighbor resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterBgpNeighbor(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgpNeighbor resource from API: %v", err)
	}
	return nil
}

func flattenRouterBgpNeighborActivate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborActivateEvpn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborActivateVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborActivateVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborActivate62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAdditionalPath2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAdditionalPathVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAdditionalPathVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAdditionalPath62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAdvAdditionalPath2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAdvAdditionalPathVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAdvAdditionalPathVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAdvAdditionalPath62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAdvertisementInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAllowasIn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAllowasInEnable2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAllowasInEnableEvpn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAllowasInEnableVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAllowasInEnableVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAllowasInEnable62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAllowasInEvpn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAllowasInVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAllowasInVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAllowasIn62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAsOverride2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAsOverride62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborAttributeUnchanged2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborAttributeUnchangedVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborAttributeUnchangedVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborAttributeUnchanged62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborAuthOptions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborBfd2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityDefaultOriginate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityDefaultOriginate62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityDynamic2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityGracefulRestart2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityGracefulRestartEvpn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityGracefulRestartVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityGracefulRestartVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityGracefulRestart62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityOrf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityOrf62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityRouteRefresh2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborConditionalAdvertise2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterBgpNeighborConditionalAdvertiseAdvertiseRoutemap2edl(i["advertise-routemap"], d, pre_append)
			tmp["advertise_routemap"] = fortiAPISubPartPatch(v, "RouterBgpNeighbor-ConditionalAdvertise-AdvertiseRoutemap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_routemap"
		if _, ok := i["condition-routemap"]; ok {
			v := flattenRouterBgpNeighborConditionalAdvertiseConditionRoutemap2edl(i["condition-routemap"], d, pre_append)
			tmp["condition_routemap"] = fortiAPISubPartPatch(v, "RouterBgpNeighbor-ConditionalAdvertise-ConditionRoutemap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_type"
		if _, ok := i["condition-type"]; ok {
			v := flattenRouterBgpNeighborConditionalAdvertiseConditionType2edl(i["condition-type"], d, pre_append)
			tmp["condition_type"] = fortiAPISubPartPatch(v, "RouterBgpNeighbor-ConditionalAdvertise-ConditionType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterBgpNeighborConditionalAdvertiseAdvertiseRoutemap2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenRouterBgpNeighborConditionalAdvertiseConditionRoutemap2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborConditionalAdvertiseConditionType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborConditionalAdvertise62edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterBgpNeighborConditionalAdvertise6AdvertiseRoutemap2edl(i["advertise-routemap"], d, pre_append)
			tmp["advertise_routemap"] = fortiAPISubPartPatch(v, "RouterBgpNeighbor-ConditionalAdvertise6-AdvertiseRoutemap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_routemap"
		if _, ok := i["condition-routemap"]; ok {
			v := flattenRouterBgpNeighborConditionalAdvertise6ConditionRoutemap2edl(i["condition-routemap"], d, pre_append)
			tmp["condition_routemap"] = fortiAPISubPartPatch(v, "RouterBgpNeighbor-ConditionalAdvertise6-ConditionRoutemap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_type"
		if _, ok := i["condition-type"]; ok {
			v := flattenRouterBgpNeighborConditionalAdvertise6ConditionType2edl(i["condition-type"], d, pre_append)
			tmp["condition_type"] = fortiAPISubPartPatch(v, "RouterBgpNeighbor-ConditionalAdvertise6-ConditionType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterBgpNeighborConditionalAdvertise6AdvertiseRoutemap2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborConditionalAdvertise6ConditionRoutemap2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborConditionalAdvertise6ConditionType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborConnectTimer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborDefaultOriginateRoutemap2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborDefaultOriginateRoutemap62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborDistributeListIn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborDistributeListInVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborDistributeListInVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborDistributeListIn62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborDistributeListOut2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborDistributeListOutVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborDistributeListOutVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborDistributeListOut62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborDontCapabilityNegotiate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborEbgpEnforceMultihop2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborEbgpMultihopTtl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborFilterListIn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborFilterListInVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborFilterListInVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborFilterListIn62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborFilterListOut2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborFilterListOutVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborFilterListOutVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborFilterListOut62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborHoldtimeTimer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborKeepAliveTimer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborLinkDownFailover2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborLocalAs2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborLocalAsNoPrepend2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborLocalAsReplaceAs2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixEvpn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixThresholdEvpn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixThresholdVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixThresholdVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixThreshold62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixWarningOnly2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixWarningOnlyEvpn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixWarningOnlyVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixWarningOnlyVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixWarningOnly62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefix62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborNextHopSelf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborNextHopSelfRr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborNextHopSelfRr62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborNextHopSelfVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborNextHopSelfVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborNextHopSelf62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborOverrideCapability2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborPassive2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborPrefixListIn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborPrefixListInVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborPrefixListInVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborPrefixListIn62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborPrefixListOut2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborPrefixListOutVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborPrefixListOutVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborPrefixListOut62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRemoteAs2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRemovePrivateAs2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRemovePrivateAsEvpn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRemovePrivateAsVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRemovePrivateAsVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRemovePrivateAs62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRestartTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRetainStaleTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteMapIn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRouteMapInEvpn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRouteMapInVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRouteMapInVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRouteMapIn62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRouteMapOut2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRouteMapOutEvpn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRouteMapOutPreferable2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRouteMapOutVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRouteMapOutVpnv4Preferable2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRouteMapOutVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRouteMapOutVpnv6Preferable2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRouteMapOut62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRouteMapOut6Preferable2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRouteReflectorClient2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteReflectorClientEvpn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteReflectorClientVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteReflectorClientVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteReflectorClient62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteServerClient2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteServerClientEvpn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteServerClientVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteServerClientVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteServerClient62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRrAttrAllowChange2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRrAttrAllowChangeEvpn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRrAttrAllowChangeVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRrAttrAllowChangeVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRrAttrAllowChange62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborSendCommunity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborSendCommunityEvpn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborSendCommunityVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborSendCommunityVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborSendCommunity62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborShutdown2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborSoftReconfiguration2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborSoftReconfigurationEvpn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborSoftReconfigurationVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborSoftReconfigurationVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborSoftReconfiguration62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborStaleRoute2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborStrictCapabilityMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborUnsuppressMap2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborUnsuppressMap62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborUpdateSource2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterBgpNeighbor(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("activate", flattenRouterBgpNeighborActivate2edl(o["activate"], d, "activate")); err != nil {
		if vv, ok := fortiAPIPatch(o["activate"], "RouterBgpNeighbor-Activate"); ok {
			if err = d.Set("activate", vv); err != nil {
				return fmt.Errorf("Error reading activate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading activate: %v", err)
		}
	}

	if err = d.Set("activate_evpn", flattenRouterBgpNeighborActivateEvpn2edl(o["activate-evpn"], d, "activate_evpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["activate-evpn"], "RouterBgpNeighbor-ActivateEvpn"); ok {
			if err = d.Set("activate_evpn", vv); err != nil {
				return fmt.Errorf("Error reading activate_evpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading activate_evpn: %v", err)
		}
	}

	if err = d.Set("activate_vpnv4", flattenRouterBgpNeighborActivateVpnv42edl(o["activate-vpnv4"], d, "activate_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["activate-vpnv4"], "RouterBgpNeighbor-ActivateVpnv4"); ok {
			if err = d.Set("activate_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading activate_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading activate_vpnv4: %v", err)
		}
	}

	if err = d.Set("activate_vpnv6", flattenRouterBgpNeighborActivateVpnv62edl(o["activate-vpnv6"], d, "activate_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["activate-vpnv6"], "RouterBgpNeighbor-ActivateVpnv6"); ok {
			if err = d.Set("activate_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading activate_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading activate_vpnv6: %v", err)
		}
	}

	if err = d.Set("activate6", flattenRouterBgpNeighborActivate62edl(o["activate6"], d, "activate6")); err != nil {
		if vv, ok := fortiAPIPatch(o["activate6"], "RouterBgpNeighbor-Activate6"); ok {
			if err = d.Set("activate6", vv); err != nil {
				return fmt.Errorf("Error reading activate6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading activate6: %v", err)
		}
	}

	if err = d.Set("additional_path", flattenRouterBgpNeighborAdditionalPath2edl(o["additional-path"], d, "additional_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["additional-path"], "RouterBgpNeighbor-AdditionalPath"); ok {
			if err = d.Set("additional_path", vv); err != nil {
				return fmt.Errorf("Error reading additional_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading additional_path: %v", err)
		}
	}

	if err = d.Set("additional_path_vpnv4", flattenRouterBgpNeighborAdditionalPathVpnv42edl(o["additional-path-vpnv4"], d, "additional_path_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["additional-path-vpnv4"], "RouterBgpNeighbor-AdditionalPathVpnv4"); ok {
			if err = d.Set("additional_path_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading additional_path_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading additional_path_vpnv4: %v", err)
		}
	}

	if err = d.Set("additional_path_vpnv6", flattenRouterBgpNeighborAdditionalPathVpnv62edl(o["additional-path-vpnv6"], d, "additional_path_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["additional-path-vpnv6"], "RouterBgpNeighbor-AdditionalPathVpnv6"); ok {
			if err = d.Set("additional_path_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading additional_path_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading additional_path_vpnv6: %v", err)
		}
	}

	if err = d.Set("additional_path6", flattenRouterBgpNeighborAdditionalPath62edl(o["additional-path6"], d, "additional_path6")); err != nil {
		if vv, ok := fortiAPIPatch(o["additional-path6"], "RouterBgpNeighbor-AdditionalPath6"); ok {
			if err = d.Set("additional_path6", vv); err != nil {
				return fmt.Errorf("Error reading additional_path6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading additional_path6: %v", err)
		}
	}

	if err = d.Set("adv_additional_path", flattenRouterBgpNeighborAdvAdditionalPath2edl(o["adv-additional-path"], d, "adv_additional_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["adv-additional-path"], "RouterBgpNeighbor-AdvAdditionalPath"); ok {
			if err = d.Set("adv_additional_path", vv); err != nil {
				return fmt.Errorf("Error reading adv_additional_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adv_additional_path: %v", err)
		}
	}

	if err = d.Set("adv_additional_path_vpnv4", flattenRouterBgpNeighborAdvAdditionalPathVpnv42edl(o["adv-additional-path-vpnv4"], d, "adv_additional_path_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["adv-additional-path-vpnv4"], "RouterBgpNeighbor-AdvAdditionalPathVpnv4"); ok {
			if err = d.Set("adv_additional_path_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading adv_additional_path_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adv_additional_path_vpnv4: %v", err)
		}
	}

	if err = d.Set("adv_additional_path_vpnv6", flattenRouterBgpNeighborAdvAdditionalPathVpnv62edl(o["adv-additional-path-vpnv6"], d, "adv_additional_path_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["adv-additional-path-vpnv6"], "RouterBgpNeighbor-AdvAdditionalPathVpnv6"); ok {
			if err = d.Set("adv_additional_path_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading adv_additional_path_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adv_additional_path_vpnv6: %v", err)
		}
	}

	if err = d.Set("adv_additional_path6", flattenRouterBgpNeighborAdvAdditionalPath62edl(o["adv-additional-path6"], d, "adv_additional_path6")); err != nil {
		if vv, ok := fortiAPIPatch(o["adv-additional-path6"], "RouterBgpNeighbor-AdvAdditionalPath6"); ok {
			if err = d.Set("adv_additional_path6", vv); err != nil {
				return fmt.Errorf("Error reading adv_additional_path6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adv_additional_path6: %v", err)
		}
	}

	if err = d.Set("advertisement_interval", flattenRouterBgpNeighborAdvertisementInterval2edl(o["advertisement-interval"], d, "advertisement_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["advertisement-interval"], "RouterBgpNeighbor-AdvertisementInterval"); ok {
			if err = d.Set("advertisement_interval", vv); err != nil {
				return fmt.Errorf("Error reading advertisement_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading advertisement_interval: %v", err)
		}
	}

	if err = d.Set("allowas_in", flattenRouterBgpNeighborAllowasIn2edl(o["allowas-in"], d, "allowas_in")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowas-in"], "RouterBgpNeighbor-AllowasIn"); ok {
			if err = d.Set("allowas_in", vv); err != nil {
				return fmt.Errorf("Error reading allowas_in: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowas_in: %v", err)
		}
	}

	if err = d.Set("allowas_in_enable", flattenRouterBgpNeighborAllowasInEnable2edl(o["allowas-in-enable"], d, "allowas_in_enable")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowas-in-enable"], "RouterBgpNeighbor-AllowasInEnable"); ok {
			if err = d.Set("allowas_in_enable", vv); err != nil {
				return fmt.Errorf("Error reading allowas_in_enable: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowas_in_enable: %v", err)
		}
	}

	if err = d.Set("allowas_in_enable_evpn", flattenRouterBgpNeighborAllowasInEnableEvpn2edl(o["allowas-in-enable-evpn"], d, "allowas_in_enable_evpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowas-in-enable-evpn"], "RouterBgpNeighbor-AllowasInEnableEvpn"); ok {
			if err = d.Set("allowas_in_enable_evpn", vv); err != nil {
				return fmt.Errorf("Error reading allowas_in_enable_evpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowas_in_enable_evpn: %v", err)
		}
	}

	if err = d.Set("allowas_in_enable_vpnv4", flattenRouterBgpNeighborAllowasInEnableVpnv42edl(o["allowas-in-enable-vpnv4"], d, "allowas_in_enable_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowas-in-enable-vpnv4"], "RouterBgpNeighbor-AllowasInEnableVpnv4"); ok {
			if err = d.Set("allowas_in_enable_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading allowas_in_enable_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowas_in_enable_vpnv4: %v", err)
		}
	}

	if err = d.Set("allowas_in_enable_vpnv6", flattenRouterBgpNeighborAllowasInEnableVpnv62edl(o["allowas-in-enable-vpnv6"], d, "allowas_in_enable_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowas-in-enable-vpnv6"], "RouterBgpNeighbor-AllowasInEnableVpnv6"); ok {
			if err = d.Set("allowas_in_enable_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading allowas_in_enable_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowas_in_enable_vpnv6: %v", err)
		}
	}

	if err = d.Set("allowas_in_enable6", flattenRouterBgpNeighborAllowasInEnable62edl(o["allowas-in-enable6"], d, "allowas_in_enable6")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowas-in-enable6"], "RouterBgpNeighbor-AllowasInEnable6"); ok {
			if err = d.Set("allowas_in_enable6", vv); err != nil {
				return fmt.Errorf("Error reading allowas_in_enable6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowas_in_enable6: %v", err)
		}
	}

	if err = d.Set("allowas_in_evpn", flattenRouterBgpNeighborAllowasInEvpn2edl(o["allowas-in-evpn"], d, "allowas_in_evpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowas-in-evpn"], "RouterBgpNeighbor-AllowasInEvpn"); ok {
			if err = d.Set("allowas_in_evpn", vv); err != nil {
				return fmt.Errorf("Error reading allowas_in_evpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowas_in_evpn: %v", err)
		}
	}

	if err = d.Set("allowas_in_vpnv4", flattenRouterBgpNeighborAllowasInVpnv42edl(o["allowas-in-vpnv4"], d, "allowas_in_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowas-in-vpnv4"], "RouterBgpNeighbor-AllowasInVpnv4"); ok {
			if err = d.Set("allowas_in_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading allowas_in_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowas_in_vpnv4: %v", err)
		}
	}

	if err = d.Set("allowas_in_vpnv6", flattenRouterBgpNeighborAllowasInVpnv62edl(o["allowas-in-vpnv6"], d, "allowas_in_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowas-in-vpnv6"], "RouterBgpNeighbor-AllowasInVpnv6"); ok {
			if err = d.Set("allowas_in_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading allowas_in_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowas_in_vpnv6: %v", err)
		}
	}

	if err = d.Set("allowas_in6", flattenRouterBgpNeighborAllowasIn62edl(o["allowas-in6"], d, "allowas_in6")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowas-in6"], "RouterBgpNeighbor-AllowasIn6"); ok {
			if err = d.Set("allowas_in6", vv); err != nil {
				return fmt.Errorf("Error reading allowas_in6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowas_in6: %v", err)
		}
	}

	if err = d.Set("as_override", flattenRouterBgpNeighborAsOverride2edl(o["as-override"], d, "as_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["as-override"], "RouterBgpNeighbor-AsOverride"); ok {
			if err = d.Set("as_override", vv); err != nil {
				return fmt.Errorf("Error reading as_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading as_override: %v", err)
		}
	}

	if err = d.Set("as_override6", flattenRouterBgpNeighborAsOverride62edl(o["as-override6"], d, "as_override6")); err != nil {
		if vv, ok := fortiAPIPatch(o["as-override6"], "RouterBgpNeighbor-AsOverride6"); ok {
			if err = d.Set("as_override6", vv); err != nil {
				return fmt.Errorf("Error reading as_override6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading as_override6: %v", err)
		}
	}

	if err = d.Set("attribute_unchanged", flattenRouterBgpNeighborAttributeUnchanged2edl(o["attribute-unchanged"], d, "attribute_unchanged")); err != nil {
		if vv, ok := fortiAPIPatch(o["attribute-unchanged"], "RouterBgpNeighbor-AttributeUnchanged"); ok {
			if err = d.Set("attribute_unchanged", vv); err != nil {
				return fmt.Errorf("Error reading attribute_unchanged: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading attribute_unchanged: %v", err)
		}
	}

	if err = d.Set("attribute_unchanged_vpnv4", flattenRouterBgpNeighborAttributeUnchangedVpnv42edl(o["attribute-unchanged-vpnv4"], d, "attribute_unchanged_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["attribute-unchanged-vpnv4"], "RouterBgpNeighbor-AttributeUnchangedVpnv4"); ok {
			if err = d.Set("attribute_unchanged_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading attribute_unchanged_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading attribute_unchanged_vpnv4: %v", err)
		}
	}

	if err = d.Set("attribute_unchanged_vpnv6", flattenRouterBgpNeighborAttributeUnchangedVpnv62edl(o["attribute-unchanged-vpnv6"], d, "attribute_unchanged_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["attribute-unchanged-vpnv6"], "RouterBgpNeighbor-AttributeUnchangedVpnv6"); ok {
			if err = d.Set("attribute_unchanged_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading attribute_unchanged_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading attribute_unchanged_vpnv6: %v", err)
		}
	}

	if err = d.Set("attribute_unchanged6", flattenRouterBgpNeighborAttributeUnchanged62edl(o["attribute-unchanged6"], d, "attribute_unchanged6")); err != nil {
		if vv, ok := fortiAPIPatch(o["attribute-unchanged6"], "RouterBgpNeighbor-AttributeUnchanged6"); ok {
			if err = d.Set("attribute_unchanged6", vv); err != nil {
				return fmt.Errorf("Error reading attribute_unchanged6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading attribute_unchanged6: %v", err)
		}
	}

	if err = d.Set("auth_options", flattenRouterBgpNeighborAuthOptions2edl(o["auth-options"], d, "auth_options")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-options"], "RouterBgpNeighbor-AuthOptions"); ok {
			if err = d.Set("auth_options", vv); err != nil {
				return fmt.Errorf("Error reading auth_options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_options: %v", err)
		}
	}

	if err = d.Set("bfd", flattenRouterBgpNeighborBfd2edl(o["bfd"], d, "bfd")); err != nil {
		if vv, ok := fortiAPIPatch(o["bfd"], "RouterBgpNeighbor-Bfd"); ok {
			if err = d.Set("bfd", vv); err != nil {
				return fmt.Errorf("Error reading bfd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	if err = d.Set("capability_default_originate", flattenRouterBgpNeighborCapabilityDefaultOriginate2edl(o["capability-default-originate"], d, "capability_default_originate")); err != nil {
		if vv, ok := fortiAPIPatch(o["capability-default-originate"], "RouterBgpNeighbor-CapabilityDefaultOriginate"); ok {
			if err = d.Set("capability_default_originate", vv); err != nil {
				return fmt.Errorf("Error reading capability_default_originate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capability_default_originate: %v", err)
		}
	}

	if err = d.Set("capability_default_originate6", flattenRouterBgpNeighborCapabilityDefaultOriginate62edl(o["capability-default-originate6"], d, "capability_default_originate6")); err != nil {
		if vv, ok := fortiAPIPatch(o["capability-default-originate6"], "RouterBgpNeighbor-CapabilityDefaultOriginate6"); ok {
			if err = d.Set("capability_default_originate6", vv); err != nil {
				return fmt.Errorf("Error reading capability_default_originate6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capability_default_originate6: %v", err)
		}
	}

	if err = d.Set("capability_dynamic", flattenRouterBgpNeighborCapabilityDynamic2edl(o["capability-dynamic"], d, "capability_dynamic")); err != nil {
		if vv, ok := fortiAPIPatch(o["capability-dynamic"], "RouterBgpNeighbor-CapabilityDynamic"); ok {
			if err = d.Set("capability_dynamic", vv); err != nil {
				return fmt.Errorf("Error reading capability_dynamic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capability_dynamic: %v", err)
		}
	}

	if err = d.Set("capability_graceful_restart", flattenRouterBgpNeighborCapabilityGracefulRestart2edl(o["capability-graceful-restart"], d, "capability_graceful_restart")); err != nil {
		if vv, ok := fortiAPIPatch(o["capability-graceful-restart"], "RouterBgpNeighbor-CapabilityGracefulRestart"); ok {
			if err = d.Set("capability_graceful_restart", vv); err != nil {
				return fmt.Errorf("Error reading capability_graceful_restart: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capability_graceful_restart: %v", err)
		}
	}

	if err = d.Set("capability_graceful_restart_evpn", flattenRouterBgpNeighborCapabilityGracefulRestartEvpn2edl(o["capability-graceful-restart-evpn"], d, "capability_graceful_restart_evpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["capability-graceful-restart-evpn"], "RouterBgpNeighbor-CapabilityGracefulRestartEvpn"); ok {
			if err = d.Set("capability_graceful_restart_evpn", vv); err != nil {
				return fmt.Errorf("Error reading capability_graceful_restart_evpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capability_graceful_restart_evpn: %v", err)
		}
	}

	if err = d.Set("capability_graceful_restart_vpnv4", flattenRouterBgpNeighborCapabilityGracefulRestartVpnv42edl(o["capability-graceful-restart-vpnv4"], d, "capability_graceful_restart_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["capability-graceful-restart-vpnv4"], "RouterBgpNeighbor-CapabilityGracefulRestartVpnv4"); ok {
			if err = d.Set("capability_graceful_restart_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading capability_graceful_restart_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capability_graceful_restart_vpnv4: %v", err)
		}
	}

	if err = d.Set("capability_graceful_restart_vpnv6", flattenRouterBgpNeighborCapabilityGracefulRestartVpnv62edl(o["capability-graceful-restart-vpnv6"], d, "capability_graceful_restart_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["capability-graceful-restart-vpnv6"], "RouterBgpNeighbor-CapabilityGracefulRestartVpnv6"); ok {
			if err = d.Set("capability_graceful_restart_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading capability_graceful_restart_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capability_graceful_restart_vpnv6: %v", err)
		}
	}

	if err = d.Set("capability_graceful_restart6", flattenRouterBgpNeighborCapabilityGracefulRestart62edl(o["capability-graceful-restart6"], d, "capability_graceful_restart6")); err != nil {
		if vv, ok := fortiAPIPatch(o["capability-graceful-restart6"], "RouterBgpNeighbor-CapabilityGracefulRestart6"); ok {
			if err = d.Set("capability_graceful_restart6", vv); err != nil {
				return fmt.Errorf("Error reading capability_graceful_restart6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capability_graceful_restart6: %v", err)
		}
	}

	if err = d.Set("capability_orf", flattenRouterBgpNeighborCapabilityOrf2edl(o["capability-orf"], d, "capability_orf")); err != nil {
		if vv, ok := fortiAPIPatch(o["capability-orf"], "RouterBgpNeighbor-CapabilityOrf"); ok {
			if err = d.Set("capability_orf", vv); err != nil {
				return fmt.Errorf("Error reading capability_orf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capability_orf: %v", err)
		}
	}

	if err = d.Set("capability_orf6", flattenRouterBgpNeighborCapabilityOrf62edl(o["capability-orf6"], d, "capability_orf6")); err != nil {
		if vv, ok := fortiAPIPatch(o["capability-orf6"], "RouterBgpNeighbor-CapabilityOrf6"); ok {
			if err = d.Set("capability_orf6", vv); err != nil {
				return fmt.Errorf("Error reading capability_orf6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capability_orf6: %v", err)
		}
	}

	if err = d.Set("capability_route_refresh", flattenRouterBgpNeighborCapabilityRouteRefresh2edl(o["capability-route-refresh"], d, "capability_route_refresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["capability-route-refresh"], "RouterBgpNeighbor-CapabilityRouteRefresh"); ok {
			if err = d.Set("capability_route_refresh", vv); err != nil {
				return fmt.Errorf("Error reading capability_route_refresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capability_route_refresh: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("conditional_advertise", flattenRouterBgpNeighborConditionalAdvertise2edl(o["conditional-advertise"], d, "conditional_advertise")); err != nil {
			if vv, ok := fortiAPIPatch(o["conditional-advertise"], "RouterBgpNeighbor-ConditionalAdvertise"); ok {
				if err = d.Set("conditional_advertise", vv); err != nil {
					return fmt.Errorf("Error reading conditional_advertise: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading conditional_advertise: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("conditional_advertise"); ok {
			if err = d.Set("conditional_advertise", flattenRouterBgpNeighborConditionalAdvertise2edl(o["conditional-advertise"], d, "conditional_advertise")); err != nil {
				if vv, ok := fortiAPIPatch(o["conditional-advertise"], "RouterBgpNeighbor-ConditionalAdvertise"); ok {
					if err = d.Set("conditional_advertise", vv); err != nil {
						return fmt.Errorf("Error reading conditional_advertise: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading conditional_advertise: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("conditional_advertise6", flattenRouterBgpNeighborConditionalAdvertise62edl(o["conditional-advertise6"], d, "conditional_advertise6")); err != nil {
			if vv, ok := fortiAPIPatch(o["conditional-advertise6"], "RouterBgpNeighbor-ConditionalAdvertise6"); ok {
				if err = d.Set("conditional_advertise6", vv); err != nil {
					return fmt.Errorf("Error reading conditional_advertise6: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading conditional_advertise6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("conditional_advertise6"); ok {
			if err = d.Set("conditional_advertise6", flattenRouterBgpNeighborConditionalAdvertise62edl(o["conditional-advertise6"], d, "conditional_advertise6")); err != nil {
				if vv, ok := fortiAPIPatch(o["conditional-advertise6"], "RouterBgpNeighbor-ConditionalAdvertise6"); ok {
					if err = d.Set("conditional_advertise6", vv); err != nil {
						return fmt.Errorf("Error reading conditional_advertise6: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading conditional_advertise6: %v", err)
				}
			}
		}
	}

	if err = d.Set("connect_timer", flattenRouterBgpNeighborConnectTimer2edl(o["connect-timer"], d, "connect_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["connect-timer"], "RouterBgpNeighbor-ConnectTimer"); ok {
			if err = d.Set("connect_timer", vv); err != nil {
				return fmt.Errorf("Error reading connect_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading connect_timer: %v", err)
		}
	}

	if err = d.Set("default_originate_routemap", flattenRouterBgpNeighborDefaultOriginateRoutemap2edl(o["default-originate-routemap"], d, "default_originate_routemap")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-originate-routemap"], "RouterBgpNeighbor-DefaultOriginateRoutemap"); ok {
			if err = d.Set("default_originate_routemap", vv); err != nil {
				return fmt.Errorf("Error reading default_originate_routemap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_originate_routemap: %v", err)
		}
	}

	if err = d.Set("default_originate_routemap6", flattenRouterBgpNeighborDefaultOriginateRoutemap62edl(o["default-originate-routemap6"], d, "default_originate_routemap6")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-originate-routemap6"], "RouterBgpNeighbor-DefaultOriginateRoutemap6"); ok {
			if err = d.Set("default_originate_routemap6", vv); err != nil {
				return fmt.Errorf("Error reading default_originate_routemap6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_originate_routemap6: %v", err)
		}
	}

	if err = d.Set("description", flattenRouterBgpNeighborDescription2edl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "RouterBgpNeighbor-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("distribute_list_in", flattenRouterBgpNeighborDistributeListIn2edl(o["distribute-list-in"], d, "distribute_list_in")); err != nil {
		if vv, ok := fortiAPIPatch(o["distribute-list-in"], "RouterBgpNeighbor-DistributeListIn"); ok {
			if err = d.Set("distribute_list_in", vv); err != nil {
				return fmt.Errorf("Error reading distribute_list_in: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distribute_list_in: %v", err)
		}
	}

	if err = d.Set("distribute_list_in_vpnv4", flattenRouterBgpNeighborDistributeListInVpnv42edl(o["distribute-list-in-vpnv4"], d, "distribute_list_in_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["distribute-list-in-vpnv4"], "RouterBgpNeighbor-DistributeListInVpnv4"); ok {
			if err = d.Set("distribute_list_in_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading distribute_list_in_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distribute_list_in_vpnv4: %v", err)
		}
	}

	if err = d.Set("distribute_list_in_vpnv6", flattenRouterBgpNeighborDistributeListInVpnv62edl(o["distribute-list-in-vpnv6"], d, "distribute_list_in_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["distribute-list-in-vpnv6"], "RouterBgpNeighbor-DistributeListInVpnv6"); ok {
			if err = d.Set("distribute_list_in_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading distribute_list_in_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distribute_list_in_vpnv6: %v", err)
		}
	}

	if err = d.Set("distribute_list_in6", flattenRouterBgpNeighborDistributeListIn62edl(o["distribute-list-in6"], d, "distribute_list_in6")); err != nil {
		if vv, ok := fortiAPIPatch(o["distribute-list-in6"], "RouterBgpNeighbor-DistributeListIn6"); ok {
			if err = d.Set("distribute_list_in6", vv); err != nil {
				return fmt.Errorf("Error reading distribute_list_in6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distribute_list_in6: %v", err)
		}
	}

	if err = d.Set("distribute_list_out", flattenRouterBgpNeighborDistributeListOut2edl(o["distribute-list-out"], d, "distribute_list_out")); err != nil {
		if vv, ok := fortiAPIPatch(o["distribute-list-out"], "RouterBgpNeighbor-DistributeListOut"); ok {
			if err = d.Set("distribute_list_out", vv); err != nil {
				return fmt.Errorf("Error reading distribute_list_out: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distribute_list_out: %v", err)
		}
	}

	if err = d.Set("distribute_list_out_vpnv4", flattenRouterBgpNeighborDistributeListOutVpnv42edl(o["distribute-list-out-vpnv4"], d, "distribute_list_out_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["distribute-list-out-vpnv4"], "RouterBgpNeighbor-DistributeListOutVpnv4"); ok {
			if err = d.Set("distribute_list_out_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading distribute_list_out_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distribute_list_out_vpnv4: %v", err)
		}
	}

	if err = d.Set("distribute_list_out_vpnv6", flattenRouterBgpNeighborDistributeListOutVpnv62edl(o["distribute-list-out-vpnv6"], d, "distribute_list_out_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["distribute-list-out-vpnv6"], "RouterBgpNeighbor-DistributeListOutVpnv6"); ok {
			if err = d.Set("distribute_list_out_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading distribute_list_out_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distribute_list_out_vpnv6: %v", err)
		}
	}

	if err = d.Set("distribute_list_out6", flattenRouterBgpNeighborDistributeListOut62edl(o["distribute-list-out6"], d, "distribute_list_out6")); err != nil {
		if vv, ok := fortiAPIPatch(o["distribute-list-out6"], "RouterBgpNeighbor-DistributeListOut6"); ok {
			if err = d.Set("distribute_list_out6", vv); err != nil {
				return fmt.Errorf("Error reading distribute_list_out6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distribute_list_out6: %v", err)
		}
	}

	if err = d.Set("dont_capability_negotiate", flattenRouterBgpNeighborDontCapabilityNegotiate2edl(o["dont-capability-negotiate"], d, "dont_capability_negotiate")); err != nil {
		if vv, ok := fortiAPIPatch(o["dont-capability-negotiate"], "RouterBgpNeighbor-DontCapabilityNegotiate"); ok {
			if err = d.Set("dont_capability_negotiate", vv); err != nil {
				return fmt.Errorf("Error reading dont_capability_negotiate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dont_capability_negotiate: %v", err)
		}
	}

	if err = d.Set("ebgp_enforce_multihop", flattenRouterBgpNeighborEbgpEnforceMultihop2edl(o["ebgp-enforce-multihop"], d, "ebgp_enforce_multihop")); err != nil {
		if vv, ok := fortiAPIPatch(o["ebgp-enforce-multihop"], "RouterBgpNeighbor-EbgpEnforceMultihop"); ok {
			if err = d.Set("ebgp_enforce_multihop", vv); err != nil {
				return fmt.Errorf("Error reading ebgp_enforce_multihop: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ebgp_enforce_multihop: %v", err)
		}
	}

	if err = d.Set("ebgp_multihop_ttl", flattenRouterBgpNeighborEbgpMultihopTtl2edl(o["ebgp-multihop-ttl"], d, "ebgp_multihop_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["ebgp-multihop-ttl"], "RouterBgpNeighbor-EbgpMultihopTtl"); ok {
			if err = d.Set("ebgp_multihop_ttl", vv); err != nil {
				return fmt.Errorf("Error reading ebgp_multihop_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ebgp_multihop_ttl: %v", err)
		}
	}

	if err = d.Set("filter_list_in", flattenRouterBgpNeighborFilterListIn2edl(o["filter-list-in"], d, "filter_list_in")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-list-in"], "RouterBgpNeighbor-FilterListIn"); ok {
			if err = d.Set("filter_list_in", vv); err != nil {
				return fmt.Errorf("Error reading filter_list_in: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_list_in: %v", err)
		}
	}

	if err = d.Set("filter_list_in_vpnv4", flattenRouterBgpNeighborFilterListInVpnv42edl(o["filter-list-in-vpnv4"], d, "filter_list_in_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-list-in-vpnv4"], "RouterBgpNeighbor-FilterListInVpnv4"); ok {
			if err = d.Set("filter_list_in_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading filter_list_in_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_list_in_vpnv4: %v", err)
		}
	}

	if err = d.Set("filter_list_in_vpnv6", flattenRouterBgpNeighborFilterListInVpnv62edl(o["filter-list-in-vpnv6"], d, "filter_list_in_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-list-in-vpnv6"], "RouterBgpNeighbor-FilterListInVpnv6"); ok {
			if err = d.Set("filter_list_in_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading filter_list_in_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_list_in_vpnv6: %v", err)
		}
	}

	if err = d.Set("filter_list_in6", flattenRouterBgpNeighborFilterListIn62edl(o["filter-list-in6"], d, "filter_list_in6")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-list-in6"], "RouterBgpNeighbor-FilterListIn6"); ok {
			if err = d.Set("filter_list_in6", vv); err != nil {
				return fmt.Errorf("Error reading filter_list_in6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_list_in6: %v", err)
		}
	}

	if err = d.Set("filter_list_out", flattenRouterBgpNeighborFilterListOut2edl(o["filter-list-out"], d, "filter_list_out")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-list-out"], "RouterBgpNeighbor-FilterListOut"); ok {
			if err = d.Set("filter_list_out", vv); err != nil {
				return fmt.Errorf("Error reading filter_list_out: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_list_out: %v", err)
		}
	}

	if err = d.Set("filter_list_out_vpnv4", flattenRouterBgpNeighborFilterListOutVpnv42edl(o["filter-list-out-vpnv4"], d, "filter_list_out_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-list-out-vpnv4"], "RouterBgpNeighbor-FilterListOutVpnv4"); ok {
			if err = d.Set("filter_list_out_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading filter_list_out_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_list_out_vpnv4: %v", err)
		}
	}

	if err = d.Set("filter_list_out_vpnv6", flattenRouterBgpNeighborFilterListOutVpnv62edl(o["filter-list-out-vpnv6"], d, "filter_list_out_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-list-out-vpnv6"], "RouterBgpNeighbor-FilterListOutVpnv6"); ok {
			if err = d.Set("filter_list_out_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading filter_list_out_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_list_out_vpnv6: %v", err)
		}
	}

	if err = d.Set("filter_list_out6", flattenRouterBgpNeighborFilterListOut62edl(o["filter-list-out6"], d, "filter_list_out6")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-list-out6"], "RouterBgpNeighbor-FilterListOut6"); ok {
			if err = d.Set("filter_list_out6", vv); err != nil {
				return fmt.Errorf("Error reading filter_list_out6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_list_out6: %v", err)
		}
	}

	if err = d.Set("holdtime_timer", flattenRouterBgpNeighborHoldtimeTimer2edl(o["holdtime-timer"], d, "holdtime_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["holdtime-timer"], "RouterBgpNeighbor-HoldtimeTimer"); ok {
			if err = d.Set("holdtime_timer", vv); err != nil {
				return fmt.Errorf("Error reading holdtime_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading holdtime_timer: %v", err)
		}
	}

	if err = d.Set("interface", flattenRouterBgpNeighborInterface2edl(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "RouterBgpNeighbor-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ip", flattenRouterBgpNeighborIp2edl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "RouterBgpNeighbor-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("keep_alive_timer", flattenRouterBgpNeighborKeepAliveTimer2edl(o["keep-alive-timer"], d, "keep_alive_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["keep-alive-timer"], "RouterBgpNeighbor-KeepAliveTimer"); ok {
			if err = d.Set("keep_alive_timer", vv); err != nil {
				return fmt.Errorf("Error reading keep_alive_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keep_alive_timer: %v", err)
		}
	}

	if err = d.Set("link_down_failover", flattenRouterBgpNeighborLinkDownFailover2edl(o["link-down-failover"], d, "link_down_failover")); err != nil {
		if vv, ok := fortiAPIPatch(o["link-down-failover"], "RouterBgpNeighbor-LinkDownFailover"); ok {
			if err = d.Set("link_down_failover", vv); err != nil {
				return fmt.Errorf("Error reading link_down_failover: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading link_down_failover: %v", err)
		}
	}

	if err = d.Set("local_as", flattenRouterBgpNeighborLocalAs2edl(o["local-as"], d, "local_as")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-as"], "RouterBgpNeighbor-LocalAs"); ok {
			if err = d.Set("local_as", vv); err != nil {
				return fmt.Errorf("Error reading local_as: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_as: %v", err)
		}
	}

	if err = d.Set("local_as_no_prepend", flattenRouterBgpNeighborLocalAsNoPrepend2edl(o["local-as-no-prepend"], d, "local_as_no_prepend")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-as-no-prepend"], "RouterBgpNeighbor-LocalAsNoPrepend"); ok {
			if err = d.Set("local_as_no_prepend", vv); err != nil {
				return fmt.Errorf("Error reading local_as_no_prepend: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_as_no_prepend: %v", err)
		}
	}

	if err = d.Set("local_as_replace_as", flattenRouterBgpNeighborLocalAsReplaceAs2edl(o["local-as-replace-as"], d, "local_as_replace_as")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-as-replace-as"], "RouterBgpNeighbor-LocalAsReplaceAs"); ok {
			if err = d.Set("local_as_replace_as", vv); err != nil {
				return fmt.Errorf("Error reading local_as_replace_as: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_as_replace_as: %v", err)
		}
	}

	if err = d.Set("maximum_prefix", flattenRouterBgpNeighborMaximumPrefix2edl(o["maximum-prefix"], d, "maximum_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-prefix"], "RouterBgpNeighbor-MaximumPrefix"); ok {
			if err = d.Set("maximum_prefix", vv); err != nil {
				return fmt.Errorf("Error reading maximum_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_prefix: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_evpn", flattenRouterBgpNeighborMaximumPrefixEvpn2edl(o["maximum-prefix-evpn"], d, "maximum_prefix_evpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-prefix-evpn"], "RouterBgpNeighbor-MaximumPrefixEvpn"); ok {
			if err = d.Set("maximum_prefix_evpn", vv); err != nil {
				return fmt.Errorf("Error reading maximum_prefix_evpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_prefix_evpn: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_threshold", flattenRouterBgpNeighborMaximumPrefixThreshold2edl(o["maximum-prefix-threshold"], d, "maximum_prefix_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-prefix-threshold"], "RouterBgpNeighbor-MaximumPrefixThreshold"); ok {
			if err = d.Set("maximum_prefix_threshold", vv); err != nil {
				return fmt.Errorf("Error reading maximum_prefix_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_prefix_threshold: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_threshold_evpn", flattenRouterBgpNeighborMaximumPrefixThresholdEvpn2edl(o["maximum-prefix-threshold-evpn"], d, "maximum_prefix_threshold_evpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-prefix-threshold-evpn"], "RouterBgpNeighbor-MaximumPrefixThresholdEvpn"); ok {
			if err = d.Set("maximum_prefix_threshold_evpn", vv); err != nil {
				return fmt.Errorf("Error reading maximum_prefix_threshold_evpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_prefix_threshold_evpn: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_threshold_vpnv4", flattenRouterBgpNeighborMaximumPrefixThresholdVpnv42edl(o["maximum-prefix-threshold-vpnv4"], d, "maximum_prefix_threshold_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-prefix-threshold-vpnv4"], "RouterBgpNeighbor-MaximumPrefixThresholdVpnv4"); ok {
			if err = d.Set("maximum_prefix_threshold_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading maximum_prefix_threshold_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_prefix_threshold_vpnv4: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_threshold_vpnv6", flattenRouterBgpNeighborMaximumPrefixThresholdVpnv62edl(o["maximum-prefix-threshold-vpnv6"], d, "maximum_prefix_threshold_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-prefix-threshold-vpnv6"], "RouterBgpNeighbor-MaximumPrefixThresholdVpnv6"); ok {
			if err = d.Set("maximum_prefix_threshold_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading maximum_prefix_threshold_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_prefix_threshold_vpnv6: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_threshold6", flattenRouterBgpNeighborMaximumPrefixThreshold62edl(o["maximum-prefix-threshold6"], d, "maximum_prefix_threshold6")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-prefix-threshold6"], "RouterBgpNeighbor-MaximumPrefixThreshold6"); ok {
			if err = d.Set("maximum_prefix_threshold6", vv); err != nil {
				return fmt.Errorf("Error reading maximum_prefix_threshold6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_prefix_threshold6: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_vpnv4", flattenRouterBgpNeighborMaximumPrefixVpnv42edl(o["maximum-prefix-vpnv4"], d, "maximum_prefix_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-prefix-vpnv4"], "RouterBgpNeighbor-MaximumPrefixVpnv4"); ok {
			if err = d.Set("maximum_prefix_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading maximum_prefix_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_prefix_vpnv4: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_vpnv6", flattenRouterBgpNeighborMaximumPrefixVpnv62edl(o["maximum-prefix-vpnv6"], d, "maximum_prefix_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-prefix-vpnv6"], "RouterBgpNeighbor-MaximumPrefixVpnv6"); ok {
			if err = d.Set("maximum_prefix_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading maximum_prefix_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_prefix_vpnv6: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_warning_only", flattenRouterBgpNeighborMaximumPrefixWarningOnly2edl(o["maximum-prefix-warning-only"], d, "maximum_prefix_warning_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-prefix-warning-only"], "RouterBgpNeighbor-MaximumPrefixWarningOnly"); ok {
			if err = d.Set("maximum_prefix_warning_only", vv); err != nil {
				return fmt.Errorf("Error reading maximum_prefix_warning_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_prefix_warning_only: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_warning_only_evpn", flattenRouterBgpNeighborMaximumPrefixWarningOnlyEvpn2edl(o["maximum-prefix-warning-only-evpn"], d, "maximum_prefix_warning_only_evpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-prefix-warning-only-evpn"], "RouterBgpNeighbor-MaximumPrefixWarningOnlyEvpn"); ok {
			if err = d.Set("maximum_prefix_warning_only_evpn", vv); err != nil {
				return fmt.Errorf("Error reading maximum_prefix_warning_only_evpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_prefix_warning_only_evpn: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_warning_only_vpnv4", flattenRouterBgpNeighborMaximumPrefixWarningOnlyVpnv42edl(o["maximum-prefix-warning-only-vpnv4"], d, "maximum_prefix_warning_only_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-prefix-warning-only-vpnv4"], "RouterBgpNeighbor-MaximumPrefixWarningOnlyVpnv4"); ok {
			if err = d.Set("maximum_prefix_warning_only_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading maximum_prefix_warning_only_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_prefix_warning_only_vpnv4: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_warning_only_vpnv6", flattenRouterBgpNeighborMaximumPrefixWarningOnlyVpnv62edl(o["maximum-prefix-warning-only-vpnv6"], d, "maximum_prefix_warning_only_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-prefix-warning-only-vpnv6"], "RouterBgpNeighbor-MaximumPrefixWarningOnlyVpnv6"); ok {
			if err = d.Set("maximum_prefix_warning_only_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading maximum_prefix_warning_only_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_prefix_warning_only_vpnv6: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_warning_only6", flattenRouterBgpNeighborMaximumPrefixWarningOnly62edl(o["maximum-prefix-warning-only6"], d, "maximum_prefix_warning_only6")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-prefix-warning-only6"], "RouterBgpNeighbor-MaximumPrefixWarningOnly6"); ok {
			if err = d.Set("maximum_prefix_warning_only6", vv); err != nil {
				return fmt.Errorf("Error reading maximum_prefix_warning_only6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_prefix_warning_only6: %v", err)
		}
	}

	if err = d.Set("maximum_prefix6", flattenRouterBgpNeighborMaximumPrefix62edl(o["maximum-prefix6"], d, "maximum_prefix6")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-prefix6"], "RouterBgpNeighbor-MaximumPrefix6"); ok {
			if err = d.Set("maximum_prefix6", vv); err != nil {
				return fmt.Errorf("Error reading maximum_prefix6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_prefix6: %v", err)
		}
	}

	if err = d.Set("next_hop_self", flattenRouterBgpNeighborNextHopSelf2edl(o["next-hop-self"], d, "next_hop_self")); err != nil {
		if vv, ok := fortiAPIPatch(o["next-hop-self"], "RouterBgpNeighbor-NextHopSelf"); ok {
			if err = d.Set("next_hop_self", vv); err != nil {
				return fmt.Errorf("Error reading next_hop_self: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading next_hop_self: %v", err)
		}
	}

	if err = d.Set("next_hop_self_rr", flattenRouterBgpNeighborNextHopSelfRr2edl(o["next-hop-self-rr"], d, "next_hop_self_rr")); err != nil {
		if vv, ok := fortiAPIPatch(o["next-hop-self-rr"], "RouterBgpNeighbor-NextHopSelfRr"); ok {
			if err = d.Set("next_hop_self_rr", vv); err != nil {
				return fmt.Errorf("Error reading next_hop_self_rr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading next_hop_self_rr: %v", err)
		}
	}

	if err = d.Set("next_hop_self_rr6", flattenRouterBgpNeighborNextHopSelfRr62edl(o["next-hop-self-rr6"], d, "next_hop_self_rr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["next-hop-self-rr6"], "RouterBgpNeighbor-NextHopSelfRr6"); ok {
			if err = d.Set("next_hop_self_rr6", vv); err != nil {
				return fmt.Errorf("Error reading next_hop_self_rr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading next_hop_self_rr6: %v", err)
		}
	}

	if err = d.Set("next_hop_self_vpnv4", flattenRouterBgpNeighborNextHopSelfVpnv42edl(o["next-hop-self-vpnv4"], d, "next_hop_self_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["next-hop-self-vpnv4"], "RouterBgpNeighbor-NextHopSelfVpnv4"); ok {
			if err = d.Set("next_hop_self_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading next_hop_self_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading next_hop_self_vpnv4: %v", err)
		}
	}

	if err = d.Set("next_hop_self_vpnv6", flattenRouterBgpNeighborNextHopSelfVpnv62edl(o["next-hop-self-vpnv6"], d, "next_hop_self_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["next-hop-self-vpnv6"], "RouterBgpNeighbor-NextHopSelfVpnv6"); ok {
			if err = d.Set("next_hop_self_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading next_hop_self_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading next_hop_self_vpnv6: %v", err)
		}
	}

	if err = d.Set("next_hop_self6", flattenRouterBgpNeighborNextHopSelf62edl(o["next-hop-self6"], d, "next_hop_self6")); err != nil {
		if vv, ok := fortiAPIPatch(o["next-hop-self6"], "RouterBgpNeighbor-NextHopSelf6"); ok {
			if err = d.Set("next_hop_self6", vv); err != nil {
				return fmt.Errorf("Error reading next_hop_self6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading next_hop_self6: %v", err)
		}
	}

	if err = d.Set("override_capability", flattenRouterBgpNeighborOverrideCapability2edl(o["override-capability"], d, "override_capability")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-capability"], "RouterBgpNeighbor-OverrideCapability"); ok {
			if err = d.Set("override_capability", vv); err != nil {
				return fmt.Errorf("Error reading override_capability: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_capability: %v", err)
		}
	}

	if err = d.Set("passive", flattenRouterBgpNeighborPassive2edl(o["passive"], d, "passive")); err != nil {
		if vv, ok := fortiAPIPatch(o["passive"], "RouterBgpNeighbor-Passive"); ok {
			if err = d.Set("passive", vv); err != nil {
				return fmt.Errorf("Error reading passive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading passive: %v", err)
		}
	}

	if err = d.Set("prefix_list_in", flattenRouterBgpNeighborPrefixListIn2edl(o["prefix-list-in"], d, "prefix_list_in")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix-list-in"], "RouterBgpNeighbor-PrefixListIn"); ok {
			if err = d.Set("prefix_list_in", vv); err != nil {
				return fmt.Errorf("Error reading prefix_list_in: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix_list_in: %v", err)
		}
	}

	if err = d.Set("prefix_list_in_vpnv4", flattenRouterBgpNeighborPrefixListInVpnv42edl(o["prefix-list-in-vpnv4"], d, "prefix_list_in_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix-list-in-vpnv4"], "RouterBgpNeighbor-PrefixListInVpnv4"); ok {
			if err = d.Set("prefix_list_in_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading prefix_list_in_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix_list_in_vpnv4: %v", err)
		}
	}

	if err = d.Set("prefix_list_in_vpnv6", flattenRouterBgpNeighborPrefixListInVpnv62edl(o["prefix-list-in-vpnv6"], d, "prefix_list_in_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix-list-in-vpnv6"], "RouterBgpNeighbor-PrefixListInVpnv6"); ok {
			if err = d.Set("prefix_list_in_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading prefix_list_in_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix_list_in_vpnv6: %v", err)
		}
	}

	if err = d.Set("prefix_list_in6", flattenRouterBgpNeighborPrefixListIn62edl(o["prefix-list-in6"], d, "prefix_list_in6")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix-list-in6"], "RouterBgpNeighbor-PrefixListIn6"); ok {
			if err = d.Set("prefix_list_in6", vv); err != nil {
				return fmt.Errorf("Error reading prefix_list_in6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix_list_in6: %v", err)
		}
	}

	if err = d.Set("prefix_list_out", flattenRouterBgpNeighborPrefixListOut2edl(o["prefix-list-out"], d, "prefix_list_out")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix-list-out"], "RouterBgpNeighbor-PrefixListOut"); ok {
			if err = d.Set("prefix_list_out", vv); err != nil {
				return fmt.Errorf("Error reading prefix_list_out: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix_list_out: %v", err)
		}
	}

	if err = d.Set("prefix_list_out_vpnv4", flattenRouterBgpNeighborPrefixListOutVpnv42edl(o["prefix-list-out-vpnv4"], d, "prefix_list_out_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix-list-out-vpnv4"], "RouterBgpNeighbor-PrefixListOutVpnv4"); ok {
			if err = d.Set("prefix_list_out_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading prefix_list_out_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix_list_out_vpnv4: %v", err)
		}
	}

	if err = d.Set("prefix_list_out_vpnv6", flattenRouterBgpNeighborPrefixListOutVpnv62edl(o["prefix-list-out-vpnv6"], d, "prefix_list_out_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix-list-out-vpnv6"], "RouterBgpNeighbor-PrefixListOutVpnv6"); ok {
			if err = d.Set("prefix_list_out_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading prefix_list_out_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix_list_out_vpnv6: %v", err)
		}
	}

	if err = d.Set("prefix_list_out6", flattenRouterBgpNeighborPrefixListOut62edl(o["prefix-list-out6"], d, "prefix_list_out6")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix-list-out6"], "RouterBgpNeighbor-PrefixListOut6"); ok {
			if err = d.Set("prefix_list_out6", vv); err != nil {
				return fmt.Errorf("Error reading prefix_list_out6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix_list_out6: %v", err)
		}
	}

	if err = d.Set("remote_as", flattenRouterBgpNeighborRemoteAs2edl(o["remote-as"], d, "remote_as")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-as"], "RouterBgpNeighbor-RemoteAs"); ok {
			if err = d.Set("remote_as", vv); err != nil {
				return fmt.Errorf("Error reading remote_as: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_as: %v", err)
		}
	}

	if err = d.Set("remove_private_as", flattenRouterBgpNeighborRemovePrivateAs2edl(o["remove-private-as"], d, "remove_private_as")); err != nil {
		if vv, ok := fortiAPIPatch(o["remove-private-as"], "RouterBgpNeighbor-RemovePrivateAs"); ok {
			if err = d.Set("remove_private_as", vv); err != nil {
				return fmt.Errorf("Error reading remove_private_as: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remove_private_as: %v", err)
		}
	}

	if err = d.Set("remove_private_as_evpn", flattenRouterBgpNeighborRemovePrivateAsEvpn2edl(o["remove-private-as-evpn"], d, "remove_private_as_evpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["remove-private-as-evpn"], "RouterBgpNeighbor-RemovePrivateAsEvpn"); ok {
			if err = d.Set("remove_private_as_evpn", vv); err != nil {
				return fmt.Errorf("Error reading remove_private_as_evpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remove_private_as_evpn: %v", err)
		}
	}

	if err = d.Set("remove_private_as_vpnv4", flattenRouterBgpNeighborRemovePrivateAsVpnv42edl(o["remove-private-as-vpnv4"], d, "remove_private_as_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["remove-private-as-vpnv4"], "RouterBgpNeighbor-RemovePrivateAsVpnv4"); ok {
			if err = d.Set("remove_private_as_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading remove_private_as_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remove_private_as_vpnv4: %v", err)
		}
	}

	if err = d.Set("remove_private_as_vpnv6", flattenRouterBgpNeighborRemovePrivateAsVpnv62edl(o["remove-private-as-vpnv6"], d, "remove_private_as_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["remove-private-as-vpnv6"], "RouterBgpNeighbor-RemovePrivateAsVpnv6"); ok {
			if err = d.Set("remove_private_as_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading remove_private_as_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remove_private_as_vpnv6: %v", err)
		}
	}

	if err = d.Set("remove_private_as6", flattenRouterBgpNeighborRemovePrivateAs62edl(o["remove-private-as6"], d, "remove_private_as6")); err != nil {
		if vv, ok := fortiAPIPatch(o["remove-private-as6"], "RouterBgpNeighbor-RemovePrivateAs6"); ok {
			if err = d.Set("remove_private_as6", vv); err != nil {
				return fmt.Errorf("Error reading remove_private_as6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remove_private_as6: %v", err)
		}
	}

	if err = d.Set("restart_time", flattenRouterBgpNeighborRestartTime2edl(o["restart-time"], d, "restart_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["restart-time"], "RouterBgpNeighbor-RestartTime"); ok {
			if err = d.Set("restart_time", vv); err != nil {
				return fmt.Errorf("Error reading restart_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restart_time: %v", err)
		}
	}

	if err = d.Set("retain_stale_time", flattenRouterBgpNeighborRetainStaleTime2edl(o["retain-stale-time"], d, "retain_stale_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["retain-stale-time"], "RouterBgpNeighbor-RetainStaleTime"); ok {
			if err = d.Set("retain_stale_time", vv); err != nil {
				return fmt.Errorf("Error reading retain_stale_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading retain_stale_time: %v", err)
		}
	}

	if err = d.Set("route_map_in", flattenRouterBgpNeighborRouteMapIn2edl(o["route-map-in"], d, "route_map_in")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-map-in"], "RouterBgpNeighbor-RouteMapIn"); ok {
			if err = d.Set("route_map_in", vv); err != nil {
				return fmt.Errorf("Error reading route_map_in: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_map_in: %v", err)
		}
	}

	if err = d.Set("route_map_in_evpn", flattenRouterBgpNeighborRouteMapInEvpn2edl(o["route-map-in-evpn"], d, "route_map_in_evpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-map-in-evpn"], "RouterBgpNeighbor-RouteMapInEvpn"); ok {
			if err = d.Set("route_map_in_evpn", vv); err != nil {
				return fmt.Errorf("Error reading route_map_in_evpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_map_in_evpn: %v", err)
		}
	}

	if err = d.Set("route_map_in_vpnv4", flattenRouterBgpNeighborRouteMapInVpnv42edl(o["route-map-in-vpnv4"], d, "route_map_in_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-map-in-vpnv4"], "RouterBgpNeighbor-RouteMapInVpnv4"); ok {
			if err = d.Set("route_map_in_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading route_map_in_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_map_in_vpnv4: %v", err)
		}
	}

	if err = d.Set("route_map_in_vpnv6", flattenRouterBgpNeighborRouteMapInVpnv62edl(o["route-map-in-vpnv6"], d, "route_map_in_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-map-in-vpnv6"], "RouterBgpNeighbor-RouteMapInVpnv6"); ok {
			if err = d.Set("route_map_in_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading route_map_in_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_map_in_vpnv6: %v", err)
		}
	}

	if err = d.Set("route_map_in6", flattenRouterBgpNeighborRouteMapIn62edl(o["route-map-in6"], d, "route_map_in6")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-map-in6"], "RouterBgpNeighbor-RouteMapIn6"); ok {
			if err = d.Set("route_map_in6", vv); err != nil {
				return fmt.Errorf("Error reading route_map_in6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_map_in6: %v", err)
		}
	}

	if err = d.Set("route_map_out", flattenRouterBgpNeighborRouteMapOut2edl(o["route-map-out"], d, "route_map_out")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-map-out"], "RouterBgpNeighbor-RouteMapOut"); ok {
			if err = d.Set("route_map_out", vv); err != nil {
				return fmt.Errorf("Error reading route_map_out: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_map_out: %v", err)
		}
	}

	if err = d.Set("route_map_out_evpn", flattenRouterBgpNeighborRouteMapOutEvpn2edl(o["route-map-out-evpn"], d, "route_map_out_evpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-map-out-evpn"], "RouterBgpNeighbor-RouteMapOutEvpn"); ok {
			if err = d.Set("route_map_out_evpn", vv); err != nil {
				return fmt.Errorf("Error reading route_map_out_evpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_map_out_evpn: %v", err)
		}
	}

	if err = d.Set("route_map_out_preferable", flattenRouterBgpNeighborRouteMapOutPreferable2edl(o["route-map-out-preferable"], d, "route_map_out_preferable")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-map-out-preferable"], "RouterBgpNeighbor-RouteMapOutPreferable"); ok {
			if err = d.Set("route_map_out_preferable", vv); err != nil {
				return fmt.Errorf("Error reading route_map_out_preferable: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_map_out_preferable: %v", err)
		}
	}

	if err = d.Set("route_map_out_vpnv4", flattenRouterBgpNeighborRouteMapOutVpnv42edl(o["route-map-out-vpnv4"], d, "route_map_out_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-map-out-vpnv4"], "RouterBgpNeighbor-RouteMapOutVpnv4"); ok {
			if err = d.Set("route_map_out_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading route_map_out_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_map_out_vpnv4: %v", err)
		}
	}

	if err = d.Set("route_map_out_vpnv4_preferable", flattenRouterBgpNeighborRouteMapOutVpnv4Preferable2edl(o["route-map-out-vpnv4-preferable"], d, "route_map_out_vpnv4_preferable")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-map-out-vpnv4-preferable"], "RouterBgpNeighbor-RouteMapOutVpnv4Preferable"); ok {
			if err = d.Set("route_map_out_vpnv4_preferable", vv); err != nil {
				return fmt.Errorf("Error reading route_map_out_vpnv4_preferable: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_map_out_vpnv4_preferable: %v", err)
		}
	}

	if err = d.Set("route_map_out_vpnv6", flattenRouterBgpNeighborRouteMapOutVpnv62edl(o["route-map-out-vpnv6"], d, "route_map_out_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-map-out-vpnv6"], "RouterBgpNeighbor-RouteMapOutVpnv6"); ok {
			if err = d.Set("route_map_out_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading route_map_out_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_map_out_vpnv6: %v", err)
		}
	}

	if err = d.Set("route_map_out_vpnv6_preferable", flattenRouterBgpNeighborRouteMapOutVpnv6Preferable2edl(o["route-map-out-vpnv6-preferable"], d, "route_map_out_vpnv6_preferable")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-map-out-vpnv6-preferable"], "RouterBgpNeighbor-RouteMapOutVpnv6Preferable"); ok {
			if err = d.Set("route_map_out_vpnv6_preferable", vv); err != nil {
				return fmt.Errorf("Error reading route_map_out_vpnv6_preferable: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_map_out_vpnv6_preferable: %v", err)
		}
	}

	if err = d.Set("route_map_out6", flattenRouterBgpNeighborRouteMapOut62edl(o["route-map-out6"], d, "route_map_out6")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-map-out6"], "RouterBgpNeighbor-RouteMapOut6"); ok {
			if err = d.Set("route_map_out6", vv); err != nil {
				return fmt.Errorf("Error reading route_map_out6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_map_out6: %v", err)
		}
	}

	if err = d.Set("route_map_out6_preferable", flattenRouterBgpNeighborRouteMapOut6Preferable2edl(o["route-map-out6-preferable"], d, "route_map_out6_preferable")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-map-out6-preferable"], "RouterBgpNeighbor-RouteMapOut6Preferable"); ok {
			if err = d.Set("route_map_out6_preferable", vv); err != nil {
				return fmt.Errorf("Error reading route_map_out6_preferable: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_map_out6_preferable: %v", err)
		}
	}

	if err = d.Set("route_reflector_client", flattenRouterBgpNeighborRouteReflectorClient2edl(o["route-reflector-client"], d, "route_reflector_client")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-reflector-client"], "RouterBgpNeighbor-RouteReflectorClient"); ok {
			if err = d.Set("route_reflector_client", vv); err != nil {
				return fmt.Errorf("Error reading route_reflector_client: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_reflector_client: %v", err)
		}
	}

	if err = d.Set("route_reflector_client_evpn", flattenRouterBgpNeighborRouteReflectorClientEvpn2edl(o["route-reflector-client-evpn"], d, "route_reflector_client_evpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-reflector-client-evpn"], "RouterBgpNeighbor-RouteReflectorClientEvpn"); ok {
			if err = d.Set("route_reflector_client_evpn", vv); err != nil {
				return fmt.Errorf("Error reading route_reflector_client_evpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_reflector_client_evpn: %v", err)
		}
	}

	if err = d.Set("route_reflector_client_vpnv4", flattenRouterBgpNeighborRouteReflectorClientVpnv42edl(o["route-reflector-client-vpnv4"], d, "route_reflector_client_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-reflector-client-vpnv4"], "RouterBgpNeighbor-RouteReflectorClientVpnv4"); ok {
			if err = d.Set("route_reflector_client_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading route_reflector_client_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_reflector_client_vpnv4: %v", err)
		}
	}

	if err = d.Set("route_reflector_client_vpnv6", flattenRouterBgpNeighborRouteReflectorClientVpnv62edl(o["route-reflector-client-vpnv6"], d, "route_reflector_client_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-reflector-client-vpnv6"], "RouterBgpNeighbor-RouteReflectorClientVpnv6"); ok {
			if err = d.Set("route_reflector_client_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading route_reflector_client_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_reflector_client_vpnv6: %v", err)
		}
	}

	if err = d.Set("route_reflector_client6", flattenRouterBgpNeighborRouteReflectorClient62edl(o["route-reflector-client6"], d, "route_reflector_client6")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-reflector-client6"], "RouterBgpNeighbor-RouteReflectorClient6"); ok {
			if err = d.Set("route_reflector_client6", vv); err != nil {
				return fmt.Errorf("Error reading route_reflector_client6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_reflector_client6: %v", err)
		}
	}

	if err = d.Set("route_server_client", flattenRouterBgpNeighborRouteServerClient2edl(o["route-server-client"], d, "route_server_client")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-server-client"], "RouterBgpNeighbor-RouteServerClient"); ok {
			if err = d.Set("route_server_client", vv); err != nil {
				return fmt.Errorf("Error reading route_server_client: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_server_client: %v", err)
		}
	}

	if err = d.Set("route_server_client_evpn", flattenRouterBgpNeighborRouteServerClientEvpn2edl(o["route-server-client-evpn"], d, "route_server_client_evpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-server-client-evpn"], "RouterBgpNeighbor-RouteServerClientEvpn"); ok {
			if err = d.Set("route_server_client_evpn", vv); err != nil {
				return fmt.Errorf("Error reading route_server_client_evpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_server_client_evpn: %v", err)
		}
	}

	if err = d.Set("route_server_client_vpnv4", flattenRouterBgpNeighborRouteServerClientVpnv42edl(o["route-server-client-vpnv4"], d, "route_server_client_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-server-client-vpnv4"], "RouterBgpNeighbor-RouteServerClientVpnv4"); ok {
			if err = d.Set("route_server_client_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading route_server_client_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_server_client_vpnv4: %v", err)
		}
	}

	if err = d.Set("route_server_client_vpnv6", flattenRouterBgpNeighborRouteServerClientVpnv62edl(o["route-server-client-vpnv6"], d, "route_server_client_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-server-client-vpnv6"], "RouterBgpNeighbor-RouteServerClientVpnv6"); ok {
			if err = d.Set("route_server_client_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading route_server_client_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_server_client_vpnv6: %v", err)
		}
	}

	if err = d.Set("route_server_client6", flattenRouterBgpNeighborRouteServerClient62edl(o["route-server-client6"], d, "route_server_client6")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-server-client6"], "RouterBgpNeighbor-RouteServerClient6"); ok {
			if err = d.Set("route_server_client6", vv); err != nil {
				return fmt.Errorf("Error reading route_server_client6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_server_client6: %v", err)
		}
	}

	if err = d.Set("rr_attr_allow_change", flattenRouterBgpNeighborRrAttrAllowChange2edl(o["rr-attr-allow-change"], d, "rr_attr_allow_change")); err != nil {
		if vv, ok := fortiAPIPatch(o["rr-attr-allow-change"], "RouterBgpNeighbor-RrAttrAllowChange"); ok {
			if err = d.Set("rr_attr_allow_change", vv); err != nil {
				return fmt.Errorf("Error reading rr_attr_allow_change: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rr_attr_allow_change: %v", err)
		}
	}

	if err = d.Set("rr_attr_allow_change_evpn", flattenRouterBgpNeighborRrAttrAllowChangeEvpn2edl(o["rr-attr-allow-change-evpn"], d, "rr_attr_allow_change_evpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["rr-attr-allow-change-evpn"], "RouterBgpNeighbor-RrAttrAllowChangeEvpn"); ok {
			if err = d.Set("rr_attr_allow_change_evpn", vv); err != nil {
				return fmt.Errorf("Error reading rr_attr_allow_change_evpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rr_attr_allow_change_evpn: %v", err)
		}
	}

	if err = d.Set("rr_attr_allow_change_vpnv4", flattenRouterBgpNeighborRrAttrAllowChangeVpnv42edl(o["rr-attr-allow-change-vpnv4"], d, "rr_attr_allow_change_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["rr-attr-allow-change-vpnv4"], "RouterBgpNeighbor-RrAttrAllowChangeVpnv4"); ok {
			if err = d.Set("rr_attr_allow_change_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading rr_attr_allow_change_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rr_attr_allow_change_vpnv4: %v", err)
		}
	}

	if err = d.Set("rr_attr_allow_change_vpnv6", flattenRouterBgpNeighborRrAttrAllowChangeVpnv62edl(o["rr-attr-allow-change-vpnv6"], d, "rr_attr_allow_change_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["rr-attr-allow-change-vpnv6"], "RouterBgpNeighbor-RrAttrAllowChangeVpnv6"); ok {
			if err = d.Set("rr_attr_allow_change_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading rr_attr_allow_change_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rr_attr_allow_change_vpnv6: %v", err)
		}
	}

	if err = d.Set("rr_attr_allow_change6", flattenRouterBgpNeighborRrAttrAllowChange62edl(o["rr-attr-allow-change6"], d, "rr_attr_allow_change6")); err != nil {
		if vv, ok := fortiAPIPatch(o["rr-attr-allow-change6"], "RouterBgpNeighbor-RrAttrAllowChange6"); ok {
			if err = d.Set("rr_attr_allow_change6", vv); err != nil {
				return fmt.Errorf("Error reading rr_attr_allow_change6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rr_attr_allow_change6: %v", err)
		}
	}

	if err = d.Set("send_community", flattenRouterBgpNeighborSendCommunity2edl(o["send-community"], d, "send_community")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-community"], "RouterBgpNeighbor-SendCommunity"); ok {
			if err = d.Set("send_community", vv); err != nil {
				return fmt.Errorf("Error reading send_community: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_community: %v", err)
		}
	}

	if err = d.Set("send_community_evpn", flattenRouterBgpNeighborSendCommunityEvpn2edl(o["send-community-evpn"], d, "send_community_evpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-community-evpn"], "RouterBgpNeighbor-SendCommunityEvpn"); ok {
			if err = d.Set("send_community_evpn", vv); err != nil {
				return fmt.Errorf("Error reading send_community_evpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_community_evpn: %v", err)
		}
	}

	if err = d.Set("send_community_vpnv4", flattenRouterBgpNeighborSendCommunityVpnv42edl(o["send-community-vpnv4"], d, "send_community_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-community-vpnv4"], "RouterBgpNeighbor-SendCommunityVpnv4"); ok {
			if err = d.Set("send_community_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading send_community_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_community_vpnv4: %v", err)
		}
	}

	if err = d.Set("send_community_vpnv6", flattenRouterBgpNeighborSendCommunityVpnv62edl(o["send-community-vpnv6"], d, "send_community_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-community-vpnv6"], "RouterBgpNeighbor-SendCommunityVpnv6"); ok {
			if err = d.Set("send_community_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading send_community_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_community_vpnv6: %v", err)
		}
	}

	if err = d.Set("send_community6", flattenRouterBgpNeighborSendCommunity62edl(o["send-community6"], d, "send_community6")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-community6"], "RouterBgpNeighbor-SendCommunity6"); ok {
			if err = d.Set("send_community6", vv); err != nil {
				return fmt.Errorf("Error reading send_community6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_community6: %v", err)
		}
	}

	if err = d.Set("shutdown", flattenRouterBgpNeighborShutdown2edl(o["shutdown"], d, "shutdown")); err != nil {
		if vv, ok := fortiAPIPatch(o["shutdown"], "RouterBgpNeighbor-Shutdown"); ok {
			if err = d.Set("shutdown", vv); err != nil {
				return fmt.Errorf("Error reading shutdown: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading shutdown: %v", err)
		}
	}

	if err = d.Set("soft_reconfiguration", flattenRouterBgpNeighborSoftReconfiguration2edl(o["soft-reconfiguration"], d, "soft_reconfiguration")); err != nil {
		if vv, ok := fortiAPIPatch(o["soft-reconfiguration"], "RouterBgpNeighbor-SoftReconfiguration"); ok {
			if err = d.Set("soft_reconfiguration", vv); err != nil {
				return fmt.Errorf("Error reading soft_reconfiguration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading soft_reconfiguration: %v", err)
		}
	}

	if err = d.Set("soft_reconfiguration_evpn", flattenRouterBgpNeighborSoftReconfigurationEvpn2edl(o["soft-reconfiguration-evpn"], d, "soft_reconfiguration_evpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["soft-reconfiguration-evpn"], "RouterBgpNeighbor-SoftReconfigurationEvpn"); ok {
			if err = d.Set("soft_reconfiguration_evpn", vv); err != nil {
				return fmt.Errorf("Error reading soft_reconfiguration_evpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading soft_reconfiguration_evpn: %v", err)
		}
	}

	if err = d.Set("soft_reconfiguration_vpnv4", flattenRouterBgpNeighborSoftReconfigurationVpnv42edl(o["soft-reconfiguration-vpnv4"], d, "soft_reconfiguration_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["soft-reconfiguration-vpnv4"], "RouterBgpNeighbor-SoftReconfigurationVpnv4"); ok {
			if err = d.Set("soft_reconfiguration_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading soft_reconfiguration_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading soft_reconfiguration_vpnv4: %v", err)
		}
	}

	if err = d.Set("soft_reconfiguration_vpnv6", flattenRouterBgpNeighborSoftReconfigurationVpnv62edl(o["soft-reconfiguration-vpnv6"], d, "soft_reconfiguration_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["soft-reconfiguration-vpnv6"], "RouterBgpNeighbor-SoftReconfigurationVpnv6"); ok {
			if err = d.Set("soft_reconfiguration_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading soft_reconfiguration_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading soft_reconfiguration_vpnv6: %v", err)
		}
	}

	if err = d.Set("soft_reconfiguration6", flattenRouterBgpNeighborSoftReconfiguration62edl(o["soft-reconfiguration6"], d, "soft_reconfiguration6")); err != nil {
		if vv, ok := fortiAPIPatch(o["soft-reconfiguration6"], "RouterBgpNeighbor-SoftReconfiguration6"); ok {
			if err = d.Set("soft_reconfiguration6", vv); err != nil {
				return fmt.Errorf("Error reading soft_reconfiguration6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading soft_reconfiguration6: %v", err)
		}
	}

	if err = d.Set("stale_route", flattenRouterBgpNeighborStaleRoute2edl(o["stale-route"], d, "stale_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["stale-route"], "RouterBgpNeighbor-StaleRoute"); ok {
			if err = d.Set("stale_route", vv); err != nil {
				return fmt.Errorf("Error reading stale_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stale_route: %v", err)
		}
	}

	if err = d.Set("strict_capability_match", flattenRouterBgpNeighborStrictCapabilityMatch2edl(o["strict-capability-match"], d, "strict_capability_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["strict-capability-match"], "RouterBgpNeighbor-StrictCapabilityMatch"); ok {
			if err = d.Set("strict_capability_match", vv); err != nil {
				return fmt.Errorf("Error reading strict_capability_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading strict_capability_match: %v", err)
		}
	}

	if err = d.Set("unsuppress_map", flattenRouterBgpNeighborUnsuppressMap2edl(o["unsuppress-map"], d, "unsuppress_map")); err != nil {
		if vv, ok := fortiAPIPatch(o["unsuppress-map"], "RouterBgpNeighbor-UnsuppressMap"); ok {
			if err = d.Set("unsuppress_map", vv); err != nil {
				return fmt.Errorf("Error reading unsuppress_map: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unsuppress_map: %v", err)
		}
	}

	if err = d.Set("unsuppress_map6", flattenRouterBgpNeighborUnsuppressMap62edl(o["unsuppress-map6"], d, "unsuppress_map6")); err != nil {
		if vv, ok := fortiAPIPatch(o["unsuppress-map6"], "RouterBgpNeighbor-UnsuppressMap6"); ok {
			if err = d.Set("unsuppress_map6", vv); err != nil {
				return fmt.Errorf("Error reading unsuppress_map6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unsuppress_map6: %v", err)
		}
	}

	if err = d.Set("update_source", flattenRouterBgpNeighborUpdateSource2edl(o["update-source"], d, "update_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-source"], "RouterBgpNeighbor-UpdateSource"); ok {
			if err = d.Set("update_source", vv); err != nil {
				return fmt.Errorf("Error reading update_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_source: %v", err)
		}
	}

	if err = d.Set("weight", flattenRouterBgpNeighborWeight2edl(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "RouterBgpNeighbor-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	return nil
}

func flattenRouterBgpNeighborFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterBgpNeighborActivate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborActivateEvpn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborActivateVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborActivateVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborActivate62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdditionalPath2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdditionalPathVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdditionalPathVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdditionalPath62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdvAdditionalPath2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdvAdditionalPathVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdvAdditionalPathVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdvAdditionalPath62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdvertisementInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasIn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasInEnable2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasInEnableEvpn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasInEnableVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasInEnableVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasInEnable62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasInEvpn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasInVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasInVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasIn62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAsOverride2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAsOverride62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAttributeUnchanged2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborAttributeUnchangedVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborAttributeUnchangedVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborAttributeUnchanged62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborAuthOptions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborBfd2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityDefaultOriginate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityDefaultOriginate62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityDynamic2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityGracefulRestart2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityGracefulRestartEvpn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityGracefulRestartVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityGracefulRestartVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityGracefulRestart62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityOrf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityOrf62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityRouteRefresh2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborConditionalAdvertise2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["advertise-routemap"], _ = expandRouterBgpNeighborConditionalAdvertiseAdvertiseRoutemap2edl(d, i["advertise_routemap"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_routemap"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["condition-routemap"], _ = expandRouterBgpNeighborConditionalAdvertiseConditionRoutemap2edl(d, i["condition_routemap"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["condition-type"], _ = expandRouterBgpNeighborConditionalAdvertiseConditionType2edl(d, i["condition_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterBgpNeighborConditionalAdvertiseAdvertiseRoutemap2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandRouterBgpNeighborConditionalAdvertiseConditionRoutemap2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborConditionalAdvertiseConditionType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborConditionalAdvertise62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["advertise-routemap"], _ = expandRouterBgpNeighborConditionalAdvertise6AdvertiseRoutemap2edl(d, i["advertise_routemap"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_routemap"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["condition-routemap"], _ = expandRouterBgpNeighborConditionalAdvertise6ConditionRoutemap2edl(d, i["condition_routemap"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["condition-type"], _ = expandRouterBgpNeighborConditionalAdvertise6ConditionType2edl(d, i["condition_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterBgpNeighborConditionalAdvertise6AdvertiseRoutemap2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborConditionalAdvertise6ConditionRoutemap2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborConditionalAdvertise6ConditionType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborConnectTimer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborDefaultOriginateRoutemap2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborDefaultOriginateRoutemap62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborDistributeListIn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborDistributeListInVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborDistributeListInVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborDistributeListIn62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborDistributeListOut2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborDistributeListOutVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborDistributeListOutVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborDistributeListOut62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborDontCapabilityNegotiate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborEbgpEnforceMultihop2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborEbgpMultihopTtl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborFilterListIn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborFilterListInVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborFilterListInVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborFilterListIn62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborFilterListOut2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborFilterListOutVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborFilterListOutVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborFilterListOut62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborHoldtimeTimer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborKeepAliveTimer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborLinkDownFailover2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborLocalAs2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborLocalAsNoPrepend2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborLocalAsReplaceAs2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixEvpn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixThresholdEvpn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixThresholdVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixThresholdVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixThreshold62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixWarningOnly2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixWarningOnlyEvpn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixWarningOnlyVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixWarningOnlyVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixWarningOnly62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefix62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborNextHopSelf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborNextHopSelfRr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborNextHopSelfRr62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborNextHopSelfVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborNextHopSelfVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborNextHopSelf62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborOverrideCapability2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborPassive2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborPassword2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborPrefixListIn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborPrefixListInVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborPrefixListInVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborPrefixListIn62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborPrefixListOut2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborPrefixListOutVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborPrefixListOutVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborPrefixListOut62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRemoteAs2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRemovePrivateAs2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRemovePrivateAsEvpn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRemovePrivateAsVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRemovePrivateAsVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRemovePrivateAs62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRestartTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRetainStaleTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteMapIn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRouteMapInEvpn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRouteMapInVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRouteMapInVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRouteMapIn62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRouteMapOut2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRouteMapOutEvpn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRouteMapOutPreferable2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRouteMapOutVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRouteMapOutVpnv4Preferable2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRouteMapOutVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRouteMapOutVpnv6Preferable2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRouteMapOut62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRouteMapOut6Preferable2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRouteReflectorClient2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteReflectorClientEvpn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteReflectorClientVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteReflectorClientVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteReflectorClient62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteServerClient2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteServerClientEvpn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteServerClientVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteServerClientVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteServerClient62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRrAttrAllowChange2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRrAttrAllowChangeEvpn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRrAttrAllowChangeVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRrAttrAllowChangeVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRrAttrAllowChange62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSendCommunity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSendCommunityEvpn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSendCommunityVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSendCommunityVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSendCommunity62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborShutdown2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSoftReconfiguration2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSoftReconfigurationEvpn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSoftReconfigurationVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSoftReconfigurationVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSoftReconfiguration62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborStaleRoute2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborStrictCapabilityMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborUnsuppressMap2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborUnsuppressMap62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborUpdateSource2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterBgpNeighbor(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("activate"); ok || d.HasChange("activate") {
		t, err := expandRouterBgpNeighborActivate2edl(d, v, "activate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["activate"] = t
		}
	}

	if v, ok := d.GetOk("activate_evpn"); ok || d.HasChange("activate_evpn") {
		t, err := expandRouterBgpNeighborActivateEvpn2edl(d, v, "activate_evpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["activate-evpn"] = t
		}
	}

	if v, ok := d.GetOk("activate_vpnv4"); ok || d.HasChange("activate_vpnv4") {
		t, err := expandRouterBgpNeighborActivateVpnv42edl(d, v, "activate_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["activate-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("activate_vpnv6"); ok || d.HasChange("activate_vpnv6") {
		t, err := expandRouterBgpNeighborActivateVpnv62edl(d, v, "activate_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["activate-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("activate6"); ok || d.HasChange("activate6") {
		t, err := expandRouterBgpNeighborActivate62edl(d, v, "activate6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["activate6"] = t
		}
	}

	if v, ok := d.GetOk("additional_path"); ok || d.HasChange("additional_path") {
		t, err := expandRouterBgpNeighborAdditionalPath2edl(d, v, "additional_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-path"] = t
		}
	}

	if v, ok := d.GetOk("additional_path_vpnv4"); ok || d.HasChange("additional_path_vpnv4") {
		t, err := expandRouterBgpNeighborAdditionalPathVpnv42edl(d, v, "additional_path_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-path-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("additional_path_vpnv6"); ok || d.HasChange("additional_path_vpnv6") {
		t, err := expandRouterBgpNeighborAdditionalPathVpnv62edl(d, v, "additional_path_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-path-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("additional_path6"); ok || d.HasChange("additional_path6") {
		t, err := expandRouterBgpNeighborAdditionalPath62edl(d, v, "additional_path6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-path6"] = t
		}
	}

	if v, ok := d.GetOk("adv_additional_path"); ok || d.HasChange("adv_additional_path") {
		t, err := expandRouterBgpNeighborAdvAdditionalPath2edl(d, v, "adv_additional_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adv-additional-path"] = t
		}
	}

	if v, ok := d.GetOk("adv_additional_path_vpnv4"); ok || d.HasChange("adv_additional_path_vpnv4") {
		t, err := expandRouterBgpNeighborAdvAdditionalPathVpnv42edl(d, v, "adv_additional_path_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adv-additional-path-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("adv_additional_path_vpnv6"); ok || d.HasChange("adv_additional_path_vpnv6") {
		t, err := expandRouterBgpNeighborAdvAdditionalPathVpnv62edl(d, v, "adv_additional_path_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adv-additional-path-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("adv_additional_path6"); ok || d.HasChange("adv_additional_path6") {
		t, err := expandRouterBgpNeighborAdvAdditionalPath62edl(d, v, "adv_additional_path6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adv-additional-path6"] = t
		}
	}

	if v, ok := d.GetOk("advertisement_interval"); ok || d.HasChange("advertisement_interval") {
		t, err := expandRouterBgpNeighborAdvertisementInterval2edl(d, v, "advertisement_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["advertisement-interval"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in"); ok || d.HasChange("allowas_in") {
		t, err := expandRouterBgpNeighborAllowasIn2edl(d, v, "allowas_in")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in_enable"); ok || d.HasChange("allowas_in_enable") {
		t, err := expandRouterBgpNeighborAllowasInEnable2edl(d, v, "allowas_in_enable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in-enable"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in_enable_evpn"); ok || d.HasChange("allowas_in_enable_evpn") {
		t, err := expandRouterBgpNeighborAllowasInEnableEvpn2edl(d, v, "allowas_in_enable_evpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in-enable-evpn"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in_enable_vpnv4"); ok || d.HasChange("allowas_in_enable_vpnv4") {
		t, err := expandRouterBgpNeighborAllowasInEnableVpnv42edl(d, v, "allowas_in_enable_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in-enable-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in_enable_vpnv6"); ok || d.HasChange("allowas_in_enable_vpnv6") {
		t, err := expandRouterBgpNeighborAllowasInEnableVpnv62edl(d, v, "allowas_in_enable_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in-enable-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in_enable6"); ok || d.HasChange("allowas_in_enable6") {
		t, err := expandRouterBgpNeighborAllowasInEnable62edl(d, v, "allowas_in_enable6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in-enable6"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in_evpn"); ok || d.HasChange("allowas_in_evpn") {
		t, err := expandRouterBgpNeighborAllowasInEvpn2edl(d, v, "allowas_in_evpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in-evpn"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in_vpnv4"); ok || d.HasChange("allowas_in_vpnv4") {
		t, err := expandRouterBgpNeighborAllowasInVpnv42edl(d, v, "allowas_in_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in_vpnv6"); ok || d.HasChange("allowas_in_vpnv6") {
		t, err := expandRouterBgpNeighborAllowasInVpnv62edl(d, v, "allowas_in_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in6"); ok || d.HasChange("allowas_in6") {
		t, err := expandRouterBgpNeighborAllowasIn62edl(d, v, "allowas_in6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in6"] = t
		}
	}

	if v, ok := d.GetOk("as_override"); ok || d.HasChange("as_override") {
		t, err := expandRouterBgpNeighborAsOverride2edl(d, v, "as_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["as-override"] = t
		}
	}

	if v, ok := d.GetOk("as_override6"); ok || d.HasChange("as_override6") {
		t, err := expandRouterBgpNeighborAsOverride62edl(d, v, "as_override6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["as-override6"] = t
		}
	}

	if v, ok := d.GetOk("attribute_unchanged"); ok || d.HasChange("attribute_unchanged") {
		t, err := expandRouterBgpNeighborAttributeUnchanged2edl(d, v, "attribute_unchanged")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attribute-unchanged"] = t
		}
	}

	if v, ok := d.GetOk("attribute_unchanged_vpnv4"); ok || d.HasChange("attribute_unchanged_vpnv4") {
		t, err := expandRouterBgpNeighborAttributeUnchangedVpnv42edl(d, v, "attribute_unchanged_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attribute-unchanged-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("attribute_unchanged_vpnv6"); ok || d.HasChange("attribute_unchanged_vpnv6") {
		t, err := expandRouterBgpNeighborAttributeUnchangedVpnv62edl(d, v, "attribute_unchanged_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attribute-unchanged-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("attribute_unchanged6"); ok || d.HasChange("attribute_unchanged6") {
		t, err := expandRouterBgpNeighborAttributeUnchanged62edl(d, v, "attribute_unchanged6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attribute-unchanged6"] = t
		}
	}

	if v, ok := d.GetOk("auth_options"); ok || d.HasChange("auth_options") {
		t, err := expandRouterBgpNeighborAuthOptions2edl(d, v, "auth_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-options"] = t
		}
	}

	if v, ok := d.GetOk("bfd"); ok || d.HasChange("bfd") {
		t, err := expandRouterBgpNeighborBfd2edl(d, v, "bfd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd"] = t
		}
	}

	if v, ok := d.GetOk("capability_default_originate"); ok || d.HasChange("capability_default_originate") {
		t, err := expandRouterBgpNeighborCapabilityDefaultOriginate2edl(d, v, "capability_default_originate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-default-originate"] = t
		}
	}

	if v, ok := d.GetOk("capability_default_originate6"); ok || d.HasChange("capability_default_originate6") {
		t, err := expandRouterBgpNeighborCapabilityDefaultOriginate62edl(d, v, "capability_default_originate6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-default-originate6"] = t
		}
	}

	if v, ok := d.GetOk("capability_dynamic"); ok || d.HasChange("capability_dynamic") {
		t, err := expandRouterBgpNeighborCapabilityDynamic2edl(d, v, "capability_dynamic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-dynamic"] = t
		}
	}

	if v, ok := d.GetOk("capability_graceful_restart"); ok || d.HasChange("capability_graceful_restart") {
		t, err := expandRouterBgpNeighborCapabilityGracefulRestart2edl(d, v, "capability_graceful_restart")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-graceful-restart"] = t
		}
	}

	if v, ok := d.GetOk("capability_graceful_restart_evpn"); ok || d.HasChange("capability_graceful_restart_evpn") {
		t, err := expandRouterBgpNeighborCapabilityGracefulRestartEvpn2edl(d, v, "capability_graceful_restart_evpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-graceful-restart-evpn"] = t
		}
	}

	if v, ok := d.GetOk("capability_graceful_restart_vpnv4"); ok || d.HasChange("capability_graceful_restart_vpnv4") {
		t, err := expandRouterBgpNeighborCapabilityGracefulRestartVpnv42edl(d, v, "capability_graceful_restart_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-graceful-restart-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("capability_graceful_restart_vpnv6"); ok || d.HasChange("capability_graceful_restart_vpnv6") {
		t, err := expandRouterBgpNeighborCapabilityGracefulRestartVpnv62edl(d, v, "capability_graceful_restart_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-graceful-restart-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("capability_graceful_restart6"); ok || d.HasChange("capability_graceful_restart6") {
		t, err := expandRouterBgpNeighborCapabilityGracefulRestart62edl(d, v, "capability_graceful_restart6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-graceful-restart6"] = t
		}
	}

	if v, ok := d.GetOk("capability_orf"); ok || d.HasChange("capability_orf") {
		t, err := expandRouterBgpNeighborCapabilityOrf2edl(d, v, "capability_orf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-orf"] = t
		}
	}

	if v, ok := d.GetOk("capability_orf6"); ok || d.HasChange("capability_orf6") {
		t, err := expandRouterBgpNeighborCapabilityOrf62edl(d, v, "capability_orf6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-orf6"] = t
		}
	}

	if v, ok := d.GetOk("capability_route_refresh"); ok || d.HasChange("capability_route_refresh") {
		t, err := expandRouterBgpNeighborCapabilityRouteRefresh2edl(d, v, "capability_route_refresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-route-refresh"] = t
		}
	}

	if v, ok := d.GetOk("conditional_advertise"); ok || d.HasChange("conditional_advertise") {
		t, err := expandRouterBgpNeighborConditionalAdvertise2edl(d, v, "conditional_advertise")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conditional-advertise"] = t
		}
	}

	if v, ok := d.GetOk("conditional_advertise6"); ok || d.HasChange("conditional_advertise6") {
		t, err := expandRouterBgpNeighborConditionalAdvertise62edl(d, v, "conditional_advertise6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conditional-advertise6"] = t
		}
	}

	if v, ok := d.GetOk("connect_timer"); ok || d.HasChange("connect_timer") {
		t, err := expandRouterBgpNeighborConnectTimer2edl(d, v, "connect_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connect-timer"] = t
		}
	}

	if v, ok := d.GetOk("default_originate_routemap"); ok || d.HasChange("default_originate_routemap") {
		t, err := expandRouterBgpNeighborDefaultOriginateRoutemap2edl(d, v, "default_originate_routemap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-originate-routemap"] = t
		}
	}

	if v, ok := d.GetOk("default_originate_routemap6"); ok || d.HasChange("default_originate_routemap6") {
		t, err := expandRouterBgpNeighborDefaultOriginateRoutemap62edl(d, v, "default_originate_routemap6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-originate-routemap6"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandRouterBgpNeighborDescription2edl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list_in"); ok || d.HasChange("distribute_list_in") {
		t, err := expandRouterBgpNeighborDistributeListIn2edl(d, v, "distribute_list_in")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-in"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list_in_vpnv4"); ok || d.HasChange("distribute_list_in_vpnv4") {
		t, err := expandRouterBgpNeighborDistributeListInVpnv42edl(d, v, "distribute_list_in_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-in-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list_in_vpnv6"); ok || d.HasChange("distribute_list_in_vpnv6") {
		t, err := expandRouterBgpNeighborDistributeListInVpnv62edl(d, v, "distribute_list_in_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-in-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list_in6"); ok || d.HasChange("distribute_list_in6") {
		t, err := expandRouterBgpNeighborDistributeListIn62edl(d, v, "distribute_list_in6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-in6"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list_out"); ok || d.HasChange("distribute_list_out") {
		t, err := expandRouterBgpNeighborDistributeListOut2edl(d, v, "distribute_list_out")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-out"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list_out_vpnv4"); ok || d.HasChange("distribute_list_out_vpnv4") {
		t, err := expandRouterBgpNeighborDistributeListOutVpnv42edl(d, v, "distribute_list_out_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-out-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list_out_vpnv6"); ok || d.HasChange("distribute_list_out_vpnv6") {
		t, err := expandRouterBgpNeighborDistributeListOutVpnv62edl(d, v, "distribute_list_out_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-out-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list_out6"); ok || d.HasChange("distribute_list_out6") {
		t, err := expandRouterBgpNeighborDistributeListOut62edl(d, v, "distribute_list_out6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-out6"] = t
		}
	}

	if v, ok := d.GetOk("dont_capability_negotiate"); ok || d.HasChange("dont_capability_negotiate") {
		t, err := expandRouterBgpNeighborDontCapabilityNegotiate2edl(d, v, "dont_capability_negotiate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dont-capability-negotiate"] = t
		}
	}

	if v, ok := d.GetOk("ebgp_enforce_multihop"); ok || d.HasChange("ebgp_enforce_multihop") {
		t, err := expandRouterBgpNeighborEbgpEnforceMultihop2edl(d, v, "ebgp_enforce_multihop")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ebgp-enforce-multihop"] = t
		}
	}

	if v, ok := d.GetOk("ebgp_multihop_ttl"); ok || d.HasChange("ebgp_multihop_ttl") {
		t, err := expandRouterBgpNeighborEbgpMultihopTtl2edl(d, v, "ebgp_multihop_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ebgp-multihop-ttl"] = t
		}
	}

	if v, ok := d.GetOk("filter_list_in"); ok || d.HasChange("filter_list_in") {
		t, err := expandRouterBgpNeighborFilterListIn2edl(d, v, "filter_list_in")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-list-in"] = t
		}
	}

	if v, ok := d.GetOk("filter_list_in_vpnv4"); ok || d.HasChange("filter_list_in_vpnv4") {
		t, err := expandRouterBgpNeighborFilterListInVpnv42edl(d, v, "filter_list_in_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-list-in-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("filter_list_in_vpnv6"); ok || d.HasChange("filter_list_in_vpnv6") {
		t, err := expandRouterBgpNeighborFilterListInVpnv62edl(d, v, "filter_list_in_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-list-in-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("filter_list_in6"); ok || d.HasChange("filter_list_in6") {
		t, err := expandRouterBgpNeighborFilterListIn62edl(d, v, "filter_list_in6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-list-in6"] = t
		}
	}

	if v, ok := d.GetOk("filter_list_out"); ok || d.HasChange("filter_list_out") {
		t, err := expandRouterBgpNeighborFilterListOut2edl(d, v, "filter_list_out")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-list-out"] = t
		}
	}

	if v, ok := d.GetOk("filter_list_out_vpnv4"); ok || d.HasChange("filter_list_out_vpnv4") {
		t, err := expandRouterBgpNeighborFilterListOutVpnv42edl(d, v, "filter_list_out_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-list-out-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("filter_list_out_vpnv6"); ok || d.HasChange("filter_list_out_vpnv6") {
		t, err := expandRouterBgpNeighborFilterListOutVpnv62edl(d, v, "filter_list_out_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-list-out-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("filter_list_out6"); ok || d.HasChange("filter_list_out6") {
		t, err := expandRouterBgpNeighborFilterListOut62edl(d, v, "filter_list_out6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-list-out6"] = t
		}
	}

	if v, ok := d.GetOk("holdtime_timer"); ok || d.HasChange("holdtime_timer") {
		t, err := expandRouterBgpNeighborHoldtimeTimer2edl(d, v, "holdtime_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["holdtime-timer"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandRouterBgpNeighborInterface2edl(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandRouterBgpNeighborIp2edl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("keep_alive_timer"); ok || d.HasChange("keep_alive_timer") {
		t, err := expandRouterBgpNeighborKeepAliveTimer2edl(d, v, "keep_alive_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keep-alive-timer"] = t
		}
	}

	if v, ok := d.GetOk("link_down_failover"); ok || d.HasChange("link_down_failover") {
		t, err := expandRouterBgpNeighborLinkDownFailover2edl(d, v, "link_down_failover")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-down-failover"] = t
		}
	}

	if v, ok := d.GetOk("local_as"); ok || d.HasChange("local_as") {
		t, err := expandRouterBgpNeighborLocalAs2edl(d, v, "local_as")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-as"] = t
		}
	}

	if v, ok := d.GetOk("local_as_no_prepend"); ok || d.HasChange("local_as_no_prepend") {
		t, err := expandRouterBgpNeighborLocalAsNoPrepend2edl(d, v, "local_as_no_prepend")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-as-no-prepend"] = t
		}
	}

	if v, ok := d.GetOk("local_as_replace_as"); ok || d.HasChange("local_as_replace_as") {
		t, err := expandRouterBgpNeighborLocalAsReplaceAs2edl(d, v, "local_as_replace_as")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-as-replace-as"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix"); ok || d.HasChange("maximum_prefix") {
		t, err := expandRouterBgpNeighborMaximumPrefix2edl(d, v, "maximum_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_evpn"); ok || d.HasChange("maximum_prefix_evpn") {
		t, err := expandRouterBgpNeighborMaximumPrefixEvpn2edl(d, v, "maximum_prefix_evpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-evpn"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_threshold"); ok || d.HasChange("maximum_prefix_threshold") {
		t, err := expandRouterBgpNeighborMaximumPrefixThreshold2edl(d, v, "maximum_prefix_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-threshold"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_threshold_evpn"); ok || d.HasChange("maximum_prefix_threshold_evpn") {
		t, err := expandRouterBgpNeighborMaximumPrefixThresholdEvpn2edl(d, v, "maximum_prefix_threshold_evpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-threshold-evpn"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_threshold_vpnv4"); ok || d.HasChange("maximum_prefix_threshold_vpnv4") {
		t, err := expandRouterBgpNeighborMaximumPrefixThresholdVpnv42edl(d, v, "maximum_prefix_threshold_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-threshold-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_threshold_vpnv6"); ok || d.HasChange("maximum_prefix_threshold_vpnv6") {
		t, err := expandRouterBgpNeighborMaximumPrefixThresholdVpnv62edl(d, v, "maximum_prefix_threshold_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-threshold-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_threshold6"); ok || d.HasChange("maximum_prefix_threshold6") {
		t, err := expandRouterBgpNeighborMaximumPrefixThreshold62edl(d, v, "maximum_prefix_threshold6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-threshold6"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_vpnv4"); ok || d.HasChange("maximum_prefix_vpnv4") {
		t, err := expandRouterBgpNeighborMaximumPrefixVpnv42edl(d, v, "maximum_prefix_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_vpnv6"); ok || d.HasChange("maximum_prefix_vpnv6") {
		t, err := expandRouterBgpNeighborMaximumPrefixVpnv62edl(d, v, "maximum_prefix_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_warning_only"); ok || d.HasChange("maximum_prefix_warning_only") {
		t, err := expandRouterBgpNeighborMaximumPrefixWarningOnly2edl(d, v, "maximum_prefix_warning_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-warning-only"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_warning_only_evpn"); ok || d.HasChange("maximum_prefix_warning_only_evpn") {
		t, err := expandRouterBgpNeighborMaximumPrefixWarningOnlyEvpn2edl(d, v, "maximum_prefix_warning_only_evpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-warning-only-evpn"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_warning_only_vpnv4"); ok || d.HasChange("maximum_prefix_warning_only_vpnv4") {
		t, err := expandRouterBgpNeighborMaximumPrefixWarningOnlyVpnv42edl(d, v, "maximum_prefix_warning_only_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-warning-only-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_warning_only_vpnv6"); ok || d.HasChange("maximum_prefix_warning_only_vpnv6") {
		t, err := expandRouterBgpNeighborMaximumPrefixWarningOnlyVpnv62edl(d, v, "maximum_prefix_warning_only_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-warning-only-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_warning_only6"); ok || d.HasChange("maximum_prefix_warning_only6") {
		t, err := expandRouterBgpNeighborMaximumPrefixWarningOnly62edl(d, v, "maximum_prefix_warning_only6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-warning-only6"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix6"); ok || d.HasChange("maximum_prefix6") {
		t, err := expandRouterBgpNeighborMaximumPrefix62edl(d, v, "maximum_prefix6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix6"] = t
		}
	}

	if v, ok := d.GetOk("next_hop_self"); ok || d.HasChange("next_hop_self") {
		t, err := expandRouterBgpNeighborNextHopSelf2edl(d, v, "next_hop_self")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-hop-self"] = t
		}
	}

	if v, ok := d.GetOk("next_hop_self_rr"); ok || d.HasChange("next_hop_self_rr") {
		t, err := expandRouterBgpNeighborNextHopSelfRr2edl(d, v, "next_hop_self_rr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-hop-self-rr"] = t
		}
	}

	if v, ok := d.GetOk("next_hop_self_rr6"); ok || d.HasChange("next_hop_self_rr6") {
		t, err := expandRouterBgpNeighborNextHopSelfRr62edl(d, v, "next_hop_self_rr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-hop-self-rr6"] = t
		}
	}

	if v, ok := d.GetOk("next_hop_self_vpnv4"); ok || d.HasChange("next_hop_self_vpnv4") {
		t, err := expandRouterBgpNeighborNextHopSelfVpnv42edl(d, v, "next_hop_self_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-hop-self-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("next_hop_self_vpnv6"); ok || d.HasChange("next_hop_self_vpnv6") {
		t, err := expandRouterBgpNeighborNextHopSelfVpnv62edl(d, v, "next_hop_self_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-hop-self-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("next_hop_self6"); ok || d.HasChange("next_hop_self6") {
		t, err := expandRouterBgpNeighborNextHopSelf62edl(d, v, "next_hop_self6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-hop-self6"] = t
		}
	}

	if v, ok := d.GetOk("override_capability"); ok || d.HasChange("override_capability") {
		t, err := expandRouterBgpNeighborOverrideCapability2edl(d, v, "override_capability")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-capability"] = t
		}
	}

	if v, ok := d.GetOk("passive"); ok || d.HasChange("passive") {
		t, err := expandRouterBgpNeighborPassive2edl(d, v, "passive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passive"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandRouterBgpNeighborPassword2edl(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("prefix_list_in"); ok || d.HasChange("prefix_list_in") {
		t, err := expandRouterBgpNeighborPrefixListIn2edl(d, v, "prefix_list_in")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-in"] = t
		}
	}

	if v, ok := d.GetOk("prefix_list_in_vpnv4"); ok || d.HasChange("prefix_list_in_vpnv4") {
		t, err := expandRouterBgpNeighborPrefixListInVpnv42edl(d, v, "prefix_list_in_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-in-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("prefix_list_in_vpnv6"); ok || d.HasChange("prefix_list_in_vpnv6") {
		t, err := expandRouterBgpNeighborPrefixListInVpnv62edl(d, v, "prefix_list_in_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-in-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("prefix_list_in6"); ok || d.HasChange("prefix_list_in6") {
		t, err := expandRouterBgpNeighborPrefixListIn62edl(d, v, "prefix_list_in6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-in6"] = t
		}
	}

	if v, ok := d.GetOk("prefix_list_out"); ok || d.HasChange("prefix_list_out") {
		t, err := expandRouterBgpNeighborPrefixListOut2edl(d, v, "prefix_list_out")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-out"] = t
		}
	}

	if v, ok := d.GetOk("prefix_list_out_vpnv4"); ok || d.HasChange("prefix_list_out_vpnv4") {
		t, err := expandRouterBgpNeighborPrefixListOutVpnv42edl(d, v, "prefix_list_out_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-out-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("prefix_list_out_vpnv6"); ok || d.HasChange("prefix_list_out_vpnv6") {
		t, err := expandRouterBgpNeighborPrefixListOutVpnv62edl(d, v, "prefix_list_out_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-out-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("prefix_list_out6"); ok || d.HasChange("prefix_list_out6") {
		t, err := expandRouterBgpNeighborPrefixListOut62edl(d, v, "prefix_list_out6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-out6"] = t
		}
	}

	if v, ok := d.GetOk("remote_as"); ok || d.HasChange("remote_as") {
		t, err := expandRouterBgpNeighborRemoteAs2edl(d, v, "remote_as")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-as"] = t
		}
	}

	if v, ok := d.GetOk("remove_private_as"); ok || d.HasChange("remove_private_as") {
		t, err := expandRouterBgpNeighborRemovePrivateAs2edl(d, v, "remove_private_as")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remove-private-as"] = t
		}
	}

	if v, ok := d.GetOk("remove_private_as_evpn"); ok || d.HasChange("remove_private_as_evpn") {
		t, err := expandRouterBgpNeighborRemovePrivateAsEvpn2edl(d, v, "remove_private_as_evpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remove-private-as-evpn"] = t
		}
	}

	if v, ok := d.GetOk("remove_private_as_vpnv4"); ok || d.HasChange("remove_private_as_vpnv4") {
		t, err := expandRouterBgpNeighborRemovePrivateAsVpnv42edl(d, v, "remove_private_as_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remove-private-as-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("remove_private_as_vpnv6"); ok || d.HasChange("remove_private_as_vpnv6") {
		t, err := expandRouterBgpNeighborRemovePrivateAsVpnv62edl(d, v, "remove_private_as_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remove-private-as-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("remove_private_as6"); ok || d.HasChange("remove_private_as6") {
		t, err := expandRouterBgpNeighborRemovePrivateAs62edl(d, v, "remove_private_as6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remove-private-as6"] = t
		}
	}

	if v, ok := d.GetOk("restart_time"); ok || d.HasChange("restart_time") {
		t, err := expandRouterBgpNeighborRestartTime2edl(d, v, "restart_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restart-time"] = t
		}
	}

	if v, ok := d.GetOk("retain_stale_time"); ok || d.HasChange("retain_stale_time") {
		t, err := expandRouterBgpNeighborRetainStaleTime2edl(d, v, "retain_stale_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retain-stale-time"] = t
		}
	}

	if v, ok := d.GetOk("route_map_in"); ok || d.HasChange("route_map_in") {
		t, err := expandRouterBgpNeighborRouteMapIn2edl(d, v, "route_map_in")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-in"] = t
		}
	}

	if v, ok := d.GetOk("route_map_in_evpn"); ok || d.HasChange("route_map_in_evpn") {
		t, err := expandRouterBgpNeighborRouteMapInEvpn2edl(d, v, "route_map_in_evpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-in-evpn"] = t
		}
	}

	if v, ok := d.GetOk("route_map_in_vpnv4"); ok || d.HasChange("route_map_in_vpnv4") {
		t, err := expandRouterBgpNeighborRouteMapInVpnv42edl(d, v, "route_map_in_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-in-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("route_map_in_vpnv6"); ok || d.HasChange("route_map_in_vpnv6") {
		t, err := expandRouterBgpNeighborRouteMapInVpnv62edl(d, v, "route_map_in_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-in-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("route_map_in6"); ok || d.HasChange("route_map_in6") {
		t, err := expandRouterBgpNeighborRouteMapIn62edl(d, v, "route_map_in6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-in6"] = t
		}
	}

	if v, ok := d.GetOk("route_map_out"); ok || d.HasChange("route_map_out") {
		t, err := expandRouterBgpNeighborRouteMapOut2edl(d, v, "route_map_out")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out"] = t
		}
	}

	if v, ok := d.GetOk("route_map_out_evpn"); ok || d.HasChange("route_map_out_evpn") {
		t, err := expandRouterBgpNeighborRouteMapOutEvpn2edl(d, v, "route_map_out_evpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out-evpn"] = t
		}
	}

	if v, ok := d.GetOk("route_map_out_preferable"); ok || d.HasChange("route_map_out_preferable") {
		t, err := expandRouterBgpNeighborRouteMapOutPreferable2edl(d, v, "route_map_out_preferable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out-preferable"] = t
		}
	}

	if v, ok := d.GetOk("route_map_out_vpnv4"); ok || d.HasChange("route_map_out_vpnv4") {
		t, err := expandRouterBgpNeighborRouteMapOutVpnv42edl(d, v, "route_map_out_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("route_map_out_vpnv4_preferable"); ok || d.HasChange("route_map_out_vpnv4_preferable") {
		t, err := expandRouterBgpNeighborRouteMapOutVpnv4Preferable2edl(d, v, "route_map_out_vpnv4_preferable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out-vpnv4-preferable"] = t
		}
	}

	if v, ok := d.GetOk("route_map_out_vpnv6"); ok || d.HasChange("route_map_out_vpnv6") {
		t, err := expandRouterBgpNeighborRouteMapOutVpnv62edl(d, v, "route_map_out_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("route_map_out_vpnv6_preferable"); ok || d.HasChange("route_map_out_vpnv6_preferable") {
		t, err := expandRouterBgpNeighborRouteMapOutVpnv6Preferable2edl(d, v, "route_map_out_vpnv6_preferable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out-vpnv6-preferable"] = t
		}
	}

	if v, ok := d.GetOk("route_map_out6"); ok || d.HasChange("route_map_out6") {
		t, err := expandRouterBgpNeighborRouteMapOut62edl(d, v, "route_map_out6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out6"] = t
		}
	}

	if v, ok := d.GetOk("route_map_out6_preferable"); ok || d.HasChange("route_map_out6_preferable") {
		t, err := expandRouterBgpNeighborRouteMapOut6Preferable2edl(d, v, "route_map_out6_preferable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out6-preferable"] = t
		}
	}

	if v, ok := d.GetOk("route_reflector_client"); ok || d.HasChange("route_reflector_client") {
		t, err := expandRouterBgpNeighborRouteReflectorClient2edl(d, v, "route_reflector_client")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-reflector-client"] = t
		}
	}

	if v, ok := d.GetOk("route_reflector_client_evpn"); ok || d.HasChange("route_reflector_client_evpn") {
		t, err := expandRouterBgpNeighborRouteReflectorClientEvpn2edl(d, v, "route_reflector_client_evpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-reflector-client-evpn"] = t
		}
	}

	if v, ok := d.GetOk("route_reflector_client_vpnv4"); ok || d.HasChange("route_reflector_client_vpnv4") {
		t, err := expandRouterBgpNeighborRouteReflectorClientVpnv42edl(d, v, "route_reflector_client_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-reflector-client-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("route_reflector_client_vpnv6"); ok || d.HasChange("route_reflector_client_vpnv6") {
		t, err := expandRouterBgpNeighborRouteReflectorClientVpnv62edl(d, v, "route_reflector_client_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-reflector-client-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("route_reflector_client6"); ok || d.HasChange("route_reflector_client6") {
		t, err := expandRouterBgpNeighborRouteReflectorClient62edl(d, v, "route_reflector_client6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-reflector-client6"] = t
		}
	}

	if v, ok := d.GetOk("route_server_client"); ok || d.HasChange("route_server_client") {
		t, err := expandRouterBgpNeighborRouteServerClient2edl(d, v, "route_server_client")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-server-client"] = t
		}
	}

	if v, ok := d.GetOk("route_server_client_evpn"); ok || d.HasChange("route_server_client_evpn") {
		t, err := expandRouterBgpNeighborRouteServerClientEvpn2edl(d, v, "route_server_client_evpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-server-client-evpn"] = t
		}
	}

	if v, ok := d.GetOk("route_server_client_vpnv4"); ok || d.HasChange("route_server_client_vpnv4") {
		t, err := expandRouterBgpNeighborRouteServerClientVpnv42edl(d, v, "route_server_client_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-server-client-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("route_server_client_vpnv6"); ok || d.HasChange("route_server_client_vpnv6") {
		t, err := expandRouterBgpNeighborRouteServerClientVpnv62edl(d, v, "route_server_client_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-server-client-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("route_server_client6"); ok || d.HasChange("route_server_client6") {
		t, err := expandRouterBgpNeighborRouteServerClient62edl(d, v, "route_server_client6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-server-client6"] = t
		}
	}

	if v, ok := d.GetOk("rr_attr_allow_change"); ok || d.HasChange("rr_attr_allow_change") {
		t, err := expandRouterBgpNeighborRrAttrAllowChange2edl(d, v, "rr_attr_allow_change")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rr-attr-allow-change"] = t
		}
	}

	if v, ok := d.GetOk("rr_attr_allow_change_evpn"); ok || d.HasChange("rr_attr_allow_change_evpn") {
		t, err := expandRouterBgpNeighborRrAttrAllowChangeEvpn2edl(d, v, "rr_attr_allow_change_evpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rr-attr-allow-change-evpn"] = t
		}
	}

	if v, ok := d.GetOk("rr_attr_allow_change_vpnv4"); ok || d.HasChange("rr_attr_allow_change_vpnv4") {
		t, err := expandRouterBgpNeighborRrAttrAllowChangeVpnv42edl(d, v, "rr_attr_allow_change_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rr-attr-allow-change-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("rr_attr_allow_change_vpnv6"); ok || d.HasChange("rr_attr_allow_change_vpnv6") {
		t, err := expandRouterBgpNeighborRrAttrAllowChangeVpnv62edl(d, v, "rr_attr_allow_change_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rr-attr-allow-change-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("rr_attr_allow_change6"); ok || d.HasChange("rr_attr_allow_change6") {
		t, err := expandRouterBgpNeighborRrAttrAllowChange62edl(d, v, "rr_attr_allow_change6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rr-attr-allow-change6"] = t
		}
	}

	if v, ok := d.GetOk("send_community"); ok || d.HasChange("send_community") {
		t, err := expandRouterBgpNeighborSendCommunity2edl(d, v, "send_community")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-community"] = t
		}
	}

	if v, ok := d.GetOk("send_community_evpn"); ok || d.HasChange("send_community_evpn") {
		t, err := expandRouterBgpNeighborSendCommunityEvpn2edl(d, v, "send_community_evpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-community-evpn"] = t
		}
	}

	if v, ok := d.GetOk("send_community_vpnv4"); ok || d.HasChange("send_community_vpnv4") {
		t, err := expandRouterBgpNeighborSendCommunityVpnv42edl(d, v, "send_community_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-community-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("send_community_vpnv6"); ok || d.HasChange("send_community_vpnv6") {
		t, err := expandRouterBgpNeighborSendCommunityVpnv62edl(d, v, "send_community_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-community-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("send_community6"); ok || d.HasChange("send_community6") {
		t, err := expandRouterBgpNeighborSendCommunity62edl(d, v, "send_community6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-community6"] = t
		}
	}

	if v, ok := d.GetOk("shutdown"); ok || d.HasChange("shutdown") {
		t, err := expandRouterBgpNeighborShutdown2edl(d, v, "shutdown")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shutdown"] = t
		}
	}

	if v, ok := d.GetOk("soft_reconfiguration"); ok || d.HasChange("soft_reconfiguration") {
		t, err := expandRouterBgpNeighborSoftReconfiguration2edl(d, v, "soft_reconfiguration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["soft-reconfiguration"] = t
		}
	}

	if v, ok := d.GetOk("soft_reconfiguration_evpn"); ok || d.HasChange("soft_reconfiguration_evpn") {
		t, err := expandRouterBgpNeighborSoftReconfigurationEvpn2edl(d, v, "soft_reconfiguration_evpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["soft-reconfiguration-evpn"] = t
		}
	}

	if v, ok := d.GetOk("soft_reconfiguration_vpnv4"); ok || d.HasChange("soft_reconfiguration_vpnv4") {
		t, err := expandRouterBgpNeighborSoftReconfigurationVpnv42edl(d, v, "soft_reconfiguration_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["soft-reconfiguration-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("soft_reconfiguration_vpnv6"); ok || d.HasChange("soft_reconfiguration_vpnv6") {
		t, err := expandRouterBgpNeighborSoftReconfigurationVpnv62edl(d, v, "soft_reconfiguration_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["soft-reconfiguration-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("soft_reconfiguration6"); ok || d.HasChange("soft_reconfiguration6") {
		t, err := expandRouterBgpNeighborSoftReconfiguration62edl(d, v, "soft_reconfiguration6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["soft-reconfiguration6"] = t
		}
	}

	if v, ok := d.GetOk("stale_route"); ok || d.HasChange("stale_route") {
		t, err := expandRouterBgpNeighborStaleRoute2edl(d, v, "stale_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stale-route"] = t
		}
	}

	if v, ok := d.GetOk("strict_capability_match"); ok || d.HasChange("strict_capability_match") {
		t, err := expandRouterBgpNeighborStrictCapabilityMatch2edl(d, v, "strict_capability_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strict-capability-match"] = t
		}
	}

	if v, ok := d.GetOk("unsuppress_map"); ok || d.HasChange("unsuppress_map") {
		t, err := expandRouterBgpNeighborUnsuppressMap2edl(d, v, "unsuppress_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unsuppress-map"] = t
		}
	}

	if v, ok := d.GetOk("unsuppress_map6"); ok || d.HasChange("unsuppress_map6") {
		t, err := expandRouterBgpNeighborUnsuppressMap62edl(d, v, "unsuppress_map6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unsuppress-map6"] = t
		}
	}

	if v, ok := d.GetOk("update_source"); ok || d.HasChange("update_source") {
		t, err := expandRouterBgpNeighborUpdateSource2edl(d, v, "update_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-source"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandRouterBgpNeighborWeight2edl(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	return &obj, nil
}
