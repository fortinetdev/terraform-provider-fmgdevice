// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: WAN optimization content delivery network rule entries.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWanoptContentDeliveryNetworkRuleRules() *schema.Resource {
	return &schema.Resource{
		Create: resourceWanoptContentDeliveryNetworkRuleRulesCreate,
		Read:   resourceWanoptContentDeliveryNetworkRuleRulesRead,
		Update: resourceWanoptContentDeliveryNetworkRuleRulesUpdate,
		Delete: resourceWanoptContentDeliveryNetworkRuleRulesDelete,

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
			"content_delivery_network_rule": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
				ForceNew: true,
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceWanoptContentDeliveryNetworkRuleRulesCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	content_delivery_network_rule := d.Get("content_delivery_network_rule").(string)
	paradict["device"] = device_name
	paradict["content_delivery_network_rule"] = content_delivery_network_rule

	obj, err := getObjectWanoptContentDeliveryNetworkRuleRules(d)
	if err != nil {
		return fmt.Errorf("Error creating WanoptContentDeliveryNetworkRuleRules resource while getting object: %v", err)
	}

	_, err = c.CreateWanoptContentDeliveryNetworkRuleRules(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WanoptContentDeliveryNetworkRuleRules resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWanoptContentDeliveryNetworkRuleRulesRead(d, m)
}

func resourceWanoptContentDeliveryNetworkRuleRulesUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	content_delivery_network_rule := d.Get("content_delivery_network_rule").(string)
	paradict["device"] = device_name
	paradict["content_delivery_network_rule"] = content_delivery_network_rule

	obj, err := getObjectWanoptContentDeliveryNetworkRuleRules(d)
	if err != nil {
		return fmt.Errorf("Error updating WanoptContentDeliveryNetworkRuleRules resource while getting object: %v", err)
	}

	_, err = c.UpdateWanoptContentDeliveryNetworkRuleRules(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WanoptContentDeliveryNetworkRuleRules resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWanoptContentDeliveryNetworkRuleRulesRead(d, m)
}

func resourceWanoptContentDeliveryNetworkRuleRulesDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	content_delivery_network_rule := d.Get("content_delivery_network_rule").(string)
	paradict["device"] = device_name
	paradict["content_delivery_network_rule"] = content_delivery_network_rule

	err = c.DeleteWanoptContentDeliveryNetworkRuleRules(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WanoptContentDeliveryNetworkRuleRules resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWanoptContentDeliveryNetworkRuleRulesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	content_delivery_network_rule := d.Get("content_delivery_network_rule").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if content_delivery_network_rule == "" {
		content_delivery_network_rule = importOptionChecking(m.(*FortiClient).Cfg, "content_delivery_network_rule")
		if content_delivery_network_rule == "" {
			return fmt.Errorf("Parameter content_delivery_network_rule is missing")
		}
		if err = d.Set("content_delivery_network_rule", content_delivery_network_rule); err != nil {
			return fmt.Errorf("Error set params content_delivery_network_rule: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["content_delivery_network_rule"] = content_delivery_network_rule

	o, err := c.ReadWanoptContentDeliveryNetworkRuleRules(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WanoptContentDeliveryNetworkRuleRules resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWanoptContentDeliveryNetworkRuleRules(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WanoptContentDeliveryNetworkRuleRules resource from API: %v", err)
	}
	return nil
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentId2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "end_direction"
	if _, ok := i["end-direction"]; ok {
		result["end_direction"] = flattenWanoptContentDeliveryNetworkRuleRulesContentIdEndDirection2edl(i["end-direction"], d, pre_append)
	}

	pre_append = pre + ".0." + "end_skip"
	if _, ok := i["end-skip"]; ok {
		result["end_skip"] = flattenWanoptContentDeliveryNetworkRuleRulesContentIdEndSkip2edl(i["end-skip"], d, pre_append)
	}

	pre_append = pre + ".0." + "end_str"
	if _, ok := i["end-str"]; ok {
		result["end_str"] = flattenWanoptContentDeliveryNetworkRuleRulesContentIdEndStr2edl(i["end-str"], d, pre_append)
	}

	pre_append = pre + ".0." + "range_str"
	if _, ok := i["range-str"]; ok {
		result["range_str"] = flattenWanoptContentDeliveryNetworkRuleRulesContentIdRangeStr2edl(i["range-str"], d, pre_append)
	}

	pre_append = pre + ".0." + "start_direction"
	if _, ok := i["start-direction"]; ok {
		result["start_direction"] = flattenWanoptContentDeliveryNetworkRuleRulesContentIdStartDirection2edl(i["start-direction"], d, pre_append)
	}

	pre_append = pre + ".0." + "start_skip"
	if _, ok := i["start-skip"]; ok {
		result["start_skip"] = flattenWanoptContentDeliveryNetworkRuleRulesContentIdStartSkip2edl(i["start-skip"], d, pre_append)
	}

	pre_append = pre + ".0." + "start_str"
	if _, ok := i["start-str"]; ok {
		result["start_str"] = flattenWanoptContentDeliveryNetworkRuleRulesContentIdStartStr2edl(i["start-str"], d, pre_append)
	}

	pre_append = pre + ".0." + "target"
	if _, ok := i["target"]; ok {
		result["target"] = flattenWanoptContentDeliveryNetworkRuleRulesContentIdTarget2edl(i["target"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdEndDirection2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdEndSkip2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdEndStr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdRangeStr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdStartDirection2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdStartSkip2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdStartStr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdTarget2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesMatchEntries2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWanoptContentDeliveryNetworkRuleRulesMatchEntriesId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WanoptContentDeliveryNetworkRuleRules-MatchEntries-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			v := flattenWanoptContentDeliveryNetworkRuleRulesMatchEntriesPattern2edl(i["pattern"], d, pre_append)
			tmp["pattern"] = fortiAPISubPartPatch(v, "WanoptContentDeliveryNetworkRuleRules-MatchEntries-Pattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := i["target"]; ok {
			v := flattenWanoptContentDeliveryNetworkRuleRulesMatchEntriesTarget2edl(i["target"], d, pre_append)
			tmp["target"] = fortiAPISubPartPatch(v, "WanoptContentDeliveryNetworkRuleRules-MatchEntries-Target")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWanoptContentDeliveryNetworkRuleRulesMatchEntriesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesMatchEntriesPattern2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWanoptContentDeliveryNetworkRuleRulesMatchEntriesTarget2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesMatchMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesSkipEntries2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWanoptContentDeliveryNetworkRuleRulesSkipEntriesId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WanoptContentDeliveryNetworkRuleRules-SkipEntries-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			v := flattenWanoptContentDeliveryNetworkRuleRulesSkipEntriesPattern2edl(i["pattern"], d, pre_append)
			tmp["pattern"] = fortiAPISubPartPatch(v, "WanoptContentDeliveryNetworkRuleRules-SkipEntries-Pattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := i["target"]; ok {
			v := flattenWanoptContentDeliveryNetworkRuleRulesSkipEntriesTarget2edl(i["target"], d, pre_append)
			tmp["target"] = fortiAPISubPartPatch(v, "WanoptContentDeliveryNetworkRuleRules-SkipEntries-Target")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWanoptContentDeliveryNetworkRuleRulesSkipEntriesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesSkipEntriesPattern2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWanoptContentDeliveryNetworkRuleRulesSkipEntriesTarget2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesSkipRuleMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWanoptContentDeliveryNetworkRuleRules(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("content_id", flattenWanoptContentDeliveryNetworkRuleRulesContentId2edl(o["content-id"], d, "content_id")); err != nil {
			if vv, ok := fortiAPIPatch(o["content-id"], "WanoptContentDeliveryNetworkRuleRules-ContentId"); ok {
				if err = d.Set("content_id", vv); err != nil {
					return fmt.Errorf("Error reading content_id: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading content_id: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("content_id"); ok {
			if err = d.Set("content_id", flattenWanoptContentDeliveryNetworkRuleRulesContentId2edl(o["content-id"], d, "content_id")); err != nil {
				if vv, ok := fortiAPIPatch(o["content-id"], "WanoptContentDeliveryNetworkRuleRules-ContentId"); ok {
					if err = d.Set("content_id", vv); err != nil {
						return fmt.Errorf("Error reading content_id: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading content_id: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("match_entries", flattenWanoptContentDeliveryNetworkRuleRulesMatchEntries2edl(o["match-entries"], d, "match_entries")); err != nil {
			if vv, ok := fortiAPIPatch(o["match-entries"], "WanoptContentDeliveryNetworkRuleRules-MatchEntries"); ok {
				if err = d.Set("match_entries", vv); err != nil {
					return fmt.Errorf("Error reading match_entries: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading match_entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("match_entries"); ok {
			if err = d.Set("match_entries", flattenWanoptContentDeliveryNetworkRuleRulesMatchEntries2edl(o["match-entries"], d, "match_entries")); err != nil {
				if vv, ok := fortiAPIPatch(o["match-entries"], "WanoptContentDeliveryNetworkRuleRules-MatchEntries"); ok {
					if err = d.Set("match_entries", vv); err != nil {
						return fmt.Errorf("Error reading match_entries: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading match_entries: %v", err)
				}
			}
		}
	}

	if err = d.Set("match_mode", flattenWanoptContentDeliveryNetworkRuleRulesMatchMode2edl(o["match-mode"], d, "match_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-mode"], "WanoptContentDeliveryNetworkRuleRules-MatchMode"); ok {
			if err = d.Set("match_mode", vv); err != nil {
				return fmt.Errorf("Error reading match_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_mode: %v", err)
		}
	}

	if err = d.Set("name", flattenWanoptContentDeliveryNetworkRuleRulesName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WanoptContentDeliveryNetworkRuleRules-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("skip_entries", flattenWanoptContentDeliveryNetworkRuleRulesSkipEntries2edl(o["skip-entries"], d, "skip_entries")); err != nil {
			if vv, ok := fortiAPIPatch(o["skip-entries"], "WanoptContentDeliveryNetworkRuleRules-SkipEntries"); ok {
				if err = d.Set("skip_entries", vv); err != nil {
					return fmt.Errorf("Error reading skip_entries: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading skip_entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("skip_entries"); ok {
			if err = d.Set("skip_entries", flattenWanoptContentDeliveryNetworkRuleRulesSkipEntries2edl(o["skip-entries"], d, "skip_entries")); err != nil {
				if vv, ok := fortiAPIPatch(o["skip-entries"], "WanoptContentDeliveryNetworkRuleRules-SkipEntries"); ok {
					if err = d.Set("skip_entries", vv); err != nil {
						return fmt.Errorf("Error reading skip_entries: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading skip_entries: %v", err)
				}
			}
		}
	}

	if err = d.Set("skip_rule_mode", flattenWanoptContentDeliveryNetworkRuleRulesSkipRuleMode2edl(o["skip-rule-mode"], d, "skip_rule_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["skip-rule-mode"], "WanoptContentDeliveryNetworkRuleRules-SkipRuleMode"); ok {
			if err = d.Set("skip_rule_mode", vv); err != nil {
				return fmt.Errorf("Error reading skip_rule_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading skip_rule_mode: %v", err)
		}
	}

	return nil
}

func flattenWanoptContentDeliveryNetworkRuleRulesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWanoptContentDeliveryNetworkRuleRulesContentId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "end_direction"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["end-direction"], _ = expandWanoptContentDeliveryNetworkRuleRulesContentIdEndDirection2edl(d, i["end_direction"], pre_append)
	}
	pre_append = pre + ".0." + "end_skip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["end-skip"], _ = expandWanoptContentDeliveryNetworkRuleRulesContentIdEndSkip2edl(d, i["end_skip"], pre_append)
	}
	pre_append = pre + ".0." + "end_str"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["end-str"], _ = expandWanoptContentDeliveryNetworkRuleRulesContentIdEndStr2edl(d, i["end_str"], pre_append)
	}
	pre_append = pre + ".0." + "range_str"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["range-str"], _ = expandWanoptContentDeliveryNetworkRuleRulesContentIdRangeStr2edl(d, i["range_str"], pre_append)
	}
	pre_append = pre + ".0." + "start_direction"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["start-direction"], _ = expandWanoptContentDeliveryNetworkRuleRulesContentIdStartDirection2edl(d, i["start_direction"], pre_append)
	}
	pre_append = pre + ".0." + "start_skip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["start-skip"], _ = expandWanoptContentDeliveryNetworkRuleRulesContentIdStartSkip2edl(d, i["start_skip"], pre_append)
	}
	pre_append = pre + ".0." + "start_str"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["start-str"], _ = expandWanoptContentDeliveryNetworkRuleRulesContentIdStartStr2edl(d, i["start_str"], pre_append)
	}
	pre_append = pre + ".0." + "target"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["target"], _ = expandWanoptContentDeliveryNetworkRuleRulesContentIdTarget2edl(d, i["target"], pre_append)
	}

	return result, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdEndDirection2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdEndSkip2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdEndStr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdRangeStr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdStartDirection2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdStartSkip2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdStartStr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdTarget2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesMatchEntries2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWanoptContentDeliveryNetworkRuleRulesMatchEntriesId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern"], _ = expandWanoptContentDeliveryNetworkRuleRulesMatchEntriesPattern2edl(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["target"], _ = expandWanoptContentDeliveryNetworkRuleRulesMatchEntriesTarget2edl(d, i["target"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesMatchEntriesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesMatchEntriesPattern2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWanoptContentDeliveryNetworkRuleRulesMatchEntriesTarget2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesMatchMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesSkipEntries2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWanoptContentDeliveryNetworkRuleRulesSkipEntriesId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern"], _ = expandWanoptContentDeliveryNetworkRuleRulesSkipEntriesPattern2edl(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["target"], _ = expandWanoptContentDeliveryNetworkRuleRulesSkipEntriesTarget2edl(d, i["target"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesSkipEntriesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesSkipEntriesPattern2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWanoptContentDeliveryNetworkRuleRulesSkipEntriesTarget2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesSkipRuleMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWanoptContentDeliveryNetworkRuleRules(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("content_id"); ok || d.HasChange("content_id") {
		t, err := expandWanoptContentDeliveryNetworkRuleRulesContentId2edl(d, v, "content_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content-id"] = t
		}
	}

	if v, ok := d.GetOk("match_entries"); ok || d.HasChange("match_entries") {
		t, err := expandWanoptContentDeliveryNetworkRuleRulesMatchEntries2edl(d, v, "match_entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-entries"] = t
		}
	}

	if v, ok := d.GetOk("match_mode"); ok || d.HasChange("match_mode") {
		t, err := expandWanoptContentDeliveryNetworkRuleRulesMatchMode2edl(d, v, "match_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-mode"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWanoptContentDeliveryNetworkRuleRulesName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("skip_entries"); ok || d.HasChange("skip_entries") {
		t, err := expandWanoptContentDeliveryNetworkRuleRulesSkipEntries2edl(d, v, "skip_entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["skip-entries"] = t
		}
	}

	if v, ok := d.GetOk("skip_rule_mode"); ok || d.HasChange("skip_rule_mode") {
		t, err := expandWanoptContentDeliveryNetworkRuleRulesSkipRuleMode2edl(d, v, "skip_rule_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["skip-rule-mode"] = t
		}
	}

	return &obj, nil
}
