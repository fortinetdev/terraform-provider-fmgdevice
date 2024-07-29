// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Coordinate federated upgrades within the Security Fabric.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemFederatedUpgrade() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemFederatedUpgradeUpdate,
		Read:   resourceSystemFederatedUpgradeRead,
		Update: resourceSystemFederatedUpgradeUpdate,
		Delete: resourceSystemFederatedUpgradeDelete,

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
			"failure_device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"next_path_index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"node_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"coordinating_fortigate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"device_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"maximum_minutes": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"serial": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"setup_time": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
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
					},
				},
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upgrade_id": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceSystemFederatedUpgradeUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemFederatedUpgrade(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemFederatedUpgrade resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemFederatedUpgrade(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemFederatedUpgrade resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemFederatedUpgrade")

	return resourceSystemFederatedUpgradeRead(d, m)
}

func resourceSystemFederatedUpgradeDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemFederatedUpgrade(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemFederatedUpgrade resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFederatedUpgradeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemFederatedUpgrade(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemFederatedUpgrade resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFederatedUpgrade(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemFederatedUpgrade resource from API: %v", err)
	}
	return nil
}

func flattenSystemFederatedUpgradeFailureDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeFailureReason(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeHaRebootController(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeIgnoreSigningErrors(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeKnownHaMembers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemFederatedUpgradeKnownHaMembersSerial(i["serial"], d, pre_append)
			tmp["serial"] = fortiAPISubPartPatch(v, "SystemFederatedUpgrade-KnownHaMembers-Serial")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemFederatedUpgradeKnownHaMembersSerial(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeNextPathIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeNodeList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "coordinating_fortigate"
		if _, ok := i["coordinating-fortigate"]; ok {
			v := flattenSystemFederatedUpgradeNodeListCoordinatingFortigate(i["coordinating-fortigate"], d, pre_append)
			tmp["coordinating_fortigate"] = fortiAPISubPartPatch(v, "SystemFederatedUpgrade-NodeList-CoordinatingFortigate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_type"
		if _, ok := i["device-type"]; ok {
			v := flattenSystemFederatedUpgradeNodeListDeviceType(i["device-type"], d, pre_append)
			tmp["device_type"] = fortiAPISubPartPatch(v, "SystemFederatedUpgrade-NodeList-DeviceType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_minutes"
		if _, ok := i["maximum-minutes"]; ok {
			v := flattenSystemFederatedUpgradeNodeListMaximumMinutes(i["maximum-minutes"], d, pre_append)
			tmp["maximum_minutes"] = fortiAPISubPartPatch(v, "SystemFederatedUpgrade-NodeList-MaximumMinutes")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial"
		if _, ok := i["serial"]; ok {
			v := flattenSystemFederatedUpgradeNodeListSerial(i["serial"], d, pre_append)
			tmp["serial"] = fortiAPISubPartPatch(v, "SystemFederatedUpgrade-NodeList-Serial")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "setup_time"
		if _, ok := i["setup-time"]; ok {
			v := flattenSystemFederatedUpgradeNodeListSetupTime(i["setup-time"], d, pre_append)
			tmp["setup_time"] = fortiAPISubPartPatch(v, "SystemFederatedUpgrade-NodeList-SetupTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "time"
		if _, ok := i["time"]; ok {
			v := flattenSystemFederatedUpgradeNodeListTime(i["time"], d, pre_append)
			tmp["time"] = fortiAPISubPartPatch(v, "SystemFederatedUpgrade-NodeList-Time")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "timing"
		if _, ok := i["timing"]; ok {
			v := flattenSystemFederatedUpgradeNodeListTiming(i["timing"], d, pre_append)
			tmp["timing"] = fortiAPISubPartPatch(v, "SystemFederatedUpgrade-NodeList-Timing")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "upgrade_path"
		if _, ok := i["upgrade-path"]; ok {
			v := flattenSystemFederatedUpgradeNodeListUpgradePath(i["upgrade-path"], d, pre_append)
			tmp["upgrade_path"] = fortiAPISubPartPatch(v, "SystemFederatedUpgrade-NodeList-UpgradePath")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemFederatedUpgradeNodeListCoordinatingFortigate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeNodeListDeviceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeNodeListMaximumMinutes(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeNodeListSerial(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeNodeListSetupTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFederatedUpgradeNodeListTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFederatedUpgradeNodeListTiming(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeNodeListUpgradePath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeUpgradeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemFederatedUpgrade(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("failure_device", flattenSystemFederatedUpgradeFailureDevice(o["failure-device"], d, "failure_device")); err != nil {
		if vv, ok := fortiAPIPatch(o["failure-device"], "SystemFederatedUpgrade-FailureDevice"); ok {
			if err = d.Set("failure_device", vv); err != nil {
				return fmt.Errorf("Error reading failure_device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading failure_device: %v", err)
		}
	}

	if err = d.Set("failure_reason", flattenSystemFederatedUpgradeFailureReason(o["failure-reason"], d, "failure_reason")); err != nil {
		if vv, ok := fortiAPIPatch(o["failure-reason"], "SystemFederatedUpgrade-FailureReason"); ok {
			if err = d.Set("failure_reason", vv); err != nil {
				return fmt.Errorf("Error reading failure_reason: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading failure_reason: %v", err)
		}
	}

	if err = d.Set("ha_reboot_controller", flattenSystemFederatedUpgradeHaRebootController(o["ha-reboot-controller"], d, "ha_reboot_controller")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-reboot-controller"], "SystemFederatedUpgrade-HaRebootController"); ok {
			if err = d.Set("ha_reboot_controller", vv); err != nil {
				return fmt.Errorf("Error reading ha_reboot_controller: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_reboot_controller: %v", err)
		}
	}

	if err = d.Set("ignore_signing_errors", flattenSystemFederatedUpgradeIgnoreSigningErrors(o["ignore-signing-errors"], d, "ignore_signing_errors")); err != nil {
		if vv, ok := fortiAPIPatch(o["ignore-signing-errors"], "SystemFederatedUpgrade-IgnoreSigningErrors"); ok {
			if err = d.Set("ignore_signing_errors", vv); err != nil {
				return fmt.Errorf("Error reading ignore_signing_errors: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ignore_signing_errors: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("known_ha_members", flattenSystemFederatedUpgradeKnownHaMembers(o["known-ha-members"], d, "known_ha_members")); err != nil {
			if vv, ok := fortiAPIPatch(o["known-ha-members"], "SystemFederatedUpgrade-KnownHaMembers"); ok {
				if err = d.Set("known_ha_members", vv); err != nil {
					return fmt.Errorf("Error reading known_ha_members: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading known_ha_members: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("known_ha_members"); ok {
			if err = d.Set("known_ha_members", flattenSystemFederatedUpgradeKnownHaMembers(o["known-ha-members"], d, "known_ha_members")); err != nil {
				if vv, ok := fortiAPIPatch(o["known-ha-members"], "SystemFederatedUpgrade-KnownHaMembers"); ok {
					if err = d.Set("known_ha_members", vv); err != nil {
						return fmt.Errorf("Error reading known_ha_members: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading known_ha_members: %v", err)
				}
			}
		}
	}

	if err = d.Set("next_path_index", flattenSystemFederatedUpgradeNextPathIndex(o["next-path-index"], d, "next_path_index")); err != nil {
		if vv, ok := fortiAPIPatch(o["next-path-index"], "SystemFederatedUpgrade-NextPathIndex"); ok {
			if err = d.Set("next_path_index", vv); err != nil {
				return fmt.Errorf("Error reading next_path_index: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading next_path_index: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("node_list", flattenSystemFederatedUpgradeNodeList(o["node-list"], d, "node_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["node-list"], "SystemFederatedUpgrade-NodeList"); ok {
				if err = d.Set("node_list", vv); err != nil {
					return fmt.Errorf("Error reading node_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading node_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("node_list"); ok {
			if err = d.Set("node_list", flattenSystemFederatedUpgradeNodeList(o["node-list"], d, "node_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["node-list"], "SystemFederatedUpgrade-NodeList"); ok {
					if err = d.Set("node_list", vv); err != nil {
						return fmt.Errorf("Error reading node_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading node_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("status", flattenSystemFederatedUpgradeStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemFederatedUpgrade-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("upgrade_id", flattenSystemFederatedUpgradeUpgradeId(o["upgrade-id"], d, "upgrade_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["upgrade-id"], "SystemFederatedUpgrade-UpgradeId"); ok {
			if err = d.Set("upgrade_id", vv); err != nil {
				return fmt.Errorf("Error reading upgrade_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upgrade_id: %v", err)
		}
	}

	return nil
}

func flattenSystemFederatedUpgradeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemFederatedUpgradeFailureDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeFailureReason(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeHaRebootController(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeIgnoreSigningErrors(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeKnownHaMembers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["serial"], _ = expandSystemFederatedUpgradeKnownHaMembersSerial(d, i["serial"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemFederatedUpgradeKnownHaMembersSerial(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeNextPathIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeNodeList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "coordinating_fortigate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["coordinating-fortigate"], _ = expandSystemFederatedUpgradeNodeListCoordinatingFortigate(d, i["coordinating_fortigate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["device-type"], _ = expandSystemFederatedUpgradeNodeListDeviceType(d, i["device_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_minutes"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-minutes"], _ = expandSystemFederatedUpgradeNodeListMaximumMinutes(d, i["maximum_minutes"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["serial"], _ = expandSystemFederatedUpgradeNodeListSerial(d, i["serial"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "setup_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["setup-time"], _ = expandSystemFederatedUpgradeNodeListSetupTime(d, i["setup_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["time"], _ = expandSystemFederatedUpgradeNodeListTime(d, i["time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "timing"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["timing"], _ = expandSystemFederatedUpgradeNodeListTiming(d, i["timing"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "upgrade_path"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["upgrade-path"], _ = expandSystemFederatedUpgradeNodeListUpgradePath(d, i["upgrade_path"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemFederatedUpgradeNodeListCoordinatingFortigate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeNodeListDeviceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeNodeListMaximumMinutes(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeNodeListSerial(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeNodeListSetupTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFederatedUpgradeNodeListTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFederatedUpgradeNodeListTiming(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeNodeListUpgradePath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeUpgradeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemFederatedUpgrade(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("failure_device"); ok || d.HasChange("failure_device") {
		t, err := expandSystemFederatedUpgradeFailureDevice(d, v, "failure_device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["failure-device"] = t
		}
	}

	if v, ok := d.GetOk("failure_reason"); ok || d.HasChange("failure_reason") {
		t, err := expandSystemFederatedUpgradeFailureReason(d, v, "failure_reason")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["failure-reason"] = t
		}
	}

	if v, ok := d.GetOk("ha_reboot_controller"); ok || d.HasChange("ha_reboot_controller") {
		t, err := expandSystemFederatedUpgradeHaRebootController(d, v, "ha_reboot_controller")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-reboot-controller"] = t
		}
	}

	if v, ok := d.GetOk("ignore_signing_errors"); ok || d.HasChange("ignore_signing_errors") {
		t, err := expandSystemFederatedUpgradeIgnoreSigningErrors(d, v, "ignore_signing_errors")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ignore-signing-errors"] = t
		}
	}

	if v, ok := d.GetOk("known_ha_members"); ok || d.HasChange("known_ha_members") {
		t, err := expandSystemFederatedUpgradeKnownHaMembers(d, v, "known_ha_members")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["known-ha-members"] = t
		}
	}

	if v, ok := d.GetOk("next_path_index"); ok || d.HasChange("next_path_index") {
		t, err := expandSystemFederatedUpgradeNextPathIndex(d, v, "next_path_index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-path-index"] = t
		}
	}

	if v, ok := d.GetOk("node_list"); ok || d.HasChange("node_list") {
		t, err := expandSystemFederatedUpgradeNodeList(d, v, "node_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["node-list"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemFederatedUpgradeStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("upgrade_id"); ok || d.HasChange("upgrade_id") {
		t, err := expandSystemFederatedUpgradeUpgradeId(d, v, "upgrade_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upgrade-id"] = t
		}
	}

	return &obj, nil
}
