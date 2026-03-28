// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure labels used by DLP blocking.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDlpLabel() *schema.Resource {
	return &schema.Resource{
		Create: resourceDlpLabelCreate,
		Read:   resourceDlpLabelRead,
		Update: resourceDlpLabelUpdate,
		Delete: resourceDlpLabelDelete,

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
			"connector": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"entries": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fortidata_label_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"guid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mpip_label_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"mpip_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"type": &schema.Schema{
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

func resourceDlpLabelCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDlpLabel(d)
	if err != nil {
		return fmt.Errorf("Error creating DlpLabel resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadDlpLabel(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateDlpLabel(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating DlpLabel resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateDlpLabel(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating DlpLabel resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceDlpLabelRead(d, m)
}

func resourceDlpLabelUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDlpLabel(d)
	if err != nil {
		return fmt.Errorf("Error updating DlpLabel resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateDlpLabel(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating DlpLabel resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceDlpLabelRead(d, m)
}

func resourceDlpLabelDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteDlpLabel(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting DlpLabel resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDlpLabelRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadDlpLabel(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading DlpLabel resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDlpLabel(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DlpLabel resource from API: %v", err)
	}
	return nil
}

func flattenDlpLabelComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpLabelConnector(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDlpLabelEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortidata_label_name"
		if _, ok := i["fortidata-label-name"]; ok {
			v := flattenDlpLabelEntriesFortidataLabelName(i["fortidata-label-name"], d, pre_append)
			tmp["fortidata_label_name"] = fortiAPISubPartPatch(v, "DlpLabel-Entries-FortidataLabelName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "guid"
		if _, ok := i["guid"]; ok {
			v := flattenDlpLabelEntriesGuid(i["guid"], d, pre_append)
			tmp["guid"] = fortiAPISubPartPatch(v, "DlpLabel-Entries-Guid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenDlpLabelEntriesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "DlpLabel-Entries-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpip_label_name"
		if _, ok := i["mpip-label-name"]; ok {
			v := flattenDlpLabelEntriesMpipLabelName(i["mpip-label-name"], d, pre_append)
			tmp["mpip_label_name"] = fortiAPISubPartPatch(v, "DlpLabel-Entries-MpipLabelName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenDlpLabelEntriesFortidataLabelName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpLabelEntriesGuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpLabelEntriesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpLabelEntriesMpipLabelName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpLabelMpipType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpLabelName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpLabelType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDlpLabel(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comment", flattenDlpLabelComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "DlpLabel-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("connector", flattenDlpLabelConnector(o["connector"], d, "connector")); err != nil {
		if vv, ok := fortiAPIPatch(o["connector"], "DlpLabel-Connector"); ok {
			if err = d.Set("connector", vv); err != nil {
				return fmt.Errorf("Error reading connector: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading connector: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("entries", flattenDlpLabelEntries(o["entries"], d, "entries")); err != nil {
			if vv, ok := fortiAPIPatch(o["entries"], "DlpLabel-Entries"); ok {
				if err = d.Set("entries", vv); err != nil {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenDlpLabelEntries(o["entries"], d, "entries")); err != nil {
				if vv, ok := fortiAPIPatch(o["entries"], "DlpLabel-Entries"); ok {
					if err = d.Set("entries", vv); err != nil {
						return fmt.Errorf("Error reading entries: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			}
		}
	}

	if err = d.Set("mpip_type", flattenDlpLabelMpipType(o["mpip-type"], d, "mpip_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["mpip-type"], "DlpLabel-MpipType"); ok {
			if err = d.Set("mpip_type", vv); err != nil {
				return fmt.Errorf("Error reading mpip_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mpip_type: %v", err)
		}
	}

	if err = d.Set("name", flattenDlpLabelName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "DlpLabel-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("type", flattenDlpLabelType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "DlpLabel-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenDlpLabelFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDlpLabelComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpLabelConnector(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpLabelEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortidata_label_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fortidata-label-name"], _ = expandDlpLabelEntriesFortidataLabelName(d, i["fortidata_label_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "guid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["guid"], _ = expandDlpLabelEntriesGuid(d, i["guid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandDlpLabelEntriesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpip_label_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mpip-label-name"], _ = expandDlpLabelEntriesMpipLabelName(d, i["mpip_label_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandDlpLabelEntriesFortidataLabelName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpLabelEntriesGuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpLabelEntriesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpLabelEntriesMpipLabelName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpLabelMpipType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpLabelName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpLabelType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDlpLabel(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandDlpLabelComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("connector"); ok || d.HasChange("connector") {
		t, err := expandDlpLabelConnector(d, v, "connector")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connector"] = t
		}
	}

	if v, ok := d.GetOk("entries"); ok || d.HasChange("entries") {
		t, err := expandDlpLabelEntries(d, v, "entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	if v, ok := d.GetOk("mpip_type"); ok || d.HasChange("mpip_type") {
		t, err := expandDlpLabelMpipType(d, v, "mpip_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpip-type"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandDlpLabelName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandDlpLabelType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
