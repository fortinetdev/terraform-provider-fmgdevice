// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device LlmProxy

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLlmProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceLlmProxyUpdate,
		Read:   resourceLlmProxyRead,
		Update: resourceLlmProxyUpdate,
		Delete: resourceLlmProxyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

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
			"hostname": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"incoming_http_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"incoming_https_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"incoming_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"incoming_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ipv6_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"srv_pool_max_concurrent_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"srv_pool_max_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"srv_pool_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ssl_certificate": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
			},
		},
	}
}

func resourceLlmProxyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLlmProxy(d, false)
	if err != nil {
		return fmt.Errorf("Error updating LlmProxy resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateLlmProxy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LlmProxy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("LlmProxy")

	return resourceLlmProxyRead(d, m)
}

func resourceLlmProxyDelete(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLlmProxy(d, true)

	if err != nil {
		return fmt.Errorf("Error updating LlmProxy resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateLlmProxy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error clearing LlmProxy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLlmProxyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLlmProxy(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading LlmProxy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLlmProxy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LlmProxy resource from API: %v", err)
	}
	return nil
}

func flattenLlmProxyHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenLlmProxyIncomingHttpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProxyIncomingHttpsPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProxyIncomingIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProxyIncomingIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProxyInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenLlmProxyIpv6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProxyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProxyProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenLlmProxySrvPoolMaxConcurrentRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProxySrvPoolMaxRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProxySrvPoolTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProxySslCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenLlmProxySslMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProxySslMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProxyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLlmProxy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("hostname", flattenLlmProxyHostname(o["hostname"], d, "hostname")); err != nil {
		if vv, ok := fortiAPIPatch(o["hostname"], "LlmProxy-Hostname"); ok {
			if err = d.Set("hostname", vv); err != nil {
				return fmt.Errorf("Error reading hostname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("incoming_http_port", flattenLlmProxyIncomingHttpPort(o["incoming-http-port"], d, "incoming_http_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["incoming-http-port"], "LlmProxy-IncomingHttpPort"); ok {
			if err = d.Set("incoming_http_port", vv); err != nil {
				return fmt.Errorf("Error reading incoming_http_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading incoming_http_port: %v", err)
		}
	}

	if err = d.Set("incoming_https_port", flattenLlmProxyIncomingHttpsPort(o["incoming-https-port"], d, "incoming_https_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["incoming-https-port"], "LlmProxy-IncomingHttpsPort"); ok {
			if err = d.Set("incoming_https_port", vv); err != nil {
				return fmt.Errorf("Error reading incoming_https_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading incoming_https_port: %v", err)
		}
	}

	if err = d.Set("incoming_ip", flattenLlmProxyIncomingIp(o["incoming-ip"], d, "incoming_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["incoming-ip"], "LlmProxy-IncomingIp"); ok {
			if err = d.Set("incoming_ip", vv); err != nil {
				return fmt.Errorf("Error reading incoming_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading incoming_ip: %v", err)
		}
	}

	if err = d.Set("incoming_ip6", flattenLlmProxyIncomingIp6(o["incoming-ip6"], d, "incoming_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["incoming-ip6"], "LlmProxy-IncomingIp6"); ok {
			if err = d.Set("incoming_ip6", vv); err != nil {
				return fmt.Errorf("Error reading incoming_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading incoming_ip6: %v", err)
		}
	}

	if err = d.Set("interface", flattenLlmProxyInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "LlmProxy-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ipv6_status", flattenLlmProxyIpv6Status(o["ipv6-status"], d, "ipv6_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-status"], "LlmProxy-Ipv6Status"); ok {
			if err = d.Set("ipv6_status", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_status: %v", err)
		}
	}

	if err = d.Set("name", flattenLlmProxyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "LlmProxy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("protocol", flattenLlmProxyProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "LlmProxy-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("srv_pool_max_concurrent_request", flattenLlmProxySrvPoolMaxConcurrentRequest(o["srv-pool-max-concurrent-request"], d, "srv_pool_max_concurrent_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["srv-pool-max-concurrent-request"], "LlmProxy-SrvPoolMaxConcurrentRequest"); ok {
			if err = d.Set("srv_pool_max_concurrent_request", vv); err != nil {
				return fmt.Errorf("Error reading srv_pool_max_concurrent_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srv_pool_max_concurrent_request: %v", err)
		}
	}

	if err = d.Set("srv_pool_max_request", flattenLlmProxySrvPoolMaxRequest(o["srv-pool-max-request"], d, "srv_pool_max_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["srv-pool-max-request"], "LlmProxy-SrvPoolMaxRequest"); ok {
			if err = d.Set("srv_pool_max_request", vv); err != nil {
				return fmt.Errorf("Error reading srv_pool_max_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srv_pool_max_request: %v", err)
		}
	}

	if err = d.Set("srv_pool_ttl", flattenLlmProxySrvPoolTtl(o["srv-pool-ttl"], d, "srv_pool_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["srv-pool-ttl"], "LlmProxy-SrvPoolTtl"); ok {
			if err = d.Set("srv_pool_ttl", vv); err != nil {
				return fmt.Errorf("Error reading srv_pool_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srv_pool_ttl: %v", err)
		}
	}

	if err = d.Set("ssl_certificate", flattenLlmProxySslCertificate(o["ssl-certificate"], d, "ssl_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-certificate"], "LlmProxy-SslCertificate"); ok {
			if err = d.Set("ssl_certificate", vv); err != nil {
				return fmt.Errorf("Error reading ssl_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_certificate: %v", err)
		}
	}

	if err = d.Set("ssl_max_version", flattenLlmProxySslMaxVersion(o["ssl-max-version"], d, "ssl_max_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-max-version"], "LlmProxy-SslMaxVersion"); ok {
			if err = d.Set("ssl_max_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_max_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_max_version: %v", err)
		}
	}

	if err = d.Set("ssl_min_version", flattenLlmProxySslMinVersion(o["ssl-min-version"], d, "ssl_min_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-min-version"], "LlmProxy-SslMinVersion"); ok {
			if err = d.Set("ssl_min_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_min_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_min_version: %v", err)
		}
	}

	if err = d.Set("status", flattenLlmProxyStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "LlmProxy-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenLlmProxyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLlmProxyHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandLlmProxyIncomingHttpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProxyIncomingHttpsPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProxyIncomingIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProxyIncomingIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProxyInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandLlmProxyIpv6Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProxyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProxyProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandLlmProxySrvPoolMaxConcurrentRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProxySrvPoolMaxRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProxySrvPoolTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProxySslCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandLlmProxySslMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProxySslMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProxyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLlmProxy(d *schema.ResourceData, bemptysontable bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("hostname"); ok || d.HasChange("hostname") {
		t, err := expandLlmProxyHostname(d, v, "hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostname"] = t
		}
	}

	if v, ok := d.GetOk("incoming_http_port"); ok || d.HasChange("incoming_http_port") {
		t, err := expandLlmProxyIncomingHttpPort(d, v, "incoming_http_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["incoming-http-port"] = t
		}
	}

	if v, ok := d.GetOk("incoming_https_port"); ok || d.HasChange("incoming_https_port") {
		t, err := expandLlmProxyIncomingHttpsPort(d, v, "incoming_https_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["incoming-https-port"] = t
		}
	}

	if v, ok := d.GetOk("incoming_ip"); ok || d.HasChange("incoming_ip") {
		t, err := expandLlmProxyIncomingIp(d, v, "incoming_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["incoming-ip"] = t
		}
	}

	if v, ok := d.GetOk("incoming_ip6"); ok || d.HasChange("incoming_ip6") {
		t, err := expandLlmProxyIncomingIp6(d, v, "incoming_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["incoming-ip6"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandLlmProxyInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_status"); ok || d.HasChange("ipv6_status") {
		t, err := expandLlmProxyIpv6Status(d, v, "ipv6_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-status"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandLlmProxyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandLlmProxyProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("srv_pool_max_concurrent_request"); ok || d.HasChange("srv_pool_max_concurrent_request") {
		t, err := expandLlmProxySrvPoolMaxConcurrentRequest(d, v, "srv_pool_max_concurrent_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srv-pool-max-concurrent-request"] = t
		}
	}

	if v, ok := d.GetOk("srv_pool_max_request"); ok || d.HasChange("srv_pool_max_request") {
		t, err := expandLlmProxySrvPoolMaxRequest(d, v, "srv_pool_max_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srv-pool-max-request"] = t
		}
	}

	if v, ok := d.GetOk("srv_pool_ttl"); ok || d.HasChange("srv_pool_ttl") {
		t, err := expandLlmProxySrvPoolTtl(d, v, "srv_pool_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srv-pool-ttl"] = t
		}
	}

	if v, ok := d.GetOk("ssl_certificate"); ok || d.HasChange("ssl_certificate") {
		t, err := expandLlmProxySslCertificate(d, v, "ssl_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-certificate"] = t
		}
	}

	if v, ok := d.GetOk("ssl_max_version"); ok || d.HasChange("ssl_max_version") {
		t, err := expandLlmProxySslMaxVersion(d, v, "ssl_max_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-max-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_version"); ok || d.HasChange("ssl_min_version") {
		t, err := expandLlmProxySslMinVersion(d, v, "ssl_min_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-version"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandLlmProxyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
