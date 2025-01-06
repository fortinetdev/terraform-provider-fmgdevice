// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Independent upgrades for managed devices.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemDeviceUpgrade() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDeviceUpgradeCreate,
		Read:   resourceSystemDeviceUpgradeRead,
		Update: resourceSystemDeviceUpgradeUpdate,
		Delete: resourceSystemDeviceUpgradeDelete,

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
			"device_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"failure_reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ha_reboot_controller": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ignore_signing_errors": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"known_ha_members": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"serial": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"maximum_minutes": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"serial": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"setup_time": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"time": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"timing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upgrade_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemDeviceUpgradeCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemDeviceUpgrade(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemDeviceUpgrade resource while getting object: %v", err)
	}

	_, err = c.CreateSystemDeviceUpgrade(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemDeviceUpgrade resource: %v", err)
	}

	d.SetId(getStringKey(d, "serial"))

	return resourceSystemDeviceUpgradeRead(d, m)
}

func resourceSystemDeviceUpgradeUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemDeviceUpgrade(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemDeviceUpgrade resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemDeviceUpgrade(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemDeviceUpgrade resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "serial"))

	return resourceSystemDeviceUpgradeRead(d, m)
}

func resourceSystemDeviceUpgradeDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	err = c.DeleteSystemDeviceUpgrade(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDeviceUpgrade resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDeviceUpgradeRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	paradict["device"] = device_name

	o, err := c.ReadSystemDeviceUpgrade(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemDeviceUpgrade resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDeviceUpgrade(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemDeviceUpgrade resource from API: %v", err)
	}
	return nil
}

func flattenSystemDeviceUpgradeDeviceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDeviceUpgradeFailureReason(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDeviceUpgradeHaRebootController(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDeviceUpgradeIgnoreSigningErrors(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDeviceUpgradeKnownHaMembers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial"
		if _, ok := i["serial"]; ok {
			v := flattenSystemDeviceUpgradeKnownHaMembersSerial(i["serial"], d, pre_append)
			tmp["serial"] = fortiAPISubPartPatch(v, "SystemDeviceUpgrade-KnownHaMembers-Serial")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemDeviceUpgradeKnownHaMembersSerial(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDeviceUpgradeMaximumMinutes(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDeviceUpgradeSerial(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDeviceUpgradeSetupTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDeviceUpgradeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDeviceUpgradeTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDeviceUpgradeTiming(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDeviceUpgradeUpgradePath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemDeviceUpgrade(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("device_type", flattenSystemDeviceUpgradeDeviceType(o["device-type"], d, "device_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-type"], "SystemDeviceUpgrade-DeviceType"); ok {
			if err = d.Set("device_type", vv); err != nil {
				return fmt.Errorf("Error reading device_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_type: %v", err)
		}
	}

	if err = d.Set("failure_reason", flattenSystemDeviceUpgradeFailureReason(o["failure-reason"], d, "failure_reason")); err != nil {
		if vv, ok := fortiAPIPatch(o["failure-reason"], "SystemDeviceUpgrade-FailureReason"); ok {
			if err = d.Set("failure_reason", vv); err != nil {
				return fmt.Errorf("Error reading failure_reason: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading failure_reason: %v", err)
		}
	}

	if err = d.Set("ha_reboot_controller", flattenSystemDeviceUpgradeHaRebootController(o["ha-reboot-controller"], d, "ha_reboot_controller")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-reboot-controller"], "SystemDeviceUpgrade-HaRebootController"); ok {
			if err = d.Set("ha_reboot_controller", vv); err != nil {
				return fmt.Errorf("Error reading ha_reboot_controller: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_reboot_controller: %v", err)
		}
	}

	if err = d.Set("ignore_signing_errors", flattenSystemDeviceUpgradeIgnoreSigningErrors(o["ignore-signing-errors"], d, "ignore_signing_errors")); err != nil {
		if vv, ok := fortiAPIPatch(o["ignore-signing-errors"], "SystemDeviceUpgrade-IgnoreSigningErrors"); ok {
			if err = d.Set("ignore_signing_errors", vv); err != nil {
				return fmt.Errorf("Error reading ignore_signing_errors: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ignore_signing_errors: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("known_ha_members", flattenSystemDeviceUpgradeKnownHaMembers(o["known-ha-members"], d, "known_ha_members")); err != nil {
			if vv, ok := fortiAPIPatch(o["known-ha-members"], "SystemDeviceUpgrade-KnownHaMembers"); ok {
				if err = d.Set("known_ha_members", vv); err != nil {
					return fmt.Errorf("Error reading known_ha_members: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading known_ha_members: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("known_ha_members"); ok {
			if err = d.Set("known_ha_members", flattenSystemDeviceUpgradeKnownHaMembers(o["known-ha-members"], d, "known_ha_members")); err != nil {
				if vv, ok := fortiAPIPatch(o["known-ha-members"], "SystemDeviceUpgrade-KnownHaMembers"); ok {
					if err = d.Set("known_ha_members", vv); err != nil {
						return fmt.Errorf("Error reading known_ha_members: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading known_ha_members: %v", err)
				}
			}
		}
	}

	if err = d.Set("maximum_minutes", flattenSystemDeviceUpgradeMaximumMinutes(o["maximum-minutes"], d, "maximum_minutes")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-minutes"], "SystemDeviceUpgrade-MaximumMinutes"); ok {
			if err = d.Set("maximum_minutes", vv); err != nil {
				return fmt.Errorf("Error reading maximum_minutes: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_minutes: %v", err)
		}
	}

	if err = d.Set("serial", flattenSystemDeviceUpgradeSerial(o["serial"], d, "serial")); err != nil {
		if vv, ok := fortiAPIPatch(o["serial"], "SystemDeviceUpgrade-Serial"); ok {
			if err = d.Set("serial", vv); err != nil {
				return fmt.Errorf("Error reading serial: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading serial: %v", err)
		}
	}

	if err = d.Set("setup_time", flattenSystemDeviceUpgradeSetupTime(o["setup-time"], d, "setup_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["setup-time"], "SystemDeviceUpgrade-SetupTime"); ok {
			if err = d.Set("setup_time", vv); err != nil {
				return fmt.Errorf("Error reading setup_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading setup_time: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemDeviceUpgradeStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemDeviceUpgrade-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("time", flattenSystemDeviceUpgradeTime(o["time"], d, "time")); err != nil {
		if vv, ok := fortiAPIPatch(o["time"], "SystemDeviceUpgrade-Time"); ok {
			if err = d.Set("time", vv); err != nil {
				return fmt.Errorf("Error reading time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading time: %v", err)
		}
	}

	if err = d.Set("timing", flattenSystemDeviceUpgradeTiming(o["timing"], d, "timing")); err != nil {
		if vv, ok := fortiAPIPatch(o["timing"], "SystemDeviceUpgrade-Timing"); ok {
			if err = d.Set("timing", vv); err != nil {
				return fmt.Errorf("Error reading timing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timing: %v", err)
		}
	}

	if err = d.Set("upgrade_path", flattenSystemDeviceUpgradeUpgradePath(o["upgrade-path"], d, "upgrade_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["upgrade-path"], "SystemDeviceUpgrade-UpgradePath"); ok {
			if err = d.Set("upgrade_path", vv); err != nil {
				return fmt.Errorf("Error reading upgrade_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upgrade_path: %v", err)
		}
	}

	return nil
}

func flattenSystemDeviceUpgradeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemDeviceUpgradeDeviceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDeviceUpgradeFailureReason(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDeviceUpgradeHaRebootController(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDeviceUpgradeIgnoreSigningErrors(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDeviceUpgradeKnownHaMembers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["serial"], _ = expandSystemDeviceUpgradeKnownHaMembersSerial(d, i["serial"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemDeviceUpgradeKnownHaMembersSerial(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDeviceUpgradeMaximumMinutes(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDeviceUpgradeSerial(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDeviceUpgradeSetupTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDeviceUpgradeStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDeviceUpgradeTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDeviceUpgradeTiming(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDeviceUpgradeUpgradePath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDeviceUpgrade(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("device_type"); ok || d.HasChange("device_type") {
		t, err := expandSystemDeviceUpgradeDeviceType(d, v, "device_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-type"] = t
		}
	}

	if v, ok := d.GetOk("failure_reason"); ok || d.HasChange("failure_reason") {
		t, err := expandSystemDeviceUpgradeFailureReason(d, v, "failure_reason")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["failure-reason"] = t
		}
	}

	if v, ok := d.GetOk("ha_reboot_controller"); ok || d.HasChange("ha_reboot_controller") {
		t, err := expandSystemDeviceUpgradeHaRebootController(d, v, "ha_reboot_controller")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-reboot-controller"] = t
		}
	}

	if v, ok := d.GetOk("ignore_signing_errors"); ok || d.HasChange("ignore_signing_errors") {
		t, err := expandSystemDeviceUpgradeIgnoreSigningErrors(d, v, "ignore_signing_errors")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ignore-signing-errors"] = t
		}
	}

	if v, ok := d.GetOk("known_ha_members"); ok || d.HasChange("known_ha_members") {
		t, err := expandSystemDeviceUpgradeKnownHaMembers(d, v, "known_ha_members")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["known-ha-members"] = t
		}
	}

	if v, ok := d.GetOk("maximum_minutes"); ok || d.HasChange("maximum_minutes") {
		t, err := expandSystemDeviceUpgradeMaximumMinutes(d, v, "maximum_minutes")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-minutes"] = t
		}
	}

	if v, ok := d.GetOk("serial"); ok || d.HasChange("serial") {
		t, err := expandSystemDeviceUpgradeSerial(d, v, "serial")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial"] = t
		}
	}

	if v, ok := d.GetOk("setup_time"); ok || d.HasChange("setup_time") {
		t, err := expandSystemDeviceUpgradeSetupTime(d, v, "setup_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["setup-time"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemDeviceUpgradeStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("time"); ok || d.HasChange("time") {
		t, err := expandSystemDeviceUpgradeTime(d, v, "time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["time"] = t
		}
	}

	if v, ok := d.GetOk("timing"); ok || d.HasChange("timing") {
		t, err := expandSystemDeviceUpgradeTiming(d, v, "timing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timing"] = t
		}
	}

	if v, ok := d.GetOk("upgrade_path"); ok || d.HasChange("upgrade_path") {
		t, err := expandSystemDeviceUpgradeUpgradePath(d, v, "upgrade_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upgrade-path"] = t
		}
	}

	return &obj, nil
}
