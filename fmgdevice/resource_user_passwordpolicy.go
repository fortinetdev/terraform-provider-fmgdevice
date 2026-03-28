// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure user password policy.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserPasswordPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserPasswordPolicyCreate,
		Read:   resourceUserPasswordPolicyRead,
		Update: resourceUserPasswordPolicyUpdate,
		Delete: resourceUserPasswordPolicyDelete,

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
			"expire_days": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"expire_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"expired_password_renewal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"min_change_characters": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"min_lower_case_letter": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"min_non_alphanumeric": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"min_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"min_upper_case_letter": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"minimum_length": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"reuse_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reuse_password_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"warn_days": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceUserPasswordPolicyCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectUserPasswordPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating UserPasswordPolicy resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadUserPasswordPolicy(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateUserPasswordPolicy(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating UserPasswordPolicy resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateUserPasswordPolicy(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating UserPasswordPolicy resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceUserPasswordPolicyRead(d, m)
}

func resourceUserPasswordPolicyUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectUserPasswordPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating UserPasswordPolicy resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateUserPasswordPolicy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating UserPasswordPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceUserPasswordPolicyRead(d, m)
}

func resourceUserPasswordPolicyDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	wsParams["adom"] = adomv

	err = c.DeleteUserPasswordPolicy(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting UserPasswordPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserPasswordPolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserPasswordPolicy(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading UserPasswordPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserPasswordPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserPasswordPolicy resource from API: %v", err)
	}
	return nil
}

func flattenUserPasswordPolicyExpireDays(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPasswordPolicyExpireStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPasswordPolicyExpiredPasswordRenewal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPasswordPolicyMinChangeCharacters(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPasswordPolicyMinLowerCaseLetter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPasswordPolicyMinNonAlphanumeric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPasswordPolicyMinNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPasswordPolicyMinUpperCaseLetter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPasswordPolicyMinimumLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPasswordPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPasswordPolicyReusePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPasswordPolicyReusePasswordLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPasswordPolicyWarnDays(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectUserPasswordPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("expire_days", flattenUserPasswordPolicyExpireDays(o["expire-days"], d, "expire_days")); err != nil {
		if vv, ok := fortiAPIPatch(o["expire-days"], "UserPasswordPolicy-ExpireDays"); ok {
			if err = d.Set("expire_days", vv); err != nil {
				return fmt.Errorf("Error reading expire_days: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expire_days: %v", err)
		}
	}

	if err = d.Set("expire_status", flattenUserPasswordPolicyExpireStatus(o["expire-status"], d, "expire_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["expire-status"], "UserPasswordPolicy-ExpireStatus"); ok {
			if err = d.Set("expire_status", vv); err != nil {
				return fmt.Errorf("Error reading expire_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expire_status: %v", err)
		}
	}

	if err = d.Set("expired_password_renewal", flattenUserPasswordPolicyExpiredPasswordRenewal(o["expired-password-renewal"], d, "expired_password_renewal")); err != nil {
		if vv, ok := fortiAPIPatch(o["expired-password-renewal"], "UserPasswordPolicy-ExpiredPasswordRenewal"); ok {
			if err = d.Set("expired_password_renewal", vv); err != nil {
				return fmt.Errorf("Error reading expired_password_renewal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expired_password_renewal: %v", err)
		}
	}

	if err = d.Set("min_change_characters", flattenUserPasswordPolicyMinChangeCharacters(o["min-change-characters"], d, "min_change_characters")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-change-characters"], "UserPasswordPolicy-MinChangeCharacters"); ok {
			if err = d.Set("min_change_characters", vv); err != nil {
				return fmt.Errorf("Error reading min_change_characters: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_change_characters: %v", err)
		}
	}

	if err = d.Set("min_lower_case_letter", flattenUserPasswordPolicyMinLowerCaseLetter(o["min-lower-case-letter"], d, "min_lower_case_letter")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-lower-case-letter"], "UserPasswordPolicy-MinLowerCaseLetter"); ok {
			if err = d.Set("min_lower_case_letter", vv); err != nil {
				return fmt.Errorf("Error reading min_lower_case_letter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_lower_case_letter: %v", err)
		}
	}

	if err = d.Set("min_non_alphanumeric", flattenUserPasswordPolicyMinNonAlphanumeric(o["min-non-alphanumeric"], d, "min_non_alphanumeric")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-non-alphanumeric"], "UserPasswordPolicy-MinNonAlphanumeric"); ok {
			if err = d.Set("min_non_alphanumeric", vv); err != nil {
				return fmt.Errorf("Error reading min_non_alphanumeric: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_non_alphanumeric: %v", err)
		}
	}

	if err = d.Set("min_number", flattenUserPasswordPolicyMinNumber(o["min-number"], d, "min_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-number"], "UserPasswordPolicy-MinNumber"); ok {
			if err = d.Set("min_number", vv); err != nil {
				return fmt.Errorf("Error reading min_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_number: %v", err)
		}
	}

	if err = d.Set("min_upper_case_letter", flattenUserPasswordPolicyMinUpperCaseLetter(o["min-upper-case-letter"], d, "min_upper_case_letter")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-upper-case-letter"], "UserPasswordPolicy-MinUpperCaseLetter"); ok {
			if err = d.Set("min_upper_case_letter", vv); err != nil {
				return fmt.Errorf("Error reading min_upper_case_letter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_upper_case_letter: %v", err)
		}
	}

	if err = d.Set("minimum_length", flattenUserPasswordPolicyMinimumLength(o["minimum-length"], d, "minimum_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["minimum-length"], "UserPasswordPolicy-MinimumLength"); ok {
			if err = d.Set("minimum_length", vv); err != nil {
				return fmt.Errorf("Error reading minimum_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading minimum_length: %v", err)
		}
	}

	if err = d.Set("name", flattenUserPasswordPolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "UserPasswordPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("reuse_password", flattenUserPasswordPolicyReusePassword(o["reuse-password"], d, "reuse_password")); err != nil {
		if vv, ok := fortiAPIPatch(o["reuse-password"], "UserPasswordPolicy-ReusePassword"); ok {
			if err = d.Set("reuse_password", vv); err != nil {
				return fmt.Errorf("Error reading reuse_password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reuse_password: %v", err)
		}
	}

	if err = d.Set("reuse_password_limit", flattenUserPasswordPolicyReusePasswordLimit(o["reuse-password-limit"], d, "reuse_password_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["reuse-password-limit"], "UserPasswordPolicy-ReusePasswordLimit"); ok {
			if err = d.Set("reuse_password_limit", vv); err != nil {
				return fmt.Errorf("Error reading reuse_password_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reuse_password_limit: %v", err)
		}
	}

	if err = d.Set("warn_days", flattenUserPasswordPolicyWarnDays(o["warn-days"], d, "warn_days")); err != nil {
		if vv, ok := fortiAPIPatch(o["warn-days"], "UserPasswordPolicy-WarnDays"); ok {
			if err = d.Set("warn_days", vv); err != nil {
				return fmt.Errorf("Error reading warn_days: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading warn_days: %v", err)
		}
	}

	return nil
}

func flattenUserPasswordPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserPasswordPolicyExpireDays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPasswordPolicyExpireStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPasswordPolicyExpiredPasswordRenewal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPasswordPolicyMinChangeCharacters(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPasswordPolicyMinLowerCaseLetter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPasswordPolicyMinNonAlphanumeric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPasswordPolicyMinNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPasswordPolicyMinUpperCaseLetter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPasswordPolicyMinimumLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPasswordPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPasswordPolicyReusePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPasswordPolicyReusePasswordLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPasswordPolicyWarnDays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectUserPasswordPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("expire_days"); ok || d.HasChange("expire_days") {
		t, err := expandUserPasswordPolicyExpireDays(d, v, "expire_days")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expire-days"] = t
		}
	}

	if v, ok := d.GetOk("expire_status"); ok || d.HasChange("expire_status") {
		t, err := expandUserPasswordPolicyExpireStatus(d, v, "expire_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expire-status"] = t
		}
	}

	if v, ok := d.GetOk("expired_password_renewal"); ok || d.HasChange("expired_password_renewal") {
		t, err := expandUserPasswordPolicyExpiredPasswordRenewal(d, v, "expired_password_renewal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expired-password-renewal"] = t
		}
	}

	if v, ok := d.GetOk("min_change_characters"); ok || d.HasChange("min_change_characters") {
		t, err := expandUserPasswordPolicyMinChangeCharacters(d, v, "min_change_characters")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-change-characters"] = t
		}
	}

	if v, ok := d.GetOk("min_lower_case_letter"); ok || d.HasChange("min_lower_case_letter") {
		t, err := expandUserPasswordPolicyMinLowerCaseLetter(d, v, "min_lower_case_letter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-lower-case-letter"] = t
		}
	}

	if v, ok := d.GetOk("min_non_alphanumeric"); ok || d.HasChange("min_non_alphanumeric") {
		t, err := expandUserPasswordPolicyMinNonAlphanumeric(d, v, "min_non_alphanumeric")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-non-alphanumeric"] = t
		}
	}

	if v, ok := d.GetOk("min_number"); ok || d.HasChange("min_number") {
		t, err := expandUserPasswordPolicyMinNumber(d, v, "min_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-number"] = t
		}
	}

	if v, ok := d.GetOk("min_upper_case_letter"); ok || d.HasChange("min_upper_case_letter") {
		t, err := expandUserPasswordPolicyMinUpperCaseLetter(d, v, "min_upper_case_letter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-upper-case-letter"] = t
		}
	}

	if v, ok := d.GetOk("minimum_length"); ok || d.HasChange("minimum_length") {
		t, err := expandUserPasswordPolicyMinimumLength(d, v, "minimum_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["minimum-length"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandUserPasswordPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("reuse_password"); ok || d.HasChange("reuse_password") {
		t, err := expandUserPasswordPolicyReusePassword(d, v, "reuse_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reuse-password"] = t
		}
	}

	if v, ok := d.GetOk("reuse_password_limit"); ok || d.HasChange("reuse_password_limit") {
		t, err := expandUserPasswordPolicyReusePasswordLimit(d, v, "reuse_password_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reuse-password-limit"] = t
		}
	}

	if v, ok := d.GetOk("warn_days"); ok || d.HasChange("warn_days") {
		t, err := expandUserPasswordPolicyWarnDays(d, v, "warn_days")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["warn-days"] = t
		}
	}

	return &obj, nil
}
