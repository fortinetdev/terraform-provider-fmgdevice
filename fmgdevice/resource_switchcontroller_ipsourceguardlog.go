// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch ip source guard log.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerIpSourceGuardLog() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerIpSourceGuardLogUpdate,
		Read:   resourceSwitchControllerIpSourceGuardLogRead,
		Update: resourceSwitchControllerIpSourceGuardLogUpdate,
		Delete: resourceSwitchControllerIpSourceGuardLogDelete,

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
			"log_violations": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"violation_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSwitchControllerIpSourceGuardLogUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerIpSourceGuardLog(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerIpSourceGuardLog resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSwitchControllerIpSourceGuardLog(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerIpSourceGuardLog resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SwitchControllerIpSourceGuardLog")

	return resourceSwitchControllerIpSourceGuardLogRead(d, m)
}

func resourceSwitchControllerIpSourceGuardLogDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerIpSourceGuardLog(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerIpSourceGuardLog resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerIpSourceGuardLogRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerIpSourceGuardLog(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SwitchControllerIpSourceGuardLog resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerIpSourceGuardLog(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerIpSourceGuardLog resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerIpSourceGuardLogLogViolations(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerIpSourceGuardLogViolationTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerIpSourceGuardLog(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("log_violations", flattenSwitchControllerIpSourceGuardLogLogViolations(o["log-violations"], d, "log_violations")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-violations"], "SwitchControllerIpSourceGuardLog-LogViolations"); ok {
			if err = d.Set("log_violations", vv); err != nil {
				return fmt.Errorf("Error reading log_violations: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_violations: %v", err)
		}
	}

	if err = d.Set("violation_timer", flattenSwitchControllerIpSourceGuardLogViolationTimer(o["violation-timer"], d, "violation_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["violation-timer"], "SwitchControllerIpSourceGuardLog-ViolationTimer"); ok {
			if err = d.Set("violation_timer", vv); err != nil {
				return fmt.Errorf("Error reading violation_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading violation_timer: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerIpSourceGuardLogFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerIpSourceGuardLogLogViolations(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerIpSourceGuardLogViolationTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerIpSourceGuardLog(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("log_violations"); ok || d.HasChange("log_violations") {
		t, err := expandSwitchControllerIpSourceGuardLogLogViolations(d, v, "log_violations")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-violations"] = t
		}
	}

	if v, ok := d.GetOk("violation_timer"); ok || d.HasChange("violation_timer") {
		t, err := expandSwitchControllerIpSourceGuardLogViolationTimer(d, v, "violation_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["violation-timer"] = t
		}
	}

	return &obj, nil
}
