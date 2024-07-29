// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Automation stitches.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAutomationStitch() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutomationStitchCreate,
		Read:   resourceSystemAutomationStitchRead,
		Update: resourceSystemAutomationStitchUpdate,
		Delete: resourceSystemAutomationStitchDelete,

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
			"action": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"actions": &schema.Schema{
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
						"delay": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"required": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"destination": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trigger": &schema.Schema{
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

func resourceSystemAutomationStitchCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemAutomationStitch(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationStitch resource while getting object: %v", err)
	}

	_, err = c.CreateSystemAutomationStitch(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationStitch resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemAutomationStitchRead(d, m)
}

func resourceSystemAutomationStitchUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemAutomationStitch(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationStitch resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAutomationStitch(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationStitch resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemAutomationStitchRead(d, m)
}

func resourceSystemAutomationStitchDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemAutomationStitch(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAutomationStitch resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutomationStitchRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemAutomationStitch(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationStitch resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutomationStitch(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationStitch resource from API: %v", err)
	}
	return nil
}

func flattenSystemAutomationStitchAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAutomationStitchActions(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemAutomationStitchActionsAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "SystemAutomationStitch-Actions-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "delay"
		if _, ok := i["delay"]; ok {
			v := flattenSystemAutomationStitchActionsDelay(i["delay"], d, pre_append)
			tmp["delay"] = fortiAPISubPartPatch(v, "SystemAutomationStitch-Actions-Delay")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemAutomationStitchActionsId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemAutomationStitch-Actions-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "required"
		if _, ok := i["required"]; ok {
			v := flattenSystemAutomationStitchActionsRequired(i["required"], d, pre_append)
			tmp["required"] = fortiAPISubPartPatch(v, "SystemAutomationStitch-Actions-Required")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemAutomationStitchActionsAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAutomationStitchActionsDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationStitchActionsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationStitchActionsRequired(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationStitchDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationStitchDestination(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAutomationStitchName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationStitchStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationStitchTrigger(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemAutomationStitch(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("action", flattenSystemAutomationStitchAction(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "SystemAutomationStitch-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("actions", flattenSystemAutomationStitchActions(o["actions"], d, "actions")); err != nil {
			if vv, ok := fortiAPIPatch(o["actions"], "SystemAutomationStitch-Actions"); ok {
				if err = d.Set("actions", vv); err != nil {
					return fmt.Errorf("Error reading actions: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading actions: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("actions"); ok {
			if err = d.Set("actions", flattenSystemAutomationStitchActions(o["actions"], d, "actions")); err != nil {
				if vv, ok := fortiAPIPatch(o["actions"], "SystemAutomationStitch-Actions"); ok {
					if err = d.Set("actions", vv); err != nil {
						return fmt.Errorf("Error reading actions: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading actions: %v", err)
				}
			}
		}
	}

	if err = d.Set("description", flattenSystemAutomationStitchDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SystemAutomationStitch-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("destination", flattenSystemAutomationStitchDestination(o["destination"], d, "destination")); err != nil {
		if vv, ok := fortiAPIPatch(o["destination"], "SystemAutomationStitch-Destination"); ok {
			if err = d.Set("destination", vv); err != nil {
				return fmt.Errorf("Error reading destination: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading destination: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemAutomationStitchName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemAutomationStitch-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemAutomationStitchStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemAutomationStitch-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("trigger", flattenSystemAutomationStitchTrigger(o["trigger"], d, "trigger")); err != nil {
		if vv, ok := fortiAPIPatch(o["trigger"], "SystemAutomationStitch-Trigger"); ok {
			if err = d.Set("trigger", vv); err != nil {
				return fmt.Errorf("Error reading trigger: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trigger: %v", err)
		}
	}

	return nil
}

func flattenSystemAutomationStitchFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAutomationStitchAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAutomationStitchActions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandSystemAutomationStitchActionsAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "delay"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["delay"], _ = expandSystemAutomationStitchActionsDelay(d, i["delay"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemAutomationStitchActionsId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "required"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["required"], _ = expandSystemAutomationStitchActionsRequired(d, i["required"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemAutomationStitchActionsAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAutomationStitchActionsDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationStitchActionsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationStitchActionsRequired(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationStitchDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationStitchDestination(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAutomationStitchName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationStitchStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationStitchTrigger(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemAutomationStitch(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandSystemAutomationStitchAction(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("actions"); ok || d.HasChange("actions") {
		t, err := expandSystemAutomationStitchActions(d, v, "actions")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["actions"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSystemAutomationStitchDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("destination"); ok || d.HasChange("destination") {
		t, err := expandSystemAutomationStitchDestination(d, v, "destination")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["destination"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemAutomationStitchName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemAutomationStitchStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("trigger"); ok || d.HasChange("trigger") {
		t, err := expandSystemAutomationStitchTrigger(d, v, "trigger")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger"] = t
		}
	}

	return &obj, nil
}
