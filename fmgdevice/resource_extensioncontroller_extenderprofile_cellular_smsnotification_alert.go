// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SMS alert list.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceExtensionControllerExtenderProfileCellularSmsNotificationAlert() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtensionControllerExtenderProfileCellularSmsNotificationAlertUpdate,
		Read:   resourceExtensionControllerExtenderProfileCellularSmsNotificationAlertRead,
		Update: resourceExtensionControllerExtenderProfileCellularSmsNotificationAlertUpdate,
		Delete: resourceExtensionControllerExtenderProfileCellularSmsNotificationAlertDelete,

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
			"data_exhausted": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fgt_backup_mode_switch": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"low_signal_strength": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mode_switch": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"os_image_fallback": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"session_disconnect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"system_reboot": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceExtensionControllerExtenderProfileCellularSmsNotificationAlertUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert(d)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerExtenderProfileCellularSmsNotificationAlert resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateExtensionControllerExtenderProfileCellularSmsNotificationAlert(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerExtenderProfileCellularSmsNotificationAlert resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ExtensionControllerExtenderProfileCellularSmsNotificationAlert")

	return resourceExtensionControllerExtenderProfileCellularSmsNotificationAlertRead(d, m)
}

func resourceExtensionControllerExtenderProfileCellularSmsNotificationAlertDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteExtensionControllerExtenderProfileCellularSmsNotificationAlert(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ExtensionControllerExtenderProfileCellularSmsNotificationAlert resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtensionControllerExtenderProfileCellularSmsNotificationAlertRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadExtensionControllerExtenderProfileCellularSmsNotificationAlert(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ExtensionControllerExtenderProfileCellularSmsNotificationAlert resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ExtensionControllerExtenderProfileCellularSmsNotificationAlert resource from API: %v", err)
	}
	return nil
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("data_exhausted", flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted4thl(o["data-exhausted"], d, "data_exhausted")); err != nil {
		if vv, ok := fortiAPIPatch(o["data-exhausted"], "ExtensionControllerExtenderProfileCellularSmsNotificationAlert-DataExhausted"); ok {
			if err = d.Set("data_exhausted", vv); err != nil {
				return fmt.Errorf("Error reading data_exhausted: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading data_exhausted: %v", err)
		}
	}

	if err = d.Set("fgt_backup_mode_switch", flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch4thl(o["fgt-backup-mode-switch"], d, "fgt_backup_mode_switch")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgt-backup-mode-switch"], "ExtensionControllerExtenderProfileCellularSmsNotificationAlert-FgtBackupModeSwitch"); ok {
			if err = d.Set("fgt_backup_mode_switch", vv); err != nil {
				return fmt.Errorf("Error reading fgt_backup_mode_switch: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgt_backup_mode_switch: %v", err)
		}
	}

	if err = d.Set("low_signal_strength", flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength4thl(o["low-signal-strength"], d, "low_signal_strength")); err != nil {
		if vv, ok := fortiAPIPatch(o["low-signal-strength"], "ExtensionControllerExtenderProfileCellularSmsNotificationAlert-LowSignalStrength"); ok {
			if err = d.Set("low_signal_strength", vv); err != nil {
				return fmt.Errorf("Error reading low_signal_strength: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading low_signal_strength: %v", err)
		}
	}

	if err = d.Set("mode_switch", flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch4thl(o["mode-switch"], d, "mode_switch")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode-switch"], "ExtensionControllerExtenderProfileCellularSmsNotificationAlert-ModeSwitch"); ok {
			if err = d.Set("mode_switch", vv); err != nil {
				return fmt.Errorf("Error reading mode_switch: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode_switch: %v", err)
		}
	}

	if err = d.Set("os_image_fallback", flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback4thl(o["os-image-fallback"], d, "os_image_fallback")); err != nil {
		if vv, ok := fortiAPIPatch(o["os-image-fallback"], "ExtensionControllerExtenderProfileCellularSmsNotificationAlert-OsImageFallback"); ok {
			if err = d.Set("os_image_fallback", vv); err != nil {
				return fmt.Errorf("Error reading os_image_fallback: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading os_image_fallback: %v", err)
		}
	}

	if err = d.Set("session_disconnect", flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect4thl(o["session-disconnect"], d, "session_disconnect")); err != nil {
		if vv, ok := fortiAPIPatch(o["session-disconnect"], "ExtensionControllerExtenderProfileCellularSmsNotificationAlert-SessionDisconnect"); ok {
			if err = d.Set("session_disconnect", vv); err != nil {
				return fmt.Errorf("Error reading session_disconnect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session_disconnect: %v", err)
		}
	}

	if err = d.Set("system_reboot", flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot4thl(o["system-reboot"], d, "system_reboot")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-reboot"], "ExtensionControllerExtenderProfileCellularSmsNotificationAlert-SystemReboot"); ok {
			if err = d.Set("system_reboot", vv); err != nil {
				return fmt.Errorf("Error reading system_reboot: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_reboot: %v", err)
		}
	}

	return nil
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("data_exhausted"); ok || d.HasChange("data_exhausted") {
		t, err := expandExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted4thl(d, v, "data_exhausted")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["data-exhausted"] = t
		}
	}

	if v, ok := d.GetOk("fgt_backup_mode_switch"); ok || d.HasChange("fgt_backup_mode_switch") {
		t, err := expandExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch4thl(d, v, "fgt_backup_mode_switch")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgt-backup-mode-switch"] = t
		}
	}

	if v, ok := d.GetOk("low_signal_strength"); ok || d.HasChange("low_signal_strength") {
		t, err := expandExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength4thl(d, v, "low_signal_strength")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["low-signal-strength"] = t
		}
	}

	if v, ok := d.GetOk("mode_switch"); ok || d.HasChange("mode_switch") {
		t, err := expandExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch4thl(d, v, "mode_switch")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode-switch"] = t
		}
	}

	if v, ok := d.GetOk("os_image_fallback"); ok || d.HasChange("os_image_fallback") {
		t, err := expandExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback4thl(d, v, "os_image_fallback")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os-image-fallback"] = t
		}
	}

	if v, ok := d.GetOk("session_disconnect"); ok || d.HasChange("session_disconnect") {
		t, err := expandExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect4thl(d, v, "session_disconnect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-disconnect"] = t
		}
	}

	if v, ok := d.GetOk("system_reboot"); ok || d.HasChange("system_reboot") {
		t, err := expandExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot4thl(d, v, "system_reboot")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-reboot"] = t
		}
	}

	return &obj, nil
}
