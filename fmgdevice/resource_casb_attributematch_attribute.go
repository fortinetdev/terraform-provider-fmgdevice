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

func resourceCasbAttributeMatchAttribute() *schema.Resource {
	return &schema.Resource{
		Create: resourceCasbAttributeMatchAttributeCreate,
		Read:   resourceCasbAttributeMatchAttributeRead,
		Update: resourceCasbAttributeMatchAttributeUpdate,
		Delete: resourceCasbAttributeMatchAttributeDelete,

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
			"attribute_match": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
				ForceNew: true,
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

func resourceCasbAttributeMatchAttributeCreate(d *schema.ResourceData, m interface{}) error {
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
	attribute_match := d.Get("attribute_match").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["attribute_match"] = attribute_match

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectCasbAttributeMatchAttribute(d)
	if err != nil {
		return fmt.Errorf("Error creating CasbAttributeMatchAttribute resource while getting object: %v", err)
	}

	_, err = c.CreateCasbAttributeMatchAttribute(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating CasbAttributeMatchAttribute resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceCasbAttributeMatchAttributeRead(d, m)
}

func resourceCasbAttributeMatchAttributeUpdate(d *schema.ResourceData, m interface{}) error {
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
	attribute_match := d.Get("attribute_match").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["attribute_match"] = attribute_match

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectCasbAttributeMatchAttribute(d)
	if err != nil {
		return fmt.Errorf("Error updating CasbAttributeMatchAttribute resource while getting object: %v", err)
	}

	_, err = c.UpdateCasbAttributeMatchAttribute(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating CasbAttributeMatchAttribute resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceCasbAttributeMatchAttributeRead(d, m)
}

func resourceCasbAttributeMatchAttributeDelete(d *schema.ResourceData, m interface{}) error {
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
	attribute_match := d.Get("attribute_match").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["attribute_match"] = attribute_match

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteCasbAttributeMatchAttribute(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting CasbAttributeMatchAttribute resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCasbAttributeMatchAttributeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadCasbAttributeMatchAttribute(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading CasbAttributeMatchAttribute resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCasbAttributeMatchAttribute(d, o)
	if err != nil {
		return fmt.Errorf("Error reading CasbAttributeMatchAttribute resource from API: %v", err)
	}
	return nil
}

func flattenCasbAttributeMatchAttributeCaseSensitive2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbAttributeMatchAttributeMatchPattern2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbAttributeMatchAttributeMatchValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbAttributeMatchAttributeName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbAttributeMatchAttributeNegate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectCasbAttributeMatchAttribute(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("case_sensitive", flattenCasbAttributeMatchAttributeCaseSensitive2edl(o["case-sensitive"], d, "case_sensitive")); err != nil {
		if vv, ok := fortiAPIPatch(o["case-sensitive"], "CasbAttributeMatchAttribute-CaseSensitive"); ok {
			if err = d.Set("case_sensitive", vv); err != nil {
				return fmt.Errorf("Error reading case_sensitive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading case_sensitive: %v", err)
		}
	}

	if err = d.Set("match_pattern", flattenCasbAttributeMatchAttributeMatchPattern2edl(o["match-pattern"], d, "match_pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-pattern"], "CasbAttributeMatchAttribute-MatchPattern"); ok {
			if err = d.Set("match_pattern", vv); err != nil {
				return fmt.Errorf("Error reading match_pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_pattern: %v", err)
		}
	}

	if err = d.Set("match_value", flattenCasbAttributeMatchAttributeMatchValue2edl(o["match-value"], d, "match_value")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-value"], "CasbAttributeMatchAttribute-MatchValue"); ok {
			if err = d.Set("match_value", vv); err != nil {
				return fmt.Errorf("Error reading match_value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_value: %v", err)
		}
	}

	if err = d.Set("name", flattenCasbAttributeMatchAttributeName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "CasbAttributeMatchAttribute-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("negate", flattenCasbAttributeMatchAttributeNegate2edl(o["negate"], d, "negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["negate"], "CasbAttributeMatchAttribute-Negate"); ok {
			if err = d.Set("negate", vv); err != nil {
				return fmt.Errorf("Error reading negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading negate: %v", err)
		}
	}

	return nil
}

func flattenCasbAttributeMatchAttributeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandCasbAttributeMatchAttributeCaseSensitive2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbAttributeMatchAttributeMatchPattern2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbAttributeMatchAttributeMatchValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbAttributeMatchAttributeName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbAttributeMatchAttributeNegate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectCasbAttributeMatchAttribute(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("case_sensitive"); ok || d.HasChange("case_sensitive") {
		t, err := expandCasbAttributeMatchAttributeCaseSensitive2edl(d, v, "case_sensitive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["case-sensitive"] = t
		}
	}

	if v, ok := d.GetOk("match_pattern"); ok || d.HasChange("match_pattern") {
		t, err := expandCasbAttributeMatchAttributeMatchPattern2edl(d, v, "match_pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-pattern"] = t
		}
	}

	if v, ok := d.GetOk("match_value"); ok || d.HasChange("match_value") {
		t, err := expandCasbAttributeMatchAttributeMatchValue2edl(d, v, "match_value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-value"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandCasbAttributeMatchAttributeName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("negate"); ok || d.HasChange("negate") {
		t, err := expandCasbAttributeMatchAttributeNegate2edl(d, v, "negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["negate"] = t
		}
	}

	return &obj, nil
}
