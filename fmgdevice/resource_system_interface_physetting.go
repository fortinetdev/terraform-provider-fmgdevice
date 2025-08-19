// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: PHY settings

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemInterfacePhySetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemInterfacePhySettingUpdate,
		Read:   resourceSystemInterfacePhySettingRead,
		Update: resourceSystemInterfacePhySettingUpdate,
		Delete: resourceSystemInterfacePhySettingDelete,

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
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"signal_ok_threshold_value": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSystemInterfacePhySettingUpdate(d *schema.ResourceData, m interface{}) error {
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
	var_interface := d.Get("interface").(string)
	paradict["device"] = device_name
	paradict["interface"] = var_interface

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemInterfacePhySetting(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterfacePhySetting resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemInterfacePhySetting(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterfacePhySetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemInterfacePhySetting")

	return resourceSystemInterfacePhySettingRead(d, m)
}

func resourceSystemInterfacePhySettingDelete(d *schema.ResourceData, m interface{}) error {
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
	var_interface := d.Get("interface").(string)
	paradict["device"] = device_name
	paradict["interface"] = var_interface

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSystemInterfacePhySetting(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemInterfacePhySetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemInterfacePhySettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	var_interface := d.Get("interface").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if var_interface == "" {
		var_interface = importOptionChecking(m.(*FortiClient).Cfg, "interface")
		if var_interface == "" {
			return fmt.Errorf("Parameter interface is missing")
		}
		if err = d.Set("interface", var_interface); err != nil {
			return fmt.Errorf("Error set params interface: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["interface"] = var_interface

	o, err := c.ReadSystemInterfacePhySetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterfacePhySetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemInterfacePhySetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterfacePhySetting resource from API: %v", err)
	}
	return nil
}

func flattenSystemInterfacePhySettingSignalOkThresholdValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemInterfacePhySetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("signal_ok_threshold_value", flattenSystemInterfacePhySettingSignalOkThresholdValue2edl(o["signal-ok-threshold-value"], d, "signal_ok_threshold_value")); err != nil {
		if vv, ok := fortiAPIPatch(o["signal-ok-threshold-value"], "SystemInterfacePhySetting-SignalOkThresholdValue"); ok {
			if err = d.Set("signal_ok_threshold_value", vv); err != nil {
				return fmt.Errorf("Error reading signal_ok_threshold_value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading signal_ok_threshold_value: %v", err)
		}
	}

	return nil
}

func flattenSystemInterfacePhySettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemInterfacePhySettingSignalOkThresholdValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemInterfacePhySetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("signal_ok_threshold_value"); ok || d.HasChange("signal_ok_threshold_value") {
		t, err := expandSystemInterfacePhySettingSignalOkThresholdValue2edl(d, v, "signal_ok_threshold_value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signal-ok-threshold-value"] = t
		}
	}

	return &obj, nil
}
