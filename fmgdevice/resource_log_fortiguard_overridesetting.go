// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Override global FortiCloud logging settings for this VDOM.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogFortiguardOverrideSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogFortiguardOverrideSettingUpdate,
		Read:   resourceLogFortiguardOverrideSettingRead,
		Update: resourceLogFortiguardOverrideSettingUpdate,
		Delete: resourceLogFortiguardOverrideSettingDelete,

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
			"access_config": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_log_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"upload_day": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"upload_interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"upload_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"upload_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceLogFortiguardOverrideSettingUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectLogFortiguardOverrideSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating LogFortiguardOverrideSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateLogFortiguardOverrideSetting(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LogFortiguardOverrideSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("LogFortiguardOverrideSetting")

	return resourceLogFortiguardOverrideSettingRead(d, m)
}

func resourceLogFortiguardOverrideSettingDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteLogFortiguardOverrideSetting(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting LogFortiguardOverrideSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogFortiguardOverrideSettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogFortiguardOverrideSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading LogFortiguardOverrideSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogFortiguardOverrideSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogFortiguardOverrideSetting resource from API: %v", err)
	}
	return nil
}

func flattenLogFortiguardOverrideSettingAccessConfig(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortiguardOverrideSettingMaxLogRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortiguardOverrideSettingOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortiguardOverrideSettingPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortiguardOverrideSettingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortiguardOverrideSettingUploadDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortiguardOverrideSettingUploadInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortiguardOverrideSettingUploadOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortiguardOverrideSettingUploadTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogFortiguardOverrideSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("access_config", flattenLogFortiguardOverrideSettingAccessConfig(o["access-config"], d, "access_config")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-config"], "LogFortiguardOverrideSetting-AccessConfig"); ok {
			if err = d.Set("access_config", vv); err != nil {
				return fmt.Errorf("Error reading access_config: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_config: %v", err)
		}
	}

	if err = d.Set("max_log_rate", flattenLogFortiguardOverrideSettingMaxLogRate(o["max-log-rate"], d, "max_log_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-log-rate"], "LogFortiguardOverrideSetting-MaxLogRate"); ok {
			if err = d.Set("max_log_rate", vv); err != nil {
				return fmt.Errorf("Error reading max_log_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_log_rate: %v", err)
		}
	}

	if err = d.Set("override", flattenLogFortiguardOverrideSettingOverride(o["override"], d, "override")); err != nil {
		if vv, ok := fortiAPIPatch(o["override"], "LogFortiguardOverrideSetting-Override"); ok {
			if err = d.Set("override", vv); err != nil {
				return fmt.Errorf("Error reading override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override: %v", err)
		}
	}

	if err = d.Set("priority", flattenLogFortiguardOverrideSettingPriority(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "LogFortiguardOverrideSetting-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("status", flattenLogFortiguardOverrideSettingStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "LogFortiguardOverrideSetting-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("upload_day", flattenLogFortiguardOverrideSettingUploadDay(o["upload-day"], d, "upload_day")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-day"], "LogFortiguardOverrideSetting-UploadDay"); ok {
			if err = d.Set("upload_day", vv); err != nil {
				return fmt.Errorf("Error reading upload_day: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_day: %v", err)
		}
	}

	if err = d.Set("upload_interval", flattenLogFortiguardOverrideSettingUploadInterval(o["upload-interval"], d, "upload_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-interval"], "LogFortiguardOverrideSetting-UploadInterval"); ok {
			if err = d.Set("upload_interval", vv); err != nil {
				return fmt.Errorf("Error reading upload_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_interval: %v", err)
		}
	}

	if err = d.Set("upload_option", flattenLogFortiguardOverrideSettingUploadOption(o["upload-option"], d, "upload_option")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-option"], "LogFortiguardOverrideSetting-UploadOption"); ok {
			if err = d.Set("upload_option", vv); err != nil {
				return fmt.Errorf("Error reading upload_option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_option: %v", err)
		}
	}

	if err = d.Set("upload_time", flattenLogFortiguardOverrideSettingUploadTime(o["upload-time"], d, "upload_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-time"], "LogFortiguardOverrideSetting-UploadTime"); ok {
			if err = d.Set("upload_time", vv); err != nil {
				return fmt.Errorf("Error reading upload_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_time: %v", err)
		}
	}

	return nil
}

func flattenLogFortiguardOverrideSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogFortiguardOverrideSettingAccessConfig(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardOverrideSettingMaxLogRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardOverrideSettingOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardOverrideSettingPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardOverrideSettingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardOverrideSettingUploadDay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardOverrideSettingUploadInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardOverrideSettingUploadOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardOverrideSettingUploadTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogFortiguardOverrideSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("access_config"); ok || d.HasChange("access_config") {
		t, err := expandLogFortiguardOverrideSettingAccessConfig(d, v, "access_config")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-config"] = t
		}
	}

	if v, ok := d.GetOk("max_log_rate"); ok || d.HasChange("max_log_rate") {
		t, err := expandLogFortiguardOverrideSettingMaxLogRate(d, v, "max_log_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-log-rate"] = t
		}
	}

	if v, ok := d.GetOk("override"); ok || d.HasChange("override") {
		t, err := expandLogFortiguardOverrideSettingOverride(d, v, "override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandLogFortiguardOverrideSettingPriority(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandLogFortiguardOverrideSettingStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("upload_day"); ok || d.HasChange("upload_day") {
		t, err := expandLogFortiguardOverrideSettingUploadDay(d, v, "upload_day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-day"] = t
		}
	}

	if v, ok := d.GetOk("upload_interval"); ok || d.HasChange("upload_interval") {
		t, err := expandLogFortiguardOverrideSettingUploadInterval(d, v, "upload_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-interval"] = t
		}
	}

	if v, ok := d.GetOk("upload_option"); ok || d.HasChange("upload_option") {
		t, err := expandLogFortiguardOverrideSettingUploadOption(d, v, "upload_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-option"] = t
		}
	}

	if v, ok := d.GetOk("upload_time"); ok || d.HasChange("upload_time") {
		t, err := expandLogFortiguardOverrideSettingUploadTime(d, v, "upload_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-time"] = t
		}
	}

	return &obj, nil
}
