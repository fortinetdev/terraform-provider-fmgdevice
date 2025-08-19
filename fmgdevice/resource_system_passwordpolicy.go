// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure password policy for locally defined administrator passwords and IPsec VPN pre-shared keys.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemPasswordPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemPasswordPolicyUpdate,
		Read:   resourceSystemPasswordPolicyRead,
		Update: resourceSystemPasswordPolicyUpdate,
		Delete: resourceSystemPasswordPolicyDelete,

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
			"apply_to": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"change_4_characters": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"expire_day": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"expire_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"login_lockout_upon_downgrade": &schema.Schema{
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
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemPasswordPolicyUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemPasswordPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemPasswordPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemPasswordPolicy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemPasswordPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemPasswordPolicy")

	return resourceSystemPasswordPolicyRead(d, m)
}

func resourceSystemPasswordPolicyDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSystemPasswordPolicy(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemPasswordPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemPasswordPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	paradict["device"] = device_name

	o, err := c.ReadSystemPasswordPolicy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemPasswordPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemPasswordPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemPasswordPolicy resource from API: %v", err)
	}
	return nil
}

func flattenSystemPasswordPolicyApplyTo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemPasswordPolicyChange4Characters(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyExpireDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyExpireStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyLoginLockoutUponDowngrade(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyMinChangeCharacters(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyMinLowerCaseLetter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyMinNonAlphanumeric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyMinNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyMinUpperCaseLetter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyMinimumLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyReusePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyReusePasswordLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemPasswordPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("apply_to", flattenSystemPasswordPolicyApplyTo(o["apply-to"], d, "apply_to")); err != nil {
		if vv, ok := fortiAPIPatch(o["apply-to"], "SystemPasswordPolicy-ApplyTo"); ok {
			if err = d.Set("apply_to", vv); err != nil {
				return fmt.Errorf("Error reading apply_to: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apply_to: %v", err)
		}
	}

	if err = d.Set("change_4_characters", flattenSystemPasswordPolicyChange4Characters(o["change-4-characters"], d, "change_4_characters")); err != nil {
		if vv, ok := fortiAPIPatch(o["change-4-characters"], "SystemPasswordPolicy-Change4Characters"); ok {
			if err = d.Set("change_4_characters", vv); err != nil {
				return fmt.Errorf("Error reading change_4_characters: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading change_4_characters: %v", err)
		}
	}

	if err = d.Set("expire_day", flattenSystemPasswordPolicyExpireDay(o["expire-day"], d, "expire_day")); err != nil {
		if vv, ok := fortiAPIPatch(o["expire-day"], "SystemPasswordPolicy-ExpireDay"); ok {
			if err = d.Set("expire_day", vv); err != nil {
				return fmt.Errorf("Error reading expire_day: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expire_day: %v", err)
		}
	}

	if err = d.Set("expire_status", flattenSystemPasswordPolicyExpireStatus(o["expire-status"], d, "expire_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["expire-status"], "SystemPasswordPolicy-ExpireStatus"); ok {
			if err = d.Set("expire_status", vv); err != nil {
				return fmt.Errorf("Error reading expire_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expire_status: %v", err)
		}
	}

	if err = d.Set("login_lockout_upon_downgrade", flattenSystemPasswordPolicyLoginLockoutUponDowngrade(o["login-lockout-upon-downgrade"], d, "login_lockout_upon_downgrade")); err != nil {
		if vv, ok := fortiAPIPatch(o["login-lockout-upon-downgrade"], "SystemPasswordPolicy-LoginLockoutUponDowngrade"); ok {
			if err = d.Set("login_lockout_upon_downgrade", vv); err != nil {
				return fmt.Errorf("Error reading login_lockout_upon_downgrade: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading login_lockout_upon_downgrade: %v", err)
		}
	}

	if err = d.Set("min_change_characters", flattenSystemPasswordPolicyMinChangeCharacters(o["min-change-characters"], d, "min_change_characters")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-change-characters"], "SystemPasswordPolicy-MinChangeCharacters"); ok {
			if err = d.Set("min_change_characters", vv); err != nil {
				return fmt.Errorf("Error reading min_change_characters: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_change_characters: %v", err)
		}
	}

	if err = d.Set("min_lower_case_letter", flattenSystemPasswordPolicyMinLowerCaseLetter(o["min-lower-case-letter"], d, "min_lower_case_letter")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-lower-case-letter"], "SystemPasswordPolicy-MinLowerCaseLetter"); ok {
			if err = d.Set("min_lower_case_letter", vv); err != nil {
				return fmt.Errorf("Error reading min_lower_case_letter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_lower_case_letter: %v", err)
		}
	}

	if err = d.Set("min_non_alphanumeric", flattenSystemPasswordPolicyMinNonAlphanumeric(o["min-non-alphanumeric"], d, "min_non_alphanumeric")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-non-alphanumeric"], "SystemPasswordPolicy-MinNonAlphanumeric"); ok {
			if err = d.Set("min_non_alphanumeric", vv); err != nil {
				return fmt.Errorf("Error reading min_non_alphanumeric: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_non_alphanumeric: %v", err)
		}
	}

	if err = d.Set("min_number", flattenSystemPasswordPolicyMinNumber(o["min-number"], d, "min_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-number"], "SystemPasswordPolicy-MinNumber"); ok {
			if err = d.Set("min_number", vv); err != nil {
				return fmt.Errorf("Error reading min_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_number: %v", err)
		}
	}

	if err = d.Set("min_upper_case_letter", flattenSystemPasswordPolicyMinUpperCaseLetter(o["min-upper-case-letter"], d, "min_upper_case_letter")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-upper-case-letter"], "SystemPasswordPolicy-MinUpperCaseLetter"); ok {
			if err = d.Set("min_upper_case_letter", vv); err != nil {
				return fmt.Errorf("Error reading min_upper_case_letter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_upper_case_letter: %v", err)
		}
	}

	if err = d.Set("minimum_length", flattenSystemPasswordPolicyMinimumLength(o["minimum-length"], d, "minimum_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["minimum-length"], "SystemPasswordPolicy-MinimumLength"); ok {
			if err = d.Set("minimum_length", vv); err != nil {
				return fmt.Errorf("Error reading minimum_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading minimum_length: %v", err)
		}
	}

	if err = d.Set("reuse_password", flattenSystemPasswordPolicyReusePassword(o["reuse-password"], d, "reuse_password")); err != nil {
		if vv, ok := fortiAPIPatch(o["reuse-password"], "SystemPasswordPolicy-ReusePassword"); ok {
			if err = d.Set("reuse_password", vv); err != nil {
				return fmt.Errorf("Error reading reuse_password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reuse_password: %v", err)
		}
	}

	if err = d.Set("reuse_password_limit", flattenSystemPasswordPolicyReusePasswordLimit(o["reuse-password-limit"], d, "reuse_password_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["reuse-password-limit"], "SystemPasswordPolicy-ReusePasswordLimit"); ok {
			if err = d.Set("reuse_password_limit", vv); err != nil {
				return fmt.Errorf("Error reading reuse_password_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reuse_password_limit: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemPasswordPolicyStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemPasswordPolicy-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSystemPasswordPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemPasswordPolicyApplyTo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemPasswordPolicyChange4Characters(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyExpireDay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyExpireStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyLoginLockoutUponDowngrade(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyMinChangeCharacters(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyMinLowerCaseLetter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyMinNonAlphanumeric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyMinNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyMinUpperCaseLetter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyMinimumLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyReusePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyReusePasswordLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemPasswordPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("apply_to"); ok || d.HasChange("apply_to") {
		t, err := expandSystemPasswordPolicyApplyTo(d, v, "apply_to")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apply-to"] = t
		}
	}

	if v, ok := d.GetOk("change_4_characters"); ok || d.HasChange("change_4_characters") {
		t, err := expandSystemPasswordPolicyChange4Characters(d, v, "change_4_characters")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["change-4-characters"] = t
		}
	}

	if v, ok := d.GetOk("expire_day"); ok || d.HasChange("expire_day") {
		t, err := expandSystemPasswordPolicyExpireDay(d, v, "expire_day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expire-day"] = t
		}
	}

	if v, ok := d.GetOk("expire_status"); ok || d.HasChange("expire_status") {
		t, err := expandSystemPasswordPolicyExpireStatus(d, v, "expire_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expire-status"] = t
		}
	}

	if v, ok := d.GetOk("login_lockout_upon_downgrade"); ok || d.HasChange("login_lockout_upon_downgrade") {
		t, err := expandSystemPasswordPolicyLoginLockoutUponDowngrade(d, v, "login_lockout_upon_downgrade")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-lockout-upon-downgrade"] = t
		}
	}

	if v, ok := d.GetOk("min_change_characters"); ok || d.HasChange("min_change_characters") {
		t, err := expandSystemPasswordPolicyMinChangeCharacters(d, v, "min_change_characters")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-change-characters"] = t
		}
	}

	if v, ok := d.GetOk("min_lower_case_letter"); ok || d.HasChange("min_lower_case_letter") {
		t, err := expandSystemPasswordPolicyMinLowerCaseLetter(d, v, "min_lower_case_letter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-lower-case-letter"] = t
		}
	}

	if v, ok := d.GetOk("min_non_alphanumeric"); ok || d.HasChange("min_non_alphanumeric") {
		t, err := expandSystemPasswordPolicyMinNonAlphanumeric(d, v, "min_non_alphanumeric")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-non-alphanumeric"] = t
		}
	}

	if v, ok := d.GetOk("min_number"); ok || d.HasChange("min_number") {
		t, err := expandSystemPasswordPolicyMinNumber(d, v, "min_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-number"] = t
		}
	}

	if v, ok := d.GetOk("min_upper_case_letter"); ok || d.HasChange("min_upper_case_letter") {
		t, err := expandSystemPasswordPolicyMinUpperCaseLetter(d, v, "min_upper_case_letter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-upper-case-letter"] = t
		}
	}

	if v, ok := d.GetOk("minimum_length"); ok || d.HasChange("minimum_length") {
		t, err := expandSystemPasswordPolicyMinimumLength(d, v, "minimum_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["minimum-length"] = t
		}
	}

	if v, ok := d.GetOk("reuse_password"); ok || d.HasChange("reuse_password") {
		t, err := expandSystemPasswordPolicyReusePassword(d, v, "reuse_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reuse-password"] = t
		}
	}

	if v, ok := d.GetOk("reuse_password_limit"); ok || d.HasChange("reuse_password_limit") {
		t, err := expandSystemPasswordPolicyReusePasswordLimit(d, v, "reuse_password_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reuse-password-limit"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemPasswordPolicyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
