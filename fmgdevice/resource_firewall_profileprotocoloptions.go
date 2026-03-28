// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure protocol options.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallProfileProtocolOptions() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallProfileProtocolOptionsCreate,
		Read:   resourceFirewallProfileProtocolOptionsRead,
		Update: resourceFirewallProfileProtocolOptionsUpdate,
		Delete: resourceFirewallProfileProtocolOptionsDelete,

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
			"cifs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain_controller": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"file_filter": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"entries": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"action": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"comment": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"direction": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"file_type": &schema.Schema{
													Type:     schema.TypeSet,
													Elem:     &schema.Schema{Type: schema.TypeString},
													Optional: true,
													Computed: true,
												},
												"filter": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"protocol": &schema.Schema{
													Type:     schema.TypeSet,
													Elem:     &schema.Schema{Type: schema.TypeString},
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"log": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ports": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"server_credential_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"server_keytab": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"keytab": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"password": &schema.Schema{
										Type:      schema.TypeSet,
										Elem:      &schema.Schema{Type: schema.TypeString},
										Optional:  true,
										Sensitive: true,
										Computed:  true,
									},
									"principal": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tcp_window_maximum": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"tcp_window_minimum": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"tcp_window_size": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"tcp_window_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
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
			"feature_set": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ftp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comfort_amount": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"comfort_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"explicit_ftp_tls": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"inspect_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ports": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_offloaded": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"stream_based_uncompressed_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"tcp_window_maximum": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"tcp_window_minimum": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"tcp_window_size": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"tcp_window_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"http": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address_ip_rating": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_page_status_code": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"comfort_amount": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"comfort_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"fortinet_bar": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortinet_bar_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"domain_fronting": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"h2c": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"http_09": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"inspect_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ports": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"post_lang": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"proxy_after_tcp_handshake": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"range_block": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"retry_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_offloaded": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"stream_based_uncompressed_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"streaming_content_bypass": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"strip_x_forwarded_for": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"switching_protocols": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tcp_window_maximum": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"tcp_window_minimum": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"tcp_window_size": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"tcp_window_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tunnel_non_http": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"unknown_content_encoding": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unknown_http_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"verify_dns_for_policy_matching": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dns_protection": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"encrypted_file": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"encrypted_file_log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"imap": &schema.Schema{
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
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
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
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_offloaded": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"address_ip_rating": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"mail_signature": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"signature": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"mapi": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ports": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nntp": &schema.Schema{
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
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
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
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"oversize_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pop3": &schema.Schema{
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
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
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
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_offloaded": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"replacemsg_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"rpc_over_http": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"smtp": &schema.Schema{
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
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
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
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"server_busy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_offloaded": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
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
						"comfort_amount": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"comfort_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_offloaded": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"stream_based_uncompressed_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"tcp_window_maximum": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"tcp_window_minimum": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"tcp_window_size": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"tcp_window_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"explicit_ftp_tls": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"switching_protocols_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_redirect": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
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
		},
	}
}

func resourceFirewallProfileProtocolOptionsCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallProfileProtocolOptions(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallProfileProtocolOptions resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallProfileProtocolOptions(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallProfileProtocolOptions(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallProfileProtocolOptions resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallProfileProtocolOptions(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallProfileProtocolOptions resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallProfileProtocolOptionsRead(d, m)
}

func resourceFirewallProfileProtocolOptionsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallProfileProtocolOptions(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallProfileProtocolOptions resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallProfileProtocolOptions(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallProfileProtocolOptions resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallProfileProtocolOptionsRead(d, m)
}

func resourceFirewallProfileProtocolOptionsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallProfileProtocolOptions(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallProfileProtocolOptions resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallProfileProtocolOptionsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallProfileProtocolOptions(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallProfileProtocolOptions resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallProfileProtocolOptions(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallProfileProtocolOptions resource from API: %v", err)
	}
	return nil
}

func flattenFirewallProfileProtocolOptionsCifs(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "domain_controller"
	if _, ok := i["domain-controller"]; ok {
		result["domain_controller"] = flattenFirewallProfileProtocolOptionsCifsDomainController(i["domain-controller"], d, pre_append)
	}

	pre_append = pre + ".0." + "file_filter"
	if _, ok := i["file-filter"]; ok {
		result["file_filter"] = flattenFirewallProfileProtocolOptionsCifsFileFilter(i["file-filter"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenFirewallProfileProtocolOptionsCifsOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = flattenFirewallProfileProtocolOptionsCifsOversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallProfileProtocolOptionsCifsPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = flattenFirewallProfileProtocolOptionsCifsScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "server_credential_type"
	if _, ok := i["server-credential-type"]; ok {
		result["server_credential_type"] = flattenFirewallProfileProtocolOptionsCifsServerCredentialType(i["server-credential-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "server_keytab"
	if _, ok := i["server-keytab"]; ok {
		result["server_keytab"] = flattenFirewallProfileProtocolOptionsCifsServerKeytab(i["server-keytab"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallProfileProtocolOptionsCifsStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_maximum"
	if _, ok := i["tcp-window-maximum"]; ok {
		result["tcp_window_maximum"] = flattenFirewallProfileProtocolOptionsCifsTcpWindowMaximum(i["tcp-window-maximum"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_minimum"
	if _, ok := i["tcp-window-minimum"]; ok {
		result["tcp_window_minimum"] = flattenFirewallProfileProtocolOptionsCifsTcpWindowMinimum(i["tcp-window-minimum"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_size"
	if _, ok := i["tcp-window-size"]; ok {
		result["tcp_window_size"] = flattenFirewallProfileProtocolOptionsCifsTcpWindowSize(i["tcp-window-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_type"
	if _, ok := i["tcp-window-type"]; ok {
		result["tcp_window_type"] = flattenFirewallProfileProtocolOptionsCifsTcpWindowType(i["tcp-window-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = flattenFirewallProfileProtocolOptionsCifsUncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = flattenFirewallProfileProtocolOptionsCifsUncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallProfileProtocolOptionsCifsDomainController(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileProtocolOptionsCifsFileFilter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "entries"
	if _, ok := i["entries"]; ok {
		result["entries"] = flattenFirewallProfileProtocolOptionsCifsFileFilterEntries(i["entries"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenFirewallProfileProtocolOptionsCifsFileFilterLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallProfileProtocolOptionsCifsFileFilterStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallProfileProtocolOptionsCifsFileFilterEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenFirewallProfileProtocolOptionsCifsFileFilterEntriesAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "FirewallProfileProtocolOptionsCifsFileFilter-Entries-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenFirewallProfileProtocolOptionsCifsFileFilterEntriesComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "FirewallProfileProtocolOptionsCifsFileFilter-Entries-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			v := flattenFirewallProfileProtocolOptionsCifsFileFilterEntriesDirection(i["direction"], d, pre_append)
			tmp["direction"] = fortiAPISubPartPatch(v, "FirewallProfileProtocolOptionsCifsFileFilter-Entries-Direction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_type"
		if _, ok := i["file-type"]; ok {
			v := flattenFirewallProfileProtocolOptionsCifsFileFilterEntriesFileType(i["file-type"], d, pre_append)
			tmp["file_type"] = fortiAPISubPartPatch(v, "FirewallProfileProtocolOptionsCifsFileFilter-Entries-FileType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := i["filter"]; ok {
			v := flattenFirewallProfileProtocolOptionsCifsFileFilterEntriesFilter(i["filter"], d, pre_append)
			tmp["filter"] = fortiAPISubPartPatch(v, "FirewallProfileProtocolOptionsCifsFileFilter-Entries-Filter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenFirewallProfileProtocolOptionsCifsFileFilterEntriesProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "FirewallProfileProtocolOptionsCifsFileFilter-Entries-Protocol")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallProfileProtocolOptionsCifsFileFilterEntriesAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsFileFilterEntriesComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsFileFilterEntriesDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsFileFilterEntriesFileType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileProtocolOptionsCifsFileFilterEntriesFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsFileFilterEntriesProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileProtocolOptionsCifsFileFilterLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsFileFilterStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileProtocolOptionsCifsOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenFirewallProfileProtocolOptionsCifsScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsServerCredentialType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsServerKeytab(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keytab"
		if _, ok := i["keytab"]; ok {
			v := flattenFirewallProfileProtocolOptionsCifsServerKeytabKeytab(i["keytab"], d, pre_append)
			tmp["keytab"] = fortiAPISubPartPatch(v, "FirewallProfileProtocolOptionsCifs-ServerKeytab-Keytab")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "principal"
		if _, ok := i["principal"]; ok {
			v := flattenFirewallProfileProtocolOptionsCifsServerKeytabPrincipal(i["principal"], d, pre_append)
			tmp["principal"] = fortiAPISubPartPatch(v, "FirewallProfileProtocolOptionsCifs-ServerKeytab-Principal")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallProfileProtocolOptionsCifsServerKeytabKeytab(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsServerKeytabPrincipal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsTcpWindowMaximum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsTcpWindowMinimum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsTcpWindowSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsTcpWindowType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsUncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsUncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsDns(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallProfileProtocolOptionsDnsPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallProfileProtocolOptionsDnsStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallProfileProtocolOptionsDnsPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenFirewallProfileProtocolOptionsDnsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsFeatureSet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsFtp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "comfort_amount"
	if _, ok := i["comfort-amount"]; ok {
		result["comfort_amount"] = flattenFirewallProfileProtocolOptionsFtpComfortAmount(i["comfort-amount"], d, pre_append)
	}

	pre_append = pre + ".0." + "comfort_interval"
	if _, ok := i["comfort-interval"]; ok {
		result["comfort_interval"] = flattenFirewallProfileProtocolOptionsFtpComfortInterval(i["comfort-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "explicit_ftp_tls"
	if _, ok := i["explicit-ftp-tls"]; ok {
		result["explicit_ftp_tls"] = flattenFirewallProfileProtocolOptionsFtpExplicitFtpTls(i["explicit-ftp-tls"], d, pre_append)
	}

	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = flattenFirewallProfileProtocolOptionsFtpInspectAll(i["inspect-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenFirewallProfileProtocolOptionsFtpOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = flattenFirewallProfileProtocolOptionsFtpOversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallProfileProtocolOptionsFtpPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = flattenFirewallProfileProtocolOptionsFtpScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_offloaded"
	if _, ok := i["ssl-offloaded"]; ok {
		result["ssl_offloaded"] = flattenFirewallProfileProtocolOptionsFtpSslOffloaded(i["ssl-offloaded"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallProfileProtocolOptionsFtpStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "stream_based_uncompressed_limit"
	if _, ok := i["stream-based-uncompressed-limit"]; ok {
		result["stream_based_uncompressed_limit"] = flattenFirewallProfileProtocolOptionsFtpStreamBasedUncompressedLimit(i["stream-based-uncompressed-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_maximum"
	if _, ok := i["tcp-window-maximum"]; ok {
		result["tcp_window_maximum"] = flattenFirewallProfileProtocolOptionsFtpTcpWindowMaximum(i["tcp-window-maximum"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_minimum"
	if _, ok := i["tcp-window-minimum"]; ok {
		result["tcp_window_minimum"] = flattenFirewallProfileProtocolOptionsFtpTcpWindowMinimum(i["tcp-window-minimum"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_size"
	if _, ok := i["tcp-window-size"]; ok {
		result["tcp_window_size"] = flattenFirewallProfileProtocolOptionsFtpTcpWindowSize(i["tcp-window-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_type"
	if _, ok := i["tcp-window-type"]; ok {
		result["tcp_window_type"] = flattenFirewallProfileProtocolOptionsFtpTcpWindowType(i["tcp-window-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = flattenFirewallProfileProtocolOptionsFtpUncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = flattenFirewallProfileProtocolOptionsFtpUncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallProfileProtocolOptionsFtpComfortAmount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsFtpComfortInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsFtpExplicitFtpTls(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsFtpInspectAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsFtpOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileProtocolOptionsFtpOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsFtpPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenFirewallProfileProtocolOptionsFtpScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsFtpSslOffloaded(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsFtpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsFtpStreamBasedUncompressedLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsFtpTcpWindowMaximum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsFtpTcpWindowMinimum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsFtpTcpWindowSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsFtpTcpWindowType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsFtpUncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsFtpUncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "address_ip_rating"
	if _, ok := i["address-ip-rating"]; ok {
		result["address_ip_rating"] = flattenFirewallProfileProtocolOptionsHttpAddressIpRating(i["address-ip-rating"], d, pre_append)
	}

	pre_append = pre + ".0." + "block_page_status_code"
	if _, ok := i["block-page-status-code"]; ok {
		result["block_page_status_code"] = flattenFirewallProfileProtocolOptionsHttpBlockPageStatusCode(i["block-page-status-code"], d, pre_append)
	}

	pre_append = pre + ".0." + "comfort_amount"
	if _, ok := i["comfort-amount"]; ok {
		result["comfort_amount"] = flattenFirewallProfileProtocolOptionsHttpComfortAmount(i["comfort-amount"], d, pre_append)
	}

	pre_append = pre + ".0." + "comfort_interval"
	if _, ok := i["comfort-interval"]; ok {
		result["comfort_interval"] = flattenFirewallProfileProtocolOptionsHttpComfortInterval(i["comfort-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortinet_bar"
	if _, ok := i["fortinet-bar"]; ok {
		result["fortinet_bar"] = flattenFirewallProfileProtocolOptionsHttpFortinetBar(i["fortinet-bar"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortinet_bar_port"
	if _, ok := i["fortinet-bar-port"]; ok {
		result["fortinet_bar_port"] = flattenFirewallProfileProtocolOptionsHttpFortinetBarPort(i["fortinet-bar-port"], d, pre_append)
	}

	pre_append = pre + ".0." + "domain_fronting"
	if _, ok := i["domain-fronting"]; ok {
		result["domain_fronting"] = flattenFirewallProfileProtocolOptionsHttpDomainFronting(i["domain-fronting"], d, pre_append)
	}

	pre_append = pre + ".0." + "h2c"
	if _, ok := i["h2c"]; ok {
		result["h2c"] = flattenFirewallProfileProtocolOptionsHttpH2C(i["h2c"], d, pre_append)
	}

	pre_append = pre + ".0." + "http_09"
	if _, ok := i["http-0.9"]; ok {
		result["http_09"] = flattenFirewallProfileProtocolOptionsHttpHttp09(i["http-0.9"], d, pre_append)
	}

	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = flattenFirewallProfileProtocolOptionsHttpInspectAll(i["inspect-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenFirewallProfileProtocolOptionsHttpOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = flattenFirewallProfileProtocolOptionsHttpOversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallProfileProtocolOptionsHttpPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "post_lang"
	if _, ok := i["post-lang"]; ok {
		result["post_lang"] = flattenFirewallProfileProtocolOptionsHttpPostLang(i["post-lang"], d, pre_append)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = flattenFirewallProfileProtocolOptionsHttpProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append)
	}

	pre_append = pre + ".0." + "range_block"
	if _, ok := i["range-block"]; ok {
		result["range_block"] = flattenFirewallProfileProtocolOptionsHttpRangeBlock(i["range-block"], d, pre_append)
	}

	pre_append = pre + ".0." + "retry_count"
	if _, ok := i["retry-count"]; ok {
		result["retry_count"] = flattenFirewallProfileProtocolOptionsHttpRetryCount(i["retry-count"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = flattenFirewallProfileProtocolOptionsHttpScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_offloaded"
	if _, ok := i["ssl-offloaded"]; ok {
		result["ssl_offloaded"] = flattenFirewallProfileProtocolOptionsHttpSslOffloaded(i["ssl-offloaded"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallProfileProtocolOptionsHttpStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "stream_based_uncompressed_limit"
	if _, ok := i["stream-based-uncompressed-limit"]; ok {
		result["stream_based_uncompressed_limit"] = flattenFirewallProfileProtocolOptionsHttpStreamBasedUncompressedLimit(i["stream-based-uncompressed-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "streaming_content_bypass"
	if _, ok := i["streaming-content-bypass"]; ok {
		result["streaming_content_bypass"] = flattenFirewallProfileProtocolOptionsHttpStreamingContentBypass(i["streaming-content-bypass"], d, pre_append)
	}

	pre_append = pre + ".0." + "strip_x_forwarded_for"
	if _, ok := i["strip-x-forwarded-for"]; ok {
		result["strip_x_forwarded_for"] = flattenFirewallProfileProtocolOptionsHttpStripXForwardedFor(i["strip-x-forwarded-for"], d, pre_append)
	}

	pre_append = pre + ".0." + "switching_protocols"
	if _, ok := i["switching-protocols"]; ok {
		result["switching_protocols"] = flattenFirewallProfileProtocolOptionsHttpSwitchingProtocols(i["switching-protocols"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_maximum"
	if _, ok := i["tcp-window-maximum"]; ok {
		result["tcp_window_maximum"] = flattenFirewallProfileProtocolOptionsHttpTcpWindowMaximum(i["tcp-window-maximum"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_minimum"
	if _, ok := i["tcp-window-minimum"]; ok {
		result["tcp_window_minimum"] = flattenFirewallProfileProtocolOptionsHttpTcpWindowMinimum(i["tcp-window-minimum"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_size"
	if _, ok := i["tcp-window-size"]; ok {
		result["tcp_window_size"] = flattenFirewallProfileProtocolOptionsHttpTcpWindowSize(i["tcp-window-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_type"
	if _, ok := i["tcp-window-type"]; ok {
		result["tcp_window_type"] = flattenFirewallProfileProtocolOptionsHttpTcpWindowType(i["tcp-window-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "tunnel_non_http"
	if _, ok := i["tunnel-non-http"]; ok {
		result["tunnel_non_http"] = flattenFirewallProfileProtocolOptionsHttpTunnelNonHttp(i["tunnel-non-http"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = flattenFirewallProfileProtocolOptionsHttpUncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = flattenFirewallProfileProtocolOptionsHttpUncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "unknown_content_encoding"
	if _, ok := i["unknown-content-encoding"]; ok {
		result["unknown_content_encoding"] = flattenFirewallProfileProtocolOptionsHttpUnknownContentEncoding(i["unknown-content-encoding"], d, pre_append)
	}

	pre_append = pre + ".0." + "unknown_http_version"
	if _, ok := i["unknown-http-version"]; ok {
		result["unknown_http_version"] = flattenFirewallProfileProtocolOptionsHttpUnknownHttpVersion(i["unknown-http-version"], d, pre_append)
	}

	pre_append = pre + ".0." + "verify_dns_for_policy_matching"
	if _, ok := i["verify-dns-for-policy-matching"]; ok {
		result["verify_dns_for_policy_matching"] = flattenFirewallProfileProtocolOptionsHttpVerifyDnsForPolicyMatching(i["verify-dns-for-policy-matching"], d, pre_append)
	}

	pre_append = pre + ".0." + "dns_protection"
	if _, ok := i["dns-protection"]; ok {
		result["dns_protection"] = flattenFirewallProfileProtocolOptionsHttpDnsProtection(i["dns-protection"], d, pre_append)
	}

	pre_append = pre + ".0." + "encrypted_file"
	if _, ok := i["encrypted-file"]; ok {
		result["encrypted_file"] = flattenFirewallProfileProtocolOptionsHttpEncryptedFile(i["encrypted-file"], d, pre_append)
	}

	pre_append = pre + ".0." + "encrypted_file_log"
	if _, ok := i["encrypted-file-log"]; ok {
		result["encrypted_file_log"] = flattenFirewallProfileProtocolOptionsHttpEncryptedFileLog(i["encrypted-file-log"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallProfileProtocolOptionsHttpAddressIpRating(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpBlockPageStatusCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpComfortAmount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpComfortInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpFortinetBar(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpFortinetBarPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpDomainFronting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpH2C(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpHttp09(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpInspectAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileProtocolOptionsHttpOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenFirewallProfileProtocolOptionsHttpPostLang(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileProtocolOptionsHttpProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpRangeBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpRetryCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpSslOffloaded(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpStreamBasedUncompressedLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpStreamingContentBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpStripXForwardedFor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpSwitchingProtocols(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpTcpWindowMaximum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpTcpWindowMinimum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpTcpWindowSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpTcpWindowType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpTunnelNonHttp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpUncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpUncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpUnknownContentEncoding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpUnknownHttpVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpVerifyDnsForPolicyMatching(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpDnsProtection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpEncryptedFile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpEncryptedFileLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsImap(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = flattenFirewallProfileProtocolOptionsImapInspectAll(i["inspect-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenFirewallProfileProtocolOptionsImapOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = flattenFirewallProfileProtocolOptionsImapOversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallProfileProtocolOptionsImapPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = flattenFirewallProfileProtocolOptionsImapProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = flattenFirewallProfileProtocolOptionsImapScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_offloaded"
	if _, ok := i["ssl-offloaded"]; ok {
		result["ssl_offloaded"] = flattenFirewallProfileProtocolOptionsImapSslOffloaded(i["ssl-offloaded"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallProfileProtocolOptionsImapStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = flattenFirewallProfileProtocolOptionsImapUncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = flattenFirewallProfileProtocolOptionsImapUncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "address_ip_rating"
	if _, ok := i["address-ip-rating"]; ok {
		result["address_ip_rating"] = flattenFirewallProfileProtocolOptionsImapAddressIpRating(i["address-ip-rating"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallProfileProtocolOptionsImapInspectAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsImapOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileProtocolOptionsImapOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsImapPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenFirewallProfileProtocolOptionsImapProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsImapScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsImapSslOffloaded(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsImapStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsImapUncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsImapUncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsImapAddressIpRating(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsMailSignature(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "signature"
	if _, ok := i["signature"]; ok {
		result["signature"] = flattenFirewallProfileProtocolOptionsMailSignatureSignature(i["signature"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallProfileProtocolOptionsMailSignatureStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallProfileProtocolOptionsMailSignatureSignature(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsMailSignatureStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsMapi(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenFirewallProfileProtocolOptionsMapiOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = flattenFirewallProfileProtocolOptionsMapiOversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallProfileProtocolOptionsMapiPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = flattenFirewallProfileProtocolOptionsMapiScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallProfileProtocolOptionsMapiStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = flattenFirewallProfileProtocolOptionsMapiUncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = flattenFirewallProfileProtocolOptionsMapiUncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallProfileProtocolOptionsMapiOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileProtocolOptionsMapiOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsMapiPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenFirewallProfileProtocolOptionsMapiScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsMapiStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsMapiUncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsMapiUncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsNntp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = flattenFirewallProfileProtocolOptionsNntpInspectAll(i["inspect-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenFirewallProfileProtocolOptionsNntpOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = flattenFirewallProfileProtocolOptionsNntpOversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallProfileProtocolOptionsNntpPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = flattenFirewallProfileProtocolOptionsNntpProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = flattenFirewallProfileProtocolOptionsNntpScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallProfileProtocolOptionsNntpStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = flattenFirewallProfileProtocolOptionsNntpUncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = flattenFirewallProfileProtocolOptionsNntpUncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallProfileProtocolOptionsNntpInspectAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsNntpOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileProtocolOptionsNntpOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsNntpPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenFirewallProfileProtocolOptionsNntpProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsNntpScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsNntpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsNntpUncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsNntpUncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsOversizeLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsPop3(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = flattenFirewallProfileProtocolOptionsPop3InspectAll(i["inspect-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenFirewallProfileProtocolOptionsPop3Options(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = flattenFirewallProfileProtocolOptionsPop3OversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallProfileProtocolOptionsPop3Ports(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = flattenFirewallProfileProtocolOptionsPop3ProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = flattenFirewallProfileProtocolOptionsPop3ScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_offloaded"
	if _, ok := i["ssl-offloaded"]; ok {
		result["ssl_offloaded"] = flattenFirewallProfileProtocolOptionsPop3SslOffloaded(i["ssl-offloaded"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallProfileProtocolOptionsPop3Status(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = flattenFirewallProfileProtocolOptionsPop3UncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = flattenFirewallProfileProtocolOptionsPop3UncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallProfileProtocolOptionsPop3InspectAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsPop3Options(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileProtocolOptionsPop3OversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsPop3Ports(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenFirewallProfileProtocolOptionsPop3ProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsPop3ScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsPop3SslOffloaded(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsPop3Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsPop3UncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsPop3UncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileProtocolOptionsRpcOverHttp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSmtp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = flattenFirewallProfileProtocolOptionsSmtpInspectAll(i["inspect-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenFirewallProfileProtocolOptionsSmtpOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = flattenFirewallProfileProtocolOptionsSmtpOversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallProfileProtocolOptionsSmtpPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = flattenFirewallProfileProtocolOptionsSmtpProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = flattenFirewallProfileProtocolOptionsSmtpScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "server_busy"
	if _, ok := i["server-busy"]; ok {
		result["server_busy"] = flattenFirewallProfileProtocolOptionsSmtpServerBusy(i["server-busy"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_offloaded"
	if _, ok := i["ssl-offloaded"]; ok {
		result["ssl_offloaded"] = flattenFirewallProfileProtocolOptionsSmtpSslOffloaded(i["ssl-offloaded"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallProfileProtocolOptionsSmtpStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = flattenFirewallProfileProtocolOptionsSmtpUncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = flattenFirewallProfileProtocolOptionsSmtpUncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallProfileProtocolOptionsSmtpInspectAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSmtpOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileProtocolOptionsSmtpOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSmtpPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenFirewallProfileProtocolOptionsSmtpProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSmtpScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSmtpServerBusy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSmtpSslOffloaded(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSmtpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSmtpUncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSmtpUncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSsh(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "comfort_amount"
	if _, ok := i["comfort-amount"]; ok {
		result["comfort_amount"] = flattenFirewallProfileProtocolOptionsSshComfortAmount(i["comfort-amount"], d, pre_append)
	}

	pre_append = pre + ".0." + "comfort_interval"
	if _, ok := i["comfort-interval"]; ok {
		result["comfort_interval"] = flattenFirewallProfileProtocolOptionsSshComfortInterval(i["comfort-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenFirewallProfileProtocolOptionsSshOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = flattenFirewallProfileProtocolOptionsSshOversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = flattenFirewallProfileProtocolOptionsSshScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_offloaded"
	if _, ok := i["ssl-offloaded"]; ok {
		result["ssl_offloaded"] = flattenFirewallProfileProtocolOptionsSshSslOffloaded(i["ssl-offloaded"], d, pre_append)
	}

	pre_append = pre + ".0." + "stream_based_uncompressed_limit"
	if _, ok := i["stream-based-uncompressed-limit"]; ok {
		result["stream_based_uncompressed_limit"] = flattenFirewallProfileProtocolOptionsSshStreamBasedUncompressedLimit(i["stream-based-uncompressed-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_maximum"
	if _, ok := i["tcp-window-maximum"]; ok {
		result["tcp_window_maximum"] = flattenFirewallProfileProtocolOptionsSshTcpWindowMaximum(i["tcp-window-maximum"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_minimum"
	if _, ok := i["tcp-window-minimum"]; ok {
		result["tcp_window_minimum"] = flattenFirewallProfileProtocolOptionsSshTcpWindowMinimum(i["tcp-window-minimum"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_size"
	if _, ok := i["tcp-window-size"]; ok {
		result["tcp_window_size"] = flattenFirewallProfileProtocolOptionsSshTcpWindowSize(i["tcp-window-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_type"
	if _, ok := i["tcp-window-type"]; ok {
		result["tcp_window_type"] = flattenFirewallProfileProtocolOptionsSshTcpWindowType(i["tcp-window-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = flattenFirewallProfileProtocolOptionsSshUncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = flattenFirewallProfileProtocolOptionsSshUncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "explicit_ftp_tls"
	if _, ok := i["explicit-ftp-tls"]; ok {
		result["explicit_ftp_tls"] = flattenFirewallProfileProtocolOptionsSshExplicitFtpTls(i["explicit-ftp-tls"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallProfileProtocolOptionsSshComfortAmount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSshComfortInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSshOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileProtocolOptionsSshOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSshScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSshSslOffloaded(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSshStreamBasedUncompressedLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSshTcpWindowMaximum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSshTcpWindowMinimum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSshTcpWindowSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSshTcpWindowType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSshUncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSshUncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSshExplicitFtpTls(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSwitchingProtocolsLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsProxyRedirect(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallProfileProtocolOptionsProxyRedirectPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallProfileProtocolOptionsProxyRedirectStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallProfileProtocolOptionsProxyRedirectPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenFirewallProfileProtocolOptionsProxyRedirectStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallProfileProtocolOptions(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if isImportTable() {
		if err = d.Set("cifs", flattenFirewallProfileProtocolOptionsCifs(o["cifs"], d, "cifs")); err != nil {
			if vv, ok := fortiAPIPatch(o["cifs"], "FirewallProfileProtocolOptions-Cifs"); ok {
				if err = d.Set("cifs", vv); err != nil {
					return fmt.Errorf("Error reading cifs: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading cifs: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("cifs"); ok {
			if err = d.Set("cifs", flattenFirewallProfileProtocolOptionsCifs(o["cifs"], d, "cifs")); err != nil {
				if vv, ok := fortiAPIPatch(o["cifs"], "FirewallProfileProtocolOptions-Cifs"); ok {
					if err = d.Set("cifs", vv); err != nil {
						return fmt.Errorf("Error reading cifs: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading cifs: %v", err)
				}
			}
		}
	}

	if err = d.Set("comment", flattenFirewallProfileProtocolOptionsComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "FirewallProfileProtocolOptions-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dns", flattenFirewallProfileProtocolOptionsDns(o["dns"], d, "dns")); err != nil {
			if vv, ok := fortiAPIPatch(o["dns"], "FirewallProfileProtocolOptions-Dns"); ok {
				if err = d.Set("dns", vv); err != nil {
					return fmt.Errorf("Error reading dns: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dns: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dns"); ok {
			if err = d.Set("dns", flattenFirewallProfileProtocolOptionsDns(o["dns"], d, "dns")); err != nil {
				if vv, ok := fortiAPIPatch(o["dns"], "FirewallProfileProtocolOptions-Dns"); ok {
					if err = d.Set("dns", vv); err != nil {
						return fmt.Errorf("Error reading dns: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dns: %v", err)
				}
			}
		}
	}

	if err = d.Set("feature_set", flattenFirewallProfileProtocolOptionsFeatureSet(o["feature-set"], d, "feature_set")); err != nil {
		if vv, ok := fortiAPIPatch(o["feature-set"], "FirewallProfileProtocolOptions-FeatureSet"); ok {
			if err = d.Set("feature_set", vv); err != nil {
				return fmt.Errorf("Error reading feature_set: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading feature_set: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ftp", flattenFirewallProfileProtocolOptionsFtp(o["ftp"], d, "ftp")); err != nil {
			if vv, ok := fortiAPIPatch(o["ftp"], "FirewallProfileProtocolOptions-Ftp"); ok {
				if err = d.Set("ftp", vv); err != nil {
					return fmt.Errorf("Error reading ftp: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ftp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ftp"); ok {
			if err = d.Set("ftp", flattenFirewallProfileProtocolOptionsFtp(o["ftp"], d, "ftp")); err != nil {
				if vv, ok := fortiAPIPatch(o["ftp"], "FirewallProfileProtocolOptions-Ftp"); ok {
					if err = d.Set("ftp", vv); err != nil {
						return fmt.Errorf("Error reading ftp: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ftp: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("http", flattenFirewallProfileProtocolOptionsHttp(o["http"], d, "http")); err != nil {
			if vv, ok := fortiAPIPatch(o["http"], "FirewallProfileProtocolOptions-Http"); ok {
				if err = d.Set("http", vv); err != nil {
					return fmt.Errorf("Error reading http: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading http: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("http"); ok {
			if err = d.Set("http", flattenFirewallProfileProtocolOptionsHttp(o["http"], d, "http")); err != nil {
				if vv, ok := fortiAPIPatch(o["http"], "FirewallProfileProtocolOptions-Http"); ok {
					if err = d.Set("http", vv); err != nil {
						return fmt.Errorf("Error reading http: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading http: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("imap", flattenFirewallProfileProtocolOptionsImap(o["imap"], d, "imap")); err != nil {
			if vv, ok := fortiAPIPatch(o["imap"], "FirewallProfileProtocolOptions-Imap"); ok {
				if err = d.Set("imap", vv); err != nil {
					return fmt.Errorf("Error reading imap: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading imap: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("imap"); ok {
			if err = d.Set("imap", flattenFirewallProfileProtocolOptionsImap(o["imap"], d, "imap")); err != nil {
				if vv, ok := fortiAPIPatch(o["imap"], "FirewallProfileProtocolOptions-Imap"); ok {
					if err = d.Set("imap", vv); err != nil {
						return fmt.Errorf("Error reading imap: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading imap: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("mail_signature", flattenFirewallProfileProtocolOptionsMailSignature(o["mail-signature"], d, "mail_signature")); err != nil {
			if vv, ok := fortiAPIPatch(o["mail-signature"], "FirewallProfileProtocolOptions-MailSignature"); ok {
				if err = d.Set("mail_signature", vv); err != nil {
					return fmt.Errorf("Error reading mail_signature: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mail_signature: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mail_signature"); ok {
			if err = d.Set("mail_signature", flattenFirewallProfileProtocolOptionsMailSignature(o["mail-signature"], d, "mail_signature")); err != nil {
				if vv, ok := fortiAPIPatch(o["mail-signature"], "FirewallProfileProtocolOptions-MailSignature"); ok {
					if err = d.Set("mail_signature", vv); err != nil {
						return fmt.Errorf("Error reading mail_signature: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mail_signature: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("mapi", flattenFirewallProfileProtocolOptionsMapi(o["mapi"], d, "mapi")); err != nil {
			if vv, ok := fortiAPIPatch(o["mapi"], "FirewallProfileProtocolOptions-Mapi"); ok {
				if err = d.Set("mapi", vv); err != nil {
					return fmt.Errorf("Error reading mapi: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mapi: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mapi"); ok {
			if err = d.Set("mapi", flattenFirewallProfileProtocolOptionsMapi(o["mapi"], d, "mapi")); err != nil {
				if vv, ok := fortiAPIPatch(o["mapi"], "FirewallProfileProtocolOptions-Mapi"); ok {
					if err = d.Set("mapi", vv); err != nil {
						return fmt.Errorf("Error reading mapi: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mapi: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenFirewallProfileProtocolOptionsName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallProfileProtocolOptions-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("nntp", flattenFirewallProfileProtocolOptionsNntp(o["nntp"], d, "nntp")); err != nil {
			if vv, ok := fortiAPIPatch(o["nntp"], "FirewallProfileProtocolOptions-Nntp"); ok {
				if err = d.Set("nntp", vv); err != nil {
					return fmt.Errorf("Error reading nntp: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading nntp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("nntp"); ok {
			if err = d.Set("nntp", flattenFirewallProfileProtocolOptionsNntp(o["nntp"], d, "nntp")); err != nil {
				if vv, ok := fortiAPIPatch(o["nntp"], "FirewallProfileProtocolOptions-Nntp"); ok {
					if err = d.Set("nntp", vv); err != nil {
						return fmt.Errorf("Error reading nntp: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading nntp: %v", err)
				}
			}
		}
	}

	if err = d.Set("oversize_log", flattenFirewallProfileProtocolOptionsOversizeLog(o["oversize-log"], d, "oversize_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["oversize-log"], "FirewallProfileProtocolOptions-OversizeLog"); ok {
			if err = d.Set("oversize_log", vv); err != nil {
				return fmt.Errorf("Error reading oversize_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oversize_log: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("pop3", flattenFirewallProfileProtocolOptionsPop3(o["pop3"], d, "pop3")); err != nil {
			if vv, ok := fortiAPIPatch(o["pop3"], "FirewallProfileProtocolOptions-Pop3"); ok {
				if err = d.Set("pop3", vv); err != nil {
					return fmt.Errorf("Error reading pop3: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading pop3: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("pop3"); ok {
			if err = d.Set("pop3", flattenFirewallProfileProtocolOptionsPop3(o["pop3"], d, "pop3")); err != nil {
				if vv, ok := fortiAPIPatch(o["pop3"], "FirewallProfileProtocolOptions-Pop3"); ok {
					if err = d.Set("pop3", vv); err != nil {
						return fmt.Errorf("Error reading pop3: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading pop3: %v", err)
				}
			}
		}
	}

	if err = d.Set("replacemsg_group", flattenFirewallProfileProtocolOptionsReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["replacemsg-group"], "FirewallProfileProtocolOptions-ReplacemsgGroup"); ok {
			if err = d.Set("replacemsg_group", vv); err != nil {
				return fmt.Errorf("Error reading replacemsg_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if err = d.Set("rpc_over_http", flattenFirewallProfileProtocolOptionsRpcOverHttp(o["rpc-over-http"], d, "rpc_over_http")); err != nil {
		if vv, ok := fortiAPIPatch(o["rpc-over-http"], "FirewallProfileProtocolOptions-RpcOverHttp"); ok {
			if err = d.Set("rpc_over_http", vv); err != nil {
				return fmt.Errorf("Error reading rpc_over_http: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rpc_over_http: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("smtp", flattenFirewallProfileProtocolOptionsSmtp(o["smtp"], d, "smtp")); err != nil {
			if vv, ok := fortiAPIPatch(o["smtp"], "FirewallProfileProtocolOptions-Smtp"); ok {
				if err = d.Set("smtp", vv); err != nil {
					return fmt.Errorf("Error reading smtp: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading smtp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("smtp"); ok {
			if err = d.Set("smtp", flattenFirewallProfileProtocolOptionsSmtp(o["smtp"], d, "smtp")); err != nil {
				if vv, ok := fortiAPIPatch(o["smtp"], "FirewallProfileProtocolOptions-Smtp"); ok {
					if err = d.Set("smtp", vv); err != nil {
						return fmt.Errorf("Error reading smtp: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading smtp: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ssh", flattenFirewallProfileProtocolOptionsSsh(o["ssh"], d, "ssh")); err != nil {
			if vv, ok := fortiAPIPatch(o["ssh"], "FirewallProfileProtocolOptions-Ssh"); ok {
				if err = d.Set("ssh", vv); err != nil {
					return fmt.Errorf("Error reading ssh: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ssh: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssh"); ok {
			if err = d.Set("ssh", flattenFirewallProfileProtocolOptionsSsh(o["ssh"], d, "ssh")); err != nil {
				if vv, ok := fortiAPIPatch(o["ssh"], "FirewallProfileProtocolOptions-Ssh"); ok {
					if err = d.Set("ssh", vv); err != nil {
						return fmt.Errorf("Error reading ssh: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ssh: %v", err)
				}
			}
		}
	}

	if err = d.Set("switching_protocols_log", flattenFirewallProfileProtocolOptionsSwitchingProtocolsLog(o["switching-protocols-log"], d, "switching_protocols_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["switching-protocols-log"], "FirewallProfileProtocolOptions-SwitchingProtocolsLog"); ok {
			if err = d.Set("switching_protocols_log", vv); err != nil {
				return fmt.Errorf("Error reading switching_protocols_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switching_protocols_log: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("proxy_redirect", flattenFirewallProfileProtocolOptionsProxyRedirect(o["proxy-redirect"], d, "proxy_redirect")); err != nil {
			if vv, ok := fortiAPIPatch(o["proxy-redirect"], "FirewallProfileProtocolOptions-ProxyRedirect"); ok {
				if err = d.Set("proxy_redirect", vv); err != nil {
					return fmt.Errorf("Error reading proxy_redirect: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading proxy_redirect: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("proxy_redirect"); ok {
			if err = d.Set("proxy_redirect", flattenFirewallProfileProtocolOptionsProxyRedirect(o["proxy-redirect"], d, "proxy_redirect")); err != nil {
				if vv, ok := fortiAPIPatch(o["proxy-redirect"], "FirewallProfileProtocolOptions-ProxyRedirect"); ok {
					if err = d.Set("proxy_redirect", vv); err != nil {
						return fmt.Errorf("Error reading proxy_redirect: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading proxy_redirect: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenFirewallProfileProtocolOptionsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallProfileProtocolOptionsCifs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "domain_controller"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["domain-controller"], _ = expandFirewallProfileProtocolOptionsCifsDomainController(d, i["domain_controller"], pre_append)
	}
	pre_append = pre + ".0." + "file_filter"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandFirewallProfileProtocolOptionsCifsFileFilter(d, i["file_filter"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["file-filter"] = t
		}
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandFirewallProfileProtocolOptionsCifsOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["oversize-limit"], _ = expandFirewallProfileProtocolOptionsCifsOversizeLimit(d, i["oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ports"], _ = expandFirewallProfileProtocolOptionsCifsPorts(d, i["ports"], pre_append)
	}
	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["scan-bzip2"], _ = expandFirewallProfileProtocolOptionsCifsScanBzip2(d, i["scan_bzip2"], pre_append)
	}
	pre_append = pre + ".0." + "server_credential_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["server-credential-type"], _ = expandFirewallProfileProtocolOptionsCifsServerCredentialType(d, i["server_credential_type"], pre_append)
	}
	pre_append = pre + ".0." + "server_keytab"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandFirewallProfileProtocolOptionsCifsServerKeytab(d, i["server_keytab"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["server-keytab"] = t
		}
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFirewallProfileProtocolOptionsCifsStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_maximum"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-window-maximum"], _ = expandFirewallProfileProtocolOptionsCifsTcpWindowMaximum(d, i["tcp_window_maximum"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_minimum"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-window-minimum"], _ = expandFirewallProfileProtocolOptionsCifsTcpWindowMinimum(d, i["tcp_window_minimum"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-window-size"], _ = expandFirewallProfileProtocolOptionsCifsTcpWindowSize(d, i["tcp_window_size"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-window-type"], _ = expandFirewallProfileProtocolOptionsCifsTcpWindowType(d, i["tcp_window_type"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["uncompressed-nest-limit"], _ = expandFirewallProfileProtocolOptionsCifsUncompressedNestLimit(d, i["uncompressed_nest_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["uncompressed-oversize-limit"], _ = expandFirewallProfileProtocolOptionsCifsUncompressedOversizeLimit(d, i["uncompressed_oversize_limit"], pre_append)
	}

	return result, nil
}

func expandFirewallProfileProtocolOptionsCifsDomainController(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsCifsFileFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "entries"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandFirewallProfileProtocolOptionsCifsFileFilterEntries(d, i["entries"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["entries"] = t
		}
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandFirewallProfileProtocolOptionsCifsFileFilterLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFirewallProfileProtocolOptionsCifsFileFilterStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandFirewallProfileProtocolOptionsCifsFileFilterEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandFirewallProfileProtocolOptionsCifsFileFilterEntriesAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandFirewallProfileProtocolOptionsCifsFileFilterEntriesComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["direction"], _ = expandFirewallProfileProtocolOptionsCifsFileFilterEntriesDirection(d, i["direction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["file-type"], _ = expandFirewallProfileProtocolOptionsCifsFileFilterEntriesFileType(d, i["file_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter"], _ = expandFirewallProfileProtocolOptionsCifsFileFilterEntriesFilter(d, i["filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandFirewallProfileProtocolOptionsCifsFileFilterEntriesProtocol(d, i["protocol"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallProfileProtocolOptionsCifsFileFilterEntriesAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsFileFilterEntriesComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsFileFilterEntriesDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsFileFilterEntriesFileType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsCifsFileFilterEntriesFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsFileFilterEntriesProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsCifsFileFilterLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsFileFilterStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsCifsOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsCifsScanBzip2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsServerCredentialType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsServerKeytab(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keytab"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["keytab"], _ = expandFirewallProfileProtocolOptionsCifsServerKeytabKeytab(d, i["keytab"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password"], _ = expandFirewallProfileProtocolOptionsCifsServerKeytabPassword(d, i["password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "principal"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["principal"], _ = expandFirewallProfileProtocolOptionsCifsServerKeytabPrincipal(d, i["principal"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallProfileProtocolOptionsCifsServerKeytabKeytab(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsServerKeytabPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsCifsServerKeytabPrincipal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsTcpWindowMaximum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsTcpWindowMinimum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsTcpWindowSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsTcpWindowType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsUncompressedNestLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsUncompressedOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsDns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ports"], _ = expandFirewallProfileProtocolOptionsDnsPorts(d, i["ports"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFirewallProfileProtocolOptionsDnsStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandFirewallProfileProtocolOptionsDnsPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsDnsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsFeatureSet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsFtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "comfort_amount"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["comfort-amount"], _ = expandFirewallProfileProtocolOptionsFtpComfortAmount(d, i["comfort_amount"], pre_append)
	}
	pre_append = pre + ".0." + "comfort_interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["comfort-interval"], _ = expandFirewallProfileProtocolOptionsFtpComfortInterval(d, i["comfort_interval"], pre_append)
	}
	pre_append = pre + ".0." + "explicit_ftp_tls"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["explicit-ftp-tls"], _ = expandFirewallProfileProtocolOptionsFtpExplicitFtpTls(d, i["explicit_ftp_tls"], pre_append)
	}
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["inspect-all"], _ = expandFirewallProfileProtocolOptionsFtpInspectAll(d, i["inspect_all"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandFirewallProfileProtocolOptionsFtpOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["oversize-limit"], _ = expandFirewallProfileProtocolOptionsFtpOversizeLimit(d, i["oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ports"], _ = expandFirewallProfileProtocolOptionsFtpPorts(d, i["ports"], pre_append)
	}
	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["scan-bzip2"], _ = expandFirewallProfileProtocolOptionsFtpScanBzip2(d, i["scan_bzip2"], pre_append)
	}
	pre_append = pre + ".0." + "ssl_offloaded"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ssl-offloaded"], _ = expandFirewallProfileProtocolOptionsFtpSslOffloaded(d, i["ssl_offloaded"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFirewallProfileProtocolOptionsFtpStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "stream_based_uncompressed_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["stream-based-uncompressed-limit"], _ = expandFirewallProfileProtocolOptionsFtpStreamBasedUncompressedLimit(d, i["stream_based_uncompressed_limit"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_maximum"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-window-maximum"], _ = expandFirewallProfileProtocolOptionsFtpTcpWindowMaximum(d, i["tcp_window_maximum"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_minimum"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-window-minimum"], _ = expandFirewallProfileProtocolOptionsFtpTcpWindowMinimum(d, i["tcp_window_minimum"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-window-size"], _ = expandFirewallProfileProtocolOptionsFtpTcpWindowSize(d, i["tcp_window_size"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-window-type"], _ = expandFirewallProfileProtocolOptionsFtpTcpWindowType(d, i["tcp_window_type"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["uncompressed-nest-limit"], _ = expandFirewallProfileProtocolOptionsFtpUncompressedNestLimit(d, i["uncompressed_nest_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["uncompressed-oversize-limit"], _ = expandFirewallProfileProtocolOptionsFtpUncompressedOversizeLimit(d, i["uncompressed_oversize_limit"], pre_append)
	}

	return result, nil
}

func expandFirewallProfileProtocolOptionsFtpComfortAmount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsFtpComfortInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsFtpExplicitFtpTls(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsFtpInspectAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsFtpOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsFtpOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsFtpPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsFtpScanBzip2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsFtpSslOffloaded(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsFtpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsFtpStreamBasedUncompressedLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsFtpTcpWindowMaximum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsFtpTcpWindowMinimum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsFtpTcpWindowSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsFtpTcpWindowType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsFtpUncompressedNestLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsFtpUncompressedOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "address_ip_rating"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["address-ip-rating"], _ = expandFirewallProfileProtocolOptionsHttpAddressIpRating(d, i["address_ip_rating"], pre_append)
	}
	pre_append = pre + ".0." + "block_page_status_code"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["block-page-status-code"], _ = expandFirewallProfileProtocolOptionsHttpBlockPageStatusCode(d, i["block_page_status_code"], pre_append)
	}
	pre_append = pre + ".0." + "comfort_amount"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["comfort-amount"], _ = expandFirewallProfileProtocolOptionsHttpComfortAmount(d, i["comfort_amount"], pre_append)
	}
	pre_append = pre + ".0." + "comfort_interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["comfort-interval"], _ = expandFirewallProfileProtocolOptionsHttpComfortInterval(d, i["comfort_interval"], pre_append)
	}
	pre_append = pre + ".0." + "fortinet_bar"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortinet-bar"], _ = expandFirewallProfileProtocolOptionsHttpFortinetBar(d, i["fortinet_bar"], pre_append)
	}
	pre_append = pre + ".0." + "fortinet_bar_port"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortinet-bar-port"], _ = expandFirewallProfileProtocolOptionsHttpFortinetBarPort(d, i["fortinet_bar_port"], pre_append)
	}
	pre_append = pre + ".0." + "domain_fronting"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["domain-fronting"], _ = expandFirewallProfileProtocolOptionsHttpDomainFronting(d, i["domain_fronting"], pre_append)
	}
	pre_append = pre + ".0." + "h2c"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["h2c"], _ = expandFirewallProfileProtocolOptionsHttpH2C(d, i["h2c"], pre_append)
	}
	pre_append = pre + ".0." + "http_09"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["http-0.9"], _ = expandFirewallProfileProtocolOptionsHttpHttp09(d, i["http_09"], pre_append)
	}
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["inspect-all"], _ = expandFirewallProfileProtocolOptionsHttpInspectAll(d, i["inspect_all"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandFirewallProfileProtocolOptionsHttpOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["oversize-limit"], _ = expandFirewallProfileProtocolOptionsHttpOversizeLimit(d, i["oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ports"], _ = expandFirewallProfileProtocolOptionsHttpPorts(d, i["ports"], pre_append)
	}
	pre_append = pre + ".0." + "post_lang"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["post-lang"], _ = expandFirewallProfileProtocolOptionsHttpPostLang(d, i["post_lang"], pre_append)
	}
	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["proxy-after-tcp-handshake"], _ = expandFirewallProfileProtocolOptionsHttpProxyAfterTcpHandshake(d, i["proxy_after_tcp_handshake"], pre_append)
	}
	pre_append = pre + ".0." + "range_block"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["range-block"], _ = expandFirewallProfileProtocolOptionsHttpRangeBlock(d, i["range_block"], pre_append)
	}
	pre_append = pre + ".0." + "retry_count"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["retry-count"], _ = expandFirewallProfileProtocolOptionsHttpRetryCount(d, i["retry_count"], pre_append)
	}
	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["scan-bzip2"], _ = expandFirewallProfileProtocolOptionsHttpScanBzip2(d, i["scan_bzip2"], pre_append)
	}
	pre_append = pre + ".0." + "ssl_offloaded"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ssl-offloaded"], _ = expandFirewallProfileProtocolOptionsHttpSslOffloaded(d, i["ssl_offloaded"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFirewallProfileProtocolOptionsHttpStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "stream_based_uncompressed_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["stream-based-uncompressed-limit"], _ = expandFirewallProfileProtocolOptionsHttpStreamBasedUncompressedLimit(d, i["stream_based_uncompressed_limit"], pre_append)
	}
	pre_append = pre + ".0." + "streaming_content_bypass"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["streaming-content-bypass"], _ = expandFirewallProfileProtocolOptionsHttpStreamingContentBypass(d, i["streaming_content_bypass"], pre_append)
	}
	pre_append = pre + ".0." + "strip_x_forwarded_for"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["strip-x-forwarded-for"], _ = expandFirewallProfileProtocolOptionsHttpStripXForwardedFor(d, i["strip_x_forwarded_for"], pre_append)
	}
	pre_append = pre + ".0." + "switching_protocols"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switching-protocols"], _ = expandFirewallProfileProtocolOptionsHttpSwitchingProtocols(d, i["switching_protocols"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_maximum"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-window-maximum"], _ = expandFirewallProfileProtocolOptionsHttpTcpWindowMaximum(d, i["tcp_window_maximum"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_minimum"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-window-minimum"], _ = expandFirewallProfileProtocolOptionsHttpTcpWindowMinimum(d, i["tcp_window_minimum"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-window-size"], _ = expandFirewallProfileProtocolOptionsHttpTcpWindowSize(d, i["tcp_window_size"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-window-type"], _ = expandFirewallProfileProtocolOptionsHttpTcpWindowType(d, i["tcp_window_type"], pre_append)
	}
	pre_append = pre + ".0." + "tunnel_non_http"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tunnel-non-http"], _ = expandFirewallProfileProtocolOptionsHttpTunnelNonHttp(d, i["tunnel_non_http"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["uncompressed-nest-limit"], _ = expandFirewallProfileProtocolOptionsHttpUncompressedNestLimit(d, i["uncompressed_nest_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["uncompressed-oversize-limit"], _ = expandFirewallProfileProtocolOptionsHttpUncompressedOversizeLimit(d, i["uncompressed_oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "unknown_content_encoding"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unknown-content-encoding"], _ = expandFirewallProfileProtocolOptionsHttpUnknownContentEncoding(d, i["unknown_content_encoding"], pre_append)
	}
	pre_append = pre + ".0." + "unknown_http_version"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unknown-http-version"], _ = expandFirewallProfileProtocolOptionsHttpUnknownHttpVersion(d, i["unknown_http_version"], pre_append)
	}
	pre_append = pre + ".0." + "verify_dns_for_policy_matching"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["verify-dns-for-policy-matching"], _ = expandFirewallProfileProtocolOptionsHttpVerifyDnsForPolicyMatching(d, i["verify_dns_for_policy_matching"], pre_append)
	}
	pre_append = pre + ".0." + "dns_protection"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dns-protection"], _ = expandFirewallProfileProtocolOptionsHttpDnsProtection(d, i["dns_protection"], pre_append)
	}
	pre_append = pre + ".0." + "encrypted_file"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["encrypted-file"], _ = expandFirewallProfileProtocolOptionsHttpEncryptedFile(d, i["encrypted_file"], pre_append)
	}
	pre_append = pre + ".0." + "encrypted_file_log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["encrypted-file-log"], _ = expandFirewallProfileProtocolOptionsHttpEncryptedFileLog(d, i["encrypted_file_log"], pre_append)
	}

	return result, nil
}

func expandFirewallProfileProtocolOptionsHttpAddressIpRating(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpBlockPageStatusCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpComfortAmount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpComfortInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpFortinetBar(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpFortinetBarPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpDomainFronting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpH2C(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpHttp09(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpInspectAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsHttpOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsHttpPostLang(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsHttpProxyAfterTcpHandshake(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpRangeBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpRetryCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpScanBzip2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpSslOffloaded(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpStreamBasedUncompressedLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpStreamingContentBypass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpStripXForwardedFor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpSwitchingProtocols(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpTcpWindowMaximum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpTcpWindowMinimum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpTcpWindowSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpTcpWindowType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpTunnelNonHttp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpUncompressedNestLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpUncompressedOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpUnknownContentEncoding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpUnknownHttpVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpVerifyDnsForPolicyMatching(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpDnsProtection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpEncryptedFile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpEncryptedFileLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsImap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["inspect-all"], _ = expandFirewallProfileProtocolOptionsImapInspectAll(d, i["inspect_all"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandFirewallProfileProtocolOptionsImapOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["oversize-limit"], _ = expandFirewallProfileProtocolOptionsImapOversizeLimit(d, i["oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ports"], _ = expandFirewallProfileProtocolOptionsImapPorts(d, i["ports"], pre_append)
	}
	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["proxy-after-tcp-handshake"], _ = expandFirewallProfileProtocolOptionsImapProxyAfterTcpHandshake(d, i["proxy_after_tcp_handshake"], pre_append)
	}
	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["scan-bzip2"], _ = expandFirewallProfileProtocolOptionsImapScanBzip2(d, i["scan_bzip2"], pre_append)
	}
	pre_append = pre + ".0." + "ssl_offloaded"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ssl-offloaded"], _ = expandFirewallProfileProtocolOptionsImapSslOffloaded(d, i["ssl_offloaded"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFirewallProfileProtocolOptionsImapStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["uncompressed-nest-limit"], _ = expandFirewallProfileProtocolOptionsImapUncompressedNestLimit(d, i["uncompressed_nest_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["uncompressed-oversize-limit"], _ = expandFirewallProfileProtocolOptionsImapUncompressedOversizeLimit(d, i["uncompressed_oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "address_ip_rating"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["address-ip-rating"], _ = expandFirewallProfileProtocolOptionsImapAddressIpRating(d, i["address_ip_rating"], pre_append)
	}

	return result, nil
}

func expandFirewallProfileProtocolOptionsImapInspectAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsImapOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsImapOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsImapPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsImapProxyAfterTcpHandshake(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsImapScanBzip2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsImapSslOffloaded(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsImapStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsImapUncompressedNestLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsImapUncompressedOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsImapAddressIpRating(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsMailSignature(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "signature"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["signature"], _ = expandFirewallProfileProtocolOptionsMailSignatureSignature(d, i["signature"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFirewallProfileProtocolOptionsMailSignatureStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandFirewallProfileProtocolOptionsMailSignatureSignature(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsMailSignatureStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsMapi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandFirewallProfileProtocolOptionsMapiOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["oversize-limit"], _ = expandFirewallProfileProtocolOptionsMapiOversizeLimit(d, i["oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ports"], _ = expandFirewallProfileProtocolOptionsMapiPorts(d, i["ports"], pre_append)
	}
	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["scan-bzip2"], _ = expandFirewallProfileProtocolOptionsMapiScanBzip2(d, i["scan_bzip2"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFirewallProfileProtocolOptionsMapiStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["uncompressed-nest-limit"], _ = expandFirewallProfileProtocolOptionsMapiUncompressedNestLimit(d, i["uncompressed_nest_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["uncompressed-oversize-limit"], _ = expandFirewallProfileProtocolOptionsMapiUncompressedOversizeLimit(d, i["uncompressed_oversize_limit"], pre_append)
	}

	return result, nil
}

func expandFirewallProfileProtocolOptionsMapiOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsMapiOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsMapiPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsMapiScanBzip2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsMapiStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsMapiUncompressedNestLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsMapiUncompressedOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsNntp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["inspect-all"], _ = expandFirewallProfileProtocolOptionsNntpInspectAll(d, i["inspect_all"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandFirewallProfileProtocolOptionsNntpOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["oversize-limit"], _ = expandFirewallProfileProtocolOptionsNntpOversizeLimit(d, i["oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ports"], _ = expandFirewallProfileProtocolOptionsNntpPorts(d, i["ports"], pre_append)
	}
	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["proxy-after-tcp-handshake"], _ = expandFirewallProfileProtocolOptionsNntpProxyAfterTcpHandshake(d, i["proxy_after_tcp_handshake"], pre_append)
	}
	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["scan-bzip2"], _ = expandFirewallProfileProtocolOptionsNntpScanBzip2(d, i["scan_bzip2"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFirewallProfileProtocolOptionsNntpStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["uncompressed-nest-limit"], _ = expandFirewallProfileProtocolOptionsNntpUncompressedNestLimit(d, i["uncompressed_nest_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["uncompressed-oversize-limit"], _ = expandFirewallProfileProtocolOptionsNntpUncompressedOversizeLimit(d, i["uncompressed_oversize_limit"], pre_append)
	}

	return result, nil
}

func expandFirewallProfileProtocolOptionsNntpInspectAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsNntpOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsNntpOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsNntpPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsNntpProxyAfterTcpHandshake(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsNntpScanBzip2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsNntpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsNntpUncompressedNestLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsNntpUncompressedOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsOversizeLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsPop3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["inspect-all"], _ = expandFirewallProfileProtocolOptionsPop3InspectAll(d, i["inspect_all"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandFirewallProfileProtocolOptionsPop3Options(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["oversize-limit"], _ = expandFirewallProfileProtocolOptionsPop3OversizeLimit(d, i["oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ports"], _ = expandFirewallProfileProtocolOptionsPop3Ports(d, i["ports"], pre_append)
	}
	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["proxy-after-tcp-handshake"], _ = expandFirewallProfileProtocolOptionsPop3ProxyAfterTcpHandshake(d, i["proxy_after_tcp_handshake"], pre_append)
	}
	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["scan-bzip2"], _ = expandFirewallProfileProtocolOptionsPop3ScanBzip2(d, i["scan_bzip2"], pre_append)
	}
	pre_append = pre + ".0." + "ssl_offloaded"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ssl-offloaded"], _ = expandFirewallProfileProtocolOptionsPop3SslOffloaded(d, i["ssl_offloaded"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFirewallProfileProtocolOptionsPop3Status(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["uncompressed-nest-limit"], _ = expandFirewallProfileProtocolOptionsPop3UncompressedNestLimit(d, i["uncompressed_nest_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["uncompressed-oversize-limit"], _ = expandFirewallProfileProtocolOptionsPop3UncompressedOversizeLimit(d, i["uncompressed_oversize_limit"], pre_append)
	}

	return result, nil
}

func expandFirewallProfileProtocolOptionsPop3InspectAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsPop3Options(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsPop3OversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsPop3Ports(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsPop3ProxyAfterTcpHandshake(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsPop3ScanBzip2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsPop3SslOffloaded(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsPop3Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsPop3UncompressedNestLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsPop3UncompressedOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsRpcOverHttp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSmtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["inspect-all"], _ = expandFirewallProfileProtocolOptionsSmtpInspectAll(d, i["inspect_all"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandFirewallProfileProtocolOptionsSmtpOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["oversize-limit"], _ = expandFirewallProfileProtocolOptionsSmtpOversizeLimit(d, i["oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ports"], _ = expandFirewallProfileProtocolOptionsSmtpPorts(d, i["ports"], pre_append)
	}
	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["proxy-after-tcp-handshake"], _ = expandFirewallProfileProtocolOptionsSmtpProxyAfterTcpHandshake(d, i["proxy_after_tcp_handshake"], pre_append)
	}
	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["scan-bzip2"], _ = expandFirewallProfileProtocolOptionsSmtpScanBzip2(d, i["scan_bzip2"], pre_append)
	}
	pre_append = pre + ".0." + "server_busy"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["server-busy"], _ = expandFirewallProfileProtocolOptionsSmtpServerBusy(d, i["server_busy"], pre_append)
	}
	pre_append = pre + ".0." + "ssl_offloaded"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ssl-offloaded"], _ = expandFirewallProfileProtocolOptionsSmtpSslOffloaded(d, i["ssl_offloaded"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFirewallProfileProtocolOptionsSmtpStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["uncompressed-nest-limit"], _ = expandFirewallProfileProtocolOptionsSmtpUncompressedNestLimit(d, i["uncompressed_nest_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["uncompressed-oversize-limit"], _ = expandFirewallProfileProtocolOptionsSmtpUncompressedOversizeLimit(d, i["uncompressed_oversize_limit"], pre_append)
	}

	return result, nil
}

func expandFirewallProfileProtocolOptionsSmtpInspectAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSmtpOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsSmtpOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSmtpPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsSmtpProxyAfterTcpHandshake(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSmtpScanBzip2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSmtpServerBusy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSmtpSslOffloaded(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSmtpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSmtpUncompressedNestLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSmtpUncompressedOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSsh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "comfort_amount"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["comfort-amount"], _ = expandFirewallProfileProtocolOptionsSshComfortAmount(d, i["comfort_amount"], pre_append)
	}
	pre_append = pre + ".0." + "comfort_interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["comfort-interval"], _ = expandFirewallProfileProtocolOptionsSshComfortInterval(d, i["comfort_interval"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandFirewallProfileProtocolOptionsSshOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["oversize-limit"], _ = expandFirewallProfileProtocolOptionsSshOversizeLimit(d, i["oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["scan-bzip2"], _ = expandFirewallProfileProtocolOptionsSshScanBzip2(d, i["scan_bzip2"], pre_append)
	}
	pre_append = pre + ".0." + "ssl_offloaded"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ssl-offloaded"], _ = expandFirewallProfileProtocolOptionsSshSslOffloaded(d, i["ssl_offloaded"], pre_append)
	}
	pre_append = pre + ".0." + "stream_based_uncompressed_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["stream-based-uncompressed-limit"], _ = expandFirewallProfileProtocolOptionsSshStreamBasedUncompressedLimit(d, i["stream_based_uncompressed_limit"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_maximum"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-window-maximum"], _ = expandFirewallProfileProtocolOptionsSshTcpWindowMaximum(d, i["tcp_window_maximum"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_minimum"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-window-minimum"], _ = expandFirewallProfileProtocolOptionsSshTcpWindowMinimum(d, i["tcp_window_minimum"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-window-size"], _ = expandFirewallProfileProtocolOptionsSshTcpWindowSize(d, i["tcp_window_size"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-window-type"], _ = expandFirewallProfileProtocolOptionsSshTcpWindowType(d, i["tcp_window_type"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["uncompressed-nest-limit"], _ = expandFirewallProfileProtocolOptionsSshUncompressedNestLimit(d, i["uncompressed_nest_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["uncompressed-oversize-limit"], _ = expandFirewallProfileProtocolOptionsSshUncompressedOversizeLimit(d, i["uncompressed_oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "explicit_ftp_tls"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["explicit-ftp-tls"], _ = expandFirewallProfileProtocolOptionsSshExplicitFtpTls(d, i["explicit_ftp_tls"], pre_append)
	}

	return result, nil
}

func expandFirewallProfileProtocolOptionsSshComfortAmount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSshComfortInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSshOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsSshOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSshScanBzip2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSshSslOffloaded(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSshStreamBasedUncompressedLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSshTcpWindowMaximum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSshTcpWindowMinimum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSshTcpWindowSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSshTcpWindowType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSshUncompressedNestLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSshUncompressedOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSshExplicitFtpTls(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSwitchingProtocolsLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsProxyRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ports"], _ = expandFirewallProfileProtocolOptionsProxyRedirectPorts(d, i["ports"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFirewallProfileProtocolOptionsProxyRedirectStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandFirewallProfileProtocolOptionsProxyRedirectPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsProxyRedirectStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallProfileProtocolOptions(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cifs"); ok || d.HasChange("cifs") {
		t, err := expandFirewallProfileProtocolOptionsCifs(d, v, "cifs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cifs"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandFirewallProfileProtocolOptionsComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("dns"); ok || d.HasChange("dns") {
		t, err := expandFirewallProfileProtocolOptionsDns(d, v, "dns")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns"] = t
		}
	}

	if v, ok := d.GetOk("feature_set"); ok || d.HasChange("feature_set") {
		t, err := expandFirewallProfileProtocolOptionsFeatureSet(d, v, "feature_set")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["feature-set"] = t
		}
	}

	if v, ok := d.GetOk("ftp"); ok || d.HasChange("ftp") {
		t, err := expandFirewallProfileProtocolOptionsFtp(d, v, "ftp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp"] = t
		}
	}

	if v, ok := d.GetOk("http"); ok || d.HasChange("http") {
		t, err := expandFirewallProfileProtocolOptionsHttp(d, v, "http")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http"] = t
		}
	}

	if v, ok := d.GetOk("imap"); ok || d.HasChange("imap") {
		t, err := expandFirewallProfileProtocolOptionsImap(d, v, "imap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["imap"] = t
		}
	}

	if v, ok := d.GetOk("mail_signature"); ok || d.HasChange("mail_signature") {
		t, err := expandFirewallProfileProtocolOptionsMailSignature(d, v, "mail_signature")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mail-signature"] = t
		}
	}

	if v, ok := d.GetOk("mapi"); ok || d.HasChange("mapi") {
		t, err := expandFirewallProfileProtocolOptionsMapi(d, v, "mapi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mapi"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallProfileProtocolOptionsName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nntp"); ok || d.HasChange("nntp") {
		t, err := expandFirewallProfileProtocolOptionsNntp(d, v, "nntp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nntp"] = t
		}
	}

	if v, ok := d.GetOk("oversize_log"); ok || d.HasChange("oversize_log") {
		t, err := expandFirewallProfileProtocolOptionsOversizeLog(d, v, "oversize_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oversize-log"] = t
		}
	}

	if v, ok := d.GetOk("pop3"); ok || d.HasChange("pop3") {
		t, err := expandFirewallProfileProtocolOptionsPop3(d, v, "pop3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pop3"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok || d.HasChange("replacemsg_group") {
		t, err := expandFirewallProfileProtocolOptionsReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("rpc_over_http"); ok || d.HasChange("rpc_over_http") {
		t, err := expandFirewallProfileProtocolOptionsRpcOverHttp(d, v, "rpc_over_http")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rpc-over-http"] = t
		}
	}

	if v, ok := d.GetOk("smtp"); ok || d.HasChange("smtp") {
		t, err := expandFirewallProfileProtocolOptionsSmtp(d, v, "smtp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smtp"] = t
		}
	}

	if v, ok := d.GetOk("ssh"); ok || d.HasChange("ssh") {
		t, err := expandFirewallProfileProtocolOptionsSsh(d, v, "ssh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh"] = t
		}
	}

	if v, ok := d.GetOk("switching_protocols_log"); ok || d.HasChange("switching_protocols_log") {
		t, err := expandFirewallProfileProtocolOptionsSwitchingProtocolsLog(d, v, "switching_protocols_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switching-protocols-log"] = t
		}
	}

	if v, ok := d.GetOk("proxy_redirect"); ok || d.HasChange("proxy_redirect") {
		t, err := expandFirewallProfileProtocolOptionsProxyRedirect(d, v, "proxy_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-redirect"] = t
		}
	}

	return &obj, nil
}
