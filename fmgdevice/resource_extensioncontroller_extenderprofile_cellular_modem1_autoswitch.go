// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: FortiExtender auto switch configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceExtensionControllerExtenderProfileCellularModem1AutoSwitch() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtensionControllerExtenderProfileCellularModem1AutoSwitchUpdate,
		Read:   resourceExtensionControllerExtenderProfileCellularModem1AutoSwitchRead,
		Update: resourceExtensionControllerExtenderProfileCellularModem1AutoSwitchUpdate,
		Delete: resourceExtensionControllerExtenderProfileCellularModem1AutoSwitchDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
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
			"extender_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dataplan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disconnect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disconnect_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"disconnect_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"signal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_back": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"switch_back_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_back_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceExtensionControllerExtenderProfileCellularModem1AutoSwitchUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	extender_profile := d.Get("extender_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["extender_profile"] = extender_profile

	obj, err := getObjectExtensionControllerExtenderProfileCellularModem1AutoSwitch(d)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerExtenderProfileCellularModem1AutoSwitch resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateExtensionControllerExtenderProfileCellularModem1AutoSwitch(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerExtenderProfileCellularModem1AutoSwitch resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ExtensionControllerExtenderProfileCellularModem1AutoSwitch")

	return resourceExtensionControllerExtenderProfileCellularModem1AutoSwitchRead(d, m)
}

func resourceExtensionControllerExtenderProfileCellularModem1AutoSwitchDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	extender_profile := d.Get("extender_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["extender_profile"] = extender_profile

	wsParams["adom"] = adomv

	err = c.DeleteExtensionControllerExtenderProfileCellularModem1AutoSwitch(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ExtensionControllerExtenderProfileCellularModem1AutoSwitch resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtensionControllerExtenderProfileCellularModem1AutoSwitchRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	extender_profile := d.Get("extender_profile").(string)
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
	if extender_profile == "" {
		extender_profile = importOptionChecking(m.(*FortiClient).Cfg, "extender_profile")
		if extender_profile == "" {
			return fmt.Errorf("Parameter extender_profile is missing")
		}
		if err = d.Set("extender_profile", extender_profile); err != nil {
			return fmt.Errorf("Error set params extender_profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["extender_profile"] = extender_profile

	o, err := c.ReadExtensionControllerExtenderProfileCellularModem1AutoSwitch(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ExtensionControllerExtenderProfileCellularModem1AutoSwitch resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectExtensionControllerExtenderProfileCellularModem1AutoSwitch(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ExtensionControllerExtenderProfileCellularModem1AutoSwitch resource from API: %v", err)
	}
	return nil
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDataplan4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnect4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSignal4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBack4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectExtensionControllerExtenderProfileCellularModem1AutoSwitch(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("dataplan", flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDataplan4thl(o["dataplan"], d, "dataplan")); err != nil {
		if vv, ok := fortiAPIPatch(o["dataplan"], "ExtensionControllerExtenderProfileCellularModem1AutoSwitch-Dataplan"); ok {
			if err = d.Set("dataplan", vv); err != nil {
				return fmt.Errorf("Error reading dataplan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dataplan: %v", err)
		}
	}

	if err = d.Set("disconnect", flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnect4thl(o["disconnect"], d, "disconnect")); err != nil {
		if vv, ok := fortiAPIPatch(o["disconnect"], "ExtensionControllerExtenderProfileCellularModem1AutoSwitch-Disconnect"); ok {
			if err = d.Set("disconnect", vv); err != nil {
				return fmt.Errorf("Error reading disconnect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading disconnect: %v", err)
		}
	}

	if err = d.Set("disconnect_period", flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod4thl(o["disconnect-period"], d, "disconnect_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["disconnect-period"], "ExtensionControllerExtenderProfileCellularModem1AutoSwitch-DisconnectPeriod"); ok {
			if err = d.Set("disconnect_period", vv); err != nil {
				return fmt.Errorf("Error reading disconnect_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading disconnect_period: %v", err)
		}
	}

	if err = d.Set("disconnect_threshold", flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold4thl(o["disconnect-threshold"], d, "disconnect_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["disconnect-threshold"], "ExtensionControllerExtenderProfileCellularModem1AutoSwitch-DisconnectThreshold"); ok {
			if err = d.Set("disconnect_threshold", vv); err != nil {
				return fmt.Errorf("Error reading disconnect_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading disconnect_threshold: %v", err)
		}
	}

	if err = d.Set("signal", flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSignal4thl(o["signal"], d, "signal")); err != nil {
		if vv, ok := fortiAPIPatch(o["signal"], "ExtensionControllerExtenderProfileCellularModem1AutoSwitch-Signal"); ok {
			if err = d.Set("signal", vv); err != nil {
				return fmt.Errorf("Error reading signal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading signal: %v", err)
		}
	}

	if err = d.Set("switch_back", flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBack4thl(o["switch-back"], d, "switch_back")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-back"], "ExtensionControllerExtenderProfileCellularModem1AutoSwitch-SwitchBack"); ok {
			if err = d.Set("switch_back", vv); err != nil {
				return fmt.Errorf("Error reading switch_back: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_back: %v", err)
		}
	}

	if err = d.Set("switch_back_time", flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime4thl(o["switch-back-time"], d, "switch_back_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-back-time"], "ExtensionControllerExtenderProfileCellularModem1AutoSwitch-SwitchBackTime"); ok {
			if err = d.Set("switch_back_time", vv); err != nil {
				return fmt.Errorf("Error reading switch_back_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_back_time: %v", err)
		}
	}

	if err = d.Set("switch_back_timer", flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer4thl(o["switch-back-timer"], d, "switch_back_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-back-timer"], "ExtensionControllerExtenderProfileCellularModem1AutoSwitch-SwitchBackTimer"); ok {
			if err = d.Set("switch_back_timer", vv); err != nil {
				return fmt.Errorf("Error reading switch_back_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_back_timer: %v", err)
		}
	}

	return nil
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDataplan4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnect4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSignal4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBack4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectExtensionControllerExtenderProfileCellularModem1AutoSwitch(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dataplan"); ok || d.HasChange("dataplan") {
		t, err := expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDataplan4thl(d, v, "dataplan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dataplan"] = t
		}
	}

	if v, ok := d.GetOk("disconnect"); ok || d.HasChange("disconnect") {
		t, err := expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnect4thl(d, v, "disconnect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disconnect"] = t
		}
	}

	if v, ok := d.GetOk("disconnect_period"); ok || d.HasChange("disconnect_period") {
		t, err := expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod4thl(d, v, "disconnect_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disconnect-period"] = t
		}
	}

	if v, ok := d.GetOk("disconnect_threshold"); ok || d.HasChange("disconnect_threshold") {
		t, err := expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold4thl(d, v, "disconnect_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disconnect-threshold"] = t
		}
	}

	if v, ok := d.GetOk("signal"); ok || d.HasChange("signal") {
		t, err := expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSignal4thl(d, v, "signal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signal"] = t
		}
	}

	if v, ok := d.GetOk("switch_back"); ok || d.HasChange("switch_back") {
		t, err := expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBack4thl(d, v, "switch_back")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-back"] = t
		}
	}

	if v, ok := d.GetOk("switch_back_time"); ok || d.HasChange("switch_back_time") {
		t, err := expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime4thl(d, v, "switch_back_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-back-time"] = t
		}
	}

	if v, ok := d.GetOk("switch_back_timer"); ok || d.HasChange("switch_back_timer") {
		t, err := expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer4thl(d, v, "switch_back_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-back-timer"] = t
		}
	}

	return &obj, nil
}
