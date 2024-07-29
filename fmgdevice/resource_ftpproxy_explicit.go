// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure explicit FTP proxy settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFtpProxyExplicit() *schema.Resource {
	return &schema.Resource{
		Create: resourceFtpProxyExplicitUpdate,
		Read:   resourceFtpProxyExplicitRead,
		Update: resourceFtpProxyExplicitUpdate,
		Delete: resourceFtpProxyExplicitDelete,

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
			"incoming_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"incoming_port": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"outgoing_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sec_default_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_data_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssl_dh_bits": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFtpProxyExplicitUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFtpProxyExplicit(d)
	if err != nil {
		return fmt.Errorf("Error updating FtpProxyExplicit resource while getting object: %v", err)
	}

	_, err = c.UpdateFtpProxyExplicit(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating FtpProxyExplicit resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FtpProxyExplicit")

	return resourceFtpProxyExplicitRead(d, m)
}

func resourceFtpProxyExplicitDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFtpProxyExplicit(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting FtpProxyExplicit resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFtpProxyExplicitRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFtpProxyExplicit(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading FtpProxyExplicit resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFtpProxyExplicit(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FtpProxyExplicit resource from API: %v", err)
	}
	return nil
}

func flattenFtpProxyExplicitIncomingIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFtpProxyExplicitIncomingPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFtpProxyExplicitOutgoingIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFtpProxyExplicitSecDefaultAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFtpProxyExplicitServerDataMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFtpProxyExplicitSsl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFtpProxyExplicitSslAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFtpProxyExplicitSslCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFtpProxyExplicitSslDhBits(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFtpProxyExplicitStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFtpProxyExplicit(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("incoming_ip", flattenFtpProxyExplicitIncomingIp(o["incoming-ip"], d, "incoming_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["incoming-ip"], "FtpProxyExplicit-IncomingIp"); ok {
			if err = d.Set("incoming_ip", vv); err != nil {
				return fmt.Errorf("Error reading incoming_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading incoming_ip: %v", err)
		}
	}

	if err = d.Set("incoming_port", flattenFtpProxyExplicitIncomingPort(o["incoming-port"], d, "incoming_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["incoming-port"], "FtpProxyExplicit-IncomingPort"); ok {
			if err = d.Set("incoming_port", vv); err != nil {
				return fmt.Errorf("Error reading incoming_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading incoming_port: %v", err)
		}
	}

	if err = d.Set("outgoing_ip", flattenFtpProxyExplicitOutgoingIp(o["outgoing-ip"], d, "outgoing_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["outgoing-ip"], "FtpProxyExplicit-OutgoingIp"); ok {
			if err = d.Set("outgoing_ip", vv); err != nil {
				return fmt.Errorf("Error reading outgoing_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outgoing_ip: %v", err)
		}
	}

	if err = d.Set("sec_default_action", flattenFtpProxyExplicitSecDefaultAction(o["sec-default-action"], d, "sec_default_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["sec-default-action"], "FtpProxyExplicit-SecDefaultAction"); ok {
			if err = d.Set("sec_default_action", vv); err != nil {
				return fmt.Errorf("Error reading sec_default_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sec_default_action: %v", err)
		}
	}

	if err = d.Set("server_data_mode", flattenFtpProxyExplicitServerDataMode(o["server-data-mode"], d, "server_data_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-data-mode"], "FtpProxyExplicit-ServerDataMode"); ok {
			if err = d.Set("server_data_mode", vv); err != nil {
				return fmt.Errorf("Error reading server_data_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_data_mode: %v", err)
		}
	}

	if err = d.Set("ssl", flattenFtpProxyExplicitSsl(o["ssl"], d, "ssl")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl"], "FtpProxyExplicit-Ssl"); ok {
			if err = d.Set("ssl", vv); err != nil {
				return fmt.Errorf("Error reading ssl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl: %v", err)
		}
	}

	if err = d.Set("ssl_algorithm", flattenFtpProxyExplicitSslAlgorithm(o["ssl-algorithm"], d, "ssl_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-algorithm"], "FtpProxyExplicit-SslAlgorithm"); ok {
			if err = d.Set("ssl_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading ssl_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_algorithm: %v", err)
		}
	}

	if err = d.Set("ssl_cert", flattenFtpProxyExplicitSslCert(o["ssl-cert"], d, "ssl_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-cert"], "FtpProxyExplicit-SslCert"); ok {
			if err = d.Set("ssl_cert", vv); err != nil {
				return fmt.Errorf("Error reading ssl_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_cert: %v", err)
		}
	}

	if err = d.Set("ssl_dh_bits", flattenFtpProxyExplicitSslDhBits(o["ssl-dh-bits"], d, "ssl_dh_bits")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-dh-bits"], "FtpProxyExplicit-SslDhBits"); ok {
			if err = d.Set("ssl_dh_bits", vv); err != nil {
				return fmt.Errorf("Error reading ssl_dh_bits: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_dh_bits: %v", err)
		}
	}

	if err = d.Set("status", flattenFtpProxyExplicitStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FtpProxyExplicit-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenFtpProxyExplicitFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFtpProxyExplicitIncomingIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFtpProxyExplicitIncomingPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFtpProxyExplicitOutgoingIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFtpProxyExplicitSecDefaultAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFtpProxyExplicitServerDataMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFtpProxyExplicitSsl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFtpProxyExplicitSslAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFtpProxyExplicitSslCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFtpProxyExplicitSslDhBits(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFtpProxyExplicitStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFtpProxyExplicit(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("incoming_ip"); ok || d.HasChange("incoming_ip") {
		t, err := expandFtpProxyExplicitIncomingIp(d, v, "incoming_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["incoming-ip"] = t
		}
	}

	if v, ok := d.GetOk("incoming_port"); ok || d.HasChange("incoming_port") {
		t, err := expandFtpProxyExplicitIncomingPort(d, v, "incoming_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["incoming-port"] = t
		}
	}

	if v, ok := d.GetOk("outgoing_ip"); ok || d.HasChange("outgoing_ip") {
		t, err := expandFtpProxyExplicitOutgoingIp(d, v, "outgoing_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outgoing-ip"] = t
		}
	}

	if v, ok := d.GetOk("sec_default_action"); ok || d.HasChange("sec_default_action") {
		t, err := expandFtpProxyExplicitSecDefaultAction(d, v, "sec_default_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sec-default-action"] = t
		}
	}

	if v, ok := d.GetOk("server_data_mode"); ok || d.HasChange("server_data_mode") {
		t, err := expandFtpProxyExplicitServerDataMode(d, v, "server_data_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-data-mode"] = t
		}
	}

	if v, ok := d.GetOk("ssl"); ok || d.HasChange("ssl") {
		t, err := expandFtpProxyExplicitSsl(d, v, "ssl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl"] = t
		}
	}

	if v, ok := d.GetOk("ssl_algorithm"); ok || d.HasChange("ssl_algorithm") {
		t, err := expandFtpProxyExplicitSslAlgorithm(d, v, "ssl_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_cert"); ok || d.HasChange("ssl_cert") {
		t, err := expandFtpProxyExplicitSslCert(d, v, "ssl_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-cert"] = t
		}
	}

	if v, ok := d.GetOk("ssl_dh_bits"); ok || d.HasChange("ssl_dh_bits") {
		t, err := expandFtpProxyExplicitSslDhBits(d, v, "ssl_dh_bits")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-dh-bits"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFtpProxyExplicitStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
