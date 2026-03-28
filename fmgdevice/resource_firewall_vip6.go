// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure virtual IP for IPv6.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallVip6() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallVip6Create,
		Read:   resourceFirewallVip6Read,
		Update: resourceFirewallVip6Update,
		Delete: resourceFirewallVip6Delete,

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
			"add_nat64_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"arp_reply": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"color": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_mapping": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"_scope": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"vdom": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"add_nat64_route": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"arp_reply": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"client_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"color": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"embedded_ipv4_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"empty_cert_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"extip": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"extport": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"h2_support": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"h3_support": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						},
						"http_ip_header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"http_ip_header_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"http_multiplex": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"http_redirect": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"https_cookie_secure": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ipv4_mappedip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipv4_mappedport": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ldb_method": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mappedip": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"mappedport": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"max_embryonic_connections": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"monitor": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"nat_source_vip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"nat64": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"nat66": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ndp_reply": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"outlook_web_access": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"persistence": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"portforward": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"realservers": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"client_ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"healthcheck": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"holddown_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
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
									},
									"max_connections": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"monitor": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"port": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"translate_host": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"verify_cert": &schema.Schema{
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
						"server_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"src_filter": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"src_vip_filter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_accept_ffdhe_groups": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_algorithm": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						},
						"ssl_client_rekey_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ssl_client_renegotiation": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_client_session_state_max": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ssl_client_session_state_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ssl_client_session_state_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_dh_bits": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_hpkp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_hpkp_age": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
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
						},
						"ssl_hsts_age": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ssl_hsts_include_subdomains": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_http_location_conversion": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_http_match_host": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_max_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_min_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_pfs": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_send_empty_frags": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_server_algorithm": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_server_max_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_server_min_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_server_renegotiation": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_server_session_state_max": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ssl_server_session_state_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ssl_server_session_state_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"user_agent_detect": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"uuid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"weblogic_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"websphere_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"embedded_ipv4_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"empty_cert_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extip": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"extport": &schema.Schema{
				Type:     schema.TypeString,
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
			"http_ip_header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_ip_header_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_multiplex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_redirect": &schema.Schema{
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
			},
			"ipv4_mappedip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_mappedport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldb_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mappedip": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"mappedport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_embryonic_connections": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"monitor": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nat_source_vip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat64": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat66": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ndp_reply": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"outlook_web_access": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"persistence": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"portforward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
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
						"client_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"healthcheck": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"holddown_interval": &schema.Schema{
							Type:     schema.TypeInt,
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
						},
						"max_connections": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"monitor": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
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
			"server_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"src_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"src_vip_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_accept_ffdhe_groups": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"ssl_client_session_state_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ssl_client_session_state_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ssl_client_session_state_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_dh_bits": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_hpkp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_hpkp_age": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"ssl_max_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_min_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_pfs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_send_empty_frags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"ssl_server_session_state_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ssl_server_session_state_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_agent_detect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weblogic_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"websphere_server": &schema.Schema{
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

func resourceFirewallVip6Create(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallVip6(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallVip6 resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallVip6(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallVip6(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallVip6 resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallVip6(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallVip6 resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallVip6Read(d, m)
}

func resourceFirewallVip6Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallVip6(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVip6 resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallVip6(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVip6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallVip6Read(d, m)
}

func resourceFirewallVip6Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallVip6(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallVip6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallVip6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallVip6(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallVip6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallVip6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallVip6 resource from API: %v", err)
	}
	return nil
}

func flattenFirewallVip6AddNat64Route(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6ArpReply(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6ClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6Color(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6Comment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := i["_scope"]; ok {
			v := flattenFirewallVip6DynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_nat64_route"
		if _, ok := i["add-nat64-route"]; ok {
			v := flattenFirewallVip6DynamicMappingAddNat64Route(i["add-nat64-route"], d, pre_append)
			tmp["add_nat64_route"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-AddNat64Route")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "arp_reply"
		if _, ok := i["arp-reply"]; ok {
			v := flattenFirewallVip6DynamicMappingArpReply(i["arp-reply"], d, pre_append)
			tmp["arp_reply"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-ArpReply")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_cert"
		if _, ok := i["client-cert"]; ok {
			v := flattenFirewallVip6DynamicMappingClientCert(i["client-cert"], d, pre_append)
			tmp["client_cert"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-ClientCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color"
		if _, ok := i["color"]; ok {
			v := flattenFirewallVip6DynamicMappingColor(i["color"], d, pre_append)
			tmp["color"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-Color")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenFirewallVip6DynamicMappingComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "embedded_ipv4_address"
		if _, ok := i["embedded-ipv4-address"]; ok {
			v := flattenFirewallVip6DynamicMappingEmbeddedIpv4Address(i["embedded-ipv4-address"], d, pre_append)
			tmp["embedded_ipv4_address"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-EmbeddedIpv4Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "empty_cert_action"
		if _, ok := i["empty-cert-action"]; ok {
			v := flattenFirewallVip6DynamicMappingEmptyCertAction(i["empty-cert-action"], d, pre_append)
			tmp["empty_cert_action"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-EmptyCertAction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extip"
		if _, ok := i["extip"]; ok {
			v := flattenFirewallVip6DynamicMappingExtip(i["extip"], d, pre_append)
			tmp["extip"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-Extip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extport"
		if _, ok := i["extport"]; ok {
			v := flattenFirewallVip6DynamicMappingExtport(i["extport"], d, pre_append)
			tmp["extport"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-Extport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h2_support"
		if _, ok := i["h2-support"]; ok {
			v := flattenFirewallVip6DynamicMappingH2Support(i["h2-support"], d, pre_append)
			tmp["h2_support"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-H2Support")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h3_support"
		if _, ok := i["h3-support"]; ok {
			v := flattenFirewallVip6DynamicMappingH3Support(i["h3-support"], d, pre_append)
			tmp["h3_support"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-H3Support")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := i["http-cookie-age"]; ok {
			v := flattenFirewallVip6DynamicMappingHttpCookieAge(i["http-cookie-age"], d, pre_append)
			tmp["http_cookie_age"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-HttpCookieAge")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := i["http-cookie-domain"]; ok {
			v := flattenFirewallVip6DynamicMappingHttpCookieDomain(i["http-cookie-domain"], d, pre_append)
			tmp["http_cookie_domain"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-HttpCookieDomain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := i["http-cookie-domain-from-host"]; ok {
			v := flattenFirewallVip6DynamicMappingHttpCookieDomainFromHost(i["http-cookie-domain-from-host"], d, pre_append)
			tmp["http_cookie_domain_from_host"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-HttpCookieDomainFromHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := i["http-cookie-generation"]; ok {
			v := flattenFirewallVip6DynamicMappingHttpCookieGeneration(i["http-cookie-generation"], d, pre_append)
			tmp["http_cookie_generation"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-HttpCookieGeneration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := i["http-cookie-path"]; ok {
			v := flattenFirewallVip6DynamicMappingHttpCookiePath(i["http-cookie-path"], d, pre_append)
			tmp["http_cookie_path"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-HttpCookiePath")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := i["http-cookie-share"]; ok {
			v := flattenFirewallVip6DynamicMappingHttpCookieShare(i["http-cookie-share"], d, pre_append)
			tmp["http_cookie_share"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-HttpCookieShare")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_ip_header"
		if _, ok := i["http-ip-header"]; ok {
			v := flattenFirewallVip6DynamicMappingHttpIpHeader(i["http-ip-header"], d, pre_append)
			tmp["http_ip_header"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-HttpIpHeader")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_ip_header_name"
		if _, ok := i["http-ip-header-name"]; ok {
			v := flattenFirewallVip6DynamicMappingHttpIpHeaderName(i["http-ip-header-name"], d, pre_append)
			tmp["http_ip_header_name"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-HttpIpHeaderName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_multiplex"
		if _, ok := i["http-multiplex"]; ok {
			v := flattenFirewallVip6DynamicMappingHttpMultiplex(i["http-multiplex"], d, pre_append)
			tmp["http_multiplex"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-HttpMultiplex")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_redirect"
		if _, ok := i["http-redirect"]; ok {
			v := flattenFirewallVip6DynamicMappingHttpRedirect(i["http-redirect"], d, pre_append)
			tmp["http_redirect"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-HttpRedirect")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := i["https-cookie-secure"]; ok {
			v := flattenFirewallVip6DynamicMappingHttpsCookieSecure(i["https-cookie-secure"], d, pre_append)
			tmp["https_cookie_secure"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-HttpsCookieSecure")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFirewallVip6DynamicMappingId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv4_mappedip"
		if _, ok := i["ipv4-mappedip"]; ok {
			v := flattenFirewallVip6DynamicMappingIpv4Mappedip(i["ipv4-mappedip"], d, pre_append)
			tmp["ipv4_mappedip"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-Ipv4Mappedip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv4_mappedport"
		if _, ok := i["ipv4-mappedport"]; ok {
			v := flattenFirewallVip6DynamicMappingIpv4Mappedport(i["ipv4-mappedport"], d, pre_append)
			tmp["ipv4_mappedport"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-Ipv4Mappedport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := i["ldb-method"]; ok {
			v := flattenFirewallVip6DynamicMappingLdbMethod(i["ldb-method"], d, pre_append)
			tmp["ldb_method"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-LdbMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedip"
		if _, ok := i["mappedip"]; ok {
			v := flattenFirewallVip6DynamicMappingMappedip(i["mappedip"], d, pre_append)
			tmp["mappedip"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-Mappedip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := i["mappedport"]; ok {
			v := flattenFirewallVip6DynamicMappingMappedport(i["mappedport"], d, pre_append)
			tmp["mappedport"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-Mappedport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_embryonic_connections"
		if _, ok := i["max-embryonic-connections"]; ok {
			v := flattenFirewallVip6DynamicMappingMaxEmbryonicConnections(i["max-embryonic-connections"], d, pre_append)
			tmp["max_embryonic_connections"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-MaxEmbryonicConnections")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := i["monitor"]; ok {
			v := flattenFirewallVip6DynamicMappingMonitor(i["monitor"], d, pre_append)
			tmp["monitor"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-Monitor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat_source_vip"
		if _, ok := i["nat-source-vip"]; ok {
			v := flattenFirewallVip6DynamicMappingNatSourceVip(i["nat-source-vip"], d, pre_append)
			tmp["nat_source_vip"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-NatSourceVip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat64"
		if _, ok := i["nat64"]; ok {
			v := flattenFirewallVip6DynamicMappingNat64(i["nat64"], d, pre_append)
			tmp["nat64"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-Nat64")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat66"
		if _, ok := i["nat66"]; ok {
			v := flattenFirewallVip6DynamicMappingNat66(i["nat66"], d, pre_append)
			tmp["nat66"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-Nat66")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ndp_reply"
		if _, ok := i["ndp-reply"]; ok {
			v := flattenFirewallVip6DynamicMappingNdpReply(i["ndp-reply"], d, pre_append)
			tmp["ndp_reply"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-NdpReply")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "outlook_web_access"
		if _, ok := i["outlook-web-access"]; ok {
			v := flattenFirewallVip6DynamicMappingOutlookWebAccess(i["outlook-web-access"], d, pre_append)
			tmp["outlook_web_access"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-OutlookWebAccess")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := i["persistence"]; ok {
			v := flattenFirewallVip6DynamicMappingPersistence(i["persistence"], d, pre_append)
			tmp["persistence"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-Persistence")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "portforward"
		if _, ok := i["portforward"]; ok {
			v := flattenFirewallVip6DynamicMappingPortforward(i["portforward"], d, pre_append)
			tmp["portforward"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-Portforward")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenFirewallVip6DynamicMappingProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := i["realservers"]; ok {
			v := flattenFirewallVip6DynamicMappingRealservers(i["realservers"], d, pre_append)
			tmp["realservers"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-Realservers")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_type"
		if _, ok := i["server-type"]; ok {
			v := flattenFirewallVip6DynamicMappingServerType(i["server-type"], d, pre_append)
			tmp["server_type"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-ServerType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_filter"
		if _, ok := i["src-filter"]; ok {
			v := flattenFirewallVip6DynamicMappingSrcFilter(i["src-filter"], d, pre_append)
			tmp["src_filter"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SrcFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_vip_filter"
		if _, ok := i["src-vip-filter"]; ok {
			v := flattenFirewallVip6DynamicMappingSrcVipFilter(i["src-vip-filter"], d, pre_append)
			tmp["src_vip_filter"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SrcVipFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_accept_ffdhe_groups"
		if _, ok := i["ssl-accept-ffdhe-groups"]; ok {
			v := flattenFirewallVip6DynamicMappingSslAcceptFfdheGroups(i["ssl-accept-ffdhe-groups"], d, pre_append)
			tmp["ssl_accept_ffdhe_groups"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslAcceptFfdheGroups")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := i["ssl-algorithm"]; ok {
			v := flattenFirewallVip6DynamicMappingSslAlgorithm(i["ssl-algorithm"], d, pre_append)
			tmp["ssl_algorithm"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslAlgorithm")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_certificate"
		if _, ok := i["ssl-certificate"]; ok {
			v := flattenFirewallVip6DynamicMappingSslCertificate(i["ssl-certificate"], d, pre_append)
			tmp["ssl_certificate"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslCertificate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := i["ssl-cipher-suites"]; ok {
			v := flattenFirewallVip6DynamicMappingSslCipherSuites(i["ssl-cipher-suites"], d, pre_append)
			tmp["ssl_cipher_suites"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslCipherSuites")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_fallback"
		if _, ok := i["ssl-client-fallback"]; ok {
			v := flattenFirewallVip6DynamicMappingSslClientFallback(i["ssl-client-fallback"], d, pre_append)
			tmp["ssl_client_fallback"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslClientFallback")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_rekey_count"
		if _, ok := i["ssl-client-rekey-count"]; ok {
			v := flattenFirewallVip6DynamicMappingSslClientRekeyCount(i["ssl-client-rekey-count"], d, pre_append)
			tmp["ssl_client_rekey_count"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslClientRekeyCount")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_renegotiation"
		if _, ok := i["ssl-client-renegotiation"]; ok {
			v := flattenFirewallVip6DynamicMappingSslClientRenegotiation(i["ssl-client-renegotiation"], d, pre_append)
			tmp["ssl_client_renegotiation"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslClientRenegotiation")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_session_state_max"
		if _, ok := i["ssl-client-session-state-max"]; ok {
			v := flattenFirewallVip6DynamicMappingSslClientSessionStateMax(i["ssl-client-session-state-max"], d, pre_append)
			tmp["ssl_client_session_state_max"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslClientSessionStateMax")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_session_state_timeout"
		if _, ok := i["ssl-client-session-state-timeout"]; ok {
			v := flattenFirewallVip6DynamicMappingSslClientSessionStateTimeout(i["ssl-client-session-state-timeout"], d, pre_append)
			tmp["ssl_client_session_state_timeout"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslClientSessionStateTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_session_state_type"
		if _, ok := i["ssl-client-session-state-type"]; ok {
			v := flattenFirewallVip6DynamicMappingSslClientSessionStateType(i["ssl-client-session-state-type"], d, pre_append)
			tmp["ssl_client_session_state_type"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslClientSessionStateType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := i["ssl-dh-bits"]; ok {
			v := flattenFirewallVip6DynamicMappingSslDhBits(i["ssl-dh-bits"], d, pre_append)
			tmp["ssl_dh_bits"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslDhBits")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp"
		if _, ok := i["ssl-hpkp"]; ok {
			v := flattenFirewallVip6DynamicMappingSslHpkp(i["ssl-hpkp"], d, pre_append)
			tmp["ssl_hpkp"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslHpkp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_age"
		if _, ok := i["ssl-hpkp-age"]; ok {
			v := flattenFirewallVip6DynamicMappingSslHpkpAge(i["ssl-hpkp-age"], d, pre_append)
			tmp["ssl_hpkp_age"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslHpkpAge")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_backup"
		if _, ok := i["ssl-hpkp-backup"]; ok {
			v := flattenFirewallVip6DynamicMappingSslHpkpBackup(i["ssl-hpkp-backup"], d, pre_append)
			tmp["ssl_hpkp_backup"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslHpkpBackup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_include_subdomains"
		if _, ok := i["ssl-hpkp-include-subdomains"]; ok {
			v := flattenFirewallVip6DynamicMappingSslHpkpIncludeSubdomains(i["ssl-hpkp-include-subdomains"], d, pre_append)
			tmp["ssl_hpkp_include_subdomains"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslHpkpIncludeSubdomains")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_primary"
		if _, ok := i["ssl-hpkp-primary"]; ok {
			v := flattenFirewallVip6DynamicMappingSslHpkpPrimary(i["ssl-hpkp-primary"], d, pre_append)
			tmp["ssl_hpkp_primary"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslHpkpPrimary")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_report_uri"
		if _, ok := i["ssl-hpkp-report-uri"]; ok {
			v := flattenFirewallVip6DynamicMappingSslHpkpReportUri(i["ssl-hpkp-report-uri"], d, pre_append)
			tmp["ssl_hpkp_report_uri"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslHpkpReportUri")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hsts"
		if _, ok := i["ssl-hsts"]; ok {
			v := flattenFirewallVip6DynamicMappingSslHsts(i["ssl-hsts"], d, pre_append)
			tmp["ssl_hsts"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslHsts")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hsts_age"
		if _, ok := i["ssl-hsts-age"]; ok {
			v := flattenFirewallVip6DynamicMappingSslHstsAge(i["ssl-hsts-age"], d, pre_append)
			tmp["ssl_hsts_age"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslHstsAge")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hsts_include_subdomains"
		if _, ok := i["ssl-hsts-include-subdomains"]; ok {
			v := flattenFirewallVip6DynamicMappingSslHstsIncludeSubdomains(i["ssl-hsts-include-subdomains"], d, pre_append)
			tmp["ssl_hsts_include_subdomains"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslHstsIncludeSubdomains")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_http_location_conversion"
		if _, ok := i["ssl-http-location-conversion"]; ok {
			v := flattenFirewallVip6DynamicMappingSslHttpLocationConversion(i["ssl-http-location-conversion"], d, pre_append)
			tmp["ssl_http_location_conversion"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslHttpLocationConversion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_http_match_host"
		if _, ok := i["ssl-http-match-host"]; ok {
			v := flattenFirewallVip6DynamicMappingSslHttpMatchHost(i["ssl-http-match-host"], d, pre_append)
			tmp["ssl_http_match_host"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslHttpMatchHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := i["ssl-max-version"]; ok {
			v := flattenFirewallVip6DynamicMappingSslMaxVersion(i["ssl-max-version"], d, pre_append)
			tmp["ssl_max_version"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslMaxVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := i["ssl-min-version"]; ok {
			v := flattenFirewallVip6DynamicMappingSslMinVersion(i["ssl-min-version"], d, pre_append)
			tmp["ssl_min_version"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslMinVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_mode"
		if _, ok := i["ssl-mode"]; ok {
			v := flattenFirewallVip6DynamicMappingSslMode(i["ssl-mode"], d, pre_append)
			tmp["ssl_mode"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_pfs"
		if _, ok := i["ssl-pfs"]; ok {
			v := flattenFirewallVip6DynamicMappingSslPfs(i["ssl-pfs"], d, pre_append)
			tmp["ssl_pfs"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslPfs")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_send_empty_frags"
		if _, ok := i["ssl-send-empty-frags"]; ok {
			v := flattenFirewallVip6DynamicMappingSslSendEmptyFrags(i["ssl-send-empty-frags"], d, pre_append)
			tmp["ssl_send_empty_frags"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslSendEmptyFrags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_algorithm"
		if _, ok := i["ssl-server-algorithm"]; ok {
			v := flattenFirewallVip6DynamicMappingSslServerAlgorithm(i["ssl-server-algorithm"], d, pre_append)
			tmp["ssl_server_algorithm"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslServerAlgorithm")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_max_version"
		if _, ok := i["ssl-server-max-version"]; ok {
			v := flattenFirewallVip6DynamicMappingSslServerMaxVersion(i["ssl-server-max-version"], d, pre_append)
			tmp["ssl_server_max_version"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslServerMaxVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_min_version"
		if _, ok := i["ssl-server-min-version"]; ok {
			v := flattenFirewallVip6DynamicMappingSslServerMinVersion(i["ssl-server-min-version"], d, pre_append)
			tmp["ssl_server_min_version"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslServerMinVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_renegotiation"
		if _, ok := i["ssl-server-renegotiation"]; ok {
			v := flattenFirewallVip6DynamicMappingSslServerRenegotiation(i["ssl-server-renegotiation"], d, pre_append)
			tmp["ssl_server_renegotiation"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslServerRenegotiation")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_session_state_max"
		if _, ok := i["ssl-server-session-state-max"]; ok {
			v := flattenFirewallVip6DynamicMappingSslServerSessionStateMax(i["ssl-server-session-state-max"], d, pre_append)
			tmp["ssl_server_session_state_max"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslServerSessionStateMax")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_session_state_timeout"
		if _, ok := i["ssl-server-session-state-timeout"]; ok {
			v := flattenFirewallVip6DynamicMappingSslServerSessionStateTimeout(i["ssl-server-session-state-timeout"], d, pre_append)
			tmp["ssl_server_session_state_timeout"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslServerSessionStateTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_session_state_type"
		if _, ok := i["ssl-server-session-state-type"]; ok {
			v := flattenFirewallVip6DynamicMappingSslServerSessionStateType(i["ssl-server-session-state-type"], d, pre_append)
			tmp["ssl_server_session_state_type"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-SslServerSessionStateType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenFirewallVip6DynamicMappingType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_agent_detect"
		if _, ok := i["user-agent-detect"]; ok {
			v := flattenFirewallVip6DynamicMappingUserAgentDetect(i["user-agent-detect"], d, pre_append)
			tmp["user_agent_detect"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-UserAgentDetect")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uuid"
		if _, ok := i["uuid"]; ok {
			v := flattenFirewallVip6DynamicMappingUuid(i["uuid"], d, pre_append)
			tmp["uuid"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-Uuid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weblogic_server"
		if _, ok := i["weblogic-server"]; ok {
			v := flattenFirewallVip6DynamicMappingWeblogicServer(i["weblogic-server"], d, pre_append)
			tmp["weblogic_server"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-WeblogicServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "websphere_server"
		if _, ok := i["websphere-server"]; ok {
			v := flattenFirewallVip6DynamicMappingWebsphereServer(i["websphere-server"], d, pre_append)
			tmp["websphere_server"] = fortiAPISubPartPatch(v, "FirewallVip6-DynamicMapping-WebsphereServer")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallVip6DynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenFirewallVip6DynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "FirewallVip6DynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenFirewallVip6DynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "FirewallVip6DynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallVip6DynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingAddNat64Route(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingArpReply(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingEmbeddedIpv4Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingEmptyCertAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingExtip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVip6DynamicMappingExtport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingH2Support(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingH3Support(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingHttpCookieAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingHttpCookieDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingHttpCookieDomainFromHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingHttpCookieGeneration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingHttpCookiePath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingHttpCookieShare(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingHttpIpHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingHttpIpHeaderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingHttpMultiplex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingHttpRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingHttpsCookieSecure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingIpv4Mappedip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingIpv4Mappedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingLdbMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingMappedip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVip6DynamicMappingMappedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingMaxEmbryonicConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVip6DynamicMappingNatSourceVip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingNat64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingNat66(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingNdpReply(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingOutlookWebAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingPersistence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingPortforward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingRealservers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_ip"
		if _, ok := i["client-ip"]; ok {
			v := flattenFirewallVip6DynamicMappingRealserversClientIp(i["client-ip"], d, pre_append)
			tmp["client_ip"] = fortiAPISubPartPatch(v, "FirewallVip6DynamicMapping-Realservers-ClientIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "healthcheck"
		if _, ok := i["healthcheck"]; ok {
			v := flattenFirewallVip6DynamicMappingRealserversHealthcheck(i["healthcheck"], d, pre_append)
			tmp["healthcheck"] = fortiAPISubPartPatch(v, "FirewallVip6DynamicMapping-Realservers-Healthcheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := i["holddown-interval"]; ok {
			v := flattenFirewallVip6DynamicMappingRealserversHolddownInterval(i["holddown-interval"], d, pre_append)
			tmp["holddown_interval"] = fortiAPISubPartPatch(v, "FirewallVip6DynamicMapping-Realservers-HolddownInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := i["http-host"]; ok {
			v := flattenFirewallVip6DynamicMappingRealserversHttpHost(i["http-host"], d, pre_append)
			tmp["http_host"] = fortiAPISubPartPatch(v, "FirewallVip6DynamicMapping-Realservers-HttpHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFirewallVip6DynamicMappingRealserversId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FirewallVip6DynamicMapping-Realservers-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenFirewallVip6DynamicMappingRealserversIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "FirewallVip6DynamicMapping-Realservers-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_connections"
		if _, ok := i["max-connections"]; ok {
			v := flattenFirewallVip6DynamicMappingRealserversMaxConnections(i["max-connections"], d, pre_append)
			tmp["max_connections"] = fortiAPISubPartPatch(v, "FirewallVip6DynamicMapping-Realservers-MaxConnections")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := i["monitor"]; ok {
			v := flattenFirewallVip6DynamicMappingRealserversMonitor(i["monitor"], d, pre_append)
			tmp["monitor"] = fortiAPISubPartPatch(v, "FirewallVip6DynamicMapping-Realservers-Monitor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenFirewallVip6DynamicMappingRealserversPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "FirewallVip6DynamicMapping-Realservers-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenFirewallVip6DynamicMappingRealserversStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "FirewallVip6DynamicMapping-Realservers-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := i["translate-host"]; ok {
			v := flattenFirewallVip6DynamicMappingRealserversTranslateHost(i["translate-host"], d, pre_append)
			tmp["translate_host"] = fortiAPISubPartPatch(v, "FirewallVip6DynamicMapping-Realservers-TranslateHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "verify_cert"
		if _, ok := i["verify-cert"]; ok {
			v := flattenFirewallVip6DynamicMappingRealserversVerifyCert(i["verify-cert"], d, pre_append)
			tmp["verify_cert"] = fortiAPISubPartPatch(v, "FirewallVip6DynamicMapping-Realservers-VerifyCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenFirewallVip6DynamicMappingRealserversWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "FirewallVip6DynamicMapping-Realservers-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallVip6DynamicMappingRealserversClientIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingRealserversHealthcheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingRealserversHolddownInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingRealserversHttpHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingRealserversId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingRealserversIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingRealserversMaxConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingRealserversMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVip6DynamicMappingRealserversPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingRealserversStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingRealserversTranslateHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingRealserversVerifyCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingRealserversWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSrcFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVip6DynamicMappingSrcVipFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslAcceptFfdheGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVip6DynamicMappingSslCipherSuites(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallVip6DynamicMappingSslCipherSuitesCipher(i["cipher"], d, pre_append)
			tmp["cipher"] = fortiAPISubPartPatch(v, "FirewallVip6DynamicMapping-SslCipherSuites-Cipher")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenFirewallVip6DynamicMappingSslCipherSuitesPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "FirewallVip6DynamicMapping-SslCipherSuites-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {
			v := flattenFirewallVip6DynamicMappingSslCipherSuitesVersions(i["versions"], d, pre_append)
			tmp["versions"] = fortiAPISubPartPatch(v, "FirewallVip6DynamicMapping-SslCipherSuites-Versions")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallVip6DynamicMappingSslCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVip6DynamicMappingSslClientFallback(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslClientRekeyCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslClientRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslClientSessionStateMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslClientSessionStateTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslClientSessionStateType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslDhBits(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslHpkp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslHpkpAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslHpkpBackup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVip6DynamicMappingSslHpkpIncludeSubdomains(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslHpkpPrimary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVip6DynamicMappingSslHpkpReportUri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslHsts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslHstsAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslHstsIncludeSubdomains(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslHttpLocationConversion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslHttpMatchHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslPfs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslSendEmptyFrags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslServerAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslServerMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslServerMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslServerRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslServerSessionStateMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslServerSessionStateTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingSslServerSessionStateType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingUserAgentDetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingWeblogicServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6DynamicMappingWebsphereServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6EmbeddedIpv4Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6EmptyCertAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6Extip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVip6Extport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6H2Support(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6H3Support(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6HttpCookieAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6HttpCookieDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6HttpCookieDomainFromHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6HttpCookieGeneration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6HttpCookiePath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6HttpCookieShare(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6HttpIpHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6HttpIpHeaderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6HttpMultiplex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6HttpRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6HttpsCookieSecure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6Ipv4Mappedip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6Ipv4Mappedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6LdbMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6Mappedip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVip6Mappedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6MaxEmbryonicConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6Monitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVip6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6NatSourceVip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6Nat64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6Nat66(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6NdpReply(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6OutlookWebAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6Persistence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6Portforward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6Protocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6Quic(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := i["ack-delay-exponent"]; ok {
		result["ack_delay_exponent"] = flattenFirewallVip6QuicAckDelayExponent(i["ack-delay-exponent"], d, pre_append)
	}

	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := i["active-connection-id-limit"]; ok {
		result["active_connection_id_limit"] = flattenFirewallVip6QuicActiveConnectionIdLimit(i["active-connection-id-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "active_migration"
	if _, ok := i["active-migration"]; ok {
		result["active_migration"] = flattenFirewallVip6QuicActiveMigration(i["active-migration"], d, pre_append)
	}

	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := i["grease-quic-bit"]; ok {
		result["grease_quic_bit"] = flattenFirewallVip6QuicGreaseQuicBit(i["grease-quic-bit"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := i["max-ack-delay"]; ok {
		result["max_ack_delay"] = flattenFirewallVip6QuicMaxAckDelay(i["max-ack-delay"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := i["max-datagram-frame-size"]; ok {
		result["max_datagram_frame_size"] = flattenFirewallVip6QuicMaxDatagramFrameSize(i["max-datagram-frame-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := i["max-idle-timeout"]; ok {
		result["max_idle_timeout"] = flattenFirewallVip6QuicMaxIdleTimeout(i["max-idle-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := i["max-udp-payload-size"]; ok {
		result["max_udp_payload_size"] = flattenFirewallVip6QuicMaxUdpPayloadSize(i["max-udp-payload-size"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallVip6QuicAckDelayExponent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6QuicActiveConnectionIdLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6QuicActiveMigration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6QuicGreaseQuicBit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6QuicMaxAckDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6QuicMaxDatagramFrameSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6QuicMaxIdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6QuicMaxUdpPayloadSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6Realservers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_ip"
		if _, ok := i["client-ip"]; ok {
			v := flattenFirewallVip6RealserversClientIp(i["client-ip"], d, pre_append)
			tmp["client_ip"] = fortiAPISubPartPatch(v, "FirewallVip6-Realservers-ClientIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "healthcheck"
		if _, ok := i["healthcheck"]; ok {
			v := flattenFirewallVip6RealserversHealthcheck(i["healthcheck"], d, pre_append)
			tmp["healthcheck"] = fortiAPISubPartPatch(v, "FirewallVip6-Realservers-Healthcheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := i["holddown-interval"]; ok {
			v := flattenFirewallVip6RealserversHolddownInterval(i["holddown-interval"], d, pre_append)
			tmp["holddown_interval"] = fortiAPISubPartPatch(v, "FirewallVip6-Realservers-HolddownInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := i["http-host"]; ok {
			v := flattenFirewallVip6RealserversHttpHost(i["http-host"], d, pre_append)
			tmp["http_host"] = fortiAPISubPartPatch(v, "FirewallVip6-Realservers-HttpHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFirewallVip6RealserversId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FirewallVip6-Realservers-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenFirewallVip6RealserversIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "FirewallVip6-Realservers-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_connections"
		if _, ok := i["max-connections"]; ok {
			v := flattenFirewallVip6RealserversMaxConnections(i["max-connections"], d, pre_append)
			tmp["max_connections"] = fortiAPISubPartPatch(v, "FirewallVip6-Realservers-MaxConnections")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := i["monitor"]; ok {
			v := flattenFirewallVip6RealserversMonitor(i["monitor"], d, pre_append)
			tmp["monitor"] = fortiAPISubPartPatch(v, "FirewallVip6-Realservers-Monitor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenFirewallVip6RealserversPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "FirewallVip6-Realservers-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenFirewallVip6RealserversStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "FirewallVip6-Realservers-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := i["translate-host"]; ok {
			v := flattenFirewallVip6RealserversTranslateHost(i["translate-host"], d, pre_append)
			tmp["translate_host"] = fortiAPISubPartPatch(v, "FirewallVip6-Realservers-TranslateHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "verify_cert"
		if _, ok := i["verify-cert"]; ok {
			v := flattenFirewallVip6RealserversVerifyCert(i["verify-cert"], d, pre_append)
			tmp["verify_cert"] = fortiAPISubPartPatch(v, "FirewallVip6-Realservers-VerifyCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenFirewallVip6RealserversWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "FirewallVip6-Realservers-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallVip6RealserversClientIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6RealserversHealthcheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6RealserversHolddownInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6RealserversHttpHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6RealserversId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6RealserversIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6RealserversMaxConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6RealserversMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVip6RealserversPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6RealserversStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6RealserversTranslateHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6RealserversVerifyCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6RealserversWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6ServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SrcFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVip6SrcVipFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslAcceptFfdheGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVip6SslCipherSuites(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallVip6SslCipherSuitesCipher(i["cipher"], d, pre_append)
			tmp["cipher"] = fortiAPISubPartPatch(v, "FirewallVip6-SslCipherSuites-Cipher")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenFirewallVip6SslCipherSuitesPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "FirewallVip6-SslCipherSuites-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {
			v := flattenFirewallVip6SslCipherSuitesVersions(i["versions"], d, pre_append)
			tmp["versions"] = fortiAPISubPartPatch(v, "FirewallVip6-SslCipherSuites-Versions")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallVip6SslCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVip6SslClientFallback(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslClientRekeyCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslClientRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslClientSessionStateMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslClientSessionStateTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslClientSessionStateType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslDhBits(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslHpkp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslHpkpAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslHpkpBackup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVip6SslHpkpIncludeSubdomains(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslHpkpPrimary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVip6SslHpkpReportUri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslHsts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslHstsAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslHstsIncludeSubdomains(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslHttpLocationConversion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslHttpMatchHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslPfs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslSendEmptyFrags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslServerAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslServerCipherSuites(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallVip6SslServerCipherSuitesCipher(i["cipher"], d, pre_append)
			tmp["cipher"] = fortiAPISubPartPatch(v, "FirewallVip6-SslServerCipherSuites-Cipher")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenFirewallVip6SslServerCipherSuitesPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "FirewallVip6-SslServerCipherSuites-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {
			v := flattenFirewallVip6SslServerCipherSuitesVersions(i["versions"], d, pre_append)
			tmp["versions"] = fortiAPISubPartPatch(v, "FirewallVip6-SslServerCipherSuites-Versions")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallVip6SslServerCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslServerCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslServerCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVip6SslServerMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslServerMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslServerRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslServerSessionStateMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslServerSessionStateTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslServerSessionStateType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6Type(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6UserAgentDetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6Uuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6WeblogicServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6WebsphereServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallVip6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("add_nat64_route", flattenFirewallVip6AddNat64Route(o["add-nat64-route"], d, "add_nat64_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["add-nat64-route"], "FirewallVip6-AddNat64Route"); ok {
			if err = d.Set("add_nat64_route", vv); err != nil {
				return fmt.Errorf("Error reading add_nat64_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading add_nat64_route: %v", err)
		}
	}

	if err = d.Set("arp_reply", flattenFirewallVip6ArpReply(o["arp-reply"], d, "arp_reply")); err != nil {
		if vv, ok := fortiAPIPatch(o["arp-reply"], "FirewallVip6-ArpReply"); ok {
			if err = d.Set("arp_reply", vv); err != nil {
				return fmt.Errorf("Error reading arp_reply: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading arp_reply: %v", err)
		}
	}

	if err = d.Set("client_cert", flattenFirewallVip6ClientCert(o["client-cert"], d, "client_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-cert"], "FirewallVip6-ClientCert"); ok {
			if err = d.Set("client_cert", vv); err != nil {
				return fmt.Errorf("Error reading client_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_cert: %v", err)
		}
	}

	if err = d.Set("color", flattenFirewallVip6Color(o["color"], d, "color")); err != nil {
		if vv, ok := fortiAPIPatch(o["color"], "FirewallVip6-Color"); ok {
			if err = d.Set("color", vv); err != nil {
				return fmt.Errorf("Error reading color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("comment", flattenFirewallVip6Comment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "FirewallVip6-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenFirewallVip6DynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "FirewallVip6-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenFirewallVip6DynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "FirewallVip6-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("embedded_ipv4_address", flattenFirewallVip6EmbeddedIpv4Address(o["embedded-ipv4-address"], d, "embedded_ipv4_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["embedded-ipv4-address"], "FirewallVip6-EmbeddedIpv4Address"); ok {
			if err = d.Set("embedded_ipv4_address", vv); err != nil {
				return fmt.Errorf("Error reading embedded_ipv4_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading embedded_ipv4_address: %v", err)
		}
	}

	if err = d.Set("empty_cert_action", flattenFirewallVip6EmptyCertAction(o["empty-cert-action"], d, "empty_cert_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["empty-cert-action"], "FirewallVip6-EmptyCertAction"); ok {
			if err = d.Set("empty_cert_action", vv); err != nil {
				return fmt.Errorf("Error reading empty_cert_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading empty_cert_action: %v", err)
		}
	}

	if err = d.Set("extip", flattenFirewallVip6Extip(o["extip"], d, "extip")); err != nil {
		if vv, ok := fortiAPIPatch(o["extip"], "FirewallVip6-Extip"); ok {
			if err = d.Set("extip", vv); err != nil {
				return fmt.Errorf("Error reading extip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extip: %v", err)
		}
	}

	if err = d.Set("extport", flattenFirewallVip6Extport(o["extport"], d, "extport")); err != nil {
		if vv, ok := fortiAPIPatch(o["extport"], "FirewallVip6-Extport"); ok {
			if err = d.Set("extport", vv); err != nil {
				return fmt.Errorf("Error reading extport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extport: %v", err)
		}
	}

	if err = d.Set("h2_support", flattenFirewallVip6H2Support(o["h2-support"], d, "h2_support")); err != nil {
		if vv, ok := fortiAPIPatch(o["h2-support"], "FirewallVip6-H2Support"); ok {
			if err = d.Set("h2_support", vv); err != nil {
				return fmt.Errorf("Error reading h2_support: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading h2_support: %v", err)
		}
	}

	if err = d.Set("h3_support", flattenFirewallVip6H3Support(o["h3-support"], d, "h3_support")); err != nil {
		if vv, ok := fortiAPIPatch(o["h3-support"], "FirewallVip6-H3Support"); ok {
			if err = d.Set("h3_support", vv); err != nil {
				return fmt.Errorf("Error reading h3_support: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading h3_support: %v", err)
		}
	}

	if err = d.Set("http_cookie_age", flattenFirewallVip6HttpCookieAge(o["http-cookie-age"], d, "http_cookie_age")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-age"], "FirewallVip6-HttpCookieAge"); ok {
			if err = d.Set("http_cookie_age", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_age: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_age: %v", err)
		}
	}

	if err = d.Set("http_cookie_domain", flattenFirewallVip6HttpCookieDomain(o["http-cookie-domain"], d, "http_cookie_domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-domain"], "FirewallVip6-HttpCookieDomain"); ok {
			if err = d.Set("http_cookie_domain", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_domain: %v", err)
		}
	}

	if err = d.Set("http_cookie_domain_from_host", flattenFirewallVip6HttpCookieDomainFromHost(o["http-cookie-domain-from-host"], d, "http_cookie_domain_from_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-domain-from-host"], "FirewallVip6-HttpCookieDomainFromHost"); ok {
			if err = d.Set("http_cookie_domain_from_host", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_domain_from_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_domain_from_host: %v", err)
		}
	}

	if err = d.Set("http_cookie_generation", flattenFirewallVip6HttpCookieGeneration(o["http-cookie-generation"], d, "http_cookie_generation")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-generation"], "FirewallVip6-HttpCookieGeneration"); ok {
			if err = d.Set("http_cookie_generation", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_generation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_generation: %v", err)
		}
	}

	if err = d.Set("http_cookie_path", flattenFirewallVip6HttpCookiePath(o["http-cookie-path"], d, "http_cookie_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-path"], "FirewallVip6-HttpCookiePath"); ok {
			if err = d.Set("http_cookie_path", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_path: %v", err)
		}
	}

	if err = d.Set("http_cookie_share", flattenFirewallVip6HttpCookieShare(o["http-cookie-share"], d, "http_cookie_share")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-share"], "FirewallVip6-HttpCookieShare"); ok {
			if err = d.Set("http_cookie_share", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_share: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_share: %v", err)
		}
	}

	if err = d.Set("http_ip_header", flattenFirewallVip6HttpIpHeader(o["http-ip-header"], d, "http_ip_header")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-ip-header"], "FirewallVip6-HttpIpHeader"); ok {
			if err = d.Set("http_ip_header", vv); err != nil {
				return fmt.Errorf("Error reading http_ip_header: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_ip_header: %v", err)
		}
	}

	if err = d.Set("http_ip_header_name", flattenFirewallVip6HttpIpHeaderName(o["http-ip-header-name"], d, "http_ip_header_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-ip-header-name"], "FirewallVip6-HttpIpHeaderName"); ok {
			if err = d.Set("http_ip_header_name", vv); err != nil {
				return fmt.Errorf("Error reading http_ip_header_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_ip_header_name: %v", err)
		}
	}

	if err = d.Set("http_multiplex", flattenFirewallVip6HttpMultiplex(o["http-multiplex"], d, "http_multiplex")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-multiplex"], "FirewallVip6-HttpMultiplex"); ok {
			if err = d.Set("http_multiplex", vv); err != nil {
				return fmt.Errorf("Error reading http_multiplex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_multiplex: %v", err)
		}
	}

	if err = d.Set("http_redirect", flattenFirewallVip6HttpRedirect(o["http-redirect"], d, "http_redirect")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-redirect"], "FirewallVip6-HttpRedirect"); ok {
			if err = d.Set("http_redirect", vv); err != nil {
				return fmt.Errorf("Error reading http_redirect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_redirect: %v", err)
		}
	}

	if err = d.Set("https_cookie_secure", flattenFirewallVip6HttpsCookieSecure(o["https-cookie-secure"], d, "https_cookie_secure")); err != nil {
		if vv, ok := fortiAPIPatch(o["https-cookie-secure"], "FirewallVip6-HttpsCookieSecure"); ok {
			if err = d.Set("https_cookie_secure", vv); err != nil {
				return fmt.Errorf("Error reading https_cookie_secure: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading https_cookie_secure: %v", err)
		}
	}

	if err = d.Set("fosid", flattenFirewallVip6Id(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FirewallVip6-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ipv4_mappedip", flattenFirewallVip6Ipv4Mappedip(o["ipv4-mappedip"], d, "ipv4_mappedip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-mappedip"], "FirewallVip6-Ipv4Mappedip"); ok {
			if err = d.Set("ipv4_mappedip", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_mappedip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_mappedip: %v", err)
		}
	}

	if err = d.Set("ipv4_mappedport", flattenFirewallVip6Ipv4Mappedport(o["ipv4-mappedport"], d, "ipv4_mappedport")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-mappedport"], "FirewallVip6-Ipv4Mappedport"); ok {
			if err = d.Set("ipv4_mappedport", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_mappedport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_mappedport: %v", err)
		}
	}

	if err = d.Set("ldb_method", flattenFirewallVip6LdbMethod(o["ldb-method"], d, "ldb_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldb-method"], "FirewallVip6-LdbMethod"); ok {
			if err = d.Set("ldb_method", vv); err != nil {
				return fmt.Errorf("Error reading ldb_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldb_method: %v", err)
		}
	}

	if err = d.Set("mappedip", flattenFirewallVip6Mappedip(o["mappedip"], d, "mappedip")); err != nil {
		if vv, ok := fortiAPIPatch(o["mappedip"], "FirewallVip6-Mappedip"); ok {
			if err = d.Set("mappedip", vv); err != nil {
				return fmt.Errorf("Error reading mappedip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mappedip: %v", err)
		}
	}

	if err = d.Set("mappedport", flattenFirewallVip6Mappedport(o["mappedport"], d, "mappedport")); err != nil {
		if vv, ok := fortiAPIPatch(o["mappedport"], "FirewallVip6-Mappedport"); ok {
			if err = d.Set("mappedport", vv); err != nil {
				return fmt.Errorf("Error reading mappedport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mappedport: %v", err)
		}
	}

	if err = d.Set("max_embryonic_connections", flattenFirewallVip6MaxEmbryonicConnections(o["max-embryonic-connections"], d, "max_embryonic_connections")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-embryonic-connections"], "FirewallVip6-MaxEmbryonicConnections"); ok {
			if err = d.Set("max_embryonic_connections", vv); err != nil {
				return fmt.Errorf("Error reading max_embryonic_connections: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_embryonic_connections: %v", err)
		}
	}

	if err = d.Set("monitor", flattenFirewallVip6Monitor(o["monitor"], d, "monitor")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor"], "FirewallVip6-Monitor"); ok {
			if err = d.Set("monitor", vv); err != nil {
				return fmt.Errorf("Error reading monitor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallVip6Name(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallVip6-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("nat_source_vip", flattenFirewallVip6NatSourceVip(o["nat-source-vip"], d, "nat_source_vip")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat-source-vip"], "FirewallVip6-NatSourceVip"); ok {
			if err = d.Set("nat_source_vip", vv); err != nil {
				return fmt.Errorf("Error reading nat_source_vip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat_source_vip: %v", err)
		}
	}

	if err = d.Set("nat64", flattenFirewallVip6Nat64(o["nat64"], d, "nat64")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat64"], "FirewallVip6-Nat64"); ok {
			if err = d.Set("nat64", vv); err != nil {
				return fmt.Errorf("Error reading nat64: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat64: %v", err)
		}
	}

	if err = d.Set("nat66", flattenFirewallVip6Nat66(o["nat66"], d, "nat66")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat66"], "FirewallVip6-Nat66"); ok {
			if err = d.Set("nat66", vv); err != nil {
				return fmt.Errorf("Error reading nat66: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat66: %v", err)
		}
	}

	if err = d.Set("ndp_reply", flattenFirewallVip6NdpReply(o["ndp-reply"], d, "ndp_reply")); err != nil {
		if vv, ok := fortiAPIPatch(o["ndp-reply"], "FirewallVip6-NdpReply"); ok {
			if err = d.Set("ndp_reply", vv); err != nil {
				return fmt.Errorf("Error reading ndp_reply: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ndp_reply: %v", err)
		}
	}

	if err = d.Set("outlook_web_access", flattenFirewallVip6OutlookWebAccess(o["outlook-web-access"], d, "outlook_web_access")); err != nil {
		if vv, ok := fortiAPIPatch(o["outlook-web-access"], "FirewallVip6-OutlookWebAccess"); ok {
			if err = d.Set("outlook_web_access", vv); err != nil {
				return fmt.Errorf("Error reading outlook_web_access: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outlook_web_access: %v", err)
		}
	}

	if err = d.Set("persistence", flattenFirewallVip6Persistence(o["persistence"], d, "persistence")); err != nil {
		if vv, ok := fortiAPIPatch(o["persistence"], "FirewallVip6-Persistence"); ok {
			if err = d.Set("persistence", vv); err != nil {
				return fmt.Errorf("Error reading persistence: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading persistence: %v", err)
		}
	}

	if err = d.Set("portforward", flattenFirewallVip6Portforward(o["portforward"], d, "portforward")); err != nil {
		if vv, ok := fortiAPIPatch(o["portforward"], "FirewallVip6-Portforward"); ok {
			if err = d.Set("portforward", vv); err != nil {
				return fmt.Errorf("Error reading portforward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading portforward: %v", err)
		}
	}

	if err = d.Set("protocol", flattenFirewallVip6Protocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "FirewallVip6-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("quic", flattenFirewallVip6Quic(o["quic"], d, "quic")); err != nil {
			if vv, ok := fortiAPIPatch(o["quic"], "FirewallVip6-Quic"); ok {
				if err = d.Set("quic", vv); err != nil {
					return fmt.Errorf("Error reading quic: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading quic: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("quic"); ok {
			if err = d.Set("quic", flattenFirewallVip6Quic(o["quic"], d, "quic")); err != nil {
				if vv, ok := fortiAPIPatch(o["quic"], "FirewallVip6-Quic"); ok {
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
		if err = d.Set("realservers", flattenFirewallVip6Realservers(o["realservers"], d, "realservers")); err != nil {
			if vv, ok := fortiAPIPatch(o["realservers"], "FirewallVip6-Realservers"); ok {
				if err = d.Set("realservers", vv); err != nil {
					return fmt.Errorf("Error reading realservers: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading realservers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("realservers"); ok {
			if err = d.Set("realservers", flattenFirewallVip6Realservers(o["realservers"], d, "realservers")); err != nil {
				if vv, ok := fortiAPIPatch(o["realservers"], "FirewallVip6-Realservers"); ok {
					if err = d.Set("realservers", vv); err != nil {
						return fmt.Errorf("Error reading realservers: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading realservers: %v", err)
				}
			}
		}
	}

	if err = d.Set("server_type", flattenFirewallVip6ServerType(o["server-type"], d, "server_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-type"], "FirewallVip6-ServerType"); ok {
			if err = d.Set("server_type", vv); err != nil {
				return fmt.Errorf("Error reading server_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if err = d.Set("src_filter", flattenFirewallVip6SrcFilter(o["src-filter"], d, "src_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-filter"], "FirewallVip6-SrcFilter"); ok {
			if err = d.Set("src_filter", vv); err != nil {
				return fmt.Errorf("Error reading src_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_filter: %v", err)
		}
	}

	if err = d.Set("src_vip_filter", flattenFirewallVip6SrcVipFilter(o["src-vip-filter"], d, "src_vip_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-vip-filter"], "FirewallVip6-SrcVipFilter"); ok {
			if err = d.Set("src_vip_filter", vv); err != nil {
				return fmt.Errorf("Error reading src_vip_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_vip_filter: %v", err)
		}
	}

	if err = d.Set("ssl_accept_ffdhe_groups", flattenFirewallVip6SslAcceptFfdheGroups(o["ssl-accept-ffdhe-groups"], d, "ssl_accept_ffdhe_groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-accept-ffdhe-groups"], "FirewallVip6-SslAcceptFfdheGroups"); ok {
			if err = d.Set("ssl_accept_ffdhe_groups", vv); err != nil {
				return fmt.Errorf("Error reading ssl_accept_ffdhe_groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_accept_ffdhe_groups: %v", err)
		}
	}

	if err = d.Set("ssl_algorithm", flattenFirewallVip6SslAlgorithm(o["ssl-algorithm"], d, "ssl_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-algorithm"], "FirewallVip6-SslAlgorithm"); ok {
			if err = d.Set("ssl_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading ssl_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_algorithm: %v", err)
		}
	}

	if err = d.Set("ssl_certificate", flattenFirewallVip6SslCertificate(o["ssl-certificate"], d, "ssl_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-certificate"], "FirewallVip6-SslCertificate"); ok {
			if err = d.Set("ssl_certificate", vv); err != nil {
				return fmt.Errorf("Error reading ssl_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_certificate: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ssl_cipher_suites", flattenFirewallVip6SslCipherSuites(o["ssl-cipher-suites"], d, "ssl_cipher_suites")); err != nil {
			if vv, ok := fortiAPIPatch(o["ssl-cipher-suites"], "FirewallVip6-SslCipherSuites"); ok {
				if err = d.Set("ssl_cipher_suites", vv); err != nil {
					return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_cipher_suites"); ok {
			if err = d.Set("ssl_cipher_suites", flattenFirewallVip6SslCipherSuites(o["ssl-cipher-suites"], d, "ssl_cipher_suites")); err != nil {
				if vv, ok := fortiAPIPatch(o["ssl-cipher-suites"], "FirewallVip6-SslCipherSuites"); ok {
					if err = d.Set("ssl_cipher_suites", vv); err != nil {
						return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_client_fallback", flattenFirewallVip6SslClientFallback(o["ssl-client-fallback"], d, "ssl_client_fallback")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-fallback"], "FirewallVip6-SslClientFallback"); ok {
			if err = d.Set("ssl_client_fallback", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_fallback: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_fallback: %v", err)
		}
	}

	if err = d.Set("ssl_client_rekey_count", flattenFirewallVip6SslClientRekeyCount(o["ssl-client-rekey-count"], d, "ssl_client_rekey_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-rekey-count"], "FirewallVip6-SslClientRekeyCount"); ok {
			if err = d.Set("ssl_client_rekey_count", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_rekey_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_rekey_count: %v", err)
		}
	}

	if err = d.Set("ssl_client_renegotiation", flattenFirewallVip6SslClientRenegotiation(o["ssl-client-renegotiation"], d, "ssl_client_renegotiation")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-renegotiation"], "FirewallVip6-SslClientRenegotiation"); ok {
			if err = d.Set("ssl_client_renegotiation", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_renegotiation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_renegotiation: %v", err)
		}
	}

	if err = d.Set("ssl_client_session_state_max", flattenFirewallVip6SslClientSessionStateMax(o["ssl-client-session-state-max"], d, "ssl_client_session_state_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-session-state-max"], "FirewallVip6-SslClientSessionStateMax"); ok {
			if err = d.Set("ssl_client_session_state_max", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_session_state_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_session_state_max: %v", err)
		}
	}

	if err = d.Set("ssl_client_session_state_timeout", flattenFirewallVip6SslClientSessionStateTimeout(o["ssl-client-session-state-timeout"], d, "ssl_client_session_state_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-session-state-timeout"], "FirewallVip6-SslClientSessionStateTimeout"); ok {
			if err = d.Set("ssl_client_session_state_timeout", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_session_state_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_session_state_timeout: %v", err)
		}
	}

	if err = d.Set("ssl_client_session_state_type", flattenFirewallVip6SslClientSessionStateType(o["ssl-client-session-state-type"], d, "ssl_client_session_state_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-session-state-type"], "FirewallVip6-SslClientSessionStateType"); ok {
			if err = d.Set("ssl_client_session_state_type", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_session_state_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_session_state_type: %v", err)
		}
	}

	if err = d.Set("ssl_dh_bits", flattenFirewallVip6SslDhBits(o["ssl-dh-bits"], d, "ssl_dh_bits")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-dh-bits"], "FirewallVip6-SslDhBits"); ok {
			if err = d.Set("ssl_dh_bits", vv); err != nil {
				return fmt.Errorf("Error reading ssl_dh_bits: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_dh_bits: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp", flattenFirewallVip6SslHpkp(o["ssl-hpkp"], d, "ssl_hpkp")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hpkp"], "FirewallVip6-SslHpkp"); ok {
			if err = d.Set("ssl_hpkp", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hpkp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hpkp: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_age", flattenFirewallVip6SslHpkpAge(o["ssl-hpkp-age"], d, "ssl_hpkp_age")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hpkp-age"], "FirewallVip6-SslHpkpAge"); ok {
			if err = d.Set("ssl_hpkp_age", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hpkp_age: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hpkp_age: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_backup", flattenFirewallVip6SslHpkpBackup(o["ssl-hpkp-backup"], d, "ssl_hpkp_backup")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hpkp-backup"], "FirewallVip6-SslHpkpBackup"); ok {
			if err = d.Set("ssl_hpkp_backup", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hpkp_backup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hpkp_backup: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_include_subdomains", flattenFirewallVip6SslHpkpIncludeSubdomains(o["ssl-hpkp-include-subdomains"], d, "ssl_hpkp_include_subdomains")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hpkp-include-subdomains"], "FirewallVip6-SslHpkpIncludeSubdomains"); ok {
			if err = d.Set("ssl_hpkp_include_subdomains", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hpkp_include_subdomains: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hpkp_include_subdomains: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_primary", flattenFirewallVip6SslHpkpPrimary(o["ssl-hpkp-primary"], d, "ssl_hpkp_primary")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hpkp-primary"], "FirewallVip6-SslHpkpPrimary"); ok {
			if err = d.Set("ssl_hpkp_primary", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hpkp_primary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hpkp_primary: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_report_uri", flattenFirewallVip6SslHpkpReportUri(o["ssl-hpkp-report-uri"], d, "ssl_hpkp_report_uri")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hpkp-report-uri"], "FirewallVip6-SslHpkpReportUri"); ok {
			if err = d.Set("ssl_hpkp_report_uri", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hpkp_report_uri: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hpkp_report_uri: %v", err)
		}
	}

	if err = d.Set("ssl_hsts", flattenFirewallVip6SslHsts(o["ssl-hsts"], d, "ssl_hsts")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hsts"], "FirewallVip6-SslHsts"); ok {
			if err = d.Set("ssl_hsts", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hsts: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hsts: %v", err)
		}
	}

	if err = d.Set("ssl_hsts_age", flattenFirewallVip6SslHstsAge(o["ssl-hsts-age"], d, "ssl_hsts_age")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hsts-age"], "FirewallVip6-SslHstsAge"); ok {
			if err = d.Set("ssl_hsts_age", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hsts_age: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hsts_age: %v", err)
		}
	}

	if err = d.Set("ssl_hsts_include_subdomains", flattenFirewallVip6SslHstsIncludeSubdomains(o["ssl-hsts-include-subdomains"], d, "ssl_hsts_include_subdomains")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hsts-include-subdomains"], "FirewallVip6-SslHstsIncludeSubdomains"); ok {
			if err = d.Set("ssl_hsts_include_subdomains", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hsts_include_subdomains: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hsts_include_subdomains: %v", err)
		}
	}

	if err = d.Set("ssl_http_location_conversion", flattenFirewallVip6SslHttpLocationConversion(o["ssl-http-location-conversion"], d, "ssl_http_location_conversion")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-http-location-conversion"], "FirewallVip6-SslHttpLocationConversion"); ok {
			if err = d.Set("ssl_http_location_conversion", vv); err != nil {
				return fmt.Errorf("Error reading ssl_http_location_conversion: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_http_location_conversion: %v", err)
		}
	}

	if err = d.Set("ssl_http_match_host", flattenFirewallVip6SslHttpMatchHost(o["ssl-http-match-host"], d, "ssl_http_match_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-http-match-host"], "FirewallVip6-SslHttpMatchHost"); ok {
			if err = d.Set("ssl_http_match_host", vv); err != nil {
				return fmt.Errorf("Error reading ssl_http_match_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_http_match_host: %v", err)
		}
	}

	if err = d.Set("ssl_max_version", flattenFirewallVip6SslMaxVersion(o["ssl-max-version"], d, "ssl_max_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-max-version"], "FirewallVip6-SslMaxVersion"); ok {
			if err = d.Set("ssl_max_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_max_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_max_version: %v", err)
		}
	}

	if err = d.Set("ssl_min_version", flattenFirewallVip6SslMinVersion(o["ssl-min-version"], d, "ssl_min_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-min-version"], "FirewallVip6-SslMinVersion"); ok {
			if err = d.Set("ssl_min_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_min_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_min_version: %v", err)
		}
	}

	if err = d.Set("ssl_mode", flattenFirewallVip6SslMode(o["ssl-mode"], d, "ssl_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-mode"], "FirewallVip6-SslMode"); ok {
			if err = d.Set("ssl_mode", vv); err != nil {
				return fmt.Errorf("Error reading ssl_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_mode: %v", err)
		}
	}

	if err = d.Set("ssl_pfs", flattenFirewallVip6SslPfs(o["ssl-pfs"], d, "ssl_pfs")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-pfs"], "FirewallVip6-SslPfs"); ok {
			if err = d.Set("ssl_pfs", vv); err != nil {
				return fmt.Errorf("Error reading ssl_pfs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_pfs: %v", err)
		}
	}

	if err = d.Set("ssl_send_empty_frags", flattenFirewallVip6SslSendEmptyFrags(o["ssl-send-empty-frags"], d, "ssl_send_empty_frags")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-send-empty-frags"], "FirewallVip6-SslSendEmptyFrags"); ok {
			if err = d.Set("ssl_send_empty_frags", vv); err != nil {
				return fmt.Errorf("Error reading ssl_send_empty_frags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_send_empty_frags: %v", err)
		}
	}

	if err = d.Set("ssl_server_algorithm", flattenFirewallVip6SslServerAlgorithm(o["ssl-server-algorithm"], d, "ssl_server_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-algorithm"], "FirewallVip6-SslServerAlgorithm"); ok {
			if err = d.Set("ssl_server_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_algorithm: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ssl_server_cipher_suites", flattenFirewallVip6SslServerCipherSuites(o["ssl-server-cipher-suites"], d, "ssl_server_cipher_suites")); err != nil {
			if vv, ok := fortiAPIPatch(o["ssl-server-cipher-suites"], "FirewallVip6-SslServerCipherSuites"); ok {
				if err = d.Set("ssl_server_cipher_suites", vv); err != nil {
					return fmt.Errorf("Error reading ssl_server_cipher_suites: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ssl_server_cipher_suites: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_server_cipher_suites"); ok {
			if err = d.Set("ssl_server_cipher_suites", flattenFirewallVip6SslServerCipherSuites(o["ssl-server-cipher-suites"], d, "ssl_server_cipher_suites")); err != nil {
				if vv, ok := fortiAPIPatch(o["ssl-server-cipher-suites"], "FirewallVip6-SslServerCipherSuites"); ok {
					if err = d.Set("ssl_server_cipher_suites", vv); err != nil {
						return fmt.Errorf("Error reading ssl_server_cipher_suites: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ssl_server_cipher_suites: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_server_max_version", flattenFirewallVip6SslServerMaxVersion(o["ssl-server-max-version"], d, "ssl_server_max_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-max-version"], "FirewallVip6-SslServerMaxVersion"); ok {
			if err = d.Set("ssl_server_max_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_max_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_max_version: %v", err)
		}
	}

	if err = d.Set("ssl_server_min_version", flattenFirewallVip6SslServerMinVersion(o["ssl-server-min-version"], d, "ssl_server_min_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-min-version"], "FirewallVip6-SslServerMinVersion"); ok {
			if err = d.Set("ssl_server_min_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_min_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_min_version: %v", err)
		}
	}

	if err = d.Set("ssl_server_renegotiation", flattenFirewallVip6SslServerRenegotiation(o["ssl-server-renegotiation"], d, "ssl_server_renegotiation")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-renegotiation"], "FirewallVip6-SslServerRenegotiation"); ok {
			if err = d.Set("ssl_server_renegotiation", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_renegotiation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_renegotiation: %v", err)
		}
	}

	if err = d.Set("ssl_server_session_state_max", flattenFirewallVip6SslServerSessionStateMax(o["ssl-server-session-state-max"], d, "ssl_server_session_state_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-session-state-max"], "FirewallVip6-SslServerSessionStateMax"); ok {
			if err = d.Set("ssl_server_session_state_max", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_session_state_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_session_state_max: %v", err)
		}
	}

	if err = d.Set("ssl_server_session_state_timeout", flattenFirewallVip6SslServerSessionStateTimeout(o["ssl-server-session-state-timeout"], d, "ssl_server_session_state_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-session-state-timeout"], "FirewallVip6-SslServerSessionStateTimeout"); ok {
			if err = d.Set("ssl_server_session_state_timeout", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_session_state_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_session_state_timeout: %v", err)
		}
	}

	if err = d.Set("ssl_server_session_state_type", flattenFirewallVip6SslServerSessionStateType(o["ssl-server-session-state-type"], d, "ssl_server_session_state_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-session-state-type"], "FirewallVip6-SslServerSessionStateType"); ok {
			if err = d.Set("ssl_server_session_state_type", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_session_state_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_session_state_type: %v", err)
		}
	}

	if err = d.Set("type", flattenFirewallVip6Type(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "FirewallVip6-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("user_agent_detect", flattenFirewallVip6UserAgentDetect(o["user-agent-detect"], d, "user_agent_detect")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-agent-detect"], "FirewallVip6-UserAgentDetect"); ok {
			if err = d.Set("user_agent_detect", vv); err != nil {
				return fmt.Errorf("Error reading user_agent_detect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_agent_detect: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallVip6Uuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "FirewallVip6-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("weblogic_server", flattenFirewallVip6WeblogicServer(o["weblogic-server"], d, "weblogic_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["weblogic-server"], "FirewallVip6-WeblogicServer"); ok {
			if err = d.Set("weblogic_server", vv); err != nil {
				return fmt.Errorf("Error reading weblogic_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weblogic_server: %v", err)
		}
	}

	if err = d.Set("websphere_server", flattenFirewallVip6WebsphereServer(o["websphere-server"], d, "websphere_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["websphere-server"], "FirewallVip6-WebsphereServer"); ok {
			if err = d.Set("websphere_server", vv); err != nil {
				return fmt.Errorf("Error reading websphere_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading websphere_server: %v", err)
		}
	}

	return nil
}

func flattenFirewallVip6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallVip6AddNat64Route(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6ArpReply(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6ClientCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Color(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Comment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandFirewallVip6DynamicMappingScope(d, i["_scope"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["_scope"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_nat64_route"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["add-nat64-route"], _ = expandFirewallVip6DynamicMappingAddNat64Route(d, i["add_nat64_route"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "arp_reply"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["arp-reply"], _ = expandFirewallVip6DynamicMappingArpReply(d, i["arp_reply"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["client-cert"], _ = expandFirewallVip6DynamicMappingClientCert(d, i["client_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["color"], _ = expandFirewallVip6DynamicMappingColor(d, i["color"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandFirewallVip6DynamicMappingComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "embedded_ipv4_address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["embedded-ipv4-address"], _ = expandFirewallVip6DynamicMappingEmbeddedIpv4Address(d, i["embedded_ipv4_address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "empty_cert_action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["empty-cert-action"], _ = expandFirewallVip6DynamicMappingEmptyCertAction(d, i["empty_cert_action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["extip"], _ = expandFirewallVip6DynamicMappingExtip(d, i["extip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["extport"], _ = expandFirewallVip6DynamicMappingExtport(d, i["extport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h2_support"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["h2-support"], _ = expandFirewallVip6DynamicMappingH2Support(d, i["h2_support"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h3_support"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["h3-support"], _ = expandFirewallVip6DynamicMappingH3Support(d, i["h3_support"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-age"], _ = expandFirewallVip6DynamicMappingHttpCookieAge(d, i["http_cookie_age"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-domain"], _ = expandFirewallVip6DynamicMappingHttpCookieDomain(d, i["http_cookie_domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-domain-from-host"], _ = expandFirewallVip6DynamicMappingHttpCookieDomainFromHost(d, i["http_cookie_domain_from_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-generation"], _ = expandFirewallVip6DynamicMappingHttpCookieGeneration(d, i["http_cookie_generation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-path"], _ = expandFirewallVip6DynamicMappingHttpCookiePath(d, i["http_cookie_path"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-share"], _ = expandFirewallVip6DynamicMappingHttpCookieShare(d, i["http_cookie_share"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_ip_header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-ip-header"], _ = expandFirewallVip6DynamicMappingHttpIpHeader(d, i["http_ip_header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_ip_header_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-ip-header-name"], _ = expandFirewallVip6DynamicMappingHttpIpHeaderName(d, i["http_ip_header_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_multiplex"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-multiplex"], _ = expandFirewallVip6DynamicMappingHttpMultiplex(d, i["http_multiplex"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_redirect"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-redirect"], _ = expandFirewallVip6DynamicMappingHttpRedirect(d, i["http_redirect"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["https-cookie-secure"], _ = expandFirewallVip6DynamicMappingHttpsCookieSecure(d, i["https_cookie_secure"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFirewallVip6DynamicMappingId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv4_mappedip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipv4-mappedip"], _ = expandFirewallVip6DynamicMappingIpv4Mappedip(d, i["ipv4_mappedip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv4_mappedport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipv4-mappedport"], _ = expandFirewallVip6DynamicMappingIpv4Mappedport(d, i["ipv4_mappedport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ldb-method"], _ = expandFirewallVip6DynamicMappingLdbMethod(d, i["ldb_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mappedip"], _ = expandFirewallVip6DynamicMappingMappedip(d, i["mappedip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mappedport"], _ = expandFirewallVip6DynamicMappingMappedport(d, i["mappedport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_embryonic_connections"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-embryonic-connections"], _ = expandFirewallVip6DynamicMappingMaxEmbryonicConnections(d, i["max_embryonic_connections"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["monitor"], _ = expandFirewallVip6DynamicMappingMonitor(d, i["monitor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat_source_vip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nat-source-vip"], _ = expandFirewallVip6DynamicMappingNatSourceVip(d, i["nat_source_vip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat64"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nat64"], _ = expandFirewallVip6DynamicMappingNat64(d, i["nat64"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat66"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nat66"], _ = expandFirewallVip6DynamicMappingNat66(d, i["nat66"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ndp_reply"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ndp-reply"], _ = expandFirewallVip6DynamicMappingNdpReply(d, i["ndp_reply"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "outlook_web_access"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["outlook-web-access"], _ = expandFirewallVip6DynamicMappingOutlookWebAccess(d, i["outlook_web_access"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["persistence"], _ = expandFirewallVip6DynamicMappingPersistence(d, i["persistence"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "portforward"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["portforward"], _ = expandFirewallVip6DynamicMappingPortforward(d, i["portforward"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandFirewallVip6DynamicMappingProtocol(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandFirewallVip6DynamicMappingRealservers(d, i["realservers"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["realservers"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server-type"], _ = expandFirewallVip6DynamicMappingServerType(d, i["server_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-filter"], _ = expandFirewallVip6DynamicMappingSrcFilter(d, i["src_filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_vip_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-vip-filter"], _ = expandFirewallVip6DynamicMappingSrcVipFilter(d, i["src_vip_filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_accept_ffdhe_groups"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-accept-ffdhe-groups"], _ = expandFirewallVip6DynamicMappingSslAcceptFfdheGroups(d, i["ssl_accept_ffdhe_groups"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-algorithm"], _ = expandFirewallVip6DynamicMappingSslAlgorithm(d, i["ssl_algorithm"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_certificate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-certificate"], _ = expandFirewallVip6DynamicMappingSslCertificate(d, i["ssl_certificate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandFirewallVip6DynamicMappingSslCipherSuites(d, i["ssl_cipher_suites"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["ssl-cipher-suites"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_fallback"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-client-fallback"], _ = expandFirewallVip6DynamicMappingSslClientFallback(d, i["ssl_client_fallback"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_rekey_count"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-client-rekey-count"], _ = expandFirewallVip6DynamicMappingSslClientRekeyCount(d, i["ssl_client_rekey_count"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_renegotiation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-client-renegotiation"], _ = expandFirewallVip6DynamicMappingSslClientRenegotiation(d, i["ssl_client_renegotiation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_session_state_max"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-client-session-state-max"], _ = expandFirewallVip6DynamicMappingSslClientSessionStateMax(d, i["ssl_client_session_state_max"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_session_state_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-client-session-state-timeout"], _ = expandFirewallVip6DynamicMappingSslClientSessionStateTimeout(d, i["ssl_client_session_state_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_session_state_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-client-session-state-type"], _ = expandFirewallVip6DynamicMappingSslClientSessionStateType(d, i["ssl_client_session_state_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-dh-bits"], _ = expandFirewallVip6DynamicMappingSslDhBits(d, i["ssl_dh_bits"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hpkp"], _ = expandFirewallVip6DynamicMappingSslHpkp(d, i["ssl_hpkp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_age"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hpkp-age"], _ = expandFirewallVip6DynamicMappingSslHpkpAge(d, i["ssl_hpkp_age"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_backup"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hpkp-backup"], _ = expandFirewallVip6DynamicMappingSslHpkpBackup(d, i["ssl_hpkp_backup"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_include_subdomains"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hpkp-include-subdomains"], _ = expandFirewallVip6DynamicMappingSslHpkpIncludeSubdomains(d, i["ssl_hpkp_include_subdomains"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_primary"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hpkp-primary"], _ = expandFirewallVip6DynamicMappingSslHpkpPrimary(d, i["ssl_hpkp_primary"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_report_uri"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hpkp-report-uri"], _ = expandFirewallVip6DynamicMappingSslHpkpReportUri(d, i["ssl_hpkp_report_uri"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hsts"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hsts"], _ = expandFirewallVip6DynamicMappingSslHsts(d, i["ssl_hsts"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hsts_age"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hsts-age"], _ = expandFirewallVip6DynamicMappingSslHstsAge(d, i["ssl_hsts_age"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hsts_include_subdomains"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hsts-include-subdomains"], _ = expandFirewallVip6DynamicMappingSslHstsIncludeSubdomains(d, i["ssl_hsts_include_subdomains"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_http_location_conversion"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-http-location-conversion"], _ = expandFirewallVip6DynamicMappingSslHttpLocationConversion(d, i["ssl_http_location_conversion"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_http_match_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-http-match-host"], _ = expandFirewallVip6DynamicMappingSslHttpMatchHost(d, i["ssl_http_match_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-max-version"], _ = expandFirewallVip6DynamicMappingSslMaxVersion(d, i["ssl_max_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-min-version"], _ = expandFirewallVip6DynamicMappingSslMinVersion(d, i["ssl_min_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-mode"], _ = expandFirewallVip6DynamicMappingSslMode(d, i["ssl_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_pfs"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-pfs"], _ = expandFirewallVip6DynamicMappingSslPfs(d, i["ssl_pfs"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_send_empty_frags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-send-empty-frags"], _ = expandFirewallVip6DynamicMappingSslSendEmptyFrags(d, i["ssl_send_empty_frags"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_algorithm"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-server-algorithm"], _ = expandFirewallVip6DynamicMappingSslServerAlgorithm(d, i["ssl_server_algorithm"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_max_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-server-max-version"], _ = expandFirewallVip6DynamicMappingSslServerMaxVersion(d, i["ssl_server_max_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_min_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-server-min-version"], _ = expandFirewallVip6DynamicMappingSslServerMinVersion(d, i["ssl_server_min_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_renegotiation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-server-renegotiation"], _ = expandFirewallVip6DynamicMappingSslServerRenegotiation(d, i["ssl_server_renegotiation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_session_state_max"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-server-session-state-max"], _ = expandFirewallVip6DynamicMappingSslServerSessionStateMax(d, i["ssl_server_session_state_max"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_session_state_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-server-session-state-timeout"], _ = expandFirewallVip6DynamicMappingSslServerSessionStateTimeout(d, i["ssl_server_session_state_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_session_state_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-server-session-state-type"], _ = expandFirewallVip6DynamicMappingSslServerSessionStateType(d, i["ssl_server_session_state_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandFirewallVip6DynamicMappingType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_agent_detect"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-agent-detect"], _ = expandFirewallVip6DynamicMappingUserAgentDetect(d, i["user_agent_detect"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uuid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uuid"], _ = expandFirewallVip6DynamicMappingUuid(d, i["uuid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weblogic_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weblogic-server"], _ = expandFirewallVip6DynamicMappingWeblogicServer(d, i["weblogic_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "websphere_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["websphere-server"], _ = expandFirewallVip6DynamicMappingWebsphereServer(d, i["websphere_server"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallVip6DynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandFirewallVip6DynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandFirewallVip6DynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallVip6DynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingAddNat64Route(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingArpReply(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingClientCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingEmbeddedIpv4Address(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingEmptyCertAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingExtip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVip6DynamicMappingExtport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingH2Support(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingH3Support(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingHttpCookieAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingHttpCookieDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingHttpCookieDomainFromHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingHttpCookieGeneration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingHttpCookiePath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingHttpCookieShare(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingHttpIpHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingHttpIpHeaderName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingHttpMultiplex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingHttpRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingHttpsCookieSecure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingIpv4Mappedip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingIpv4Mappedport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingLdbMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingMappedip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVip6DynamicMappingMappedport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingMaxEmbryonicConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVip6DynamicMappingNatSourceVip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingNat64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingNat66(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingNdpReply(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingOutlookWebAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingPersistence(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingPortforward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingRealservers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["client-ip"], _ = expandFirewallVip6DynamicMappingRealserversClientIp(d, i["client_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "healthcheck"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["healthcheck"], _ = expandFirewallVip6DynamicMappingRealserversHealthcheck(d, i["healthcheck"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["holddown-interval"], _ = expandFirewallVip6DynamicMappingRealserversHolddownInterval(d, i["holddown_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-host"], _ = expandFirewallVip6DynamicMappingRealserversHttpHost(d, i["http_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFirewallVip6DynamicMappingRealserversId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandFirewallVip6DynamicMappingRealserversIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_connections"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-connections"], _ = expandFirewallVip6DynamicMappingRealserversMaxConnections(d, i["max_connections"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["monitor"], _ = expandFirewallVip6DynamicMappingRealserversMonitor(d, i["monitor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandFirewallVip6DynamicMappingRealserversPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandFirewallVip6DynamicMappingRealserversStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["translate-host"], _ = expandFirewallVip6DynamicMappingRealserversTranslateHost(d, i["translate_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "verify_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["verify-cert"], _ = expandFirewallVip6DynamicMappingRealserversVerifyCert(d, i["verify_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandFirewallVip6DynamicMappingRealserversWeight(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallVip6DynamicMappingRealserversClientIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingRealserversHealthcheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingRealserversHolddownInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingRealserversHttpHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingRealserversId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingRealserversIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingRealserversMaxConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingRealserversMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVip6DynamicMappingRealserversPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingRealserversStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingRealserversTranslateHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingRealserversVerifyCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingRealserversWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSrcFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVip6DynamicMappingSrcVipFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslAcceptFfdheGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVip6DynamicMappingSslCipherSuites(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["cipher"], _ = expandFirewallVip6DynamicMappingSslCipherSuitesCipher(d, i["cipher"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandFirewallVip6DynamicMappingSslCipherSuitesPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["versions"], _ = expandFirewallVip6DynamicMappingSslCipherSuitesVersions(d, i["versions"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallVip6DynamicMappingSslCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVip6DynamicMappingSslClientFallback(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslClientRekeyCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslClientRenegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslClientSessionStateMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslClientSessionStateTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslClientSessionStateType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslDhBits(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslHpkp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslHpkpAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslHpkpBackup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVip6DynamicMappingSslHpkpIncludeSubdomains(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslHpkpPrimary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVip6DynamicMappingSslHpkpReportUri(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslHsts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslHstsAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslHstsIncludeSubdomains(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslHttpLocationConversion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslHttpMatchHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslPfs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslSendEmptyFrags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslServerAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslServerMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslServerMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslServerRenegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslServerSessionStateMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslServerSessionStateTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingSslServerSessionStateType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingUserAgentDetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingWeblogicServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6DynamicMappingWebsphereServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6EmbeddedIpv4Address(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6EmptyCertAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Extip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVip6Extport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6H2Support(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6H3Support(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6HttpCookieAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6HttpCookieDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6HttpCookieDomainFromHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6HttpCookieGeneration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6HttpCookiePath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6HttpCookieShare(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6HttpIpHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6HttpIpHeaderName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6HttpMultiplex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6HttpRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6HttpsCookieSecure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Ipv4Mappedip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Ipv4Mappedport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6LdbMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Mappedip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVip6Mappedport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6MaxEmbryonicConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Monitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVip6Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6NatSourceVip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Nat64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Nat66(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6NdpReply(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6OutlookWebAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Persistence(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Portforward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Protocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Quic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ack-delay-exponent"], _ = expandFirewallVip6QuicAckDelayExponent(d, i["ack_delay_exponent"], pre_append)
	}
	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["active-connection-id-limit"], _ = expandFirewallVip6QuicActiveConnectionIdLimit(d, i["active_connection_id_limit"], pre_append)
	}
	pre_append = pre + ".0." + "active_migration"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["active-migration"], _ = expandFirewallVip6QuicActiveMigration(d, i["active_migration"], pre_append)
	}
	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["grease-quic-bit"], _ = expandFirewallVip6QuicGreaseQuicBit(d, i["grease_quic_bit"], pre_append)
	}
	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-ack-delay"], _ = expandFirewallVip6QuicMaxAckDelay(d, i["max_ack_delay"], pre_append)
	}
	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-datagram-frame-size"], _ = expandFirewallVip6QuicMaxDatagramFrameSize(d, i["max_datagram_frame_size"], pre_append)
	}
	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-idle-timeout"], _ = expandFirewallVip6QuicMaxIdleTimeout(d, i["max_idle_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-udp-payload-size"], _ = expandFirewallVip6QuicMaxUdpPayloadSize(d, i["max_udp_payload_size"], pre_append)
	}

	return result, nil
}

func expandFirewallVip6QuicAckDelayExponent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6QuicActiveConnectionIdLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6QuicActiveMigration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6QuicGreaseQuicBit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6QuicMaxAckDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6QuicMaxDatagramFrameSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6QuicMaxIdleTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6QuicMaxUdpPayloadSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Realservers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["client-ip"], _ = expandFirewallVip6RealserversClientIp(d, i["client_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "healthcheck"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["healthcheck"], _ = expandFirewallVip6RealserversHealthcheck(d, i["healthcheck"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["holddown-interval"], _ = expandFirewallVip6RealserversHolddownInterval(d, i["holddown_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-host"], _ = expandFirewallVip6RealserversHttpHost(d, i["http_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFirewallVip6RealserversId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandFirewallVip6RealserversIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_connections"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-connections"], _ = expandFirewallVip6RealserversMaxConnections(d, i["max_connections"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["monitor"], _ = expandFirewallVip6RealserversMonitor(d, i["monitor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandFirewallVip6RealserversPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandFirewallVip6RealserversStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["translate-host"], _ = expandFirewallVip6RealserversTranslateHost(d, i["translate_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "verify_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["verify-cert"], _ = expandFirewallVip6RealserversVerifyCert(d, i["verify_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandFirewallVip6RealserversWeight(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallVip6RealserversClientIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversHealthcheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversHolddownInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversHttpHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversMaxConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVip6RealserversPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversTranslateHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversVerifyCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6ServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SrcFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVip6SrcVipFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslAcceptFfdheGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVip6SslCipherSuites(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["cipher"], _ = expandFirewallVip6SslCipherSuitesCipher(d, i["cipher"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandFirewallVip6SslCipherSuitesPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["versions"], _ = expandFirewallVip6SslCipherSuitesVersions(d, i["versions"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallVip6SslCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVip6SslClientFallback(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslClientRekeyCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslClientRenegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslClientSessionStateMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslClientSessionStateTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslClientSessionStateType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslDhBits(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslHpkp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslHpkpAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslHpkpBackup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVip6SslHpkpIncludeSubdomains(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslHpkpPrimary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVip6SslHpkpReportUri(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslHsts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslHstsAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslHstsIncludeSubdomains(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslHttpLocationConversion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslHttpMatchHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslPfs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslSendEmptyFrags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslServerAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslServerCipherSuites(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["cipher"], _ = expandFirewallVip6SslServerCipherSuitesCipher(d, i["cipher"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandFirewallVip6SslServerCipherSuitesPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["versions"], _ = expandFirewallVip6SslServerCipherSuitesVersions(d, i["versions"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallVip6SslServerCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslServerCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslServerCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVip6SslServerMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslServerMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslServerRenegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslServerSessionStateMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslServerSessionStateTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslServerSessionStateType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Type(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6UserAgentDetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Uuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6WeblogicServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6WebsphereServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallVip6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("add_nat64_route"); ok || d.HasChange("add_nat64_route") {
		t, err := expandFirewallVip6AddNat64Route(d, v, "add_nat64_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-nat64-route"] = t
		}
	}

	if v, ok := d.GetOk("arp_reply"); ok || d.HasChange("arp_reply") {
		t, err := expandFirewallVip6ArpReply(d, v, "arp_reply")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arp-reply"] = t
		}
	}

	if v, ok := d.GetOk("client_cert"); ok || d.HasChange("client_cert") {
		t, err := expandFirewallVip6ClientCert(d, v, "client_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert"] = t
		}
	}

	if v, ok := d.GetOk("color"); ok || d.HasChange("color") {
		t, err := expandFirewallVip6Color(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandFirewallVip6Comment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandFirewallVip6DynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("embedded_ipv4_address"); ok || d.HasChange("embedded_ipv4_address") {
		t, err := expandFirewallVip6EmbeddedIpv4Address(d, v, "embedded_ipv4_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["embedded-ipv4-address"] = t
		}
	}

	if v, ok := d.GetOk("empty_cert_action"); ok || d.HasChange("empty_cert_action") {
		t, err := expandFirewallVip6EmptyCertAction(d, v, "empty_cert_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["empty-cert-action"] = t
		}
	}

	if v, ok := d.GetOk("extip"); ok || d.HasChange("extip") {
		t, err := expandFirewallVip6Extip(d, v, "extip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extip"] = t
		}
	}

	if v, ok := d.GetOk("extport"); ok || d.HasChange("extport") {
		t, err := expandFirewallVip6Extport(d, v, "extport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extport"] = t
		}
	}

	if v, ok := d.GetOk("h2_support"); ok || d.HasChange("h2_support") {
		t, err := expandFirewallVip6H2Support(d, v, "h2_support")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h2-support"] = t
		}
	}

	if v, ok := d.GetOk("h3_support"); ok || d.HasChange("h3_support") {
		t, err := expandFirewallVip6H3Support(d, v, "h3_support")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-support"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_age"); ok || d.HasChange("http_cookie_age") {
		t, err := expandFirewallVip6HttpCookieAge(d, v, "http_cookie_age")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-age"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_domain"); ok || d.HasChange("http_cookie_domain") {
		t, err := expandFirewallVip6HttpCookieDomain(d, v, "http_cookie_domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-domain"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_domain_from_host"); ok || d.HasChange("http_cookie_domain_from_host") {
		t, err := expandFirewallVip6HttpCookieDomainFromHost(d, v, "http_cookie_domain_from_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-domain-from-host"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_generation"); ok || d.HasChange("http_cookie_generation") {
		t, err := expandFirewallVip6HttpCookieGeneration(d, v, "http_cookie_generation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-generation"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_path"); ok || d.HasChange("http_cookie_path") {
		t, err := expandFirewallVip6HttpCookiePath(d, v, "http_cookie_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-path"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_share"); ok || d.HasChange("http_cookie_share") {
		t, err := expandFirewallVip6HttpCookieShare(d, v, "http_cookie_share")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-share"] = t
		}
	}

	if v, ok := d.GetOk("http_ip_header"); ok || d.HasChange("http_ip_header") {
		t, err := expandFirewallVip6HttpIpHeader(d, v, "http_ip_header")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-ip-header"] = t
		}
	}

	if v, ok := d.GetOk("http_ip_header_name"); ok || d.HasChange("http_ip_header_name") {
		t, err := expandFirewallVip6HttpIpHeaderName(d, v, "http_ip_header_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-ip-header-name"] = t
		}
	}

	if v, ok := d.GetOk("http_multiplex"); ok || d.HasChange("http_multiplex") {
		t, err := expandFirewallVip6HttpMultiplex(d, v, "http_multiplex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-multiplex"] = t
		}
	}

	if v, ok := d.GetOk("http_redirect"); ok || d.HasChange("http_redirect") {
		t, err := expandFirewallVip6HttpRedirect(d, v, "http_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-redirect"] = t
		}
	}

	if v, ok := d.GetOk("https_cookie_secure"); ok || d.HasChange("https_cookie_secure") {
		t, err := expandFirewallVip6HttpsCookieSecure(d, v, "https_cookie_secure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-cookie-secure"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandFirewallVip6Id(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_mappedip"); ok || d.HasChange("ipv4_mappedip") {
		t, err := expandFirewallVip6Ipv4Mappedip(d, v, "ipv4_mappedip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-mappedip"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_mappedport"); ok || d.HasChange("ipv4_mappedport") {
		t, err := expandFirewallVip6Ipv4Mappedport(d, v, "ipv4_mappedport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-mappedport"] = t
		}
	}

	if v, ok := d.GetOk("ldb_method"); ok || d.HasChange("ldb_method") {
		t, err := expandFirewallVip6LdbMethod(d, v, "ldb_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldb-method"] = t
		}
	}

	if v, ok := d.GetOk("mappedip"); ok || d.HasChange("mappedip") {
		t, err := expandFirewallVip6Mappedip(d, v, "mappedip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mappedip"] = t
		}
	}

	if v, ok := d.GetOk("mappedport"); ok || d.HasChange("mappedport") {
		t, err := expandFirewallVip6Mappedport(d, v, "mappedport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mappedport"] = t
		}
	}

	if v, ok := d.GetOk("max_embryonic_connections"); ok || d.HasChange("max_embryonic_connections") {
		t, err := expandFirewallVip6MaxEmbryonicConnections(d, v, "max_embryonic_connections")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-embryonic-connections"] = t
		}
	}

	if v, ok := d.GetOk("monitor"); ok || d.HasChange("monitor") {
		t, err := expandFirewallVip6Monitor(d, v, "monitor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallVip6Name(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nat_source_vip"); ok || d.HasChange("nat_source_vip") {
		t, err := expandFirewallVip6NatSourceVip(d, v, "nat_source_vip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat-source-vip"] = t
		}
	}

	if v, ok := d.GetOk("nat64"); ok || d.HasChange("nat64") {
		t, err := expandFirewallVip6Nat64(d, v, "nat64")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat64"] = t
		}
	}

	if v, ok := d.GetOk("nat66"); ok || d.HasChange("nat66") {
		t, err := expandFirewallVip6Nat66(d, v, "nat66")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat66"] = t
		}
	}

	if v, ok := d.GetOk("ndp_reply"); ok || d.HasChange("ndp_reply") {
		t, err := expandFirewallVip6NdpReply(d, v, "ndp_reply")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ndp-reply"] = t
		}
	}

	if v, ok := d.GetOk("outlook_web_access"); ok || d.HasChange("outlook_web_access") {
		t, err := expandFirewallVip6OutlookWebAccess(d, v, "outlook_web_access")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outlook-web-access"] = t
		}
	}

	if v, ok := d.GetOk("persistence"); ok || d.HasChange("persistence") {
		t, err := expandFirewallVip6Persistence(d, v, "persistence")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["persistence"] = t
		}
	}

	if v, ok := d.GetOk("portforward"); ok || d.HasChange("portforward") {
		t, err := expandFirewallVip6Portforward(d, v, "portforward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portforward"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandFirewallVip6Protocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("quic"); ok || d.HasChange("quic") {
		t, err := expandFirewallVip6Quic(d, v, "quic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quic"] = t
		}
	}

	if v, ok := d.GetOk("realservers"); ok || d.HasChange("realservers") {
		t, err := expandFirewallVip6Realservers(d, v, "realservers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["realservers"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok || d.HasChange("server_type") {
		t, err := expandFirewallVip6ServerType(d, v, "server_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}

	if v, ok := d.GetOk("src_filter"); ok || d.HasChange("src_filter") {
		t, err := expandFirewallVip6SrcFilter(d, v, "src_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-filter"] = t
		}
	}

	if v, ok := d.GetOk("src_vip_filter"); ok || d.HasChange("src_vip_filter") {
		t, err := expandFirewallVip6SrcVipFilter(d, v, "src_vip_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-vip-filter"] = t
		}
	}

	if v, ok := d.GetOk("ssl_accept_ffdhe_groups"); ok || d.HasChange("ssl_accept_ffdhe_groups") {
		t, err := expandFirewallVip6SslAcceptFfdheGroups(d, v, "ssl_accept_ffdhe_groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-accept-ffdhe-groups"] = t
		}
	}

	if v, ok := d.GetOk("ssl_algorithm"); ok || d.HasChange("ssl_algorithm") {
		t, err := expandFirewallVip6SslAlgorithm(d, v, "ssl_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_certificate"); ok || d.HasChange("ssl_certificate") {
		t, err := expandFirewallVip6SslCertificate(d, v, "ssl_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-certificate"] = t
		}
	}

	if v, ok := d.GetOk("ssl_cipher_suites"); ok || d.HasChange("ssl_cipher_suites") {
		t, err := expandFirewallVip6SslCipherSuites(d, v, "ssl_cipher_suites")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-cipher-suites"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_fallback"); ok || d.HasChange("ssl_client_fallback") {
		t, err := expandFirewallVip6SslClientFallback(d, v, "ssl_client_fallback")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-fallback"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_rekey_count"); ok || d.HasChange("ssl_client_rekey_count") {
		t, err := expandFirewallVip6SslClientRekeyCount(d, v, "ssl_client_rekey_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-rekey-count"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_renegotiation"); ok || d.HasChange("ssl_client_renegotiation") {
		t, err := expandFirewallVip6SslClientRenegotiation(d, v, "ssl_client_renegotiation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-renegotiation"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_session_state_max"); ok || d.HasChange("ssl_client_session_state_max") {
		t, err := expandFirewallVip6SslClientSessionStateMax(d, v, "ssl_client_session_state_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-session-state-max"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_session_state_timeout"); ok || d.HasChange("ssl_client_session_state_timeout") {
		t, err := expandFirewallVip6SslClientSessionStateTimeout(d, v, "ssl_client_session_state_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-session-state-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_session_state_type"); ok || d.HasChange("ssl_client_session_state_type") {
		t, err := expandFirewallVip6SslClientSessionStateType(d, v, "ssl_client_session_state_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-session-state-type"] = t
		}
	}

	if v, ok := d.GetOk("ssl_dh_bits"); ok || d.HasChange("ssl_dh_bits") {
		t, err := expandFirewallVip6SslDhBits(d, v, "ssl_dh_bits")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-dh-bits"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp"); ok || d.HasChange("ssl_hpkp") {
		t, err := expandFirewallVip6SslHpkp(d, v, "ssl_hpkp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_age"); ok || d.HasChange("ssl_hpkp_age") {
		t, err := expandFirewallVip6SslHpkpAge(d, v, "ssl_hpkp_age")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-age"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_backup"); ok || d.HasChange("ssl_hpkp_backup") {
		t, err := expandFirewallVip6SslHpkpBackup(d, v, "ssl_hpkp_backup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-backup"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_include_subdomains"); ok || d.HasChange("ssl_hpkp_include_subdomains") {
		t, err := expandFirewallVip6SslHpkpIncludeSubdomains(d, v, "ssl_hpkp_include_subdomains")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-include-subdomains"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_primary"); ok || d.HasChange("ssl_hpkp_primary") {
		t, err := expandFirewallVip6SslHpkpPrimary(d, v, "ssl_hpkp_primary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-primary"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_report_uri"); ok || d.HasChange("ssl_hpkp_report_uri") {
		t, err := expandFirewallVip6SslHpkpReportUri(d, v, "ssl_hpkp_report_uri")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-report-uri"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hsts"); ok || d.HasChange("ssl_hsts") {
		t, err := expandFirewallVip6SslHsts(d, v, "ssl_hsts")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hsts"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hsts_age"); ok || d.HasChange("ssl_hsts_age") {
		t, err := expandFirewallVip6SslHstsAge(d, v, "ssl_hsts_age")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hsts-age"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hsts_include_subdomains"); ok || d.HasChange("ssl_hsts_include_subdomains") {
		t, err := expandFirewallVip6SslHstsIncludeSubdomains(d, v, "ssl_hsts_include_subdomains")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hsts-include-subdomains"] = t
		}
	}

	if v, ok := d.GetOk("ssl_http_location_conversion"); ok || d.HasChange("ssl_http_location_conversion") {
		t, err := expandFirewallVip6SslHttpLocationConversion(d, v, "ssl_http_location_conversion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-http-location-conversion"] = t
		}
	}

	if v, ok := d.GetOk("ssl_http_match_host"); ok || d.HasChange("ssl_http_match_host") {
		t, err := expandFirewallVip6SslHttpMatchHost(d, v, "ssl_http_match_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-http-match-host"] = t
		}
	}

	if v, ok := d.GetOk("ssl_max_version"); ok || d.HasChange("ssl_max_version") {
		t, err := expandFirewallVip6SslMaxVersion(d, v, "ssl_max_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-max-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_version"); ok || d.HasChange("ssl_min_version") {
		t, err := expandFirewallVip6SslMinVersion(d, v, "ssl_min_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_mode"); ok || d.HasChange("ssl_mode") {
		t, err := expandFirewallVip6SslMode(d, v, "ssl_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-mode"] = t
		}
	}

	if v, ok := d.GetOk("ssl_pfs"); ok || d.HasChange("ssl_pfs") {
		t, err := expandFirewallVip6SslPfs(d, v, "ssl_pfs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-pfs"] = t
		}
	}

	if v, ok := d.GetOk("ssl_send_empty_frags"); ok || d.HasChange("ssl_send_empty_frags") {
		t, err := expandFirewallVip6SslSendEmptyFrags(d, v, "ssl_send_empty_frags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-send-empty-frags"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_algorithm"); ok || d.HasChange("ssl_server_algorithm") {
		t, err := expandFirewallVip6SslServerAlgorithm(d, v, "ssl_server_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_cipher_suites"); ok || d.HasChange("ssl_server_cipher_suites") {
		t, err := expandFirewallVip6SslServerCipherSuites(d, v, "ssl_server_cipher_suites")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-cipher-suites"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_max_version"); ok || d.HasChange("ssl_server_max_version") {
		t, err := expandFirewallVip6SslServerMaxVersion(d, v, "ssl_server_max_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-max-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_min_version"); ok || d.HasChange("ssl_server_min_version") {
		t, err := expandFirewallVip6SslServerMinVersion(d, v, "ssl_server_min_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-min-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_renegotiation"); ok || d.HasChange("ssl_server_renegotiation") {
		t, err := expandFirewallVip6SslServerRenegotiation(d, v, "ssl_server_renegotiation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-renegotiation"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_session_state_max"); ok || d.HasChange("ssl_server_session_state_max") {
		t, err := expandFirewallVip6SslServerSessionStateMax(d, v, "ssl_server_session_state_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-session-state-max"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_session_state_timeout"); ok || d.HasChange("ssl_server_session_state_timeout") {
		t, err := expandFirewallVip6SslServerSessionStateTimeout(d, v, "ssl_server_session_state_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-session-state-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_session_state_type"); ok || d.HasChange("ssl_server_session_state_type") {
		t, err := expandFirewallVip6SslServerSessionStateType(d, v, "ssl_server_session_state_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-session-state-type"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandFirewallVip6Type(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("user_agent_detect"); ok || d.HasChange("user_agent_detect") {
		t, err := expandFirewallVip6UserAgentDetect(d, v, "user_agent_detect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-agent-detect"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandFirewallVip6Uuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("weblogic_server"); ok || d.HasChange("weblogic_server") {
		t, err := expandFirewallVip6WeblogicServer(d, v, "weblogic_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weblogic-server"] = t
		}
	}

	if v, ok := d.GetOk("websphere_server"); ok || d.HasChange("websphere_server") {
		t, err := expandFirewallVip6WebsphereServer(d, v, "websphere_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["websphere-server"] = t
		}
	}

	return &obj, nil
}
