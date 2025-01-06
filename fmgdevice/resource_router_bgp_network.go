// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: BGP network table.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterBgpNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterBgpNetworkCreate,
		Read:   resourceRouterBgpNetworkRead,
		Update: resourceRouterBgpNetworkUpdate,
		Delete: resourceRouterBgpNetworkDelete,

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
			"backdoor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"network_import_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prefix": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"prefix_name": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"route_map": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRouterBgpNetworkCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterBgpNetwork(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterBgpNetwork resource while getting object: %v", err)
	}

	v, err := c.CreateRouterBgpNetwork(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating RouterBgpNetwork resource: %v", err)
	}

	if v != nil && v["id"] != nil {
		if vidn, ok := v["id"].(float64); ok {
			d.SetId(strconv.Itoa(int(vidn)))
			return resourceRouterBgpNetworkRead(d, m)
		} else {
			return fmt.Errorf("Error creating RouterBgpNetwork resource: %v", err)
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterBgpNetworkRead(d, m)
}

func resourceRouterBgpNetworkUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterBgpNetwork(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgpNetwork resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterBgpNetwork(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgpNetwork resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterBgpNetworkRead(d, m)
}

func resourceRouterBgpNetworkDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterBgpNetwork(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting RouterBgpNetwork resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterBgpNetworkRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterBgpNetwork(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgpNetwork resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterBgpNetwork(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgpNetwork resource from API: %v", err)
	}
	return nil
}

func flattenRouterBgpNetworkBackdoor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNetworkId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNetworkNetworkImportCheck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNetworkPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNetworkPrefixName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpNetworkRouteMap2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectRouterBgpNetwork(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("backdoor", flattenRouterBgpNetworkBackdoor2edl(o["backdoor"], d, "backdoor")); err != nil {
		if vv, ok := fortiAPIPatch(o["backdoor"], "RouterBgpNetwork-Backdoor"); ok {
			if err = d.Set("backdoor", vv); err != nil {
				return fmt.Errorf("Error reading backdoor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading backdoor: %v", err)
		}
	}

	if err = d.Set("fosid", flattenRouterBgpNetworkId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "RouterBgpNetwork-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("network_import_check", flattenRouterBgpNetworkNetworkImportCheck2edl(o["network-import-check"], d, "network_import_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["network-import-check"], "RouterBgpNetwork-NetworkImportCheck"); ok {
			if err = d.Set("network_import_check", vv); err != nil {
				return fmt.Errorf("Error reading network_import_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading network_import_check: %v", err)
		}
	}

	if err = d.Set("prefix", flattenRouterBgpNetworkPrefix2edl(o["prefix"], d, "prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix"], "RouterBgpNetwork-Prefix"); ok {
			if err = d.Set("prefix", vv); err != nil {
				return fmt.Errorf("Error reading prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix: %v", err)
		}
	}

	if err = d.Set("prefix_name", flattenRouterBgpNetworkPrefixName2edl(o["prefix-name"], d, "prefix_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix-name"], "RouterBgpNetwork-PrefixName"); ok {
			if err = d.Set("prefix_name", vv); err != nil {
				return fmt.Errorf("Error reading prefix_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix_name: %v", err)
		}
	}

	if err = d.Set("route_map", flattenRouterBgpNetworkRouteMap2edl(o["route-map"], d, "route_map")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-map"], "RouterBgpNetwork-RouteMap"); ok {
			if err = d.Set("route_map", vv); err != nil {
				return fmt.Errorf("Error reading route_map: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_map: %v", err)
		}
	}

	return nil
}

func flattenRouterBgpNetworkFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterBgpNetworkBackdoor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetworkId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetworkNetworkImportCheck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetworkPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandRouterBgpNetworkPrefixName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpNetworkRouteMap2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectRouterBgpNetwork(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("backdoor"); ok || d.HasChange("backdoor") {
		t, err := expandRouterBgpNetworkBackdoor2edl(d, v, "backdoor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backdoor"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandRouterBgpNetworkId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("network_import_check"); ok || d.HasChange("network_import_check") {
		t, err := expandRouterBgpNetworkNetworkImportCheck2edl(d, v, "network_import_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-import-check"] = t
		}
	}

	if v, ok := d.GetOk("prefix"); ok || d.HasChange("prefix") {
		t, err := expandRouterBgpNetworkPrefix2edl(d, v, "prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix"] = t
		}
	}

	if v, ok := d.GetOk("prefix_name"); ok || d.HasChange("prefix_name") {
		t, err := expandRouterBgpNetworkPrefixName2edl(d, v, "prefix_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-name"] = t
		}
	}

	if v, ok := d.GetOk("route_map"); ok || d.HasChange("route_map") {
		t, err := expandRouterBgpNetworkRouteMap2edl(d, v, "route_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map"] = t
		}
	}

	return &obj, nil
}
