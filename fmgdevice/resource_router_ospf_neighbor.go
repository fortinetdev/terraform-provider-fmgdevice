// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: OSPF neighbor configuration are used when OSPF runs on non-broadcast media.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterOspfNeighbor() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterOspfNeighborCreate,
		Read:   resourceRouterOspfNeighborRead,
		Update: resourceRouterOspfNeighborUpdate,
		Delete: resourceRouterOspfNeighborDelete,

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
			"cost": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func resourceRouterOspfNeighborCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterOspfNeighbor(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspfNeighbor resource while getting object: %v", err)
	}

	_, err = c.CreateRouterOspfNeighbor(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspfNeighbor resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterOspfNeighborRead(d, m)
}

func resourceRouterOspfNeighborUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterOspfNeighbor(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspfNeighbor resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterOspfNeighbor(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspfNeighbor resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterOspfNeighborRead(d, m)
}

func resourceRouterOspfNeighborDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterOspfNeighbor(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterOspfNeighbor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterOspfNeighborRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterOspfNeighbor(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspfNeighbor resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterOspfNeighbor(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspfNeighbor resource from API: %v", err)
	}
	return nil
}

func flattenRouterOspfNeighborCost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfNeighborId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfNeighborIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfNeighborPollInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfNeighborPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterOspfNeighbor(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("cost", flattenRouterOspfNeighborCost2edl(o["cost"], d, "cost")); err != nil {
		if vv, ok := fortiAPIPatch(o["cost"], "RouterOspfNeighbor-Cost"); ok {
			if err = d.Set("cost", vv); err != nil {
				return fmt.Errorf("Error reading cost: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cost: %v", err)
		}
	}

	if err = d.Set("fosid", flattenRouterOspfNeighborId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "RouterOspfNeighbor-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip", flattenRouterOspfNeighborIp2edl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "RouterOspfNeighbor-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("poll_interval", flattenRouterOspfNeighborPollInterval2edl(o["poll-interval"], d, "poll_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["poll-interval"], "RouterOspfNeighbor-PollInterval"); ok {
			if err = d.Set("poll_interval", vv); err != nil {
				return fmt.Errorf("Error reading poll_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poll_interval: %v", err)
		}
	}

	if err = d.Set("priority", flattenRouterOspfNeighborPriority2edl(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "RouterOspfNeighbor-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	return nil
}

func flattenRouterOspfNeighborFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterOspfNeighborCost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfNeighborId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfNeighborIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfNeighborPollInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfNeighborPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterOspfNeighbor(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cost"); ok || d.HasChange("cost") {
		t, err := expandRouterOspfNeighborCost2edl(d, v, "cost")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cost"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandRouterOspfNeighborId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandRouterOspfNeighborIp2edl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("poll_interval"); ok || d.HasChange("poll_interval") {
		t, err := expandRouterOspfNeighborPollInterval2edl(d, v, "poll_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poll-interval"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandRouterOspfNeighborPriority2edl(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	return &obj, nil
}
