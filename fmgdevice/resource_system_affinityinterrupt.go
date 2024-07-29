// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure interrupt affinity.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAffinityInterrupt() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAffinityInterruptCreate,
		Read:   resourceSystemAffinityInterruptRead,
		Update: resourceSystemAffinityInterruptUpdate,
		Delete: resourceSystemAffinityInterruptDelete,

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
			"affinity_cpumask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"default_affinity_cpumask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"interrupt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemAffinityInterruptCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemAffinityInterrupt(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAffinityInterrupt resource while getting object: %v", err)
	}

	_, err = c.CreateSystemAffinityInterrupt(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemAffinityInterrupt resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemAffinityInterruptRead(d, m)
}

func resourceSystemAffinityInterruptUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemAffinityInterrupt(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAffinityInterrupt resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAffinityInterrupt(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemAffinityInterrupt resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemAffinityInterruptRead(d, m)
}

func resourceSystemAffinityInterruptDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemAffinityInterrupt(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAffinityInterrupt resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAffinityInterruptRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemAffinityInterrupt(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemAffinityInterrupt resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAffinityInterrupt(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAffinityInterrupt resource from API: %v", err)
	}
	return nil
}

func flattenSystemAffinityInterruptAffinityCpumask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAffinityInterruptDefaultAffinityCpumask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAffinityInterruptId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAffinityInterruptInterrupt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAffinityInterrupt(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("affinity_cpumask", flattenSystemAffinityInterruptAffinityCpumask(o["affinity-cpumask"], d, "affinity_cpumask")); err != nil {
		if vv, ok := fortiAPIPatch(o["affinity-cpumask"], "SystemAffinityInterrupt-AffinityCpumask"); ok {
			if err = d.Set("affinity_cpumask", vv); err != nil {
				return fmt.Errorf("Error reading affinity_cpumask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading affinity_cpumask: %v", err)
		}
	}

	if err = d.Set("default_affinity_cpumask", flattenSystemAffinityInterruptDefaultAffinityCpumask(o["default-affinity-cpumask"], d, "default_affinity_cpumask")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-affinity-cpumask"], "SystemAffinityInterrupt-DefaultAffinityCpumask"); ok {
			if err = d.Set("default_affinity_cpumask", vv); err != nil {
				return fmt.Errorf("Error reading default_affinity_cpumask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_affinity_cpumask: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemAffinityInterruptId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemAffinityInterrupt-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("interrupt", flattenSystemAffinityInterruptInterrupt(o["interrupt"], d, "interrupt")); err != nil {
		if vv, ok := fortiAPIPatch(o["interrupt"], "SystemAffinityInterrupt-Interrupt"); ok {
			if err = d.Set("interrupt", vv); err != nil {
				return fmt.Errorf("Error reading interrupt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interrupt: %v", err)
		}
	}

	return nil
}

func flattenSystemAffinityInterruptFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAffinityInterruptAffinityCpumask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAffinityInterruptDefaultAffinityCpumask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAffinityInterruptId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAffinityInterruptInterrupt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAffinityInterrupt(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("affinity_cpumask"); ok || d.HasChange("affinity_cpumask") {
		t, err := expandSystemAffinityInterruptAffinityCpumask(d, v, "affinity_cpumask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["affinity-cpumask"] = t
		}
	}

	if v, ok := d.GetOk("default_affinity_cpumask"); ok || d.HasChange("default_affinity_cpumask") {
		t, err := expandSystemAffinityInterruptDefaultAffinityCpumask(d, v, "default_affinity_cpumask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-affinity-cpumask"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemAffinityInterruptId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("interrupt"); ok || d.HasChange("interrupt") {
		t, err := expandSystemAffinityInterruptInterrupt(d, v, "interrupt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interrupt"] = t
		}
	}

	return &obj, nil
}
