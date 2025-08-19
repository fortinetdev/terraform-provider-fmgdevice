// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure CASB SaaS application.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCasbAttributeMatch() *schema.Resource {
	return &schema.Resource{
		Create: resourceCasbAttributeMatchCreate,
		Read:   resourceCasbAttributeMatchRead,
		Update: resourceCasbAttributeMatchUpdate,
		Delete: resourceCasbAttributeMatchDelete,

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
			"application": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"attribute": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"case_sensitive": &schema.Schema{
							Type:     schema.TypeString,
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
						"name": &schema.Schema{
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
			"match_strategy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceCasbAttributeMatchCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectCasbAttributeMatch(d)
	if err != nil {
		return fmt.Errorf("Error creating CasbAttributeMatch resource while getting object: %v", err)
	}

	_, err = c.CreateCasbAttributeMatch(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating CasbAttributeMatch resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceCasbAttributeMatchRead(d, m)
}

func resourceCasbAttributeMatchUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectCasbAttributeMatch(d)
	if err != nil {
		return fmt.Errorf("Error updating CasbAttributeMatch resource while getting object: %v", err)
	}

	_, err = c.UpdateCasbAttributeMatch(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating CasbAttributeMatch resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceCasbAttributeMatchRead(d, m)
}

func resourceCasbAttributeMatchDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteCasbAttributeMatch(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting CasbAttributeMatch resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCasbAttributeMatchRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadCasbAttributeMatch(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading CasbAttributeMatch resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCasbAttributeMatch(d, o)
	if err != nil {
		return fmt.Errorf("Error reading CasbAttributeMatch resource from API: %v", err)
	}
	return nil
}

func flattenCasbAttributeMatchApplication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenCasbAttributeMatchAttribute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitive"
		if _, ok := i["case-sensitive"]; ok {
			v := flattenCasbAttributeMatchAttributeCaseSensitive(i["case-sensitive"], d, pre_append)
			tmp["case_sensitive"] = fortiAPISubPartPatch(v, "CasbAttributeMatch-Attribute-CaseSensitive")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_pattern"
		if _, ok := i["match-pattern"]; ok {
			v := flattenCasbAttributeMatchAttributeMatchPattern(i["match-pattern"], d, pre_append)
			tmp["match_pattern"] = fortiAPISubPartPatch(v, "CasbAttributeMatch-Attribute-MatchPattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_value"
		if _, ok := i["match-value"]; ok {
			v := flattenCasbAttributeMatchAttributeMatchValue(i["match-value"], d, pre_append)
			tmp["match_value"] = fortiAPISubPartPatch(v, "CasbAttributeMatch-Attribute-MatchValue")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenCasbAttributeMatchAttributeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "CasbAttributeMatch-Attribute-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "negate"
		if _, ok := i["negate"]; ok {
			v := flattenCasbAttributeMatchAttributeNegate(i["negate"], d, pre_append)
			tmp["negate"] = fortiAPISubPartPatch(v, "CasbAttributeMatch-Attribute-Negate")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbAttributeMatchAttributeCaseSensitive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbAttributeMatchAttributeMatchPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbAttributeMatchAttributeMatchValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbAttributeMatchAttributeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbAttributeMatchAttributeNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbAttributeMatchMatchStrategy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbAttributeMatchName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectCasbAttributeMatch(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("application", flattenCasbAttributeMatchApplication(o["application"], d, "application")); err != nil {
		if vv, ok := fortiAPIPatch(o["application"], "CasbAttributeMatch-Application"); ok {
			if err = d.Set("application", vv); err != nil {
				return fmt.Errorf("Error reading application: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("attribute", flattenCasbAttributeMatchAttribute(o["attribute"], d, "attribute")); err != nil {
			if vv, ok := fortiAPIPatch(o["attribute"], "CasbAttributeMatch-Attribute"); ok {
				if err = d.Set("attribute", vv); err != nil {
					return fmt.Errorf("Error reading attribute: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading attribute: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("attribute"); ok {
			if err = d.Set("attribute", flattenCasbAttributeMatchAttribute(o["attribute"], d, "attribute")); err != nil {
				if vv, ok := fortiAPIPatch(o["attribute"], "CasbAttributeMatch-Attribute"); ok {
					if err = d.Set("attribute", vv); err != nil {
						return fmt.Errorf("Error reading attribute: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading attribute: %v", err)
				}
			}
		}
	}

	if err = d.Set("match_strategy", flattenCasbAttributeMatchMatchStrategy(o["match-strategy"], d, "match_strategy")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-strategy"], "CasbAttributeMatch-MatchStrategy"); ok {
			if err = d.Set("match_strategy", vv); err != nil {
				return fmt.Errorf("Error reading match_strategy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_strategy: %v", err)
		}
	}

	if err = d.Set("name", flattenCasbAttributeMatchName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "CasbAttributeMatch-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenCasbAttributeMatchFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandCasbAttributeMatchApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandCasbAttributeMatchAttribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitive"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["case-sensitive"], _ = expandCasbAttributeMatchAttributeCaseSensitive(d, i["case_sensitive"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-pattern"], _ = expandCasbAttributeMatchAttributeMatchPattern(d, i["match_pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-value"], _ = expandCasbAttributeMatchAttributeMatchValue(d, i["match_value"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandCasbAttributeMatchAttributeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "negate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["negate"], _ = expandCasbAttributeMatchAttributeNegate(d, i["negate"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandCasbAttributeMatchAttributeCaseSensitive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbAttributeMatchAttributeMatchPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbAttributeMatchAttributeMatchValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbAttributeMatchAttributeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbAttributeMatchAttributeNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbAttributeMatchMatchStrategy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbAttributeMatchName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectCasbAttributeMatch(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("application"); ok || d.HasChange("application") {
		t, err := expandCasbAttributeMatchApplication(d, v, "application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("attribute"); ok || d.HasChange("attribute") {
		t, err := expandCasbAttributeMatchAttribute(d, v, "attribute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attribute"] = t
		}
	}

	if v, ok := d.GetOk("match_strategy"); ok || d.HasChange("match_strategy") {
		t, err := expandCasbAttributeMatchMatchStrategy(d, v, "match_strategy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-strategy"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandCasbAttributeMatchName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
