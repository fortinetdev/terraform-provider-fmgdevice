// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Method restriction.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWafProfileMethod() *schema.Resource {
	return &schema.Resource{
		Create: resourceWafProfileMethodUpdate,
		Read:   resourceWafProfileMethodRead,
		Update: resourceWafProfileMethodUpdate,
		Delete: resourceWafProfileMethodDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

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
			"default_allowed_methods": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"method_policy": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"allowed_methods": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"pattern": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"regex": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceWafProfileMethodUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWafProfileMethod(d)
	if err != nil {
		return fmt.Errorf("Error updating WafProfileMethod resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWafProfileMethod(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WafProfileMethod resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WafProfileMethod")

	return resourceWafProfileMethodRead(d, m)
}

func resourceWafProfileMethodDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWafProfileMethod(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WafProfileMethod resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafProfileMethodRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWafProfileMethod(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WafProfileMethod resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafProfileMethod(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WafProfileMethod resource from API: %v", err)
	}
	return nil
}

func flattenWafProfileMethodDefaultAllowedMethods2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWafProfileMethodLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileMethodMethodPolicy2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {
			v := flattenWafProfileMethodMethodPolicyAddress2edl(i["address"], d, pre_append)
			tmp["address"] = fortiAPISubPartPatch(v, "WafProfileMethod-MethodPolicy-Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowed_methods"
		if _, ok := i["allowed-methods"]; ok {
			v := flattenWafProfileMethodMethodPolicyAllowedMethods2edl(i["allowed-methods"], d, pre_append)
			tmp["allowed_methods"] = fortiAPISubPartPatch(v, "WafProfileMethod-MethodPolicy-AllowedMethods")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWafProfileMethodMethodPolicyId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WafProfileMethod-MethodPolicy-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			v := flattenWafProfileMethodMethodPolicyPattern2edl(i["pattern"], d, pre_append)
			tmp["pattern"] = fortiAPISubPartPatch(v, "WafProfileMethod-MethodPolicy-Pattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := i["regex"]; ok {
			v := flattenWafProfileMethodMethodPolicyRegex2edl(i["regex"], d, pre_append)
			tmp["regex"] = fortiAPISubPartPatch(v, "WafProfileMethod-MethodPolicy-Regex")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWafProfileMethodMethodPolicyAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWafProfileMethodMethodPolicyAllowedMethods2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWafProfileMethodMethodPolicyId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileMethodMethodPolicyPattern2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileMethodMethodPolicyRegex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileMethodSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileMethodStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWafProfileMethod(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("default_allowed_methods", flattenWafProfileMethodDefaultAllowedMethods2edl(o["default-allowed-methods"], d, "default_allowed_methods")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-allowed-methods"], "WafProfileMethod-DefaultAllowedMethods"); ok {
			if err = d.Set("default_allowed_methods", vv); err != nil {
				return fmt.Errorf("Error reading default_allowed_methods: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_allowed_methods: %v", err)
		}
	}

	if err = d.Set("log", flattenWafProfileMethodLog2edl(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "WafProfileMethod-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("method_policy", flattenWafProfileMethodMethodPolicy2edl(o["method-policy"], d, "method_policy")); err != nil {
			if vv, ok := fortiAPIPatch(o["method-policy"], "WafProfileMethod-MethodPolicy"); ok {
				if err = d.Set("method_policy", vv); err != nil {
					return fmt.Errorf("Error reading method_policy: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading method_policy: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("method_policy"); ok {
			if err = d.Set("method_policy", flattenWafProfileMethodMethodPolicy2edl(o["method-policy"], d, "method_policy")); err != nil {
				if vv, ok := fortiAPIPatch(o["method-policy"], "WafProfileMethod-MethodPolicy"); ok {
					if err = d.Set("method_policy", vv); err != nil {
						return fmt.Errorf("Error reading method_policy: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading method_policy: %v", err)
				}
			}
		}
	}

	if err = d.Set("severity", flattenWafProfileMethodSeverity2edl(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "WafProfileMethod-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("status", flattenWafProfileMethodStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WafProfileMethod-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenWafProfileMethodFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWafProfileMethodDefaultAllowedMethods2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWafProfileMethodLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileMethodMethodPolicy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["address"], _ = expandWafProfileMethodMethodPolicyAddress2edl(d, i["address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowed_methods"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowed-methods"], _ = expandWafProfileMethodMethodPolicyAllowedMethods2edl(d, i["allowed_methods"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWafProfileMethodMethodPolicyId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern"], _ = expandWafProfileMethodMethodPolicyPattern2edl(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["regex"], _ = expandWafProfileMethodMethodPolicyRegex2edl(d, i["regex"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWafProfileMethodMethodPolicyAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWafProfileMethodMethodPolicyAllowedMethods2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWafProfileMethodMethodPolicyId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileMethodMethodPolicyPattern2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileMethodMethodPolicyRegex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileMethodSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileMethodStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWafProfileMethod(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("default_allowed_methods"); ok || d.HasChange("default_allowed_methods") {
		t, err := expandWafProfileMethodDefaultAllowedMethods2edl(d, v, "default_allowed_methods")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-allowed-methods"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandWafProfileMethodLog2edl(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("method_policy"); ok || d.HasChange("method_policy") {
		t, err := expandWafProfileMethodMethodPolicy2edl(d, v, "method_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["method-policy"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandWafProfileMethodSeverity2edl(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWafProfileMethodStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
