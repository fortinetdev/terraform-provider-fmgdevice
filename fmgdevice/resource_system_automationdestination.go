// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Automation destinations.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAutomationDestination() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutomationDestinationCreate,
		Read:   resourceSystemAutomationDestinationRead,
		Update: resourceSystemAutomationDestinationUpdate,
		Delete: resourceSystemAutomationDestinationDelete,

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
			"destination": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ha_group_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemAutomationDestinationCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemAutomationDestination(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationDestination resource while getting object: %v", err)
	}

	_, err = c.CreateSystemAutomationDestination(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationDestination resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemAutomationDestinationRead(d, m)
}

func resourceSystemAutomationDestinationUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemAutomationDestination(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationDestination resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAutomationDestination(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationDestination resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemAutomationDestinationRead(d, m)
}

func resourceSystemAutomationDestinationDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemAutomationDestination(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAutomationDestination resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutomationDestinationRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemAutomationDestination(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationDestination resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutomationDestination(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationDestination resource from API: %v", err)
	}
	return nil
}

func flattenSystemAutomationDestinationDestination(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAutomationDestinationHaGroupId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationDestinationName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationDestinationType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAutomationDestination(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("destination", flattenSystemAutomationDestinationDestination(o["destination"], d, "destination")); err != nil {
		if vv, ok := fortiAPIPatch(o["destination"], "SystemAutomationDestination-Destination"); ok {
			if err = d.Set("destination", vv); err != nil {
				return fmt.Errorf("Error reading destination: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading destination: %v", err)
		}
	}

	if err = d.Set("ha_group_id", flattenSystemAutomationDestinationHaGroupId(o["ha-group-id"], d, "ha_group_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-group-id"], "SystemAutomationDestination-HaGroupId"); ok {
			if err = d.Set("ha_group_id", vv); err != nil {
				return fmt.Errorf("Error reading ha_group_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_group_id: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemAutomationDestinationName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemAutomationDestination-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemAutomationDestinationType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "SystemAutomationDestination-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenSystemAutomationDestinationFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAutomationDestinationDestination(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAutomationDestinationHaGroupId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationDestinationName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationDestinationType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAutomationDestination(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("destination"); ok || d.HasChange("destination") {
		t, err := expandSystemAutomationDestinationDestination(d, v, "destination")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["destination"] = t
		}
	}

	if v, ok := d.GetOk("ha_group_id"); ok || d.HasChange("ha_group_id") {
		t, err := expandSystemAutomationDestinationHaGroupId(d, v, "ha_group_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-group-id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemAutomationDestinationName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandSystemAutomationDestinationType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
