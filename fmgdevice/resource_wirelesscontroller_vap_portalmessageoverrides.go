// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Individual message overrides.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerVapPortalMessageOverrides() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerVapPortalMessageOverridesUpdate,
		Read:   resourceWirelessControllerVapPortalMessageOverridesRead,
		Update: resourceWirelessControllerVapPortalMessageOverridesUpdate,
		Delete: resourceWirelessControllerVapPortalMessageOverridesDelete,

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
			"vap": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"auth_disclaimer_page": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_login_failed_page": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_login_page": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_reject_page": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceWirelessControllerVapPortalMessageOverridesUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	vap := d.Get("vap").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vap"] = vap

	obj, err := getObjectWirelessControllerVapPortalMessageOverrides(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerVapPortalMessageOverrides resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerVapPortalMessageOverrides(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerVapPortalMessageOverrides resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WirelessControllerVapPortalMessageOverrides")

	return resourceWirelessControllerVapPortalMessageOverridesRead(d, m)
}

func resourceWirelessControllerVapPortalMessageOverridesDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	vap := d.Get("vap").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vap"] = vap

	err = c.DeleteWirelessControllerVapPortalMessageOverrides(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerVapPortalMessageOverrides resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerVapPortalMessageOverridesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	vap := d.Get("vap").(string)
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
	if vap == "" {
		vap = importOptionChecking(m.(*FortiClient).Cfg, "vap")
		if vap == "" {
			return fmt.Errorf("Parameter vap is missing")
		}
		if err = d.Set("vap", vap); err != nil {
			return fmt.Errorf("Error set params vap: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vap"] = vap

	o, err := c.ReadWirelessControllerVapPortalMessageOverrides(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerVapPortalMessageOverrides resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerVapPortalMessageOverrides(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerVapPortalMessageOverrides resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerVapPortalMessageOverridesAuthDisclaimerPage2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapPortalMessageOverridesAuthLoginFailedPage2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapPortalMessageOverridesAuthLoginPage2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapPortalMessageOverridesAuthRejectPage2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerVapPortalMessageOverrides(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auth_disclaimer_page", flattenWirelessControllerVapPortalMessageOverridesAuthDisclaimerPage2edl(o["auth-disclaimer-page"], d, "auth_disclaimer_page")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-disclaimer-page"], "WirelessControllerVapPortalMessageOverrides-AuthDisclaimerPage"); ok {
			if err = d.Set("auth_disclaimer_page", vv); err != nil {
				return fmt.Errorf("Error reading auth_disclaimer_page: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_disclaimer_page: %v", err)
		}
	}

	if err = d.Set("auth_login_failed_page", flattenWirelessControllerVapPortalMessageOverridesAuthLoginFailedPage2edl(o["auth-login-failed-page"], d, "auth_login_failed_page")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-login-failed-page"], "WirelessControllerVapPortalMessageOverrides-AuthLoginFailedPage"); ok {
			if err = d.Set("auth_login_failed_page", vv); err != nil {
				return fmt.Errorf("Error reading auth_login_failed_page: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_login_failed_page: %v", err)
		}
	}

	if err = d.Set("auth_login_page", flattenWirelessControllerVapPortalMessageOverridesAuthLoginPage2edl(o["auth-login-page"], d, "auth_login_page")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-login-page"], "WirelessControllerVapPortalMessageOverrides-AuthLoginPage"); ok {
			if err = d.Set("auth_login_page", vv); err != nil {
				return fmt.Errorf("Error reading auth_login_page: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_login_page: %v", err)
		}
	}

	if err = d.Set("auth_reject_page", flattenWirelessControllerVapPortalMessageOverridesAuthRejectPage2edl(o["auth-reject-page"], d, "auth_reject_page")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-reject-page"], "WirelessControllerVapPortalMessageOverrides-AuthRejectPage"); ok {
			if err = d.Set("auth_reject_page", vv); err != nil {
				return fmt.Errorf("Error reading auth_reject_page: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_reject_page: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerVapPortalMessageOverridesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerVapPortalMessageOverridesAuthDisclaimerPage2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPortalMessageOverridesAuthLoginFailedPage2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPortalMessageOverridesAuthLoginPage2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPortalMessageOverridesAuthRejectPage2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerVapPortalMessageOverrides(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_disclaimer_page"); ok || d.HasChange("auth_disclaimer_page") {
		t, err := expandWirelessControllerVapPortalMessageOverridesAuthDisclaimerPage2edl(d, v, "auth_disclaimer_page")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-disclaimer-page"] = t
		}
	}

	if v, ok := d.GetOk("auth_login_failed_page"); ok || d.HasChange("auth_login_failed_page") {
		t, err := expandWirelessControllerVapPortalMessageOverridesAuthLoginFailedPage2edl(d, v, "auth_login_failed_page")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-login-failed-page"] = t
		}
	}

	if v, ok := d.GetOk("auth_login_page"); ok || d.HasChange("auth_login_page") {
		t, err := expandWirelessControllerVapPortalMessageOverridesAuthLoginPage2edl(d, v, "auth_login_page")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-login-page"] = t
		}
	}

	if v, ok := d.GetOk("auth_reject_page"); ok || d.HasChange("auth_reject_page") {
		t, err := expandWirelessControllerVapPortalMessageOverridesAuthRejectPage2edl(d, v, "auth_reject_page")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-reject-page"] = t
		}
	}

	return &obj, nil
}
