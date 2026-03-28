// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Rule.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterRouteMapRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterRouteMapRuleCreate,
		Read:   resourceRouterRouteMapRuleRead,
		Update: resourceRouterRouteMapRuleUpdate,
		Delete: resourceRouterRouteMapRuleDelete,

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
			"route_map": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
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
	}
}

func resourceRouterRouteMapRuleCreate(d *schema.ResourceData, m interface{}) error {
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
	route_map := d.Get("route_map").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["route_map"] = route_map

	obj, err := getObjectRouterRouteMapRule(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterRouteMapRule resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadRouterRouteMapRule(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateRouterRouteMapRule(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating RouterRouteMapRule resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateRouterRouteMapRule(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating RouterRouteMapRule resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceRouterRouteMapRuleRead(d, m)
			} else {
				return fmt.Errorf("Error creating RouterRouteMapRule resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterRouteMapRuleRead(d, m)
}

func resourceRouterRouteMapRuleUpdate(d *schema.ResourceData, m interface{}) error {
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
	route_map := d.Get("route_map").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["route_map"] = route_map

	obj, err := getObjectRouterRouteMapRule(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterRouteMapRule resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateRouterRouteMapRule(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterRouteMapRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterRouteMapRuleRead(d, m)
}

func resourceRouterRouteMapRuleDelete(d *schema.ResourceData, m interface{}) error {
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
	route_map := d.Get("route_map").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["route_map"] = route_map

	wsParams["adom"] = adomv

	err = c.DeleteRouterRouteMapRule(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterRouteMapRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterRouteMapRuleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	route_map := d.Get("route_map").(string)
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
	if route_map == "" {
		route_map = importOptionChecking(m.(*FortiClient).Cfg, "route_map")
		if route_map == "" {
			return fmt.Errorf("Parameter route_map is missing")
		}
		if err = d.Set("route_map", route_map); err != nil {
			return fmt.Errorf("Error set params route_map: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["route_map"] = route_map

	o, err := c.ReadRouterRouteMapRule(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading RouterRouteMapRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterRouteMapRule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterRouteMapRule resource from API: %v", err)
	}
	return nil
}

func flattenRouterRouteMapRuleAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleMatchAsPath2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRouteMapRuleMatchCommunity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRouteMapRuleMatchCommunityExact2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleMatchExtcommunity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRouteMapRuleMatchExtcommunityExact2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleMatchFlags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleMatchInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRouteMapRuleMatchIpAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRouteMapRuleMatchIpNexthop2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRouteMapRuleMatchIp6Address2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRouteMapRuleMatchIp6Nexthop2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRouteMapRuleMatchMetric2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleMatchOrigin2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleMatchRouteType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleMatchTag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleMatchVrf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetAggregatorAs2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetAggregatorIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetAspath2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRouteMapRuleSetAspathAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetAtomicAggregate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetCommunity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRouteMapRuleSetCommunityAdditive2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetCommunityDelete2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRouteMapRuleSetDampeningMaxSuppress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetDampeningReachabilityHalfLife2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetDampeningReuse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetDampeningSuppress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetDampeningUnreachabilityHalfLife2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetExtcommunityRt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRouteMapRuleSetExtcommunitySoo2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRouteMapRuleSetFlags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetIpNexthop2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetIpPrefsrc2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetIp6Nexthop2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetIp6NexthopLocal2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetLocalPreference2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetMetric2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetMetricType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetOrigin2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetOriginatorId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetRouteTag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetTag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetVpnv4Nexthop2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetVpnv6Nexthop2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetVpnv6NexthopLocal2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterRouteMapRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenRouterRouteMapRuleAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "RouterRouteMapRule-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("fosid", flattenRouterRouteMapRuleId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "RouterRouteMapRule-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("match_as_path", flattenRouterRouteMapRuleMatchAsPath2edl(o["match-as-path"], d, "match_as_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-as-path"], "RouterRouteMapRule-MatchAsPath"); ok {
			if err = d.Set("match_as_path", vv); err != nil {
				return fmt.Errorf("Error reading match_as_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_as_path: %v", err)
		}
	}

	if err = d.Set("match_community", flattenRouterRouteMapRuleMatchCommunity2edl(o["match-community"], d, "match_community")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-community"], "RouterRouteMapRule-MatchCommunity"); ok {
			if err = d.Set("match_community", vv); err != nil {
				return fmt.Errorf("Error reading match_community: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_community: %v", err)
		}
	}

	if err = d.Set("match_community_exact", flattenRouterRouteMapRuleMatchCommunityExact2edl(o["match-community-exact"], d, "match_community_exact")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-community-exact"], "RouterRouteMapRule-MatchCommunityExact"); ok {
			if err = d.Set("match_community_exact", vv); err != nil {
				return fmt.Errorf("Error reading match_community_exact: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_community_exact: %v", err)
		}
	}

	if err = d.Set("match_extcommunity", flattenRouterRouteMapRuleMatchExtcommunity2edl(o["match-extcommunity"], d, "match_extcommunity")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-extcommunity"], "RouterRouteMapRule-MatchExtcommunity"); ok {
			if err = d.Set("match_extcommunity", vv); err != nil {
				return fmt.Errorf("Error reading match_extcommunity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_extcommunity: %v", err)
		}
	}

	if err = d.Set("match_extcommunity_exact", flattenRouterRouteMapRuleMatchExtcommunityExact2edl(o["match-extcommunity-exact"], d, "match_extcommunity_exact")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-extcommunity-exact"], "RouterRouteMapRule-MatchExtcommunityExact"); ok {
			if err = d.Set("match_extcommunity_exact", vv); err != nil {
				return fmt.Errorf("Error reading match_extcommunity_exact: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_extcommunity_exact: %v", err)
		}
	}

	if err = d.Set("match_flags", flattenRouterRouteMapRuleMatchFlags2edl(o["match-flags"], d, "match_flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-flags"], "RouterRouteMapRule-MatchFlags"); ok {
			if err = d.Set("match_flags", vv); err != nil {
				return fmt.Errorf("Error reading match_flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_flags: %v", err)
		}
	}

	if err = d.Set("match_interface", flattenRouterRouteMapRuleMatchInterface2edl(o["match-interface"], d, "match_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-interface"], "RouterRouteMapRule-MatchInterface"); ok {
			if err = d.Set("match_interface", vv); err != nil {
				return fmt.Errorf("Error reading match_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_interface: %v", err)
		}
	}

	if err = d.Set("match_ip_address", flattenRouterRouteMapRuleMatchIpAddress2edl(o["match-ip-address"], d, "match_ip_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-ip-address"], "RouterRouteMapRule-MatchIpAddress"); ok {
			if err = d.Set("match_ip_address", vv); err != nil {
				return fmt.Errorf("Error reading match_ip_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_ip_address: %v", err)
		}
	}

	if err = d.Set("match_ip_nexthop", flattenRouterRouteMapRuleMatchIpNexthop2edl(o["match-ip-nexthop"], d, "match_ip_nexthop")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-ip-nexthop"], "RouterRouteMapRule-MatchIpNexthop"); ok {
			if err = d.Set("match_ip_nexthop", vv); err != nil {
				return fmt.Errorf("Error reading match_ip_nexthop: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_ip_nexthop: %v", err)
		}
	}

	if err = d.Set("match_ip6_address", flattenRouterRouteMapRuleMatchIp6Address2edl(o["match-ip6-address"], d, "match_ip6_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-ip6-address"], "RouterRouteMapRule-MatchIp6Address"); ok {
			if err = d.Set("match_ip6_address", vv); err != nil {
				return fmt.Errorf("Error reading match_ip6_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_ip6_address: %v", err)
		}
	}

	if err = d.Set("match_ip6_nexthop", flattenRouterRouteMapRuleMatchIp6Nexthop2edl(o["match-ip6-nexthop"], d, "match_ip6_nexthop")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-ip6-nexthop"], "RouterRouteMapRule-MatchIp6Nexthop"); ok {
			if err = d.Set("match_ip6_nexthop", vv); err != nil {
				return fmt.Errorf("Error reading match_ip6_nexthop: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_ip6_nexthop: %v", err)
		}
	}

	if err = d.Set("match_metric", flattenRouterRouteMapRuleMatchMetric2edl(o["match-metric"], d, "match_metric")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-metric"], "RouterRouteMapRule-MatchMetric"); ok {
			if err = d.Set("match_metric", vv); err != nil {
				return fmt.Errorf("Error reading match_metric: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_metric: %v", err)
		}
	}

	if err = d.Set("match_origin", flattenRouterRouteMapRuleMatchOrigin2edl(o["match-origin"], d, "match_origin")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-origin"], "RouterRouteMapRule-MatchOrigin"); ok {
			if err = d.Set("match_origin", vv); err != nil {
				return fmt.Errorf("Error reading match_origin: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_origin: %v", err)
		}
	}

	if err = d.Set("match_route_type", flattenRouterRouteMapRuleMatchRouteType2edl(o["match-route-type"], d, "match_route_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-route-type"], "RouterRouteMapRule-MatchRouteType"); ok {
			if err = d.Set("match_route_type", vv); err != nil {
				return fmt.Errorf("Error reading match_route_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_route_type: %v", err)
		}
	}

	if err = d.Set("match_tag", flattenRouterRouteMapRuleMatchTag2edl(o["match-tag"], d, "match_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-tag"], "RouterRouteMapRule-MatchTag"); ok {
			if err = d.Set("match_tag", vv); err != nil {
				return fmt.Errorf("Error reading match_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_tag: %v", err)
		}
	}

	if err = d.Set("match_vrf", flattenRouterRouteMapRuleMatchVrf2edl(o["match-vrf"], d, "match_vrf")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-vrf"], "RouterRouteMapRule-MatchVrf"); ok {
			if err = d.Set("match_vrf", vv); err != nil {
				return fmt.Errorf("Error reading match_vrf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_vrf: %v", err)
		}
	}

	if err = d.Set("set_aggregator_as", flattenRouterRouteMapRuleSetAggregatorAs2edl(o["set-aggregator-as"], d, "set_aggregator_as")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-aggregator-as"], "RouterRouteMapRule-SetAggregatorAs"); ok {
			if err = d.Set("set_aggregator_as", vv); err != nil {
				return fmt.Errorf("Error reading set_aggregator_as: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_aggregator_as: %v", err)
		}
	}

	if err = d.Set("set_aggregator_ip", flattenRouterRouteMapRuleSetAggregatorIp2edl(o["set-aggregator-ip"], d, "set_aggregator_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-aggregator-ip"], "RouterRouteMapRule-SetAggregatorIp"); ok {
			if err = d.Set("set_aggregator_ip", vv); err != nil {
				return fmt.Errorf("Error reading set_aggregator_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_aggregator_ip: %v", err)
		}
	}

	if err = d.Set("set_aspath", flattenRouterRouteMapRuleSetAspath2edl(o["set-aspath"], d, "set_aspath")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-aspath"], "RouterRouteMapRule-SetAspath"); ok {
			if err = d.Set("set_aspath", vv); err != nil {
				return fmt.Errorf("Error reading set_aspath: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_aspath: %v", err)
		}
	}

	if err = d.Set("set_aspath_action", flattenRouterRouteMapRuleSetAspathAction2edl(o["set-aspath-action"], d, "set_aspath_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-aspath-action"], "RouterRouteMapRule-SetAspathAction"); ok {
			if err = d.Set("set_aspath_action", vv); err != nil {
				return fmt.Errorf("Error reading set_aspath_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_aspath_action: %v", err)
		}
	}

	if err = d.Set("set_atomic_aggregate", flattenRouterRouteMapRuleSetAtomicAggregate2edl(o["set-atomic-aggregate"], d, "set_atomic_aggregate")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-atomic-aggregate"], "RouterRouteMapRule-SetAtomicAggregate"); ok {
			if err = d.Set("set_atomic_aggregate", vv); err != nil {
				return fmt.Errorf("Error reading set_atomic_aggregate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_atomic_aggregate: %v", err)
		}
	}

	if err = d.Set("set_community", flattenRouterRouteMapRuleSetCommunity2edl(o["set-community"], d, "set_community")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-community"], "RouterRouteMapRule-SetCommunity"); ok {
			if err = d.Set("set_community", vv); err != nil {
				return fmt.Errorf("Error reading set_community: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_community: %v", err)
		}
	}

	if err = d.Set("set_community_additive", flattenRouterRouteMapRuleSetCommunityAdditive2edl(o["set-community-additive"], d, "set_community_additive")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-community-additive"], "RouterRouteMapRule-SetCommunityAdditive"); ok {
			if err = d.Set("set_community_additive", vv); err != nil {
				return fmt.Errorf("Error reading set_community_additive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_community_additive: %v", err)
		}
	}

	if err = d.Set("set_community_delete", flattenRouterRouteMapRuleSetCommunityDelete2edl(o["set-community-delete"], d, "set_community_delete")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-community-delete"], "RouterRouteMapRule-SetCommunityDelete"); ok {
			if err = d.Set("set_community_delete", vv); err != nil {
				return fmt.Errorf("Error reading set_community_delete: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_community_delete: %v", err)
		}
	}

	if err = d.Set("set_dampening_max_suppress", flattenRouterRouteMapRuleSetDampeningMaxSuppress2edl(o["set-dampening-max-suppress"], d, "set_dampening_max_suppress")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-dampening-max-suppress"], "RouterRouteMapRule-SetDampeningMaxSuppress"); ok {
			if err = d.Set("set_dampening_max_suppress", vv); err != nil {
				return fmt.Errorf("Error reading set_dampening_max_suppress: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_dampening_max_suppress: %v", err)
		}
	}

	if err = d.Set("set_dampening_reachability_half_life", flattenRouterRouteMapRuleSetDampeningReachabilityHalfLife2edl(o["set-dampening-reachability-half-life"], d, "set_dampening_reachability_half_life")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-dampening-reachability-half-life"], "RouterRouteMapRule-SetDampeningReachabilityHalfLife"); ok {
			if err = d.Set("set_dampening_reachability_half_life", vv); err != nil {
				return fmt.Errorf("Error reading set_dampening_reachability_half_life: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_dampening_reachability_half_life: %v", err)
		}
	}

	if err = d.Set("set_dampening_reuse", flattenRouterRouteMapRuleSetDampeningReuse2edl(o["set-dampening-reuse"], d, "set_dampening_reuse")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-dampening-reuse"], "RouterRouteMapRule-SetDampeningReuse"); ok {
			if err = d.Set("set_dampening_reuse", vv); err != nil {
				return fmt.Errorf("Error reading set_dampening_reuse: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_dampening_reuse: %v", err)
		}
	}

	if err = d.Set("set_dampening_suppress", flattenRouterRouteMapRuleSetDampeningSuppress2edl(o["set-dampening-suppress"], d, "set_dampening_suppress")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-dampening-suppress"], "RouterRouteMapRule-SetDampeningSuppress"); ok {
			if err = d.Set("set_dampening_suppress", vv); err != nil {
				return fmt.Errorf("Error reading set_dampening_suppress: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_dampening_suppress: %v", err)
		}
	}

	if err = d.Set("set_dampening_unreachability_half_life", flattenRouterRouteMapRuleSetDampeningUnreachabilityHalfLife2edl(o["set-dampening-unreachability-half-life"], d, "set_dampening_unreachability_half_life")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-dampening-unreachability-half-life"], "RouterRouteMapRule-SetDampeningUnreachabilityHalfLife"); ok {
			if err = d.Set("set_dampening_unreachability_half_life", vv); err != nil {
				return fmt.Errorf("Error reading set_dampening_unreachability_half_life: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_dampening_unreachability_half_life: %v", err)
		}
	}

	if err = d.Set("set_extcommunity_rt", flattenRouterRouteMapRuleSetExtcommunityRt2edl(o["set-extcommunity-rt"], d, "set_extcommunity_rt")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-extcommunity-rt"], "RouterRouteMapRule-SetExtcommunityRt"); ok {
			if err = d.Set("set_extcommunity_rt", vv); err != nil {
				return fmt.Errorf("Error reading set_extcommunity_rt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_extcommunity_rt: %v", err)
		}
	}

	if err = d.Set("set_extcommunity_soo", flattenRouterRouteMapRuleSetExtcommunitySoo2edl(o["set-extcommunity-soo"], d, "set_extcommunity_soo")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-extcommunity-soo"], "RouterRouteMapRule-SetExtcommunitySoo"); ok {
			if err = d.Set("set_extcommunity_soo", vv); err != nil {
				return fmt.Errorf("Error reading set_extcommunity_soo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_extcommunity_soo: %v", err)
		}
	}

	if err = d.Set("set_flags", flattenRouterRouteMapRuleSetFlags2edl(o["set-flags"], d, "set_flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-flags"], "RouterRouteMapRule-SetFlags"); ok {
			if err = d.Set("set_flags", vv); err != nil {
				return fmt.Errorf("Error reading set_flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_flags: %v", err)
		}
	}

	if err = d.Set("set_ip_nexthop", flattenRouterRouteMapRuleSetIpNexthop2edl(o["set-ip-nexthop"], d, "set_ip_nexthop")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-ip-nexthop"], "RouterRouteMapRule-SetIpNexthop"); ok {
			if err = d.Set("set_ip_nexthop", vv); err != nil {
				return fmt.Errorf("Error reading set_ip_nexthop: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_ip_nexthop: %v", err)
		}
	}

	if err = d.Set("set_ip_prefsrc", flattenRouterRouteMapRuleSetIpPrefsrc2edl(o["set-ip-prefsrc"], d, "set_ip_prefsrc")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-ip-prefsrc"], "RouterRouteMapRule-SetIpPrefsrc"); ok {
			if err = d.Set("set_ip_prefsrc", vv); err != nil {
				return fmt.Errorf("Error reading set_ip_prefsrc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_ip_prefsrc: %v", err)
		}
	}

	if err = d.Set("set_ip6_nexthop", flattenRouterRouteMapRuleSetIp6Nexthop2edl(o["set-ip6-nexthop"], d, "set_ip6_nexthop")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-ip6-nexthop"], "RouterRouteMapRule-SetIp6Nexthop"); ok {
			if err = d.Set("set_ip6_nexthop", vv); err != nil {
				return fmt.Errorf("Error reading set_ip6_nexthop: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_ip6_nexthop: %v", err)
		}
	}

	if err = d.Set("set_ip6_nexthop_local", flattenRouterRouteMapRuleSetIp6NexthopLocal2edl(o["set-ip6-nexthop-local"], d, "set_ip6_nexthop_local")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-ip6-nexthop-local"], "RouterRouteMapRule-SetIp6NexthopLocal"); ok {
			if err = d.Set("set_ip6_nexthop_local", vv); err != nil {
				return fmt.Errorf("Error reading set_ip6_nexthop_local: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_ip6_nexthop_local: %v", err)
		}
	}

	if err = d.Set("set_local_preference", flattenRouterRouteMapRuleSetLocalPreference2edl(o["set-local-preference"], d, "set_local_preference")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-local-preference"], "RouterRouteMapRule-SetLocalPreference"); ok {
			if err = d.Set("set_local_preference", vv); err != nil {
				return fmt.Errorf("Error reading set_local_preference: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_local_preference: %v", err)
		}
	}

	if err = d.Set("set_metric", flattenRouterRouteMapRuleSetMetric2edl(o["set-metric"], d, "set_metric")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-metric"], "RouterRouteMapRule-SetMetric"); ok {
			if err = d.Set("set_metric", vv); err != nil {
				return fmt.Errorf("Error reading set_metric: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_metric: %v", err)
		}
	}

	if err = d.Set("set_metric_type", flattenRouterRouteMapRuleSetMetricType2edl(o["set-metric-type"], d, "set_metric_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-metric-type"], "RouterRouteMapRule-SetMetricType"); ok {
			if err = d.Set("set_metric_type", vv); err != nil {
				return fmt.Errorf("Error reading set_metric_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_metric_type: %v", err)
		}
	}

	if err = d.Set("set_origin", flattenRouterRouteMapRuleSetOrigin2edl(o["set-origin"], d, "set_origin")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-origin"], "RouterRouteMapRule-SetOrigin"); ok {
			if err = d.Set("set_origin", vv); err != nil {
				return fmt.Errorf("Error reading set_origin: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_origin: %v", err)
		}
	}

	if err = d.Set("set_originator_id", flattenRouterRouteMapRuleSetOriginatorId2edl(o["set-originator-id"], d, "set_originator_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-originator-id"], "RouterRouteMapRule-SetOriginatorId"); ok {
			if err = d.Set("set_originator_id", vv); err != nil {
				return fmt.Errorf("Error reading set_originator_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_originator_id: %v", err)
		}
	}

	if err = d.Set("set_priority", flattenRouterRouteMapRuleSetPriority2edl(o["set-priority"], d, "set_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-priority"], "RouterRouteMapRule-SetPriority"); ok {
			if err = d.Set("set_priority", vv); err != nil {
				return fmt.Errorf("Error reading set_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_priority: %v", err)
		}
	}

	if err = d.Set("set_route_tag", flattenRouterRouteMapRuleSetRouteTag2edl(o["set-route-tag"], d, "set_route_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-route-tag"], "RouterRouteMapRule-SetRouteTag"); ok {
			if err = d.Set("set_route_tag", vv); err != nil {
				return fmt.Errorf("Error reading set_route_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_route_tag: %v", err)
		}
	}

	if err = d.Set("set_tag", flattenRouterRouteMapRuleSetTag2edl(o["set-tag"], d, "set_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-tag"], "RouterRouteMapRule-SetTag"); ok {
			if err = d.Set("set_tag", vv); err != nil {
				return fmt.Errorf("Error reading set_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_tag: %v", err)
		}
	}

	if err = d.Set("set_vpnv4_nexthop", flattenRouterRouteMapRuleSetVpnv4Nexthop2edl(o["set-vpnv4-nexthop"], d, "set_vpnv4_nexthop")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-vpnv4-nexthop"], "RouterRouteMapRule-SetVpnv4Nexthop"); ok {
			if err = d.Set("set_vpnv4_nexthop", vv); err != nil {
				return fmt.Errorf("Error reading set_vpnv4_nexthop: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_vpnv4_nexthop: %v", err)
		}
	}

	if err = d.Set("set_vpnv6_nexthop", flattenRouterRouteMapRuleSetVpnv6Nexthop2edl(o["set-vpnv6-nexthop"], d, "set_vpnv6_nexthop")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-vpnv6-nexthop"], "RouterRouteMapRule-SetVpnv6Nexthop"); ok {
			if err = d.Set("set_vpnv6_nexthop", vv); err != nil {
				return fmt.Errorf("Error reading set_vpnv6_nexthop: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_vpnv6_nexthop: %v", err)
		}
	}

	if err = d.Set("set_vpnv6_nexthop_local", flattenRouterRouteMapRuleSetVpnv6NexthopLocal2edl(o["set-vpnv6-nexthop-local"], d, "set_vpnv6_nexthop_local")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-vpnv6-nexthop-local"], "RouterRouteMapRule-SetVpnv6NexthopLocal"); ok {
			if err = d.Set("set_vpnv6_nexthop_local", vv); err != nil {
				return fmt.Errorf("Error reading set_vpnv6_nexthop_local: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_vpnv6_nexthop_local: %v", err)
		}
	}

	if err = d.Set("set_weight", flattenRouterRouteMapRuleSetWeight2edl(o["set-weight"], d, "set_weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-weight"], "RouterRouteMapRule-SetWeight"); ok {
			if err = d.Set("set_weight", vv); err != nil {
				return fmt.Errorf("Error reading set_weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_weight: %v", err)
		}
	}

	return nil
}

func flattenRouterRouteMapRuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterRouteMapRuleAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchAsPath2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRouteMapRuleMatchCommunity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRouteMapRuleMatchCommunityExact2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchExtcommunity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRouteMapRuleMatchExtcommunityExact2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchFlags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRouteMapRuleMatchIpAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRouteMapRuleMatchIpNexthop2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRouteMapRuleMatchIp6Address2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRouteMapRuleMatchIp6Nexthop2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRouteMapRuleMatchMetric2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchOrigin2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchRouteType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchTag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchVrf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetAggregatorAs2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetAggregatorIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetAspath2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRouteMapRuleSetAspathAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetAtomicAggregate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetCommunity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRouteMapRuleSetCommunityAdditive2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetCommunityDelete2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRouteMapRuleSetDampeningMaxSuppress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetDampeningReachabilityHalfLife2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetDampeningReuse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetDampeningSuppress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetDampeningUnreachabilityHalfLife2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetExtcommunityRt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRouteMapRuleSetExtcommunitySoo2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRouteMapRuleSetFlags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetIpNexthop2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetIpPrefsrc2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetIp6Nexthop2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetIp6NexthopLocal2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetLocalPreference2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetMetric2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetMetricType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetOrigin2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetOriginatorId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetRouteTag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetTag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetVpnv4Nexthop2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetVpnv6Nexthop2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetVpnv6NexthopLocal2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterRouteMapRule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandRouterRouteMapRuleAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandRouterRouteMapRuleId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("match_as_path"); ok || d.HasChange("match_as_path") {
		t, err := expandRouterRouteMapRuleMatchAsPath2edl(d, v, "match_as_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-as-path"] = t
		}
	}

	if v, ok := d.GetOk("match_community"); ok || d.HasChange("match_community") {
		t, err := expandRouterRouteMapRuleMatchCommunity2edl(d, v, "match_community")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-community"] = t
		}
	}

	if v, ok := d.GetOk("match_community_exact"); ok || d.HasChange("match_community_exact") {
		t, err := expandRouterRouteMapRuleMatchCommunityExact2edl(d, v, "match_community_exact")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-community-exact"] = t
		}
	}

	if v, ok := d.GetOk("match_extcommunity"); ok || d.HasChange("match_extcommunity") {
		t, err := expandRouterRouteMapRuleMatchExtcommunity2edl(d, v, "match_extcommunity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-extcommunity"] = t
		}
	}

	if v, ok := d.GetOk("match_extcommunity_exact"); ok || d.HasChange("match_extcommunity_exact") {
		t, err := expandRouterRouteMapRuleMatchExtcommunityExact2edl(d, v, "match_extcommunity_exact")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-extcommunity-exact"] = t
		}
	}

	if v, ok := d.GetOk("match_flags"); ok || d.HasChange("match_flags") {
		t, err := expandRouterRouteMapRuleMatchFlags2edl(d, v, "match_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-flags"] = t
		}
	}

	if v, ok := d.GetOk("match_interface"); ok || d.HasChange("match_interface") {
		t, err := expandRouterRouteMapRuleMatchInterface2edl(d, v, "match_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-interface"] = t
		}
	}

	if v, ok := d.GetOk("match_ip_address"); ok || d.HasChange("match_ip_address") {
		t, err := expandRouterRouteMapRuleMatchIpAddress2edl(d, v, "match_ip_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-ip-address"] = t
		}
	}

	if v, ok := d.GetOk("match_ip_nexthop"); ok || d.HasChange("match_ip_nexthop") {
		t, err := expandRouterRouteMapRuleMatchIpNexthop2edl(d, v, "match_ip_nexthop")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-ip-nexthop"] = t
		}
	}

	if v, ok := d.GetOk("match_ip6_address"); ok || d.HasChange("match_ip6_address") {
		t, err := expandRouterRouteMapRuleMatchIp6Address2edl(d, v, "match_ip6_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-ip6-address"] = t
		}
	}

	if v, ok := d.GetOk("match_ip6_nexthop"); ok || d.HasChange("match_ip6_nexthop") {
		t, err := expandRouterRouteMapRuleMatchIp6Nexthop2edl(d, v, "match_ip6_nexthop")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-ip6-nexthop"] = t
		}
	}

	if v, ok := d.GetOk("match_metric"); ok || d.HasChange("match_metric") {
		t, err := expandRouterRouteMapRuleMatchMetric2edl(d, v, "match_metric")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-metric"] = t
		}
	}

	if v, ok := d.GetOk("match_origin"); ok || d.HasChange("match_origin") {
		t, err := expandRouterRouteMapRuleMatchOrigin2edl(d, v, "match_origin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-origin"] = t
		}
	}

	if v, ok := d.GetOk("match_route_type"); ok || d.HasChange("match_route_type") {
		t, err := expandRouterRouteMapRuleMatchRouteType2edl(d, v, "match_route_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-route-type"] = t
		}
	}

	if v, ok := d.GetOk("match_tag"); ok || d.HasChange("match_tag") {
		t, err := expandRouterRouteMapRuleMatchTag2edl(d, v, "match_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-tag"] = t
		}
	}

	if v, ok := d.GetOk("match_vrf"); ok || d.HasChange("match_vrf") {
		t, err := expandRouterRouteMapRuleMatchVrf2edl(d, v, "match_vrf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-vrf"] = t
		}
	}

	if v, ok := d.GetOk("set_aggregator_as"); ok || d.HasChange("set_aggregator_as") {
		t, err := expandRouterRouteMapRuleSetAggregatorAs2edl(d, v, "set_aggregator_as")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-aggregator-as"] = t
		}
	}

	if v, ok := d.GetOk("set_aggregator_ip"); ok || d.HasChange("set_aggregator_ip") {
		t, err := expandRouterRouteMapRuleSetAggregatorIp2edl(d, v, "set_aggregator_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-aggregator-ip"] = t
		}
	}

	if v, ok := d.GetOk("set_aspath"); ok || d.HasChange("set_aspath") {
		t, err := expandRouterRouteMapRuleSetAspath2edl(d, v, "set_aspath")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-aspath"] = t
		}
	}

	if v, ok := d.GetOk("set_aspath_action"); ok || d.HasChange("set_aspath_action") {
		t, err := expandRouterRouteMapRuleSetAspathAction2edl(d, v, "set_aspath_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-aspath-action"] = t
		}
	}

	if v, ok := d.GetOk("set_atomic_aggregate"); ok || d.HasChange("set_atomic_aggregate") {
		t, err := expandRouterRouteMapRuleSetAtomicAggregate2edl(d, v, "set_atomic_aggregate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-atomic-aggregate"] = t
		}
	}

	if v, ok := d.GetOk("set_community"); ok || d.HasChange("set_community") {
		t, err := expandRouterRouteMapRuleSetCommunity2edl(d, v, "set_community")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-community"] = t
		}
	}

	if v, ok := d.GetOk("set_community_additive"); ok || d.HasChange("set_community_additive") {
		t, err := expandRouterRouteMapRuleSetCommunityAdditive2edl(d, v, "set_community_additive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-community-additive"] = t
		}
	}

	if v, ok := d.GetOk("set_community_delete"); ok || d.HasChange("set_community_delete") {
		t, err := expandRouterRouteMapRuleSetCommunityDelete2edl(d, v, "set_community_delete")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-community-delete"] = t
		}
	}

	if v, ok := d.GetOk("set_dampening_max_suppress"); ok || d.HasChange("set_dampening_max_suppress") {
		t, err := expandRouterRouteMapRuleSetDampeningMaxSuppress2edl(d, v, "set_dampening_max_suppress")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-dampening-max-suppress"] = t
		}
	}

	if v, ok := d.GetOk("set_dampening_reachability_half_life"); ok || d.HasChange("set_dampening_reachability_half_life") {
		t, err := expandRouterRouteMapRuleSetDampeningReachabilityHalfLife2edl(d, v, "set_dampening_reachability_half_life")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-dampening-reachability-half-life"] = t
		}
	}

	if v, ok := d.GetOk("set_dampening_reuse"); ok || d.HasChange("set_dampening_reuse") {
		t, err := expandRouterRouteMapRuleSetDampeningReuse2edl(d, v, "set_dampening_reuse")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-dampening-reuse"] = t
		}
	}

	if v, ok := d.GetOk("set_dampening_suppress"); ok || d.HasChange("set_dampening_suppress") {
		t, err := expandRouterRouteMapRuleSetDampeningSuppress2edl(d, v, "set_dampening_suppress")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-dampening-suppress"] = t
		}
	}

	if v, ok := d.GetOk("set_dampening_unreachability_half_life"); ok || d.HasChange("set_dampening_unreachability_half_life") {
		t, err := expandRouterRouteMapRuleSetDampeningUnreachabilityHalfLife2edl(d, v, "set_dampening_unreachability_half_life")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-dampening-unreachability-half-life"] = t
		}
	}

	if v, ok := d.GetOk("set_extcommunity_rt"); ok || d.HasChange("set_extcommunity_rt") {
		t, err := expandRouterRouteMapRuleSetExtcommunityRt2edl(d, v, "set_extcommunity_rt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-extcommunity-rt"] = t
		}
	}

	if v, ok := d.GetOk("set_extcommunity_soo"); ok || d.HasChange("set_extcommunity_soo") {
		t, err := expandRouterRouteMapRuleSetExtcommunitySoo2edl(d, v, "set_extcommunity_soo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-extcommunity-soo"] = t
		}
	}

	if v, ok := d.GetOk("set_flags"); ok || d.HasChange("set_flags") {
		t, err := expandRouterRouteMapRuleSetFlags2edl(d, v, "set_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-flags"] = t
		}
	}

	if v, ok := d.GetOk("set_ip_nexthop"); ok || d.HasChange("set_ip_nexthop") {
		t, err := expandRouterRouteMapRuleSetIpNexthop2edl(d, v, "set_ip_nexthop")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-ip-nexthop"] = t
		}
	}

	if v, ok := d.GetOk("set_ip_prefsrc"); ok || d.HasChange("set_ip_prefsrc") {
		t, err := expandRouterRouteMapRuleSetIpPrefsrc2edl(d, v, "set_ip_prefsrc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-ip-prefsrc"] = t
		}
	}

	if v, ok := d.GetOk("set_ip6_nexthop"); ok || d.HasChange("set_ip6_nexthop") {
		t, err := expandRouterRouteMapRuleSetIp6Nexthop2edl(d, v, "set_ip6_nexthop")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-ip6-nexthop"] = t
		}
	}

	if v, ok := d.GetOk("set_ip6_nexthop_local"); ok || d.HasChange("set_ip6_nexthop_local") {
		t, err := expandRouterRouteMapRuleSetIp6NexthopLocal2edl(d, v, "set_ip6_nexthop_local")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-ip6-nexthop-local"] = t
		}
	}

	if v, ok := d.GetOk("set_local_preference"); ok || d.HasChange("set_local_preference") {
		t, err := expandRouterRouteMapRuleSetLocalPreference2edl(d, v, "set_local_preference")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-local-preference"] = t
		}
	}

	if v, ok := d.GetOk("set_metric"); ok || d.HasChange("set_metric") {
		t, err := expandRouterRouteMapRuleSetMetric2edl(d, v, "set_metric")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-metric"] = t
		}
	}

	if v, ok := d.GetOk("set_metric_type"); ok || d.HasChange("set_metric_type") {
		t, err := expandRouterRouteMapRuleSetMetricType2edl(d, v, "set_metric_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-metric-type"] = t
		}
	}

	if v, ok := d.GetOk("set_origin"); ok || d.HasChange("set_origin") {
		t, err := expandRouterRouteMapRuleSetOrigin2edl(d, v, "set_origin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-origin"] = t
		}
	}

	if v, ok := d.GetOk("set_originator_id"); ok || d.HasChange("set_originator_id") {
		t, err := expandRouterRouteMapRuleSetOriginatorId2edl(d, v, "set_originator_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-originator-id"] = t
		}
	}

	if v, ok := d.GetOk("set_priority"); ok || d.HasChange("set_priority") {
		t, err := expandRouterRouteMapRuleSetPriority2edl(d, v, "set_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-priority"] = t
		}
	}

	if v, ok := d.GetOk("set_route_tag"); ok || d.HasChange("set_route_tag") {
		t, err := expandRouterRouteMapRuleSetRouteTag2edl(d, v, "set_route_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-route-tag"] = t
		}
	}

	if v, ok := d.GetOk("set_tag"); ok || d.HasChange("set_tag") {
		t, err := expandRouterRouteMapRuleSetTag2edl(d, v, "set_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-tag"] = t
		}
	}

	if v, ok := d.GetOk("set_vpnv4_nexthop"); ok || d.HasChange("set_vpnv4_nexthop") {
		t, err := expandRouterRouteMapRuleSetVpnv4Nexthop2edl(d, v, "set_vpnv4_nexthop")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-vpnv4-nexthop"] = t
		}
	}

	if v, ok := d.GetOk("set_vpnv6_nexthop"); ok || d.HasChange("set_vpnv6_nexthop") {
		t, err := expandRouterRouteMapRuleSetVpnv6Nexthop2edl(d, v, "set_vpnv6_nexthop")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-vpnv6-nexthop"] = t
		}
	}

	if v, ok := d.GetOk("set_vpnv6_nexthop_local"); ok || d.HasChange("set_vpnv6_nexthop_local") {
		t, err := expandRouterRouteMapRuleSetVpnv6NexthopLocal2edl(d, v, "set_vpnv6_nexthop_local")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-vpnv6-nexthop-local"] = t
		}
	}

	if v, ok := d.GetOk("set_weight"); ok || d.HasChange("set_weight") {
		t, err := expandRouterRouteMapRuleSetWeight2edl(d, v, "set_weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-weight"] = t
		}
	}

	return &obj, nil
}
