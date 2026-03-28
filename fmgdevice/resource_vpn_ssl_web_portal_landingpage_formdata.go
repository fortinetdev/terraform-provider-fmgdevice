// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Form data.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnSslWebPortalLandingPageFormData() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnSslWebPortalLandingPageFormDataCreate,
		Read:   resourceVpnSslWebPortalLandingPageFormDataRead,
		Update: resourceVpnSslWebPortalLandingPageFormDataUpdate,
		Delete: resourceVpnSslWebPortalLandingPageFormDataDelete,

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
			"portal": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceVpnSslWebPortalLandingPageFormDataCreate(d *schema.ResourceData, m interface{}) error {
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
	portal := d.Get("portal").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["portal"] = portal

	obj, err := getObjectVpnSslWebPortalLandingPageFormData(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnSslWebPortalLandingPageFormData resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadVpnSslWebPortalLandingPageFormData(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateVpnSslWebPortalLandingPageFormData(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating VpnSslWebPortalLandingPageFormData resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateVpnSslWebPortalLandingPageFormData(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating VpnSslWebPortalLandingPageFormData resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceVpnSslWebPortalLandingPageFormDataRead(d, m)
}

func resourceVpnSslWebPortalLandingPageFormDataUpdate(d *schema.ResourceData, m interface{}) error {
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
	portal := d.Get("portal").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["portal"] = portal

	obj, err := getObjectVpnSslWebPortalLandingPageFormData(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebPortalLandingPageFormData resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateVpnSslWebPortalLandingPageFormData(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebPortalLandingPageFormData resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceVpnSslWebPortalLandingPageFormDataRead(d, m)
}

func resourceVpnSslWebPortalLandingPageFormDataDelete(d *schema.ResourceData, m interface{}) error {
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
	portal := d.Get("portal").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["portal"] = portal

	wsParams["adom"] = adomv

	err = c.DeleteVpnSslWebPortalLandingPageFormData(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VpnSslWebPortalLandingPageFormData resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnSslWebPortalLandingPageFormDataRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	portal := d.Get("portal").(string)
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
	if portal == "" {
		portal = importOptionChecking(m.(*FortiClient).Cfg, "portal")
		if portal == "" {
			return fmt.Errorf("Parameter portal is missing")
		}
		if err = d.Set("portal", portal); err != nil {
			return fmt.Errorf("Error set params portal: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["portal"] = portal

	o, err := c.ReadVpnSslWebPortalLandingPageFormData(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading VpnSslWebPortalLandingPageFormData resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnSslWebPortalLandingPageFormData(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebPortalLandingPageFormData resource from API: %v", err)
	}
	return nil
}

func flattenVpnSslWebPortalLandingPageFormDataName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalLandingPageFormDataValue3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnSslWebPortalLandingPageFormData(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenVpnSslWebPortalLandingPageFormDataName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "VpnSslWebPortalLandingPageFormData-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("value", flattenVpnSslWebPortalLandingPageFormDataValue3rdl(o["value"], d, "value")); err != nil {
		if vv, ok := fortiAPIPatch(o["value"], "VpnSslWebPortalLandingPageFormData-Value"); ok {
			if err = d.Set("value", vv); err != nil {
				return fmt.Errorf("Error reading value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	return nil
}

func flattenVpnSslWebPortalLandingPageFormDataFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnSslWebPortalLandingPageFormDataName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalLandingPageFormDataValue3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnSslWebPortalLandingPageFormData(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandVpnSslWebPortalLandingPageFormDataName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok || d.HasChange("value") {
		t, err := expandVpnSslWebPortalLandingPageFormDataValue3rdl(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
