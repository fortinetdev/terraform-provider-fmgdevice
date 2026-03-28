// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i>

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceZtnaServiceConnector() *schema.Resource {
	return &schema.Resource{
		Create: resourceZtnaServiceConnectorCreate,
		Read:   resourceZtnaServiceConnectorRead,
		Update: resourceZtnaServiceConnectorUpdate,
		Delete: resourceZtnaServiceConnectorDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
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
			"certificate": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"connection_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"encryption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forward_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"forward_destination_cn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"forward_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"health_check_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"relay_dev_info": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"relay_user_info": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_max_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_min_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusted_ca": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"url_map": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceZtnaServiceConnectorCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

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

	obj, err := getObjectZtnaServiceConnector(d)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaServiceConnector resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadZtnaServiceConnector(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateZtnaServiceConnector(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ZtnaServiceConnector resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateZtnaServiceConnector(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ZtnaServiceConnector resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceZtnaServiceConnectorRead(d, m)
}

func resourceZtnaServiceConnectorUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

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

	obj, err := getObjectZtnaServiceConnector(d)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaServiceConnector resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateZtnaServiceConnector(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaServiceConnector resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceZtnaServiceConnectorRead(d, m)
}

func resourceZtnaServiceConnectorDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

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

	wsParams["adom"] = adomv

	err = c.DeleteZtnaServiceConnector(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ZtnaServiceConnector resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceZtnaServiceConnectorRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadZtnaServiceConnector(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ZtnaServiceConnector resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectZtnaServiceConnector(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaServiceConnector resource from API: %v", err)
	}
	return nil
}

func flattenZtnaServiceConnectorCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaServiceConnectorConnectionMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaServiceConnectorEncryption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaServiceConnectorForwardAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaServiceConnectorForwardDestinationCn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaServiceConnectorForwardPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaServiceConnectorHealthCheckInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaServiceConnectorLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaServiceConnectorName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaServiceConnectorRelayDevInfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaServiceConnectorRelayUserInfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaServiceConnectorSslMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaServiceConnectorSslMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaServiceConnectorStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaServiceConnectorTrustedCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaServiceConnectorUrlMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectZtnaServiceConnector(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("certificate", flattenZtnaServiceConnectorCertificate(o["certificate"], d, "certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["certificate"], "ZtnaServiceConnector-Certificate"); ok {
			if err = d.Set("certificate", vv); err != nil {
				return fmt.Errorf("Error reading certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("connection_mode", flattenZtnaServiceConnectorConnectionMode(o["connection-mode"], d, "connection_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["connection-mode"], "ZtnaServiceConnector-ConnectionMode"); ok {
			if err = d.Set("connection_mode", vv); err != nil {
				return fmt.Errorf("Error reading connection_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading connection_mode: %v", err)
		}
	}

	if err = d.Set("encryption", flattenZtnaServiceConnectorEncryption(o["encryption"], d, "encryption")); err != nil {
		if vv, ok := fortiAPIPatch(o["encryption"], "ZtnaServiceConnector-Encryption"); ok {
			if err = d.Set("encryption", vv); err != nil {
				return fmt.Errorf("Error reading encryption: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading encryption: %v", err)
		}
	}

	if err = d.Set("forward_address", flattenZtnaServiceConnectorForwardAddress(o["forward-address"], d, "forward_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["forward-address"], "ZtnaServiceConnector-ForwardAddress"); ok {
			if err = d.Set("forward_address", vv); err != nil {
				return fmt.Errorf("Error reading forward_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forward_address: %v", err)
		}
	}

	if err = d.Set("forward_destination_cn", flattenZtnaServiceConnectorForwardDestinationCn(o["forward-destination-cn"], d, "forward_destination_cn")); err != nil {
		if vv, ok := fortiAPIPatch(o["forward-destination-cn"], "ZtnaServiceConnector-ForwardDestinationCn"); ok {
			if err = d.Set("forward_destination_cn", vv); err != nil {
				return fmt.Errorf("Error reading forward_destination_cn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forward_destination_cn: %v", err)
		}
	}

	if err = d.Set("forward_port", flattenZtnaServiceConnectorForwardPort(o["forward-port"], d, "forward_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["forward-port"], "ZtnaServiceConnector-ForwardPort"); ok {
			if err = d.Set("forward_port", vv); err != nil {
				return fmt.Errorf("Error reading forward_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forward_port: %v", err)
		}
	}

	if err = d.Set("health_check_interval", flattenZtnaServiceConnectorHealthCheckInterval(o["health-check-interval"], d, "health_check_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["health-check-interval"], "ZtnaServiceConnector-HealthCheckInterval"); ok {
			if err = d.Set("health_check_interval", vv); err != nil {
				return fmt.Errorf("Error reading health_check_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading health_check_interval: %v", err)
		}
	}

	if err = d.Set("log", flattenZtnaServiceConnectorLog(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "ZtnaServiceConnector-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("name", flattenZtnaServiceConnectorName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ZtnaServiceConnector-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("relay_dev_info", flattenZtnaServiceConnectorRelayDevInfo(o["relay-dev-info"], d, "relay_dev_info")); err != nil {
		if vv, ok := fortiAPIPatch(o["relay-dev-info"], "ZtnaServiceConnector-RelayDevInfo"); ok {
			if err = d.Set("relay_dev_info", vv); err != nil {
				return fmt.Errorf("Error reading relay_dev_info: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading relay_dev_info: %v", err)
		}
	}

	if err = d.Set("relay_user_info", flattenZtnaServiceConnectorRelayUserInfo(o["relay-user-info"], d, "relay_user_info")); err != nil {
		if vv, ok := fortiAPIPatch(o["relay-user-info"], "ZtnaServiceConnector-RelayUserInfo"); ok {
			if err = d.Set("relay_user_info", vv); err != nil {
				return fmt.Errorf("Error reading relay_user_info: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading relay_user_info: %v", err)
		}
	}

	if err = d.Set("ssl_max_version", flattenZtnaServiceConnectorSslMaxVersion(o["ssl-max-version"], d, "ssl_max_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-max-version"], "ZtnaServiceConnector-SslMaxVersion"); ok {
			if err = d.Set("ssl_max_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_max_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_max_version: %v", err)
		}
	}

	if err = d.Set("ssl_min_version", flattenZtnaServiceConnectorSslMinVersion(o["ssl-min-version"], d, "ssl_min_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-min-version"], "ZtnaServiceConnector-SslMinVersion"); ok {
			if err = d.Set("ssl_min_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_min_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_min_version: %v", err)
		}
	}

	if err = d.Set("status", flattenZtnaServiceConnectorStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ZtnaServiceConnector-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("trusted_ca", flattenZtnaServiceConnectorTrustedCa(o["trusted-ca"], d, "trusted_ca")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusted-ca"], "ZtnaServiceConnector-TrustedCa"); ok {
			if err = d.Set("trusted_ca", vv); err != nil {
				return fmt.Errorf("Error reading trusted_ca: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusted_ca: %v", err)
		}
	}

	if err = d.Set("url_map", flattenZtnaServiceConnectorUrlMap(o["url-map"], d, "url_map")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-map"], "ZtnaServiceConnector-UrlMap"); ok {
			if err = d.Set("url_map", vv); err != nil {
				return fmt.Errorf("Error reading url_map: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_map: %v", err)
		}
	}

	return nil
}

func flattenZtnaServiceConnectorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandZtnaServiceConnectorCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaServiceConnectorConnectionMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaServiceConnectorEncryption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaServiceConnectorForwardAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaServiceConnectorForwardDestinationCn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaServiceConnectorForwardPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaServiceConnectorHealthCheckInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaServiceConnectorLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaServiceConnectorName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaServiceConnectorRelayDevInfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaServiceConnectorRelayUserInfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaServiceConnectorSslMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaServiceConnectorSslMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaServiceConnectorStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaServiceConnectorTrustedCa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaServiceConnectorUrlMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectZtnaServiceConnector(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("certificate"); ok || d.HasChange("certificate") {
		t, err := expandZtnaServiceConnectorCertificate(d, v, "certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("connection_mode"); ok || d.HasChange("connection_mode") {
		t, err := expandZtnaServiceConnectorConnectionMode(d, v, "connection_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connection-mode"] = t
		}
	}

	if v, ok := d.GetOk("encryption"); ok || d.HasChange("encryption") {
		t, err := expandZtnaServiceConnectorEncryption(d, v, "encryption")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encryption"] = t
		}
	}

	if v, ok := d.GetOk("forward_address"); ok || d.HasChange("forward_address") {
		t, err := expandZtnaServiceConnectorForwardAddress(d, v, "forward_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward-address"] = t
		}
	}

	if v, ok := d.GetOk("forward_destination_cn"); ok || d.HasChange("forward_destination_cn") {
		t, err := expandZtnaServiceConnectorForwardDestinationCn(d, v, "forward_destination_cn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward-destination-cn"] = t
		}
	}

	if v, ok := d.GetOk("forward_port"); ok || d.HasChange("forward_port") {
		t, err := expandZtnaServiceConnectorForwardPort(d, v, "forward_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward-port"] = t
		}
	}

	if v, ok := d.GetOk("health_check_interval"); ok || d.HasChange("health_check_interval") {
		t, err := expandZtnaServiceConnectorHealthCheckInterval(d, v, "health_check_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check-interval"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandZtnaServiceConnectorLog(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandZtnaServiceConnectorName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("relay_dev_info"); ok || d.HasChange("relay_dev_info") {
		t, err := expandZtnaServiceConnectorRelayDevInfo(d, v, "relay_dev_info")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["relay-dev-info"] = t
		}
	}

	if v, ok := d.GetOk("relay_user_info"); ok || d.HasChange("relay_user_info") {
		t, err := expandZtnaServiceConnectorRelayUserInfo(d, v, "relay_user_info")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["relay-user-info"] = t
		}
	}

	if v, ok := d.GetOk("ssl_max_version"); ok || d.HasChange("ssl_max_version") {
		t, err := expandZtnaServiceConnectorSslMaxVersion(d, v, "ssl_max_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-max-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_version"); ok || d.HasChange("ssl_min_version") {
		t, err := expandZtnaServiceConnectorSslMinVersion(d, v, "ssl_min_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-version"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandZtnaServiceConnectorStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("trusted_ca"); ok || d.HasChange("trusted_ca") {
		t, err := expandZtnaServiceConnectorTrustedCa(d, v, "trusted_ca")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusted-ca"] = t
		}
	}

	if v, ok := d.GetOk("url_map"); ok || d.HasChange("url_map") {
		t, err := expandZtnaServiceConnectorUrlMap(d, v, "url_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-map"] = t
		}
	}

	return &obj, nil
}
