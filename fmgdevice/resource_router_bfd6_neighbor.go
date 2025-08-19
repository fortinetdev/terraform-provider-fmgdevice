// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure neighbor of IPv6 BFD.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterBfd6Neighbor() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterBfd6NeighborCreate,
		Read:   resourceRouterBfd6NeighborRead,
		Update: resourceRouterBfd6NeighborUpdate,
		Delete: resourceRouterBfd6NeighborDelete,

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
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ip6_address": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceRouterBfd6NeighborCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterBfd6Neighbor(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterBfd6Neighbor resource while getting object: %v", err)
	}

	_, err = c.CreateRouterBfd6Neighbor(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating RouterBfd6Neighbor resource: %v", err)
	}

	d.SetId(getStringKey(d, "ip6_address"))

	return resourceRouterBfd6NeighborRead(d, m)
}

func resourceRouterBfd6NeighborUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterBfd6Neighbor(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterBfd6Neighbor resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterBfd6Neighbor(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterBfd6Neighbor resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "ip6_address"))

	return resourceRouterBfd6NeighborRead(d, m)
}

func resourceRouterBfd6NeighborDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterBfd6Neighbor(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterBfd6Neighbor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterBfd6NeighborRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterBfd6Neighbor(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterBfd6Neighbor resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterBfd6Neighbor(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterBfd6Neighbor resource from API: %v", err)
	}
	return nil
}

func flattenRouterBfd6NeighborInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBfd6NeighborIp6Address2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterBfd6Neighbor(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("interface", flattenRouterBfd6NeighborInterface2edl(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "RouterBfd6Neighbor-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ip6_address", flattenRouterBfd6NeighborIp6Address2edl(o["ip6-address"], d, "ip6_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-address"], "RouterBfd6Neighbor-Ip6Address"); ok {
			if err = d.Set("ip6_address", vv); err != nil {
				return fmt.Errorf("Error reading ip6_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_address: %v", err)
		}
	}

	return nil
}

func flattenRouterBfd6NeighborFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterBfd6NeighborInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBfd6NeighborIp6Address2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterBfd6Neighbor(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandRouterBfd6NeighborInterface2edl(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ip6_address"); ok || d.HasChange("ip6_address") {
		t, err := expandRouterBfd6NeighborIp6Address2edl(d, v, "ip6_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-address"] = t
		}
	}

	return &obj, nil
}
