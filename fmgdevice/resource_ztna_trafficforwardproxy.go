// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure ZTNA traffic forward proxy.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceZtnaTrafficForwardProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceZtnaTrafficForwardProxyCreate,
		Read:   resourceZtnaTrafficForwardProxyRead,
		Update: resourceZtnaTrafficForwardProxyUpdate,
		Delete: resourceZtnaTrafficForwardProxyDelete,

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
			"auth_portal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"empty_cert_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"h3_support": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"log_blocked_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"quic": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ack_delay_exponent": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"active_connection_id_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"active_migration": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"grease_quic_bit": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_ack_delay": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"max_datagram_frame_size": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"max_idle_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"max_udp_payload_size": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ssl_accept_ffdhe_groups": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_certificate": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssl_cipher_suites": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cipher": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"versions": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ssl_client_fallback": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_client_rekey_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ssl_client_renegotiation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_client_session_state_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ssl_client_session_state_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ssl_client_session_state_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_dh_bits": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_hpkp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_hpkp_age": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ssl_hpkp_backup": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssl_hpkp_include_subdomains": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_hpkp_primary": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssl_hpkp_report_uri": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_hsts": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_hsts_age": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ssl_hsts_include_subdomains": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_http_location_conversion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_http_match_host": &schema.Schema{
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
			"ssl_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_pfs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_send_empty_frags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_server_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_server_cipher_suites": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cipher": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"versions": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ssl_server_max_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_server_min_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_server_renegotiation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_server_session_state_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ssl_server_session_state_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ssl_server_session_state_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"svr_pool_multiplex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"svr_pool_server_max_concurrent_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"svr_pool_server_max_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"svr_pool_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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

func resourceZtnaTrafficForwardProxyCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectZtnaTrafficForwardProxy(d)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaTrafficForwardProxy resource while getting object: %v", err)
	}

	_, err = c.CreateZtnaTrafficForwardProxy(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ZtnaTrafficForwardProxy resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceZtnaTrafficForwardProxyRead(d, m)
}

func resourceZtnaTrafficForwardProxyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectZtnaTrafficForwardProxy(d)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaTrafficForwardProxy resource while getting object: %v", err)
	}

	_, err = c.UpdateZtnaTrafficForwardProxy(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaTrafficForwardProxy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceZtnaTrafficForwardProxyRead(d, m)
}

func resourceZtnaTrafficForwardProxyDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteZtnaTrafficForwardProxy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ZtnaTrafficForwardProxy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceZtnaTrafficForwardProxyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadZtnaTrafficForwardProxy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaTrafficForwardProxy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectZtnaTrafficForwardProxy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaTrafficForwardProxy resource from API: %v", err)
	}
	return nil
}

func flattenZtnaTrafficForwardProxyAuthPortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyEmptyCertAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyH3Support(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaTrafficForwardProxyLogBlockedTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyQuic(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := i["ack-delay-exponent"]; ok {
		result["ack_delay_exponent"] = flattenZtnaTrafficForwardProxyQuicAckDelayExponent(i["ack-delay-exponent"], d, pre_append)
	}

	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := i["active-connection-id-limit"]; ok {
		result["active_connection_id_limit"] = flattenZtnaTrafficForwardProxyQuicActiveConnectionIdLimit(i["active-connection-id-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "active_migration"
	if _, ok := i["active-migration"]; ok {
		result["active_migration"] = flattenZtnaTrafficForwardProxyQuicActiveMigration(i["active-migration"], d, pre_append)
	}

	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := i["grease-quic-bit"]; ok {
		result["grease_quic_bit"] = flattenZtnaTrafficForwardProxyQuicGreaseQuicBit(i["grease-quic-bit"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := i["max-ack-delay"]; ok {
		result["max_ack_delay"] = flattenZtnaTrafficForwardProxyQuicMaxAckDelay(i["max-ack-delay"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := i["max-datagram-frame-size"]; ok {
		result["max_datagram_frame_size"] = flattenZtnaTrafficForwardProxyQuicMaxDatagramFrameSize(i["max-datagram-frame-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := i["max-idle-timeout"]; ok {
		result["max_idle_timeout"] = flattenZtnaTrafficForwardProxyQuicMaxIdleTimeout(i["max-idle-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := i["max-udp-payload-size"]; ok {
		result["max_udp_payload_size"] = flattenZtnaTrafficForwardProxyQuicMaxUdpPayloadSize(i["max-udp-payload-size"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenZtnaTrafficForwardProxyQuicAckDelayExponent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyQuicActiveConnectionIdLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyQuicActiveMigration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyQuicGreaseQuicBit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyQuicMaxAckDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyQuicMaxDatagramFrameSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyQuicMaxIdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyQuicMaxUdpPayloadSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslAcceptFfdheGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaTrafficForwardProxySslCipherSuites(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := i["cipher"]; ok {
			v := flattenZtnaTrafficForwardProxySslCipherSuitesCipher(i["cipher"], d, pre_append)
			tmp["cipher"] = fortiAPISubPartPatch(v, "ZtnaTrafficForwardProxy-SslCipherSuites-Cipher")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenZtnaTrafficForwardProxySslCipherSuitesPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "ZtnaTrafficForwardProxy-SslCipherSuites-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {
			v := flattenZtnaTrafficForwardProxySslCipherSuitesVersions(i["versions"], d, pre_append)
			tmp["versions"] = fortiAPISubPartPatch(v, "ZtnaTrafficForwardProxy-SslCipherSuites-Versions")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenZtnaTrafficForwardProxySslCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaTrafficForwardProxySslClientFallback(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslClientRekeyCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslClientRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslClientSessionStateMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslClientSessionStateTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslClientSessionStateType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslDhBits(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslHpkp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslHpkpAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslHpkpBackup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaTrafficForwardProxySslHpkpIncludeSubdomains(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslHpkpPrimary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaTrafficForwardProxySslHpkpReportUri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslHsts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslHstsAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslHstsIncludeSubdomains(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslHttpLocationConversion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslHttpMatchHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslPfs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslSendEmptyFrags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslServerAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslServerCipherSuites(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := i["cipher"]; ok {
			v := flattenZtnaTrafficForwardProxySslServerCipherSuitesCipher(i["cipher"], d, pre_append)
			tmp["cipher"] = fortiAPISubPartPatch(v, "ZtnaTrafficForwardProxy-SslServerCipherSuites-Cipher")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenZtnaTrafficForwardProxySslServerCipherSuitesPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "ZtnaTrafficForwardProxy-SslServerCipherSuites-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {
			v := flattenZtnaTrafficForwardProxySslServerCipherSuitesVersions(i["versions"], d, pre_append)
			tmp["versions"] = fortiAPISubPartPatch(v, "ZtnaTrafficForwardProxy-SslServerCipherSuites-Versions")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenZtnaTrafficForwardProxySslServerCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslServerCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslServerCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaTrafficForwardProxySslServerMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslServerMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslServerRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslServerSessionStateMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslServerSessionStateTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslServerSessionStateType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySvrPoolMultiplex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySvrPoolServerMaxConcurrentRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySvrPoolServerMaxRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySvrPoolTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyUserAgentDetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectZtnaTrafficForwardProxy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("auth_portal", flattenZtnaTrafficForwardProxyAuthPortal(o["auth-portal"], d, "auth_portal")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-portal"], "ZtnaTrafficForwardProxy-AuthPortal"); ok {
			if err = d.Set("auth_portal", vv); err != nil {
				return fmt.Errorf("Error reading auth_portal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_portal: %v", err)
		}
	}

	if err = d.Set("client_cert", flattenZtnaTrafficForwardProxyClientCert(o["client-cert"], d, "client_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-cert"], "ZtnaTrafficForwardProxy-ClientCert"); ok {
			if err = d.Set("client_cert", vv); err != nil {
				return fmt.Errorf("Error reading client_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_cert: %v", err)
		}
	}

	if err = d.Set("comment", flattenZtnaTrafficForwardProxyComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ZtnaTrafficForwardProxy-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("empty_cert_action", flattenZtnaTrafficForwardProxyEmptyCertAction(o["empty-cert-action"], d, "empty_cert_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["empty-cert-action"], "ZtnaTrafficForwardProxy-EmptyCertAction"); ok {
			if err = d.Set("empty_cert_action", vv); err != nil {
				return fmt.Errorf("Error reading empty_cert_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading empty_cert_action: %v", err)
		}
	}

	if err = d.Set("h3_support", flattenZtnaTrafficForwardProxyH3Support(o["h3-support"], d, "h3_support")); err != nil {
		if vv, ok := fortiAPIPatch(o["h3-support"], "ZtnaTrafficForwardProxy-H3Support"); ok {
			if err = d.Set("h3_support", vv); err != nil {
				return fmt.Errorf("Error reading h3_support: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading h3_support: %v", err)
		}
	}

	if err = d.Set("interface", flattenZtnaTrafficForwardProxyInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "ZtnaTrafficForwardProxy-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("log_blocked_traffic", flattenZtnaTrafficForwardProxyLogBlockedTraffic(o["log-blocked-traffic"], d, "log_blocked_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-blocked-traffic"], "ZtnaTrafficForwardProxy-LogBlockedTraffic"); ok {
			if err = d.Set("log_blocked_traffic", vv); err != nil {
				return fmt.Errorf("Error reading log_blocked_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_blocked_traffic: %v", err)
		}
	}

	if err = d.Set("name", flattenZtnaTrafficForwardProxyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ZtnaTrafficForwardProxy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenZtnaTrafficForwardProxyPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ZtnaTrafficForwardProxy-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("quic", flattenZtnaTrafficForwardProxyQuic(o["quic"], d, "quic")); err != nil {
			if vv, ok := fortiAPIPatch(o["quic"], "ZtnaTrafficForwardProxy-Quic"); ok {
				if err = d.Set("quic", vv); err != nil {
					return fmt.Errorf("Error reading quic: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading quic: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("quic"); ok {
			if err = d.Set("quic", flattenZtnaTrafficForwardProxyQuic(o["quic"], d, "quic")); err != nil {
				if vv, ok := fortiAPIPatch(o["quic"], "ZtnaTrafficForwardProxy-Quic"); ok {
					if err = d.Set("quic", vv); err != nil {
						return fmt.Errorf("Error reading quic: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading quic: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_accept_ffdhe_groups", flattenZtnaTrafficForwardProxySslAcceptFfdheGroups(o["ssl-accept-ffdhe-groups"], d, "ssl_accept_ffdhe_groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-accept-ffdhe-groups"], "ZtnaTrafficForwardProxy-SslAcceptFfdheGroups"); ok {
			if err = d.Set("ssl_accept_ffdhe_groups", vv); err != nil {
				return fmt.Errorf("Error reading ssl_accept_ffdhe_groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_accept_ffdhe_groups: %v", err)
		}
	}

	if err = d.Set("ssl_algorithm", flattenZtnaTrafficForwardProxySslAlgorithm(o["ssl-algorithm"], d, "ssl_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-algorithm"], "ZtnaTrafficForwardProxy-SslAlgorithm"); ok {
			if err = d.Set("ssl_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading ssl_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_algorithm: %v", err)
		}
	}

	if err = d.Set("ssl_certificate", flattenZtnaTrafficForwardProxySslCertificate(o["ssl-certificate"], d, "ssl_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-certificate"], "ZtnaTrafficForwardProxy-SslCertificate"); ok {
			if err = d.Set("ssl_certificate", vv); err != nil {
				return fmt.Errorf("Error reading ssl_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_certificate: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ssl_cipher_suites", flattenZtnaTrafficForwardProxySslCipherSuites(o["ssl-cipher-suites"], d, "ssl_cipher_suites")); err != nil {
			if vv, ok := fortiAPIPatch(o["ssl-cipher-suites"], "ZtnaTrafficForwardProxy-SslCipherSuites"); ok {
				if err = d.Set("ssl_cipher_suites", vv); err != nil {
					return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_cipher_suites"); ok {
			if err = d.Set("ssl_cipher_suites", flattenZtnaTrafficForwardProxySslCipherSuites(o["ssl-cipher-suites"], d, "ssl_cipher_suites")); err != nil {
				if vv, ok := fortiAPIPatch(o["ssl-cipher-suites"], "ZtnaTrafficForwardProxy-SslCipherSuites"); ok {
					if err = d.Set("ssl_cipher_suites", vv); err != nil {
						return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_client_fallback", flattenZtnaTrafficForwardProxySslClientFallback(o["ssl-client-fallback"], d, "ssl_client_fallback")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-fallback"], "ZtnaTrafficForwardProxy-SslClientFallback"); ok {
			if err = d.Set("ssl_client_fallback", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_fallback: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_fallback: %v", err)
		}
	}

	if err = d.Set("ssl_client_rekey_count", flattenZtnaTrafficForwardProxySslClientRekeyCount(o["ssl-client-rekey-count"], d, "ssl_client_rekey_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-rekey-count"], "ZtnaTrafficForwardProxy-SslClientRekeyCount"); ok {
			if err = d.Set("ssl_client_rekey_count", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_rekey_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_rekey_count: %v", err)
		}
	}

	if err = d.Set("ssl_client_renegotiation", flattenZtnaTrafficForwardProxySslClientRenegotiation(o["ssl-client-renegotiation"], d, "ssl_client_renegotiation")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-renegotiation"], "ZtnaTrafficForwardProxy-SslClientRenegotiation"); ok {
			if err = d.Set("ssl_client_renegotiation", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_renegotiation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_renegotiation: %v", err)
		}
	}

	if err = d.Set("ssl_client_session_state_max", flattenZtnaTrafficForwardProxySslClientSessionStateMax(o["ssl-client-session-state-max"], d, "ssl_client_session_state_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-session-state-max"], "ZtnaTrafficForwardProxy-SslClientSessionStateMax"); ok {
			if err = d.Set("ssl_client_session_state_max", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_session_state_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_session_state_max: %v", err)
		}
	}

	if err = d.Set("ssl_client_session_state_timeout", flattenZtnaTrafficForwardProxySslClientSessionStateTimeout(o["ssl-client-session-state-timeout"], d, "ssl_client_session_state_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-session-state-timeout"], "ZtnaTrafficForwardProxy-SslClientSessionStateTimeout"); ok {
			if err = d.Set("ssl_client_session_state_timeout", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_session_state_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_session_state_timeout: %v", err)
		}
	}

	if err = d.Set("ssl_client_session_state_type", flattenZtnaTrafficForwardProxySslClientSessionStateType(o["ssl-client-session-state-type"], d, "ssl_client_session_state_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-session-state-type"], "ZtnaTrafficForwardProxy-SslClientSessionStateType"); ok {
			if err = d.Set("ssl_client_session_state_type", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_session_state_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_session_state_type: %v", err)
		}
	}

	if err = d.Set("ssl_dh_bits", flattenZtnaTrafficForwardProxySslDhBits(o["ssl-dh-bits"], d, "ssl_dh_bits")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-dh-bits"], "ZtnaTrafficForwardProxy-SslDhBits"); ok {
			if err = d.Set("ssl_dh_bits", vv); err != nil {
				return fmt.Errorf("Error reading ssl_dh_bits: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_dh_bits: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp", flattenZtnaTrafficForwardProxySslHpkp(o["ssl-hpkp"], d, "ssl_hpkp")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hpkp"], "ZtnaTrafficForwardProxy-SslHpkp"); ok {
			if err = d.Set("ssl_hpkp", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hpkp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hpkp: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_age", flattenZtnaTrafficForwardProxySslHpkpAge(o["ssl-hpkp-age"], d, "ssl_hpkp_age")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hpkp-age"], "ZtnaTrafficForwardProxy-SslHpkpAge"); ok {
			if err = d.Set("ssl_hpkp_age", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hpkp_age: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hpkp_age: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_backup", flattenZtnaTrafficForwardProxySslHpkpBackup(o["ssl-hpkp-backup"], d, "ssl_hpkp_backup")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hpkp-backup"], "ZtnaTrafficForwardProxy-SslHpkpBackup"); ok {
			if err = d.Set("ssl_hpkp_backup", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hpkp_backup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hpkp_backup: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_include_subdomains", flattenZtnaTrafficForwardProxySslHpkpIncludeSubdomains(o["ssl-hpkp-include-subdomains"], d, "ssl_hpkp_include_subdomains")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hpkp-include-subdomains"], "ZtnaTrafficForwardProxy-SslHpkpIncludeSubdomains"); ok {
			if err = d.Set("ssl_hpkp_include_subdomains", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hpkp_include_subdomains: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hpkp_include_subdomains: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_primary", flattenZtnaTrafficForwardProxySslHpkpPrimary(o["ssl-hpkp-primary"], d, "ssl_hpkp_primary")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hpkp-primary"], "ZtnaTrafficForwardProxy-SslHpkpPrimary"); ok {
			if err = d.Set("ssl_hpkp_primary", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hpkp_primary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hpkp_primary: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_report_uri", flattenZtnaTrafficForwardProxySslHpkpReportUri(o["ssl-hpkp-report-uri"], d, "ssl_hpkp_report_uri")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hpkp-report-uri"], "ZtnaTrafficForwardProxy-SslHpkpReportUri"); ok {
			if err = d.Set("ssl_hpkp_report_uri", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hpkp_report_uri: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hpkp_report_uri: %v", err)
		}
	}

	if err = d.Set("ssl_hsts", flattenZtnaTrafficForwardProxySslHsts(o["ssl-hsts"], d, "ssl_hsts")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hsts"], "ZtnaTrafficForwardProxy-SslHsts"); ok {
			if err = d.Set("ssl_hsts", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hsts: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hsts: %v", err)
		}
	}

	if err = d.Set("ssl_hsts_age", flattenZtnaTrafficForwardProxySslHstsAge(o["ssl-hsts-age"], d, "ssl_hsts_age")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hsts-age"], "ZtnaTrafficForwardProxy-SslHstsAge"); ok {
			if err = d.Set("ssl_hsts_age", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hsts_age: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hsts_age: %v", err)
		}
	}

	if err = d.Set("ssl_hsts_include_subdomains", flattenZtnaTrafficForwardProxySslHstsIncludeSubdomains(o["ssl-hsts-include-subdomains"], d, "ssl_hsts_include_subdomains")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hsts-include-subdomains"], "ZtnaTrafficForwardProxy-SslHstsIncludeSubdomains"); ok {
			if err = d.Set("ssl_hsts_include_subdomains", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hsts_include_subdomains: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hsts_include_subdomains: %v", err)
		}
	}

	if err = d.Set("ssl_http_location_conversion", flattenZtnaTrafficForwardProxySslHttpLocationConversion(o["ssl-http-location-conversion"], d, "ssl_http_location_conversion")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-http-location-conversion"], "ZtnaTrafficForwardProxy-SslHttpLocationConversion"); ok {
			if err = d.Set("ssl_http_location_conversion", vv); err != nil {
				return fmt.Errorf("Error reading ssl_http_location_conversion: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_http_location_conversion: %v", err)
		}
	}

	if err = d.Set("ssl_http_match_host", flattenZtnaTrafficForwardProxySslHttpMatchHost(o["ssl-http-match-host"], d, "ssl_http_match_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-http-match-host"], "ZtnaTrafficForwardProxy-SslHttpMatchHost"); ok {
			if err = d.Set("ssl_http_match_host", vv); err != nil {
				return fmt.Errorf("Error reading ssl_http_match_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_http_match_host: %v", err)
		}
	}

	if err = d.Set("ssl_max_version", flattenZtnaTrafficForwardProxySslMaxVersion(o["ssl-max-version"], d, "ssl_max_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-max-version"], "ZtnaTrafficForwardProxy-SslMaxVersion"); ok {
			if err = d.Set("ssl_max_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_max_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_max_version: %v", err)
		}
	}

	if err = d.Set("ssl_min_version", flattenZtnaTrafficForwardProxySslMinVersion(o["ssl-min-version"], d, "ssl_min_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-min-version"], "ZtnaTrafficForwardProxy-SslMinVersion"); ok {
			if err = d.Set("ssl_min_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_min_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_min_version: %v", err)
		}
	}

	if err = d.Set("ssl_mode", flattenZtnaTrafficForwardProxySslMode(o["ssl-mode"], d, "ssl_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-mode"], "ZtnaTrafficForwardProxy-SslMode"); ok {
			if err = d.Set("ssl_mode", vv); err != nil {
				return fmt.Errorf("Error reading ssl_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_mode: %v", err)
		}
	}

	if err = d.Set("ssl_pfs", flattenZtnaTrafficForwardProxySslPfs(o["ssl-pfs"], d, "ssl_pfs")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-pfs"], "ZtnaTrafficForwardProxy-SslPfs"); ok {
			if err = d.Set("ssl_pfs", vv); err != nil {
				return fmt.Errorf("Error reading ssl_pfs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_pfs: %v", err)
		}
	}

	if err = d.Set("ssl_send_empty_frags", flattenZtnaTrafficForwardProxySslSendEmptyFrags(o["ssl-send-empty-frags"], d, "ssl_send_empty_frags")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-send-empty-frags"], "ZtnaTrafficForwardProxy-SslSendEmptyFrags"); ok {
			if err = d.Set("ssl_send_empty_frags", vv); err != nil {
				return fmt.Errorf("Error reading ssl_send_empty_frags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_send_empty_frags: %v", err)
		}
	}

	if err = d.Set("ssl_server_algorithm", flattenZtnaTrafficForwardProxySslServerAlgorithm(o["ssl-server-algorithm"], d, "ssl_server_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-algorithm"], "ZtnaTrafficForwardProxy-SslServerAlgorithm"); ok {
			if err = d.Set("ssl_server_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_algorithm: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ssl_server_cipher_suites", flattenZtnaTrafficForwardProxySslServerCipherSuites(o["ssl-server-cipher-suites"], d, "ssl_server_cipher_suites")); err != nil {
			if vv, ok := fortiAPIPatch(o["ssl-server-cipher-suites"], "ZtnaTrafficForwardProxy-SslServerCipherSuites"); ok {
				if err = d.Set("ssl_server_cipher_suites", vv); err != nil {
					return fmt.Errorf("Error reading ssl_server_cipher_suites: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ssl_server_cipher_suites: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_server_cipher_suites"); ok {
			if err = d.Set("ssl_server_cipher_suites", flattenZtnaTrafficForwardProxySslServerCipherSuites(o["ssl-server-cipher-suites"], d, "ssl_server_cipher_suites")); err != nil {
				if vv, ok := fortiAPIPatch(o["ssl-server-cipher-suites"], "ZtnaTrafficForwardProxy-SslServerCipherSuites"); ok {
					if err = d.Set("ssl_server_cipher_suites", vv); err != nil {
						return fmt.Errorf("Error reading ssl_server_cipher_suites: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ssl_server_cipher_suites: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_server_max_version", flattenZtnaTrafficForwardProxySslServerMaxVersion(o["ssl-server-max-version"], d, "ssl_server_max_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-max-version"], "ZtnaTrafficForwardProxy-SslServerMaxVersion"); ok {
			if err = d.Set("ssl_server_max_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_max_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_max_version: %v", err)
		}
	}

	if err = d.Set("ssl_server_min_version", flattenZtnaTrafficForwardProxySslServerMinVersion(o["ssl-server-min-version"], d, "ssl_server_min_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-min-version"], "ZtnaTrafficForwardProxy-SslServerMinVersion"); ok {
			if err = d.Set("ssl_server_min_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_min_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_min_version: %v", err)
		}
	}

	if err = d.Set("ssl_server_renegotiation", flattenZtnaTrafficForwardProxySslServerRenegotiation(o["ssl-server-renegotiation"], d, "ssl_server_renegotiation")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-renegotiation"], "ZtnaTrafficForwardProxy-SslServerRenegotiation"); ok {
			if err = d.Set("ssl_server_renegotiation", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_renegotiation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_renegotiation: %v", err)
		}
	}

	if err = d.Set("ssl_server_session_state_max", flattenZtnaTrafficForwardProxySslServerSessionStateMax(o["ssl-server-session-state-max"], d, "ssl_server_session_state_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-session-state-max"], "ZtnaTrafficForwardProxy-SslServerSessionStateMax"); ok {
			if err = d.Set("ssl_server_session_state_max", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_session_state_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_session_state_max: %v", err)
		}
	}

	if err = d.Set("ssl_server_session_state_timeout", flattenZtnaTrafficForwardProxySslServerSessionStateTimeout(o["ssl-server-session-state-timeout"], d, "ssl_server_session_state_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-session-state-timeout"], "ZtnaTrafficForwardProxy-SslServerSessionStateTimeout"); ok {
			if err = d.Set("ssl_server_session_state_timeout", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_session_state_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_session_state_timeout: %v", err)
		}
	}

	if err = d.Set("ssl_server_session_state_type", flattenZtnaTrafficForwardProxySslServerSessionStateType(o["ssl-server-session-state-type"], d, "ssl_server_session_state_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-session-state-type"], "ZtnaTrafficForwardProxy-SslServerSessionStateType"); ok {
			if err = d.Set("ssl_server_session_state_type", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_session_state_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_session_state_type: %v", err)
		}
	}

	if err = d.Set("status", flattenZtnaTrafficForwardProxyStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ZtnaTrafficForwardProxy-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("svr_pool_multiplex", flattenZtnaTrafficForwardProxySvrPoolMultiplex(o["svr-pool-multiplex"], d, "svr_pool_multiplex")); err != nil {
		if vv, ok := fortiAPIPatch(o["svr-pool-multiplex"], "ZtnaTrafficForwardProxy-SvrPoolMultiplex"); ok {
			if err = d.Set("svr_pool_multiplex", vv); err != nil {
				return fmt.Errorf("Error reading svr_pool_multiplex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading svr_pool_multiplex: %v", err)
		}
	}

	if err = d.Set("svr_pool_server_max_concurrent_request", flattenZtnaTrafficForwardProxySvrPoolServerMaxConcurrentRequest(o["svr-pool-server-max-concurrent-request"], d, "svr_pool_server_max_concurrent_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["svr-pool-server-max-concurrent-request"], "ZtnaTrafficForwardProxy-SvrPoolServerMaxConcurrentRequest"); ok {
			if err = d.Set("svr_pool_server_max_concurrent_request", vv); err != nil {
				return fmt.Errorf("Error reading svr_pool_server_max_concurrent_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading svr_pool_server_max_concurrent_request: %v", err)
		}
	}

	if err = d.Set("svr_pool_server_max_request", flattenZtnaTrafficForwardProxySvrPoolServerMaxRequest(o["svr-pool-server-max-request"], d, "svr_pool_server_max_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["svr-pool-server-max-request"], "ZtnaTrafficForwardProxy-SvrPoolServerMaxRequest"); ok {
			if err = d.Set("svr_pool_server_max_request", vv); err != nil {
				return fmt.Errorf("Error reading svr_pool_server_max_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading svr_pool_server_max_request: %v", err)
		}
	}

	if err = d.Set("svr_pool_ttl", flattenZtnaTrafficForwardProxySvrPoolTtl(o["svr-pool-ttl"], d, "svr_pool_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["svr-pool-ttl"], "ZtnaTrafficForwardProxy-SvrPoolTtl"); ok {
			if err = d.Set("svr_pool_ttl", vv); err != nil {
				return fmt.Errorf("Error reading svr_pool_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading svr_pool_ttl: %v", err)
		}
	}

	if err = d.Set("user_agent_detect", flattenZtnaTrafficForwardProxyUserAgentDetect(o["user-agent-detect"], d, "user_agent_detect")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-agent-detect"], "ZtnaTrafficForwardProxy-UserAgentDetect"); ok {
			if err = d.Set("user_agent_detect", vv); err != nil {
				return fmt.Errorf("Error reading user_agent_detect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_agent_detect: %v", err)
		}
	}

	return nil
}

func flattenZtnaTrafficForwardProxyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandZtnaTrafficForwardProxyAuthPortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyClientCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyEmptyCertAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyH3Support(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaTrafficForwardProxyLogBlockedTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyQuic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ack-delay-exponent"], _ = expandZtnaTrafficForwardProxyQuicAckDelayExponent(d, i["ack_delay_exponent"], pre_append)
	}
	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["active-connection-id-limit"], _ = expandZtnaTrafficForwardProxyQuicActiveConnectionIdLimit(d, i["active_connection_id_limit"], pre_append)
	}
	pre_append = pre + ".0." + "active_migration"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["active-migration"], _ = expandZtnaTrafficForwardProxyQuicActiveMigration(d, i["active_migration"], pre_append)
	}
	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["grease-quic-bit"], _ = expandZtnaTrafficForwardProxyQuicGreaseQuicBit(d, i["grease_quic_bit"], pre_append)
	}
	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-ack-delay"], _ = expandZtnaTrafficForwardProxyQuicMaxAckDelay(d, i["max_ack_delay"], pre_append)
	}
	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-datagram-frame-size"], _ = expandZtnaTrafficForwardProxyQuicMaxDatagramFrameSize(d, i["max_datagram_frame_size"], pre_append)
	}
	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-idle-timeout"], _ = expandZtnaTrafficForwardProxyQuicMaxIdleTimeout(d, i["max_idle_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-udp-payload-size"], _ = expandZtnaTrafficForwardProxyQuicMaxUdpPayloadSize(d, i["max_udp_payload_size"], pre_append)
	}

	return result, nil
}

func expandZtnaTrafficForwardProxyQuicAckDelayExponent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyQuicActiveConnectionIdLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyQuicActiveMigration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyQuicGreaseQuicBit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyQuicMaxAckDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyQuicMaxDatagramFrameSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyQuicMaxIdleTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyQuicMaxUdpPayloadSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslAcceptFfdheGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaTrafficForwardProxySslCipherSuites(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cipher"], _ = expandZtnaTrafficForwardProxySslCipherSuitesCipher(d, i["cipher"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandZtnaTrafficForwardProxySslCipherSuitesPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["versions"], _ = expandZtnaTrafficForwardProxySslCipherSuitesVersions(d, i["versions"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandZtnaTrafficForwardProxySslCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaTrafficForwardProxySslClientFallback(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslClientRekeyCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslClientRenegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslClientSessionStateMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslClientSessionStateTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslClientSessionStateType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslDhBits(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslHpkp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslHpkpAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslHpkpBackup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaTrafficForwardProxySslHpkpIncludeSubdomains(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslHpkpPrimary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaTrafficForwardProxySslHpkpReportUri(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslHsts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslHstsAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslHstsIncludeSubdomains(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslHttpLocationConversion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslHttpMatchHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslPfs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslSendEmptyFrags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslServerAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslServerCipherSuites(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cipher"], _ = expandZtnaTrafficForwardProxySslServerCipherSuitesCipher(d, i["cipher"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandZtnaTrafficForwardProxySslServerCipherSuitesPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["versions"], _ = expandZtnaTrafficForwardProxySslServerCipherSuitesVersions(d, i["versions"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandZtnaTrafficForwardProxySslServerCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslServerCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslServerCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaTrafficForwardProxySslServerMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslServerMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslServerRenegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslServerSessionStateMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslServerSessionStateTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslServerSessionStateType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySvrPoolMultiplex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySvrPoolServerMaxConcurrentRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySvrPoolServerMaxRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySvrPoolTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyUserAgentDetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectZtnaTrafficForwardProxy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_portal"); ok || d.HasChange("auth_portal") {
		t, err := expandZtnaTrafficForwardProxyAuthPortal(d, v, "auth_portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-portal"] = t
		}
	}

	if v, ok := d.GetOk("client_cert"); ok || d.HasChange("client_cert") {
		t, err := expandZtnaTrafficForwardProxyClientCert(d, v, "client_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandZtnaTrafficForwardProxyComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("empty_cert_action"); ok || d.HasChange("empty_cert_action") {
		t, err := expandZtnaTrafficForwardProxyEmptyCertAction(d, v, "empty_cert_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["empty-cert-action"] = t
		}
	}

	if v, ok := d.GetOk("h3_support"); ok || d.HasChange("h3_support") {
		t, err := expandZtnaTrafficForwardProxyH3Support(d, v, "h3_support")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-support"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandZtnaTrafficForwardProxyInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("log_blocked_traffic"); ok || d.HasChange("log_blocked_traffic") {
		t, err := expandZtnaTrafficForwardProxyLogBlockedTraffic(d, v, "log_blocked_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-blocked-traffic"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandZtnaTrafficForwardProxyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandZtnaTrafficForwardProxyPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("quic"); ok || d.HasChange("quic") {
		t, err := expandZtnaTrafficForwardProxyQuic(d, v, "quic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quic"] = t
		}
	}

	if v, ok := d.GetOk("ssl_accept_ffdhe_groups"); ok || d.HasChange("ssl_accept_ffdhe_groups") {
		t, err := expandZtnaTrafficForwardProxySslAcceptFfdheGroups(d, v, "ssl_accept_ffdhe_groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-accept-ffdhe-groups"] = t
		}
	}

	if v, ok := d.GetOk("ssl_algorithm"); ok || d.HasChange("ssl_algorithm") {
		t, err := expandZtnaTrafficForwardProxySslAlgorithm(d, v, "ssl_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_certificate"); ok || d.HasChange("ssl_certificate") {
		t, err := expandZtnaTrafficForwardProxySslCertificate(d, v, "ssl_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-certificate"] = t
		}
	}

	if v, ok := d.GetOk("ssl_cipher_suites"); ok || d.HasChange("ssl_cipher_suites") {
		t, err := expandZtnaTrafficForwardProxySslCipherSuites(d, v, "ssl_cipher_suites")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-cipher-suites"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_fallback"); ok || d.HasChange("ssl_client_fallback") {
		t, err := expandZtnaTrafficForwardProxySslClientFallback(d, v, "ssl_client_fallback")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-fallback"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_rekey_count"); ok || d.HasChange("ssl_client_rekey_count") {
		t, err := expandZtnaTrafficForwardProxySslClientRekeyCount(d, v, "ssl_client_rekey_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-rekey-count"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_renegotiation"); ok || d.HasChange("ssl_client_renegotiation") {
		t, err := expandZtnaTrafficForwardProxySslClientRenegotiation(d, v, "ssl_client_renegotiation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-renegotiation"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_session_state_max"); ok || d.HasChange("ssl_client_session_state_max") {
		t, err := expandZtnaTrafficForwardProxySslClientSessionStateMax(d, v, "ssl_client_session_state_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-session-state-max"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_session_state_timeout"); ok || d.HasChange("ssl_client_session_state_timeout") {
		t, err := expandZtnaTrafficForwardProxySslClientSessionStateTimeout(d, v, "ssl_client_session_state_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-session-state-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_session_state_type"); ok || d.HasChange("ssl_client_session_state_type") {
		t, err := expandZtnaTrafficForwardProxySslClientSessionStateType(d, v, "ssl_client_session_state_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-session-state-type"] = t
		}
	}

	if v, ok := d.GetOk("ssl_dh_bits"); ok || d.HasChange("ssl_dh_bits") {
		t, err := expandZtnaTrafficForwardProxySslDhBits(d, v, "ssl_dh_bits")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-dh-bits"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp"); ok || d.HasChange("ssl_hpkp") {
		t, err := expandZtnaTrafficForwardProxySslHpkp(d, v, "ssl_hpkp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_age"); ok || d.HasChange("ssl_hpkp_age") {
		t, err := expandZtnaTrafficForwardProxySslHpkpAge(d, v, "ssl_hpkp_age")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-age"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_backup"); ok || d.HasChange("ssl_hpkp_backup") {
		t, err := expandZtnaTrafficForwardProxySslHpkpBackup(d, v, "ssl_hpkp_backup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-backup"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_include_subdomains"); ok || d.HasChange("ssl_hpkp_include_subdomains") {
		t, err := expandZtnaTrafficForwardProxySslHpkpIncludeSubdomains(d, v, "ssl_hpkp_include_subdomains")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-include-subdomains"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_primary"); ok || d.HasChange("ssl_hpkp_primary") {
		t, err := expandZtnaTrafficForwardProxySslHpkpPrimary(d, v, "ssl_hpkp_primary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-primary"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_report_uri"); ok || d.HasChange("ssl_hpkp_report_uri") {
		t, err := expandZtnaTrafficForwardProxySslHpkpReportUri(d, v, "ssl_hpkp_report_uri")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-report-uri"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hsts"); ok || d.HasChange("ssl_hsts") {
		t, err := expandZtnaTrafficForwardProxySslHsts(d, v, "ssl_hsts")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hsts"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hsts_age"); ok || d.HasChange("ssl_hsts_age") {
		t, err := expandZtnaTrafficForwardProxySslHstsAge(d, v, "ssl_hsts_age")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hsts-age"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hsts_include_subdomains"); ok || d.HasChange("ssl_hsts_include_subdomains") {
		t, err := expandZtnaTrafficForwardProxySslHstsIncludeSubdomains(d, v, "ssl_hsts_include_subdomains")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hsts-include-subdomains"] = t
		}
	}

	if v, ok := d.GetOk("ssl_http_location_conversion"); ok || d.HasChange("ssl_http_location_conversion") {
		t, err := expandZtnaTrafficForwardProxySslHttpLocationConversion(d, v, "ssl_http_location_conversion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-http-location-conversion"] = t
		}
	}

	if v, ok := d.GetOk("ssl_http_match_host"); ok || d.HasChange("ssl_http_match_host") {
		t, err := expandZtnaTrafficForwardProxySslHttpMatchHost(d, v, "ssl_http_match_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-http-match-host"] = t
		}
	}

	if v, ok := d.GetOk("ssl_max_version"); ok || d.HasChange("ssl_max_version") {
		t, err := expandZtnaTrafficForwardProxySslMaxVersion(d, v, "ssl_max_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-max-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_version"); ok || d.HasChange("ssl_min_version") {
		t, err := expandZtnaTrafficForwardProxySslMinVersion(d, v, "ssl_min_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_mode"); ok || d.HasChange("ssl_mode") {
		t, err := expandZtnaTrafficForwardProxySslMode(d, v, "ssl_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-mode"] = t
		}
	}

	if v, ok := d.GetOk("ssl_pfs"); ok || d.HasChange("ssl_pfs") {
		t, err := expandZtnaTrafficForwardProxySslPfs(d, v, "ssl_pfs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-pfs"] = t
		}
	}

	if v, ok := d.GetOk("ssl_send_empty_frags"); ok || d.HasChange("ssl_send_empty_frags") {
		t, err := expandZtnaTrafficForwardProxySslSendEmptyFrags(d, v, "ssl_send_empty_frags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-send-empty-frags"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_algorithm"); ok || d.HasChange("ssl_server_algorithm") {
		t, err := expandZtnaTrafficForwardProxySslServerAlgorithm(d, v, "ssl_server_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_cipher_suites"); ok || d.HasChange("ssl_server_cipher_suites") {
		t, err := expandZtnaTrafficForwardProxySslServerCipherSuites(d, v, "ssl_server_cipher_suites")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-cipher-suites"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_max_version"); ok || d.HasChange("ssl_server_max_version") {
		t, err := expandZtnaTrafficForwardProxySslServerMaxVersion(d, v, "ssl_server_max_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-max-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_min_version"); ok || d.HasChange("ssl_server_min_version") {
		t, err := expandZtnaTrafficForwardProxySslServerMinVersion(d, v, "ssl_server_min_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-min-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_renegotiation"); ok || d.HasChange("ssl_server_renegotiation") {
		t, err := expandZtnaTrafficForwardProxySslServerRenegotiation(d, v, "ssl_server_renegotiation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-renegotiation"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_session_state_max"); ok || d.HasChange("ssl_server_session_state_max") {
		t, err := expandZtnaTrafficForwardProxySslServerSessionStateMax(d, v, "ssl_server_session_state_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-session-state-max"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_session_state_timeout"); ok || d.HasChange("ssl_server_session_state_timeout") {
		t, err := expandZtnaTrafficForwardProxySslServerSessionStateTimeout(d, v, "ssl_server_session_state_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-session-state-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_session_state_type"); ok || d.HasChange("ssl_server_session_state_type") {
		t, err := expandZtnaTrafficForwardProxySslServerSessionStateType(d, v, "ssl_server_session_state_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-session-state-type"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandZtnaTrafficForwardProxyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("svr_pool_multiplex"); ok || d.HasChange("svr_pool_multiplex") {
		t, err := expandZtnaTrafficForwardProxySvrPoolMultiplex(d, v, "svr_pool_multiplex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["svr-pool-multiplex"] = t
		}
	}

	if v, ok := d.GetOk("svr_pool_server_max_concurrent_request"); ok || d.HasChange("svr_pool_server_max_concurrent_request") {
		t, err := expandZtnaTrafficForwardProxySvrPoolServerMaxConcurrentRequest(d, v, "svr_pool_server_max_concurrent_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["svr-pool-server-max-concurrent-request"] = t
		}
	}

	if v, ok := d.GetOk("svr_pool_server_max_request"); ok || d.HasChange("svr_pool_server_max_request") {
		t, err := expandZtnaTrafficForwardProxySvrPoolServerMaxRequest(d, v, "svr_pool_server_max_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["svr-pool-server-max-request"] = t
		}
	}

	if v, ok := d.GetOk("svr_pool_ttl"); ok || d.HasChange("svr_pool_ttl") {
		t, err := expandZtnaTrafficForwardProxySvrPoolTtl(d, v, "svr_pool_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["svr-pool-ttl"] = t
		}
	}

	if v, ok := d.GetOk("user_agent_detect"); ok || d.HasChange("user_agent_detect") {
		t, err := expandZtnaTrafficForwardProxyUserAgentDetect(d, v, "user_agent_detect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-agent-detect"] = t
		}
	}

	return &obj, nil
}
