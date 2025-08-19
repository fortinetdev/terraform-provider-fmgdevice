// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Multicast-flow entries.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterMulticastFlowFlows() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterMulticastFlowFlowsCreate,
		Read:   resourceRouterMulticastFlowFlowsRead,
		Update: resourceRouterMulticastFlowFlowsUpdate,
		Delete: resourceRouterMulticastFlowFlowsDelete,

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
			"multicast_flow": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"group_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"source_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRouterMulticastFlowFlowsCreate(d *schema.ResourceData, m interface{}) error {
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
	multicast_flow := d.Get("multicast_flow").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["multicast_flow"] = multicast_flow

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectRouterMulticastFlowFlows(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterMulticastFlowFlows resource while getting object: %v", err)
	}

	_, err = c.CreateRouterMulticastFlowFlows(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating RouterMulticastFlowFlows resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterMulticastFlowFlowsRead(d, m)
}

func resourceRouterMulticastFlowFlowsUpdate(d *schema.ResourceData, m interface{}) error {
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
	multicast_flow := d.Get("multicast_flow").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["multicast_flow"] = multicast_flow

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectRouterMulticastFlowFlows(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticastFlowFlows resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterMulticastFlowFlows(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticastFlowFlows resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterMulticastFlowFlowsRead(d, m)
}

func resourceRouterMulticastFlowFlowsDelete(d *schema.ResourceData, m interface{}) error {
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
	multicast_flow := d.Get("multicast_flow").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["multicast_flow"] = multicast_flow

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteRouterMulticastFlowFlows(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterMulticastFlowFlows resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterMulticastFlowFlowsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	multicast_flow := d.Get("multicast_flow").(string)
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
	if multicast_flow == "" {
		multicast_flow = importOptionChecking(m.(*FortiClient).Cfg, "multicast_flow")
		if multicast_flow == "" {
			return fmt.Errorf("Parameter multicast_flow is missing")
		}
		if err = d.Set("multicast_flow", multicast_flow); err != nil {
			return fmt.Errorf("Error set params multicast_flow: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["multicast_flow"] = multicast_flow

	o, err := c.ReadRouterMulticastFlowFlows(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticastFlowFlows resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterMulticastFlowFlows(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticastFlowFlows resource from API: %v", err)
	}
	return nil
}

func flattenRouterMulticastFlowFlowsGroupAddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastFlowFlowsId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastFlowFlowsSourceAddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterMulticastFlowFlows(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("group_addr", flattenRouterMulticastFlowFlowsGroupAddr2edl(o["group-addr"], d, "group_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-addr"], "RouterMulticastFlowFlows-GroupAddr"); ok {
			if err = d.Set("group_addr", vv); err != nil {
				return fmt.Errorf("Error reading group_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_addr: %v", err)
		}
	}

	if err = d.Set("fosid", flattenRouterMulticastFlowFlowsId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "RouterMulticastFlowFlows-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("source_addr", flattenRouterMulticastFlowFlowsSourceAddr2edl(o["source-addr"], d, "source_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-addr"], "RouterMulticastFlowFlows-SourceAddr"); ok {
			if err = d.Set("source_addr", vv); err != nil {
				return fmt.Errorf("Error reading source_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_addr: %v", err)
		}
	}

	return nil
}

func flattenRouterMulticastFlowFlowsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterMulticastFlowFlowsGroupAddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastFlowFlowsId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastFlowFlowsSourceAddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterMulticastFlowFlows(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("group_addr"); ok || d.HasChange("group_addr") {
		t, err := expandRouterMulticastFlowFlowsGroupAddr2edl(d, v, "group_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-addr"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandRouterMulticastFlowFlowsId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("source_addr"); ok || d.HasChange("source_addr") {
		t, err := expandRouterMulticastFlowFlowsSourceAddr2edl(d, v, "source_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-addr"] = t
		}
	}

	return &obj, nil
}
