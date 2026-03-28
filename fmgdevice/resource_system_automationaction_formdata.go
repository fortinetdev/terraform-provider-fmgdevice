// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Form data parts for content type multipart/form-data.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAutomationActionFormData() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutomationActionFormDataCreate,
		Read:   resourceSystemAutomationActionFormDataRead,
		Update: resourceSystemAutomationActionFormDataUpdate,
		Delete: resourceSystemAutomationActionFormDataDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

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
			"automation_action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"key": &schema.Schema{
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

func resourceSystemAutomationActionFormDataCreate(d *schema.ResourceData, m interface{}) error {
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
	automation_action := d.Get("automation_action").(string)
	paradict["device"] = device_name
	paradict["automation_action"] = automation_action

	obj, err := getObjectSystemAutomationActionFormData(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationActionFormData resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemAutomationActionFormData(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemAutomationActionFormData(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemAutomationActionFormData resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateSystemAutomationActionFormData(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemAutomationActionFormData resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceSystemAutomationActionFormDataRead(d, m)
			} else {
				return fmt.Errorf("Error creating SystemAutomationActionFormData resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemAutomationActionFormDataRead(d, m)
}

func resourceSystemAutomationActionFormDataUpdate(d *schema.ResourceData, m interface{}) error {
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
	automation_action := d.Get("automation_action").(string)
	paradict["device"] = device_name
	paradict["automation_action"] = automation_action

	obj, err := getObjectSystemAutomationActionFormData(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationActionFormData resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemAutomationActionFormData(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationActionFormData resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemAutomationActionFormDataRead(d, m)
}

func resourceSystemAutomationActionFormDataDelete(d *schema.ResourceData, m interface{}) error {
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
	automation_action := d.Get("automation_action").(string)
	paradict["device"] = device_name
	paradict["automation_action"] = automation_action

	wsParams["adom"] = adomv

	err = c.DeleteSystemAutomationActionFormData(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAutomationActionFormData resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutomationActionFormDataRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	automation_action := d.Get("automation_action").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if automation_action == "" {
		automation_action = importOptionChecking(m.(*FortiClient).Cfg, "automation_action")
		if automation_action == "" {
			return fmt.Errorf("Parameter automation_action is missing")
		}
		if err = d.Set("automation_action", automation_action); err != nil {
			return fmt.Errorf("Error set params automation_action: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["automation_action"] = automation_action

	o, err := c.ReadSystemAutomationActionFormData(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemAutomationActionFormData resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutomationActionFormData(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationActionFormData resource from API: %v", err)
	}
	return nil
}

func flattenSystemAutomationActionFormDataId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionFormDataKey2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionFormDataValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAutomationActionFormData(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenSystemAutomationActionFormDataId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemAutomationActionFormData-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("key", flattenSystemAutomationActionFormDataKey2edl(o["key"], d, "key")); err != nil {
		if vv, ok := fortiAPIPatch(o["key"], "SystemAutomationActionFormData-Key"); ok {
			if err = d.Set("key", vv); err != nil {
				return fmt.Errorf("Error reading key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading key: %v", err)
		}
	}

	if err = d.Set("value", flattenSystemAutomationActionFormDataValue2edl(o["value"], d, "value")); err != nil {
		if vv, ok := fortiAPIPatch(o["value"], "SystemAutomationActionFormData-Value"); ok {
			if err = d.Set("value", vv); err != nil {
				return fmt.Errorf("Error reading value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	return nil
}

func flattenSystemAutomationActionFormDataFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAutomationActionFormDataId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionFormDataKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionFormDataValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAutomationActionFormData(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemAutomationActionFormDataId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("key"); ok || d.HasChange("key") {
		t, err := expandSystemAutomationActionFormDataKey2edl(d, v, "key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok || d.HasChange("value") {
		t, err := expandSystemAutomationActionFormDataValue2edl(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
