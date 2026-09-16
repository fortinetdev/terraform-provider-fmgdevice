// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure custom log format.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogCustomFormat() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogCustomFormatCreate,
		Read:   resourceLogCustomFormatRead,
		Update: resourceLogCustomFormatUpdate,
		Delete: resourceLogCustomFormatDelete,

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
			"empty_value_indicator": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"field_exclusion_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"log_templates": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"subtypes": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"template": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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

func resourceLogCustomFormatCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	obj, err := getObjectLogCustomFormat(d)
	if err != nil {
		return fmt.Errorf("Error creating LogCustomFormat resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadLogCustomFormat(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateLogCustomFormat(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating LogCustomFormat resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateLogCustomFormat(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating LogCustomFormat resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceLogCustomFormatRead(d, m)
}

func resourceLogCustomFormatUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	obj, err := getObjectLogCustomFormat(d)
	if err != nil {
		return fmt.Errorf("Error updating LogCustomFormat resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateLogCustomFormat(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LogCustomFormat resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceLogCustomFormatRead(d, m)
}

func resourceLogCustomFormatDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	wsParams["adom"] = adomv

	err = c.DeleteLogCustomFormat(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting LogCustomFormat resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogCustomFormatRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogCustomFormat(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading LogCustomFormat resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogCustomFormat(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogCustomFormat resource from API: %v", err)
	}
	return nil
}

func flattenLogCustomFormatEmptyValueIndicator(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogCustomFormatFieldExclusionList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenLogCustomFormatLogTemplates(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			v := flattenLogCustomFormatLogTemplatesCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "LogCustomFormat-LogTemplates-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenLogCustomFormatLogTemplatesName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "LogCustomFormat-LogTemplates-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subtypes"
		if _, ok := i["subtypes"]; ok {
			v := flattenLogCustomFormatLogTemplatesSubtypes(i["subtypes"], d, pre_append)
			tmp["subtypes"] = fortiAPISubPartPatch(v, "LogCustomFormat-LogTemplates-Subtypes")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "template"
		if _, ok := i["template"]; ok {
			v := flattenLogCustomFormatLogTemplatesTemplate(i["template"], d, pre_append)
			tmp["template"] = fortiAPISubPartPatch(v, "LogCustomFormat-LogTemplates-Template")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenLogCustomFormatLogTemplatesCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogCustomFormatLogTemplatesName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogCustomFormatLogTemplatesSubtypes(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenLogCustomFormatLogTemplatesTemplate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogCustomFormatName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogCustomFormat(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("empty_value_indicator", flattenLogCustomFormatEmptyValueIndicator(o["empty-value-indicator"], d, "empty_value_indicator")); err != nil {
		if vv, ok := fortiAPIPatch(o["empty-value-indicator"], "LogCustomFormat-EmptyValueIndicator"); ok {
			if err = d.Set("empty_value_indicator", vv); err != nil {
				return fmt.Errorf("Error reading empty_value_indicator: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading empty_value_indicator: %v", err)
		}
	}

	if err = d.Set("field_exclusion_list", flattenLogCustomFormatFieldExclusionList(o["field-exclusion-list"], d, "field_exclusion_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["field-exclusion-list"], "LogCustomFormat-FieldExclusionList"); ok {
			if err = d.Set("field_exclusion_list", vv); err != nil {
				return fmt.Errorf("Error reading field_exclusion_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading field_exclusion_list: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("log_templates", flattenLogCustomFormatLogTemplates(o["log-templates"], d, "log_templates")); err != nil {
			if vv, ok := fortiAPIPatch(o["log-templates"], "LogCustomFormat-LogTemplates"); ok {
				if err = d.Set("log_templates", vv); err != nil {
					return fmt.Errorf("Error reading log_templates: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading log_templates: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("log_templates"); ok {
			if err = d.Set("log_templates", flattenLogCustomFormatLogTemplates(o["log-templates"], d, "log_templates")); err != nil {
				if vv, ok := fortiAPIPatch(o["log-templates"], "LogCustomFormat-LogTemplates"); ok {
					if err = d.Set("log_templates", vv); err != nil {
						return fmt.Errorf("Error reading log_templates: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading log_templates: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenLogCustomFormatName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "LogCustomFormat-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenLogCustomFormatFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogCustomFormatEmptyValueIndicator(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogCustomFormatFieldExclusionList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandLogCustomFormatLogTemplates(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["category"], _ = expandLogCustomFormatLogTemplatesCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandLogCustomFormatLogTemplatesName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subtypes"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["subtypes"], _ = expandLogCustomFormatLogTemplatesSubtypes(d, i["subtypes"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "template"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["template"], _ = expandLogCustomFormatLogTemplatesTemplate(d, i["template"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandLogCustomFormatLogTemplatesCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogCustomFormatLogTemplatesName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogCustomFormatLogTemplatesSubtypes(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandLogCustomFormatLogTemplatesTemplate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogCustomFormatName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogCustomFormat(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("empty_value_indicator"); ok || d.HasChange("empty_value_indicator") {
		t, err := expandLogCustomFormatEmptyValueIndicator(d, v, "empty_value_indicator")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["empty-value-indicator"] = t
		}
	}

	if v, ok := d.GetOk("field_exclusion_list"); ok || d.HasChange("field_exclusion_list") {
		t, err := expandLogCustomFormatFieldExclusionList(d, v, "field_exclusion_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["field-exclusion-list"] = t
		}
	}

	if v, ok := d.GetOk("log_templates"); ok || d.HasChange("log_templates") {
		t, err := expandLogCustomFormatLogTemplates(d, v, "log_templates")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-templates"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandLogCustomFormatName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
