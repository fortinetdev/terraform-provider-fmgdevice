// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Bookmark table.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnSslWebUserGroupBookmarkBookmarks() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnSslWebUserGroupBookmarkBookmarksCreate,
		Read:   resourceVpnSslWebUserGroupBookmarkBookmarksRead,
		Update: resourceVpnSslWebUserGroupBookmarkBookmarksUpdate,
		Delete: resourceVpnSslWebUserGroupBookmarkBookmarksDelete,

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
			"user_group_bookmark": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
				ForceNew: true,
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceVpnSslWebUserGroupBookmarkBookmarksCreate(d *schema.ResourceData, m interface{}) error {
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
	user_group_bookmark := d.Get("user_group_bookmark").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["user_group_bookmark"] = user_group_bookmark

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectVpnSslWebUserGroupBookmarkBookmarks(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnSslWebUserGroupBookmarkBookmarks resource while getting object: %v", err)
	}

	_, err = c.CreateVpnSslWebUserGroupBookmarkBookmarks(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating VpnSslWebUserGroupBookmarkBookmarks resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceVpnSslWebUserGroupBookmarkBookmarksRead(d, m)
}

func resourceVpnSslWebUserGroupBookmarkBookmarksUpdate(d *schema.ResourceData, m interface{}) error {
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
	user_group_bookmark := d.Get("user_group_bookmark").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["user_group_bookmark"] = user_group_bookmark

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectVpnSslWebUserGroupBookmarkBookmarks(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebUserGroupBookmarkBookmarks resource while getting object: %v", err)
	}

	_, err = c.UpdateVpnSslWebUserGroupBookmarkBookmarks(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebUserGroupBookmarkBookmarks resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceVpnSslWebUserGroupBookmarkBookmarksRead(d, m)
}

func resourceVpnSslWebUserGroupBookmarkBookmarksDelete(d *schema.ResourceData, m interface{}) error {
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
	user_group_bookmark := d.Get("user_group_bookmark").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["user_group_bookmark"] = user_group_bookmark

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteVpnSslWebUserGroupBookmarkBookmarks(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VpnSslWebUserGroupBookmarkBookmarks resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnSslWebUserGroupBookmarkBookmarksRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	user_group_bookmark := d.Get("user_group_bookmark").(string)
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
	if user_group_bookmark == "" {
		user_group_bookmark = importOptionChecking(m.(*FortiClient).Cfg, "user_group_bookmark")
		if user_group_bookmark == "" {
			return fmt.Errorf("Parameter user_group_bookmark is missing")
		}
		if err = d.Set("user_group_bookmark", user_group_bookmark); err != nil {
			return fmt.Errorf("Error set params user_group_bookmark: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["user_group_bookmark"] = user_group_bookmark

	o, err := c.ReadVpnSslWebUserGroupBookmarkBookmarks(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebUserGroupBookmarkBookmarks resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnSslWebUserGroupBookmarkBookmarks(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebUserGroupBookmarkBookmarks resource from API: %v", err)
	}
	return nil
}

func flattenVpnSslWebUserGroupBookmarkBookmarksAdditionalParams2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksApptype2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksColorDepth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksDomain2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksFolder2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksFormData2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenVpnSslWebUserGroupBookmarkBookmarksFormDataName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmarkBookmarks-FormData-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenVpnSslWebUserGroupBookmarkBookmarksFormDataValue2edl(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "VpnSslWebUserGroupBookmarkBookmarks-FormData-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenVpnSslWebUserGroupBookmarkBookmarksFormDataName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksFormDataValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksHeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksHost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksKeyboardLayout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksListeningPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksLoadBalancingInfo2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksLogonUser2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksPreconnectionBlob2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksPreconnectionId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksRemotePort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksRestrictedAdmin2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksSecurity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksSendPreconnectionId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksServerLayout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksShowStatusWindow2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksSso2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksSsoCredential2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksSsoCredentialSentOnce2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksSsoUsername2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksVncKeyboardLayout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksWidth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnSslWebUserGroupBookmarkBookmarks(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("additional_params", flattenVpnSslWebUserGroupBookmarkBookmarksAdditionalParams2edl(o["additional-params"], d, "additional_params")); err != nil {
		if vv, ok := fortiAPIPatch(o["additional-params"], "VpnSslWebUserGroupBookmarkBookmarks-AdditionalParams"); ok {
			if err = d.Set("additional_params", vv); err != nil {
				return fmt.Errorf("Error reading additional_params: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading additional_params: %v", err)
		}
	}

	if err = d.Set("apptype", flattenVpnSslWebUserGroupBookmarkBookmarksApptype2edl(o["apptype"], d, "apptype")); err != nil {
		if vv, ok := fortiAPIPatch(o["apptype"], "VpnSslWebUserGroupBookmarkBookmarks-Apptype"); ok {
			if err = d.Set("apptype", vv); err != nil {
				return fmt.Errorf("Error reading apptype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apptype: %v", err)
		}
	}

	if err = d.Set("color_depth", flattenVpnSslWebUserGroupBookmarkBookmarksColorDepth2edl(o["color-depth"], d, "color_depth")); err != nil {
		if vv, ok := fortiAPIPatch(o["color-depth"], "VpnSslWebUserGroupBookmarkBookmarks-ColorDepth"); ok {
			if err = d.Set("color_depth", vv); err != nil {
				return fmt.Errorf("Error reading color_depth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color_depth: %v", err)
		}
	}

	if err = d.Set("description", flattenVpnSslWebUserGroupBookmarkBookmarksDescription2edl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "VpnSslWebUserGroupBookmarkBookmarks-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("domain", flattenVpnSslWebUserGroupBookmarkBookmarksDomain2edl(o["domain"], d, "domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain"], "VpnSslWebUserGroupBookmarkBookmarks-Domain"); ok {
			if err = d.Set("domain", vv); err != nil {
				return fmt.Errorf("Error reading domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("folder", flattenVpnSslWebUserGroupBookmarkBookmarksFolder2edl(o["folder"], d, "folder")); err != nil {
		if vv, ok := fortiAPIPatch(o["folder"], "VpnSslWebUserGroupBookmarkBookmarks-Folder"); ok {
			if err = d.Set("folder", vv); err != nil {
				return fmt.Errorf("Error reading folder: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading folder: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("form_data", flattenVpnSslWebUserGroupBookmarkBookmarksFormData2edl(o["form-data"], d, "form_data")); err != nil {
			if vv, ok := fortiAPIPatch(o["form-data"], "VpnSslWebUserGroupBookmarkBookmarks-FormData"); ok {
				if err = d.Set("form_data", vv); err != nil {
					return fmt.Errorf("Error reading form_data: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading form_data: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("form_data"); ok {
			if err = d.Set("form_data", flattenVpnSslWebUserGroupBookmarkBookmarksFormData2edl(o["form-data"], d, "form_data")); err != nil {
				if vv, ok := fortiAPIPatch(o["form-data"], "VpnSslWebUserGroupBookmarkBookmarks-FormData"); ok {
					if err = d.Set("form_data", vv); err != nil {
						return fmt.Errorf("Error reading form_data: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading form_data: %v", err)
				}
			}
		}
	}

	if err = d.Set("height", flattenVpnSslWebUserGroupBookmarkBookmarksHeight2edl(o["height"], d, "height")); err != nil {
		if vv, ok := fortiAPIPatch(o["height"], "VpnSslWebUserGroupBookmarkBookmarks-Height"); ok {
			if err = d.Set("height", vv); err != nil {
				return fmt.Errorf("Error reading height: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading height: %v", err)
		}
	}

	if err = d.Set("host", flattenVpnSslWebUserGroupBookmarkBookmarksHost2edl(o["host"], d, "host")); err != nil {
		if vv, ok := fortiAPIPatch(o["host"], "VpnSslWebUserGroupBookmarkBookmarks-Host"); ok {
			if err = d.Set("host", vv); err != nil {
				return fmt.Errorf("Error reading host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("keyboard_layout", flattenVpnSslWebUserGroupBookmarkBookmarksKeyboardLayout2edl(o["keyboard-layout"], d, "keyboard_layout")); err != nil {
		if vv, ok := fortiAPIPatch(o["keyboard-layout"], "VpnSslWebUserGroupBookmarkBookmarks-KeyboardLayout"); ok {
			if err = d.Set("keyboard_layout", vv); err != nil {
				return fmt.Errorf("Error reading keyboard_layout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keyboard_layout: %v", err)
		}
	}

	if err = d.Set("listening_port", flattenVpnSslWebUserGroupBookmarkBookmarksListeningPort2edl(o["listening-port"], d, "listening_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["listening-port"], "VpnSslWebUserGroupBookmarkBookmarks-ListeningPort"); ok {
			if err = d.Set("listening_port", vv); err != nil {
				return fmt.Errorf("Error reading listening_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading listening_port: %v", err)
		}
	}

	if err = d.Set("load_balancing_info", flattenVpnSslWebUserGroupBookmarkBookmarksLoadBalancingInfo2edl(o["load-balancing-info"], d, "load_balancing_info")); err != nil {
		if vv, ok := fortiAPIPatch(o["load-balancing-info"], "VpnSslWebUserGroupBookmarkBookmarks-LoadBalancingInfo"); ok {
			if err = d.Set("load_balancing_info", vv); err != nil {
				return fmt.Errorf("Error reading load_balancing_info: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading load_balancing_info: %v", err)
		}
	}

	if err = d.Set("logon_user", flattenVpnSslWebUserGroupBookmarkBookmarksLogonUser2edl(o["logon-user"], d, "logon_user")); err != nil {
		if vv, ok := fortiAPIPatch(o["logon-user"], "VpnSslWebUserGroupBookmarkBookmarks-LogonUser"); ok {
			if err = d.Set("logon_user", vv); err != nil {
				return fmt.Errorf("Error reading logon_user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logon_user: %v", err)
		}
	}

	if err = d.Set("name", flattenVpnSslWebUserGroupBookmarkBookmarksName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "VpnSslWebUserGroupBookmarkBookmarks-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenVpnSslWebUserGroupBookmarkBookmarksPort2edl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "VpnSslWebUserGroupBookmarkBookmarks-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("preconnection_blob", flattenVpnSslWebUserGroupBookmarkBookmarksPreconnectionBlob2edl(o["preconnection-blob"], d, "preconnection_blob")); err != nil {
		if vv, ok := fortiAPIPatch(o["preconnection-blob"], "VpnSslWebUserGroupBookmarkBookmarks-PreconnectionBlob"); ok {
			if err = d.Set("preconnection_blob", vv); err != nil {
				return fmt.Errorf("Error reading preconnection_blob: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preconnection_blob: %v", err)
		}
	}

	if err = d.Set("preconnection_id", flattenVpnSslWebUserGroupBookmarkBookmarksPreconnectionId2edl(o["preconnection-id"], d, "preconnection_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["preconnection-id"], "VpnSslWebUserGroupBookmarkBookmarks-PreconnectionId"); ok {
			if err = d.Set("preconnection_id", vv); err != nil {
				return fmt.Errorf("Error reading preconnection_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preconnection_id: %v", err)
		}
	}

	if err = d.Set("remote_port", flattenVpnSslWebUserGroupBookmarkBookmarksRemotePort2edl(o["remote-port"], d, "remote_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-port"], "VpnSslWebUserGroupBookmarkBookmarks-RemotePort"); ok {
			if err = d.Set("remote_port", vv); err != nil {
				return fmt.Errorf("Error reading remote_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_port: %v", err)
		}
	}

	if err = d.Set("restricted_admin", flattenVpnSslWebUserGroupBookmarkBookmarksRestrictedAdmin2edl(o["restricted-admin"], d, "restricted_admin")); err != nil {
		if vv, ok := fortiAPIPatch(o["restricted-admin"], "VpnSslWebUserGroupBookmarkBookmarks-RestrictedAdmin"); ok {
			if err = d.Set("restricted_admin", vv); err != nil {
				return fmt.Errorf("Error reading restricted_admin: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restricted_admin: %v", err)
		}
	}

	if err = d.Set("security", flattenVpnSslWebUserGroupBookmarkBookmarksSecurity2edl(o["security"], d, "security")); err != nil {
		if vv, ok := fortiAPIPatch(o["security"], "VpnSslWebUserGroupBookmarkBookmarks-Security"); ok {
			if err = d.Set("security", vv); err != nil {
				return fmt.Errorf("Error reading security: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security: %v", err)
		}
	}

	if err = d.Set("send_preconnection_id", flattenVpnSslWebUserGroupBookmarkBookmarksSendPreconnectionId2edl(o["send-preconnection-id"], d, "send_preconnection_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-preconnection-id"], "VpnSslWebUserGroupBookmarkBookmarks-SendPreconnectionId"); ok {
			if err = d.Set("send_preconnection_id", vv); err != nil {
				return fmt.Errorf("Error reading send_preconnection_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_preconnection_id: %v", err)
		}
	}

	if err = d.Set("server_layout", flattenVpnSslWebUserGroupBookmarkBookmarksServerLayout2edl(o["server-layout"], d, "server_layout")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-layout"], "VpnSslWebUserGroupBookmarkBookmarks-ServerLayout"); ok {
			if err = d.Set("server_layout", vv); err != nil {
				return fmt.Errorf("Error reading server_layout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_layout: %v", err)
		}
	}

	if err = d.Set("show_status_window", flattenVpnSslWebUserGroupBookmarkBookmarksShowStatusWindow2edl(o["show-status-window"], d, "show_status_window")); err != nil {
		if vv, ok := fortiAPIPatch(o["show-status-window"], "VpnSslWebUserGroupBookmarkBookmarks-ShowStatusWindow"); ok {
			if err = d.Set("show_status_window", vv); err != nil {
				return fmt.Errorf("Error reading show_status_window: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading show_status_window: %v", err)
		}
	}

	if err = d.Set("sso", flattenVpnSslWebUserGroupBookmarkBookmarksSso2edl(o["sso"], d, "sso")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso"], "VpnSslWebUserGroupBookmarkBookmarks-Sso"); ok {
			if err = d.Set("sso", vv); err != nil {
				return fmt.Errorf("Error reading sso: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso: %v", err)
		}
	}

	if err = d.Set("sso_credential", flattenVpnSslWebUserGroupBookmarkBookmarksSsoCredential2edl(o["sso-credential"], d, "sso_credential")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-credential"], "VpnSslWebUserGroupBookmarkBookmarks-SsoCredential"); ok {
			if err = d.Set("sso_credential", vv); err != nil {
				return fmt.Errorf("Error reading sso_credential: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_credential: %v", err)
		}
	}

	if err = d.Set("sso_credential_sent_once", flattenVpnSslWebUserGroupBookmarkBookmarksSsoCredentialSentOnce2edl(o["sso-credential-sent-once"], d, "sso_credential_sent_once")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-credential-sent-once"], "VpnSslWebUserGroupBookmarkBookmarks-SsoCredentialSentOnce"); ok {
			if err = d.Set("sso_credential_sent_once", vv); err != nil {
				return fmt.Errorf("Error reading sso_credential_sent_once: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_credential_sent_once: %v", err)
		}
	}

	if err = d.Set("sso_username", flattenVpnSslWebUserGroupBookmarkBookmarksSsoUsername2edl(o["sso-username"], d, "sso_username")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-username"], "VpnSslWebUserGroupBookmarkBookmarks-SsoUsername"); ok {
			if err = d.Set("sso_username", vv); err != nil {
				return fmt.Errorf("Error reading sso_username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_username: %v", err)
		}
	}

	if err = d.Set("url", flattenVpnSslWebUserGroupBookmarkBookmarksUrl2edl(o["url"], d, "url")); err != nil {
		if vv, ok := fortiAPIPatch(o["url"], "VpnSslWebUserGroupBookmarkBookmarks-Url"); ok {
			if err = d.Set("url", vv); err != nil {
				return fmt.Errorf("Error reading url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}

	if err = d.Set("vnc_keyboard_layout", flattenVpnSslWebUserGroupBookmarkBookmarksVncKeyboardLayout2edl(o["vnc-keyboard-layout"], d, "vnc_keyboard_layout")); err != nil {
		if vv, ok := fortiAPIPatch(o["vnc-keyboard-layout"], "VpnSslWebUserGroupBookmarkBookmarks-VncKeyboardLayout"); ok {
			if err = d.Set("vnc_keyboard_layout", vv); err != nil {
				return fmt.Errorf("Error reading vnc_keyboard_layout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vnc_keyboard_layout: %v", err)
		}
	}

	if err = d.Set("width", flattenVpnSslWebUserGroupBookmarkBookmarksWidth2edl(o["width"], d, "width")); err != nil {
		if vv, ok := fortiAPIPatch(o["width"], "VpnSslWebUserGroupBookmarkBookmarks-Width"); ok {
			if err = d.Set("width", vv); err != nil {
				return fmt.Errorf("Error reading width: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading width: %v", err)
		}
	}

	return nil
}

func flattenVpnSslWebUserGroupBookmarkBookmarksFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnSslWebUserGroupBookmarkBookmarksAdditionalParams2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksApptype2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksColorDepth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksDomain2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksFolder2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksFormData2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandVpnSslWebUserGroupBookmarkBookmarksFormDataName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandVpnSslWebUserGroupBookmarkBookmarksFormDataValue2edl(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksFormDataName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksFormDataValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksHeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksHost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksKeyboardLayout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksListeningPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksLoadBalancingInfo2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksLogonPassword2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksLogonUser2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksPreconnectionBlob2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksPreconnectionId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksRemotePort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksRestrictedAdmin2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksSecurity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksSendPreconnectionId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksServerLayout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksShowStatusWindow2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksSso2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksSsoCredential2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksSsoCredentialSentOnce2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksSsoPassword2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksSsoUsername2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksVncKeyboardLayout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksWidth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnSslWebUserGroupBookmarkBookmarks(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("additional_params"); ok || d.HasChange("additional_params") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksAdditionalParams2edl(d, v, "additional_params")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-params"] = t
		}
	}

	if v, ok := d.GetOk("apptype"); ok || d.HasChange("apptype") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksApptype2edl(d, v, "apptype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apptype"] = t
		}
	}

	if v, ok := d.GetOk("color_depth"); ok || d.HasChange("color_depth") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksColorDepth2edl(d, v, "color_depth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color-depth"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksDescription2edl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok || d.HasChange("domain") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksDomain2edl(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("folder"); ok || d.HasChange("folder") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksFolder2edl(d, v, "folder")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["folder"] = t
		}
	}

	if v, ok := d.GetOk("form_data"); ok || d.HasChange("form_data") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksFormData2edl(d, v, "form_data")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["form-data"] = t
		}
	}

	if v, ok := d.GetOk("height"); ok || d.HasChange("height") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksHeight2edl(d, v, "height")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["height"] = t
		}
	}

	if v, ok := d.GetOk("host"); ok || d.HasChange("host") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksHost2edl(d, v, "host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}

	if v, ok := d.GetOk("keyboard_layout"); ok || d.HasChange("keyboard_layout") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksKeyboardLayout2edl(d, v, "keyboard_layout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keyboard-layout"] = t
		}
	}

	if v, ok := d.GetOk("listening_port"); ok || d.HasChange("listening_port") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksListeningPort2edl(d, v, "listening_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["listening-port"] = t
		}
	}

	if v, ok := d.GetOk("load_balancing_info"); ok || d.HasChange("load_balancing_info") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksLoadBalancingInfo2edl(d, v, "load_balancing_info")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["load-balancing-info"] = t
		}
	}

	if v, ok := d.GetOk("logon_password"); ok || d.HasChange("logon_password") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksLogonPassword2edl(d, v, "logon_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logon-password"] = t
		}
	}

	if v, ok := d.GetOk("logon_user"); ok || d.HasChange("logon_user") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksLogonUser2edl(d, v, "logon_user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logon-user"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksPort2edl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("preconnection_blob"); ok || d.HasChange("preconnection_blob") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksPreconnectionBlob2edl(d, v, "preconnection_blob")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preconnection-blob"] = t
		}
	}

	if v, ok := d.GetOk("preconnection_id"); ok || d.HasChange("preconnection_id") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksPreconnectionId2edl(d, v, "preconnection_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preconnection-id"] = t
		}
	}

	if v, ok := d.GetOk("remote_port"); ok || d.HasChange("remote_port") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksRemotePort2edl(d, v, "remote_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-port"] = t
		}
	}

	if v, ok := d.GetOk("restricted_admin"); ok || d.HasChange("restricted_admin") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksRestrictedAdmin2edl(d, v, "restricted_admin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restricted-admin"] = t
		}
	}

	if v, ok := d.GetOk("security"); ok || d.HasChange("security") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksSecurity2edl(d, v, "security")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security"] = t
		}
	}

	if v, ok := d.GetOk("send_preconnection_id"); ok || d.HasChange("send_preconnection_id") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksSendPreconnectionId2edl(d, v, "send_preconnection_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-preconnection-id"] = t
		}
	}

	if v, ok := d.GetOk("server_layout"); ok || d.HasChange("server_layout") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksServerLayout2edl(d, v, "server_layout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-layout"] = t
		}
	}

	if v, ok := d.GetOk("show_status_window"); ok || d.HasChange("show_status_window") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksShowStatusWindow2edl(d, v, "show_status_window")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["show-status-window"] = t
		}
	}

	if v, ok := d.GetOk("sso"); ok || d.HasChange("sso") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksSso2edl(d, v, "sso")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso"] = t
		}
	}

	if v, ok := d.GetOk("sso_credential"); ok || d.HasChange("sso_credential") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksSsoCredential2edl(d, v, "sso_credential")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-credential"] = t
		}
	}

	if v, ok := d.GetOk("sso_credential_sent_once"); ok || d.HasChange("sso_credential_sent_once") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksSsoCredentialSentOnce2edl(d, v, "sso_credential_sent_once")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-credential-sent-once"] = t
		}
	}

	if v, ok := d.GetOk("sso_password"); ok || d.HasChange("sso_password") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksSsoPassword2edl(d, v, "sso_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-password"] = t
		}
	}

	if v, ok := d.GetOk("sso_username"); ok || d.HasChange("sso_username") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksSsoUsername2edl(d, v, "sso_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-username"] = t
		}
	}

	if v, ok := d.GetOk("url"); ok || d.HasChange("url") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksUrl2edl(d, v, "url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}

	if v, ok := d.GetOk("vnc_keyboard_layout"); ok || d.HasChange("vnc_keyboard_layout") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksVncKeyboardLayout2edl(d, v, "vnc_keyboard_layout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vnc-keyboard-layout"] = t
		}
	}

	if v, ok := d.GetOk("width"); ok || d.HasChange("width") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksWidth2edl(d, v, "width")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["width"] = t
		}
	}

	return &obj, nil
}
