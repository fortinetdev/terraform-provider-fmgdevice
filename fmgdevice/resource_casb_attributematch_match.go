// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: CASB tenant match rules.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCasbAttributeMatchMatch() *schema.Resource {
	return &schema.Resource{
		Create: resourceCasbAttributeMatchMatchCreate,
		Read:   resourceCasbAttributeMatchMatchRead,
		Update: resourceCasbAttributeMatchMatchUpdate,
		Delete: resourceCasbAttributeMatchMatchDelete,

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
			"attribute_match": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attribute": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"case_sensitive": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"match_pattern": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"match_value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"negate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"rule_strategy": &schema.Schema{
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

func resourceCasbAttributeMatchMatchCreate(d *schema.ResourceData, m interface{}) error {
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
	attribute_match := d.Get("attribute_match").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["attribute_match"] = attribute_match

	obj, err := getObjectCasbAttributeMatchMatch(d)
	if err != nil {
		return fmt.Errorf("Error creating CasbAttributeMatchMatch resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadCasbAttributeMatchMatch(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateCasbAttributeMatchMatch(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating CasbAttributeMatchMatch resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateCasbAttributeMatchMatch(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating CasbAttributeMatchMatch resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceCasbAttributeMatchMatchRead(d, m)
			} else {
				return fmt.Errorf("Error creating CasbAttributeMatchMatch resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceCasbAttributeMatchMatchRead(d, m)
}

func resourceCasbAttributeMatchMatchUpdate(d *schema.ResourceData, m interface{}) error {
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
	attribute_match := d.Get("attribute_match").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["attribute_match"] = attribute_match

	obj, err := getObjectCasbAttributeMatchMatch(d)
	if err != nil {
		return fmt.Errorf("Error updating CasbAttributeMatchMatch resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateCasbAttributeMatchMatch(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating CasbAttributeMatchMatch resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceCasbAttributeMatchMatchRead(d, m)
}

func resourceCasbAttributeMatchMatchDelete(d *schema.ResourceData, m interface{}) error {
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
	attribute_match := d.Get("attribute_match").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["attribute_match"] = attribute_match

	wsParams["adom"] = adomv

	err = c.DeleteCasbAttributeMatchMatch(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting CasbAttributeMatchMatch resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCasbAttributeMatchMatchRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	attribute_match := d.Get("attribute_match").(string)
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
	if attribute_match == "" {
		attribute_match = importOptionChecking(m.(*FortiClient).Cfg, "attribute_match")
		if attribute_match == "" {
			return fmt.Errorf("Parameter attribute_match is missing")
		}
		if err = d.Set("attribute_match", attribute_match); err != nil {
			return fmt.Errorf("Error set params attribute_match: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["attribute_match"] = attribute_match

	o, err := c.ReadCasbAttributeMatchMatch(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading CasbAttributeMatchMatch resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCasbAttributeMatchMatch(d, o)
	if err != nil {
		return fmt.Errorf("Error reading CasbAttributeMatchMatch resource from API: %v", err)
	}
	return nil
}

func flattenCasbAttributeMatchMatchId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbAttributeMatchMatchRule2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute"
		if _, ok := i["attribute"]; ok {
			v := flattenCasbAttributeMatchMatchRuleAttribute2edl(i["attribute"], d, pre_append)
			tmp["attribute"] = fortiAPISubPartPatch(v, "CasbAttributeMatchMatch-Rule-Attribute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitive"
		if _, ok := i["case-sensitive"]; ok {
			v := flattenCasbAttributeMatchMatchRuleCaseSensitive2edl(i["case-sensitive"], d, pre_append)
			tmp["case_sensitive"] = fortiAPISubPartPatch(v, "CasbAttributeMatchMatch-Rule-CaseSensitive")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenCasbAttributeMatchMatchRuleId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "CasbAttributeMatchMatch-Rule-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_pattern"
		if _, ok := i["match-pattern"]; ok {
			v := flattenCasbAttributeMatchMatchRuleMatchPattern2edl(i["match-pattern"], d, pre_append)
			tmp["match_pattern"] = fortiAPISubPartPatch(v, "CasbAttributeMatchMatch-Rule-MatchPattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_value"
		if _, ok := i["match-value"]; ok {
			v := flattenCasbAttributeMatchMatchRuleMatchValue2edl(i["match-value"], d, pre_append)
			tmp["match_value"] = fortiAPISubPartPatch(v, "CasbAttributeMatchMatch-Rule-MatchValue")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "negate"
		if _, ok := i["negate"]; ok {
			v := flattenCasbAttributeMatchMatchRuleNegate2edl(i["negate"], d, pre_append)
			tmp["negate"] = fortiAPISubPartPatch(v, "CasbAttributeMatchMatch-Rule-Negate")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbAttributeMatchMatchRuleAttribute2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbAttributeMatchMatchRuleCaseSensitive2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbAttributeMatchMatchRuleId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbAttributeMatchMatchRuleMatchPattern2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbAttributeMatchMatchRuleMatchValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbAttributeMatchMatchRuleNegate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbAttributeMatchMatchRuleStrategy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectCasbAttributeMatchMatch(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("fosid", flattenCasbAttributeMatchMatchId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "CasbAttributeMatchMatch-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("rule", flattenCasbAttributeMatchMatchRule2edl(o["rule"], d, "rule")); err != nil {
			if vv, ok := fortiAPIPatch(o["rule"], "CasbAttributeMatchMatch-Rule"); ok {
				if err = d.Set("rule", vv); err != nil {
					return fmt.Errorf("Error reading rule: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading rule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("rule"); ok {
			if err = d.Set("rule", flattenCasbAttributeMatchMatchRule2edl(o["rule"], d, "rule")); err != nil {
				if vv, ok := fortiAPIPatch(o["rule"], "CasbAttributeMatchMatch-Rule"); ok {
					if err = d.Set("rule", vv); err != nil {
						return fmt.Errorf("Error reading rule: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading rule: %v", err)
				}
			}
		}
	}

	if err = d.Set("rule_strategy", flattenCasbAttributeMatchMatchRuleStrategy2edl(o["rule-strategy"], d, "rule_strategy")); err != nil {
		if vv, ok := fortiAPIPatch(o["rule-strategy"], "CasbAttributeMatchMatch-RuleStrategy"); ok {
			if err = d.Set("rule_strategy", vv); err != nil {
				return fmt.Errorf("Error reading rule_strategy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rule_strategy: %v", err)
		}
	}

	return nil
}

func flattenCasbAttributeMatchMatchFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandCasbAttributeMatchMatchId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbAttributeMatchMatchRule2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["attribute"], _ = expandCasbAttributeMatchMatchRuleAttribute2edl(d, i["attribute"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitive"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["case-sensitive"], _ = expandCasbAttributeMatchMatchRuleCaseSensitive2edl(d, i["case_sensitive"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandCasbAttributeMatchMatchRuleId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-pattern"], _ = expandCasbAttributeMatchMatchRuleMatchPattern2edl(d, i["match_pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-value"], _ = expandCasbAttributeMatchMatchRuleMatchValue2edl(d, i["match_value"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "negate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["negate"], _ = expandCasbAttributeMatchMatchRuleNegate2edl(d, i["negate"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandCasbAttributeMatchMatchRuleAttribute2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbAttributeMatchMatchRuleCaseSensitive2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbAttributeMatchMatchRuleId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbAttributeMatchMatchRuleMatchPattern2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbAttributeMatchMatchRuleMatchValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbAttributeMatchMatchRuleNegate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbAttributeMatchMatchRuleStrategy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectCasbAttributeMatchMatch(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandCasbAttributeMatchMatchId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("rule"); ok || d.HasChange("rule") {
		t, err := expandCasbAttributeMatchMatchRule2edl(d, v, "rule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule"] = t
		}
	}

	if v, ok := d.GetOk("rule_strategy"); ok || d.HasChange("rule_strategy") {
		t, err := expandCasbAttributeMatchMatchRuleStrategy2edl(d, v, "rule_strategy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule-strategy"] = t
		}
	}

	return &obj, nil
}
