// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Connector Remote server

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceZtnaTrafficForwardProxyReverseServiceRemoteServers() *schema.Resource {
	return &schema.Resource{
		Create: resourceZtnaTrafficForwardProxyReverseServiceRemoteServersCreate,
		Read:   resourceZtnaTrafficForwardProxyReverseServiceRemoteServersRead,
		Update: resourceZtnaTrafficForwardProxyReverseServiceRemoteServersUpdate,
		Delete: resourceZtnaTrafficForwardProxyReverseServiceRemoteServersDelete,

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
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"certificate": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"health_check_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ssl_max_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusted_server_ca": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceZtnaTrafficForwardProxyReverseServiceRemoteServersCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectZtnaTrafficForwardProxyReverseServiceRemoteServers(d)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaTrafficForwardProxyReverseServiceRemoteServers resource while getting object: %v", err)
	}

	_, err = c.CreateZtnaTrafficForwardProxyReverseServiceRemoteServers(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaTrafficForwardProxyReverseServiceRemoteServers resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceZtnaTrafficForwardProxyReverseServiceRemoteServersRead(d, m)
}

func resourceZtnaTrafficForwardProxyReverseServiceRemoteServersUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectZtnaTrafficForwardProxyReverseServiceRemoteServers(d)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaTrafficForwardProxyReverseServiceRemoteServers resource while getting object: %v", err)
	}

	_, err = c.UpdateZtnaTrafficForwardProxyReverseServiceRemoteServers(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaTrafficForwardProxyReverseServiceRemoteServers resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceZtnaTrafficForwardProxyReverseServiceRemoteServersRead(d, m)
}

func resourceZtnaTrafficForwardProxyReverseServiceRemoteServersDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteZtnaTrafficForwardProxyReverseServiceRemoteServers(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ZtnaTrafficForwardProxyReverseServiceRemoteServers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceZtnaTrafficForwardProxyReverseServiceRemoteServersRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadZtnaTrafficForwardProxyReverseServiceRemoteServers(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaTrafficForwardProxyReverseServiceRemoteServers resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectZtnaTrafficForwardProxyReverseServiceRemoteServers(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaTrafficForwardProxyReverseServiceRemoteServers resource from API: %v", err)
	}
	return nil
}

func flattenZtnaTrafficForwardProxyReverseServiceRemoteServersAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyReverseServiceRemoteServersCertificate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaTrafficForwardProxyReverseServiceRemoteServersHealthCheckInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyReverseServiceRemoteServersName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyReverseServiceRemoteServersPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyReverseServiceRemoteServersSslMaxVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyReverseServiceRemoteServersStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyReverseServiceRemoteServersTrustedServerCa2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectZtnaTrafficForwardProxyReverseServiceRemoteServers(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("address", flattenZtnaTrafficForwardProxyReverseServiceRemoteServersAddress2edl(o["address"], d, "address")); err != nil {
		if vv, ok := fortiAPIPatch(o["address"], "ZtnaTrafficForwardProxyReverseServiceRemoteServers-Address"); ok {
			if err = d.Set("address", vv); err != nil {
				return fmt.Errorf("Error reading address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("certificate", flattenZtnaTrafficForwardProxyReverseServiceRemoteServersCertificate2edl(o["certificate"], d, "certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["certificate"], "ZtnaTrafficForwardProxyReverseServiceRemoteServers-Certificate"); ok {
			if err = d.Set("certificate", vv); err != nil {
				return fmt.Errorf("Error reading certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("health_check_interval", flattenZtnaTrafficForwardProxyReverseServiceRemoteServersHealthCheckInterval2edl(o["health-check-interval"], d, "health_check_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["health-check-interval"], "ZtnaTrafficForwardProxyReverseServiceRemoteServers-HealthCheckInterval"); ok {
			if err = d.Set("health_check_interval", vv); err != nil {
				return fmt.Errorf("Error reading health_check_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading health_check_interval: %v", err)
		}
	}

	if err = d.Set("name", flattenZtnaTrafficForwardProxyReverseServiceRemoteServersName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ZtnaTrafficForwardProxyReverseServiceRemoteServers-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenZtnaTrafficForwardProxyReverseServiceRemoteServersPort2edl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ZtnaTrafficForwardProxyReverseServiceRemoteServers-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("ssl_max_version", flattenZtnaTrafficForwardProxyReverseServiceRemoteServersSslMaxVersion2edl(o["ssl-max-version"], d, "ssl_max_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-max-version"], "ZtnaTrafficForwardProxyReverseServiceRemoteServers-SslMaxVersion"); ok {
			if err = d.Set("ssl_max_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_max_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_max_version: %v", err)
		}
	}

	if err = d.Set("status", flattenZtnaTrafficForwardProxyReverseServiceRemoteServersStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ZtnaTrafficForwardProxyReverseServiceRemoteServers-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("trusted_server_ca", flattenZtnaTrafficForwardProxyReverseServiceRemoteServersTrustedServerCa2edl(o["trusted-server-ca"], d, "trusted_server_ca")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusted-server-ca"], "ZtnaTrafficForwardProxyReverseServiceRemoteServers-TrustedServerCa"); ok {
			if err = d.Set("trusted_server_ca", vv); err != nil {
				return fmt.Errorf("Error reading trusted_server_ca: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusted_server_ca: %v", err)
		}
	}

	return nil
}

func flattenZtnaTrafficForwardProxyReverseServiceRemoteServersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandZtnaTrafficForwardProxyReverseServiceRemoteServersAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyReverseServiceRemoteServersCertificate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaTrafficForwardProxyReverseServiceRemoteServersHealthCheckInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyReverseServiceRemoteServersName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyReverseServiceRemoteServersPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyReverseServiceRemoteServersSslMaxVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyReverseServiceRemoteServersStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyReverseServiceRemoteServersTrustedServerCa2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectZtnaTrafficForwardProxyReverseServiceRemoteServers(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address"); ok || d.HasChange("address") {
		t, err := expandZtnaTrafficForwardProxyReverseServiceRemoteServersAddress2edl(d, v, "address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok || d.HasChange("certificate") {
		t, err := expandZtnaTrafficForwardProxyReverseServiceRemoteServersCertificate2edl(d, v, "certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("health_check_interval"); ok || d.HasChange("health_check_interval") {
		t, err := expandZtnaTrafficForwardProxyReverseServiceRemoteServersHealthCheckInterval2edl(d, v, "health_check_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check-interval"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandZtnaTrafficForwardProxyReverseServiceRemoteServersName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandZtnaTrafficForwardProxyReverseServiceRemoteServersPort2edl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("ssl_max_version"); ok || d.HasChange("ssl_max_version") {
		t, err := expandZtnaTrafficForwardProxyReverseServiceRemoteServersSslMaxVersion2edl(d, v, "ssl_max_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-max-version"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandZtnaTrafficForwardProxyReverseServiceRemoteServersStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("trusted_server_ca"); ok || d.HasChange("trusted_server_ca") {
		t, err := expandZtnaTrafficForwardProxyReverseServiceRemoteServersTrustedServerCa2edl(d, v, "trusted_server_ca")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusted-server-ca"] = t
		}
	}

	return &obj, nil
}
