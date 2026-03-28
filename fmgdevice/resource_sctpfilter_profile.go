// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure SCTP filter profiles.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSctpFilterProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceSctpFilterProfileCreate,
		Read:   resourceSctpFilterProfileRead,
		Update: resourceSctpFilterProfileUpdate,
		Delete: resourceSctpFilterProfileDelete,

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
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"ppid_filters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ppid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSctpFilterProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSctpFilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating SctpFilterProfile resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSctpFilterProfile(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSctpFilterProfile(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SctpFilterProfile resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSctpFilterProfile(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SctpFilterProfile resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceSctpFilterProfileRead(d, m)
}

func resourceSctpFilterProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSctpFilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating SctpFilterProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSctpFilterProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SctpFilterProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSctpFilterProfileRead(d, m)
}

func resourceSctpFilterProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSctpFilterProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SctpFilterProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSctpFilterProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSctpFilterProfile(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SctpFilterProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSctpFilterProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SctpFilterProfile resource from API: %v", err)
	}
	return nil
}

func flattenSctpFilterProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSctpFilterProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSctpFilterProfilePpidFilters(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSctpFilterProfilePpidFiltersAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "SctpFilterProfile-PpidFilters-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenSctpFilterProfilePpidFiltersComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "SctpFilterProfile-PpidFilters-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSctpFilterProfilePpidFiltersId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SctpFilterProfile-PpidFilters-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ppid"
		if _, ok := i["ppid"]; ok {
			v := flattenSctpFilterProfilePpidFiltersPpid(i["ppid"], d, pre_append)
			tmp["ppid"] = fortiAPISubPartPatch(v, "SctpFilterProfile-PpidFilters-Ppid")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSctpFilterProfilePpidFiltersAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSctpFilterProfilePpidFiltersComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSctpFilterProfilePpidFiltersId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSctpFilterProfilePpidFiltersPpid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSctpFilterProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comment", flattenSctpFilterProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "SctpFilterProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("name", flattenSctpFilterProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SctpFilterProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ppid_filters", flattenSctpFilterProfilePpidFilters(o["ppid-filters"], d, "ppid_filters")); err != nil {
			if vv, ok := fortiAPIPatch(o["ppid-filters"], "SctpFilterProfile-PpidFilters"); ok {
				if err = d.Set("ppid_filters", vv); err != nil {
					return fmt.Errorf("Error reading ppid_filters: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ppid_filters: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ppid_filters"); ok {
			if err = d.Set("ppid_filters", flattenSctpFilterProfilePpidFilters(o["ppid-filters"], d, "ppid_filters")); err != nil {
				if vv, ok := fortiAPIPatch(o["ppid-filters"], "SctpFilterProfile-PpidFilters"); ok {
					if err = d.Set("ppid_filters", vv); err != nil {
						return fmt.Errorf("Error reading ppid_filters: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ppid_filters: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSctpFilterProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSctpFilterProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSctpFilterProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSctpFilterProfilePpidFilters(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandSctpFilterProfilePpidFiltersAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandSctpFilterProfilePpidFiltersComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSctpFilterProfilePpidFiltersId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ppid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ppid"], _ = expandSctpFilterProfilePpidFiltersPpid(d, i["ppid"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSctpFilterProfilePpidFiltersAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSctpFilterProfilePpidFiltersComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSctpFilterProfilePpidFiltersId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSctpFilterProfilePpidFiltersPpid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSctpFilterProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandSctpFilterProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSctpFilterProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("ppid_filters"); ok || d.HasChange("ppid_filters") {
		t, err := expandSctpFilterProfilePpidFilters(d, v, "ppid_filters")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppid-filters"] = t
		}
	}

	return &obj, nil
}
