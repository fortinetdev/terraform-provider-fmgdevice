// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device LogSyslogdOverrideSettingLogTemplates

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogSyslogdOverrideSettingLogTemplates() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogSyslogdOverrideSettingLogTemplatesCreate,
		Read:   resourceLogSyslogdOverrideSettingLogTemplatesRead,
		Update: resourceLogSyslogdOverrideSettingLogTemplatesUpdate,
		Delete: resourceLogSyslogdOverrideSettingLogTemplatesDelete,

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
			"category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"empty_value_indicator": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"template": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceLogSyslogdOverrideSettingLogTemplatesCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLogSyslogdOverrideSettingLogTemplates(d)
	if err != nil {
		return fmt.Errorf("Error creating LogSyslogdOverrideSettingLogTemplates resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadLogSyslogdOverrideSettingLogTemplates(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateLogSyslogdOverrideSettingLogTemplates(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating LogSyslogdOverrideSettingLogTemplates resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateLogSyslogdOverrideSettingLogTemplates(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating LogSyslogdOverrideSettingLogTemplates resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceLogSyslogdOverrideSettingLogTemplatesRead(d, m)
			} else {
				return fmt.Errorf("Error creating LogSyslogdOverrideSettingLogTemplates resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceLogSyslogdOverrideSettingLogTemplatesRead(d, m)
}

func resourceLogSyslogdOverrideSettingLogTemplatesUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLogSyslogdOverrideSettingLogTemplates(d)
	if err != nil {
		return fmt.Errorf("Error updating LogSyslogdOverrideSettingLogTemplates resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateLogSyslogdOverrideSettingLogTemplates(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LogSyslogdOverrideSettingLogTemplates resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceLogSyslogdOverrideSettingLogTemplatesRead(d, m)
}

func resourceLogSyslogdOverrideSettingLogTemplatesDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteLogSyslogdOverrideSettingLogTemplates(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting LogSyslogdOverrideSettingLogTemplates resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogSyslogdOverrideSettingLogTemplatesRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogSyslogdOverrideSettingLogTemplates(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading LogSyslogdOverrideSettingLogTemplates resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogSyslogdOverrideSettingLogTemplates(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogSyslogdOverrideSettingLogTemplates resource from API: %v", err)
	}
	return nil
}

func flattenLogSyslogdOverrideSettingLogTemplatesCategory2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogdOverrideSettingLogTemplatesEmptyValueIndicator2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogdOverrideSettingLogTemplatesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogdOverrideSettingLogTemplatesTemplate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogSyslogdOverrideSettingLogTemplates(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("category", flattenLogSyslogdOverrideSettingLogTemplatesCategory2edl(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "LogSyslogdOverrideSettingLogTemplates-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("empty_value_indicator", flattenLogSyslogdOverrideSettingLogTemplatesEmptyValueIndicator2edl(o["empty-value-indicator"], d, "empty_value_indicator")); err != nil {
		if vv, ok := fortiAPIPatch(o["empty-value-indicator"], "LogSyslogdOverrideSettingLogTemplates-EmptyValueIndicator"); ok {
			if err = d.Set("empty_value_indicator", vv); err != nil {
				return fmt.Errorf("Error reading empty_value_indicator: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading empty_value_indicator: %v", err)
		}
	}

	if err = d.Set("fosid", flattenLogSyslogdOverrideSettingLogTemplatesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "LogSyslogdOverrideSettingLogTemplates-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("template", flattenLogSyslogdOverrideSettingLogTemplatesTemplate2edl(o["template"], d, "template")); err != nil {
		if vv, ok := fortiAPIPatch(o["template"], "LogSyslogdOverrideSettingLogTemplates-Template"); ok {
			if err = d.Set("template", vv); err != nil {
				return fmt.Errorf("Error reading template: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading template: %v", err)
		}
	}

	return nil
}

func flattenLogSyslogdOverrideSettingLogTemplatesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogSyslogdOverrideSettingLogTemplatesCategory2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogdOverrideSettingLogTemplatesEmptyValueIndicator2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogdOverrideSettingLogTemplatesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogdOverrideSettingLogTemplatesTemplate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogSyslogdOverrideSettingLogTemplates(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandLogSyslogdOverrideSettingLogTemplatesCategory2edl(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("empty_value_indicator"); ok || d.HasChange("empty_value_indicator") {
		t, err := expandLogSyslogdOverrideSettingLogTemplatesEmptyValueIndicator2edl(d, v, "empty_value_indicator")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["empty-value-indicator"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandLogSyslogdOverrideSettingLogTemplatesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("template"); ok || d.HasChange("template") {
		t, err := expandLogSyslogdOverrideSettingLogTemplatesTemplate2edl(d, v, "template")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["template"] = t
		}
	}

	return &obj, nil
}
