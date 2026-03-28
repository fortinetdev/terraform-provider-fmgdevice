// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> SaaS application input attributes.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCasbSaasApplicationInputAttributes() *schema.Resource {
	return &schema.Resource{
		Create: resourceCasbSaasApplicationInputAttributesCreate,
		Read:   resourceCasbSaasApplicationInputAttributesRead,
		Update: resourceCasbSaasApplicationInputAttributesUpdate,
		Delete: resourceCasbSaasApplicationInputAttributesDelete,

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
			"saas_application": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"attr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"default": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fallback_input": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"required": &schema.Schema{
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

func resourceCasbSaasApplicationInputAttributesCreate(d *schema.ResourceData, m interface{}) error {
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
	saas_application := d.Get("saas_application").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["saas_application"] = saas_application

	obj, err := getObjectCasbSaasApplicationInputAttributes(d)
	if err != nil {
		return fmt.Errorf("Error creating CasbSaasApplicationInputAttributes resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadCasbSaasApplicationInputAttributes(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateCasbSaasApplicationInputAttributes(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating CasbSaasApplicationInputAttributes resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateCasbSaasApplicationInputAttributes(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating CasbSaasApplicationInputAttributes resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceCasbSaasApplicationInputAttributesRead(d, m)
}

func resourceCasbSaasApplicationInputAttributesUpdate(d *schema.ResourceData, m interface{}) error {
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
	saas_application := d.Get("saas_application").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["saas_application"] = saas_application

	obj, err := getObjectCasbSaasApplicationInputAttributes(d)
	if err != nil {
		return fmt.Errorf("Error updating CasbSaasApplicationInputAttributes resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateCasbSaasApplicationInputAttributes(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating CasbSaasApplicationInputAttributes resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceCasbSaasApplicationInputAttributesRead(d, m)
}

func resourceCasbSaasApplicationInputAttributesDelete(d *schema.ResourceData, m interface{}) error {
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
	saas_application := d.Get("saas_application").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["saas_application"] = saas_application

	wsParams["adom"] = adomv

	err = c.DeleteCasbSaasApplicationInputAttributes(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting CasbSaasApplicationInputAttributes resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCasbSaasApplicationInputAttributesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	saas_application := d.Get("saas_application").(string)
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
	if saas_application == "" {
		saas_application = importOptionChecking(m.(*FortiClient).Cfg, "saas_application")
		if saas_application == "" {
			return fmt.Errorf("Parameter saas_application is missing")
		}
		if err = d.Set("saas_application", saas_application); err != nil {
			return fmt.Errorf("Error set params saas_application: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["saas_application"] = saas_application

	o, err := c.ReadCasbSaasApplicationInputAttributes(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading CasbSaasApplicationInputAttributes resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCasbSaasApplicationInputAttributes(d, o)
	if err != nil {
		return fmt.Errorf("Error reading CasbSaasApplicationInputAttributes resource from API: %v", err)
	}
	return nil
}

func flattenCasbSaasApplicationInputAttributesAttrType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbSaasApplicationInputAttributesDefault2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbSaasApplicationInputAttributesDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbSaasApplicationInputAttributesFallbackInput2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbSaasApplicationInputAttributesName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbSaasApplicationInputAttributesRequired2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbSaasApplicationInputAttributesType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectCasbSaasApplicationInputAttributes(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("attr_type", flattenCasbSaasApplicationInputAttributesAttrType2edl(o["attr-type"], d, "attr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["attr-type"], "CasbSaasApplicationInputAttributes-AttrType"); ok {
			if err = d.Set("attr_type", vv); err != nil {
				return fmt.Errorf("Error reading attr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading attr_type: %v", err)
		}
	}

	if err = d.Set("default", flattenCasbSaasApplicationInputAttributesDefault2edl(o["default"], d, "default")); err != nil {
		if vv, ok := fortiAPIPatch(o["default"], "CasbSaasApplicationInputAttributes-Default"); ok {
			if err = d.Set("default", vv); err != nil {
				return fmt.Errorf("Error reading default: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default: %v", err)
		}
	}

	if err = d.Set("description", flattenCasbSaasApplicationInputAttributesDescription2edl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "CasbSaasApplicationInputAttributes-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("fallback_input", flattenCasbSaasApplicationInputAttributesFallbackInput2edl(o["fallback-input"], d, "fallback_input")); err != nil {
		if vv, ok := fortiAPIPatch(o["fallback-input"], "CasbSaasApplicationInputAttributes-FallbackInput"); ok {
			if err = d.Set("fallback_input", vv); err != nil {
				return fmt.Errorf("Error reading fallback_input: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fallback_input: %v", err)
		}
	}

	if err = d.Set("name", flattenCasbSaasApplicationInputAttributesName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "CasbSaasApplicationInputAttributes-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("required", flattenCasbSaasApplicationInputAttributesRequired2edl(o["required"], d, "required")); err != nil {
		if vv, ok := fortiAPIPatch(o["required"], "CasbSaasApplicationInputAttributes-Required"); ok {
			if err = d.Set("required", vv); err != nil {
				return fmt.Errorf("Error reading required: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading required: %v", err)
		}
	}

	if err = d.Set("type", flattenCasbSaasApplicationInputAttributesType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "CasbSaasApplicationInputAttributes-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenCasbSaasApplicationInputAttributesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandCasbSaasApplicationInputAttributesAttrType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationInputAttributesDefault2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationInputAttributesDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationInputAttributesFallbackInput2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationInputAttributesName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationInputAttributesRequired2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationInputAttributesType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectCasbSaasApplicationInputAttributes(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("attr_type"); ok || d.HasChange("attr_type") {
		t, err := expandCasbSaasApplicationInputAttributesAttrType2edl(d, v, "attr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attr-type"] = t
		}
	}

	if v, ok := d.GetOk("default"); ok || d.HasChange("default") {
		t, err := expandCasbSaasApplicationInputAttributesDefault2edl(d, v, "default")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandCasbSaasApplicationInputAttributesDescription2edl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("fallback_input"); ok || d.HasChange("fallback_input") {
		t, err := expandCasbSaasApplicationInputAttributesFallbackInput2edl(d, v, "fallback_input")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fallback-input"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandCasbSaasApplicationInputAttributesName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("required"); ok || d.HasChange("required") {
		t, err := expandCasbSaasApplicationInputAttributesRequired2edl(d, v, "required")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["required"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandCasbSaasApplicationInputAttributesType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
