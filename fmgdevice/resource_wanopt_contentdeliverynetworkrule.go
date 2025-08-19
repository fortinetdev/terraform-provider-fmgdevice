// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure WAN optimization content delivery network rules.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWanoptContentDeliveryNetworkRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceWanoptContentDeliveryNetworkRuleCreate,
		Read:   resourceWanoptContentDeliveryNetworkRuleRead,
		Update: resourceWanoptContentDeliveryNetworkRuleUpdate,
		Delete: resourceWanoptContentDeliveryNetworkRuleDelete,

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
			"category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_domain_name_suffix": &schema.Schema{
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
			"request_cache_control": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"response_cache_control": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"response_expires": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rules": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"content_id": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end_direction": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"end_skip": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"end_str": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"range_str": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"start_direction": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"start_skip": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"start_str": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"target": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"match_entries": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"pattern": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"target": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"match_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"skip_entries": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"pattern": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"target": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"skip_rule_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"text_response_vcache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"updateserver": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceWanoptContentDeliveryNetworkRuleCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWanoptContentDeliveryNetworkRule(d)
	if err != nil {
		return fmt.Errorf("Error creating WanoptContentDeliveryNetworkRule resource while getting object: %v", err)
	}

	_, err = c.CreateWanoptContentDeliveryNetworkRule(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WanoptContentDeliveryNetworkRule resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWanoptContentDeliveryNetworkRuleRead(d, m)
}

func resourceWanoptContentDeliveryNetworkRuleUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWanoptContentDeliveryNetworkRule(d)
	if err != nil {
		return fmt.Errorf("Error updating WanoptContentDeliveryNetworkRule resource while getting object: %v", err)
	}

	_, err = c.UpdateWanoptContentDeliveryNetworkRule(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WanoptContentDeliveryNetworkRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWanoptContentDeliveryNetworkRuleRead(d, m)
}

func resourceWanoptContentDeliveryNetworkRuleDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteWanoptContentDeliveryNetworkRule(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WanoptContentDeliveryNetworkRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWanoptContentDeliveryNetworkRuleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	paradict["device"] = device_name

	o, err := c.ReadWanoptContentDeliveryNetworkRule(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WanoptContentDeliveryNetworkRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWanoptContentDeliveryNetworkRule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WanoptContentDeliveryNetworkRule resource from API: %v", err)
	}
	return nil
}

func flattenWanoptContentDeliveryNetworkRuleCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleHostDomainNameSuffix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWanoptContentDeliveryNetworkRuleName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRequestCacheControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleResponseCacheControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleResponseExpires(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRules(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content_id"
		if _, ok := i["content-id"]; ok {
			v := flattenWanoptContentDeliveryNetworkRuleRulesContentId(i["content-id"], d, pre_append)
			tmp["content_id"] = fortiAPISubPartPatch(v, "WanoptContentDeliveryNetworkRule-Rules-ContentId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_entries"
		if _, ok := i["match-entries"]; ok {
			v := flattenWanoptContentDeliveryNetworkRuleRulesMatchEntries(i["match-entries"], d, pre_append)
			tmp["match_entries"] = fortiAPISubPartPatch(v, "WanoptContentDeliveryNetworkRule-Rules-MatchEntries")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_mode"
		if _, ok := i["match-mode"]; ok {
			v := flattenWanoptContentDeliveryNetworkRuleRulesMatchMode(i["match-mode"], d, pre_append)
			tmp["match_mode"] = fortiAPISubPartPatch(v, "WanoptContentDeliveryNetworkRule-Rules-MatchMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenWanoptContentDeliveryNetworkRuleRulesName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "WanoptContentDeliveryNetworkRule-Rules-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "skip_entries"
		if _, ok := i["skip-entries"]; ok {
			v := flattenWanoptContentDeliveryNetworkRuleRulesSkipEntries(i["skip-entries"], d, pre_append)
			tmp["skip_entries"] = fortiAPISubPartPatch(v, "WanoptContentDeliveryNetworkRule-Rules-SkipEntries")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "skip_rule_mode"
		if _, ok := i["skip-rule-mode"]; ok {
			v := flattenWanoptContentDeliveryNetworkRuleRulesSkipRuleMode(i["skip-rule-mode"], d, pre_append)
			tmp["skip_rule_mode"] = fortiAPISubPartPatch(v, "WanoptContentDeliveryNetworkRule-Rules-SkipRuleMode")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentId(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "end_direction"
	if _, ok := i["end-direction"]; ok {
		result["end_direction"] = flattenWanoptContentDeliveryNetworkRuleRulesContentIdEndDirection(i["end-direction"], d, pre_append)
	}

	pre_append = pre + ".0." + "end_skip"
	if _, ok := i["end-skip"]; ok {
		result["end_skip"] = flattenWanoptContentDeliveryNetworkRuleRulesContentIdEndSkip(i["end-skip"], d, pre_append)
	}

	pre_append = pre + ".0." + "end_str"
	if _, ok := i["end-str"]; ok {
		result["end_str"] = flattenWanoptContentDeliveryNetworkRuleRulesContentIdEndStr(i["end-str"], d, pre_append)
	}

	pre_append = pre + ".0." + "range_str"
	if _, ok := i["range-str"]; ok {
		result["range_str"] = flattenWanoptContentDeliveryNetworkRuleRulesContentIdRangeStr(i["range-str"], d, pre_append)
	}

	pre_append = pre + ".0." + "start_direction"
	if _, ok := i["start-direction"]; ok {
		result["start_direction"] = flattenWanoptContentDeliveryNetworkRuleRulesContentIdStartDirection(i["start-direction"], d, pre_append)
	}

	pre_append = pre + ".0." + "start_skip"
	if _, ok := i["start-skip"]; ok {
		result["start_skip"] = flattenWanoptContentDeliveryNetworkRuleRulesContentIdStartSkip(i["start-skip"], d, pre_append)
	}

	pre_append = pre + ".0." + "start_str"
	if _, ok := i["start-str"]; ok {
		result["start_str"] = flattenWanoptContentDeliveryNetworkRuleRulesContentIdStartStr(i["start-str"], d, pre_append)
	}

	pre_append = pre + ".0." + "target"
	if _, ok := i["target"]; ok {
		result["target"] = flattenWanoptContentDeliveryNetworkRuleRulesContentIdTarget(i["target"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdEndDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdEndSkip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdEndStr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdRangeStr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdStartDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdStartSkip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdStartStr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdTarget(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesMatchEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWanoptContentDeliveryNetworkRuleRulesMatchEntriesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WanoptContentDeliveryNetworkRuleRules-MatchEntries-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			v := flattenWanoptContentDeliveryNetworkRuleRulesMatchEntriesPattern(i["pattern"], d, pre_append)
			tmp["pattern"] = fortiAPISubPartPatch(v, "WanoptContentDeliveryNetworkRuleRules-MatchEntries-Pattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := i["target"]; ok {
			v := flattenWanoptContentDeliveryNetworkRuleRulesMatchEntriesTarget(i["target"], d, pre_append)
			tmp["target"] = fortiAPISubPartPatch(v, "WanoptContentDeliveryNetworkRuleRules-MatchEntries-Target")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWanoptContentDeliveryNetworkRuleRulesMatchEntriesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesMatchEntriesPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWanoptContentDeliveryNetworkRuleRulesMatchEntriesTarget(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesMatchMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesSkipEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWanoptContentDeliveryNetworkRuleRulesSkipEntriesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WanoptContentDeliveryNetworkRuleRules-SkipEntries-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			v := flattenWanoptContentDeliveryNetworkRuleRulesSkipEntriesPattern(i["pattern"], d, pre_append)
			tmp["pattern"] = fortiAPISubPartPatch(v, "WanoptContentDeliveryNetworkRuleRules-SkipEntries-Pattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := i["target"]; ok {
			v := flattenWanoptContentDeliveryNetworkRuleRulesSkipEntriesTarget(i["target"], d, pre_append)
			tmp["target"] = fortiAPISubPartPatch(v, "WanoptContentDeliveryNetworkRuleRules-SkipEntries-Target")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWanoptContentDeliveryNetworkRuleRulesSkipEntriesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesSkipEntriesPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWanoptContentDeliveryNetworkRuleRulesSkipEntriesTarget(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesSkipRuleMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleTextResponseVcache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleUpdateserver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWanoptContentDeliveryNetworkRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("category", flattenWanoptContentDeliveryNetworkRuleCategory(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "WanoptContentDeliveryNetworkRule-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("comment", flattenWanoptContentDeliveryNetworkRuleComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "WanoptContentDeliveryNetworkRule-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("host_domain_name_suffix", flattenWanoptContentDeliveryNetworkRuleHostDomainNameSuffix(o["host-domain-name-suffix"], d, "host_domain_name_suffix")); err != nil {
		if vv, ok := fortiAPIPatch(o["host-domain-name-suffix"], "WanoptContentDeliveryNetworkRule-HostDomainNameSuffix"); ok {
			if err = d.Set("host_domain_name_suffix", vv); err != nil {
				return fmt.Errorf("Error reading host_domain_name_suffix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host_domain_name_suffix: %v", err)
		}
	}

	if err = d.Set("name", flattenWanoptContentDeliveryNetworkRuleName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WanoptContentDeliveryNetworkRule-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("request_cache_control", flattenWanoptContentDeliveryNetworkRuleRequestCacheControl(o["request-cache-control"], d, "request_cache_control")); err != nil {
		if vv, ok := fortiAPIPatch(o["request-cache-control"], "WanoptContentDeliveryNetworkRule-RequestCacheControl"); ok {
			if err = d.Set("request_cache_control", vv); err != nil {
				return fmt.Errorf("Error reading request_cache_control: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading request_cache_control: %v", err)
		}
	}

	if err = d.Set("response_cache_control", flattenWanoptContentDeliveryNetworkRuleResponseCacheControl(o["response-cache-control"], d, "response_cache_control")); err != nil {
		if vv, ok := fortiAPIPatch(o["response-cache-control"], "WanoptContentDeliveryNetworkRule-ResponseCacheControl"); ok {
			if err = d.Set("response_cache_control", vv); err != nil {
				return fmt.Errorf("Error reading response_cache_control: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading response_cache_control: %v", err)
		}
	}

	if err = d.Set("response_expires", flattenWanoptContentDeliveryNetworkRuleResponseExpires(o["response-expires"], d, "response_expires")); err != nil {
		if vv, ok := fortiAPIPatch(o["response-expires"], "WanoptContentDeliveryNetworkRule-ResponseExpires"); ok {
			if err = d.Set("response_expires", vv); err != nil {
				return fmt.Errorf("Error reading response_expires: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading response_expires: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("rules", flattenWanoptContentDeliveryNetworkRuleRules(o["rules"], d, "rules")); err != nil {
			if vv, ok := fortiAPIPatch(o["rules"], "WanoptContentDeliveryNetworkRule-Rules"); ok {
				if err = d.Set("rules", vv); err != nil {
					return fmt.Errorf("Error reading rules: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading rules: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("rules"); ok {
			if err = d.Set("rules", flattenWanoptContentDeliveryNetworkRuleRules(o["rules"], d, "rules")); err != nil {
				if vv, ok := fortiAPIPatch(o["rules"], "WanoptContentDeliveryNetworkRule-Rules"); ok {
					if err = d.Set("rules", vv); err != nil {
						return fmt.Errorf("Error reading rules: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading rules: %v", err)
				}
			}
		}
	}

	if err = d.Set("status", flattenWanoptContentDeliveryNetworkRuleStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WanoptContentDeliveryNetworkRule-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("text_response_vcache", flattenWanoptContentDeliveryNetworkRuleTextResponseVcache(o["text-response-vcache"], d, "text_response_vcache")); err != nil {
		if vv, ok := fortiAPIPatch(o["text-response-vcache"], "WanoptContentDeliveryNetworkRule-TextResponseVcache"); ok {
			if err = d.Set("text_response_vcache", vv); err != nil {
				return fmt.Errorf("Error reading text_response_vcache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading text_response_vcache: %v", err)
		}
	}

	if err = d.Set("updateserver", flattenWanoptContentDeliveryNetworkRuleUpdateserver(o["updateserver"], d, "updateserver")); err != nil {
		if vv, ok := fortiAPIPatch(o["updateserver"], "WanoptContentDeliveryNetworkRule-Updateserver"); ok {
			if err = d.Set("updateserver", vv); err != nil {
				return fmt.Errorf("Error reading updateserver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading updateserver: %v", err)
		}
	}

	return nil
}

func flattenWanoptContentDeliveryNetworkRuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWanoptContentDeliveryNetworkRuleCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleHostDomainNameSuffix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWanoptContentDeliveryNetworkRuleName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRequestCacheControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleResponseCacheControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleResponseExpires(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRules(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandWanoptContentDeliveryNetworkRuleRulesContentId(d, i["content_id"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["content-id"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_entries"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandWanoptContentDeliveryNetworkRuleRulesMatchEntries(d, i["match_entries"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["match-entries"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-mode"], _ = expandWanoptContentDeliveryNetworkRuleRulesMatchMode(d, i["match_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandWanoptContentDeliveryNetworkRuleRulesName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "skip_entries"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandWanoptContentDeliveryNetworkRuleRulesSkipEntries(d, i["skip_entries"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["skip-entries"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "skip_rule_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["skip-rule-mode"], _ = expandWanoptContentDeliveryNetworkRuleRulesSkipRuleMode(d, i["skip_rule_mode"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "end_direction"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["end-direction"], _ = expandWanoptContentDeliveryNetworkRuleRulesContentIdEndDirection(d, i["end_direction"], pre_append)
	}
	pre_append = pre + ".0." + "end_skip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["end-skip"], _ = expandWanoptContentDeliveryNetworkRuleRulesContentIdEndSkip(d, i["end_skip"], pre_append)
	}
	pre_append = pre + ".0." + "end_str"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["end-str"], _ = expandWanoptContentDeliveryNetworkRuleRulesContentIdEndStr(d, i["end_str"], pre_append)
	}
	pre_append = pre + ".0." + "range_str"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["range-str"], _ = expandWanoptContentDeliveryNetworkRuleRulesContentIdRangeStr(d, i["range_str"], pre_append)
	}
	pre_append = pre + ".0." + "start_direction"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["start-direction"], _ = expandWanoptContentDeliveryNetworkRuleRulesContentIdStartDirection(d, i["start_direction"], pre_append)
	}
	pre_append = pre + ".0." + "start_skip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["start-skip"], _ = expandWanoptContentDeliveryNetworkRuleRulesContentIdStartSkip(d, i["start_skip"], pre_append)
	}
	pre_append = pre + ".0." + "start_str"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["start-str"], _ = expandWanoptContentDeliveryNetworkRuleRulesContentIdStartStr(d, i["start_str"], pre_append)
	}
	pre_append = pre + ".0." + "target"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["target"], _ = expandWanoptContentDeliveryNetworkRuleRulesContentIdTarget(d, i["target"], pre_append)
	}

	return result, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdEndDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdEndSkip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdEndStr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdRangeStr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdStartDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdStartSkip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdStartStr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesMatchEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWanoptContentDeliveryNetworkRuleRulesMatchEntriesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern"], _ = expandWanoptContentDeliveryNetworkRuleRulesMatchEntriesPattern(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["target"], _ = expandWanoptContentDeliveryNetworkRuleRulesMatchEntriesTarget(d, i["target"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesMatchEntriesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesMatchEntriesPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWanoptContentDeliveryNetworkRuleRulesMatchEntriesTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesMatchMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesSkipEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWanoptContentDeliveryNetworkRuleRulesSkipEntriesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern"], _ = expandWanoptContentDeliveryNetworkRuleRulesSkipEntriesPattern(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["target"], _ = expandWanoptContentDeliveryNetworkRuleRulesSkipEntriesTarget(d, i["target"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesSkipEntriesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesSkipEntriesPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWanoptContentDeliveryNetworkRuleRulesSkipEntriesTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesSkipRuleMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleTextResponseVcache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleUpdateserver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWanoptContentDeliveryNetworkRule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandWanoptContentDeliveryNetworkRuleCategory(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandWanoptContentDeliveryNetworkRuleComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("host_domain_name_suffix"); ok || d.HasChange("host_domain_name_suffix") {
		t, err := expandWanoptContentDeliveryNetworkRuleHostDomainNameSuffix(d, v, "host_domain_name_suffix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-domain-name-suffix"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWanoptContentDeliveryNetworkRuleName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("request_cache_control"); ok || d.HasChange("request_cache_control") {
		t, err := expandWanoptContentDeliveryNetworkRuleRequestCacheControl(d, v, "request_cache_control")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-cache-control"] = t
		}
	}

	if v, ok := d.GetOk("response_cache_control"); ok || d.HasChange("response_cache_control") {
		t, err := expandWanoptContentDeliveryNetworkRuleResponseCacheControl(d, v, "response_cache_control")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["response-cache-control"] = t
		}
	}

	if v, ok := d.GetOk("response_expires"); ok || d.HasChange("response_expires") {
		t, err := expandWanoptContentDeliveryNetworkRuleResponseExpires(d, v, "response_expires")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["response-expires"] = t
		}
	}

	if v, ok := d.GetOk("rules"); ok || d.HasChange("rules") {
		t, err := expandWanoptContentDeliveryNetworkRuleRules(d, v, "rules")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rules"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWanoptContentDeliveryNetworkRuleStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("text_response_vcache"); ok || d.HasChange("text_response_vcache") {
		t, err := expandWanoptContentDeliveryNetworkRuleTextResponseVcache(d, v, "text_response_vcache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["text-response-vcache"] = t
		}
	}

	if v, ok := d.GetOk("updateserver"); ok || d.HasChange("updateserver") {
		t, err := expandWanoptContentDeliveryNetworkRuleUpdateserver(d, v, "updateserver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["updateserver"] = t
		}
	}

	return &obj, nil
}
