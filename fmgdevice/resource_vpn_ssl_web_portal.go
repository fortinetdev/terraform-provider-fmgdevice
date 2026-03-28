// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Portal.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnSslWebPortal() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnSslWebPortalCreate,
		Read:   resourceVpnSslWebPortalRead,
		Update: resourceVpnSslWebPortalUpdate,
		Delete: resourceVpnSslWebPortalDelete,

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
			"allow_user_access": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"auto_connect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"bookmark_group": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bookmarks": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_params": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"apptype": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"color_depth": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"description": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"domain": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"folder": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"form_data": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"height": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"host": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"keyboard_layout": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"listening_port": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"load_balancing_info": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"logon_password": &schema.Schema{
										Type:      schema.TypeSet,
										Elem:      &schema.Schema{Type: schema.TypeString},
										Optional:  true,
										Sensitive: true,
										Computed:  true,
									},
									"logon_user": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"port": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"preconnection_blob": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"preconnection_id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"remote_port": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"restricted_admin": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"security": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"send_preconnection_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"server_layout": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"show_status_window": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"sso": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"sso_credential": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"sso_credential_sent_once": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"sso_password": &schema.Schema{
										Type:      schema.TypeSet,
										Elem:      &schema.Schema{Type: schema.TypeString},
										Optional:  true,
										Sensitive: true,
										Computed:  true,
									},
									"sso_username": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"url": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"vnc_keyboard_layout": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"width": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"client_src_range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"clipboard": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"custom_lang": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"customize_forticlient_download_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"default_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_window_height": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"default_window_width": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dhcp_ip_overlap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp_ra_giaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp_reservation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp6_ra_linkaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_bookmark": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_connection_tools": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_history": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_suffix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"exclusive_routing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"focus_bookmark": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forticlient_download": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"forticlient_download_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"heading": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hide_sso_credential": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_check_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"host_check_policy": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ip_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_pools": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ipv6_dns_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_dns_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_exclusive_routing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_pools": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ipv6_service_restriction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_split_tunneling": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_split_tunneling_routing_address": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ipv6_split_tunneling_routing_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_tunnel_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_wins_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_wins_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"keep_alive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"landing_page": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"form_data": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"value": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"sso": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sso_credential": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sso_password": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"sso_username": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"landing_page_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"limit_user_logins": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_addr_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mac_addr_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mac_addr_check_rule": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mac_addr_list": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"mac_addr_mask": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"macos_forticlient_download_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"os_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"os_check_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"latest_patch_level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"minor_version": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tolerance": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"prefer_ipv6_dns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redir_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rewrite_ip_uri_ui": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"save_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_restriction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"skip_check_for_browser": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"skip_check_for_unsupported_os": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"smb_max_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"smb_min_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"smb_ntlmv1_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"split_dns": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dns_server1": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dns_server2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"domains": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ipv6_dns_server1": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipv6_dns_server2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"split_tunneling": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"split_tunneling_routing_address": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"split_tunneling_routing_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"theme": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tunnel_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"use_sdwan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_bookmark": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_group_bookmark": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"web_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"windows_forticlient_download_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wins_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wins_server2": &schema.Schema{
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

func resourceVpnSslWebPortalCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnSslWebPortal(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnSslWebPortal resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadVpnSslWebPortal(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateVpnSslWebPortal(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating VpnSslWebPortal resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateVpnSslWebPortal(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating VpnSslWebPortal resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceVpnSslWebPortalRead(d, m)
}

func resourceVpnSslWebPortalUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnSslWebPortal(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebPortal resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateVpnSslWebPortal(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebPortal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceVpnSslWebPortalRead(d, m)
}

func resourceVpnSslWebPortalDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteVpnSslWebPortal(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VpnSslWebPortal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnSslWebPortalRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnSslWebPortal(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading VpnSslWebPortal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnSslWebPortal(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebPortal resource from API: %v", err)
	}
	return nil
}

func flattenVpnSslWebPortalAllowUserAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslWebPortalAutoConnect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bookmarks"
		if _, ok := i["bookmarks"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarks(i["bookmarks"], d, pre_append)
			tmp["bookmarks"] = fortiAPISubPartPatch(v, "VpnSslWebPortal-BookmarkGroup-Bookmarks")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "VpnSslWebPortal-BookmarkGroup-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenVpnSslWebPortalBookmarkGroupBookmarks(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_params"
		if _, ok := i["additional-params"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksAdditionalParams(i["additional-params"], d, pre_append)
			tmp["additional_params"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroup-Bookmarks-AdditionalParams")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apptype"
		if _, ok := i["apptype"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksApptype(i["apptype"], d, pre_append)
			tmp["apptype"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroup-Bookmarks-Apptype")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color_depth"
		if _, ok := i["color-depth"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksColorDepth(i["color-depth"], d, pre_append)
			tmp["color_depth"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroup-Bookmarks-ColorDepth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroup-Bookmarks-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := i["domain"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksDomain(i["domain"], d, pre_append)
			tmp["domain"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroup-Bookmarks-Domain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "folder"
		if _, ok := i["folder"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksFolder(i["folder"], d, pre_append)
			tmp["folder"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroup-Bookmarks-Folder")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "form_data"
		if _, ok := i["form-data"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksFormData(i["form-data"], d, pre_append)
			tmp["form_data"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroup-Bookmarks-FormData")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "height"
		if _, ok := i["height"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksHeight(i["height"], d, pre_append)
			tmp["height"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroup-Bookmarks-Height")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := i["host"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksHost(i["host"], d, pre_append)
			tmp["host"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroup-Bookmarks-Host")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keyboard_layout"
		if _, ok := i["keyboard-layout"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksKeyboardLayout(i["keyboard-layout"], d, pre_append)
			tmp["keyboard_layout"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroup-Bookmarks-KeyboardLayout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listening_port"
		if _, ok := i["listening-port"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksListeningPort(i["listening-port"], d, pre_append)
			tmp["listening_port"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroup-Bookmarks-ListeningPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "load_balancing_info"
		if _, ok := i["load-balancing-info"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksLoadBalancingInfo(i["load-balancing-info"], d, pre_append)
			tmp["load_balancing_info"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroup-Bookmarks-LoadBalancingInfo")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_user"
		if _, ok := i["logon-user"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksLogonUser(i["logon-user"], d, pre_append)
			tmp["logon_user"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroup-Bookmarks-LogonUser")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroup-Bookmarks-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroup-Bookmarks-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_blob"
		if _, ok := i["preconnection-blob"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksPreconnectionBlob(i["preconnection-blob"], d, pre_append)
			tmp["preconnection_blob"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroup-Bookmarks-PreconnectionBlob")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_id"
		if _, ok := i["preconnection-id"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksPreconnectionId(i["preconnection-id"], d, pre_append)
			tmp["preconnection_id"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroup-Bookmarks-PreconnectionId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_port"
		if _, ok := i["remote-port"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksRemotePort(i["remote-port"], d, pre_append)
			tmp["remote_port"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroup-Bookmarks-RemotePort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restricted_admin"
		if _, ok := i["restricted-admin"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksRestrictedAdmin(i["restricted-admin"], d, pre_append)
			tmp["restricted_admin"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroup-Bookmarks-RestrictedAdmin")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security"
		if _, ok := i["security"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksSecurity(i["security"], d, pre_append)
			tmp["security"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroup-Bookmarks-Security")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_preconnection_id"
		if _, ok := i["send-preconnection-id"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksSendPreconnectionId(i["send-preconnection-id"], d, pre_append)
			tmp["send_preconnection_id"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroup-Bookmarks-SendPreconnectionId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_layout"
		if _, ok := i["server-layout"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksServerLayout(i["server-layout"], d, pre_append)
			tmp["server_layout"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroup-Bookmarks-ServerLayout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "show_status_window"
		if _, ok := i["show-status-window"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksShowStatusWindow(i["show-status-window"], d, pre_append)
			tmp["show_status_window"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroup-Bookmarks-ShowStatusWindow")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso"
		if _, ok := i["sso"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksSso(i["sso"], d, pre_append)
			tmp["sso"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroup-Bookmarks-Sso")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential"
		if _, ok := i["sso-credential"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksSsoCredential(i["sso-credential"], d, pre_append)
			tmp["sso_credential"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroup-Bookmarks-SsoCredential")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential_sent_once"
		if _, ok := i["sso-credential-sent-once"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksSsoCredentialSentOnce(i["sso-credential-sent-once"], d, pre_append)
			tmp["sso_credential_sent_once"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroup-Bookmarks-SsoCredentialSentOnce")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_username"
		if _, ok := i["sso-username"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksSsoUsername(i["sso-username"], d, pre_append)
			tmp["sso_username"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroup-Bookmarks-SsoUsername")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := i["url"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksUrl(i["url"], d, pre_append)
			tmp["url"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroup-Bookmarks-Url")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vnc_keyboard_layout"
		if _, ok := i["vnc-keyboard-layout"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksVncKeyboardLayout(i["vnc-keyboard-layout"], d, pre_append)
			tmp["vnc_keyboard_layout"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroup-Bookmarks-VncKeyboardLayout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "width"
		if _, ok := i["width"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksWidth(i["width"], d, pre_append)
			tmp["width"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroup-Bookmarks-Width")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksAdditionalParams(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksApptype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksColorDepth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksFolder(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksFormData(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksFormDataName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroupBookmarks-FormData-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenVpnSslWebPortalBookmarkGroupBookmarksFormDataValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "VpnSslWebPortalBookmarkGroupBookmarks-FormData-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksFormDataName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksFormDataValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksHeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksKeyboardLayout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksListeningPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksLoadBalancingInfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksLogonUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksPreconnectionBlob(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksPreconnectionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksRemotePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksRestrictedAdmin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksSecurity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksSendPreconnectionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksServerLayout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksShowStatusWindow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksSso(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksSsoCredential(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksSsoCredentialSentOnce(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksSsoUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksVncKeyboardLayout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksWidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalClientSrcRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalClipboard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalCustomLang(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslWebPortalCustomizeForticlientDownloadUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalDefaultProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalDefaultWindowHeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalDefaultWindowWidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalDhcpIpOverlap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalDhcpRaGiaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalDhcpReservation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalDhcp6RaLinkaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalDisplayBookmark(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalDisplayConnectionTools(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalDisplayHistory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalDisplayStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalDnsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalDnsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalDnsSuffix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalExclusiveRouting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalFocusBookmark(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalForticlientDownload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalForticlientDownloadMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalHeading(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalHideSsoCredential(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalHostCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalHostCheckInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalHostCheckPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslWebPortalIpMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalIpPools(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslWebPortalIpv6DnsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalIpv6DnsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalIpv6ExclusiveRouting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalIpv6Pools(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslWebPortalIpv6ServiceRestriction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalIpv6SplitTunneling(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalIpv6SplitTunnelingRoutingAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslWebPortalIpv6SplitTunnelingRoutingNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalIpv6TunnelMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalIpv6WinsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalIpv6WinsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalKeepAlive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalLandingPage(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "form_data"
	if _, ok := i["form-data"]; ok {
		result["form_data"] = flattenVpnSslWebPortalLandingPageFormData(i["form-data"], d, pre_append)
	}

	pre_append = pre + ".0." + "sso"
	if _, ok := i["sso"]; ok {
		result["sso"] = flattenVpnSslWebPortalLandingPageSso(i["sso"], d, pre_append)
	}

	pre_append = pre + ".0." + "sso_credential"
	if _, ok := i["sso-credential"]; ok {
		result["sso_credential"] = flattenVpnSslWebPortalLandingPageSsoCredential(i["sso-credential"], d, pre_append)
	}

	pre_append = pre + ".0." + "sso_username"
	if _, ok := i["sso-username"]; ok {
		result["sso_username"] = flattenVpnSslWebPortalLandingPageSsoUsername(i["sso-username"], d, pre_append)
	}

	pre_append = pre + ".0." + "url"
	if _, ok := i["url"]; ok {
		result["url"] = flattenVpnSslWebPortalLandingPageUrl(i["url"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenVpnSslWebPortalLandingPageFormData(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenVpnSslWebPortalLandingPageFormDataName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "VpnSslWebPortalLandingPage-FormData-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenVpnSslWebPortalLandingPageFormDataValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "VpnSslWebPortalLandingPage-FormData-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenVpnSslWebPortalLandingPageFormDataName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalLandingPageFormDataValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalLandingPageSso(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalLandingPageSsoCredential(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalLandingPageSsoUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalLandingPageUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalLandingPageMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalLimitUserLogins(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalMacAddrAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalMacAddrCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalMacAddrCheckRule(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_addr_list"
		if _, ok := i["mac-addr-list"]; ok {
			v := flattenVpnSslWebPortalMacAddrCheckRuleMacAddrList(i["mac-addr-list"], d, pre_append)
			tmp["mac_addr_list"] = fortiAPISubPartPatch(v, "VpnSslWebPortal-MacAddrCheckRule-MacAddrList")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_addr_mask"
		if _, ok := i["mac-addr-mask"]; ok {
			v := flattenVpnSslWebPortalMacAddrCheckRuleMacAddrMask(i["mac-addr-mask"], d, pre_append)
			tmp["mac_addr_mask"] = fortiAPISubPartPatch(v, "VpnSslWebPortal-MacAddrCheckRule-MacAddrMask")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenVpnSslWebPortalMacAddrCheckRuleName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "VpnSslWebPortal-MacAddrCheckRule-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenVpnSslWebPortalMacAddrCheckRuleMacAddrList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslWebPortalMacAddrCheckRuleMacAddrMask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalMacAddrCheckRuleName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalMacosForticlientDownloadUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalOsCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalOsCheckList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenVpnSslWebPortalOsCheckListAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "latest_patch_level"
	if _, ok := i["latest-patch-level"]; ok {
		result["latest_patch_level"] = flattenVpnSslWebPortalOsCheckListLatestPatchLevel(i["latest-patch-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "minor_version"
	if _, ok := i["minor-version"]; ok {
		result["minor_version"] = flattenVpnSslWebPortalOsCheckListMinorVersion(i["minor-version"], d, pre_append)
	}

	pre_append = pre + ".0." + "name"
	if _, ok := i["name"]; ok {
		result["name"] = flattenVpnSslWebPortalOsCheckListName(i["name"], d, pre_append)
	}

	pre_append = pre + ".0." + "tolerance"
	if _, ok := i["tolerance"]; ok {
		result["tolerance"] = flattenVpnSslWebPortalOsCheckListTolerance(i["tolerance"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenVpnSslWebPortalOsCheckListAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalOsCheckListLatestPatchLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalOsCheckListMinorVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalOsCheckListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalOsCheckListTolerance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalPreferIpv6Dns(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalRedirUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalRewriteIpUriUi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalSavePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalServiceRestriction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalSkipCheckForBrowser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalSkipCheckForUnsupportedOs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalSmbMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalSmbMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalSmbNtlmv1Auth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalSplitDns(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_server1"
		if _, ok := i["dns-server1"]; ok {
			v := flattenVpnSslWebPortalSplitDnsDnsServer1(i["dns-server1"], d, pre_append)
			tmp["dns_server1"] = fortiAPISubPartPatch(v, "VpnSslWebPortal-SplitDns-DnsServer1")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_server2"
		if _, ok := i["dns-server2"]; ok {
			v := flattenVpnSslWebPortalSplitDnsDnsServer2(i["dns-server2"], d, pre_append)
			tmp["dns_server2"] = fortiAPISubPartPatch(v, "VpnSslWebPortal-SplitDns-DnsServer2")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domains"
		if _, ok := i["domains"]; ok {
			v := flattenVpnSslWebPortalSplitDnsDomains(i["domains"], d, pre_append)
			tmp["domains"] = fortiAPISubPartPatch(v, "VpnSslWebPortal-SplitDns-Domains")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenVpnSslWebPortalSplitDnsId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "VpnSslWebPortal-SplitDns-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_dns_server1"
		if _, ok := i["ipv6-dns-server1"]; ok {
			v := flattenVpnSslWebPortalSplitDnsIpv6DnsServer1(i["ipv6-dns-server1"], d, pre_append)
			tmp["ipv6_dns_server1"] = fortiAPISubPartPatch(v, "VpnSslWebPortal-SplitDns-Ipv6DnsServer1")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_dns_server2"
		if _, ok := i["ipv6-dns-server2"]; ok {
			v := flattenVpnSslWebPortalSplitDnsIpv6DnsServer2(i["ipv6-dns-server2"], d, pre_append)
			tmp["ipv6_dns_server2"] = fortiAPISubPartPatch(v, "VpnSslWebPortal-SplitDns-Ipv6DnsServer2")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenVpnSslWebPortalSplitDnsDnsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalSplitDnsDnsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalSplitDnsDomains(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalSplitDnsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalSplitDnsIpv6DnsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalSplitDnsIpv6DnsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalSplitTunneling(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalSplitTunnelingRoutingAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslWebPortalSplitTunnelingRoutingNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalTheme(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalTunnelMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalUseSdwan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalUserBookmark(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalUserGroupBookmark(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalWebMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalWindowsForticlientDownloadUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalWinsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalWinsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnSslWebPortal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("allow_user_access", flattenVpnSslWebPortalAllowUserAccess(o["allow-user-access"], d, "allow_user_access")); err != nil {
		if vv, ok := fortiAPIPatch(o["allow-user-access"], "VpnSslWebPortal-AllowUserAccess"); ok {
			if err = d.Set("allow_user_access", vv); err != nil {
				return fmt.Errorf("Error reading allow_user_access: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allow_user_access: %v", err)
		}
	}

	if err = d.Set("auto_connect", flattenVpnSslWebPortalAutoConnect(o["auto-connect"], d, "auto_connect")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-connect"], "VpnSslWebPortal-AutoConnect"); ok {
			if err = d.Set("auto_connect", vv); err != nil {
				return fmt.Errorf("Error reading auto_connect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_connect: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("bookmark_group", flattenVpnSslWebPortalBookmarkGroup(o["bookmark-group"], d, "bookmark_group")); err != nil {
			if vv, ok := fortiAPIPatch(o["bookmark-group"], "VpnSslWebPortal-BookmarkGroup"); ok {
				if err = d.Set("bookmark_group", vv); err != nil {
					return fmt.Errorf("Error reading bookmark_group: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading bookmark_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("bookmark_group"); ok {
			if err = d.Set("bookmark_group", flattenVpnSslWebPortalBookmarkGroup(o["bookmark-group"], d, "bookmark_group")); err != nil {
				if vv, ok := fortiAPIPatch(o["bookmark-group"], "VpnSslWebPortal-BookmarkGroup"); ok {
					if err = d.Set("bookmark_group", vv); err != nil {
						return fmt.Errorf("Error reading bookmark_group: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading bookmark_group: %v", err)
				}
			}
		}
	}

	if err = d.Set("client_src_range", flattenVpnSslWebPortalClientSrcRange(o["client-src-range"], d, "client_src_range")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-src-range"], "VpnSslWebPortal-ClientSrcRange"); ok {
			if err = d.Set("client_src_range", vv); err != nil {
				return fmt.Errorf("Error reading client_src_range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_src_range: %v", err)
		}
	}

	if err = d.Set("clipboard", flattenVpnSslWebPortalClipboard(o["clipboard"], d, "clipboard")); err != nil {
		if vv, ok := fortiAPIPatch(o["clipboard"], "VpnSslWebPortal-Clipboard"); ok {
			if err = d.Set("clipboard", vv); err != nil {
				return fmt.Errorf("Error reading clipboard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading clipboard: %v", err)
		}
	}

	if err = d.Set("custom_lang", flattenVpnSslWebPortalCustomLang(o["custom-lang"], d, "custom_lang")); err != nil {
		if vv, ok := fortiAPIPatch(o["custom-lang"], "VpnSslWebPortal-CustomLang"); ok {
			if err = d.Set("custom_lang", vv); err != nil {
				return fmt.Errorf("Error reading custom_lang: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading custom_lang: %v", err)
		}
	}

	if err = d.Set("customize_forticlient_download_url", flattenVpnSslWebPortalCustomizeForticlientDownloadUrl(o["customize-forticlient-download-url"], d, "customize_forticlient_download_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["customize-forticlient-download-url"], "VpnSslWebPortal-CustomizeForticlientDownloadUrl"); ok {
			if err = d.Set("customize_forticlient_download_url", vv); err != nil {
				return fmt.Errorf("Error reading customize_forticlient_download_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading customize_forticlient_download_url: %v", err)
		}
	}

	if err = d.Set("default_protocol", flattenVpnSslWebPortalDefaultProtocol(o["default-protocol"], d, "default_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-protocol"], "VpnSslWebPortal-DefaultProtocol"); ok {
			if err = d.Set("default_protocol", vv); err != nil {
				return fmt.Errorf("Error reading default_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_protocol: %v", err)
		}
	}

	if err = d.Set("default_window_height", flattenVpnSslWebPortalDefaultWindowHeight(o["default-window-height"], d, "default_window_height")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-window-height"], "VpnSslWebPortal-DefaultWindowHeight"); ok {
			if err = d.Set("default_window_height", vv); err != nil {
				return fmt.Errorf("Error reading default_window_height: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_window_height: %v", err)
		}
	}

	if err = d.Set("default_window_width", flattenVpnSslWebPortalDefaultWindowWidth(o["default-window-width"], d, "default_window_width")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-window-width"], "VpnSslWebPortal-DefaultWindowWidth"); ok {
			if err = d.Set("default_window_width", vv); err != nil {
				return fmt.Errorf("Error reading default_window_width: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_window_width: %v", err)
		}
	}

	if err = d.Set("dhcp_ip_overlap", flattenVpnSslWebPortalDhcpIpOverlap(o["dhcp-ip-overlap"], d, "dhcp_ip_overlap")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-ip-overlap"], "VpnSslWebPortal-DhcpIpOverlap"); ok {
			if err = d.Set("dhcp_ip_overlap", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_ip_overlap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_ip_overlap: %v", err)
		}
	}

	if err = d.Set("dhcp_ra_giaddr", flattenVpnSslWebPortalDhcpRaGiaddr(o["dhcp-ra-giaddr"], d, "dhcp_ra_giaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-ra-giaddr"], "VpnSslWebPortal-DhcpRaGiaddr"); ok {
			if err = d.Set("dhcp_ra_giaddr", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_ra_giaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_ra_giaddr: %v", err)
		}
	}

	if err = d.Set("dhcp_reservation", flattenVpnSslWebPortalDhcpReservation(o["dhcp-reservation"], d, "dhcp_reservation")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-reservation"], "VpnSslWebPortal-DhcpReservation"); ok {
			if err = d.Set("dhcp_reservation", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_reservation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_reservation: %v", err)
		}
	}

	if err = d.Set("dhcp6_ra_linkaddr", flattenVpnSslWebPortalDhcp6RaLinkaddr(o["dhcp6-ra-linkaddr"], d, "dhcp6_ra_linkaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp6-ra-linkaddr"], "VpnSslWebPortal-Dhcp6RaLinkaddr"); ok {
			if err = d.Set("dhcp6_ra_linkaddr", vv); err != nil {
				return fmt.Errorf("Error reading dhcp6_ra_linkaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp6_ra_linkaddr: %v", err)
		}
	}

	if err = d.Set("display_bookmark", flattenVpnSslWebPortalDisplayBookmark(o["display-bookmark"], d, "display_bookmark")); err != nil {
		if vv, ok := fortiAPIPatch(o["display-bookmark"], "VpnSslWebPortal-DisplayBookmark"); ok {
			if err = d.Set("display_bookmark", vv); err != nil {
				return fmt.Errorf("Error reading display_bookmark: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading display_bookmark: %v", err)
		}
	}

	if err = d.Set("display_connection_tools", flattenVpnSslWebPortalDisplayConnectionTools(o["display-connection-tools"], d, "display_connection_tools")); err != nil {
		if vv, ok := fortiAPIPatch(o["display-connection-tools"], "VpnSslWebPortal-DisplayConnectionTools"); ok {
			if err = d.Set("display_connection_tools", vv); err != nil {
				return fmt.Errorf("Error reading display_connection_tools: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading display_connection_tools: %v", err)
		}
	}

	if err = d.Set("display_history", flattenVpnSslWebPortalDisplayHistory(o["display-history"], d, "display_history")); err != nil {
		if vv, ok := fortiAPIPatch(o["display-history"], "VpnSslWebPortal-DisplayHistory"); ok {
			if err = d.Set("display_history", vv); err != nil {
				return fmt.Errorf("Error reading display_history: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading display_history: %v", err)
		}
	}

	if err = d.Set("display_status", flattenVpnSslWebPortalDisplayStatus(o["display-status"], d, "display_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["display-status"], "VpnSslWebPortal-DisplayStatus"); ok {
			if err = d.Set("display_status", vv); err != nil {
				return fmt.Errorf("Error reading display_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading display_status: %v", err)
		}
	}

	if err = d.Set("dns_server1", flattenVpnSslWebPortalDnsServer1(o["dns-server1"], d, "dns_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server1"], "VpnSslWebPortal-DnsServer1"); ok {
			if err = d.Set("dns_server1", vv); err != nil {
				return fmt.Errorf("Error reading dns_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server1: %v", err)
		}
	}

	if err = d.Set("dns_server2", flattenVpnSslWebPortalDnsServer2(o["dns-server2"], d, "dns_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server2"], "VpnSslWebPortal-DnsServer2"); ok {
			if err = d.Set("dns_server2", vv); err != nil {
				return fmt.Errorf("Error reading dns_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server2: %v", err)
		}
	}

	if err = d.Set("dns_suffix", flattenVpnSslWebPortalDnsSuffix(o["dns-suffix"], d, "dns_suffix")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-suffix"], "VpnSslWebPortal-DnsSuffix"); ok {
			if err = d.Set("dns_suffix", vv); err != nil {
				return fmt.Errorf("Error reading dns_suffix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_suffix: %v", err)
		}
	}

	if err = d.Set("exclusive_routing", flattenVpnSslWebPortalExclusiveRouting(o["exclusive-routing"], d, "exclusive_routing")); err != nil {
		if vv, ok := fortiAPIPatch(o["exclusive-routing"], "VpnSslWebPortal-ExclusiveRouting"); ok {
			if err = d.Set("exclusive_routing", vv); err != nil {
				return fmt.Errorf("Error reading exclusive_routing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exclusive_routing: %v", err)
		}
	}

	if err = d.Set("focus_bookmark", flattenVpnSslWebPortalFocusBookmark(o["focus-bookmark"], d, "focus_bookmark")); err != nil {
		if vv, ok := fortiAPIPatch(o["focus-bookmark"], "VpnSslWebPortal-FocusBookmark"); ok {
			if err = d.Set("focus_bookmark", vv); err != nil {
				return fmt.Errorf("Error reading focus_bookmark: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading focus_bookmark: %v", err)
		}
	}

	if err = d.Set("forticlient_download", flattenVpnSslWebPortalForticlientDownload(o["forticlient-download"], d, "forticlient_download")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticlient-download"], "VpnSslWebPortal-ForticlientDownload"); ok {
			if err = d.Set("forticlient_download", vv); err != nil {
				return fmt.Errorf("Error reading forticlient_download: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticlient_download: %v", err)
		}
	}

	if err = d.Set("forticlient_download_method", flattenVpnSslWebPortalForticlientDownloadMethod(o["forticlient-download-method"], d, "forticlient_download_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticlient-download-method"], "VpnSslWebPortal-ForticlientDownloadMethod"); ok {
			if err = d.Set("forticlient_download_method", vv); err != nil {
				return fmt.Errorf("Error reading forticlient_download_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticlient_download_method: %v", err)
		}
	}

	if err = d.Set("heading", flattenVpnSslWebPortalHeading(o["heading"], d, "heading")); err != nil {
		if vv, ok := fortiAPIPatch(o["heading"], "VpnSslWebPortal-Heading"); ok {
			if err = d.Set("heading", vv); err != nil {
				return fmt.Errorf("Error reading heading: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading heading: %v", err)
		}
	}

	if err = d.Set("hide_sso_credential", flattenVpnSslWebPortalHideSsoCredential(o["hide-sso-credential"], d, "hide_sso_credential")); err != nil {
		if vv, ok := fortiAPIPatch(o["hide-sso-credential"], "VpnSslWebPortal-HideSsoCredential"); ok {
			if err = d.Set("hide_sso_credential", vv); err != nil {
				return fmt.Errorf("Error reading hide_sso_credential: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hide_sso_credential: %v", err)
		}
	}

	if err = d.Set("host_check", flattenVpnSslWebPortalHostCheck(o["host-check"], d, "host_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["host-check"], "VpnSslWebPortal-HostCheck"); ok {
			if err = d.Set("host_check", vv); err != nil {
				return fmt.Errorf("Error reading host_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host_check: %v", err)
		}
	}

	if err = d.Set("host_check_interval", flattenVpnSslWebPortalHostCheckInterval(o["host-check-interval"], d, "host_check_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["host-check-interval"], "VpnSslWebPortal-HostCheckInterval"); ok {
			if err = d.Set("host_check_interval", vv); err != nil {
				return fmt.Errorf("Error reading host_check_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host_check_interval: %v", err)
		}
	}

	if err = d.Set("host_check_policy", flattenVpnSslWebPortalHostCheckPolicy(o["host-check-policy"], d, "host_check_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["host-check-policy"], "VpnSslWebPortal-HostCheckPolicy"); ok {
			if err = d.Set("host_check_policy", vv); err != nil {
				return fmt.Errorf("Error reading host_check_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host_check_policy: %v", err)
		}
	}

	if err = d.Set("ip_mode", flattenVpnSslWebPortalIpMode(o["ip-mode"], d, "ip_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-mode"], "VpnSslWebPortal-IpMode"); ok {
			if err = d.Set("ip_mode", vv); err != nil {
				return fmt.Errorf("Error reading ip_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_mode: %v", err)
		}
	}

	if err = d.Set("ip_pools", flattenVpnSslWebPortalIpPools(o["ip-pools"], d, "ip_pools")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-pools"], "VpnSslWebPortal-IpPools"); ok {
			if err = d.Set("ip_pools", vv); err != nil {
				return fmt.Errorf("Error reading ip_pools: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_pools: %v", err)
		}
	}

	if err = d.Set("ipv6_dns_server1", flattenVpnSslWebPortalIpv6DnsServer1(o["ipv6-dns-server1"], d, "ipv6_dns_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-dns-server1"], "VpnSslWebPortal-Ipv6DnsServer1"); ok {
			if err = d.Set("ipv6_dns_server1", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_dns_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_dns_server1: %v", err)
		}
	}

	if err = d.Set("ipv6_dns_server2", flattenVpnSslWebPortalIpv6DnsServer2(o["ipv6-dns-server2"], d, "ipv6_dns_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-dns-server2"], "VpnSslWebPortal-Ipv6DnsServer2"); ok {
			if err = d.Set("ipv6_dns_server2", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_dns_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_dns_server2: %v", err)
		}
	}

	if err = d.Set("ipv6_exclusive_routing", flattenVpnSslWebPortalIpv6ExclusiveRouting(o["ipv6-exclusive-routing"], d, "ipv6_exclusive_routing")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-exclusive-routing"], "VpnSslWebPortal-Ipv6ExclusiveRouting"); ok {
			if err = d.Set("ipv6_exclusive_routing", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_exclusive_routing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_exclusive_routing: %v", err)
		}
	}

	if err = d.Set("ipv6_pools", flattenVpnSslWebPortalIpv6Pools(o["ipv6-pools"], d, "ipv6_pools")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-pools"], "VpnSslWebPortal-Ipv6Pools"); ok {
			if err = d.Set("ipv6_pools", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_pools: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_pools: %v", err)
		}
	}

	if err = d.Set("ipv6_service_restriction", flattenVpnSslWebPortalIpv6ServiceRestriction(o["ipv6-service-restriction"], d, "ipv6_service_restriction")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-service-restriction"], "VpnSslWebPortal-Ipv6ServiceRestriction"); ok {
			if err = d.Set("ipv6_service_restriction", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_service_restriction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_service_restriction: %v", err)
		}
	}

	if err = d.Set("ipv6_split_tunneling", flattenVpnSslWebPortalIpv6SplitTunneling(o["ipv6-split-tunneling"], d, "ipv6_split_tunneling")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-split-tunneling"], "VpnSslWebPortal-Ipv6SplitTunneling"); ok {
			if err = d.Set("ipv6_split_tunneling", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_split_tunneling: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_split_tunneling: %v", err)
		}
	}

	if err = d.Set("ipv6_split_tunneling_routing_address", flattenVpnSslWebPortalIpv6SplitTunnelingRoutingAddress(o["ipv6-split-tunneling-routing-address"], d, "ipv6_split_tunneling_routing_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-split-tunneling-routing-address"], "VpnSslWebPortal-Ipv6SplitTunnelingRoutingAddress"); ok {
			if err = d.Set("ipv6_split_tunneling_routing_address", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_split_tunneling_routing_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_split_tunneling_routing_address: %v", err)
		}
	}

	if err = d.Set("ipv6_split_tunneling_routing_negate", flattenVpnSslWebPortalIpv6SplitTunnelingRoutingNegate(o["ipv6-split-tunneling-routing-negate"], d, "ipv6_split_tunneling_routing_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-split-tunneling-routing-negate"], "VpnSslWebPortal-Ipv6SplitTunnelingRoutingNegate"); ok {
			if err = d.Set("ipv6_split_tunneling_routing_negate", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_split_tunneling_routing_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_split_tunneling_routing_negate: %v", err)
		}
	}

	if err = d.Set("ipv6_tunnel_mode", flattenVpnSslWebPortalIpv6TunnelMode(o["ipv6-tunnel-mode"], d, "ipv6_tunnel_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-tunnel-mode"], "VpnSslWebPortal-Ipv6TunnelMode"); ok {
			if err = d.Set("ipv6_tunnel_mode", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_tunnel_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_tunnel_mode: %v", err)
		}
	}

	if err = d.Set("ipv6_wins_server1", flattenVpnSslWebPortalIpv6WinsServer1(o["ipv6-wins-server1"], d, "ipv6_wins_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-wins-server1"], "VpnSslWebPortal-Ipv6WinsServer1"); ok {
			if err = d.Set("ipv6_wins_server1", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_wins_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_wins_server1: %v", err)
		}
	}

	if err = d.Set("ipv6_wins_server2", flattenVpnSslWebPortalIpv6WinsServer2(o["ipv6-wins-server2"], d, "ipv6_wins_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-wins-server2"], "VpnSslWebPortal-Ipv6WinsServer2"); ok {
			if err = d.Set("ipv6_wins_server2", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_wins_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_wins_server2: %v", err)
		}
	}

	if err = d.Set("keep_alive", flattenVpnSslWebPortalKeepAlive(o["keep-alive"], d, "keep_alive")); err != nil {
		if vv, ok := fortiAPIPatch(o["keep-alive"], "VpnSslWebPortal-KeepAlive"); ok {
			if err = d.Set("keep_alive", vv); err != nil {
				return fmt.Errorf("Error reading keep_alive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keep_alive: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("landing_page", flattenVpnSslWebPortalLandingPage(o["landing-page"], d, "landing_page")); err != nil {
			if vv, ok := fortiAPIPatch(o["landing-page"], "VpnSslWebPortal-LandingPage"); ok {
				if err = d.Set("landing_page", vv); err != nil {
					return fmt.Errorf("Error reading landing_page: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading landing_page: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("landing_page"); ok {
			if err = d.Set("landing_page", flattenVpnSslWebPortalLandingPage(o["landing-page"], d, "landing_page")); err != nil {
				if vv, ok := fortiAPIPatch(o["landing-page"], "VpnSslWebPortal-LandingPage"); ok {
					if err = d.Set("landing_page", vv); err != nil {
						return fmt.Errorf("Error reading landing_page: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading landing_page: %v", err)
				}
			}
		}
	}

	if err = d.Set("landing_page_mode", flattenVpnSslWebPortalLandingPageMode(o["landing-page-mode"], d, "landing_page_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["landing-page-mode"], "VpnSslWebPortal-LandingPageMode"); ok {
			if err = d.Set("landing_page_mode", vv); err != nil {
				return fmt.Errorf("Error reading landing_page_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading landing_page_mode: %v", err)
		}
	}

	if err = d.Set("limit_user_logins", flattenVpnSslWebPortalLimitUserLogins(o["limit-user-logins"], d, "limit_user_logins")); err != nil {
		if vv, ok := fortiAPIPatch(o["limit-user-logins"], "VpnSslWebPortal-LimitUserLogins"); ok {
			if err = d.Set("limit_user_logins", vv); err != nil {
				return fmt.Errorf("Error reading limit_user_logins: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading limit_user_logins: %v", err)
		}
	}

	if err = d.Set("mac_addr_action", flattenVpnSslWebPortalMacAddrAction(o["mac-addr-action"], d, "mac_addr_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-addr-action"], "VpnSslWebPortal-MacAddrAction"); ok {
			if err = d.Set("mac_addr_action", vv); err != nil {
				return fmt.Errorf("Error reading mac_addr_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_addr_action: %v", err)
		}
	}

	if err = d.Set("mac_addr_check", flattenVpnSslWebPortalMacAddrCheck(o["mac-addr-check"], d, "mac_addr_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-addr-check"], "VpnSslWebPortal-MacAddrCheck"); ok {
			if err = d.Set("mac_addr_check", vv); err != nil {
				return fmt.Errorf("Error reading mac_addr_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_addr_check: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("mac_addr_check_rule", flattenVpnSslWebPortalMacAddrCheckRule(o["mac-addr-check-rule"], d, "mac_addr_check_rule")); err != nil {
			if vv, ok := fortiAPIPatch(o["mac-addr-check-rule"], "VpnSslWebPortal-MacAddrCheckRule"); ok {
				if err = d.Set("mac_addr_check_rule", vv); err != nil {
					return fmt.Errorf("Error reading mac_addr_check_rule: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mac_addr_check_rule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mac_addr_check_rule"); ok {
			if err = d.Set("mac_addr_check_rule", flattenVpnSslWebPortalMacAddrCheckRule(o["mac-addr-check-rule"], d, "mac_addr_check_rule")); err != nil {
				if vv, ok := fortiAPIPatch(o["mac-addr-check-rule"], "VpnSslWebPortal-MacAddrCheckRule"); ok {
					if err = d.Set("mac_addr_check_rule", vv); err != nil {
						return fmt.Errorf("Error reading mac_addr_check_rule: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mac_addr_check_rule: %v", err)
				}
			}
		}
	}

	if err = d.Set("macos_forticlient_download_url", flattenVpnSslWebPortalMacosForticlientDownloadUrl(o["macos-forticlient-download-url"], d, "macos_forticlient_download_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["macos-forticlient-download-url"], "VpnSslWebPortal-MacosForticlientDownloadUrl"); ok {
			if err = d.Set("macos_forticlient_download_url", vv); err != nil {
				return fmt.Errorf("Error reading macos_forticlient_download_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading macos_forticlient_download_url: %v", err)
		}
	}

	if err = d.Set("name", flattenVpnSslWebPortalName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "VpnSslWebPortal-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("os_check", flattenVpnSslWebPortalOsCheck(o["os-check"], d, "os_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["os-check"], "VpnSslWebPortal-OsCheck"); ok {
			if err = d.Set("os_check", vv); err != nil {
				return fmt.Errorf("Error reading os_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading os_check: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("os_check_list", flattenVpnSslWebPortalOsCheckList(o["os-check-list"], d, "os_check_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["os-check-list"], "VpnSslWebPortal-OsCheckList"); ok {
				if err = d.Set("os_check_list", vv); err != nil {
					return fmt.Errorf("Error reading os_check_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading os_check_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("os_check_list"); ok {
			if err = d.Set("os_check_list", flattenVpnSslWebPortalOsCheckList(o["os-check-list"], d, "os_check_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["os-check-list"], "VpnSslWebPortal-OsCheckList"); ok {
					if err = d.Set("os_check_list", vv); err != nil {
						return fmt.Errorf("Error reading os_check_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading os_check_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("prefer_ipv6_dns", flattenVpnSslWebPortalPreferIpv6Dns(o["prefer-ipv6-dns"], d, "prefer_ipv6_dns")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefer-ipv6-dns"], "VpnSslWebPortal-PreferIpv6Dns"); ok {
			if err = d.Set("prefer_ipv6_dns", vv); err != nil {
				return fmt.Errorf("Error reading prefer_ipv6_dns: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefer_ipv6_dns: %v", err)
		}
	}

	if err = d.Set("redir_url", flattenVpnSslWebPortalRedirUrl(o["redir-url"], d, "redir_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["redir-url"], "VpnSslWebPortal-RedirUrl"); ok {
			if err = d.Set("redir_url", vv); err != nil {
				return fmt.Errorf("Error reading redir_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redir_url: %v", err)
		}
	}

	if err = d.Set("rewrite_ip_uri_ui", flattenVpnSslWebPortalRewriteIpUriUi(o["rewrite-ip-uri-ui"], d, "rewrite_ip_uri_ui")); err != nil {
		if vv, ok := fortiAPIPatch(o["rewrite-ip-uri-ui"], "VpnSslWebPortal-RewriteIpUriUi"); ok {
			if err = d.Set("rewrite_ip_uri_ui", vv); err != nil {
				return fmt.Errorf("Error reading rewrite_ip_uri_ui: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rewrite_ip_uri_ui: %v", err)
		}
	}

	if err = d.Set("save_password", flattenVpnSslWebPortalSavePassword(o["save-password"], d, "save_password")); err != nil {
		if vv, ok := fortiAPIPatch(o["save-password"], "VpnSslWebPortal-SavePassword"); ok {
			if err = d.Set("save_password", vv); err != nil {
				return fmt.Errorf("Error reading save_password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading save_password: %v", err)
		}
	}

	if err = d.Set("service_restriction", flattenVpnSslWebPortalServiceRestriction(o["service-restriction"], d, "service_restriction")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-restriction"], "VpnSslWebPortal-ServiceRestriction"); ok {
			if err = d.Set("service_restriction", vv); err != nil {
				return fmt.Errorf("Error reading service_restriction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_restriction: %v", err)
		}
	}

	if err = d.Set("skip_check_for_browser", flattenVpnSslWebPortalSkipCheckForBrowser(o["skip-check-for-browser"], d, "skip_check_for_browser")); err != nil {
		if vv, ok := fortiAPIPatch(o["skip-check-for-browser"], "VpnSslWebPortal-SkipCheckForBrowser"); ok {
			if err = d.Set("skip_check_for_browser", vv); err != nil {
				return fmt.Errorf("Error reading skip_check_for_browser: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading skip_check_for_browser: %v", err)
		}
	}

	if err = d.Set("skip_check_for_unsupported_os", flattenVpnSslWebPortalSkipCheckForUnsupportedOs(o["skip-check-for-unsupported-os"], d, "skip_check_for_unsupported_os")); err != nil {
		if vv, ok := fortiAPIPatch(o["skip-check-for-unsupported-os"], "VpnSslWebPortal-SkipCheckForUnsupportedOs"); ok {
			if err = d.Set("skip_check_for_unsupported_os", vv); err != nil {
				return fmt.Errorf("Error reading skip_check_for_unsupported_os: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading skip_check_for_unsupported_os: %v", err)
		}
	}

	if err = d.Set("smb_max_version", flattenVpnSslWebPortalSmbMaxVersion(o["smb-max-version"], d, "smb_max_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["smb-max-version"], "VpnSslWebPortal-SmbMaxVersion"); ok {
			if err = d.Set("smb_max_version", vv); err != nil {
				return fmt.Errorf("Error reading smb_max_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading smb_max_version: %v", err)
		}
	}

	if err = d.Set("smb_min_version", flattenVpnSslWebPortalSmbMinVersion(o["smb-min-version"], d, "smb_min_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["smb-min-version"], "VpnSslWebPortal-SmbMinVersion"); ok {
			if err = d.Set("smb_min_version", vv); err != nil {
				return fmt.Errorf("Error reading smb_min_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading smb_min_version: %v", err)
		}
	}

	if err = d.Set("smb_ntlmv1_auth", flattenVpnSslWebPortalSmbNtlmv1Auth(o["smb-ntlmv1-auth"], d, "smb_ntlmv1_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["smb-ntlmv1-auth"], "VpnSslWebPortal-SmbNtlmv1Auth"); ok {
			if err = d.Set("smb_ntlmv1_auth", vv); err != nil {
				return fmt.Errorf("Error reading smb_ntlmv1_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading smb_ntlmv1_auth: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("split_dns", flattenVpnSslWebPortalSplitDns(o["split-dns"], d, "split_dns")); err != nil {
			if vv, ok := fortiAPIPatch(o["split-dns"], "VpnSslWebPortal-SplitDns"); ok {
				if err = d.Set("split_dns", vv); err != nil {
					return fmt.Errorf("Error reading split_dns: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading split_dns: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("split_dns"); ok {
			if err = d.Set("split_dns", flattenVpnSslWebPortalSplitDns(o["split-dns"], d, "split_dns")); err != nil {
				if vv, ok := fortiAPIPatch(o["split-dns"], "VpnSslWebPortal-SplitDns"); ok {
					if err = d.Set("split_dns", vv); err != nil {
						return fmt.Errorf("Error reading split_dns: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading split_dns: %v", err)
				}
			}
		}
	}

	if err = d.Set("split_tunneling", flattenVpnSslWebPortalSplitTunneling(o["split-tunneling"], d, "split_tunneling")); err != nil {
		if vv, ok := fortiAPIPatch(o["split-tunneling"], "VpnSslWebPortal-SplitTunneling"); ok {
			if err = d.Set("split_tunneling", vv); err != nil {
				return fmt.Errorf("Error reading split_tunneling: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading split_tunneling: %v", err)
		}
	}

	if err = d.Set("split_tunneling_routing_address", flattenVpnSslWebPortalSplitTunnelingRoutingAddress(o["split-tunneling-routing-address"], d, "split_tunneling_routing_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["split-tunneling-routing-address"], "VpnSslWebPortal-SplitTunnelingRoutingAddress"); ok {
			if err = d.Set("split_tunneling_routing_address", vv); err != nil {
				return fmt.Errorf("Error reading split_tunneling_routing_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading split_tunneling_routing_address: %v", err)
		}
	}

	if err = d.Set("split_tunneling_routing_negate", flattenVpnSslWebPortalSplitTunnelingRoutingNegate(o["split-tunneling-routing-negate"], d, "split_tunneling_routing_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["split-tunneling-routing-negate"], "VpnSslWebPortal-SplitTunnelingRoutingNegate"); ok {
			if err = d.Set("split_tunneling_routing_negate", vv); err != nil {
				return fmt.Errorf("Error reading split_tunneling_routing_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading split_tunneling_routing_negate: %v", err)
		}
	}

	if err = d.Set("theme", flattenVpnSslWebPortalTheme(o["theme"], d, "theme")); err != nil {
		if vv, ok := fortiAPIPatch(o["theme"], "VpnSslWebPortal-Theme"); ok {
			if err = d.Set("theme", vv); err != nil {
				return fmt.Errorf("Error reading theme: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading theme: %v", err)
		}
	}

	if err = d.Set("tunnel_mode", flattenVpnSslWebPortalTunnelMode(o["tunnel-mode"], d, "tunnel_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-mode"], "VpnSslWebPortal-TunnelMode"); ok {
			if err = d.Set("tunnel_mode", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_mode: %v", err)
		}
	}

	if err = d.Set("use_sdwan", flattenVpnSslWebPortalUseSdwan(o["use-sdwan"], d, "use_sdwan")); err != nil {
		if vv, ok := fortiAPIPatch(o["use-sdwan"], "VpnSslWebPortal-UseSdwan"); ok {
			if err = d.Set("use_sdwan", vv); err != nil {
				return fmt.Errorf("Error reading use_sdwan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading use_sdwan: %v", err)
		}
	}

	if err = d.Set("user_bookmark", flattenVpnSslWebPortalUserBookmark(o["user-bookmark"], d, "user_bookmark")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-bookmark"], "VpnSslWebPortal-UserBookmark"); ok {
			if err = d.Set("user_bookmark", vv); err != nil {
				return fmt.Errorf("Error reading user_bookmark: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_bookmark: %v", err)
		}
	}

	if err = d.Set("user_group_bookmark", flattenVpnSslWebPortalUserGroupBookmark(o["user-group-bookmark"], d, "user_group_bookmark")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-group-bookmark"], "VpnSslWebPortal-UserGroupBookmark"); ok {
			if err = d.Set("user_group_bookmark", vv); err != nil {
				return fmt.Errorf("Error reading user_group_bookmark: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_group_bookmark: %v", err)
		}
	}

	if err = d.Set("web_mode", flattenVpnSslWebPortalWebMode(o["web-mode"], d, "web_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-mode"], "VpnSslWebPortal-WebMode"); ok {
			if err = d.Set("web_mode", vv); err != nil {
				return fmt.Errorf("Error reading web_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_mode: %v", err)
		}
	}

	if err = d.Set("windows_forticlient_download_url", flattenVpnSslWebPortalWindowsForticlientDownloadUrl(o["windows-forticlient-download-url"], d, "windows_forticlient_download_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["windows-forticlient-download-url"], "VpnSslWebPortal-WindowsForticlientDownloadUrl"); ok {
			if err = d.Set("windows_forticlient_download_url", vv); err != nil {
				return fmt.Errorf("Error reading windows_forticlient_download_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading windows_forticlient_download_url: %v", err)
		}
	}

	if err = d.Set("wins_server1", flattenVpnSslWebPortalWinsServer1(o["wins-server1"], d, "wins_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["wins-server1"], "VpnSslWebPortal-WinsServer1"); ok {
			if err = d.Set("wins_server1", vv); err != nil {
				return fmt.Errorf("Error reading wins_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wins_server1: %v", err)
		}
	}

	if err = d.Set("wins_server2", flattenVpnSslWebPortalWinsServer2(o["wins-server2"], d, "wins_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["wins-server2"], "VpnSslWebPortal-WinsServer2"); ok {
			if err = d.Set("wins_server2", vv); err != nil {
				return fmt.Errorf("Error reading wins_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wins_server2: %v", err)
		}
	}

	return nil
}

func flattenVpnSslWebPortalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnSslWebPortalAllowUserAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslWebPortalAutoConnect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bookmarks"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandVpnSslWebPortalBookmarkGroupBookmarks(d, i["bookmarks"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["bookmarks"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandVpnSslWebPortalBookmarkGroupName(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarks(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_params"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["additional-params"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksAdditionalParams(d, i["additional_params"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apptype"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["apptype"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksApptype(d, i["apptype"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color_depth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["color-depth"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksColorDepth(d, i["color_depth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["domain"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksDomain(d, i["domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "folder"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["folder"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksFolder(d, i["folder"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "form_data"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandVpnSslWebPortalBookmarkGroupBookmarksFormData(d, i["form_data"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["form-data"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "height"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["height"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksHeight(d, i["height"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["host"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksHost(d, i["host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keyboard_layout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["keyboard-layout"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksKeyboardLayout(d, i["keyboard_layout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listening_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["listening-port"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksListeningPort(d, i["listening_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "load_balancing_info"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["load-balancing-info"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksLoadBalancingInfo(d, i["load_balancing_info"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["logon-password"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksLogonPassword(d, i["logon_password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_user"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["logon-user"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksLogonUser(d, i["logon_user"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_blob"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preconnection-blob"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksPreconnectionBlob(d, i["preconnection_blob"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preconnection-id"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksPreconnectionId(d, i["preconnection_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remote-port"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksRemotePort(d, i["remote_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restricted_admin"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["restricted-admin"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksRestrictedAdmin(d, i["restricted_admin"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["security"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksSecurity(d, i["security"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_preconnection_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["send-preconnection-id"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksSendPreconnectionId(d, i["send_preconnection_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_layout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server-layout"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksServerLayout(d, i["server_layout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "show_status_window"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["show-status-window"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksShowStatusWindow(d, i["show_status_window"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksSso(d, i["sso"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso-credential"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksSsoCredential(d, i["sso_credential"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential_sent_once"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso-credential-sent-once"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksSsoCredentialSentOnce(d, i["sso_credential_sent_once"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso-password"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksSsoPassword(d, i["sso_password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_username"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso-username"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksSsoUsername(d, i["sso_username"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["url"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksUrl(d, i["url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vnc_keyboard_layout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vnc-keyboard-layout"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksVncKeyboardLayout(d, i["vnc_keyboard_layout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "width"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["width"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksWidth(d, i["width"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksAdditionalParams(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksApptype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksColorDepth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksFolder(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksFormData(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksFormDataName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksFormDataValue(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksFormDataName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksFormDataValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksHeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksKeyboardLayout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksListeningPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksLoadBalancingInfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksLogonPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksLogonUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksPreconnectionBlob(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksPreconnectionId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksRemotePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksRestrictedAdmin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksSecurity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksSendPreconnectionId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksServerLayout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksShowStatusWindow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksSso(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksSsoCredential(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksSsoCredentialSentOnce(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksSsoPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksSsoUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksVncKeyboardLayout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksWidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalClientSrcRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalClipboard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalCustomLang(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslWebPortalCustomizeForticlientDownloadUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalDefaultProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalDefaultWindowHeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalDefaultWindowWidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalDhcpIpOverlap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalDhcpRaGiaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalDhcpReservation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalDhcp6RaLinkaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalDisplayBookmark(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalDisplayConnectionTools(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalDisplayHistory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalDisplayStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalDnsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalDnsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalDnsSuffix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalExclusiveRouting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalFocusBookmark(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalForticlientDownload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalForticlientDownloadMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalHeading(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalHideSsoCredential(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalHostCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalHostCheckInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalHostCheckPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslWebPortalIpMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalIpPools(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslWebPortalIpv6DnsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalIpv6DnsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalIpv6ExclusiveRouting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalIpv6Pools(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslWebPortalIpv6ServiceRestriction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalIpv6SplitTunneling(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalIpv6SplitTunnelingRoutingAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslWebPortalIpv6SplitTunnelingRoutingNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalIpv6TunnelMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalIpv6WinsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalIpv6WinsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalKeepAlive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalLandingPage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "form_data"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandVpnSslWebPortalLandingPageFormData(d, i["form_data"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["form-data"] = t
		}
	}
	pre_append = pre + ".0." + "sso"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sso"], _ = expandVpnSslWebPortalLandingPageSso(d, i["sso"], pre_append)
	}
	pre_append = pre + ".0." + "sso_credential"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sso-credential"], _ = expandVpnSslWebPortalLandingPageSsoCredential(d, i["sso_credential"], pre_append)
	}
	pre_append = pre + ".0." + "sso_password"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sso-password"], _ = expandVpnSslWebPortalLandingPageSsoPassword(d, i["sso_password"], pre_append)
	}
	pre_append = pre + ".0." + "sso_username"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sso-username"], _ = expandVpnSslWebPortalLandingPageSsoUsername(d, i["sso_username"], pre_append)
	}
	pre_append = pre + ".0." + "url"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["url"], _ = expandVpnSslWebPortalLandingPageUrl(d, i["url"], pre_append)
	}

	return result, nil
}

func expandVpnSslWebPortalLandingPageFormData(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandVpnSslWebPortalLandingPageFormDataName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandVpnSslWebPortalLandingPageFormDataValue(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandVpnSslWebPortalLandingPageFormDataName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalLandingPageFormDataValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalLandingPageSso(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalLandingPageSsoCredential(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalLandingPageSsoPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslWebPortalLandingPageSsoUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalLandingPageUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalLandingPageMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalLimitUserLogins(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalMacAddrAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalMacAddrCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalMacAddrCheckRule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_addr_list"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac-addr-list"], _ = expandVpnSslWebPortalMacAddrCheckRuleMacAddrList(d, i["mac_addr_list"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_addr_mask"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac-addr-mask"], _ = expandVpnSslWebPortalMacAddrCheckRuleMacAddrMask(d, i["mac_addr_mask"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandVpnSslWebPortalMacAddrCheckRuleName(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandVpnSslWebPortalMacAddrCheckRuleMacAddrList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslWebPortalMacAddrCheckRuleMacAddrMask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalMacAddrCheckRuleName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalMacosForticlientDownloadUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalOsCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalOsCheckList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandVpnSslWebPortalOsCheckListAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "latest_patch_level"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["latest-patch-level"], _ = expandVpnSslWebPortalOsCheckListLatestPatchLevel(d, i["latest_patch_level"], pre_append)
	}
	pre_append = pre + ".0." + "minor_version"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["minor-version"], _ = expandVpnSslWebPortalOsCheckListMinorVersion(d, i["minor_version"], pre_append)
	}
	pre_append = pre + ".0." + "name"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["name"], _ = expandVpnSslWebPortalOsCheckListName(d, i["name"], pre_append)
	}
	pre_append = pre + ".0." + "tolerance"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tolerance"], _ = expandVpnSslWebPortalOsCheckListTolerance(d, i["tolerance"], pre_append)
	}

	return result, nil
}

func expandVpnSslWebPortalOsCheckListAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalOsCheckListLatestPatchLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalOsCheckListMinorVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalOsCheckListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalOsCheckListTolerance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalPreferIpv6Dns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalRedirUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalRewriteIpUriUi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSavePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalServiceRestriction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSkipCheckForBrowser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSkipCheckForUnsupportedOs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSmbMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSmbMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSmbNtlmv1Auth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSplitDns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_server1"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dns-server1"], _ = expandVpnSslWebPortalSplitDnsDnsServer1(d, i["dns_server1"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_server2"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dns-server2"], _ = expandVpnSslWebPortalSplitDnsDnsServer2(d, i["dns_server2"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domains"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["domains"], _ = expandVpnSslWebPortalSplitDnsDomains(d, i["domains"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandVpnSslWebPortalSplitDnsId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_dns_server1"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipv6-dns-server1"], _ = expandVpnSslWebPortalSplitDnsIpv6DnsServer1(d, i["ipv6_dns_server1"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_dns_server2"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipv6-dns-server2"], _ = expandVpnSslWebPortalSplitDnsIpv6DnsServer2(d, i["ipv6_dns_server2"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandVpnSslWebPortalSplitDnsDnsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSplitDnsDnsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSplitDnsDomains(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSplitDnsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSplitDnsIpv6DnsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSplitDnsIpv6DnsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSplitTunneling(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSplitTunnelingRoutingAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslWebPortalSplitTunnelingRoutingNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalTheme(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalTunnelMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalUseSdwan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalUserBookmark(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalUserGroupBookmark(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalWebMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalWindowsForticlientDownloadUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalWinsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalWinsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnSslWebPortal(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("allow_user_access"); ok || d.HasChange("allow_user_access") {
		t, err := expandVpnSslWebPortalAllowUserAccess(d, v, "allow_user_access")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-user-access"] = t
		}
	}

	if v, ok := d.GetOk("auto_connect"); ok || d.HasChange("auto_connect") {
		t, err := expandVpnSslWebPortalAutoConnect(d, v, "auto_connect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-connect"] = t
		}
	}

	if v, ok := d.GetOk("bookmark_group"); ok || d.HasChange("bookmark_group") {
		t, err := expandVpnSslWebPortalBookmarkGroup(d, v, "bookmark_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bookmark-group"] = t
		}
	}

	if v, ok := d.GetOk("client_src_range"); ok || d.HasChange("client_src_range") {
		t, err := expandVpnSslWebPortalClientSrcRange(d, v, "client_src_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-src-range"] = t
		}
	}

	if v, ok := d.GetOk("clipboard"); ok || d.HasChange("clipboard") {
		t, err := expandVpnSslWebPortalClipboard(d, v, "clipboard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clipboard"] = t
		}
	}

	if v, ok := d.GetOk("custom_lang"); ok || d.HasChange("custom_lang") {
		t, err := expandVpnSslWebPortalCustomLang(d, v, "custom_lang")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-lang"] = t
		}
	}

	if v, ok := d.GetOk("customize_forticlient_download_url"); ok || d.HasChange("customize_forticlient_download_url") {
		t, err := expandVpnSslWebPortalCustomizeForticlientDownloadUrl(d, v, "customize_forticlient_download_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["customize-forticlient-download-url"] = t
		}
	}

	if v, ok := d.GetOk("default_protocol"); ok || d.HasChange("default_protocol") {
		t, err := expandVpnSslWebPortalDefaultProtocol(d, v, "default_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-protocol"] = t
		}
	}

	if v, ok := d.GetOk("default_window_height"); ok || d.HasChange("default_window_height") {
		t, err := expandVpnSslWebPortalDefaultWindowHeight(d, v, "default_window_height")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-window-height"] = t
		}
	}

	if v, ok := d.GetOk("default_window_width"); ok || d.HasChange("default_window_width") {
		t, err := expandVpnSslWebPortalDefaultWindowWidth(d, v, "default_window_width")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-window-width"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_ip_overlap"); ok || d.HasChange("dhcp_ip_overlap") {
		t, err := expandVpnSslWebPortalDhcpIpOverlap(d, v, "dhcp_ip_overlap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-ip-overlap"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_ra_giaddr"); ok || d.HasChange("dhcp_ra_giaddr") {
		t, err := expandVpnSslWebPortalDhcpRaGiaddr(d, v, "dhcp_ra_giaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-ra-giaddr"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_reservation"); ok || d.HasChange("dhcp_reservation") {
		t, err := expandVpnSslWebPortalDhcpReservation(d, v, "dhcp_reservation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-reservation"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_ra_linkaddr"); ok || d.HasChange("dhcp6_ra_linkaddr") {
		t, err := expandVpnSslWebPortalDhcp6RaLinkaddr(d, v, "dhcp6_ra_linkaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-ra-linkaddr"] = t
		}
	}

	if v, ok := d.GetOk("display_bookmark"); ok || d.HasChange("display_bookmark") {
		t, err := expandVpnSslWebPortalDisplayBookmark(d, v, "display_bookmark")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["display-bookmark"] = t
		}
	}

	if v, ok := d.GetOk("display_connection_tools"); ok || d.HasChange("display_connection_tools") {
		t, err := expandVpnSslWebPortalDisplayConnectionTools(d, v, "display_connection_tools")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["display-connection-tools"] = t
		}
	}

	if v, ok := d.GetOk("display_history"); ok || d.HasChange("display_history") {
		t, err := expandVpnSslWebPortalDisplayHistory(d, v, "display_history")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["display-history"] = t
		}
	}

	if v, ok := d.GetOk("display_status"); ok || d.HasChange("display_status") {
		t, err := expandVpnSslWebPortalDisplayStatus(d, v, "display_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["display-status"] = t
		}
	}

	if v, ok := d.GetOk("dns_server1"); ok || d.HasChange("dns_server1") {
		t, err := expandVpnSslWebPortalDnsServer1(d, v, "dns_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("dns_server2"); ok || d.HasChange("dns_server2") {
		t, err := expandVpnSslWebPortalDnsServer2(d, v, "dns_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("dns_suffix"); ok || d.HasChange("dns_suffix") {
		t, err := expandVpnSslWebPortalDnsSuffix(d, v, "dns_suffix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-suffix"] = t
		}
	}

	if v, ok := d.GetOk("exclusive_routing"); ok || d.HasChange("exclusive_routing") {
		t, err := expandVpnSslWebPortalExclusiveRouting(d, v, "exclusive_routing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exclusive-routing"] = t
		}
	}

	if v, ok := d.GetOk("focus_bookmark"); ok || d.HasChange("focus_bookmark") {
		t, err := expandVpnSslWebPortalFocusBookmark(d, v, "focus_bookmark")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["focus-bookmark"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_download"); ok || d.HasChange("forticlient_download") {
		t, err := expandVpnSslWebPortalForticlientDownload(d, v, "forticlient_download")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-download"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_download_method"); ok || d.HasChange("forticlient_download_method") {
		t, err := expandVpnSslWebPortalForticlientDownloadMethod(d, v, "forticlient_download_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-download-method"] = t
		}
	}

	if v, ok := d.GetOk("heading"); ok || d.HasChange("heading") {
		t, err := expandVpnSslWebPortalHeading(d, v, "heading")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["heading"] = t
		}
	}

	if v, ok := d.GetOk("hide_sso_credential"); ok || d.HasChange("hide_sso_credential") {
		t, err := expandVpnSslWebPortalHideSsoCredential(d, v, "hide_sso_credential")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hide-sso-credential"] = t
		}
	}

	if v, ok := d.GetOk("host_check"); ok || d.HasChange("host_check") {
		t, err := expandVpnSslWebPortalHostCheck(d, v, "host_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-check"] = t
		}
	}

	if v, ok := d.GetOk("host_check_interval"); ok || d.HasChange("host_check_interval") {
		t, err := expandVpnSslWebPortalHostCheckInterval(d, v, "host_check_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-check-interval"] = t
		}
	}

	if v, ok := d.GetOk("host_check_policy"); ok || d.HasChange("host_check_policy") {
		t, err := expandVpnSslWebPortalHostCheckPolicy(d, v, "host_check_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-check-policy"] = t
		}
	}

	if v, ok := d.GetOk("ip_mode"); ok || d.HasChange("ip_mode") {
		t, err := expandVpnSslWebPortalIpMode(d, v, "ip_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-mode"] = t
		}
	}

	if v, ok := d.GetOk("ip_pools"); ok || d.HasChange("ip_pools") {
		t, err := expandVpnSslWebPortalIpPools(d, v, "ip_pools")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-pools"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_dns_server1"); ok || d.HasChange("ipv6_dns_server1") {
		t, err := expandVpnSslWebPortalIpv6DnsServer1(d, v, "ipv6_dns_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_dns_server2"); ok || d.HasChange("ipv6_dns_server2") {
		t, err := expandVpnSslWebPortalIpv6DnsServer2(d, v, "ipv6_dns_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_exclusive_routing"); ok || d.HasChange("ipv6_exclusive_routing") {
		t, err := expandVpnSslWebPortalIpv6ExclusiveRouting(d, v, "ipv6_exclusive_routing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-exclusive-routing"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_pools"); ok || d.HasChange("ipv6_pools") {
		t, err := expandVpnSslWebPortalIpv6Pools(d, v, "ipv6_pools")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-pools"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_service_restriction"); ok || d.HasChange("ipv6_service_restriction") {
		t, err := expandVpnSslWebPortalIpv6ServiceRestriction(d, v, "ipv6_service_restriction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-service-restriction"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_split_tunneling"); ok || d.HasChange("ipv6_split_tunneling") {
		t, err := expandVpnSslWebPortalIpv6SplitTunneling(d, v, "ipv6_split_tunneling")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-split-tunneling"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_split_tunneling_routing_address"); ok || d.HasChange("ipv6_split_tunneling_routing_address") {
		t, err := expandVpnSslWebPortalIpv6SplitTunnelingRoutingAddress(d, v, "ipv6_split_tunneling_routing_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-split-tunneling-routing-address"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_split_tunneling_routing_negate"); ok || d.HasChange("ipv6_split_tunneling_routing_negate") {
		t, err := expandVpnSslWebPortalIpv6SplitTunnelingRoutingNegate(d, v, "ipv6_split_tunneling_routing_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-split-tunneling-routing-negate"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_tunnel_mode"); ok || d.HasChange("ipv6_tunnel_mode") {
		t, err := expandVpnSslWebPortalIpv6TunnelMode(d, v, "ipv6_tunnel_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-tunnel-mode"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_wins_server1"); ok || d.HasChange("ipv6_wins_server1") {
		t, err := expandVpnSslWebPortalIpv6WinsServer1(d, v, "ipv6_wins_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-wins-server1"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_wins_server2"); ok || d.HasChange("ipv6_wins_server2") {
		t, err := expandVpnSslWebPortalIpv6WinsServer2(d, v, "ipv6_wins_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-wins-server2"] = t
		}
	}

	if v, ok := d.GetOk("keep_alive"); ok || d.HasChange("keep_alive") {
		t, err := expandVpnSslWebPortalKeepAlive(d, v, "keep_alive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keep-alive"] = t
		}
	}

	if v, ok := d.GetOk("landing_page"); ok || d.HasChange("landing_page") {
		t, err := expandVpnSslWebPortalLandingPage(d, v, "landing_page")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["landing-page"] = t
		}
	}

	if v, ok := d.GetOk("landing_page_mode"); ok || d.HasChange("landing_page_mode") {
		t, err := expandVpnSslWebPortalLandingPageMode(d, v, "landing_page_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["landing-page-mode"] = t
		}
	}

	if v, ok := d.GetOk("limit_user_logins"); ok || d.HasChange("limit_user_logins") {
		t, err := expandVpnSslWebPortalLimitUserLogins(d, v, "limit_user_logins")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["limit-user-logins"] = t
		}
	}

	if v, ok := d.GetOk("mac_addr_action"); ok || d.HasChange("mac_addr_action") {
		t, err := expandVpnSslWebPortalMacAddrAction(d, v, "mac_addr_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-addr-action"] = t
		}
	}

	if v, ok := d.GetOk("mac_addr_check"); ok || d.HasChange("mac_addr_check") {
		t, err := expandVpnSslWebPortalMacAddrCheck(d, v, "mac_addr_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-addr-check"] = t
		}
	}

	if v, ok := d.GetOk("mac_addr_check_rule"); ok || d.HasChange("mac_addr_check_rule") {
		t, err := expandVpnSslWebPortalMacAddrCheckRule(d, v, "mac_addr_check_rule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-addr-check-rule"] = t
		}
	}

	if v, ok := d.GetOk("macos_forticlient_download_url"); ok || d.HasChange("macos_forticlient_download_url") {
		t, err := expandVpnSslWebPortalMacosForticlientDownloadUrl(d, v, "macos_forticlient_download_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["macos-forticlient-download-url"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandVpnSslWebPortalName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("os_check"); ok || d.HasChange("os_check") {
		t, err := expandVpnSslWebPortalOsCheck(d, v, "os_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os-check"] = t
		}
	}

	if v, ok := d.GetOk("os_check_list"); ok || d.HasChange("os_check_list") {
		t, err := expandVpnSslWebPortalOsCheckList(d, v, "os_check_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os-check-list"] = t
		}
	}

	if v, ok := d.GetOk("prefer_ipv6_dns"); ok || d.HasChange("prefer_ipv6_dns") {
		t, err := expandVpnSslWebPortalPreferIpv6Dns(d, v, "prefer_ipv6_dns")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefer-ipv6-dns"] = t
		}
	}

	if v, ok := d.GetOk("redir_url"); ok || d.HasChange("redir_url") {
		t, err := expandVpnSslWebPortalRedirUrl(d, v, "redir_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redir-url"] = t
		}
	}

	if v, ok := d.GetOk("rewrite_ip_uri_ui"); ok || d.HasChange("rewrite_ip_uri_ui") {
		t, err := expandVpnSslWebPortalRewriteIpUriUi(d, v, "rewrite_ip_uri_ui")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rewrite-ip-uri-ui"] = t
		}
	}

	if v, ok := d.GetOk("save_password"); ok || d.HasChange("save_password") {
		t, err := expandVpnSslWebPortalSavePassword(d, v, "save_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["save-password"] = t
		}
	}

	if v, ok := d.GetOk("service_restriction"); ok || d.HasChange("service_restriction") {
		t, err := expandVpnSslWebPortalServiceRestriction(d, v, "service_restriction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-restriction"] = t
		}
	}

	if v, ok := d.GetOk("skip_check_for_browser"); ok || d.HasChange("skip_check_for_browser") {
		t, err := expandVpnSslWebPortalSkipCheckForBrowser(d, v, "skip_check_for_browser")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["skip-check-for-browser"] = t
		}
	}

	if v, ok := d.GetOk("skip_check_for_unsupported_os"); ok || d.HasChange("skip_check_for_unsupported_os") {
		t, err := expandVpnSslWebPortalSkipCheckForUnsupportedOs(d, v, "skip_check_for_unsupported_os")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["skip-check-for-unsupported-os"] = t
		}
	}

	if v, ok := d.GetOk("smb_max_version"); ok || d.HasChange("smb_max_version") {
		t, err := expandVpnSslWebPortalSmbMaxVersion(d, v, "smb_max_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smb-max-version"] = t
		}
	}

	if v, ok := d.GetOk("smb_min_version"); ok || d.HasChange("smb_min_version") {
		t, err := expandVpnSslWebPortalSmbMinVersion(d, v, "smb_min_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smb-min-version"] = t
		}
	}

	if v, ok := d.GetOk("smb_ntlmv1_auth"); ok || d.HasChange("smb_ntlmv1_auth") {
		t, err := expandVpnSslWebPortalSmbNtlmv1Auth(d, v, "smb_ntlmv1_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smb-ntlmv1-auth"] = t
		}
	}

	if v, ok := d.GetOk("split_dns"); ok || d.HasChange("split_dns") {
		t, err := expandVpnSslWebPortalSplitDns(d, v, "split_dns")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-dns"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling"); ok || d.HasChange("split_tunneling") {
		t, err := expandVpnSslWebPortalSplitTunneling(d, v, "split_tunneling")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling_routing_address"); ok || d.HasChange("split_tunneling_routing_address") {
		t, err := expandVpnSslWebPortalSplitTunnelingRoutingAddress(d, v, "split_tunneling_routing_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling-routing-address"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling_routing_negate"); ok || d.HasChange("split_tunneling_routing_negate") {
		t, err := expandVpnSslWebPortalSplitTunnelingRoutingNegate(d, v, "split_tunneling_routing_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling-routing-negate"] = t
		}
	}

	if v, ok := d.GetOk("theme"); ok || d.HasChange("theme") {
		t, err := expandVpnSslWebPortalTheme(d, v, "theme")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["theme"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_mode"); ok || d.HasChange("tunnel_mode") {
		t, err := expandVpnSslWebPortalTunnelMode(d, v, "tunnel_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-mode"] = t
		}
	}

	if v, ok := d.GetOk("use_sdwan"); ok || d.HasChange("use_sdwan") {
		t, err := expandVpnSslWebPortalUseSdwan(d, v, "use_sdwan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-sdwan"] = t
		}
	}

	if v, ok := d.GetOk("user_bookmark"); ok || d.HasChange("user_bookmark") {
		t, err := expandVpnSslWebPortalUserBookmark(d, v, "user_bookmark")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-bookmark"] = t
		}
	}

	if v, ok := d.GetOk("user_group_bookmark"); ok || d.HasChange("user_group_bookmark") {
		t, err := expandVpnSslWebPortalUserGroupBookmark(d, v, "user_group_bookmark")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-group-bookmark"] = t
		}
	}

	if v, ok := d.GetOk("web_mode"); ok || d.HasChange("web_mode") {
		t, err := expandVpnSslWebPortalWebMode(d, v, "web_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-mode"] = t
		}
	}

	if v, ok := d.GetOk("windows_forticlient_download_url"); ok || d.HasChange("windows_forticlient_download_url") {
		t, err := expandVpnSslWebPortalWindowsForticlientDownloadUrl(d, v, "windows_forticlient_download_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["windows-forticlient-download-url"] = t
		}
	}

	if v, ok := d.GetOk("wins_server1"); ok || d.HasChange("wins_server1") {
		t, err := expandVpnSslWebPortalWinsServer1(d, v, "wins_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wins-server1"] = t
		}
	}

	if v, ok := d.GetOk("wins_server2"); ok || d.HasChange("wins_server2") {
		t, err := expandVpnSslWebPortalWinsServer2(d, v, "wins_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wins-server2"] = t
		}
	}

	return &obj, nil
}
