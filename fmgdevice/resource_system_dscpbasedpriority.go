// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure DSCP based priority table.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemDscpBasedPriority() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDscpBasedPriorityCreate,
		Read:   resourceSystemDscpBasedPriorityRead,
		Update: resourceSystemDscpBasedPriorityUpdate,
		Delete: resourceSystemDscpBasedPriorityDelete,

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
			"ds": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemDscpBasedPriorityCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemDscpBasedPriority(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemDscpBasedPriority resource while getting object: %v", err)
	}

	_, err = c.CreateSystemDscpBasedPriority(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemDscpBasedPriority resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemDscpBasedPriorityRead(d, m)
}

func resourceSystemDscpBasedPriorityUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemDscpBasedPriority(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemDscpBasedPriority resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemDscpBasedPriority(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemDscpBasedPriority resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemDscpBasedPriorityRead(d, m)
}

func resourceSystemDscpBasedPriorityDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemDscpBasedPriority(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDscpBasedPriority resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDscpBasedPriorityRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemDscpBasedPriority(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemDscpBasedPriority resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDscpBasedPriority(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemDscpBasedPriority resource from API: %v", err)
	}
	return nil
}

func flattenSystemDscpBasedPriorityDs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDscpBasedPriorityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDscpBasedPriorityPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemDscpBasedPriority(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ds", flattenSystemDscpBasedPriorityDs(o["ds"], d, "ds")); err != nil {
		if vv, ok := fortiAPIPatch(o["ds"], "SystemDscpBasedPriority-Ds"); ok {
			if err = d.Set("ds", vv); err != nil {
				return fmt.Errorf("Error reading ds: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ds: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemDscpBasedPriorityId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemDscpBasedPriority-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("priority", flattenSystemDscpBasedPriorityPriority(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "SystemDscpBasedPriority-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	return nil
}

func flattenSystemDscpBasedPriorityFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemDscpBasedPriorityDs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDscpBasedPriorityId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDscpBasedPriorityPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDscpBasedPriority(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ds"); ok || d.HasChange("ds") {
		t, err := expandSystemDscpBasedPriorityDs(d, v, "ds")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ds"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemDscpBasedPriorityId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandSystemDscpBasedPriorityPriority(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	return &obj, nil
}
