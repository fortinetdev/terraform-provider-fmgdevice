// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure ztna web-portal bookmark.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceZtnaWebPortalBookmark() *schema.Resource {
	return &schema.Resource{
		Create: resourceZtnaWebPortalBookmarkCreate,
		Read:   resourceZtnaWebPortalBookmarkRead,
		Update: resourceZtnaWebPortalBookmarkUpdate,
		Delete: resourceZtnaWebPortalBookmarkDelete,

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
				},
			},
			"groups": &schema.Schema{
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
			"users": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceZtnaWebPortalBookmarkCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectZtnaWebPortalBookmark(d)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaWebPortalBookmark resource while getting object: %v", err)
	}

	_, err = c.CreateZtnaWebPortalBookmark(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaWebPortalBookmark resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceZtnaWebPortalBookmarkRead(d, m)
}

func resourceZtnaWebPortalBookmarkUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectZtnaWebPortalBookmark(d)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaWebPortalBookmark resource while getting object: %v", err)
	}

	_, err = c.UpdateZtnaWebPortalBookmark(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaWebPortalBookmark resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceZtnaWebPortalBookmarkRead(d, m)
}

func resourceZtnaWebPortalBookmarkDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteZtnaWebPortalBookmark(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ZtnaWebPortalBookmark resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceZtnaWebPortalBookmarkRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadZtnaWebPortalBookmark(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaWebPortalBookmark resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectZtnaWebPortalBookmark(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaWebPortalBookmark resource from API: %v", err)
	}
	return nil
}

func flattenZtnaWebPortalBookmarkBookmarks(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apptype"
		if _, ok := i["apptype"]; ok {
			v := flattenZtnaWebPortalBookmarkBookmarksApptype(i["apptype"], d, pre_append)
			tmp["apptype"] = fortiAPISubPartPatch(v, "ZtnaWebPortalBookmark-Bookmarks-Apptype")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color_depth"
		if _, ok := i["color-depth"]; ok {
			v := flattenZtnaWebPortalBookmarkBookmarksColorDepth(i["color-depth"], d, pre_append)
			tmp["color_depth"] = fortiAPISubPartPatch(v, "ZtnaWebPortalBookmark-Bookmarks-ColorDepth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenZtnaWebPortalBookmarkBookmarksDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "ZtnaWebPortalBookmark-Bookmarks-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := i["domain"]; ok {
			v := flattenZtnaWebPortalBookmarkBookmarksDomain(i["domain"], d, pre_append)
			tmp["domain"] = fortiAPISubPartPatch(v, "ZtnaWebPortalBookmark-Bookmarks-Domain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "folder"
		if _, ok := i["folder"]; ok {
			v := flattenZtnaWebPortalBookmarkBookmarksFolder(i["folder"], d, pre_append)
			tmp["folder"] = fortiAPISubPartPatch(v, "ZtnaWebPortalBookmark-Bookmarks-Folder")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "height"
		if _, ok := i["height"]; ok {
			v := flattenZtnaWebPortalBookmarkBookmarksHeight(i["height"], d, pre_append)
			tmp["height"] = fortiAPISubPartPatch(v, "ZtnaWebPortalBookmark-Bookmarks-Height")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := i["host"]; ok {
			v := flattenZtnaWebPortalBookmarkBookmarksHost(i["host"], d, pre_append)
			tmp["host"] = fortiAPISubPartPatch(v, "ZtnaWebPortalBookmark-Bookmarks-Host")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keyboard_layout"
		if _, ok := i["keyboard-layout"]; ok {
			v := flattenZtnaWebPortalBookmarkBookmarksKeyboardLayout(i["keyboard-layout"], d, pre_append)
			tmp["keyboard_layout"] = fortiAPISubPartPatch(v, "ZtnaWebPortalBookmark-Bookmarks-KeyboardLayout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "load_balancing_info"
		if _, ok := i["load-balancing-info"]; ok {
			v := flattenZtnaWebPortalBookmarkBookmarksLoadBalancingInfo(i["load-balancing-info"], d, pre_append)
			tmp["load_balancing_info"] = fortiAPISubPartPatch(v, "ZtnaWebPortalBookmark-Bookmarks-LoadBalancingInfo")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_user"
		if _, ok := i["logon-user"]; ok {
			v := flattenZtnaWebPortalBookmarkBookmarksLogonUser(i["logon-user"], d, pre_append)
			tmp["logon_user"] = fortiAPISubPartPatch(v, "ZtnaWebPortalBookmark-Bookmarks-LogonUser")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenZtnaWebPortalBookmarkBookmarksName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ZtnaWebPortalBookmark-Bookmarks-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenZtnaWebPortalBookmarkBookmarksPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "ZtnaWebPortalBookmark-Bookmarks-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_blob"
		if _, ok := i["preconnection-blob"]; ok {
			v := flattenZtnaWebPortalBookmarkBookmarksPreconnectionBlob(i["preconnection-blob"], d, pre_append)
			tmp["preconnection_blob"] = fortiAPISubPartPatch(v, "ZtnaWebPortalBookmark-Bookmarks-PreconnectionBlob")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_id"
		if _, ok := i["preconnection-id"]; ok {
			v := flattenZtnaWebPortalBookmarkBookmarksPreconnectionId(i["preconnection-id"], d, pre_append)
			tmp["preconnection_id"] = fortiAPISubPartPatch(v, "ZtnaWebPortalBookmark-Bookmarks-PreconnectionId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restricted_admin"
		if _, ok := i["restricted-admin"]; ok {
			v := flattenZtnaWebPortalBookmarkBookmarksRestrictedAdmin(i["restricted-admin"], d, pre_append)
			tmp["restricted_admin"] = fortiAPISubPartPatch(v, "ZtnaWebPortalBookmark-Bookmarks-RestrictedAdmin")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security"
		if _, ok := i["security"]; ok {
			v := flattenZtnaWebPortalBookmarkBookmarksSecurity(i["security"], d, pre_append)
			tmp["security"] = fortiAPISubPartPatch(v, "ZtnaWebPortalBookmark-Bookmarks-Security")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_preconnection_id"
		if _, ok := i["send-preconnection-id"]; ok {
			v := flattenZtnaWebPortalBookmarkBookmarksSendPreconnectionId(i["send-preconnection-id"], d, pre_append)
			tmp["send_preconnection_id"] = fortiAPISubPartPatch(v, "ZtnaWebPortalBookmark-Bookmarks-SendPreconnectionId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso"
		if _, ok := i["sso"]; ok {
			v := flattenZtnaWebPortalBookmarkBookmarksSso(i["sso"], d, pre_append)
			tmp["sso"] = fortiAPISubPartPatch(v, "ZtnaWebPortalBookmark-Bookmarks-Sso")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := i["url"]; ok {
			v := flattenZtnaWebPortalBookmarkBookmarksUrl(i["url"], d, pre_append)
			tmp["url"] = fortiAPISubPartPatch(v, "ZtnaWebPortalBookmark-Bookmarks-Url")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vnc_keyboard_layout"
		if _, ok := i["vnc-keyboard-layout"]; ok {
			v := flattenZtnaWebPortalBookmarkBookmarksVncKeyboardLayout(i["vnc-keyboard-layout"], d, pre_append)
			tmp["vnc_keyboard_layout"] = fortiAPISubPartPatch(v, "ZtnaWebPortalBookmark-Bookmarks-VncKeyboardLayout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "width"
		if _, ok := i["width"]; ok {
			v := flattenZtnaWebPortalBookmarkBookmarksWidth(i["width"], d, pre_append)
			tmp["width"] = fortiAPISubPartPatch(v, "ZtnaWebPortalBookmark-Bookmarks-Width")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenZtnaWebPortalBookmarkBookmarksApptype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksColorDepth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksFolder(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksHeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksKeyboardLayout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksLoadBalancingInfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksLogonUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksPreconnectionBlob(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksPreconnectionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksRestrictedAdmin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksSecurity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksSendPreconnectionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksSso(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksVncKeyboardLayout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksWidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaWebPortalBookmarkName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkUsers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectZtnaWebPortalBookmark(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("bookmarks", flattenZtnaWebPortalBookmarkBookmarks(o["bookmarks"], d, "bookmarks")); err != nil {
			if vv, ok := fortiAPIPatch(o["bookmarks"], "ZtnaWebPortalBookmark-Bookmarks"); ok {
				if err = d.Set("bookmarks", vv); err != nil {
					return fmt.Errorf("Error reading bookmarks: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading bookmarks: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("bookmarks"); ok {
			if err = d.Set("bookmarks", flattenZtnaWebPortalBookmarkBookmarks(o["bookmarks"], d, "bookmarks")); err != nil {
				if vv, ok := fortiAPIPatch(o["bookmarks"], "ZtnaWebPortalBookmark-Bookmarks"); ok {
					if err = d.Set("bookmarks", vv); err != nil {
						return fmt.Errorf("Error reading bookmarks: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading bookmarks: %v", err)
				}
			}
		}
	}

	if err = d.Set("groups", flattenZtnaWebPortalBookmarkGroups(o["groups"], d, "groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["groups"], "ZtnaWebPortalBookmark-Groups"); ok {
			if err = d.Set("groups", vv); err != nil {
				return fmt.Errorf("Error reading groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading groups: %v", err)
		}
	}

	if err = d.Set("name", flattenZtnaWebPortalBookmarkName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ZtnaWebPortalBookmark-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("users", flattenZtnaWebPortalBookmarkUsers(o["users"], d, "users")); err != nil {
		if vv, ok := fortiAPIPatch(o["users"], "ZtnaWebPortalBookmark-Users"); ok {
			if err = d.Set("users", vv); err != nil {
				return fmt.Errorf("Error reading users: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading users: %v", err)
		}
	}

	return nil
}

func flattenZtnaWebPortalBookmarkFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandZtnaWebPortalBookmarkBookmarks(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apptype"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["apptype"], _ = expandZtnaWebPortalBookmarkBookmarksApptype(d, i["apptype"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color_depth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["color-depth"], _ = expandZtnaWebPortalBookmarkBookmarksColorDepth(d, i["color_depth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandZtnaWebPortalBookmarkBookmarksDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["domain"], _ = expandZtnaWebPortalBookmarkBookmarksDomain(d, i["domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "folder"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["folder"], _ = expandZtnaWebPortalBookmarkBookmarksFolder(d, i["folder"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "height"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["height"], _ = expandZtnaWebPortalBookmarkBookmarksHeight(d, i["height"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["host"], _ = expandZtnaWebPortalBookmarkBookmarksHost(d, i["host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keyboard_layout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["keyboard-layout"], _ = expandZtnaWebPortalBookmarkBookmarksKeyboardLayout(d, i["keyboard_layout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "load_balancing_info"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["load-balancing-info"], _ = expandZtnaWebPortalBookmarkBookmarksLoadBalancingInfo(d, i["load_balancing_info"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["logon-password"], _ = expandZtnaWebPortalBookmarkBookmarksLogonPassword(d, i["logon_password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_user"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["logon-user"], _ = expandZtnaWebPortalBookmarkBookmarksLogonUser(d, i["logon_user"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandZtnaWebPortalBookmarkBookmarksName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandZtnaWebPortalBookmarkBookmarksPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_blob"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preconnection-blob"], _ = expandZtnaWebPortalBookmarkBookmarksPreconnectionBlob(d, i["preconnection_blob"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preconnection-id"], _ = expandZtnaWebPortalBookmarkBookmarksPreconnectionId(d, i["preconnection_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restricted_admin"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["restricted-admin"], _ = expandZtnaWebPortalBookmarkBookmarksRestrictedAdmin(d, i["restricted_admin"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["security"], _ = expandZtnaWebPortalBookmarkBookmarksSecurity(d, i["security"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_preconnection_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["send-preconnection-id"], _ = expandZtnaWebPortalBookmarkBookmarksSendPreconnectionId(d, i["send_preconnection_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso"], _ = expandZtnaWebPortalBookmarkBookmarksSso(d, i["sso"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["url"], _ = expandZtnaWebPortalBookmarkBookmarksUrl(d, i["url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vnc_keyboard_layout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vnc-keyboard-layout"], _ = expandZtnaWebPortalBookmarkBookmarksVncKeyboardLayout(d, i["vnc_keyboard_layout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "width"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["width"], _ = expandZtnaWebPortalBookmarkBookmarksWidth(d, i["width"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandZtnaWebPortalBookmarkBookmarksApptype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksColorDepth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksFolder(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksHeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksKeyboardLayout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksLoadBalancingInfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksLogonPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaWebPortalBookmarkBookmarksLogonUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksPreconnectionBlob(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksPreconnectionId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksRestrictedAdmin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksSecurity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksSendPreconnectionId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksSso(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksVncKeyboardLayout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksWidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaWebPortalBookmarkName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkUsers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectZtnaWebPortalBookmark(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bookmarks"); ok || d.HasChange("bookmarks") {
		t, err := expandZtnaWebPortalBookmarkBookmarks(d, v, "bookmarks")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bookmarks"] = t
		}
	}

	if v, ok := d.GetOk("groups"); ok || d.HasChange("groups") {
		t, err := expandZtnaWebPortalBookmarkGroups(d, v, "groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["groups"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandZtnaWebPortalBookmarkName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("users"); ok || d.HasChange("users") {
		t, err := expandZtnaWebPortalBookmarkUsers(d, v, "users")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["users"] = t
		}
	}

	return &obj, nil
}
