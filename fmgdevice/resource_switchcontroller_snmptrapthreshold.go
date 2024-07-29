// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch SNMP trap threshold values globally.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerSnmpTrapThreshold() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerSnmpTrapThresholdUpdate,
		Read:   resourceSwitchControllerSnmpTrapThresholdRead,
		Update: resourceSwitchControllerSnmpTrapThresholdUpdate,
		Delete: resourceSwitchControllerSnmpTrapThresholdDelete,

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
			"trap_high_cpu_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trap_log_full_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trap_low_memory_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerSnmpTrapThresholdUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	obj, err := getObjectSwitchControllerSnmpTrapThreshold(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSnmpTrapThreshold resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerSnmpTrapThreshold(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSnmpTrapThreshold resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SwitchControllerSnmpTrapThreshold")

	return resourceSwitchControllerSnmpTrapThresholdRead(d, m)
}

func resourceSwitchControllerSnmpTrapThresholdDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	err = c.DeleteSwitchControllerSnmpTrapThreshold(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerSnmpTrapThreshold resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerSnmpTrapThresholdRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerSnmpTrapThreshold(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSnmpTrapThreshold resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerSnmpTrapThreshold(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSnmpTrapThreshold resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerSnmpTrapThresholdTrapHighCpuThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSnmpTrapThresholdTrapLogFullThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSnmpTrapThresholdTrapLowMemoryThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerSnmpTrapThreshold(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("trap_high_cpu_threshold", flattenSwitchControllerSnmpTrapThresholdTrapHighCpuThreshold(o["trap-high-cpu-threshold"], d, "trap_high_cpu_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-high-cpu-threshold"], "SwitchControllerSnmpTrapThreshold-TrapHighCpuThreshold"); ok {
			if err = d.Set("trap_high_cpu_threshold", vv); err != nil {
				return fmt.Errorf("Error reading trap_high_cpu_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_high_cpu_threshold: %v", err)
		}
	}

	if err = d.Set("trap_log_full_threshold", flattenSwitchControllerSnmpTrapThresholdTrapLogFullThreshold(o["trap-log-full-threshold"], d, "trap_log_full_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-log-full-threshold"], "SwitchControllerSnmpTrapThreshold-TrapLogFullThreshold"); ok {
			if err = d.Set("trap_log_full_threshold", vv); err != nil {
				return fmt.Errorf("Error reading trap_log_full_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_log_full_threshold: %v", err)
		}
	}

	if err = d.Set("trap_low_memory_threshold", flattenSwitchControllerSnmpTrapThresholdTrapLowMemoryThreshold(o["trap-low-memory-threshold"], d, "trap_low_memory_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-low-memory-threshold"], "SwitchControllerSnmpTrapThreshold-TrapLowMemoryThreshold"); ok {
			if err = d.Set("trap_low_memory_threshold", vv); err != nil {
				return fmt.Errorf("Error reading trap_low_memory_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_low_memory_threshold: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerSnmpTrapThresholdFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerSnmpTrapThresholdTrapHighCpuThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpTrapThresholdTrapLogFullThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpTrapThresholdTrapLowMemoryThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerSnmpTrapThreshold(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("trap_high_cpu_threshold"); ok || d.HasChange("trap_high_cpu_threshold") {
		t, err := expandSwitchControllerSnmpTrapThresholdTrapHighCpuThreshold(d, v, "trap_high_cpu_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-high-cpu-threshold"] = t
		}
	}

	if v, ok := d.GetOk("trap_log_full_threshold"); ok || d.HasChange("trap_log_full_threshold") {
		t, err := expandSwitchControllerSnmpTrapThresholdTrapLogFullThreshold(d, v, "trap_log_full_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-log-full-threshold"] = t
		}
	}

	if v, ok := d.GetOk("trap_low_memory_threshold"); ok || d.HasChange("trap_low_memory_threshold") {
		t, err := expandSwitchControllerSnmpTrapThresholdTrapLowMemoryThreshold(d, v, "trap_low_memory_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-low-memory-threshold"] = t
		}
	}

	return &obj, nil
}
