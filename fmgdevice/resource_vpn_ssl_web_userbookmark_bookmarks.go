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

func resourceVpnSslWebUserBookmarkBookmarks() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnSslWebUserBookmarkBookmarksCreate,
		Read:   resourceVpnSslWebUserBookmarkBookmarksRead,
		Update: resourceVpnSslWebUserBookmarkBookmarksUpdate,
		Delete: resourceVpnSslWebUserBookmarkBookmarksDelete,

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
			"user_bookmark": &schema.Schema{
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

func resourceVpnSslWebUserBookmarkBookmarksCreate(d *schema.ResourceData, m interface{}) error {
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
	user_bookmark := d.Get("user_bookmark").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["user_bookmark"] = user_bookmark

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectVpnSslWebUserBookmarkBookmarks(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnSslWebUserBookmarkBookmarks resource while getting object: %v", err)
	}

	_, err = c.CreateVpnSslWebUserBookmarkBookmarks(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating VpnSslWebUserBookmarkBookmarks resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceVpnSslWebUserBookmarkBookmarksRead(d, m)
}

func resourceVpnSslWebUserBookmarkBookmarksUpdate(d *schema.ResourceData, m interface{}) error {
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
	user_bookmark := d.Get("user_bookmark").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["user_bookmark"] = user_bookmark

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectVpnSslWebUserBookmarkBookmarks(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebUserBookmarkBookmarks resource while getting object: %v", err)
	}

	_, err = c.UpdateVpnSslWebUserBookmarkBookmarks(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebUserBookmarkBookmarks resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceVpnSslWebUserBookmarkBookmarksRead(d, m)
}

func resourceVpnSslWebUserBookmarkBookmarksDelete(d *schema.ResourceData, m interface{}) error {
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
	user_bookmark := d.Get("user_bookmark").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["user_bookmark"] = user_bookmark

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteVpnSslWebUserBookmarkBookmarks(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VpnSslWebUserBookmarkBookmarks resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnSslWebUserBookmarkBookmarksRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	user_bookmark := d.Get("user_bookmark").(string)
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
	if user_bookmark == "" {
		user_bookmark = importOptionChecking(m.(*FortiClient).Cfg, "user_bookmark")
		if user_bookmark == "" {
			return fmt.Errorf("Parameter user_bookmark is missing")
		}
		if err = d.Set("user_bookmark", user_bookmark); err != nil {
			return fmt.Errorf("Error set params user_bookmark: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["user_bookmark"] = user_bookmark

	o, err := c.ReadVpnSslWebUserBookmarkBookmarks(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebUserBookmarkBookmarks resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnSslWebUserBookmarkBookmarks(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebUserBookmarkBookmarks resource from API: %v", err)
	}
	return nil
}

func flattenVpnSslWebUserBookmarkBookmarksAdditionalParams2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksApptype2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksColorDepth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksDomain2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksFolder2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksFormData2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenVpnSslWebUserBookmarkBookmarksFormDataName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmarkBookmarks-FormData-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenVpnSslWebUserBookmarkBookmarksFormDataValue2edl(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "VpnSslWebUserBookmarkBookmarks-FormData-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenVpnSslWebUserBookmarkBookmarksFormDataName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksFormDataValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksHeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksHost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksKeyboardLayout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksListeningPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksLoadBalancingInfo2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksLogonUser2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksPreconnectionBlob2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksPreconnectionId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksRemotePort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksRestrictedAdmin2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksSecurity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksSendPreconnectionId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksServerLayout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksShowStatusWindow2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksSso2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksSsoCredential2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksSsoCredentialSentOnce2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksSsoUsername2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksVncKeyboardLayout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksWidth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnSslWebUserBookmarkBookmarks(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("additional_params", flattenVpnSslWebUserBookmarkBookmarksAdditionalParams2edl(o["additional-params"], d, "additional_params")); err != nil {
		if vv, ok := fortiAPIPatch(o["additional-params"], "VpnSslWebUserBookmarkBookmarks-AdditionalParams"); ok {
			if err = d.Set("additional_params", vv); err != nil {
				return fmt.Errorf("Error reading additional_params: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading additional_params: %v", err)
		}
	}

	if err = d.Set("apptype", flattenVpnSslWebUserBookmarkBookmarksApptype2edl(o["apptype"], d, "apptype")); err != nil {
		if vv, ok := fortiAPIPatch(o["apptype"], "VpnSslWebUserBookmarkBookmarks-Apptype"); ok {
			if err = d.Set("apptype", vv); err != nil {
				return fmt.Errorf("Error reading apptype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apptype: %v", err)
		}
	}

	if err = d.Set("color_depth", flattenVpnSslWebUserBookmarkBookmarksColorDepth2edl(o["color-depth"], d, "color_depth")); err != nil {
		if vv, ok := fortiAPIPatch(o["color-depth"], "VpnSslWebUserBookmarkBookmarks-ColorDepth"); ok {
			if err = d.Set("color_depth", vv); err != nil {
				return fmt.Errorf("Error reading color_depth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color_depth: %v", err)
		}
	}

	if err = d.Set("description", flattenVpnSslWebUserBookmarkBookmarksDescription2edl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "VpnSslWebUserBookmarkBookmarks-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("domain", flattenVpnSslWebUserBookmarkBookmarksDomain2edl(o["domain"], d, "domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain"], "VpnSslWebUserBookmarkBookmarks-Domain"); ok {
			if err = d.Set("domain", vv); err != nil {
				return fmt.Errorf("Error reading domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("folder", flattenVpnSslWebUserBookmarkBookmarksFolder2edl(o["folder"], d, "folder")); err != nil {
		if vv, ok := fortiAPIPatch(o["folder"], "VpnSslWebUserBookmarkBookmarks-Folder"); ok {
			if err = d.Set("folder", vv); err != nil {
				return fmt.Errorf("Error reading folder: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading folder: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("form_data", flattenVpnSslWebUserBookmarkBookmarksFormData2edl(o["form-data"], d, "form_data")); err != nil {
			if vv, ok := fortiAPIPatch(o["form-data"], "VpnSslWebUserBookmarkBookmarks-FormData"); ok {
				if err = d.Set("form_data", vv); err != nil {
					return fmt.Errorf("Error reading form_data: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading form_data: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("form_data"); ok {
			if err = d.Set("form_data", flattenVpnSslWebUserBookmarkBookmarksFormData2edl(o["form-data"], d, "form_data")); err != nil {
				if vv, ok := fortiAPIPatch(o["form-data"], "VpnSslWebUserBookmarkBookmarks-FormData"); ok {
					if err = d.Set("form_data", vv); err != nil {
						return fmt.Errorf("Error reading form_data: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading form_data: %v", err)
				}
			}
		}
	}

	if err = d.Set("height", flattenVpnSslWebUserBookmarkBookmarksHeight2edl(o["height"], d, "height")); err != nil {
		if vv, ok := fortiAPIPatch(o["height"], "VpnSslWebUserBookmarkBookmarks-Height"); ok {
			if err = d.Set("height", vv); err != nil {
				return fmt.Errorf("Error reading height: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading height: %v", err)
		}
	}

	if err = d.Set("host", flattenVpnSslWebUserBookmarkBookmarksHost2edl(o["host"], d, "host")); err != nil {
		if vv, ok := fortiAPIPatch(o["host"], "VpnSslWebUserBookmarkBookmarks-Host"); ok {
			if err = d.Set("host", vv); err != nil {
				return fmt.Errorf("Error reading host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("keyboard_layout", flattenVpnSslWebUserBookmarkBookmarksKeyboardLayout2edl(o["keyboard-layout"], d, "keyboard_layout")); err != nil {
		if vv, ok := fortiAPIPatch(o["keyboard-layout"], "VpnSslWebUserBookmarkBookmarks-KeyboardLayout"); ok {
			if err = d.Set("keyboard_layout", vv); err != nil {
				return fmt.Errorf("Error reading keyboard_layout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keyboard_layout: %v", err)
		}
	}

	if err = d.Set("listening_port", flattenVpnSslWebUserBookmarkBookmarksListeningPort2edl(o["listening-port"], d, "listening_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["listening-port"], "VpnSslWebUserBookmarkBookmarks-ListeningPort"); ok {
			if err = d.Set("listening_port", vv); err != nil {
				return fmt.Errorf("Error reading listening_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading listening_port: %v", err)
		}
	}

	if err = d.Set("load_balancing_info", flattenVpnSslWebUserBookmarkBookmarksLoadBalancingInfo2edl(o["load-balancing-info"], d, "load_balancing_info")); err != nil {
		if vv, ok := fortiAPIPatch(o["load-balancing-info"], "VpnSslWebUserBookmarkBookmarks-LoadBalancingInfo"); ok {
			if err = d.Set("load_balancing_info", vv); err != nil {
				return fmt.Errorf("Error reading load_balancing_info: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading load_balancing_info: %v", err)
		}
	}

	if err = d.Set("logon_user", flattenVpnSslWebUserBookmarkBookmarksLogonUser2edl(o["logon-user"], d, "logon_user")); err != nil {
		if vv, ok := fortiAPIPatch(o["logon-user"], "VpnSslWebUserBookmarkBookmarks-LogonUser"); ok {
			if err = d.Set("logon_user", vv); err != nil {
				return fmt.Errorf("Error reading logon_user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logon_user: %v", err)
		}
	}

	if err = d.Set("name", flattenVpnSslWebUserBookmarkBookmarksName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "VpnSslWebUserBookmarkBookmarks-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenVpnSslWebUserBookmarkBookmarksPort2edl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "VpnSslWebUserBookmarkBookmarks-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("preconnection_blob", flattenVpnSslWebUserBookmarkBookmarksPreconnectionBlob2edl(o["preconnection-blob"], d, "preconnection_blob")); err != nil {
		if vv, ok := fortiAPIPatch(o["preconnection-blob"], "VpnSslWebUserBookmarkBookmarks-PreconnectionBlob"); ok {
			if err = d.Set("preconnection_blob", vv); err != nil {
				return fmt.Errorf("Error reading preconnection_blob: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preconnection_blob: %v", err)
		}
	}

	if err = d.Set("preconnection_id", flattenVpnSslWebUserBookmarkBookmarksPreconnectionId2edl(o["preconnection-id"], d, "preconnection_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["preconnection-id"], "VpnSslWebUserBookmarkBookmarks-PreconnectionId"); ok {
			if err = d.Set("preconnection_id", vv); err != nil {
				return fmt.Errorf("Error reading preconnection_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preconnection_id: %v", err)
		}
	}

	if err = d.Set("remote_port", flattenVpnSslWebUserBookmarkBookmarksRemotePort2edl(o["remote-port"], d, "remote_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-port"], "VpnSslWebUserBookmarkBookmarks-RemotePort"); ok {
			if err = d.Set("remote_port", vv); err != nil {
				return fmt.Errorf("Error reading remote_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_port: %v", err)
		}
	}

	if err = d.Set("restricted_admin", flattenVpnSslWebUserBookmarkBookmarksRestrictedAdmin2edl(o["restricted-admin"], d, "restricted_admin")); err != nil {
		if vv, ok := fortiAPIPatch(o["restricted-admin"], "VpnSslWebUserBookmarkBookmarks-RestrictedAdmin"); ok {
			if err = d.Set("restricted_admin", vv); err != nil {
				return fmt.Errorf("Error reading restricted_admin: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restricted_admin: %v", err)
		}
	}

	if err = d.Set("security", flattenVpnSslWebUserBookmarkBookmarksSecurity2edl(o["security"], d, "security")); err != nil {
		if vv, ok := fortiAPIPatch(o["security"], "VpnSslWebUserBookmarkBookmarks-Security"); ok {
			if err = d.Set("security", vv); err != nil {
				return fmt.Errorf("Error reading security: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security: %v", err)
		}
	}

	if err = d.Set("send_preconnection_id", flattenVpnSslWebUserBookmarkBookmarksSendPreconnectionId2edl(o["send-preconnection-id"], d, "send_preconnection_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-preconnection-id"], "VpnSslWebUserBookmarkBookmarks-SendPreconnectionId"); ok {
			if err = d.Set("send_preconnection_id", vv); err != nil {
				return fmt.Errorf("Error reading send_preconnection_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_preconnection_id: %v", err)
		}
	}

	if err = d.Set("server_layout", flattenVpnSslWebUserBookmarkBookmarksServerLayout2edl(o["server-layout"], d, "server_layout")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-layout"], "VpnSslWebUserBookmarkBookmarks-ServerLayout"); ok {
			if err = d.Set("server_layout", vv); err != nil {
				return fmt.Errorf("Error reading server_layout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_layout: %v", err)
		}
	}

	if err = d.Set("show_status_window", flattenVpnSslWebUserBookmarkBookmarksShowStatusWindow2edl(o["show-status-window"], d, "show_status_window")); err != nil {
		if vv, ok := fortiAPIPatch(o["show-status-window"], "VpnSslWebUserBookmarkBookmarks-ShowStatusWindow"); ok {
			if err = d.Set("show_status_window", vv); err != nil {
				return fmt.Errorf("Error reading show_status_window: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading show_status_window: %v", err)
		}
	}

	if err = d.Set("sso", flattenVpnSslWebUserBookmarkBookmarksSso2edl(o["sso"], d, "sso")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso"], "VpnSslWebUserBookmarkBookmarks-Sso"); ok {
			if err = d.Set("sso", vv); err != nil {
				return fmt.Errorf("Error reading sso: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso: %v", err)
		}
	}

	if err = d.Set("sso_credential", flattenVpnSslWebUserBookmarkBookmarksSsoCredential2edl(o["sso-credential"], d, "sso_credential")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-credential"], "VpnSslWebUserBookmarkBookmarks-SsoCredential"); ok {
			if err = d.Set("sso_credential", vv); err != nil {
				return fmt.Errorf("Error reading sso_credential: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_credential: %v", err)
		}
	}

	if err = d.Set("sso_credential_sent_once", flattenVpnSslWebUserBookmarkBookmarksSsoCredentialSentOnce2edl(o["sso-credential-sent-once"], d, "sso_credential_sent_once")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-credential-sent-once"], "VpnSslWebUserBookmarkBookmarks-SsoCredentialSentOnce"); ok {
			if err = d.Set("sso_credential_sent_once", vv); err != nil {
				return fmt.Errorf("Error reading sso_credential_sent_once: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_credential_sent_once: %v", err)
		}
	}

	if err = d.Set("sso_username", flattenVpnSslWebUserBookmarkBookmarksSsoUsername2edl(o["sso-username"], d, "sso_username")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-username"], "VpnSslWebUserBookmarkBookmarks-SsoUsername"); ok {
			if err = d.Set("sso_username", vv); err != nil {
				return fmt.Errorf("Error reading sso_username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_username: %v", err)
		}
	}

	if err = d.Set("url", flattenVpnSslWebUserBookmarkBookmarksUrl2edl(o["url"], d, "url")); err != nil {
		if vv, ok := fortiAPIPatch(o["url"], "VpnSslWebUserBookmarkBookmarks-Url"); ok {
			if err = d.Set("url", vv); err != nil {
				return fmt.Errorf("Error reading url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}

	if err = d.Set("vnc_keyboard_layout", flattenVpnSslWebUserBookmarkBookmarksVncKeyboardLayout2edl(o["vnc-keyboard-layout"], d, "vnc_keyboard_layout")); err != nil {
		if vv, ok := fortiAPIPatch(o["vnc-keyboard-layout"], "VpnSslWebUserBookmarkBookmarks-VncKeyboardLayout"); ok {
			if err = d.Set("vnc_keyboard_layout", vv); err != nil {
				return fmt.Errorf("Error reading vnc_keyboard_layout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vnc_keyboard_layout: %v", err)
		}
	}

	if err = d.Set("width", flattenVpnSslWebUserBookmarkBookmarksWidth2edl(o["width"], d, "width")); err != nil {
		if vv, ok := fortiAPIPatch(o["width"], "VpnSslWebUserBookmarkBookmarks-Width"); ok {
			if err = d.Set("width", vv); err != nil {
				return fmt.Errorf("Error reading width: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading width: %v", err)
		}
	}

	return nil
}

func flattenVpnSslWebUserBookmarkBookmarksFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnSslWebUserBookmarkBookmarksAdditionalParams2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksApptype2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksColorDepth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksDomain2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksFolder2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksFormData2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandVpnSslWebUserBookmarkBookmarksFormDataName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandVpnSslWebUserBookmarkBookmarksFormDataValue2edl(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandVpnSslWebUserBookmarkBookmarksFormDataName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksFormDataValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksHeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksHost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksKeyboardLayout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksListeningPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksLoadBalancingInfo2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksLogonPassword2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslWebUserBookmarkBookmarksLogonUser2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksPreconnectionBlob2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksPreconnectionId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksRemotePort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksRestrictedAdmin2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksSecurity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksSendPreconnectionId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksServerLayout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksShowStatusWindow2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksSso2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksSsoCredential2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksSsoCredentialSentOnce2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksSsoPassword2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslWebUserBookmarkBookmarksSsoUsername2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksVncKeyboardLayout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksWidth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnSslWebUserBookmarkBookmarks(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("additional_params"); ok || d.HasChange("additional_params") {
		t, err := expandVpnSslWebUserBookmarkBookmarksAdditionalParams2edl(d, v, "additional_params")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-params"] = t
		}
	}

	if v, ok := d.GetOk("apptype"); ok || d.HasChange("apptype") {
		t, err := expandVpnSslWebUserBookmarkBookmarksApptype2edl(d, v, "apptype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apptype"] = t
		}
	}

	if v, ok := d.GetOk("color_depth"); ok || d.HasChange("color_depth") {
		t, err := expandVpnSslWebUserBookmarkBookmarksColorDepth2edl(d, v, "color_depth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color-depth"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandVpnSslWebUserBookmarkBookmarksDescription2edl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok || d.HasChange("domain") {
		t, err := expandVpnSslWebUserBookmarkBookmarksDomain2edl(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("folder"); ok || d.HasChange("folder") {
		t, err := expandVpnSslWebUserBookmarkBookmarksFolder2edl(d, v, "folder")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["folder"] = t
		}
	}

	if v, ok := d.GetOk("form_data"); ok || d.HasChange("form_data") {
		t, err := expandVpnSslWebUserBookmarkBookmarksFormData2edl(d, v, "form_data")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["form-data"] = t
		}
	}

	if v, ok := d.GetOk("height"); ok || d.HasChange("height") {
		t, err := expandVpnSslWebUserBookmarkBookmarksHeight2edl(d, v, "height")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["height"] = t
		}
	}

	if v, ok := d.GetOk("host"); ok || d.HasChange("host") {
		t, err := expandVpnSslWebUserBookmarkBookmarksHost2edl(d, v, "host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}

	if v, ok := d.GetOk("keyboard_layout"); ok || d.HasChange("keyboard_layout") {
		t, err := expandVpnSslWebUserBookmarkBookmarksKeyboardLayout2edl(d, v, "keyboard_layout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keyboard-layout"] = t
		}
	}

	if v, ok := d.GetOk("listening_port"); ok || d.HasChange("listening_port") {
		t, err := expandVpnSslWebUserBookmarkBookmarksListeningPort2edl(d, v, "listening_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["listening-port"] = t
		}
	}

	if v, ok := d.GetOk("load_balancing_info"); ok || d.HasChange("load_balancing_info") {
		t, err := expandVpnSslWebUserBookmarkBookmarksLoadBalancingInfo2edl(d, v, "load_balancing_info")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["load-balancing-info"] = t
		}
	}

	if v, ok := d.GetOk("logon_password"); ok || d.HasChange("logon_password") {
		t, err := expandVpnSslWebUserBookmarkBookmarksLogonPassword2edl(d, v, "logon_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logon-password"] = t
		}
	}

	if v, ok := d.GetOk("logon_user"); ok || d.HasChange("logon_user") {
		t, err := expandVpnSslWebUserBookmarkBookmarksLogonUser2edl(d, v, "logon_user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logon-user"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandVpnSslWebUserBookmarkBookmarksName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandVpnSslWebUserBookmarkBookmarksPort2edl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("preconnection_blob"); ok || d.HasChange("preconnection_blob") {
		t, err := expandVpnSslWebUserBookmarkBookmarksPreconnectionBlob2edl(d, v, "preconnection_blob")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preconnection-blob"] = t
		}
	}

	if v, ok := d.GetOk("preconnection_id"); ok || d.HasChange("preconnection_id") {
		t, err := expandVpnSslWebUserBookmarkBookmarksPreconnectionId2edl(d, v, "preconnection_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preconnection-id"] = t
		}
	}

	if v, ok := d.GetOk("remote_port"); ok || d.HasChange("remote_port") {
		t, err := expandVpnSslWebUserBookmarkBookmarksRemotePort2edl(d, v, "remote_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-port"] = t
		}
	}

	if v, ok := d.GetOk("restricted_admin"); ok || d.HasChange("restricted_admin") {
		t, err := expandVpnSslWebUserBookmarkBookmarksRestrictedAdmin2edl(d, v, "restricted_admin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restricted-admin"] = t
		}
	}

	if v, ok := d.GetOk("security"); ok || d.HasChange("security") {
		t, err := expandVpnSslWebUserBookmarkBookmarksSecurity2edl(d, v, "security")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security"] = t
		}
	}

	if v, ok := d.GetOk("send_preconnection_id"); ok || d.HasChange("send_preconnection_id") {
		t, err := expandVpnSslWebUserBookmarkBookmarksSendPreconnectionId2edl(d, v, "send_preconnection_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-preconnection-id"] = t
		}
	}

	if v, ok := d.GetOk("server_layout"); ok || d.HasChange("server_layout") {
		t, err := expandVpnSslWebUserBookmarkBookmarksServerLayout2edl(d, v, "server_layout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-layout"] = t
		}
	}

	if v, ok := d.GetOk("show_status_window"); ok || d.HasChange("show_status_window") {
		t, err := expandVpnSslWebUserBookmarkBookmarksShowStatusWindow2edl(d, v, "show_status_window")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["show-status-window"] = t
		}
	}

	if v, ok := d.GetOk("sso"); ok || d.HasChange("sso") {
		t, err := expandVpnSslWebUserBookmarkBookmarksSso2edl(d, v, "sso")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso"] = t
		}
	}

	if v, ok := d.GetOk("sso_credential"); ok || d.HasChange("sso_credential") {
		t, err := expandVpnSslWebUserBookmarkBookmarksSsoCredential2edl(d, v, "sso_credential")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-credential"] = t
		}
	}

	if v, ok := d.GetOk("sso_credential_sent_once"); ok || d.HasChange("sso_credential_sent_once") {
		t, err := expandVpnSslWebUserBookmarkBookmarksSsoCredentialSentOnce2edl(d, v, "sso_credential_sent_once")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-credential-sent-once"] = t
		}
	}

	if v, ok := d.GetOk("sso_password"); ok || d.HasChange("sso_password") {
		t, err := expandVpnSslWebUserBookmarkBookmarksSsoPassword2edl(d, v, "sso_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-password"] = t
		}
	}

	if v, ok := d.GetOk("sso_username"); ok || d.HasChange("sso_username") {
		t, err := expandVpnSslWebUserBookmarkBookmarksSsoUsername2edl(d, v, "sso_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-username"] = t
		}
	}

	if v, ok := d.GetOk("url"); ok || d.HasChange("url") {
		t, err := expandVpnSslWebUserBookmarkBookmarksUrl2edl(d, v, "url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}

	if v, ok := d.GetOk("vnc_keyboard_layout"); ok || d.HasChange("vnc_keyboard_layout") {
		t, err := expandVpnSslWebUserBookmarkBookmarksVncKeyboardLayout2edl(d, v, "vnc_keyboard_layout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vnc-keyboard-layout"] = t
		}
	}

	if v, ok := d.GetOk("width"); ok || d.HasChange("width") {
		t, err := expandVpnSslWebUserBookmarkBookmarksWidth2edl(d, v, "width")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["width"] = t
		}
	}

	return &obj, nil
}
