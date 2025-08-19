// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Form data.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnSslWebUserGroupBookmarkBookmarksFormData() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnSslWebUserGroupBookmarkBookmarksFormDataCreate,
		Read:   resourceVpnSslWebUserGroupBookmarkBookmarksFormDataRead,
		Update: resourceVpnSslWebUserGroupBookmarkBookmarksFormDataUpdate,
		Delete: resourceVpnSslWebUserGroupBookmarkBookmarksFormDataDelete,

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
			"bookmarks": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceVpnSslWebUserGroupBookmarkBookmarksFormDataCreate(d *schema.ResourceData, m interface{}) error {
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
	bookmarks := d.Get("bookmarks").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["user_group_bookmark"] = user_group_bookmark
	paradict["bookmarks"] = bookmarks

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectVpnSslWebUserGroupBookmarkBookmarksFormData(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnSslWebUserGroupBookmarkBookmarksFormData resource while getting object: %v", err)
	}

	_, err = c.CreateVpnSslWebUserGroupBookmarkBookmarksFormData(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating VpnSslWebUserGroupBookmarkBookmarksFormData resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceVpnSslWebUserGroupBookmarkBookmarksFormDataRead(d, m)
}

func resourceVpnSslWebUserGroupBookmarkBookmarksFormDataUpdate(d *schema.ResourceData, m interface{}) error {
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
	bookmarks := d.Get("bookmarks").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["user_group_bookmark"] = user_group_bookmark
	paradict["bookmarks"] = bookmarks

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectVpnSslWebUserGroupBookmarkBookmarksFormData(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebUserGroupBookmarkBookmarksFormData resource while getting object: %v", err)
	}

	_, err = c.UpdateVpnSslWebUserGroupBookmarkBookmarksFormData(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebUserGroupBookmarkBookmarksFormData resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceVpnSslWebUserGroupBookmarkBookmarksFormDataRead(d, m)
}

func resourceVpnSslWebUserGroupBookmarkBookmarksFormDataDelete(d *schema.ResourceData, m interface{}) error {
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
	bookmarks := d.Get("bookmarks").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["user_group_bookmark"] = user_group_bookmark
	paradict["bookmarks"] = bookmarks

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteVpnSslWebUserGroupBookmarkBookmarksFormData(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VpnSslWebUserGroupBookmarkBookmarksFormData resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnSslWebUserGroupBookmarkBookmarksFormDataRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	user_group_bookmark := d.Get("user_group_bookmark").(string)
	bookmarks := d.Get("bookmarks").(string)
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
	if bookmarks == "" {
		bookmarks = importOptionChecking(m.(*FortiClient).Cfg, "bookmarks")
		if bookmarks == "" {
			return fmt.Errorf("Parameter bookmarks is missing")
		}
		if err = d.Set("bookmarks", bookmarks); err != nil {
			return fmt.Errorf("Error set params bookmarks: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["user_group_bookmark"] = user_group_bookmark
	paradict["bookmarks"] = bookmarks

	o, err := c.ReadVpnSslWebUserGroupBookmarkBookmarksFormData(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebUserGroupBookmarkBookmarksFormData resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnSslWebUserGroupBookmarkBookmarksFormData(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebUserGroupBookmarkBookmarksFormData resource from API: %v", err)
	}
	return nil
}

func flattenVpnSslWebUserGroupBookmarkBookmarksFormDataName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksFormDataValue3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnSslWebUserGroupBookmarkBookmarksFormData(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenVpnSslWebUserGroupBookmarkBookmarksFormDataName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "VpnSslWebUserGroupBookmarkBookmarksFormData-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("value", flattenVpnSslWebUserGroupBookmarkBookmarksFormDataValue3rdl(o["value"], d, "value")); err != nil {
		if vv, ok := fortiAPIPatch(o["value"], "VpnSslWebUserGroupBookmarkBookmarksFormData-Value"); ok {
			if err = d.Set("value", vv); err != nil {
				return fmt.Errorf("Error reading value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	return nil
}

func flattenVpnSslWebUserGroupBookmarkBookmarksFormDataFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnSslWebUserGroupBookmarkBookmarksFormDataName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksFormDataValue3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnSslWebUserGroupBookmarkBookmarksFormData(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksFormDataName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok || d.HasChange("value") {
		t, err := expandVpnSslWebUserGroupBookmarkBookmarksFormDataValue3rdl(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
