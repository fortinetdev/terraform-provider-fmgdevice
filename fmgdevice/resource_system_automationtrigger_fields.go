// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Customized trigger field settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAutomationTriggerFields() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutomationTriggerFieldsCreate,
		Read:   resourceSystemAutomationTriggerFieldsRead,
		Update: resourceSystemAutomationTriggerFieldsUpdate,
		Delete: resourceSystemAutomationTriggerFieldsDelete,

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
			"automation_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemAutomationTriggerFieldsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	automation_trigger := d.Get("automation_trigger").(string)
	paradict["device"] = device_name
	paradict["automation_trigger"] = automation_trigger

	obj, err := getObjectSystemAutomationTriggerFields(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationTriggerFields resource while getting object: %v", err)
	}

	_, err = c.CreateSystemAutomationTriggerFields(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationTriggerFields resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemAutomationTriggerFieldsRead(d, m)
}

func resourceSystemAutomationTriggerFieldsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	automation_trigger := d.Get("automation_trigger").(string)
	paradict["device"] = device_name
	paradict["automation_trigger"] = automation_trigger

	obj, err := getObjectSystemAutomationTriggerFields(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationTriggerFields resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAutomationTriggerFields(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationTriggerFields resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemAutomationTriggerFieldsRead(d, m)
}

func resourceSystemAutomationTriggerFieldsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	automation_trigger := d.Get("automation_trigger").(string)
	paradict["device"] = device_name
	paradict["automation_trigger"] = automation_trigger

	err = c.DeleteSystemAutomationTriggerFields(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAutomationTriggerFields resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutomationTriggerFieldsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	automation_trigger := d.Get("automation_trigger").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if automation_trigger == "" {
		automation_trigger = importOptionChecking(m.(*FortiClient).Cfg, "automation_trigger")
		if automation_trigger == "" {
			return fmt.Errorf("Parameter automation_trigger is missing")
		}
		if err = d.Set("automation_trigger", automation_trigger); err != nil {
			return fmt.Errorf("Error set params automation_trigger: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["automation_trigger"] = automation_trigger

	o, err := c.ReadSystemAutomationTriggerFields(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationTriggerFields resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutomationTriggerFields(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationTriggerFields resource from API: %v", err)
	}
	return nil
}

func flattenSystemAutomationTriggerFieldsId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerFieldsName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerFieldsValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAutomationTriggerFields(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenSystemAutomationTriggerFieldsId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemAutomationTriggerFields-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemAutomationTriggerFieldsName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemAutomationTriggerFields-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("value", flattenSystemAutomationTriggerFieldsValue2edl(o["value"], d, "value")); err != nil {
		if vv, ok := fortiAPIPatch(o["value"], "SystemAutomationTriggerFields-Value"); ok {
			if err = d.Set("value", vv); err != nil {
				return fmt.Errorf("Error reading value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	return nil
}

func flattenSystemAutomationTriggerFieldsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAutomationTriggerFieldsId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerFieldsName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerFieldsValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAutomationTriggerFields(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemAutomationTriggerFieldsId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemAutomationTriggerFieldsName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok || d.HasChange("value") {
		t, err := expandSystemAutomationTriggerFieldsValue2edl(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
