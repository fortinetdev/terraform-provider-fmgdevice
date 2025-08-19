// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure router settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterSettingUpdate,
		Read:   resourceRouterSettingRead,
		Update: resourceRouterSettingUpdate,
		Delete: resourceRouterSettingDelete,

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
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"kernel_route_distance": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"show_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRouterSettingUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterSetting(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("RouterSetting")

	return resourceRouterSettingRead(d, m)
}

func resourceRouterSettingDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterSetting(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterSettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterSetting resource from API: %v", err)
	}
	return nil
}

func flattenRouterSettingHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterSettingKernelRouteDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterSettingShowFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectRouterSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("hostname", flattenRouterSettingHostname(o["hostname"], d, "hostname")); err != nil {
		if vv, ok := fortiAPIPatch(o["hostname"], "RouterSetting-Hostname"); ok {
			if err = d.Set("hostname", vv); err != nil {
				return fmt.Errorf("Error reading hostname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("kernel_route_distance", flattenRouterSettingKernelRouteDistance(o["kernel-route-distance"], d, "kernel_route_distance")); err != nil {
		if vv, ok := fortiAPIPatch(o["kernel-route-distance"], "RouterSetting-KernelRouteDistance"); ok {
			if err = d.Set("kernel_route_distance", vv); err != nil {
				return fmt.Errorf("Error reading kernel_route_distance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading kernel_route_distance: %v", err)
		}
	}

	if err = d.Set("show_filter", flattenRouterSettingShowFilter(o["show-filter"], d, "show_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["show-filter"], "RouterSetting-ShowFilter"); ok {
			if err = d.Set("show_filter", vv); err != nil {
				return fmt.Errorf("Error reading show_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading show_filter: %v", err)
		}
	}

	return nil
}

func flattenRouterSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterSettingHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingKernelRouteDistance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingShowFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectRouterSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("hostname"); ok || d.HasChange("hostname") {
		t, err := expandRouterSettingHostname(d, v, "hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostname"] = t
		}
	}

	if v, ok := d.GetOk("kernel_route_distance"); ok || d.HasChange("kernel_route_distance") {
		t, err := expandRouterSettingKernelRouteDistance(d, v, "kernel_route_distance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["kernel-route-distance"] = t
		}
	}

	if v, ok := d.GetOk("show_filter"); ok || d.HasChange("show_filter") {
		t, err := expandRouterSettingShowFilter(d, v, "show_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["show-filter"] = t
		}
	}

	return &obj, nil
}
