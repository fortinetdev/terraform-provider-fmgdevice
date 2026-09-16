// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IPv6 multicast-flow entries.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterMulticast6FlowFlows() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterMulticast6FlowFlowsCreate,
		Read:   resourceRouterMulticast6FlowFlowsRead,
		Update: resourceRouterMulticast6FlowFlowsUpdate,
		Delete: resourceRouterMulticast6FlowFlowsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
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
			"multicast6_flow": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"group_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceRouterMulticast6FlowFlowsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	multicast6_flow := d.Get("multicast6_flow").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["multicast6_flow"] = multicast6_flow

	obj, err := getObjectRouterMulticast6FlowFlows(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterMulticast6FlowFlows resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadRouterMulticast6FlowFlows(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateRouterMulticast6FlowFlows(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating RouterMulticast6FlowFlows resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateRouterMulticast6FlowFlows(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating RouterMulticast6FlowFlows resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterMulticast6FlowFlowsRead(d, m)
}

func resourceRouterMulticast6FlowFlowsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	multicast6_flow := d.Get("multicast6_flow").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["multicast6_flow"] = multicast6_flow

	obj, err := getObjectRouterMulticast6FlowFlows(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticast6FlowFlows resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateRouterMulticast6FlowFlows(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticast6FlowFlows resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterMulticast6FlowFlowsRead(d, m)
}

func resourceRouterMulticast6FlowFlowsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	multicast6_flow := d.Get("multicast6_flow").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["multicast6_flow"] = multicast6_flow

	wsParams["adom"] = adomv

	err = c.DeleteRouterMulticast6FlowFlows(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterMulticast6FlowFlows resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterMulticast6FlowFlowsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	multicast6_flow := d.Get("multicast6_flow").(string)
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
	if multicast6_flow == "" {
		multicast6_flow = importOptionChecking(m.(*FortiClient).Cfg, "multicast6_flow")
		if multicast6_flow == "" {
			return fmt.Errorf("Parameter multicast6_flow is missing")
		}
		if err = d.Set("multicast6_flow", multicast6_flow); err != nil {
			return fmt.Errorf("Error set params multicast6_flow: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["multicast6_flow"] = multicast6_flow

	o, err := c.ReadRouterMulticast6FlowFlows(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading RouterMulticast6FlowFlows resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterMulticast6FlowFlows(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticast6FlowFlows resource from API: %v", err)
	}
	return nil
}

func flattenRouterMulticast6FlowFlowsGroupAddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6FlowFlowsId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterMulticast6FlowFlows(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("group_addr", flattenRouterMulticast6FlowFlowsGroupAddr2edl(o["group-addr"], d, "group_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-addr"], "RouterMulticast6FlowFlows-GroupAddr"); ok {
			if err = d.Set("group_addr", vv); err != nil {
				return fmt.Errorf("Error reading group_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_addr: %v", err)
		}
	}

	if err = d.Set("fosid", flattenRouterMulticast6FlowFlowsId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "RouterMulticast6FlowFlows-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenRouterMulticast6FlowFlowsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterMulticast6FlowFlowsGroupAddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6FlowFlowsId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterMulticast6FlowFlows(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("group_addr"); ok || d.HasChange("group_addr") {
		t, err := expandRouterMulticast6FlowFlowsGroupAddr2edl(d, v, "group_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-addr"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandRouterMulticast6FlowFlowsId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
