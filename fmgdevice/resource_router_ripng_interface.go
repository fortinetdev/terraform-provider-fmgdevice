// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: RIPng interface configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterRipngInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterRipngInterfaceCreate,
		Read:   resourceRouterRipngInterfaceRead,
		Update: resourceRouterRipngInterfaceUpdate,
		Delete: resourceRouterRipngInterfaceDelete,

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
			"flags": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"split_horizon": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"split_horizon_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRouterRipngInterfaceCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterRipngInterface(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterRipngInterface resource while getting object: %v", err)
	}

	_, err = c.CreateRouterRipngInterface(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating RouterRipngInterface resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceRouterRipngInterfaceRead(d, m)
}

func resourceRouterRipngInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterRipngInterface(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterRipngInterface resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterRipngInterface(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterRipngInterface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceRouterRipngInterfaceRead(d, m)
}

func resourceRouterRipngInterfaceDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterRipngInterface(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterRipngInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterRipngInterfaceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterRipngInterface(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterRipngInterface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterRipngInterface(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterRipngInterface resource from API: %v", err)
	}
	return nil
}

func flattenRouterRipngInterfaceFlags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngInterfaceName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngInterfaceSplitHorizon2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngInterfaceSplitHorizonStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterRipngInterface(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("flags", flattenRouterRipngInterfaceFlags2edl(o["flags"], d, "flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["flags"], "RouterRipngInterface-Flags"); ok {
			if err = d.Set("flags", vv); err != nil {
				return fmt.Errorf("Error reading flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flags: %v", err)
		}
	}

	if err = d.Set("name", flattenRouterRipngInterfaceName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "RouterRipngInterface-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("split_horizon", flattenRouterRipngInterfaceSplitHorizon2edl(o["split-horizon"], d, "split_horizon")); err != nil {
		if vv, ok := fortiAPIPatch(o["split-horizon"], "RouterRipngInterface-SplitHorizon"); ok {
			if err = d.Set("split_horizon", vv); err != nil {
				return fmt.Errorf("Error reading split_horizon: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading split_horizon: %v", err)
		}
	}

	if err = d.Set("split_horizon_status", flattenRouterRipngInterfaceSplitHorizonStatus2edl(o["split-horizon-status"], d, "split_horizon_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["split-horizon-status"], "RouterRipngInterface-SplitHorizonStatus"); ok {
			if err = d.Set("split_horizon_status", vv); err != nil {
				return fmt.Errorf("Error reading split_horizon_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading split_horizon_status: %v", err)
		}
	}

	return nil
}

func flattenRouterRipngInterfaceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterRipngInterfaceFlags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngInterfaceName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngInterfaceSplitHorizon2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngInterfaceSplitHorizonStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterRipngInterface(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("flags"); ok || d.HasChange("flags") {
		t, err := expandRouterRipngInterfaceFlags2edl(d, v, "flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flags"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandRouterRipngInterfaceName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("split_horizon"); ok || d.HasChange("split_horizon") {
		t, err := expandRouterRipngInterfaceSplitHorizon2edl(d, v, "split_horizon")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-horizon"] = t
		}
	}

	if v, ok := d.GetOk("split_horizon_status"); ok || d.HasChange("split_horizon_status") {
		t, err := expandRouterRipngInterfaceSplitHorizonStatus2edl(d, v, "split_horizon_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-horizon-status"] = t
		}
	}

	return &obj, nil
}
