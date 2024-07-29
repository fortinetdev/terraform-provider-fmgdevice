// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Table for mapping VLAN name to VLAN ID.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerVapVlanName() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerVapVlanNameCreate,
		Read:   resourceWirelessControllerVapVlanNameRead,
		Update: resourceWirelessControllerVapVlanNameUpdate,
		Delete: resourceWirelessControllerVapVlanNameDelete,

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
			"vap": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"vlan_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceWirelessControllerVapVlanNameCreate(d *schema.ResourceData, m interface{}) error {
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
	vap := d.Get("vap").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vap"] = vap

	obj, err := getObjectWirelessControllerVapVlanName(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerVapVlanName resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerVapVlanName(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerVapVlanName resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerVapVlanNameRead(d, m)
}

func resourceWirelessControllerVapVlanNameUpdate(d *schema.ResourceData, m interface{}) error {
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
	vap := d.Get("vap").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vap"] = vap

	obj, err := getObjectWirelessControllerVapVlanName(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerVapVlanName resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerVapVlanName(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerVapVlanName resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerVapVlanNameRead(d, m)
}

func resourceWirelessControllerVapVlanNameDelete(d *schema.ResourceData, m interface{}) error {
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
	vap := d.Get("vap").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vap"] = vap

	err = c.DeleteWirelessControllerVapVlanName(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerVapVlanName resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerVapVlanNameRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	vap := d.Get("vap").(string)
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
	if vap == "" {
		vap = importOptionChecking(m.(*FortiClient).Cfg, "vap")
		if vap == "" {
			return fmt.Errorf("Parameter vap is missing")
		}
		if err = d.Set("vap", vap); err != nil {
			return fmt.Errorf("Error set params vap: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vap"] = vap

	o, err := c.ReadWirelessControllerVapVlanName(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerVapVlanName resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerVapVlanName(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerVapVlanName resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerVapVlanNameName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapVlanNameVlanId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerVapVlanName(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenWirelessControllerVapVlanNameName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerVapVlanName-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("vlan_id", flattenWirelessControllerVapVlanNameVlanId2edl(o["vlan-id"], d, "vlan_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-id"], "WirelessControllerVapVlanName-VlanId"); ok {
			if err = d.Set("vlan_id", vv); err != nil {
				return fmt.Errorf("Error reading vlan_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_id: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerVapVlanNameFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerVapVlanNameName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapVlanNameVlanId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerVapVlanName(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerVapVlanNameName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("vlan_id"); ok || d.HasChange("vlan_id") {
		t, err := expandWirelessControllerVapVlanNameVlanId2edl(d, v, "vlan_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-id"] = t
		}
	}

	return &obj, nil
}
