// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i>

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebProxyRedirectProfileEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebProxyRedirectProfileEntriesCreate,
		Read:   resourceWebProxyRedirectProfileEntriesRead,
		Update: resourceWebProxyRedirectProfileEntriesUpdate,
		Delete: resourceWebProxyRedirectProfileEntriesDelete,

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
			"redirect_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"redirect_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirect_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceWebProxyRedirectProfileEntriesCreate(d *schema.ResourceData, m interface{}) error {
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
	redirect_profile := d.Get("redirect_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["redirect_profile"] = redirect_profile

	obj, err := getObjectWebProxyRedirectProfileEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating WebProxyRedirectProfileEntries resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWebProxyRedirectProfileEntries(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWebProxyRedirectProfileEntries(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WebProxyRedirectProfileEntries resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateWebProxyRedirectProfileEntries(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WebProxyRedirectProfileEntries resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceWebProxyRedirectProfileEntriesRead(d, m)
			} else {
				return fmt.Errorf("Error creating WebProxyRedirectProfileEntries resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWebProxyRedirectProfileEntriesRead(d, m)
}

func resourceWebProxyRedirectProfileEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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
	redirect_profile := d.Get("redirect_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["redirect_profile"] = redirect_profile

	obj, err := getObjectWebProxyRedirectProfileEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyRedirectProfileEntries resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWebProxyRedirectProfileEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyRedirectProfileEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWebProxyRedirectProfileEntriesRead(d, m)
}

func resourceWebProxyRedirectProfileEntriesDelete(d *schema.ResourceData, m interface{}) error {
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
	redirect_profile := d.Get("redirect_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["redirect_profile"] = redirect_profile

	wsParams["adom"] = adomv

	err = c.DeleteWebProxyRedirectProfileEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebProxyRedirectProfileEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyRedirectProfileEntriesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	redirect_profile := d.Get("redirect_profile").(string)
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
	if redirect_profile == "" {
		redirect_profile = importOptionChecking(m.(*FortiClient).Cfg, "redirect_profile")
		if redirect_profile == "" {
			return fmt.Errorf("Parameter redirect_profile is missing")
		}
		if err = d.Set("redirect_profile", redirect_profile); err != nil {
			return fmt.Errorf("Error set params redirect_profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["redirect_profile"] = redirect_profile

	o, err := c.ReadWebProxyRedirectProfileEntries(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WebProxyRedirectProfileEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebProxyRedirectProfileEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyRedirectProfileEntries resource from API: %v", err)
	}
	return nil
}

func flattenWebProxyRedirectProfileEntriesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyRedirectProfileEntriesRedirectCode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyRedirectProfileEntriesRedirectUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyRedirectProfileEntriesType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyRedirectProfileEntriesUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebProxyRedirectProfileEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenWebProxyRedirectProfileEntriesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "WebProxyRedirectProfileEntries-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("redirect_code", flattenWebProxyRedirectProfileEntriesRedirectCode2edl(o["redirect-code"], d, "redirect_code")); err != nil {
		if vv, ok := fortiAPIPatch(o["redirect-code"], "WebProxyRedirectProfileEntries-RedirectCode"); ok {
			if err = d.Set("redirect_code", vv); err != nil {
				return fmt.Errorf("Error reading redirect_code: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redirect_code: %v", err)
		}
	}

	if err = d.Set("redirect_url", flattenWebProxyRedirectProfileEntriesRedirectUrl2edl(o["redirect-url"], d, "redirect_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["redirect-url"], "WebProxyRedirectProfileEntries-RedirectUrl"); ok {
			if err = d.Set("redirect_url", vv); err != nil {
				return fmt.Errorf("Error reading redirect_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redirect_url: %v", err)
		}
	}

	if err = d.Set("type", flattenWebProxyRedirectProfileEntriesType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "WebProxyRedirectProfileEntries-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("url", flattenWebProxyRedirectProfileEntriesUrl2edl(o["url"], d, "url")); err != nil {
		if vv, ok := fortiAPIPatch(o["url"], "WebProxyRedirectProfileEntries-Url"); ok {
			if err = d.Set("url", vv); err != nil {
				return fmt.Errorf("Error reading url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}

	return nil
}

func flattenWebProxyRedirectProfileEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebProxyRedirectProfileEntriesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyRedirectProfileEntriesRedirectCode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyRedirectProfileEntriesRedirectUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyRedirectProfileEntriesType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyRedirectProfileEntriesUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebProxyRedirectProfileEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandWebProxyRedirectProfileEntriesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("redirect_code"); ok || d.HasChange("redirect_code") {
		t, err := expandWebProxyRedirectProfileEntriesRedirectCode2edl(d, v, "redirect_code")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redirect-code"] = t
		}
	}

	if v, ok := d.GetOk("redirect_url"); ok || d.HasChange("redirect_url") {
		t, err := expandWebProxyRedirectProfileEntriesRedirectUrl2edl(d, v, "redirect_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redirect-url"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandWebProxyRedirectProfileEntriesType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("url"); ok || d.HasChange("url") {
		t, err := expandWebProxyRedirectProfileEntriesUrl2edl(d, v, "url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}

	return &obj, nil
}
