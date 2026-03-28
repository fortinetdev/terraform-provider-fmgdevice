// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure FSSO groups.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserAdgrp() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserAdgrpCreate,
		Read:   resourceUserAdgrpRead,
		Update: resourceUserAdgrpUpdate,
		Delete: resourceUserAdgrpDelete,

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
			"connector_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"server_name": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceUserAdgrpCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectUserAdgrp(d)
	if err != nil {
		return fmt.Errorf("Error creating UserAdgrp resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadUserAdgrp(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateUserAdgrp(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating UserAdgrp resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateUserAdgrp(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating UserAdgrp resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceUserAdgrpRead(d, m)
}

func resourceUserAdgrpUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectUserAdgrp(d)
	if err != nil {
		return fmt.Errorf("Error updating UserAdgrp resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateUserAdgrp(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating UserAdgrp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceUserAdgrpRead(d, m)
}

func resourceUserAdgrpDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	wsParams["adom"] = adomv

	err = c.DeleteUserAdgrp(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting UserAdgrp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserAdgrpRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserAdgrp(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading UserAdgrp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserAdgrp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserAdgrp resource from API: %v", err)
	}
	return nil
}

func flattenUserAdgrpConnectorSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserAdgrpId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserAdgrpName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserAdgrpServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectUserAdgrp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("connector_source", flattenUserAdgrpConnectorSource(o["connector-source"], d, "connector_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["connector-source"], "UserAdgrp-ConnectorSource"); ok {
			if err = d.Set("connector_source", vv); err != nil {
				return fmt.Errorf("Error reading connector_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading connector_source: %v", err)
		}
	}

	if err = d.Set("fosid", flattenUserAdgrpId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "UserAdgrp-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenUserAdgrpName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "UserAdgrp-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("server_name", flattenUserAdgrpServerName(o["server-name"], d, "server_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-name"], "UserAdgrp-ServerName"); ok {
			if err = d.Set("server_name", vv); err != nil {
				return fmt.Errorf("Error reading server_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_name: %v", err)
		}
	}

	return nil
}

func flattenUserAdgrpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserAdgrpConnectorSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserAdgrpId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserAdgrpName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserAdgrpServerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectUserAdgrp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("connector_source"); ok || d.HasChange("connector_source") {
		t, err := expandUserAdgrpConnectorSource(d, v, "connector_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connector-source"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandUserAdgrpId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandUserAdgrpName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("server_name"); ok || d.HasChange("server_name") {
		t, err := expandUserAdgrpServerName(d, v, "server_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-name"] = t
		}
	}

	return &obj, nil
}
