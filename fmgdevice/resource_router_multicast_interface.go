// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: PIM interfaces.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterMulticastInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterMulticastInterfaceCreate,
		Read:   resourceRouterMulticastInterfaceRead,
		Update: resourceRouterMulticastInterfaceUpdate,
		Delete: resourceRouterMulticastInterfaceDelete,

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
			"bfd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cisco_exclude_genid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dr_priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hello_holdtime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hello_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"igmp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_group": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"immediate_leave_group": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"last_member_query_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"last_member_query_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"query_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"query_max_response_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"query_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"router_alert_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"join_group": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"multicast_flow": &schema.Schema{
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
			"neighbour_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"passive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pim_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"propagation_delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rp_candidate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rp_candidate_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"rp_candidate_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rp_candidate_priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rpf_nbr_fail_back": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rpf_nbr_fail_back_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"state_refresh_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"static_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ttl_threshold": &schema.Schema{
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

func resourceRouterMulticastInterfaceCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterMulticastInterface(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterMulticastInterface resource while getting object: %v", err)
	}

	_, err = c.CreateRouterMulticastInterface(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating RouterMulticastInterface resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceRouterMulticastInterfaceRead(d, m)
}

func resourceRouterMulticastInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterMulticastInterface(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticastInterface resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterMulticastInterface(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticastInterface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceRouterMulticastInterfaceRead(d, m)
}

func resourceRouterMulticastInterfaceDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterMulticastInterface(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterMulticastInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterMulticastInterfaceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterMulticastInterface(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticastInterface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterMulticastInterface(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticastInterface resource from API: %v", err)
	}
	return nil
}

func flattenRouterMulticastInterfaceBfd2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceCiscoExcludeGenid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceDrPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceHelloHoldtime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceHelloInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceIgmp2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "access_group"
	if _, ok := i["access-group"]; ok {
		result["access_group"] = flattenRouterMulticastInterfaceIgmpAccessGroup2edl(i["access-group"], d, pre_append)
	}

	pre_append = pre + ".0." + "immediate_leave_group"
	if _, ok := i["immediate-leave-group"]; ok {
		result["immediate_leave_group"] = flattenRouterMulticastInterfaceIgmpImmediateLeaveGroup2edl(i["immediate-leave-group"], d, pre_append)
	}

	pre_append = pre + ".0." + "last_member_query_count"
	if _, ok := i["last-member-query-count"]; ok {
		result["last_member_query_count"] = flattenRouterMulticastInterfaceIgmpLastMemberQueryCount2edl(i["last-member-query-count"], d, pre_append)
	}

	pre_append = pre + ".0." + "last_member_query_interval"
	if _, ok := i["last-member-query-interval"]; ok {
		result["last_member_query_interval"] = flattenRouterMulticastInterfaceIgmpLastMemberQueryInterval2edl(i["last-member-query-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "query_interval"
	if _, ok := i["query-interval"]; ok {
		result["query_interval"] = flattenRouterMulticastInterfaceIgmpQueryInterval2edl(i["query-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "query_max_response_time"
	if _, ok := i["query-max-response-time"]; ok {
		result["query_max_response_time"] = flattenRouterMulticastInterfaceIgmpQueryMaxResponseTime2edl(i["query-max-response-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "query_timeout"
	if _, ok := i["query-timeout"]; ok {
		result["query_timeout"] = flattenRouterMulticastInterfaceIgmpQueryTimeout2edl(i["query-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "router_alert_check"
	if _, ok := i["router-alert-check"]; ok {
		result["router_alert_check"] = flattenRouterMulticastInterfaceIgmpRouterAlertCheck2edl(i["router-alert-check"], d, pre_append)
	}

	pre_append = pre + ".0." + "version"
	if _, ok := i["version"]; ok {
		result["version"] = flattenRouterMulticastInterfaceIgmpVersion2edl(i["version"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenRouterMulticastInterfaceIgmpAccessGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastInterfaceIgmpImmediateLeaveGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastInterfaceIgmpLastMemberQueryCount2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceIgmpLastMemberQueryInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceIgmpQueryInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceIgmpQueryMaxResponseTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceIgmpQueryTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceIgmpRouterAlertCheck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceIgmpVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceJoinGroup2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {
			v := flattenRouterMulticastInterfaceJoinGroupAddress2edl(i["address"], d, pre_append)
			tmp["address"] = fortiAPISubPartPatch(v, "RouterMulticastInterface-JoinGroup-Address")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterMulticastInterfaceJoinGroupAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceMulticastFlow2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastInterfaceName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenRouterMulticastInterfaceNeighbourFilter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastInterfacePassive2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfacePimMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfacePropagationDelay2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceRpCandidate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceRpCandidateGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastInterfaceRpCandidateInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceRpCandidatePriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceRpfNbrFailBack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceRpfNbrFailBackFilter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastInterfaceStateRefreshInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceStaticGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastInterfaceTtlThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterMulticastInterface(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("bfd", flattenRouterMulticastInterfaceBfd2edl(o["bfd"], d, "bfd")); err != nil {
		if vv, ok := fortiAPIPatch(o["bfd"], "RouterMulticastInterface-Bfd"); ok {
			if err = d.Set("bfd", vv); err != nil {
				return fmt.Errorf("Error reading bfd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	if err = d.Set("cisco_exclude_genid", flattenRouterMulticastInterfaceCiscoExcludeGenid2edl(o["cisco-exclude-genid"], d, "cisco_exclude_genid")); err != nil {
		if vv, ok := fortiAPIPatch(o["cisco-exclude-genid"], "RouterMulticastInterface-CiscoExcludeGenid"); ok {
			if err = d.Set("cisco_exclude_genid", vv); err != nil {
				return fmt.Errorf("Error reading cisco_exclude_genid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cisco_exclude_genid: %v", err)
		}
	}

	if err = d.Set("dr_priority", flattenRouterMulticastInterfaceDrPriority2edl(o["dr-priority"], d, "dr_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["dr-priority"], "RouterMulticastInterface-DrPriority"); ok {
			if err = d.Set("dr_priority", vv); err != nil {
				return fmt.Errorf("Error reading dr_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dr_priority: %v", err)
		}
	}

	if err = d.Set("hello_holdtime", flattenRouterMulticastInterfaceHelloHoldtime2edl(o["hello-holdtime"], d, "hello_holdtime")); err != nil {
		if vv, ok := fortiAPIPatch(o["hello-holdtime"], "RouterMulticastInterface-HelloHoldtime"); ok {
			if err = d.Set("hello_holdtime", vv); err != nil {
				return fmt.Errorf("Error reading hello_holdtime: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hello_holdtime: %v", err)
		}
	}

	if err = d.Set("hello_interval", flattenRouterMulticastInterfaceHelloInterval2edl(o["hello-interval"], d, "hello_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["hello-interval"], "RouterMulticastInterface-HelloInterval"); ok {
			if err = d.Set("hello_interval", vv); err != nil {
				return fmt.Errorf("Error reading hello_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hello_interval: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("igmp", flattenRouterMulticastInterfaceIgmp2edl(o["igmp"], d, "igmp")); err != nil {
			if vv, ok := fortiAPIPatch(o["igmp"], "RouterMulticastInterface-Igmp"); ok {
				if err = d.Set("igmp", vv); err != nil {
					return fmt.Errorf("Error reading igmp: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading igmp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("igmp"); ok {
			if err = d.Set("igmp", flattenRouterMulticastInterfaceIgmp2edl(o["igmp"], d, "igmp")); err != nil {
				if vv, ok := fortiAPIPatch(o["igmp"], "RouterMulticastInterface-Igmp"); ok {
					if err = d.Set("igmp", vv); err != nil {
						return fmt.Errorf("Error reading igmp: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading igmp: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("join_group", flattenRouterMulticastInterfaceJoinGroup2edl(o["join-group"], d, "join_group")); err != nil {
			if vv, ok := fortiAPIPatch(o["join-group"], "RouterMulticastInterface-JoinGroup"); ok {
				if err = d.Set("join_group", vv); err != nil {
					return fmt.Errorf("Error reading join_group: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading join_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("join_group"); ok {
			if err = d.Set("join_group", flattenRouterMulticastInterfaceJoinGroup2edl(o["join-group"], d, "join_group")); err != nil {
				if vv, ok := fortiAPIPatch(o["join-group"], "RouterMulticastInterface-JoinGroup"); ok {
					if err = d.Set("join_group", vv); err != nil {
						return fmt.Errorf("Error reading join_group: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading join_group: %v", err)
				}
			}
		}
	}

	if err = d.Set("multicast_flow", flattenRouterMulticastInterfaceMulticastFlow2edl(o["multicast-flow"], d, "multicast_flow")); err != nil {
		if vv, ok := fortiAPIPatch(o["multicast-flow"], "RouterMulticastInterface-MulticastFlow"); ok {
			if err = d.Set("multicast_flow", vv); err != nil {
				return fmt.Errorf("Error reading multicast_flow: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multicast_flow: %v", err)
		}
	}

	if err = d.Set("name", flattenRouterMulticastInterfaceName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "RouterMulticastInterface-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("neighbour_filter", flattenRouterMulticastInterfaceNeighbourFilter2edl(o["neighbour-filter"], d, "neighbour_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["neighbour-filter"], "RouterMulticastInterface-NeighbourFilter"); ok {
			if err = d.Set("neighbour_filter", vv); err != nil {
				return fmt.Errorf("Error reading neighbour_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading neighbour_filter: %v", err)
		}
	}

	if err = d.Set("passive", flattenRouterMulticastInterfacePassive2edl(o["passive"], d, "passive")); err != nil {
		if vv, ok := fortiAPIPatch(o["passive"], "RouterMulticastInterface-Passive"); ok {
			if err = d.Set("passive", vv); err != nil {
				return fmt.Errorf("Error reading passive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading passive: %v", err)
		}
	}

	if err = d.Set("pim_mode", flattenRouterMulticastInterfacePimMode2edl(o["pim-mode"], d, "pim_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["pim-mode"], "RouterMulticastInterface-PimMode"); ok {
			if err = d.Set("pim_mode", vv); err != nil {
				return fmt.Errorf("Error reading pim_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pim_mode: %v", err)
		}
	}

	if err = d.Set("propagation_delay", flattenRouterMulticastInterfacePropagationDelay2edl(o["propagation-delay"], d, "propagation_delay")); err != nil {
		if vv, ok := fortiAPIPatch(o["propagation-delay"], "RouterMulticastInterface-PropagationDelay"); ok {
			if err = d.Set("propagation_delay", vv); err != nil {
				return fmt.Errorf("Error reading propagation_delay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading propagation_delay: %v", err)
		}
	}

	if err = d.Set("rp_candidate", flattenRouterMulticastInterfaceRpCandidate2edl(o["rp-candidate"], d, "rp_candidate")); err != nil {
		if vv, ok := fortiAPIPatch(o["rp-candidate"], "RouterMulticastInterface-RpCandidate"); ok {
			if err = d.Set("rp_candidate", vv); err != nil {
				return fmt.Errorf("Error reading rp_candidate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rp_candidate: %v", err)
		}
	}

	if err = d.Set("rp_candidate_group", flattenRouterMulticastInterfaceRpCandidateGroup2edl(o["rp-candidate-group"], d, "rp_candidate_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["rp-candidate-group"], "RouterMulticastInterface-RpCandidateGroup"); ok {
			if err = d.Set("rp_candidate_group", vv); err != nil {
				return fmt.Errorf("Error reading rp_candidate_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rp_candidate_group: %v", err)
		}
	}

	if err = d.Set("rp_candidate_interval", flattenRouterMulticastInterfaceRpCandidateInterval2edl(o["rp-candidate-interval"], d, "rp_candidate_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["rp-candidate-interval"], "RouterMulticastInterface-RpCandidateInterval"); ok {
			if err = d.Set("rp_candidate_interval", vv); err != nil {
				return fmt.Errorf("Error reading rp_candidate_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rp_candidate_interval: %v", err)
		}
	}

	if err = d.Set("rp_candidate_priority", flattenRouterMulticastInterfaceRpCandidatePriority2edl(o["rp-candidate-priority"], d, "rp_candidate_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["rp-candidate-priority"], "RouterMulticastInterface-RpCandidatePriority"); ok {
			if err = d.Set("rp_candidate_priority", vv); err != nil {
				return fmt.Errorf("Error reading rp_candidate_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rp_candidate_priority: %v", err)
		}
	}

	if err = d.Set("rpf_nbr_fail_back", flattenRouterMulticastInterfaceRpfNbrFailBack2edl(o["rpf-nbr-fail-back"], d, "rpf_nbr_fail_back")); err != nil {
		if vv, ok := fortiAPIPatch(o["rpf-nbr-fail-back"], "RouterMulticastInterface-RpfNbrFailBack"); ok {
			if err = d.Set("rpf_nbr_fail_back", vv); err != nil {
				return fmt.Errorf("Error reading rpf_nbr_fail_back: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rpf_nbr_fail_back: %v", err)
		}
	}

	if err = d.Set("rpf_nbr_fail_back_filter", flattenRouterMulticastInterfaceRpfNbrFailBackFilter2edl(o["rpf-nbr-fail-back-filter"], d, "rpf_nbr_fail_back_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["rpf-nbr-fail-back-filter"], "RouterMulticastInterface-RpfNbrFailBackFilter"); ok {
			if err = d.Set("rpf_nbr_fail_back_filter", vv); err != nil {
				return fmt.Errorf("Error reading rpf_nbr_fail_back_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rpf_nbr_fail_back_filter: %v", err)
		}
	}

	if err = d.Set("state_refresh_interval", flattenRouterMulticastInterfaceStateRefreshInterval2edl(o["state-refresh-interval"], d, "state_refresh_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["state-refresh-interval"], "RouterMulticastInterface-StateRefreshInterval"); ok {
			if err = d.Set("state_refresh_interval", vv); err != nil {
				return fmt.Errorf("Error reading state_refresh_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading state_refresh_interval: %v", err)
		}
	}

	if err = d.Set("static_group", flattenRouterMulticastInterfaceStaticGroup2edl(o["static-group"], d, "static_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["static-group"], "RouterMulticastInterface-StaticGroup"); ok {
			if err = d.Set("static_group", vv); err != nil {
				return fmt.Errorf("Error reading static_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading static_group: %v", err)
		}
	}

	if err = d.Set("ttl_threshold", flattenRouterMulticastInterfaceTtlThreshold2edl(o["ttl-threshold"], d, "ttl_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["ttl-threshold"], "RouterMulticastInterface-TtlThreshold"); ok {
			if err = d.Set("ttl_threshold", vv); err != nil {
				return fmt.Errorf("Error reading ttl_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ttl_threshold: %v", err)
		}
	}

	return nil
}

func flattenRouterMulticastInterfaceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterMulticastInterfaceBfd2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceCiscoExcludeGenid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceDrPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceHelloHoldtime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceHelloInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "access_group"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["access-group"], _ = expandRouterMulticastInterfaceIgmpAccessGroup2edl(d, i["access_group"], pre_append)
	}
	pre_append = pre + ".0." + "immediate_leave_group"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["immediate-leave-group"], _ = expandRouterMulticastInterfaceIgmpImmediateLeaveGroup2edl(d, i["immediate_leave_group"], pre_append)
	}
	pre_append = pre + ".0." + "last_member_query_count"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["last-member-query-count"], _ = expandRouterMulticastInterfaceIgmpLastMemberQueryCount2edl(d, i["last_member_query_count"], pre_append)
	}
	pre_append = pre + ".0." + "last_member_query_interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["last-member-query-interval"], _ = expandRouterMulticastInterfaceIgmpLastMemberQueryInterval2edl(d, i["last_member_query_interval"], pre_append)
	}
	pre_append = pre + ".0." + "query_interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["query-interval"], _ = expandRouterMulticastInterfaceIgmpQueryInterval2edl(d, i["query_interval"], pre_append)
	}
	pre_append = pre + ".0." + "query_max_response_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["query-max-response-time"], _ = expandRouterMulticastInterfaceIgmpQueryMaxResponseTime2edl(d, i["query_max_response_time"], pre_append)
	}
	pre_append = pre + ".0." + "query_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["query-timeout"], _ = expandRouterMulticastInterfaceIgmpQueryTimeout2edl(d, i["query_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "router_alert_check"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["router-alert-check"], _ = expandRouterMulticastInterfaceIgmpRouterAlertCheck2edl(d, i["router_alert_check"], pre_append)
	}
	pre_append = pre + ".0." + "version"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["version"], _ = expandRouterMulticastInterfaceIgmpVersion2edl(d, i["version"], pre_append)
	}

	return result, nil
}

func expandRouterMulticastInterfaceIgmpAccessGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastInterfaceIgmpImmediateLeaveGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastInterfaceIgmpLastMemberQueryCount2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmpLastMemberQueryInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmpQueryInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmpQueryMaxResponseTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmpQueryTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmpRouterAlertCheck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmpVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceJoinGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["address"], _ = expandRouterMulticastInterfaceJoinGroupAddress2edl(d, i["address"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterMulticastInterfaceJoinGroupAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceMulticastFlow2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastInterfaceName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandRouterMulticastInterfaceNeighbourFilter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastInterfacePassive2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfacePimMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfacePropagationDelay2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceRpCandidate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceRpCandidateGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastInterfaceRpCandidateInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceRpCandidatePriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceRpfNbrFailBack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceRpfNbrFailBackFilter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastInterfaceStateRefreshInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceStaticGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastInterfaceTtlThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterMulticastInterface(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bfd"); ok || d.HasChange("bfd") {
		t, err := expandRouterMulticastInterfaceBfd2edl(d, v, "bfd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd"] = t
		}
	}

	if v, ok := d.GetOk("cisco_exclude_genid"); ok || d.HasChange("cisco_exclude_genid") {
		t, err := expandRouterMulticastInterfaceCiscoExcludeGenid2edl(d, v, "cisco_exclude_genid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cisco-exclude-genid"] = t
		}
	}

	if v, ok := d.GetOk("dr_priority"); ok || d.HasChange("dr_priority") {
		t, err := expandRouterMulticastInterfaceDrPriority2edl(d, v, "dr_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dr-priority"] = t
		}
	}

	if v, ok := d.GetOk("hello_holdtime"); ok || d.HasChange("hello_holdtime") {
		t, err := expandRouterMulticastInterfaceHelloHoldtime2edl(d, v, "hello_holdtime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hello-holdtime"] = t
		}
	}

	if v, ok := d.GetOk("hello_interval"); ok || d.HasChange("hello_interval") {
		t, err := expandRouterMulticastInterfaceHelloInterval2edl(d, v, "hello_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hello-interval"] = t
		}
	}

	if v, ok := d.GetOk("igmp"); ok || d.HasChange("igmp") {
		t, err := expandRouterMulticastInterfaceIgmp2edl(d, v, "igmp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["igmp"] = t
		}
	}

	if v, ok := d.GetOk("join_group"); ok || d.HasChange("join_group") {
		t, err := expandRouterMulticastInterfaceJoinGroup2edl(d, v, "join_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["join-group"] = t
		}
	}

	if v, ok := d.GetOk("multicast_flow"); ok || d.HasChange("multicast_flow") {
		t, err := expandRouterMulticastInterfaceMulticastFlow2edl(d, v, "multicast_flow")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-flow"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandRouterMulticastInterfaceName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("neighbour_filter"); ok || d.HasChange("neighbour_filter") {
		t, err := expandRouterMulticastInterfaceNeighbourFilter2edl(d, v, "neighbour_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbour-filter"] = t
		}
	}

	if v, ok := d.GetOk("passive"); ok || d.HasChange("passive") {
		t, err := expandRouterMulticastInterfacePassive2edl(d, v, "passive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passive"] = t
		}
	}

	if v, ok := d.GetOk("pim_mode"); ok || d.HasChange("pim_mode") {
		t, err := expandRouterMulticastInterfacePimMode2edl(d, v, "pim_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pim-mode"] = t
		}
	}

	if v, ok := d.GetOk("propagation_delay"); ok || d.HasChange("propagation_delay") {
		t, err := expandRouterMulticastInterfacePropagationDelay2edl(d, v, "propagation_delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["propagation-delay"] = t
		}
	}

	if v, ok := d.GetOk("rp_candidate"); ok || d.HasChange("rp_candidate") {
		t, err := expandRouterMulticastInterfaceRpCandidate2edl(d, v, "rp_candidate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rp-candidate"] = t
		}
	}

	if v, ok := d.GetOk("rp_candidate_group"); ok || d.HasChange("rp_candidate_group") {
		t, err := expandRouterMulticastInterfaceRpCandidateGroup2edl(d, v, "rp_candidate_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rp-candidate-group"] = t
		}
	}

	if v, ok := d.GetOk("rp_candidate_interval"); ok || d.HasChange("rp_candidate_interval") {
		t, err := expandRouterMulticastInterfaceRpCandidateInterval2edl(d, v, "rp_candidate_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rp-candidate-interval"] = t
		}
	}

	if v, ok := d.GetOk("rp_candidate_priority"); ok || d.HasChange("rp_candidate_priority") {
		t, err := expandRouterMulticastInterfaceRpCandidatePriority2edl(d, v, "rp_candidate_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rp-candidate-priority"] = t
		}
	}

	if v, ok := d.GetOk("rpf_nbr_fail_back"); ok || d.HasChange("rpf_nbr_fail_back") {
		t, err := expandRouterMulticastInterfaceRpfNbrFailBack2edl(d, v, "rpf_nbr_fail_back")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rpf-nbr-fail-back"] = t
		}
	}

	if v, ok := d.GetOk("rpf_nbr_fail_back_filter"); ok || d.HasChange("rpf_nbr_fail_back_filter") {
		t, err := expandRouterMulticastInterfaceRpfNbrFailBackFilter2edl(d, v, "rpf_nbr_fail_back_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rpf-nbr-fail-back-filter"] = t
		}
	}

	if v, ok := d.GetOk("state_refresh_interval"); ok || d.HasChange("state_refresh_interval") {
		t, err := expandRouterMulticastInterfaceStateRefreshInterval2edl(d, v, "state_refresh_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["state-refresh-interval"] = t
		}
	}

	if v, ok := d.GetOk("static_group"); ok || d.HasChange("static_group") {
		t, err := expandRouterMulticastInterfaceStaticGroup2edl(d, v, "static_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["static-group"] = t
		}
	}

	if v, ok := d.GetOk("ttl_threshold"); ok || d.HasChange("ttl_threshold") {
		t, err := expandRouterMulticastInterfaceTtlThreshold2edl(d, v, "ttl_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ttl-threshold"] = t
		}
	}

	return &obj, nil
}
