// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure DHCP snooping option 82 override.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerManagedSwitchPortsDhcpSnoopOption82Override() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideCreate,
		Read:   resourceSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideRead,
		Update: resourceSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideUpdate,
		Delete: resourceSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideDelete,

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
			"managed_switch": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ports": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"circuit_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vlan_name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideCreate(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	ports := d.Get("ports").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch
	paradict["ports"] = ports

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82Override(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerManagedSwitchPortsDhcpSnoopOption82Override resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerManagedSwitchPortsDhcpSnoopOption82Override(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerManagedSwitchPortsDhcpSnoopOption82Override resource: %v", err)
	}

	d.SetId(getStringKey(d, "vlan_name"))

	return resourceSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideRead(d, m)
}

func resourceSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideUpdate(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	ports := d.Get("ports").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch
	paradict["ports"] = ports

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82Override(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchPortsDhcpSnoopOption82Override resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerManagedSwitchPortsDhcpSnoopOption82Override(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchPortsDhcpSnoopOption82Override resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "vlan_name"))

	return resourceSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideRead(d, m)
}

func resourceSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideDelete(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	ports := d.Get("ports").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch
	paradict["ports"] = ports

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSwitchControllerManagedSwitchPortsDhcpSnoopOption82Override(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerManagedSwitchPortsDhcpSnoopOption82Override resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	managed_switch := d.Get("managed_switch").(string)
	ports := d.Get("ports").(string)
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
	if managed_switch == "" {
		managed_switch = importOptionChecking(m.(*FortiClient).Cfg, "managed_switch")
		if managed_switch == "" {
			return fmt.Errorf("Parameter managed_switch is missing")
		}
		if err = d.Set("managed_switch", managed_switch); err != nil {
			return fmt.Errorf("Error set params managed_switch: %v", err)
		}
	}
	if ports == "" {
		ports = importOptionChecking(m.(*FortiClient).Cfg, "ports")
		if ports == "" {
			return fmt.Errorf("Parameter ports is missing")
		}
		if err = d.Set("ports", ports); err != nil {
			return fmt.Errorf("Error set params ports: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch
	paradict["ports"] = ports

	o, err := c.ReadSwitchControllerManagedSwitchPortsDhcpSnoopOption82Override(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchPortsDhcpSnoopOption82Override resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82Override(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchPortsDhcpSnoopOption82Override resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideCircuitId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideRemoteId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideVlanName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82Override(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("circuit_id", flattenSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideCircuitId3rdl(o["circuit-id"], d, "circuit_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["circuit-id"], "SwitchControllerManagedSwitchPortsDhcpSnoopOption82Override-CircuitId"); ok {
			if err = d.Set("circuit_id", vv); err != nil {
				return fmt.Errorf("Error reading circuit_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading circuit_id: %v", err)
		}
	}

	if err = d.Set("remote_id", flattenSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideRemoteId3rdl(o["remote-id"], d, "remote_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-id"], "SwitchControllerManagedSwitchPortsDhcpSnoopOption82Override-RemoteId"); ok {
			if err = d.Set("remote_id", vv); err != nil {
				return fmt.Errorf("Error reading remote_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_id: %v", err)
		}
	}

	if err = d.Set("vlan_name", flattenSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideVlanName3rdl(o["vlan-name"], d, "vlan_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-name"], "SwitchControllerManagedSwitchPortsDhcpSnoopOption82Override-VlanName"); ok {
			if err = d.Set("vlan_name", vv); err != nil {
				return fmt.Errorf("Error reading vlan_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_name: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideCircuitId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideRemoteId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideVlanName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82Override(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("circuit_id"); ok || d.HasChange("circuit_id") {
		t, err := expandSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideCircuitId3rdl(d, v, "circuit_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["circuit-id"] = t
		}
	}

	if v, ok := d.GetOk("remote_id"); ok || d.HasChange("remote_id") {
		t, err := expandSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideRemoteId3rdl(d, v, "remote_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-id"] = t
		}
	}

	if v, ok := d.GetOk("vlan_name"); ok || d.HasChange("vlan_name") {
		t, err := expandSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideVlanName3rdl(d, v, "vlan_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-name"] = t
		}
	}

	return &obj, nil
}
