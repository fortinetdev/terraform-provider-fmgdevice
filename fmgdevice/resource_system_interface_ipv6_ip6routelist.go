// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Advertised route list.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemInterfaceIpv6Ip6RouteList() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemInterfaceIpv6Ip6RouteListCreate,
		Read:   resourceSystemInterfaceIpv6Ip6RouteListRead,
		Update: resourceSystemInterfaceIpv6Ip6RouteListUpdate,
		Delete: resourceSystemInterfaceIpv6Ip6RouteListDelete,

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
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"route": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"route_life_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"route_pref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemInterfaceIpv6Ip6RouteListCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	var_interface := d.Get("interface").(string)
	paradict["device"] = device_name
	paradict["interface"] = var_interface

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemInterfaceIpv6Ip6RouteList(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemInterfaceIpv6Ip6RouteList resource while getting object: %v", err)
	}

	_, err = c.CreateSystemInterfaceIpv6Ip6RouteList(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemInterfaceIpv6Ip6RouteList resource: %v", err)
	}

	d.SetId(getStringKey(d, "route"))

	return resourceSystemInterfaceIpv6Ip6RouteListRead(d, m)
}

func resourceSystemInterfaceIpv6Ip6RouteListUpdate(d *schema.ResourceData, m interface{}) error {
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
	var_interface := d.Get("interface").(string)
	paradict["device"] = device_name
	paradict["interface"] = var_interface

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemInterfaceIpv6Ip6RouteList(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterfaceIpv6Ip6RouteList resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemInterfaceIpv6Ip6RouteList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterfaceIpv6Ip6RouteList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "route"))

	return resourceSystemInterfaceIpv6Ip6RouteListRead(d, m)
}

func resourceSystemInterfaceIpv6Ip6RouteListDelete(d *schema.ResourceData, m interface{}) error {
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
	var_interface := d.Get("interface").(string)
	paradict["device"] = device_name
	paradict["interface"] = var_interface

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSystemInterfaceIpv6Ip6RouteList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemInterfaceIpv6Ip6RouteList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemInterfaceIpv6Ip6RouteListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	var_interface := d.Get("interface").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if var_interface == "" {
		var_interface = importOptionChecking(m.(*FortiClient).Cfg, "interface")
		if var_interface == "" {
			return fmt.Errorf("Parameter interface is missing")
		}
		if err = d.Set("interface", var_interface); err != nil {
			return fmt.Errorf("Error set params interface: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["interface"] = var_interface

	o, err := c.ReadSystemInterfaceIpv6Ip6RouteList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterfaceIpv6Ip6RouteList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemInterfaceIpv6Ip6RouteList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterfaceIpv6Ip6RouteList resource from API: %v", err)
	}
	return nil
}

func flattenSystemInterfaceIpv6Ip6RouteListRoute3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6RouteListRouteLifeTime3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6RouteListRoutePref3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemInterfaceIpv6Ip6RouteList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("route", flattenSystemInterfaceIpv6Ip6RouteListRoute3rdl(o["route"], d, "route")); err != nil {
		if vv, ok := fortiAPIPatch(o["route"], "SystemInterfaceIpv6Ip6RouteList-Route"); ok {
			if err = d.Set("route", vv); err != nil {
				return fmt.Errorf("Error reading route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route: %v", err)
		}
	}

	if err = d.Set("route_life_time", flattenSystemInterfaceIpv6Ip6RouteListRouteLifeTime3rdl(o["route-life-time"], d, "route_life_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-life-time"], "SystemInterfaceIpv6Ip6RouteList-RouteLifeTime"); ok {
			if err = d.Set("route_life_time", vv); err != nil {
				return fmt.Errorf("Error reading route_life_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_life_time: %v", err)
		}
	}

	if err = d.Set("route_pref", flattenSystemInterfaceIpv6Ip6RouteListRoutePref3rdl(o["route-pref"], d, "route_pref")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-pref"], "SystemInterfaceIpv6Ip6RouteList-RoutePref"); ok {
			if err = d.Set("route_pref", vv); err != nil {
				return fmt.Errorf("Error reading route_pref: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_pref: %v", err)
		}
	}

	return nil
}

func flattenSystemInterfaceIpv6Ip6RouteListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemInterfaceIpv6Ip6RouteListRoute3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6RouteListRouteLifeTime3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6RouteListRoutePref3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemInterfaceIpv6Ip6RouteList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("route"); ok || d.HasChange("route") {
		t, err := expandSystemInterfaceIpv6Ip6RouteListRoute3rdl(d, v, "route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route"] = t
		}
	}

	if v, ok := d.GetOk("route_life_time"); ok || d.HasChange("route_life_time") {
		t, err := expandSystemInterfaceIpv6Ip6RouteListRouteLifeTime3rdl(d, v, "route_life_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-life-time"] = t
		}
	}

	if v, ok := d.GetOk("route_pref"); ok || d.HasChange("route_pref") {
		t, err := expandSystemInterfaceIpv6Ip6RouteListRoutePref3rdl(d, v, "route_pref")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-pref"] = t
		}
	}

	return &obj, nil
}
