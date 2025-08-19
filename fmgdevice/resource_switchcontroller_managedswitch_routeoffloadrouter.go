// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure route offload MCLAG IP address.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerManagedSwitchRouteOffloadRouter() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerManagedSwitchRouteOffloadRouterCreate,
		Read:   resourceSwitchControllerManagedSwitchRouteOffloadRouterRead,
		Update: resourceSwitchControllerManagedSwitchRouteOffloadRouterUpdate,
		Delete: resourceSwitchControllerManagedSwitchRouteOffloadRouterDelete,

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
			"managed_switch": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"router_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan_name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceSwitchControllerManagedSwitchRouteOffloadRouterCreate(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerManagedSwitchRouteOffloadRouter(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerManagedSwitchRouteOffloadRouter resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerManagedSwitchRouteOffloadRouter(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerManagedSwitchRouteOffloadRouter resource: %v", err)
	}

	d.SetId(getStringKey(d, "vlan_name"))

	return resourceSwitchControllerManagedSwitchRouteOffloadRouterRead(d, m)
}

func resourceSwitchControllerManagedSwitchRouteOffloadRouterUpdate(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerManagedSwitchRouteOffloadRouter(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchRouteOffloadRouter resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerManagedSwitchRouteOffloadRouter(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchRouteOffloadRouter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "vlan_name"))

	return resourceSwitchControllerManagedSwitchRouteOffloadRouterRead(d, m)
}

func resourceSwitchControllerManagedSwitchRouteOffloadRouterDelete(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSwitchControllerManagedSwitchRouteOffloadRouter(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerManagedSwitchRouteOffloadRouter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerManagedSwitchRouteOffloadRouterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	managed_switch := d.Get("managed_switch").(string)
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
	if managed_switch == "" {
		managed_switch = importOptionChecking(m.(*FortiClient).Cfg, "managed_switch")
		if managed_switch == "" {
			return fmt.Errorf("Parameter managed_switch is missing")
		}
		if err = d.Set("managed_switch", managed_switch); err != nil {
			return fmt.Errorf("Error set params managed_switch: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	o, err := c.ReadSwitchControllerManagedSwitchRouteOffloadRouter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchRouteOffloadRouter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerManagedSwitchRouteOffloadRouter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchRouteOffloadRouter resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerManagedSwitchRouteOffloadRouterRouterIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRouteOffloadRouterVlanName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerManagedSwitchRouteOffloadRouter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("router_ip", flattenSwitchControllerManagedSwitchRouteOffloadRouterRouterIp2edl(o["router-ip"], d, "router_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["router-ip"], "SwitchControllerManagedSwitchRouteOffloadRouter-RouterIp"); ok {
			if err = d.Set("router_ip", vv); err != nil {
				return fmt.Errorf("Error reading router_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading router_ip: %v", err)
		}
	}

	if err = d.Set("vlan_name", flattenSwitchControllerManagedSwitchRouteOffloadRouterVlanName2edl(o["vlan-name"], d, "vlan_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-name"], "SwitchControllerManagedSwitchRouteOffloadRouter-VlanName"); ok {
			if err = d.Set("vlan_name", vv); err != nil {
				return fmt.Errorf("Error reading vlan_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_name: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerManagedSwitchRouteOffloadRouterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerManagedSwitchRouteOffloadRouterRouterIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRouteOffloadRouterVlanName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerManagedSwitchRouteOffloadRouter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("router_ip"); ok || d.HasChange("router_ip") {
		t, err := expandSwitchControllerManagedSwitchRouteOffloadRouterRouterIp2edl(d, v, "router_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["router-ip"] = t
		}
	}

	if v, ok := d.GetOk("vlan_name"); ok || d.HasChange("vlan_name") {
		t, err := expandSwitchControllerManagedSwitchRouteOffloadRouterVlanName2edl(d, v, "vlan_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-name"] = t
		}
	}

	return &obj, nil
}
