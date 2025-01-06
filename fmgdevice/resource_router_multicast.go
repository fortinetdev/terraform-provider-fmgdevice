// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure router multicast.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterMulticast() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterMulticastUpdate,
		Read:   resourceRouterMulticastRead,
		Update: resourceRouterMulticastUpdate,
		Delete: resourceRouterMulticastDelete,

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
			"interface": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
					},
				},
			},
			"multicast_routing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pim_sm_global": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accept_register_list": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"accept_source_list": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"bsr_allow_quick_refresh": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bsr_candidate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bsr_hash": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"bsr_interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"bsr_priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"cisco_crp_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cisco_ignore_rp_set_priority": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cisco_register_checksum": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cisco_register_checksum_group": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"join_prune_holdtime": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"message_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"null_register_retries": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"pim_use_sdwan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"register_rate_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"register_rp_reachability": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"register_source": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"register_source_interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"register_source_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"register_supression": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"rp_address": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"group": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ip_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"rp_register_keepalive": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"spt_threshold": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"spt_threshold_group": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"ssm": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssm_range": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"route_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"route_threshold": &schema.Schema{
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

func resourceRouterMulticastUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterMulticast(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticast resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterMulticast(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticast resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("RouterMulticast")

	return resourceRouterMulticastRead(d, m)
}

func resourceRouterMulticastDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterMulticast(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting RouterMulticast resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterMulticastRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterMulticast(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticast resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterMulticast(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticast resource from API: %v", err)
	}
	return nil
}

func flattenRouterMulticastInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := i["bfd"]; ok {
			v := flattenRouterMulticastInterfaceBfd(i["bfd"], d, pre_append)
			tmp["bfd"] = fortiAPISubPartPatch(v, "RouterMulticast-Interface-Bfd")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cisco_exclude_genid"
		if _, ok := i["cisco-exclude-genid"]; ok {
			v := flattenRouterMulticastInterfaceCiscoExcludeGenid(i["cisco-exclude-genid"], d, pre_append)
			tmp["cisco_exclude_genid"] = fortiAPISubPartPatch(v, "RouterMulticast-Interface-CiscoExcludeGenid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dr_priority"
		if _, ok := i["dr-priority"]; ok {
			v := flattenRouterMulticastInterfaceDrPriority(i["dr-priority"], d, pre_append)
			tmp["dr_priority"] = fortiAPISubPartPatch(v, "RouterMulticast-Interface-DrPriority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_holdtime"
		if _, ok := i["hello-holdtime"]; ok {
			v := flattenRouterMulticastInterfaceHelloHoldtime(i["hello-holdtime"], d, pre_append)
			tmp["hello_holdtime"] = fortiAPISubPartPatch(v, "RouterMulticast-Interface-HelloHoldtime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := i["hello-interval"]; ok {
			v := flattenRouterMulticastInterfaceHelloInterval(i["hello-interval"], d, pre_append)
			tmp["hello_interval"] = fortiAPISubPartPatch(v, "RouterMulticast-Interface-HelloInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmp"
		if _, ok := i["igmp"]; ok {
			v := flattenRouterMulticastInterfaceIgmp(i["igmp"], d, pre_append)
			tmp["igmp"] = fortiAPISubPartPatch(v, "RouterMulticast-Interface-Igmp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "join_group"
		if _, ok := i["join-group"]; ok {
			v := flattenRouterMulticastInterfaceJoinGroup(i["join-group"], d, pre_append)
			tmp["join_group"] = fortiAPISubPartPatch(v, "RouterMulticast-Interface-JoinGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "multicast_flow"
		if _, ok := i["multicast-flow"]; ok {
			v := flattenRouterMulticastInterfaceMulticastFlow(i["multicast-flow"], d, pre_append)
			tmp["multicast_flow"] = fortiAPISubPartPatch(v, "RouterMulticast-Interface-MulticastFlow")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenRouterMulticastInterfaceName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "RouterMulticast-Interface-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbour_filter"
		if _, ok := i["neighbour-filter"]; ok {
			v := flattenRouterMulticastInterfaceNeighbourFilter(i["neighbour-filter"], d, pre_append)
			tmp["neighbour_filter"] = fortiAPISubPartPatch(v, "RouterMulticast-Interface-NeighbourFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive"
		if _, ok := i["passive"]; ok {
			v := flattenRouterMulticastInterfacePassive(i["passive"], d, pre_append)
			tmp["passive"] = fortiAPISubPartPatch(v, "RouterMulticast-Interface-Passive")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pim_mode"
		if _, ok := i["pim-mode"]; ok {
			v := flattenRouterMulticastInterfacePimMode(i["pim-mode"], d, pre_append)
			tmp["pim_mode"] = fortiAPISubPartPatch(v, "RouterMulticast-Interface-PimMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "propagation_delay"
		if _, ok := i["propagation-delay"]; ok {
			v := flattenRouterMulticastInterfacePropagationDelay(i["propagation-delay"], d, pre_append)
			tmp["propagation_delay"] = fortiAPISubPartPatch(v, "RouterMulticast-Interface-PropagationDelay")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate"
		if _, ok := i["rp-candidate"]; ok {
			v := flattenRouterMulticastInterfaceRpCandidate(i["rp-candidate"], d, pre_append)
			tmp["rp_candidate"] = fortiAPISubPartPatch(v, "RouterMulticast-Interface-RpCandidate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate_group"
		if _, ok := i["rp-candidate-group"]; ok {
			v := flattenRouterMulticastInterfaceRpCandidateGroup(i["rp-candidate-group"], d, pre_append)
			tmp["rp_candidate_group"] = fortiAPISubPartPatch(v, "RouterMulticast-Interface-RpCandidateGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate_interval"
		if _, ok := i["rp-candidate-interval"]; ok {
			v := flattenRouterMulticastInterfaceRpCandidateInterval(i["rp-candidate-interval"], d, pre_append)
			tmp["rp_candidate_interval"] = fortiAPISubPartPatch(v, "RouterMulticast-Interface-RpCandidateInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate_priority"
		if _, ok := i["rp-candidate-priority"]; ok {
			v := flattenRouterMulticastInterfaceRpCandidatePriority(i["rp-candidate-priority"], d, pre_append)
			tmp["rp_candidate_priority"] = fortiAPISubPartPatch(v, "RouterMulticast-Interface-RpCandidatePriority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rpf_nbr_fail_back"
		if _, ok := i["rpf-nbr-fail-back"]; ok {
			v := flattenRouterMulticastInterfaceRpfNbrFailBack(i["rpf-nbr-fail-back"], d, pre_append)
			tmp["rpf_nbr_fail_back"] = fortiAPISubPartPatch(v, "RouterMulticast-Interface-RpfNbrFailBack")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rpf_nbr_fail_back_filter"
		if _, ok := i["rpf-nbr-fail-back-filter"]; ok {
			v := flattenRouterMulticastInterfaceRpfNbrFailBackFilter(i["rpf-nbr-fail-back-filter"], d, pre_append)
			tmp["rpf_nbr_fail_back_filter"] = fortiAPISubPartPatch(v, "RouterMulticast-Interface-RpfNbrFailBackFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "state_refresh_interval"
		if _, ok := i["state-refresh-interval"]; ok {
			v := flattenRouterMulticastInterfaceStateRefreshInterval(i["state-refresh-interval"], d, pre_append)
			tmp["state_refresh_interval"] = fortiAPISubPartPatch(v, "RouterMulticast-Interface-StateRefreshInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "static_group"
		if _, ok := i["static-group"]; ok {
			v := flattenRouterMulticastInterfaceStaticGroup(i["static-group"], d, pre_append)
			tmp["static_group"] = fortiAPISubPartPatch(v, "RouterMulticast-Interface-StaticGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ttl_threshold"
		if _, ok := i["ttl-threshold"]; ok {
			v := flattenRouterMulticastInterfaceTtlThreshold(i["ttl-threshold"], d, pre_append)
			tmp["ttl_threshold"] = fortiAPISubPartPatch(v, "RouterMulticast-Interface-TtlThreshold")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterMulticastInterfaceBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceCiscoExcludeGenid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceDrPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceHelloHoldtime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceHelloInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceIgmp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "access_group"
	if _, ok := i["access-group"]; ok {
		result["access_group"] = flattenRouterMulticastInterfaceIgmpAccessGroup(i["access-group"], d, pre_append)
	}

	pre_append = pre + ".0." + "immediate_leave_group"
	if _, ok := i["immediate-leave-group"]; ok {
		result["immediate_leave_group"] = flattenRouterMulticastInterfaceIgmpImmediateLeaveGroup(i["immediate-leave-group"], d, pre_append)
	}

	pre_append = pre + ".0." + "last_member_query_count"
	if _, ok := i["last-member-query-count"]; ok {
		result["last_member_query_count"] = flattenRouterMulticastInterfaceIgmpLastMemberQueryCount(i["last-member-query-count"], d, pre_append)
	}

	pre_append = pre + ".0." + "last_member_query_interval"
	if _, ok := i["last-member-query-interval"]; ok {
		result["last_member_query_interval"] = flattenRouterMulticastInterfaceIgmpLastMemberQueryInterval(i["last-member-query-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "query_interval"
	if _, ok := i["query-interval"]; ok {
		result["query_interval"] = flattenRouterMulticastInterfaceIgmpQueryInterval(i["query-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "query_max_response_time"
	if _, ok := i["query-max-response-time"]; ok {
		result["query_max_response_time"] = flattenRouterMulticastInterfaceIgmpQueryMaxResponseTime(i["query-max-response-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "query_timeout"
	if _, ok := i["query-timeout"]; ok {
		result["query_timeout"] = flattenRouterMulticastInterfaceIgmpQueryTimeout(i["query-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "router_alert_check"
	if _, ok := i["router-alert-check"]; ok {
		result["router_alert_check"] = flattenRouterMulticastInterfaceIgmpRouterAlertCheck(i["router-alert-check"], d, pre_append)
	}

	pre_append = pre + ".0." + "version"
	if _, ok := i["version"]; ok {
		result["version"] = flattenRouterMulticastInterfaceIgmpVersion(i["version"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenRouterMulticastInterfaceIgmpAccessGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastInterfaceIgmpImmediateLeaveGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastInterfaceIgmpLastMemberQueryCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceIgmpLastMemberQueryInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceIgmpQueryInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceIgmpQueryMaxResponseTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceIgmpQueryTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceIgmpRouterAlertCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceIgmpVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceJoinGroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterMulticastInterfaceJoinGroupAddress(i["address"], d, pre_append)
			tmp["address"] = fortiAPISubPartPatch(v, "RouterMulticastInterface-JoinGroup-Address")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterMulticastInterfaceJoinGroupAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceMulticastFlow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenRouterMulticastInterfaceNeighbourFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastInterfacePassive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfacePimMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfacePropagationDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceRpCandidate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceRpCandidateGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastInterfaceRpCandidateInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceRpCandidatePriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceRpfNbrFailBack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceRpfNbrFailBackFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastInterfaceStateRefreshInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceStaticGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastInterfaceTtlThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastMulticastRouting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobal(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "accept_register_list"
	if _, ok := i["accept-register-list"]; ok {
		result["accept_register_list"] = flattenRouterMulticastPimSmGlobalAcceptRegisterList(i["accept-register-list"], d, pre_append)
	}

	pre_append = pre + ".0." + "accept_source_list"
	if _, ok := i["accept-source-list"]; ok {
		result["accept_source_list"] = flattenRouterMulticastPimSmGlobalAcceptSourceList(i["accept-source-list"], d, pre_append)
	}

	pre_append = pre + ".0." + "bsr_allow_quick_refresh"
	if _, ok := i["bsr-allow-quick-refresh"]; ok {
		result["bsr_allow_quick_refresh"] = flattenRouterMulticastPimSmGlobalBsrAllowQuickRefresh(i["bsr-allow-quick-refresh"], d, pre_append)
	}

	pre_append = pre + ".0." + "bsr_candidate"
	if _, ok := i["bsr-candidate"]; ok {
		result["bsr_candidate"] = flattenRouterMulticastPimSmGlobalBsrCandidate(i["bsr-candidate"], d, pre_append)
	}

	pre_append = pre + ".0." + "bsr_hash"
	if _, ok := i["bsr-hash"]; ok {
		result["bsr_hash"] = flattenRouterMulticastPimSmGlobalBsrHash(i["bsr-hash"], d, pre_append)
	}

	pre_append = pre + ".0." + "bsr_interface"
	if _, ok := i["bsr-interface"]; ok {
		result["bsr_interface"] = flattenRouterMulticastPimSmGlobalBsrInterface(i["bsr-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "bsr_priority"
	if _, ok := i["bsr-priority"]; ok {
		result["bsr_priority"] = flattenRouterMulticastPimSmGlobalBsrPriority(i["bsr-priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "cisco_crp_prefix"
	if _, ok := i["cisco-crp-prefix"]; ok {
		result["cisco_crp_prefix"] = flattenRouterMulticastPimSmGlobalCiscoCrpPrefix(i["cisco-crp-prefix"], d, pre_append)
	}

	pre_append = pre + ".0." + "cisco_ignore_rp_set_priority"
	if _, ok := i["cisco-ignore-rp-set-priority"]; ok {
		result["cisco_ignore_rp_set_priority"] = flattenRouterMulticastPimSmGlobalCiscoIgnoreRpSetPriority(i["cisco-ignore-rp-set-priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "cisco_register_checksum"
	if _, ok := i["cisco-register-checksum"]; ok {
		result["cisco_register_checksum"] = flattenRouterMulticastPimSmGlobalCiscoRegisterChecksum(i["cisco-register-checksum"], d, pre_append)
	}

	pre_append = pre + ".0." + "cisco_register_checksum_group"
	if _, ok := i["cisco-register-checksum-group"]; ok {
		result["cisco_register_checksum_group"] = flattenRouterMulticastPimSmGlobalCiscoRegisterChecksumGroup(i["cisco-register-checksum-group"], d, pre_append)
	}

	pre_append = pre + ".0." + "join_prune_holdtime"
	if _, ok := i["join-prune-holdtime"]; ok {
		result["join_prune_holdtime"] = flattenRouterMulticastPimSmGlobalJoinPruneHoldtime(i["join-prune-holdtime"], d, pre_append)
	}

	pre_append = pre + ".0." + "message_interval"
	if _, ok := i["message-interval"]; ok {
		result["message_interval"] = flattenRouterMulticastPimSmGlobalMessageInterval(i["message-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "null_register_retries"
	if _, ok := i["null-register-retries"]; ok {
		result["null_register_retries"] = flattenRouterMulticastPimSmGlobalNullRegisterRetries(i["null-register-retries"], d, pre_append)
	}

	pre_append = pre + ".0." + "pim_use_sdwan"
	if _, ok := i["pim-use-sdwan"]; ok {
		result["pim_use_sdwan"] = flattenRouterMulticastPimSmGlobalPimUseSdwan(i["pim-use-sdwan"], d, pre_append)
	}

	pre_append = pre + ".0." + "register_rate_limit"
	if _, ok := i["register-rate-limit"]; ok {
		result["register_rate_limit"] = flattenRouterMulticastPimSmGlobalRegisterRateLimit(i["register-rate-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "register_rp_reachability"
	if _, ok := i["register-rp-reachability"]; ok {
		result["register_rp_reachability"] = flattenRouterMulticastPimSmGlobalRegisterRpReachability(i["register-rp-reachability"], d, pre_append)
	}

	pre_append = pre + ".0." + "register_source"
	if _, ok := i["register-source"]; ok {
		result["register_source"] = flattenRouterMulticastPimSmGlobalRegisterSource(i["register-source"], d, pre_append)
	}

	pre_append = pre + ".0." + "register_source_interface"
	if _, ok := i["register-source-interface"]; ok {
		result["register_source_interface"] = flattenRouterMulticastPimSmGlobalRegisterSourceInterface(i["register-source-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "register_source_ip"
	if _, ok := i["register-source-ip"]; ok {
		result["register_source_ip"] = flattenRouterMulticastPimSmGlobalRegisterSourceIp(i["register-source-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "register_supression"
	if _, ok := i["register-supression"]; ok {
		result["register_supression"] = flattenRouterMulticastPimSmGlobalRegisterSupression(i["register-supression"], d, pre_append)
	}

	pre_append = pre + ".0." + "rp_address"
	if _, ok := i["rp-address"]; ok {
		result["rp_address"] = flattenRouterMulticastPimSmGlobalRpAddress(i["rp-address"], d, pre_append)
	}

	pre_append = pre + ".0." + "rp_register_keepalive"
	if _, ok := i["rp-register-keepalive"]; ok {
		result["rp_register_keepalive"] = flattenRouterMulticastPimSmGlobalRpRegisterKeepalive(i["rp-register-keepalive"], d, pre_append)
	}

	pre_append = pre + ".0." + "spt_threshold"
	if _, ok := i["spt-threshold"]; ok {
		result["spt_threshold"] = flattenRouterMulticastPimSmGlobalSptThreshold(i["spt-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "spt_threshold_group"
	if _, ok := i["spt-threshold-group"]; ok {
		result["spt_threshold_group"] = flattenRouterMulticastPimSmGlobalSptThresholdGroup(i["spt-threshold-group"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssm"
	if _, ok := i["ssm"]; ok {
		result["ssm"] = flattenRouterMulticastPimSmGlobalSsm(i["ssm"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssm_range"
	if _, ok := i["ssm-range"]; ok {
		result["ssm_range"] = flattenRouterMulticastPimSmGlobalSsmRange(i["ssm-range"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenRouterMulticastPimSmGlobalAcceptRegisterList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastPimSmGlobalAcceptSourceList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastPimSmGlobalBsrAllowQuickRefresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalBsrCandidate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalBsrHash(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalBsrInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastPimSmGlobalBsrPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalCiscoCrpPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalCiscoIgnoreRpSetPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalCiscoRegisterChecksum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalCiscoRegisterChecksumGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastPimSmGlobalJoinPruneHoldtime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalMessageInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalNullRegisterRetries(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalPimUseSdwan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalRegisterRateLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalRegisterRpReachability(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalRegisterSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalRegisterSourceInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastPimSmGlobalRegisterSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalRegisterSupression(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalRpAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
		if _, ok := i["group"]; ok {
			v := flattenRouterMulticastPimSmGlobalRpAddressGroup(i["group"], d, pre_append)
			tmp["group"] = fortiAPISubPartPatch(v, "RouterMulticastPimSmGlobal-RpAddress-Group")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterMulticastPimSmGlobalRpAddressId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterMulticastPimSmGlobal-RpAddress-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_address"
		if _, ok := i["ip-address"]; ok {
			v := flattenRouterMulticastPimSmGlobalRpAddressIpAddress(i["ip-address"], d, pre_append)
			tmp["ip_address"] = fortiAPISubPartPatch(v, "RouterMulticastPimSmGlobal-RpAddress-IpAddress")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterMulticastPimSmGlobalRpAddressGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastPimSmGlobalRpAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalRpAddressIpAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalRpRegisterKeepalive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalSptThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalSptThresholdGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastPimSmGlobalSsm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalSsmRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastRouteLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastRouteThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterMulticast(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("interface", flattenRouterMulticastInterface(o["interface"], d, "interface")); err != nil {
			if vv, ok := fortiAPIPatch(o["interface"], "RouterMulticast-Interface"); ok {
				if err = d.Set("interface", vv); err != nil {
					return fmt.Errorf("Error reading interface: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("interface"); ok {
			if err = d.Set("interface", flattenRouterMulticastInterface(o["interface"], d, "interface")); err != nil {
				if vv, ok := fortiAPIPatch(o["interface"], "RouterMulticast-Interface"); ok {
					if err = d.Set("interface", vv); err != nil {
						return fmt.Errorf("Error reading interface: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading interface: %v", err)
				}
			}
		}
	}

	if err = d.Set("multicast_routing", flattenRouterMulticastMulticastRouting(o["multicast-routing"], d, "multicast_routing")); err != nil {
		if vv, ok := fortiAPIPatch(o["multicast-routing"], "RouterMulticast-MulticastRouting"); ok {
			if err = d.Set("multicast_routing", vv); err != nil {
				return fmt.Errorf("Error reading multicast_routing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multicast_routing: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("pim_sm_global", flattenRouterMulticastPimSmGlobal(o["pim-sm-global"], d, "pim_sm_global")); err != nil {
			if vv, ok := fortiAPIPatch(o["pim-sm-global"], "RouterMulticast-PimSmGlobal"); ok {
				if err = d.Set("pim_sm_global", vv); err != nil {
					return fmt.Errorf("Error reading pim_sm_global: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading pim_sm_global: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("pim_sm_global"); ok {
			if err = d.Set("pim_sm_global", flattenRouterMulticastPimSmGlobal(o["pim-sm-global"], d, "pim_sm_global")); err != nil {
				if vv, ok := fortiAPIPatch(o["pim-sm-global"], "RouterMulticast-PimSmGlobal"); ok {
					if err = d.Set("pim_sm_global", vv); err != nil {
						return fmt.Errorf("Error reading pim_sm_global: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading pim_sm_global: %v", err)
				}
			}
		}
	}

	if err = d.Set("route_limit", flattenRouterMulticastRouteLimit(o["route-limit"], d, "route_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-limit"], "RouterMulticast-RouteLimit"); ok {
			if err = d.Set("route_limit", vv); err != nil {
				return fmt.Errorf("Error reading route_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_limit: %v", err)
		}
	}

	if err = d.Set("route_threshold", flattenRouterMulticastRouteThreshold(o["route-threshold"], d, "route_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-threshold"], "RouterMulticast-RouteThreshold"); ok {
			if err = d.Set("route_threshold", vv); err != nil {
				return fmt.Errorf("Error reading route_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_threshold: %v", err)
		}
	}

	return nil
}

func flattenRouterMulticastFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterMulticastInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bfd"], _ = expandRouterMulticastInterfaceBfd(d, i["bfd"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cisco_exclude_genid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cisco-exclude-genid"], _ = expandRouterMulticastInterfaceCiscoExcludeGenid(d, i["cisco_exclude_genid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dr_priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dr-priority"], _ = expandRouterMulticastInterfaceDrPriority(d, i["dr_priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_holdtime"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hello-holdtime"], _ = expandRouterMulticastInterfaceHelloHoldtime(d, i["hello_holdtime"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hello-interval"], _ = expandRouterMulticastInterfaceHelloInterval(d, i["hello_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandRouterMulticastInterfaceIgmp(d, i["igmp"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["igmp"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "join_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandRouterMulticastInterfaceJoinGroup(d, i["join_group"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["join-group"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "multicast_flow"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["multicast-flow"], _ = expandRouterMulticastInterfaceMulticastFlow(d, i["multicast_flow"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandRouterMulticastInterfaceName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbour_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["neighbour-filter"], _ = expandRouterMulticastInterfaceNeighbourFilter(d, i["neighbour_filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["passive"], _ = expandRouterMulticastInterfacePassive(d, i["passive"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pim_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pim-mode"], _ = expandRouterMulticastInterfacePimMode(d, i["pim_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "propagation_delay"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["propagation-delay"], _ = expandRouterMulticastInterfacePropagationDelay(d, i["propagation_delay"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rp-candidate"], _ = expandRouterMulticastInterfaceRpCandidate(d, i["rp_candidate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rp-candidate-group"], _ = expandRouterMulticastInterfaceRpCandidateGroup(d, i["rp_candidate_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rp-candidate-interval"], _ = expandRouterMulticastInterfaceRpCandidateInterval(d, i["rp_candidate_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate_priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rp-candidate-priority"], _ = expandRouterMulticastInterfaceRpCandidatePriority(d, i["rp_candidate_priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rpf_nbr_fail_back"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rpf-nbr-fail-back"], _ = expandRouterMulticastInterfaceRpfNbrFailBack(d, i["rpf_nbr_fail_back"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rpf_nbr_fail_back_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rpf-nbr-fail-back-filter"], _ = expandRouterMulticastInterfaceRpfNbrFailBackFilter(d, i["rpf_nbr_fail_back_filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "state_refresh_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["state-refresh-interval"], _ = expandRouterMulticastInterfaceStateRefreshInterval(d, i["state_refresh_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "static_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["static-group"], _ = expandRouterMulticastInterfaceStaticGroup(d, i["static_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ttl_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ttl-threshold"], _ = expandRouterMulticastInterfaceTtlThreshold(d, i["ttl_threshold"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterMulticastInterfaceBfd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceCiscoExcludeGenid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceDrPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceHelloHoldtime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceHelloInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "access_group"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["access-group"], _ = expandRouterMulticastInterfaceIgmpAccessGroup(d, i["access_group"], pre_append)
	}
	pre_append = pre + ".0." + "immediate_leave_group"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["immediate-leave-group"], _ = expandRouterMulticastInterfaceIgmpImmediateLeaveGroup(d, i["immediate_leave_group"], pre_append)
	}
	pre_append = pre + ".0." + "last_member_query_count"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["last-member-query-count"], _ = expandRouterMulticastInterfaceIgmpLastMemberQueryCount(d, i["last_member_query_count"], pre_append)
	}
	pre_append = pre + ".0." + "last_member_query_interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["last-member-query-interval"], _ = expandRouterMulticastInterfaceIgmpLastMemberQueryInterval(d, i["last_member_query_interval"], pre_append)
	}
	pre_append = pre + ".0." + "query_interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["query-interval"], _ = expandRouterMulticastInterfaceIgmpQueryInterval(d, i["query_interval"], pre_append)
	}
	pre_append = pre + ".0." + "query_max_response_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["query-max-response-time"], _ = expandRouterMulticastInterfaceIgmpQueryMaxResponseTime(d, i["query_max_response_time"], pre_append)
	}
	pre_append = pre + ".0." + "query_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["query-timeout"], _ = expandRouterMulticastInterfaceIgmpQueryTimeout(d, i["query_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "router_alert_check"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["router-alert-check"], _ = expandRouterMulticastInterfaceIgmpRouterAlertCheck(d, i["router_alert_check"], pre_append)
	}
	pre_append = pre + ".0." + "version"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["version"], _ = expandRouterMulticastInterfaceIgmpVersion(d, i["version"], pre_append)
	}

	return result, nil
}

func expandRouterMulticastInterfaceIgmpAccessGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastInterfaceIgmpImmediateLeaveGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastInterfaceIgmpLastMemberQueryCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmpLastMemberQueryInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmpQueryInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmpQueryMaxResponseTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmpQueryTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmpRouterAlertCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmpVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceJoinGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["address"], _ = expandRouterMulticastInterfaceJoinGroupAddress(d, i["address"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterMulticastInterfaceJoinGroupAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceMulticastFlow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastInterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandRouterMulticastInterfaceNeighbourFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastInterfacePassive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfacePimMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfacePropagationDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceRpCandidate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceRpCandidateGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastInterfaceRpCandidateInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceRpCandidatePriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceRpfNbrFailBack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceRpfNbrFailBackFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastInterfaceStateRefreshInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceStaticGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastInterfaceTtlThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastMulticastRouting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "accept_register_list"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["accept-register-list"], _ = expandRouterMulticastPimSmGlobalAcceptRegisterList(d, i["accept_register_list"], pre_append)
	}
	pre_append = pre + ".0." + "accept_source_list"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["accept-source-list"], _ = expandRouterMulticastPimSmGlobalAcceptSourceList(d, i["accept_source_list"], pre_append)
	}
	pre_append = pre + ".0." + "bsr_allow_quick_refresh"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bsr-allow-quick-refresh"], _ = expandRouterMulticastPimSmGlobalBsrAllowQuickRefresh(d, i["bsr_allow_quick_refresh"], pre_append)
	}
	pre_append = pre + ".0." + "bsr_candidate"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bsr-candidate"], _ = expandRouterMulticastPimSmGlobalBsrCandidate(d, i["bsr_candidate"], pre_append)
	}
	pre_append = pre + ".0." + "bsr_hash"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bsr-hash"], _ = expandRouterMulticastPimSmGlobalBsrHash(d, i["bsr_hash"], pre_append)
	}
	pre_append = pre + ".0." + "bsr_interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bsr-interface"], _ = expandRouterMulticastPimSmGlobalBsrInterface(d, i["bsr_interface"], pre_append)
	}
	pre_append = pre + ".0." + "bsr_priority"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bsr-priority"], _ = expandRouterMulticastPimSmGlobalBsrPriority(d, i["bsr_priority"], pre_append)
	}
	pre_append = pre + ".0." + "cisco_crp_prefix"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cisco-crp-prefix"], _ = expandRouterMulticastPimSmGlobalCiscoCrpPrefix(d, i["cisco_crp_prefix"], pre_append)
	}
	pre_append = pre + ".0." + "cisco_ignore_rp_set_priority"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cisco-ignore-rp-set-priority"], _ = expandRouterMulticastPimSmGlobalCiscoIgnoreRpSetPriority(d, i["cisco_ignore_rp_set_priority"], pre_append)
	}
	pre_append = pre + ".0." + "cisco_register_checksum"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cisco-register-checksum"], _ = expandRouterMulticastPimSmGlobalCiscoRegisterChecksum(d, i["cisco_register_checksum"], pre_append)
	}
	pre_append = pre + ".0." + "cisco_register_checksum_group"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cisco-register-checksum-group"], _ = expandRouterMulticastPimSmGlobalCiscoRegisterChecksumGroup(d, i["cisco_register_checksum_group"], pre_append)
	}
	pre_append = pre + ".0." + "join_prune_holdtime"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["join-prune-holdtime"], _ = expandRouterMulticastPimSmGlobalJoinPruneHoldtime(d, i["join_prune_holdtime"], pre_append)
	}
	pre_append = pre + ".0." + "message_interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["message-interval"], _ = expandRouterMulticastPimSmGlobalMessageInterval(d, i["message_interval"], pre_append)
	}
	pre_append = pre + ".0." + "null_register_retries"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["null-register-retries"], _ = expandRouterMulticastPimSmGlobalNullRegisterRetries(d, i["null_register_retries"], pre_append)
	}
	pre_append = pre + ".0." + "pim_use_sdwan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pim-use-sdwan"], _ = expandRouterMulticastPimSmGlobalPimUseSdwan(d, i["pim_use_sdwan"], pre_append)
	}
	pre_append = pre + ".0." + "register_rate_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["register-rate-limit"], _ = expandRouterMulticastPimSmGlobalRegisterRateLimit(d, i["register_rate_limit"], pre_append)
	}
	pre_append = pre + ".0." + "register_rp_reachability"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["register-rp-reachability"], _ = expandRouterMulticastPimSmGlobalRegisterRpReachability(d, i["register_rp_reachability"], pre_append)
	}
	pre_append = pre + ".0." + "register_source"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["register-source"], _ = expandRouterMulticastPimSmGlobalRegisterSource(d, i["register_source"], pre_append)
	}
	pre_append = pre + ".0." + "register_source_interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["register-source-interface"], _ = expandRouterMulticastPimSmGlobalRegisterSourceInterface(d, i["register_source_interface"], pre_append)
	}
	pre_append = pre + ".0." + "register_source_ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["register-source-ip"], _ = expandRouterMulticastPimSmGlobalRegisterSourceIp(d, i["register_source_ip"], pre_append)
	}
	pre_append = pre + ".0." + "register_supression"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["register-supression"], _ = expandRouterMulticastPimSmGlobalRegisterSupression(d, i["register_supression"], pre_append)
	}
	pre_append = pre + ".0." + "rp_address"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandRouterMulticastPimSmGlobalRpAddress(d, i["rp_address"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["rp-address"] = t
		}
	}
	pre_append = pre + ".0." + "rp_register_keepalive"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["rp-register-keepalive"], _ = expandRouterMulticastPimSmGlobalRpRegisterKeepalive(d, i["rp_register_keepalive"], pre_append)
	}
	pre_append = pre + ".0." + "spt_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["spt-threshold"], _ = expandRouterMulticastPimSmGlobalSptThreshold(d, i["spt_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "spt_threshold_group"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["spt-threshold-group"], _ = expandRouterMulticastPimSmGlobalSptThresholdGroup(d, i["spt_threshold_group"], pre_append)
	}
	pre_append = pre + ".0." + "ssm"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ssm"], _ = expandRouterMulticastPimSmGlobalSsm(d, i["ssm"], pre_append)
	}
	pre_append = pre + ".0." + "ssm_range"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ssm-range"], _ = expandRouterMulticastPimSmGlobalSsmRange(d, i["ssm_range"], pre_append)
	}

	return result, nil
}

func expandRouterMulticastPimSmGlobalAcceptRegisterList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastPimSmGlobalAcceptSourceList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastPimSmGlobalBsrAllowQuickRefresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalBsrCandidate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalBsrHash(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalBsrInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastPimSmGlobalBsrPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalCiscoCrpPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalCiscoIgnoreRpSetPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalCiscoRegisterChecksum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalCiscoRegisterChecksumGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastPimSmGlobalJoinPruneHoldtime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalMessageInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalNullRegisterRetries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalPimUseSdwan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalRegisterRateLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalRegisterRpReachability(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalRegisterSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalRegisterSourceInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastPimSmGlobalRegisterSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalRegisterSupression(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalRpAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group"], _ = expandRouterMulticastPimSmGlobalRpAddressGroup(d, i["group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterMulticastPimSmGlobalRpAddressId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip-address"], _ = expandRouterMulticastPimSmGlobalRpAddressIpAddress(d, i["ip_address"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterMulticastPimSmGlobalRpAddressGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastPimSmGlobalRpAddressId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalRpAddressIpAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalRpRegisterKeepalive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalSptThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalSptThresholdGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastPimSmGlobalSsm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalSsmRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastRouteLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastRouteThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterMulticast(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandRouterMulticastInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("multicast_routing"); ok || d.HasChange("multicast_routing") {
		t, err := expandRouterMulticastMulticastRouting(d, v, "multicast_routing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-routing"] = t
		}
	}

	if v, ok := d.GetOk("pim_sm_global"); ok || d.HasChange("pim_sm_global") {
		t, err := expandRouterMulticastPimSmGlobal(d, v, "pim_sm_global")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pim-sm-global"] = t
		}
	}

	if v, ok := d.GetOk("route_limit"); ok || d.HasChange("route_limit") {
		t, err := expandRouterMulticastRouteLimit(d, v, "route_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-limit"] = t
		}
	}

	if v, ok := d.GetOk("route_threshold"); ok || d.HasChange("route_threshold") {
		t, err := expandRouterMulticastRouteThreshold(d, v, "route_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-threshold"] = t
		}
	}

	return &obj, nil
}
