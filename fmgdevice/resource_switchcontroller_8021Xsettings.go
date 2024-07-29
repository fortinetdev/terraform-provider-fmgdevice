// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure global 802.1X settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchController8021XSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchController8021XSettingsUpdate,
		Read:   resourceSwitchController8021XSettingsRead,
		Update: resourceSwitchController8021XSettingsUpdate,
		Delete: resourceSwitchController8021XSettingsDelete,

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
			"link_down_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mab_reauth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_called_station_delimiter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_calling_station_delimiter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_case": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_password_delimiter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_username_delimiter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_reauth_attempt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"reauth_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tx_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchController8021XSettingsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchController8021XSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchController8021XSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchController8021XSettings(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SwitchController8021XSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SwitchController8021XSettings")

	return resourceSwitchController8021XSettingsRead(d, m)
}

func resourceSwitchController8021XSettingsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchController8021XSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchController8021XSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchController8021XSettingsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchController8021XSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchController8021XSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchController8021XSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchController8021XSettings resource from API: %v", err)
	}
	return nil
}

func flattenSwitchController8021XSettingsLinkDownAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchController8021XSettingsMabReauth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchController8021XSettingsMacCalledStationDelimiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchController8021XSettingsMacCallingStationDelimiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchController8021XSettingsMacCase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchController8021XSettingsMacPasswordDelimiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchController8021XSettingsMacUsernameDelimiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchController8021XSettingsMaxReauthAttempt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchController8021XSettingsReauthPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchController8021XSettingsTxPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchController8021XSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("link_down_auth", flattenSwitchController8021XSettingsLinkDownAuth(o["link-down-auth"], d, "link_down_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["link-down-auth"], "SwitchController8021XSettings-LinkDownAuth"); ok {
			if err = d.Set("link_down_auth", vv); err != nil {
				return fmt.Errorf("Error reading link_down_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading link_down_auth: %v", err)
		}
	}

	if err = d.Set("mab_reauth", flattenSwitchController8021XSettingsMabReauth(o["mab-reauth"], d, "mab_reauth")); err != nil {
		if vv, ok := fortiAPIPatch(o["mab-reauth"], "SwitchController8021XSettings-MabReauth"); ok {
			if err = d.Set("mab_reauth", vv); err != nil {
				return fmt.Errorf("Error reading mab_reauth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mab_reauth: %v", err)
		}
	}

	if err = d.Set("mac_called_station_delimiter", flattenSwitchController8021XSettingsMacCalledStationDelimiter(o["mac-called-station-delimiter"], d, "mac_called_station_delimiter")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-called-station-delimiter"], "SwitchController8021XSettings-MacCalledStationDelimiter"); ok {
			if err = d.Set("mac_called_station_delimiter", vv); err != nil {
				return fmt.Errorf("Error reading mac_called_station_delimiter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_called_station_delimiter: %v", err)
		}
	}

	if err = d.Set("mac_calling_station_delimiter", flattenSwitchController8021XSettingsMacCallingStationDelimiter(o["mac-calling-station-delimiter"], d, "mac_calling_station_delimiter")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-calling-station-delimiter"], "SwitchController8021XSettings-MacCallingStationDelimiter"); ok {
			if err = d.Set("mac_calling_station_delimiter", vv); err != nil {
				return fmt.Errorf("Error reading mac_calling_station_delimiter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_calling_station_delimiter: %v", err)
		}
	}

	if err = d.Set("mac_case", flattenSwitchController8021XSettingsMacCase(o["mac-case"], d, "mac_case")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-case"], "SwitchController8021XSettings-MacCase"); ok {
			if err = d.Set("mac_case", vv); err != nil {
				return fmt.Errorf("Error reading mac_case: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_case: %v", err)
		}
	}

	if err = d.Set("mac_password_delimiter", flattenSwitchController8021XSettingsMacPasswordDelimiter(o["mac-password-delimiter"], d, "mac_password_delimiter")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-password-delimiter"], "SwitchController8021XSettings-MacPasswordDelimiter"); ok {
			if err = d.Set("mac_password_delimiter", vv); err != nil {
				return fmt.Errorf("Error reading mac_password_delimiter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_password_delimiter: %v", err)
		}
	}

	if err = d.Set("mac_username_delimiter", flattenSwitchController8021XSettingsMacUsernameDelimiter(o["mac-username-delimiter"], d, "mac_username_delimiter")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-username-delimiter"], "SwitchController8021XSettings-MacUsernameDelimiter"); ok {
			if err = d.Set("mac_username_delimiter", vv); err != nil {
				return fmt.Errorf("Error reading mac_username_delimiter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_username_delimiter: %v", err)
		}
	}

	if err = d.Set("max_reauth_attempt", flattenSwitchController8021XSettingsMaxReauthAttempt(o["max-reauth-attempt"], d, "max_reauth_attempt")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-reauth-attempt"], "SwitchController8021XSettings-MaxReauthAttempt"); ok {
			if err = d.Set("max_reauth_attempt", vv); err != nil {
				return fmt.Errorf("Error reading max_reauth_attempt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_reauth_attempt: %v", err)
		}
	}

	if err = d.Set("reauth_period", flattenSwitchController8021XSettingsReauthPeriod(o["reauth-period"], d, "reauth_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["reauth-period"], "SwitchController8021XSettings-ReauthPeriod"); ok {
			if err = d.Set("reauth_period", vv); err != nil {
				return fmt.Errorf("Error reading reauth_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reauth_period: %v", err)
		}
	}

	if err = d.Set("tx_period", flattenSwitchController8021XSettingsTxPeriod(o["tx-period"], d, "tx_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["tx-period"], "SwitchController8021XSettings-TxPeriod"); ok {
			if err = d.Set("tx_period", vv); err != nil {
				return fmt.Errorf("Error reading tx_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tx_period: %v", err)
		}
	}

	return nil
}

func flattenSwitchController8021XSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchController8021XSettingsLinkDownAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchController8021XSettingsMabReauth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchController8021XSettingsMacCalledStationDelimiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchController8021XSettingsMacCallingStationDelimiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchController8021XSettingsMacCase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchController8021XSettingsMacPasswordDelimiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchController8021XSettingsMacUsernameDelimiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchController8021XSettingsMaxReauthAttempt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchController8021XSettingsReauthPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchController8021XSettingsTxPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchController8021XSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("link_down_auth"); ok || d.HasChange("link_down_auth") {
		t, err := expandSwitchController8021XSettingsLinkDownAuth(d, v, "link_down_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-down-auth"] = t
		}
	}

	if v, ok := d.GetOk("mab_reauth"); ok || d.HasChange("mab_reauth") {
		t, err := expandSwitchController8021XSettingsMabReauth(d, v, "mab_reauth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mab-reauth"] = t
		}
	}

	if v, ok := d.GetOk("mac_called_station_delimiter"); ok || d.HasChange("mac_called_station_delimiter") {
		t, err := expandSwitchController8021XSettingsMacCalledStationDelimiter(d, v, "mac_called_station_delimiter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-called-station-delimiter"] = t
		}
	}

	if v, ok := d.GetOk("mac_calling_station_delimiter"); ok || d.HasChange("mac_calling_station_delimiter") {
		t, err := expandSwitchController8021XSettingsMacCallingStationDelimiter(d, v, "mac_calling_station_delimiter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-calling-station-delimiter"] = t
		}
	}

	if v, ok := d.GetOk("mac_case"); ok || d.HasChange("mac_case") {
		t, err := expandSwitchController8021XSettingsMacCase(d, v, "mac_case")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-case"] = t
		}
	}

	if v, ok := d.GetOk("mac_password_delimiter"); ok || d.HasChange("mac_password_delimiter") {
		t, err := expandSwitchController8021XSettingsMacPasswordDelimiter(d, v, "mac_password_delimiter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-password-delimiter"] = t
		}
	}

	if v, ok := d.GetOk("mac_username_delimiter"); ok || d.HasChange("mac_username_delimiter") {
		t, err := expandSwitchController8021XSettingsMacUsernameDelimiter(d, v, "mac_username_delimiter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-username-delimiter"] = t
		}
	}

	if v, ok := d.GetOk("max_reauth_attempt"); ok || d.HasChange("max_reauth_attempt") {
		t, err := expandSwitchController8021XSettingsMaxReauthAttempt(d, v, "max_reauth_attempt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-reauth-attempt"] = t
		}
	}

	if v, ok := d.GetOk("reauth_period"); ok || d.HasChange("reauth_period") {
		t, err := expandSwitchController8021XSettingsReauthPeriod(d, v, "reauth_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reauth-period"] = t
		}
	}

	if v, ok := d.GetOk("tx_period"); ok || d.HasChange("tx_period") {
		t, err := expandSwitchController8021XSettingsTxPeriod(d, v, "tx_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tx-period"] = t
		}
	}

	return &obj, nil
}
