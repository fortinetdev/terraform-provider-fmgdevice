// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure names for shaping classes.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallTrafficClass() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallTrafficClassCreate,
		Read:   resourceFirewallTrafficClassRead,
		Update: resourceFirewallTrafficClassUpdate,
		Delete: resourceFirewallTrafficClassDelete,

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
			"class_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"class_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceFirewallTrafficClassCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallTrafficClass(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallTrafficClass resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("class_id")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallTrafficClass(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallTrafficClass(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallTrafficClass resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallTrafficClass(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallTrafficClass resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "class_id")))

	return resourceFirewallTrafficClassRead(d, m)
}

func resourceFirewallTrafficClassUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallTrafficClass(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallTrafficClass resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallTrafficClass(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallTrafficClass resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "class_id")))

	return resourceFirewallTrafficClassRead(d, m)
}

func resourceFirewallTrafficClassDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallTrafficClass(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallTrafficClass resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallTrafficClassRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallTrafficClass(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallTrafficClass resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallTrafficClass(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallTrafficClass resource from API: %v", err)
	}
	return nil
}

func flattenFirewallTrafficClassClassId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallTrafficClassClassName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallTrafficClass(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("class_id", flattenFirewallTrafficClassClassId(o["class-id"], d, "class_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["class-id"], "FirewallTrafficClass-ClassId"); ok {
			if err = d.Set("class_id", vv); err != nil {
				return fmt.Errorf("Error reading class_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading class_id: %v", err)
		}
	}

	if err = d.Set("class_name", flattenFirewallTrafficClassClassName(o["class-name"], d, "class_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["class-name"], "FirewallTrafficClass-ClassName"); ok {
			if err = d.Set("class_name", vv); err != nil {
				return fmt.Errorf("Error reading class_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading class_name: %v", err)
		}
	}

	return nil
}

func flattenFirewallTrafficClassFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallTrafficClassClassId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallTrafficClassClassName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallTrafficClass(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("class_id"); ok || d.HasChange("class_id") {
		t, err := expandFirewallTrafficClassClassId(d, v, "class_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["class-id"] = t
		}
	}

	if v, ok := d.GetOk("class_name"); ok || d.HasChange("class_name") {
		t, err := expandFirewallTrafficClassClassName(d, v, "class_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["class-name"] = t
		}
	}

	return &obj, nil
}
