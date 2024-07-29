// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure explicit Web proxy settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebProxyExplicit() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebProxyExplicitUpdate,
		Read:   resourceWebProxyExplicitRead,
		Update: resourceWebProxyExplicitUpdate,
		Delete: resourceWebProxyExplicitDelete,

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
			"client_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"https_replacement_message": &schema.Schema{
				Type:     schema.TypeString,
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
				Computed: true,
			},
			"ipv6_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"message_upon_server_error": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"outgoing_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"outgoing_ip6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
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
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
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
			"pac_policy": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comments": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dstaddr": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"pac_file_data": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pac_file_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"policyid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"srcaddr": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"srcaddr6": &schema.Schema{
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
					},
				},
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
			"strict_guest": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trace_auth_no_rsp": &schema.Schema{
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceWebProxyExplicitUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebProxyExplicit(d)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyExplicit resource while getting object: %v", err)
	}

	_, err = c.UpdateWebProxyExplicit(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyExplicit resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WebProxyExplicit")

	return resourceWebProxyExplicitRead(d, m)
}

func resourceWebProxyExplicitDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWebProxyExplicit(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WebProxyExplicit resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyExplicitRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebProxyExplicit(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyExplicit resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebProxyExplicit(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyExplicit resource from API: %v", err)
	}
	return nil
}

func flattenWebProxyExplicitClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitEmptyCertAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitFtpIncomingPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyExplicitFtpOverHttp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitHttpConnectionMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitHttpIncomingPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyExplicitHttpsIncomingPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyExplicitHttpsReplacementMessage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitIncomingIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitIncomingIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitIpv6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitMessageUponServerError(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitOutgoingIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyExplicitOutgoingIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyExplicitPacFileData(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitPacFileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitPacFileServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyExplicitPacFileServerStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitPacFileThroughHttps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitPacFileUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitPacPolicy(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
		if _, ok := i["comments"]; ok {
			v := flattenWebProxyExplicitPacPolicyComments(i["comments"], d, pre_append)
			tmp["comments"] = fortiAPISubPartPatch(v, "WebProxyExplicit-PacPolicy-Comments")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
		if _, ok := i["dstaddr"]; ok {
			v := flattenWebProxyExplicitPacPolicyDstaddr(i["dstaddr"], d, pre_append)
			tmp["dstaddr"] = fortiAPISubPartPatch(v, "WebProxyExplicit-PacPolicy-Dstaddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pac_file_data"
		if _, ok := i["pac-file-data"]; ok {
			v := flattenWebProxyExplicitPacPolicyPacFileData(i["pac-file-data"], d, pre_append)
			tmp["pac_file_data"] = fortiAPISubPartPatch(v, "WebProxyExplicit-PacPolicy-PacFileData")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pac_file_name"
		if _, ok := i["pac-file-name"]; ok {
			v := flattenWebProxyExplicitPacPolicyPacFileName(i["pac-file-name"], d, pre_append)
			tmp["pac_file_name"] = fortiAPISubPartPatch(v, "WebProxyExplicit-PacPolicy-PacFileName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "policyid"
		if _, ok := i["policyid"]; ok {
			v := flattenWebProxyExplicitPacPolicyPolicyid(i["policyid"], d, pre_append)
			tmp["policyid"] = fortiAPISubPartPatch(v, "WebProxyExplicit-PacPolicy-Policyid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := i["srcaddr"]; ok {
			v := flattenWebProxyExplicitPacPolicySrcaddr(i["srcaddr"], d, pre_append)
			tmp["srcaddr"] = fortiAPISubPartPatch(v, "WebProxyExplicit-PacPolicy-Srcaddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr6"
		if _, ok := i["srcaddr6"]; ok {
			v := flattenWebProxyExplicitPacPolicySrcaddr6(i["srcaddr6"], d, pre_append)
			tmp["srcaddr6"] = fortiAPISubPartPatch(v, "WebProxyExplicit-PacPolicy-Srcaddr6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenWebProxyExplicitPacPolicyStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "WebProxyExplicit-PacPolicy-Status")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWebProxyExplicitPacPolicyComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitPacPolicyDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyExplicitPacPolicyPacFileData(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitPacPolicyPacFileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitPacPolicyPolicyid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitPacPolicySrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyExplicitPacPolicySrcaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyExplicitPacPolicyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitPrefDnsResult(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitRealm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitSecDefaultAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitSecureWebProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitSecureWebProxyCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyExplicitSocks(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitSocksIncomingPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyExplicitSslAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitSslDhBits(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitStrictGuest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitTraceAuthNoRsp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitUnknownHttpVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyExplicitUserAgentDetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebProxyExplicit(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("client_cert", flattenWebProxyExplicitClientCert(o["client-cert"], d, "client_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-cert"], "WebProxyExplicit-ClientCert"); ok {
			if err = d.Set("client_cert", vv); err != nil {
				return fmt.Errorf("Error reading client_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_cert: %v", err)
		}
	}

	if err = d.Set("empty_cert_action", flattenWebProxyExplicitEmptyCertAction(o["empty-cert-action"], d, "empty_cert_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["empty-cert-action"], "WebProxyExplicit-EmptyCertAction"); ok {
			if err = d.Set("empty_cert_action", vv); err != nil {
				return fmt.Errorf("Error reading empty_cert_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading empty_cert_action: %v", err)
		}
	}

	if err = d.Set("ftp_incoming_port", flattenWebProxyExplicitFtpIncomingPort(o["ftp-incoming-port"], d, "ftp_incoming_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["ftp-incoming-port"], "WebProxyExplicit-FtpIncomingPort"); ok {
			if err = d.Set("ftp_incoming_port", vv); err != nil {
				return fmt.Errorf("Error reading ftp_incoming_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ftp_incoming_port: %v", err)
		}
	}

	if err = d.Set("ftp_over_http", flattenWebProxyExplicitFtpOverHttp(o["ftp-over-http"], d, "ftp_over_http")); err != nil {
		if vv, ok := fortiAPIPatch(o["ftp-over-http"], "WebProxyExplicit-FtpOverHttp"); ok {
			if err = d.Set("ftp_over_http", vv); err != nil {
				return fmt.Errorf("Error reading ftp_over_http: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ftp_over_http: %v", err)
		}
	}

	if err = d.Set("http_connection_mode", flattenWebProxyExplicitHttpConnectionMode(o["http-connection-mode"], d, "http_connection_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-connection-mode"], "WebProxyExplicit-HttpConnectionMode"); ok {
			if err = d.Set("http_connection_mode", vv); err != nil {
				return fmt.Errorf("Error reading http_connection_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_connection_mode: %v", err)
		}
	}

	if err = d.Set("http_incoming_port", flattenWebProxyExplicitHttpIncomingPort(o["http-incoming-port"], d, "http_incoming_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-incoming-port"], "WebProxyExplicit-HttpIncomingPort"); ok {
			if err = d.Set("http_incoming_port", vv); err != nil {
				return fmt.Errorf("Error reading http_incoming_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_incoming_port: %v", err)
		}
	}

	if err = d.Set("https_incoming_port", flattenWebProxyExplicitHttpsIncomingPort(o["https-incoming-port"], d, "https_incoming_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["https-incoming-port"], "WebProxyExplicit-HttpsIncomingPort"); ok {
			if err = d.Set("https_incoming_port", vv); err != nil {
				return fmt.Errorf("Error reading https_incoming_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading https_incoming_port: %v", err)
		}
	}

	if err = d.Set("https_replacement_message", flattenWebProxyExplicitHttpsReplacementMessage(o["https-replacement-message"], d, "https_replacement_message")); err != nil {
		if vv, ok := fortiAPIPatch(o["https-replacement-message"], "WebProxyExplicit-HttpsReplacementMessage"); ok {
			if err = d.Set("https_replacement_message", vv); err != nil {
				return fmt.Errorf("Error reading https_replacement_message: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading https_replacement_message: %v", err)
		}
	}

	if err = d.Set("incoming_ip", flattenWebProxyExplicitIncomingIp(o["incoming-ip"], d, "incoming_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["incoming-ip"], "WebProxyExplicit-IncomingIp"); ok {
			if err = d.Set("incoming_ip", vv); err != nil {
				return fmt.Errorf("Error reading incoming_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading incoming_ip: %v", err)
		}
	}

	if err = d.Set("incoming_ip6", flattenWebProxyExplicitIncomingIp6(o["incoming-ip6"], d, "incoming_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["incoming-ip6"], "WebProxyExplicit-IncomingIp6"); ok {
			if err = d.Set("incoming_ip6", vv); err != nil {
				return fmt.Errorf("Error reading incoming_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading incoming_ip6: %v", err)
		}
	}

	if err = d.Set("ipv6_status", flattenWebProxyExplicitIpv6Status(o["ipv6-status"], d, "ipv6_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-status"], "WebProxyExplicit-Ipv6Status"); ok {
			if err = d.Set("ipv6_status", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_status: %v", err)
		}
	}

	if err = d.Set("message_upon_server_error", flattenWebProxyExplicitMessageUponServerError(o["message-upon-server-error"], d, "message_upon_server_error")); err != nil {
		if vv, ok := fortiAPIPatch(o["message-upon-server-error"], "WebProxyExplicit-MessageUponServerError"); ok {
			if err = d.Set("message_upon_server_error", vv); err != nil {
				return fmt.Errorf("Error reading message_upon_server_error: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading message_upon_server_error: %v", err)
		}
	}

	if err = d.Set("outgoing_ip", flattenWebProxyExplicitOutgoingIp(o["outgoing-ip"], d, "outgoing_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["outgoing-ip"], "WebProxyExplicit-OutgoingIp"); ok {
			if err = d.Set("outgoing_ip", vv); err != nil {
				return fmt.Errorf("Error reading outgoing_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outgoing_ip: %v", err)
		}
	}

	if err = d.Set("outgoing_ip6", flattenWebProxyExplicitOutgoingIp6(o["outgoing-ip6"], d, "outgoing_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["outgoing-ip6"], "WebProxyExplicit-OutgoingIp6"); ok {
			if err = d.Set("outgoing_ip6", vv); err != nil {
				return fmt.Errorf("Error reading outgoing_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outgoing_ip6: %v", err)
		}
	}

	if err = d.Set("pac_file_data", flattenWebProxyExplicitPacFileData(o["pac-file-data"], d, "pac_file_data")); err != nil {
		if vv, ok := fortiAPIPatch(o["pac-file-data"], "WebProxyExplicit-PacFileData"); ok {
			if err = d.Set("pac_file_data", vv); err != nil {
				return fmt.Errorf("Error reading pac_file_data: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pac_file_data: %v", err)
		}
	}

	if err = d.Set("pac_file_name", flattenWebProxyExplicitPacFileName(o["pac-file-name"], d, "pac_file_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["pac-file-name"], "WebProxyExplicit-PacFileName"); ok {
			if err = d.Set("pac_file_name", vv); err != nil {
				return fmt.Errorf("Error reading pac_file_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pac_file_name: %v", err)
		}
	}

	if err = d.Set("pac_file_server_port", flattenWebProxyExplicitPacFileServerPort(o["pac-file-server-port"], d, "pac_file_server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["pac-file-server-port"], "WebProxyExplicit-PacFileServerPort"); ok {
			if err = d.Set("pac_file_server_port", vv); err != nil {
				return fmt.Errorf("Error reading pac_file_server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pac_file_server_port: %v", err)
		}
	}

	if err = d.Set("pac_file_server_status", flattenWebProxyExplicitPacFileServerStatus(o["pac-file-server-status"], d, "pac_file_server_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["pac-file-server-status"], "WebProxyExplicit-PacFileServerStatus"); ok {
			if err = d.Set("pac_file_server_status", vv); err != nil {
				return fmt.Errorf("Error reading pac_file_server_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pac_file_server_status: %v", err)
		}
	}

	if err = d.Set("pac_file_through_https", flattenWebProxyExplicitPacFileThroughHttps(o["pac-file-through-https"], d, "pac_file_through_https")); err != nil {
		if vv, ok := fortiAPIPatch(o["pac-file-through-https"], "WebProxyExplicit-PacFileThroughHttps"); ok {
			if err = d.Set("pac_file_through_https", vv); err != nil {
				return fmt.Errorf("Error reading pac_file_through_https: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pac_file_through_https: %v", err)
		}
	}

	if err = d.Set("pac_file_url", flattenWebProxyExplicitPacFileUrl(o["pac-file-url"], d, "pac_file_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["pac-file-url"], "WebProxyExplicit-PacFileUrl"); ok {
			if err = d.Set("pac_file_url", vv); err != nil {
				return fmt.Errorf("Error reading pac_file_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pac_file_url: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("pac_policy", flattenWebProxyExplicitPacPolicy(o["pac-policy"], d, "pac_policy")); err != nil {
			if vv, ok := fortiAPIPatch(o["pac-policy"], "WebProxyExplicit-PacPolicy"); ok {
				if err = d.Set("pac_policy", vv); err != nil {
					return fmt.Errorf("Error reading pac_policy: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading pac_policy: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("pac_policy"); ok {
			if err = d.Set("pac_policy", flattenWebProxyExplicitPacPolicy(o["pac-policy"], d, "pac_policy")); err != nil {
				if vv, ok := fortiAPIPatch(o["pac-policy"], "WebProxyExplicit-PacPolicy"); ok {
					if err = d.Set("pac_policy", vv); err != nil {
						return fmt.Errorf("Error reading pac_policy: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading pac_policy: %v", err)
				}
			}
		}
	}

	if err = d.Set("pref_dns_result", flattenWebProxyExplicitPrefDnsResult(o["pref-dns-result"], d, "pref_dns_result")); err != nil {
		if vv, ok := fortiAPIPatch(o["pref-dns-result"], "WebProxyExplicit-PrefDnsResult"); ok {
			if err = d.Set("pref_dns_result", vv); err != nil {
				return fmt.Errorf("Error reading pref_dns_result: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pref_dns_result: %v", err)
		}
	}

	if err = d.Set("realm", flattenWebProxyExplicitRealm(o["realm"], d, "realm")); err != nil {
		if vv, ok := fortiAPIPatch(o["realm"], "WebProxyExplicit-Realm"); ok {
			if err = d.Set("realm", vv); err != nil {
				return fmt.Errorf("Error reading realm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading realm: %v", err)
		}
	}

	if err = d.Set("sec_default_action", flattenWebProxyExplicitSecDefaultAction(o["sec-default-action"], d, "sec_default_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["sec-default-action"], "WebProxyExplicit-SecDefaultAction"); ok {
			if err = d.Set("sec_default_action", vv); err != nil {
				return fmt.Errorf("Error reading sec_default_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sec_default_action: %v", err)
		}
	}

	if err = d.Set("secure_web_proxy", flattenWebProxyExplicitSecureWebProxy(o["secure-web-proxy"], d, "secure_web_proxy")); err != nil {
		if vv, ok := fortiAPIPatch(o["secure-web-proxy"], "WebProxyExplicit-SecureWebProxy"); ok {
			if err = d.Set("secure_web_proxy", vv); err != nil {
				return fmt.Errorf("Error reading secure_web_proxy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secure_web_proxy: %v", err)
		}
	}

	if err = d.Set("secure_web_proxy_cert", flattenWebProxyExplicitSecureWebProxyCert(o["secure-web-proxy-cert"], d, "secure_web_proxy_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["secure-web-proxy-cert"], "WebProxyExplicit-SecureWebProxyCert"); ok {
			if err = d.Set("secure_web_proxy_cert", vv); err != nil {
				return fmt.Errorf("Error reading secure_web_proxy_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secure_web_proxy_cert: %v", err)
		}
	}

	if err = d.Set("socks", flattenWebProxyExplicitSocks(o["socks"], d, "socks")); err != nil {
		if vv, ok := fortiAPIPatch(o["socks"], "WebProxyExplicit-Socks"); ok {
			if err = d.Set("socks", vv); err != nil {
				return fmt.Errorf("Error reading socks: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading socks: %v", err)
		}
	}

	if err = d.Set("socks_incoming_port", flattenWebProxyExplicitSocksIncomingPort(o["socks-incoming-port"], d, "socks_incoming_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["socks-incoming-port"], "WebProxyExplicit-SocksIncomingPort"); ok {
			if err = d.Set("socks_incoming_port", vv); err != nil {
				return fmt.Errorf("Error reading socks_incoming_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading socks_incoming_port: %v", err)
		}
	}

	if err = d.Set("ssl_algorithm", flattenWebProxyExplicitSslAlgorithm(o["ssl-algorithm"], d, "ssl_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-algorithm"], "WebProxyExplicit-SslAlgorithm"); ok {
			if err = d.Set("ssl_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading ssl_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_algorithm: %v", err)
		}
	}

	if err = d.Set("ssl_dh_bits", flattenWebProxyExplicitSslDhBits(o["ssl-dh-bits"], d, "ssl_dh_bits")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-dh-bits"], "WebProxyExplicit-SslDhBits"); ok {
			if err = d.Set("ssl_dh_bits", vv); err != nil {
				return fmt.Errorf("Error reading ssl_dh_bits: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_dh_bits: %v", err)
		}
	}

	if err = d.Set("status", flattenWebProxyExplicitStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WebProxyExplicit-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("strict_guest", flattenWebProxyExplicitStrictGuest(o["strict-guest"], d, "strict_guest")); err != nil {
		if vv, ok := fortiAPIPatch(o["strict-guest"], "WebProxyExplicit-StrictGuest"); ok {
			if err = d.Set("strict_guest", vv); err != nil {
				return fmt.Errorf("Error reading strict_guest: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading strict_guest: %v", err)
		}
	}

	if err = d.Set("trace_auth_no_rsp", flattenWebProxyExplicitTraceAuthNoRsp(o["trace-auth-no-rsp"], d, "trace_auth_no_rsp")); err != nil {
		if vv, ok := fortiAPIPatch(o["trace-auth-no-rsp"], "WebProxyExplicit-TraceAuthNoRsp"); ok {
			if err = d.Set("trace_auth_no_rsp", vv); err != nil {
				return fmt.Errorf("Error reading trace_auth_no_rsp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trace_auth_no_rsp: %v", err)
		}
	}

	if err = d.Set("unknown_http_version", flattenWebProxyExplicitUnknownHttpVersion(o["unknown-http-version"], d, "unknown_http_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["unknown-http-version"], "WebProxyExplicit-UnknownHttpVersion"); ok {
			if err = d.Set("unknown_http_version", vv); err != nil {
				return fmt.Errorf("Error reading unknown_http_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unknown_http_version: %v", err)
		}
	}

	if err = d.Set("user_agent_detect", flattenWebProxyExplicitUserAgentDetect(o["user-agent-detect"], d, "user_agent_detect")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-agent-detect"], "WebProxyExplicit-UserAgentDetect"); ok {
			if err = d.Set("user_agent_detect", vv); err != nil {
				return fmt.Errorf("Error reading user_agent_detect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_agent_detect: %v", err)
		}
	}

	return nil
}

func flattenWebProxyExplicitFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebProxyExplicitClientCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitEmptyCertAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitFtpIncomingPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyExplicitFtpOverHttp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitHttpConnectionMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitHttpIncomingPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyExplicitHttpsIncomingPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyExplicitHttpsReplacementMessage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitIncomingIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitIncomingIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitIpv6Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitMessageUponServerError(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitOutgoingIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyExplicitOutgoingIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyExplicitPacFileData(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitPacFileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitPacFileServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyExplicitPacFileServerStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitPacFileThroughHttps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitPacFileUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitPacPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comments"], _ = expandWebProxyExplicitPacPolicyComments(d, i["comments"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dstaddr"], _ = expandWebProxyExplicitPacPolicyDstaddr(d, i["dstaddr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pac_file_data"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pac-file-data"], _ = expandWebProxyExplicitPacPolicyPacFileData(d, i["pac_file_data"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pac_file_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pac-file-name"], _ = expandWebProxyExplicitPacPolicyPacFileName(d, i["pac_file_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "policyid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["policyid"], _ = expandWebProxyExplicitPacPolicyPolicyid(d, i["policyid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["srcaddr"], _ = expandWebProxyExplicitPacPolicySrcaddr(d, i["srcaddr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["srcaddr6"], _ = expandWebProxyExplicitPacPolicySrcaddr6(d, i["srcaddr6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandWebProxyExplicitPacPolicyStatus(d, i["status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWebProxyExplicitPacPolicyComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitPacPolicyDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyExplicitPacPolicyPacFileData(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitPacPolicyPacFileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitPacPolicyPolicyid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitPacPolicySrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyExplicitPacPolicySrcaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyExplicitPacPolicyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitPrefDnsResult(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitRealm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitSecDefaultAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitSecureWebProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitSecureWebProxyCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyExplicitSocks(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitSocksIncomingPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyExplicitSslAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitSslDhBits(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitStrictGuest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitTraceAuthNoRsp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitUnknownHttpVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitUserAgentDetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebProxyExplicit(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("client_cert"); ok || d.HasChange("client_cert") {
		t, err := expandWebProxyExplicitClientCert(d, v, "client_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert"] = t
		}
	}

	if v, ok := d.GetOk("empty_cert_action"); ok || d.HasChange("empty_cert_action") {
		t, err := expandWebProxyExplicitEmptyCertAction(d, v, "empty_cert_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["empty-cert-action"] = t
		}
	}

	if v, ok := d.GetOk("ftp_incoming_port"); ok || d.HasChange("ftp_incoming_port") {
		t, err := expandWebProxyExplicitFtpIncomingPort(d, v, "ftp_incoming_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp-incoming-port"] = t
		}
	}

	if v, ok := d.GetOk("ftp_over_http"); ok || d.HasChange("ftp_over_http") {
		t, err := expandWebProxyExplicitFtpOverHttp(d, v, "ftp_over_http")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp-over-http"] = t
		}
	}

	if v, ok := d.GetOk("http_connection_mode"); ok || d.HasChange("http_connection_mode") {
		t, err := expandWebProxyExplicitHttpConnectionMode(d, v, "http_connection_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-connection-mode"] = t
		}
	}

	if v, ok := d.GetOk("http_incoming_port"); ok || d.HasChange("http_incoming_port") {
		t, err := expandWebProxyExplicitHttpIncomingPort(d, v, "http_incoming_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-incoming-port"] = t
		}
	}

	if v, ok := d.GetOk("https_incoming_port"); ok || d.HasChange("https_incoming_port") {
		t, err := expandWebProxyExplicitHttpsIncomingPort(d, v, "https_incoming_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-incoming-port"] = t
		}
	}

	if v, ok := d.GetOk("https_replacement_message"); ok || d.HasChange("https_replacement_message") {
		t, err := expandWebProxyExplicitHttpsReplacementMessage(d, v, "https_replacement_message")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-replacement-message"] = t
		}
	}

	if v, ok := d.GetOk("incoming_ip"); ok || d.HasChange("incoming_ip") {
		t, err := expandWebProxyExplicitIncomingIp(d, v, "incoming_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["incoming-ip"] = t
		}
	}

	if v, ok := d.GetOk("incoming_ip6"); ok || d.HasChange("incoming_ip6") {
		t, err := expandWebProxyExplicitIncomingIp6(d, v, "incoming_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["incoming-ip6"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_status"); ok || d.HasChange("ipv6_status") {
		t, err := expandWebProxyExplicitIpv6Status(d, v, "ipv6_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-status"] = t
		}
	}

	if v, ok := d.GetOk("message_upon_server_error"); ok || d.HasChange("message_upon_server_error") {
		t, err := expandWebProxyExplicitMessageUponServerError(d, v, "message_upon_server_error")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["message-upon-server-error"] = t
		}
	}

	if v, ok := d.GetOk("outgoing_ip"); ok || d.HasChange("outgoing_ip") {
		t, err := expandWebProxyExplicitOutgoingIp(d, v, "outgoing_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outgoing-ip"] = t
		}
	}

	if v, ok := d.GetOk("outgoing_ip6"); ok || d.HasChange("outgoing_ip6") {
		t, err := expandWebProxyExplicitOutgoingIp6(d, v, "outgoing_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outgoing-ip6"] = t
		}
	}

	if v, ok := d.GetOk("pac_file_data"); ok || d.HasChange("pac_file_data") {
		t, err := expandWebProxyExplicitPacFileData(d, v, "pac_file_data")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pac-file-data"] = t
		}
	}

	if v, ok := d.GetOk("pac_file_name"); ok || d.HasChange("pac_file_name") {
		t, err := expandWebProxyExplicitPacFileName(d, v, "pac_file_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pac-file-name"] = t
		}
	}

	if v, ok := d.GetOk("pac_file_server_port"); ok || d.HasChange("pac_file_server_port") {
		t, err := expandWebProxyExplicitPacFileServerPort(d, v, "pac_file_server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pac-file-server-port"] = t
		}
	}

	if v, ok := d.GetOk("pac_file_server_status"); ok || d.HasChange("pac_file_server_status") {
		t, err := expandWebProxyExplicitPacFileServerStatus(d, v, "pac_file_server_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pac-file-server-status"] = t
		}
	}

	if v, ok := d.GetOk("pac_file_through_https"); ok || d.HasChange("pac_file_through_https") {
		t, err := expandWebProxyExplicitPacFileThroughHttps(d, v, "pac_file_through_https")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pac-file-through-https"] = t
		}
	}

	if v, ok := d.GetOk("pac_file_url"); ok || d.HasChange("pac_file_url") {
		t, err := expandWebProxyExplicitPacFileUrl(d, v, "pac_file_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pac-file-url"] = t
		}
	}

	if v, ok := d.GetOk("pac_policy"); ok || d.HasChange("pac_policy") {
		t, err := expandWebProxyExplicitPacPolicy(d, v, "pac_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pac-policy"] = t
		}
	}

	if v, ok := d.GetOk("pref_dns_result"); ok || d.HasChange("pref_dns_result") {
		t, err := expandWebProxyExplicitPrefDnsResult(d, v, "pref_dns_result")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pref-dns-result"] = t
		}
	}

	if v, ok := d.GetOk("realm"); ok || d.HasChange("realm") {
		t, err := expandWebProxyExplicitRealm(d, v, "realm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["realm"] = t
		}
	}

	if v, ok := d.GetOk("sec_default_action"); ok || d.HasChange("sec_default_action") {
		t, err := expandWebProxyExplicitSecDefaultAction(d, v, "sec_default_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sec-default-action"] = t
		}
	}

	if v, ok := d.GetOk("secure_web_proxy"); ok || d.HasChange("secure_web_proxy") {
		t, err := expandWebProxyExplicitSecureWebProxy(d, v, "secure_web_proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure-web-proxy"] = t
		}
	}

	if v, ok := d.GetOk("secure_web_proxy_cert"); ok || d.HasChange("secure_web_proxy_cert") {
		t, err := expandWebProxyExplicitSecureWebProxyCert(d, v, "secure_web_proxy_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure-web-proxy-cert"] = t
		}
	}

	if v, ok := d.GetOk("socks"); ok || d.HasChange("socks") {
		t, err := expandWebProxyExplicitSocks(d, v, "socks")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["socks"] = t
		}
	}

	if v, ok := d.GetOk("socks_incoming_port"); ok || d.HasChange("socks_incoming_port") {
		t, err := expandWebProxyExplicitSocksIncomingPort(d, v, "socks_incoming_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["socks-incoming-port"] = t
		}
	}

	if v, ok := d.GetOk("ssl_algorithm"); ok || d.HasChange("ssl_algorithm") {
		t, err := expandWebProxyExplicitSslAlgorithm(d, v, "ssl_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_dh_bits"); ok || d.HasChange("ssl_dh_bits") {
		t, err := expandWebProxyExplicitSslDhBits(d, v, "ssl_dh_bits")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-dh-bits"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWebProxyExplicitStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("strict_guest"); ok || d.HasChange("strict_guest") {
		t, err := expandWebProxyExplicitStrictGuest(d, v, "strict_guest")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strict-guest"] = t
		}
	}

	if v, ok := d.GetOk("trace_auth_no_rsp"); ok || d.HasChange("trace_auth_no_rsp") {
		t, err := expandWebProxyExplicitTraceAuthNoRsp(d, v, "trace_auth_no_rsp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trace-auth-no-rsp"] = t
		}
	}

	if v, ok := d.GetOk("unknown_http_version"); ok || d.HasChange("unknown_http_version") {
		t, err := expandWebProxyExplicitUnknownHttpVersion(d, v, "unknown_http_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-http-version"] = t
		}
	}

	if v, ok := d.GetOk("user_agent_detect"); ok || d.HasChange("user_agent_detect") {
		t, err := expandWebProxyExplicitUserAgentDetect(d, v, "user_agent_detect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-agent-detect"] = t
		}
	}

	return &obj, nil
}
