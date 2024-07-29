// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configuration method to edit Simple Network Management Protocol (SNMP) trap threshold values.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerManagedSwitchSnmpTrapThreshold() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerManagedSwitchSnmpTrapThresholdUpdate,
		Read:   resourceSwitchControllerManagedSwitchSnmpTrapThresholdRead,
		Update: resourceSwitchControllerManagedSwitchSnmpTrapThresholdUpdate,
		Delete: resourceSwitchControllerManagedSwitchSnmpTrapThresholdDelete,

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
			"managed_switch": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"trap_high_cpu_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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

func resourceSwitchControllerManagedSwitchSnmpTrapThresholdUpdate(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	obj, err := getObjectSwitchControllerManagedSwitchSnmpTrapThreshold(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchSnmpTrapThreshold resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerManagedSwitchSnmpTrapThreshold(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchSnmpTrapThreshold resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SwitchControllerManagedSwitchSnmpTrapThreshold")

	return resourceSwitchControllerManagedSwitchSnmpTrapThresholdRead(d, m)
}

func resourceSwitchControllerManagedSwitchSnmpTrapThresholdDelete(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	err = c.DeleteSwitchControllerManagedSwitchSnmpTrapThreshold(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerManagedSwitchSnmpTrapThreshold resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerManagedSwitchSnmpTrapThresholdRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	managed_switch := d.Get("managed_switch").(string)
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
	if managed_switch == "" {
		managed_switch = importOptionChecking(m.(*FortiClient).Cfg, "managed_switch")
		if managed_switch == "" {
			return fmt.Errorf("Parameter managed_switch is missing")
		}
		if err = d.Set("managed_switch", managed_switch); err != nil {
			return fmt.Errorf("Error set params managed_switch: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	o, err := c.ReadSwitchControllerManagedSwitchSnmpTrapThreshold(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchSnmpTrapThreshold resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerManagedSwitchSnmpTrapThreshold(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchSnmpTrapThreshold resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerManagedSwitchSnmpTrapThresholdTrapHighCpuThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpTrapThresholdTrapLogFullThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpTrapThresholdTrapLowMemoryThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerManagedSwitchSnmpTrapThreshold(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("trap_high_cpu_threshold", flattenSwitchControllerManagedSwitchSnmpTrapThresholdTrapHighCpuThreshold2edl(o["trap-high-cpu-threshold"], d, "trap_high_cpu_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-high-cpu-threshold"], "SwitchControllerManagedSwitchSnmpTrapThreshold-TrapHighCpuThreshold"); ok {
			if err = d.Set("trap_high_cpu_threshold", vv); err != nil {
				return fmt.Errorf("Error reading trap_high_cpu_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_high_cpu_threshold: %v", err)
		}
	}

	if err = d.Set("trap_log_full_threshold", flattenSwitchControllerManagedSwitchSnmpTrapThresholdTrapLogFullThreshold2edl(o["trap-log-full-threshold"], d, "trap_log_full_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-log-full-threshold"], "SwitchControllerManagedSwitchSnmpTrapThreshold-TrapLogFullThreshold"); ok {
			if err = d.Set("trap_log_full_threshold", vv); err != nil {
				return fmt.Errorf("Error reading trap_log_full_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_log_full_threshold: %v", err)
		}
	}

	if err = d.Set("trap_low_memory_threshold", flattenSwitchControllerManagedSwitchSnmpTrapThresholdTrapLowMemoryThreshold2edl(o["trap-low-memory-threshold"], d, "trap_low_memory_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-low-memory-threshold"], "SwitchControllerManagedSwitchSnmpTrapThreshold-TrapLowMemoryThreshold"); ok {
			if err = d.Set("trap_low_memory_threshold", vv); err != nil {
				return fmt.Errorf("Error reading trap_low_memory_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_low_memory_threshold: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerManagedSwitchSnmpTrapThresholdFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerManagedSwitchSnmpTrapThresholdTrapHighCpuThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpTrapThresholdTrapLogFullThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpTrapThresholdTrapLowMemoryThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerManagedSwitchSnmpTrapThreshold(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("trap_high_cpu_threshold"); ok || d.HasChange("trap_high_cpu_threshold") {
		t, err := expandSwitchControllerManagedSwitchSnmpTrapThresholdTrapHighCpuThreshold2edl(d, v, "trap_high_cpu_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-high-cpu-threshold"] = t
		}
	}

	if v, ok := d.GetOk("trap_log_full_threshold"); ok || d.HasChange("trap_log_full_threshold") {
		t, err := expandSwitchControllerManagedSwitchSnmpTrapThresholdTrapLogFullThreshold2edl(d, v, "trap_log_full_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-log-full-threshold"] = t
		}
	}

	if v, ok := d.GetOk("trap_low_memory_threshold"); ok || d.HasChange("trap_low_memory_threshold") {
		t, err := expandSwitchControllerManagedSwitchSnmpTrapThresholdTrapLowMemoryThreshold2edl(d, v, "trap_low_memory_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-low-memory-threshold"] = t
		}
	}

	return &obj, nil
}
