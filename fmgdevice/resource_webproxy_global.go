// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Web proxy global settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebProxyGlobal() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebProxyGlobalUpdate,
		Read:   resourceWebProxyGlobalRead,
		Update: resourceWebProxyGlobalUpdate,
		Delete: resourceWebProxyGlobalDelete,

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
			"always_learn_client_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fast_policy_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forward_proxy_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forward_server_affinity_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http2_client_window_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http2_server_window_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ldap_user_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"learn_client_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"learn_client_ip_from_header": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"learn_client_ip_srcaddr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"learn_client_ip_srcaddr6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"log_app_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_forward_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_policy_pending": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_message_length": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_request_length": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_waf_body_cache_length": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"policy_category_deep_inspect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_transparent_cert_inspection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"request_obs_fold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_affinity_exempt_addr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"src_affinity_exempt_addr6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssl_ca_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssl_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"strict_web_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tunnel_non_http": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"unknown_http_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"webproxy_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"address_ip_rating": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"explicit_outgoing_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"explicit_outgoing_ip6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"extended_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"https_replacement_message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_http_transaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"message_upon_server_error": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"realm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"strict_guest": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"trace_auth_no_rsp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"use_dynamic_pkey": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceWebProxyGlobalUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebProxyGlobal(d)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyGlobal resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWebProxyGlobal(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyGlobal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WebProxyGlobal")

	return resourceWebProxyGlobalRead(d, m)
}

func resourceWebProxyGlobalDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWebProxyGlobal(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebProxyGlobal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyGlobalRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebProxyGlobal(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WebProxyGlobal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebProxyGlobal(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyGlobal resource from API: %v", err)
	}
	return nil
}

func flattenWebProxyGlobalAlwaysLearnClientIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyGlobalFastPolicyMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyGlobalForwardProxyAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyGlobalForwardServerAffinityTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyGlobalHttp2ClientWindowSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyGlobalHttp2ServerWindowSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyGlobalLdapUserCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyGlobalLearnClientIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyGlobalLearnClientIpFromHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyGlobalLearnClientIpSrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyGlobalLearnClientIpSrcaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyGlobalLogAppId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyGlobalLogForwardServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyGlobalLogPolicyPending(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyGlobalMaxMessageLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyGlobalMaxRequestLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyGlobalMaxWafBodyCacheLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyGlobalPolicyCategoryDeepInspect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyGlobalProxyFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyGlobalProxyTransparentCertInspection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyGlobalRequestObsFold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyGlobalSrcAffinityExemptAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyGlobalSrcAffinityExemptAddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyGlobalSslCaCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyGlobalSslCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyGlobalStrictWebCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyGlobalTunnelNonHttp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyGlobalUnknownHttpVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyGlobalWebproxyProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyGlobalAddressIpRating(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyGlobalExplicitOutgoingIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyGlobalExplicitOutgoingIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyGlobalExtendedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyGlobalHttpsReplacementMessage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyGlobalLogHttpTransaction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyGlobalMessageUponServerError(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyGlobalRealm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyGlobalStrictGuest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyGlobalTraceAuthNoRsp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyGlobalUseDynamicPkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebProxyGlobal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("always_learn_client_ip", flattenWebProxyGlobalAlwaysLearnClientIp(o["always-learn-client-ip"], d, "always_learn_client_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["always-learn-client-ip"], "WebProxyGlobal-AlwaysLearnClientIp"); ok {
			if err = d.Set("always_learn_client_ip", vv); err != nil {
				return fmt.Errorf("Error reading always_learn_client_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading always_learn_client_ip: %v", err)
		}
	}

	if err = d.Set("fast_policy_match", flattenWebProxyGlobalFastPolicyMatch(o["fast-policy-match"], d, "fast_policy_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["fast-policy-match"], "WebProxyGlobal-FastPolicyMatch"); ok {
			if err = d.Set("fast_policy_match", vv); err != nil {
				return fmt.Errorf("Error reading fast_policy_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fast_policy_match: %v", err)
		}
	}

	if err = d.Set("forward_proxy_auth", flattenWebProxyGlobalForwardProxyAuth(o["forward-proxy-auth"], d, "forward_proxy_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["forward-proxy-auth"], "WebProxyGlobal-ForwardProxyAuth"); ok {
			if err = d.Set("forward_proxy_auth", vv); err != nil {
				return fmt.Errorf("Error reading forward_proxy_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forward_proxy_auth: %v", err)
		}
	}

	if err = d.Set("forward_server_affinity_timeout", flattenWebProxyGlobalForwardServerAffinityTimeout(o["forward-server-affinity-timeout"], d, "forward_server_affinity_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["forward-server-affinity-timeout"], "WebProxyGlobal-ForwardServerAffinityTimeout"); ok {
			if err = d.Set("forward_server_affinity_timeout", vv); err != nil {
				return fmt.Errorf("Error reading forward_server_affinity_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forward_server_affinity_timeout: %v", err)
		}
	}

	if err = d.Set("http2_client_window_size", flattenWebProxyGlobalHttp2ClientWindowSize(o["http2-client-window-size"], d, "http2_client_window_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["http2-client-window-size"], "WebProxyGlobal-Http2ClientWindowSize"); ok {
			if err = d.Set("http2_client_window_size", vv); err != nil {
				return fmt.Errorf("Error reading http2_client_window_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http2_client_window_size: %v", err)
		}
	}

	if err = d.Set("http2_server_window_size", flattenWebProxyGlobalHttp2ServerWindowSize(o["http2-server-window-size"], d, "http2_server_window_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["http2-server-window-size"], "WebProxyGlobal-Http2ServerWindowSize"); ok {
			if err = d.Set("http2_server_window_size", vv); err != nil {
				return fmt.Errorf("Error reading http2_server_window_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http2_server_window_size: %v", err)
		}
	}

	if err = d.Set("ldap_user_cache", flattenWebProxyGlobalLdapUserCache(o["ldap-user-cache"], d, "ldap_user_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-user-cache"], "WebProxyGlobal-LdapUserCache"); ok {
			if err = d.Set("ldap_user_cache", vv); err != nil {
				return fmt.Errorf("Error reading ldap_user_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_user_cache: %v", err)
		}
	}

	if err = d.Set("learn_client_ip", flattenWebProxyGlobalLearnClientIp(o["learn-client-ip"], d, "learn_client_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["learn-client-ip"], "WebProxyGlobal-LearnClientIp"); ok {
			if err = d.Set("learn_client_ip", vv); err != nil {
				return fmt.Errorf("Error reading learn_client_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading learn_client_ip: %v", err)
		}
	}

	if err = d.Set("learn_client_ip_from_header", flattenWebProxyGlobalLearnClientIpFromHeader(o["learn-client-ip-from-header"], d, "learn_client_ip_from_header")); err != nil {
		if vv, ok := fortiAPIPatch(o["learn-client-ip-from-header"], "WebProxyGlobal-LearnClientIpFromHeader"); ok {
			if err = d.Set("learn_client_ip_from_header", vv); err != nil {
				return fmt.Errorf("Error reading learn_client_ip_from_header: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading learn_client_ip_from_header: %v", err)
		}
	}

	if err = d.Set("learn_client_ip_srcaddr", flattenWebProxyGlobalLearnClientIpSrcaddr(o["learn-client-ip-srcaddr"], d, "learn_client_ip_srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["learn-client-ip-srcaddr"], "WebProxyGlobal-LearnClientIpSrcaddr"); ok {
			if err = d.Set("learn_client_ip_srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading learn_client_ip_srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading learn_client_ip_srcaddr: %v", err)
		}
	}

	if err = d.Set("learn_client_ip_srcaddr6", flattenWebProxyGlobalLearnClientIpSrcaddr6(o["learn-client-ip-srcaddr6"], d, "learn_client_ip_srcaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["learn-client-ip-srcaddr6"], "WebProxyGlobal-LearnClientIpSrcaddr6"); ok {
			if err = d.Set("learn_client_ip_srcaddr6", vv); err != nil {
				return fmt.Errorf("Error reading learn_client_ip_srcaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading learn_client_ip_srcaddr6: %v", err)
		}
	}

	if err = d.Set("log_app_id", flattenWebProxyGlobalLogAppId(o["log-app-id"], d, "log_app_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-app-id"], "WebProxyGlobal-LogAppId"); ok {
			if err = d.Set("log_app_id", vv); err != nil {
				return fmt.Errorf("Error reading log_app_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_app_id: %v", err)
		}
	}

	if err = d.Set("log_forward_server", flattenWebProxyGlobalLogForwardServer(o["log-forward-server"], d, "log_forward_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-forward-server"], "WebProxyGlobal-LogForwardServer"); ok {
			if err = d.Set("log_forward_server", vv); err != nil {
				return fmt.Errorf("Error reading log_forward_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_forward_server: %v", err)
		}
	}

	if err = d.Set("log_policy_pending", flattenWebProxyGlobalLogPolicyPending(o["log-policy-pending"], d, "log_policy_pending")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-policy-pending"], "WebProxyGlobal-LogPolicyPending"); ok {
			if err = d.Set("log_policy_pending", vv); err != nil {
				return fmt.Errorf("Error reading log_policy_pending: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_policy_pending: %v", err)
		}
	}

	if err = d.Set("max_message_length", flattenWebProxyGlobalMaxMessageLength(o["max-message-length"], d, "max_message_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-message-length"], "WebProxyGlobal-MaxMessageLength"); ok {
			if err = d.Set("max_message_length", vv); err != nil {
				return fmt.Errorf("Error reading max_message_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_message_length: %v", err)
		}
	}

	if err = d.Set("max_request_length", flattenWebProxyGlobalMaxRequestLength(o["max-request-length"], d, "max_request_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-request-length"], "WebProxyGlobal-MaxRequestLength"); ok {
			if err = d.Set("max_request_length", vv); err != nil {
				return fmt.Errorf("Error reading max_request_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_request_length: %v", err)
		}
	}

	if err = d.Set("max_waf_body_cache_length", flattenWebProxyGlobalMaxWafBodyCacheLength(o["max-waf-body-cache-length"], d, "max_waf_body_cache_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-waf-body-cache-length"], "WebProxyGlobal-MaxWafBodyCacheLength"); ok {
			if err = d.Set("max_waf_body_cache_length", vv); err != nil {
				return fmt.Errorf("Error reading max_waf_body_cache_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_waf_body_cache_length: %v", err)
		}
	}

	if err = d.Set("policy_category_deep_inspect", flattenWebProxyGlobalPolicyCategoryDeepInspect(o["policy-category-deep-inspect"], d, "policy_category_deep_inspect")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-category-deep-inspect"], "WebProxyGlobal-PolicyCategoryDeepInspect"); ok {
			if err = d.Set("policy_category_deep_inspect", vv); err != nil {
				return fmt.Errorf("Error reading policy_category_deep_inspect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_category_deep_inspect: %v", err)
		}
	}

	if err = d.Set("proxy_fqdn", flattenWebProxyGlobalProxyFqdn(o["proxy-fqdn"], d, "proxy_fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-fqdn"], "WebProxyGlobal-ProxyFqdn"); ok {
			if err = d.Set("proxy_fqdn", vv); err != nil {
				return fmt.Errorf("Error reading proxy_fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_fqdn: %v", err)
		}
	}

	if err = d.Set("proxy_transparent_cert_inspection", flattenWebProxyGlobalProxyTransparentCertInspection(o["proxy-transparent-cert-inspection"], d, "proxy_transparent_cert_inspection")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-transparent-cert-inspection"], "WebProxyGlobal-ProxyTransparentCertInspection"); ok {
			if err = d.Set("proxy_transparent_cert_inspection", vv); err != nil {
				return fmt.Errorf("Error reading proxy_transparent_cert_inspection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_transparent_cert_inspection: %v", err)
		}
	}

	if err = d.Set("request_obs_fold", flattenWebProxyGlobalRequestObsFold(o["request-obs-fold"], d, "request_obs_fold")); err != nil {
		if vv, ok := fortiAPIPatch(o["request-obs-fold"], "WebProxyGlobal-RequestObsFold"); ok {
			if err = d.Set("request_obs_fold", vv); err != nil {
				return fmt.Errorf("Error reading request_obs_fold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading request_obs_fold: %v", err)
		}
	}

	if err = d.Set("src_affinity_exempt_addr", flattenWebProxyGlobalSrcAffinityExemptAddr(o["src-affinity-exempt-addr"], d, "src_affinity_exempt_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-affinity-exempt-addr"], "WebProxyGlobal-SrcAffinityExemptAddr"); ok {
			if err = d.Set("src_affinity_exempt_addr", vv); err != nil {
				return fmt.Errorf("Error reading src_affinity_exempt_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_affinity_exempt_addr: %v", err)
		}
	}

	if err = d.Set("src_affinity_exempt_addr6", flattenWebProxyGlobalSrcAffinityExemptAddr6(o["src-affinity-exempt-addr6"], d, "src_affinity_exempt_addr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-affinity-exempt-addr6"], "WebProxyGlobal-SrcAffinityExemptAddr6"); ok {
			if err = d.Set("src_affinity_exempt_addr6", vv); err != nil {
				return fmt.Errorf("Error reading src_affinity_exempt_addr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_affinity_exempt_addr6: %v", err)
		}
	}

	if err = d.Set("ssl_ca_cert", flattenWebProxyGlobalSslCaCert(o["ssl-ca-cert"], d, "ssl_ca_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-ca-cert"], "WebProxyGlobal-SslCaCert"); ok {
			if err = d.Set("ssl_ca_cert", vv); err != nil {
				return fmt.Errorf("Error reading ssl_ca_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_ca_cert: %v", err)
		}
	}

	if err = d.Set("ssl_cert", flattenWebProxyGlobalSslCert(o["ssl-cert"], d, "ssl_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-cert"], "WebProxyGlobal-SslCert"); ok {
			if err = d.Set("ssl_cert", vv); err != nil {
				return fmt.Errorf("Error reading ssl_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_cert: %v", err)
		}
	}

	if err = d.Set("strict_web_check", flattenWebProxyGlobalStrictWebCheck(o["strict-web-check"], d, "strict_web_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["strict-web-check"], "WebProxyGlobal-StrictWebCheck"); ok {
			if err = d.Set("strict_web_check", vv); err != nil {
				return fmt.Errorf("Error reading strict_web_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading strict_web_check: %v", err)
		}
	}

	if err = d.Set("tunnel_non_http", flattenWebProxyGlobalTunnelNonHttp(o["tunnel-non-http"], d, "tunnel_non_http")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-non-http"], "WebProxyGlobal-TunnelNonHttp"); ok {
			if err = d.Set("tunnel_non_http", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_non_http: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_non_http: %v", err)
		}
	}

	if err = d.Set("unknown_http_version", flattenWebProxyGlobalUnknownHttpVersion(o["unknown-http-version"], d, "unknown_http_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["unknown-http-version"], "WebProxyGlobal-UnknownHttpVersion"); ok {
			if err = d.Set("unknown_http_version", vv); err != nil {
				return fmt.Errorf("Error reading unknown_http_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unknown_http_version: %v", err)
		}
	}

	if err = d.Set("webproxy_profile", flattenWebProxyGlobalWebproxyProfile(o["webproxy-profile"], d, "webproxy_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["webproxy-profile"], "WebProxyGlobal-WebproxyProfile"); ok {
			if err = d.Set("webproxy_profile", vv); err != nil {
				return fmt.Errorf("Error reading webproxy_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webproxy_profile: %v", err)
		}
	}

	if err = d.Set("address_ip_rating", flattenWebProxyGlobalAddressIpRating(o["address-ip-rating"], d, "address_ip_rating")); err != nil {
		if vv, ok := fortiAPIPatch(o["address-ip-rating"], "WebProxyGlobal-AddressIpRating"); ok {
			if err = d.Set("address_ip_rating", vv); err != nil {
				return fmt.Errorf("Error reading address_ip_rating: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address_ip_rating: %v", err)
		}
	}

	if err = d.Set("explicit_outgoing_ip", flattenWebProxyGlobalExplicitOutgoingIp(o["explicit-outgoing-ip"], d, "explicit_outgoing_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["explicit-outgoing-ip"], "WebProxyGlobal-ExplicitOutgoingIp"); ok {
			if err = d.Set("explicit_outgoing_ip", vv); err != nil {
				return fmt.Errorf("Error reading explicit_outgoing_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading explicit_outgoing_ip: %v", err)
		}
	}

	if err = d.Set("explicit_outgoing_ip6", flattenWebProxyGlobalExplicitOutgoingIp6(o["explicit-outgoing-ip6"], d, "explicit_outgoing_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["explicit-outgoing-ip6"], "WebProxyGlobal-ExplicitOutgoingIp6"); ok {
			if err = d.Set("explicit_outgoing_ip6", vv); err != nil {
				return fmt.Errorf("Error reading explicit_outgoing_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading explicit_outgoing_ip6: %v", err)
		}
	}

	if err = d.Set("extended_log", flattenWebProxyGlobalExtendedLog(o["extended-log"], d, "extended_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["extended-log"], "WebProxyGlobal-ExtendedLog"); ok {
			if err = d.Set("extended_log", vv); err != nil {
				return fmt.Errorf("Error reading extended_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extended_log: %v", err)
		}
	}

	if err = d.Set("https_replacement_message", flattenWebProxyGlobalHttpsReplacementMessage(o["https-replacement-message"], d, "https_replacement_message")); err != nil {
		if vv, ok := fortiAPIPatch(o["https-replacement-message"], "WebProxyGlobal-HttpsReplacementMessage"); ok {
			if err = d.Set("https_replacement_message", vv); err != nil {
				return fmt.Errorf("Error reading https_replacement_message: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading https_replacement_message: %v", err)
		}
	}

	if err = d.Set("log_http_transaction", flattenWebProxyGlobalLogHttpTransaction(o["log-http-transaction"], d, "log_http_transaction")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-http-transaction"], "WebProxyGlobal-LogHttpTransaction"); ok {
			if err = d.Set("log_http_transaction", vv); err != nil {
				return fmt.Errorf("Error reading log_http_transaction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_http_transaction: %v", err)
		}
	}

	if err = d.Set("message_upon_server_error", flattenWebProxyGlobalMessageUponServerError(o["message-upon-server-error"], d, "message_upon_server_error")); err != nil {
		if vv, ok := fortiAPIPatch(o["message-upon-server-error"], "WebProxyGlobal-MessageUponServerError"); ok {
			if err = d.Set("message_upon_server_error", vv); err != nil {
				return fmt.Errorf("Error reading message_upon_server_error: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading message_upon_server_error: %v", err)
		}
	}

	if err = d.Set("realm", flattenWebProxyGlobalRealm(o["realm"], d, "realm")); err != nil {
		if vv, ok := fortiAPIPatch(o["realm"], "WebProxyGlobal-Realm"); ok {
			if err = d.Set("realm", vv); err != nil {
				return fmt.Errorf("Error reading realm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading realm: %v", err)
		}
	}

	if err = d.Set("strict_guest", flattenWebProxyGlobalStrictGuest(o["strict-guest"], d, "strict_guest")); err != nil {
		if vv, ok := fortiAPIPatch(o["strict-guest"], "WebProxyGlobal-StrictGuest"); ok {
			if err = d.Set("strict_guest", vv); err != nil {
				return fmt.Errorf("Error reading strict_guest: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading strict_guest: %v", err)
		}
	}

	if err = d.Set("trace_auth_no_rsp", flattenWebProxyGlobalTraceAuthNoRsp(o["trace-auth-no-rsp"], d, "trace_auth_no_rsp")); err != nil {
		if vv, ok := fortiAPIPatch(o["trace-auth-no-rsp"], "WebProxyGlobal-TraceAuthNoRsp"); ok {
			if err = d.Set("trace_auth_no_rsp", vv); err != nil {
				return fmt.Errorf("Error reading trace_auth_no_rsp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trace_auth_no_rsp: %v", err)
		}
	}

	if err = d.Set("use_dynamic_pkey", flattenWebProxyGlobalUseDynamicPkey(o["use-dynamic-pkey"], d, "use_dynamic_pkey")); err != nil {
		if vv, ok := fortiAPIPatch(o["use-dynamic-pkey"], "WebProxyGlobal-UseDynamicPkey"); ok {
			if err = d.Set("use_dynamic_pkey", vv); err != nil {
				return fmt.Errorf("Error reading use_dynamic_pkey: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading use_dynamic_pkey: %v", err)
		}
	}

	return nil
}

func flattenWebProxyGlobalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebProxyGlobalAlwaysLearnClientIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalFastPolicyMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalForwardProxyAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalForwardServerAffinityTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalHttp2ClientWindowSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalHttp2ServerWindowSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalLdapUserCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalLearnClientIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalLearnClientIpFromHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyGlobalLearnClientIpSrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyGlobalLearnClientIpSrcaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyGlobalLogAppId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalLogForwardServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalLogPolicyPending(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalMaxMessageLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalMaxRequestLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalMaxWafBodyCacheLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalPolicyCategoryDeepInspect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalProxyFqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalProxyTransparentCertInspection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalRequestObsFold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalSrcAffinityExemptAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyGlobalSrcAffinityExemptAddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyGlobalSslCaCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyGlobalSslCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyGlobalStrictWebCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalTunnelNonHttp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalUnknownHttpVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalWebproxyProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyGlobalAddressIpRating(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalExplicitOutgoingIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyGlobalExplicitOutgoingIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyGlobalExtendedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalHttpsReplacementMessage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalLogHttpTransaction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalMessageUponServerError(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalRealm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalStrictGuest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalTraceAuthNoRsp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalUseDynamicPkey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebProxyGlobal(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("always_learn_client_ip"); ok || d.HasChange("always_learn_client_ip") {
		t, err := expandWebProxyGlobalAlwaysLearnClientIp(d, v, "always_learn_client_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["always-learn-client-ip"] = t
		}
	}

	if v, ok := d.GetOk("fast_policy_match"); ok || d.HasChange("fast_policy_match") {
		t, err := expandWebProxyGlobalFastPolicyMatch(d, v, "fast_policy_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fast-policy-match"] = t
		}
	}

	if v, ok := d.GetOk("forward_proxy_auth"); ok || d.HasChange("forward_proxy_auth") {
		t, err := expandWebProxyGlobalForwardProxyAuth(d, v, "forward_proxy_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward-proxy-auth"] = t
		}
	}

	if v, ok := d.GetOk("forward_server_affinity_timeout"); ok || d.HasChange("forward_server_affinity_timeout") {
		t, err := expandWebProxyGlobalForwardServerAffinityTimeout(d, v, "forward_server_affinity_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward-server-affinity-timeout"] = t
		}
	}

	if v, ok := d.GetOk("http2_client_window_size"); ok || d.HasChange("http2_client_window_size") {
		t, err := expandWebProxyGlobalHttp2ClientWindowSize(d, v, "http2_client_window_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http2-client-window-size"] = t
		}
	}

	if v, ok := d.GetOk("http2_server_window_size"); ok || d.HasChange("http2_server_window_size") {
		t, err := expandWebProxyGlobalHttp2ServerWindowSize(d, v, "http2_server_window_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http2-server-window-size"] = t
		}
	}

	if v, ok := d.GetOk("ldap_user_cache"); ok || d.HasChange("ldap_user_cache") {
		t, err := expandWebProxyGlobalLdapUserCache(d, v, "ldap_user_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-user-cache"] = t
		}
	}

	if v, ok := d.GetOk("learn_client_ip"); ok || d.HasChange("learn_client_ip") {
		t, err := expandWebProxyGlobalLearnClientIp(d, v, "learn_client_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["learn-client-ip"] = t
		}
	}

	if v, ok := d.GetOk("learn_client_ip_from_header"); ok || d.HasChange("learn_client_ip_from_header") {
		t, err := expandWebProxyGlobalLearnClientIpFromHeader(d, v, "learn_client_ip_from_header")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["learn-client-ip-from-header"] = t
		}
	}

	if v, ok := d.GetOk("learn_client_ip_srcaddr"); ok || d.HasChange("learn_client_ip_srcaddr") {
		t, err := expandWebProxyGlobalLearnClientIpSrcaddr(d, v, "learn_client_ip_srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["learn-client-ip-srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("learn_client_ip_srcaddr6"); ok || d.HasChange("learn_client_ip_srcaddr6") {
		t, err := expandWebProxyGlobalLearnClientIpSrcaddr6(d, v, "learn_client_ip_srcaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["learn-client-ip-srcaddr6"] = t
		}
	}

	if v, ok := d.GetOk("log_app_id"); ok || d.HasChange("log_app_id") {
		t, err := expandWebProxyGlobalLogAppId(d, v, "log_app_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-app-id"] = t
		}
	}

	if v, ok := d.GetOk("log_forward_server"); ok || d.HasChange("log_forward_server") {
		t, err := expandWebProxyGlobalLogForwardServer(d, v, "log_forward_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-forward-server"] = t
		}
	}

	if v, ok := d.GetOk("log_policy_pending"); ok || d.HasChange("log_policy_pending") {
		t, err := expandWebProxyGlobalLogPolicyPending(d, v, "log_policy_pending")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-policy-pending"] = t
		}
	}

	if v, ok := d.GetOk("max_message_length"); ok || d.HasChange("max_message_length") {
		t, err := expandWebProxyGlobalMaxMessageLength(d, v, "max_message_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-message-length"] = t
		}
	}

	if v, ok := d.GetOk("max_request_length"); ok || d.HasChange("max_request_length") {
		t, err := expandWebProxyGlobalMaxRequestLength(d, v, "max_request_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-request-length"] = t
		}
	}

	if v, ok := d.GetOk("max_waf_body_cache_length"); ok || d.HasChange("max_waf_body_cache_length") {
		t, err := expandWebProxyGlobalMaxWafBodyCacheLength(d, v, "max_waf_body_cache_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-waf-body-cache-length"] = t
		}
	}

	if v, ok := d.GetOk("policy_category_deep_inspect"); ok || d.HasChange("policy_category_deep_inspect") {
		t, err := expandWebProxyGlobalPolicyCategoryDeepInspect(d, v, "policy_category_deep_inspect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-category-deep-inspect"] = t
		}
	}

	if v, ok := d.GetOk("proxy_fqdn"); ok || d.HasChange("proxy_fqdn") {
		t, err := expandWebProxyGlobalProxyFqdn(d, v, "proxy_fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-fqdn"] = t
		}
	}

	if v, ok := d.GetOk("proxy_transparent_cert_inspection"); ok || d.HasChange("proxy_transparent_cert_inspection") {
		t, err := expandWebProxyGlobalProxyTransparentCertInspection(d, v, "proxy_transparent_cert_inspection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-transparent-cert-inspection"] = t
		}
	}

	if v, ok := d.GetOk("request_obs_fold"); ok || d.HasChange("request_obs_fold") {
		t, err := expandWebProxyGlobalRequestObsFold(d, v, "request_obs_fold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-obs-fold"] = t
		}
	}

	if v, ok := d.GetOk("src_affinity_exempt_addr"); ok || d.HasChange("src_affinity_exempt_addr") {
		t, err := expandWebProxyGlobalSrcAffinityExemptAddr(d, v, "src_affinity_exempt_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-affinity-exempt-addr"] = t
		}
	}

	if v, ok := d.GetOk("src_affinity_exempt_addr6"); ok || d.HasChange("src_affinity_exempt_addr6") {
		t, err := expandWebProxyGlobalSrcAffinityExemptAddr6(d, v, "src_affinity_exempt_addr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-affinity-exempt-addr6"] = t
		}
	}

	if v, ok := d.GetOk("ssl_ca_cert"); ok || d.HasChange("ssl_ca_cert") {
		t, err := expandWebProxyGlobalSslCaCert(d, v, "ssl_ca_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-ca-cert"] = t
		}
	}

	if v, ok := d.GetOk("ssl_cert"); ok || d.HasChange("ssl_cert") {
		t, err := expandWebProxyGlobalSslCert(d, v, "ssl_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-cert"] = t
		}
	}

	if v, ok := d.GetOk("strict_web_check"); ok || d.HasChange("strict_web_check") {
		t, err := expandWebProxyGlobalStrictWebCheck(d, v, "strict_web_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strict-web-check"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_non_http"); ok || d.HasChange("tunnel_non_http") {
		t, err := expandWebProxyGlobalTunnelNonHttp(d, v, "tunnel_non_http")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-non-http"] = t
		}
	}

	if v, ok := d.GetOk("unknown_http_version"); ok || d.HasChange("unknown_http_version") {
		t, err := expandWebProxyGlobalUnknownHttpVersion(d, v, "unknown_http_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-http-version"] = t
		}
	}

	if v, ok := d.GetOk("webproxy_profile"); ok || d.HasChange("webproxy_profile") {
		t, err := expandWebProxyGlobalWebproxyProfile(d, v, "webproxy_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webproxy-profile"] = t
		}
	}

	if v, ok := d.GetOk("address_ip_rating"); ok || d.HasChange("address_ip_rating") {
		t, err := expandWebProxyGlobalAddressIpRating(d, v, "address_ip_rating")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address-ip-rating"] = t
		}
	}

	if v, ok := d.GetOk("explicit_outgoing_ip"); ok || d.HasChange("explicit_outgoing_ip") {
		t, err := expandWebProxyGlobalExplicitOutgoingIp(d, v, "explicit_outgoing_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["explicit-outgoing-ip"] = t
		}
	}

	if v, ok := d.GetOk("explicit_outgoing_ip6"); ok || d.HasChange("explicit_outgoing_ip6") {
		t, err := expandWebProxyGlobalExplicitOutgoingIp6(d, v, "explicit_outgoing_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["explicit-outgoing-ip6"] = t
		}
	}

	if v, ok := d.GetOk("extended_log"); ok || d.HasChange("extended_log") {
		t, err := expandWebProxyGlobalExtendedLog(d, v, "extended_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-log"] = t
		}
	}

	if v, ok := d.GetOk("https_replacement_message"); ok || d.HasChange("https_replacement_message") {
		t, err := expandWebProxyGlobalHttpsReplacementMessage(d, v, "https_replacement_message")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-replacement-message"] = t
		}
	}

	if v, ok := d.GetOk("log_http_transaction"); ok || d.HasChange("log_http_transaction") {
		t, err := expandWebProxyGlobalLogHttpTransaction(d, v, "log_http_transaction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-http-transaction"] = t
		}
	}

	if v, ok := d.GetOk("message_upon_server_error"); ok || d.HasChange("message_upon_server_error") {
		t, err := expandWebProxyGlobalMessageUponServerError(d, v, "message_upon_server_error")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["message-upon-server-error"] = t
		}
	}

	if v, ok := d.GetOk("realm"); ok || d.HasChange("realm") {
		t, err := expandWebProxyGlobalRealm(d, v, "realm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["realm"] = t
		}
	}

	if v, ok := d.GetOk("strict_guest"); ok || d.HasChange("strict_guest") {
		t, err := expandWebProxyGlobalStrictGuest(d, v, "strict_guest")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strict-guest"] = t
		}
	}

	if v, ok := d.GetOk("trace_auth_no_rsp"); ok || d.HasChange("trace_auth_no_rsp") {
		t, err := expandWebProxyGlobalTraceAuthNoRsp(d, v, "trace_auth_no_rsp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trace-auth-no-rsp"] = t
		}
	}

	if v, ok := d.GetOk("use_dynamic_pkey"); ok || d.HasChange("use_dynamic_pkey") {
		t, err := expandWebProxyGlobalUseDynamicPkey(d, v, "use_dynamic_pkey")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-dynamic-pkey"] = t
		}
	}

	return &obj, nil
}
