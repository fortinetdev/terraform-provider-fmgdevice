// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device LogSyslogd3OverrideSettingLogTemplates

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogSyslogd3OverrideSettingLogTemplates() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogSyslogd3OverrideSettingLogTemplatesCreate,
		Read:   resourceLogSyslogd3OverrideSettingLogTemplatesRead,
		Update: resourceLogSyslogd3OverrideSettingLogTemplatesUpdate,
		Delete: resourceLogSyslogd3OverrideSettingLogTemplatesDelete,

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

func resourceLogSyslogd3OverrideSettingLogTemplatesCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLogSyslogd3OverrideSettingLogTemplates(d)
	if err != nil {
		return fmt.Errorf("Error creating LogSyslogd3OverrideSettingLogTemplates resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadLogSyslogd3OverrideSettingLogTemplates(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateLogSyslogd3OverrideSettingLogTemplates(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating LogSyslogd3OverrideSettingLogTemplates resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateLogSyslogd3OverrideSettingLogTemplates(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating LogSyslogd3OverrideSettingLogTemplates resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceLogSyslogd3OverrideSettingLogTemplatesRead(d, m)
			} else {
				return fmt.Errorf("Error creating LogSyslogd3OverrideSettingLogTemplates resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceLogSyslogd3OverrideSettingLogTemplatesRead(d, m)
}

func resourceLogSyslogd3OverrideSettingLogTemplatesUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLogSyslogd3OverrideSettingLogTemplates(d)
	if err != nil {
		return fmt.Errorf("Error updating LogSyslogd3OverrideSettingLogTemplates resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateLogSyslogd3OverrideSettingLogTemplates(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LogSyslogd3OverrideSettingLogTemplates resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceLogSyslogd3OverrideSettingLogTemplatesRead(d, m)
}

func resourceLogSyslogd3OverrideSettingLogTemplatesDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteLogSyslogd3OverrideSettingLogTemplates(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting LogSyslogd3OverrideSettingLogTemplates resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogSyslogd3OverrideSettingLogTemplatesRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogSyslogd3OverrideSettingLogTemplates(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading LogSyslogd3OverrideSettingLogTemplates resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogSyslogd3OverrideSettingLogTemplates(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogSyslogd3OverrideSettingLogTemplates resource from API: %v", err)
	}
	return nil
}

func flattenLogSyslogd3OverrideSettingLogTemplatesCategory2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3OverrideSettingLogTemplatesEmptyValueIndicator2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3OverrideSettingLogTemplatesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3OverrideSettingLogTemplatesTemplate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogSyslogd3OverrideSettingLogTemplates(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("category", flattenLogSyslogd3OverrideSettingLogTemplatesCategory2edl(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "LogSyslogd3OverrideSettingLogTemplates-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("empty_value_indicator", flattenLogSyslogd3OverrideSettingLogTemplatesEmptyValueIndicator2edl(o["empty-value-indicator"], d, "empty_value_indicator")); err != nil {
		if vv, ok := fortiAPIPatch(o["empty-value-indicator"], "LogSyslogd3OverrideSettingLogTemplates-EmptyValueIndicator"); ok {
			if err = d.Set("empty_value_indicator", vv); err != nil {
				return fmt.Errorf("Error reading empty_value_indicator: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading empty_value_indicator: %v", err)
		}
	}

	if err = d.Set("fosid", flattenLogSyslogd3OverrideSettingLogTemplatesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "LogSyslogd3OverrideSettingLogTemplates-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("template", flattenLogSyslogd3OverrideSettingLogTemplatesTemplate2edl(o["template"], d, "template")); err != nil {
		if vv, ok := fortiAPIPatch(o["template"], "LogSyslogd3OverrideSettingLogTemplates-Template"); ok {
			if err = d.Set("template", vv); err != nil {
				return fmt.Errorf("Error reading template: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading template: %v", err)
		}
	}

	return nil
}

func flattenLogSyslogd3OverrideSettingLogTemplatesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogSyslogd3OverrideSettingLogTemplatesCategory2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3OverrideSettingLogTemplatesEmptyValueIndicator2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3OverrideSettingLogTemplatesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3OverrideSettingLogTemplatesTemplate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogSyslogd3OverrideSettingLogTemplates(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandLogSyslogd3OverrideSettingLogTemplatesCategory2edl(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("empty_value_indicator"); ok || d.HasChange("empty_value_indicator") {
		t, err := expandLogSyslogd3OverrideSettingLogTemplatesEmptyValueIndicator2edl(d, v, "empty_value_indicator")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["empty-value-indicator"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandLogSyslogd3OverrideSettingLogTemplatesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("template"); ok || d.HasChange("template") {
		t, err := expandLogSyslogd3OverrideSettingLogTemplatesTemplate2edl(d, v, "template")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["template"] = t
		}
	}

	return &obj, nil
}
