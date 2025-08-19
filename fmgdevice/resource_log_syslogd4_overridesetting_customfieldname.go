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

func resourceLogSyslogd4OverrideSettingCustomFieldName() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogSyslogd4OverrideSettingCustomFieldNameCreate,
		Read:   resourceLogSyslogd4OverrideSettingCustomFieldNameRead,
		Update: resourceLogSyslogd4OverrideSettingCustomFieldNameUpdate,
		Delete: resourceLogSyslogd4OverrideSettingCustomFieldNameDelete,

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

func resourceLogSyslogd4OverrideSettingCustomFieldNameCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

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

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectLogSyslogd4OverrideSettingCustomFieldName(d)
	if err != nil {
		return fmt.Errorf("Error creating LogSyslogd4OverrideSettingCustomFieldName resource while getting object: %v", err)
	}

	_, err = c.CreateLogSyslogd4OverrideSettingCustomFieldName(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating LogSyslogd4OverrideSettingCustomFieldName resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceLogSyslogd4OverrideSettingCustomFieldNameRead(d, m)
}

func resourceLogSyslogd4OverrideSettingCustomFieldNameUpdate(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectLogSyslogd4OverrideSettingCustomFieldName(d)
	if err != nil {
		return fmt.Errorf("Error updating LogSyslogd4OverrideSettingCustomFieldName resource while getting object: %v", err)
	}

	_, err = c.UpdateLogSyslogd4OverrideSettingCustomFieldName(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LogSyslogd4OverrideSettingCustomFieldName resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceLogSyslogd4OverrideSettingCustomFieldNameRead(d, m)
}

func resourceLogSyslogd4OverrideSettingCustomFieldNameDelete(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteLogSyslogd4OverrideSettingCustomFieldName(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting LogSyslogd4OverrideSettingCustomFieldName resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogSyslogd4OverrideSettingCustomFieldNameRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogSyslogd4OverrideSettingCustomFieldName(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading LogSyslogd4OverrideSettingCustomFieldName resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogSyslogd4OverrideSettingCustomFieldName(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogSyslogd4OverrideSettingCustomFieldName resource from API: %v", err)
	}
	return nil
}

func flattenLogSyslogd4OverrideSettingCustomFieldNameCustom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideSettingCustomFieldNameId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideSettingCustomFieldNameName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogSyslogd4OverrideSettingCustomFieldName(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("custom", flattenLogSyslogd4OverrideSettingCustomFieldNameCustom2edl(o["custom"], d, "custom")); err != nil {
		if vv, ok := fortiAPIPatch(o["custom"], "LogSyslogd4OverrideSettingCustomFieldName-Custom"); ok {
			if err = d.Set("custom", vv); err != nil {
				return fmt.Errorf("Error reading custom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading custom: %v", err)
		}
	}

	if err = d.Set("fosid", flattenLogSyslogd4OverrideSettingCustomFieldNameId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "LogSyslogd4OverrideSettingCustomFieldName-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenLogSyslogd4OverrideSettingCustomFieldNameName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "LogSyslogd4OverrideSettingCustomFieldName-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenLogSyslogd4OverrideSettingCustomFieldNameFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogSyslogd4OverrideSettingCustomFieldNameCustom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideSettingCustomFieldNameId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideSettingCustomFieldNameName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogSyslogd4OverrideSettingCustomFieldName(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("custom"); ok || d.HasChange("custom") {
		t, err := expandLogSyslogd4OverrideSettingCustomFieldNameCustom2edl(d, v, "custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandLogSyslogd4OverrideSettingCustomFieldNameId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandLogSyslogd4OverrideSettingCustomFieldNameName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
