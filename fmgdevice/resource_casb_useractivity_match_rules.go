// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> CASB user activity rules.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCasbUserActivityMatchRules() *schema.Resource {
	return &schema.Resource{
		Create: resourceCasbUserActivityMatchRulesCreate,
		Read:   resourceCasbUserActivityMatchRulesRead,
		Update: resourceCasbUserActivityMatchRulesUpdate,
		Delete: resourceCasbUserActivityMatchRulesDelete,

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
			"user_activity": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"match": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"body_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"case_sensitive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domains": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"header_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"jq": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"methods": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceCasbUserActivityMatchRulesCreate(d *schema.ResourceData, m interface{}) error {
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
	user_activity := d.Get("user_activity").(string)
	match := d.Get("match").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["user_activity"] = user_activity
	paradict["match"] = match

	obj, err := getObjectCasbUserActivityMatchRules(d)
	if err != nil {
		return fmt.Errorf("Error creating CasbUserActivityMatchRules resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadCasbUserActivityMatchRules(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateCasbUserActivityMatchRules(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating CasbUserActivityMatchRules resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateCasbUserActivityMatchRules(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating CasbUserActivityMatchRules resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceCasbUserActivityMatchRulesRead(d, m)
			} else {
				return fmt.Errorf("Error creating CasbUserActivityMatchRules resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceCasbUserActivityMatchRulesRead(d, m)
}

func resourceCasbUserActivityMatchRulesUpdate(d *schema.ResourceData, m interface{}) error {
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
	user_activity := d.Get("user_activity").(string)
	match := d.Get("match").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["user_activity"] = user_activity
	paradict["match"] = match

	obj, err := getObjectCasbUserActivityMatchRules(d)
	if err != nil {
		return fmt.Errorf("Error updating CasbUserActivityMatchRules resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateCasbUserActivityMatchRules(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating CasbUserActivityMatchRules resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceCasbUserActivityMatchRulesRead(d, m)
}

func resourceCasbUserActivityMatchRulesDelete(d *schema.ResourceData, m interface{}) error {
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
	user_activity := d.Get("user_activity").(string)
	match := d.Get("match").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["user_activity"] = user_activity
	paradict["match"] = match

	wsParams["adom"] = adomv

	err = c.DeleteCasbUserActivityMatchRules(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting CasbUserActivityMatchRules resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCasbUserActivityMatchRulesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	user_activity := d.Get("user_activity").(string)
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
	if user_activity == "" {
		user_activity = importOptionChecking(m.(*FortiClient).Cfg, "user_activity")
		if user_activity == "" {
			return fmt.Errorf("Parameter user_activity is missing")
		}
		if err = d.Set("user_activity", user_activity); err != nil {
			return fmt.Errorf("Error set params user_activity: %v", err)
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
	paradict["user_activity"] = user_activity
	paradict["match"] = match

	o, err := c.ReadCasbUserActivityMatchRules(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading CasbUserActivityMatchRules resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCasbUserActivityMatchRules(d, o)
	if err != nil {
		return fmt.Errorf("Error reading CasbUserActivityMatchRules resource from API: %v", err)
	}
	return nil
}

func flattenCasbUserActivityMatchRulesBodyType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesCaseSensitive3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesDomains3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenCasbUserActivityMatchRulesHeaderName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesJq3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesMatchPattern3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesMatchValue3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesMethods3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenCasbUserActivityMatchRulesNegate3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectCasbUserActivityMatchRules(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("body_type", flattenCasbUserActivityMatchRulesBodyType3rdl(o["body-type"], d, "body_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["body-type"], "CasbUserActivityMatchRules-BodyType"); ok {
			if err = d.Set("body_type", vv); err != nil {
				return fmt.Errorf("Error reading body_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading body_type: %v", err)
		}
	}

	if err = d.Set("case_sensitive", flattenCasbUserActivityMatchRulesCaseSensitive3rdl(o["case-sensitive"], d, "case_sensitive")); err != nil {
		if vv, ok := fortiAPIPatch(o["case-sensitive"], "CasbUserActivityMatchRules-CaseSensitive"); ok {
			if err = d.Set("case_sensitive", vv); err != nil {
				return fmt.Errorf("Error reading case_sensitive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading case_sensitive: %v", err)
		}
	}

	if err = d.Set("domains", flattenCasbUserActivityMatchRulesDomains3rdl(o["domains"], d, "domains")); err != nil {
		if vv, ok := fortiAPIPatch(o["domains"], "CasbUserActivityMatchRules-Domains"); ok {
			if err = d.Set("domains", vv); err != nil {
				return fmt.Errorf("Error reading domains: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domains: %v", err)
		}
	}

	if err = d.Set("header_name", flattenCasbUserActivityMatchRulesHeaderName3rdl(o["header-name"], d, "header_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-name"], "CasbUserActivityMatchRules-HeaderName"); ok {
			if err = d.Set("header_name", vv); err != nil {
				return fmt.Errorf("Error reading header_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_name: %v", err)
		}
	}

	if err = d.Set("fosid", flattenCasbUserActivityMatchRulesId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "CasbUserActivityMatchRules-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("jq", flattenCasbUserActivityMatchRulesJq3rdl(o["jq"], d, "jq")); err != nil {
		if vv, ok := fortiAPIPatch(o["jq"], "CasbUserActivityMatchRules-Jq"); ok {
			if err = d.Set("jq", vv); err != nil {
				return fmt.Errorf("Error reading jq: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading jq: %v", err)
		}
	}

	if err = d.Set("match_pattern", flattenCasbUserActivityMatchRulesMatchPattern3rdl(o["match-pattern"], d, "match_pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-pattern"], "CasbUserActivityMatchRules-MatchPattern"); ok {
			if err = d.Set("match_pattern", vv); err != nil {
				return fmt.Errorf("Error reading match_pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_pattern: %v", err)
		}
	}

	if err = d.Set("match_value", flattenCasbUserActivityMatchRulesMatchValue3rdl(o["match-value"], d, "match_value")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-value"], "CasbUserActivityMatchRules-MatchValue"); ok {
			if err = d.Set("match_value", vv); err != nil {
				return fmt.Errorf("Error reading match_value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_value: %v", err)
		}
	}

	if err = d.Set("methods", flattenCasbUserActivityMatchRulesMethods3rdl(o["methods"], d, "methods")); err != nil {
		if vv, ok := fortiAPIPatch(o["methods"], "CasbUserActivityMatchRules-Methods"); ok {
			if err = d.Set("methods", vv); err != nil {
				return fmt.Errorf("Error reading methods: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading methods: %v", err)
		}
	}

	if err = d.Set("negate", flattenCasbUserActivityMatchRulesNegate3rdl(o["negate"], d, "negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["negate"], "CasbUserActivityMatchRules-Negate"); ok {
			if err = d.Set("negate", vv); err != nil {
				return fmt.Errorf("Error reading negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading negate: %v", err)
		}
	}

	if err = d.Set("type", flattenCasbUserActivityMatchRulesType3rdl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "CasbUserActivityMatchRules-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenCasbUserActivityMatchRulesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandCasbUserActivityMatchRulesBodyType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesCaseSensitive3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesDomains3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandCasbUserActivityMatchRulesHeaderName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesJq3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesMatchPattern3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesMatchValue3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesMethods3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandCasbUserActivityMatchRulesNegate3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectCasbUserActivityMatchRules(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("body_type"); ok || d.HasChange("body_type") {
		t, err := expandCasbUserActivityMatchRulesBodyType3rdl(d, v, "body_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["body-type"] = t
		}
	}

	if v, ok := d.GetOk("case_sensitive"); ok || d.HasChange("case_sensitive") {
		t, err := expandCasbUserActivityMatchRulesCaseSensitive3rdl(d, v, "case_sensitive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["case-sensitive"] = t
		}
	}

	if v, ok := d.GetOk("domains"); ok || d.HasChange("domains") {
		t, err := expandCasbUserActivityMatchRulesDomains3rdl(d, v, "domains")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domains"] = t
		}
	}

	if v, ok := d.GetOk("header_name"); ok || d.HasChange("header_name") {
		t, err := expandCasbUserActivityMatchRulesHeaderName3rdl(d, v, "header_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-name"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandCasbUserActivityMatchRulesId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("jq"); ok || d.HasChange("jq") {
		t, err := expandCasbUserActivityMatchRulesJq3rdl(d, v, "jq")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["jq"] = t
		}
	}

	if v, ok := d.GetOk("match_pattern"); ok || d.HasChange("match_pattern") {
		t, err := expandCasbUserActivityMatchRulesMatchPattern3rdl(d, v, "match_pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-pattern"] = t
		}
	}

	if v, ok := d.GetOk("match_value"); ok || d.HasChange("match_value") {
		t, err := expandCasbUserActivityMatchRulesMatchValue3rdl(d, v, "match_value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-value"] = t
		}
	}

	if v, ok := d.GetOk("methods"); ok || d.HasChange("methods") {
		t, err := expandCasbUserActivityMatchRulesMethods3rdl(d, v, "methods")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["methods"] = t
		}
	}

	if v, ok := d.GetOk("negate"); ok || d.HasChange("negate") {
		t, err := expandCasbUserActivityMatchRulesNegate3rdl(d, v, "negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["negate"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandCasbUserActivityMatchRulesType3rdl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
