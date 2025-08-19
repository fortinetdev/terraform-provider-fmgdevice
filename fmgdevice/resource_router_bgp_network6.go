// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: BGP IPv6 network table.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterBgpNetwork6() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterBgpNetwork6Create,
		Read:   resourceRouterBgpNetwork6Read,
		Update: resourceRouterBgpNetwork6Update,
		Delete: resourceRouterBgpNetwork6Delete,

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
			},
			"network_import_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prefix6": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceRouterBgpNetwork6Create(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterBgpNetwork6(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterBgpNetwork6 resource while getting object: %v", err)
	}

	_, err = c.CreateRouterBgpNetwork6(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating RouterBgpNetwork6 resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterBgpNetwork6Read(d, m)
}

func resourceRouterBgpNetwork6Update(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterBgpNetwork6(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgpNetwork6 resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterBgpNetwork6(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgpNetwork6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterBgpNetwork6Read(d, m)
}

func resourceRouterBgpNetwork6Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterBgpNetwork6(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterBgpNetwork6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterBgpNetwork6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterBgpNetwork6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgpNetwork6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterBgpNetwork6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgpNetwork6 resource from API: %v", err)
	}
	return nil
}

func flattenRouterBgpNetwork6Backdoor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNetwork6Id2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNetwork6NetworkImportCheck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNetwork6Prefix62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpNetwork6RouteMap2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectRouterBgpNetwork6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("backdoor", flattenRouterBgpNetwork6Backdoor2edl(o["backdoor"], d, "backdoor")); err != nil {
		if vv, ok := fortiAPIPatch(o["backdoor"], "RouterBgpNetwork6-Backdoor"); ok {
			if err = d.Set("backdoor", vv); err != nil {
				return fmt.Errorf("Error reading backdoor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading backdoor: %v", err)
		}
	}

	if err = d.Set("fosid", flattenRouterBgpNetwork6Id2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "RouterBgpNetwork6-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("network_import_check", flattenRouterBgpNetwork6NetworkImportCheck2edl(o["network-import-check"], d, "network_import_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["network-import-check"], "RouterBgpNetwork6-NetworkImportCheck"); ok {
			if err = d.Set("network_import_check", vv); err != nil {
				return fmt.Errorf("Error reading network_import_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading network_import_check: %v", err)
		}
	}

	if err = d.Set("prefix6", flattenRouterBgpNetwork6Prefix62edl(o["prefix6"], d, "prefix6")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix6"], "RouterBgpNetwork6-Prefix6"); ok {
			if err = d.Set("prefix6", vv); err != nil {
				return fmt.Errorf("Error reading prefix6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix6: %v", err)
		}
	}

	if err = d.Set("route_map", flattenRouterBgpNetwork6RouteMap2edl(o["route-map"], d, "route_map")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-map"], "RouterBgpNetwork6-RouteMap"); ok {
			if err = d.Set("route_map", vv); err != nil {
				return fmt.Errorf("Error reading route_map: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_map: %v", err)
		}
	}

	return nil
}

func flattenRouterBgpNetwork6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterBgpNetwork6Backdoor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetwork6Id2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetwork6NetworkImportCheck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetwork6Prefix62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetwork6RouteMap2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectRouterBgpNetwork6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("backdoor"); ok || d.HasChange("backdoor") {
		t, err := expandRouterBgpNetwork6Backdoor2edl(d, v, "backdoor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backdoor"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandRouterBgpNetwork6Id2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("network_import_check"); ok || d.HasChange("network_import_check") {
		t, err := expandRouterBgpNetwork6NetworkImportCheck2edl(d, v, "network_import_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-import-check"] = t
		}
	}

	if v, ok := d.GetOk("prefix6"); ok || d.HasChange("prefix6") {
		t, err := expandRouterBgpNetwork6Prefix62edl(d, v, "prefix6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix6"] = t
		}
	}

	if v, ok := d.GetOk("route_map"); ok || d.HasChange("route_map") {
		t, err := expandRouterBgpNetwork6RouteMap2edl(d, v, "route_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map"] = t
		}
	}

	return &obj, nil
}
