// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure DHCP server access list.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemInterfaceDhcpSnoopingServerList() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemInterfaceDhcpSnoopingServerListCreate,
		Read:   resourceSystemInterfaceDhcpSnoopingServerListRead,
		Update: resourceSystemInterfaceDhcpSnoopingServerListUpdate,
		Delete: resourceSystemInterfaceDhcpSnoopingServerListDelete,

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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemInterfaceDhcpSnoopingServerListCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	var_interface := d.Get("interface").(string)
	paradict["device"] = device_name
	paradict["interface"] = var_interface

	obj, err := getObjectSystemInterfaceDhcpSnoopingServerList(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemInterfaceDhcpSnoopingServerList resource while getting object: %v", err)
	}

	_, err = c.CreateSystemInterfaceDhcpSnoopingServerList(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemInterfaceDhcpSnoopingServerList resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemInterfaceDhcpSnoopingServerListRead(d, m)
}

func resourceSystemInterfaceDhcpSnoopingServerListUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["interface"] = var_interface

	obj, err := getObjectSystemInterfaceDhcpSnoopingServerList(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterfaceDhcpSnoopingServerList resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemInterfaceDhcpSnoopingServerList(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterfaceDhcpSnoopingServerList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemInterfaceDhcpSnoopingServerListRead(d, m)
}

func resourceSystemInterfaceDhcpSnoopingServerListDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["interface"] = var_interface

	err = c.DeleteSystemInterfaceDhcpSnoopingServerList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemInterfaceDhcpSnoopingServerList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemInterfaceDhcpSnoopingServerListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
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
	paradict["interface"] = var_interface

	o, err := c.ReadSystemInterfaceDhcpSnoopingServerList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterfaceDhcpSnoopingServerList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemInterfaceDhcpSnoopingServerList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterfaceDhcpSnoopingServerList resource from API: %v", err)
	}
	return nil
}

func flattenSystemInterfaceDhcpSnoopingServerListName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceDhcpSnoopingServerListServerIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemInterfaceDhcpSnoopingServerList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSystemInterfaceDhcpSnoopingServerListName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemInterfaceDhcpSnoopingServerList-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("server_ip", flattenSystemInterfaceDhcpSnoopingServerListServerIp2edl(o["server-ip"], d, "server_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-ip"], "SystemInterfaceDhcpSnoopingServerList-ServerIp"); ok {
			if err = d.Set("server_ip", vv); err != nil {
				return fmt.Errorf("Error reading server_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_ip: %v", err)
		}
	}

	return nil
}

func flattenSystemInterfaceDhcpSnoopingServerListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemInterfaceDhcpSnoopingServerListName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDhcpSnoopingServerListServerIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemInterfaceDhcpSnoopingServerList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemInterfaceDhcpSnoopingServerListName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("server_ip"); ok || d.HasChange("server_ip") {
		t, err := expandSystemInterfaceDhcpSnoopingServerListServerIp2edl(d, v, "server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-ip"] = t
		}
	}

	return &obj, nil
}
