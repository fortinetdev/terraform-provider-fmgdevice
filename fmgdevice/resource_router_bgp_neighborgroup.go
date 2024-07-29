// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: BGP neighbor group table.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterBgpNeighborGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterBgpNeighborGroupCreate,
		Read:   resourceRouterBgpNeighborGroupRead,
		Update: resourceRouterBgpNeighborGroupUpdate,
		Delete: resourceRouterBgpNeighborGroupDelete,

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
				ForceNew: true,
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
	}
}

func resourceRouterBgpNeighborGroupCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterBgpNeighborGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterBgpNeighborGroup resource while getting object: %v", err)
	}

	_, err = c.CreateRouterBgpNeighborGroup(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating RouterBgpNeighborGroup resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceRouterBgpNeighborGroupRead(d, m)
}

func resourceRouterBgpNeighborGroupUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterBgpNeighborGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgpNeighborGroup resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterBgpNeighborGroup(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgpNeighborGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceRouterBgpNeighborGroupRead(d, m)
}

func resourceRouterBgpNeighborGroupDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterBgpNeighborGroup(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting RouterBgpNeighborGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterBgpNeighborGroupRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterBgpNeighborGroup(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgpNeighborGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterBgpNeighborGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgpNeighborGroup resource from API: %v", err)
	}
	return nil
}

func flattenRouterBgpNeighborGroupActivate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupActivateEvpn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupActivateVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupActivateVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupActivate62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAdditionalPath2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAdditionalPathVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAdditionalPathVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAdditionalPath62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAdvAdditionalPath2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAdvAdditionalPathVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAdvAdditionalPathVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAdvAdditionalPath62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAdvertisementInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAllowasIn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAllowasInEnable2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAllowasInEnableEvpn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAllowasInEnableVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAllowasInEnableVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAllowasInEnable62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAllowasInEvpn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAllowasInVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAllowasInVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAllowasIn62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAsOverride2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAsOverride62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAttributeUnchanged2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupAttributeUnchangedVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupAttributeUnchangedVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupAttributeUnchanged62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupAuthOptions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupBfd2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityDefaultOriginate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityDefaultOriginate62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityDynamic2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityGracefulRestart2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityGracefulRestartEvpn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityGracefulRestartVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityGracefulRestartVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityGracefulRestart62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityOrf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityOrf62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityRouteRefresh2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupConnectTimer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupDefaultOriginateRoutemap2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupDefaultOriginateRoutemap62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupDistributeListIn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupDistributeListInVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupDistributeListInVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupDistributeListIn62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupDistributeListOut2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupDistributeListOutVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupDistributeListOutVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupDistributeListOut62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupDontCapabilityNegotiate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupEbgpEnforceMultihop2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupEbgpMultihopTtl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupFilterListIn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupFilterListInVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupFilterListInVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupFilterListIn62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupFilterListOut2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupFilterListOutVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupFilterListOutVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupFilterListOut62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupHoldtimeTimer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupKeepAliveTimer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupLinkDownFailover2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupLocalAs2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupLocalAsNoPrepend2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupLocalAsReplaceAs2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixEvpn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixThresholdEvpn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixThresholdVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixThresholdVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixThreshold62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixWarningOnly2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixWarningOnlyEvpn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixWarningOnlyVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixWarningOnlyVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixWarningOnly62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefix62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupNextHopSelf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupNextHopSelfRr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupNextHopSelfRr62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupNextHopSelfVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupNextHopSelfVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupNextHopSelf62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupOverrideCapability2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupPassive2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupPrefixListIn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupPrefixListInVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupPrefixListInVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupPrefixListIn62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupPrefixListOut2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupPrefixListOutVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupPrefixListOutVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupPrefixListOut62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRemoteAs2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRemoteAsFilter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRemovePrivateAs2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRemovePrivateAsEvpn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRemovePrivateAsVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRemovePrivateAsVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRemovePrivateAs62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRestartTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRetainStaleTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteMapIn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRouteMapInEvpn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRouteMapInVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRouteMapInVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRouteMapIn62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRouteMapOut2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRouteMapOutEvpn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRouteMapOutPreferable2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRouteMapOutVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRouteMapOutVpnv4Preferable2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRouteMapOutVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRouteMapOutVpnv6Preferable2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRouteMapOut62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRouteMapOut6Preferable2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupRouteReflectorClient2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteReflectorClientEvpn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteReflectorClientVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteReflectorClientVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteReflectorClient62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteServerClient2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteServerClientEvpn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteServerClientVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteServerClientVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteServerClient62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSendCommunity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSendCommunityEvpn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSendCommunityVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSendCommunityVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSendCommunity62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupShutdown2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSoftReconfiguration2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSoftReconfigurationEvpn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSoftReconfigurationVpnv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSoftReconfigurationVpnv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSoftReconfiguration62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupStaleRoute2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupStrictCapabilityMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupUnsuppressMap2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupUnsuppressMap62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupUpdateSource2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborGroupWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterBgpNeighborGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("activate", flattenRouterBgpNeighborGroupActivate2edl(o["activate"], d, "activate")); err != nil {
		if vv, ok := fortiAPIPatch(o["activate"], "RouterBgpNeighborGroup-Activate"); ok {
			if err = d.Set("activate", vv); err != nil {
				return fmt.Errorf("Error reading activate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading activate: %v", err)
		}
	}

	if err = d.Set("activate_evpn", flattenRouterBgpNeighborGroupActivateEvpn2edl(o["activate-evpn"], d, "activate_evpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["activate-evpn"], "RouterBgpNeighborGroup-ActivateEvpn"); ok {
			if err = d.Set("activate_evpn", vv); err != nil {
				return fmt.Errorf("Error reading activate_evpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading activate_evpn: %v", err)
		}
	}

	if err = d.Set("activate_vpnv4", flattenRouterBgpNeighborGroupActivateVpnv42edl(o["activate-vpnv4"], d, "activate_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["activate-vpnv4"], "RouterBgpNeighborGroup-ActivateVpnv4"); ok {
			if err = d.Set("activate_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading activate_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading activate_vpnv4: %v", err)
		}
	}

	if err = d.Set("activate_vpnv6", flattenRouterBgpNeighborGroupActivateVpnv62edl(o["activate-vpnv6"], d, "activate_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["activate-vpnv6"], "RouterBgpNeighborGroup-ActivateVpnv6"); ok {
			if err = d.Set("activate_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading activate_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading activate_vpnv6: %v", err)
		}
	}

	if err = d.Set("activate6", flattenRouterBgpNeighborGroupActivate62edl(o["activate6"], d, "activate6")); err != nil {
		if vv, ok := fortiAPIPatch(o["activate6"], "RouterBgpNeighborGroup-Activate6"); ok {
			if err = d.Set("activate6", vv); err != nil {
				return fmt.Errorf("Error reading activate6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading activate6: %v", err)
		}
	}

	if err = d.Set("additional_path", flattenRouterBgpNeighborGroupAdditionalPath2edl(o["additional-path"], d, "additional_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["additional-path"], "RouterBgpNeighborGroup-AdditionalPath"); ok {
			if err = d.Set("additional_path", vv); err != nil {
				return fmt.Errorf("Error reading additional_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading additional_path: %v", err)
		}
	}

	if err = d.Set("additional_path_vpnv4", flattenRouterBgpNeighborGroupAdditionalPathVpnv42edl(o["additional-path-vpnv4"], d, "additional_path_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["additional-path-vpnv4"], "RouterBgpNeighborGroup-AdditionalPathVpnv4"); ok {
			if err = d.Set("additional_path_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading additional_path_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading additional_path_vpnv4: %v", err)
		}
	}

	if err = d.Set("additional_path_vpnv6", flattenRouterBgpNeighborGroupAdditionalPathVpnv62edl(o["additional-path-vpnv6"], d, "additional_path_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["additional-path-vpnv6"], "RouterBgpNeighborGroup-AdditionalPathVpnv6"); ok {
			if err = d.Set("additional_path_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading additional_path_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading additional_path_vpnv6: %v", err)
		}
	}

	if err = d.Set("additional_path6", flattenRouterBgpNeighborGroupAdditionalPath62edl(o["additional-path6"], d, "additional_path6")); err != nil {
		if vv, ok := fortiAPIPatch(o["additional-path6"], "RouterBgpNeighborGroup-AdditionalPath6"); ok {
			if err = d.Set("additional_path6", vv); err != nil {
				return fmt.Errorf("Error reading additional_path6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading additional_path6: %v", err)
		}
	}

	if err = d.Set("adv_additional_path", flattenRouterBgpNeighborGroupAdvAdditionalPath2edl(o["adv-additional-path"], d, "adv_additional_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["adv-additional-path"], "RouterBgpNeighborGroup-AdvAdditionalPath"); ok {
			if err = d.Set("adv_additional_path", vv); err != nil {
				return fmt.Errorf("Error reading adv_additional_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adv_additional_path: %v", err)
		}
	}

	if err = d.Set("adv_additional_path_vpnv4", flattenRouterBgpNeighborGroupAdvAdditionalPathVpnv42edl(o["adv-additional-path-vpnv4"], d, "adv_additional_path_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["adv-additional-path-vpnv4"], "RouterBgpNeighborGroup-AdvAdditionalPathVpnv4"); ok {
			if err = d.Set("adv_additional_path_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading adv_additional_path_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adv_additional_path_vpnv4: %v", err)
		}
	}

	if err = d.Set("adv_additional_path_vpnv6", flattenRouterBgpNeighborGroupAdvAdditionalPathVpnv62edl(o["adv-additional-path-vpnv6"], d, "adv_additional_path_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["adv-additional-path-vpnv6"], "RouterBgpNeighborGroup-AdvAdditionalPathVpnv6"); ok {
			if err = d.Set("adv_additional_path_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading adv_additional_path_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adv_additional_path_vpnv6: %v", err)
		}
	}

	if err = d.Set("adv_additional_path6", flattenRouterBgpNeighborGroupAdvAdditionalPath62edl(o["adv-additional-path6"], d, "adv_additional_path6")); err != nil {
		if vv, ok := fortiAPIPatch(o["adv-additional-path6"], "RouterBgpNeighborGroup-AdvAdditionalPath6"); ok {
			if err = d.Set("adv_additional_path6", vv); err != nil {
				return fmt.Errorf("Error reading adv_additional_path6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adv_additional_path6: %v", err)
		}
	}

	if err = d.Set("advertisement_interval", flattenRouterBgpNeighborGroupAdvertisementInterval2edl(o["advertisement-interval"], d, "advertisement_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["advertisement-interval"], "RouterBgpNeighborGroup-AdvertisementInterval"); ok {
			if err = d.Set("advertisement_interval", vv); err != nil {
				return fmt.Errorf("Error reading advertisement_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading advertisement_interval: %v", err)
		}
	}

	if err = d.Set("allowas_in", flattenRouterBgpNeighborGroupAllowasIn2edl(o["allowas-in"], d, "allowas_in")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowas-in"], "RouterBgpNeighborGroup-AllowasIn"); ok {
			if err = d.Set("allowas_in", vv); err != nil {
				return fmt.Errorf("Error reading allowas_in: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowas_in: %v", err)
		}
	}

	if err = d.Set("allowas_in_enable", flattenRouterBgpNeighborGroupAllowasInEnable2edl(o["allowas-in-enable"], d, "allowas_in_enable")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowas-in-enable"], "RouterBgpNeighborGroup-AllowasInEnable"); ok {
			if err = d.Set("allowas_in_enable", vv); err != nil {
				return fmt.Errorf("Error reading allowas_in_enable: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowas_in_enable: %v", err)
		}
	}

	if err = d.Set("allowas_in_enable_evpn", flattenRouterBgpNeighborGroupAllowasInEnableEvpn2edl(o["allowas-in-enable-evpn"], d, "allowas_in_enable_evpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowas-in-enable-evpn"], "RouterBgpNeighborGroup-AllowasInEnableEvpn"); ok {
			if err = d.Set("allowas_in_enable_evpn", vv); err != nil {
				return fmt.Errorf("Error reading allowas_in_enable_evpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowas_in_enable_evpn: %v", err)
		}
	}

	if err = d.Set("allowas_in_enable_vpnv4", flattenRouterBgpNeighborGroupAllowasInEnableVpnv42edl(o["allowas-in-enable-vpnv4"], d, "allowas_in_enable_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowas-in-enable-vpnv4"], "RouterBgpNeighborGroup-AllowasInEnableVpnv4"); ok {
			if err = d.Set("allowas_in_enable_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading allowas_in_enable_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowas_in_enable_vpnv4: %v", err)
		}
	}

	if err = d.Set("allowas_in_enable_vpnv6", flattenRouterBgpNeighborGroupAllowasInEnableVpnv62edl(o["allowas-in-enable-vpnv6"], d, "allowas_in_enable_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowas-in-enable-vpnv6"], "RouterBgpNeighborGroup-AllowasInEnableVpnv6"); ok {
			if err = d.Set("allowas_in_enable_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading allowas_in_enable_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowas_in_enable_vpnv6: %v", err)
		}
	}

	if err = d.Set("allowas_in_enable6", flattenRouterBgpNeighborGroupAllowasInEnable62edl(o["allowas-in-enable6"], d, "allowas_in_enable6")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowas-in-enable6"], "RouterBgpNeighborGroup-AllowasInEnable6"); ok {
			if err = d.Set("allowas_in_enable6", vv); err != nil {
				return fmt.Errorf("Error reading allowas_in_enable6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowas_in_enable6: %v", err)
		}
	}

	if err = d.Set("allowas_in_evpn", flattenRouterBgpNeighborGroupAllowasInEvpn2edl(o["allowas-in-evpn"], d, "allowas_in_evpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowas-in-evpn"], "RouterBgpNeighborGroup-AllowasInEvpn"); ok {
			if err = d.Set("allowas_in_evpn", vv); err != nil {
				return fmt.Errorf("Error reading allowas_in_evpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowas_in_evpn: %v", err)
		}
	}

	if err = d.Set("allowas_in_vpnv4", flattenRouterBgpNeighborGroupAllowasInVpnv42edl(o["allowas-in-vpnv4"], d, "allowas_in_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowas-in-vpnv4"], "RouterBgpNeighborGroup-AllowasInVpnv4"); ok {
			if err = d.Set("allowas_in_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading allowas_in_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowas_in_vpnv4: %v", err)
		}
	}

	if err = d.Set("allowas_in_vpnv6", flattenRouterBgpNeighborGroupAllowasInVpnv62edl(o["allowas-in-vpnv6"], d, "allowas_in_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowas-in-vpnv6"], "RouterBgpNeighborGroup-AllowasInVpnv6"); ok {
			if err = d.Set("allowas_in_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading allowas_in_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowas_in_vpnv6: %v", err)
		}
	}

	if err = d.Set("allowas_in6", flattenRouterBgpNeighborGroupAllowasIn62edl(o["allowas-in6"], d, "allowas_in6")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowas-in6"], "RouterBgpNeighborGroup-AllowasIn6"); ok {
			if err = d.Set("allowas_in6", vv); err != nil {
				return fmt.Errorf("Error reading allowas_in6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowas_in6: %v", err)
		}
	}

	if err = d.Set("as_override", flattenRouterBgpNeighborGroupAsOverride2edl(o["as-override"], d, "as_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["as-override"], "RouterBgpNeighborGroup-AsOverride"); ok {
			if err = d.Set("as_override", vv); err != nil {
				return fmt.Errorf("Error reading as_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading as_override: %v", err)
		}
	}

	if err = d.Set("as_override6", flattenRouterBgpNeighborGroupAsOverride62edl(o["as-override6"], d, "as_override6")); err != nil {
		if vv, ok := fortiAPIPatch(o["as-override6"], "RouterBgpNeighborGroup-AsOverride6"); ok {
			if err = d.Set("as_override6", vv); err != nil {
				return fmt.Errorf("Error reading as_override6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading as_override6: %v", err)
		}
	}

	if err = d.Set("attribute_unchanged", flattenRouterBgpNeighborGroupAttributeUnchanged2edl(o["attribute-unchanged"], d, "attribute_unchanged")); err != nil {
		if vv, ok := fortiAPIPatch(o["attribute-unchanged"], "RouterBgpNeighborGroup-AttributeUnchanged"); ok {
			if err = d.Set("attribute_unchanged", vv); err != nil {
				return fmt.Errorf("Error reading attribute_unchanged: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading attribute_unchanged: %v", err)
		}
	}

	if err = d.Set("attribute_unchanged_vpnv4", flattenRouterBgpNeighborGroupAttributeUnchangedVpnv42edl(o["attribute-unchanged-vpnv4"], d, "attribute_unchanged_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["attribute-unchanged-vpnv4"], "RouterBgpNeighborGroup-AttributeUnchangedVpnv4"); ok {
			if err = d.Set("attribute_unchanged_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading attribute_unchanged_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading attribute_unchanged_vpnv4: %v", err)
		}
	}

	if err = d.Set("attribute_unchanged_vpnv6", flattenRouterBgpNeighborGroupAttributeUnchangedVpnv62edl(o["attribute-unchanged-vpnv6"], d, "attribute_unchanged_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["attribute-unchanged-vpnv6"], "RouterBgpNeighborGroup-AttributeUnchangedVpnv6"); ok {
			if err = d.Set("attribute_unchanged_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading attribute_unchanged_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading attribute_unchanged_vpnv6: %v", err)
		}
	}

	if err = d.Set("attribute_unchanged6", flattenRouterBgpNeighborGroupAttributeUnchanged62edl(o["attribute-unchanged6"], d, "attribute_unchanged6")); err != nil {
		if vv, ok := fortiAPIPatch(o["attribute-unchanged6"], "RouterBgpNeighborGroup-AttributeUnchanged6"); ok {
			if err = d.Set("attribute_unchanged6", vv); err != nil {
				return fmt.Errorf("Error reading attribute_unchanged6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading attribute_unchanged6: %v", err)
		}
	}

	if err = d.Set("auth_options", flattenRouterBgpNeighborGroupAuthOptions2edl(o["auth-options"], d, "auth_options")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-options"], "RouterBgpNeighborGroup-AuthOptions"); ok {
			if err = d.Set("auth_options", vv); err != nil {
				return fmt.Errorf("Error reading auth_options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_options: %v", err)
		}
	}

	if err = d.Set("bfd", flattenRouterBgpNeighborGroupBfd2edl(o["bfd"], d, "bfd")); err != nil {
		if vv, ok := fortiAPIPatch(o["bfd"], "RouterBgpNeighborGroup-Bfd"); ok {
			if err = d.Set("bfd", vv); err != nil {
				return fmt.Errorf("Error reading bfd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	if err = d.Set("capability_default_originate", flattenRouterBgpNeighborGroupCapabilityDefaultOriginate2edl(o["capability-default-originate"], d, "capability_default_originate")); err != nil {
		if vv, ok := fortiAPIPatch(o["capability-default-originate"], "RouterBgpNeighborGroup-CapabilityDefaultOriginate"); ok {
			if err = d.Set("capability_default_originate", vv); err != nil {
				return fmt.Errorf("Error reading capability_default_originate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capability_default_originate: %v", err)
		}
	}

	if err = d.Set("capability_default_originate6", flattenRouterBgpNeighborGroupCapabilityDefaultOriginate62edl(o["capability-default-originate6"], d, "capability_default_originate6")); err != nil {
		if vv, ok := fortiAPIPatch(o["capability-default-originate6"], "RouterBgpNeighborGroup-CapabilityDefaultOriginate6"); ok {
			if err = d.Set("capability_default_originate6", vv); err != nil {
				return fmt.Errorf("Error reading capability_default_originate6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capability_default_originate6: %v", err)
		}
	}

	if err = d.Set("capability_dynamic", flattenRouterBgpNeighborGroupCapabilityDynamic2edl(o["capability-dynamic"], d, "capability_dynamic")); err != nil {
		if vv, ok := fortiAPIPatch(o["capability-dynamic"], "RouterBgpNeighborGroup-CapabilityDynamic"); ok {
			if err = d.Set("capability_dynamic", vv); err != nil {
				return fmt.Errorf("Error reading capability_dynamic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capability_dynamic: %v", err)
		}
	}

	if err = d.Set("capability_graceful_restart", flattenRouterBgpNeighborGroupCapabilityGracefulRestart2edl(o["capability-graceful-restart"], d, "capability_graceful_restart")); err != nil {
		if vv, ok := fortiAPIPatch(o["capability-graceful-restart"], "RouterBgpNeighborGroup-CapabilityGracefulRestart"); ok {
			if err = d.Set("capability_graceful_restart", vv); err != nil {
				return fmt.Errorf("Error reading capability_graceful_restart: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capability_graceful_restart: %v", err)
		}
	}

	if err = d.Set("capability_graceful_restart_evpn", flattenRouterBgpNeighborGroupCapabilityGracefulRestartEvpn2edl(o["capability-graceful-restart-evpn"], d, "capability_graceful_restart_evpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["capability-graceful-restart-evpn"], "RouterBgpNeighborGroup-CapabilityGracefulRestartEvpn"); ok {
			if err = d.Set("capability_graceful_restart_evpn", vv); err != nil {
				return fmt.Errorf("Error reading capability_graceful_restart_evpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capability_graceful_restart_evpn: %v", err)
		}
	}

	if err = d.Set("capability_graceful_restart_vpnv4", flattenRouterBgpNeighborGroupCapabilityGracefulRestartVpnv42edl(o["capability-graceful-restart-vpnv4"], d, "capability_graceful_restart_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["capability-graceful-restart-vpnv4"], "RouterBgpNeighborGroup-CapabilityGracefulRestartVpnv4"); ok {
			if err = d.Set("capability_graceful_restart_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading capability_graceful_restart_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capability_graceful_restart_vpnv4: %v", err)
		}
	}

	if err = d.Set("capability_graceful_restart_vpnv6", flattenRouterBgpNeighborGroupCapabilityGracefulRestartVpnv62edl(o["capability-graceful-restart-vpnv6"], d, "capability_graceful_restart_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["capability-graceful-restart-vpnv6"], "RouterBgpNeighborGroup-CapabilityGracefulRestartVpnv6"); ok {
			if err = d.Set("capability_graceful_restart_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading capability_graceful_restart_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capability_graceful_restart_vpnv6: %v", err)
		}
	}

	if err = d.Set("capability_graceful_restart6", flattenRouterBgpNeighborGroupCapabilityGracefulRestart62edl(o["capability-graceful-restart6"], d, "capability_graceful_restart6")); err != nil {
		if vv, ok := fortiAPIPatch(o["capability-graceful-restart6"], "RouterBgpNeighborGroup-CapabilityGracefulRestart6"); ok {
			if err = d.Set("capability_graceful_restart6", vv); err != nil {
				return fmt.Errorf("Error reading capability_graceful_restart6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capability_graceful_restart6: %v", err)
		}
	}

	if err = d.Set("capability_orf", flattenRouterBgpNeighborGroupCapabilityOrf2edl(o["capability-orf"], d, "capability_orf")); err != nil {
		if vv, ok := fortiAPIPatch(o["capability-orf"], "RouterBgpNeighborGroup-CapabilityOrf"); ok {
			if err = d.Set("capability_orf", vv); err != nil {
				return fmt.Errorf("Error reading capability_orf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capability_orf: %v", err)
		}
	}

	if err = d.Set("capability_orf6", flattenRouterBgpNeighborGroupCapabilityOrf62edl(o["capability-orf6"], d, "capability_orf6")); err != nil {
		if vv, ok := fortiAPIPatch(o["capability-orf6"], "RouterBgpNeighborGroup-CapabilityOrf6"); ok {
			if err = d.Set("capability_orf6", vv); err != nil {
				return fmt.Errorf("Error reading capability_orf6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capability_orf6: %v", err)
		}
	}

	if err = d.Set("capability_route_refresh", flattenRouterBgpNeighborGroupCapabilityRouteRefresh2edl(o["capability-route-refresh"], d, "capability_route_refresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["capability-route-refresh"], "RouterBgpNeighborGroup-CapabilityRouteRefresh"); ok {
			if err = d.Set("capability_route_refresh", vv); err != nil {
				return fmt.Errorf("Error reading capability_route_refresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capability_route_refresh: %v", err)
		}
	}

	if err = d.Set("connect_timer", flattenRouterBgpNeighborGroupConnectTimer2edl(o["connect-timer"], d, "connect_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["connect-timer"], "RouterBgpNeighborGroup-ConnectTimer"); ok {
			if err = d.Set("connect_timer", vv); err != nil {
				return fmt.Errorf("Error reading connect_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading connect_timer: %v", err)
		}
	}

	if err = d.Set("default_originate_routemap", flattenRouterBgpNeighborGroupDefaultOriginateRoutemap2edl(o["default-originate-routemap"], d, "default_originate_routemap")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-originate-routemap"], "RouterBgpNeighborGroup-DefaultOriginateRoutemap"); ok {
			if err = d.Set("default_originate_routemap", vv); err != nil {
				return fmt.Errorf("Error reading default_originate_routemap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_originate_routemap: %v", err)
		}
	}

	if err = d.Set("default_originate_routemap6", flattenRouterBgpNeighborGroupDefaultOriginateRoutemap62edl(o["default-originate-routemap6"], d, "default_originate_routemap6")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-originate-routemap6"], "RouterBgpNeighborGroup-DefaultOriginateRoutemap6"); ok {
			if err = d.Set("default_originate_routemap6", vv); err != nil {
				return fmt.Errorf("Error reading default_originate_routemap6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_originate_routemap6: %v", err)
		}
	}

	if err = d.Set("description", flattenRouterBgpNeighborGroupDescription2edl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "RouterBgpNeighborGroup-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("distribute_list_in", flattenRouterBgpNeighborGroupDistributeListIn2edl(o["distribute-list-in"], d, "distribute_list_in")); err != nil {
		if vv, ok := fortiAPIPatch(o["distribute-list-in"], "RouterBgpNeighborGroup-DistributeListIn"); ok {
			if err = d.Set("distribute_list_in", vv); err != nil {
				return fmt.Errorf("Error reading distribute_list_in: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distribute_list_in: %v", err)
		}
	}

	if err = d.Set("distribute_list_in_vpnv4", flattenRouterBgpNeighborGroupDistributeListInVpnv42edl(o["distribute-list-in-vpnv4"], d, "distribute_list_in_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["distribute-list-in-vpnv4"], "RouterBgpNeighborGroup-DistributeListInVpnv4"); ok {
			if err = d.Set("distribute_list_in_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading distribute_list_in_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distribute_list_in_vpnv4: %v", err)
		}
	}

	if err = d.Set("distribute_list_in_vpnv6", flattenRouterBgpNeighborGroupDistributeListInVpnv62edl(o["distribute-list-in-vpnv6"], d, "distribute_list_in_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["distribute-list-in-vpnv6"], "RouterBgpNeighborGroup-DistributeListInVpnv6"); ok {
			if err = d.Set("distribute_list_in_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading distribute_list_in_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distribute_list_in_vpnv6: %v", err)
		}
	}

	if err = d.Set("distribute_list_in6", flattenRouterBgpNeighborGroupDistributeListIn62edl(o["distribute-list-in6"], d, "distribute_list_in6")); err != nil {
		if vv, ok := fortiAPIPatch(o["distribute-list-in6"], "RouterBgpNeighborGroup-DistributeListIn6"); ok {
			if err = d.Set("distribute_list_in6", vv); err != nil {
				return fmt.Errorf("Error reading distribute_list_in6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distribute_list_in6: %v", err)
		}
	}

	if err = d.Set("distribute_list_out", flattenRouterBgpNeighborGroupDistributeListOut2edl(o["distribute-list-out"], d, "distribute_list_out")); err != nil {
		if vv, ok := fortiAPIPatch(o["distribute-list-out"], "RouterBgpNeighborGroup-DistributeListOut"); ok {
			if err = d.Set("distribute_list_out", vv); err != nil {
				return fmt.Errorf("Error reading distribute_list_out: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distribute_list_out: %v", err)
		}
	}

	if err = d.Set("distribute_list_out_vpnv4", flattenRouterBgpNeighborGroupDistributeListOutVpnv42edl(o["distribute-list-out-vpnv4"], d, "distribute_list_out_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["distribute-list-out-vpnv4"], "RouterBgpNeighborGroup-DistributeListOutVpnv4"); ok {
			if err = d.Set("distribute_list_out_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading distribute_list_out_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distribute_list_out_vpnv4: %v", err)
		}
	}

	if err = d.Set("distribute_list_out_vpnv6", flattenRouterBgpNeighborGroupDistributeListOutVpnv62edl(o["distribute-list-out-vpnv6"], d, "distribute_list_out_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["distribute-list-out-vpnv6"], "RouterBgpNeighborGroup-DistributeListOutVpnv6"); ok {
			if err = d.Set("distribute_list_out_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading distribute_list_out_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distribute_list_out_vpnv6: %v", err)
		}
	}

	if err = d.Set("distribute_list_out6", flattenRouterBgpNeighborGroupDistributeListOut62edl(o["distribute-list-out6"], d, "distribute_list_out6")); err != nil {
		if vv, ok := fortiAPIPatch(o["distribute-list-out6"], "RouterBgpNeighborGroup-DistributeListOut6"); ok {
			if err = d.Set("distribute_list_out6", vv); err != nil {
				return fmt.Errorf("Error reading distribute_list_out6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distribute_list_out6: %v", err)
		}
	}

	if err = d.Set("dont_capability_negotiate", flattenRouterBgpNeighborGroupDontCapabilityNegotiate2edl(o["dont-capability-negotiate"], d, "dont_capability_negotiate")); err != nil {
		if vv, ok := fortiAPIPatch(o["dont-capability-negotiate"], "RouterBgpNeighborGroup-DontCapabilityNegotiate"); ok {
			if err = d.Set("dont_capability_negotiate", vv); err != nil {
				return fmt.Errorf("Error reading dont_capability_negotiate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dont_capability_negotiate: %v", err)
		}
	}

	if err = d.Set("ebgp_enforce_multihop", flattenRouterBgpNeighborGroupEbgpEnforceMultihop2edl(o["ebgp-enforce-multihop"], d, "ebgp_enforce_multihop")); err != nil {
		if vv, ok := fortiAPIPatch(o["ebgp-enforce-multihop"], "RouterBgpNeighborGroup-EbgpEnforceMultihop"); ok {
			if err = d.Set("ebgp_enforce_multihop", vv); err != nil {
				return fmt.Errorf("Error reading ebgp_enforce_multihop: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ebgp_enforce_multihop: %v", err)
		}
	}

	if err = d.Set("ebgp_multihop_ttl", flattenRouterBgpNeighborGroupEbgpMultihopTtl2edl(o["ebgp-multihop-ttl"], d, "ebgp_multihop_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["ebgp-multihop-ttl"], "RouterBgpNeighborGroup-EbgpMultihopTtl"); ok {
			if err = d.Set("ebgp_multihop_ttl", vv); err != nil {
				return fmt.Errorf("Error reading ebgp_multihop_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ebgp_multihop_ttl: %v", err)
		}
	}

	if err = d.Set("filter_list_in", flattenRouterBgpNeighborGroupFilterListIn2edl(o["filter-list-in"], d, "filter_list_in")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-list-in"], "RouterBgpNeighborGroup-FilterListIn"); ok {
			if err = d.Set("filter_list_in", vv); err != nil {
				return fmt.Errorf("Error reading filter_list_in: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_list_in: %v", err)
		}
	}

	if err = d.Set("filter_list_in_vpnv4", flattenRouterBgpNeighborGroupFilterListInVpnv42edl(o["filter-list-in-vpnv4"], d, "filter_list_in_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-list-in-vpnv4"], "RouterBgpNeighborGroup-FilterListInVpnv4"); ok {
			if err = d.Set("filter_list_in_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading filter_list_in_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_list_in_vpnv4: %v", err)
		}
	}

	if err = d.Set("filter_list_in_vpnv6", flattenRouterBgpNeighborGroupFilterListInVpnv62edl(o["filter-list-in-vpnv6"], d, "filter_list_in_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-list-in-vpnv6"], "RouterBgpNeighborGroup-FilterListInVpnv6"); ok {
			if err = d.Set("filter_list_in_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading filter_list_in_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_list_in_vpnv6: %v", err)
		}
	}

	if err = d.Set("filter_list_in6", flattenRouterBgpNeighborGroupFilterListIn62edl(o["filter-list-in6"], d, "filter_list_in6")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-list-in6"], "RouterBgpNeighborGroup-FilterListIn6"); ok {
			if err = d.Set("filter_list_in6", vv); err != nil {
				return fmt.Errorf("Error reading filter_list_in6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_list_in6: %v", err)
		}
	}

	if err = d.Set("filter_list_out", flattenRouterBgpNeighborGroupFilterListOut2edl(o["filter-list-out"], d, "filter_list_out")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-list-out"], "RouterBgpNeighborGroup-FilterListOut"); ok {
			if err = d.Set("filter_list_out", vv); err != nil {
				return fmt.Errorf("Error reading filter_list_out: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_list_out: %v", err)
		}
	}

	if err = d.Set("filter_list_out_vpnv4", flattenRouterBgpNeighborGroupFilterListOutVpnv42edl(o["filter-list-out-vpnv4"], d, "filter_list_out_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-list-out-vpnv4"], "RouterBgpNeighborGroup-FilterListOutVpnv4"); ok {
			if err = d.Set("filter_list_out_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading filter_list_out_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_list_out_vpnv4: %v", err)
		}
	}

	if err = d.Set("filter_list_out_vpnv6", flattenRouterBgpNeighborGroupFilterListOutVpnv62edl(o["filter-list-out-vpnv6"], d, "filter_list_out_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-list-out-vpnv6"], "RouterBgpNeighborGroup-FilterListOutVpnv6"); ok {
			if err = d.Set("filter_list_out_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading filter_list_out_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_list_out_vpnv6: %v", err)
		}
	}

	if err = d.Set("filter_list_out6", flattenRouterBgpNeighborGroupFilterListOut62edl(o["filter-list-out6"], d, "filter_list_out6")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-list-out6"], "RouterBgpNeighborGroup-FilterListOut6"); ok {
			if err = d.Set("filter_list_out6", vv); err != nil {
				return fmt.Errorf("Error reading filter_list_out6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_list_out6: %v", err)
		}
	}

	if err = d.Set("holdtime_timer", flattenRouterBgpNeighborGroupHoldtimeTimer2edl(o["holdtime-timer"], d, "holdtime_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["holdtime-timer"], "RouterBgpNeighborGroup-HoldtimeTimer"); ok {
			if err = d.Set("holdtime_timer", vv); err != nil {
				return fmt.Errorf("Error reading holdtime_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading holdtime_timer: %v", err)
		}
	}

	if err = d.Set("interface", flattenRouterBgpNeighborGroupInterface2edl(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "RouterBgpNeighborGroup-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("keep_alive_timer", flattenRouterBgpNeighborGroupKeepAliveTimer2edl(o["keep-alive-timer"], d, "keep_alive_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["keep-alive-timer"], "RouterBgpNeighborGroup-KeepAliveTimer"); ok {
			if err = d.Set("keep_alive_timer", vv); err != nil {
				return fmt.Errorf("Error reading keep_alive_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keep_alive_timer: %v", err)
		}
	}

	if err = d.Set("link_down_failover", flattenRouterBgpNeighborGroupLinkDownFailover2edl(o["link-down-failover"], d, "link_down_failover")); err != nil {
		if vv, ok := fortiAPIPatch(o["link-down-failover"], "RouterBgpNeighborGroup-LinkDownFailover"); ok {
			if err = d.Set("link_down_failover", vv); err != nil {
				return fmt.Errorf("Error reading link_down_failover: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading link_down_failover: %v", err)
		}
	}

	if err = d.Set("local_as", flattenRouterBgpNeighborGroupLocalAs2edl(o["local-as"], d, "local_as")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-as"], "RouterBgpNeighborGroup-LocalAs"); ok {
			if err = d.Set("local_as", vv); err != nil {
				return fmt.Errorf("Error reading local_as: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_as: %v", err)
		}
	}

	if err = d.Set("local_as_no_prepend", flattenRouterBgpNeighborGroupLocalAsNoPrepend2edl(o["local-as-no-prepend"], d, "local_as_no_prepend")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-as-no-prepend"], "RouterBgpNeighborGroup-LocalAsNoPrepend"); ok {
			if err = d.Set("local_as_no_prepend", vv); err != nil {
				return fmt.Errorf("Error reading local_as_no_prepend: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_as_no_prepend: %v", err)
		}
	}

	if err = d.Set("local_as_replace_as", flattenRouterBgpNeighborGroupLocalAsReplaceAs2edl(o["local-as-replace-as"], d, "local_as_replace_as")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-as-replace-as"], "RouterBgpNeighborGroup-LocalAsReplaceAs"); ok {
			if err = d.Set("local_as_replace_as", vv); err != nil {
				return fmt.Errorf("Error reading local_as_replace_as: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_as_replace_as: %v", err)
		}
	}

	if err = d.Set("maximum_prefix", flattenRouterBgpNeighborGroupMaximumPrefix2edl(o["maximum-prefix"], d, "maximum_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-prefix"], "RouterBgpNeighborGroup-MaximumPrefix"); ok {
			if err = d.Set("maximum_prefix", vv); err != nil {
				return fmt.Errorf("Error reading maximum_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_prefix: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_evpn", flattenRouterBgpNeighborGroupMaximumPrefixEvpn2edl(o["maximum-prefix-evpn"], d, "maximum_prefix_evpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-prefix-evpn"], "RouterBgpNeighborGroup-MaximumPrefixEvpn"); ok {
			if err = d.Set("maximum_prefix_evpn", vv); err != nil {
				return fmt.Errorf("Error reading maximum_prefix_evpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_prefix_evpn: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_threshold", flattenRouterBgpNeighborGroupMaximumPrefixThreshold2edl(o["maximum-prefix-threshold"], d, "maximum_prefix_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-prefix-threshold"], "RouterBgpNeighborGroup-MaximumPrefixThreshold"); ok {
			if err = d.Set("maximum_prefix_threshold", vv); err != nil {
				return fmt.Errorf("Error reading maximum_prefix_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_prefix_threshold: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_threshold_evpn", flattenRouterBgpNeighborGroupMaximumPrefixThresholdEvpn2edl(o["maximum-prefix-threshold-evpn"], d, "maximum_prefix_threshold_evpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-prefix-threshold-evpn"], "RouterBgpNeighborGroup-MaximumPrefixThresholdEvpn"); ok {
			if err = d.Set("maximum_prefix_threshold_evpn", vv); err != nil {
				return fmt.Errorf("Error reading maximum_prefix_threshold_evpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_prefix_threshold_evpn: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_threshold_vpnv4", flattenRouterBgpNeighborGroupMaximumPrefixThresholdVpnv42edl(o["maximum-prefix-threshold-vpnv4"], d, "maximum_prefix_threshold_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-prefix-threshold-vpnv4"], "RouterBgpNeighborGroup-MaximumPrefixThresholdVpnv4"); ok {
			if err = d.Set("maximum_prefix_threshold_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading maximum_prefix_threshold_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_prefix_threshold_vpnv4: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_threshold_vpnv6", flattenRouterBgpNeighborGroupMaximumPrefixThresholdVpnv62edl(o["maximum-prefix-threshold-vpnv6"], d, "maximum_prefix_threshold_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-prefix-threshold-vpnv6"], "RouterBgpNeighborGroup-MaximumPrefixThresholdVpnv6"); ok {
			if err = d.Set("maximum_prefix_threshold_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading maximum_prefix_threshold_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_prefix_threshold_vpnv6: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_threshold6", flattenRouterBgpNeighborGroupMaximumPrefixThreshold62edl(o["maximum-prefix-threshold6"], d, "maximum_prefix_threshold6")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-prefix-threshold6"], "RouterBgpNeighborGroup-MaximumPrefixThreshold6"); ok {
			if err = d.Set("maximum_prefix_threshold6", vv); err != nil {
				return fmt.Errorf("Error reading maximum_prefix_threshold6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_prefix_threshold6: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_vpnv4", flattenRouterBgpNeighborGroupMaximumPrefixVpnv42edl(o["maximum-prefix-vpnv4"], d, "maximum_prefix_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-prefix-vpnv4"], "RouterBgpNeighborGroup-MaximumPrefixVpnv4"); ok {
			if err = d.Set("maximum_prefix_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading maximum_prefix_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_prefix_vpnv4: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_vpnv6", flattenRouterBgpNeighborGroupMaximumPrefixVpnv62edl(o["maximum-prefix-vpnv6"], d, "maximum_prefix_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-prefix-vpnv6"], "RouterBgpNeighborGroup-MaximumPrefixVpnv6"); ok {
			if err = d.Set("maximum_prefix_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading maximum_prefix_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_prefix_vpnv6: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_warning_only", flattenRouterBgpNeighborGroupMaximumPrefixWarningOnly2edl(o["maximum-prefix-warning-only"], d, "maximum_prefix_warning_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-prefix-warning-only"], "RouterBgpNeighborGroup-MaximumPrefixWarningOnly"); ok {
			if err = d.Set("maximum_prefix_warning_only", vv); err != nil {
				return fmt.Errorf("Error reading maximum_prefix_warning_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_prefix_warning_only: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_warning_only_evpn", flattenRouterBgpNeighborGroupMaximumPrefixWarningOnlyEvpn2edl(o["maximum-prefix-warning-only-evpn"], d, "maximum_prefix_warning_only_evpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-prefix-warning-only-evpn"], "RouterBgpNeighborGroup-MaximumPrefixWarningOnlyEvpn"); ok {
			if err = d.Set("maximum_prefix_warning_only_evpn", vv); err != nil {
				return fmt.Errorf("Error reading maximum_prefix_warning_only_evpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_prefix_warning_only_evpn: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_warning_only_vpnv4", flattenRouterBgpNeighborGroupMaximumPrefixWarningOnlyVpnv42edl(o["maximum-prefix-warning-only-vpnv4"], d, "maximum_prefix_warning_only_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-prefix-warning-only-vpnv4"], "RouterBgpNeighborGroup-MaximumPrefixWarningOnlyVpnv4"); ok {
			if err = d.Set("maximum_prefix_warning_only_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading maximum_prefix_warning_only_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_prefix_warning_only_vpnv4: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_warning_only_vpnv6", flattenRouterBgpNeighborGroupMaximumPrefixWarningOnlyVpnv62edl(o["maximum-prefix-warning-only-vpnv6"], d, "maximum_prefix_warning_only_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-prefix-warning-only-vpnv6"], "RouterBgpNeighborGroup-MaximumPrefixWarningOnlyVpnv6"); ok {
			if err = d.Set("maximum_prefix_warning_only_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading maximum_prefix_warning_only_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_prefix_warning_only_vpnv6: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_warning_only6", flattenRouterBgpNeighborGroupMaximumPrefixWarningOnly62edl(o["maximum-prefix-warning-only6"], d, "maximum_prefix_warning_only6")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-prefix-warning-only6"], "RouterBgpNeighborGroup-MaximumPrefixWarningOnly6"); ok {
			if err = d.Set("maximum_prefix_warning_only6", vv); err != nil {
				return fmt.Errorf("Error reading maximum_prefix_warning_only6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_prefix_warning_only6: %v", err)
		}
	}

	if err = d.Set("maximum_prefix6", flattenRouterBgpNeighborGroupMaximumPrefix62edl(o["maximum-prefix6"], d, "maximum_prefix6")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-prefix6"], "RouterBgpNeighborGroup-MaximumPrefix6"); ok {
			if err = d.Set("maximum_prefix6", vv); err != nil {
				return fmt.Errorf("Error reading maximum_prefix6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_prefix6: %v", err)
		}
	}

	if err = d.Set("name", flattenRouterBgpNeighborGroupName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "RouterBgpNeighborGroup-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("next_hop_self", flattenRouterBgpNeighborGroupNextHopSelf2edl(o["next-hop-self"], d, "next_hop_self")); err != nil {
		if vv, ok := fortiAPIPatch(o["next-hop-self"], "RouterBgpNeighborGroup-NextHopSelf"); ok {
			if err = d.Set("next_hop_self", vv); err != nil {
				return fmt.Errorf("Error reading next_hop_self: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading next_hop_self: %v", err)
		}
	}

	if err = d.Set("next_hop_self_rr", flattenRouterBgpNeighborGroupNextHopSelfRr2edl(o["next-hop-self-rr"], d, "next_hop_self_rr")); err != nil {
		if vv, ok := fortiAPIPatch(o["next-hop-self-rr"], "RouterBgpNeighborGroup-NextHopSelfRr"); ok {
			if err = d.Set("next_hop_self_rr", vv); err != nil {
				return fmt.Errorf("Error reading next_hop_self_rr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading next_hop_self_rr: %v", err)
		}
	}

	if err = d.Set("next_hop_self_rr6", flattenRouterBgpNeighborGroupNextHopSelfRr62edl(o["next-hop-self-rr6"], d, "next_hop_self_rr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["next-hop-self-rr6"], "RouterBgpNeighborGroup-NextHopSelfRr6"); ok {
			if err = d.Set("next_hop_self_rr6", vv); err != nil {
				return fmt.Errorf("Error reading next_hop_self_rr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading next_hop_self_rr6: %v", err)
		}
	}

	if err = d.Set("next_hop_self_vpnv4", flattenRouterBgpNeighborGroupNextHopSelfVpnv42edl(o["next-hop-self-vpnv4"], d, "next_hop_self_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["next-hop-self-vpnv4"], "RouterBgpNeighborGroup-NextHopSelfVpnv4"); ok {
			if err = d.Set("next_hop_self_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading next_hop_self_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading next_hop_self_vpnv4: %v", err)
		}
	}

	if err = d.Set("next_hop_self_vpnv6", flattenRouterBgpNeighborGroupNextHopSelfVpnv62edl(o["next-hop-self-vpnv6"], d, "next_hop_self_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["next-hop-self-vpnv6"], "RouterBgpNeighborGroup-NextHopSelfVpnv6"); ok {
			if err = d.Set("next_hop_self_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading next_hop_self_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading next_hop_self_vpnv6: %v", err)
		}
	}

	if err = d.Set("next_hop_self6", flattenRouterBgpNeighborGroupNextHopSelf62edl(o["next-hop-self6"], d, "next_hop_self6")); err != nil {
		if vv, ok := fortiAPIPatch(o["next-hop-self6"], "RouterBgpNeighborGroup-NextHopSelf6"); ok {
			if err = d.Set("next_hop_self6", vv); err != nil {
				return fmt.Errorf("Error reading next_hop_self6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading next_hop_self6: %v", err)
		}
	}

	if err = d.Set("override_capability", flattenRouterBgpNeighborGroupOverrideCapability2edl(o["override-capability"], d, "override_capability")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-capability"], "RouterBgpNeighborGroup-OverrideCapability"); ok {
			if err = d.Set("override_capability", vv); err != nil {
				return fmt.Errorf("Error reading override_capability: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_capability: %v", err)
		}
	}

	if err = d.Set("passive", flattenRouterBgpNeighborGroupPassive2edl(o["passive"], d, "passive")); err != nil {
		if vv, ok := fortiAPIPatch(o["passive"], "RouterBgpNeighborGroup-Passive"); ok {
			if err = d.Set("passive", vv); err != nil {
				return fmt.Errorf("Error reading passive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading passive: %v", err)
		}
	}

	if err = d.Set("prefix_list_in", flattenRouterBgpNeighborGroupPrefixListIn2edl(o["prefix-list-in"], d, "prefix_list_in")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix-list-in"], "RouterBgpNeighborGroup-PrefixListIn"); ok {
			if err = d.Set("prefix_list_in", vv); err != nil {
				return fmt.Errorf("Error reading prefix_list_in: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix_list_in: %v", err)
		}
	}

	if err = d.Set("prefix_list_in_vpnv4", flattenRouterBgpNeighborGroupPrefixListInVpnv42edl(o["prefix-list-in-vpnv4"], d, "prefix_list_in_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix-list-in-vpnv4"], "RouterBgpNeighborGroup-PrefixListInVpnv4"); ok {
			if err = d.Set("prefix_list_in_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading prefix_list_in_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix_list_in_vpnv4: %v", err)
		}
	}

	if err = d.Set("prefix_list_in_vpnv6", flattenRouterBgpNeighborGroupPrefixListInVpnv62edl(o["prefix-list-in-vpnv6"], d, "prefix_list_in_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix-list-in-vpnv6"], "RouterBgpNeighborGroup-PrefixListInVpnv6"); ok {
			if err = d.Set("prefix_list_in_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading prefix_list_in_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix_list_in_vpnv6: %v", err)
		}
	}

	if err = d.Set("prefix_list_in6", flattenRouterBgpNeighborGroupPrefixListIn62edl(o["prefix-list-in6"], d, "prefix_list_in6")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix-list-in6"], "RouterBgpNeighborGroup-PrefixListIn6"); ok {
			if err = d.Set("prefix_list_in6", vv); err != nil {
				return fmt.Errorf("Error reading prefix_list_in6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix_list_in6: %v", err)
		}
	}

	if err = d.Set("prefix_list_out", flattenRouterBgpNeighborGroupPrefixListOut2edl(o["prefix-list-out"], d, "prefix_list_out")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix-list-out"], "RouterBgpNeighborGroup-PrefixListOut"); ok {
			if err = d.Set("prefix_list_out", vv); err != nil {
				return fmt.Errorf("Error reading prefix_list_out: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix_list_out: %v", err)
		}
	}

	if err = d.Set("prefix_list_out_vpnv4", flattenRouterBgpNeighborGroupPrefixListOutVpnv42edl(o["prefix-list-out-vpnv4"], d, "prefix_list_out_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix-list-out-vpnv4"], "RouterBgpNeighborGroup-PrefixListOutVpnv4"); ok {
			if err = d.Set("prefix_list_out_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading prefix_list_out_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix_list_out_vpnv4: %v", err)
		}
	}

	if err = d.Set("prefix_list_out_vpnv6", flattenRouterBgpNeighborGroupPrefixListOutVpnv62edl(o["prefix-list-out-vpnv6"], d, "prefix_list_out_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix-list-out-vpnv6"], "RouterBgpNeighborGroup-PrefixListOutVpnv6"); ok {
			if err = d.Set("prefix_list_out_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading prefix_list_out_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix_list_out_vpnv6: %v", err)
		}
	}

	if err = d.Set("prefix_list_out6", flattenRouterBgpNeighborGroupPrefixListOut62edl(o["prefix-list-out6"], d, "prefix_list_out6")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix-list-out6"], "RouterBgpNeighborGroup-PrefixListOut6"); ok {
			if err = d.Set("prefix_list_out6", vv); err != nil {
				return fmt.Errorf("Error reading prefix_list_out6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix_list_out6: %v", err)
		}
	}

	if err = d.Set("remote_as", flattenRouterBgpNeighborGroupRemoteAs2edl(o["remote-as"], d, "remote_as")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-as"], "RouterBgpNeighborGroup-RemoteAs"); ok {
			if err = d.Set("remote_as", vv); err != nil {
				return fmt.Errorf("Error reading remote_as: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_as: %v", err)
		}
	}

	if err = d.Set("remote_as_filter", flattenRouterBgpNeighborGroupRemoteAsFilter2edl(o["remote-as-filter"], d, "remote_as_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-as-filter"], "RouterBgpNeighborGroup-RemoteAsFilter"); ok {
			if err = d.Set("remote_as_filter", vv); err != nil {
				return fmt.Errorf("Error reading remote_as_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_as_filter: %v", err)
		}
	}

	if err = d.Set("remove_private_as", flattenRouterBgpNeighborGroupRemovePrivateAs2edl(o["remove-private-as"], d, "remove_private_as")); err != nil {
		if vv, ok := fortiAPIPatch(o["remove-private-as"], "RouterBgpNeighborGroup-RemovePrivateAs"); ok {
			if err = d.Set("remove_private_as", vv); err != nil {
				return fmt.Errorf("Error reading remove_private_as: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remove_private_as: %v", err)
		}
	}

	if err = d.Set("remove_private_as_evpn", flattenRouterBgpNeighborGroupRemovePrivateAsEvpn2edl(o["remove-private-as-evpn"], d, "remove_private_as_evpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["remove-private-as-evpn"], "RouterBgpNeighborGroup-RemovePrivateAsEvpn"); ok {
			if err = d.Set("remove_private_as_evpn", vv); err != nil {
				return fmt.Errorf("Error reading remove_private_as_evpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remove_private_as_evpn: %v", err)
		}
	}

	if err = d.Set("remove_private_as_vpnv4", flattenRouterBgpNeighborGroupRemovePrivateAsVpnv42edl(o["remove-private-as-vpnv4"], d, "remove_private_as_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["remove-private-as-vpnv4"], "RouterBgpNeighborGroup-RemovePrivateAsVpnv4"); ok {
			if err = d.Set("remove_private_as_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading remove_private_as_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remove_private_as_vpnv4: %v", err)
		}
	}

	if err = d.Set("remove_private_as_vpnv6", flattenRouterBgpNeighborGroupRemovePrivateAsVpnv62edl(o["remove-private-as-vpnv6"], d, "remove_private_as_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["remove-private-as-vpnv6"], "RouterBgpNeighborGroup-RemovePrivateAsVpnv6"); ok {
			if err = d.Set("remove_private_as_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading remove_private_as_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remove_private_as_vpnv6: %v", err)
		}
	}

	if err = d.Set("remove_private_as6", flattenRouterBgpNeighborGroupRemovePrivateAs62edl(o["remove-private-as6"], d, "remove_private_as6")); err != nil {
		if vv, ok := fortiAPIPatch(o["remove-private-as6"], "RouterBgpNeighborGroup-RemovePrivateAs6"); ok {
			if err = d.Set("remove_private_as6", vv); err != nil {
				return fmt.Errorf("Error reading remove_private_as6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remove_private_as6: %v", err)
		}
	}

	if err = d.Set("restart_time", flattenRouterBgpNeighborGroupRestartTime2edl(o["restart-time"], d, "restart_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["restart-time"], "RouterBgpNeighborGroup-RestartTime"); ok {
			if err = d.Set("restart_time", vv); err != nil {
				return fmt.Errorf("Error reading restart_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restart_time: %v", err)
		}
	}

	if err = d.Set("retain_stale_time", flattenRouterBgpNeighborGroupRetainStaleTime2edl(o["retain-stale-time"], d, "retain_stale_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["retain-stale-time"], "RouterBgpNeighborGroup-RetainStaleTime"); ok {
			if err = d.Set("retain_stale_time", vv); err != nil {
				return fmt.Errorf("Error reading retain_stale_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading retain_stale_time: %v", err)
		}
	}

	if err = d.Set("route_map_in", flattenRouterBgpNeighborGroupRouteMapIn2edl(o["route-map-in"], d, "route_map_in")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-map-in"], "RouterBgpNeighborGroup-RouteMapIn"); ok {
			if err = d.Set("route_map_in", vv); err != nil {
				return fmt.Errorf("Error reading route_map_in: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_map_in: %v", err)
		}
	}

	if err = d.Set("route_map_in_evpn", flattenRouterBgpNeighborGroupRouteMapInEvpn2edl(o["route-map-in-evpn"], d, "route_map_in_evpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-map-in-evpn"], "RouterBgpNeighborGroup-RouteMapInEvpn"); ok {
			if err = d.Set("route_map_in_evpn", vv); err != nil {
				return fmt.Errorf("Error reading route_map_in_evpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_map_in_evpn: %v", err)
		}
	}

	if err = d.Set("route_map_in_vpnv4", flattenRouterBgpNeighborGroupRouteMapInVpnv42edl(o["route-map-in-vpnv4"], d, "route_map_in_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-map-in-vpnv4"], "RouterBgpNeighborGroup-RouteMapInVpnv4"); ok {
			if err = d.Set("route_map_in_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading route_map_in_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_map_in_vpnv4: %v", err)
		}
	}

	if err = d.Set("route_map_in_vpnv6", flattenRouterBgpNeighborGroupRouteMapInVpnv62edl(o["route-map-in-vpnv6"], d, "route_map_in_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-map-in-vpnv6"], "RouterBgpNeighborGroup-RouteMapInVpnv6"); ok {
			if err = d.Set("route_map_in_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading route_map_in_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_map_in_vpnv6: %v", err)
		}
	}

	if err = d.Set("route_map_in6", flattenRouterBgpNeighborGroupRouteMapIn62edl(o["route-map-in6"], d, "route_map_in6")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-map-in6"], "RouterBgpNeighborGroup-RouteMapIn6"); ok {
			if err = d.Set("route_map_in6", vv); err != nil {
				return fmt.Errorf("Error reading route_map_in6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_map_in6: %v", err)
		}
	}

	if err = d.Set("route_map_out", flattenRouterBgpNeighborGroupRouteMapOut2edl(o["route-map-out"], d, "route_map_out")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-map-out"], "RouterBgpNeighborGroup-RouteMapOut"); ok {
			if err = d.Set("route_map_out", vv); err != nil {
				return fmt.Errorf("Error reading route_map_out: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_map_out: %v", err)
		}
	}

	if err = d.Set("route_map_out_evpn", flattenRouterBgpNeighborGroupRouteMapOutEvpn2edl(o["route-map-out-evpn"], d, "route_map_out_evpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-map-out-evpn"], "RouterBgpNeighborGroup-RouteMapOutEvpn"); ok {
			if err = d.Set("route_map_out_evpn", vv); err != nil {
				return fmt.Errorf("Error reading route_map_out_evpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_map_out_evpn: %v", err)
		}
	}

	if err = d.Set("route_map_out_preferable", flattenRouterBgpNeighborGroupRouteMapOutPreferable2edl(o["route-map-out-preferable"], d, "route_map_out_preferable")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-map-out-preferable"], "RouterBgpNeighborGroup-RouteMapOutPreferable"); ok {
			if err = d.Set("route_map_out_preferable", vv); err != nil {
				return fmt.Errorf("Error reading route_map_out_preferable: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_map_out_preferable: %v", err)
		}
	}

	if err = d.Set("route_map_out_vpnv4", flattenRouterBgpNeighborGroupRouteMapOutVpnv42edl(o["route-map-out-vpnv4"], d, "route_map_out_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-map-out-vpnv4"], "RouterBgpNeighborGroup-RouteMapOutVpnv4"); ok {
			if err = d.Set("route_map_out_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading route_map_out_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_map_out_vpnv4: %v", err)
		}
	}

	if err = d.Set("route_map_out_vpnv4_preferable", flattenRouterBgpNeighborGroupRouteMapOutVpnv4Preferable2edl(o["route-map-out-vpnv4-preferable"], d, "route_map_out_vpnv4_preferable")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-map-out-vpnv4-preferable"], "RouterBgpNeighborGroup-RouteMapOutVpnv4Preferable"); ok {
			if err = d.Set("route_map_out_vpnv4_preferable", vv); err != nil {
				return fmt.Errorf("Error reading route_map_out_vpnv4_preferable: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_map_out_vpnv4_preferable: %v", err)
		}
	}

	if err = d.Set("route_map_out_vpnv6", flattenRouterBgpNeighborGroupRouteMapOutVpnv62edl(o["route-map-out-vpnv6"], d, "route_map_out_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-map-out-vpnv6"], "RouterBgpNeighborGroup-RouteMapOutVpnv6"); ok {
			if err = d.Set("route_map_out_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading route_map_out_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_map_out_vpnv6: %v", err)
		}
	}

	if err = d.Set("route_map_out_vpnv6_preferable", flattenRouterBgpNeighborGroupRouteMapOutVpnv6Preferable2edl(o["route-map-out-vpnv6-preferable"], d, "route_map_out_vpnv6_preferable")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-map-out-vpnv6-preferable"], "RouterBgpNeighborGroup-RouteMapOutVpnv6Preferable"); ok {
			if err = d.Set("route_map_out_vpnv6_preferable", vv); err != nil {
				return fmt.Errorf("Error reading route_map_out_vpnv6_preferable: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_map_out_vpnv6_preferable: %v", err)
		}
	}

	if err = d.Set("route_map_out6", flattenRouterBgpNeighborGroupRouteMapOut62edl(o["route-map-out6"], d, "route_map_out6")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-map-out6"], "RouterBgpNeighborGroup-RouteMapOut6"); ok {
			if err = d.Set("route_map_out6", vv); err != nil {
				return fmt.Errorf("Error reading route_map_out6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_map_out6: %v", err)
		}
	}

	if err = d.Set("route_map_out6_preferable", flattenRouterBgpNeighborGroupRouteMapOut6Preferable2edl(o["route-map-out6-preferable"], d, "route_map_out6_preferable")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-map-out6-preferable"], "RouterBgpNeighborGroup-RouteMapOut6Preferable"); ok {
			if err = d.Set("route_map_out6_preferable", vv); err != nil {
				return fmt.Errorf("Error reading route_map_out6_preferable: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_map_out6_preferable: %v", err)
		}
	}

	if err = d.Set("route_reflector_client", flattenRouterBgpNeighborGroupRouteReflectorClient2edl(o["route-reflector-client"], d, "route_reflector_client")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-reflector-client"], "RouterBgpNeighborGroup-RouteReflectorClient"); ok {
			if err = d.Set("route_reflector_client", vv); err != nil {
				return fmt.Errorf("Error reading route_reflector_client: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_reflector_client: %v", err)
		}
	}

	if err = d.Set("route_reflector_client_evpn", flattenRouterBgpNeighborGroupRouteReflectorClientEvpn2edl(o["route-reflector-client-evpn"], d, "route_reflector_client_evpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-reflector-client-evpn"], "RouterBgpNeighborGroup-RouteReflectorClientEvpn"); ok {
			if err = d.Set("route_reflector_client_evpn", vv); err != nil {
				return fmt.Errorf("Error reading route_reflector_client_evpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_reflector_client_evpn: %v", err)
		}
	}

	if err = d.Set("route_reflector_client_vpnv4", flattenRouterBgpNeighborGroupRouteReflectorClientVpnv42edl(o["route-reflector-client-vpnv4"], d, "route_reflector_client_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-reflector-client-vpnv4"], "RouterBgpNeighborGroup-RouteReflectorClientVpnv4"); ok {
			if err = d.Set("route_reflector_client_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading route_reflector_client_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_reflector_client_vpnv4: %v", err)
		}
	}

	if err = d.Set("route_reflector_client_vpnv6", flattenRouterBgpNeighborGroupRouteReflectorClientVpnv62edl(o["route-reflector-client-vpnv6"], d, "route_reflector_client_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-reflector-client-vpnv6"], "RouterBgpNeighborGroup-RouteReflectorClientVpnv6"); ok {
			if err = d.Set("route_reflector_client_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading route_reflector_client_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_reflector_client_vpnv6: %v", err)
		}
	}

	if err = d.Set("route_reflector_client6", flattenRouterBgpNeighborGroupRouteReflectorClient62edl(o["route-reflector-client6"], d, "route_reflector_client6")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-reflector-client6"], "RouterBgpNeighborGroup-RouteReflectorClient6"); ok {
			if err = d.Set("route_reflector_client6", vv); err != nil {
				return fmt.Errorf("Error reading route_reflector_client6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_reflector_client6: %v", err)
		}
	}

	if err = d.Set("route_server_client", flattenRouterBgpNeighborGroupRouteServerClient2edl(o["route-server-client"], d, "route_server_client")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-server-client"], "RouterBgpNeighborGroup-RouteServerClient"); ok {
			if err = d.Set("route_server_client", vv); err != nil {
				return fmt.Errorf("Error reading route_server_client: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_server_client: %v", err)
		}
	}

	if err = d.Set("route_server_client_evpn", flattenRouterBgpNeighborGroupRouteServerClientEvpn2edl(o["route-server-client-evpn"], d, "route_server_client_evpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-server-client-evpn"], "RouterBgpNeighborGroup-RouteServerClientEvpn"); ok {
			if err = d.Set("route_server_client_evpn", vv); err != nil {
				return fmt.Errorf("Error reading route_server_client_evpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_server_client_evpn: %v", err)
		}
	}

	if err = d.Set("route_server_client_vpnv4", flattenRouterBgpNeighborGroupRouteServerClientVpnv42edl(o["route-server-client-vpnv4"], d, "route_server_client_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-server-client-vpnv4"], "RouterBgpNeighborGroup-RouteServerClientVpnv4"); ok {
			if err = d.Set("route_server_client_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading route_server_client_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_server_client_vpnv4: %v", err)
		}
	}

	if err = d.Set("route_server_client_vpnv6", flattenRouterBgpNeighborGroupRouteServerClientVpnv62edl(o["route-server-client-vpnv6"], d, "route_server_client_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-server-client-vpnv6"], "RouterBgpNeighborGroup-RouteServerClientVpnv6"); ok {
			if err = d.Set("route_server_client_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading route_server_client_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_server_client_vpnv6: %v", err)
		}
	}

	if err = d.Set("route_server_client6", flattenRouterBgpNeighborGroupRouteServerClient62edl(o["route-server-client6"], d, "route_server_client6")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-server-client6"], "RouterBgpNeighborGroup-RouteServerClient6"); ok {
			if err = d.Set("route_server_client6", vv); err != nil {
				return fmt.Errorf("Error reading route_server_client6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_server_client6: %v", err)
		}
	}

	if err = d.Set("send_community", flattenRouterBgpNeighborGroupSendCommunity2edl(o["send-community"], d, "send_community")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-community"], "RouterBgpNeighborGroup-SendCommunity"); ok {
			if err = d.Set("send_community", vv); err != nil {
				return fmt.Errorf("Error reading send_community: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_community: %v", err)
		}
	}

	if err = d.Set("send_community_evpn", flattenRouterBgpNeighborGroupSendCommunityEvpn2edl(o["send-community-evpn"], d, "send_community_evpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-community-evpn"], "RouterBgpNeighborGroup-SendCommunityEvpn"); ok {
			if err = d.Set("send_community_evpn", vv); err != nil {
				return fmt.Errorf("Error reading send_community_evpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_community_evpn: %v", err)
		}
	}

	if err = d.Set("send_community_vpnv4", flattenRouterBgpNeighborGroupSendCommunityVpnv42edl(o["send-community-vpnv4"], d, "send_community_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-community-vpnv4"], "RouterBgpNeighborGroup-SendCommunityVpnv4"); ok {
			if err = d.Set("send_community_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading send_community_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_community_vpnv4: %v", err)
		}
	}

	if err = d.Set("send_community_vpnv6", flattenRouterBgpNeighborGroupSendCommunityVpnv62edl(o["send-community-vpnv6"], d, "send_community_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-community-vpnv6"], "RouterBgpNeighborGroup-SendCommunityVpnv6"); ok {
			if err = d.Set("send_community_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading send_community_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_community_vpnv6: %v", err)
		}
	}

	if err = d.Set("send_community6", flattenRouterBgpNeighborGroupSendCommunity62edl(o["send-community6"], d, "send_community6")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-community6"], "RouterBgpNeighborGroup-SendCommunity6"); ok {
			if err = d.Set("send_community6", vv); err != nil {
				return fmt.Errorf("Error reading send_community6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_community6: %v", err)
		}
	}

	if err = d.Set("shutdown", flattenRouterBgpNeighborGroupShutdown2edl(o["shutdown"], d, "shutdown")); err != nil {
		if vv, ok := fortiAPIPatch(o["shutdown"], "RouterBgpNeighborGroup-Shutdown"); ok {
			if err = d.Set("shutdown", vv); err != nil {
				return fmt.Errorf("Error reading shutdown: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading shutdown: %v", err)
		}
	}

	if err = d.Set("soft_reconfiguration", flattenRouterBgpNeighborGroupSoftReconfiguration2edl(o["soft-reconfiguration"], d, "soft_reconfiguration")); err != nil {
		if vv, ok := fortiAPIPatch(o["soft-reconfiguration"], "RouterBgpNeighborGroup-SoftReconfiguration"); ok {
			if err = d.Set("soft_reconfiguration", vv); err != nil {
				return fmt.Errorf("Error reading soft_reconfiguration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading soft_reconfiguration: %v", err)
		}
	}

	if err = d.Set("soft_reconfiguration_evpn", flattenRouterBgpNeighborGroupSoftReconfigurationEvpn2edl(o["soft-reconfiguration-evpn"], d, "soft_reconfiguration_evpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["soft-reconfiguration-evpn"], "RouterBgpNeighborGroup-SoftReconfigurationEvpn"); ok {
			if err = d.Set("soft_reconfiguration_evpn", vv); err != nil {
				return fmt.Errorf("Error reading soft_reconfiguration_evpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading soft_reconfiguration_evpn: %v", err)
		}
	}

	if err = d.Set("soft_reconfiguration_vpnv4", flattenRouterBgpNeighborGroupSoftReconfigurationVpnv42edl(o["soft-reconfiguration-vpnv4"], d, "soft_reconfiguration_vpnv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["soft-reconfiguration-vpnv4"], "RouterBgpNeighborGroup-SoftReconfigurationVpnv4"); ok {
			if err = d.Set("soft_reconfiguration_vpnv4", vv); err != nil {
				return fmt.Errorf("Error reading soft_reconfiguration_vpnv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading soft_reconfiguration_vpnv4: %v", err)
		}
	}

	if err = d.Set("soft_reconfiguration_vpnv6", flattenRouterBgpNeighborGroupSoftReconfigurationVpnv62edl(o["soft-reconfiguration-vpnv6"], d, "soft_reconfiguration_vpnv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["soft-reconfiguration-vpnv6"], "RouterBgpNeighborGroup-SoftReconfigurationVpnv6"); ok {
			if err = d.Set("soft_reconfiguration_vpnv6", vv); err != nil {
				return fmt.Errorf("Error reading soft_reconfiguration_vpnv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading soft_reconfiguration_vpnv6: %v", err)
		}
	}

	if err = d.Set("soft_reconfiguration6", flattenRouterBgpNeighborGroupSoftReconfiguration62edl(o["soft-reconfiguration6"], d, "soft_reconfiguration6")); err != nil {
		if vv, ok := fortiAPIPatch(o["soft-reconfiguration6"], "RouterBgpNeighborGroup-SoftReconfiguration6"); ok {
			if err = d.Set("soft_reconfiguration6", vv); err != nil {
				return fmt.Errorf("Error reading soft_reconfiguration6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading soft_reconfiguration6: %v", err)
		}
	}

	if err = d.Set("stale_route", flattenRouterBgpNeighborGroupStaleRoute2edl(o["stale-route"], d, "stale_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["stale-route"], "RouterBgpNeighborGroup-StaleRoute"); ok {
			if err = d.Set("stale_route", vv); err != nil {
				return fmt.Errorf("Error reading stale_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stale_route: %v", err)
		}
	}

	if err = d.Set("strict_capability_match", flattenRouterBgpNeighborGroupStrictCapabilityMatch2edl(o["strict-capability-match"], d, "strict_capability_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["strict-capability-match"], "RouterBgpNeighborGroup-StrictCapabilityMatch"); ok {
			if err = d.Set("strict_capability_match", vv); err != nil {
				return fmt.Errorf("Error reading strict_capability_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading strict_capability_match: %v", err)
		}
	}

	if err = d.Set("unsuppress_map", flattenRouterBgpNeighborGroupUnsuppressMap2edl(o["unsuppress-map"], d, "unsuppress_map")); err != nil {
		if vv, ok := fortiAPIPatch(o["unsuppress-map"], "RouterBgpNeighborGroup-UnsuppressMap"); ok {
			if err = d.Set("unsuppress_map", vv); err != nil {
				return fmt.Errorf("Error reading unsuppress_map: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unsuppress_map: %v", err)
		}
	}

	if err = d.Set("unsuppress_map6", flattenRouterBgpNeighborGroupUnsuppressMap62edl(o["unsuppress-map6"], d, "unsuppress_map6")); err != nil {
		if vv, ok := fortiAPIPatch(o["unsuppress-map6"], "RouterBgpNeighborGroup-UnsuppressMap6"); ok {
			if err = d.Set("unsuppress_map6", vv); err != nil {
				return fmt.Errorf("Error reading unsuppress_map6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unsuppress_map6: %v", err)
		}
	}

	if err = d.Set("update_source", flattenRouterBgpNeighborGroupUpdateSource2edl(o["update-source"], d, "update_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-source"], "RouterBgpNeighborGroup-UpdateSource"); ok {
			if err = d.Set("update_source", vv); err != nil {
				return fmt.Errorf("Error reading update_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_source: %v", err)
		}
	}

	if err = d.Set("weight", flattenRouterBgpNeighborGroupWeight2edl(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "RouterBgpNeighborGroup-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	return nil
}

func flattenRouterBgpNeighborGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterBgpNeighborGroupActivate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupActivateEvpn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupActivateVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupActivateVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupActivate62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdditionalPath2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdditionalPathVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdditionalPathVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdditionalPath62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdvAdditionalPath2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdvAdditionalPathVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdvAdditionalPathVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdvAdditionalPath62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdvertisementInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasIn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasInEnable2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasInEnableEvpn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasInEnableVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasInEnableVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasInEnable62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasInEvpn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasInVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasInVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasIn62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAsOverride2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAsOverride62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAttributeUnchanged2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupAttributeUnchangedVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupAttributeUnchangedVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupAttributeUnchanged62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupAuthOptions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupBfd2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityDefaultOriginate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityDefaultOriginate62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityDynamic2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityGracefulRestart2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityGracefulRestartEvpn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityGracefulRestartVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityGracefulRestartVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityGracefulRestart62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityOrf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityOrf62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityRouteRefresh2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupConnectTimer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupDefaultOriginateRoutemap2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupDefaultOriginateRoutemap62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupDistributeListIn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupDistributeListInVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupDistributeListInVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupDistributeListIn62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupDistributeListOut2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupDistributeListOutVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupDistributeListOutVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupDistributeListOut62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupDontCapabilityNegotiate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupEbgpEnforceMultihop2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupEbgpMultihopTtl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupFilterListIn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupFilterListInVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupFilterListInVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupFilterListIn62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupFilterListOut2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupFilterListOutVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupFilterListOutVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupFilterListOut62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupHoldtimeTimer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupKeepAliveTimer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupLinkDownFailover2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupLocalAs2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupLocalAsNoPrepend2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupLocalAsReplaceAs2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixEvpn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixThresholdEvpn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixThresholdVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixThresholdVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixThreshold62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixWarningOnly2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixWarningOnlyEvpn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixWarningOnlyVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixWarningOnlyVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixWarningOnly62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefix62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupNextHopSelf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupNextHopSelfRr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupNextHopSelfRr62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupNextHopSelfVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupNextHopSelfVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupNextHopSelf62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupOverrideCapability2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupPassive2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupPassword2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupPrefixListIn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupPrefixListInVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupPrefixListInVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupPrefixListIn62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupPrefixListOut2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupPrefixListOutVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupPrefixListOutVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupPrefixListOut62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRemoteAs2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRemoteAsFilter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRemovePrivateAs2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRemovePrivateAsEvpn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRemovePrivateAsVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRemovePrivateAsVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRemovePrivateAs62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRestartTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRetainStaleTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteMapIn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRouteMapInEvpn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRouteMapInVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRouteMapInVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRouteMapIn62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRouteMapOut2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRouteMapOutEvpn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRouteMapOutPreferable2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRouteMapOutVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRouteMapOutVpnv4Preferable2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRouteMapOutVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRouteMapOutVpnv6Preferable2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRouteMapOut62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRouteMapOut6Preferable2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupRouteReflectorClient2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteReflectorClientEvpn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteReflectorClientVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteReflectorClientVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteReflectorClient62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteServerClient2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteServerClientEvpn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteServerClientVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteServerClientVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteServerClient62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSendCommunity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSendCommunityEvpn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSendCommunityVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSendCommunityVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSendCommunity62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupShutdown2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSoftReconfiguration2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSoftReconfigurationEvpn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSoftReconfigurationVpnv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSoftReconfigurationVpnv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSoftReconfiguration62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupStaleRoute2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupStrictCapabilityMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupUnsuppressMap2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupUnsuppressMap62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupUpdateSource2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborGroupWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterBgpNeighborGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("activate"); ok || d.HasChange("activate") {
		t, err := expandRouterBgpNeighborGroupActivate2edl(d, v, "activate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["activate"] = t
		}
	}

	if v, ok := d.GetOk("activate_evpn"); ok || d.HasChange("activate_evpn") {
		t, err := expandRouterBgpNeighborGroupActivateEvpn2edl(d, v, "activate_evpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["activate-evpn"] = t
		}
	}

	if v, ok := d.GetOk("activate_vpnv4"); ok || d.HasChange("activate_vpnv4") {
		t, err := expandRouterBgpNeighborGroupActivateVpnv42edl(d, v, "activate_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["activate-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("activate_vpnv6"); ok || d.HasChange("activate_vpnv6") {
		t, err := expandRouterBgpNeighborGroupActivateVpnv62edl(d, v, "activate_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["activate-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("activate6"); ok || d.HasChange("activate6") {
		t, err := expandRouterBgpNeighborGroupActivate62edl(d, v, "activate6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["activate6"] = t
		}
	}

	if v, ok := d.GetOk("additional_path"); ok || d.HasChange("additional_path") {
		t, err := expandRouterBgpNeighborGroupAdditionalPath2edl(d, v, "additional_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-path"] = t
		}
	}

	if v, ok := d.GetOk("additional_path_vpnv4"); ok || d.HasChange("additional_path_vpnv4") {
		t, err := expandRouterBgpNeighborGroupAdditionalPathVpnv42edl(d, v, "additional_path_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-path-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("additional_path_vpnv6"); ok || d.HasChange("additional_path_vpnv6") {
		t, err := expandRouterBgpNeighborGroupAdditionalPathVpnv62edl(d, v, "additional_path_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-path-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("additional_path6"); ok || d.HasChange("additional_path6") {
		t, err := expandRouterBgpNeighborGroupAdditionalPath62edl(d, v, "additional_path6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-path6"] = t
		}
	}

	if v, ok := d.GetOk("adv_additional_path"); ok || d.HasChange("adv_additional_path") {
		t, err := expandRouterBgpNeighborGroupAdvAdditionalPath2edl(d, v, "adv_additional_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adv-additional-path"] = t
		}
	}

	if v, ok := d.GetOk("adv_additional_path_vpnv4"); ok || d.HasChange("adv_additional_path_vpnv4") {
		t, err := expandRouterBgpNeighborGroupAdvAdditionalPathVpnv42edl(d, v, "adv_additional_path_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adv-additional-path-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("adv_additional_path_vpnv6"); ok || d.HasChange("adv_additional_path_vpnv6") {
		t, err := expandRouterBgpNeighborGroupAdvAdditionalPathVpnv62edl(d, v, "adv_additional_path_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adv-additional-path-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("adv_additional_path6"); ok || d.HasChange("adv_additional_path6") {
		t, err := expandRouterBgpNeighborGroupAdvAdditionalPath62edl(d, v, "adv_additional_path6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adv-additional-path6"] = t
		}
	}

	if v, ok := d.GetOk("advertisement_interval"); ok || d.HasChange("advertisement_interval") {
		t, err := expandRouterBgpNeighborGroupAdvertisementInterval2edl(d, v, "advertisement_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["advertisement-interval"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in"); ok || d.HasChange("allowas_in") {
		t, err := expandRouterBgpNeighborGroupAllowasIn2edl(d, v, "allowas_in")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in_enable"); ok || d.HasChange("allowas_in_enable") {
		t, err := expandRouterBgpNeighborGroupAllowasInEnable2edl(d, v, "allowas_in_enable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in-enable"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in_enable_evpn"); ok || d.HasChange("allowas_in_enable_evpn") {
		t, err := expandRouterBgpNeighborGroupAllowasInEnableEvpn2edl(d, v, "allowas_in_enable_evpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in-enable-evpn"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in_enable_vpnv4"); ok || d.HasChange("allowas_in_enable_vpnv4") {
		t, err := expandRouterBgpNeighborGroupAllowasInEnableVpnv42edl(d, v, "allowas_in_enable_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in-enable-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in_enable_vpnv6"); ok || d.HasChange("allowas_in_enable_vpnv6") {
		t, err := expandRouterBgpNeighborGroupAllowasInEnableVpnv62edl(d, v, "allowas_in_enable_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in-enable-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in_enable6"); ok || d.HasChange("allowas_in_enable6") {
		t, err := expandRouterBgpNeighborGroupAllowasInEnable62edl(d, v, "allowas_in_enable6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in-enable6"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in_evpn"); ok || d.HasChange("allowas_in_evpn") {
		t, err := expandRouterBgpNeighborGroupAllowasInEvpn2edl(d, v, "allowas_in_evpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in-evpn"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in_vpnv4"); ok || d.HasChange("allowas_in_vpnv4") {
		t, err := expandRouterBgpNeighborGroupAllowasInVpnv42edl(d, v, "allowas_in_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in_vpnv6"); ok || d.HasChange("allowas_in_vpnv6") {
		t, err := expandRouterBgpNeighborGroupAllowasInVpnv62edl(d, v, "allowas_in_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in6"); ok || d.HasChange("allowas_in6") {
		t, err := expandRouterBgpNeighborGroupAllowasIn62edl(d, v, "allowas_in6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in6"] = t
		}
	}

	if v, ok := d.GetOk("as_override"); ok || d.HasChange("as_override") {
		t, err := expandRouterBgpNeighborGroupAsOverride2edl(d, v, "as_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["as-override"] = t
		}
	}

	if v, ok := d.GetOk("as_override6"); ok || d.HasChange("as_override6") {
		t, err := expandRouterBgpNeighborGroupAsOverride62edl(d, v, "as_override6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["as-override6"] = t
		}
	}

	if v, ok := d.GetOk("attribute_unchanged"); ok || d.HasChange("attribute_unchanged") {
		t, err := expandRouterBgpNeighborGroupAttributeUnchanged2edl(d, v, "attribute_unchanged")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attribute-unchanged"] = t
		}
	}

	if v, ok := d.GetOk("attribute_unchanged_vpnv4"); ok || d.HasChange("attribute_unchanged_vpnv4") {
		t, err := expandRouterBgpNeighborGroupAttributeUnchangedVpnv42edl(d, v, "attribute_unchanged_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attribute-unchanged-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("attribute_unchanged_vpnv6"); ok || d.HasChange("attribute_unchanged_vpnv6") {
		t, err := expandRouterBgpNeighborGroupAttributeUnchangedVpnv62edl(d, v, "attribute_unchanged_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attribute-unchanged-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("attribute_unchanged6"); ok || d.HasChange("attribute_unchanged6") {
		t, err := expandRouterBgpNeighborGroupAttributeUnchanged62edl(d, v, "attribute_unchanged6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attribute-unchanged6"] = t
		}
	}

	if v, ok := d.GetOk("auth_options"); ok || d.HasChange("auth_options") {
		t, err := expandRouterBgpNeighborGroupAuthOptions2edl(d, v, "auth_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-options"] = t
		}
	}

	if v, ok := d.GetOk("bfd"); ok || d.HasChange("bfd") {
		t, err := expandRouterBgpNeighborGroupBfd2edl(d, v, "bfd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd"] = t
		}
	}

	if v, ok := d.GetOk("capability_default_originate"); ok || d.HasChange("capability_default_originate") {
		t, err := expandRouterBgpNeighborGroupCapabilityDefaultOriginate2edl(d, v, "capability_default_originate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-default-originate"] = t
		}
	}

	if v, ok := d.GetOk("capability_default_originate6"); ok || d.HasChange("capability_default_originate6") {
		t, err := expandRouterBgpNeighborGroupCapabilityDefaultOriginate62edl(d, v, "capability_default_originate6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-default-originate6"] = t
		}
	}

	if v, ok := d.GetOk("capability_dynamic"); ok || d.HasChange("capability_dynamic") {
		t, err := expandRouterBgpNeighborGroupCapabilityDynamic2edl(d, v, "capability_dynamic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-dynamic"] = t
		}
	}

	if v, ok := d.GetOk("capability_graceful_restart"); ok || d.HasChange("capability_graceful_restart") {
		t, err := expandRouterBgpNeighborGroupCapabilityGracefulRestart2edl(d, v, "capability_graceful_restart")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-graceful-restart"] = t
		}
	}

	if v, ok := d.GetOk("capability_graceful_restart_evpn"); ok || d.HasChange("capability_graceful_restart_evpn") {
		t, err := expandRouterBgpNeighborGroupCapabilityGracefulRestartEvpn2edl(d, v, "capability_graceful_restart_evpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-graceful-restart-evpn"] = t
		}
	}

	if v, ok := d.GetOk("capability_graceful_restart_vpnv4"); ok || d.HasChange("capability_graceful_restart_vpnv4") {
		t, err := expandRouterBgpNeighborGroupCapabilityGracefulRestartVpnv42edl(d, v, "capability_graceful_restart_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-graceful-restart-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("capability_graceful_restart_vpnv6"); ok || d.HasChange("capability_graceful_restart_vpnv6") {
		t, err := expandRouterBgpNeighborGroupCapabilityGracefulRestartVpnv62edl(d, v, "capability_graceful_restart_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-graceful-restart-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("capability_graceful_restart6"); ok || d.HasChange("capability_graceful_restart6") {
		t, err := expandRouterBgpNeighborGroupCapabilityGracefulRestart62edl(d, v, "capability_graceful_restart6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-graceful-restart6"] = t
		}
	}

	if v, ok := d.GetOk("capability_orf"); ok || d.HasChange("capability_orf") {
		t, err := expandRouterBgpNeighborGroupCapabilityOrf2edl(d, v, "capability_orf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-orf"] = t
		}
	}

	if v, ok := d.GetOk("capability_orf6"); ok || d.HasChange("capability_orf6") {
		t, err := expandRouterBgpNeighborGroupCapabilityOrf62edl(d, v, "capability_orf6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-orf6"] = t
		}
	}

	if v, ok := d.GetOk("capability_route_refresh"); ok || d.HasChange("capability_route_refresh") {
		t, err := expandRouterBgpNeighborGroupCapabilityRouteRefresh2edl(d, v, "capability_route_refresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-route-refresh"] = t
		}
	}

	if v, ok := d.GetOk("connect_timer"); ok || d.HasChange("connect_timer") {
		t, err := expandRouterBgpNeighborGroupConnectTimer2edl(d, v, "connect_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connect-timer"] = t
		}
	}

	if v, ok := d.GetOk("default_originate_routemap"); ok || d.HasChange("default_originate_routemap") {
		t, err := expandRouterBgpNeighborGroupDefaultOriginateRoutemap2edl(d, v, "default_originate_routemap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-originate-routemap"] = t
		}
	}

	if v, ok := d.GetOk("default_originate_routemap6"); ok || d.HasChange("default_originate_routemap6") {
		t, err := expandRouterBgpNeighborGroupDefaultOriginateRoutemap62edl(d, v, "default_originate_routemap6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-originate-routemap6"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandRouterBgpNeighborGroupDescription2edl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list_in"); ok || d.HasChange("distribute_list_in") {
		t, err := expandRouterBgpNeighborGroupDistributeListIn2edl(d, v, "distribute_list_in")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-in"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list_in_vpnv4"); ok || d.HasChange("distribute_list_in_vpnv4") {
		t, err := expandRouterBgpNeighborGroupDistributeListInVpnv42edl(d, v, "distribute_list_in_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-in-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list_in_vpnv6"); ok || d.HasChange("distribute_list_in_vpnv6") {
		t, err := expandRouterBgpNeighborGroupDistributeListInVpnv62edl(d, v, "distribute_list_in_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-in-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list_in6"); ok || d.HasChange("distribute_list_in6") {
		t, err := expandRouterBgpNeighborGroupDistributeListIn62edl(d, v, "distribute_list_in6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-in6"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list_out"); ok || d.HasChange("distribute_list_out") {
		t, err := expandRouterBgpNeighborGroupDistributeListOut2edl(d, v, "distribute_list_out")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-out"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list_out_vpnv4"); ok || d.HasChange("distribute_list_out_vpnv4") {
		t, err := expandRouterBgpNeighborGroupDistributeListOutVpnv42edl(d, v, "distribute_list_out_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-out-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list_out_vpnv6"); ok || d.HasChange("distribute_list_out_vpnv6") {
		t, err := expandRouterBgpNeighborGroupDistributeListOutVpnv62edl(d, v, "distribute_list_out_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-out-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list_out6"); ok || d.HasChange("distribute_list_out6") {
		t, err := expandRouterBgpNeighborGroupDistributeListOut62edl(d, v, "distribute_list_out6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-out6"] = t
		}
	}

	if v, ok := d.GetOk("dont_capability_negotiate"); ok || d.HasChange("dont_capability_negotiate") {
		t, err := expandRouterBgpNeighborGroupDontCapabilityNegotiate2edl(d, v, "dont_capability_negotiate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dont-capability-negotiate"] = t
		}
	}

	if v, ok := d.GetOk("ebgp_enforce_multihop"); ok || d.HasChange("ebgp_enforce_multihop") {
		t, err := expandRouterBgpNeighborGroupEbgpEnforceMultihop2edl(d, v, "ebgp_enforce_multihop")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ebgp-enforce-multihop"] = t
		}
	}

	if v, ok := d.GetOk("ebgp_multihop_ttl"); ok || d.HasChange("ebgp_multihop_ttl") {
		t, err := expandRouterBgpNeighborGroupEbgpMultihopTtl2edl(d, v, "ebgp_multihop_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ebgp-multihop-ttl"] = t
		}
	}

	if v, ok := d.GetOk("filter_list_in"); ok || d.HasChange("filter_list_in") {
		t, err := expandRouterBgpNeighborGroupFilterListIn2edl(d, v, "filter_list_in")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-list-in"] = t
		}
	}

	if v, ok := d.GetOk("filter_list_in_vpnv4"); ok || d.HasChange("filter_list_in_vpnv4") {
		t, err := expandRouterBgpNeighborGroupFilterListInVpnv42edl(d, v, "filter_list_in_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-list-in-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("filter_list_in_vpnv6"); ok || d.HasChange("filter_list_in_vpnv6") {
		t, err := expandRouterBgpNeighborGroupFilterListInVpnv62edl(d, v, "filter_list_in_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-list-in-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("filter_list_in6"); ok || d.HasChange("filter_list_in6") {
		t, err := expandRouterBgpNeighborGroupFilterListIn62edl(d, v, "filter_list_in6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-list-in6"] = t
		}
	}

	if v, ok := d.GetOk("filter_list_out"); ok || d.HasChange("filter_list_out") {
		t, err := expandRouterBgpNeighborGroupFilterListOut2edl(d, v, "filter_list_out")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-list-out"] = t
		}
	}

	if v, ok := d.GetOk("filter_list_out_vpnv4"); ok || d.HasChange("filter_list_out_vpnv4") {
		t, err := expandRouterBgpNeighborGroupFilterListOutVpnv42edl(d, v, "filter_list_out_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-list-out-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("filter_list_out_vpnv6"); ok || d.HasChange("filter_list_out_vpnv6") {
		t, err := expandRouterBgpNeighborGroupFilterListOutVpnv62edl(d, v, "filter_list_out_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-list-out-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("filter_list_out6"); ok || d.HasChange("filter_list_out6") {
		t, err := expandRouterBgpNeighborGroupFilterListOut62edl(d, v, "filter_list_out6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-list-out6"] = t
		}
	}

	if v, ok := d.GetOk("holdtime_timer"); ok || d.HasChange("holdtime_timer") {
		t, err := expandRouterBgpNeighborGroupHoldtimeTimer2edl(d, v, "holdtime_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["holdtime-timer"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandRouterBgpNeighborGroupInterface2edl(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("keep_alive_timer"); ok || d.HasChange("keep_alive_timer") {
		t, err := expandRouterBgpNeighborGroupKeepAliveTimer2edl(d, v, "keep_alive_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keep-alive-timer"] = t
		}
	}

	if v, ok := d.GetOk("link_down_failover"); ok || d.HasChange("link_down_failover") {
		t, err := expandRouterBgpNeighborGroupLinkDownFailover2edl(d, v, "link_down_failover")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-down-failover"] = t
		}
	}

	if v, ok := d.GetOk("local_as"); ok || d.HasChange("local_as") {
		t, err := expandRouterBgpNeighborGroupLocalAs2edl(d, v, "local_as")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-as"] = t
		}
	}

	if v, ok := d.GetOk("local_as_no_prepend"); ok || d.HasChange("local_as_no_prepend") {
		t, err := expandRouterBgpNeighborGroupLocalAsNoPrepend2edl(d, v, "local_as_no_prepend")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-as-no-prepend"] = t
		}
	}

	if v, ok := d.GetOk("local_as_replace_as"); ok || d.HasChange("local_as_replace_as") {
		t, err := expandRouterBgpNeighborGroupLocalAsReplaceAs2edl(d, v, "local_as_replace_as")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-as-replace-as"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix"); ok || d.HasChange("maximum_prefix") {
		t, err := expandRouterBgpNeighborGroupMaximumPrefix2edl(d, v, "maximum_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_evpn"); ok || d.HasChange("maximum_prefix_evpn") {
		t, err := expandRouterBgpNeighborGroupMaximumPrefixEvpn2edl(d, v, "maximum_prefix_evpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-evpn"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_threshold"); ok || d.HasChange("maximum_prefix_threshold") {
		t, err := expandRouterBgpNeighborGroupMaximumPrefixThreshold2edl(d, v, "maximum_prefix_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-threshold"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_threshold_evpn"); ok || d.HasChange("maximum_prefix_threshold_evpn") {
		t, err := expandRouterBgpNeighborGroupMaximumPrefixThresholdEvpn2edl(d, v, "maximum_prefix_threshold_evpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-threshold-evpn"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_threshold_vpnv4"); ok || d.HasChange("maximum_prefix_threshold_vpnv4") {
		t, err := expandRouterBgpNeighborGroupMaximumPrefixThresholdVpnv42edl(d, v, "maximum_prefix_threshold_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-threshold-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_threshold_vpnv6"); ok || d.HasChange("maximum_prefix_threshold_vpnv6") {
		t, err := expandRouterBgpNeighborGroupMaximumPrefixThresholdVpnv62edl(d, v, "maximum_prefix_threshold_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-threshold-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_threshold6"); ok || d.HasChange("maximum_prefix_threshold6") {
		t, err := expandRouterBgpNeighborGroupMaximumPrefixThreshold62edl(d, v, "maximum_prefix_threshold6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-threshold6"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_vpnv4"); ok || d.HasChange("maximum_prefix_vpnv4") {
		t, err := expandRouterBgpNeighborGroupMaximumPrefixVpnv42edl(d, v, "maximum_prefix_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_vpnv6"); ok || d.HasChange("maximum_prefix_vpnv6") {
		t, err := expandRouterBgpNeighborGroupMaximumPrefixVpnv62edl(d, v, "maximum_prefix_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_warning_only"); ok || d.HasChange("maximum_prefix_warning_only") {
		t, err := expandRouterBgpNeighborGroupMaximumPrefixWarningOnly2edl(d, v, "maximum_prefix_warning_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-warning-only"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_warning_only_evpn"); ok || d.HasChange("maximum_prefix_warning_only_evpn") {
		t, err := expandRouterBgpNeighborGroupMaximumPrefixWarningOnlyEvpn2edl(d, v, "maximum_prefix_warning_only_evpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-warning-only-evpn"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_warning_only_vpnv4"); ok || d.HasChange("maximum_prefix_warning_only_vpnv4") {
		t, err := expandRouterBgpNeighborGroupMaximumPrefixWarningOnlyVpnv42edl(d, v, "maximum_prefix_warning_only_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-warning-only-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_warning_only_vpnv6"); ok || d.HasChange("maximum_prefix_warning_only_vpnv6") {
		t, err := expandRouterBgpNeighborGroupMaximumPrefixWarningOnlyVpnv62edl(d, v, "maximum_prefix_warning_only_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-warning-only-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_warning_only6"); ok || d.HasChange("maximum_prefix_warning_only6") {
		t, err := expandRouterBgpNeighborGroupMaximumPrefixWarningOnly62edl(d, v, "maximum_prefix_warning_only6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-warning-only6"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix6"); ok || d.HasChange("maximum_prefix6") {
		t, err := expandRouterBgpNeighborGroupMaximumPrefix62edl(d, v, "maximum_prefix6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix6"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandRouterBgpNeighborGroupName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("next_hop_self"); ok || d.HasChange("next_hop_self") {
		t, err := expandRouterBgpNeighborGroupNextHopSelf2edl(d, v, "next_hop_self")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-hop-self"] = t
		}
	}

	if v, ok := d.GetOk("next_hop_self_rr"); ok || d.HasChange("next_hop_self_rr") {
		t, err := expandRouterBgpNeighborGroupNextHopSelfRr2edl(d, v, "next_hop_self_rr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-hop-self-rr"] = t
		}
	}

	if v, ok := d.GetOk("next_hop_self_rr6"); ok || d.HasChange("next_hop_self_rr6") {
		t, err := expandRouterBgpNeighborGroupNextHopSelfRr62edl(d, v, "next_hop_self_rr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-hop-self-rr6"] = t
		}
	}

	if v, ok := d.GetOk("next_hop_self_vpnv4"); ok || d.HasChange("next_hop_self_vpnv4") {
		t, err := expandRouterBgpNeighborGroupNextHopSelfVpnv42edl(d, v, "next_hop_self_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-hop-self-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("next_hop_self_vpnv6"); ok || d.HasChange("next_hop_self_vpnv6") {
		t, err := expandRouterBgpNeighborGroupNextHopSelfVpnv62edl(d, v, "next_hop_self_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-hop-self-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("next_hop_self6"); ok || d.HasChange("next_hop_self6") {
		t, err := expandRouterBgpNeighborGroupNextHopSelf62edl(d, v, "next_hop_self6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-hop-self6"] = t
		}
	}

	if v, ok := d.GetOk("override_capability"); ok || d.HasChange("override_capability") {
		t, err := expandRouterBgpNeighborGroupOverrideCapability2edl(d, v, "override_capability")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-capability"] = t
		}
	}

	if v, ok := d.GetOk("passive"); ok || d.HasChange("passive") {
		t, err := expandRouterBgpNeighborGroupPassive2edl(d, v, "passive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passive"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandRouterBgpNeighborGroupPassword2edl(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("prefix_list_in"); ok || d.HasChange("prefix_list_in") {
		t, err := expandRouterBgpNeighborGroupPrefixListIn2edl(d, v, "prefix_list_in")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-in"] = t
		}
	}

	if v, ok := d.GetOk("prefix_list_in_vpnv4"); ok || d.HasChange("prefix_list_in_vpnv4") {
		t, err := expandRouterBgpNeighborGroupPrefixListInVpnv42edl(d, v, "prefix_list_in_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-in-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("prefix_list_in_vpnv6"); ok || d.HasChange("prefix_list_in_vpnv6") {
		t, err := expandRouterBgpNeighborGroupPrefixListInVpnv62edl(d, v, "prefix_list_in_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-in-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("prefix_list_in6"); ok || d.HasChange("prefix_list_in6") {
		t, err := expandRouterBgpNeighborGroupPrefixListIn62edl(d, v, "prefix_list_in6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-in6"] = t
		}
	}

	if v, ok := d.GetOk("prefix_list_out"); ok || d.HasChange("prefix_list_out") {
		t, err := expandRouterBgpNeighborGroupPrefixListOut2edl(d, v, "prefix_list_out")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-out"] = t
		}
	}

	if v, ok := d.GetOk("prefix_list_out_vpnv4"); ok || d.HasChange("prefix_list_out_vpnv4") {
		t, err := expandRouterBgpNeighborGroupPrefixListOutVpnv42edl(d, v, "prefix_list_out_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-out-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("prefix_list_out_vpnv6"); ok || d.HasChange("prefix_list_out_vpnv6") {
		t, err := expandRouterBgpNeighborGroupPrefixListOutVpnv62edl(d, v, "prefix_list_out_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-out-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("prefix_list_out6"); ok || d.HasChange("prefix_list_out6") {
		t, err := expandRouterBgpNeighborGroupPrefixListOut62edl(d, v, "prefix_list_out6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-out6"] = t
		}
	}

	if v, ok := d.GetOk("remote_as"); ok || d.HasChange("remote_as") {
		t, err := expandRouterBgpNeighborGroupRemoteAs2edl(d, v, "remote_as")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-as"] = t
		}
	}

	if v, ok := d.GetOk("remote_as_filter"); ok || d.HasChange("remote_as_filter") {
		t, err := expandRouterBgpNeighborGroupRemoteAsFilter2edl(d, v, "remote_as_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-as-filter"] = t
		}
	}

	if v, ok := d.GetOk("remove_private_as"); ok || d.HasChange("remove_private_as") {
		t, err := expandRouterBgpNeighborGroupRemovePrivateAs2edl(d, v, "remove_private_as")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remove-private-as"] = t
		}
	}

	if v, ok := d.GetOk("remove_private_as_evpn"); ok || d.HasChange("remove_private_as_evpn") {
		t, err := expandRouterBgpNeighborGroupRemovePrivateAsEvpn2edl(d, v, "remove_private_as_evpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remove-private-as-evpn"] = t
		}
	}

	if v, ok := d.GetOk("remove_private_as_vpnv4"); ok || d.HasChange("remove_private_as_vpnv4") {
		t, err := expandRouterBgpNeighborGroupRemovePrivateAsVpnv42edl(d, v, "remove_private_as_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remove-private-as-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("remove_private_as_vpnv6"); ok || d.HasChange("remove_private_as_vpnv6") {
		t, err := expandRouterBgpNeighborGroupRemovePrivateAsVpnv62edl(d, v, "remove_private_as_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remove-private-as-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("remove_private_as6"); ok || d.HasChange("remove_private_as6") {
		t, err := expandRouterBgpNeighborGroupRemovePrivateAs62edl(d, v, "remove_private_as6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remove-private-as6"] = t
		}
	}

	if v, ok := d.GetOk("restart_time"); ok || d.HasChange("restart_time") {
		t, err := expandRouterBgpNeighborGroupRestartTime2edl(d, v, "restart_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restart-time"] = t
		}
	}

	if v, ok := d.GetOk("retain_stale_time"); ok || d.HasChange("retain_stale_time") {
		t, err := expandRouterBgpNeighborGroupRetainStaleTime2edl(d, v, "retain_stale_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retain-stale-time"] = t
		}
	}

	if v, ok := d.GetOk("route_map_in"); ok || d.HasChange("route_map_in") {
		t, err := expandRouterBgpNeighborGroupRouteMapIn2edl(d, v, "route_map_in")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-in"] = t
		}
	}

	if v, ok := d.GetOk("route_map_in_evpn"); ok || d.HasChange("route_map_in_evpn") {
		t, err := expandRouterBgpNeighborGroupRouteMapInEvpn2edl(d, v, "route_map_in_evpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-in-evpn"] = t
		}
	}

	if v, ok := d.GetOk("route_map_in_vpnv4"); ok || d.HasChange("route_map_in_vpnv4") {
		t, err := expandRouterBgpNeighborGroupRouteMapInVpnv42edl(d, v, "route_map_in_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-in-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("route_map_in_vpnv6"); ok || d.HasChange("route_map_in_vpnv6") {
		t, err := expandRouterBgpNeighborGroupRouteMapInVpnv62edl(d, v, "route_map_in_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-in-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("route_map_in6"); ok || d.HasChange("route_map_in6") {
		t, err := expandRouterBgpNeighborGroupRouteMapIn62edl(d, v, "route_map_in6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-in6"] = t
		}
	}

	if v, ok := d.GetOk("route_map_out"); ok || d.HasChange("route_map_out") {
		t, err := expandRouterBgpNeighborGroupRouteMapOut2edl(d, v, "route_map_out")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out"] = t
		}
	}

	if v, ok := d.GetOk("route_map_out_evpn"); ok || d.HasChange("route_map_out_evpn") {
		t, err := expandRouterBgpNeighborGroupRouteMapOutEvpn2edl(d, v, "route_map_out_evpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out-evpn"] = t
		}
	}

	if v, ok := d.GetOk("route_map_out_preferable"); ok || d.HasChange("route_map_out_preferable") {
		t, err := expandRouterBgpNeighborGroupRouteMapOutPreferable2edl(d, v, "route_map_out_preferable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out-preferable"] = t
		}
	}

	if v, ok := d.GetOk("route_map_out_vpnv4"); ok || d.HasChange("route_map_out_vpnv4") {
		t, err := expandRouterBgpNeighborGroupRouteMapOutVpnv42edl(d, v, "route_map_out_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("route_map_out_vpnv4_preferable"); ok || d.HasChange("route_map_out_vpnv4_preferable") {
		t, err := expandRouterBgpNeighborGroupRouteMapOutVpnv4Preferable2edl(d, v, "route_map_out_vpnv4_preferable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out-vpnv4-preferable"] = t
		}
	}

	if v, ok := d.GetOk("route_map_out_vpnv6"); ok || d.HasChange("route_map_out_vpnv6") {
		t, err := expandRouterBgpNeighborGroupRouteMapOutVpnv62edl(d, v, "route_map_out_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("route_map_out_vpnv6_preferable"); ok || d.HasChange("route_map_out_vpnv6_preferable") {
		t, err := expandRouterBgpNeighborGroupRouteMapOutVpnv6Preferable2edl(d, v, "route_map_out_vpnv6_preferable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out-vpnv6-preferable"] = t
		}
	}

	if v, ok := d.GetOk("route_map_out6"); ok || d.HasChange("route_map_out6") {
		t, err := expandRouterBgpNeighborGroupRouteMapOut62edl(d, v, "route_map_out6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out6"] = t
		}
	}

	if v, ok := d.GetOk("route_map_out6_preferable"); ok || d.HasChange("route_map_out6_preferable") {
		t, err := expandRouterBgpNeighborGroupRouteMapOut6Preferable2edl(d, v, "route_map_out6_preferable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out6-preferable"] = t
		}
	}

	if v, ok := d.GetOk("route_reflector_client"); ok || d.HasChange("route_reflector_client") {
		t, err := expandRouterBgpNeighborGroupRouteReflectorClient2edl(d, v, "route_reflector_client")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-reflector-client"] = t
		}
	}

	if v, ok := d.GetOk("route_reflector_client_evpn"); ok || d.HasChange("route_reflector_client_evpn") {
		t, err := expandRouterBgpNeighborGroupRouteReflectorClientEvpn2edl(d, v, "route_reflector_client_evpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-reflector-client-evpn"] = t
		}
	}

	if v, ok := d.GetOk("route_reflector_client_vpnv4"); ok || d.HasChange("route_reflector_client_vpnv4") {
		t, err := expandRouterBgpNeighborGroupRouteReflectorClientVpnv42edl(d, v, "route_reflector_client_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-reflector-client-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("route_reflector_client_vpnv6"); ok || d.HasChange("route_reflector_client_vpnv6") {
		t, err := expandRouterBgpNeighborGroupRouteReflectorClientVpnv62edl(d, v, "route_reflector_client_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-reflector-client-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("route_reflector_client6"); ok || d.HasChange("route_reflector_client6") {
		t, err := expandRouterBgpNeighborGroupRouteReflectorClient62edl(d, v, "route_reflector_client6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-reflector-client6"] = t
		}
	}

	if v, ok := d.GetOk("route_server_client"); ok || d.HasChange("route_server_client") {
		t, err := expandRouterBgpNeighborGroupRouteServerClient2edl(d, v, "route_server_client")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-server-client"] = t
		}
	}

	if v, ok := d.GetOk("route_server_client_evpn"); ok || d.HasChange("route_server_client_evpn") {
		t, err := expandRouterBgpNeighborGroupRouteServerClientEvpn2edl(d, v, "route_server_client_evpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-server-client-evpn"] = t
		}
	}

	if v, ok := d.GetOk("route_server_client_vpnv4"); ok || d.HasChange("route_server_client_vpnv4") {
		t, err := expandRouterBgpNeighborGroupRouteServerClientVpnv42edl(d, v, "route_server_client_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-server-client-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("route_server_client_vpnv6"); ok || d.HasChange("route_server_client_vpnv6") {
		t, err := expandRouterBgpNeighborGroupRouteServerClientVpnv62edl(d, v, "route_server_client_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-server-client-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("route_server_client6"); ok || d.HasChange("route_server_client6") {
		t, err := expandRouterBgpNeighborGroupRouteServerClient62edl(d, v, "route_server_client6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-server-client6"] = t
		}
	}

	if v, ok := d.GetOk("send_community"); ok || d.HasChange("send_community") {
		t, err := expandRouterBgpNeighborGroupSendCommunity2edl(d, v, "send_community")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-community"] = t
		}
	}

	if v, ok := d.GetOk("send_community_evpn"); ok || d.HasChange("send_community_evpn") {
		t, err := expandRouterBgpNeighborGroupSendCommunityEvpn2edl(d, v, "send_community_evpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-community-evpn"] = t
		}
	}

	if v, ok := d.GetOk("send_community_vpnv4"); ok || d.HasChange("send_community_vpnv4") {
		t, err := expandRouterBgpNeighborGroupSendCommunityVpnv42edl(d, v, "send_community_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-community-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("send_community_vpnv6"); ok || d.HasChange("send_community_vpnv6") {
		t, err := expandRouterBgpNeighborGroupSendCommunityVpnv62edl(d, v, "send_community_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-community-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("send_community6"); ok || d.HasChange("send_community6") {
		t, err := expandRouterBgpNeighborGroupSendCommunity62edl(d, v, "send_community6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-community6"] = t
		}
	}

	if v, ok := d.GetOk("shutdown"); ok || d.HasChange("shutdown") {
		t, err := expandRouterBgpNeighborGroupShutdown2edl(d, v, "shutdown")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shutdown"] = t
		}
	}

	if v, ok := d.GetOk("soft_reconfiguration"); ok || d.HasChange("soft_reconfiguration") {
		t, err := expandRouterBgpNeighborGroupSoftReconfiguration2edl(d, v, "soft_reconfiguration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["soft-reconfiguration"] = t
		}
	}

	if v, ok := d.GetOk("soft_reconfiguration_evpn"); ok || d.HasChange("soft_reconfiguration_evpn") {
		t, err := expandRouterBgpNeighborGroupSoftReconfigurationEvpn2edl(d, v, "soft_reconfiguration_evpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["soft-reconfiguration-evpn"] = t
		}
	}

	if v, ok := d.GetOk("soft_reconfiguration_vpnv4"); ok || d.HasChange("soft_reconfiguration_vpnv4") {
		t, err := expandRouterBgpNeighborGroupSoftReconfigurationVpnv42edl(d, v, "soft_reconfiguration_vpnv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["soft-reconfiguration-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("soft_reconfiguration_vpnv6"); ok || d.HasChange("soft_reconfiguration_vpnv6") {
		t, err := expandRouterBgpNeighborGroupSoftReconfigurationVpnv62edl(d, v, "soft_reconfiguration_vpnv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["soft-reconfiguration-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("soft_reconfiguration6"); ok || d.HasChange("soft_reconfiguration6") {
		t, err := expandRouterBgpNeighborGroupSoftReconfiguration62edl(d, v, "soft_reconfiguration6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["soft-reconfiguration6"] = t
		}
	}

	if v, ok := d.GetOk("stale_route"); ok || d.HasChange("stale_route") {
		t, err := expandRouterBgpNeighborGroupStaleRoute2edl(d, v, "stale_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stale-route"] = t
		}
	}

	if v, ok := d.GetOk("strict_capability_match"); ok || d.HasChange("strict_capability_match") {
		t, err := expandRouterBgpNeighborGroupStrictCapabilityMatch2edl(d, v, "strict_capability_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strict-capability-match"] = t
		}
	}

	if v, ok := d.GetOk("unsuppress_map"); ok || d.HasChange("unsuppress_map") {
		t, err := expandRouterBgpNeighborGroupUnsuppressMap2edl(d, v, "unsuppress_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unsuppress-map"] = t
		}
	}

	if v, ok := d.GetOk("unsuppress_map6"); ok || d.HasChange("unsuppress_map6") {
		t, err := expandRouterBgpNeighborGroupUnsuppressMap62edl(d, v, "unsuppress_map6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unsuppress-map6"] = t
		}
	}

	if v, ok := d.GetOk("update_source"); ok || d.HasChange("update_source") {
		t, err := expandRouterBgpNeighborGroupUpdateSource2edl(d, v, "update_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-source"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandRouterBgpNeighborGroupWeight2edl(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	return &obj, nil
}
