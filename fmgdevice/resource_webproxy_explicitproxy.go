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

func resourceWebProxyExplicitProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebProxyExplicitProxyCreate,
		Read:   resourceWebProxyExplicitProxyRead,
		Update: resourceWebProxyExplicitProxyUpdate,
		Delete: resourceWebProxyExplicitProxyDelete,

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
			"detect_https_in_http_request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstport_from_incoming": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"empty_cert_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ftp_incoming_port": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ftp_over_http": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"header_proxy_agent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_connection_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_incoming_port": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"https_incoming_port": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"incoming_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"learn_dst_from_sni": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"pac_file_data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pac_file_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pac_file_server_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pac_file_server_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pac_file_through_https": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pac_file_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pref_dns_result": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"realm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"return_to_sender": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sec_default_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"secure_web_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secure_web_proxy_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"socks": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"socks_incoming_port": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssl_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_dh_bits": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unknown_http_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_agent_detect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWebProxyExplicitProxyCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebProxyExplicitProxy(d)
	if err != nil {
		return fmt.Errorf("Error creating WebProxyExplicitProxy resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWebProxyExplicitProxy(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWebProxyExplicitProxy(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WebProxyExplicitProxy resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateWebProxyExplicitProxy(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WebProxyExplicitProxy resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceWebProxyExplicitProxyRead(d, m)
}

func resourceWebProxyExplicitProxyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebProxyExplicitProxy(d)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyExplicitProxy resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWebProxyExplicitProxy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyExplicitProxy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWebProxyExplicitProxyRead(d, m)
}

func resourceWebProxyExplicitProxyDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWebProxyExplicitProxy(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebProxyExplicitProxy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyExplicitProxyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebProxyExplicitProxy(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WebProxyExplicitProxy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebProxyExplicitProxy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyExplicitProxy resource from API: %v", err)
	}
	return nil
}

func flattenWebProxyExplicitProxyDetectHttpsInHttpRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitProxyClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitProxyDnsMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitProxyDstportFromIncoming(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitProxyEmptyCertAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitProxyFtpIncomingPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyExplicitProxyFtpOverHttp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitProxyHeaderProxyAgent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitProxyHttp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitProxyHttpConnectionMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitProxyHttpIncomingPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyExplicitProxyHttpsIncomingPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyExplicitProxyIncomingIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitProxyIncomingIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitProxyInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyExplicitProxyIpv6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitProxyLearnDstFromSni(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitProxyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitProxyPacFileData(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitProxyPacFileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitProxyPacFileServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitProxyPacFileServerStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitProxyPacFileThroughHttps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitProxyPacFileUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitProxyPrefDnsResult(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitProxyRealm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitProxyReturnToSender(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitProxySecDefaultAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitProxySecureWebProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitProxySecureWebProxyCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyExplicitProxySocks(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitProxySocksIncomingPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyExplicitProxySslAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitProxySslDhBits(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitProxyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitProxyUnknownHttpVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitProxyUserAgentDetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebProxyExplicitProxy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("detect_https_in_http_request", flattenWebProxyExplicitProxyDetectHttpsInHttpRequest(o["detect-https-in-http-request"], d, "detect_https_in_http_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["detect-https-in-http-request"], "WebProxyExplicitProxy-DetectHttpsInHttpRequest"); ok {
			if err = d.Set("detect_https_in_http_request", vv); err != nil {
				return fmt.Errorf("Error reading detect_https_in_http_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading detect_https_in_http_request: %v", err)
		}
	}

	if err = d.Set("client_cert", flattenWebProxyExplicitProxyClientCert(o["client-cert"], d, "client_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-cert"], "WebProxyExplicitProxy-ClientCert"); ok {
			if err = d.Set("client_cert", vv); err != nil {
				return fmt.Errorf("Error reading client_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_cert: %v", err)
		}
	}

	if err = d.Set("dns_mode", flattenWebProxyExplicitProxyDnsMode(o["dns-mode"], d, "dns_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-mode"], "WebProxyExplicitProxy-DnsMode"); ok {
			if err = d.Set("dns_mode", vv); err != nil {
				return fmt.Errorf("Error reading dns_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_mode: %v", err)
		}
	}

	if err = d.Set("dstport_from_incoming", flattenWebProxyExplicitProxyDstportFromIncoming(o["dstport-from-incoming"], d, "dstport_from_incoming")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstport-from-incoming"], "WebProxyExplicitProxy-DstportFromIncoming"); ok {
			if err = d.Set("dstport_from_incoming", vv); err != nil {
				return fmt.Errorf("Error reading dstport_from_incoming: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstport_from_incoming: %v", err)
		}
	}

	if err = d.Set("empty_cert_action", flattenWebProxyExplicitProxyEmptyCertAction(o["empty-cert-action"], d, "empty_cert_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["empty-cert-action"], "WebProxyExplicitProxy-EmptyCertAction"); ok {
			if err = d.Set("empty_cert_action", vv); err != nil {
				return fmt.Errorf("Error reading empty_cert_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading empty_cert_action: %v", err)
		}
	}

	if err = d.Set("ftp_incoming_port", flattenWebProxyExplicitProxyFtpIncomingPort(o["ftp-incoming-port"], d, "ftp_incoming_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["ftp-incoming-port"], "WebProxyExplicitProxy-FtpIncomingPort"); ok {
			if err = d.Set("ftp_incoming_port", vv); err != nil {
				return fmt.Errorf("Error reading ftp_incoming_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ftp_incoming_port: %v", err)
		}
	}

	if err = d.Set("ftp_over_http", flattenWebProxyExplicitProxyFtpOverHttp(o["ftp-over-http"], d, "ftp_over_http")); err != nil {
		if vv, ok := fortiAPIPatch(o["ftp-over-http"], "WebProxyExplicitProxy-FtpOverHttp"); ok {
			if err = d.Set("ftp_over_http", vv); err != nil {
				return fmt.Errorf("Error reading ftp_over_http: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ftp_over_http: %v", err)
		}
	}

	if err = d.Set("header_proxy_agent", flattenWebProxyExplicitProxyHeaderProxyAgent(o["header-proxy-agent"], d, "header_proxy_agent")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-proxy-agent"], "WebProxyExplicitProxy-HeaderProxyAgent"); ok {
			if err = d.Set("header_proxy_agent", vv); err != nil {
				return fmt.Errorf("Error reading header_proxy_agent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_proxy_agent: %v", err)
		}
	}

	if err = d.Set("http", flattenWebProxyExplicitProxyHttp(o["http"], d, "http")); err != nil {
		if vv, ok := fortiAPIPatch(o["http"], "WebProxyExplicitProxy-Http"); ok {
			if err = d.Set("http", vv); err != nil {
				return fmt.Errorf("Error reading http: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http: %v", err)
		}
	}

	if err = d.Set("http_connection_mode", flattenWebProxyExplicitProxyHttpConnectionMode(o["http-connection-mode"], d, "http_connection_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-connection-mode"], "WebProxyExplicitProxy-HttpConnectionMode"); ok {
			if err = d.Set("http_connection_mode", vv); err != nil {
				return fmt.Errorf("Error reading http_connection_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_connection_mode: %v", err)
		}
	}

	if err = d.Set("http_incoming_port", flattenWebProxyExplicitProxyHttpIncomingPort(o["http-incoming-port"], d, "http_incoming_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-incoming-port"], "WebProxyExplicitProxy-HttpIncomingPort"); ok {
			if err = d.Set("http_incoming_port", vv); err != nil {
				return fmt.Errorf("Error reading http_incoming_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_incoming_port: %v", err)
		}
	}

	if err = d.Set("https_incoming_port", flattenWebProxyExplicitProxyHttpsIncomingPort(o["https-incoming-port"], d, "https_incoming_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["https-incoming-port"], "WebProxyExplicitProxy-HttpsIncomingPort"); ok {
			if err = d.Set("https_incoming_port", vv); err != nil {
				return fmt.Errorf("Error reading https_incoming_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading https_incoming_port: %v", err)
		}
	}

	if err = d.Set("incoming_ip", flattenWebProxyExplicitProxyIncomingIp(o["incoming-ip"], d, "incoming_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["incoming-ip"], "WebProxyExplicitProxy-IncomingIp"); ok {
			if err = d.Set("incoming_ip", vv); err != nil {
				return fmt.Errorf("Error reading incoming_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading incoming_ip: %v", err)
		}
	}

	if err = d.Set("incoming_ip6", flattenWebProxyExplicitProxyIncomingIp6(o["incoming-ip6"], d, "incoming_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["incoming-ip6"], "WebProxyExplicitProxy-IncomingIp6"); ok {
			if err = d.Set("incoming_ip6", vv); err != nil {
				return fmt.Errorf("Error reading incoming_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading incoming_ip6: %v", err)
		}
	}

	if err = d.Set("interface", flattenWebProxyExplicitProxyInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "WebProxyExplicitProxy-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ipv6_status", flattenWebProxyExplicitProxyIpv6Status(o["ipv6-status"], d, "ipv6_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-status"], "WebProxyExplicitProxy-Ipv6Status"); ok {
			if err = d.Set("ipv6_status", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_status: %v", err)
		}
	}

	if err = d.Set("learn_dst_from_sni", flattenWebProxyExplicitProxyLearnDstFromSni(o["learn-dst-from-sni"], d, "learn_dst_from_sni")); err != nil {
		if vv, ok := fortiAPIPatch(o["learn-dst-from-sni"], "WebProxyExplicitProxy-LearnDstFromSni"); ok {
			if err = d.Set("learn_dst_from_sni", vv); err != nil {
				return fmt.Errorf("Error reading learn_dst_from_sni: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading learn_dst_from_sni: %v", err)
		}
	}

	if err = d.Set("name", flattenWebProxyExplicitProxyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WebProxyExplicitProxy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("pac_file_data", flattenWebProxyExplicitProxyPacFileData(o["pac-file-data"], d, "pac_file_data")); err != nil {
		if vv, ok := fortiAPIPatch(o["pac-file-data"], "WebProxyExplicitProxy-PacFileData"); ok {
			if err = d.Set("pac_file_data", vv); err != nil {
				return fmt.Errorf("Error reading pac_file_data: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pac_file_data: %v", err)
		}
	}

	if err = d.Set("pac_file_name", flattenWebProxyExplicitProxyPacFileName(o["pac-file-name"], d, "pac_file_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["pac-file-name"], "WebProxyExplicitProxy-PacFileName"); ok {
			if err = d.Set("pac_file_name", vv); err != nil {
				return fmt.Errorf("Error reading pac_file_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pac_file_name: %v", err)
		}
	}

	if err = d.Set("pac_file_server_port", flattenWebProxyExplicitProxyPacFileServerPort(o["pac-file-server-port"], d, "pac_file_server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["pac-file-server-port"], "WebProxyExplicitProxy-PacFileServerPort"); ok {
			if err = d.Set("pac_file_server_port", vv); err != nil {
				return fmt.Errorf("Error reading pac_file_server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pac_file_server_port: %v", err)
		}
	}

	if err = d.Set("pac_file_server_status", flattenWebProxyExplicitProxyPacFileServerStatus(o["pac-file-server-status"], d, "pac_file_server_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["pac-file-server-status"], "WebProxyExplicitProxy-PacFileServerStatus"); ok {
			if err = d.Set("pac_file_server_status", vv); err != nil {
				return fmt.Errorf("Error reading pac_file_server_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pac_file_server_status: %v", err)
		}
	}

	if err = d.Set("pac_file_through_https", flattenWebProxyExplicitProxyPacFileThroughHttps(o["pac-file-through-https"], d, "pac_file_through_https")); err != nil {
		if vv, ok := fortiAPIPatch(o["pac-file-through-https"], "WebProxyExplicitProxy-PacFileThroughHttps"); ok {
			if err = d.Set("pac_file_through_https", vv); err != nil {
				return fmt.Errorf("Error reading pac_file_through_https: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pac_file_through_https: %v", err)
		}
	}

	if err = d.Set("pac_file_url", flattenWebProxyExplicitProxyPacFileUrl(o["pac-file-url"], d, "pac_file_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["pac-file-url"], "WebProxyExplicitProxy-PacFileUrl"); ok {
			if err = d.Set("pac_file_url", vv); err != nil {
				return fmt.Errorf("Error reading pac_file_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pac_file_url: %v", err)
		}
	}

	if err = d.Set("pref_dns_result", flattenWebProxyExplicitProxyPrefDnsResult(o["pref-dns-result"], d, "pref_dns_result")); err != nil {
		if vv, ok := fortiAPIPatch(o["pref-dns-result"], "WebProxyExplicitProxy-PrefDnsResult"); ok {
			if err = d.Set("pref_dns_result", vv); err != nil {
				return fmt.Errorf("Error reading pref_dns_result: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pref_dns_result: %v", err)
		}
	}

	if err = d.Set("realm", flattenWebProxyExplicitProxyRealm(o["realm"], d, "realm")); err != nil {
		if vv, ok := fortiAPIPatch(o["realm"], "WebProxyExplicitProxy-Realm"); ok {
			if err = d.Set("realm", vv); err != nil {
				return fmt.Errorf("Error reading realm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading realm: %v", err)
		}
	}

	if err = d.Set("return_to_sender", flattenWebProxyExplicitProxyReturnToSender(o["return-to-sender"], d, "return_to_sender")); err != nil {
		if vv, ok := fortiAPIPatch(o["return-to-sender"], "WebProxyExplicitProxy-ReturnToSender"); ok {
			if err = d.Set("return_to_sender", vv); err != nil {
				return fmt.Errorf("Error reading return_to_sender: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading return_to_sender: %v", err)
		}
	}

	if err = d.Set("sec_default_action", flattenWebProxyExplicitProxySecDefaultAction(o["sec-default-action"], d, "sec_default_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["sec-default-action"], "WebProxyExplicitProxy-SecDefaultAction"); ok {
			if err = d.Set("sec_default_action", vv); err != nil {
				return fmt.Errorf("Error reading sec_default_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sec_default_action: %v", err)
		}
	}

	if err = d.Set("secure_web_proxy", flattenWebProxyExplicitProxySecureWebProxy(o["secure-web-proxy"], d, "secure_web_proxy")); err != nil {
		if vv, ok := fortiAPIPatch(o["secure-web-proxy"], "WebProxyExplicitProxy-SecureWebProxy"); ok {
			if err = d.Set("secure_web_proxy", vv); err != nil {
				return fmt.Errorf("Error reading secure_web_proxy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secure_web_proxy: %v", err)
		}
	}

	if err = d.Set("secure_web_proxy_cert", flattenWebProxyExplicitProxySecureWebProxyCert(o["secure-web-proxy-cert"], d, "secure_web_proxy_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["secure-web-proxy-cert"], "WebProxyExplicitProxy-SecureWebProxyCert"); ok {
			if err = d.Set("secure_web_proxy_cert", vv); err != nil {
				return fmt.Errorf("Error reading secure_web_proxy_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secure_web_proxy_cert: %v", err)
		}
	}

	if err = d.Set("socks", flattenWebProxyExplicitProxySocks(o["socks"], d, "socks")); err != nil {
		if vv, ok := fortiAPIPatch(o["socks"], "WebProxyExplicitProxy-Socks"); ok {
			if err = d.Set("socks", vv); err != nil {
				return fmt.Errorf("Error reading socks: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading socks: %v", err)
		}
	}

	if err = d.Set("socks_incoming_port", flattenWebProxyExplicitProxySocksIncomingPort(o["socks-incoming-port"], d, "socks_incoming_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["socks-incoming-port"], "WebProxyExplicitProxy-SocksIncomingPort"); ok {
			if err = d.Set("socks_incoming_port", vv); err != nil {
				return fmt.Errorf("Error reading socks_incoming_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading socks_incoming_port: %v", err)
		}
	}

	if err = d.Set("ssl_algorithm", flattenWebProxyExplicitProxySslAlgorithm(o["ssl-algorithm"], d, "ssl_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-algorithm"], "WebProxyExplicitProxy-SslAlgorithm"); ok {
			if err = d.Set("ssl_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading ssl_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_algorithm: %v", err)
		}
	}

	if err = d.Set("ssl_dh_bits", flattenWebProxyExplicitProxySslDhBits(o["ssl-dh-bits"], d, "ssl_dh_bits")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-dh-bits"], "WebProxyExplicitProxy-SslDhBits"); ok {
			if err = d.Set("ssl_dh_bits", vv); err != nil {
				return fmt.Errorf("Error reading ssl_dh_bits: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_dh_bits: %v", err)
		}
	}

	if err = d.Set("status", flattenWebProxyExplicitProxyStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WebProxyExplicitProxy-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("unknown_http_version", flattenWebProxyExplicitProxyUnknownHttpVersion(o["unknown-http-version"], d, "unknown_http_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["unknown-http-version"], "WebProxyExplicitProxy-UnknownHttpVersion"); ok {
			if err = d.Set("unknown_http_version", vv); err != nil {
				return fmt.Errorf("Error reading unknown_http_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unknown_http_version: %v", err)
		}
	}

	if err = d.Set("user_agent_detect", flattenWebProxyExplicitProxyUserAgentDetect(o["user-agent-detect"], d, "user_agent_detect")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-agent-detect"], "WebProxyExplicitProxy-UserAgentDetect"); ok {
			if err = d.Set("user_agent_detect", vv); err != nil {
				return fmt.Errorf("Error reading user_agent_detect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_agent_detect: %v", err)
		}
	}

	return nil
}

func flattenWebProxyExplicitProxyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebProxyExplicitProxyDetectHttpsInHttpRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitProxyClientCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitProxyDnsMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitProxyDstportFromIncoming(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitProxyEmptyCertAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitProxyFtpIncomingPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyExplicitProxyFtpOverHttp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitProxyHeaderProxyAgent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitProxyHttp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitProxyHttpConnectionMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitProxyHttpIncomingPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyExplicitProxyHttpsIncomingPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyExplicitProxyIncomingIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitProxyIncomingIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitProxyInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyExplicitProxyIpv6Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitProxyLearnDstFromSni(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitProxyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitProxyPacFileData(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitProxyPacFileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitProxyPacFileServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitProxyPacFileServerStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitProxyPacFileThroughHttps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitProxyPacFileUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitProxyPrefDnsResult(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitProxyRealm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitProxyReturnToSender(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitProxySecDefaultAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitProxySecureWebProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitProxySecureWebProxyCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyExplicitProxySocks(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitProxySocksIncomingPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyExplicitProxySslAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitProxySslDhBits(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitProxyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitProxyUnknownHttpVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitProxyUserAgentDetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebProxyExplicitProxy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("detect_https_in_http_request"); ok || d.HasChange("detect_https_in_http_request") {
		t, err := expandWebProxyExplicitProxyDetectHttpsInHttpRequest(d, v, "detect_https_in_http_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["detect-https-in-http-request"] = t
		}
	}

	if v, ok := d.GetOk("client_cert"); ok || d.HasChange("client_cert") {
		t, err := expandWebProxyExplicitProxyClientCert(d, v, "client_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert"] = t
		}
	}

	if v, ok := d.GetOk("dns_mode"); ok || d.HasChange("dns_mode") {
		t, err := expandWebProxyExplicitProxyDnsMode(d, v, "dns_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-mode"] = t
		}
	}

	if v, ok := d.GetOk("dstport_from_incoming"); ok || d.HasChange("dstport_from_incoming") {
		t, err := expandWebProxyExplicitProxyDstportFromIncoming(d, v, "dstport_from_incoming")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstport-from-incoming"] = t
		}
	}

	if v, ok := d.GetOk("empty_cert_action"); ok || d.HasChange("empty_cert_action") {
		t, err := expandWebProxyExplicitProxyEmptyCertAction(d, v, "empty_cert_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["empty-cert-action"] = t
		}
	}

	if v, ok := d.GetOk("ftp_incoming_port"); ok || d.HasChange("ftp_incoming_port") {
		t, err := expandWebProxyExplicitProxyFtpIncomingPort(d, v, "ftp_incoming_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp-incoming-port"] = t
		}
	}

	if v, ok := d.GetOk("ftp_over_http"); ok || d.HasChange("ftp_over_http") {
		t, err := expandWebProxyExplicitProxyFtpOverHttp(d, v, "ftp_over_http")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp-over-http"] = t
		}
	}

	if v, ok := d.GetOk("header_proxy_agent"); ok || d.HasChange("header_proxy_agent") {
		t, err := expandWebProxyExplicitProxyHeaderProxyAgent(d, v, "header_proxy_agent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-proxy-agent"] = t
		}
	}

	if v, ok := d.GetOk("http"); ok || d.HasChange("http") {
		t, err := expandWebProxyExplicitProxyHttp(d, v, "http")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http"] = t
		}
	}

	if v, ok := d.GetOk("http_connection_mode"); ok || d.HasChange("http_connection_mode") {
		t, err := expandWebProxyExplicitProxyHttpConnectionMode(d, v, "http_connection_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-connection-mode"] = t
		}
	}

	if v, ok := d.GetOk("http_incoming_port"); ok || d.HasChange("http_incoming_port") {
		t, err := expandWebProxyExplicitProxyHttpIncomingPort(d, v, "http_incoming_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-incoming-port"] = t
		}
	}

	if v, ok := d.GetOk("https_incoming_port"); ok || d.HasChange("https_incoming_port") {
		t, err := expandWebProxyExplicitProxyHttpsIncomingPort(d, v, "https_incoming_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-incoming-port"] = t
		}
	}

	if v, ok := d.GetOk("incoming_ip"); ok || d.HasChange("incoming_ip") {
		t, err := expandWebProxyExplicitProxyIncomingIp(d, v, "incoming_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["incoming-ip"] = t
		}
	}

	if v, ok := d.GetOk("incoming_ip6"); ok || d.HasChange("incoming_ip6") {
		t, err := expandWebProxyExplicitProxyIncomingIp6(d, v, "incoming_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["incoming-ip6"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandWebProxyExplicitProxyInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_status"); ok || d.HasChange("ipv6_status") {
		t, err := expandWebProxyExplicitProxyIpv6Status(d, v, "ipv6_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-status"] = t
		}
	}

	if v, ok := d.GetOk("learn_dst_from_sni"); ok || d.HasChange("learn_dst_from_sni") {
		t, err := expandWebProxyExplicitProxyLearnDstFromSni(d, v, "learn_dst_from_sni")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["learn-dst-from-sni"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWebProxyExplicitProxyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("pac_file_data"); ok || d.HasChange("pac_file_data") {
		t, err := expandWebProxyExplicitProxyPacFileData(d, v, "pac_file_data")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pac-file-data"] = t
		}
	}

	if v, ok := d.GetOk("pac_file_name"); ok || d.HasChange("pac_file_name") {
		t, err := expandWebProxyExplicitProxyPacFileName(d, v, "pac_file_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pac-file-name"] = t
		}
	}

	if v, ok := d.GetOk("pac_file_server_port"); ok || d.HasChange("pac_file_server_port") {
		t, err := expandWebProxyExplicitProxyPacFileServerPort(d, v, "pac_file_server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pac-file-server-port"] = t
		}
	}

	if v, ok := d.GetOk("pac_file_server_status"); ok || d.HasChange("pac_file_server_status") {
		t, err := expandWebProxyExplicitProxyPacFileServerStatus(d, v, "pac_file_server_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pac-file-server-status"] = t
		}
	}

	if v, ok := d.GetOk("pac_file_through_https"); ok || d.HasChange("pac_file_through_https") {
		t, err := expandWebProxyExplicitProxyPacFileThroughHttps(d, v, "pac_file_through_https")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pac-file-through-https"] = t
		}
	}

	if v, ok := d.GetOk("pac_file_url"); ok || d.HasChange("pac_file_url") {
		t, err := expandWebProxyExplicitProxyPacFileUrl(d, v, "pac_file_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pac-file-url"] = t
		}
	}

	if v, ok := d.GetOk("pref_dns_result"); ok || d.HasChange("pref_dns_result") {
		t, err := expandWebProxyExplicitProxyPrefDnsResult(d, v, "pref_dns_result")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pref-dns-result"] = t
		}
	}

	if v, ok := d.GetOk("realm"); ok || d.HasChange("realm") {
		t, err := expandWebProxyExplicitProxyRealm(d, v, "realm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["realm"] = t
		}
	}

	if v, ok := d.GetOk("return_to_sender"); ok || d.HasChange("return_to_sender") {
		t, err := expandWebProxyExplicitProxyReturnToSender(d, v, "return_to_sender")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["return-to-sender"] = t
		}
	}

	if v, ok := d.GetOk("sec_default_action"); ok || d.HasChange("sec_default_action") {
		t, err := expandWebProxyExplicitProxySecDefaultAction(d, v, "sec_default_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sec-default-action"] = t
		}
	}

	if v, ok := d.GetOk("secure_web_proxy"); ok || d.HasChange("secure_web_proxy") {
		t, err := expandWebProxyExplicitProxySecureWebProxy(d, v, "secure_web_proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure-web-proxy"] = t
		}
	}

	if v, ok := d.GetOk("secure_web_proxy_cert"); ok || d.HasChange("secure_web_proxy_cert") {
		t, err := expandWebProxyExplicitProxySecureWebProxyCert(d, v, "secure_web_proxy_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure-web-proxy-cert"] = t
		}
	}

	if v, ok := d.GetOk("socks"); ok || d.HasChange("socks") {
		t, err := expandWebProxyExplicitProxySocks(d, v, "socks")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["socks"] = t
		}
	}

	if v, ok := d.GetOk("socks_incoming_port"); ok || d.HasChange("socks_incoming_port") {
		t, err := expandWebProxyExplicitProxySocksIncomingPort(d, v, "socks_incoming_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["socks-incoming-port"] = t
		}
	}

	if v, ok := d.GetOk("ssl_algorithm"); ok || d.HasChange("ssl_algorithm") {
		t, err := expandWebProxyExplicitProxySslAlgorithm(d, v, "ssl_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_dh_bits"); ok || d.HasChange("ssl_dh_bits") {
		t, err := expandWebProxyExplicitProxySslDhBits(d, v, "ssl_dh_bits")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-dh-bits"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWebProxyExplicitProxyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("unknown_http_version"); ok || d.HasChange("unknown_http_version") {
		t, err := expandWebProxyExplicitProxyUnknownHttpVersion(d, v, "unknown_http_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-http-version"] = t
		}
	}

	if v, ok := d.GetOk("user_agent_detect"); ok || d.HasChange("user_agent_detect") {
		t, err := expandWebProxyExplicitProxyUserAgentDetect(d, v, "user_agent_detect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-agent-detect"] = t
		}
	}

	return &obj, nil
}
