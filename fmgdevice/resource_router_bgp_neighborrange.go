// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: BGP neighbor range table.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterBgpNeighborRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterBgpNeighborRangeCreate,
		Read:   resourceRouterBgpNeighborRangeRead,
		Update: resourceRouterBgpNeighborRangeUpdate,
		Delete: resourceRouterBgpNeighborRangeDelete,

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
			"prefix": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRouterBgpNeighborRangeCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterBgpNeighborRange(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterBgpNeighborRange resource while getting object: %v", err)
	}

	_, err = c.CreateRouterBgpNeighborRange(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating RouterBgpNeighborRange resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterBgpNeighborRangeRead(d, m)
}

func resourceRouterBgpNeighborRangeUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterBgpNeighborRange(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgpNeighborRange resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterBgpNeighborRange(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgpNeighborRange resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterBgpNeighborRangeRead(d, m)
}

func resourceRouterBgpNeighborRangeDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterBgpNeighborRange(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterBgpNeighborRange resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterBgpNeighborRangeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterBgpNeighborRange(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgpNeighborRange resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterBgpNeighborRange(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgpNeighborRange resource from API: %v", err)
	}
	return nil
}

func flattenRouterBgpNeighborRangeId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRangeMaxNeighborNum2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNeighborRangeNeighborGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNeighborRangePrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectRouterBgpNeighborRange(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenRouterBgpNeighborRangeId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "RouterBgpNeighborRange-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("max_neighbor_num", flattenRouterBgpNeighborRangeMaxNeighborNum2edl(o["max-neighbor-num"], d, "max_neighbor_num")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-neighbor-num"], "RouterBgpNeighborRange-MaxNeighborNum"); ok {
			if err = d.Set("max_neighbor_num", vv); err != nil {
				return fmt.Errorf("Error reading max_neighbor_num: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_neighbor_num: %v", err)
		}
	}

	if err = d.Set("neighbor_group", flattenRouterBgpNeighborRangeNeighborGroup2edl(o["neighbor-group"], d, "neighbor_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["neighbor-group"], "RouterBgpNeighborRange-NeighborGroup"); ok {
			if err = d.Set("neighbor_group", vv); err != nil {
				return fmt.Errorf("Error reading neighbor_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading neighbor_group: %v", err)
		}
	}

	if err = d.Set("prefix", flattenRouterBgpNeighborRangePrefix2edl(o["prefix"], d, "prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix"], "RouterBgpNeighborRange-Prefix"); ok {
			if err = d.Set("prefix", vv); err != nil {
				return fmt.Errorf("Error reading prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix: %v", err)
		}
	}

	return nil
}

func flattenRouterBgpNeighborRangeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterBgpNeighborRangeId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRangeMaxNeighborNum2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRangeNeighborGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNeighborRangePrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func getObjectRouterBgpNeighborRange(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandRouterBgpNeighborRangeId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("max_neighbor_num"); ok || d.HasChange("max_neighbor_num") {
		t, err := expandRouterBgpNeighborRangeMaxNeighborNum2edl(d, v, "max_neighbor_num")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-neighbor-num"] = t
		}
	}

	if v, ok := d.GetOk("neighbor_group"); ok || d.HasChange("neighbor_group") {
		t, err := expandRouterBgpNeighborRangeNeighborGroup2edl(d, v, "neighbor_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-group"] = t
		}
	}

	if v, ok := d.GetOk("prefix"); ok || d.HasChange("prefix") {
		t, err := expandRouterBgpNeighborRangePrefix2edl(d, v, "prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix"] = t
		}
	}

	return &obj, nil
}
