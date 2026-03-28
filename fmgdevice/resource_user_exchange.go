// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure MS Exchange server entries.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserExchange() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserExchangeCreate,
		Read:   resourceUserExchangeRead,
		Update: resourceUserExchangeUpdate,
		Delete: resourceUserExchangeDelete,

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
			"auth_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_discover_kdc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connect_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"kdc_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"server_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_min_proto_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"validate_server_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceUserExchangeCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserExchange(d)
	if err != nil {
		return fmt.Errorf("Error creating UserExchange resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadUserExchange(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateUserExchange(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating UserExchange resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateUserExchange(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating UserExchange resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceUserExchangeRead(d, m)
}

func resourceUserExchangeUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserExchange(d)
	if err != nil {
		return fmt.Errorf("Error updating UserExchange resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateUserExchange(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating UserExchange resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceUserExchangeRead(d, m)
}

func resourceUserExchangeDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteUserExchange(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting UserExchange resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserExchangeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserExchange(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading UserExchange resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserExchange(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserExchange resource from API: %v", err)
	}
	return nil
}

func flattenUserExchangeAuthLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserExchangeAuthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserExchangeAutoDiscoverKdc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserExchangeConnectProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserExchangeDomainName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserExchangeHttpAuthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserExchangeIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserExchangeKdcIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserExchangeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserExchangeServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserExchangeSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserExchangeUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserExchangeValidateServerCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectUserExchange(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auth_level", flattenUserExchangeAuthLevel(o["auth-level"], d, "auth_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-level"], "UserExchange-AuthLevel"); ok {
			if err = d.Set("auth_level", vv); err != nil {
				return fmt.Errorf("Error reading auth_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_level: %v", err)
		}
	}

	if err = d.Set("auth_type", flattenUserExchangeAuthType(o["auth-type"], d, "auth_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-type"], "UserExchange-AuthType"); ok {
			if err = d.Set("auth_type", vv); err != nil {
				return fmt.Errorf("Error reading auth_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}

	if err = d.Set("auto_discover_kdc", flattenUserExchangeAutoDiscoverKdc(o["auto-discover-kdc"], d, "auto_discover_kdc")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-discover-kdc"], "UserExchange-AutoDiscoverKdc"); ok {
			if err = d.Set("auto_discover_kdc", vv); err != nil {
				return fmt.Errorf("Error reading auto_discover_kdc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_discover_kdc: %v", err)
		}
	}

	if err = d.Set("connect_protocol", flattenUserExchangeConnectProtocol(o["connect-protocol"], d, "connect_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["connect-protocol"], "UserExchange-ConnectProtocol"); ok {
			if err = d.Set("connect_protocol", vv); err != nil {
				return fmt.Errorf("Error reading connect_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading connect_protocol: %v", err)
		}
	}

	if err = d.Set("domain_name", flattenUserExchangeDomainName(o["domain-name"], d, "domain_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain-name"], "UserExchange-DomainName"); ok {
			if err = d.Set("domain_name", vv); err != nil {
				return fmt.Errorf("Error reading domain_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain_name: %v", err)
		}
	}

	if err = d.Set("http_auth_type", flattenUserExchangeHttpAuthType(o["http-auth-type"], d, "http_auth_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-auth-type"], "UserExchange-HttpAuthType"); ok {
			if err = d.Set("http_auth_type", vv); err != nil {
				return fmt.Errorf("Error reading http_auth_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_auth_type: %v", err)
		}
	}

	if err = d.Set("ip", flattenUserExchangeIp(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "UserExchange-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("kdc_ip", flattenUserExchangeKdcIp(o["kdc-ip"], d, "kdc_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["kdc-ip"], "UserExchange-KdcIp"); ok {
			if err = d.Set("kdc_ip", vv); err != nil {
				return fmt.Errorf("Error reading kdc_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading kdc_ip: %v", err)
		}
	}

	if err = d.Set("name", flattenUserExchangeName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "UserExchange-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("server_name", flattenUserExchangeServerName(o["server-name"], d, "server_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-name"], "UserExchange-ServerName"); ok {
			if err = d.Set("server_name", vv); err != nil {
				return fmt.Errorf("Error reading server_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_name: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", flattenUserExchangeSslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-min-proto-version"], "UserExchange-SslMinProtoVersion"); ok {
			if err = d.Set("ssl_min_proto_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	if err = d.Set("username", flattenUserExchangeUsername(o["username"], d, "username")); err != nil {
		if vv, ok := fortiAPIPatch(o["username"], "UserExchange-Username"); ok {
			if err = d.Set("username", vv); err != nil {
				return fmt.Errorf("Error reading username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("validate_server_certificate", flattenUserExchangeValidateServerCertificate(o["validate-server-certificate"], d, "validate_server_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["validate-server-certificate"], "UserExchange-ValidateServerCertificate"); ok {
			if err = d.Set("validate_server_certificate", vv); err != nil {
				return fmt.Errorf("Error reading validate_server_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading validate_server_certificate: %v", err)
		}
	}

	return nil
}

func flattenUserExchangeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserExchangeAuthLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserExchangeAuthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserExchangeAutoDiscoverKdc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserExchangeConnectProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserExchangeDomainName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserExchangeHttpAuthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserExchangeIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserExchangeKdcIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserExchangeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserExchangePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserExchangeServerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserExchangeSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserExchangeUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserExchangeValidateServerCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectUserExchange(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_level"); ok || d.HasChange("auth_level") {
		t, err := expandUserExchangeAuthLevel(d, v, "auth_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-level"] = t
		}
	}

	if v, ok := d.GetOk("auth_type"); ok || d.HasChange("auth_type") {
		t, err := expandUserExchangeAuthType(d, v, "auth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-type"] = t
		}
	}

	if v, ok := d.GetOk("auto_discover_kdc"); ok || d.HasChange("auto_discover_kdc") {
		t, err := expandUserExchangeAutoDiscoverKdc(d, v, "auto_discover_kdc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-discover-kdc"] = t
		}
	}

	if v, ok := d.GetOk("connect_protocol"); ok || d.HasChange("connect_protocol") {
		t, err := expandUserExchangeConnectProtocol(d, v, "connect_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connect-protocol"] = t
		}
	}

	if v, ok := d.GetOk("domain_name"); ok || d.HasChange("domain_name") {
		t, err := expandUserExchangeDomainName(d, v, "domain_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-name"] = t
		}
	}

	if v, ok := d.GetOk("http_auth_type"); ok || d.HasChange("http_auth_type") {
		t, err := expandUserExchangeHttpAuthType(d, v, "http_auth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-auth-type"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandUserExchangeIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("kdc_ip"); ok || d.HasChange("kdc_ip") {
		t, err := expandUserExchangeKdcIp(d, v, "kdc_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["kdc-ip"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandUserExchangeName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandUserExchangePassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("server_name"); ok || d.HasChange("server_name") {
		t, err := expandUserExchangeServerName(d, v, "server_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-name"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_version"); ok || d.HasChange("ssl_min_proto_version") {
		t, err := expandUserExchangeSslMinProtoVersion(d, v, "ssl_min_proto_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-proto-version"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok || d.HasChange("username") {
		t, err := expandUserExchangeUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("validate_server_certificate"); ok || d.HasChange("validate_server_certificate") {
		t, err := expandUserExchangeValidateServerCertificate(d, v, "validate_server_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["validate-server-certificate"] = t
		}
	}

	return &obj, nil
}
