// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch SNMP system information globally.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerSnmpSysinfo() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerSnmpSysinfoUpdate,
		Read:   resourceSwitchControllerSnmpSysinfoRead,
		Update: resourceSwitchControllerSnmpSysinfoUpdate,
		Delete: resourceSwitchControllerSnmpSysinfoDelete,

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
			"location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerSnmpSysinfoUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerSnmpSysinfo(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSnmpSysinfo resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerSnmpSysinfo(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSnmpSysinfo resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SwitchControllerSnmpSysinfo")

	return resourceSwitchControllerSnmpSysinfoRead(d, m)
}

func resourceSwitchControllerSnmpSysinfoDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSwitchControllerSnmpSysinfo(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerSnmpSysinfo resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerSnmpSysinfoRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	o, err := c.ReadSwitchControllerSnmpSysinfo(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSnmpSysinfo resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerSnmpSysinfo(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSnmpSysinfo resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerSnmpSysinfoContactInfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSnmpSysinfoDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSnmpSysinfoEngineId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSnmpSysinfoLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSnmpSysinfoStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerSnmpSysinfo(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("contact_info", flattenSwitchControllerSnmpSysinfoContactInfo(o["contact-info"], d, "contact_info")); err != nil {
		if vv, ok := fortiAPIPatch(o["contact-info"], "SwitchControllerSnmpSysinfo-ContactInfo"); ok {
			if err = d.Set("contact_info", vv); err != nil {
				return fmt.Errorf("Error reading contact_info: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading contact_info: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchControllerSnmpSysinfoDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SwitchControllerSnmpSysinfo-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("engine_id", flattenSwitchControllerSnmpSysinfoEngineId(o["engine-id"], d, "engine_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["engine-id"], "SwitchControllerSnmpSysinfo-EngineId"); ok {
			if err = d.Set("engine_id", vv); err != nil {
				return fmt.Errorf("Error reading engine_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading engine_id: %v", err)
		}
	}

	if err = d.Set("location", flattenSwitchControllerSnmpSysinfoLocation(o["location"], d, "location")); err != nil {
		if vv, ok := fortiAPIPatch(o["location"], "SwitchControllerSnmpSysinfo-Location"); ok {
			if err = d.Set("location", vv); err != nil {
				return fmt.Errorf("Error reading location: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading location: %v", err)
		}
	}

	if err = d.Set("status", flattenSwitchControllerSnmpSysinfoStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SwitchControllerSnmpSysinfo-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerSnmpSysinfoFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerSnmpSysinfoContactInfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpSysinfoDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpSysinfoEngineId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpSysinfoLocation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpSysinfoStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerSnmpSysinfo(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("contact_info"); ok || d.HasChange("contact_info") {
		t, err := expandSwitchControllerSnmpSysinfoContactInfo(d, v, "contact_info")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["contact-info"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSwitchControllerSnmpSysinfoDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("engine_id"); ok || d.HasChange("engine_id") {
		t, err := expandSwitchControllerSnmpSysinfoEngineId(d, v, "engine_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["engine-id"] = t
		}
	}

	if v, ok := d.GetOk("location"); ok || d.HasChange("location") {
		t, err := expandSwitchControllerSnmpSysinfoLocation(d, v, "location")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["location"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSwitchControllerSnmpSysinfoStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
