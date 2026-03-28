// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i>

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIsolatorProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceIsolatorProfileCreate,
		Read:   resourceIsolatorProfileRead,
		Update: resourceIsolatorProfileUpdate,
		Delete: resourceIsolatorProfileDelete,

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
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"disclaimer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"entries": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"copy_paste": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"proxy_address": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"right_click": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
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

func resourceIsolatorProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectIsolatorProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating IsolatorProfile resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadIsolatorProfile(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateIsolatorProfile(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating IsolatorProfile resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateIsolatorProfile(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating IsolatorProfile resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceIsolatorProfileRead(d, m)
}

func resourceIsolatorProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectIsolatorProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating IsolatorProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateIsolatorProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating IsolatorProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceIsolatorProfileRead(d, m)
}

func resourceIsolatorProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteIsolatorProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting IsolatorProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIsolatorProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadIsolatorProfile(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading IsolatorProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIsolatorProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading IsolatorProfile resource from API: %v", err)
	}
	return nil
}

func flattenIsolatorProfileComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIsolatorProfileDisclaimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIsolatorProfileEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenIsolatorProfileEntriesAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "IsolatorProfile-Entries-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "copy_paste"
		if _, ok := i["copy-paste"]; ok {
			v := flattenIsolatorProfileEntriesCopyPaste(i["copy-paste"], d, pre_append)
			tmp["copy_paste"] = fortiAPISubPartPatch(v, "IsolatorProfile-Entries-CopyPaste")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenIsolatorProfileEntriesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "IsolatorProfile-Entries-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "proxy_address"
		if _, ok := i["proxy-address"]; ok {
			v := flattenIsolatorProfileEntriesProxyAddress(i["proxy-address"], d, pre_append)
			tmp["proxy_address"] = fortiAPISubPartPatch(v, "IsolatorProfile-Entries-ProxyAddress")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "right_click"
		if _, ok := i["right-click"]; ok {
			v := flattenIsolatorProfileEntriesRightClick(i["right-click"], d, pre_append)
			tmp["right_click"] = fortiAPISubPartPatch(v, "IsolatorProfile-Entries-RightClick")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenIsolatorProfileEntriesStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "IsolatorProfile-Entries-Status")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenIsolatorProfileEntriesAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIsolatorProfileEntriesCopyPaste(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIsolatorProfileEntriesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIsolatorProfileEntriesProxyAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIsolatorProfileEntriesRightClick(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIsolatorProfileEntriesStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIsolatorProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectIsolatorProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comments", flattenIsolatorProfileComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "IsolatorProfile-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("disclaimer", flattenIsolatorProfileDisclaimer(o["disclaimer"], d, "disclaimer")); err != nil {
		if vv, ok := fortiAPIPatch(o["disclaimer"], "IsolatorProfile-Disclaimer"); ok {
			if err = d.Set("disclaimer", vv); err != nil {
				return fmt.Errorf("Error reading disclaimer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading disclaimer: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("entries", flattenIsolatorProfileEntries(o["entries"], d, "entries")); err != nil {
			if vv, ok := fortiAPIPatch(o["entries"], "IsolatorProfile-Entries"); ok {
				if err = d.Set("entries", vv); err != nil {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenIsolatorProfileEntries(o["entries"], d, "entries")); err != nil {
				if vv, ok := fortiAPIPatch(o["entries"], "IsolatorProfile-Entries"); ok {
					if err = d.Set("entries", vv); err != nil {
						return fmt.Errorf("Error reading entries: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenIsolatorProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "IsolatorProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenIsolatorProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandIsolatorProfileComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIsolatorProfileDisclaimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIsolatorProfileEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandIsolatorProfileEntriesAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "copy_paste"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["copy-paste"], _ = expandIsolatorProfileEntriesCopyPaste(d, i["copy_paste"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandIsolatorProfileEntriesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "proxy_address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["proxy-address"], _ = expandIsolatorProfileEntriesProxyAddress(d, i["proxy_address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "right_click"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["right-click"], _ = expandIsolatorProfileEntriesRightClick(d, i["right_click"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandIsolatorProfileEntriesStatus(d, i["status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandIsolatorProfileEntriesAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIsolatorProfileEntriesCopyPaste(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIsolatorProfileEntriesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIsolatorProfileEntriesProxyAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIsolatorProfileEntriesRightClick(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIsolatorProfileEntriesStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIsolatorProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectIsolatorProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandIsolatorProfileComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("disclaimer"); ok || d.HasChange("disclaimer") {
		t, err := expandIsolatorProfileDisclaimer(d, v, "disclaimer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disclaimer"] = t
		}
	}

	if v, ok := d.GetOk("entries"); ok || d.HasChange("entries") {
		t, err := expandIsolatorProfileEntries(d, v, "entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandIsolatorProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
