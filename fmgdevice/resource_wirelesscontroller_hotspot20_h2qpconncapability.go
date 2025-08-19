// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure connection capability.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerHotspot20H2QpConnCapability() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20H2QpConnCapabilityCreate,
		Read:   resourceWirelessControllerHotspot20H2QpConnCapabilityRead,
		Update: resourceWirelessControllerHotspot20H2QpConnCapabilityUpdate,
		Delete: resourceWirelessControllerHotspot20H2QpConnCapabilityDelete,

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
			"esp_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ftp_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icmp_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ikev2_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ikev2_xx_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"pptp_vpn_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tls_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"voip_tcp_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"voip_udp_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerHotspot20H2QpConnCapabilityCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWirelessControllerHotspot20H2QpConnCapability(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20H2QpConnCapability resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerHotspot20H2QpConnCapability(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20H2QpConnCapability resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerHotspot20H2QpConnCapabilityRead(d, m)
}

func resourceWirelessControllerHotspot20H2QpConnCapabilityUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWirelessControllerHotspot20H2QpConnCapability(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20H2QpConnCapability resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerHotspot20H2QpConnCapability(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20H2QpConnCapability resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerHotspot20H2QpConnCapabilityRead(d, m)
}

func resourceWirelessControllerHotspot20H2QpConnCapabilityDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWirelessControllerHotspot20H2QpConnCapability(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20H2QpConnCapability resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20H2QpConnCapabilityRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerHotspot20H2QpConnCapability(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20H2QpConnCapability resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20H2QpConnCapability(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20H2QpConnCapability resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20H2QpConnCapabilityEspPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpConnCapabilityFtpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpConnCapabilityHttpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpConnCapabilityIcmpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpConnCapabilityIkev2Port(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpConnCapabilityIkev2XxPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpConnCapabilityName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpConnCapabilityPptpVpnPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpConnCapabilitySshPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpConnCapabilityTlsPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpConnCapabilityVoipTcpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpConnCapabilityVoipUdpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20H2QpConnCapability(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("esp_port", flattenWirelessControllerHotspot20H2QpConnCapabilityEspPort(o["esp-port"], d, "esp_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["esp-port"], "WirelessControllerHotspot20H2QpConnCapability-EspPort"); ok {
			if err = d.Set("esp_port", vv); err != nil {
				return fmt.Errorf("Error reading esp_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading esp_port: %v", err)
		}
	}

	if err = d.Set("ftp_port", flattenWirelessControllerHotspot20H2QpConnCapabilityFtpPort(o["ftp-port"], d, "ftp_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["ftp-port"], "WirelessControllerHotspot20H2QpConnCapability-FtpPort"); ok {
			if err = d.Set("ftp_port", vv); err != nil {
				return fmt.Errorf("Error reading ftp_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ftp_port: %v", err)
		}
	}

	if err = d.Set("http_port", flattenWirelessControllerHotspot20H2QpConnCapabilityHttpPort(o["http-port"], d, "http_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-port"], "WirelessControllerHotspot20H2QpConnCapability-HttpPort"); ok {
			if err = d.Set("http_port", vv); err != nil {
				return fmt.Errorf("Error reading http_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_port: %v", err)
		}
	}

	if err = d.Set("icmp_port", flattenWirelessControllerHotspot20H2QpConnCapabilityIcmpPort(o["icmp-port"], d, "icmp_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["icmp-port"], "WirelessControllerHotspot20H2QpConnCapability-IcmpPort"); ok {
			if err = d.Set("icmp_port", vv); err != nil {
				return fmt.Errorf("Error reading icmp_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icmp_port: %v", err)
		}
	}

	if err = d.Set("ikev2_port", flattenWirelessControllerHotspot20H2QpConnCapabilityIkev2Port(o["ikev2-port"], d, "ikev2_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["ikev2-port"], "WirelessControllerHotspot20H2QpConnCapability-Ikev2Port"); ok {
			if err = d.Set("ikev2_port", vv); err != nil {
				return fmt.Errorf("Error reading ikev2_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ikev2_port: %v", err)
		}
	}

	if err = d.Set("ikev2_xx_port", flattenWirelessControllerHotspot20H2QpConnCapabilityIkev2XxPort(o["ikev2-xx-port"], d, "ikev2_xx_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["ikev2-xx-port"], "WirelessControllerHotspot20H2QpConnCapability-Ikev2XxPort"); ok {
			if err = d.Set("ikev2_xx_port", vv); err != nil {
				return fmt.Errorf("Error reading ikev2_xx_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ikev2_xx_port: %v", err)
		}
	}

	if err = d.Set("name", flattenWirelessControllerHotspot20H2QpConnCapabilityName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerHotspot20H2QpConnCapability-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("pptp_vpn_port", flattenWirelessControllerHotspot20H2QpConnCapabilityPptpVpnPort(o["pptp-vpn-port"], d, "pptp_vpn_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["pptp-vpn-port"], "WirelessControllerHotspot20H2QpConnCapability-PptpVpnPort"); ok {
			if err = d.Set("pptp_vpn_port", vv); err != nil {
				return fmt.Errorf("Error reading pptp_vpn_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pptp_vpn_port: %v", err)
		}
	}

	if err = d.Set("ssh_port", flattenWirelessControllerHotspot20H2QpConnCapabilitySshPort(o["ssh-port"], d, "ssh_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-port"], "WirelessControllerHotspot20H2QpConnCapability-SshPort"); ok {
			if err = d.Set("ssh_port", vv); err != nil {
				return fmt.Errorf("Error reading ssh_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_port: %v", err)
		}
	}

	if err = d.Set("tls_port", flattenWirelessControllerHotspot20H2QpConnCapabilityTlsPort(o["tls-port"], d, "tls_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["tls-port"], "WirelessControllerHotspot20H2QpConnCapability-TlsPort"); ok {
			if err = d.Set("tls_port", vv); err != nil {
				return fmt.Errorf("Error reading tls_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tls_port: %v", err)
		}
	}

	if err = d.Set("voip_tcp_port", flattenWirelessControllerHotspot20H2QpConnCapabilityVoipTcpPort(o["voip-tcp-port"], d, "voip_tcp_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["voip-tcp-port"], "WirelessControllerHotspot20H2QpConnCapability-VoipTcpPort"); ok {
			if err = d.Set("voip_tcp_port", vv); err != nil {
				return fmt.Errorf("Error reading voip_tcp_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading voip_tcp_port: %v", err)
		}
	}

	if err = d.Set("voip_udp_port", flattenWirelessControllerHotspot20H2QpConnCapabilityVoipUdpPort(o["voip-udp-port"], d, "voip_udp_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["voip-udp-port"], "WirelessControllerHotspot20H2QpConnCapability-VoipUdpPort"); ok {
			if err = d.Set("voip_udp_port", vv); err != nil {
				return fmt.Errorf("Error reading voip_udp_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading voip_udp_port: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20H2QpConnCapabilityFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerHotspot20H2QpConnCapabilityEspPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpConnCapabilityFtpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpConnCapabilityHttpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpConnCapabilityIcmpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpConnCapabilityIkev2Port(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpConnCapabilityIkev2XxPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpConnCapabilityName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpConnCapabilityPptpVpnPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpConnCapabilitySshPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpConnCapabilityTlsPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpConnCapabilityVoipTcpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpConnCapabilityVoipUdpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20H2QpConnCapability(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("esp_port"); ok || d.HasChange("esp_port") {
		t, err := expandWirelessControllerHotspot20H2QpConnCapabilityEspPort(d, v, "esp_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["esp-port"] = t
		}
	}

	if v, ok := d.GetOk("ftp_port"); ok || d.HasChange("ftp_port") {
		t, err := expandWirelessControllerHotspot20H2QpConnCapabilityFtpPort(d, v, "ftp_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp-port"] = t
		}
	}

	if v, ok := d.GetOk("http_port"); ok || d.HasChange("http_port") {
		t, err := expandWirelessControllerHotspot20H2QpConnCapabilityHttpPort(d, v, "http_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-port"] = t
		}
	}

	if v, ok := d.GetOk("icmp_port"); ok || d.HasChange("icmp_port") {
		t, err := expandWirelessControllerHotspot20H2QpConnCapabilityIcmpPort(d, v, "icmp_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmp-port"] = t
		}
	}

	if v, ok := d.GetOk("ikev2_port"); ok || d.HasChange("ikev2_port") {
		t, err := expandWirelessControllerHotspot20H2QpConnCapabilityIkev2Port(d, v, "ikev2_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ikev2-port"] = t
		}
	}

	if v, ok := d.GetOk("ikev2_xx_port"); ok || d.HasChange("ikev2_xx_port") {
		t, err := expandWirelessControllerHotspot20H2QpConnCapabilityIkev2XxPort(d, v, "ikev2_xx_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ikev2-xx-port"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerHotspot20H2QpConnCapabilityName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("pptp_vpn_port"); ok || d.HasChange("pptp_vpn_port") {
		t, err := expandWirelessControllerHotspot20H2QpConnCapabilityPptpVpnPort(d, v, "pptp_vpn_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pptp-vpn-port"] = t
		}
	}

	if v, ok := d.GetOk("ssh_port"); ok || d.HasChange("ssh_port") {
		t, err := expandWirelessControllerHotspot20H2QpConnCapabilitySshPort(d, v, "ssh_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-port"] = t
		}
	}

	if v, ok := d.GetOk("tls_port"); ok || d.HasChange("tls_port") {
		t, err := expandWirelessControllerHotspot20H2QpConnCapabilityTlsPort(d, v, "tls_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tls-port"] = t
		}
	}

	if v, ok := d.GetOk("voip_tcp_port"); ok || d.HasChange("voip_tcp_port") {
		t, err := expandWirelessControllerHotspot20H2QpConnCapabilityVoipTcpPort(d, v, "voip_tcp_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voip-tcp-port"] = t
		}
	}

	if v, ok := d.GetOk("voip_udp_port"); ok || d.HasChange("voip_udp_port") {
		t, err := expandWirelessControllerHotspot20H2QpConnCapabilityVoipUdpPort(d, v, "voip_udp_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voip-udp-port"] = t
		}
	}

	return &obj, nil
}
