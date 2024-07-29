// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure quarantine support.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserQuarantine() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserQuarantineUpdate,
		Read:   resourceUserQuarantineRead,
		Update: resourceUserQuarantineUpdate,
		Delete: resourceUserQuarantineDelete,

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
			"firewall_groups": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"quarantine": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"targets": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"entry": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"macs": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"description": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"drop": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"entry_id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"mac": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"parent": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"traffic_policy": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceUserQuarantineUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserQuarantine(d)
	if err != nil {
		return fmt.Errorf("Error updating UserQuarantine resource while getting object: %v", err)
	}

	_, err = c.UpdateUserQuarantine(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating UserQuarantine resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("UserQuarantine")

	return resourceUserQuarantineRead(d, m)
}

func resourceUserQuarantineDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteUserQuarantine(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting UserQuarantine resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserQuarantineRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserQuarantine(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading UserQuarantine resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserQuarantine(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserQuarantine resource from API: %v", err)
	}
	return nil
}

func flattenUserQuarantineFirewallGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserQuarantineQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserQuarantineTargets(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenUserQuarantineTargetsDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "UserQuarantine-Targets-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entry"
		if _, ok := i["entry"]; ok {
			v := flattenUserQuarantineTargetsEntry(i["entry"], d, pre_append)
			tmp["entry"] = fortiAPISubPartPatch(v, "UserQuarantine-Targets-Entry")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "macs"
		if _, ok := i["macs"]; ok {
			v := flattenUserQuarantineTargetsMacs(i["macs"], d, pre_append)
			tmp["macs"] = fortiAPISubPartPatch(v, "UserQuarantine-Targets-Macs")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenUserQuarantineTargetsDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserQuarantineTargetsEntry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserQuarantineTargetsMacs(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenUserQuarantineTargetsMacsDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "UserQuarantineTargets-Macs-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "drop"
		if _, ok := i["drop"]; ok {
			v := flattenUserQuarantineTargetsMacsDrop(i["drop"], d, pre_append)
			tmp["drop"] = fortiAPISubPartPatch(v, "UserQuarantineTargets-Macs-Drop")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entry_id"
		if _, ok := i["entry-id"]; ok {
			v := flattenUserQuarantineTargetsMacsEntryId(i["entry-id"], d, pre_append)
			tmp["entry_id"] = fortiAPISubPartPatch(v, "UserQuarantineTargets-Macs-EntryId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			v := flattenUserQuarantineTargetsMacsMac(i["mac"], d, pre_append)
			tmp["mac"] = fortiAPISubPartPatch(v, "UserQuarantineTargets-Macs-Mac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "parent"
		if _, ok := i["parent"]; ok {
			v := flattenUserQuarantineTargetsMacsParent(i["parent"], d, pre_append)
			tmp["parent"] = fortiAPISubPartPatch(v, "UserQuarantineTargets-Macs-Parent")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenUserQuarantineTargetsMacsDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserQuarantineTargetsMacsDrop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserQuarantineTargetsMacsEntryId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserQuarantineTargetsMacsMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return case_insensitive(v, d.Get(pre))
}

func flattenUserQuarantineTargetsMacsParent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserQuarantineTrafficPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectUserQuarantine(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("firewall_groups", flattenUserQuarantineFirewallGroups(o["firewall-groups"], d, "firewall_groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["firewall-groups"], "UserQuarantine-FirewallGroups"); ok {
			if err = d.Set("firewall_groups", vv); err != nil {
				return fmt.Errorf("Error reading firewall_groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading firewall_groups: %v", err)
		}
	}

	if err = d.Set("quarantine", flattenUserQuarantineQuarantine(o["quarantine"], d, "quarantine")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine"], "UserQuarantine-Quarantine"); ok {
			if err = d.Set("quarantine", vv); err != nil {
				return fmt.Errorf("Error reading quarantine: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("targets", flattenUserQuarantineTargets(o["targets"], d, "targets")); err != nil {
			if vv, ok := fortiAPIPatch(o["targets"], "UserQuarantine-Targets"); ok {
				if err = d.Set("targets", vv); err != nil {
					return fmt.Errorf("Error reading targets: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading targets: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("targets"); ok {
			if err = d.Set("targets", flattenUserQuarantineTargets(o["targets"], d, "targets")); err != nil {
				if vv, ok := fortiAPIPatch(o["targets"], "UserQuarantine-Targets"); ok {
					if err = d.Set("targets", vv); err != nil {
						return fmt.Errorf("Error reading targets: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading targets: %v", err)
				}
			}
		}
	}

	if err = d.Set("traffic_policy", flattenUserQuarantineTrafficPolicy(o["traffic-policy"], d, "traffic_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-policy"], "UserQuarantine-TrafficPolicy"); ok {
			if err = d.Set("traffic_policy", vv); err != nil {
				return fmt.Errorf("Error reading traffic_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_policy: %v", err)
		}
	}

	return nil
}

func flattenUserQuarantineFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserQuarantineFirewallGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserQuarantineQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTargets(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandUserQuarantineTargetsDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entry"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["entry"], _ = expandUserQuarantineTargetsEntry(d, i["entry"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "macs"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandUserQuarantineTargetsMacs(d, i["macs"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["macs"] = t
			}
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandUserQuarantineTargetsDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTargetsEntry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTargetsMacs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandUserQuarantineTargetsMacsDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "drop"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["drop"], _ = expandUserQuarantineTargetsMacsDrop(d, i["drop"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entry_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["entry-id"], _ = expandUserQuarantineTargetsMacsEntryId(d, i["entry_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac"], _ = expandUserQuarantineTargetsMacsMac(d, i["mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "parent"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["parent"], _ = expandUserQuarantineTargetsMacsParent(d, i["parent"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandUserQuarantineTargetsMacsDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTargetsMacsDrop(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTargetsMacsEntryId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTargetsMacsMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTargetsMacsParent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTrafficPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectUserQuarantine(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("firewall_groups"); ok || d.HasChange("firewall_groups") {
		t, err := expandUserQuarantineFirewallGroups(d, v, "firewall_groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firewall-groups"] = t
		}
	}

	if v, ok := d.GetOk("quarantine"); ok || d.HasChange("quarantine") {
		t, err := expandUserQuarantineQuarantine(d, v, "quarantine")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine"] = t
		}
	}

	if v, ok := d.GetOk("targets"); ok || d.HasChange("targets") {
		t, err := expandUserQuarantineTargets(d, v, "targets")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["targets"] = t
		}
	}

	if v, ok := d.GetOk("traffic_policy"); ok || d.HasChange("traffic_policy") {
		t, err := expandUserQuarantineTrafficPolicy(d, v, "traffic_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-policy"] = t
		}
	}

	return &obj, nil
}
