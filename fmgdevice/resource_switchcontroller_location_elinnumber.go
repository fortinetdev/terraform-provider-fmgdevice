// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure location ELIN number.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerLocationElinNumber() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerLocationElinNumberUpdate,
		Read:   resourceSwitchControllerLocationElinNumberRead,
		Update: resourceSwitchControllerLocationElinNumberUpdate,
		Delete: resourceSwitchControllerLocationElinNumberDelete,

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
			"location": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"elin_num": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSwitchControllerLocationElinNumberUpdate(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	location := d.Get("location").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["location"] = location

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerLocationElinNumber(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerLocationElinNumber resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerLocationElinNumber(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerLocationElinNumber resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SwitchControllerLocationElinNumber")

	return resourceSwitchControllerLocationElinNumberRead(d, m)
}

func resourceSwitchControllerLocationElinNumberDelete(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	location := d.Get("location").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["location"] = location

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSwitchControllerLocationElinNumber(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerLocationElinNumber resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerLocationElinNumberRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	location := d.Get("location").(string)
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
	if location == "" {
		location = importOptionChecking(m.(*FortiClient).Cfg, "location")
		if location == "" {
			return fmt.Errorf("Parameter location is missing")
		}
		if err = d.Set("location", location); err != nil {
			return fmt.Errorf("Error set params location: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["location"] = location

	o, err := c.ReadSwitchControllerLocationElinNumber(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerLocationElinNumber resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerLocationElinNumber(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerLocationElinNumber resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerLocationElinNumberElinNum2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationElinNumberParentKey2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerLocationElinNumber(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("elin_num", flattenSwitchControllerLocationElinNumberElinNum2edl(o["elin-num"], d, "elin_num")); err != nil {
		if vv, ok := fortiAPIPatch(o["elin-num"], "SwitchControllerLocationElinNumber-ElinNum"); ok {
			if err = d.Set("elin_num", vv); err != nil {
				return fmt.Errorf("Error reading elin_num: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading elin_num: %v", err)
		}
	}

	if err = d.Set("parent_key", flattenSwitchControllerLocationElinNumberParentKey2edl(o["parent-key"], d, "parent_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["parent-key"], "SwitchControllerLocationElinNumber-ParentKey"); ok {
			if err = d.Set("parent_key", vv); err != nil {
				return fmt.Errorf("Error reading parent_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading parent_key: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerLocationElinNumberFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerLocationElinNumberElinNum2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationElinNumberParentKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerLocationElinNumber(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("elin_num"); ok || d.HasChange("elin_num") {
		t, err := expandSwitchControllerLocationElinNumberElinNum2edl(d, v, "elin_num")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["elin-num"] = t
		}
	}

	if v, ok := d.GetOk("parent_key"); ok || d.HasChange("parent_key") {
		t, err := expandSwitchControllerLocationElinNumberParentKey2edl(d, v, "parent_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parent-key"] = t
		}
	}

	return &obj, nil
}
