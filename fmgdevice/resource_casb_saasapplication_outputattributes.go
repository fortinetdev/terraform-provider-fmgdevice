// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> SaaS application output attributes.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCasbSaasApplicationOutputAttributes() *schema.Resource {
	return &schema.Resource{
		Create: resourceCasbSaasApplicationOutputAttributesCreate,
		Read:   resourceCasbSaasApplicationOutputAttributesRead,
		Update: resourceCasbSaasApplicationOutputAttributesUpdate,
		Delete: resourceCasbSaasApplicationOutputAttributesDelete,

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
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"optional": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"required": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceCasbSaasApplicationOutputAttributesCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectCasbSaasApplicationOutputAttributes(d)
	if err != nil {
		return fmt.Errorf("Error creating CasbSaasApplicationOutputAttributes resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadCasbSaasApplicationOutputAttributes(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateCasbSaasApplicationOutputAttributes(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating CasbSaasApplicationOutputAttributes resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateCasbSaasApplicationOutputAttributes(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating CasbSaasApplicationOutputAttributes resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceCasbSaasApplicationOutputAttributesRead(d, m)
}

func resourceCasbSaasApplicationOutputAttributesUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectCasbSaasApplicationOutputAttributes(d)
	if err != nil {
		return fmt.Errorf("Error updating CasbSaasApplicationOutputAttributes resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateCasbSaasApplicationOutputAttributes(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating CasbSaasApplicationOutputAttributes resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceCasbSaasApplicationOutputAttributesRead(d, m)
}

func resourceCasbSaasApplicationOutputAttributesDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteCasbSaasApplicationOutputAttributes(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting CasbSaasApplicationOutputAttributes resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCasbSaasApplicationOutputAttributesRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadCasbSaasApplicationOutputAttributes(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading CasbSaasApplicationOutputAttributes resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCasbSaasApplicationOutputAttributes(d, o)
	if err != nil {
		return fmt.Errorf("Error reading CasbSaasApplicationOutputAttributes resource from API: %v", err)
	}
	return nil
}

func flattenCasbSaasApplicationOutputAttributesAttrType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbSaasApplicationOutputAttributesDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbSaasApplicationOutputAttributesName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbSaasApplicationOutputAttributesOptional2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbSaasApplicationOutputAttributesRequired2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbSaasApplicationOutputAttributesType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectCasbSaasApplicationOutputAttributes(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("attr_type", flattenCasbSaasApplicationOutputAttributesAttrType2edl(o["attr-type"], d, "attr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["attr-type"], "CasbSaasApplicationOutputAttributes-AttrType"); ok {
			if err = d.Set("attr_type", vv); err != nil {
				return fmt.Errorf("Error reading attr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading attr_type: %v", err)
		}
	}

	if err = d.Set("description", flattenCasbSaasApplicationOutputAttributesDescription2edl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "CasbSaasApplicationOutputAttributes-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("name", flattenCasbSaasApplicationOutputAttributesName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "CasbSaasApplicationOutputAttributes-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("optional", flattenCasbSaasApplicationOutputAttributesOptional2edl(o["optional"], d, "optional")); err != nil {
		if vv, ok := fortiAPIPatch(o["optional"], "CasbSaasApplicationOutputAttributes-Optional"); ok {
			if err = d.Set("optional", vv); err != nil {
				return fmt.Errorf("Error reading optional: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading optional: %v", err)
		}
	}

	if err = d.Set("required", flattenCasbSaasApplicationOutputAttributesRequired2edl(o["required"], d, "required")); err != nil {
		if vv, ok := fortiAPIPatch(o["required"], "CasbSaasApplicationOutputAttributes-Required"); ok {
			if err = d.Set("required", vv); err != nil {
				return fmt.Errorf("Error reading required: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading required: %v", err)
		}
	}

	if err = d.Set("type", flattenCasbSaasApplicationOutputAttributesType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "CasbSaasApplicationOutputAttributes-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenCasbSaasApplicationOutputAttributesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandCasbSaasApplicationOutputAttributesAttrType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationOutputAttributesDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationOutputAttributesName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationOutputAttributesOptional2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationOutputAttributesRequired2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationOutputAttributesType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectCasbSaasApplicationOutputAttributes(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("attr_type"); ok || d.HasChange("attr_type") {
		t, err := expandCasbSaasApplicationOutputAttributesAttrType2edl(d, v, "attr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attr-type"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandCasbSaasApplicationOutputAttributesDescription2edl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandCasbSaasApplicationOutputAttributesName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("optional"); ok || d.HasChange("optional") {
		t, err := expandCasbSaasApplicationOutputAttributesOptional2edl(d, v, "optional")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["optional"] = t
		}
	}

	if v, ok := d.GetOk("required"); ok || d.HasChange("required") {
		t, err := expandCasbSaasApplicationOutputAttributesRequired2edl(d, v, "required")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["required"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandCasbSaasApplicationOutputAttributesType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
