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

func resourceWebfilterUrlListEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterUrlListEntriesCreate,
		Read:   resourceWebfilterUrlListEntriesRead,
		Update: resourceWebfilterUrlListEntriesUpdate,
		Delete: resourceWebfilterUrlListEntriesDelete,

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
			"url_list": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceWebfilterUrlListEntriesCreate(d *schema.ResourceData, m interface{}) error {
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
	url_list := d.Get("url_list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["url_list"] = url_list

	obj, err := getObjectWebfilterUrlListEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating WebfilterUrlListEntries resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("url")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWebfilterUrlListEntries(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWebfilterUrlListEntries(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WebfilterUrlListEntries resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateWebfilterUrlListEntries(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WebfilterUrlListEntries resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "url"))

	return resourceWebfilterUrlListEntriesRead(d, m)
}

func resourceWebfilterUrlListEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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
	url_list := d.Get("url_list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["url_list"] = url_list

	obj, err := getObjectWebfilterUrlListEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterUrlListEntries resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWebfilterUrlListEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterUrlListEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "url"))

	return resourceWebfilterUrlListEntriesRead(d, m)
}

func resourceWebfilterUrlListEntriesDelete(d *schema.ResourceData, m interface{}) error {
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
	url_list := d.Get("url_list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["url_list"] = url_list

	wsParams["adom"] = adomv

	err = c.DeleteWebfilterUrlListEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterUrlListEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterUrlListEntriesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	url_list := d.Get("url_list").(string)
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
	if url_list == "" {
		url_list = importOptionChecking(m.(*FortiClient).Cfg, "url_list")
		if url_list == "" {
			return fmt.Errorf("Parameter url_list is missing")
		}
		if err = d.Set("url_list", url_list); err != nil {
			return fmt.Errorf("Error set params url_list: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["url_list"] = url_list

	o, err := c.ReadWebfilterUrlListEntries(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WebfilterUrlListEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterUrlListEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterUrlListEntries resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterUrlListEntriesUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebfilterUrlListEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("url", flattenWebfilterUrlListEntriesUrl2edl(o["url"], d, "url")); err != nil {
		if vv, ok := fortiAPIPatch(o["url"], "WebfilterUrlListEntries-Url"); ok {
			if err = d.Set("url", vv); err != nil {
				return fmt.Errorf("Error reading url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}

	return nil
}

func flattenWebfilterUrlListEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebfilterUrlListEntriesUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterUrlListEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("url"); ok || d.HasChange("url") {
		t, err := expandWebfilterUrlListEntriesUrl2edl(d, v, "url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}

	return &obj, nil
}
