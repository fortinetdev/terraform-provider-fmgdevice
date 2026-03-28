// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure YouTube API keys.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVideofilterYoutubeKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceVideofilterYoutubeKeyCreate,
		Read:   resourceVideofilterYoutubeKeyRead,
		Update: resourceVideofilterYoutubeKeyUpdate,
		Delete: resourceVideofilterYoutubeKeyDelete,

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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceVideofilterYoutubeKeyCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectVideofilterYoutubeKey(d)
	if err != nil {
		return fmt.Errorf("Error creating VideofilterYoutubeKey resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadVideofilterYoutubeKey(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateVideofilterYoutubeKey(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating VideofilterYoutubeKey resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateVideofilterYoutubeKey(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating VideofilterYoutubeKey resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceVideofilterYoutubeKeyRead(d, m)
			} else {
				return fmt.Errorf("Error creating VideofilterYoutubeKey resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceVideofilterYoutubeKeyRead(d, m)
}

func resourceVideofilterYoutubeKeyUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectVideofilterYoutubeKey(d)
	if err != nil {
		return fmt.Errorf("Error updating VideofilterYoutubeKey resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateVideofilterYoutubeKey(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VideofilterYoutubeKey resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceVideofilterYoutubeKeyRead(d, m)
}

func resourceVideofilterYoutubeKeyDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	wsParams["adom"] = adomv

	err = c.DeleteVideofilterYoutubeKey(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VideofilterYoutubeKey resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVideofilterYoutubeKeyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVideofilterYoutubeKey(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading VideofilterYoutubeKey resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVideofilterYoutubeKey(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VideofilterYoutubeKey resource from API: %v", err)
	}
	return nil
}

func flattenVideofilterYoutubeKeyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVideofilterYoutubeKeyKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVideofilterYoutubeKeyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVideofilterYoutubeKey(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenVideofilterYoutubeKeyId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "VideofilterYoutubeKey-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("key", flattenVideofilterYoutubeKeyKey(o["key"], d, "key")); err != nil {
		if vv, ok := fortiAPIPatch(o["key"], "VideofilterYoutubeKey-Key"); ok {
			if err = d.Set("key", vv); err != nil {
				return fmt.Errorf("Error reading key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading key: %v", err)
		}
	}

	if err = d.Set("status", flattenVideofilterYoutubeKeyStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "VideofilterYoutubeKey-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenVideofilterYoutubeKeyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVideofilterYoutubeKeyId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVideofilterYoutubeKeyKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVideofilterYoutubeKeyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVideofilterYoutubeKey(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandVideofilterYoutubeKeyId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("key"); ok || d.HasChange("key") {
		t, err := expandVideofilterYoutubeKeyKey(d, v, "key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandVideofilterYoutubeKeyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
