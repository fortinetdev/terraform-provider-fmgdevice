// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure dictionaries used by DLP blocking.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDlpDictionary() *schema.Resource {
	return &schema.Resource{
		Create: resourceDlpDictionaryCreate,
		Read:   resourceDlpDictionaryRead,
		Update: resourceDlpDictionaryUpdate,
		Delete: resourceDlpDictionaryDelete,

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
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ignore_case": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pattern": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"repeat": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"fgd_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"match_around": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"match_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceDlpDictionaryCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDlpDictionary(d)
	if err != nil {
		return fmt.Errorf("Error creating DlpDictionary resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadDlpDictionary(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateDlpDictionary(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating DlpDictionary resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateDlpDictionary(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating DlpDictionary resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceDlpDictionaryRead(d, m)
}

func resourceDlpDictionaryUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDlpDictionary(d)
	if err != nil {
		return fmt.Errorf("Error updating DlpDictionary resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateDlpDictionary(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating DlpDictionary resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceDlpDictionaryRead(d, m)
}

func resourceDlpDictionaryDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteDlpDictionary(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting DlpDictionary resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDlpDictionaryRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadDlpDictionary(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading DlpDictionary resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDlpDictionary(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DlpDictionary resource from API: %v", err)
	}
	return nil
}

func flattenDlpDictionaryComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpDictionaryEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenDlpDictionaryEntriesComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "DlpDictionary-Entries-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenDlpDictionaryEntriesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "DlpDictionary-Entries-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ignore_case"
		if _, ok := i["ignore-case"]; ok {
			v := flattenDlpDictionaryEntriesIgnoreCase(i["ignore-case"], d, pre_append)
			tmp["ignore_case"] = fortiAPISubPartPatch(v, "DlpDictionary-Entries-IgnoreCase")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			v := flattenDlpDictionaryEntriesPattern(i["pattern"], d, pre_append)
			tmp["pattern"] = fortiAPISubPartPatch(v, "DlpDictionary-Entries-Pattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "repeat"
		if _, ok := i["repeat"]; ok {
			v := flattenDlpDictionaryEntriesRepeat(i["repeat"], d, pre_append)
			tmp["repeat"] = fortiAPISubPartPatch(v, "DlpDictionary-Entries-Repeat")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenDlpDictionaryEntriesStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "DlpDictionary-Entries-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenDlpDictionaryEntriesType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "DlpDictionary-Entries-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenDlpDictionaryEntriesComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpDictionaryEntriesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpDictionaryEntriesIgnoreCase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpDictionaryEntriesPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpDictionaryEntriesRepeat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpDictionaryEntriesStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpDictionaryEntriesType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDlpDictionaryFgdId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpDictionaryMatchAround(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpDictionaryMatchType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpDictionaryName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpDictionaryUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDlpDictionary(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comment", flattenDlpDictionaryComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "DlpDictionary-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("entries", flattenDlpDictionaryEntries(o["entries"], d, "entries")); err != nil {
			if vv, ok := fortiAPIPatch(o["entries"], "DlpDictionary-Entries"); ok {
				if err = d.Set("entries", vv); err != nil {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenDlpDictionaryEntries(o["entries"], d, "entries")); err != nil {
				if vv, ok := fortiAPIPatch(o["entries"], "DlpDictionary-Entries"); ok {
					if err = d.Set("entries", vv); err != nil {
						return fmt.Errorf("Error reading entries: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			}
		}
	}

	if err = d.Set("fgd_id", flattenDlpDictionaryFgdId(o["fgd-id"], d, "fgd_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgd-id"], "DlpDictionary-FgdId"); ok {
			if err = d.Set("fgd_id", vv); err != nil {
				return fmt.Errorf("Error reading fgd_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgd_id: %v", err)
		}
	}

	if err = d.Set("match_around", flattenDlpDictionaryMatchAround(o["match-around"], d, "match_around")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-around"], "DlpDictionary-MatchAround"); ok {
			if err = d.Set("match_around", vv); err != nil {
				return fmt.Errorf("Error reading match_around: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_around: %v", err)
		}
	}

	if err = d.Set("match_type", flattenDlpDictionaryMatchType(o["match-type"], d, "match_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-type"], "DlpDictionary-MatchType"); ok {
			if err = d.Set("match_type", vv); err != nil {
				return fmt.Errorf("Error reading match_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_type: %v", err)
		}
	}

	if err = d.Set("name", flattenDlpDictionaryName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "DlpDictionary-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("uuid", flattenDlpDictionaryUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "DlpDictionary-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	return nil
}

func flattenDlpDictionaryFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDlpDictionaryComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpDictionaryEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandDlpDictionaryEntriesComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandDlpDictionaryEntriesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ignore_case"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ignore-case"], _ = expandDlpDictionaryEntriesIgnoreCase(d, i["ignore_case"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern"], _ = expandDlpDictionaryEntriesPattern(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "repeat"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["repeat"], _ = expandDlpDictionaryEntriesRepeat(d, i["repeat"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandDlpDictionaryEntriesStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandDlpDictionaryEntriesType(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandDlpDictionaryEntriesComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpDictionaryEntriesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpDictionaryEntriesIgnoreCase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpDictionaryEntriesPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpDictionaryEntriesRepeat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpDictionaryEntriesStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpDictionaryEntriesType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpDictionaryFgdId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpDictionaryMatchAround(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpDictionaryMatchType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpDictionaryName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpDictionaryUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDlpDictionary(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandDlpDictionaryComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("entries"); ok || d.HasChange("entries") {
		t, err := expandDlpDictionaryEntries(d, v, "entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	if v, ok := d.GetOk("fgd_id"); ok || d.HasChange("fgd_id") {
		t, err := expandDlpDictionaryFgdId(d, v, "fgd_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgd-id"] = t
		}
	}

	if v, ok := d.GetOk("match_around"); ok || d.HasChange("match_around") {
		t, err := expandDlpDictionaryMatchAround(d, v, "match_around")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-around"] = t
		}
	}

	if v, ok := d.GetOk("match_type"); ok || d.HasChange("match_type") {
		t, err := expandDlpDictionaryMatchType(d, v, "match_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-type"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandDlpDictionaryName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandDlpDictionaryUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	return &obj, nil
}
