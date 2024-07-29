// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Worker blade used by this group.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLoadBalanceSettingWorkers() *schema.Resource {
	return &schema.Resource{
		Create: resourceLoadBalanceSettingWorkersCreate,
		Read:   resourceLoadBalanceSettingWorkersRead,
		Update: resourceLoadBalanceSettingWorkersUpdate,
		Delete: resourceLoadBalanceSettingWorkersDelete,

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
			"slot": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLoadBalanceSettingWorkersCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectLoadBalanceSettingWorkers(d)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceSettingWorkers resource while getting object: %v", err)
	}

	_, err = c.CreateLoadBalanceSettingWorkers(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceSettingWorkers resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "slot")))

	return resourceLoadBalanceSettingWorkersRead(d, m)
}

func resourceLoadBalanceSettingWorkersUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLoadBalanceSettingWorkers(d)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceSettingWorkers resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceSettingWorkers(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceSettingWorkers resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "slot")))

	return resourceLoadBalanceSettingWorkersRead(d, m)
}

func resourceLoadBalanceSettingWorkersDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteLoadBalanceSettingWorkers(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceSettingWorkers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLoadBalanceSettingWorkersRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLoadBalanceSettingWorkers(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceSettingWorkers resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceSettingWorkers(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceSettingWorkers resource from API: %v", err)
	}
	return nil
}

func flattenLoadBalanceSettingWorkersSlot2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLoadBalanceSettingWorkersStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLoadBalanceSettingWorkersWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLoadBalanceSettingWorkers(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("slot", flattenLoadBalanceSettingWorkersSlot2edl(o["slot"], d, "slot")); err != nil {
		if vv, ok := fortiAPIPatch(o["slot"], "LoadBalanceSettingWorkers-Slot"); ok {
			if err = d.Set("slot", vv); err != nil {
				return fmt.Errorf("Error reading slot: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading slot: %v", err)
		}
	}

	if err = d.Set("status", flattenLoadBalanceSettingWorkersStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "LoadBalanceSettingWorkers-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("weight", flattenLoadBalanceSettingWorkersWeight2edl(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "LoadBalanceSettingWorkers-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	return nil
}

func flattenLoadBalanceSettingWorkersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLoadBalanceSettingWorkersSlot2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceSettingWorkersStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLoadBalanceSettingWorkersWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceSettingWorkers(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("slot"); ok || d.HasChange("slot") {
		t, err := expandLoadBalanceSettingWorkersSlot2edl(d, v, "slot")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["slot"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandLoadBalanceSettingWorkersStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandLoadBalanceSettingWorkersWeight2edl(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	return &obj, nil
}
