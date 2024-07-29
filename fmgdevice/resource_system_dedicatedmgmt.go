// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure dedicated management.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemDedicatedMgmt() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDedicatedMgmtUpdate,
		Read:   resourceSystemDedicatedMgmtRead,
		Update: resourceSystemDedicatedMgmtUpdate,
		Delete: resourceSystemDedicatedMgmtDelete,

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
			"default_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp_end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp_netmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp_start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemDedicatedMgmtUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemDedicatedMgmt(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemDedicatedMgmt resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemDedicatedMgmt(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemDedicatedMgmt resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemDedicatedMgmt")

	return resourceSystemDedicatedMgmtRead(d, m)
}

func resourceSystemDedicatedMgmtDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	err = c.DeleteSystemDedicatedMgmt(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDedicatedMgmt resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDedicatedMgmtRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	paradict["device"] = device_name

	o, err := c.ReadSystemDedicatedMgmt(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemDedicatedMgmt resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDedicatedMgmt(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemDedicatedMgmt resource from API: %v", err)
	}
	return nil
}

func flattenSystemDedicatedMgmtDefaultGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDedicatedMgmtDhcpEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDedicatedMgmtDhcpNetmask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDedicatedMgmtDhcpServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDedicatedMgmtDhcpStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDedicatedMgmtInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDedicatedMgmtStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemDedicatedMgmt(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("default_gateway", flattenSystemDedicatedMgmtDefaultGateway(o["default-gateway"], d, "default_gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-gateway"], "SystemDedicatedMgmt-DefaultGateway"); ok {
			if err = d.Set("default_gateway", vv); err != nil {
				return fmt.Errorf("Error reading default_gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_gateway: %v", err)
		}
	}

	if err = d.Set("dhcp_end_ip", flattenSystemDedicatedMgmtDhcpEndIp(o["dhcp-end-ip"], d, "dhcp_end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-end-ip"], "SystemDedicatedMgmt-DhcpEndIp"); ok {
			if err = d.Set("dhcp_end_ip", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_end_ip: %v", err)
		}
	}

	if err = d.Set("dhcp_netmask", flattenSystemDedicatedMgmtDhcpNetmask(o["dhcp-netmask"], d, "dhcp_netmask")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-netmask"], "SystemDedicatedMgmt-DhcpNetmask"); ok {
			if err = d.Set("dhcp_netmask", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_netmask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_netmask: %v", err)
		}
	}

	if err = d.Set("dhcp_server", flattenSystemDedicatedMgmtDhcpServer(o["dhcp-server"], d, "dhcp_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-server"], "SystemDedicatedMgmt-DhcpServer"); ok {
			if err = d.Set("dhcp_server", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_server: %v", err)
		}
	}

	if err = d.Set("dhcp_start_ip", flattenSystemDedicatedMgmtDhcpStartIp(o["dhcp-start-ip"], d, "dhcp_start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-start-ip"], "SystemDedicatedMgmt-DhcpStartIp"); ok {
			if err = d.Set("dhcp_start_ip", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_start_ip: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemDedicatedMgmtInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemDedicatedMgmt-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemDedicatedMgmtStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemDedicatedMgmt-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSystemDedicatedMgmtFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemDedicatedMgmtDefaultGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDedicatedMgmtDhcpEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDedicatedMgmtDhcpNetmask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDedicatedMgmtDhcpServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDedicatedMgmtDhcpStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDedicatedMgmtInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDedicatedMgmtStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDedicatedMgmt(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("default_gateway"); ok || d.HasChange("default_gateway") {
		t, err := expandSystemDedicatedMgmtDefaultGateway(d, v, "default_gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-gateway"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_end_ip"); ok || d.HasChange("dhcp_end_ip") {
		t, err := expandSystemDedicatedMgmtDhcpEndIp(d, v, "dhcp_end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-end-ip"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_netmask"); ok || d.HasChange("dhcp_netmask") {
		t, err := expandSystemDedicatedMgmtDhcpNetmask(d, v, "dhcp_netmask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-netmask"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_server"); ok || d.HasChange("dhcp_server") {
		t, err := expandSystemDedicatedMgmtDhcpServer(d, v, "dhcp_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-server"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_start_ip"); ok || d.HasChange("dhcp_start_ip") {
		t, err := expandSystemDedicatedMgmtDhcpStartIp(d, v, "dhcp_start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-start-ip"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemDedicatedMgmtInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemDedicatedMgmtStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
