// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configuration method to edit FortiSwitch storm control for measuring traffic activity using data rates to prevent traffic disruption.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerManagedSwitchStormControl() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerManagedSwitchStormControlUpdate,
		Read:   resourceSwitchControllerManagedSwitchStormControlRead,
		Update: resourceSwitchControllerManagedSwitchStormControlUpdate,
		Delete: resourceSwitchControllerManagedSwitchStormControlDelete,

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
			"managed_switch": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"broadcast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"local_override": &schema.Schema{
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

func resourceSwitchControllerManagedSwitchStormControlUpdate(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerManagedSwitchStormControl(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchStormControl resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerManagedSwitchStormControl(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchStormControl resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SwitchControllerManagedSwitchStormControl")

	return resourceSwitchControllerManagedSwitchStormControlRead(d, m)
}

func resourceSwitchControllerManagedSwitchStormControlDelete(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSwitchControllerManagedSwitchStormControl(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerManagedSwitchStormControl resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerManagedSwitchStormControlRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	managed_switch := d.Get("managed_switch").(string)
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
	if managed_switch == "" {
		managed_switch = importOptionChecking(m.(*FortiClient).Cfg, "managed_switch")
		if managed_switch == "" {
			return fmt.Errorf("Parameter managed_switch is missing")
		}
		if err = d.Set("managed_switch", managed_switch); err != nil {
			return fmt.Errorf("Error set params managed_switch: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	o, err := c.ReadSwitchControllerManagedSwitchStormControl(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchStormControl resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerManagedSwitchStormControl(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchStormControl resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerManagedSwitchStormControlBroadcast2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStormControlLocalOverride2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStormControlRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStormControlUnknownMulticast2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStormControlUnknownUnicast2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerManagedSwitchStormControl(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("broadcast", flattenSwitchControllerManagedSwitchStormControlBroadcast2edl(o["broadcast"], d, "broadcast")); err != nil {
		if vv, ok := fortiAPIPatch(o["broadcast"], "SwitchControllerManagedSwitchStormControl-Broadcast"); ok {
			if err = d.Set("broadcast", vv); err != nil {
				return fmt.Errorf("Error reading broadcast: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading broadcast: %v", err)
		}
	}

	if err = d.Set("local_override", flattenSwitchControllerManagedSwitchStormControlLocalOverride2edl(o["local-override"], d, "local_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-override"], "SwitchControllerManagedSwitchStormControl-LocalOverride"); ok {
			if err = d.Set("local_override", vv); err != nil {
				return fmt.Errorf("Error reading local_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_override: %v", err)
		}
	}

	if err = d.Set("rate", flattenSwitchControllerManagedSwitchStormControlRate2edl(o["rate"], d, "rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate"], "SwitchControllerManagedSwitchStormControl-Rate"); ok {
			if err = d.Set("rate", vv); err != nil {
				return fmt.Errorf("Error reading rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate: %v", err)
		}
	}

	if err = d.Set("unknown_multicast", flattenSwitchControllerManagedSwitchStormControlUnknownMulticast2edl(o["unknown-multicast"], d, "unknown_multicast")); err != nil {
		if vv, ok := fortiAPIPatch(o["unknown-multicast"], "SwitchControllerManagedSwitchStormControl-UnknownMulticast"); ok {
			if err = d.Set("unknown_multicast", vv); err != nil {
				return fmt.Errorf("Error reading unknown_multicast: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unknown_multicast: %v", err)
		}
	}

	if err = d.Set("unknown_unicast", flattenSwitchControllerManagedSwitchStormControlUnknownUnicast2edl(o["unknown-unicast"], d, "unknown_unicast")); err != nil {
		if vv, ok := fortiAPIPatch(o["unknown-unicast"], "SwitchControllerManagedSwitchStormControl-UnknownUnicast"); ok {
			if err = d.Set("unknown_unicast", vv); err != nil {
				return fmt.Errorf("Error reading unknown_unicast: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unknown_unicast: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerManagedSwitchStormControlFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerManagedSwitchStormControlBroadcast2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStormControlLocalOverride2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStormControlRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStormControlUnknownMulticast2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStormControlUnknownUnicast2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerManagedSwitchStormControl(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("broadcast"); ok || d.HasChange("broadcast") {
		t, err := expandSwitchControllerManagedSwitchStormControlBroadcast2edl(d, v, "broadcast")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["broadcast"] = t
		}
	}

	if v, ok := d.GetOk("local_override"); ok || d.HasChange("local_override") {
		t, err := expandSwitchControllerManagedSwitchStormControlLocalOverride2edl(d, v, "local_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-override"] = t
		}
	}

	if v, ok := d.GetOk("rate"); ok || d.HasChange("rate") {
		t, err := expandSwitchControllerManagedSwitchStormControlRate2edl(d, v, "rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate"] = t
		}
	}

	if v, ok := d.GetOk("unknown_multicast"); ok || d.HasChange("unknown_multicast") {
		t, err := expandSwitchControllerManagedSwitchStormControlUnknownMulticast2edl(d, v, "unknown_multicast")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-multicast"] = t
		}
	}

	if v, ok := d.GetOk("unknown_unicast"); ok || d.HasChange("unknown_unicast") {
		t, err := expandSwitchControllerManagedSwitchStormControlUnknownUnicast2edl(d, v, "unknown_unicast")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-unicast"] = t
		}
	}

	return &obj, nil
}
