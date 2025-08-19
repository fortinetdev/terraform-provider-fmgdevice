// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Proxy destination connection fast-fallback.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebProxyFastFallback() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebProxyFastFallbackCreate,
		Read:   resourceWebProxyFastFallbackRead,
		Update: resourceWebProxyFastFallbackUpdate,
		Delete: resourceWebProxyFastFallbackDelete,

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
			"connection_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connection_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWebProxyFastFallbackCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWebProxyFastFallback(d)
	if err != nil {
		return fmt.Errorf("Error creating WebProxyFastFallback resource while getting object: %v", err)
	}

	_, err = c.CreateWebProxyFastFallback(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WebProxyFastFallback resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWebProxyFastFallbackRead(d, m)
}

func resourceWebProxyFastFallbackUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWebProxyFastFallback(d)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyFastFallback resource while getting object: %v", err)
	}

	_, err = c.UpdateWebProxyFastFallback(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyFastFallback resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWebProxyFastFallbackRead(d, m)
}

func resourceWebProxyFastFallbackDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWebProxyFastFallback(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebProxyFastFallback resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyFastFallbackRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebProxyFastFallback(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyFastFallback resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebProxyFastFallback(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyFastFallback resource from API: %v", err)
	}
	return nil
}

func flattenWebProxyFastFallbackConnectionMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyFastFallbackConnectionTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyFastFallbackName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyFastFallbackProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyFastFallbackStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebProxyFastFallback(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("connection_mode", flattenWebProxyFastFallbackConnectionMode(o["connection-mode"], d, "connection_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["connection-mode"], "WebProxyFastFallback-ConnectionMode"); ok {
			if err = d.Set("connection_mode", vv); err != nil {
				return fmt.Errorf("Error reading connection_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading connection_mode: %v", err)
		}
	}

	if err = d.Set("connection_timeout", flattenWebProxyFastFallbackConnectionTimeout(o["connection-timeout"], d, "connection_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["connection-timeout"], "WebProxyFastFallback-ConnectionTimeout"); ok {
			if err = d.Set("connection_timeout", vv); err != nil {
				return fmt.Errorf("Error reading connection_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading connection_timeout: %v", err)
		}
	}

	if err = d.Set("name", flattenWebProxyFastFallbackName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WebProxyFastFallback-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("protocol", flattenWebProxyFastFallbackProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "WebProxyFastFallback-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("status", flattenWebProxyFastFallbackStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WebProxyFastFallback-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenWebProxyFastFallbackFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebProxyFastFallbackConnectionMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyFastFallbackConnectionTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyFastFallbackName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyFastFallbackProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyFastFallbackStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebProxyFastFallback(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("connection_mode"); ok || d.HasChange("connection_mode") {
		t, err := expandWebProxyFastFallbackConnectionMode(d, v, "connection_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connection-mode"] = t
		}
	}

	if v, ok := d.GetOk("connection_timeout"); ok || d.HasChange("connection_timeout") {
		t, err := expandWebProxyFastFallbackConnectionTimeout(d, v, "connection_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connection-timeout"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWebProxyFastFallbackName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandWebProxyFastFallbackProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWebProxyFastFallbackStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
