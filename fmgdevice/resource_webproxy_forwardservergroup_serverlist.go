// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Add web forward servers to a list to form a server group. Optionally assign weights to each server.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebProxyForwardServerGroupServerList() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebProxyForwardServerGroupServerListCreate,
		Read:   resourceWebProxyForwardServerGroupServerListRead,
		Update: resourceWebProxyForwardServerGroupServerListUpdate,
		Delete: resourceWebProxyForwardServerGroupServerListDelete,

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
			"forward_server_group": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWebProxyForwardServerGroupServerListCreate(d *schema.ResourceData, m interface{}) error {
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
	forward_server_group := d.Get("forward_server_group").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["forward_server_group"] = forward_server_group

	obj, err := getObjectWebProxyForwardServerGroupServerList(d)
	if err != nil {
		return fmt.Errorf("Error creating WebProxyForwardServerGroupServerList resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWebProxyForwardServerGroupServerList(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWebProxyForwardServerGroupServerList(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WebProxyForwardServerGroupServerList resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateWebProxyForwardServerGroupServerList(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WebProxyForwardServerGroupServerList resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceWebProxyForwardServerGroupServerListRead(d, m)
}

func resourceWebProxyForwardServerGroupServerListUpdate(d *schema.ResourceData, m interface{}) error {
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
	forward_server_group := d.Get("forward_server_group").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["forward_server_group"] = forward_server_group

	obj, err := getObjectWebProxyForwardServerGroupServerList(d)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyForwardServerGroupServerList resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWebProxyForwardServerGroupServerList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyForwardServerGroupServerList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWebProxyForwardServerGroupServerListRead(d, m)
}

func resourceWebProxyForwardServerGroupServerListDelete(d *schema.ResourceData, m interface{}) error {
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
	forward_server_group := d.Get("forward_server_group").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["forward_server_group"] = forward_server_group

	wsParams["adom"] = adomv

	err = c.DeleteWebProxyForwardServerGroupServerList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebProxyForwardServerGroupServerList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyForwardServerGroupServerListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	forward_server_group := d.Get("forward_server_group").(string)
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
	if forward_server_group == "" {
		forward_server_group = importOptionChecking(m.(*FortiClient).Cfg, "forward_server_group")
		if forward_server_group == "" {
			return fmt.Errorf("Parameter forward_server_group is missing")
		}
		if err = d.Set("forward_server_group", forward_server_group); err != nil {
			return fmt.Errorf("Error set params forward_server_group: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["forward_server_group"] = forward_server_group

	o, err := c.ReadWebProxyForwardServerGroupServerList(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WebProxyForwardServerGroupServerList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebProxyForwardServerGroupServerList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyForwardServerGroupServerList resource from API: %v", err)
	}
	return nil
}

func flattenWebProxyForwardServerGroupServerListName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenWebProxyForwardServerGroupServerListWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebProxyForwardServerGroupServerList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenWebProxyForwardServerGroupServerListName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WebProxyForwardServerGroupServerList-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("weight", flattenWebProxyForwardServerGroupServerListWeight2edl(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "WebProxyForwardServerGroupServerList-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	return nil
}

func flattenWebProxyForwardServerGroupServerListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebProxyForwardServerGroupServerListName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyForwardServerGroupServerListWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebProxyForwardServerGroupServerList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWebProxyForwardServerGroupServerListName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandWebProxyForwardServerGroupServerListWeight2edl(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	return &obj, nil
}
