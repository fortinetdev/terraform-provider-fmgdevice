// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: VRRP Proxy ARP configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemInterfaceVrrpProxyArp() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemInterfaceVrrpProxyArpCreate,
		Read:   resourceSystemInterfaceVrrpProxyArpRead,
		Update: resourceSystemInterfaceVrrpProxyArpUpdate,
		Delete: resourceSystemInterfaceVrrpProxyArpDelete,

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
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vrrp": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemInterfaceVrrpProxyArpCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	var_interface := d.Get("interface").(string)
	vrrp := d.Get("vrrp").(string)
	paradict["device"] = device_name
	paradict["interface"] = var_interface
	paradict["vrrp"] = vrrp

	obj, err := getObjectSystemInterfaceVrrpProxyArp(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemInterfaceVrrpProxyArp resource while getting object: %v", err)
	}

	_, err = c.CreateSystemInterfaceVrrpProxyArp(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemInterfaceVrrpProxyArp resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemInterfaceVrrpProxyArpRead(d, m)
}

func resourceSystemInterfaceVrrpProxyArpUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	var_interface := d.Get("interface").(string)
	vrrp := d.Get("vrrp").(string)
	paradict["device"] = device_name
	paradict["interface"] = var_interface
	paradict["vrrp"] = vrrp

	obj, err := getObjectSystemInterfaceVrrpProxyArp(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterfaceVrrpProxyArp resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemInterfaceVrrpProxyArp(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterfaceVrrpProxyArp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemInterfaceVrrpProxyArpRead(d, m)
}

func resourceSystemInterfaceVrrpProxyArpDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	var_interface := d.Get("interface").(string)
	vrrp := d.Get("vrrp").(string)
	paradict["device"] = device_name
	paradict["interface"] = var_interface
	paradict["vrrp"] = vrrp

	err = c.DeleteSystemInterfaceVrrpProxyArp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemInterfaceVrrpProxyArp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemInterfaceVrrpProxyArpRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	var_interface := d.Get("interface").(string)
	vrrp := d.Get("vrrp").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
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
	if vrrp == "" {
		vrrp = importOptionChecking(m.(*FortiClient).Cfg, "vrrp")
		if vrrp == "" {
			return fmt.Errorf("Parameter vrrp is missing")
		}
		if err = d.Set("vrrp", vrrp); err != nil {
			return fmt.Errorf("Error set params vrrp: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["interface"] = var_interface
	paradict["vrrp"] = vrrp

	o, err := c.ReadSystemInterfaceVrrpProxyArp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterfaceVrrpProxyArp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemInterfaceVrrpProxyArp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterfaceVrrpProxyArp resource from API: %v", err)
	}
	return nil
}

func flattenSystemInterfaceVrrpProxyArpId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVrrpProxyArpIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemInterfaceVrrpProxyArp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenSystemInterfaceVrrpProxyArpId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemInterfaceVrrpProxyArp-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip", flattenSystemInterfaceVrrpProxyArpIp3rdl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "SystemInterfaceVrrpProxyArp-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	return nil
}

func flattenSystemInterfaceVrrpProxyArpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemInterfaceVrrpProxyArpId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVrrpProxyArpIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemInterfaceVrrpProxyArp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemInterfaceVrrpProxyArpId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandSystemInterfaceVrrpProxyArpIp3rdl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	return &obj, nil
}
