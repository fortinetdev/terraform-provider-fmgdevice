// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure SSL-VPN user group bookmark.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnSslWebUserGroupBookmark() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnSslWebUserGroupBookmarkCreate,
		Read:   resourceVpnSslWebUserGroupBookmarkRead,
		Update: resourceVpnSslWebUserGroupBookmarkUpdate,
		Delete: resourceVpnSslWebUserGroupBookmarkDelete,

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

func resourceVpnSslWebUserGroupBookmarkCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectVpnSslWebUserGroupBookmark(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnSslWebUserGroupBookmark resource while getting object: %v", err)
	}

	_, err = c.CreateVpnSslWebUserGroupBookmark(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating VpnSslWebUserGroupBookmark resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceVpnSslWebUserGroupBookmarkRead(d, m)
}

func resourceVpnSslWebUserGroupBookmarkUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectVpnSslWebUserGroupBookmark(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebUserGroupBookmark resource while getting object: %v", err)
	}

	_, err = c.UpdateVpnSslWebUserGroupBookmark(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebUserGroupBookmark resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceVpnSslWebUserGroupBookmarkRead(d, m)
}

func resourceVpnSslWebUserGroupBookmarkDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteVpnSslWebUserGroupBookmark(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VpnSslWebUserGroupBookmark resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnSslWebUserGroupBookmarkRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnSslWebUserGroupBookmark(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebUserGroupBookmark resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnSslWebUserGroupBookmark(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebUserGroupBookmark resource from API: %v", err)
	}
	return nil
}

func flattenVpnSslWebUserGroupBookmarkBookmarks(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenVpnSslWebUserGroupBookmarkBookmarksAdditionalParams(i["additional-params"], d, pre_append)
			tmp["additional_params"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmark-Bookmarks-AdditionalParams")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apptype"
		if _, ok := i["apptype"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksApptype(i["apptype"], d, pre_append)
			tmp["apptype"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmark-Bookmarks-Apptype")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color_depth"
		if _, ok := i["color-depth"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksColorDepth(i["color-depth"], d, pre_append)
			tmp["color_depth"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmark-Bookmarks-ColorDepth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmark-Bookmarks-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := i["domain"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksDomain(i["domain"], d, pre_append)
			tmp["domain"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmark-Bookmarks-Domain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "folder"
		if _, ok := i["folder"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksFolder(i["folder"], d, pre_append)
			tmp["folder"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmark-Bookmarks-Folder")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "form_data"
		if _, ok := i["form-data"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksFormData(i["form-data"], d, pre_append)
			tmp["form_data"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmark-Bookmarks-FormData")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "height"
		if _, ok := i["height"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksHeight(i["height"], d, pre_append)
			tmp["height"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmark-Bookmarks-Height")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := i["host"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksHost(i["host"], d, pre_append)
			tmp["host"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmark-Bookmarks-Host")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keyboard_layout"
		if _, ok := i["keyboard-layout"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksKeyboardLayout(i["keyboard-layout"], d, pre_append)
			tmp["keyboard_layout"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmark-Bookmarks-KeyboardLayout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listening_port"
		if _, ok := i["listening-port"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksListeningPort(i["listening-port"], d, pre_append)
			tmp["listening_port"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmark-Bookmarks-ListeningPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "load_balancing_info"
		if _, ok := i["load-balancing-info"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksLoadBalancingInfo(i["load-balancing-info"], d, pre_append)
			tmp["load_balancing_info"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmark-Bookmarks-LoadBalancingInfo")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_user"
		if _, ok := i["logon-user"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksLogonUser(i["logon-user"], d, pre_append)
			tmp["logon_user"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmark-Bookmarks-LogonUser")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmark-Bookmarks-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmark-Bookmarks-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_blob"
		if _, ok := i["preconnection-blob"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksPreconnectionBlob(i["preconnection-blob"], d, pre_append)
			tmp["preconnection_blob"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmark-Bookmarks-PreconnectionBlob")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_id"
		if _, ok := i["preconnection-id"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksPreconnectionId(i["preconnection-id"], d, pre_append)
			tmp["preconnection_id"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmark-Bookmarks-PreconnectionId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_port"
		if _, ok := i["remote-port"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksRemotePort(i["remote-port"], d, pre_append)
			tmp["remote_port"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmark-Bookmarks-RemotePort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restricted_admin"
		if _, ok := i["restricted-admin"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksRestrictedAdmin(i["restricted-admin"], d, pre_append)
			tmp["restricted_admin"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmark-Bookmarks-RestrictedAdmin")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security"
		if _, ok := i["security"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksSecurity(i["security"], d, pre_append)
			tmp["security"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmark-Bookmarks-Security")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_preconnection_id"
		if _, ok := i["send-preconnection-id"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksSendPreconnectionId(i["send-preconnection-id"], d, pre_append)
			tmp["send_preconnection_id"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmark-Bookmarks-SendPreconnectionId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_layout"
		if _, ok := i["server-layout"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksServerLayout(i["server-layout"], d, pre_append)
			tmp["server_layout"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmark-Bookmarks-ServerLayout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "show_status_window"
		if _, ok := i["show-status-window"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksShowStatusWindow(i["show-status-window"], d, pre_append)
			tmp["show_status_window"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmark-Bookmarks-ShowStatusWindow")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso"
		if _, ok := i["sso"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksSso(i["sso"], d, pre_append)
			tmp["sso"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmark-Bookmarks-Sso")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential"
		if _, ok := i["sso-credential"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksSsoCredential(i["sso-credential"], d, pre_append)
			tmp["sso_credential"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmark-Bookmarks-SsoCredential")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential_sent_once"
		if _, ok := i["sso-credential-sent-once"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksSsoCredentialSentOnce(i["sso-credential-sent-once"], d, pre_append)
			tmp["sso_credential_sent_once"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmark-Bookmarks-SsoCredentialSentOnce")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_username"
		if _, ok := i["sso-username"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksSsoUsername(i["sso-username"], d, pre_append)
			tmp["sso_username"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmark-Bookmarks-SsoUsername")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := i["url"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksUrl(i["url"], d, pre_append)
			tmp["url"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmark-Bookmarks-Url")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vnc_keyboard_layout"
		if _, ok := i["vnc-keyboard-layout"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksVncKeyboardLayout(i["vnc-keyboard-layout"], d, pre_append)
			tmp["vnc_keyboard_layout"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmark-Bookmarks-VncKeyboardLayout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "width"
		if _, ok := i["width"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksWidth(i["width"], d, pre_append)
			tmp["width"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmark-Bookmarks-Width")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenVpnSslWebUserGroupBookmarkBookmarksAdditionalParams(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksApptype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksColorDepth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksFolder(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksFormData(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenVpnSslWebUserGroupBookmarkBookmarksFormDataName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmarkBookmarks-FormData-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksFormDataValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmarkBookmarks-FormData-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenVpnSslWebUserGroupBookmarkBookmarksFormDataName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksFormDataValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksHeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksKeyboardLayout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksListeningPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksLoadBalancingInfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksLogonUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksPreconnectionBlob(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksPreconnectionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksRemotePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksRestrictedAdmin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksSecurity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksSendPreconnectionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksServerLayout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksShowStatusWindow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksSso(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksSsoCredential(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksSsoCredentialSentOnce(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksSsoUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksVncKeyboardLayout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksWidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnSslWebUserGroupBookmark(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("bookmarks", flattenVpnSslWebUserGroupBookmarkBookmarks(o["bookmarks"], d, "bookmarks")); err != nil {
			if vv, ok := fortiAPIPatch(o["bookmarks"], "VpnSslWebUserGroupBookmark-Bookmarks"); ok {
				if err = d.Set("bookmarks", vv); err != nil {
					return fmt.Errorf("Error reading bookmarks: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading bookmarks: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("bookmarks"); ok {
			if err = d.Set("bookmarks", flattenVpnSslWebUserGroupBookmarkBookmarks(o["bookmarks"], d, "bookmarks")); err != nil {
				if vv, ok := fortiAPIPatch(o["bookmarks"], "VpnSslWebUserGroupBookmark-Bookmarks"); ok {
					if err = d.Set("bookmarks", vv); err != nil {
						return fmt.Errorf("Error reading bookmarks: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading bookmarks: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenVpnSslWebUserGroupBookmarkName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "VpnSslWebUserGroupBookmark-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenVpnSslWebUserGroupBookmarkFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnSslWebUserGroupBookmarkBookmarks(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["additional-params"], _ = expandVpnSslWebUserGroupBookmarkBookmarksAdditionalParams(d, i["additional_params"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apptype"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["apptype"], _ = expandVpnSslWebUserGroupBookmarkBookmarksApptype(d, i["apptype"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color_depth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["color-depth"], _ = expandVpnSslWebUserGroupBookmarkBookmarksColorDepth(d, i["color_depth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandVpnSslWebUserGroupBookmarkBookmarksDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["domain"], _ = expandVpnSslWebUserGroupBookmarkBookmarksDomain(d, i["domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "folder"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["folder"], _ = expandVpnSslWebUserGroupBookmarkBookmarksFolder(d, i["folder"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "form_data"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandVpnSslWebUserGroupBookmarkBookmarksFormData(d, i["form_data"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["form-data"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "height"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["height"], _ = expandVpnSslWebUserGroupBookmarkBookmarksHeight(d, i["height"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["host"], _ = expandVpnSslWebUserGroupBookmarkBookmarksHost(d, i["host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keyboard_layout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["keyboard-layout"], _ = expandVpnSslWebUserGroupBookmarkBookmarksKeyboardLayout(d, i["keyboard_layout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listening_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["listening-port"], _ = expandVpnSslWebUserGroupBookmarkBookmarksListeningPort(d, i["listening_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "load_balancing_info"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["load-balancing-info"], _ = expandVpnSslWebUserGroupBookmarkBookmarksLoadBalancingInfo(d, i["load_balancing_info"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["logon-password"], _ = expandVpnSslWebUserGroupBookmarkBookmarksLogonPassword(d, i["logon_password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_user"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["logon-user"], _ = expandVpnSslWebUserGroupBookmarkBookmarksLogonUser(d, i["logon_user"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandVpnSslWebUserGroupBookmarkBookmarksName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandVpnSslWebUserGroupBookmarkBookmarksPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_blob"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preconnection-blob"], _ = expandVpnSslWebUserGroupBookmarkBookmarksPreconnectionBlob(d, i["preconnection_blob"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preconnection-id"], _ = expandVpnSslWebUserGroupBookmarkBookmarksPreconnectionId(d, i["preconnection_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remote-port"], _ = expandVpnSslWebUserGroupBookmarkBookmarksRemotePort(d, i["remote_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restricted_admin"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["restricted-admin"], _ = expandVpnSslWebUserGroupBookmarkBookmarksRestrictedAdmin(d, i["restricted_admin"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["security"], _ = expandVpnSslWebUserGroupBookmarkBookmarksSecurity(d, i["security"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_preconnection_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["send-preconnection-id"], _ = expandVpnSslWebUserGroupBookmarkBookmarksSendPreconnectionId(d, i["send_preconnection_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_layout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server-layout"], _ = expandVpnSslWebUserGroupBookmarkBookmarksServerLayout(d, i["server_layout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "show_status_window"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["show-status-window"], _ = expandVpnSslWebUserGroupBookmarkBookmarksShowStatusWindow(d, i["show_status_window"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso"], _ = expandVpnSslWebUserGroupBookmarkBookmarksSso(d, i["sso"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso-credential"], _ = expandVpnSslWebUserGroupBookmarkBookmarksSsoCredential(d, i["sso_credential"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential_sent_once"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso-credential-sent-once"], _ = expandVpnSslWebUserGroupBookmarkBookmarksSsoCredentialSentOnce(d, i["sso_credential_sent_once"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso-password"], _ = expandVpnSslWebUserGroupBookmarkBookmarksSsoPassword(d, i["sso_password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_username"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso-username"], _ = expandVpnSslWebUserGroupBookmarkBookmarksSsoUsername(d, i["sso_username"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["url"], _ = expandVpnSslWebUserGroupBookmarkBookmarksUrl(d, i["url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vnc_keyboard_layout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vnc-keyboard-layout"], _ = expandVpnSslWebUserGroupBookmarkBookmarksVncKeyboardLayout(d, i["vnc_keyboard_layout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "width"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["width"], _ = expandVpnSslWebUserGroupBookmarkBookmarksWidth(d, i["width"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksAdditionalParams(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksApptype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksColorDepth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksFolder(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksFormData(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandVpnSslWebUserGroupBookmarkBookmarksFormDataName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandVpnSslWebUserGroupBookmarkBookmarksFormDataValue(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksFormDataName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksFormDataValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksHeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksKeyboardLayout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksListeningPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksLoadBalancingInfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksLogonPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksLogonUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksPreconnectionBlob(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksPreconnectionId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksRemotePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksRestrictedAdmin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksSecurity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksSendPreconnectionId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksServerLayout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksShowStatusWindow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksSso(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksSsoCredential(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksSsoCredentialSentOnce(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksSsoPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksSsoUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksVncKeyboardLayout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksWidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnSslWebUserGroupBookmark(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bookmarks"); ok || d.HasChange("bookmarks") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarks(d, v, "bookmarks")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bookmarks"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandVpnSslWebUserGroupBookmarkName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
