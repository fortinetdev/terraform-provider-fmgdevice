// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPv6/IPv4 in IPv6 tunnel.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemIpv6Tunnel() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemIpv6TunnelCreate,
		Read:   resourceSystemIpv6TunnelRead,
		Update: resourceSystemIpv6TunnelUpdate,
		Delete: resourceSystemIpv6TunnelDelete,

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
			"destination": &schema.Schema{
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"source": &schema.Schema{
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

func resourceSystemIpv6TunnelCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemIpv6Tunnel(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemIpv6Tunnel resource while getting object: %v", err)
	}

	_, err = c.CreateSystemIpv6Tunnel(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemIpv6Tunnel resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemIpv6TunnelRead(d, m)
}

func resourceSystemIpv6TunnelUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemIpv6Tunnel(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemIpv6Tunnel resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemIpv6Tunnel(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemIpv6Tunnel resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemIpv6TunnelRead(d, m)
}

func resourceSystemIpv6TunnelDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemIpv6Tunnel(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemIpv6Tunnel resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemIpv6TunnelRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemIpv6Tunnel(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemIpv6Tunnel resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemIpv6Tunnel(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemIpv6Tunnel resource from API: %v", err)
	}
	return nil
}

func flattenSystemIpv6TunnelAutoAsicOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpv6TunnelDestination(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpv6TunnelInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemIpv6TunnelName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpv6TunnelSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpv6TunnelUseSdwan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemIpv6Tunnel(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auto_asic_offload", flattenSystemIpv6TunnelAutoAsicOffload(o["auto-asic-offload"], d, "auto_asic_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-asic-offload"], "SystemIpv6Tunnel-AutoAsicOffload"); ok {
			if err = d.Set("auto_asic_offload", vv); err != nil {
				return fmt.Errorf("Error reading auto_asic_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_asic_offload: %v", err)
		}
	}

	if err = d.Set("destination", flattenSystemIpv6TunnelDestination(o["destination"], d, "destination")); err != nil {
		if vv, ok := fortiAPIPatch(o["destination"], "SystemIpv6Tunnel-Destination"); ok {
			if err = d.Set("destination", vv); err != nil {
				return fmt.Errorf("Error reading destination: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading destination: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemIpv6TunnelInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemIpv6Tunnel-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemIpv6TunnelName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemIpv6Tunnel-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("source", flattenSystemIpv6TunnelSource(o["source"], d, "source")); err != nil {
		if vv, ok := fortiAPIPatch(o["source"], "SystemIpv6Tunnel-Source"); ok {
			if err = d.Set("source", vv); err != nil {
				return fmt.Errorf("Error reading source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	if err = d.Set("use_sdwan", flattenSystemIpv6TunnelUseSdwan(o["use-sdwan"], d, "use_sdwan")); err != nil {
		if vv, ok := fortiAPIPatch(o["use-sdwan"], "SystemIpv6Tunnel-UseSdwan"); ok {
			if err = d.Set("use_sdwan", vv); err != nil {
				return fmt.Errorf("Error reading use_sdwan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading use_sdwan: %v", err)
		}
	}

	return nil
}

func flattenSystemIpv6TunnelFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemIpv6TunnelAutoAsicOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpv6TunnelDestination(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpv6TunnelInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemIpv6TunnelName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpv6TunnelSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpv6TunnelUseSdwan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemIpv6Tunnel(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auto_asic_offload"); ok || d.HasChange("auto_asic_offload") {
		t, err := expandSystemIpv6TunnelAutoAsicOffload(d, v, "auto_asic_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-asic-offload"] = t
		}
	}

	if v, ok := d.GetOk("destination"); ok || d.HasChange("destination") {
		t, err := expandSystemIpv6TunnelDestination(d, v, "destination")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["destination"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemIpv6TunnelInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemIpv6TunnelName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok || d.HasChange("source") {
		t, err := expandSystemIpv6TunnelSource(d, v, "source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	if v, ok := d.GetOk("use_sdwan"); ok || d.HasChange("use_sdwan") {
		t, err := expandSystemIpv6TunnelUseSdwan(d, v, "use_sdwan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-sdwan"] = t
		}
	}

	return &obj, nil
}
