// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> CASB custom control option.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCasbProfileSaasApplicationCustomControlOption() *schema.Resource {
	return &schema.Resource{
		Create: resourceCasbProfileSaasApplicationCustomControlOptionCreate,
		Read:   resourceCasbProfileSaasApplicationCustomControlOptionRead,
		Update: resourceCasbProfileSaasApplicationCustomControlOptionUpdate,
		Delete: resourceCasbProfileSaasApplicationCustomControlOptionDelete,

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
			"saas_application": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"custom_control": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"user_input": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceCasbProfileSaasApplicationCustomControlOptionCreate(d *schema.ResourceData, m interface{}) error {
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
	saas_application := d.Get("saas_application").(string)
	custom_control := d.Get("custom_control").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application
	paradict["custom_control"] = custom_control

	obj, err := getObjectCasbProfileSaasApplicationCustomControlOption(d)
	if err != nil {
		return fmt.Errorf("Error creating CasbProfileSaasApplicationCustomControlOption resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadCasbProfileSaasApplicationCustomControlOption(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateCasbProfileSaasApplicationCustomControlOption(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating CasbProfileSaasApplicationCustomControlOption resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateCasbProfileSaasApplicationCustomControlOption(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating CasbProfileSaasApplicationCustomControlOption resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceCasbProfileSaasApplicationCustomControlOptionRead(d, m)
}

func resourceCasbProfileSaasApplicationCustomControlOptionUpdate(d *schema.ResourceData, m interface{}) error {
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
	saas_application := d.Get("saas_application").(string)
	custom_control := d.Get("custom_control").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application
	paradict["custom_control"] = custom_control

	obj, err := getObjectCasbProfileSaasApplicationCustomControlOption(d)
	if err != nil {
		return fmt.Errorf("Error updating CasbProfileSaasApplicationCustomControlOption resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateCasbProfileSaasApplicationCustomControlOption(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating CasbProfileSaasApplicationCustomControlOption resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceCasbProfileSaasApplicationCustomControlOptionRead(d, m)
}

func resourceCasbProfileSaasApplicationCustomControlOptionDelete(d *schema.ResourceData, m interface{}) error {
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
	saas_application := d.Get("saas_application").(string)
	custom_control := d.Get("custom_control").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application
	paradict["custom_control"] = custom_control

	wsParams["adom"] = adomv

	err = c.DeleteCasbProfileSaasApplicationCustomControlOption(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting CasbProfileSaasApplicationCustomControlOption resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCasbProfileSaasApplicationCustomControlOptionRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	profile := d.Get("profile").(string)
	saas_application := d.Get("saas_application").(string)
	custom_control := d.Get("custom_control").(string)
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
	if saas_application == "" {
		saas_application = importOptionChecking(m.(*FortiClient).Cfg, "saas_application")
		if saas_application == "" {
			return fmt.Errorf("Parameter saas_application is missing")
		}
		if err = d.Set("saas_application", saas_application); err != nil {
			return fmt.Errorf("Error set params saas_application: %v", err)
		}
	}
	if custom_control == "" {
		custom_control = importOptionChecking(m.(*FortiClient).Cfg, "custom_control")
		if custom_control == "" {
			return fmt.Errorf("Parameter custom_control is missing")
		}
		if err = d.Set("custom_control", custom_control); err != nil {
			return fmt.Errorf("Error set params custom_control: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application
	paradict["custom_control"] = custom_control

	o, err := c.ReadCasbProfileSaasApplicationCustomControlOption(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading CasbProfileSaasApplicationCustomControlOption resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCasbProfileSaasApplicationCustomControlOption(d, o)
	if err != nil {
		return fmt.Errorf("Error reading CasbProfileSaasApplicationCustomControlOption resource from API: %v", err)
	}
	return nil
}

func flattenCasbProfileSaasApplicationCustomControlOptionName4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationCustomControlOptionUserInput4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectCasbProfileSaasApplicationCustomControlOption(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenCasbProfileSaasApplicationCustomControlOptionName4thl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "CasbProfileSaasApplicationCustomControlOption-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("user_input", flattenCasbProfileSaasApplicationCustomControlOptionUserInput4thl(o["user-input"], d, "user_input")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-input"], "CasbProfileSaasApplicationCustomControlOption-UserInput"); ok {
			if err = d.Set("user_input", vv); err != nil {
				return fmt.Errorf("Error reading user_input: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_input: %v", err)
		}
	}

	return nil
}

func flattenCasbProfileSaasApplicationCustomControlOptionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandCasbProfileSaasApplicationCustomControlOptionName4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationCustomControlOptionUserInput4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectCasbProfileSaasApplicationCustomControlOption(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandCasbProfileSaasApplicationCustomControlOptionName4thl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("user_input"); ok || d.HasChange("user_input") {
		t, err := expandCasbProfileSaasApplicationCustomControlOptionUserInput4thl(d, v, "user_input")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-input"] = t
		}
	}

	return &obj, nil
}
