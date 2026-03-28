// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Set IPv4 API Gateway.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallAccessProxyApiGateway() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallAccessProxyApiGatewayCreate,
		Read:   resourceFirewallAccessProxyApiGatewayRead,
		Update: resourceFirewallAccessProxyApiGatewayUpdate,
		Delete: resourceFirewallAccessProxyApiGatewayDelete,

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
			"access_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"application": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"h2_support": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"h3_support": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_cookie_age": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"http_cookie_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_cookie_domain_from_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_cookie_generation": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"http_cookie_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_cookie_share": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"https_cookie_secure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ldb_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"persistence": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"realservers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"addr_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"address": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"domain": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"external_auth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"health_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"health_check_proto": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"holddown_interval": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"http_host": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mappedport": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ssh_client_cert": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"ssh_host_key": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"ssh_host_key_validation": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"translate_host": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tunnel_encryption": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"verify_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"saml_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"saml_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_algorithm": &schema.Schema{
				Type:     schema.TypeString,
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
			"ssl_dh_bits": &schema.Schema{
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
			"ssl_renegotiation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_vpn_web_portal": &schema.Schema{
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
			"url_map_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"virtual_host": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceFirewallAccessProxyApiGatewayCreate(d *schema.ResourceData, m interface{}) error {
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
	access_proxy := d.Get("access_proxy").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_proxy"] = access_proxy

	obj, err := getObjectFirewallAccessProxyApiGateway(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallAccessProxyApiGateway resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallAccessProxyApiGateway(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallAccessProxyApiGateway(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallAccessProxyApiGateway resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateFirewallAccessProxyApiGateway(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallAccessProxyApiGateway resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceFirewallAccessProxyApiGatewayRead(d, m)
			} else {
				return fmt.Errorf("Error creating FirewallAccessProxyApiGateway resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallAccessProxyApiGatewayRead(d, m)
}

func resourceFirewallAccessProxyApiGatewayUpdate(d *schema.ResourceData, m interface{}) error {
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
	access_proxy := d.Get("access_proxy").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_proxy"] = access_proxy

	obj, err := getObjectFirewallAccessProxyApiGateway(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAccessProxyApiGateway resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallAccessProxyApiGateway(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAccessProxyApiGateway resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallAccessProxyApiGatewayRead(d, m)
}

func resourceFirewallAccessProxyApiGatewayDelete(d *schema.ResourceData, m interface{}) error {
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
	access_proxy := d.Get("access_proxy").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_proxy"] = access_proxy

	wsParams["adom"] = adomv

	err = c.DeleteFirewallAccessProxyApiGateway(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallAccessProxyApiGateway resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAccessProxyApiGatewayRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	access_proxy := d.Get("access_proxy").(string)
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
	if access_proxy == "" {
		access_proxy = importOptionChecking(m.(*FortiClient).Cfg, "access_proxy")
		if access_proxy == "" {
			return fmt.Errorf("Parameter access_proxy is missing")
		}
		if err = d.Set("access_proxy", access_proxy); err != nil {
			return fmt.Errorf("Error set params access_proxy: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_proxy"] = access_proxy

	o, err := c.ReadFirewallAccessProxyApiGateway(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallAccessProxyApiGateway resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallAccessProxyApiGateway(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAccessProxyApiGateway resource from API: %v", err)
	}
	return nil
}

func flattenFirewallAccessProxyApiGatewayApplication2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyApiGatewayH2Support2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayH3Support2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayHttpCookieAge2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayHttpCookieDomain2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayHttpCookieDomainFromHost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayHttpCookieGeneration2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayHttpCookiePath2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayHttpCookieShare2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayHttpsCookieSecure2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayLdbMethod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayPersistence2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayQuic2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := i["ack-delay-exponent"]; ok {
		result["ack_delay_exponent"] = flattenFirewallAccessProxyApiGatewayQuicAckDelayExponent2edl(i["ack-delay-exponent"], d, pre_append)
	}

	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := i["active-connection-id-limit"]; ok {
		result["active_connection_id_limit"] = flattenFirewallAccessProxyApiGatewayQuicActiveConnectionIdLimit2edl(i["active-connection-id-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "active_migration"
	if _, ok := i["active-migration"]; ok {
		result["active_migration"] = flattenFirewallAccessProxyApiGatewayQuicActiveMigration2edl(i["active-migration"], d, pre_append)
	}

	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := i["grease-quic-bit"]; ok {
		result["grease_quic_bit"] = flattenFirewallAccessProxyApiGatewayQuicGreaseQuicBit2edl(i["grease-quic-bit"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := i["max-ack-delay"]; ok {
		result["max_ack_delay"] = flattenFirewallAccessProxyApiGatewayQuicMaxAckDelay2edl(i["max-ack-delay"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := i["max-datagram-frame-size"]; ok {
		result["max_datagram_frame_size"] = flattenFirewallAccessProxyApiGatewayQuicMaxDatagramFrameSize2edl(i["max-datagram-frame-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := i["max-idle-timeout"]; ok {
		result["max_idle_timeout"] = flattenFirewallAccessProxyApiGatewayQuicMaxIdleTimeout2edl(i["max-idle-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := i["max-udp-payload-size"]; ok {
		result["max_udp_payload_size"] = flattenFirewallAccessProxyApiGatewayQuicMaxUdpPayloadSize2edl(i["max-udp-payload-size"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallAccessProxyApiGatewayQuicAckDelayExponent2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayQuicActiveConnectionIdLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayQuicActiveMigration2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayQuicGreaseQuicBit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayQuicMaxAckDelay2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayQuicMaxDatagramFrameSize2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayQuicMaxIdleTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayQuicMaxUdpPayloadSize2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealservers2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := i["addr-type"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversAddrType2edl(i["addr-type"], d, pre_append)
			tmp["addr_type"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-AddrType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversAddress2edl(i["address"], d, pre_append)
			tmp["address"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := i["domain"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversDomain2edl(i["domain"], d, pre_append)
			tmp["domain"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-Domain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_auth"
		if _, ok := i["external-auth"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversExternalAuth2edl(i["external-auth"], d, pre_append)
			tmp["external_auth"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-ExternalAuth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := i["health-check"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversHealthCheck2edl(i["health-check"], d, pre_append)
			tmp["health_check"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-HealthCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := i["health-check-proto"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversHealthCheckProto2edl(i["health-check-proto"], d, pre_append)
			tmp["health_check_proto"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-HealthCheckProto")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := i["holddown-interval"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversHolddownInterval2edl(i["holddown-interval"], d, pre_append)
			tmp["holddown_interval"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-HolddownInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := i["http-host"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversHttpHost2edl(i["http-host"], d, pre_append)
			tmp["http_host"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-HttpHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversIp2edl(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := i["mappedport"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversMappedport2edl(i["mappedport"], d, pre_append)
			tmp["mappedport"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-Mappedport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversPort2edl(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_client_cert"
		if _, ok := i["ssh-client-cert"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversSshClientCert2edl(i["ssh-client-cert"], d, pre_append)
			tmp["ssh_client_cert"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-SshClientCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key"
		if _, ok := i["ssh-host-key"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversSshHostKey2edl(i["ssh-host-key"], d, pre_append)
			tmp["ssh_host_key"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-SshHostKey")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key_validation"
		if _, ok := i["ssh-host-key-validation"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversSshHostKeyValidation2edl(i["ssh-host-key-validation"], d, pre_append)
			tmp["ssh_host_key_validation"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-SshHostKeyValidation")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversStatus2edl(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := i["translate-host"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversTranslateHost2edl(i["translate-host"], d, pre_append)
			tmp["translate_host"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-TranslateHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tunnel_encryption"
		if _, ok := i["tunnel-encryption"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversTunnelEncryption2edl(i["tunnel-encryption"], d, pre_append)
			tmp["tunnel_encryption"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-TunnelEncryption")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversType2edl(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "verify_cert"
		if _, ok := i["verify-cert"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversVerifyCert2edl(i["verify-cert"], d, pre_append)
			tmp["verify_cert"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-VerifyCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversWeight2edl(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAccessProxyApiGatewayRealserversAddrType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyApiGatewayRealserversDomain2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversExternalAuth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversHealthCheck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversHealthCheckProto2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversHolddownInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversHttpHost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversMappedport2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyApiGatewayRealserversPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversSshClientCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyApiGatewayRealserversSshHostKey2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyApiGatewayRealserversSshHostKeyValidation2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversTranslateHost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversTunnelEncryption2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversVerifyCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySamlRedirect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySamlServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyApiGatewayService2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslAlgorithm2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslCipherSuites2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallAccessProxyApiGatewaySslCipherSuitesCipher2edl(i["cipher"], d, pre_append)
			tmp["cipher"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-SslCipherSuites-Cipher")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenFirewallAccessProxyApiGatewaySslCipherSuitesPriority2edl(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-SslCipherSuites-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {
			v := flattenFirewallAccessProxyApiGatewaySslCipherSuitesVersions2edl(i["versions"], d, pre_append)
			tmp["versions"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-SslCipherSuites-Versions")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAccessProxyApiGatewaySslCipherSuitesCipher2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslCipherSuitesPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslCipherSuitesVersions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyApiGatewaySslDhBits2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslMaxVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslMinVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslRenegotiation2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslVpnWebPortal2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyApiGatewayUrlMap2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayUrlMapType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayVirtualHost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectFirewallAccessProxyApiGateway(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("application", flattenFirewallAccessProxyApiGatewayApplication2edl(o["application"], d, "application")); err != nil {
		if vv, ok := fortiAPIPatch(o["application"], "FirewallAccessProxyApiGateway-Application"); ok {
			if err = d.Set("application", vv); err != nil {
				return fmt.Errorf("Error reading application: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if err = d.Set("h2_support", flattenFirewallAccessProxyApiGatewayH2Support2edl(o["h2-support"], d, "h2_support")); err != nil {
		if vv, ok := fortiAPIPatch(o["h2-support"], "FirewallAccessProxyApiGateway-H2Support"); ok {
			if err = d.Set("h2_support", vv); err != nil {
				return fmt.Errorf("Error reading h2_support: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading h2_support: %v", err)
		}
	}

	if err = d.Set("h3_support", flattenFirewallAccessProxyApiGatewayH3Support2edl(o["h3-support"], d, "h3_support")); err != nil {
		if vv, ok := fortiAPIPatch(o["h3-support"], "FirewallAccessProxyApiGateway-H3Support"); ok {
			if err = d.Set("h3_support", vv); err != nil {
				return fmt.Errorf("Error reading h3_support: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading h3_support: %v", err)
		}
	}

	if err = d.Set("http_cookie_age", flattenFirewallAccessProxyApiGatewayHttpCookieAge2edl(o["http-cookie-age"], d, "http_cookie_age")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-age"], "FirewallAccessProxyApiGateway-HttpCookieAge"); ok {
			if err = d.Set("http_cookie_age", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_age: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_age: %v", err)
		}
	}

	if err = d.Set("http_cookie_domain", flattenFirewallAccessProxyApiGatewayHttpCookieDomain2edl(o["http-cookie-domain"], d, "http_cookie_domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-domain"], "FirewallAccessProxyApiGateway-HttpCookieDomain"); ok {
			if err = d.Set("http_cookie_domain", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_domain: %v", err)
		}
	}

	if err = d.Set("http_cookie_domain_from_host", flattenFirewallAccessProxyApiGatewayHttpCookieDomainFromHost2edl(o["http-cookie-domain-from-host"], d, "http_cookie_domain_from_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-domain-from-host"], "FirewallAccessProxyApiGateway-HttpCookieDomainFromHost"); ok {
			if err = d.Set("http_cookie_domain_from_host", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_domain_from_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_domain_from_host: %v", err)
		}
	}

	if err = d.Set("http_cookie_generation", flattenFirewallAccessProxyApiGatewayHttpCookieGeneration2edl(o["http-cookie-generation"], d, "http_cookie_generation")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-generation"], "FirewallAccessProxyApiGateway-HttpCookieGeneration"); ok {
			if err = d.Set("http_cookie_generation", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_generation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_generation: %v", err)
		}
	}

	if err = d.Set("http_cookie_path", flattenFirewallAccessProxyApiGatewayHttpCookiePath2edl(o["http-cookie-path"], d, "http_cookie_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-path"], "FirewallAccessProxyApiGateway-HttpCookiePath"); ok {
			if err = d.Set("http_cookie_path", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_path: %v", err)
		}
	}

	if err = d.Set("http_cookie_share", flattenFirewallAccessProxyApiGatewayHttpCookieShare2edl(o["http-cookie-share"], d, "http_cookie_share")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-share"], "FirewallAccessProxyApiGateway-HttpCookieShare"); ok {
			if err = d.Set("http_cookie_share", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_share: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_share: %v", err)
		}
	}

	if err = d.Set("https_cookie_secure", flattenFirewallAccessProxyApiGatewayHttpsCookieSecure2edl(o["https-cookie-secure"], d, "https_cookie_secure")); err != nil {
		if vv, ok := fortiAPIPatch(o["https-cookie-secure"], "FirewallAccessProxyApiGateway-HttpsCookieSecure"); ok {
			if err = d.Set("https_cookie_secure", vv); err != nil {
				return fmt.Errorf("Error reading https_cookie_secure: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading https_cookie_secure: %v", err)
		}
	}

	if err = d.Set("fosid", flattenFirewallAccessProxyApiGatewayId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FirewallAccessProxyApiGateway-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ldb_method", flattenFirewallAccessProxyApiGatewayLdbMethod2edl(o["ldb-method"], d, "ldb_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldb-method"], "FirewallAccessProxyApiGateway-LdbMethod"); ok {
			if err = d.Set("ldb_method", vv); err != nil {
				return fmt.Errorf("Error reading ldb_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldb_method: %v", err)
		}
	}

	if err = d.Set("persistence", flattenFirewallAccessProxyApiGatewayPersistence2edl(o["persistence"], d, "persistence")); err != nil {
		if vv, ok := fortiAPIPatch(o["persistence"], "FirewallAccessProxyApiGateway-Persistence"); ok {
			if err = d.Set("persistence", vv); err != nil {
				return fmt.Errorf("Error reading persistence: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading persistence: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("quic", flattenFirewallAccessProxyApiGatewayQuic2edl(o["quic"], d, "quic")); err != nil {
			if vv, ok := fortiAPIPatch(o["quic"], "FirewallAccessProxyApiGateway-Quic"); ok {
				if err = d.Set("quic", vv); err != nil {
					return fmt.Errorf("Error reading quic: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading quic: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("quic"); ok {
			if err = d.Set("quic", flattenFirewallAccessProxyApiGatewayQuic2edl(o["quic"], d, "quic")); err != nil {
				if vv, ok := fortiAPIPatch(o["quic"], "FirewallAccessProxyApiGateway-Quic"); ok {
					if err = d.Set("quic", vv); err != nil {
						return fmt.Errorf("Error reading quic: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading quic: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("realservers", flattenFirewallAccessProxyApiGatewayRealservers2edl(o["realservers"], d, "realservers")); err != nil {
			if vv, ok := fortiAPIPatch(o["realservers"], "FirewallAccessProxyApiGateway-Realservers"); ok {
				if err = d.Set("realservers", vv); err != nil {
					return fmt.Errorf("Error reading realservers: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading realservers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("realservers"); ok {
			if err = d.Set("realservers", flattenFirewallAccessProxyApiGatewayRealservers2edl(o["realservers"], d, "realservers")); err != nil {
				if vv, ok := fortiAPIPatch(o["realservers"], "FirewallAccessProxyApiGateway-Realservers"); ok {
					if err = d.Set("realservers", vv); err != nil {
						return fmt.Errorf("Error reading realservers: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading realservers: %v", err)
				}
			}
		}
	}

	if err = d.Set("saml_redirect", flattenFirewallAccessProxyApiGatewaySamlRedirect2edl(o["saml-redirect"], d, "saml_redirect")); err != nil {
		if vv, ok := fortiAPIPatch(o["saml-redirect"], "FirewallAccessProxyApiGateway-SamlRedirect"); ok {
			if err = d.Set("saml_redirect", vv); err != nil {
				return fmt.Errorf("Error reading saml_redirect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading saml_redirect: %v", err)
		}
	}

	if err = d.Set("saml_server", flattenFirewallAccessProxyApiGatewaySamlServer2edl(o["saml-server"], d, "saml_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["saml-server"], "FirewallAccessProxyApiGateway-SamlServer"); ok {
			if err = d.Set("saml_server", vv); err != nil {
				return fmt.Errorf("Error reading saml_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading saml_server: %v", err)
		}
	}

	if err = d.Set("service", flattenFirewallAccessProxyApiGatewayService2edl(o["service"], d, "service")); err != nil {
		if vv, ok := fortiAPIPatch(o["service"], "FirewallAccessProxyApiGateway-Service"); ok {
			if err = d.Set("service", vv); err != nil {
				return fmt.Errorf("Error reading service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("ssl_algorithm", flattenFirewallAccessProxyApiGatewaySslAlgorithm2edl(o["ssl-algorithm"], d, "ssl_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-algorithm"], "FirewallAccessProxyApiGateway-SslAlgorithm"); ok {
			if err = d.Set("ssl_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading ssl_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_algorithm: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ssl_cipher_suites", flattenFirewallAccessProxyApiGatewaySslCipherSuites2edl(o["ssl-cipher-suites"], d, "ssl_cipher_suites")); err != nil {
			if vv, ok := fortiAPIPatch(o["ssl-cipher-suites"], "FirewallAccessProxyApiGateway-SslCipherSuites"); ok {
				if err = d.Set("ssl_cipher_suites", vv); err != nil {
					return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_cipher_suites"); ok {
			if err = d.Set("ssl_cipher_suites", flattenFirewallAccessProxyApiGatewaySslCipherSuites2edl(o["ssl-cipher-suites"], d, "ssl_cipher_suites")); err != nil {
				if vv, ok := fortiAPIPatch(o["ssl-cipher-suites"], "FirewallAccessProxyApiGateway-SslCipherSuites"); ok {
					if err = d.Set("ssl_cipher_suites", vv); err != nil {
						return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_dh_bits", flattenFirewallAccessProxyApiGatewaySslDhBits2edl(o["ssl-dh-bits"], d, "ssl_dh_bits")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-dh-bits"], "FirewallAccessProxyApiGateway-SslDhBits"); ok {
			if err = d.Set("ssl_dh_bits", vv); err != nil {
				return fmt.Errorf("Error reading ssl_dh_bits: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_dh_bits: %v", err)
		}
	}

	if err = d.Set("ssl_max_version", flattenFirewallAccessProxyApiGatewaySslMaxVersion2edl(o["ssl-max-version"], d, "ssl_max_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-max-version"], "FirewallAccessProxyApiGateway-SslMaxVersion"); ok {
			if err = d.Set("ssl_max_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_max_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_max_version: %v", err)
		}
	}

	if err = d.Set("ssl_min_version", flattenFirewallAccessProxyApiGatewaySslMinVersion2edl(o["ssl-min-version"], d, "ssl_min_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-min-version"], "FirewallAccessProxyApiGateway-SslMinVersion"); ok {
			if err = d.Set("ssl_min_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_min_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_min_version: %v", err)
		}
	}

	if err = d.Set("ssl_renegotiation", flattenFirewallAccessProxyApiGatewaySslRenegotiation2edl(o["ssl-renegotiation"], d, "ssl_renegotiation")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-renegotiation"], "FirewallAccessProxyApiGateway-SslRenegotiation"); ok {
			if err = d.Set("ssl_renegotiation", vv); err != nil {
				return fmt.Errorf("Error reading ssl_renegotiation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_renegotiation: %v", err)
		}
	}

	if err = d.Set("ssl_vpn_web_portal", flattenFirewallAccessProxyApiGatewaySslVpnWebPortal2edl(o["ssl-vpn-web-portal"], d, "ssl_vpn_web_portal")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-vpn-web-portal"], "FirewallAccessProxyApiGateway-SslVpnWebPortal"); ok {
			if err = d.Set("ssl_vpn_web_portal", vv); err != nil {
				return fmt.Errorf("Error reading ssl_vpn_web_portal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_vpn_web_portal: %v", err)
		}
	}

	if err = d.Set("url_map", flattenFirewallAccessProxyApiGatewayUrlMap2edl(o["url-map"], d, "url_map")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-map"], "FirewallAccessProxyApiGateway-UrlMap"); ok {
			if err = d.Set("url_map", vv); err != nil {
				return fmt.Errorf("Error reading url_map: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_map: %v", err)
		}
	}

	if err = d.Set("url_map_type", flattenFirewallAccessProxyApiGatewayUrlMapType2edl(o["url-map-type"], d, "url_map_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-map-type"], "FirewallAccessProxyApiGateway-UrlMapType"); ok {
			if err = d.Set("url_map_type", vv); err != nil {
				return fmt.Errorf("Error reading url_map_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_map_type: %v", err)
		}
	}

	if err = d.Set("virtual_host", flattenFirewallAccessProxyApiGatewayVirtualHost2edl(o["virtual-host"], d, "virtual_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["virtual-host"], "FirewallAccessProxyApiGateway-VirtualHost"); ok {
			if err = d.Set("virtual_host", vv); err != nil {
				return fmt.Errorf("Error reading virtual_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading virtual_host: %v", err)
		}
	}

	return nil
}

func flattenFirewallAccessProxyApiGatewayFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallAccessProxyApiGatewayApplication2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyApiGatewayH2Support2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayH3Support2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayHttpCookieAge2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayHttpCookieDomain2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayHttpCookieDomainFromHost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayHttpCookieGeneration2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayHttpCookiePath2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayHttpCookieShare2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayHttpsCookieSecure2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayLdbMethod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayPersistence2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayQuic2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ack-delay-exponent"], _ = expandFirewallAccessProxyApiGatewayQuicAckDelayExponent2edl(d, i["ack_delay_exponent"], pre_append)
	}
	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["active-connection-id-limit"], _ = expandFirewallAccessProxyApiGatewayQuicActiveConnectionIdLimit2edl(d, i["active_connection_id_limit"], pre_append)
	}
	pre_append = pre + ".0." + "active_migration"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["active-migration"], _ = expandFirewallAccessProxyApiGatewayQuicActiveMigration2edl(d, i["active_migration"], pre_append)
	}
	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["grease-quic-bit"], _ = expandFirewallAccessProxyApiGatewayQuicGreaseQuicBit2edl(d, i["grease_quic_bit"], pre_append)
	}
	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-ack-delay"], _ = expandFirewallAccessProxyApiGatewayQuicMaxAckDelay2edl(d, i["max_ack_delay"], pre_append)
	}
	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-datagram-frame-size"], _ = expandFirewallAccessProxyApiGatewayQuicMaxDatagramFrameSize2edl(d, i["max_datagram_frame_size"], pre_append)
	}
	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-idle-timeout"], _ = expandFirewallAccessProxyApiGatewayQuicMaxIdleTimeout2edl(d, i["max_idle_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-udp-payload-size"], _ = expandFirewallAccessProxyApiGatewayQuicMaxUdpPayloadSize2edl(d, i["max_udp_payload_size"], pre_append)
	}

	return result, nil
}

func expandFirewallAccessProxyApiGatewayQuicAckDelayExponent2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayQuicActiveConnectionIdLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayQuicActiveMigration2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayQuicGreaseQuicBit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayQuicMaxAckDelay2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayQuicMaxDatagramFrameSize2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayQuicMaxIdleTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayQuicMaxUdpPayloadSize2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealservers2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["addr-type"], _ = expandFirewallAccessProxyApiGatewayRealserversAddrType2edl(d, i["addr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["address"], _ = expandFirewallAccessProxyApiGatewayRealserversAddress2edl(d, i["address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["domain"], _ = expandFirewallAccessProxyApiGatewayRealserversDomain2edl(d, i["domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_auth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["external-auth"], _ = expandFirewallAccessProxyApiGatewayRealserversExternalAuth2edl(d, i["external_auth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check"], _ = expandFirewallAccessProxyApiGatewayRealserversHealthCheck2edl(d, i["health_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check-proto"], _ = expandFirewallAccessProxyApiGatewayRealserversHealthCheckProto2edl(d, i["health_check_proto"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["holddown-interval"], _ = expandFirewallAccessProxyApiGatewayRealserversHolddownInterval2edl(d, i["holddown_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-host"], _ = expandFirewallAccessProxyApiGatewayRealserversHttpHost2edl(d, i["http_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFirewallAccessProxyApiGatewayRealserversId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandFirewallAccessProxyApiGatewayRealserversIp2edl(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mappedport"], _ = expandFirewallAccessProxyApiGatewayRealserversMappedport2edl(d, i["mappedport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandFirewallAccessProxyApiGatewayRealserversPort2edl(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_client_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssh-client-cert"], _ = expandFirewallAccessProxyApiGatewayRealserversSshClientCert2edl(d, i["ssh_client_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssh-host-key"], _ = expandFirewallAccessProxyApiGatewayRealserversSshHostKey2edl(d, i["ssh_host_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key_validation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssh-host-key-validation"], _ = expandFirewallAccessProxyApiGatewayRealserversSshHostKeyValidation2edl(d, i["ssh_host_key_validation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandFirewallAccessProxyApiGatewayRealserversStatus2edl(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["translate-host"], _ = expandFirewallAccessProxyApiGatewayRealserversTranslateHost2edl(d, i["translate_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tunnel_encryption"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tunnel-encryption"], _ = expandFirewallAccessProxyApiGatewayRealserversTunnelEncryption2edl(d, i["tunnel_encryption"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandFirewallAccessProxyApiGatewayRealserversType2edl(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "verify_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["verify-cert"], _ = expandFirewallAccessProxyApiGatewayRealserversVerifyCert2edl(d, i["verify_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandFirewallAccessProxyApiGatewayRealserversWeight2edl(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxyApiGatewayRealserversAddrType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyApiGatewayRealserversDomain2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversExternalAuth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversHealthCheck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversHealthCheckProto2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversHolddownInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversHttpHost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversMappedport2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyApiGatewayRealserversPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversSshClientCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyApiGatewayRealserversSshHostKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyApiGatewayRealserversSshHostKeyValidation2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversTranslateHost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversTunnelEncryption2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversVerifyCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySamlRedirect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySamlServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyApiGatewayService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslAlgorithm2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslCipherSuites2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["cipher"], _ = expandFirewallAccessProxyApiGatewaySslCipherSuitesCipher2edl(d, i["cipher"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandFirewallAccessProxyApiGatewaySslCipherSuitesPriority2edl(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["versions"], _ = expandFirewallAccessProxyApiGatewaySslCipherSuitesVersions2edl(d, i["versions"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxyApiGatewaySslCipherSuitesCipher2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslCipherSuitesPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslCipherSuitesVersions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyApiGatewaySslDhBits2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslMaxVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslMinVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslRenegotiation2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslVpnWebPortal2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyApiGatewayUrlMap2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayUrlMapType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayVirtualHost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectFirewallAccessProxyApiGateway(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("application"); ok || d.HasChange("application") {
		t, err := expandFirewallAccessProxyApiGatewayApplication2edl(d, v, "application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("h2_support"); ok || d.HasChange("h2_support") {
		t, err := expandFirewallAccessProxyApiGatewayH2Support2edl(d, v, "h2_support")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h2-support"] = t
		}
	}

	if v, ok := d.GetOk("h3_support"); ok || d.HasChange("h3_support") {
		t, err := expandFirewallAccessProxyApiGatewayH3Support2edl(d, v, "h3_support")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-support"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_age"); ok || d.HasChange("http_cookie_age") {
		t, err := expandFirewallAccessProxyApiGatewayHttpCookieAge2edl(d, v, "http_cookie_age")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-age"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_domain"); ok || d.HasChange("http_cookie_domain") {
		t, err := expandFirewallAccessProxyApiGatewayHttpCookieDomain2edl(d, v, "http_cookie_domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-domain"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_domain_from_host"); ok || d.HasChange("http_cookie_domain_from_host") {
		t, err := expandFirewallAccessProxyApiGatewayHttpCookieDomainFromHost2edl(d, v, "http_cookie_domain_from_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-domain-from-host"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_generation"); ok || d.HasChange("http_cookie_generation") {
		t, err := expandFirewallAccessProxyApiGatewayHttpCookieGeneration2edl(d, v, "http_cookie_generation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-generation"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_path"); ok || d.HasChange("http_cookie_path") {
		t, err := expandFirewallAccessProxyApiGatewayHttpCookiePath2edl(d, v, "http_cookie_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-path"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_share"); ok || d.HasChange("http_cookie_share") {
		t, err := expandFirewallAccessProxyApiGatewayHttpCookieShare2edl(d, v, "http_cookie_share")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-share"] = t
		}
	}

	if v, ok := d.GetOk("https_cookie_secure"); ok || d.HasChange("https_cookie_secure") {
		t, err := expandFirewallAccessProxyApiGatewayHttpsCookieSecure2edl(d, v, "https_cookie_secure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-cookie-secure"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandFirewallAccessProxyApiGatewayId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ldb_method"); ok || d.HasChange("ldb_method") {
		t, err := expandFirewallAccessProxyApiGatewayLdbMethod2edl(d, v, "ldb_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldb-method"] = t
		}
	}

	if v, ok := d.GetOk("persistence"); ok || d.HasChange("persistence") {
		t, err := expandFirewallAccessProxyApiGatewayPersistence2edl(d, v, "persistence")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["persistence"] = t
		}
	}

	if v, ok := d.GetOk("quic"); ok || d.HasChange("quic") {
		t, err := expandFirewallAccessProxyApiGatewayQuic2edl(d, v, "quic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quic"] = t
		}
	}

	if v, ok := d.GetOk("realservers"); ok || d.HasChange("realservers") {
		t, err := expandFirewallAccessProxyApiGatewayRealservers2edl(d, v, "realservers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["realservers"] = t
		}
	}

	if v, ok := d.GetOk("saml_redirect"); ok || d.HasChange("saml_redirect") {
		t, err := expandFirewallAccessProxyApiGatewaySamlRedirect2edl(d, v, "saml_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["saml-redirect"] = t
		}
	}

	if v, ok := d.GetOk("saml_server"); ok || d.HasChange("saml_server") {
		t, err := expandFirewallAccessProxyApiGatewaySamlServer2edl(d, v, "saml_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["saml-server"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandFirewallAccessProxyApiGatewayService2edl(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("ssl_algorithm"); ok || d.HasChange("ssl_algorithm") {
		t, err := expandFirewallAccessProxyApiGatewaySslAlgorithm2edl(d, v, "ssl_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_cipher_suites"); ok || d.HasChange("ssl_cipher_suites") {
		t, err := expandFirewallAccessProxyApiGatewaySslCipherSuites2edl(d, v, "ssl_cipher_suites")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-cipher-suites"] = t
		}
	}

	if v, ok := d.GetOk("ssl_dh_bits"); ok || d.HasChange("ssl_dh_bits") {
		t, err := expandFirewallAccessProxyApiGatewaySslDhBits2edl(d, v, "ssl_dh_bits")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-dh-bits"] = t
		}
	}

	if v, ok := d.GetOk("ssl_max_version"); ok || d.HasChange("ssl_max_version") {
		t, err := expandFirewallAccessProxyApiGatewaySslMaxVersion2edl(d, v, "ssl_max_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-max-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_version"); ok || d.HasChange("ssl_min_version") {
		t, err := expandFirewallAccessProxyApiGatewaySslMinVersion2edl(d, v, "ssl_min_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_renegotiation"); ok || d.HasChange("ssl_renegotiation") {
		t, err := expandFirewallAccessProxyApiGatewaySslRenegotiation2edl(d, v, "ssl_renegotiation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-renegotiation"] = t
		}
	}

	if v, ok := d.GetOk("ssl_vpn_web_portal"); ok || d.HasChange("ssl_vpn_web_portal") {
		t, err := expandFirewallAccessProxyApiGatewaySslVpnWebPortal2edl(d, v, "ssl_vpn_web_portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-vpn-web-portal"] = t
		}
	}

	if v, ok := d.GetOk("url_map"); ok || d.HasChange("url_map") {
		t, err := expandFirewallAccessProxyApiGatewayUrlMap2edl(d, v, "url_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-map"] = t
		}
	}

	if v, ok := d.GetOk("url_map_type"); ok || d.HasChange("url_map_type") {
		t, err := expandFirewallAccessProxyApiGatewayUrlMapType2edl(d, v, "url_map_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-map-type"] = t
		}
	}

	if v, ok := d.GetOk("virtual_host"); ok || d.HasChange("virtual_host") {
		t, err := expandFirewallAccessProxyApiGatewayVirtualHost2edl(d, v, "virtual_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-host"] = t
		}
	}

	return &obj, nil
}
