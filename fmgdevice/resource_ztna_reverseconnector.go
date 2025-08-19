// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure ZTNA Reverse-Connector.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceZtnaReverseConnector() *schema.Resource {
	return &schema.Resource{
		Create: resourceZtnaReverseConnectorCreate,
		Read:   resourceZtnaReverseConnectorRead,
		Update: resourceZtnaReverseConnectorUpdate,
		Delete: resourceZtnaReverseConnectorDelete,

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

func resourceZtnaReverseConnectorCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectZtnaReverseConnector(d)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaReverseConnector resource while getting object: %v", err)
	}

	_, err = c.CreateZtnaReverseConnector(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaReverseConnector resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceZtnaReverseConnectorRead(d, m)
}

func resourceZtnaReverseConnectorUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectZtnaReverseConnector(d)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaReverseConnector resource while getting object: %v", err)
	}

	_, err = c.UpdateZtnaReverseConnector(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaReverseConnector resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceZtnaReverseConnectorRead(d, m)
}

func resourceZtnaReverseConnectorDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteZtnaReverseConnector(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ZtnaReverseConnector resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceZtnaReverseConnectorRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadZtnaReverseConnector(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaReverseConnector resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectZtnaReverseConnector(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaReverseConnector resource from API: %v", err)
	}
	return nil
}

func flattenZtnaReverseConnectorAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaReverseConnectorCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaReverseConnectorHealthCheckInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaReverseConnectorName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaReverseConnectorPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaReverseConnectorSslMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaReverseConnectorStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaReverseConnectorTrustedServerCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectZtnaReverseConnector(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("address", flattenZtnaReverseConnectorAddress(o["address"], d, "address")); err != nil {
		if vv, ok := fortiAPIPatch(o["address"], "ZtnaReverseConnector-Address"); ok {
			if err = d.Set("address", vv); err != nil {
				return fmt.Errorf("Error reading address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("certificate", flattenZtnaReverseConnectorCertificate(o["certificate"], d, "certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["certificate"], "ZtnaReverseConnector-Certificate"); ok {
			if err = d.Set("certificate", vv); err != nil {
				return fmt.Errorf("Error reading certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("health_check_interval", flattenZtnaReverseConnectorHealthCheckInterval(o["health-check-interval"], d, "health_check_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["health-check-interval"], "ZtnaReverseConnector-HealthCheckInterval"); ok {
			if err = d.Set("health_check_interval", vv); err != nil {
				return fmt.Errorf("Error reading health_check_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading health_check_interval: %v", err)
		}
	}

	if err = d.Set("name", flattenZtnaReverseConnectorName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ZtnaReverseConnector-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenZtnaReverseConnectorPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ZtnaReverseConnector-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("ssl_max_version", flattenZtnaReverseConnectorSslMaxVersion(o["ssl-max-version"], d, "ssl_max_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-max-version"], "ZtnaReverseConnector-SslMaxVersion"); ok {
			if err = d.Set("ssl_max_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_max_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_max_version: %v", err)
		}
	}

	if err = d.Set("status", flattenZtnaReverseConnectorStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ZtnaReverseConnector-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("trusted_server_ca", flattenZtnaReverseConnectorTrustedServerCa(o["trusted-server-ca"], d, "trusted_server_ca")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusted-server-ca"], "ZtnaReverseConnector-TrustedServerCa"); ok {
			if err = d.Set("trusted_server_ca", vv); err != nil {
				return fmt.Errorf("Error reading trusted_server_ca: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusted_server_ca: %v", err)
		}
	}

	return nil
}

func flattenZtnaReverseConnectorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandZtnaReverseConnectorAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaReverseConnectorCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaReverseConnectorHealthCheckInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaReverseConnectorName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaReverseConnectorPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaReverseConnectorSslMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaReverseConnectorStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaReverseConnectorTrustedServerCa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectZtnaReverseConnector(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address"); ok || d.HasChange("address") {
		t, err := expandZtnaReverseConnectorAddress(d, v, "address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok || d.HasChange("certificate") {
		t, err := expandZtnaReverseConnectorCertificate(d, v, "certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("health_check_interval"); ok || d.HasChange("health_check_interval") {
		t, err := expandZtnaReverseConnectorHealthCheckInterval(d, v, "health_check_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check-interval"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandZtnaReverseConnectorName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandZtnaReverseConnectorPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("ssl_max_version"); ok || d.HasChange("ssl_max_version") {
		t, err := expandZtnaReverseConnectorSslMaxVersion(d, v, "ssl_max_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-max-version"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandZtnaReverseConnectorStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("trusted_server_ca"); ok || d.HasChange("trusted_server_ca") {
		t, err := expandZtnaReverseConnectorTrustedServerCa(d, v, "trusted_server_ca")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusted-server-ca"] = t
		}
	}

	return &obj, nil
}
