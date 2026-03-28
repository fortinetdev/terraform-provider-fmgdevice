// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure virtual-patch profile.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVirtualPatchProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceVirtualPatchProfileCreate,
		Read:   resourceVirtualPatchProfileRead,
		Update: resourceVirtualPatchProfileUpdate,
		Delete: resourceVirtualPatchProfileDelete,

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
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"exemption": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"device": &schema.Schema{
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
						"rule": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"severity": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceVirtualPatchProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVirtualPatchProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating VirtualPatchProfile resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadVirtualPatchProfile(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateVirtualPatchProfile(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating VirtualPatchProfile resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateVirtualPatchProfile(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating VirtualPatchProfile resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceVirtualPatchProfileRead(d, m)
}

func resourceVirtualPatchProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVirtualPatchProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating VirtualPatchProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateVirtualPatchProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VirtualPatchProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceVirtualPatchProfileRead(d, m)
}

func resourceVirtualPatchProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteVirtualPatchProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VirtualPatchProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVirtualPatchProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVirtualPatchProfile(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading VirtualPatchProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVirtualPatchProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VirtualPatchProfile resource from API: %v", err)
	}
	return nil
}

func flattenVirtualPatchProfileAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVirtualPatchProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVirtualPatchProfileExemption(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device"
		if _, ok := i["device"]; ok {
			v := flattenVirtualPatchProfileExemptionDevice(i["device"], d, pre_append)
			tmp["device"] = fortiAPISubPartPatch(v, "VirtualPatchProfile-Exemption-Device")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenVirtualPatchProfileExemptionId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "VirtualPatchProfile-Exemption-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule"
		if _, ok := i["rule"]; ok {
			v := flattenVirtualPatchProfileExemptionRule(i["rule"], d, pre_append)
			tmp["rule"] = fortiAPISubPartPatch(v, "VirtualPatchProfile-Exemption-Rule")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenVirtualPatchProfileExemptionStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "VirtualPatchProfile-Exemption-Status")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenVirtualPatchProfileExemptionDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVirtualPatchProfileExemptionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVirtualPatchProfileExemptionRule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenVirtualPatchProfileExemptionStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVirtualPatchProfileLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVirtualPatchProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVirtualPatchProfileSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectVirtualPatchProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("action", flattenVirtualPatchProfileAction(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "VirtualPatchProfile-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("comment", flattenVirtualPatchProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "VirtualPatchProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("exemption", flattenVirtualPatchProfileExemption(o["exemption"], d, "exemption")); err != nil {
			if vv, ok := fortiAPIPatch(o["exemption"], "VirtualPatchProfile-Exemption"); ok {
				if err = d.Set("exemption", vv); err != nil {
					return fmt.Errorf("Error reading exemption: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading exemption: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("exemption"); ok {
			if err = d.Set("exemption", flattenVirtualPatchProfileExemption(o["exemption"], d, "exemption")); err != nil {
				if vv, ok := fortiAPIPatch(o["exemption"], "VirtualPatchProfile-Exemption"); ok {
					if err = d.Set("exemption", vv); err != nil {
						return fmt.Errorf("Error reading exemption: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading exemption: %v", err)
				}
			}
		}
	}

	if err = d.Set("log", flattenVirtualPatchProfileLog(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "VirtualPatchProfile-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("name", flattenVirtualPatchProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "VirtualPatchProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("severity", flattenVirtualPatchProfileSeverity(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "VirtualPatchProfile-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	return nil
}

func flattenVirtualPatchProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVirtualPatchProfileAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVirtualPatchProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVirtualPatchProfileExemption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["device"], _ = expandVirtualPatchProfileExemptionDevice(d, i["device"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandVirtualPatchProfileExemptionId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rule"], _ = expandVirtualPatchProfileExemptionRule(d, i["rule"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandVirtualPatchProfileExemptionStatus(d, i["status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandVirtualPatchProfileExemptionDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVirtualPatchProfileExemptionId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVirtualPatchProfileExemptionRule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandVirtualPatchProfileExemptionStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVirtualPatchProfileLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVirtualPatchProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVirtualPatchProfileSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectVirtualPatchProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandVirtualPatchProfileAction(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandVirtualPatchProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("exemption"); ok || d.HasChange("exemption") {
		t, err := expandVirtualPatchProfileExemption(d, v, "exemption")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exemption"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandVirtualPatchProfileLog(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandVirtualPatchProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandVirtualPatchProfileSeverity(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	return &obj, nil
}
