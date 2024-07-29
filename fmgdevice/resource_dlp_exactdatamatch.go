// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure exact-data-match template used by DLP scan.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDlpExactDataMatch() *schema.Resource {
	return &schema.Resource{
		Create: resourceDlpExactDataMatchCreate,
		Read:   resourceDlpExactDataMatchRead,
		Update: resourceDlpExactDataMatchUpdate,
		Delete: resourceDlpExactDataMatchDelete,

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
			"columns": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"index": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"optional": &schema.Schema{
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
			"data": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"optional": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceDlpExactDataMatchCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	obj, err := getObjectDlpExactDataMatch(d)
	if err != nil {
		return fmt.Errorf("Error creating DlpExactDataMatch resource while getting object: %v", err)
	}

	_, err = c.CreateDlpExactDataMatch(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating DlpExactDataMatch resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceDlpExactDataMatchRead(d, m)
}

func resourceDlpExactDataMatchUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	obj, err := getObjectDlpExactDataMatch(d)
	if err != nil {
		return fmt.Errorf("Error updating DlpExactDataMatch resource while getting object: %v", err)
	}

	_, err = c.UpdateDlpExactDataMatch(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating DlpExactDataMatch resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceDlpExactDataMatchRead(d, m)
}

func resourceDlpExactDataMatchDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	err = c.DeleteDlpExactDataMatch(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting DlpExactDataMatch resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDlpExactDataMatchRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadDlpExactDataMatch(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading DlpExactDataMatch resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDlpExactDataMatch(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DlpExactDataMatch resource from API: %v", err)
	}
	return nil
}

func flattenDlpExactDataMatchColumns(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := i["index"]; ok {
			v := flattenDlpExactDataMatchColumnsIndex(i["index"], d, pre_append)
			tmp["index"] = fortiAPISubPartPatch(v, "DlpExactDataMatch-Columns-Index")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "optional"
		if _, ok := i["optional"]; ok {
			v := flattenDlpExactDataMatchColumnsOptional(i["optional"], d, pre_append)
			tmp["optional"] = fortiAPISubPartPatch(v, "DlpExactDataMatch-Columns-Optional")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenDlpExactDataMatchColumnsType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "DlpExactDataMatch-Columns-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenDlpExactDataMatchColumnsIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpExactDataMatchColumnsOptional(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpExactDataMatchColumnsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDlpExactDataMatchData(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDlpExactDataMatchName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpExactDataMatchOptional(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDlpExactDataMatch(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("columns", flattenDlpExactDataMatchColumns(o["columns"], d, "columns")); err != nil {
			if vv, ok := fortiAPIPatch(o["columns"], "DlpExactDataMatch-Columns"); ok {
				if err = d.Set("columns", vv); err != nil {
					return fmt.Errorf("Error reading columns: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading columns: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("columns"); ok {
			if err = d.Set("columns", flattenDlpExactDataMatchColumns(o["columns"], d, "columns")); err != nil {
				if vv, ok := fortiAPIPatch(o["columns"], "DlpExactDataMatch-Columns"); ok {
					if err = d.Set("columns", vv); err != nil {
						return fmt.Errorf("Error reading columns: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading columns: %v", err)
				}
			}
		}
	}

	if err = d.Set("data", flattenDlpExactDataMatchData(o["data"], d, "data")); err != nil {
		if vv, ok := fortiAPIPatch(o["data"], "DlpExactDataMatch-Data"); ok {
			if err = d.Set("data", vv); err != nil {
				return fmt.Errorf("Error reading data: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading data: %v", err)
		}
	}

	if err = d.Set("name", flattenDlpExactDataMatchName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "DlpExactDataMatch-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("optional", flattenDlpExactDataMatchOptional(o["optional"], d, "optional")); err != nil {
		if vv, ok := fortiAPIPatch(o["optional"], "DlpExactDataMatch-Optional"); ok {
			if err = d.Set("optional", vv); err != nil {
				return fmt.Errorf("Error reading optional: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading optional: %v", err)
		}
	}

	return nil
}

func flattenDlpExactDataMatchFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDlpExactDataMatchColumns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["index"], _ = expandDlpExactDataMatchColumnsIndex(d, i["index"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "optional"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["optional"], _ = expandDlpExactDataMatchColumnsOptional(d, i["optional"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandDlpExactDataMatchColumnsType(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandDlpExactDataMatchColumnsIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpExactDataMatchColumnsOptional(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpExactDataMatchColumnsType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpExactDataMatchData(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpExactDataMatchName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpExactDataMatchOptional(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDlpExactDataMatch(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("columns"); ok || d.HasChange("columns") {
		t, err := expandDlpExactDataMatchColumns(d, v, "columns")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["columns"] = t
		}
	}

	if v, ok := d.GetOk("data"); ok || d.HasChange("data") {
		t, err := expandDlpExactDataMatchData(d, v, "data")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["data"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandDlpExactDataMatchName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("optional"); ok || d.HasChange("optional") {
		t, err := expandDlpExactDataMatchOptional(d, v, "optional")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["optional"] = t
		}
	}

	return &obj, nil
}
