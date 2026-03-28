// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IPS threat weight settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogThreatWeightIps() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogThreatWeightIpsUpdate,
		Read:   resourceLogThreatWeightIpsRead,
		Update: resourceLogThreatWeightIpsUpdate,
		Delete: resourceLogThreatWeightIpsDelete,

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
			"critical_severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"high_severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"info_severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"low_severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"medium_severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLogThreatWeightIpsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLogThreatWeightIps(d)
	if err != nil {
		return fmt.Errorf("Error updating LogThreatWeightIps resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateLogThreatWeightIps(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LogThreatWeightIps resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("LogThreatWeightIps")

	return resourceLogThreatWeightIpsRead(d, m)
}

func resourceLogThreatWeightIpsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteLogThreatWeightIps(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting LogThreatWeightIps resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogThreatWeightIpsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogThreatWeightIps(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading LogThreatWeightIps resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogThreatWeightIps(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogThreatWeightIps resource from API: %v", err)
	}
	return nil
}

func flattenLogThreatWeightIpsCriticalSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightIpsHighSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightIpsInfoSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightIpsLowSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightIpsMediumSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogThreatWeightIps(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("critical_severity", flattenLogThreatWeightIpsCriticalSeverity2edl(o["critical-severity"], d, "critical_severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["critical-severity"], "LogThreatWeightIps-CriticalSeverity"); ok {
			if err = d.Set("critical_severity", vv); err != nil {
				return fmt.Errorf("Error reading critical_severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading critical_severity: %v", err)
		}
	}

	if err = d.Set("high_severity", flattenLogThreatWeightIpsHighSeverity2edl(o["high-severity"], d, "high_severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["high-severity"], "LogThreatWeightIps-HighSeverity"); ok {
			if err = d.Set("high_severity", vv); err != nil {
				return fmt.Errorf("Error reading high_severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading high_severity: %v", err)
		}
	}

	if err = d.Set("info_severity", flattenLogThreatWeightIpsInfoSeverity2edl(o["info-severity"], d, "info_severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["info-severity"], "LogThreatWeightIps-InfoSeverity"); ok {
			if err = d.Set("info_severity", vv); err != nil {
				return fmt.Errorf("Error reading info_severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading info_severity: %v", err)
		}
	}

	if err = d.Set("low_severity", flattenLogThreatWeightIpsLowSeverity2edl(o["low-severity"], d, "low_severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["low-severity"], "LogThreatWeightIps-LowSeverity"); ok {
			if err = d.Set("low_severity", vv); err != nil {
				return fmt.Errorf("Error reading low_severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading low_severity: %v", err)
		}
	}

	if err = d.Set("medium_severity", flattenLogThreatWeightIpsMediumSeverity2edl(o["medium-severity"], d, "medium_severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["medium-severity"], "LogThreatWeightIps-MediumSeverity"); ok {
			if err = d.Set("medium_severity", vv); err != nil {
				return fmt.Errorf("Error reading medium_severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading medium_severity: %v", err)
		}
	}

	return nil
}

func flattenLogThreatWeightIpsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogThreatWeightIpsCriticalSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightIpsHighSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightIpsInfoSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightIpsLowSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightIpsMediumSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogThreatWeightIps(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("critical_severity"); ok || d.HasChange("critical_severity") {
		t, err := expandLogThreatWeightIpsCriticalSeverity2edl(d, v, "critical_severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["critical-severity"] = t
		}
	}

	if v, ok := d.GetOk("high_severity"); ok || d.HasChange("high_severity") {
		t, err := expandLogThreatWeightIpsHighSeverity2edl(d, v, "high_severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["high-severity"] = t
		}
	}

	if v, ok := d.GetOk("info_severity"); ok || d.HasChange("info_severity") {
		t, err := expandLogThreatWeightIpsInfoSeverity2edl(d, v, "info_severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["info-severity"] = t
		}
	}

	if v, ok := d.GetOk("low_severity"); ok || d.HasChange("low_severity") {
		t, err := expandLogThreatWeightIpsLowSeverity2edl(d, v, "low_severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["low-severity"] = t
		}
	}

	if v, ok := d.GetOk("medium_severity"); ok || d.HasChange("medium_severity") {
		t, err := expandLogThreatWeightIpsMediumSeverity2edl(d, v, "medium_severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["medium-severity"] = t
		}
	}

	return &obj, nil
}
