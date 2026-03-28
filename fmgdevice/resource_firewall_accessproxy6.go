// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure IPv6 access proxy.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallAccessProxy6() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallAccessProxy6Create,
		Read:   resourceFirewallAccessProxy6Read,
		Update: resourceFirewallAccessProxy6Update,
		Delete: resourceFirewallAccessProxy6Delete,

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
			"http_supported_max_version": &schema.Schema{
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceFirewallAccessProxy6Create(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallAccessProxy6(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallAccessProxy6 resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallAccessProxy6(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallAccessProxy6(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallAccessProxy6 resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallAccessProxy6(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallAccessProxy6 resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallAccessProxy6Read(d, m)
}

func resourceFirewallAccessProxy6Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallAccessProxy6(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAccessProxy6 resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallAccessProxy6(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAccessProxy6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallAccessProxy6Read(d, m)
}

func resourceFirewallAccessProxy6Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallAccessProxy6(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallAccessProxy6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAccessProxy6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallAccessProxy6(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallAccessProxy6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallAccessProxy6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAccessProxy6 resource from API: %v", err)
	}
	return nil
}

func flattenFirewallAccessProxy6AddVhostDomainToDnsdb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallAccessProxy6ApiGatewayApplication(i["application"], d, pre_append)
			tmp["application"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway-Application")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h2_support"
		if _, ok := i["h2-support"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayH2Support(i["h2-support"], d, pre_append)
			tmp["h2_support"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway-H2Support")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h3_support"
		if _, ok := i["h3-support"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayH3Support(i["h3-support"], d, pre_append)
			tmp["h3_support"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway-H3Support")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := i["http-cookie-age"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayHttpCookieAge(i["http-cookie-age"], d, pre_append)
			tmp["http_cookie_age"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway-HttpCookieAge")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := i["http-cookie-domain"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayHttpCookieDomain(i["http-cookie-domain"], d, pre_append)
			tmp["http_cookie_domain"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway-HttpCookieDomain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := i["http-cookie-domain-from-host"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayHttpCookieDomainFromHost(i["http-cookie-domain-from-host"], d, pre_append)
			tmp["http_cookie_domain_from_host"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway-HttpCookieDomainFromHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := i["http-cookie-generation"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayHttpCookieGeneration(i["http-cookie-generation"], d, pre_append)
			tmp["http_cookie_generation"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway-HttpCookieGeneration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := i["http-cookie-path"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayHttpCookiePath(i["http-cookie-path"], d, pre_append)
			tmp["http_cookie_path"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway-HttpCookiePath")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := i["http-cookie-share"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayHttpCookieShare(i["http-cookie-share"], d, pre_append)
			tmp["http_cookie_share"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway-HttpCookieShare")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := i["https-cookie-secure"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayHttpsCookieSecure(i["https-cookie-secure"], d, pre_append)
			tmp["https_cookie_secure"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway-HttpsCookieSecure")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := i["ldb-method"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayLdbMethod(i["ldb-method"], d, pre_append)
			tmp["ldb_method"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway-LdbMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := i["persistence"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayPersistence(i["persistence"], d, pre_append)
			tmp["persistence"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway-Persistence")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quic"
		if _, ok := i["quic"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayQuic(i["quic"], d, pre_append)
			tmp["quic"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway-Quic")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := i["realservers"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealservers(i["realservers"], d, pre_append)
			tmp["realservers"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway-Realservers")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_redirect"
		if _, ok := i["saml-redirect"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewaySamlRedirect(i["saml-redirect"], d, pre_append)
			tmp["saml_redirect"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway-SamlRedirect")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_server"
		if _, ok := i["saml-server"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewaySamlServer(i["saml-server"], d, pre_append)
			tmp["saml_server"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway-SamlServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := i["service"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayService(i["service"], d, pre_append)
			tmp["service"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway-Service")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := i["ssl-algorithm"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewaySslAlgorithm(i["ssl-algorithm"], d, pre_append)
			tmp["ssl_algorithm"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway-SslAlgorithm")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := i["ssl-cipher-suites"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewaySslCipherSuites(i["ssl-cipher-suites"], d, pre_append)
			tmp["ssl_cipher_suites"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway-SslCipherSuites")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := i["ssl-dh-bits"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewaySslDhBits(i["ssl-dh-bits"], d, pre_append)
			tmp["ssl_dh_bits"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway-SslDhBits")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := i["ssl-max-version"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewaySslMaxVersion(i["ssl-max-version"], d, pre_append)
			tmp["ssl_max_version"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway-SslMaxVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := i["ssl-min-version"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewaySslMinVersion(i["ssl-min-version"], d, pre_append)
			tmp["ssl_min_version"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway-SslMinVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_renegotiation"
		if _, ok := i["ssl-renegotiation"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewaySslRenegotiation(i["ssl-renegotiation"], d, pre_append)
			tmp["ssl_renegotiation"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway-SslRenegotiation")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_vpn_web_portal"
		if _, ok := i["ssl-vpn-web-portal"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewaySslVpnWebPortal(i["ssl-vpn-web-portal"], d, pre_append)
			tmp["ssl_vpn_web_portal"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway-SslVpnWebPortal")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map"
		if _, ok := i["url-map"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayUrlMap(i["url-map"], d, pre_append)
			tmp["url_map"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway-UrlMap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map_type"
		if _, ok := i["url-map-type"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayUrlMapType(i["url-map-type"], d, pre_append)
			tmp["url_map_type"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway-UrlMapType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_host"
		if _, ok := i["virtual-host"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayVirtualHost(i["virtual-host"], d, pre_append)
			tmp["virtual_host"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway-VirtualHost")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAccessProxy6ApiGatewayApplication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGatewayH2Support(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayH3Support(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayHttpCookieAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayHttpCookieDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayHttpCookieDomainFromHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayHttpCookieGeneration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayHttpCookiePath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayHttpCookieShare(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayHttpsCookieSecure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayLdbMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayPersistence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayQuic(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := i["ack-delay-exponent"]; ok {
		result["ack_delay_exponent"] = flattenFirewallAccessProxy6ApiGatewayQuicAckDelayExponent(i["ack-delay-exponent"], d, pre_append)
	}

	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := i["active-connection-id-limit"]; ok {
		result["active_connection_id_limit"] = flattenFirewallAccessProxy6ApiGatewayQuicActiveConnectionIdLimit(i["active-connection-id-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "active_migration"
	if _, ok := i["active-migration"]; ok {
		result["active_migration"] = flattenFirewallAccessProxy6ApiGatewayQuicActiveMigration(i["active-migration"], d, pre_append)
	}

	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := i["grease-quic-bit"]; ok {
		result["grease_quic_bit"] = flattenFirewallAccessProxy6ApiGatewayQuicGreaseQuicBit(i["grease-quic-bit"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := i["max-ack-delay"]; ok {
		result["max_ack_delay"] = flattenFirewallAccessProxy6ApiGatewayQuicMaxAckDelay(i["max-ack-delay"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := i["max-datagram-frame-size"]; ok {
		result["max_datagram_frame_size"] = flattenFirewallAccessProxy6ApiGatewayQuicMaxDatagramFrameSize(i["max-datagram-frame-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := i["max-idle-timeout"]; ok {
		result["max_idle_timeout"] = flattenFirewallAccessProxy6ApiGatewayQuicMaxIdleTimeout(i["max-idle-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := i["max-udp-payload-size"]; ok {
		result["max_udp_payload_size"] = flattenFirewallAccessProxy6ApiGatewayQuicMaxUdpPayloadSize(i["max-udp-payload-size"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallAccessProxy6ApiGatewayQuicAckDelayExponent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayQuicActiveConnectionIdLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayQuicActiveMigration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayQuicGreaseQuicBit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayQuicMaxAckDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayQuicMaxDatagramFrameSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayQuicMaxIdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayQuicMaxUdpPayloadSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealservers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallAccessProxy6ApiGatewayRealserversAddrType(i["addr-type"], d, pre_append)
			tmp["addr_type"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-AddrType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversAddress(i["address"], d, pre_append)
			tmp["address"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := i["domain"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversDomain(i["domain"], d, pre_append)
			tmp["domain"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-Domain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_auth"
		if _, ok := i["external-auth"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversExternalAuth(i["external-auth"], d, pre_append)
			tmp["external_auth"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-ExternalAuth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := i["health-check"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversHealthCheck(i["health-check"], d, pre_append)
			tmp["health_check"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-HealthCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := i["health-check-proto"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversHealthCheckProto(i["health-check-proto"], d, pre_append)
			tmp["health_check_proto"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-HealthCheckProto")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := i["holddown-interval"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversHolddownInterval(i["holddown-interval"], d, pre_append)
			tmp["holddown_interval"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-HolddownInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := i["http-host"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversHttpHost(i["http-host"], d, pre_append)
			tmp["http_host"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-HttpHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := i["mappedport"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversMappedport(i["mappedport"], d, pre_append)
			tmp["mappedport"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-Mappedport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_client_cert"
		if _, ok := i["ssh-client-cert"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversSshClientCert(i["ssh-client-cert"], d, pre_append)
			tmp["ssh_client_cert"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-SshClientCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key"
		if _, ok := i["ssh-host-key"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversSshHostKey(i["ssh-host-key"], d, pre_append)
			tmp["ssh_host_key"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-SshHostKey")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key_validation"
		if _, ok := i["ssh-host-key-validation"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversSshHostKeyValidation(i["ssh-host-key-validation"], d, pre_append)
			tmp["ssh_host_key_validation"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-SshHostKeyValidation")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := i["translate-host"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversTranslateHost(i["translate-host"], d, pre_append)
			tmp["translate_host"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-TranslateHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tunnel_encryption"
		if _, ok := i["tunnel-encryption"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversTunnelEncryption(i["tunnel-encryption"], d, pre_append)
			tmp["tunnel_encryption"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-TunnelEncryption")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "verify_cert"
		if _, ok := i["verify-cert"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversVerifyCert(i["verify-cert"], d, pre_append)
			tmp["verify_cert"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-VerifyCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewayRealserversWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-Realservers-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAccessProxy6ApiGatewayRealserversAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGatewayRealserversDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversExternalAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversHealthCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversHealthCheckProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversHolddownInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversHttpHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversMappedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGatewayRealserversPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversSshClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGatewayRealserversSshHostKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGatewayRealserversSshHostKeyValidation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversTranslateHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversTunnelEncryption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversVerifyCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewaySamlRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewaySamlServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGatewayService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewaySslAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewaySslCipherSuites(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallAccessProxy6ApiGatewaySslCipherSuitesCipher(i["cipher"], d, pre_append)
			tmp["cipher"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-SslCipherSuites-Cipher")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewaySslCipherSuitesPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-SslCipherSuites-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {
			v := flattenFirewallAccessProxy6ApiGatewaySslCipherSuitesVersions(i["versions"], d, pre_append)
			tmp["versions"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway-SslCipherSuites-Versions")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAccessProxy6ApiGatewaySslCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewaySslCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewaySslCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGatewaySslDhBits(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewaySslMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewaySslMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewaySslRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewaySslVpnWebPortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGatewayUrlMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayUrlMapType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayVirtualHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGateway6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallAccessProxy6ApiGateway6Application(i["application"], d, pre_append)
			tmp["application"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway6-Application")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h2_support"
		if _, ok := i["h2-support"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6H2Support(i["h2-support"], d, pre_append)
			tmp["h2_support"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway6-H2Support")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h3_support"
		if _, ok := i["h3-support"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6H3Support(i["h3-support"], d, pre_append)
			tmp["h3_support"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway6-H3Support")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := i["http-cookie-age"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6HttpCookieAge(i["http-cookie-age"], d, pre_append)
			tmp["http_cookie_age"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway6-HttpCookieAge")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := i["http-cookie-domain"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6HttpCookieDomain(i["http-cookie-domain"], d, pre_append)
			tmp["http_cookie_domain"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway6-HttpCookieDomain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := i["http-cookie-domain-from-host"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6HttpCookieDomainFromHost(i["http-cookie-domain-from-host"], d, pre_append)
			tmp["http_cookie_domain_from_host"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway6-HttpCookieDomainFromHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := i["http-cookie-generation"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6HttpCookieGeneration(i["http-cookie-generation"], d, pre_append)
			tmp["http_cookie_generation"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway6-HttpCookieGeneration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := i["http-cookie-path"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6HttpCookiePath(i["http-cookie-path"], d, pre_append)
			tmp["http_cookie_path"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway6-HttpCookiePath")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := i["http-cookie-share"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6HttpCookieShare(i["http-cookie-share"], d, pre_append)
			tmp["http_cookie_share"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway6-HttpCookieShare")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := i["https-cookie-secure"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6HttpsCookieSecure(i["https-cookie-secure"], d, pre_append)
			tmp["https_cookie_secure"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway6-HttpsCookieSecure")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6Id(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway6-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := i["ldb-method"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6LdbMethod(i["ldb-method"], d, pre_append)
			tmp["ldb_method"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway6-LdbMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := i["persistence"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6Persistence(i["persistence"], d, pre_append)
			tmp["persistence"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway6-Persistence")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quic"
		if _, ok := i["quic"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6Quic(i["quic"], d, pre_append)
			tmp["quic"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway6-Quic")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := i["realservers"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6Realservers(i["realservers"], d, pre_append)
			tmp["realservers"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway6-Realservers")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_redirect"
		if _, ok := i["saml-redirect"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6SamlRedirect(i["saml-redirect"], d, pre_append)
			tmp["saml_redirect"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway6-SamlRedirect")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_server"
		if _, ok := i["saml-server"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6SamlServer(i["saml-server"], d, pre_append)
			tmp["saml_server"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway6-SamlServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := i["service"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6Service(i["service"], d, pre_append)
			tmp["service"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway6-Service")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := i["ssl-algorithm"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6SslAlgorithm(i["ssl-algorithm"], d, pre_append)
			tmp["ssl_algorithm"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway6-SslAlgorithm")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := i["ssl-cipher-suites"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6SslCipherSuites(i["ssl-cipher-suites"], d, pre_append)
			tmp["ssl_cipher_suites"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway6-SslCipherSuites")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := i["ssl-dh-bits"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6SslDhBits(i["ssl-dh-bits"], d, pre_append)
			tmp["ssl_dh_bits"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway6-SslDhBits")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := i["ssl-max-version"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6SslMaxVersion(i["ssl-max-version"], d, pre_append)
			tmp["ssl_max_version"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway6-SslMaxVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := i["ssl-min-version"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6SslMinVersion(i["ssl-min-version"], d, pre_append)
			tmp["ssl_min_version"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway6-SslMinVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_renegotiation"
		if _, ok := i["ssl-renegotiation"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6SslRenegotiation(i["ssl-renegotiation"], d, pre_append)
			tmp["ssl_renegotiation"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway6-SslRenegotiation")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_vpn_web_portal"
		if _, ok := i["ssl-vpn-web-portal"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6SslVpnWebPortal(i["ssl-vpn-web-portal"], d, pre_append)
			tmp["ssl_vpn_web_portal"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway6-SslVpnWebPortal")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map"
		if _, ok := i["url-map"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6UrlMap(i["url-map"], d, pre_append)
			tmp["url_map"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway6-UrlMap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map_type"
		if _, ok := i["url-map-type"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6UrlMapType(i["url-map-type"], d, pre_append)
			tmp["url_map_type"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway6-UrlMapType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_host"
		if _, ok := i["virtual-host"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6VirtualHost(i["virtual-host"], d, pre_append)
			tmp["virtual_host"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6-ApiGateway6-VirtualHost")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAccessProxy6ApiGateway6Application(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGateway6H2Support(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6H3Support(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6HttpCookieAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6HttpCookieDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6HttpCookieDomainFromHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6HttpCookieGeneration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6HttpCookiePath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6HttpCookieShare(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6HttpsCookieSecure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6LdbMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6Persistence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6Quic(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := i["ack-delay-exponent"]; ok {
		result["ack_delay_exponent"] = flattenFirewallAccessProxy6ApiGateway6QuicAckDelayExponent(i["ack-delay-exponent"], d, pre_append)
	}

	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := i["active-connection-id-limit"]; ok {
		result["active_connection_id_limit"] = flattenFirewallAccessProxy6ApiGateway6QuicActiveConnectionIdLimit(i["active-connection-id-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "active_migration"
	if _, ok := i["active-migration"]; ok {
		result["active_migration"] = flattenFirewallAccessProxy6ApiGateway6QuicActiveMigration(i["active-migration"], d, pre_append)
	}

	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := i["grease-quic-bit"]; ok {
		result["grease_quic_bit"] = flattenFirewallAccessProxy6ApiGateway6QuicGreaseQuicBit(i["grease-quic-bit"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := i["max-ack-delay"]; ok {
		result["max_ack_delay"] = flattenFirewallAccessProxy6ApiGateway6QuicMaxAckDelay(i["max-ack-delay"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := i["max-datagram-frame-size"]; ok {
		result["max_datagram_frame_size"] = flattenFirewallAccessProxy6ApiGateway6QuicMaxDatagramFrameSize(i["max-datagram-frame-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := i["max-idle-timeout"]; ok {
		result["max_idle_timeout"] = flattenFirewallAccessProxy6ApiGateway6QuicMaxIdleTimeout(i["max-idle-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := i["max-udp-payload-size"]; ok {
		result["max_udp_payload_size"] = flattenFirewallAccessProxy6ApiGateway6QuicMaxUdpPayloadSize(i["max-udp-payload-size"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallAccessProxy6ApiGateway6QuicAckDelayExponent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6QuicActiveConnectionIdLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6QuicActiveMigration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6QuicGreaseQuicBit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6QuicMaxAckDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6QuicMaxDatagramFrameSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6QuicMaxIdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6QuicMaxUdpPayloadSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6Realservers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallAccessProxy6ApiGateway6RealserversAddrType(i["addr-type"], d, pre_append)
			tmp["addr_type"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway6-Realservers-AddrType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6RealserversAddress(i["address"], d, pre_append)
			tmp["address"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway6-Realservers-Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := i["domain"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6RealserversDomain(i["domain"], d, pre_append)
			tmp["domain"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway6-Realservers-Domain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_auth"
		if _, ok := i["external-auth"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6RealserversExternalAuth(i["external-auth"], d, pre_append)
			tmp["external_auth"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway6-Realservers-ExternalAuth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := i["health-check"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6RealserversHealthCheck(i["health-check"], d, pre_append)
			tmp["health_check"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway6-Realservers-HealthCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := i["health-check-proto"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6RealserversHealthCheckProto(i["health-check-proto"], d, pre_append)
			tmp["health_check_proto"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway6-Realservers-HealthCheckProto")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := i["holddown-interval"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6RealserversHolddownInterval(i["holddown-interval"], d, pre_append)
			tmp["holddown_interval"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway6-Realservers-HolddownInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := i["http-host"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6RealserversHttpHost(i["http-host"], d, pre_append)
			tmp["http_host"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway6-Realservers-HttpHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6RealserversId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway6-Realservers-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6RealserversIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway6-Realservers-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := i["mappedport"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6RealserversMappedport(i["mappedport"], d, pre_append)
			tmp["mappedport"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway6-Realservers-Mappedport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6RealserversPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway6-Realservers-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_client_cert"
		if _, ok := i["ssh-client-cert"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6RealserversSshClientCert(i["ssh-client-cert"], d, pre_append)
			tmp["ssh_client_cert"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway6-Realservers-SshClientCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key"
		if _, ok := i["ssh-host-key"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6RealserversSshHostKey(i["ssh-host-key"], d, pre_append)
			tmp["ssh_host_key"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway6-Realservers-SshHostKey")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key_validation"
		if _, ok := i["ssh-host-key-validation"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6RealserversSshHostKeyValidation(i["ssh-host-key-validation"], d, pre_append)
			tmp["ssh_host_key_validation"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway6-Realservers-SshHostKeyValidation")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6RealserversStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway6-Realservers-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := i["translate-host"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6RealserversTranslateHost(i["translate-host"], d, pre_append)
			tmp["translate_host"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway6-Realservers-TranslateHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tunnel_encryption"
		if _, ok := i["tunnel-encryption"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6RealserversTunnelEncryption(i["tunnel-encryption"], d, pre_append)
			tmp["tunnel_encryption"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway6-Realservers-TunnelEncryption")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6RealserversType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway6-Realservers-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "verify_cert"
		if _, ok := i["verify-cert"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6RealserversVerifyCert(i["verify-cert"], d, pre_append)
			tmp["verify_cert"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway6-Realservers-VerifyCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6RealserversWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway6-Realservers-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAccessProxy6ApiGateway6RealserversAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGateway6RealserversDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversExternalAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversHealthCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversHealthCheckProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversHolddownInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversHttpHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversMappedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGateway6RealserversPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversSshClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGateway6RealserversSshHostKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGateway6RealserversSshHostKeyValidation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversTranslateHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversTunnelEncryption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversVerifyCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6SamlRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6SamlServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGateway6Service(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6SslAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6SslCipherSuites(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallAccessProxy6ApiGateway6SslCipherSuitesCipher(i["cipher"], d, pre_append)
			tmp["cipher"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway6-SslCipherSuites-Cipher")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6SslCipherSuitesPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway6-SslCipherSuites-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {
			v := flattenFirewallAccessProxy6ApiGateway6SslCipherSuitesVersions(i["versions"], d, pre_append)
			tmp["versions"] = fortiAPISubPartPatch(v, "FirewallAccessProxy6ApiGateway6-SslCipherSuites-Versions")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAccessProxy6ApiGateway6SslCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6SslCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6SslCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGateway6SslDhBits(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6SslMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6SslMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6SslRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6SslVpnWebPortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGateway6UrlMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6UrlMapType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6VirtualHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6AuthPortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6AuthVirtualHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6DecryptedTrafficMirror(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6EmptyCertAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6HttpSupportedMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6LogBlockedTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6SvrPoolMultiplex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6SvrPoolServerMaxConcurrentRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6SvrPoolServerMaxRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6SvrPoolTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6UserAgentDetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6Vip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectFirewallAccessProxy6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("add_vhost_domain_to_dnsdb", flattenFirewallAccessProxy6AddVhostDomainToDnsdb(o["add-vhost-domain-to-dnsdb"], d, "add_vhost_domain_to_dnsdb")); err != nil {
		if vv, ok := fortiAPIPatch(o["add-vhost-domain-to-dnsdb"], "FirewallAccessProxy6-AddVhostDomainToDnsdb"); ok {
			if err = d.Set("add_vhost_domain_to_dnsdb", vv); err != nil {
				return fmt.Errorf("Error reading add_vhost_domain_to_dnsdb: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading add_vhost_domain_to_dnsdb: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("api_gateway", flattenFirewallAccessProxy6ApiGateway(o["api-gateway"], d, "api_gateway")); err != nil {
			if vv, ok := fortiAPIPatch(o["api-gateway"], "FirewallAccessProxy6-ApiGateway"); ok {
				if err = d.Set("api_gateway", vv); err != nil {
					return fmt.Errorf("Error reading api_gateway: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading api_gateway: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("api_gateway"); ok {
			if err = d.Set("api_gateway", flattenFirewallAccessProxy6ApiGateway(o["api-gateway"], d, "api_gateway")); err != nil {
				if vv, ok := fortiAPIPatch(o["api-gateway"], "FirewallAccessProxy6-ApiGateway"); ok {
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
		if err = d.Set("api_gateway6", flattenFirewallAccessProxy6ApiGateway6(o["api-gateway6"], d, "api_gateway6")); err != nil {
			if vv, ok := fortiAPIPatch(o["api-gateway6"], "FirewallAccessProxy6-ApiGateway6"); ok {
				if err = d.Set("api_gateway6", vv); err != nil {
					return fmt.Errorf("Error reading api_gateway6: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading api_gateway6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("api_gateway6"); ok {
			if err = d.Set("api_gateway6", flattenFirewallAccessProxy6ApiGateway6(o["api-gateway6"], d, "api_gateway6")); err != nil {
				if vv, ok := fortiAPIPatch(o["api-gateway6"], "FirewallAccessProxy6-ApiGateway6"); ok {
					if err = d.Set("api_gateway6", vv); err != nil {
						return fmt.Errorf("Error reading api_gateway6: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading api_gateway6: %v", err)
				}
			}
		}
	}

	if err = d.Set("auth_portal", flattenFirewallAccessProxy6AuthPortal(o["auth-portal"], d, "auth_portal")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-portal"], "FirewallAccessProxy6-AuthPortal"); ok {
			if err = d.Set("auth_portal", vv); err != nil {
				return fmt.Errorf("Error reading auth_portal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_portal: %v", err)
		}
	}

	if err = d.Set("auth_virtual_host", flattenFirewallAccessProxy6AuthVirtualHost(o["auth-virtual-host"], d, "auth_virtual_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-virtual-host"], "FirewallAccessProxy6-AuthVirtualHost"); ok {
			if err = d.Set("auth_virtual_host", vv); err != nil {
				return fmt.Errorf("Error reading auth_virtual_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_virtual_host: %v", err)
		}
	}

	if err = d.Set("client_cert", flattenFirewallAccessProxy6ClientCert(o["client-cert"], d, "client_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-cert"], "FirewallAccessProxy6-ClientCert"); ok {
			if err = d.Set("client_cert", vv); err != nil {
				return fmt.Errorf("Error reading client_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_cert: %v", err)
		}
	}

	if err = d.Set("decrypted_traffic_mirror", flattenFirewallAccessProxy6DecryptedTrafficMirror(o["decrypted-traffic-mirror"], d, "decrypted_traffic_mirror")); err != nil {
		if vv, ok := fortiAPIPatch(o["decrypted-traffic-mirror"], "FirewallAccessProxy6-DecryptedTrafficMirror"); ok {
			if err = d.Set("decrypted_traffic_mirror", vv); err != nil {
				return fmt.Errorf("Error reading decrypted_traffic_mirror: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading decrypted_traffic_mirror: %v", err)
		}
	}

	if err = d.Set("empty_cert_action", flattenFirewallAccessProxy6EmptyCertAction(o["empty-cert-action"], d, "empty_cert_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["empty-cert-action"], "FirewallAccessProxy6-EmptyCertAction"); ok {
			if err = d.Set("empty_cert_action", vv); err != nil {
				return fmt.Errorf("Error reading empty_cert_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading empty_cert_action: %v", err)
		}
	}

	if err = d.Set("http_supported_max_version", flattenFirewallAccessProxy6HttpSupportedMaxVersion(o["http-supported-max-version"], d, "http_supported_max_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-supported-max-version"], "FirewallAccessProxy6-HttpSupportedMaxVersion"); ok {
			if err = d.Set("http_supported_max_version", vv); err != nil {
				return fmt.Errorf("Error reading http_supported_max_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_supported_max_version: %v", err)
		}
	}

	if err = d.Set("log_blocked_traffic", flattenFirewallAccessProxy6LogBlockedTraffic(o["log-blocked-traffic"], d, "log_blocked_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-blocked-traffic"], "FirewallAccessProxy6-LogBlockedTraffic"); ok {
			if err = d.Set("log_blocked_traffic", vv); err != nil {
				return fmt.Errorf("Error reading log_blocked_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_blocked_traffic: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallAccessProxy6Name(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallAccessProxy6-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("svr_pool_multiplex", flattenFirewallAccessProxy6SvrPoolMultiplex(o["svr-pool-multiplex"], d, "svr_pool_multiplex")); err != nil {
		if vv, ok := fortiAPIPatch(o["svr-pool-multiplex"], "FirewallAccessProxy6-SvrPoolMultiplex"); ok {
			if err = d.Set("svr_pool_multiplex", vv); err != nil {
				return fmt.Errorf("Error reading svr_pool_multiplex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading svr_pool_multiplex: %v", err)
		}
	}

	if err = d.Set("svr_pool_server_max_concurrent_request", flattenFirewallAccessProxy6SvrPoolServerMaxConcurrentRequest(o["svr-pool-server-max-concurrent-request"], d, "svr_pool_server_max_concurrent_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["svr-pool-server-max-concurrent-request"], "FirewallAccessProxy6-SvrPoolServerMaxConcurrentRequest"); ok {
			if err = d.Set("svr_pool_server_max_concurrent_request", vv); err != nil {
				return fmt.Errorf("Error reading svr_pool_server_max_concurrent_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading svr_pool_server_max_concurrent_request: %v", err)
		}
	}

	if err = d.Set("svr_pool_server_max_request", flattenFirewallAccessProxy6SvrPoolServerMaxRequest(o["svr-pool-server-max-request"], d, "svr_pool_server_max_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["svr-pool-server-max-request"], "FirewallAccessProxy6-SvrPoolServerMaxRequest"); ok {
			if err = d.Set("svr_pool_server_max_request", vv); err != nil {
				return fmt.Errorf("Error reading svr_pool_server_max_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading svr_pool_server_max_request: %v", err)
		}
	}

	if err = d.Set("svr_pool_ttl", flattenFirewallAccessProxy6SvrPoolTtl(o["svr-pool-ttl"], d, "svr_pool_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["svr-pool-ttl"], "FirewallAccessProxy6-SvrPoolTtl"); ok {
			if err = d.Set("svr_pool_ttl", vv); err != nil {
				return fmt.Errorf("Error reading svr_pool_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading svr_pool_ttl: %v", err)
		}
	}

	if err = d.Set("user_agent_detect", flattenFirewallAccessProxy6UserAgentDetect(o["user-agent-detect"], d, "user_agent_detect")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-agent-detect"], "FirewallAccessProxy6-UserAgentDetect"); ok {
			if err = d.Set("user_agent_detect", vv); err != nil {
				return fmt.Errorf("Error reading user_agent_detect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_agent_detect: %v", err)
		}
	}

	if err = d.Set("vip", flattenFirewallAccessProxy6Vip(o["vip"], d, "vip")); err != nil {
		if vv, ok := fortiAPIPatch(o["vip"], "FirewallAccessProxy6-Vip"); ok {
			if err = d.Set("vip", vv); err != nil {
				return fmt.Errorf("Error reading vip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vip: %v", err)
		}
	}

	return nil
}

func flattenFirewallAccessProxy6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallAccessProxy6AddVhostDomainToDnsdb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["application"], _ = expandFirewallAccessProxy6ApiGatewayApplication(d, i["application"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h2_support"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["h2-support"], _ = expandFirewallAccessProxy6ApiGatewayH2Support(d, i["h2_support"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h3_support"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["h3-support"], _ = expandFirewallAccessProxy6ApiGatewayH3Support(d, i["h3_support"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-age"], _ = expandFirewallAccessProxy6ApiGatewayHttpCookieAge(d, i["http_cookie_age"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-domain"], _ = expandFirewallAccessProxy6ApiGatewayHttpCookieDomain(d, i["http_cookie_domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-domain-from-host"], _ = expandFirewallAccessProxy6ApiGatewayHttpCookieDomainFromHost(d, i["http_cookie_domain_from_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-generation"], _ = expandFirewallAccessProxy6ApiGatewayHttpCookieGeneration(d, i["http_cookie_generation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-path"], _ = expandFirewallAccessProxy6ApiGatewayHttpCookiePath(d, i["http_cookie_path"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-share"], _ = expandFirewallAccessProxy6ApiGatewayHttpCookieShare(d, i["http_cookie_share"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["https-cookie-secure"], _ = expandFirewallAccessProxy6ApiGatewayHttpsCookieSecure(d, i["https_cookie_secure"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFirewallAccessProxy6ApiGatewayId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ldb-method"], _ = expandFirewallAccessProxy6ApiGatewayLdbMethod(d, i["ldb_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["persistence"], _ = expandFirewallAccessProxy6ApiGatewayPersistence(d, i["persistence"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quic"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandFirewallAccessProxy6ApiGatewayQuic(d, i["quic"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["quic"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandFirewallAccessProxy6ApiGatewayRealservers(d, i["realservers"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["realservers"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_redirect"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["saml-redirect"], _ = expandFirewallAccessProxy6ApiGatewaySamlRedirect(d, i["saml_redirect"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["saml-server"], _ = expandFirewallAccessProxy6ApiGatewaySamlServer(d, i["saml_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service"], _ = expandFirewallAccessProxy6ApiGatewayService(d, i["service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-algorithm"], _ = expandFirewallAccessProxy6ApiGatewaySslAlgorithm(d, i["ssl_algorithm"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandFirewallAccessProxy6ApiGatewaySslCipherSuites(d, i["ssl_cipher_suites"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["ssl-cipher-suites"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-dh-bits"], _ = expandFirewallAccessProxy6ApiGatewaySslDhBits(d, i["ssl_dh_bits"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-max-version"], _ = expandFirewallAccessProxy6ApiGatewaySslMaxVersion(d, i["ssl_max_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-min-version"], _ = expandFirewallAccessProxy6ApiGatewaySslMinVersion(d, i["ssl_min_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_renegotiation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-renegotiation"], _ = expandFirewallAccessProxy6ApiGatewaySslRenegotiation(d, i["ssl_renegotiation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_vpn_web_portal"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-vpn-web-portal"], _ = expandFirewallAccessProxy6ApiGatewaySslVpnWebPortal(d, i["ssl_vpn_web_portal"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["url-map"], _ = expandFirewallAccessProxy6ApiGatewayUrlMap(d, i["url_map"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["url-map-type"], _ = expandFirewallAccessProxy6ApiGatewayUrlMapType(d, i["url_map_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["virtual-host"], _ = expandFirewallAccessProxy6ApiGatewayVirtualHost(d, i["virtual_host"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxy6ApiGatewayApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGatewayH2Support(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayH3Support(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayHttpCookieAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayHttpCookieDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayHttpCookieDomainFromHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayHttpCookieGeneration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayHttpCookiePath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayHttpCookieShare(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayHttpsCookieSecure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayLdbMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayPersistence(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayQuic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ack-delay-exponent"], _ = expandFirewallAccessProxy6ApiGatewayQuicAckDelayExponent(d, i["ack_delay_exponent"], pre_append)
	}
	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["active-connection-id-limit"], _ = expandFirewallAccessProxy6ApiGatewayQuicActiveConnectionIdLimit(d, i["active_connection_id_limit"], pre_append)
	}
	pre_append = pre + ".0." + "active_migration"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["active-migration"], _ = expandFirewallAccessProxy6ApiGatewayQuicActiveMigration(d, i["active_migration"], pre_append)
	}
	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["grease-quic-bit"], _ = expandFirewallAccessProxy6ApiGatewayQuicGreaseQuicBit(d, i["grease_quic_bit"], pre_append)
	}
	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-ack-delay"], _ = expandFirewallAccessProxy6ApiGatewayQuicMaxAckDelay(d, i["max_ack_delay"], pre_append)
	}
	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-datagram-frame-size"], _ = expandFirewallAccessProxy6ApiGatewayQuicMaxDatagramFrameSize(d, i["max_datagram_frame_size"], pre_append)
	}
	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-idle-timeout"], _ = expandFirewallAccessProxy6ApiGatewayQuicMaxIdleTimeout(d, i["max_idle_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-udp-payload-size"], _ = expandFirewallAccessProxy6ApiGatewayQuicMaxUdpPayloadSize(d, i["max_udp_payload_size"], pre_append)
	}

	return result, nil
}

func expandFirewallAccessProxy6ApiGatewayQuicAckDelayExponent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayQuicActiveConnectionIdLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayQuicActiveMigration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayQuicGreaseQuicBit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayQuicMaxAckDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayQuicMaxDatagramFrameSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayQuicMaxIdleTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayQuicMaxUdpPayloadSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealservers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["addr-type"], _ = expandFirewallAccessProxy6ApiGatewayRealserversAddrType(d, i["addr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["address"], _ = expandFirewallAccessProxy6ApiGatewayRealserversAddress(d, i["address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["domain"], _ = expandFirewallAccessProxy6ApiGatewayRealserversDomain(d, i["domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_auth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["external-auth"], _ = expandFirewallAccessProxy6ApiGatewayRealserversExternalAuth(d, i["external_auth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check"], _ = expandFirewallAccessProxy6ApiGatewayRealserversHealthCheck(d, i["health_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check-proto"], _ = expandFirewallAccessProxy6ApiGatewayRealserversHealthCheckProto(d, i["health_check_proto"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["holddown-interval"], _ = expandFirewallAccessProxy6ApiGatewayRealserversHolddownInterval(d, i["holddown_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-host"], _ = expandFirewallAccessProxy6ApiGatewayRealserversHttpHost(d, i["http_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFirewallAccessProxy6ApiGatewayRealserversId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandFirewallAccessProxy6ApiGatewayRealserversIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mappedport"], _ = expandFirewallAccessProxy6ApiGatewayRealserversMappedport(d, i["mappedport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandFirewallAccessProxy6ApiGatewayRealserversPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_client_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssh-client-cert"], _ = expandFirewallAccessProxy6ApiGatewayRealserversSshClientCert(d, i["ssh_client_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssh-host-key"], _ = expandFirewallAccessProxy6ApiGatewayRealserversSshHostKey(d, i["ssh_host_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key_validation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssh-host-key-validation"], _ = expandFirewallAccessProxy6ApiGatewayRealserversSshHostKeyValidation(d, i["ssh_host_key_validation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandFirewallAccessProxy6ApiGatewayRealserversStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["translate-host"], _ = expandFirewallAccessProxy6ApiGatewayRealserversTranslateHost(d, i["translate_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tunnel_encryption"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tunnel-encryption"], _ = expandFirewallAccessProxy6ApiGatewayRealserversTunnelEncryption(d, i["tunnel_encryption"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandFirewallAccessProxy6ApiGatewayRealserversType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "verify_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["verify-cert"], _ = expandFirewallAccessProxy6ApiGatewayRealserversVerifyCert(d, i["verify_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandFirewallAccessProxy6ApiGatewayRealserversWeight(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversExternalAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversHealthCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversHealthCheckProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversHolddownInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversHttpHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversMappedport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversSshClientCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversSshHostKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversSshHostKeyValidation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversTranslateHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversTunnelEncryption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversVerifyCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySamlRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySamlServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGatewayService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySslAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySslCipherSuites(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["cipher"], _ = expandFirewallAccessProxy6ApiGatewaySslCipherSuitesCipher(d, i["cipher"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandFirewallAccessProxy6ApiGatewaySslCipherSuitesPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["versions"], _ = expandFirewallAccessProxy6ApiGatewaySslCipherSuitesVersions(d, i["versions"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxy6ApiGatewaySslCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySslCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySslCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGatewaySslDhBits(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySslMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySslMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySslRenegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySslVpnWebPortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGatewayUrlMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayUrlMapType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayVirtualHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGateway6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["application"], _ = expandFirewallAccessProxy6ApiGateway6Application(d, i["application"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h2_support"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["h2-support"], _ = expandFirewallAccessProxy6ApiGateway6H2Support(d, i["h2_support"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h3_support"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["h3-support"], _ = expandFirewallAccessProxy6ApiGateway6H3Support(d, i["h3_support"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-age"], _ = expandFirewallAccessProxy6ApiGateway6HttpCookieAge(d, i["http_cookie_age"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-domain"], _ = expandFirewallAccessProxy6ApiGateway6HttpCookieDomain(d, i["http_cookie_domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-domain-from-host"], _ = expandFirewallAccessProxy6ApiGateway6HttpCookieDomainFromHost(d, i["http_cookie_domain_from_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-generation"], _ = expandFirewallAccessProxy6ApiGateway6HttpCookieGeneration(d, i["http_cookie_generation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-path"], _ = expandFirewallAccessProxy6ApiGateway6HttpCookiePath(d, i["http_cookie_path"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-share"], _ = expandFirewallAccessProxy6ApiGateway6HttpCookieShare(d, i["http_cookie_share"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["https-cookie-secure"], _ = expandFirewallAccessProxy6ApiGateway6HttpsCookieSecure(d, i["https_cookie_secure"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFirewallAccessProxy6ApiGateway6Id(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ldb-method"], _ = expandFirewallAccessProxy6ApiGateway6LdbMethod(d, i["ldb_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["persistence"], _ = expandFirewallAccessProxy6ApiGateway6Persistence(d, i["persistence"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quic"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandFirewallAccessProxy6ApiGateway6Quic(d, i["quic"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["quic"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandFirewallAccessProxy6ApiGateway6Realservers(d, i["realservers"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["realservers"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_redirect"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["saml-redirect"], _ = expandFirewallAccessProxy6ApiGateway6SamlRedirect(d, i["saml_redirect"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["saml-server"], _ = expandFirewallAccessProxy6ApiGateway6SamlServer(d, i["saml_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service"], _ = expandFirewallAccessProxy6ApiGateway6Service(d, i["service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-algorithm"], _ = expandFirewallAccessProxy6ApiGateway6SslAlgorithm(d, i["ssl_algorithm"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandFirewallAccessProxy6ApiGateway6SslCipherSuites(d, i["ssl_cipher_suites"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["ssl-cipher-suites"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-dh-bits"], _ = expandFirewallAccessProxy6ApiGateway6SslDhBits(d, i["ssl_dh_bits"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-max-version"], _ = expandFirewallAccessProxy6ApiGateway6SslMaxVersion(d, i["ssl_max_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-min-version"], _ = expandFirewallAccessProxy6ApiGateway6SslMinVersion(d, i["ssl_min_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_renegotiation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-renegotiation"], _ = expandFirewallAccessProxy6ApiGateway6SslRenegotiation(d, i["ssl_renegotiation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_vpn_web_portal"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-vpn-web-portal"], _ = expandFirewallAccessProxy6ApiGateway6SslVpnWebPortal(d, i["ssl_vpn_web_portal"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["url-map"], _ = expandFirewallAccessProxy6ApiGateway6UrlMap(d, i["url_map"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["url-map-type"], _ = expandFirewallAccessProxy6ApiGateway6UrlMapType(d, i["url_map_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["virtual-host"], _ = expandFirewallAccessProxy6ApiGateway6VirtualHost(d, i["virtual_host"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxy6ApiGateway6Application(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGateway6H2Support(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6H3Support(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6HttpCookieAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6HttpCookieDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6HttpCookieDomainFromHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6HttpCookieGeneration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6HttpCookiePath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6HttpCookieShare(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6HttpsCookieSecure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6LdbMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6Persistence(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6Quic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ack-delay-exponent"], _ = expandFirewallAccessProxy6ApiGateway6QuicAckDelayExponent(d, i["ack_delay_exponent"], pre_append)
	}
	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["active-connection-id-limit"], _ = expandFirewallAccessProxy6ApiGateway6QuicActiveConnectionIdLimit(d, i["active_connection_id_limit"], pre_append)
	}
	pre_append = pre + ".0." + "active_migration"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["active-migration"], _ = expandFirewallAccessProxy6ApiGateway6QuicActiveMigration(d, i["active_migration"], pre_append)
	}
	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["grease-quic-bit"], _ = expandFirewallAccessProxy6ApiGateway6QuicGreaseQuicBit(d, i["grease_quic_bit"], pre_append)
	}
	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-ack-delay"], _ = expandFirewallAccessProxy6ApiGateway6QuicMaxAckDelay(d, i["max_ack_delay"], pre_append)
	}
	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-datagram-frame-size"], _ = expandFirewallAccessProxy6ApiGateway6QuicMaxDatagramFrameSize(d, i["max_datagram_frame_size"], pre_append)
	}
	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-idle-timeout"], _ = expandFirewallAccessProxy6ApiGateway6QuicMaxIdleTimeout(d, i["max_idle_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-udp-payload-size"], _ = expandFirewallAccessProxy6ApiGateway6QuicMaxUdpPayloadSize(d, i["max_udp_payload_size"], pre_append)
	}

	return result, nil
}

func expandFirewallAccessProxy6ApiGateway6QuicAckDelayExponent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6QuicActiveConnectionIdLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6QuicActiveMigration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6QuicGreaseQuicBit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6QuicMaxAckDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6QuicMaxDatagramFrameSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6QuicMaxIdleTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6QuicMaxUdpPayloadSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6Realservers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["addr-type"], _ = expandFirewallAccessProxy6ApiGateway6RealserversAddrType(d, i["addr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["address"], _ = expandFirewallAccessProxy6ApiGateway6RealserversAddress(d, i["address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["domain"], _ = expandFirewallAccessProxy6ApiGateway6RealserversDomain(d, i["domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_auth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["external-auth"], _ = expandFirewallAccessProxy6ApiGateway6RealserversExternalAuth(d, i["external_auth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check"], _ = expandFirewallAccessProxy6ApiGateway6RealserversHealthCheck(d, i["health_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check-proto"], _ = expandFirewallAccessProxy6ApiGateway6RealserversHealthCheckProto(d, i["health_check_proto"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["holddown-interval"], _ = expandFirewallAccessProxy6ApiGateway6RealserversHolddownInterval(d, i["holddown_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-host"], _ = expandFirewallAccessProxy6ApiGateway6RealserversHttpHost(d, i["http_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFirewallAccessProxy6ApiGateway6RealserversId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandFirewallAccessProxy6ApiGateway6RealserversIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mappedport"], _ = expandFirewallAccessProxy6ApiGateway6RealserversMappedport(d, i["mappedport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandFirewallAccessProxy6ApiGateway6RealserversPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_client_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssh-client-cert"], _ = expandFirewallAccessProxy6ApiGateway6RealserversSshClientCert(d, i["ssh_client_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssh-host-key"], _ = expandFirewallAccessProxy6ApiGateway6RealserversSshHostKey(d, i["ssh_host_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key_validation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssh-host-key-validation"], _ = expandFirewallAccessProxy6ApiGateway6RealserversSshHostKeyValidation(d, i["ssh_host_key_validation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandFirewallAccessProxy6ApiGateway6RealserversStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["translate-host"], _ = expandFirewallAccessProxy6ApiGateway6RealserversTranslateHost(d, i["translate_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tunnel_encryption"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tunnel-encryption"], _ = expandFirewallAccessProxy6ApiGateway6RealserversTunnelEncryption(d, i["tunnel_encryption"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandFirewallAccessProxy6ApiGateway6RealserversType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "verify_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["verify-cert"], _ = expandFirewallAccessProxy6ApiGateway6RealserversVerifyCert(d, i["verify_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandFirewallAccessProxy6ApiGateway6RealserversWeight(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversExternalAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversHealthCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversHealthCheckProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversHolddownInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversHttpHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversMappedport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversSshClientCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversSshHostKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversSshHostKeyValidation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversTranslateHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversTunnelEncryption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversVerifyCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6SamlRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6SamlServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGateway6Service(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6SslAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6SslCipherSuites(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["cipher"], _ = expandFirewallAccessProxy6ApiGateway6SslCipherSuitesCipher(d, i["cipher"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandFirewallAccessProxy6ApiGateway6SslCipherSuitesPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["versions"], _ = expandFirewallAccessProxy6ApiGateway6SslCipherSuitesVersions(d, i["versions"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxy6ApiGateway6SslCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6SslCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6SslCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGateway6SslDhBits(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6SslMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6SslMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6SslRenegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6SslVpnWebPortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGateway6UrlMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6UrlMapType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6VirtualHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6AuthPortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6AuthVirtualHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ClientCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6DecryptedTrafficMirror(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6EmptyCertAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6HttpSupportedMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6LogBlockedTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6SvrPoolMultiplex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6SvrPoolServerMaxConcurrentRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6SvrPoolServerMaxRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6SvrPoolTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6UserAgentDetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6Vip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectFirewallAccessProxy6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("add_vhost_domain_to_dnsdb"); ok || d.HasChange("add_vhost_domain_to_dnsdb") {
		t, err := expandFirewallAccessProxy6AddVhostDomainToDnsdb(d, v, "add_vhost_domain_to_dnsdb")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-vhost-domain-to-dnsdb"] = t
		}
	}

	if v, ok := d.GetOk("api_gateway"); ok || d.HasChange("api_gateway") {
		t, err := expandFirewallAccessProxy6ApiGateway(d, v, "api_gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["api-gateway"] = t
		}
	}

	if v, ok := d.GetOk("api_gateway6"); ok || d.HasChange("api_gateway6") {
		t, err := expandFirewallAccessProxy6ApiGateway6(d, v, "api_gateway6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["api-gateway6"] = t
		}
	}

	if v, ok := d.GetOk("auth_portal"); ok || d.HasChange("auth_portal") {
		t, err := expandFirewallAccessProxy6AuthPortal(d, v, "auth_portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-portal"] = t
		}
	}

	if v, ok := d.GetOk("auth_virtual_host"); ok || d.HasChange("auth_virtual_host") {
		t, err := expandFirewallAccessProxy6AuthVirtualHost(d, v, "auth_virtual_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-virtual-host"] = t
		}
	}

	if v, ok := d.GetOk("client_cert"); ok || d.HasChange("client_cert") {
		t, err := expandFirewallAccessProxy6ClientCert(d, v, "client_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert"] = t
		}
	}

	if v, ok := d.GetOk("decrypted_traffic_mirror"); ok || d.HasChange("decrypted_traffic_mirror") {
		t, err := expandFirewallAccessProxy6DecryptedTrafficMirror(d, v, "decrypted_traffic_mirror")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["decrypted-traffic-mirror"] = t
		}
	}

	if v, ok := d.GetOk("empty_cert_action"); ok || d.HasChange("empty_cert_action") {
		t, err := expandFirewallAccessProxy6EmptyCertAction(d, v, "empty_cert_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["empty-cert-action"] = t
		}
	}

	if v, ok := d.GetOk("http_supported_max_version"); ok || d.HasChange("http_supported_max_version") {
		t, err := expandFirewallAccessProxy6HttpSupportedMaxVersion(d, v, "http_supported_max_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-supported-max-version"] = t
		}
	}

	if v, ok := d.GetOk("log_blocked_traffic"); ok || d.HasChange("log_blocked_traffic") {
		t, err := expandFirewallAccessProxy6LogBlockedTraffic(d, v, "log_blocked_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-blocked-traffic"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallAccessProxy6Name(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("svr_pool_multiplex"); ok || d.HasChange("svr_pool_multiplex") {
		t, err := expandFirewallAccessProxy6SvrPoolMultiplex(d, v, "svr_pool_multiplex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["svr-pool-multiplex"] = t
		}
	}

	if v, ok := d.GetOk("svr_pool_server_max_concurrent_request"); ok || d.HasChange("svr_pool_server_max_concurrent_request") {
		t, err := expandFirewallAccessProxy6SvrPoolServerMaxConcurrentRequest(d, v, "svr_pool_server_max_concurrent_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["svr-pool-server-max-concurrent-request"] = t
		}
	}

	if v, ok := d.GetOk("svr_pool_server_max_request"); ok || d.HasChange("svr_pool_server_max_request") {
		t, err := expandFirewallAccessProxy6SvrPoolServerMaxRequest(d, v, "svr_pool_server_max_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["svr-pool-server-max-request"] = t
		}
	}

	if v, ok := d.GetOk("svr_pool_ttl"); ok || d.HasChange("svr_pool_ttl") {
		t, err := expandFirewallAccessProxy6SvrPoolTtl(d, v, "svr_pool_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["svr-pool-ttl"] = t
		}
	}

	if v, ok := d.GetOk("user_agent_detect"); ok || d.HasChange("user_agent_detect") {
		t, err := expandFirewallAccessProxy6UserAgentDetect(d, v, "user_agent_detect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-agent-detect"] = t
		}
	}

	if v, ok := d.GetOk("vip"); ok || d.HasChange("vip") {
		t, err := expandFirewallAccessProxy6Vip(d, v, "vip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vip"] = t
		}
	}

	return &obj, nil
}
