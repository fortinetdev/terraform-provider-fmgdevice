// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Quarantine entry to hold multiple MACs.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserQuarantineTargets() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserQuarantineTargetsCreate,
		Read:   resourceUserQuarantineTargetsRead,
		Update: resourceUserQuarantineTargetsUpdate,
		Delete: resourceUserQuarantineTargetsDelete,

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
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"entry": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceUserQuarantineTargetsCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectUserQuarantineTargets(d)
	if err != nil {
		return fmt.Errorf("Error creating UserQuarantineTargets resource while getting object: %v", err)
	}

	_, err = c.CreateUserQuarantineTargets(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating UserQuarantineTargets resource: %v", err)
	}

	d.SetId(getStringKey(d, "entry"))

	return resourceUserQuarantineTargetsRead(d, m)
}

func resourceUserQuarantineTargetsUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectUserQuarantineTargets(d)
	if err != nil {
		return fmt.Errorf("Error updating UserQuarantineTargets resource while getting object: %v", err)
	}

	_, err = c.UpdateUserQuarantineTargets(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating UserQuarantineTargets resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "entry"))

	return resourceUserQuarantineTargetsRead(d, m)
}

func resourceUserQuarantineTargetsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteUserQuarantineTargets(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting UserQuarantineTargets resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserQuarantineTargetsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserQuarantineTargets(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading UserQuarantineTargets resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserQuarantineTargets(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserQuarantineTargets resource from API: %v", err)
	}
	return nil
}

func flattenUserQuarantineTargetsDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserQuarantineTargetsEntry2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserQuarantineTargetsMacs2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenUserQuarantineTargetsMacsDescription2edl(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "UserQuarantineTargets-Macs-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "drop"
		if _, ok := i["drop"]; ok {
			v := flattenUserQuarantineTargetsMacsDrop2edl(i["drop"], d, pre_append)
			tmp["drop"] = fortiAPISubPartPatch(v, "UserQuarantineTargets-Macs-Drop")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entry_id"
		if _, ok := i["entry-id"]; ok {
			v := flattenUserQuarantineTargetsMacsEntryId2edl(i["entry-id"], d, pre_append)
			tmp["entry_id"] = fortiAPISubPartPatch(v, "UserQuarantineTargets-Macs-EntryId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			v := flattenUserQuarantineTargetsMacsMac2edl(i["mac"], d, pre_append)
			tmp["mac"] = fortiAPISubPartPatch(v, "UserQuarantineTargets-Macs-Mac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "parent"
		if _, ok := i["parent"]; ok {
			v := flattenUserQuarantineTargetsMacsParent2edl(i["parent"], d, pre_append)
			tmp["parent"] = fortiAPISubPartPatch(v, "UserQuarantineTargets-Macs-Parent")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenUserQuarantineTargetsMacsDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserQuarantineTargetsMacsDrop2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserQuarantineTargetsMacsEntryId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserQuarantineTargetsMacsMac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return case_insensitive(v, d.Get(pre))
}

func flattenUserQuarantineTargetsMacsParent2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectUserQuarantineTargets(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("description", flattenUserQuarantineTargetsDescription2edl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "UserQuarantineTargets-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("entry", flattenUserQuarantineTargetsEntry2edl(o["entry"], d, "entry")); err != nil {
		if vv, ok := fortiAPIPatch(o["entry"], "UserQuarantineTargets-Entry"); ok {
			if err = d.Set("entry", vv); err != nil {
				return fmt.Errorf("Error reading entry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading entry: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("macs", flattenUserQuarantineTargetsMacs2edl(o["macs"], d, "macs")); err != nil {
			if vv, ok := fortiAPIPatch(o["macs"], "UserQuarantineTargets-Macs"); ok {
				if err = d.Set("macs", vv); err != nil {
					return fmt.Errorf("Error reading macs: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading macs: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("macs"); ok {
			if err = d.Set("macs", flattenUserQuarantineTargetsMacs2edl(o["macs"], d, "macs")); err != nil {
				if vv, ok := fortiAPIPatch(o["macs"], "UserQuarantineTargets-Macs"); ok {
					if err = d.Set("macs", vv); err != nil {
						return fmt.Errorf("Error reading macs: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading macs: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenUserQuarantineTargetsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserQuarantineTargetsDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTargetsEntry2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTargetsMacs2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["description"], _ = expandUserQuarantineTargetsMacsDescription2edl(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "drop"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["drop"], _ = expandUserQuarantineTargetsMacsDrop2edl(d, i["drop"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entry_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["entry-id"], _ = expandUserQuarantineTargetsMacsEntryId2edl(d, i["entry_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac"], _ = expandUserQuarantineTargetsMacsMac2edl(d, i["mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "parent"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["parent"], _ = expandUserQuarantineTargetsMacsParent2edl(d, i["parent"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandUserQuarantineTargetsMacsDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTargetsMacsDrop2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTargetsMacsEntryId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTargetsMacsMac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTargetsMacsParent2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectUserQuarantineTargets(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandUserQuarantineTargetsDescription2edl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("entry"); ok || d.HasChange("entry") {
		t, err := expandUserQuarantineTargetsEntry2edl(d, v, "entry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entry"] = t
		}
	}

	if v, ok := d.GetOk("macs"); ok || d.HasChange("macs") {
		t, err := expandUserQuarantineTargetsMacs2edl(d, v, "macs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["macs"] = t
		}
	}

	return &obj, nil
}
