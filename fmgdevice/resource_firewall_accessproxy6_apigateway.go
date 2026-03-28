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

func resourceFirewallAccessProxy6ApiGateway() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallAccessProxy6ApiGatewayCreate,
		Read:   resourceFirewallAccessProxy6ApiGatewayRead,
		Update: resourceFirewallAccessProxy6ApiGatewayUpdate,
		Delete: resourceFirewallAccessProxy6ApiGatewayDelete,

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
			"access_proxy6": &schema.Schema{
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
				Computed: true,
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
							Computed: true,
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
							Computed: true,
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
							Computed: true,
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
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
				Computed: true,
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

func resourceFirewallAccessProxy6ApiGatewayCreate(d *schema.ResourceData, m interface{}) error {
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
	access_proxy6 := d.Get("access_proxy6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_proxy6"] = access_proxy6

	obj, err := getObjectFirewallAccessProxy6ApiGateway(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallAccessProxy6ApiGateway resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallAccessProxy6ApiGateway(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallAccessProxy6ApiGateway(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallAccessProxy6ApiGateway resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateFirewallAccessProxy6ApiGateway(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallAccessProxy6ApiGateway resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceFirewallAccessProxy6ApiGatewayRead(d, m)
			} else {
				return fmt.Errorf("Error creating FirewallAccessProxy6ApiGateway resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallAccessProxy6ApiGatewayRead(d, m)
}

func resourceFirewallAccessProxy6ApiGatewayUpdate(d *schema.ResourceData, m interface{}) error {
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
	access_proxy6 := d.Get("access_proxy6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_proxy6"] = access_proxy6

	obj, err := getObjectFirewallAccessProxy6ApiGateway(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAccessProxy6ApiGateway resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallAccessProxy6ApiGateway(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAccessProxy6ApiGateway resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallAccessProxy6ApiGatewayRead(d, m)
}

func resourceFirewallAccessProxy6ApiGatewayDelete(d *schema.ResourceData, m interface{}) error {
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
	access_proxy6 := d.Get("access_proxy6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_proxy6"] = access_proxy6

	wsParams["adom"] = adomv

	err = c.DeleteFirewallAccessProxy6ApiGateway(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallAccessProxy6ApiGateway resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAccessProxy6ApiGatewayRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	access_proxy6 := d.Get("access_proxy6").(string)
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
	if access_proxy6 == "" {
		access_proxy6 = importOptionChecking(m.(*FortiClient).Cfg, "access_proxy6")
		if access_proxy6 == "" {
			return fmt.Errorf("Parameter access_proxy6 is missing")
		}
		if err = d.Set("access_proxy6", access_proxy6); err != nil {
			return fmt.Errorf("Error set params access_proxy6: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_proxy6"] = access_proxy6

	o, err := c.ReadFirewallAccessProxy6ApiGateway(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallAccessProxy6ApiGateway resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallAccessProxy6ApiGateway(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAccessProxy6ApiGateway resource from API: %v", err)
	}
	return nil
}

func flattenFirewallAccessProxy6ApiGatewayApplication2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGatewayH2Support2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayH3Support2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayHttpCookieAge2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayHttpCookieDomain2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayHttpCookieDomainFromHost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayHttpCookieGeneration2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayHttpCookiePath2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayHttpCookieShare2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayHttpsCookieSecure2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayLdbMethod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayPersistence2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayQuic2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := i["ack-delay-exponent"]; ok {
		result["ack_delay_exponent"] = flattenFirewallAccessProxy6ApiGatewayQuicAckDelayExponent2edl(i["ack-delay-exponent"], d, pre_append)
	}

	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := i["active-connection-id-limit"]; ok {
		result["active_connection_id_limit"] = flattenFirewallAccessProxy6ApiGatewayQuicActiveConnectionIdLimit2edl(i["active-connection-id-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "active_migration"
	if _, ok := i["active-migration"]; ok {
		result["active_migration"] = flattenFirewallAccessProxy6ApiGatewayQuicActiveMigration2edl(i["active-migration"], d, pre_append)
	}

	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := i["grease-quic-bit"]; ok {
		result["grease_quic_bit"] = flattenFirewallAccessProxy6ApiGatewayQuicGreaseQuicBit2edl(i["grease-quic-bit"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := i["max-ack-delay"]; ok {
		result["max_ack_delay"] = flattenFirewallAccessProxy6ApiGatewayQuicMaxAckDelay2edl(i["max-ack-delay"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := i["max-datagram-frame-size"]; ok {
		result["max_datagram_frame_size"] = flattenFirewallAccessProxy6ApiGatewayQuicMaxDatagramFrameSize2edl(i["max-datagram-frame-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := i["max-idle-timeout"]; ok {
		result["max_idle_timeout"] = flattenFirewallAccessProxy6ApiGatewayQuicMaxIdleTimeout2edl(i["max-idle-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := i["max-udp-payload-size"]; ok {
		result["max_udp_payload_size"] = flattenFirewallAccessProxy6ApiGatewayQuicMaxUdpPayloadSize2edl(i["max-udp-payload-size"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallAccessProxy6ApiGatewayQuicAckDelayExponent2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayQuicActiveConnectionIdLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayQuicActiveMigration2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayQuicGreaseQuicBit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayQuicMaxAckDelay2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayQuicMaxDatagramFrameSize2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayQuicMaxIdleTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayQuicMaxUdpPayloadSize2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealservers2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallAccessProxy6ApiGatewayRealserversAddrType2edl(i["addr-type"], d, pre_append)
			tmp["addr_type"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-AddrType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversAddress2edl(i["address"], d, pre_append)
			tmp["address"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := i["domain"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversDomain2edl(i["domain"], d, pre_append)
			tmp["domain"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-Domain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_auth"
		if _, ok := i["external-auth"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversExternalAuth2edl(i["external-auth"], d, pre_append)
			tmp["external_auth"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-ExternalAuth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := i["health-check"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversHealthCheck2edl(i["health-check"], d, pre_append)
			tmp["health_check"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-HealthCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := i["health-check-proto"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversHealthCheckProto2edl(i["health-check-proto"], d, pre_append)
			tmp["health_check_proto"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-HealthCheckProto")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := i["holddown-interval"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversHolddownInterval2edl(i["holddown-interval"], d, pre_append)
			tmp["holddown_interval"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-HolddownInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := i["http-host"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversHttpHost2edl(i["http-host"], d, pre_append)
			tmp["http_host"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-HttpHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversIp2edl(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := i["mappedport"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversMappedport2edl(i["mappedport"], d, pre_append)
			tmp["mappedport"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-Mappedport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversPort2edl(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_client_cert"
		if _, ok := i["ssh-client-cert"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversSshClientCert2edl(i["ssh-client-cert"], d, pre_append)
			tmp["ssh_client_cert"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-SshClientCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key"
		if _, ok := i["ssh-host-key"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversSshHostKey2edl(i["ssh-host-key"], d, pre_append)
			tmp["ssh_host_key"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-SshHostKey")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key_validation"
		if _, ok := i["ssh-host-key-validation"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversSshHostKeyValidation2edl(i["ssh-host-key-validation"], d, pre_append)
			tmp["ssh_host_key_validation"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-SshHostKeyValidation")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversStatus2edl(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := i["translate-host"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversTranslateHost2edl(i["translate-host"], d, pre_append)
			tmp["translate_host"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-TranslateHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tunnel_encryption"
		if _, ok := i["tunnel-encryption"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversTunnelEncryption2edl(i["tunnel-encryption"], d, pre_append)
			tmp["tunnel_encryption"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-TunnelEncryption")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversType2edl(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "verify_cert"
		if _, ok := i["verify-cert"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversVerifyCert2edl(i["verify-cert"], d, pre_append)
			tmp["verify_cert"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-VerifyCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversWeight2edl(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAccessProxy6ApiGatewayRealserversAddrType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGatewayRealserversDomain2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversExternalAuth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversHealthCheck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversHealthCheckProto2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversHolddownInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversHttpHost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversMappedport2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGatewayRealserversPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversSshClientCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGatewayRealserversSshHostKey2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGatewayRealserversSshHostKeyValidation2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversTranslateHost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversTunnelEncryption2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversVerifyCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewaySamlRedirect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewaySamlServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGatewayService2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewaySslAlgorithm2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewaySslCipherSuites2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallAccessProxy6ApiGatewaySslCipherSuitesCipher2edl(i["cipher"], d, pre_append)
			tmp["cipher"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-SslCipherSuites-Cipher")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewaySslCipherSuitesPriority2edl(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-SslCipherSuites-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewaySslCipherSuitesVersions2edl(i["versions"], d, pre_append)
			tmp["versions"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-SslCipherSuites-Versions")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAccessProxy6ApiGatewaySslCipherSuitesCipher2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewaySslCipherSuitesPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewaySslCipherSuitesVersions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGatewaySslDhBits2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewaySslMaxVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewaySslMinVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewaySslRenegotiation2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewaySslVpnWebPortal2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGatewayUrlMap2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayUrlMapType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayVirtualHost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectFirewallAccessProxy6ApiGateway(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("application", flattenFirewallAccessProxy6ApiGatewayApplication2edl(o["application"], d, "application")); err != nil {
		if vv, ok := fortiAPIPatch(o["application"], "FirewallAccessProxy6ApiGateway-Application"); ok {
			if err = d.Set("application", vv); err != nil {
				return fmt.Errorf("Error reading application: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if err = d.Set("h2_support", flattenFirewallAccessProxy6ApiGatewayH2Support2edl(o["h2-support"], d, "h2_support")); err != nil {
		if vv, ok := fortiAPIPatch(o["h2-support"], "FirewallAccessProxy6ApiGateway-H2Support"); ok {
			if err = d.Set("h2_support", vv); err != nil {
				return fmt.Errorf("Error reading h2_support: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading h2_support: %v", err)
		}
	}

	if err = d.Set("h3_support", flattenFirewallAccessProxy6ApiGatewayH3Support2edl(o["h3-support"], d, "h3_support")); err != nil {
		if vv, ok := fortiAPIPatch(o["h3-support"], "FirewallAccessProxy6ApiGateway-H3Support"); ok {
			if err = d.Set("h3_support", vv); err != nil {
				return fmt.Errorf("Error reading h3_support: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading h3_support: %v", err)
		}
	}

	if err = d.Set("http_cookie_age", flattenFirewallAccessProxy6ApiGatewayHttpCookieAge2edl(o["http-cookie-age"], d, "http_cookie_age")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-age"], "FirewallAccessProxy6ApiGateway-HttpCookieAge"); ok {
			if err = d.Set("http_cookie_age", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_age: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_age: %v", err)
		}
	}

	if err = d.Set("http_cookie_domain", flattenFirewallAccessProxy6ApiGatewayHttpCookieDomain2edl(o["http-cookie-domain"], d, "http_cookie_domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-domain"], "FirewallAccessProxy6ApiGateway-HttpCookieDomain"); ok {
			if err = d.Set("http_cookie_domain", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_domain: %v", err)
		}
	}

	if err = d.Set("http_cookie_domain_from_host", flattenFirewallAccessProxy6ApiGatewayHttpCookieDomainFromHost2edl(o["http-cookie-domain-from-host"], d, "http_cookie_domain_from_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-domain-from-host"], "FirewallAccessProxy6ApiGateway-HttpCookieDomainFromHost"); ok {
			if err = d.Set("http_cookie_domain_from_host", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_domain_from_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_domain_from_host: %v", err)
		}
	}

	if err = d.Set("http_cookie_generation", flattenFirewallAccessProxy6ApiGatewayHttpCookieGeneration2edl(o["http-cookie-generation"], d, "http_cookie_generation")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-generation"], "FirewallAccessProxy6ApiGateway-HttpCookieGeneration"); ok {
			if err = d.Set("http_cookie_generation", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_generation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_generation: %v", err)
		}
	}

	if err = d.Set("http_cookie_path", flattenFirewallAccessProxy6ApiGatewayHttpCookiePath2edl(o["http-cookie-path"], d, "http_cookie_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-path"], "FirewallAccessProxy6ApiGateway-HttpCookiePath"); ok {
			if err = d.Set("http_cookie_path", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_path: %v", err)
		}
	}

	if err = d.Set("http_cookie_share", flattenFirewallAccessProxy6ApiGatewayHttpCookieShare2edl(o["http-cookie-share"], d, "http_cookie_share")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-share"], "FirewallAccessProxy6ApiGateway-HttpCookieShare"); ok {
			if err = d.Set("http_cookie_share", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_share: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_share: %v", err)
		}
	}

	if err = d.Set("https_cookie_secure", flattenFirewallAccessProxy6ApiGatewayHttpsCookieSecure2edl(o["https-cookie-secure"], d, "https_cookie_secure")); err != nil {
		if vv, ok := fortiAPIPatch(o["https-cookie-secure"], "FirewallAccessProxy6ApiGateway-HttpsCookieSecure"); ok {
			if err = d.Set("https_cookie_secure", vv); err != nil {
				return fmt.Errorf("Error reading https_cookie_secure: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading https_cookie_secure: %v", err)
		}
	}

	if err = d.Set("fosid", flattenFirewallAccessProxy6ApiGatewayId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FirewallAccessProxy6ApiGateway-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ldb_method", flattenFirewallAccessProxy6ApiGatewayLdbMethod2edl(o["ldb-method"], d, "ldb_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldb-method"], "FirewallAccessProxy6ApiGateway-LdbMethod"); ok {
			if err = d.Set("ldb_method", vv); err != nil {
				return fmt.Errorf("Error reading ldb_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldb_method: %v", err)
		}
	}

	if err = d.Set("persistence", flattenFirewallAccessProxy6ApiGatewayPersistence2edl(o["persistence"], d, "persistence")); err != nil {
		if vv, ok := fortiAPIPatch(o["persistence"], "FirewallAccessProxy6ApiGateway-Persistence"); ok {
			if err = d.Set("persistence", vv); err != nil {
				return fmt.Errorf("Error reading persistence: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading persistence: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("quic", flattenFirewallAccessProxy6ApiGatewayQuic2edl(o["quic"], d, "quic")); err != nil {
			if vv, ok := fortiAPIPatch(o["quic"], "FirewallAccessProxy6ApiGateway-Quic"); ok {
				if err = d.Set("quic", vv); err != nil {
					return fmt.Errorf("Error reading quic: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading quic: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("quic"); ok {
			if err = d.Set("quic", flattenFirewallAccessProxy6ApiGatewayQuic2edl(o["quic"], d, "quic")); err != nil {
				if vv, ok := fortiAPIPatch(o["quic"], "FirewallAccessProxy6ApiGateway-Quic"); ok {
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
		if err = d.Set("realservers", flattenFirewallAccessProxy6ApiGatewayRealservers2edl(o["realservers"], d, "realservers")); err != nil {
			if vv, ok := fortiAPIPatch(o["realservers"], "FirewallAccessProxy6ApiGateway-Realservers"); ok {
				if err = d.Set("realservers", vv); err != nil {
					return fmt.Errorf("Error reading realservers: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading realservers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("realservers"); ok {
			if err = d.Set("realservers", flattenFirewallAccessProxy6ApiGatewayRealservers2edl(o["realservers"], d, "realservers")); err != nil {
				if vv, ok := fortiAPIPatch(o["realservers"], "FirewallAccessProxy6ApiGateway-Realservers"); ok {
					if err = d.Set("realservers", vv); err != nil {
						return fmt.Errorf("Error reading realservers: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading realservers: %v", err)
				}
			}
		}
	}

	if err = d.Set("saml_redirect", flattenFirewallAccessProxy6ApiGatewaySamlRedirect2edl(o["saml-redirect"], d, "saml_redirect")); err != nil {
		if vv, ok := fortiAPIPatch(o["saml-redirect"], "FirewallAccessProxy6ApiGateway-SamlRedirect"); ok {
			if err = d.Set("saml_redirect", vv); err != nil {
				return fmt.Errorf("Error reading saml_redirect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading saml_redirect: %v", err)
		}
	}

	if err = d.Set("saml_server", flattenFirewallAccessProxy6ApiGatewaySamlServer2edl(o["saml-server"], d, "saml_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["saml-server"], "FirewallAccessProxy6ApiGateway-SamlServer"); ok {
			if err = d.Set("saml_server", vv); err != nil {
				return fmt.Errorf("Error reading saml_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading saml_server: %v", err)
		}
	}

	if err = d.Set("service", flattenFirewallAccessProxy6ApiGatewayService2edl(o["service"], d, "service")); err != nil {
		if vv, ok := fortiAPIPatch(o["service"], "FirewallAccessProxy6ApiGateway-Service"); ok {
			if err = d.Set("service", vv); err != nil {
				return fmt.Errorf("Error reading service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("ssl_algorithm", flattenFirewallAccessProxy6ApiGatewaySslAlgorithm2edl(o["ssl-algorithm"], d, "ssl_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-algorithm"], "FirewallAccessProxy6ApiGateway-SslAlgorithm"); ok {
			if err = d.Set("ssl_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading ssl_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_algorithm: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ssl_cipher_suites", flattenFirewallAccessProxy6ApiGatewaySslCipherSuites2edl(o["ssl-cipher-suites"], d, "ssl_cipher_suites")); err != nil {
			if vv, ok := fortiAPIPatch(o["ssl-cipher-suites"], "FirewallAccessProxy6ApiGateway-SslCipherSuites"); ok {
				if err = d.Set("ssl_cipher_suites", vv); err != nil {
					return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_cipher_suites"); ok {
			if err = d.Set("ssl_cipher_suites", flattenFirewallAccessProxy6ApiGatewaySslCipherSuites2edl(o["ssl-cipher-suites"], d, "ssl_cipher_suites")); err != nil {
				if vv, ok := fortiAPIPatch(o["ssl-cipher-suites"], "FirewallAccessProxy6ApiGateway-SslCipherSuites"); ok {
					if err = d.Set("ssl_cipher_suites", vv); err != nil {
						return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_dh_bits", flattenFirewallAccessProxy6ApiGatewaySslDhBits2edl(o["ssl-dh-bits"], d, "ssl_dh_bits")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-dh-bits"], "FirewallAccessProxy6ApiGateway-SslDhBits"); ok {
			if err = d.Set("ssl_dh_bits", vv); err != nil {
				return fmt.Errorf("Error reading ssl_dh_bits: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_dh_bits: %v", err)
		}
	}

	if err = d.Set("ssl_max_version", flattenFirewallAccessProxy6ApiGatewaySslMaxVersion2edl(o["ssl-max-version"], d, "ssl_max_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-max-version"], "FirewallAccessProxy6ApiGateway-SslMaxVersion"); ok {
			if err = d.Set("ssl_max_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_max_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_max_version: %v", err)
		}
	}

	if err = d.Set("ssl_min_version", flattenFirewallAccessProxy6ApiGatewaySslMinVersion2edl(o["ssl-min-version"], d, "ssl_min_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-min-version"], "FirewallAccessProxy6ApiGateway-SslMinVersion"); ok {
			if err = d.Set("ssl_min_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_min_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_min_version: %v", err)
		}
	}

	if err = d.Set("ssl_renegotiation", flattenFirewallAccessProxy6ApiGatewaySslRenegotiation2edl(o["ssl-renegotiation"], d, "ssl_renegotiation")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-renegotiation"], "FirewallAccessProxy6ApiGateway-SslRenegotiation"); ok {
			if err = d.Set("ssl_renegotiation", vv); err != nil {
				return fmt.Errorf("Error reading ssl_renegotiation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_renegotiation: %v", err)
		}
	}

	if err = d.Set("ssl_vpn_web_portal", flattenFirewallAccessProxy6ApiGatewaySslVpnWebPortal2edl(o["ssl-vpn-web-portal"], d, "ssl_vpn_web_portal")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-vpn-web-portal"], "FirewallAccessProxy6ApiGateway-SslVpnWebPortal"); ok {
			if err = d.Set("ssl_vpn_web_portal", vv); err != nil {
				return fmt.Errorf("Error reading ssl_vpn_web_portal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_vpn_web_portal: %v", err)
		}
	}

	if err = d.Set("url_map", flattenFirewallAccessProxy6ApiGatewayUrlMap2edl(o["url-map"], d, "url_map")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-map"], "FirewallAccessProxy6ApiGateway-UrlMap"); ok {
			if err = d.Set("url_map", vv); err != nil {
				return fmt.Errorf("Error reading url_map: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_map: %v", err)
		}
	}

	if err = d.Set("url_map_type", flattenFirewallAccessProxy6ApiGatewayUrlMapType2edl(o["url-map-type"], d, "url_map_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-map-type"], "FirewallAccessProxy6ApiGateway-UrlMapType"); ok {
			if err = d.Set("url_map_type", vv); err != nil {
				return fmt.Errorf("Error reading url_map_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_map_type: %v", err)
		}
	}

	if err = d.Set("virtual_host", flattenFirewallAccessProxy6ApiGatewayVirtualHost2edl(o["virtual-host"], d, "virtual_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["virtual-host"], "FirewallAccessProxy6ApiGateway-VirtualHost"); ok {
			if err = d.Set("virtual_host", vv); err != nil {
				return fmt.Errorf("Error reading virtual_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading virtual_host: %v", err)
		}
	}

	return nil
}

func flattenFirewallAccessProxy6ApiGatewayFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallAccessProxy6ApiGatewayApplication2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGatewayH2Support2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayH3Support2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayHttpCookieAge2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayHttpCookieDomain2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayHttpCookieDomainFromHost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayHttpCookieGeneration2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayHttpCookiePath2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayHttpCookieShare2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayHttpsCookieSecure2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayLdbMethod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayPersistence2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayQuic2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ack-delay-exponent"], _ = expandFirewallAccessProxy6ApiGatewayQuicAckDelayExponent2edl(d, i["ack_delay_exponent"], pre_append)
	}
	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["active-connection-id-limit"], _ = expandFirewallAccessProxy6ApiGatewayQuicActiveConnectionIdLimit2edl(d, i["active_connection_id_limit"], pre_append)
	}
	pre_append = pre + ".0." + "active_migration"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["active-migration"], _ = expandFirewallAccessProxy6ApiGatewayQuicActiveMigration2edl(d, i["active_migration"], pre_append)
	}
	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["grease-quic-bit"], _ = expandFirewallAccessProxy6ApiGatewayQuicGreaseQuicBit2edl(d, i["grease_quic_bit"], pre_append)
	}
	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-ack-delay"], _ = expandFirewallAccessProxy6ApiGatewayQuicMaxAckDelay2edl(d, i["max_ack_delay"], pre_append)
	}
	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-datagram-frame-size"], _ = expandFirewallAccessProxy6ApiGatewayQuicMaxDatagramFrameSize2edl(d, i["max_datagram_frame_size"], pre_append)
	}
	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-idle-timeout"], _ = expandFirewallAccessProxy6ApiGatewayQuicMaxIdleTimeout2edl(d, i["max_idle_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-udp-payload-size"], _ = expandFirewallAccessProxy6ApiGatewayQuicMaxUdpPayloadSize2edl(d, i["max_udp_payload_size"], pre_append)
	}

	return result, nil
}

func expandFirewallAccessProxy6ApiGatewayQuicAckDelayExponent2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayQuicActiveConnectionIdLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayQuicActiveMigration2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayQuicGreaseQuicBit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayQuicMaxAckDelay2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayQuicMaxDatagramFrameSize2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayQuicMaxIdleTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayQuicMaxUdpPayloadSize2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealservers2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["addr-type"], _ = expandFirewallAccessProxy6ApiGatewayRealserversAddrType2edl(d, i["addr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["address"], _ = expandFirewallAccessProxy6ApiGatewayRealserversAddress2edl(d, i["address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["domain"], _ = expandFirewallAccessProxy6ApiGatewayRealserversDomain2edl(d, i["domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_auth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["external-auth"], _ = expandFirewallAccessProxy6ApiGatewayRealserversExternalAuth2edl(d, i["external_auth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check"], _ = expandFirewallAccessProxy6ApiGatewayRealserversHealthCheck2edl(d, i["health_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check-proto"], _ = expandFirewallAccessProxy6ApiGatewayRealserversHealthCheckProto2edl(d, i["health_check_proto"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["holddown-interval"], _ = expandFirewallAccessProxy6ApiGatewayRealserversHolddownInterval2edl(d, i["holddown_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-host"], _ = expandFirewallAccessProxy6ApiGatewayRealserversHttpHost2edl(d, i["http_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFirewallAccessProxy6ApiGatewayRealserversId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandFirewallAccessProxy6ApiGatewayRealserversIp2edl(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mappedport"], _ = expandFirewallAccessProxy6ApiGatewayRealserversMappedport2edl(d, i["mappedport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandFirewallAccessProxy6ApiGatewayRealserversPort2edl(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_client_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssh-client-cert"], _ = expandFirewallAccessProxy6ApiGatewayRealserversSshClientCert2edl(d, i["ssh_client_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssh-host-key"], _ = expandFirewallAccessProxy6ApiGatewayRealserversSshHostKey2edl(d, i["ssh_host_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key_validation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssh-host-key-validation"], _ = expandFirewallAccessProxy6ApiGatewayRealserversSshHostKeyValidation2edl(d, i["ssh_host_key_validation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandFirewallAccessProxy6ApiGatewayRealserversStatus2edl(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["translate-host"], _ = expandFirewallAccessProxy6ApiGatewayRealserversTranslateHost2edl(d, i["translate_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tunnel_encryption"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tunnel-encryption"], _ = expandFirewallAccessProxy6ApiGatewayRealserversTunnelEncryption2edl(d, i["tunnel_encryption"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandFirewallAccessProxy6ApiGatewayRealserversType2edl(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "verify_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["verify-cert"], _ = expandFirewallAccessProxy6ApiGatewayRealserversVerifyCert2edl(d, i["verify_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandFirewallAccessProxy6ApiGatewayRealserversWeight2edl(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversAddrType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversDomain2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversExternalAuth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversHealthCheck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversHealthCheckProto2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversHolddownInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversHttpHost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversMappedport2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversSshClientCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversSshHostKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversSshHostKeyValidation2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversTranslateHost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversTunnelEncryption2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversVerifyCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySamlRedirect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySamlServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGatewayService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySslAlgorithm2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySslCipherSuites2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["cipher"], _ = expandFirewallAccessProxy6ApiGatewaySslCipherSuitesCipher2edl(d, i["cipher"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandFirewallAccessProxy6ApiGatewaySslCipherSuitesPriority2edl(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["versions"], _ = expandFirewallAccessProxy6ApiGatewaySslCipherSuitesVersions2edl(d, i["versions"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxy6ApiGatewaySslCipherSuitesCipher2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySslCipherSuitesPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySslCipherSuitesVersions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGatewaySslDhBits2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySslMaxVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySslMinVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySslRenegotiation2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySslVpnWebPortal2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGatewayUrlMap2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayUrlMapType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayVirtualHost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectFirewallAccessProxy6ApiGateway(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("application"); ok || d.HasChange("application") {
		t, err := expandFirewallAccessProxy6ApiGatewayApplication2edl(d, v, "application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("h2_support"); ok || d.HasChange("h2_support") {
		t, err := expandFirewallAccessProxy6ApiGatewayH2Support2edl(d, v, "h2_support")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h2-support"] = t
		}
	}

	if v, ok := d.GetOk("h3_support"); ok || d.HasChange("h3_support") {
		t, err := expandFirewallAccessProxy6ApiGatewayH3Support2edl(d, v, "h3_support")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-support"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_age"); ok || d.HasChange("http_cookie_age") {
		t, err := expandFirewallAccessProxy6ApiGatewayHttpCookieAge2edl(d, v, "http_cookie_age")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-age"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_domain"); ok || d.HasChange("http_cookie_domain") {
		t, err := expandFirewallAccessProxy6ApiGatewayHttpCookieDomain2edl(d, v, "http_cookie_domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-domain"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_domain_from_host"); ok || d.HasChange("http_cookie_domain_from_host") {
		t, err := expandFirewallAccessProxy6ApiGatewayHttpCookieDomainFromHost2edl(d, v, "http_cookie_domain_from_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-domain-from-host"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_generation"); ok || d.HasChange("http_cookie_generation") {
		t, err := expandFirewallAccessProxy6ApiGatewayHttpCookieGeneration2edl(d, v, "http_cookie_generation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-generation"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_path"); ok || d.HasChange("http_cookie_path") {
		t, err := expandFirewallAccessProxy6ApiGatewayHttpCookiePath2edl(d, v, "http_cookie_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-path"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_share"); ok || d.HasChange("http_cookie_share") {
		t, err := expandFirewallAccessProxy6ApiGatewayHttpCookieShare2edl(d, v, "http_cookie_share")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-share"] = t
		}
	}

	if v, ok := d.GetOk("https_cookie_secure"); ok || d.HasChange("https_cookie_secure") {
		t, err := expandFirewallAccessProxy6ApiGatewayHttpsCookieSecure2edl(d, v, "https_cookie_secure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-cookie-secure"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandFirewallAccessProxy6ApiGatewayId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ldb_method"); ok || d.HasChange("ldb_method") {
		t, err := expandFirewallAccessProxy6ApiGatewayLdbMethod2edl(d, v, "ldb_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldb-method"] = t
		}
	}

	if v, ok := d.GetOk("persistence"); ok || d.HasChange("persistence") {
		t, err := expandFirewallAccessProxy6ApiGatewayPersistence2edl(d, v, "persistence")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["persistence"] = t
		}
	}

	if v, ok := d.GetOk("quic"); ok || d.HasChange("quic") {
		t, err := expandFirewallAccessProxy6ApiGatewayQuic2edl(d, v, "quic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quic"] = t
		}
	}

	if v, ok := d.GetOk("realservers"); ok || d.HasChange("realservers") {
		t, err := expandFirewallAccessProxy6ApiGatewayRealservers2edl(d, v, "realservers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["realservers"] = t
		}
	}

	if v, ok := d.GetOk("saml_redirect"); ok || d.HasChange("saml_redirect") {
		t, err := expandFirewallAccessProxy6ApiGatewaySamlRedirect2edl(d, v, "saml_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["saml-redirect"] = t
		}
	}

	if v, ok := d.GetOk("saml_server"); ok || d.HasChange("saml_server") {
		t, err := expandFirewallAccessProxy6ApiGatewaySamlServer2edl(d, v, "saml_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["saml-server"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandFirewallAccessProxy6ApiGatewayService2edl(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("ssl_algorithm"); ok || d.HasChange("ssl_algorithm") {
		t, err := expandFirewallAccessProxy6ApiGatewaySslAlgorithm2edl(d, v, "ssl_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_cipher_suites"); ok || d.HasChange("ssl_cipher_suites") {
		t, err := expandFirewallAccessProxy6ApiGatewaySslCipherSuites2edl(d, v, "ssl_cipher_suites")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-cipher-suites"] = t
		}
	}

	if v, ok := d.GetOk("ssl_dh_bits"); ok || d.HasChange("ssl_dh_bits") {
		t, err := expandFirewallAccessProxy6ApiGatewaySslDhBits2edl(d, v, "ssl_dh_bits")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-dh-bits"] = t
		}
	}

	if v, ok := d.GetOk("ssl_max_version"); ok || d.HasChange("ssl_max_version") {
		t, err := expandFirewallAccessProxy6ApiGatewaySslMaxVersion2edl(d, v, "ssl_max_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-max-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_version"); ok || d.HasChange("ssl_min_version") {
		t, err := expandFirewallAccessProxy6ApiGatewaySslMinVersion2edl(d, v, "ssl_min_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_renegotiation"); ok || d.HasChange("ssl_renegotiation") {
		t, err := expandFirewallAccessProxy6ApiGatewaySslRenegotiation2edl(d, v, "ssl_renegotiation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-renegotiation"] = t
		}
	}

	if v, ok := d.GetOk("ssl_vpn_web_portal"); ok || d.HasChange("ssl_vpn_web_portal") {
		t, err := expandFirewallAccessProxy6ApiGatewaySslVpnWebPortal2edl(d, v, "ssl_vpn_web_portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-vpn-web-portal"] = t
		}
	}

	if v, ok := d.GetOk("url_map"); ok || d.HasChange("url_map") {
		t, err := expandFirewallAccessProxy6ApiGatewayUrlMap2edl(d, v, "url_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-map"] = t
		}
	}

	if v, ok := d.GetOk("url_map_type"); ok || d.HasChange("url_map_type") {
		t, err := expandFirewallAccessProxy6ApiGatewayUrlMapType2edl(d, v, "url_map_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-map-type"] = t
		}
	}

	if v, ok := d.GetOk("virtual_host"); ok || d.HasChange("virtual_host") {
		t, err := expandFirewallAccessProxy6ApiGatewayVirtualHost2edl(d, v, "virtual_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-host"] = t
		}
	}

	return &obj, nil
}
