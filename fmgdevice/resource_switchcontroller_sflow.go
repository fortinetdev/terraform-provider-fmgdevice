// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch sFlow.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerSflow() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerSflowUpdate,
		Read:   resourceSwitchControllerSflowRead,
		Update: resourceSwitchControllerSflowUpdate,
		Delete: resourceSwitchControllerSflowDelete,

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
			"collector_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"collector_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerSflowUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerSflow(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSflow resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerSflow(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSflow resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SwitchControllerSflow")

	return resourceSwitchControllerSflowRead(d, m)
}

func resourceSwitchControllerSflowDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerSflow(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerSflow resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerSflowRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerSflow(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSflow resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerSflow(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSflow resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerSflowCollectorIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSflowCollectorPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerSflow(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("collector_ip", flattenSwitchControllerSflowCollectorIp(o["collector-ip"], d, "collector_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["collector-ip"], "SwitchControllerSflow-CollectorIp"); ok {
			if err = d.Set("collector_ip", vv); err != nil {
				return fmt.Errorf("Error reading collector_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading collector_ip: %v", err)
		}
	}

	if err = d.Set("collector_port", flattenSwitchControllerSflowCollectorPort(o["collector-port"], d, "collector_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["collector-port"], "SwitchControllerSflow-CollectorPort"); ok {
			if err = d.Set("collector_port", vv); err != nil {
				return fmt.Errorf("Error reading collector_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading collector_port: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerSflowFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerSflowCollectorIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSflowCollectorPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerSflow(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("collector_ip"); ok || d.HasChange("collector_ip") {
		t, err := expandSwitchControllerSflowCollectorIp(d, v, "collector_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["collector-ip"] = t
		}
	}

	if v, ok := d.GetOk("collector_port"); ok || d.HasChange("collector_port") {
		t, err := expandSwitchControllerSflowCollectorPort(d, v, "collector_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["collector-port"] = t
		}
	}

	return &obj, nil
}
