// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure debug URL addresses.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebProxyDebugUrl() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebProxyDebugUrlCreate,
		Read:   resourceWebProxyDebugUrlRead,
		Update: resourceWebProxyDebugUrlUpdate,
		Delete: resourceWebProxyDebugUrlDelete,

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
			"exact": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url_pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceWebProxyDebugUrlCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectWebProxyDebugUrl(d)
	if err != nil {
		return fmt.Errorf("Error creating WebProxyDebugUrl resource while getting object: %v", err)
	}

	_, err = c.CreateWebProxyDebugUrl(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WebProxyDebugUrl resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWebProxyDebugUrlRead(d, m)
}

func resourceWebProxyDebugUrlUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectWebProxyDebugUrl(d)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyDebugUrl resource while getting object: %v", err)
	}

	_, err = c.UpdateWebProxyDebugUrl(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyDebugUrl resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWebProxyDebugUrlRead(d, m)
}

func resourceWebProxyDebugUrlDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	err = c.DeleteWebProxyDebugUrl(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WebProxyDebugUrl resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyDebugUrlRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebProxyDebugUrl(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyDebugUrl resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebProxyDebugUrl(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyDebugUrl resource from API: %v", err)
	}
	return nil
}

func flattenWebProxyDebugUrlExact(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyDebugUrlName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyDebugUrlStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyDebugUrlUrlPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebProxyDebugUrl(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("exact", flattenWebProxyDebugUrlExact(o["exact"], d, "exact")); err != nil {
		if vv, ok := fortiAPIPatch(o["exact"], "WebProxyDebugUrl-Exact"); ok {
			if err = d.Set("exact", vv); err != nil {
				return fmt.Errorf("Error reading exact: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exact: %v", err)
		}
	}

	if err = d.Set("name", flattenWebProxyDebugUrlName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WebProxyDebugUrl-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", flattenWebProxyDebugUrlStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WebProxyDebugUrl-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("url_pattern", flattenWebProxyDebugUrlUrlPattern(o["url-pattern"], d, "url_pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-pattern"], "WebProxyDebugUrl-UrlPattern"); ok {
			if err = d.Set("url_pattern", vv); err != nil {
				return fmt.Errorf("Error reading url_pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_pattern: %v", err)
		}
	}

	return nil
}

func flattenWebProxyDebugUrlFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebProxyDebugUrlExact(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyDebugUrlName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyDebugUrlStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyDebugUrlUrlPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebProxyDebugUrl(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("exact"); ok || d.HasChange("exact") {
		t, err := expandWebProxyDebugUrlExact(d, v, "exact")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exact"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWebProxyDebugUrlName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWebProxyDebugUrlStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("url_pattern"); ok || d.HasChange("url_pattern") {
		t, err := expandWebProxyDebugUrlUrlPattern(d, v, "url_pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-pattern"] = t
		}
	}

	return &obj, nil
}
