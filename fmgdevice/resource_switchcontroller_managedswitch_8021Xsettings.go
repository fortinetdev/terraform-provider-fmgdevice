// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configuration method to edit FortiSwitch 802.1X global settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerManagedSwitch8021XSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerManagedSwitch8021XSettingsUpdate,
		Read:   resourceSwitchControllerManagedSwitch8021XSettingsRead,
		Update: resourceSwitchControllerManagedSwitch8021XSettingsUpdate,
		Delete: resourceSwitchControllerManagedSwitch8021XSettingsDelete,

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
			"link_down_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_override": &schema.Schema{
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

func resourceSwitchControllerManagedSwitch8021XSettingsUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerManagedSwitch8021XSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitch8021XSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerManagedSwitch8021XSettings(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitch8021XSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SwitchControllerManagedSwitch8021XSettings")

	return resourceSwitchControllerManagedSwitch8021XSettingsRead(d, m)
}

func resourceSwitchControllerManagedSwitch8021XSettingsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerManagedSwitch8021XSettings(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerManagedSwitch8021XSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerManagedSwitch8021XSettingsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerManagedSwitch8021XSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitch8021XSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerManagedSwitch8021XSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitch8021XSettings resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerManagedSwitch8021XSettingsLinkDownAuth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettingsLocalOverride2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettingsMabReauth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettingsMacCalledStationDelimiter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettingsMacCallingStationDelimiter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettingsMacCase2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettingsMacPasswordDelimiter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettingsMacUsernameDelimiter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettingsMaxReauthAttempt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettingsReauthPeriod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettingsTxPeriod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerManagedSwitch8021XSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("link_down_auth", flattenSwitchControllerManagedSwitch8021XSettingsLinkDownAuth2edl(o["link-down-auth"], d, "link_down_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["link-down-auth"], "SwitchControllerManagedSwitch8021XSettings-LinkDownAuth"); ok {
			if err = d.Set("link_down_auth", vv); err != nil {
				return fmt.Errorf("Error reading link_down_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading link_down_auth: %v", err)
		}
	}

	if err = d.Set("local_override", flattenSwitchControllerManagedSwitch8021XSettingsLocalOverride2edl(o["local-override"], d, "local_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-override"], "SwitchControllerManagedSwitch8021XSettings-LocalOverride"); ok {
			if err = d.Set("local_override", vv); err != nil {
				return fmt.Errorf("Error reading local_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_override: %v", err)
		}
	}

	if err = d.Set("mab_reauth", flattenSwitchControllerManagedSwitch8021XSettingsMabReauth2edl(o["mab-reauth"], d, "mab_reauth")); err != nil {
		if vv, ok := fortiAPIPatch(o["mab-reauth"], "SwitchControllerManagedSwitch8021XSettings-MabReauth"); ok {
			if err = d.Set("mab_reauth", vv); err != nil {
				return fmt.Errorf("Error reading mab_reauth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mab_reauth: %v", err)
		}
	}

	if err = d.Set("mac_called_station_delimiter", flattenSwitchControllerManagedSwitch8021XSettingsMacCalledStationDelimiter2edl(o["mac-called-station-delimiter"], d, "mac_called_station_delimiter")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-called-station-delimiter"], "SwitchControllerManagedSwitch8021XSettings-MacCalledStationDelimiter"); ok {
			if err = d.Set("mac_called_station_delimiter", vv); err != nil {
				return fmt.Errorf("Error reading mac_called_station_delimiter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_called_station_delimiter: %v", err)
		}
	}

	if err = d.Set("mac_calling_station_delimiter", flattenSwitchControllerManagedSwitch8021XSettingsMacCallingStationDelimiter2edl(o["mac-calling-station-delimiter"], d, "mac_calling_station_delimiter")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-calling-station-delimiter"], "SwitchControllerManagedSwitch8021XSettings-MacCallingStationDelimiter"); ok {
			if err = d.Set("mac_calling_station_delimiter", vv); err != nil {
				return fmt.Errorf("Error reading mac_calling_station_delimiter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_calling_station_delimiter: %v", err)
		}
	}

	if err = d.Set("mac_case", flattenSwitchControllerManagedSwitch8021XSettingsMacCase2edl(o["mac-case"], d, "mac_case")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-case"], "SwitchControllerManagedSwitch8021XSettings-MacCase"); ok {
			if err = d.Set("mac_case", vv); err != nil {
				return fmt.Errorf("Error reading mac_case: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_case: %v", err)
		}
	}

	if err = d.Set("mac_password_delimiter", flattenSwitchControllerManagedSwitch8021XSettingsMacPasswordDelimiter2edl(o["mac-password-delimiter"], d, "mac_password_delimiter")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-password-delimiter"], "SwitchControllerManagedSwitch8021XSettings-MacPasswordDelimiter"); ok {
			if err = d.Set("mac_password_delimiter", vv); err != nil {
				return fmt.Errorf("Error reading mac_password_delimiter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_password_delimiter: %v", err)
		}
	}

	if err = d.Set("mac_username_delimiter", flattenSwitchControllerManagedSwitch8021XSettingsMacUsernameDelimiter2edl(o["mac-username-delimiter"], d, "mac_username_delimiter")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-username-delimiter"], "SwitchControllerManagedSwitch8021XSettings-MacUsernameDelimiter"); ok {
			if err = d.Set("mac_username_delimiter", vv); err != nil {
				return fmt.Errorf("Error reading mac_username_delimiter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_username_delimiter: %v", err)
		}
	}

	if err = d.Set("max_reauth_attempt", flattenSwitchControllerManagedSwitch8021XSettingsMaxReauthAttempt2edl(o["max-reauth-attempt"], d, "max_reauth_attempt")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-reauth-attempt"], "SwitchControllerManagedSwitch8021XSettings-MaxReauthAttempt"); ok {
			if err = d.Set("max_reauth_attempt", vv); err != nil {
				return fmt.Errorf("Error reading max_reauth_attempt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_reauth_attempt: %v", err)
		}
	}

	if err = d.Set("reauth_period", flattenSwitchControllerManagedSwitch8021XSettingsReauthPeriod2edl(o["reauth-period"], d, "reauth_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["reauth-period"], "SwitchControllerManagedSwitch8021XSettings-ReauthPeriod"); ok {
			if err = d.Set("reauth_period", vv); err != nil {
				return fmt.Errorf("Error reading reauth_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reauth_period: %v", err)
		}
	}

	if err = d.Set("tx_period", flattenSwitchControllerManagedSwitch8021XSettingsTxPeriod2edl(o["tx-period"], d, "tx_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["tx-period"], "SwitchControllerManagedSwitch8021XSettings-TxPeriod"); ok {
			if err = d.Set("tx_period", vv); err != nil {
				return fmt.Errorf("Error reading tx_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tx_period: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerManagedSwitch8021XSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerManagedSwitch8021XSettingsLinkDownAuth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsLocalOverride2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsMabReauth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsMacCalledStationDelimiter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsMacCallingStationDelimiter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsMacCase2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsMacPasswordDelimiter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsMacUsernameDelimiter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsMaxReauthAttempt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsReauthPeriod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsTxPeriod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerManagedSwitch8021XSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("link_down_auth"); ok || d.HasChange("link_down_auth") {
		t, err := expandSwitchControllerManagedSwitch8021XSettingsLinkDownAuth2edl(d, v, "link_down_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-down-auth"] = t
		}
	}

	if v, ok := d.GetOk("local_override"); ok || d.HasChange("local_override") {
		t, err := expandSwitchControllerManagedSwitch8021XSettingsLocalOverride2edl(d, v, "local_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-override"] = t
		}
	}

	if v, ok := d.GetOk("mab_reauth"); ok || d.HasChange("mab_reauth") {
		t, err := expandSwitchControllerManagedSwitch8021XSettingsMabReauth2edl(d, v, "mab_reauth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mab-reauth"] = t
		}
	}

	if v, ok := d.GetOk("mac_called_station_delimiter"); ok || d.HasChange("mac_called_station_delimiter") {
		t, err := expandSwitchControllerManagedSwitch8021XSettingsMacCalledStationDelimiter2edl(d, v, "mac_called_station_delimiter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-called-station-delimiter"] = t
		}
	}

	if v, ok := d.GetOk("mac_calling_station_delimiter"); ok || d.HasChange("mac_calling_station_delimiter") {
		t, err := expandSwitchControllerManagedSwitch8021XSettingsMacCallingStationDelimiter2edl(d, v, "mac_calling_station_delimiter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-calling-station-delimiter"] = t
		}
	}

	if v, ok := d.GetOk("mac_case"); ok || d.HasChange("mac_case") {
		t, err := expandSwitchControllerManagedSwitch8021XSettingsMacCase2edl(d, v, "mac_case")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-case"] = t
		}
	}

	if v, ok := d.GetOk("mac_password_delimiter"); ok || d.HasChange("mac_password_delimiter") {
		t, err := expandSwitchControllerManagedSwitch8021XSettingsMacPasswordDelimiter2edl(d, v, "mac_password_delimiter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-password-delimiter"] = t
		}
	}

	if v, ok := d.GetOk("mac_username_delimiter"); ok || d.HasChange("mac_username_delimiter") {
		t, err := expandSwitchControllerManagedSwitch8021XSettingsMacUsernameDelimiter2edl(d, v, "mac_username_delimiter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-username-delimiter"] = t
		}
	}

	if v, ok := d.GetOk("max_reauth_attempt"); ok || d.HasChange("max_reauth_attempt") {
		t, err := expandSwitchControllerManagedSwitch8021XSettingsMaxReauthAttempt2edl(d, v, "max_reauth_attempt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-reauth-attempt"] = t
		}
	}

	if v, ok := d.GetOk("reauth_period"); ok || d.HasChange("reauth_period") {
		t, err := expandSwitchControllerManagedSwitch8021XSettingsReauthPeriod2edl(d, v, "reauth_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reauth-period"] = t
		}
	}

	if v, ok := d.GetOk("tx_period"); ok || d.HasChange("tx_period") {
		t, err := expandSwitchControllerManagedSwitch8021XSettingsTxPeriod2edl(d, v, "tx_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tx-period"] = t
		}
	}

	return &obj, nil
}
