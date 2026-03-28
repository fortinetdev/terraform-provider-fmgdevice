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

func resourceIsolatorProfileEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceIsolatorProfileEntriesCreate,
		Read:   resourceIsolatorProfileEntriesRead,
		Update: resourceIsolatorProfileEntriesUpdate,
		Delete: resourceIsolatorProfileEntriesDelete,

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
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"copy_paste": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"proxy_address": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"right_click": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceIsolatorProfileEntriesCreate(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	obj, err := getObjectIsolatorProfileEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating IsolatorProfileEntries resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadIsolatorProfileEntries(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateIsolatorProfileEntries(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating IsolatorProfileEntries resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateIsolatorProfileEntries(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating IsolatorProfileEntries resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceIsolatorProfileEntriesRead(d, m)
}

func resourceIsolatorProfileEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	obj, err := getObjectIsolatorProfileEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating IsolatorProfileEntries resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateIsolatorProfileEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating IsolatorProfileEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceIsolatorProfileEntriesRead(d, m)
}

func resourceIsolatorProfileEntriesDelete(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	wsParams["adom"] = adomv

	err = c.DeleteIsolatorProfileEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting IsolatorProfileEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIsolatorProfileEntriesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	profile := d.Get("profile").(string)
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
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	o, err := c.ReadIsolatorProfileEntries(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading IsolatorProfileEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIsolatorProfileEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading IsolatorProfileEntries resource from API: %v", err)
	}
	return nil
}

func flattenIsolatorProfileEntriesAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIsolatorProfileEntriesCopyPaste2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIsolatorProfileEntriesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIsolatorProfileEntriesProxyAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIsolatorProfileEntriesRightClick2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIsolatorProfileEntriesStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectIsolatorProfileEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenIsolatorProfileEntriesAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "IsolatorProfileEntries-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("copy_paste", flattenIsolatorProfileEntriesCopyPaste2edl(o["copy-paste"], d, "copy_paste")); err != nil {
		if vv, ok := fortiAPIPatch(o["copy-paste"], "IsolatorProfileEntries-CopyPaste"); ok {
			if err = d.Set("copy_paste", vv); err != nil {
				return fmt.Errorf("Error reading copy_paste: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading copy_paste: %v", err)
		}
	}

	if err = d.Set("fosid", flattenIsolatorProfileEntriesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "IsolatorProfileEntries-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("proxy_address", flattenIsolatorProfileEntriesProxyAddress2edl(o["proxy-address"], d, "proxy_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-address"], "IsolatorProfileEntries-ProxyAddress"); ok {
			if err = d.Set("proxy_address", vv); err != nil {
				return fmt.Errorf("Error reading proxy_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_address: %v", err)
		}
	}

	if err = d.Set("right_click", flattenIsolatorProfileEntriesRightClick2edl(o["right-click"], d, "right_click")); err != nil {
		if vv, ok := fortiAPIPatch(o["right-click"], "IsolatorProfileEntries-RightClick"); ok {
			if err = d.Set("right_click", vv); err != nil {
				return fmt.Errorf("Error reading right_click: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading right_click: %v", err)
		}
	}

	if err = d.Set("status", flattenIsolatorProfileEntriesStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "IsolatorProfileEntries-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenIsolatorProfileEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandIsolatorProfileEntriesAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIsolatorProfileEntriesCopyPaste2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIsolatorProfileEntriesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIsolatorProfileEntriesProxyAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIsolatorProfileEntriesRightClick2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIsolatorProfileEntriesStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectIsolatorProfileEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandIsolatorProfileEntriesAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("copy_paste"); ok || d.HasChange("copy_paste") {
		t, err := expandIsolatorProfileEntriesCopyPaste2edl(d, v, "copy_paste")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["copy-paste"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandIsolatorProfileEntriesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("proxy_address"); ok || d.HasChange("proxy_address") {
		t, err := expandIsolatorProfileEntriesProxyAddress2edl(d, v, "proxy_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-address"] = t
		}
	}

	if v, ok := d.GetOk("right_click"); ok || d.HasChange("right_click") {
		t, err := expandIsolatorProfileEntriesRightClick2edl(d, v, "right_click")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["right-click"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandIsolatorProfileEntriesStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
