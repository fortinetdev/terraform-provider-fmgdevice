// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure virtual IP for IPv4.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallVip() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallVipCreate,
		Read:   resourceFirewallVipRead,
		Update: resourceFirewallVipUpdate,
		Delete: resourceFirewallVipDelete,

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
			"add_nat46_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"arp_reply": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"dns_mapping_ttl": &schema.Schema{
				Type:     schema.TypeInt,
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
						"add_nat46_route": &schema.Schema{
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
						"dns_mapping_ttl": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"empty_cert_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"extaddr": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"extintf": &schema.Schema{
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
						"gratuitous_arp_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"gslb_domain_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"gslb_hostname": &schema.Schema{
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
						"http_multiplex_max_concurrent_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"http_multiplex_max_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"http_multiplex_ttl": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"http_redirect": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"http_supported_max_version": &schema.Schema{
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
						"ipv6_mappedip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipv6_mappedport": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ldb_method": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mapped_addr": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
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
						"nat44": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"nat46": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"one_click_gslb_server": &schema.Schema{
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
						"portmapping_type": &schema.Schema{
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
									"address": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"client_ip": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"health_check_proto": &schema.Schema{
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
									"seq": &schema.Schema{
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
									"type": &schema.Schema{
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
						"service": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
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
						"srcintf_filter": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
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
									"id": &schema.Schema{
										Type:     schema.TypeInt,
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
						"status": &schema.Schema{
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
			"empty_cert_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extaddr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"extintf": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
			},
			"gratuitous_arp_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"gslb_domain_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gslb_hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gslb_public_ips": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"index": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
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
			"http_multiplex_max_concurrent_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"http_multiplex_max_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"http_multiplex_ttl": &schema.Schema{
				Type:     schema.TypeInt,
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
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ipv6_mappedip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_mappedport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ldb_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mapped_addr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nat_source_vip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat44": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat46": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"one_click_gslb_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Computed: true,
			},
			"portmapping_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": &schema.Schema{
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
			"realservers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"client_ip": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
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
						"type": &schema.Schema{
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
						"health_check_proto": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"seq": &schema.Schema{
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
			"service": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
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
			"srcintf_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
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
			},
			"status": &schema.Schema{
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
			},
			"websphere_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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

func resourceFirewallVipCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallVip(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallVip resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallVip(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallVip(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallVip resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallVip(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallVip resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallVipRead(d, m)
}

func resourceFirewallVipUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallVip(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVip resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallVip(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVip resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallVipRead(d, m)
}

func resourceFirewallVipDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallVip(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallVip resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallVipRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallVip(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallVip resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallVip(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallVip resource from API: %v", err)
	}
	return nil
}

func flattenFirewallVipAddNat46Route(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipArpReply(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDnsMappingTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallVipDynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_nat46_route"
		if _, ok := i["add-nat46-route"]; ok {
			v := flattenFirewallVipDynamicMappingAddNat46Route(i["add-nat46-route"], d, pre_append)
			tmp["add_nat46_route"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-AddNat46Route")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "arp_reply"
		if _, ok := i["arp-reply"]; ok {
			v := flattenFirewallVipDynamicMappingArpReply(i["arp-reply"], d, pre_append)
			tmp["arp_reply"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-ArpReply")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_cert"
		if _, ok := i["client-cert"]; ok {
			v := flattenFirewallVipDynamicMappingClientCert(i["client-cert"], d, pre_append)
			tmp["client_cert"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-ClientCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color"
		if _, ok := i["color"]; ok {
			v := flattenFirewallVipDynamicMappingColor(i["color"], d, pre_append)
			tmp["color"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-Color")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenFirewallVipDynamicMappingComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_mapping_ttl"
		if _, ok := i["dns-mapping-ttl"]; ok {
			v := flattenFirewallVipDynamicMappingDnsMappingTtl(i["dns-mapping-ttl"], d, pre_append)
			tmp["dns_mapping_ttl"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-DnsMappingTtl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "empty_cert_action"
		if _, ok := i["empty-cert-action"]; ok {
			v := flattenFirewallVipDynamicMappingEmptyCertAction(i["empty-cert-action"], d, pre_append)
			tmp["empty_cert_action"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-EmptyCertAction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extaddr"
		if _, ok := i["extaddr"]; ok {
			v := flattenFirewallVipDynamicMappingExtaddr(i["extaddr"], d, pre_append)
			tmp["extaddr"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-Extaddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extintf"
		if _, ok := i["extintf"]; ok {
			v := flattenFirewallVipDynamicMappingExtintf(i["extintf"], d, pre_append)
			tmp["extintf"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-Extintf")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extip"
		if _, ok := i["extip"]; ok {
			v := flattenFirewallVipDynamicMappingExtip(i["extip"], d, pre_append)
			tmp["extip"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-Extip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extport"
		if _, ok := i["extport"]; ok {
			v := flattenFirewallVipDynamicMappingExtport(i["extport"], d, pre_append)
			tmp["extport"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-Extport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gratuitous_arp_interval"
		if _, ok := i["gratuitous-arp-interval"]; ok {
			v := flattenFirewallVipDynamicMappingGratuitousArpInterval(i["gratuitous-arp-interval"], d, pre_append)
			tmp["gratuitous_arp_interval"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-GratuitousArpInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gslb_domain_name"
		if _, ok := i["gslb-domain-name"]; ok {
			v := flattenFirewallVipDynamicMappingGslbDomainName(i["gslb-domain-name"], d, pre_append)
			tmp["gslb_domain_name"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-GslbDomainName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gslb_hostname"
		if _, ok := i["gslb-hostname"]; ok {
			v := flattenFirewallVipDynamicMappingGslbHostname(i["gslb-hostname"], d, pre_append)
			tmp["gslb_hostname"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-GslbHostname")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h2_support"
		if _, ok := i["h2-support"]; ok {
			v := flattenFirewallVipDynamicMappingH2Support(i["h2-support"], d, pre_append)
			tmp["h2_support"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-H2Support")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h3_support"
		if _, ok := i["h3-support"]; ok {
			v := flattenFirewallVipDynamicMappingH3Support(i["h3-support"], d, pre_append)
			tmp["h3_support"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-H3Support")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := i["http-cookie-age"]; ok {
			v := flattenFirewallVipDynamicMappingHttpCookieAge(i["http-cookie-age"], d, pre_append)
			tmp["http_cookie_age"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-HttpCookieAge")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := i["http-cookie-domain"]; ok {
			v := flattenFirewallVipDynamicMappingHttpCookieDomain(i["http-cookie-domain"], d, pre_append)
			tmp["http_cookie_domain"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-HttpCookieDomain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := i["http-cookie-domain-from-host"]; ok {
			v := flattenFirewallVipDynamicMappingHttpCookieDomainFromHost(i["http-cookie-domain-from-host"], d, pre_append)
			tmp["http_cookie_domain_from_host"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-HttpCookieDomainFromHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := i["http-cookie-generation"]; ok {
			v := flattenFirewallVipDynamicMappingHttpCookieGeneration(i["http-cookie-generation"], d, pre_append)
			tmp["http_cookie_generation"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-HttpCookieGeneration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := i["http-cookie-path"]; ok {
			v := flattenFirewallVipDynamicMappingHttpCookiePath(i["http-cookie-path"], d, pre_append)
			tmp["http_cookie_path"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-HttpCookiePath")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := i["http-cookie-share"]; ok {
			v := flattenFirewallVipDynamicMappingHttpCookieShare(i["http-cookie-share"], d, pre_append)
			tmp["http_cookie_share"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-HttpCookieShare")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_ip_header"
		if _, ok := i["http-ip-header"]; ok {
			v := flattenFirewallVipDynamicMappingHttpIpHeader(i["http-ip-header"], d, pre_append)
			tmp["http_ip_header"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-HttpIpHeader")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_ip_header_name"
		if _, ok := i["http-ip-header-name"]; ok {
			v := flattenFirewallVipDynamicMappingHttpIpHeaderName(i["http-ip-header-name"], d, pre_append)
			tmp["http_ip_header_name"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-HttpIpHeaderName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_multiplex"
		if _, ok := i["http-multiplex"]; ok {
			v := flattenFirewallVipDynamicMappingHttpMultiplex(i["http-multiplex"], d, pre_append)
			tmp["http_multiplex"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-HttpMultiplex")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_multiplex_max_concurrent_request"
		if _, ok := i["http-multiplex-max-concurrent-request"]; ok {
			v := flattenFirewallVipDynamicMappingHttpMultiplexMaxConcurrentRequest(i["http-multiplex-max-concurrent-request"], d, pre_append)
			tmp["http_multiplex_max_concurrent_request"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-HttpMultiplexMaxConcurrentRequest")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_multiplex_max_request"
		if _, ok := i["http-multiplex-max-request"]; ok {
			v := flattenFirewallVipDynamicMappingHttpMultiplexMaxRequest(i["http-multiplex-max-request"], d, pre_append)
			tmp["http_multiplex_max_request"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-HttpMultiplexMaxRequest")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_multiplex_ttl"
		if _, ok := i["http-multiplex-ttl"]; ok {
			v := flattenFirewallVipDynamicMappingHttpMultiplexTtl(i["http-multiplex-ttl"], d, pre_append)
			tmp["http_multiplex_ttl"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-HttpMultiplexTtl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_redirect"
		if _, ok := i["http-redirect"]; ok {
			v := flattenFirewallVipDynamicMappingHttpRedirect(i["http-redirect"], d, pre_append)
			tmp["http_redirect"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-HttpRedirect")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_supported_max_version"
		if _, ok := i["http-supported-max-version"]; ok {
			v := flattenFirewallVipDynamicMappingHttpSupportedMaxVersion(i["http-supported-max-version"], d, pre_append)
			tmp["http_supported_max_version"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-HttpSupportedMaxVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := i["https-cookie-secure"]; ok {
			v := flattenFirewallVipDynamicMappingHttpsCookieSecure(i["https-cookie-secure"], d, pre_append)
			tmp["https_cookie_secure"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-HttpsCookieSecure")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFirewallVipDynamicMappingId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_mappedip"
		if _, ok := i["ipv6-mappedip"]; ok {
			v := flattenFirewallVipDynamicMappingIpv6Mappedip(i["ipv6-mappedip"], d, pre_append)
			tmp["ipv6_mappedip"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-Ipv6Mappedip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_mappedport"
		if _, ok := i["ipv6-mappedport"]; ok {
			v := flattenFirewallVipDynamicMappingIpv6Mappedport(i["ipv6-mappedport"], d, pre_append)
			tmp["ipv6_mappedport"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-Ipv6Mappedport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := i["ldb-method"]; ok {
			v := flattenFirewallVipDynamicMappingLdbMethod(i["ldb-method"], d, pre_append)
			tmp["ldb_method"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-LdbMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mapped_addr"
		if _, ok := i["mapped-addr"]; ok {
			v := flattenFirewallVipDynamicMappingMappedAddr(i["mapped-addr"], d, pre_append)
			tmp["mapped_addr"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-MappedAddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedip"
		if _, ok := i["mappedip"]; ok {
			v := flattenFirewallVipDynamicMappingMappedip(i["mappedip"], d, pre_append)
			tmp["mappedip"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-Mappedip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := i["mappedport"]; ok {
			v := flattenFirewallVipDynamicMappingMappedport(i["mappedport"], d, pre_append)
			tmp["mappedport"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-Mappedport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_embryonic_connections"
		if _, ok := i["max-embryonic-connections"]; ok {
			v := flattenFirewallVipDynamicMappingMaxEmbryonicConnections(i["max-embryonic-connections"], d, pre_append)
			tmp["max_embryonic_connections"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-MaxEmbryonicConnections")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := i["monitor"]; ok {
			v := flattenFirewallVipDynamicMappingMonitor(i["monitor"], d, pre_append)
			tmp["monitor"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-Monitor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat_source_vip"
		if _, ok := i["nat-source-vip"]; ok {
			v := flattenFirewallVipDynamicMappingNatSourceVip(i["nat-source-vip"], d, pre_append)
			tmp["nat_source_vip"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-NatSourceVip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat44"
		if _, ok := i["nat44"]; ok {
			v := flattenFirewallVipDynamicMappingNat44(i["nat44"], d, pre_append)
			tmp["nat44"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-Nat44")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat46"
		if _, ok := i["nat46"]; ok {
			v := flattenFirewallVipDynamicMappingNat46(i["nat46"], d, pre_append)
			tmp["nat46"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-Nat46")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "one_click_gslb_server"
		if _, ok := i["one-click-gslb-server"]; ok {
			v := flattenFirewallVipDynamicMappingOneClickGslbServer(i["one-click-gslb-server"], d, pre_append)
			tmp["one_click_gslb_server"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-OneClickGslbServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "outlook_web_access"
		if _, ok := i["outlook-web-access"]; ok {
			v := flattenFirewallVipDynamicMappingOutlookWebAccess(i["outlook-web-access"], d, pre_append)
			tmp["outlook_web_access"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-OutlookWebAccess")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := i["persistence"]; ok {
			v := flattenFirewallVipDynamicMappingPersistence(i["persistence"], d, pre_append)
			tmp["persistence"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-Persistence")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "portforward"
		if _, ok := i["portforward"]; ok {
			v := flattenFirewallVipDynamicMappingPortforward(i["portforward"], d, pre_append)
			tmp["portforward"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-Portforward")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "portmapping_type"
		if _, ok := i["portmapping-type"]; ok {
			v := flattenFirewallVipDynamicMappingPortmappingType(i["portmapping-type"], d, pre_append)
			tmp["portmapping_type"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-PortmappingType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenFirewallVipDynamicMappingProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := i["realservers"]; ok {
			v := flattenFirewallVipDynamicMappingRealservers(i["realservers"], d, pre_append)
			tmp["realservers"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-Realservers")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_type"
		if _, ok := i["server-type"]; ok {
			v := flattenFirewallVipDynamicMappingServerType(i["server-type"], d, pre_append)
			tmp["server_type"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-ServerType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := i["service"]; ok {
			v := flattenFirewallVipDynamicMappingService(i["service"], d, pre_append)
			tmp["service"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-Service")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_filter"
		if _, ok := i["src-filter"]; ok {
			v := flattenFirewallVipDynamicMappingSrcFilter(i["src-filter"], d, pre_append)
			tmp["src_filter"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SrcFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_vip_filter"
		if _, ok := i["src-vip-filter"]; ok {
			v := flattenFirewallVipDynamicMappingSrcVipFilter(i["src-vip-filter"], d, pre_append)
			tmp["src_vip_filter"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SrcVipFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcintf_filter"
		if _, ok := i["srcintf-filter"]; ok {
			v := flattenFirewallVipDynamicMappingSrcintfFilter(i["srcintf-filter"], d, pre_append)
			tmp["srcintf_filter"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SrcintfFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_accept_ffdhe_groups"
		if _, ok := i["ssl-accept-ffdhe-groups"]; ok {
			v := flattenFirewallVipDynamicMappingSslAcceptFfdheGroups(i["ssl-accept-ffdhe-groups"], d, pre_append)
			tmp["ssl_accept_ffdhe_groups"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslAcceptFfdheGroups")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := i["ssl-algorithm"]; ok {
			v := flattenFirewallVipDynamicMappingSslAlgorithm(i["ssl-algorithm"], d, pre_append)
			tmp["ssl_algorithm"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslAlgorithm")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_certificate"
		if _, ok := i["ssl-certificate"]; ok {
			v := flattenFirewallVipDynamicMappingSslCertificate(i["ssl-certificate"], d, pre_append)
			tmp["ssl_certificate"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslCertificate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := i["ssl-cipher-suites"]; ok {
			v := flattenFirewallVipDynamicMappingSslCipherSuites(i["ssl-cipher-suites"], d, pre_append)
			tmp["ssl_cipher_suites"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslCipherSuites")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_fallback"
		if _, ok := i["ssl-client-fallback"]; ok {
			v := flattenFirewallVipDynamicMappingSslClientFallback(i["ssl-client-fallback"], d, pre_append)
			tmp["ssl_client_fallback"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslClientFallback")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_rekey_count"
		if _, ok := i["ssl-client-rekey-count"]; ok {
			v := flattenFirewallVipDynamicMappingSslClientRekeyCount(i["ssl-client-rekey-count"], d, pre_append)
			tmp["ssl_client_rekey_count"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslClientRekeyCount")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_renegotiation"
		if _, ok := i["ssl-client-renegotiation"]; ok {
			v := flattenFirewallVipDynamicMappingSslClientRenegotiation(i["ssl-client-renegotiation"], d, pre_append)
			tmp["ssl_client_renegotiation"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslClientRenegotiation")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_session_state_max"
		if _, ok := i["ssl-client-session-state-max"]; ok {
			v := flattenFirewallVipDynamicMappingSslClientSessionStateMax(i["ssl-client-session-state-max"], d, pre_append)
			tmp["ssl_client_session_state_max"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslClientSessionStateMax")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_session_state_timeout"
		if _, ok := i["ssl-client-session-state-timeout"]; ok {
			v := flattenFirewallVipDynamicMappingSslClientSessionStateTimeout(i["ssl-client-session-state-timeout"], d, pre_append)
			tmp["ssl_client_session_state_timeout"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslClientSessionStateTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_session_state_type"
		if _, ok := i["ssl-client-session-state-type"]; ok {
			v := flattenFirewallVipDynamicMappingSslClientSessionStateType(i["ssl-client-session-state-type"], d, pre_append)
			tmp["ssl_client_session_state_type"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslClientSessionStateType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := i["ssl-dh-bits"]; ok {
			v := flattenFirewallVipDynamicMappingSslDhBits(i["ssl-dh-bits"], d, pre_append)
			tmp["ssl_dh_bits"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslDhBits")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp"
		if _, ok := i["ssl-hpkp"]; ok {
			v := flattenFirewallVipDynamicMappingSslHpkp(i["ssl-hpkp"], d, pre_append)
			tmp["ssl_hpkp"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslHpkp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_age"
		if _, ok := i["ssl-hpkp-age"]; ok {
			v := flattenFirewallVipDynamicMappingSslHpkpAge(i["ssl-hpkp-age"], d, pre_append)
			tmp["ssl_hpkp_age"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslHpkpAge")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_backup"
		if _, ok := i["ssl-hpkp-backup"]; ok {
			v := flattenFirewallVipDynamicMappingSslHpkpBackup(i["ssl-hpkp-backup"], d, pre_append)
			tmp["ssl_hpkp_backup"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslHpkpBackup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_include_subdomains"
		if _, ok := i["ssl-hpkp-include-subdomains"]; ok {
			v := flattenFirewallVipDynamicMappingSslHpkpIncludeSubdomains(i["ssl-hpkp-include-subdomains"], d, pre_append)
			tmp["ssl_hpkp_include_subdomains"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslHpkpIncludeSubdomains")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_primary"
		if _, ok := i["ssl-hpkp-primary"]; ok {
			v := flattenFirewallVipDynamicMappingSslHpkpPrimary(i["ssl-hpkp-primary"], d, pre_append)
			tmp["ssl_hpkp_primary"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslHpkpPrimary")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_report_uri"
		if _, ok := i["ssl-hpkp-report-uri"]; ok {
			v := flattenFirewallVipDynamicMappingSslHpkpReportUri(i["ssl-hpkp-report-uri"], d, pre_append)
			tmp["ssl_hpkp_report_uri"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslHpkpReportUri")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hsts"
		if _, ok := i["ssl-hsts"]; ok {
			v := flattenFirewallVipDynamicMappingSslHsts(i["ssl-hsts"], d, pre_append)
			tmp["ssl_hsts"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslHsts")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hsts_age"
		if _, ok := i["ssl-hsts-age"]; ok {
			v := flattenFirewallVipDynamicMappingSslHstsAge(i["ssl-hsts-age"], d, pre_append)
			tmp["ssl_hsts_age"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslHstsAge")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hsts_include_subdomains"
		if _, ok := i["ssl-hsts-include-subdomains"]; ok {
			v := flattenFirewallVipDynamicMappingSslHstsIncludeSubdomains(i["ssl-hsts-include-subdomains"], d, pre_append)
			tmp["ssl_hsts_include_subdomains"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslHstsIncludeSubdomains")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_http_location_conversion"
		if _, ok := i["ssl-http-location-conversion"]; ok {
			v := flattenFirewallVipDynamicMappingSslHttpLocationConversion(i["ssl-http-location-conversion"], d, pre_append)
			tmp["ssl_http_location_conversion"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslHttpLocationConversion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_http_match_host"
		if _, ok := i["ssl-http-match-host"]; ok {
			v := flattenFirewallVipDynamicMappingSslHttpMatchHost(i["ssl-http-match-host"], d, pre_append)
			tmp["ssl_http_match_host"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslHttpMatchHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := i["ssl-max-version"]; ok {
			v := flattenFirewallVipDynamicMappingSslMaxVersion(i["ssl-max-version"], d, pre_append)
			tmp["ssl_max_version"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslMaxVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := i["ssl-min-version"]; ok {
			v := flattenFirewallVipDynamicMappingSslMinVersion(i["ssl-min-version"], d, pre_append)
			tmp["ssl_min_version"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslMinVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_mode"
		if _, ok := i["ssl-mode"]; ok {
			v := flattenFirewallVipDynamicMappingSslMode(i["ssl-mode"], d, pre_append)
			tmp["ssl_mode"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_pfs"
		if _, ok := i["ssl-pfs"]; ok {
			v := flattenFirewallVipDynamicMappingSslPfs(i["ssl-pfs"], d, pre_append)
			tmp["ssl_pfs"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslPfs")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_send_empty_frags"
		if _, ok := i["ssl-send-empty-frags"]; ok {
			v := flattenFirewallVipDynamicMappingSslSendEmptyFrags(i["ssl-send-empty-frags"], d, pre_append)
			tmp["ssl_send_empty_frags"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslSendEmptyFrags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_algorithm"
		if _, ok := i["ssl-server-algorithm"]; ok {
			v := flattenFirewallVipDynamicMappingSslServerAlgorithm(i["ssl-server-algorithm"], d, pre_append)
			tmp["ssl_server_algorithm"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslServerAlgorithm")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_max_version"
		if _, ok := i["ssl-server-max-version"]; ok {
			v := flattenFirewallVipDynamicMappingSslServerMaxVersion(i["ssl-server-max-version"], d, pre_append)
			tmp["ssl_server_max_version"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslServerMaxVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_min_version"
		if _, ok := i["ssl-server-min-version"]; ok {
			v := flattenFirewallVipDynamicMappingSslServerMinVersion(i["ssl-server-min-version"], d, pre_append)
			tmp["ssl_server_min_version"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslServerMinVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_renegotiation"
		if _, ok := i["ssl-server-renegotiation"]; ok {
			v := flattenFirewallVipDynamicMappingSslServerRenegotiation(i["ssl-server-renegotiation"], d, pre_append)
			tmp["ssl_server_renegotiation"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslServerRenegotiation")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_session_state_max"
		if _, ok := i["ssl-server-session-state-max"]; ok {
			v := flattenFirewallVipDynamicMappingSslServerSessionStateMax(i["ssl-server-session-state-max"], d, pre_append)
			tmp["ssl_server_session_state_max"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslServerSessionStateMax")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_session_state_timeout"
		if _, ok := i["ssl-server-session-state-timeout"]; ok {
			v := flattenFirewallVipDynamicMappingSslServerSessionStateTimeout(i["ssl-server-session-state-timeout"], d, pre_append)
			tmp["ssl_server_session_state_timeout"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslServerSessionStateTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_session_state_type"
		if _, ok := i["ssl-server-session-state-type"]; ok {
			v := flattenFirewallVipDynamicMappingSslServerSessionStateType(i["ssl-server-session-state-type"], d, pre_append)
			tmp["ssl_server_session_state_type"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-SslServerSessionStateType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenFirewallVipDynamicMappingStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenFirewallVipDynamicMappingType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_agent_detect"
		if _, ok := i["user-agent-detect"]; ok {
			v := flattenFirewallVipDynamicMappingUserAgentDetect(i["user-agent-detect"], d, pre_append)
			tmp["user_agent_detect"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-UserAgentDetect")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uuid"
		if _, ok := i["uuid"]; ok {
			v := flattenFirewallVipDynamicMappingUuid(i["uuid"], d, pre_append)
			tmp["uuid"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-Uuid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weblogic_server"
		if _, ok := i["weblogic-server"]; ok {
			v := flattenFirewallVipDynamicMappingWeblogicServer(i["weblogic-server"], d, pre_append)
			tmp["weblogic_server"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-WeblogicServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "websphere_server"
		if _, ok := i["websphere-server"]; ok {
			v := flattenFirewallVipDynamicMappingWebsphereServer(i["websphere-server"], d, pre_append)
			tmp["websphere_server"] = fortiAPISubPartPatch(v, "FirewallVip-DynamicMapping-WebsphereServer")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallVipDynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallVipDynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "FirewallVipDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenFirewallVipDynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "FirewallVipDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallVipDynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingAddNat46Route(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingArpReply(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingDnsMappingTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingEmptyCertAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingExtaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipDynamicMappingExtintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingExtip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipDynamicMappingExtport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingGratuitousArpInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingGslbDomainName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingGslbHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingH2Support(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingH3Support(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingHttpCookieAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingHttpCookieDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingHttpCookieDomainFromHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingHttpCookieGeneration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingHttpCookiePath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingHttpCookieShare(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingHttpIpHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingHttpIpHeaderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingHttpMultiplex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingHttpMultiplexMaxConcurrentRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingHttpMultiplexMaxRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingHttpMultiplexTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingHttpRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingHttpSupportedMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingHttpsCookieSecure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingIpv6Mappedip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingIpv6Mappedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingLdbMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingMappedAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipDynamicMappingMappedip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipDynamicMappingMappedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingMaxEmbryonicConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipDynamicMappingNatSourceVip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingNat44(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingNat46(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingOneClickGslbServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingOutlookWebAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingPersistence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingPortforward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingPortmappingType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingRealservers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {
			v := flattenFirewallVipDynamicMappingRealserversAddress(i["address"], d, pre_append)
			tmp["address"] = fortiAPISubPartPatch(v, "FirewallVipDynamicMapping-Realservers-Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_ip"
		if _, ok := i["client-ip"]; ok {
			v := flattenFirewallVipDynamicMappingRealserversClientIp(i["client-ip"], d, pre_append)
			tmp["client_ip"] = fortiAPISubPartPatch(v, "FirewallVipDynamicMapping-Realservers-ClientIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := i["health-check-proto"]; ok {
			v := flattenFirewallVipDynamicMappingRealserversHealthCheckProto(i["health-check-proto"], d, pre_append)
			tmp["health_check_proto"] = fortiAPISubPartPatch(v, "FirewallVipDynamicMapping-Realservers-HealthCheckProto")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "healthcheck"
		if _, ok := i["healthcheck"]; ok {
			v := flattenFirewallVipDynamicMappingRealserversHealthcheck(i["healthcheck"], d, pre_append)
			tmp["healthcheck"] = fortiAPISubPartPatch(v, "FirewallVipDynamicMapping-Realservers-Healthcheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := i["holddown-interval"]; ok {
			v := flattenFirewallVipDynamicMappingRealserversHolddownInterval(i["holddown-interval"], d, pre_append)
			tmp["holddown_interval"] = fortiAPISubPartPatch(v, "FirewallVipDynamicMapping-Realservers-HolddownInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := i["http-host"]; ok {
			v := flattenFirewallVipDynamicMappingRealserversHttpHost(i["http-host"], d, pre_append)
			tmp["http_host"] = fortiAPISubPartPatch(v, "FirewallVipDynamicMapping-Realservers-HttpHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFirewallVipDynamicMappingRealserversId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FirewallVipDynamicMapping-Realservers-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenFirewallVipDynamicMappingRealserversIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "FirewallVipDynamicMapping-Realservers-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_connections"
		if _, ok := i["max-connections"]; ok {
			v := flattenFirewallVipDynamicMappingRealserversMaxConnections(i["max-connections"], d, pre_append)
			tmp["max_connections"] = fortiAPISubPartPatch(v, "FirewallVipDynamicMapping-Realservers-MaxConnections")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := i["monitor"]; ok {
			v := flattenFirewallVipDynamicMappingRealserversMonitor(i["monitor"], d, pre_append)
			tmp["monitor"] = fortiAPISubPartPatch(v, "FirewallVipDynamicMapping-Realservers-Monitor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenFirewallVipDynamicMappingRealserversPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "FirewallVipDynamicMapping-Realservers-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq"
		if _, ok := i["seq"]; ok {
			v := flattenFirewallVipDynamicMappingRealserversSeq(i["seq"], d, pre_append)
			tmp["seq"] = fortiAPISubPartPatch(v, "FirewallVipDynamicMapping-Realservers-Seq")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenFirewallVipDynamicMappingRealserversStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "FirewallVipDynamicMapping-Realservers-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := i["translate-host"]; ok {
			v := flattenFirewallVipDynamicMappingRealserversTranslateHost(i["translate-host"], d, pre_append)
			tmp["translate_host"] = fortiAPISubPartPatch(v, "FirewallVipDynamicMapping-Realservers-TranslateHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenFirewallVipDynamicMappingRealserversType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "FirewallVipDynamicMapping-Realservers-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "verify_cert"
		if _, ok := i["verify-cert"]; ok {
			v := flattenFirewallVipDynamicMappingRealserversVerifyCert(i["verify-cert"], d, pre_append)
			tmp["verify_cert"] = fortiAPISubPartPatch(v, "FirewallVipDynamicMapping-Realservers-VerifyCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenFirewallVipDynamicMappingRealserversWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "FirewallVipDynamicMapping-Realservers-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallVipDynamicMappingRealserversAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipDynamicMappingRealserversClientIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipDynamicMappingRealserversHealthCheckProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingRealserversHealthcheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingRealserversHolddownInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingRealserversHttpHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingRealserversId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingRealserversIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingRealserversMaxConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingRealserversMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipDynamicMappingRealserversPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingRealserversSeq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingRealserversStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingRealserversTranslateHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingRealserversType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingRealserversVerifyCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingRealserversWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipDynamicMappingSrcFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipDynamicMappingSrcVipFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSrcintfFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipDynamicMappingSslAcceptFfdheGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipDynamicMappingSslCipherSuites(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallVipDynamicMappingSslCipherSuitesCipher(i["cipher"], d, pre_append)
			tmp["cipher"] = fortiAPISubPartPatch(v, "FirewallVipDynamicMapping-SslCipherSuites-Cipher")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFirewallVipDynamicMappingSslCipherSuitesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FirewallVipDynamicMapping-SslCipherSuites-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenFirewallVipDynamicMappingSslCipherSuitesPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "FirewallVipDynamicMapping-SslCipherSuites-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {
			v := flattenFirewallVipDynamicMappingSslCipherSuitesVersions(i["versions"], d, pre_append)
			tmp["versions"] = fortiAPISubPartPatch(v, "FirewallVipDynamicMapping-SslCipherSuites-Versions")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallVipDynamicMappingSslCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslCipherSuitesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipDynamicMappingSslClientFallback(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslClientRekeyCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslClientRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslClientSessionStateMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslClientSessionStateTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslClientSessionStateType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslDhBits(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslHpkp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslHpkpAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslHpkpBackup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipDynamicMappingSslHpkpIncludeSubdomains(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslHpkpPrimary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipDynamicMappingSslHpkpReportUri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslHsts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslHstsAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslHstsIncludeSubdomains(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslHttpLocationConversion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslHttpMatchHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslPfs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslSendEmptyFrags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslServerAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslServerMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslServerMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslServerRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslServerSessionStateMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslServerSessionStateTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingSslServerSessionStateType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingUserAgentDetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingWeblogicServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipDynamicMappingWebsphereServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipEmptyCertAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipExtaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipExtintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipExtip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipExtport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipGratuitousArpInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipGslbDomainName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipGslbHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipGslbPublicIps(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := i["index"]; ok {
			v := flattenFirewallVipGslbPublicIpsIndex(i["index"], d, pre_append)
			tmp["index"] = fortiAPISubPartPatch(v, "FirewallVip-GslbPublicIps-Index")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenFirewallVipGslbPublicIpsIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "FirewallVip-GslbPublicIps-Ip")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallVipGslbPublicIpsIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipGslbPublicIpsIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipH2Support(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipH3Support(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipHttpCookieAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipHttpCookieDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipHttpCookieDomainFromHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipHttpCookieGeneration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipHttpCookiePath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipHttpCookieShare(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipHttpIpHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipHttpIpHeaderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipHttpMultiplex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipHttpMultiplexMaxConcurrentRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipHttpMultiplexMaxRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipHttpMultiplexTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipHttpRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipHttpsCookieSecure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipIpv6Mappedip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipIpv6Mappedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipLdbMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipMappedAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipMappedip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipMappedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipMaxEmbryonicConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipNatSourceVip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipNat44(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipNat46(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipOneClickGslbServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipOutlookWebAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipPersistence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipPortforward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipPortmappingType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipQuic(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := i["ack-delay-exponent"]; ok {
		result["ack_delay_exponent"] = flattenFirewallVipQuicAckDelayExponent(i["ack-delay-exponent"], d, pre_append)
	}

	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := i["active-connection-id-limit"]; ok {
		result["active_connection_id_limit"] = flattenFirewallVipQuicActiveConnectionIdLimit(i["active-connection-id-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "active_migration"
	if _, ok := i["active-migration"]; ok {
		result["active_migration"] = flattenFirewallVipQuicActiveMigration(i["active-migration"], d, pre_append)
	}

	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := i["grease-quic-bit"]; ok {
		result["grease_quic_bit"] = flattenFirewallVipQuicGreaseQuicBit(i["grease-quic-bit"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := i["max-ack-delay"]; ok {
		result["max_ack_delay"] = flattenFirewallVipQuicMaxAckDelay(i["max-ack-delay"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := i["max-datagram-frame-size"]; ok {
		result["max_datagram_frame_size"] = flattenFirewallVipQuicMaxDatagramFrameSize(i["max-datagram-frame-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := i["max-idle-timeout"]; ok {
		result["max_idle_timeout"] = flattenFirewallVipQuicMaxIdleTimeout(i["max-idle-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := i["max-udp-payload-size"]; ok {
		result["max_udp_payload_size"] = flattenFirewallVipQuicMaxUdpPayloadSize(i["max-udp-payload-size"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallVipQuicAckDelayExponent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipQuicActiveConnectionIdLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipQuicActiveMigration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipQuicGreaseQuicBit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipQuicMaxAckDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipQuicMaxDatagramFrameSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipQuicMaxIdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipQuicMaxUdpPayloadSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipRealservers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {
			v := flattenFirewallVipRealserversAddress(i["address"], d, pre_append)
			tmp["address"] = fortiAPISubPartPatch(v, "FirewallVip-Realservers-Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_ip"
		if _, ok := i["client-ip"]; ok {
			v := flattenFirewallVipRealserversClientIp(i["client-ip"], d, pre_append)
			tmp["client_ip"] = fortiAPISubPartPatch(v, "FirewallVip-Realservers-ClientIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "healthcheck"
		if _, ok := i["healthcheck"]; ok {
			v := flattenFirewallVipRealserversHealthcheck(i["healthcheck"], d, pre_append)
			tmp["healthcheck"] = fortiAPISubPartPatch(v, "FirewallVip-Realservers-Healthcheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := i["holddown-interval"]; ok {
			v := flattenFirewallVipRealserversHolddownInterval(i["holddown-interval"], d, pre_append)
			tmp["holddown_interval"] = fortiAPISubPartPatch(v, "FirewallVip-Realservers-HolddownInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := i["http-host"]; ok {
			v := flattenFirewallVipRealserversHttpHost(i["http-host"], d, pre_append)
			tmp["http_host"] = fortiAPISubPartPatch(v, "FirewallVip-Realservers-HttpHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFirewallVipRealserversId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FirewallVip-Realservers-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenFirewallVipRealserversIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "FirewallVip-Realservers-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_connections"
		if _, ok := i["max-connections"]; ok {
			v := flattenFirewallVipRealserversMaxConnections(i["max-connections"], d, pre_append)
			tmp["max_connections"] = fortiAPISubPartPatch(v, "FirewallVip-Realservers-MaxConnections")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := i["monitor"]; ok {
			v := flattenFirewallVipRealserversMonitor(i["monitor"], d, pre_append)
			tmp["monitor"] = fortiAPISubPartPatch(v, "FirewallVip-Realservers-Monitor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenFirewallVipRealserversPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "FirewallVip-Realservers-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenFirewallVipRealserversStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "FirewallVip-Realservers-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := i["translate-host"]; ok {
			v := flattenFirewallVipRealserversTranslateHost(i["translate-host"], d, pre_append)
			tmp["translate_host"] = fortiAPISubPartPatch(v, "FirewallVip-Realservers-TranslateHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenFirewallVipRealserversType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "FirewallVip-Realservers-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "verify_cert"
		if _, ok := i["verify-cert"]; ok {
			v := flattenFirewallVipRealserversVerifyCert(i["verify-cert"], d, pre_append)
			tmp["verify_cert"] = fortiAPISubPartPatch(v, "FirewallVip-Realservers-VerifyCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenFirewallVipRealserversWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "FirewallVip-Realservers-Weight")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := i["health-check-proto"]; ok {
			v := flattenFirewallVipRealserversHealthCheckProto(i["health-check-proto"], d, pre_append)
			tmp["health_check_proto"] = fortiAPISubPartPatch(v, "FirewallVip-Realservers-HealthCheckProto")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq"
		if _, ok := i["seq"]; ok {
			v := flattenFirewallVipRealserversSeq(i["seq"], d, pre_append)
			tmp["seq"] = fortiAPISubPartPatch(v, "FirewallVip-Realservers-Seq")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallVipRealserversAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipRealserversClientIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipRealserversHealthcheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipRealserversHolddownInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipRealserversHttpHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipRealserversId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipRealserversIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipRealserversMaxConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipRealserversMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipRealserversPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipRealserversStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipRealserversTranslateHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipRealserversType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipRealserversVerifyCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipRealserversWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipRealserversHealthCheckProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipRealserversSeq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipSrcFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipSrcVipFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSrcintfFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipSslAcceptFfdheGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipSslCipherSuites(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallVipSslCipherSuitesCipher(i["cipher"], d, pre_append)
			tmp["cipher"] = fortiAPISubPartPatch(v, "FirewallVip-SslCipherSuites-Cipher")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenFirewallVipSslCipherSuitesPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "FirewallVip-SslCipherSuites-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {
			v := flattenFirewallVipSslCipherSuitesVersions(i["versions"], d, pre_append)
			tmp["versions"] = fortiAPISubPartPatch(v, "FirewallVip-SslCipherSuites-Versions")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFirewallVipSslCipherSuitesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FirewallVip-SslCipherSuites-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallVipSslCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipSslCipherSuitesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslClientFallback(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslClientRekeyCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslClientRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslClientSessionStateMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslClientSessionStateTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslClientSessionStateType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslDhBits(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslHpkp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslHpkpAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslHpkpBackup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipSslHpkpIncludeSubdomains(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslHpkpPrimary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipSslHpkpReportUri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslHsts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslHstsAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslHstsIncludeSubdomains(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslHttpLocationConversion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslHttpMatchHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslPfs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslSendEmptyFrags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslServerAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslServerCipherSuites(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallVipSslServerCipherSuitesCipher(i["cipher"], d, pre_append)
			tmp["cipher"] = fortiAPISubPartPatch(v, "FirewallVip-SslServerCipherSuites-Cipher")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenFirewallVipSslServerCipherSuitesPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "FirewallVip-SslServerCipherSuites-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {
			v := flattenFirewallVipSslServerCipherSuitesVersions(i["versions"], d, pre_append)
			tmp["versions"] = fortiAPISubPartPatch(v, "FirewallVip-SslServerCipherSuites-Versions")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallVipSslServerCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslServerCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslServerCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipSslServerMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslServerMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslServerRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslServerSessionStateMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslServerSessionStateTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslServerSessionStateType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipUserAgentDetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipWeblogicServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipWebsphereServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipHttpSupportedMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallVip(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("add_nat46_route", flattenFirewallVipAddNat46Route(o["add-nat46-route"], d, "add_nat46_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["add-nat46-route"], "FirewallVip-AddNat46Route"); ok {
			if err = d.Set("add_nat46_route", vv); err != nil {
				return fmt.Errorf("Error reading add_nat46_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading add_nat46_route: %v", err)
		}
	}

	if err = d.Set("arp_reply", flattenFirewallVipArpReply(o["arp-reply"], d, "arp_reply")); err != nil {
		if vv, ok := fortiAPIPatch(o["arp-reply"], "FirewallVip-ArpReply"); ok {
			if err = d.Set("arp_reply", vv); err != nil {
				return fmt.Errorf("Error reading arp_reply: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading arp_reply: %v", err)
		}
	}

	if err = d.Set("client_cert", flattenFirewallVipClientCert(o["client-cert"], d, "client_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-cert"], "FirewallVip-ClientCert"); ok {
			if err = d.Set("client_cert", vv); err != nil {
				return fmt.Errorf("Error reading client_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_cert: %v", err)
		}
	}

	if err = d.Set("color", flattenFirewallVipColor(o["color"], d, "color")); err != nil {
		if vv, ok := fortiAPIPatch(o["color"], "FirewallVip-Color"); ok {
			if err = d.Set("color", vv); err != nil {
				return fmt.Errorf("Error reading color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("comment", flattenFirewallVipComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "FirewallVip-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("dns_mapping_ttl", flattenFirewallVipDnsMappingTtl(o["dns-mapping-ttl"], d, "dns_mapping_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-mapping-ttl"], "FirewallVip-DnsMappingTtl"); ok {
			if err = d.Set("dns_mapping_ttl", vv); err != nil {
				return fmt.Errorf("Error reading dns_mapping_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_mapping_ttl: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenFirewallVipDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "FirewallVip-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenFirewallVipDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "FirewallVip-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("empty_cert_action", flattenFirewallVipEmptyCertAction(o["empty-cert-action"], d, "empty_cert_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["empty-cert-action"], "FirewallVip-EmptyCertAction"); ok {
			if err = d.Set("empty_cert_action", vv); err != nil {
				return fmt.Errorf("Error reading empty_cert_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading empty_cert_action: %v", err)
		}
	}

	if err = d.Set("extaddr", flattenFirewallVipExtaddr(o["extaddr"], d, "extaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["extaddr"], "FirewallVip-Extaddr"); ok {
			if err = d.Set("extaddr", vv); err != nil {
				return fmt.Errorf("Error reading extaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extaddr: %v", err)
		}
	}

	if err = d.Set("extintf", flattenFirewallVipExtintf(o["extintf"], d, "extintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["extintf"], "FirewallVip-Extintf"); ok {
			if err = d.Set("extintf", vv); err != nil {
				return fmt.Errorf("Error reading extintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extintf: %v", err)
		}
	}

	if err = d.Set("extip", flattenFirewallVipExtip(o["extip"], d, "extip")); err != nil {
		if vv, ok := fortiAPIPatch(o["extip"], "FirewallVip-Extip"); ok {
			if err = d.Set("extip", vv); err != nil {
				return fmt.Errorf("Error reading extip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extip: %v", err)
		}
	}

	if err = d.Set("extport", flattenFirewallVipExtport(o["extport"], d, "extport")); err != nil {
		if vv, ok := fortiAPIPatch(o["extport"], "FirewallVip-Extport"); ok {
			if err = d.Set("extport", vv); err != nil {
				return fmt.Errorf("Error reading extport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extport: %v", err)
		}
	}

	if err = d.Set("gratuitous_arp_interval", flattenFirewallVipGratuitousArpInterval(o["gratuitous-arp-interval"], d, "gratuitous_arp_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["gratuitous-arp-interval"], "FirewallVip-GratuitousArpInterval"); ok {
			if err = d.Set("gratuitous_arp_interval", vv); err != nil {
				return fmt.Errorf("Error reading gratuitous_arp_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gratuitous_arp_interval: %v", err)
		}
	}

	if err = d.Set("gslb_domain_name", flattenFirewallVipGslbDomainName(o["gslb-domain-name"], d, "gslb_domain_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["gslb-domain-name"], "FirewallVip-GslbDomainName"); ok {
			if err = d.Set("gslb_domain_name", vv); err != nil {
				return fmt.Errorf("Error reading gslb_domain_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gslb_domain_name: %v", err)
		}
	}

	if err = d.Set("gslb_hostname", flattenFirewallVipGslbHostname(o["gslb-hostname"], d, "gslb_hostname")); err != nil {
		if vv, ok := fortiAPIPatch(o["gslb-hostname"], "FirewallVip-GslbHostname"); ok {
			if err = d.Set("gslb_hostname", vv); err != nil {
				return fmt.Errorf("Error reading gslb_hostname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gslb_hostname: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("gslb_public_ips", flattenFirewallVipGslbPublicIps(o["gslb-public-ips"], d, "gslb_public_ips")); err != nil {
			if vv, ok := fortiAPIPatch(o["gslb-public-ips"], "FirewallVip-GslbPublicIps"); ok {
				if err = d.Set("gslb_public_ips", vv); err != nil {
					return fmt.Errorf("Error reading gslb_public_ips: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading gslb_public_ips: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("gslb_public_ips"); ok {
			if err = d.Set("gslb_public_ips", flattenFirewallVipGslbPublicIps(o["gslb-public-ips"], d, "gslb_public_ips")); err != nil {
				if vv, ok := fortiAPIPatch(o["gslb-public-ips"], "FirewallVip-GslbPublicIps"); ok {
					if err = d.Set("gslb_public_ips", vv); err != nil {
						return fmt.Errorf("Error reading gslb_public_ips: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading gslb_public_ips: %v", err)
				}
			}
		}
	}

	if err = d.Set("h2_support", flattenFirewallVipH2Support(o["h2-support"], d, "h2_support")); err != nil {
		if vv, ok := fortiAPIPatch(o["h2-support"], "FirewallVip-H2Support"); ok {
			if err = d.Set("h2_support", vv); err != nil {
				return fmt.Errorf("Error reading h2_support: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading h2_support: %v", err)
		}
	}

	if err = d.Set("h3_support", flattenFirewallVipH3Support(o["h3-support"], d, "h3_support")); err != nil {
		if vv, ok := fortiAPIPatch(o["h3-support"], "FirewallVip-H3Support"); ok {
			if err = d.Set("h3_support", vv); err != nil {
				return fmt.Errorf("Error reading h3_support: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading h3_support: %v", err)
		}
	}

	if err = d.Set("http_cookie_age", flattenFirewallVipHttpCookieAge(o["http-cookie-age"], d, "http_cookie_age")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-age"], "FirewallVip-HttpCookieAge"); ok {
			if err = d.Set("http_cookie_age", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_age: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_age: %v", err)
		}
	}

	if err = d.Set("http_cookie_domain", flattenFirewallVipHttpCookieDomain(o["http-cookie-domain"], d, "http_cookie_domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-domain"], "FirewallVip-HttpCookieDomain"); ok {
			if err = d.Set("http_cookie_domain", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_domain: %v", err)
		}
	}

	if err = d.Set("http_cookie_domain_from_host", flattenFirewallVipHttpCookieDomainFromHost(o["http-cookie-domain-from-host"], d, "http_cookie_domain_from_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-domain-from-host"], "FirewallVip-HttpCookieDomainFromHost"); ok {
			if err = d.Set("http_cookie_domain_from_host", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_domain_from_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_domain_from_host: %v", err)
		}
	}

	if err = d.Set("http_cookie_generation", flattenFirewallVipHttpCookieGeneration(o["http-cookie-generation"], d, "http_cookie_generation")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-generation"], "FirewallVip-HttpCookieGeneration"); ok {
			if err = d.Set("http_cookie_generation", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_generation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_generation: %v", err)
		}
	}

	if err = d.Set("http_cookie_path", flattenFirewallVipHttpCookiePath(o["http-cookie-path"], d, "http_cookie_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-path"], "FirewallVip-HttpCookiePath"); ok {
			if err = d.Set("http_cookie_path", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_path: %v", err)
		}
	}

	if err = d.Set("http_cookie_share", flattenFirewallVipHttpCookieShare(o["http-cookie-share"], d, "http_cookie_share")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-share"], "FirewallVip-HttpCookieShare"); ok {
			if err = d.Set("http_cookie_share", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_share: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_share: %v", err)
		}
	}

	if err = d.Set("http_ip_header", flattenFirewallVipHttpIpHeader(o["http-ip-header"], d, "http_ip_header")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-ip-header"], "FirewallVip-HttpIpHeader"); ok {
			if err = d.Set("http_ip_header", vv); err != nil {
				return fmt.Errorf("Error reading http_ip_header: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_ip_header: %v", err)
		}
	}

	if err = d.Set("http_ip_header_name", flattenFirewallVipHttpIpHeaderName(o["http-ip-header-name"], d, "http_ip_header_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-ip-header-name"], "FirewallVip-HttpIpHeaderName"); ok {
			if err = d.Set("http_ip_header_name", vv); err != nil {
				return fmt.Errorf("Error reading http_ip_header_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_ip_header_name: %v", err)
		}
	}

	if err = d.Set("http_multiplex", flattenFirewallVipHttpMultiplex(o["http-multiplex"], d, "http_multiplex")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-multiplex"], "FirewallVip-HttpMultiplex"); ok {
			if err = d.Set("http_multiplex", vv); err != nil {
				return fmt.Errorf("Error reading http_multiplex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_multiplex: %v", err)
		}
	}

	if err = d.Set("http_multiplex_max_concurrent_request", flattenFirewallVipHttpMultiplexMaxConcurrentRequest(o["http-multiplex-max-concurrent-request"], d, "http_multiplex_max_concurrent_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-multiplex-max-concurrent-request"], "FirewallVip-HttpMultiplexMaxConcurrentRequest"); ok {
			if err = d.Set("http_multiplex_max_concurrent_request", vv); err != nil {
				return fmt.Errorf("Error reading http_multiplex_max_concurrent_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_multiplex_max_concurrent_request: %v", err)
		}
	}

	if err = d.Set("http_multiplex_max_request", flattenFirewallVipHttpMultiplexMaxRequest(o["http-multiplex-max-request"], d, "http_multiplex_max_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-multiplex-max-request"], "FirewallVip-HttpMultiplexMaxRequest"); ok {
			if err = d.Set("http_multiplex_max_request", vv); err != nil {
				return fmt.Errorf("Error reading http_multiplex_max_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_multiplex_max_request: %v", err)
		}
	}

	if err = d.Set("http_multiplex_ttl", flattenFirewallVipHttpMultiplexTtl(o["http-multiplex-ttl"], d, "http_multiplex_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-multiplex-ttl"], "FirewallVip-HttpMultiplexTtl"); ok {
			if err = d.Set("http_multiplex_ttl", vv); err != nil {
				return fmt.Errorf("Error reading http_multiplex_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_multiplex_ttl: %v", err)
		}
	}

	if err = d.Set("http_redirect", flattenFirewallVipHttpRedirect(o["http-redirect"], d, "http_redirect")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-redirect"], "FirewallVip-HttpRedirect"); ok {
			if err = d.Set("http_redirect", vv); err != nil {
				return fmt.Errorf("Error reading http_redirect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_redirect: %v", err)
		}
	}

	if err = d.Set("https_cookie_secure", flattenFirewallVipHttpsCookieSecure(o["https-cookie-secure"], d, "https_cookie_secure")); err != nil {
		if vv, ok := fortiAPIPatch(o["https-cookie-secure"], "FirewallVip-HttpsCookieSecure"); ok {
			if err = d.Set("https_cookie_secure", vv); err != nil {
				return fmt.Errorf("Error reading https_cookie_secure: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading https_cookie_secure: %v", err)
		}
	}

	if err = d.Set("fosid", flattenFirewallVipId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FirewallVip-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ipv6_mappedip", flattenFirewallVipIpv6Mappedip(o["ipv6-mappedip"], d, "ipv6_mappedip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-mappedip"], "FirewallVip-Ipv6Mappedip"); ok {
			if err = d.Set("ipv6_mappedip", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_mappedip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_mappedip: %v", err)
		}
	}

	if err = d.Set("ipv6_mappedport", flattenFirewallVipIpv6Mappedport(o["ipv6-mappedport"], d, "ipv6_mappedport")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-mappedport"], "FirewallVip-Ipv6Mappedport"); ok {
			if err = d.Set("ipv6_mappedport", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_mappedport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_mappedport: %v", err)
		}
	}

	if err = d.Set("ldb_method", flattenFirewallVipLdbMethod(o["ldb-method"], d, "ldb_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldb-method"], "FirewallVip-LdbMethod"); ok {
			if err = d.Set("ldb_method", vv); err != nil {
				return fmt.Errorf("Error reading ldb_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldb_method: %v", err)
		}
	}

	if err = d.Set("mapped_addr", flattenFirewallVipMappedAddr(o["mapped-addr"], d, "mapped_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["mapped-addr"], "FirewallVip-MappedAddr"); ok {
			if err = d.Set("mapped_addr", vv); err != nil {
				return fmt.Errorf("Error reading mapped_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mapped_addr: %v", err)
		}
	}

	if err = d.Set("mappedip", flattenFirewallVipMappedip(o["mappedip"], d, "mappedip")); err != nil {
		if vv, ok := fortiAPIPatch(o["mappedip"], "FirewallVip-Mappedip"); ok {
			if err = d.Set("mappedip", vv); err != nil {
				return fmt.Errorf("Error reading mappedip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mappedip: %v", err)
		}
	}

	if err = d.Set("mappedport", flattenFirewallVipMappedport(o["mappedport"], d, "mappedport")); err != nil {
		if vv, ok := fortiAPIPatch(o["mappedport"], "FirewallVip-Mappedport"); ok {
			if err = d.Set("mappedport", vv); err != nil {
				return fmt.Errorf("Error reading mappedport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mappedport: %v", err)
		}
	}

	if err = d.Set("max_embryonic_connections", flattenFirewallVipMaxEmbryonicConnections(o["max-embryonic-connections"], d, "max_embryonic_connections")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-embryonic-connections"], "FirewallVip-MaxEmbryonicConnections"); ok {
			if err = d.Set("max_embryonic_connections", vv); err != nil {
				return fmt.Errorf("Error reading max_embryonic_connections: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_embryonic_connections: %v", err)
		}
	}

	if err = d.Set("monitor", flattenFirewallVipMonitor(o["monitor"], d, "monitor")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor"], "FirewallVip-Monitor"); ok {
			if err = d.Set("monitor", vv); err != nil {
				return fmt.Errorf("Error reading monitor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallVipName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallVip-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("nat_source_vip", flattenFirewallVipNatSourceVip(o["nat-source-vip"], d, "nat_source_vip")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat-source-vip"], "FirewallVip-NatSourceVip"); ok {
			if err = d.Set("nat_source_vip", vv); err != nil {
				return fmt.Errorf("Error reading nat_source_vip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat_source_vip: %v", err)
		}
	}

	if err = d.Set("nat44", flattenFirewallVipNat44(o["nat44"], d, "nat44")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat44"], "FirewallVip-Nat44"); ok {
			if err = d.Set("nat44", vv); err != nil {
				return fmt.Errorf("Error reading nat44: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat44: %v", err)
		}
	}

	if err = d.Set("nat46", flattenFirewallVipNat46(o["nat46"], d, "nat46")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat46"], "FirewallVip-Nat46"); ok {
			if err = d.Set("nat46", vv); err != nil {
				return fmt.Errorf("Error reading nat46: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat46: %v", err)
		}
	}

	if err = d.Set("one_click_gslb_server", flattenFirewallVipOneClickGslbServer(o["one-click-gslb-server"], d, "one_click_gslb_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["one-click-gslb-server"], "FirewallVip-OneClickGslbServer"); ok {
			if err = d.Set("one_click_gslb_server", vv); err != nil {
				return fmt.Errorf("Error reading one_click_gslb_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading one_click_gslb_server: %v", err)
		}
	}

	if err = d.Set("outlook_web_access", flattenFirewallVipOutlookWebAccess(o["outlook-web-access"], d, "outlook_web_access")); err != nil {
		if vv, ok := fortiAPIPatch(o["outlook-web-access"], "FirewallVip-OutlookWebAccess"); ok {
			if err = d.Set("outlook_web_access", vv); err != nil {
				return fmt.Errorf("Error reading outlook_web_access: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outlook_web_access: %v", err)
		}
	}

	if err = d.Set("persistence", flattenFirewallVipPersistence(o["persistence"], d, "persistence")); err != nil {
		if vv, ok := fortiAPIPatch(o["persistence"], "FirewallVip-Persistence"); ok {
			if err = d.Set("persistence", vv); err != nil {
				return fmt.Errorf("Error reading persistence: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading persistence: %v", err)
		}
	}

	if err = d.Set("portforward", flattenFirewallVipPortforward(o["portforward"], d, "portforward")); err != nil {
		if vv, ok := fortiAPIPatch(o["portforward"], "FirewallVip-Portforward"); ok {
			if err = d.Set("portforward", vv); err != nil {
				return fmt.Errorf("Error reading portforward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading portforward: %v", err)
		}
	}

	if err = d.Set("portmapping_type", flattenFirewallVipPortmappingType(o["portmapping-type"], d, "portmapping_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["portmapping-type"], "FirewallVip-PortmappingType"); ok {
			if err = d.Set("portmapping_type", vv); err != nil {
				return fmt.Errorf("Error reading portmapping_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading portmapping_type: %v", err)
		}
	}

	if err = d.Set("protocol", flattenFirewallVipProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "FirewallVip-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("quic", flattenFirewallVipQuic(o["quic"], d, "quic")); err != nil {
			if vv, ok := fortiAPIPatch(o["quic"], "FirewallVip-Quic"); ok {
				if err = d.Set("quic", vv); err != nil {
					return fmt.Errorf("Error reading quic: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading quic: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("quic"); ok {
			if err = d.Set("quic", flattenFirewallVipQuic(o["quic"], d, "quic")); err != nil {
				if vv, ok := fortiAPIPatch(o["quic"], "FirewallVip-Quic"); ok {
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
		if err = d.Set("realservers", flattenFirewallVipRealservers(o["realservers"], d, "realservers")); err != nil {
			if vv, ok := fortiAPIPatch(o["realservers"], "FirewallVip-Realservers"); ok {
				if err = d.Set("realservers", vv); err != nil {
					return fmt.Errorf("Error reading realservers: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading realservers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("realservers"); ok {
			if err = d.Set("realservers", flattenFirewallVipRealservers(o["realservers"], d, "realservers")); err != nil {
				if vv, ok := fortiAPIPatch(o["realservers"], "FirewallVip-Realservers"); ok {
					if err = d.Set("realservers", vv); err != nil {
						return fmt.Errorf("Error reading realservers: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading realservers: %v", err)
				}
			}
		}
	}

	if err = d.Set("server_type", flattenFirewallVipServerType(o["server-type"], d, "server_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-type"], "FirewallVip-ServerType"); ok {
			if err = d.Set("server_type", vv); err != nil {
				return fmt.Errorf("Error reading server_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if err = d.Set("service", flattenFirewallVipService(o["service"], d, "service")); err != nil {
		if vv, ok := fortiAPIPatch(o["service"], "FirewallVip-Service"); ok {
			if err = d.Set("service", vv); err != nil {
				return fmt.Errorf("Error reading service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("src_filter", flattenFirewallVipSrcFilter(o["src-filter"], d, "src_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-filter"], "FirewallVip-SrcFilter"); ok {
			if err = d.Set("src_filter", vv); err != nil {
				return fmt.Errorf("Error reading src_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_filter: %v", err)
		}
	}

	if err = d.Set("src_vip_filter", flattenFirewallVipSrcVipFilter(o["src-vip-filter"], d, "src_vip_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-vip-filter"], "FirewallVip-SrcVipFilter"); ok {
			if err = d.Set("src_vip_filter", vv); err != nil {
				return fmt.Errorf("Error reading src_vip_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_vip_filter: %v", err)
		}
	}

	if err = d.Set("srcintf_filter", flattenFirewallVipSrcintfFilter(o["srcintf-filter"], d, "srcintf_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcintf-filter"], "FirewallVip-SrcintfFilter"); ok {
			if err = d.Set("srcintf_filter", vv); err != nil {
				return fmt.Errorf("Error reading srcintf_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcintf_filter: %v", err)
		}
	}

	if err = d.Set("ssl_accept_ffdhe_groups", flattenFirewallVipSslAcceptFfdheGroups(o["ssl-accept-ffdhe-groups"], d, "ssl_accept_ffdhe_groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-accept-ffdhe-groups"], "FirewallVip-SslAcceptFfdheGroups"); ok {
			if err = d.Set("ssl_accept_ffdhe_groups", vv); err != nil {
				return fmt.Errorf("Error reading ssl_accept_ffdhe_groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_accept_ffdhe_groups: %v", err)
		}
	}

	if err = d.Set("ssl_algorithm", flattenFirewallVipSslAlgorithm(o["ssl-algorithm"], d, "ssl_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-algorithm"], "FirewallVip-SslAlgorithm"); ok {
			if err = d.Set("ssl_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading ssl_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_algorithm: %v", err)
		}
	}

	if err = d.Set("ssl_certificate", flattenFirewallVipSslCertificate(o["ssl-certificate"], d, "ssl_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-certificate"], "FirewallVip-SslCertificate"); ok {
			if err = d.Set("ssl_certificate", vv); err != nil {
				return fmt.Errorf("Error reading ssl_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_certificate: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ssl_cipher_suites", flattenFirewallVipSslCipherSuites(o["ssl-cipher-suites"], d, "ssl_cipher_suites")); err != nil {
			if vv, ok := fortiAPIPatch(o["ssl-cipher-suites"], "FirewallVip-SslCipherSuites"); ok {
				if err = d.Set("ssl_cipher_suites", vv); err != nil {
					return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_cipher_suites"); ok {
			if err = d.Set("ssl_cipher_suites", flattenFirewallVipSslCipherSuites(o["ssl-cipher-suites"], d, "ssl_cipher_suites")); err != nil {
				if vv, ok := fortiAPIPatch(o["ssl-cipher-suites"], "FirewallVip-SslCipherSuites"); ok {
					if err = d.Set("ssl_cipher_suites", vv); err != nil {
						return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_client_fallback", flattenFirewallVipSslClientFallback(o["ssl-client-fallback"], d, "ssl_client_fallback")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-fallback"], "FirewallVip-SslClientFallback"); ok {
			if err = d.Set("ssl_client_fallback", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_fallback: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_fallback: %v", err)
		}
	}

	if err = d.Set("ssl_client_rekey_count", flattenFirewallVipSslClientRekeyCount(o["ssl-client-rekey-count"], d, "ssl_client_rekey_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-rekey-count"], "FirewallVip-SslClientRekeyCount"); ok {
			if err = d.Set("ssl_client_rekey_count", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_rekey_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_rekey_count: %v", err)
		}
	}

	if err = d.Set("ssl_client_renegotiation", flattenFirewallVipSslClientRenegotiation(o["ssl-client-renegotiation"], d, "ssl_client_renegotiation")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-renegotiation"], "FirewallVip-SslClientRenegotiation"); ok {
			if err = d.Set("ssl_client_renegotiation", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_renegotiation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_renegotiation: %v", err)
		}
	}

	if err = d.Set("ssl_client_session_state_max", flattenFirewallVipSslClientSessionStateMax(o["ssl-client-session-state-max"], d, "ssl_client_session_state_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-session-state-max"], "FirewallVip-SslClientSessionStateMax"); ok {
			if err = d.Set("ssl_client_session_state_max", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_session_state_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_session_state_max: %v", err)
		}
	}

	if err = d.Set("ssl_client_session_state_timeout", flattenFirewallVipSslClientSessionStateTimeout(o["ssl-client-session-state-timeout"], d, "ssl_client_session_state_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-session-state-timeout"], "FirewallVip-SslClientSessionStateTimeout"); ok {
			if err = d.Set("ssl_client_session_state_timeout", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_session_state_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_session_state_timeout: %v", err)
		}
	}

	if err = d.Set("ssl_client_session_state_type", flattenFirewallVipSslClientSessionStateType(o["ssl-client-session-state-type"], d, "ssl_client_session_state_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-session-state-type"], "FirewallVip-SslClientSessionStateType"); ok {
			if err = d.Set("ssl_client_session_state_type", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_session_state_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_session_state_type: %v", err)
		}
	}

	if err = d.Set("ssl_dh_bits", flattenFirewallVipSslDhBits(o["ssl-dh-bits"], d, "ssl_dh_bits")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-dh-bits"], "FirewallVip-SslDhBits"); ok {
			if err = d.Set("ssl_dh_bits", vv); err != nil {
				return fmt.Errorf("Error reading ssl_dh_bits: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_dh_bits: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp", flattenFirewallVipSslHpkp(o["ssl-hpkp"], d, "ssl_hpkp")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hpkp"], "FirewallVip-SslHpkp"); ok {
			if err = d.Set("ssl_hpkp", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hpkp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hpkp: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_age", flattenFirewallVipSslHpkpAge(o["ssl-hpkp-age"], d, "ssl_hpkp_age")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hpkp-age"], "FirewallVip-SslHpkpAge"); ok {
			if err = d.Set("ssl_hpkp_age", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hpkp_age: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hpkp_age: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_backup", flattenFirewallVipSslHpkpBackup(o["ssl-hpkp-backup"], d, "ssl_hpkp_backup")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hpkp-backup"], "FirewallVip-SslHpkpBackup"); ok {
			if err = d.Set("ssl_hpkp_backup", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hpkp_backup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hpkp_backup: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_include_subdomains", flattenFirewallVipSslHpkpIncludeSubdomains(o["ssl-hpkp-include-subdomains"], d, "ssl_hpkp_include_subdomains")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hpkp-include-subdomains"], "FirewallVip-SslHpkpIncludeSubdomains"); ok {
			if err = d.Set("ssl_hpkp_include_subdomains", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hpkp_include_subdomains: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hpkp_include_subdomains: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_primary", flattenFirewallVipSslHpkpPrimary(o["ssl-hpkp-primary"], d, "ssl_hpkp_primary")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hpkp-primary"], "FirewallVip-SslHpkpPrimary"); ok {
			if err = d.Set("ssl_hpkp_primary", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hpkp_primary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hpkp_primary: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_report_uri", flattenFirewallVipSslHpkpReportUri(o["ssl-hpkp-report-uri"], d, "ssl_hpkp_report_uri")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hpkp-report-uri"], "FirewallVip-SslHpkpReportUri"); ok {
			if err = d.Set("ssl_hpkp_report_uri", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hpkp_report_uri: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hpkp_report_uri: %v", err)
		}
	}

	if err = d.Set("ssl_hsts", flattenFirewallVipSslHsts(o["ssl-hsts"], d, "ssl_hsts")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hsts"], "FirewallVip-SslHsts"); ok {
			if err = d.Set("ssl_hsts", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hsts: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hsts: %v", err)
		}
	}

	if err = d.Set("ssl_hsts_age", flattenFirewallVipSslHstsAge(o["ssl-hsts-age"], d, "ssl_hsts_age")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hsts-age"], "FirewallVip-SslHstsAge"); ok {
			if err = d.Set("ssl_hsts_age", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hsts_age: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hsts_age: %v", err)
		}
	}

	if err = d.Set("ssl_hsts_include_subdomains", flattenFirewallVipSslHstsIncludeSubdomains(o["ssl-hsts-include-subdomains"], d, "ssl_hsts_include_subdomains")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hsts-include-subdomains"], "FirewallVip-SslHstsIncludeSubdomains"); ok {
			if err = d.Set("ssl_hsts_include_subdomains", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hsts_include_subdomains: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hsts_include_subdomains: %v", err)
		}
	}

	if err = d.Set("ssl_http_location_conversion", flattenFirewallVipSslHttpLocationConversion(o["ssl-http-location-conversion"], d, "ssl_http_location_conversion")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-http-location-conversion"], "FirewallVip-SslHttpLocationConversion"); ok {
			if err = d.Set("ssl_http_location_conversion", vv); err != nil {
				return fmt.Errorf("Error reading ssl_http_location_conversion: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_http_location_conversion: %v", err)
		}
	}

	if err = d.Set("ssl_http_match_host", flattenFirewallVipSslHttpMatchHost(o["ssl-http-match-host"], d, "ssl_http_match_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-http-match-host"], "FirewallVip-SslHttpMatchHost"); ok {
			if err = d.Set("ssl_http_match_host", vv); err != nil {
				return fmt.Errorf("Error reading ssl_http_match_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_http_match_host: %v", err)
		}
	}

	if err = d.Set("ssl_max_version", flattenFirewallVipSslMaxVersion(o["ssl-max-version"], d, "ssl_max_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-max-version"], "FirewallVip-SslMaxVersion"); ok {
			if err = d.Set("ssl_max_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_max_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_max_version: %v", err)
		}
	}

	if err = d.Set("ssl_min_version", flattenFirewallVipSslMinVersion(o["ssl-min-version"], d, "ssl_min_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-min-version"], "FirewallVip-SslMinVersion"); ok {
			if err = d.Set("ssl_min_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_min_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_min_version: %v", err)
		}
	}

	if err = d.Set("ssl_mode", flattenFirewallVipSslMode(o["ssl-mode"], d, "ssl_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-mode"], "FirewallVip-SslMode"); ok {
			if err = d.Set("ssl_mode", vv); err != nil {
				return fmt.Errorf("Error reading ssl_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_mode: %v", err)
		}
	}

	if err = d.Set("ssl_pfs", flattenFirewallVipSslPfs(o["ssl-pfs"], d, "ssl_pfs")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-pfs"], "FirewallVip-SslPfs"); ok {
			if err = d.Set("ssl_pfs", vv); err != nil {
				return fmt.Errorf("Error reading ssl_pfs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_pfs: %v", err)
		}
	}

	if err = d.Set("ssl_send_empty_frags", flattenFirewallVipSslSendEmptyFrags(o["ssl-send-empty-frags"], d, "ssl_send_empty_frags")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-send-empty-frags"], "FirewallVip-SslSendEmptyFrags"); ok {
			if err = d.Set("ssl_send_empty_frags", vv); err != nil {
				return fmt.Errorf("Error reading ssl_send_empty_frags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_send_empty_frags: %v", err)
		}
	}

	if err = d.Set("ssl_server_algorithm", flattenFirewallVipSslServerAlgorithm(o["ssl-server-algorithm"], d, "ssl_server_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-algorithm"], "FirewallVip-SslServerAlgorithm"); ok {
			if err = d.Set("ssl_server_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_algorithm: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ssl_server_cipher_suites", flattenFirewallVipSslServerCipherSuites(o["ssl-server-cipher-suites"], d, "ssl_server_cipher_suites")); err != nil {
			if vv, ok := fortiAPIPatch(o["ssl-server-cipher-suites"], "FirewallVip-SslServerCipherSuites"); ok {
				if err = d.Set("ssl_server_cipher_suites", vv); err != nil {
					return fmt.Errorf("Error reading ssl_server_cipher_suites: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ssl_server_cipher_suites: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_server_cipher_suites"); ok {
			if err = d.Set("ssl_server_cipher_suites", flattenFirewallVipSslServerCipherSuites(o["ssl-server-cipher-suites"], d, "ssl_server_cipher_suites")); err != nil {
				if vv, ok := fortiAPIPatch(o["ssl-server-cipher-suites"], "FirewallVip-SslServerCipherSuites"); ok {
					if err = d.Set("ssl_server_cipher_suites", vv); err != nil {
						return fmt.Errorf("Error reading ssl_server_cipher_suites: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ssl_server_cipher_suites: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_server_max_version", flattenFirewallVipSslServerMaxVersion(o["ssl-server-max-version"], d, "ssl_server_max_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-max-version"], "FirewallVip-SslServerMaxVersion"); ok {
			if err = d.Set("ssl_server_max_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_max_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_max_version: %v", err)
		}
	}

	if err = d.Set("ssl_server_min_version", flattenFirewallVipSslServerMinVersion(o["ssl-server-min-version"], d, "ssl_server_min_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-min-version"], "FirewallVip-SslServerMinVersion"); ok {
			if err = d.Set("ssl_server_min_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_min_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_min_version: %v", err)
		}
	}

	if err = d.Set("ssl_server_renegotiation", flattenFirewallVipSslServerRenegotiation(o["ssl-server-renegotiation"], d, "ssl_server_renegotiation")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-renegotiation"], "FirewallVip-SslServerRenegotiation"); ok {
			if err = d.Set("ssl_server_renegotiation", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_renegotiation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_renegotiation: %v", err)
		}
	}

	if err = d.Set("ssl_server_session_state_max", flattenFirewallVipSslServerSessionStateMax(o["ssl-server-session-state-max"], d, "ssl_server_session_state_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-session-state-max"], "FirewallVip-SslServerSessionStateMax"); ok {
			if err = d.Set("ssl_server_session_state_max", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_session_state_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_session_state_max: %v", err)
		}
	}

	if err = d.Set("ssl_server_session_state_timeout", flattenFirewallVipSslServerSessionStateTimeout(o["ssl-server-session-state-timeout"], d, "ssl_server_session_state_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-session-state-timeout"], "FirewallVip-SslServerSessionStateTimeout"); ok {
			if err = d.Set("ssl_server_session_state_timeout", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_session_state_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_session_state_timeout: %v", err)
		}
	}

	if err = d.Set("ssl_server_session_state_type", flattenFirewallVipSslServerSessionStateType(o["ssl-server-session-state-type"], d, "ssl_server_session_state_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-session-state-type"], "FirewallVip-SslServerSessionStateType"); ok {
			if err = d.Set("ssl_server_session_state_type", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_session_state_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_session_state_type: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallVipStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FirewallVip-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("type", flattenFirewallVipType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "FirewallVip-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("user_agent_detect", flattenFirewallVipUserAgentDetect(o["user-agent-detect"], d, "user_agent_detect")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-agent-detect"], "FirewallVip-UserAgentDetect"); ok {
			if err = d.Set("user_agent_detect", vv); err != nil {
				return fmt.Errorf("Error reading user_agent_detect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_agent_detect: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallVipUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "FirewallVip-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("weblogic_server", flattenFirewallVipWeblogicServer(o["weblogic-server"], d, "weblogic_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["weblogic-server"], "FirewallVip-WeblogicServer"); ok {
			if err = d.Set("weblogic_server", vv); err != nil {
				return fmt.Errorf("Error reading weblogic_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weblogic_server: %v", err)
		}
	}

	if err = d.Set("websphere_server", flattenFirewallVipWebsphereServer(o["websphere-server"], d, "websphere_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["websphere-server"], "FirewallVip-WebsphereServer"); ok {
			if err = d.Set("websphere_server", vv); err != nil {
				return fmt.Errorf("Error reading websphere_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading websphere_server: %v", err)
		}
	}

	if err = d.Set("http_supported_max_version", flattenFirewallVipHttpSupportedMaxVersion(o["http-supported-max-version"], d, "http_supported_max_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-supported-max-version"], "FirewallVip-HttpSupportedMaxVersion"); ok {
			if err = d.Set("http_supported_max_version", vv); err != nil {
				return fmt.Errorf("Error reading http_supported_max_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_supported_max_version: %v", err)
		}
	}

	return nil
}

func flattenFirewallVipFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallVipAddNat46Route(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipArpReply(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipClientCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDnsMappingTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			t, err := expandFirewallVipDynamicMappingScope(d, i["_scope"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["_scope"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_nat46_route"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["add-nat46-route"], _ = expandFirewallVipDynamicMappingAddNat46Route(d, i["add_nat46_route"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "arp_reply"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["arp-reply"], _ = expandFirewallVipDynamicMappingArpReply(d, i["arp_reply"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["client-cert"], _ = expandFirewallVipDynamicMappingClientCert(d, i["client_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["color"], _ = expandFirewallVipDynamicMappingColor(d, i["color"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandFirewallVipDynamicMappingComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_mapping_ttl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dns-mapping-ttl"], _ = expandFirewallVipDynamicMappingDnsMappingTtl(d, i["dns_mapping_ttl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "empty_cert_action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["empty-cert-action"], _ = expandFirewallVipDynamicMappingEmptyCertAction(d, i["empty_cert_action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extaddr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["extaddr"], _ = expandFirewallVipDynamicMappingExtaddr(d, i["extaddr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extintf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["extintf"], _ = expandFirewallVipDynamicMappingExtintf(d, i["extintf"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["extip"], _ = expandFirewallVipDynamicMappingExtip(d, i["extip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["extport"], _ = expandFirewallVipDynamicMappingExtport(d, i["extport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gratuitous_arp_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gratuitous-arp-interval"], _ = expandFirewallVipDynamicMappingGratuitousArpInterval(d, i["gratuitous_arp_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gslb_domain_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gslb-domain-name"], _ = expandFirewallVipDynamicMappingGslbDomainName(d, i["gslb_domain_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gslb_hostname"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gslb-hostname"], _ = expandFirewallVipDynamicMappingGslbHostname(d, i["gslb_hostname"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h2_support"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["h2-support"], _ = expandFirewallVipDynamicMappingH2Support(d, i["h2_support"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h3_support"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["h3-support"], _ = expandFirewallVipDynamicMappingH3Support(d, i["h3_support"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-age"], _ = expandFirewallVipDynamicMappingHttpCookieAge(d, i["http_cookie_age"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-domain"], _ = expandFirewallVipDynamicMappingHttpCookieDomain(d, i["http_cookie_domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-domain-from-host"], _ = expandFirewallVipDynamicMappingHttpCookieDomainFromHost(d, i["http_cookie_domain_from_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-generation"], _ = expandFirewallVipDynamicMappingHttpCookieGeneration(d, i["http_cookie_generation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-path"], _ = expandFirewallVipDynamicMappingHttpCookiePath(d, i["http_cookie_path"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-share"], _ = expandFirewallVipDynamicMappingHttpCookieShare(d, i["http_cookie_share"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_ip_header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-ip-header"], _ = expandFirewallVipDynamicMappingHttpIpHeader(d, i["http_ip_header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_ip_header_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-ip-header-name"], _ = expandFirewallVipDynamicMappingHttpIpHeaderName(d, i["http_ip_header_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_multiplex"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-multiplex"], _ = expandFirewallVipDynamicMappingHttpMultiplex(d, i["http_multiplex"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_multiplex_max_concurrent_request"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-multiplex-max-concurrent-request"], _ = expandFirewallVipDynamicMappingHttpMultiplexMaxConcurrentRequest(d, i["http_multiplex_max_concurrent_request"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_multiplex_max_request"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-multiplex-max-request"], _ = expandFirewallVipDynamicMappingHttpMultiplexMaxRequest(d, i["http_multiplex_max_request"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_multiplex_ttl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-multiplex-ttl"], _ = expandFirewallVipDynamicMappingHttpMultiplexTtl(d, i["http_multiplex_ttl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_redirect"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-redirect"], _ = expandFirewallVipDynamicMappingHttpRedirect(d, i["http_redirect"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_supported_max_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-supported-max-version"], _ = expandFirewallVipDynamicMappingHttpSupportedMaxVersion(d, i["http_supported_max_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["https-cookie-secure"], _ = expandFirewallVipDynamicMappingHttpsCookieSecure(d, i["https_cookie_secure"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFirewallVipDynamicMappingId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_mappedip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipv6-mappedip"], _ = expandFirewallVipDynamicMappingIpv6Mappedip(d, i["ipv6_mappedip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_mappedport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipv6-mappedport"], _ = expandFirewallVipDynamicMappingIpv6Mappedport(d, i["ipv6_mappedport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ldb-method"], _ = expandFirewallVipDynamicMappingLdbMethod(d, i["ldb_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mapped_addr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mapped-addr"], _ = expandFirewallVipDynamicMappingMappedAddr(d, i["mapped_addr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mappedip"], _ = expandFirewallVipDynamicMappingMappedip(d, i["mappedip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mappedport"], _ = expandFirewallVipDynamicMappingMappedport(d, i["mappedport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_embryonic_connections"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-embryonic-connections"], _ = expandFirewallVipDynamicMappingMaxEmbryonicConnections(d, i["max_embryonic_connections"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["monitor"], _ = expandFirewallVipDynamicMappingMonitor(d, i["monitor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat_source_vip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nat-source-vip"], _ = expandFirewallVipDynamicMappingNatSourceVip(d, i["nat_source_vip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat44"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nat44"], _ = expandFirewallVipDynamicMappingNat44(d, i["nat44"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat46"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nat46"], _ = expandFirewallVipDynamicMappingNat46(d, i["nat46"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "one_click_gslb_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["one-click-gslb-server"], _ = expandFirewallVipDynamicMappingOneClickGslbServer(d, i["one_click_gslb_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "outlook_web_access"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["outlook-web-access"], _ = expandFirewallVipDynamicMappingOutlookWebAccess(d, i["outlook_web_access"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["persistence"], _ = expandFirewallVipDynamicMappingPersistence(d, i["persistence"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "portforward"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["portforward"], _ = expandFirewallVipDynamicMappingPortforward(d, i["portforward"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "portmapping_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["portmapping-type"], _ = expandFirewallVipDynamicMappingPortmappingType(d, i["portmapping_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandFirewallVipDynamicMappingProtocol(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandFirewallVipDynamicMappingRealservers(d, i["realservers"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["realservers"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server-type"], _ = expandFirewallVipDynamicMappingServerType(d, i["server_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service"], _ = expandFirewallVipDynamicMappingService(d, i["service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-filter"], _ = expandFirewallVipDynamicMappingSrcFilter(d, i["src_filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_vip_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-vip-filter"], _ = expandFirewallVipDynamicMappingSrcVipFilter(d, i["src_vip_filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcintf_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["srcintf-filter"], _ = expandFirewallVipDynamicMappingSrcintfFilter(d, i["srcintf_filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_accept_ffdhe_groups"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-accept-ffdhe-groups"], _ = expandFirewallVipDynamicMappingSslAcceptFfdheGroups(d, i["ssl_accept_ffdhe_groups"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-algorithm"], _ = expandFirewallVipDynamicMappingSslAlgorithm(d, i["ssl_algorithm"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_certificate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-certificate"], _ = expandFirewallVipDynamicMappingSslCertificate(d, i["ssl_certificate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandFirewallVipDynamicMappingSslCipherSuites(d, i["ssl_cipher_suites"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["ssl-cipher-suites"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_fallback"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-client-fallback"], _ = expandFirewallVipDynamicMappingSslClientFallback(d, i["ssl_client_fallback"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_rekey_count"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-client-rekey-count"], _ = expandFirewallVipDynamicMappingSslClientRekeyCount(d, i["ssl_client_rekey_count"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_renegotiation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-client-renegotiation"], _ = expandFirewallVipDynamicMappingSslClientRenegotiation(d, i["ssl_client_renegotiation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_session_state_max"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-client-session-state-max"], _ = expandFirewallVipDynamicMappingSslClientSessionStateMax(d, i["ssl_client_session_state_max"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_session_state_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-client-session-state-timeout"], _ = expandFirewallVipDynamicMappingSslClientSessionStateTimeout(d, i["ssl_client_session_state_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_session_state_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-client-session-state-type"], _ = expandFirewallVipDynamicMappingSslClientSessionStateType(d, i["ssl_client_session_state_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-dh-bits"], _ = expandFirewallVipDynamicMappingSslDhBits(d, i["ssl_dh_bits"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hpkp"], _ = expandFirewallVipDynamicMappingSslHpkp(d, i["ssl_hpkp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_age"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hpkp-age"], _ = expandFirewallVipDynamicMappingSslHpkpAge(d, i["ssl_hpkp_age"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_backup"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hpkp-backup"], _ = expandFirewallVipDynamicMappingSslHpkpBackup(d, i["ssl_hpkp_backup"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_include_subdomains"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hpkp-include-subdomains"], _ = expandFirewallVipDynamicMappingSslHpkpIncludeSubdomains(d, i["ssl_hpkp_include_subdomains"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_primary"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hpkp-primary"], _ = expandFirewallVipDynamicMappingSslHpkpPrimary(d, i["ssl_hpkp_primary"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_report_uri"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hpkp-report-uri"], _ = expandFirewallVipDynamicMappingSslHpkpReportUri(d, i["ssl_hpkp_report_uri"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hsts"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hsts"], _ = expandFirewallVipDynamicMappingSslHsts(d, i["ssl_hsts"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hsts_age"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hsts-age"], _ = expandFirewallVipDynamicMappingSslHstsAge(d, i["ssl_hsts_age"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hsts_include_subdomains"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hsts-include-subdomains"], _ = expandFirewallVipDynamicMappingSslHstsIncludeSubdomains(d, i["ssl_hsts_include_subdomains"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_http_location_conversion"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-http-location-conversion"], _ = expandFirewallVipDynamicMappingSslHttpLocationConversion(d, i["ssl_http_location_conversion"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_http_match_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-http-match-host"], _ = expandFirewallVipDynamicMappingSslHttpMatchHost(d, i["ssl_http_match_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-max-version"], _ = expandFirewallVipDynamicMappingSslMaxVersion(d, i["ssl_max_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-min-version"], _ = expandFirewallVipDynamicMappingSslMinVersion(d, i["ssl_min_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-mode"], _ = expandFirewallVipDynamicMappingSslMode(d, i["ssl_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_pfs"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-pfs"], _ = expandFirewallVipDynamicMappingSslPfs(d, i["ssl_pfs"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_send_empty_frags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-send-empty-frags"], _ = expandFirewallVipDynamicMappingSslSendEmptyFrags(d, i["ssl_send_empty_frags"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_algorithm"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-server-algorithm"], _ = expandFirewallVipDynamicMappingSslServerAlgorithm(d, i["ssl_server_algorithm"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_max_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-server-max-version"], _ = expandFirewallVipDynamicMappingSslServerMaxVersion(d, i["ssl_server_max_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_min_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-server-min-version"], _ = expandFirewallVipDynamicMappingSslServerMinVersion(d, i["ssl_server_min_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_renegotiation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-server-renegotiation"], _ = expandFirewallVipDynamicMappingSslServerRenegotiation(d, i["ssl_server_renegotiation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_session_state_max"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-server-session-state-max"], _ = expandFirewallVipDynamicMappingSslServerSessionStateMax(d, i["ssl_server_session_state_max"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_session_state_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-server-session-state-timeout"], _ = expandFirewallVipDynamicMappingSslServerSessionStateTimeout(d, i["ssl_server_session_state_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_session_state_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-server-session-state-type"], _ = expandFirewallVipDynamicMappingSslServerSessionStateType(d, i["ssl_server_session_state_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandFirewallVipDynamicMappingStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandFirewallVipDynamicMappingType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_agent_detect"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-agent-detect"], _ = expandFirewallVipDynamicMappingUserAgentDetect(d, i["user_agent_detect"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uuid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uuid"], _ = expandFirewallVipDynamicMappingUuid(d, i["uuid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weblogic_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weblogic-server"], _ = expandFirewallVipDynamicMappingWeblogicServer(d, i["weblogic_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "websphere_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["websphere-server"], _ = expandFirewallVipDynamicMappingWebsphereServer(d, i["websphere_server"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallVipDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallVipDynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandFirewallVipDynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallVipDynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingAddNat46Route(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingArpReply(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingClientCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingDnsMappingTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingEmptyCertAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingExtaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipDynamicMappingExtintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingExtip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipDynamicMappingExtport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingGratuitousArpInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingGslbDomainName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingGslbHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingH2Support(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingH3Support(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingHttpCookieAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingHttpCookieDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingHttpCookieDomainFromHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingHttpCookieGeneration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingHttpCookiePath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingHttpCookieShare(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingHttpIpHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingHttpIpHeaderName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingHttpMultiplex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingHttpMultiplexMaxConcurrentRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingHttpMultiplexMaxRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingHttpMultiplexTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingHttpRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingHttpSupportedMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingHttpsCookieSecure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingIpv6Mappedip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingIpv6Mappedport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingLdbMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingMappedAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipDynamicMappingMappedip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipDynamicMappingMappedport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingMaxEmbryonicConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipDynamicMappingNatSourceVip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingNat44(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingNat46(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingOneClickGslbServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingOutlookWebAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingPersistence(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingPortforward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingPortmappingType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingRealservers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["address"], _ = expandFirewallVipDynamicMappingRealserversAddress(d, i["address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["client-ip"], _ = expandFirewallVipDynamicMappingRealserversClientIp(d, i["client_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check-proto"], _ = expandFirewallVipDynamicMappingRealserversHealthCheckProto(d, i["health_check_proto"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "healthcheck"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["healthcheck"], _ = expandFirewallVipDynamicMappingRealserversHealthcheck(d, i["healthcheck"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["holddown-interval"], _ = expandFirewallVipDynamicMappingRealserversHolddownInterval(d, i["holddown_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-host"], _ = expandFirewallVipDynamicMappingRealserversHttpHost(d, i["http_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFirewallVipDynamicMappingRealserversId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandFirewallVipDynamicMappingRealserversIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_connections"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-connections"], _ = expandFirewallVipDynamicMappingRealserversMaxConnections(d, i["max_connections"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["monitor"], _ = expandFirewallVipDynamicMappingRealserversMonitor(d, i["monitor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandFirewallVipDynamicMappingRealserversPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["seq"], _ = expandFirewallVipDynamicMappingRealserversSeq(d, i["seq"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandFirewallVipDynamicMappingRealserversStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["translate-host"], _ = expandFirewallVipDynamicMappingRealserversTranslateHost(d, i["translate_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandFirewallVipDynamicMappingRealserversType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "verify_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["verify-cert"], _ = expandFirewallVipDynamicMappingRealserversVerifyCert(d, i["verify_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandFirewallVipDynamicMappingRealserversWeight(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallVipDynamicMappingRealserversAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipDynamicMappingRealserversClientIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipDynamicMappingRealserversHealthCheckProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingRealserversHealthcheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingRealserversHolddownInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingRealserversHttpHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingRealserversId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingRealserversIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingRealserversMaxConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingRealserversMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipDynamicMappingRealserversPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingRealserversSeq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingRealserversStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingRealserversTranslateHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingRealserversType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingRealserversVerifyCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingRealserversWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipDynamicMappingSrcFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipDynamicMappingSrcVipFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSrcintfFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipDynamicMappingSslAcceptFfdheGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipDynamicMappingSslCipherSuites(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["cipher"], _ = expandFirewallVipDynamicMappingSslCipherSuitesCipher(d, i["cipher"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFirewallVipDynamicMappingSslCipherSuitesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandFirewallVipDynamicMappingSslCipherSuitesPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["versions"], _ = expandFirewallVipDynamicMappingSslCipherSuitesVersions(d, i["versions"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallVipDynamicMappingSslCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslCipherSuitesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipDynamicMappingSslClientFallback(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslClientRekeyCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslClientRenegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslClientSessionStateMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslClientSessionStateTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslClientSessionStateType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslDhBits(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslHpkp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslHpkpAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslHpkpBackup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipDynamicMappingSslHpkpIncludeSubdomains(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslHpkpPrimary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipDynamicMappingSslHpkpReportUri(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslHsts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslHstsAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslHstsIncludeSubdomains(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslHttpLocationConversion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslHttpMatchHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslPfs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslSendEmptyFrags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslServerAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslServerMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslServerMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslServerRenegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslServerSessionStateMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslServerSessionStateTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingSslServerSessionStateType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingUserAgentDetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingWeblogicServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDynamicMappingWebsphereServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipEmptyCertAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipExtaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipExtintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipExtip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipExtport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipGratuitousArpInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipGslbDomainName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipGslbHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipGslbPublicIps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["index"], _ = expandFirewallVipGslbPublicIpsIndex(d, i["index"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandFirewallVipGslbPublicIpsIp(d, i["ip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallVipGslbPublicIpsIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipGslbPublicIpsIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipH2Support(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipH3Support(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipHttpCookieAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipHttpCookieDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipHttpCookieDomainFromHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipHttpCookieGeneration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipHttpCookiePath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipHttpCookieShare(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipHttpIpHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipHttpIpHeaderName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipHttpMultiplex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipHttpMultiplexMaxConcurrentRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipHttpMultiplexMaxRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipHttpMultiplexTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipHttpRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipHttpsCookieSecure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipIpv6Mappedip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipIpv6Mappedport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipLdbMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipMappedAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipMappedip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipMappedport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipMaxEmbryonicConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipNatSourceVip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipNat44(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipNat46(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipOneClickGslbServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipOutlookWebAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipPersistence(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipPortforward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipPortmappingType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipQuic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ack-delay-exponent"], _ = expandFirewallVipQuicAckDelayExponent(d, i["ack_delay_exponent"], pre_append)
	}
	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["active-connection-id-limit"], _ = expandFirewallVipQuicActiveConnectionIdLimit(d, i["active_connection_id_limit"], pre_append)
	}
	pre_append = pre + ".0." + "active_migration"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["active-migration"], _ = expandFirewallVipQuicActiveMigration(d, i["active_migration"], pre_append)
	}
	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["grease-quic-bit"], _ = expandFirewallVipQuicGreaseQuicBit(d, i["grease_quic_bit"], pre_append)
	}
	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-ack-delay"], _ = expandFirewallVipQuicMaxAckDelay(d, i["max_ack_delay"], pre_append)
	}
	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-datagram-frame-size"], _ = expandFirewallVipQuicMaxDatagramFrameSize(d, i["max_datagram_frame_size"], pre_append)
	}
	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-idle-timeout"], _ = expandFirewallVipQuicMaxIdleTimeout(d, i["max_idle_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-udp-payload-size"], _ = expandFirewallVipQuicMaxUdpPayloadSize(d, i["max_udp_payload_size"], pre_append)
	}

	return result, nil
}

func expandFirewallVipQuicAckDelayExponent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipQuicActiveConnectionIdLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipQuicActiveMigration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipQuicGreaseQuicBit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipQuicMaxAckDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipQuicMaxDatagramFrameSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipQuicMaxIdleTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipQuicMaxUdpPayloadSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealservers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["address"], _ = expandFirewallVipRealserversAddress(d, i["address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["client-ip"], _ = expandFirewallVipRealserversClientIp(d, i["client_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "healthcheck"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["healthcheck"], _ = expandFirewallVipRealserversHealthcheck(d, i["healthcheck"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["holddown-interval"], _ = expandFirewallVipRealserversHolddownInterval(d, i["holddown_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-host"], _ = expandFirewallVipRealserversHttpHost(d, i["http_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFirewallVipRealserversId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandFirewallVipRealserversIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_connections"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-connections"], _ = expandFirewallVipRealserversMaxConnections(d, i["max_connections"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["monitor"], _ = expandFirewallVipRealserversMonitor(d, i["monitor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandFirewallVipRealserversPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandFirewallVipRealserversStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["translate-host"], _ = expandFirewallVipRealserversTranslateHost(d, i["translate_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandFirewallVipRealserversType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "verify_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["verify-cert"], _ = expandFirewallVipRealserversVerifyCert(d, i["verify_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandFirewallVipRealserversWeight(d, i["weight"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check-proto"], _ = expandFirewallVipRealserversHealthCheckProto(d, i["health_check_proto"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["seq"], _ = expandFirewallVipRealserversSeq(d, i["seq"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallVipRealserversAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipRealserversClientIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipRealserversHealthcheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversHolddownInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversHttpHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversMaxConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipRealserversPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversTranslateHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversVerifyCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversHealthCheckProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversSeq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipSrcFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipSrcVipFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSrcintfFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipSslAcceptFfdheGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipSslCipherSuites(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["cipher"], _ = expandFirewallVipSslCipherSuitesCipher(d, i["cipher"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandFirewallVipSslCipherSuitesPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["versions"], _ = expandFirewallVipSslCipherSuitesVersions(d, i["versions"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFirewallVipSslCipherSuitesId(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallVipSslCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipSslCipherSuitesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslClientFallback(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslClientRekeyCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslClientRenegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslClientSessionStateMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslClientSessionStateTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslClientSessionStateType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslDhBits(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslHpkp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslHpkpAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslHpkpBackup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipSslHpkpIncludeSubdomains(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslHpkpPrimary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipSslHpkpReportUri(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslHsts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslHstsAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslHstsIncludeSubdomains(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslHttpLocationConversion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslHttpMatchHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslPfs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslSendEmptyFrags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslServerAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslServerCipherSuites(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["cipher"], _ = expandFirewallVipSslServerCipherSuitesCipher(d, i["cipher"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandFirewallVipSslServerCipherSuitesPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["versions"], _ = expandFirewallVipSslServerCipherSuitesVersions(d, i["versions"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallVipSslServerCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslServerCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslServerCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipSslServerMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslServerMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslServerRenegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslServerSessionStateMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslServerSessionStateTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslServerSessionStateType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipUserAgentDetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipWeblogicServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipWebsphereServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipHttpSupportedMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallVip(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("add_nat46_route"); ok || d.HasChange("add_nat46_route") {
		t, err := expandFirewallVipAddNat46Route(d, v, "add_nat46_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-nat46-route"] = t
		}
	}

	if v, ok := d.GetOk("arp_reply"); ok || d.HasChange("arp_reply") {
		t, err := expandFirewallVipArpReply(d, v, "arp_reply")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arp-reply"] = t
		}
	}

	if v, ok := d.GetOk("client_cert"); ok || d.HasChange("client_cert") {
		t, err := expandFirewallVipClientCert(d, v, "client_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert"] = t
		}
	}

	if v, ok := d.GetOk("color"); ok || d.HasChange("color") {
		t, err := expandFirewallVipColor(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandFirewallVipComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("dns_mapping_ttl"); ok || d.HasChange("dns_mapping_ttl") {
		t, err := expandFirewallVipDnsMappingTtl(d, v, "dns_mapping_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-mapping-ttl"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandFirewallVipDynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("empty_cert_action"); ok || d.HasChange("empty_cert_action") {
		t, err := expandFirewallVipEmptyCertAction(d, v, "empty_cert_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["empty-cert-action"] = t
		}
	}

	if v, ok := d.GetOk("extaddr"); ok || d.HasChange("extaddr") {
		t, err := expandFirewallVipExtaddr(d, v, "extaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extaddr"] = t
		}
	}

	if v, ok := d.GetOk("extintf"); ok || d.HasChange("extintf") {
		t, err := expandFirewallVipExtintf(d, v, "extintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extintf"] = t
		}
	}

	if v, ok := d.GetOk("extip"); ok || d.HasChange("extip") {
		t, err := expandFirewallVipExtip(d, v, "extip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extip"] = t
		}
	}

	if v, ok := d.GetOk("extport"); ok || d.HasChange("extport") {
		t, err := expandFirewallVipExtport(d, v, "extport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extport"] = t
		}
	}

	if v, ok := d.GetOk("gratuitous_arp_interval"); ok || d.HasChange("gratuitous_arp_interval") {
		t, err := expandFirewallVipGratuitousArpInterval(d, v, "gratuitous_arp_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gratuitous-arp-interval"] = t
		}
	}

	if v, ok := d.GetOk("gslb_domain_name"); ok || d.HasChange("gslb_domain_name") {
		t, err := expandFirewallVipGslbDomainName(d, v, "gslb_domain_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gslb-domain-name"] = t
		}
	}

	if v, ok := d.GetOk("gslb_hostname"); ok || d.HasChange("gslb_hostname") {
		t, err := expandFirewallVipGslbHostname(d, v, "gslb_hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gslb-hostname"] = t
		}
	}

	if v, ok := d.GetOk("gslb_public_ips"); ok || d.HasChange("gslb_public_ips") {
		t, err := expandFirewallVipGslbPublicIps(d, v, "gslb_public_ips")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gslb-public-ips"] = t
		}
	}

	if v, ok := d.GetOk("h2_support"); ok || d.HasChange("h2_support") {
		t, err := expandFirewallVipH2Support(d, v, "h2_support")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h2-support"] = t
		}
	}

	if v, ok := d.GetOk("h3_support"); ok || d.HasChange("h3_support") {
		t, err := expandFirewallVipH3Support(d, v, "h3_support")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-support"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_age"); ok || d.HasChange("http_cookie_age") {
		t, err := expandFirewallVipHttpCookieAge(d, v, "http_cookie_age")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-age"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_domain"); ok || d.HasChange("http_cookie_domain") {
		t, err := expandFirewallVipHttpCookieDomain(d, v, "http_cookie_domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-domain"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_domain_from_host"); ok || d.HasChange("http_cookie_domain_from_host") {
		t, err := expandFirewallVipHttpCookieDomainFromHost(d, v, "http_cookie_domain_from_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-domain-from-host"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_generation"); ok || d.HasChange("http_cookie_generation") {
		t, err := expandFirewallVipHttpCookieGeneration(d, v, "http_cookie_generation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-generation"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_path"); ok || d.HasChange("http_cookie_path") {
		t, err := expandFirewallVipHttpCookiePath(d, v, "http_cookie_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-path"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_share"); ok || d.HasChange("http_cookie_share") {
		t, err := expandFirewallVipHttpCookieShare(d, v, "http_cookie_share")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-share"] = t
		}
	}

	if v, ok := d.GetOk("http_ip_header"); ok || d.HasChange("http_ip_header") {
		t, err := expandFirewallVipHttpIpHeader(d, v, "http_ip_header")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-ip-header"] = t
		}
	}

	if v, ok := d.GetOk("http_ip_header_name"); ok || d.HasChange("http_ip_header_name") {
		t, err := expandFirewallVipHttpIpHeaderName(d, v, "http_ip_header_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-ip-header-name"] = t
		}
	}

	if v, ok := d.GetOk("http_multiplex"); ok || d.HasChange("http_multiplex") {
		t, err := expandFirewallVipHttpMultiplex(d, v, "http_multiplex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-multiplex"] = t
		}
	}

	if v, ok := d.GetOk("http_multiplex_max_concurrent_request"); ok || d.HasChange("http_multiplex_max_concurrent_request") {
		t, err := expandFirewallVipHttpMultiplexMaxConcurrentRequest(d, v, "http_multiplex_max_concurrent_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-multiplex-max-concurrent-request"] = t
		}
	}

	if v, ok := d.GetOk("http_multiplex_max_request"); ok || d.HasChange("http_multiplex_max_request") {
		t, err := expandFirewallVipHttpMultiplexMaxRequest(d, v, "http_multiplex_max_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-multiplex-max-request"] = t
		}
	}

	if v, ok := d.GetOk("http_multiplex_ttl"); ok || d.HasChange("http_multiplex_ttl") {
		t, err := expandFirewallVipHttpMultiplexTtl(d, v, "http_multiplex_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-multiplex-ttl"] = t
		}
	}

	if v, ok := d.GetOk("http_redirect"); ok || d.HasChange("http_redirect") {
		t, err := expandFirewallVipHttpRedirect(d, v, "http_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-redirect"] = t
		}
	}

	if v, ok := d.GetOk("https_cookie_secure"); ok || d.HasChange("https_cookie_secure") {
		t, err := expandFirewallVipHttpsCookieSecure(d, v, "https_cookie_secure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-cookie-secure"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandFirewallVipId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_mappedip"); ok || d.HasChange("ipv6_mappedip") {
		t, err := expandFirewallVipIpv6Mappedip(d, v, "ipv6_mappedip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-mappedip"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_mappedport"); ok || d.HasChange("ipv6_mappedport") {
		t, err := expandFirewallVipIpv6Mappedport(d, v, "ipv6_mappedport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-mappedport"] = t
		}
	}

	if v, ok := d.GetOk("ldb_method"); ok || d.HasChange("ldb_method") {
		t, err := expandFirewallVipLdbMethod(d, v, "ldb_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldb-method"] = t
		}
	}

	if v, ok := d.GetOk("mapped_addr"); ok || d.HasChange("mapped_addr") {
		t, err := expandFirewallVipMappedAddr(d, v, "mapped_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mapped-addr"] = t
		}
	}

	if v, ok := d.GetOk("mappedip"); ok || d.HasChange("mappedip") {
		t, err := expandFirewallVipMappedip(d, v, "mappedip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mappedip"] = t
		}
	}

	if v, ok := d.GetOk("mappedport"); ok || d.HasChange("mappedport") {
		t, err := expandFirewallVipMappedport(d, v, "mappedport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mappedport"] = t
		}
	}

	if v, ok := d.GetOk("max_embryonic_connections"); ok || d.HasChange("max_embryonic_connections") {
		t, err := expandFirewallVipMaxEmbryonicConnections(d, v, "max_embryonic_connections")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-embryonic-connections"] = t
		}
	}

	if v, ok := d.GetOk("monitor"); ok || d.HasChange("monitor") {
		t, err := expandFirewallVipMonitor(d, v, "monitor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallVipName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nat_source_vip"); ok || d.HasChange("nat_source_vip") {
		t, err := expandFirewallVipNatSourceVip(d, v, "nat_source_vip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat-source-vip"] = t
		}
	}

	if v, ok := d.GetOk("nat44"); ok || d.HasChange("nat44") {
		t, err := expandFirewallVipNat44(d, v, "nat44")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat44"] = t
		}
	}

	if v, ok := d.GetOk("nat46"); ok || d.HasChange("nat46") {
		t, err := expandFirewallVipNat46(d, v, "nat46")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat46"] = t
		}
	}

	if v, ok := d.GetOk("one_click_gslb_server"); ok || d.HasChange("one_click_gslb_server") {
		t, err := expandFirewallVipOneClickGslbServer(d, v, "one_click_gslb_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["one-click-gslb-server"] = t
		}
	}

	if v, ok := d.GetOk("outlook_web_access"); ok || d.HasChange("outlook_web_access") {
		t, err := expandFirewallVipOutlookWebAccess(d, v, "outlook_web_access")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outlook-web-access"] = t
		}
	}

	if v, ok := d.GetOk("persistence"); ok || d.HasChange("persistence") {
		t, err := expandFirewallVipPersistence(d, v, "persistence")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["persistence"] = t
		}
	}

	if v, ok := d.GetOk("portforward"); ok || d.HasChange("portforward") {
		t, err := expandFirewallVipPortforward(d, v, "portforward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portforward"] = t
		}
	}

	if v, ok := d.GetOk("portmapping_type"); ok || d.HasChange("portmapping_type") {
		t, err := expandFirewallVipPortmappingType(d, v, "portmapping_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portmapping-type"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandFirewallVipProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("quic"); ok || d.HasChange("quic") {
		t, err := expandFirewallVipQuic(d, v, "quic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quic"] = t
		}
	}

	if v, ok := d.GetOk("realservers"); ok || d.HasChange("realservers") {
		t, err := expandFirewallVipRealservers(d, v, "realservers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["realservers"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok || d.HasChange("server_type") {
		t, err := expandFirewallVipServerType(d, v, "server_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandFirewallVipService(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("src_filter"); ok || d.HasChange("src_filter") {
		t, err := expandFirewallVipSrcFilter(d, v, "src_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-filter"] = t
		}
	}

	if v, ok := d.GetOk("src_vip_filter"); ok || d.HasChange("src_vip_filter") {
		t, err := expandFirewallVipSrcVipFilter(d, v, "src_vip_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-vip-filter"] = t
		}
	}

	if v, ok := d.GetOk("srcintf_filter"); ok || d.HasChange("srcintf_filter") {
		t, err := expandFirewallVipSrcintfFilter(d, v, "srcintf_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf-filter"] = t
		}
	}

	if v, ok := d.GetOk("ssl_accept_ffdhe_groups"); ok || d.HasChange("ssl_accept_ffdhe_groups") {
		t, err := expandFirewallVipSslAcceptFfdheGroups(d, v, "ssl_accept_ffdhe_groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-accept-ffdhe-groups"] = t
		}
	}

	if v, ok := d.GetOk("ssl_algorithm"); ok || d.HasChange("ssl_algorithm") {
		t, err := expandFirewallVipSslAlgorithm(d, v, "ssl_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_certificate"); ok || d.HasChange("ssl_certificate") {
		t, err := expandFirewallVipSslCertificate(d, v, "ssl_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-certificate"] = t
		}
	}

	if v, ok := d.GetOk("ssl_cipher_suites"); ok || d.HasChange("ssl_cipher_suites") {
		t, err := expandFirewallVipSslCipherSuites(d, v, "ssl_cipher_suites")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-cipher-suites"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_fallback"); ok || d.HasChange("ssl_client_fallback") {
		t, err := expandFirewallVipSslClientFallback(d, v, "ssl_client_fallback")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-fallback"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_rekey_count"); ok || d.HasChange("ssl_client_rekey_count") {
		t, err := expandFirewallVipSslClientRekeyCount(d, v, "ssl_client_rekey_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-rekey-count"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_renegotiation"); ok || d.HasChange("ssl_client_renegotiation") {
		t, err := expandFirewallVipSslClientRenegotiation(d, v, "ssl_client_renegotiation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-renegotiation"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_session_state_max"); ok || d.HasChange("ssl_client_session_state_max") {
		t, err := expandFirewallVipSslClientSessionStateMax(d, v, "ssl_client_session_state_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-session-state-max"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_session_state_timeout"); ok || d.HasChange("ssl_client_session_state_timeout") {
		t, err := expandFirewallVipSslClientSessionStateTimeout(d, v, "ssl_client_session_state_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-session-state-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_session_state_type"); ok || d.HasChange("ssl_client_session_state_type") {
		t, err := expandFirewallVipSslClientSessionStateType(d, v, "ssl_client_session_state_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-session-state-type"] = t
		}
	}

	if v, ok := d.GetOk("ssl_dh_bits"); ok || d.HasChange("ssl_dh_bits") {
		t, err := expandFirewallVipSslDhBits(d, v, "ssl_dh_bits")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-dh-bits"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp"); ok || d.HasChange("ssl_hpkp") {
		t, err := expandFirewallVipSslHpkp(d, v, "ssl_hpkp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_age"); ok || d.HasChange("ssl_hpkp_age") {
		t, err := expandFirewallVipSslHpkpAge(d, v, "ssl_hpkp_age")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-age"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_backup"); ok || d.HasChange("ssl_hpkp_backup") {
		t, err := expandFirewallVipSslHpkpBackup(d, v, "ssl_hpkp_backup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-backup"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_include_subdomains"); ok || d.HasChange("ssl_hpkp_include_subdomains") {
		t, err := expandFirewallVipSslHpkpIncludeSubdomains(d, v, "ssl_hpkp_include_subdomains")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-include-subdomains"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_primary"); ok || d.HasChange("ssl_hpkp_primary") {
		t, err := expandFirewallVipSslHpkpPrimary(d, v, "ssl_hpkp_primary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-primary"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_report_uri"); ok || d.HasChange("ssl_hpkp_report_uri") {
		t, err := expandFirewallVipSslHpkpReportUri(d, v, "ssl_hpkp_report_uri")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-report-uri"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hsts"); ok || d.HasChange("ssl_hsts") {
		t, err := expandFirewallVipSslHsts(d, v, "ssl_hsts")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hsts"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hsts_age"); ok || d.HasChange("ssl_hsts_age") {
		t, err := expandFirewallVipSslHstsAge(d, v, "ssl_hsts_age")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hsts-age"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hsts_include_subdomains"); ok || d.HasChange("ssl_hsts_include_subdomains") {
		t, err := expandFirewallVipSslHstsIncludeSubdomains(d, v, "ssl_hsts_include_subdomains")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hsts-include-subdomains"] = t
		}
	}

	if v, ok := d.GetOk("ssl_http_location_conversion"); ok || d.HasChange("ssl_http_location_conversion") {
		t, err := expandFirewallVipSslHttpLocationConversion(d, v, "ssl_http_location_conversion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-http-location-conversion"] = t
		}
	}

	if v, ok := d.GetOk("ssl_http_match_host"); ok || d.HasChange("ssl_http_match_host") {
		t, err := expandFirewallVipSslHttpMatchHost(d, v, "ssl_http_match_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-http-match-host"] = t
		}
	}

	if v, ok := d.GetOk("ssl_max_version"); ok || d.HasChange("ssl_max_version") {
		t, err := expandFirewallVipSslMaxVersion(d, v, "ssl_max_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-max-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_version"); ok || d.HasChange("ssl_min_version") {
		t, err := expandFirewallVipSslMinVersion(d, v, "ssl_min_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_mode"); ok || d.HasChange("ssl_mode") {
		t, err := expandFirewallVipSslMode(d, v, "ssl_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-mode"] = t
		}
	}

	if v, ok := d.GetOk("ssl_pfs"); ok || d.HasChange("ssl_pfs") {
		t, err := expandFirewallVipSslPfs(d, v, "ssl_pfs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-pfs"] = t
		}
	}

	if v, ok := d.GetOk("ssl_send_empty_frags"); ok || d.HasChange("ssl_send_empty_frags") {
		t, err := expandFirewallVipSslSendEmptyFrags(d, v, "ssl_send_empty_frags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-send-empty-frags"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_algorithm"); ok || d.HasChange("ssl_server_algorithm") {
		t, err := expandFirewallVipSslServerAlgorithm(d, v, "ssl_server_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_cipher_suites"); ok || d.HasChange("ssl_server_cipher_suites") {
		t, err := expandFirewallVipSslServerCipherSuites(d, v, "ssl_server_cipher_suites")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-cipher-suites"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_max_version"); ok || d.HasChange("ssl_server_max_version") {
		t, err := expandFirewallVipSslServerMaxVersion(d, v, "ssl_server_max_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-max-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_min_version"); ok || d.HasChange("ssl_server_min_version") {
		t, err := expandFirewallVipSslServerMinVersion(d, v, "ssl_server_min_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-min-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_renegotiation"); ok || d.HasChange("ssl_server_renegotiation") {
		t, err := expandFirewallVipSslServerRenegotiation(d, v, "ssl_server_renegotiation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-renegotiation"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_session_state_max"); ok || d.HasChange("ssl_server_session_state_max") {
		t, err := expandFirewallVipSslServerSessionStateMax(d, v, "ssl_server_session_state_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-session-state-max"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_session_state_timeout"); ok || d.HasChange("ssl_server_session_state_timeout") {
		t, err := expandFirewallVipSslServerSessionStateTimeout(d, v, "ssl_server_session_state_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-session-state-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_session_state_type"); ok || d.HasChange("ssl_server_session_state_type") {
		t, err := expandFirewallVipSslServerSessionStateType(d, v, "ssl_server_session_state_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-session-state-type"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFirewallVipStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandFirewallVipType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("user_agent_detect"); ok || d.HasChange("user_agent_detect") {
		t, err := expandFirewallVipUserAgentDetect(d, v, "user_agent_detect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-agent-detect"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandFirewallVipUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("weblogic_server"); ok || d.HasChange("weblogic_server") {
		t, err := expandFirewallVipWeblogicServer(d, v, "weblogic_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weblogic-server"] = t
		}
	}

	if v, ok := d.GetOk("websphere_server"); ok || d.HasChange("websphere_server") {
		t, err := expandFirewallVipWebsphereServer(d, v, "websphere_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["websphere-server"] = t
		}
	}

	if v, ok := d.GetOk("http_supported_max_version"); ok || d.HasChange("http_supported_max_version") {
		t, err := expandFirewallVipHttpSupportedMaxVersion(d, v, "http_supported_max_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-supported-max-version"] = t
		}
	}

	return &obj, nil
}
