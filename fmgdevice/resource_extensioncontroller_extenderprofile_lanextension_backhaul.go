// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: LAN extension backhaul tunnel configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceExtensionControllerExtenderProfileLanExtensionBackhaul() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtensionControllerExtenderProfileLanExtensionBackhaulCreate,
		Read:   resourceExtensionControllerExtenderProfileLanExtensionBackhaulRead,
		Update: resourceExtensionControllerExtenderProfileLanExtensionBackhaulUpdate,
		Delete: resourceExtensionControllerExtenderProfileLanExtensionBackhaulDelete,

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
			"extender_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceExtensionControllerExtenderProfileLanExtensionBackhaulCreate(d *schema.ResourceData, m interface{}) error {
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
	extender_profile := d.Get("extender_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["extender_profile"] = extender_profile

	obj, err := getObjectExtensionControllerExtenderProfileLanExtensionBackhaul(d)
	if err != nil {
		return fmt.Errorf("Error creating ExtensionControllerExtenderProfileLanExtensionBackhaul resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadExtensionControllerExtenderProfileLanExtensionBackhaul(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateExtensionControllerExtenderProfileLanExtensionBackhaul(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ExtensionControllerExtenderProfileLanExtensionBackhaul resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateExtensionControllerExtenderProfileLanExtensionBackhaul(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ExtensionControllerExtenderProfileLanExtensionBackhaul resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceExtensionControllerExtenderProfileLanExtensionBackhaulRead(d, m)
}

func resourceExtensionControllerExtenderProfileLanExtensionBackhaulUpdate(d *schema.ResourceData, m interface{}) error {
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
	extender_profile := d.Get("extender_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["extender_profile"] = extender_profile

	obj, err := getObjectExtensionControllerExtenderProfileLanExtensionBackhaul(d)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerExtenderProfileLanExtensionBackhaul resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateExtensionControllerExtenderProfileLanExtensionBackhaul(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerExtenderProfileLanExtensionBackhaul resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceExtensionControllerExtenderProfileLanExtensionBackhaulRead(d, m)
}

func resourceExtensionControllerExtenderProfileLanExtensionBackhaulDelete(d *schema.ResourceData, m interface{}) error {
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
	extender_profile := d.Get("extender_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["extender_profile"] = extender_profile

	wsParams["adom"] = adomv

	err = c.DeleteExtensionControllerExtenderProfileLanExtensionBackhaul(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ExtensionControllerExtenderProfileLanExtensionBackhaul resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtensionControllerExtenderProfileLanExtensionBackhaulRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	extender_profile := d.Get("extender_profile").(string)
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
	if extender_profile == "" {
		extender_profile = importOptionChecking(m.(*FortiClient).Cfg, "extender_profile")
		if extender_profile == "" {
			return fmt.Errorf("Parameter extender_profile is missing")
		}
		if err = d.Set("extender_profile", extender_profile); err != nil {
			return fmt.Errorf("Error set params extender_profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["extender_profile"] = extender_profile

	o, err := c.ReadExtensionControllerExtenderProfileLanExtensionBackhaul(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ExtensionControllerExtenderProfileLanExtensionBackhaul resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectExtensionControllerExtenderProfileLanExtensionBackhaul(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ExtensionControllerExtenderProfileLanExtensionBackhaul resource from API: %v", err)
	}
	return nil
}

func flattenExtensionControllerExtenderProfileLanExtensionBackhaulName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionBackhaulPort3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionBackhaulRole3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionBackhaulWeight3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectExtensionControllerExtenderProfileLanExtensionBackhaul(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenExtensionControllerExtenderProfileLanExtensionBackhaulName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ExtensionControllerExtenderProfileLanExtensionBackhaul-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenExtensionControllerExtenderProfileLanExtensionBackhaulPort3rdl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ExtensionControllerExtenderProfileLanExtensionBackhaul-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("role", flattenExtensionControllerExtenderProfileLanExtensionBackhaulRole3rdl(o["role"], d, "role")); err != nil {
		if vv, ok := fortiAPIPatch(o["role"], "ExtensionControllerExtenderProfileLanExtensionBackhaul-Role"); ok {
			if err = d.Set("role", vv); err != nil {
				return fmt.Errorf("Error reading role: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading role: %v", err)
		}
	}

	if err = d.Set("weight", flattenExtensionControllerExtenderProfileLanExtensionBackhaulWeight3rdl(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "ExtensionControllerExtenderProfileLanExtensionBackhaul-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	return nil
}

func flattenExtensionControllerExtenderProfileLanExtensionBackhaulFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandExtensionControllerExtenderProfileLanExtensionBackhaulName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionBackhaulPort3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionBackhaulRole3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionBackhaulWeight3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectExtensionControllerExtenderProfileLanExtensionBackhaul(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandExtensionControllerExtenderProfileLanExtensionBackhaulName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandExtensionControllerExtenderProfileLanExtensionBackhaulPort3rdl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("role"); ok || d.HasChange("role") {
		t, err := expandExtensionControllerExtenderProfileLanExtensionBackhaulRole3rdl(d, v, "role")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["role"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandExtensionControllerExtenderProfileLanExtensionBackhaulWeight3rdl(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	return &obj, nil
}
