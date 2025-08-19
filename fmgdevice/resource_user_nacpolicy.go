// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure NAC policy matching pattern to identify matching NAC devices.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserNacPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserNacPolicyCreate,
		Read:   resourceUserNacPolicyRead,
		Update: resourceUserNacPolicyUpdate,
		Delete: resourceUserNacPolicyDelete,

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
			"category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ems_tag": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"family": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hw_vendor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hw_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"os": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sw_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"switch_auto_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"switch_fortilink": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"switch_mac_policy": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"switch_port_policy": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"switch_scope": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceUserNacPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

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

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectUserNacPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating UserNacPolicy resource while getting object: %v", err)
	}

	_, err = c.CreateUserNacPolicy(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating UserNacPolicy resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceUserNacPolicyRead(d, m)
}

func resourceUserNacPolicyUpdate(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectUserNacPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating UserNacPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateUserNacPolicy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating UserNacPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceUserNacPolicyRead(d, m)
}

func resourceUserNacPolicyDelete(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteUserNacPolicy(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting UserNacPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserNacPolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserNacPolicy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading UserNacPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserNacPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserNacPolicy resource from API: %v", err)
	}
	return nil
}

func flattenUserNacPolicyCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserNacPolicyDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserNacPolicyEmsTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserNacPolicyFamily(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserNacPolicyHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserNacPolicyHwVendor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserNacPolicyHwVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserNacPolicyMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserNacPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserNacPolicyOs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserNacPolicySrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserNacPolicyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserNacPolicySwVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserNacPolicySwitchAutoAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserNacPolicySwitchFortilink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserNacPolicySwitchMacPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserNacPolicySwitchPortPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserNacPolicySwitchScope(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserNacPolicyType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserNacPolicyUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserNacPolicyUserGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectUserNacPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("category", flattenUserNacPolicyCategory(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "UserNacPolicy-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("description", flattenUserNacPolicyDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "UserNacPolicy-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("ems_tag", flattenUserNacPolicyEmsTag(o["ems-tag"], d, "ems_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["ems-tag"], "UserNacPolicy-EmsTag"); ok {
			if err = d.Set("ems_tag", vv); err != nil {
				return fmt.Errorf("Error reading ems_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ems_tag: %v", err)
		}
	}

	if err = d.Set("family", flattenUserNacPolicyFamily(o["family"], d, "family")); err != nil {
		if vv, ok := fortiAPIPatch(o["family"], "UserNacPolicy-Family"); ok {
			if err = d.Set("family", vv); err != nil {
				return fmt.Errorf("Error reading family: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading family: %v", err)
		}
	}

	if err = d.Set("host", flattenUserNacPolicyHost(o["host"], d, "host")); err != nil {
		if vv, ok := fortiAPIPatch(o["host"], "UserNacPolicy-Host"); ok {
			if err = d.Set("host", vv); err != nil {
				return fmt.Errorf("Error reading host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("hw_vendor", flattenUserNacPolicyHwVendor(o["hw-vendor"], d, "hw_vendor")); err != nil {
		if vv, ok := fortiAPIPatch(o["hw-vendor"], "UserNacPolicy-HwVendor"); ok {
			if err = d.Set("hw_vendor", vv); err != nil {
				return fmt.Errorf("Error reading hw_vendor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hw_vendor: %v", err)
		}
	}

	if err = d.Set("hw_version", flattenUserNacPolicyHwVersion(o["hw-version"], d, "hw_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["hw-version"], "UserNacPolicy-HwVersion"); ok {
			if err = d.Set("hw_version", vv); err != nil {
				return fmt.Errorf("Error reading hw_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hw_version: %v", err)
		}
	}

	if err = d.Set("mac", flattenUserNacPolicyMac(o["mac"], d, "mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac"], "UserNacPolicy-Mac"); ok {
			if err = d.Set("mac", vv); err != nil {
				return fmt.Errorf("Error reading mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac: %v", err)
		}
	}

	if err = d.Set("name", flattenUserNacPolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "UserNacPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("os", flattenUserNacPolicyOs(o["os"], d, "os")); err != nil {
		if vv, ok := fortiAPIPatch(o["os"], "UserNacPolicy-Os"); ok {
			if err = d.Set("os", vv); err != nil {
				return fmt.Errorf("Error reading os: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading os: %v", err)
		}
	}

	if err = d.Set("src", flattenUserNacPolicySrc(o["src"], d, "src")); err != nil {
		if vv, ok := fortiAPIPatch(o["src"], "UserNacPolicy-Src"); ok {
			if err = d.Set("src", vv); err != nil {
				return fmt.Errorf("Error reading src: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src: %v", err)
		}
	}

	if err = d.Set("status", flattenUserNacPolicyStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "UserNacPolicy-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("sw_version", flattenUserNacPolicySwVersion(o["sw-version"], d, "sw_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["sw-version"], "UserNacPolicy-SwVersion"); ok {
			if err = d.Set("sw_version", vv); err != nil {
				return fmt.Errorf("Error reading sw_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sw_version: %v", err)
		}
	}

	if err = d.Set("switch_auto_auth", flattenUserNacPolicySwitchAutoAuth(o["switch-auto-auth"], d, "switch_auto_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-auto-auth"], "UserNacPolicy-SwitchAutoAuth"); ok {
			if err = d.Set("switch_auto_auth", vv); err != nil {
				return fmt.Errorf("Error reading switch_auto_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_auto_auth: %v", err)
		}
	}

	if err = d.Set("switch_fortilink", flattenUserNacPolicySwitchFortilink(o["switch-fortilink"], d, "switch_fortilink")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-fortilink"], "UserNacPolicy-SwitchFortilink"); ok {
			if err = d.Set("switch_fortilink", vv); err != nil {
				return fmt.Errorf("Error reading switch_fortilink: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_fortilink: %v", err)
		}
	}

	if err = d.Set("switch_mac_policy", flattenUserNacPolicySwitchMacPolicy(o["switch-mac-policy"], d, "switch_mac_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-mac-policy"], "UserNacPolicy-SwitchMacPolicy"); ok {
			if err = d.Set("switch_mac_policy", vv); err != nil {
				return fmt.Errorf("Error reading switch_mac_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_mac_policy: %v", err)
		}
	}

	if err = d.Set("switch_port_policy", flattenUserNacPolicySwitchPortPolicy(o["switch-port-policy"], d, "switch_port_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-port-policy"], "UserNacPolicy-SwitchPortPolicy"); ok {
			if err = d.Set("switch_port_policy", vv); err != nil {
				return fmt.Errorf("Error reading switch_port_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_port_policy: %v", err)
		}
	}

	if err = d.Set("switch_scope", flattenUserNacPolicySwitchScope(o["switch-scope"], d, "switch_scope")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-scope"], "UserNacPolicy-SwitchScope"); ok {
			if err = d.Set("switch_scope", vv); err != nil {
				return fmt.Errorf("Error reading switch_scope: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_scope: %v", err)
		}
	}

	if err = d.Set("type", flattenUserNacPolicyType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "UserNacPolicy-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("user", flattenUserNacPolicyUser(o["user"], d, "user")); err != nil {
		if vv, ok := fortiAPIPatch(o["user"], "UserNacPolicy-User"); ok {
			if err = d.Set("user", vv); err != nil {
				return fmt.Errorf("Error reading user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	if err = d.Set("user_group", flattenUserNacPolicyUserGroup(o["user-group"], d, "user_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-group"], "UserNacPolicy-UserGroup"); ok {
			if err = d.Set("user_group", vv); err != nil {
				return fmt.Errorf("Error reading user_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_group: %v", err)
		}
	}

	return nil
}

func flattenUserNacPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserNacPolicyCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicyDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicyEmsTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserNacPolicyFamily(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicyHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicyHwVendor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicyHwVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicyMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicyOs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicySrc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicySwVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicySwitchAutoAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicySwitchFortilink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserNacPolicySwitchMacPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserNacPolicySwitchPortPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserNacPolicySwitchScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserNacPolicyType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicyUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicyUserGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectUserNacPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandUserNacPolicyCategory(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandUserNacPolicyDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("ems_tag"); ok || d.HasChange("ems_tag") {
		t, err := expandUserNacPolicyEmsTag(d, v, "ems_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ems-tag"] = t
		}
	}

	if v, ok := d.GetOk("family"); ok || d.HasChange("family") {
		t, err := expandUserNacPolicyFamily(d, v, "family")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["family"] = t
		}
	}

	if v, ok := d.GetOk("host"); ok || d.HasChange("host") {
		t, err := expandUserNacPolicyHost(d, v, "host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}

	if v, ok := d.GetOk("hw_vendor"); ok || d.HasChange("hw_vendor") {
		t, err := expandUserNacPolicyHwVendor(d, v, "hw_vendor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hw-vendor"] = t
		}
	}

	if v, ok := d.GetOk("hw_version"); ok || d.HasChange("hw_version") {
		t, err := expandUserNacPolicyHwVersion(d, v, "hw_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hw-version"] = t
		}
	}

	if v, ok := d.GetOk("mac"); ok || d.HasChange("mac") {
		t, err := expandUserNacPolicyMac(d, v, "mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandUserNacPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("os"); ok || d.HasChange("os") {
		t, err := expandUserNacPolicyOs(d, v, "os")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os"] = t
		}
	}

	if v, ok := d.GetOk("src"); ok || d.HasChange("src") {
		t, err := expandUserNacPolicySrc(d, v, "src")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandUserNacPolicyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("sw_version"); ok || d.HasChange("sw_version") {
		t, err := expandUserNacPolicySwVersion(d, v, "sw_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sw-version"] = t
		}
	}

	if v, ok := d.GetOk("switch_auto_auth"); ok || d.HasChange("switch_auto_auth") {
		t, err := expandUserNacPolicySwitchAutoAuth(d, v, "switch_auto_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-auto-auth"] = t
		}
	}

	if v, ok := d.GetOk("switch_fortilink"); ok || d.HasChange("switch_fortilink") {
		t, err := expandUserNacPolicySwitchFortilink(d, v, "switch_fortilink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-fortilink"] = t
		}
	}

	if v, ok := d.GetOk("switch_mac_policy"); ok || d.HasChange("switch_mac_policy") {
		t, err := expandUserNacPolicySwitchMacPolicy(d, v, "switch_mac_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-mac-policy"] = t
		}
	}

	if v, ok := d.GetOk("switch_port_policy"); ok || d.HasChange("switch_port_policy") {
		t, err := expandUserNacPolicySwitchPortPolicy(d, v, "switch_port_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-port-policy"] = t
		}
	}

	if v, ok := d.GetOk("switch_scope"); ok || d.HasChange("switch_scope") {
		t, err := expandUserNacPolicySwitchScope(d, v, "switch_scope")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-scope"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandUserNacPolicyType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok || d.HasChange("user") {
		t, err := expandUserNacPolicyUser(d, v, "user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	if v, ok := d.GetOk("user_group"); ok || d.HasChange("user_group") {
		t, err := expandUserNacPolicyUserGroup(d, v, "user_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-group"] = t
		}
	}

	return &obj, nil
}
