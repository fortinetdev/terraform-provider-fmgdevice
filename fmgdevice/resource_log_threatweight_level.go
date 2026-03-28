// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Score mapping for threat weight levels.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogThreatWeightLevel() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogThreatWeightLevelUpdate,
		Read:   resourceLogThreatWeightLevelRead,
		Update: resourceLogThreatWeightLevelUpdate,
		Delete: resourceLogThreatWeightLevelDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

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
			"critical": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"high": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"low": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"medium": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLogThreatWeightLevelUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLogThreatWeightLevel(d)
	if err != nil {
		return fmt.Errorf("Error updating LogThreatWeightLevel resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateLogThreatWeightLevel(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LogThreatWeightLevel resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("LogThreatWeightLevel")

	return resourceLogThreatWeightLevelRead(d, m)
}

func resourceLogThreatWeightLevelDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteLogThreatWeightLevel(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting LogThreatWeightLevel resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogThreatWeightLevelRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogThreatWeightLevel(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading LogThreatWeightLevel resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogThreatWeightLevel(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogThreatWeightLevel resource from API: %v", err)
	}
	return nil
}

func flattenLogThreatWeightLevelCritical2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightLevelHigh2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightLevelLow2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightLevelMedium2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogThreatWeightLevel(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("critical", flattenLogThreatWeightLevelCritical2edl(o["critical"], d, "critical")); err != nil {
		if vv, ok := fortiAPIPatch(o["critical"], "LogThreatWeightLevel-Critical"); ok {
			if err = d.Set("critical", vv); err != nil {
				return fmt.Errorf("Error reading critical: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading critical: %v", err)
		}
	}

	if err = d.Set("high", flattenLogThreatWeightLevelHigh2edl(o["high"], d, "high")); err != nil {
		if vv, ok := fortiAPIPatch(o["high"], "LogThreatWeightLevel-High"); ok {
			if err = d.Set("high", vv); err != nil {
				return fmt.Errorf("Error reading high: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading high: %v", err)
		}
	}

	if err = d.Set("low", flattenLogThreatWeightLevelLow2edl(o["low"], d, "low")); err != nil {
		if vv, ok := fortiAPIPatch(o["low"], "LogThreatWeightLevel-Low"); ok {
			if err = d.Set("low", vv); err != nil {
				return fmt.Errorf("Error reading low: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading low: %v", err)
		}
	}

	if err = d.Set("medium", flattenLogThreatWeightLevelMedium2edl(o["medium"], d, "medium")); err != nil {
		if vv, ok := fortiAPIPatch(o["medium"], "LogThreatWeightLevel-Medium"); ok {
			if err = d.Set("medium", vv); err != nil {
				return fmt.Errorf("Error reading medium: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading medium: %v", err)
		}
	}

	return nil
}

func flattenLogThreatWeightLevelFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogThreatWeightLevelCritical2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightLevelHigh2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightLevelLow2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightLevelMedium2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogThreatWeightLevel(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("critical"); ok || d.HasChange("critical") {
		t, err := expandLogThreatWeightLevelCritical2edl(d, v, "critical")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["critical"] = t
		}
	}

	if v, ok := d.GetOk("high"); ok || d.HasChange("high") {
		t, err := expandLogThreatWeightLevelHigh2edl(d, v, "high")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["high"] = t
		}
	}

	if v, ok := d.GetOk("low"); ok || d.HasChange("low") {
		t, err := expandLogThreatWeightLevelLow2edl(d, v, "low")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["low"] = t
		}
	}

	if v, ok := d.GetOk("medium"); ok || d.HasChange("medium") {
		t, err := expandLogThreatWeightLevelMedium2edl(d, v, "medium")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["medium"] = t
		}
	}

	return &obj, nil
}
