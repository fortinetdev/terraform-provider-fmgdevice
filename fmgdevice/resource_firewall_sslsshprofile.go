// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure SSL/SSH protocol options.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallSslSshProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallSslSshProfileCreate,
		Read:   resourceFirewallSslSshProfileRead,
		Update: resourceFirewallSslSshProfileUpdate,
		Delete: resourceFirewallSslSshProfileDelete,

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
			"block_blacklisted_certificates": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"allowlist": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_blocklisted_certificates": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"caname": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dot": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_validation_failure": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cert_validation_timeout": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"client_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"expired_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"min_allowed_ssl_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"proxy_after_tcp_handshake": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quic": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"revoked_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sni_server_cert_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"udp_not_quic": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_cipher": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_negotiation": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"untrusted_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ech_outer_sni": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sni": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ftps": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_validation_failure": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cert_validation_timeout": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"client_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"expired_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"min_allowed_ssl_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ports": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"revoked_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sni_server_cert_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_cipher": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_negotiation": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"untrusted_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"https": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_probe_failure": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cert_validation_failure": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cert_validation_timeout": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"client_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"encrypted_client_hello": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"expired_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"min_allowed_ssl_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ports": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"proxy_after_tcp_handshake": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quic": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"revoked_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sni_server_cert_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"udp_not_quic": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_cipher": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_negotiation": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"untrusted_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"imaps": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_validation_failure": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cert_validation_timeout": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"client_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"expired_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"min_allowed_ssl_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ports": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"proxy_after_tcp_handshake": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"revoked_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sni_server_cert_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_cipher": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_negotiation": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"untrusted_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"mapi_over_https": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pop3s": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_validation_failure": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cert_validation_timeout": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"client_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"expired_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"min_allowed_ssl_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ports": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"proxy_after_tcp_handshake": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"revoked_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sni_server_cert_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_cipher": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_negotiation": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"untrusted_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"rpc_over_https": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"server_cert_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"smtps": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_validation_failure": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cert_validation_timeout": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"client_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"expired_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"min_allowed_ssl_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ports": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"proxy_after_tcp_handshake": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"revoked_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sni_server_cert_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_cipher": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_negotiation": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"untrusted_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ssh": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inspect_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ports": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"proxy_after_tcp_handshake": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssh_algorithm": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssh_tun_policy_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ssl": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_probe_failure": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cert_validation_failure": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cert_validation_timeout": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"client_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"encrypted_client_hello": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"expired_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"inspect_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"min_allowed_ssl_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"revoked_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sni_server_cert_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"unsupported_ssl_cipher": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_negotiation": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"untrusted_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ssl_anomalies_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_anomaly_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_exempt": &schema.Schema{
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
						"address6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"fortiguard_category": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"regex": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"wildcard_fqdn": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"finger_print_category": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ssl_exemption_ip_rating": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_exemption_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_exemptions_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_handshake_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_negotiation_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_server": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ftps_client_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"https_client_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"imaps_client_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pop3s_client_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"smtps_client_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_other_client_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ssl_server_cert_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"supported_alpn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"untrusted_caname": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"use_ssl_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"whitelist": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_client_certificate": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"caname": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"cert": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"keyring_list": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceFirewallSslSshProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallSslSshProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallSslSshProfile resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallSslSshProfile(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallSslSshProfile(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallSslSshProfile resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallSslSshProfile(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallSslSshProfile resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallSslSshProfileRead(d, m)
}

func resourceFirewallSslSshProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallSslSshProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSslSshProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallSslSshProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSslSshProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallSslSshProfileRead(d, m)
}

func resourceFirewallSslSshProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallSslSshProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallSslSshProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallSslSshProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallSslSshProfile(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallSslSshProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallSslSshProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSslSshProfile resource from API: %v", err)
	}
	return nil
}

func flattenFirewallSslSshProfileBlockBlacklistedCertificates(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileAllowlist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileBlockBlocklistedCertificates(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileCaname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSslSshProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDot(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := i["cert-validation-failure"]; ok {
		result["cert_validation_failure"] = flattenFirewallSslSshProfileDotCertValidationFailure(i["cert-validation-failure"], d, pre_append)
	}

	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := i["cert-validation-timeout"]; ok {
		result["cert_validation_timeout"] = flattenFirewallSslSshProfileDotCertValidationTimeout(i["cert-validation-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "client_certificate"
	if _, ok := i["client-certificate"]; ok {
		result["client_certificate"] = flattenFirewallSslSshProfileDotClientCertificate(i["client-certificate"], d, pre_append)
	}

	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := i["expired-server-cert"]; ok {
		result["expired_server_cert"] = flattenFirewallSslSshProfileDotExpiredServerCert(i["expired-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "min_allowed_ssl_version"
	if _, ok := i["min-allowed-ssl-version"]; ok {
		result["min_allowed_ssl_version"] = flattenFirewallSslSshProfileDotMinAllowedSslVersion(i["min-allowed-ssl-version"], d, pre_append)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = flattenFirewallSslSshProfileDotProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append)
	}

	pre_append = pre + ".0." + "quic"
	if _, ok := i["quic"]; ok {
		result["quic"] = flattenFirewallSslSshProfileDotQuic(i["quic"], d, pre_append)
	}

	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := i["revoked-server-cert"]; ok {
		result["revoked_server_cert"] = flattenFirewallSslSshProfileDotRevokedServerCert(i["revoked-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := i["sni-server-cert-check"]; ok {
		result["sni_server_cert_check"] = flattenFirewallSslSshProfileDotSniServerCertCheck(i["sni-server-cert-check"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallSslSshProfileDotStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "udp_not_quic"
	if _, ok := i["udp-not-quic"]; ok {
		result["udp_not_quic"] = flattenFirewallSslSshProfileDotUdpNotQuic(i["udp-not-quic"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := i["unsupported-ssl-cipher"]; ok {
		result["unsupported_ssl_cipher"] = flattenFirewallSslSshProfileDotUnsupportedSslCipher(i["unsupported-ssl-cipher"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := i["unsupported-ssl-negotiation"]; ok {
		result["unsupported_ssl_negotiation"] = flattenFirewallSslSshProfileDotUnsupportedSslNegotiation(i["unsupported-ssl-negotiation"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_version"
	if _, ok := i["unsupported-ssl-version"]; ok {
		result["unsupported_ssl_version"] = flattenFirewallSslSshProfileDotUnsupportedSslVersion(i["unsupported-ssl-version"], d, pre_append)
	}

	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := i["untrusted-server-cert"]; ok {
		result["untrusted_server_cert"] = flattenFirewallSslSshProfileDotUntrustedServerCert(i["untrusted-server-cert"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallSslSshProfileDotCertValidationFailure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotCertValidationTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotClientCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotExpiredServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotMinAllowedSslVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotQuic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotRevokedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotSniServerCertCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotUdpNotQuic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotUnsupportedSslCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotUnsupportedSslNegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotUnsupportedSslVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotUntrustedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileEchOuterSni(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallSslSshProfileEchOuterSniName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "FirewallSslSshProfile-EchOuterSni-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sni"
		if _, ok := i["sni"]; ok {
			v := flattenFirewallSslSshProfileEchOuterSniSni(i["sni"], d, pre_append)
			tmp["sni"] = fortiAPISubPartPatch(v, "FirewallSslSshProfile-EchOuterSni-Sni")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallSslSshProfileEchOuterSniName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileEchOuterSniSni(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtps(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := i["cert-validation-failure"]; ok {
		result["cert_validation_failure"] = flattenFirewallSslSshProfileFtpsCertValidationFailure(i["cert-validation-failure"], d, pre_append)
	}

	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := i["cert-validation-timeout"]; ok {
		result["cert_validation_timeout"] = flattenFirewallSslSshProfileFtpsCertValidationTimeout(i["cert-validation-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "client_certificate"
	if _, ok := i["client-certificate"]; ok {
		result["client_certificate"] = flattenFirewallSslSshProfileFtpsClientCertificate(i["client-certificate"], d, pre_append)
	}

	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := i["expired-server-cert"]; ok {
		result["expired_server_cert"] = flattenFirewallSslSshProfileFtpsExpiredServerCert(i["expired-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "min_allowed_ssl_version"
	if _, ok := i["min-allowed-ssl-version"]; ok {
		result["min_allowed_ssl_version"] = flattenFirewallSslSshProfileFtpsMinAllowedSslVersion(i["min-allowed-ssl-version"], d, pre_append)
	}

	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallSslSshProfileFtpsPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := i["revoked-server-cert"]; ok {
		result["revoked_server_cert"] = flattenFirewallSslSshProfileFtpsRevokedServerCert(i["revoked-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := i["sni-server-cert-check"]; ok {
		result["sni_server_cert_check"] = flattenFirewallSslSshProfileFtpsSniServerCertCheck(i["sni-server-cert-check"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallSslSshProfileFtpsStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := i["unsupported-ssl-cipher"]; ok {
		result["unsupported_ssl_cipher"] = flattenFirewallSslSshProfileFtpsUnsupportedSslCipher(i["unsupported-ssl-cipher"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := i["unsupported-ssl-negotiation"]; ok {
		result["unsupported_ssl_negotiation"] = flattenFirewallSslSshProfileFtpsUnsupportedSslNegotiation(i["unsupported-ssl-negotiation"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_version"
	if _, ok := i["unsupported-ssl-version"]; ok {
		result["unsupported_ssl_version"] = flattenFirewallSslSshProfileFtpsUnsupportedSslVersion(i["unsupported-ssl-version"], d, pre_append)
	}

	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := i["untrusted-server-cert"]; ok {
		result["untrusted_server_cert"] = flattenFirewallSslSshProfileFtpsUntrustedServerCert(i["untrusted-server-cert"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallSslSshProfileFtpsCertValidationFailure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtpsCertValidationTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtpsClientCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtpsExpiredServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtpsMinAllowedSslVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtpsPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenFirewallSslSshProfileFtpsRevokedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtpsSniServerCertCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtpsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtpsUnsupportedSslCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtpsUnsupportedSslNegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtpsUnsupportedSslVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtpsUntrustedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttps(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cert_probe_failure"
	if _, ok := i["cert-probe-failure"]; ok {
		result["cert_probe_failure"] = flattenFirewallSslSshProfileHttpsCertProbeFailure(i["cert-probe-failure"], d, pre_append)
	}

	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := i["cert-validation-failure"]; ok {
		result["cert_validation_failure"] = flattenFirewallSslSshProfileHttpsCertValidationFailure(i["cert-validation-failure"], d, pre_append)
	}

	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := i["cert-validation-timeout"]; ok {
		result["cert_validation_timeout"] = flattenFirewallSslSshProfileHttpsCertValidationTimeout(i["cert-validation-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "client_certificate"
	if _, ok := i["client-certificate"]; ok {
		result["client_certificate"] = flattenFirewallSslSshProfileHttpsClientCertificate(i["client-certificate"], d, pre_append)
	}

	pre_append = pre + ".0." + "encrypted_client_hello"
	if _, ok := i["encrypted-client-hello"]; ok {
		result["encrypted_client_hello"] = flattenFirewallSslSshProfileHttpsEncryptedClientHello(i["encrypted-client-hello"], d, pre_append)
	}

	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := i["expired-server-cert"]; ok {
		result["expired_server_cert"] = flattenFirewallSslSshProfileHttpsExpiredServerCert(i["expired-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "min_allowed_ssl_version"
	if _, ok := i["min-allowed-ssl-version"]; ok {
		result["min_allowed_ssl_version"] = flattenFirewallSslSshProfileHttpsMinAllowedSslVersion(i["min-allowed-ssl-version"], d, pre_append)
	}

	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallSslSshProfileHttpsPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = flattenFirewallSslSshProfileHttpsProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append)
	}

	pre_append = pre + ".0." + "quic"
	if _, ok := i["quic"]; ok {
		result["quic"] = flattenFirewallSslSshProfileHttpsQuic(i["quic"], d, pre_append)
	}

	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := i["revoked-server-cert"]; ok {
		result["revoked_server_cert"] = flattenFirewallSslSshProfileHttpsRevokedServerCert(i["revoked-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := i["sni-server-cert-check"]; ok {
		result["sni_server_cert_check"] = flattenFirewallSslSshProfileHttpsSniServerCertCheck(i["sni-server-cert-check"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallSslSshProfileHttpsStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "udp_not_quic"
	if _, ok := i["udp-not-quic"]; ok {
		result["udp_not_quic"] = flattenFirewallSslSshProfileHttpsUdpNotQuic(i["udp-not-quic"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := i["unsupported-ssl-cipher"]; ok {
		result["unsupported_ssl_cipher"] = flattenFirewallSslSshProfileHttpsUnsupportedSslCipher(i["unsupported-ssl-cipher"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := i["unsupported-ssl-negotiation"]; ok {
		result["unsupported_ssl_negotiation"] = flattenFirewallSslSshProfileHttpsUnsupportedSslNegotiation(i["unsupported-ssl-negotiation"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_version"
	if _, ok := i["unsupported-ssl-version"]; ok {
		result["unsupported_ssl_version"] = flattenFirewallSslSshProfileHttpsUnsupportedSslVersion(i["unsupported-ssl-version"], d, pre_append)
	}

	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := i["untrusted-server-cert"]; ok {
		result["untrusted_server_cert"] = flattenFirewallSslSshProfileHttpsUntrustedServerCert(i["untrusted-server-cert"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallSslSshProfileHttpsCertProbeFailure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsCertValidationFailure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsCertValidationTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsClientCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsEncryptedClientHello(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsExpiredServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsMinAllowedSslVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenFirewallSslSshProfileHttpsProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsQuic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsRevokedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsSniServerCertCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsUdpNotQuic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsUnsupportedSslCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsUnsupportedSslNegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsUnsupportedSslVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsUntrustedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImaps(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := i["cert-validation-failure"]; ok {
		result["cert_validation_failure"] = flattenFirewallSslSshProfileImapsCertValidationFailure(i["cert-validation-failure"], d, pre_append)
	}

	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := i["cert-validation-timeout"]; ok {
		result["cert_validation_timeout"] = flattenFirewallSslSshProfileImapsCertValidationTimeout(i["cert-validation-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "client_certificate"
	if _, ok := i["client-certificate"]; ok {
		result["client_certificate"] = flattenFirewallSslSshProfileImapsClientCertificate(i["client-certificate"], d, pre_append)
	}

	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := i["expired-server-cert"]; ok {
		result["expired_server_cert"] = flattenFirewallSslSshProfileImapsExpiredServerCert(i["expired-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "min_allowed_ssl_version"
	if _, ok := i["min-allowed-ssl-version"]; ok {
		result["min_allowed_ssl_version"] = flattenFirewallSslSshProfileImapsMinAllowedSslVersion(i["min-allowed-ssl-version"], d, pre_append)
	}

	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallSslSshProfileImapsPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = flattenFirewallSslSshProfileImapsProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append)
	}

	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := i["revoked-server-cert"]; ok {
		result["revoked_server_cert"] = flattenFirewallSslSshProfileImapsRevokedServerCert(i["revoked-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := i["sni-server-cert-check"]; ok {
		result["sni_server_cert_check"] = flattenFirewallSslSshProfileImapsSniServerCertCheck(i["sni-server-cert-check"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallSslSshProfileImapsStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := i["unsupported-ssl-cipher"]; ok {
		result["unsupported_ssl_cipher"] = flattenFirewallSslSshProfileImapsUnsupportedSslCipher(i["unsupported-ssl-cipher"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := i["unsupported-ssl-negotiation"]; ok {
		result["unsupported_ssl_negotiation"] = flattenFirewallSslSshProfileImapsUnsupportedSslNegotiation(i["unsupported-ssl-negotiation"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_version"
	if _, ok := i["unsupported-ssl-version"]; ok {
		result["unsupported_ssl_version"] = flattenFirewallSslSshProfileImapsUnsupportedSslVersion(i["unsupported-ssl-version"], d, pre_append)
	}

	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := i["untrusted-server-cert"]; ok {
		result["untrusted_server_cert"] = flattenFirewallSslSshProfileImapsUntrustedServerCert(i["untrusted-server-cert"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallSslSshProfileImapsCertValidationFailure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsCertValidationTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsClientCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsExpiredServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsMinAllowedSslVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenFirewallSslSshProfileImapsProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsRevokedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsSniServerCertCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsUnsupportedSslCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsUnsupportedSslNegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsUnsupportedSslVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsUntrustedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileMapiOverHttps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3S(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := i["cert-validation-failure"]; ok {
		result["cert_validation_failure"] = flattenFirewallSslSshProfilePop3SCertValidationFailure(i["cert-validation-failure"], d, pre_append)
	}

	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := i["cert-validation-timeout"]; ok {
		result["cert_validation_timeout"] = flattenFirewallSslSshProfilePop3SCertValidationTimeout(i["cert-validation-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "client_certificate"
	if _, ok := i["client-certificate"]; ok {
		result["client_certificate"] = flattenFirewallSslSshProfilePop3SClientCertificate(i["client-certificate"], d, pre_append)
	}

	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := i["expired-server-cert"]; ok {
		result["expired_server_cert"] = flattenFirewallSslSshProfilePop3SExpiredServerCert(i["expired-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "min_allowed_ssl_version"
	if _, ok := i["min-allowed-ssl-version"]; ok {
		result["min_allowed_ssl_version"] = flattenFirewallSslSshProfilePop3SMinAllowedSslVersion(i["min-allowed-ssl-version"], d, pre_append)
	}

	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallSslSshProfilePop3SPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = flattenFirewallSslSshProfilePop3SProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append)
	}

	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := i["revoked-server-cert"]; ok {
		result["revoked_server_cert"] = flattenFirewallSslSshProfilePop3SRevokedServerCert(i["revoked-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := i["sni-server-cert-check"]; ok {
		result["sni_server_cert_check"] = flattenFirewallSslSshProfilePop3SSniServerCertCheck(i["sni-server-cert-check"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallSslSshProfilePop3SStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := i["unsupported-ssl-cipher"]; ok {
		result["unsupported_ssl_cipher"] = flattenFirewallSslSshProfilePop3SUnsupportedSslCipher(i["unsupported-ssl-cipher"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := i["unsupported-ssl-negotiation"]; ok {
		result["unsupported_ssl_negotiation"] = flattenFirewallSslSshProfilePop3SUnsupportedSslNegotiation(i["unsupported-ssl-negotiation"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_version"
	if _, ok := i["unsupported-ssl-version"]; ok {
		result["unsupported_ssl_version"] = flattenFirewallSslSshProfilePop3SUnsupportedSslVersion(i["unsupported-ssl-version"], d, pre_append)
	}

	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := i["untrusted-server-cert"]; ok {
		result["untrusted_server_cert"] = flattenFirewallSslSshProfilePop3SUntrustedServerCert(i["untrusted-server-cert"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallSslSshProfilePop3SCertValidationFailure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SCertValidationTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SClientCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SExpiredServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SMinAllowedSslVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenFirewallSslSshProfilePop3SProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SRevokedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SSniServerCertCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SUnsupportedSslCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SUnsupportedSslNegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SUnsupportedSslVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SUntrustedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileRpcOverHttps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSslSshProfileServerCertMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtps(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := i["cert-validation-failure"]; ok {
		result["cert_validation_failure"] = flattenFirewallSslSshProfileSmtpsCertValidationFailure(i["cert-validation-failure"], d, pre_append)
	}

	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := i["cert-validation-timeout"]; ok {
		result["cert_validation_timeout"] = flattenFirewallSslSshProfileSmtpsCertValidationTimeout(i["cert-validation-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "client_certificate"
	if _, ok := i["client-certificate"]; ok {
		result["client_certificate"] = flattenFirewallSslSshProfileSmtpsClientCertificate(i["client-certificate"], d, pre_append)
	}

	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := i["expired-server-cert"]; ok {
		result["expired_server_cert"] = flattenFirewallSslSshProfileSmtpsExpiredServerCert(i["expired-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "min_allowed_ssl_version"
	if _, ok := i["min-allowed-ssl-version"]; ok {
		result["min_allowed_ssl_version"] = flattenFirewallSslSshProfileSmtpsMinAllowedSslVersion(i["min-allowed-ssl-version"], d, pre_append)
	}

	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallSslSshProfileSmtpsPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = flattenFirewallSslSshProfileSmtpsProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append)
	}

	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := i["revoked-server-cert"]; ok {
		result["revoked_server_cert"] = flattenFirewallSslSshProfileSmtpsRevokedServerCert(i["revoked-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := i["sni-server-cert-check"]; ok {
		result["sni_server_cert_check"] = flattenFirewallSslSshProfileSmtpsSniServerCertCheck(i["sni-server-cert-check"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallSslSshProfileSmtpsStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := i["unsupported-ssl-cipher"]; ok {
		result["unsupported_ssl_cipher"] = flattenFirewallSslSshProfileSmtpsUnsupportedSslCipher(i["unsupported-ssl-cipher"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := i["unsupported-ssl-negotiation"]; ok {
		result["unsupported_ssl_negotiation"] = flattenFirewallSslSshProfileSmtpsUnsupportedSslNegotiation(i["unsupported-ssl-negotiation"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_version"
	if _, ok := i["unsupported-ssl-version"]; ok {
		result["unsupported_ssl_version"] = flattenFirewallSslSshProfileSmtpsUnsupportedSslVersion(i["unsupported-ssl-version"], d, pre_append)
	}

	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := i["untrusted-server-cert"]; ok {
		result["untrusted_server_cert"] = flattenFirewallSslSshProfileSmtpsUntrustedServerCert(i["untrusted-server-cert"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallSslSshProfileSmtpsCertValidationFailure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsCertValidationTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsClientCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsExpiredServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsMinAllowedSslVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenFirewallSslSshProfileSmtpsProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsRevokedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsSniServerCertCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsUnsupportedSslCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsUnsupportedSslNegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsUnsupportedSslVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsUntrustedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSsh(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = flattenFirewallSslSshProfileSshInspectAll(i["inspect-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallSslSshProfileSshPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = flattenFirewallSslSshProfileSshProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssh_algorithm"
	if _, ok := i["ssh-algorithm"]; ok {
		result["ssh_algorithm"] = flattenFirewallSslSshProfileSshSshAlgorithm(i["ssh-algorithm"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssh_tun_policy_check"
	if _, ok := i["ssh-tun-policy-check"]; ok {
		result["ssh_tun_policy_check"] = flattenFirewallSslSshProfileSshSshTunPolicyCheck(i["ssh-tun-policy-check"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallSslSshProfileSshStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_version"
	if _, ok := i["unsupported-version"]; ok {
		result["unsupported_version"] = flattenFirewallSslSshProfileSshUnsupportedVersion(i["unsupported-version"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallSslSshProfileSshInspectAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSshPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenFirewallSslSshProfileSshProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSshSshAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSshSshTunPolicyCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSshStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSshUnsupportedVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSsl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cert_probe_failure"
	if _, ok := i["cert-probe-failure"]; ok {
		result["cert_probe_failure"] = flattenFirewallSslSshProfileSslCertProbeFailure(i["cert-probe-failure"], d, pre_append)
	}

	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := i["cert-validation-failure"]; ok {
		result["cert_validation_failure"] = flattenFirewallSslSshProfileSslCertValidationFailure(i["cert-validation-failure"], d, pre_append)
	}

	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := i["cert-validation-timeout"]; ok {
		result["cert_validation_timeout"] = flattenFirewallSslSshProfileSslCertValidationTimeout(i["cert-validation-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "client_certificate"
	if _, ok := i["client-certificate"]; ok {
		result["client_certificate"] = flattenFirewallSslSshProfileSslClientCertificateU(i["client-certificate"], d, pre_append)
	}

	pre_append = pre + ".0." + "encrypted_client_hello"
	if _, ok := i["encrypted-client-hello"]; ok {
		result["encrypted_client_hello"] = flattenFirewallSslSshProfileSslEncryptedClientHello(i["encrypted-client-hello"], d, pre_append)
	}

	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := i["expired-server-cert"]; ok {
		result["expired_server_cert"] = flattenFirewallSslSshProfileSslExpiredServerCert(i["expired-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = flattenFirewallSslSshProfileSslInspectAll(i["inspect-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "min_allowed_ssl_version"
	if _, ok := i["min-allowed-ssl-version"]; ok {
		result["min_allowed_ssl_version"] = flattenFirewallSslSshProfileSslMinAllowedSslVersion(i["min-allowed-ssl-version"], d, pre_append)
	}

	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := i["revoked-server-cert"]; ok {
		result["revoked_server_cert"] = flattenFirewallSslSshProfileSslRevokedServerCert(i["revoked-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := i["sni-server-cert-check"]; ok {
		result["sni_server_cert_check"] = flattenFirewallSslSshProfileSslSniServerCertCheck(i["sni-server-cert-check"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := i["unsupported-ssl-cipher"]; ok {
		result["unsupported_ssl_cipher"] = flattenFirewallSslSshProfileSslUnsupportedSslCipher(i["unsupported-ssl-cipher"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := i["unsupported-ssl-negotiation"]; ok {
		result["unsupported_ssl_negotiation"] = flattenFirewallSslSshProfileSslUnsupportedSslNegotiation(i["unsupported-ssl-negotiation"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_version"
	if _, ok := i["unsupported-ssl-version"]; ok {
		result["unsupported_ssl_version"] = flattenFirewallSslSshProfileSslUnsupportedSslVersion(i["unsupported-ssl-version"], d, pre_append)
	}

	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := i["untrusted-server-cert"]; ok {
		result["untrusted_server_cert"] = flattenFirewallSslSshProfileSslUntrustedServerCert(i["untrusted-server-cert"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallSslSshProfileSslCertProbeFailure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslCertValidationFailure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslCertValidationTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslClientCertificateU(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslEncryptedClientHello(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExpiredServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslInspectAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslMinAllowedSslVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslRevokedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslSniServerCertCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslUnsupportedSslCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslUnsupportedSslNegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslUnsupportedSslVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslUntrustedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslAnomaliesLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslAnomalyLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExempt(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallSslSshProfileSslExemptAddress(i["address"], d, pre_append)
			tmp["address"] = fortiAPISubPartPatch(v, "FirewallSslSshProfile-SslExempt-Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address6"
		if _, ok := i["address6"]; ok {
			v := flattenFirewallSslSshProfileSslExemptAddress6(i["address6"], d, pre_append)
			tmp["address6"] = fortiAPISubPartPatch(v, "FirewallSslSshProfile-SslExempt-Address6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortiguard_category"
		if _, ok := i["fortiguard-category"]; ok {
			v := flattenFirewallSslSshProfileSslExemptFortiguardCategory(i["fortiguard-category"], d, pre_append)
			tmp["fortiguard_category"] = fortiAPISubPartPatch(v, "FirewallSslSshProfile-SslExempt-FortiguardCategory")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFirewallSslSshProfileSslExemptId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FirewallSslSshProfile-SslExempt-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := i["regex"]; ok {
			v := flattenFirewallSslSshProfileSslExemptRegex(i["regex"], d, pre_append)
			tmp["regex"] = fortiAPISubPartPatch(v, "FirewallSslSshProfile-SslExempt-Regex")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenFirewallSslSshProfileSslExemptType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "FirewallSslSshProfile-SslExempt-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wildcard_fqdn"
		if _, ok := i["wildcard-fqdn"]; ok {
			v := flattenFirewallSslSshProfileSslExemptWildcardFqdn(i["wildcard-fqdn"], d, pre_append)
			tmp["wildcard_fqdn"] = fortiAPISubPartPatch(v, "FirewallSslSshProfile-SslExempt-WildcardFqdn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "finger_print_category"
		if _, ok := i["finger-print-category"]; ok {
			v := flattenFirewallSslSshProfileSslExemptFingerPrintCategory(i["finger-print-category"], d, pre_append)
			tmp["finger_print_category"] = fortiAPISubPartPatch(v, "FirewallSslSshProfile-SslExempt-FingerPrintCategory")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallSslSshProfileSslExemptAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSslSshProfileSslExemptAddress6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSslSshProfileSslExemptFortiguardCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSslSshProfileSslExemptId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExemptRegex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExemptType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExemptWildcardFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSslSshProfileSslExemptFingerPrintCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExemptionIpRating(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExemptionLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExemptionsLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslHandshakeLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslNegotiationLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServer(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftps_client_certificate"
		if _, ok := i["ftps-client-certificate"]; ok {
			v := flattenFirewallSslSshProfileSslServerFtpsClientCertificate(i["ftps-client-certificate"], d, pre_append)
			tmp["ftps_client_certificate"] = fortiAPISubPartPatch(v, "FirewallSslSshProfile-SslServer-FtpsClientCertificate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_client_certificate"
		if _, ok := i["https-client-certificate"]; ok {
			v := flattenFirewallSslSshProfileSslServerHttpsClientCertificate(i["https-client-certificate"], d, pre_append)
			tmp["https_client_certificate"] = fortiAPISubPartPatch(v, "FirewallSslSshProfile-SslServer-HttpsClientCertificate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFirewallSslSshProfileSslServerId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FirewallSslSshProfile-SslServer-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "imaps_client_certificate"
		if _, ok := i["imaps-client-certificate"]; ok {
			v := flattenFirewallSslSshProfileSslServerImapsClientCertificate(i["imaps-client-certificate"], d, pre_append)
			tmp["imaps_client_certificate"] = fortiAPISubPartPatch(v, "FirewallSslSshProfile-SslServer-ImapsClientCertificate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenFirewallSslSshProfileSslServerIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "FirewallSslSshProfile-SslServer-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pop3s_client_certificate"
		if _, ok := i["pop3s-client-certificate"]; ok {
			v := flattenFirewallSslSshProfileSslServerPop3SClientCertificate(i["pop3s-client-certificate"], d, pre_append)
			tmp["pop3s_client_certificate"] = fortiAPISubPartPatch(v, "FirewallSslSshProfile-SslServer-Pop3SClientCertificate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "smtps_client_certificate"
		if _, ok := i["smtps-client-certificate"]; ok {
			v := flattenFirewallSslSshProfileSslServerSmtpsClientCertificate(i["smtps-client-certificate"], d, pre_append)
			tmp["smtps_client_certificate"] = fortiAPISubPartPatch(v, "FirewallSslSshProfile-SslServer-SmtpsClientCertificate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_other_client_certificate"
		if _, ok := i["ssl-other-client-certificate"]; ok {
			v := flattenFirewallSslSshProfileSslServerSslOtherClientCertificate(i["ssl-other-client-certificate"], d, pre_append)
			tmp["ssl_other_client_certificate"] = fortiAPISubPartPatch(v, "FirewallSslSshProfile-SslServer-SslOtherClientCertificate")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallSslSshProfileSslServerFtpsClientCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerHttpsClientCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerImapsClientCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerPop3SClientCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerSmtpsClientCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerSslOtherClientCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerCertLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSupportedAlpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileUntrustedCaname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSslSshProfileUseSslServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileWhitelist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslClientCertificate(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "caname"
	if _, ok := i["caname"]; ok {
		result["caname"] = flattenFirewallSslSshProfileSslClientCertificateCaname(i["caname"], d, pre_append)
	}

	pre_append = pre + ".0." + "cert"
	if _, ok := i["cert"]; ok {
		result["cert"] = flattenFirewallSslSshProfileSslClientCertificateCert(i["cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "keyring_list"
	if _, ok := i["keyring-list"]; ok {
		result["keyring_list"] = flattenFirewallSslSshProfileSslClientCertificateKeyringList(i["keyring-list"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallSslSshProfileSslClientCertificateStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallSslSshProfileSslClientCertificateCaname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSslSshProfileSslClientCertificateCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSslSshProfileSslClientCertificateKeyringList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSslSshProfileSslClientCertificateStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallSslSshProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("block_blacklisted_certificates", flattenFirewallSslSshProfileBlockBlacklistedCertificates(o["block-blacklisted-certificates"], d, "block_blacklisted_certificates")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-blacklisted-certificates"], "FirewallSslSshProfile-BlockBlacklistedCertificates"); ok {
			if err = d.Set("block_blacklisted_certificates", vv); err != nil {
				return fmt.Errorf("Error reading block_blacklisted_certificates: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_blacklisted_certificates: %v", err)
		}
	}

	if err = d.Set("allowlist", flattenFirewallSslSshProfileAllowlist(o["allowlist"], d, "allowlist")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowlist"], "FirewallSslSshProfile-Allowlist"); ok {
			if err = d.Set("allowlist", vv); err != nil {
				return fmt.Errorf("Error reading allowlist: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowlist: %v", err)
		}
	}

	if err = d.Set("block_blocklisted_certificates", flattenFirewallSslSshProfileBlockBlocklistedCertificates(o["block-blocklisted-certificates"], d, "block_blocklisted_certificates")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-blocklisted-certificates"], "FirewallSslSshProfile-BlockBlocklistedCertificates"); ok {
			if err = d.Set("block_blocklisted_certificates", vv); err != nil {
				return fmt.Errorf("Error reading block_blocklisted_certificates: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_blocklisted_certificates: %v", err)
		}
	}

	if err = d.Set("caname", flattenFirewallSslSshProfileCaname(o["caname"], d, "caname")); err != nil {
		if vv, ok := fortiAPIPatch(o["caname"], "FirewallSslSshProfile-Caname"); ok {
			if err = d.Set("caname", vv); err != nil {
				return fmt.Errorf("Error reading caname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading caname: %v", err)
		}
	}

	if err = d.Set("comment", flattenFirewallSslSshProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "FirewallSslSshProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dot", flattenFirewallSslSshProfileDot(o["dot"], d, "dot")); err != nil {
			if vv, ok := fortiAPIPatch(o["dot"], "FirewallSslSshProfile-Dot"); ok {
				if err = d.Set("dot", vv); err != nil {
					return fmt.Errorf("Error reading dot: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dot: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dot"); ok {
			if err = d.Set("dot", flattenFirewallSslSshProfileDot(o["dot"], d, "dot")); err != nil {
				if vv, ok := fortiAPIPatch(o["dot"], "FirewallSslSshProfile-Dot"); ok {
					if err = d.Set("dot", vv); err != nil {
						return fmt.Errorf("Error reading dot: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dot: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ech_outer_sni", flattenFirewallSslSshProfileEchOuterSni(o["ech-outer-sni"], d, "ech_outer_sni")); err != nil {
			if vv, ok := fortiAPIPatch(o["ech-outer-sni"], "FirewallSslSshProfile-EchOuterSni"); ok {
				if err = d.Set("ech_outer_sni", vv); err != nil {
					return fmt.Errorf("Error reading ech_outer_sni: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ech_outer_sni: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ech_outer_sni"); ok {
			if err = d.Set("ech_outer_sni", flattenFirewallSslSshProfileEchOuterSni(o["ech-outer-sni"], d, "ech_outer_sni")); err != nil {
				if vv, ok := fortiAPIPatch(o["ech-outer-sni"], "FirewallSslSshProfile-EchOuterSni"); ok {
					if err = d.Set("ech_outer_sni", vv); err != nil {
						return fmt.Errorf("Error reading ech_outer_sni: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ech_outer_sni: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ftps", flattenFirewallSslSshProfileFtps(o["ftps"], d, "ftps")); err != nil {
			if vv, ok := fortiAPIPatch(o["ftps"], "FirewallSslSshProfile-Ftps"); ok {
				if err = d.Set("ftps", vv); err != nil {
					return fmt.Errorf("Error reading ftps: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ftps: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ftps"); ok {
			if err = d.Set("ftps", flattenFirewallSslSshProfileFtps(o["ftps"], d, "ftps")); err != nil {
				if vv, ok := fortiAPIPatch(o["ftps"], "FirewallSslSshProfile-Ftps"); ok {
					if err = d.Set("ftps", vv); err != nil {
						return fmt.Errorf("Error reading ftps: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ftps: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("https", flattenFirewallSslSshProfileHttps(o["https"], d, "https")); err != nil {
			if vv, ok := fortiAPIPatch(o["https"], "FirewallSslSshProfile-Https"); ok {
				if err = d.Set("https", vv); err != nil {
					return fmt.Errorf("Error reading https: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading https: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("https"); ok {
			if err = d.Set("https", flattenFirewallSslSshProfileHttps(o["https"], d, "https")); err != nil {
				if vv, ok := fortiAPIPatch(o["https"], "FirewallSslSshProfile-Https"); ok {
					if err = d.Set("https", vv); err != nil {
						return fmt.Errorf("Error reading https: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading https: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("imaps", flattenFirewallSslSshProfileImaps(o["imaps"], d, "imaps")); err != nil {
			if vv, ok := fortiAPIPatch(o["imaps"], "FirewallSslSshProfile-Imaps"); ok {
				if err = d.Set("imaps", vv); err != nil {
					return fmt.Errorf("Error reading imaps: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading imaps: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("imaps"); ok {
			if err = d.Set("imaps", flattenFirewallSslSshProfileImaps(o["imaps"], d, "imaps")); err != nil {
				if vv, ok := fortiAPIPatch(o["imaps"], "FirewallSslSshProfile-Imaps"); ok {
					if err = d.Set("imaps", vv); err != nil {
						return fmt.Errorf("Error reading imaps: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading imaps: %v", err)
				}
			}
		}
	}

	if err = d.Set("mapi_over_https", flattenFirewallSslSshProfileMapiOverHttps(o["mapi-over-https"], d, "mapi_over_https")); err != nil {
		if vv, ok := fortiAPIPatch(o["mapi-over-https"], "FirewallSslSshProfile-MapiOverHttps"); ok {
			if err = d.Set("mapi_over_https", vv); err != nil {
				return fmt.Errorf("Error reading mapi_over_https: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mapi_over_https: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallSslSshProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallSslSshProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("pop3s", flattenFirewallSslSshProfilePop3S(o["pop3s"], d, "pop3s")); err != nil {
			if vv, ok := fortiAPIPatch(o["pop3s"], "FirewallSslSshProfile-Pop3S"); ok {
				if err = d.Set("pop3s", vv); err != nil {
					return fmt.Errorf("Error reading pop3s: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading pop3s: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("pop3s"); ok {
			if err = d.Set("pop3s", flattenFirewallSslSshProfilePop3S(o["pop3s"], d, "pop3s")); err != nil {
				if vv, ok := fortiAPIPatch(o["pop3s"], "FirewallSslSshProfile-Pop3S"); ok {
					if err = d.Set("pop3s", vv); err != nil {
						return fmt.Errorf("Error reading pop3s: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading pop3s: %v", err)
				}
			}
		}
	}

	if err = d.Set("rpc_over_https", flattenFirewallSslSshProfileRpcOverHttps(o["rpc-over-https"], d, "rpc_over_https")); err != nil {
		if vv, ok := fortiAPIPatch(o["rpc-over-https"], "FirewallSslSshProfile-RpcOverHttps"); ok {
			if err = d.Set("rpc_over_https", vv); err != nil {
				return fmt.Errorf("Error reading rpc_over_https: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rpc_over_https: %v", err)
		}
	}

	if err = d.Set("server_cert", flattenFirewallSslSshProfileServerCert(o["server-cert"], d, "server_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-cert"], "FirewallSslSshProfile-ServerCert"); ok {
			if err = d.Set("server_cert", vv); err != nil {
				return fmt.Errorf("Error reading server_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_cert: %v", err)
		}
	}

	if err = d.Set("server_cert_mode", flattenFirewallSslSshProfileServerCertMode(o["server-cert-mode"], d, "server_cert_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-cert-mode"], "FirewallSslSshProfile-ServerCertMode"); ok {
			if err = d.Set("server_cert_mode", vv); err != nil {
				return fmt.Errorf("Error reading server_cert_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_cert_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("smtps", flattenFirewallSslSshProfileSmtps(o["smtps"], d, "smtps")); err != nil {
			if vv, ok := fortiAPIPatch(o["smtps"], "FirewallSslSshProfile-Smtps"); ok {
				if err = d.Set("smtps", vv); err != nil {
					return fmt.Errorf("Error reading smtps: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading smtps: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("smtps"); ok {
			if err = d.Set("smtps", flattenFirewallSslSshProfileSmtps(o["smtps"], d, "smtps")); err != nil {
				if vv, ok := fortiAPIPatch(o["smtps"], "FirewallSslSshProfile-Smtps"); ok {
					if err = d.Set("smtps", vv); err != nil {
						return fmt.Errorf("Error reading smtps: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading smtps: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ssh", flattenFirewallSslSshProfileSsh(o["ssh"], d, "ssh")); err != nil {
			if vv, ok := fortiAPIPatch(o["ssh"], "FirewallSslSshProfile-Ssh"); ok {
				if err = d.Set("ssh", vv); err != nil {
					return fmt.Errorf("Error reading ssh: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ssh: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssh"); ok {
			if err = d.Set("ssh", flattenFirewallSslSshProfileSsh(o["ssh"], d, "ssh")); err != nil {
				if vv, ok := fortiAPIPatch(o["ssh"], "FirewallSslSshProfile-Ssh"); ok {
					if err = d.Set("ssh", vv); err != nil {
						return fmt.Errorf("Error reading ssh: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ssh: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ssl", flattenFirewallSslSshProfileSsl(o["ssl"], d, "ssl")); err != nil {
			if vv, ok := fortiAPIPatch(o["ssl"], "FirewallSslSshProfile-Ssl"); ok {
				if err = d.Set("ssl", vv); err != nil {
					return fmt.Errorf("Error reading ssl: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ssl: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl"); ok {
			if err = d.Set("ssl", flattenFirewallSslSshProfileSsl(o["ssl"], d, "ssl")); err != nil {
				if vv, ok := fortiAPIPatch(o["ssl"], "FirewallSslSshProfile-Ssl"); ok {
					if err = d.Set("ssl", vv); err != nil {
						return fmt.Errorf("Error reading ssl: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ssl: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_anomalies_log", flattenFirewallSslSshProfileSslAnomaliesLog(o["ssl-anomalies-log"], d, "ssl_anomalies_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-anomalies-log"], "FirewallSslSshProfile-SslAnomaliesLog"); ok {
			if err = d.Set("ssl_anomalies_log", vv); err != nil {
				return fmt.Errorf("Error reading ssl_anomalies_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_anomalies_log: %v", err)
		}
	}

	if err = d.Set("ssl_anomaly_log", flattenFirewallSslSshProfileSslAnomalyLog(o["ssl-anomaly-log"], d, "ssl_anomaly_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-anomaly-log"], "FirewallSslSshProfile-SslAnomalyLog"); ok {
			if err = d.Set("ssl_anomaly_log", vv); err != nil {
				return fmt.Errorf("Error reading ssl_anomaly_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_anomaly_log: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ssl_exempt", flattenFirewallSslSshProfileSslExempt(o["ssl-exempt"], d, "ssl_exempt")); err != nil {
			if vv, ok := fortiAPIPatch(o["ssl-exempt"], "FirewallSslSshProfile-SslExempt"); ok {
				if err = d.Set("ssl_exempt", vv); err != nil {
					return fmt.Errorf("Error reading ssl_exempt: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ssl_exempt: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_exempt"); ok {
			if err = d.Set("ssl_exempt", flattenFirewallSslSshProfileSslExempt(o["ssl-exempt"], d, "ssl_exempt")); err != nil {
				if vv, ok := fortiAPIPatch(o["ssl-exempt"], "FirewallSslSshProfile-SslExempt"); ok {
					if err = d.Set("ssl_exempt", vv); err != nil {
						return fmt.Errorf("Error reading ssl_exempt: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ssl_exempt: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_exemption_ip_rating", flattenFirewallSslSshProfileSslExemptionIpRating(o["ssl-exemption-ip-rating"], d, "ssl_exemption_ip_rating")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-exemption-ip-rating"], "FirewallSslSshProfile-SslExemptionIpRating"); ok {
			if err = d.Set("ssl_exemption_ip_rating", vv); err != nil {
				return fmt.Errorf("Error reading ssl_exemption_ip_rating: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_exemption_ip_rating: %v", err)
		}
	}

	if err = d.Set("ssl_exemption_log", flattenFirewallSslSshProfileSslExemptionLog(o["ssl-exemption-log"], d, "ssl_exemption_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-exemption-log"], "FirewallSslSshProfile-SslExemptionLog"); ok {
			if err = d.Set("ssl_exemption_log", vv); err != nil {
				return fmt.Errorf("Error reading ssl_exemption_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_exemption_log: %v", err)
		}
	}

	if err = d.Set("ssl_exemptions_log", flattenFirewallSslSshProfileSslExemptionsLog(o["ssl-exemptions-log"], d, "ssl_exemptions_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-exemptions-log"], "FirewallSslSshProfile-SslExemptionsLog"); ok {
			if err = d.Set("ssl_exemptions_log", vv); err != nil {
				return fmt.Errorf("Error reading ssl_exemptions_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_exemptions_log: %v", err)
		}
	}

	if err = d.Set("ssl_handshake_log", flattenFirewallSslSshProfileSslHandshakeLog(o["ssl-handshake-log"], d, "ssl_handshake_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-handshake-log"], "FirewallSslSshProfile-SslHandshakeLog"); ok {
			if err = d.Set("ssl_handshake_log", vv); err != nil {
				return fmt.Errorf("Error reading ssl_handshake_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_handshake_log: %v", err)
		}
	}

	if err = d.Set("ssl_negotiation_log", flattenFirewallSslSshProfileSslNegotiationLog(o["ssl-negotiation-log"], d, "ssl_negotiation_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-negotiation-log"], "FirewallSslSshProfile-SslNegotiationLog"); ok {
			if err = d.Set("ssl_negotiation_log", vv); err != nil {
				return fmt.Errorf("Error reading ssl_negotiation_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_negotiation_log: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ssl_server", flattenFirewallSslSshProfileSslServer(o["ssl-server"], d, "ssl_server")); err != nil {
			if vv, ok := fortiAPIPatch(o["ssl-server"], "FirewallSslSshProfile-SslServer"); ok {
				if err = d.Set("ssl_server", vv); err != nil {
					return fmt.Errorf("Error reading ssl_server: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ssl_server: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_server"); ok {
			if err = d.Set("ssl_server", flattenFirewallSslSshProfileSslServer(o["ssl-server"], d, "ssl_server")); err != nil {
				if vv, ok := fortiAPIPatch(o["ssl-server"], "FirewallSslSshProfile-SslServer"); ok {
					if err = d.Set("ssl_server", vv); err != nil {
						return fmt.Errorf("Error reading ssl_server: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ssl_server: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_server_cert_log", flattenFirewallSslSshProfileSslServerCertLog(o["ssl-server-cert-log"], d, "ssl_server_cert_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-cert-log"], "FirewallSslSshProfile-SslServerCertLog"); ok {
			if err = d.Set("ssl_server_cert_log", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_cert_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_cert_log: %v", err)
		}
	}

	if err = d.Set("supported_alpn", flattenFirewallSslSshProfileSupportedAlpn(o["supported-alpn"], d, "supported_alpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["supported-alpn"], "FirewallSslSshProfile-SupportedAlpn"); ok {
			if err = d.Set("supported_alpn", vv); err != nil {
				return fmt.Errorf("Error reading supported_alpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading supported_alpn: %v", err)
		}
	}

	if err = d.Set("untrusted_caname", flattenFirewallSslSshProfileUntrustedCaname(o["untrusted-caname"], d, "untrusted_caname")); err != nil {
		if vv, ok := fortiAPIPatch(o["untrusted-caname"], "FirewallSslSshProfile-UntrustedCaname"); ok {
			if err = d.Set("untrusted_caname", vv); err != nil {
				return fmt.Errorf("Error reading untrusted_caname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading untrusted_caname: %v", err)
		}
	}

	if err = d.Set("use_ssl_server", flattenFirewallSslSshProfileUseSslServer(o["use-ssl-server"], d, "use_ssl_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["use-ssl-server"], "FirewallSslSshProfile-UseSslServer"); ok {
			if err = d.Set("use_ssl_server", vv); err != nil {
				return fmt.Errorf("Error reading use_ssl_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading use_ssl_server: %v", err)
		}
	}

	if err = d.Set("whitelist", flattenFirewallSslSshProfileWhitelist(o["whitelist"], d, "whitelist")); err != nil {
		if vv, ok := fortiAPIPatch(o["whitelist"], "FirewallSslSshProfile-Whitelist"); ok {
			if err = d.Set("whitelist", vv); err != nil {
				return fmt.Errorf("Error reading whitelist: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading whitelist: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ssl_client_certificate", flattenFirewallSslSshProfileSslClientCertificate(o["ssl-client-certificate"], d, "ssl_client_certificate")); err != nil {
			if vv, ok := fortiAPIPatch(o["ssl-client-certificate"], "FirewallSslSshProfile-SslClientCertificate"); ok {
				if err = d.Set("ssl_client_certificate", vv); err != nil {
					return fmt.Errorf("Error reading ssl_client_certificate: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ssl_client_certificate: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_client_certificate"); ok {
			if err = d.Set("ssl_client_certificate", flattenFirewallSslSshProfileSslClientCertificate(o["ssl-client-certificate"], d, "ssl_client_certificate")); err != nil {
				if vv, ok := fortiAPIPatch(o["ssl-client-certificate"], "FirewallSslSshProfile-SslClientCertificate"); ok {
					if err = d.Set("ssl_client_certificate", vv); err != nil {
						return fmt.Errorf("Error reading ssl_client_certificate: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ssl_client_certificate: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenFirewallSslSshProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallSslSshProfileBlockBlacklistedCertificates(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileAllowlist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileBlockBlocklistedCertificates(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileCaname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSslSshProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDot(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cert-validation-failure"], _ = expandFirewallSslSshProfileDotCertValidationFailure(d, i["cert_validation_failure"], pre_append)
	}
	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cert-validation-timeout"], _ = expandFirewallSslSshProfileDotCertValidationTimeout(d, i["cert_validation_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "client_certificate"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["client-certificate"], _ = expandFirewallSslSshProfileDotClientCertificate(d, i["client_certificate"], pre_append)
	}
	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["expired-server-cert"], _ = expandFirewallSslSshProfileDotExpiredServerCert(d, i["expired_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "min_allowed_ssl_version"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["min-allowed-ssl-version"], _ = expandFirewallSslSshProfileDotMinAllowedSslVersion(d, i["min_allowed_ssl_version"], pre_append)
	}
	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["proxy-after-tcp-handshake"], _ = expandFirewallSslSshProfileDotProxyAfterTcpHandshake(d, i["proxy_after_tcp_handshake"], pre_append)
	}
	pre_append = pre + ".0." + "quic"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["quic"], _ = expandFirewallSslSshProfileDotQuic(d, i["quic"], pre_append)
	}
	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["revoked-server-cert"], _ = expandFirewallSslSshProfileDotRevokedServerCert(d, i["revoked_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sni-server-cert-check"], _ = expandFirewallSslSshProfileDotSniServerCertCheck(d, i["sni_server_cert_check"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFirewallSslSshProfileDotStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "udp_not_quic"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["udp-not-quic"], _ = expandFirewallSslSshProfileDotUdpNotQuic(d, i["udp_not_quic"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unsupported-ssl-cipher"], _ = expandFirewallSslSshProfileDotUnsupportedSslCipher(d, i["unsupported_ssl_cipher"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unsupported-ssl-negotiation"], _ = expandFirewallSslSshProfileDotUnsupportedSslNegotiation(d, i["unsupported_ssl_negotiation"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_version"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unsupported-ssl-version"], _ = expandFirewallSslSshProfileDotUnsupportedSslVersion(d, i["unsupported_ssl_version"], pre_append)
	}
	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["untrusted-server-cert"], _ = expandFirewallSslSshProfileDotUntrustedServerCert(d, i["untrusted_server_cert"], pre_append)
	}

	return result, nil
}

func expandFirewallSslSshProfileDotCertValidationFailure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotCertValidationTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotClientCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotExpiredServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotMinAllowedSslVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotProxyAfterTcpHandshake(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotQuic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotRevokedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotSniServerCertCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotUdpNotQuic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotUnsupportedSslCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotUnsupportedSslNegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotUnsupportedSslVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotUntrustedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileEchOuterSni(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallSslSshProfileEchOuterSniName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sni"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sni"], _ = expandFirewallSslSshProfileEchOuterSniSni(d, i["sni"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallSslSshProfileEchOuterSniName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileEchOuterSniSni(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cert-validation-failure"], _ = expandFirewallSslSshProfileFtpsCertValidationFailure(d, i["cert_validation_failure"], pre_append)
	}
	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cert-validation-timeout"], _ = expandFirewallSslSshProfileFtpsCertValidationTimeout(d, i["cert_validation_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "client_certificate"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["client-certificate"], _ = expandFirewallSslSshProfileFtpsClientCertificate(d, i["client_certificate"], pre_append)
	}
	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["expired-server-cert"], _ = expandFirewallSslSshProfileFtpsExpiredServerCert(d, i["expired_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "min_allowed_ssl_version"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["min-allowed-ssl-version"], _ = expandFirewallSslSshProfileFtpsMinAllowedSslVersion(d, i["min_allowed_ssl_version"], pre_append)
	}
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ports"], _ = expandFirewallSslSshProfileFtpsPorts(d, i["ports"], pre_append)
	}
	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["revoked-server-cert"], _ = expandFirewallSslSshProfileFtpsRevokedServerCert(d, i["revoked_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sni-server-cert-check"], _ = expandFirewallSslSshProfileFtpsSniServerCertCheck(d, i["sni_server_cert_check"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFirewallSslSshProfileFtpsStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unsupported-ssl-cipher"], _ = expandFirewallSslSshProfileFtpsUnsupportedSslCipher(d, i["unsupported_ssl_cipher"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unsupported-ssl-negotiation"], _ = expandFirewallSslSshProfileFtpsUnsupportedSslNegotiation(d, i["unsupported_ssl_negotiation"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_version"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unsupported-ssl-version"], _ = expandFirewallSslSshProfileFtpsUnsupportedSslVersion(d, i["unsupported_ssl_version"], pre_append)
	}
	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["untrusted-server-cert"], _ = expandFirewallSslSshProfileFtpsUntrustedServerCert(d, i["untrusted_server_cert"], pre_append)
	}

	return result, nil
}

func expandFirewallSslSshProfileFtpsCertValidationFailure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsCertValidationTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsClientCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsExpiredServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsMinAllowedSslVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandFirewallSslSshProfileFtpsRevokedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsSniServerCertCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsUnsupportedSslCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsUnsupportedSslNegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsUnsupportedSslVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsUntrustedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cert_probe_failure"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cert-probe-failure"], _ = expandFirewallSslSshProfileHttpsCertProbeFailure(d, i["cert_probe_failure"], pre_append)
	}
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cert-validation-failure"], _ = expandFirewallSslSshProfileHttpsCertValidationFailure(d, i["cert_validation_failure"], pre_append)
	}
	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cert-validation-timeout"], _ = expandFirewallSslSshProfileHttpsCertValidationTimeout(d, i["cert_validation_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "client_certificate"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["client-certificate"], _ = expandFirewallSslSshProfileHttpsClientCertificate(d, i["client_certificate"], pre_append)
	}
	pre_append = pre + ".0." + "encrypted_client_hello"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["encrypted-client-hello"], _ = expandFirewallSslSshProfileHttpsEncryptedClientHello(d, i["encrypted_client_hello"], pre_append)
	}
	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["expired-server-cert"], _ = expandFirewallSslSshProfileHttpsExpiredServerCert(d, i["expired_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "min_allowed_ssl_version"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["min-allowed-ssl-version"], _ = expandFirewallSslSshProfileHttpsMinAllowedSslVersion(d, i["min_allowed_ssl_version"], pre_append)
	}
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ports"], _ = expandFirewallSslSshProfileHttpsPorts(d, i["ports"], pre_append)
	}
	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["proxy-after-tcp-handshake"], _ = expandFirewallSslSshProfileHttpsProxyAfterTcpHandshake(d, i["proxy_after_tcp_handshake"], pre_append)
	}
	pre_append = pre + ".0." + "quic"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["quic"], _ = expandFirewallSslSshProfileHttpsQuic(d, i["quic"], pre_append)
	}
	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["revoked-server-cert"], _ = expandFirewallSslSshProfileHttpsRevokedServerCert(d, i["revoked_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sni-server-cert-check"], _ = expandFirewallSslSshProfileHttpsSniServerCertCheck(d, i["sni_server_cert_check"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFirewallSslSshProfileHttpsStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "udp_not_quic"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["udp-not-quic"], _ = expandFirewallSslSshProfileHttpsUdpNotQuic(d, i["udp_not_quic"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unsupported-ssl-cipher"], _ = expandFirewallSslSshProfileHttpsUnsupportedSslCipher(d, i["unsupported_ssl_cipher"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unsupported-ssl-negotiation"], _ = expandFirewallSslSshProfileHttpsUnsupportedSslNegotiation(d, i["unsupported_ssl_negotiation"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_version"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unsupported-ssl-version"], _ = expandFirewallSslSshProfileHttpsUnsupportedSslVersion(d, i["unsupported_ssl_version"], pre_append)
	}
	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["untrusted-server-cert"], _ = expandFirewallSslSshProfileHttpsUntrustedServerCert(d, i["untrusted_server_cert"], pre_append)
	}

	return result, nil
}

func expandFirewallSslSshProfileHttpsCertProbeFailure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsCertValidationFailure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsCertValidationTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsClientCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsEncryptedClientHello(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsExpiredServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsMinAllowedSslVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandFirewallSslSshProfileHttpsProxyAfterTcpHandshake(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsQuic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsRevokedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsSniServerCertCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsUdpNotQuic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsUnsupportedSslCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsUnsupportedSslNegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsUnsupportedSslVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsUntrustedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImaps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cert-validation-failure"], _ = expandFirewallSslSshProfileImapsCertValidationFailure(d, i["cert_validation_failure"], pre_append)
	}
	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cert-validation-timeout"], _ = expandFirewallSslSshProfileImapsCertValidationTimeout(d, i["cert_validation_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "client_certificate"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["client-certificate"], _ = expandFirewallSslSshProfileImapsClientCertificate(d, i["client_certificate"], pre_append)
	}
	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["expired-server-cert"], _ = expandFirewallSslSshProfileImapsExpiredServerCert(d, i["expired_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "min_allowed_ssl_version"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["min-allowed-ssl-version"], _ = expandFirewallSslSshProfileImapsMinAllowedSslVersion(d, i["min_allowed_ssl_version"], pre_append)
	}
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ports"], _ = expandFirewallSslSshProfileImapsPorts(d, i["ports"], pre_append)
	}
	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["proxy-after-tcp-handshake"], _ = expandFirewallSslSshProfileImapsProxyAfterTcpHandshake(d, i["proxy_after_tcp_handshake"], pre_append)
	}
	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["revoked-server-cert"], _ = expandFirewallSslSshProfileImapsRevokedServerCert(d, i["revoked_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sni-server-cert-check"], _ = expandFirewallSslSshProfileImapsSniServerCertCheck(d, i["sni_server_cert_check"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFirewallSslSshProfileImapsStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unsupported-ssl-cipher"], _ = expandFirewallSslSshProfileImapsUnsupportedSslCipher(d, i["unsupported_ssl_cipher"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unsupported-ssl-negotiation"], _ = expandFirewallSslSshProfileImapsUnsupportedSslNegotiation(d, i["unsupported_ssl_negotiation"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_version"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unsupported-ssl-version"], _ = expandFirewallSslSshProfileImapsUnsupportedSslVersion(d, i["unsupported_ssl_version"], pre_append)
	}
	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["untrusted-server-cert"], _ = expandFirewallSslSshProfileImapsUntrustedServerCert(d, i["untrusted_server_cert"], pre_append)
	}

	return result, nil
}

func expandFirewallSslSshProfileImapsCertValidationFailure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsCertValidationTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsClientCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsExpiredServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsMinAllowedSslVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandFirewallSslSshProfileImapsProxyAfterTcpHandshake(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsRevokedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsSniServerCertCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsUnsupportedSslCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsUnsupportedSslNegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsUnsupportedSslVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsUntrustedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileMapiOverHttps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3S(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cert-validation-failure"], _ = expandFirewallSslSshProfilePop3SCertValidationFailure(d, i["cert_validation_failure"], pre_append)
	}
	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cert-validation-timeout"], _ = expandFirewallSslSshProfilePop3SCertValidationTimeout(d, i["cert_validation_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "client_certificate"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["client-certificate"], _ = expandFirewallSslSshProfilePop3SClientCertificate(d, i["client_certificate"], pre_append)
	}
	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["expired-server-cert"], _ = expandFirewallSslSshProfilePop3SExpiredServerCert(d, i["expired_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "min_allowed_ssl_version"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["min-allowed-ssl-version"], _ = expandFirewallSslSshProfilePop3SMinAllowedSslVersion(d, i["min_allowed_ssl_version"], pre_append)
	}
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ports"], _ = expandFirewallSslSshProfilePop3SPorts(d, i["ports"], pre_append)
	}
	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["proxy-after-tcp-handshake"], _ = expandFirewallSslSshProfilePop3SProxyAfterTcpHandshake(d, i["proxy_after_tcp_handshake"], pre_append)
	}
	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["revoked-server-cert"], _ = expandFirewallSslSshProfilePop3SRevokedServerCert(d, i["revoked_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sni-server-cert-check"], _ = expandFirewallSslSshProfilePop3SSniServerCertCheck(d, i["sni_server_cert_check"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFirewallSslSshProfilePop3SStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unsupported-ssl-cipher"], _ = expandFirewallSslSshProfilePop3SUnsupportedSslCipher(d, i["unsupported_ssl_cipher"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unsupported-ssl-negotiation"], _ = expandFirewallSslSshProfilePop3SUnsupportedSslNegotiation(d, i["unsupported_ssl_negotiation"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_version"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unsupported-ssl-version"], _ = expandFirewallSslSshProfilePop3SUnsupportedSslVersion(d, i["unsupported_ssl_version"], pre_append)
	}
	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["untrusted-server-cert"], _ = expandFirewallSslSshProfilePop3SUntrustedServerCert(d, i["untrusted_server_cert"], pre_append)
	}

	return result, nil
}

func expandFirewallSslSshProfilePop3SCertValidationFailure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SCertValidationTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SClientCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SExpiredServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SMinAllowedSslVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandFirewallSslSshProfilePop3SProxyAfterTcpHandshake(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SRevokedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SSniServerCertCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SUnsupportedSslCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SUnsupportedSslNegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SUnsupportedSslVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SUntrustedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileRpcOverHttps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSslSshProfileServerCertMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cert-validation-failure"], _ = expandFirewallSslSshProfileSmtpsCertValidationFailure(d, i["cert_validation_failure"], pre_append)
	}
	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cert-validation-timeout"], _ = expandFirewallSslSshProfileSmtpsCertValidationTimeout(d, i["cert_validation_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "client_certificate"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["client-certificate"], _ = expandFirewallSslSshProfileSmtpsClientCertificate(d, i["client_certificate"], pre_append)
	}
	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["expired-server-cert"], _ = expandFirewallSslSshProfileSmtpsExpiredServerCert(d, i["expired_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "min_allowed_ssl_version"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["min-allowed-ssl-version"], _ = expandFirewallSslSshProfileSmtpsMinAllowedSslVersion(d, i["min_allowed_ssl_version"], pre_append)
	}
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ports"], _ = expandFirewallSslSshProfileSmtpsPorts(d, i["ports"], pre_append)
	}
	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["proxy-after-tcp-handshake"], _ = expandFirewallSslSshProfileSmtpsProxyAfterTcpHandshake(d, i["proxy_after_tcp_handshake"], pre_append)
	}
	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["revoked-server-cert"], _ = expandFirewallSslSshProfileSmtpsRevokedServerCert(d, i["revoked_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sni-server-cert-check"], _ = expandFirewallSslSshProfileSmtpsSniServerCertCheck(d, i["sni_server_cert_check"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFirewallSslSshProfileSmtpsStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unsupported-ssl-cipher"], _ = expandFirewallSslSshProfileSmtpsUnsupportedSslCipher(d, i["unsupported_ssl_cipher"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unsupported-ssl-negotiation"], _ = expandFirewallSslSshProfileSmtpsUnsupportedSslNegotiation(d, i["unsupported_ssl_negotiation"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_version"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unsupported-ssl-version"], _ = expandFirewallSslSshProfileSmtpsUnsupportedSslVersion(d, i["unsupported_ssl_version"], pre_append)
	}
	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["untrusted-server-cert"], _ = expandFirewallSslSshProfileSmtpsUntrustedServerCert(d, i["untrusted_server_cert"], pre_append)
	}

	return result, nil
}

func expandFirewallSslSshProfileSmtpsCertValidationFailure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsCertValidationTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsClientCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsExpiredServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsMinAllowedSslVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandFirewallSslSshProfileSmtpsProxyAfterTcpHandshake(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsRevokedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsSniServerCertCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsUnsupportedSslCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsUnsupportedSslNegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsUnsupportedSslVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsUntrustedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSsh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["inspect-all"], _ = expandFirewallSslSshProfileSshInspectAll(d, i["inspect_all"], pre_append)
	}
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ports"], _ = expandFirewallSslSshProfileSshPorts(d, i["ports"], pre_append)
	}
	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["proxy-after-tcp-handshake"], _ = expandFirewallSslSshProfileSshProxyAfterTcpHandshake(d, i["proxy_after_tcp_handshake"], pre_append)
	}
	pre_append = pre + ".0." + "ssh_algorithm"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ssh-algorithm"], _ = expandFirewallSslSshProfileSshSshAlgorithm(d, i["ssh_algorithm"], pre_append)
	}
	pre_append = pre + ".0." + "ssh_tun_policy_check"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ssh-tun-policy-check"], _ = expandFirewallSslSshProfileSshSshTunPolicyCheck(d, i["ssh_tun_policy_check"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFirewallSslSshProfileSshStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_version"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unsupported-version"], _ = expandFirewallSslSshProfileSshUnsupportedVersion(d, i["unsupported_version"], pre_append)
	}

	return result, nil
}

func expandFirewallSslSshProfileSshInspectAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSshPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandFirewallSslSshProfileSshProxyAfterTcpHandshake(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSshSshAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSshSshTunPolicyCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSshStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSshUnsupportedVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSsl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cert_probe_failure"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cert-probe-failure"], _ = expandFirewallSslSshProfileSslCertProbeFailure(d, i["cert_probe_failure"], pre_append)
	}
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cert-validation-failure"], _ = expandFirewallSslSshProfileSslCertValidationFailure(d, i["cert_validation_failure"], pre_append)
	}
	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cert-validation-timeout"], _ = expandFirewallSslSshProfileSslCertValidationTimeout(d, i["cert_validation_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "client_certificate"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["client-certificate"], _ = expandFirewallSslSshProfileSslClientCertificateU(d, i["client_certificate"], pre_append)
	}
	pre_append = pre + ".0." + "encrypted_client_hello"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["encrypted-client-hello"], _ = expandFirewallSslSshProfileSslEncryptedClientHello(d, i["encrypted_client_hello"], pre_append)
	}
	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["expired-server-cert"], _ = expandFirewallSslSshProfileSslExpiredServerCert(d, i["expired_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["inspect-all"], _ = expandFirewallSslSshProfileSslInspectAll(d, i["inspect_all"], pre_append)
	}
	pre_append = pre + ".0." + "min_allowed_ssl_version"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["min-allowed-ssl-version"], _ = expandFirewallSslSshProfileSslMinAllowedSslVersion(d, i["min_allowed_ssl_version"], pre_append)
	}
	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["revoked-server-cert"], _ = expandFirewallSslSshProfileSslRevokedServerCert(d, i["revoked_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sni-server-cert-check"], _ = expandFirewallSslSshProfileSslSniServerCertCheck(d, i["sni_server_cert_check"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unsupported-ssl-cipher"], _ = expandFirewallSslSshProfileSslUnsupportedSslCipher(d, i["unsupported_ssl_cipher"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unsupported-ssl-negotiation"], _ = expandFirewallSslSshProfileSslUnsupportedSslNegotiation(d, i["unsupported_ssl_negotiation"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_version"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unsupported-ssl-version"], _ = expandFirewallSslSshProfileSslUnsupportedSslVersion(d, i["unsupported_ssl_version"], pre_append)
	}
	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["untrusted-server-cert"], _ = expandFirewallSslSshProfileSslUntrustedServerCert(d, i["untrusted_server_cert"], pre_append)
	}

	return result, nil
}

func expandFirewallSslSshProfileSslCertProbeFailure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslCertValidationFailure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslCertValidationTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslClientCertificateU(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslEncryptedClientHello(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExpiredServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslInspectAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslMinAllowedSslVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslRevokedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslSniServerCertCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslUnsupportedSslCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslUnsupportedSslNegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslUnsupportedSslVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslUntrustedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslAnomaliesLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslAnomalyLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExempt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["address"], _ = expandFirewallSslSshProfileSslExemptAddress(d, i["address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["address6"], _ = expandFirewallSslSshProfileSslExemptAddress6(d, i["address6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortiguard_category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fortiguard-category"], _ = expandFirewallSslSshProfileSslExemptFortiguardCategory(d, i["fortiguard_category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFirewallSslSshProfileSslExemptId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["regex"], _ = expandFirewallSslSshProfileSslExemptRegex(d, i["regex"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandFirewallSslSshProfileSslExemptType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wildcard_fqdn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["wildcard-fqdn"], _ = expandFirewallSslSshProfileSslExemptWildcardFqdn(d, i["wildcard_fqdn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "finger_print_category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["finger-print-category"], _ = expandFirewallSslSshProfileSslExemptFingerPrintCategory(d, i["finger_print_category"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallSslSshProfileSslExemptAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSslSshProfileSslExemptAddress6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSslSshProfileSslExemptFortiguardCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSslSshProfileSslExemptId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExemptRegex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExemptType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExemptWildcardFqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSslSshProfileSslExemptFingerPrintCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExemptionIpRating(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExemptionLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExemptionsLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslHandshakeLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslNegotiationLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftps_client_certificate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ftps-client-certificate"], _ = expandFirewallSslSshProfileSslServerFtpsClientCertificate(d, i["ftps_client_certificate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_client_certificate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["https-client-certificate"], _ = expandFirewallSslSshProfileSslServerHttpsClientCertificate(d, i["https_client_certificate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFirewallSslSshProfileSslServerId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "imaps_client_certificate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["imaps-client-certificate"], _ = expandFirewallSslSshProfileSslServerImapsClientCertificate(d, i["imaps_client_certificate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandFirewallSslSshProfileSslServerIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pop3s_client_certificate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pop3s-client-certificate"], _ = expandFirewallSslSshProfileSslServerPop3SClientCertificate(d, i["pop3s_client_certificate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "smtps_client_certificate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["smtps-client-certificate"], _ = expandFirewallSslSshProfileSslServerSmtpsClientCertificate(d, i["smtps_client_certificate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_other_client_certificate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-other-client-certificate"], _ = expandFirewallSslSshProfileSslServerSslOtherClientCertificate(d, i["ssl_other_client_certificate"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallSslSshProfileSslServerFtpsClientCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerHttpsClientCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerImapsClientCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerPop3SClientCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerSmtpsClientCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerSslOtherClientCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerCertLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSupportedAlpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileUntrustedCaname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSslSshProfileUseSslServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileWhitelist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslClientCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "caname"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["caname"], _ = expandFirewallSslSshProfileSslClientCertificateCaname(d, i["caname"], pre_append)
	}
	pre_append = pre + ".0." + "cert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cert"], _ = expandFirewallSslSshProfileSslClientCertificateCert(d, i["cert"], pre_append)
	}
	pre_append = pre + ".0." + "keyring_list"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keyring-list"], _ = expandFirewallSslSshProfileSslClientCertificateKeyringList(d, i["keyring_list"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFirewallSslSshProfileSslClientCertificateStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandFirewallSslSshProfileSslClientCertificateCaname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSslSshProfileSslClientCertificateCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSslSshProfileSslClientCertificateKeyringList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSslSshProfileSslClientCertificateStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallSslSshProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("block_blacklisted_certificates"); ok || d.HasChange("block_blacklisted_certificates") {
		t, err := expandFirewallSslSshProfileBlockBlacklistedCertificates(d, v, "block_blacklisted_certificates")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-blacklisted-certificates"] = t
		}
	}

	if v, ok := d.GetOk("allowlist"); ok || d.HasChange("allowlist") {
		t, err := expandFirewallSslSshProfileAllowlist(d, v, "allowlist")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowlist"] = t
		}
	}

	if v, ok := d.GetOk("block_blocklisted_certificates"); ok || d.HasChange("block_blocklisted_certificates") {
		t, err := expandFirewallSslSshProfileBlockBlocklistedCertificates(d, v, "block_blocklisted_certificates")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-blocklisted-certificates"] = t
		}
	}

	if v, ok := d.GetOk("caname"); ok || d.HasChange("caname") {
		t, err := expandFirewallSslSshProfileCaname(d, v, "caname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["caname"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandFirewallSslSshProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("dot"); ok || d.HasChange("dot") {
		t, err := expandFirewallSslSshProfileDot(d, v, "dot")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dot"] = t
		}
	}

	if v, ok := d.GetOk("ech_outer_sni"); ok || d.HasChange("ech_outer_sni") {
		t, err := expandFirewallSslSshProfileEchOuterSni(d, v, "ech_outer_sni")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ech-outer-sni"] = t
		}
	}

	if v, ok := d.GetOk("ftps"); ok || d.HasChange("ftps") {
		t, err := expandFirewallSslSshProfileFtps(d, v, "ftps")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftps"] = t
		}
	}

	if v, ok := d.GetOk("https"); ok || d.HasChange("https") {
		t, err := expandFirewallSslSshProfileHttps(d, v, "https")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https"] = t
		}
	}

	if v, ok := d.GetOk("imaps"); ok || d.HasChange("imaps") {
		t, err := expandFirewallSslSshProfileImaps(d, v, "imaps")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["imaps"] = t
		}
	}

	if v, ok := d.GetOk("mapi_over_https"); ok || d.HasChange("mapi_over_https") {
		t, err := expandFirewallSslSshProfileMapiOverHttps(d, v, "mapi_over_https")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mapi-over-https"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallSslSshProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("pop3s"); ok || d.HasChange("pop3s") {
		t, err := expandFirewallSslSshProfilePop3S(d, v, "pop3s")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pop3s"] = t
		}
	}

	if v, ok := d.GetOk("rpc_over_https"); ok || d.HasChange("rpc_over_https") {
		t, err := expandFirewallSslSshProfileRpcOverHttps(d, v, "rpc_over_https")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rpc-over-https"] = t
		}
	}

	if v, ok := d.GetOk("server_cert"); ok || d.HasChange("server_cert") {
		t, err := expandFirewallSslSshProfileServerCert(d, v, "server_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-cert"] = t
		}
	}

	if v, ok := d.GetOk("server_cert_mode"); ok || d.HasChange("server_cert_mode") {
		t, err := expandFirewallSslSshProfileServerCertMode(d, v, "server_cert_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-cert-mode"] = t
		}
	}

	if v, ok := d.GetOk("smtps"); ok || d.HasChange("smtps") {
		t, err := expandFirewallSslSshProfileSmtps(d, v, "smtps")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smtps"] = t
		}
	}

	if v, ok := d.GetOk("ssh"); ok || d.HasChange("ssh") {
		t, err := expandFirewallSslSshProfileSsh(d, v, "ssh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh"] = t
		}
	}

	if v, ok := d.GetOk("ssl"); ok || d.HasChange("ssl") {
		t, err := expandFirewallSslSshProfileSsl(d, v, "ssl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl"] = t
		}
	}

	if v, ok := d.GetOk("ssl_anomalies_log"); ok || d.HasChange("ssl_anomalies_log") {
		t, err := expandFirewallSslSshProfileSslAnomaliesLog(d, v, "ssl_anomalies_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-anomalies-log"] = t
		}
	}

	if v, ok := d.GetOk("ssl_anomaly_log"); ok || d.HasChange("ssl_anomaly_log") {
		t, err := expandFirewallSslSshProfileSslAnomalyLog(d, v, "ssl_anomaly_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-anomaly-log"] = t
		}
	}

	if v, ok := d.GetOk("ssl_exempt"); ok || d.HasChange("ssl_exempt") {
		t, err := expandFirewallSslSshProfileSslExempt(d, v, "ssl_exempt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-exempt"] = t
		}
	}

	if v, ok := d.GetOk("ssl_exemption_ip_rating"); ok || d.HasChange("ssl_exemption_ip_rating") {
		t, err := expandFirewallSslSshProfileSslExemptionIpRating(d, v, "ssl_exemption_ip_rating")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-exemption-ip-rating"] = t
		}
	}

	if v, ok := d.GetOk("ssl_exemption_log"); ok || d.HasChange("ssl_exemption_log") {
		t, err := expandFirewallSslSshProfileSslExemptionLog(d, v, "ssl_exemption_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-exemption-log"] = t
		}
	}

	if v, ok := d.GetOk("ssl_exemptions_log"); ok || d.HasChange("ssl_exemptions_log") {
		t, err := expandFirewallSslSshProfileSslExemptionsLog(d, v, "ssl_exemptions_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-exemptions-log"] = t
		}
	}

	if v, ok := d.GetOk("ssl_handshake_log"); ok || d.HasChange("ssl_handshake_log") {
		t, err := expandFirewallSslSshProfileSslHandshakeLog(d, v, "ssl_handshake_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-handshake-log"] = t
		}
	}

	if v, ok := d.GetOk("ssl_negotiation_log"); ok || d.HasChange("ssl_negotiation_log") {
		t, err := expandFirewallSslSshProfileSslNegotiationLog(d, v, "ssl_negotiation_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-negotiation-log"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server"); ok || d.HasChange("ssl_server") {
		t, err := expandFirewallSslSshProfileSslServer(d, v, "ssl_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_cert_log"); ok || d.HasChange("ssl_server_cert_log") {
		t, err := expandFirewallSslSshProfileSslServerCertLog(d, v, "ssl_server_cert_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-cert-log"] = t
		}
	}

	if v, ok := d.GetOk("supported_alpn"); ok || d.HasChange("supported_alpn") {
		t, err := expandFirewallSslSshProfileSupportedAlpn(d, v, "supported_alpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["supported-alpn"] = t
		}
	}

	if v, ok := d.GetOk("untrusted_caname"); ok || d.HasChange("untrusted_caname") {
		t, err := expandFirewallSslSshProfileUntrustedCaname(d, v, "untrusted_caname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["untrusted-caname"] = t
		}
	}

	if v, ok := d.GetOk("use_ssl_server"); ok || d.HasChange("use_ssl_server") {
		t, err := expandFirewallSslSshProfileUseSslServer(d, v, "use_ssl_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-ssl-server"] = t
		}
	}

	if v, ok := d.GetOk("whitelist"); ok || d.HasChange("whitelist") {
		t, err := expandFirewallSslSshProfileWhitelist(d, v, "whitelist")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["whitelist"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_certificate"); ok || d.HasChange("ssl_client_certificate") {
		t, err := expandFirewallSslSshProfileSslClientCertificate(d, v, "ssl_client_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-certificate"] = t
		}
	}

	return &obj, nil
}
