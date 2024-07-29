// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Exempt URLs from web proxy forwarding and caching.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebProxyUrlMatch() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebProxyUrlMatchCreate,
		Read:   resourceWebProxyUrlMatchRead,
		Update: resourceWebProxyUrlMatchUpdate,
		Delete: resourceWebProxyUrlMatchDelete,

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
			"cache_exemption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fast_fallback": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"forward_server": &schema.Schema{
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

func resourceWebProxyUrlMatchCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebProxyUrlMatch(d)
	if err != nil {
		return fmt.Errorf("Error creating WebProxyUrlMatch resource while getting object: %v", err)
	}

	_, err = c.CreateWebProxyUrlMatch(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WebProxyUrlMatch resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWebProxyUrlMatchRead(d, m)
}

func resourceWebProxyUrlMatchUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebProxyUrlMatch(d)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyUrlMatch resource while getting object: %v", err)
	}

	_, err = c.UpdateWebProxyUrlMatch(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyUrlMatch resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWebProxyUrlMatchRead(d, m)
}

func resourceWebProxyUrlMatchDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWebProxyUrlMatch(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WebProxyUrlMatch resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyUrlMatchRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebProxyUrlMatch(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyUrlMatch resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebProxyUrlMatch(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyUrlMatch resource from API: %v", err)
	}
	return nil
}

func flattenWebProxyUrlMatchCacheExemption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyUrlMatchComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyUrlMatchFastFallback(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyUrlMatchForwardServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyUrlMatchName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyUrlMatchStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyUrlMatchUrlPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebProxyUrlMatch(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("cache_exemption", flattenWebProxyUrlMatchCacheExemption(o["cache-exemption"], d, "cache_exemption")); err != nil {
		if vv, ok := fortiAPIPatch(o["cache-exemption"], "WebProxyUrlMatch-CacheExemption"); ok {
			if err = d.Set("cache_exemption", vv); err != nil {
				return fmt.Errorf("Error reading cache_exemption: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cache_exemption: %v", err)
		}
	}

	if err = d.Set("comment", flattenWebProxyUrlMatchComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "WebProxyUrlMatch-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("fast_fallback", flattenWebProxyUrlMatchFastFallback(o["fast-fallback"], d, "fast_fallback")); err != nil {
		if vv, ok := fortiAPIPatch(o["fast-fallback"], "WebProxyUrlMatch-FastFallback"); ok {
			if err = d.Set("fast_fallback", vv); err != nil {
				return fmt.Errorf("Error reading fast_fallback: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fast_fallback: %v", err)
		}
	}

	if err = d.Set("forward_server", flattenWebProxyUrlMatchForwardServer(o["forward-server"], d, "forward_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["forward-server"], "WebProxyUrlMatch-ForwardServer"); ok {
			if err = d.Set("forward_server", vv); err != nil {
				return fmt.Errorf("Error reading forward_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forward_server: %v", err)
		}
	}

	if err = d.Set("name", flattenWebProxyUrlMatchName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WebProxyUrlMatch-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", flattenWebProxyUrlMatchStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WebProxyUrlMatch-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("url_pattern", flattenWebProxyUrlMatchUrlPattern(o["url-pattern"], d, "url_pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-pattern"], "WebProxyUrlMatch-UrlPattern"); ok {
			if err = d.Set("url_pattern", vv); err != nil {
				return fmt.Errorf("Error reading url_pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_pattern: %v", err)
		}
	}

	return nil
}

func flattenWebProxyUrlMatchFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebProxyUrlMatchCacheExemption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyUrlMatchComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyUrlMatchFastFallback(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyUrlMatchForwardServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyUrlMatchName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyUrlMatchStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyUrlMatchUrlPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebProxyUrlMatch(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cache_exemption"); ok || d.HasChange("cache_exemption") {
		t, err := expandWebProxyUrlMatchCacheExemption(d, v, "cache_exemption")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-exemption"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandWebProxyUrlMatchComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("fast_fallback"); ok || d.HasChange("fast_fallback") {
		t, err := expandWebProxyUrlMatchFastFallback(d, v, "fast_fallback")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fast-fallback"] = t
		}
	}

	if v, ok := d.GetOk("forward_server"); ok || d.HasChange("forward_server") {
		t, err := expandWebProxyUrlMatchForwardServer(d, v, "forward_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward-server"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWebProxyUrlMatchName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWebProxyUrlMatchStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("url_pattern"); ok || d.HasChange("url_pattern") {
		t, err := expandWebProxyUrlMatchUrlPattern(d, v, "url_pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-pattern"] = t
		}
	}

	return &obj, nil
}
