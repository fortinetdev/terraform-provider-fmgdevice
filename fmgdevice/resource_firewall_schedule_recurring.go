// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Recurring schedule configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallScheduleRecurring() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallScheduleRecurringCreate,
		Read:   resourceFirewallScheduleRecurringRead,
		Update: resourceFirewallScheduleRecurringUpdate,
		Delete: resourceFirewallScheduleRecurringDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

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
			"color": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"day": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"end": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fabric_object": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"global_object": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"start": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallScheduleRecurringCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectFirewallScheduleRecurring(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallScheduleRecurring resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallScheduleRecurring(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallScheduleRecurring(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallScheduleRecurring resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallScheduleRecurring(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallScheduleRecurring resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallScheduleRecurringRead(d, m)
}

func resourceFirewallScheduleRecurringUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectFirewallScheduleRecurring(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallScheduleRecurring resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallScheduleRecurring(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallScheduleRecurring resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallScheduleRecurringRead(d, m)
}

func resourceFirewallScheduleRecurringDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	wsParams["adom"] = adomv

	err = c.DeleteFirewallScheduleRecurring(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallScheduleRecurring resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallScheduleRecurringRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallScheduleRecurring(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallScheduleRecurring resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallScheduleRecurring(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallScheduleRecurring resource from API: %v", err)
	}
	return nil
}

func flattenFirewallScheduleRecurringColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallScheduleRecurringDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallScheduleRecurringEnd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallScheduleRecurringFabricObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallScheduleRecurringGlobalObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallScheduleRecurringName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallScheduleRecurringStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallScheduleRecurringUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallScheduleRecurring(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("color", flattenFirewallScheduleRecurringColor(o["color"], d, "color")); err != nil {
		if vv, ok := fortiAPIPatch(o["color"], "FirewallScheduleRecurring-Color"); ok {
			if err = d.Set("color", vv); err != nil {
				return fmt.Errorf("Error reading color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("day", flattenFirewallScheduleRecurringDay(o["day"], d, "day")); err != nil {
		if vv, ok := fortiAPIPatch(o["day"], "FirewallScheduleRecurring-Day"); ok {
			if err = d.Set("day", vv); err != nil {
				return fmt.Errorf("Error reading day: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading day: %v", err)
		}
	}

	if err = d.Set("end", flattenFirewallScheduleRecurringEnd(o["end"], d, "end")); err != nil {
		if vv, ok := fortiAPIPatch(o["end"], "FirewallScheduleRecurring-End"); ok {
			if err = d.Set("end", vv); err != nil {
				return fmt.Errorf("Error reading end: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end: %v", err)
		}
	}

	if err = d.Set("fabric_object", flattenFirewallScheduleRecurringFabricObject(o["fabric-object"], d, "fabric_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-object"], "FirewallScheduleRecurring-FabricObject"); ok {
			if err = d.Set("fabric_object", vv); err != nil {
				return fmt.Errorf("Error reading fabric_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_object: %v", err)
		}
	}

	if err = d.Set("global_object", flattenFirewallScheduleRecurringGlobalObject(o["global-object"], d, "global_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["global-object"], "FirewallScheduleRecurring-GlobalObject"); ok {
			if err = d.Set("global_object", vv); err != nil {
				return fmt.Errorf("Error reading global_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading global_object: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallScheduleRecurringName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallScheduleRecurring-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("start", flattenFirewallScheduleRecurringStart(o["start"], d, "start")); err != nil {
		if vv, ok := fortiAPIPatch(o["start"], "FirewallScheduleRecurring-Start"); ok {
			if err = d.Set("start", vv); err != nil {
				return fmt.Errorf("Error reading start: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallScheduleRecurringUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "FirewallScheduleRecurring-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	return nil
}

func flattenFirewallScheduleRecurringFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallScheduleRecurringColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallScheduleRecurringDay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallScheduleRecurringEnd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallScheduleRecurringFabricObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallScheduleRecurringGlobalObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallScheduleRecurringName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallScheduleRecurringStart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallScheduleRecurringUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallScheduleRecurring(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("color"); ok || d.HasChange("color") {
		t, err := expandFirewallScheduleRecurringColor(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("day"); ok || d.HasChange("day") {
		t, err := expandFirewallScheduleRecurringDay(d, v, "day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["day"] = t
		}
	}

	if v, ok := d.GetOk("end"); ok || d.HasChange("end") {
		t, err := expandFirewallScheduleRecurringEnd(d, v, "end")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object"); ok || d.HasChange("fabric_object") {
		t, err := expandFirewallScheduleRecurringFabricObject(d, v, "fabric_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object"] = t
		}
	}

	if v, ok := d.GetOk("global_object"); ok || d.HasChange("global_object") {
		t, err := expandFirewallScheduleRecurringGlobalObject(d, v, "global_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["global-object"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallScheduleRecurringName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("start"); ok || d.HasChange("start") {
		t, err := expandFirewallScheduleRecurringStart(d, v, "start")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandFirewallScheduleRecurringUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	return &obj, nil
}
