// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure AntiSpam MIME header.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceEmailfilterMheader() *schema.Resource {
	return &schema.Resource{
		Create: resourceEmailfilterMheaderCreate,
		Read:   resourceEmailfilterMheaderRead,
		Update: resourceEmailfilterMheaderUpdate,
		Delete: resourceEmailfilterMheaderDelete,

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
						"fieldbody": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fieldname": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"pattern_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"name": &schema.Schema{
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

func resourceEmailfilterMheaderCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectEmailfilterMheader(d)
	if err != nil {
		return fmt.Errorf("Error creating EmailfilterMheader resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadEmailfilterMheader(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateEmailfilterMheader(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating EmailfilterMheader resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateEmailfilterMheader(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating EmailfilterMheader resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceEmailfilterMheaderRead(d, m)
}

func resourceEmailfilterMheaderUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectEmailfilterMheader(d)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterMheader resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateEmailfilterMheader(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterMheader resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceEmailfilterMheaderRead(d, m)
}

func resourceEmailfilterMheaderDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteEmailfilterMheader(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting EmailfilterMheader resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEmailfilterMheaderRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadEmailfilterMheader(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading EmailfilterMheader resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectEmailfilterMheader(d, o)
	if err != nil {
		return fmt.Errorf("Error reading EmailfilterMheader resource from API: %v", err)
	}
	return nil
}

func flattenEmailfilterMheaderComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterMheaderEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenEmailfilterMheaderEntriesAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "EmailfilterMheader-Entries-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fieldbody"
		if _, ok := i["fieldbody"]; ok {
			v := flattenEmailfilterMheaderEntriesFieldbody(i["fieldbody"], d, pre_append)
			tmp["fieldbody"] = fortiAPISubPartPatch(v, "EmailfilterMheader-Entries-Fieldbody")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fieldname"
		if _, ok := i["fieldname"]; ok {
			v := flattenEmailfilterMheaderEntriesFieldname(i["fieldname"], d, pre_append)
			tmp["fieldname"] = fortiAPISubPartPatch(v, "EmailfilterMheader-Entries-Fieldname")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenEmailfilterMheaderEntriesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "EmailfilterMheader-Entries-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern_type"
		if _, ok := i["pattern-type"]; ok {
			v := flattenEmailfilterMheaderEntriesPatternType(i["pattern-type"], d, pre_append)
			tmp["pattern_type"] = fortiAPISubPartPatch(v, "EmailfilterMheader-Entries-PatternType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenEmailfilterMheaderEntriesStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "EmailfilterMheader-Entries-Status")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenEmailfilterMheaderEntriesAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterMheaderEntriesFieldbody(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterMheaderEntriesFieldname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterMheaderEntriesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterMheaderEntriesPatternType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterMheaderEntriesStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterMheaderId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterMheaderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectEmailfilterMheader(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comment", flattenEmailfilterMheaderComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "EmailfilterMheader-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("entries", flattenEmailfilterMheaderEntries(o["entries"], d, "entries")); err != nil {
			if vv, ok := fortiAPIPatch(o["entries"], "EmailfilterMheader-Entries"); ok {
				if err = d.Set("entries", vv); err != nil {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenEmailfilterMheaderEntries(o["entries"], d, "entries")); err != nil {
				if vv, ok := fortiAPIPatch(o["entries"], "EmailfilterMheader-Entries"); ok {
					if err = d.Set("entries", vv); err != nil {
						return fmt.Errorf("Error reading entries: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			}
		}
	}

	if err = d.Set("fosid", flattenEmailfilterMheaderId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "EmailfilterMheader-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenEmailfilterMheaderName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "EmailfilterMheader-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenEmailfilterMheaderFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandEmailfilterMheaderComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterMheaderEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandEmailfilterMheaderEntriesAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fieldbody"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fieldbody"], _ = expandEmailfilterMheaderEntriesFieldbody(d, i["fieldbody"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fieldname"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fieldname"], _ = expandEmailfilterMheaderEntriesFieldname(d, i["fieldname"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandEmailfilterMheaderEntriesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern-type"], _ = expandEmailfilterMheaderEntriesPatternType(d, i["pattern_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandEmailfilterMheaderEntriesStatus(d, i["status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandEmailfilterMheaderEntriesAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterMheaderEntriesFieldbody(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterMheaderEntriesFieldname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterMheaderEntriesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterMheaderEntriesPatternType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterMheaderEntriesStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterMheaderId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterMheaderName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectEmailfilterMheader(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandEmailfilterMheaderComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("entries"); ok || d.HasChange("entries") {
		t, err := expandEmailfilterMheaderEntries(d, v, "entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandEmailfilterMheaderId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandEmailfilterMheaderName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
