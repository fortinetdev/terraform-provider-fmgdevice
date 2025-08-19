// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Global settings for memory logging.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogMemoryGlobalSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogMemoryGlobalSettingUpdate,
		Read:   resourceLogMemoryGlobalSettingRead,
		Update: resourceLogMemoryGlobalSettingUpdate,
		Delete: resourceLogMemoryGlobalSettingDelete,

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
			"full_final_warning_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"full_first_warning_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"full_second_warning_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLogMemoryGlobalSettingUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectLogMemoryGlobalSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating LogMemoryGlobalSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateLogMemoryGlobalSetting(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LogMemoryGlobalSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("LogMemoryGlobalSetting")

	return resourceLogMemoryGlobalSettingRead(d, m)
}

func resourceLogMemoryGlobalSettingDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteLogMemoryGlobalSetting(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting LogMemoryGlobalSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogMemoryGlobalSettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogMemoryGlobalSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading LogMemoryGlobalSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogMemoryGlobalSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogMemoryGlobalSetting resource from API: %v", err)
	}
	return nil
}

func flattenLogMemoryGlobalSettingFullFinalWarningThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryGlobalSettingFullFirstWarningThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryGlobalSettingFullSecondWarningThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryGlobalSettingMaxSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogMemoryGlobalSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("full_final_warning_threshold", flattenLogMemoryGlobalSettingFullFinalWarningThreshold(o["full-final-warning-threshold"], d, "full_final_warning_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["full-final-warning-threshold"], "LogMemoryGlobalSetting-FullFinalWarningThreshold"); ok {
			if err = d.Set("full_final_warning_threshold", vv); err != nil {
				return fmt.Errorf("Error reading full_final_warning_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading full_final_warning_threshold: %v", err)
		}
	}

	if err = d.Set("full_first_warning_threshold", flattenLogMemoryGlobalSettingFullFirstWarningThreshold(o["full-first-warning-threshold"], d, "full_first_warning_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["full-first-warning-threshold"], "LogMemoryGlobalSetting-FullFirstWarningThreshold"); ok {
			if err = d.Set("full_first_warning_threshold", vv); err != nil {
				return fmt.Errorf("Error reading full_first_warning_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading full_first_warning_threshold: %v", err)
		}
	}

	if err = d.Set("full_second_warning_threshold", flattenLogMemoryGlobalSettingFullSecondWarningThreshold(o["full-second-warning-threshold"], d, "full_second_warning_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["full-second-warning-threshold"], "LogMemoryGlobalSetting-FullSecondWarningThreshold"); ok {
			if err = d.Set("full_second_warning_threshold", vv); err != nil {
				return fmt.Errorf("Error reading full_second_warning_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading full_second_warning_threshold: %v", err)
		}
	}

	if err = d.Set("max_size", flattenLogMemoryGlobalSettingMaxSize(o["max-size"], d, "max_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-size"], "LogMemoryGlobalSetting-MaxSize"); ok {
			if err = d.Set("max_size", vv); err != nil {
				return fmt.Errorf("Error reading max_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_size: %v", err)
		}
	}

	return nil
}

func flattenLogMemoryGlobalSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogMemoryGlobalSettingFullFinalWarningThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryGlobalSettingFullFirstWarningThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryGlobalSettingFullSecondWarningThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryGlobalSettingMaxSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogMemoryGlobalSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("full_final_warning_threshold"); ok || d.HasChange("full_final_warning_threshold") {
		t, err := expandLogMemoryGlobalSettingFullFinalWarningThreshold(d, v, "full_final_warning_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["full-final-warning-threshold"] = t
		}
	}

	if v, ok := d.GetOk("full_first_warning_threshold"); ok || d.HasChange("full_first_warning_threshold") {
		t, err := expandLogMemoryGlobalSettingFullFirstWarningThreshold(d, v, "full_first_warning_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["full-first-warning-threshold"] = t
		}
	}

	if v, ok := d.GetOk("full_second_warning_threshold"); ok || d.HasChange("full_second_warning_threshold") {
		t, err := expandLogMemoryGlobalSettingFullSecondWarningThreshold(d, v, "full_second_warning_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["full-second-warning-threshold"] = t
		}
	}

	if v, ok := d.GetOk("max_size"); ok || d.HasChange("max_size") {
		t, err := expandLogMemoryGlobalSettingMaxSize(d, v, "max_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-size"] = t
		}
	}

	return &obj, nil
}
