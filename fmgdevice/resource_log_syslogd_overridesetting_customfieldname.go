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

func resourceLogSyslogdOverrideSettingCustomFieldName() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogSyslogdOverrideSettingCustomFieldNameCreate,
		Read:   resourceLogSyslogdOverrideSettingCustomFieldNameRead,
		Update: resourceLogSyslogdOverrideSettingCustomFieldNameUpdate,
		Delete: resourceLogSyslogdOverrideSettingCustomFieldNameDelete,

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

func resourceLogSyslogdOverrideSettingCustomFieldNameCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectLogSyslogdOverrideSettingCustomFieldName(d)
	if err != nil {
		return fmt.Errorf("Error creating LogSyslogdOverrideSettingCustomFieldName resource while getting object: %v", err)
	}

	_, err = c.CreateLogSyslogdOverrideSettingCustomFieldName(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating LogSyslogdOverrideSettingCustomFieldName resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceLogSyslogdOverrideSettingCustomFieldNameRead(d, m)
}

func resourceLogSyslogdOverrideSettingCustomFieldNameUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectLogSyslogdOverrideSettingCustomFieldName(d)
	if err != nil {
		return fmt.Errorf("Error updating LogSyslogdOverrideSettingCustomFieldName resource while getting object: %v", err)
	}

	_, err = c.UpdateLogSyslogdOverrideSettingCustomFieldName(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LogSyslogdOverrideSettingCustomFieldName resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceLogSyslogdOverrideSettingCustomFieldNameRead(d, m)
}

func resourceLogSyslogdOverrideSettingCustomFieldNameDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteLogSyslogdOverrideSettingCustomFieldName(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting LogSyslogdOverrideSettingCustomFieldName resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogSyslogdOverrideSettingCustomFieldNameRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogSyslogdOverrideSettingCustomFieldName(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading LogSyslogdOverrideSettingCustomFieldName resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogSyslogdOverrideSettingCustomFieldName(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogSyslogdOverrideSettingCustomFieldName resource from API: %v", err)
	}
	return nil
}

func flattenLogSyslogdOverrideSettingCustomFieldNameCustom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogdOverrideSettingCustomFieldNameId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogdOverrideSettingCustomFieldNameName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogSyslogdOverrideSettingCustomFieldName(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("custom", flattenLogSyslogdOverrideSettingCustomFieldNameCustom2edl(o["custom"], d, "custom")); err != nil {
		if vv, ok := fortiAPIPatch(o["custom"], "LogSyslogdOverrideSettingCustomFieldName-Custom"); ok {
			if err = d.Set("custom", vv); err != nil {
				return fmt.Errorf("Error reading custom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading custom: %v", err)
		}
	}

	if err = d.Set("fosid", flattenLogSyslogdOverrideSettingCustomFieldNameId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "LogSyslogdOverrideSettingCustomFieldName-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenLogSyslogdOverrideSettingCustomFieldNameName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "LogSyslogdOverrideSettingCustomFieldName-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenLogSyslogdOverrideSettingCustomFieldNameFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogSyslogdOverrideSettingCustomFieldNameCustom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogdOverrideSettingCustomFieldNameId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogdOverrideSettingCustomFieldNameName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogSyslogdOverrideSettingCustomFieldName(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("custom"); ok || d.HasChange("custom") {
		t, err := expandLogSyslogdOverrideSettingCustomFieldNameCustom2edl(d, v, "custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandLogSyslogdOverrideSettingCustomFieldNameId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandLogSyslogdOverrideSettingCustomFieldNameName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
