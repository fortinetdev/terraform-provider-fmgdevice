// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure IPv4 access proxy.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallAccessProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallAccessProxyCreate,
		Read:   resourceFirewallAccessProxyRead,
		Update: resourceFirewallAccessProxyUpdate,
		Delete: resourceFirewallAccessProxyDelete,

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
			"add_vhost_domain_to_dnsdb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"api_gateway": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"id": &schema.Schema{
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
					},
				},
			},
			"api_gateway6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"id": &schema.Schema{
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
			"client_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"decrypted_traffic_mirror": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"empty_cert_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldb_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_blocked_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"realservers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"server_pubkey_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_pubkey_auth_settings": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_ca": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"cert_extension": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"critical": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"data": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"permit_agent_forwarding": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"permit_port_forwarding": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"permit_pty": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"permit_user_rc": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"permit_x11_forwarding": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"source_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
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
			"vip": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"http_supported_max_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceFirewallAccessProxyCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallAccessProxy(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallAccessProxy resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallAccessProxy(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallAccessProxy(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallAccessProxy resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallAccessProxy(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallAccessProxy resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallAccessProxyRead(d, m)
}

func resourceFirewallAccessProxyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallAccessProxy(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAccessProxy resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallAccessProxy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAccessProxy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallAccessProxyRead(d, m)
}

func resourceFirewallAccessProxyDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallAccessProxy(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallAccessProxy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAccessProxyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallAccessProxy(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallAccessProxy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallAccessProxy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAccessProxy resource from API: %v", err)
	}
	return nil
}

func flattenFirewallAccessProxyAddVhostDomainToDnsdb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application"
		if _, ok := i["application"]; ok {
			v := flattenFirewallAccessProxyApiGatewayApplication(i["application"], d, pre_append)
			tmp["application"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway-Application")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h2_support"
		if _, ok := i["h2-support"]; ok {
			v := flattenFirewallAccessProxyApiGatewayH2Support(i["h2-support"], d, pre_append)
			tmp["h2_support"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway-H2Support")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h3_support"
		if _, ok := i["h3-support"]; ok {
			v := flattenFirewallAccessProxyApiGatewayH3Support(i["h3-support"], d, pre_append)
			tmp["h3_support"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway-H3Support")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := i["http-cookie-age"]; ok {
			v := flattenFirewallAccessProxyApiGatewayHttpCookieAge(i["http-cookie-age"], d, pre_append)
			tmp["http_cookie_age"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway-HttpCookieAge")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := i["http-cookie-domain"]; ok {
			v := flattenFirewallAccessProxyApiGatewayHttpCookieDomain(i["http-cookie-domain"], d, pre_append)
			tmp["http_cookie_domain"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway-HttpCookieDomain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := i["http-cookie-domain-from-host"]; ok {
			v := flattenFirewallAccessProxyApiGatewayHttpCookieDomainFromHost(i["http-cookie-domain-from-host"], d, pre_append)
			tmp["http_cookie_domain_from_host"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway-HttpCookieDomainFromHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := i["http-cookie-generation"]; ok {
			v := flattenFirewallAccessProxyApiGatewayHttpCookieGeneration(i["http-cookie-generation"], d, pre_append)
			tmp["http_cookie_generation"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway-HttpCookieGeneration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := i["http-cookie-path"]; ok {
			v := flattenFirewallAccessProxyApiGatewayHttpCookiePath(i["http-cookie-path"], d, pre_append)
			tmp["http_cookie_path"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway-HttpCookiePath")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := i["http-cookie-share"]; ok {
			v := flattenFirewallAccessProxyApiGatewayHttpCookieShare(i["http-cookie-share"], d, pre_append)
			tmp["http_cookie_share"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway-HttpCookieShare")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := i["https-cookie-secure"]; ok {
			v := flattenFirewallAccessProxyApiGatewayHttpsCookieSecure(i["https-cookie-secure"], d, pre_append)
			tmp["https_cookie_secure"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway-HttpsCookieSecure")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFirewallAccessProxyApiGatewayId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := i["ldb-method"]; ok {
			v := flattenFirewallAccessProxyApiGatewayLdbMethod(i["ldb-method"], d, pre_append)
			tmp["ldb_method"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway-LdbMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := i["persistence"]; ok {
			v := flattenFirewallAccessProxyApiGatewayPersistence(i["persistence"], d, pre_append)
			tmp["persistence"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway-Persistence")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quic"
		if _, ok := i["quic"]; ok {
			v := flattenFirewallAccessProxyApiGatewayQuic(i["quic"], d, pre_append)
			tmp["quic"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway-Quic")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := i["realservers"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealservers(i["realservers"], d, pre_append)
			tmp["realservers"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway-Realservers")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_redirect"
		if _, ok := i["saml-redirect"]; ok {
			v := flattenFirewallAccessProxyApiGatewaySamlRedirect(i["saml-redirect"], d, pre_append)
			tmp["saml_redirect"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway-SamlRedirect")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_server"
		if _, ok := i["saml-server"]; ok {
			v := flattenFirewallAccessProxyApiGatewaySamlServer(i["saml-server"], d, pre_append)
			tmp["saml_server"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway-SamlServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := i["service"]; ok {
			v := flattenFirewallAccessProxyApiGatewayService(i["service"], d, pre_append)
			tmp["service"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway-Service")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := i["ssl-algorithm"]; ok {
			v := flattenFirewallAccessProxyApiGatewaySslAlgorithm(i["ssl-algorithm"], d, pre_append)
			tmp["ssl_algorithm"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway-SslAlgorithm")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := i["ssl-cipher-suites"]; ok {
			v := flattenFirewallAccessProxyApiGatewaySslCipherSuites(i["ssl-cipher-suites"], d, pre_append)
			tmp["ssl_cipher_suites"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway-SslCipherSuites")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := i["ssl-dh-bits"]; ok {
			v := flattenFirewallAccessProxyApiGatewaySslDhBits(i["ssl-dh-bits"], d, pre_append)
			tmp["ssl_dh_bits"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway-SslDhBits")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := i["ssl-max-version"]; ok {
			v := flattenFirewallAccessProxyApiGatewaySslMaxVersion(i["ssl-max-version"], d, pre_append)
			tmp["ssl_max_version"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway-SslMaxVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := i["ssl-min-version"]; ok {
			v := flattenFirewallAccessProxyApiGatewaySslMinVersion(i["ssl-min-version"], d, pre_append)
			tmp["ssl_min_version"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway-SslMinVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_renegotiation"
		if _, ok := i["ssl-renegotiation"]; ok {
			v := flattenFirewallAccessProxyApiGatewaySslRenegotiation(i["ssl-renegotiation"], d, pre_append)
			tmp["ssl_renegotiation"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway-SslRenegotiation")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_vpn_web_portal"
		if _, ok := i["ssl-vpn-web-portal"]; ok {
			v := flattenFirewallAccessProxyApiGatewaySslVpnWebPortal(i["ssl-vpn-web-portal"], d, pre_append)
			tmp["ssl_vpn_web_portal"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway-SslVpnWebPortal")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map"
		if _, ok := i["url-map"]; ok {
			v := flattenFirewallAccessProxyApiGatewayUrlMap(i["url-map"], d, pre_append)
			tmp["url_map"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway-UrlMap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map_type"
		if _, ok := i["url-map-type"]; ok {
			v := flattenFirewallAccessProxyApiGatewayUrlMapType(i["url-map-type"], d, pre_append)
			tmp["url_map_type"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway-UrlMapType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_host"
		if _, ok := i["virtual-host"]; ok {
			v := flattenFirewallAccessProxyApiGatewayVirtualHost(i["virtual-host"], d, pre_append)
			tmp["virtual_host"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway-VirtualHost")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAccessProxyApiGatewayApplication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyApiGatewayH2Support(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayH3Support(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayHttpCookieAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayHttpCookieDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayHttpCookieDomainFromHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayHttpCookieGeneration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayHttpCookiePath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayHttpCookieShare(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayHttpsCookieSecure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayLdbMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayPersistence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayQuic(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := i["ack-delay-exponent"]; ok {
		result["ack_delay_exponent"] = flattenFirewallAccessProxyApiGatewayQuicAckDelayExponent(i["ack-delay-exponent"], d, pre_append)
	}

	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := i["active-connection-id-limit"]; ok {
		result["active_connection_id_limit"] = flattenFirewallAccessProxyApiGatewayQuicActiveConnectionIdLimit(i["active-connection-id-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "active_migration"
	if _, ok := i["active-migration"]; ok {
		result["active_migration"] = flattenFirewallAccessProxyApiGatewayQuicActiveMigration(i["active-migration"], d, pre_append)
	}

	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := i["grease-quic-bit"]; ok {
		result["grease_quic_bit"] = flattenFirewallAccessProxyApiGatewayQuicGreaseQuicBit(i["grease-quic-bit"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := i["max-ack-delay"]; ok {
		result["max_ack_delay"] = flattenFirewallAccessProxyApiGatewayQuicMaxAckDelay(i["max-ack-delay"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := i["max-datagram-frame-size"]; ok {
		result["max_datagram_frame_size"] = flattenFirewallAccessProxyApiGatewayQuicMaxDatagramFrameSize(i["max-datagram-frame-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := i["max-idle-timeout"]; ok {
		result["max_idle_timeout"] = flattenFirewallAccessProxyApiGatewayQuicMaxIdleTimeout(i["max-idle-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := i["max-udp-payload-size"]; ok {
		result["max_udp_payload_size"] = flattenFirewallAccessProxyApiGatewayQuicMaxUdpPayloadSize(i["max-udp-payload-size"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallAccessProxyApiGatewayQuicAckDelayExponent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayQuicActiveConnectionIdLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayQuicActiveMigration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayQuicGreaseQuicBit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayQuicMaxAckDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayQuicMaxDatagramFrameSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayQuicMaxIdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayQuicMaxUdpPayloadSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealservers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallAccessProxyApiGatewayRealserversAddrType(i["addr-type"], d, pre_append)
			tmp["addr_type"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-AddrType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversAddress(i["address"], d, pre_append)
			tmp["address"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := i["domain"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversDomain(i["domain"], d, pre_append)
			tmp["domain"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-Domain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_auth"
		if _, ok := i["external-auth"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversExternalAuth(i["external-auth"], d, pre_append)
			tmp["external_auth"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-ExternalAuth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := i["health-check"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversHealthCheck(i["health-check"], d, pre_append)
			tmp["health_check"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-HealthCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := i["health-check-proto"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversHealthCheckProto(i["health-check-proto"], d, pre_append)
			tmp["health_check_proto"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-HealthCheckProto")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := i["holddown-interval"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversHolddownInterval(i["holddown-interval"], d, pre_append)
			tmp["holddown_interval"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-HolddownInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := i["http-host"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversHttpHost(i["http-host"], d, pre_append)
			tmp["http_host"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-HttpHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := i["mappedport"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversMappedport(i["mappedport"], d, pre_append)
			tmp["mappedport"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-Mappedport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_client_cert"
		if _, ok := i["ssh-client-cert"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversSshClientCert(i["ssh-client-cert"], d, pre_append)
			tmp["ssh_client_cert"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-SshClientCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key"
		if _, ok := i["ssh-host-key"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversSshHostKey(i["ssh-host-key"], d, pre_append)
			tmp["ssh_host_key"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-SshHostKey")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key_validation"
		if _, ok := i["ssh-host-key-validation"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversSshHostKeyValidation(i["ssh-host-key-validation"], d, pre_append)
			tmp["ssh_host_key_validation"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-SshHostKeyValidation")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := i["translate-host"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversTranslateHost(i["translate-host"], d, pre_append)
			tmp["translate_host"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-TranslateHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tunnel_encryption"
		if _, ok := i["tunnel-encryption"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversTunnelEncryption(i["tunnel-encryption"], d, pre_append)
			tmp["tunnel_encryption"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-TunnelEncryption")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "verify_cert"
		if _, ok := i["verify-cert"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversVerifyCert(i["verify-cert"], d, pre_append)
			tmp["verify_cert"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-VerifyCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenFirewallAccessProxyApiGatewayRealserversWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-Realservers-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAccessProxyApiGatewayRealserversAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyApiGatewayRealserversDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversExternalAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversHealthCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversHealthCheckProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversHolddownInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversHttpHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversMappedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyApiGatewayRealserversPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversSshClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyApiGatewayRealserversSshHostKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyApiGatewayRealserversSshHostKeyValidation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversTranslateHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversTunnelEncryption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversVerifyCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySamlRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySamlServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyApiGatewayService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslCipherSuites(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallAccessProxyApiGatewaySslCipherSuitesCipher(i["cipher"], d, pre_append)
			tmp["cipher"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-SslCipherSuites-Cipher")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenFirewallAccessProxyApiGatewaySslCipherSuitesPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-SslCipherSuites-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {
			v := flattenFirewallAccessProxyApiGatewaySslCipherSuitesVersions(i["versions"], d, pre_append)
			tmp["versions"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway-SslCipherSuites-Versions")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAccessProxyApiGatewaySslCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyApiGatewaySslDhBits(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslVpnWebPortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyApiGatewayUrlMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayUrlMapType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayVirtualHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyApiGateway6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application"
		if _, ok := i["application"]; ok {
			v := flattenFirewallAccessProxyApiGateway6Application(i["application"], d, pre_append)
			tmp["application"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway6-Application")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h2_support"
		if _, ok := i["h2-support"]; ok {
			v := flattenFirewallAccessProxyApiGateway6H2Support(i["h2-support"], d, pre_append)
			tmp["h2_support"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway6-H2Support")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h3_support"
		if _, ok := i["h3-support"]; ok {
			v := flattenFirewallAccessProxyApiGateway6H3Support(i["h3-support"], d, pre_append)
			tmp["h3_support"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway6-H3Support")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := i["http-cookie-age"]; ok {
			v := flattenFirewallAccessProxyApiGateway6HttpCookieAge(i["http-cookie-age"], d, pre_append)
			tmp["http_cookie_age"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway6-HttpCookieAge")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := i["http-cookie-domain"]; ok {
			v := flattenFirewallAccessProxyApiGateway6HttpCookieDomain(i["http-cookie-domain"], d, pre_append)
			tmp["http_cookie_domain"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway6-HttpCookieDomain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := i["http-cookie-domain-from-host"]; ok {
			v := flattenFirewallAccessProxyApiGateway6HttpCookieDomainFromHost(i["http-cookie-domain-from-host"], d, pre_append)
			tmp["http_cookie_domain_from_host"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway6-HttpCookieDomainFromHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := i["http-cookie-generation"]; ok {
			v := flattenFirewallAccessProxyApiGateway6HttpCookieGeneration(i["http-cookie-generation"], d, pre_append)
			tmp["http_cookie_generation"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway6-HttpCookieGeneration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := i["http-cookie-path"]; ok {
			v := flattenFirewallAccessProxyApiGateway6HttpCookiePath(i["http-cookie-path"], d, pre_append)
			tmp["http_cookie_path"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway6-HttpCookiePath")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := i["http-cookie-share"]; ok {
			v := flattenFirewallAccessProxyApiGateway6HttpCookieShare(i["http-cookie-share"], d, pre_append)
			tmp["http_cookie_share"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway6-HttpCookieShare")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := i["https-cookie-secure"]; ok {
			v := flattenFirewallAccessProxyApiGateway6HttpsCookieSecure(i["https-cookie-secure"], d, pre_append)
			tmp["https_cookie_secure"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway6-HttpsCookieSecure")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFirewallAccessProxyApiGateway6Id(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway6-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := i["ldb-method"]; ok {
			v := flattenFirewallAccessProxyApiGateway6LdbMethod(i["ldb-method"], d, pre_append)
			tmp["ldb_method"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway6-LdbMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := i["persistence"]; ok {
			v := flattenFirewallAccessProxyApiGateway6Persistence(i["persistence"], d, pre_append)
			tmp["persistence"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway6-Persistence")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quic"
		if _, ok := i["quic"]; ok {
			v := flattenFirewallAccessProxyApiGateway6Quic(i["quic"], d, pre_append)
			tmp["quic"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway6-Quic")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := i["realservers"]; ok {
			v := flattenFirewallAccessProxyApiGateway6Realservers(i["realservers"], d, pre_append)
			tmp["realservers"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway6-Realservers")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_redirect"
		if _, ok := i["saml-redirect"]; ok {
			v := flattenFirewallAccessProxyApiGateway6SamlRedirect(i["saml-redirect"], d, pre_append)
			tmp["saml_redirect"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway6-SamlRedirect")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_server"
		if _, ok := i["saml-server"]; ok {
			v := flattenFirewallAccessProxyApiGateway6SamlServer(i["saml-server"], d, pre_append)
			tmp["saml_server"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway6-SamlServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := i["service"]; ok {
			v := flattenFirewallAccessProxyApiGateway6Service(i["service"], d, pre_append)
			tmp["service"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway6-Service")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := i["ssl-algorithm"]; ok {
			v := flattenFirewallAccessProxyApiGateway6SslAlgorithm(i["ssl-algorithm"], d, pre_append)
			tmp["ssl_algorithm"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway6-SslAlgorithm")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := i["ssl-cipher-suites"]; ok {
			v := flattenFirewallAccessProxyApiGateway6SslCipherSuites(i["ssl-cipher-suites"], d, pre_append)
			tmp["ssl_cipher_suites"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway6-SslCipherSuites")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := i["ssl-dh-bits"]; ok {
			v := flattenFirewallAccessProxyApiGateway6SslDhBits(i["ssl-dh-bits"], d, pre_append)
			tmp["ssl_dh_bits"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway6-SslDhBits")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := i["ssl-max-version"]; ok {
			v := flattenFirewallAccessProxyApiGateway6SslMaxVersion(i["ssl-max-version"], d, pre_append)
			tmp["ssl_max_version"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway6-SslMaxVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := i["ssl-min-version"]; ok {
			v := flattenFirewallAccessProxyApiGateway6SslMinVersion(i["ssl-min-version"], d, pre_append)
			tmp["ssl_min_version"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway6-SslMinVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_renegotiation"
		if _, ok := i["ssl-renegotiation"]; ok {
			v := flattenFirewallAccessProxyApiGateway6SslRenegotiation(i["ssl-renegotiation"], d, pre_append)
			tmp["ssl_renegotiation"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway6-SslRenegotiation")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_vpn_web_portal"
		if _, ok := i["ssl-vpn-web-portal"]; ok {
			v := flattenFirewallAccessProxyApiGateway6SslVpnWebPortal(i["ssl-vpn-web-portal"], d, pre_append)
			tmp["ssl_vpn_web_portal"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway6-SslVpnWebPortal")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map"
		if _, ok := i["url-map"]; ok {
			v := flattenFirewallAccessProxyApiGateway6UrlMap(i["url-map"], d, pre_append)
			tmp["url_map"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway6-UrlMap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map_type"
		if _, ok := i["url-map-type"]; ok {
			v := flattenFirewallAccessProxyApiGateway6UrlMapType(i["url-map-type"], d, pre_append)
			tmp["url_map_type"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway6-UrlMapType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_host"
		if _, ok := i["virtual-host"]; ok {
			v := flattenFirewallAccessProxyApiGateway6VirtualHost(i["virtual-host"], d, pre_append)
			tmp["virtual_host"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-ApiGateway6-VirtualHost")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAccessProxyApiGateway6Application(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyApiGateway6H2Support(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6H3Support(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6HttpCookieAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6HttpCookieDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6HttpCookieDomainFromHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6HttpCookieGeneration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6HttpCookiePath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6HttpCookieShare(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6HttpsCookieSecure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6LdbMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6Persistence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6Quic(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := i["ack-delay-exponent"]; ok {
		result["ack_delay_exponent"] = flattenFirewallAccessProxyApiGateway6QuicAckDelayExponent(i["ack-delay-exponent"], d, pre_append)
	}

	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := i["active-connection-id-limit"]; ok {
		result["active_connection_id_limit"] = flattenFirewallAccessProxyApiGateway6QuicActiveConnectionIdLimit(i["active-connection-id-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "active_migration"
	if _, ok := i["active-migration"]; ok {
		result["active_migration"] = flattenFirewallAccessProxyApiGateway6QuicActiveMigration(i["active-migration"], d, pre_append)
	}

	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := i["grease-quic-bit"]; ok {
		result["grease_quic_bit"] = flattenFirewallAccessProxyApiGateway6QuicGreaseQuicBit(i["grease-quic-bit"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := i["max-ack-delay"]; ok {
		result["max_ack_delay"] = flattenFirewallAccessProxyApiGateway6QuicMaxAckDelay(i["max-ack-delay"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := i["max-datagram-frame-size"]; ok {
		result["max_datagram_frame_size"] = flattenFirewallAccessProxyApiGateway6QuicMaxDatagramFrameSize(i["max-datagram-frame-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := i["max-idle-timeout"]; ok {
		result["max_idle_timeout"] = flattenFirewallAccessProxyApiGateway6QuicMaxIdleTimeout(i["max-idle-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := i["max-udp-payload-size"]; ok {
		result["max_udp_payload_size"] = flattenFirewallAccessProxyApiGateway6QuicMaxUdpPayloadSize(i["max-udp-payload-size"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallAccessProxyApiGateway6QuicAckDelayExponent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6QuicActiveConnectionIdLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6QuicActiveMigration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6QuicGreaseQuicBit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6QuicMaxAckDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6QuicMaxDatagramFrameSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6QuicMaxIdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6QuicMaxUdpPayloadSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6Realservers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallAccessProxyApiGateway6RealserversAddrType(i["addr-type"], d, pre_append)
			tmp["addr_type"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway6-Realservers-AddrType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {
			v := flattenFirewallAccessProxyApiGateway6RealserversAddress(i["address"], d, pre_append)
			tmp["address"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway6-Realservers-Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := i["domain"]; ok {
			v := flattenFirewallAccessProxyApiGateway6RealserversDomain(i["domain"], d, pre_append)
			tmp["domain"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway6-Realservers-Domain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_auth"
		if _, ok := i["external-auth"]; ok {
			v := flattenFirewallAccessProxyApiGateway6RealserversExternalAuth(i["external-auth"], d, pre_append)
			tmp["external_auth"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway6-Realservers-ExternalAuth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := i["health-check"]; ok {
			v := flattenFirewallAccessProxyApiGateway6RealserversHealthCheck(i["health-check"], d, pre_append)
			tmp["health_check"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway6-Realservers-HealthCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := i["health-check-proto"]; ok {
			v := flattenFirewallAccessProxyApiGateway6RealserversHealthCheckProto(i["health-check-proto"], d, pre_append)
			tmp["health_check_proto"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway6-Realservers-HealthCheckProto")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := i["holddown-interval"]; ok {
			v := flattenFirewallAccessProxyApiGateway6RealserversHolddownInterval(i["holddown-interval"], d, pre_append)
			tmp["holddown_interval"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway6-Realservers-HolddownInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := i["http-host"]; ok {
			v := flattenFirewallAccessProxyApiGateway6RealserversHttpHost(i["http-host"], d, pre_append)
			tmp["http_host"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway6-Realservers-HttpHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFirewallAccessProxyApiGateway6RealserversId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway6-Realservers-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenFirewallAccessProxyApiGateway6RealserversIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway6-Realservers-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := i["mappedport"]; ok {
			v := flattenFirewallAccessProxyApiGateway6RealserversMappedport(i["mappedport"], d, pre_append)
			tmp["mappedport"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway6-Realservers-Mappedport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenFirewallAccessProxyApiGateway6RealserversPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway6-Realservers-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_client_cert"
		if _, ok := i["ssh-client-cert"]; ok {
			v := flattenFirewallAccessProxyApiGateway6RealserversSshClientCert(i["ssh-client-cert"], d, pre_append)
			tmp["ssh_client_cert"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway6-Realservers-SshClientCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key"
		if _, ok := i["ssh-host-key"]; ok {
			v := flattenFirewallAccessProxyApiGateway6RealserversSshHostKey(i["ssh-host-key"], d, pre_append)
			tmp["ssh_host_key"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway6-Realservers-SshHostKey")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key_validation"
		if _, ok := i["ssh-host-key-validation"]; ok {
			v := flattenFirewallAccessProxyApiGateway6RealserversSshHostKeyValidation(i["ssh-host-key-validation"], d, pre_append)
			tmp["ssh_host_key_validation"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway6-Realservers-SshHostKeyValidation")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenFirewallAccessProxyApiGateway6RealserversStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway6-Realservers-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := i["translate-host"]; ok {
			v := flattenFirewallAccessProxyApiGateway6RealserversTranslateHost(i["translate-host"], d, pre_append)
			tmp["translate_host"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway6-Realservers-TranslateHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tunnel_encryption"
		if _, ok := i["tunnel-encryption"]; ok {
			v := flattenFirewallAccessProxyApiGateway6RealserversTunnelEncryption(i["tunnel-encryption"], d, pre_append)
			tmp["tunnel_encryption"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway6-Realservers-TunnelEncryption")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenFirewallAccessProxyApiGateway6RealserversType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway6-Realservers-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "verify_cert"
		if _, ok := i["verify-cert"]; ok {
			v := flattenFirewallAccessProxyApiGateway6RealserversVerifyCert(i["verify-cert"], d, pre_append)
			tmp["verify_cert"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway6-Realservers-VerifyCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenFirewallAccessProxyApiGateway6RealserversWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway6-Realservers-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAccessProxyApiGateway6RealserversAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyApiGateway6RealserversDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversExternalAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversHealthCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversHealthCheckProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversHolddownInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversHttpHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversMappedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyApiGateway6RealserversPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversSshClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyApiGateway6RealserversSshHostKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyApiGateway6RealserversSshHostKeyValidation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversTranslateHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversTunnelEncryption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversVerifyCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6SamlRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6SamlServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyApiGateway6Service(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6SslAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6SslCipherSuites(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallAccessProxyApiGateway6SslCipherSuitesCipher(i["cipher"], d, pre_append)
			tmp["cipher"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway6-SslCipherSuites-Cipher")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenFirewallAccessProxyApiGateway6SslCipherSuitesPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway6-SslCipherSuites-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {
			v := flattenFirewallAccessProxyApiGateway6SslCipherSuitesVersions(i["versions"], d, pre_append)
			tmp["versions"] = fortiAPISubPartPatch(v, "FirewallAccessProxyApiGateway6-SslCipherSuites-Versions")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAccessProxyApiGateway6SslCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6SslCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6SslCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyApiGateway6SslDhBits(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6SslMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6SslMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6SslRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6SslVpnWebPortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyApiGateway6UrlMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6UrlMapType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6VirtualHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyAuthPortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyAuthVirtualHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyDecryptedTrafficMirror(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyEmptyCertAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyLdbMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyLogBlockedTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyRealservers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFirewallAccessProxyRealserversId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-Realservers-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenFirewallAccessProxyRealserversIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-Realservers-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenFirewallAccessProxyRealserversPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-Realservers-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenFirewallAccessProxyRealserversStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-Realservers-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenFirewallAccessProxyRealserversWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "FirewallAccessProxy-Realservers-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAccessProxyRealserversId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyRealserversIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyRealserversPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyRealserversStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyRealserversWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyServerPubkeyAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyServerPubkeyAuthSettings(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auth_ca"
	if _, ok := i["auth-ca"]; ok {
		result["auth_ca"] = flattenFirewallAccessProxyServerPubkeyAuthSettingsAuthCa(i["auth-ca"], d, pre_append)
	}

	pre_append = pre + ".0." + "cert_extension"
	if _, ok := i["cert-extension"]; ok {
		result["cert_extension"] = flattenFirewallAccessProxyServerPubkeyAuthSettingsCertExtension(i["cert-extension"], d, pre_append)
	}

	pre_append = pre + ".0." + "permit_agent_forwarding"
	if _, ok := i["permit-agent-forwarding"]; ok {
		result["permit_agent_forwarding"] = flattenFirewallAccessProxyServerPubkeyAuthSettingsPermitAgentForwarding(i["permit-agent-forwarding"], d, pre_append)
	}

	pre_append = pre + ".0." + "permit_port_forwarding"
	if _, ok := i["permit-port-forwarding"]; ok {
		result["permit_port_forwarding"] = flattenFirewallAccessProxyServerPubkeyAuthSettingsPermitPortForwarding(i["permit-port-forwarding"], d, pre_append)
	}

	pre_append = pre + ".0." + "permit_pty"
	if _, ok := i["permit-pty"]; ok {
		result["permit_pty"] = flattenFirewallAccessProxyServerPubkeyAuthSettingsPermitPty(i["permit-pty"], d, pre_append)
	}

	pre_append = pre + ".0." + "permit_user_rc"
	if _, ok := i["permit-user-rc"]; ok {
		result["permit_user_rc"] = flattenFirewallAccessProxyServerPubkeyAuthSettingsPermitUserRc(i["permit-user-rc"], d, pre_append)
	}

	pre_append = pre + ".0." + "permit_x11_forwarding"
	if _, ok := i["permit-x11-forwarding"]; ok {
		result["permit_x11_forwarding"] = flattenFirewallAccessProxyServerPubkeyAuthSettingsPermitX11Forwarding(i["permit-x11-forwarding"], d, pre_append)
	}

	pre_append = pre + ".0." + "source_address"
	if _, ok := i["source-address"]; ok {
		result["source_address"] = flattenFirewallAccessProxyServerPubkeyAuthSettingsSourceAddress(i["source-address"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallAccessProxyServerPubkeyAuthSettingsAuthCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyServerPubkeyAuthSettingsCertExtension(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "critical"
		if _, ok := i["critical"]; ok {
			v := flattenFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionCritical(i["critical"], d, pre_append)
			tmp["critical"] = fortiAPISubPartPatch(v, "FirewallAccessProxyServerPubkeyAuthSettings-CertExtension-Critical")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "data"
		if _, ok := i["data"]; ok {
			v := flattenFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionData(i["data"], d, pre_append)
			tmp["data"] = fortiAPISubPartPatch(v, "FirewallAccessProxyServerPubkeyAuthSettings-CertExtension-Data")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "FirewallAccessProxyServerPubkeyAuthSettings-CertExtension-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "FirewallAccessProxyServerPubkeyAuthSettings-CertExtension-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionCritical(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionData(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyServerPubkeyAuthSettingsPermitAgentForwarding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyServerPubkeyAuthSettingsPermitPortForwarding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyServerPubkeyAuthSettingsPermitPty(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyServerPubkeyAuthSettingsPermitUserRc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyServerPubkeyAuthSettingsPermitX11Forwarding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyServerPubkeyAuthSettingsSourceAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxySvrPoolMultiplex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxySvrPoolServerMaxConcurrentRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxySvrPoolServerMaxRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxySvrPoolTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyUserAgentDetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyVip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyHttpSupportedMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallAccessProxy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("add_vhost_domain_to_dnsdb", flattenFirewallAccessProxyAddVhostDomainToDnsdb(o["add-vhost-domain-to-dnsdb"], d, "add_vhost_domain_to_dnsdb")); err != nil {
		if vv, ok := fortiAPIPatch(o["add-vhost-domain-to-dnsdb"], "FirewallAccessProxy-AddVhostDomainToDnsdb"); ok {
			if err = d.Set("add_vhost_domain_to_dnsdb", vv); err != nil {
				return fmt.Errorf("Error reading add_vhost_domain_to_dnsdb: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading add_vhost_domain_to_dnsdb: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("api_gateway", flattenFirewallAccessProxyApiGateway(o["api-gateway"], d, "api_gateway")); err != nil {
			if vv, ok := fortiAPIPatch(o["api-gateway"], "FirewallAccessProxy-ApiGateway"); ok {
				if err = d.Set("api_gateway", vv); err != nil {
					return fmt.Errorf("Error reading api_gateway: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading api_gateway: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("api_gateway"); ok {
			if err = d.Set("api_gateway", flattenFirewallAccessProxyApiGateway(o["api-gateway"], d, "api_gateway")); err != nil {
				if vv, ok := fortiAPIPatch(o["api-gateway"], "FirewallAccessProxy-ApiGateway"); ok {
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
		if err = d.Set("api_gateway6", flattenFirewallAccessProxyApiGateway6(o["api-gateway6"], d, "api_gateway6")); err != nil {
			if vv, ok := fortiAPIPatch(o["api-gateway6"], "FirewallAccessProxy-ApiGateway6"); ok {
				if err = d.Set("api_gateway6", vv); err != nil {
					return fmt.Errorf("Error reading api_gateway6: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading api_gateway6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("api_gateway6"); ok {
			if err = d.Set("api_gateway6", flattenFirewallAccessProxyApiGateway6(o["api-gateway6"], d, "api_gateway6")); err != nil {
				if vv, ok := fortiAPIPatch(o["api-gateway6"], "FirewallAccessProxy-ApiGateway6"); ok {
					if err = d.Set("api_gateway6", vv); err != nil {
						return fmt.Errorf("Error reading api_gateway6: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading api_gateway6: %v", err)
				}
			}
		}
	}

	if err = d.Set("auth_portal", flattenFirewallAccessProxyAuthPortal(o["auth-portal"], d, "auth_portal")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-portal"], "FirewallAccessProxy-AuthPortal"); ok {
			if err = d.Set("auth_portal", vv); err != nil {
				return fmt.Errorf("Error reading auth_portal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_portal: %v", err)
		}
	}

	if err = d.Set("auth_virtual_host", flattenFirewallAccessProxyAuthVirtualHost(o["auth-virtual-host"], d, "auth_virtual_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-virtual-host"], "FirewallAccessProxy-AuthVirtualHost"); ok {
			if err = d.Set("auth_virtual_host", vv); err != nil {
				return fmt.Errorf("Error reading auth_virtual_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_virtual_host: %v", err)
		}
	}

	if err = d.Set("client_cert", flattenFirewallAccessProxyClientCert(o["client-cert"], d, "client_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-cert"], "FirewallAccessProxy-ClientCert"); ok {
			if err = d.Set("client_cert", vv); err != nil {
				return fmt.Errorf("Error reading client_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_cert: %v", err)
		}
	}

	if err = d.Set("decrypted_traffic_mirror", flattenFirewallAccessProxyDecryptedTrafficMirror(o["decrypted-traffic-mirror"], d, "decrypted_traffic_mirror")); err != nil {
		if vv, ok := fortiAPIPatch(o["decrypted-traffic-mirror"], "FirewallAccessProxy-DecryptedTrafficMirror"); ok {
			if err = d.Set("decrypted_traffic_mirror", vv); err != nil {
				return fmt.Errorf("Error reading decrypted_traffic_mirror: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading decrypted_traffic_mirror: %v", err)
		}
	}

	if err = d.Set("empty_cert_action", flattenFirewallAccessProxyEmptyCertAction(o["empty-cert-action"], d, "empty_cert_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["empty-cert-action"], "FirewallAccessProxy-EmptyCertAction"); ok {
			if err = d.Set("empty_cert_action", vv); err != nil {
				return fmt.Errorf("Error reading empty_cert_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading empty_cert_action: %v", err)
		}
	}

	if err = d.Set("ldb_method", flattenFirewallAccessProxyLdbMethod(o["ldb-method"], d, "ldb_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldb-method"], "FirewallAccessProxy-LdbMethod"); ok {
			if err = d.Set("ldb_method", vv); err != nil {
				return fmt.Errorf("Error reading ldb_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldb_method: %v", err)
		}
	}

	if err = d.Set("log_blocked_traffic", flattenFirewallAccessProxyLogBlockedTraffic(o["log-blocked-traffic"], d, "log_blocked_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-blocked-traffic"], "FirewallAccessProxy-LogBlockedTraffic"); ok {
			if err = d.Set("log_blocked_traffic", vv); err != nil {
				return fmt.Errorf("Error reading log_blocked_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_blocked_traffic: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallAccessProxyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallAccessProxy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("realservers", flattenFirewallAccessProxyRealservers(o["realservers"], d, "realservers")); err != nil {
			if vv, ok := fortiAPIPatch(o["realservers"], "FirewallAccessProxy-Realservers"); ok {
				if err = d.Set("realservers", vv); err != nil {
					return fmt.Errorf("Error reading realservers: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading realservers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("realservers"); ok {
			if err = d.Set("realservers", flattenFirewallAccessProxyRealservers(o["realservers"], d, "realservers")); err != nil {
				if vv, ok := fortiAPIPatch(o["realservers"], "FirewallAccessProxy-Realservers"); ok {
					if err = d.Set("realservers", vv); err != nil {
						return fmt.Errorf("Error reading realservers: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading realservers: %v", err)
				}
			}
		}
	}

	if err = d.Set("server_pubkey_auth", flattenFirewallAccessProxyServerPubkeyAuth(o["server-pubkey-auth"], d, "server_pubkey_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-pubkey-auth"], "FirewallAccessProxy-ServerPubkeyAuth"); ok {
			if err = d.Set("server_pubkey_auth", vv); err != nil {
				return fmt.Errorf("Error reading server_pubkey_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_pubkey_auth: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("server_pubkey_auth_settings", flattenFirewallAccessProxyServerPubkeyAuthSettings(o["server-pubkey-auth-settings"], d, "server_pubkey_auth_settings")); err != nil {
			if vv, ok := fortiAPIPatch(o["server-pubkey-auth-settings"], "FirewallAccessProxy-ServerPubkeyAuthSettings"); ok {
				if err = d.Set("server_pubkey_auth_settings", vv); err != nil {
					return fmt.Errorf("Error reading server_pubkey_auth_settings: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading server_pubkey_auth_settings: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("server_pubkey_auth_settings"); ok {
			if err = d.Set("server_pubkey_auth_settings", flattenFirewallAccessProxyServerPubkeyAuthSettings(o["server-pubkey-auth-settings"], d, "server_pubkey_auth_settings")); err != nil {
				if vv, ok := fortiAPIPatch(o["server-pubkey-auth-settings"], "FirewallAccessProxy-ServerPubkeyAuthSettings"); ok {
					if err = d.Set("server_pubkey_auth_settings", vv); err != nil {
						return fmt.Errorf("Error reading server_pubkey_auth_settings: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading server_pubkey_auth_settings: %v", err)
				}
			}
		}
	}

	if err = d.Set("svr_pool_multiplex", flattenFirewallAccessProxySvrPoolMultiplex(o["svr-pool-multiplex"], d, "svr_pool_multiplex")); err != nil {
		if vv, ok := fortiAPIPatch(o["svr-pool-multiplex"], "FirewallAccessProxy-SvrPoolMultiplex"); ok {
			if err = d.Set("svr_pool_multiplex", vv); err != nil {
				return fmt.Errorf("Error reading svr_pool_multiplex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading svr_pool_multiplex: %v", err)
		}
	}

	if err = d.Set("svr_pool_server_max_concurrent_request", flattenFirewallAccessProxySvrPoolServerMaxConcurrentRequest(o["svr-pool-server-max-concurrent-request"], d, "svr_pool_server_max_concurrent_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["svr-pool-server-max-concurrent-request"], "FirewallAccessProxy-SvrPoolServerMaxConcurrentRequest"); ok {
			if err = d.Set("svr_pool_server_max_concurrent_request", vv); err != nil {
				return fmt.Errorf("Error reading svr_pool_server_max_concurrent_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading svr_pool_server_max_concurrent_request: %v", err)
		}
	}

	if err = d.Set("svr_pool_server_max_request", flattenFirewallAccessProxySvrPoolServerMaxRequest(o["svr-pool-server-max-request"], d, "svr_pool_server_max_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["svr-pool-server-max-request"], "FirewallAccessProxy-SvrPoolServerMaxRequest"); ok {
			if err = d.Set("svr_pool_server_max_request", vv); err != nil {
				return fmt.Errorf("Error reading svr_pool_server_max_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading svr_pool_server_max_request: %v", err)
		}
	}

	if err = d.Set("svr_pool_ttl", flattenFirewallAccessProxySvrPoolTtl(o["svr-pool-ttl"], d, "svr_pool_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["svr-pool-ttl"], "FirewallAccessProxy-SvrPoolTtl"); ok {
			if err = d.Set("svr_pool_ttl", vv); err != nil {
				return fmt.Errorf("Error reading svr_pool_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading svr_pool_ttl: %v", err)
		}
	}

	if err = d.Set("user_agent_detect", flattenFirewallAccessProxyUserAgentDetect(o["user-agent-detect"], d, "user_agent_detect")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-agent-detect"], "FirewallAccessProxy-UserAgentDetect"); ok {
			if err = d.Set("user_agent_detect", vv); err != nil {
				return fmt.Errorf("Error reading user_agent_detect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_agent_detect: %v", err)
		}
	}

	if err = d.Set("vip", flattenFirewallAccessProxyVip(o["vip"], d, "vip")); err != nil {
		if vv, ok := fortiAPIPatch(o["vip"], "FirewallAccessProxy-Vip"); ok {
			if err = d.Set("vip", vv); err != nil {
				return fmt.Errorf("Error reading vip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vip: %v", err)
		}
	}

	if err = d.Set("http_supported_max_version", flattenFirewallAccessProxyHttpSupportedMaxVersion(o["http-supported-max-version"], d, "http_supported_max_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-supported-max-version"], "FirewallAccessProxy-HttpSupportedMaxVersion"); ok {
			if err = d.Set("http_supported_max_version", vv); err != nil {
				return fmt.Errorf("Error reading http_supported_max_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_supported_max_version: %v", err)
		}
	}

	return nil
}

func flattenFirewallAccessProxyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallAccessProxyAddVhostDomainToDnsdb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["application"], _ = expandFirewallAccessProxyApiGatewayApplication(d, i["application"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h2_support"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["h2-support"], _ = expandFirewallAccessProxyApiGatewayH2Support(d, i["h2_support"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h3_support"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["h3-support"], _ = expandFirewallAccessProxyApiGatewayH3Support(d, i["h3_support"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-age"], _ = expandFirewallAccessProxyApiGatewayHttpCookieAge(d, i["http_cookie_age"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-domain"], _ = expandFirewallAccessProxyApiGatewayHttpCookieDomain(d, i["http_cookie_domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-domain-from-host"], _ = expandFirewallAccessProxyApiGatewayHttpCookieDomainFromHost(d, i["http_cookie_domain_from_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-generation"], _ = expandFirewallAccessProxyApiGatewayHttpCookieGeneration(d, i["http_cookie_generation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-path"], _ = expandFirewallAccessProxyApiGatewayHttpCookiePath(d, i["http_cookie_path"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-share"], _ = expandFirewallAccessProxyApiGatewayHttpCookieShare(d, i["http_cookie_share"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["https-cookie-secure"], _ = expandFirewallAccessProxyApiGatewayHttpsCookieSecure(d, i["https_cookie_secure"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFirewallAccessProxyApiGatewayId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ldb-method"], _ = expandFirewallAccessProxyApiGatewayLdbMethod(d, i["ldb_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["persistence"], _ = expandFirewallAccessProxyApiGatewayPersistence(d, i["persistence"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quic"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandFirewallAccessProxyApiGatewayQuic(d, i["quic"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["quic"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandFirewallAccessProxyApiGatewayRealservers(d, i["realservers"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["realservers"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_redirect"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["saml-redirect"], _ = expandFirewallAccessProxyApiGatewaySamlRedirect(d, i["saml_redirect"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["saml-server"], _ = expandFirewallAccessProxyApiGatewaySamlServer(d, i["saml_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service"], _ = expandFirewallAccessProxyApiGatewayService(d, i["service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-algorithm"], _ = expandFirewallAccessProxyApiGatewaySslAlgorithm(d, i["ssl_algorithm"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandFirewallAccessProxyApiGatewaySslCipherSuites(d, i["ssl_cipher_suites"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["ssl-cipher-suites"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-dh-bits"], _ = expandFirewallAccessProxyApiGatewaySslDhBits(d, i["ssl_dh_bits"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-max-version"], _ = expandFirewallAccessProxyApiGatewaySslMaxVersion(d, i["ssl_max_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-min-version"], _ = expandFirewallAccessProxyApiGatewaySslMinVersion(d, i["ssl_min_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_renegotiation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-renegotiation"], _ = expandFirewallAccessProxyApiGatewaySslRenegotiation(d, i["ssl_renegotiation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_vpn_web_portal"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-vpn-web-portal"], _ = expandFirewallAccessProxyApiGatewaySslVpnWebPortal(d, i["ssl_vpn_web_portal"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["url-map"], _ = expandFirewallAccessProxyApiGatewayUrlMap(d, i["url_map"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["url-map-type"], _ = expandFirewallAccessProxyApiGatewayUrlMapType(d, i["url_map_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["virtual-host"], _ = expandFirewallAccessProxyApiGatewayVirtualHost(d, i["virtual_host"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxyApiGatewayApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyApiGatewayH2Support(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayH3Support(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayHttpCookieAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayHttpCookieDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayHttpCookieDomainFromHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayHttpCookieGeneration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayHttpCookiePath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayHttpCookieShare(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayHttpsCookieSecure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayLdbMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayPersistence(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayQuic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ack-delay-exponent"], _ = expandFirewallAccessProxyApiGatewayQuicAckDelayExponent(d, i["ack_delay_exponent"], pre_append)
	}
	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["active-connection-id-limit"], _ = expandFirewallAccessProxyApiGatewayQuicActiveConnectionIdLimit(d, i["active_connection_id_limit"], pre_append)
	}
	pre_append = pre + ".0." + "active_migration"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["active-migration"], _ = expandFirewallAccessProxyApiGatewayQuicActiveMigration(d, i["active_migration"], pre_append)
	}
	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["grease-quic-bit"], _ = expandFirewallAccessProxyApiGatewayQuicGreaseQuicBit(d, i["grease_quic_bit"], pre_append)
	}
	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-ack-delay"], _ = expandFirewallAccessProxyApiGatewayQuicMaxAckDelay(d, i["max_ack_delay"], pre_append)
	}
	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-datagram-frame-size"], _ = expandFirewallAccessProxyApiGatewayQuicMaxDatagramFrameSize(d, i["max_datagram_frame_size"], pre_append)
	}
	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-idle-timeout"], _ = expandFirewallAccessProxyApiGatewayQuicMaxIdleTimeout(d, i["max_idle_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-udp-payload-size"], _ = expandFirewallAccessProxyApiGatewayQuicMaxUdpPayloadSize(d, i["max_udp_payload_size"], pre_append)
	}

	return result, nil
}

func expandFirewallAccessProxyApiGatewayQuicAckDelayExponent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayQuicActiveConnectionIdLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayQuicActiveMigration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayQuicGreaseQuicBit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayQuicMaxAckDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayQuicMaxDatagramFrameSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayQuicMaxIdleTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayQuicMaxUdpPayloadSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealservers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["addr-type"], _ = expandFirewallAccessProxyApiGatewayRealserversAddrType(d, i["addr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["address"], _ = expandFirewallAccessProxyApiGatewayRealserversAddress(d, i["address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["domain"], _ = expandFirewallAccessProxyApiGatewayRealserversDomain(d, i["domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_auth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["external-auth"], _ = expandFirewallAccessProxyApiGatewayRealserversExternalAuth(d, i["external_auth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check"], _ = expandFirewallAccessProxyApiGatewayRealserversHealthCheck(d, i["health_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check-proto"], _ = expandFirewallAccessProxyApiGatewayRealserversHealthCheckProto(d, i["health_check_proto"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["holddown-interval"], _ = expandFirewallAccessProxyApiGatewayRealserversHolddownInterval(d, i["holddown_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-host"], _ = expandFirewallAccessProxyApiGatewayRealserversHttpHost(d, i["http_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFirewallAccessProxyApiGatewayRealserversId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandFirewallAccessProxyApiGatewayRealserversIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mappedport"], _ = expandFirewallAccessProxyApiGatewayRealserversMappedport(d, i["mappedport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandFirewallAccessProxyApiGatewayRealserversPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_client_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssh-client-cert"], _ = expandFirewallAccessProxyApiGatewayRealserversSshClientCert(d, i["ssh_client_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssh-host-key"], _ = expandFirewallAccessProxyApiGatewayRealserversSshHostKey(d, i["ssh_host_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key_validation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssh-host-key-validation"], _ = expandFirewallAccessProxyApiGatewayRealserversSshHostKeyValidation(d, i["ssh_host_key_validation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandFirewallAccessProxyApiGatewayRealserversStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["translate-host"], _ = expandFirewallAccessProxyApiGatewayRealserversTranslateHost(d, i["translate_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tunnel_encryption"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tunnel-encryption"], _ = expandFirewallAccessProxyApiGatewayRealserversTunnelEncryption(d, i["tunnel_encryption"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandFirewallAccessProxyApiGatewayRealserversType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "verify_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["verify-cert"], _ = expandFirewallAccessProxyApiGatewayRealserversVerifyCert(d, i["verify_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandFirewallAccessProxyApiGatewayRealserversWeight(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxyApiGatewayRealserversAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyApiGatewayRealserversDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversExternalAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversHealthCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversHealthCheckProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversHolddownInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversHttpHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversMappedport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyApiGatewayRealserversPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversSshClientCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyApiGatewayRealserversSshHostKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyApiGatewayRealserversSshHostKeyValidation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversTranslateHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversTunnelEncryption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversVerifyCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySamlRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySamlServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyApiGatewayService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslCipherSuites(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["cipher"], _ = expandFirewallAccessProxyApiGatewaySslCipherSuitesCipher(d, i["cipher"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandFirewallAccessProxyApiGatewaySslCipherSuitesPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["versions"], _ = expandFirewallAccessProxyApiGatewaySslCipherSuitesVersions(d, i["versions"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxyApiGatewaySslCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyApiGatewaySslDhBits(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslRenegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslVpnWebPortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyApiGatewayUrlMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayUrlMapType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayVirtualHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyApiGateway6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["application"], _ = expandFirewallAccessProxyApiGateway6Application(d, i["application"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h2_support"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["h2-support"], _ = expandFirewallAccessProxyApiGateway6H2Support(d, i["h2_support"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h3_support"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["h3-support"], _ = expandFirewallAccessProxyApiGateway6H3Support(d, i["h3_support"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-age"], _ = expandFirewallAccessProxyApiGateway6HttpCookieAge(d, i["http_cookie_age"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-domain"], _ = expandFirewallAccessProxyApiGateway6HttpCookieDomain(d, i["http_cookie_domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-domain-from-host"], _ = expandFirewallAccessProxyApiGateway6HttpCookieDomainFromHost(d, i["http_cookie_domain_from_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-generation"], _ = expandFirewallAccessProxyApiGateway6HttpCookieGeneration(d, i["http_cookie_generation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-path"], _ = expandFirewallAccessProxyApiGateway6HttpCookiePath(d, i["http_cookie_path"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-share"], _ = expandFirewallAccessProxyApiGateway6HttpCookieShare(d, i["http_cookie_share"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["https-cookie-secure"], _ = expandFirewallAccessProxyApiGateway6HttpsCookieSecure(d, i["https_cookie_secure"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFirewallAccessProxyApiGateway6Id(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ldb-method"], _ = expandFirewallAccessProxyApiGateway6LdbMethod(d, i["ldb_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["persistence"], _ = expandFirewallAccessProxyApiGateway6Persistence(d, i["persistence"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quic"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandFirewallAccessProxyApiGateway6Quic(d, i["quic"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["quic"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandFirewallAccessProxyApiGateway6Realservers(d, i["realservers"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["realservers"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_redirect"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["saml-redirect"], _ = expandFirewallAccessProxyApiGateway6SamlRedirect(d, i["saml_redirect"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["saml-server"], _ = expandFirewallAccessProxyApiGateway6SamlServer(d, i["saml_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service"], _ = expandFirewallAccessProxyApiGateway6Service(d, i["service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-algorithm"], _ = expandFirewallAccessProxyApiGateway6SslAlgorithm(d, i["ssl_algorithm"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandFirewallAccessProxyApiGateway6SslCipherSuites(d, i["ssl_cipher_suites"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["ssl-cipher-suites"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-dh-bits"], _ = expandFirewallAccessProxyApiGateway6SslDhBits(d, i["ssl_dh_bits"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-max-version"], _ = expandFirewallAccessProxyApiGateway6SslMaxVersion(d, i["ssl_max_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-min-version"], _ = expandFirewallAccessProxyApiGateway6SslMinVersion(d, i["ssl_min_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_renegotiation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-renegotiation"], _ = expandFirewallAccessProxyApiGateway6SslRenegotiation(d, i["ssl_renegotiation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_vpn_web_portal"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-vpn-web-portal"], _ = expandFirewallAccessProxyApiGateway6SslVpnWebPortal(d, i["ssl_vpn_web_portal"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["url-map"], _ = expandFirewallAccessProxyApiGateway6UrlMap(d, i["url_map"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["url-map-type"], _ = expandFirewallAccessProxyApiGateway6UrlMapType(d, i["url_map_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["virtual-host"], _ = expandFirewallAccessProxyApiGateway6VirtualHost(d, i["virtual_host"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxyApiGateway6Application(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyApiGateway6H2Support(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6H3Support(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6HttpCookieAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6HttpCookieDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6HttpCookieDomainFromHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6HttpCookieGeneration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6HttpCookiePath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6HttpCookieShare(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6HttpsCookieSecure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6LdbMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6Persistence(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6Quic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ack-delay-exponent"], _ = expandFirewallAccessProxyApiGateway6QuicAckDelayExponent(d, i["ack_delay_exponent"], pre_append)
	}
	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["active-connection-id-limit"], _ = expandFirewallAccessProxyApiGateway6QuicActiveConnectionIdLimit(d, i["active_connection_id_limit"], pre_append)
	}
	pre_append = pre + ".0." + "active_migration"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["active-migration"], _ = expandFirewallAccessProxyApiGateway6QuicActiveMigration(d, i["active_migration"], pre_append)
	}
	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["grease-quic-bit"], _ = expandFirewallAccessProxyApiGateway6QuicGreaseQuicBit(d, i["grease_quic_bit"], pre_append)
	}
	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-ack-delay"], _ = expandFirewallAccessProxyApiGateway6QuicMaxAckDelay(d, i["max_ack_delay"], pre_append)
	}
	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-datagram-frame-size"], _ = expandFirewallAccessProxyApiGateway6QuicMaxDatagramFrameSize(d, i["max_datagram_frame_size"], pre_append)
	}
	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-idle-timeout"], _ = expandFirewallAccessProxyApiGateway6QuicMaxIdleTimeout(d, i["max_idle_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-udp-payload-size"], _ = expandFirewallAccessProxyApiGateway6QuicMaxUdpPayloadSize(d, i["max_udp_payload_size"], pre_append)
	}

	return result, nil
}

func expandFirewallAccessProxyApiGateway6QuicAckDelayExponent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6QuicActiveConnectionIdLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6QuicActiveMigration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6QuicGreaseQuicBit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6QuicMaxAckDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6QuicMaxDatagramFrameSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6QuicMaxIdleTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6QuicMaxUdpPayloadSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6Realservers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["addr-type"], _ = expandFirewallAccessProxyApiGateway6RealserversAddrType(d, i["addr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["address"], _ = expandFirewallAccessProxyApiGateway6RealserversAddress(d, i["address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["domain"], _ = expandFirewallAccessProxyApiGateway6RealserversDomain(d, i["domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_auth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["external-auth"], _ = expandFirewallAccessProxyApiGateway6RealserversExternalAuth(d, i["external_auth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check"], _ = expandFirewallAccessProxyApiGateway6RealserversHealthCheck(d, i["health_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check-proto"], _ = expandFirewallAccessProxyApiGateway6RealserversHealthCheckProto(d, i["health_check_proto"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["holddown-interval"], _ = expandFirewallAccessProxyApiGateway6RealserversHolddownInterval(d, i["holddown_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-host"], _ = expandFirewallAccessProxyApiGateway6RealserversHttpHost(d, i["http_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFirewallAccessProxyApiGateway6RealserversId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandFirewallAccessProxyApiGateway6RealserversIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mappedport"], _ = expandFirewallAccessProxyApiGateway6RealserversMappedport(d, i["mappedport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandFirewallAccessProxyApiGateway6RealserversPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_client_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssh-client-cert"], _ = expandFirewallAccessProxyApiGateway6RealserversSshClientCert(d, i["ssh_client_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssh-host-key"], _ = expandFirewallAccessProxyApiGateway6RealserversSshHostKey(d, i["ssh_host_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key_validation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssh-host-key-validation"], _ = expandFirewallAccessProxyApiGateway6RealserversSshHostKeyValidation(d, i["ssh_host_key_validation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandFirewallAccessProxyApiGateway6RealserversStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["translate-host"], _ = expandFirewallAccessProxyApiGateway6RealserversTranslateHost(d, i["translate_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tunnel_encryption"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tunnel-encryption"], _ = expandFirewallAccessProxyApiGateway6RealserversTunnelEncryption(d, i["tunnel_encryption"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandFirewallAccessProxyApiGateway6RealserversType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "verify_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["verify-cert"], _ = expandFirewallAccessProxyApiGateway6RealserversVerifyCert(d, i["verify_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandFirewallAccessProxyApiGateway6RealserversWeight(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxyApiGateway6RealserversAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyApiGateway6RealserversDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversExternalAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversHealthCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversHealthCheckProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversHolddownInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversHttpHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversMappedport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyApiGateway6RealserversPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversSshClientCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyApiGateway6RealserversSshHostKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyApiGateway6RealserversSshHostKeyValidation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversTranslateHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversTunnelEncryption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversVerifyCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6SamlRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6SamlServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyApiGateway6Service(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6SslAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6SslCipherSuites(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["cipher"], _ = expandFirewallAccessProxyApiGateway6SslCipherSuitesCipher(d, i["cipher"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandFirewallAccessProxyApiGateway6SslCipherSuitesPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["versions"], _ = expandFirewallAccessProxyApiGateway6SslCipherSuitesVersions(d, i["versions"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxyApiGateway6SslCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6SslCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6SslCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyApiGateway6SslDhBits(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6SslMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6SslMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6SslRenegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6SslVpnWebPortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyApiGateway6UrlMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6UrlMapType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6VirtualHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyAuthPortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyAuthVirtualHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyClientCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyDecryptedTrafficMirror(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyEmptyCertAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyLdbMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyLogBlockedTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyRealservers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFirewallAccessProxyRealserversId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandFirewallAccessProxyRealserversIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandFirewallAccessProxyRealserversPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandFirewallAccessProxyRealserversStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandFirewallAccessProxyRealserversWeight(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxyRealserversId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyRealserversIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyRealserversPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyRealserversStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyRealserversWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyServerPubkeyAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyServerPubkeyAuthSettings(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auth_ca"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auth-ca"], _ = expandFirewallAccessProxyServerPubkeyAuthSettingsAuthCa(d, i["auth_ca"], pre_append)
	}
	pre_append = pre + ".0." + "cert_extension"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandFirewallAccessProxyServerPubkeyAuthSettingsCertExtension(d, i["cert_extension"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["cert-extension"] = t
		}
	}
	pre_append = pre + ".0." + "permit_agent_forwarding"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["permit-agent-forwarding"], _ = expandFirewallAccessProxyServerPubkeyAuthSettingsPermitAgentForwarding(d, i["permit_agent_forwarding"], pre_append)
	}
	pre_append = pre + ".0." + "permit_port_forwarding"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["permit-port-forwarding"], _ = expandFirewallAccessProxyServerPubkeyAuthSettingsPermitPortForwarding(d, i["permit_port_forwarding"], pre_append)
	}
	pre_append = pre + ".0." + "permit_pty"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["permit-pty"], _ = expandFirewallAccessProxyServerPubkeyAuthSettingsPermitPty(d, i["permit_pty"], pre_append)
	}
	pre_append = pre + ".0." + "permit_user_rc"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["permit-user-rc"], _ = expandFirewallAccessProxyServerPubkeyAuthSettingsPermitUserRc(d, i["permit_user_rc"], pre_append)
	}
	pre_append = pre + ".0." + "permit_x11_forwarding"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["permit-x11-forwarding"], _ = expandFirewallAccessProxyServerPubkeyAuthSettingsPermitX11Forwarding(d, i["permit_x11_forwarding"], pre_append)
	}
	pre_append = pre + ".0." + "source_address"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["source-address"], _ = expandFirewallAccessProxyServerPubkeyAuthSettingsSourceAddress(d, i["source_address"], pre_append)
	}

	return result, nil
}

func expandFirewallAccessProxyServerPubkeyAuthSettingsAuthCa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyServerPubkeyAuthSettingsCertExtension(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "critical"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["critical"], _ = expandFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionCritical(d, i["critical"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "data"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["data"], _ = expandFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionData(d, i["data"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionType(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionCritical(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionData(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyServerPubkeyAuthSettingsPermitAgentForwarding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyServerPubkeyAuthSettingsPermitPortForwarding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyServerPubkeyAuthSettingsPermitPty(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyServerPubkeyAuthSettingsPermitUserRc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyServerPubkeyAuthSettingsPermitX11Forwarding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyServerPubkeyAuthSettingsSourceAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxySvrPoolMultiplex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxySvrPoolServerMaxConcurrentRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxySvrPoolServerMaxRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxySvrPoolTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyUserAgentDetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyVip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyHttpSupportedMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallAccessProxy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("add_vhost_domain_to_dnsdb"); ok || d.HasChange("add_vhost_domain_to_dnsdb") {
		t, err := expandFirewallAccessProxyAddVhostDomainToDnsdb(d, v, "add_vhost_domain_to_dnsdb")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-vhost-domain-to-dnsdb"] = t
		}
	}

	if v, ok := d.GetOk("api_gateway"); ok || d.HasChange("api_gateway") {
		t, err := expandFirewallAccessProxyApiGateway(d, v, "api_gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["api-gateway"] = t
		}
	}

	if v, ok := d.GetOk("api_gateway6"); ok || d.HasChange("api_gateway6") {
		t, err := expandFirewallAccessProxyApiGateway6(d, v, "api_gateway6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["api-gateway6"] = t
		}
	}

	if v, ok := d.GetOk("auth_portal"); ok || d.HasChange("auth_portal") {
		t, err := expandFirewallAccessProxyAuthPortal(d, v, "auth_portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-portal"] = t
		}
	}

	if v, ok := d.GetOk("auth_virtual_host"); ok || d.HasChange("auth_virtual_host") {
		t, err := expandFirewallAccessProxyAuthVirtualHost(d, v, "auth_virtual_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-virtual-host"] = t
		}
	}

	if v, ok := d.GetOk("client_cert"); ok || d.HasChange("client_cert") {
		t, err := expandFirewallAccessProxyClientCert(d, v, "client_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert"] = t
		}
	}

	if v, ok := d.GetOk("decrypted_traffic_mirror"); ok || d.HasChange("decrypted_traffic_mirror") {
		t, err := expandFirewallAccessProxyDecryptedTrafficMirror(d, v, "decrypted_traffic_mirror")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["decrypted-traffic-mirror"] = t
		}
	}

	if v, ok := d.GetOk("empty_cert_action"); ok || d.HasChange("empty_cert_action") {
		t, err := expandFirewallAccessProxyEmptyCertAction(d, v, "empty_cert_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["empty-cert-action"] = t
		}
	}

	if v, ok := d.GetOk("ldb_method"); ok || d.HasChange("ldb_method") {
		t, err := expandFirewallAccessProxyLdbMethod(d, v, "ldb_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldb-method"] = t
		}
	}

	if v, ok := d.GetOk("log_blocked_traffic"); ok || d.HasChange("log_blocked_traffic") {
		t, err := expandFirewallAccessProxyLogBlockedTraffic(d, v, "log_blocked_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-blocked-traffic"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallAccessProxyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("realservers"); ok || d.HasChange("realservers") {
		t, err := expandFirewallAccessProxyRealservers(d, v, "realservers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["realservers"] = t
		}
	}

	if v, ok := d.GetOk("server_pubkey_auth"); ok || d.HasChange("server_pubkey_auth") {
		t, err := expandFirewallAccessProxyServerPubkeyAuth(d, v, "server_pubkey_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-pubkey-auth"] = t
		}
	}

	if v, ok := d.GetOk("server_pubkey_auth_settings"); ok || d.HasChange("server_pubkey_auth_settings") {
		t, err := expandFirewallAccessProxyServerPubkeyAuthSettings(d, v, "server_pubkey_auth_settings")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-pubkey-auth-settings"] = t
		}
	}

	if v, ok := d.GetOk("svr_pool_multiplex"); ok || d.HasChange("svr_pool_multiplex") {
		t, err := expandFirewallAccessProxySvrPoolMultiplex(d, v, "svr_pool_multiplex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["svr-pool-multiplex"] = t
		}
	}

	if v, ok := d.GetOk("svr_pool_server_max_concurrent_request"); ok || d.HasChange("svr_pool_server_max_concurrent_request") {
		t, err := expandFirewallAccessProxySvrPoolServerMaxConcurrentRequest(d, v, "svr_pool_server_max_concurrent_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["svr-pool-server-max-concurrent-request"] = t
		}
	}

	if v, ok := d.GetOk("svr_pool_server_max_request"); ok || d.HasChange("svr_pool_server_max_request") {
		t, err := expandFirewallAccessProxySvrPoolServerMaxRequest(d, v, "svr_pool_server_max_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["svr-pool-server-max-request"] = t
		}
	}

	if v, ok := d.GetOk("svr_pool_ttl"); ok || d.HasChange("svr_pool_ttl") {
		t, err := expandFirewallAccessProxySvrPoolTtl(d, v, "svr_pool_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["svr-pool-ttl"] = t
		}
	}

	if v, ok := d.GetOk("user_agent_detect"); ok || d.HasChange("user_agent_detect") {
		t, err := expandFirewallAccessProxyUserAgentDetect(d, v, "user_agent_detect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-agent-detect"] = t
		}
	}

	if v, ok := d.GetOk("vip"); ok || d.HasChange("vip") {
		t, err := expandFirewallAccessProxyVip(d, v, "vip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vip"] = t
		}
	}

	if v, ok := d.GetOk("http_supported_max_version"); ok || d.HasChange("http_supported_max_version") {
		t, err := expandFirewallAccessProxyHttpSupportedMaxVersion(d, v, "http_supported_max_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-supported-max-version"] = t
		}
	}

	return &obj, nil
}
