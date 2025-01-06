// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure the password policy for guest administrators.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemPasswordPolicyGuestAdmin() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemPasswordPolicyGuestAdminUpdate,
		Read:   resourceSystemPasswordPolicyGuestAdminRead,
		Update: resourceSystemPasswordPolicyGuestAdminUpdate,
		Delete: resourceSystemPasswordPolicyGuestAdminDelete,

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

func resourceSystemPasswordPolicyGuestAdminUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemPasswordPolicyGuestAdmin(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemPasswordPolicyGuestAdmin resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemPasswordPolicyGuestAdmin(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemPasswordPolicyGuestAdmin resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemPasswordPolicyGuestAdmin")

	return resourceSystemPasswordPolicyGuestAdminRead(d, m)
}

func resourceSystemPasswordPolicyGuestAdminDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	err = c.DeleteSystemPasswordPolicyGuestAdmin(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemPasswordPolicyGuestAdmin resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemPasswordPolicyGuestAdminRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemPasswordPolicyGuestAdmin(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemPasswordPolicyGuestAdmin resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemPasswordPolicyGuestAdmin(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemPasswordPolicyGuestAdmin resource from API: %v", err)
	}
	return nil
}

func flattenSystemPasswordPolicyGuestAdminApplyTo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemPasswordPolicyGuestAdminChange4Characters(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyGuestAdminExpireDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyGuestAdminExpireStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyGuestAdminMinChangeCharacters(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyGuestAdminMinLowerCaseLetter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyGuestAdminMinNonAlphanumeric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyGuestAdminMinNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyGuestAdminMinUpperCaseLetter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyGuestAdminMinimumLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyGuestAdminReusePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyGuestAdminReusePasswordLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyGuestAdminStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemPasswordPolicyGuestAdmin(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("apply_to", flattenSystemPasswordPolicyGuestAdminApplyTo(o["apply-to"], d, "apply_to")); err != nil {
		if vv, ok := fortiAPIPatch(o["apply-to"], "SystemPasswordPolicyGuestAdmin-ApplyTo"); ok {
			if err = d.Set("apply_to", vv); err != nil {
				return fmt.Errorf("Error reading apply_to: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apply_to: %v", err)
		}
	}

	if err = d.Set("change_4_characters", flattenSystemPasswordPolicyGuestAdminChange4Characters(o["change-4-characters"], d, "change_4_characters")); err != nil {
		if vv, ok := fortiAPIPatch(o["change-4-characters"], "SystemPasswordPolicyGuestAdmin-Change4Characters"); ok {
			if err = d.Set("change_4_characters", vv); err != nil {
				return fmt.Errorf("Error reading change_4_characters: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading change_4_characters: %v", err)
		}
	}

	if err = d.Set("expire_day", flattenSystemPasswordPolicyGuestAdminExpireDay(o["expire-day"], d, "expire_day")); err != nil {
		if vv, ok := fortiAPIPatch(o["expire-day"], "SystemPasswordPolicyGuestAdmin-ExpireDay"); ok {
			if err = d.Set("expire_day", vv); err != nil {
				return fmt.Errorf("Error reading expire_day: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expire_day: %v", err)
		}
	}

	if err = d.Set("expire_status", flattenSystemPasswordPolicyGuestAdminExpireStatus(o["expire-status"], d, "expire_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["expire-status"], "SystemPasswordPolicyGuestAdmin-ExpireStatus"); ok {
			if err = d.Set("expire_status", vv); err != nil {
				return fmt.Errorf("Error reading expire_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expire_status: %v", err)
		}
	}

	if err = d.Set("min_change_characters", flattenSystemPasswordPolicyGuestAdminMinChangeCharacters(o["min-change-characters"], d, "min_change_characters")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-change-characters"], "SystemPasswordPolicyGuestAdmin-MinChangeCharacters"); ok {
			if err = d.Set("min_change_characters", vv); err != nil {
				return fmt.Errorf("Error reading min_change_characters: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_change_characters: %v", err)
		}
	}

	if err = d.Set("min_lower_case_letter", flattenSystemPasswordPolicyGuestAdminMinLowerCaseLetter(o["min-lower-case-letter"], d, "min_lower_case_letter")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-lower-case-letter"], "SystemPasswordPolicyGuestAdmin-MinLowerCaseLetter"); ok {
			if err = d.Set("min_lower_case_letter", vv); err != nil {
				return fmt.Errorf("Error reading min_lower_case_letter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_lower_case_letter: %v", err)
		}
	}

	if err = d.Set("min_non_alphanumeric", flattenSystemPasswordPolicyGuestAdminMinNonAlphanumeric(o["min-non-alphanumeric"], d, "min_non_alphanumeric")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-non-alphanumeric"], "SystemPasswordPolicyGuestAdmin-MinNonAlphanumeric"); ok {
			if err = d.Set("min_non_alphanumeric", vv); err != nil {
				return fmt.Errorf("Error reading min_non_alphanumeric: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_non_alphanumeric: %v", err)
		}
	}

	if err = d.Set("min_number", flattenSystemPasswordPolicyGuestAdminMinNumber(o["min-number"], d, "min_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-number"], "SystemPasswordPolicyGuestAdmin-MinNumber"); ok {
			if err = d.Set("min_number", vv); err != nil {
				return fmt.Errorf("Error reading min_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_number: %v", err)
		}
	}

	if err = d.Set("min_upper_case_letter", flattenSystemPasswordPolicyGuestAdminMinUpperCaseLetter(o["min-upper-case-letter"], d, "min_upper_case_letter")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-upper-case-letter"], "SystemPasswordPolicyGuestAdmin-MinUpperCaseLetter"); ok {
			if err = d.Set("min_upper_case_letter", vv); err != nil {
				return fmt.Errorf("Error reading min_upper_case_letter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_upper_case_letter: %v", err)
		}
	}

	if err = d.Set("minimum_length", flattenSystemPasswordPolicyGuestAdminMinimumLength(o["minimum-length"], d, "minimum_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["minimum-length"], "SystemPasswordPolicyGuestAdmin-MinimumLength"); ok {
			if err = d.Set("minimum_length", vv); err != nil {
				return fmt.Errorf("Error reading minimum_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading minimum_length: %v", err)
		}
	}

	if err = d.Set("reuse_password", flattenSystemPasswordPolicyGuestAdminReusePassword(o["reuse-password"], d, "reuse_password")); err != nil {
		if vv, ok := fortiAPIPatch(o["reuse-password"], "SystemPasswordPolicyGuestAdmin-ReusePassword"); ok {
			if err = d.Set("reuse_password", vv); err != nil {
				return fmt.Errorf("Error reading reuse_password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reuse_password: %v", err)
		}
	}

	if err = d.Set("reuse_password_limit", flattenSystemPasswordPolicyGuestAdminReusePasswordLimit(o["reuse-password-limit"], d, "reuse_password_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["reuse-password-limit"], "SystemPasswordPolicyGuestAdmin-ReusePasswordLimit"); ok {
			if err = d.Set("reuse_password_limit", vv); err != nil {
				return fmt.Errorf("Error reading reuse_password_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reuse_password_limit: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemPasswordPolicyGuestAdminStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemPasswordPolicyGuestAdmin-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSystemPasswordPolicyGuestAdminFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemPasswordPolicyGuestAdminApplyTo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemPasswordPolicyGuestAdminChange4Characters(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyGuestAdminExpireDay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyGuestAdminExpireStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyGuestAdminMinChangeCharacters(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyGuestAdminMinLowerCaseLetter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyGuestAdminMinNonAlphanumeric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyGuestAdminMinNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyGuestAdminMinUpperCaseLetter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyGuestAdminMinimumLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyGuestAdminReusePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyGuestAdminReusePasswordLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyGuestAdminStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemPasswordPolicyGuestAdmin(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("apply_to"); ok || d.HasChange("apply_to") {
		t, err := expandSystemPasswordPolicyGuestAdminApplyTo(d, v, "apply_to")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apply-to"] = t
		}
	}

	if v, ok := d.GetOk("change_4_characters"); ok || d.HasChange("change_4_characters") {
		t, err := expandSystemPasswordPolicyGuestAdminChange4Characters(d, v, "change_4_characters")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["change-4-characters"] = t
		}
	}

	if v, ok := d.GetOk("expire_day"); ok || d.HasChange("expire_day") {
		t, err := expandSystemPasswordPolicyGuestAdminExpireDay(d, v, "expire_day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expire-day"] = t
		}
	}

	if v, ok := d.GetOk("expire_status"); ok || d.HasChange("expire_status") {
		t, err := expandSystemPasswordPolicyGuestAdminExpireStatus(d, v, "expire_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expire-status"] = t
		}
	}

	if v, ok := d.GetOk("min_change_characters"); ok || d.HasChange("min_change_characters") {
		t, err := expandSystemPasswordPolicyGuestAdminMinChangeCharacters(d, v, "min_change_characters")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-change-characters"] = t
		}
	}

	if v, ok := d.GetOk("min_lower_case_letter"); ok || d.HasChange("min_lower_case_letter") {
		t, err := expandSystemPasswordPolicyGuestAdminMinLowerCaseLetter(d, v, "min_lower_case_letter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-lower-case-letter"] = t
		}
	}

	if v, ok := d.GetOk("min_non_alphanumeric"); ok || d.HasChange("min_non_alphanumeric") {
		t, err := expandSystemPasswordPolicyGuestAdminMinNonAlphanumeric(d, v, "min_non_alphanumeric")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-non-alphanumeric"] = t
		}
	}

	if v, ok := d.GetOk("min_number"); ok || d.HasChange("min_number") {
		t, err := expandSystemPasswordPolicyGuestAdminMinNumber(d, v, "min_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-number"] = t
		}
	}

	if v, ok := d.GetOk("min_upper_case_letter"); ok || d.HasChange("min_upper_case_letter") {
		t, err := expandSystemPasswordPolicyGuestAdminMinUpperCaseLetter(d, v, "min_upper_case_letter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-upper-case-letter"] = t
		}
	}

	if v, ok := d.GetOk("minimum_length"); ok || d.HasChange("minimum_length") {
		t, err := expandSystemPasswordPolicyGuestAdminMinimumLength(d, v, "minimum_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["minimum-length"] = t
		}
	}

	if v, ok := d.GetOk("reuse_password"); ok || d.HasChange("reuse_password") {
		t, err := expandSystemPasswordPolicyGuestAdminReusePassword(d, v, "reuse_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reuse-password"] = t
		}
	}

	if v, ok := d.GetOk("reuse_password_limit"); ok || d.HasChange("reuse_password_limit") {
		t, err := expandSystemPasswordPolicyGuestAdminReusePasswordLimit(d, v, "reuse_password_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reuse-password-limit"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemPasswordPolicyGuestAdminStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
