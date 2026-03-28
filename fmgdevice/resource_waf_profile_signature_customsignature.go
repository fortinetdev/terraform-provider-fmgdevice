// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Custom signature.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWafProfileSignatureCustomSignature() *schema.Resource {
	return &schema.Resource{
		Create: resourceWafProfileSignatureCustomSignatureCreate,
		Read:   resourceWafProfileSignatureCustomSignatureRead,
		Update: resourceWafProfileSignatureCustomSignatureUpdate,
		Delete: resourceWafProfileSignatureCustomSignatureDelete,

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
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"case_sensitivity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"direction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"target": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWafProfileSignatureCustomSignatureCreate(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	obj, err := getObjectWafProfileSignatureCustomSignature(d)
	if err != nil {
		return fmt.Errorf("Error creating WafProfileSignatureCustomSignature resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWafProfileSignatureCustomSignature(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWafProfileSignatureCustomSignature(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WafProfileSignatureCustomSignature resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateWafProfileSignatureCustomSignature(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WafProfileSignatureCustomSignature resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceWafProfileSignatureCustomSignatureRead(d, m)
}

func resourceWafProfileSignatureCustomSignatureUpdate(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	obj, err := getObjectWafProfileSignatureCustomSignature(d)
	if err != nil {
		return fmt.Errorf("Error updating WafProfileSignatureCustomSignature resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWafProfileSignatureCustomSignature(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WafProfileSignatureCustomSignature resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWafProfileSignatureCustomSignatureRead(d, m)
}

func resourceWafProfileSignatureCustomSignatureDelete(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	wsParams["adom"] = adomv

	err = c.DeleteWafProfileSignatureCustomSignature(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WafProfileSignatureCustomSignature resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafProfileSignatureCustomSignatureRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	profile := d.Get("profile").(string)
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
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	o, err := c.ReadWafProfileSignatureCustomSignature(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WafProfileSignatureCustomSignature resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafProfileSignatureCustomSignature(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WafProfileSignatureCustomSignature resource from API: %v", err)
	}
	return nil
}

func flattenWafProfileSignatureCustomSignatureAction3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignatureCaseSensitivity3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignatureDirection3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignatureLog3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignatureName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignaturePattern3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignatureSeverity3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignatureStatus3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignatureTarget3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectWafProfileSignatureCustomSignature(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenWafProfileSignatureCustomSignatureAction3rdl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "WafProfileSignatureCustomSignature-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("case_sensitivity", flattenWafProfileSignatureCustomSignatureCaseSensitivity3rdl(o["case-sensitivity"], d, "case_sensitivity")); err != nil {
		if vv, ok := fortiAPIPatch(o["case-sensitivity"], "WafProfileSignatureCustomSignature-CaseSensitivity"); ok {
			if err = d.Set("case_sensitivity", vv); err != nil {
				return fmt.Errorf("Error reading case_sensitivity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading case_sensitivity: %v", err)
		}
	}

	if err = d.Set("direction", flattenWafProfileSignatureCustomSignatureDirection3rdl(o["direction"], d, "direction")); err != nil {
		if vv, ok := fortiAPIPatch(o["direction"], "WafProfileSignatureCustomSignature-Direction"); ok {
			if err = d.Set("direction", vv); err != nil {
				return fmt.Errorf("Error reading direction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading direction: %v", err)
		}
	}

	if err = d.Set("log", flattenWafProfileSignatureCustomSignatureLog3rdl(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "WafProfileSignatureCustomSignature-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("name", flattenWafProfileSignatureCustomSignatureName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WafProfileSignatureCustomSignature-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("pattern", flattenWafProfileSignatureCustomSignaturePattern3rdl(o["pattern"], d, "pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["pattern"], "WafProfileSignatureCustomSignature-Pattern"); ok {
			if err = d.Set("pattern", vv); err != nil {
				return fmt.Errorf("Error reading pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pattern: %v", err)
		}
	}

	if err = d.Set("severity", flattenWafProfileSignatureCustomSignatureSeverity3rdl(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "WafProfileSignatureCustomSignature-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("status", flattenWafProfileSignatureCustomSignatureStatus3rdl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WafProfileSignatureCustomSignature-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("target", flattenWafProfileSignatureCustomSignatureTarget3rdl(o["target"], d, "target")); err != nil {
		if vv, ok := fortiAPIPatch(o["target"], "WafProfileSignatureCustomSignature-Target"); ok {
			if err = d.Set("target", vv); err != nil {
				return fmt.Errorf("Error reading target: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading target: %v", err)
		}
	}

	return nil
}

func flattenWafProfileSignatureCustomSignatureFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWafProfileSignatureCustomSignatureAction3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignatureCaseSensitivity3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignatureDirection3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignatureLog3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignatureName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignaturePattern3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignatureSeverity3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignatureStatus3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignatureTarget3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectWafProfileSignatureCustomSignature(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandWafProfileSignatureCustomSignatureAction3rdl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("case_sensitivity"); ok || d.HasChange("case_sensitivity") {
		t, err := expandWafProfileSignatureCustomSignatureCaseSensitivity3rdl(d, v, "case_sensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["case-sensitivity"] = t
		}
	}

	if v, ok := d.GetOk("direction"); ok || d.HasChange("direction") {
		t, err := expandWafProfileSignatureCustomSignatureDirection3rdl(d, v, "direction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["direction"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandWafProfileSignatureCustomSignatureLog3rdl(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWafProfileSignatureCustomSignatureName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("pattern"); ok || d.HasChange("pattern") {
		t, err := expandWafProfileSignatureCustomSignaturePattern3rdl(d, v, "pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandWafProfileSignatureCustomSignatureSeverity3rdl(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWafProfileSignatureCustomSignatureStatus3rdl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("target"); ok || d.HasChange("target") {
		t, err := expandWafProfileSignatureCustomSignatureTarget3rdl(d, v, "target")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target"] = t
		}
	}

	return &obj, nil
}
