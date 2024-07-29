// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IP in IP Tunneling.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemIpipTunnel() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemIpipTunnelCreate,
		Read:   resourceSystemIpipTunnelRead,
		Update: resourceSystemIpipTunnelUpdate,
		Delete: resourceSystemIpipTunnelDelete,

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
			"auto_asic_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"local_gw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"remote_gw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_sdwan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemIpipTunnelCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectSystemIpipTunnel(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemIpipTunnel resource while getting object: %v", err)
	}

	_, err = c.CreateSystemIpipTunnel(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemIpipTunnel resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemIpipTunnelRead(d, m)
}

func resourceSystemIpipTunnelUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectSystemIpipTunnel(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemIpipTunnel resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemIpipTunnel(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemIpipTunnel resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemIpipTunnelRead(d, m)
}

func resourceSystemIpipTunnelDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	err = c.DeleteSystemIpipTunnel(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemIpipTunnel resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemIpipTunnelRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemIpipTunnel(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemIpipTunnel resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemIpipTunnel(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemIpipTunnel resource from API: %v", err)
	}
	return nil
}

func flattenSystemIpipTunnelAutoAsicOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpipTunnelInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemIpipTunnelLocalGw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpipTunnelName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpipTunnelRemoteGw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpipTunnelUseSdwan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemIpipTunnel(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auto_asic_offload", flattenSystemIpipTunnelAutoAsicOffload(o["auto-asic-offload"], d, "auto_asic_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-asic-offload"], "SystemIpipTunnel-AutoAsicOffload"); ok {
			if err = d.Set("auto_asic_offload", vv); err != nil {
				return fmt.Errorf("Error reading auto_asic_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_asic_offload: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemIpipTunnelInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemIpipTunnel-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("local_gw", flattenSystemIpipTunnelLocalGw(o["local-gw"], d, "local_gw")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-gw"], "SystemIpipTunnel-LocalGw"); ok {
			if err = d.Set("local_gw", vv); err != nil {
				return fmt.Errorf("Error reading local_gw: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_gw: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemIpipTunnelName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemIpipTunnel-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("remote_gw", flattenSystemIpipTunnelRemoteGw(o["remote-gw"], d, "remote_gw")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-gw"], "SystemIpipTunnel-RemoteGw"); ok {
			if err = d.Set("remote_gw", vv); err != nil {
				return fmt.Errorf("Error reading remote_gw: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_gw: %v", err)
		}
	}

	if err = d.Set("use_sdwan", flattenSystemIpipTunnelUseSdwan(o["use-sdwan"], d, "use_sdwan")); err != nil {
		if vv, ok := fortiAPIPatch(o["use-sdwan"], "SystemIpipTunnel-UseSdwan"); ok {
			if err = d.Set("use_sdwan", vv); err != nil {
				return fmt.Errorf("Error reading use_sdwan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading use_sdwan: %v", err)
		}
	}

	return nil
}

func flattenSystemIpipTunnelFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemIpipTunnelAutoAsicOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpipTunnelInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemIpipTunnelLocalGw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpipTunnelName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpipTunnelRemoteGw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpipTunnelUseSdwan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemIpipTunnel(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auto_asic_offload"); ok || d.HasChange("auto_asic_offload") {
		t, err := expandSystemIpipTunnelAutoAsicOffload(d, v, "auto_asic_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-asic-offload"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemIpipTunnelInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("local_gw"); ok || d.HasChange("local_gw") {
		t, err := expandSystemIpipTunnelLocalGw(d, v, "local_gw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-gw"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemIpipTunnelName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw"); ok || d.HasChange("remote_gw") {
		t, err := expandSystemIpipTunnelRemoteGw(d, v, "remote_gw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw"] = t
		}
	}

	if v, ok := d.GetOk("use_sdwan"); ok || d.HasChange("use_sdwan") {
		t, err := expandSystemIpipTunnelUseSdwan(d, v, "use_sdwan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-sdwan"] = t
		}
	}

	return &obj, nil
}
