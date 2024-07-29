// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Conditional advertisement.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterBgpNeighborConditionalAdvertise() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterBgpNeighborConditionalAdvertiseCreate,
		Read:   resourceRouterBgpNeighborConditionalAdvertiseRead,
		Update: resourceRouterBgpNeighborConditionalAdvertiseUpdate,
		Delete: resourceRouterBgpNeighborConditionalAdvertiseDelete,

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
			"neighbor": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"advertise_routemap": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"condition_routemap": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"condition_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRouterBgpNeighborConditionalAdvertiseCreate(d *schema.ResourceData, m interface{}) error {
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
	neighbor := d.Get("neighbor").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["neighbor"] = neighbor

	obj, err := getObjectRouterBgpNeighborConditionalAdvertise(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterBgpNeighborConditionalAdvertise resource while getting object: %v", err)
	}

	_, err = c.CreateRouterBgpNeighborConditionalAdvertise(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating RouterBgpNeighborConditionalAdvertise resource: %v", err)
	}

	d.SetId(getStringKey(d, "advertise_routemap"))

	return resourceRouterBgpNeighborConditionalAdvertiseRead(d, m)
}

func resourceRouterBgpNeighborConditionalAdvertiseUpdate(d *schema.ResourceData, m interface{}) error {
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
	neighbor := d.Get("neighbor").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["neighbor"] = neighbor

	obj, err := getObjectRouterBgpNeighborConditionalAdvertise(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgpNeighborConditionalAdvertise resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterBgpNeighborConditionalAdvertise(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgpNeighborConditionalAdvertise resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "advertise_routemap"))

	return resourceRouterBgpNeighborConditionalAdvertiseRead(d, m)
}

func resourceRouterBgpNeighborConditionalAdvertiseDelete(d *schema.ResourceData, m interface{}) error {
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
	neighbor := d.Get("neighbor").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["neighbor"] = neighbor

	err = c.DeleteRouterBgpNeighborConditionalAdvertise(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting RouterBgpNeighborConditionalAdvertise resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterBgpNeighborConditionalAdvertiseRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	neighbor := d.Get("neighbor").(string)
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
	if neighbor == "" {
		neighbor = importOptionChecking(m.(*FortiClient).Cfg, "neighbor")
		if neighbor == "" {
			return fmt.Errorf("Parameter neighbor is missing")
		}
		if err = d.Set("neighbor", neighbor); err != nil {
			return fmt.Errorf("Error set params neighbor: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["neighbor"] = neighbor

	o, err := c.ReadRouterBgpNeighborConditionalAdvertise(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgpNeighborConditionalAdvertise resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterBgpNeighborConditionalAdvertise(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgpNeighborConditionalAdvertise resource from API: %v", err)
	}
	return nil
}

func flattenRouterBgpNeighborConditionalAdvertiseAdvertiseRoutemap3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborConditionalAdvertiseConditionRoutemap3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborConditionalAdvertiseConditionType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterBgpNeighborConditionalAdvertise(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("advertise_routemap", flattenRouterBgpNeighborConditionalAdvertiseAdvertiseRoutemap3rdl(o["advertise-routemap"], d, "advertise_routemap")); err != nil {
		if vv, ok := fortiAPIPatch(o["advertise-routemap"], "RouterBgpNeighborConditionalAdvertise-AdvertiseRoutemap"); ok {
			if err = d.Set("advertise_routemap", vv); err != nil {
				return fmt.Errorf("Error reading advertise_routemap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading advertise_routemap: %v", err)
		}
	}

	if err = d.Set("condition_routemap", flattenRouterBgpNeighborConditionalAdvertiseConditionRoutemap3rdl(o["condition-routemap"], d, "condition_routemap")); err != nil {
		if vv, ok := fortiAPIPatch(o["condition-routemap"], "RouterBgpNeighborConditionalAdvertise-ConditionRoutemap"); ok {
			if err = d.Set("condition_routemap", vv); err != nil {
				return fmt.Errorf("Error reading condition_routemap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading condition_routemap: %v", err)
		}
	}

	if err = d.Set("condition_type", flattenRouterBgpNeighborConditionalAdvertiseConditionType3rdl(o["condition-type"], d, "condition_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["condition-type"], "RouterBgpNeighborConditionalAdvertise-ConditionType"); ok {
			if err = d.Set("condition_type", vv); err != nil {
				return fmt.Errorf("Error reading condition_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading condition_type: %v", err)
		}
	}

	return nil
}

func flattenRouterBgpNeighborConditionalAdvertiseFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterBgpNeighborConditionalAdvertiseAdvertiseRoutemap3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborConditionalAdvertiseConditionRoutemap3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborConditionalAdvertiseConditionType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterBgpNeighborConditionalAdvertise(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("advertise_routemap"); ok || d.HasChange("advertise_routemap") {
		t, err := expandRouterBgpNeighborConditionalAdvertiseAdvertiseRoutemap3rdl(d, v, "advertise_routemap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["advertise-routemap"] = t
		}
	}

	if v, ok := d.GetOk("condition_routemap"); ok || d.HasChange("condition_routemap") {
		t, err := expandRouterBgpNeighborConditionalAdvertiseConditionRoutemap3rdl(d, v, "condition_routemap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["condition-routemap"] = t
		}
	}

	if v, ok := d.GetOk("condition_type"); ok || d.HasChange("condition_type") {
		t, err := expandRouterBgpNeighborConditionalAdvertiseConditionType3rdl(d, v, "condition_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["condition-type"] = t
		}
	}

	return &obj, nil
}
