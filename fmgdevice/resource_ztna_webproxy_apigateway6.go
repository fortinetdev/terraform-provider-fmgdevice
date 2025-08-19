// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Set IPv6 API Gateway.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceZtnaWebProxyApiGateway6() *schema.Resource {
	return &schema.Resource{
		Create: resourceZtnaWebProxyApiGateway6Create,
		Read:   resourceZtnaWebProxyApiGateway6Read,
		Update: resourceZtnaWebProxyApiGateway6Update,
		Delete: resourceZtnaWebProxyApiGateway6Delete,

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
			"web_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
				ForceNew: true,
				Optional: true,
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
						"port": &schema.Schema{
							Type:     schema.TypeInt,
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
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceZtnaWebProxyApiGateway6Create(d *schema.ResourceData, m interface{}) error {
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
	web_proxy := d.Get("web_proxy").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["web_proxy"] = web_proxy

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectZtnaWebProxyApiGateway6(d)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaWebProxyApiGateway6 resource while getting object: %v", err)
	}

	_, err = c.CreateZtnaWebProxyApiGateway6(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaWebProxyApiGateway6 resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceZtnaWebProxyApiGateway6Read(d, m)
}

func resourceZtnaWebProxyApiGateway6Update(d *schema.ResourceData, m interface{}) error {
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
	web_proxy := d.Get("web_proxy").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["web_proxy"] = web_proxy

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectZtnaWebProxyApiGateway6(d)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaWebProxyApiGateway6 resource while getting object: %v", err)
	}

	_, err = c.UpdateZtnaWebProxyApiGateway6(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaWebProxyApiGateway6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceZtnaWebProxyApiGateway6Read(d, m)
}

func resourceZtnaWebProxyApiGateway6Delete(d *schema.ResourceData, m interface{}) error {
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
	web_proxy := d.Get("web_proxy").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["web_proxy"] = web_proxy

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteZtnaWebProxyApiGateway6(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ZtnaWebProxyApiGateway6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceZtnaWebProxyApiGateway6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	web_proxy := d.Get("web_proxy").(string)
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
	if web_proxy == "" {
		web_proxy = importOptionChecking(m.(*FortiClient).Cfg, "web_proxy")
		if web_proxy == "" {
			return fmt.Errorf("Parameter web_proxy is missing")
		}
		if err = d.Set("web_proxy", web_proxy); err != nil {
			return fmt.Errorf("Error set params web_proxy: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["web_proxy"] = web_proxy

	o, err := c.ReadZtnaWebProxyApiGateway6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaWebProxyApiGateway6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectZtnaWebProxyApiGateway6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaWebProxyApiGateway6 resource from API: %v", err)
	}
	return nil
}

func flattenZtnaWebProxyApiGateway6H2Support2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6H3Support2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6HttpCookieAge2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6HttpCookieDomain2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6HttpCookieDomainFromHost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6HttpCookieGeneration2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6HttpCookiePath2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6HttpCookieShare2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6HttpsCookieSecure2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6Id2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6LdbMethod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6Persistence2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6Quic2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := i["ack-delay-exponent"]; ok {
		result["ack_delay_exponent"] = flattenZtnaWebProxyApiGateway6QuicAckDelayExponent2edl(i["ack-delay-exponent"], d, pre_append)
	}

	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := i["active-connection-id-limit"]; ok {
		result["active_connection_id_limit"] = flattenZtnaWebProxyApiGateway6QuicActiveConnectionIdLimit2edl(i["active-connection-id-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "active_migration"
	if _, ok := i["active-migration"]; ok {
		result["active_migration"] = flattenZtnaWebProxyApiGateway6QuicActiveMigration2edl(i["active-migration"], d, pre_append)
	}

	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := i["grease-quic-bit"]; ok {
		result["grease_quic_bit"] = flattenZtnaWebProxyApiGateway6QuicGreaseQuicBit2edl(i["grease-quic-bit"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := i["max-ack-delay"]; ok {
		result["max_ack_delay"] = flattenZtnaWebProxyApiGateway6QuicMaxAckDelay2edl(i["max-ack-delay"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := i["max-datagram-frame-size"]; ok {
		result["max_datagram_frame_size"] = flattenZtnaWebProxyApiGateway6QuicMaxDatagramFrameSize2edl(i["max-datagram-frame-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := i["max-idle-timeout"]; ok {
		result["max_idle_timeout"] = flattenZtnaWebProxyApiGateway6QuicMaxIdleTimeout2edl(i["max-idle-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := i["max-udp-payload-size"]; ok {
		result["max_udp_payload_size"] = flattenZtnaWebProxyApiGateway6QuicMaxUdpPayloadSize2edl(i["max-udp-payload-size"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenZtnaWebProxyApiGateway6QuicAckDelayExponent2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6QuicActiveConnectionIdLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6QuicActiveMigration2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6QuicGreaseQuicBit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6QuicMaxAckDelay2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6QuicMaxDatagramFrameSize2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6QuicMaxIdleTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6QuicMaxUdpPayloadSize2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6Realservers2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenZtnaWebProxyApiGateway6RealserversAddrType2edl(i["addr-type"], d, pre_append)
			tmp["addr_type"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway6-Realservers-AddrType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {
			v := flattenZtnaWebProxyApiGateway6RealserversAddress2edl(i["address"], d, pre_append)
			tmp["address"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway6-Realservers-Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := i["health-check"]; ok {
			v := flattenZtnaWebProxyApiGateway6RealserversHealthCheck2edl(i["health-check"], d, pre_append)
			tmp["health_check"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway6-Realservers-HealthCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := i["health-check-proto"]; ok {
			v := flattenZtnaWebProxyApiGateway6RealserversHealthCheckProto2edl(i["health-check-proto"], d, pre_append)
			tmp["health_check_proto"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway6-Realservers-HealthCheckProto")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := i["holddown-interval"]; ok {
			v := flattenZtnaWebProxyApiGateway6RealserversHolddownInterval2edl(i["holddown-interval"], d, pre_append)
			tmp["holddown_interval"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway6-Realservers-HolddownInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := i["http-host"]; ok {
			v := flattenZtnaWebProxyApiGateway6RealserversHttpHost2edl(i["http-host"], d, pre_append)
			tmp["http_host"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway6-Realservers-HttpHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenZtnaWebProxyApiGateway6RealserversId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway6-Realservers-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenZtnaWebProxyApiGateway6RealserversIp2edl(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway6-Realservers-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenZtnaWebProxyApiGateway6RealserversPort2edl(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway6-Realservers-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenZtnaWebProxyApiGateway6RealserversStatus2edl(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway6-Realservers-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := i["translate-host"]; ok {
			v := flattenZtnaWebProxyApiGateway6RealserversTranslateHost2edl(i["translate-host"], d, pre_append)
			tmp["translate_host"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway6-Realservers-TranslateHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenZtnaWebProxyApiGateway6RealserversWeight2edl(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway6-Realservers-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenZtnaWebProxyApiGateway6RealserversAddrType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaWebProxyApiGateway6RealserversHealthCheck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversHealthCheckProto2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversHolddownInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversHttpHost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversTranslateHost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6Service2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6SslAlgorithm2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6SslCipherSuites2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenZtnaWebProxyApiGateway6SslCipherSuitesCipher2edl(i["cipher"], d, pre_append)
			tmp["cipher"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway6-SslCipherSuites-Cipher")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenZtnaWebProxyApiGateway6SslCipherSuitesPriority2edl(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway6-SslCipherSuites-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {
			v := flattenZtnaWebProxyApiGateway6SslCipherSuitesVersions2edl(i["versions"], d, pre_append)
			tmp["versions"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway6-SslCipherSuites-Versions")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenZtnaWebProxyApiGateway6SslCipherSuitesCipher2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6SslCipherSuitesPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6SslCipherSuitesVersions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaWebProxyApiGateway6SslDhBits2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6SslMaxVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6SslMinVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6SslRenegotiation2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6UrlMap2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6UrlMapType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectZtnaWebProxyApiGateway6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("h2_support", flattenZtnaWebProxyApiGateway6H2Support2edl(o["h2-support"], d, "h2_support")); err != nil {
		if vv, ok := fortiAPIPatch(o["h2-support"], "ZtnaWebProxyApiGateway6-H2Support"); ok {
			if err = d.Set("h2_support", vv); err != nil {
				return fmt.Errorf("Error reading h2_support: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading h2_support: %v", err)
		}
	}

	if err = d.Set("h3_support", flattenZtnaWebProxyApiGateway6H3Support2edl(o["h3-support"], d, "h3_support")); err != nil {
		if vv, ok := fortiAPIPatch(o["h3-support"], "ZtnaWebProxyApiGateway6-H3Support"); ok {
			if err = d.Set("h3_support", vv); err != nil {
				return fmt.Errorf("Error reading h3_support: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading h3_support: %v", err)
		}
	}

	if err = d.Set("http_cookie_age", flattenZtnaWebProxyApiGateway6HttpCookieAge2edl(o["http-cookie-age"], d, "http_cookie_age")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-age"], "ZtnaWebProxyApiGateway6-HttpCookieAge"); ok {
			if err = d.Set("http_cookie_age", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_age: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_age: %v", err)
		}
	}

	if err = d.Set("http_cookie_domain", flattenZtnaWebProxyApiGateway6HttpCookieDomain2edl(o["http-cookie-domain"], d, "http_cookie_domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-domain"], "ZtnaWebProxyApiGateway6-HttpCookieDomain"); ok {
			if err = d.Set("http_cookie_domain", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_domain: %v", err)
		}
	}

	if err = d.Set("http_cookie_domain_from_host", flattenZtnaWebProxyApiGateway6HttpCookieDomainFromHost2edl(o["http-cookie-domain-from-host"], d, "http_cookie_domain_from_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-domain-from-host"], "ZtnaWebProxyApiGateway6-HttpCookieDomainFromHost"); ok {
			if err = d.Set("http_cookie_domain_from_host", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_domain_from_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_domain_from_host: %v", err)
		}
	}

	if err = d.Set("http_cookie_generation", flattenZtnaWebProxyApiGateway6HttpCookieGeneration2edl(o["http-cookie-generation"], d, "http_cookie_generation")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-generation"], "ZtnaWebProxyApiGateway6-HttpCookieGeneration"); ok {
			if err = d.Set("http_cookie_generation", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_generation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_generation: %v", err)
		}
	}

	if err = d.Set("http_cookie_path", flattenZtnaWebProxyApiGateway6HttpCookiePath2edl(o["http-cookie-path"], d, "http_cookie_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-path"], "ZtnaWebProxyApiGateway6-HttpCookiePath"); ok {
			if err = d.Set("http_cookie_path", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_path: %v", err)
		}
	}

	if err = d.Set("http_cookie_share", flattenZtnaWebProxyApiGateway6HttpCookieShare2edl(o["http-cookie-share"], d, "http_cookie_share")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-share"], "ZtnaWebProxyApiGateway6-HttpCookieShare"); ok {
			if err = d.Set("http_cookie_share", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_share: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_share: %v", err)
		}
	}

	if err = d.Set("https_cookie_secure", flattenZtnaWebProxyApiGateway6HttpsCookieSecure2edl(o["https-cookie-secure"], d, "https_cookie_secure")); err != nil {
		if vv, ok := fortiAPIPatch(o["https-cookie-secure"], "ZtnaWebProxyApiGateway6-HttpsCookieSecure"); ok {
			if err = d.Set("https_cookie_secure", vv); err != nil {
				return fmt.Errorf("Error reading https_cookie_secure: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading https_cookie_secure: %v", err)
		}
	}

	if err = d.Set("fosid", flattenZtnaWebProxyApiGateway6Id2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ZtnaWebProxyApiGateway6-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ldb_method", flattenZtnaWebProxyApiGateway6LdbMethod2edl(o["ldb-method"], d, "ldb_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldb-method"], "ZtnaWebProxyApiGateway6-LdbMethod"); ok {
			if err = d.Set("ldb_method", vv); err != nil {
				return fmt.Errorf("Error reading ldb_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldb_method: %v", err)
		}
	}

	if err = d.Set("persistence", flattenZtnaWebProxyApiGateway6Persistence2edl(o["persistence"], d, "persistence")); err != nil {
		if vv, ok := fortiAPIPatch(o["persistence"], "ZtnaWebProxyApiGateway6-Persistence"); ok {
			if err = d.Set("persistence", vv); err != nil {
				return fmt.Errorf("Error reading persistence: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading persistence: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("quic", flattenZtnaWebProxyApiGateway6Quic2edl(o["quic"], d, "quic")); err != nil {
			if vv, ok := fortiAPIPatch(o["quic"], "ZtnaWebProxyApiGateway6-Quic"); ok {
				if err = d.Set("quic", vv); err != nil {
					return fmt.Errorf("Error reading quic: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading quic: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("quic"); ok {
			if err = d.Set("quic", flattenZtnaWebProxyApiGateway6Quic2edl(o["quic"], d, "quic")); err != nil {
				if vv, ok := fortiAPIPatch(o["quic"], "ZtnaWebProxyApiGateway6-Quic"); ok {
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
		if err = d.Set("realservers", flattenZtnaWebProxyApiGateway6Realservers2edl(o["realservers"], d, "realservers")); err != nil {
			if vv, ok := fortiAPIPatch(o["realservers"], "ZtnaWebProxyApiGateway6-Realservers"); ok {
				if err = d.Set("realservers", vv); err != nil {
					return fmt.Errorf("Error reading realservers: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading realservers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("realservers"); ok {
			if err = d.Set("realservers", flattenZtnaWebProxyApiGateway6Realservers2edl(o["realservers"], d, "realservers")); err != nil {
				if vv, ok := fortiAPIPatch(o["realservers"], "ZtnaWebProxyApiGateway6-Realservers"); ok {
					if err = d.Set("realservers", vv); err != nil {
						return fmt.Errorf("Error reading realservers: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading realservers: %v", err)
				}
			}
		}
	}

	if err = d.Set("service", flattenZtnaWebProxyApiGateway6Service2edl(o["service"], d, "service")); err != nil {
		if vv, ok := fortiAPIPatch(o["service"], "ZtnaWebProxyApiGateway6-Service"); ok {
			if err = d.Set("service", vv); err != nil {
				return fmt.Errorf("Error reading service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("ssl_algorithm", flattenZtnaWebProxyApiGateway6SslAlgorithm2edl(o["ssl-algorithm"], d, "ssl_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-algorithm"], "ZtnaWebProxyApiGateway6-SslAlgorithm"); ok {
			if err = d.Set("ssl_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading ssl_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_algorithm: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ssl_cipher_suites", flattenZtnaWebProxyApiGateway6SslCipherSuites2edl(o["ssl-cipher-suites"], d, "ssl_cipher_suites")); err != nil {
			if vv, ok := fortiAPIPatch(o["ssl-cipher-suites"], "ZtnaWebProxyApiGateway6-SslCipherSuites"); ok {
				if err = d.Set("ssl_cipher_suites", vv); err != nil {
					return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_cipher_suites"); ok {
			if err = d.Set("ssl_cipher_suites", flattenZtnaWebProxyApiGateway6SslCipherSuites2edl(o["ssl-cipher-suites"], d, "ssl_cipher_suites")); err != nil {
				if vv, ok := fortiAPIPatch(o["ssl-cipher-suites"], "ZtnaWebProxyApiGateway6-SslCipherSuites"); ok {
					if err = d.Set("ssl_cipher_suites", vv); err != nil {
						return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_dh_bits", flattenZtnaWebProxyApiGateway6SslDhBits2edl(o["ssl-dh-bits"], d, "ssl_dh_bits")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-dh-bits"], "ZtnaWebProxyApiGateway6-SslDhBits"); ok {
			if err = d.Set("ssl_dh_bits", vv); err != nil {
				return fmt.Errorf("Error reading ssl_dh_bits: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_dh_bits: %v", err)
		}
	}

	if err = d.Set("ssl_max_version", flattenZtnaWebProxyApiGateway6SslMaxVersion2edl(o["ssl-max-version"], d, "ssl_max_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-max-version"], "ZtnaWebProxyApiGateway6-SslMaxVersion"); ok {
			if err = d.Set("ssl_max_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_max_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_max_version: %v", err)
		}
	}

	if err = d.Set("ssl_min_version", flattenZtnaWebProxyApiGateway6SslMinVersion2edl(o["ssl-min-version"], d, "ssl_min_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-min-version"], "ZtnaWebProxyApiGateway6-SslMinVersion"); ok {
			if err = d.Set("ssl_min_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_min_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_min_version: %v", err)
		}
	}

	if err = d.Set("ssl_renegotiation", flattenZtnaWebProxyApiGateway6SslRenegotiation2edl(o["ssl-renegotiation"], d, "ssl_renegotiation")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-renegotiation"], "ZtnaWebProxyApiGateway6-SslRenegotiation"); ok {
			if err = d.Set("ssl_renegotiation", vv); err != nil {
				return fmt.Errorf("Error reading ssl_renegotiation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_renegotiation: %v", err)
		}
	}

	if err = d.Set("url_map", flattenZtnaWebProxyApiGateway6UrlMap2edl(o["url-map"], d, "url_map")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-map"], "ZtnaWebProxyApiGateway6-UrlMap"); ok {
			if err = d.Set("url_map", vv); err != nil {
				return fmt.Errorf("Error reading url_map: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_map: %v", err)
		}
	}

	if err = d.Set("url_map_type", flattenZtnaWebProxyApiGateway6UrlMapType2edl(o["url-map-type"], d, "url_map_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-map-type"], "ZtnaWebProxyApiGateway6-UrlMapType"); ok {
			if err = d.Set("url_map_type", vv); err != nil {
				return fmt.Errorf("Error reading url_map_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_map_type: %v", err)
		}
	}

	return nil
}

func flattenZtnaWebProxyApiGateway6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandZtnaWebProxyApiGateway6H2Support2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6H3Support2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6HttpCookieAge2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6HttpCookieDomain2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6HttpCookieDomainFromHost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6HttpCookieGeneration2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6HttpCookiePath2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6HttpCookieShare2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6HttpsCookieSecure2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6Id2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6LdbMethod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6Persistence2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6Quic2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ack-delay-exponent"], _ = expandZtnaWebProxyApiGateway6QuicAckDelayExponent2edl(d, i["ack_delay_exponent"], pre_append)
	}
	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["active-connection-id-limit"], _ = expandZtnaWebProxyApiGateway6QuicActiveConnectionIdLimit2edl(d, i["active_connection_id_limit"], pre_append)
	}
	pre_append = pre + ".0." + "active_migration"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["active-migration"], _ = expandZtnaWebProxyApiGateway6QuicActiveMigration2edl(d, i["active_migration"], pre_append)
	}
	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["grease-quic-bit"], _ = expandZtnaWebProxyApiGateway6QuicGreaseQuicBit2edl(d, i["grease_quic_bit"], pre_append)
	}
	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-ack-delay"], _ = expandZtnaWebProxyApiGateway6QuicMaxAckDelay2edl(d, i["max_ack_delay"], pre_append)
	}
	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-datagram-frame-size"], _ = expandZtnaWebProxyApiGateway6QuicMaxDatagramFrameSize2edl(d, i["max_datagram_frame_size"], pre_append)
	}
	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-idle-timeout"], _ = expandZtnaWebProxyApiGateway6QuicMaxIdleTimeout2edl(d, i["max_idle_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-udp-payload-size"], _ = expandZtnaWebProxyApiGateway6QuicMaxUdpPayloadSize2edl(d, i["max_udp_payload_size"], pre_append)
	}

	return result, nil
}

func expandZtnaWebProxyApiGateway6QuicAckDelayExponent2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6QuicActiveConnectionIdLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6QuicActiveMigration2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6QuicGreaseQuicBit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6QuicMaxAckDelay2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6QuicMaxDatagramFrameSize2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6QuicMaxIdleTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6QuicMaxUdpPayloadSize2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6Realservers2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["addr-type"], _ = expandZtnaWebProxyApiGateway6RealserversAddrType2edl(d, i["addr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["address"], _ = expandZtnaWebProxyApiGateway6RealserversAddress2edl(d, i["address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check"], _ = expandZtnaWebProxyApiGateway6RealserversHealthCheck2edl(d, i["health_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check-proto"], _ = expandZtnaWebProxyApiGateway6RealserversHealthCheckProto2edl(d, i["health_check_proto"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["holddown-interval"], _ = expandZtnaWebProxyApiGateway6RealserversHolddownInterval2edl(d, i["holddown_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-host"], _ = expandZtnaWebProxyApiGateway6RealserversHttpHost2edl(d, i["http_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandZtnaWebProxyApiGateway6RealserversId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandZtnaWebProxyApiGateway6RealserversIp2edl(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandZtnaWebProxyApiGateway6RealserversPort2edl(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandZtnaWebProxyApiGateway6RealserversStatus2edl(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["translate-host"], _ = expandZtnaWebProxyApiGateway6RealserversTranslateHost2edl(d, i["translate_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandZtnaWebProxyApiGateway6RealserversWeight2edl(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandZtnaWebProxyApiGateway6RealserversAddrType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaWebProxyApiGateway6RealserversHealthCheck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversHealthCheckProto2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversHolddownInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversHttpHost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversTranslateHost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6Service2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6SslAlgorithm2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6SslCipherSuites2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["cipher"], _ = expandZtnaWebProxyApiGateway6SslCipherSuitesCipher2edl(d, i["cipher"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandZtnaWebProxyApiGateway6SslCipherSuitesPriority2edl(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["versions"], _ = expandZtnaWebProxyApiGateway6SslCipherSuitesVersions2edl(d, i["versions"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandZtnaWebProxyApiGateway6SslCipherSuitesCipher2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6SslCipherSuitesPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6SslCipherSuitesVersions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaWebProxyApiGateway6SslDhBits2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6SslMaxVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6SslMinVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6SslRenegotiation2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6UrlMap2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6UrlMapType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectZtnaWebProxyApiGateway6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("h2_support"); ok || d.HasChange("h2_support") {
		t, err := expandZtnaWebProxyApiGateway6H2Support2edl(d, v, "h2_support")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h2-support"] = t
		}
	}

	if v, ok := d.GetOk("h3_support"); ok || d.HasChange("h3_support") {
		t, err := expandZtnaWebProxyApiGateway6H3Support2edl(d, v, "h3_support")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-support"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_age"); ok || d.HasChange("http_cookie_age") {
		t, err := expandZtnaWebProxyApiGateway6HttpCookieAge2edl(d, v, "http_cookie_age")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-age"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_domain"); ok || d.HasChange("http_cookie_domain") {
		t, err := expandZtnaWebProxyApiGateway6HttpCookieDomain2edl(d, v, "http_cookie_domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-domain"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_domain_from_host"); ok || d.HasChange("http_cookie_domain_from_host") {
		t, err := expandZtnaWebProxyApiGateway6HttpCookieDomainFromHost2edl(d, v, "http_cookie_domain_from_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-domain-from-host"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_generation"); ok || d.HasChange("http_cookie_generation") {
		t, err := expandZtnaWebProxyApiGateway6HttpCookieGeneration2edl(d, v, "http_cookie_generation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-generation"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_path"); ok || d.HasChange("http_cookie_path") {
		t, err := expandZtnaWebProxyApiGateway6HttpCookiePath2edl(d, v, "http_cookie_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-path"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_share"); ok || d.HasChange("http_cookie_share") {
		t, err := expandZtnaWebProxyApiGateway6HttpCookieShare2edl(d, v, "http_cookie_share")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-share"] = t
		}
	}

	if v, ok := d.GetOk("https_cookie_secure"); ok || d.HasChange("https_cookie_secure") {
		t, err := expandZtnaWebProxyApiGateway6HttpsCookieSecure2edl(d, v, "https_cookie_secure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-cookie-secure"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandZtnaWebProxyApiGateway6Id2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ldb_method"); ok || d.HasChange("ldb_method") {
		t, err := expandZtnaWebProxyApiGateway6LdbMethod2edl(d, v, "ldb_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldb-method"] = t
		}
	}

	if v, ok := d.GetOk("persistence"); ok || d.HasChange("persistence") {
		t, err := expandZtnaWebProxyApiGateway6Persistence2edl(d, v, "persistence")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["persistence"] = t
		}
	}

	if v, ok := d.GetOk("quic"); ok || d.HasChange("quic") {
		t, err := expandZtnaWebProxyApiGateway6Quic2edl(d, v, "quic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quic"] = t
		}
	}

	if v, ok := d.GetOk("realservers"); ok || d.HasChange("realservers") {
		t, err := expandZtnaWebProxyApiGateway6Realservers2edl(d, v, "realservers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["realservers"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandZtnaWebProxyApiGateway6Service2edl(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("ssl_algorithm"); ok || d.HasChange("ssl_algorithm") {
		t, err := expandZtnaWebProxyApiGateway6SslAlgorithm2edl(d, v, "ssl_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_cipher_suites"); ok || d.HasChange("ssl_cipher_suites") {
		t, err := expandZtnaWebProxyApiGateway6SslCipherSuites2edl(d, v, "ssl_cipher_suites")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-cipher-suites"] = t
		}
	}

	if v, ok := d.GetOk("ssl_dh_bits"); ok || d.HasChange("ssl_dh_bits") {
		t, err := expandZtnaWebProxyApiGateway6SslDhBits2edl(d, v, "ssl_dh_bits")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-dh-bits"] = t
		}
	}

	if v, ok := d.GetOk("ssl_max_version"); ok || d.HasChange("ssl_max_version") {
		t, err := expandZtnaWebProxyApiGateway6SslMaxVersion2edl(d, v, "ssl_max_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-max-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_version"); ok || d.HasChange("ssl_min_version") {
		t, err := expandZtnaWebProxyApiGateway6SslMinVersion2edl(d, v, "ssl_min_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_renegotiation"); ok || d.HasChange("ssl_renegotiation") {
		t, err := expandZtnaWebProxyApiGateway6SslRenegotiation2edl(d, v, "ssl_renegotiation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-renegotiation"] = t
		}
	}

	if v, ok := d.GetOk("url_map"); ok || d.HasChange("url_map") {
		t, err := expandZtnaWebProxyApiGateway6UrlMap2edl(d, v, "url_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-map"] = t
		}
	}

	if v, ok := d.GetOk("url_map_type"); ok || d.HasChange("url_map_type") {
		t, err := expandZtnaWebProxyApiGateway6UrlMapType2edl(d, v, "url_map_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-map-type"] = t
		}
	}

	return &obj, nil
}
