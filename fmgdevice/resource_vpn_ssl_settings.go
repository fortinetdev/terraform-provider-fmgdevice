// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure SSL-VPN.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnSslSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnSslSettingsUpdate,
		Read:   resourceVpnSslSettingsRead,
		Update: resourceVpnSslSettingsUpdate,
		Delete: resourceVpnSslSettingsDelete,

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
			"algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_session_check_source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"authentication_rule": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cipher": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"client_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"groups": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"portal": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"realm": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"source_address": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"source_address_negate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"source_address6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"source_address6_negate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"source_interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"user_peer": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"users": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"auto_tunnel_static_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"banned_cipher": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"browser_language_detection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"check_referer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ciphersuite": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"client_sigalgs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_portal": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"deflate_compression_level": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"deflate_min_data_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dns_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_suffix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dtls_heartbeat_fail_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dtls_heartbeat_idle_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dtls_heartbeat_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dtls_hello_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dtls_max_proto_ver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dtls_min_proto_ver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dtls_tunnel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dual_stack_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encode_2f_sequence": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encrypt_and_store_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"force_two_factor_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header_x_forwarded_for": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hsts_include_subdomains": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_compression": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_only_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_request_body_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http_request_header_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"https_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"idle_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ipv6_dns_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_dns_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_wins_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_wins_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"login_attempt_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"login_block_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"login_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"port_precedence": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reqclientcert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"saml_redirect_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server_hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"servercert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"source_address": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"source_address_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_address6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"source_address6_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssl_client_renegotiation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_insert_empty_fragment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_max_proto_ver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_min_proto_ver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"transform_backward_slashes": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tunnel_addr_assigned_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tunnel_connect_without_reauth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tunnel_ip_pools": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"tunnel_ipv6_pools": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"tunnel_user_session_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"unsafe_legacy_renegotiation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url_obscuration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_peer": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"web_mode_snat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wins_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wins_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"x_content_type_options": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ztna_trusted_client": &schema.Schema{
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

func resourceVpnSslSettingsUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectVpnSslSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateVpnSslSettings(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("VpnSslSettings")

	return resourceVpnSslSettingsRead(d, m)
}

func resourceVpnSslSettingsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteVpnSslSettings(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VpnSslSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnSslSettingsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnSslSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnSslSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslSettings resource from API: %v", err)
	}
	return nil
}

func flattenVpnSslSettingsAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthSessionCheckSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthenticationRule(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth"
		if _, ok := i["auth"]; ok {
			v := flattenVpnSslSettingsAuthenticationRuleAuth(i["auth"], d, pre_append)
			tmp["auth"] = fortiAPISubPartPatch(v, "VpnSslSettings-AuthenticationRule-Auth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := i["cipher"]; ok {
			v := flattenVpnSslSettingsAuthenticationRuleCipher(i["cipher"], d, pre_append)
			tmp["cipher"] = fortiAPISubPartPatch(v, "VpnSslSettings-AuthenticationRule-Cipher")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_cert"
		if _, ok := i["client-cert"]; ok {
			v := flattenVpnSslSettingsAuthenticationRuleClientCert(i["client-cert"], d, pre_append)
			tmp["client_cert"] = fortiAPISubPartPatch(v, "VpnSslSettings-AuthenticationRule-ClientCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "groups"
		if _, ok := i["groups"]; ok {
			v := flattenVpnSslSettingsAuthenticationRuleGroups(i["groups"], d, pre_append)
			tmp["groups"] = fortiAPISubPartPatch(v, "VpnSslSettings-AuthenticationRule-Groups")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenVpnSslSettingsAuthenticationRuleId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "VpnSslSettings-AuthenticationRule-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "portal"
		if _, ok := i["portal"]; ok {
			v := flattenVpnSslSettingsAuthenticationRulePortal(i["portal"], d, pre_append)
			tmp["portal"] = fortiAPISubPartPatch(v, "VpnSslSettings-AuthenticationRule-Portal")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realm"
		if _, ok := i["realm"]; ok {
			v := flattenVpnSslSettingsAuthenticationRuleRealm(i["realm"], d, pre_append)
			tmp["realm"] = fortiAPISubPartPatch(v, "VpnSslSettings-AuthenticationRule-Realm")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_address"
		if _, ok := i["source-address"]; ok {
			v := flattenVpnSslSettingsAuthenticationRuleSourceAddress(i["source-address"], d, pre_append)
			tmp["source_address"] = fortiAPISubPartPatch(v, "VpnSslSettings-AuthenticationRule-SourceAddress")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_address_negate"
		if _, ok := i["source-address-negate"]; ok {
			v := flattenVpnSslSettingsAuthenticationRuleSourceAddressNegate(i["source-address-negate"], d, pre_append)
			tmp["source_address_negate"] = fortiAPISubPartPatch(v, "VpnSslSettings-AuthenticationRule-SourceAddressNegate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_address6"
		if _, ok := i["source-address6"]; ok {
			v := flattenVpnSslSettingsAuthenticationRuleSourceAddress6(i["source-address6"], d, pre_append)
			tmp["source_address6"] = fortiAPISubPartPatch(v, "VpnSslSettings-AuthenticationRule-SourceAddress6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_address6_negate"
		if _, ok := i["source-address6-negate"]; ok {
			v := flattenVpnSslSettingsAuthenticationRuleSourceAddress6Negate(i["source-address6-negate"], d, pre_append)
			tmp["source_address6_negate"] = fortiAPISubPartPatch(v, "VpnSslSettings-AuthenticationRule-SourceAddress6Negate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_interface"
		if _, ok := i["source-interface"]; ok {
			v := flattenVpnSslSettingsAuthenticationRuleSourceInterface(i["source-interface"], d, pre_append)
			tmp["source_interface"] = fortiAPISubPartPatch(v, "VpnSslSettings-AuthenticationRule-SourceInterface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_peer"
		if _, ok := i["user-peer"]; ok {
			v := flattenVpnSslSettingsAuthenticationRuleUserPeer(i["user-peer"], d, pre_append)
			tmp["user_peer"] = fortiAPISubPartPatch(v, "VpnSslSettings-AuthenticationRule-UserPeer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "users"
		if _, ok := i["users"]; ok {
			v := flattenVpnSslSettingsAuthenticationRuleUsers(i["users"], d, pre_append)
			tmp["users"] = fortiAPISubPartPatch(v, "VpnSslSettings-AuthenticationRule-Users")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenVpnSslSettingsAuthenticationRuleAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthenticationRuleCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthenticationRuleClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthenticationRuleGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslSettingsAuthenticationRuleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthenticationRulePortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslSettingsAuthenticationRuleRealm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslSettingsAuthenticationRuleSourceAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslSettingsAuthenticationRuleSourceAddressNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthenticationRuleSourceAddress6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslSettingsAuthenticationRuleSourceAddress6Negate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthenticationRuleSourceInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslSettingsAuthenticationRuleUserPeer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslSettingsAuthenticationRuleUsers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslSettingsAutoTunnelStaticRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsBannedCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslSettingsBrowserLanguageDetection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsCheckReferer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsCiphersuite(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslSettingsClientSigalgs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsDefaultPortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslSettingsDeflateCompressionLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsDeflateMinDataSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsDnsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsDnsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsDnsSuffix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsDtlsHeartbeatFailCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsDtlsHeartbeatIdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsDtlsHeartbeatInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsDtlsHelloTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsDtlsMaxProtoVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsDtlsMinProtoVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsDtlsTunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsDualStackMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsEncode2FSequence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsEncryptAndStorePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsForceTwoFactorAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsHeaderXForwardedFor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsHstsIncludeSubdomains(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsHttpCompression(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsHttpOnlyCookie(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsHttpRequestBodyTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsHttpRequestHeaderTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsHttpsRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsIdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsIpv6DnsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsIpv6DnsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsIpv6WinsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsIpv6WinsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsLoginAttemptLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsLoginBlockTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsLoginTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsPortPrecedence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsReqclientcert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsSamlRedirectPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsServerHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsServercert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslSettingsSourceAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslSettingsSourceAddressNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsSourceAddress6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslSettingsSourceAddress6Negate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsSourceInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslSettingsSslClientRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsSslInsertEmptyFragment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsSslMaxProtoVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsSslMinProtoVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsTransformBackwardSlashes(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsTunnelAddrAssignedMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsTunnelConnectWithoutReauth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsTunnelIpPools(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslSettingsTunnelIpv6Pools(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslSettingsTunnelUserSessionTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsUnsafeLegacyRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsUrlObscuration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsUserPeer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslSettingsWebModeSnat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsWinsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsWinsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsXContentTypeOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsZtnaTrustedClient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnSslSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("algorithm", flattenVpnSslSettingsAlgorithm(o["algorithm"], d, "algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["algorithm"], "VpnSslSettings-Algorithm"); ok {
			if err = d.Set("algorithm", vv); err != nil {
				return fmt.Errorf("Error reading algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading algorithm: %v", err)
		}
	}

	if err = d.Set("auth_session_check_source_ip", flattenVpnSslSettingsAuthSessionCheckSourceIp(o["auth-session-check-source-ip"], d, "auth_session_check_source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-session-check-source-ip"], "VpnSslSettings-AuthSessionCheckSourceIp"); ok {
			if err = d.Set("auth_session_check_source_ip", vv); err != nil {
				return fmt.Errorf("Error reading auth_session_check_source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_session_check_source_ip: %v", err)
		}
	}

	if err = d.Set("auth_timeout", flattenVpnSslSettingsAuthTimeout(o["auth-timeout"], d, "auth_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-timeout"], "VpnSslSettings-AuthTimeout"); ok {
			if err = d.Set("auth_timeout", vv); err != nil {
				return fmt.Errorf("Error reading auth_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_timeout: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("authentication_rule", flattenVpnSslSettingsAuthenticationRule(o["authentication-rule"], d, "authentication_rule")); err != nil {
			if vv, ok := fortiAPIPatch(o["authentication-rule"], "VpnSslSettings-AuthenticationRule"); ok {
				if err = d.Set("authentication_rule", vv); err != nil {
					return fmt.Errorf("Error reading authentication_rule: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading authentication_rule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("authentication_rule"); ok {
			if err = d.Set("authentication_rule", flattenVpnSslSettingsAuthenticationRule(o["authentication-rule"], d, "authentication_rule")); err != nil {
				if vv, ok := fortiAPIPatch(o["authentication-rule"], "VpnSslSettings-AuthenticationRule"); ok {
					if err = d.Set("authentication_rule", vv); err != nil {
						return fmt.Errorf("Error reading authentication_rule: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading authentication_rule: %v", err)
				}
			}
		}
	}

	if err = d.Set("auto_tunnel_static_route", flattenVpnSslSettingsAutoTunnelStaticRoute(o["auto-tunnel-static-route"], d, "auto_tunnel_static_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-tunnel-static-route"], "VpnSslSettings-AutoTunnelStaticRoute"); ok {
			if err = d.Set("auto_tunnel_static_route", vv); err != nil {
				return fmt.Errorf("Error reading auto_tunnel_static_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_tunnel_static_route: %v", err)
		}
	}

	if err = d.Set("banned_cipher", flattenVpnSslSettingsBannedCipher(o["banned-cipher"], d, "banned_cipher")); err != nil {
		if vv, ok := fortiAPIPatch(o["banned-cipher"], "VpnSslSettings-BannedCipher"); ok {
			if err = d.Set("banned_cipher", vv); err != nil {
				return fmt.Errorf("Error reading banned_cipher: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading banned_cipher: %v", err)
		}
	}

	if err = d.Set("browser_language_detection", flattenVpnSslSettingsBrowserLanguageDetection(o["browser-language-detection"], d, "browser_language_detection")); err != nil {
		if vv, ok := fortiAPIPatch(o["browser-language-detection"], "VpnSslSettings-BrowserLanguageDetection"); ok {
			if err = d.Set("browser_language_detection", vv); err != nil {
				return fmt.Errorf("Error reading browser_language_detection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading browser_language_detection: %v", err)
		}
	}

	if err = d.Set("check_referer", flattenVpnSslSettingsCheckReferer(o["check-referer"], d, "check_referer")); err != nil {
		if vv, ok := fortiAPIPatch(o["check-referer"], "VpnSslSettings-CheckReferer"); ok {
			if err = d.Set("check_referer", vv); err != nil {
				return fmt.Errorf("Error reading check_referer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading check_referer: %v", err)
		}
	}

	if err = d.Set("ciphersuite", flattenVpnSslSettingsCiphersuite(o["ciphersuite"], d, "ciphersuite")); err != nil {
		if vv, ok := fortiAPIPatch(o["ciphersuite"], "VpnSslSettings-Ciphersuite"); ok {
			if err = d.Set("ciphersuite", vv); err != nil {
				return fmt.Errorf("Error reading ciphersuite: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ciphersuite: %v", err)
		}
	}

	if err = d.Set("client_sigalgs", flattenVpnSslSettingsClientSigalgs(o["client-sigalgs"], d, "client_sigalgs")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-sigalgs"], "VpnSslSettings-ClientSigalgs"); ok {
			if err = d.Set("client_sigalgs", vv); err != nil {
				return fmt.Errorf("Error reading client_sigalgs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_sigalgs: %v", err)
		}
	}

	if err = d.Set("default_portal", flattenVpnSslSettingsDefaultPortal(o["default-portal"], d, "default_portal")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-portal"], "VpnSslSettings-DefaultPortal"); ok {
			if err = d.Set("default_portal", vv); err != nil {
				return fmt.Errorf("Error reading default_portal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_portal: %v", err)
		}
	}

	if err = d.Set("deflate_compression_level", flattenVpnSslSettingsDeflateCompressionLevel(o["deflate-compression-level"], d, "deflate_compression_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["deflate-compression-level"], "VpnSslSettings-DeflateCompressionLevel"); ok {
			if err = d.Set("deflate_compression_level", vv); err != nil {
				return fmt.Errorf("Error reading deflate_compression_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading deflate_compression_level: %v", err)
		}
	}

	if err = d.Set("deflate_min_data_size", flattenVpnSslSettingsDeflateMinDataSize(o["deflate-min-data-size"], d, "deflate_min_data_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["deflate-min-data-size"], "VpnSslSettings-DeflateMinDataSize"); ok {
			if err = d.Set("deflate_min_data_size", vv); err != nil {
				return fmt.Errorf("Error reading deflate_min_data_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading deflate_min_data_size: %v", err)
		}
	}

	if err = d.Set("dns_server1", flattenVpnSslSettingsDnsServer1(o["dns-server1"], d, "dns_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server1"], "VpnSslSettings-DnsServer1"); ok {
			if err = d.Set("dns_server1", vv); err != nil {
				return fmt.Errorf("Error reading dns_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server1: %v", err)
		}
	}

	if err = d.Set("dns_server2", flattenVpnSslSettingsDnsServer2(o["dns-server2"], d, "dns_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server2"], "VpnSslSettings-DnsServer2"); ok {
			if err = d.Set("dns_server2", vv); err != nil {
				return fmt.Errorf("Error reading dns_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server2: %v", err)
		}
	}

	if err = d.Set("dns_suffix", flattenVpnSslSettingsDnsSuffix(o["dns-suffix"], d, "dns_suffix")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-suffix"], "VpnSslSettings-DnsSuffix"); ok {
			if err = d.Set("dns_suffix", vv); err != nil {
				return fmt.Errorf("Error reading dns_suffix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_suffix: %v", err)
		}
	}

	if err = d.Set("dtls_heartbeat_fail_count", flattenVpnSslSettingsDtlsHeartbeatFailCount(o["dtls-heartbeat-fail-count"], d, "dtls_heartbeat_fail_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["dtls-heartbeat-fail-count"], "VpnSslSettings-DtlsHeartbeatFailCount"); ok {
			if err = d.Set("dtls_heartbeat_fail_count", vv); err != nil {
				return fmt.Errorf("Error reading dtls_heartbeat_fail_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dtls_heartbeat_fail_count: %v", err)
		}
	}

	if err = d.Set("dtls_heartbeat_idle_timeout", flattenVpnSslSettingsDtlsHeartbeatIdleTimeout(o["dtls-heartbeat-idle-timeout"], d, "dtls_heartbeat_idle_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["dtls-heartbeat-idle-timeout"], "VpnSslSettings-DtlsHeartbeatIdleTimeout"); ok {
			if err = d.Set("dtls_heartbeat_idle_timeout", vv); err != nil {
				return fmt.Errorf("Error reading dtls_heartbeat_idle_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dtls_heartbeat_idle_timeout: %v", err)
		}
	}

	if err = d.Set("dtls_heartbeat_interval", flattenVpnSslSettingsDtlsHeartbeatInterval(o["dtls-heartbeat-interval"], d, "dtls_heartbeat_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["dtls-heartbeat-interval"], "VpnSslSettings-DtlsHeartbeatInterval"); ok {
			if err = d.Set("dtls_heartbeat_interval", vv); err != nil {
				return fmt.Errorf("Error reading dtls_heartbeat_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dtls_heartbeat_interval: %v", err)
		}
	}

	if err = d.Set("dtls_hello_timeout", flattenVpnSslSettingsDtlsHelloTimeout(o["dtls-hello-timeout"], d, "dtls_hello_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["dtls-hello-timeout"], "VpnSslSettings-DtlsHelloTimeout"); ok {
			if err = d.Set("dtls_hello_timeout", vv); err != nil {
				return fmt.Errorf("Error reading dtls_hello_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dtls_hello_timeout: %v", err)
		}
	}

	if err = d.Set("dtls_max_proto_ver", flattenVpnSslSettingsDtlsMaxProtoVer(o["dtls-max-proto-ver"], d, "dtls_max_proto_ver")); err != nil {
		if vv, ok := fortiAPIPatch(o["dtls-max-proto-ver"], "VpnSslSettings-DtlsMaxProtoVer"); ok {
			if err = d.Set("dtls_max_proto_ver", vv); err != nil {
				return fmt.Errorf("Error reading dtls_max_proto_ver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dtls_max_proto_ver: %v", err)
		}
	}

	if err = d.Set("dtls_min_proto_ver", flattenVpnSslSettingsDtlsMinProtoVer(o["dtls-min-proto-ver"], d, "dtls_min_proto_ver")); err != nil {
		if vv, ok := fortiAPIPatch(o["dtls-min-proto-ver"], "VpnSslSettings-DtlsMinProtoVer"); ok {
			if err = d.Set("dtls_min_proto_ver", vv); err != nil {
				return fmt.Errorf("Error reading dtls_min_proto_ver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dtls_min_proto_ver: %v", err)
		}
	}

	if err = d.Set("dtls_tunnel", flattenVpnSslSettingsDtlsTunnel(o["dtls-tunnel"], d, "dtls_tunnel")); err != nil {
		if vv, ok := fortiAPIPatch(o["dtls-tunnel"], "VpnSslSettings-DtlsTunnel"); ok {
			if err = d.Set("dtls_tunnel", vv); err != nil {
				return fmt.Errorf("Error reading dtls_tunnel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dtls_tunnel: %v", err)
		}
	}

	if err = d.Set("dual_stack_mode", flattenVpnSslSettingsDualStackMode(o["dual-stack-mode"], d, "dual_stack_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["dual-stack-mode"], "VpnSslSettings-DualStackMode"); ok {
			if err = d.Set("dual_stack_mode", vv); err != nil {
				return fmt.Errorf("Error reading dual_stack_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dual_stack_mode: %v", err)
		}
	}

	if err = d.Set("encode_2f_sequence", flattenVpnSslSettingsEncode2FSequence(o["encode-2f-sequence"], d, "encode_2f_sequence")); err != nil {
		if vv, ok := fortiAPIPatch(o["encode-2f-sequence"], "VpnSslSettings-Encode2FSequence"); ok {
			if err = d.Set("encode_2f_sequence", vv); err != nil {
				return fmt.Errorf("Error reading encode_2f_sequence: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading encode_2f_sequence: %v", err)
		}
	}

	if err = d.Set("encrypt_and_store_password", flattenVpnSslSettingsEncryptAndStorePassword(o["encrypt-and-store-password"], d, "encrypt_and_store_password")); err != nil {
		if vv, ok := fortiAPIPatch(o["encrypt-and-store-password"], "VpnSslSettings-EncryptAndStorePassword"); ok {
			if err = d.Set("encrypt_and_store_password", vv); err != nil {
				return fmt.Errorf("Error reading encrypt_and_store_password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading encrypt_and_store_password: %v", err)
		}
	}

	if err = d.Set("force_two_factor_auth", flattenVpnSslSettingsForceTwoFactorAuth(o["force-two-factor-auth"], d, "force_two_factor_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["force-two-factor-auth"], "VpnSslSettings-ForceTwoFactorAuth"); ok {
			if err = d.Set("force_two_factor_auth", vv); err != nil {
				return fmt.Errorf("Error reading force_two_factor_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading force_two_factor_auth: %v", err)
		}
	}

	if err = d.Set("header_x_forwarded_for", flattenVpnSslSettingsHeaderXForwardedFor(o["header-x-forwarded-for"], d, "header_x_forwarded_for")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-x-forwarded-for"], "VpnSslSettings-HeaderXForwardedFor"); ok {
			if err = d.Set("header_x_forwarded_for", vv); err != nil {
				return fmt.Errorf("Error reading header_x_forwarded_for: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_x_forwarded_for: %v", err)
		}
	}

	if err = d.Set("hsts_include_subdomains", flattenVpnSslSettingsHstsIncludeSubdomains(o["hsts-include-subdomains"], d, "hsts_include_subdomains")); err != nil {
		if vv, ok := fortiAPIPatch(o["hsts-include-subdomains"], "VpnSslSettings-HstsIncludeSubdomains"); ok {
			if err = d.Set("hsts_include_subdomains", vv); err != nil {
				return fmt.Errorf("Error reading hsts_include_subdomains: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hsts_include_subdomains: %v", err)
		}
	}

	if err = d.Set("http_compression", flattenVpnSslSettingsHttpCompression(o["http-compression"], d, "http_compression")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-compression"], "VpnSslSettings-HttpCompression"); ok {
			if err = d.Set("http_compression", vv); err != nil {
				return fmt.Errorf("Error reading http_compression: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_compression: %v", err)
		}
	}

	if err = d.Set("http_only_cookie", flattenVpnSslSettingsHttpOnlyCookie(o["http-only-cookie"], d, "http_only_cookie")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-only-cookie"], "VpnSslSettings-HttpOnlyCookie"); ok {
			if err = d.Set("http_only_cookie", vv); err != nil {
				return fmt.Errorf("Error reading http_only_cookie: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_only_cookie: %v", err)
		}
	}

	if err = d.Set("http_request_body_timeout", flattenVpnSslSettingsHttpRequestBodyTimeout(o["http-request-body-timeout"], d, "http_request_body_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-request-body-timeout"], "VpnSslSettings-HttpRequestBodyTimeout"); ok {
			if err = d.Set("http_request_body_timeout", vv); err != nil {
				return fmt.Errorf("Error reading http_request_body_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_request_body_timeout: %v", err)
		}
	}

	if err = d.Set("http_request_header_timeout", flattenVpnSslSettingsHttpRequestHeaderTimeout(o["http-request-header-timeout"], d, "http_request_header_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-request-header-timeout"], "VpnSslSettings-HttpRequestHeaderTimeout"); ok {
			if err = d.Set("http_request_header_timeout", vv); err != nil {
				return fmt.Errorf("Error reading http_request_header_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_request_header_timeout: %v", err)
		}
	}

	if err = d.Set("https_redirect", flattenVpnSslSettingsHttpsRedirect(o["https-redirect"], d, "https_redirect")); err != nil {
		if vv, ok := fortiAPIPatch(o["https-redirect"], "VpnSslSettings-HttpsRedirect"); ok {
			if err = d.Set("https_redirect", vv); err != nil {
				return fmt.Errorf("Error reading https_redirect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading https_redirect: %v", err)
		}
	}

	if err = d.Set("idle_timeout", flattenVpnSslSettingsIdleTimeout(o["idle-timeout"], d, "idle_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["idle-timeout"], "VpnSslSettings-IdleTimeout"); ok {
			if err = d.Set("idle_timeout", vv); err != nil {
				return fmt.Errorf("Error reading idle_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idle_timeout: %v", err)
		}
	}

	if err = d.Set("ipv6_dns_server1", flattenVpnSslSettingsIpv6DnsServer1(o["ipv6-dns-server1"], d, "ipv6_dns_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-dns-server1"], "VpnSslSettings-Ipv6DnsServer1"); ok {
			if err = d.Set("ipv6_dns_server1", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_dns_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_dns_server1: %v", err)
		}
	}

	if err = d.Set("ipv6_dns_server2", flattenVpnSslSettingsIpv6DnsServer2(o["ipv6-dns-server2"], d, "ipv6_dns_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-dns-server2"], "VpnSslSettings-Ipv6DnsServer2"); ok {
			if err = d.Set("ipv6_dns_server2", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_dns_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_dns_server2: %v", err)
		}
	}

	if err = d.Set("ipv6_wins_server1", flattenVpnSslSettingsIpv6WinsServer1(o["ipv6-wins-server1"], d, "ipv6_wins_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-wins-server1"], "VpnSslSettings-Ipv6WinsServer1"); ok {
			if err = d.Set("ipv6_wins_server1", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_wins_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_wins_server1: %v", err)
		}
	}

	if err = d.Set("ipv6_wins_server2", flattenVpnSslSettingsIpv6WinsServer2(o["ipv6-wins-server2"], d, "ipv6_wins_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-wins-server2"], "VpnSslSettings-Ipv6WinsServer2"); ok {
			if err = d.Set("ipv6_wins_server2", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_wins_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_wins_server2: %v", err)
		}
	}

	if err = d.Set("login_attempt_limit", flattenVpnSslSettingsLoginAttemptLimit(o["login-attempt-limit"], d, "login_attempt_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["login-attempt-limit"], "VpnSslSettings-LoginAttemptLimit"); ok {
			if err = d.Set("login_attempt_limit", vv); err != nil {
				return fmt.Errorf("Error reading login_attempt_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading login_attempt_limit: %v", err)
		}
	}

	if err = d.Set("login_block_time", flattenVpnSslSettingsLoginBlockTime(o["login-block-time"], d, "login_block_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["login-block-time"], "VpnSslSettings-LoginBlockTime"); ok {
			if err = d.Set("login_block_time", vv); err != nil {
				return fmt.Errorf("Error reading login_block_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading login_block_time: %v", err)
		}
	}

	if err = d.Set("login_timeout", flattenVpnSslSettingsLoginTimeout(o["login-timeout"], d, "login_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["login-timeout"], "VpnSslSettings-LoginTimeout"); ok {
			if err = d.Set("login_timeout", vv); err != nil {
				return fmt.Errorf("Error reading login_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading login_timeout: %v", err)
		}
	}

	if err = d.Set("port", flattenVpnSslSettingsPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "VpnSslSettings-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("port_precedence", flattenVpnSslSettingsPortPrecedence(o["port-precedence"], d, "port_precedence")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-precedence"], "VpnSslSettings-PortPrecedence"); ok {
			if err = d.Set("port_precedence", vv); err != nil {
				return fmt.Errorf("Error reading port_precedence: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_precedence: %v", err)
		}
	}

	if err = d.Set("reqclientcert", flattenVpnSslSettingsReqclientcert(o["reqclientcert"], d, "reqclientcert")); err != nil {
		if vv, ok := fortiAPIPatch(o["reqclientcert"], "VpnSslSettings-Reqclientcert"); ok {
			if err = d.Set("reqclientcert", vv); err != nil {
				return fmt.Errorf("Error reading reqclientcert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reqclientcert: %v", err)
		}
	}

	if err = d.Set("saml_redirect_port", flattenVpnSslSettingsSamlRedirectPort(o["saml-redirect-port"], d, "saml_redirect_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["saml-redirect-port"], "VpnSslSettings-SamlRedirectPort"); ok {
			if err = d.Set("saml_redirect_port", vv); err != nil {
				return fmt.Errorf("Error reading saml_redirect_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading saml_redirect_port: %v", err)
		}
	}

	if err = d.Set("server_hostname", flattenVpnSslSettingsServerHostname(o["server-hostname"], d, "server_hostname")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-hostname"], "VpnSslSettings-ServerHostname"); ok {
			if err = d.Set("server_hostname", vv); err != nil {
				return fmt.Errorf("Error reading server_hostname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_hostname: %v", err)
		}
	}

	if err = d.Set("servercert", flattenVpnSslSettingsServercert(o["servercert"], d, "servercert")); err != nil {
		if vv, ok := fortiAPIPatch(o["servercert"], "VpnSslSettings-Servercert"); ok {
			if err = d.Set("servercert", vv); err != nil {
				return fmt.Errorf("Error reading servercert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading servercert: %v", err)
		}
	}

	if err = d.Set("source_address", flattenVpnSslSettingsSourceAddress(o["source-address"], d, "source_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-address"], "VpnSslSettings-SourceAddress"); ok {
			if err = d.Set("source_address", vv); err != nil {
				return fmt.Errorf("Error reading source_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_address: %v", err)
		}
	}

	if err = d.Set("source_address_negate", flattenVpnSslSettingsSourceAddressNegate(o["source-address-negate"], d, "source_address_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-address-negate"], "VpnSslSettings-SourceAddressNegate"); ok {
			if err = d.Set("source_address_negate", vv); err != nil {
				return fmt.Errorf("Error reading source_address_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_address_negate: %v", err)
		}
	}

	if err = d.Set("source_address6", flattenVpnSslSettingsSourceAddress6(o["source-address6"], d, "source_address6")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-address6"], "VpnSslSettings-SourceAddress6"); ok {
			if err = d.Set("source_address6", vv); err != nil {
				return fmt.Errorf("Error reading source_address6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_address6: %v", err)
		}
	}

	if err = d.Set("source_address6_negate", flattenVpnSslSettingsSourceAddress6Negate(o["source-address6-negate"], d, "source_address6_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-address6-negate"], "VpnSslSettings-SourceAddress6Negate"); ok {
			if err = d.Set("source_address6_negate", vv); err != nil {
				return fmt.Errorf("Error reading source_address6_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_address6_negate: %v", err)
		}
	}

	if err = d.Set("source_interface", flattenVpnSslSettingsSourceInterface(o["source-interface"], d, "source_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-interface"], "VpnSslSettings-SourceInterface"); ok {
			if err = d.Set("source_interface", vv); err != nil {
				return fmt.Errorf("Error reading source_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_interface: %v", err)
		}
	}

	if err = d.Set("ssl_client_renegotiation", flattenVpnSslSettingsSslClientRenegotiation(o["ssl-client-renegotiation"], d, "ssl_client_renegotiation")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-renegotiation"], "VpnSslSettings-SslClientRenegotiation"); ok {
			if err = d.Set("ssl_client_renegotiation", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_renegotiation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_renegotiation: %v", err)
		}
	}

	if err = d.Set("ssl_insert_empty_fragment", flattenVpnSslSettingsSslInsertEmptyFragment(o["ssl-insert-empty-fragment"], d, "ssl_insert_empty_fragment")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-insert-empty-fragment"], "VpnSslSettings-SslInsertEmptyFragment"); ok {
			if err = d.Set("ssl_insert_empty_fragment", vv); err != nil {
				return fmt.Errorf("Error reading ssl_insert_empty_fragment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_insert_empty_fragment: %v", err)
		}
	}

	if err = d.Set("ssl_max_proto_ver", flattenVpnSslSettingsSslMaxProtoVer(o["ssl-max-proto-ver"], d, "ssl_max_proto_ver")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-max-proto-ver"], "VpnSslSettings-SslMaxProtoVer"); ok {
			if err = d.Set("ssl_max_proto_ver", vv); err != nil {
				return fmt.Errorf("Error reading ssl_max_proto_ver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_max_proto_ver: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_ver", flattenVpnSslSettingsSslMinProtoVer(o["ssl-min-proto-ver"], d, "ssl_min_proto_ver")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-min-proto-ver"], "VpnSslSettings-SslMinProtoVer"); ok {
			if err = d.Set("ssl_min_proto_ver", vv); err != nil {
				return fmt.Errorf("Error reading ssl_min_proto_ver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_min_proto_ver: %v", err)
		}
	}

	if err = d.Set("status", flattenVpnSslSettingsStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "VpnSslSettings-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("transform_backward_slashes", flattenVpnSslSettingsTransformBackwardSlashes(o["transform-backward-slashes"], d, "transform_backward_slashes")); err != nil {
		if vv, ok := fortiAPIPatch(o["transform-backward-slashes"], "VpnSslSettings-TransformBackwardSlashes"); ok {
			if err = d.Set("transform_backward_slashes", vv); err != nil {
				return fmt.Errorf("Error reading transform_backward_slashes: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading transform_backward_slashes: %v", err)
		}
	}

	if err = d.Set("tunnel_addr_assigned_method", flattenVpnSslSettingsTunnelAddrAssignedMethod(o["tunnel-addr-assigned-method"], d, "tunnel_addr_assigned_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-addr-assigned-method"], "VpnSslSettings-TunnelAddrAssignedMethod"); ok {
			if err = d.Set("tunnel_addr_assigned_method", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_addr_assigned_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_addr_assigned_method: %v", err)
		}
	}

	if err = d.Set("tunnel_connect_without_reauth", flattenVpnSslSettingsTunnelConnectWithoutReauth(o["tunnel-connect-without-reauth"], d, "tunnel_connect_without_reauth")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-connect-without-reauth"], "VpnSslSettings-TunnelConnectWithoutReauth"); ok {
			if err = d.Set("tunnel_connect_without_reauth", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_connect_without_reauth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_connect_without_reauth: %v", err)
		}
	}

	if err = d.Set("tunnel_ip_pools", flattenVpnSslSettingsTunnelIpPools(o["tunnel-ip-pools"], d, "tunnel_ip_pools")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-ip-pools"], "VpnSslSettings-TunnelIpPools"); ok {
			if err = d.Set("tunnel_ip_pools", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_ip_pools: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_ip_pools: %v", err)
		}
	}

	if err = d.Set("tunnel_ipv6_pools", flattenVpnSslSettingsTunnelIpv6Pools(o["tunnel-ipv6-pools"], d, "tunnel_ipv6_pools")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-ipv6-pools"], "VpnSslSettings-TunnelIpv6Pools"); ok {
			if err = d.Set("tunnel_ipv6_pools", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_ipv6_pools: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_ipv6_pools: %v", err)
		}
	}

	if err = d.Set("tunnel_user_session_timeout", flattenVpnSslSettingsTunnelUserSessionTimeout(o["tunnel-user-session-timeout"], d, "tunnel_user_session_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-user-session-timeout"], "VpnSslSettings-TunnelUserSessionTimeout"); ok {
			if err = d.Set("tunnel_user_session_timeout", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_user_session_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_user_session_timeout: %v", err)
		}
	}

	if err = d.Set("unsafe_legacy_renegotiation", flattenVpnSslSettingsUnsafeLegacyRenegotiation(o["unsafe-legacy-renegotiation"], d, "unsafe_legacy_renegotiation")); err != nil {
		if vv, ok := fortiAPIPatch(o["unsafe-legacy-renegotiation"], "VpnSslSettings-UnsafeLegacyRenegotiation"); ok {
			if err = d.Set("unsafe_legacy_renegotiation", vv); err != nil {
				return fmt.Errorf("Error reading unsafe_legacy_renegotiation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unsafe_legacy_renegotiation: %v", err)
		}
	}

	if err = d.Set("url_obscuration", flattenVpnSslSettingsUrlObscuration(o["url-obscuration"], d, "url_obscuration")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-obscuration"], "VpnSslSettings-UrlObscuration"); ok {
			if err = d.Set("url_obscuration", vv); err != nil {
				return fmt.Errorf("Error reading url_obscuration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_obscuration: %v", err)
		}
	}

	if err = d.Set("user_peer", flattenVpnSslSettingsUserPeer(o["user-peer"], d, "user_peer")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-peer"], "VpnSslSettings-UserPeer"); ok {
			if err = d.Set("user_peer", vv); err != nil {
				return fmt.Errorf("Error reading user_peer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_peer: %v", err)
		}
	}

	if err = d.Set("web_mode_snat", flattenVpnSslSettingsWebModeSnat(o["web-mode-snat"], d, "web_mode_snat")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-mode-snat"], "VpnSslSettings-WebModeSnat"); ok {
			if err = d.Set("web_mode_snat", vv); err != nil {
				return fmt.Errorf("Error reading web_mode_snat: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_mode_snat: %v", err)
		}
	}

	if err = d.Set("wins_server1", flattenVpnSslSettingsWinsServer1(o["wins-server1"], d, "wins_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["wins-server1"], "VpnSslSettings-WinsServer1"); ok {
			if err = d.Set("wins_server1", vv); err != nil {
				return fmt.Errorf("Error reading wins_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wins_server1: %v", err)
		}
	}

	if err = d.Set("wins_server2", flattenVpnSslSettingsWinsServer2(o["wins-server2"], d, "wins_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["wins-server2"], "VpnSslSettings-WinsServer2"); ok {
			if err = d.Set("wins_server2", vv); err != nil {
				return fmt.Errorf("Error reading wins_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wins_server2: %v", err)
		}
	}

	if err = d.Set("x_content_type_options", flattenVpnSslSettingsXContentTypeOptions(o["x-content-type-options"], d, "x_content_type_options")); err != nil {
		if vv, ok := fortiAPIPatch(o["x-content-type-options"], "VpnSslSettings-XContentTypeOptions"); ok {
			if err = d.Set("x_content_type_options", vv); err != nil {
				return fmt.Errorf("Error reading x_content_type_options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading x_content_type_options: %v", err)
		}
	}

	if err = d.Set("ztna_trusted_client", flattenVpnSslSettingsZtnaTrustedClient(o["ztna-trusted-client"], d, "ztna_trusted_client")); err != nil {
		if vv, ok := fortiAPIPatch(o["ztna-trusted-client"], "VpnSslSettings-ZtnaTrustedClient"); ok {
			if err = d.Set("ztna_trusted_client", vv); err != nil {
				return fmt.Errorf("Error reading ztna_trusted_client: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ztna_trusted_client: %v", err)
		}
	}

	return nil
}

func flattenVpnSslSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnSslSettingsAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthSessionCheckSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthenticationRule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth"], _ = expandVpnSslSettingsAuthenticationRuleAuth(d, i["auth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cipher"], _ = expandVpnSslSettingsAuthenticationRuleCipher(d, i["cipher"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["client-cert"], _ = expandVpnSslSettingsAuthenticationRuleClientCert(d, i["client_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "groups"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["groups"], _ = expandVpnSslSettingsAuthenticationRuleGroups(d, i["groups"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandVpnSslSettingsAuthenticationRuleId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "portal"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["portal"], _ = expandVpnSslSettingsAuthenticationRulePortal(d, i["portal"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realm"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["realm"], _ = expandVpnSslSettingsAuthenticationRuleRealm(d, i["realm"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-address"], _ = expandVpnSslSettingsAuthenticationRuleSourceAddress(d, i["source_address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_address_negate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-address-negate"], _ = expandVpnSslSettingsAuthenticationRuleSourceAddressNegate(d, i["source_address_negate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_address6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-address6"], _ = expandVpnSslSettingsAuthenticationRuleSourceAddress6(d, i["source_address6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_address6_negate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-address6-negate"], _ = expandVpnSslSettingsAuthenticationRuleSourceAddress6Negate(d, i["source_address6_negate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-interface"], _ = expandVpnSslSettingsAuthenticationRuleSourceInterface(d, i["source_interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_peer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-peer"], _ = expandVpnSslSettingsAuthenticationRuleUserPeer(d, i["user_peer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "users"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["users"], _ = expandVpnSslSettingsAuthenticationRuleUsers(d, i["users"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandVpnSslSettingsAuthenticationRuleAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthenticationRuleCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthenticationRuleClientCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthenticationRuleGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslSettingsAuthenticationRuleId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthenticationRulePortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslSettingsAuthenticationRuleRealm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslSettingsAuthenticationRuleSourceAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslSettingsAuthenticationRuleSourceAddressNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthenticationRuleSourceAddress6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslSettingsAuthenticationRuleSourceAddress6Negate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthenticationRuleSourceInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslSettingsAuthenticationRuleUserPeer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslSettingsAuthenticationRuleUsers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslSettingsAutoTunnelStaticRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsBannedCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslSettingsBrowserLanguageDetection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsCheckReferer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsCiphersuite(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslSettingsClientSigalgs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsDefaultPortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslSettingsDeflateCompressionLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsDeflateMinDataSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsDnsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsDnsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsDnsSuffix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsDtlsHeartbeatFailCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsDtlsHeartbeatIdleTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsDtlsHeartbeatInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsDtlsHelloTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsDtlsMaxProtoVer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsDtlsMinProtoVer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsDtlsTunnel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsDualStackMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsEncode2FSequence(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsEncryptAndStorePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsForceTwoFactorAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsHeaderXForwardedFor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsHstsIncludeSubdomains(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsHttpCompression(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsHttpOnlyCookie(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsHttpRequestBodyTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsHttpRequestHeaderTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsHttpsRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsIdleTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsIpv6DnsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsIpv6DnsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsIpv6WinsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsIpv6WinsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsLoginAttemptLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsLoginBlockTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsLoginTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsPortPrecedence(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsReqclientcert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsSamlRedirectPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsServerHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsServercert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslSettingsSourceAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslSettingsSourceAddressNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsSourceAddress6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslSettingsSourceAddress6Negate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsSourceInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslSettingsSslClientRenegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsSslInsertEmptyFragment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsSslMaxProtoVer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsSslMinProtoVer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsTransformBackwardSlashes(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsTunnelAddrAssignedMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsTunnelConnectWithoutReauth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsTunnelIpPools(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslSettingsTunnelIpv6Pools(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslSettingsTunnelUserSessionTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsUnsafeLegacyRenegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsUrlObscuration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsUserPeer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslSettingsWebModeSnat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsWinsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsWinsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsXContentTypeOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsZtnaTrustedClient(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnSslSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("algorithm"); ok || d.HasChange("algorithm") {
		t, err := expandVpnSslSettingsAlgorithm(d, v, "algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["algorithm"] = t
		}
	}

	if v, ok := d.GetOk("auth_session_check_source_ip"); ok || d.HasChange("auth_session_check_source_ip") {
		t, err := expandVpnSslSettingsAuthSessionCheckSourceIp(d, v, "auth_session_check_source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-session-check-source-ip"] = t
		}
	}

	if v, ok := d.GetOk("auth_timeout"); ok || d.HasChange("auth_timeout") {
		t, err := expandVpnSslSettingsAuthTimeout(d, v, "auth_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-timeout"] = t
		}
	}

	if v, ok := d.GetOk("authentication_rule"); ok || d.HasChange("authentication_rule") {
		t, err := expandVpnSslSettingsAuthenticationRule(d, v, "authentication_rule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication-rule"] = t
		}
	}

	if v, ok := d.GetOk("auto_tunnel_static_route"); ok || d.HasChange("auto_tunnel_static_route") {
		t, err := expandVpnSslSettingsAutoTunnelStaticRoute(d, v, "auto_tunnel_static_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-tunnel-static-route"] = t
		}
	}

	if v, ok := d.GetOk("banned_cipher"); ok || d.HasChange("banned_cipher") {
		t, err := expandVpnSslSettingsBannedCipher(d, v, "banned_cipher")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["banned-cipher"] = t
		}
	}

	if v, ok := d.GetOk("browser_language_detection"); ok || d.HasChange("browser_language_detection") {
		t, err := expandVpnSslSettingsBrowserLanguageDetection(d, v, "browser_language_detection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["browser-language-detection"] = t
		}
	}

	if v, ok := d.GetOk("check_referer"); ok || d.HasChange("check_referer") {
		t, err := expandVpnSslSettingsCheckReferer(d, v, "check_referer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["check-referer"] = t
		}
	}

	if v, ok := d.GetOk("ciphersuite"); ok || d.HasChange("ciphersuite") {
		t, err := expandVpnSslSettingsCiphersuite(d, v, "ciphersuite")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ciphersuite"] = t
		}
	}

	if v, ok := d.GetOk("client_sigalgs"); ok || d.HasChange("client_sigalgs") {
		t, err := expandVpnSslSettingsClientSigalgs(d, v, "client_sigalgs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-sigalgs"] = t
		}
	}

	if v, ok := d.GetOk("default_portal"); ok || d.HasChange("default_portal") {
		t, err := expandVpnSslSettingsDefaultPortal(d, v, "default_portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-portal"] = t
		}
	}

	if v, ok := d.GetOk("deflate_compression_level"); ok || d.HasChange("deflate_compression_level") {
		t, err := expandVpnSslSettingsDeflateCompressionLevel(d, v, "deflate_compression_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deflate-compression-level"] = t
		}
	}

	if v, ok := d.GetOk("deflate_min_data_size"); ok || d.HasChange("deflate_min_data_size") {
		t, err := expandVpnSslSettingsDeflateMinDataSize(d, v, "deflate_min_data_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deflate-min-data-size"] = t
		}
	}

	if v, ok := d.GetOk("dns_server1"); ok || d.HasChange("dns_server1") {
		t, err := expandVpnSslSettingsDnsServer1(d, v, "dns_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("dns_server2"); ok || d.HasChange("dns_server2") {
		t, err := expandVpnSslSettingsDnsServer2(d, v, "dns_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("dns_suffix"); ok || d.HasChange("dns_suffix") {
		t, err := expandVpnSslSettingsDnsSuffix(d, v, "dns_suffix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-suffix"] = t
		}
	}

	if v, ok := d.GetOk("dtls_heartbeat_fail_count"); ok || d.HasChange("dtls_heartbeat_fail_count") {
		t, err := expandVpnSslSettingsDtlsHeartbeatFailCount(d, v, "dtls_heartbeat_fail_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dtls-heartbeat-fail-count"] = t
		}
	}

	if v, ok := d.GetOk("dtls_heartbeat_idle_timeout"); ok || d.HasChange("dtls_heartbeat_idle_timeout") {
		t, err := expandVpnSslSettingsDtlsHeartbeatIdleTimeout(d, v, "dtls_heartbeat_idle_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dtls-heartbeat-idle-timeout"] = t
		}
	}

	if v, ok := d.GetOk("dtls_heartbeat_interval"); ok || d.HasChange("dtls_heartbeat_interval") {
		t, err := expandVpnSslSettingsDtlsHeartbeatInterval(d, v, "dtls_heartbeat_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dtls-heartbeat-interval"] = t
		}
	}

	if v, ok := d.GetOk("dtls_hello_timeout"); ok || d.HasChange("dtls_hello_timeout") {
		t, err := expandVpnSslSettingsDtlsHelloTimeout(d, v, "dtls_hello_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dtls-hello-timeout"] = t
		}
	}

	if v, ok := d.GetOk("dtls_max_proto_ver"); ok || d.HasChange("dtls_max_proto_ver") {
		t, err := expandVpnSslSettingsDtlsMaxProtoVer(d, v, "dtls_max_proto_ver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dtls-max-proto-ver"] = t
		}
	}

	if v, ok := d.GetOk("dtls_min_proto_ver"); ok || d.HasChange("dtls_min_proto_ver") {
		t, err := expandVpnSslSettingsDtlsMinProtoVer(d, v, "dtls_min_proto_ver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dtls-min-proto-ver"] = t
		}
	}

	if v, ok := d.GetOk("dtls_tunnel"); ok || d.HasChange("dtls_tunnel") {
		t, err := expandVpnSslSettingsDtlsTunnel(d, v, "dtls_tunnel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dtls-tunnel"] = t
		}
	}

	if v, ok := d.GetOk("dual_stack_mode"); ok || d.HasChange("dual_stack_mode") {
		t, err := expandVpnSslSettingsDualStackMode(d, v, "dual_stack_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dual-stack-mode"] = t
		}
	}

	if v, ok := d.GetOk("encode_2f_sequence"); ok || d.HasChange("encode_2f_sequence") {
		t, err := expandVpnSslSettingsEncode2FSequence(d, v, "encode_2f_sequence")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encode-2f-sequence"] = t
		}
	}

	if v, ok := d.GetOk("encrypt_and_store_password"); ok || d.HasChange("encrypt_and_store_password") {
		t, err := expandVpnSslSettingsEncryptAndStorePassword(d, v, "encrypt_and_store_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encrypt-and-store-password"] = t
		}
	}

	if v, ok := d.GetOk("force_two_factor_auth"); ok || d.HasChange("force_two_factor_auth") {
		t, err := expandVpnSslSettingsForceTwoFactorAuth(d, v, "force_two_factor_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["force-two-factor-auth"] = t
		}
	}

	if v, ok := d.GetOk("header_x_forwarded_for"); ok || d.HasChange("header_x_forwarded_for") {
		t, err := expandVpnSslSettingsHeaderXForwardedFor(d, v, "header_x_forwarded_for")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-x-forwarded-for"] = t
		}
	}

	if v, ok := d.GetOk("hsts_include_subdomains"); ok || d.HasChange("hsts_include_subdomains") {
		t, err := expandVpnSslSettingsHstsIncludeSubdomains(d, v, "hsts_include_subdomains")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hsts-include-subdomains"] = t
		}
	}

	if v, ok := d.GetOk("http_compression"); ok || d.HasChange("http_compression") {
		t, err := expandVpnSslSettingsHttpCompression(d, v, "http_compression")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-compression"] = t
		}
	}

	if v, ok := d.GetOk("http_only_cookie"); ok || d.HasChange("http_only_cookie") {
		t, err := expandVpnSslSettingsHttpOnlyCookie(d, v, "http_only_cookie")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-only-cookie"] = t
		}
	}

	if v, ok := d.GetOk("http_request_body_timeout"); ok || d.HasChange("http_request_body_timeout") {
		t, err := expandVpnSslSettingsHttpRequestBodyTimeout(d, v, "http_request_body_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-request-body-timeout"] = t
		}
	}

	if v, ok := d.GetOk("http_request_header_timeout"); ok || d.HasChange("http_request_header_timeout") {
		t, err := expandVpnSslSettingsHttpRequestHeaderTimeout(d, v, "http_request_header_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-request-header-timeout"] = t
		}
	}

	if v, ok := d.GetOk("https_redirect"); ok || d.HasChange("https_redirect") {
		t, err := expandVpnSslSettingsHttpsRedirect(d, v, "https_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-redirect"] = t
		}
	}

	if v, ok := d.GetOk("idle_timeout"); ok || d.HasChange("idle_timeout") {
		t, err := expandVpnSslSettingsIdleTimeout(d, v, "idle_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idle-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_dns_server1"); ok || d.HasChange("ipv6_dns_server1") {
		t, err := expandVpnSslSettingsIpv6DnsServer1(d, v, "ipv6_dns_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_dns_server2"); ok || d.HasChange("ipv6_dns_server2") {
		t, err := expandVpnSslSettingsIpv6DnsServer2(d, v, "ipv6_dns_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_wins_server1"); ok || d.HasChange("ipv6_wins_server1") {
		t, err := expandVpnSslSettingsIpv6WinsServer1(d, v, "ipv6_wins_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-wins-server1"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_wins_server2"); ok || d.HasChange("ipv6_wins_server2") {
		t, err := expandVpnSslSettingsIpv6WinsServer2(d, v, "ipv6_wins_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-wins-server2"] = t
		}
	}

	if v, ok := d.GetOk("login_attempt_limit"); ok || d.HasChange("login_attempt_limit") {
		t, err := expandVpnSslSettingsLoginAttemptLimit(d, v, "login_attempt_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-attempt-limit"] = t
		}
	}

	if v, ok := d.GetOk("login_block_time"); ok || d.HasChange("login_block_time") {
		t, err := expandVpnSslSettingsLoginBlockTime(d, v, "login_block_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-block-time"] = t
		}
	}

	if v, ok := d.GetOk("login_timeout"); ok || d.HasChange("login_timeout") {
		t, err := expandVpnSslSettingsLoginTimeout(d, v, "login_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-timeout"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandVpnSslSettingsPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("port_precedence"); ok || d.HasChange("port_precedence") {
		t, err := expandVpnSslSettingsPortPrecedence(d, v, "port_precedence")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-precedence"] = t
		}
	}

	if v, ok := d.GetOk("reqclientcert"); ok || d.HasChange("reqclientcert") {
		t, err := expandVpnSslSettingsReqclientcert(d, v, "reqclientcert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reqclientcert"] = t
		}
	}

	if v, ok := d.GetOk("saml_redirect_port"); ok || d.HasChange("saml_redirect_port") {
		t, err := expandVpnSslSettingsSamlRedirectPort(d, v, "saml_redirect_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["saml-redirect-port"] = t
		}
	}

	if v, ok := d.GetOk("server_hostname"); ok || d.HasChange("server_hostname") {
		t, err := expandVpnSslSettingsServerHostname(d, v, "server_hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-hostname"] = t
		}
	}

	if v, ok := d.GetOk("servercert"); ok || d.HasChange("servercert") {
		t, err := expandVpnSslSettingsServercert(d, v, "servercert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["servercert"] = t
		}
	}

	if v, ok := d.GetOk("source_address"); ok || d.HasChange("source_address") {
		t, err := expandVpnSslSettingsSourceAddress(d, v, "source_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-address"] = t
		}
	}

	if v, ok := d.GetOk("source_address_negate"); ok || d.HasChange("source_address_negate") {
		t, err := expandVpnSslSettingsSourceAddressNegate(d, v, "source_address_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-address-negate"] = t
		}
	}

	if v, ok := d.GetOk("source_address6"); ok || d.HasChange("source_address6") {
		t, err := expandVpnSslSettingsSourceAddress6(d, v, "source_address6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-address6"] = t
		}
	}

	if v, ok := d.GetOk("source_address6_negate"); ok || d.HasChange("source_address6_negate") {
		t, err := expandVpnSslSettingsSourceAddress6Negate(d, v, "source_address6_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-address6-negate"] = t
		}
	}

	if v, ok := d.GetOk("source_interface"); ok || d.HasChange("source_interface") {
		t, err := expandVpnSslSettingsSourceInterface(d, v, "source_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-interface"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_renegotiation"); ok || d.HasChange("ssl_client_renegotiation") {
		t, err := expandVpnSslSettingsSslClientRenegotiation(d, v, "ssl_client_renegotiation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-renegotiation"] = t
		}
	}

	if v, ok := d.GetOk("ssl_insert_empty_fragment"); ok || d.HasChange("ssl_insert_empty_fragment") {
		t, err := expandVpnSslSettingsSslInsertEmptyFragment(d, v, "ssl_insert_empty_fragment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-insert-empty-fragment"] = t
		}
	}

	if v, ok := d.GetOk("ssl_max_proto_ver"); ok || d.HasChange("ssl_max_proto_ver") {
		t, err := expandVpnSslSettingsSslMaxProtoVer(d, v, "ssl_max_proto_ver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-max-proto-ver"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_ver"); ok || d.HasChange("ssl_min_proto_ver") {
		t, err := expandVpnSslSettingsSslMinProtoVer(d, v, "ssl_min_proto_ver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-proto-ver"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandVpnSslSettingsStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("transform_backward_slashes"); ok || d.HasChange("transform_backward_slashes") {
		t, err := expandVpnSslSettingsTransformBackwardSlashes(d, v, "transform_backward_slashes")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transform-backward-slashes"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_addr_assigned_method"); ok || d.HasChange("tunnel_addr_assigned_method") {
		t, err := expandVpnSslSettingsTunnelAddrAssignedMethod(d, v, "tunnel_addr_assigned_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-addr-assigned-method"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_connect_without_reauth"); ok || d.HasChange("tunnel_connect_without_reauth") {
		t, err := expandVpnSslSettingsTunnelConnectWithoutReauth(d, v, "tunnel_connect_without_reauth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-connect-without-reauth"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_ip_pools"); ok || d.HasChange("tunnel_ip_pools") {
		t, err := expandVpnSslSettingsTunnelIpPools(d, v, "tunnel_ip_pools")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-ip-pools"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_ipv6_pools"); ok || d.HasChange("tunnel_ipv6_pools") {
		t, err := expandVpnSslSettingsTunnelIpv6Pools(d, v, "tunnel_ipv6_pools")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-ipv6-pools"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_user_session_timeout"); ok || d.HasChange("tunnel_user_session_timeout") {
		t, err := expandVpnSslSettingsTunnelUserSessionTimeout(d, v, "tunnel_user_session_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-user-session-timeout"] = t
		}
	}

	if v, ok := d.GetOk("unsafe_legacy_renegotiation"); ok || d.HasChange("unsafe_legacy_renegotiation") {
		t, err := expandVpnSslSettingsUnsafeLegacyRenegotiation(d, v, "unsafe_legacy_renegotiation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unsafe-legacy-renegotiation"] = t
		}
	}

	if v, ok := d.GetOk("url_obscuration"); ok || d.HasChange("url_obscuration") {
		t, err := expandVpnSslSettingsUrlObscuration(d, v, "url_obscuration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-obscuration"] = t
		}
	}

	if v, ok := d.GetOk("user_peer"); ok || d.HasChange("user_peer") {
		t, err := expandVpnSslSettingsUserPeer(d, v, "user_peer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-peer"] = t
		}
	}

	if v, ok := d.GetOk("web_mode_snat"); ok || d.HasChange("web_mode_snat") {
		t, err := expandVpnSslSettingsWebModeSnat(d, v, "web_mode_snat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-mode-snat"] = t
		}
	}

	if v, ok := d.GetOk("wins_server1"); ok || d.HasChange("wins_server1") {
		t, err := expandVpnSslSettingsWinsServer1(d, v, "wins_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wins-server1"] = t
		}
	}

	if v, ok := d.GetOk("wins_server2"); ok || d.HasChange("wins_server2") {
		t, err := expandVpnSslSettingsWinsServer2(d, v, "wins_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wins-server2"] = t
		}
	}

	if v, ok := d.GetOk("x_content_type_options"); ok || d.HasChange("x_content_type_options") {
		t, err := expandVpnSslSettingsXContentTypeOptions(d, v, "x_content_type_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["x-content-type-options"] = t
		}
	}

	if v, ok := d.GetOk("ztna_trusted_client"); ok || d.HasChange("ztna_trusted_client") {
		t, err := expandVpnSslSettingsZtnaTrustedClient(d, v, "ztna_trusted_client")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna-trusted-client"] = t
		}
	}

	return &obj, nil
}
