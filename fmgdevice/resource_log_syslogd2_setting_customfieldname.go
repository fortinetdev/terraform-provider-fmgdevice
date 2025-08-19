// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Custom field name for CEF format logging.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogSyslogd2SettingCustomFieldName() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogSyslogd2SettingCustomFieldNameCreate,
		Read:   resourceLogSyslogd2SettingCustomFieldNameRead,
		Update: resourceLogSyslogd2SettingCustomFieldNameUpdate,
		Delete: resourceLogSyslogd2SettingCustomFieldNameDelete,

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
			"custom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
		},
	}
}

func resourceLogSyslogd2SettingCustomFieldNameCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectLogSyslogd2SettingCustomFieldName(d)
	if err != nil {
		return fmt.Errorf("Error creating LogSyslogd2SettingCustomFieldName resource while getting object: %v", err)
	}

	_, err = c.CreateLogSyslogd2SettingCustomFieldName(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating LogSyslogd2SettingCustomFieldName resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceLogSyslogd2SettingCustomFieldNameRead(d, m)
}

func resourceLogSyslogd2SettingCustomFieldNameUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectLogSyslogd2SettingCustomFieldName(d)
	if err != nil {
		return fmt.Errorf("Error updating LogSyslogd2SettingCustomFieldName resource while getting object: %v", err)
	}

	_, err = c.UpdateLogSyslogd2SettingCustomFieldName(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LogSyslogd2SettingCustomFieldName resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceLogSyslogd2SettingCustomFieldNameRead(d, m)
}

func resourceLogSyslogd2SettingCustomFieldNameDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteLogSyslogd2SettingCustomFieldName(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting LogSyslogd2SettingCustomFieldName resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogSyslogd2SettingCustomFieldNameRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogSyslogd2SettingCustomFieldName(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading LogSyslogd2SettingCustomFieldName resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogSyslogd2SettingCustomFieldName(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogSyslogd2SettingCustomFieldName resource from API: %v", err)
	}
	return nil
}

func flattenLogSyslogd2SettingCustomFieldNameCustom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2SettingCustomFieldNameId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2SettingCustomFieldNameName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogSyslogd2SettingCustomFieldName(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("custom", flattenLogSyslogd2SettingCustomFieldNameCustom2edl(o["custom"], d, "custom")); err != nil {
		if vv, ok := fortiAPIPatch(o["custom"], "LogSyslogd2SettingCustomFieldName-Custom"); ok {
			if err = d.Set("custom", vv); err != nil {
				return fmt.Errorf("Error reading custom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading custom: %v", err)
		}
	}

	if err = d.Set("fosid", flattenLogSyslogd2SettingCustomFieldNameId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "LogSyslogd2SettingCustomFieldName-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenLogSyslogd2SettingCustomFieldNameName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "LogSyslogd2SettingCustomFieldName-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenLogSyslogd2SettingCustomFieldNameFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogSyslogd2SettingCustomFieldNameCustom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2SettingCustomFieldNameId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2SettingCustomFieldNameName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogSyslogd2SettingCustomFieldName(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("custom"); ok || d.HasChange("custom") {
		t, err := expandLogSyslogd2SettingCustomFieldNameCustom2edl(d, v, "custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandLogSyslogd2SettingCustomFieldNameId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandLogSyslogd2SettingCustomFieldNameName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
