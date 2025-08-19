// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Join multicast groups.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterMulticastInterfaceJoinGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterMulticastInterfaceJoinGroupCreate,
		Read:   resourceRouterMulticastInterfaceJoinGroupRead,
		Update: resourceRouterMulticastInterfaceJoinGroupUpdate,
		Delete: resourceRouterMulticastInterfaceJoinGroupDelete,

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
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceRouterMulticastInterfaceJoinGroupCreate(d *schema.ResourceData, m interface{}) error {
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
	var_interface := d.Get("interface").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["interface"] = var_interface

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectRouterMulticastInterfaceJoinGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterMulticastInterfaceJoinGroup resource while getting object: %v", err)
	}

	_, err = c.CreateRouterMulticastInterfaceJoinGroup(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating RouterMulticastInterfaceJoinGroup resource: %v", err)
	}

	d.SetId(getStringKey(d, "address"))

	return resourceRouterMulticastInterfaceJoinGroupRead(d, m)
}

func resourceRouterMulticastInterfaceJoinGroupUpdate(d *schema.ResourceData, m interface{}) error {
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
	var_interface := d.Get("interface").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["interface"] = var_interface

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectRouterMulticastInterfaceJoinGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticastInterfaceJoinGroup resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterMulticastInterfaceJoinGroup(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticastInterfaceJoinGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "address"))

	return resourceRouterMulticastInterfaceJoinGroupRead(d, m)
}

func resourceRouterMulticastInterfaceJoinGroupDelete(d *schema.ResourceData, m interface{}) error {
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
	var_interface := d.Get("interface").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["interface"] = var_interface

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteRouterMulticastInterfaceJoinGroup(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterMulticastInterfaceJoinGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterMulticastInterfaceJoinGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
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
	if device_vdom == "" {
		device_vdom = importOptionChecking(m.(*FortiClient).Cfg, "device_vdom")
		if device_vdom == "" {
			return fmt.Errorf("Parameter device_vdom is missing")
		}
		if err = d.Set("device_vdom", device_vdom); err != nil {
			return fmt.Errorf("Error set params device_vdom: %v", err)
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
	paradict["vdom"] = device_vdom
	paradict["interface"] = var_interface

	o, err := c.ReadRouterMulticastInterfaceJoinGroup(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticastInterfaceJoinGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterMulticastInterfaceJoinGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticastInterfaceJoinGroup resource from API: %v", err)
	}
	return nil
}

func flattenRouterMulticastInterfaceJoinGroupAddress3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterMulticastInterfaceJoinGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("address", flattenRouterMulticastInterfaceJoinGroupAddress3rdl(o["address"], d, "address")); err != nil {
		if vv, ok := fortiAPIPatch(o["address"], "RouterMulticastInterfaceJoinGroup-Address"); ok {
			if err = d.Set("address", vv); err != nil {
				return fmt.Errorf("Error reading address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	return nil
}

func flattenRouterMulticastInterfaceJoinGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterMulticastInterfaceJoinGroupAddress3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterMulticastInterfaceJoinGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address"); ok || d.HasChange("address") {
		t, err := expandRouterMulticastInterfaceJoinGroupAddress3rdl(d, v, "address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	return &obj, nil
}
