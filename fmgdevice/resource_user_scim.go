// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure SCIM client entries.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserScim() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserScimCreate,
		Read:   resourceUserScimRead,
		Update: resourceUserScimUpdate,
		Delete: resourceUserScimDelete,

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
			"base_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"certificate": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"client_authentication_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_identity_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_secret_token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceUserScimCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectUserScim(d)
	if err != nil {
		return fmt.Errorf("Error creating UserScim resource while getting object: %v", err)
	}

	_, err = c.CreateUserScim(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating UserScim resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceUserScimRead(d, m)
}

func resourceUserScimUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectUserScim(d)
	if err != nil {
		return fmt.Errorf("Error updating UserScim resource while getting object: %v", err)
	}

	_, err = c.UpdateUserScim(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating UserScim resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceUserScimRead(d, m)
}

func resourceUserScimDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	err = c.DeleteUserScim(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting UserScim resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserScimRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserScim(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading UserScim resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserScim(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserScim resource from API: %v", err)
	}
	return nil
}

func flattenUserScimBaseUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserScimCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserScimClientAuthenticationMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserScimClientIdentityCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserScimClientSecretToken(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserScimId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserScimName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserScimStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectUserScim(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("base_url", flattenUserScimBaseUrl(o["base-url"], d, "base_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["base-url"], "UserScim-BaseUrl"); ok {
			if err = d.Set("base_url", vv); err != nil {
				return fmt.Errorf("Error reading base_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading base_url: %v", err)
		}
	}

	if err = d.Set("certificate", flattenUserScimCertificate(o["certificate"], d, "certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["certificate"], "UserScim-Certificate"); ok {
			if err = d.Set("certificate", vv); err != nil {
				return fmt.Errorf("Error reading certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("client_authentication_method", flattenUserScimClientAuthenticationMethod(o["client-authentication-method"], d, "client_authentication_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-authentication-method"], "UserScim-ClientAuthenticationMethod"); ok {
			if err = d.Set("client_authentication_method", vv); err != nil {
				return fmt.Errorf("Error reading client_authentication_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_authentication_method: %v", err)
		}
	}

	if err = d.Set("client_identity_check", flattenUserScimClientIdentityCheck(o["client-identity-check"], d, "client_identity_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-identity-check"], "UserScim-ClientIdentityCheck"); ok {
			if err = d.Set("client_identity_check", vv); err != nil {
				return fmt.Errorf("Error reading client_identity_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_identity_check: %v", err)
		}
	}

	if err = d.Set("client_secret_token", flattenUserScimClientSecretToken(o["client-secret-token"], d, "client_secret_token")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-secret-token"], "UserScim-ClientSecretToken"); ok {
			if err = d.Set("client_secret_token", vv); err != nil {
				return fmt.Errorf("Error reading client_secret_token: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_secret_token: %v", err)
		}
	}

	if err = d.Set("fosid", flattenUserScimId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "UserScim-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenUserScimName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "UserScim-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", flattenUserScimStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "UserScim-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenUserScimFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserScimBaseUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserScimCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserScimClientAuthenticationMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserScimClientIdentityCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserScimClientSecretToken(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserScimId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserScimName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserScimStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectUserScim(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("base_url"); ok || d.HasChange("base_url") {
		t, err := expandUserScimBaseUrl(d, v, "base_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["base-url"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok || d.HasChange("certificate") {
		t, err := expandUserScimCertificate(d, v, "certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("client_authentication_method"); ok || d.HasChange("client_authentication_method") {
		t, err := expandUserScimClientAuthenticationMethod(d, v, "client_authentication_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-authentication-method"] = t
		}
	}

	if v, ok := d.GetOk("client_identity_check"); ok || d.HasChange("client_identity_check") {
		t, err := expandUserScimClientIdentityCheck(d, v, "client_identity_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-identity-check"] = t
		}
	}

	if v, ok := d.GetOk("client_secret_token"); ok || d.HasChange("client_secret_token") {
		t, err := expandUserScimClientSecretToken(d, v, "client_secret_token")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-secret-token"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandUserScimId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandUserScimName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandUserScimStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
