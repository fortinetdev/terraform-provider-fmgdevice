// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure network monitor settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerNetworkMonitorSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerNetworkMonitorSettingsUpdate,
		Read:   resourceSwitchControllerNetworkMonitorSettingsRead,
		Update: resourceSwitchControllerNetworkMonitorSettingsUpdate,
		Delete: resourceSwitchControllerNetworkMonitorSettingsDelete,

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
			"network_monitoring": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerNetworkMonitorSettingsUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerNetworkMonitorSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerNetworkMonitorSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerNetworkMonitorSettings(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerNetworkMonitorSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SwitchControllerNetworkMonitorSettings")

	return resourceSwitchControllerNetworkMonitorSettingsRead(d, m)
}

func resourceSwitchControllerNetworkMonitorSettingsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerNetworkMonitorSettings(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerNetworkMonitorSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerNetworkMonitorSettingsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerNetworkMonitorSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerNetworkMonitorSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerNetworkMonitorSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerNetworkMonitorSettings resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerNetworkMonitorSettingsNetworkMonitoring(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerNetworkMonitorSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("network_monitoring", flattenSwitchControllerNetworkMonitorSettingsNetworkMonitoring(o["network-monitoring"], d, "network_monitoring")); err != nil {
		if vv, ok := fortiAPIPatch(o["network-monitoring"], "SwitchControllerNetworkMonitorSettings-NetworkMonitoring"); ok {
			if err = d.Set("network_monitoring", vv); err != nil {
				return fmt.Errorf("Error reading network_monitoring: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading network_monitoring: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerNetworkMonitorSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerNetworkMonitorSettingsNetworkMonitoring(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerNetworkMonitorSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("network_monitoring"); ok || d.HasChange("network_monitoring") {
		t, err := expandSwitchControllerNetworkMonitorSettingsNetworkMonitoring(d, v, "network_monitoring")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-monitoring"] = t
		}
	}

	return &obj, nil
}
