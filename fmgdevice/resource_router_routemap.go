// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure route maps.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterRouteMap() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterRouteMapCreate,
		Read:   resourceRouterRouteMapRead,
		Update: resourceRouterRouteMapUpdate,
		Delete: resourceRouterRouteMapDelete,

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
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"match_as_path": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"match_community": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"match_community_exact": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"match_extcommunity": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"match_extcommunity_exact": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"match_flags": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"match_interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"match_ip_address": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"match_ip_nexthop": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"match_ip6_address": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"match_ip6_nexthop": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"match_metric": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"match_origin": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"match_route_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"match_tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"match_vrf": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"set_aggregator_as": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"set_aggregator_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"set_aspath": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"set_aspath_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"set_atomic_aggregate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"set_community": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"set_community_additive": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"set_community_delete": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"set_dampening_max_suppress": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"set_dampening_reachability_half_life": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"set_dampening_reuse": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"set_dampening_suppress": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"set_dampening_unreachability_half_life": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"set_extcommunity_rt": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"set_extcommunity_soo": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"set_flags": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"set_ip_nexthop": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"set_ip_prefsrc": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"set_ip6_nexthop": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"set_ip6_nexthop_local": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"set_local_preference": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"set_metric": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"set_metric_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"set_origin": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"set_originator_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"set_priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"set_route_tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"set_tag": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"set_vpnv4_nexthop": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"set_vpnv6_nexthop": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"set_vpnv6_nexthop_local": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"set_weight": &schema.Schema{
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

func resourceRouterRouteMapCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterRouteMap(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterRouteMap resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadRouterRouteMap(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateRouterRouteMap(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating RouterRouteMap resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateRouterRouteMap(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating RouterRouteMap resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceRouterRouteMapRead(d, m)
}

func resourceRouterRouteMapUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterRouteMap(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterRouteMap resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateRouterRouteMap(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterRouteMap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceRouterRouteMapRead(d, m)
}

func resourceRouterRouteMapDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterRouteMap(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterRouteMap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterRouteMapRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterRouteMap(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading RouterRouteMap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterRouteMap(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterRouteMap resource from API: %v", err)
	}
	return nil
}

func flattenRouterRouteMapComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRule(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenRouterRouteMapRuleAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterRouteMapRuleId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_as_path"
		if _, ok := i["match-as-path"]; ok {
			v := flattenRouterRouteMapRuleMatchAsPath(i["match-as-path"], d, pre_append)
			tmp["match_as_path"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-MatchAsPath")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_community"
		if _, ok := i["match-community"]; ok {
			v := flattenRouterRouteMapRuleMatchCommunity(i["match-community"], d, pre_append)
			tmp["match_community"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-MatchCommunity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_community_exact"
		if _, ok := i["match-community-exact"]; ok {
			v := flattenRouterRouteMapRuleMatchCommunityExact(i["match-community-exact"], d, pre_append)
			tmp["match_community_exact"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-MatchCommunityExact")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_extcommunity"
		if _, ok := i["match-extcommunity"]; ok {
			v := flattenRouterRouteMapRuleMatchExtcommunity(i["match-extcommunity"], d, pre_append)
			tmp["match_extcommunity"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-MatchExtcommunity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_extcommunity_exact"
		if _, ok := i["match-extcommunity-exact"]; ok {
			v := flattenRouterRouteMapRuleMatchExtcommunityExact(i["match-extcommunity-exact"], d, pre_append)
			tmp["match_extcommunity_exact"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-MatchExtcommunityExact")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_flags"
		if _, ok := i["match-flags"]; ok {
			v := flattenRouterRouteMapRuleMatchFlags(i["match-flags"], d, pre_append)
			tmp["match_flags"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-MatchFlags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_interface"
		if _, ok := i["match-interface"]; ok {
			v := flattenRouterRouteMapRuleMatchInterface(i["match-interface"], d, pre_append)
			tmp["match_interface"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-MatchInterface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_ip_address"
		if _, ok := i["match-ip-address"]; ok {
			v := flattenRouterRouteMapRuleMatchIpAddress(i["match-ip-address"], d, pre_append)
			tmp["match_ip_address"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-MatchIpAddress")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_ip_nexthop"
		if _, ok := i["match-ip-nexthop"]; ok {
			v := flattenRouterRouteMapRuleMatchIpNexthop(i["match-ip-nexthop"], d, pre_append)
			tmp["match_ip_nexthop"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-MatchIpNexthop")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_ip6_address"
		if _, ok := i["match-ip6-address"]; ok {
			v := flattenRouterRouteMapRuleMatchIp6Address(i["match-ip6-address"], d, pre_append)
			tmp["match_ip6_address"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-MatchIp6Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_ip6_nexthop"
		if _, ok := i["match-ip6-nexthop"]; ok {
			v := flattenRouterRouteMapRuleMatchIp6Nexthop(i["match-ip6-nexthop"], d, pre_append)
			tmp["match_ip6_nexthop"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-MatchIp6Nexthop")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_metric"
		if _, ok := i["match-metric"]; ok {
			v := flattenRouterRouteMapRuleMatchMetric(i["match-metric"], d, pre_append)
			tmp["match_metric"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-MatchMetric")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_origin"
		if _, ok := i["match-origin"]; ok {
			v := flattenRouterRouteMapRuleMatchOrigin(i["match-origin"], d, pre_append)
			tmp["match_origin"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-MatchOrigin")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_route_type"
		if _, ok := i["match-route-type"]; ok {
			v := flattenRouterRouteMapRuleMatchRouteType(i["match-route-type"], d, pre_append)
			tmp["match_route_type"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-MatchRouteType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_tag"
		if _, ok := i["match-tag"]; ok {
			v := flattenRouterRouteMapRuleMatchTag(i["match-tag"], d, pre_append)
			tmp["match_tag"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-MatchTag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_vrf"
		if _, ok := i["match-vrf"]; ok {
			v := flattenRouterRouteMapRuleMatchVrf(i["match-vrf"], d, pre_append)
			tmp["match_vrf"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-MatchVrf")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_aggregator_as"
		if _, ok := i["set-aggregator-as"]; ok {
			v := flattenRouterRouteMapRuleSetAggregatorAs(i["set-aggregator-as"], d, pre_append)
			tmp["set_aggregator_as"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetAggregatorAs")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_aggregator_ip"
		if _, ok := i["set-aggregator-ip"]; ok {
			v := flattenRouterRouteMapRuleSetAggregatorIp(i["set-aggregator-ip"], d, pre_append)
			tmp["set_aggregator_ip"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetAggregatorIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_aspath"
		if _, ok := i["set-aspath"]; ok {
			v := flattenRouterRouteMapRuleSetAspath(i["set-aspath"], d, pre_append)
			tmp["set_aspath"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetAspath")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_aspath_action"
		if _, ok := i["set-aspath-action"]; ok {
			v := flattenRouterRouteMapRuleSetAspathAction(i["set-aspath-action"], d, pre_append)
			tmp["set_aspath_action"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetAspathAction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_atomic_aggregate"
		if _, ok := i["set-atomic-aggregate"]; ok {
			v := flattenRouterRouteMapRuleSetAtomicAggregate(i["set-atomic-aggregate"], d, pre_append)
			tmp["set_atomic_aggregate"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetAtomicAggregate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_community"
		if _, ok := i["set-community"]; ok {
			v := flattenRouterRouteMapRuleSetCommunity(i["set-community"], d, pre_append)
			tmp["set_community"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetCommunity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_community_additive"
		if _, ok := i["set-community-additive"]; ok {
			v := flattenRouterRouteMapRuleSetCommunityAdditive(i["set-community-additive"], d, pre_append)
			tmp["set_community_additive"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetCommunityAdditive")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_community_delete"
		if _, ok := i["set-community-delete"]; ok {
			v := flattenRouterRouteMapRuleSetCommunityDelete(i["set-community-delete"], d, pre_append)
			tmp["set_community_delete"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetCommunityDelete")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_max_suppress"
		if _, ok := i["set-dampening-max-suppress"]; ok {
			v := flattenRouterRouteMapRuleSetDampeningMaxSuppress(i["set-dampening-max-suppress"], d, pre_append)
			tmp["set_dampening_max_suppress"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetDampeningMaxSuppress")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_reachability_half_life"
		if _, ok := i["set-dampening-reachability-half-life"]; ok {
			v := flattenRouterRouteMapRuleSetDampeningReachabilityHalfLife(i["set-dampening-reachability-half-life"], d, pre_append)
			tmp["set_dampening_reachability_half_life"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetDampeningReachabilityHalfLife")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_reuse"
		if _, ok := i["set-dampening-reuse"]; ok {
			v := flattenRouterRouteMapRuleSetDampeningReuse(i["set-dampening-reuse"], d, pre_append)
			tmp["set_dampening_reuse"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetDampeningReuse")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_suppress"
		if _, ok := i["set-dampening-suppress"]; ok {
			v := flattenRouterRouteMapRuleSetDampeningSuppress(i["set-dampening-suppress"], d, pre_append)
			tmp["set_dampening_suppress"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetDampeningSuppress")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_unreachability_half_life"
		if _, ok := i["set-dampening-unreachability-half-life"]; ok {
			v := flattenRouterRouteMapRuleSetDampeningUnreachabilityHalfLife(i["set-dampening-unreachability-half-life"], d, pre_append)
			tmp["set_dampening_unreachability_half_life"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetDampeningUnreachabilityHalfLife")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_extcommunity_rt"
		if _, ok := i["set-extcommunity-rt"]; ok {
			v := flattenRouterRouteMapRuleSetExtcommunityRt(i["set-extcommunity-rt"], d, pre_append)
			tmp["set_extcommunity_rt"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetExtcommunityRt")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_extcommunity_soo"
		if _, ok := i["set-extcommunity-soo"]; ok {
			v := flattenRouterRouteMapRuleSetExtcommunitySoo(i["set-extcommunity-soo"], d, pre_append)
			tmp["set_extcommunity_soo"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetExtcommunitySoo")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_flags"
		if _, ok := i["set-flags"]; ok {
			v := flattenRouterRouteMapRuleSetFlags(i["set-flags"], d, pre_append)
			tmp["set_flags"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetFlags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_ip_nexthop"
		if _, ok := i["set-ip-nexthop"]; ok {
			v := flattenRouterRouteMapRuleSetIpNexthop(i["set-ip-nexthop"], d, pre_append)
			tmp["set_ip_nexthop"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetIpNexthop")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_ip_prefsrc"
		if _, ok := i["set-ip-prefsrc"]; ok {
			v := flattenRouterRouteMapRuleSetIpPrefsrc(i["set-ip-prefsrc"], d, pre_append)
			tmp["set_ip_prefsrc"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetIpPrefsrc")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_ip6_nexthop"
		if _, ok := i["set-ip6-nexthop"]; ok {
			v := flattenRouterRouteMapRuleSetIp6Nexthop(i["set-ip6-nexthop"], d, pre_append)
			tmp["set_ip6_nexthop"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetIp6Nexthop")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_ip6_nexthop_local"
		if _, ok := i["set-ip6-nexthop-local"]; ok {
			v := flattenRouterRouteMapRuleSetIp6NexthopLocal(i["set-ip6-nexthop-local"], d, pre_append)
			tmp["set_ip6_nexthop_local"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetIp6NexthopLocal")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_local_preference"
		if _, ok := i["set-local-preference"]; ok {
			v := flattenRouterRouteMapRuleSetLocalPreference(i["set-local-preference"], d, pre_append)
			tmp["set_local_preference"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetLocalPreference")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_metric"
		if _, ok := i["set-metric"]; ok {
			v := flattenRouterRouteMapRuleSetMetric(i["set-metric"], d, pre_append)
			tmp["set_metric"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetMetric")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_metric_type"
		if _, ok := i["set-metric-type"]; ok {
			v := flattenRouterRouteMapRuleSetMetricType(i["set-metric-type"], d, pre_append)
			tmp["set_metric_type"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetMetricType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_origin"
		if _, ok := i["set-origin"]; ok {
			v := flattenRouterRouteMapRuleSetOrigin(i["set-origin"], d, pre_append)
			tmp["set_origin"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetOrigin")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_originator_id"
		if _, ok := i["set-originator-id"]; ok {
			v := flattenRouterRouteMapRuleSetOriginatorId(i["set-originator-id"], d, pre_append)
			tmp["set_originator_id"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetOriginatorId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_priority"
		if _, ok := i["set-priority"]; ok {
			v := flattenRouterRouteMapRuleSetPriority(i["set-priority"], d, pre_append)
			tmp["set_priority"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetPriority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_route_tag"
		if _, ok := i["set-route-tag"]; ok {
			v := flattenRouterRouteMapRuleSetRouteTag(i["set-route-tag"], d, pre_append)
			tmp["set_route_tag"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetRouteTag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_tag"
		if _, ok := i["set-tag"]; ok {
			v := flattenRouterRouteMapRuleSetTag(i["set-tag"], d, pre_append)
			tmp["set_tag"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetTag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_vpnv4_nexthop"
		if _, ok := i["set-vpnv4-nexthop"]; ok {
			v := flattenRouterRouteMapRuleSetVpnv4Nexthop(i["set-vpnv4-nexthop"], d, pre_append)
			tmp["set_vpnv4_nexthop"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetVpnv4Nexthop")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_vpnv6_nexthop"
		if _, ok := i["set-vpnv6-nexthop"]; ok {
			v := flattenRouterRouteMapRuleSetVpnv6Nexthop(i["set-vpnv6-nexthop"], d, pre_append)
			tmp["set_vpnv6_nexthop"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetVpnv6Nexthop")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_vpnv6_nexthop_local"
		if _, ok := i["set-vpnv6-nexthop-local"]; ok {
			v := flattenRouterRouteMapRuleSetVpnv6NexthopLocal(i["set-vpnv6-nexthop-local"], d, pre_append)
			tmp["set_vpnv6_nexthop_local"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetVpnv6NexthopLocal")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_weight"
		if _, ok := i["set-weight"]; ok {
			v := flattenRouterRouteMapRuleSetWeight(i["set-weight"], d, pre_append)
			tmp["set_weight"] = fortiAPISubPartPatch(v, "RouterRouteMap-Rule-SetWeight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterRouteMapRuleAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleMatchAsPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRouteMapRuleMatchCommunity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRouteMapRuleMatchCommunityExact(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleMatchExtcommunity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRouteMapRuleMatchExtcommunityExact(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleMatchFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleMatchInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRouteMapRuleMatchIpAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRouteMapRuleMatchIpNexthop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRouteMapRuleMatchIp6Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRouteMapRuleMatchIp6Nexthop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRouteMapRuleMatchMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleMatchOrigin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleMatchRouteType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleMatchTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleMatchVrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetAggregatorAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetAggregatorIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetAspath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRouteMapRuleSetAspathAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetAtomicAggregate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetCommunity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRouteMapRuleSetCommunityAdditive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetCommunityDelete(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRouteMapRuleSetDampeningMaxSuppress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetDampeningReachabilityHalfLife(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetDampeningReuse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetDampeningSuppress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetDampeningUnreachabilityHalfLife(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetExtcommunityRt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRouteMapRuleSetExtcommunitySoo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRouteMapRuleSetFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetIpNexthop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetIpPrefsrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetIp6Nexthop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetIp6NexthopLocal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetLocalPreference(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetOrigin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetOriginatorId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetRouteTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetVpnv4Nexthop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetVpnv6Nexthop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetVpnv6NexthopLocal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterRouteMap(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comments", flattenRouterRouteMapComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "RouterRouteMap-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("name", flattenRouterRouteMapName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "RouterRouteMap-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("rule", flattenRouterRouteMapRule(o["rule"], d, "rule")); err != nil {
			if vv, ok := fortiAPIPatch(o["rule"], "RouterRouteMap-Rule"); ok {
				if err = d.Set("rule", vv); err != nil {
					return fmt.Errorf("Error reading rule: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading rule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("rule"); ok {
			if err = d.Set("rule", flattenRouterRouteMapRule(o["rule"], d, "rule")); err != nil {
				if vv, ok := fortiAPIPatch(o["rule"], "RouterRouteMap-Rule"); ok {
					if err = d.Set("rule", vv); err != nil {
						return fmt.Errorf("Error reading rule: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading rule: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenRouterRouteMapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterRouteMapComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandRouterRouteMapRuleAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterRouteMapRuleId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_as_path"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-as-path"], _ = expandRouterRouteMapRuleMatchAsPath(d, i["match_as_path"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_community"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-community"], _ = expandRouterRouteMapRuleMatchCommunity(d, i["match_community"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_community_exact"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-community-exact"], _ = expandRouterRouteMapRuleMatchCommunityExact(d, i["match_community_exact"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_extcommunity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-extcommunity"], _ = expandRouterRouteMapRuleMatchExtcommunity(d, i["match_extcommunity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_extcommunity_exact"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-extcommunity-exact"], _ = expandRouterRouteMapRuleMatchExtcommunityExact(d, i["match_extcommunity_exact"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_flags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-flags"], _ = expandRouterRouteMapRuleMatchFlags(d, i["match_flags"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-interface"], _ = expandRouterRouteMapRuleMatchInterface(d, i["match_interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_ip_address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-ip-address"], _ = expandRouterRouteMapRuleMatchIpAddress(d, i["match_ip_address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_ip_nexthop"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-ip-nexthop"], _ = expandRouterRouteMapRuleMatchIpNexthop(d, i["match_ip_nexthop"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_ip6_address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-ip6-address"], _ = expandRouterRouteMapRuleMatchIp6Address(d, i["match_ip6_address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_ip6_nexthop"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-ip6-nexthop"], _ = expandRouterRouteMapRuleMatchIp6Nexthop(d, i["match_ip6_nexthop"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_metric"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-metric"], _ = expandRouterRouteMapRuleMatchMetric(d, i["match_metric"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_origin"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-origin"], _ = expandRouterRouteMapRuleMatchOrigin(d, i["match_origin"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_route_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-route-type"], _ = expandRouterRouteMapRuleMatchRouteType(d, i["match_route_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_tag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-tag"], _ = expandRouterRouteMapRuleMatchTag(d, i["match_tag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_vrf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-vrf"], _ = expandRouterRouteMapRuleMatchVrf(d, i["match_vrf"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_aggregator_as"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-aggregator-as"], _ = expandRouterRouteMapRuleSetAggregatorAs(d, i["set_aggregator_as"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_aggregator_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-aggregator-ip"], _ = expandRouterRouteMapRuleSetAggregatorIp(d, i["set_aggregator_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_aspath"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-aspath"], _ = expandRouterRouteMapRuleSetAspath(d, i["set_aspath"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_aspath_action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-aspath-action"], _ = expandRouterRouteMapRuleSetAspathAction(d, i["set_aspath_action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_atomic_aggregate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-atomic-aggregate"], _ = expandRouterRouteMapRuleSetAtomicAggregate(d, i["set_atomic_aggregate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_community"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-community"], _ = expandRouterRouteMapRuleSetCommunity(d, i["set_community"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_community_additive"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-community-additive"], _ = expandRouterRouteMapRuleSetCommunityAdditive(d, i["set_community_additive"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_community_delete"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-community-delete"], _ = expandRouterRouteMapRuleSetCommunityDelete(d, i["set_community_delete"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_max_suppress"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-dampening-max-suppress"], _ = expandRouterRouteMapRuleSetDampeningMaxSuppress(d, i["set_dampening_max_suppress"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_reachability_half_life"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-dampening-reachability-half-life"], _ = expandRouterRouteMapRuleSetDampeningReachabilityHalfLife(d, i["set_dampening_reachability_half_life"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_reuse"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-dampening-reuse"], _ = expandRouterRouteMapRuleSetDampeningReuse(d, i["set_dampening_reuse"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_suppress"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-dampening-suppress"], _ = expandRouterRouteMapRuleSetDampeningSuppress(d, i["set_dampening_suppress"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_unreachability_half_life"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-dampening-unreachability-half-life"], _ = expandRouterRouteMapRuleSetDampeningUnreachabilityHalfLife(d, i["set_dampening_unreachability_half_life"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_extcommunity_rt"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-extcommunity-rt"], _ = expandRouterRouteMapRuleSetExtcommunityRt(d, i["set_extcommunity_rt"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_extcommunity_soo"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-extcommunity-soo"], _ = expandRouterRouteMapRuleSetExtcommunitySoo(d, i["set_extcommunity_soo"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_flags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-flags"], _ = expandRouterRouteMapRuleSetFlags(d, i["set_flags"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_ip_nexthop"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-ip-nexthop"], _ = expandRouterRouteMapRuleSetIpNexthop(d, i["set_ip_nexthop"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_ip_prefsrc"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-ip-prefsrc"], _ = expandRouterRouteMapRuleSetIpPrefsrc(d, i["set_ip_prefsrc"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_ip6_nexthop"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-ip6-nexthop"], _ = expandRouterRouteMapRuleSetIp6Nexthop(d, i["set_ip6_nexthop"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_ip6_nexthop_local"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-ip6-nexthop-local"], _ = expandRouterRouteMapRuleSetIp6NexthopLocal(d, i["set_ip6_nexthop_local"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_local_preference"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-local-preference"], _ = expandRouterRouteMapRuleSetLocalPreference(d, i["set_local_preference"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_metric"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-metric"], _ = expandRouterRouteMapRuleSetMetric(d, i["set_metric"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_metric_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-metric-type"], _ = expandRouterRouteMapRuleSetMetricType(d, i["set_metric_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_origin"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-origin"], _ = expandRouterRouteMapRuleSetOrigin(d, i["set_origin"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_originator_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-originator-id"], _ = expandRouterRouteMapRuleSetOriginatorId(d, i["set_originator_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-priority"], _ = expandRouterRouteMapRuleSetPriority(d, i["set_priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_route_tag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-route-tag"], _ = expandRouterRouteMapRuleSetRouteTag(d, i["set_route_tag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_tag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-tag"], _ = expandRouterRouteMapRuleSetTag(d, i["set_tag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_vpnv4_nexthop"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-vpnv4-nexthop"], _ = expandRouterRouteMapRuleSetVpnv4Nexthop(d, i["set_vpnv4_nexthop"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_vpnv6_nexthop"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-vpnv6-nexthop"], _ = expandRouterRouteMapRuleSetVpnv6Nexthop(d, i["set_vpnv6_nexthop"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_vpnv6_nexthop_local"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-vpnv6-nexthop-local"], _ = expandRouterRouteMapRuleSetVpnv6NexthopLocal(d, i["set_vpnv6_nexthop_local"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-weight"], _ = expandRouterRouteMapRuleSetWeight(d, i["set_weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterRouteMapRuleAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchAsPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRouteMapRuleMatchCommunity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRouteMapRuleMatchCommunityExact(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchExtcommunity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRouteMapRuleMatchExtcommunityExact(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRouteMapRuleMatchIpAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRouteMapRuleMatchIpNexthop(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRouteMapRuleMatchIp6Address(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRouteMapRuleMatchIp6Nexthop(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRouteMapRuleMatchMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchOrigin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchRouteType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchVrf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetAggregatorAs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetAggregatorIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetAspath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRouteMapRuleSetAspathAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetAtomicAggregate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetCommunity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRouteMapRuleSetCommunityAdditive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetCommunityDelete(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRouteMapRuleSetDampeningMaxSuppress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetDampeningReachabilityHalfLife(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetDampeningReuse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetDampeningSuppress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetDampeningUnreachabilityHalfLife(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetExtcommunityRt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRouteMapRuleSetExtcommunitySoo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRouteMapRuleSetFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetIpNexthop(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetIpPrefsrc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetIp6Nexthop(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetIp6NexthopLocal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetLocalPreference(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetMetricType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetOrigin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetOriginatorId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetRouteTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetVpnv4Nexthop(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetVpnv6Nexthop(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetVpnv6NexthopLocal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterRouteMap(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandRouterRouteMapComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandRouterRouteMapName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("rule"); ok || d.HasChange("rule") {
		t, err := expandRouterRouteMapRule(d, v, "rule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule"] = t
		}
	}

	return &obj, nil
}
