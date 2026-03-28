// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device IcapRemoteServerGroupServerList

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIcapRemoteServerGroupServerList() *schema.Resource {
	return &schema.Resource{
		Create: resourceIcapRemoteServerGroupServerListCreate,
		Read:   resourceIcapRemoteServerGroupServerListRead,
		Update: resourceIcapRemoteServerGroupServerListUpdate,
		Delete: resourceIcapRemoteServerGroupServerListDelete,

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
			"remote_server_group": &schema.Schema{
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

func resourceIcapRemoteServerGroupServerListCreate(d *schema.ResourceData, m interface{}) error {
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
	remote_server_group := d.Get("remote_server_group").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["remote_server_group"] = remote_server_group

	obj, err := getObjectIcapRemoteServerGroupServerList(d)
	if err != nil {
		return fmt.Errorf("Error creating IcapRemoteServerGroupServerList resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadIcapRemoteServerGroupServerList(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateIcapRemoteServerGroupServerList(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating IcapRemoteServerGroupServerList resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateIcapRemoteServerGroupServerList(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating IcapRemoteServerGroupServerList resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceIcapRemoteServerGroupServerListRead(d, m)
}

func resourceIcapRemoteServerGroupServerListUpdate(d *schema.ResourceData, m interface{}) error {
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
	remote_server_group := d.Get("remote_server_group").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["remote_server_group"] = remote_server_group

	obj, err := getObjectIcapRemoteServerGroupServerList(d)
	if err != nil {
		return fmt.Errorf("Error updating IcapRemoteServerGroupServerList resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateIcapRemoteServerGroupServerList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating IcapRemoteServerGroupServerList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceIcapRemoteServerGroupServerListRead(d, m)
}

func resourceIcapRemoteServerGroupServerListDelete(d *schema.ResourceData, m interface{}) error {
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
	remote_server_group := d.Get("remote_server_group").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["remote_server_group"] = remote_server_group

	wsParams["adom"] = adomv

	err = c.DeleteIcapRemoteServerGroupServerList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting IcapRemoteServerGroupServerList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIcapRemoteServerGroupServerListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	remote_server_group := d.Get("remote_server_group").(string)
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
	if remote_server_group == "" {
		remote_server_group = importOptionChecking(m.(*FortiClient).Cfg, "remote_server_group")
		if remote_server_group == "" {
			return fmt.Errorf("Parameter remote_server_group is missing")
		}
		if err = d.Set("remote_server_group", remote_server_group); err != nil {
			return fmt.Errorf("Error set params remote_server_group: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["remote_server_group"] = remote_server_group

	o, err := c.ReadIcapRemoteServerGroupServerList(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading IcapRemoteServerGroupServerList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIcapRemoteServerGroupServerList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading IcapRemoteServerGroupServerList resource from API: %v", err)
	}
	return nil
}

func flattenIcapRemoteServerGroupServerListName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenIcapRemoteServerGroupServerListWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectIcapRemoteServerGroupServerList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenIcapRemoteServerGroupServerListName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "IcapRemoteServerGroupServerList-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("weight", flattenIcapRemoteServerGroupServerListWeight2edl(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "IcapRemoteServerGroupServerList-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	return nil
}

func flattenIcapRemoteServerGroupServerListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandIcapRemoteServerGroupServerListName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapRemoteServerGroupServerListWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectIcapRemoteServerGroupServerList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandIcapRemoteServerGroupServerListName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandIcapRemoteServerGroupServerListWeight2edl(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	return &obj, nil
}
