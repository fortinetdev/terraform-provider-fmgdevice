// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: FortiClient user authentication groups.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnOcvpnForticlientAccessAuthGroups() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnOcvpnForticlientAccessAuthGroupsCreate,
		Read:   resourceVpnOcvpnForticlientAccessAuthGroupsRead,
		Update: resourceVpnOcvpnForticlientAccessAuthGroupsUpdate,
		Delete: resourceVpnOcvpnForticlientAccessAuthGroupsDelete,

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
			"auth_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"overlays": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceVpnOcvpnForticlientAccessAuthGroupsCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectVpnOcvpnForticlientAccessAuthGroups(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnOcvpnForticlientAccessAuthGroups resource while getting object: %v", err)
	}

	_, err = c.CreateVpnOcvpnForticlientAccessAuthGroups(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating VpnOcvpnForticlientAccessAuthGroups resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceVpnOcvpnForticlientAccessAuthGroupsRead(d, m)
}

func resourceVpnOcvpnForticlientAccessAuthGroupsUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectVpnOcvpnForticlientAccessAuthGroups(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnOcvpnForticlientAccessAuthGroups resource while getting object: %v", err)
	}

	_, err = c.UpdateVpnOcvpnForticlientAccessAuthGroups(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VpnOcvpnForticlientAccessAuthGroups resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceVpnOcvpnForticlientAccessAuthGroupsRead(d, m)
}

func resourceVpnOcvpnForticlientAccessAuthGroupsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteVpnOcvpnForticlientAccessAuthGroups(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VpnOcvpnForticlientAccessAuthGroups resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnOcvpnForticlientAccessAuthGroupsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnOcvpnForticlientAccessAuthGroups(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading VpnOcvpnForticlientAccessAuthGroups resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnOcvpnForticlientAccessAuthGroups(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnOcvpnForticlientAccessAuthGroups resource from API: %v", err)
	}
	return nil
}

func flattenVpnOcvpnForticlientAccessAuthGroupsAuthGroup3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnOcvpnForticlientAccessAuthGroupsName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnForticlientAccessAuthGroupsOverlays3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectVpnOcvpnForticlientAccessAuthGroups(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auth_group", flattenVpnOcvpnForticlientAccessAuthGroupsAuthGroup3rdl(o["auth-group"], d, "auth_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-group"], "VpnOcvpnForticlientAccessAuthGroups-AuthGroup"); ok {
			if err = d.Set("auth_group", vv); err != nil {
				return fmt.Errorf("Error reading auth_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_group: %v", err)
		}
	}

	if err = d.Set("name", flattenVpnOcvpnForticlientAccessAuthGroupsName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "VpnOcvpnForticlientAccessAuthGroups-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("overlays", flattenVpnOcvpnForticlientAccessAuthGroupsOverlays3rdl(o["overlays"], d, "overlays")); err != nil {
		if vv, ok := fortiAPIPatch(o["overlays"], "VpnOcvpnForticlientAccessAuthGroups-Overlays"); ok {
			if err = d.Set("overlays", vv); err != nil {
				return fmt.Errorf("Error reading overlays: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading overlays: %v", err)
		}
	}

	return nil
}

func flattenVpnOcvpnForticlientAccessAuthGroupsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnOcvpnForticlientAccessAuthGroupsAuthGroup3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnOcvpnForticlientAccessAuthGroupsName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnForticlientAccessAuthGroupsOverlays3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectVpnOcvpnForticlientAccessAuthGroups(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_group"); ok || d.HasChange("auth_group") {
		t, err := expandVpnOcvpnForticlientAccessAuthGroupsAuthGroup3rdl(d, v, "auth_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-group"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandVpnOcvpnForticlientAccessAuthGroupsName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("overlays"); ok || d.HasChange("overlays") {
		t, err := expandVpnOcvpnForticlientAccessAuthGroupsOverlays3rdl(d, v, "overlays")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["overlays"] = t
		}
	}

	return &obj, nil
}
