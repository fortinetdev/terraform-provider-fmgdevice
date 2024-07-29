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

func resourceVpnSslWebUserBookmarkBookmarksFormData() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnSslWebUserBookmarkBookmarksFormDataCreate,
		Read:   resourceVpnSslWebUserBookmarkBookmarksFormDataRead,
		Update: resourceVpnSslWebUserBookmarkBookmarksFormDataUpdate,
		Delete: resourceVpnSslWebUserBookmarkBookmarksFormDataDelete,

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

func resourceVpnSslWebUserBookmarkBookmarksFormDataCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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
	bookmarks := d.Get("bookmarks").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["user_bookmark"] = user_bookmark
	paradict["bookmarks"] = bookmarks

	obj, err := getObjectVpnSslWebUserBookmarkBookmarksFormData(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnSslWebUserBookmarkBookmarksFormData resource while getting object: %v", err)
	}

	_, err = c.CreateVpnSslWebUserBookmarkBookmarksFormData(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating VpnSslWebUserBookmarkBookmarksFormData resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceVpnSslWebUserBookmarkBookmarksFormDataRead(d, m)
}

func resourceVpnSslWebUserBookmarkBookmarksFormDataUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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
	bookmarks := d.Get("bookmarks").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["user_bookmark"] = user_bookmark
	paradict["bookmarks"] = bookmarks

	obj, err := getObjectVpnSslWebUserBookmarkBookmarksFormData(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebUserBookmarkBookmarksFormData resource while getting object: %v", err)
	}

	_, err = c.UpdateVpnSslWebUserBookmarkBookmarksFormData(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebUserBookmarkBookmarksFormData resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceVpnSslWebUserBookmarkBookmarksFormDataRead(d, m)
}

func resourceVpnSslWebUserBookmarkBookmarksFormDataDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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
	bookmarks := d.Get("bookmarks").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["user_bookmark"] = user_bookmark
	paradict["bookmarks"] = bookmarks

	err = c.DeleteVpnSslWebUserBookmarkBookmarksFormData(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting VpnSslWebUserBookmarkBookmarksFormData resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnSslWebUserBookmarkBookmarksFormDataRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	user_bookmark := d.Get("user_bookmark").(string)
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
	if user_bookmark == "" {
		user_bookmark = importOptionChecking(m.(*FortiClient).Cfg, "user_bookmark")
		if user_bookmark == "" {
			return fmt.Errorf("Parameter user_bookmark is missing")
		}
		if err = d.Set("user_bookmark", user_bookmark); err != nil {
			return fmt.Errorf("Error set params user_bookmark: %v", err)
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
	paradict["user_bookmark"] = user_bookmark
	paradict["bookmarks"] = bookmarks

	o, err := c.ReadVpnSslWebUserBookmarkBookmarksFormData(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebUserBookmarkBookmarksFormData resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnSslWebUserBookmarkBookmarksFormData(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebUserBookmarkBookmarksFormData resource from API: %v", err)
	}
	return nil
}

func flattenVpnSslWebUserBookmarkBookmarksFormDataName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksFormDataValue3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnSslWebUserBookmarkBookmarksFormData(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenVpnSslWebUserBookmarkBookmarksFormDataName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "VpnSslWebUserBookmarkBookmarksFormData-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("value", flattenVpnSslWebUserBookmarkBookmarksFormDataValue3rdl(o["value"], d, "value")); err != nil {
		if vv, ok := fortiAPIPatch(o["value"], "VpnSslWebUserBookmarkBookmarksFormData-Value"); ok {
			if err = d.Set("value", vv); err != nil {
				return fmt.Errorf("Error reading value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	return nil
}

func flattenVpnSslWebUserBookmarkBookmarksFormDataFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnSslWebUserBookmarkBookmarksFormDataName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksFormDataValue3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnSslWebUserBookmarkBookmarksFormData(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandVpnSslWebUserBookmarkBookmarksFormDataName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok || d.HasChange("value") {
		t, err := expandVpnSslWebUserBookmarkBookmarksFormDataValue3rdl(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
