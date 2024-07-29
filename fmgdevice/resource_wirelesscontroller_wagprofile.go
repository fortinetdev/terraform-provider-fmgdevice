// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure wireless access gateway (WAG) profiles used for tunnels on AP.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerWagProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerWagProfileCreate,
		Read:   resourceWirelessControllerWagProfileRead,
		Update: resourceWirelessControllerWagProfileUpdate,
		Delete: resourceWirelessControllerWagProfileDelete,

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
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp_ip_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"ping_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ping_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"return_packet_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tunnel_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wag_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wag_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerWagProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerWagProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerWagProfile resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerWagProfile(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerWagProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerWagProfileRead(d, m)
}

func resourceWirelessControllerWagProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerWagProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWagProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerWagProfile(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWagProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerWagProfileRead(d, m)
}

func resourceWirelessControllerWagProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWirelessControllerWagProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerWagProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerWagProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerWagProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWagProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerWagProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWagProfile resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerWagProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWagProfileDhcpIpAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWagProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWagProfilePingInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWagProfilePingNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWagProfileReturnPacketTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWagProfileTunnelType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWagProfileWagIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWagProfileWagPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerWagProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("comment", flattenWirelessControllerWagProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "WirelessControllerWagProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("dhcp_ip_addr", flattenWirelessControllerWagProfileDhcpIpAddr(o["dhcp-ip-addr"], d, "dhcp_ip_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-ip-addr"], "WirelessControllerWagProfile-DhcpIpAddr"); ok {
			if err = d.Set("dhcp_ip_addr", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_ip_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_ip_addr: %v", err)
		}
	}

	if err = d.Set("name", flattenWirelessControllerWagProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerWagProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("ping_interval", flattenWirelessControllerWagProfilePingInterval(o["ping-interval"], d, "ping_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["ping-interval"], "WirelessControllerWagProfile-PingInterval"); ok {
			if err = d.Set("ping_interval", vv); err != nil {
				return fmt.Errorf("Error reading ping_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ping_interval: %v", err)
		}
	}

	if err = d.Set("ping_number", flattenWirelessControllerWagProfilePingNumber(o["ping-number"], d, "ping_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["ping-number"], "WirelessControllerWagProfile-PingNumber"); ok {
			if err = d.Set("ping_number", vv); err != nil {
				return fmt.Errorf("Error reading ping_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ping_number: %v", err)
		}
	}

	if err = d.Set("return_packet_timeout", flattenWirelessControllerWagProfileReturnPacketTimeout(o["return-packet-timeout"], d, "return_packet_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["return-packet-timeout"], "WirelessControllerWagProfile-ReturnPacketTimeout"); ok {
			if err = d.Set("return_packet_timeout", vv); err != nil {
				return fmt.Errorf("Error reading return_packet_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading return_packet_timeout: %v", err)
		}
	}

	if err = d.Set("tunnel_type", flattenWirelessControllerWagProfileTunnelType(o["tunnel-type"], d, "tunnel_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-type"], "WirelessControllerWagProfile-TunnelType"); ok {
			if err = d.Set("tunnel_type", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_type: %v", err)
		}
	}

	if err = d.Set("wag_ip", flattenWirelessControllerWagProfileWagIp(o["wag-ip"], d, "wag_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["wag-ip"], "WirelessControllerWagProfile-WagIp"); ok {
			if err = d.Set("wag_ip", vv); err != nil {
				return fmt.Errorf("Error reading wag_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wag_ip: %v", err)
		}
	}

	if err = d.Set("wag_port", flattenWirelessControllerWagProfileWagPort(o["wag-port"], d, "wag_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["wag-port"], "WirelessControllerWagProfile-WagPort"); ok {
			if err = d.Set("wag_port", vv); err != nil {
				return fmt.Errorf("Error reading wag_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wag_port: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerWagProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerWagProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWagProfileDhcpIpAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWagProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWagProfilePingInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWagProfilePingNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWagProfileReturnPacketTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWagProfileTunnelType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWagProfileWagIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWagProfileWagPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerWagProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandWirelessControllerWagProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_ip_addr"); ok || d.HasChange("dhcp_ip_addr") {
		t, err := expandWirelessControllerWagProfileDhcpIpAddr(d, v, "dhcp_ip_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-ip-addr"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerWagProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("ping_interval"); ok || d.HasChange("ping_interval") {
		t, err := expandWirelessControllerWagProfilePingInterval(d, v, "ping_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ping-interval"] = t
		}
	}

	if v, ok := d.GetOk("ping_number"); ok || d.HasChange("ping_number") {
		t, err := expandWirelessControllerWagProfilePingNumber(d, v, "ping_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ping-number"] = t
		}
	}

	if v, ok := d.GetOk("return_packet_timeout"); ok || d.HasChange("return_packet_timeout") {
		t, err := expandWirelessControllerWagProfileReturnPacketTimeout(d, v, "return_packet_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["return-packet-timeout"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_type"); ok || d.HasChange("tunnel_type") {
		t, err := expandWirelessControllerWagProfileTunnelType(d, v, "tunnel_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-type"] = t
		}
	}

	if v, ok := d.GetOk("wag_ip"); ok || d.HasChange("wag_ip") {
		t, err := expandWirelessControllerWagProfileWagIp(d, v, "wag_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wag-ip"] = t
		}
	}

	if v, ok := d.GetOk("wag_port"); ok || d.HasChange("wag_port") {
		t, err := expandWirelessControllerWagProfileWagPort(d, v, "wag_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wag-port"] = t
		}
	}

	return &obj, nil
}
