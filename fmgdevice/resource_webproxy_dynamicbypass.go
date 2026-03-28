// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device WebProxyDynamicBypass

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebProxyDynamicBypass() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebProxyDynamicBypassUpdate,
		Read:   resourceWebProxyDynamicBypassRead,
		Update: resourceWebProxyDynamicBypassUpdate,
		Delete: resourceWebProxyDynamicBypassDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

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
			"errors": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"server_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"total_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWebProxyDynamicBypassUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebProxyDynamicBypass(d)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyDynamicBypass resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWebProxyDynamicBypass(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyDynamicBypass resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WebProxyDynamicBypass")

	return resourceWebProxyDynamicBypassRead(d, m)
}

func resourceWebProxyDynamicBypassDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWebProxyDynamicBypass(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebProxyDynamicBypass resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyDynamicBypassRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebProxyDynamicBypass(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WebProxyDynamicBypass resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebProxyDynamicBypass(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyDynamicBypass resource from API: %v", err)
	}
	return nil
}

func flattenWebProxyDynamicBypassErrors(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyDynamicBypassServerMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyDynamicBypassStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyDynamicBypassTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyDynamicBypassTotalMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebProxyDynamicBypass(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("errors", flattenWebProxyDynamicBypassErrors(o["errors"], d, "errors")); err != nil {
		if vv, ok := fortiAPIPatch(o["errors"], "WebProxyDynamicBypass-Errors"); ok {
			if err = d.Set("errors", vv); err != nil {
				return fmt.Errorf("Error reading errors: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading errors: %v", err)
		}
	}

	if err = d.Set("server_max", flattenWebProxyDynamicBypassServerMax(o["server-max"], d, "server_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-max"], "WebProxyDynamicBypass-ServerMax"); ok {
			if err = d.Set("server_max", vv); err != nil {
				return fmt.Errorf("Error reading server_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_max: %v", err)
		}
	}

	if err = d.Set("status", flattenWebProxyDynamicBypassStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WebProxyDynamicBypass-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("timeout", flattenWebProxyDynamicBypassTimeout(o["timeout"], d, "timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["timeout"], "WebProxyDynamicBypass-Timeout"); ok {
			if err = d.Set("timeout", vv); err != nil {
				return fmt.Errorf("Error reading timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	if err = d.Set("total_max", flattenWebProxyDynamicBypassTotalMax(o["total-max"], d, "total_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["total-max"], "WebProxyDynamicBypass-TotalMax"); ok {
			if err = d.Set("total_max", vv); err != nil {
				return fmt.Errorf("Error reading total_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading total_max: %v", err)
		}
	}

	return nil
}

func flattenWebProxyDynamicBypassFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebProxyDynamicBypassErrors(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyDynamicBypassServerMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyDynamicBypassStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyDynamicBypassTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyDynamicBypassTotalMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebProxyDynamicBypass(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("errors"); ok || d.HasChange("errors") {
		t, err := expandWebProxyDynamicBypassErrors(d, v, "errors")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["errors"] = t
		}
	}

	if v, ok := d.GetOk("server_max"); ok || d.HasChange("server_max") {
		t, err := expandWebProxyDynamicBypassServerMax(d, v, "server_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-max"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWebProxyDynamicBypassStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("timeout"); ok || d.HasChange("timeout") {
		t, err := expandWebProxyDynamicBypassTimeout(d, v, "timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout"] = t
		}
	}

	if v, ok := d.GetOk("total_max"); ok || d.HasChange("total_max") {
		t, err := expandWebProxyDynamicBypassTotalMax(d, v, "total_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["total-max"] = t
		}
	}

	return &obj, nil
}
