// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: OSPFv3 neighbors are used when OSPFv3 runs on non-broadcast media.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterOspf6Ospf6InterfaceNeighbor() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterOspf6Ospf6InterfaceNeighborCreate,
		Read:   resourceRouterOspf6Ospf6InterfaceNeighborRead,
		Update: resourceRouterOspf6Ospf6InterfaceNeighborUpdate,
		Delete: resourceRouterOspf6Ospf6InterfaceNeighborDelete,

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
			"ospf6_interface": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cost": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"poll_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRouterOspf6Ospf6InterfaceNeighborCreate(d *schema.ResourceData, m interface{}) error {
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
	ospf6_interface := d.Get("ospf6_interface").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["ospf6_interface"] = ospf6_interface

	obj, err := getObjectRouterOspf6Ospf6InterfaceNeighbor(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspf6Ospf6InterfaceNeighbor resource while getting object: %v", err)
	}

	_, err = c.CreateRouterOspf6Ospf6InterfaceNeighbor(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating RouterOspf6Ospf6InterfaceNeighbor resource: %v", err)
	}

	d.SetId(getStringKey(d, "ip6"))

	return resourceRouterOspf6Ospf6InterfaceNeighborRead(d, m)
}

func resourceRouterOspf6Ospf6InterfaceNeighborUpdate(d *schema.ResourceData, m interface{}) error {
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
	ospf6_interface := d.Get("ospf6_interface").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["ospf6_interface"] = ospf6_interface

	obj, err := getObjectRouterOspf6Ospf6InterfaceNeighbor(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspf6Ospf6InterfaceNeighbor resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterOspf6Ospf6InterfaceNeighbor(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspf6Ospf6InterfaceNeighbor resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "ip6"))

	return resourceRouterOspf6Ospf6InterfaceNeighborRead(d, m)
}

func resourceRouterOspf6Ospf6InterfaceNeighborDelete(d *schema.ResourceData, m interface{}) error {
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
	ospf6_interface := d.Get("ospf6_interface").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["ospf6_interface"] = ospf6_interface

	err = c.DeleteRouterOspf6Ospf6InterfaceNeighbor(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting RouterOspf6Ospf6InterfaceNeighbor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterOspf6Ospf6InterfaceNeighborRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	ospf6_interface := d.Get("ospf6_interface").(string)
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
	if ospf6_interface == "" {
		ospf6_interface = importOptionChecking(m.(*FortiClient).Cfg, "ospf6_interface")
		if ospf6_interface == "" {
			return fmt.Errorf("Parameter ospf6_interface is missing")
		}
		if err = d.Set("ospf6_interface", ospf6_interface); err != nil {
			return fmt.Errorf("Error set params ospf6_interface: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["ospf6_interface"] = ospf6_interface

	o, err := c.ReadRouterOspf6Ospf6InterfaceNeighbor(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspf6Ospf6InterfaceNeighbor resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterOspf6Ospf6InterfaceNeighbor(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspf6Ospf6InterfaceNeighbor resource from API: %v", err)
	}
	return nil
}

func flattenRouterOspf6Ospf6InterfaceNeighborCost3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceNeighborIp63rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceNeighborPollInterval3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceNeighborPriority3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterOspf6Ospf6InterfaceNeighbor(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("cost", flattenRouterOspf6Ospf6InterfaceNeighborCost3rdl(o["cost"], d, "cost")); err != nil {
		if vv, ok := fortiAPIPatch(o["cost"], "RouterOspf6Ospf6InterfaceNeighbor-Cost"); ok {
			if err = d.Set("cost", vv); err != nil {
				return fmt.Errorf("Error reading cost: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cost: %v", err)
		}
	}

	if err = d.Set("ip6", flattenRouterOspf6Ospf6InterfaceNeighborIp63rdl(o["ip6"], d, "ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6"], "RouterOspf6Ospf6InterfaceNeighbor-Ip6"); ok {
			if err = d.Set("ip6", vv); err != nil {
				return fmt.Errorf("Error reading ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("poll_interval", flattenRouterOspf6Ospf6InterfaceNeighborPollInterval3rdl(o["poll-interval"], d, "poll_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["poll-interval"], "RouterOspf6Ospf6InterfaceNeighbor-PollInterval"); ok {
			if err = d.Set("poll_interval", vv); err != nil {
				return fmt.Errorf("Error reading poll_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poll_interval: %v", err)
		}
	}

	if err = d.Set("priority", flattenRouterOspf6Ospf6InterfaceNeighborPriority3rdl(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "RouterOspf6Ospf6InterfaceNeighbor-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	return nil
}

func flattenRouterOspf6Ospf6InterfaceNeighborFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterOspf6Ospf6InterfaceNeighborCost3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceNeighborIp63rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceNeighborPollInterval3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceNeighborPriority3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterOspf6Ospf6InterfaceNeighbor(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cost"); ok || d.HasChange("cost") {
		t, err := expandRouterOspf6Ospf6InterfaceNeighborCost3rdl(d, v, "cost")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cost"] = t
		}
	}

	if v, ok := d.GetOk("ip6"); ok || d.HasChange("ip6") {
		t, err := expandRouterOspf6Ospf6InterfaceNeighborIp63rdl(d, v, "ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6"] = t
		}
	}

	if v, ok := d.GetOk("poll_interval"); ok || d.HasChange("poll_interval") {
		t, err := expandRouterOspf6Ospf6InterfaceNeighborPollInterval3rdl(d, v, "poll_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poll-interval"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandRouterOspf6Ospf6InterfaceNeighborPriority3rdl(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	return &obj, nil
}
