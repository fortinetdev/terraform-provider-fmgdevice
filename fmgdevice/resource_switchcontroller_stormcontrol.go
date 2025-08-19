// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch storm control.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerStormControl() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerStormControlUpdate,
		Read:   resourceSwitchControllerStormControlRead,
		Update: resourceSwitchControllerStormControlUpdate,
		Delete: resourceSwitchControllerStormControlDelete,

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
			"rate": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceSwitchControllerStormControlUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerStormControl(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerStormControl resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerStormControl(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerStormControl resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SwitchControllerStormControl")

	return resourceSwitchControllerStormControlRead(d, m)
}

func resourceSwitchControllerStormControlDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerStormControl(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerStormControl resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerStormControlRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerStormControl(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerStormControl resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerStormControl(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerStormControl resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerStormControlBroadcast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerStormControlRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerStormControlUnknownMulticast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerStormControlUnknownUnicast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerStormControl(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("broadcast", flattenSwitchControllerStormControlBroadcast(o["broadcast"], d, "broadcast")); err != nil {
		if vv, ok := fortiAPIPatch(o["broadcast"], "SwitchControllerStormControl-Broadcast"); ok {
			if err = d.Set("broadcast", vv); err != nil {
				return fmt.Errorf("Error reading broadcast: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading broadcast: %v", err)
		}
	}

	if err = d.Set("rate", flattenSwitchControllerStormControlRate(o["rate"], d, "rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate"], "SwitchControllerStormControl-Rate"); ok {
			if err = d.Set("rate", vv); err != nil {
				return fmt.Errorf("Error reading rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate: %v", err)
		}
	}

	if err = d.Set("unknown_multicast", flattenSwitchControllerStormControlUnknownMulticast(o["unknown-multicast"], d, "unknown_multicast")); err != nil {
		if vv, ok := fortiAPIPatch(o["unknown-multicast"], "SwitchControllerStormControl-UnknownMulticast"); ok {
			if err = d.Set("unknown_multicast", vv); err != nil {
				return fmt.Errorf("Error reading unknown_multicast: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unknown_multicast: %v", err)
		}
	}

	if err = d.Set("unknown_unicast", flattenSwitchControllerStormControlUnknownUnicast(o["unknown-unicast"], d, "unknown_unicast")); err != nil {
		if vv, ok := fortiAPIPatch(o["unknown-unicast"], "SwitchControllerStormControl-UnknownUnicast"); ok {
			if err = d.Set("unknown_unicast", vv); err != nil {
				return fmt.Errorf("Error reading unknown_unicast: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unknown_unicast: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerStormControlFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerStormControlBroadcast(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStormControlRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStormControlUnknownMulticast(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStormControlUnknownUnicast(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerStormControl(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("broadcast"); ok || d.HasChange("broadcast") {
		t, err := expandSwitchControllerStormControlBroadcast(d, v, "broadcast")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["broadcast"] = t
		}
	}

	if v, ok := d.GetOk("rate"); ok || d.HasChange("rate") {
		t, err := expandSwitchControllerStormControlRate(d, v, "rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate"] = t
		}
	}

	if v, ok := d.GetOk("unknown_multicast"); ok || d.HasChange("unknown_multicast") {
		t, err := expandSwitchControllerStormControlUnknownMulticast(d, v, "unknown_multicast")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-multicast"] = t
		}
	}

	if v, ok := d.GetOk("unknown_unicast"); ok || d.HasChange("unknown_unicast") {
		t, err := expandSwitchControllerStormControlUnknownUnicast(d, v, "unknown_unicast")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-unicast"] = t
		}
	}

	return &obj, nil
}
