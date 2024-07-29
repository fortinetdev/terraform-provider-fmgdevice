// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure virtual network enabler tunnel.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemVneTunnel() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemVneTunnelUpdate,
		Read:   resourceSystemVneTunnelRead,
		Update: resourceSystemVneTunnelUpdate,
		Delete: resourceSystemVneTunnelDelete,

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
			},
			"bmr_hostname": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"br": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"http_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ipv4_address": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_certificate": &schema.Schema{
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
			"update_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemVneTunnelUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemVneTunnel(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemVneTunnel resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemVneTunnel(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemVneTunnel resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemVneTunnel")

	return resourceSystemVneTunnelRead(d, m)
}

func resourceSystemVneTunnelDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemVneTunnel(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemVneTunnel resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVneTunnelRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemVneTunnel(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemVneTunnel resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemVneTunnel(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemVneTunnel resource from API: %v", err)
	}
	return nil
}

func flattenSystemVneTunnelAutoAsicOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVneTunnelBr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVneTunnelHttpUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVneTunnelInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemVneTunnelIpv4Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemVneTunnelMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVneTunnelSslCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemVneTunnelStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVneTunnelUpdateUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemVneTunnel(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auto_asic_offload", flattenSystemVneTunnelAutoAsicOffload(o["auto-asic-offload"], d, "auto_asic_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-asic-offload"], "SystemVneTunnel-AutoAsicOffload"); ok {
			if err = d.Set("auto_asic_offload", vv); err != nil {
				return fmt.Errorf("Error reading auto_asic_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_asic_offload: %v", err)
		}
	}

	if err = d.Set("br", flattenSystemVneTunnelBr(o["br"], d, "br")); err != nil {
		if vv, ok := fortiAPIPatch(o["br"], "SystemVneTunnel-Br"); ok {
			if err = d.Set("br", vv); err != nil {
				return fmt.Errorf("Error reading br: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading br: %v", err)
		}
	}

	if err = d.Set("http_username", flattenSystemVneTunnelHttpUsername(o["http-username"], d, "http_username")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-username"], "SystemVneTunnel-HttpUsername"); ok {
			if err = d.Set("http_username", vv); err != nil {
				return fmt.Errorf("Error reading http_username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_username: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemVneTunnelInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemVneTunnel-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ipv4_address", flattenSystemVneTunnelIpv4Address(o["ipv4-address"], d, "ipv4_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-address"], "SystemVneTunnel-Ipv4Address"); ok {
			if err = d.Set("ipv4_address", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_address: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemVneTunnelMode(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "SystemVneTunnel-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("ssl_certificate", flattenSystemVneTunnelSslCertificate(o["ssl-certificate"], d, "ssl_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-certificate"], "SystemVneTunnel-SslCertificate"); ok {
			if err = d.Set("ssl_certificate", vv); err != nil {
				return fmt.Errorf("Error reading ssl_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_certificate: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemVneTunnelStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemVneTunnel-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("update_url", flattenSystemVneTunnelUpdateUrl(o["update-url"], d, "update_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-url"], "SystemVneTunnel-UpdateUrl"); ok {
			if err = d.Set("update_url", vv); err != nil {
				return fmt.Errorf("Error reading update_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_url: %v", err)
		}
	}

	return nil
}

func flattenSystemVneTunnelFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemVneTunnelAutoAsicOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVneTunnelBmrHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemVneTunnelBr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVneTunnelHttpPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemVneTunnelHttpUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVneTunnelInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemVneTunnelIpv4Address(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemVneTunnelMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVneTunnelSslCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemVneTunnelStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVneTunnelUpdateUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemVneTunnel(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auto_asic_offload"); ok || d.HasChange("auto_asic_offload") {
		t, err := expandSystemVneTunnelAutoAsicOffload(d, v, "auto_asic_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-asic-offload"] = t
		}
	}

	if v, ok := d.GetOk("bmr_hostname"); ok || d.HasChange("bmr_hostname") {
		t, err := expandSystemVneTunnelBmrHostname(d, v, "bmr_hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bmr-hostname"] = t
		}
	}

	if v, ok := d.GetOk("br"); ok || d.HasChange("br") {
		t, err := expandSystemVneTunnelBr(d, v, "br")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["br"] = t
		}
	}

	if v, ok := d.GetOk("http_password"); ok || d.HasChange("http_password") {
		t, err := expandSystemVneTunnelHttpPassword(d, v, "http_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-password"] = t
		}
	}

	if v, ok := d.GetOk("http_username"); ok || d.HasChange("http_username") {
		t, err := expandSystemVneTunnelHttpUsername(d, v, "http_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-username"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemVneTunnelInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_address"); ok || d.HasChange("ipv4_address") {
		t, err := expandSystemVneTunnelIpv4Address(d, v, "ipv4_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-address"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandSystemVneTunnelMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("ssl_certificate"); ok || d.HasChange("ssl_certificate") {
		t, err := expandSystemVneTunnelSslCertificate(d, v, "ssl_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-certificate"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemVneTunnelStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("update_url"); ok || d.HasChange("update_url") {
		t, err := expandSystemVneTunnelUpdateUrl(d, v, "update_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-url"] = t
		}
	}

	return &obj, nil
}
