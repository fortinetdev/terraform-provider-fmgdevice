// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: VDOM wireless controller configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerSettingUpdate,
		Read:   resourceWirelessControllerSettingRead,
		Update: resourceWirelessControllerSettingUpdate,
		Delete: resourceWirelessControllerSettingDelete,

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
			"account_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"country": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"darrp_optimize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"darrp_optimize_schedules": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"device_holdoff": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"device_idle": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"device_weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"duplicate_ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fake_ssid_action": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fapc_compatibility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"firmware_provision_on_authorization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"offending_ssid": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ssid_pattern": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"phishing_ssid_detect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rolling_wtp_upgrade": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wfa_compatibility": &schema.Schema{
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

func resourceWirelessControllerSettingUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWirelessControllerSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerSetting(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WirelessControllerSetting")

	return resourceWirelessControllerSettingRead(d, m)
}

func resourceWirelessControllerSettingDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWirelessControllerSetting(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerSettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerSetting resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerSettingAccountId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSettingCountry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSettingDarrpOptimize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSettingDarrpOptimizeSchedules(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerSettingDeviceHoldoff(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSettingDeviceIdle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSettingDeviceWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSettingDuplicateSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSettingFakeSsidAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerSettingFapcCompatibility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSettingFirmwareProvisionOnAuthorization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSettingOffendingSsid(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenWirelessControllerSettingOffendingSsidAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "WirelessControllerSetting-OffendingSsid-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWirelessControllerSettingOffendingSsidId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WirelessControllerSetting-OffendingSsid-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssid_pattern"
		if _, ok := i["ssid-pattern"]; ok {
			v := flattenWirelessControllerSettingOffendingSsidSsidPattern(i["ssid-pattern"], d, pre_append)
			tmp["ssid_pattern"] = fortiAPISubPartPatch(v, "WirelessControllerSetting-OffendingSsid-SsidPattern")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerSettingOffendingSsidAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerSettingOffendingSsidId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSettingOffendingSsidSsidPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSettingPhishingSsidDetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSettingRollingWtpUpgrade(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSettingWfaCompatibility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("account_id", flattenWirelessControllerSettingAccountId(o["account-id"], d, "account_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["account-id"], "WirelessControllerSetting-AccountId"); ok {
			if err = d.Set("account_id", vv); err != nil {
				return fmt.Errorf("Error reading account_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading account_id: %v", err)
		}
	}

	if err = d.Set("country", flattenWirelessControllerSettingCountry(o["country"], d, "country")); err != nil {
		if vv, ok := fortiAPIPatch(o["country"], "WirelessControllerSetting-Country"); ok {
			if err = d.Set("country", vv); err != nil {
				return fmt.Errorf("Error reading country: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading country: %v", err)
		}
	}

	if err = d.Set("darrp_optimize", flattenWirelessControllerSettingDarrpOptimize(o["darrp-optimize"], d, "darrp_optimize")); err != nil {
		if vv, ok := fortiAPIPatch(o["darrp-optimize"], "WirelessControllerSetting-DarrpOptimize"); ok {
			if err = d.Set("darrp_optimize", vv); err != nil {
				return fmt.Errorf("Error reading darrp_optimize: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading darrp_optimize: %v", err)
		}
	}

	if err = d.Set("darrp_optimize_schedules", flattenWirelessControllerSettingDarrpOptimizeSchedules(o["darrp-optimize-schedules"], d, "darrp_optimize_schedules")); err != nil {
		if vv, ok := fortiAPIPatch(o["darrp-optimize-schedules"], "WirelessControllerSetting-DarrpOptimizeSchedules"); ok {
			if err = d.Set("darrp_optimize_schedules", vv); err != nil {
				return fmt.Errorf("Error reading darrp_optimize_schedules: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading darrp_optimize_schedules: %v", err)
		}
	}

	if err = d.Set("device_holdoff", flattenWirelessControllerSettingDeviceHoldoff(o["device-holdoff"], d, "device_holdoff")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-holdoff"], "WirelessControllerSetting-DeviceHoldoff"); ok {
			if err = d.Set("device_holdoff", vv); err != nil {
				return fmt.Errorf("Error reading device_holdoff: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_holdoff: %v", err)
		}
	}

	if err = d.Set("device_idle", flattenWirelessControllerSettingDeviceIdle(o["device-idle"], d, "device_idle")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-idle"], "WirelessControllerSetting-DeviceIdle"); ok {
			if err = d.Set("device_idle", vv); err != nil {
				return fmt.Errorf("Error reading device_idle: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_idle: %v", err)
		}
	}

	if err = d.Set("device_weight", flattenWirelessControllerSettingDeviceWeight(o["device-weight"], d, "device_weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-weight"], "WirelessControllerSetting-DeviceWeight"); ok {
			if err = d.Set("device_weight", vv); err != nil {
				return fmt.Errorf("Error reading device_weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_weight: %v", err)
		}
	}

	if err = d.Set("duplicate_ssid", flattenWirelessControllerSettingDuplicateSsid(o["duplicate-ssid"], d, "duplicate_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["duplicate-ssid"], "WirelessControllerSetting-DuplicateSsid"); ok {
			if err = d.Set("duplicate_ssid", vv); err != nil {
				return fmt.Errorf("Error reading duplicate_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading duplicate_ssid: %v", err)
		}
	}

	if err = d.Set("fake_ssid_action", flattenWirelessControllerSettingFakeSsidAction(o["fake-ssid-action"], d, "fake_ssid_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["fake-ssid-action"], "WirelessControllerSetting-FakeSsidAction"); ok {
			if err = d.Set("fake_ssid_action", vv); err != nil {
				return fmt.Errorf("Error reading fake_ssid_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fake_ssid_action: %v", err)
		}
	}

	if err = d.Set("fapc_compatibility", flattenWirelessControllerSettingFapcCompatibility(o["fapc-compatibility"], d, "fapc_compatibility")); err != nil {
		if vv, ok := fortiAPIPatch(o["fapc-compatibility"], "WirelessControllerSetting-FapcCompatibility"); ok {
			if err = d.Set("fapc_compatibility", vv); err != nil {
				return fmt.Errorf("Error reading fapc_compatibility: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fapc_compatibility: %v", err)
		}
	}

	if err = d.Set("firmware_provision_on_authorization", flattenWirelessControllerSettingFirmwareProvisionOnAuthorization(o["firmware-provision-on-authorization"], d, "firmware_provision_on_authorization")); err != nil {
		if vv, ok := fortiAPIPatch(o["firmware-provision-on-authorization"], "WirelessControllerSetting-FirmwareProvisionOnAuthorization"); ok {
			if err = d.Set("firmware_provision_on_authorization", vv); err != nil {
				return fmt.Errorf("Error reading firmware_provision_on_authorization: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading firmware_provision_on_authorization: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("offending_ssid", flattenWirelessControllerSettingOffendingSsid(o["offending-ssid"], d, "offending_ssid")); err != nil {
			if vv, ok := fortiAPIPatch(o["offending-ssid"], "WirelessControllerSetting-OffendingSsid"); ok {
				if err = d.Set("offending_ssid", vv); err != nil {
					return fmt.Errorf("Error reading offending_ssid: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading offending_ssid: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("offending_ssid"); ok {
			if err = d.Set("offending_ssid", flattenWirelessControllerSettingOffendingSsid(o["offending-ssid"], d, "offending_ssid")); err != nil {
				if vv, ok := fortiAPIPatch(o["offending-ssid"], "WirelessControllerSetting-OffendingSsid"); ok {
					if err = d.Set("offending_ssid", vv); err != nil {
						return fmt.Errorf("Error reading offending_ssid: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading offending_ssid: %v", err)
				}
			}
		}
	}

	if err = d.Set("phishing_ssid_detect", flattenWirelessControllerSettingPhishingSsidDetect(o["phishing-ssid-detect"], d, "phishing_ssid_detect")); err != nil {
		if vv, ok := fortiAPIPatch(o["phishing-ssid-detect"], "WirelessControllerSetting-PhishingSsidDetect"); ok {
			if err = d.Set("phishing_ssid_detect", vv); err != nil {
				return fmt.Errorf("Error reading phishing_ssid_detect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading phishing_ssid_detect: %v", err)
		}
	}

	if err = d.Set("rolling_wtp_upgrade", flattenWirelessControllerSettingRollingWtpUpgrade(o["rolling-wtp-upgrade"], d, "rolling_wtp_upgrade")); err != nil {
		if vv, ok := fortiAPIPatch(o["rolling-wtp-upgrade"], "WirelessControllerSetting-RollingWtpUpgrade"); ok {
			if err = d.Set("rolling_wtp_upgrade", vv); err != nil {
				return fmt.Errorf("Error reading rolling_wtp_upgrade: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rolling_wtp_upgrade: %v", err)
		}
	}

	if err = d.Set("wfa_compatibility", flattenWirelessControllerSettingWfaCompatibility(o["wfa-compatibility"], d, "wfa_compatibility")); err != nil {
		if vv, ok := fortiAPIPatch(o["wfa-compatibility"], "WirelessControllerSetting-WfaCompatibility"); ok {
			if err = d.Set("wfa_compatibility", vv); err != nil {
				return fmt.Errorf("Error reading wfa_compatibility: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wfa_compatibility: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerSettingAccountId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingCountry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingDarrpOptimize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingDarrpOptimizeSchedules(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerSettingDeviceHoldoff(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingDeviceIdle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingDeviceWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingDuplicateSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingFakeSsidAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerSettingFapcCompatibility(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingFirmwareProvisionOnAuthorization(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingOffendingSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandWirelessControllerSettingOffendingSsidAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWirelessControllerSettingOffendingSsidId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssid_pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssid-pattern"], _ = expandWirelessControllerSettingOffendingSsidSsidPattern(d, i["ssid_pattern"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerSettingOffendingSsidAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerSettingOffendingSsidId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingOffendingSsidSsidPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingPhishingSsidDetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingRollingWtpUpgrade(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingWfaCompatibility(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("account_id"); ok || d.HasChange("account_id") {
		t, err := expandWirelessControllerSettingAccountId(d, v, "account_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["account-id"] = t
		}
	}

	if v, ok := d.GetOk("country"); ok || d.HasChange("country") {
		t, err := expandWirelessControllerSettingCountry(d, v, "country")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["country"] = t
		}
	}

	if v, ok := d.GetOk("darrp_optimize"); ok || d.HasChange("darrp_optimize") {
		t, err := expandWirelessControllerSettingDarrpOptimize(d, v, "darrp_optimize")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["darrp-optimize"] = t
		}
	}

	if v, ok := d.GetOk("darrp_optimize_schedules"); ok || d.HasChange("darrp_optimize_schedules") {
		t, err := expandWirelessControllerSettingDarrpOptimizeSchedules(d, v, "darrp_optimize_schedules")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["darrp-optimize-schedules"] = t
		}
	}

	if v, ok := d.GetOk("device_holdoff"); ok || d.HasChange("device_holdoff") {
		t, err := expandWirelessControllerSettingDeviceHoldoff(d, v, "device_holdoff")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-holdoff"] = t
		}
	}

	if v, ok := d.GetOk("device_idle"); ok || d.HasChange("device_idle") {
		t, err := expandWirelessControllerSettingDeviceIdle(d, v, "device_idle")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-idle"] = t
		}
	}

	if v, ok := d.GetOk("device_weight"); ok || d.HasChange("device_weight") {
		t, err := expandWirelessControllerSettingDeviceWeight(d, v, "device_weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-weight"] = t
		}
	}

	if v, ok := d.GetOk("duplicate_ssid"); ok || d.HasChange("duplicate_ssid") {
		t, err := expandWirelessControllerSettingDuplicateSsid(d, v, "duplicate_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["duplicate-ssid"] = t
		}
	}

	if v, ok := d.GetOk("fake_ssid_action"); ok || d.HasChange("fake_ssid_action") {
		t, err := expandWirelessControllerSettingFakeSsidAction(d, v, "fake_ssid_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fake-ssid-action"] = t
		}
	}

	if v, ok := d.GetOk("fapc_compatibility"); ok || d.HasChange("fapc_compatibility") {
		t, err := expandWirelessControllerSettingFapcCompatibility(d, v, "fapc_compatibility")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fapc-compatibility"] = t
		}
	}

	if v, ok := d.GetOk("firmware_provision_on_authorization"); ok || d.HasChange("firmware_provision_on_authorization") {
		t, err := expandWirelessControllerSettingFirmwareProvisionOnAuthorization(d, v, "firmware_provision_on_authorization")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firmware-provision-on-authorization"] = t
		}
	}

	if v, ok := d.GetOk("offending_ssid"); ok || d.HasChange("offending_ssid") {
		t, err := expandWirelessControllerSettingOffendingSsid(d, v, "offending_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["offending-ssid"] = t
		}
	}

	if v, ok := d.GetOk("phishing_ssid_detect"); ok || d.HasChange("phishing_ssid_detect") {
		t, err := expandWirelessControllerSettingPhishingSsidDetect(d, v, "phishing_ssid_detect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["phishing-ssid-detect"] = t
		}
	}

	if v, ok := d.GetOk("rolling_wtp_upgrade"); ok || d.HasChange("rolling_wtp_upgrade") {
		t, err := expandWirelessControllerSettingRollingWtpUpgrade(d, v, "rolling_wtp_upgrade")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rolling-wtp-upgrade"] = t
		}
	}

	if v, ok := d.GetOk("wfa_compatibility"); ok || d.HasChange("wfa_compatibility") {
		t, err := expandWirelessControllerSettingWfaCompatibility(d, v, "wfa_compatibility")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wfa-compatibility"] = t
		}
	}

	return &obj, nil
}
