// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure ZTNA web-proxy.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceZtnaWebProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceZtnaWebProxyCreate,
		Read:   resourceZtnaWebProxyRead,
		Update: resourceZtnaWebProxyUpdate,
		Delete: resourceZtnaWebProxyDelete,

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
			"api_gateway": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"id": &schema.Schema{
							Type:     schema.TypeInt,
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
					},
				},
			},
			"api_gateway6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"id": &schema.Schema{
							Type:     schema.TypeInt,
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
					},
				},
			},
			"auth_portal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_virtual_host": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"decrypted_traffic_mirror": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"host": &schema.Schema{
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
			"vip": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vip6": &schema.Schema{
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

func resourceZtnaWebProxyCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectZtnaWebProxy(d)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaWebProxy resource while getting object: %v", err)
	}

	_, err = c.CreateZtnaWebProxy(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaWebProxy resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceZtnaWebProxyRead(d, m)
}

func resourceZtnaWebProxyUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectZtnaWebProxy(d)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaWebProxy resource while getting object: %v", err)
	}

	_, err = c.UpdateZtnaWebProxy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaWebProxy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceZtnaWebProxyRead(d, m)
}

func resourceZtnaWebProxyDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteZtnaWebProxy(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ZtnaWebProxy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceZtnaWebProxyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadZtnaWebProxy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaWebProxy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectZtnaWebProxy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaWebProxy resource from API: %v", err)
	}
	return nil
}

func flattenZtnaWebProxyApiGateway(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h2_support"
		if _, ok := i["h2-support"]; ok {
			v := flattenZtnaWebProxyApiGatewayH2Support(i["h2-support"], d, pre_append)
			tmp["h2_support"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway-H2Support")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h3_support"
		if _, ok := i["h3-support"]; ok {
			v := flattenZtnaWebProxyApiGatewayH3Support(i["h3-support"], d, pre_append)
			tmp["h3_support"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway-H3Support")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := i["http-cookie-age"]; ok {
			v := flattenZtnaWebProxyApiGatewayHttpCookieAge(i["http-cookie-age"], d, pre_append)
			tmp["http_cookie_age"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway-HttpCookieAge")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := i["http-cookie-domain"]; ok {
			v := flattenZtnaWebProxyApiGatewayHttpCookieDomain(i["http-cookie-domain"], d, pre_append)
			tmp["http_cookie_domain"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway-HttpCookieDomain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := i["http-cookie-domain-from-host"]; ok {
			v := flattenZtnaWebProxyApiGatewayHttpCookieDomainFromHost(i["http-cookie-domain-from-host"], d, pre_append)
			tmp["http_cookie_domain_from_host"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway-HttpCookieDomainFromHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := i["http-cookie-generation"]; ok {
			v := flattenZtnaWebProxyApiGatewayHttpCookieGeneration(i["http-cookie-generation"], d, pre_append)
			tmp["http_cookie_generation"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway-HttpCookieGeneration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := i["http-cookie-path"]; ok {
			v := flattenZtnaWebProxyApiGatewayHttpCookiePath(i["http-cookie-path"], d, pre_append)
			tmp["http_cookie_path"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway-HttpCookiePath")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := i["http-cookie-share"]; ok {
			v := flattenZtnaWebProxyApiGatewayHttpCookieShare(i["http-cookie-share"], d, pre_append)
			tmp["http_cookie_share"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway-HttpCookieShare")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := i["https-cookie-secure"]; ok {
			v := flattenZtnaWebProxyApiGatewayHttpsCookieSecure(i["https-cookie-secure"], d, pre_append)
			tmp["https_cookie_secure"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway-HttpsCookieSecure")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenZtnaWebProxyApiGatewayId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := i["ldb-method"]; ok {
			v := flattenZtnaWebProxyApiGatewayLdbMethod(i["ldb-method"], d, pre_append)
			tmp["ldb_method"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway-LdbMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := i["persistence"]; ok {
			v := flattenZtnaWebProxyApiGatewayPersistence(i["persistence"], d, pre_append)
			tmp["persistence"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway-Persistence")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quic"
		if _, ok := i["quic"]; ok {
			v := flattenZtnaWebProxyApiGatewayQuic(i["quic"], d, pre_append)
			tmp["quic"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway-Quic")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := i["realservers"]; ok {
			v := flattenZtnaWebProxyApiGatewayRealservers(i["realservers"], d, pre_append)
			tmp["realservers"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway-Realservers")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := i["service"]; ok {
			v := flattenZtnaWebProxyApiGatewayService(i["service"], d, pre_append)
			tmp["service"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway-Service")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := i["ssl-algorithm"]; ok {
			v := flattenZtnaWebProxyApiGatewaySslAlgorithm(i["ssl-algorithm"], d, pre_append)
			tmp["ssl_algorithm"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway-SslAlgorithm")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := i["ssl-cipher-suites"]; ok {
			v := flattenZtnaWebProxyApiGatewaySslCipherSuites(i["ssl-cipher-suites"], d, pre_append)
			tmp["ssl_cipher_suites"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway-SslCipherSuites")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := i["ssl-dh-bits"]; ok {
			v := flattenZtnaWebProxyApiGatewaySslDhBits(i["ssl-dh-bits"], d, pre_append)
			tmp["ssl_dh_bits"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway-SslDhBits")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := i["ssl-max-version"]; ok {
			v := flattenZtnaWebProxyApiGatewaySslMaxVersion(i["ssl-max-version"], d, pre_append)
			tmp["ssl_max_version"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway-SslMaxVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := i["ssl-min-version"]; ok {
			v := flattenZtnaWebProxyApiGatewaySslMinVersion(i["ssl-min-version"], d, pre_append)
			tmp["ssl_min_version"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway-SslMinVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_renegotiation"
		if _, ok := i["ssl-renegotiation"]; ok {
			v := flattenZtnaWebProxyApiGatewaySslRenegotiation(i["ssl-renegotiation"], d, pre_append)
			tmp["ssl_renegotiation"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway-SslRenegotiation")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map"
		if _, ok := i["url-map"]; ok {
			v := flattenZtnaWebProxyApiGatewayUrlMap(i["url-map"], d, pre_append)
			tmp["url_map"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway-UrlMap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map_type"
		if _, ok := i["url-map-type"]; ok {
			v := flattenZtnaWebProxyApiGatewayUrlMapType(i["url-map-type"], d, pre_append)
			tmp["url_map_type"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway-UrlMapType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenZtnaWebProxyApiGatewayH2Support(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayH3Support(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayHttpCookieAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayHttpCookieDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayHttpCookieDomainFromHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayHttpCookieGeneration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayHttpCookiePath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayHttpCookieShare(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayHttpsCookieSecure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayLdbMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayPersistence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayQuic(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := i["ack-delay-exponent"]; ok {
		result["ack_delay_exponent"] = flattenZtnaWebProxyApiGatewayQuicAckDelayExponent(i["ack-delay-exponent"], d, pre_append)
	}

	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := i["active-connection-id-limit"]; ok {
		result["active_connection_id_limit"] = flattenZtnaWebProxyApiGatewayQuicActiveConnectionIdLimit(i["active-connection-id-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "active_migration"
	if _, ok := i["active-migration"]; ok {
		result["active_migration"] = flattenZtnaWebProxyApiGatewayQuicActiveMigration(i["active-migration"], d, pre_append)
	}

	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := i["grease-quic-bit"]; ok {
		result["grease_quic_bit"] = flattenZtnaWebProxyApiGatewayQuicGreaseQuicBit(i["grease-quic-bit"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := i["max-ack-delay"]; ok {
		result["max_ack_delay"] = flattenZtnaWebProxyApiGatewayQuicMaxAckDelay(i["max-ack-delay"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := i["max-datagram-frame-size"]; ok {
		result["max_datagram_frame_size"] = flattenZtnaWebProxyApiGatewayQuicMaxDatagramFrameSize(i["max-datagram-frame-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := i["max-idle-timeout"]; ok {
		result["max_idle_timeout"] = flattenZtnaWebProxyApiGatewayQuicMaxIdleTimeout(i["max-idle-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := i["max-udp-payload-size"]; ok {
		result["max_udp_payload_size"] = flattenZtnaWebProxyApiGatewayQuicMaxUdpPayloadSize(i["max-udp-payload-size"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenZtnaWebProxyApiGatewayQuicAckDelayExponent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayQuicActiveConnectionIdLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayQuicActiveMigration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayQuicGreaseQuicBit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayQuicMaxAckDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayQuicMaxDatagramFrameSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayQuicMaxIdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayQuicMaxUdpPayloadSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealservers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenZtnaWebProxyApiGatewayRealserversAddrType(i["addr-type"], d, pre_append)
			tmp["addr_type"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway-Realservers-AddrType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {
			v := flattenZtnaWebProxyApiGatewayRealserversAddress(i["address"], d, pre_append)
			tmp["address"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway-Realservers-Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := i["health-check"]; ok {
			v := flattenZtnaWebProxyApiGatewayRealserversHealthCheck(i["health-check"], d, pre_append)
			tmp["health_check"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway-Realservers-HealthCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := i["health-check-proto"]; ok {
			v := flattenZtnaWebProxyApiGatewayRealserversHealthCheckProto(i["health-check-proto"], d, pre_append)
			tmp["health_check_proto"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway-Realservers-HealthCheckProto")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := i["holddown-interval"]; ok {
			v := flattenZtnaWebProxyApiGatewayRealserversHolddownInterval(i["holddown-interval"], d, pre_append)
			tmp["holddown_interval"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway-Realservers-HolddownInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := i["http-host"]; ok {
			v := flattenZtnaWebProxyApiGatewayRealserversHttpHost(i["http-host"], d, pre_append)
			tmp["http_host"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway-Realservers-HttpHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenZtnaWebProxyApiGatewayRealserversId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway-Realservers-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenZtnaWebProxyApiGatewayRealserversIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway-Realservers-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenZtnaWebProxyApiGatewayRealserversPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway-Realservers-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenZtnaWebProxyApiGatewayRealserversStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway-Realservers-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := i["translate-host"]; ok {
			v := flattenZtnaWebProxyApiGatewayRealserversTranslateHost(i["translate-host"], d, pre_append)
			tmp["translate_host"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway-Realservers-TranslateHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenZtnaWebProxyApiGatewayRealserversWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway-Realservers-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenZtnaWebProxyApiGatewayRealserversAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealserversAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaWebProxyApiGatewayRealserversHealthCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealserversHealthCheckProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealserversHolddownInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealserversHttpHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealserversId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealserversIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealserversPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealserversStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealserversTranslateHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealserversWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewaySslAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewaySslCipherSuites(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenZtnaWebProxyApiGatewaySslCipherSuitesCipher(i["cipher"], d, pre_append)
			tmp["cipher"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway-SslCipherSuites-Cipher")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenZtnaWebProxyApiGatewaySslCipherSuitesPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway-SslCipherSuites-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {
			v := flattenZtnaWebProxyApiGatewaySslCipherSuitesVersions(i["versions"], d, pre_append)
			tmp["versions"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway-SslCipherSuites-Versions")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenZtnaWebProxyApiGatewaySslCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewaySslCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewaySslCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaWebProxyApiGatewaySslDhBits(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewaySslMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewaySslMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewaySslRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayUrlMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayUrlMapType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h2_support"
		if _, ok := i["h2-support"]; ok {
			v := flattenZtnaWebProxyApiGateway6H2Support(i["h2-support"], d, pre_append)
			tmp["h2_support"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway6-H2Support")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h3_support"
		if _, ok := i["h3-support"]; ok {
			v := flattenZtnaWebProxyApiGateway6H3Support(i["h3-support"], d, pre_append)
			tmp["h3_support"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway6-H3Support")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := i["http-cookie-age"]; ok {
			v := flattenZtnaWebProxyApiGateway6HttpCookieAge(i["http-cookie-age"], d, pre_append)
			tmp["http_cookie_age"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway6-HttpCookieAge")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := i["http-cookie-domain"]; ok {
			v := flattenZtnaWebProxyApiGateway6HttpCookieDomain(i["http-cookie-domain"], d, pre_append)
			tmp["http_cookie_domain"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway6-HttpCookieDomain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := i["http-cookie-domain-from-host"]; ok {
			v := flattenZtnaWebProxyApiGateway6HttpCookieDomainFromHost(i["http-cookie-domain-from-host"], d, pre_append)
			tmp["http_cookie_domain_from_host"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway6-HttpCookieDomainFromHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := i["http-cookie-generation"]; ok {
			v := flattenZtnaWebProxyApiGateway6HttpCookieGeneration(i["http-cookie-generation"], d, pre_append)
			tmp["http_cookie_generation"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway6-HttpCookieGeneration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := i["http-cookie-path"]; ok {
			v := flattenZtnaWebProxyApiGateway6HttpCookiePath(i["http-cookie-path"], d, pre_append)
			tmp["http_cookie_path"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway6-HttpCookiePath")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := i["http-cookie-share"]; ok {
			v := flattenZtnaWebProxyApiGateway6HttpCookieShare(i["http-cookie-share"], d, pre_append)
			tmp["http_cookie_share"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway6-HttpCookieShare")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := i["https-cookie-secure"]; ok {
			v := flattenZtnaWebProxyApiGateway6HttpsCookieSecure(i["https-cookie-secure"], d, pre_append)
			tmp["https_cookie_secure"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway6-HttpsCookieSecure")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenZtnaWebProxyApiGateway6Id(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway6-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := i["ldb-method"]; ok {
			v := flattenZtnaWebProxyApiGateway6LdbMethod(i["ldb-method"], d, pre_append)
			tmp["ldb_method"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway6-LdbMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := i["persistence"]; ok {
			v := flattenZtnaWebProxyApiGateway6Persistence(i["persistence"], d, pre_append)
			tmp["persistence"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway6-Persistence")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quic"
		if _, ok := i["quic"]; ok {
			v := flattenZtnaWebProxyApiGateway6Quic(i["quic"], d, pre_append)
			tmp["quic"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway6-Quic")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := i["realservers"]; ok {
			v := flattenZtnaWebProxyApiGateway6Realservers(i["realservers"], d, pre_append)
			tmp["realservers"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway6-Realservers")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := i["service"]; ok {
			v := flattenZtnaWebProxyApiGateway6Service(i["service"], d, pre_append)
			tmp["service"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway6-Service")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := i["ssl-algorithm"]; ok {
			v := flattenZtnaWebProxyApiGateway6SslAlgorithm(i["ssl-algorithm"], d, pre_append)
			tmp["ssl_algorithm"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway6-SslAlgorithm")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := i["ssl-cipher-suites"]; ok {
			v := flattenZtnaWebProxyApiGateway6SslCipherSuites(i["ssl-cipher-suites"], d, pre_append)
			tmp["ssl_cipher_suites"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway6-SslCipherSuites")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := i["ssl-dh-bits"]; ok {
			v := flattenZtnaWebProxyApiGateway6SslDhBits(i["ssl-dh-bits"], d, pre_append)
			tmp["ssl_dh_bits"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway6-SslDhBits")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := i["ssl-max-version"]; ok {
			v := flattenZtnaWebProxyApiGateway6SslMaxVersion(i["ssl-max-version"], d, pre_append)
			tmp["ssl_max_version"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway6-SslMaxVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := i["ssl-min-version"]; ok {
			v := flattenZtnaWebProxyApiGateway6SslMinVersion(i["ssl-min-version"], d, pre_append)
			tmp["ssl_min_version"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway6-SslMinVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_renegotiation"
		if _, ok := i["ssl-renegotiation"]; ok {
			v := flattenZtnaWebProxyApiGateway6SslRenegotiation(i["ssl-renegotiation"], d, pre_append)
			tmp["ssl_renegotiation"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway6-SslRenegotiation")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map"
		if _, ok := i["url-map"]; ok {
			v := flattenZtnaWebProxyApiGateway6UrlMap(i["url-map"], d, pre_append)
			tmp["url_map"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway6-UrlMap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map_type"
		if _, ok := i["url-map-type"]; ok {
			v := flattenZtnaWebProxyApiGateway6UrlMapType(i["url-map-type"], d, pre_append)
			tmp["url_map_type"] = fortiAPISubPartPatch(v, "ZtnaWebProxy-ApiGateway6-UrlMapType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenZtnaWebProxyApiGateway6H2Support(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6H3Support(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6HttpCookieAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6HttpCookieDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6HttpCookieDomainFromHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6HttpCookieGeneration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6HttpCookiePath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6HttpCookieShare(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6HttpsCookieSecure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6LdbMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6Persistence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6Quic(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := i["ack-delay-exponent"]; ok {
		result["ack_delay_exponent"] = flattenZtnaWebProxyApiGateway6QuicAckDelayExponent(i["ack-delay-exponent"], d, pre_append)
	}

	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := i["active-connection-id-limit"]; ok {
		result["active_connection_id_limit"] = flattenZtnaWebProxyApiGateway6QuicActiveConnectionIdLimit(i["active-connection-id-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "active_migration"
	if _, ok := i["active-migration"]; ok {
		result["active_migration"] = flattenZtnaWebProxyApiGateway6QuicActiveMigration(i["active-migration"], d, pre_append)
	}

	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := i["grease-quic-bit"]; ok {
		result["grease_quic_bit"] = flattenZtnaWebProxyApiGateway6QuicGreaseQuicBit(i["grease-quic-bit"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := i["max-ack-delay"]; ok {
		result["max_ack_delay"] = flattenZtnaWebProxyApiGateway6QuicMaxAckDelay(i["max-ack-delay"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := i["max-datagram-frame-size"]; ok {
		result["max_datagram_frame_size"] = flattenZtnaWebProxyApiGateway6QuicMaxDatagramFrameSize(i["max-datagram-frame-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := i["max-idle-timeout"]; ok {
		result["max_idle_timeout"] = flattenZtnaWebProxyApiGateway6QuicMaxIdleTimeout(i["max-idle-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := i["max-udp-payload-size"]; ok {
		result["max_udp_payload_size"] = flattenZtnaWebProxyApiGateway6QuicMaxUdpPayloadSize(i["max-udp-payload-size"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenZtnaWebProxyApiGateway6QuicAckDelayExponent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6QuicActiveConnectionIdLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6QuicActiveMigration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6QuicGreaseQuicBit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6QuicMaxAckDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6QuicMaxDatagramFrameSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6QuicMaxIdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6QuicMaxUdpPayloadSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6Realservers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenZtnaWebProxyApiGateway6RealserversAddrType(i["addr-type"], d, pre_append)
			tmp["addr_type"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway6-Realservers-AddrType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {
			v := flattenZtnaWebProxyApiGateway6RealserversAddress(i["address"], d, pre_append)
			tmp["address"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway6-Realservers-Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := i["health-check"]; ok {
			v := flattenZtnaWebProxyApiGateway6RealserversHealthCheck(i["health-check"], d, pre_append)
			tmp["health_check"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway6-Realservers-HealthCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := i["health-check-proto"]; ok {
			v := flattenZtnaWebProxyApiGateway6RealserversHealthCheckProto(i["health-check-proto"], d, pre_append)
			tmp["health_check_proto"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway6-Realservers-HealthCheckProto")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := i["holddown-interval"]; ok {
			v := flattenZtnaWebProxyApiGateway6RealserversHolddownInterval(i["holddown-interval"], d, pre_append)
			tmp["holddown_interval"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway6-Realservers-HolddownInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := i["http-host"]; ok {
			v := flattenZtnaWebProxyApiGateway6RealserversHttpHost(i["http-host"], d, pre_append)
			tmp["http_host"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway6-Realservers-HttpHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenZtnaWebProxyApiGateway6RealserversId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway6-Realservers-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenZtnaWebProxyApiGateway6RealserversIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway6-Realservers-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenZtnaWebProxyApiGateway6RealserversPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway6-Realservers-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenZtnaWebProxyApiGateway6RealserversStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway6-Realservers-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := i["translate-host"]; ok {
			v := flattenZtnaWebProxyApiGateway6RealserversTranslateHost(i["translate-host"], d, pre_append)
			tmp["translate_host"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway6-Realservers-TranslateHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenZtnaWebProxyApiGateway6RealserversWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway6-Realservers-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenZtnaWebProxyApiGateway6RealserversAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaWebProxyApiGateway6RealserversHealthCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversHealthCheckProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversHolddownInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversHttpHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversTranslateHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6Service(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6SslAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6SslCipherSuites(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenZtnaWebProxyApiGateway6SslCipherSuitesCipher(i["cipher"], d, pre_append)
			tmp["cipher"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway6-SslCipherSuites-Cipher")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenZtnaWebProxyApiGateway6SslCipherSuitesPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway6-SslCipherSuites-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {
			v := flattenZtnaWebProxyApiGateway6SslCipherSuitesVersions(i["versions"], d, pre_append)
			tmp["versions"] = fortiAPISubPartPatch(v, "ZtnaWebProxyApiGateway6-SslCipherSuites-Versions")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenZtnaWebProxyApiGateway6SslCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6SslCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6SslCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaWebProxyApiGateway6SslDhBits(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6SslMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6SslMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6SslRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6UrlMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6UrlMapType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyAuthPortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyAuthVirtualHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaWebProxyDecryptedTrafficMirror(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaWebProxyHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaWebProxyLogBlockedTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxySvrPoolMultiplex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxySvrPoolServerMaxConcurrentRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxySvrPoolServerMaxRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxySvrPoolTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyVip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaWebProxyVip6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectZtnaWebProxy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("api_gateway", flattenZtnaWebProxyApiGateway(o["api-gateway"], d, "api_gateway")); err != nil {
			if vv, ok := fortiAPIPatch(o["api-gateway"], "ZtnaWebProxy-ApiGateway"); ok {
				if err = d.Set("api_gateway", vv); err != nil {
					return fmt.Errorf("Error reading api_gateway: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading api_gateway: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("api_gateway"); ok {
			if err = d.Set("api_gateway", flattenZtnaWebProxyApiGateway(o["api-gateway"], d, "api_gateway")); err != nil {
				if vv, ok := fortiAPIPatch(o["api-gateway"], "ZtnaWebProxy-ApiGateway"); ok {
					if err = d.Set("api_gateway", vv); err != nil {
						return fmt.Errorf("Error reading api_gateway: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading api_gateway: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("api_gateway6", flattenZtnaWebProxyApiGateway6(o["api-gateway6"], d, "api_gateway6")); err != nil {
			if vv, ok := fortiAPIPatch(o["api-gateway6"], "ZtnaWebProxy-ApiGateway6"); ok {
				if err = d.Set("api_gateway6", vv); err != nil {
					return fmt.Errorf("Error reading api_gateway6: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading api_gateway6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("api_gateway6"); ok {
			if err = d.Set("api_gateway6", flattenZtnaWebProxyApiGateway6(o["api-gateway6"], d, "api_gateway6")); err != nil {
				if vv, ok := fortiAPIPatch(o["api-gateway6"], "ZtnaWebProxy-ApiGateway6"); ok {
					if err = d.Set("api_gateway6", vv); err != nil {
						return fmt.Errorf("Error reading api_gateway6: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading api_gateway6: %v", err)
				}
			}
		}
	}

	if err = d.Set("auth_portal", flattenZtnaWebProxyAuthPortal(o["auth-portal"], d, "auth_portal")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-portal"], "ZtnaWebProxy-AuthPortal"); ok {
			if err = d.Set("auth_portal", vv); err != nil {
				return fmt.Errorf("Error reading auth_portal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_portal: %v", err)
		}
	}

	if err = d.Set("auth_virtual_host", flattenZtnaWebProxyAuthVirtualHost(o["auth-virtual-host"], d, "auth_virtual_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-virtual-host"], "ZtnaWebProxy-AuthVirtualHost"); ok {
			if err = d.Set("auth_virtual_host", vv); err != nil {
				return fmt.Errorf("Error reading auth_virtual_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_virtual_host: %v", err)
		}
	}

	if err = d.Set("decrypted_traffic_mirror", flattenZtnaWebProxyDecryptedTrafficMirror(o["decrypted-traffic-mirror"], d, "decrypted_traffic_mirror")); err != nil {
		if vv, ok := fortiAPIPatch(o["decrypted-traffic-mirror"], "ZtnaWebProxy-DecryptedTrafficMirror"); ok {
			if err = d.Set("decrypted_traffic_mirror", vv); err != nil {
				return fmt.Errorf("Error reading decrypted_traffic_mirror: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading decrypted_traffic_mirror: %v", err)
		}
	}

	if err = d.Set("host", flattenZtnaWebProxyHost(o["host"], d, "host")); err != nil {
		if vv, ok := fortiAPIPatch(o["host"], "ZtnaWebProxy-Host"); ok {
			if err = d.Set("host", vv); err != nil {
				return fmt.Errorf("Error reading host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("log_blocked_traffic", flattenZtnaWebProxyLogBlockedTraffic(o["log-blocked-traffic"], d, "log_blocked_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-blocked-traffic"], "ZtnaWebProxy-LogBlockedTraffic"); ok {
			if err = d.Set("log_blocked_traffic", vv); err != nil {
				return fmt.Errorf("Error reading log_blocked_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_blocked_traffic: %v", err)
		}
	}

	if err = d.Set("name", flattenZtnaWebProxyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ZtnaWebProxy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("svr_pool_multiplex", flattenZtnaWebProxySvrPoolMultiplex(o["svr-pool-multiplex"], d, "svr_pool_multiplex")); err != nil {
		if vv, ok := fortiAPIPatch(o["svr-pool-multiplex"], "ZtnaWebProxy-SvrPoolMultiplex"); ok {
			if err = d.Set("svr_pool_multiplex", vv); err != nil {
				return fmt.Errorf("Error reading svr_pool_multiplex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading svr_pool_multiplex: %v", err)
		}
	}

	if err = d.Set("svr_pool_server_max_concurrent_request", flattenZtnaWebProxySvrPoolServerMaxConcurrentRequest(o["svr-pool-server-max-concurrent-request"], d, "svr_pool_server_max_concurrent_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["svr-pool-server-max-concurrent-request"], "ZtnaWebProxy-SvrPoolServerMaxConcurrentRequest"); ok {
			if err = d.Set("svr_pool_server_max_concurrent_request", vv); err != nil {
				return fmt.Errorf("Error reading svr_pool_server_max_concurrent_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading svr_pool_server_max_concurrent_request: %v", err)
		}
	}

	if err = d.Set("svr_pool_server_max_request", flattenZtnaWebProxySvrPoolServerMaxRequest(o["svr-pool-server-max-request"], d, "svr_pool_server_max_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["svr-pool-server-max-request"], "ZtnaWebProxy-SvrPoolServerMaxRequest"); ok {
			if err = d.Set("svr_pool_server_max_request", vv); err != nil {
				return fmt.Errorf("Error reading svr_pool_server_max_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading svr_pool_server_max_request: %v", err)
		}
	}

	if err = d.Set("svr_pool_ttl", flattenZtnaWebProxySvrPoolTtl(o["svr-pool-ttl"], d, "svr_pool_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["svr-pool-ttl"], "ZtnaWebProxy-SvrPoolTtl"); ok {
			if err = d.Set("svr_pool_ttl", vv); err != nil {
				return fmt.Errorf("Error reading svr_pool_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading svr_pool_ttl: %v", err)
		}
	}

	if err = d.Set("vip", flattenZtnaWebProxyVip(o["vip"], d, "vip")); err != nil {
		if vv, ok := fortiAPIPatch(o["vip"], "ZtnaWebProxy-Vip"); ok {
			if err = d.Set("vip", vv); err != nil {
				return fmt.Errorf("Error reading vip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vip: %v", err)
		}
	}

	if err = d.Set("vip6", flattenZtnaWebProxyVip6(o["vip6"], d, "vip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["vip6"], "ZtnaWebProxy-Vip6"); ok {
			if err = d.Set("vip6", vv); err != nil {
				return fmt.Errorf("Error reading vip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vip6: %v", err)
		}
	}

	return nil
}

func flattenZtnaWebProxyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandZtnaWebProxyApiGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h2_support"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["h2-support"], _ = expandZtnaWebProxyApiGatewayH2Support(d, i["h2_support"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h3_support"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["h3-support"], _ = expandZtnaWebProxyApiGatewayH3Support(d, i["h3_support"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-age"], _ = expandZtnaWebProxyApiGatewayHttpCookieAge(d, i["http_cookie_age"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-domain"], _ = expandZtnaWebProxyApiGatewayHttpCookieDomain(d, i["http_cookie_domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-domain-from-host"], _ = expandZtnaWebProxyApiGatewayHttpCookieDomainFromHost(d, i["http_cookie_domain_from_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-generation"], _ = expandZtnaWebProxyApiGatewayHttpCookieGeneration(d, i["http_cookie_generation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-path"], _ = expandZtnaWebProxyApiGatewayHttpCookiePath(d, i["http_cookie_path"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-share"], _ = expandZtnaWebProxyApiGatewayHttpCookieShare(d, i["http_cookie_share"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["https-cookie-secure"], _ = expandZtnaWebProxyApiGatewayHttpsCookieSecure(d, i["https_cookie_secure"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandZtnaWebProxyApiGatewayId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ldb-method"], _ = expandZtnaWebProxyApiGatewayLdbMethod(d, i["ldb_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["persistence"], _ = expandZtnaWebProxyApiGatewayPersistence(d, i["persistence"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quic"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandZtnaWebProxyApiGatewayQuic(d, i["quic"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["quic"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandZtnaWebProxyApiGatewayRealservers(d, i["realservers"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["realservers"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service"], _ = expandZtnaWebProxyApiGatewayService(d, i["service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-algorithm"], _ = expandZtnaWebProxyApiGatewaySslAlgorithm(d, i["ssl_algorithm"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandZtnaWebProxyApiGatewaySslCipherSuites(d, i["ssl_cipher_suites"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["ssl-cipher-suites"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-dh-bits"], _ = expandZtnaWebProxyApiGatewaySslDhBits(d, i["ssl_dh_bits"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-max-version"], _ = expandZtnaWebProxyApiGatewaySslMaxVersion(d, i["ssl_max_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-min-version"], _ = expandZtnaWebProxyApiGatewaySslMinVersion(d, i["ssl_min_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_renegotiation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-renegotiation"], _ = expandZtnaWebProxyApiGatewaySslRenegotiation(d, i["ssl_renegotiation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["url-map"], _ = expandZtnaWebProxyApiGatewayUrlMap(d, i["url_map"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["url-map-type"], _ = expandZtnaWebProxyApiGatewayUrlMapType(d, i["url_map_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandZtnaWebProxyApiGatewayH2Support(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayH3Support(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayHttpCookieAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayHttpCookieDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayHttpCookieDomainFromHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayHttpCookieGeneration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayHttpCookiePath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayHttpCookieShare(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayHttpsCookieSecure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayLdbMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayPersistence(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayQuic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ack-delay-exponent"], _ = expandZtnaWebProxyApiGatewayQuicAckDelayExponent(d, i["ack_delay_exponent"], pre_append)
	}
	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["active-connection-id-limit"], _ = expandZtnaWebProxyApiGatewayQuicActiveConnectionIdLimit(d, i["active_connection_id_limit"], pre_append)
	}
	pre_append = pre + ".0." + "active_migration"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["active-migration"], _ = expandZtnaWebProxyApiGatewayQuicActiveMigration(d, i["active_migration"], pre_append)
	}
	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["grease-quic-bit"], _ = expandZtnaWebProxyApiGatewayQuicGreaseQuicBit(d, i["grease_quic_bit"], pre_append)
	}
	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-ack-delay"], _ = expandZtnaWebProxyApiGatewayQuicMaxAckDelay(d, i["max_ack_delay"], pre_append)
	}
	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-datagram-frame-size"], _ = expandZtnaWebProxyApiGatewayQuicMaxDatagramFrameSize(d, i["max_datagram_frame_size"], pre_append)
	}
	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-idle-timeout"], _ = expandZtnaWebProxyApiGatewayQuicMaxIdleTimeout(d, i["max_idle_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-udp-payload-size"], _ = expandZtnaWebProxyApiGatewayQuicMaxUdpPayloadSize(d, i["max_udp_payload_size"], pre_append)
	}

	return result, nil
}

func expandZtnaWebProxyApiGatewayQuicAckDelayExponent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayQuicActiveConnectionIdLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayQuicActiveMigration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayQuicGreaseQuicBit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayQuicMaxAckDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayQuicMaxDatagramFrameSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayQuicMaxIdleTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayQuicMaxUdpPayloadSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealservers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["addr-type"], _ = expandZtnaWebProxyApiGatewayRealserversAddrType(d, i["addr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["address"], _ = expandZtnaWebProxyApiGatewayRealserversAddress(d, i["address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check"], _ = expandZtnaWebProxyApiGatewayRealserversHealthCheck(d, i["health_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check-proto"], _ = expandZtnaWebProxyApiGatewayRealserversHealthCheckProto(d, i["health_check_proto"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["holddown-interval"], _ = expandZtnaWebProxyApiGatewayRealserversHolddownInterval(d, i["holddown_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-host"], _ = expandZtnaWebProxyApiGatewayRealserversHttpHost(d, i["http_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandZtnaWebProxyApiGatewayRealserversId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandZtnaWebProxyApiGatewayRealserversIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandZtnaWebProxyApiGatewayRealserversPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandZtnaWebProxyApiGatewayRealserversStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["translate-host"], _ = expandZtnaWebProxyApiGatewayRealserversTranslateHost(d, i["translate_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandZtnaWebProxyApiGatewayRealserversWeight(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandZtnaWebProxyApiGatewayRealserversAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaWebProxyApiGatewayRealserversHealthCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversHealthCheckProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversHolddownInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversHttpHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversTranslateHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewaySslAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewaySslCipherSuites(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["cipher"], _ = expandZtnaWebProxyApiGatewaySslCipherSuitesCipher(d, i["cipher"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandZtnaWebProxyApiGatewaySslCipherSuitesPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["versions"], _ = expandZtnaWebProxyApiGatewaySslCipherSuitesVersions(d, i["versions"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandZtnaWebProxyApiGatewaySslCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewaySslCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewaySslCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaWebProxyApiGatewaySslDhBits(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewaySslMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewaySslMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewaySslRenegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayUrlMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayUrlMapType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h2_support"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["h2-support"], _ = expandZtnaWebProxyApiGateway6H2Support(d, i["h2_support"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h3_support"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["h3-support"], _ = expandZtnaWebProxyApiGateway6H3Support(d, i["h3_support"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-age"], _ = expandZtnaWebProxyApiGateway6HttpCookieAge(d, i["http_cookie_age"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-domain"], _ = expandZtnaWebProxyApiGateway6HttpCookieDomain(d, i["http_cookie_domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-domain-from-host"], _ = expandZtnaWebProxyApiGateway6HttpCookieDomainFromHost(d, i["http_cookie_domain_from_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-generation"], _ = expandZtnaWebProxyApiGateway6HttpCookieGeneration(d, i["http_cookie_generation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-path"], _ = expandZtnaWebProxyApiGateway6HttpCookiePath(d, i["http_cookie_path"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-share"], _ = expandZtnaWebProxyApiGateway6HttpCookieShare(d, i["http_cookie_share"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["https-cookie-secure"], _ = expandZtnaWebProxyApiGateway6HttpsCookieSecure(d, i["https_cookie_secure"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandZtnaWebProxyApiGateway6Id(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ldb-method"], _ = expandZtnaWebProxyApiGateway6LdbMethod(d, i["ldb_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["persistence"], _ = expandZtnaWebProxyApiGateway6Persistence(d, i["persistence"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quic"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandZtnaWebProxyApiGateway6Quic(d, i["quic"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["quic"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandZtnaWebProxyApiGateway6Realservers(d, i["realservers"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["realservers"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service"], _ = expandZtnaWebProxyApiGateway6Service(d, i["service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-algorithm"], _ = expandZtnaWebProxyApiGateway6SslAlgorithm(d, i["ssl_algorithm"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandZtnaWebProxyApiGateway6SslCipherSuites(d, i["ssl_cipher_suites"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["ssl-cipher-suites"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-dh-bits"], _ = expandZtnaWebProxyApiGateway6SslDhBits(d, i["ssl_dh_bits"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-max-version"], _ = expandZtnaWebProxyApiGateway6SslMaxVersion(d, i["ssl_max_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-min-version"], _ = expandZtnaWebProxyApiGateway6SslMinVersion(d, i["ssl_min_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_renegotiation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-renegotiation"], _ = expandZtnaWebProxyApiGateway6SslRenegotiation(d, i["ssl_renegotiation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["url-map"], _ = expandZtnaWebProxyApiGateway6UrlMap(d, i["url_map"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["url-map-type"], _ = expandZtnaWebProxyApiGateway6UrlMapType(d, i["url_map_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandZtnaWebProxyApiGateway6H2Support(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6H3Support(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6HttpCookieAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6HttpCookieDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6HttpCookieDomainFromHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6HttpCookieGeneration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6HttpCookiePath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6HttpCookieShare(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6HttpsCookieSecure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6LdbMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6Persistence(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6Quic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ack-delay-exponent"], _ = expandZtnaWebProxyApiGateway6QuicAckDelayExponent(d, i["ack_delay_exponent"], pre_append)
	}
	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["active-connection-id-limit"], _ = expandZtnaWebProxyApiGateway6QuicActiveConnectionIdLimit(d, i["active_connection_id_limit"], pre_append)
	}
	pre_append = pre + ".0." + "active_migration"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["active-migration"], _ = expandZtnaWebProxyApiGateway6QuicActiveMigration(d, i["active_migration"], pre_append)
	}
	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["grease-quic-bit"], _ = expandZtnaWebProxyApiGateway6QuicGreaseQuicBit(d, i["grease_quic_bit"], pre_append)
	}
	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-ack-delay"], _ = expandZtnaWebProxyApiGateway6QuicMaxAckDelay(d, i["max_ack_delay"], pre_append)
	}
	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-datagram-frame-size"], _ = expandZtnaWebProxyApiGateway6QuicMaxDatagramFrameSize(d, i["max_datagram_frame_size"], pre_append)
	}
	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-idle-timeout"], _ = expandZtnaWebProxyApiGateway6QuicMaxIdleTimeout(d, i["max_idle_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-udp-payload-size"], _ = expandZtnaWebProxyApiGateway6QuicMaxUdpPayloadSize(d, i["max_udp_payload_size"], pre_append)
	}

	return result, nil
}

func expandZtnaWebProxyApiGateway6QuicAckDelayExponent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6QuicActiveConnectionIdLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6QuicActiveMigration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6QuicGreaseQuicBit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6QuicMaxAckDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6QuicMaxDatagramFrameSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6QuicMaxIdleTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6QuicMaxUdpPayloadSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6Realservers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["addr-type"], _ = expandZtnaWebProxyApiGateway6RealserversAddrType(d, i["addr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["address"], _ = expandZtnaWebProxyApiGateway6RealserversAddress(d, i["address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check"], _ = expandZtnaWebProxyApiGateway6RealserversHealthCheck(d, i["health_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check-proto"], _ = expandZtnaWebProxyApiGateway6RealserversHealthCheckProto(d, i["health_check_proto"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["holddown-interval"], _ = expandZtnaWebProxyApiGateway6RealserversHolddownInterval(d, i["holddown_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-host"], _ = expandZtnaWebProxyApiGateway6RealserversHttpHost(d, i["http_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandZtnaWebProxyApiGateway6RealserversId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandZtnaWebProxyApiGateway6RealserversIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandZtnaWebProxyApiGateway6RealserversPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandZtnaWebProxyApiGateway6RealserversStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["translate-host"], _ = expandZtnaWebProxyApiGateway6RealserversTranslateHost(d, i["translate_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandZtnaWebProxyApiGateway6RealserversWeight(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandZtnaWebProxyApiGateway6RealserversAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaWebProxyApiGateway6RealserversHealthCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversHealthCheckProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversHolddownInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversHttpHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversTranslateHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6Service(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6SslAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6SslCipherSuites(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["cipher"], _ = expandZtnaWebProxyApiGateway6SslCipherSuitesCipher(d, i["cipher"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandZtnaWebProxyApiGateway6SslCipherSuitesPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["versions"], _ = expandZtnaWebProxyApiGateway6SslCipherSuitesVersions(d, i["versions"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandZtnaWebProxyApiGateway6SslCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6SslCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6SslCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaWebProxyApiGateway6SslDhBits(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6SslMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6SslMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6SslRenegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6UrlMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6UrlMapType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyAuthPortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyAuthVirtualHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaWebProxyDecryptedTrafficMirror(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaWebProxyHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaWebProxyLogBlockedTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxySvrPoolMultiplex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxySvrPoolServerMaxConcurrentRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxySvrPoolServerMaxRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxySvrPoolTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyVip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaWebProxyVip6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectZtnaWebProxy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("api_gateway"); ok || d.HasChange("api_gateway") {
		t, err := expandZtnaWebProxyApiGateway(d, v, "api_gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["api-gateway"] = t
		}
	}

	if v, ok := d.GetOk("api_gateway6"); ok || d.HasChange("api_gateway6") {
		t, err := expandZtnaWebProxyApiGateway6(d, v, "api_gateway6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["api-gateway6"] = t
		}
	}

	if v, ok := d.GetOk("auth_portal"); ok || d.HasChange("auth_portal") {
		t, err := expandZtnaWebProxyAuthPortal(d, v, "auth_portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-portal"] = t
		}
	}

	if v, ok := d.GetOk("auth_virtual_host"); ok || d.HasChange("auth_virtual_host") {
		t, err := expandZtnaWebProxyAuthVirtualHost(d, v, "auth_virtual_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-virtual-host"] = t
		}
	}

	if v, ok := d.GetOk("decrypted_traffic_mirror"); ok || d.HasChange("decrypted_traffic_mirror") {
		t, err := expandZtnaWebProxyDecryptedTrafficMirror(d, v, "decrypted_traffic_mirror")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["decrypted-traffic-mirror"] = t
		}
	}

	if v, ok := d.GetOk("host"); ok || d.HasChange("host") {
		t, err := expandZtnaWebProxyHost(d, v, "host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}

	if v, ok := d.GetOk("log_blocked_traffic"); ok || d.HasChange("log_blocked_traffic") {
		t, err := expandZtnaWebProxyLogBlockedTraffic(d, v, "log_blocked_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-blocked-traffic"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandZtnaWebProxyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("svr_pool_multiplex"); ok || d.HasChange("svr_pool_multiplex") {
		t, err := expandZtnaWebProxySvrPoolMultiplex(d, v, "svr_pool_multiplex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["svr-pool-multiplex"] = t
		}
	}

	if v, ok := d.GetOk("svr_pool_server_max_concurrent_request"); ok || d.HasChange("svr_pool_server_max_concurrent_request") {
		t, err := expandZtnaWebProxySvrPoolServerMaxConcurrentRequest(d, v, "svr_pool_server_max_concurrent_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["svr-pool-server-max-concurrent-request"] = t
		}
	}

	if v, ok := d.GetOk("svr_pool_server_max_request"); ok || d.HasChange("svr_pool_server_max_request") {
		t, err := expandZtnaWebProxySvrPoolServerMaxRequest(d, v, "svr_pool_server_max_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["svr-pool-server-max-request"] = t
		}
	}

	if v, ok := d.GetOk("svr_pool_ttl"); ok || d.HasChange("svr_pool_ttl") {
		t, err := expandZtnaWebProxySvrPoolTtl(d, v, "svr_pool_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["svr-pool-ttl"] = t
		}
	}

	if v, ok := d.GetOk("vip"); ok || d.HasChange("vip") {
		t, err := expandZtnaWebProxyVip(d, v, "vip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vip"] = t
		}
	}

	if v, ok := d.GetOk("vip6"); ok || d.HasChange("vip6") {
		t, err := expandZtnaWebProxyVip6(d, v, "vip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vip6"] = t
		}
	}

	return &obj, nil
}
