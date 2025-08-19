// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure SSL-VPN user bookmark.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnSslWebUserBookmark() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnSslWebUserBookmarkCreate,
		Read:   resourceVpnSslWebUserBookmarkRead,
		Update: resourceVpnSslWebUserBookmarkUpdate,
		Delete: resourceVpnSslWebUserBookmarkDelete,

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
							Computed: true,
						},
						"color_depth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
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
							Computed: true,
						},
						"security": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"send_preconnection_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
						},
						"sso_credential": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sso_credential_sent_once": &schema.Schema{
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
						"vnc_keyboard_layout": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"width": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"custom_lang": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
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

func resourceVpnSslWebUserBookmarkCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectVpnSslWebUserBookmark(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnSslWebUserBookmark resource while getting object: %v", err)
	}

	_, err = c.CreateVpnSslWebUserBookmark(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating VpnSslWebUserBookmark resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceVpnSslWebUserBookmarkRead(d, m)
}

func resourceVpnSslWebUserBookmarkUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectVpnSslWebUserBookmark(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebUserBookmark resource while getting object: %v", err)
	}

	_, err = c.UpdateVpnSslWebUserBookmark(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebUserBookmark resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceVpnSslWebUserBookmarkRead(d, m)
}

func resourceVpnSslWebUserBookmarkDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteVpnSslWebUserBookmark(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VpnSslWebUserBookmark resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnSslWebUserBookmarkRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnSslWebUserBookmark(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebUserBookmark resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnSslWebUserBookmark(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebUserBookmark resource from API: %v", err)
	}
	return nil
}

func flattenVpnSslWebUserBookmarkBookmarks(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenVpnSslWebUserBookmarkBookmarksAdditionalParams(i["additional-params"], d, pre_append)
			tmp["additional_params"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmark-Bookmarks-AdditionalParams")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apptype"
		if _, ok := i["apptype"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksApptype(i["apptype"], d, pre_append)
			tmp["apptype"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmark-Bookmarks-Apptype")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color_depth"
		if _, ok := i["color-depth"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksColorDepth(i["color-depth"], d, pre_append)
			tmp["color_depth"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmark-Bookmarks-ColorDepth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmark-Bookmarks-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := i["domain"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksDomain(i["domain"], d, pre_append)
			tmp["domain"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmark-Bookmarks-Domain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "folder"
		if _, ok := i["folder"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksFolder(i["folder"], d, pre_append)
			tmp["folder"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmark-Bookmarks-Folder")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "form_data"
		if _, ok := i["form-data"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksFormData(i["form-data"], d, pre_append)
			tmp["form_data"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmark-Bookmarks-FormData")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "height"
		if _, ok := i["height"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksHeight(i["height"], d, pre_append)
			tmp["height"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmark-Bookmarks-Height")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := i["host"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksHost(i["host"], d, pre_append)
			tmp["host"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmark-Bookmarks-Host")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keyboard_layout"
		if _, ok := i["keyboard-layout"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksKeyboardLayout(i["keyboard-layout"], d, pre_append)
			tmp["keyboard_layout"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmark-Bookmarks-KeyboardLayout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listening_port"
		if _, ok := i["listening-port"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksListeningPort(i["listening-port"], d, pre_append)
			tmp["listening_port"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmark-Bookmarks-ListeningPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "load_balancing_info"
		if _, ok := i["load-balancing-info"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksLoadBalancingInfo(i["load-balancing-info"], d, pre_append)
			tmp["load_balancing_info"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmark-Bookmarks-LoadBalancingInfo")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_user"
		if _, ok := i["logon-user"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksLogonUser(i["logon-user"], d, pre_append)
			tmp["logon_user"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmark-Bookmarks-LogonUser")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmark-Bookmarks-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmark-Bookmarks-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_blob"
		if _, ok := i["preconnection-blob"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksPreconnectionBlob(i["preconnection-blob"], d, pre_append)
			tmp["preconnection_blob"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmark-Bookmarks-PreconnectionBlob")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_id"
		if _, ok := i["preconnection-id"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksPreconnectionId(i["preconnection-id"], d, pre_append)
			tmp["preconnection_id"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmark-Bookmarks-PreconnectionId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_port"
		if _, ok := i["remote-port"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksRemotePort(i["remote-port"], d, pre_append)
			tmp["remote_port"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmark-Bookmarks-RemotePort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restricted_admin"
		if _, ok := i["restricted-admin"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksRestrictedAdmin(i["restricted-admin"], d, pre_append)
			tmp["restricted_admin"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmark-Bookmarks-RestrictedAdmin")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security"
		if _, ok := i["security"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksSecurity(i["security"], d, pre_append)
			tmp["security"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmark-Bookmarks-Security")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_preconnection_id"
		if _, ok := i["send-preconnection-id"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksSendPreconnectionId(i["send-preconnection-id"], d, pre_append)
			tmp["send_preconnection_id"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmark-Bookmarks-SendPreconnectionId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_layout"
		if _, ok := i["server-layout"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksServerLayout(i["server-layout"], d, pre_append)
			tmp["server_layout"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmark-Bookmarks-ServerLayout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "show_status_window"
		if _, ok := i["show-status-window"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksShowStatusWindow(i["show-status-window"], d, pre_append)
			tmp["show_status_window"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmark-Bookmarks-ShowStatusWindow")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso"
		if _, ok := i["sso"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksSso(i["sso"], d, pre_append)
			tmp["sso"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmark-Bookmarks-Sso")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential"
		if _, ok := i["sso-credential"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksSsoCredential(i["sso-credential"], d, pre_append)
			tmp["sso_credential"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmark-Bookmarks-SsoCredential")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential_sent_once"
		if _, ok := i["sso-credential-sent-once"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksSsoCredentialSentOnce(i["sso-credential-sent-once"], d, pre_append)
			tmp["sso_credential_sent_once"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmark-Bookmarks-SsoCredentialSentOnce")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_username"
		if _, ok := i["sso-username"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksSsoUsername(i["sso-username"], d, pre_append)
			tmp["sso_username"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmark-Bookmarks-SsoUsername")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := i["url"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksUrl(i["url"], d, pre_append)
			tmp["url"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmark-Bookmarks-Url")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vnc_keyboard_layout"
		if _, ok := i["vnc-keyboard-layout"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksVncKeyboardLayout(i["vnc-keyboard-layout"], d, pre_append)
			tmp["vnc_keyboard_layout"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmark-Bookmarks-VncKeyboardLayout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "width"
		if _, ok := i["width"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksWidth(i["width"], d, pre_append)
			tmp["width"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmark-Bookmarks-Width")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenVpnSslWebUserBookmarkBookmarksAdditionalParams(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksApptype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksColorDepth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksFolder(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksFormData(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenVpnSslWebUserBookmarkBookmarksFormDataName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmarkBookmarks-FormData-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksFormDataValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmarkBookmarks-FormData-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenVpnSslWebUserBookmarkBookmarksFormDataName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksFormDataValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksHeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksKeyboardLayout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksListeningPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksLoadBalancingInfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksLogonUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksPreconnectionBlob(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksPreconnectionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksRemotePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksRestrictedAdmin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksSecurity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksSendPreconnectionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksServerLayout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksShowStatusWindow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksSso(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksSsoCredential(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksSsoCredentialSentOnce(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksSsoUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksVncKeyboardLayout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksWidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkCustomLang(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslWebUserBookmarkName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnSslWebUserBookmark(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("bookmarks", flattenVpnSslWebUserBookmarkBookmarks(o["bookmarks"], d, "bookmarks")); err != nil {
			if vv, ok := fortiAPIPatch(o["bookmarks"], "VpnSslWebUserBookmark-Bookmarks"); ok {
				if err = d.Set("bookmarks", vv); err != nil {
					return fmt.Errorf("Error reading bookmarks: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading bookmarks: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("bookmarks"); ok {
			if err = d.Set("bookmarks", flattenVpnSslWebUserBookmarkBookmarks(o["bookmarks"], d, "bookmarks")); err != nil {
				if vv, ok := fortiAPIPatch(o["bookmarks"], "VpnSslWebUserBookmark-Bookmarks"); ok {
					if err = d.Set("bookmarks", vv); err != nil {
						return fmt.Errorf("Error reading bookmarks: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading bookmarks: %v", err)
				}
			}
		}
	}

	if err = d.Set("custom_lang", flattenVpnSslWebUserBookmarkCustomLang(o["custom-lang"], d, "custom_lang")); err != nil {
		if vv, ok := fortiAPIPatch(o["custom-lang"], "VpnSslWebUserBookmark-CustomLang"); ok {
			if err = d.Set("custom_lang", vv); err != nil {
				return fmt.Errorf("Error reading custom_lang: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading custom_lang: %v", err)
		}
	}

	if err = d.Set("name", flattenVpnSslWebUserBookmarkName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "VpnSslWebUserBookmark-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenVpnSslWebUserBookmarkFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnSslWebUserBookmarkBookmarks(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["additional-params"], _ = expandVpnSslWebUserBookmarkBookmarksAdditionalParams(d, i["additional_params"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apptype"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["apptype"], _ = expandVpnSslWebUserBookmarkBookmarksApptype(d, i["apptype"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color_depth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["color-depth"], _ = expandVpnSslWebUserBookmarkBookmarksColorDepth(d, i["color_depth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandVpnSslWebUserBookmarkBookmarksDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["domain"], _ = expandVpnSslWebUserBookmarkBookmarksDomain(d, i["domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "folder"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["folder"], _ = expandVpnSslWebUserBookmarkBookmarksFolder(d, i["folder"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "form_data"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandVpnSslWebUserBookmarkBookmarksFormData(d, i["form_data"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["form-data"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "height"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["height"], _ = expandVpnSslWebUserBookmarkBookmarksHeight(d, i["height"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["host"], _ = expandVpnSslWebUserBookmarkBookmarksHost(d, i["host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keyboard_layout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["keyboard-layout"], _ = expandVpnSslWebUserBookmarkBookmarksKeyboardLayout(d, i["keyboard_layout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listening_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["listening-port"], _ = expandVpnSslWebUserBookmarkBookmarksListeningPort(d, i["listening_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "load_balancing_info"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["load-balancing-info"], _ = expandVpnSslWebUserBookmarkBookmarksLoadBalancingInfo(d, i["load_balancing_info"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["logon-password"], _ = expandVpnSslWebUserBookmarkBookmarksLogonPassword(d, i["logon_password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_user"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["logon-user"], _ = expandVpnSslWebUserBookmarkBookmarksLogonUser(d, i["logon_user"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandVpnSslWebUserBookmarkBookmarksName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandVpnSslWebUserBookmarkBookmarksPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_blob"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preconnection-blob"], _ = expandVpnSslWebUserBookmarkBookmarksPreconnectionBlob(d, i["preconnection_blob"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preconnection-id"], _ = expandVpnSslWebUserBookmarkBookmarksPreconnectionId(d, i["preconnection_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remote-port"], _ = expandVpnSslWebUserBookmarkBookmarksRemotePort(d, i["remote_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restricted_admin"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["restricted-admin"], _ = expandVpnSslWebUserBookmarkBookmarksRestrictedAdmin(d, i["restricted_admin"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["security"], _ = expandVpnSslWebUserBookmarkBookmarksSecurity(d, i["security"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_preconnection_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["send-preconnection-id"], _ = expandVpnSslWebUserBookmarkBookmarksSendPreconnectionId(d, i["send_preconnection_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_layout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server-layout"], _ = expandVpnSslWebUserBookmarkBookmarksServerLayout(d, i["server_layout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "show_status_window"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["show-status-window"], _ = expandVpnSslWebUserBookmarkBookmarksShowStatusWindow(d, i["show_status_window"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso"], _ = expandVpnSslWebUserBookmarkBookmarksSso(d, i["sso"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso-credential"], _ = expandVpnSslWebUserBookmarkBookmarksSsoCredential(d, i["sso_credential"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential_sent_once"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso-credential-sent-once"], _ = expandVpnSslWebUserBookmarkBookmarksSsoCredentialSentOnce(d, i["sso_credential_sent_once"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso-password"], _ = expandVpnSslWebUserBookmarkBookmarksSsoPassword(d, i["sso_password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_username"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso-username"], _ = expandVpnSslWebUserBookmarkBookmarksSsoUsername(d, i["sso_username"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["url"], _ = expandVpnSslWebUserBookmarkBookmarksUrl(d, i["url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vnc_keyboard_layout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vnc-keyboard-layout"], _ = expandVpnSslWebUserBookmarkBookmarksVncKeyboardLayout(d, i["vnc_keyboard_layout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "width"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["width"], _ = expandVpnSslWebUserBookmarkBookmarksWidth(d, i["width"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandVpnSslWebUserBookmarkBookmarksAdditionalParams(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksApptype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksColorDepth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksFolder(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksFormData(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandVpnSslWebUserBookmarkBookmarksFormDataName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandVpnSslWebUserBookmarkBookmarksFormDataValue(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandVpnSslWebUserBookmarkBookmarksFormDataName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksFormDataValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksHeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksKeyboardLayout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksListeningPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksLoadBalancingInfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksLogonPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslWebUserBookmarkBookmarksLogonUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksPreconnectionBlob(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksPreconnectionId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksRemotePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksRestrictedAdmin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksSecurity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksSendPreconnectionId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksServerLayout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksShowStatusWindow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksSso(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksSsoCredential(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksSsoCredentialSentOnce(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksSsoPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslWebUserBookmarkBookmarksSsoUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksVncKeyboardLayout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksWidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkCustomLang(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslWebUserBookmarkName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnSslWebUserBookmark(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bookmarks"); ok || d.HasChange("bookmarks") {
		t, err := expandVpnSslWebUserBookmarkBookmarks(d, v, "bookmarks")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bookmarks"] = t
		}
	}

	if v, ok := d.GetOk("custom_lang"); ok || d.HasChange("custom_lang") {
		t, err := expandVpnSslWebUserBookmarkCustomLang(d, v, "custom_lang")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-lang"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandVpnSslWebUserBookmarkName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
