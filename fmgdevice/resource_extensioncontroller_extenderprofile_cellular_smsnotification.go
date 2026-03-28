// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: FortiExtender cellular SMS notification configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceExtensionControllerExtenderProfileCellularSmsNotification() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtensionControllerExtenderProfileCellularSmsNotificationUpdate,
		Read:   resourceExtensionControllerExtenderProfileCellularSmsNotificationRead,
		Update: resourceExtensionControllerExtenderProfileCellularSmsNotificationUpdate,
		Delete: resourceExtensionControllerExtenderProfileCellularSmsNotificationDelete,

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
			"alert": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				},
			},
			"receiver": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alert": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"phone_number": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceExtensionControllerExtenderProfileCellularSmsNotificationUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectExtensionControllerExtenderProfileCellularSmsNotification(d)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerExtenderProfileCellularSmsNotification resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateExtensionControllerExtenderProfileCellularSmsNotification(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerExtenderProfileCellularSmsNotification resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ExtensionControllerExtenderProfileCellularSmsNotification")

	return resourceExtensionControllerExtenderProfileCellularSmsNotificationRead(d, m)
}

func resourceExtensionControllerExtenderProfileCellularSmsNotificationDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteExtensionControllerExtenderProfileCellularSmsNotification(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ExtensionControllerExtenderProfileCellularSmsNotification resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtensionControllerExtenderProfileCellularSmsNotificationRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadExtensionControllerExtenderProfileCellularSmsNotification(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ExtensionControllerExtenderProfileCellularSmsNotification resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectExtensionControllerExtenderProfileCellularSmsNotification(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ExtensionControllerExtenderProfileCellularSmsNotification resource from API: %v", err)
	}
	return nil
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlert3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "data_exhausted"
	if _, ok := i["data-exhausted"]; ok {
		result["data_exhausted"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted3rdl(i["data-exhausted"], d, pre_append)
	}

	pre_append = pre + ".0." + "fgt_backup_mode_switch"
	if _, ok := i["fgt-backup-mode-switch"]; ok {
		result["fgt_backup_mode_switch"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch3rdl(i["fgt-backup-mode-switch"], d, pre_append)
	}

	pre_append = pre + ".0." + "low_signal_strength"
	if _, ok := i["low-signal-strength"]; ok {
		result["low_signal_strength"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength3rdl(i["low-signal-strength"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode_switch"
	if _, ok := i["mode-switch"]; ok {
		result["mode_switch"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch3rdl(i["mode-switch"], d, pre_append)
	}

	pre_append = pre + ".0." + "os_image_fallback"
	if _, ok := i["os-image-fallback"]; ok {
		result["os_image_fallback"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback3rdl(i["os-image-fallback"], d, pre_append)
	}

	pre_append = pre + ".0." + "session_disconnect"
	if _, ok := i["session-disconnect"]; ok {
		result["session_disconnect"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect3rdl(i["session-disconnect"], d, pre_append)
	}

	pre_append = pre + ".0." + "system_reboot"
	if _, ok := i["system-reboot"]; ok {
		result["system_reboot"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot3rdl(i["system-reboot"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiver3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "alert"
		if _, ok := i["alert"]; ok {
			v := flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert3rdl(i["alert"], d, pre_append)
			tmp["alert"] = fortiAPISubPartPatch(v, "ExtensionControllerExtenderProfileCellularSmsNotification-Receiver-Alert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverName3rdl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ExtensionControllerExtenderProfileCellularSmsNotification-Receiver-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "phone_number"
		if _, ok := i["phone-number"]; ok {
			v := flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber3rdl(i["phone-number"], d, pre_append)
			tmp["phone_number"] = fortiAPISubPartPatch(v, "ExtensionControllerExtenderProfileCellularSmsNotification-Receiver-PhoneNumber")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus3rdl(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ExtensionControllerExtenderProfileCellularSmsNotification-Receiver-Status")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationStatus3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectExtensionControllerExtenderProfileCellularSmsNotification(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("alert", flattenExtensionControllerExtenderProfileCellularSmsNotificationAlert3rdl(o["alert"], d, "alert")); err != nil {
			if vv, ok := fortiAPIPatch(o["alert"], "ExtensionControllerExtenderProfileCellularSmsNotification-Alert"); ok {
				if err = d.Set("alert", vv); err != nil {
					return fmt.Errorf("Error reading alert: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading alert: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("alert"); ok {
			if err = d.Set("alert", flattenExtensionControllerExtenderProfileCellularSmsNotificationAlert3rdl(o["alert"], d, "alert")); err != nil {
				if vv, ok := fortiAPIPatch(o["alert"], "ExtensionControllerExtenderProfileCellularSmsNotification-Alert"); ok {
					if err = d.Set("alert", vv); err != nil {
						return fmt.Errorf("Error reading alert: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading alert: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("receiver", flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiver3rdl(o["receiver"], d, "receiver")); err != nil {
			if vv, ok := fortiAPIPatch(o["receiver"], "ExtensionControllerExtenderProfileCellularSmsNotification-Receiver"); ok {
				if err = d.Set("receiver", vv); err != nil {
					return fmt.Errorf("Error reading receiver: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading receiver: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("receiver"); ok {
			if err = d.Set("receiver", flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiver3rdl(o["receiver"], d, "receiver")); err != nil {
				if vv, ok := fortiAPIPatch(o["receiver"], "ExtensionControllerExtenderProfileCellularSmsNotification-Receiver"); ok {
					if err = d.Set("receiver", vv); err != nil {
						return fmt.Errorf("Error reading receiver: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading receiver: %v", err)
				}
			}
		}
	}

	if err = d.Set("status", flattenExtensionControllerExtenderProfileCellularSmsNotificationStatus3rdl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ExtensionControllerExtenderProfileCellularSmsNotification-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlert3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "data_exhausted"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["data-exhausted"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted3rdl(d, i["data_exhausted"], pre_append)
	}
	pre_append = pre + ".0." + "fgt_backup_mode_switch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fgt-backup-mode-switch"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch3rdl(d, i["fgt_backup_mode_switch"], pre_append)
	}
	pre_append = pre + ".0." + "low_signal_strength"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["low-signal-strength"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength3rdl(d, i["low_signal_strength"], pre_append)
	}
	pre_append = pre + ".0." + "mode_switch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mode-switch"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch3rdl(d, i["mode_switch"], pre_append)
	}
	pre_append = pre + ".0." + "os_image_fallback"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["os-image-fallback"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback3rdl(d, i["os_image_fallback"], pre_append)
	}
	pre_append = pre + ".0." + "session_disconnect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["session-disconnect"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect3rdl(d, i["session_disconnect"], pre_append)
	}
	pre_append = pre + ".0." + "system_reboot"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["system-reboot"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot3rdl(d, i["system_reboot"], pre_append)
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationReceiver3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "alert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["alert"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert3rdl(d, i["alert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverName3rdl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "phone_number"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["phone-number"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber3rdl(d, i["phone_number"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus3rdl(d, i["status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationStatus3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectExtensionControllerExtenderProfileCellularSmsNotification(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("alert"); ok || d.HasChange("alert") {
		t, err := expandExtensionControllerExtenderProfileCellularSmsNotificationAlert3rdl(d, v, "alert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alert"] = t
		}
	}

	if v, ok := d.GetOk("receiver"); ok || d.HasChange("receiver") {
		t, err := expandExtensionControllerExtenderProfileCellularSmsNotificationReceiver3rdl(d, v, "receiver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["receiver"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandExtensionControllerExtenderProfileCellularSmsNotificationStatus3rdl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
