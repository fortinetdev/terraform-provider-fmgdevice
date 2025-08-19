// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch storm control policy to be applied on managed-switch ports.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerStormControlPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerStormControlPolicyCreate,
		Read:   resourceSwitchControllerStormControlPolicyRead,
		Update: resourceSwitchControllerStormControlPolicyUpdate,
		Delete: resourceSwitchControllerStormControlPolicyDelete,

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
			"broadcast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"storm_control_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unknown_multicast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unknown_unicast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerStormControlPolicyCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerStormControlPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerStormControlPolicy resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerStormControlPolicy(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerStormControlPolicy resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerStormControlPolicyRead(d, m)
}

func resourceSwitchControllerStormControlPolicyUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerStormControlPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerStormControlPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerStormControlPolicy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerStormControlPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerStormControlPolicyRead(d, m)
}

func resourceSwitchControllerStormControlPolicyDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerStormControlPolicy(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerStormControlPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerStormControlPolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerStormControlPolicy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerStormControlPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerStormControlPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerStormControlPolicy resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerStormControlPolicyBroadcast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerStormControlPolicyDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerStormControlPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerStormControlPolicyRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerStormControlPolicyStormControlMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerStormControlPolicyUnknownMulticast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerStormControlPolicyUnknownUnicast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerStormControlPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("broadcast", flattenSwitchControllerStormControlPolicyBroadcast(o["broadcast"], d, "broadcast")); err != nil {
		if vv, ok := fortiAPIPatch(o["broadcast"], "SwitchControllerStormControlPolicy-Broadcast"); ok {
			if err = d.Set("broadcast", vv); err != nil {
				return fmt.Errorf("Error reading broadcast: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading broadcast: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchControllerStormControlPolicyDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SwitchControllerStormControlPolicy-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchControllerStormControlPolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerStormControlPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("rate", flattenSwitchControllerStormControlPolicyRate(o["rate"], d, "rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate"], "SwitchControllerStormControlPolicy-Rate"); ok {
			if err = d.Set("rate", vv); err != nil {
				return fmt.Errorf("Error reading rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate: %v", err)
		}
	}

	if err = d.Set("storm_control_mode", flattenSwitchControllerStormControlPolicyStormControlMode(o["storm-control-mode"], d, "storm_control_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["storm-control-mode"], "SwitchControllerStormControlPolicy-StormControlMode"); ok {
			if err = d.Set("storm_control_mode", vv); err != nil {
				return fmt.Errorf("Error reading storm_control_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading storm_control_mode: %v", err)
		}
	}

	if err = d.Set("unknown_multicast", flattenSwitchControllerStormControlPolicyUnknownMulticast(o["unknown-multicast"], d, "unknown_multicast")); err != nil {
		if vv, ok := fortiAPIPatch(o["unknown-multicast"], "SwitchControllerStormControlPolicy-UnknownMulticast"); ok {
			if err = d.Set("unknown_multicast", vv); err != nil {
				return fmt.Errorf("Error reading unknown_multicast: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unknown_multicast: %v", err)
		}
	}

	if err = d.Set("unknown_unicast", flattenSwitchControllerStormControlPolicyUnknownUnicast(o["unknown-unicast"], d, "unknown_unicast")); err != nil {
		if vv, ok := fortiAPIPatch(o["unknown-unicast"], "SwitchControllerStormControlPolicy-UnknownUnicast"); ok {
			if err = d.Set("unknown_unicast", vv); err != nil {
				return fmt.Errorf("Error reading unknown_unicast: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unknown_unicast: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerStormControlPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerStormControlPolicyBroadcast(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStormControlPolicyDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStormControlPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStormControlPolicyRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStormControlPolicyStormControlMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStormControlPolicyUnknownMulticast(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStormControlPolicyUnknownUnicast(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerStormControlPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("broadcast"); ok || d.HasChange("broadcast") {
		t, err := expandSwitchControllerStormControlPolicyBroadcast(d, v, "broadcast")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["broadcast"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSwitchControllerStormControlPolicyDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerStormControlPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("rate"); ok || d.HasChange("rate") {
		t, err := expandSwitchControllerStormControlPolicyRate(d, v, "rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate"] = t
		}
	}

	if v, ok := d.GetOk("storm_control_mode"); ok || d.HasChange("storm_control_mode") {
		t, err := expandSwitchControllerStormControlPolicyStormControlMode(d, v, "storm_control_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["storm-control-mode"] = t
		}
	}

	if v, ok := d.GetOk("unknown_multicast"); ok || d.HasChange("unknown_multicast") {
		t, err := expandSwitchControllerStormControlPolicyUnknownMulticast(d, v, "unknown_multicast")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-multicast"] = t
		}
	}

	if v, ok := d.GetOk("unknown_unicast"); ok || d.HasChange("unknown_unicast") {
		t, err := expandSwitchControllerStormControlPolicyUnknownUnicast(d, v, "unknown_unicast")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-unicast"] = t
		}
	}

	return &obj, nil
}
