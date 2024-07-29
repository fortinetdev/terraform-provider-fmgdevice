// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch LLDP settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerLldpSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerLldpSettingsUpdate,
		Read:   resourceSwitchControllerLldpSettingsRead,
		Update: resourceSwitchControllerLldpSettingsUpdate,
		Delete: resourceSwitchControllerLldpSettingsDelete,

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
			"device_detection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fast_start_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"management_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tx_hold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tx_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerLldpSettingsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	obj, err := getObjectSwitchControllerLldpSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerLldpSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerLldpSettings(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerLldpSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SwitchControllerLldpSettings")

	return resourceSwitchControllerLldpSettingsRead(d, m)
}

func resourceSwitchControllerLldpSettingsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	err = c.DeleteSwitchControllerLldpSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerLldpSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerLldpSettingsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerLldpSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerLldpSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerLldpSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerLldpSettings resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerLldpSettingsDeviceDetection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpSettingsFastStartInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpSettingsManagementInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpSettingsTxHold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpSettingsTxInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerLldpSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("device_detection", flattenSwitchControllerLldpSettingsDeviceDetection(o["device-detection"], d, "device_detection")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-detection"], "SwitchControllerLldpSettings-DeviceDetection"); ok {
			if err = d.Set("device_detection", vv); err != nil {
				return fmt.Errorf("Error reading device_detection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_detection: %v", err)
		}
	}

	if err = d.Set("fast_start_interval", flattenSwitchControllerLldpSettingsFastStartInterval(o["fast-start-interval"], d, "fast_start_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["fast-start-interval"], "SwitchControllerLldpSettings-FastStartInterval"); ok {
			if err = d.Set("fast_start_interval", vv); err != nil {
				return fmt.Errorf("Error reading fast_start_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fast_start_interval: %v", err)
		}
	}

	if err = d.Set("management_interface", flattenSwitchControllerLldpSettingsManagementInterface(o["management-interface"], d, "management_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["management-interface"], "SwitchControllerLldpSettings-ManagementInterface"); ok {
			if err = d.Set("management_interface", vv); err != nil {
				return fmt.Errorf("Error reading management_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading management_interface: %v", err)
		}
	}

	if err = d.Set("tx_hold", flattenSwitchControllerLldpSettingsTxHold(o["tx-hold"], d, "tx_hold")); err != nil {
		if vv, ok := fortiAPIPatch(o["tx-hold"], "SwitchControllerLldpSettings-TxHold"); ok {
			if err = d.Set("tx_hold", vv); err != nil {
				return fmt.Errorf("Error reading tx_hold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tx_hold: %v", err)
		}
	}

	if err = d.Set("tx_interval", flattenSwitchControllerLldpSettingsTxInterval(o["tx-interval"], d, "tx_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["tx-interval"], "SwitchControllerLldpSettings-TxInterval"); ok {
			if err = d.Set("tx_interval", vv); err != nil {
				return fmt.Errorf("Error reading tx_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tx_interval: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerLldpSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerLldpSettingsDeviceDetection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpSettingsFastStartInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpSettingsManagementInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpSettingsTxHold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpSettingsTxInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerLldpSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("device_detection"); ok || d.HasChange("device_detection") {
		t, err := expandSwitchControllerLldpSettingsDeviceDetection(d, v, "device_detection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-detection"] = t
		}
	}

	if v, ok := d.GetOk("fast_start_interval"); ok || d.HasChange("fast_start_interval") {
		t, err := expandSwitchControllerLldpSettingsFastStartInterval(d, v, "fast_start_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fast-start-interval"] = t
		}
	}

	if v, ok := d.GetOk("management_interface"); ok || d.HasChange("management_interface") {
		t, err := expandSwitchControllerLldpSettingsManagementInterface(d, v, "management_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["management-interface"] = t
		}
	}

	if v, ok := d.GetOk("tx_hold"); ok || d.HasChange("tx_hold") {
		t, err := expandSwitchControllerLldpSettingsTxHold(d, v, "tx_hold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tx-hold"] = t
		}
	}

	if v, ok := d.GetOk("tx_interval"); ok || d.HasChange("tx_interval") {
		t, err := expandSwitchControllerLldpSettingsTxInterval(d, v, "tx_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tx-interval"] = t
		}
	}

	return &obj, nil
}
