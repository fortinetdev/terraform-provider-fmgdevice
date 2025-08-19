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

func resourceZtnaWebPortalBookmarkBookmarks() *schema.Resource {
	return &schema.Resource{
		Create: resourceZtnaWebPortalBookmarkBookmarksCreate,
		Read:   resourceZtnaWebPortalBookmarkBookmarksRead,
		Update: resourceZtnaWebPortalBookmarkBookmarksUpdate,
		Delete: resourceZtnaWebPortalBookmarkBookmarksDelete,

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
			"web_portal_bookmark": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
			"restricted_admin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"send_preconnection_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sso": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
	}
}

func resourceZtnaWebPortalBookmarkBookmarksCreate(d *schema.ResourceData, m interface{}) error {
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
	web_portal_bookmark := d.Get("web_portal_bookmark").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["web_portal_bookmark"] = web_portal_bookmark

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectZtnaWebPortalBookmarkBookmarks(d)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaWebPortalBookmarkBookmarks resource while getting object: %v", err)
	}

	_, err = c.CreateZtnaWebPortalBookmarkBookmarks(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaWebPortalBookmarkBookmarks resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceZtnaWebPortalBookmarkBookmarksRead(d, m)
}

func resourceZtnaWebPortalBookmarkBookmarksUpdate(d *schema.ResourceData, m interface{}) error {
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
	web_portal_bookmark := d.Get("web_portal_bookmark").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["web_portal_bookmark"] = web_portal_bookmark

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectZtnaWebPortalBookmarkBookmarks(d)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaWebPortalBookmarkBookmarks resource while getting object: %v", err)
	}

	_, err = c.UpdateZtnaWebPortalBookmarkBookmarks(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaWebPortalBookmarkBookmarks resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceZtnaWebPortalBookmarkBookmarksRead(d, m)
}

func resourceZtnaWebPortalBookmarkBookmarksDelete(d *schema.ResourceData, m interface{}) error {
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
	web_portal_bookmark := d.Get("web_portal_bookmark").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["web_portal_bookmark"] = web_portal_bookmark

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteZtnaWebPortalBookmarkBookmarks(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ZtnaWebPortalBookmarkBookmarks resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceZtnaWebPortalBookmarkBookmarksRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	web_portal_bookmark := d.Get("web_portal_bookmark").(string)
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
	if web_portal_bookmark == "" {
		web_portal_bookmark = importOptionChecking(m.(*FortiClient).Cfg, "web_portal_bookmark")
		if web_portal_bookmark == "" {
			return fmt.Errorf("Parameter web_portal_bookmark is missing")
		}
		if err = d.Set("web_portal_bookmark", web_portal_bookmark); err != nil {
			return fmt.Errorf("Error set params web_portal_bookmark: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["web_portal_bookmark"] = web_portal_bookmark

	o, err := c.ReadZtnaWebPortalBookmarkBookmarks(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaWebPortalBookmarkBookmarks resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectZtnaWebPortalBookmarkBookmarks(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaWebPortalBookmarkBookmarks resource from API: %v", err)
	}
	return nil
}

func flattenZtnaWebPortalBookmarkBookmarksApptype2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksColorDepth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksDomain2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksFolder2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksHeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksHost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksKeyboardLayout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksLoadBalancingInfo2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksLogonUser2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksPreconnectionBlob2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksPreconnectionId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksRestrictedAdmin2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksSecurity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksSendPreconnectionId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksSso2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksVncKeyboardLayout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksWidth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectZtnaWebPortalBookmarkBookmarks(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("apptype", flattenZtnaWebPortalBookmarkBookmarksApptype2edl(o["apptype"], d, "apptype")); err != nil {
		if vv, ok := fortiAPIPatch(o["apptype"], "ZtnaWebPortalBookmarkBookmarks-Apptype"); ok {
			if err = d.Set("apptype", vv); err != nil {
				return fmt.Errorf("Error reading apptype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apptype: %v", err)
		}
	}

	if err = d.Set("color_depth", flattenZtnaWebPortalBookmarkBookmarksColorDepth2edl(o["color-depth"], d, "color_depth")); err != nil {
		if vv, ok := fortiAPIPatch(o["color-depth"], "ZtnaWebPortalBookmarkBookmarks-ColorDepth"); ok {
			if err = d.Set("color_depth", vv); err != nil {
				return fmt.Errorf("Error reading color_depth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color_depth: %v", err)
		}
	}

	if err = d.Set("description", flattenZtnaWebPortalBookmarkBookmarksDescription2edl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ZtnaWebPortalBookmarkBookmarks-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("domain", flattenZtnaWebPortalBookmarkBookmarksDomain2edl(o["domain"], d, "domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain"], "ZtnaWebPortalBookmarkBookmarks-Domain"); ok {
			if err = d.Set("domain", vv); err != nil {
				return fmt.Errorf("Error reading domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("folder", flattenZtnaWebPortalBookmarkBookmarksFolder2edl(o["folder"], d, "folder")); err != nil {
		if vv, ok := fortiAPIPatch(o["folder"], "ZtnaWebPortalBookmarkBookmarks-Folder"); ok {
			if err = d.Set("folder", vv); err != nil {
				return fmt.Errorf("Error reading folder: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading folder: %v", err)
		}
	}

	if err = d.Set("height", flattenZtnaWebPortalBookmarkBookmarksHeight2edl(o["height"], d, "height")); err != nil {
		if vv, ok := fortiAPIPatch(o["height"], "ZtnaWebPortalBookmarkBookmarks-Height"); ok {
			if err = d.Set("height", vv); err != nil {
				return fmt.Errorf("Error reading height: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading height: %v", err)
		}
	}

	if err = d.Set("host", flattenZtnaWebPortalBookmarkBookmarksHost2edl(o["host"], d, "host")); err != nil {
		if vv, ok := fortiAPIPatch(o["host"], "ZtnaWebPortalBookmarkBookmarks-Host"); ok {
			if err = d.Set("host", vv); err != nil {
				return fmt.Errorf("Error reading host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("keyboard_layout", flattenZtnaWebPortalBookmarkBookmarksKeyboardLayout2edl(o["keyboard-layout"], d, "keyboard_layout")); err != nil {
		if vv, ok := fortiAPIPatch(o["keyboard-layout"], "ZtnaWebPortalBookmarkBookmarks-KeyboardLayout"); ok {
			if err = d.Set("keyboard_layout", vv); err != nil {
				return fmt.Errorf("Error reading keyboard_layout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keyboard_layout: %v", err)
		}
	}

	if err = d.Set("load_balancing_info", flattenZtnaWebPortalBookmarkBookmarksLoadBalancingInfo2edl(o["load-balancing-info"], d, "load_balancing_info")); err != nil {
		if vv, ok := fortiAPIPatch(o["load-balancing-info"], "ZtnaWebPortalBookmarkBookmarks-LoadBalancingInfo"); ok {
			if err = d.Set("load_balancing_info", vv); err != nil {
				return fmt.Errorf("Error reading load_balancing_info: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading load_balancing_info: %v", err)
		}
	}

	if err = d.Set("logon_user", flattenZtnaWebPortalBookmarkBookmarksLogonUser2edl(o["logon-user"], d, "logon_user")); err != nil {
		if vv, ok := fortiAPIPatch(o["logon-user"], "ZtnaWebPortalBookmarkBookmarks-LogonUser"); ok {
			if err = d.Set("logon_user", vv); err != nil {
				return fmt.Errorf("Error reading logon_user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logon_user: %v", err)
		}
	}

	if err = d.Set("name", flattenZtnaWebPortalBookmarkBookmarksName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ZtnaWebPortalBookmarkBookmarks-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenZtnaWebPortalBookmarkBookmarksPort2edl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ZtnaWebPortalBookmarkBookmarks-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("preconnection_blob", flattenZtnaWebPortalBookmarkBookmarksPreconnectionBlob2edl(o["preconnection-blob"], d, "preconnection_blob")); err != nil {
		if vv, ok := fortiAPIPatch(o["preconnection-blob"], "ZtnaWebPortalBookmarkBookmarks-PreconnectionBlob"); ok {
			if err = d.Set("preconnection_blob", vv); err != nil {
				return fmt.Errorf("Error reading preconnection_blob: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preconnection_blob: %v", err)
		}
	}

	if err = d.Set("preconnection_id", flattenZtnaWebPortalBookmarkBookmarksPreconnectionId2edl(o["preconnection-id"], d, "preconnection_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["preconnection-id"], "ZtnaWebPortalBookmarkBookmarks-PreconnectionId"); ok {
			if err = d.Set("preconnection_id", vv); err != nil {
				return fmt.Errorf("Error reading preconnection_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preconnection_id: %v", err)
		}
	}

	if err = d.Set("restricted_admin", flattenZtnaWebPortalBookmarkBookmarksRestrictedAdmin2edl(o["restricted-admin"], d, "restricted_admin")); err != nil {
		if vv, ok := fortiAPIPatch(o["restricted-admin"], "ZtnaWebPortalBookmarkBookmarks-RestrictedAdmin"); ok {
			if err = d.Set("restricted_admin", vv); err != nil {
				return fmt.Errorf("Error reading restricted_admin: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restricted_admin: %v", err)
		}
	}

	if err = d.Set("security", flattenZtnaWebPortalBookmarkBookmarksSecurity2edl(o["security"], d, "security")); err != nil {
		if vv, ok := fortiAPIPatch(o["security"], "ZtnaWebPortalBookmarkBookmarks-Security"); ok {
			if err = d.Set("security", vv); err != nil {
				return fmt.Errorf("Error reading security: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security: %v", err)
		}
	}

	if err = d.Set("send_preconnection_id", flattenZtnaWebPortalBookmarkBookmarksSendPreconnectionId2edl(o["send-preconnection-id"], d, "send_preconnection_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-preconnection-id"], "ZtnaWebPortalBookmarkBookmarks-SendPreconnectionId"); ok {
			if err = d.Set("send_preconnection_id", vv); err != nil {
				return fmt.Errorf("Error reading send_preconnection_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_preconnection_id: %v", err)
		}
	}

	if err = d.Set("sso", flattenZtnaWebPortalBookmarkBookmarksSso2edl(o["sso"], d, "sso")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso"], "ZtnaWebPortalBookmarkBookmarks-Sso"); ok {
			if err = d.Set("sso", vv); err != nil {
				return fmt.Errorf("Error reading sso: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso: %v", err)
		}
	}

	if err = d.Set("url", flattenZtnaWebPortalBookmarkBookmarksUrl2edl(o["url"], d, "url")); err != nil {
		if vv, ok := fortiAPIPatch(o["url"], "ZtnaWebPortalBookmarkBookmarks-Url"); ok {
			if err = d.Set("url", vv); err != nil {
				return fmt.Errorf("Error reading url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}

	if err = d.Set("vnc_keyboard_layout", flattenZtnaWebPortalBookmarkBookmarksVncKeyboardLayout2edl(o["vnc-keyboard-layout"], d, "vnc_keyboard_layout")); err != nil {
		if vv, ok := fortiAPIPatch(o["vnc-keyboard-layout"], "ZtnaWebPortalBookmarkBookmarks-VncKeyboardLayout"); ok {
			if err = d.Set("vnc_keyboard_layout", vv); err != nil {
				return fmt.Errorf("Error reading vnc_keyboard_layout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vnc_keyboard_layout: %v", err)
		}
	}

	if err = d.Set("width", flattenZtnaWebPortalBookmarkBookmarksWidth2edl(o["width"], d, "width")); err != nil {
		if vv, ok := fortiAPIPatch(o["width"], "ZtnaWebPortalBookmarkBookmarks-Width"); ok {
			if err = d.Set("width", vv); err != nil {
				return fmt.Errorf("Error reading width: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading width: %v", err)
		}
	}

	return nil
}

func flattenZtnaWebPortalBookmarkBookmarksFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandZtnaWebPortalBookmarkBookmarksApptype2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksColorDepth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksDomain2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksFolder2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksHeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksHost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksKeyboardLayout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksLoadBalancingInfo2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksLogonPassword2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaWebPortalBookmarkBookmarksLogonUser2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksPreconnectionBlob2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksPreconnectionId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksRestrictedAdmin2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksSecurity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksSendPreconnectionId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksSso2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksVncKeyboardLayout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksWidth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectZtnaWebPortalBookmarkBookmarks(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("apptype"); ok || d.HasChange("apptype") {
		t, err := expandZtnaWebPortalBookmarkBookmarksApptype2edl(d, v, "apptype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apptype"] = t
		}
	}

	if v, ok := d.GetOk("color_depth"); ok || d.HasChange("color_depth") {
		t, err := expandZtnaWebPortalBookmarkBookmarksColorDepth2edl(d, v, "color_depth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color-depth"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandZtnaWebPortalBookmarkBookmarksDescription2edl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok || d.HasChange("domain") {
		t, err := expandZtnaWebPortalBookmarkBookmarksDomain2edl(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("folder"); ok || d.HasChange("folder") {
		t, err := expandZtnaWebPortalBookmarkBookmarksFolder2edl(d, v, "folder")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["folder"] = t
		}
	}

	if v, ok := d.GetOk("height"); ok || d.HasChange("height") {
		t, err := expandZtnaWebPortalBookmarkBookmarksHeight2edl(d, v, "height")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["height"] = t
		}
	}

	if v, ok := d.GetOk("host"); ok || d.HasChange("host") {
		t, err := expandZtnaWebPortalBookmarkBookmarksHost2edl(d, v, "host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}

	if v, ok := d.GetOk("keyboard_layout"); ok || d.HasChange("keyboard_layout") {
		t, err := expandZtnaWebPortalBookmarkBookmarksKeyboardLayout2edl(d, v, "keyboard_layout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keyboard-layout"] = t
		}
	}

	if v, ok := d.GetOk("load_balancing_info"); ok || d.HasChange("load_balancing_info") {
		t, err := expandZtnaWebPortalBookmarkBookmarksLoadBalancingInfo2edl(d, v, "load_balancing_info")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["load-balancing-info"] = t
		}
	}

	if v, ok := d.GetOk("logon_password"); ok || d.HasChange("logon_password") {
		t, err := expandZtnaWebPortalBookmarkBookmarksLogonPassword2edl(d, v, "logon_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logon-password"] = t
		}
	}

	if v, ok := d.GetOk("logon_user"); ok || d.HasChange("logon_user") {
		t, err := expandZtnaWebPortalBookmarkBookmarksLogonUser2edl(d, v, "logon_user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logon-user"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandZtnaWebPortalBookmarkBookmarksName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandZtnaWebPortalBookmarkBookmarksPort2edl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("preconnection_blob"); ok || d.HasChange("preconnection_blob") {
		t, err := expandZtnaWebPortalBookmarkBookmarksPreconnectionBlob2edl(d, v, "preconnection_blob")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preconnection-blob"] = t
		}
	}

	if v, ok := d.GetOk("preconnection_id"); ok || d.HasChange("preconnection_id") {
		t, err := expandZtnaWebPortalBookmarkBookmarksPreconnectionId2edl(d, v, "preconnection_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preconnection-id"] = t
		}
	}

	if v, ok := d.GetOk("restricted_admin"); ok || d.HasChange("restricted_admin") {
		t, err := expandZtnaWebPortalBookmarkBookmarksRestrictedAdmin2edl(d, v, "restricted_admin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restricted-admin"] = t
		}
	}

	if v, ok := d.GetOk("security"); ok || d.HasChange("security") {
		t, err := expandZtnaWebPortalBookmarkBookmarksSecurity2edl(d, v, "security")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security"] = t
		}
	}

	if v, ok := d.GetOk("send_preconnection_id"); ok || d.HasChange("send_preconnection_id") {
		t, err := expandZtnaWebPortalBookmarkBookmarksSendPreconnectionId2edl(d, v, "send_preconnection_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-preconnection-id"] = t
		}
	}

	if v, ok := d.GetOk("sso"); ok || d.HasChange("sso") {
		t, err := expandZtnaWebPortalBookmarkBookmarksSso2edl(d, v, "sso")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso"] = t
		}
	}

	if v, ok := d.GetOk("url"); ok || d.HasChange("url") {
		t, err := expandZtnaWebPortalBookmarkBookmarksUrl2edl(d, v, "url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}

	if v, ok := d.GetOk("vnc_keyboard_layout"); ok || d.HasChange("vnc_keyboard_layout") {
		t, err := expandZtnaWebPortalBookmarkBookmarksVncKeyboardLayout2edl(d, v, "vnc_keyboard_layout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vnc-keyboard-layout"] = t
		}
	}

	if v, ok := d.GetOk("width"); ok || d.HasChange("width") {
		t, err := expandZtnaWebPortalBookmarkBookmarksWidth2edl(d, v, "width")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["width"] = t
		}
	}

	return &obj, nil
}
