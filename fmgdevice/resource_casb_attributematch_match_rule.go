// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: CASB attribute match rule.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCasbAttributeMatchMatchRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceCasbAttributeMatchMatchRuleCreate,
		Read:   resourceCasbAttributeMatchMatchRuleRead,
		Update: resourceCasbAttributeMatchMatchRuleUpdate,
		Delete: resourceCasbAttributeMatchMatchRuleDelete,

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
			"match": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"case_sensitive": &schema.Schema{
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
	}
}

func resourceCasbAttributeMatchMatchRuleCreate(d *schema.ResourceData, m interface{}) error {
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
	match := d.Get("match").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["attribute_match"] = attribute_match
	paradict["match"] = match

	obj, err := getObjectCasbAttributeMatchMatchRule(d)
	if err != nil {
		return fmt.Errorf("Error creating CasbAttributeMatchMatchRule resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadCasbAttributeMatchMatchRule(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateCasbAttributeMatchMatchRule(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating CasbAttributeMatchMatchRule resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateCasbAttributeMatchMatchRule(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating CasbAttributeMatchMatchRule resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceCasbAttributeMatchMatchRuleRead(d, m)
			} else {
				return fmt.Errorf("Error creating CasbAttributeMatchMatchRule resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceCasbAttributeMatchMatchRuleRead(d, m)
}

func resourceCasbAttributeMatchMatchRuleUpdate(d *schema.ResourceData, m interface{}) error {
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
	match := d.Get("match").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["attribute_match"] = attribute_match
	paradict["match"] = match

	obj, err := getObjectCasbAttributeMatchMatchRule(d)
	if err != nil {
		return fmt.Errorf("Error updating CasbAttributeMatchMatchRule resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateCasbAttributeMatchMatchRule(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating CasbAttributeMatchMatchRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceCasbAttributeMatchMatchRuleRead(d, m)
}

func resourceCasbAttributeMatchMatchRuleDelete(d *schema.ResourceData, m interface{}) error {
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
	match := d.Get("match").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["attribute_match"] = attribute_match
	paradict["match"] = match

	wsParams["adom"] = adomv

	err = c.DeleteCasbAttributeMatchMatchRule(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting CasbAttributeMatchMatchRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCasbAttributeMatchMatchRuleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	attribute_match := d.Get("attribute_match").(string)
	match := d.Get("match").(string)
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
	if match == "" {
		match = importOptionChecking(m.(*FortiClient).Cfg, "match")
		if match == "" {
			return fmt.Errorf("Parameter match is missing")
		}
		if err = d.Set("match", match); err != nil {
			return fmt.Errorf("Error set params match: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["attribute_match"] = attribute_match
	paradict["match"] = match

	o, err := c.ReadCasbAttributeMatchMatchRule(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading CasbAttributeMatchMatchRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCasbAttributeMatchMatchRule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading CasbAttributeMatchMatchRule resource from API: %v", err)
	}
	return nil
}

func flattenCasbAttributeMatchMatchRuleAttribute3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbAttributeMatchMatchRuleCaseSensitive3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbAttributeMatchMatchRuleId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbAttributeMatchMatchRuleMatchPattern3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbAttributeMatchMatchRuleMatchValue3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbAttributeMatchMatchRuleNegate3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectCasbAttributeMatchMatchRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("attribute", flattenCasbAttributeMatchMatchRuleAttribute3rdl(o["attribute"], d, "attribute")); err != nil {
		if vv, ok := fortiAPIPatch(o["attribute"], "CasbAttributeMatchMatchRule-Attribute"); ok {
			if err = d.Set("attribute", vv); err != nil {
				return fmt.Errorf("Error reading attribute: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading attribute: %v", err)
		}
	}

	if err = d.Set("case_sensitive", flattenCasbAttributeMatchMatchRuleCaseSensitive3rdl(o["case-sensitive"], d, "case_sensitive")); err != nil {
		if vv, ok := fortiAPIPatch(o["case-sensitive"], "CasbAttributeMatchMatchRule-CaseSensitive"); ok {
			if err = d.Set("case_sensitive", vv); err != nil {
				return fmt.Errorf("Error reading case_sensitive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading case_sensitive: %v", err)
		}
	}

	if err = d.Set("fosid", flattenCasbAttributeMatchMatchRuleId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "CasbAttributeMatchMatchRule-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("match_pattern", flattenCasbAttributeMatchMatchRuleMatchPattern3rdl(o["match-pattern"], d, "match_pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-pattern"], "CasbAttributeMatchMatchRule-MatchPattern"); ok {
			if err = d.Set("match_pattern", vv); err != nil {
				return fmt.Errorf("Error reading match_pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_pattern: %v", err)
		}
	}

	if err = d.Set("match_value", flattenCasbAttributeMatchMatchRuleMatchValue3rdl(o["match-value"], d, "match_value")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-value"], "CasbAttributeMatchMatchRule-MatchValue"); ok {
			if err = d.Set("match_value", vv); err != nil {
				return fmt.Errorf("Error reading match_value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_value: %v", err)
		}
	}

	if err = d.Set("negate", flattenCasbAttributeMatchMatchRuleNegate3rdl(o["negate"], d, "negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["negate"], "CasbAttributeMatchMatchRule-Negate"); ok {
			if err = d.Set("negate", vv); err != nil {
				return fmt.Errorf("Error reading negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading negate: %v", err)
		}
	}

	return nil
}

func flattenCasbAttributeMatchMatchRuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandCasbAttributeMatchMatchRuleAttribute3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbAttributeMatchMatchRuleCaseSensitive3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbAttributeMatchMatchRuleId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbAttributeMatchMatchRuleMatchPattern3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbAttributeMatchMatchRuleMatchValue3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbAttributeMatchMatchRuleNegate3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectCasbAttributeMatchMatchRule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("attribute"); ok || d.HasChange("attribute") {
		t, err := expandCasbAttributeMatchMatchRuleAttribute3rdl(d, v, "attribute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attribute"] = t
		}
	}

	if v, ok := d.GetOk("case_sensitive"); ok || d.HasChange("case_sensitive") {
		t, err := expandCasbAttributeMatchMatchRuleCaseSensitive3rdl(d, v, "case_sensitive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["case-sensitive"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandCasbAttributeMatchMatchRuleId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("match_pattern"); ok || d.HasChange("match_pattern") {
		t, err := expandCasbAttributeMatchMatchRuleMatchPattern3rdl(d, v, "match_pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-pattern"] = t
		}
	}

	if v, ok := d.GetOk("match_value"); ok || d.HasChange("match_value") {
		t, err := expandCasbAttributeMatchMatchRuleMatchValue3rdl(d, v, "match_value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-value"] = t
		}
	}

	if v, ok := d.GetOk("negate"); ok || d.HasChange("negate") {
		t, err := expandCasbAttributeMatchMatchRuleNegate3rdl(d, v, "negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["negate"] = t
		}
	}

	return &obj, nil
}
