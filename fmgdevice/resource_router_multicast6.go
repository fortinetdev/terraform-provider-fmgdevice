// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPv6 multicast.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterMulticast6() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterMulticast6Update,
		Read:   resourceRouterMulticast6Read,
		Update: resourceRouterMulticast6Update,
		Delete: resourceRouterMulticast6Delete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

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
			"interface": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hello_holdtime": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"hello_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rp_candidate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						"static_group": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"multicast_pmtu": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
						"bsr_allow_quick_refresh": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"bsr_candidate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						},
						"cisco_ignore_rp_set_priority": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
									"ip6_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"spt_threshold": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"spt_threshold_group": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"pim_sm_global_vrf": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
							Computed: true,
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
									"ip6_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"vrf": &schema.Schema{
							Type:     schema.TypeInt,
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

func resourceRouterMulticast6Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterMulticast6(d, false)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticast6 resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateRouterMulticast6(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticast6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("RouterMulticast6")

	return resourceRouterMulticast6Read(d, m)
}

func resourceRouterMulticast6Delete(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterMulticast6(d, true)

	if err != nil {
		return fmt.Errorf("Error updating RouterMulticast6 resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateRouterMulticast6(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error clearing RouterMulticast6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterMulticast6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterMulticast6(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading RouterMulticast6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterMulticast6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticast6 resource from API: %v", err)
	}
	return nil
}

func flattenRouterMulticast6Interface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_holdtime"
		if _, ok := i["hello-holdtime"]; ok {
			v := flattenRouterMulticast6InterfaceHelloHoldtime(i["hello-holdtime"], d, pre_append)
			tmp["hello_holdtime"] = fortiAPISubPartPatch(v, "RouterMulticast6-Interface-HelloHoldtime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := i["hello-interval"]; ok {
			v := flattenRouterMulticast6InterfaceHelloInterval(i["hello-interval"], d, pre_append)
			tmp["hello_interval"] = fortiAPISubPartPatch(v, "RouterMulticast6-Interface-HelloInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenRouterMulticast6InterfaceName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "RouterMulticast6-Interface-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate"
		if _, ok := i["rp-candidate"]; ok {
			v := flattenRouterMulticast6InterfaceRpCandidate(i["rp-candidate"], d, pre_append)
			tmp["rp_candidate"] = fortiAPISubPartPatch(v, "RouterMulticast6-Interface-RpCandidate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate_group"
		if _, ok := i["rp-candidate-group"]; ok {
			v := flattenRouterMulticast6InterfaceRpCandidateGroup(i["rp-candidate-group"], d, pre_append)
			tmp["rp_candidate_group"] = fortiAPISubPartPatch(v, "RouterMulticast6-Interface-RpCandidateGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate_interval"
		if _, ok := i["rp-candidate-interval"]; ok {
			v := flattenRouterMulticast6InterfaceRpCandidateInterval(i["rp-candidate-interval"], d, pre_append)
			tmp["rp_candidate_interval"] = fortiAPISubPartPatch(v, "RouterMulticast6-Interface-RpCandidateInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate_priority"
		if _, ok := i["rp-candidate-priority"]; ok {
			v := flattenRouterMulticast6InterfaceRpCandidatePriority(i["rp-candidate-priority"], d, pre_append)
			tmp["rp_candidate_priority"] = fortiAPISubPartPatch(v, "RouterMulticast6-Interface-RpCandidatePriority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "static_group"
		if _, ok := i["static-group"]; ok {
			v := flattenRouterMulticast6InterfaceStaticGroup(i["static-group"], d, pre_append)
			tmp["static_group"] = fortiAPISubPartPatch(v, "RouterMulticast6-Interface-StaticGroup")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterMulticast6InterfaceHelloHoldtime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6InterfaceHelloInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6InterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenRouterMulticast6InterfaceRpCandidate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6InterfaceRpCandidateGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticast6InterfaceRpCandidateInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6InterfaceRpCandidatePriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6InterfaceStaticGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticast6MulticastPmtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6MulticastRouting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobal(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "bsr_allow_quick_refresh"
	if _, ok := i["bsr-allow-quick-refresh"]; ok {
		result["bsr_allow_quick_refresh"] = flattenRouterMulticast6PimSmGlobalBsrAllowQuickRefresh(i["bsr-allow-quick-refresh"], d, pre_append)
	}

	pre_append = pre + ".0." + "bsr_candidate"
	if _, ok := i["bsr-candidate"]; ok {
		result["bsr_candidate"] = flattenRouterMulticast6PimSmGlobalBsrCandidate(i["bsr-candidate"], d, pre_append)
	}

	pre_append = pre + ".0." + "bsr_hash"
	if _, ok := i["bsr-hash"]; ok {
		result["bsr_hash"] = flattenRouterMulticast6PimSmGlobalBsrHash(i["bsr-hash"], d, pre_append)
	}

	pre_append = pre + ".0." + "bsr_interface"
	if _, ok := i["bsr-interface"]; ok {
		result["bsr_interface"] = flattenRouterMulticast6PimSmGlobalBsrInterface(i["bsr-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "bsr_priority"
	if _, ok := i["bsr-priority"]; ok {
		result["bsr_priority"] = flattenRouterMulticast6PimSmGlobalBsrPriority(i["bsr-priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "cisco_crp_prefix"
	if _, ok := i["cisco-crp-prefix"]; ok {
		result["cisco_crp_prefix"] = flattenRouterMulticast6PimSmGlobalCiscoCrpPrefix(i["cisco-crp-prefix"], d, pre_append)
	}

	pre_append = pre + ".0." + "cisco_ignore_rp_set_priority"
	if _, ok := i["cisco-ignore-rp-set-priority"]; ok {
		result["cisco_ignore_rp_set_priority"] = flattenRouterMulticast6PimSmGlobalCiscoIgnoreRpSetPriority(i["cisco-ignore-rp-set-priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "pim_use_sdwan"
	if _, ok := i["pim-use-sdwan"]; ok {
		result["pim_use_sdwan"] = flattenRouterMulticast6PimSmGlobalPimUseSdwan(i["pim-use-sdwan"], d, pre_append)
	}

	pre_append = pre + ".0." + "register_rate_limit"
	if _, ok := i["register-rate-limit"]; ok {
		result["register_rate_limit"] = flattenRouterMulticast6PimSmGlobalRegisterRateLimit(i["register-rate-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "rp_address"
	if _, ok := i["rp-address"]; ok {
		result["rp_address"] = flattenRouterMulticast6PimSmGlobalRpAddress(i["rp-address"], d, pre_append)
	}

	pre_append = pre + ".0." + "spt_threshold"
	if _, ok := i["spt-threshold"]; ok {
		result["spt_threshold"] = flattenRouterMulticast6PimSmGlobalSptThreshold(i["spt-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "spt_threshold_group"
	if _, ok := i["spt-threshold-group"]; ok {
		result["spt_threshold_group"] = flattenRouterMulticast6PimSmGlobalSptThresholdGroup(i["spt-threshold-group"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenRouterMulticast6PimSmGlobalBsrAllowQuickRefresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalBsrCandidate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalBsrHash(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalBsrInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticast6PimSmGlobalBsrPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalCiscoCrpPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalCiscoIgnoreRpSetPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalPimUseSdwan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalRegisterRateLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalRpAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterMulticast6PimSmGlobalRpAddressGroup(i["group"], d, pre_append)
			tmp["group"] = fortiAPISubPartPatch(v, "RouterMulticast6PimSmGlobal-RpAddress-Group")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterMulticast6PimSmGlobalRpAddressId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterMulticast6PimSmGlobal-RpAddress-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_address"
		if _, ok := i["ip6-address"]; ok {
			v := flattenRouterMulticast6PimSmGlobalRpAddressIp6Address(i["ip6-address"], d, pre_append)
			tmp["ip6_address"] = fortiAPISubPartPatch(v, "RouterMulticast6PimSmGlobal-RpAddress-Ip6Address")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterMulticast6PimSmGlobalRpAddressGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticast6PimSmGlobalRpAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalRpAddressIp6Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalSptThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalSptThresholdGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticast6PimSmGlobalVrf(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_allow_quick_refresh"
		if _, ok := i["bsr-allow-quick-refresh"]; ok {
			v := flattenRouterMulticast6PimSmGlobalVrfBsrAllowQuickRefresh(i["bsr-allow-quick-refresh"], d, pre_append)
			tmp["bsr_allow_quick_refresh"] = fortiAPISubPartPatch(v, "RouterMulticast6-PimSmGlobalVrf-BsrAllowQuickRefresh")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_candidate"
		if _, ok := i["bsr-candidate"]; ok {
			v := flattenRouterMulticast6PimSmGlobalVrfBsrCandidate(i["bsr-candidate"], d, pre_append)
			tmp["bsr_candidate"] = fortiAPISubPartPatch(v, "RouterMulticast6-PimSmGlobalVrf-BsrCandidate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_hash"
		if _, ok := i["bsr-hash"]; ok {
			v := flattenRouterMulticast6PimSmGlobalVrfBsrHash(i["bsr-hash"], d, pre_append)
			tmp["bsr_hash"] = fortiAPISubPartPatch(v, "RouterMulticast6-PimSmGlobalVrf-BsrHash")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_interface"
		if _, ok := i["bsr-interface"]; ok {
			v := flattenRouterMulticast6PimSmGlobalVrfBsrInterface(i["bsr-interface"], d, pre_append)
			tmp["bsr_interface"] = fortiAPISubPartPatch(v, "RouterMulticast6-PimSmGlobalVrf-BsrInterface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_priority"
		if _, ok := i["bsr-priority"]; ok {
			v := flattenRouterMulticast6PimSmGlobalVrfBsrPriority(i["bsr-priority"], d, pre_append)
			tmp["bsr_priority"] = fortiAPISubPartPatch(v, "RouterMulticast6-PimSmGlobalVrf-BsrPriority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cisco_crp_prefix"
		if _, ok := i["cisco-crp-prefix"]; ok {
			v := flattenRouterMulticast6PimSmGlobalVrfCiscoCrpPrefix(i["cisco-crp-prefix"], d, pre_append)
			tmp["cisco_crp_prefix"] = fortiAPISubPartPatch(v, "RouterMulticast6-PimSmGlobalVrf-CiscoCrpPrefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_address"
		if _, ok := i["rp-address"]; ok {
			v := flattenRouterMulticast6PimSmGlobalVrfRpAddress(i["rp-address"], d, pre_append)
			tmp["rp_address"] = fortiAPISubPartPatch(v, "RouterMulticast6-PimSmGlobalVrf-RpAddress")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := i["vrf"]; ok {
			v := flattenRouterMulticast6PimSmGlobalVrfVrf(i["vrf"], d, pre_append)
			tmp["vrf"] = fortiAPISubPartPatch(v, "RouterMulticast6-PimSmGlobalVrf-Vrf")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterMulticast6PimSmGlobalVrfBsrAllowQuickRefresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalVrfBsrCandidate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalVrfBsrHash(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalVrfBsrInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticast6PimSmGlobalVrfBsrPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalVrfCiscoCrpPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalVrfRpAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterMulticast6PimSmGlobalVrfRpAddressGroup(i["group"], d, pre_append)
			tmp["group"] = fortiAPISubPartPatch(v, "RouterMulticast6PimSmGlobalVrf-RpAddress-Group")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterMulticast6PimSmGlobalVrfRpAddressId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterMulticast6PimSmGlobalVrf-RpAddress-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_address"
		if _, ok := i["ip6-address"]; ok {
			v := flattenRouterMulticast6PimSmGlobalVrfRpAddressIp6Address(i["ip6-address"], d, pre_append)
			tmp["ip6_address"] = fortiAPISubPartPatch(v, "RouterMulticast6PimSmGlobalVrf-RpAddress-Ip6Address")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterMulticast6PimSmGlobalVrfRpAddressGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticast6PimSmGlobalVrfRpAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalVrfRpAddressIp6Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalVrfVrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterMulticast6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("interface", flattenRouterMulticast6Interface(o["interface"], d, "interface")); err != nil {
			if vv, ok := fortiAPIPatch(o["interface"], "RouterMulticast6-Interface"); ok {
				if err = d.Set("interface", vv); err != nil {
					return fmt.Errorf("Error reading interface: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("interface"); ok {
			if err = d.Set("interface", flattenRouterMulticast6Interface(o["interface"], d, "interface")); err != nil {
				if vv, ok := fortiAPIPatch(o["interface"], "RouterMulticast6-Interface"); ok {
					if err = d.Set("interface", vv); err != nil {
						return fmt.Errorf("Error reading interface: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading interface: %v", err)
				}
			}
		}
	}

	if err = d.Set("multicast_pmtu", flattenRouterMulticast6MulticastPmtu(o["multicast-pmtu"], d, "multicast_pmtu")); err != nil {
		if vv, ok := fortiAPIPatch(o["multicast-pmtu"], "RouterMulticast6-MulticastPmtu"); ok {
			if err = d.Set("multicast_pmtu", vv); err != nil {
				return fmt.Errorf("Error reading multicast_pmtu: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multicast_pmtu: %v", err)
		}
	}

	if err = d.Set("multicast_routing", flattenRouterMulticast6MulticastRouting(o["multicast-routing"], d, "multicast_routing")); err != nil {
		if vv, ok := fortiAPIPatch(o["multicast-routing"], "RouterMulticast6-MulticastRouting"); ok {
			if err = d.Set("multicast_routing", vv); err != nil {
				return fmt.Errorf("Error reading multicast_routing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multicast_routing: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("pim_sm_global", flattenRouterMulticast6PimSmGlobal(o["pim-sm-global"], d, "pim_sm_global")); err != nil {
			if vv, ok := fortiAPIPatch(o["pim-sm-global"], "RouterMulticast6-PimSmGlobal"); ok {
				if err = d.Set("pim_sm_global", vv); err != nil {
					return fmt.Errorf("Error reading pim_sm_global: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading pim_sm_global: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("pim_sm_global"); ok {
			if err = d.Set("pim_sm_global", flattenRouterMulticast6PimSmGlobal(o["pim-sm-global"], d, "pim_sm_global")); err != nil {
				if vv, ok := fortiAPIPatch(o["pim-sm-global"], "RouterMulticast6-PimSmGlobal"); ok {
					if err = d.Set("pim_sm_global", vv); err != nil {
						return fmt.Errorf("Error reading pim_sm_global: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading pim_sm_global: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("pim_sm_global_vrf", flattenRouterMulticast6PimSmGlobalVrf(o["pim-sm-global-vrf"], d, "pim_sm_global_vrf")); err != nil {
			if vv, ok := fortiAPIPatch(o["pim-sm-global-vrf"], "RouterMulticast6-PimSmGlobalVrf"); ok {
				if err = d.Set("pim_sm_global_vrf", vv); err != nil {
					return fmt.Errorf("Error reading pim_sm_global_vrf: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading pim_sm_global_vrf: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("pim_sm_global_vrf"); ok {
			if err = d.Set("pim_sm_global_vrf", flattenRouterMulticast6PimSmGlobalVrf(o["pim-sm-global-vrf"], d, "pim_sm_global_vrf")); err != nil {
				if vv, ok := fortiAPIPatch(o["pim-sm-global-vrf"], "RouterMulticast6-PimSmGlobalVrf"); ok {
					if err = d.Set("pim_sm_global_vrf", vv); err != nil {
						return fmt.Errorf("Error reading pim_sm_global_vrf: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading pim_sm_global_vrf: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenRouterMulticast6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterMulticast6Interface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_holdtime"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hello-holdtime"], _ = expandRouterMulticast6InterfaceHelloHoldtime(d, i["hello_holdtime"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hello-interval"], _ = expandRouterMulticast6InterfaceHelloInterval(d, i["hello_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandRouterMulticast6InterfaceName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rp-candidate"], _ = expandRouterMulticast6InterfaceRpCandidate(d, i["rp_candidate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rp-candidate-group"], _ = expandRouterMulticast6InterfaceRpCandidateGroup(d, i["rp_candidate_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rp-candidate-interval"], _ = expandRouterMulticast6InterfaceRpCandidateInterval(d, i["rp_candidate_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate_priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rp-candidate-priority"], _ = expandRouterMulticast6InterfaceRpCandidatePriority(d, i["rp_candidate_priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "static_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["static-group"], _ = expandRouterMulticast6InterfaceStaticGroup(d, i["static_group"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterMulticast6InterfaceHelloHoldtime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6InterfaceHelloInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6InterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandRouterMulticast6InterfaceRpCandidate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6InterfaceRpCandidateGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticast6InterfaceRpCandidateInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6InterfaceRpCandidatePriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6InterfaceStaticGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticast6MulticastPmtu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6MulticastRouting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "bsr_allow_quick_refresh"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bsr-allow-quick-refresh"], _ = expandRouterMulticast6PimSmGlobalBsrAllowQuickRefresh(d, i["bsr_allow_quick_refresh"], pre_append)
	}
	pre_append = pre + ".0." + "bsr_candidate"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bsr-candidate"], _ = expandRouterMulticast6PimSmGlobalBsrCandidate(d, i["bsr_candidate"], pre_append)
	}
	pre_append = pre + ".0." + "bsr_hash"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bsr-hash"], _ = expandRouterMulticast6PimSmGlobalBsrHash(d, i["bsr_hash"], pre_append)
	}
	pre_append = pre + ".0." + "bsr_interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bsr-interface"], _ = expandRouterMulticast6PimSmGlobalBsrInterface(d, i["bsr_interface"], pre_append)
	}
	pre_append = pre + ".0." + "bsr_priority"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bsr-priority"], _ = expandRouterMulticast6PimSmGlobalBsrPriority(d, i["bsr_priority"], pre_append)
	}
	pre_append = pre + ".0." + "cisco_crp_prefix"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cisco-crp-prefix"], _ = expandRouterMulticast6PimSmGlobalCiscoCrpPrefix(d, i["cisco_crp_prefix"], pre_append)
	}
	pre_append = pre + ".0." + "cisco_ignore_rp_set_priority"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cisco-ignore-rp-set-priority"], _ = expandRouterMulticast6PimSmGlobalCiscoIgnoreRpSetPriority(d, i["cisco_ignore_rp_set_priority"], pre_append)
	}
	pre_append = pre + ".0." + "pim_use_sdwan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pim-use-sdwan"], _ = expandRouterMulticast6PimSmGlobalPimUseSdwan(d, i["pim_use_sdwan"], pre_append)
	}
	pre_append = pre + ".0." + "register_rate_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["register-rate-limit"], _ = expandRouterMulticast6PimSmGlobalRegisterRateLimit(d, i["register_rate_limit"], pre_append)
	}
	pre_append = pre + ".0." + "rp_address"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandRouterMulticast6PimSmGlobalRpAddress(d, i["rp_address"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["rp-address"] = t
		}
	}
	pre_append = pre + ".0." + "spt_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["spt-threshold"], _ = expandRouterMulticast6PimSmGlobalSptThreshold(d, i["spt_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "spt_threshold_group"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["spt-threshold-group"], _ = expandRouterMulticast6PimSmGlobalSptThresholdGroup(d, i["spt_threshold_group"], pre_append)
	}

	return result, nil
}

func expandRouterMulticast6PimSmGlobalBsrAllowQuickRefresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalBsrCandidate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalBsrHash(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalBsrInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticast6PimSmGlobalBsrPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalCiscoCrpPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalCiscoIgnoreRpSetPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalPimUseSdwan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalRegisterRateLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalRpAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["group"], _ = expandRouterMulticast6PimSmGlobalRpAddressGroup(d, i["group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterMulticast6PimSmGlobalRpAddressId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip6-address"], _ = expandRouterMulticast6PimSmGlobalRpAddressIp6Address(d, i["ip6_address"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterMulticast6PimSmGlobalRpAddressGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticast6PimSmGlobalRpAddressId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalRpAddressIp6Address(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalSptThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalSptThresholdGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticast6PimSmGlobalVrf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_allow_quick_refresh"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bsr-allow-quick-refresh"], _ = expandRouterMulticast6PimSmGlobalVrfBsrAllowQuickRefresh(d, i["bsr_allow_quick_refresh"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_candidate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bsr-candidate"], _ = expandRouterMulticast6PimSmGlobalVrfBsrCandidate(d, i["bsr_candidate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_hash"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bsr-hash"], _ = expandRouterMulticast6PimSmGlobalVrfBsrHash(d, i["bsr_hash"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bsr-interface"], _ = expandRouterMulticast6PimSmGlobalVrfBsrInterface(d, i["bsr_interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bsr-priority"], _ = expandRouterMulticast6PimSmGlobalVrfBsrPriority(d, i["bsr_priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cisco_crp_prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cisco-crp-prefix"], _ = expandRouterMulticast6PimSmGlobalVrfCiscoCrpPrefix(d, i["cisco_crp_prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandRouterMulticast6PimSmGlobalVrfRpAddress(d, i["rp_address"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["rp-address"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrf"], _ = expandRouterMulticast6PimSmGlobalVrfVrf(d, i["vrf"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterMulticast6PimSmGlobalVrfBsrAllowQuickRefresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalVrfBsrCandidate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalVrfBsrHash(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalVrfBsrInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticast6PimSmGlobalVrfBsrPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalVrfCiscoCrpPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalVrfRpAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["group"], _ = expandRouterMulticast6PimSmGlobalVrfRpAddressGroup(d, i["group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterMulticast6PimSmGlobalVrfRpAddressId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip6-address"], _ = expandRouterMulticast6PimSmGlobalVrfRpAddressIp6Address(d, i["ip6_address"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterMulticast6PimSmGlobalVrfRpAddressGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticast6PimSmGlobalVrfRpAddressId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalVrfRpAddressIp6Address(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalVrfVrf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterMulticast6(d *schema.ResourceData, bemptysontable bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if bemptysontable {
		obj["interface"] = make([]struct{}, 0)
	} else {
		if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
			t, err := expandRouterMulticast6Interface(d, v, "interface")
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface"] = t
			}
		}
	}

	if v, ok := d.GetOk("multicast_pmtu"); ok || d.HasChange("multicast_pmtu") {
		t, err := expandRouterMulticast6MulticastPmtu(d, v, "multicast_pmtu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-pmtu"] = t
		}
	}

	if v, ok := d.GetOk("multicast_routing"); ok || d.HasChange("multicast_routing") {
		t, err := expandRouterMulticast6MulticastRouting(d, v, "multicast_routing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-routing"] = t
		}
	}

	if v, ok := d.GetOk("pim_sm_global"); ok || d.HasChange("pim_sm_global") {
		t, err := expandRouterMulticast6PimSmGlobal(d, v, "pim_sm_global")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pim-sm-global"] = t
		}
	}

	if bemptysontable {
		obj["pim-sm-global-vrf"] = make([]struct{}, 0)
	} else {
		if v, ok := d.GetOk("pim_sm_global_vrf"); ok || d.HasChange("pim_sm_global_vrf") {
			t, err := expandRouterMulticast6PimSmGlobalVrf(d, v, "pim_sm_global_vrf")
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pim-sm-global-vrf"] = t
			}
		}
	}

	return &obj, nil
}
