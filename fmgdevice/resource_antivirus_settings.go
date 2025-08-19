// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure AntiVirus settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAntivirusSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceAntivirusSettingsUpdate,
		Read:   resourceAntivirusSettingsRead,
		Update: resourceAntivirusSettingsUpdate,
		Delete: resourceAntivirusSettingsDelete,

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
			"default_db": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cache_clean_result": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cache_infected_result": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"grayware": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"machine_learning_detection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"use_extreme_db": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceAntivirusSettingsUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectAntivirusSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating AntivirusSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateAntivirusSettings(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating AntivirusSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("AntivirusSettings")

	return resourceAntivirusSettingsRead(d, m)
}

func resourceAntivirusSettingsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteAntivirusSettings(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting AntivirusSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAntivirusSettingsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadAntivirusSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading AntivirusSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectAntivirusSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading AntivirusSettings resource from API: %v", err)
	}
	return nil
}

func flattenAntivirusSettingsDefaultDb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusSettingsCacheCleanResult(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusSettingsCacheInfectedResult(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusSettingsGrayware(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusSettingsMachineLearningDetection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusSettingsOverrideTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusSettingsUseExtremeDb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectAntivirusSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("default_db", flattenAntivirusSettingsDefaultDb(o["default-db"], d, "default_db")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-db"], "AntivirusSettings-DefaultDb"); ok {
			if err = d.Set("default_db", vv); err != nil {
				return fmt.Errorf("Error reading default_db: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_db: %v", err)
		}
	}

	if err = d.Set("cache_clean_result", flattenAntivirusSettingsCacheCleanResult(o["cache-clean-result"], d, "cache_clean_result")); err != nil {
		if vv, ok := fortiAPIPatch(o["cache-clean-result"], "AntivirusSettings-CacheCleanResult"); ok {
			if err = d.Set("cache_clean_result", vv); err != nil {
				return fmt.Errorf("Error reading cache_clean_result: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cache_clean_result: %v", err)
		}
	}

	if err = d.Set("cache_infected_result", flattenAntivirusSettingsCacheInfectedResult(o["cache-infected-result"], d, "cache_infected_result")); err != nil {
		if vv, ok := fortiAPIPatch(o["cache-infected-result"], "AntivirusSettings-CacheInfectedResult"); ok {
			if err = d.Set("cache_infected_result", vv); err != nil {
				return fmt.Errorf("Error reading cache_infected_result: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cache_infected_result: %v", err)
		}
	}

	if err = d.Set("grayware", flattenAntivirusSettingsGrayware(o["grayware"], d, "grayware")); err != nil {
		if vv, ok := fortiAPIPatch(o["grayware"], "AntivirusSettings-Grayware"); ok {
			if err = d.Set("grayware", vv); err != nil {
				return fmt.Errorf("Error reading grayware: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading grayware: %v", err)
		}
	}

	if err = d.Set("machine_learning_detection", flattenAntivirusSettingsMachineLearningDetection(o["machine-learning-detection"], d, "machine_learning_detection")); err != nil {
		if vv, ok := fortiAPIPatch(o["machine-learning-detection"], "AntivirusSettings-MachineLearningDetection"); ok {
			if err = d.Set("machine_learning_detection", vv); err != nil {
				return fmt.Errorf("Error reading machine_learning_detection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading machine_learning_detection: %v", err)
		}
	}

	if err = d.Set("override_timeout", flattenAntivirusSettingsOverrideTimeout(o["override-timeout"], d, "override_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-timeout"], "AntivirusSettings-OverrideTimeout"); ok {
			if err = d.Set("override_timeout", vv); err != nil {
				return fmt.Errorf("Error reading override_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_timeout: %v", err)
		}
	}

	if err = d.Set("use_extreme_db", flattenAntivirusSettingsUseExtremeDb(o["use-extreme-db"], d, "use_extreme_db")); err != nil {
		if vv, ok := fortiAPIPatch(o["use-extreme-db"], "AntivirusSettings-UseExtremeDb"); ok {
			if err = d.Set("use_extreme_db", vv); err != nil {
				return fmt.Errorf("Error reading use_extreme_db: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading use_extreme_db: %v", err)
		}
	}

	return nil
}

func flattenAntivirusSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandAntivirusSettingsDefaultDb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusSettingsCacheCleanResult(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusSettingsCacheInfectedResult(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusSettingsGrayware(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusSettingsMachineLearningDetection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusSettingsOverrideTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusSettingsUseExtremeDb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectAntivirusSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("default_db"); ok || d.HasChange("default_db") {
		t, err := expandAntivirusSettingsDefaultDb(d, v, "default_db")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-db"] = t
		}
	}

	if v, ok := d.GetOk("cache_clean_result"); ok || d.HasChange("cache_clean_result") {
		t, err := expandAntivirusSettingsCacheCleanResult(d, v, "cache_clean_result")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-clean-result"] = t
		}
	}

	if v, ok := d.GetOk("cache_infected_result"); ok || d.HasChange("cache_infected_result") {
		t, err := expandAntivirusSettingsCacheInfectedResult(d, v, "cache_infected_result")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-infected-result"] = t
		}
	}

	if v, ok := d.GetOk("grayware"); ok || d.HasChange("grayware") {
		t, err := expandAntivirusSettingsGrayware(d, v, "grayware")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["grayware"] = t
		}
	}

	if v, ok := d.GetOk("machine_learning_detection"); ok || d.HasChange("machine_learning_detection") {
		t, err := expandAntivirusSettingsMachineLearningDetection(d, v, "machine_learning_detection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["machine-learning-detection"] = t
		}
	}

	if v, ok := d.GetOk("override_timeout"); ok || d.HasChange("override_timeout") {
		t, err := expandAntivirusSettingsOverrideTimeout(d, v, "override_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-timeout"] = t
		}
	}

	if v, ok := d.GetOk("use_extreme_db"); ok || d.HasChange("use_extreme_db") {
		t, err := expandAntivirusSettingsUseExtremeDb(d, v, "use_extreme_db")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-extreme-db"] = t
		}
	}

	return &obj, nil
}
