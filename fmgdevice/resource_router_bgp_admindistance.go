// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Administrative distance modifications.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterBgpAdminDistance() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterBgpAdminDistanceCreate,
		Read:   resourceRouterBgpAdminDistanceRead,
		Update: resourceRouterBgpAdminDistanceUpdate,
		Delete: resourceRouterBgpAdminDistanceDelete,

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
			"distance": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"neighbour_prefix": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"route_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRouterBgpAdminDistanceCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterBgpAdminDistance(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterBgpAdminDistance resource while getting object: %v", err)
	}

	v, err := c.CreateRouterBgpAdminDistance(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating RouterBgpAdminDistance resource: %v", err)
	}

	if v != nil && v["id"] != nil {
		if vidn, ok := v["id"].(float64); ok {
			d.SetId(strconv.Itoa(int(vidn)))
			return resourceRouterBgpAdminDistanceRead(d, m)
		} else {
			return fmt.Errorf("Error creating RouterBgpAdminDistance resource: %v", err)
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterBgpAdminDistanceRead(d, m)
}

func resourceRouterBgpAdminDistanceUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterBgpAdminDistance(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgpAdminDistance resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterBgpAdminDistance(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgpAdminDistance resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterBgpAdminDistanceRead(d, m)
}

func resourceRouterBgpAdminDistanceDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterBgpAdminDistance(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterBgpAdminDistance resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterBgpAdminDistanceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterBgpAdminDistance(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgpAdminDistance resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterBgpAdminDistance(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgpAdminDistance resource from API: %v", err)
	}
	return nil
}

func flattenRouterBgpAdminDistanceDistance2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpAdminDistanceId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpAdminDistanceNeighbourPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpAdminDistanceRouteList2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectRouterBgpAdminDistance(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("distance", flattenRouterBgpAdminDistanceDistance2edl(o["distance"], d, "distance")); err != nil {
		if vv, ok := fortiAPIPatch(o["distance"], "RouterBgpAdminDistance-Distance"); ok {
			if err = d.Set("distance", vv); err != nil {
				return fmt.Errorf("Error reading distance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	if err = d.Set("fosid", flattenRouterBgpAdminDistanceId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "RouterBgpAdminDistance-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("neighbour_prefix", flattenRouterBgpAdminDistanceNeighbourPrefix2edl(o["neighbour-prefix"], d, "neighbour_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["neighbour-prefix"], "RouterBgpAdminDistance-NeighbourPrefix"); ok {
			if err = d.Set("neighbour_prefix", vv); err != nil {
				return fmt.Errorf("Error reading neighbour_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading neighbour_prefix: %v", err)
		}
	}

	if err = d.Set("route_list", flattenRouterBgpAdminDistanceRouteList2edl(o["route-list"], d, "route_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-list"], "RouterBgpAdminDistance-RouteList"); ok {
			if err = d.Set("route_list", vv); err != nil {
				return fmt.Errorf("Error reading route_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_list: %v", err)
		}
	}

	return nil
}

func flattenRouterBgpAdminDistanceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterBgpAdminDistanceDistance2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdminDistanceId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdminDistanceNeighbourPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandRouterBgpAdminDistanceRouteList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectRouterBgpAdminDistance(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("distance"); ok || d.HasChange("distance") {
		t, err := expandRouterBgpAdminDistanceDistance2edl(d, v, "distance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandRouterBgpAdminDistanceId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("neighbour_prefix"); ok || d.HasChange("neighbour_prefix") {
		t, err := expandRouterBgpAdminDistanceNeighbourPrefix2edl(d, v, "neighbour_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbour-prefix"] = t
		}
	}

	if v, ok := d.GetOk("route_list"); ok || d.HasChange("route_list") {
		t, err := expandRouterBgpAdminDistanceRouteList2edl(d, v, "route_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-list"] = t
		}
	}

	return &obj, nil
}
