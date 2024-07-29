// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch QoS 802.1p.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerQosDot1PMap() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerQosDot1PMapCreate,
		Read:   resourceSwitchControllerQosDot1PMapRead,
		Update: resourceSwitchControllerQosDot1PMapUpdate,
		Delete: resourceSwitchControllerQosDot1PMapDelete,

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
			"egress_pri_tagging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"priority_0": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_7": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerQosDot1PMapCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerQosDot1PMap(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerQosDot1PMap resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerQosDot1PMap(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerQosDot1PMap resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerQosDot1PMapRead(d, m)
}

func resourceSwitchControllerQosDot1PMapUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerQosDot1PMap(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerQosDot1PMap resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerQosDot1PMap(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerQosDot1PMap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerQosDot1PMapRead(d, m)
}

func resourceSwitchControllerQosDot1PMapDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerQosDot1PMap(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerQosDot1PMap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerQosDot1PMapRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerQosDot1PMap(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerQosDot1PMap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerQosDot1PMap(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerQosDot1PMap resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerQosDot1PMapDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosDot1PMapEgressPriTagging(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosDot1PMapName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosDot1PMapPriority0(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosDot1PMapPriority1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosDot1PMapPriority2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosDot1PMapPriority3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosDot1PMapPriority4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosDot1PMapPriority5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosDot1PMapPriority6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosDot1PMapPriority7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerQosDot1PMap(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("description", flattenSwitchControllerQosDot1PMapDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SwitchControllerQosDot1PMap-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("egress_pri_tagging", flattenSwitchControllerQosDot1PMapEgressPriTagging(o["egress-pri-tagging"], d, "egress_pri_tagging")); err != nil {
		if vv, ok := fortiAPIPatch(o["egress-pri-tagging"], "SwitchControllerQosDot1PMap-EgressPriTagging"); ok {
			if err = d.Set("egress_pri_tagging", vv); err != nil {
				return fmt.Errorf("Error reading egress_pri_tagging: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading egress_pri_tagging: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchControllerQosDot1PMapName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerQosDot1PMap-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("priority_0", flattenSwitchControllerQosDot1PMapPriority0(o["priority-0"], d, "priority_0")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-0"], "SwitchControllerQosDot1PMap-Priority0"); ok {
			if err = d.Set("priority_0", vv); err != nil {
				return fmt.Errorf("Error reading priority_0: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_0: %v", err)
		}
	}

	if err = d.Set("priority_1", flattenSwitchControllerQosDot1PMapPriority1(o["priority-1"], d, "priority_1")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-1"], "SwitchControllerQosDot1PMap-Priority1"); ok {
			if err = d.Set("priority_1", vv); err != nil {
				return fmt.Errorf("Error reading priority_1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_1: %v", err)
		}
	}

	if err = d.Set("priority_2", flattenSwitchControllerQosDot1PMapPriority2(o["priority-2"], d, "priority_2")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-2"], "SwitchControllerQosDot1PMap-Priority2"); ok {
			if err = d.Set("priority_2", vv); err != nil {
				return fmt.Errorf("Error reading priority_2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_2: %v", err)
		}
	}

	if err = d.Set("priority_3", flattenSwitchControllerQosDot1PMapPriority3(o["priority-3"], d, "priority_3")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-3"], "SwitchControllerQosDot1PMap-Priority3"); ok {
			if err = d.Set("priority_3", vv); err != nil {
				return fmt.Errorf("Error reading priority_3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_3: %v", err)
		}
	}

	if err = d.Set("priority_4", flattenSwitchControllerQosDot1PMapPriority4(o["priority-4"], d, "priority_4")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-4"], "SwitchControllerQosDot1PMap-Priority4"); ok {
			if err = d.Set("priority_4", vv); err != nil {
				return fmt.Errorf("Error reading priority_4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_4: %v", err)
		}
	}

	if err = d.Set("priority_5", flattenSwitchControllerQosDot1PMapPriority5(o["priority-5"], d, "priority_5")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-5"], "SwitchControllerQosDot1PMap-Priority5"); ok {
			if err = d.Set("priority_5", vv); err != nil {
				return fmt.Errorf("Error reading priority_5: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_5: %v", err)
		}
	}

	if err = d.Set("priority_6", flattenSwitchControllerQosDot1PMapPriority6(o["priority-6"], d, "priority_6")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-6"], "SwitchControllerQosDot1PMap-Priority6"); ok {
			if err = d.Set("priority_6", vv); err != nil {
				return fmt.Errorf("Error reading priority_6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_6: %v", err)
		}
	}

	if err = d.Set("priority_7", flattenSwitchControllerQosDot1PMapPriority7(o["priority-7"], d, "priority_7")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-7"], "SwitchControllerQosDot1PMap-Priority7"); ok {
			if err = d.Set("priority_7", vv); err != nil {
				return fmt.Errorf("Error reading priority_7: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_7: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerQosDot1PMapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerQosDot1PMapDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosDot1PMapEgressPriTagging(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosDot1PMapName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosDot1PMapPriority0(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosDot1PMapPriority1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosDot1PMapPriority2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosDot1PMapPriority3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosDot1PMapPriority4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosDot1PMapPriority5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosDot1PMapPriority6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosDot1PMapPriority7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerQosDot1PMap(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSwitchControllerQosDot1PMapDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("egress_pri_tagging"); ok || d.HasChange("egress_pri_tagging") {
		t, err := expandSwitchControllerQosDot1PMapEgressPriTagging(d, v, "egress_pri_tagging")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["egress-pri-tagging"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerQosDot1PMapName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("priority_0"); ok || d.HasChange("priority_0") {
		t, err := expandSwitchControllerQosDot1PMapPriority0(d, v, "priority_0")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-0"] = t
		}
	}

	if v, ok := d.GetOk("priority_1"); ok || d.HasChange("priority_1") {
		t, err := expandSwitchControllerQosDot1PMapPriority1(d, v, "priority_1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-1"] = t
		}
	}

	if v, ok := d.GetOk("priority_2"); ok || d.HasChange("priority_2") {
		t, err := expandSwitchControllerQosDot1PMapPriority2(d, v, "priority_2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-2"] = t
		}
	}

	if v, ok := d.GetOk("priority_3"); ok || d.HasChange("priority_3") {
		t, err := expandSwitchControllerQosDot1PMapPriority3(d, v, "priority_3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-3"] = t
		}
	}

	if v, ok := d.GetOk("priority_4"); ok || d.HasChange("priority_4") {
		t, err := expandSwitchControllerQosDot1PMapPriority4(d, v, "priority_4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-4"] = t
		}
	}

	if v, ok := d.GetOk("priority_5"); ok || d.HasChange("priority_5") {
		t, err := expandSwitchControllerQosDot1PMapPriority5(d, v, "priority_5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-5"] = t
		}
	}

	if v, ok := d.GetOk("priority_6"); ok || d.HasChange("priority_6") {
		t, err := expandSwitchControllerQosDot1PMapPriority6(d, v, "priority_6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-6"] = t
		}
	}

	if v, ok := d.GetOk("priority_7"); ok || d.HasChange("priority_7") {
		t, err := expandSwitchControllerQosDot1PMapPriority7(d, v, "priority_7")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-7"] = t
		}
	}

	return &obj, nil
}
