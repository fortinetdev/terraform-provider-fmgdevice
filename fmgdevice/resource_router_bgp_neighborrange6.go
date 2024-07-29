// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: BGP IPv6 neighbor range table.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterBgpNeighborRange6() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterBgpNeighborRange6Create,
		Read:   resourceRouterBgpNeighborRange6Read,
		Update: resourceRouterBgpNeighborRange6Update,
		Delete: resourceRouterBgpNeighborRange6Delete,

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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"max_neighbor_num": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"neighbor_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"prefix6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRouterBgpNeighborRange6Create(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterBgpNeighborRange6(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterBgpNeighborRange6 resource while getting object: %v", err)
	}

	_, err = c.CreateRouterBgpNeighborRange6(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating RouterBgpNeighborRange6 resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterBgpNeighborRange6Read(d, m)
}

func resourceRouterBgpNeighborRange6Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterBgpNeighborRange6(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgpNeighborRange6 resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterBgpNeighborRange6(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgpNeighborRange6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterBgpNeighborRange6Read(d, m)
}

func resourceRouterBgpNeighborRange6Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterBgpNeighborRange6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting RouterBgpNeighborRange6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterBgpNeighborRange6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterBgpNeighborRange6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgpNeighborRange6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterBgpNeighborRange6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgpNeighborRange6 resource from API: %v", err)
	}
	return nil
}

func flattenRouterBgpNeighborRange6Id2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRange6MaxNeighborNum2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRange6NeighborGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRange6Prefix62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterBgpNeighborRange6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenRouterBgpNeighborRange6Id2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "RouterBgpNeighborRange6-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("max_neighbor_num", flattenRouterBgpNeighborRange6MaxNeighborNum2edl(o["max-neighbor-num"], d, "max_neighbor_num")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-neighbor-num"], "RouterBgpNeighborRange6-MaxNeighborNum"); ok {
			if err = d.Set("max_neighbor_num", vv); err != nil {
				return fmt.Errorf("Error reading max_neighbor_num: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_neighbor_num: %v", err)
		}
	}

	if err = d.Set("neighbor_group", flattenRouterBgpNeighborRange6NeighborGroup2edl(o["neighbor-group"], d, "neighbor_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["neighbor-group"], "RouterBgpNeighborRange6-NeighborGroup"); ok {
			if err = d.Set("neighbor_group", vv); err != nil {
				return fmt.Errorf("Error reading neighbor_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading neighbor_group: %v", err)
		}
	}

	if err = d.Set("prefix6", flattenRouterBgpNeighborRange6Prefix62edl(o["prefix6"], d, "prefix6")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix6"], "RouterBgpNeighborRange6-Prefix6"); ok {
			if err = d.Set("prefix6", vv); err != nil {
				return fmt.Errorf("Error reading prefix6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix6: %v", err)
		}
	}

	return nil
}

func flattenRouterBgpNeighborRange6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterBgpNeighborRange6Id2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRange6MaxNeighborNum2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRange6NeighborGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRange6Prefix62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterBgpNeighborRange6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandRouterBgpNeighborRange6Id2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("max_neighbor_num"); ok || d.HasChange("max_neighbor_num") {
		t, err := expandRouterBgpNeighborRange6MaxNeighborNum2edl(d, v, "max_neighbor_num")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-neighbor-num"] = t
		}
	}

	if v, ok := d.GetOk("neighbor_group"); ok || d.HasChange("neighbor_group") {
		t, err := expandRouterBgpNeighborRange6NeighborGroup2edl(d, v, "neighbor_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-group"] = t
		}
	}

	if v, ok := d.GetOk("prefix6"); ok || d.HasChange("prefix6") {
		t, err := expandRouterBgpNeighborRange6Prefix62edl(d, v, "prefix6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix6"] = t
		}
	}

	return &obj, nil
}
