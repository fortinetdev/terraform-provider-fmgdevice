// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure update schedule.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAutoupdateSchedule() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutoupdateScheduleUpdate,
		Read:   resourceSystemAutoupdateScheduleRead,
		Update: resourceSystemAutoupdateScheduleUpdate,
		Delete: resourceSystemAutoupdateScheduleDelete,

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
			"day": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"frequency": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemAutoupdateScheduleUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemAutoupdateSchedule(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoupdateSchedule resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAutoupdateSchedule(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoupdateSchedule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemAutoupdateSchedule")

	return resourceSystemAutoupdateScheduleRead(d, m)
}

func resourceSystemAutoupdateScheduleDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemAutoupdateSchedule(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAutoupdateSchedule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutoupdateScheduleRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemAutoupdateSchedule(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoupdateSchedule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutoupdateSchedule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoupdateSchedule resource from API: %v", err)
	}
	return nil
}

func flattenSystemAutoupdateScheduleDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoupdateScheduleFrequency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoupdateScheduleStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoupdateScheduleTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAutoupdateSchedule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("day", flattenSystemAutoupdateScheduleDay(o["day"], d, "day")); err != nil {
		if vv, ok := fortiAPIPatch(o["day"], "SystemAutoupdateSchedule-Day"); ok {
			if err = d.Set("day", vv); err != nil {
				return fmt.Errorf("Error reading day: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading day: %v", err)
		}
	}

	if err = d.Set("frequency", flattenSystemAutoupdateScheduleFrequency(o["frequency"], d, "frequency")); err != nil {
		if vv, ok := fortiAPIPatch(o["frequency"], "SystemAutoupdateSchedule-Frequency"); ok {
			if err = d.Set("frequency", vv); err != nil {
				return fmt.Errorf("Error reading frequency: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading frequency: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemAutoupdateScheduleStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemAutoupdateSchedule-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("time", flattenSystemAutoupdateScheduleTime(o["time"], d, "time")); err != nil {
		if vv, ok := fortiAPIPatch(o["time"], "SystemAutoupdateSchedule-Time"); ok {
			if err = d.Set("time", vv); err != nil {
				return fmt.Errorf("Error reading time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading time: %v", err)
		}
	}

	return nil
}

func flattenSystemAutoupdateScheduleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAutoupdateScheduleDay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoupdateScheduleFrequency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoupdateScheduleStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoupdateScheduleTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAutoupdateSchedule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("day"); ok || d.HasChange("day") {
		t, err := expandSystemAutoupdateScheduleDay(d, v, "day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["day"] = t
		}
	}

	if v, ok := d.GetOk("frequency"); ok || d.HasChange("frequency") {
		t, err := expandSystemAutoupdateScheduleFrequency(d, v, "frequency")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["frequency"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemAutoupdateScheduleStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("time"); ok || d.HasChange("time") {
		t, err := expandSystemAutoupdateScheduleTime(d, v, "time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["time"] = t
		}
	}

	return &obj, nil
}
