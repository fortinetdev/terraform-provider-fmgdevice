// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SNMP system info configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSnmpSysinfo() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSnmpSysinfoUpdate,
		Read:   resourceSystemSnmpSysinfoRead,
		Update: resourceSystemSnmpSysinfoUpdate,
		Delete: resourceSystemSnmpSysinfoDelete,

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
			"append_index": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"contact_info": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"engine_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"engine_id_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trap_free_memory_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trap_freeable_memory_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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

func resourceSystemSnmpSysinfoUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemSnmpSysinfo(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpSysinfo resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSnmpSysinfo(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpSysinfo resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemSnmpSysinfo")

	return resourceSystemSnmpSysinfoRead(d, m)
}

func resourceSystemSnmpSysinfoDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	err = c.DeleteSystemSnmpSysinfo(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSnmpSysinfo resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSnmpSysinfoRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSnmpSysinfo(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpSysinfo resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSnmpSysinfo(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpSysinfo resource from API: %v", err)
	}
	return nil
}

func flattenSystemSnmpSysinfoAppendIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoContactInfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoEngineId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoEngineIdType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoTrapFreeMemoryThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoTrapFreeableMemoryThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoTrapHighCpuThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoTrapLogFullThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoTrapLowMemoryThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSnmpSysinfo(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("append_index", flattenSystemSnmpSysinfoAppendIndex(o["append-index"], d, "append_index")); err != nil {
		if vv, ok := fortiAPIPatch(o["append-index"], "SystemSnmpSysinfo-AppendIndex"); ok {
			if err = d.Set("append_index", vv); err != nil {
				return fmt.Errorf("Error reading append_index: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading append_index: %v", err)
		}
	}

	if err = d.Set("contact_info", flattenSystemSnmpSysinfoContactInfo(o["contact-info"], d, "contact_info")); err != nil {
		if vv, ok := fortiAPIPatch(o["contact-info"], "SystemSnmpSysinfo-ContactInfo"); ok {
			if err = d.Set("contact_info", vv); err != nil {
				return fmt.Errorf("Error reading contact_info: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading contact_info: %v", err)
		}
	}

	if err = d.Set("description", flattenSystemSnmpSysinfoDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SystemSnmpSysinfo-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("engine_id", flattenSystemSnmpSysinfoEngineId(o["engine-id"], d, "engine_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["engine-id"], "SystemSnmpSysinfo-EngineId"); ok {
			if err = d.Set("engine_id", vv); err != nil {
				return fmt.Errorf("Error reading engine_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading engine_id: %v", err)
		}
	}

	if err = d.Set("engine_id_type", flattenSystemSnmpSysinfoEngineIdType(o["engine-id-type"], d, "engine_id_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["engine-id-type"], "SystemSnmpSysinfo-EngineIdType"); ok {
			if err = d.Set("engine_id_type", vv); err != nil {
				return fmt.Errorf("Error reading engine_id_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading engine_id_type: %v", err)
		}
	}

	if err = d.Set("location", flattenSystemSnmpSysinfoLocation(o["location"], d, "location")); err != nil {
		if vv, ok := fortiAPIPatch(o["location"], "SystemSnmpSysinfo-Location"); ok {
			if err = d.Set("location", vv); err != nil {
				return fmt.Errorf("Error reading location: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading location: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemSnmpSysinfoStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemSnmpSysinfo-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("trap_free_memory_threshold", flattenSystemSnmpSysinfoTrapFreeMemoryThreshold(o["trap-free-memory-threshold"], d, "trap_free_memory_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-free-memory-threshold"], "SystemSnmpSysinfo-TrapFreeMemoryThreshold"); ok {
			if err = d.Set("trap_free_memory_threshold", vv); err != nil {
				return fmt.Errorf("Error reading trap_free_memory_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_free_memory_threshold: %v", err)
		}
	}

	if err = d.Set("trap_freeable_memory_threshold", flattenSystemSnmpSysinfoTrapFreeableMemoryThreshold(o["trap-freeable-memory-threshold"], d, "trap_freeable_memory_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-freeable-memory-threshold"], "SystemSnmpSysinfo-TrapFreeableMemoryThreshold"); ok {
			if err = d.Set("trap_freeable_memory_threshold", vv); err != nil {
				return fmt.Errorf("Error reading trap_freeable_memory_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_freeable_memory_threshold: %v", err)
		}
	}

	if err = d.Set("trap_high_cpu_threshold", flattenSystemSnmpSysinfoTrapHighCpuThreshold(o["trap-high-cpu-threshold"], d, "trap_high_cpu_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-high-cpu-threshold"], "SystemSnmpSysinfo-TrapHighCpuThreshold"); ok {
			if err = d.Set("trap_high_cpu_threshold", vv); err != nil {
				return fmt.Errorf("Error reading trap_high_cpu_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_high_cpu_threshold: %v", err)
		}
	}

	if err = d.Set("trap_log_full_threshold", flattenSystemSnmpSysinfoTrapLogFullThreshold(o["trap-log-full-threshold"], d, "trap_log_full_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-log-full-threshold"], "SystemSnmpSysinfo-TrapLogFullThreshold"); ok {
			if err = d.Set("trap_log_full_threshold", vv); err != nil {
				return fmt.Errorf("Error reading trap_log_full_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_log_full_threshold: %v", err)
		}
	}

	if err = d.Set("trap_low_memory_threshold", flattenSystemSnmpSysinfoTrapLowMemoryThreshold(o["trap-low-memory-threshold"], d, "trap_low_memory_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-low-memory-threshold"], "SystemSnmpSysinfo-TrapLowMemoryThreshold"); ok {
			if err = d.Set("trap_low_memory_threshold", vv); err != nil {
				return fmt.Errorf("Error reading trap_low_memory_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_low_memory_threshold: %v", err)
		}
	}

	return nil
}

func flattenSystemSnmpSysinfoFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSnmpSysinfoAppendIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoContactInfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoEngineId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoEngineIdType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoLocation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoTrapFreeMemoryThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoTrapFreeableMemoryThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoTrapHighCpuThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoTrapLogFullThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoTrapLowMemoryThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSnmpSysinfo(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("append_index"); ok || d.HasChange("append_index") {
		t, err := expandSystemSnmpSysinfoAppendIndex(d, v, "append_index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["append-index"] = t
		}
	}

	if v, ok := d.GetOk("contact_info"); ok || d.HasChange("contact_info") {
		t, err := expandSystemSnmpSysinfoContactInfo(d, v, "contact_info")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["contact-info"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSystemSnmpSysinfoDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("engine_id"); ok || d.HasChange("engine_id") {
		t, err := expandSystemSnmpSysinfoEngineId(d, v, "engine_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["engine-id"] = t
		}
	}

	if v, ok := d.GetOk("engine_id_type"); ok || d.HasChange("engine_id_type") {
		t, err := expandSystemSnmpSysinfoEngineIdType(d, v, "engine_id_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["engine-id-type"] = t
		}
	}

	if v, ok := d.GetOk("location"); ok || d.HasChange("location") {
		t, err := expandSystemSnmpSysinfoLocation(d, v, "location")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["location"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemSnmpSysinfoStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("trap_free_memory_threshold"); ok || d.HasChange("trap_free_memory_threshold") {
		t, err := expandSystemSnmpSysinfoTrapFreeMemoryThreshold(d, v, "trap_free_memory_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-free-memory-threshold"] = t
		}
	}

	if v, ok := d.GetOk("trap_freeable_memory_threshold"); ok || d.HasChange("trap_freeable_memory_threshold") {
		t, err := expandSystemSnmpSysinfoTrapFreeableMemoryThreshold(d, v, "trap_freeable_memory_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-freeable-memory-threshold"] = t
		}
	}

	if v, ok := d.GetOk("trap_high_cpu_threshold"); ok || d.HasChange("trap_high_cpu_threshold") {
		t, err := expandSystemSnmpSysinfoTrapHighCpuThreshold(d, v, "trap_high_cpu_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-high-cpu-threshold"] = t
		}
	}

	if v, ok := d.GetOk("trap_log_full_threshold"); ok || d.HasChange("trap_log_full_threshold") {
		t, err := expandSystemSnmpSysinfoTrapLogFullThreshold(d, v, "trap_log_full_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-log-full-threshold"] = t
		}
	}

	if v, ok := d.GetOk("trap_low_memory_threshold"); ok || d.HasChange("trap_low_memory_threshold") {
		t, err := expandSystemSnmpSysinfoTrapLowMemoryThreshold(d, v, "trap_low_memory_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-low-memory-threshold"] = t
		}
	}

	return &obj, nil
}
